// api/internal/bnetapi/media.go
package bnetapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// trustedMediaHost is the only host character-media asset URLs are ever
// returned for. Blizzard's own render CDN is stable; anything else would
// mean either an API change worth investigating or an unexpected
// response, and this site never embeds an arbitrary host on the account
// or character page (.superpowers/account-visual-brief.md §A2).
const trustedMediaHost = "https://render.worldofwarcraft.com/"

// Media is a character's Blizzard-hosted images: a small square avatar
// and a full-body render, when Blizzard has generated them.
type Media struct {
	AvatarURL string
	RenderURL string
}

type characterMediaResponse struct {
	Assets []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"assets"`
}

// CharacterMedia reads GET .../character-media and returns both the
// typed URLs and the raw response body (stored verbatim as
// characters.bnet_media, the same capture shape Character and Equipment
// use). AvatarURL comes from the "avatar" asset; RenderURL comes from
// "main-raw", falling back to "main" when Blizzard has not generated the
// raw render. A URL under any host but trustedMediaHost is dropped (and
// logged, op=media_host_rejected) rather than returned, so a caller can
// never end up embedding an untrusted host.
func (c *Client) CharacterMedia(ctx context.Context, region, realmSlug, name string) (Media, json.RawMessage, error) {
	token, err := c.AppToken(ctx)
	if err != nil {
		return Media{}, nil, fmt.Errorf("bnetapi: character_media: %w", err)
	}
	u := c.APIHost(region) + "/profile/wow/character/" + url.PathEscape(realmSlug) + "/" +
		url.PathEscape(strings.ToLower(name)) + "/character-media?namespace=" + c.ProfileNamespace(region)
	body, err := c.getBytes(ctx, "character_media", u, token)
	if err != nil {
		return Media{}, nil, err
	}
	var res characterMediaResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return Media{}, nil, fmt.Errorf("bnetapi: character_media: decode: %w", err)
	}

	var media Media
	var main string
	for _, asset := range res.Assets {
		switch asset.Key {
		case "avatar":
			media.AvatarURL = c.trustedMediaURL(asset.Value)
		case "main-raw":
			media.RenderURL = c.trustedMediaURL(asset.Value)
		case "main":
			main = c.trustedMediaURL(asset.Value)
		}
	}
	if media.RenderURL == "" {
		media.RenderURL = main
	}
	return media, json.RawMessage(body), nil
}

// trustedMediaURL returns raw unchanged when it is under trustedMediaHost,
// or "" (after logging op=media_host_rejected) otherwise.
func (c *Client) trustedMediaURL(raw string) string {
	if strings.HasPrefix(raw, trustedMediaHost) {
		return raw
	}
	if raw != "" {
		c.logger().Warn("bnetapi", "op", "media_host_rejected", "url", raw)
	}
	return ""
}
