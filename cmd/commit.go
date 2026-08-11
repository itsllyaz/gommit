package cmd

import (
	"fmt"

	"github.com/itsllyaz/gommit/utils"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Generate a commit message using AI",
	Long: `✨ Gommit - Smart Commit Assistant

This command generates an AI-powered commit message based on your git status and repository context.`,
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Println("Generating commit message using AI...")
		}
		message := utils.HandleGemini()
		fmt.Printf("\n💬: %v\n", message)
		fmt.Println("\n✅ Your commit message is ready! Copy and use it below 👇")
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}