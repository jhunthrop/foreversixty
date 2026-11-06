package coach

import (
	"slices"
	"testing"
)

const (
	f2DemoShoutID   = 11556 // Demoralizing Shout, the warrior's top trained rank
	f2DemoShoutName = "Demoralizing Shout"
	f2DemoRoarID    = 9898 // Demoralizing Roar, a druid's: the same slot on the boss, they do not stack
	f2DruidGUID     = "Player-1136-0CCCCCCC"
)

// protDemoShout is the Protection warrior whose rule file judges Demoralizing Shout on the
// boss, covered by coveredBy, who shouts for shoutMS from the pull while a druid's Roar holds
// the slot from roarFromMS to the end.
func protDemoShout(t *testing.T, shoutMS, roarFromMS int64, coveredBy []int64) Input {
	t.Helper()
	in := f2Steady(protInput(t))
	in.Tables.Rules.Auras = append(slices.Clone(in.Tables.Rules.Auras), AuraRule{
		SpellID: f2DemoShoutID, Name: f2DemoShoutName, Kind: AuraDefensive, CoveredBy: coveredBy, Verified: fixtureVerified,
	})
	if shoutMS > 0 {
		in.Fight = withAura(in.Fight, f2BossGUID, f2DemoShoutID, f2DemoShoutName, []string{furyGUID}, span(0, shoutMS))
	}
	if roarFromMS < f2FightMS {
		in.Fight = withAura(in.Fight, f2BossGUID, f2DemoRoarID, "Demoralizing Roar", []string{f2DruidGUID}, span(roarFromMS, f2FightMS))
	}
	return withPool(in, poolDefensive(f2DemoShoutID, 90, 25))
}

func TestUptime_ARoarHoldingTheSlotCountsForTheShout(t *testing.T) {
	// He shouts for the first fifth, then the druid's Roar holds the slot: the boss is
	// demoralized the whole fight, which is what the guide asks of him.
	in := protDemoShout(t, f2FightMS/5, f2FightMS/5, []int64{f2DemoRoarID})
	if f := uptime(in, CountedTime(in)).Finding; f != nil && f.Title == "Demoralizing Shout up 20% of the fight" {
		t.Fatalf("the Roar's time read as a gap: %+v", f)
	}
}

func TestUptime_ARoarAloneCoversAShoutNeverCast(t *testing.T) {
	in := protDemoShout(t, 0, 0, []int64{f2DemoRoarID})
	if f := uptime(in, CountedTime(in)).Finding; f != nil && f.Title == "Demoralizing Shout up 0% of the fight" {
		t.Fatalf("a Roar all fight read as no Shout: %+v", f)
	}
}

func TestUptime_UncoveredGapStillCounts(t *testing.T) {
	// Without a covering aura in the rule, the same fight is the low uptime it always was.
	in := protDemoShout(t, f2FightMS/5, f2FightMS/5, nil)
	f := uptime(in, CountedTime(in)).Finding
	if f == nil || f.Title != "Demoralizing Shout up 20% of the fight" {
		t.Fatalf("finding %+v", f)
	}
}
