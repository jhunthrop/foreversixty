// api/internal/auth/build_fields.go
package auth

import (
	"github.com/jhunthrop/foreversixty/api/internal/fs1"
	"github.com/jhunthrop/foreversixty/api/internal/spec"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// buildFieldsFromExport decodes c.Build's stored export string (fs1.Decode) and fills in
// everything GET /v1/me adds on top of Build's existing source/captured_at: the build's
// gear, talent split, level and data build, the character's own level/race/faction/realm
// wherever Battle.net has not already set them (Battle.net always wins -- this only ever
// fills a gap), and the character's primary spec, named from the decoded talent split.
//
// c.Build is already non-nil when this is called (the caller only calls it once it has
// one). A decode failure (fs1.Decode's ok=false) leaves every one of those fields exactly
// as it already was and logs at debug -- the one thing this function must never do is turn
// a malformed or out-of-date export into a failed GET /v1/me.
func (s *Store) buildFieldsFromExport(c *Character, export string) {
	decoded, ok := fs1.Decode(export)
	if !ok {
		if export != "" {
			s.logger().Debug("auth", "op", "build_fields", "character_key", c.Key,
				"reason", "stored export does not decode as FS1")
		}
		return
	}

	if len(decoded.Gear) > 0 {
		c.Build.Gear = decoded.Gear
	}
	c.Build.Talents = &BuildTalents{Trees: decoded.Talents.Trees, Points: decoded.Talents.Points}
	c.Build.DataBuild = decoded.DataBuild
	if decoded.HasLevel {
		level := decoded.Level
		c.Build.Level = &level
	}

	if c.Level == nil && decoded.HasLevel {
		level := decoded.Level
		c.Level = &level
	}
	if c.Realm == "" && decoded.HasWho && decoded.Realm != "" {
		c.Realm = decoded.Realm
	}

	build := s.treesBuildFor(decoded.DataBuild)
	if build == nil {
		return
	}
	if c.Race == "" || c.Faction == "" {
		if race, ok := raceBySlug(build, decoded.RaceSlug); ok {
			if c.Race == "" {
				c.Race = race.Name
			}
			if c.Faction == "" {
				c.Faction = race.Faction
			}
		}
	}
	c.Spec = primarySpec(build, decoded.ClassSlug, decoded.Talents.Points)
}

// treesBuildFor is the client build decoded.DataBuild names, or the newest loaded build
// when that exact version is not loaded (an export a data refresh has moved past, or one
// with no data build recorded at all). Trees being nil -- no client data loaded at all --
// is the only case this returns nil.
func (s *Store) treesBuildFor(dataBuild string) *trees.Build {
	if s.Trees == nil {
		return nil
	}
	if dataBuild != "" {
		if b, ok := s.Trees.Build(dataBuild); ok {
			return b
		}
	}
	b, _ := s.Trees.Latest()
	return b
}

// raceBySlug finds a build's race by its site slug ("troll"), the same vocabulary
// Export.RACE_SLUGS (the addon) and bnetbuild's RaceTable both write into an FS1 head's
// race field -- not a second slug table, the one trees.Data already loads from
// data/builds/<version>/races.json.
func raceBySlug(build *trees.Build, slug string) (trees.Race, bool) {
	for _, r := range build.Races() {
		if r.Slug == slug {
			return r, true
		}
	}
	return trees.Race{}, false
}

// primarySpec is the character's primary talent tree, named from a decoded point split:
// the tree with strictly the most points. It answers "" -- "not known yet" to the web --
// when the class slug does not resolve, when the split does not have exactly one entry
// per tree the build has, when no points are spent, or when two or more trees tie for the
// most. spec.Inferrer.Name breaks that same tie toward the earlier tree, because a ranking
// row must always show something; GET /v1/me is not that contract and would rather omit
// Spec than guess at a 0/0/0 or an even 20/20/0 split.
func primarySpec(build *trees.Build, classSlug string, points []int) string {
	classID, ok := spec.New(build).ClassID(classSlug)
	if !ok {
		return ""
	}
	treeList := build.Trees(classID)
	if len(treeList) != len(points) {
		return ""
	}
	best, most, ties := -1, 0, 0
	for i, p := range points {
		switch {
		case p > most:
			best, most, ties = i, p, 1
		case p == most && p > 0:
			ties++
		}
	}
	if best < 0 || most == 0 || ties > 1 {
		return ""
	}
	return treeList[best].Name
}
