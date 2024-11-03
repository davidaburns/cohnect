package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	Name = "cohnect"
	Version = "0.0.1"
	BuildNumber = 0
	BuildTime = ""
	BuildType = ""
	Commit = ""
)

func main() {
	versionFlag := flag.Bool("version", false, "Display version information")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("Version: %s-%s-%s:%d [%s]\n", Name, Version, BuildType, BuildNumber, Commit)
		os.Exit(0)
	}
}