package gitbot

type PayloadIssues struct {
	Action string      `json:"action"`     // The action that was performed. Can be one of “opened”, “closed”, or “reopened”.
	Issue  *Issue      `json:"issue"`      // The issue itself.
	Repo   *Repository `json:"repository"` // Repo
	Sender *User       `json:"sender"`     // Sender
}
