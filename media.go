package patreon

// MediaAttributes is all fields in the Media Attributes struct
var MediaAttributes = []string{
	"CreatedAt", "DownloadURL", "FileName", "ImageURLs",
	"Metadata", "Mimetype", "OwnerID", "OwnerRelationship",
	"OwnerType", "SizeBytes", "State", "UploadExpiresAt",
	"UploadParameters", "UploadUrl",
}

// Media is a file uploaded to patreon.com, usually an image.
type Media struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		CreatedAt         NullTime    `json:"created_at"`
		DownloadURL       string      `json:"download_url"`
		FileName          string      `json:"file_name"`
		ImageURLs         interface{} `json:"image_urls"`
		Metadata          interface{} `json:"metadata"`
		Mimetype          string      `json:"mimetype"`
		OwnerID           string      `json:"owner_id"`
		OwnerRelationship string      `json:"owner_relationship"`
		OwnerType         string      `json:"owner_type"`
		SizeBytes         int         `json:"size_bytes"`
		State             string      `json:"state"`
		UploadExpiresAt   NullTime    `json:"upload_expires_at"`
		UploadParameters  interface{} `json:"upload_parameters"`
		UploadUrl         string      `json:"upload_url"`
	} `json:"attributes"`
}
