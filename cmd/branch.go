package cmd

import (
	"github.com/itsllyaz/gommit/utils"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Show the current branch name",
	Long: `✨ Gommit - Smart Commit Assistant

This command shows the current branch name and allows you to switch between branches.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.HandleBranch()
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
}