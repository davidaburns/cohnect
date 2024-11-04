package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	name = "cohnect"
	version = "0.0.1"
	buildNumber = 0
	buildTime = ""
	buildType = ""
	commit = ""
)

type BuildInfo struct {
	Name string
	Version string
	BuildNumber int
	BuildTime string
	BuildType string
	BuildCommit string
}

func NewBuildInfo() *BuildInfo {
	return &BuildInfo {
		Name: name,
		Version: version,
		BuildNumber: buildNumber,
		BuildTime: buildTime,
		BuildType: buildType,
		BuildCommit: commit,
	}
}

func (build *BuildInfo) ToString() string {
	return fmt.Sprintf("%s-%s.%s:b%d[%s]", build.Name, build.Version, build.BuildType, build.BuildNumber, build.BuildCommit)
}

type Config struct {
	Server struct {
		Host string `mapstructure:"host"`
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Logger struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"logger"`
}

func LoadConfigFile(file string) (*Config, error) {
	var config Config
	viper.SetConfigFile(file)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}