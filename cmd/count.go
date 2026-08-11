package cmd

import (
	"fmt"

	"github.com/itsllyaz/gommit/utils"
	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Count the total commits in the current branch",
	Long: `✨ Gommit - Smart Commit Assistant

Displays the total number of commits in the current branch.`,
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Println("Counting commits...")
		}
		utils.HandleCount()
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
}