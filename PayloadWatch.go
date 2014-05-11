package gitbot

import "fmt"

type PayloadWatch struct {
	Action string      `json:"action"`     // The action that was performed. Currently, can only be started.
	Repo   *Repository `json:"repository"` // Repo
	Sender *User       `json:"sender"`     // Sender
}

func (s PayloadWatch) String() string {
	if s.Action == "started" {
		return fmt.Sprintf("[%s] was starred by %s", s.Repo, s.Sender)
	}
}
