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
	"bytes"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/terakoya76/populator/cmd"
	"github.com/terakoya76/populator/config"
	"github.com/terakoya76/populator/database"
)

func Test_BuildCreateTableStmt_Columns(t *testing.T) {
	viper.SetConfigType("yaml")

	cases := []struct {
		name string
		yaml []byte
		sql  string
		err  error
	}{
		{
			name: "boolean",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: boolean
                      order: 0
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 boolean\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "boolean w/ all option",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: boolean
                      order: 0
                      precision: 0
                      unsigned: true
                      notNull: true
                      default: true
                      primary: true
                      autoIncrement: true
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 boolean NOT NULL DEFAULT(true) PRIMARY KEY\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "tinyint",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: tinyint
                      order: 4
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 tinyint(4)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "smallint(6)",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: smallint
                      order: 6
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 smallint(6)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "mediumint(9)",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: mediumint
                      order: 9
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 mediumint(9)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "int(11)",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: int
                      order: 11
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 int(11)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "bigint(20)",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: bigint
                      order: 20
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 bigint(20)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "numeric w/ all option",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: bigint
                      order: 20
                      precision: 0
                      unsigned: true
                      notNull: true
                      default: 1000
                      primary: true
                      autoIncrement: true
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 bigint(20) UNSIGNED AUTO_INCREMENT NOT NULL DEFAULT(1000) PRIMARY KEY\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "decimal(5,2)",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: decimal
                      order: 5
                      precision: 2
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 decimal(5, 2)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "float(5,2)",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: float
                      order: 5
                      precision: 2
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 float(5, 2)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "real(5,2)",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: real
                      order: 5
                      precision: 2
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 real(5, 2)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "double(5,2)",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: double
                      order: 5
                      precision: 2
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 double(5, 2)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "float w/ all options",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: float
                      order: 5
                      precision: 2
                      unsigned: true
                      notNull: true
                      default: 123.45
                      primary: true
                      autoIncrement: true
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 float(5, 2) UNSIGNED NOT NULL DEFAULT(123.45) PRIMARY KEY\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "bit(8)",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: bit
                      order: 8
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 bit(8)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "bit w/ all options",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: bit
                      order: 8
                      precision: 0
                      unsigned: true
                      notNull: true
                      default: b'01010101'
                      primary: true
                      autoIncrement: true
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 bit(8) NOT NULL DEFAULT \"b'01010101'\" PRIMARY KEY\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "date",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: date
                      order: 0
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 date\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "datetime",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: datetime
                      order: 0
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 datetime\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "timestamp",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: timestamp
                      order: 0
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 timestamp\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "time",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: time
                      order: 0
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 time\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "year",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: year
                      order: 4
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 year(4)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "date w/ all options",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: date
                      order: 0
                      precision: 0
                      unsigned: true
                      notNull: true
                      default: 2000-12-01
                      primary: true
                      autoIncrement: true
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 date NOT NULL DEFAULT \"2000-12-01\" PRIMARY KEY\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "char",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: char
                      order: 20
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 char(20)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "varchar",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: varchar
                      order: 20
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 varchar(20)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "binary",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: binary
                      order: 20
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 binary(20)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "varbinary",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: varbinary
                      order: 20
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 varbinary(20)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "tinyblob",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: tinyblob
                      order: 0
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 tinyblob\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "tinytext",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: tinytext
                      order: 0
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 tinytext\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "blob",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: blob
                      order: 65535
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 blob(65535)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "text",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: text
                      order: 65535
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 text(65535)\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "mediumblob",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: mediumblob
                      order: 0
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 mediumblob\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "mediumtext",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: mediumtext
                      order: 0
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 mediumtext\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "longblob",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: longblob
                      order: 0
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 longblob\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "longtext",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: longtext
                      order: 0
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 longtext\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},

		{
			name: "text w/ all options",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: text
                      order: 65535
                      precision: 0
                      unsigned: true
                      notNull: true
                      default: "hoge"
                      primary: true
                      autoIncrement: true
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "CREATE TABLE IF NOT EXISTS table_a (\n    col_1 text(65535) NOT NULL PRIMARY KEY\n) DEFAULT CHARSET=utf8mb4",
			err: nil,
		},
	}

	for _, c := range cases {
		if err := viper.ReadConfig(bytes.NewBuffer(c.yaml)); err != nil {
			t.Errorf("case: %s is failed, err: %s\n", c.name, err)
		}

		err := cmd.LoadConfig()
		if !assert.Equal(t, c.err, err) {
			t.Errorf("case: %s is failed, err: %s\n", c.name, err)
		}
		config.Instance.CompleteWithDefault()

		for _, table := range config.Instance.Tables {
			client := database.MySQLClient{}
			sql := client.BuildCreateTableStmt(table)
			if !assert.Equal(t, c.sql, sql) {
				t.Errorf("case: %s is failed, expected: %+v, actual: %+v\n", c.name, c.sql, sql)
			}
		}

		// reset global variable
		config.Instance = nil
	}
}

func Test_BuildOrderDesc(t *testing.T) {
	viper.SetConfigType("yaml")

	cases := []struct {
		name   string
		yaml   []byte
		result string
		err    error
	}{
		{
			name: "tinyint default",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: tinyint
                      order:
                      precision:
                  charset: utf8mb4
                  record: 100000
            `),
			result: "(4)",
			err:    nil,
		},

		{
			name: "tinyint",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: tinyint
                      order: 2
                      precision:
                  charset: utf8mb4
                  record: 100000
            `),
			result: "(2)",
			err:    nil,
		},

		{
			name: "smallint default",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: smallint
                      order:
                      precision:
                  charset: utf8mb4
                  record: 100000
            `),
			result: "(6)",
			err:    nil,
		},

		{
			name: "smallint",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: smallint
                      order: 4
                      precision:
                  charset: utf8mb4
                  record: 100000
            `),
			result: "(4)",
			err:    nil,
		},

		{
			name: "mediumint default",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: mediumint
                      order:
                      precision:
                  charset: utf8mb4
                  record: 100000
            `),
			result: "(9)",
			err:    nil,
		},

		{
			name: "mediumint",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: mediumint
                      order: 6
                      precision:
                  charset: utf8mb4
                  record: 100000
            `),
			result: "(6)",
			err:    nil,
		},

		{
			name: "int default",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: int
                      order:
                      precision:
                  charset: utf8mb4
                  record: 100000
            `),
			result: "(11)",
			err:    nil,
		},

		{
			name: "int",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: int
                      order: 9
                      precision:
                  charset: utf8mb4
                  record: 100000
            `),
			result: "(9)",
			err:    nil,
		},

		{
			name: "bigint default",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: bigint
                      order:
                      precision:
                  charset: utf8mb4
                  record: 100000
            `),
			result: "(20)",
			err:    nil,
		},

		{
			name: "bigint",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: bigint
                      order: 11
                      precision:
                  charset: utf8mb4
                  record: 100000
            `),
			result: "(11)",
			err:    nil,
		},

		{
			name: "bit default",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: bit
                      order:
                      precision:
                  charset: utf8mb4
                  record: 100000
            `),
			result: "(1)",
			err:    nil,
		},

		{
			name: "bit",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: bit
                      order: 8
                      precision:
                  charset: utf8mb4
                  record: 100000
            `),
			result: "(8)",
			err:    nil,
		},

		{
			name: "other default",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: varchar
                      order:
                      precision:
                  charset: utf8mb4
                  record: 100000
            `),
			result: "(0)",
			err:    nil,
		},

		{
			name: "other",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: varchar
                      order: 8
                      precision:
                  charset: utf8mb4
                  record: 100000
            `),
			result: "(8)",
			err:    nil,
		},
	}

	for _, c := range cases {
		if err := viper.ReadConfig(bytes.NewBuffer(c.yaml)); err != nil {
			t.Errorf("case: %s is failed, err: %s\n", c.name, err)
		}

		err := cmd.LoadConfig()
		if !assert.Equal(t, c.err, err) {
			t.Errorf("case: %s is failed, err: %s\n", c.name, err)
		}
		config.Instance.CompleteWithDefault()

		for _, table := range config.Instance.Tables {
			client := database.MySQLClient{}
			for _, column := range table.Columns {
				result := client.BuildOrderDesc(column)
				if !assert.Equal(t, c.result, result) {
					t.Errorf("case: %s is failed, expected: %+v, actual: %+v\n", c.name, c.result, result)
				}
			}
		}

		// reset global variable
		config.Instance = nil
	}
}

func Test_BuildPrecisionDesc(t *testing.T) {
	viper.SetConfigType("yaml")

	cases := []struct {
		name   string
		yaml   []byte
		result string
		err    error
	}{
		{
			name: "decimal default",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: decimal
                      order:
                      precision:
                  charset: utf8mb4
                  record: 100000
            `),
			result: "(10, 0)",
			err:    nil,
		},

		{
			name: "decimal",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: decimal
                      order: 6
                      precision: 3
                  charset: utf8mb4
                  record: 100000
            `),
			result: "(6, 3)",
			err:    nil,
		},

		{
			name: "other default",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: float
                      order:
                      precision:
                  charset: utf8mb4
                  record: 100000
            `),
			result: "(0, 0)",
			err:    nil,
		},

		{
			name: "other",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: float
                      order: 6
                      precision: 3
                  charset: utf8mb4
                  record: 100000
            `),
			result: "(6, 3)",
			err:    nil,
		},
	}

	for _, c := range cases {
		if err := viper.ReadConfig(bytes.NewBuffer(c.yaml)); err != nil {
			t.Errorf("case: %s is failed, err: %s\n", c.name, err)
		}

		err := cmd.LoadConfig()
		if !assert.Equal(t, c.err, err) {
			t.Errorf("case: %s is failed, err: %s\n", c.name, err)
		}
		config.Instance.CompleteWithDefault()

		for _, table := range config.Instance.Tables {
			client := database.MySQLClient{}
			for _, column := range table.Columns {
				result := client.BuildPrecisionDesc(column)
				if !assert.Equal(t, c.result, result) {
					t.Errorf("case: %s is failed, expected: %+v, actual: %+v\n", c.name, c.result, result)
				}
			}
		}

		// reset global variable
		config.Instance = nil
	}
}

func Test_BuildDefaultDesc(t *testing.T) {
	viper.SetConfigType("yaml")

	cases := []struct {
		name   string
		yaml   []byte
		result string
		err    error
	}{
		{
			name: "string",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: varchar
                      order:
                      precision:
                      default: hoge
                  charset: utf8mb4
                  record: 100000
            `),
			result: " DEFAULT \"hoge\"",
			err:    nil,
		},

		{
			name: "other",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: decimal
                      order: 6
                      precision: 3
                      default: 123.456
                  charset: utf8mb4
                  record: 100000
            `),
			result: " DEFAULT(123.456)",
			err:    nil,
		},
	}

	for _, c := range cases {
		if err := viper.ReadConfig(bytes.NewBuffer(c.yaml)); err != nil {
			t.Errorf("case: %s is failed, err: %s\n", c.name, err)
		}

		err := cmd.LoadConfig()
		if !assert.Equal(t, c.err, err) {
			t.Errorf("case: %s is failed, err: %s\n", c.name, err)
		}
		config.Instance.CompleteWithDefault()

		for _, table := range config.Instance.Tables {
			client := database.MySQLClient{}
			for _, column := range table.Columns {
				result := client.BuildDefaultDesc(column)
				if !assert.Equal(t, c.result, result) {
					t.Errorf("case: %s is failed, expected: %+v, actual: %+v\n", c.name, c.result, result)
				}
			}
		}

		// reset global variable
		config.Instance = nil
	}
}

func Test_BuildDropTableStmt(t *testing.T) {
	viper.SetConfigType("yaml")

	cases := []struct {
		name string
		yaml []byte
		sql  string
		err  error
	}{
		{
			name: "boolean",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: boolean
                      order: 0
                      precision: 0
                  charset: utf8mb4
                  record: 100000
            `),
			sql: "DROP TABLE IF EXISTS table_a",
			err: nil,
		},
	}

	for _, c := range cases {
		if err := viper.ReadConfig(bytes.NewBuffer(c.yaml)); err != nil {
			t.Errorf("case: %s is failed, err: %s\n", c.name, err)
		}

		err := cmd.LoadConfig()
		if !assert.Equal(t, c.err, err) {
			t.Errorf("case: %s is failed, err: %s\n", c.name, err)
		}
		config.Instance.CompleteWithDefault()

		for _, table := range config.Instance.Tables {
			client := database.MySQLClient{}
			sql := client.BuildDropTableStmt(table)
			if !assert.Equal(t, c.sql, sql) {
				t.Errorf("case: %s is failed, expected: %+v, actual: %+v\n", c.name, c.sql, sql)
			}
		}

		// reset global variable
		config.Instance = nil
	}
}
