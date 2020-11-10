/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

// mtoyCmd represents the mtoy command
var mtoyCmd = &cobra.Command{
	Use:   "mtoy",
	Short: "Meters to Yards",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Meters: " + args[0])
		meters, _ := strconv.ParseFloat(args[0], 64)
		yards := meters * 1.094
		fmt.Println("Yards: ", yards)
	},
}

func init() {
	rootCmd.AddCommand(mtoyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mtoyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mtoyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
