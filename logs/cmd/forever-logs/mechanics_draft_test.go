// logs/cmd/forever-logs/mechanics_draft_test.go
package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
)

func TestMechanicsDraftProposesATableFromTheFixtureLog(t *testing.T) {
	var out, errOut bytes.Buffer
	err := run([]string{"mechanics-draft", "-encounter", "9001", "../../../web/src/fixtures/report/fixture.log"}, &out, &errOut)
	if err != nil {
		t.Fatal(err, errOut.String())
	}
	table, err := mechanics.Parse(out.Bytes())
	if err != nil {
		t.Fatalf("the draft must be a valid table: %v\n%s", err, out.String())
	}
	if table.EncounterID != 9001 {
		t.Fatalf("encounter = %d", table.EncounterID)
	}
	var lash *mechanics.Mechanic
	for i := range table.Mechanics {
		if table.Mechanics[i].SpellID == 334660 {
			lash = &table.Mechanics[i]
		}
	}
	if lash == nil || !strings.HasPrefix(lash.Note, "draft:") {
		t.Fatalf("Anima Lash must be drafted with evidence: %+v", table.Mechanics)
	}
	_ = json.Valid
}

// TestSpellDraftClassifiesByHitNotByPlayer exercises spellDraft.draft
// through spellDraft.hit, the same path addFight uses. It is table-driven
// over the classification rules, including a player who tanked on one pull
// of an encounter and did not on another: tankness must be judged per hit,
// not overwritten for the whole player, or a tank swap between pulls would
// misclassify every one of that player's hits by whichever pull came last.
func TestSpellDraftClassifiesByHitNotByPlayer(t *testing.T) {
	type call struct {
		guid      string
		hits      int64
		effective int64
		isTank    bool
	}
	tests := []struct {
		name             string
		calls            []call
		totalDamageTaken int64
		wantKind         mechanics.Kind
		wantNote         string
	}{
		{
			name: "avoidable: hit two non-tank players",
			calls: []call{
				{guid: "p1", hits: 1, effective: 100, isTank: false},
				{guid: "p2", hits: 1, effective: 100, isTank: false},
			},
			totalDamageTaken: 200,
			wantKind:         mechanics.Avoidable,
			wantNote:         "draft: hit 2 players (2 non-tanks), 100% of damage taken, killed 0",
		},
		{
			name: "avoidable: hit one non-tank player twice",
			calls: []call{
				{guid: "p1", hits: 2, effective: 100, isTank: false},
			},
			totalDamageTaken: 100,
			wantKind:         mechanics.Avoidable,
			wantNote:         "draft: hit 1 players (1 non-tanks), 100% of damage taken, killed 0",
		},
		{
			name: "unavoidable: every hit landed on a tank",
			calls: []call{
				{guid: "tank", hits: 3, effective: 300, isTank: true},
			},
			totalDamageTaken: 300,
			wantKind:         mechanics.Unavoidable,
			wantNote:         "draft: hit 1 players (0 non-tanks), 100% of damage taken, killed 0",
		},
		{
			name: "unclassified: one non-tank hit once",
			calls: []call{
				{guid: "p1", hits: 1, effective: 50, isTank: false},
			},
			totalDamageTaken: 500,
			wantKind:         mechanics.Avoidable,
			wantNote:         "unclassified: check",
		},
		{
			name: "tank on one pull, not the tank on another: hits counted on each side",
			calls: []call{
				{guid: "p1", hits: 1, effective: 100, isTank: true},  // pull 1: p1 tanked it
				{guid: "p1", hits: 1, effective: 100, isTank: false}, // pull 2: someone else tanked; p1 took it as a non-tank
			},
			totalDamageTaken: 200,
			// One non-tank hit, once: too thin to call, same as any other
			// single non-tank hit. Before the fix, the second call's
			// isTank=false overwrote the first call's isTank=true for the
			// whole player, so both hits were counted as non-tank hits on
			// one player and this classified avoidable instead.
			wantKind: mechanics.Avoidable,
			wantNote: "unclassified: check",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sd := &spellDraft{name: "Test Spell", players: map[string]*playerHits{}}
			for _, c := range tt.calls {
				sd.hit(c.guid, c.hits, c.effective, c.isTank)
			}
			m := sd.draft(1, tt.totalDamageTaken)
			if m.Kind != tt.wantKind {
				t.Errorf("kind = %q, want %q", m.Kind, tt.wantKind)
			}
			if m.Note != tt.wantNote {
				t.Errorf("note = %q, want %q", m.Note, tt.wantNote)
			}
		})
	}
}

func TestPercentGuardsAZeroDenominator(t *testing.T) {
	if got := percent(50, 0); got != 0 {
		t.Fatalf("percent(50, 0) = %d, want 0", got)
	}
	if got := percent(50, 200); got != 25 {
		t.Fatalf("percent(50, 200) = %d, want 25", got)
	}
}
