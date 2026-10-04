// api/internal/guilds/bis_lookup.go
package guilds

import "github.com/jhunthrop/foreversixty/api/internal/bis"

// bisDataBuild is the client build the guild control centre's own BiS reads use - the same
// build api/cmd/seedguild/gear.go's own bisBuild constant names, and the one real,
// alternatives-bearing BiS catalogue in this repo today. Hardcoded rather than read from
// config, matching that tool's own choice: there is no "current active build" setting
// anywhere in this service yet (checked api/internal/config), so inventing one for this one
// reader would be new config surface the contract never asked for. A future lane that adds
// a real active-build setting should thread it through here instead of this constant.
const bisDataBuild = "1.60.1.70009"

// bisBandLevel is the level band every raider on this page is assumed to be - level 60,
// design spec §4.E's own assumption throughout the guild control centre.
const bisBandLevel = 60

// loadBandFor is FileSlugFor+LoadBand, folded into "do we have a band for this character at
// all" - any failure (unknown spec, missing file, read/parse error) reads as false, never an
// error: a BiS data gap must never fail the roster, standing, readiness or loot endpoints,
// the same "a data gap is an honest null, not a 500" rule fs1.Decode's own package comment
// states.
func (s *Store) loadBandFor(classSlug, specName, faction string) (bis.Band, bool) {
	if specName == "" || faction == "" {
		return bis.Band{}, false
	}
	fileSlug, ok := bis.FileSlugFor(classSlug, specName)
	if !ok {
		return bis.Band{}, false
	}
	band, ok, err := bis.LoadBand(s.DataDir, bisDataBuild, fileSlug, bisBandLevel, faction)
	if err != nil || !ok {
		return bis.Band{}, false
	}
	return band, true
}
