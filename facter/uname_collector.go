package facter

import (
	"strings"
)

type uname_collector struct {
}

func NewUnameCollector() Collector {
	return &uname_collector{}
}

func (uc *uname_collector) Collect() (facts [][2]string) {
	facts = append(facts, [2]string{"hardwareisa", run_uname("-p")})
	facts = append(facts, [2]string{"hardwaremodel", run_uname("-m")})
	facts = append(facts, [2]string{"kernel", run_uname("-s")})
	release := run_uname("-r")
	version := strings.Split(release, "-")[0]
	facts = append(facts, [2]string{"kernelrelease", release})
	facts = append(facts, [2]string{"kernelversion", version})
	return
}
