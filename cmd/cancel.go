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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cancelCmd represents the cancel command
var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "will cancel the particular task by index",
	Long: `This command helps you cancel a specific task or a list of task. It is achived via knowing passing index of the task(s)
	You can know index by : sparkle list --cat
	. For example:
		sparkle cancel 07
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cancel called")
		if len(args) == 0 {
			cmd.Help()
			return
		}
		//more than one args mean deleting multiple tasks!
		if len(args) > 0 {
			log.Info(args)
			idx, err := strconv.Atoi(args[0])
			if err != nil {
				log.Error("Please pass valid argument - ", err)
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
