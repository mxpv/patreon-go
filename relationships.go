package patreon

// Data represents a link to entity.
type Data struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// Meta represents extra information about relationship.
type Meta struct {
	Count int `json:"count"`
}

// CategoriesRelationship represents 'categories' include.
type CategoriesRelationship struct {
	Data []Data `json:"data"`
}

// CreatorRelationship represents 'creator' include.
type CreatorRelationship struct {
	Data  Data `json:"data"`
	Links struct {
		Related string `json:"related"`
	} `json:"links"`
}

// UserRelationship represents 'user' include
type UserRelationship struct {
	Data  Data `json:"data"`
	Links struct {
		Related string `json:"related"`
	} `json:"links"`
}

// PledgesRelationship represents 'pledges' include.
type PledgesRelationship struct {
	Data  []Data `json:"data"`
	Links struct {
		First string `json:"first"`
		Next  string `json:"next"`
	} `json:"links"`
	Meta Meta `json:"meta"`
}

// GoalsRelationship represents 'goals' include.
type GoalsRelationship struct {
	Data []Data `json:"data"`
}

// RewardsRelationship represents 'rewards' include.
type RewardsRelationship struct {
	Data []Data `json:"data"`
}

// RewardRelationship represents 'reward' include.
type RewardRelationship struct {
	Data  Data `json:"data"`
	Links struct {
		Related string `json:"related"`
	} `json:"links"`
}

// PostAggregationRelationship represents 'post_aggregation' include.
type PostAggregationRelationship struct {
	Data  Data `json:"data"`
	Links struct {
		Related string `json:"related"`
	} `json:"links"`
}

// CampaignRelationship represents 'campaign' include.
type CampaignRelationship struct {
	Data  Data `json:"data"`
	Links struct {
		Related string `json:"related"`
	} `json:"links"`
}

// PatronRelationship represents 'patron' include.
type PatronRelationship struct {
	Data  Data `json:"data"`
	Links struct {
		Related string `json:"related"`
	} `json:"links"`
}
