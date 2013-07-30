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
	return &f
}

func (f *facter) Collect() {
	for _, c := range f.collectors {
		fact_name, fact_value := c.Collect()
		f.facts[fact_name] = fact_value
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
