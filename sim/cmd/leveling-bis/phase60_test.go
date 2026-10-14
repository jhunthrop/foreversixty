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

// TestTheBand60PicksComeOnlyFromContentOpenAtLaunch is task L1.7's ruling
// 2, the phase rule at band 60, committed as a test (fix round 1): over
// the active build's published bis/*.json, band 60, both factions, every
// preset, no pick, alternative or runner-up is a raid drop, and none is
// named by a dungeon or world source that opens only in a later phase -
// at least one loot source with the published label and no `opens` must
// give the item. A later phase's content (a world boss, a raid's
// reputation, a summoned boss) is never a level-60 character's pick at
// launch.
func TestTheBand60PicksComeOnlyFromContentOpenAtLaunch(t *testing.T) {
	build, err := leveling.ReadActiveBuild(publishedRepoRoot)
	if err != nil {
		t.Fatalf("ReadActiveBuild: %v", err)
	}
	buildDir := filepath.Join(publishedRepoRoot, "data", "builds", build)
	files, err := filepath.Glob(filepath.Join(buildDir, "bis", "*.json"))
	if err != nil {
		t.Fatalf("glob %s/bis/*.json: %v", buildDir, err)
	}
	if len(files) == 0 {
		t.Skipf("no bis files published for the active build %s", build)
	}
	idx, _, err := loadLootIndex(buildDir, nil)
	if err != nil {
		t.Fatalf("loadLootIndex(%s): %v", buildDir, err)
	}
	checked := 0
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
			where := fmt.Sprintf("%s band 60 %s %s", report.Spec, band.Faction, band.Preset)
			for _, row := range band.Slots {
				if row.ItemID == 0 {
					continue
				}
				entries := []runnerUpRow{{ItemID: row.ItemID, SourceKind: row.SourceKind, Source: row.Source}}
				for _, a := range row.Alternatives {
					entries = append(entries, runnerUpRow{ItemID: a.ItemID, SourceKind: a.SourceKind, Source: a.Source})
				}
				entries = append(entries, row.RunnersUp...)
				for _, e := range entries {
					checked++
					if msg := band60PhaseError(idx, e); msg != "" {
						t.Errorf("%s %s: %s", where, row.Slot, msg)
					}
				}
			}
		}
	}
	if checked == 0 || !factions["alliance"] || !factions["horde"] {
		t.Fatalf("checked %d band-60 entries over factions %v, want both factions' picks", checked, factions)
	}
}

// band60PhaseError is the phase rule for one published band-60 entry (""
// when it holds): never a raid drop, and a dungeon or world label must
// name at least one source with no `opens`.
func band60PhaseError(idx lootIndex, e runnerUpRow) string {
	if e.SourceKind == "raid" {
		return fmt.Sprintf("item %d is a raid drop (%s)", e.ItemID, e.Source)
	}
	if e.SourceKind != "dungeon" && e.SourceKind != "world" {
		return ""
	}
	var opens []string
	for _, s := range idx[e.ItemID] {
		if s.Kind != e.SourceKind || s.Label != e.Source {
			continue
		}
		if s.Opens == "" {
			return ""
		}
		opens = append(opens, s.Opens)
	}
	if len(opens) == 0 {
		return fmt.Sprintf("item %d: no %s source labelled %q gives it", e.ItemID, e.SourceKind, e.Source)
	}
	return fmt.Sprintf("item %d: its only %s source %q opens later (%s)", e.ItemID, e.SourceKind, e.Source, strings.Join(opens, ", "))
}

func TestTheBand60PhaseRuleRefusesARaidDropAndALaterPhasesSource(t *testing.T) {
	idx := lootIndex{
		1: {{Kind: "world", Label: "Lord Kazzak", Opens: "raids-1"}},
		2: {{Kind: "world", Label: "Lord Skwol"}, {Kind: "world", Label: "Lord Kazzak", Opens: "raids-1"}},
		3: {{Kind: "dungeon", Label: "Stratholme: Baron Rivendare"}},
	}
	cases := []struct {
		e  runnerUpRow
		ok bool
	}{
		{runnerUpRow{ItemID: 1, SourceKind: "world", Source: "Lord Kazzak"}, false},
		{runnerUpRow{ItemID: 2, SourceKind: "world", Source: "Lord Skwol"}, true},
		{runnerUpRow{ItemID: 3, SourceKind: "dungeon", Source: "Stratholme: Baron Rivendare"}, true},
		{runnerUpRow{ItemID: 3, SourceKind: "dungeon", Source: "Scholomance"}, false},
		{runnerUpRow{ItemID: 4, SourceKind: "raid", Source: "Molten Core: Ragnaros"}, false},
		{runnerUpRow{ItemID: 5, SourceKind: "crafted", Source: "Blacksmithing"}, true},
	}
	for _, c := range cases {
		if msg := band60PhaseError(idx, c.e); (msg == "") != c.ok {
			t.Errorf("band60PhaseError(%+v) = %q, want ok=%v", c.e, msg, c.ok)
		}
	}
}
