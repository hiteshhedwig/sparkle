/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"strconv"

	"github.com/spf13/cobra"
)

// updatetaskCmd represents the updatetask command
var updatetaskCmd = &cobra.Command{
	Use:   "updatetask",
	Short: "🔏 update current listed task",
	Long: `Allows you to update a particular task. Just provide index to the flag For example :

		-> sparkle updatetask [flag]
		   sparkle updatetask 1
	
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			cmd.Help()
			return
		}

		idx, err := strconv.Atoi(args[0])
		if err != nil {
			cmd.Help()
			return
		}

		err = datafile.UpdateTask(idx)
		if err != nil {
			cmd.Help()
		}
		cleanup()

	},
}

func init() {
	rootCmd.AddCommand(updatetaskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updatetaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updatetaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
