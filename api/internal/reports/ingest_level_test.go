package reports

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	testCharacterLevel = 60
	testItemLevel      = 187
)

func rosterRow(level, itemLevel int64) summary.Summary {
	return summary.Summary{Roster: []summary.RosterRow{{GUID: "Player-1", Name: "A", Level: level, ItemLevel: itemLevel}}}
}

func TestOverlayTakesTheCharacterLevelFromTheRecord(t *testing.T) {
	posted := rosterRow(0, 0)
	overlayDerivedRoster(&posted, rosterRow(0, 0), map[string]int64{"Player-1": testCharacterLevel})
	if got := posted.Roster[0].Level; got != testCharacterLevel {
		t.Fatalf("level = %d, want %d from the character record", got, testCharacterLevel)
	}
}

func TestOverlayNeverTakesTheCharacterLevelFromTheLog(t *testing.T) {
	posted := rosterRow(0, 0)
	overlayDerivedRoster(&posted, rosterRow(testCharacterLevel, 0), nil)
	if got := posted.Roster[0].Level; got != 0 {
		t.Fatalf("level = %d, want 0 with no matched character", got)
	}
}

func TestOverlayFillsItemLevelFromTheRebuildAndKeepsAPostedOne(t *testing.T) {
	blank, kept := rosterRow(0, 0), rosterRow(0, 150)
	overlayDerivedRoster(&blank, rosterRow(0, testItemLevel), nil)
	overlayDerivedRoster(&kept, rosterRow(0, testItemLevel), nil)
	if blank.Roster[0].ItemLevel != testItemLevel || kept.Roster[0].ItemLevel != 150 {
		t.Fatalf("item levels = %d and %d, want %d and 150",
			blank.Roster[0].ItemLevel, kept.Roster[0].ItemLevel, testItemLevel)
	}
}

// fakeLevels answers character levels by character key.
type fakeLevels map[string]int64

func (f fakeLevels) CharacterLevels(_ context.Context, keys []string) (map[string]int64, error) {
	out := map[string]int64{}
	for _, k := range keys {
		if f[k] > 0 {
			out[k] = f[k]
		}
	}
	return out, nil
}

// erroringLevels fails every lookup.
type erroringLevels struct{}

func (erroringLevels) CharacterLevels(context.Context, []string) (map[string]int64, error) {
	return nil, errors.New("levels store down")
}

// levelsUnder keys testCharacterLevel for every bundle roster name, in the region of
// report id: under the ruleset the name resolves to, or, when otherRuleset, under a
// different one.
func levelsUnder(t *testing.T, h *harness, id string, otherRuleset bool) fakeLevels {
	t.Helper()
	rep, err := h.store.Get(t.Context(), id)
	if err != nil {
		t.Fatal(err)
	}
	region, ruleset := ReportRealm(rep)
	levels := fakeLevels{}
	for _, row := range makeBundle(t, 1, nil).rows {
		levels[levelKey(region, ruleset, row.Name, otherRuleset)] = testCharacterLevel
	}
	return levels
}

// levelKey is the key the ingest derives for unit, or one under a different ruleset.
func levelKey(region, ruleset, unit string, otherRuleset bool) string {
	derived := character.KeyFromUnit(region, ruleset, unit)
	if !otherRuleset {
		return derived
	}
	name, _ := character.SplitUnit(unit)
	for _, r := range character.Rulesets {
		if key := character.Key(region, r, name); key != derived {
			return key
		}
	}
	return derived
}

// ingestStoredRoster ingests fight 1 with the harness's Levels hook as set, the roster
// posted with item level testItemLevel and no level, and returns the stored roster.
func ingestStoredRoster(t *testing.T, h *harness, id string) []summary.RosterRow {
	t.Helper()
	b := makeBundleWith(t, func(parts map[string][]byte) {
		var sum summary.Summary
		if err := json.Unmarshal(parts["summary"], &sum); err != nil {
			t.Fatal(err)
		}
		for j := range sum.Roster {
			sum.Roster[j].ItemLevel, sum.Roster[j].Level = testItemLevel, 0
		}
		parts["summary"] = mustJSON(t, sum)
	})
	res := h.do(http.MethodPut, fmt.Sprintf("/v1/reports/%s/fights/1", id), b.contentType, b.body)
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("status = %d, want 201", res.StatusCode)
	}
	var stored summary.Summary
	decodeInto(t, h.readFile(store.Keys{ReportID: id}.FightSummary(1)), &stored)
	if len(stored.Roster) == 0 {
		t.Fatal("stored roster is empty")
	}
	return stored.Roster
}

func requireLevels(t *testing.T, roster []summary.RosterRow, wantLevel int64) {
	t.Helper()
	for _, row := range roster {
		if row.Level != wantLevel || row.ItemLevel != testItemLevel {
			t.Errorf("%s stored level %d item_level %d, want %d and %d",
				row.Name, row.Level, row.ItemLevel, wantLevel, testItemLevel)
		}
	}
}

func TestAFullIngestStoresTheCharacterLevelAndItemLevel(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	h.ingest.Levels = levelsUnder(t, h, id, false)
	requireLevels(t, ingestStoredRoster(t, h, id), testCharacterLevel)
}

func TestASameNameCharacterUnderAnotherRulesetGivesNoLevel(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	h.ingest.Levels = levelsUnder(t, h, id, true)
	requireLevels(t, ingestStoredRoster(t, h, id), 0)
}

func TestALevelsLookupErrorGivesNoLevelAndTheIngestStillSucceeds(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	h.ingest.Levels = erroringLevels{}
	requireLevels(t, ingestStoredRoster(t, h, id), 0)
}

func TestANilLevelsHookGivesNoLevel(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	h.ingest.Levels = nil
	requireLevels(t, ingestStoredRoster(t, h, id), 0)
}
