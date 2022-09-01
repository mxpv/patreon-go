package patreon

// GoalAttributes is all fields in the Goal Attributes struct
var GoalAttributes = []string{
	"AmountCents", "CompletedPercentage", "CreatedAt",
	"Description", "ReachedAt", "Title",
}

// Goal is the funding goal in USD set by a creator on a campaign.
type Goal struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		AmountCents         int      `json:"amount_cents"`
		CompletedPercentage int      `json:"completed_percentage"`
		CreatedAt           NullTime `json:"created_at"`
		Description         string   `json:"description"`
		ReachedAt           NullTime `json:"reached_at"`
		Title               string   `json:"title"`
	} `json:"attributes"`
	Relationships struct {
		Campaign *CampaignRelationship `json:"campaign,omitempty"`
	} `json:"relationships"`
}
