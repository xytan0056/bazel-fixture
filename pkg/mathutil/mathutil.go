// Package mathutil provides small numeric helpers.
package mathutil

func Clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func Sum(xs []int) int {
	total := 0
	for _, x := range xs {
		total += x
	}
	return total
}

func Max(xs []int) (int, bool) {
	if len(xs) == 0 {
		return 0, false
	}
	m := xs[0]
	for _, x := range xs[1:] {
		if x > m {
			m = x
		}
	}
	return m, true
}
