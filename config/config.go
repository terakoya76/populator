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
	Driver   Driver
	Database *Database
	Tables   []*Table
}

// CompleteWithDefault complete config value which is not required but configurable
func (c *config) CompleteWithDefault() {
	c.Driver.CompleteWithDefault()
	c.Database.CompleteWithDefault()
	for _, table := range c.Tables {
		table.CompleteWithDefault()
	}

}

// Validate validates config
func (c *config) Validate() error {
	if err := c.Driver.Validate(); err != nil {
		return err
	}

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

// Driver represents database driver
type Driver string

// CompleteWithDefault complete config value which is not required but configurable
func (d *Driver) CompleteWithDefault() {
}

// Validate validates database driver config
func (d Driver) Validate() error {
	// TODO: adopt more
	switch d {
	case "mysql":
		return nil
	default:
		return errors.New("database driver is invalid or non-supported")
	}
}

// Database represents information for connecting DB
type Database struct {
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
