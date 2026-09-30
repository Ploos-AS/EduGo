package noteservice

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAndRead(t *testing.T) {
	ts := httptest.NewServer(Handler(NewMemoryStore()))
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/api/v1/notes", "application/json", bytes.NewBufferString(`{"text":"learn Go"}`))
	if err != nil { t.Fatal(err) }
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated { t.Fatalf("create status = %d", resp.StatusCode) }

	var created Note
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil { t.Fatal(err) }

	get, err := http.Get(ts.URL + "/api/v1/notes/1")
	if err != nil { t.Fatal(err) }
	defer get.Body.Close()
	if get.StatusCode != http.StatusOK { t.Fatalf("get status = %d", get.StatusCode) }

	var got Note
	if err := json.NewDecoder(get.Body).Decode(&got); err != nil { t.Fatal(err) }
	if got != created { t.Fatalf("got %#v, want %#v", got, created) }
}
