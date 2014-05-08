//Structures for git
package gitbot

import "time"

type PushCommit struct {
	Added     []string
	Modified  []string
	Removed   []string
	Author    *PushPerson
	Committer *PushPerson
	Distinct  bool
	Id        string
	Message   string
	Timestamp time.Time
	Url       string
}

type PushPerson struct {
	Email    string
	Name     string
	Username string
}

type PushRepository struct {
	Created_at    uint64
	Description   string
	Fork          bool
	Forks         uint
	Has_downloads bool
	Has_issues    bool
	Has_wiki      bool
	Homepage      string
	Id            uint
	Language      string
	Master_branch string
	Name          string
	Open_issues   uint
	Owner         *PushPerson
	Private       bool
	Pushed_at     uint64
	Size          uint
	Stargazers    uint
	Url           string
	Watchers      uint
}

type PushPayload struct {
	Before      string
	After       string
	Ref         string
	Compare     string
	Created     bool
	Deleted     bool
	Forced      bool
	Commits     []*PushCommit
	Head_commit *PushCommit
	Pusher      *PushPerson //may not have Username, beware
	Repository  *PushRepository
}
