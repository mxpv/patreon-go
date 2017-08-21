package patreon

import "time"

// Card represents Patreon's credit card or paypal account.
type Card struct {
	Type       string `json:"type"`
	Id         string `json:"id"`
	Attributes struct {
		// PayPal
		CardType         string    `json:"card_type"`
		CreatedAt        time.Time `json:"created_at"`
		ExpirationDate   time.Time `json:"expiration_date"`
		HasFailedPayment bool      `json:"has_a_failed_payment"`
		IsVerified       bool      `json:"is_verified"`
		Number           string    `json:"number"`
		PaymentToken     string    `json:"payment_token"`
		PaymentTokenID   int       `json:"payment_token_id"`
	} `json:"attributes"`
	Relationships struct {
		User *UserRelationship `json:"user"`
	} `json:"relationships"`
}
