package patreon

import "time"

// User represents a user
type User struct {
	Type       string `json:"type"`
	Id         string `json:"id"`
	Attributes struct {
		FirstName       string    `json:"first_name"`
		LastName        string    `json:"last_name"`
		FullName        string    `json:"full_name"`
		Vanity          string    `json:"vanity"`
		Email           string    `json:"email"`
		About           string    `json:"about"`
		FacebookId      string    `json:"facebook_id"`
		Gender          int       `json:"gender"`
		HasPassword     bool      `json:"has_password"`
		ImageURL        string    `json:"image_url"`
		ThumbURL        string    `json:"thumb_url"`
		YouTube         string    `json:"youtube"`
		Twitter         string    `json:"twitter"`
		Facebook        string    `json:"facebook"`
		IsEmailVerified bool      `json:"is_email_verified"`
		IsSuspended     bool      `json:"is_suspended"`
		IsDeleted       bool      `json:"is_deleted"`
		IsNuked         bool      `json:"is_nuked"`
		Created         time.Time `json:"created"`
		URL             string    `json:"url"`
	} `json:"attributes"`
}

type UserResponse struct {
	Data  User `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}
