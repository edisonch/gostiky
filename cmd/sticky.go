/*
Copyright © 2021 Edison Ch <edisonch@yahoo.com>

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

// stickyCmd represents the sticky command
var stickyCmd = &cobra.Command{
	Use:   "sticky",
	Short: "To make file immutable/appendable or to undo it",
	Long: `To prevent file being modified or deleted, we make it IMMUTABLE
mode into the file or recursively in directory.
To make file/directory appendable, we make it APPEND mode into the file or recursively in directory.
	
gostiky sticky -d (directory) -f (suffix file/dir) -m (mode)
	
For example:
	gostiky sticky -d /home/test  -f .php -m IMMUTABLE
	
	Mode available to use is as follow:
		APPEND : mode that only append to file
		UNAPPEND : mode that undo append 
		IMMUTABLE : mode that makes file to be immune to any command
		UNIMMUTABLE : mode that undo 
		IMMUTE_RECURSIVE : find the directory and apply the mode to all files
		UNIMMUTE_RECURSIVE : mode that undo command above
		APPEND_RECURSIVE : find the directory and apply the mode to all files
		UNAPPEND_RECURSIVE : mode that undo command above
	.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sticky called")

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
		val.StickyMode = mode_sticky

		err = util.StickIt(val)
		if err != nil {
			panic(err)
		}
		util.ClearBashHistory()
	},
}

func init() {
	stickyCmd.Flags().StringP("dir", "d", "", "search directory")
	stickyCmd.Flags().StringP("suffix", "f", "",
		"suffix file/dir")
	stickyCmd.Flags().StringP("mode", "m", "", "mode of sticky")
	stickyCmd.MarkFlagRequired("dir")
	stickyCmd.MarkFlagRequired("suffix")
	stickyCmd.MarkFlagRequired("mode")
	rootCmd.AddCommand(stickyCmd)
}
