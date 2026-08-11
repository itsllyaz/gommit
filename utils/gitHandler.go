package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func PrintHeader(commandTitle string) {
	fmt.Println(Logo)
	fmt.Printf("     ***GOMMIT YOUR PERSONAL GIT ASSISTANT***\n\n")
	fmt.Printf("%v \n\n", commandTitle)
}

func HandleBranch() {
	PrintHeader("🌿 Gommit - Branch Information")
	err := IsGitRepo()
	if err != nil {
		fmt.Println("❌ This is not a git repository")
		return
	}
	output, err := GetCurrentBranch()
	if err != nil {
		fmt.Println("No Branch found")
	} else {
		fmt.Printf("👉 *%v*  current branch\n", output)
	}
}

func HandleInfo() {
	PrintHeader("📊 Gommit - Repository Information")
	err := IsGitRepo()
	if err != nil {
		fmt.Println("❌ This is not a git repository")
		return
	}

	wd, _ := os.Getwd()
	repoName := filepath.Base(wd)
	fmt.Printf("📁 Repository: %v\n", repoName)

	output, err := GetGitOrigin()
	if err != nil {
		fmt.Println("🌐 Origin: Not configured")
	} else {
		fmt.Printf("🌐 Origin: %v\n", output)
	}

	output, err = GetCurrentBranch()
	if err != nil {
		fmt.Println("Couldn't retrive the branch")
	} else {
		fmt.Printf("📍 Current branch: %v\n", output)
	}

	output, err = GitLastCommit()
	if err != nil {
		fmt.Println("📝 Last commit: No commits yet")
	} else {
		fmt.Printf("📝 Last commit: %v\n", output)
	}

	output, err = GitTotalCommit()
	if err != nil {
		fmt.Println("📈 Total commits: 0")
	} else {
		fmt.Printf("📈 Total commits: %v\n", output)
	}
}

func HandleStatus() {
	PrintHeader("🌿 Gommit - Status")
	output, err := GetCurrentBranch()
	if err != nil {
		fmt.Println("No Branch found")
		return
	}
	fmt.Printf("📍 current branch  %v\n", output)

	output, err = GitStatus()
	if err != nil {
		fmt.Println("❌ Could not retrieve status:")
		fmt.Println(err)
		return
	}
	if output == "" {
		fmt.Println("✅ Working directory clean")
	} else {
		fmt.Println("📝 Changes detected TO:")
		color.Red(" %v", output)
	}
}

func HandleLatest() {
	PrintHeader("📄 Gommit - Latest Commit")
	err := IsGitRepo()
	if err != nil {
		fmt.Println("❌ This is not a git repository")
		return
	}
	output, err := GitLastCommit()
	if err != nil {
		fmt.Println("📝 Last commit: No commits yet")
		return
	}
	fmt.Printf("📝 Latest commit: %v\n", output)
}

func HandleRemote() {
	PrintHeader("🌐 Gommit - Remote")
	err := IsGitRepo()
	if err != nil {
		fmt.Println("❌ This is not a git repository")
		return
	}
	output, err := GetGitOrigin()
	if err != nil {
		fmt.Println("🌐 Origin: Not configured")
	} else {
		fmt.Printf("🌐 Origin: %v\n", output)
	}
}

func HandleCount() {
	PrintHeader("📈 Gommit - Commit Count")
	err := IsGitRepo()
	if err != nil {
		fmt.Println("❌ This is not a git repository")
		return
	}
	output, err := GitTotalCommit()
	if err != nil {
		fmt.Println("📈 Total commits: 0")
		return
	}
	fmt.Printf("📈 Total commits: %v\n", output)
}