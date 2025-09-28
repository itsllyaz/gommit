package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/llyas36/gommit/internal"
)

const openRouterURL = "https://openrouter.ai/api/v1/chat/completions"

type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type RequestBody struct {
    Model    string    `json:"model"`
    Messages []Message `json:"messages"`
}


type CompletionResponse struct {
    Choices []struct {
        Message struct {
            Content string `json:"content"`
        } `json:"message"`
    } `json:"choices"`
}




var commitInfo = internal.HandleGitCommitMessage()
var prompt = internal.FormatPrompt((commitInfo))

func HandleRequest() string{
    apiKey := os.Getenv("OPENROUTER_API_KEY") // Set your OpenRouter API key as an environment variable

    if apiKey == "" {
        fmt.Println("Please set the OPENROUTER_API_KEY environment variable.")
        os.Exit(1)
    }

    reqBody := RequestBody{
        Model: "cognitivecomputations/dolphin-mistral-24b-venice-edition:free",
        Messages: []Message{
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
    req.Header.Set("HTTP-Referer", "https://yourapp.example.com") // Replace with your site or CLI tool name
    req.Header.Set("X-Title", "My CLI Commit Generator")           // Optional: name your app

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response from Dolphin Mistral via OpenRouter:")
   // fmt.Println(string(body))


    return string(body)
}

func ExtractContent(jsonInput string) string {
    var response CompletionResponse
    err := json.Unmarshal([]byte(jsonInput), &response)
    if err != nil {
        return fmt.Sprintf("❌ Failed to parse JSON: %v", err)
    }

    if len(response.Choices) == 0 {
        return "⚠️ No choices found in response."
    }

    return response.Choices[0].Message.Content
}


// func main(){
// 	fmt.Println(">>>>")
// 	jsonContent := HandleRequest()
// 	content := ExtractContent(jsonContent)

// 	fmt.Println("..........")
// 	fmt.Println(content)
// 	fmt.Println("..........")

// }
