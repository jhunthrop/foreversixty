-- addon/ForeverSixty/Rules.lua
-- The Forever rules engine: "can this character use this item, and when".
--
-- Every verdict the addon makes about an item goes through here, so a rule
-- that differs from Classic (a shaman who cannot dual wield, mail and plate
-- arriving at 40, hit and crit as one stat, six new race and class pairs)
-- is written down once. Pure functions over client ids and tokens: nothing
-- here reads the client, so a busted spec pins every rule without a mock.
--
-- A rule this module does not know is never a refusal. An item it cannot
-- classify, a class it has no row for, a level that is not a number all
-- answer "ok": a wrong "not for you" is worse than saying nothing.
local _, ns = ...
ns = type(ns) == "table" and ns or {}

local Rules = {}

--- Armor types, lowest first. A class wears its own type and every type
--- below it, so the rank is all the "below its own type" rule needs.
local ARMOR_RANK = { cloth = 1, leather = 2, mail = 3, plate = 4 }

--- The level each class starts to wear each armor type. Only the types a
--- class trains are listed; anything lower in ARMOR_RANK is worn anyway.
Rules.ARMOR_AT = {
	WARRIOR = { mail = 1, plate = 40 },
	PALADIN = { mail = 1, plate = 40 },
	HUNTER = { leather = 1, mail = 40 },
	SHAMAN = { leather = 1, mail = 40 },
	ROGUE = { leather = 1 },
	DRUID = { leather = 1 },
	PRIEST = { cloth = 1 },
	MAGE = { cloth = 1 },
	WARLOCK = { cloth = 1 },
}

--- Melee weapons, shields and held-in-off-hand items per class, and whether
--- the class may carry a second weapon. SHAMAN and ROGUE are Forever's own
--- rows (shamans train axes, maces, fists, staves and daggers but cannot dual
--- wield; rogues train maces beside daggers, swords and fists). The other
--- classes are Classic's training, which Forever has not been seen to change.
--- A held-in-off-hand item is allowed unless a row says `offhand = false`.
Rules.WEAPONS = {
	SHAMAN = {
		axe1 = true, axe2 = true, mace1 = true, mace2 = true, fist = true, staff = true,
		dagger = true, shield = true, offhand = true, dualWield = false,
	},
	ROGUE = { dagger = true, sword1 = true, mace1 = true, fist = true, dualWield = true },
	WARRIOR = {
		axe1 = true, axe2 = true, mace1 = true, mace2 = true, sword1 = true, sword2 = true,
		dagger = true, fist = true, polearm = true, staff = true, shield = true, dualWield = true,
	},
	PALADIN = {
		axe1 = true, axe2 = true, mace1 = true, mace2 = true, sword1 = true, sword2 = true,
		polearm = true, shield = true, dualWield = false,
	},
	HUNTER = {
		axe1 = true, axe2 = true, sword1 = true, sword2 = true, dagger = true, fist = true,
		polearm = true, staff = true, dualWield = true,
	},
	DRUID = { mace1 = true, mace2 = true, dagger = true, fist = true, polearm = true, staff = true, dualWield = false },
	PRIEST = { mace1 = true, dagger = true, staff = true, dualWield = false },
	MAGE = { sword1 = true, dagger = true, staff = true, dualWield = false },
	WARLOCK = { sword1 = true, dagger = true, staff = true, dualWield = false },
}

--- Ranged weapons per class, kept apart from WEAPONS because they take the
--- ranged slot, never a hand.
Rules.RANGED = {
	WARRIOR = { bow = true, gun = true, crossbow = true, thrown = true },
	HUNTER = { bow = true, gun = true, crossbow = true, thrown = true },
	ROGUE = { bow = true, gun = true, crossbow = true, thrown = true },
	PRIEST = { wand = true },
	MAGE = { wand = true },
	WARLOCK = { wand = true },
}

--- The one relic kind each class carries in the relic slot.
Rules.RELICS = { PALADIN = "libram", DRUID = "idol", SHAMAN = "totem" }

--- The client's class ids (GetItemInfoInstant's classID) and their English
--- names (its itemType, the fallback where no id is given).
local CLASS_WEAPON, CLASS_ARMOR = 2, 4
local KIND_BY_CLASS = {
	[CLASS_WEAPON] = "weapon", Weapon = "weapon",
	[CLASS_ARMOR] = "armor", Armor = "armor",
}

--- Weapon subclass id / English itemSubType -> the key the WEAPONS and RANGED
--- tables use. Subclasses with no key (miscellaneous, fishing poles, spears)
--- are left out and so are never refused.
local WEAPON_KEY = {
	[0] = "axe1", [1] = "axe2", [2] = "bow", [3] = "gun", [4] = "mace1", [5] = "mace2",
	[6] = "polearm", [7] = "sword1", [8] = "sword2", [10] = "staff", [13] = "fist",
	[15] = "dagger", [16] = "thrown", [18] = "crossbow", [19] = "wand",
	["One-Handed Axes"] = "axe1", ["Two-Handed Axes"] = "axe2", ["Bows"] = "bow",
	["Guns"] = "gun", ["One-Handed Maces"] = "mace1", ["Two-Handed Maces"] = "mace2",
	["Polearms"] = "polearm", ["One-Handed Swords"] = "sword1", ["Two-Handed Swords"] = "sword2",
	["Staves"] = "staff", ["Fist Weapons"] = "fist", ["Daggers"] = "dagger",
	["Thrown"] = "thrown", ["Crossbows"] = "crossbow", ["Wands"] = "wand",
}

--- The same keys as the plain words a refusal names ("Not for you: ...").
local WEAPON_NAME = {
	axe1 = "one-handed axes", axe2 = "two-handed axes", bow = "bows", gun = "guns",
	mace1 = "one-handed maces", mace2 = "two-handed maces", polearm = "polearms",
	sword1 = "one-handed swords", sword2 = "two-handed swords", staff = "staves",
	fist = "fist weapons", dagger = "daggers", thrown = "thrown weapons",
	crossbow = "crossbows", wand = "wands",
}

--- Armor subclass id / English itemSubType -> "cloth" .. "plate", "shield",
--- or a relic kind. Miscellaneous (rings, necks, held-in-off-hand) is absent.
local ARMOR_KEY = {
	[1] = "cloth", [2] = "leather", [3] = "mail", [4] = "plate", [6] = "shield",
	[7] = "libram", [8] = "idol", [9] = "totem",
	["Cloth"] = "cloth", ["Leather"] = "leather", ["Mail"] = "mail", ["Plate"] = "plate",
	["Shields"] = "shield", ["Librams"] = "libram", ["Idols"] = "idol", ["Totems"] = "totem",
}

local RELIC_NAME = { libram = "librams", idol = "idols", totem = "totems" }

--- Slots whose items are not armor-type proficiency items, whatever the
--- subclass says (a cloak is "cloth" and every class wears it).
local FREE_LOCATIONS = {
	INVTYPE_CLOAK = true, INVTYPE_FINGER = true, INVTYPE_NECK = true, INVTYPE_TRINKET = true,
	INVTYPE_BODY = true, INVTYPE_TABARD = true, INVTYPE_HOLDABLE = true,
}

local function classRow(classToken)
	return type(classToken) == "string" and classToken:upper() or nil
end

--- "ok", or "until" with the level, or "never", for an armor-type item.
local function armorVerdict(class, level, key)
	local rank = ARMOR_RANK[key]
	local ladder = Rules.ARMOR_AT[class]
	local best, unlock = 0, nil
	for armorType, from in pairs(ladder) do
		local typeRank = ARMOR_RANK[armorType]
		if from <= level and typeRank > best then
			best = typeRank
		end
		-- The earliest level at which the class wears this type or a higher
		-- one: a higher type covers the item as well, since lower is worn.
		if typeRank >= rank and (unlock == nil or from < unlock) then
			unlock = from
		end
	end
	if rank <= best then
		return "ok"
	end
	if unlock ~= nil then
		return "until", unlock, key .. " armor"
	end
	return "never", nil, key .. " armor"
end

local function weaponVerdict(class, key, equipLoc, slot)
	local ranged = Rules.RANGED[class] or {}
	local melee = Rules.WEAPONS[class]
	local isRanged = key == "bow" or key == "gun" or key == "crossbow" or key == "thrown" or key == "wand"
	if isRanged then
		if ranged[key] ~= true then
			return "never", nil, WEAPON_NAME[key]
		end
		return "ok"
	end
	if melee[key] ~= true then
		return "never", nil, WEAPON_NAME[key]
	end
	-- A weapon judged for the off hand is a second weapon: either the item
	-- can only go there, or the caller says it is being judged for that slot.
	local second = equipLoc == "INVTYPE_WEAPONOFFHAND"
		or (equipLoc == "INVTYPE_WEAPON" and slot == "off_hand")
	if second and melee.dualWield ~= true then
		return "never", nil, "dual wielding"
	end
	return "ok"
end

--- Can a character of `classToken` at `level` use this item, and when?
--- itemClass and itemSubClass are the client's numeric ids, or the English
--- itemType and itemSubType strings where the client gave none. equipLoc is
--- the INVTYPE_* location. `slot` is optional: pass "off_hand" to judge a
--- one-hand weapon as the second weapon (a shaman cannot carry one).
--- @return "ok" | "until" | "never", level|nil, reason string|nil
function Rules.canEquip(classToken, level, itemClass, itemSubClass, equipLoc, slot)
	local class = classRow(classToken)
	if class == nil or Rules.ARMOR_AT[class] == nil or type(level) ~= "number" then
		return "ok"
	end
	local kind = KIND_BY_CLASS[itemClass]
	if kind == "weapon" then
		local key = WEAPON_KEY[itemSubClass]
		if key == nil then
			return "ok"
		end
		return weaponVerdict(class, key, equipLoc, slot)
	end
	if kind ~= "armor" then
		return "ok"
	end
	local key = ARMOR_KEY[itemSubClass]
	if equipLoc == "INVTYPE_SHIELD" or key == "shield" then
		if Rules.WEAPONS[class].shield ~= true then
			return "never", nil, "shields"
		end
		return "ok"
	end
	if equipLoc == "INVTYPE_HOLDABLE" then
		if Rules.WEAPONS[class].offhand == false then
			return "never", nil, "items held in the off hand"
		end
		return "ok"
	end
	if RELIC_NAME[key] ~= nil then
		if Rules.RELICS[class] ~= key then
			return "never", nil, RELIC_NAME[key]
		end
		return "ok"
	end
	if ARMOR_RANK[key] == nil or FREE_LOCATIONS[equipLoc] then
		return "ok"
	end
	return armorVerdict(class, level, key)
end

--- GetItemStats' own key names -> contract 10.8's stat names, which is also
--- what Data.lua's weights are keyed by. Hit and crit are one stat each in
--- Forever, applying to spells and melee alike, so the melee, ranged and spell
--- variants a client may return all fold into "hit" and "crit". Spike check 11
--- confirms the set the beta client returns; a key absent from this table is
--- left out of the score rather than guessed at, because a wrong stat is a
--- wrong recommendation and a missing one is only a conservative score.
Rules.STAT_KEYS = {
	ITEM_MOD_STRENGTH_SHORT = "strength",
	ITEM_MOD_AGILITY_SHORT = "agility",
	ITEM_MOD_STAMINA_SHORT = "stamina",
	ITEM_MOD_INTELLECT_SHORT = "intellect",
	ITEM_MOD_SPIRIT_SHORT = "spirit",
	ITEM_MOD_ATTACK_POWER_SHORT = "attack_power",
	ITEM_MOD_RANGED_ATTACK_POWER_SHORT = "ranged_attack_power",
	ITEM_MOD_SPELL_POWER_SHORT = "spell_power",
	ITEM_MOD_SPELL_DAMAGE_DONE_SHORT = "spell_damage",
	ITEM_MOD_SPELL_HEALING_DONE_SHORT = "healing_power",
	ITEM_MOD_HIT_RATING_SHORT = "hit",
	ITEM_MOD_HIT_MELEE_RATING_SHORT = "hit",
	ITEM_MOD_HIT_RANGED_RATING_SHORT = "hit",
	ITEM_MOD_HIT_SPELL_RATING_SHORT = "hit",
	ITEM_MOD_CRIT_RATING_SHORT = "crit",
	ITEM_MOD_CRIT_MELEE_RATING_SHORT = "crit",
	ITEM_MOD_CRIT_RANGED_RATING_SHORT = "crit",
	ITEM_MOD_CRIT_SPELL_RATING_SHORT = "crit",
	ITEM_MOD_HASTE_RATING_SHORT = "melee_haste",
	ITEM_MOD_EXPERTISE_RATING_SHORT = "expertise",
	ITEM_MOD_DEFENSE_SKILL_RATING_SHORT = "defense",
	ITEM_MOD_DODGE_RATING_SHORT = "dodge",
	ITEM_MOD_PARRY_RATING_SHORT = "parry",
	ITEM_MOD_BLOCK_RATING_SHORT = "block",
	ITEM_MOD_BLOCK_VALUE_SHORT = "block_value",
	ITEM_MOD_POWER_REGEN0_SHORT = "mp5",
	ITEM_MOD_SPELL_PENETRATION_SHORT = "spell_penetration",
	ITEM_MOD_ARMOR_PENETRATION_RATING_SHORT = "armor_penetration",
	RESISTANCE0_NAME = "armor",
	-- A weapon's own damage per second (a client that omits the key gives none,
	-- and the weapon is scored on its stat line alone). No weight carries it:
	-- Progress.itemScore converts it at the sim's 14 attack power per DPS.
	ITEM_MOD_DAMAGE_PER_SECOND_SHORT = "weapon_dps",
}

--- hit and crit are one stat in Forever; the key the weights table uses, or
--- nil for a client key this module does not know.
function Rules.statKey(itemStatKey)
	if itemStatKey == nil then
		return nil
	end
	return Rules.STAT_KEYS[itemStatKey]
end

--- Race and class pairs. Classic's pairs, plus the six Forever adds; no
--- others. Keyed by the race with its case and spacing stripped, so the
--- client's token ("NightElf"), the English name and the site's slug read alike.
local PAIRS = {
	human = { WARRIOR = true, PALADIN = true, ROGUE = true, PRIEST = true, MAGE = true, WARLOCK = true },
	dwarf = { WARRIOR = true, PALADIN = true, HUNTER = true, ROGUE = true, PRIEST = true },
	nightelf = { WARRIOR = true, HUNTER = true, ROGUE = true, PRIEST = true, DRUID = true },
	gnome = { WARRIOR = true, ROGUE = true, MAGE = true, WARLOCK = true },
	orc = { WARRIOR = true, HUNTER = true, ROGUE = true, SHAMAN = true, WARLOCK = true },
	undead = { WARRIOR = true, ROGUE = true, PRIEST = true, MAGE = true, WARLOCK = true },
	tauren = { WARRIOR = true, HUNTER = true, SHAMAN = true, DRUID = true },
	troll = { WARRIOR = true, HUNTER = true, ROGUE = true, PRIEST = true, MAGE = true, SHAMAN = true },
}

--- The six pairs Forever adds on top.
local NEW_PAIRS = {
	{ "gnome", "PRIEST" }, { "human", "HUNTER" }, { "dwarf", "SHAMAN" },
	{ "orc", "MAGE" }, { "troll", "WARLOCK" }, { "undead", "PALADIN" },
}
for _, pair in ipairs(NEW_PAIRS) do
	PAIRS[pair[1]][pair[2]] = true
end

--- The client's own token for undead is "Scourge".
local RACE_ALIAS = { scourge = "undead" }

function Rules.racePairAllowed(raceToken, classToken)
	if type(raceToken) ~= "string" or type(classToken) ~= "string" then
		return false
	end
	local race = raceToken:lower():gsub("[^%a]", "")
	race = RACE_ALIAS[race] or race
	local classes = PAIRS[race]
	return classes ~= nil and classes[classToken:upper()] == true
end

--- A quest reward counts as obtainable from three levels below its quest.
function Rules.questRewardFloor(questLevel)
	return questLevel - 3
end

ns.Rules = Rules
return Rules
