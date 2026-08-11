/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gommit",
	Short: "Gommit is a minimal Git assistant CLI",
	Long: `✨ Gommit - Smart Commit Assistant

This command helps you craft meaningful, well-structured commit messages.
You'll be guided through a series of prompts to describe:
  🔧 Type of change (Feature, Bug fix, etc.)
  ✏️  What was changed
  💡 Why it was changed
  📍 Scope of impact
  ⚠️  Any breaking changes

Let's make your commit history beautiful and informative! 🚀`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

// persistentFlags defines flags available to all subcommands
var (
	quiet   bool
	verbose bool
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Suppress output")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Show verbose output")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}