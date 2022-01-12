package impl

type Stat interface {
	Add(key string, value []byte)
	Del(key string, value []byte)
}

type Stater struct {
	Count     int64
	KeySize   int64
	ValueSize int64
}

func (s *Stater) Add(k string, v []byte) {
	s.Count += 1
	s.KeySize += int64(len(k))
	s.ValueSize += int64(len(v))
}

func (s *Stater) Del(k string, v []byte) {
	s.Count -= 1
	s.KeySize -= int64(len(k))
	s.ValueSize -= int64(len(v))
}
