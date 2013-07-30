package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	APP_VERSION = "0.1"
)

// command-line flags
var (
	versionFlag = flag.Bool("v", false, "Print the version number.")
)

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}
	os.Exit(0)
}
