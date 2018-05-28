package patreon

// Address represents a Patreon's address.
type Address struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		Addressee   string `json:"addressee"`
		City        string `json:"city"`
		Country     string `json:"country"`
		Line1       string `json:"line_1"`
		Line2       string `json:"line_2"`
		PhoneNumber string `json:"phone_number"`
		PostalCode  string `json:"postal_code"`
		State       string `json:"state"`
	} `json:"attributes"`
}
