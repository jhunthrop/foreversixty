// Command leveling-bis ranks the best leveling gear for a spec, once
// per spec and band, never per character (the design doc's own words:
// docs/superpowers/specs/2026-09-28-leveling-bis-design.md). For every
// band and both factions it ranks every eligible item per slot by the
// spec's own stat weights, picks the best per slot (trinkets by an
// engine-verified top-item-level ranking instead, since they carry no
// scorable stats - trinkets.go), verifies the pick against its
// runner-up with a real sim, and writes the result as
// data/builds/<build>/bis/<spec>.json (lane bis-web's read contract -
// report.go's specReport) plus a readable markdown table for humans.
//
// Usage (from the repository root):
//
//	go run ./sim/cmd/leveling-bis -spec hunter-marksmanship -bands 20,30,40,60
//	go run ./sim/cmd/leveling-bis -all
//
// See docs/superpowers/specs/2026-09-28-leveling-bis-design.md for the
// design this implements and this lane's own brief (repeated in the
// lane report) for the exact rules.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/sim/leveling"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("leveling-bis: %v", err)
	}
}

// defaultBands is the design doc's own leveling-bis band list: 10
// through 60, step 5 - "the best you can wear at level L", one list
// per rung a leveling character actually stops on to gear up.
const defaultBandsFlag = "10,15,20,25,30,35,40,45,50,55,60"

func run() error {
	repoRoot := flag.String("repo-root", ".", "the site repository root (data/curated/specs.json must be under it)")
	spec := flag.String("spec", "hunter-marksmanship", "the spec to rank (data/curated/specs.json's spec slug); ignored when -all is set")
	all := flag.Bool("all", false, "rank every spec in data/curated/specs.json with a written rotation (data/curated/apl/<spec>.json state == \"written\"), one output file per spec - what make bis and the nightly workflow run")
	bandsFlag := flag.String("bands", defaultBandsFlag, "comma-separated level bands")
	build := flag.String("build", "", "data build to read from data/builds/<build>; defaults to web/src/data/active-build.json's build")
	weightsIterations := flag.Int("weights-iterations", 100, "iterations PER DIRECTION the weights sweep runs (multiplied by the engine's own WeightsIterationsFactor); the brief allows reducing this to stay under the time budget")
	out := flag.String("out", "", "output directory; defaults to data/builds/<build>/bis under -repo-root (lane bis-web's read contract)")
	flag.Parse()

	if _, err := os.Stat(filepath.Join(*repoRoot, "data", "curated", "specs.json")); err != nil {
		return fmt.Errorf("-repo-root %q does not look like the site repository (data/curated/specs.json not found): %w", *repoRoot, err)
	}

	activeBuild := *build
	if activeBuild == "" {
		var err error
		activeBuild, err = readActiveBuild(*repoRoot)
		if err != nil {
			return err
		}
	}
	buildDir := filepath.Join(*repoRoot, "data", "builds", activeBuild)

	outDir := *out
	if outDir == "" {
		outDir = filepath.Join(*repoRoot, "data", "builds", activeBuild, "bis")
	}

	bands, err := parseBands(*bandsFlag)
	if err != nil {
		return err
	}

	specs := []string{*spec}
	if *all {
		specs, err = writtenSpecs(*repoRoot)
		if err != nil {
			return err
		}
	}

	overallStart := time.Now()
	for _, s := range specs {
		specStart := time.Now()
		if err := runSpec(*repoRoot, buildDir, activeBuild, outDir, s, bands, *weightsIterations); err != nil {
			return fmt.Errorf("spec %s: %w", s, err)
		}
		log.Printf("leveling-bis: spec %s done in %.1fs", s, time.Since(specStart).Seconds())
	}
	log.Printf("leveling-bis: %d spec(s), total run time %.1fs", len(specs), time.Since(overallStart).Seconds())
	return nil
}

// runSpec ranks one spec across every band and both factions and writes
// its two output files (json, md) under outDir.
func runSpec(repoRoot, buildDir, activeBuild, outDir, spec string, bands []int, weightsIterations int) error {
	specInfo, err := loadSpec(repoRoot, spec)
	if err != nil {
		return err
	}
	guide, err := loadGuideRaces(repoRoot, specInfo.ClassSlug, specInfo.SpecSlug)
	if err != nil {
		return err
	}
	// The shared talent-truncation rule (sim/leveling, moved verbatim
	// out of sim/request/ladder.go by this lane): the guide's level-60
	// build, read by stable talent id, re-resolved against the
	// ACTIVE build's own trees so a talent the active build no longer
	// carries is dropped rather than silently misaligning every digit
	// after it (sim/leveling's own GuideTalentTargets doc explains why
	// - the same paladin build-drift finding sim/request/ladder_test.go
	// guards against). One read per spec, not per band: the trees and
	// targets do not change across bands, only how many points
	// LadderTalentString spends from them.
	guideBuild, treeDigits, err := leveling.GuideBuildTalents(repoRoot, specInfo.ClassSlug, specInfo.SpecSlug)
	if err != nil {
		return err
	}
	guideTrees, err := leveling.LoadTalentTrees(repoRoot, guideBuild, specInfo.ClassSlug)
	if err != nil {
		return err
	}
	activeTrees, err := leveling.LoadTalentTrees(repoRoot, activeBuild, specInfo.ClassSlug)
	if err != nil {
		return err
	}
	talentTargets := leveling.GuideTalentTargets(guideTrees, treeDigits)

	items, missing, err := loadCandidates(buildDir, specInfo.ClassSlug)
	if err != nil {
		return err
	}
	for _, m := range missing {
		log.Printf("leveling-bis: %s: %s", spec, m)
	}
	lootIdx, err := loadLootIndex(buildDir)
	if err != nil {
		return err
	}

	factions := []struct{ name, race string }{
		{"alliance", guide.AllianceRace},
		{"horde", guide.HordeRace},
	}
	previous := map[string]map[string]slotPick{"alliance": nil, "horde": nil}

	var reports []bandReport
	for _, band := range bands {
		talents := leveling.LadderTalentString(activeTrees, talentTargets, specInfo.TreeIndex, band)
		talentPoints := talentPointsSpent(talents)

		weapon := ladderWeapon(items, band)
		ladderCh := ladderCharacter(guide.AllianceRace, specInfo.ClassSlug, band, talents, weapon)

		weightsStart := time.Now()
		wreq := weightsRequest(specInfo, ladderCh, weightsIterations, 3)
		wresult, err := runWeights(wreq)
		if err != nil {
			return fmt.Errorf("band %d weights run: %w", band, err)
		}
		weightsSeconds := time.Since(weightsStart).Seconds()
		weights := make(map[string]float64, len(wresult))
		for stat, w := range wresult {
			weights[stat] = w.Weight
		}
		log.Printf("leveling-bis: %s band %d weights (%.1fs): %s", spec, band, weightsSeconds, formatWeights(specInfo.WeightStats, weights))

		for _, f := range factions {
			pool := buildBandPool(items, lootIdx, specInfo.ClassSlug, band, f.name, weights)
			bySlot := candidatesBySlot(pool.Scored)
			picks := pick(bySlot)

			// Trinkets carry no scorable stats (score.go's own doc), so
			// pick()'s score-based choice for trinket1/trinket2 is
			// really just "lowest item id" - replace it with an
			// engine-verified ranking of the top item-level candidates
			// (trinkets.go; this lane's brief). trinket1 first so
			// trinket2's own ranking sees trinket1's final pick, not
			// its score-based placeholder.
			for _, slot := range []string{"trinket1", "trinket2"} {
				var notes []string
				picks, notes = rankTrinketSlot(specInfo, f.race, specInfo.ClassSlug, band, picks, bySlot, slot)
				for _, n := range notes {
					log.Printf("leveling-bis: %s band %d %s: %s", spec, band, f.name, n)
				}
			}

			verifyStart := time.Now()
			setDPS, swaps, verifyErrors, err := verifyBand(specInfo, f.race, specInfo.ClassSlug, band, picks)
			if err != nil {
				return fmt.Errorf("band %d %s verify run (baseline): %w", band, f.name, err)
			}
			verifySeconds := time.Since(verifyStart).Seconds()
			for _, e := range verifyErrors {
				log.Printf("leveling-bis: %s band %d %s: could not verify %s", spec, band, f.name, e)
			}

			report := buildReport(specInfo, band, f.name, f.race, talents, talentPoints, weights, specInfo.WeightStats, picks, setDPS, swaps, pool.NoSource, previous[f.name], weightsSeconds, verifySeconds, verifyErrors)
			reports = append(reports, report)
			previous[f.name] = picks

			log.Printf("leveling-bis: %s band %d %s: set DPS %.1f, verify %.1fs, %d no-source, %d cross-class set item(s) excluded, %d weapon candidate(s) with no dps (lane data-weapons' gap), %d verify errors", spec, band, f.name, setDPS, verifySeconds, len(pool.NoSource), len(pool.CrossClassSet), len(pool.NoDPSWeapon), len(verifyErrors))
		}
	}

	jsonPath := filepath.Join(outDir, spec+".json")
	if err := writeSpecReport(jsonPath, spec, activeBuild, reports); err != nil {
		return fmt.Errorf("writing %s: %w", jsonPath, err)
	}
	mdPath := filepath.Join(outDir, spec+".md")
	if err := writeMarkdown(mdPath, specInfo, reports); err != nil {
		return fmt.Errorf("writing %s: %w", mdPath, err)
	}
	log.Printf("leveling-bis: wrote %s and %s", jsonPath, mdPath)
	return nil
}

func formatWeights(order []string, weights map[string]float64) string {
	parts := make([]string, len(order))
	for i, id := range order {
		parts[i] = fmt.Sprintf("%s=%.3f", id, weights[id])
	}
	return strings.Join(parts, ", ")
}

// talentPointsSpent sums a leveling.LadderTalentString result's own
// digits rather than recomputing level-9: LadderTalentString may spend
// fewer points than the budget allows (its own doc: "If the guide
// build itself spends fewer than level-9 points... the remainder is
// left unspent") - reading the count back off the string it actually
// produced reports what was truly spent, with no second formula that
// could drift from the first.
func talentPointsSpent(talents string) int {
	total := 0
	for _, r := range talents {
		if r >= '0' && r <= '9' {
			total += int(r - '0')
		}
	}
	return total
}

func parseBands(s string) ([]int, error) {
	parts := strings.Split(s, ",")
	out := make([]int, 0, len(parts))
	for _, p := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			return nil, fmt.Errorf("-bands: %q is not a number: %w", p, err)
		}
		out = append(out, n)
	}
	return out, nil
}

func readActiveBuild(repoRoot string) (string, error) {
	type activeBuildFile struct {
		Build string `json:"build"`
	}
	path := filepath.Join(repoRoot, "web", "src", "data", "active-build.json")
	b, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("reading %s: %w", path, err)
	}
	var f activeBuildFile
	if err := json.Unmarshal(b, &f); err != nil {
		return "", fmt.Errorf("decoding %s: %w", path, err)
	}
	if f.Build == "" {
		return "", fmt.Errorf("%s carries no build", path)
	}
	return f.Build, nil
}
