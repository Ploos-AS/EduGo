package job

import (
	"errors"
	"strings"
	"sync"
)

type Status string

const (
	Pending Status = "pending"
	Done    Status = "done"
)

type Job struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Status Status `json:"status"`
}

type Store struct {
	mu     sync.RWMutex
	nextID int
	jobs   map[int]Job
}

func NewStore() *Store {
	return &Store{nextID: 1, jobs: make(map[int]Job)}
}

func (s *Store) Create(text string) (Job, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return Job{}, errors.New("text is required")
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	j := Job{ID: s.nextID, Text: text, Status: Pending}
	s.nextID++
	s.jobs[j.ID] = j
	return j, nil
}

func (s *Store) Get(id int) (Job, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	j, ok := s.jobs[id]
	return j, ok
}

func (s *Store) MarkDone(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	j, ok := s.jobs[id]
	if !ok {
		return false
	}
	j.Status = Done
	s.jobs[id] = j
	return true
}
