package azure

func AllReleaseMergeRequests() []PullRequest {
	return ReleaseMergeRequests().PullRequests
}
