package impl

type Scan interface {
	Scan() bool
	Key() string
	Value() []byte
	Close()
}

type Scanner struct {

}

func (s *Scanner) Scan() bool {

	return true
}

func (s *Scanner) Key() string {

	return ""
}

func (s *Scanner) Value() []byte {

	return nil
}

func (s *Scanner) Close() {

}
