package store

import "os"

type FileStore interface {
	Save(os.File) error
	Open(string, *os.File) error
	List() error
}
