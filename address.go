package patreon

// AddressAttributes is all fields in the Address Attributes struct
var AddressAttributes = []string{
	"Addressee", "City", "Country", "CreatedAt", "Line1",
	"Line2", "PhoneNumber", "PostalCode", "State",
}

// Address represents a Patreon's shipping address.
type Address struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		Addressee   string   `json:"addressee"`
		City        string   `json:"city"`
		Country     string   `json:"country"`
		CreatedAt   NullTime `json:"created_at"`
		Line1       string   `json:"line_1"`
		Line2       string   `json:"line_2"`
		PhoneNumber string   `json:"phone_number"`
		PostalCode  string   `json:"postal_code"`
		State       string   `json:"state"`
	} `json:"attributes"`
	Relationships struct {
		Campaigns *CampaignsRelationship `json:"campaigns,omitempty"`
		User      *UserRelationship      `json:"user,omitempty"`
	} `json:"relationships"`
}
