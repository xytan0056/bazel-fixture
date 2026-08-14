package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/xytan0056/bazel-fixture/pkg/logger"
	"github.com/xytan0056/bazel-fixture/service/audit"
	"github.com/xytan0056/bazel-fixture/service/config"
	"github.com/xytan0056/bazel-fixture/service/store"
)

func TestEcho(t *testing.T) {
	cfg := config.Config{
		Features:       config.Features{Echo: true, Store: true},
		MaxConnections: 100,
	}
	log := logger.New()
	h := New(cfg, store.New(log), audit.NewRecorder(log), log)
	out, err := h.Echo("alice", "hello & world")
	require.NoError(t, err)
	assert.Equal(t, "Hello And World", out)
}
