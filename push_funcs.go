package gitbot

import (
	"fmt"
	"github.com/aarondl/gitio"
	"strings"
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
	} else {
		return c.Message
	}
}

func (c PushCommit) String() string {
	return fmt.Sprintf(COMMIT_FORMAT, c.ShortId(), c.User(), c.Msg())
}

func (pl PushPayload) Branch() string {
	refs := strings.Split(pl.Ref, "/")
	return refs[len(refs)-1]
}

func (pl PushPayload) Name() string {
	if pl.Repository == nil || pl.Repository.Name == "" {
		return UNKNOWN
	}
	return pl.Repository.Name
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
	out = fmt.Sprintf(PAYLOAD_FORMAT, pl.Name(), pl.PusherName(),
		pl.NumCommits(), pl.Branch(), comp)
	return
}
