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
	"github.com/terakoya76/populator/utils"
)

// NaturalNumberDataType accepts unsigned/increment options
var NaturalNumberDataType = []string{
	"tinyint",
	"smallint",
	"mediumint",
	"int",
	"bigint",
}

// OrderRequiredDataTypes require DataType(Order) like sql
var OrderRequiredDataTypes = []string{
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
var PrecisionRequiredDataTypes = []string{
	"decimal",
	"float",
	"real",
	"double",
}

// MustNullableDataTypes must allow null value
var MustNullableDataTypes = []string{
	"blob",
	"text",
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

	if utils.Contains(NaturalNumberDataType, cfg.Type) && cfg.Unsigned {
		sb.WriteString(" UNSIGNED")
	}
	if utils.Contains(NaturalNumberDataType, cfg.Type) && cfg.Increment {
		sb.WriteString(" AUTO_INCREMENT")
	}

	if !utils.Contains(MustNullableDataTypes, cfg.Type) && !cfg.Nullable {
		sb.WriteString(" NOT NULL")
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
	fmt.Printf("name: %+v\n", cfg.Name)
	fmt.Printf("columns: %+v\n", cfg.Columns)
	fmt.Printf("indexes: %+v\n", cfg.Indexes)
	fmt.Printf("record: %+v\n", cfg.Record)

	return nil
}
