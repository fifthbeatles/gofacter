package facter

import (
	"strconv"
)

func KBToMB(KB int64) (MB string) {
	MB = strconv.Itoa(int(KB)/1024) + " MB"
	return
}
