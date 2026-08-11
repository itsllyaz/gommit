// Package observability provides metrics, tracing, and error tracking hooks.
package observability

import (
	"fmt"
	"sync"
	"time"
)

// Counter tracks monotonically increasing values.
type Counter struct {
	name  string
	value int64
	mu    sync.Mutex
}

func NewCounter(name string) *Counter { return &Counter{name: name} }

func (c *Counter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) Add(n int64) {
	c.mu.Lock()
	c.value += n
	c.mu.Unlock()
}

func (c *Counter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func (c *Counter) Name() string { return c.name }

// Histogram records latency samples in milliseconds.
type Histogram struct {
	name    string
	samples []float64
	mu      sync.Mutex
}

func NewHistogram(name string) *Histogram {
	return &Histogram{name: name, samples: make([]float64, 0, 128)}
}

func (h *Histogram) Observe(ms float64) {
	h.mu.Lock()
	h.samples = append(h.samples, ms)
	h.mu.Unlock()
}

func (h *Histogram) ObserveDuration(d time.Duration) {
	h.Observe(float64(d.Milliseconds()))
}

func (h *Histogram) Count() int {
	h.mu.Lock()
	defer h.mu.Unlock()
	return len(h.samples)
}

func (h *Histogram) P50() float64 { return h.percentile(0.5) }

func (h *Histogram) P95() float64 { return h.percentile(0.95) }

func (h *Histogram) percentile(p float64) float64 {
	h.mu.Lock()
	defer h.mu.Unlock()
	if len(h.samples) == 0 {
		return 0
	}
	idx := int(float64(len(h.samples)-1) * p)
	if idx < 0 {
		idx = 0
	}
	return h.samples[idx]
}

// Span represents a traced operation.
type Span struct {
	Name      string
	Start     time.Time
	End       time.Time
	Labels    map[string]string
	Events    []string
	Parent    *Span
	mu        sync.Mutex
}

func StartSpan(name string) *Span {
	return &Span{Name: name, Start: time.Now().UTC(), Labels: map[string]string{}}
}

func (s *Span) SetLabel(k, v string) {
	s.mu.Lock()
	s.Labels[k] = v
	s.mu.Unlock()
}

func (s *Span) AddEvent(msg string) {
	s.mu.Lock()
	s.Events = append(s.Events, msg)
	s.mu.Unlock()
}

func (s *Span) Finish() {
	s.mu.Lock()
	s.End = time.Now().UTC()
	s.mu.Unlock()
}

func (s *Span) Duration() time.Duration {
	s.mu.Lock()
	defer s.mu.Unlock()
	end := s.End
	if end.IsZero() {
		end = time.Now().UTC()
	}
	return end.Sub(s.Start)
}

// Tracer collects spans for a session.
type Tracer struct {
	spans []*Span
	mu    sync.Mutex
}

func NewTracer() *Tracer { return &Tracer{} }

func (t *Tracer) Record(span *Span) {
	t.mu.Lock()
	t.spans = append(t.spans, span)
	t.mu.Unlock()
}

func (t *Tracer) Spans() []*Span {
	t.mu.Lock()
	defer t.mu.Unlock()
	out := make([]*Span, len(t.spans))
	copy(out, t.spans)
	return out
}

// Registry holds named metrics.
type Registry struct {
	counters    map[string]*Counter
	histograms  map[string]*Histogram
	mu          sync.RWMutex
}

func NewRegistry() *Registry {
	return &Registry{
		counters:   map[string]*Counter{},
		histograms: map[string]*Histogram{},
	}
}

func (r *Registry) Counter(name string) *Counter {
	r.mu.Lock()
	defer r.mu.Unlock()
	if c, ok := r.counters[name]; ok {
		return c
	}
	c := NewCounter(name)
	r.counters[name] = c
	return c
}

func (r *Registry) Histogram(name string) *Histogram {
	r.mu.Lock()
	defer r.mu.Unlock()
	if h, ok := r.histograms[name]; ok {
		return h
	}
	h := NewHistogram(name)
	r.histograms[name] = h
	return h
}

// ErrorTracker records handled errors.
type ErrorTracker struct {
	errors []TrackedError
	mu     sync.Mutex
}

type TrackedError struct {
	When    time.Time
	Message string
	Command string
}

func NewErrorTracker() *ErrorTracker { return &ErrorTracker{} }

func (e *ErrorTracker) Track(command, message string) {
	e.mu.Lock()
	e.errors = append(e.errors, TrackedError{
		When: time.Now().UTC(), Message: message, Command: command,
	})
	e.mu.Unlock()
}

func (e *ErrorTracker) Recent(n int) []TrackedError {
	e.mu.Lock()
	defer e.mu.Unlock()
	if n <= 0 || n > len(e.errors) {
		n = len(e.errors)
	}
	start := len(e.errors) - n
	out := make([]TrackedError, n)
	copy(out, e.errors[start:])
	return out
}

// Report renders a text summary of metrics.
func (r *Registry) Report() string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var b string
	for name, c := range r.counters {
		b += fmt.Sprintf("counter %s=%d\n", name, c.Value())
	}
	for name, h := range r.histograms {
		b += fmt.Sprintf("histogram %s count=%d p50=%.2fms p95=%.2fms\n",
			name, h.Count(), h.P50(), h.P95())
	}
	return b
}
