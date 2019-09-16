// Package gitbot parses git webhooks and sends them to cinotify
package gitbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aarondl/cinotify"
	"github.com/aarondl/gitio"
)

// Constants
const (
	Name      = "gitbot"
	CommitMsg = "[%s/%s] %s"
	NoMsg     = "No message"
	Unknown   = "Unknown"
)

func init() {
	cinotify.Register(Name, gitHandler{})
}

type shortener interface {
	Shorten(string) (string, error)
}
type shortenFunc func(string) (string, error)

func (s shortenFunc) Shorten(url string) (string, error) {
	return s(url)
}

var shortURL shortener = shortenFunc(gitio.Shorten)

// gitHandler implements cinotify.Handler
type gitHandler struct {
}

// Handle the event by returning a stringer, if nil the event was not parseable
// or otherwise understood
func (gitHandler) Handle(r *http.Request) fmt.Stringer {
	path := r.URL.Path == "/"
	method := r.Method == http.MethodPost
	contentType := r.Header.Get("Content-Type") == "application/json"
	xGithubEvent := len(r.Header.Get("X-GITHUB-EVENT")) > 0
	if !path || !method || !contentType || !xGithubEvent {
		return nil
	}

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
		return nil
	case "release": //Triggered when a release is published.
		payload = &PayloadRelease{}
	case "team_add": //Triggered when a user is added to a team or when a repository is added to a team.
		payload = &PayloadTeamAdd{}
	case "star":
		payload = &PayloadStar{}
	case "watch":
		// Watch doesn't make sense anymore, ignore it intentionally
		return nil
	default:
		cinotify.DoLogf("cinotify/gitbot/%s: unknown event", r.Header.Get("X-GITHUB-EVENT"))
		return nil
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil
	}

	if err := json.Unmarshal(b, payload); err != nil {
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
