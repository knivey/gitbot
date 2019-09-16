package gitbot

import "fmt"

// PullRequest object from github API
type PullRequest struct {
	Head              *Head  `json:"head"`
	Base              *Head  `json:"base"`
	Links             *Links `json:"_links"`
	User              *User  `json:"user"`
	MergedBy          *User  `json:"merged_by"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	DiffURL           string `json:"diff_url"`
	PatchURL          string `json:"patch_url"`
	IssueURL          string `json:"issue_url"`
	CommitsURL        string `json:"commits_url"`
	ReviewCommentsURL string `json:"review_comments_url"`
	ReviewCommentURL  string `json:"review_comment_url"`
	CommentsURL       string `json:"comments_url"`
	StatusesURL       string `json:"statuses_url"`
	Number            int    `json:"number"`
	State             string `json:"state"`
	Title             string `json:"title"`
	Body              string `json:"body"`
	//CreatedAt         NullTime `json:"created_at"`
	//UpdatedAt         NullTime `json:"updated_at"`
	//ClosedAt          NullTime `json:"closed_at"`
	//MergedAt          NullTime `json:"merged_at"`
	MergeCommitSha string `json:"merge_commit_sha"`
	Merged         bool   `json:"merged"`
	Mergeable      bool   `json:"mergeable"`
	Comments       int    `json:"comments"`
	Commits        int    `json:"commits"`
	Additions      int    `json:"additions"`
	Deletions      int    `json:"deletions"`
	ChangedFiles   int    `json:"changed_files"`
}

func (s PullRequest) String() string {
	return fmt.Sprintf("#%v %s (%s)", s.Number, s.Title, s.HTMLURL)
}
