package patreon

// CampaignDefaultRelations specifies default includes for Campaign.
const CampaignDefaultRelations = "rewards,creator,goals"

// Campaign represents Patreon's campaign.
// Valid relationships: rewards, creator, goals, pledges, current_user_pledge, post_aggregation, categories, preview_token.
type Campaign struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		Summary                       string   `json:"summary"`
		CreationName                  string   `json:"creation_name"`
		DisplayPatronGoals            bool     `json:"display_patron_goals"`
		PayPerName                    string   `json:"pay_per_name"`
		OneLiner                      string   `json:"one_liner"`
		MainVideoEmbed                string   `json:"main_video_embed"`
		MainVideoURL                  string   `json:"main_video_url"`
		ImageSmallURL                 string   `json:"image_small_url"`
		ImageURL                      string   `json:"image_url"`
		ThanksVideoURL                string   `json:"thanks_video_url"`
		ThanksEmbed                   string   `json:"thanks_embed"`
		ThanksMsg                     string   `json:"thanks_msg"`
		IsChargedImmediately          bool     `json:"is_charged_immediately"`
		IsMonthly                     bool     `json:"is_monthly"`
		IsNsfw                        bool     `json:"is_nsfw"`
		IsPlural                      bool     `json:"is_plural"`
		CreatedAt                     NullTime `json:"created_at"`
		PublishedAt                   NullTime `json:"published_at"`
		PledgeURL                     string   `json:"pledge_url"`
		PledgeSum                     int      `json:"pledge_sum"`
		PatronCount                   int      `json:"patron_count"`
		CreationCount                 int      `json:"creation_count"`
		OutstandingPaymentAmountCents int      `json:"outstanding_payment_amount_cents"`
	} `json:"attributes"`
	Relationships struct {
		Categories      *CategoriesRelationship      `json:"categories,omitempty"`
		Creator         *CreatorRelationship         `json:"creator,omitempty"`
		Rewards         *RewardsRelationship         `json:"rewards,omitempty"`
		Goals           *GoalsRelationship           `json:"goals,omitempty"`
		Pledges         *PledgesRelationship         `json:"pledges,omitempty"`
		PostAggregation *PostAggregationRelationship `json:"post_aggregation,omitempty"`
	} `json:"relationships"`
}

// CampaignResponse wraps Patreon's campaign API response
type CampaignResponse struct {
	Data     []Campaign `json:"data"`
	Included Includes   `json:"included"`
}
