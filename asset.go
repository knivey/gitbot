package gitbot

import "time"

type Asset struct {
	Url           string    `json:"url"`
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	Label         string    `json:"label"`
	State         string    `json:"state"`
	ContentType   string    `json:"content_type"`
	Size          int       `json:"size"`
	DownloadCount int       `json:"download_count"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Uploader      *User     `json:"uploader"`
}
