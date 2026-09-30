package authtoken

import (
	"errors"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	key := []byte("0123456789abcdef0123456789abcdef")
	token, err := New(key)
	if err != nil {
		t.Fatal(err)
	}
	if err := Verify(key, token); err != nil {
		t.Fatal(err)
	}
}

func TestTamperedToken(t *testing.T) {
	key := []byte("0123456789abcdef0123456789abcdef")
	token, err := New(key)
	if err != nil {
		t.Fatal(err)
	}

	last := token[len(token)-1]
	replacement := byte('A')
	if last == replacement {
		replacement = 'B'
	}
	tampered := token[:len(token)-1] + string(replacement)

	if err := Verify(key, tampered); !errors.Is(err, ErrInvalid) {
		t.Fatalf("got %v", err)
	}
}

func TestRejectsShortKey(t *testing.T) {
	if _, err := New([]byte("short")); err == nil {
		t.Fatal("expected error")
	}
}
