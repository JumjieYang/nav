/*
Copyright © 2022 Junjie Yang junjie@jyang.dev
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// regCmd represents the reg command
var regCmd = &cobra.Command{
	Use:   "reg",
	Short: "this cmdlet is to register a local directory or a web page to CLI",
	Long:  "this cmdlet is to register a local directory or a web page to CLI",

	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(regCmd)

}
