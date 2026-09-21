package adapter

import (
	"bytes"
	"encoding/json"
	"errors"
	"math"
	"reflect"
	"slices"
	"sort"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/wowsims/classic/sim/core/proto"
)

// oneAction builds a UnitMetrics with a single ability whose numbers are
// chosen so that the per-iteration division is visible: 200 casts over 100
// iterations is 2 casts per fight, not 200.
func oneAction() *proto.UnitMetrics {
	return &proto.UnitMetrics{
		Name:      "Fury",
		UnitIndex: 0,
		Dps:       &proto.DistributionMetrics{Avg: 1791.1, Stdev: 120.5, Max: 2100, Min: 1400},
		Actions: []*proto.ActionMetrics{{
			Id:          &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 23894}},
			IsMelee:     true,
			SpellSchool: 1,
			Targets: []*proto.TargetedActionMetrics{{
				UnitIndex:      1,
				Casts:          200,
				Hits:           160,
				Crits:          40,
				Misses:         20,
				Dodges:         10,
				Parries:        5,
				Glances:        5,
				Damage:         100000,
				CritDamage:     40000,
				ResistedDamage: 2000,
			}},
		}},
		Auras: []*proto.AuraMetrics{{
			Id:               &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 12966}},
			UptimeSecondsAvg: 140.4,
			ProcsAvg:         31.5,
		}},
		Resources: []*proto.ResourceMetrics{{
			Id:         &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 23894}},
			Type:       proto.ResourceType_ResourceTypeRage,
			Events:     200,
			Gain:       -6000,
			ActualGain: -6000,
		}},
	}
}

func resultWith(u *proto.UnitMetrics, iterations int32) *proto.RaidSimResult {
	return &proto.RaidSimResult{
		RaidMetrics: &proto.RaidMetrics{
			Dps:     &proto.DistributionMetrics{Avg: 1791.1, Stdev: 120.5, Max: 2100, Min: 1400},
			Parties: []*proto.PartyMetrics{{Players: []*proto.UnitMetrics{u}}},
		},
		AvgIterationDuration: 180.0,
		IterationsDone:       iterations,
	}
}

func req() api.SimRequest {
	return api.SimRequest{
		EngineVersion: enginever.Version,
		Spec:          "warrior-fury",
		Source:        api.CharacterSource{Kind: api.SourceManual},
		Character:     api.CharacterSpec{Name: "Fury", Race: "orc", Class: "warrior", Level: 60},
		Encounter:     api.DefaultEncounter(),
		Iterations:    100,
	}
}

func TestSummarizeHeader(t *testing.T) {
	got, err := Summarize(resultWith(oneAction(), 100), req())
	if err != nil {
		t.Fatal(err)
	}
	if want := "sim:" + enginever.Version; got.EngineVersion != want {
		t.Errorf("EngineVersion = %q, want %q", got.EngineVersion, want)
	}
	if got.FightIndex != 1 {
		t.Errorf("FightIndex = %d, want 1", got.FightIndex)
	}
	// DurationMS is now DERIVED (task 3), not the measured
	// avg_iteration_duration: it is the duration that makes total
	// damage / duration equal DPS(res).Mean exactly. oneAction()'s
	// actor totals 1000 (100000 damage over 100 iterations) and
	// resultWith wires RaidMetrics.Dps.Avg to 1791.1, so the derived
	// duration is round(1000/1791.1*1000) = 558, not the 180000 that
	// avg_iteration_duration (180s) alone would give - this fixture's
	// two numbers were never meant to agree; TestGoldenSummaries is
	// where headline-equals-table is actually proved, against real
	// engine output.
	if got.DurationMS != 558 {
		t.Errorf("DurationMS = %d, want 558 (derived: total damage / DPS.Mean)", got.DurationMS)
	}
}

// The engine accumulates across iterations; the summary describes one
// fight. Everything countable is divided by IterationsDone.
func TestSummarizeDividesByIterations(t *testing.T) {
	got, err := Summarize(resultWith(oneAction(), 100), req())
	if err != nil {
		t.Fatal(err)
	}
	if len(got.DamageDone) != 1 {
		t.Fatalf("DamageDone has %d actors, want 1", len(got.DamageDone))
	}
	a := got.DamageDone[0]
	if a.Name != "Fury" {
		t.Errorf("actor Name = %q, want %q", a.Name, "Fury")
	}
	if a.Class != "warrior" {
		t.Errorf("actor Class = %q, want %q; the roster and the damage table must agree", a.Class, "warrior")
	}
	// 100000 damage over 100 iterations. The brief's draft expected 1420
	// here, from damage + crit_damage + resisted_damage; the engine's
	// crit and resisted figures are breakdowns of damage rather than
	// additions to it, so that sum counts the same crit twice. sim/core
	// itself totals a unit's DPS from damage alone.
	if a.Total != 1000 {
		t.Errorf("actor Total = %d, want 1000 (damage is already every outcome's)", a.Total)
	}
	if a.Effective != a.Total {
		t.Errorf("Effective = %d, want it equal to Total (%d)", a.Effective, a.Total)
	}
	// See TestSummarizeHeader: the fight duration is now derived, not
	// avg_iteration_duration, so ActiveMS follows the same 558.
	if a.ActiveMS != 558 {
		t.Errorf("ActiveMS = %d, want the derived fight duration 558", a.ActiveMS)
	}
	if len(a.Abilities) != 1 {
		t.Fatalf("actor has %d abilities, want 1", len(a.Abilities))
	}
	ab := a.Abilities[0]
	if ab.SpellID != 23894 {
		t.Errorf("SpellID = %d, want 23894", ab.SpellID)
	}
	if ab.Name != "spell:23894" {
		t.Errorf("Name = %q; the engine carries no names, so the adapter emits the id and the web resolves it", ab.Name)
	}
	// Hits is every landing: 160 plain + 40 crits + 5 glances = 205
	// over 100 iterations, which rounds to 2. See
	// TestHitsAndCritsUseTheLogsEngineDefinitions for why.
	if ab.Hits != 2 {
		t.Errorf("Hits = %d, want 2 ((160+40+5)/100 rounded)", ab.Hits)
	}
	if ab.Crits != 0 { // 40/100 = 0.4, rounds to 0
		t.Errorf("Crits = %d, want 0 (40/100 rounds down)", ab.Crits)
	}
	// Resisted means the damage a resist took away, and the engine
	// tracks no such figure: its resisted_damage is what a partially
	// resisted cast still dealt, which is a different number with the
	// opposite sense. The adapter reports nothing rather than that.
	if ab.Resisted != 0 {
		t.Errorf("Resisted = %d, want 0; the engine reports no resisted amount", ab.Resisted)
	}
	// The keys are the combat log's own uppercase miss types, so a sim
	// and a real fight aggregate together; a count that divides to zero
	// writes no key at all.
	if ab.Misses["MISS"] != 0 { // 20/100 = 0.2
		t.Errorf("misses = %d, want 0", ab.Misses["MISS"])
	}
	if _, ok := ab.Misses["DODGE"]; ok { // 10/100 = 0.1
		t.Error(`Misses carries a "DODGE" key for an outcome that divides to zero`)
	}
	if len(a.Targets) != 1 || a.Targets[0].Total != 1000 {
		t.Errorf("Targets = %+v, want one entry totalling 1000", a.Targets)
	}
}

// Rounding must not lose a whole ability: an ability cast every third
// fight must still appear, with its damage, rather than vanish.
func TestSummarizeKeepsRareAbilities(t *testing.T) {
	u := oneAction()
	u.Actions = append(u.Actions, &proto.ActionMetrics{
		Id:      &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 20572}},
		IsMelee: false,
		Targets: []*proto.TargetedActionMetrics{{UnitIndex: 1, Casts: 33, Hits: 33, Damage: 3300}},
	})
	got, err := Summarize(resultWith(u, 100), req())
	if err != nil {
		t.Fatal(err)
	}
	var found bool
	for _, ab := range got.DamageDone[0].Abilities {
		if ab.SpellID == 20572 {
			found = true
			if ab.Total != 33 {
				t.Errorf("rare ability Total = %d, want 33", ab.Total)
			}
		}
	}
	if !found {
		t.Error("an ability used in a third of iterations vanished from the table")
	}
}

func TestSummarizeAurasCastsAndResources(t *testing.T) {
	got, err := Summarize(resultWith(oneAction(), 100), req())
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Auras) != 1 {
		t.Fatalf("Auras has %d entries, want 1", len(got.Auras))
	}
	au := got.Auras[0]
	if au.SpellID != 12966 {
		t.Errorf("aura SpellID = %d, want 12966", au.SpellID)
	}
	if au.UptimeMS != 140400 {
		t.Errorf("aura UptimeMS = %d, want 140400", au.UptimeMS)
	}
	if au.Applications != 32 { // 31.5 rounds to 32
		t.Errorf("aura Applications = %d, want 32", au.Applications)
	}
	if au.Segments == nil || len(au.Segments) != 0 {
		t.Errorf("aura Segments = %v; the engine reports no application timeline, so this must be empty and non-nil", au.Segments)
	}
	// The logs engine's vocabulary is upper case - its accumulator
	// writes "BUFF" and "DEBUFF" - and the report's shared AuraTable
	// filters on those two strings. A lower-case "buff" here renders as
	// an aura of no known type and the filter drops the row.
	if au.Type != "BUFF" {
		t.Errorf("aura Type = %q, want %q; the web's AuraTable filters on the logs engine's upper-case vocabulary", au.Type, "BUFF")
	}

	if len(got.Casts) != 1 {
		t.Fatalf("Casts has %d entries, want 1", len(got.Casts))
	}
	c := got.Casts[0]
	if c.SpellID != 23894 || c.Started != 2 || c.Succeeded != 2 {
		t.Errorf("cast row = %+v, want spell 23894 started and succeeded 2", c)
	}
	if c.OwnerGUID != c.GUID {
		t.Errorf("cast OwnerGUID = %q, want the caster's own GUID %q for a player row", c.OwnerGUID, c.GUID)
	}
	if c.Sequence == nil || len(c.Sequence) != 0 {
		t.Errorf("cast Sequence = %v; the engine reports no cast timestamps, so this must be empty and non-nil", c.Sequence)
	}

	if len(got.Resources) != 1 {
		t.Fatalf("Resources has %d entries, want 1", len(got.Resources))
	}
	r := got.Resources[0]
	// The LOG's own power-type id for Rage (1), not the engine's own enum
	// value (proto.ResourceType_ResourceTypeRage == 3) - 2026-09-21
	// result-page review round 3, E7: this assertion used to pin the raw
	// engine value straight through, which is exactly how a Fury Warrior's
	// RESOURCES tab came to read "Energy" for rage (POWER_NAMES is keyed
	// by the log's numbering, and 3 there is Energy). See
	// engineResourceTypeToLogPowerType's own comment (adapter.go) for both
	// enums in full.
	if r.PowerType != 1 {
		t.Errorf("resource PowerType = %d, want 1 (the log's own Rage id)", r.PowerType)
	}
	if r.Spent != 60 {
		t.Errorf("resource Spent = %d, want 60 (6000 spent over 100 iterations)", r.Spent)
	}
	if r.Gained != 0 {
		t.Errorf("resource Gained = %d, want 0", r.Gained)
	}
	if r.Max != 0 || r.AtMaxMS != 0 || r.ZeroMS != 0 {
		t.Errorf("resource cap fields = %d/%d/%d; the engine reports no cap or timeline, so all three stay zero", r.Max, r.AtMaxMS, r.ZeroMS)
	}
}

// gain minus actual_gain is exactly the resource the engine threw away
// over the cap, which is what ResourceTrack.Wasted means.
func TestSummarizeReportsWastedResource(t *testing.T) {
	u := oneAction()
	u.Resources = []*proto.ResourceMetrics{{
		Id:         &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 2687}},
		Type:       proto.ResourceType_ResourceTypeRage,
		Events:     100,
		Gain:       5000,
		ActualGain: 4000,
	}}
	got, err := Summarize(resultWith(u, 100), req())
	if err != nil {
		t.Fatal(err)
	}
	r := got.Resources[0]
	if r.Gained != 40 {
		t.Errorf("Gained = %d, want 40 (actual_gain over 100 iterations)", r.Gained)
	}
	if r.Wasted != 10 {
		t.Errorf("Wasted = %d, want 10 ((gain - actual_gain) over 100 iterations)", r.Wasted)
	}
}

// 2026-09-21 result-page review round 3, E7: PowerType must be the LOG's
// own power-type numbering (what web/src/components/report/
// ResourceGraphs.svelte's POWER_NAMES is keyed by), never the engine's
// own proto.ResourceType value handed through unchanged - the two enums
// disagree on every resource but ComboPoints, and Rage/Energy are each
// other's engine id, which is exactly how a Fury Warrior's own rage read
// as "Energy". Covers every resource a class in this ruleset actually
// uses: Rage (warrior/druid feral/rogue-adjacent... actually rogues and
// cat druids use Energy, warriors and bear druids use Rage), Energy,
// Mana (every caster, and hunters too - sim/hunter/hunter.go's own
// EnableManaBar, not Focus, in this ruleset) and ComboPoints. Focus is
// exercised directly against the translation table alone (TestLogPowerType
// below), not through a real class, because none uses it here.
func TestSummarizeTranslatesEveryResourceTypeToTheLogsOwnNumbering(t *testing.T) {
	cases := []struct {
		name        string
		engineType  proto.ResourceType
		wantLogType int64
	}{
		{"Mana", proto.ResourceType_ResourceTypeMana, 0},
		{"Rage", proto.ResourceType_ResourceTypeRage, 1},
		{"Energy", proto.ResourceType_ResourceTypeEnergy, 3},
		{"ComboPoints", proto.ResourceType_ResourceTypeComboPoints, 4},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			u := oneAction()
			u.Resources = []*proto.ResourceMetrics{{
				Id:         &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 2687}},
				Type:       c.engineType,
				Events:     100,
				Gain:       500,
				ActualGain: 500,
			}}
			got, err := Summarize(resultWith(u, 100), req())
			if err != nil {
				t.Fatal(err)
			}
			if len(got.Resources) != 1 {
				t.Fatalf("Resources has %d entries, want 1", len(got.Resources))
			}
			if pt := got.Resources[0].PowerType; pt != c.wantLogType {
				t.Errorf("engine %s (%d) -> PowerType %d, want the log's own %d",
					c.name, c.engineType, pt, c.wantLogType)
			}
		})
	}
}

// The translation table directly, including Focus - unreachable through a
// real class in this ruleset (no spec's PowerBarOptions selects it,
// sim/hunter/hunter.go included), but engineResourceTypeToLogPowerType
// carries it anyway for the day that changes, and this pins it does not
// silently regress to "unrecognised, passed through unchanged" (5, which
// the log's own POWER_NAMES would read as "Runic power").
func TestLogPowerType(t *testing.T) {
	cases := []struct {
		engine proto.ResourceType
		want   int64
	}{
		{proto.ResourceType_ResourceTypeMana, 0},
		{proto.ResourceType_ResourceTypeRage, 1},
		{proto.ResourceType_ResourceTypeFocus, 2},
		{proto.ResourceType_ResourceTypeEnergy, 3},
		{proto.ResourceType_ResourceTypeComboPoints, 4},
		// Health has no entry (see engineResourceTypeToLogPowerType's own
		// comment): an unrecognised type passes through unchanged.
		{proto.ResourceType_ResourceTypeHealth, int64(proto.ResourceType_ResourceTypeHealth)},
	}
	for _, c := range cases {
		if got := logPowerType(c.engine); got != c.want {
			t.Errorf("logPowerType(%v) = %d, want %d", c.engine, got, c.want)
		}
	}
}

// A rage-gain event travels through the same ActionMetrics channel a real
// ability does, but it is the engine's own bookkeeping, not something the
// player cast - the persona review that found this defect saw 98
// "Rage gain" rows in the Casts tab. It must not appear as a cast, an
// auto attack (a real cast) must survive alongside it, and the rage
// figures the filter is not responsible for must still reach Resources.
func TestCastsExcludeResourcePseudoActions(t *testing.T) {
	u := &proto.UnitMetrics{
		Name: "Fury",
		Actions: []*proto.ActionMetrics{
			{
				Id:      &proto.ActionID{RawId: &proto.ActionID_OtherId{OtherId: proto.OtherAction_OtherActionRageGain}},
				Targets: []*proto.TargetedActionMetrics{{UnitIndex: 1, Casts: 144}},
			},
			{
				// Auto attack, tag 1: a real cast, and must survive the filter.
				Id: &proto.ActionID{
					RawId: &proto.ActionID_OtherId{OtherId: proto.OtherAction_OtherActionAttack},
					Tag:   1,
				},
				Targets: []*proto.TargetedActionMetrics{{UnitIndex: 1, Casts: 200, Hits: 200, Damage: 10000}},
			},
		},
		Resources: []*proto.ResourceMetrics{{
			Id:         &proto.ActionID{RawId: &proto.ActionID_OtherId{OtherId: proto.OtherAction_OtherActionRageGain}},
			Type:       proto.ResourceType_ResourceTypeRage,
			Events:     144,
			Gain:       1440,
			ActualGain: 1440,
		}},
	}
	got, err := Summarize(resultWith(u, 100), req())
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range got.Casts {
		if c.SpellName == "other:rage_gain" {
			t.Errorf("Casts has a row named %q; rage gain is the engine's own bookkeeping, not a cast", c.SpellName)
		}
	}
	var sawAttack bool
	for _, c := range got.Casts {
		if c.SpellName == "other:attack/1" {
			sawAttack = true
			if c.Succeeded != 2 { // 200 casts / 100 iterations
				t.Errorf("other:attack/1 Succeeded = %d, want 2", c.Succeeded)
			}
		}
	}
	if !sawAttack {
		t.Error("Casts is missing other:attack/1; a real auto-attack cast must survive the pseudo-action filter")
	}
	if len(got.Resources) != 1 {
		t.Fatalf("Resources has %d entries, want 1; the filter must not touch Resources", len(got.Resources))
	}
	if r := got.Resources[0]; r.Gained != 14 { // 1440/100 rounded
		t.Errorf("resource Gained = %d, want 14; the rage-gain event still belongs in Resources", r.Gained)
	}
}

func TestSummarizeRoster(t *testing.T) {
	got, err := Summarize(resultWith(oneAction(), 100), req())
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Roster) != 1 {
		t.Fatalf("Roster has %d rows, want 1", len(got.Roster))
	}
	r := got.Roster[0]
	if r.Name != "Fury" || r.Class != "warrior" || r.Spec != "fury" {
		t.Errorf("roster row = %+v, want name Fury, class warrior, spec fury", r)
	}
	if r.Role != "dps" {
		t.Errorf("roster Role = %q, want %q", r.Role, "dps")
	}
	if math.Abs(r.DPS-1791.1) > 0.001 {
		t.Errorf("roster DPS = %v, want 1791.1", r.DPS)
	}
}

// A pet is a second actor in the damage table, exactly as it is in a real
// fight's summary, and its cast rows hang off its owner.
func TestSummarizePets(t *testing.T) {
	u := oneAction()
	u.Pets = []*proto.UnitMetrics{{
		Name: "Fury - Pet",
		Dps:  &proto.DistributionMetrics{Avg: 100},
		Actions: []*proto.ActionMetrics{{
			Id:      &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 3110}},
			IsMelee: true,
			Targets: []*proto.TargetedActionMetrics{{UnitIndex: 1, Casts: 100, Hits: 100, Damage: 20000}},
		}},
	}}
	got, err := Summarize(resultWith(u, 100), req())
	if err != nil {
		t.Fatal(err)
	}
	if len(got.DamageDone) != 2 {
		t.Fatalf("DamageDone has %d actors, want 2 (player and pet)", len(got.DamageDone))
	}
	if got.DamageDone[1].Name != "Fury - Pet" {
		t.Errorf("second actor = %q, want the pet", got.DamageDone[1].Name)
	}
	if got.DamageDone[1].Total != 200 {
		t.Errorf("pet Total = %d, want 200", got.DamageDone[1].Total)
	}
	var petCast *summary.CastRow
	for i := range got.Casts {
		if got.Casts[i].SpellID == 3110 {
			petCast = &got.Casts[i]
		}
	}
	if petCast == nil {
		t.Fatal("the pet's cast row is missing")
	}
	if petCast.OwnerGUID != playerGUID {
		t.Errorf("pet cast OwnerGUID = %q, want the player's %q", petCast.OwnerGUID, playerGUID)
	}
	if petCast.GUID == playerGUID {
		t.Error("the pet's cast row took the player's GUID; the row stays the pet's")
	}
	if petCast.Name != "Fury - Pet" {
		t.Errorf("pet cast Name = %q, want the pet's own name", petCast.Name)
	}
}

// The four fields logs engine 0.5.3 added are all things a sim cannot
// know. Every one is still present and empty, because the report
// components read them and a nil slice marshals as null.
func TestSummarizeSetsTheFieldsASimCannotKnow(t *testing.T) {
	got, err := Summarize(resultWith(oneAction(), 100), req())
	if err != nil {
		t.Fatal(err)
	}
	empties := map[string]int{
		"damage_taken":     len(got.DamageTaken),
		"healing":          len(got.Healing),
		"healing_taken":    len(got.HealingTaken),
		"deaths":           len(got.Deaths),
		"interrupts":       len(got.Interrupts),
		"dispels":          len(got.Dispels),
		"threat":           len(got.Threat),
		"threat_by_target": len(got.ThreatByTarget),
		"taunts":           len(got.Taunts),
		"combatants":       len(got.Combatants),
		"phases":           len(got.Phases),
		"mechanics.rows":   len(got.Mechanics.Rows),
	}
	for name, n := range empties {
		if n != 0 {
			t.Errorf("%s has %d entries, want 0", name, n)
		}
	}
	if got.Mechanics.TableFound {
		t.Error("Mechanics.TableFound is true; a sim has no curated mechanics table")
	}

	// Empty is not the same as null: the web renders a list.
	b, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}
	for _, k := range []string{
		"damage_done", "damage_taken", "healing", "healing_taken", "deaths",
		"auras", "casts", "interrupts", "dispels", "resources", "threat",
		"threat_by_target", "taunts", "combatants", "roster", "phases",
	} {
		raw, ok := m[k]
		if !ok {
			t.Errorf("the summary has no %q key; logs engine 0.5.3 expects it", k)
			continue
		}
		if string(raw) == "null" {
			t.Errorf("%q marshalled as null; it must be [] so the report components render an empty table", k)
		}
	}
}

func TestSummarizeRejectsFailures(t *testing.T) {
	bad := resultWith(oneAction(), 100)
	bad.Error = &proto.ErrorOutcome{Message: "boom"}
	if _, err := Summarize(bad, req()); err == nil {
		t.Fatal("a result carrying an ErrorOutcome was summarized without error")
	}

	empty := &proto.RaidSimResult{RaidMetrics: &proto.RaidMetrics{}, IterationsDone: 100}
	if _, err := Summarize(empty, req()); err == nil {
		t.Fatal("a result with no player was summarized without error")
	}

	zero := resultWith(oneAction(), 0)
	if _, err := Summarize(zero, req()); err == nil {
		t.Fatal("a result with zero iterations was summarized without error")
	}
}

// The api lane calls these two directly, and a nil result is what a
// failed run hands back.
func TestExportedHelpersSurviveANilResult(t *testing.T) {
	if _, err := PlayerMetrics(nil); !errors.Is(err, ErrNoPlayer) {
		t.Errorf("PlayerMetrics(nil) = %v, want ErrNoPlayer", err)
	}
	if got := DPS(nil); got != (api.Estimate{}) {
		t.Errorf("DPS(nil) = %+v, want the zero estimate", got)
	}
	if _, err := Summarize(nil, req()); !errors.Is(err, ErrSimFailed) {
		t.Errorf("Summarize(nil) = %v, want ErrSimFailed", err)
	}
}

func TestDPS(t *testing.T) {
	got := DPS(resultWith(oneAction(), 100))
	if got.Mean != 1791.1 || got.StdDev != 120.5 || got.Min != 1400 || got.Max != 2100 {
		t.Errorf("DPS() = %+v", got)
	}
	// standard error of the mean = stdev / sqrt(n)
	want := 120.5 / math.Sqrt(100)
	if math.Abs(got.Error-want) > 1e-9 {
		t.Errorf("DPS().Error = %v, want %v", got.Error, want)
	}
}

// The summary's Hits, Crits and Ticks are the LOGS ENGINE's counters,
// not sim/core's, and the two conventions are different in a way no
// golden can catch on its own: a golden pins whatever the adapter does
// today, which is how the disagreement survived a review.
//
// logs/engine/summary/damage.go's fold is the definition:
//
//	if periodic { ab.Ticks++ } else { ab.Hits++ }
//	if e.Critical.V { ab.Crits++ }
//
// so Hits counts EVERY non-periodic landing whatever its outcome, Ticks
// every periodic one, and Crits is a second count over both. The report
// computes crit % as crits/hits against exactly that.
//
// sim/core's counters are disjoint: a crit never increments Hits, and
// Glances, Crushes, Blocks and BlockedCrits are counted separately
// again (spell_outcome.go). Copying them straight across printed a
// glancing-heavy warrior as barely hitting, a dot as never critting,
// and a crit rate over the wrong denominator - a sim and a real fight
// of identical shape disagreeing inside one component.
func TestHitsAndCritsUseTheLogsEngineDefinitions(t *testing.T) {
	// Every landed outcome the engine counts, one apiece per iteration
	// so nothing is lost to rounding, with distinct values so a swapped
	// term shows up as a wrong number rather than a coincidence.
	const iters = 1
	u := &proto.UnitMetrics{
		Name: "Fury",
		Dps:  &proto.DistributionMetrics{Avg: 1, Stdev: 0, Max: 1, Min: 1},
		Actions: []*proto.ActionMetrics{{
			Id:      &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 23894}},
			IsMelee: true,
			Targets: []*proto.TargetedActionMetrics{{
				UnitIndex:    1,
				Casts:        1000,
				Hits:         100,
				Crits:        40,
				Glances:      20,
				Crushes:      10,
				Blocks:       5,
				BlockedCrits: 2,
				Ticks:        7,
				CritTicks:    3,
				Misses:       11,
				Dodges:       6,
				Parries:      4,
				Damage:       1000,
			}},
		}},
	}
	got, err := Summarize(resultWith(u, iters), req())
	if err != nil {
		t.Fatal(err)
	}
	ab := got.DamageDone[0].Abilities[0]

	// Every non-periodic landing: 100 + 40 + 20 + 10 + 5 + 2.
	if want := int64(177); ab.Hits != want {
		t.Errorf("Hits = %d, want %d: a crit, a glance, a crush and a block all LANDED, and the logs engine counts each of them a hit", ab.Hits, want)
	}
	// Every critical landing, periodic included: 40 + 2 + 3.
	if want := int64(45); ab.Crits != want {
		t.Errorf("Crits = %d, want %d: a blocked crit and a critical tick are crits, and the report's crit %% is crits/hits", ab.Crits, want)
	}
	// Every periodic landing: 7 + 3.
	if want := int64(10); ab.Ticks != want {
		t.Errorf("Ticks = %d, want %d", ab.Ticks, want)
	}
	// Crits are a subset of the landings, never more of them, or the
	// report prints a crit rate over 100%.
	if ab.Crits > ab.Hits+ab.Ticks {
		t.Errorf("crits %d exceed landings %d", ab.Crits, ab.Hits+ab.Ticks)
	}
	// The outcome keys are a cross-lane vocabulary and every one of
	// them is exercised here with a count that survives the division,
	// because a map read on a nil map returns zero and an assertion
	// that only ever compares zero to zero cannot fail.
	for key, want := range map[string]int64{
		"MISS": 11, "DODGE": 6, "PARRY": 4,
		"BLOCK":    7, // Blocks + BlockedCrits
		"GLANCING": 20, "CRUSHING": 10,
	} {
		if ab.Misses[key] != want {
			t.Errorf("Misses[%q] = %d, want %d", key, ab.Misses[key], want)
		}
	}
}

// An abort is not a corrupt result. The engine reports one with an
// ErrorOutcome whose Type is ErrorOutcomeAborted and whose Message is
// EMPTY, so every guard written as `Error != nil && Message != ""` let
// it through and the user pressing Stop was told the result had zero
// iterations.
func TestAnAbortIsReportedAsAnAbort(t *testing.T) {
	aborted := &proto.RaidSimResult{Error: &proto.ErrorOutcome{Type: proto.ErrorOutcomeType_ErrorOutcomeAborted}}
	if err := ResultError(aborted); !errors.Is(err, ErrAborted) {
		t.Errorf("ResultError(aborted) = %v, want ErrAborted", err)
	}
	if _, err := Summarize(aborted, req()); !errors.Is(err, ErrAborted) {
		t.Errorf("Summarize(aborted) = %v, want ErrAborted", err)
	}
	if errors.Is(ResultError(aborted), ErrSimFailed) {
		t.Error("an abort must not read as a failure: the user asked for it")
	}

	failed := &proto.RaidSimResult{Error: &proto.ErrorOutcome{Message: "boom"}}
	if err := ResultError(failed); !errors.Is(err, ErrSimFailed) {
		t.Errorf("ResultError(failed) = %v, want ErrSimFailed", err)
	}
	// An error outcome with neither a type nor a message is still an
	// error, not a success.
	if err := ResultError(&proto.RaidSimResult{Error: &proto.ErrorOutcome{}}); !errors.Is(err, ErrSimFailed) {
		t.Errorf("ResultError(empty outcome) = %v, want ErrSimFailed", err)
	}
	if err := ResultError(resultWith(oneAction(), 100)); err != nil {
		t.Errorf("ResultError(a good result) = %v, want nil", err)
	}
}

// The spec slug splits where the canonical list says it does, not at
// the first hyphen: "hunter-beast-mastery" is a hunter, spec
// "beast-mastery".
func TestSpecSlugSplitsOnTheCanonicalList(t *testing.T) {
	for slug, want := range map[string][2]string{
		"warrior-fury":         {"warrior", "fury"},
		"hunter-beast-mastery": {"hunter", "beast-mastery"},
		"nonsense":             {"nonsense", ""},
		"not-a-spec":           {"not", "a-spec"},
	} {
		class, spec := splitSpecSlug(slug)
		if class != want[0] || spec != want[1] {
			t.Errorf("splitSpecSlug(%q) = (%q, %q), want (%q, %q)", slug, class, spec, want[0], want[1])
		}
	}
}

// A nil Go slice marshals as null, so a summary built as a zero value
// hands the web `"damage_done": null` where a finished run hands it
// `[]`. That is a null check per key on the path least likely to be
// exercised - the user pressed Stop - so the abort path carries
// EmptySummary() and this walks the marshalled JSON to prove it.
//
// The walk is by REFLECTION over summary.Summary rather than against a
// list written down here: a list field added to that struct later must
// fail this test rather than reach the web as a null.
func TestEverySummaryListIsEmptyNotNull(t *testing.T) {
	res := api.SimResult{
		EngineVersion: enginever.Version,
		Aborted:       true,
		IterationsRun: 7,
		Summary:       EmptySummary(),
	}
	b, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	var envelope struct {
		Summary json.RawMessage `json:"summary"`
	}
	if err := json.Unmarshal(b, &envelope); err != nil {
		t.Fatal(err)
	}
	decoded := map[string]json.RawMessage{}
	if err := json.Unmarshal(envelope.Summary, &decoded); err != nil {
		t.Fatal(err)
	}

	typ := reflect.TypeOf(summary.Summary{})
	lists := 0
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Type.Kind() != reflect.Slice {
			continue
		}
		lists++
		key := strings.Split(field.Tag.Get("json"), ",")[0]
		raw, ok := decoded[key]
		if !ok {
			t.Errorf("an aborted summary has no %q key (field %s)", key, field.Name)
			continue
		}
		if string(raw) != "[]" {
			t.Errorf("an aborted summary's %q is %s, want []", key, raw)
		}
	}
	if lists == 0 {
		t.Fatal("the reflection walk found no list fields; summary.Summary changed shape")
	}

	// The one list that is nested rather than top level.
	var mechanics struct {
		Rows json.RawMessage `json:"rows"`
	}
	if err := json.Unmarshal(decoded["mechanics"], &mechanics); err != nil {
		t.Fatal(err)
	}
	if string(mechanics.Rows) != "[]" {
		t.Errorf("an aborted summary's mechanics.rows is %s, want []", mechanics.Rows)
	}

	// And the summary holds no null at all, at any depth, which is the
	// claim the web lane actually relies on. (The echoed request may
	// hold nulls - those are the caller's own optional lists, and this
	// synthetic request omits them.)
	if bytes.Contains(envelope.Summary, []byte("null")) {
		t.Errorf("an aborted summary marshals a null: %s", envelope.Summary)
	}
}

// The empty shape and the finished one must carry the same keys, or
// "the same shape as a completed result" is only true of the lists.
func TestTheEmptyAndFinishedSummariesHaveTheSameKeys(t *testing.T) {
	full, err := Fixture("warrior-fury")
	if err != nil {
		t.Fatal(err)
	}
	finished, err := Summarize(full, req())
	if err != nil {
		t.Fatal(err)
	}
	keys := func(s summary.Summary) []string {
		b, err := json.Marshal(s)
		if err != nil {
			t.Fatal(err)
		}
		var m map[string]json.RawMessage
		if err := json.Unmarshal(b, &m); err != nil {
			t.Fatal(err)
		}
		out := make([]string, 0, len(m))
		for k := range m {
			out = append(out, k)
		}
		sort.Strings(out)
		return out
	}
	if got, want := keys(EmptySummary()), keys(finished); !slices.Equal(got, want) {
		t.Errorf("the empty summary's keys are %v, the finished one's %v", got, want)
	}
}

// The engine answers with a UnitStats array indexed by proto.Stat; the
// envelope answers with named rows in the order the request asked for
// them, normalised so the reference stat is exactly 1.
func TestWeightsMapping(t *testing.T) {
	req := api.SimRequest{Weights: &api.WeightsSpec{
		Stats:     []string{"agility", "attack_power", "crit"},
		Reference: "attack_power",
	}}
	stats := make([]float64, len(proto.Stat_name))
	stdev := make([]float64, len(proto.Stat_name))
	stats[proto.Stat_StatAgility] = 1.1
	stats[proto.Stat_StatAttackPower] = 1.0
	stats[proto.Stat_StatCrit] = 12.0
	stdev[proto.Stat_StatAgility] = 0.02
	stdev[proto.Stat_StatAttackPower] = 0.01
	stdev[proto.Stat_StatCrit] = 0.30

	res := &proto.StatWeightsResult{Dps: &proto.StatWeightValues{
		Weights:      &proto.UnitStats{Stats: stats},
		WeightsStdev: &proto.UnitStats{Stats: stdev},
	}}
	got, err := Weights(res, req)
	if err != nil {
		t.Fatal(err)
	}
	want := []api.StatWeight{
		{Stat: "agility", Weight: 1.1, Error: 0.02, Insignificant: false},
		{Stat: "attack_power", Weight: 1.0, Error: 0.01, Insignificant: false},
		{Stat: "crit", Weight: 12.0, Error: 0.30, Insignificant: false},
	}
	if len(got) != len(want) {
		t.Fatalf("got %d weights, want %d: %+v", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("weight %d = %+v, want %+v", i, got[i], want[i])
		}
	}
}

// TestWeightsConvertsPopulationStdevToStandardError pins the
// denominator sim/api.WeightsIterationsFactor's doc documents:
// sim/core/statweight.go's WeightsStdev is a population standard
// deviation, and Weights divides it by sqrt(req.Iterations *
// api.WeightsIterationsFactor) - the multiplied count
// sim/request.BuildWeights actually runs the engine's sweep at - the
// same way adapter.DPS divides by sqrt(IterationsDone) for the
// headline number.
func TestWeightsConvertsPopulationStdevToStandardError(t *testing.T) {
	req := api.SimRequest{
		Iterations: 3000,
		Weights: &api.WeightsSpec{
			Stats:     []string{"attack_power", "crit"},
			Reference: "attack_power",
		},
	}
	stats := make([]float64, len(proto.Stat_name))
	stdev := make([]float64, len(proto.Stat_name))
	stats[proto.Stat_StatAttackPower] = 1.0
	stats[proto.Stat_StatCrit] = 5.0
	stdev[proto.Stat_StatAttackPower] = 13.0
	stdev[proto.Stat_StatCrit] = 6.0

	res := &proto.StatWeightsResult{Dps: &proto.StatWeightValues{
		Weights:      &proto.UnitStats{Stats: stats},
		WeightsStdev: &proto.UnitStats{Stats: stdev},
	}}
	got, err := Weights(res, req)
	if err != nil {
		t.Fatal(err)
	}
	n := math.Sqrt(float64(req.Iterations * api.WeightsIterationsFactor))
	want := []float64{13.0 / n, 6.0 / n}
	for i, w := range want {
		if diff := got[i].Error - w; diff > 1e-9 || diff < -1e-9 {
			t.Errorf("weight %d error = %v, want %v (raw stdev / sqrt(%d*%d))",
				i, got[i].Error, w, req.Iterations, api.WeightsIterationsFactor)
		}
	}
}

// TestInsignificant is the table task 5(b2) pins: error compared to
// the absolute value of the weight, error >= weight (not just >)
// flags a row the boundary case included, and the reference stat's
// own exactly-1 weight gets whatever the rule gives it like any
// other row.
func TestInsignificant(t *testing.T) {
	cases := []struct {
		name   string
		weight float64
		stdev  float64
		want   bool
	}{
		{"error well under weight is significant", 6.9, 4.0, false},
		{"error over weight is insignificant", 1.0, 13.0, true},
		{"error exactly equal to a nonzero weight is insignificant", 5.0, 5.0, true},
		{"hard-capped 0 +/- 0 is insignificant", 0.0, 0.0, true},
		{"a real negative weight bigger than its error is significant", -6.0, 2.0, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := api.SimRequest{Weights: &api.WeightsSpec{
				Stats:     []string{"attack_power", "crit"},
				Reference: "attack_power",
			}}
			stats := make([]float64, len(proto.Stat_name))
			stdev := make([]float64, len(proto.Stat_name))
			stats[proto.Stat_StatAttackPower] = 1.0
			stdev[proto.Stat_StatAttackPower] = 0.0
			stats[proto.Stat_StatCrit] = c.weight
			stdev[proto.Stat_StatCrit] = c.stdev

			res := &proto.StatWeightsResult{Dps: &proto.StatWeightValues{
				Weights:      &proto.UnitStats{Stats: stats},
				WeightsStdev: &proto.UnitStats{Stats: stdev},
			}}
			got, err := Weights(res, req)
			if err != nil {
				t.Fatal(err)
			}
			// got[0] is attack_power, the reference: weight 1/1=1,
			// error 0/1=0 - not insignificant since 0 < 1.
			if got[0].Insignificant {
				t.Errorf("reference stat came back insignificant: %+v", got[0])
			}
			if got[1].Insignificant != c.want {
				t.Errorf("crit.Insignificant = %v, want %v (weight %v, error %v)",
					got[1].Insignificant, c.want, got[1].Weight, got[1].Error)
			}
		})
	}
}

// TestInsignificantOnTheReferenceStatItself pins the brief's other
// named case: the reference stat's own weight is always exactly 1
// (it is normalised against itself), and its Insignificant flag gets
// whatever error>=1 gives it like any other row, not a hardcoded
// false.
func TestInsignificantOnTheReferenceStatItself(t *testing.T) {
	for _, c := range []struct {
		name              string
		refStdev          float64
		refRaw            float64
		wantInsignificant bool
	}{
		{"reference error well under its own weight of 1", 2.0, 20.0, false},
		{"reference error over its own weight of 1", 15.0, 10.0, true},
	} {
		t.Run(c.name, func(t *testing.T) {
			req := api.SimRequest{Weights: &api.WeightsSpec{
				Stats:     []string{"attack_power"},
				Reference: "attack_power",
			}}
			stats := make([]float64, len(proto.Stat_name))
			stdev := make([]float64, len(proto.Stat_name))
			stats[proto.Stat_StatAttackPower] = c.refRaw
			stdev[proto.Stat_StatAttackPower] = c.refStdev

			res := &proto.StatWeightsResult{Dps: &proto.StatWeightValues{
				Weights:      &proto.UnitStats{Stats: stats},
				WeightsStdev: &proto.UnitStats{Stats: stdev},
			}}
			got, err := Weights(res, req)
			if err != nil {
				t.Fatal(err)
			}
			if got[0].Weight != 1.0 {
				t.Fatalf("reference weight = %v, want exactly 1", got[0].Weight)
			}
			if got[0].Insignificant != c.wantInsignificant {
				t.Errorf("reference Insignificant = %v, want %v (error %v)",
					got[0].Insignificant, c.wantInsignificant, got[0].Error)
			}
		})
	}
}

// A reference weight of zero cannot be normalised against, and
// reporting infinities would render as a table of blanks.
func TestWeightsRefusesAZeroReference(t *testing.T) {
	req := api.SimRequest{Weights: &api.WeightsSpec{Stats: []string{"crit"}, Reference: "crit"}}
	stats := make([]float64, len(proto.Stat_name))
	res := &proto.StatWeightsResult{Dps: &proto.StatWeightValues{
		Weights:      &proto.UnitStats{Stats: stats},
		WeightsStdev: &proto.UnitStats{Stats: stats},
	}}
	if _, err := Weights(res, req); err == nil {
		t.Error("a zero reference weight was normalised")
	}
}

func TestWeightsRefusals(t *testing.T) {
	req := api.SimRequest{Weights: &api.WeightsSpec{Stats: []string{"crit"}, Reference: "crit"}}
	if _, err := Weights(nil, req); !errors.Is(err, ErrNoWeights) {
		t.Error("a nil result was accepted")
	}
	if _, err := Weights(&proto.StatWeightsResult{}, req); !errors.Is(err, ErrNoWeights) {
		t.Error("a result with no dps block was accepted")
	}
	failed := &proto.StatWeightsResult{Error: &proto.ErrorOutcome{Message: "the engine died"}}
	if _, err := Weights(failed, req); err == nil {
		t.Error("an engine failure was reported as weights")
	}
	if _, err := Weights(&proto.StatWeightsResult{}, api.SimRequest{}); err == nil {
		t.Error("a request with no weights block was accepted")
	}
}

// One iteration's casts, in order, with the pre-pull negative and each
// row carrying the summary's ACTION KEY rather than a display name -
// the page resolves the name with resolveActionName exactly as it
// does for a cast row (contract A12).
func TestSampleMapsTheEnginesCastLog(t *testing.T) {
	res := &proto.RaidSimResult{SampleIteration: &proto.SampleIteration{
		Dps:             1038.66,
		DurationSeconds: 60,
		Casts: []*proto.SampleCast{
			{AtMs: -1500, ActionId: &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 1719}}},
			{AtMs: 0, ActionId: &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 23881}},
				Target: "Target Dummy", Resources: map[string]int32{"rage": 26}},
			{AtMs: 1502, ActionId: &proto.ActionID{RawId: &proto.ActionID_ItemId{ItemId: 13442}}},
			{AtMs: 2000, ActionId: &proto.ActionID{RawId: &proto.ActionID_OtherId{OtherId: proto.OtherAction_OtherActionAttack}}},
		},
	}}

	got := Sample(res)
	if len(got) != 4 {
		t.Fatalf("got %d casts, want 4", len(got))
	}
	if got[0].AtMS != -1500 {
		t.Errorf("the pre-pull cast is at %d ms, want -1500", got[0].AtMS)
	}
	if got[0].Action != "spell:1719" {
		t.Errorf("an untagged spell is %q, want its action key", got[0].Action)
	}
	if got[1].Target != "Target Dummy" || got[1].Resources["rage"] != 26 {
		t.Errorf("cast 1 = %+v", got[1])
	}
	if got[2].AtMS != 1502 {
		t.Errorf("cast 2 is at %d ms, want 1502", got[2].AtMS)
	}
	// An item and an "other" action take the same key form the cast
	// table's rows take, so the page resolves one name one way.
	_, wantItem := ActionName(&proto.ActionID{RawId: &proto.ActionID_ItemId{ItemId: 13442}})
	if got[2].Action != wantItem {
		t.Errorf("the item cast is %q, want %q", got[2].Action, wantItem)
	}
	if got[3].Action != "other:attack" {
		t.Errorf("the white swing is %q", got[3].Action)
	}
	// No display name anywhere: the page resolves them, and a name
	// baked in here would be a second vocabulary.
	for i, c := range got {
		if c.Action == "" {
			t.Errorf("cast %d has no action key", i)
		}
	}
}

// The sample table is a cast log too, so it must exclude the same
// resource pseudo-actions summary.Casts does: a rage-gain event in the
// sample timeline is exactly as much not a cast as it is in the Casts
// tab.
func TestSampleExcludesResourcePseudoActions(t *testing.T) {
	res := &proto.RaidSimResult{SampleIteration: &proto.SampleIteration{
		Casts: []*proto.SampleCast{
			{AtMs: 0, ActionId: &proto.ActionID{RawId: &proto.ActionID_OtherId{OtherId: proto.OtherAction_OtherActionRageGain}}},
			{AtMs: 100, ActionId: &proto.ActionID{
				RawId: &proto.ActionID_OtherId{OtherId: proto.OtherAction_OtherActionAttack}, Tag: 1,
			}},
		},
	}}
	got := Sample(res)
	if len(got) != 1 {
		t.Fatalf("got %d sample casts, want 1 (the rage-gain event must be filtered)", len(got))
	}
	if got[0].Action != "other:attack/1" {
		t.Errorf("surviving sample cast = %q, want %q", got[0].Action, "other:attack/1")
	}
}

// A result with no sample is not an error: an aborted run, a bulk
// stage and an older engine all produce one, and the page renders the
// card only when there are rows.
func TestSampleOfNothingIsNothing(t *testing.T) {
	if got := Sample(nil); got != nil {
		t.Errorf("Sample(nil) = %+v", got)
	}
	if got := Sample(&proto.RaidSimResult{}); got != nil {
		t.Errorf("a result with no sample produced %+v", got)
	}
	if got := Sample(&proto.RaidSimResult{SampleIteration: &proto.SampleIteration{}}); got != nil {
		t.Errorf("an empty sample produced %+v", got)
	}
}

// The checked-in fixtures are real engine output, so the mapping is
// proved against one rather than only against hand-built messages.
func TestSampleOfTheWarriorFixture(t *testing.T) {
	res, err := Fixture("warrior-fury")
	if err != nil {
		t.Fatal(err)
	}
	got := Sample(res)
	if len(got) == 0 {
		t.Skip("the checked-in fixture predates sample_iteration; refresh it with `forever-sim -out-proto`")
	}
	var last int64 = -1 << 62
	for i, c := range got {
		if c.AtMS < last {
			t.Errorf("cast %d is at %d ms, after %d; the sample is in cast order", i, c.AtMS, last)
		}
		last = c.AtMS
		if c.Action == "" {
			t.Errorf("cast %d has no action key", i)
		}
	}
}

// combatLogSchool must map the engine's core school bitmask onto the
// combat log's, bit by bit: the persona review that found this bug saw
// every warrior ability, and only warrior abilities, render as school
// "Holy" (engine 2, Physical, misread as log 2, Holy) because the two
// masks assign the same schools to different bit positions.
func TestCombatLogSchool(t *testing.T) {
	tests := []struct {
		name   string
		engine int32
		want   int64
	}{
		{"none maps to none, not to Physical's bit", 0, 0},
		{"warrior physical ability: the bug this task fixes", 2, 1},
		// Frost is the one school the two masks happen to place on the
		// same bit. It is still listed explicitly in
		// engineSchoolToLogBit and tested here so nobody later
		// "simplifies" the table by dropping the identity entry.
		{"frost mage frost ability: coincidentally identical on both masks", 16, 16},
		{"holy", 32, 2},
		{"arcane", 4, 64},
		{"fire", 8, 4},
		{"nature", 64, 8},
		{"shadow", 128, 32},
		{"combined mask ORs both translated bits: Frost|Shadow", 16 | 128, 16 | 32},
		{"a bit the table does not know is dropped, not mapped to a wrong school", 1 << 20, 0},
		{"an unmapped bit combined with a known one drops only the unknown bit", 2 | (1 << 20), 1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := combatLogSchool(tc.engine); got != tc.want {
				t.Errorf("combatLogSchool(%d) = %d, want %d", tc.engine, got, tc.want)
			}
		})
	}
}

// The mapping has to be wired into ability(), not just correct in
// isolation: this pins Summarize's output for a warrior's physical
// ability (the exact case the persona review flagged) and a frost
// mage's frost ability (the coincidentally-identical case, which would
// pass even if ability() still wrote am.SpellSchool straight through).
func TestSummarizeTranslatesAbilitySchool(t *testing.T) {
	const iters = 1
	tests := []struct {
		name       string
		spellID    int32
		engineMask int32
		wantSchool int64
	}{
		{"warrior physical ability", 23894, 2, 1},
		{"frost mage frost ability", 25304, 16, 16},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			u := &proto.UnitMetrics{
				Name: "Sim",
				Dps:  &proto.DistributionMetrics{Avg: 1, Stdev: 0, Max: 1, Min: 1},
				Actions: []*proto.ActionMetrics{{
					Id:          &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: tc.spellID}},
					SpellSchool: tc.engineMask,
					Targets: []*proto.TargetedActionMetrics{{
						UnitIndex: 1,
						Hits:      1,
						Damage:    100,
					}},
				}},
			}
			got, err := Summarize(resultWith(u, iters), req())
			if err != nil {
				t.Fatal(err)
			}
			ab := got.DamageDone[0].Abilities[0]
			if ab.School != tc.wantSchool {
				t.Errorf("School = %d, want %d", ab.School, tc.wantSchool)
			}
		})
	}
}
