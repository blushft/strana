package store

import (
	"context"
	"database/sql/driver"

	"contrib.go.opencensus.io/integrations/ocsql"
	"github.com/blushft/strana/platform/config"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type SQLStore struct {
	conf config.Database
}

func NewSQL(conf config.Database) (*SQLStore, error) {
	return &SQLStore{
		conf: conf,
	}, nil
}

func (s *SQLStore) Dialect() string {
	return s.conf.EntDialect()
}

func (s *SQLStore) Connect(context.Context) (driver.Conn, error) {
	return s.Driver().Open(s.conf.ConnString())
}

func (s *SQLStore) Driver() driver.Driver {
	return ocsql.Wrap(
		s.conf.Driver(),
		ocsql.WithAllTraceOptions(),
		ocsql.WithRowsClose(false),
		ocsql.WithRowsNext(false),
		ocsql.WithDisableErrSkip(true),
	)
}

func (s *SQLStore) Mount(fn func(*SQLStore) error) error {
	return fn(s)
}
