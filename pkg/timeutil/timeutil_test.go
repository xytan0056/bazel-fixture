package timeutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUnixMilliRoundTrip(t *testing.T) {
	orig := time.Date(2026, 8, 13, 12, 34, 56, 789_000_000, time.UTC)
	ms := UnixMilli(orig)
	assert.Equal(t, orig, FromUnixMilli(ms))
}

func TestFormatRFC3339(t *testing.T) {
	tm := time.Date(2026, 1, 2, 3, 4, 5, 6_000_000, time.UTC)
	assert.Equal(t, "2026-01-02T03:04:05.006Z", FormatRFC3339(tm))
}
