package request

import "fmt"

// The two factions, as the BiS files and spec-breakdown's -faction name them.
const (
	AllianceFaction = "alliance"
	HordeFaction    = "horde"
)

// raceFactions is the faction of every race slug in races. Skyborne is one
// race whose faction is chosen at creation; the client carries it as two rows,
// and so do races and this map.
var raceFactions = map[string]string{
	"dwarf":               AllianceFaction,
	"gnome":               AllianceFaction,
	"human":               AllianceFaction,
	"night-elf":           AllianceFaction,
	"high-order-skyborne": AllianceFaction,
	"orc":                 HordeFaction,
	"tauren":              HordeFaction,
	"troll":               HordeFaction,
	"undead":              HordeFaction,
	"windshaper-skyborne": HordeFaction,
}

// FactionOfRace is the faction a race slug plays for; the rank choices a
// character casts follow it.
func FactionOfRace(race string) (string, error) {
	faction, ok := raceFactions[race]
	if !ok {
		return "", fmt.Errorf("request: race %q has no faction", race)
	}
	return faction, nil
}
