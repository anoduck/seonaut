package crawler

import (
	"sync"
)

type URLStorage struct {
	seen map[string]bool
	lock *sync.RWMutex
}

func NewURLStorage() *URLStorage {
	return &URLStorage{
		seen: make(map[string]bool),
		lock: &sync.RWMutex{},
	}
}
func (s *URLStorage) Seen(u string) bool {
	s.lock.RLock()
	_, ok := s.seen[u]
	s.lock.RUnlock()

	return ok
}

func (s *URLStorage) Add(u string) {
	s.lock.Lock()
	s.seen[u] = true
	s.lock.Unlock()
}
