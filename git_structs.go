//Structures for git
package gitbot

type Commit struct {
	Added     []string
	Modified  []string
	Removed   []string
	Author    *Person
	Committer *Person
	Distinct  bool
	Id        string
	Message   string
	Timestamp string //Might be able to use time here
	Url       string
}

type Person struct {
	Email    string
	Name     string
	Username string
}

type Repository struct {
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
	Owner         *Person
	Private       bool
	Pushed_at     uint64
	Size          uint
	Stargazers    uint
	Url           string
	Watchers      uint
}

type Payload struct {
	Before      string
	After       string
	Ref         string
	Compare     string
	Created     bool
	Deleted     bool
	Forced      bool
	Commits     []*Commit
	Head_commit *Commit
	Pusher      *Person //may not have Username, beware
	Repository  *Repository
}
