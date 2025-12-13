/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/itsllyaz/gommit/utils"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "A brief description of your command",
	Long: `✨ Gommit - Smart Commit Assistant

This command helps you craft meaningful, well-structured commit messages.
You'll be guided through a series of prompts to describe:
  🔧 Type of change (Feature, Bug fix, etc.)
  ✏️  What was changed
  💡 Why it was changed
  📍 Scope of impact
  ⚠️  Any breaking changes

Let’s make your commit history beautiful and informative! 🚀`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.HandleStatus()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
