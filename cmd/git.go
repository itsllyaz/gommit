package cmd

import (
	"fmt"
	"os"

	"github.com/llyas36/gommit/utils"
	"github.com/spf13/cobra"
)




var gitCmd = &cobra.Command{
	Use: "git",
	Short: "is this wd contain .git? ",
	Long: "to check if .git found",

	Run: func(cmd *cobra.Command, args[]string){
		// err := utils.IsGitRepo()
		// if err != nil{
		// 	fmt.Println(".git doesn't found")
		// }else{
		// 	fmt.Println(".git found")
		// }
		// ///......
		// err := utils.GetCurrentBranch()
		// if err != nil{
		// 	fmt.Fprintln(os.Stderr, "Error:", err)
		// 	os.Exit(1)
		// }
		//
		// ...
		err := utils.GetRepoInfo()
		if err != nil{
			fmt.Fprintln(os.Stderr, "Error: ", err)
			os.Exit(1)
		}

	},

}

func init(){
	rootCmd.AddCommand(gitCmd)
}
