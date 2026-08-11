package cmd

import (
	"fmt"

	"github.com/itsllyaz/gommit/utils"
	"github.com/spf13/cobra"
)

var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Show the remote repository URL",
	Long: `✨ Gommit - Smart Commit Assistant

Displays the remote repository URL.`,
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Println("Fetching remote URL...")
		}
		utils.HandleRemote()
	},
}

func init() {
	rootCmd.AddCommand(remoteCmd)
}