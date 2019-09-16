package gitbot

import "fmt"

// PayloadFork when someone forks a repo
type PayloadFork struct {
	Forkee *Repository `json:"forkee"`     // The created repository.
	Repo   *Repository `json:"repository"` // Repo
	Sender *User       `json:"sender"`     // Sender
}

func (s PayloadFork) String() string {
	return fmt.Sprintf("[%s] %s has forked the repository (%s)", s.Repo, s.Sender, s.Forkee)
}
