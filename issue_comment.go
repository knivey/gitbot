package gitbot

import (
	"fmt"
)

// IssueComment is a comment on an issue
type IssueComment struct {
	User    *User  `json:"user"`
	ID      int    `json:"id"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
	Body    string `json:"body"`
	//CreatedAt NullTime `json:"created_at"`
	//UpdatedAt NullTime `json:"updated_at"`
}

func (s IssueComment) String() string {
	return fmt.Sprintf("%s", s.HTMLURL)
}
