package streamcopy

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestCopy(t *testing.T) {
	var dst bytes.Buffer
	n, err := Copy(&dst, strings.NewReader("hello"), 5)
	if err != nil {
		t.Fatal(err)
	}
	if n != 5 || dst.String() != "hello" {
		t.Fatalf("n=%d output=%q", n, dst.String())
	}
}

func TestLimit(t *testing.T) {
	var dst bytes.Buffer
	n, err := Copy(&dst, strings.NewReader("abcdef"), 5)
	if !errors.Is(err, ErrLimitExceeded) {
		t.Fatalf("got %v", err)
	}
	if n != 6 {
		t.Fatalf("n=%d", n)
	}
}

func TestNegativeLimit(t *testing.T) {
	var dst bytes.Buffer
	if _, err := Copy(&dst, strings.NewReader("x"), -1); err == nil {
		t.Fatal("expected error")
	}
}
