package gitbot

// Page refers to github pages
type Page struct {
	PageName string `json:"page_name"` // The name of the page.
	Title    string `json:"title"`     // The current page title.
	Action   string `json:"action"`    // The action that was performed on the page. Can be “created” or “edited”.
	SHA      string `json:"sha"`       // The latest commit SHA of the page.
	HTMLURL  string `json:"html_url"`  // Points to the HTML wiki page.
}
