package facter

import (
	"fmt"
)

type facter struct {
	facts      map[string]string
	collectors []Collector
}

func NewFacter() *facter {
	f := facter{facts: make(map[string]string)}
	f.collectors = append(f.collectors, NewPathCollector())
	f.collectors = append(f.collectors, NewInterfaceCollector())
	f.collectors = append(f.collectors, NewMemoryCollector())
	f.collectors = append(f.collectors, NewUptimeCollector())
	f.collectors = append(f.collectors, NewTimezoneCollector())
	f.collectors = append(f.collectors, NewHostnameCollector())
	return &f
}

func (f *facter) Collect() {
	for _, c := range f.collectors {
		for _, fact_pair := range c.Collect() {
			f.facts[fact_pair[0]] = fact_pair[1]
		}
	}
}

func (f *facter) Value(name string) (value string, ok bool) {
	value, ok = f.facts[name]
	return
}

func (f *facter) SetValue(name, value string) {
	f.facts[name] = value
}

func (f *facter) PrintAllValues() {
	for name, value := range f.facts {
		fmt.Println(name, "=>", value)
	}
}
