package gitbot

type Team struct {
	Url          string        `json:"url"`
	Name         string        `json:"name"`
	Id           string        `json:"id"`
	Permission   string        `json:"permission"`
	MembersCount int           `json:"members_count"`
	ReposCount   int           `json:"repos_count"`
	Organization *Organization `json:"organization"`
}
