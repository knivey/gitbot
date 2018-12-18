package gitbot

import "time"

type Repository struct {
	Owner           *User     `json:"owner"`
	Id              int       `json:"id"`
	Name            string    `json:"name"`
	FullName        string    `json:"full_name"`
	Description     string    `json:"description"`
	Private         bool      `json:"private"`
	Fork            bool      `json:"fork"`
	Url             string    `json:"url"`
	HtmlUrl         string    `json:"html_url"`
	CloneUrl        string    `json:"clone_url"`
	GitUrl          string    `json:"git_url"`
	SshUrl          string    `json:"ssh_url"`
	SvnUrl          string    `json:"svn_url"`
	MirrorUrl       string    `json:"mirror_url"`
	Homepage        string    `json:"homepage"`
	Language        string    `json:"language"`
	ForksCount      int       `json:"forks_count"`
	StargazersCount int       `json:"stargazers_count"`
	WatchersCount   int       `json:"watchers_count"`
	Size            int       `json:"size"`
	DefaultBranch   string    `json:"default_branch"`
	OpenIssuesCount int       `json:"open_issues_count"`
	HasIssues       bool      `json:"has_issues"`
	HasWiki         bool      `json:"has_wiki"`
	HasDownloads    bool      `json:"has_downloads"`
	PushedAt        time.Time `json:"pushed_at"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (s Repository) String() string {
	return s.FullName
}
