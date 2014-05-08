package gitbot

type PayloadCommitComment struct {
	Comment *CommitComment `json:"comment"`    // The comment itself.
	Repo    *Repository    `json:"repository"` // Repo
	Sender  *User          `json:"sender"`     // Sender
}
