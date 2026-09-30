package job

import (
	"sync"
	"testing"
)

func TestLifecycle(t *testing.T) {
	store := NewStore()
	j, err := store.Create("work")
	if err != nil {
		t.Fatal(err)
	}
	if j.Status != Pending {
		t.Fatalf("status=%q", j.Status)
	}
	if !store.MarkDone(j.ID) {
		t.Fatal("mark done failed")
	}
	got, ok := store.Get(j.ID)
	if !ok || got.Status != Done {
		t.Fatalf("got %#v ok=%v", got, ok)
	}
}

func TestConcurrentCreate(t *testing.T) {
	store := NewStore()
	var wg sync.WaitGroup
	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := store.Create("work"); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()
}
