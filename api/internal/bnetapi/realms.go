// api/internal/bnetapi/realms.go
package bnetapi

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

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

// realmWorkerCount bounds how many realm-detail requests Realms has in
// flight at once. Fetched serially a full Era region (66 realms) took
// ~7s from a laptop — long enough that the importer's 5s budget expired
// partway through on the first production sign-in (2026-09-22), leaving
// every later realm's Type empty and (before the never-cache-a-partial-
// list fix below) caching that partial list for 24h. 8 in flight keeps a
// full region's fetch under a second.
const realmWorkerCount = 8

// warmRealmsTimeout bounds WarmRealms's own startup fetch, independent of
// whatever context main.go passes it, so a slow Blizzard response at
// service start can never block indefinitely.
const warmRealmsTimeout = 30 * time.Second

// Realms lists every realm in region's dynamic namespace, with its type
// and category, caching the result in memory for realmCacheTTL. An
// unrecognised realm type is logged at WARN (spec §3) and mapped to
// "normal" by RulesetOf. A list in which any realm's detail fetch failed
// is never cached: Realms returns the error instead, so a caller's own
// per-run fallback (bnetimport resolves an unknown realm to "normal" for
// that run only) applies, rather than a bad result being cached for
// realmCacheTTL.
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

	realms := make([]Realm, len(idx.Realms))
	for i, r := range idx.Realms {
		realms[i] = Realm{ID: r.ID, Slug: r.Slug, Name: r.Name}
	}

	if err := c.fetchRealmDetails(ctx, region, realms); err != nil {
		return nil, err
	}

	c.realmMu.Lock()
	c.realmCache[region] = realmCacheEntry{realms: realms, cachedAt: c.clock()}
	c.realmMu.Unlock()
	return realms, nil
}

// fetchRealmDetails fills in each of realms' Type and Category
// concurrently, realmWorkerCount at a time. Each worker only ever writes
// the index it was handed, so no synchronization is needed on realms
// itself. It waits for every in-flight request to finish (rather than
// canceling the rest on the first failure) so a single bad realm never
// leaves the others half-fetched, then returns the first error seen, if
// any.
func (c *Client) fetchRealmDetails(ctx context.Context, region string, realms []Realm) error {
	type job struct {
		index int
		slug  string
	}
	jobs := make(chan job)
	errs := make([]error, len(realms))
	var wg sync.WaitGroup

	for w := 0; w < realmWorkerCount; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				var detail realmDetailResponse
				detailURL := c.APIHost(region) + "/data/wow/realm/" + j.slug + "?namespace=" + c.DynamicNamespace(region)
				if err := c.getJSON(ctx, "realm_detail", detailURL, &detail); err != nil {
					c.logger().Warn("bnetapi", "op", "realm_detail", "region", region, "realm", j.slug, "err", err)
					errs[j.index] = err
					continue
				}
				realms[j.index].Type = detail.Type.Type
				realms[j.index].Category = detail.Category
				if realms[j.index].Type != "" && !knownRealmTypes[strings.ToUpper(realms[j.index].Type)] {
					c.logger().Warn("bnetapi", "op", "realm_type", "region", region, "realm", j.slug,
						"type", realms[j.index].Type, "mapped_to", RulesetOf(realms[j.index]))
				}
			}
		}()
	}
	for i, r := range realms {
		jobs <- job{index: i, slug: r.Slug}
	}
	close(jobs)
	wg.Wait()

	for _, err := range errs {
		if err != nil {
			return fmt.Errorf("bnetapi: realms: %s: %w", region, err)
		}
	}
	return nil
}

// WarmRealms fetches and caches every one of regions' realm lists before
// the service takes traffic (spec A1), so a first sign-in's import never
// pays the cost of a cold per-region realm fetch inside its 5s budget. A
// region that fails to warm is logged and simply fetched cold on first
// use, like any cache miss.
func (c *Client) WarmRealms(ctx context.Context, regions []string) {
	ctx, cancel := context.WithTimeout(ctx, warmRealmsTimeout)
	defer cancel()
	for _, region := range regions {
		if _, err := c.Realms(ctx, region); err != nil {
			c.logger().Warn("bnetapi", "op", "warm_realms", "region", region, "err", err)
		}
	}
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
