package gitbot

import (
	"fmt"
)

type Issue struct {
	User        *User        `json:"user"`
	Assignee    *User        `json:"assignee"`
	Milestone   *Milestone   `json:"milestone"`
	PullRequest *PullRequest `json:"pull_request"`
	ClosedBy    *User        `json:"closed_by"`
	Labels      []*Label     `json:"labels"`
	Url         string       `json:"url"`
	HtmlUrl     string       `json:"html_url"`
	Number      int          `json:"number"`
	State       string       `json:"state"`
	Title       string       `json:"title"`
	Body        string       `json:"body"`
	Comments    int          `json:"comments"`
	ClosedAt    string       `json:"closed_at"`
	//CreatedAt   NullTime     `json:"created_at"`
	//UpdatedAt   NullTime     `json:"updated_at"`
}

func (s Issue) String() string {
	return fmt.Sprintf("#%v (%s)", s.Number, s.Title)
}
