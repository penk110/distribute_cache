package stathat_consistent

import "stathat.com/c/consistent"

type Consistent struct {
	*consistent.Consistent
}

func NewConsistent() *Consistent {
	return &Consistent{
		consistent.New(),
	}
}
