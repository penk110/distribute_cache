package rocksdb

// #include <stdlib.h>
// #include "rocksdb/c.h"
// #cgo CFLAGS: -I/usr/local/Cellar/rocksdb/6.22.1/include
// #cgo LDFLAGS: -L/usr/local/Cellar/rocksdb/6.22.1 -lrocksdb -lz -lpthread -lsnappy -lstdc++ -lm -O3
import "C"
import (
	"errors"
	"runtime"
	"time"
	"unsafe"
)

const BatchSize = 100

type Rocksdb struct {
	db *C.rocksdb_t
	ro *C.rocksdb_readoptions_t
	wo *C.rocksdb_writeoptions_t
	e  *C.char
	ch chan *pair
}

type pair struct {
	k string
	v []byte
}

func flushBatch(db *C.rocksdb_t, b *C.rocksdb_writebatch_t, o *C.rocksdb_writeoptions_t) {
	var e *C.char
	C.rocksdb_write(db, o, b, &e)
	if e != nil {
		panic(C.GoString(e))
	}
	C.rocksdb_writebatch_clear(b)
}

func writer(db *C.rocksdb_t, c chan *pair, o *C.rocksdb_writeoptions_t) {
	count := 0

	// TODO: 每秒保存
	t := time.NewTimer(time.Millisecond * 600)
	b := C.rocksdb_writebatch_create()
	for {
		select {
		case p := <-c:
			count++
			key := C.CString(p.k)
			value := C.CBytes(p.v)
			C.rocksdb_writebatch_put(b, key, C.size_t(len(p.k)), (*C.char)(value), C.size_t(len(p.v)))
			C.free(unsafe.Pointer(key))
			C.free(value)
			if count == BatchSize {
				flushBatch(db, b, o)
				count = 0
			}
			if !t.Stop() {
				<-t.C
			}
			t.Reset(time.Second)
		case <-t.C:
			if count != 0 {
				flushBatch(db, b, o)
				count = 0
			}
			t.Reset(time.Second)
		}
	}
}

func (c *Rocksdb) Set(key string, value []byte) error {
	c.ch <- &pair{key, value}
	return nil
}

func (c *Rocksdb) Get(key string) ([]byte, error) {
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))
	var length C.size_t
	v := C.rocksdb_get(c.db, c.ro, k, C.size_t(len(key)), &length, &c.e)
	if c.e != nil {
		return nil, errors.New(C.GoString(c.e))
	}
	defer C.free(unsafe.Pointer(v))
	return C.GoBytes(unsafe.Pointer(v), C.int(length)), nil
}

func NewRocksdb(ttl int) *Rocksdb {
	options := C.rocksdb_options_create()
	C.rocksdb_options_increase_parallelism(options, C.int(runtime.NumCPU()))
	C.rocksdb_options_set_create_if_missing(options, 1)
	var e *C.char

	// TODO: get db path from env or config
	// creat db with ttl
	db := C.rocksdb_open_with_ttl(options, C.CString("/Users/zhang/gopath/src/github.com/penk110/distribute_cache/test/rocksdb"), C.int(ttl), &e)
	if e != nil {
		panic(C.GoString(e))
	}
	C.rocksdb_options_destroy(options)
	c := make(chan *pair, 5000)
	wo := C.rocksdb_writeoptions_create()

	go writer(db, c, wo)

	return &Rocksdb{db, C.rocksdb_readoptions_create(), wo, e, c}
}
