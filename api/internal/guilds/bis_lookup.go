// api/internal/guilds/bis_lookup.go
package guilds

import "github.com/jhunthrop/foreversixty/api/internal/bis"

// BisBuild is the client build every BiS read of this service uses (the guild control
// centre's readiness, loot and roster reads, and the Discord bot's lookups): the site BiS
// page's own rule (bis.SourceBuild of the newest loaded build, web/scripts/
// bis-source-build.mjs), so the API and the site never read two different BiS builds. Empty
// when no build data is loaded, which every caller reads as "no band".
func (s *Store) BisBuild() string {
	if s.Trees == nil {
		return ""
	}
	latest, ok := s.Trees.Latest()
	if !ok {
		return ""
	}
	return bis.SourceBuild(s.DataDir, latest.Version)
}

// BisBuild is Store.BisBuild.
func (s *Service) BisBuild() string { return s.Store.BisBuild() }

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
	band, ok, err := bis.LoadBand(s.DataDir, s.BisBuild(), fileSlug, bisBandLevel, faction)
	if err != nil || !ok {
		return bis.Band{}, false
	}
	return band, true
}

// factionForRace resolves an FS1 export's race slug ("human", "undead", ...) to its
// faction via the same build data loot.go's per-class item table already reads
// (Store.Trees, data/builds/<build>/races.json) - never a second hardcoded race table that
// could drift from that file. fight_metrics carries no faction column at all (never
// written by a real ingest or by api/cmd/seedguild - see that package's own
// insertFightsAndMetrics comment), so this is the only source of faction HomeRoster has;
// ok is false for an empty slug, a nil Trees (a harness that never loaded build data), or a
// slug the build's races.json does not name - every case this package already treats as
// "no band available," never an error.
func (s *Store) factionForRace(raceSlug string) (string, bool) {
	if raceSlug == "" || s.Trees == nil {
		return "", false
	}
	build, ok := s.Trees.Build(s.BisBuild())
	if !ok {
		return "", false
	}
	for _, r := range build.Races() {
		if r.Slug == raceSlug {
			return r.Faction, true
		}
	}
	return "", false
}
