package main

import (
	"fmt"
)

// new fields may be added later
type Facter struct {
	facts map[string]string
}

func NewFacter() *Facter {
	return &Facter{make(map[string]string)}
}

func (facter *Facter) Value(name string) (value string, ok bool) {
	value, ok = facter.facts[name]
	return
}

func (facter *Facter) SetValue(name, value string) {
	facter.facts[name] = value
}

func (facter *Facter) PrintAllValues() {
	for name, value := range facter.facts {
		fmt.Println(name, "=>", value)
	}
}
