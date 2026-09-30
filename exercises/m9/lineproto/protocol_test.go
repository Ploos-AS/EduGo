package lineproto

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestProtocol(t *testing.T) {
	var out bytes.Buffer
	err := Serve(strings.NewReader("PING\nECHO hello\nNOPE\n"), &out, 1024)
	if err != nil {
		t.Fatal(err)
	}
	want := "PONG\nhello\nERR unknown command\n"
	if out.String() != want {
		t.Fatalf("got %q, want %q", out.String(), want)
	}
}

func TestFragmentedInput(t *testing.T) {
	var out bytes.Buffer
	r := &chunkReader{chunks: []string{"PI", "NG\nE", "CHO hi", "\n"}}
	if err := Serve(r, &out, 1024); err != nil {
		t.Fatal(err)
	}
	if out.String() != "PONG\nhi\n" {
		t.Fatalf("got %q", out.String())
	}
}

type chunkReader struct {
	chunks []string
}

func (r *chunkReader) Read(p []byte) (int, error) {
	if len(r.chunks) == 0 {
		return 0, io.EOF
	}
	chunk := r.chunks[0]
	r.chunks = r.chunks[1:]
	return copy(p, chunk), nil
}
