/*
Copyright © 2022 Junjie Yang junjie@jyang.dev
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/JumjieYang/nav/data"
	"github.com/spf13/cobra"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "this cmdlet is to register a remote webpage",
	Long:  `this cmdlet is to register a remote webpage.`,
	Run: func(cmd *cobra.Command, args []string) {
		if newWebPage == "" && deleteWebPage == "" && !listWeb {
			cmd.Help()
			return
		}

		if newWebPage != "" || deleteWebPage != "" {
			if webSource == "" {
				cmd.Help()
				return
			}
		}

		if newWebPage != "" {
			registerWeb(newWebPage, webSource)
			return
		}

		if deleteWebPage != "" {
			deleteWeb(deleteRepoName)
			return
		}

		if listWeb {
			listAllWeb()
		}
	},
}

func registerWeb(webName, webSource string) {

	data.InsertLink(webName, webSource, "web")
}

func deleteWeb(webName string) {
	data.DeleteLink(webName, "web")
}

func listAllWeb() {
	res, err := data.ListLink("web")

	if err != nil {
		log.Fatalln(err)
	}

	for _, s := range res {
		fmt.Println(s)
	}
}

var newWebPage string
var webSource string
var deleteWebPage string
var listWeb bool

func init() {
	regCmd.AddCommand(webCmd)
	webCmd.Flags().StringVarP(&newWebPage, "new", "n", "", "the alias of target repo")
	webCmd.Flags().StringVarP(&webSource, "source", "s", "", "the source of the repo")
	webCmd.Flags().StringVarP(&deleteWebPage, "remove", "r", "", "delete the alias")
	webCmd.Flags().BoolVarP(&listWeb, "list", "l", false, "list the repos")

	webCmd.MarkFlagsMutuallyExclusive("new", "remove", "list")
}
