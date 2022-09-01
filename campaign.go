package patreon

var (
	// CampaignDefaultIncludes specifies default includes for Campaign.
	CampaignDefaultIncludes = []string{"tiers", "creator", "benefits", "goals"}
	// CampaignAttributes is all fields in the Campaign Attributes struct
	CampaignAttributes = []string{
		"CreatedAt", "CreationName", "DiscordServerID", "GoogleAnalyticsID",
		"HasRSS", "HasSentRSSNotify", "ImageSmallURL", "ImageURL",
		"IsChargedImmediately", "IsMonthly", "IsNsfw", "MainVideoEmbed",
		"MainVideoURL", "OneLiner", "PatronCount", "PayPerName", "PledgeURL",
		"PublishedAt", "ResArtworkURL", "RSSFeedTitle", "ShowEarnings", "Summary",
		"ThanksEmbed", "ThanksMsg", "ThanksVideoURL", "URL", "Vanity",
	}
)

// Campaign is the creator's page, and the top-level object for accessing lists of members, tiers, etc.
type Campaign struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		CreatedAt            NullTime `json:"created_at"`
		CreationName         string   `json:"creation_name"`
		DiscordServerID      string   `json:"discord_server_id"`
		GoogleAnalyticsID    string   `json:"google_analytics_id"`
		HasRSS               bool     `json:"has_rss"`
		HasSentRSSNotify     bool     `json:"has_sent_rss_notify"`
		ImageSmallURL        string   `json:"image_small_url"`
		ImageURL             string   `json:"image_url"`
		IsChargedImmediately bool     `json:"is_charged_immediately"`
		IsMonthly            bool     `json:"is_monthly"`
		IsNsfw               bool     `json:"is_nsfw"`
		MainVideoEmbed       string   `json:"main_video_embed"`
		MainVideoURL         string   `json:"main_video_url"`
		OneLiner             string   `json:"one_liner"`
		PatronCount          int      `json:"patron_count"`
		PayPerName           string   `json:"pay_per_name"`
		PledgeURL            string   `json:"pledge_url"`
		PublishedAt          NullTime `json:"published_at"`
		ResArtworkURL        bool     `json:"res_artwork_url"`
		RSSFeedTitle         string   `json:"rss_feed_title"`
		ShowEarnings         bool     `json:"show_earnings"`
		Summary              string   `json:"summary"`
		ThanksEmbed          string   `json:"thanks_embed"`
		ThanksMsg            string   `json:"thanks_msg"`
		ThanksVideoURL       string   `json:"thanks_video_url"`
		URL                  string   `json:"url"`
		Vanity               string   `json:"vanity"`
	} `json:"attributes"`
	Relationships struct {
		Benefits              *BenefitsRelationship   `json:"benefits,omitempty"`
		CampaignInstallations *interface{}            `json:"campaign_installations,omitempty"`
		Categories            *CategoriesRelationship `json:"categories,omitempty"`
		Creator               *CreatorRelationship    `json:"creator,omitempty"`
		Goals                 *GoalsRelationship      `json:"goals,omitempty"`
		Tiers                 *TiersRelationship      `json:"tiers,omitempty"`
	} `json:"relationships"`
}

// CampaignV2Response wraps Patreon's campaign API response
type CampaignResponse struct {
	Data     Campaign `json:"data"`
	Included Includes `json:"included"`
}

// CampaignsV2Response wraps Patreon's campaign API response
type CampaignsResponse struct {
	Data     []Campaign `json:"data"`
	Included Includes   `json:"included"`
}
