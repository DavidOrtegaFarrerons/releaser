package jira

type Response struct {
	Tickets []Ticket `json:"issues"`
}

type Ticket struct {
	Id     string `json:"id"`
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
}

type Fields struct {
	Assignee Assignee `json:"assignee"`
	Status   Status   `json:"status"`
}

type Assignee struct {
	DisplayName string `json:"displayName"`
}

type Status struct {
	Name string `json:"name"`
}
