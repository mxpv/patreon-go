package patreon_go

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	AuthorizationURL = "https://www.patreon.com/oauth2/authorize"
	AccessTokenURL   = "https://api.patreon.com/oauth2/token"
	BaseURL          = "https://api.patreon.com"
)

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

// Fetch a patron's profile info
// This API returns a representation of the user who granted your OAuth client the provided access_token.
// It is most typically used in the OAuth "Log in with Patreon" flow to create or update the user's account on your site.
func (c *Client) FetchUser() (*UserResponse, error) {
	resp := &UserResponse{}
	err := c.get("/oauth2/api/current_user", resp)
	return resp, err
}

// Fetch your own profile and campaign info
// This API returns a representation of the user's campaign, including its rewards and goals, and the pledges to it.
// If there are more than twenty pledges to the campaign, the first twenty will be returned, along with a link to the
// next page of pledges.
func (c *Client) FetchCampaign() (*CampaignResponse, error) {
	resp := &CampaignResponse{}
	err := c.get("/oauth2/api/current_user/campaigns", resp)
	return resp, err
}

// Paging through a list of pledges to you
// This API returns a list of pledges to the provided campaignId. They are sorted by the date the pledge was made,
// and provide relationship references to the users who made each respective pledge. The API response will also contain
// a links section which may be used to fetch the next page of pledges, or go back to the first page.
func (c *Client) FetchPledges(campaignId string) (*PledgeResponse, error) {
	resp := &PledgeResponse{}
	path := fmt.Sprintf("/oauth2/api/campaigns/%s/pledges", campaignId)
	err := c.get(path, resp)
	return resp, err
}

func (c *Client) get(path string, v interface{}) error {
	resp, err := c.httpClient.Get(c.baseURL + path)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		errs := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(errs); err != nil {
			return err
		}

		return errs
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
