package gitbot

import (
	"net/http"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	payload := `
{
  "action": "closed",
  "issue": {
    "url": "https://api.github.com/repos/aarondl/ultimateq/issues/29",
    "labels_url": "https://api.github.com/repos/aarondl/ultimateq/issues/29/labels{/name}",
    "comments_url": "https://api.github.com/repos/aarondl/ultimateq/issues/29/comments",
    "events_url": "https://api.github.com/repos/aarondl/ultimateq/issues/29/events",
    "html_url": "https://github.com/aarondl/ultimateq/issues/29",
    "id": 32118394,
    "number": 29,
    "title": "Move all message splitting/repeating/filtering into inet.",
    "user": {
      "login": "aarondl",
      "id": 699706,
      "avatar_url": "https://avatars.githubusercontent.com/u/699706?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/aarondl",
      "html_url": "https://github.com/aarondl",
      "followers_url": "https://api.github.com/users/aarondl/followers",
      "following_url": "https://api.github.com/users/aarondl/following{/other_user}",
      "gists_url": "https://api.github.com/users/aarondl/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/aarondl/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/aarondl/subscriptions",
      "organizations_url": "https://api.github.com/users/aarondl/orgs",
      "repos_url": "https://api.github.com/users/aarondl/repos",
      "events_url": "https://api.github.com/users/aarondl/events{/privacy}",
      "received_events_url": "https://api.github.com/users/aarondl/received_events",
      "type": "User",
      "site_admin": false
    },
    "labels": [
      {
        "url": "https://api.github.com/repos/aarondl/ultimateq/labels/bug",
        "name": "bug",
        "color": "fc2929"
      },
      {
        "url": "https://api.github.com/repos/aarondl/ultimateq/labels/enhancement",
        "name": "enhancement",
        "color": "84b6eb"
      }
    ],
    "state": "closed",
    "locked": false,
    "assignee": null,
    "milestone": null,
    "comments": 0,
    "created_at": "2014-04-24T03:49:38Z",
    "updated_at": "2015-03-07T19:50:04Z",
    "closed_at": "2015-03-07T19:50:04Z",
    "body": "Currently irc.Helper is doing a lot of work that doesn't belong at that level. It should all be placed inside the inet.IrcClient's methods when it is dealing with written data.\r\n\r\nAs this issue is addressed, dissallowed characters in irc messages (\\r\\n\\0) must be filtered to prevent exploitation of the bot."
  },
  "repository": {
    "id": 9417576,
    "name": "ultimateq",
    "full_name": "aarondl/ultimateq",
    "owner": {
      "login": "aarondl",
      "id": 699706,
      "avatar_url": "https://avatars.githubusercontent.com/u/699706?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/aarondl",
      "html_url": "https://github.com/aarondl",
      "followers_url": "https://api.github.com/users/aarondl/followers",
      "following_url": "https://api.github.com/users/aarondl/following{/other_user}",
      "gists_url": "https://api.github.com/users/aarondl/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/aarondl/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/aarondl/subscriptions",
      "organizations_url": "https://api.github.com/users/aarondl/orgs",
      "repos_url": "https://api.github.com/users/aarondl/repos",
      "events_url": "https://api.github.com/users/aarondl/events{/privacy}",
      "received_events_url": "https://api.github.com/users/aarondl/received_events",
      "type": "User",
      "site_admin": false
    },
    "private": false,
    "html_url": "https://github.com/aarondl/ultimateq",
    "description": "An irc bot written in Go.",
    "fork": false,
    "url": "https://api.github.com/repos/aarondl/ultimateq",
    "forks_url": "https://api.github.com/repos/aarondl/ultimateq/forks",
    "keys_url": "https://api.github.com/repos/aarondl/ultimateq/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/aarondl/ultimateq/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/aarondl/ultimateq/teams",
    "hooks_url": "https://api.github.com/repos/aarondl/ultimateq/hooks",
    "issue_events_url": "https://api.github.com/repos/aarondl/ultimateq/issues/events{/number}",
    "events_url": "https://api.github.com/repos/aarondl/ultimateq/events",
    "assignees_url": "https://api.github.com/repos/aarondl/ultimateq/assignees{/user}",
    "branches_url": "https://api.github.com/repos/aarondl/ultimateq/branches{/branch}",
    "tags_url": "https://api.github.com/repos/aarondl/ultimateq/tags",
    "blobs_url": "https://api.github.com/repos/aarondl/ultimateq/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/aarondl/ultimateq/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/aarondl/ultimateq/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/aarondl/ultimateq/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/aarondl/ultimateq/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/aarondl/ultimateq/languages",
    "stargazers_url": "https://api.github.com/repos/aarondl/ultimateq/stargazers",
    "contributors_url": "https://api.github.com/repos/aarondl/ultimateq/contributors",
    "subscribers_url": "https://api.github.com/repos/aarondl/ultimateq/subscribers",
    "subscription_url": "https://api.github.com/repos/aarondl/ultimateq/subscription",
    "commits_url": "https://api.github.com/repos/aarondl/ultimateq/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/aarondl/ultimateq/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/aarondl/ultimateq/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/aarondl/ultimateq/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/aarondl/ultimateq/contents/{+path}",
    "compare_url": "https://api.github.com/repos/aarondl/ultimateq/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/aarondl/ultimateq/merges",
    "archive_url": "https://api.github.com/repos/aarondl/ultimateq/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/aarondl/ultimateq/downloads",
    "issues_url": "https://api.github.com/repos/aarondl/ultimateq/issues{/number}",
    "pulls_url": "https://api.github.com/repos/aarondl/ultimateq/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/aarondl/ultimateq/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/aarondl/ultimateq/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/aarondl/ultimateq/labels{/name}",
    "releases_url": "https://api.github.com/repos/aarondl/ultimateq/releases{/id}",
    "created_at": "2013-04-13T17:50:39Z",
    "updated_at": "2015-01-07T13:47:16Z",
    "pushed_at": "2015-01-23T06:02:50Z",
    "git_url": "git://github.com/aarondl/ultimateq.git",
    "ssh_url": "git@github.com:aarondl/ultimateq.git",
    "clone_url": "https://github.com/aarondl/ultimateq.git",
    "svn_url": "https://github.com/aarondl/ultimateq",
    "homepage": null,
    "size": 1819,
    "stargazers_count": 9,
    "watchers_count": 9,
    "language": "Go",
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": false,
    "forks_count": 4,
    "mirror_url": null,
    "open_issues_count": 5,
    "forks": 4,
    "open_issues": 5,
    "watchers": 9,
    "default_branch": "master"
  },
  "sender": {
    "login": "aarondl",
    "id": 699706,
    "avatar_url": "https://avatars.githubusercontent.com/u/699706?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/aarondl",
    "html_url": "https://github.com/aarondl",
    "followers_url": "https://api.github.com/users/aarondl/followers",
    "following_url": "https://api.github.com/users/aarondl/following{/other_user}",
    "gists_url": "https://api.github.com/users/aarondl/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/aarondl/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/aarondl/subscriptions",
    "organizations_url": "https://api.github.com/users/aarondl/orgs",
    "repos_url": "https://api.github.com/users/aarondl/repos",
    "events_url": "https://api.github.com/users/aarondl/events{/privacy}",
    "received_events_url": "https://api.github.com/users/aarondl/received_events",
    "type": "User",
    "site_admin": false
  }
}`

	module := gitHandler{}
	r, _ := http.NewRequest("GET", "/", strings.NewReader(payload))
	r.Header.Set("X-GITHUB-EVENT", "issues")
	ret := module.Handle(r)
	if ret == nil {
		t.Fatal("It should have generated a response")
	}

	if str := ret.String(); str != "[aarondl/ultimateq] aarondl closed issue #29 (Move all message splitting/repeating/filtering into inet.)" {
		t.Error("String was wrong:", str)
	}
}
