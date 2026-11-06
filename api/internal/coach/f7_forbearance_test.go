package coach

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// The paladin defensives F7's Forbearance tests read (1.60.1.70790/spellconst/paladin.json).
const (
	divineShieldR2ID     = int64(1020)
	divineProtectionR2ID = int64(5573)
	paladinDefensiveCDMS = int64(300_000)
)

// retDeath is a level 60 Retribution paladin dying to Onyxia's melee at atMS in a fight of
// durationMS, with Divine Shield and Divine Protection as the rule file's defensives (each
// a five minute cooldown) and forbearance the segments the player carried Forbearance.
func retDeath(t *testing.T, durationMS, atMS int64, forbearance ...summary.Segment) Input {
	t.Helper()
	in := f7Level(protDeathShieldWallReady(t, durationMS, atMS), 60)
	in.Fight = withRole(in.Fight, in.Row.GUID, "Paladin", "Retribution", RoleDPS)
	in.Row = rosterRowOf(t, in.Fight, in.Row.GUID)
	in = f7Defensives(in,
		RuleEntry{Family: "Divine Shield", SpellIDs: []int64{642, divineShieldR2ID}},
		RuleEntry{Family: "Divine Protection", SpellIDs: []int64{498, divineProtectionR2ID}})
	in.Tables.Spells = map[int64]SpellConst{
		divineShieldR2ID:     {Name: "Divine Shield", Rank: 2, CooldownMS: paladinDefensiveCDMS},
		divineProtectionR2ID: {Name: "Divine Protection", Rank: 2, CooldownMS: paladinDefensiveCDMS},
	}
	in.Tables.Trainable = map[string][]TrainableRank{
		"Divine Shield":     {{ID: divineShieldR2ID, Rank: 2, Level: 50}},
		"Divine Protection": {{ID: divineProtectionR2ID, Rank: 2, Level: 18}},
	}
	if len(forbearance) > 0 {
		in.Fight = withAura(in.Fight, in.Row.GUID, forbearanceSpellID, "Forbearance", []string{in.Row.GUID}, forbearance...)
	}
	return in
}

// Control: with no Forbearance, Divine Protection five minutes into the pull is ready.
func TestDeath_DivineProtectionWithoutForbearanceIsReady(t *testing.T) {
	in := f7Cast(retDeath(t, 600_000, 330_000), divineShieldR2ID, "Divine Shield", 300_000)
	if f := death(in, CountedTime(in)).Finding; f == nil || f.Title != "Died at 5:30 with Divine Protection ready" {
		t.Fatalf("got %+v", f)
	}
}

// A paladin who bubbled at 5:00 and died at 5:30 carried Forbearance: neither Divine
// Shield nor Divine Protection could be cast.
func TestDeath_ForbearanceBlocksDivineProtection(t *testing.T) {
	in := retDeath(t, 600_000, 330_000, summary.Segment{StartMS: 300_000, EndMS: 330_000})
	in = f7Cast(in, divineShieldR2ID, "Divine Shield", 300_000)
	if f := death(in, CountedTime(in)).Finding; f == nil || f.Title != "Died at 5:30 to Onyxia's melee" {
		t.Fatalf("Forbearance was up at the death: got %+v", f)
	}
}

// Forbearance from another paladin's Blessing of Protection blocks a bubble never cast.
func TestDeath_ForbearanceBlocksANeverCastBubble(t *testing.T) {
	in := retDeath(t, 600_000, 100_000, summary.Segment{StartMS: 80_000, EndMS: 140_000})
	if f := death(in, CountedTime(in)).Finding; f == nil || f.Title != "Died at 1:40 to Onyxia's melee" {
		t.Fatalf("Forbearance was up at the death: got %+v", f)
	}
}

// Forbearance that ran out before the death blocks nothing.
func TestDeath_ForbearanceThatEndedBlocksNothing(t *testing.T) {
	in := retDeath(t, 600_000, 100_000, summary.Segment{StartMS: 10_000, EndMS: 70_000})
	if f := death(in, CountedTime(in)).Finding; f == nil || f.Title != "Died at 1:40, Divine Shield not used" {
		t.Fatalf("got %+v", f)
	}
}
