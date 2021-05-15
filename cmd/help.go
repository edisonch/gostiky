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

	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "to show this help",
	Long: `If you see this, it means you need guide how to use this program.
	gostiky needs 3 parameters to run and they are as follow:
	--dir    or -d : directory to search to files immutable
	--suffix or -f : suffix of file/directory that you want to apply mode for
	--mode   or -m : mode of file/folder that you want to do action
	
	MODE available to use is as follow:
		APPEND : mode that only append to file
		UNAPPEND : mode that undo append 
		IMMUTABLE : mode that makes file to be immune to any command
		UNIMMUTABLE : mode that undo 
		IMMUTE_RECURSIVE : find the directory and apply the mode to all files
		UNIMMUTE_RECURSIVE : mode that undo command above
		APPEND_RECURSIVE : find the directory and apply the mode to all files
		UNAPPEND_RECURSIVE : mode that undo command above.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("help called")
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
