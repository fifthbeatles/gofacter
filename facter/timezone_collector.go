package facter

import (
	"strconv"
	"time"
)

type timezone_collector struct {
}

func NewTimezoneCollector() Collector {
	return &timezone_collector{}
}

func (tc *timezone_collector) Collect() (facts []Fact) {
	now := time.Now()
	zone, offset := now.Zone()
	facts = append(facts, Fact{"timezone", zone})
	facts = append(facts, Fact{"timezone_offset", strconv.Itoa(offset)})
	return
}
