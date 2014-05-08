package gitbot

type PayloadFork struct {
	Forkee *Repository `json:"forkee"`     // The created repository.
	Repo   *Repository `json:"repository"` // Repo
	Sender *User       `json:"sender"`     // Sender
}
