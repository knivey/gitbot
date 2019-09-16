package gitbot

// Label for a repository
type Label struct {
	URL   string `json:"url"`
	Name  string `json:"name"`
	Color string `json:"color"`
}
