package netintf

import (
	"net"

	"github.com/pkg/errors"
)

func PrivateCIDRs() []*net.IPNet {
	pbs := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"100.64.0.0/10",
		"fd00::/8",
	}

	res := make([]*net.IPNet, 0, len(pbs))
	for _, b := range pbs {
		if _, block, err := net.ParseCIDR(b); err == nil {
			res = append(res, block)
		}
	}

	return res
}

func IsPrivate(addr net.IP) bool {
	for _, priv := range PrivateCIDRs() {
		if priv.Contains(addr) {
			return true
		}
	}

	return false
}

func AvailableAddresses() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, errors.Wrap(err, "failed to list interfaces")
	}

	var addrs []string
	for _, iface := range ifaces {
		iaddrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, ia := range iaddrs {
			var ip net.IP

			switch v := ia.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			default:
				continue
			}

			if ip == nil {
				continue
			}

			addrs = append(addrs, ip.String())
		}
	}

	return addrs, nil
}
