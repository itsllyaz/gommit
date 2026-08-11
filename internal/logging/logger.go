// Package logging provides structured logging helpers for gommit.
package logging

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// Level represents log severity.
type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// Logger writes leveled log lines.
type Logger struct {
	out   io.Writer
	level Level
	mu    sync.Mutex
}

// New creates a logger writing to stderr at info level.
func New() *Logger {
	return &Logger{out: os.Stderr, level: LevelInfo}
}

func (l *Logger) SetLevel(level Level) { l.level = level }
func (l *Logger) SetOutput(w io.Writer) { l.out = w }

func (l *Logger) log(level Level, format string, args ...any) {
	if level < l.level {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	ts := time.Now().UTC().Format(time.RFC3339)
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(l.out, "%s [%s] %s\n", ts, level.String(), msg)
}

func (l *Logger) Debug(format string, args ...any) { l.log(LevelDebug, format, args...) }
func (l *Logger) Info(format string, args ...any)  { l.log(LevelInfo, format, args...) }
func (l *Logger) Warn(format string, args ...any)  { l.log(LevelWarn, format, args...) }
func (l *Logger) Error(format string, args ...any) { l.log(LevelError, format, args...) }

// Default is the package-level logger.
var Default = New()
