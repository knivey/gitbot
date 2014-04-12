// gitbot project gitbot.go
package gitbot

import (
	"encoding/json"
	"fmt"
	"github.com/aarondl/gitio"
	"github.com/aarondl/ultimateq/data"
	"log"
	"net/http"
	"strings"
)

var (
	Endpoint *data.DataEndpoint
	Channel  string
)

const (
	MSG1 = "[%s] %s pushed %v commits to %s %s"
	MSG2 = "[%s/%s] %s: %s: %s"
)

func (c Commit) User() string {
	if c.Committer == nil || c.Committer.Username == "" {
		return "\x02Unknown\x02"
	}
	return c.Committer.Username
}

func (c Commit) ShortId() string {
	if c.Id == "" {
		return "\x02Unknown\x02"
	}
	if len(c.Id) < 7 {
		return c.Id
	}
	return c.Id[:7]
}

func (c Commit) Msg() string {
	if c.Message == "" {
		return "No message"
	}
	return c.Message
}

func (pl Payload) Branch() string {
	refs := strings.Split(pl.Ref, "/")
	return refs[len(refs)-1]
}

func (pl Payload) Name() string {
	if pl.Repository == nil || pl.Repository.Name == "" {
		return "\x02Unknown\x02"
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
		return "\x02Unknown\x02"
	}
	return pl.Pusher.Name
}

func alertChan(pl Payload) {
	name := pl.Name()
	branch := pl.Branch()
	comp, err := gitio.Shorten(pl.Compare)
	if err != nil {
		comp = pl.Compare
	}

	msg := fmt.Sprintf(MSG1, name, pl.PusherName(), pl.NumCommits(), branch, comp)
	Endpoint.Privmsg(Channel, msg)

	var c *Commit
	for i := 0; i < pl.NumCommits(); i++ {
		c = pl.Commits[i]
		msg = fmt.Sprintf(MSG2, name, branch, c.ShortId(), c.User(), c.Msg())
		Endpoint.Privmsg(Channel, msg)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	payload := r.FormValue("payload")
	var Pl Payload
	err := json.Unmarshal([]byte(payload), &Pl)
	if err != nil {
		log.Println(err)
		return
	}
	alertChan(Pl)
}

//Start the http listener
func Listen(port, channel string, endpoint *data.DataEndpoint) (err error) {
	Channel = channel
	Endpoint = endpoint
	http.HandleFunc("/", handler)
	err = http.ListenAndServe(":"+port, nil)
	return
}
