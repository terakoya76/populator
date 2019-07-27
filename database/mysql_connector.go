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

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/terakoya76/populator/config"
)

// MaxConnections holds max_connections var for memory use control
var MaxConnections int

// MySQLConnector is an implementation of DBConnector for MySQL
type MySQLConnector struct{}

// NewMySQLConnector constructs MySQLConnector instance
func NewMySQLConnector() *MySQLConnector {
	return &MySQLConnector{}
}

// Connect find_or_create database w/ given database name, then connect it
func (c *MySQLConnector) Connect(cfg *config.Database) (DBClient, error) {
	ci := c.connectInfo(cfg)
	db, err := sqlx.Open("mysql", ci)
	if err != nil {
		return nil, fmt.Errorf("failed to setup database %s on mysql: %+v", cfg.Name, err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + cfg.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to setup database %s on mysql: %+v", cfg.Name, err)
	}
	db.Close()

	db, err = sqlx.Connect("mysql", ci+cfg.Name)
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

func (c *MySQLConnector) connectInfo(cfg *config.Database) string {
	return cfg.User + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + fmt.Sprint(cfg.Port) + ")/"
}
