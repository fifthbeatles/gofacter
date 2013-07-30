package facter

import (
	"fmt"
)

type facter struct {
	facts map[string]string
}

func NewFacter() *facter {
	return &facter{make(map[string]string)}
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
