package facter

import (
	"strings"
)

type unameCollector struct {
}

func NewUnameCollector() Collector {
	return &unameCollector{}
}

func (uc *unameCollector) Collect() (facts []Fact) {
	facts = append(facts, Fact{"hardwareisa", runUname("-p")})
	facts = append(facts, Fact{"hardwaremodel", runUname("-m")})
	facts = append(facts, Fact{"kernel", runUname("-s")})
	release := runUname("-r")
	version := strings.Split(release, "-")[0]
	facts = append(facts, Fact{"kernelrelease", release})
	facts = append(facts, Fact{"kernelversion", version})
	return
}
