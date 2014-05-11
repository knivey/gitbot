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
	Name       = "gitbot"
	COMMIT_MSG = "[%s/%s] %s"
	NOMSG      = "No message"
	UNKNOWN    = "Unknown"
)

func AlertChan(pl PushPayload, target string, endpoint *data.DataEndpoint) {
	var msg string
	var c *PushCommit
	name := pl.Name()
	branch := pl.Branch()

	endpoint.Privmsg(target, pl)

	for i := 0; i < pl.NumCommits(); i++ {
		c = pl.Commits[i]
		msg = fmt.Sprintf(COMMIT_MSG, name, branch, c)
		endpoint.Privmsg(target, msg)
	}
}

func init() {
	cinotify.Register(Name, gitHandler{})
}

// gitHandler implements cinotify.Handler
type gitHandler struct {
}

func (_ gitHandler) Handle(r *http.Request) fmt.Stringer {
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	var err error

	switch r.Header.Get("X-GITHUB-EVENT") {
	//Triggered when a repository branch is pushed to.
	case "push":
		var pushPayload *PushPayload
		if err = dec.Decode(pushPayload); err == nil {
			return pushPayload
		}

	// Triggered when a commit comment is created.
	case "commit_comment":
		var pl *PayloadCommitComment
		if err = dec.Decode(pl); err == nil {
			return pl
		}

	//Represents a deleted branch or tag.
	case "delete":
		var pl *PayloadDelete
		if err = dec.Decode(pl); err == nil {
			return pl
		}

	//Triggered when a user forks a repository.
	case "fork":
		var pl *PayloadFork
		if err = dec.Decode(pl); err == nil {
			return pl
		}

	//Triggered when a Wiki page is created or updated.
	case "gollum":
		var pl *PayloadGollum
		if err = dec.Decode(pl); err == nil {
			return pl
		}

	//Triggered when an issue comment is created.
	case "issue_comment":
		var pl *PayloadIssueComment
		if err = dec.Decode(pl); err == nil {
			return pl
		}

	//Triggered when an issue is created, closed or reopened.
	case "issues":
		var pl *PayloadIssues
		if err = dec.Decode(pl); err == nil {
			return pl
		}

	//Triggered when a user is added as a collaborator to a repository.
	case "member":
		var pl *PayloadMember
		if err = dec.Decode(pl); err == nil {
			return pl
		}

	//Triggered when a pull request is created, closed, reopened or synchronized.
	case "pull_request":
		var pl *PayloadPullRequest
		if err = dec.Decode(pl); err == nil {
			return pl
		}

	//Triggered when a comment is created on a portion of the unified diff of a pull request.
	case "pull_request_review_comment":

	//Triggered when a release is published.
	case "release":
		var pl *PayloadRelease
		if err = dec.Decode(pl); err == nil {
			return pl
		}

	//Triggered when a user is added to a team or when a repository is added to a team.
	case "team_add":
		var pl *PayloadTeamAdd
		if err = dec.Decode(pl); err == nil {
			return pl
		}

	//The WatchEvent is related to starring a repository, not watching.
	case "watch":
		var pl *PayloadWatch
		if err = dec.Decode(pl); err == nil {
			return pl
		}
	}

	if err != nil {
		cinotify.DoLog("cinotify/gitbot/%s: Failed to decode json payload: %v",
			r.Header.Get("X-GITHUB-EVENT"), err)
	}
	return nil
}

func (_ gitHandler) Route(r *mux.Route) {
	r.Path("/").Methods("POST").Headers(
		"Content-Type", "application/json",
		"X-GITHUB-EVENT", "",
	)
}
