package main

import "github.com/jhunthrop/foreversixty/sim/leveling"

import "sort"

// scored is one candidate with its score and source already resolved,
// the unit pick() ranks.
type scored struct {
	candidate
	Score     float64
	Source    itemSource // zero value when HasSource is false
	HasSource bool
}

// slotOrder is the order sim/api.GearSlots and the engine's own
// equipment array use. pick() walks it in this order because the
// ring/trinket and one-hand/two-hand rules below are stateful across
// slots: finger2 needs to know what finger1 chose, and off_hand needs
// to know whether main_hand chose a two-hander.
var slotOrder = []string{
	"head", "neck", "shoulder", "back", "chest", "wrist", "hands", "waist",
	"legs", "feet", "finger1", "finger2", "trinket1", "trinket2",
	"main_hand", "off_hand", "ranged",
}

// slotPick is one slot's answer: the best eligible item and the
// runner-up, either of which may be nil when the slot had zero or one
// candidate.
type slotPick struct {
	Item     *scored
	RunnerUp *scored
}

// candidatesBySlot fans a scored pool out by every planner slot each
// item occupies (a ring lands in both "finger1" and "finger2"; a
// one-hander in both "main_hand" and "off_hand"), sorted best score
// first within each slot. Ties break on item id, so the order - and
// therefore which of two equally-scored items is "the pick" versus
// "the runner-up" - is stable across runs.
func candidatesBySlot(pool []scored) map[string][]scored {
	out := map[string][]scored{}
	for _, s := range pool {
		for _, slot := range s.Slots {
			out[slot] = append(out[slot], s)
		}
	}
	for slot, list := range out {
		sort.SliceStable(list, func(i, j int) bool {
			if list[i].Score != list[j].Score {
				return list[i].Score > list[j].Score
			}
			return list[i].ID < list[j].ID
		})
		out[slot] = list
	}
	return out
}

// pick chooses one item per planner slot, greedily, best score first,
// with the two rules this lane's brief calls out from sim/bulk/expand.go
// (unexported there - valid() and the finger/trinket dedup in it, plus
// gearCombinations' one-hander-in-either-hand shape - so mirrored here
// rather than imported; see this function's own rule-by-rule comments
// for exactly what each one is):
//
//   - finger1/finger2 and trinket1/trinket2 are a PAIR: the same item
//     id, or two items sharing a name (a lower- and higher-quality
//     version of "the same ring"), cannot fill both slots of a pair.
//     bulk's valid() enforces the general form of this rule (any two
//     of a gear list's finger/trinket slots); pick only ever fills
//     two such slots per pair, so checking the second against the
//     first already chosen covers it.
//   - a two-handed main_hand leaves off_hand empty: bulk's valid()
//     refuses `mainHand.HandType == HandTwo && offHand.ID != 0` as a
//     GEAR LIST, which for a greedy picker translates to "do not pick
//     an off_hand item at all" once main_hand's pick is two-handed.
func pick(spec string, bySlot map[string][]scored) map[string]slotPick {
	out := make(map[string]slotPick, len(slotOrder))
	var fingerUsedID int
	var fingerUsedName string
	var fingerHasPick bool
	var trinketUsedID int
	var trinketUsedName string
	var trinketHasPick bool
	var mainHandTwoHand bool
	var mainHandID int
	var mainHandName string
	var mainHandHasPick bool

	for _, slot := range slotOrder {
		list := bySlot[slot]
		switch slot {
		case "finger2":
			if fingerHasPick {
				list = excludePaired(list, fingerUsedID, fingerUsedName)
			}
		case "trinket2":
			if trinketHasPick {
				list = excludePaired(list, trinketUsedID, trinketUsedName)
			}
		case "off_hand":
			if mainHandTwoHand {
				out[slot] = slotPick{}
				continue
			}
			// A one-hander already worn in the main hand is not
			// offered again for the off hand: this lane's prototype
			// assumes the character owns one copy of any BiS
			// recommendation, the same reasoning bulk/expand.go's
			// sameWeaponTwice comment gives for refusing one physical
			// one-hander in both hands within a single combination.
			// If it is genuinely the only one-hander eligible, that
			// exclusion empties the list and off_hand is correctly
			// left unpicked rather than silently dual-wielding a
			// second copy nobody said the character has.
			if mainHandHasPick {
				list = excludePaired(list, mainHandID, mainHandName)
			}
			if leveling.DualWieldSpecs[spec] {
				list = weaponsOnly(list)
			}
		}
		var sp slotPick
		if len(list) > 0 {
			item := list[0]
			sp.Item = &item
		}
		if len(list) > 1 {
			runnerUp := list[1]
			sp.RunnerUp = &runnerUp
		}
		out[slot] = sp

		switch slot {
		case "finger1":
			if sp.Item != nil {
				fingerUsedID, fingerUsedName, fingerHasPick = sp.Item.ID, sp.Item.Name, true
			}
		case "trinket1":
			if sp.Item != nil {
				trinketUsedID, trinketUsedName, trinketHasPick = sp.Item.ID, sp.Item.Name, true
			}
		case "main_hand":
			if sp.Item != nil {
				mainHandTwoHand = sp.Item.TwoHand
				mainHandID, mainHandName, mainHandHasPick = sp.Item.ID, sp.Item.Name, true
			}
		}
	}
	return out
}

// excludePaired drops an id or a name already used by this slot's
// pair-mate from a candidate list, preserving order.
func excludePaired(list []scored, usedID int, usedName string) []scored {
	out := make([]scored, 0, len(list))
	for _, s := range list {
		if s.ID == usedID || s.Name == usedName {
			continue
		}
		out = append(out, s)
	}
	return out
}

// weaponsOnly keeps the candidates that are weapons (the client's item
// class 2): a dual-wielder's off hand is never a held item or a shield.
func weaponsOnly(list []scored) []scored {
	out := make([]scored, 0, len(list))
	for _, sc := range list {
		if sc.ClassID == itemClassWeapon {
			out = append(out, sc)
		}
	}
	return out
}
