// Package replay records and replays CLI sessions for debugging.
package replay

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Event is a single recorded CLI interaction.
type Event struct {
	Timestamp time.Time         `json:"timestamp"`
	Command   string            `json:"command"`
	Args      []string          `json:"args"`
	ExitCode  int               `json:"exit_code"`
	Output    string            `json:"output,omitempty"`
	Meta      map[string]string `json:"meta,omitempty"`
}

// Session is a recorded gommit CLI session.
type Session struct {
	ID        string    `json:"id"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at,omitempty"`
	Events    []Event   `json:"events"`
}

// Recorder captures session events.
type Recorder struct {
	session Session
	mu      sync.Mutex
	enabled bool
}

// NewRecorder starts a new session recorder.
func NewRecorder(id string) *Recorder {
	if id == "" {
		id = fmt.Sprintf("session-%d", time.Now().Unix())
	}
	return &Recorder{
		enabled: true,
		session: Session{ID: id, StartedAt: time.Now().UTC()},
	}
}

// Record appends an event to the session.
func (r *Recorder) Record(ev Event) {
	if !r.enabled {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if ev.Timestamp.IsZero() {
		ev.Timestamp = time.Now().UTC()
	}
	r.session.Events = append(r.session.Events, ev)
}

// Finish marks the session complete.
func (r *Recorder) Finish() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.session.EndedAt = time.Now().UTC()
}

// Session returns a copy of the current session.
func (r *Recorder) Session() Session {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.session
}

// Save writes the session to a JSON file.
func (r *Recorder) Save(path string) error {
	r.Finish()
	data, err := json.MarshalIndent(r.session, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// LoadSession reads a session from disk.
func LoadSession(path string) (Session, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Session{}, err
	}
	var s Session
	if err := json.Unmarshal(data, &s); err != nil {
		return Session{}, err
	}
	return s, nil
}

// Replayer simulates executing recorded events.
type Replayer struct {
	DryRun bool
}

// Replay walks session events and invokes handler for each.
func (rp *Replayer) Replay(s Session, handler func(Event) error) error {
	for i, ev := range s.Events {
		if handler == nil {
			fmt.Printf("[%d] %s %v\n", i, ev.Command, ev.Args)
			continue
		}
		if err := handler(ev); err != nil {
			return fmt.Errorf("event %d: %w", i, err)
		}
	}
	return nil
}

// Filter returns events matching command prefix.
func Filter(s Session, commandPrefix string) []Event {
	var out []Event
	for _, ev := range s.Events {
		if commandPrefix == "" || ev.Command == commandPrefix {
			out = append(out, ev)
		}
	}
	return out
}

// Stats summarizes a session.
type Stats struct {
	EventCount   int
	CommandCount map[string]int
	Duration     time.Duration
	Failures     int
}

func Summarize(s Session) Stats {
	st := Stats{CommandCount: map[string]int{}}
	st.EventCount = len(s.Events)
	for _, ev := range s.Events {
		st.CommandCount[ev.Command]++
		if ev.ExitCode != 0 {
			st.Failures++
		}
	}
	if !s.EndedAt.IsZero() {
		st.Duration = s.EndedAt.Sub(s.StartedAt)
	}
	return st
}
