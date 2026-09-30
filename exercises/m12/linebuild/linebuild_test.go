package linebuild

import "testing"

var parts = []string{"alpha", "-", "beta", "-", "gamma", "-", "delta"}

func TestSameResult(t *testing.T) {
	if Plus(parts) != Builder(parts) {
		t.Fatal("implementations differ")
	}
}

func BenchmarkPlus(b *testing.B) {
	for b.Loop() {
		_ = Plus(parts)
	}
}

func BenchmarkBuilder(b *testing.B) {
	for b.Loop() {
		_ = Builder(parts)
	}
}
