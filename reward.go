package patreon

// Reward represents a Patreon's reward.
type Reward struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		Amount           int      `json:"amount"`
		AmountCents      int      `json:"amount_cents"`
		CreatedAt        NullTime `json:"created_at"`
		DeletedAt        NullTime `json:"deleted_at"`
		EditedAt         NullTime `json:"edited_at"`
		Description      string   `json:"description"`
		ImageURL         string   `json:"image_url"`
		PatronCount      int      `json:"patron_count"`
		PostCount        int      `json:"post_count"`
		Published        bool     `json:"published"`
		PublishedAt      NullTime `json:"published_at"`
		RequiresShipping bool     `json:"requires_shipping"`
		Title            string   `json:"title"`
		UnpublishedAt    NullTime `json:"unpublished_at"`
		URL              string   `json:"url"`
	} `json:"attributes"`
}
