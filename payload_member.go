package gitbot

import "fmt"

// PayloadMember when membership of an repo changes
type PayloadMember struct {
	Member *User       `json:"member"`     // The user that was added.
	Action string      `json:"action"`     // The action that was performed. Currently, can only be “added”.
	Repo   *Repository `json:"repository"` // Repo
	Sender *User       `json:"sender"`     // Sender
}

func (s PayloadMember) String() string {
	return fmt.Sprintf("[%s] %s has %s %s as a collaborator", s.Repo, s.Sender, s.Action, s.Member)
}
