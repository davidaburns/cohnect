package config

var (
	name = "cohnect"
	version = "0.0.1"
	buildNumber = 0
	buildTime = ""
	buildType = ""
	commit = ""
)

type BuildConfig struct {
	Name string
	Version string
	BuildNumber int
	BuildTime string
	BuildType string
	BuildCommit string
}

func NewBuildConfig() *BuildConfig {
	return &BuildConfig {
		Name: name,
		Version: version,
		BuildNumber: buildNumber,
		BuildTime: buildTime,
		BuildType: buildType,
		BuildCommit: commit,
	}
}