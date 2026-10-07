package main

// The tank-only steps of a band: the weights character a tank is swept
// on, and the published figures. Everything else a tank band does is the
// ordinary band pipeline running on tankEngine (score_tank.go).

import (
	"fmt"
	"sort"

	"github.com/jhunthrop/foreversixty/sim/api"
)

const (
	// armorWeightStat is the weight_stats id of a gear piece's own armor.
	armorWeightStat = "armor"

	// tankBaselineMinQuality is the lowest item quality the weights
	// character wears: green. Grey and white pieces are not what a tank
	// of that level is in.
	tankBaselineMinQuality = 2
	// tankBaselineItemLevelWindow is how far below the slot's best item
	// level a piece may be and still count as one the character could be
	// wearing; within it the most armored piece wins.
	tankBaselineItemLevelWindow = 8
)

// tankBaselineSlots are the slots the weights character is dressed in.
// Trinkets and the ranged slot are left empty: their value is in effects
// the weights should not be taken around, and the ranker decides them in
// its own tournaments.
var tankBaselineSlots = []string{
	"head", "neck", "shoulder", "back", "chest", "wrist", "hands", "waist",
	"legs", "feet", "finger1", "finger2", "off_hand",
}

// tankLadderCharacter dresses a tank's weights character. A DPS spec's
// weights are measured on a bare, armed character; a tank's would be
// measured at an armor, health and avoidance no real tank has, where a
// point of armor is worth far more than it is at the gear the set it
// feeds will have. So a tank is swept on a representative set instead:
// in every slot, among the pieces its level can wear that are within a
// few item levels of the best, the most armored (then the sturdiest). The
// weapon the ladder already chose stays. An off hand with no armor (a
// held item, which a druid wears) is left empty: only a shield counts.
func tankLadderCharacter(spec specInfo, items []candidate, level int, ch api.CharacterSpec) api.CharacterSpec {
	if spec.Role != roleTank {
		return ch
	}
	worn := make(map[string]bool, len(ch.Gear))
	taken := make(map[int]bool, len(ch.Gear))
	for _, g := range ch.Gear {
		worn[g.Slot] = true
		taken[g.ItemID] = true
	}
	gear := append([]api.GearSlot(nil), ch.Gear...)
	for _, slot := range tankBaselineSlots {
		if worn[slot] {
			continue
		}
		if pick := tankBaselinePick(items, slot, level, taken); pick != nil {
			taken[pick.ID] = true
			gear = append(gear, api.GearSlot{Slot: slot, ItemID: pick.ID})
		}
	}
	ch.Gear = gear
	return ch
}

// tankBaselinePick is the representative piece for one slot, or nil.
func tankBaselinePick(items []candidate, slot string, level int, taken map[int]bool) *candidate {
	var pool []candidate
	best := 0
	for _, c := range items {
		if taken[c.ID] || c.RequiredLevel > level || c.Quality < tankBaselineMinQuality || !offersSlot(c, slot) {
			continue
		}
		pool = append(pool, c)
		best = max(best, c.ItemLevel)
	}
	var near []candidate
	for _, c := range pool {
		if c.ItemLevel >= best-tankBaselineItemLevelWindow {
			near = append(near, c)
		}
	}
	if len(near) == 0 {
		return nil
	}
	sort.Slice(near, func(i, j int) bool {
		a, b := near[i], near[j]
		if a.Armor != b.Armor {
			return a.Armor > b.Armor
		}
		if a.Stats["stamina"] != b.Stats["stamina"] {
			return a.Stats["stamina"] > b.Stats["stamina"]
		}
		return a.ID < b.ID
	})
	if slot == "off_hand" && near[0].Armor == 0 {
		return nil
	}
	return &near[0]
}

func offersSlot(c candidate, slot string) bool {
	for _, s := range c.Slots {
		if s == slot {
			return true
		}
	}
	return false
}

// annotateTankBand publishes a tank band's figures: one longer run of the
// band's final set, whose damage taken, TMI, chance of death, threat and
// own damage replace a DPS number as the band's headline. SetDPS becomes
// the tank's own damage so nothing that reads it breaks, and the score
// unit says every sim-decided figure on the band is tank score.
func annotateTankBand(runner engineRunner, spec specInfo, race string, level int, talents string, picks map[string]slotPick, rep *bandReport) error {
	if spec.Role != roleTank {
		return nil
	}
	tank, ok := runner.(tankRunner)
	if !ok {
		return fmt.Errorf("%s is a tank but its runner is not a tank runner", spec.Spec)
	}
	req := plainRequest(spec, bandCharacter("metrics", race, spec.ClassSlug, spec.Spec, level, talents, buildGear(picks)), tankMetricsIterations, tankMetricsSeed)
	figures, err := tank.RunTankFigures(req)
	if err != nil {
		return err
	}
	rep.Role = roleTank
	rep.Metrics = figures.report()
	rep.SetDPS = figures.DPS
	rep.ScoreUnit = scoreUnitTankScore
	return nil
}
