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
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// finishCmd represents the finish command
var finishCmd = &cobra.Command{
	Use:   "finish",
	Short: "will allow you to change status of a task to `finished`",
	Long: `For example: You can use below command to change the status of a task at index 1 

sparkle finish 1
	
Make sure you are providing index of the task you wish to finish.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("finish called")
		if len(args) == 0 {
			cmd.Help()
			return
		}
		//more than one args mean deleting multiple tasks!
		if len(args) > 0 {
			idx, err := strconv.Atoi(args[0])
			if err != nil {
				log.Error("Please pass valid argument - ", err)
				cmd.Help()
				return
			}
			err = datafile.CompletedTask(idx)
			if err != nil {
				cmd.Help()
			}
			cleanup()
		}

	},
}

func cleanup() {
	fmt.Printf("%s\n", string(colorWhite))
}
func init() {
	rootCmd.AddCommand(finishCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//finishCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// finishCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
