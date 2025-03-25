package release

import (
	"github.com/spf13/viper"
	"regexp"
	"release-handler/config"
	"release-handler/internal/jira"
	"release-handler/internal/models"
	"release-handler/internal/scm/azure"
	"release-handler/internal/ui"
)

func CreateReleaseTable() {
	issues := jira.AllReleaseIssues()
	mergeRequests := azure.AllReleaseMergeRequests()

	ticketsAndMergeRequestsMap := make(map[string]models.TableTicket)

	createTicketMap(issues, ticketsAndMergeRequestsMap)
	fillTicketsWithMergeRequests(mergeRequests, ticketsAndMergeRequestsMap)

	ui.ReleaseTable(ticketsAndMergeRequestsMap)

}

func fillTicketsWithMergeRequests(mergeRequests []azure.PullRequest, m map[string]models.TableTicket) {
	for _, mr := range mergeRequests {
		ticketPattern := ticketPattern(mr.BranchName)
		if ticket, exists := m[ticketPattern]; exists {
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

func ticketPattern(value string) string {
	pattern := viper.GetString(config.TicketPrefix) + `-\d+`
	re := regexp.MustCompile(pattern)

	return re.FindString(value)
}
