package cmd

import (
	"fmt"
	"os"

	"github.com/nathfavour/gitbruhh/pkg/gh"
	"github.com/nathfavour/gitbruhh/pkg/ui"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo [owner/repo]",
	Short: "Get information about a GitHub repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoPath := args[0]
		client := gh.NewClient()
		repo, err := client.GetRepo(repoPath)
		if err != nil {
			fmt.Printf("Error fetching repository %s: %v\n", repoPath, err)
			os.Exit(1)
		}
		ui.PrintRepo(repo)
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
}

