package coach

import (
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// The rogue defensives F7's melee-only tests read (1.60.1.70790/spellconst/rogue.json):
// Evasion raises dodge only (aura 49), Vanish does not.
const (
	evasionID       = int64(5277)
	vanishR2ID      = int64(1857)
	rogueDefenseCD  = int64(300_000)
	fireballBoltID  = int64(18392) // Onyxia's Fireball
	schoolFire      = int64(0x4)
	knockAwayID     = int64(19633) // Onyxia's Knock Away, a physical ability
	rogueDeathLevel = 60
)

// rogueDeath is a level 60 rogue dying at 6:40 of a ten-minute pull to spellID (school,
// as the player's damage-taken row records it), with Evasion and Vanish never cast.
func rogueDeath(t *testing.T, spellID int64, name string, school int64) Input {
	t.Helper()
	in := f7Level(f7Died(protInput(t), 600_000, 400_000, spellID, name, 2_400), rogueDeathLevel)
	in.Fight = withRole(in.Fight, in.Row.GUID, "Rogue", "Combat", RoleDPS)
	in.Row = rosterRowOf(t, in.Fight, in.Row.GUID)
	in.Fight = withDamageTaken(in.Fight, in.Row.GUID, spellID, name, school)
	in = f7Defensives(in,
		RuleEntry{Family: "Evasion", SpellIDs: []int64{evasionID}},
		RuleEntry{Family: "Vanish", SpellIDs: []int64{1856, vanishR2ID}})
	in.Tables.Spells = map[int64]SpellConst{
		evasionID:  {Name: "Evasion", CooldownMS: rogueDefenseCD, MeleeOnly: true},
		vanishR2ID: {Name: "Vanish", Rank: 2, CooldownMS: rogueDefenseCD},
	}
	in.Tables.Trainable = map[string][]TrainableRank{
		"Evasion": {{ID: evasionID, Rank: 1, Level: 8}},
		"Vanish":  {{ID: vanishR2ID, Rank: 2, Level: 42}},
	}
	return in
}

// withDamageTaken is s with guid's damage-taken row holding one hit of spellID at school.
func withDamageTaken(s summary.Summary, guid string, spellID int64, name string, school int64) summary.Summary {
	s.DamageTaken = slices.DeleteFunc(slices.Clone(s.DamageTaken), func(a summary.Actor) bool { return a.GUID == guid })
	s.DamageTaken = append(s.DamageTaken, summary.Actor{GUID: guid, Abilities: []summary.Ability{
		{SpellID: spellID, Name: name, School: school, Total: 2_400, Hits: 1},
	}})
	return s
}

// Evasion only raises dodge: a rogue killed by a Fireball is never told Evasion would
// have saved them. Vanish, which drops the fight, is still named.
func TestDeath_EvasionIsNotOfferedAtASpellDeath(t *testing.T) {
	in := rogueDeath(t, fireballBoltID, "Fireball", schoolFire)
	f := death(in, CountedTime(in)).Finding
	if f == nil || f.Title != "Died at 6:40 with Vanish ready" {
		t.Fatalf("a fire death: got %+v", f)
	}
}

// A melee swing is what Evasion stops: it is offered first, as the rule file lists it.
func TestDeath_EvasionIsOfferedAtAMeleeDeath(t *testing.T) {
	in := rogueDeath(t, meleeSpellID, "Melee", schoolPhysical)
	f := death(in, CountedTime(in)).Finding
	if f == nil || f.Title != "Died at 6:40 with Evasion ready" {
		t.Fatalf("a melee death: got %+v", f)
	}
}

// A physical ability is an attack Evasion can dodge.
func TestDeath_EvasionIsOfferedAtAPhysicalAbilityDeath(t *testing.T) {
	in := rogueDeath(t, knockAwayID, "Knock Away", schoolPhysical)
	f := death(in, CountedTime(in)).Finding
	if f == nil || f.Title != "Died at 6:40 with Evasion ready" {
		t.Fatalf("a physical ability death: got %+v", f)
	}
}

// A killing spell whose school the log did not keep is not proven physical: Evasion is
// not offered on a guess.
func TestDeath_EvasionIsNotOfferedWhenTheSchoolIsUnknown(t *testing.T) {
	in := rogueDeath(t, fireballBoltID, "Fireball", schoolFire)
	in.Fight.DamageTaken = nil
	f := death(in, CountedTime(in)).Finding
	if f == nil || f.Title != "Died at 6:40 with Vanish ready" {
		t.Fatalf("an unknown school: got %+v", f)
	}
}

// The client data decides which defensives are melee-only: a spell whose every effect
// raises dodge or parry (Evasion, Deterrence), never one with any other effect.
func TestParseSpellConstMarksDodgeAndParryOnlySpellsMeleeOnly(t *testing.T) {
	spells, err := ParseSpellConst([]byte(`{"spells":{
		"5277":{"name":"Evasion","category_cooldown_ms":300000,"effects":[{"effect":6,"aura":49}]},
		"19263":{"name":"Deterrence","cooldown_ms":300000,"effects":[{"effect":6,"aura":47},{"effect":6,"aura":49}]},
		"871":{"name":"Shield Wall","cooldown_ms":900000,"effects":[{"effect":6,"aura":87}]},
		"1856":{"name":"Vanish","category_cooldown_ms":300000,"effects":[{"effect":64},{"effect":79}]},
		"1":{"name":"No effects"}}}`))
	if err != nil {
		t.Fatal(err)
	}
	for id, want := range map[int64]bool{5277: true, 19263: true, 871: false, 1856: false, 1: false} {
		if got := spells[id].MeleeOnly; got != want {
			t.Errorf("%s MeleeOnly = %v, want %v", spells[id].Name, got, want)
		}
	}
}
