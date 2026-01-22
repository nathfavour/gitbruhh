package cmd

import (
	"fmt"
	"os"

	"github.com/nathfavour/gitbruhh/pkg/gh"
	"github.com/nathfavour/gitbruhh/pkg/ui"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search for GitHub repositories",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := ""
		for _, arg := range args {
			query += arg + " "
		}
		client := gh.NewClient()
		repos, err := client.SearchRepos(query)
		if err != nil {
			fmt.Printf("Error searching for %s: %v\n", query, err)
			os.Exit(1)
		}
		ui.PrintRepos(repos)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

