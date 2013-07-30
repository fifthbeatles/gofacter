package facter

type Collector interface {
	Collect() (fact_name, fact_value string)
}
