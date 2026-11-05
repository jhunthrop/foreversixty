// api/internal/guilds/raids_zones.go
//
// The curated zones a guild can schedule (design/specs/2026-11-05-guild-events.md §3.1):
// dungeon runs and guild events first (raids_zones_catalogue.go, generated with
// data/curated/raids/events.json from the client's Group Finder tables), then the raids of
// data/curated/raids/first-tier.json. raids_zones_test.go pins both files to these slices.
// Curated data is not shipped in the API image, so the slices are the runtime source.
package guilds

import (
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/phase"
)

// Zone kinds: what an officer posts (spec §3.1).
const (
	ZoneRaid    = "raid"
	ZoneDungeon = "dungeon"
	ZoneEvent   = "event"
)

// raidKeyPrefix is every raid zone's key prefix; the raid cap's count query filters on it.
const raidKeyPrefix = "raid:"

// KeyItem is the item a run's door asks for.
type KeyItem struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Quality int    `json:"quality"`
	// Door is the one door the key opens when the run has another way in with no lock
	// (Stratholme: Undead's service entrance); empty when the run needs the key.
	Door string `json:"door,omitempty"`
}

// Zone is one schedulable zone. Size, LogZone, the level range and Opens are nil until a
// sourced fact names them; a guild event's Size is the officer's pick of key.
type Zone struct {
	Key      string     `json:"key"`
	Name     string     `json:"name"`
	Kind     string     `json:"kind"`
	Size     *int       `json:"size"`
	LogZone  *string    `json:"log_zone"`
	Opens    *time.Time `json:"opens"`
	MinLevel *int       `json:"min_level"`
	MaxLevel *int       `json:"max_level"`
	KeyItem  *KeyItem   `json:"key_item"`
}

// IsRaid reports whether the zone is a raid, the only kind the 12-ahead cap and Readiness read.
func (z Zone) IsRaid() bool { return z.Kind == ZoneRaid }

// OpenedBy reports whether a run starting at startsAt, read on the guild's calendar in loc,
// falls on or after the zone's opening day. A zone with no Opens is always open.
func (z Zone) OpenedBy(startsAt time.Time, loc *time.Location) bool {
	if z.Opens == nil {
		return true
	}
	y, m, d := startsAt.In(loc).Date()
	oy, om, od := z.Opens.UTC().Date()
	return !time.Date(y, m, d, 0, 0, 0, 0, time.UTC).Before(time.Date(oy, om, od, 0, 0, 0, 0, time.UTC))
}

func intPtr(n int) *int { return &n }

var (
	onyxiasLairLogZone = "Onyxia's Lair"
	firstTierOpens     = phaseStart(phase.Raids1)
)

// raidZones are first-tier.json's raids, each opening with its phase.
var raidZones = []Zone{
	{Key: "raid:barrow-deeps", Name: "Barrow Deeps", Kind: ZoneRaid, Opens: firstTierOpens},
	{Key: "raid:hyjal-summit", Name: "Hyjal Summit", Kind: ZoneRaid, Opens: firstTierOpens},
	{
		Key: "raid:onyxias-lair", Name: "Onyxia's Lair", Kind: ZoneRaid, Size: intPtr(40),
		LogZone: &onyxiasLairLogZone, Opens: firstTierOpens,
	},
}

// curatedZones is the form's order: dungeons in level order, guild events, raids.
var curatedZones = append(append([]Zone{}, eventZones...), raidZones...)

// phaseStart is the instant the named phase opens; it panics on a name phase does not know,
// which raids_zones_test.go would catch first.
func phaseStart(name string) *time.Time {
	for _, b := range phase.Boundaries {
		if b.Name == name {
			t := b.Start
			return &t
		}
	}
	panic("guilds: unknown phase " + name)
}

// Zones returns the curated zones in their curated order.
func Zones() []Zone {
	out := make([]Zone, len(curatedZones))
	copy(out, curatedZones)
	return out
}

// ZoneByKey finds a curated zone by its key.
func ZoneByKey(key string) (Zone, bool) {
	for _, z := range curatedZones {
		if z.Key == key {
			return z, true
		}
	}
	return Zone{}, false
}
