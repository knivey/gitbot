package gitbot

import "fmt"

type PayloadTeamAdd struct {
	Team   *Team       `json:"team"`       // The team that was modified.
	User   *User       `json:"user"`       // The user that was added to this team.
	Repo   *Repository `json:"repository"` // The repository that was added to this team.
	Sender *User       `json:"sender"`     // Sender
}

func (s PayloadTeamAdd) String() string {
	return fmt.Sprintf("[%s] %s added team %s", s.Repo, s.Sender, s.Team)
}
