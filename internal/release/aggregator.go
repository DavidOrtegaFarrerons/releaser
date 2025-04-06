package release

import (
	"github.com/spf13/viper"
	"regexp"
	"release-handler/config"
	"release-handler/internal/jira"
	"release-handler/internal/models"
	"release-handler/internal/scm/azure"
)

func MergeTickets() []models.TableTicket {
	tickets := jira.AllReleaseIssues()
	pullRequests := azure.AllReleaseMergeRequests()

	ticketsAndMergeRequestsMap := make(map[string]models.TableTicket)

	createTicketMap(tickets, ticketsAndMergeRequestsMap)
	AddPullRequestsToTickets(pullRequests, ticketsAndMergeRequestsMap)

	orderedTickets := make([]models.TableTicket, 0)

	for _, issue := range tickets {
		if ticket, exists := ticketsAndMergeRequestsMap[issue.Key]; exists {
			orderedTickets = append(orderedTickets, ticket)
		}
	}

	return orderedTickets
}

func AddPullRequestsToTickets(pullRequests []azure.PullRequest, m map[string]models.TableTicket) {
	for _, mr := range pullRequests {
		ticketPrefix := ticketPrefix(mr.BranchName)
		if ticket, exists := m[ticketPrefix]; exists {
			m[mr.BranchName] = models.TableTicket{
				PullRequest: &mr,
				Ticket:      ticket.Ticket,
			}
		}
	}
}

func createTicketMap(issues []jira.Ticket, m map[string]models.TableTicket) {
	for _, issue := range issues {
		m[issue.Key] = models.TableTicket{Ticket: &issue}
	}
}

func ticketPrefix(value string) string {
	pattern := viper.GetString(config.TicketPrefix) + `-\d+`
	re := regexp.MustCompile(pattern)

	return re.FindString(value)
}
