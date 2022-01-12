package rocksdb

// #include <stdlib.h>
// #include "rocksdb/c.h"
// #cgo CFLAGS: -I/usr/local/Cellar/rocksdb/6.22.1/include
// #cgo LDFLAGS: -L/usr/local/Cellar/rocksdb/6.22.1 -lrocksdb -lz -lpthread -lsnappy -lstdc++ -lm -O3
import "C"
import (
	"os"
	"runtime"
	"time"
	"unsafe"

	"github.com/penk110/distribute_cache/cache/impl"
)

const (
	BatchSize = 100
)

// TODO: test code

var TestRocksDB *RocksDB

func init() {
	TestRocksDB = NewRocksDB(30)
}

func NewRocksDB(ttl int64) *RocksDB {
	var (
		rocksDB *RocksDB
		dbPath  string
		e       *C.char
	)
	dbPath = os.Getenv("ROCKSDB_PATH")
	if dbPath == "" {
		panic("invalid rocksdb path")
	}
	options := C.rocksdb_options_create()
	C.rocksdb_options_increase_parallelism(options, C.int(runtime.NumCPU()))
	C.rocksdb_options_set_create_if_missing(options, 1)
	db := C.rocksdb_open_with_ttl(options, C.CString(dbPath), C.int(ttl), &e)
	if e != nil {
		panic(C.GoString(e))
	}
	C.rocksdb_options_destroy(options)
	rocksDB = &RocksDB{
		stat: &impl.Stater{},
		scan: &impl.Scanner{},
		ttl:  ttl,
		ch:   make(chan *impl.Pair, 2000),
		db:   db,
		wo:   C.rocksdb_writeoptions_create(),
		ro:   C.rocksdb_readoptions_create(),
		e:    e,
	}

	go writer(db, rocksDB.ch, rocksDB.wo)

	return rocksDB
}

func (rdb *RocksDB) NewScanner() impl.Scan {

	return rdb.scan
}

func flushBatch(db *C.rocksdb_t, b *C.rocksdb_writebatch_t, o *C.rocksdb_writeoptions_t) {
	var e *C.char
	C.rocksdb_write(db, o, b, &e)
	if e != nil {
		panic(C.GoString(e))
	}
	C.rocksdb_writebatch_clear(b)
}

func writer(db *C.rocksdb_t, c chan *impl.Pair, o *C.rocksdb_writeoptions_t) {
	count := 0
	t := time.NewTimer(time.Millisecond * 600)
	b := C.rocksdb_writebatch_create()
	// TODO: 批量或者定时写入 会导致 立即写后立即读取到的是空值
	for {
		select {
		case p := <-c:
			count++
			key := C.CString(p.K)
			value := C.CBytes(p.V)
			C.rocksdb_writebatch_put(b, key, C.size_t(len(p.K)), (*C.char)(value), C.size_t(len(p.V)))
			C.free(unsafe.Pointer(key))
			C.free(value)
			if count == BatchSize {
				flushBatch(db, b, o)
				count = 0
			}
			if !t.Stop() {
				<-t.C
			}
			t.Reset(time.Millisecond * 600)
		case <-t.C:
			if count != 0 {
				flushBatch(db, b, o)
				count = 0
			}
			t.Reset(time.Millisecond * 600)
		}
	}
}
