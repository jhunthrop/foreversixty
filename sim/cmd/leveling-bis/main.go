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
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/leveling"
)

func main() {
	if err := run(os.Args[0], os.Args[1:]); err != nil {
		log.Fatalf("leveling-bis: %v", err)
	}
}

// defaultBands is the design doc's own leveling-bis band list: 10
// through 60, step 5 - "the best you can wear at level L", one list
// per rung a leveling character actually stops on to gear up.
const defaultBandsFlag = "10,15,20,25,30,35,40,45,50,55,60"

// run takes execPath and args explicitly (main passes os.Args[0] and
// os.Args[1:]) rather than reading the process's own os.Args and the
// package-global flag.CommandLine directly: a dedicated flag.FlagSet
// per call is what lets a test invoke run more than once in the same
// process (flag.CommandLine is shared package state - a second
// flag.String("repo-root", ...) against it panics with "flag
// redefined") - this codebase's own rule, explicit dependencies over
// globals, applied to the one dependency run() has that a fake could
// not otherwise replace: which binary runAllSpecsIsolated re-execs.
func run(execPath string, args []string) error {
	fs := flag.NewFlagSet("leveling-bis", flag.ContinueOnError)
	repoRoot := fs.String("repo-root", ".", "the site repository root (data/curated/specs.json must be under it)")
	spec := fs.String("spec", "hunter-marksmanship", "the spec to rank (data/curated/specs.json's spec slug); ignored when -all is set")
	all := fs.Bool("all", false, "rank every spec in data/curated/specs.json with a written rotation (data/curated/apl/<spec>.json state == \"written\"), one output file per spec - what make bis and the nightly workflow run")
	bandsFlag := fs.String("bands", defaultBandsFlag, "comma-separated level bands")
	build := fs.String("build", "", "data build to read from data/builds/<build>; defaults to web/src/data/active-build.json's build")
	weightsIterations := fs.Int("weights-iterations", 100, "iterations PER DIRECTION the weights sweep runs (multiplied by the engine's own WeightsIterationsFactor); the brief allows reducing this to stay under the time budget")
	out := fs.String("out", "", "output directory; defaults to data/builds/<build>/bis under -repo-root (lane bis-web's read contract)")
	memProfile := fs.String("memprofile", "", "write a heap profile to this path (a per-spec suffix is added under -all: <path>.<spec>) - diagnostic only, go tool pprof -top <file>")
	if err := fs.Parse(args); err != nil {
		return err
	}

	if *memProfile != "" && !*all {
		defer writeHeapProfile(*memProfile)
	}

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

	// -all runs each spec in its OWN subprocess (runSpecSubprocess)
	// rather than looping runSpec in this process. This is a direct
	// fix for a real incident: a full -all run (20 written specs, the
	// default 11 bands, both factions) was killed after 1h52min at
	// ~35GB RSS, starving every other lane on the machine, against
	// this brief's 30-minute/well-under-memory budget. Bisecting it
	// (see the lane report) found no single pathological spec - every
	// spec sampled alone finished in under a minute at a few hundred
	// MB RSS - which points at slow, sustained growth ACROSS a single
	// long-lived process issuing many thousands of one-shot engine
	// sims (weights + verify + trinket-rank, per band, per faction,
	// per spec), not a bug in any one spec's own data. A subprocess
	// per spec makes each spec's peak memory the WHOLE process's peak
	// memory - the OS reclaims everything the moment that subprocess
	// exits, the same guarantee an explicit in-process GC/
	// FreeOSMemory call cannot make if something really is being held
	// reachable across specs. It also gives each spec a hard wall-
	// clock ceiling (specTimeout) so one hung spec cannot silently
	// re-create the same incident.
	if *all {
		return runAllSpecsIsolated(execPath, specTimeout, *repoRoot, activeBuild, outDir, *bandsFlag, *weightsIterations, *memProfile)
	}
	return runSpec(realEngine{}, *repoRoot, buildDir, activeBuild, outDir, *spec, bands, *weightsIterations)
}

// specTimeout bounds one spec's subprocess: generous next to every
// measured single-spec run in this lane's report (under a minute
// each), but short enough that a hang is caught and reported rather
// than repeating the incident this function's caller documents.
const specTimeout = 5 * time.Minute

// runAllSpecsIsolated runs writtenSpecs, one subprocess per spec, and
// reports which (if any) failed or hung - see run()'s own doc for why
// this is a subprocess loop and not an in-process one. execPath and
// timeout are run()'s own explicit dependencies threaded one level
// further (execPath is normally os.Args[0]; timeout is normally
// specTimeout) so a test can point both at a fast, scripted stand-in
// process instead of re-execing the real, slow ranking binary.
func runAllSpecsIsolated(execPath string, timeout time.Duration, repoRoot, activeBuild, outDir, bandsFlag string, weightsIterations int, memProfile string) error {
	specs, err := writtenSpecs(repoRoot)
	if err != nil {
		return err
	}

	overallStart := time.Now()
	var failed []string
	for _, s := range specs {
		specStart := time.Now()
		if err := runSpecSubprocess(execPath, timeout, repoRoot, activeBuild, outDir, s, bandsFlag, weightsIterations, memProfile); err != nil {
			// -all is a nightly batch of independent units of work - a
			// mage bug returning no engine data this run genuinely
			// cannot rank should not cost every OTHER spec its BiS
			// list too (a real failure mode this lane's own -all dry
			// run hit: hunter-survival's reference stat, character.go's
			// referenceStatOverride doc). Log it, keep going, and fail
			// the whole run at the end if anything did not make it -
			// visible in CI, but never at the cost of the specs that
			// succeeded.
			log.Printf("leveling-bis: spec %s FAILED, skipping: %v", s, err)
			failed = append(failed, s)
			continue
		}
		log.Printf("leveling-bis: spec %s done in %.1fs", s, time.Since(specStart).Seconds())
	}
	log.Printf("leveling-bis: %d spec(s) attempted, %d failed, total run time %.1fs", len(specs), len(failed), time.Since(overallStart).Seconds())
	if len(failed) > 0 {
		return fmt.Errorf("%d of %d spec(s) failed: %s", len(failed), len(specs), strings.Join(failed, ", "))
	}
	return nil
}

// runSpecSubprocess re-execs execPath (run()'s own os.Args[0] - a real
// executable in both `go run` (go run builds one to a temp path first)
// and a built binary, so this works identically in dev and in `make
// bis`/the nightly workflow) for exactly one spec, forwarding its
// stdout/stderr live so the parent's log stays one continuous stream.
// -all is deliberately NOT forwarded (this call always names -spec),
// which is what keeps this from recursing.
func runSpecSubprocess(execPath string, timeout time.Duration, repoRoot, build, outDir, spec, bandsFlag string, weightsIterations int, memProfile string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	args := []string{
		"-repo-root", repoRoot,
		"-spec", spec,
		"-bands", bandsFlag,
		"-build", build,
		"-out", outDir,
		"-weights-iterations", strconv.Itoa(weightsIterations),
	}
	if memProfile != "" {
		args = append(args, "-memprofile", memProfile+"."+spec)
	}
	cmd := exec.CommandContext(ctx, execPath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// The spec runs in its own process group so a timeout kills every
	// descendant, not just the direct child: a grandchild left holding
	// stdout kept CI's test binary waiting a full minute ("Test I/O
	// incomplete 1m0s after exiting"). WaitDelay bounds that wait for
	// anything the group kill still misses.
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Cancel = func() error { return syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL) }
	cmd.WaitDelay = 5 * time.Second
	err := cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		return fmt.Errorf("spec %s did not finish within %s (killed) - the incident this lane's report documents; this is the safety net, not the fix", spec, timeout)
	}
	return err
}

// runSpec ranks one spec across every band and both factions and writes
// its two output files (json, md) under outDir.
func runSpec(runner engineRunner, repoRoot, buildDir, activeBuild, outDir, spec string, bands []int, weightsIterations int) error {
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
	// itemFactionRestriction: item id -> its own client-stated
	// faction_restriction, for correctedRepSource's general check
	// (data.go's own doc: a mined rep source's Side is wrong for a real
	// handful of WSG honored-tier items, and the item's own hard
	// restriction outranks it whenever they disagree).
	itemFactionRestriction := make(map[int]string, len(items))
	for _, c := range items {
		if c.FactionRestriction != "" {
			itemFactionRestriction[c.ID] = c.FactionRestriction
		}
	}
	lootIdx, questFloors, err := loadLootIndex(buildDir, itemFactionRestriction)
	if err != nil {
		return err
	}
	// 2026-09-28 quest-levels lane: resolve each candidate's REAL level
	// gate (quest min_level / crafted item-level proxy, not just the
	// item's own required_level, which a quest or crafted reward almost
	// always states as 0) once, before any band uses eligible().
	items = applyEffectiveRequiredLevels(items, lootIdx, questFloors)

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
		wresult, err := runner.RunWeights(wreq)
		if err != nil {
			return fmt.Errorf("band %d weights run: %w", band, err)
		}
		weightsSeconds := time.Since(weightsStart).Seconds()
		// buildBandPool/score() only ever need the plain number (a
		// candidate's stats dotted against it), and only for a weight
		// this command's own significance bar trusts -- effectiveWeights
		// (weights.go) zeros the rest, so an insignificant (possibly
		// negative) weight cannot move a ranking. wresult itself (with
		// Error and Insignificant) still rides through to buildReport
		// unchanged, so the published JSON keeps publishing what
		// isWeightSignificant says about each one instead of a bare,
		// unqualified number.
		weights := effectiveWeights(wresult)
		log.Printf("leveling-bis: %s band %d weights (%.1fs): %s", spec, band, weightsSeconds, formatWeights(specInfo.WeightStats, wresult))

		for _, f := range factions {
			pool := buildBandPool(items, lootIdx, specInfo.ClassSlug, band, f.name, weights)
			bySlot := candidatesBySlot(pool.Scored)
			picks := pick(spec, bySlot)

			// Trinkets carry no scorable stats (score.go's own doc), so
			// pick()'s score-based choice for trinket1/trinket2 is
			// really just "lowest item id" - replace it with an
			// engine-verified ranking of the top item-level candidates
			// (trinkets.go; this lane's brief). trinket1 first so
			// trinket2's own ranking sees trinket1's final pick, not
			// its score-based placeholder.
			trinketStart := time.Now()
			for _, slot := range []string{"trinket1", "trinket2"} {
				var notes []string
				picks, notes = rankTrinketSlot(runner, specInfo, f.race, specInfo.ClassSlug, band, talents, picks, bySlot, slot)
				for _, n := range notes {
					log.Printf("leveling-bis: %s band %d %s: %s", spec, band, f.name, n)
				}
			}
			trinketSeconds := time.Since(trinketStart).Seconds()

			// Every other slot with an engine-implemented effect
			// candidate (rank.go; this lane's brief, item 3): score()
			// cannot see a proc at all, so a slot score() would
			// otherwise decide on stats alone gets a real verify pass
			// against its own implemented-effect candidates.
			effectStart := time.Now()
			for _, slot := range slotsNeedingEffectVerification(bySlot) {
				var notes []string
				picks, notes = rankSlotWithEffects(runner, specInfo, f.race, specInfo.ClassSlug, band, talents, picks, bySlot, slot)
				for _, n := range notes {
					log.Printf("leveling-bis: %s band %d %s: %s", spec, band, f.name, n)
				}
			}

			// A pick that would complete an engine-implemented 2- or
			// 3-piece set is tried together and kept only if it
			// verifies ahead of the independently-scored picks (sets.go;
			// this lane's brief, item 3's second half).
			var setNotes []string
			picks, setNotes = trySetCompletion(runner, specInfo, f.race, specInfo.ClassSlug, band, talents, picks, bySlot)
			for _, n := range setNotes {
				log.Printf("leveling-bis: %s band %d %s: %s", spec, band, f.name, n)
			}
			effectSeconds := time.Since(effectStart).Seconds()

			// Re-assert pick()'s own two-hand/off-hand rule: either of
			// the two passes just above can replace main_hand's pick
			// with a two-hander without knowing off_hand exists (see
			// pick.go's enforceTwoHandOffHandInvariant doc - this
			// lane's report names every spec it found the gap on).
			picks = enforceTwoHandOffHandInvariant(picks)

			verifyStart := time.Now()
			setDPS, swaps, verifyErrors, err := verifyBand(runner, specInfo, f.race, specInfo.ClassSlug, band, talents, picks)
			if err != nil {
				return fmt.Errorf("band %d %s verify run (baseline): %w", band, f.name, err)
			}
			for _, e := range verifyErrors {
				log.Printf("leveling-bis: %s band %d %s: could not verify %s", spec, band, f.name, e)
			}
			// A runner-up the sim measured ahead of the scored pick IS the
			// pick: swap it into the slot and re-measure the whole set once,
			// so the published row, the set DPS and the next band's diff all
			// name the item a player should actually wear.
			picks, setDPS, swaps, err = applySwaps(runner, specInfo, f.race, specInfo.ClassSlug, band, talents, picks, swaps, setDPS)
			if err != nil {
				return fmt.Errorf("band %d %s verify run (after swaps): %w", band, f.name, err)
			}
			verifySeconds := time.Since(verifyStart).Seconds()

			report := buildReport(specInfo, band, f.name, f.race, talents, talentPoints, wresult, specInfo.WeightStats, picks, setDPS, swaps, pool.NoSource, previous[f.name], weightsSeconds, verifySeconds, verifyErrors, pool.Coverage)
			reports = append(reports, report)
			previous[f.name] = picks

			// This is the per-spec/band/faction breakdown the controller
			// asked for after the memory incident: weights (once per
			// band, logged above), trinket-rank and verify seconds
			// separately per faction, so a slow band/spec is visible
			// without re-deriving it from timestamps.
			log.Printf("leveling-bis: %s band %d %s: set DPS %.1f, trinket-rank %.1fs, effect-rank+set-completion %.1fs, verify %.1fs, %d no-source, %d cross-class set item(s) excluded, %d weapon candidate(s) with no dps (lane data-weapons' gap), %d verify errors", spec, band, f.name, setDPS, trinketSeconds, effectSeconds, verifySeconds, len(pool.NoSource), len(pool.CrossClassSet), len(pool.NoDPSWeapon), len(verifyErrors))
			// lane rank-guardrails, guardrail A: one line per band+faction
			// naming how much of the slot table a reader is actually
			// looking at versus how much the ranker could see at all -
			// the same coverage the published JSON's own Coverage field
			// carries (report.go's coverageSummary), so a nightly log
			// reader sees the honesty gap without opening the JSON.
			log.Printf("leveling-bis: %s band %d %s: coverage %s", spec, band, f.name, coverageSummary(pool.Coverage))
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

// formatWeights renders the log line a nightly run's own console shows
// per band: every weight with its ± error, and "not significant" for
// one report.go's isWeightSignificant would grey out on the page - so
// a reader watching the run does not have to open the JSON to see the
// same honesty the page shows.
func formatWeights(order []string, weights map[string]api.StatWeight) string {
	parts := make([]string, len(order))
	for i, id := range order {
		w := weights[id]
		if !isWeightSignificant(w) {
			parts[i] = fmt.Sprintf("%s=not significant (%.3f ± %.3f)", id, w.Weight, w.Error)
			continue
		}
		parts[i] = fmt.Sprintf("%s=%.3f ± %.3f", id, w.Weight, w.Error)
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

// writeHeapProfile writes a pprof heap snapshot to path - `go tool
// pprof -top <path>` afterward ranks what is still reachable when the
// process exits (deferred from run(), so this fires on both a clean
// finish and the early return -all's per-spec failure path can still
// take). It is diagnostic only: nightly runs do not pass -memprofile,
// and a failure to write one is logged, not fatal - losing a profile
// should never be why an otherwise-successful ranking run reports
// itself as failed.
func writeHeapProfile(path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Printf("leveling-bis: -memprofile: creating %s: %v", path, err)
		return
	}
	defer f.Close()
	runtime.GC()
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Printf("leveling-bis: -memprofile: writing %s: %v", path, err)
	}
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
