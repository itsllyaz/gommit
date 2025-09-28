package cmd

import (
	"fmt"

	"github.com/llyas36/gommit/utils"
	"github.com/spf13/cobra"
)

// type MessageFlow struct{
// 	TypeOfChange string
// 	TheChange string
// 	WhyItChanged string
// 	ScopeAffected string
// 	BreakingChange string
// }

// var container []MessageFlow

// func HandleGitCommitMessage() []MessageFlow {
//     typeOfChangeReader := bufio.NewReader(os.Stdin)
//     fmt.Print(`
// 🔧 What type of change is this?
//    Choose one of the following:
//    🛠️  Features
//    🐞  Bug fix
//    🔄  Refactor
//    📚  Documentation
//    🧪  Test
//    🧹  Chore
// 👉 Type your choice: `)
//     typeOfChangeInput, _ := typeOfChangeReader.ReadString('\n')
//     typeOfChangeInput = strings.TrimSpace(typeOfChangeInput)

//     theChangeReader := bufio.NewReader(os.Stdin)
//     fmt.Print("\n✏️  Briefly describe what you changed or added:\n👉 ")

//     theChangeInput, _ := theChangeReader.ReadString('\n')
//     theChangeInput = strings.TrimSpace(theChangeInput)

//     whyItChangedReader := bufio.NewReader(os.Stdin)
//     fmt.Print("\n💡 Why did you make this change? (Optional)\n👉 ")

//     whyItChangedInput, _ := whyItChangedReader.ReadString('\n')
//     whyItChangedInput = strings.TrimSpace(whyItChangedInput)

//     scopeAffectedReader := bufio.NewReader(os.Stdin)
//     fmt.Print("\n📍 Which part of the codebase does this affect? (Optional)\n👉 ")

//     scopeAffectedInput, _ := scopeAffectedReader.ReadString('\n')
//     scopeAffectedInput = strings.TrimSpace(scopeAffectedInput)

//     breakingChangeReader := bufio.NewReader(os.Stdin)
//     fmt.Print("\n⚠️  Does this change break anything or require updates elsewhere?\n👉 ")

//     breakingChangeInput, _ := breakingChangeReader.ReadString('\n')
//     breakingChangeInput = strings.TrimSpace(breakingChangeInput)

//     container = append(container,
//         MessageFlow{
//             TypeOfChange:    typeOfChangeInput,
//             TheChange:       theChangeInput,
//             WhyItChanged:    whyItChangedInput,
//             ScopeAffected:   scopeAffectedInput,
//             BreakingChange:  breakingChangeInput,
//         },
//     )

//     fmt.Println("\n✅ Change details captured successfully!")

//     return container
// }

// func FormatPrompt(info []MessageFlow) string {
//     return fmt.Sprintf(
//         "Generate a commit message for the following change:\nType: %s\nChange: %s\nReason: %s\nScope: %s\nBreaking change: %s",
//         info[0].TypeOfChange,
//         info[0].TheChange,
//         info[0].WhyItChanged,
//         info[0].ScopeAffected,
//         info[0].BreakingChange,
//     )
// }
var showDetails bool
var commitCmd = &cobra.Command{
    Use:   "commit",
    Short: "📝 Create a new commit with style!",
    Long: `✨ Gommit - Smart Commit Assistant

This command helps you craft meaningful, well-structured commit messages.
You'll be guided through a series of prompts to describe:
  🔧 Type of change (Feature, Bug fix, etc.)
  ✏️  What was changed
  💡 Why it was changed
  📍 Scope of impact
  ⚠️  Any breaking changes

Let’s make your commit history beautiful and informative! 🚀`,


	Run: func(cmd *cobra.Command, args []string){

		if showDetails{
			PrintDetails()
			return
		}

		jsonContent := utils.HandleRequest()
		content := utils.ExtractContent(jsonContent)
		fmt.Println("\n✅ Your commit message is ready! Copy and use it below 👇")
		fmt.Printf("\n💬: %v\n",content)
	},
}


// Run: func(cmd *cobra.Command, args []string) {
//     // 🧠 Process the user input and generate commit message
//     jsonContent := utils.HandleRequest()
//     content := utils.ExtractContent(jsonContent)

//     // 📦 Display the final commit message
//     fmt.Println("\n✅ Your commit message is ready! Copy and use it below 👇")
//     fmt.Printf("\n💬 %s\n", content)
// },


func PrintDetails(){
	fmt.Println(`📌 **Commit Types Overview**

🛠️ **Feature** – New capability or enhancement
🐞 **Bug Fix** – Fixes broken or unexpected behavior
🔄 **Refactor** – Internal code cleanup, no behavior change
📚 **Documentation** – Updates to README, comments, or guides
🧪 **Test** – Adds or updates tests
🧹 **Chore** – Maintenance tasks like dependency updates

💡 Use these to guide your commit message structure!
`)
}


func init(){
	commitCmd.Flags().BoolVarP(&showDetails, "details", "d", false, "show commit type details and example")
	rootCmd.AddCommand(commitCmd)
}
