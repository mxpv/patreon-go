package patreon

// BenefitAttributes is all fields in the Benefit Attributes struct
var BenefitAttributes = []string{
	"AppExternalID", "AppMeta", "BenefitType", "CreatedAt",
	"DeliverablesDueTodayCount", "DeliveredDeliverablesCount",
	"Description", "IsDeleted", "IsEnded", "IsPublished",
	"NextDeliverableDueDate", "NotDeliveredDeliverablesCount",
	"RuleType", "TiersCount", "Title",
}

// Benefit is a benefit added to the campaign, which can be added to a tier to be delivered to the patron.
type Benefit struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		AppExternalID                 string                 `json:"app_external_id"`
		AppMeta                       map[string]interface{} `json:"app_meta"`
		BenefitType                   string                 `json:"benefit_type"`
		CreatedAt                     NullTime               `json:"created_at"`
		DeliverablesDueTodayCount     int                    `json:"deliverables_due_today_count"`
		DeliveredDeliverablesCount    int                    `json:"delivered_deliverables_count"`
		Description                   string                 `json:"description"`
		IsDeleted                     bool                   `json:"is_deleted"`
		IsEnded                       bool                   `json:"is_ended"`
		IsPublished                   bool                   `json:"is_published"`
		NextDeliverableDueDate        NullTime               `json:"next_deliverable_due_date"`
		NotDeliveredDeliverablesCount int                    `json:"not_delivered_deliverables_count"`
		RuleType                      bool                   `json:"rule_type"`
		TiersCount                    int                    `json:"tiers_count"`
		Title                         string                 `json:"title"`
	} `json:"attributes"`
	Relationships struct {
		Campaign              *CampaignRelationship     `json:"campaign,omitempty"`
		CampaignInstallations interface{}               `json:"campaign_installation"` // I don't know what this is.. Couldn't find any docs / examples
		Deliverables          *DeliverablesRelationship `json:"deliverables,omitempty"`
		Tiers                 *TiersRelationship        `json:"tiers,omitempty"`
	} `json:"relationships"`
}
