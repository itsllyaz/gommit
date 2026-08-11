package cmd

import (
	"fmt"

	"github.com/itsllyaz/gommit/utils"
	"github.com/spf13/cobra"
)

var latestCmd = &cobra.Command{
	Use:   "latest",
	Short: "Show the latest commit in the current branch",
	Long: `✨ Gommit - Smart Commit Assistant

Displays the latest commit information for the current branch.`,
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Println("Fetching latest commit...")
		}
		utils.HandleLatest()
	},
}

func init() {
	rootCmd.AddCommand(latestCmd)
}