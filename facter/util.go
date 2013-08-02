package facter

import (
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var (
	re_domain   = `.+\..+`
	re_hostname = regexp.MustCompile(`(.*?)\.(.+\..+$)`)
)

func KBToMB(KB int64) (MB string) {
	MB = strconv.Itoa(int(KB)/1024) + " MB"
	return
}

func run_ifconfig(arg ...string) (output string, err error) {
	os.Setenv("LANG", "C")
	output_bytes, err := exec.Command("/sbin/ifconfig", arg...).Output()
	if err == nil {
		output = string(output_bytes)
	}
	return
}

func get_faces() (faces []string) {
	output, err := run_ifconfig()
	if err != nil {
		return
	}
	for _, line := range strings.Split(output, "\n") {
		if matches := re_interface.FindStringSubmatch(line); len(matches) >= 2 && matches[1] != "lo" {
			faces = append(faces, matches[1])
		}
	}
	return
}

func isLocalIp6(ip6 string) bool {
	return ip6 == "::1" || strings.HasPrefix(strings.ToLower(ip6), "fe80")
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

func run_uname(arg ...string) (output string) {
	os.Setenv("LANG", "C")
	output_bytes, err := exec.Command("/bin/uname", arg...).Output()
	if err == nil {
		output = strings.TrimRight(string(output_bytes), "\n")
	}
	return
}
