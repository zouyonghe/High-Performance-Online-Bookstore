package version

import (
	"fmt"
	"runtime"
	"time"
)

var (
	gitTag       = getLastGitTag()
	gitCommit    = getLastCommit()                          // sha1 from git, output of $(git rev-parse HEAD)
	gitTreeState = "not a git tree"                         // state of git tree, either "clean" or "dirty"
	buildDate    = time.Now().Format("2006-01-02 15:04:05") // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
	goVersion    = runtime.Version()
	compiler     = runtime.Compiler
	platform     = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)
