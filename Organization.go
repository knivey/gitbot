package gitbot

type Organization struct {
	Login     string `json:"login"`
	Id        int    `json:"id"`
	Url       string `json:"url"`
	AvatarUrl string `json:"avatar_url"`
}
