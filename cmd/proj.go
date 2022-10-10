/*
Copyright © 2022 Junjie Yang junjie@jyang.dev
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/JumjieYang/nav/data"
	"github.com/spf13/cobra"
)

// projCmd represents the proj command
var projCmd = &cobra.Command{
	Use:   "proj",
	Short: "this cmdlet is to register a local directory",
	Long:  `this cmdlet is to register a local directory`,
	Run: func(cmd *cobra.Command, args []string) {
		if newRepoName == "" && deleteRepoName == "" && !listRepos {
			cmd.Help()
			return
		}

		if newRepoName != "" {
			registerRepo(newRepoName)
			return
		}

		if deleteRepoName != "" {
			deleteRepo(deleteRepoName)
			return
		}

		if listRepos {
			listAllRepo()
		}
	},
}

var newRepoName string
var deleteRepoName string
var listRepos bool

func registerRepo(repoName string) {
	path, err := os.Getwd()

	if err != nil {
		log.Fatalln(err)
	}

	data.InsertLink(repoName, path, "repo")
}

func deleteRepo(repoName string) {
	data.DeleteLink(repoName, "repo")
}

func listAllRepo() {
	res, err := data.ListLink("repo")

	if err != nil {
		log.Fatalln(err)
	}

	for _, s := range res {
		fmt.Println(s)
	}
}

func init() {
	regCmd.AddCommand(projCmd)
	projCmd.Flags().StringVarP(&newRepoName, "new", "n", "", "the alias of target repo")
	projCmd.Flags().StringVarP(&deleteRepoName, "remove", "r", "", "delete the alias")
	projCmd.Flags().BoolVarP(&listRepos, "list", "l", false, "list the repos")
	projCmd.MarkFlagsMutuallyExclusive("new", "remove", "list")
}
