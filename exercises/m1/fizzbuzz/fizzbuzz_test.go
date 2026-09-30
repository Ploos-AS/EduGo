package fizzbuzz

import "testing"

func TestLabel(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{1, "1"}, {3, "Fizz"}, {5, "Buzz"}, {15, "FizzBuzz"}, {16, "16"},
	}
	for _, tt := range tests {
		if got := Label(tt.n); got != tt.want {
			t.Errorf("Label(%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}
