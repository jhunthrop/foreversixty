package coach

import (
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	prepFightMS   = 60_000 // short enough that one combat potion is all the fight could hold
	slotChest     = 4
	slotFeet      = 7
	slotBack      = 14
	slotMainHand  = 15
	gearSlotCount = 16
	itemEnchanted = 1900
)

// preparedGear is a full gear list with the four checked slots enchanted, and the main
// hand too when mainHand is true.
func preparedGear(mainHand bool) []event.Item {
	gear := make([]event.Item, gearSlotCount)
	for _, slot := range []int{slotChest, slotFeet, slotBack, slotMainHand} {
		gear[slot] = event.Item{ID: 19364, ItemLevel: 83, Enchants: []int64{itemEnchanted}}
	}
	if !mainHand {
		gear[slotMainHand].Enchants = nil
	}
	return gear
}

func withGear(in Input, gear []event.Item) Input {
	rows := slices.Clone(in.Fight.Combatants)
	for i := range rows {
		if rows[i].GUID == in.Row.GUID {
			rows[i].Gear = gear
		}
	}
	in.Fight.Combatants = rows
	return in
}

// withUnenchantedMainHand is in with the player's main-hand entry carrying no permanent
// enchant.
func withUnenchantedMainHand(in Input) Input { return withGear(in, preparedGear(false)) }

// seedPotion marks the catalogue's first combat potion as seeded on guid at the pull.
func seedPotion(s summary.Summary, guid string) summary.Summary {
	e := consumables.Load().AllRoles().CombatPotion.Entries[0]
	s = withAura(s, guid, e.SpellID, e.Name, []string{guid}, span(0, s.DurationMS))
	tracks := slices.Clone(s.Auras)
	for i := range tracks {
		if tracks[i].TargetGUID == guid && tracks[i].SpellID == e.SpellID {
			tracks[i].Seeded = true
		}
	}
	s.Auras = tracks
	return s
}

// furyPrep is a Fury warrior alone in a raid of the given size, fully enchanted, with
// Battle Shout and the given seeds.
func furyPrep(t *testing.T, size int, ss seeds) Input {
	t.Helper()
	in := furyInput(t)
	in.Size = size
	s := withDuration(in.Fight, prepFightMS)
	s = withRoster(s, in.Row)
	s = withSeeds(s, in.Row.GUID, ss)
	in.Fight = s
	return withRaidBuff(t, withGear(in, preparedGear(true)), "Battle Shout")
}

// rosterOf is in with the player first and rows after.
func rosterOf(in Input, rows ...summary.RosterRow) Input {
	in.Fight = withRoster(in.Fight, append([]summary.RosterRow{in.Row}, rows...)...)
	return in
}

func raidBuffByName(t *testing.T, in Input, name string) consumables.RaidBuff {
	t.Helper()
	for _, b := range in.Tables.RaidBuffs {
		if b.Name == name {
			return b
		}
	}
	t.Fatalf("no raid buff %q", name)
	return consumables.RaidBuff{}
}

// withRaidBuff is in with the player carrying the buff at the pull.
func withRaidBuff(t *testing.T, in Input, name string) Input {
	t.Helper()
	b := raidBuffByName(t, in, name)
	in.Fight = withAura(in.Fight, in.Row.GUID, b.SpellIDs[0], b.Name, []string{in.Row.GUID}, span(0, in.Fight.DurationMS))
	return in
}

func rosterClass(guid, class string) summary.RosterRow {
	return summary.RosterRow{GUID: guid, Name: guid, Class: class, Role: RoleDPS}
}

// tenManNoPriest is a Beast Mastery hunter in a 10-player roster of warriors, rogues and
// one mage, with no priest: flask, food and a potion up, no raid buffs.
func tenManNoPriest(t *testing.T) Input {
	t.Helper()
	in := furyPrep(t, 10, seededAt(flask("Flask of the Titans"), food("Grilled Squid")))
	in.Row = withRole(in.Fight, in.Row.GUID, "Hunter", "Beast Mastery", RoleDPS).Roster[0]
	rows := []summary.RosterRow{rosterClass("mage", "Mage")}
	for i := range 4 {
		rows = append(rows, rosterClass("warrior"+string(rune('a'+i)), "Warrior"), rosterClass("rogue"+string(rune('a'+i)), "Rogue"))
	}
	return rosterOf(in, rows...)
}

func catalogueWorldBuff(t *testing.T) consumables.Entry {
	t.Helper()
	wb := consumables.Load().AllRoles().WorldBuffs
	if len(wb) == 0 {
		t.Fatal("no world buffs in the catalogue")
	}
	return wb[0]
}
