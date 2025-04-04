/*
Package database ...

Copyright © 2019 hajime-terasawa <terako.studio@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package database

import (
	"fmt"
	"strings"
	"sync"

	// MySQL Driver.
	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"

	"github.com/terakoya76/populator/config"
	"github.com/terakoya76/populator/rand"
	"github.com/terakoya76/populator/utils"
)

const (
	tinyBlobSize   int = 255
	tinyTextSize   int = 255
	blobSize       int = 1000
	textSize       int = 1000
	mediumBlobSize int = 3000
	mediumTextSize int = 3000
	longBlobSize   int = 5000
	longTextSize   int = 5000
)

// MaxConnections holds max_connections var for memory use control.
var MaxConnections int

// MySQLClient is an implementation of DBClient for MySQL.
type MySQLClient struct {
	*sqlx.DB
}

// SetupMySQLDB find_or_create database w/ given database name, then connect it.
func SetupMySQLDB(cfg *config.Database) error {
	ci := buildConnectInfo(cfg)
	db, err := sqlx.Open("mysql", ci)

	if err != nil {
		return fmt.Errorf("failed to setup database %s on mysql: %+v", cfg.Name, err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + cfg.Name)
	if err != nil {
		return fmt.Errorf("failed to setup database %s on mysql: %+v", cfg.Name, err)
	}

	db.Close()

	return nil
}

// BuildMySQLClient returns MySQLClient.
func BuildMySQLClient(cfg *config.Database) (*MySQLClient, error) {
	ci := buildConnectInfo(cfg)
	db, err := sqlx.Connect("mysql", ci+cfg.Name)

	if err != nil {
		return nil, fmt.Errorf("failed to setup database %s on mysql: %+v", cfg.Name, err)
	}

	var _name string
	err = db.QueryRow("show variables like \"%max_connections%\"").Scan(&_name, &MaxConnections)

	if err != nil {
		fmt.Println(err)
	}

	return &MySQLClient{db}, nil
}

func buildConnectInfo(cfg *config.Database) string {
	return cfg.User + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + fmt.Sprint(cfg.Port) + ")/"
}

// Verbose displays sql from cobra.
var Verbose bool

// UnsignedableDataType accepts unsigned options.
var UnsignedableDataType = []interface{}{
	"tinyint",
	"smallint",
	"mediumint",
	"int",
	"bigint",
	"decimal",
	"float",
	"real",
	"double",
}

// IncrementableDataType accepts increment options.
var IncrementableDataType = []interface{}{
	"tinyint",
	"smallint",
	"mediumint",
	"int",
	"bigint",
}

// OrderRequiredDataTypes require DataType(Order) like sql.
var OrderRequiredDataTypes = []interface{}{
	"tinyint",
	"smallint",
	"mediumint",
	"int",
	"bigint",
	"bit",
	"blob",
	"text",
	"year",
	"char",
	"varchar",
	"binary",
	"varbinary",
}

// PrecisionRequiredDataTypes require DataType(Order, Precision) like sql.
var PrecisionRequiredDataTypes = []interface{}{
	"decimal",
	"float",
	"real",
	"double",
}

// ProhibitDefaultDataTypes must allow null value.
var ProhibitDefaultDataTypes = []interface{}{
	"tinyblob",
	"tinytext",
	"blob",
	"text",
	"mediumblob",
	"mediumtext",
	"longblob",
	"longtext",
}

// CreateTable does CreateTable statement for MySQL.
func (db *MySQLClient) CreateTable(cfg *config.Table) error {
	sql := db.BuildCreateTableStmt(cfg)

	if Verbose {
		fmt.Println(sql)
	}

	if _, err := db.Exec(sql); err != nil {
		return err
	}

	return nil
}

// BuildCreateTableStmt generate create_table_stmt sql for MySQL.
func (db *MySQLClient) BuildCreateTableStmt(cfg *config.Table) string {
	var sb strings.Builder

	sb.WriteString(
		fmt.Sprintf(
			"CREATE TABLE IF NOT EXISTS %s (\n",
			cfg.Name,
		),
	)

	regCol := make([]string, 0, len(cfg.Columns))
	for _, column := range cfg.Columns {
		regCol = append(regCol, db.buildCreateTableStmtColumn(column))
	}

	sb.WriteString(strings.Join(regCol, ",\n"))

	if len(cfg.Indexes) > 0 {
		sb.WriteString(",\n")
	}

	regIdx := make([]string, 0, len(cfg.Indexes))
	for _, index := range cfg.Indexes {
		regIdx = append(regIdx, db.BuildIndexDesc(index))
	}

	sb.WriteString(strings.Join(regIdx, ",\n"))

	sb.WriteString(
		fmt.Sprintf(
			"\n) DEFAULT CHARSET=%s",
			cfg.Charset,
		),
	)

	return sb.String()
}

func (db *MySQLClient) buildCreateTableStmtColumn(cfg *config.Column) string {
	var sb strings.Builder

	sb.WriteString(
		fmt.Sprintf(
			"    %s %s",
			cfg.Name,
			cfg.Type,
		),
	)

	if utils.Contains(OrderRequiredDataTypes, cfg.Type) {
		sb.WriteString(fmt.Sprintf("(%d)", cfg.Order))
	}

	if utils.Contains(PrecisionRequiredDataTypes, cfg.Type) {
		sb.WriteString(fmt.Sprintf("(%d, %d)", cfg.Order, cfg.Precision))
	}

	if utils.Contains(UnsignedableDataType, cfg.Type) && cfg.Unsigned {
		sb.WriteString(" UNSIGNED")
	}

	if utils.Contains(IncrementableDataType, cfg.Type) && cfg.AutoIncrement {
		sb.WriteString(" AUTO_INCREMENT")
	}

	if cfg.NotNull {
		sb.WriteString(" NOT NULL")
	}

	if !utils.Contains(ProhibitDefaultDataTypes, cfg.Type) && cfg.Default != nil {
		sb.WriteString(db.BuildDefaultDesc(cfg))
	}

	if cfg.Primary {
		sb.WriteString(" PRIMARY KEY")
	}

	return sb.String()
}

// BuildDefaultDesc generate a default desc part of sql for MySQL.
func (db *MySQLClient) BuildDefaultDesc(cfg *config.Column) string {
	switch cfg.Default.(type) {
	case string:
		return fmt.Sprintf(" DEFAULT %q", cfg.Default)
	default:
		return fmt.Sprintf(" DEFAULT(%v)", cfg.Default)
	}
}

// BuildIndexDesc generate an index desc part of sql for MySQL.
func (db *MySQLClient) BuildIndexDesc(cfg *config.Index) string {
	var sb strings.Builder

	if cfg.Primary {
		sb.WriteString("    PRIMARY KEY ")
	} else if cfg.Uniq {
		sb.WriteString("    UNIQUE ")
	} else {
		sb.WriteString("    INDEX ")
	}

	if cfg.Name != "" {
		sb.WriteString(cfg.Name)
	}

	sb.WriteString(" (")
	sb.WriteString(strings.Join(cfg.Columns, ", "))
	sb.WriteString(")")

	return sb.String()
}

// DropTable does DropTable statement for MySQL.
func (db *MySQLClient) DropTable(cfg *config.Table) error {
	sql := db.BuildDropTableStmt(cfg)

	if Verbose {
		fmt.Println(sql)
	}

	if _, err := db.Exec(sql); err != nil {
		return err
	}

	return nil
}

// BuildDropTableStmt generate drop_table_stmt sql for MySQL.
func (db *MySQLClient) BuildDropTableStmt(cfg *config.Table) string {
	return fmt.Sprintf(
		"DROP TABLE IF EXISTS %s",
		cfg.Name,
	)
}

// Populate does Insert statement for MySQL.
func (db *MySQLClient) Populate(cfg *config.Table) error {
	var wg sync.WaitGroup

	otherConnections := 100
	batchSize := 200

	if cfg.Record < batchSize {
		batchSize = cfg.Record
	}

	i := 0
	for i < cfg.Record {
		// Not try to exec query
		// it would return "Error 1040: Too many connections"
		var currentConnections int
		err := db.QueryRow("select count(*) from information_schema.PROCESSLIST").Scan(&currentConnections)

		if err != nil {
			fmt.Println(err)

			currentConnections = MaxConnections
		}

		if currentConnections+otherConnections < MaxConnections {
			wg.Add(1)

			rows := make([]string, batchSize)
			for j := 0; j < batchSize; j++ {
				rows[j] = db.generateInsertRow(cfg)
			}

			go func() {
				if err := db.execInsertStmt(cfg, rows); err != nil {
					fmt.Println(err)
				}

				wg.Done()
			}()

			i += batchSize
		}
	}

	wg.Wait()

	return nil
}

func (db *MySQLClient) execInsertStmt(cfg *config.Table, values []string) error {
	sql := db.BuildInsertStmt(cfg, values)

	if Verbose {
		fmt.Println(sql)
	}

	if _, err := db.Exec(sql); err != nil {
		return err
	}

	return nil
}

// BuildInsertStmt generate insert_stmt sql for MySQL.
func (db *MySQLClient) BuildInsertStmt(cfg *config.Table, values []string) string {
	var sb strings.Builder

	reg := make([]string, 0, len(cfg.Columns))
	for _, column := range cfg.Columns {
		reg = append(reg, column.Name)
	}

	sb.WriteString(
		fmt.Sprintf(
			"INSERT INTO %s (%s) VALUES (\n",
			cfg.Name,
			strings.Join(reg, ", "),
		),
	)

	sb.WriteString(strings.Join(values, "\n), (\n"))
	sb.WriteString("\n)")

	return sb.String()
}

func (db *MySQLClient) generateInsertRow(cfg *config.Table) string {
	// generate insert values
	reg := make([]string, 0, len(cfg.Columns))

	for _, column := range cfg.Columns {
		value := db.generateValue(column)
		switch value := value.(type) {
		case string:
			reg = append(reg, fmt.Sprintf("   '%v'", value))
		case float32, float64:
			var sb strings.Builder

			sb.WriteString("   %.")
			sb.WriteString(fmt.Sprintf("%d", column.Precision))
			sb.WriteString("f")
			reg = append(reg, fmt.Sprintf(sb.String(), value))
		default:
			reg = append(reg, fmt.Sprintf("   %v", value))
		}
	}

	return strings.Join(reg, ",\n")
}

//nolint:gocyclo,funlen
func (db *MySQLClient) generateValue(cfg *config.Column) interface{} {
	if cfg.AutoIncrement {
		return 0
	}

	if len(cfg.Values) > 0 {
		return utils.Shuffle(cfg.Values)
	}

	switch cfg.Type {
	case "boolean":
		return rand.Boolean()

	case "tinyint":
		if cfg.Unsigned {
			return rand.UnsignedTinyInt()
		}

		return rand.TinyInt()

	case "smallint":
		if cfg.Unsigned {
			return rand.UnsignedSmallInt()
		}

		return rand.SmallInt()

	case "mediumint":
		if cfg.Unsigned {
			return rand.UnsignedMediumInt()
		}

		return rand.MediumInt()

	case "int":
		if cfg.Unsigned {
			return rand.UnsignedInt()
		}

		return rand.Int()

	case "bigint":
		if cfg.Unsigned {
			return rand.UnsignedBigInt()
		}

		return rand.BigInt()

	case "decimal":
		if cfg.Unsigned {
			return rand.UnsignedDecimal(cfg.Order, cfg.Precision)
		}

		return rand.Decimal(cfg.Order, cfg.Precision)

	case "float":
		if cfg.Unsigned {
			return rand.UnsignedFloat(cfg.Order, cfg.Precision)
		}

		return rand.Float(cfg.Order, cfg.Precision)

	case "real":
		if cfg.Unsigned {
			return rand.UnsignedReal(cfg.Order, cfg.Precision)
		}

		return rand.Real(cfg.Order, cfg.Precision)

	case "double":
		if cfg.Unsigned {
			return rand.UnsignedDouble(cfg.Order, cfg.Precision)
		}

		return rand.Double(cfg.Order, cfg.Precision)

	case "bit":
		return rand.Bit(cfg.Order)

	case "date":
		return rand.Date()

	case "datetime":
		return rand.DateTime()

	case "timestamp":
		return rand.Timestamp()

	case "time":
		return rand.Time()

	case "year":
		//nolint:mnd
		if cfg.Order == 4 {
			return rand.Year4()
		}

		return rand.Year2()

	case "char":
		return rand.Char(cfg.Order)

	case "varchar":
		return rand.VarChar(cfg.Order)

	case "binary":
		return rand.Binary(cfg.Order)

	case "varbinary":
		return rand.VarBinary(cfg.Order)

	case "tinyblob":
		return rand.TinyBlob(tinyBlobSize)

	case "tinytext":
		return rand.TinyText(tinyTextSize)

	case "blob":
		return rand.Blob(blobSize)

	case "text":
		return rand.Text(textSize)

	case "mediumblob":
		return rand.MediumBlob(mediumBlobSize)

	case "mediumtext":
		return rand.MediumText(mediumTextSize)

	case "longblob":
		return rand.LongBlob(longBlobSize)

	case "longtext":
		return rand.LongText(longTextSize)

	default:
		return 0
	}
}
