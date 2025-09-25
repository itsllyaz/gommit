package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var testCmd = &cobra.Command{
	Use: "test",
	Short: "Get short test functinality....",
	Long: "Get long test ....",

	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("Hello world this is test")
	},
}

func init(){
	rootCmd.AddCommand(testCmd)
}
