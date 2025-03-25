package models

import (
	"release-handler/internal/jira"
	"release-handler/internal/scm/azure"
)

type TableTicket struct {
	PullRequest *azure.PullRequest
	Ticket      *jira.Ticket
}
