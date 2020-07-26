package store

import (
	"context"

	"github.com/blushft/strana/platform/config"
	"github.com/blushft/strana/platform/store/ent"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	client *ent.Client
}

func NewStore(conf config.Database) (*Store, error) {
	driver, conn := conf.ConnString()

	c, err := ent.Open(driver, conn)
	if err != nil {
		return nil, err
	}

	s := &Store{
		client: c,
	}

	if err := s.setup(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Store) Client() *ent.Client {
	return s.client
}

func (s *Store) Mount(fn func(*Store)) {
	fn(s)
}

func (s *Store) setup() error {
	if err := s.client.Schema.Create(context.TODO()); err != nil {
		return err
	}

	scnt, err := s.client.App.Query().Count(context.TODO())
	if err != nil {
		return err
	}

	if scnt == 0 {
		_, err := s.client.App.Create().SetName("Default").SetTrackingID("default").Save(context.TODO())
		if err != nil {
			return err
		}
	}

	return nil
}
