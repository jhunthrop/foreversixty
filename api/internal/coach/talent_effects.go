package coach

import "slices"

// PlayerTalents is the player's talent ranks by talent name, read from the character's
// stored build (the addon export or the Battle.net import): the log never carries talents.
// Nil means unknown, and the checks then fall back to what holds without them.
type PlayerTalents map[string]int

// cooldownTalent is a talent that shortens a rule-file family's cooldown: by a flat amount
// a rank, or by a share of the base cooldown a rank. Every value is the client's own talent
// text in 1.60.1.70790 (talents/<class>.json; TestCooldownTalentsMatchTheClientText).
type cooldownTalent struct {
	Talent     string
	Families   []string
	CutMSRank  int64   // milliseconds off a rank
	CutPctRank float64 // percent of the base cooldown off a rank
}

// cooldownTalents is every talent that shortens a cooldown a coach rule file judges (F1's
// core and efficiency lists, F7's defensives).
var cooldownTalents = []cooldownTalent{
	{Talent: "Rapid Killing", Families: []string{"Rapid Fire"}, CutMSRank: 60_000},
	{Talent: "Survivalist's Discipline", Families: []string{"Deterrence"}, CutPctRank: 20},
	{Talent: "Improved Fire Blast", Families: []string{"Fire Blast"}, CutMSRank: 1_000},
	{Talent: "Improved Judgement", Families: []string{"Judgement"}, CutMSRank: 1_000},
	{Talent: "Shieldsworn", Families: []string{"Holy Strike"}, CutMSRank: 4_000},
	{Talent: "Sacred Duty", Families: []string{"Divine Shield", "Divine Protection", "Templar's Bulwark"}, CutMSRank: 30_000},
	{Talent: "Endurance", Families: []string{"Evasion"}, CutPctRank: 30},
	{Talent: "Elusiveness", Families: []string{"Vanish"}, CutMSRank: 45_000},
	{Talent: "Improved Whirlwind", Families: []string{"Whirlwind"}, CutMSRank: 1_000},
	{Talent: "Improved Shield Wall", Families: []string{"Shield Wall"}, CutMSRank: 330_000},
}

// cooldownFor is family's cooldown for this player: the spell data's base cooldown less
// what the player's talents take off. shortenable is true when the talents are unknown and
// a talent could shorten it, so the base cooldown is only the longest it can be.
func cooldownFor(in Input, family string, sc SpellConst) (ms int64, shortenable bool) {
	ms = sc.CooldownMS
	for _, ct := range cooldownTalents {
		if !slices.Contains(ct.Families, family) {
			continue
		}
		if in.Tables.Talents == nil {
			shortenable = true
			continue
		}
		rank := int64(in.Tables.Talents[ct.Talent])
		ms -= rank*ct.CutMSRank + int64(float64(sc.CooldownMS)*ct.CutPctRank*float64(rank)/100)
	}
	return max(ms, 0), shortenable
}

// shardGate is a core ability that costs a Soul Shard. The log carries no shard count, so
// it is judged only where every cast gives its shard back: the talent at its full rank, or,
// with the talents unknown, from the level both shipped orders fill it.
type shardGate struct {
	Talent string
	Rank   int
	Level  int64
}

// shardGates: Shadowburn costs a Soul Shard; Shadow and Flame refunds it on 20% of casts a
// rank (426316, 100% at 5/5), which both Destruction orders fill by 44 (raid 42, leveling
// 44; TestShardGateLevelIsWhereTheShippedOrdersFillTheTalent).
var shardGates = map[string]shardGate{
	"Shadowburn": {Talent: "Shadow and Flame", Rank: 5, Level: 44},
}

// shardShort is true when family costs a Soul Shard and nothing shows the player had one
// for every cast: the refund talent below full rank, or, with talents unknown, a level
// unknown or under the gate's.
func shardShort(in Input, family string) bool {
	g, ok := shardGates[family]
	if !ok {
		return false
	}
	if in.Tables.Talents != nil {
		return in.Tables.Talents[g.Talent] < g.Rank
	}
	return in.Row.Level < g.Level
}
