package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	Name = "cohnect"
	Version = "0.0.1"
	Type = "dev"
	BuildNumber = 0
	BuildTime = ""
	Commit = ""
)

func main() {
	versionFlag := flag.Bool("version", false, "Display version information")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("Version: %s-%s.%d [%s]\n", Version, Type, BuildNumber, Commit)
		os.Exit(0)
	}
}