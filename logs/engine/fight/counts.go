// logs/engine/fight/counts.go
package fight

import (
	"sort"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// notMobDeaths are the death lines that end a unit without a mob dying: a
// totem destroyed, a summon dissipating.
var notMobDeaths = map[string]bool{"UNIT_DESTROYED": true, "UNIT_DISSIPATES": true}

// Counter tallies who took part in a fight and who died in it. The
// segmenter keeps one per fight; the API's rebuild of an uploaded fight
// uses one directly, so both count the same events the same way.
//
//   - Every hostile NPC that died counts once, however many lines say so
//     (PARTY_KILL and UNIT_DIED for one kill are one death).
//   - A death is the group's when a PARTY_KILL names it, or when the last
//     damage the unit took came from a unit flagged MINE, PARTY or RAID
//     (which covers the group's pets).
//   - The group is the players flagged MINE, PARTY or RAID on any line of
//     this fight, so someone who leaves the party mid-night is in the group
//     for the pulls before and out of it for the pulls after.
type Counter struct {
	players     map[string]bool
	group       map[string]bool
	died        map[string]bool
	groupKilled map[string]bool
	// lastHitByGroup is, for each hostile unit hit so far, whether the last
	// damage it took came from the group.
	lastHitByGroup map[string]bool
	deaths         int
	groupDeaths    int
	// The counts a restored fight already held, which the sets above
	// cannot re-derive.
	baseMobDeaths, baseGroupKills int
}

// NewCounter returns an empty counter.
func NewCounter() *Counter {
	return &Counter{
		players: map[string]bool{}, group: map[string]bool{},
		died: map[string]bool{}, groupKilled: map[string]bool{},
		lastHitByGroup: map[string]bool{},
	}
}

// restoreCounter seeds a counter from a fight saved mid-way. The per-unit
// death sets are not saved, so a death logged again after the restore
// counts a second time; the session never saves an open fight (it replays
// it), so only a caller restoring a segmenter by hand can see that.
func restoreCounter(f Fight) *Counter {
	c := NewCounter()
	for _, g := range f.Players {
		c.players[g] = true
	}
	for _, g := range f.Group {
		c.group[g] = true
	}
	c.deaths, c.groupDeaths = f.Deaths, f.GroupDeaths
	c.baseMobDeaths, c.baseGroupKills = f.MobDeaths, f.GroupKills
	return c
}

// Observe counts one event.
func (c *Counter) Observe(e event.Event) {
	for _, u := range [...]event.Unit{e.Source, e.Dest} {
		if u.GUID == "" || u.GUID == units.NoGUID || units.Parse(u.GUID).Kind != units.KindPlayer {
			continue
		}
		c.players[u.GUID] = true
		if units.InGroup(u.Flags) {
			c.group[u.GUID] = true
		}
	}
	switch e.Kind {
	case event.Damage:
		if units.Hostile(e.Dest.Flags) {
			c.lastHitByGroup[e.Dest.GUID] = units.InGroup(e.Source.Flags)
		}
	case event.PartyKill:
		if isMob(e.Dest) {
			c.died[e.Dest.GUID] = true
			c.groupKilled[e.Dest.GUID] = true
		}
	case event.Death:
		c.observeDeath(e)
	}
}

func (c *Counter) observeDeath(e event.Event) {
	if units.Parse(e.Dest.GUID).Kind == units.KindPlayer {
		c.deaths++
		if units.InGroup(e.Dest.Flags) || c.group[e.Dest.GUID] {
			c.groupDeaths++
		}
		return
	}
	if notMobDeaths[e.Name] || !isMob(e.Dest) {
		return
	}
	c.died[e.Dest.GUID] = true
	if c.lastHitByGroup[e.Dest.GUID] {
		c.groupKilled[e.Dest.GUID] = true
	}
}

// isMob is a hostile unit that is not a player.
func isMob(u event.Unit) bool {
	return u.GUID != "" && u.GUID != units.NoGUID && units.Hostile(u.Flags) &&
		units.Parse(u.GUID).Kind != units.KindPlayer
}

// Apply writes the counts onto f, replacing what f held.
func (c *Counter) Apply(f *Fight) {
	f.Players, f.Group = sortedKeys(c.players), sortedKeys(c.group)
	f.Deaths, f.GroupDeaths = c.deaths, c.groupDeaths
	f.MobDeaths, f.GroupKills = c.baseMobDeaths+len(c.died), c.baseGroupKills+len(c.groupKilled)
}

func sortedKeys(m map[string]bool) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
