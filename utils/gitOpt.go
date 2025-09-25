package utils

import (
	"fmt"
	"os"
	"os/exec"
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

func GetCurrentBranch() error {
	// run shell command(git command)
	//str := "git rev-parse --abbrev-ref HEAD"
	// str:= "ls"

	//cmd := exec.Command(str)
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil{
		//fmt.Println("there is not branch found")
		return err
		//os.Exit(1)
	}

	return nil


}

func GetRepoInfo() error{
	err := IsGitRepo()
	if err != nil{
		fmt.Println("❌ This is not a git repository")
	}

    cmd := exec.Command("git", "remote", "get-url", "origin")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err = cmd.Run()
    if err != nil {
    return  err
   // os.Exit(1)
    }
    return nil
}

func GitStatus(){
	// check if .git found
	err := IsGitRepo()
	if err != nil{
		fmt.Println("❌ This is not a git repository")
	}

	// get current branch
}
