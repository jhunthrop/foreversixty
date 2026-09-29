package main

// A permanent sanity test over the PUBLISHED output, not a fixture:
// data/builds/<build>/bis/<spec>.json is what the site's page actually
// reads (report.go's own specReport doc: "lane bis-web's read
// contract"), and tonight's audits found real defects in it only by
// eye (a level-38 rep bow on a level-20 list, Ragefire Chasm loot for
// Alliance, an item from the other faction's rep, a two-hander with
// an off hand, the same trinket twice, level-60 quest rewards at 20).
// This file re-derives the same facts main.go's own pipeline used to
// build each pick - eligibility, source, faction, weapon shape - from
// the SAME loaders (data.go/band.go) against the SAME build's item
// and loot data every committed bis/*.json names in its own "build"
// field, and checks every published pick against them, so a future
// regression in either the ranker or the data fails a test instead of
// waiting for another eyeball pass.
//
// Every other test in this package reads testdata/reporoot, a small
// hand-built stand-in (data_test.go's own doc). This one deliberately
// reads the real, committed repository instead: the whole point is a
// test over what the nightly run actually publishes, which a fixture
// could drift from silently.

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// publishedRepoRoot is this repository's own real root: three levels
// up from this package (sim/cmd/leveling-bis -> sim/cmd -> sim ->
// repo root).
const publishedRepoRoot = "../../.."

// TestPublishedBISFilesPassSanityChecks walks every
// data/builds/<build>/bis/<spec>.json the repository currently
// carries. A build with no bis directory at all (not every
// data/builds/<build> has been ranked - most never will be) is
// skipped cleanly: filepath.Glob against a directory that does not
// exist returns no matches and no error, not a failure.
func TestPublishedBISFilesPassSanityChecks(t *testing.T) {
	// The committed bis files are the NIGHTLY's output: a ranker fix
	// lands on main before the files that carry it, so this check
	// runs where fresh output exists -- bis.yml right after `make bis`
	// (FOREVER_BIS_SANITY=1) -- and is skipped in an ordinary test run.
	if os.Getenv("FOREVER_BIS_SANITY") == "" {
		t.Skip("set FOREVER_BIS_SANITY=1 to check the committed bis files (the nightly does)")
	}
	buildDirs, err := filepath.Glob(filepath.Join(publishedRepoRoot, "data", "builds", "*"))
	if err != nil {
		t.Fatalf("glob data/builds/*: %v", err)
	}
	if len(buildDirs) == 0 {
		t.Fatalf("no data/builds/* found under %s - publishedRepoRoot looks wrong", publishedRepoRoot)
	}

	tested := 0
	for _, buildDir := range buildDirs {
		files, err := filepath.Glob(filepath.Join(buildDir, "bis", "*.json"))
		if err != nil {
			t.Fatalf("glob %s/bis/*.json: %v", buildDir, err)
		}
		for _, f := range files {
			f := f
			t.Run(filepath.Base(buildDir)+"/"+strings.TrimSuffix(filepath.Base(f), ".json"), func(t *testing.T) {
				checkPublishedSpecFile(t, f)
			})
			tested++
		}
	}
	if tested == 0 {
		t.Skip("no data/builds/*/bis/*.json published yet - nothing to check")
	}
}

// checkPublishedSpecFile loads one published specReport and its
// build's own item/loot data (data.go's own loaders - the same ones
// main.go's runSpec used to write it), then applies every check this
// lane's brief lists to every band.
func checkPublishedSpecFile(t *testing.T, path string) {
	t.Helper()

	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading %s: %v", path, err)
	}
	var report specReport
	if err := json.Unmarshal(b, &report); err != nil {
		t.Fatalf("decoding %s: %v", path, err)
	}
	if len(report.Bands) == 0 {
		t.Fatalf("%s: published with zero bands", path)
	}

	spec, err := loadSpec(publishedRepoRoot, report.Spec)
	if err != nil {
		t.Fatalf("%s: loadSpec(%q): %v", path, report.Spec, err)
	}

	buildDir := filepath.Join(publishedRepoRoot, "data", "builds", report.Build)
	items, _, err := loadCandidates(buildDir, spec.ClassSlug)
	if err != nil {
		t.Fatalf("%s: loadCandidates(%s, %s): %v", path, buildDir, spec.ClassSlug, err)
	}
	// itemFactionRestriction mirrors main.go's own pre-pass exactly
	// (runSpec's own comment): correctedRepSource needs it before
	// loadLootIndex can build a faction-corrected source index.
	itemFactionRestriction := make(map[int]string, len(items))
	for _, c := range items {
		if c.FactionRestriction != "" {
			itemFactionRestriction[c.ID] = c.FactionRestriction
		}
	}
	lootIdx, questFloors, err := loadLootIndex(buildDir, itemFactionRestriction)
	if err != nil {
		t.Fatalf("%s: loadLootIndex: %v", path, err)
	}
	items = applyEffectiveRequiredLevels(items, lootIdx, questFloors)
	byID := make(map[int]candidate, len(items))
	for _, c := range items {
		byID[c.ID] = c
	}

	checkHunterSurvivalWeightsMeleeNotRanged(t, path, spec)

	for _, band := range report.Bands {
		checkPublishedBand(t, path, band, spec, byID, lootIdx)
	}
}

// checkPublishedBand applies checks 1-5, 7 and 8 (this lane's brief)
// to one band+faction's published slots.
func checkPublishedBand(t *testing.T, path string, band bandReport, spec specInfo, byID map[int]candidate, lootIdx lootIndex) {
	t.Helper()
	bandLabel := fmt.Sprintf("%s: %s band %d %s", path, band.Spec, band.Band, band.Faction)
	slotLabel := func(slot string) string { return bandLabel + " " + slot }

	bySlot := make(map[string]slotRow, len(band.Slots))
	for _, row := range band.Slots {
		bySlot[row.Slot] = row
		if row.ItemID == 0 {
			continue
		}

		c, ok := byID[row.ItemID]
		if !ok {
			t.Errorf("%s: item %d (%s) is not in %s's item file at all", slotLabel(row.Slot), row.ItemID, row.ItemName, spec.ClassSlug)
			continue
		}

		// (1) EffectiveRequiredLevel <= the band it was published at.
		if c.EffectiveRequiredLevel > band.Band {
			t.Errorf("%s: item %d (%s) needs level %d (sim/leveling.EffectiveRequiredLevel), published at band %d", slotLabel(row.Slot), row.ItemID, row.ItemName, c.EffectiveRequiredLevel, band.Band)
		}

		// (2) faction_restriction empty or matches this band's faction.
		if c.FactionRestriction != "" && c.FactionRestriction != band.Faction {
			t.Errorf("%s: item %d (%s) is restricted to %s, published for %s", slotLabel(row.Slot), row.ItemID, row.ItemName, c.FactionRestriction, band.Faction)
		}

		// (3) sourceFor returns an obtainable source, and the published
		// source/source_kind agree with it.
		src, ok := sourceFor(row.ItemID, band.Band, band.Faction, c.FactionRestriction, lootIdx)
		if !ok {
			t.Errorf("%s: item %d (%s) has no source obtainable by a %s character at level %d (sourceFor), but was published as a pick", slotLabel(row.Slot), row.ItemID, row.ItemName, band.Faction, band.Band)
		} else {
			if row.Source != src.Label {
				t.Errorf("%s: item %d (%s) published source %q, sourceFor says %q", slotLabel(row.Slot), row.ItemID, row.ItemName, row.Source, src.Label)
			}
			if row.SourceKind != src.Kind {
				t.Errorf("%s: item %d (%s) published source_kind %q, sourceFor says %q", slotLabel(row.Slot), row.ItemID, row.ItemName, row.SourceKind, src.Kind)
			}
		}

		// (7, second half) report.go's own omitempty drops an exact
		// 0.0 score, so any score field a published row DOES carry is
		// never legitimately <= 0 - a decoded negative (or exact zero,
		// which omitempty should have already dropped) score means an
		// item whose stats actively hurt this spec's weights out-
		// ranked leaving the slot empty, or a writer/decoder mismatch
		// with report.go's own contract.
		if row.Score < 0 {
			t.Errorf("%s: item %d (%s) published a negative score (%.4f) - a net-negative item should never outrank leaving this slot empty", slotLabel(row.Slot), row.ItemID, row.ItemName, row.Score)
		}
	}

	// (4) finger1/finger2 and trinket1/trinket2 (and, the same shape,
	// main_hand/off_hand for a dual-wield spec - verify.go's own
	// pairSlot) never carry the same item id or name.
	checkPairNotDuplicated(t, slotLabel, bySlot, "finger1", "finger2")
	checkPairNotDuplicated(t, slotLabel, bySlot, "trinket1", "trinket2")
	checkPairNotDuplicated(t, slotLabel, bySlot, "main_hand", "off_hand")

	// (5) off_hand's shape: absent for a two-handed main_hand; a
	// weapon for a dual-wield spec; a shield/held item (never a
	// weapon) for everyone else.
	checkOffHandShape(t, slotLabel, spec, byID, bySlot)

	// (7, first half) a band with a rotation (every published spec
	// here is - writtenSpecs' own "written" gate) must have measured a
	// positive set DPS.
	if band.SetDPS <= 0 {
		t.Errorf("%s: set_dps = %v, want > 0", bandLabel, band.SetDPS)
	}

	// (8) the published weights list is exactly this spec's own
	// WeightStats (data/curated/specs.json), in the same order -
	// neither missing a stat the spec is scored on nor carrying one it
	// is not.
	checkWeightsCoverSpecStats(t, bandLabel, band, spec)

	// (9) this lane's brief, item 1: every published alternative is
	// itself eligible, sourced (with the source/source_kind sourceFor
	// says it has), distinct from the pick and from every other
	// alternative, never the slot's own pair-mate, and bounded at
	// alternativesLimit.
	checkAlternatives(t, slotLabel, byID, lootIdx, band, bySlot)
}

// checkAlternatives applies check (9) to every filled slot's
// Alternatives: each one is a real, eligible, sourced candidate at
// this band+faction (the same three facts checks 1-3 already demand
// of the pick itself), never the pick, never listed twice, never the
// slot's own pair-mate (pairSlot, verify.go), and the list never
// exceeds alternativesLimit (report.go).
func checkAlternatives(t *testing.T, slotLabel func(string) string, byID map[int]candidate, lootIdx lootIndex, band bandReport, bySlot map[string]slotRow) {
	t.Helper()
	for _, row := range bySlot {
		if row.ItemID == 0 || len(row.Alternatives) == 0 {
			continue
		}
		if len(row.Alternatives) > alternativesLimit {
			t.Errorf("%s: %d alternatives published, want at most %d (alternativesLimit)", slotLabel(row.Slot), len(row.Alternatives), alternativesLimit)
		}
		mateRow, hasMate := bySlot[pairSlot[row.Slot]]
		seen := map[int]bool{row.ItemID: true}
		for _, alt := range row.Alternatives {
			if seen[alt.ItemID] {
				t.Errorf("%s: alternative %d (%s) published more than once, or duplicates the pick itself", slotLabel(row.Slot), alt.ItemID, alt.ItemName)
			}
			seen[alt.ItemID] = true

			c, ok := byID[alt.ItemID]
			if !ok {
				t.Errorf("%s: alternative %d (%s) is not in %s's item file at all", slotLabel(row.Slot), alt.ItemID, alt.ItemName, band.Spec)
				continue
			}
			if c.EffectiveRequiredLevel > band.Band {
				t.Errorf("%s: alternative %d (%s) needs level %d (sim/leveling.EffectiveRequiredLevel), published at band %d", slotLabel(row.Slot), alt.ItemID, alt.ItemName, c.EffectiveRequiredLevel, band.Band)
			}
			if c.FactionRestriction != "" && c.FactionRestriction != band.Faction {
				t.Errorf("%s: alternative %d (%s) is restricted to %s, published for %s", slotLabel(row.Slot), alt.ItemID, alt.ItemName, c.FactionRestriction, band.Faction)
			}
			src, ok := sourceFor(alt.ItemID, band.Band, band.Faction, c.FactionRestriction, lootIdx)
			if !ok {
				t.Errorf("%s: alternative %d (%s) has no source obtainable by a %s character at level %d (sourceFor), but was published as an alternative", slotLabel(row.Slot), alt.ItemID, alt.ItemName, band.Faction, band.Band)
			} else if alt.Source != src.Label || alt.SourceKind != src.Kind {
				t.Errorf("%s: alternative %d (%s) published source %q/%q, sourceFor says %q/%q", slotLabel(row.Slot), alt.ItemID, alt.ItemName, alt.Source, alt.SourceKind, src.Label, src.Kind)
			}

			if hasMate && mateRow.ItemID != 0 && (alt.ItemID == mateRow.ItemID || (alt.ItemName != "" && alt.ItemName == mateRow.ItemName)) {
				t.Errorf("%s: alternative %d (%s) is this slot's own pair-mate (%s, id %d) - the same physical item cannot be offered as a fallback here", slotLabel(row.Slot), alt.ItemID, alt.ItemName, mateRow.ItemName, mateRow.ItemID)
			}

			// (10) this lane's brief (bis-ranker-integrity, 2026-09-29),
			// item 1: a positive dps_delta is a claim that this
			// alternative actually beats the verified pick - the ONE
			// claim tenet 8 demands real evidence for. buildAlternatives'
			// own cap (report.go) is supposed to guarantee this never
			// happens for a row verify.go did not actually sim against
			// the pick; this check holds that guarantee to the published
			// output itself; not the field the ranker set out to
			// produce, so a future regression in either place fails
			// here instead of waiting for another wow-player pass to
			// catch it by eye (mage-fire band 60 trinket1's own Neltharion's
			// Tear alternative, +345 "DPS", unverified, is exactly the
			// defect this guards against).
			if alt.DPSDelta > 0 && !alt.Verified {
				t.Errorf("%s: alternative %d (%s) publishes a positive dps_delta (%.2f) with verified omitted - an unverified candidate must never claim to beat the pick", slotLabel(row.Slot), alt.ItemID, alt.ItemName, alt.DPSDelta)
			}
		}
	}
}

// checkPairNotDuplicated reports a violation when slots a and b (both
// populated) name the same physical item: the same id, or the same
// name (a lower/higher-quality reprint of "the same ring/trinket/
// weapon"), matching pick.go's own excludePaired rule and verify.go's
// own pairSlot doc for why a duplicate can otherwise slip past it
// (this lane's report: the applySwaps fix).
func checkPairNotDuplicated(t *testing.T, slotLabel func(string) string, bySlot map[string]slotRow, a, b string) {
	t.Helper()
	ra, oka := bySlot[a]
	rb, okb := bySlot[b]
	if !oka || !okb || ra.ItemID == 0 || rb.ItemID == 0 {
		return
	}
	if ra.ItemID == rb.ItemID {
		t.Errorf("%s and %s: both wear item %d (%s) - the same physical item cannot fill both slots", slotLabel(a), slotLabel(b), ra.ItemID, ra.ItemName)
		return
	}
	if ra.ItemName != "" && ra.ItemName == rb.ItemName {
		t.Errorf("%s and %s: both named %q (ids %d and %d) - a lower/higher-quality reprint cannot fill both slots either", slotLabel(a), slotLabel(b), ra.ItemName, ra.ItemID, rb.ItemID)
	}
}

// checkOffHandShape applies check (5): a two-handed main_hand leaves
// off_hand empty (buildGear's own defended rule); otherwise off_hand,
// if populated, must be a weapon (item class 2, itemClassWeapon) for
// a leveling.DualWieldSpecs member and must NOT be a weapon (a
// shield or held item instead) for everyone else.
func checkOffHandShape(t *testing.T, slotLabel func(string) string, spec specInfo, byID map[int]candidate, bySlot map[string]slotRow) {
	t.Helper()
	mh := bySlot["main_hand"]
	oh := bySlot["off_hand"]
	if mh.ItemID == 0 {
		return
	}
	mhItem, ok := byID[mh.ItemID]
	if !ok {
		return // already reported as missing-from-item-file above
	}

	if mhItem.TwoHand {
		if oh.ItemID != 0 {
			t.Errorf("%s: main_hand %d (%s) is two-handed, but off_hand still carries %d (%s) - the sim (buildGear) never actually wore both at once", slotLabel("main_hand/off_hand"), mh.ItemID, mh.ItemName, oh.ItemID, oh.ItemName)
		}
		return
	}
	if oh.ItemID == 0 {
		return // an empty off_hand beside a one-handed main_hand is always fine
	}
	ohItem, ok := byID[oh.ItemID]
	if !ok {
		return
	}

	dualWield := leveling.DualWieldSpecs[spec.Spec]
	isWeapon := ohItem.ClassID == itemClassWeapon
	switch {
	case dualWield && !isWeapon:
		t.Errorf("%s: %s dual-wields (leveling.DualWieldSpecs), but off_hand %d (%s) is not a weapon", slotLabel("off_hand"), spec.Spec, oh.ItemID, oh.ItemName)
	case !dualWield && isWeapon:
		t.Errorf("%s: %s does not dual-wield, but off_hand %d (%s) is a weapon (want a shield or held item)", slotLabel("off_hand"), spec.Spec, oh.ItemID, oh.ItemName)
	}
}

// checkWeightsCoverSpecStats applies check (8): the published weights
// array names exactly spec.WeightStats, in the same order - a spec
// with a stat missing from the page's aside, or an extra one nobody
// asked to be weighted, both mean this report and specs.json have
// drifted apart.
func checkWeightsCoverSpecStats(t *testing.T, bandLabel string, band bandReport, spec specInfo) {
	t.Helper()
	got := make([]string, len(band.Weights))
	for i, w := range band.Weights {
		got[i] = w.Stat
	}
	if len(got) != len(spec.WeightStats) {
		t.Errorf("%s: published weights %v (%d stats), want exactly specs.json's weight_stats %v (%d stats)", bandLabel, got, len(got), spec.WeightStats, len(spec.WeightStats))
		return
	}
	for i, want := range spec.WeightStats {
		if got[i] != want {
			t.Errorf("%s: weights[%d] = %q, want %q (specs.json's weight_stats, in order)", bandLabel, i, got[i], want)
		}
	}
}

// checkHunterSurvivalWeightsMeleeNotRanged applies check (6): Forever's
// survival hunter fights in melee (this repo's own memory/tenets), so
// its ranker must never treat the ranged slot as the spec's damage
// weapon. rank.go's request/ladder.go sibling encodes this as a
// ladderGearProfile/weapon-type table, but that table is an unexported
// var in a different package (sim/request) this command cannot import
// (data.go's own aplState doc gives the same reason for its own small
// duplication); specs.json's weight_stats/reference_stat is this
// spec's own equivalent fact and IS available here - score.go's own
// weaponAPStat map converts a slot's weapon DPS into the weight_stats
// entry that slot's own AP stat maps to, so a survival hunter whose
// weight_stats never mentions ranged_attack_power is a hunter whose
// ranged slot's DPS never contributes to its score at all, exactly
// the "never picks a ranged weapon as its damage weapon" this check
// is asking for.
func checkHunterSurvivalWeightsMeleeNotRanged(t *testing.T, path string, spec specInfo) {
	t.Helper()
	if spec.Spec != "hunter-survival" {
		return
	}
	for _, stat := range spec.WeightStats {
		if stat == "ranged_attack_power" {
			t.Errorf("%s: hunter-survival's weight_stats carries ranged_attack_power - Forever's survival hunter fights in melee; weighting ranged_attack_power would make score.go convert the ranged slot's DPS into the spec's damage stat, exactly the regression this check exists to catch", path)
		}
	}
	if spec.ReferenceStat != "attack_power" {
		t.Errorf("%s: hunter-survival's reference_stat is %q, want attack_power (melee) - see this function's own doc", path, spec.ReferenceStat)
	}
}
