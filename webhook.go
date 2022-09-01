package patreon

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
)

const (
	// EventCreatePledge specifies a create pledge event
	EventCreatePledge = "pledges:create"

	// EventUpdatePledge specifies an update pledge event
	EventUpdatePledge = "pledges:update"

	// EventDeletePledge specifies a delete pledge event
	EventDeletePledge = "pledges:delete"
)

const (
	// HeaderEventType specifies an event type HTTP header name
	HeaderEventType = "X-Patreon-Event"

	// HeaderEventSignature specifies message signature HTTP header name to verify message body
	HeaderSignature = "X-Patreon-Signature"
)

// VerifySignature verifies the sender of the message
func VerifySignature(message []byte, secret string, signature string) (bool, error) {
	hash := hmac.New(md5.New, []byte(secret))
	if _, err := hash.Write(message); err != nil {
		return false, err
	}

	sum := hash.Sum(nil)
	expectedSignature := hex.EncodeToString(sum)

	return expectedSignature == signature, nil
}

// WebhookAttributes is all fields in the Webhook Attributes struct
var WebhookAttributes = []string{
	"LastAttemptedAt", "NumConsecutiveTimesFailed", "Paused",
	"Secret", "Triggers", "URI",
}

// Webhook is fired based on events happening on a particular campaign.
type Webhook struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		LastAttemptedAt           NullTime    `json:"last_attempted_at"`
		NumConsecutiveTimesFailed int         `json:"num_consecutive_times_failed"`
		Paused                    bool        `json:"paused"`
		Secret                    string      `json:"secret"`
		Triggers                  interface{} `json:"triggers"`
		URI                       string      `json:"uri"`
	} `json:"attributes"`
	Relationships struct {
		Campaign    *CampaignRelationship    `json:"campaign,omitempty"`
		Memberships *MembershipsRelationship `json:"memberships,omitempty"`
	} `json:"relationships"`
}

// WebhookResponse wraps Patreon's fetch user API response
type WebhookResponse struct {
	Data     Webhook  `json:"data"`
	Included Includes `json:"included"`
	Links    struct {
		Self string `json:"self"`
	} `json:"links"`
}
