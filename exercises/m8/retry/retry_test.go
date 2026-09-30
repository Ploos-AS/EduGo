package retry

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestEventuallySucceeds(t *testing.T) {
	calls := 0
	waits := 0
	err := Do(context.Background(), 3, time.Second, func(context.Context, time.Duration) error {
		waits++
		return nil
	}, func() error {
		calls++
		if calls < 3 {
			return errors.New("temporary")
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if calls != 3 || waits != 2 {
		t.Fatalf("calls=%d waits=%d", calls, waits)
	}
}

func TestCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	called := false
	err := Do(ctx, 3, time.Second, nil, func() error {
		called = true
		return nil
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("got %v", err)
	}
	if called {
		t.Fatal("operation ran after cancellation")
	}
}
