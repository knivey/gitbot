package gitbot

import (
	"fmt"
)

type CommitComment struct {
	HtmlUrl  string `json:"html_url"`
	Url      string `json:"url"`
	Id       int    `json:"id"`
	Body     string `json:"body"`
	Path     string `json:"path"`
	Position int    `json:"position"`
	Line     int    `json:"line"`
	CommitId string `json:"commit_id"`
	User     *User  `json:"user"`
	//CreatedAt NullTime `json:"created_at"`
	//UpdatedAt NullTime `json:"updated_at"`
}

func (s CommitComment) String() string {
	if s.Path != "" {
		return fmt.Sprintf("on %s:%d by %s %s", s.Path, s.Line, s.User, s.HtmlUrl)
	}
	return fmt.Sprintf("by %s %s", s.User, s.HtmlUrl)
}
