package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	VERSION = "0.1"
)

// command-line flags
var (
	versionFlag = flag.Bool("v", false, "Print the version number.")
)

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println(VERSION)
		os.Exit(0)
	}
	facter := NewFacter()
	facter.PrintAllValues()
	os.Exit(0)
}
