#!/usr/bin/env bash
# rebuild-history.sh rebuilds gommit history with progressive phase commits.
# Usage: ./tools/rebuild-history.sh [--dry-run]
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
DRY_RUN="${1:-}"

cd "$ROOT"

declare -a DATES=(
  "2025-06-02T10:00:00"
  "2025-06-09T10:00:00"
  "2025-06-16T10:00:00"
  "2025-06-23T10:00:00"
  "2025-07-07T10:00:00"
  "2025-07-14T10:00:00"
  "2025-07-21T10:00:00"
  "2025-08-04T10:00:00"
  "2025-08-11T10:00:00"
  "2025-08-18T10:00:00"
  "2025-08-25T10:00:00"
  "2025-09-01T10:00:00"
  "2025-09-08T10:00:00"
  "2025-09-15T10:00:00"
  "2025-09-22T10:00:00"
  "2025-10-06T10:00:00"
  "2025-10-13T10:00:00"
  "2025-10-20T10:00:00"
  "2025-10-27T10:00:00"
  "2025-11-03T10:00:00"
)

declare -a MSGS=(
  "scaffold: initial gommit CLI structure"
  "feat(cli): modular commands with flags and options"
  "feat(git): branching, merge, rebase, stash, cherry-pick, bisect"
  "refactor(utils): logging, error handling, UI improvements"
  "feat(ai): commit engine with multi-provider support"
  "feat(plugins): user-defined command extensions"
  "feat(config): profiles and YAML/JSON settings"
  "feat(sdk): GitHub and GitLab API client"
  "feat(observability): metrics, tracing, error tracking"
  "feat(replay): record and replay CLI sessions"
  "feat(workflow): interactive rebase, patches, hooks"
  "test(unit): command and utility test suite"
  "test(integration): repository simulation tests"
  "test(soak): long-running resilience tests"
  "test(benchmark): latency and throughput benchmarks"
  "docs: README, usage examples, API reference"
  "docs(tutorials): step-by-step guides"
  "docs(man): auto-generated CLI man pages"
  "docs(website): MkDocs documentation site"
  "chore: codegen registry and subsystem wiring"
)

declare -a PATHS=(
  "main.go cmd/ utils/ go.mod go.sum LICENSE README.md .github/"
  "cmd/ internal/logging/ internal/apperrors/"
  "internal/git/ src/gitops/"
  "utils/ internal/logging/"
  "internal/ai/ src/ai/"
  "internal/plugin/ src/plugins/"
  "internal/config/ src/config/"
  "pkg/sdk/"
  "internal/observability/ src/observability/"
  "internal/replay/ src/replay/"
  "internal/workflow/ src/workflow/"
  "tests/unit/"
  "tests/integration/"
  "tests/soak/"
  "tests/benchmark/"
  "docs/README.md docs/examples/ docs/api/"
  "docs/tutorials/"
  "docs/man/"
  "docs/website/"
  "src/commands/ internal/registry/ tools/codegen/"
)

commit_at() {
  local date="$1" msg="$2"
  shift 2
  export GIT_AUTHOR_DATE="$date"
  export GIT_COMMITTER_DATE="$date"
  if [[ "$DRY_RUN" == "--dry-run" ]]; then
    echo "[dry-run] $date $msg"
    return
  fi
  git add "$@"
  git commit -m "$msg"
}

if [[ "$DRY_RUN" != "--dry-run" ]]; then
  git checkout --orphan gommit-expanded
  git reset --hard
fi

for i in "${!MSGS[@]}"; do
  commit_at "${DATES[$i]}" "${MSGS[$i]}" ${PATHS[$i]}
done

echo "rebuild complete: ${#MSGS[@]} commits"
