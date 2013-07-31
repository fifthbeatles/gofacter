package facter

import (
	"bytes"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	re_inet_addr    = regexp.MustCompile(`inet addr:([0-9]+\.[0-9]+\.[0-9]+\.[0-9]+)`)
	re_inet6_addr   = regexp.MustCompile(`inet6 addr: ((?:[0-9,a-f,A-F]*\:{1,2})+[0-9,a-f,A-F]{0,4})/`)
	ifconfig_output = run_ifconfig()
)

func run_ifconfig() (output string) {
	os.Setenv("LANG", "C")
	cmd := exec.Command("/sbin/ifconfig")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return
	}
	output = out.String()
	return
}

type ip_collector struct {
}

func NewIpCollector() Collector {
	return &ip_collector{}
}

func (ic *ip_collector) Collect() (facts [][2]string) {
	// ipv4
	for _, matches := range re_inet_addr.FindAllStringSubmatch(ifconfig_output, -1) {
		if !strings.HasPrefix(matches[1], "127.") {
			facts = append(facts, [2]string{"ipaddress", matches[1]})
			break
		}
	}
	// ipv6
	for _, matches := range re_inet6_addr.FindAllStringSubmatch(ifconfig_output, -1) {
		if matches[1] != "::1" && !strings.HasPrefix(strings.ToLower(matches[1]), "fe80") {
			facts = append(facts, [2]string{"ipaddress6", matches[1]})
			break
		}
	}
	return
}
