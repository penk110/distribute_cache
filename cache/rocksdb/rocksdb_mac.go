// Package rocksdb +build !linux
package rocksdb

// #include <stdlib.h>
// #include "rocksdb/c.h"
// #cgo CFLAGS: -I/usr/local/Cellar/rocksdb/6.22.1/include
// #cgo LDFLAGS: -L/usr/local/Cellar/rocksdb/6.22.1 -lrocksdb -lz -lpthread -lsnappy -lstdc++ -lm -O3
import "C"
import (
	"errors"
	"log"
	"unsafe"

	"github.com/penk110/distribute_cache/cache/impl"
)

type RocksDB struct {
	stat impl.Stat
	scan impl.Scan
	ttl  int64
	db   *C.rocksdb_t // TODO: C API
	ro   *C.rocksdb_readoptions_t
	wo   *C.rocksdb_writeoptions_t
	e    *C.char
	ch   chan *impl.Pair
}

func (rdb *RocksDB) Set(key string, value []byte) error {
	rdb.ch <- &impl.Pair{K: key, V: value}
	return nil
}

func (rdb *RocksDB) SetEX(key string, ttl int64, value []byte) error {

	return nil
}

func (rdb *RocksDB) GetStat() impl.Stat {

	return rdb.stat
}

func (rdb *RocksDB) Get(key string) ([]byte, error) {
	defer func() {
		e := recover()
		if e != nil {
			log.Printf("e: %v\n", e)
		}
	}()
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))
	var length C.size_t
	v := C.rocksdb_get(rdb.db, rdb.ro, k, C.size_t(len(key)), &length, &rdb.e)
	if rdb.e != nil {
		return nil, errors.New(C.GoString(rdb.e))
	}
	defer C.free(unsafe.Pointer(v))
	return C.GoBytes(unsafe.Pointer(v), C.int(length)), nil
}

func (rdb *RocksDB) Del(key string) error {

	return nil
}
