package gitbot

import "fmt"

type PullRequest struct {
	Head              *Head    `json:"head"`
	Base              *Head    `json:"base"`
	Links             *Links   `json:"_links"`
	User              *User    `json:"user"`
	MergedBy          *User    `json:"merged_by"`
	Url               string   `json:"url"`
	HtmlUrl           string   `json:"html_url"`
	DiffUrl           string   `json:"diff_url"`
	PatchUrl          string   `json:"patch_url"`
	IssueUrl          string   `json:"issue_url"`
	CommitsUrl        string   `json:"commits_url"`
	ReviewCommentsUrl string   `json:"review_comments_url"`
	ReviewCommentUrl  string   `json:"review_comment_url"`
	CommentsUrl       string   `json:"comments_url"`
	StatusesUrl       string   `json:"statuses_url"`
	Number            int      `json:"number"`
	State             string   `json:"state"`
	Title             string   `json:"title"`
	Body              string   `json:"body"`
	CreatedAt         NullTime `json:"created_at"`
	UpdatedAt         NullTime `json:"updated_at"`
	ClosedAt          NullTime `json:"closed_at"`
	MergedAt          NullTime `json:"merged_at"`
	MergeCommitSha    string   `json:"merge_commit_sha"`
	Merged            bool     `json:"merged"`
	Mergeable         bool     `json:"mergeable"`
	Comments          int      `json:"comments"`
	Commits           int      `json:"commits"`
	Additions         int      `json:"additions"`
	Deletions         int      `json:"deletions"`
	ChangedFiles      int      `json:"changed_files"`
}

func (s PullRequest) String() string {
	return fmt.Sprintf("#%v %s (%s)", s.Number, s.Title, s.HtmlUrl)
}
