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
	"gostiky/util"

	"github.com/spf13/cobra"
)

// sqlinjectionCmd represents the sqlinjection command
var sqlinjectionCmd = &cobra.Command{
	Use:   "sqlinjection",
	Short: "scan php sql injection",
	Long: `Scan php sql injection by looking for php prepare statement
in coding with statement like INSERT,UPDATE,SELECT,DELETE. If those
	sql statement present, 
	then it would look for this pattern of string ->prepare(
	For example:
	go stiky sqlinjection -d <directory to recursively search> .`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sqlinjection called")

		var err error

		dir_flag, err := cmd.Flags().GetString("dir")
		if dir_flag == "" {
			panic(errors.New("must enter the search directory"))
		}
		suffix_flag, err := cmd.Flags().GetString("suffix")
		if suffix_flag == "" {
			panic(errors.New("must enter file/dir suffix"))
		}

		params := []string{dir_flag, suffix_flag}
		val, err := util.CountFiles(params, util.MODE_STICK_UNKNOWN)
		if err != nil {
			panic(err)
		}
		val.StickyMode = util.MODE_STICK_UNKNOWN

		result, err := util.ProcessFiles(val)
		if err != nil {
			panic(err)
		}
		for _, res := range result {
			println("no prepare statement at ", res)
		}
	},
}

func init() {
	rootCmd.AddCommand(sqlinjectionCmd)
	sqlinjectionCmd.Flags().StringP("dir", "d", "", "search directory")
	sqlinjectionCmd.Flags().StringP("suffix", "f", "",
		"suffix file/dir")
	sqlinjectionCmd.MarkFlagRequired("dir")
	sqlinjectionCmd.MarkFlagRequired("suffix")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sqlinjectionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sqlinjectionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
