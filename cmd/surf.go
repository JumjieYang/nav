/*
Copyright © 2022 Junjie Yang junjie@jyang.dev
*/
package cmd

import (
	"log"
	"os/exec"

	"github.com/JumjieYang/nav/data"
	"github.com/spf13/cobra"
)

// surfCmd represents the surf command
var surfCmd = &cobra.Command{
	Use:   "surf",
	Short: "This cmdlet will spawn a new browser page at the website",
	Long:  `This cmdlet will spawn a new browser page at the website.`,
	Run: func(cmd *cobra.Command, args []string) {
		source, err := data.GetLink(webName, "web")

		if err != nil {
			log.Fatalln(err)
			return
		}

		if source == "" {
			log.Fatalln("no such website")
			return
		}

		cmdlet := exec.Command("open", source)

		cmdlet.Run()
	},
}
var webName string

func init() {
	rootCmd.AddCommand(surfCmd)
	surfCmd.Flags().StringVarP(&webName, "name", "n", "", "web page to jump")
}
