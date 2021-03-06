package gitbot

import "fmt"

// PayloadWatch when a watcher is added
// Currently Github is confusing and sends this as well when a star occurs
type PayloadWatch struct {
	Action string      `json:"action"`     // The action that was performed. Currently, can only be started.
	Repo   *Repository `json:"repository"` // Repo
	Sender *User       `json:"sender"`     // Sender
}

func (s PayloadWatch) String() string {
	if s.Action == "started" {
		return fmt.Sprintf("[%s] was starred by %s", s.Repo, s.Sender)
	}
	return fmt.Sprintf("[%s] %s has performed an unknown watch action", s.Repo, s.Sender)
}
