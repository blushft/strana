package minio

import (
	"os"

	"github.com/minio/minio-go/v6"
)

type Options struct {
	Bucket   string `json:"bucket" yaml:"bucket"`
	Location string `json:"location" yaml:"location"`
	Prefix   string `json:"prefix" yaml:"prefix"`
	URL      string `json:"url" yaml:"url"`
	Access   string `json:"access" yaml:"access"`
	Secret   string `json:"secret" yaml:"secret"`
	Secure   bool   `json:"secure" yaml:"secure"`
}

type Option func(*Options)

type minioStore struct {
	opts Options
}

func (s *minioStore) client() (*minio.Client, error) {
	return minio.New(s.opts.URL, s.opts.Access, s.opts.Secret, s.opts.Secure)
}

func (s *minioStore) Save(f *os.File) error {
	mc, err := s.client()
	if err != nil {
		return err
	}

	if err := mc.MakeBucket(s.opts.Bucket, s.opts.Location); err != nil {
		ex, xerr := mc.BucketExists(s.opts.Bucket)
		if xerr != nil {
			return xerr
		}

		if !ex {
			return err
		}
	}

	_, err = mc.FPutObject(s.opts.Bucket, f.Name(), f.Name(), minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (s *minioStore) Open(p string, f *os.File) error {
	mc, err := s.client()
	if err != nil {
		return err
	}

	return mc.FGetObject(s.opts.Bucket, p, f.Name(), minio.GetObjectOptions{})
}
