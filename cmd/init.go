/*
Copyright © 2022 Junjie Yang junjie@jyang.dev
*/
package cmd

import (
	"github.com/JumjieYang/nav/data"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new database and table for nav CLI",
	Long:  "Initialize a new database and table for nav CLI",

	Run: func(cmd *cobra.Command, args []string) {
		data.CreateTable()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
