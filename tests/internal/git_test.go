package cmd

import (
	"github.com/llyas36/gommit/utils"
	"github.com/spf13/cobra"
)




var gitCmd = &cobra.Command{


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
		//...
		// output, err := utils.GitTotalCommit()
		// if err != nil{
		// 	fmt.Println("🌐 Origin: Not configured")

		// }else{
		// 	fmt.Printf("🌐 Origin: %v\n", output)
		// }
		//
		utils.IsGitRepo()

	},

}
