package patreon

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewClient(nil)
	client.baseURL = server.URL
}

func teardown() {
	server.Close()
}

func TestBuildURL(t *testing.T) {
	client := NewClient(nil)

	url, err := client.buildURL("/path",
		WithIncludes("patron", "reward", "creator"),
		WithFields("pledge", "total_historical_amount_cents", "unread_count"),
		WithPageSize(10),
		WithCursor("123"),
	)

	require.NoError(t, err)
	require.Equal(t, "https://www.patreon.com/path?fields%5Bpledge%5D=total_historical_amount_cents%2Cunread_count&include=patron%2Creward%2Ccreator&page%5Bcount%5D=10&page%5Bcursor%5D=123", url)
}

func TestBuildURLWithInvalidPath(t *testing.T) {
	client := &Client{}

	url, err := client.buildURL("")
	require.Error(t, err)
	require.Empty(t, url)
}

func TestOAuthClient(t *testing.T) {
	tc := oauth2.NewClient(context.Background(), nil)
	client := NewClient(tc)
	require.Equal(t, tc, client.Client())
}

const testIdentityResponse = `
{
    "data": {
        "attributes": {
            "about": "",
            "created": "2016-02-02T19:56:14+00:00",
            "email": "mail@gmail.com",
            "first_name": "Max",
            "full_name": "",
            "image_url": "https://c10.patreonusercontent.com/3/eyJ3IjoyMDB9/patreon-media/p/user/2822191/8c0bc222ed0c43e68c620fa191b0a0c5/1?token-time=2145916800&token-hash=efJ0dqJhGOR-FtExvoH3ELw8CPpHo5I3Ce6ixNaJmuk%3D",
            "last_name": "",
            "social_connections": {
                "deviantart": null,
                "discord": null,
                "facebook": null,
                "reddit": null,
                "spotify": null,
                "twitch": null,
                "twitter": null,
                "youtube": null
            },
            "thumb_url": "https://c10.patreonusercontent.com/3/eyJoIjoxMDAsInciOjEwMH0%3D/patreon-media/p/user/2822191/8c0bc222ed0c43e68c620fa191b0a0c5/1?token-time=2145916800&token-hash=NoMxBCUckp3EBrPgADzsEdUI3uFV13EB_wRx4LqIh4I%3D",
            "url": "https://www.patreon.com/podsync",
            "vanity": "podsync"
        },
        "id": "2822191",
        "relationships": {
            "campaign": {
                "data": {
                    "id": "278915",
                    "type": "campaign"
                },
                "links": {
                    "related": "https://www.patreon.com/api/oauth2/v2/campaigns/278915"
                }
            },
            "memberships": {
                "data": []
            }
        },
        "type": "user"
    },
    "included": [{
        "attributes": {
            "created_at": null,
            "creation_name": "new podcasting experience - Podsync",
            "discord_server_id": null,
            "google_analytics_id": null,
            "has_rss": false,
            "has_sent_rss_notify": false,
            "image_small_url": "https://c10.patreonusercontent.com/3/eyJoIjo2NDAsInciOjY0MH0%3D/patreon-media/p/campaign/278915/c17598520740471ca6c0ffe553ade639/1?token-time=2145916800&token-hash=Au6pSGBsM1mQ4D3YFFtbrJHit_G99uOvIyJs_C9uT7E%3D",
            "image_url": "https://c10.patreonusercontent.com/3/eyJ3IjoxOTIwfQ%3D%3D/patreon-media/p/campaign/278915/c17598520740471ca6c0ffe553ade639/1?token-time=2145916800&token-hash=thO-8NggSSPFYnLMeW9YowqCYKgsfTtwah_eoak6tVc%3D",
            "is_charged_immediately": true,
            "is_monthly": true,
            "is_nsfw": false,
            "main_video_embed": "",
            "main_video_url": "",
            "one_liner": null,
            "patron_count": 482,
            "pay_per_name": "month",
            "pledge_url": "/join/podsync",
            "published_at": "2016-02-02T20:11:19+00:00",
            "rss_artwork_url": null,
            "rss_feed_title": null,
            "summary": "<a href=\"http://podsync.net/\" rel=\"nofollow\">Podsync</a> - is a simple, free service that lets you listen to any YouTube / Vimeo channels, playlists or user videos in podcast format.<br><br><strong>Idea:</strong><br>Podcast applications have a rich functionality for content delivery - automatic download of new episodes, remembering last played position, sync between devices and offline listening. This functionality is not available on YouTube and Vimeo. So the aim of\u00a0<a href=\"http://podsync.net/\" rel=\"nofollow\">Podsync</a> is to make your life easier and enable you to view/listen to content on any device in podcast client.<br><br>It's my hobby project, so to continue to support and improve it, I need your help. Your money will go into paying my server bills and adding new features.<br><br>",
            "thanks_embed": "",
            "thanks_msg": "",
            "thanks_video_url": null
        },
        "id": "278915",
        "type": "campaign"
    }]
}
`

func TestClient_GetIdentity(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/oauth2/v2/identity", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, testIdentityResponse)
		assert.NoError(t, err)
	})

	user, err := client.GetIdentity()
	require.NoError(t, err)
	require.NotNil(t, user)

	assert.Equal(t, "2822191", user.ID)
	assert.Equal(t, "mail@gmail.com", user.Email)
	assert.Equal(t, "Max", user.FirstName)
	assert.Equal(t, "https://www.patreon.com/podsync", user.URL)
	assert.Equal(t, "podsync", user.Vanity)

	assert.Equal(t, time.Date(2016, 02, 02, 19, 56, 14, 0, time.UTC).Unix(), user.Created.Unix())

	require.NotNil(t, user.Campaign)

	assert.Equal(t, "278915", user.Campaign.ID)
	assert.Equal(t, "new podcasting experience - Podsync", user.Campaign.CreationName)
	assert.EqualValues(t, 482, user.Campaign.PatronCount)
	assert.Equal(t, "/join/podsync", user.Campaign.PledgeURL)
	assert.True(t, user.Campaign.IsMonthly)
	assert.True(t, user.Campaign.IsChargedImmediately)
}

const testIdentityResponseWithEmptyAttributes = `
{
    "data": {
        "attributes": {},
        "id": "2822191",
        "relationships": {
            "campaign": {
                "data": {
                    "id": "278915",
                    "type": "campaign"
                },
                "links": {
                    "related": "https://www.patreon.com/api/oauth2/v2/campaigns/278915"
                }
            },
            "memberships": {
                "data": []
            }
        },
        "type": "user"
    },
    "included": [{
        "attributes": {},
        "id": "278915",
        "type": "campaign"
    }]
}
`

func TestClient_GetIdentityEmptyAttributes(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/oauth2/v2/identity", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, testIdentityResponseWithEmptyAttributes)
		assert.NoError(t, err)
	})

	user, err := client.GetIdentity()
	require.NoError(t, err)
	require.NotNil(t, user)

	assert.Equal(t, "2822191", user.ID)
}

const testCampaignListResponse = `
{
    "data": [{
        "attributes": {
            "created_at": "2016-02-02T19:58:18+00:00",
            "creation_name": "new podcasting experience - Podsync",
            "discord_server_id": null,
            "google_analytics_id": null,
            "has_rss": false,
            "has_sent_rss_notify": false,
            "image_small_url": "https://c10.patreonusercontent.com/3/eyJoIjo2NDAsInciOjY0MH0%3D/patreon-media/p/campaign/278915/c17598520740471ca6c0ffe553ade639/1?token-time=2145916800&token-hash=Au6pSGBsM1mQ4D3YFFtbrJHit_G99uOvIyJs_C9uT7E%3D",
            "image_url": "https://c10.patreonusercontent.com/3/eyJ3IjoxOTIwfQ%3D%3D/patreon-media/p/campaign/278915/c17598520740471ca6c0ffe553ade639/1?token-time=2145916800&token-hash=thO-8NggSSPFYnLMeW9YowqCYKgsfTtwah_eoak6tVc%3D",
            "is_charged_immediately": true,
            "is_monthly": true,
            "is_nsfw": false,
            "main_video_embed": "",
            "main_video_url": "",
            "one_liner": null,
            "patron_count": 482,
            "pay_per_name": "month",
            "pledge_url": "/join/podsync",
            "published_at": "2016-02-02T20:11:19+00:00",
            "rss_artwork_url": null,
            "rss_feed_title": null,
            "summary": "<a href=\"http://podsync.net/\" rel=\"nofollow\">Podsync</a> - is a simple, free service that lets you listen to any YouTube / Vimeo channels, playlists or user videos in podcast format.<br><br><strong>Idea:</strong><br>Podcast applications have a rich functionality for content delivery - automatic download of new episodes, remembering last played position, sync between devices and offline listening. This functionality is not available on YouTube and Vimeo. So the aim of\u00a0<a href=\"http://podsync.net/\" rel=\"nofollow\">Podsync</a> is to make your life easier and enable you to view/listen to content on any device in podcast client.<br><br>It's my hobby project, so to continue to support and improve it, I need your help. Your money will go into paying my server bills and adding new features.<br><br>",
            "thanks_embed": "",
            "thanks_msg": "You are awesome! Thank you so much!<br>",
            "thanks_video_url": null
        },
        "id": "278915",
        "relationships": {
            "benefits": {
                "data": []
            },
            "creator": {
                "data": {
                    "id": "2822191",
                    "type": "user"
                }
            },
            "goals": {
                "data": [{
                    "id": "342492",
                    "type": "goal"
                }, {
                    "id": "605110",
                    "type": "goal"
                }, {
                    "id": "605111",
                    "type": "goal"
                }]
            },
            "tiers": {
                "data": [{
                    "id": "1048240",
                    "type": "tier"
                }, {
                    "id": "2140517",
                    "type": "tier"
                }]
            }
        },
        "type": "campaign"
    }],
    "included": [{
        "attributes": {
            "about": "",
            "created": "2016-02-02T19:56:14+00:00",
            "email": "mail@gmail.com",
            "first_name": "Max",
            "full_name": "Max",
            "image_url": "https://c10.patreonusercontent.com/3/eyJ3IjoyMDB9/patreon-media/p/user/2822191/8c0bc222ed0c43e68c620fa191b0a0c5/1?token-time=2145916800&token-hash=efJ0dqJhGOR-FtExvoH3ELw8CPpHo5I3Ce6ixNaJmuk%3D",
            "last_name": "",
            "social_connections": {
                "deviantart": null,
                "discord": null,
                "facebook": null,
                "reddit": null,
                "spotify": null,
                "twitch": null,
                "twitter": null,
                "youtube": null
            },
            "thumb_url": "https://c10.patreonusercontent.com/3/eyJoIjoxMDAsInciOjEwMH0%3D/patreon-media/p/user/2822191/8c0bc222ed0c43e68c620fa191b0a0c5/1?token-time=2145916800&token-hash=NoMxBCUckp3EBrPgADzsEdUI3uFV13EB_wRx4LqIh4I%3D",
            "url": "https://www.patreon.com/podsync",
            "vanity": "podsync"
        },
        "id": "2822191",
        "type": "user"
    }, {
        "attributes": {},
        "id": "1048240",
        "type": "tier"
    }, {
        "attributes": {},
        "id": "2140517",
        "type": "tier"
    }, {
        "attributes": {},
        "id": "342492",
        "type": "goal"
    }, {
        "attributes": {},
        "id": "605110",
        "type": "goal"
    }, {
        "attributes": {},
        "id": "605111",
        "type": "goal"
    }],
    "meta": {
        "pagination": {
            "total": 1
        }
    }
}
`

func TestClient_GetCampaigns(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/oauth2/v2/campaigns", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, testCampaignListResponse)
		assert.NoError(t, err)
	})

	campaigns, err := client.GetCampaigns()
	require.NoError(t, err)
	require.NotNil(t, campaigns)
	require.Len(t, campaigns, 1)

	campaign := campaigns[0]
	assert.Equal(t, "new podcasting experience - Podsync", campaign.CreationName)

	assert.NotNil(t, campaign.Creator)
	assert.Equal(t, "2822191", campaign.Creator.ID)
	assert.Equal(t, "Max", campaign.Creator.FullName)

	assert.Len(t, campaign.Tiers, 2)
	assert.Equal(t, "1048240", campaign.Tiers[0].ID)
	assert.Equal(t, "2140517", campaign.Tiers[1].ID)

	assert.Nil(t, campaign.Tiers[0].TierAttributes)
	assert.Nil(t, campaign.Tiers[1].TierAttributes)

	assert.Len(t, campaign.Goals, 3)
	assert.EqualValues(t, "342492", campaign.Goals[0].ID)
	assert.EqualValues(t, "605110", campaign.Goals[1].ID)
	assert.EqualValues(t, "605111", campaign.Goals[2].ID)

	assert.Nil(t, campaign.Goals[0].GoalAttributes)
	assert.Nil(t, campaign.Goals[1].GoalAttributes)
	assert.Nil(t, campaign.Goals[2].GoalAttributes)

	assert.Len(t, campaign.Benefits, 0)
}

const testCampaignResponse = `
{
    "data": {
        "attributes": {
            "created_at": "2016-02-02T19:58:18+00:00",
            "creation_name": "new podcasting experience - Podsync",
            "discord_server_id": null,
            "google_analytics_id": null,
            "has_rss": false,
            "has_sent_rss_notify": false,
            "image_small_url": "https://c10.patreonusercontent.com/3/eyJoIjo2NDAsInciOjY0MH0%3D/patreon-media/p/campaign/278915/c17598520740471ca6c0ffe553ade639/1?token-time=2145916800&token-hash=Au6pSGBsM1mQ4D3YFFtbrJHit_G99uOvIyJs_C9uT7E%3D",
            "image_url": "https://c10.patreonusercontent.com/3/eyJ3IjoxOTIwfQ%3D%3D/patreon-media/p/campaign/278915/c17598520740471ca6c0ffe553ade639/1?token-time=2145916800&token-hash=thO-8NggSSPFYnLMeW9YowqCYKgsfTtwah_eoak6tVc%3D",
            "is_charged_immediately": true,
            "is_monthly": true,
            "is_nsfw": false,
            "main_video_embed": "",
            "main_video_url": "",
            "one_liner": null,
            "patron_count": 468,
            "pay_per_name": "month",
            "pledge_url": "/join/podsync",
            "published_at": "2016-02-02T20:11:19+00:00",
            "rss_artwork_url": null,
            "rss_feed_title": null,
            "summary": "<a href=\"http://podsync.net/\" rel=\"nofollow\">Podsync</a> - is a simple, free service that lets you listen to any YouTube / Vimeo channels, playlists or user videos in podcast format.<br><br><strong>Idea:</strong><br>Podcast applications have a rich functionality for content delivery - automatic download of new episodes, remembering last played position, sync between devices and offline listening. This functionality is not available on YouTube and Vimeo. So the aim of\u00a0<a href=\"http://podsync.net/\" rel=\"nofollow\">Podsync</a> is to make your life easier and enable you to view/listen to content on any device in podcast client.<br><br>It's my hobby project, so to continue to support and improve it, I need your help. Your money will go into paying my server bills and adding new features.<br><br>",
            "thanks_embed": "",
            "thanks_msg": "You are awesome! Thank you so much!<br>",
            "thanks_video_url": null
        },
        "id": "278915",
        "relationships": {
            "benefits": {
                "data": []
            },
            "creator": {
                "data": {
                    "id": "2822191",
                    "type": "user"
                }
            },
            "goals": {
                "data": [{
                    "id": "342492",
                    "type": "goal"
                }, {
                    "id": "605110",
                    "type": "goal"
                }, {
                    "id": "605111",
                    "type": "goal"
                }]
            },
            "tiers": {
                "data": [{
                    "id": "1048240",
                    "type": "tier"
                }, {
                    "id": "2140517",
                    "type": "tier"
                }]
            }
        },
        "type": "campaign"
    },
    "included": [{
        "attributes": {
            "about": "",
            "created": "2016-02-02T19:56:14+00:00",
            "email": "mail@gmail.com",
            "first_name": "Max",
            "full_name": "Max",
            "image_url": "https://c10.patreonusercontent.com/3/eyJ3IjoyMDB9/patreon-media/p/user/2822191/8c0bc222ed0c43e68c620fa191b0a0c5/1?token-time=2145916800&token-hash=efJ0dqJhGOR-FtExvoH3ELw8CPpHo5I3Ce6ixNaJmuk%3D",
            "last_name": "",
            "social_connections": {
                "deviantart": null,
                "discord": null,
                "facebook": null,
                "reddit": null,
                "spotify": null,
                "twitch": null,
                "twitter": null,
                "youtube": null
            },
            "thumb_url": "https://c10.patreonusercontent.com/3/eyJoIjoxMDAsInciOjEwMH0%3D/patreon-media/p/user/2822191/8c0bc222ed0c43e68c620fa191b0a0c5/1?token-time=2145916800&token-hash=NoMxBCUckp3EBrPgADzsEdUI3uFV13EB_wRx4LqIh4I%3D",
            "url": "https://www.patreon.com/podsync",
            "vanity": "podsync"
        },
        "id": "2822191",
        "type": "user"
    }, {
        "attributes": {},
        "id": "1048240",
        "type": "tier"
    }, {
        "attributes": {},
        "id": "2140517",
        "type": "tier"
    }, {
        "attributes": {},
        "id": "342492",
        "type": "goal"
    }, {
        "attributes": {},
        "id": "605110",
        "type": "goal"
    }, {
        "attributes": {},
        "id": "605111",
        "type": "goal"
    }],
    "links": {
        "self": "https://www.patreon.com/api/oauth2/v2/campaigns/278915"
    }
}
`

func TestClient_GetCampaignByID(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/oauth2/v2/campaigns/2", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, testCampaignResponse)
		assert.NoError(t, err)
	})

	campaign, err := client.GetCampaignByID("2")
	require.NoError(t, err)
	require.NotNil(t, campaign)
	require.NotNil(t, campaign.CampaignAttributes)

	assert.Equal(t, "278915", campaign.ID)
	assert.Equal(t, "new podcasting experience - Podsync", campaign.CreationName)
	assert.EqualValues(t, 468, campaign.PatronCount)
	assert.Equal(t, "/join/podsync", campaign.PledgeURL)
	assert.True(t, campaign.IsMonthly)
	assert.True(t, campaign.IsChargedImmediately)
}

const testGetMembersByCampaignIDResponse = `
{
    "data": [{
        "attributes": {},
        "id": "00478b66-a597-4a6a-b8c1-111111111111",
        "type": "member"
    }, {
        "attributes": {},
        "id": "007fb93d-f393-4b01-afae-222222222222",
        "type": "member"
    }],
    "links": {
        "next": "https://www.patreon.com/api/oauth2/v2/campaigns/278915/members?page[count]=2&page[cursor]=Prn1AE3MjZONYS111111114"
    },
    "meta": {
        "pagination": {
            "cursors": {
                "next": "Prn1AE3MjZONYS111111114"
            },
            "total": 1035
        }
    }
}
`

func TestClient_GetMembersByCampaignID(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/oauth2/v2/campaigns/278915/members", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, testGetMembersByCampaignIDResponse)
		assert.NoError(t, err)
	})

	members, err := client.GetMembersByCampaignID("278915")
	require.NoError(t, err)
	require.Len(t, members, 2)

	assert.Equal(t, "00478b66-a597-4a6a-b8c1-111111111111", members[0].ID)
	assert.Equal(t, "007fb93d-f393-4b01-afae-222222222222", members[1].ID)
}

const testMemberResponse = `
{
    "data": {
        "attributes": {},
        "id": "00478b66-a597-4a6a-b8c1-0e1a395ab613",
        "relationships": {
            "address": {
                "data": null
            },
            "campaign": {
                "data": {
                    "id": "278915",
                    "type": "campaign"
                },
                "links": {
                    "related": "https://www.patreon.com/api/oauth2/v2/campaigns/278915"
                }
            },
            "currently_entitled_tiers": {
                "data": [{
                    "id": "1048240",
                    "type": "tier"
                }]
            },
            "user": {
                "data": {
                    "id": "75985",
                    "type": "user"
                }
            }
        },
        "type": "member"
    },
    "included": [{
        "attributes": {
            "created_at": "2016-02-02T19:58:18+00:00",
            "creation_name": "new podcasting experience - Podsync",
            "discord_server_id": null,
            "google_analytics_id": "UA-22222222-2",
            "has_rss": false,
            "has_sent_rss_notify": false,
            "image_small_url": "https://c10.patreonusercontent.com/3/eyJoIjo2NDAsInciOjY0MH0%3D/patreon-media/p/campaign/278915/c17598520740471ca6c0ffe553ade639/1?token-time=2145916800&token-hash=Au6pSGBsM1mQ4D3YFFtbrJHit_G99uOvIyJs_C9uT7E%3D",
            "image_url": "https://c10.patreonusercontent.com/3/eyJ3IjoxOTIwfQ%3D%3D/patreon-media/p/campaign/278915/c17598520740471ca6c0ffe553ade639/1?token-time=2145916800&token-hash=thO-8NggSSPFYnLMeW9YowqCYKgsfTtwah_eoak6tVc%3D",
            "is_charged_immediately": true,
            "is_monthly": true,
            "is_nsfw": false,
            "main_video_embed": "",
            "main_video_url": "",
            "one_liner": null,
            "patron_count": 464,
            "pay_per_name": "month",
            "pledge_url": "/join/podsync",
            "published_at": "2016-02-02T20:11:19+00:00",
            "rss_artwork_url": null,
            "rss_feed_title": null,
            "summary": "<a href=\"http://podsync.net/\" rel=\"nofollow\">Podsync</a> - is a simple, free service that lets you listen to any YouTube / Vimeo channels, playlists or user videos in podcast format.<br><br><strong>Idea:</strong><br>Podcast applications have a rich functionality for content delivery - automatic download of new episodes, remembering last played position, sync between devices and offline listening. This functionality is not available on YouTube and Vimeo. So the aim of\u00a0<a href=\"http://podsync.net/\" rel=\"nofollow\">Podsync</a> is to make your life easier and enable you to view/listen to content on any device in podcast client.<br><br>It's my hobby project, so to continue to support and improve it, I need your help. Your money will go into paying my server bills and adding new features.<br><br>",
            "thanks_embed": "",
            "thanks_msg": "You are awesome!",
            "thanks_video_url": null
        },
        "id": "278915",
        "type": "campaign"
    }, {
        "attributes": {
            "about": null,
            "created": "2017-08-30T14:07:37+00:00",
            "first_name": "Marcel",
            "full_name": "Marcel",
            "image_url": "https://c8.patreon.com/2/200/75985",
            "last_name": "",
            "social_connections": {
                "deviantart": null,
                "discord": null,
                "facebook": null,
                "reddit": null,
                "spotify": null,
                "twitch": null,
                "twitter": null,
                "youtube": null
            },
            "thumb_url": "",
            "url": "https://www.patreon.com/user?u=75985",
            "vanity": null
        },
        "id": "75985",
        "type": "user"
    }, {
        "attributes": {},
        "id": "1048240",
        "type": "tier"
    }],
    "links": {
        "self": "https://www.patreon.com/api/oauth2/v2/members/00478b66-a597-4a6a-b8c1-0e1a395ab613"
    }
}
`

func TestClient_GetMemberByID(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/oauth2/v2/members/00478b66-a597-4a6a-b8c1-0e1a395ab613", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, testMemberResponse)
		assert.NoError(t, err)
	})

	member, err := client.GetMemberByID("00478b66-a597-4a6a-b8c1-0e1a395ab613")
	require.NoError(t, err)
	require.NotNil(t, member)

	assert.Equal(t, "00478b66-a597-4a6a-b8c1-0e1a395ab613", member.ID)

	require.NotNil(t, member.Campaign)
	require.NotNil(t, member.Campaign.CampaignAttributes)

	assert.Equal(t, "278915", member.Campaign.ID)
	assert.Equal(t, "UA-22222222-2", member.Campaign.GoogleAnalyticsID)
	assert.Equal(t, "You are awesome!", member.Campaign.ThanksMsg)

	require.NotNil(t, member.User)
	require.NotNil(t, member.User.UserAttributes)

	assert.Equal(t, "Marcel", member.User.FirstName)
	assert.Equal(t, "Marcel", member.User.FullName)
	assert.Equal(t, "https://www.patreon.com/user?u=75985", member.User.URL)
	assert.Equal(t, "https://c8.patreon.com/2/200/75985", member.User.ImageURL)

	require.Len(t, member.CurrentlyEntitledTiers, 1)
	assert.Equal(t, "1048240", member.CurrentlyEntitledTiers[0].ID)
	assert.Nil(t, member.CurrentlyEntitledTiers[0].TierAttributes)
}
