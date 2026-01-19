package payment

import "sync"

type Store struct {
	mu sync.Mutex
	m  map[string]bool
}

func NewStore() *Store {
	return &Store{m: make(map[string]bool)}
}

func (s *Store) Used(tx string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.m[tx]
}

func (s *Store) Mark(tx string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[tx] = true
}
