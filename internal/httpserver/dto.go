package httpserver

import (
	"release-handler/internal/helper"
	"release-handler/internal/jira"
	"release-handler/internal/models"
	"release-handler/internal/scm/azure"
	"strings"
	"time"
)

type PullRequestDTO struct {
	Id           int    `json:"id"`
	Status       string `json:"status"`
	CreatedBy    string `json:"createdBy"`
	BranchName   string `json:"branchName"`
	Url          string `json:"url"`
	CreationDate string `json:"creationDate"`
	MergeStatus  string `json:"mergeStatus"`
}

type ReviewerDTO struct {
	Vote string `json:"vote"`
}

type TicketDTO struct {
	Id       string            `json:"id"`
	Key      string            `json:"key"`
	Assignee TicketAssigneeDTO `json:"assignee"`
	Url      string            `json:"url"`
	Status   string            `json:"status"`
}

type TicketAssigneeDTO struct {
	DisplayName  string `json:"displayName"`
	ProfileImage string `json:"profileImage"`
}

type TableTicketDTO struct {
	PullRequest *PullRequestDTO `json:"pullRequest,omitempty"`
	Ticket      *TicketDTO      `json:"ticket,omitempty"`
}

func ToPullRequestDTO(pr *azure.PullRequest) *PullRequestDTO {
	if pr == nil {
		return nil
	}

	return &PullRequestDTO{
		Id:           pr.Id,
		MergeStatus:  pr.MergeStatus,
		CreatedBy:    pr.CreatedBy.DisplayName,
		BranchName:   strings.Replace(pr.BranchName, "refs/heads/", "", 1),
		Url:          helper.GeneratePullRequestUrl(pr),
		CreationDate: pr.CreationDate.Format(time.RFC3339),
		Status:       azure.GetFinalReviewStatus(pr.Reviewers),
	}
}

func ToTicketDTO(ticket *jira.Ticket) *TicketDTO {
	if ticket == nil {
		return nil
	}

	return &TicketDTO{
		Id:       ticket.Id,
		Key:      ticket.Key,
		Assignee: toTicketAssigneeDto(ticket.Fields.Assignee),
		Url:      helper.GenerateTicketUrl(ticket),
		Status:   ticket.Fields.Status.Name,
	}
}

func ToTableTicketDTO(tableTicket models.TableTicket) TableTicketDTO {
	return TableTicketDTO{
		PullRequest: ToPullRequestDTO(tableTicket.PullRequest),
		Ticket:      ToTicketDTO(tableTicket.Ticket),
	}
}

func toTicketAssigneeDto(assignee jira.Assignee) TicketAssigneeDTO {
	return TicketAssigneeDTO{
		DisplayName:  assignee.DisplayName,
		ProfileImage: assignee.AvatarUrls.Size32,
	}
}
