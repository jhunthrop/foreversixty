package upgrades

import "github.com/jhunthrop/foreversixty/api/internal/bis"

// LevelBands are the BiS data's level bands.
var LevelBands = []int{20, 30, 40, 50, 60}

// TopBand is the last level band, the one the Phase 1 raid is ranked at.
var TopBand = LevelBands[len(LevelBands)-1]

// LevelBand is the band a level falls in (20 to 29 is 20); false under the first band.
func LevelBand(level *int) (int, bool) {
	if level == nil || *level < LevelBands[0] {
		return 0, false
	}
	return min(*level/10*10, TopBand), true
}

// PresetFor is the preset a character's upgrades read against at band: the Phase 1 raid
// (raid-ready) at 60, bare below (the site's default preset at 60 is raid-ready).
func PresetFor(band int) string {
	if band == TopBand {
		return bis.RaidPreset
	}
	return bis.BarePreset
}
