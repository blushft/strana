package tracker

import (
	"encoding/json"
	"time"

	"github.com/dgraph-io/badger/v2"
)

type StoreEvent struct {
	Event       *Event
	Attempted   bool
	Attempts    int
	LastAttempt time.Time
}

func (se *StoreEvent) JSON() ([]byte, error) {
	return json.Marshal(se)
}

type Store interface {
	Set(*Event) error
	Update(*StoreEvent) error
	Remove(*StoreEvent) error
	Get(string) (*StoreEvent, error)
	GetAll() ([]*StoreEvent, error)
}

type memStore struct {
	db *badger.DB
}

func NewMemStore() (Store, error) {
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		return nil, err
	}

	return &memStore{
		db: db,
	}, nil
}

func (s *memStore) Set(evt *Event) error {
	se := &StoreEvent{
		Event:     evt,
		Attempted: false,
		Attempts:  0,
	}

	return s.Update(se)
}

func (s *memStore) Update(se *StoreEvent) error {
	b, err := se.JSON()
	if err != nil {
		return err
	}

	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(se.Event.ID), b)
	})
}

func (s *memStore) Remove(se *StoreEvent) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(se.Event.ID))
	})
}

func (s *memStore) Get(id string) (*StoreEvent, error) {
	var val []byte
	err := s.db.View(func(txn *badger.Txn) error {
		itm, err := txn.Get([]byte(id))
		if err != nil {
			return err
		}

		if verr := itm.Value(func(v []byte) error {
			copy(val, v)
			return nil
		}); verr != nil {
			return verr
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	var se *StoreEvent
	if err := json.Unmarshal(val, &se); err != nil {
		return nil, err
	}

	return se, nil
}

func (s *memStore) GetAll() ([]*StoreEvent, error) {
	var ses []*StoreEvent
	if err := s.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			var val []byte

			itm := it.Item()
			if verr := itm.Value(func(v []byte) error {
				val = append([]byte{}, v...)
				return nil
			}); verr != nil {
				return verr
			}

			var se *StoreEvent
			if err := json.Unmarshal(val, &se); err != nil {
				return err
			}

			ses = append(ses, se)

			return nil
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return ses, nil
}
