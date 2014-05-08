package gitbot

//Head or Base
type Head struct {
	User  *User       `json:"user"`
	Repo  *Repository `json:"repo"`
	Label string      `json:"label"`
	Ref   string      `json:"ref"`
	Sha   string      `json:"sha"`
}
