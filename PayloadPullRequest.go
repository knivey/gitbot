package gitbot

type PayloadPullRequest struct {
	Action      string       `json:"action"`       // The action that was performed. Can be one of “opened”, “closed”, “synchronize”, or “reopened”.
	Number      int          `json:"number"`       // The pull request number.
	PullRequest *PullRequest `json:"pull_request"` // The pull request itself.
	Repo        *Repository  `json:"repository"`   // Repo
	Sender      *User        `json:"sender"`       // Sender
}
