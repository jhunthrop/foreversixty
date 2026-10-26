// api/internal/guilds/raids_zones.go
//
// The curated raid zones a guild can schedule. The same values live in
// data/curated/raids/first-tier.json; raids_zones_test.go pins the two together. Curated
// data is not shipped in the API image, so this slice is the runtime source.
package guilds

// Zone is one schedulable raid zone. Size and LogZone are nil until a sourced fact names them.
type Zone struct {
	Key     string  `json:"key"`
	Name    string  `json:"name"`
	Size    *int    `json:"size"`
	LogZone *string `json:"log_zone"`
}

var (
	onyxiasLairSize    = 40
	onyxiasLairLogZone = "Onyxia's Lair"
)

var curatedZones = []Zone{
	{Key: "raid:barrow-deeps", Name: "Barrow Deeps"},
	{Key: "raid:hyjal-summit", Name: "Hyjal Summit"},
	{Key: "raid:onyxias-lair", Name: "Onyxia's Lair", Size: &onyxiasLairSize, LogZone: &onyxiasLairLogZone},
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
