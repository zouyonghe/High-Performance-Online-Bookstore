package version

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"log"
	"runtime"
	"strings"
	"time"
)

// Info contains versioning information.
type Info struct {
	GitTag       string    `json:"gitTag"`
	GitCommit    string    `json:"gitCommit"`
	GitTreeState string    `json:"gitTreeState"`
	BuildDate    time.Time `json:"buildDate"`
	GoVersion    string    `json:"goVersion"`
	Compiler     string    `json:"compiler"`
	Platform     string    `json:"platform"`
}

// String returns info as a human-friendly version string.
func (info Info) String() string {
	return info.GitTag
}

func Get() Info {
	return Info{
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// GetGitInfo returns last git commit information.
func getLastCommit() string {
	r, _ := git.PlainOpen(".")

	// 获取 HEAD 指向的分支
	ref, _ := r.Head()

	// 获取由 ref 指向的提交对象
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		log.Fatal(err)
	}

	/*	repository, _ := git.PlainOpen(".")
		headRef, err := repository.Head()
		if err != nil {
			log.Fatal(err)
		}
		headSha := headRef.Hash().String()

		return headSha*/
	commitMessage := commit.Message
	commitMessage = strings.Replace(commitMessage, "\n", "", -1)
	return commitMessage
}

// getGitTag returns the git tag.
func getLastGitTag() string {
	repository, _ := git.PlainOpen(".")
	tagRefs, err := repository.Tags()
	if err != nil {
		log.Fatal(err)
	}

	var lastTagCommit *object.Commit
	var lastTagName string
	err = tagRefs.ForEach(func(tagRef *plumbing.Reference) error {
		revision := plumbing.Revision(tagRef.Name().String())
		tagCommitHash, err := repository.ResolveRevision(revision)
		if err != nil {
			return err
		}

		commit, err := repository.CommitObject(*tagCommitHash)
		if err != nil {
			return err
		}

		if lastTagCommit == nil {
			lastTagCommit = commit
			lastTagName = tagRef.Name().String()
		}

		if commit.Committer.When.After(lastTagCommit.Committer.When) {
			lastTagCommit = commit
			lastTagName = tagRef.Name().String()
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return lastTagName
}
