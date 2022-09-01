package patreon

// UserDefaultIncludes specifies default includes for User.
const UserDefaultIncludes = "campaign,memberships"

// UserAttributes is all fields in the User Attributes struct
var UserAttributes = []string{
	"About", "CanSeeNSFW", "Created", "Email", "FirstName",
	"FullName", "HidePledges", "ImageURL", "IsEmailVerified",
	"LastName", "LikeCount", "SocialConnections", "ThumbURL",
	"URL", "Vanity",
}

// User is the Patreon user, which can be both patron and creator.
type User struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		About             string      `json:"about"`
		CanSeeNSFW        bool        `json:"can_see_nsfw"`
		Created           NullTime    `json:"created"`
		Email             string      `json:"email"`
		FirstName         string      `json:"first_name"`
		FullName          string      `json:"full_name"`
		HidePledges       bool        `json:"hide_pledges"`
		ImageURL          string      `json:"image_url"`
		IsEmailVerified   bool        `json:"is_email_verified"`
		LastName          string      `json:"last_name"`
		LikeCount         int         `json:"like_count"`
		SocialConnections interface{} `json:"social_connections"`
		ThumbURL          string      `json:"thumb_url"`
		URL               string      `json:"url"`
		Vanity            string      `json:"vanity"`
	} `json:"attributes"`
	Relationships struct {
		Campaign    *CampaignRelationship    `json:"campaign,omitempty"`
		Memberships *MembershipsRelationship `json:"memberships,omitempty"`
	} `json:"relationships"`
}

// UserResponse wraps Patreon's fetch user API response
type UserResponse struct {
	Data     User     `json:"data"`
	Included Includes `json:"included"`
	Links    struct {
		Self string `json:"self"`
	} `json:"links"`
}
