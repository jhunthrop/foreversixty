// Package zoneicons is the one table of the raid zones' HD icons (the home mock's set, the same
// names as web/src/lib/first-raids.ts), by zone key and by name: zones.json, which the data
// pipeline's loot run reads too (the dungeon and raid sources' icon in loot.json). Each is a
// name under /data/{build}/icons. A zone with no icon of its own reads as Generic.
package zoneicons

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

// Generic is the icon of a zone with none of its own; zones.json's generic says the same.
const Generic = "inv_misc_book_09"

//go:embed zones.json
var tableJSON []byte

type zone struct{ Key, Name, Icon string }

type table struct {
	Generic string `json:"generic"`
	Zones   []zone `json:"zones"`
}

// zones are the first raids, in zones.json's order.
var zones = mustLoad(tableJSON)

// mustLoad parses the table, failing at start-up on a broken one.
func mustLoad(raw []byte) []zone {
	var t table
	if err := json.Unmarshal(raw, &t); err != nil {
		panic(fmt.Sprintf("zoneicons: zones.json: %v", err))
	}
	if t.Generic != Generic || len(t.Zones) == 0 {
		panic("zoneicons: zones.json names another generic icon or no zones")
	}
	return t.Zones
}

// Names is the first raids' names, in that order.
func Names() []string {
	out := make([]string, len(zones))
	for i, z := range zones {
		out[i] = z.Name
	}
	return out
}

// For is the icon of the zone with this key ("raid:onyxias-lair"), Generic for any other.
func For(key string) string {
	for _, z := range zones {
		if z.Key == key {
			return z.Icon
		}
	}
	return Generic
}

// ForName is the icon of the zone with this name ("Onyxia's Lair"), Generic for any other.
func ForName(name string) string {
	for _, z := range zones {
		if z.Name == name {
			return z.Icon
		}
	}
	return Generic
}
