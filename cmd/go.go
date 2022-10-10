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

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "This cmdlet will spawn a new text editor at the project root",
	Long: `This cmdlet will spawn a new text editor at the project root
		By default, new terminal will be open in new iterm tab
	`,
	Run: func(cmd *cobra.Command, args []string) {

		loc, err := data.GetLink(repoName, "repo")

		if err != nil {
			log.Fatalln(err)
			return
		}

		if loc == "" {
			log.Fatalln("no such repo")
			return
		}

		var cmdlet *exec.Cmd
		if editorName == "" || editorName == "code" {
			cmdlet = exec.Command("code", loc)
		} else {
			cmdlet = exec.Command("open", "-a", "iTerm", loc)
		}

		cmdlet.Run()
	},
}

var repoName string
var editorName string

func init() {
	rootCmd.AddCommand(goCmd)
	goCmd.Flags().StringVarP(&repoName, "name", "n", "", "repo name to jump")
	goCmd.Flags().StringVarP(&editorName, "editor", "e", "iTerm", "editor to open")
	goCmd.MarkFlagRequired("name")
}
