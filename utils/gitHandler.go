package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func HandleInfo(){
	PrintHeader("📊 Gommit - Repository Information")

	err := IsGitRepo()
	if err != nil{
		fmt.Println("❌ This is not a git repository")
	}

	// repository name
	wd, _ := os.Getwd()
	repoName := filepath.Base(wd)
	fmt.Printf("📁 Repository: %v\n", repoName)

	// remote origin
	output, err := GetGitOrigin()
	if err != nil{
		fmt.Println("🌐 Origin: Not configured")

	}else{
		fmt.Printf("🌐 Origin: %v\n", output)
	}

	// current branch
	output, err = GetCurrentBranch()
	if err != nil{
		fmt.Println("Couldn't retrive the branch")
	}else{
		fmt.Printf("📍 Current branch: %v\n", output)
	}

	// last commit
	output, err = GitLastCommit()
	if err != nil{
		fmt.Println("📝 Last commit: No commits yet")
	}else{
		fmt.Printf("📝 Last commit: %v\n", output)
	}

	// total commit
	output, err = GitTotalCommit()
	if err != nil{
		fmt.Println("📈 Total commits: 0")
	}else{
		fmt.Printf("📈 Total commits: %v\n", output)
	}


}


func HandleBranch(){
	PrintHeader("🌿 Gommit - Branch Information")
	err := IsGitRepo()
	if err != nil{
		fmt.Println("❌ This is not a git repository")

	}

	output, err := GetCurrentBranch()
	if err != nil{
		fmt.Println("No Branch found")
	} else{
		fmt.Printf("👉 *%v*  current branch\n", output)
	}
}



func HandleStatus(){
	PrintHeader("🌿 Gommit - Status")

	// get current branch
	output, err := GetCurrentBranch()
	if err != nil{
		fmt.Println("No Branch found")
	} else{
		fmt.Printf("📍 current branch  %v\n", output)
	}

	// getting status
	output, err = GitStatus()
	if err != nil{

		fmt.Println("working directory clean")
	}else{
		fmt.Println("📝 Changes detected TO:")
		color.Red(" %v", output)
		//fmt.Printf("📝TO:  %v\n", output)

	}
}
