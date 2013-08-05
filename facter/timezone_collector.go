package facter

import (
	"strconv"
	"time"
)

type timezoneCollector struct {
}

func NewTimezoneCollector() Collector {
	return &timezoneCollector{}
}

func (tc *timezoneCollector) Collect() (facts []Fact) {
	now := time.Now()
	zone, offset := now.Zone()
	facts = append(facts, Fact{"timezone", zone})
	facts = append(facts, Fact{"timezone_offset", strconv.Itoa(offset)})
	return
}
