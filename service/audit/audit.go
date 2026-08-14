// Package audit records handler events with monotonically-ordered timestamps.
package audit

import (
	"sync"
	"time"

	"github.com/xytan0056/bazel-fixture/pkg/logger"
	"github.com/xytan0056/bazel-fixture/pkg/timeutil"
)

type Event struct {
	Actor    string
	Action   string
	AtUnixMs int64
}

type Recorder struct {
	mu     sync.Mutex
	events []Event
	log    *logger.Logger
}

func NewRecorder(log *logger.Logger) *Recorder { return &Recorder{log: log} }

// Record appends a new event with the current time.
func (r *Recorder) Record(actor, action string) Event {
	e := Event{
		Actor:    actor,
		Action:   action,
		AtUnixMs: timeutil.UnixMilli(time.Now()),
	}
	r.mu.Lock()
	r.events = append(r.events, e)
	r.mu.Unlock()
	r.log.Info("audit " + actor + " " + action + " @ " + timeutil.FormatRFC3339(timeutil.FromUnixMilli(e.AtUnixMs)))
	return e
}

// Events returns a snapshot copy of the recorded events.
func (r *Recorder) Events() []Event {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := make([]Event, len(r.events))
	copy(out, r.events)
	return out
}
