package gitbot

import (
	"fmt"
	"strings"
)

const (
	fmtPushMsg = "[%s] %s pushed %v commits to %s %s"
)

// User that pushed
func (c PushCommit) User() string {
	if c.Committer == nil || c.Committer.Username == "" {
		return Unknown
	}
	return c.Committer.Username
}

// ShortID is the short git commit ID
func (c PushCommit) ShortID() string {
	if c.ID == "" {
		return Unknown
	}
	if len(c.ID) < 7 {
		return c.ID
	}
	return c.ID[:7]
}

// Msg is the last commit message from the pushed commits
func (c PushCommit) Msg() string {
	if c.Message == "" {
		return NoMsg
	}

	n := strings.Index(c.Message, "\n")

	if n > 0 {
		return c.Message[:n-1]
	}
	return c.Message
}

func (c PushCommit) String() string {
	return fmt.Sprintf("%s: %s: %s", c.ShortID(), c.User(), c.Msg())
}

// Branch of the push
func (pl PushPayload) Branch() string {
	refs := strings.Split(pl.Ref, "/")
	return refs[len(refs)-1]
}

// Name of the repo
func (pl PushPayload) Name() string {
	if pl.Repository == nil {
		return Unknown
	}
	return pl.Repository.String()
}

// NumCommits that occurred in the push
func (pl PushPayload) NumCommits() int {
	if pl.Commits == nil {
		return 0
	}
	return len(pl.Commits)
}

// PusherName is the person who pushed
func (pl PushPayload) PusherName() string {
	if pl.Pusher == nil || pl.Pusher.Name == "" {
		return Unknown
	}
	return pl.Pusher.Name
}

func (pl PushPayload) String() (out string) {
	comp, err := shortURL.Shorten(pl.Compare)
	if err != nil {
		comp = pl.Compare
	}
	out = fmt.Sprintf(fmtPushMsg, pl.Name(), pl.PusherName(),
		pl.NumCommits(), pl.Branch(), comp)
	return
}
