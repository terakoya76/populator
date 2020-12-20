/*
Package cmd ...
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
package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/terakoya76/populator/config"
	"github.com/terakoya76/populator/database"
)

func Test_BuildCreateTableStmt_Columns(t *testing.T) {
	var nilValues []interface{}

	cases := []struct {
		name string
		cfg  *config.Table
		sql  string
		err  error
	}{
		{
			name: "boolean",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "boolean",
						Order:         0,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 boolean\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "boolean w/ all option",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "boolean",
						Order:         0,
						Precision:     0,
						Unsigned:      true,
						NotNull:       true,
						Default:       true,
						Primary:       true,
						AutoIncrement: true,
						Values: []interface{}{
							true,
						},
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 boolean NOT NULL DEFAULT(true) PRIMARY KEY\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "tinyint",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "tinyint",
						Order:         2,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 tinyint(2)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "smallint",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "smallint",
						Order:         4,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 smallint(4)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "mediumint",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "mediumint",
						Order:         6,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 mediumint(6)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "int",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "int",
						Order:         9,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 int(9)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "bigint",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "bigint",
						Order:         11,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 bigint(11)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "numeric w/ all option",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "bigint",
						Order:         20,
						Precision:     0,
						Unsigned:      true,
						NotNull:       true,
						Default:       1000,
						Primary:       true,
						AutoIncrement: true,
						Values: []interface{}{
							1,
						},
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			// nolint:lll
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 bigint(20) UNSIGNED AUTO_INCREMENT NOT NULL DEFAULT(1000) PRIMARY KEY\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "decimal",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "decimal",
						Order:         5,
						Precision:     2,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 decimal(5, 2)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "float",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "float",
						Order:         10,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 float(10, 0)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "real",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "real",
						Order:         5,
						Precision:     2,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 real(5, 2)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "double",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "double",
						Order:         5,
						Precision:     2,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 double(5, 2)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "float w/ all options",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "float",
						Order:         5,
						Precision:     2,
						Unsigned:      true,
						NotNull:       true,
						Default:       123.45,
						Primary:       true,
						AutoIncrement: true,
						Values: []interface{}{
							123.45,
						},
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			// nolint:lll
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 float(5, 2) UNSIGNED NOT NULL DEFAULT(123.45) PRIMARY KEY\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "bit",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "bit",
						Order:         8,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 bit(8)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "bit w/ all options",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "bit",
						Order:         1,
						Precision:     0,
						Unsigned:      true,
						NotNull:       true,
						Default:       "b'01010101'",
						Primary:       true,
						AutoIncrement: true,
						Values: []interface{}{
							"b'01010101'",
						},
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 bit(1) NOT NULL DEFAULT \"b'01010101'\" PRIMARY KEY\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "date",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "date",
						Order:         0,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 date\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "datetime",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "datetime",
						Order:         0,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 datetime\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "timestamp",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "timestamp",
						Order:         0,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 timestamp\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "time",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "time",
						Order:         0,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 time\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "year",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "year",
						Order:         4,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 year(4)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "date w/ all options",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "date",
						Order:         0,
						Precision:     0,
						Unsigned:      true,
						NotNull:       true,
						Default:       "2000-12-01",
						Primary:       true,
						AutoIncrement: true,
						Values: []interface{}{
							"2000-12-01",
						},
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 date NOT NULL DEFAULT \"2000-12-01\" PRIMARY KEY\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "char",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "char",
						Order:         20,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 char(20)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "varchar",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "varchar",
						Order:         20,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 varchar(20)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "binary",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "binary",
						Order:         20,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 binary(20)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "varbinary",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "varbinary",
						Order:         20,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 varbinary(20)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "tinyblob",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "tinyblob",
						Order:         0,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 tinyblob\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "tinytext",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "tinytext",
						Order:         0,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 tinytext\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "blob",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "blob",
						Order:         100,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 blob(100)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "text",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "text",
						Order:         100,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 text(100)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "mediumblob",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "mediumblob",
						Order:         65535,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 mediumblob\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "mediumtext",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "mediumtext",
						Order:         65535,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 mediumtext\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "longblob",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "longblob",
						Order:         65535,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 longblob\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "longtext",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "longtext",
						Order:         65535,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
						Values:        nilValues,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 longtext\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "text w/ all options",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "text",
						Order:         65535,
						Precision:     0,
						Unsigned:      true,
						NotNull:       true,
						Default:       "hoge",
						Primary:       true,
						AutoIncrement: true,
						Values: []interface{}{
							"hoge",
						},
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 text(65535) NOT NULL PRIMARY KEY\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},
	}

	for _, c := range cases {
		client := database.MySQLClient{}
		sql := client.BuildCreateTableStmt(c.cfg)
		if !assert.Equal(t, c.sql, sql) {
			t.Errorf("case: %s is failed, expected: %+v, actual: %+v\n", c.name, c.sql, sql)
		}
	}
}

func Test_BuildDefaultDesc(t *testing.T) {
	cases := []struct {
		name   string
		cfg    *config.Column
		result string
		err    error
	}{
		{
			name: "string",
			cfg: &config.Column{
				Name:          "col_1",
				Type:          "varchar",
				Order:         0,
				Precision:     0,
				Unsigned:      false,
				NotNull:       false,
				Default:       "hoge",
				Primary:       false,
				AutoIncrement: false,
			},
			result: " DEFAULT \"hoge\"",
			err:    nil,
		},

		{
			name: "other",
			cfg: &config.Column{
				Name:          "col_1",
				Type:          "decimal",
				Order:         6,
				Precision:     3,
				Unsigned:      false,
				NotNull:       false,
				Default:       123.456,
				Primary:       false,
				AutoIncrement: false,
			},
			result: " DEFAULT(123.456)",
			err:    nil,
		},
	}

	for _, c := range cases {
		client := database.MySQLClient{}
		result := client.BuildDefaultDesc(c.cfg)
		if !assert.Equal(t, c.result, result) {
			t.Errorf("case: %s is failed, expected: %+v, actual: %+v\n", c.name, c.result, result)
		}
	}
}

func Test_BuildIndexDesc(t *testing.T) {
	cases := []struct {
		name   string
		cfg    *config.Index
		result string
		err    error
	}{
		{
			name: "normal",
			cfg: &config.Index{
				Name:    "idx_1",
				Primary: false,
				Uniq:    false,
				Columns: []string{
					"col_1",
				},
			},
			result: "    INDEX idx_1 (col_1)",
			err:    nil,
		},

		{
			name: "covering",
			cfg: &config.Index{
				Name:    "idx_1",
				Primary: false,
				Uniq:    false,
				Columns: []string{
					"col_1",
					"col_2",
				},
			},
			result: "    INDEX idx_1 (col_1, col_2)",
			err:    nil,
		},

		{
			name: "primary key",
			cfg: &config.Index{
				Name:    "",
				Primary: true,
				Uniq:    false,
				Columns: []string{
					"col_1",
				},
			},
			result: "    PRIMARY KEY  (col_1)",
			err:    nil,
		},

		{
			name: "covering primary key",
			cfg: &config.Index{
				Name:    "",
				Primary: true,
				Uniq:    false,
				Columns: []string{
					"col_1",
					"col_2",
				},
			},
			result: "    PRIMARY KEY  (col_1, col_2)",
			err:    nil,
		},

		{
			name: "uniq key",
			cfg: &config.Index{
				Name:    "idx_1",
				Primary: false,
				Uniq:    true,
				Columns: []string{
					"col_1",
				},
			},
			result: "    UNIQUE idx_1 (col_1)",
			err:    nil,
		},

		{
			name: "covering uniq key",
			cfg: &config.Index{
				Name:    "idx_1",
				Primary: false,
				Uniq:    true,
				Columns: []string{
					"col_1",
					"col_2",
				},
			},
			result: "    UNIQUE idx_1 (col_1, col_2)",
			err:    nil,
		},
	}

	for _, c := range cases {
		client := database.MySQLClient{}
		result := client.BuildIndexDesc(c.cfg)
		if !assert.Equal(t, c.result, result) {
			t.Errorf("case: %s is failed, expected: %+v, actual: %+v\n", c.name, c.result, result)
		}
	}
}

func Test_BuildDropTableStmt(t *testing.T) {
	cases := []struct {
		name string
		cfg  *config.Table
		sql  string
		err  error
	}{
		{
			name: "boolean",
			cfg: &config.Table{
				Name: "table_a",
				Columns: []*config.Column{
					{
						Name:          "col_1",
						Type:          "boolean",
						Order:         0,
						Precision:     0,
						Unsigned:      false,
						NotNull:       false,
						Default:       nil,
						Primary:       false,
						AutoIncrement: false,
					},
				},
				Indexes: []*config.Index{},
				Charset: "utf8mb4",
				Record:  100000,
			},
			sql: "DROP TABLE IF EXISTS table_a",
			err: nil,
		},
	}

	for _, c := range cases {
		client := database.MySQLClient{}
		sql := client.BuildDropTableStmt(c.cfg)
		if !assert.Equal(t, c.sql, sql) {
			t.Errorf("case: %s is failed, expected: %+v, actual: %+v\n", c.name, c.sql, sql)
		}
	}
}
