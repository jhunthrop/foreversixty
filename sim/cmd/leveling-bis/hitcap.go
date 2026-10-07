package main

import "github.com/wowsims/classic/sim/core"

// hitToCap is how far a band's weights character is from the miss-table
// caps, in engine percent (one point is one percent of hit chance, the unit
// the published hit weight is measured in per percent). The site reads it to
// say "hit to cap first": hit is worth its full weight only up to the cap,
// and nothing past it.
type hitToCap struct {
	// Baseline is the weights character's own hit, from gear, talents and
	// buffs.
	Baseline float64 `json:"baseline"`
	// Specials is the hit still worth full value to special attacks:
	// the cap (base miss plus the level-gap suppression) less Baseline.
	Specials float64 `json:"specials"`
	// White is the hit still worth anything to white swings. It is set only
	// for a dual wielder, whose white swings miss 19 points more than a
	// special does; everyone else's white swings share the special cap.
	White *float64 `json:"white,omitempty"`
}

// hitToCapFromProfile reads the engine's hit profile into the published
// figure. A spec that does not swing a weapon (a caster) has no physical
// miss table to cap against, so it publishes none.
func hitToCapFromProfile(profile core.HitProfile) *hitToCap {
	if !profile.Physical {
		return nil
	}
	out := &hitToCap{Baseline: profile.Hit, Specials: profile.ToSpecialCap()}
	if profile.DualWielding {
		white := profile.ToWhiteCap()
		out.White = &white
	}
	return out
}
