package version

import "time"

var (
	gitTag       string    = ""
	gitCommit    string    = "$Format:%H$"    // sha1 from git, output of $(git rev-parse HEAD)
	gitTreeState string    = "not a git tree" // state of git tree, either "clean" or "dirty"
	buildDate    time.Time = time.Now()       // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)
