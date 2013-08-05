package facter

import (
	"strings"
)

type uname_collector struct {
}

func NewUnameCollector() Collector {
	return &uname_collector{}
}

func (uc *uname_collector) Collect() (facts []Fact) {
	facts = append(facts, Fact{"hardwareisa", run_uname("-p")})
	facts = append(facts, Fact{"hardwaremodel", run_uname("-m")})
	facts = append(facts, Fact{"kernel", run_uname("-s")})
	release := run_uname("-r")
	version := strings.Split(release, "-")[0]
	facts = append(facts, Fact{"kernelrelease", release})
	facts = append(facts, Fact{"kernelversion", version})
	return
}
