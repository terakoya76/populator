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
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/terakoya76/populator/config"
	"github.com/terakoya76/populator/database"
)

func init() {
	rootCmd.AddCommand(populateCmd)
}

var populateCmd = &cobra.Command{
	Use:   "populate",
	Short: "Populate given tables' w/ seed data",
	Long:  "Populate given tables' w/ seed data",
	Run: func(cmd *cobra.Command, args []string) {
		if err := populate(); err != nil {
			fmt.Println(err)
		}
	},
}

func populate() error {
	db := database.DB()
	cfg := config.Instance
	for _, table := range cfg.Tables {
		if err := db.CreateTable(table); err != nil {
			return err
		}

		db.Insert(table)
	}

	return nil
}
