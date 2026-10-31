package upgrades

import (
	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// Items looks an item up in the build's table for the character's class.
type Items func(itemID int) (trees.Item, bool)

// Row is one slot whose worn item is not the pick.
type Row struct {
	Slot string
	Pick bis.Slot
	// WornID is the worn item, 0 for an empty slot; WornKnown is false when the item data
	// lacks it (always false for an empty slot).
	WornID    int
	WornKnown bool
	Verdict
	Group Group
}

// Best is a slot already wearing the pick, or a measured tie with it.
type Best struct {
	Slot     string
	ItemID   int
	ItemName string
}

// Comparison is a character's gear against one band (web/src/lib/home/upgrades.ts
// upgradesFor), its upgrades in the bot's order.
type Comparison struct {
	Upgrades    []Row
	AlreadyBest []Best
	// TotalGain is the sum of every upgrade's gain that has a figure.
	TotalGain float64
	// NotSimChecked counts the upgrades that are weapons with no listed alternative.
	NotSimChecked int
	Unit          Unit
}

// Compare is gear (slot -> worn item id) against band. A slot the band has no pick for is
// neither an upgrade nor already best.
func Compare(band bis.Band, gear map[string]int, items Items) Comparison {
	out := Comparison{Unit: UnitOf(band)}
	weights := band.WeightsByStat()
	score := func(id int) (float64, bool) {
		it, ok := items(id)
		if !ok {
			return 0, false
		}
		return ScoreItem(it, weights), true
	}
	picks := band.BySlot()
	var rows []Row
	for _, s := range Slots {
		pick, ok := picks[s.Key]
		if !ok || pick.ItemID <= 0 {
			continue
		}
		worn := gear[s.Key]
		v := Judge(band, pick, worn, score)
		if v.Tie {
			out.AlreadyBest = append(out.AlreadyBest, best(pick, worn, items))
			continue
		}
		_, known := items(worn)
		row := Row{Slot: s.Key, Pick: pick, WornID: worn, WornKnown: worn > 0 && known, Verdict: v, Group: groupOf(pick, worn, v)}
		if v.Gain != nil {
			out.TotalGain += *v.Gain
		}
		if v.NotSimChecked {
			out.NotSimChecked++
		}
		rows = append(rows, row)
	}
	out.Upgrades = Order(rows, func(r Row) (Group, float64) { return r.Group, gainOf(r.Verdict) })
	return out
}

// best is a tied slot's worn item, named from the item data, else the pick's name.
func best(pick bis.Slot, worn int, items Items) Best {
	name := pick.ItemName
	if it, ok := items(worn); ok && it.Name != "" {
		name = it.Name
	}
	return Best{Slot: pick.Slot, ItemID: worn, ItemName: name}
}

// groupOf is a row's place in the order: an empty slot, a measured gain, or a pick with no
// figure or an effect the sim leaves out.
func groupOf(pick bis.Slot, worn int, v Verdict) Group {
	switch {
	case worn <= 0:
		return GroupEmpty
	case v.Gain != nil && !pick.EffectUnmodelled:
		return GroupMeasured
	}
	return GroupUnmodelled
}

func gainOf(v Verdict) float64 {
	if v.Gain == nil {
		return 0
	}
	return *v.Gain
}
