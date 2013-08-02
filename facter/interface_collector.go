package facter

import (
	"regexp"
	"strings"
)

var (
	re_interface  = regexp.MustCompile(`(^\S+)`)
	re_ipaddress  = regexp.MustCompile(`inet addr:([0-9]+\.[0-9]+\.[0-9]+\.[0-9]+)`)
	re_ipaddress6 = regexp.MustCompile(`inet6 addr: ((?:[0-9,a-f,A-F]*\:{1,2})+[0-9,a-f,A-F]{0,4})/`)
	re_macaddress = regexp.MustCompile(`(?:ether|HWaddr)\s+(\w{1,2}:\w{1,2}:\w{1,2}:\w{1,2}:\w{1,2}:\w{1,2})`)
	re_netmask    = regexp.MustCompile(`Mask:([0-9]+\.[0-9]+\.[0-9]+\.[0-9]+)`)
)

type interface_collector struct {
}

func NewInterfaceCollector() Collector {
	return &interface_collector{}
}

func (ic *interface_collector) Collect() (facts [][2]string) {
	faces := get_faces()
	if len(faces) == 0 {
		return
	}
	facts = append(facts, [2]string{"interfaces", strings.Join(faces, ",")})

	var ipaddress, ipaddress6, netmask, macaddress []string
	for _, face := range faces {
		output, err := run_ifconfig(face)
		if err != nil {
			continue
		}
		var matches []string
		// ipv4
		matches = re_ipaddress.FindStringSubmatch(output)
		if len(matches) >= 2 {
			facts = append(facts, [2]string{"ipaddress_" + face, matches[1]})
			ipaddress = append(ipaddress, matches[1])
		}
		// ipv6
		matches = re_ipaddress6.FindStringSubmatch(output)
		if len(matches) >= 2 {
			if !isLocalIp6(matches[1]) {
				facts = append(facts, [2]string{"ipaddress6_" + face, matches[1]})
				ipaddress6 = append(ipaddress6, matches[1])
			}
		}
		// netmask
		matches = re_netmask.FindStringSubmatch(output)
		if len(matches) >= 2 {
			facts = append(facts, [2]string{"netmask_" + face, matches[1]})
			netmask = append(netmask, matches[1])
		}
		// macaddress
		matches = re_macaddress.FindStringSubmatch(output)
		if len(matches) >= 2 {
			facts = append(facts, [2]string{"macaddress_" + face, matches[1]})
			macaddress = append(macaddress, matches[1])
		}
	}
	if len(ipaddress) > 0 {
		facts = append(facts, [2]string{"ipaddress", ipaddress[0]})
	}
	if len(ipaddress6) > 0 {
		facts = append(facts, [2]string{"ipaddress6", ipaddress6[0]})
	}
	if len(netmask) > 0 {
		facts = append(facts, [2]string{"netmask", netmask[0]})
	}
	if len(macaddress) > 0 {
		facts = append(facts, [2]string{"macaddress", macaddress[0]})
	}
	return
}
