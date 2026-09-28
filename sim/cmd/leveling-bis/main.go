// Command leveling-bis is lane bis-proto's prototype: for one spec, at
// a handful of level bands and both factions, it ranks every eligible
// item per slot by the spec's own stat weights, picks the best per
// slot, verifies the pick against its runner-up with a real sim, and
// writes the result as JSON and a readable markdown table.
//
// Usage (from the repository root):
//
//	go run ./sim/cmd/leveling-bis -spec hunter-marksmanship -bands 20,30,40,60
//
// See docs/superpowers/specs/2026-09-28-leveling-bis-design.md for the
// design this prototypes and this lane's own brief (repeated in the
// lane report) for the exact rules implemented.
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
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("leveling-bis: %v", err)
	}
}

func run() error {
	repoRoot := flag.String("repo-root", ".", "the site repository root (data/curated/specs.json must be under it)")
	spec := flag.String("spec", "hunter-marksmanship", "the spec to rank (data/curated/specs.json's spec slug)")
	bandsFlag := flag.String("bands", "20,30,40,60", "comma-separated level bands")
	build := flag.String("build", "", "data build to read from data/builds/<build>; defaults to web/src/data/active-build.json's build")
	weightsIterations := flag.Int("weights-iterations", 100, "iterations PER DIRECTION the weights sweep runs (multiplied by the engine's own WeightsIterationsFactor); the brief allows reducing this to stay under the time budget")
	out := flag.String("out", "", "output directory; defaults to sim/cmd/leveling-bis/out under -repo-root")
	flag.Parse()

	if _, err := os.Stat(filepath.Join(*repoRoot, "data", "curated", "specs.json")); err != nil {
		return fmt.Errorf("-repo-root %q does not look like the site repository (data/curated/specs.json not found): %w", *repoRoot, err)
	}
	outDir := *out
	if outDir == "" {
		outDir = filepath.Join(*repoRoot, "sim", "cmd", "leveling-bis", "out")
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

	bands, err := parseBands(*bandsFlag)
	if err != nil {
		return err
	}

	specInfo, err := loadSpec(*repoRoot, *spec)
	if err != nil {
		return err
	}
	guide, err := loadGuideBuild(*repoRoot, specInfo.ClassSlug, specInfo.SpecSlug)
	if err != nil {
		return err
	}
	items, missing, err := loadCandidates(buildDir, specInfo.ClassSlug)
	if err != nil {
		return err
	}
	for _, m := range missing {
		log.Printf("leveling-bis: %s", m)
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
	overallStart := time.Now()
	for _, band := range bands {
		trees := guide.Trees
		budget := levelBudget(band)
		talentPoints := budget
		if talentPoints < 0 {
			talentPoints = 0
		}
		// truncateTalents is a no-op once budget already covers every
		// point the build spends (band 60's 51), so it is always safe
		// to call rather than special-cased per band.
		talents := talentString(truncateTalents(trees, budget))

		weapon := ladderWeapon(items, band)
		ladderCh := ladderCharacter(guide.AllianceRace, specInfo.ClassSlug, band, talents, weapon)

		weightsStart := time.Now()
		wreq := weightsRequest(specInfo, ladderCh, *weightsIterations, 3)
		wresult, err := runWeights(wreq)
		if err != nil {
			return fmt.Errorf("band %d weights run: %w", band, err)
		}
		weightsSeconds := time.Since(weightsStart).Seconds()
		weights := make(map[string]float64, len(wresult))
		for stat, w := range wresult {
			weights[stat] = w.Weight
		}
		log.Printf("leveling-bis: band %d weights (%.1fs): %s", band, weightsSeconds, formatWeights(specInfo.WeightStats, weights))

		for _, f := range factions {
			pool := buildBandPool(items, lootIdx, specInfo.ClassSlug, band, f.name, weights)
			bySlot := candidatesBySlot(pool.Scored)
			picks := pick(bySlot)

			verifyStart := time.Now()
			setDPS, swaps, verifyErrors, err := verifyBand(specInfo, f.race, specInfo.ClassSlug, band, picks)
			if err != nil {
				return fmt.Errorf("band %d %s verify run (baseline): %w", band, f.name, err)
			}
			verifySeconds := time.Since(verifyStart).Seconds()
			for _, e := range verifyErrors {
				log.Printf("leveling-bis: band %d %s: could not verify %s", band, f.name, e)
			}

			report := buildReport(specInfo, band, f.name, f.race, talents, talentPoints, weights, specInfo.WeightStats, picks, setDPS, swaps, pool.NoSource, previous[f.name], weightsSeconds, verifySeconds, verifyErrors)
			reports = append(reports, report)
			previous[f.name] = picks

			log.Printf("leveling-bis: band %d %s: set DPS %.1f, verify %.1fs, %d no-source, %d verify errors", band, f.name, setDPS, verifySeconds, len(pool.NoSource), len(verifyErrors))
		}
	}
	log.Printf("leveling-bis: total run time %.1fs", time.Since(overallStart).Seconds())

	jsonPath := filepath.Join(outDir, *spec+".json")
	if err := writeJSON(jsonPath, reports); err != nil {
		return fmt.Errorf("writing %s: %w", jsonPath, err)
	}
	mdPath := filepath.Join(outDir, *spec+".md")
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
