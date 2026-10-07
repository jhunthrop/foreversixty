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
	"flag"
	"fmt"
	"log"
	"maps"
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
	"github.com/jhunthrop/foreversixty/sim/internal/enginetalents"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/request"
)

func main() {
	if err := run(os.Args[0], os.Args[1:]); err != nil {
		log.Fatalf("leveling-bis: %v", err)
	}
}

// defaultBandsFlag is the leveling-bis band list: 20 through 60, step
// 10 - "the best you can wear at level L", one list per rung a leveling
// character actually stops on to gear up. Owner ruling 2026-09-29: five
// bands, not the earlier 10..60 step 5 (a list every five levels was
// noise, and nothing below 20 is worth gearing for).
const defaultBandsFlag = "20,30,40,50,60"

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
		activeBuild, err = leveling.ReadActiveBuild(*repoRoot)
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
	return runSpec(realEngine{}, *repoRoot, buildDir, activeBuild, outDir, *spec, bands, *weightsIterations, resolveTalentLayout)
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

// factionWork is one faction's own working state for one band, carried
// from the first half of runSpec's own per-faction loop (buildBandPool
// through rankTrinketSlot) to reconcileFactionTrinkets
// (faction_trinkets.go, this lane's brief bis-ranker-integrity-15) and
// on into the loop's second half (rankSlotWithEffects onward) - split
// into two passes so both factions' own trinket tournaments are
// finished, and reconcilable against each other, before either one's
// picks continue into the rest of the pipeline.
type factionWork struct {
	faction        string
	race           string
	pool           bandPool
	bySlot         map[string][]scored
	pickBySlot     map[string][]scored
	picks          map[string]slotPick
	trinketSeconds float64
}

// goodWeights is a band's trusted weights, kept for the next band's
// fallback (see runSpec's lastGood).
type goodWeights struct {
	weights              map[string]float64
	referenceDPSPerPoint float64
	band                 int
}

// talentLayout is enginetalents.Layout's own Reposition method,
// abstracted the same way engineRunner abstracts the real engine: a
// test supplies a layout matching its own synthetic talent ids
// (identityTalentLayout, testhelpers_test.go) instead of resolving,
// and validating every talent against, the real compiled engine.
type talentLayout interface {
	Reposition(trees []leveling.TalentTree, s string) (string, error)
}

// talentLayoutResolver resolves a class's talentLayout given the site
// repository root - runSpec's own explicit dependency for where that
// layout comes from, the same role engineRunner plays for where a sim
// result comes from.
type talentLayoutResolver func(repoRoot, class string) (talentLayout, error)

// resolveTalentLayout is run()'s production talentLayoutResolver: the
// compiled engine's own talent-string layout for class, read from its
// proto source under repoRoot/sim (sim/internal/enginetalents' own
// doc).
func resolveTalentLayout(repoRoot, class string) (talentLayout, error) {
	engineDir, err := enginetalents.SourceDir(filepath.Join(repoRoot, "sim"))
	if err != nil {
		return nil, err
	}
	layout, err := enginetalents.ForClass(engineDir, class)
	if err != nil {
		return nil, err
	}
	return layout, nil
}

// runSpec ranks one spec across every band and both factions and writes
// its two output files (json, md) under outDir.
func runSpec(runner engineRunner, repoRoot, buildDir, activeBuild, outDir, spec string, bands []int, weightsIterations int, resolveLayout talentLayoutResolver) error {
	specInfo, err := loadSpec(repoRoot, spec)
	if err != nil {
		return err
	}
	raidPreset, err := resolveRaidPreset(repoRoot, specInfo)
	if err != nil {
		return fmt.Errorf("resolving the %s preset for %s: %w", presetRaid, spec, err)
	}
	guide, err := loadGuideRaces(repoRoot, buildDir, specInfo.ClassSlug, specInfo.SpecSlug)
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
	// GuideTalentTargets (below) maps treeDigits onto talent ids by
	// walking guideBuild's own trees POSITIONALLY; that read is only
	// right if guideBuild's own tree shape is the one treeDigits was
	// actually authored against. Every guide's stamp briefly said
	// otherwise (lane guide-codes-70009's own finding - the digits
	// were always in 1.60.1.70009's order, every stamp just said
	// 1.60.1.69893), which would have silently misaligned a digit for
	// any class whose tree shape moved between the two builds. This
	// guard is the same invariant RequireGuideBuildMatchesActive's own
	// doc already explains, checked before a mismatch can reach
	// GuideTalentTargets at all.
	if err := leveling.RequireGuideBuildMatchesActive(guideBuild, activeBuild); err != nil {
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

	// engineLayout: the COMPILED engine's own talent string layout for
	// this class, by stable talent id (sim/internal/enginetalents' own
	// doc). The engine's proto was last regenerated from a client build
	// that is not always the site's active one - paladin and shaman
	// both drift - so every talent string this spec hands the engine
	// below is written in engineLayout's own field order, never
	// activeTrees' positional order directly. resolveLayout is an
	// explicit dependency (like runner, above) rather than a direct
	// enginetalents.SourceDir/ForClass call, so a test can supply a
	// layout for its own synthetic talent ids instead of the real
	// compiled engine's (testhelpers_test.go's identityTalentLayout).
	engineLayout, err := resolveLayout(repoRoot, specInfo.ClassSlug)
	if err != nil {
		return err
	}

	items, missing, err := loadCandidates(buildDir, specInfo.ClassSlug)
	if err != nil {
		return err
	}
	for _, m := range missing {
		log.Printf("leveling-bis: %s: %s", spec, m)
	}
	// rating-units lane: items/<class>.json states hit/crit/dodge/
	// parry/block/defense as the RATING number the client's own
	// tooltip shows (ItemModType 31/32/12-15), but every weight
	// score() (score.go) dots a candidate's stats against was measured
	// by the weights sweep per SIM UNIT (percent) - the same unit
	// data/pipeline/simdb/ratings.py already divides item/enchant
	// stats into before they reach simdb.bin. Loaded once per spec run
	// (this build's own gametables/combatratings.txt level-60 row
	// never changes within a run) and applied to every candidate right
	// here, before any band's score() ever sees one, so every
	// downstream consumer of a candidate's Stats already agrees with
	// the engine's own units.
	ratingFactorsForBuild, err := loadRatingFactors(buildDir)
	if err != nil {
		return err
	}
	items = convertCandidateRatings(items, ratingFactorsForBuild)
	// This lane's brief, item 1: markNotInSimDB (data.go) reads this
	// build's own embedded item database once per spec, the same way
	// ratingFactorsForBuild just did - every candidate report.go can
	// ever publish as a final pick already carries the flag before any
	// band's own eligible()/pick()/tournament pass runs.
	items = markNotInSimDB(items)
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

	// This lane's brief, item 5's second half: read once per spec (the
	// rotation does not change per band/faction) whether this spec's own
	// APL casts Backstab or Ambush anywhere - weapon_requirements.go's
	// own doc for why that, not a hand-maintained spec list, is what
	// decides whether main_hand/off_hand get restricted to daggers below.
	requiresDagger, err := aplRotationRequiresDagger(repoRoot, spec)
	if err != nil {
		return fmt.Errorf("checking %s's rotation for a dagger requirement: %w", spec, err)
	}

	// Caster sweep, bis-ranker-integrity-4 lane, item 2: read once per
	// spec, the same way requiresDagger is, whether this spec's own
	// rotation casts Shoot at all - score.go's own doc for why a
	// wand's flat DPS may only ever count for a spec whose rotation
	// actually fires it (mage/priest-shadow/warlock do; shaman-
	// elemental and druid-balance do not).
	castsShoot, err := aplRotationCastsShoot(repoRoot, spec)
	if err != nil {
		return fmt.Errorf("checking %s's rotation for a Shoot cast: %w", spec, err)
	}

	// This lane's brief (bis-ranker-integrity-10), item 3: read once
	// per spec, the same way requiresDagger/castsShoot are (the class's
	// weapon proficiency does not change per band/faction) - the
	// Go-side safety net so a data regression that puts an illegal
	// weapon subclass in <class>.json (a paladin axe, a druid polearm)
	// never reaches a published pick, whatever score() thinks of its
	// stats. weaponSubclassesSource is logged once so the nightly log
	// says whether this run read a published table or fell back to
	// weapon_requirements.go's own static one.
	weaponSubclasses, weaponSubclassesSource, err := loadWeaponSubclasses(buildDir, specInfo.ClassSlug)
	if err != nil {
		return fmt.Errorf("loading %s's weapon proficiency: %w", specInfo.ClassSlug, err)
	}
	log.Printf("leveling-bis: %s: weapon proficiency source: %s", spec, weaponSubclassesSource)

	factions := []struct{ name, race string }{
		{"alliance", guide.AllianceRace},
		{"horde", guide.HordeRace},
	}
	previous := map[string]map[string]slotPick{"alliance": nil, "horde": nil}
	// notInSimWarned: this lane's brief, item 1's last sentence - "log
	// every stripped id once per spec run at warning level so the
	// nightly log names them". A single id can carry a SimStatus of
	// notInSimReason on many bands/factions in this same spec run (the
	// same relic is often BiS at several levels in a row); this map
	// dedupes so the nightly log names each such id exactly once per
	// spec, not once per band+faction it happened to win in.
	notInSimWarned := make(map[int]bool)

	// anchorStat/anchorOK: this lane's brief (ranker-weights-anchor),
	// item 1/2 - the weight_stats row id normalizeScaleFactors
	// (weights.go) anchors this spec's published scale-factor table
	// to, resolved once per spec (primaryAnchorStat's own doc,
	// primary_stat.go) rather than once per band, since it never
	// varies by band. anchorOK is false only for a spec this command's
	// own primaryStatBySpec table has no entry for - unreachable in
	// production (TestPrimaryStatCoversEverySpec, primary_stat_test.go)
	// but handled the same way an absent primary row always is: an
	// empty anchorStat, which both normalizeScaleFactors and
	// primaryStatSignificanceCheck treat as "nothing to anchor on,
	// fall back to the pre-existing rule".
	anchorStat, anchorOK := primaryAnchorStat(specInfo)
	if !anchorOK {
		log.Printf("leveling-bis: %s: no primary_stat entry for this spec - scale factors fall back to the largest-significant-weight rule", spec)
		anchorStat = ""
	}

	var reports []bandReport
	// lastGood tracks, per pass (a raid pass never falls back to bare
	// weights), the most recent LOWER band whose own sweep measured its reference
	// stat positive beyond its own error - this lane's brief, item 1's
	// fallback: a band whose own sweep (even re-run at
	// weightsRetryIterationsFactor iterations) still cannot be trusted
	// ranks and verifies its picks against this band's weights instead
	// of an empty map, rather than publishing nine empty slots over a
	// noisy sweep (warlock-destruction band 60's own repro). Bands run
	// in ascending order (defaultBandsFlag's own doc), so "the nearest
	// lower band" is simply the last one stored.
	lastGood := map[string]goodWeights{}
	for _, band := range bands {
		// priorPicks is every faction's picks at the previous band, which
		// both passes of this band diff against; only the bare pass
		// advances it, so the raid pass at level 60 reads the same
		// previous band the bare one does.
		priorPicks := maps.Clone(previous)
		for _, pass := range passesFor(specInfo, band, raidPreset) {
			specInfo := pass.spec
			lg := lastGood[pass.name]
			// talents is the published band string, kept in the site's own
			// (active-build) layout - talentPoints and buildReport's own
			// "talents" field both read it unconverted, since that is the
			// layout the web planner decodes. engineTalents is the SAME
			// build, repositioned onto the compiled engine's own field
			// order (engineLayout, above) - every character this band
			// builds for the engine (ladderCh and every rank/verify/set
			// character below) is spent from engineTalents, never talents,
			// or paladin and shaman misread every talent at and after the
			// first talent whose tree position moved between the engine's
			// proto build and the active one (this lane's own brief).
			talents, engineTalents, err := bandTalentStrings(activeTrees, talentTargets, specInfo.TreeIndex, band, engineLayout)
			if err != nil {
				return fmt.Errorf("band %d: %w", band, err)
			}
			talentPoints := talentPointsSpent(talents)

			weapon := ladderWeapon(items, band)
			ladderCh := ladderCharacter(guide.AllianceRace, specInfo.ClassSlug, band, engineTalents, weapon, ladderMeleeWeapons(items, band, specInfo.Spec)...)

			weightsStart := time.Now()
			wreq := weightsRequest(specInfo, ladderCh, weightsIterations, 3)
			wresult, referenceDPSPerPoint, err := runner.RunWeights(wreq)
			if err != nil {
				return fmt.Errorf("band %d weights run: %w", band, err)
			}
			weightsReason := referenceMeasurementReason(specInfo.ReferenceStat, wresult, referenceDPSPerPoint)
			if weightsReason != "" {
				// This lane's brief, item 1's guard: a band whose reference
				// stat is not positive beyond its own error re-runs the
				// sweep ONCE at weightsRetryIterationsFactor iterations
				// before giving up on it - warlock-destruction band 60's
				// own repro (bis-ranker-integrity-11) found this tightens
				// the raw standard error (±0.1332 at 100 iterations/
				// direction to ±0.0642 at 400) but does not always flip an
				// actually-negative measurement positive, so the fallback
				// below still has to exist for when this retry alone is
				// not enough.
				retryReq := weightsRequest(specInfo, ladderCh, weightsIterations*weightsRetryIterationsFactor, 3)
				retryResult, retryReferenceDPSPerPoint, retryErr := runner.RunWeights(retryReq)
				if retryErr != nil {
					return fmt.Errorf("band %d weights retry run: %w", band, retryErr)
				}
				retryReason := referenceMeasurementReason(specInfo.ReferenceStat, retryResult, retryReferenceDPSPerPoint)
				log.Printf("leveling-bis: %s band %d: sweep at %d iterations/direction was not significant (%s); re-ran at %dx", spec, band, weightsIterations, weightsReason, weightsRetryIterationsFactor)
				wresult, referenceDPSPerPoint, weightsReason = retryResult, retryReferenceDPSPerPoint, retryReason
			}
			// This lane's brief (ranker-weights-anchor), item 3's own
			// guard: separate from the reference-stat retry just above,
			// re-measure the PRIMARY-anchor row once at
			// primaryStatRetryIterationsFactor iterations if it came back
			// insignificant - but only on a band the reference-stat guard
			// already trusts (weightsReason == ""); an already-untrustworthy
			// band publishes no anchor at all regardless (normalizeScaleFactors'
			// own doc), so spending the extra sim time here would be wasted.
			weightsLowConfidence := false
			if weightsReason == "" && anchorStat != "" {
				retried, lowConf, pErr := primaryStatSignificanceCheck(runner, specInfo, ladderCh, anchorStat, weightsIterations, 3, wresult)
				if pErr != nil {
					return fmt.Errorf("band %d primary-stat (%s) weights retry: %w", band, anchorStat, pErr)
				}
				if lowConf {
					log.Printf("leveling-bis: %s band %d: primary stat %s still not significant after a %dx retry; publishing the anchor anyway (weights_low_confidence)", spec, band, anchorStat, primaryStatRetryIterationsFactor)
				}
				wresult = retried
				weightsLowConfidence = lowConf
			}
			// weightsSeconds covers the whole band, including either
			// retry above when one ran - a single per-band number, the
			// same field buildReport has always taken one of.
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
			bandReferenceDPSPerPoint := referenceDPSPerPoint
			if weightsReason != "" {
				if lg.weights != nil {
					// The fallback half of the guard: rank and verify this
					// band's picks against the nearest lower band's own
					// trusted weights instead of an empty map, and say
					// exactly that in the published weights_reason - an
					// empty slot is only ever published when no candidate
					// exists, never because a sweep was noisy (this lane's
					// brief).
					weights = lg.weights
					bandReferenceDPSPerPoint = lg.referenceDPSPerPoint
					weightsReason = fallbackWeightsReason(specInfo.ReferenceStat, wresult, referenceDPSPerPoint, band, lg.band)
				} else {
					// No earlier band to fall back to (this is the lowest
					// band run, or every band so far has been untrustworthy)
					// - the whole sweep is untrustworthy (see
					// referenceMeasurementReason's own doc) and nothing it
					// measured may rank an item or convert a wand's flat
					// DPS into score units for this band, so both feeds a
					// corrupted reference could poison are cleared the
					// same way an unmeasured band always reads: no weight,
					// no reference to convert against.
					weights = map[string]float64{}
					bandReferenceDPSPerPoint = 0
				}
				log.Printf("leveling-bis: %s band %d: %s", spec, band, weightsReason)
			} else {
				lastGood[pass.name] = goodWeights{weights: weights, referenceDPSPerPoint: bandReferenceDPSPerPoint, band: band}
			}
			log.Printf("leveling-bis: %s band %d weights (%.1fs): %s", spec, band, weightsSeconds, formatWeights(specInfo.WeightStats, wresult))

			work := make(map[string]*factionWork, len(factions))
			for _, f := range factions {
				pool := buildBandPool(items, lootIdx, specInfo.ClassSlug, band, f.name, weights, bandReferenceDPSPerPoint, castsShoot)
				bySlot := candidatesBySlot(pool.Scored)
				// This lane's brief (bis-ranker-integrity-10), item 3:
				// applied before the dagger/ranged-type restrictions below
				// (a narrower, spec- or ranged-specific gate), to every
				// weapon slot this class actually equips a weapon in - the
				// general class-legality gate eligible.go's own doc says
				// is otherwise "NOT checked here" at all.
				bySlot["main_hand"] = restrictToProficientWeapons(bySlot["main_hand"], weaponSubclasses)
				bySlot["off_hand"] = restrictToProficientWeapons(bySlot["off_hand"], weaponSubclasses)
				bySlot["ranged"] = restrictToProficientWeapons(bySlot["ranged"], weaponSubclasses)
				if requiresDagger {
					// weapon_requirements.go's own doc: a mace or sword is a
					// real, legally-equippable item this class file already
					// passed (eligible.go delegates weapon proficiency to the
					// per-class file entirely), but this spec's own rotation
					// cannot cast its dagger-only opener/builder without one -
					// restricted here, before pick() or any later pass ever
					// sees either weapon slot, so a dual-wielder's off_hand
					// (pick()'s own case, which merges main_hand's one-handers
					// in) inherits the restriction for free.
					bySlot["main_hand"] = restrictToDaggers(bySlot["main_hand"])
					bySlot["off_hand"] = restrictToDaggers(bySlot["off_hand"])
				}
				// This lane's brief (bis-ranker-integrity-6), item 9:
				// restrictRangedByProficiency's own doc (weapon_requirements.go)
				// - a caster's ranged slot is only ever a real wand, never a
				// thrown weapon or bow score()'s own Shoot fallback cannot
				// tell apart from one today. Applied unconditionally (every
				// classSlug, not gated behind a spec flag the way the dagger
				// restriction is) since every spec of a given class shares
				// the identical ranged-weapon proficiency.
				bySlot["ranged"] = restrictRangedByProficiency(bySlot["ranged"], specInfo.ClassSlug)

				// pickBySlot is bySlot's own candidates, further narrowed for
				// the DECISION passes only (pick(), rankTrinketSlot,
				// rankSlotWithEffects, trySetCompletion) - bySlot itself stays
				// unfiltered because buildReport (below) reads it for
				// buildAlternatives, and a PvP reward above band.go's
				// pvpRankCap must still be able to appear there, labelled by
				// rank, even though it must never be a DEFAULT pick (this
				// lane's brief, item 3). promoteLowValueWeapon additionally
				// reorders a weapon slot whose every candidate scored exactly
				// 0 (pick.go's own doc; this lane's brief, item 1's second
				// half) - reordering only ever changes which zero-scoring
				// candidate wins a tie, so running it on bySlot too would be
				// harmless, but pickBySlot is the one map every decision pass
				// actually reads, so that is the only copy that needs it.
				pickBySlot := make(map[string][]scored, len(bySlot))
				for slot, list := range bySlot {
					filtered := excludeAbovePvpRankCap(list)
					if weaponSlots[slot] {
						filtered = promoteLowValueWeapon(filtered, specInfo.WeightStats)
					}
					pickBySlot[slot] = filtered
				}
				picks := pick(spec, pickBySlot)

				// Trinkets carry no scorable stats (score.go's own doc), so
				// pick()'s score-based choice for trinket1/trinket2 is
				// really just "lowest item id" - replace it with an
				// engine-verified ranking of the top item-level candidates
				// (trinkets.go; this lane's brief). trinket1 first so
				// trinket2's own ranking sees trinket1's final pick, not
				// its score-based placeholder.
				//
				// This lane's brief (bis-ranker-integrity-6), item 5: the
				// comment above only ever protected trinket2's OWN view of
				// trinket1 - it never noticed that trinket1's OWN
				// rankTrinketSlot call (running first) still reads
				// picks["trinket2"] as its own pair-mate to exclude, and at
				// that point picks["trinket2"] is STILL pick()'s bare
				// score()-based placeholder, not a real decision -
				// clearTrinketPlaceholders' own doc (pick.go) has the full
				// repro and reasoning.
				picks = clearTrinketPlaceholders(picks)
				trinketStart := time.Now()
				for _, slot := range []string{"trinket1", "trinket2"} {
					var notes []string
					picks, notes = rankTrinketSlot(runner, specInfo, f.race, specInfo.ClassSlug, band, engineTalents, picks, pickBySlot, slot, weights)
					for _, n := range notes {
						log.Printf("leveling-bis: %s band %d %s: %s", spec, band, f.name, n)
					}
				}
				trinketSeconds := time.Since(trinketStart).Seconds()
				work[f.name] = &factionWork{faction: f.name, race: f.race, pool: pool, bySlot: bySlot, pickBySlot: pickBySlot, picks: picks, trinketSeconds: trinketSeconds}
			}

			// This lane's brief (bis-ranker-integrity-15, twelfth sweep): both
			// factions' own rankTrinketSlot tournaments just above are now
			// finished for this band - reconcile trinket1/trinket2 across
			// them before either faction's picks continue into
			// rankSlotWithEffects/trySetCompletion/verifyBand below, so a
			// faction-neutral item's own published verdict never differs
			// between Alliance and Horde for a reason that is really just
			// one side's own sim noise landing on the wrong side of a shared
			// bar (druid-feral band 50, druid-balance band 50 - this lane's
			// own repro; faction_trinkets.go's own doc has the full design).
			allianceWork, hordeWork := work["alliance"], work["horde"]
			var reconcileNotes []string
			allianceWork.picks, hordeWork.picks, reconcileNotes = reconcileFactionTrinkets(
				runner, specInfo, specInfo.ClassSlug, band, engineTalents, lootIdx,
				factionTrinketInputs{Faction: "alliance", Race: allianceWork.race, BySlot: allianceWork.bySlot}, allianceWork.picks,
				factionTrinketInputs{Faction: "horde", Race: hordeWork.race, BySlot: hordeWork.bySlot}, hordeWork.picks,
			)
			for _, n := range reconcileNotes {
				log.Printf("leveling-bis: %s band %d faction reconcile: %s", spec, band, n)
			}

			for _, f := range factions {
				fw := work[f.name]
				pool := fw.pool
				bySlot := fw.bySlot
				pickBySlot := fw.pickBySlot
				picks := fw.picks
				trinketSeconds := fw.trinketSeconds

				// Every other slot with an engine-implemented effect
				// candidate (rank.go; this lane's brief, item 3): score()
				// cannot see a proc at all, so a slot score() would
				// otherwise decide on stats alone gets a real verify pass
				// against its own implemented-effect candidates.
				effectStart := time.Now()
				for _, slot := range slotsNeedingEffectVerification(pickBySlot) {
					var notes []string
					picks, notes = rankSlotWithEffects(runner, specInfo, f.race, specInfo.ClassSlug, band, engineTalents, picks, pickBySlot, slot)
					for _, n := range notes {
						log.Printf("leveling-bis: %s band %d %s: %s", spec, band, f.name, n)
					}
				}

				// A pick that would complete an engine-implemented 2- or
				// 3-piece set is tried together and kept only if it
				// verifies ahead of the independently-scored picks (sets.go;
				// this lane's brief, item 3's second half).
				var setNotes []string
				picks, setNotes = trySetCompletion(runner, specInfo, f.race, specInfo.ClassSlug, band, engineTalents, picks, pickBySlot)
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
				setDPS, swaps, verifyErrors, err := verifyBand(runner, specInfo, f.race, specInfo.ClassSlug, band, engineTalents, picks)
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
				picks, setDPS, swaps, err = applySwaps(runner, specInfo, f.race, specInfo.ClassSlug, band, engineTalents, picks, swaps, setDPS)
				if err != nil {
					return fmt.Errorf("band %d %s verify run (after swaps): %w", band, f.name, err)
				}
				verifySeconds := time.Since(verifyStart).Seconds()

				report := buildReport(specInfo, band, f.name, f.race, talents, talentPoints, wresult, specInfo.WeightStats, picks, setDPS, swaps, pool.NoSource, priorPicks[f.name], weightsSeconds, verifySeconds, verifyErrors, pool.Coverage, bySlot, bandReferenceDPSPerPoint, weightsReason)
				// This lane's brief (ranker-weights-anchor), item 3:
				// weights_low_confidence is this band's own flag (set
				// above, once per band, before the faction loop) - not
				// per-faction data, but published on every faction's own
				// report the same way every other band-level field here is.
				report.WeightsLowConfidence = weightsLowConfidence
				report.Preset = pass.name
				// "No primary stat ever published as 'not significant'"
				// (this lane's brief, item 3) - clears Insignificant on
				// exactly the anchor row, on an otherwise-trustworthy band
				// (see forceAnchorRowSignificant's own doc for why an
				// already-untrustworthy band is excluded). Must run before
				// normalizeScaleFactors below, which trusts this flag when
				// deciding whether the anchor row is even usable as a
				// divisor.
				report.Weights = forceAnchorRowSignificant(report.Weights, anchorStat, weightsReason)
				// This lane's brief, item 3: what the site publishes is per
				// RATING point (what the item's own tooltip shows), not per
				// sim unit (percent) - publishWeightRatingUnits (report.go)
				// converts exactly the rating-family rows, keeping the raw
				// sim-unit weight under weight_per_percent. Applied here,
				// once per band+faction, rather than inside buildReport
				// itself - see that function's own doc for why.
				report.Weights = publishWeightRatingUnits(report.Weights, ratingFactorsForBuild)
				// This lane's brief (bis-weights-simc, extended by
				// ranker-weights-anchor): republish the same rows again,
				// this time in the SimulationCraft/Pawn-familiar
				// scale-factor convention (per point, normalized to the
				// spec's own PRIMARY stat = 1.00, anchorStat) - see
				// normalizeScaleFactors' own doc (weights.go) for why this
				// runs after, not instead of, publishWeightRatingUnits
				// above (it needs the already-converted per-rating-point
				// Weight, not the raw per-percent one).
				report.Weights, report.ScaleReferenceStat = normalizeScaleFactors(report.Weights, report.ReferenceDPSPerPoint, anchorStat)
				// Owner correction, 2026-09-30, after player review: haste
				// is not a table row on the site's own weight rail any
				// more, only a one-line caption built from this one number
				// - see bandReport.HasteScaleFactor's own doc.
				report.HasteScaleFactor = hasteScaleFactorFromRows(report.Weights, report.ScaleReferenceStat)
				// Owner correction, 2026-09-30, after the caption's own
				// doubled-suffix bug was found on screenshot review: a
				// plain inventory check (does this band's own eligible
				// pool carry a haste stat at all), independent of whether
				// the sweep's own sample happened to land significant -
				// see bandReport.HasteOnItems' own doc.
				report.HasteOnItems = bandHasHasteCandidate(pool.Scored, pool.NoSource)
				reports = append(reports, report)
				if pass.name == presetBare {
					previous[f.name] = picks
				}

				// This lane's brief, item 1's last sentence: name every id
				// this band published with SimStatus "not_in_sim" once per
				// spec run, at warning level, so the nightly log tells the
				// Python data lane exactly which ids its own simdb.bin
				// rebuild needs to carry a row for (report.go's own
				// SimStatus/SetDPSPartial doc has the full reasoning).
				for _, row := range report.Slots {
					if row.SimStatus != notInSimReason || notInSimWarned[row.ItemID] {
						continue
					}
					notInSimWarned[row.ItemID] = true
					log.Printf("leveling-bis: %s: WARNING item %d (%s) is not in this build's simdb.bin (simdb.Known false) - stripped by simdb.Attach's UnequipUnknown before every sim, published score-decided with sim_status=not_in_sim", spec, row.ItemID, row.ItemName)
				}

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
	}

	jsonPath := filepath.Join(outDir, spec+".json")
	presets := map[string]request.ResolvedPreset{presetRaid: raidPreset}
	if err := writeSpecReport(jsonPath, spec, activeBuild, reports, presets); err != nil {
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
// bandTalentStrings is one band's talent build in both layouts this
// pipeline needs: site - the published band `talents` field, in the
// active build's own (tier, column) order, truncated to this band's
// points - and engine - the SAME points, repositioned onto the
// compiled engine's own field order (sim/internal/enginetalents' own
// doc) so a character this pipeline hands the engine reads correctly
// even for a class (paladin, shaman - this lane's own brief) whose
// tree shape moved between the engine's proto build and the active
// one. Every rank/verify/set/weights character this pipeline builds
// must be spent from engine, never site; pulling the pairing out of
// the band loop into its own function gives a test one place to pin
// that invariant for a real class the two layouts are known to
// differ on, rather than trusting every call site to keep passing the
// right one of the two strings by hand.
func bandTalentStrings(activeTrees []leveling.TalentTree, targets map[int]int, treeIndex, band int, layout talentLayout) (site, engine string, err error) {
	site = leveling.LadderTalentString(activeTrees, targets, treeIndex, band)
	engine, err = layout.Reposition(activeTrees, site)
	if err != nil {
		return "", "", fmt.Errorf("converting talents to the engine's own layout: %w", err)
	}
	return site, engine, nil
}

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
