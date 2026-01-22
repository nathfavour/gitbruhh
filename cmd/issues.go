package cmd

import (
	"fmt"
	"os"

	"github.com/nathfavour/gitbruhh/pkg/gh"
	"github.com/nathfavour/gitbruhh/pkg/ui"
	"github.com/spf13/cobra"
)

var issuesCmd = &cobra.Command{
	Use:   "issues [owner/repo]",
	Short: "Get recent issues for a GitHub repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoPath := args[0]
		client := gh.NewClient()
		issues, err := client.GetIssues(repoPath)
		if err != nil {
			fmt.Printf("Error fetching issues for %s: %v\n", repoPath, err)
			os.Exit(1)
		}
		ui.PrintIssues(issues)
	},
}

func init() {
	rootCmd.AddCommand(issuesCmd)
}

