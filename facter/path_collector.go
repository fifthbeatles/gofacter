package facter

import "os"

var env_path = os.Getenv("PATH")

type path_collector struct {
}

func NewPathCollector() Collector {
	return &path_collector{}
}

func (pc *path_collector) Collect() (facts [][2]string) {
	facts = append(facts, [2]string{"path", env_path})
	return
}
