package version

import (
	"fmt"
	"runtime"
)

// 构建时的版本信息
type VersionInfo struct {
	GitTag       string `json:"git_tag"`
	GitCommit    string `json:"git_commit"`
	GitTreeState string `json:"git_tree_state"`
	BuildDate    string `json:"build_date"`
	GoVersion    string `json:"go_version"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func (info VersionInfo) String() string {
	return info.GitTag
}

func Get() VersionInfo {
	return VersionInfo{
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
