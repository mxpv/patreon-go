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

const (
	baseURL = "https://api.patreon.com"
)

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

	return &Client{httpClient: httpClient, baseURL: baseURL}
}

// Client returns the HTTP client configured for this client.
func (c *Client) Client() *http.Client {
	return c.httpClient
}

// FetchIdentity fetches a patron's profile info.
// This is the endpoint for accessing information about the current User with reference to the oauth token.
// It is most typically used in the OAuth "Log in with Patreon" flow to create or update the user's account on your site.
func (c *Client) FetchIdentity(opts ...requestOption) (*UserResponse, error) {
	resp := &UserResponse{}
	err := c.get("/api/oauth2/v2/identity", resp, opts...)
	return resp, err
}

// FetchCampaign is the single resource endpoint returns information about a single Campaign, fetched by campaign ID.
// Requires the campaigns scope.
func (c *Client) FetchCampaign(campaignID string, opts ...requestOption) (*CampaignResponse, error) {
	resp := &CampaignResponse{}
	err := c.get(fmt.Sprintf("/api/oauth2/v2/campaigns/%s", campaignID), resp, opts...)
	return resp, err
}

// FetchCampaigns Returns a list of Campaigns owned by the authorized user.
// Requires the campaigns scope.
func (c *Client) FetchCampaigns(opts ...requestOption) (*CampaignsResponse, error) {
	resp := &CampaignsResponse{}
	err := c.get("/api/oauth2/v2/campaigns", resp, opts...)
	return resp, err
}

// FetchCampaignMembers gets the Members for a given Campaign.
// Requires the campaigns.members scope.
func (c *Client) FetchCampaignMembers(campaignID string, opts ...requestOption) (*MembersResponse, error) {
	resp := &MembersResponse{}
	err := c.get(fmt.Sprintf("/api/oauth2/v2/campaigns/%s/members", campaignID), resp, opts...)
	return resp, err
}

// FetchCampaignMember gets a particular member by id.
// Requires the campaigns.members scope.
func (c *Client) FetchCampaignMember(memberID string, opts ...requestOption) (*MemberResponse, error) {
	resp := &MemberResponse{}
	err := c.get(fmt.Sprintf("/api/oauth2/v2/members/%s", memberID), resp, opts...)
	return resp, err
}

// FetchCampaignPosts gets a list of all the Posts on a given Campaign by campaign ID.
// Requires the campaigns.posts scope.
func (c *Client) FetchCampaignPosts(campaignID string, opts ...requestOption) (*PostsResponse, error) {
	resp := &PostsResponse{}
	err := c.get(fmt.Sprintf("/api/oauth2/v2/campaigns/%s/posts", campaignID), resp, opts...)
	return resp, err
}

// FetchCampaignPost gets a particular Post by ID.
// Requires the campaigns.posts scope.
func (c *Client) FetchCampaignPost(postID string, opts ...requestOption) (*PostResponse, error) {
	resp := &PostResponse{}
	err := c.get(fmt.Sprintf("/api/oauth2/v2/posts/%s", postID), resp, opts...)
	return resp, err
}

func (c *Client) buildURL(path string, opts ...requestOption) (string, error) {
	cfg := getOptions(opts...)

	u, err := url.ParseRequestURI(c.baseURL + path)
	if err != nil {
		return "", err
	}

	q := url.Values{}
	if cfg.include != "" {
		q.Set("include", cfg.include)
	}

	if len(cfg.fields) > 0 {
		for resource, fields := range cfg.fields {
			key := fmt.Sprintf("fields[%s]", resource)
			q.Set(key, fields)
		}
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

func (c *Client) get(path string, v interface{}, opts ...requestOption) error {
	addr, err := c.buildURL(path, opts...)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Get(addr)
	if err != nil {
		return err
	}

	// body, err := ioutil.ReadAll(resp.Body)

	// if err != nil {
	// 	return err
	// }

	// fmt.Println(string(body))

	if resp.StatusCode != http.StatusOK {
		errs := ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(&errs); err != nil {
			return err
		}

		return errs
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
