package adapter

import (
	"encoding/json"
	"errors"
	"math"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/sim/api"
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
		EngineVersion: "7779ebb",
		Spec:          "warrior-fury",
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
	if got.EngineVersion != "sim:7779ebb" {
		t.Errorf("EngineVersion = %q, want %q", got.EngineVersion, "sim:7779ebb")
	}
	if got.FightIndex != 1 {
		t.Errorf("FightIndex = %d, want 1", got.FightIndex)
	}
	if got.DurationMS != 180000 {
		t.Errorf("DurationMS = %d, want 180000 (avg_iteration_duration * 1000)", got.DurationMS)
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
	if a.ActiveMS != 180000 {
		t.Errorf("ActiveMS = %d, want the fight duration 180000", a.ActiveMS)
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
	if ab.Hits != 2 { // 160/100 = 1.6, rounds to 2
		t.Errorf("Hits = %d, want 2 (160/100 rounded)", ab.Hits)
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
	if r.PowerType != int64(proto.ResourceType_ResourceTypeRage) {
		t.Errorf("resource PowerType = %d, want %d", r.PowerType, proto.ResourceType_ResourceTypeRage)
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
