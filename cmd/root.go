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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/terakoya76/populator/config"
	"github.com/terakoya76/populator/database"
)

var CfgFile string
var ReCreate bool

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "populator",
	Short: "Populate given tables' w/ seed data",
	Long:  "Populate given tables' w/ seed data",
	Run: func(_ *cobra.Command, _ []string) {
		if err := populate(); err != nil {
			fmt.Println(err)
		}
	},
}

func populate() error {
	db := database.DB()
	cfg := config.Instance

	for _, table := range cfg.Tables {
		if ReCreate {
			if err := db.DropTable(table); err != nil {
				return err
			}
		}

		if err := db.CreateTable(table); err != nil {
			return err
		}

		if err := db.Populate(table); err != nil {
			return err
		}
	}

	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
	} else {
		// find current working directory.
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(dir)
		viper.SetConfigName("populator")
		viper.SetConfigType("yaml")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		// yaml parsing error
		fmt.Println(err)
		os.Exit(1)
	}

	if err := LoadConfig(); err != nil {
		fmt.Printf("config file is invalid: %s", err)
		os.Exit(1)
	}

	config.Instance.CompleteWithDefault()
}

// LoadConfig assigns the configuration input to config.Instance.
func LoadConfig() error {
	err := viper.Unmarshal(&config.Instance)
	if err != nil {
		return err
	}

	return config.Instance.Validate()
}
