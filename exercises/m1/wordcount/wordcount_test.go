package wordcount

import "testing"

func TestCount(t *testing.T) {
	got := Count("Go go learn")
	if got["go"] != 2 || got["learn"] != 1 {
		t.Fatalf("Count returned %#v", got)
	}
}
