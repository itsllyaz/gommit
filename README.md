
---

## 🧰 Git Assistant CLI

A lightweight command-line tool to help you interact with Git more easily. Built with [Cobra](https://github.com/spf13/cobra), this assistant provides quick access to common Git tasks like checking the current branch, viewing the latest commit, and getting the remote URL.

### ✨ Features

- Show current branch name  
- Display latest commit info  
- Get remote repository URL  
- Count total commits in the current branch  

### 🚀 Installation

Clone the repo and build the binary:

```bash
git clone https://github.com/yourusername/gommit.git
cd gommit
go build -o gommit
```

### 🛠 Usage

```bash
./gommit branch     # Show current branch
./gommit latest     # Show latest commit
./gommit remote     # Show remote URL
./gommit count      # Show total number of commits
```

### 📦 Built With

- [Go](https://golang.org/)
- [Cobra](https://github.com/spf13/cobra)

---

Let me know if you want to add badges, license info, or examples with colored output!
