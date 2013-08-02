/*
A tool to collect information of the system, similar to facter writen in ruby
*/
package main

import (
	"flag"
	"fmt"
	"github.com/fifthbeatles/gofacter/facter"
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

	f := facter.NewFacter()
	f.Collect()
	f.PrintAllValues()
	os.Exit(0)
}
