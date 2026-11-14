// api/internal/bnetapi/directory.go
package bnetapi

import (
	"context"
	"encoding/json"
	"fmt"
)

// CharacterStatus is GET .../character/{slug}/status: the validation Blizzard's API terms
// ask for at least every 30 days (section 2(s)). Valid is false for a character that no
// longer exists under that id; a 404 means the same or a profile not public yet.
type CharacterStatus struct {
	ID    int64
	Valid bool
}

type characterStatusResponse struct {
	ID      int64 `json:"id"`
	IsValid bool  `json:"is_valid"`
}

// Status reads a character's validation status.
func (c *Client) Status(ctx context.Context, region, realmSlug, name string) (CharacterStatus, error) {
	raw, err := c.characterResource(ctx, "status", region, realmSlug, name)
	if err != nil {
		return CharacterStatus{}, err
	}
	var res characterStatusResponse
	if err := json.Unmarshal(raw, &res); err != nil {
		return CharacterStatus{}, fmt.Errorf("bnetapi: status: decode: %w", err)
	}
	return CharacterStatus{ID: res.ID, Valid: res.IsValid}, nil
}

// Professions reads a character's professions verbatim (primaries and secondaries, each
// with tiers and skill points).
func (c *Client) Professions(ctx context.Context, region, realmSlug, name string) (json.RawMessage, error) {
	return c.characterResource(ctx, "professions", region, realmSlug, name)
}

// playableClassSlugs is Blizzard's playable class id to the site's class slug: the ids are
// fixed across every WoW game (the playable-class index, testdata/forever).
var playableClassSlugs = map[int64]string{
	1: "warrior", 2: "paladin", 3: "hunter", 4: "rogue", 5: "priest",
	7: "shaman", 8: "mage", 9: "warlock", 11: "druid",
}

// ClassSlugByID is the class slug of a playable class id; "" for an id it does not know.
func ClassSlugByID(id int64) string { return playableClassSlugs[id] }
