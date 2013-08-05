package facter

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

var (
	re = regexp.MustCompile(`(\S+):\s+(\d+).*\n`)
)

type memoryCollector struct {
}

func NewMemoryCollector() Collector {
	return &memoryCollector{}
}

func (mc *memoryCollector) Collect() (facts []Fact) {
	f, err := os.Open("/proc/meminfo")
	if err != nil {
		return
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	var memTotal, memFree, swapTotal, swapFree int64
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		matches := re.FindStringSubmatch(line)
		switch matches[1] {
		case "MemTotal":
			memTotal, _ = strconv.ParseInt(matches[2], 10, 64)
		case "MemFree", "Buffers", "Cached":
			temp, _ := strconv.ParseInt(matches[2], 10, 64)
			memFree += temp
		case "SwapTotal":
			swapTotal, _ = strconv.ParseInt(matches[2], 10, 64)
		case "SwapFree":
			swapFree, _ = strconv.ParseInt(matches[2], 10, 64)
		}
	}
	facts = append(facts, Fact{"memorysize", KBToMB(memTotal)})
	facts = append(facts, Fact{"memoryfree", KBToMB(memFree)})
	facts = append(facts, Fact{"swapsize", KBToMB(swapTotal)})
	facts = append(facts, Fact{"swapfree", KBToMB(swapFree)})
	return
}
