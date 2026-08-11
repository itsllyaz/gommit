# Gommit Documentation

Gommit is a standalone Git assistant CLI with AI-powered commit messages,
extended Git workflows, plugin support, and forge integrations.

## Quick start

```bash
go build -o gommit .
./gommit branch
./gommit status
./gommit commit
```

## Subsystems

| Subsystem | Path | Description |
|-----------|------|-------------|
| CLI | `cmd/` | Cobra commands and flags |
| Git ops | `internal/git/`, `src/gitops/` | Branch, merge, rebase, stash |
| AI engine | `internal/ai/`, `src/ai/` | Multi-provider commit generation |
| Plugins | `internal/plugin/`, `src/plugins/` | User extensions |
| Config | `internal/config/`, `src/config/` | Profiles and settings |
| SDK | `pkg/sdk/` | GitHub/GitLab API client |
| Observability | `internal/observability/` | Metrics and tracing |
| Replay | `internal/replay/` | Session record/replay |
| Workflows | `internal/workflow/` | Advanced Git automation |

## Configuration

Global config: `~/.gommit/config.json`  
Local overrides: `.gommit.local.json`

Environment variables:

- `GOMMIT_PROFILE` — active profile name
- `GOMMIT_AI_PROVIDER` — override AI provider
- `AI_API_KEY` — Gemini API key

## Codegen

Regenerate subsystem modules:

```bash
go run ./tools/codegen
```

## Testing

```bash
go test ./...
```

## License

See LICENSE in repository root.
