package upgrades

import "math"

// EquippedItemLevel is the one item-level figure the guild roster, the readiness board and the
// bot's /char card all show (guild round 3: they read 61, 61 and 59.4 for one character): the
// mean item level of every equipped item, rounded half up to a whole level. levels holds one
// entry per equipped item (an empty slot is not equipped and has no entry; callers pass the
// gear slots only, never a shirt or tabard), nil for an item the data does not know. ok is
// false when nothing is equipped or any equipped item is unknown: the unknown items are the
// newest, so a mean of the rest would be wrong in exactly the case that matters.
func EquippedItemLevel(levels []*int) (int, bool) {
	if len(levels) == 0 {
		return 0, false
	}
	total := 0
	for _, l := range levels {
		if l == nil {
			return 0, false
		}
		total += *l
	}
	return int(math.Floor(float64(total)/float64(len(levels)) + 0.5)), true
}
