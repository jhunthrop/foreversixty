package upgrades

import (
	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// Verdict is one slot's comparison of what is worn with the band's pick.
type Verdict struct {
	// Tie is true when the worn item is the pick or a listed alternative measured equal to it.
	Tie bool
	// Gain is the pick's advantage over the worn item in the band's unit; nil when no
	// honest figure exists.
	Gain *float64
	// NotSimChecked is true for a weapon slot whose worn item the pick does not list.
	NotSimChecked bool
}

// ItemScore is a worn item's stat-weight score under the band's weights; known is false when
// the item data lacks the item. A nil ItemScore leaves a worn item the pick does not list
// unmeasured.
type ItemScore func(itemID int) (score float64, known bool)

// Judge is one slot's verdict: wornID is the worn item, 0 when the slot is empty.
func Judge(band bis.Band, pick bis.Slot, wornID int, score ItemScore) Verdict {
	alt, listed := alternative(pick, wornID)
	switch {
	case wornID > 0 && (pick.ItemID == wornID || listed && alt.DpsDelta == 0):
		return Verdict{Tie: true}
	case listed:
		return Verdict{Gain: ptrTo(-alt.DpsDelta)}
	case weaponSlots[pick.Slot]:
		return Verdict{NotSimChecked: true}
	case pick.Score == nil || band.ReferenceDpsPerPoint == nil:
		return Verdict{}
	}
	worn := 0.0
	if wornID > 0 {
		var known bool
		if score == nil {
			return Verdict{}
		}
		if worn, known = score(wornID); !known {
			return Verdict{}
		}
	}
	return Verdict{Gain: ptrTo((*pick.Score - worn) * *band.ReferenceDpsPerPoint)}
}

// alternative is the pick's listed alternative for wornID; false when it lists none.
func alternative(pick bis.Slot, wornID int) (bis.Alternative, bool) {
	if wornID <= 0 {
		return bis.Alternative{}, false
	}
	for _, alt := range pick.Alternatives {
		if alt.ItemID == wornID {
			return alt, true
		}
	}
	return bis.Alternative{}, false
}

// contractStatNames are the item data's stat keys that the weights spell out in full
// (web/src/lib/addon/score.ts CONTRACT_STAT_NAMES).
var contractStatNames = map[string]string{
	"healing": "healing_power", "fire_res": "fire_resistance", "frost_res": "frost_resistance",
	"nature_res": "nature_resistance", "shadow_res": "shadow_resistance", "arcane_res": "arcane_resistance",
}

// ScoreItem is an item's stat-weight score (score.ts scoreItem): each stat times its weight,
// plus armor times the armor weight.
func ScoreItem(item trees.Item, weights map[string]float64) float64 {
	total := float64(item.Armor) * weights["armor"]
	for key, value := range item.Stats {
		name := key
		if full, ok := contractStatNames[key]; ok {
			name = full
		}
		total += float64(value) * weights[name]
	}
	return total
}

func ptrTo(v float64) *float64 { return &v }
