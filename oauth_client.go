package patreon

import "golang.org/x/oauth2"

// OauthClientAttributes is all fields in the OauthClient Attributes struct
var OauthClientAttributes = []string{
	"AuthorName", "ClientSecret", "DefaultScopes",
	"Description", "Domain", "IconURL", "Name",
	"PrivacyPolicyURL", "RedirectURIs", "TosURL", "Version",
}

// OauthClient is a client created by a developer, used for getting OAuth2 access tokens.
type OauthClient struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		AuthorName       string `json:"author_name"`
		ClientSecret     string `json:"client_secret"`
		DefaultScopes    string `json:"default_scopes"`
		Description      string `json:"description"`
		Domain           string `json:"domain"`
		IconURL          string `json:"icon_url"`
		Name             string `json:"name"`
		PrivacyPolicyURL string `json:"privacy_policy_url"`
		RedirectURIs     string `json:"redirect_uris"`
		TosURL           string `json:"tos_url"`
		Version          string `json:"version"`
	} `json:"attributes"`
	Relationships struct {
		// Apps *AppsRelationship  `json:"apps"`
		Campaign     *CampaignRelationship `json:"campaign,omitempty"`
		CreatorToken *oauth2.Token         `json:"creator_token,omitempty"`
		User         *UserRelationship     `json:"user,omitempty"`
	} `json:"relationships"`
}
