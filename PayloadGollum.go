package gitbot

type PayloadGollum struct {
	Pages  []*Page     `json:"pages"`      // The pages that were updated.
	Repo   *Repository `json:"repository"` // Repo
	Sender *User       `json:"sender"`     // Sender
}
