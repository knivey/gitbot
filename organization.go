package gitbot

// Organization in github
type Organization struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	URL       string `json:"url"`
	AvatarURL string `json:"avatar_url"`
}
