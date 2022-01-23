/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"strconv"

	"github.com/spf13/cobra"
)

// cancelCmd represents the cancel command
var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "✖️ cancel the particular task by index",
	Long: `This command helps you cancel a specific task. Just pass the index in argument.
	You can know index by : sparkle list --cat
	. For example:
		sparkle cancel 7

	NOTE : it always cancel task from current list of task. If you wanna do remove task from other categories, consider
			sparkle purge --help
		`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		//more than one args mean deleting multiple tasks!
		if len(args) == 1 {
			idx, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Errorf("Please pass valid argument - %v", err)
				cmd.Help()
				return
			}
			datafile.Cancel(idx)

		}

	},
}

func init() {
	rootCmd.AddCommand(cancelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cancelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cancelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
