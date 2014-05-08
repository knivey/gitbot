package gitbot

type PayloadIssueComment struct {
	Action  string        `json:"action"`     // The action that was performed on the comment. Currently, can only be “created”.
	Issue   *Issue        `json:"issue"`      // The issue the comment belongs to.
	Comment *IssueComment `json:"comment"`    // The comment itself
	Repo    *Repository   `json:"repository"` // Repo
	Sender  *User         `json:"sender"`     // Sender
}
