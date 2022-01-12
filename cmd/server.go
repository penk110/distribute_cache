package main

import (
	"flag"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/penk110/distribute_cache/cache"
	"github.com/penk110/distribute_cache/cache/impl"
	"github.com/penk110/distribute_cache/cluster"
	"github.com/penk110/distribute_cache/pkg/network"
	"github.com/penk110/distribute_cache/server/http"
)

var (
	st   string
	port int
	ttl  int64
	node string
	ca   string
)

func init() {
	flag.StringVar(&st, "st", "memory", "service type")
	flag.IntVar(&port, "port", 8080, "listen port")
	flag.Int64Var(&ttl, "ttl", 60, "ttl")
	flag.StringVar(&ca, "ca", "", "cluster address")
	flag.StringVar(&node, "node", "127.0.0.1", "node address")

	log.Printf("ct: %s, port: %d, ttl: %d, ca: %s\n", st, port, ttl, ca)
}

// go run service.go -st memory -port 8080 -ttl 30 -ca 127.0.0.1
// main entry
func main() {
	var (
		c        impl.Cache
		n        cluster.Cluster
		listener net.Listener
		addr     string
		err      error
	)
	if listener, err = network.GetLocalTCPListener(port); err != nil {
		panic(err)
	}

	c = cache.NewCache(st, ttl)
	addr = listener.Addr().String()
	log.Printf("listen address: %s\n", addr)
	port, _ = strconv.Atoi(strings.Split(listener.Addr().String(), ":")[1])
	if n, err = cluster.NewNode(node, port, ca); err != nil {
		panic(err)
	}

	go runTCP(listener)
	http.NewServer(c, n).Listen(listener)
}

func runTCP(listener net.Listener) {
}
