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

	"github.com/spf13/viper"

	"github.com/terakoya76/populator/config"
)

// DBClient is an interface for DB Querying
type DBClient interface {
	CreateTable(cfg []*config.Table)
	CreateIndex(cfg []*config.Table)
	Insert(cfg []*config.Table)
}

// DBConnector is an interface for DB adapter
type DBConnector interface {
	Connect(cfg *config.Database) (DBClient, error)
	connectInfo(cfg *config.Database) string
}

var client DBClient

// onceDB is used for Mutex Lock when initializing an instance
var onceDB sync.Once

// DB provides instance of DB client
func DB() DBClient {
	onceDB.Do(func() {
		initialize()
	})
	return client
}

func initialize() {
	var c DBConnector
	if viper.GetString("driver") == "mysql" {
		c = NewMySQLConnector()
	}

	var err error
	cfg := config.Instance
	client, err = c.Connect(cfg.Database)
	if err != nil {
		fmt.Println(err)
	}
}
