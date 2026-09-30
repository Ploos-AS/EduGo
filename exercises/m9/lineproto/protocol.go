package lineproto

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

var ErrTooLong = errors.New("message too long")

func Serve(r io.Reader, w io.Writer, max int) error {
	if max < 1 {
		return errors.New("max must be positive")
	}

	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 0, min(max, 4096)), max)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case line == "PING":
			fmt.Fprintln(w, "PONG")
		case strings.HasPrefix(line, "ECHO "):
			fmt.Fprintln(w, strings.TrimPrefix(line, "ECHO "))
		default:
			fmt.Fprintln(w, "ERR unknown command")
		}
	}
	if err := scanner.Err(); err != nil {
		if errors.Is(err, bufio.ErrTooLong) {
			return ErrTooLong
		}
		return err
	}
	return nil
}
