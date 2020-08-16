package tracker

import (
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/blushft/strana/pkg/event"
)

type StoreEvent struct {
	ID          string
	Event       []byte
	Attempted   bool
	Attempts    int
	LastAttempt time.Time
}

type Store interface {
	Set(*event.Event) error
	Update(*StoreEvent) error
	Remove(*StoreEvent) error
	Get(string) (*StoreEvent, error)
	GetAll() ([]*StoreEvent, error)
}

type memStore struct {
	m  map[string][]byte
	mu sync.RWMutex
}

func NewMemStore() (Store, error) {
	return &memStore{
		m: make(map[string][]byte),
	}, nil
}

func (s *memStore) Set(evt *event.Event) error {
	pl, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	se := &StoreEvent{
		ID:        evt.ID,
		Event:     pl,
		Attempted: false,
		Attempts:  0,
	}

	return s.Update(se)
}

func (s *memStore) Update(se *StoreEvent) error {
	b, err := json.Marshal(se)
	if err != nil {
		return err
	}

	s.mu.Lock()
	s.m[se.ID] = b
	s.mu.Unlock()

	return nil
}

func (s *memStore) Remove(se *StoreEvent) error {
	s.mu.Lock()
	delete(s.m, se.ID)
	s.mu.Unlock()

	return nil
}

func (s *memStore) Get(id string) (*StoreEvent, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	b, ok := s.m[id]
	if !ok {
		return nil, errors.New("not found")
	}

	var se *StoreEvent
	if err := json.Unmarshal(b, &se); err != nil {
		return nil, err
	}

	return se, nil
}

func (s *memStore) GetAll() ([]*StoreEvent, error) {
	var ses []*StoreEvent

	for k := range s.m {
		se, err := s.Get(k)
		if err != nil {
			return nil, err
		}

		ses = append(ses, se)
	}

	return ses, nil
}
