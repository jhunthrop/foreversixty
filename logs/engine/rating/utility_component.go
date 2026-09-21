// logs/engine/rating/utility_component.go
package rating

import (
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/utility"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// targetEnemy is utility.Entry.Target's value for a debuff applied to an
// enemy; anything else ("ally", "" or a cooldown/talent-modifier with no
// target) is treated as applied to allies.
const targetEnemy = "enemy"

// scoreUtility implements spec §1.3's Utility component: percentile-
// within-bracket of mean uptime-share (or use-share for a cooldown) across
// every buff/debuff/cooldown the spec's utility table lists as owned.
// talent-modifier entries are informational only and never scored (spec
// §3.2).
func scoreUtility(fight summary.Summary, player string, bracket Bracket, table *utility.Table, src PercentileSource) Component {
	bracket.Component = ComponentNameUtility
	c := Component{Name: ComponentNameUtility}
	if table == nil || len(table.Owned) == 0 {
		c.Excluded, c.Reason = true, ReasonNoUtilityTable
		return c
	}

	var shares []float64
	var moments []Moment
	for _, e := range table.Owned {
		switch e.Kind {
		case utility.Buff, utility.Debuff:
			share, tracked := auraUptimeShare(fight, player, e)
			shares = append(shares, share)
			if tracked && share < 100 {
				moments = append(moments, Moment{Kind: "utility_uptime", SpellID: e.SpellID, SpellName: e.Name})
			}
		case utility.Cooldown:
			shares = append(shares, cooldownUseShare(fight, player, e))
		default:
			// talent-modifier: informational only, never scored.
		}
	}
	if len(shares) == 0 {
		c.Excluded, c.Reason = true, ReasonNoUtilityTable
		return c
	}

	var sum float64
	for _, s := range shares {
		sum += s
	}
	mean := clamp(sum/float64(len(shares)), 0, 100)
	c = placeBoundedOrSelf(c, src, bracket, mean)
	c.Moments = moments
	return c
}

// auraUptimeShare is one owned entry's uptime share, 0-100, for the tracks
// this player is an applier of: the single best-covered enemy target for a
// debuff (a raid-wide multi-target fight still credits "kept it up" if any
// one target had it), or the mean across every ally recipient for a buff
// (a self/party buff applied to many raiders). tracked is false when the
// player never applied this spell at all this fight, which still
// contributes a legitimate 0, not a skip (spec §1.3's "mean rewards broad
// coverage").
func auraUptimeShare(fight summary.Summary, player string, e utility.Entry) (float64, bool) {
	if fight.DurationMS <= 0 {
		return 0, false
	}
	var tracks []summary.AuraTrack
	for _, tr := range fight.Auras {
		if tr.SpellID != e.SpellID || !containsGUID(tr.Appliers, player) {
			continue
		}
		tracks = append(tracks, tr)
	}
	if len(tracks) == 0 {
		return 0, false
	}
	if e.Target == targetEnemy {
		var best int64
		for _, tr := range tracks {
			if tr.UptimeMS > best {
				best = tr.UptimeMS
			}
		}
		return clamp(float64(best)/float64(fight.DurationMS)*100, 0, 100), true
	}
	var sum float64
	for _, tr := range tracks {
		sum += float64(tr.UptimeMS) / float64(fight.DurationMS) * 100
	}
	return clamp(sum/float64(len(tracks)), 0, 100), true
}

// cooldownUseShare is a coarse proxy for a utility cooldown's use (spec
// §1.3: "scored by cast count against fight length rather than by
// AuraTrack.UptimeMS"): the utility table schema (§3.2) does not curate an
// expected-uses-per-fight number for any cooldown entry, so a magic
// interval is not invented here -- "cast at least once" is the minimum bar
// a cooldown's use can assert without one, documented as a coarse proxy.
func cooldownUseShare(fight summary.Summary, player string, e utility.Entry) float64 {
	for _, cr := range fight.Casts {
		if cr.OwnerGUID == player && cr.SpellID == e.SpellID && cr.Succeeded > 0 {
			return 100
		}
	}
	return 0
}

func containsGUID(guids []string, guid string) bool {
	for _, g := range guids {
		if g == guid {
			return true
		}
	}
	return false
}
