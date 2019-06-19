package gitbot

import (
	"fmt"
	"strings"

	"github.com/aarondl/gitio"
)

const (
	PUSH_MSG = "[%s] %s pushed %v commits to %s %s"
)

func (c PushCommit) User() string {
	if c.Committer == nil || c.Committer.Username == "" {
		return UNKNOWN
	}
	return c.Committer.Username
}

func (c PushCommit) ShortId() string {
	if c.Id == "" {
		return UNKNOWN
	}
	if len(c.Id) < 7 {
		return c.Id
	}
	return c.Id[:7]
}

func (c PushCommit) Msg() string {
	if c.Message == "" {
		return NOMSG
	}

	n := strings.Index(c.Message, "\n")

	if n > 0 {
		return c.Message[:n-1]
	}
	return c.Message
}

func (c PushCommit) String() string {
	return fmt.Sprintf("%s: %s: %s", c.ShortId(), c.User(), c.Msg())
}

func (pl PushPayload) Branch() string {
	refs := strings.Split(pl.Ref, "/")
	return refs[len(refs)-1]
}

func (pl PushPayload) Name() string {
	if pl.Repository == nil {
		return UNKNOWN
	}
	return pl.Repository.String()
}

func (pl PushPayload) NumCommits() int {
	if pl.Commits == nil {
		return 0
	}
	return len(pl.Commits)
}

func (pl PushPayload) PusherName() string {
	if pl.Pusher == nil || pl.Pusher.Name == "" {
		return UNKNOWN
	}
	return pl.Pusher.Name
}

func (pl PushPayload) String() (out string) {
	comp, err := gitio.Shorten(pl.Compare)
	if err != nil {
		comp = pl.Compare
	}
	out = fmt.Sprintf(PUSH_MSG, pl.Name(), pl.PusherName(),
		pl.NumCommits(), pl.Branch(), comp)
	return
}
