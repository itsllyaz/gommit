// // utils/print.go
package utils

// import (
//     "fmt"
//     "github.com/fatih/color"
// )

// // Red prints red-colored text
// func Red(format string, a ...interface{}) {
//     color.New(color.FgRed).Printf(format+"\n", a...)
// }

// // Green prints green-colored text
// func Green(format string, a ...interface{}) {
//     color.New(color.FgGreen).Printf(format+"\n", a...)
// }

// // Yellow prints yellow-colored text
// func Yellow(format string, a ...interface{}) {
//     color.New(color.FgYellow).Printf(format+"\n", a...)
// }

// // Cyan prints cyan-colored text
// func Cyan(format string, a ...interface{}) {
//     color.New(color.FgCyan).Printf(format+"\n", a...)
// }

// import (
//     "github.com/yourusername/yourproject/utils"
// )

// output, err := GetCurrentBranch()
// if err != nil {
//     utils.Red("No Branch found")
// } else {
//     utils.Red("📍 current branch %v", output)
// }

// output, err = GitStatus()
// if err != nil {
//     utils.Green("working directory clean")
// } else {
//     utils.Yellow("📝 Changes detected:")
//     utils.Red("📝TO: %v", output)
// }
