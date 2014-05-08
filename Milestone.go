package gitbot

import "time"

type Milestone struct {
	Creator      *User     `json:"creator"`
	Url          string    `json:"url"`
	Number       int       `json:"number"`
	State        string    `json:"state"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	OpenIssues   int       `json:"open_issues"`
	ClosedIssues int       `json:"closed_issues"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DueOn        string    `json:"due_on"`
}
