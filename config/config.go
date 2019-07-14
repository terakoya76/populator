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
	"fmt"
)

// Instance represents the both information of the connecting database and the tables schema to be populated w/ seed data
var Instance *config

type config struct {
	Driver   *Driver
	Database *Database
	Tables   []*Table
}

// Driver represents database driver
type Driver string

// IsValid validates given database driver is acceptable or not
func (d *Driver) IsValid() bool {
	// TODO: adopt more
	if *d != "mysql" {
		fmt.Println("supported database driver is only mysql")
		return false
	}
	return true
}

// Database represents information for connecting DB
type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

// Table represents a single table schema
type Table struct {
	Name    string
	Columns []*Column
	Indexes []*Index
	Record  int
}

// Column represents a single column schema
type Column struct {
	Name    string
	Type    string
	Primary bool
	Null    bool
	Charset string
}

// Index represents a single index schema
type Index struct {
	Name    string
	Uniq    bool
	Columns []string
}
