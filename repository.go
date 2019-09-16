package gitbot

// Repository is the general repository object in the github API
type Repository struct {
	Owner           *User  `json:"owner"`
	ID              int    `json:"id"`
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	Description     string `json:"description"`
	Private         bool   `json:"private"`
	Fork            bool   `json:"fork"`
	URL             string `json:"url"`
	HTMLURL         string `json:"html_url"`
	CloneURL        string `json:"clone_url"`
	GitURL          string `json:"git_url"`
	SSHURL          string `json:"ssh_url"`
	SVNURL          string `json:"svn_url"`
	MirrorURL       string `json:"mirror_url"`
	Homepage        string `json:"homepage"`
	Language        string `json:"language"`
	ForksCount      int    `json:"forks_count"`
	StargazersCount int    `json:"stargazers_count"`
	WatchersCount   int    `json:"watchers_count"`
	Size            int    `json:"size"`
	DefaultBranch   string `json:"default_branch"`
	OpenIssuesCount int    `json:"open_issues_count"`
	HasIssues       bool   `json:"has_issues"`
	HasWiki         bool   `json:"has_wiki"`
	HasDownloads    bool   `json:"has_downloads"`
	//PushedAt        NullTime `json:"pushed_at"`
	//CreatedAt       NullTime `json:"created_at"`
	//UpdatedAt       NullTime `json:"updated_at"`
}

func (s Repository) String() string {
	return s.FullName
}
