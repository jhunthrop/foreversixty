// logs/engine/fight/counts_test.go
package fight

import (
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
)

// Unit flags as the game writes them: reaction, type and affiliation.
const (
	flagMine     uint32 = 0x511 // a friendly player, the logger
	flagParty    uint32 = 0x512 // a friendly player in the logger's party
	flagRaid     uint32 = 0x514 // a friendly player in the logger's raid
	flagOutsider uint32 = 0x518 // a friendly player outside the group
	flagHostile  uint32 = 0xa48 // a hostile NPC
	flagNeutral  uint32 = 0xa28 // a neutral NPC: a yellow beast, or a critter
)

const (
	logger   = "Player-4184-000000B1"
	stranger = "Player-4184-000000C9"
	mob2     = "Creature-0-2085-2284-7855-169753-0000AA0002"
	critter  = "Creature-0-2085-2284-7855-61071-0000AA0003"
	beast    = "Creature-0-2085-2284-7855-250926-0000AA0005"
	beast2   = "Creature-0-2085-2284-7855-250926-0000AA0006"
	totem    = "Creature-0-2085-2284-7855-5925-0000AA0004"
)

func died(sec float64, guid string, flags uint32) event.Event {
	return event.Event{Time: at(sec), Kind: event.Death, Name: "UNIT_DIED",
		Dest: event.Unit{GUID: guid, Flags: flags}}
}

func partyKill(sec float64, src string, srcFlags uint32, dst string, dstFlags uint32) event.Event {
	return event.Event{Time: at(sec), Kind: event.PartyKill, Name: "PARTY_KILL",
		Source: event.Unit{GUID: src, Flags: srcFlags},
		Dest:   event.Unit{GUID: dst, Flags: dstFlags}}
}

func closeOne(t *testing.T, evs []event.Event, end float64) *Fight {
	t.Helper()
	s := NewSegmenter(Options{MinTrash: 0})
	run(s, evs)
	f := s.Flush(at(end))
	if f == nil {
		t.Fatal("no fight")
	}
	return f
}

func TestAGroupKillCountsOnceAsADeathAndOnceAsTheGroups(t *testing.T) {
	f := closeOne(t, []event.Event{
		hit(0, logger, mob, flagMine, flagHostile),
		partyKill(1, logger, flagMine, mob, flagHostile),
		died(1, mob, flagHostile),
	}, 2)
	if f.MobDeaths != 1 || f.GroupKills != 1 {
		t.Errorf("mob deaths = %d, group kills = %d; want 1 and 1 (PARTY_KILL and UNIT_DIED are one death)",
			f.MobDeaths, f.GroupKills)
	}
}

func TestAMobSomeoneElseKilledDiedButIsNotTheGroups(t *testing.T) {
	f := closeOne(t, []event.Event{
		hit(0, logger, mob, flagMine, flagHostile),
		hit(1, stranger, mob2, flagOutsider, flagHostile),
		partyKill(2, logger, flagMine, mob, flagHostile),
		died(2, mob, flagHostile),
		died(3, mob2, flagHostile),
	}, 4)
	if f.MobDeaths != 2 || f.GroupKills != 1 {
		t.Errorf("mob deaths = %d, group kills = %d; want 2 and 1", f.MobDeaths, f.GroupKills)
	}
}

func TestTheGroupsKillingBlowCountsWithoutAPartyKill(t *testing.T) {
	// An outsider opened on the mob; a raid member's hit was the last it took.
	f := closeOne(t, []event.Event{
		hit(0, stranger, mob, flagOutsider, flagHostile),
		hit(1, mage, mob, flagRaid, flagHostile),
		died(2, mob, flagHostile),
	}, 3)
	if f.MobDeaths != 1 || f.GroupKills != 1 {
		t.Errorf("mob deaths = %d, group kills = %d; want 1 and 1", f.MobDeaths, f.GroupKills)
	}
}

func TestAnOutsidersKillingBlowIsNotTheGroups(t *testing.T) {
	f := closeOne(t, []event.Event{
		hit(0, mage, mob, flagParty, flagHostile),
		hit(1, stranger, mob, flagOutsider, flagHostile),
		died(2, mob, flagHostile),
	}, 3)
	if f.MobDeaths != 1 || f.GroupKills != 0 {
		t.Errorf("mob deaths = %d, group kills = %d; want 1 and 0", f.MobDeaths, f.GroupKills)
	}
}

func TestTotemsAreNotMobDeaths(t *testing.T) {
	destroyed := died(2, totem, flagHostile)
	destroyed.Name = "UNIT_DESTROYED"
	f := closeOne(t, []event.Event{hit(0, logger, mob, flagMine, flagHostile), destroyed}, 3)
	if f.MobDeaths != 0 || f.GroupKills != 0 {
		t.Errorf("mob deaths = %d, group kills = %d; want 0 and 0", f.MobDeaths, f.GroupKills)
	}
}

// A neutral (yellow) beast keeps its neutral flags while it fights the group and on the
// lines that end it. Levelling zones are full of them, so the group's kill counts.
func TestANeutralBeastTheGroupKilledCounts(t *testing.T) {
	f := closeOne(t, []event.Event{
		withMaxHP(hit(0, logger, beast, flagMine, flagNeutral), beast, 102),
		partyKill(1, logger, flagMine, beast, flagNeutral),
		died(1, beast, flagNeutral),
		withMaxHP(hit(2, logger, beast2, flagMine, flagNeutral), beast2, 120),
		died(3, beast2, flagNeutral), // no PARTY_KILL: the group's last hit decides
	}, 4)
	if f.MobDeaths != 2 || f.GroupKills != 2 {
		t.Errorf("mob deaths = %d, group kills = %d; want 2 and 2", f.MobDeaths, f.GroupKills)
	}
}

// A neutral unit the group never touched is no mob of theirs: someone else's pull, or a
// critter caught in someone's area spell.
func TestANeutralUnitTheGroupNeverHitIsNotCounted(t *testing.T) {
	f := closeOne(t, []event.Event{
		hit(0, logger, mob, flagMine, flagHostile),
		hit(1, stranger, beast, flagOutsider, flagNeutral),
		died(2, beast, flagNeutral),
	}, 3)
	if f.MobDeaths != 0 || f.GroupKills != 0 {
		t.Errorf("mob deaths = %d, group kills = %d; want 0 and 0", f.MobDeaths, f.GroupKills)
	}
}

// A critter is a neutral unit with next to no health. The group killing one, even with a
// PARTY_KILL, is not a mob death.
func TestACritterTheGroupKilledIsNotCounted(t *testing.T) {
	f := closeOne(t, []event.Event{
		hit(0, logger, mob, flagMine, flagHostile),
		withMaxHP(hit(1, logger, critter, flagMine, flagNeutral), critter, 8),
		partyKill(1, logger, flagMine, critter, flagNeutral),
		died(1, critter, flagNeutral),
	}, 2)
	if f.MobDeaths != 0 || f.GroupKills != 0 {
		t.Errorf("mob deaths = %d, group kills = %d; want 0 and 0", f.MobDeaths, f.GroupKills)
	}
}

func TestTheCritterCutoffIsConfigurable(t *testing.T) {
	c := NewCounter(Options{CritterMaxHP: 200})
	for _, e := range []event.Event{
		withMaxHP(hit(0, logger, beast, flagMine, flagNeutral), beast, 102),
		died(1, beast, flagNeutral),
	} {
		c.Observe(e)
	}
	var f Fight
	c.Apply(&f)
	if f.MobDeaths != 0 {
		t.Errorf("mob deaths = %d, want 0 under a 200 HP cutoff", f.MobDeaths)
	}
}

// A mob a group member mind-controlled carries friendly, group flags when it dies. It was
// hostile earlier in the fight, so it is still a mob death.
func TestAMobThatDiesCharmedStillCounts(t *testing.T) {
	charmed := uint32(0x1111)
	f := closeOne(t, []event.Event{
		hit(0, logger, mob, flagMine, flagHostile),
		hit(1, mob2, stranger, charmed, flagOutsider), // a stray hit while charmed, never hostile
		hit(2, mob, logger, charmed, flagMine),
		died(3, mob, charmed),
		died(3, mob2, charmed),
	}, 4)
	if f.MobDeaths != 1 {
		t.Errorf("mob deaths = %d, want 1: the charmed mob seen hostile, not the one never hostile", f.MobDeaths)
	}
}

// If the client logs a hunter's Feign Death as UNIT_DIED, the hunter acting again in the same
// fight undoes it. A real death, followed by a resurrection and the player acting, stays.
func TestAFeignDeathIsUndoneWhenTheHunterActsAgain(t *testing.T) {
	cast := func(sec float64, src string, flags uint32, name string) event.Event {
		return event.Event{Time: at(sec), Kind: event.CastSuccess, Name: name,
			Source: event.Unit{GUID: src, Flags: flags}}
	}
	swing := hit(5, mage, mob, flagParty, flagHostile)
	swing.Name = "SWING_DAMAGE"
	dot := hit(4, logger, mob, flagMine, flagHostile)
	dot.Name = "SPELL_PERIODIC_DAMAGE" // a DoT ticking after its caster died is no sign of life
	f := closeOne(t, []event.Event{
		hit(0, logger, mob, flagMine, flagHostile),
		cast(1, mage, flagParty, "SPELL_CAST_SUCCESS"), // Feign Death
		died(1, mage, flagParty),
		died(2, logger, flagMine),
		dot,
		swing, // the hunter is up: the feign was no death
		{Time: at(6), Kind: event.Resurrect, Name: "SPELL_RESURRECT",
			Source: event.Unit{GUID: stranger, Flags: flagOutsider}, Dest: event.Unit{GUID: logger, Flags: flagMine}},
		cast(7, logger, flagMine, "SPELL_CAST_SUCCESS"), // the raised logger acting
	}, 8)
	if f.Deaths != 1 || f.GroupDeaths != 1 {
		t.Errorf("deaths = %d, group deaths = %d; want 1 and 1 (the logger's, not the feign)", f.Deaths, f.GroupDeaths)
	}
}

func withMaxHP(e event.Event, guid string, maxHP int64) event.Event {
	e.Adv = event.Advanced{OK: true, InfoGUID: guid, CurrentHP: maxHP, MaxHP: maxHP}
	return e
}

func TestADeathLoggedTwiceCountsOnce(t *testing.T) {
	f := closeOne(t, []event.Event{
		hit(0, logger, mob, flagMine, flagHostile),
		died(1, mob, flagHostile),
		died(1, mob, flagHostile),
	}, 2)
	if f.MobDeaths != 1 {
		t.Errorf("mob deaths = %d, want 1", f.MobDeaths)
	}
}

func TestGroupDeathsCountOnlyTheGroup(t *testing.T) {
	f := closeOne(t, []event.Event{
		hit(0, logger, mob, flagMine, flagHostile),
		hit(0, stranger, mob, flagOutsider, flagHostile),
		died(1, logger, flagMine),
		died(2, stranger, flagOutsider),
	}, 3)
	if f.Deaths != 2 || f.GroupDeaths != 1 {
		t.Errorf("deaths = %d, group deaths = %d; want 2 and 1", f.Deaths, f.GroupDeaths)
	}
}

func TestTheGroupIsWhoWasInItDuringEachFight(t *testing.T) {
	// The mage is in the party for the first pull and has left it by the second.
	s := NewSegmenter(Options{MinTrash: 0, Gap: 5e9})
	closed := run(s, []event.Event{
		hit(0, logger, mob, flagMine, flagHostile),
		hit(1, mage, mob, flagParty, flagHostile),
		hit(2, stranger, mob, flagOutsider, flagHostile),
		hit(60, logger, mob2, flagMine, flagHostile),
		hit(61, mage, mob2, flagOutsider, flagHostile),
	})
	second := s.Flush(at(62))
	if len(closed) != 1 || second == nil {
		t.Fatalf("closed %d fights then %v, want two fights", len(closed), second)
	}
	if want := []string{logger, mage}; !slices.Equal(closed[0].Group, sorted(want)) {
		t.Errorf("first pull's group = %v, want %v", closed[0].Group, sorted(want))
	}
	if want := []string{logger}; !slices.Equal(second.Group, want) {
		t.Errorf("second pull's group = %v, want %v", second.Group, want)
	}
	if !slices.Contains(second.Players, mage) {
		t.Errorf("the mage was still in range in the second pull: players = %v", second.Players)
	}
}

func TestPetsAreNotGroupMembers(t *testing.T) {
	pet := "Pet-0-4184-0-0-416-0100000001"
	f := closeOne(t, []event.Event{
		hit(0, pet, mob, 0x1111, flagHostile),
		hit(1, logger, mob, flagMine, flagHostile),
	}, 2)
	if !slices.Equal(f.Group, []string{logger}) {
		t.Errorf("group = %v, want only the logger", f.Group)
	}
}

func TestAPetsKillingBlowIsTheGroups(t *testing.T) {
	pet := "Pet-0-4184-0-0-416-0100000001"
	f := closeOne(t, []event.Event{
		hit(0, pet, mob, 0x1111, flagHostile),
		died(1, mob, flagHostile),
	}, 2)
	if f.GroupKills != 1 {
		t.Errorf("group kills = %d, want 1 (the logger's pet carries the logger's affiliation)", f.GroupKills)
	}
}

func TestCountsSurviveTheOpenSnapshot(t *testing.T) {
	s := NewSegmenter(Options{MinTrash: 0})
	run(s, []event.Event{
		hit(0, logger, mob, flagMine, flagHostile),
		partyKill(1, logger, flagMine, mob, flagHostile),
		died(1, mob, flagHostile),
	})
	open := s.Open()
	if open == nil || open.MobDeaths != 1 || open.GroupKills != 1 || len(open.Group) != 1 {
		t.Fatalf("open fight = %+v, want its counts so far", open)
	}
}

func TestCounterAppliesToAFightBuiltByHand(t *testing.T) {
	c := NewCounter(Options{})
	for _, e := range []event.Event{
		hit(0, logger, mob, flagMine, flagHostile),
		partyKill(1, logger, flagMine, mob, flagHostile),
		died(1, mob, flagHostile),
		died(2, logger, flagMine),
	} {
		c.Observe(e)
	}
	var f Fight
	c.Apply(&f)
	if f.MobDeaths != 1 || f.GroupKills != 1 || f.Deaths != 1 || f.GroupDeaths != 1 {
		t.Errorf("counts = %d/%d/%d/%d, want 1/1/1/1", f.MobDeaths, f.GroupKills, f.Deaths, f.GroupDeaths)
	}
	if !slices.Equal(f.Players, []string{logger}) || !slices.Equal(f.Group, []string{logger}) {
		t.Errorf("players = %v, group = %v", f.Players, f.Group)
	}
}

func sorted(s []string) []string {
	out := slices.Clone(s)
	slices.Sort(out)
	return out
}
