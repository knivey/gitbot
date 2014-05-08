package gitbot

import "time"

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
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
