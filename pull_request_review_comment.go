package gitbot

// PullRequestReviewComment review for pull requests
type PullRequestReviewComment struct {
	Comment *IssueComment `json:"comment"`    // The comment itself
	Repo    *Repository   `json:"repository"` // Repo
	Sender  *User         `json:"sender"`     // Sender
}
