package patreon

import "time"

// Reward represents a Patreon's reward.
type Reward struct {
	Type       string `json:"type"`
	Id         string `json:"id"`
	Attributes struct {
		Amount           int       `json:"amount"`
		AmountCents      int       `json:"amount_cents"`
		CreatedAt        time.Time `json:"created_at"`
		DeletedAt        time.Time `json:"deleted_at"`
		EditedAt         time.Time `json:"edited_at"`
		Description      string    `json:"description"`
		ImageURL         string    `json:"image_url"`
		PatronCount      int       `json:"patron_count"`
		PostCount        int       `json:"post_count"`
		Published        bool      `json:"published"`
		PublishedAt      time.Time `json:"published_at"`
		RequiresShipping bool      `json:"requires_shipping"`
		Title            string    `json:"title"`
		UnpublishedAt    time.Time `json:"unpublished_at"`
		URL              string    `json:"url"`
	} `json:"attributes"`
}
