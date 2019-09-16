package gitbot

// Team object from github api
type Team struct {
	URL          string        `json:"url"`
	Name         string        `json:"name"`
	ID           string        `json:"id"`
	Permission   string        `json:"permission"`
	MembersCount int           `json:"members_count"`
	ReposCount   int           `json:"repos_count"`
	Organization *Organization `json:"organization"`
}

func (s Team) String() string {
	return s.Name
}
