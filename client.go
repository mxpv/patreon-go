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

// Identity fetches the User resource.
// Top-level includes: memberships, campaign.
// This is the endpoint for accessing information about the current User with reference to the oauth token.
// With the basic scope of identity, you will receive the user’s public profile information.
// If you have the identity[email] scope, you will also get the user’s email address.
// You will not receive email address without that scope.
// See https://docs.patreon.com/#get-api-oauth2-v2-identity
func (c *Client) Identity(opts ...RequestOpt) (*User, error) {
	var resp = struct {
		Data struct {
			ID         string          `json:"id"`
			Attributes *UserAttributes `json:"attributes"`
		} `json:"data"`
		Included includedItems `json:"included"`
	}{}

	if err := c.get("/api/oauth2/v2/identity", &resp, opts...); err != nil {
		return nil, err
	}

	user := User{
		ID:             resp.Data.ID,
		UserAttributes: resp.Data.Attributes,
	}

	for _, item := range resp.Included.Items {
		switch item.Type {
		case "campaign":
			if campaign, err := item.toCampaign(); err != nil {
				return nil, err
			} else {
				user.Campaign = campaign
			}
		case "memberships":
			if member, err := item.toMember(); err != nil {
				return nil, err
			} else {
				user.Memberships = append(user.Memberships, member)
			}
		default:
			return nil, fmt.Errorf("unexpected include type %q", item.Type)
		}
	}

	return &user, nil
}

// Campaigns returns a list of Campaigns owned by the authorized user.
// Requires the campaigns scope.
// Top-level includes: tiers, creator, benefits, goals.
// See https://docs.patreon.com/#get-api-oauth2-v2-campaigns
func (c *Client) Campaigns(opts ...RequestOpt) ([]*Campaign, error) {
	var resp = struct {
		Data []struct {
			ID            string              `json:"id"`
			Attributes    *CampaignAttributes `json:"attributes"`
			Relationships struct {
				Benefits dataArray `json:"benefits"`
				Creator  data      `json:"creator"`
				Goals    dataArray `json:"goals"`
				Tiers    dataArray `json:"tiers"`
			} `json:"relationships"`
		} `json:"data"`
		Included includedItems `json:"included"`
	}{}

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

		if creator, err := resp.Included.findBy(item.Relationships.Creator.Data); err != nil {
			return nil, err
		} else {
			user, err := creator.toUser()
			if err != nil {
				return nil, err
			}

			campaign.Creator = user
		}

		for _, relation := range item.Relationships.Benefits.Data {
			data, err := resp.Included.findBy(relation)
			if err != nil {
				return nil, err
			}

			benefit, err := data.toBenefit()
			if err != nil {
				return nil, err
			}

			campaign.Benefits = append(campaign.Benefits, benefit)
		}

		for _, relation := range item.Relationships.Goals.Data {
			data, err := resp.Included.findBy(relation)
			if err != nil {
				return nil, err
			}

			goal, err := data.toGoal()
			if err != nil {
				return nil, err
			}

			campaign.Goals = append(campaign.Goals, goal)
		}

		for _, relation := range item.Relationships.Tiers.Data {
			data, err := resp.Included.findBy(relation)
			if err != nil {
				return nil, err
			}

			tier, err := data.toTier()
			if err != nil {
				return nil, err
			}

			campaign.Tiers = append(campaign.Tiers, tier)
		}

		campaigns[idx] = campaign
	}

	return campaigns, nil
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
