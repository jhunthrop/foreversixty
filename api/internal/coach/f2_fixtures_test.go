package coach

import (
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	f2RendID         = 11977
	f2FightMS        = 232_000
	f2FirstCastMS    = 0
	f2CastSpellID    = 23894
	f2CastSpellName  = "Bloodthirst"
	f2AuraSegmentMS  = 3
	f2SecondaryGUID  = "Player-1136-0AAAAAAA"
	f2ThirdGUID      = "Player-1136-0BBBBBBB"
	f2BossGUID       = "Creature-0-1-1084-1-10184-0000000001"
	f2BattleShoutID  = 6673
	f2SunderArmorID  = 11597
	f2ShieldBlockID  = 2565
	f2RendName       = "Rend"
	f2BattleShoutNme = "Battle Shout"
)

// f2Steady is in with a fight of f2FightMS in which the player lives the whole time, casts
// from the pull and has no curated downtime, so an aura's uptime is its plain share of
// the fight.
func f2Steady(in Input) Input {
	s := withDuration(in.Fight, f2FightMS)
	s.Deaths = slices.DeleteFunc(slices.Clone(s.Deaths), func(d summary.Death) bool { return d.GUID == in.Row.GUID })
	s = withoutCasts(s, in.Row.GUID)
	s = withCasts(s, in.Row.GUID, f2CastSpellID, f2CastSpellName, f2FirstCastMS)
	in.Fight = s
	in.Tables.Mechanics = nil
	return in
}

// f2Segments is upMS of uptime cut into f2AuraSegmentMS equal segments spread over the
// fight, the first starting at the pull.
func f2Segments(upMS int64) []summary.Segment {
	each := upMS / f2AuraSegmentMS
	stride := int64(f2FightMS) / f2AuraSegmentMS
	segs := make([]summary.Segment, f2AuraSegmentMS)
	for i := range segs {
		from := int64(i) * stride
		segs[i] = span(from, from+each)
	}
	return segs
}

// protShieldBlock is the Protection warrior with Shield Block up upMS of durationMS.
func protShieldBlock(t *testing.T, upMS, durationMS int64, refs Refs) Input {
	t.Helper()
	in := f2Steady(protInput(t))
	if durationMS != f2FightMS {
		t.Fatalf("protShieldBlock: the fixture fight is %d ms", f2FightMS)
	}
	in.Fight = withAura(in.Fight, furyGUID, f2ShieldBlockID, "Shield Block", []string{furyGUID}, f2Segments(upMS)...)
	return withPool(in, refs)
}

// furyBattleShout is the Fury warrior with Battle Shout up the fraction of the fight.
func furyBattleShout(t *testing.T, fraction float64, refs Refs) Input {
	t.Helper()
	in := f2Steady(furyInput(t))
	in.Fight = withAura(in.Fight, furyGUID, f2BattleShoutID, f2BattleShoutNme, []string{furyGUID},
		f2Segments(int64(fraction*f2FightMS))...)
	return withPool(in, refs)
}

// furySunderShared is the Fury warrior whose Sunder Armor on the boss was applied by
// appliers warriors, himself one of them.
func furySunderShared(t *testing.T, appliers int) Input {
	t.Helper()
	in := f2Steady(furyInput(t))
	names := []string{furyGUID, f2SecondaryGUID, f2ThirdGUID}
	in.Fight = withAura(in.Fight, f2BossGUID, f2SunderArmorID, "Sunder Armor", names[:appliers], span(0, f2FightMS))
	return withPool(in, poolUptime(f2SunderArmorID, 90, PoolTopN))
}

const (
	f2ThunderClapID   = 11581
	f2ThunderClapName = "Thunder Clap"
)

// protThunderClap is the Protection warrior with Thunder Clap, listed as a defensive in his
// rule file, applied by him alone to the boss for upMS of the fight.
func protThunderClap(t *testing.T, upMS int64, refs Refs) Input {
	t.Helper()
	in := f2Steady(protInput(t))
	in.Tables.Rules.Auras = append(slices.Clone(in.Tables.Rules.Auras),
		AuraRule{SpellID: f2ThunderClapID, Name: f2ThunderClapName, Kind: AuraDefensive, Verified: fixtureVerified})
	in.Fight = withAura(in.Fight, f2BossGUID, f2ThunderClapID, f2ThunderClapName, []string{furyGUID}, f2Segments(upMS)...)
	return withPool(in, refs)
}

// furyRendLow is the Fury warrior with a Rend rule and Rend up a fifth of the fight, on
// refs.
func furyRendLow(t *testing.T, refs Refs) Input {
	t.Helper()
	in := f2Steady(furyInput(t))
	in.Tables.Rules.Auras = append(slices.Clone(in.Tables.Rules.Auras),
		AuraRule{SpellID: f2RendID, Name: f2RendName, Kind: AuraDot, Verified: fixtureVerified})
	in.Fight = withAura(in.Fight, f2BossGUID, f2RendID, f2RendName, []string{furyGUID}, f2Segments(f2FightMS/5)...)
	return withPool(in, refs)
}
