package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/leveling"
)

// launchBannedPath lists the items no level-60 pick or alternative may name
// at launch, each with the content-phase gate that bans it.
const launchBannedPath = "testdata/launch-banned.json"

// bannedWhileClosed is launch-banned.json's item id -> gate key, for every
// gate data/curated/content-phases.json still closes at launch: a gate the
// owner opens unbans its items with no edit to the list. It also checks
// that the list names only real gates and gives every closed gate an
// example.
func bannedWhileClosed(t *testing.T) map[int]string {
	t.Helper()
	b, err := os.ReadFile(launchBannedPath)
	if err != nil {
		t.Fatalf("reading %s: %v", launchBannedPath, err)
	}
	var list struct {
		Items []struct {
			ID   int    `json:"id"`
			Gate string `json:"gate"`
		} `json:"items"`
	}
	if err := json.Unmarshal(b, &list); err != nil {
		t.Fatalf("decoding %s: %v", launchBannedPath, err)
	}
	g, err := loadContentGates(publishedRepoRoot)
	if err != nil {
		t.Fatalf("loadContentGates: %v", err)
	}
	examples := map[string]bool{}
	banned := map[int]string{}
	for _, row := range list.Items {
		closed, known := g.gateClosed[row.Gate]
		if !known {
			t.Fatalf("%s bans item %d under unknown gate %q", launchBannedPath, row.ID, row.Gate)
		}
		examples[row.Gate] = true
		if closed {
			banned[row.ID] = row.Gate
		}
	}
	for key, closed := range g.gateClosed {
		if closed && !examples[key] {
			t.Errorf("%s has no example item for the closed gate %q", launchBannedPath, key)
		}
	}
	return banned
}

func activeBuildDir(t *testing.T) string {
	t.Helper()
	build, err := leveling.ReadActiveBuild(publishedRepoRoot)
	if err != nil {
		t.Fatalf("ReadActiveBuild: %v", err)
	}
	return filepath.Join(publishedRepoRoot, "data", "builds", build)
}

// band60Entry is one published band-60 pick, alternative or runner-up.
type band60Entry struct {
	where string
	row   runnerUpRow
}

// publishedBand60Entries reads the active build's bis/*.json and returns
// every band-60 entry, both factions, every preset. The committed files
// are the nightly's output, so the two tests that read them run in its
// sanity step (FOREVER_BIS_SANITY=1), after the ranking that applied
// tonight's gates; TestNoBannedItemIsObtainableAtLaunch guards the
// ranker's own input on every run.
func publishedBand60Entries(t *testing.T) []band60Entry {
	t.Helper()
	if os.Getenv("FOREVER_BIS_SANITY") == "" {
		t.Skip("set FOREVER_BIS_SANITY=1 to check the committed bis files (the nightly does)")
	}
	files, err := filepath.Glob(filepath.Join(activeBuildDir(t), "bis", "*.json"))
	if err != nil {
		t.Fatalf("glob bis/*.json: %v", err)
	}
	if len(files) == 0 {
		t.Skip("no bis files published for the active build")
	}
	var entries []band60Entry
	factions := map[string]bool{}
	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatalf("reading %s: %v", f, err)
		}
		var report specReport
		if err := json.Unmarshal(b, &report); err != nil {
			t.Fatalf("decoding %s: %v", f, err)
		}
		for _, band := range report.Bands {
			if band.Band != 60 {
				continue
			}
			factions[band.Faction] = true
			for _, row := range band.Slots {
				if row.ItemID == 0 {
					continue
				}
				where := fmt.Sprintf("%s band 60 %s %s %s", report.Spec, band.Faction, band.Preset, row.Slot)
				entries = append(entries, band60Entry{where, runnerUpRow{ItemID: row.ItemID, SourceKind: row.SourceKind, Source: row.Source}})
				for _, a := range row.Alternatives {
					entries = append(entries, band60Entry{where, runnerUpRow{ItemID: a.ItemID, SourceKind: a.SourceKind, Source: a.Source}})
				}
				for _, r := range row.RunnersUp {
					entries = append(entries, band60Entry{where, r})
				}
			}
		}
	}
	if len(entries) == 0 || !factions["alliance"] || !factions["horde"] {
		t.Fatalf("read %d band-60 entries over factions %v, want both factions' picks", len(entries), factions)
	}
	return entries
}

// TestNoBannedItemIsAPublishedBand60Pick: no published band-60 pick,
// alternative or runner-up, in any spec, preset or faction, is an item
// launch-banned.json bans while its gate is closed.
func TestNoBannedItemIsAPublishedBand60Pick(t *testing.T) {
	banned := bannedWhileClosed(t)
	for _, e := range publishedBand60Entries(t) {
		if gate, ok := banned[e.row.ItemID]; ok {
			t.Errorf("%s: item %d (%s %q) is banned at launch by gate %s", e.where, e.row.ItemID, e.row.SourceKind, e.row.Source, gate)
		}
	}
}

// TestTheBand60PicksComeOnlyFromContentOpenAtLaunch is task L1.7's ruling
// 2, the phase rule at band 60: no published band-60 entry is a raid drop,
// and every one is named by at least one source that is still open once
// the curated content-phase gates are applied - not merely one whose
// loot.json `opens` happens to be empty.
func TestTheBand60PicksComeOnlyFromContentOpenAtLaunch(t *testing.T) {
	entries := publishedBand60Entries(t)
	idx := activeGatedLootIndex(t)
	for _, e := range entries {
		if msg := band60PhaseError(idx, e.row); msg != "" {
			t.Errorf("%s: %s", e.where, msg)
		}
	}
}

// band60PhaseError is the phase rule for one published band-60 entry (""
// when it holds): never a raid drop, and at least one source of the
// published kind - the one with the published label when there is one -
// is open after the content-phase gates.
func band60PhaseError(idx lootIndex, e runnerUpRow) string {
	if e.SourceKind == "raid" {
		return fmt.Sprintf("item %d is a raid drop (%s)", e.ItemID, e.Source)
	}
	var labelled, ofKind []itemSource
	for _, s := range idx[e.ItemID] {
		if s.Kind != e.SourceKind {
			continue
		}
		ofKind = append(ofKind, s)
		if s.Label == e.Source {
			labelled = append(labelled, s)
		}
	}
	candidates := labelled
	if len(candidates) == 0 {
		candidates = ofKind
	}
	if len(candidates) == 0 {
		return fmt.Sprintf("item %d: no %s source gives it", e.ItemID, e.SourceKind)
	}
	var opens []string
	for _, s := range candidates {
		if s.Opens == "" {
			return ""
		}
		opens = append(opens, s.Opens)
	}
	return fmt.Sprintf("item %d: every %s source %q opens after launch (%s)", e.ItemID, e.SourceKind, e.Source, strings.Join(opens, ", "))
}

func TestTheBand60PhaseRuleRefusesARaidDropAndALaterPhasesSource(t *testing.T) {
	idx := lootIndex{
		1: {{Kind: "world", Label: "Lord Kazzak", Opens: "raids-1"}},
		2: {{Kind: "world", Label: "Lord Skwol"}, {Kind: "world", Label: "Lord Kazzak", Opens: "raids-1"}},
		3: {{Kind: "dungeon", Label: "Stratholme: Baron Rivendare"}, {Kind: "dungeon", Label: "Stratholme: Balzaphon", Opens: "later"}},
		6: {{Kind: "vendor", Label: "Anachronos", Opens: "later"}},
		7: {{Kind: "quest", Label: "Treasure of the Timeless One"}, {Kind: "quest", Label: "Treasure of the Timeless One", Opens: "later"}},
	}
	cases := []struct {
		e  runnerUpRow
		ok bool
	}{
		{runnerUpRow{ItemID: 1, SourceKind: "world", Source: "Lord Kazzak"}, false},
		{runnerUpRow{ItemID: 2, SourceKind: "world", Source: "Lord Skwol"}, true},
		{runnerUpRow{ItemID: 3, SourceKind: "dungeon", Source: "Stratholme: Baron Rivendare"}, true},
		{runnerUpRow{ItemID: 3, SourceKind: "dungeon", Source: "Stratholme: Balzaphon"}, false},
		{runnerUpRow{ItemID: 4, SourceKind: "raid", Source: "Molten Core: Ragnaros"}, false},
		{runnerUpRow{ItemID: 5, SourceKind: "crafted", Source: "Blacksmithing"}, false},
		{runnerUpRow{ItemID: 6, SourceKind: "vendor", Source: "Anachronos"}, false},
		// A bare duplicate is still open in a raw index; loadGatedLootIndex
		// is what closes it (TestContentGatesCloseEveryCopyOfADuplicateQuest).
		{runnerUpRow{ItemID: 7, SourceKind: "quest", Source: "Treasure of the Timeless One"}, true},
	}
	for _, c := range cases {
		if msg := band60PhaseError(idx, c.e); (msg == "") != c.ok {
			t.Errorf("band60PhaseError(%+v) = %q, want ok=%v", c.e, msg, c.ok)
		}
	}
}
