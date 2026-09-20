-- addon/ForeverSixty/Gear.lua
-- Scoring what the player is carrying against the build they loaded.
--
-- The addon never needs the site's item database: the addon code carries the
-- planned items' stats, and everything in the bags is read from the client
-- with GetItemStats. The two meet in one vocabulary -- parity contract
-- 10.8's, which is also what Data.lua's weights are keyed by -- so no
-- translation table sits between a weight and a stat.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Export = ns.Export or require("Export")
local Talents = ns.Talents or require("Talents")

local Gear = {}

--- GetItemStats' own key names -> contract 10.8's stat names. Spike check 11
--- confirms the set the beta client returns; a key absent from this table is
--- left out of the score rather than guessed at, because a wrong stat is a
--- wrong recommendation and a missing one is only a conservative score.
Gear.STAT_KEYS = {
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
	ITEM_MOD_CRIT_RATING_SHORT = "crit",
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
}

function Gear.statsOf(link)
	local raw = link and GetItemStats(link) or nil
	local stats = {}
	if raw == nil then
		return stats
	end
	for key, value in pairs(raw) do
		local name = Gear.STAT_KEYS[key]
		-- Skip the private keys the spec mock uses to stand in for
		-- GetItemInfoInstant, and any client key this table does not know.
		if name ~= nil and type(value) == "number" then
			stats[name] = (stats[name] or 0) + value
		end
	end
	return stats
end

function Gear.score(stats, weights)
	local total = 0
	for name, value in pairs(stats or {}) do
		total = total + value * ((weights or {})[name] or 0)
	end
	return total
end

local function slugify(name)
	if name == nil or name == "" then
		return nil
	end
	return (name:lower():gsub("[^%w]+", "-"):gsub("^%-+", ""):gsub("%-+$", ""))
end

--- The spec key: the tree with the most points, ties to the first tree.
--- Design "Stat weights" states exactly this rule; the site's own spec
--- derivation uses it too, so a build reads the same in both places.
function Gear.specOf(data, classSlug, ranks)
	local class = data.classes[classSlug]
	if class == nil then
		return nil
	end
	local best, bestTab = -1, 1
	for tab = 1, #class.tabs do
		local total = 0
		for _, rank in pairs(ranks[tab] or {}) do
			total = total + rank
		end
		if total > best then
			best, bestTab = total, tab
		end
	end
	-- The spec slug is the tab name lowercased, which is how
	-- curated/specs.json's spec_slug is built from the tree name.
	return classSlug .. "-" .. slugify(class.tabs[bestTab].name)
end

--- The client's INVTYPE_* -> the site's slot names an item may occupy.
Gear.SLOTS_BY_EQUIP_LOCATION = {
	INVTYPE_HEAD = { "head" },
	INVTYPE_NECK = { "neck" },
	INVTYPE_SHOULDER = { "shoulder" },
	INVTYPE_CLOAK = { "back" },
	INVTYPE_CHEST = { "chest" },
	INVTYPE_ROBE = { "chest" },
	INVTYPE_WRIST = { "wrist" },
	INVTYPE_HAND = { "hands" },
	INVTYPE_WAIST = { "waist" },
	INVTYPE_LEGS = { "legs" },
	INVTYPE_FEET = { "feet" },
	INVTYPE_FINGER = { "finger1", "finger2" },
	INVTYPE_TRINKET = { "trinket1", "trinket2" },
	INVTYPE_WEAPON = { "main_hand", "off_hand" },
	INVTYPE_WEAPONMAINHAND = { "main_hand" },
	INVTYPE_WEAPONOFFHAND = { "off_hand" },
	INVTYPE_2HWEAPON = { "main_hand" },
	INVTYPE_SHIELD = { "off_hand" },
	INVTYPE_HOLDABLE = { "off_hand" },
	INVTYPE_RANGED = { "ranged" },
	INVTYPE_RANGEDRIGHT = { "ranged" },
	INVTYPE_THROWN = { "ranged" },
	INVTYPE_RELIC = { "ranged" },
}

local function plannedBySlot(build)
	local planned = {}
	for _, entry in ipairs(build.gear or {}) do
		planned[entry.slot] = entry
	end
	return planned
end

local function hasAnyStat(stats)
	return next(stats or {}) ~= nil
end

--- Every item the player can reach, equipped or bagged, with its own link.
--- The design scores both: an item already on the character can still beat the
--- planned one, and hiding that would tell a player to swap something they
--- should keep.
function Gear.candidates()
	local links = {}
	for _, entry in ipairs(Export.INVENTORY_SLOTS) do
		local link = GetInventoryItemLink("player", entry.id)
		if link ~= nil then
			links[#links + 1] = link
		end
	end
	for _, bag in ipairs(Export.CARRIED_BAGS) do
		for slot = 1, Export.containerSize(bag) do
			local link = Export.containerLink(bag, slot)
			if link ~= nil then
				links[#links + 1] = link
			end
		end
	end
	return links
end

--- Everything the player is carrying or wearing that beats the planned item in
--- a slot it fits.
function Gear.upgrades(data, build)
	local ranks = Talents.readRanks()
	local spec = Gear.specOf(data, build.classSlug, ranks)
	local weights = spec and data.weights[spec] or nil
	if weights == nil then
		return {}
	end
	local planned = plannedBySlot(build)
	local found = {}
	for _, link in ipairs(Gear.candidates()) do
		local id, _, _, equipLocation = GetItemInfoInstant(link)
		for _, target in ipairs(Gear.SLOTS_BY_EQUIP_LOCATION[equipLocation] or {}) do
			local against = planned[target]
			-- A planned item with no stats is an FS1 code, which carries none.
			-- Scoring it at zero would call every item an upgrade; saying
			-- nothing about that slot is honest. An item that IS the planned
			-- item scores a delta of exactly zero and is dropped by the test
			-- below, so a correctly geared slot never lists itself.
			if against ~= nil and hasAnyStat(against.stats) then
				local delta = Gear.score(Gear.statsOf(link), weights)
					- Gear.score(against.stats, weights)
				if delta > 0 then
					found[#found + 1] = {
						slot = target,
						itemId = tonumber(id),
						delta = delta,
						againstItemId = against.itemId,
					}
				end
			end
		end
	end
	table.sort(found, function(left, right)
		if left.delta ~= right.delta then
			return left.delta > right.delta
		end
		return left.slot < right.slot
	end)
	return found
end

function Gear.lines(data, build)
	if build == nil then
		return { L.gearNoBuild }
	end
	local spec = Gear.specOf(data, build.classSlug, Talents.readRanks())
	if spec == nil or data.weights[spec] == nil then
		return { L.gearNoWeights }
	end
	local upgrades = Gear.upgrades(data, build)
	if #upgrades == 0 then
		return { L.gearNone }
	end
	local lines = {}
	for index, upgrade in ipairs(upgrades) do
		lines[index] = string.format(L.gearUpgrade, upgrade.slot, upgrade.delta, upgrade.againstItemId)
	end
	return lines
end

ns.Gear = Gear
return Gear
