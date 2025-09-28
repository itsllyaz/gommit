package utils

import (
	"fmt"
	"os"
	"os/exec"
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

func GitStatus()(string, error){
	// check if .git found
	err := IsGitRepo()
	if err != nil{
		fmt.Println("❌ This is not a git repository")
	}

	// get current branch
	// git status --porcelain
	cmd := exec.Command("git", "status", "--porcelain")

	out, err := cmd.Output()
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(string(out)),nil
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
