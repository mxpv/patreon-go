package patreon

// PostAttributes is all fields in the Post Attributes struct
var PostAttributes = []string{
	"AppID", "AppStatus", "Content", "EmbedData",
	"EmbedURL", "IsPaid", "IsPublic", "PublishedAt",
	"Title", "URL",
}

// Post is content posted by a creator on a campaign page.
type Post struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		AppID       int         `json:"app_id"`
		AppStatus   string      `json:"app_status"`
		Content     string      `json:"content"`
		EmbedData   interface{} `json:"embed_data"`
		EmbedURL    string      `json:"embed_url"`
		IsPaid      bool        `json:"is_paid"`
		IsPublic    bool        `json:"is_public"`
		PublishedAt NullTime    `json:"published_at"`
		Title       string      `json:"title"`
		URL         string      `json:""`
	} `json:"attributes"`
	Relationships struct {
		User     *UserRelationship     `json:"user,omitempty"`
		Campaign *CampaignRelationship `json:"campaign,omitempty"`
	} `json:"relationships"`
}

// PostResponse wraps Patreon's fetch benefit API response
type PostResponse struct {
	Data     Post     `json:"data"`
	Included Includes `json:"included"`
	Links    struct {
		Self string `json:"self"`
	} `json:"links"`
}

// PostsResponse wraps Patreon's fetch benefit API response
type PostsResponse struct {
	Data     []Post   `json:"data"`
	Included Includes `json:"included"`
	Links    struct {
		Self string `json:"self"`
	} `json:"links"`
}
