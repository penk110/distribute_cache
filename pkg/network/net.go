package network

import (
	"fmt"
	"log"
	"net"
)

func GetLocalTCPListener(port int) (net.Listener, error) {
	addresses := privateIPs()
	if len(addresses) == 0 {
		log.Fatal("[GetLocalTCPListener] get private IP address failed.")
		return nil, nil
	}
	// TODO: 默认取第一个地址作为监听地址
	addr1 := fmt.Sprintf("%s:%d", addresses[0], port)

	return net.Listen("tcp", addr1)
}

// GetLocalListenerWithNET get local listener with net(TCP or unix)
func GetLocalListenerWithNET(network string, port int) (net.Listener, error) {
	addresses := privateIPs()
	if len(addresses) == 0 {
		log.Fatal("[GetLocalListenerWithNET] get private IP address failed.")
		return nil, nil
	}
	addr1 := addresses[0]
	if port > 0 {
		addr1 = fmt.Sprintf("%s:%d", addr1, port)
	}

	return net.Listen(network, addr1)
}
