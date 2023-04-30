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
	"errors"
	"fmt"
	"sync"

	"github.com/terakoya76/populator/config"
)

// DBClient is an interface for DB Querying.
type DBClient interface {
	CreateTable(cfg *config.Table) error
	DropTable(cfg *config.Table) error
	Populate(cfg *config.Table) error
}

var client DBClient

// onceDB is used for Mutex Lock when initializing an instance.
var onceDB sync.Once

// DB provides instance of DB client.
func DB() DBClient {
	onceDB.Do(func() {
		initialize()
	})

	return client
}

func initialize() {
	var err error

	cfg := config.Instance
	client, err = BuildClient(cfg.Database)

	if err != nil {
		fmt.Println(err)
	}
}

// BuildClient builds DBClient for abstraction.
func BuildClient(cfg *config.Database) (DBClient, error) {
	switch cfg.Driver {
	case "mysql":
		if err := SetupMySQLDB(cfg); err != nil {
			return nil, err
		}

		return BuildMySQLClient(cfg)
	default:
		return nil, errors.New("not supported database driver")
	}
}
