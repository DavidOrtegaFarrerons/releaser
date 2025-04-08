package helper

import (
	"fmt"
	"github.com/spf13/viper"
	"release-handler/config"
	"release-handler/internal/jira"
	"release-handler/internal/models"
	"release-handler/internal/scm/azure"
)

func GenerateUrl(urlType string, ticket models.TableTicket) string {
	switch urlType {
	case "azure":
		if ticket.PullRequest != nil {
			GeneratePullRequestUrl(ticket.PullRequest)
		}
	case "jira":
		if ticket.Ticket != nil {
			GenerateTicketUrl(ticket.Ticket)
		}
	}
	return ""
}

func GenerateTicketUrl(ticket *jira.Ticket) string {
	return fmt.Sprintf("https://%s.atlassian.net/browse/%s",
		viper.GetString(config.JiraDomain),
		ticket.Key,
	)
}

func GeneratePullRequestUrl(pr *azure.PullRequest) string {
	return fmt.Sprintf("https://dev.azure.com/%s/%s/_git/%s/pullrequest/%d",
		viper.GetString(config.AzureOrganization),
		viper.GetString(config.AzureProject),
		viper.GetString(config.AzureRepositoryId),
		pr.Id,
	)
}
