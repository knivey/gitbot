package gitbot

import "fmt"

type PayloadRelease struct {
	Action  string      `json:"action"`     // The action that was performed. Currently, can only be “published”.
	Release *Release    `json:"release"`    // The release itself.
	Repo    *Repository `json:"repository"` // Repo
	Sender  *User       `json:"sender"`     // Sender
}

func (s PayloadRelease) String() string {
	return fmt.Sprintf("[%s] %s has %s a release: %s", s.Repo, s.Sender, s.Action, s.Release)
}
