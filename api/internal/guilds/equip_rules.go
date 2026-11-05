// api/internal/guilds/equip_rules.go
//
// Who can take a drop (guild loot spec 2026-11-05 §3.2): the item's own class restriction, the
// class's weapon and armour skills at the character's level, Dual Wield for a one-hander in the
// off hand, and, for the item-level tier only, whether the spec wants the item's stats. The skills
// table is generated from the client (equip_rules_table.go, from data/builds/<build>/
// equip-skills.json); the stat families are a rule, kept here, because data/curated is not in the
// API image.
package guilds

import "strings"

//go:generate go run ../../cmd/genequiprules -build 1.60.1.70790

// equipClass is one class's equip skills: item subclass -> the level it is learned at, for armour
// and for weapons, and the level Dual Wield is learned at (0: never; Forever's shamans).
type equipClass struct {
	armor, weapons map[int]int
	dualWieldLevel int
}

// Item.ClassID values a drop can carry.
const (
	itemClassWeapon = 2
	itemClassArmor  = 4
)

// armorMiscellaneous is Item.SubclassID 0 under armour: rings, necks, trinkets and items held in
// the off hand. No skill gates it.
const armorMiscellaneous = 0

// Item.InventoryType values the loot rules read.
const (
	invOneHand       = 13
	invHeldInOffHand = 23
)

// classMaskBit is a class slug's bit in ItemSparse.AllowableClass (1 << (ChrClasses.ID - 1)).
var classMaskBit = map[string]int{
	"warrior": 1, "paladin": 2, "hunter": 4, "rogue": 8, "priest": 16,
	"shaman": 64, "mage": 128, "warlock": 256, "druid": 1024,
}

// allowsClass reports whether an AllowableClass mask lets class take the item. -1, 0 and a mask
// naming every class all mean no restriction.
func allowsClass(mask int, class string) bool {
	if mask <= 0 {
		return true
	}
	bit, ok := classMaskBit[class]
	return ok && mask&bit != 0
}

// canEquip reports whether a character of class at level can equip item: the item's class
// restriction, then the armour or weapon skill for its subclass learned by level. Miscellaneous
// armour (a ring, a trinket, a held-in-off-hand item) needs no skill.
func canEquip(class string, level int, item lootItem) bool {
	if !allowsClass(item.AllowableClass, class) {
		return false
	}
	skills, ok := equipSkills[class]
	if !ok {
		return false
	}
	var learnedAt map[int]int
	switch item.ItemClass {
	case itemClassArmor:
		if item.Subclass == armorMiscellaneous {
			return true
		}
		learnedAt = skills.armor
	case itemClassWeapon:
		learnedAt = skills.weapons
	default:
		return false
	}
	at, ok := learnedAt[item.Subclass]
	return ok && level >= at
}

// canDualWield reports whether a character of class at level has learned Dual Wield.
func canDualWield(class string, level int) bool {
	at := equipSkills[class].dualWieldLevel
	return at > 0 && level >= at
}

// statFamily is what a spec wants from an item's stats (spec §3.2 rule 4).
type statFamily int

const (
	familyPhysical statFamily = iota
	familyCaster
	familyHealer
	familyTank
)

// familyStats is each family's stats, keyed as items/<class>.json's stats are. Hit and crit are
// in none of them: every family values both (sharedStats).
var familyStats = map[statFamily][]string{
	familyPhysical: {"strength", "agility", "attack_power", "ranged_attack_power", "feral_attack_power"},
	familyCaster: {"intellect", "spell_power", "arcane_power", "fire_power", "frost_power",
		"holy_power", "nature_power", "shadow_power", "spell_penetration"},
	familyHealer: {"intellect", "spirit", "healing", "mp5"},
	familyTank: {"strength", "agility", "attack_power", "defense", "dodge", "parry", "block",
		"block_value"},
}

// sharedStats are the stats every family values.
var sharedStats = []string{"hit", "crit"}

// specFamilies maps "class/spec" (lower case, the spec as fight_metrics names it) to its family.
var specFamilies = map[string]statFamily{
	"warrior/arms": familyPhysical, "warrior/fury": familyPhysical, "warrior/protection": familyTank,
	"rogue/assassination": familyPhysical, "rogue/combat": familyPhysical, "rogue/subtlety": familyPhysical,
	"hunter/beast mastery": familyPhysical, "hunter/marksmanship": familyPhysical, "hunter/survival": familyPhysical,
	"druid/feral": familyPhysical, "druid/balance": familyCaster, "druid/restoration": familyHealer,
	"shaman/enhancement": familyPhysical, "shaman/elemental": familyCaster, "shaman/restoration": familyHealer,
	"paladin/retribution": familyPhysical, "paladin/holy": familyHealer, "paladin/protection": familyTank,
	"priest/holy": familyHealer, "priest/discipline": familyHealer, "priest/shadow": familyCaster,
	"mage/arcane": familyCaster, "mage/fire": familyCaster, "mage/frost": familyCaster,
	"warlock/affliction": familyCaster, "warlock/demonology": familyCaster, "warlock/destruction": familyCaster,
}

// classFamilies is every family a class can play, for a character whose spec is not known yet.
var classFamilies = map[string][]statFamily{
	"warrior": {familyPhysical, familyTank},
	"rogue":   {familyPhysical},
	"hunter":  {familyPhysical},
	"druid":   {familyPhysical, familyTank, familyCaster, familyHealer},
	"shaman":  {familyPhysical, familyCaster, familyHealer},
	"paladin": {familyPhysical, familyTank, familyHealer},
	"priest":  {familyHealer, familyCaster},
	"mage":    {familyCaster},
	"warlock": {familyCaster},
}

// familiesFor is the family a character plays: its spec's, a tank role reading as tank (a bear
// druid logs as Feral), else every family its class can play when the spec is unknown.
func familiesFor(class, spec, role string) []statFamily {
	if role == "tank" {
		return []statFamily{familyTank}
	}
	if f, ok := specFamilies[class+"/"+strings.ToLower(spec)]; ok {
		return []statFamily{f}
	}
	return classFamilies[class]
}

// suitsFamily reports whether an item with stats suits one of families: it carries a stat of the
// family or hit or crit, or carries no family stat at all (Stamina, armour and resistances only).
// A weapon with no family stat (Deathbringer, Vis'kag the Bloodletter: a chance on hit and nothing
// else) is a physical weapon: its damage is the point.
func suitsFamily(stats map[string]int, isWeapon bool, families []statFamily) bool {
	if hasAny(stats, sharedStats) {
		return true
	}
	carriesFamilyStat := false
	for _, family := range []statFamily{familyPhysical, familyCaster, familyHealer, familyTank} {
		if hasAny(stats, familyStats[family]) {
			carriesFamilyStat = true
			break
		}
	}
	for _, family := range families {
		switch {
		case carriesFamilyStat && hasAny(stats, familyStats[family]):
			return true
		case !carriesFamilyStat && (!isWeapon || family == familyPhysical || family == familyTank):
			return true
		}
	}
	return false
}

func hasAny(stats map[string]int, keys []string) bool {
	for _, k := range keys {
		if stats[k] > 0 {
			return true
		}
	}
	return false
}
