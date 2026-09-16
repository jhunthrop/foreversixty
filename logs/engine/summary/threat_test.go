// logs/engine/summary/threat_test.go
package summary

import (
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// Two enemies, one healer. Damage threat lands on the unit hit; healing threat is
// spread over the enemies engaged in the last ten seconds; a taunt is recorded.
func TestThreatIsCreditedPerTargetAndTauntsAreKept(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	events := []event.Event{
		dmg(1, tank, boss, 1, "Melee", 1000, -1),
		dmg(2, mage, bossB, 116, "Frostbolt", 500, -1),
		heal(3, healer, tank, 2050, "Holy Light", 400, 0),
		cast(4, tank, boss, 355, "Taunt"),
	}
	for _, e := range events {
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(fight.Fight{Index: 1, Kind: fight.Encounter, Start: at(0), End: at(5), Players: []string{tank, mage, healer}}, "test")

	pair := func(guid, target string) float64 {
		for _, p := range s.ThreatByTarget {
			if p.GUID == guid && p.TargetGUID == target {
				return p.Threat
			}
		}
		return -1
	}
	if got := pair(tank, boss); got != 1000 {
		t.Errorf("tank on boss = %v, want 1000 (a point of damage is a point of threat)", got)
	}
	if got := pair(mage, bossB); got != 500 {
		t.Errorf("mage on bossB = %v, want 500", got)
	}
	// 400 effective healing at 0.5 = 200 threat, over the two engaged enemies: 100 each.
	if got := pair(healer, boss); got != 100 {
		t.Errorf("healer on boss = %v, want 100", got)
	}
	if got := pair(healer, bossB); got != 100 {
		t.Errorf("healer on bossB = %v, want 100", got)
	}
	// The per-player totals are untouched by attribution.
	for _, row := range s.Threat {
		if row.GUID == healer && row.Threat != 200 {
			t.Errorf("healer total = %v, want 200", row.Threat)
		}
	}
	if len(s.Taunts) != 1 || s.Taunts[0].SourceGUID != tank || s.Taunts[0].TargetGUID != boss || s.Taunts[0].AtMS != 4000 {
		t.Fatalf("taunts = %+v", s.Taunts)
	}
}

// A neutral-reaction enemy (the sample's Tormented Souls, flags 0xa28) is not
// friendly, so its damage is damage done and its threat is real threat. It must
// reach the per-target table and the healers' spread, not only the player's total.
func TestNeutralEnemiesGetTheirOwnThreatPair(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	neutral := event.Unit{GUID: bossB, Name: "Tormented Soul", Flags: 0xa28}
	events := []event.Event{
		{Time: at(1), Kind: event.Damage, Name: "SPELL_DAMAGE",
			Source: event.Unit{GUID: mage, Flags: 0x512}, Dest: neutral,
			Spell:    event.Spell{ID: 116, Name: "Frostbolt", School: 0x10},
			Amount:   event.OptInt{V: 600, OK: true},
			Overkill: event.OptInt{V: -1, OK: true}},
		heal(2, healer, tank, 2050, "Holy Light", 400, 0),
	}
	for _, e := range events {
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(fight.Fight{Index: 1, Kind: fight.Encounter, Start: at(0), End: at(5), Players: []string{tank, mage, healer}}, "test")

	pair := func(guid, target string) float64 {
		for _, p := range s.ThreatByTarget {
			if p.GUID == guid && p.TargetGUID == target {
				return p.Threat
			}
		}
		return -1
	}
	if got := pair(mage, bossB); got != 600 {
		t.Errorf("mage on the neutral unit = %v, want 600: a pair is credited wherever the total is", got)
	}
	// The neutral unit is the only thing engaged, so it takes the whole 200 of spread
	// healing threat: 400 effective at 0.5.
	if got := pair(healer, bossB); got != 200 {
		t.Errorf("healer on the neutral unit = %v, want 200: a neutral enemy engages", got)
	}
}

// The environment (a fall, a fire) has the nil GUID and no reaction bit at all: it is
// nobody's enemy, so it neither takes a pair nor engages for the healing spread.
func TestTheEnvironmentIsNobodysEnemy(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	events := []event.Event{
		{Time: at(1), Kind: event.Damage, Name: "ENVIRONMENTAL_DAMAGE",
			Source: event.Unit{GUID: units.NoGUID, Flags: 0x80000000}, Dest: event.Unit{GUID: tank, Flags: 0x511},
			Spell:    event.Spell{ID: -1, Name: "Falling", School: 1},
			Amount:   event.OptInt{V: 1000, OK: true},
			Overkill: event.OptInt{V: -1, OK: true}},
		heal(2, healer, tank, 2050, "Holy Light", 400, 0),
	}
	for _, e := range events {
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(fight.Fight{Index: 1, Kind: fight.Encounter, Start: at(0), End: at(5), Players: []string{tank, healer}}, "test")
	for _, p := range s.ThreatByTarget {
		if p.TargetGUID == units.NoGUID || p.TargetName == "Environment" {
			t.Errorf("the environment took a threat pair: %+v", p)
		}
	}
}

// A heal long after the last blow lands on nobody: the engaged window has expired,
// so the threat stays on the player's total and no pair is invented for it.
func TestHealingSpreadsOverNobodyOnceTheWindowExpires(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	// DefaultOptions' EngagedWindow is ten seconds; the heal is thirty seconds later.
	events := []event.Event{
		dmg(1, mage, boss, 116, "Frostbolt", 500, -1),
		heal(31, healer, tank, 2050, "Holy Light", 400, 0),
	}
	for _, e := range events {
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(fight.Fight{Index: 1, Kind: fight.Encounter, Start: at(0), End: at(35), Players: []string{tank, mage, healer}}, "test")

	for _, p := range s.ThreatByTarget {
		if p.GUID == healer {
			t.Errorf("the healer holds a pair %+v, want none: nothing was engaged when the heal landed", p)
		}
	}
	for _, row := range s.Threat {
		if row.GUID == healer && row.Threat != 200 {
			t.Errorf("healer total = %v, want 200: the total is untouched by having nowhere to spread", row.Threat)
		}
	}
}

// A taunt on the nil GUID is the game saying "no target". It renders as
// "Environment", which is nothing a tank taunted, so it is not a taunt row.
func TestTauntsWithNoDestinationAreNotKept(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	nilTaunt := cast(1, tank, units.NoGUID, 355, "Taunt")
	nilTaunt.Dest.Flags = 0
	events := []event.Event{
		nilTaunt,
		cast(2, tank, "", 355, "Taunt"),
		cast(3, tank, boss, 355, "Taunt"),
	}
	for _, e := range events {
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(fight.Fight{Index: 1, Kind: fight.Encounter, Start: at(0), End: at(5), Players: []string{tank}}, "test")
	if len(s.Taunts) != 1 || s.Taunts[0].TargetGUID != boss {
		t.Fatalf("taunts = %+v, want only the one aimed at a real unit", s.Taunts)
	}
}

// The three area and ally-cast taunts are out of the list until Forever's own ids
// land: a taunt row names the enemy it pulled, and these name nothing or a raider.
func TestUntargetedAndAllyCastTauntsAreNotKept(t *testing.T) {
	for _, spell := range []struct {
		id   int64
		name string
	}{
		{1161, "Challenging Shout"},
		{5209, "Challenging Roar"},
		{31789, "Righteous Defense"},
	} {
		if tauntSpells[spell.id] {
			t.Errorf("%s (%d) is in tauntSpells; it names the wrong target", spell.name, spell.id)
		}
	}
}

// A taunt cast on the pull lands its debuff after the fight's first event and its
// cast before it, so the debuff is the only trace the fight holds. Inside the fight
// the cast is the record and its debuff, a millisecond later, is not a second taunt.
func TestTauntSeenOnlyByItsDebuffIsKeptOnce(t *testing.T) {
	o, reg := opts(t)
	a := New(o)
	a.Start(at(0))
	debuff := func(sec float64) event.Event {
		e := cast(sec, tank, boss, 355, "Taunt")
		e.Kind, e.Name = event.AuraApplied, "SPELL_AURA_APPLIED"
		return e
	}
	events := []event.Event{
		debuff(0.001),
		cast(3, tank, boss, 355, "Taunt"),
		debuff(3.002),
	}
	for _, e := range events {
		reg.Observe(e)
		a.Add(e)
	}
	s := a.Snapshot(fight.Fight{Index: 1, Kind: fight.Encounter, Start: at(0), End: at(5), Players: []string{tank}}, "test")
	if len(s.Taunts) != 2 || s.Taunts[0].AtMS != 1 || s.Taunts[1].AtMS != 3000 {
		t.Fatalf("taunts = %+v, want the opening debuff and the later cast, once each", s.Taunts)
	}
}
