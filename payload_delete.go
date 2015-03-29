package gitbot

import "fmt"

type PayloadDelete struct {
	RefType string      `json:"ref_type"`   //The object that was deleted. Can be “branch” or “tag”.
	Ref     string      `json:"ref"`        //The full git ref.
	Repo    *Repository `json:"repository"` // Repo
	Sender  *User       `json:"sender"`     // Sender
}

func (s PayloadDelete) String() string {
	return fmt.Sprintf("[%s] %s has deleted %s %s", s.Repo, s.Sender, s.RefType, s.Ref)
}
