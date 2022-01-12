package cache

import (
	"github.com/penk110/distribute_cache/cache/impl"
	"github.com/penk110/distribute_cache/cache/memory"
	"github.com/penk110/distribute_cache/cache/rocksdb"
)

const (
	_ = iota
	MEMORY
	ROCKSDB
)

func NewCache(ct string, ttl int64) impl.Cache {
	var cache impl.Cache
	switch ct {
	case "memory":
		cache = memory.NewMemoryCache(ttl)
	case "rocksdb":
		cache = rocksdb.NewRocksDB(ttl)
	default:
		cache = memory.NewMemoryCache(ttl)
	}

	return cache
}
