/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package tests

import (
	"github.com/llyas36/gommit/utils"
	"github.com/spf13/cobra"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{

	Run: func(cmd *cobra.Command, args []string) {
		utils.HandleBranch()

	},
}

// func init() {
// 	rootCmd.AddCommand(branchCmd)

// 	// Here you will define your flags and configuration settings.

// 	// Cobra supports Persistent Flags which will work for this command
// 	// and all subcommands, e.g.:
// 	// branchCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	// Cobra supports local flags which will only run when this command
// 	// is called directly, e.g.:
// 	// branchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
