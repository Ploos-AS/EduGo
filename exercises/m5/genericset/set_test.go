package genericset

import "testing"

func TestSet(t *testing.T) {
	s := New("go", "c")
	if !s.Contains("go") || s.Contains("python") {
		t.Fatalf("unexpected set: %#v", s)
	}
}

func TestMap(t *testing.T) {
	got := Map([]int{1, 2, 3}, func(n int) int { return n * n })
	want := []int{1, 4, 9}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}
