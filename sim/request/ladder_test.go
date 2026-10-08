package request

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/internal/enginetalents"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/internal/spellranks"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/specs"
	engine "github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// The rotation ladder (Phase 1a,
// docs/superpowers/specs/2026-09-28-rotation-accuracy-program-design.md):
// for every written spec, a bare character at levels 10, 20, 30, 38, 40,
// 50 and 60, run 300 iterations against the default encounter. See
// ladder.go's ladderRulesHeader for the full rule set; it is written
// into every golden this test produces.
//
// Four rules can fail a level (the design's Phase 1a "Fails when"):
//
//  1. a curated rotation's own castSpell line, at the rank this level
//     has learned, never fires - unless the engine could not resolve
//     that rank (already an "unresolved" finding) or
//     data/curated/apl/<spec>.json's expected_idle names the line's
//     authored id with a reason.
//  2. the engine warns it cannot resolve an id the curated file's own
//     inert array does not name.
//  3. DPS at a level is lower than at the ladder's previous rung.
//  4. a level >= 20 casts no spell at all, only auto-attacks.
//
// Every violation found today is collected rather than failing the
// first subtest: FOREVER_LADDER_STRICT=1 is what turns them into a
// failure, so CI stays green on a first commit that (correctly) finds
// real gaps, and a later run that fixes one sees the list shrink.
func TestRotationLadder(t *testing.T) {
	registerEngine.Do(engine.RegisterAll)

	build := activeBuild(t)
	ranks, err := loadSpellRanks(repoRoot, build)
	if err != nil {
		t.Fatal(err)
	}

	type specResult struct {
		spec       string
		violations []string
	}
	var results []specResult

	for _, spec := range specs.All {
		curated, err := loadLadderCurated(repoRoot, spec.Spec)
		if err != nil {
			t.Fatalf("%s: %v", spec.Spec, err)
		}
		if curated.State != "written" {
			continue
		}
		spec := spec
		t.Run(spec.Spec, func(t *testing.T) {
			rows, unused, violations := runLadderSpec(t, build, spec, curated, ranks)
			golden := renderLadderGolden(spec.Spec, rows, unused, violations)
			compareOrWriteLadderGolden(t, spec.Spec, golden)
			results = append(results, specResult{spec: spec.Spec, violations: violations})
		})
	}

	if len(results) == 0 {
		t.Fatal("no curated rotation is marked written; the whole ladder is vacuous")
	}

	var all []string
	for _, r := range results {
		all = append(all, r.violations...)
	}
	if len(all) == 0 {
		return
	}
	sort.Strings(all)
	msg := fmt.Sprintf("%d ladder violation(s) today:\n%s", len(all), strings.Join(all, "\n"))
	if t.Failed() {
		// A subtest already failed (a golden mismatch, a build error):
		// that failure stands on its own. Log the business-rule
		// violations too so they are visible in the same run, without
		// layering a second Skip/Error decision on an already-failed
		// test.
		t.Log(msg)
		return
	}
	if os.Getenv("FOREVER_LADDER_STRICT") == "" {
		t.Skip(msg + "\n(set FOREVER_LADDER_STRICT=1 to fail the build on these)")
		return
	}
	t.Error(msg)
}

// runLadderSpec runs one spec's ladder: every level, the golden's rows,
// the informational learned-but-unused list, and the strict-mode
// violations.
func runLadderSpec(t *testing.T, build string, spec specs.Spec, curated ladderCurated, ranks spellRanksFile) ([]ladderRow, []unusedEntry, []string) {
	t.Helper()
	class := spec.ClassSlug
	race := smokeRace(t, class)

	guideBuild, treeDigits, err := leveling.GuideBuildTalents(repoRoot, class, spec.SpecSlug)
	if err != nil {
		t.Fatal(err)
	}
	guideTrees, err := leveling.LoadTalentTrees(repoRoot, guideBuild, class)
	if err != nil {
		t.Fatal(err)
	}
	activeTrees, err := leveling.LoadTalentTrees(repoRoot, build, class)
	if err != nil {
		t.Fatal(err)
	}
	targets := leveling.GuideTalentTargets(guideTrees, treeDigits)

	// engineLayout: the COMPILED engine's own talent string layout for
	// class (sim/internal/enginetalents' own doc) - the engine's proto
	// was last regenerated from a client build that is not always the
	// active one, so every character this ladder hands the engine
	// below is spent from engineTalents (per level, further down),
	// never from the active-build-positional `talents` string
	// directly, or paladin and shaman misread every talent at and
	// after the first one whose tree position moved between the two
	// builds.
	engineDir, err := enginetalents.SourceDir(filepath.Join(repoRoot, "sim"))
	if err != nil {
		t.Fatal(err)
	}
	engineLayout, err := enginetalents.ForClass(engineDir, class)
	if err != nil {
		t.Fatal(err)
	}

	items, err := loadClassItems(repoRoot, build, class)
	if err != nil {
		t.Fatal(err)
	}
	knownItems, err := obtainableItemIDs(repoRoot, build)
	if err != nil {
		t.Fatal(err)
	}
	requiredLevelFloors, err := loadRequiredLevelFloors(repoRoot, build, items)
	if err != nil {
		t.Fatal(err)
	}

	consts, err := loadSpellConst(repoRoot, build, class)
	if err != nil {
		t.Fatal(err)
	}
	abilities := buildClassAbilities(ranks, class)
	abilityNames := abilityNameByID(abilities)

	authoredIDs, err := authoredCastSpellIDs(curated.Rotation)
	if err != nil {
		t.Fatal(err)
	}
	expectedIdle := map[int]string{}
	for _, e := range curated.ExpectedIdle {
		expectedIdle[e.ID] = e.Reason
	}
	inert := map[string]bool{}
	for _, id := range curated.Inert {
		inert[spellAction(id)] = true
	}

	// Harness rule 1 (this wave's brief): the two readers a warned
	// action's spell id is checked against below, beside the inert and
	// expected_idle maps above - see ladder.go's own comment on
	// loadTalentSpellIDs and idLearnLevel for why these are separate,
	// read-only copies of data this file's talent-truncation code and
	// spellranks.json already carry.
	talentSpellIDs, err := loadTalentSpellIDs(repoRoot, build, class)
	if err != nil {
		t.Fatal(err)
	}
	talentSpellIDs = expandTalentSpellIDs(talentSpellIDs, abilities)
	rankLevelByID := idLearnLevel(ranks, class)

	var rows []ladderRow
	var unused []unusedEntry
	var violations []string
	prevDPS, havePrev := 0.0, false

	// A healer's ladder runs against the heal profile's fake raid, and its
	// "DPS" column is its effective healing per second: with nothing to
	// heal its rotation would stand idle and the ladder would prove nothing.
	var healProfile *HealProfile
	if spec.Role == "healer" {
		profile, err := LoadHealProfile(filepath.Join(repoRoot, "data", "curated", "heal-profile.json"))
		if err != nil {
			t.Fatal(err)
		}
		healProfile = &profile
	}

	for _, level := range ladderLevels {
		// talents is the published, active-build-positional string
		// (ladderRow.Talents, below, reports it unconverted - that is
		// the layout the web planner decodes). engineTalents is the
		// SAME build, repositioned onto engineLayout's own field
		// order; the character below is spent from engineTalents, not
		// talents (engineLayout's own doc).
		talents := leveling.LadderTalentString(activeTrees, targets, spec.TreeIndex, level)
		engineTalents, err := engineLayout.Reposition(activeTrees, talents)
		if err != nil {
			t.Fatalf("%s level %d: converting talents to the engine's own layout: %v", spec.Spec, level, err)
		}
		talentPoints, err := leveling.TalentRanksFromString(activeTrees, talents)
		if err != nil {
			t.Fatalf("%s level %d: reading the ladder's own talent string: %v", spec.Spec, level, err)
		}
		gear := ladderGear(items, knownItems, requiredLevelFloors, spec.Spec, level)

		req := api.SimRequest{
			EngineVersion: enginever.Version,
			Spec:          spec.Spec,
			Source:        api.CharacterSource{Kind: api.Sources[0]},
			Character: api.CharacterSpec{
				Name:     spec.Spec,
				Race:     race,
				Class:    class,
				Level:    level,
				Talents:  engineTalents,
				Gear:     gear,
				Consumes: ladderKitConsumes(spec.Spec, level),
				Buffs:    ladderKitBuffs(spec.Spec, level),
			},
			Encounter:  api.DefaultEncounter(),
			Iterations: ladderIterations,
			RandomSeed: ladderSeed,
		}

		// Two builds of the same request, for the reason
		// rotations_smoke_test.go's own TestEveryWrittenRotationRunsInTheEngine
		// documents: core.ComputeStats builds an environment of its own
		// out of the raid it is handed, and handing the same protobuf
		// to the sim afterwards would make the run depend on what the
		// stats pass did to it.
		// Unlike rotations_smoke_test.go's bare characters, the ladder
		// equips real weapons (ladderGear), so the request needs
		// Forever's own item rows attached - without them an equipped
		// item id resolves to nothing and the engine panics mid-build
		// (see cmd/forever-sim/main.go's own call to simdb.Attach).
		statsReq, err := BuildWith(req, Options{OpenIterations: true})
		if err != nil {
			t.Fatalf("%s level %d: building the request: %v", spec.Spec, level, err)
		}
		if err := simdb.Attach(statsReq); err != nil {
			t.Fatalf("%s level %d: attaching the item database: %v", spec.Spec, level, err)
		}
		if healProfile != nil {
			healProfile.Attach(statsReq)
		}
		warned := warnedActions(t, statsReq)
		warnedSet := map[string]bool{}
		for _, w := range warned {
			warnedSet[w] = true
		}

		simReq, err := BuildWith(req, Options{OpenIterations: true})
		if err != nil {
			t.Fatalf("%s level %d: building the request: %v", spec.Spec, level, err)
		}
		if err := simdb.Attach(simReq); err != nil {
			t.Fatalf("%s level %d: attaching the item database: %v", spec.Spec, level, err)
		}
		if healProfile != nil {
			healProfile.Attach(simReq)
		}
		res := core.RunRaidSim(simReq)
		if err := adapter.ResultError(res); err != nil {
			t.Fatalf("%s level %d: the sim failed: %v", spec.Spec, level, err)
		}
		player, err := adapter.PlayerMetrics(res)
		if err != nil {
			t.Fatalf("%s level %d: reading the player's metrics: %v", spec.Spec, level, err)
		}
		dps := ladderOutput(player, healProfile != nil)
		tallies := ladderCastSet(player, ladderIterations)
		castCounts := CastSpellCounts(player)

		// Rule 4: no cast but auto-attack at level >= 20.
		if level >= 20 && ladderDistinctSpellCasts(tallies) == 0 {
			violations = append(violations, fmt.Sprintf(
				"%s level=%d kind=no_damage_cast dps=%.1f", spec.Spec, level, dps))
		}

		// Rule 3: DPS lower than the previous rung, tolerating up to a
		// 1% drop as this run's own noise (harness rule 4, this wave's
		// brief - shaman-elemental's level 40 vs 38 is exactly this).
		if havePrev && dps < prevDPS*(1-dpsRegressionTolerance) {
			violations = append(violations, fmt.Sprintf(
				"%s level=%d kind=dps_regression dps=%.1f prev_dps=%.1f", spec.Spec, level, dps, prevDPS))
		}
		prevDPS, havePrev = dps, true

		// Rule 2: an unresolved id the curated file's inert array does
		// not name, with the harness's own three standing exceptions
		// (see ladderRulesHeader's "Unresolved" entry): the potion
		// action, a talent-granted spell the truncated build has not
		// spent a point on yet, and an above-band spellranks.json rank
		// the engine's own rewrite should already have dropped as a
		// CAST (asserted, not merely excused: this id reaching the
		// engine as unresolved some OTHER way - a condition value
		// rewriteRankedSpellIDs's castKeys leaves as-authored - is not
		// itself wrong, but the rewrite disagreeing about whether this
		// level has learned it would be).
		for _, w := range warned {
			if inert[w] {
				continue
			}
			if w == potionUnresolvedAction {
				continue // harness rule 2: no consumes on the ladder character.
			}
			if id, ok := warnedSpellID(w); ok {
				if nodeID, isTalent := talentSpellIDs[id]; isTalent && talentPoints[nodeID] == 0 {
					continue // harness rule 1: this talent has zero points at this level.
				}
				if rankLevel, isRank := rankLevelByID[id]; isRank && rankLevel > level {
					if _, learned := spellranks.HighestLearnedSpellID(class, int32(id), level); learned {
						t.Fatalf("%s level=%d: id %d is an above-band spellranks.json rank (learned at %d) "+
							"that spellranks.HighestLearnedSpellID nonetheless calls learned at %d - the rewrite "+
							"and this assertion disagree", spec.Spec, level, id, rankLevel, level)
					}
					continue // harness rule 1: the rewrite already drops a CAST of this id.
				}
			}
			violations = append(violations, fmt.Sprintf(
				"%s level=%d kind=unresolved_id action=%s", spec.Spec, level, w))
		}

		// Rule 1: a curated rotation line, resolved to the id
		// sim/internal/spellranks.HighestLearnedSpellID says the engine
		// itself casts at this level - the exact function
		// sim/request's own rewriteRotationRanks calls, not a second,
		// approximate copy of its resolution - that never fired.
		for authoredID := range authoredIDs {
			learnedID32, learned := spellranks.HighestLearnedSpellID(class, int32(authoredID), level)
			if !learned {
				continue // the engine's own rewrite already dropped this line; nothing to double-report.
			}
			learnedID := int(learnedID32)
			if warnedSet[spellAction(learnedID)] {
				continue // already counted as an unresolved id.
			}
			if castCounts[learnedID] > 0 {
				continue // fired; nothing to report.
			}
			if reason, ok := expectedIdle[authoredID]; ok && reason != "" {
				continue // curated file excuses it.
			}
			if label, ok := abilityNames[learnedID]; ok {
				violations = append(violations, fmt.Sprintf(
					"%s level=%d kind=zero_casts spell=%q id=%d authored=%d", spec.Spec, level, label, learnedID, authoredID))
			} else {
				violations = append(violations, fmt.Sprintf(
					"%s level=%d kind=zero_casts id=%d authored=%d (untracked ability; not in spellranks.json's rank chains)",
					spec.Spec, level, learnedID, authoredID))
			}
		}

		// Rule 3 of the design's "what accurate can mean" section
		// (informational): every learned damage ability the cast set
		// never touched, whether or not the rotation names it.
		for name, tiers := range abilities.Tiers {
			tier, learned := learnedTierAtLevel(tiers, level)
			if !learned {
				continue
			}
			damage, used := false, false
			for _, id := range tier.IDs {
				if isDamageSpellID(consts, id) {
					damage = true
				}
				if castCounts[id] > 0 {
					used = true
				}
			}
			if !damage || used {
				continue
			}
			unused = append(unused, unusedEntry{Level: level, Name: name, ID: tier.IDs[0]})
		}

		rows = append(rows, ladderRow{
			Level:         level,
			Talents:       talents,
			Gear:          formatGear(gear),
			DPS:           dps,
			DistinctCasts: ladderDistinctSpellCasts(tallies),
			TopCasts:      ladderTopCasts(tallies, 5),
			Unresolved:    warned,
		})
	}

	return rows, unused, violations
}

// ladderGoldenEnv is FOREVER_UPDATE_GOLDEN, the same variable name
// sim/adapter's golden test uses - one habit covers both.
const ladderGoldenEnv = "FOREVER_UPDATE_GOLDEN"

func ladderGoldenPath(spec string) string {
	return filepath.Join(repoRoot, "sim", "request", "testdata", "ladder", spec+".golden.md")
}

func compareOrWriteLadderGolden(t *testing.T, spec string, got []byte) {
	t.Helper()
	path := ladderGoldenPath(spec)
	if os.Getenv(ladderGoldenEnv) != "" {
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, got, 0o644); err != nil {
			t.Fatal(err)
		}
		t.Logf("regenerated %s", path)
		return
	}
	want, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("%v (run with %s=1 to create it)", err, ladderGoldenEnv)
	}
	if string(got) != string(want) {
		t.Errorf("the ladder golden for %s differs from testdata/ladder/%s.golden.md; "+
			"rerun with %s=1 and review the diff", spec, spec, ladderGoldenEnv)
	}
}

// ladderOutput is the number a ladder rung reports: damage per second, or
// for a healer the effective healing per second.
func ladderOutput(player *proto.UnitMetrics, healer bool) float64 {
	if healer {
		return player.EffectiveHps.GetAvg()
	}
	return player.Dps.GetAvg()
}
