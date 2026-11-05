package coach

import "testing"

// fixtureHammerID is Hammer of the Righteous: a Protection paladin's trainer spell at 40 on
// a 6 second cooldown (1.60.1.70790/spellconst/paladin.json#407632).
const fixtureHammerID = int64(407632)

// hammerTankAt is a tank at level whose rule file lists Hammer of the Righteous as a core
// trainer spell (not a talent), learned at 40, who never casts it in a 60 second fight.
func hammerTankAt(t *testing.T, level int64) Input {
	t.Helper()
	in := protInput(t)
	in.Fight = withLevel(withDuration(withoutCasts(in.Fight, furyGUID), 60_000), furyGUID, level)
	in.Row.Level = level
	in.Tables.Rules = SpecRules{
		Spec: "paladin-protection", Build: fixtureVerified,
		Core: []RuleEntry{{Family: "Hammer of the Righteous", SpellIDs: []int64{fixtureHammerID}, Verified: fixtureVerified}},
	}
	in.Tables.Spells = map[int64]SpellConst{
		fixtureHammerID: {Name: "Hammer of the Righteous", Rank: 1, CooldownMS: 6000, GCDMS: 1500, CostKnown: true},
	}
	in.Tables.Trainable = map[string][]TrainableRank{
		"Hammer of the Righteous": {{ID: fixtureHammerID, Rank: 1, Level: 40}},
	}
	return in
}

func TestCoreCasts_TrainerSpellIsJudgedOnlyFromTheLevelItIsLearned(t *testing.T) {
	for _, level := range []int64{30, 35, 39} {
		in := hammerTankAt(t, level)
		if o := coreCasts(in, CountedTime(in)); o.Finding != nil {
			t.Errorf("level %d: a spell learned at 40 was judged: %+v", level, o.Finding)
		}
	}
	in := hammerTankAt(t, 45)
	o := coreCasts(in, CountedTime(in))
	if o.Finding == nil || o.Finding.Code != CodeCoreCasts {
		t.Fatalf("level 45: never casting a spell learned at 40 is a finding, got %+v", o.Finding)
	}
}

func TestCoreCasts_UnknownLevelKeepsJudgingTrainerSpells(t *testing.T) {
	in := hammerTankAt(t, 0)
	if o := coreCasts(in, CountedTime(in)); o.Finding == nil {
		t.Fatal("with no level in the log the trainer spell is still judged, as before")
	}
}
