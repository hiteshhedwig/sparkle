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
	"bytes"
	"fmt"
	"os/exec"

	"github.com/kyokomi/emoji"
	"github.com/spf13/cobra"
)

// stashCmd represents the stash command
var stashCmd = &cobra.Command{
	Use:   "stash",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		datafile.ListAllCurrentTasks()

		files := args[0]
		var stdout bytes.Buffer
		var stderr bytes.Buffer

		command := fmt.Sprintf("git add %s", files)

		gitcmd := exec.Command("bash", "-c", command)
		err := gitcmd.Run()
		gitcmd.Stdout = &stdout
		gitcmd.Stderr = &stderr

		if err != nil {
			fmt.Println("Error ", err)
			return
		}
		disp := emoji.Sprintf(":hook: Files has been stashed. Please proceed with git commit -")
		fmt.Print(disp)
		fmt.Printf("\n")

	},
}

func init() {
	rootCmd.AddCommand(stashCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stashCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stashCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
