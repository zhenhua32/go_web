package version

// 这些值应该是从外部传入的
var (
	gitTag       string = ""
	gitCommit    string = "$Format:%H$"          // sha1 from git, output of $(git rev-parse HEAD)
	gitTreeState string = "not a git tree"       // state of git tree, either "clean" or "dirty"
	buildDate    string = "1970-01-01T00:00:00Z" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)
