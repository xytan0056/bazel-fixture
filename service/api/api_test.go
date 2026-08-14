package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/xytan0056/bazel-fixture/pkg/logger"
	"github.com/xytan0056/bazel-fixture/service/config"
	"github.com/xytan0056/bazel-fixture/service/handlers"
	"github.com/xytan0056/bazel-fixture/service/store"
)

func TestServeEcho(t *testing.T) {
	cfg := config.Config{
		Features:       config.Features{Echo: true, Store: true},
		MaxConnections: 100,
	}
	log := logger.New()
	a := New(handlers.New(cfg, store.New(log), log), log)
	out, err := a.Serve("echo", "hello world")
	require.NoError(t, err)
	assert.Equal(t, "Hello World", out)
}
