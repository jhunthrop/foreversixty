package leveling

// EffectiveRequiredLevel is the level gate a gear candidate really has,
// shared by the leveling BiS pipeline (sim/cmd/leveling-bis's
// eligible.go/data.go) and the rotation ladder (sim/request/ladder.go's
// pickGearItem) so a level-10 character is never handed a level-60
// quest reward or crafted item either place.
//
// 2026-09-28 quest-levels finding: 848 of
// data/builds/<build>/loot.json's 1,140 quest-reward items (and every
// one of its 31 crafted items) carry client required_level 0 -- the
// item itself states no level gate because the QUEST (or, for a
// crafted item, the recipe) gates it instead. itemRequiredLevel is the
// item's own client value; floor is the caller's best estimate of that
// quest/recipe gate for this item (0 means "no floor known", which
// leaves itemRequiredLevel unchanged) -- see
// data.go/ladder.go's own callers for how floor is built:
//
//   - a quest reward's floor is the LOWEST min_level among the quests
//     that award it (data/pipeline/wowhead_quests.py; loot.json's
//     quests map) -- any one of them suffices to obtain the item, so
//     the easiest one is the real gate.
//   - a crafted item's floor is the recipe's skill level when
//     loot.json states one (it does not today -- the engine fork's own
//     database carries no skill-level field on a crafted source), else
//     wowhead_quests.py's same item_level_proxy formula, computed
//     independently in Go by ItemLevelProxyRequiredLevel below (both
//     land on min(60, item_level-5); they cannot share code across
//     languages, only the formula).
//
// The floor only ever RAISES the gate, never lowers itemRequiredLevel:
// a floor of 0 (nothing this candidate's kind of source ever floors, or
// no quest/recipe data at all) is a no-op, and a floor lower than the
// item's own stated required_level never overrides a HIGHER client
// value.
func EffectiveRequiredLevel(itemRequiredLevel, floor int) int {
	if floor > itemRequiredLevel {
		return floor
	}
	return itemRequiredLevel
}

// LowestFloor is the smallest positive level among floors, or 0 if
// floors is empty or every entry is <= 0 -- "any one quest suffices"
// reduced to one call, for a caller building EffectiveRequiredLevel's
// floor argument from several QuestSource entries on the same item.
// Zero and negative entries are ignored rather than winning as the
// lowest: a floor of 0 means "no gate", not "obtainable at level 0",
// and letting one propagate would erase every other quest's real floor.
func LowestFloor(floors []int) int {
	lowest := 0
	for _, f := range floors {
		if f <= 0 {
			continue
		}
		if lowest == 0 || f < lowest {
			lowest = f
		}
	}
	return lowest
}

// ItemLevelProxyRequiredLevel mirrors
// data/pipeline/wowhead_quests.py's item_level_proxy formula for a
// crafted item Go itself must estimate a level floor for (loot.json's
// crafted source carries no recipe skill level today): min(60,
// item_level-5), floored at 0. Kept as its own named function rather
// than inlined at each call site so the two Go callers (leveling-bis's
// data.go, the ladder's ladder.go) and the Python formula stay
// visibly, deliberately the same number.
func ItemLevelProxyRequiredLevel(itemLevel int) int {
	proxy := itemLevel - 5
	if proxy > 60 {
		return 60
	}
	if proxy < 0 {
		return 0
	}
	return proxy
}

// QuestRewardLevelSlack is how far below a quest's own level a character
// is still assumed able to finish it (a "yellow" quest): a level-30 quest
// is realistic at 27, not at the level-20 the quest lets you accept it
// at. Owner ruling, 2026-09-29 (tightened from 5 the same day): a quest
// reward more than this many levels above the character is not a
// leveling-BiS pick -- a level-30 quest reward never heads a level-20 list.
const QuestRewardLevelSlack = 3

// QuestFloor is the level a character must reach before a quest's
// reward counts as obtainable: the quest's own accept level, raised to
// its quest level minus QuestRewardLevelSlack when that level is known
// (0 = unknown, then the accept level alone gates).
func QuestFloor(minLevel, questLevel int) int {
	if questLevel <= 0 {
		return minLevel
	}
	if by := questLevel - QuestRewardLevelSlack; by > minLevel {
		return by
	}
	return minLevel
}
