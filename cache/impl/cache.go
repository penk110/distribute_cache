package impl

type Cache interface {
	Set(key string, value []byte) error
	SetEX(key string, ttl int64, value []byte) error
	Get(key string) ([]byte, error)
	Del(key string) error
	GetStat() Stat
	NewScanner() Scan
}
