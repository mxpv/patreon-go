package patreon

import "time"

type Campaign struct {
	Type       string `json:"type"`
	Id         string `json:"id"`
	Attributes struct {
		Summary                       string    `json:"summary"`
		CreationName                  string    `json:"creation_name"`
		DisplayPatronGoals            bool      `json:"display_patron_goals"`
		PayPerName                    string    `json:"pay_per_name"`
		OneLiner                      string    `json:"one_liner"`
		MainVideoEmbed                string    `json:"main_video_embed"`
		MainVideoURL                  string    `json:"main_video_url"`
		ImageSmallURL                 string    `json:"image_small_url"`
		ImageURL                      string    `json:"image_url"`
		ThanksVideoURL                string    `json:"thanks_video_url"`
		ThanksEmbed                   string    `json:"thanks_embed"`
		ThanksMsg                     string    `json:"thanks_msg"`
		IsChargedImmediately          bool      `json:"is_charged_immediately"`
		IsMonthly                     bool      `json:"is_monthly"`
		IsNsfw                        bool      `json:"is_nsfw"`
		IsPlural                      bool      `json:"is_plural"`
		CreatedAt                     time.Time `json:"created_at"`
		PublishedAt                   time.Time `json:"published_at"`
		PledgeURL                     string    `json:"pledge_url"`
		PledgeSum                     int       `json:"pledge_sum"`
		PatronCount                   int       `json:"patron_count"`
		CreationCount                 int       `json:"creation_count"`
		OutstandingPaymentAmountCents int       `json:"outstanding_payment_amount_cents"`
	} `json:"attributes"`
}

type CampaignResponse struct {
	Data []Campaign `json:"data"`
}
