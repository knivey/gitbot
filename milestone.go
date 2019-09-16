package gitbot

// Milestone for a repository
type Milestone struct {
	Creator      *User  `json:"creator"`
	URL          string `json:"url"`
	Number       int    `json:"number"`
	State        string `json:"state"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	OpenIssues   int    `json:"open_issues"`
	ClosedIssues int    `json:"closed_issues"`
	//CreatedAt    NullTime `json:"created_at"`
	//UpdatedAt    NullTime `json:"updated_at"`
	DueOn string `json:"due_on"`
}
