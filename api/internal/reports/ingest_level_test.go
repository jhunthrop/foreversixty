package reports

import (
	"context"
	"encoding/json"
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

func TestAFullIngestStoresTheCharacterLevelAndItemLevel(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	rep, err := h.store.Get(t.Context(), id)
	if err != nil {
		t.Fatal(err)
	}
	region, ruleset := ReportRealm(rep)
	levels := fakeLevels{}
	for _, row := range makeBundle(t, 1, nil).rows {
		levels[character.KeyFromUnit(region, ruleset, row.Name)] = testCharacterLevel
	}
	h.ingest.Levels = levels

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
	for _, row := range stored.Roster {
		if row.Level != testCharacterLevel || row.ItemLevel != testItemLevel {
			t.Errorf("%s stored level %d item_level %d, want %d and %d",
				row.Name, row.Level, row.ItemLevel, testCharacterLevel, testItemLevel)
		}
	}
}
