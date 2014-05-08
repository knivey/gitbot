package gitbot

type PayloadTeamAdd struct {
	Team   *Team       `json:"team"`       // The team that was modified.
	User   *User       `json:"user"`       // The user that was added to this team.
	Repo   *Repository `json:"repository"` // The repository that was added to this team.
	Sender *User       `json:"sender"`     // Sender
}
