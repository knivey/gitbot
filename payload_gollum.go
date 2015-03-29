package gitbot

import "fmt"

type PayloadGollum struct {
	Pages  []*Page     `json:"pages"`      // The pages that were updated.
	Repo   *Repository `json:"repository"` // Repo
	Sender *User       `json:"sender"`     // Sender
}

func (s PayloadGollum) String() string {
	return fmt.Sprintf("[%s] %s updated pages", s.Repo, s.Sender)
}
