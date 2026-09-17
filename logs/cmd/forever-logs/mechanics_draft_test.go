// logs/cmd/forever-logs/mechanics_draft_test.go
package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
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
			// The evidence stays on the line: "check" without it tells the curator
			// nothing to check against.
			wantNote: "unclassified: check — draft: hit 1 players (1 non-tanks), 10% of damage taken, killed 0",
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
			wantNote: "unclassified: check — draft: hit 1 players (1 non-tanks), 100% of damage taken, killed 0",
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

// TestAddFightCountsPeriodicTicks pins the one thing the summary hands the
// drafter that is easy to miss: periodic damage lands in Ability.Ticks, not
// Ability.Hits. A DoT or a ground effect ticking twice on two non-tanks is
// avoidable; counting Hits alone saw "0 non-tanks" and drafted it as the
// fight's own damage.
func TestAddFightCountsPeriodicTicks(t *testing.T) {
	const p1, p2 = "Player-1-00000001", "Player-1-00000002"
	reg := units.NewRegistry(units.Options{})
	for _, guid := range []string{p1, p2} {
		reg.Observe(event.Event{
			Time:   time.Unix(0, 0),
			Kind:   event.Damage,
			Name:   "SPELL_DAMAGE",
			Source: event.Unit{GUID: guid, Name: guid, Flags: 0x512},
			Dest:   event.Unit{GUID: "Creature-0-1-1-1-1-1", Flags: 0xa48},
		})
	}
	taken := func(guid string) summary.Actor {
		return summary.Actor{GUID: guid, Name: guid, Effective: 400, Abilities: []summary.Ability{
			// A ground effect: every landing is a tick, none of them a Hit.
			{SpellID: 319685, Name: "Severing Smash", Effective: 400, Ticks: 2},
		}}
	}
	sum := summary.Summary{
		Roster:      []summary.RosterRow{{GUID: p1, Role: "dps"}, {GUID: p2, Role: "dps"}},
		DamageTaken: []summary.Actor{taken(p1), taken(p2)},
	}
	ed := &encounterDraft{name: "Kryxis the Voracious", spells: map[int64]*spellDraft{}}
	ed.addFight(sum, reg)

	sd := ed.spells[319685]
	if sd == nil {
		t.Fatalf("the ticking ability must be drafted: %+v", ed.spells)
	}
	if got := sd.players[p1].nonTankHits; got != 2 {
		t.Errorf("non-tank hits for %s = %d, want the 2 ticks", p1, got)
	}
	m := ed.table(2360).Mechanics[0]
	if m.Kind != mechanics.Avoidable {
		t.Errorf("kind = %q, want avoidable: it ticked on two non-tanks", m.Kind)
	}
	want := "draft: hit 2 players (2 non-tanks), 100% of damage taken, killed 0"
	if m.Note != want {
		t.Errorf("note = %q, want %q", m.Note, want)
	}
}

// The fixture log has one pull of Warden Kelthas. Its Anima Surge cast starts once and
// only once, so it is exactly the shape a phase trigger has; Frostbolt is a player's and
// is not a candidate at all.
func TestMechanicsDraftProposesPhaseCandidates(t *testing.T) {
	var out, errOut bytes.Buffer
	err := run([]string{"mechanics-draft", "-encounter", "9001",
		"../../../web/src/fixtures/report/fixture.log"}, &out, &errOut)
	if err != nil {
		t.Fatal(err, errOut.String())
	}
	table, err := mechanics.Parse(out.Bytes())
	if err != nil {
		t.Fatalf("the draft must be a valid table: %v\n%s", err, out.String())
	}
	var surge *mechanics.Phase
	for i := range table.Phases {
		if table.Phases[i].Starts.SpellID == 334653 {
			surge = &table.Phases[i]
		}
	}
	if surge == nil {
		t.Fatalf("Anima Surge must be a phase candidate: %+v", table.Phases)
	}
	if surge.Starts.On != mechanics.OnCastStart {
		t.Errorf("a channel's candidate must key on its start, not its success: %+v", surge.Starts)
	}
	if surge.Name != "Phase 2" {
		t.Errorf("the first candidate after the pull is Phase 2, got %q", surge.Name)
	}
	for _, p := range table.Phases {
		if p.Starts.SpellID == 116 {
			t.Errorf("a player's spell is not a phase candidate: %+v", p)
		}
	}
	if !strings.Contains(errOut.String(), "Phase 2") {
		t.Errorf("the evidence for each candidate must reach stderr:\n%s", errOut.String())
	}
}
