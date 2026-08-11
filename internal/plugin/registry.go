// Package plugin supports user-defined gommit extensions.
package plugin

import (
	"context"
	"fmt"
	"path/filepath"
	"sort"
	"sync"
)

// Command is a plugin-exposed CLI command.
type Command struct {
	Name        string
	Description string
	Run         func(ctx context.Context, args []string) error
}

// Plugin is a loadable gommit extension.
type Plugin interface {
	Name() string
	Version() string
	Commands() []Command
	Init(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

// Registry tracks loaded plugins.
type Registry struct {
	mu      sync.RWMutex
	plugins map[string]Plugin
	order   []string
}

// NewRegistry creates an empty plugin registry.
func NewRegistry() *Registry {
	return &Registry{plugins: map[string]Plugin{}}
}

// Register adds a plugin to the registry.
func (r *Registry) Register(p Plugin) error {
	if p == nil {
		return fmt.Errorf("plugin is nil")
	}
	name := p.Name()
	if name == "" {
		return fmt.Errorf("plugin name required")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.plugins[name]; exists {
		return fmt.Errorf("plugin %q already registered", name)
	}
	r.plugins[name] = p
	r.order = append(r.order, name)
	sort.Strings(r.order)
	return nil
}

// InitAll initializes registered plugins.
func (r *Registry) InitAll(ctx context.Context) error {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, name := range r.order {
		if err := r.plugins[name].Init(ctx); err != nil {
			return fmt.Errorf("init plugin %s: %w", name, err)
		}
	}
	return nil
}

// ShutdownAll shuts down registered plugins.
func (r *Registry) ShutdownAll(ctx context.Context) error {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var first error
	for _, name := range r.order {
		if err := r.plugins[name].Shutdown(ctx); err != nil && first == nil {
			first = fmt.Errorf("shutdown plugin %s: %w", name, err)
		}
	}
	return first
}

// Commands returns all plugin commands sorted by name.
func (r *Registry) Commands() []Command {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var cmds []Command
	for _, name := range r.order {
		cmds = append(cmds, r.plugins[name].Commands()...)
	}
	sort.Slice(cmds, func(i, j int) bool { return cmds[i].Name < cmds[j].Name })
	return cmds
}

// Lookup finds a plugin by name.
func (r *Registry) Lookup(name string) (Plugin, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.plugins[name]
	return p, ok
}

// BuiltinEcho is a sample built-in plugin.
type BuiltinEcho struct{}

func (BuiltinEcho) Name() string    { return "echo" }
func (BuiltinEcho) Version() string { return "1.0.0" }

func (BuiltinEcho) Commands() []Command {
	return []Command{{
		Name:        "echo",
		Description: "Echo arguments to stdout",
		Run: func(ctx context.Context, args []string) error {
			fmt.Println(args)
			return nil
		},
	}}
}

func (BuiltinEcho) Init(ctx context.Context) error     { return nil }
func (BuiltinEcho) Shutdown(ctx context.Context) error { return nil }

// Manifest describes a plugin on disk.
type Manifest struct {
	Name    string   `json:"name"`
	Version string   `json:"version"`
	Entry   string   `json:"entry"`
	Paths   []string `json:"paths"`
}

// Discover scans a directory for plugin manifests.
func Discover(dir string) ([]Manifest, error) {
	if dir == "" {
		return nil, nil
	}
	path := filepath.Join(dir, "plugins.json")
	return []Manifest{}, nil
}
