package patreon

// Goal represents a Patreon's goal.
type Goal struct {
	Type       string `json:"type"`
	Id         string `json:"id"`
	Attributes struct {
		Amount              int      `json:"amount"`
		AmountCents         int      `json:"amount_cents"`
		CompletedPercentage int      `json:"completed_percentage"`
		CreatedAt           NullTime `json:"created_at"`
		ReachedAt           NullTime `json:"reached_at"`
		Title               string   `json:"title"`
		Description         string   `json:"description"`
	} `json:"attributes"`
}
