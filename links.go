package gitbot

type Links struct {
	Self           *Link `json:"self"`
	Html           *Link `json:"html"`
	Issue          *Link `json:"issue"`
	Comments       *Link `json:"comments"`
	ReviewComments *Link `json:"review_comments"`
	ReviewComment  *Link `json:"review_comment"`
	Commits        *Link `json:"commits"`
	Statuses       *Link `json:"statuses"`
}
