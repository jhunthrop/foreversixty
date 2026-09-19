package adapter

// Folding rows onto their key.
//
// A summary has one row per identity: the logs engine accumulates a
// fight into maps keyed by {actor, spell, via}, {caster, spell} and
// {target, spell}, so a real fight can never produce two rows that
// share one. An engine result can: it carries two AuraMetrics for one
// aura the spec registered twice, and would carry two ActionMetrics for
// one action if a spec ever registered that twice too. Deriving a
// distinct id for each would invent a difference the fight did not
// have, so identical rows are added together here instead, which is
// exactly what the logs engine does with two events of one spell.

import (
	"slices"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// foldAbilities adds together the rows of one actor that share a spell
// id and a pet, keeping first-seen order.
func foldAbilities(rows []summary.Ability) []summary.Ability {
	type key struct {
		id  int64
		via string
	}
	out := make([]summary.Ability, 0, len(rows))
	at := map[key]int{}
	for _, r := range rows {
		k := key{r.SpellID, r.Via}
		i, ok := at[k]
		if !ok {
			at[k] = len(out)
			out = append(out, r)
			continue
		}
		dst := &out[i]
		dst.Total += r.Total
		dst.Effective += r.Effective
		dst.Resisted += r.Resisted
		dst.Blocked += r.Blocked
		dst.Hits += r.Hits
		dst.Crits += r.Crits
		dst.Ticks += r.Ticks
		for name, n := range r.Misses {
			if dst.Misses == nil {
				dst.Misses = map[string]int64{}
			}
			dst.Misses[name] += n
		}
	}
	return out
}

// foldCasts adds together the rows of one caster that share a spell id.
func foldCasts(rows []summary.CastRow) []summary.CastRow {
	type key struct {
		guid string
		id   int64
	}
	out := make([]summary.CastRow, 0, len(rows))
	at := map[key]int{}
	for _, r := range rows {
		k := key{r.GUID, r.SpellID}
		i, ok := at[k]
		if !ok {
			at[k] = len(out)
			out = append(out, r)
			continue
		}
		dst := &out[i]
		dst.Started += r.Started
		dst.Succeeded += r.Succeeded
		dst.Failed += r.Failed
	}
	return out
}

// foldAuras adds together the tracks on one target that share a spell
// id. Uptime and applications add; the appliers are a set.
func foldAuras(rows []summary.AuraTrack) []summary.AuraTrack {
	type key struct {
		target string
		id     int64
	}
	out := make([]summary.AuraTrack, 0, len(rows))
	at := map[key]int{}
	for _, r := range rows {
		k := key{r.TargetGUID, r.SpellID}
		i, ok := at[k]
		if !ok {
			at[k] = len(out)
			out = append(out, r)
			continue
		}
		dst := &out[i]
		dst.UptimeMS += r.UptimeMS
		dst.Applications += r.Applications
		if r.MaxStacks > dst.MaxStacks {
			dst.MaxStacks = r.MaxStacks
		}
		for _, name := range r.Appliers {
			if !slices.Contains(dst.Appliers, name) {
				dst.Appliers = append(dst.Appliers, name)
			}
		}
	}
	return out
}
