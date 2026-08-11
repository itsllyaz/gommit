# Usage Examples

## Branch information

```bash
gommit branch
gommit branch -v
```

## Repository status

```bash
gommit status
gommit count
gommit latest
gommit remote
```

## AI commit message

```bash
export AI_API_KEY=your-key
gommit commit
gommit commit -v
```

## Git operations (via internal runner)

Extended operations are available through the `internal/git` package and
generated `src/gitops` modules for merge, rebase, stash, cherry-pick, and bisect.

## Profiles

```json
{
  "current_profile": "offline",
  "profiles": {
    "offline": {
      "name": "offline",
      "ai_provider": "static",
      "quiet": true
    }
  }
}
```

## SDK

```go
svc := sdk.NewGitHubService(os.Getenv("GITHUB_TOKEN"))
repo, err := svc.GetRepo(ctx, "owner", "repo")
```
