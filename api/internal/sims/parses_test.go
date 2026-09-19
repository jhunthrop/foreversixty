package sims

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// seedParseRow writes one ranked, killing fight row in the given
// phase, and the stored summary TopParses reads the combatant and
// casts from - the ranking half and the bucket half of one parse.
func seedParseRow(h *harness, reportID string, index int, key, name, spec, phase string, dps float64, sum summary.Summary) {
	h.t.Helper()
	if _, err := h.store.Pool.Exec(h.t.Context(),
		`insert into fight_metrics (report_id, fight_index, player_key, player_name, class, spec,
		   role, metric_dps, duration_ms, kill, phase, fought_at, state)
		 values ($1, $2, $3, $4, 'warrior', $5, 'dps', $6, 180000, true, $7, now(), 'ok')`,
		reportID, index, key, name, spec, dps, phase); err != nil {
		h.t.Fatal(err)
	}
	sum.FightIndex = index
	body, err := json.Marshal(sum)
	if err != nil {
		h.t.Fatal(err)
	}
	if err := h.files.Put(h.t.Context(), store.Keys{ReportID: reportID}.FightSummary(index),
		body, ResultPut); err != nil {
		h.t.Fatal(err)
	}
}

// aParseSummary is a fight's stored summary carrying one combatant
// and one cast row for them, the shape TopParses reads.
func aParseSummary(name string) summary.Summary {
	return summary.Summary{
		Combatants: []summary.CombatantRow{{GUID: "Player-1", Name: name}},
		Casts: []summary.CastRow{
			{GUID: "Player-1", SpellID: 23881, SpellName: "Bloodthirst", Succeeded: 41},
		},
	}
}

func TestTopParsesReadsRankedFightsHighestDPSFirst(t *testing.T) {
	h := newHarness(t)
	ensureMetricsPartition(h)
	seedParseRow(h, "rep1", 1, "us/normal/a", "Baelgrim", "warrior-fury", "raids-1", 1000, aParseSummary("Baelgrim"))
	seedParseRow(h, "rep2", 1, "us/normal/b", "Otherguy", "warrior-fury", "raids-1", 1200, aParseSummary("Otherguy"))
	// A different phase's fight is not counted, even on the same spec.
	seedParseRow(h, "rep3", 1, "us/normal/c", "Thirdguy", "warrior-fury", "beta", 5000, aParseSummary("Thirdguy"))

	reader := &ParseReader{Pool: h.store.Pool, Get: dirGetter{root: h.dir}}
	got, err := reader.TopParses(t.Context(), "warrior-fury", "raids-1", 50)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("%d parses, want 2", len(got))
	}
	if got[0].ActualDPS != 1200 || got[0].PlayerKey != "us/normal/b" {
		t.Fatalf("highest dps should sort first: %+v", got[0])
	}
	if got[0].DurationSec != 180 {
		t.Errorf("duration %d, want 180", got[0].DurationSec)
	}
	if got[0].Spec != "warrior-fury" || got[0].Class != "warrior" {
		t.Errorf("spec/class %+v", got[0])
	}
	if got[0].Combatant.Name != "Otherguy" {
		t.Errorf("combatant %+v", got[0].Combatant)
	}
	c, ok := got[0].ActualCasts[23881]
	if !ok || c.Casts != 41 || c.Name != "Bloodthirst" {
		t.Errorf("actual casts %+v", got[0].ActualCasts)
	}
}

func TestTopParsesSkipsAFightWithNoStoredSummary(t *testing.T) {
	h := newHarness(t)
	ensureMetricsPartition(h)
	// A ranked row with no object ever written for it: the bucket
	// read fails, and the parse is skipped rather than failing the
	// whole read.
	if _, err := h.store.Pool.Exec(h.t.Context(),
		`insert into fight_metrics (report_id, fight_index, player_key, player_name, class, spec,
		   role, metric_dps, duration_ms, kill, phase, fought_at, state)
		 values ('repmissing', 1, 'us/normal/gone', 'Ghost', 'warrior', 'warrior-fury', 'dps', 900,
		   180000, true, 'raids-1', now(), 'ok')`); err != nil {
		t.Fatal(err)
	}
	reader := &ParseReader{Pool: h.store.Pool, Get: dirGetter{root: h.dir}}
	got, err := reader.TopParses(t.Context(), "warrior-fury", "raids-1", 50)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Fatalf("%d parses, want 0: the summary could not be read", len(got))
	}
}

func TestTopParsesSkipsAFightWhoseSummaryHasNoCombatantForThePlayer(t *testing.T) {
	h := newHarness(t)
	ensureMetricsPartition(h)
	seedParseRow(h, "repempty", 1, "us/normal/nobody", "Nobody", "warrior-fury", "raids-1", 1000,
		summary.Summary{Combatants: []summary.CombatantRow{{GUID: "Player-1", Name: "SomeoneElse"}}})
	reader := &ParseReader{Pool: h.store.Pool, Get: dirGetter{root: h.dir}}
	got, err := reader.TopParses(t.Context(), "warrior-fury", "raids-1", 50)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Fatalf("%d parses, want 0: no combatant row for the player", len(got))
	}
}

func TestTopParsesRespectsTheLimit(t *testing.T) {
	h := newHarness(t)
	ensureMetricsPartition(h)
	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("Name%d", i)
		seedParseRow(h, fmt.Sprintf("repn%d", i), 1, fmt.Sprintf("us/normal/n%d", i), name,
			"warrior-fury", "raids-1", float64(1000+i), aParseSummary(name))
	}
	reader := &ParseReader{Pool: h.store.Pool, Get: dirGetter{root: h.dir}}
	got, err := reader.TopParses(t.Context(), "warrior-fury", "raids-1", 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("%d parses, want 2 (the limit)", len(got))
	}
}
