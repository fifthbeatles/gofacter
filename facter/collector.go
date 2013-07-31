package facter

type Collector interface {
	// return [[key1, value1], [key2, value2],...]
	Collect() (facts [][2]string)
}
