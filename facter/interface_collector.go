package facter

import (
	"regexp"
	"strings"
)

var (
	reIpaddress  = regexp.MustCompile(`inet addr:([0-9]+\.[0-9]+\.[0-9]+\.[0-9]+)`)
	reIpaddress6 = regexp.MustCompile(`inet6 addr: ((?:[0-9,a-f,A-F]*\:{1,2})+[0-9,a-f,A-F]{0,4})/`)
	reMacaddress = regexp.MustCompile(`(?:ether|HWaddr)\s+(\w{1,2}:\w{1,2}:\w{1,2}:\w{1,2}:\w{1,2}:\w{1,2})`)
	reNetmask    = regexp.MustCompile(`Mask:([0-9]+\.[0-9]+\.[0-9]+\.[0-9]+)`)
)

type interfaceCollector struct {
}

func NewInterfaceCollector() Collector {
	return &interfaceCollector{}
}

func (ic *interfaceCollector) Collect() (facts []Fact) {
	faces := getFaces()
	if len(faces) == 0 {
		return
	}
	facts = append(facts, Fact{"interfaces", strings.Join(faces, ",")})

	var ipaddress, ipaddress6, netmask, macaddress []string
	for _, face := range faces {
		output, err := runIfconfig(face)
		if err != nil {
			continue
		}
		var matches []string
		// ipv4
		matches = reIpaddress.FindStringSubmatch(output)
		if len(matches) >= 2 {
			facts = append(facts, Fact{"ipaddress_" + face, matches[1]})
			ipaddress = append(ipaddress, matches[1])
		}
		// ipv6
		matches = reIpaddress6.FindStringSubmatch(output)
		if len(matches) >= 2 {
			if !isLocalIp6(matches[1]) {
				facts = append(facts, Fact{"ipaddress6_" + face, matches[1]})
				ipaddress6 = append(ipaddress6, matches[1])
			}
		}
		// netmask
		matches = reNetmask.FindStringSubmatch(output)
		if len(matches) >= 2 {
			facts = append(facts, Fact{"netmask_" + face, matches[1]})
			netmask = append(netmask, matches[1])
		}
		// macaddress
		matches = reMacaddress.FindStringSubmatch(output)
		if len(matches) >= 2 {
			facts = append(facts, Fact{"macaddress_" + face, matches[1]})
			macaddress = append(macaddress, matches[1])
		}
	}
	if len(ipaddress) > 0 {
		facts = append(facts, Fact{"ipaddress", ipaddress[0]})
	}
	if len(ipaddress6) > 0 {
		facts = append(facts, Fact{"ipaddress6", ipaddress6[0]})
	}
	if len(netmask) > 0 {
		facts = append(facts, Fact{"netmask", netmask[0]})
	}
	if len(macaddress) > 0 {
		facts = append(facts, Fact{"macaddress", macaddress[0]})
	}
	return
}
