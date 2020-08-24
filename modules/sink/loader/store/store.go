package store

import (
	"context"
	"database/sql"

	"github.com/blushft/strana/modules/sink/loader/store/ent"
	"github.com/blushft/strana/platform/store"
	entsql "github.com/facebook/ent/dialect/sql"
)

type Store struct {
	client *ent.Client
}

func New(dbs *store.SQLStore) (*Store, error) {
	db := sql.OpenDB(dbs)
	drv := entsql.OpenDB(dbs.Dialect(), db)

	s := &Store{
		client: ent.NewClient(ent.Driver(drv)),
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
	return s.client.Schema.Create(context.TODO())
}
