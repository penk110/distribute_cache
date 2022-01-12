package rocksdb

import (
	"strconv"
	"testing"
	"time"
)

func TestRocksDB(t *testing.T) {
	var (
		db        *Rocksdb
		key       string
		valueByte []byte
		valueStr  string
		err       error
	)
	key = "/distrubte_cache/dev/" + strconv.FormatInt(time.Now().Unix(), 10)
	db = NewRocksdb(30)
	err = db.Set(key, []byte(strconv.FormatInt(time.Now().Unix(), 10)))
	if err != nil {
		t.Error(err.Error())
		return
	}

	// 确保已经刷新到硬盘
	time.Sleep(time.Second * 2)

	valueByte, err = db.Get(key)
	if err != nil {
		t.Error(err.Error())
		return
	}
	valueStr = string(valueByte)
	t.Logf("valueStr: %s\n", valueStr)

}
