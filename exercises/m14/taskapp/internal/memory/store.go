package memory

import (
	"sync"

	"github.com/Ploos-AS/EduGo/exercises/m14/taskapp/task"
)

type Store struct {
	mu     sync.Mutex
	nextID int
}

func New() *Store {
	return &Store{nextID: 1}
}

func (s *Store) Create(text string) task.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := task.Task{ID: s.nextID, Text: text}
	s.nextID++
	return result
}
