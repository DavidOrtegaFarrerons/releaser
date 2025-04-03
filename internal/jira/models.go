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
	DisplayName string     `json:"displayName"`
	AvatarUrls  AvatarUrls `json:"avatarUrls"`
}

type AvatarUrls struct {
	Size48 string `json:"48x48"`
	Size32 string `json:"32x32"`
	Size24 string `json:"24x24"`
	Size16 string `json:"16x16"`
}

type Status struct {
	Name string `json:"name"`
}
