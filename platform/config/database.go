package config

import (
	"fmt"
	"strings"
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
		Database: "strana",
	}
}

func (dbc Database) ConnString() (string, string) {
	var dsn string

	if dbc.URL != "" {
		return dbc.Dialect, dbc.URL
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

		return "postgres", dsn
	case "sqlite":
		dsn = dbc.Database + "?_busy_timeout=10000&cache=shared&_fk=1"

		return "sqlite3", dsn
	case "memory":
		dsn = fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", dbc.Database)
		return "sqlite3", dsn
	default:
		return "", ""
	}
}
