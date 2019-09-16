package gitbot

import (
	"fmt"
)

// Release object
type Release struct {
	URL             string `json:"url"`
	HTMLURL         string `json:"html_url"`
	AssetsURL       string `json:"assets_url"`
	UploadURL       string `json:"upload_url"`
	TarballURL      string `json:"tarball_url"`
	ZipballURL      string `json:"zipball_url"` // lol zipball
	ID              int    `json:"id"`
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
	return fmt.Sprintf("%s %s", s.Name, s.HTMLURL)
}
