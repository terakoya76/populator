/*
Package config ...

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
package config

import (
	"errors"
)

// Instance represents the both information of the connecting database and the tables schema to be populated w/ seed data
var Instance *config

type config struct {
	Database *Database
	Tables   []*Table
}

// CompleteWithDefault complete config value which is not required but configurable
func (c *config) CompleteWithDefault() {
	c.Database.CompleteWithDefault()
	for _, table := range c.Tables {
		table.CompleteWithDefault()
	}

}

// Validate validates config
func (c *config) Validate() error {
	if err := c.Database.Validate(); err != nil {
		return err
	}

	if c.Tables == nil {
		return errors.New("tables definition is required")
	}
	for _, table := range c.Tables {
		if err := table.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Database represents information for connecting DB
type Database struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

// Validate validates database config
func (db *Database) Validate() error {
	if db == nil {
		return errors.New("database connection information is required")
	}

	// TODO: adopt more
	switch db.Driver {
	case "mysql":
	default:
		return errors.New("database driver is invalid or non-supported")
	}

	if db.User == "" {
		return errors.New("database user is required")
	}
	if db.Name == "" {
		return errors.New("database name is required")
	}
	return nil
}

// CompleteWithDefault complete config value which is not required but configurable
func (db *Database) CompleteWithDefault() {
	if db.Host == "" {
		db.Host = "127.0.0.1"
	}
	if db.Port == 0 {
		db.Port = 3306
	}
}

// Table represents a single table schema
type Table struct {
	Name    string
	Columns []*Column
	Indexes []*Index
	Charset string
	Record  int
}

// CompleteWithDefault complete config value which is not required but configurable
func (t *Table) CompleteWithDefault() {
	for _, column := range t.Columns {
		column.CompleteWithDefault()
	}
	for _, index := range t.Indexes {
		index.CompleteWithDefault()
	}
}

// Validate validates table config
func (t *Table) Validate() error {
	if t == nil {
		return errors.New("table definition is invalid")
	}
	if t.Name == "" {
		return errors.New("table name is required")
	}
	for _, column := range t.Columns {
		if err := column.Validate(); err != nil {
			return err
		}
	}
	for _, index := range t.Indexes {
		if err := index.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Column represents a single column schema
type Column struct {
	Name          string
	Type          string
	Order         int
	Precision     int
	Unsigned      bool
	NotNull       bool
	Default       interface{}
	Primary       bool
	AutoIncrement bool
	Values        []interface{}
}

// CompleteWithDefault complete config value which is not required but configurable
func (c *Column) CompleteWithDefault() {
	if c.Order == 0 {
		switch c.Type {
		case "tinyint":
			c.Order = 4
		case "smallint":
			c.Order = 6
		case "mediumint":
			c.Order = 9
		case "int":
			c.Order = 11
		case "bigint":
			c.Order = 20
		case "bit":
			c.Order = 1
		case "blob":
			// Unofficial Value
			c.Order = 65535
		case "text":
			// Unofficial Value
			c.Order = 65535
		case "year":
			// Unofficial Value
			c.Order = 4
		case "char":
			// Unofficial Value
			c.Order = 255
		case "varchar":
			// Unofficial Value
			c.Order = 255
		case "binary":
			// Unofficial Value
			c.Order = 255
		case "varbinary":
			// Unofficial Value
			c.Order = 255
		default:
		}
	}

	if c.Order == 0 && c.Precision == 0 {
		switch c.Type {
		case "decimal":
			c.Order = 10
			c.Precision = 0
		case "float":
			// Unofficial Value
			c.Order = 5
			c.Precision = 2
		case "real":
			// Unofficial Value
			c.Order = 5
			c.Precision = 10
		case "double":
			// Unofficial Value
			c.Order = 5
			c.Precision = 10
		default:
		}
	}
}

// Validate validates column config
func (c *Column) Validate() error {
	return nil
}

// Index represents a single index schema
type Index struct {
	Name    string
	Primary bool
	Uniq    bool
	Columns []string
}

// CompleteWithDefault complete config value which is not required but configurable
func (i *Index) CompleteWithDefault() {
}

// Validate validates index config
func (i *Index) Validate() error {
	if i.Primary {
		if i.Name != "" {
			return errors.New("primary key index cannot be named")
		}
		if i.Uniq {
			return errors.New("both of primary key and unique key cannot be enabled")

		}
	}
	return nil
}
