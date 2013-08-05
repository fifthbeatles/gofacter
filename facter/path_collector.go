package facter

import "os"

type pathCollector struct {
}

func NewPathCollector() Collector {
	return &pathCollector{}
}

func (pc *pathCollector) Collect() (facts []Fact) {
	facts = append(facts, Fact{"path", os.Getenv("PATH")})
	return
}
