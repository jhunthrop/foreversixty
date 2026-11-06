package leveling

import "fmt"

// armorProficiencyLevels is the level a class's armor proficiency for an
// armor subclass (items.json subclass_id under class_id 4) opens, keyed by
// "<class>:<subclass id>". An absent pair is worn from level 1.
//
// Owner-stated rule (Justin, 2026-09-30, the single source of truth): shamans
// and hunters wear mail from level 40 (leather before); warriors and paladins
// wear plate from level 40 (mail before). Nothing else in Forever is
// level-gated.
var armorProficiencyLevels = map[string]int{
	"hunter:3":  40, // mail
	"shaman:3":  40, // mail
	"warrior:4": 40, // plate
	"paladin:4": 40, // plate
}

// ArmorProficiencyLevel is the level from which class wears armor of
// subclass; 0 when it is never gated.
func ArmorProficiencyLevel(class string, subclass int) int {
	return armorProficiencyLevels[fmt.Sprintf("%s:%d", class, subclass)]
}
