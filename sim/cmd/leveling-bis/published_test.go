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
	checkSetDPSMonotonicAcrossBands(t, path, report.Bands)

	for _, band := range report.Bands {
		checkPublishedBand(t, path, band, spec, byID, lootIdx)
	}
}

// checkSetDPSMonotonicAcrossBands is this lane's brief (bis-ranker-
// integrity-10), item 1: rogue-assassination's own published set_dps
// fell from band 30 (47.7/47.2, Alliance/Horde) to band 40
// (36.7/36.5) - the ninth wow-player sweep's melee-ranged finding 1,
// and the only non-monotonic band pair this repository's own bis
// files carried across 20 specs - while band 40's own weapons (Gut
// Ripper 2164, dps 33.89; Ardent Custodian 868, dps 32.86) genuinely
// out-DPS band 30's (Royal Diplomatic Scepter 9457, dps 23.04;
// Ironspine's Fist 7687, dps 22.92). Traced with a local, uncommitted
// `go run ./sim/cmd/leveling-bis -spec rogue-assassination -bands
// 30,40` against this build's own simdb.bin (copied from the main
// checkout for the trace only, per this lane's own instructions -
// never committed): the engine's own log named the cause directly -
// "item 2164 (Gut Ripper) is not in this build's simdb.bin... stripped
// by simdb.Attach's UnequipUnknown before every sim" - and band 30's
// own main_hand (9457) is unknown to simdb.bin exactly the same way.
// Both bands' main_hand is silently unequipped before every verify
// sim SetDPS is measured from, so both bands' SetDPS was always "this
// set, minus its main-hand weapon" (bandReport.SetDPSPartial's own
// doc), never a real dual-wield measurement of the set either band
// actually publishes - band 30's off-hand (Ironspine's Fist) simply
// carries the set further alone than band 40's off-hand (Ardent
// Custodian) does, which is not the same claim as "band 40 is a worse
// set". main.go's runSpec already detects this (markNotInSimDB,
// ranker-integrity-9 lane) and publishes it honestly: SetDPSPartial is
// true on every affected band, exactly the "published reason" this
// lane's brief asks for - the gap this function closes is that
// nothing enforced that reason ever actually gets published instead
// of a silent drop reaching site visitors. (The repository's own
// currently-committed data/builds/1.60.1.70009/bis/
// rogue-assassination.json predates this ranker-9 feature actually
// running against it - Nightly owns generated data, so this lane
// never regenerates or commits it; the next nightly run republishes
// it with SetDPSPartial already true on both bands, which is what
// lets this exact check pass once that happens.)
//
// The rule: within one faction, walking bands in ascending order (the
// order runSpec's own band loop writes them, per faction, in - this
// lane's brief's own words, "for every written spec"), a lower SetDPS
// than the band before it is only legitimate when at least one of the
// two bands involved is SetDPSPartial - a genuine full-set comparison
// (both bands fully simmable) must never regress, but a comparison
// where either side's own measurement is honestly incomplete is not
// evidence of a real regression at all, so it is not one of this
// check's failures - it is exactly the case SetDPSPartial exists to
// name instead of hiding.
func checkSetDPSMonotonicAcrossBands(t testing.TB, path string, bands []bandReport) {
	t.Helper()
	previous := map[string]bandReport{}
	for _, band := range bands {
		prev, ok := previous[band.Faction]
		if ok && band.SetDPS < prev.SetDPS && !prev.SetDPSPartial && !band.SetDPSPartial {
			t.Errorf("%s: %s %s: set_dps fell from %.2f (band %d) to %.2f (band %d) with no published reason (set_dps_partial is false on both bands) - tenet 8: a real drop needs a reason, never silent", path, band.Spec, band.Faction, prev.SetDPS, prev.Band, band.SetDPS, band.Band)
		}
		// A raid entry is compared against the bare entry of its own band
		// (written just before it) and never becomes the baseline the
		// next band is measured against.
		if band.Preset != presetRaid {
			previous[band.Faction] = band
		}
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

		// (12) this lane's brief (bis-ranker-integrity-16), item 3: a
		// faction_note that claims "this faction's own pick X" must
		// name the item this row actually publishes, on a FILLED row
		// (checked before the ItemID == 0 skip below, since this is
		// exactly the case an emptied row must never carry).
		checkFactionNoteNamesPublishedItem(t, slotLabel(row.Slot), row)

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

		// (11) armorAvailableLevel's own gate (eligible.go, this
		// lane's brief item 2): no band below 40 may publish a mail
		// pick for shaman/hunter or a plate pick for warrior/paladin.
		checkArmorProficiencyGate(t, slotLabel(row.Slot), spec.ClassSlug, band.Band, c, row.ItemID, row.ItemName)

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
	checkAlternatives(t, slotLabel, spec.ClassSlug, byID, lootIdx, band, bySlot)
}

// checkAlternatives applies check (9) to every filled slot's
// Alternatives: each one is a real, eligible, sourced candidate at
// this band+faction (the same three facts checks 1-3 already demand
// of the pick itself), never the pick, never listed twice, never the
// slot's own pair-mate (pairSlot, verify.go), and the list never
// exceeds alternativesLimit (report.go).
func checkAlternatives(t *testing.T, slotLabel func(string) string, classSlug string, byID map[int]candidate, lootIdx lootIndex, band bandReport, bySlot map[string]slotRow) {
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
			checkArmorProficiencyGate(t, slotLabel(row.Slot)+" alternative", classSlug, band.Band, c, alt.ItemID, alt.ItemName)
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

// checkArmorProficiencyGate applies armorAvailableLevel's own rule
// (eligible.go, this lane's brief item 2) directly to a published pick
// or alternative: mail below 40 for shaman/hunter, or plate below 40
// for warrior/paladin, must never be published, regardless of what
// eligible() would have said at ranking time - a stale committed
// bis/*.json (Nightly owns generated data: this lane never regenerates
// or commits it) is exactly what this test exists to catch before the
// next nightly republishes it.
func checkArmorProficiencyGate(t *testing.T, label, classSlug string, band int, c candidate, itemID int, itemName string) {
	t.Helper()
	if c.ClassID != armorClassID {
		return
	}
	key := fmt.Sprintf("%s:%d", classSlug, c.SubclassID)
	if opens, gated := armorAvailableLevel[key]; gated && band < opens {
		t.Errorf("%s: item %d (%s) is subclass %d, gated to level %d for %s (armorAvailableLevel), published at band %d", label, itemID, itemName, c.SubclassID, opens, classSlug, band)
	}
}

// factionNoteOwnPickMarker is the exact phrase
// reconcileTrinketDirection's own gainsIndistinguishable branch
// (faction_trinkets.go) always writes right before the item name it
// is claiming as this row's own pick - the one FactionNote shape
// whose truth depends on what this row actually publishes (pick.go's
// own FactionNoteNeedsPick doc). The OTHER branch that writes a
// FactionNote (negativeBeyondError, "racial") never uses this phrase:
// it names only the rejected crossing candidate, a claim that holds
// regardless of what this row publishes.
const factionNoteOwnPickMarker = "own pick "

// checkFactionNoteNamesPublishedItem applies this lane's brief
// (bis-ranker-integrity-16), item 3: shaman-elemental's Alliance band
// 50 trinket2 published EMPTY while its own faction_note still said
// "alliance's own pick Molten Heart of the Mountain" - an item named
// nowhere in the row (report.go's own trinketLowGain gate hid it
// after the note was already written). A faction_note naming an "own
// pick" must name the item this row actually shows as its pick, never
// an empty row and never a different item.
func checkFactionNoteNamesPublishedItem(t *testing.T, label string, row slotRow) {
	t.Helper()
	idx := strings.Index(row.FactionNote, factionNoteOwnPickMarker)
	if idx == -1 {
		return
	}
	named := strings.TrimSpace(row.FactionNote[idx+len(factionNoteOwnPickMarker):])
	if row.ItemID == 0 {
		t.Errorf("%s: faction_note %q claims an \"own pick\" but this row published empty (empty_reason %q) - the note must describe what publishes, or be omitted", label, row.FactionNote, row.EmptyReason)
		return
	}
	if row.ItemName != "" && !strings.HasPrefix(named, row.ItemName) {
		t.Errorf("%s: faction_note %q claims \"own pick\" %s, but this row's own published pick is %q (id %d)", label, row.FactionNote, named, row.ItemName, row.ItemID)
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

// fakeTB is a minimal testing.TB stand-in that only implements
// Helper()/Errorf() - everything checkSetDPSMonotonicAcrossBands
// actually calls - so its unit tests below can assert on exactly what
// was flagged without a real failing subtest bubbling up and marking
// this file's own test binary as failed (t.Run's failure always
// propagates to its parent, which is the wrong tool for asserting "did
// this checker correctly flag a synthetic violation"). The embedded
// nil testing.TB satisfies every other method by panicking if this
// function ever called one of them, which would itself fail whichever
// test reached it - exactly as it should.
type fakeTB struct {
	testing.TB
	errors []string
}

func (f *fakeTB) Helper() {}
func (f *fakeTB) Errorf(format string, args ...any) {
	f.errors = append(f.errors, fmt.Sprintf(format, args...))
}

// TestCheckSetDPSMonotonicAcrossBandsFlagsAnUnexplainedDrop is this
// lane's brief (bis-ranker-integrity-10), item 1's own contract test,
// as a synthetic unit test (checkSetDPSMonotonicAcrossBands' own doc
// explains the real rogue-assassination repro this table is modelled
// on) so it runs under an ordinary `go test`, not only
// FOREVER_BIS_SANITY's nightly pass over the committed bis files.
func TestCheckSetDPSMonotonicAcrossBandsFlagsAnUnexplainedDrop(t *testing.T) {
	bands := []bandReport{
		{Spec: "rogue-assassination", Band: 30, Faction: "alliance", SetDPS: 47.68},
		{Spec: "rogue-assassination", Band: 40, Faction: "alliance", SetDPS: 36.71},
	}
	fake := &fakeTB{}
	checkSetDPSMonotonicAcrossBands(fake, "testdata/synthetic.json", bands)
	if len(fake.errors) != 1 {
		t.Fatalf("checkSetDPSMonotonicAcrossBands errors = %v, want exactly one flagged drop (neither band is set_dps_partial)", fake.errors)
	}
}

// TestCheckSetDPSMonotonicAcrossBandsAllowsAPartialExplainedDrop is
// the same shape, with both bands honestly marked SetDPSPartial
// (report.go's own doc: a stripped, not-in-sim main_hand left each
// band's own measurement incomplete) - the published reason this
// lane's brief asks for, so this must not be flagged.
func TestCheckSetDPSMonotonicAcrossBandsAllowsAPartialExplainedDrop(t *testing.T) {
	bands := []bandReport{
		{Spec: "rogue-assassination", Band: 30, Faction: "alliance", SetDPS: 47.68, SetDPSPartial: true},
		{Spec: "rogue-assassination", Band: 40, Faction: "alliance", SetDPS: 36.71, SetDPSPartial: true},
	}
	fake := &fakeTB{}
	checkSetDPSMonotonicAcrossBands(fake, "testdata/synthetic.json", bands)
	if len(fake.errors) != 0 {
		t.Fatalf("checkSetDPSMonotonicAcrossBands errors = %v, want none: both bands are set_dps_partial, a published reason for the drop", fake.errors)
	}
}

// TestCheckSetDPSMonotonicAcrossBandsIgnoresOtherFactionsAndIncreases
// confirms two more cases in the same table: a genuine increase never
// flags (whatever SetDPSPartial says), and two factions are compared
// independently - a horde band's own lower SetDPS never gets compared
// against an unrelated alliance band's.
func TestCheckSetDPSMonotonicAcrossBandsIgnoresOtherFactionsAndIncreases(t *testing.T) {
	bands := []bandReport{
		{Spec: "hunter-marksmanship", Band: 20, Faction: "alliance", SetDPS: 100},
		{Spec: "hunter-marksmanship", Band: 20, Faction: "horde", SetDPS: 10},
		{Spec: "hunter-marksmanship", Band: 30, Faction: "alliance", SetDPS: 150},
		{Spec: "hunter-marksmanship", Band: 30, Faction: "horde", SetDPS: 20},
	}
	fake := &fakeTB{}
	checkSetDPSMonotonicAcrossBands(fake, "testdata/synthetic.json", bands)
	if len(fake.errors) != 0 {
		t.Fatalf("checkSetDPSMonotonicAcrossBands errors = %v, want none: every band increased within its own faction", fake.errors)
	}
}

// TestCheckSetDPSMonotonicAcrossBandsComparesARaidEntryToItsBareTwin: the
// raid entry is held to the bare entry of the same band, and does not
// become the baseline the next band is compared against.
func TestCheckSetDPSMonotonicAcrossBandsComparesARaidEntryToItsBareTwin(t *testing.T) {
	bands := []bandReport{
		{Spec: "hunter-marksmanship", Band: 60, Faction: "alliance", Preset: presetBare, SetDPS: 100},
		{Spec: "hunter-marksmanship", Band: 60, Faction: "alliance", Preset: presetRaid, SetDPS: 150},
	}
	fake := &fakeTB{}
	checkSetDPSMonotonicAcrossBands(fake, "testdata/synthetic.json", bands)
	if len(fake.errors) != 0 {
		t.Fatalf("errors = %v, want none: the raid entry is above its bare twin", fake.errors)
	}
	bands[1].SetDPS = 80
	fake = &fakeTB{}
	checkSetDPSMonotonicAcrossBands(fake, "testdata/synthetic.json", bands)
	if len(fake.errors) != 1 {
		t.Fatalf("errors = %v, want one: a raid entry below its bare twin is a bug", fake.errors)
	}
}
