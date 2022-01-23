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
	"github.com/spf13/cobra"
)

// projectnameCmd represents the projectname command
var projectnameCmd = &cobra.Command{
	Use:   "projectname",
	Short: "reset project name",
	Long: `allows you to reset project name. For example : 

			sparkle projectname "thunderbird"
	
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			cmd.Help()
			return
		}
		if len(args) == 1 {
			err := datafile.SetProjectName(args[0])
			if err != nil {
				cmd.Help()
			}
			cleanup()
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(projectnameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectnameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectnameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
