package tracker

import (
	"errors"
	"time"

	"github.com/nedscode/memdb"
)

type StoreEvent struct {
	ID          string
	Event       *Event
	Attempted   bool
	Attempts    int
	LastAttempt time.Time
}

func (se *StoreEvent) IsExpired() bool {
	return se.Attempted && se.Attempts > 3
}

type Store interface {
	Set(*Event) error
	Update(*StoreEvent) error
	Remove(*StoreEvent) error
	Get(string) (*StoreEvent, error)
	GetAll() ([]*StoreEvent, error)
}

type memStore struct {
	mdb *memdb.Store
}

func NewMemStore() (Store, error) {
	mdb := memdb.NewStore().PrimaryKey("id")

	go func() {
		for {
			select {
			case <-time.After(time.Second * 10):
				mdb.Expire()
			}
		}
	}()

	return &memStore{
		mdb: mdb,
	}, nil
}

func (s *memStore) Set(evt *Event) error {
	se := &StoreEvent{
		ID:        evt.ID,
		Event:     evt,
		Attempted: false,
		Attempts:  0,
	}

	_, err := s.mdb.Put(se)

	return err
}

func (s *memStore) Update(se *StoreEvent) error {
	_, err := s.mdb.Put(se)
	return err
}

func (s *memStore) Remove(se *StoreEvent) error {
	_, err := s.mdb.Delete(se)
	return err
}

func (s *memStore) Get(id string) (*StoreEvent, error) {
	obj := s.mdb.Get(&StoreEvent{ID: id})

	se, ok := obj.(*StoreEvent)
	if !ok {
		return nil, errors.New("not found")
	}

	return se, nil
}

func (s *memStore) GetAll() ([]*StoreEvent, error) {
	var ses []*StoreEvent
	s.mdb.In("id").Each(func(itm interface{}) bool {
		se, ok := itm.(*StoreEvent)
		if !ok {
			return false
		}

		ses = append(ses, se)

		return true
	})

	return ses, nil
}
