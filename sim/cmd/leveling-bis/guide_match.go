package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// The guide's front matter is the one source of truth for what a spec's
// BiS bands wear: recommendedRaces names the race per faction, build the
// leveling talents every leveling band and the bare band 60 spend from,
// raidBuild the talents of the raid preset's band 60. The ranker reads
// them (loadGuideRaces, loadPassTargets); guideExpectation re-derives the
// same facts through the same readers so a published file that no longer
// matches its guide fails loudly instead of publishing a stale race or
// talent string. On 2026-11-05 the published hunter files still said
// Dwarf two days after the guides moved to Night Elf, and Affliction,
// Arcane, Feral, Bear, Arms and BM/MM carried old talent strings.

// publishedBand is the identity of one published band: which band,
// preset and faction it is, and the race and talents it was ranked with.
type publishedBand struct {
	Band    int    `json:"band"`
	Preset  string `json:"preset"`
	Faction string `json:"faction"`
	Race    string `json:"race"`
	Talents string `json:"talents"`
}

// publishedSpecFile is the part of bis/<spec>.json this check reads.
type publishedSpecFile struct {
	Spec  string          `json:"spec"`
	Build string          `json:"build"`
	Bands []publishedBand `json:"bands"`
}

// guideExpectation is what a spec's guide says its bands wear.
type guideExpectation struct {
	spec    specInfo
	races   map[string]string
	trees   []leveling.TalentTree
	targets passTargets
}

func loadGuideExpectation(repoRoot, buildDir, activeBuild, spec string) (guideExpectation, error) {
	info, err := loadSpec(repoRoot, spec)
	if err != nil {
		return guideExpectation{}, err
	}
	races, err := loadGuideRaces(repoRoot, buildDir, info.ClassSlug, info.SpecSlug)
	if err != nil {
		return guideExpectation{}, err
	}
	targets, err := loadPassTargets(repoRoot, activeBuild, info)
	if err != nil {
		return guideExpectation{}, err
	}
	trees, err := leveling.LoadTalentTrees(repoRoot, activeBuild, info.ClassSlug)
	if err != nil {
		return guideExpectation{}, err
	}
	return guideExpectation{
		spec:    info,
		races:   map[string]string{"alliance": races.AllianceRace, "horde": races.HordeRace},
		trees:   trees,
		targets: targets,
	}, nil
}

// talentsFor is the site-layout talent string the ranker publishes for a
// band and preset: the same LadderTalentString call bandTalentStrings makes.
func (g guideExpectation) talentsFor(band int, preset string) string {
	return leveling.LadderTalentString(g.trees, g.targets.forPass(preset), g.spec.TreeIndex, band, leveling.ForSpec(g.spec.Spec))
}

// guideDivergences lists every band whose race or talents differ from
// the guide, one line each; an empty list means the bands match it.
func guideDivergences(g guideExpectation, bands []publishedBand) []string {
	var out []string
	for _, b := range bands {
		where := fmt.Sprintf("%s band %d %s %s", g.spec.Spec, b.Band, b.Preset, b.Faction)
		if want, ok := g.races[b.Faction]; !ok {
			out = append(out, fmt.Sprintf("%s: unknown faction", where))
		} else if b.Race != want {
			out = append(out, fmt.Sprintf("%s: race %q, the guide's recommendedRaces say %q", where, b.Race, want))
		}
		if want := g.talentsFor(b.Band, b.Preset); b.Talents != want {
			out = append(out, fmt.Sprintf("%s: talents %s, the guide's %s spends %s", where, b.Talents, guideBuildKey(b.Preset), want))
		}
	}
	return out
}

func guideBuildKey(preset string) string {
	if preset == presetRaid {
		return "raidBuild"
	}
	return "build"
}

// requireBandsMatchGuide is the ranker's own last check before it writes
// a spec's file: the reports must wear the guide's races and builds.
func requireBandsMatchGuide(repoRoot, buildDir, activeBuild, spec string, reports []bandReport) error {
	g, err := loadGuideExpectation(repoRoot, buildDir, activeBuild, spec)
	if err != nil {
		return err
	}
	bands := make([]publishedBand, 0, len(reports))
	for _, r := range reports {
		bands = append(bands, publishedBand{Band: r.Band, Preset: r.Preset, Faction: r.Faction, Race: r.Race, Talents: r.Talents})
	}
	if diverged := guideDivergences(g, bands); len(diverged) > 0 {
		return fmt.Errorf("%s: %d band(s) diverge from the guide front matter, first: %s", spec, len(diverged), diverged[0])
	}
	return nil
}

// readPublishedSpecFile reads the identity fields of a committed bis file.
func readPublishedSpecFile(path string) (publishedSpecFile, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return publishedSpecFile{}, err
	}
	var f publishedSpecFile
	if err := json.Unmarshal(b, &f); err != nil {
		return publishedSpecFile{}, fmt.Errorf("%s: %w", filepath.Base(path), err)
	}
	return f, nil
}
