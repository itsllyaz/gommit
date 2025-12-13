package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type MessageFlow struct{
	TypeOfChange string
	TheChange string
	WhyItChanged string
	ScopeAffected string
	BreakingChange string
}

var container []MessageFlow
func HandleGitCommitMessage() []MessageFlow {
    typeOfChangeReader := bufio.NewReader(os.Stdin)
    fmt.Print(`
🔧 What type of change is this?
   Choose one of the following:
   🛠️  Features ------- New capability or enhancement
   🐞  Bug fix  ------- Fixes broken or unexpected behavior
   🔄  Refactor ------- Internal code cleanup, no behavior change
   📚  Documentation -- Updates to README, comments, or guides
   🧪  Test     ------- Adds or updates tests
   🧹  Chore    ------- Maintenance tasks like dependency updates
👉 Type your choice: `)
    typeOfChangeInput, _ := typeOfChangeReader.ReadString('\n')
    typeOfChangeInput = strings.TrimSpace(typeOfChangeInput)

    theChangeReader := bufio.NewReader(os.Stdin)
    fmt.Print("\n✏️  Briefly describe what you changed or added:\n👉 ")

    theChangeInput, _ := theChangeReader.ReadString('\n')
    theChangeInput = strings.TrimSpace(theChangeInput)

    whyItChangedReader := bufio.NewReader(os.Stdin)
    fmt.Print("\n💡 Why did you make this change? (Optional)\n👉 ")

    whyItChangedInput, _ := whyItChangedReader.ReadString('\n')
    whyItChangedInput = strings.TrimSpace(whyItChangedInput)

    scopeAffectedReader := bufio.NewReader(os.Stdin)
    fmt.Print("\n📍 Which part of the codebase does this affect? (Optional)\n👉 ")

    scopeAffectedInput, _ := scopeAffectedReader.ReadString('\n')
    scopeAffectedInput = strings.TrimSpace(scopeAffectedInput)

    breakingChangeReader := bufio.NewReader(os.Stdin)
    fmt.Print("\n⚠️  Does this change break anything or require updates elsewhere?\n👉 ")

    breakingChangeInput, _ := breakingChangeReader.ReadString('\n')
    breakingChangeInput = strings.TrimSpace(breakingChangeInput)

    container = append(container,
        MessageFlow{
            TypeOfChange:    typeOfChangeInput,
            TheChange:       theChangeInput,
            WhyItChanged:    whyItChangedInput,
            ScopeAffected:   scopeAffectedInput,
            BreakingChange:  breakingChangeInput,
        },
    )

    fmt.Println("\n✅ Change details captured successfully!")

    return container
}


func FormatPrompt(info []MessageFlow) string {
    return fmt.Sprintf(
        "Generate a commit message for the following change(max ~60 characters):\nType: %s\nChange: %s\nReason: %s\nScope: %s\nBreaking change: %s",
        info[0].TypeOfChange,
        info[0].TheChange,
        info[0].WhyItChanged,
        info[0].ScopeAffected,
        info[0].BreakingChange,
    )
}
