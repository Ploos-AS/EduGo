package workerpool

import (
	"context"
	"testing"
)

func TestRun(t *testing.T) {
	jobs := []Job{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	got := Run(context.Background(), 3, jobs)

	byID := make(map[int]int)
	for _, result := range got {
		byID[result.ID] = result.Square
	}

	want := map[int]int{1: 4, 2: 9, 3: 16, 4: 25}
	for id, square := range want {
		if byID[id] != square {
			t.Fatalf("job %d = %d, want %d", id, byID[id], square)
		}
	}
}

func TestRunCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if got := Run(ctx, 2, []Job{{1, 2}}); len(got) != 0 {
		t.Fatalf("cancelled Run returned %v", got)
	}
}
