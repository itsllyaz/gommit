
---

##  Git Assistant CLI

A lightweight command-line tool to help you interact with Git more easily. Built with [Cobra](https://github.com/spf13/cobra), this assistant provides quick access to common Git tasks like checking the current branch, viewing the latest commit, and getting the remote URL.

###  Features

- Show current branch name  
- Display latest commit info  
- Get remote repository URL  
- Count total commits in the current branch
- AI-powered commit messages

###  Installation

Clone the repo and build the binary:

```bash
git clone https://github.com/itsllyaz/gommit.git
cd gommit
go build -o gommit
```
Make it globally accessible(Optional)
if you want to use GOMMIT from anywhere in you terminal, move it to a directory in you $PATH: 
```bash
    sudo mv gommit /usr/local/bin/
```
### Setup 
To enable the AI-powered commit messages, you'll need an API KEY, We have to options but the most easy one is from [Google AI Studio](https://aistudio.google.com/). 
  1. Get an API KEY
      * Sign up at [Google AI Studio](https://aistudio.google.com/)
      * Generate a new API KEY from you studio
  2. Export you API KEY as an environment variable
      ```bash
          export AI_API_KEY="your_api_key_here"
      ```
### Usage

```bash
./gommit branch     # Show current branch
./gommit latest     # Show latest commit
./gommit remote     # Show remote URL
./gommit count      # Show total number of commits
```
### Screenshots

<img width="1366" height="768" alt="gommit" src="https://github.com/user-attachments/assets/9ff580aa-c606-4888-8eb9-3523406a5026" />

<img width="1366" height="768" alt="gommit-branch" src="https://github.com/user-attachments/assets/278ffad6-cb73-4903-96db-416c309969c9" />

<img width="1366" height="768" alt="gommit-info" src="https://github.com/user-attachments/assets/af9dba47-14ec-4717-9ef7-f6af396ce8ee" />

<img width="1366" height="768" alt="gommit-status" src="https://github.com/user-attachments/assets/d9e9776e-cc47-42b9-8454-68fa6433ee60" />

###  Built With

- [Go](https://golang.org/)
- [Cobra](https://github.com/spf13/cobra)

---

made with :heart:


