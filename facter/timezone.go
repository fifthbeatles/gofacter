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

func (tc *timezone_collector) Collect() (facts [][2]string) {
	now := time.Now()
	zone, offset := now.Zone()
	facts = append(facts, [2]string{"timezone", zone})
	facts = append(facts, [2]string{"timezone_offset", strconv.Itoa(offset)})
	return
}
