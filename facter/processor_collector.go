package facter

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

var (
	reProcessor = regexp.MustCompile(`model name\s+:\s+(.*)\s*`)
)

type processorCollector struct {
}

func NewProcessorCollector() Collector {
	return &processorCollector{}
}

func (pc *processorCollector) Collect() (facts []Fact) {
	f, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	count := 0
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		matches := reProcessor.FindStringSubmatch(string(line))
		if len(matches) >= 2 {
			facts = append(facts, Fact{"processor" + strconv.Itoa(count), matches[1]})
			count++
		}
	}
	facts = append(facts, Fact{"processor_count", strconv.Itoa(count)})
	return
}
