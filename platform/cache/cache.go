package cache

import (
	"time"

	"github.com/allegro/bigcache"
	"github.com/blushft/strana/platform/config"
	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
)

type Cache struct {
	*cache.Cache
}

func NewCache(conf config.Cache) (*Cache, error) {
	if conf.DefaultExpiration == 0 {
		conf.DefaultExpiration = 15
	}

	de := time.Duration(conf.DefaultExpiration)
	bcc, err := bigcache.NewBigCache(bigcache.DefaultConfig(de * time.Minute))
	if err != nil {
		return nil, err
	}

	bsc := store.NewBigcache(bcc, nil)

	return &Cache{
		Cache: cache.New(bsc),
	}, nil
}
