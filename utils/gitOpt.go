package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func IsGitRepo() error {

	// first check if the current working directory is contain .git directory
	wd, _ := os.Getwd()
	fullpath := fmt.Sprintf("%v/.git", wd)
	_, err := os.Stat(fullpath)
	if err != nil{
		return err
	}



	return nil
}

func GetCurrentBranch() (string, error) {
	// run shell command(git command)
	//str := "git rev-parse --abbrev-ref HEAD"
	// str:= "ls"

	//cmd := exec.Command(str)
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")


	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil{

		return "", err

	}

	return strings.TrimSpace(string(out)), err


}

func GetGitOrigin () (string, error){
 cmd := exec.Command("git", "remote", "get-url", "origin")

    cmd.Stderr = os.Stderr

    out, err := cmd.Output()
    if err != nil{
    	return "", err
    }

   return strings.TrimSpace(string(out)), nil
}

func GetRepoInfo() (string,error){

    cmd := exec.Command("git", "remote", "get-url", "origin")

    cmd.Stderr = os.Stderr

    out, err := cmd.Output()
    if err != nil{
    	return "", err
    }

   return strings.TrimSpace(string(out)), nil

   //  err := cmd.Run()
   //  if err != nil {
   //  return  nil, err
   // // os.Exit(1)
   //  }
   //  return nil
}

func GitStatus(){
	// check if .git found
	err := IsGitRepo()
	if err != nil{
		fmt.Println("❌ This is not a git repository")
	}

	// get current branch
}


func GitTotalCommit() (string, error){

    cmd := exec.Command("git", "rev-list", "--count", "HEAD")
    out, err := cmd.Output()
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(string(out)),nil
}

func GitLastCommit() (string, error){
 cmd := exec.Command("git", "log", "-1", `--pretty=format:%h - %s (%an, %ar)`)
    out, err := cmd.Output()
    if err != nil {
       //  fmt.Println("Error:", err)
        return "", err
    }

   // fmt.Println("Latest commit:", strings.TrimSpace(string(out)))
   return strings.TrimSpace(string(out)), err
}



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
		fmt.Printf("👉 %v  current\n", output)
	}
}
