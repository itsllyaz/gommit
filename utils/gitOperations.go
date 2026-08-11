package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func IsGitRepo() error {
	wd, _ := os.Getwd()
	fullpath := fmt.Sprintf("%v/.git", wd)
	_, err := os.Stat(fullpath)
	if err != nil {
		return err
	}
	return nil
}

func GetCurrentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func GetGitOrigin() (string, error) {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func GitStatus() (string, error) {
	if err := IsGitRepo(); err != nil {
		return "", err
	}
	cmd := exec.Command("git", "status", "--porcelain")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func GitTotalCommit() (string, error) {
	cmd := exec.Command("git", "rev-list", "--count", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func GitLastCommit() (string, error) {
	cmd := exec.Command("git", "log", "-1", "--pretty=format:%h - %s (%an, %ar)")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func GetRepoInfo() (string, error) {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}