package gitbot

import (
	"fmt"
)

// CommitComment is a bespoke comment on a commit
type CommitComment struct {
	HTMLURL  string `json:"html_url"`
	URL      string `json:"url"`
	ID       int    `json:"id"`
	Body     string `json:"body"`
	Path     string `json:"path"`
	Position int    `json:"position"`
	Line     int    `json:"line"`
	CommitID string `json:"commit_id"`
	User     *User  `json:"user"`
	//CreatedAt NullTime `json:"created_at"`
	//UpdatedAt NullTime `json:"updated_at"`
}

func (s CommitComment) String() string {
	if s.Path != "" {
		return fmt.Sprintf("on %s:%d by %s %s", s.Path, s.Line, s.User, s.HTMLURL)
	}
	return fmt.Sprintf("by %s %s", s.User, s.HTMLURL)
}
