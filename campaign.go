package patreon_go

import "time"

type Campaign struct {
	Type       string `json:"type"`
	Id         string `json:"id"`
	Attributes struct {
		Summary                       string    `json:"summary"`
		Creation_name                 string    `json:"creation_name"`
		PayPerName                    string    `json:"pay_per_name"`
		OneLiner                      string    `json:"one_liner"`
		MainVideoEmbed                string    `json:"main_video_embed"`
		MainVideoURL                  string    `json:"main_video_url"`
		ImageSmallURL                 string    `json:"image_small_url"`
		ImageURL                      string    `json:"image_url"`
		ThanksVideoURL                string    `json:"thanks_video_url"`
		ThanksEmbed                   string    `json:"thanks_embed"`
		ThanksMsg                     string    `json:"thanks_msg"`
		IsMonthly                     bool      `json:"is_monthly"`
		IsNsfw                        bool      `json:"is_nsfw"`
		CreatedAt                     time.Time `json:"created_at"`
		Published_at                  time.Time `json:"published_at"`
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
