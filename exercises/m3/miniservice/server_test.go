package miniservice

import (
	"testing"
	"time"
)

func TestServerTimeouts(t *testing.T) {
	srv := Server(":0")
	if srv.ReadHeaderTimeout <= 0 || srv.ReadTimeout <= 0 || srv.WriteTimeout <= 0 || srv.IdleTimeout <= 0 {
		t.Fatal("server must define positive timeouts")
	}
	if srv.ReadHeaderTimeout > 30*time.Second {
		t.Fatal("read header timeout is unexpectedly large")
	}
}
