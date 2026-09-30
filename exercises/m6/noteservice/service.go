package noteservice

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Note struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type Store interface {
	Create(text string) Note
	Get(id int) (Note, bool)
}

type MemoryStore struct {
	mu     sync.RWMutex
	nextID int
	notes  map[int]Note
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{nextID: 1, notes: make(map[int]Note)}
}

func (s *MemoryStore) Create(text string) Note {
	s.mu.Lock()
	defer s.mu.Unlock()

	n := Note{ID: s.nextID, Text: text}
	s.nextID++
	s.notes[n.ID] = n
	return n
}

func (s *MemoryStore) Get(id int) (Note, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	n, ok := s.notes[id]
	return n, ok
}

func Handler(store Store) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/notes", func(w http.ResponseWriter, r *http.Request) {
		var in struct {
			Text string `json:"text"`
		}
		if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20)).Decode(&in); err != nil || strings.TrimSpace(in.Text) == "" {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}
		writeJSON(w, http.StatusCreated, store.Create(in.Text))
	})
	mux.HandleFunc("GET /api/v1/notes/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		n, ok := store.Get(id)
		if !ok {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		writeJSON(w, http.StatusOK, n)
	})
	return mux
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
