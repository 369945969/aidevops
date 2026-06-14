package handlers

import (
	"encoding/json"
	"sync"
)

// ============================================================
// PipelineEventBus — SSE event broadcasting for real-time progress
// ============================================================

// PipelineEvent represents a stage progress event pushed to clients.
type PipelineEvent struct {
	PipelineID uint   `json:"pipeline_id"`
	Stage      string `json:"stage"`
	Type       string `json:"type"` // stage_start, stage_done, stage_failed, pipeline_timeout, pipeline_completed, pipeline_failed, review_pause
	Message    string `json:"message"`
	DurationMs int64  `json:"duration_ms,omitempty"`
	Status     string `json:"status,omitempty"`
	Error      string `json:"error,omitempty"`
}

// PipelineEventBus broadcasts pipeline progress events to SSE subscribers.
type PipelineEventBus struct {
	mu      sync.RWMutex
	chans   map[uint][]chan PipelineEvent // pipelineID → list of subscriber channels
	global  []chan PipelineEvent          // global subscribers (all pipelines)
}

// NewPipelineEventBus creates a new event bus.
func NewPipelineEventBus() *PipelineEventBus {
	return &PipelineEventBus{
		chans:  make(map[uint][]chan PipelineEvent),
		global: []chan PipelineEvent{},
	}
}

// Subscribe creates a channel for receiving events about a specific pipeline.
// If pipelineID is 0, subscribes to all pipeline events.
func (bus *PipelineEventBus) Subscribe(pipelineID uint) chan PipelineEvent {
	ch := make(chan PipelineEvent, 64)
	bus.mu.Lock()
	defer bus.mu.Unlock()
	if pipelineID == 0 {
		bus.global = append(bus.global, ch)
	} else {
		bus.chans[pipelineID] = append(bus.chans[pipelineID], ch)
	}
	return ch
}

// Unsubscribe removes a channel from the event bus.
func (bus *PipelineEventBus) Unsubscribe(pipelineID uint, ch chan PipelineEvent) {
	bus.mu.Lock()
	defer bus.mu.Unlock()
	if pipelineID == 0 {
		for i, c := range bus.global {
			if c == ch {
				bus.global = append(bus.global[:i], bus.global[i+1:]...)
				close(ch)
				return
			}
		}
	} else {
		chans := bus.chans[pipelineID]
		for i, c := range chans {
			if c == ch {
				bus.chans[pipelineID] = append(chans[:i], chans[i+1:]...)
				close(ch)
				return
			}
		}
	}
}

// Publish sends an event to all subscribers of the given pipeline and global subscribers.
func (bus *PipelineEventBus) Publish(event PipelineEvent) {
	bus.mu.RLock()
	defer bus.mu.RUnlock()

	// Send to pipeline-specific subscribers
	for _, ch := range bus.chans[event.PipelineID] {
		select {
		case ch <- event:
		default: // skip if channel is full (client lagging)
		}
	}

	// Send to global subscribers
	for _, ch := range bus.global {
		select {
		case ch <- event:
		default:
		}
	}
}

// ToJSON converts the event to a JSON string for SSE data field.
func (e PipelineEvent) ToJSON() string {
	b, _ := json.Marshal(e)
	return string(b)
}
