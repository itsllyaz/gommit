// Package ai provides multi-provider commit message generation.
package ai

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Provider generates commit messages from repository context.
type Provider interface {
	Name() string
	Generate(ctx context.Context, req Request) (Response, error)
}

// Request carries context for commit message generation.
type Request struct {
	Branch     string
	Status     string
	LastCommit string
	Remote     string
	Diff       string
	Template   string
	Locale     string
}

// Response is a generated commit message result.
type Response struct {
	Message   string
	Provider  string
	Cached    bool
	Latency   time.Duration
	Tokens    int
	CreatedAt time.Time
}

// Engine orchestrates providers and offline cache.
type Engine struct {
	providers []Provider
	cache     *Cache
	mu        sync.RWMutex
}

// NewEngine creates an AI engine with optional cache directory.
func NewEngine(cacheDir string, providers ...Provider) (*Engine, error) {
	cache, err := NewCache(cacheDir)
	if err != nil {
		return nil, err
	}
	return &Engine{providers: providers, cache: cache}, nil
}

// Register adds a provider to the engine.
func (e *Engine) Register(p Provider) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.providers = append(e.providers, p)
}

// Providers returns registered provider names.
func (e *Engine) Providers() []string {
	e.mu.RLock()
	defer e.mu.RUnlock()
	names := make([]string, len(e.providers))
	for i, p := range e.providers {
		names[i] = p.Name()
	}
	return names
}

// Generate uses the named provider or the first available provider.
func (e *Engine) Generate(ctx context.Context, providerName string, req Request) (Response, error) {
	key := cacheKey(providerName, req)
	if cached, ok := e.cache.Get(key); ok {
		cached.Cached = true
		return cached, nil
	}

	p, err := e.pick(providerName)
	if err != nil {
		return Response{}, err
	}

	start := time.Now()
	resp, err := p.Generate(ctx, req)
	if err != nil {
		return Response{}, err
	}
	resp.Provider = p.Name()
	resp.Latency = time.Since(start)
	resp.CreatedAt = time.Now().UTC()
	_ = e.cache.Set(key, resp)
	return resp, nil
}

func (e *Engine) pick(name string) (Provider, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if len(e.providers) == 0 {
		return nil, fmt.Errorf("no AI providers registered")
	}
	if name == "" {
		return e.providers[0], nil
	}
	for _, p := range e.providers {
		if p.Name() == name {
			return p, nil
		}
	}
	return nil, fmt.Errorf("provider %q not found", name)
}

func cacheKey(provider string, req Request) string {
	h := sha256.New()
	fmt.Fprintf(h, "%s|%s|%s|%s|%s", provider, req.Branch, req.Status, req.LastCommit, req.Diff)
	return hex.EncodeToString(h.Sum(nil))
}

// Cache stores offline commit message responses.
type Cache struct {
	dir string
	mu  sync.RWMutex
}

// NewCache creates or opens a cache directory.
func NewCache(dir string) (*Cache, error) {
	if dir == "" {
		dir = filepath.Join(os.TempDir(), "gommit-ai-cache")
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	return &Cache{dir: dir}, nil
}

// Get loads a cached response by key.
func (c *Cache) Get(key string) (Response, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	data, err := os.ReadFile(filepath.Join(c.dir, key+".json"))
	if err != nil {
		return Response{}, false
	}
	var resp Response
	if err := json.Unmarshal(data, &resp); err != nil {
		return Response{}, false
	}
	return resp, true
}

// Set stores a response in cache.
func (c *Cache) Set(key string, resp Response) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(c.dir, key+".json"), data, 0o644)
}

// Clear removes all cache entries.
func (c *Cache) Clear() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	entries, err := os.ReadDir(c.dir)
	if err != nil {
		return err
	}
	for _, e := range entries {
		_ = os.Remove(filepath.Join(c.dir, e.Name()))
	}
	return nil
}

// PromptBuilder builds structured prompts from git context.
type PromptBuilder struct {
	Prefix  string
	Include []string
}

// Build renders a prompt from a request.
func (b *PromptBuilder) Build(req Request) string {
	var sb strings.Builder
	if b.Prefix != "" {
		sb.WriteString(b.Prefix)
		sb.WriteString("\n")
	}
	sb.WriteString("Generate a concise, conventional commit message.\n")
	if req.Branch != "" {
		sb.WriteString("Branch: ")
		sb.WriteString(req.Branch)
		sb.WriteString("\n")
	}
	if req.Remote != "" {
		sb.WriteString("Remote: ")
		sb.WriteString(req.Remote)
		sb.WriteString("\n")
	}
	if req.Status != "" {
		sb.WriteString("Status:\n")
		sb.WriteString(req.Status)
		sb.WriteString("\n")
	}
	if req.LastCommit != "" {
		sb.WriteString("Last commit: ")
		sb.WriteString(req.LastCommit)
		sb.WriteString("\n")
	}
	if req.Diff != "" {
		sb.WriteString("Diff:\n")
		sb.WriteString(req.Diff)
		sb.WriteString("\n")
	}
	if req.Template != "" {
		sb.WriteString("Template: ")
		sb.WriteString(req.Template)
		sb.WriteString("\n")
	}
	return sb.String()
}

// StaticProvider returns deterministic messages for tests/offline mode.
type StaticProvider struct {
	Label string
}

func (s StaticProvider) Name() string {
	if s.Label == "" {
		return "static"
	}
	return s.Label
}

func (s StaticProvider) Generate(ctx context.Context, req Request) (Response, error) {
	select {
	case <-ctx.Done():
		return Response{}, ctx.Err()
	default:
	}
	msg := fmt.Sprintf("feat: update %s", strings.TrimSpace(req.Branch))
	if msg == "feat: update " {
		msg = "feat: update repository"
	}
	return Response{Message: msg, Tokens: len(msg)}, nil
}

// GeminiProvider wraps Google Gemini generation (stub when no SDK call).
type GeminiProvider struct {
	APIKey string
	Model  string
}

func (g GeminiProvider) Name() string { return "gemini" }

func (g GeminiProvider) Generate(ctx context.Context, req Request) (Response, error) {
	if g.APIKey == "" {
		return Response{}, fmt.Errorf("gemini: missing API key")
	}
	prompt := (&PromptBuilder{Prefix: "Gemini"}).Build(req)
	msg := fmt.Sprintf("feat(scope): %s", summarize(prompt))
	return Response{Message: msg, Tokens: len(msg)}, nil
}

// OpenAIProvider stub for OpenAI-compatible APIs.
type OpenAIProvider struct {
	APIKey string
	Model  string
	Base   string
}

func (o OpenAIProvider) Name() string { return "openai" }

func (o OpenAIProvider) Generate(ctx context.Context, req Request) (Response, error) {
	if o.APIKey == "" {
		return Response{}, fmt.Errorf("openai: missing API key")
	}
	msg := fmt.Sprintf("fix: %s", firstLine(req.Status))
	if msg == "fix: " {
		msg = "fix: resolve issue"
	}
	return Response{Message: msg, Tokens: len(msg)}, nil
}

// AnthropicProvider stub for Anthropic APIs.
type AnthropicProvider struct {
	APIKey string
	Model  string
}

func (a AnthropicProvider) Name() string { return "anthropic" }

func (a AnthropicProvider) Generate(ctx context.Context, req Request) (Response, error) {
	if a.APIKey == "" {
		return Response{}, fmt.Errorf("anthropic: missing API key")
	}
	msg := fmt.Sprintf("chore: maintain %s", req.Branch)
	return Response{Message: msg, Tokens: len(msg)}, nil
}

func summarize(s string) string {
	s = strings.TrimSpace(s)
	if len(s) > 48 {
		return s[:48]
	}
	return s
}

func firstLine(s string) string {
	if idx := strings.IndexByte(s, '\n'); idx >= 0 {
		return strings.TrimSpace(s[:idx])
	}
	return strings.TrimSpace(s)
}
