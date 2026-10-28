package coach

import (
	"github.com/jhunthrop/foreversixty/api/internal/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
)

const (
	millisPerMinute = 60_000.0
	percentScale    = 100.0
)

// profileCategories are the consumable categories a profile records at the pull.
var profileCategories = []string{
	consumables.CategoryFlask, consumables.CategoryFood,
	consumables.CategoryWeaponEnchant, consumables.CategoryCombatPotion,
}

// BuildProfile measures one player's fight in the units the pack compares (§11.1). Pure.
func BuildProfile(in Input) Profile {
	c := CountedTime(in)
	p := Profile{
		AliveMS: c.AliveMS, CountedMS: c.CountedMS, ActivePct: countedActivePct(in, c),
		CastsPerMin: castsPerMinute(in, c), Uptime: map[string]float64{},
		OverhealByRank: overhealByRank(in), DefensiveUptime: map[string]float64{},
		Consumables: consumablesAtPull(in),
	}
	measureAuras(in, c, &p)
	measureResources(in, c, &p)
	if in.Row.Role != RoleHealer && in.Tables.RotationOK {
		p.Adherence = rating.FightAdherence(in.Fight, in.Row, in.Tables.Rotation)
	}
	return p
}

// castsPerMinute is each core (efficiency for a healer) family's counted casts per counted
// minute. A talent never cast is left out: one cast proves the talent (§6.1).
func castsPerMinute(in Input, c Counted) map[string]float64 {
	out := map[string]float64{}
	if c.CountedMS <= 0 {
		return out
	}
	for _, e := range coreEntries(in) {
		seq, _ := castsOf(in.Fight, in.Row.GUID, e.SpellIDs)
		counted := countedCasts(seq, c)
		if counted == 0 && e.Talent {
			continue
		}
		out[e.Family] = float64(counted) * millisPerMinute / float64(c.CountedMS)
	}
	return out
}

// countedCasts is how many of the cast instants fall in counted time.
func countedCasts(seq []int64, c Counted) int {
	n := 0
	for _, at := range seq {
		if at < c.AliveMS && countedBetween(c.Excluded, at, at+1) > 0 {
			n++
		}
	}
	return n
}

// measureAuras fills Uptime (owned auras, sole applier only) and DefensiveUptime.
func measureAuras(in Input, c Counted, p *Profile) {
	for _, rule := range in.Tables.Rules.Auras {
		track, appliers := ownTrack(in, rule)
		if track == nil || appliers > 1 {
			continue
		}
		up, den, ok := measureUptime(*track, c)
		if !ok {
			continue
		}
		target := p.Uptime
		if rule.Kind == AuraDefensive {
			target = p.DefensiveUptime
		}
		target[SpellKey(rule.SpellID)] = float64(up) / float64(den) * percentScale
	}
}

// measureResources fills the capped share (rage and energy, the worse bar) and the mana
// stretch under ManaLowPct, over counted seconds only.
func measureResources(in Input, c Counted, p *Profile) {
	if c.CountedMS <= 0 {
		return
	}
	for _, power := range []int64{PowerRage, PowerEnergy} {
		if track, ok := trackOf(in.Fight, in.Row.GUID, power); ok && track.Max > 0 {
			runs := countedRuns(track.Series, c, func(v int64) bool { return v >= track.Max })
			p.CappedPct = max(p.CappedPct, float64(totalMS(runs))*percentScale/float64(c.CountedMS))
		}
	}
	track, ok := trackOf(in.Fight, in.Row.GUID, PowerMana)
	if !ok || track.Max <= 0 {
		return
	}
	floor := ManaLowPct * float64(track.Max)
	runs := countedRuns(track.Series, c, func(v int64) bool { return float64(v)*percentScale < floor })
	p.ManaLowMS = totalMS(runs)
	if len(runs) > 0 {
		first := runs[0].FromMS
		p.ManaLowFirstMS = &first
	}
}

// overhealByRank is overheal % per heal spell id for a healer's ranked heals cast often
// enough to compare (HealRankMinCasts).
func overhealByRank(in Input) map[string]float64 {
	out := map[string]float64{}
	if in.Row.Role != RoleHealer {
		return out
	}
	for _, a := range playerHealing(in) {
		if a.Via != "" || a.Total <= 0 || succeededCasts(in, a.SpellID) < HealRankMinCasts {
			continue
		}
		out[SpellKey(a.SpellID)] = float64(a.Overheal) / float64(a.Total) * percentScale
	}
	return out
}

// consumablesAtPull is category -> present at the pull, for the categories the role's
// catalogue lists; nil when the log has no character sheet for the player.
func consumablesAtPull(in Input) map[string]bool {
	pull := consumables.AtPull(in.Fight, in.Row.GUID, in.Tables.Consumables)
	if !pull.Known {
		return nil
	}
	out := map[string]bool{}
	for _, cat := range profileCategories {
		if len(in.Tables.Consumables.Entries(cat)) > 0 {
			out[cat] = pull.Has(cat)
		}
	}
	return out
}
