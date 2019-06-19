package gitbot

import (
	"fmt"
)

type IssueComment struct {
	User    *User  `json:"user"`
	Id      int    `json:"id"`
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
	Body    string `json:"body"`
	//CreatedAt NullTime `json:"created_at"`
	//UpdatedAt NullTime `json:"updated_at"`
}

func (s IssueComment) String() string {
	return fmt.Sprintf("%s", s.HtmlUrl)
}
