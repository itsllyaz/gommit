// Package git provides extended Git operations for gommit.
package git

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Runner executes git commands in a repository.
type Runner struct {
	Dir    string
	Env    []string
	DryRun bool
}

// NewRunner creates a git runner for the current or specified directory.
func NewRunner(dir string) *Runner {
	if dir == "" {
		dir, _ = os.Getwd()
	}
	return &Runner{Dir: dir}
}

// Run executes a git subcommand and returns combined output.
func (r *Runner) Run(ctx context.Context, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = r.Dir
	cmd.Env = append(os.Environ(), r.Env...)
	if r.DryRun {
		return fmt.Sprintf("dry-run: git %s", strings.Join(args, " ")), nil
	}
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, stderr.String())
	}
	return strings.TrimSpace(string(out)), nil
}

// Branch returns the current branch name.
func (r *Runner) Branch(ctx context.Context) (string, error) {
	return r.Run(ctx, "rev-parse", "--abbrev-ref", "HEAD")
}

// ListBranches returns local branch names.
func (r *Runner) ListBranches(ctx context.Context) ([]string, error) {
	out, err := r.Run(ctx, "branch", "--format=%(refname:short)")
	if err != nil {
		return nil, err
	}
	if out == "" {
		return nil, nil
	}
	return strings.Split(out, "\n"), nil
}

// CreateBranch creates and optionally checks out a branch.
func (r *Runner) CreateBranch(ctx context.Context, name string, checkout bool) error {
	args := []string{"branch", name}
	if checkout {
		args = []string{"checkout", "-b", name}
	}
	_, err := r.Run(ctx, args...)
	return err
}

// DeleteBranch deletes a local branch.
func (r *Runner) DeleteBranch(ctx context.Context, name string, force bool) error {
	flag := "-d"
	if force {
		flag = "-D"
	}
	_, err := r.Run(ctx, "branch", flag, name)
	return err
}

// Merge merges the given branch into HEAD.
func (r *Runner) Merge(ctx context.Context, branch string, noFF bool) error {
	args := []string{"merge", branch}
	if noFF {
		args = append(args, "--no-ff")
	}
	_, err := r.Run(ctx, args...)
	return err
}

// Rebase rebases current branch onto target.
func (r *Runner) Rebase(ctx context.Context, onto string, interactive bool) error {
	args := []string{"rebase", onto}
	if interactive {
		args = []string{"rebase", "-i", onto}
	}
	_, err := r.Run(ctx, args...)
	return err
}

// AbortRebase aborts an in-progress rebase.
func (r *Runner) AbortRebase(ctx context.Context) error {
	_, err := r.Run(ctx, "rebase", "--abort")
	return err
}

// Stash pushes working tree changes to stash.
func (r *Runner) Stash(ctx context.Context, message string) error {
	args := []string{"stash", "push"}
	if message != "" {
		args = append(args, "-m", message)
	}
	_, err := r.Run(ctx, args...)
	return err
}

// StashPop applies and removes the latest stash entry.
func (r *Runner) StashPop(ctx context.Context) error {
	_, err := r.Run(ctx, "stash", "pop")
	return err
}

// StashList returns stash entries.
func (r *Runner) StashList(ctx context.Context) ([]string, error) {
	out, err := r.Run(ctx, "stash", "list")
	if err != nil {
		return nil, err
	}
	if out == "" {
		return nil, nil
	}
	return strings.Split(out, "\n"), nil
}

// CherryPick applies a commit by hash.
func (r *Runner) CherryPick(ctx context.Context, commit string) error {
	_, err := r.Run(ctx, "cherry-pick", commit)
	return err
}

// BisectStart starts a bisect session.
func (r *Runner) BisectStart(ctx context.Context) error {
	_, err := r.Run(ctx, "bisect", "start")
	return err
}

// BisectGood marks commit as good.
func (r *Runner) BisectGood(ctx context.Context, commit string) error {
	args := []string{"bisect", "good"}
	if commit != "" {
		args = append(args, commit)
	}
	_, err := r.Run(ctx, args...)
	return err
}

// BisectBad marks commit as bad.
func (r *Runner) BisectBad(ctx context.Context, commit string) error {
	args := []string{"bisect", "bad"}
	if commit != "" {
		args = append(args, commit)
	}
	_, err := r.Run(ctx, args...)
	return err
}

// BisectReset ends bisect session.
func (r *Runner) BisectReset(ctx context.Context) error {
	_, err := r.Run(ctx, "bisect", "reset")
	return err
}

// Status returns porcelain status output.
func (r *Runner) Status(ctx context.Context) (string, error) {
	return r.Run(ctx, "status", "--porcelain")
}

// Log returns formatted log lines.
func (r *Runner) Log(ctx context.Context, count int, format string) ([]string, error) {
	if format == "" {
		format = "%h %s"
	}
	args := []string{"log", fmt.Sprintf("--pretty=format:%s", format)}
	if count > 0 {
		args = append(args, fmt.Sprintf("-n%d", count))
	}
	out, err := r.Run(ctx, args...)
	if err != nil {
		return nil, err
	}
	if out == "" {
		return nil, nil
	}
	return strings.Split(out, "\n"), nil
}

// RemoteURL returns origin URL if configured.
func (r *Runner) RemoteURL(ctx context.Context, name string) (string, error) {
	if name == "" {
		name = "origin"
	}
	return r.Run(ctx, "remote", "get-url", name)
}

// CommitCount returns total commits reachable from HEAD.
func (r *Runner) CommitCount(ctx context.Context) (int, error) {
	out, err := r.Run(ctx, "rev-list", "--count", "HEAD")
	if err != nil {
		return 0, err
	}
	var n int
	_, err = fmt.Sscanf(out, "%d", &n)
	return n, err
}

// IsRepo checks whether Dir contains a git repository.
func (r *Runner) IsRepo(ctx context.Context) bool {
	_, err := r.Run(ctx, "rev-parse", "--git-dir")
	return err == nil
}

// Diff returns diff output for optional pathspec.
func (r *Runner) Diff(ctx context.Context, cached bool, pathspec ...string) (string, error) {
	args := []string{"diff"}
	if cached {
		args = append(args, "--cached")
	}
	args = append(args, pathspec...)
	return r.Run(ctx, args...)
}

// ApplyPatch applies a patch file or stdin content.
func (r *Runner) ApplyPatch(ctx context.Context, patch string, cached bool) error {
	args := []string{"apply"}
	if cached {
		args = append(args, "--cached")
	}
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = r.Dir
	cmd.Stdin = strings.NewReader(patch)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git apply: %w: %s", err, stderr.String())
	}
	return nil
}

// HookPath returns the path to a git hook by name.
func (r *Runner) HookPath(name string) (string, error) {
	gitDir, err := r.Run(context.Background(), "rev-parse", "--git-dir")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/hooks/%s", gitDir, name), nil
}

// WorktreeList lists registered worktrees.
func (r *Runner) WorktreeList(ctx context.Context) ([]string, error) {
	out, err := r.Run(ctx, "worktree", "list", "--porcelain")
	if err != nil {
		return nil, err
	}
	var paths []string
	for _, line := range strings.Split(out, "\n") {
		if strings.HasPrefix(line, "worktree ") {
			paths = append(paths, strings.TrimPrefix(line, "worktree "))
		}
	}
	return paths, nil
}

// TagList returns tag names.
func (r *Runner) TagList(ctx context.Context) ([]string, error) {
	out, err := r.Run(ctx, "tag", "-l")
	if err != nil {
		return nil, err
	}
	if out == "" {
		return nil, nil
	}
	return strings.Split(out, "\n"), nil
}

// Fetch fetches from remote.
func (r *Runner) Fetch(ctx context.Context, remote string, prune bool) error {
	if remote == "" {
		remote = "origin"
	}
	args := []string{"fetch", remote}
	if prune {
		args = append(args, "--prune")
	}
	_, err := r.Run(ctx, args...)
	return err
}

// Pull pulls from remote branch.
func (r *Runner) Pull(ctx context.Context, remote, branch string) error {
	target := "origin"
	if remote != "" {
		target = remote
	}
	if branch != "" {
		target = target + " " + branch
	}
	_, err := r.Run(ctx, "pull", target)
	return err
}

// Push pushes to remote.
func (r *Runner) Push(ctx context.Context, remote, refspec string, force bool) error {
	args := []string{"push"}
	if force {
		args = append(args, "--force-with-lease")
	}
	if remote == "" {
		remote = "origin"
	}
	args = append(args, remote)
	if refspec != "" {
		args = append(args, refspec)
	}
	_, err := r.Run(ctx, args...)
	return err
}

// ConfigGet reads a git config value.
func (r *Runner) ConfigGet(ctx context.Context, key string) (string, error) {
	return r.Run(ctx, "config", "--get", key)
}

// ConfigSet writes a git config value.
func (r *Runner) ConfigSet(ctx context.Context, key, value string, global bool) error {
	args := []string{"config"}
	if global {
		args = append(args, "--global")
	}
	args = append(args, key, value)
	_, err := r.Run(ctx, args...)
	return err
}

// OperationSummary summarizes recent git activity.
type OperationSummary struct {
	Branch      string
	CommitCount int
	StatusLines int
	Remotes     []string
	UpdatedAt   time.Time
}

// Summarize collects repository summary information.
func (r *Runner) Summarize(ctx context.Context) (*OperationSummary, error) {
	s := &OperationSummary{UpdatedAt: time.Now().UTC()}
	var err error
	s.Branch, err = r.Branch(ctx)
	if err != nil {
		return nil, err
	}
	s.CommitCount, err = r.CommitCount(ctx)
	if err != nil {
		return nil, err
	}
	status, err := r.Status(ctx)
	if err != nil {
		return nil, err
	}
	if status != "" {
		s.StatusLines = len(strings.Split(status, "\n"))
	}
	remotes, err := r.Run(ctx, "remote")
	if err == nil && remotes != "" {
		s.Remotes = strings.Split(remotes, "\n")
	}
	return s, nil
}
