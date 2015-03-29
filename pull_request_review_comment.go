package gitbot

// TODO make PullRequestComment struct...
type PullRequestReviewComment struct {
	Comment *IssueComment `json:"comment"`    // The comment itself
	Repo    *Repository   `json:"repository"` // Repo
	Sender  *User         `json:"sender"`     // Sender
}
