package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/xytan0056/bazel-fixture/pkg/logger"
	"github.com/xytan0056/bazel-fixture/service/config"
	"github.com/xytan0056/bazel-fixture/service/store"
)

func TestEcho(t *testing.T) {
	cfg := config.Config{
		Features:       config.Features{Echo: true, Store: true},
		MaxConnections: 100,
	}
	h := New(cfg, store.New(logger.New()), logger.New())
	out, err := h.Echo("hello & world")
	require.NoError(t, err)
	assert.Equal(t, "Hello And World", out)
}
