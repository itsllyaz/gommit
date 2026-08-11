package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"google.golang.org/genai"
)

const (
	geminiAPIKeyEnv = "AI_API_KEY"
	openRouterURL   = "https://openrouter.ai/api/v1/chat/completions"
)

type GeminiResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// HandleGemini generates a commit message using Google Gemini API.
func HandleGemini() string {
	apiKey := os.Getenv(geminiAPIKeyEnv)
	if apiKey == "" {
		fmt.Println("Please set the AI_API_KEY environment variable.")
		os.Exit(1)
	}

	commitInfo := GetGitInfo()
	prompt := FormatPrompt(commitInfo)

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating Gemini client: %v\n", err)
		os.Exit(1)
	}

	result, err := client.Models.GenerateContent(ctx, "gemini-2.5-flash", genai.Text(prompt), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating content: %v\n", err)
		os.Exit(1)
	}

	return result.Text()
}

// HandleRequest generates a commit message using OpenRouter API.
func HandleRequest() string {
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		fmt.Println("Please set the OPENROUTER_API_KEY environment variable.")
		os.Exit(1)
	}

	commitInfo := GetGitInfo()
	prompt := FormatPrompt(commitInfo)

	reqBody := struct {
		Model    string `json:"model"`
		Messages []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"messages"`
	}{
		Model: "cognitivecomputations/dolphin-mistral-24b-venice-edition:free",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{Role: "user", Content: prompt},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", openRouterURL, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("HTTP-Referer", "https://yourapp.example.com")
	req.Header.Set("X-Title", "My CLI Commit Generator")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

// GetGitInfo returns git information needed for commit message generation.
func GetGitInfo() []string {
	var info []string
	branch, err := GetCurrentBranch()
	if err == nil && branch != "" {
		info = append(info, "branch: "+branch)
	}

	origin, err := GetGitOrigin()
	if err == nil && origin != "" {
		info = append(info, "remote: "+origin)
	}

	status, err := GitStatus()
	if err == nil && status != "" {
		info = append(info, "status: "+status)
	}

	lastCommit, err := GitLastCommit()
	if err == nil && lastCommit != "" {
		info = append(info, "last commit: "+lastCommit)
	}

	return info
}

// FormatPrompt formats the commit message prompt based on git information.
func FormatPrompt(commitInfo []string) string {
	var prompt strings.Builder
	prompt.WriteString("Generate a commit message for the following change:\n")
	for _, line := range commitInfo {
		prompt.WriteString(line + "\n")
	}
	prompt.WriteString("Type: feature\nChange: description\nReason: describe why\nScope: impact\nBreaking change: no\n")
	return prompt.String()
}