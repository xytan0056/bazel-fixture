package mathutil

import "testing"

func TestClamp(t *testing.T) {
	if Clamp(5, 0, 3) != 3 {
		t.Fatal("clamp high failed")
	}
	if Clamp(-1, 0, 3) != 0 {
		t.Fatal("clamp low failed")
	}
	if Clamp(2, 0, 3) != 2 {
		t.Fatal("clamp mid failed")
	}
}

func TestSum(t *testing.T) {
	if Sum([]int{1, 2, 3}) != 6 {
		t.Fatal("sum wrong")
	}
}

func TestMax(t *testing.T) {
	if _, ok := Max(nil); ok {
		t.Fatal("expected empty")
	}
	m, ok := Max([]int{3, 7, 1})
	if !ok || m != 7 {
		t.Fatal("max wrong")
	}
}
