# API Reference

## Package `internal/git`

- `Runner` — executes git commands in a repository
- `Runner.Branch`, `Merge`, `Rebase`, `Stash`, `CherryPick`, `BisectStart`

## Package `internal/ai`

- `Engine` — orchestrates AI providers with offline cache
- `Provider` — interface for Gemini, OpenAI, Anthropic, static

## Package `internal/config`

- `Loader` — loads global and local JSON config
- `Profile` — named settings (provider, model, plugins)

## Package `internal/plugin`

- `Registry` — register and run plugin commands

## Package `pkg/sdk`

- `GitHubService` — repos, PRs, issues
- `GitLabService` — projects, merge requests

## Package `internal/observability`

- `Registry` — counters and histograms
- `Tracer` — span collection

## Package `internal/replay`

- `Recorder` — record CLI session events
- `Replayer` — replay sessions for debugging

## Package `internal/workflow`

- `InteractiveRebase`, `PatchManager`, `HookManager`, `FlowRunner`

## Generated modules

Codegen emits modules under `src/` and `tests/` registered via
`internal/registry`.
