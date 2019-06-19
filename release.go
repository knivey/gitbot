package gitbot

import (
	"fmt"
)

type Release struct {
	Url             string `json:"url"`
	HtmlUrl         string `json:"html_url"`
	AssetsUrl       string `json:"assets_url"`
	UploadUrl       string `json:"upload_url"`
	TarballUrl      string `json:"tarball_url"`
	ZipballUrl      string `json:"zipball_url"` // lol zipball
	Id              int    `json:"id"`
	TagName         string `json:"tag_name"`
	TargetCommitish string `json:"target_commitish"`
	Name            string `json:"name"`
	Body            string `json:"body"`
	Draft           bool   `json:"draft"`
	PreRelease      bool   `json:"prerelease"`
	//CreatedAt       NullTime `json:"created_at"`
	//PublishedAt     NullTime `json:"published_at"`
	Author *User    `json:"author"`
	Assets []*Asset `json:"assets"`
}

func (s Release) String() string {
	return fmt.Sprintf("%s %s", s.Name, s.HtmlUrl)
}
