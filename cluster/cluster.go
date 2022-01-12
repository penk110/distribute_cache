package cluster

import (
	"github.com/hashicorp/memberlist"
	"io/ioutil"
	"log"
	"stathat.com/c/consistent"
	"time"
)

type Cluster interface {
	Process(key string) (string, bool, error)
	GetMembers() []string
	GetAddr() string
}

// TODO: create node and join to clusters

func NewNode(addr string, port int, cluster string) (*Node, error) {
	var (
		node     *Node
		conf     *memberlist.Config
		members  *memberlist.Memberlist
		circle   *consistent.Consistent
		clusters []string
		err      error
	)
	node = new(Node)
	conf = memberlist.DefaultLANConfig()
	conf.Name = addr
	conf.BindAddr = addr
	conf.BindPort = port
	conf.LogOutput = ioutil.Discard

	if members, err = memberlist.Create(conf); err != nil {
		return nil, err
	}
	if cluster == "" {
		cluster = addr
	}
	clusters = []string{cluster}

	_, err = members.Join(clusters)
	if err != nil {
		log.Printf("join to clusters failed, err: %v\n", err.Error())
		return nil, err
	}

	// To change the number of replicas, set NumberOfReplicas before adding entries.
	circle = consistent.New()
	// TODO: get number of replicas from env or config
	circle.NumberOfReplicas = 256

	// loop
	go func() {
		var ticker *time.Ticker
		ticker = time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				m := members.Members()
				nodes := make([]string, len(m))
				for i, n := range m {
					nodes[i] = n.Name
				}
				circle.Set(nodes)
			}
		}
	}()

	node.Consistent = circle
	node.addr = addr

	return node, nil
}
