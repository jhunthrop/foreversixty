// api/internal/guilds/bis_lookup.go
package guilds

import (
	"sync"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
	"github.com/jhunthrop/foreversixty/api/internal/upgrades"
)

// BisBuild is the client build every BiS read of this service uses (the guild control
// centre's readiness, loot and roster reads, and the Discord bot's lookups): the site BiS
// page's own rule (bis.SourceBuild of the active build, web/scripts/
// bis-source-build.mjs), so the API and the site never read two different BiS builds. Empty
// when no build data is loaded, which every caller reads as "no band". SourceBuild reads the
// filesystem, and this runs per roster row and per loot item, so the answer is kept for as
// long as DataDir and Trees stay the ones it was worked out from (both are set at boot).
func (s *Store) BisBuild() string {
	s.bisBuild.mu.Lock()
	defer s.bisBuild.mu.Unlock()
	c := &s.bisBuild
	if !c.set || c.dataDir != s.DataDir || c.trees != s.Trees {
		c.build, c.dataDir, c.trees, c.set = sourceBuildOf(s.DataDir, s.Trees), s.DataDir, s.Trees, true
	}
	return c.build
}

// bisBuildCache is BisBuild's kept answer and the inputs it came from.
type bisBuildCache struct {
	mu      sync.Mutex
	set     bool
	dataDir string
	trees   *trees.Data
	build   string
}

func sourceBuildOf(dataDir string, data *trees.Data) string {
	if data == nil {
		return ""
	}
	active, ok := data.Active()
	if !ok {
		return ""
	}
	return bis.SourceBuild(dataDir, active.Version)
}

// BisBuild is Store.BisBuild.
func (s *Service) BisBuild() string { return s.Store.BisBuild() }

// bisBandLevel is the level band a raider is measured against: raidLevel. Only graded
// characters read it for readiness; one under 60 is not graded (computeReadiness).
const bisBandLevel = raidLevel

// loadBandFor is FileSlugFor+LoadPresetBand under the preset the site and /char read at the band
// (upgrades.PresetFor: raid at 60), folded into "do we have a band for this character at
// all" - any failure (unknown spec, missing file, read/parse error) reads as false, never an
// error: a BiS data gap must never fail the roster, standing, readiness or loot endpoints,
// the same "a data gap is an honest null, not a 500" rule fs1.Decode's own package comment
// states.
func (s *Store) loadBandFor(classSlug, specName, faction string) (bis.Band, bool) {
	return s.loadBandAt(classSlug, specName, faction, bisBandLevel)
}

// loadBandAt is loadBandFor at level band (20, 30, ... 60), under that band's preset.
func (s *Store) loadBandAt(classSlug, specName, faction string, level int) (bis.Band, bool) {
	if specName == "" || faction == "" {
		return bis.Band{}, false
	}
	fileSlug, ok := bis.FileSlugFor(classSlug, specName)
	if !ok {
		return bis.Band{}, false
	}
	band, ok, err := bis.LoadPresetBand(s.DataDir, s.BisBuild(), fileSlug, level, faction, upgrades.PresetFor(level))
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
