/*
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
package main

import (
	"github.com/spf13/cobra"

	"github.com/terakoya76/populator/cmd"
	"github.com/terakoya76/populator/database"
)

func main() {
	cobra.OnInitialize(cmd.InitConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	cmd.RootCmd.PersistentFlags().StringVarP(&cmd.CfgFile, "config", "c", "", "config file (default is ./populator.yaml)")
	cmd.RootCmd.PersistentFlags().BoolVarP(&cmd.ReCreate, "recreate", "r", false, "drop tables then create them from scratch")
	cmd.RootCmd.PersistentFlags().BoolVarP(&database.Verbose, "verbose", "v", false, "show executed sql")
	cmd.RootCmd.DisableSuggestions = true

	cmd.Execute()
}
