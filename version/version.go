package version

import "fmt"

var (

	// URL is the git URL for the repository
	URL = "github.com/p9c/parallelcoin"
	// GitRef is the gitref, as in refs/heads/branchname
	GitRef = "refs/heads/main"
	// GitCommit is the commit hash of the current HEAD
	GitCommit = "bcb8389b84181b36e032444f9a3c55cec17c0a7f"
	// BuildTime stores the time when the current binary was built
	BuildTime = "2021-04-01T18:50:18+02:00"
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
