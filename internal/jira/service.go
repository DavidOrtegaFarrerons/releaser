package jira

func AllReleaseIssues() []Ticket {
	return releaseVersionIssues().Tickets
}
