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
	"sync"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	"github.com/terakoya76/populator/config"
)

// DBConnector is adapter interface for connecting DB
type DBConnector interface {
	Connect(cfg *config.DatabaseConfig) (*sqlx.DB, error)
	connectInfo(cfg *config.DatabaseConfig) string
}

var instance *sqlx.DB

// onceMySQL is used for Mutex Lock when initialize instanceMySQL
var onceDB sync.Once

// DB provides instance of DB client
func DB() *sqlx.DB {
	onceDB.Do(func() {
		initialize()
	})
	return instance
}

func initialize() {
	cfg := viper.Sub("database")
	dbCfg := config.NewDatabaseConfig(cfg)
	fmt.Printf("Database config file: %+v\n", dbCfg)

	var c DBConnector
	if dbCfg.Driver == "mysql" {
		c = NewMySQLConnector()
	}

	var err error
	instance, err = c.Connect(dbCfg)
	if err != nil {
		fmt.Println("Failed to connect database: ", err)
	}
}
