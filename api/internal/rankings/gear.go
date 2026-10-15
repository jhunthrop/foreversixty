package rankings

import (
	"context"
	"math"
)

// The character page's gear panel (Task L6.7): what the character wears in each of the 17
// export slots, the band's pick for that slot, and the score points the pick would add. The
// rows are computed by api/internal/rating (rating.Service.ScoreAndGear), where the Gear
// component's scale lives; this package carries them and owns the one sum, GearGapTotal.

// GearWorn is the item a character wears in a slot, from its newest addon export. Name and
// ItemLevel are null when the export's data build does not have the item.
type GearWorn struct {
	ID        int     `json:"id"`
	Name      *string `json:"name"`
	ItemLevel *int    `json:"item_level"`
}

// GearPick is the band's pick for a slot (sims.BandPicks). Chance is the drop fraction, null
// for a source with no stated rate (a promoted pick carries its own source's chance).
type GearPick struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Source string   `json:"source"`
	Chance *float64 `json:"chance"`
}

// GearSlot is one row of the gear panel. Gap is the score points the pick would add over
// the worn item, one decimal; null when either side cannot be scored.
type GearSlot struct {
	Slot string    `json:"slot"`
	Worn *GearWorn `json:"worn"`
	Pick *GearPick `json:"pick"`
	Gap  *float64  `json:"gap"`
}

// CharacterGear is the gear panel and its total. Slots is empty, never nil, when the
// character has no export.
type CharacterGear struct {
	Slots    []GearSlot
	GapTotal *float64
}

// ScoreGearReader reads a character's Forever Score and gear panel together by character
// key, so the two share one read of the stored export and one of the rated fights.
// rating.Service is the one implementation. The panel is always returned (empty when it
// cannot be read, which the implementation logs); the error is the score's, as Scorer's.
type ScoreGearReader interface {
	ScoreAndGear(ctx context.Context, key string) (ForeverScore, CharacterGear, error)
}

// GearGapTotal is gear_gap_total: the sum of slots' non-null gaps, one decimal. It is the
// points the band's picks would add over what the character wears, not the Gear component's
// distance to its top (that is measured against the band's best set across presets and
// factions). Null when no slot has a gap.
func GearGapTotal(slots []GearSlot) *float64 {
	var total float64
	found := false
	for _, s := range slots {
		if s.Gap != nil {
			total += *s.Gap
			found = true
		}
	}
	if !found {
		return nil
	}
	total = math.Round(total*10) / 10
	return &total
}
