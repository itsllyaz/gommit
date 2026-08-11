// Package config loads global and repository-local gommit settings.
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	GlobalFileName = ".gommit/config.json"
	LocalFileName  = ".gommit.local.json"
)

// Profile represents a named configuration profile.
type Profile struct {
	Name       string            `json:"name"`
	AIProvider string            `json:"ai_provider"`
	AIModel    string            `json:"ai_model"`
	Quiet      bool              `json:"quiet"`
	Verbose    bool              `json:"verbose"`
	Plugins    []string          `json:"plugins"`
	Extra      map[string]string `json:"extra,omitempty"`
}

// Settings is the root configuration document.
type Settings struct {
	CurrentProfile string             `json:"current_profile"`
	Profiles       map[string]Profile `json:"profiles"`
	Defaults       map[string]string  `json:"defaults,omitempty"`
}

// Loader reads and merges config layers.
type Loader struct {
	HomeDir string
	RepoDir string
	mu      sync.RWMutex
	cache   *Settings
}

// NewLoader creates a config loader.
func NewLoader(home, repo string) *Loader {
	if home == "" {
		home, _ = os.UserHomeDir()
	}
	if repo == "" {
		repo, _ = os.Getwd()
	}
	return &Loader{HomeDir: home, RepoDir: repo}
}

// Load merges global, profile, and local settings.
func (l *Loader) Load() (*Settings, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	settings := defaultSettings()
	globalPath := filepath.Join(l.HomeDir, GlobalFileName)
	if err := mergeFile(globalPath, settings); err != nil {
		return nil, err
	}
	localPath := filepath.Join(l.RepoDir, LocalFileName)
	if err := mergeFile(localPath, settings); err != nil {
		return nil, err
	}
	l.cache = settings
	return settings, nil
}

// SaveGlobal writes settings to the global config path.
func (l *Loader) SaveGlobal(s *Settings) error {
	path := filepath.Join(l.HomeDir, GlobalFileName)
	return writeJSON(path, s)
}

// SaveLocal writes repository-local overrides.
func (l *Loader) SaveLocal(s *Settings) error {
	path := filepath.Join(l.RepoDir, LocalFileName)
	return writeJSON(path, s)
}

// Active returns the active profile or a default profile.
func (l *Loader) Active(s *Settings) Profile {
	if s == nil {
		return defaultSettings().Profiles["default"]
	}
	name := s.CurrentProfile
	if name == "" {
		name = "default"
	}
	if p, ok := s.Profiles[name]; ok {
		return p
	}
	return Profile{Name: name, AIProvider: "gemini"}
}

func defaultSettings() *Settings {
	return &Settings{
		CurrentProfile: "default",
		Profiles: map[string]Profile{
			"default": {
				Name:       "default",
				AIProvider: "gemini",
				AIModel:    "gemini-2.5-flash",
			},
			"offline": {
				Name:       "offline",
				AIProvider: "static",
				Quiet:      true,
			},
			"verbose": {
				Name:       "verbose",
				AIProvider: "gemini",
				Verbose:    true,
			},
		},
		Defaults: map[string]string{
			"locale": "en",
		},
	}
}

func mergeFile(path string, s *Settings) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	var overlay Settings
	if err := json.Unmarshal(data, &overlay); err != nil {
		return fmt.Errorf("parse %s: %w", path, err)
	}
	if overlay.CurrentProfile != "" {
		s.CurrentProfile = overlay.CurrentProfile
	}
	for k, v := range overlay.Profiles {
		s.Profiles[k] = v
	}
	for k, v := range overlay.Defaults {
		s.Defaults[k] = v
	}
	return nil
}

func writeJSON(path string, s *Settings) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// ParseYAML is a lightweight YAML-like parser for simple key: value files.
func ParseYAML(raw string) map[string]string {
	out := map[string]string{}
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		out[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}
	return out
}

// Validate checks settings for required fields.
func Validate(s *Settings) error {
	if s == nil {
		return fmt.Errorf("settings is nil")
	}
	if len(s.Profiles) == 0 {
		return fmt.Errorf("at least one profile required")
	}
	if _, ok := s.Profiles[s.CurrentProfile]; s.CurrentProfile != "" && !ok {
		return fmt.Errorf("current profile %q not found", s.CurrentProfile)
	}
	return nil
}

// EnvOverlay applies environment variable overrides.
func EnvOverlay(s *Settings) {
	if v := os.Getenv("GOMMIT_AI_PROVIDER"); v != "" {
		p := s.Profiles[s.CurrentProfile]
		p.AIProvider = v
		s.Profiles[s.CurrentProfile] = p
	}
	if v := os.Getenv("GOMMIT_PROFILE"); v != "" {
		s.CurrentProfile = v
	}
}
