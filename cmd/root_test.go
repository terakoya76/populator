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

func Test_LoadConfig_Driver(t *testing.T) {
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
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: config.Driver("mysql"),
			err:    nil,
		},

		{
			name: "non-supported driver",
			yaml: []byte(`
                driver: postgres
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
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: config.Driver("postgres"),
			err:    errors.New("database driver is invalid or non-supported"),
		},

		{
			name: "missing driver in yaml",
			yaml: []byte(`
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
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: config.Driver(""),
			err:    errors.New("database driver is invalid or non-supported"),
		},

		{
			name: "yaml syntax error",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: int
                      order: 11
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: config.Driver(""),
			err:    errors.New("yaml syntax error"),
		},
	}

	for _, c := range cases {
		viper.ReadConfig(bytes.NewBuffer(c.yaml))
		err := cmd.LoadConfig()
		if !assert.Equal(t, c.err, err) {
			t.Errorf("case: %s is failed, expected: %s, actual: %s\n", c.name, c.err, err)
		}

		cfg := config.Instance
		if !assert.Equal(t, c.config, cfg.Driver) {
			t.Errorf("case: %s is failed, expected: %+v, actual: %+v\n", c.name, c.config, cfg.Driver)
		}

		// reset global variable
		config.Instance = nil
	}
}

func Test_LoadConfig_Database(t *testing.T) {
	viper.SetConfigType("yaml")

	var nilDatabase *config.Database
	cases := []struct {
		name   string
		yaml   []byte
		config interface{}
		err    error
	}{
		{
			name: "normal case",
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
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: &config.Database{
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
                driver: mysql
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: int
                      order: 11
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: nilDatabase,
			err:    errors.New("database connection information is required"),
		},

		{
			name: "missing a host and port of database part in yaml",
			yaml: []byte(`
                driver: mysql
                database:
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: int
                      order: 11
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: &config.Database{
				Host:     "127.0.0.1",
				Port:     3306,
				User:     "root",
				Password: "root",
				Name:     "testdb",
			},
			err: nil,
		},

		{
			name: "missing a user of database part in yaml",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: int
                      order: 11
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: &config.Database{
				Host:     "127.0.0.1",
				Port:     3306,
				User:     "",
				Password: "root",
				Name:     "testdb",
			},
			err: errors.New("database user is required"),
		},

		{
			name: "missing a password of database part in yaml",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: int
                      order: 11
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: &config.Database{
				Host:     "127.0.0.1",
				Port:     3306,
				User:     "root",
				Password: "",
				Name:     "testdb",
			},
			err: errors.New("database password is required"),
		},

		{
			name: "missing a name of database part in yaml",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: int
                      order: 11
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: &config.Database{
				Host:     "127.0.0.1",
				Port:     3306,
				User:     "root",
				Password: "root",
				Name:     "",
			},
			err: errors.New("database name is required"),
		},

		{
			name: "yaml syntax error",
			yaml: []byte(`
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user root
                  password: root
                  name: testdb
                tables:
                - name: table_a
                  columns:
                    - name: col_1
                      type: int
                      order: 11
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: nilDatabase,
			err:    errors.New("yaml syntax error"),
		},
	}

	for _, c := range cases {
		viper.ReadConfig(bytes.NewBuffer(c.yaml))
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

func Test_LoadConfig_Tables(t *testing.T) {
	viper.SetConfigType("yaml")

	var nilTables []*config.Table
	var nilColumns []*config.Column
	var nilIndexes []*config.Index
	cases := []struct {
		name   string
		yaml   []byte
		config interface{}
		err    error
	}{
		{
			name: "normal case",
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
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				&config.Table{
					Name: "table_a",
					Columns: []*config.Column{
						&config.Column{
							Name:    "col_1",
							Type:    "int",
							Order:   11,
							Null:    false,
							Primary: true,
						},
					},
					Indexes: []*config.Index{
						&config.Index{
							Name: "index_1_on_table_a",
							Uniq: true,
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
                driver: mysql
                database:
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
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - columns:
                    - name: col_1
                      type: int
                      order: 11
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				&config.Table{
					Name: "",
					Columns: []*config.Column{
						&config.Column{
							Name:    "col_1",
							Type:    "int",
							Order:   11,
							Null:    false,
							Primary: true,
						},
					},
					Indexes: []*config.Index{
						&config.Index{
							Name: "index_1_on_table_a",
							Uniq: true,
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
                driver: mysql
                database:
                  host: 127.0.0.1
                  port: 3306
                  user: root
                  password: root
                  name: testdb
                tables:
                - name: "table_a"
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				&config.Table{
					Name:    "table_a",
					Columns: nilColumns,
					Indexes: []*config.Index{
						&config.Index{
							Name: "index_1_on_table_a",
							Uniq: true,
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
                      null: false
                      primary: true
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				&config.Table{
					Name: "table_a",
					Columns: []*config.Column{
						&config.Column{
							Name:    "col_1",
							Type:    "int",
							Order:   11,
							Null:    false,
							Primary: true,
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
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  record: 100000
            `),
			config: []*config.Table{
				&config.Table{
					Name: "table_a",
					Columns: []*config.Column{
						&config.Column{
							Name:    "col_1",
							Type:    "int",
							Order:   11,
							Null:    false,
							Primary: true,
						},
					},
					Indexes: []*config.Index{
						&config.Index{
							Name: "index_1_on_table_a",
							Uniq: true,
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
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
            `),
			config: []*config.Table{
				&config.Table{
					Name: "table_a",
					Columns: []*config.Column{
						&config.Column{
							Name:    "col_1",
							Type:    "int",
							Order:   11,
							Null:    false,
							Primary: true,
						},
					},
					Indexes: []*config.Index{
						&config.Index{
							Name: "index_1_on_table_a",
							Uniq: true,
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

		{
			name: "yaml syntax error",
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
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset utf8mb4
                  record: 100000
            `),
			config: nilTables,
			err:    errors.New("yaml syntax error"),
		},
	}

	for _, c := range cases {
		viper.ReadConfig(bytes.NewBuffer(c.yaml))
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

func Test_LoadConfig_Columns(t *testing.T) {
	viper.SetConfigType("yaml")

	var nilTables []*config.Table
	cases := []struct {
		name   string
		yaml   []byte
		config interface{}
		err    error
	}{
		{
			name: "missing a name of column in Columns part",
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
                    - type: int
                      order: 11
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				&config.Table{
					Name: "table_a",
					Columns: []*config.Column{
						&config.Column{
							Name:    "",
							Type:    "int",
							Order:   11,
							Null:    false,
							Primary: true,
						},
					},
					Indexes: []*config.Index{
						&config.Index{
							Name: "index_1_on_table_a",
							Uniq: true,
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
                      order: 11
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				&config.Table{
					Name: "table_a",
					Columns: []*config.Column{
						&config.Column{
							Name:    "col_1",
							Type:    "",
							Order:   11,
							Null:    false,
							Primary: true,
						},
					},
					Indexes: []*config.Index{
						&config.Index{
							Name: "index_1_on_table_a",
							Uniq: true,
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
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				&config.Table{
					Name: "table_a",
					Columns: []*config.Column{
						&config.Column{
							Name:    "col_1",
							Type:    "int",
							Order:   0,
							Null:    false,
							Primary: true,
						},
					},
					Indexes: []*config.Index{
						&config.Index{
							Name: "index_1_on_table_a",
							Uniq: true,
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
			name: "missing a null of column in columns part",
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
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				&config.Table{
					Name: "table_a",
					Columns: []*config.Column{
						&config.Column{
							Name:    "col_1",
							Type:    "int",
							Order:   11,
							Null:    false,
							Primary: true,
						},
					},
					Indexes: []*config.Index{
						&config.Index{
							Name: "index_1_on_table_a",
							Uniq: true,
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
                      null: false
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				&config.Table{
					Name: "table_a",
					Columns: []*config.Column{
						&config.Column{
							Name:    "col_1",
							Type:    "int",
							Order:   11,
							Null:    false,
							Primary: false,
						},
					},
					Indexes: []*config.Index{
						&config.Index{
							Name: "index_1_on_table_a",
							Uniq: true,
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
			name: "yaml syntax error",
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
                      type int
                      order: 11
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: nilTables,
			err:    errors.New("yaml syntax error"),
		},
	}

	for _, c := range cases {
		viper.ReadConfig(bytes.NewBuffer(c.yaml))
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

func Test_LoadConfig_Indexes(t *testing.T) {
	viper.SetConfigType("yaml")

	var nilTables []*config.Table
	var nilColumns []string
	cases := []struct {
		name   string
		yaml   []byte
		config interface{}
		err    error
	}{
		{
			name: "missing a name of index in indexes part",
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
                      null: false
                      primary: true
                  indexes:
                    - uniq: true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				&config.Table{
					Name: "table_a",
					Columns: []*config.Column{
						&config.Column{
							Name:    "col_1",
							Type:    "int",
							Order:   11,
							Null:    false,
							Primary: true,
						},
					},
					Indexes: []*config.Index{
						&config.Index{
							Name: "",
							Uniq: true,
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
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				&config.Table{
					Name: "table_a",
					Columns: []*config.Column{
						&config.Column{
							Name:    "col_1",
							Type:    "int",
							Order:   11,
							Null:    false,
							Primary: true,
						},
					},
					Indexes: []*config.Index{
						&config.Index{
							Name: "index_1_on_table_a",
							Uniq: false,
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
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq: true
                  charset: utf8mb4
                  record: 100000
            `),
			config: []*config.Table{
				&config.Table{
					Name: "table_a",
					Columns: []*config.Column{
						&config.Column{
							Name:    "col_1",
							Type:    "int",
							Order:   11,
							Null:    false,
							Primary: true,
						},
					},
					Indexes: []*config.Index{
						&config.Index{
							Name:    "index_1_on_table_a",
							Uniq:    true,
							Columns: nilColumns,
						},
					},
					Charset: "utf8mb4",
					Record:  100000,
				},
			},
			err: nil,
		},

		{
			name: "yaml syntax error",
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
                      null: false
                      primary: true
                  indexes:
                    - name: index_1_on_table_a
                      uniq true
                      columns:
                        - col_1
                  charset: utf8mb4
                  record: 100000
            `),
			config: nilTables,
			err:    errors.New("yaml syntax error"),
		},
	}

	for _, c := range cases {
		viper.ReadConfig(bytes.NewBuffer(c.yaml))
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
