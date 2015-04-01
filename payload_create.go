package gitbot

import "fmt"

type PayloadCreate struct {
	Ref          string      `json:"ref"`
	RefType      string      `json:"ref_type"`
	MasterBranch string      `json:"master_branch"`
	Description  string      `json:"description"`
	PusherType   string      `json:"pusher_type"`
	Repo         *Repository `json:"repository"` // Repo
	Sender       *User       `json:"sender"`     // Sender
}

func (p PayloadCreate) String() string {
	return fmt.Sprintf("[%s] %s created %s %s", p.Repo, p.Sender, p.RefType, p.Ref)
}
