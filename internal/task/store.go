package task

import (
	"fmt"
	"sync"
)

type Task struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type Store interface {
	List() []Task
	Create(t Task) Task
	Update(id string, t Task) (Task, bool)
	Delete(id string) bool
}

type InMemoryStore struct {
	mu    sync.Mutex
	tasks map[string]Task
	next  int
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{tasks: make(map[string]Task)}
}

func (s *InMemoryStore) List() []Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		out = append(out, t)
	}
	return out
}

func (s *InMemoryStore) Create(t Task) Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.next++
	t.ID = fmt.Sprintf("%d", s.next)
	s.tasks[t.ID] = t
	return t
}

func (s *InMemoryStore) Update(id string, t Task) (Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.tasks[id]
	if !ok {
		return Task{}, false
	}
	t.ID = id
	s.tasks[id] = t
	return t, true
}

func (s *InMemoryStore) Delete(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.tasks[id]; !ok {
		return false
	}
	delete(s.tasks, id)
	return true
}
