package network

import "net"

// LocalAddr local address
func LocalAddr() (addr string) {
	netInterfaceAddresses, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, netInterfaceAddress := range netInterfaceAddresses {
		ip, ok := netInterfaceAddress.(*net.IPNet)
		if !ok {
			continue
		}
		if ip.IP.To4() == nil || ip.IP.IsLoopback() {
			continue
		}
		addr = ip.IP.String()
		if addr != "<nil>" {
			return addr
		}

	}
	return ""
}
