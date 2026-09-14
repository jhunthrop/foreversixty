// logs/engine/units/specs.go
package units

// RetailSpecClass maps a retail COMBATANT_INFO spec id to its class. The
// table is transcribed from wowcoach.gg/docs/combat-log/spec.yaml (enum
// spec_ids) and is the only class lookup that ships in this module: every
// other mapping is injected through Options so no spell id is hard-coded.
//
// Forever's own spec ids, if it writes any, become a second table here when
// the first beta log arrives.
var RetailSpecClass = map[int64]string{
	250: "Death Knight", 251: "Death Knight", 252: "Death Knight",
	577: "Demon Hunter", 581: "Demon Hunter", 1480: "Demon Hunter",
	102: "Druid", 103: "Druid", 104: "Druid", 105: "Druid",
	1467: "Evoker", 1468: "Evoker", 1473: "Evoker",
	253: "Hunter", 254: "Hunter", 255: "Hunter",
	62: "Mage", 63: "Mage", 64: "Mage",
	268: "Monk", 269: "Monk", 270: "Monk",
	65: "Paladin", 66: "Paladin", 70: "Paladin",
	256: "Priest", 257: "Priest", 258: "Priest",
	259: "Rogue", 260: "Rogue", 261: "Rogue",
	262: "Shaman", 263: "Shaman", 264: "Shaman",
	265: "Warlock", 266: "Warlock", 267: "Warlock",
	71: "Warrior", 72: "Warrior", 73: "Warrior",
}

// RetailSpecName maps a retail spec id to the spec's name, for the roster
// and the ranking metrics rows. Same source as RetailSpecClass.
var RetailSpecName = map[int64]string{
	250: "Blood", 251: "Frost", 252: "Unholy",
	577: "Havoc", 581: "Vengeance", 1480: "Devourer",
	102: "Balance", 103: "Feral", 104: "Guardian", 105: "Restoration",
	1467: "Devastation", 1468: "Preservation", 1473: "Augmentation",
	253: "Beast Mastery", 254: "Marksmanship", 255: "Survival",
	62: "Arcane", 63: "Fire", 64: "Frost",
	268: "Brewmaster", 269: "Windwalker", 270: "Mistweaver",
	65: "Holy", 66: "Protection", 70: "Retribution",
	256: "Discipline", 257: "Holy", 258: "Shadow",
	259: "Assassination", 260: "Outlaw", 261: "Subtlety",
	262: "Elemental", 263: "Enhancement", 264: "Restoration",
	265: "Affliction", 266: "Demonology", 267: "Destruction",
	71: "Arms", 72: "Fury", 73: "Protection",
}
