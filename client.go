package patreon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const (
	// AuthorizationURL specifies Patreon's OAuth2 authorization endpoint (see https://tools.ietf.org/html/rfc6749#section-3.1).
	// See Example_refreshToken for examples.
	AuthorizationURL = "https://www.patreon.com/oauth2/authorize"

	// AccessTokenURL specifies Patreon's OAuth2 token endpoint (see https://tools.ietf.org/html/rfc6749#section-3.2).
	// See Example_refreshToken for examples.
	AccessTokenURL = "https://api.patreon.com/oauth2/token"
)

const BaseURL = "https://www.patreon.com"

// Client manages communication with Patreon API.
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient returns a new Patreon API client. If a nil httpClient is
// provided, http.DefaultClient will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{httpClient: httpClient, baseURL: BaseURL}
}

// Client returns the HTTP client configured for this client.
func (c *Client) Client() *http.Client {
	return c.httpClient
}

// GetIdentity fetches the User resource.
// Top-level includes: memberships, campaign.
// This is the endpoint for accessing information about the current User with reference to the oauth token.
// With the basic scope of identity, you will receive the user’s public profile information.
// If you have the identity[email] scope, you will also get the user’s email address.
// You will not receive email address without that scope.
// See https://docs.patreon.com/#get-api-oauth2-v2-identity
func (c *Client) GetIdentity(opts ...RequestOpt) (*User, error) {
	var resp = identityResponse{}
	if err := c.get("/api/oauth2/v2/identity", &resp, opts...); err != nil {
		return nil, err
	}

	user := User{
		ID:             resp.Data.ID,
		UserAttributes: resp.Data.Attributes,
	}

	if resp.Data.Relationships.Campaign.Data != nil {
		user.Campaign = resp.Included.campaigns[resp.Data.Relationships.Campaign.Data.ID]
	}

	for _, value := range resp.Included.memberships {
		user.Memberships = append(user.Memberships, value)
	}

	return &user, nil
}

// GetCampaigns returns a list of GetCampaigns owned by the authorized user.
// Requires the campaigns scope.
// Top-level includes: tiers, creator, benefits, goals.
// See https://docs.patreon.com/#get-api-oauth2-v2-campaigns
func (c *Client) GetCampaigns(opts ...RequestOpt) ([]*Campaign, error) {
	var resp campaignListResponse
	if err := c.get("/api/oauth2/v2/campaigns", &resp, opts...); err != nil {
		return nil, err
	}

	// Read 'data' array
	campaigns := make([]*Campaign, len(resp.Data))
	for idx, item := range resp.Data {
		campaign := &Campaign{
			ID: item.ID,
		}

		if item.Attributes != nil {
			campaign.CampaignAttributes = item.Attributes
		}

		// Read 'relationships' fields and link 'included' items

		if item.Relationships.Creator.Data != nil {
			campaign.Creator = resp.Included.users[item.Relationships.Creator.Data.ID]
		}

		for _, relation := range item.Relationships.Benefits.Data {
			campaign.Benefits = append(campaign.Benefits, resp.Included.benefits[relation.ID])
		}

		for _, relation := range item.Relationships.Goals.Data {
			campaign.Goals = append(campaign.Goals, resp.Included.goals[relation.ID])
		}

		for _, relation := range item.Relationships.Tiers.Data {
			campaign.Tiers = append(campaign.Tiers, resp.Included.tiers[relation.ID])
		}

		campaigns[idx] = campaign
	}

	return campaigns, nil
}

func (c *Client) GetCampaignByID(id string, opts ...RequestOpt) (*Campaign, error) {
	var resp campaignResponse
	if err := c.get("/api/oauth2/v2/campaigns/"+id, &resp, opts...); err != nil {
		return nil, err
	}

	campaign := &Campaign{
		ID: resp.Data.ID,
	}

	if resp.Data.Attributes != nil {
		campaign.CampaignAttributes = resp.Data.Attributes
	}

	if resp.Data.Relationships.Creator.Data != nil {
		campaign.Creator = resp.Included.users[resp.Data.Relationships.Creator.Data.ID]
	}

	for _, relation := range resp.Data.Relationships.Benefits.Data {
		campaign.Benefits = append(campaign.Benefits, resp.Included.benefits[relation.ID])
	}

	for _, relation := range resp.Data.Relationships.Goals.Data {
		campaign.Goals = append(campaign.Goals, resp.Included.goals[relation.ID])
	}

	for _, relation := range resp.Data.Relationships.Tiers.Data {
		campaign.Tiers = append(campaign.Tiers, resp.Included.tiers[relation.ID])
	}

	return campaign, nil
}

func (c *Client) buildURL(path string, opts ...RequestOpt) (string, error) {
	cfg := getOptions(opts...)

	u, err := url.ParseRequestURI(c.baseURL + path)
	if err != nil {
		return "", err
	}

	q := url.Values{}
	if cfg.include != "" {
		q.Set("include", cfg.include)
	}

	for resource, fields := range cfg.fields {
		key := fmt.Sprintf("fields[%s]", resource)
		q.Set(key, fields)
	}

	if cfg.size != 0 {
		q.Set("page[count]", strconv.Itoa(cfg.size))
	}

	if cfg.cursor != "" {
		q.Set("page[cursor]", cfg.cursor)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}

func (c *Client) get(path string, v interface{}, opts ...RequestOpt) error {
	addr, err := c.buildURL(path, opts...)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Get(addr)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		errs := ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(&errs); err != nil {
			return err
		}

		return errs
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
