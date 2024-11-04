package main

import (
	"fmt"
	"os"

	"github.com/davidaburns/cohnect/config"
)

func main() {
	build := config.NewBuildConfig();
	fmt.Printf("Running: %s-%s.%s:b%d[%s]", build.Name, build.Version, build.BuildType, build.BuildNumber, build.BuildCommit)
	os.Exit(0)
}