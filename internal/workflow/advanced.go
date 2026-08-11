// Package workflow implements advanced Git workflow helpers.
package workflow

import (
	"context"
	"fmt"
	"strings"

	"github.com/itsllyaz/gommit/internal/git"
)

// InteractiveRebase prepares an interactive rebase plan description.
type InteractiveRebase struct {
	Runner *git.Runner
	Steps  []RebaseStep
}

type RebaseStep struct {
	Action  string
	Commit  string
	Message string
}

// Plan builds a textual rebase plan from steps.
func (ir *InteractiveRebase) Plan() string {
	var b strings.Builder
	for _, step := range ir.Steps {
		fmt.Fprintf(&b, "%s %s", step.Action, step.Commit)
		if step.Message != "" {
			fmt.Fprintf(&b, " # %s", step.Message)
		}
		b.WriteString("\n")
	}
	return b.String()
}

// Start launches interactive rebase onto base.
func (ir *InteractiveRebase) Start(ctx context.Context, base string) error {
	if ir.Runner == nil {
		return fmt.Errorf("runner required")
	}
	return ir.Runner.Rebase(ctx, base, true)
}

// PatchManager handles patch series apply/export.
type PatchManager struct {
	Runner *git.Runner
	Series []string
}

func (pm *PatchManager) Export(ctx context.Context, commits []string) ([]string, error) {
	var patches []string
	for _, c := range commits {
		out, err := pm.Runner.Run(ctx, "show", c, "--format=email")
		if err != nil {
			return nil, err
		}
		patches = append(patches, out)
	}
	pm.Series = patches
	return patches, nil
}

func (pm *PatchManager) ApplyAll(ctx context.Context) error {
	for i, p := range pm.Series {
		if err := pm.Runner.ApplyPatch(ctx, p, false); err != nil {
			return fmt.Errorf("patch %d: %w", i, err)
		}
	}
	return nil
}

// HookManager installs and validates git hooks.
type HookManager struct {
	Runner *git.Runner
}

var defaultHooks = []string{"pre-commit", "commit-msg", "pre-push", "post-checkout"}

func (hm *HookManager) ValidateInstalled() (map[string]bool, error) {
	out := map[string]bool{}
	for _, name := range defaultHooks {
		path, err := hm.Runner.HookPath(name)
		if err != nil {
			return nil, err
		}
		_, err = hm.Runner.Run(context.Background(), "rev-parse", "--show-toplevel")
		out[name] = path != "" && err == nil
	}
	return out, nil
}

// FlowRunner coordinates multi-step git workflows.
type FlowRunner struct {
	Runner *git.Runner
	Steps  []FlowStep
}

type FlowStep struct {
	Name string
	Run  func(ctx context.Context, r *git.Runner) error
}

func (fr *FlowRunner) Execute(ctx context.Context) error {
	for i, step := range fr.Steps {
		if step.Run == nil {
			continue
		}
		if err := step.Run(ctx, fr.Runner); err != nil {
			return fmt.Errorf("step %d (%s): %w", i, step.Name, err)
		}
	}
	return nil
}

// ReleaseFlow is a sample workflow: fetch, rebase, push tag.
func ReleaseFlow(r *git.Runner, tag string) *FlowRunner {
	return &FlowRunner{
		Runner: r,
		Steps: []FlowStep{
			{Name: "fetch", Run: func(ctx context.Context, gr *git.Runner) error {
				return gr.Fetch(ctx, "origin", true)
			}},
			{Name: "status", Run: func(ctx context.Context, gr *git.Runner) error {
				_, err := gr.Status(ctx)
				return err
			}},
			{Name: "tag", Run: func(ctx context.Context, gr *git.Runner) error {
				if tag == "" {
					return nil
				}
				_, err := gr.Run(ctx, "tag", tag)
				return err
			}},
		},
	}
}

// ConflictHints suggests commands for common conflict states.
func ConflictHints(status string) []string {
	if !strings.Contains(status, "UU") && !strings.Contains(status, "AA") {
		return nil
	}
	return []string{
		"git status",
		"git diff",
		"git add <resolved-files>",
		"git rebase --continue OR git merge --continue",
		"git rebase --abort OR git merge --abort",
	}
}
