// gitbot project gitbot.go
package gitbot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aarondl/cinotify"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	Name       = "gitbot"
	COMMIT_MSG = "[%s/%s] %s"
	NOMSG      = "No message"
	UNKNOWN    = "Unknown"
)

func setupLogLevel() {
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	// LOG_LEVEL not set, let's default to debug
	if !ok {
		lvl = "debug"
	}
	// parse string, this is built-in feature of logrus
	ll, err := logrus.ParseLevel(lvl)
	if err != nil {
		ll = logrus.DebugLevel
	}
	// set global log level
	logrus.SetLevel(ll)
}

func init() {
	setupLogLevel()
	cinotify.Register(Name, gitHandler{})
}

// gitHandler implements cinotify.Handler
type gitHandler struct {
}

func (_ gitHandler) Handle(r *http.Request) fmt.Stringer {
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)

	var payload interface{}

	switch r.Header.Get("X-GITHUB-EVENT") {
	case "create": // Branch creation
		payload = &PayloadCreate{}
	case "push": //Triggered when a repository branch is pushed to.
		payload = &PushPayload{}
	case "commit_comment": // Triggered when a commit comment is created.
		payload = &PayloadCommitComment{}
	case "delete": //Represents a deleted branch or tag.
		payload = &PayloadDelete{}
	case "fork": //Triggered when a user forks a repository.
		payload = &PayloadFork{}
	case "gollum": //Triggered when a Wiki page is created or updated.
		payload = &PayloadGollum{}
	case "issue_comment": //Triggered when an issue comment is created.
		payload = &PayloadIssueComment{}
	case "issues": //Triggered when an issue is created, closed or reopened.
		payload = &PayloadIssues{}
	case "member": //Triggered when a user is added as a collaborator to a repository.
		payload = &PayloadMember{}
	case "pull_request": //Triggered when a pull request is created, closed, reopened or synchronized.
		payload = &PayloadPullRequest{}
	case "pull_request_review_comment": //Triggered when a comment is created on a portion of the unified diff of a pull request.

	case "release": //Triggered when a release is published.
		payload = &PayloadRelease{}
	case "team_add": //Triggered when a user is added to a team or when a repository is added to a team.
		payload = &PayloadTeamAdd{}
	case "watch": //The WatchEvent is related to starring a repository, not watching.
		payload = &PayloadWatch{}
	default:
		return nil
	}

	if err := dec.Decode(payload); err != nil {
		logrus.Debugf("cinotify/gitbot/%s: Failed to decode json payload: %v",
			r.Header.Get("X-GITHUB-EVENT"), err)
		cinotify.DoLogf("cinotify/gitbot/%s: Failed to decode json payload: %v",
			r.Header.Get("X-GITHUB-EVENT"), err)

		return nil
	}

	if payload != nil {
		if stringer, ok := payload.(fmt.Stringer); ok {
			return stringer
		}
	}
	return nil
}

func (_ gitHandler) Route(r *mux.Route) {
	r.Path("/").Methods("POST").Headers(
		"Content-Type", "application/json",
		"X-GITHUB-EVENT", "",
	)
}
