package facter

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type uptime_collector struct {
}

func NewUptimeCollector() Collector {
	return &uptime_collector{}
}

func (uc *uptime_collector) Collect() (facts [][2]string) {
	content, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		return
	}
	uptime_seconds_float, err := strconv.ParseFloat(strings.Split(string(content), " ")[0], 64)
	if err != nil {
		return
	}
	uptime_seconds := int64(uptime_seconds_float)
	uptime_hours := uptime_seconds / 3600
	uptime_days := uptime_hours / 24
	facts = append(facts, [2]string{"uptime_seconds", strconv.Itoa(int(uptime_seconds))})
	facts = append(facts, [2]string{"uptime_hours", strconv.Itoa(int(uptime_hours))})
	facts = append(facts, [2]string{"uptime_days", strconv.Itoa(int(uptime_days))})

	hours := uptime_hours % 24
	minutes := (uptime_seconds / 60) % 60
	var uptime_str string
	switch uptime_days {
	case 0:
	case 1:
		uptime_str += "1 day "
	default:
		uptime_str = uptime_str + strconv.Itoa(int(uptime_days)) + " days "
	}
	uptime_str = uptime_str + strconv.Itoa(int(hours)) + ":" + strconv.Itoa(int(minutes))
	facts = append(facts, [2]string{"uptime", uptime_str})
	return
}
