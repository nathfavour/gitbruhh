package cmd

import (
	"fmt"
	"os"

	"github.com/nathfavour/gitbruhh/pkg/gh"
	"github.com/nathfavour/gitbruhh/pkg/ui"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user [username]",
	Short: "Get information about a GitHub user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		client := gh.NewClient(gh.ProviderType(provider))
		user, err := client.GetUser(username)
		if err != nil {
			fmt.Printf("Error fetching user %s: %v\n", username, err)
			os.Exit(1)
		}
		ui.PrintUser(user)
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}

