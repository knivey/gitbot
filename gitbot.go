// gitbot project gitbot.go
package gitbot

import (
	"encoding/json"
	"fmt"
	"github.com/aarondl/cinotify"
	"github.com/aarondl/ultimateq/data"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	Name           = "gitbot"
	PAYLOAD_FORMAT = "[%s] %s pushed %v commits to %s %s"
	COMMIT_FORMAT  = "%s: %s: %s"
	MSG            = "[%s/%s] %s"
	NOMSG          = "No message"
	UNKNOWN        = "Unknown"
)

func AlertChan(pl PushPayload, target string, endpoint *data.DataEndpoint) {
	var msg string
	var c *PushCommit
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
	var Pl PushPayload
	err := json.Unmarshal([]byte(payload), &Pl)
	if err != nil {
		cinotify.DoLog("cinotify/gitbot/push: Failed to decode json payload: ", err)
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
