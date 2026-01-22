package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var provider string

var rootCmd = &cobra.Command{
	Use:   "xoy",
	Short: "xoy is a robust CLI for GitHub",
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
	rootCmd.PersistentFlags().StringVarP(&provider, "provider", "p", "auto", "Provider to use (auto, api, scraper)")
}
