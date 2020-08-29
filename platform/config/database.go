package config

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/lib/pq"
	"github.com/mattn/go-sqlite3"
)

type Database struct {
	Dialect  string `json:"dialect" yaml:"dialect"`
	URL      string `json:"url" yaml:"url"`
	Host     string `json:"host" yaml:"host"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
	SSLMode  string `json:"ssl_mode" yaml:"ssl_mode"`
}

func DefaultDatabaseConfig() Database {
	return Database{
		Dialect:  "sqlite",
		Database: "strana.db",
	}
}

func (dbc Database) ConnString() string {
	var dsn string

	if dbc.URL != "" {
		return dbc.URL
	}

	switch dbc.Dialect {
	case "postgres":
		if dbc.Host != "" {
			dsn += " host=" + dbc.Host
		}
		if dbc.Database != "" {
			dsn += " dbname=" + dbc.Database
		}
		if dbc.User != "" {
			dsn += " user=" + dbc.User
		}
		if dbc.Password != "" {
			dsn += " password=" + dbc.Password
		}
		if dbc.SSLMode != "" {
			dsn += " sslmode=" + dbc.SSLMode
		} else {
			dsn += " sslmode=disable"
		}

		dsn = strings.TrimSpace(dsn)

		return dsn
	case "sqlite":
		dsn = dbc.Database + "?_busy_timeout=10000&cache=shared&_fk=1"

		return dsn
	case "memory":
		dsn = fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", dbc.Database)

		return dsn
	default:
		return ""
	}
}

func (dbc Database) EntDialect() string {
	switch dbc.Dialect {
	case "postgres", "pgx":
		return dialect.Postgres
	case "sqlite", "memory":
		return dialect.SQLite
	default:
		return ""
	}
}

func (dbc Database) Driver() driver.Driver {
	switch dbc.Dialect {
	case "postgres":
		return &pq.Driver{}
	case "sqlite", "memory":
		return &sqlite3.SQLiteDriver{}
	case "pgx":
		return &stdlib.Driver{}
	default:
		return nil
	}
}
