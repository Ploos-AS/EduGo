package health

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthIndependentOfReadiness(t *testing.T) {
	h := New()
	h.SetReady(false)
	server := httptest.NewServer(h.Routes())
	defer server.Close()

	health, err := http.Get(server.URL + "/healthz")
	if err != nil {
		t.Fatal(err)
	}
	defer health.Body.Close()
	if health.StatusCode != http.StatusOK {
		t.Fatalf("health status=%d", health.StatusCode)
	}

	ready, err := http.Get(server.URL + "/readyz")
	if err != nil {
		t.Fatal(err)
	}
	defer ready.Body.Close()
	if ready.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("ready status=%d", ready.StatusCode)
	}
}
