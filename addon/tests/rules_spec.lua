-- addon/tests/rules_spec.lua
-- Forever's rules engine. Pure functions only: this spec never touches the
-- client mock, and every verdict below is one the in-game surfaces will show.
local Rules = require("Rules")
local L = require("Locale")

-- The client's own ids (GetItemInfoInstant's classID and subClassID).
local WEAPON, ARMOR = 2, 4
local AXE1, AXE2, MACE1, MACE2, SWORD1, SWORD2 = 0, 1, 4, 5, 7, 8
local POLEARM, STAFF, FIST, DAGGER, THROWN, WAND = 6, 10, 13, 15, 16, 19
local BOW, GUN, CROSSBOW = 2, 3, 18
local CLOTH, LEATHER, MAIL, PLATE, SHIELD, LIBRAM, IDOL, TOTEM = 1, 2, 3, 4, 6, 7, 8, 9

describe("Rules.canEquip armor", function()
	it("keeps a level 20 shaman out of mail until 40", function()
		local verdict, level, reason = Rules.canEquip("SHAMAN", 20, ARMOR, MAIL, "INVTYPE_CHEST")
		assert.are.equal("until", verdict)
		assert.are.equal(40, level)
		assert.are.equal("mail armor", reason)
	end)

	it("lets a shaman wear mail at 40 and leather at 1", function()
		assert.are.equal("ok", (Rules.canEquip("SHAMAN", 40, ARMOR, MAIL, "INVTYPE_CHEST")))
		assert.are.equal("ok", (Rules.canEquip("SHAMAN", 1, ARMOR, LEATHER, "INVTYPE_CHEST")))
	end)

	it("keeps a level 39 warrior out of plate until 40", function()
		local verdict, level = Rules.canEquip("WARRIOR", 39, ARMOR, PLATE, "INVTYPE_LEGS")
		assert.are.equal("until", verdict)
		assert.are.equal(40, level)
		assert.are.equal("ok", (Rules.canEquip("WARRIOR", 40, ARMOR, PLATE, "INVTYPE_LEGS")))
	end)

	it("treats a level 20 warrior in mail as fine: mail is level 1", function()
		assert.are.equal("ok", (Rules.canEquip("WARRIOR", 20, ARMOR, MAIL, "INVTYPE_CHEST")))
	end)

	it("applies the same plate and mail gate to paladins and hunters", function()
		assert.are.equal("until", (Rules.canEquip("PALADIN", 30, ARMOR, PLATE, "INVTYPE_HEAD")))
		assert.are.equal("until", (Rules.canEquip("HUNTER", 30, ARMOR, MAIL, "INVTYPE_HEAD")))
		assert.are.equal("ok", (Rules.canEquip("HUNTER", 40, ARMOR, MAIL, "INVTYPE_HEAD")))
	end)

	it("lets any class wear armor below its own type", function()
		assert.are.equal("ok", (Rules.canEquip("WARRIOR", 20, ARMOR, LEATHER, "INVTYPE_CHEST")))
		assert.are.equal("ok", (Rules.canEquip("WARRIOR", 20, ARMOR, CLOTH, "INVTYPE_CHEST")))
		assert.are.equal("ok", (Rules.canEquip("DRUID", 20, ARMOR, CLOTH, "INVTYPE_CHEST")))
		assert.are.equal("ok", (Rules.canEquip("HUNTER", 20, ARMOR, CLOTH, "INVTYPE_CHEST")))
	end)

	it("never lets a class wear armor above anything it will ever train", function()
		local verdict, level, reason = Rules.canEquip("ROGUE", 60, ARMOR, MAIL, "INVTYPE_CHEST")
		assert.are.equal("never", verdict)
		assert.is_nil(level)
		assert.are.equal("mail armor", reason)
		assert.are.equal("never", (Rules.canEquip("MAGE", 60, ARMOR, LEATHER, "INVTYPE_CHEST")))
		assert.are.equal("never", (Rules.canEquip("PRIEST", 60, ARMOR, PLATE, "INVTYPE_CHEST")))
		assert.are.equal("never", (Rules.canEquip("SHAMAN", 60, ARMOR, PLATE, "INVTYPE_CHEST")))
		assert.are.equal("never", (Rules.canEquip("DRUID", 60, ARMOR, MAIL, "INVTYPE_CHEST")))
	end)

	it("does not gate cloaks, rings, necks, trinkets or shirts by armor type", function()
		assert.are.equal("ok", (Rules.canEquip("WARRIOR", 1, ARMOR, CLOTH, "INVTYPE_CLOAK")))
		assert.are.equal("ok", (Rules.canEquip("MAGE", 1, ARMOR, 0, "INVTYPE_FINGER")))
		assert.are.equal("ok", (Rules.canEquip("MAGE", 1, ARMOR, 0, "INVTYPE_NECK")))
		assert.are.equal("ok", (Rules.canEquip("MAGE", 1, ARMOR, 0, "INVTYPE_TRINKET")))
	end)

	it("reads the client's English type names when no ids are given", function()
		local verdict, level = Rules.canEquip("SHAMAN", 20, "Armor", "Mail", "INVTYPE_CHEST")
		assert.are.equal("until", verdict)
		assert.are.equal(40, level)
		assert.are.equal("ok", (Rules.canEquip("SHAMAN", 20, "Armor", "Leather", "INVTYPE_CHEST")))
	end)

	it("says nothing about an item it cannot classify", function()
		assert.are.equal("ok", (Rules.canEquip("SHAMAN", 20, 0, 0, "")))
		assert.are.equal("ok", (Rules.canEquip("SHAMAN", 20, nil, nil, nil)))
		assert.are.equal("ok", (Rules.canEquip("SHAMAN", nil, ARMOR, MAIL, "INVTYPE_CHEST")))
		assert.are.equal("ok", (Rules.canEquip("DEATHKNIGHT", 20, ARMOR, PLATE, "INVTYPE_CHEST")))
	end)
end)

describe("Rules.canEquip weapons", function()
	it("trains a shaman in axes, maces, fists, staves and daggers", function()
		for _, sub in ipairs({ AXE1, AXE2, MACE1, MACE2, FIST, STAFF, DAGGER }) do
			assert.are.equal("ok", (Rules.canEquip("SHAMAN", 1, WEAPON, sub, "INVTYPE_WEAPONMAINHAND")), sub)
		end
	end)

	it("does not train a shaman in swords or polearms", function()
		local verdict, level, reason = Rules.canEquip("SHAMAN", 60, WEAPON, SWORD2, "INVTYPE_2HWEAPON")
		assert.are.equal("never", verdict)
		assert.is_nil(level)
		assert.are.equal("two-handed swords", reason)
		assert.are.equal("never", (Rules.canEquip("SHAMAN", 60, WEAPON, SWORD1, "INVTYPE_WEAPON")))
		assert.are.equal("never", (Rules.canEquip("SHAMAN", 60, WEAPON, POLEARM, "INVTYPE_2HWEAPON")))
	end)

	it("lets a shaman carry a shield or a held-in-off-hand item", function()
		assert.are.equal("ok", (Rules.canEquip("SHAMAN", 20, ARMOR, SHIELD, "INVTYPE_SHIELD")))
		assert.are.equal("ok", (Rules.canEquip("SHAMAN", 20, ARMOR, 0, "INVTYPE_HOLDABLE")))
	end)

	it("never lets a shaman dual wield: a second one-hand weapon is refused", function()
		local verdict, level, reason = Rules.canEquip("SHAMAN", 60, WEAPON, MACE1, "INVTYPE_WEAPON", "off_hand")
		assert.are.equal("never", verdict)
		assert.is_nil(level)
		assert.are.equal("dual wielding", reason)
		-- An off-hand-only weapon is a second weapon by its own location.
		assert.are.equal("never", (Rules.canEquip("SHAMAN", 60, WEAPON, DAGGER, "INVTYPE_WEAPONOFFHAND")))
	end)

	it("lets the same shaman use that one-hand weapon in the main hand", function()
		assert.are.equal("ok", (Rules.canEquip("SHAMAN", 60, WEAPON, MACE1, "INVTYPE_WEAPON", "main_hand")))
		assert.are.equal("ok", (Rules.canEquip("SHAMAN", 60, WEAPON, MACE1, "INVTYPE_WEAPON")))
	end)

	it("trains a rogue in maces, daggers, swords and fists, and lets it dual wield", function()
		assert.are.equal("ok", (Rules.canEquip("ROGUE", 20, WEAPON, MACE1, "INVTYPE_WEAPON")))
		assert.are.equal("ok", (Rules.canEquip("ROGUE", 20, WEAPON, MACE1, "INVTYPE_WEAPON", "off_hand")))
		assert.are.equal("ok", (Rules.canEquip("ROGUE", 20, WEAPON, DAGGER, "INVTYPE_WEAPONOFFHAND")))
		assert.are.equal("ok", (Rules.canEquip("ROGUE", 20, WEAPON, SWORD1, "INVTYPE_WEAPONMAINHAND")))
		assert.are.equal("ok", (Rules.canEquip("ROGUE", 20, WEAPON, FIST, "INVTYPE_WEAPON")))
	end)

	it("keeps a rogue out of two-handers and axes", function()
		assert.are.equal("never", (Rules.canEquip("ROGUE", 60, WEAPON, MACE2, "INVTYPE_2HWEAPON")))
		assert.are.equal("never", (Rules.canEquip("ROGUE", 60, WEAPON, AXE1, "INVTYPE_WEAPON")))
	end)

	it("reads the client's English weapon names when no ids are given", function()
		assert.are.equal("ok", (Rules.canEquip("ROGUE", 20, "Weapon", "One-Handed Maces", "INVTYPE_WEAPON")))
		assert.are.equal("never", (Rules.canEquip("ROGUE", 20, "Weapon", "Two-Handed Maces", "INVTYPE_2HWEAPON")))
		assert.are.equal("ok", (Rules.canEquip("SHAMAN", 20, "Weapon", "Fist Weapons", "INVTYPE_WEAPON")))
		assert.are.equal("never", (Rules.canEquip(
			"SHAMAN", 20, "Weapon", "One-Handed Maces", "INVTYPE_WEAPON", "off_hand")))
	end)

	it("keeps the other classes' tables at their vanilla training", function()
		assert.are.equal("ok", (Rules.canEquip("WARRIOR", 20, WEAPON, SWORD2, "INVTYPE_2HWEAPON")))
		assert.are.equal("ok", (Rules.canEquip("WARRIOR", 20, WEAPON, AXE1, "INVTYPE_WEAPON", "off_hand")))
		assert.are.equal("ok", (Rules.canEquip("WARRIOR", 20, ARMOR, SHIELD, "INVTYPE_SHIELD")))
		assert.are.equal("ok", (Rules.canEquip("PALADIN", 20, WEAPON, MACE2, "INVTYPE_2HWEAPON")))
		assert.are.equal("never", (Rules.canEquip("PALADIN", 20, WEAPON, MACE1, "INVTYPE_WEAPON", "off_hand")))
		assert.are.equal("never", (Rules.canEquip("PALADIN", 20, WEAPON, DAGGER, "INVTYPE_WEAPON")))
		assert.are.equal("ok", (Rules.canEquip("HUNTER", 20, WEAPON, POLEARM, "INVTYPE_2HWEAPON")))
		assert.are.equal("never", (Rules.canEquip("HUNTER", 20, WEAPON, MACE1, "INVTYPE_WEAPON")))
		assert.are.equal("ok", (Rules.canEquip("DRUID", 20, WEAPON, STAFF, "INVTYPE_2HWEAPON")))
		assert.are.equal("never", (Rules.canEquip("DRUID", 20, ARMOR, SHIELD, "INVTYPE_SHIELD")))
		assert.are.equal("ok", (Rules.canEquip("PRIEST", 20, WEAPON, MACE1, "INVTYPE_WEAPON")))
		assert.are.equal("never", (Rules.canEquip("PRIEST", 20, WEAPON, SWORD1, "INVTYPE_WEAPON")))
		assert.are.equal("ok", (Rules.canEquip("MAGE", 20, WEAPON, SWORD1, "INVTYPE_WEAPON")))
		assert.are.equal("ok", (Rules.canEquip("WARLOCK", 20, WEAPON, DAGGER, "INVTYPE_WEAPON")))
		assert.are.equal("never", (Rules.canEquip("MAGE", 20, WEAPON, MACE1, "INVTYPE_WEAPON")))
	end)

	it("sorts ranged weapons by class", function()
		assert.are.equal("ok", (Rules.canEquip("HUNTER", 20, WEAPON, BOW, "INVTYPE_RANGED")))
		assert.are.equal("ok", (Rules.canEquip("HUNTER", 20, WEAPON, GUN, "INVTYPE_RANGEDRIGHT")))
		assert.are.equal("ok", (Rules.canEquip("WARRIOR", 20, WEAPON, CROSSBOW, "INVTYPE_RANGED")))
		assert.are.equal("ok", (Rules.canEquip("ROGUE", 20, WEAPON, THROWN, "INVTYPE_THROWN")))
		assert.are.equal("ok", (Rules.canEquip("MAGE", 20, WEAPON, WAND, "INVTYPE_RANGEDRIGHT")))
		assert.are.equal("never", (Rules.canEquip("SHAMAN", 20, WEAPON, BOW, "INVTYPE_RANGED")))
		assert.are.equal("never", (Rules.canEquip("WARRIOR", 20, WEAPON, WAND, "INVTYPE_RANGEDRIGHT")))
		assert.are.equal("never", (Rules.canEquip("PALADIN", 20, WEAPON, GUN, "INVTYPE_RANGED")))
	end)

	it("sorts librams, idols and totems by class", function()
		assert.are.equal("ok", (Rules.canEquip("PALADIN", 20, ARMOR, LIBRAM, "INVTYPE_RELIC")))
		assert.are.equal("ok", (Rules.canEquip("DRUID", 20, ARMOR, IDOL, "INVTYPE_RELIC")))
		assert.are.equal("ok", (Rules.canEquip("SHAMAN", 20, ARMOR, TOTEM, "INVTYPE_RELIC")))
		assert.are.equal("never", (Rules.canEquip("SHAMAN", 20, ARMOR, LIBRAM, "INVTYPE_RELIC")))
		assert.are.equal("never", (Rules.canEquip("WARRIOR", 20, ARMOR, TOTEM, "INVTYPE_RELIC")))
	end)
end)

describe("Rules tables", function()
	it("carries the armor ladder as the design gives it", function()
		assert.are.same({ mail = 1, plate = 40 }, Rules.ARMOR_AT.WARRIOR)
		assert.are.same({ mail = 1, plate = 40 }, Rules.ARMOR_AT.PALADIN)
		assert.are.same({ leather = 1, mail = 40 }, Rules.ARMOR_AT.HUNTER)
		assert.are.same({ leather = 1, mail = 40 }, Rules.ARMOR_AT.SHAMAN)
		assert.are.same({ leather = 1 }, Rules.ARMOR_AT.ROGUE)
		assert.are.same({ leather = 1 }, Rules.ARMOR_AT.DRUID)
		assert.are.same({ cloth = 1 }, Rules.ARMOR_AT.PRIEST)
		assert.are.same({ cloth = 1 }, Rules.ARMOR_AT.MAGE)
		assert.are.same({ cloth = 1 }, Rules.ARMOR_AT.WARLOCK)
	end)

	it("carries the shaman and rogue weapon tables as the design gives them", function()
		assert.are.same({
			axe1 = true, axe2 = true, mace1 = true, mace2 = true, fist = true, staff = true,
			dagger = true, shield = true, offhand = true, dualWield = false,
		}, Rules.WEAPONS.SHAMAN)
		assert.are.same({
			dagger = true, sword1 = true, mace1 = true, fist = true, dualWield = true,
		}, Rules.WEAPONS.ROGUE)
	end)

	it("has a weapon table for every class that has an armor row", function()
		for class in pairs(Rules.ARMOR_AT) do
			assert.is_table(Rules.WEAPONS[class], class)
		end
	end)

	it("only ever lets warriors, hunters and rogues dual wield", function()
		local dual = {}
		for class, row in pairs(Rules.WEAPONS) do
			if row.dualWield then
				dual[#dual + 1] = class
			end
		end
		table.sort(dual)
		assert.are.same({ "HUNTER", "ROGUE", "WARRIOR" }, dual)
	end)

	it("has the two player-facing refusal lines", function()
		assert.are.equal("Not for you until %d: %s", L.rulesUntilLevel)
		assert.are.equal("Not for you: %s", L.rulesNever)
	end)
end)

describe("Rules.statKey", function()
	it("maps hit and crit to the one stat each the weights table uses", function()
		assert.are.equal("hit", Rules.statKey("ITEM_MOD_HIT_RATING_SHORT"))
		assert.are.equal("crit", Rules.statKey("ITEM_MOD_CRIT_RATING_SHORT"))
	end)

	it("folds the melee, ranged and spell variants into the same stat", function()
		for _, kind in ipairs({ "MELEE", "RANGED", "SPELL" }) do
			assert.are.equal("hit", Rules.statKey("ITEM_MOD_HIT_" .. kind .. "_RATING_SHORT"), kind)
			assert.are.equal("crit", Rules.statKey("ITEM_MOD_CRIT_" .. kind .. "_RATING_SHORT"), kind)
		end
	end)

	it("maps the primary stats and armor", function()
		assert.are.equal("strength", Rules.statKey("ITEM_MOD_STRENGTH_SHORT"))
		assert.are.equal("spell_power", Rules.statKey("ITEM_MOD_SPELL_POWER_SHORT"))
		assert.are.equal("armor", Rules.statKey("RESISTANCE0_NAME"))
	end)

	it("answers nil for a key it does not know, rather than guessing", function()
		assert.is_nil(Rules.statKey("ITEM_MOD_SOMETHING_NEW_SHORT"))
		assert.is_nil(Rules.statKey(nil))
	end)
end)

describe("Rules.racePairAllowed", function()
	it("allows the six new pairs", function()
		assert.is_true(Rules.racePairAllowed("Gnome", "PRIEST"))
		assert.is_true(Rules.racePairAllowed("Human", "HUNTER"))
		assert.is_true(Rules.racePairAllowed("Dwarf", "SHAMAN"))
		assert.is_true(Rules.racePairAllowed("Orc", "MAGE"))
		assert.is_true(Rules.racePairAllowed("Troll", "WARLOCK"))
		assert.is_true(Rules.racePairAllowed("Undead", "PALADIN"))
	end)

	it("takes the client's own race token for undead and night elves", function()
		assert.is_true(Rules.racePairAllowed("Scourge", "PALADIN"))
		assert.is_true(Rules.racePairAllowed("NightElf", "DRUID"))
		assert.is_false(Rules.racePairAllowed("NightElf", "MAGE"))
	end)

	it("keeps every vanilla pair", function()
		local vanilla = {
			Human = { "WARRIOR", "PALADIN", "ROGUE", "PRIEST", "MAGE", "WARLOCK" },
			Dwarf = { "WARRIOR", "PALADIN", "HUNTER", "ROGUE", "PRIEST" },
			NightElf = { "WARRIOR", "HUNTER", "ROGUE", "PRIEST", "DRUID" },
			Gnome = { "WARRIOR", "ROGUE", "MAGE", "WARLOCK" },
			Orc = { "WARRIOR", "HUNTER", "ROGUE", "SHAMAN", "WARLOCK" },
			Undead = { "WARRIOR", "ROGUE", "PRIEST", "MAGE", "WARLOCK" },
			Tauren = { "WARRIOR", "HUNTER", "SHAMAN", "DRUID" },
			Troll = { "WARRIOR", "HUNTER", "ROGUE", "PRIEST", "MAGE", "SHAMAN" },
		}
		for race, classes in pairs(vanilla) do
			for _, class in ipairs(classes) do
				assert.is_true(Rules.racePairAllowed(race, class), race .. " " .. class)
			end
		end
	end)

	it("allows no pair beyond those: exactly 40 vanilla and 6 new of the 72", function()
		local races = { "Human", "Dwarf", "NightElf", "Gnome", "Orc", "Undead", "Tauren", "Troll" }
		local classes = { "WARRIOR", "PALADIN", "HUNTER", "ROGUE", "PRIEST", "SHAMAN", "MAGE", "WARLOCK", "DRUID" }
		local allowed = 0
		for _, race in ipairs(races) do
			for _, class in ipairs(classes) do
				if Rules.racePairAllowed(race, class) then
					allowed = allowed + 1
				end
			end
		end
		assert.are.equal(40 + 6, allowed)
	end)

	it("refuses the pairs the six do not cover", function()
		assert.is_false(Rules.racePairAllowed("Tauren", "MAGE"))
		assert.is_false(Rules.racePairAllowed("Tauren", "PALADIN"))
		assert.is_false(Rules.racePairAllowed("Gnome", "HUNTER"))
		assert.is_false(Rules.racePairAllowed("Human", "SHAMAN"))
		assert.is_false(Rules.racePairAllowed("Undead", "HUNTER"))
		assert.is_false(Rules.racePairAllowed("Undead", "DRUID"))
	end)

	it("refuses a race or class it does not know", function()
		assert.is_false(Rules.racePairAllowed("Draenei", "PALADIN"))
		assert.is_false(Rules.racePairAllowed("Human", "DEATHKNIGHT"))
		assert.is_false(Rules.racePairAllowed(nil, "MAGE"))
		assert.is_false(Rules.racePairAllowed("Human", nil))
	end)
end)

describe("Rules.questRewardFloor", function()
	it("counts a quest reward from three levels below the quest", function()
		assert.are.equal(20, Rules.questRewardFloor(23))
		assert.are.equal(57, Rules.questRewardFloor(60))
		assert.are.equal(1, Rules.questRewardFloor(4))
	end)
end)
