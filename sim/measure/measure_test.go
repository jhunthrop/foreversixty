package measure

import (
	"math"
	"testing"
	"time"
)

const fixture = "testdata/planted-v22.log"

var fixtureBase = time.Date(2026, 9, 17, 0, 0, 0, 0, time.UTC)

func load(t *testing.T) Input {
	t.Helper()
	events, err := Load(fixture, fixtureBase)
	if err != nil {
		t.Fatal(err)
	}
	if len(events) == 0 {
		t.Fatal("the fixture produced no events")
	}
	return Input{
		Events:     events,
		Actor:      "Testwarrior-Beta",
		SpellPower: 500,
		MinSamples: 10,
	}
}

func near(t *testing.T, label string, got, want, tol float64) {
	t.Helper()
	if math.Abs(got-want) > tol {
		t.Errorf("%s = %v, want %v (tolerance %v)", label, got, want, tol)
	}
}

// (1) Whether periodic damage crits at all is a server rule, not a
// per-spell field, so the only way to know is to look for the flag.
func TestPeriodicDamageCanCrit(t *testing.T) {
	rep, err := Run(load(t))
	if err != nil {
		t.Fatal(err)
	}
	if len(rep.Periodic) != 2 {
		t.Fatalf("Periodic has %d rows, want 2 (Rend and Corruption)", len(rep.Periodic))
	}
	for _, p := range rep.Periodic {
		if !p.CanCrit {
			t.Errorf("%s (%d): CanCrit is false, but the fixture plants critical ticks", p.SpellName, p.SpellID)
		}
		if p.Ticks != 40 {
			t.Errorf("%s: Ticks = %d, want 40", p.SpellName, p.Ticks)
		}
		if p.CritTicks != 10 {
			t.Errorf("%s: CritTicks = %d, want 10", p.SpellName, p.CritTicks)
		}
	}
}

// (2) The multiplier a critical tick uses. Planted at exactly 2.0, and
// reported separately for physical and magic because Forever may differ.
func TestPeriodicCritMultiplier(t *testing.T) {
	rep, err := Run(load(t))
	if err != nil {
		t.Fatal(err)
	}
	var sawPhysical, sawMagic bool
	for _, p := range rep.Periodic {
		near(t, p.SpellName+" mean normal tick", p.MeanNormal, 100, 0.001)
		near(t, p.SpellName+" mean critical tick", p.MeanCrit, 200, 0.001)
		near(t, p.SpellName+" multiplier", p.Multiplier, 2.0, 0.001)
		if !p.Enough {
			t.Errorf("%s: Enough is false with 40 ticks and MinSamples 10", p.SpellName)
		}
		if p.Physical {
			sawPhysical = true
		} else {
			sawMagic = true
		}
	}
	if !sawPhysical {
		t.Error("no physical periodic row; Rend is school 1")
	}
	if !sawMagic {
		t.Error("no magic periodic row; Corruption is school 32")
	}
}

// (3) The attack table, which with a known character sheet is what
// settles how unified Hit interacts with the weapon-skill miss table.
func TestAttackTableRates(t *testing.T) {
	rep, err := Run(load(t))
	if err != nil {
		t.Fatal(err)
	}
	if len(rep.AttackTable) != 1 {
		t.Fatalf("AttackTable has %d rows, want 1 (melee against the dummy)", len(rep.AttackTable))
	}
	row := rep.AttackTable[0]
	if row.Swings != 100 {
		t.Fatalf("Swings = %d, want 100", row.Swings)
	}
	near(t, "miss", row.Miss, 0.10, 1e-9)
	near(t, "dodge", row.Dodge, 0.10, 1e-9)
	near(t, "parry", row.Parry, 0.10, 1e-9)
	near(t, "glance", row.Glance, 0.20, 1e-9)
	near(t, "crit", row.Crit, 0.20, 1e-9)
	if !row.Enough {
		t.Error("Enough is false with 100 swings")
	}
}

// (4) Proc rate and internal cooldown. Neither is in DB2 at all.
func TestProcRateAndInternalCooldown(t *testing.T) {
	rep, err := Run(load(t))
	if err != nil {
		t.Fatal(err)
	}
	if len(rep.Procs) != 1 {
		t.Fatalf("Procs has %d rows, want 1", len(rep.Procs))
	}
	p := rep.Procs[0]
	if p.SpellID != 9345 {
		t.Errorf("SpellID = %d, want 9345", p.SpellID)
	}
	if p.Procs != 10 {
		t.Errorf("Procs = %d, want 10", p.Procs)
	}
	if p.Swings != 100 {
		t.Errorf("Swings = %d, want 100", p.Swings)
	}
	near(t, "per swing", p.PerSwing, 0.10, 1e-9)
	if p.MinGap != 45*time.Second {
		t.Errorf("MinGap = %v, want 45s: that is the internal cooldown", p.MinGap)
	}
}

// (5) Observed coefficients, because EffectBonusCoefficient is routinely
// 0 or wrong for Classic-lineage spells.
func TestObservedCoefficient(t *testing.T) {
	rep, err := Run(load(t))
	if err != nil {
		t.Fatal(err)
	}
	var fb *Coefficient
	for i := range rep.Coefficients {
		if rep.Coefficients[i].SpellID == 25304 {
			fb = &rep.Coefficients[i]
		}
	}
	if fb == nil {
		t.Fatal("no coefficient row for Frostbolt (25304)")
	}
	if fb.Hits != 40 {
		t.Errorf("Hits = %d, want 40", fb.Hits)
	}
	near(t, "mean damage", fb.MeanDamage, 500, 0.001)
	near(t, "spread", fb.StdDev, 50, 0.001)
}

// A figure from too few samples is worse than no figure: it goes into a
// constants file and nobody re-checks it.
func TestInsufficientDataIsNotANumber(t *testing.T) {
	in := load(t)
	in.MinSamples = 1000
	rep, err := Run(in)
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range rep.Periodic {
		if p.Enough {
			t.Errorf("%s: Enough is true with MinSamples 1000 and %d ticks", p.SpellName, p.Ticks)
		}
	}
	for _, r := range rep.AttackTable {
		if r.Enough {
			t.Errorf("attack table: Enough is true with MinSamples 1000 and %d swings", r.Swings)
		}
	}
	out := rep.Table()
	if !contains(out, "insufficient data") {
		t.Errorf("the table prints numbers it should have withheld:\n%s", out)
	}
}

// The report must name what it measured, so a figure can be traced back
// to the log it came from.
func TestReportCarriesItsProvenance(t *testing.T) {
	rep, err := Run(load(t))
	if err != nil {
		t.Fatal(err)
	}
	if rep.Actor != "Testwarrior-Beta" {
		t.Errorf("Actor = %q", rep.Actor)
	}
	if rep.Events == 0 {
		t.Error("Events is zero")
	}
}

func TestRunRejectsAnEmptyInput(t *testing.T) {
	if _, err := Run(Input{}); err == nil {
		t.Fatal("an empty input was accepted")
	}
	if _, err := Load("testdata/nope.log", fixtureBase); err == nil {
		t.Fatal("a missing log was accepted")
	}
}

func contains(h, n string) bool {
	for i := 0; i+len(n) <= len(h); i++ {
		if h[i:i+len(n)] == n {
			return true
		}
	}
	return false
}
