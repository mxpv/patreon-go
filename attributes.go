package patreon

// AddressAttributes represent patron's shipping address attributes
type AddressAttributes struct {
	// Full recipient name. Can be null.
	Addressee string `json:"addressee"`
	// First line of street address. Can be null.
	Line1 string `json:"line_1"`
	// Second line of street address. Can be null.
	Line2 string `json:"line_2"`
	// Postal or zip code. Can be null.
	PostalCode string `json:"postal_code"`
	// Patron's city
	City string `json:"city"`
	// State or province name. Can be null
	State string `json:"state"`
	// Patron's country
	Country string `json:"country"`
	// Telephone number. Specified for non-US addresses. Can be null.
	PhoneNumber string `json:"phone_number"`
	// Datetime address was first created.
	CreatedAt NullTime `json:"created_at"`
}

// BenefitAttributes represent a benefit attributes added to the campaign.
type BenefitAttributes struct {
	// Benefit display title
	Title string `json:"title"`
	// Benefit display description. Can be null.
	Description string `json:"description"`
	// Type of benefit, such as custom for creator-defined benefits. Can be null.
	BenefitType string `json:"benefit_type"`
	// A rule type designation, such as eom_monthly or one_time_immediate. Can be null.
	RuleType string `json:"rule_type"`
	// Datetime this benefit was created
	CreatedAt NullTime `json:"created_at"`
	// Number of deliverables for this benefit that have been marked complete
	DeliveredDeliverablesCount int `json:"delivered_deliverables_count"`
	// Number of deliverables for this benefit that are due, for all dates.
	NotDeliveredDeliverablesCount int `json:"not_delivered_deliverables_count"`
	// Number of deliverables for this benefit that are due today specifically.
	DeliverablesDueTodayCount int `json:"deliverables_due_today_count"`
	// The next due date (after EOD today) for this benefit. Can be null.
	NextDeliverableDueDate NullTime `json:"next_deliverable_due_date"`
	// Number of tiers containing this benefit.
	TiersCount int `json:"tiers_count"`
	// true if this benefit has been deleted.
	IsDeleted bool `json:"is_deleted"`
	// true if this benefit is ready to be shown to patrons.
	IsPublished bool `json:"is_published"`
	// The third-party external ID this reward is associated with, if any. Can be null.
	AppExternalId string `json:"app_external_id"`
	// Any metadata the third-party app included with this benefit on creation. Can be null.
	AppMeta interface{} `json:"app_meta"`
}

// CampaignAttributes represents the creator's page attributes.
type CampaignAttributes struct {
	// The creator's summary of their campaign. Can be null.
	Summary string `json:"summary"`
	// The type of content the creator is creating, as in "vanity is creating creation_name". Can be null.
	CreationName string `json:"creation_name"`
	// The thing which patrons are paying per, as in "vanity is making $1000 per pay_per_name". Can be null.
	PayPerName string `json:"pay_per_name"`
	// Pithy one-liner for this campaign, displayed on the creator page. Can be null.
	OneLiner string `json:"one_liner"`
	// Can be null.
	MainVideoEmbed string `json:"main_video_embed"`
	// Can be null.
	MainVideoURL string `json:"main_video_url"`
	// Banner image URL for the campaign.
	ImageURL string `json:"image_url"`
	// URL for the campaign's profile image.
	ImageSmallURL string `json:"image_small_url"`
	// URL for the video shown to patrons after they pledge to this campaign. Can be null.
	ThanksVideoURL string `json:"thanks_video_url"`
	// Can be null.
	ThanksEmbed string `json:"thanks_embed"`
	// Thank you message shown to patrons after they pledge to this campaign. Can be null.
	ThanksMsg string `json:"thanks_msg"`
	// true if the campaign charges per month, false if the campaign charges per-post.
	IsMonthly bool `json:"is_monthly"`
	// Whether this user has opted-in to rss feeds.
	HasRSS bool `json:"has_rss"`
	// Whether or not the creator has sent a one-time rss notification email.
	HasSentRSSNotify bool `json:"has_sent_rss_notify"`
	// The title of the campaigns rss feed.
	RSSFeedTitle string `json:"rss_feed_title"`
	// The url for the rss album artwork. Can be null.
	RSSArtworkURL string `json:"rss_artwork_url"`
	// true if the creator has marked the campaign as containing nsfw content.
	IsNSFW bool `json:"is_nsfw"`
	// true if the campaign charges upfront, false otherwise. Can be null.
	IsChargedImmediately bool `json:"is_charged_immediately"`
	// Datetime that the creator first began the campaign creation process.
	CreatedAt NullTime `json:"created_at"`
	// Datetime that the creator most recently published (made publicly visible) the campaign. Can be null.
	PublishedAt NullTime `json:"published_at"`
	// Relative (to patreon.com) URL for the pledge checkout flow for this campaign.
	PledgeURL string `json:"pledge_url"`
	// Number of patrons pledging to this creator.
	PatronCount int `json:"patron_count"`
	// The ID of the external discord server that is linked to this campaign. Can be null.
	DiscordServerID string `json:"discord_server_id"`
	// The ID of the Google Analytics tracker that the creator wants metrics to be sent to. Can be null.
	GoogleAnalyticsID string `json:"google_analytics_id"`
	// The visibility of the campaign's total earnings. One of private, public, patrons_only.
	EarningsVisibility string `json:"earnings_visibility"`
}

// DeliverableAttributes represents deliverable attributes.
type DeliverableAttributes struct {
	// When the creator marked the deliverable as completed or fulfilled to the patron. Can be null.
	CreatedAt NullTime `json:"created_at"`
	// One of delivered, not_delivered, wont_deliver.
	DeliveryStatus string `json:"delivery_status"`
	// When the deliverable is due to the patron.
	DueAt NullTime `json:"due_at"`
}

// GoalAttributes represents a funding goal in USD set by a creator on a campaign.
type GoalAttributes struct {
	// Goal amount in USD cents.
	AmountCents int `json:"amount_cents"`
	// Goal title.
	Title string `json:"title"`
	// Goal description. Can be null.
	Description string `json:"description"`
	// When the goal was created for the campaign.
	CreatedAt NullTime `json:"created_at"`
	// When the campaign reached the goal. Can be null.
	ReachedAt NullTime `json:"reached_at"`
	// Equal to (pledge_sum/goal amount)*100, helpful when a creator
	CompletedPercentage int `json:"completed_percentage"`
}

// MediaAttributes represents a file's attributes uploaded to patreon.com.
type MediaAttributes struct {
	// File name.
	FileName string `json:"file_name"`
	// Size of file in bytes.
	SizeBytes int `json:"size_bytes"`
	// Mimetype of uploaded file, eg: "application/jpeg".
	Mimetype string `json:"mimetype"`
	// Upload availability state of the file.
	State string `json:"state"`
	// Type of the resource that owns the file.
	OwnerType string `json:"owner_type"`
	// Ownership id (See also owner_type).
	OwnerID string `json:"owner_id"`
	// Ownership relationship type for multi-relationship medias.
	OwnerRelationship string `json:"owner_relationship"`
	// When the upload URL expires.
	UploadExpiresAt NullTime `json:"upload_expires_at"`
	// The URL to perform a POST request to in order to upload the media file.
	UploadURL string `json:"upload_url"`
	// All the parameters that have to be added to the upload form request.
	UploadParameters interface{} `json:"upload_parameters"`
	// The URL to download this media. Valid for 24 hours.
	DownloadURL string `json:"download_url"`
	// When the file was created.
	CreatedAt NullTime `json:"created_at"`
	// Metadata related to the file. Can be null.
	Metadata interface{} `json:"metadata"`
}

// MemberAttributes represents membership attributes.
type MemberAttributes struct {
	// One of active_patron, declined_patron, former_patron. Can be null.
	PatronStatus string `json:"patron_status"`
	// The user is not a pledging patron but has subscribed to updates about public posts.
	IsFollower bool `json:"is_follower"`
	// Full name of the member user.
	FullName string `json:"full_name"`
	// The member's email address. Requires the campaigns.members[email] scope.
	Email string `json:"email"`
	// Datetime of beginning of most recent pledge chainfrom this member to the campaign.
	// Pledge updates do not change this value. Can be null.
	PledgeRelationshipStart NullTime `json:"pledge_relationship_start"`
	// The total amount that the member has ever paid to the campaign. 0 if never paid.
	LifetimeSupportCents int `json:"lifetime_support_cents"`
	// The amount in cents that the member is entitled to.This includes a current pledge, or
	// payment that covers the current payment period.
	CurrentlyEntitledAmountCents int `json:"currently_entitled_amount_cents"`
	// Datetime of last attempted charge. null if never charged. Can be null.
	LastChargeDate NullTime `json:"last_charge_date"`
	// The result of the last attempted charge. The only successful status is Paid.
	// null if never charged.
	// One of Paid, Declined, Deleted, Pending, Refunded, Fraud, Other. Can be null.
	LastChargeStatus string `json:"last_charge_status"`
	// The creator's notes on the member.
	Note string `json:"note"`
	// The amount in cents the user will pay at the next pay cycle.
	WillPayAmountCents int `json:"will_pay_amount_cents"`
}

// OAuthClientAttributes represents a client's attributes.
type OAuthClientAttributes struct {
	// The client's secret.
	ClientSecret string `json:"client_secret"`
	// The name provided during client setup.
	Name string `json:"name"`
	// The description provided during client setup.
	Description string `json:"description"`
	// The author name provided during client setup. Can be null.
	AuthorName string `json:"author_name"`
	// The domain provided during client setup. Can be null.
	Domain string `json:"domain"`
	// The Patreon API version the client is targeting.
	Version int `json:"version"`
	// The URL of the icon used in the OAuth authorization flow. Can be null.
	IconURL string `json:"icon_url"`
	// The URL of the privacy policy provided during client setup. Can be null.
	PrivacyPolicyURL string `json:"privacy_policy_url"`
	// The URL of the terms of service provided during client setup. Can be null.
	TOSURL string `json:"tos_url"`
	// The allowable redirect URIs for the OAuth authorization flow.
	RedirectURIs string `json:"redirect_uris"`
	// The client's default OAuth scopes for the authorization flow (Deprecated in APIv2).
	DefaultScopes string `json:"default_scopes"`
}

// TierAttributes represents a membership level attributes.
type TierAttributes struct {
	// Monetary amount associated with this tier (in U.S. cents).
	AmountCents int `json:"amount_cents"`
	// Maximum number of patrons this tier is limited to, if applicable. Can be null.
	UserLimit int `json:"user_limit"`
	// Remaining number of patrons who may subscribe, if there is a user_limit. Can be null.
	Remaining int `json:"remaining"`
	// Tier display description.
	Description string `json:"description"`
	// true if this tier requires a shipping address from patrons.
	RequiresShipping bool `json:"requires_shipping"`
	// Datetime this tier was created.
	CreatedAt NullTime `json:"created_at"`
	// Fully qualified URL associated with this tier.
	URL string `json:"url"`
	// Number of patrons currently registered for this tier.
	PatronCount int `json:"patron_count"`
	// Number of posts published to this tier. Can be null.
	PostCount int `json:"post_count"`
	// The discord role IDs granted by this tier. Can be null.
	DiscordRoleIDs interface{} `json:"discord_role_ids"`
	// Tier display title.
	Title string `json:"title"`
	// Full qualified image URL associated with this tier. Can be null.
	ImageURL string `json:"image_url"`
	// Datetime tier was last modified.
	EditedAt NullTime `json:"edited_at"`
	// true if the tier is currently published.
	Published bool `json:"published"`
	// Datetime this tier was last published. Can be null.
	PublishedAt NullTime `json:"published_at"`
	// Datetime tier was unpublished, while applicable. Can be null.
	UnpublishedAt NullTime `json:"unpublished_at"`
}

// UserAttributes represent the Patreon user attributes.
type UserAttributes struct {
	// The user's email address. Requires certain scopes to access. See the scopes section of this documentation.
	Email string `json:"email"`
	// First name. Can be null.
	FirstName string `json:"first_name"`
	// Last name. Can be null.
	LastName string `json:"last_name"`
	// Combined first and last name.
	FullName string `json:"full_name"`
	// true if the user has confirmed their email.
	IsEmailVerified bool `json:"is_email_verified"`
	// The public "username" of the user. patreon.com/ goes to this user's creator page.
	// Non-creator users might not have a vanity. Can be null.
	Vanity string `json:"vanity"`
	// The user's about text, which appears on their profile. Can be null.
	About string `json:"about"`
	// The user's profile picture URL, scaled to width 400px.
	ImageURL string `json:"image_url"`
	// The user's profile picture URL, scaled to a square of size 100x100px.
	ThumbURL string `json:"thumb_url"`
	// true if this user can view nsfw content. Can be null.
	CanSeeNSFW bool `json:"can_see_nsfw"`
	// Datetime of this user's account creation.
	Created NullTime `json:"created"`
	// URL of this user's creator or patron profile.
	URL string `json:"url"`
	// How many posts this user has liked.
	LikeCount int `json:"like_count"`
	// true if the user has chosen to keep private which creators they pledge to. Can be null.
	HidePledges bool `json:"hide_pledges"`
	// Mapping from user's connected app names to external user id on the respective app.
	SocialConnections interface{} `json:"social_connections"`
}

// WebhookAttributes represent webhook attributes.
type WebhookAttributes struct {
	// List of events that will trigger this webhook.
	Triggers []string `json:"triggers"`
	// Fully qualified uri where webhook will be sent (e.g. https://www.example.com/webhooks/incoming).
	URI string `json:"uri"`
	// true if the webhook is paused as a result of repeated failed attempts to post to uri.
	// Set to false to attempt to re-enable a previously failing webhook.
	Paused bool `json:"paused"`
	// Last date that the webhook was attempted or used.
	LastAttemptedAt NullTime `json:"last_attempted_at"`
	// Number of times the webhook has failed consecutively, when in an error state.
	NumConsecutiveTimesFailed int `json:"num_consecutive_times_failed"`
	// Secret used to sign your webhook message body, so you can validate authenticity upon receipt.
	Secret string `json:"secret"`
}
