package gitbot

import "time"

type CommitComment struct {
	HtmlUrl   string    `json:"html_url"`
	Url       string    `json:"url"`
	Id        int       `json:"id"`
	Body      string    `json:"body"`
	Path      string    `json:"path"`
	Position  int       `json:"position"`
	Line      int       `json:"line"`
	CommitId  string    `json:"commit_id"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
