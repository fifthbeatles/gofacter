package facter

import (
	"os/exec"
	"regexp"
	"strings"
)

var (
	re_hostname = regexp.MustCompile(`(.*?)\.(.+\..+$)`)
	re_domain   = `.+\..+`
)

type hostname_collector struct {
}

func NewHostnameCollector() Collector {
	return &hostname_collector{}
}

func (hc *hostname_collector) Collect() (facts [][2]string) {
	hostname, domain := getHostnameAndDomain()
	if len(hostname) > 0 {
		facts = append(facts, [2]string{"hostname", hostname})
		if len(domain) > 0 {
			facts = append(facts, [2]string{"domain", domain})
			facts = append(facts, [2]string{"fqdn", hostname + "." + domain})
		}
	}
	return
}

func getHostnameAndDomain() (hostname, domain string) {
	output, err := exec.Command("hostname").Output()
	if err != nil {
		return
	}
	hostname = strings.TrimRight(string(output), "\n")
	matches := re_hostname.FindStringSubmatch(hostname)
	if len(matches) >= 3 {
		hostname = matches[1]
		domain = matches[2]
		return
	}
	output, err = exec.Command("dnsdomainname").Output()
	if err == nil {
		domain = strings.TrimRight(string(output), "\n")
		if ok, _ := regexp.MatchString(re_domain, domain); ok {
			return
		}
	}
	return
}
