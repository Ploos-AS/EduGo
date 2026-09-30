package retry

import (
	"context"
	"errors"
	"time"
)

type WaitFunc func(context.Context, time.Duration) error

func Wait(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func Do(ctx context.Context, attempts int, delay time.Duration, wait WaitFunc, fn func() error) error {
	if attempts < 1 {
		return errors.New("attempts must be positive")
	}
	if wait == nil {
		wait = Wait
	}

	var last error
	for attempt := 0; attempt < attempts; attempt++ {
		if err := ctx.Err(); err != nil {
			return err
		}
		if err := fn(); err == nil {
			return nil
		} else {
			last = err
		}

		if attempt+1 < attempts {
			if err := wait(ctx, delay); err != nil {
				return err
			}
		}
	}
	return last
}
