package timeutil

import "time"

func StartOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func DaysBetween(a, b time.Time) int {
	a, b = StartOfDay(a), StartOfDay(b)
	return int(b.Sub(a).Hours() / 24)
}
