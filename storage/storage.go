package storage

import (
	"fmt"
	"sync"
)

type Storage interface{
	Push([]byte) (int, error) 
	Fetch(int) ([]byte, error)
}

type MemoryStore struct{
	mutex sync.RWMutex  //race conditions
	data [][]byte
}

func NewMemoryStore() Storage{
	return &MemoryStore{
		data: make([][]byte, 0),
	}
}

func (s *MemoryStore) Push(b []byte) (int, error){
	s.mutex.Lock()
	
	defer s.mutex.Unlock()

	s.data = append(s.data, b)
	return len(s.data) - 1, nil
}

func (s *MemoryStore) Fetch(offset int) ([]byte, error){
	s.mutex.RLock()
	
	defer s.mutex.RUnlock()

	if len(s.data) < offset{
		return nil, fmt.Errorf("Offset (%d) is too high", offset)
	}
	return s.data[offset], nil
}