package task

import (
	"errors"
	"strings"
)

type Task struct {
	ID   int
	Text string
}

type Store interface {
	Create(text string) Task
}

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{store: store}
}

func (s *Service) Create(text string) (Task, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return Task{}, errors.New("text is required")
	}
	return s.store.Create(text), nil
}
