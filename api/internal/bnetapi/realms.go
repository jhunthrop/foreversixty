// api/internal/bnetapi/realms.go
package bnetapi

import (
	"context"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/character"
)

// Realm is one realm's identity and the fields RulesetOf needs.
type Realm struct {
	ID       int64
	Slug     string
	Name     string
	Type     string // Blizzard's own realm type: NORMAL, PVP, RP, RP_PVP
	Category string
}

type realmIndexResponse struct {
	Realms []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"realms"`
}

type realmDetailResponse struct {
	Type struct {
		Type string `json:"type"`
	} `json:"type"`
	Category string `json:"category"`
}

// knownRealmTypes is every realm type value RulesetOf maps explicitly.
var knownRealmTypes = map[string]bool{"NORMAL": true, "PVP": true, "RP": true, "RP_PVP": true}

// Realms lists every realm in region's dynamic namespace, with its type
// and category, caching the result in memory for realmCacheTTL. An
// unrecognised realm type is logged at WARN (spec §3) and mapped to
// "normal" by RulesetOf.
func (c *Client) Realms(ctx context.Context, region string) ([]Realm, error) {
	c.realmMu.Lock()
	if entry, ok := c.realmCache[region]; ok && c.clock().Sub(entry.cachedAt) < realmCacheTTL {
		c.realmMu.Unlock()
		return entry.realms, nil
	}
	c.realmMu.Unlock()

	var idx realmIndexResponse
	indexURL := c.APIHost(region) + "/data/wow/realm/index?namespace=" + c.DynamicNamespace(region)
	if err := c.getJSON(ctx, "realms", indexURL, &idx); err != nil {
		return nil, err
	}

	realms := make([]Realm, 0, len(idx.Realms))
	for _, r := range idx.Realms {
		realm := Realm{ID: r.ID, Slug: r.Slug, Name: r.Name}
		var detail realmDetailResponse
		detailURL := c.APIHost(region) + "/data/wow/realm/" + r.Slug + "?namespace=" + c.DynamicNamespace(region)
		if err := c.getJSON(ctx, "realm_detail", detailURL, &detail); err != nil {
			c.logger().Warn("bnetapi", "op", "realm_detail", "region", region, "realm", r.Slug, "err", err)
		} else {
			realm.Type = detail.Type.Type
			realm.Category = detail.Category
			if realm.Type != "" && !knownRealmTypes[strings.ToUpper(realm.Type)] {
				c.logger().Warn("bnetapi", "op", "realm_type", "region", region, "realm", r.Slug,
					"type", realm.Type, "mapped_to", RulesetOf(realm))
			}
		}
		realms = append(realms, realm)
	}

	c.realmMu.Lock()
	c.realmCache[region] = realmCacheEntry{realms: realms, cachedAt: c.clock()}
	c.realmMu.Unlock()
	return realms, nil
}

// RulesetOf maps a realm's Blizzard type (and, for hardcore, its name or
// category) onto one of the site's four rulesets. A pure function so it
// can be table-tested with no client or network.
func RulesetOf(realm Realm) string {
	if strings.Contains(strings.ToLower(realm.Category), "hardcore") ||
		strings.Contains(strings.ToLower(realm.Name), "hardcore") {
		return character.RulesetHardcore
	}
	switch strings.ToUpper(realm.Type) {
	case "PVP":
		return character.RulesetPvP
	case "RP", "RP_PVP":
		return character.RulesetRP
	default:
		return character.RulesetNormal
	}
}
