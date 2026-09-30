package miniservice

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status,omitempty"`
	Text   string `json:"text,omitempty"`
}

func Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, Response{Status: "ok"})
	})
	mux.HandleFunc("GET /api/v1/echo", func(w http.ResponseWriter, r *http.Request) {
		text := r.URL.Query().Get("text")
		if text == "" {
			writeJSON(w, http.StatusBadRequest, Response{Status: "missing text"})
			return
		}
		writeJSON(w, http.StatusOK, Response{Text: text})
	})
	return mux
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
