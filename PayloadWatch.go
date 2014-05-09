package gitbot

import "fmt"

type PayloadWatch struct {
	Action string      `json:"action"`     // The action that was performed. Currently, can only be started.
	Repo   *Repository `json:"repository"` // Repo
	Sender *User       `json:"sender"`     // Sender
}

func (s PayloadWatch) String() string {
	return fmt.Sprintf("[%s] %s has %s", s.Repo, s.Sender, s.Action)
}
