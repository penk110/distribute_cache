package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/penk110/distribute_cache/cache/rocksdb"
	"log"
	"net/http"
)

var CacheRouter *CacheRoute

type CacheRoute struct {
}

func GetCacheRouter() *CacheRoute {
	return CacheRouter
}

const (
	PrefixKey = "/distribute_cache/dev/"
)

func (ch *CacheRoute) Retrieve(ctx *gin.Context) {
	var (
		value []byte
		key   string
		err   error
	)
	key = ctx.Param("key")
	if key == "" {
		ctx.JSON(http.StatusOK, gin.H{"msg": err.Error(), "result": "InvalidKey"})
	}
	// key: /distribute_cache/dev/ + key
	if value, err = rocksdb.TestRocksDB.Get(PrefixKey + key); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"msg": err.Error(), "result": "Get Failed"})
		return
	}
	log.Println("value: " + string(value))

	ctx.JSON(http.StatusOK, gin.H{"msg": "success", "result": "value: " + string(value)})
}

/*
	{
		"key": "/distribute_cache/dev/001",
		"value": "2021-10-14 02:33:17"
	}
*/

type CData struct {
	Key   string          `json:"key" binding:"required,max=64"`
	Value json.RawMessage `json:"value" binding:"required"`
}

func (ch *CacheRoute) Create(ctx *gin.Context) {
	var (
		opts *CData
		err  error
	)
	opts = new(CData)
	if err = ctx.BindJSON(&opts); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"msg": err.Error(), "result": "SetFailed"})
		return
	}

	if err = rocksdb.TestRocksDB.Set(PrefixKey+opts.Key, opts.Value); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"msg": err.Error(), "result": "Set Failed"})
		return
	}
	log.Println("value: " + string(opts.Value))
	ctx.JSON(http.StatusOK, gin.H{"msg": "success", "result": "value: " + string(opts.Value)})
	return
}
