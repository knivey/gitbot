//Structures for git
package gitbot

type PushCommit struct {
	Added     []string
	Modified  []string
	Removed   []string
	Author    *PushPerson
	Committer *PushPerson
	Distinct  bool
	Id        string
	Message   string
	//Timestamp *NullTime
	Url string
}

type PushPerson struct {
	Email    string
	Name     string
	Username string
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
	Repository  *Repository
}
