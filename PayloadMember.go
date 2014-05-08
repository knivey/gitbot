package gitbot

type PayloadMember struct {
	Member *User       `json:"member"`     // The user that was added.
	Action string      `json:"action"`     // The action that was performed. Currently, can only be “added”.
	Repo   *Repository `json:"repository"` // Repo
	Sender *User       `json:"sender"`     // Sender
}
