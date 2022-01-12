package network

import (
	"bytes"
	"net"
)

const (
	// PUBLIC PRIVATE 公有/私有网段类型
	PUBLIC = iota
	PRIVATE
)

// var AddrType bool = os.Getenv("ADDR_TYPE") == "PUBLIC"

var (
	// ABC 三类私有网段
	// _AIPRange A类私有网段 2^24
	_AIPRange = IPRange{
		start: net.IPv4(10, 0, 0, 0),
		end:   net.IPv4(10, 255, 255, 255),
	}
	// _BIPRange B类私有网段 2^16
	_BIPRange = IPRange{
		start: net.IPv4(172, 16, 0, 0),
		end:   net.IPv4(172, 31, 255, 255),
	}
	// _CIPRange C类私有网段 2^16
	_CIPRange = IPRange{
		start: net.IPv4(192, 168, 0, 0),
		end:   net.IPv4(192, 168, 255, 255),
	}
	// _ABCIPRange 私有地址汇总
	_ABCIPRange = []IPRange{
		_AIPRange,
		_BIPRange,
		_CIPRange,
	}
)

type IPRange struct {
	start net.IP
	end   net.IP
}

func (ir IPRange) contains(ip net.IP) (flag bool) {
	if bytes.Compare(ip, ir.start) >= 0 && bytes.Compare(ip, ir.end) <= 0 {
		flag = true
	}
	return flag
}

func isPrivateIP(ip net.IP) (flag bool) {
	for _, ips := range _ABCIPRange {
		if ips.contains(ip) {
			return true
		}
	}
	return flag
}

func privateIPs() (ips []string) {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, addr := range addresses {
		ip, _, err := net.ParseCIDR(addr.String())
		if err != nil {
			continue
		}
		// To4 converts the IPv4 address ip to a 4-byte representation.
		// If ip is not an IPv4 address or is a loop back address, To4 returns nil.
		if ip.To4() == nil || ip.IsLoopback() {
			continue
		}
		private := isPrivateIP(ip)
		if private {
			ips = append(ips, ip.String())
		}
	}

	return
}
