package facter

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type uptimeCollector struct {
}

func NewUptimeCollector() Collector {
	return &uptimeCollector{}
}

func (uc *uptimeCollector) Collect() (facts []Fact) {
	content, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		return
	}
	uptimeSecondsFloat, err := strconv.ParseFloat(strings.Split(string(content), " ")[0], 64)
	if err != nil {
		return
	}
	uptimeSeconds := int64(uptimeSecondsFloat)
	uptimeHours := uptimeSeconds / 3600
	uptimeDays := uptimeHours / 24
	facts = append(facts, Fact{"uptimeSeconds", strconv.Itoa(int(uptimeSeconds))})
	facts = append(facts, Fact{"uptimeHours", strconv.Itoa(int(uptimeHours))})
	facts = append(facts, Fact{"uptimeDays", strconv.Itoa(int(uptimeDays))})

	hours := uptimeHours % 24
	minutes := (uptimeSeconds / 60) % 60
	var uptimeStr string
	switch uptimeDays {
	case 0:
	case 1:
		uptimeStr += "1 day "
	default:
		uptimeStr = uptimeStr + strconv.Itoa(int(uptimeDays)) + " days "
	}
	uptimeStr = uptimeStr + strconv.Itoa(int(hours)) + ":" + strconv.Itoa(int(minutes))
	facts = append(facts, Fact{"uptime", uptimeStr})
	return
}
