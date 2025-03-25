package azure

import "time"

type Response struct {
	PullRequests []PullRequest `json:"value"`
}

type PullRequest struct {
	Id           int        `json:"pullRequestId"`
	Status       string     `json:"status"`
	CreatedBy    User       `json:"createdBy"`
	BranchName   string     `json:"sourceRefName"`
	Url          string     `json:"url"`
	CreationDate time.Time  `json:"creationDate"`
	Reviewers    []Reviewer `json:"reviewers"`
}

type User struct {
	DisplayName string `json:"displayName"`
}

type Reviewer struct {
	Vote int `json:"vote"`
}
