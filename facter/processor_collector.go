package facter

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

var (
	re_processor_desc = regexp.MustCompile(`model name\s+:\s+(.*)\s*`)
)

type processor_collector struct {
}

func NewProcessorCollector() Collector {
	return &processor_collector{}
}

func (pc *processor_collector) Collect() (facts [][2]string) {
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
		matches := re_processor_desc.FindStringSubmatch(string(line))
		if len(matches) >= 2 {
			facts = append(facts, [2]string{"processor" + strconv.Itoa(count), matches[1]})
			count++
		}
	}
	facts = append(facts, [2]string{"processor_count", strconv.Itoa(count)})
	return
}
