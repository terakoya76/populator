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

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/terakoya76/populator/config"
	"github.com/terakoya76/populator/rand"
	"github.com/terakoya76/populator/utils"
)

// UnsignedableDataType accepts unsigned options
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

// IncrementableDataType accepts increment options
var IncrementableDataType = []interface{}{
	"tinyint",
	"smallint",
	"mediumint",
	"int",
	"bigint",
}

// OrderRequiredDataTypes require DataType(Order) like sql
var OrderRequiredDataTypes = []interface{}{
	"tinyint",
	"smallint",
	"mediumint",
	"int",
	"bigint",
	"bit",
	"year",
	"char",
	"varchar",
	"binary",
	"varbinary",
}

// PrecisionRequiredDataTypes require DataType(Order, Precision) like sql
var PrecisionRequiredDataTypes = []interface{}{
	"decimal",
	"float",
	"real",
	"double",
}

// ProhibitDefaultDataTypes must allow null value
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

// MySQLClient is an implementation of DBClient for MySQL
type MySQLClient struct {
	*sqlx.DB
}

// CreateTable does CreateTable statement for MySQL
func (db *MySQLClient) CreateTable(cfg *config.Table) error {
	sql := db.BuildCreateTableStmt(cfg)
	if _, err := db.Exec(sql); err != nil {
		return err
	}
	return nil
}

// BuildCreateTableStmt generate create_table_stmt sql for MySQL
func (db *MySQLClient) BuildCreateTableStmt(cfg *config.Table) string {
	var sb strings.Builder
	sb.WriteString(
		fmt.Sprintf(
			"CREATE TABLE IF NOT EXISTS %s (\n",
			cfg.Name,
		),
	)

	var regCol []string
	for _, column := range cfg.Columns {
		regCol = append(regCol, db.buildCreateTableStmtColumn(column))
	}
	sb.WriteString(strings.Join(regCol, ",\n"))

	if len(cfg.Indexes) > 0 {
		sb.WriteString(",\n")
	} else {
		sb.WriteString("\n")
	}

	var regIdx []string
	for _, index := range cfg.Indexes {
		regIdx = append(regIdx, db.buildCreateTableStmtIndex(index))
	}
	sb.WriteString(strings.Join(regIdx, ",\n"))

	if len(cfg.Indexes) > 0 {
		sb.WriteString("\n")
	}

	sb.WriteString(
		fmt.Sprintf(
			") DEFAULT CHARSET=%s",
			cfg.Charset,
		),
	)

	return sb.String()
}

func (db *MySQLClient) buildCreateTableStmtColumn(cfg *config.Column) string {
	var sb strings.Builder
	sb.WriteString("    ")
	sb.WriteString(cfg.Name)
	sb.WriteString(" ")
	sb.WriteString(cfg.Type)

	if utils.Contains(OrderRequiredDataTypes, cfg.Type) {
		sb.WriteString(
			fmt.Sprintf("(%d)", cfg.Order),
		)
	}
	if utils.Contains(PrecisionRequiredDataTypes, cfg.Type) {
		sb.WriteString(
			fmt.Sprintf("(%d, %d)", cfg.Order, cfg.Precision),
		)
	}

	if utils.Contains(UnsignedableDataType, cfg.Type) && cfg.Unsigned {
		sb.WriteString(" UNSIGNED")
	}
	if utils.Contains(IncrementableDataType, cfg.Type) && cfg.Increment {
		sb.WriteString(" AUTO_INCREMENT")
	}

	if !cfg.Nullable {
		sb.WriteString(" NOT NULL")
	}
	if !utils.Contains(ProhibitDefaultDataTypes, cfg.Type) && cfg.Default != nil {
		sb.WriteString(fmt.Sprintf(" DEFAULT(%v)", cfg.Default))
	}
	if cfg.Primary {
		sb.WriteString(" PRIMARY KEY")
	}

	return sb.String()
}

func (db *MySQLClient) buildCreateTableStmtIndex(cfg *config.Index) string {
	var sb strings.Builder
	sb.WriteString("    ")
	sb.WriteString("INDEX ")
	sb.WriteString(cfg.Name)
	sb.WriteString(" (")

	var reg []string
	for _, column := range cfg.Columns {
		reg = append(reg, fmt.Sprint(column))
	}
	sb.WriteString(strings.Join(reg, ", "))
	sb.WriteString(")")

	return sb.String()
}

// Populate does Insert statement for MySQL
func (db *MySQLClient) Populate(cfg *config.Table) error {
	err := utils.BatchTimes(
		db.buildInsertStmt(cfg),
		db.generateInsertRow(cfg),
		cfg.Record,
		100,
	)
	if err != nil {
		return err
	}

	return nil
}

func (db *MySQLClient) buildInsertStmt(cfg *config.Table) func([]string) error {
	return func(values []string) error {
		sql := db.BuildInsertStmt(cfg, values)
		fmt.Println(sql)

		if _, err := db.Exec(sql); err != nil {
			return err
		}
		return nil
	}
}

// BuildInsertStmt generate insert_stmt sql for MySQL
func (db *MySQLClient) BuildInsertStmt(cfg *config.Table, values []string) string {
	var sb strings.Builder

	var reg []string
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

func (db *MySQLClient) generateInsertRow(cfg *config.Table) func() string {
	return func() string {
		// generate insert values
		var reg []string
		for _, column := range cfg.Columns {
			value := db.generateValue(column)
			switch value := value.(type) {
			case string:
				reg = append(reg, fmt.Sprintf("   '%v'", value))
			default:
				reg = append(reg, fmt.Sprintf("   %v", value))

			}
		}
		return strings.Join(reg, ",\n")
	}
}

func (db *MySQLClient) generateValue(cfg *config.Column) interface{} {
	if cfg.Increment {
		return 0
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
			return rand.UnsignedDecimal(cfg.Order)
		}
		return rand.Decimal(cfg.Order)

	case "float":
		if cfg.Unsigned {
			return rand.UnsignedFloat(cfg.Order)
		}
		return rand.Float(cfg.Order)

	case "real":
		if cfg.Unsigned {
			return rand.UnsignedReal(cfg.Order)
		}
		return rand.Real(cfg.Order)

	case "double":
		if cfg.Unsigned {
			return rand.UnsignedDouble(cfg.Order)
		}
		return rand.Double(cfg.Order)

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
		return rand.TinyBlob(1000)

	case "tinytext":
		return rand.TinyText(1000)

	case "blob":
		return rand.Blob(3000)

	case "text":
		return rand.Text(3000)

	case "mediumblob":
		return rand.MediumBlob(7000)

	case "mediumtext":
		return rand.MediumText(7000)

	case "longblob":
		return rand.LongBlob(10000)

	case "longtext":
		return rand.LongText(10000)

	default:
		return rand.Boolean()
	}
}
