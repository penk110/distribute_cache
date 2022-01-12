package memory

import (
	"errors"
	"sync"
	"time"

	"github.com/penk110/distribute_cache/cache/impl"
)

type memoryValue struct {
	value     []byte
	ac        int64 // TODO: ac ?
	ttl       int64
	createdAt int64
}

type Cache struct {
	values map[string]memoryValue
	mutex  sync.RWMutex
	stat   impl.Stat
	scan   impl.Scan
	ttl    int64
}

func NewMemoryCache(ttl int64) *Cache {
	var mc *Cache
	mc = &Cache{
		values: make(map[string]memoryValue),
		mutex:  sync.RWMutex{},
		stat:   &impl.Stater{},
		scan:   &impl.Scanner{},
		ttl:    ttl,
	}

	go mc.expireHandler()
	return mc
}

func (mc *Cache) Set(key string, value []byte) error {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	mc.values[key] = memoryValue{
		value:     value,
		createdAt: time.Now().Unix(),
	}
	mc.stat.Add(key, value)
	return nil
}

// TODO: SetEx 是否必须

func (mc *Cache) SetEX(key string, ttl int64, value []byte) error {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	mc.values[key] = memoryValue{
		value:     value,
		createdAt: time.Now().Unix(),
		ttl:       ttl,
	}
	mc.stat.Add(key, value)
	return nil
}

func (mc *Cache) Get(key string) ([]byte, error) {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	v, ok := mc.values[key]
	if !ok {
		return nil, errors.New("key<" + key + "> don't exists")
	}
	return v.value, nil
}

func (mc *Cache) Del(key string) error {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	_, ok := mc.values[key]
	if !ok {
		return errors.New("key<" + key + "> don't exists")
	}
	delete(mc.values, key)

	return nil
}

func (mc *Cache) GetStat() impl.Stat {
	return mc.stat
}

func (mc *Cache) NewScanner() impl.Scan {
	return mc.scan
}

func (mc *Cache) expireHandler() {
	var (
		ticker *time.Ticker
	)

	ticker = time.NewTicker(time.Second * 3)
	for true {
		select {
		case <-ticker.C:
			// TODO: 过期清理
		}
	}
}
