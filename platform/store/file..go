package store

type FileStore interface {
	Save() error
	Get() error
}
