package loader

import (
	"context"
	"database/sql"

	"github.com/blushft/strana/platform/store"
	"github.com/blushft/strana/platform/store/loader/ent"
	entsql "github.com/facebook/ent/dialect/sql"
)

type Store struct {
	client *ent.Client
}

func NewStore(dbs *store.SQLStore) (*Store, error) {
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
