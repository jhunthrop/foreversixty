package coach

// F7's fixtures (Task 12). Every death is hand-built on the v22 golden (a shape fixture,
// fixtures_test.go): none of these is a known-good value.

import (
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// The fixture spell ids F7's tests cast.
const (
	flameBreathID     = 18435
	shieldWallID      = 871
	lastStandID       = 12975
	majorHealingPotID = 17534 // "Healing Potion", the Major Healing Potion's use
	majorHealthstone  = 23477 // "Major Healthstone"
	meleeSpellID      = 0
)

// f7KillingBlowAmount is s with the amount of guid's death at atMS set: withDeath leaves
// the killing blow's amount out.
func f7KillingBlowAmount(s summary.Summary, guid string, atMS, amount int64) summary.Summary {
	s.Deaths = slices.Clone(s.Deaths)
	for i, d := range s.Deaths {
		if d.GUID == guid && d.AtMS == atMS && d.KillingBlow != nil {
			blow := *d.KillingBlow
			blow.Amount = amount
			s.Deaths[i].KillingBlow = &blow
		}
	}
	return s
}

// f7Died is in with its player dying at atMS to spellID for amount, in a fight lasting
// durationMS.
func f7Died(in Input, durationMS, atMS, spellID int64, name string, amount int64) Input {
	guid := in.Row.GUID
	s := withDeath(withDuration(in.Fight, durationMS), guid, atMS, spellID, name)
	in.Fight = f7KillingBlowAmount(s, guid, atMS, amount)
	return in
}

// holyDeath is the golden's Holy priest dying at atMS to spellID, hit for amount.
func holyDeath(t *testing.T, durationMS, atMS, spellID int64, name string, amount int64) Input {
	t.Helper()
	return f7Died(holyInput(t), durationMS, atMS, spellID, name, amount)
}

// protDeathShieldWallReady is the golden's Protection warrior dying to Onyxia's melee at
// atMS, with Shield Wall (baseline, trainable at 28) never cast and Last Stand (talent)
// never cast.
func protDeathShieldWallReady(t *testing.T, durationMS, atMS int64) Input {
	t.Helper()
	return f7Died(protInput(t), durationMS, atMS, meleeSpellID, "Melee", 1_830)
}

// wipeDeath is the golden's Fury warrior dying to Flame Breath at atMS on a wipe lasting
// durationMS.
func wipeDeath(t *testing.T, durationMS, atMS int64) Input {
	t.Helper()
	in := f7Died(furyInput(t), durationMS, atMS, flameBreathID, "Flame Breath", 2_140)
	in.Fight.Kill = false
	return in
}

// f7Cast is in with its player's row for spellID replaced by casts at seqMS.
func f7Cast(in Input, spellID int64, name string, seqMS ...int64) Input {
	in.Fight = withCasts(in.Fight, in.Row.GUID, spellID, name, seqMS...)
	return in
}

// f7Level is in with its player's level set, in the fight and the row.
func f7Level(in Input, level int64) Input {
	in.Fight = withLevel(in.Fight, in.Row.GUID, level)
	in.Row.Level = level
	return in
}

// f7Defensives is in with its rule file's defensives replaced.
func f7Defensives(in Input, entries ...RuleEntry) Input {
	in.Tables.Rules.Defensives = slices.Clone(entries)
	return in
}
