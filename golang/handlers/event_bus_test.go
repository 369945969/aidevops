package handlers

import (
	"encoding/json"
	"sync"
	"testing"
)

func TestNewPipelineEventBus(t *testing.T) {
	bus := NewPipelineEventBus()
	if bus.chans == nil {
		t.Error("chans map should be initialized")
	}
	if bus.global == nil {
		t.Error("global slice should be initialized")
	}
}

func TestSubscribeAndUnsubscribe(t *testing.T) {
	bus := NewPipelineEventBus()
	ch := bus.Subscribe(1)

	if len(bus.chans[1]) != 1 {
		t.Errorf("chans[1] length = %d, want 1", len(bus.chans[1]))
	}

	bus.Unsubscribe(1, ch)
	if len(bus.chans[1]) != 0 {
		t.Errorf("after unsubscribe, chans[1] length = %d, want 0", len(bus.chans[1]))
	}

	// Verify channel is closed
	_, ok := <-ch
	if ok {
		t.Error("channel should be closed after unsubscribe")
	}
}

func TestSubscribeGlobal(t *testing.T) {
	bus := NewPipelineEventBus()
	ch := bus.Subscribe(0)
	if len(bus.global) != 1 {
		t.Errorf("global length = %d, want 1", len(bus.global))
	}

	bus.Unsubscribe(0, ch)
	if len(bus.global) != 0 {
		t.Errorf("after unsubscribe global length = %d, want 0", len(bus.global))
	}
}

func TestPublishDeliversToSubscribers(t *testing.T) {
	bus := NewPipelineEventBus()
	ch1 := bus.Subscribe(1)
	ch2 := bus.Subscribe(1)

	event := PipelineEvent{
		PipelineID: 1,
		Stage:      "architect",
		Type:       "stage_done",
		Message:    "completed",
		Status:     "completed",
		DurationMs: 5000,
	}

	bus.Publish(event)

	e1 := <-ch1
	e2 := <-ch2

	if e1.Type != "stage_done" || e2.Type != "stage_done" {
		t.Error("both subscribers should receive the event")
	}
	if e1.PipelineID != 1 || e2.PipelineID != 1 {
		t.Error("pipeline ID should match")
	}
}

func TestPublishDoesNotDeliverToOtherPipelines(t *testing.T) {
	bus := NewPipelineEventBus()
	ch := bus.Subscribe(2)

	bus.Publish(PipelineEvent{PipelineID: 1, Type: "stage_done"})

	select {
	case <-ch:
		t.Error("subscriber of pipeline 2 should not receive pipeline 1 events")
	default:
		// expected — no event
	}
}

func TestPublishAlsoDeliversToGlobal(t *testing.T) {
	bus := NewPipelineEventBus()
	specific := bus.Subscribe(1)
	global := bus.Subscribe(0)

	bus.Publish(PipelineEvent{PipelineID: 1, Type: "pipeline_started"})

	<-specific // should receive
	<-global   // should also receive
}

func TestPublishSkipsFullChannel(t *testing.T) {
	bus := NewPipelineEventBus()
	ch := bus.Subscribe(1)

	// Fill the channel (buffer is 64)
	for i := 0; i < 64; i++ {
		ch <- PipelineEvent{Type: "overflow"}
	}

	// This should not block — Publish uses select/default
	bus.Publish(PipelineEvent{PipelineID: 1, Type: "should_skip"})

	// Verify we can only drain 64 (the 65th was skipped)
	count := 0
	for {
		select {
		case <-ch:
			count++
		default:
			goto done
		}
	}
done:
	if count != 64 {
		t.Errorf("drained %d events from full channel, want 64 (1 should be skipped)", count)
	}
}

func TestConcurrentPublish(t *testing.T) {
	bus := NewPipelineEventBus()
	ch := bus.Subscribe(1)
	var wg sync.WaitGroup

	// Use fewer events than channel buffer (64) to guarantee no drops
	const numEvents = 50
	for i := 0; i < numEvents; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			bus.Publish(PipelineEvent{PipelineID: 1, Type: "event", Message: "hello"})
		}(i)
	}

	wg.Wait()

	count := 0
	for {
		select {
		case <-ch:
			count++
		default:
			goto done2
		}
	}
done2:
	if count != numEvents {
		t.Errorf("received %d events, want %d from concurrent publish", count, numEvents)
	}
}

func TestPipelineEventToJSON(t *testing.T) {
	e := PipelineEvent{
		PipelineID: 42,
		Stage:      "tester",
		Type:       "stage_failed",
		Message:    "tests failed",
		DurationMs: 3000,
		Status:     "failed",
		Error:      "assertion error",
	}
	jsonStr := e.ToJSON()
	if !json.Valid([]byte(jsonStr)) {
		t.Error("ToJSON should produce valid JSON")
	}
	var parsed PipelineEvent
	if err := json.Unmarshal([]byte(jsonStr), &parsed); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if parsed.PipelineID != 42 || parsed.Stage != "tester" {
		t.Errorf("round-trip mismatch: %+v", parsed)
	}
}

func TestUnsubscribeNonExistentChannel(t *testing.T) {
	bus := NewPipelineEventBus()
	ch := make(chan PipelineEvent, 1)

	// Should not panic
	bus.Unsubscribe(1, ch)
	bus.Unsubscribe(0, ch)
}

func TestMultipleSubscribersSamePipeline(t *testing.T) {
	bus := NewPipelineEventBus()
	chA := bus.Subscribe(1)
	chB := bus.Subscribe(1)
	chC := bus.Subscribe(1)

	bus.Publish(PipelineEvent{PipelineID: 1, Type: "test"})

	<-chA
	<-chB
	<-chC

	bus.Unsubscribe(1, chA)
	bus.Unsubscribe(1, chB)
	bus.Unsubscribe(1, chC)
}
