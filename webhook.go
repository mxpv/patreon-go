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

type WebhookPledge struct {
	Data Pledge `json:"data"`
}

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
