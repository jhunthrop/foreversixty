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
	// Ties is every OTHER candidate in this slot's score-sorted list
	// that scored identically to Item (this lane's brief, defect 4):
	// pick()'s tie-break is item id, ascending, which otherwise
	// silently picks one of several equally-good items - three
	// one-handers scoring 8.28 apiece, with the report showing only
	// the lowest-id one as though it were uniquely best - with no
	// record that the others existed. Populated by pick() only: the
	// trinket/effect/set-completion passes (rank.go, trinkets.go,
	// sets.go) replace a slot's whole slotPick with their own
	// real-sim-ranked choice, which is not a score tie in this sense,
	// so a slot any of them overwrites reports no ties.
	Ties []scored
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
		case "main_hand":
			// score() (see its own comment) converts a weapon's DPS to
			// attack power per slot in isolation, with no term for the
			// off hand a two-hander forfeits - so a two-hander's higher
			// raw dps routinely outscores a SINGLE one-hander here even
			// though a dual-wielder loses an entire second weapon's
			// worth of attack power, and for a spec whose kit assumes
			// two imbued weapons (Enhancement's Windfury/Rockbiter,
			// leveling.KitConsumes) also loses the off-hand imbue
			// entirely. This kept picking Smite's Mighty Hammer (a
			// two-hand hammer, item 7230) for shaman-enhancement's
			// level-20 main hand, leaving off_hand permanently empty
			// rather than the dual-wield set a real Enhancement shaman
			// runs.
			//
			// A blanket exclusion is wrong for a spec whose weapon is a
			// stat stick rather than its damage source, though (this
			// lane's brief, defect 3): hunter-beast-mastery/-marksmanship
			// are in DualWieldSpecs (pick.go treats "never offers a
			// two-hander for main_hand" and "offers main_hand's
			// one-handers to off_hand" as one rule), but their weight
			// runs zero out melee attack_power entirely (character.go's
			// weightsRequest doc), so score()'s dps-derived term is not
			// inflating a two-hander's score the way it does for a
			// melee dual-wielder - Impaling Harpoon (a two-hand polearm,
			// scored on its flat agility alone) legitimately outscored
			// Goblin Screwdriver+Poniard (two one-handers) this way.
			// twoHandBeatsPair compares the best two-hander against the
			// best LEGAL pair (main one-hander + its own best off-hand
			// partner, not the two-hander's score against a single
			// one-hander alone - the exact "no term for the forfeited
			// off hand" gap the exclusion above exists to guard
			// against), so a two-hander only wins here when it is
			// ahead of the full pair it would replace. Ties, and any
			// case this score-level approximation gets wrong, are
			// settled by the real sim in verify.go's own swap pass,
			// same as every other pick() decision.
			if leveling.DualWieldSpecs[spec] && !twoHandBeatsPair(bySlot) {
				list = excludeTwoHand(list)
			}
		case "off_hand":
			if mainHandTwoHand {
				out[slot] = slotPick{}
				continue
			}
			// A dual-wielder's off hand is a weapon (DualWieldSpecs'
			// own doc), but a one-hander's raw item row carries only
			// "main_hand" as its .Slot - data.go's plannerSlots fans
			// finger/trinket into their numbered pair, never a
			// one-hander into its second equippable hand - so
			// bySlot["off_hand"] never held a weapon at all; every
			// dual-wield spec's off hand came back unpicked at every
			// band (audit-rogue, 2026-09-28: assassination, combat and
			// subtlety all showed an empty off_hand at 20, 30 and 40).
			// Fixed here, not in plannerSlots, because only a
			// DualWieldSpecs member ever wants a weapon in both hands
			// and plannerSlots has no spec to check against.
			if leveling.DualWieldSpecs[spec] {
				list = mergeByScore(list, oneHandedWeapons(bySlot["main_hand"]))
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
		if sp.Item != nil {
			// list is best-score-first (candidatesBySlot's own
			// contract), so every tie with the winner is contiguous
			// starting right after index 0 - stop at the first lower
			// score.
			for i := 1; i < len(list); i++ {
				if list[i].Score != sp.Item.Score {
					break
				}
				sp.Ties = append(sp.Ties, list[i])
			}
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

// oneHandedWeapons keeps the candidates from a "main_hand" list that
// are weapons AND not two-handed - the pool the off_hand case above
// borrows from for a dual-wield spec. A two-hander is excluded here
// even though the mainHandTwoHand check above already stops this case
// once the CHOSEN main-hand item is two-handed: this list holds every
// eligible main_hand candidate, chosen or not, and a two-hander among
// the others offered no off-hand slot in the real game either.
func oneHandedWeapons(list []scored) []scored {
	out := make([]scored, 0, len(list))
	for _, sc := range list {
		if sc.ClassID == itemClassWeapon && !sc.TwoHand {
			out = append(out, sc)
		}
	}
	return out
}

// excludeTwoHand drops two-handed weapons from a dual-wielder's main-
// hand candidates: see pick()'s own "main_hand" case for why score()'s
// per-slot heuristic cannot be trusted to make this call by itself.
func excludeTwoHand(list []scored) []scored {
	out := make([]scored, 0, len(list))
	for _, sc := range list {
		if !sc.TwoHand {
			out = append(out, sc)
		}
	}
	return out
}

// twoHandBeatsPair reports whether bySlot's best two-handed main_hand
// candidate outscores the best legal one-handed main_hand+off_hand
// PAIR - this lane's brief, defect 3. mainList is already best-score-
// first (candidatesBySlot's own contract), so the first two-handed
// entry found walking it is the single highest-scoring two-hander
// regardless of how many one-handers sit above it in the list.
//
// The pair side mirrors pick()'s own off_hand case: the best one-hand
// main_hand candidate, plus the best off_hand candidate once that
// one-hander is excluded from its own pool (excludePaired) - a
// dual-wielder's off hand pool is main_hand's one-handers merged with
// off_hand's own list, weapons only, exactly as pick() builds it. A
// pair with no viable off-hand partner (empty pool) is still scored on
// its main-hand item alone, since a bare main-hander is what a
// dual-wielder gets when nothing else is eligible.
func twoHandBeatsPair(bySlot map[string][]scored) bool {
	mainList := bySlot["main_hand"]
	var bestTwoHand *scored
	for i := range mainList {
		if mainList[i].TwoHand {
			c := mainList[i]
			bestTwoHand = &c
			break
		}
	}
	if bestTwoHand == nil {
		return false
	}

	oneHanders := oneHandedWeapons(mainList)
	if len(oneHanders) == 0 {
		// No legal one-hand main_hand candidate at all: the two-hander
		// has nothing to lose to.
		return true
	}
	bestMain := oneHanders[0]
	pairScore := bestMain.Score

	offPool := weaponsOnly(excludePaired(mergeByScore(bySlot["off_hand"], oneHanders), bestMain.ID, bestMain.Name))
	if len(offPool) > 0 {
		pairScore += offPool[0].Score
	}

	return bestTwoHand.Score > pairScore
}

// mergeByScore concatenates two already best-score-first lists (the
// same ordering candidatesBySlot gives every bySlot[slot] entry) and
// re-sorts the result, since interleaving two independently sorted
// lists is not itself sorted. Ties break on item id, matching
// candidatesBySlot's own tiebreak, so which of two equally-scored
// items is "the pick" stays stable across runs.
func mergeByScore(a, b []scored) []scored {
	out := make([]scored, 0, len(a)+len(b))
	out = append(out, a...)
	out = append(out, b...)
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Score != out[j].Score {
			return out[i].Score > out[j].Score
		}
		return out[i].ID < out[j].ID
	})
	return out
}

// enforceTwoHandOffHandInvariant returns picks unchanged, or a copy
// with off_hand cleared, so that "a two-handed main_hand means an
// empty off_hand" - pick()'s own rule, right above - still holds after
// every LATER pass that can replace main_hand's pick without knowing
// off_hand exists: rankSlotWithEffects (an implemented-proc two-hander
// can out-measure the one-handed pick) and trySetCompletion (a
// two-hand set piece can complete a set the independently-scored
// picks did not). Both run once per slot, blind to what any OTHER
// slot's own pass just decided, so neither can maintain this
// invariant on its own - main.go's runSpec calls this once, after
// every pass, right before verifyBand/buildReport read picks for
// good.
//
// This is not limited to leveling.DualWieldSpecs: this lane's own
// dogfood run found it on priest-shadow and every warlock spec too (a
// two-hand staff's implemented proc winning main_hand over the
// original one-hand pick, with a held off-hand item pick()
// legitimately gave them for THAT original pick left standing) - any
// spec whose main_hand pool can contain both one- and two-handers is
// exposed, not only the ones that dual-wield.
func enforceTwoHandOffHandInvariant(picks map[string]slotPick) map[string]slotPick {
	mh := picks["main_hand"]
	if mh.Item == nil || !mh.Item.TwoHand || picks["off_hand"].Item == nil {
		return picks
	}
	out := make(map[string]slotPick, len(picks))
	for k, v := range picks {
		out[k] = v
	}
	out["off_hand"] = slotPick{}
	return out
}
