package patreon

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
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
		campaign := makeCampaign(&item, &resp.Included)
		campaigns[idx] = campaign
	}

	return campaigns, nil
}

// GetCampaignByID returns information about a single Campaign, fetched by campaign ID
// Requires the campaigns scope.
// Top-level includes: tiers, creator, benefits, goals.
// https://docs.patreon.com/#get-api-oauth2-v2-campaigns-campaign_id
func (c *Client) GetCampaignByID(id string, opts ...RequestOpt) (*Campaign, error) {
	if id == "" {
		return nil, errors.New("invalid campaign id")
	}

	var resp campaignResponse
	if err := c.get("/api/oauth2/v2/campaigns/"+id, &resp, opts...); err != nil {
		return nil, err
	}

	return makeCampaign(&resp.Data, &resp.Included), nil
}

func makeCampaign(data *campaignData, included *includes) *Campaign {
	campaign := &Campaign{
		ID: data.ID,
	}

	if data.Attributes != nil {
		campaign.CampaignAttributes = data.Attributes
	}

	relationships := &data.Relationships

	if relationships.Creator.Data != nil {
		campaign.Creator = included.users[relationships.Creator.Data.ID]
	}

	for _, relation := range relationships.Benefits.Data {
		campaign.Benefits = append(campaign.Benefits, included.benefits[relation.ID])
	}

	for _, relation := range relationships.Goals.Data {
		campaign.Goals = append(campaign.Goals, included.goals[relation.ID])
	}

	for _, relation := range relationships.Tiers.Data {
		campaign.Tiers = append(campaign.Tiers, included.tiers[relation.ID])
	}

	return campaign

}

// GetMembersByCampaignID gets the Members for a given Campaign by id.
// Requires the campaigns.members scope.
// Top-level includes: address (requires campaign.members.address scope), campaign, currently_entitled_tiers, user.
// We recommend using currently_entitled_tiers to see exactly what a Member is entitled to,
// either as an include on the members list or on the member get.
// See https://docs.patreon.com/#get-api-oauth2-v2-campaigns-campaign_id-members
func (c *Client) GetMembersByCampaignID(id string, opts ...RequestOpt) ([]*Member, error) {
	if id == "" {
		return nil, errors.New("invalid campaign id")
	}

	var resp memberListResponse
	path := fmt.Sprintf("/api/oauth2/v2/campaigns/%s/members", id)
	if err := c.get(path, &resp, opts...); err != nil {
		return nil, err
	}

	// Read 'data' array
	members := make([]*Member, len(resp.Data))

	for idx, item := range resp.Data {
		member := makeMember(&item, &resp.Included)
		members[idx] = member
	}

	return members, nil
}

// GetMemberByID gets a particular member by id.
// Requires the campaigns.members scope.
// Top-level includes: address (requires campaign.members.address scope), campaign, currently_entitled_tiers, user.
// We recommend using currently_entitled_tiers to see exactly what a member is entitled to,
// either as an include on the members list or on the member get.
// See https://docs.patreon.com/#get-api-oauth2-v2-members-id
func (c *Client) GetMemberByID(id string, opts ...RequestOpt) (*Member, error) {
	if id == "" {
		return nil, errors.New("invalid member id")
	}

	var resp memberResponse
	if err := c.get("/api/oauth2/v2/members/"+id, &resp, opts...); err != nil {
		return nil, err
	}

	return makeMember(&resp.Data, &resp.Included), nil
}

func makeMember(data *memberData, included *includes) *Member {
	member := &Member{
		ID: data.ID,
	}

	relationships := &data.Relationships

	if data.Attributes != nil {
		member.MemberAttributes = data.Attributes
	}

	if relationships.Address.Data != nil {
		member.Address = included.addresses[relationships.Address.Data.ID]
	}

	if relationships.Campaign.Data != nil {
		member.Campaign = included.campaigns[relationships.Campaign.Data.ID]
	}

	if relationships.User.Data != nil {
		member.User = included.users[relationships.User.Data.ID]
	}

	for _, item := range included.tiers {
		member.CurrentlyEntitledTiers = append(member.CurrentlyEntitledTiers, item)
	}

	return member
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
