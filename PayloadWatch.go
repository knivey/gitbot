package gitbot

type PayloadWatch struct {
	Action string      `json:"action"`     // The action that was performed. Currently, can only be started.
	Repo   *Repository `json:"repository"` // Repo
	Sender *User       `json:"sender"`     // Sender
}
