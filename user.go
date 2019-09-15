package patreon

// UserDefaultRelations specifies default includes for User.
const UserDefaultRelations = "campaign,pledges"

// User represents a Patreon's user.
// Valid relationships: pledges, cards, follows, campaign, presence, session, locations, current_user_follow, pledge_to_current_user.
type User struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		FirstName       string   `json:"first_name"`
		LastName        string   `json:"last_name"`
		FullName        string   `json:"full_name"`
		Vanity          string   `json:"vanity"`
		Email           string   `json:"email"`
		About           string   `json:"about"`
		FacebookId      string   `json:"facebook_id"`
		Gender          int      `json:"gender"`
		HasPassword     bool     `json:"has_password"`
		ImageURL        string   `json:"image_url"`
		ThumbURL        string   `json:"thumb_url"`
		YouTube         string   `json:"youtube"`
		Twitter         string   `json:"twitter"`
		Facebook        string   `json:"facebook"`
		IsEmailVerified bool     `json:"is_email_verified"`
		IsSuspended     bool     `json:"is_suspended"`
		IsDeleted       bool     `json:"is_deleted"`
		IsNuked         bool     `json:"is_nuked"`
		Created         NullTime `json:"created"`
		URL             string   `json:"url"`
		DiscordId	string 	 `json:"discord_id"`
	} `json:"attributes"`
	Relationships struct {
		Pledges *PledgesRelationship `json:"pledges,omitempty"`
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
