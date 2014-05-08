package gitbot

import "time"

type IssueComment struct {
	User      *User     `json:"user"`
	Id        int       `json:"id"`
	Url       string    `json:"url"`
	HtmlUrl   string    `json:"html_url"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
