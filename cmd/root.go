package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitbruhh",
	Short: "gitbruhh is a robust CLI for GitHub",
	Long:  `A modular and intelligent CLI tool to fetch information from GitHub without leaving your terminal.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Root flags can be defined here
}
