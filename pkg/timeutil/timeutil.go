// Package timeutil provides simple time helpers used by audit paths.
package timeutil

import "time"

// UnixMilli returns the given time as Unix milliseconds.
func UnixMilli(t time.Time) int64 { return t.UnixNano() / int64(time.Millisecond) }

// FromUnixMilli reconstructs a UTC time from Unix milliseconds.
func FromUnixMilli(ms int64) time.Time {
	return time.Unix(ms/1000, (ms%1000)*int64(time.Millisecond)).UTC()
}

// FormatRFC3339 formats t in RFC3339 with millisecond precision.
func FormatRFC3339(t time.Time) string { return t.UTC().Format("2006-01-02T15:04:05.000Z07:00") }
