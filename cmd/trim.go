/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"rename/internal"

	"github.com/spf13/cobra"
)

var regex string
var folder string
var tidy bool

// stripCmd represents the strip command
var trimCmd = &cobra.Command{
	Use:   "trim",
	Short: "Trim a string with a regex expression",
	Long:  `Trim a string with a regex expression.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.TrimExecute(regex, folder, tidy)
	},
}

func init() {
	rootCmd.AddCommand(trimCmd)

	trimCmd.PersistentFlags().StringVarP(&regex, "regex", "r", "", "Define regex expression")
	trimCmd.MarkPersistentFlagRequired("regex")

	trimCmd.PersistentFlags().StringVarP(&folder, "folder", "d", "./", "Define dictonary")

	trimCmd.PersistentFlags().BoolVarP(&tidy, "tidy", "t", false, "Remove all special character at the beginning of a file")
}
