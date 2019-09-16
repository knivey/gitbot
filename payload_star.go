package gitbot

import "fmt"

// PayloadStar when a star is added or removed
// StarredAt will be nil when it's a delete action
// Currently Github is confusing and sends this as well when a star occurs
type PayloadStar struct {
	Action    string      `json:"action"`     // The action that was performed: created | deleted
	StarredAt NullTime    `json:"starred_at"` // Null when deleted
	Repo      *Repository `json:"repository"` // Repo
	Sender    *User       `json:"sender"`     // Sender
}

func (s PayloadStar) String() string {
	switch s.Action {
	case "created":
		return fmt.Sprintf("[%s] was starred by %s", s.Repo, s.Sender)
	case "deleted":
		return fmt.Sprintf("[%s] was unstarred by %s", s.Repo, s.Sender)
	}
	return fmt.Sprintf("[%s] %s has performed an unknown starring action", s.Repo, s.Sender)
}
