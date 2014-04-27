// gitbot project gitbot.go
package gitbot

import (
	"encoding/json"
	"fmt"
	"github.com/aarondl/gitio"
	"github.com/aarondl/cinotify"
	"github.com/aarondl/ultimateq/data"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

const (
	Name           = "gitbot"
	PAYLOAD_FORMAT = "[%s] %s pushed %v commits to %s %s"
	COMMIT_FORMAT  = "%s: %s: %s"
	MSG            = "[%s/%s] %s"
	// I have made this bold here to distinguish from actual
	// data that contains Unknown
	UNKNOWN        = "\x02Unknown\x02"
	NOMSG          = "No message"
)

func (c Commit) User() string {
	if c.Committer == nil || c.Committer.Username == "" {
		return UNKNOWN
	}
	return c.Committer.Username
}

func (c Commit) ShortId() string {
	if c.Id == "" {
		return UNKNOWN
	}
	if len(c.Id) < 7 {
		return c.Id
	}
	return c.Id[:7]
}

func (c Commit) Msg() string {
	if c.Message == "" {
		return NOMSG
	}
	n := strings.IndexOf(c.Message, "\n")
	if n > 0 {
		return c.Message[:n-1]
	} else {
		return c.Message
	}
}

func (c Commit) String() string {
	return fmt.Sprintf(COMMIT_FORMAT, c.ShortId(), c.User(), c.Msg())
}

func (pl Payload) Branch() string {
	refs := strings.Split(pl.Ref, "/")
	return refs[len(refs)-1]
}

func (pl Payload) Name() string {
	if pl.Repository == nil || pl.Repository.Name == "" {
		return UNKNOWN
	}
	return pl.Repository.Name
}

func (pl Payload) NumCommits() int {
	if pl.Commits == nil {
		return 0
	}
	return len(pl.Commits)
}

func (pl Payload) PusherName() string {
	if pl.Pusher == nil || pl.Pusher.Name == "" {
		return UNKNOWN
	}
	return pl.Pusher.Name
}

func (pl Payload) String() (out string) {
	comp, err := gitio.Shorten(pl.Compare)
	if err != nil {
		comp = pl.Compare
	}
	out = fmt.Sprintf(PAYLOAD_FORMAT, pl.Name(), pl.PusherName(),
		pl.NumCommits(), pl.Branch(), comp)
	return
}

func AlertChan(pl Payload, target string, endpoint *data.DataEndpoint) {
	var msg string
	var c *Commit
	name := pl.Name()
	branch := pl.Branch()

	endpoint.Privmsg(target, pl)

	for i := 0; i < pl.NumCommits(); i++ {
		c = pl.Commits[i]
		msg = fmt.Sprintf(MSG, name, branch, c)
		endpoint.Privmsg(target, msg)
	}
}

func pushHandler(r *http.Request) fmt.Stringer {
	payload := r.FormValue("payload")
	var Pl Payload
	err := json.Unmarshal([]byte(payload), &Pl)
	if err != nil {
		cinotify.DoLog("cinotify/gitbot: Failed to decode json payload: ", err)
		return nil
	}
	return Pl
}


func init() {
	cinotify.Register(Name, gitHandler{})
}

// gitHandler implements cinotify.Handler
type gitHandler struct {
}

func (_ gitHandler) Handle(r *http.Request) fmt.Stringer {
	defer r.Body.Close()
	
	switch r.Header.Get("X-GITHUB-EVENT") {
		case "push":
			return pushHandler(r)
	}

	return nil
}

func (_ gitHandler) Route(r *mux.Route) {
	r.Path("/").Methods("POST").Headers(
		"Content-Type", "application/x-www-form-urlencoded",
		"X-GITHUB-EVENT", "",
	)
}
