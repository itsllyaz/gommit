package cmd

import (
	"github.com/itsllyaz/gommit/utils"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show git repository status",
	Long: `✨ Gommit - Smart Commit Assistant

This command displays the current branch and any uncommitted changes in the working directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.HandleStatus()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}