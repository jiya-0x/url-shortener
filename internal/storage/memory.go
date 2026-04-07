package storage

import (
	"fmt"
	"sync"
)

type MemoryStore struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]string),
	}
}

func (s *MemoryStore) Save(shortCode, originalURL string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[shortCode]; exists {
		return fmt.Errorf("short code '%s' already exists", shortCode)
	}
	s.data[shortCode] = originalURL
	return nil
}

func (s *MemoryStore) Get(shortCode string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	url, exists := s.data[shortCode]
	return url, exists
}
