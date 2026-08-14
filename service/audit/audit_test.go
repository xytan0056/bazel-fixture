package audit

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xytan0056/bazel-fixture/pkg/logger"
)

func TestRecord(t *testing.T) {
	r := NewRecorder(logger.New())
	r.Record("alice", "put")
	r.Record("bob", "echo")
	events := r.Events()
	assert.Len(t, events, 2)
	assert.Equal(t, "alice", events[0].Actor)
	assert.Equal(t, "echo", events[1].Action)
	assert.NotZero(t, events[0].AtUnixMs)
}
