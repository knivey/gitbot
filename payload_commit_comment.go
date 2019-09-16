package gitbot

import "fmt"

// PayloadCommitComment is a commit comment
type PayloadCommitComment struct {
	Comment *CommitComment `json:"comment"`    // The comment itself.
	Repo    *Repository    `json:"repository"` // Repo
	Sender  *User          `json:"sender"`     // Sender
}

func (s PayloadCommitComment) String() string {
	return fmt.Sprintf("[%s] New CommitComment %s", s.Repo, s.Comment)
}
