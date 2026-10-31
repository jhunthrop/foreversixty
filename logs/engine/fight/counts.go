// logs/engine/fight/counts.go
package fight

import (
	"sort"
	"strings"
	"time"

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
//   - A mob is a non-player unit that was hostile on some line of this fight
//     (so one that dies mind-controlled still counts), or a neutral one the
//     group damaged or killed: the yellow beasts of the levelling zones keep
//     their neutral flags while they fight. A neutral unit whose known maximum
//     health is at most Options.CritterMaxHP is a critter, never a mob.
//   - Every mob that died counts once, however many lines say so
//     (PARTY_KILL and UNIT_DIED for one kill are one death).
//   - A death is the group's when a PARTY_KILL names it, or when the last
//     damage the unit took came from a unit flagged MINE, PARTY or RAID
//     (which covers the group's pets).
//   - Every player UNIT_DIED counts. The one exception is a possible Feign
//     Death: a death within Options.FeignDeathWindow of that player's own
//     Feign Death cast is taken back if the player then casts or swings. Any
//     other death stands, whatever follows it: a corpse run, Reincarnation and
//     a Soulstone all log the dead player acting again with no SPELL_RESURRECT.
//   - The group is the players flagged MINE, PARTY or RAID on any line of
//     this fight, so someone who leaves the party mid-night is in the group
//     for the pulls before and out of it for the pulls after.
type Counter struct {
	critterMaxHP int64
	feignSpells  map[int64]bool
	feignWindow  time.Duration
	// lastFeign is when each player last cast Feign Death.
	lastFeign   map[string]time.Time
	players     map[string]bool
	group       map[string]bool
	died        map[string]bool
	groupKilled map[string]bool
	// lastHitByGroup is, for each NPC hit so far, whether the last damage it
	// took came from the group; groupHit whether the group hit it at all.
	lastHitByGroup map[string]bool
	groupHit       map[string]bool
	// hostileSeen is every NPC seen with hostile flags in this fight.
	hostileSeen map[string]bool
	// maxHP is each NPC's maximum health, from advanced blocks describing it.
	maxHP map[string]int64
	// pendingDeaths is each player whose latest death followed their own
	// Feign Death cast, and whether it was counted as the group's.
	pendingDeaths map[string]bool
	deaths        int
	groupDeaths   int
	// The counts a restored fight already held, which the sets above
	// cannot re-derive.
	baseMobDeaths, baseGroupKills int
}

// NewCounter returns an empty counter. Zero fields in o take their defaults.
func NewCounter(o Options) *Counter {
	o = o.withDefaults()
	feign := make(map[int64]bool, len(o.FeignDeathSpells))
	for _, id := range o.FeignDeathSpells {
		feign[id] = true
	}
	return &Counter{
		critterMaxHP: o.CritterMaxHP, feignSpells: feign, feignWindow: o.FeignDeathWindow,
		lastFeign: map[string]time.Time{},
		players:   map[string]bool{}, group: map[string]bool{},
		died: map[string]bool{}, groupKilled: map[string]bool{},
		lastHitByGroup: map[string]bool{}, groupHit: map[string]bool{},
		hostileSeen: map[string]bool{}, maxHP: map[string]int64{},
		pendingDeaths: map[string]bool{},
	}
}

// restoreCounter seeds a counter from a fight saved mid-way. The per-unit
// sets are not saved, so a death logged again after the restore counts a
// second time; the session never saves an open fight (it replays it), so
// only a caller restoring a segmenter by hand can see that.
func restoreCounter(o Options, f Fight) *Counter {
	c := NewCounter(o)
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
	c.observeUnits(e)
	switch e.Kind {
	case event.Damage:
		if isNPC(e.Dest) {
			byGroup := units.InGroup(e.Source.Flags)
			c.lastHitByGroup[e.Dest.GUID] = byGroup
			c.groupHit[e.Dest.GUID] = c.groupHit[e.Dest.GUID] || byGroup
		}
	case event.PartyKill:
		if isNPC(e.Dest) {
			c.groupHit[e.Dest.GUID] = true
			if c.isMob(e.Dest) {
				c.died[e.Dest.GUID] = true
				c.groupKilled[e.Dest.GUID] = true
			}
		}
	case event.CastSuccess:
		if c.feignSpells[e.Spell.ID] {
			c.lastFeign[e.Source.GUID] = e.Time
		}
	case event.Death:
		c.observeDeath(e)
	}
	if signOfLife(e) {
		c.undoDeath(e.Source.GUID)
	}
}

// observeUnits records the players and the group, and what the NPCs on
// this line are: hostile, and how much health they have.
func (c *Counter) observeUnits(e event.Event) {
	for _, u := range [...]event.Unit{e.Source, e.Dest} {
		switch {
		case u.GUID == "" || u.GUID == units.NoGUID:
		case units.Parse(u.GUID).Kind == units.KindPlayer:
			c.players[u.GUID] = true
			if units.InGroup(u.Flags) {
				c.group[u.GUID] = true
			}
		default:
			if units.Hostile(u.Flags) {
				c.hostileSeen[u.GUID] = true
			}
			if e.Adv.OK && e.Adv.MaxHP > 0 && e.Adv.InfoGUID == u.GUID {
				c.maxHP[u.GUID] = e.Adv.MaxHP
			}
		}
	}
}

func (c *Counter) observeDeath(e event.Event) {
	if units.Parse(e.Dest.GUID).Kind == units.KindPlayer {
		ours := units.InGroup(e.Dest.Flags) || c.group[e.Dest.GUID]
		c.deaths++
		if ours {
			c.groupDeaths++
		}
		if c.justFeigned(e.Dest.GUID, e.Time) {
			c.pendingDeaths[e.Dest.GUID] = ours
		}
		return
	}
	if notMobDeaths[e.Name] || !isNPC(e.Dest) || !c.isMob(e.Dest) {
		return
	}
	c.died[e.Dest.GUID] = true
	if c.lastHitByGroup[e.Dest.GUID] {
		c.groupKilled[e.Dest.GUID] = true
	}
}

// justFeigned is whether a death at this time follows the player's own
// Feign Death cast closely enough to be that feign.
func (c *Counter) justFeigned(guid string, at time.Time) bool {
	cast, ok := c.lastFeign[guid]
	return ok && !at.Before(cast) && at.Sub(cast) <= c.feignWindow
}

// undoDeath takes back a player's death that a sign of life has shown was
// no death.
func (c *Counter) undoDeath(guid string) {
	ours, pending := c.pendingDeaths[guid]
	if !pending {
		return
	}
	delete(c.pendingDeaths, guid)
	c.deaths--
	if ours {
		c.groupDeaths--
	}
}

// signOfLife is an event only a living player can be the source of: a cast
// or a melee swing. A damage-over-time tick or a pet's hit is not one, since
// both carry on after their owner dies.
func signOfLife(e event.Event) bool {
	if units.Parse(e.Source.GUID).Kind != units.KindPlayer {
		return false
	}
	switch e.Kind {
	case event.CastStart, event.CastSuccess:
		return true
	case event.Damage, event.Missed:
		return strings.HasPrefix(e.Name, "SWING_")
	}
	return false
}

// isMob decides at its death whether an NPC was a mob (see Counter).
func (c *Counter) isMob(u event.Unit) bool {
	if units.Hostile(u.Flags) || c.hostileSeen[u.GUID] {
		return true
	}
	if !units.Neutral(u.Flags) || !c.groupHit[u.GUID] {
		return false
	}
	hp, known := c.maxHP[u.GUID]
	return !known || hp > c.critterMaxHP
}

// isNPC is a unit with a GUID that is not a player.
func isNPC(u event.Unit) bool {
	return u.GUID != "" && u.GUID != units.NoGUID && units.Parse(u.GUID).Kind != units.KindPlayer
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
