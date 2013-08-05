package facter

import "os"

var env_path = os.Getenv("PATH")

type path_collector struct {
}

func NewPathCollector() Collector {
	return &path_collector{}
}

func (pc *path_collector) Collect() (facts []Fact) {
	facts = append(facts, Fact{"path", env_path})
	return
}
