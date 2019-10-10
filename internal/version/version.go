package version

// GitCommit the git commit SHA
var GitCommit string

// CurrentVersion returns the current version
func CurrentVersion() string {
	return GitCommit
}
