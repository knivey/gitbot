package gitbot

import (
	"encoding/json"
	"testing"
)

func TestPayloadCreate(t *testing.T) {
	payload := `
{
  "ref": "reaper",
  "ref_type": "branch",
  "master_branch": "master",
  "description": "An irc bot written in Go.",
  "pusher_type": "user",
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
    "issue_comment_url": "https://api.github.com/repos/aarondl/ultimateq/issues/comments/{number}",
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
    "size": 1995,
    "stargazers_count": 9,
    "watchers_count": 9,
    "language": "Go",
    "has_issues": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": false,
    "forks_count": 4,
    "mirror_url": null,
    "open_issues_count": 6,
    "forks": 4,
    "open_issues": 6,
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

	var p PayloadCreate
	if err := json.Unmarshal([]byte(payload), &p); err != nil {
		t.Error(err)
	}

	if str := p.String(); str != "[aarondl/ultimateq] aarondl created branch reaper" {
		t.Error("String is wrong:", str)
	}
}
