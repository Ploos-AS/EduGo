package streamcopy

import (
	"errors"
	"io"
)

var ErrLimitExceeded = errors.New("input limit exceeded")

func Copy(dst io.Writer, src io.Reader, max int64) (int64, error) {
	if max < 0 {
		return 0, errors.New("max must not be negative")
	}

	limited := io.LimitReader(src, max+1)
	n, err := io.Copy(dst, limited)
	if err != nil {
		return n, err
	}
	if n > max {
		return n, ErrLimitExceeded
	}
	return n, nil
}
