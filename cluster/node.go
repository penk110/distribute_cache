package cluster

import (
	"stathat.com/c/consistent"
)

type Node struct {
	*consistent.Consistent // circle
	addr                   string
}

func (n *Node) GetAddr() string {
	return n.addr
}

func (n *Node) Process(name string) (string, bool, error) {
	var (
		addr string
		err  error
	)

	if addr, err = n.Get(name); err != nil {
		return "", false, err
	}

	// current node
	return addr, addr == n.addr, nil
}

func (n *Node) GetMembers() []string {
	return n.Members()
}