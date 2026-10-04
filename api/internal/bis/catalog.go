// api/internal/bis/catalog.go
//
// Package bis is the guild control centre's own reader of the leveling-BiS tables
// (data/builds/<build>/bis/<spec>.json) and the gear-gap gain rule the contract
// (docs/contracts/2026-10-04-guild-centre-api.md) ports from web/src/lib/home/upgrades.ts.
// Used by the readiness, loot and roster-standing endpoints in api/internal/guilds.
package bis

import (
	"strings"

	"github.com/jhunthrop/foreversixty/sim/specs"
)

// FileSlugFor turns a roster row's display class ("hunter", already lowercase in every
// table this package reads it from) and spec name ("Marksmanship") into the BiS file's own
// basename ("hunter-marksmanship") - sim/specs.All, the generated copy of
// data/curated/specs.json every other lane already imports this same way
// (api/internal/sims/score.go's own specSlugFor), rather than a second runtime read of the
// JSON file itself (data/curated/ is not copied into the API's Docker image - only
// data/builds/ is, see api/Dockerfile - so a literal file read here would 404 in
// production). ok is false for a class/spec this catalogue does not name, which this
// package's callers treat as "no gear table for this character," never an error.
func FileSlugFor(classSlug, specName string) (string, bool) {
	classSlug = strings.ToLower(strings.TrimSpace(classSlug))
	specName = strings.TrimSpace(specName)
	for _, s := range specs.All {
		if s.ClassSlug == classSlug && strings.EqualFold(s.Name, specName) {
			return s.Spec, true
		}
	}
	return "", false
}
