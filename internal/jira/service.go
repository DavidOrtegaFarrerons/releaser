package jira

func AllReleaseIssues(releaseName string) []Ticket {
	return releaseVersionIssues(releaseName).Tickets
}
