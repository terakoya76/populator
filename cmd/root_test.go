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
package cmd_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/terakoya76/populator/cmd"
	"github.com/terakoya76/populator/config"
)

var (
	nilDatabase *config.Database
	nilTables   []*config.Table
	nilColumns  []*config.Column
	nilIndexes  []*config.Index
	nilValues   []interface{}
)

//nolint:funlen
func Test_LoadConfig_Database(t *testing.T) {
	viper.SetConfigType("yaml")

	cases := []struct {
		name   string
		yaml   []byte
		config interface{}
		err    error
	}{
		{
			name: "normal case",
			yaml: []byte(`
                database:
                  driver: mysql
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
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: &config.Database{
				Driver:   "mysql",
				Host:     "127.0.0.1",
				Port:     3306,
				User:     "root",
				Password: "root",
				Name:     "testdb",
			},
			err: nil,
		},

		{
			name: "missing a whole database part in yaml",
			yaml: []byte(`
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: float
                      order: 5
                      precision: 2
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: nilDatabase,
			err:    errors.New("database connection information is required"),
		},

		{
			name: "missing a driver of database part in yaml",
			yaml: []byte(`
                driver: mysql
                database:
                  host: "127.0.0.1"
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
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: &config.Database{
				Driver:   "",
				Host:     "127.0.0.1",
				Port:     3306,
				User:     "root",
				Password: "root",
				Name:     "testdb",
			},
			err: errors.New("database driver is invalid or non-supported"),
		},

		{
			name: "an unsupported driver of database part in yaml",
			yaml: []byte(`
                database:
                  driver: hoge
                  host: "127.0.0.1"
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
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: &config.Database{
				Driver:   "hoge",
				Host:     "127.0.0.1",
				Port:     3306,
				User:     "root",
				Password: "root",
				Name:     "testdb",
			},
			err: errors.New("database driver is invalid or non-supported"),
		},

		{
			name: "missing a host of database part in yaml",
			yaml: []byte(`
                database:
                  driver: mysql
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
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: &config.Database{
				Driver:   "mysql",
				Host:     "",
				Port:     3306,
				User:     "root",
				Password: "root",
				Name:     "testdb",
			},
			err: nil,
		},

		{
			name: "missing a port of database part in yaml",
			yaml: []byte(`
                database:
                  driver: mysql
                  host: "127.0.0.1"
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
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: &config.Database{
				Driver:   "mysql",
				Host:     "127.0.0.1",
				Port:     0,
				User:     "root",
				Password: "root",
				Name:     "testdb",
			},
			err: nil,
		},

		{
			name: "missing a user of database part in yaml",
			yaml: []byte(`
                database:
                  driver: mysql
                  host: 127.0.0.1
                  port: 3306
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: float
                      order: 5
                      precision: 2
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: &config.Database{
				Driver:   "mysql",
				Host:     "127.0.0.1",
				Port:     3306,
				User:     "",
				Password: "root",
				Name:     "testdb",
			},
			err: errors.New("database user is required"),
		},

		{
			name: "an empty password of database part in yaml is allowed",
			yaml: []byte(`
                database:
                  driver: mysql
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: float
                      order: 5
                      precision: 2
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: &config.Database{
				Driver:   "mysql",
				Host:     "127.0.0.1",
				Port:     3306,
				User:     "root",
				Password: "",
				Name:     "testdb",
			},
			err: nil,
		},

		{
			name: "missing a name of database part in yaml",
			yaml: []byte(`
                database:
                  driver: mysql
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: float
                      order: 5
                      precision: 2
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: &config.Database{
				Driver:   "mysql",
				Host:     "127.0.0.1",
				Port:     3306,
				User:     "root",
				Password: "root",
				Name:     "",
			},
			err: errors.New("database name is required"),
		},
	}

	//nolint:dupl
	for _, c := range cases {
		if err := viper.ReadConfig(bytes.NewBuffer(c.yaml)); err != nil {
			t.Errorf("case: %s is failed, err: %s\n", c.name, err)
		}

		err := cmd.LoadConfig()
		if !assert.Equal(t, c.err, err) {
			t.Errorf("case: %s is failed, expected: %s, actual: %s\n", c.name, c.err, err)
		}

		cfg := config.Instance
		if !assert.Equal(t, c.config, cfg.Database) {
			t.Errorf("case: %s is failed, expected: %+v, actual: %+v\n", c.name, c.config, cfg.Database)
		}

		// reset global variable
		config.Instance = nil
	}
}

//nolint:funlen
func Test_LoadConfig_Tables(t *testing.T) {
	viper.SetConfigType("yaml")

	cases := []struct {
		name   string
		yaml   []byte
		config interface{}
		err    error
	}{
		{
			name: "normal case",
			yaml: []byte(`
                database:
                  driver: mysql
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
                      values:
                        - 123.45
                        - 678.90
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
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
								678.90,
							},
						},
					},
					Indexes: []*config.Index{
						{
							Name:    "index_1_on_table_a",
							Primary: false,
							Uniq:    false,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: nil,
		},

		{
			name: "missing a whole tables part in yaml",
			yaml: []byte(`
                database:
                  driver: mysql
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
            `),
			config: nilTables,
			err:    errors.New("tables definition is required"),
		},

		{
			name: "missing a name of table part in yaml",
			yaml: []byte(`
                database:
                  driver: mysql
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - columns:
                    - name: col_1
                      type: float
                      order: 5
                      precision: 2
                      unsigned: true
                      notNull: true
                      default: 123.45
                      primary: true
                      autoIncrement: true
                      values:
                        - 123.45
                        - 678.90
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
					Name: "",
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
								678.90,
							},
						},
					},
					Indexes: []*config.Index{
						{
							Name:    "index_1_on_table_a",
							Primary: false,
							Uniq:    false,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: errors.New("table name is required"),
		},

		{
			name: "missing a columns of table part in yaml",
			yaml: []byte(`
                database:
                  driver: mysql
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: "table_a"
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
					Name:    "table_a",
					Columns: nilColumns,
					Indexes: []*config.Index{
						{
							Name:    "index_1_on_table_a",
							Primary: false,
							Uniq:    false,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: nil,
		},

		{
			name: "missing a indexes of table part in yaml",
			yaml: []byte(`
                database:
                  driver: mysql
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
                      values:
                        - 123.45
                        - 678.90
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
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
								678.90,
							},
						},
					},
					Indexes: nilIndexes,
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: nil,
		},

		{
			name: "missing a charset of table part in yaml",
			yaml: []byte(`
                database:
                  driver: mysql
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
                      values:
                        - 123.45
                        - 678.90
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  record: 100000
            `),
			config: []*config.Table{
				{
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
								678.90,
							},
						},
					},
					Indexes: []*config.Index{
						{
							Name:    "index_1_on_table_a",
							Primary: false,
							Uniq:    false,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "",
					Record:  100000,
				},
			},
			err: nil,
		},

		{
			name: "missing a record of table part in yaml",
			yaml: []byte(`
                database:
                  driver: mysql
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
                      values:
                        - 123.45
                        - 678.90
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
            `),
			config: []*config.Table{
				{
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
								678.90,
							},
						},
					},
					Indexes: []*config.Index{
						{
							Name:    "index_1_on_table_a",
							Primary: false,
							Uniq:    false,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "utf8mb4",
					Record:  0,
				},
			},
			err: nil,
		},
	}

	//nolint:dupl
	for _, c := range cases {
		if err := viper.ReadConfig(bytes.NewBuffer(c.yaml)); err != nil {
			t.Errorf("case: %s is failed, err: %s\n", c.name, err)
		}

		err := cmd.LoadConfig()
		if !assert.Equal(t, c.err, err) {
			t.Errorf("case: %s is failed, expected: %s, actual: %s\n", c.name, c.err, err)
		}

		cfg := config.Instance
		if !assert.Equal(t, c.config, cfg.Tables) {
			t.Errorf("case: %s is failed, expected: %+v, actual: %+v\n", c.name, c.config, cfg.Tables)
		}

		// reset global variable
		config.Instance = nil
	}
}

//nolint:funlen
func Test_LoadConfig_Columns(t *testing.T) {
	viper.SetConfigType("yaml")

	cases := []struct {
		name   string
		yaml   []byte
		config interface{}
		err    error
	}{
		{
			name: "missing a name of column in Columns part",
			yaml: []byte(`
                database:
                  driver: mysql
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - type: float
                      order: 5
                      precision: 2
                      unsigned: true
                      notNull: true
                      default: 123.45
                      primary: true
                      autoIncrement: true
                      values:
                        - 123.45
                        - 678.90
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
					Name: "table_a",
					Columns: []*config.Column{
						{
							Name:          "",
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
								678.90,
							},
						},
					},
					Indexes: []*config.Index{
						{
							Name:    "index_1_on_table_a",
							Primary: false,
							Uniq:    false,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: nil,
		},

		{
			name: "missing a type of column in columns part",
			yaml: []byte(`
                database:
                  driver: mysql
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      order: 5
                      precision: 2
                      unsigned: true
                      notNull: true
                      default: 123.45
                      primary: true
                      autoIncrement: true
                      values:
                        - 123.45
                        - 678.90
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
					Name: "table_a",
					Columns: []*config.Column{
						{
							Name:          "col_1",
							Type:          "",
							Order:         5,
							Precision:     2,
							Unsigned:      true,
							NotNull:       true,
							Default:       123.45,
							Primary:       true,
							AutoIncrement: true,
							Values: []interface{}{
								123.45,
								678.90,
							},
						},
					},
					Indexes: []*config.Index{
						{
							Name:    "index_1_on_table_a",
							Primary: false,
							Uniq:    false,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: nil,
		},

		{
			name: "missing a order of column in columns part",
			yaml: []byte(`
                database:
                  driver: mysql
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
                      precision: 2
                      unsigned: true
                      notNull: true
                      default: 123.45
                      primary: true
                      autoIncrement: true
                      values:
                        - 123.45
                        - 678.90
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
					Name: "table_a",
					Columns: []*config.Column{
						{
							Name:          "col_1",
							Type:          "float",
							Order:         0,
							Precision:     2,
							Unsigned:      true,
							NotNull:       true,
							Default:       123.45,
							Primary:       true,
							AutoIncrement: true,
							Values: []interface{}{
								123.45,
								678.90,
							},
						},
					},
					Indexes: []*config.Index{
						{
							Name:    "index_1_on_table_a",
							Primary: false,
							Uniq:    false,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: nil,
		},

		{
			name: "missing a precision of column in columns part",
			yaml: []byte(`
                database:
                  driver: mysql
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
                      unsigned: true
                      notNull: true
                      default: 123.45
                      primary: true
                      autoIncrement: true
                      values:
                        - 123.45
                        - 678.90
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
					Name: "table_a",
					Columns: []*config.Column{
						{
							Name:          "col_1",
							Type:          "float",
							Order:         5,
							Precision:     0,
							Unsigned:      true,
							NotNull:       true,
							Default:       123.45,
							Primary:       true,
							AutoIncrement: true,
							Values: []interface{}{
								123.45,
								678.90,
							},
						},
					},
					Indexes: []*config.Index{
						{
							Name:    "index_1_on_table_a",
							Primary: false,
							Uniq:    false,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: nil,
		},

		{
			name: "missing a unsigned of column in columns part",
			yaml: []byte(`
                database:
                  driver: mysql
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
                      notNull: true
                      default: 123.45
                      primary: true
                      autoIncrement: true
                      values:
                        - 123.45
                        - 678.90
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
					Name: "table_a",
					Columns: []*config.Column{
						{
							Name:          "col_1",
							Type:          "float",
							Order:         5,
							Precision:     2,
							Unsigned:      false,
							NotNull:       true,
							Default:       123.45,
							Primary:       true,
							AutoIncrement: true,
							Values: []interface{}{
								123.45,
								678.90,
							},
						},
					},
					Indexes: []*config.Index{
						{
							Name:    "index_1_on_table_a",
							Primary: false,
							Uniq:    false,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: nil,
		},

		{
			name: "missing a notNull of column in columns part",
			yaml: []byte(`
                database:
                  driver: mysql
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
                      default: 123.45
                      primary: true
                      autoIncrement: true
                      values:
                        - 123.45
                        - 678.90
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
					Name: "table_a",
					Columns: []*config.Column{
						{
							Name:          "col_1",
							Type:          "float",
							Order:         5,
							Precision:     2,
							Unsigned:      true,
							NotNull:       false,
							Default:       123.45,
							Primary:       true,
							AutoIncrement: true,
							Values: []interface{}{
								123.45,
								678.90,
							},
						},
					},
					Indexes: []*config.Index{
						{
							Name:    "index_1_on_table_a",
							Primary: false,
							Uniq:    false,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: nil,
		},

		{
			name: "missing a default of column in columns part",
			yaml: []byte(`
                database:
                  driver: mysql
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
                      primary: true
                      autoIncrement: true
                      values:
                        - 123.45
                        - 678.90
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
					Name: "table_a",
					Columns: []*config.Column{
						{
							Name:          "col_1",
							Type:          "float",
							Order:         5,
							Precision:     2,
							Unsigned:      true,
							NotNull:       true,
							Default:       nil,
							Primary:       true,
							AutoIncrement: true,
							Values: []interface{}{
								123.45,
								678.90,
							},
						},
					},
					Indexes: []*config.Index{
						{
							Name:    "index_1_on_table_a",
							Primary: false,
							Uniq:    false,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: nil,
		},

		{
			name: "missing a primary of column in columns part",
			yaml: []byte(`
                database:
                  driver: mysql
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
                      autoIncrement: true
                      values:
                        - 123.45
                        - 678.90
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
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
							Primary:       false,
							AutoIncrement: true,
							Values: []interface{}{
								123.45,
								678.90,
							},
						},
					},
					Indexes: []*config.Index{
						{
							Name:    "index_1_on_table_a",
							Primary: false,
							Uniq:    false,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: nil,
		},

		{
			name: "missing a autoIncrement of column in columns part",
			yaml: []byte(`
                database:
                  driver: mysql
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
                      values:
                        - 123.45
                        - 678.90
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
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
							AutoIncrement: false,
							Values: []interface{}{
								123.45,
								678.90,
							},
						},
					},
					Indexes: []*config.Index{
						{
							Name:    "index_1_on_table_a",
							Primary: false,
							Uniq:    false,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: nil,
		},

		{
			name: "missing a values of column in columns part",
			yaml: []byte(`
                database:
                  driver: mysql
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
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
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
							Values:        nilValues,
						},
					},
					Indexes: []*config.Index{
						{
							Name:    "index_1_on_table_a",
							Primary: false,
							Uniq:    false,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: nil,
		},
	}

	//nolint:dupl
	for _, c := range cases {
		if err := viper.ReadConfig(bytes.NewBuffer(c.yaml)); err != nil {
			t.Errorf("case: %s is failed, err: %s\n", c.name, err)
		}

		err := cmd.LoadConfig()
		if !assert.Equal(t, c.err, err) {
			t.Errorf("case: %s is failed, expected: %s, actual: %s\n", c.name, c.err, err)
		}

		cfg := config.Instance
		if !assert.Equal(t, c.config, cfg.Tables) {
			t.Errorf("case: %s is failed, expected: %+v, actual: %+v\n", c.name, c.config, cfg.Tables)
		}

		// reset global variable
		config.Instance = nil
	}
}

//nolint:funlen
func Test_LoadConfig_Indexes(t *testing.T) {
	viper.SetConfigType("yaml")

	cases := []struct {
		name   string
		yaml   []byte
		config interface{}
		err    error
	}{
		{
			name: "missing a name of index in indexes part",
			yaml: []byte(`
                database:
                  driver: mysql
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
                  indexes:
                    - primary: false
                      uniq: false
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
					Name: "table_a",
					Columns: []*config.Column{
						{
							Name:          "col_1",
							Type:          "float",
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
					Indexes: []*config.Index{
						{
							Name:    "",
							Primary: false,
							Uniq:    false,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: nil,
		},

		{
			name: "missing a uniq of index in indexes part",
			yaml: []byte(`
                database:
                  driver: mysql
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
                  indexes:
                    - primary: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
					Name: "table_a",
					Columns: []*config.Column{
						{
							Name:          "col_1",
							Type:          "float",
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
					Indexes: []*config.Index{
						{
							Name:    "",
							Primary: true,
							Uniq:    false,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: nil,
		},

		{
			name: "missing a primary of index in indexes part",
			yaml: []byte(`
                database:
                  driver: mysql
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
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
					Name: "table_a",
					Columns: []*config.Column{
						{
							Name:          "col_1",
							Type:          "float",
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
					Indexes: []*config.Index{
						{
							Name:    "index_1_on_table_a",
							Primary: false,
							Uniq:    true,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: nil,
		},

		{
			name: "missing a columns of index in indexes part",
			yaml: []byte(`
                database:
                  driver: mysql
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
                  indexes:
                    - name: index_1_on_table_a
                      primary: false
                      uniq: false
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
					Name: "table_a",
					Columns: []*config.Column{
						{
							Name:          "col_1",
							Type:          "float",
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
					Indexes: []*config.Index{
						{
							Name:    "index_1_on_table_a",
							Primary: false,
							Uniq:    false,
							Columns: nil,
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: nil,
		},

		{
			name: "primary key w/ name",
			yaml: []byte(`
                database:
                  driver: mysql
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
                      values:
                  indexes:
                    - name: index_1_on_table_a
                      primary: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
					Name: "table_a",
					Columns: []*config.Column{
						{
							Name:          "col_1",
							Type:          "float",
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
					Indexes: []*config.Index{
						{
							Name:    "index_1_on_table_a",
							Primary: true,
							Uniq:    false,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: errors.New("primary key index cannot be named"),
		},

		{
			name: "primary key w/ unique key",
			yaml: []byte(`
                database:
                  driver: mysql
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
                      values:
                  indexes:
                    - primary: true
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				{
					Name: "table_a",
					Columns: []*config.Column{
						{
							Name:          "col_1",
							Type:          "float",
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
					Indexes: []*config.Index{
						{
							Name:    "",
							Primary: true,
							Uniq:    true,
							Columns: []string{
								"col_1",
							},
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: errors.New("both of primary key and unique key cannot be enabled"),
		},
	}

	//nolint:dupl
	for _, c := range cases {
		if err := viper.ReadConfig(bytes.NewBuffer(c.yaml)); err != nil {
			t.Errorf("case: %s is failed, err: %s\n", c.name, err)
		}

		err := cmd.LoadConfig()
		if !assert.Equal(t, c.err, err) {
			t.Errorf("case: %s is failed, expected: %s, actual: %s\n", c.name, c.err, err)
		}

		cfg := config.Instance
		if !assert.Equal(t, c.config, cfg.Tables) {
			t.Errorf("case: %s is failed, expected: %+v, actual: %+v\n", c.name, c.config, cfg.Tables)
		}

		// reset global variable
		config.Instance = nil
	}
}
