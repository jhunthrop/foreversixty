// Package zoneicons is the one table of the raid zones' HD icons (the home mock's set, the same
// names as web/src/lib/first-raids.ts), by zone key and by name. Each is a name under
// /data/{build}/icons. A zone with no icon of its own reads as Generic.
package zoneicons

// Generic is the icon of a zone with none of its own.
const Generic = "inv_misc_book_09"

type zone struct{ key, name, icon string }

var zones = []zone{
	{"raid:barrow-deeps", "Barrow Deeps", "spell_shadow_raisedead"},
	{"raid:hyjal-summit", "Hyjal Summit", "spell_nature_timestop"},
	{"raid:onyxias-lair", "Onyxia's Lair", "inv_misc_head_dragon_black"},
}

// For is the icon of the zone with this key ("raid:onyxias-lair"), Generic for any other.
func For(key string) string {
	for _, z := range zones {
		if z.key == key {
			return z.icon
		}
	}
	return Generic
}

// ForName is the icon of the zone with this name ("Onyxia's Lair"), Generic for any other.
func ForName(name string) string {
	for _, z := range zones {
		if z.name == name {
			return z.icon
		}
	}
	return Generic
}
