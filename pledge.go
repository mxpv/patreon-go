package patreon_go

import "time"

type Pledge struct {
	Type       string `json:"type"`
	Id         string `json:"id"`
	Attributes struct {
		AmountCents    int       `json:"amount_cents"`
		CreatedAt      time.Time `json:"created_at"`
		PledgeCapCents int       `json:"pledge_cap_cents"`
		PatronPaysFees bool      `json:"patron_pays_fees"`
	} `json:"attributes"`
}

type PledgeResponse struct {
	Data  []Pledge `json:"data"`
	Links struct {
		First string `json:"first"`
		Next  string `json:"next"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}
