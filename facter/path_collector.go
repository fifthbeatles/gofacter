package facter

import "os"

var env_path = os.Getenv("PATH")

type path_collector struct {
}

func NewPathCollector() Collector {
	return &path_collector{}
}

func (pc *path_collector) Collect() (fact_name, fact_value string) {
	fact_name = "path"
	fact_value = env_path
	return
}
