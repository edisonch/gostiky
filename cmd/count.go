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
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"gostiky/util"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "To count specific file/folder given the mode",
	Long: `Count is to count recursively given the suffic file or directory
	that we want to count. Here is the format
	
gostiky count -d (directory) -f (suffix file/dir) -m (mode)
	
For example: 
	./gosticky count -d /home/test/ -f .php -m IMMUTABLE 
	
Mode available to use is as follow:
	APPEND : mode that only append to file
	UNAPPEND : mode that undo append 
	IMMUTABLE : mode that makes file to be immune to any command
	UNIMMUTABLE : mode that undo 
	IMMUTE_RECURSIVE : find the directory and apply the mode to all files
	UNIMMUTE_RECURSIVE : mode that undo command above
	APPEND_RECURSIVE : find the directory and apply the mode to all files
	UNAPPEND_RECURSIVE : mode that undo command above
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("count called")


		var err error

		dir_flag, err := cmd.Flags().GetString("dir")
		if dir_flag == "" {
			panic(errors.New("must enter the search directory"))
		}
		suffix_flag, err := cmd.Flags().GetString("suffix")
		if suffix_flag == "" {
			panic(errors.New("must enter file/dir suffix"))
		}
		mode_flag,err := cmd.Flags().GetString("mode")
		if mode_flag == "" {
			panic(errors.New("must enter the mode"))
		}
		mode_sticky := util.ConvertCLIModeToInt(mode_flag)
		if mode_sticky == util.MODE_STICK_UNKNOWN {
			panic(errors.New("UNKNOWN STICKY MODE"))
		}

		params := []string{dir_flag,suffix_flag}
		val, err := util.CountFiles(params, mode_sticky)
		if err != nil {
			panic(err)
		}
		val.Print()
		val.PrintDetail()
	},
}

func init() {
	countCmd.Flags().StringP("dir","d","","search directory")
	countCmd.Flags().StringP("suffix","f","","suffix file/dir")
	countCmd.Flags().StringP("mode","m","","mode of sticky")
	countCmd.MarkFlagRequired("dir")
	countCmd.MarkFlagRequired("suffix")
	countCmd.MarkFlagRequired("mode")
	rootCmd.AddCommand(countCmd)
}
