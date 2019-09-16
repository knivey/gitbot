package gitbot

// PushCommit is a commit in a push
type PushCommit struct {
	Added     []string
	Modified  []string
	Removed   []string
	Author    *PushPerson
	Committer *PushPerson
	Distinct  bool
	ID        string
	Message   string
	//Timestamp *NullTime
	URL string
}

// PushPerson is the person who pushed
type PushPerson struct {
	Email    string
	Name     string
	Username string
}

// PushPayload comes from a push to a branch
type PushPayload struct {
	Before     string
	After      string
	Ref        string
	Compare    string
	Created    bool
	Deleted    bool
	Forced     bool
	Commits    []*PushCommit
	HeadCommit *PushCommit
	Pusher     *PushPerson //may not have Username, beware
	Repository *Repository
}
