package version

import "fmt"

var (

	// URL is the git URL for the repository
	URL = "github.com/p9c/parallelcoin"
	// GitRef is the gitref, as in refs/heads/branchname
	GitRef = "refs/heads/main"
	// GitCommit is the commit hash of the current HEAD
	GitCommit = "a28dec888e7a24dae796ef202da722dc5ea80693"
	// BuildTime stores the time when the current binary was built
	BuildTime = "2021-04-01T18:03:53+02:00"
	// Tag lists the Tag on the p9build, adding a + to the newest Tag if the commit is
	// not that commit
	Tag = "+"
	// PathBase is the path base returned from runtime caller
	PathBase = "/home/loki/src/github.com/p9c/parallelcoin/"
)

// Get returns a pretty printed version information string
func Get() string {
	return fmt.Sprint(
		"ParallelCoin Pod\n"+
		"	git repository: "+URL+"\n",
		"	branch: "+GitRef+"\n"+
		"	commit: "+GitCommit+"\n"+
		"	built: "+BuildTime+"\n"+
		"	Tag: "+Tag+"\n",
	)
}
