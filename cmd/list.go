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
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Prints out all the current task in the bucket",
	Long: `Prints out all the current task in the bucket
	For example:
		sparkle list   //Prints out all the tasks
		sparkle list --cat //Prints out all the tasks categorywise
	`,
	Run: func(cmd *cobra.Command, args []string) {
		allstatus, _ := cmd.Flags().GetBool("cat")
		//category, _ := cmd.Flags().GetBool("cat")
		if !allstatus {
			print("All status \n")
			datafile.ListAllCurrentTasks()
		} else {
			print("No status \n")
			datafile.ListTasksCatWise()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//listCmd.Flags().Bool("all", false, "Lists all the tasks in the bucket history")

	listCmd.Flags().Bool("cat", false, "Lists all the tasks category wise")

	//show list prioritywise. Assign priority! like @critical, @inoneHour something like these

}
