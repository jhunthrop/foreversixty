package coach

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	// fixtureJudgementID is Judgement: a 10 second cooldown costing 6% of base mana
	// (1.60.1.70790/spellconst/paladin.json#20271).
	fixtureJudgementID = int64(20271)
	// fixturePaladinBaseMana60 is a level 60 paladin's base mana
	// (1.60.1.70790/levels.json, gametables/basemp.txt).
	fixturePaladinBaseMana60 = int64(1512)
	fixtureManaMax           = int64(5000)
)

// judgementPaladin is a level 60 Retribution paladin whose rule file lists Judgement as a
// core ability, who judges at 0 and 10 s of a 60 s fight and never again, with mana per
// second and the class's base mana at the row's level (0 is unknown).
func judgementPaladin(t *testing.T, mana []int64, baseMana int64) Input {
	t.Helper()
	in := protInput(t)
	in.Fight = withRole(in.Fight, furyGUID, "Paladin", "Retribution", RoleDPS)
	in.Fight = withDuration(withoutCasts(in.Fight, furyGUID), 60_000)
	in.Fight = withCasts(in.Fight, furyGUID, fixtureJudgementID, "Judgement", 0, 10_000)
	in.Fight = withResource(in.Fight, furyGUID, PowerMana, fixtureManaMax, mana...)
	in.Row = rosterRowOf(t, in.Fight, furyGUID)
	in.Tables.Rules = SpecRules{Spec: "paladin-retribution", Build: fixtureVerified,
		Core: []RuleEntry{{Family: "Judgement", SpellIDs: []int64{fixtureJudgementID}, Verified: fixtureVerified}}}
	in.Tables.Spells = map[int64]SpellConst{
		fixtureJudgementID: {Name: "Judgement", CooldownMS: 10_000, CostType: PowerMana, CostPct: 6},
	}
	in.Tables.Trainable = nil
	in.Tables.BaseMana = baseMana
	return in
}

// A paladin out of mana could not judge: no lost cast, and no "had the mana" claim.
func TestCoreCasts_PercentCostOnAnEmptyBarIsNoLostCast(t *testing.T) {
	in := judgementPaladin(t, repeat(0, 60), fixturePaladinBaseMana60)
	if o := coreCasts(in, CountedTime(in)); o.Finding != nil {
		t.Fatalf("Judgement costs 6%% of base mana and the bar was empty: %+v", o.Finding)
	}
}

// Under the cost (6% of 1512 is 90) is still too little to judge.
func TestCoreCasts_PercentCostReadsBaseManaAtTheLevel(t *testing.T) {
	in := judgementPaladin(t, repeat(89, 60), fixturePaladinBaseMana60)
	if o := coreCasts(in, CountedTime(in)); o.Finding != nil {
		t.Fatalf("89 mana is under Judgement's 90: %+v", o.Finding)
	}
}

// With the mana there, the idle Judgements are lost casts and the coach says nothing
// about not being able to check.
func TestCoreCasts_PercentCostWithTheManaIsALostCast(t *testing.T) {
	in := judgementPaladin(t, repeat(2000, 60), fixturePaladinBaseMana60)
	o := coreCasts(in, CountedTime(in))
	if o.Finding == nil || o.Finding.Code != CodeCoreCasts {
		t.Fatalf("idle Judgements with a full bar are lost casts: %+v", o.Finding)
	}
	if strings.Contains(o.Finding.Detail, "Couldn't check") {
		t.Errorf("the mana was checked: %q", o.Finding.Detail)
	}
}

// With no base mana (no level, or a build without levels.json) the cost is unknown, and
// the coach says it couldn't check the mana.
func TestCoreCasts_PercentCostWithoutBaseManaCouldNotCheck(t *testing.T) {
	in := judgementPaladin(t, repeat(0, 60), 0)
	o := coreCasts(in, CountedTime(in))
	if o.Finding == nil || !strings.Contains(o.Finding.Detail, "Couldn't check whether you had the mana.") {
		t.Fatalf("want the couldn't-check line: %+v", o.Finding)
	}
}

func TestParseSpellConst_ReadsThePercentCost(t *testing.T) {
	data, err := ParseSpellConst(buildFile(t, "spellconst", "paladin.json"))
	if err != nil {
		t.Fatal(err)
	}
	if j := data[fixtureJudgementID]; j.CostKnown || j.CostPct != 6 || j.CostType != PowerMana {
		t.Fatalf("Judgement %+v: want 6%% of base mana", j)
	}
}

func TestTablesFor_BaseManaAtTheRowLevel(t *testing.T) {
	ft := &FileTables{BuildsDir: repoBuilds, CoachDir: t.TempDir(), Specs: fakeSpecs{}}
	row := summary.RosterRow{Class: "Paladin", Spec: "Retribution", Role: "dps", Level: 60}
	tb, err := ft.TablesFor(context.Background(), row, furyEncounter, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	if tb.BaseMana != fixturePaladinBaseMana60 {
		t.Fatalf("BaseMana = %d, want %d", tb.BaseMana, fixturePaladinBaseMana60)
	}
	row.Level = 0
	if tb, _ := ft.TablesFor(context.Background(), row, furyEncounter, time.Now()); tb.BaseMana != 0 {
		t.Fatalf("no level: BaseMana = %d, want 0 (unknown)", tb.BaseMana)
	}
}
