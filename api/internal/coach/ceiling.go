package coach

import "strconv"

// ceilingDetail is F1's line when neither top logs nor the sim compares: the most casts the
// spell's own cooldown allows in the fight.
const ceilingDetail = "Spell data's ceiling for {Spell} is {ceiling} casts in this fight."

// CeilingCasts is spell data's own ceiling: floor(duration / cooldown) + 1, the first cast
// at the pull and one more each time the cooldown comes round. A spell with no cooldown has
// no ceiling (0), and so does a fight with no length.
func CeilingCasts(durationMS, cooldownMS int64) int {
	if durationMS <= 0 || cooldownMS <= 0 {
		return 0
	}
	return int(durationMS/cooldownMS) + 1
}

// withCeiling adds the ceiling line and its number to a finding that has no comparison.
func withCeiling(f *Finding, in Input, worst coreAbility) {
	ceiling := CeilingCasts(in.Fight.DurationMS, worst.spell.CooldownMS)
	if worst.cmp != nil || ceiling == 0 {
		return
	}
	f.Detail += " " + Fill(ceilingDetail, map[string]string{"Spell": worst.spell.Name, "ceiling": strconv.Itoa(ceiling)})
	f.Facts.Numbers = append(f.Facts.Numbers, castCount(int64(ceiling), worst.spell.Name))
}
