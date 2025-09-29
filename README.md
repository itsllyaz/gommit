[![Go Report Card](https://goreportcard.com/badge/github.com/llyas36/gommit)](https://goreportcard.com/report/github.com/llyas36/gommit)  
![Latest Release](https://img.shields.io/github/v/release/llyas36/gommit)
---

##  Git Assistant CLI

A lightweight command-line tool to help you interact with Git more easily. Built with [Cobra](https://github.com/spf13/cobra), this assistant provides quick access to common Git tasks like checking the current branch, viewing the latest commit, and getting the remote URL.

### ✨ Features

- Show current branch name  
- Display latest commit info  
- Get remote repository URL  
- Count total commits in the current branch  
- Generate **AI-powered commit messages** for your changes  
- Get detailed repository status (staged, unstaged, untracked files)  


###  Installation

Clone the repo and build the binary:

```bash
git clone https://github.com/llyas36/gommit.git
cd gommit
go build -o gommit
```
Make it globally accessible (Optional)

If you want to use GOMMIT from anywhere in your terminal, move it to a directory in your $PATH:
```
sudo mv gommit /usr/local/bin/
```

### ⚙️ Setup

To enable the **AI-powered commit messages**, you’ll need an API key from [OpenRouter](https://openrouter.ai).

1. **Get an API key**  
   - Sign up at [OpenRouter](https://openrouter.ai).  
   - Generate a new API key from your dashboard.

2. **Export your API key as an environment variable**  

   ```bash
   export OPENROUTER_API_KEY="your_api_key_here"

3. ** Model used
   - GOMMIT uses the following model for commit message generation
  ```bash
  cognitivecomputations/dolphin-mistral-24b-venice-edition:free
``` 

### 🛠 Usage

Use GOMMIT commands in your Git repositories to quickly get information or perform tasks:

```bash
./gommit branch     # Show the current Git branch
./gommit commit     # Generate an AI-powered commit message or perform commits
./gommit info       # Get detailed repository info (repository, origin, last commit, total commit)
./gommit status     # Display the latest status 
./gommit git        # Check if the current directory is a Git repository
```
If you installed GOMMIT globally with go install, you can just run:
```
gommit branch
gommit commit
gommit info
gommit status
gommit git
```

### 📦 Built With

- [Go](https://golang.org/)
- [Cobra](https://github.com/spf13/cobra)

---



### Contributing

Contributions are welcome! If you see an issue, have a suggestion, or want to improve something:

Fork the repo

Create a new branch e.g. feature/awesome-new-stuff

Make your changes & commit them

Open a Pull Request

### 📄 License

This project is licensed under the MIT License. See the LICENSE


