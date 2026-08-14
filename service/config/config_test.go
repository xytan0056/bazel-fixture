package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	c, err := Load()
	require.NoError(t, err)
	assert.Equal(t, "fixture", c.ServiceName)
	assert.Equal(t, "0.0.0.0:8080", c.ListenAddress)
	assert.True(t, c.Features.Echo)
}
