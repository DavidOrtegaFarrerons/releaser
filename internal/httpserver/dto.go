package httpserver

import (
	"fmt"
	"release-handler/internal/jira"
	"release-handler/internal/models"
	"release-handler/internal/scm/azure"
	"time"
)

type PullRequestDTO struct {
	Id           int    `json:"pullRequestId"`
	Status       string `json:"status"`
	CreatedBy    string `json:"createdBy"`
	BranchName   string `json:"sourceRefName"`
	Url          string `json:"url"`
	CreationDate string `json:"creationDate"`
	MergeStatus  string `json:"mergeStatus"`
}

type ReviewerDTO struct {
	Vote string `json:"vote"`
}

type TicketDTO struct {
	Id       string `json:"id"`
	Key      string `json:"key"`
	Assignee string `json:"assignee"`
	Status   string `json:"status"`
}

type TableTicketDTO struct {
	PullRequest *PullRequestDTO `json:"pullRequest,omitempty"`
	Ticket      *TicketDTO      `json:"ticket,omitempty"`
}

func ToPullRequestDTO(pr *azure.PullRequest) *PullRequestDTO {
	fmt.Println(pr)
	if pr == nil {
		return nil
	}

	return &PullRequestDTO{
		Id:           pr.Id,
		Status:       pr.Status,
		CreatedBy:    pr.CreatedBy.DisplayName,
		BranchName:   pr.BranchName,
		Url:          pr.Url,
		CreationDate: pr.CreationDate.Format(time.RFC3339),
		MergeStatus:  azure.GetFinalReviewStatus(pr.Reviewers),
	}
}

func ToTicketDTO(ticket *jira.Ticket) *TicketDTO {
	if ticket == nil {
		return nil
	}

	return &TicketDTO{
		Id:       ticket.Id,
		Key:      ticket.Key,
		Assignee: ticket.Fields.Assignee.DisplayName,
		Status:   ticket.Fields.Status.Name,
	}
}

func ToTableTicketDTO(tableTicket models.TableTicket) TableTicketDTO {
	fmt.Println(tableTicket)
	return TableTicketDTO{
		PullRequest: ToPullRequestDTO(tableTicket.PullRequest),
		Ticket:      ToTicketDTO(tableTicket.Ticket),
	}
}
