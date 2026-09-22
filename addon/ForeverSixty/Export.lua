-- addon/ForeverSixty/Export.lua
-- The character, as the string the planner reads.
--
-- Version 2 (parity contract 7, corrected by 10.5): the version-1 head plus
-- bags, bank and professions. `sets` and `loadouts` are left out -- nothing
-- in the API this addon targets states either, and inventing one would be a
-- lie the planner would then display. Codec.encodeFS1 encodes both sections
-- already, so the day the beta shows an API for them this file grows two
-- lines and the format does not move.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Codec = ns.Codec or require("Codec")
local Talents = ns.Talents or require("Talents")
local Compat = ns.Compat or require("Compat")

local Export = {}

--- The site's slot name -> the client's inventory slot id.
--- finger1/finger2 and trinket1/trinket2 are 11/12 and 13/14; the site
--- numbers them from 1 and the client does too, so the pairing is direct.
Export.INVENTORY_SLOTS = {
	{ slot = "head", id = 1 },
	{ slot = "neck", id = 2 },
	{ slot = "shoulder", id = 3 },
	{ slot = "back", id = 15 },
	{ slot = "chest", id = 5 },
	{ slot = "wrist", id = 9 },
	{ slot = "hands", id = 10 },
	{ slot = "waist", id = 6 },
	{ slot = "legs", id = 7 },
	{ slot = "feet", id = 8 },
	{ slot = "finger1", id = 11 },
	{ slot = "finger2", id = 12 },
	{ slot = "trinket1", id = 13 },
	{ slot = "trinket2", id = 14 },
	{ slot = "main_hand", id = 16 },
	{ slot = "off_hand", id = 17 },
	{ slot = "ranged", id = 18 },
}

--- Bags the player carries. 0 is the backpack; 1..4 are the equipped bags.
Export.CARRIED_BAGS = { 0, 1, 2, 3, 4 }
--- The bank. -1 is the bank window itself; 5..11 are the purchased bag slots.
Export.BANK_BAGS = { -1, 5, 6, 7, 8, 9, 10, 11 }

--- GetCurrentRegion's numbering. Spike check 13 confirms it.
Export.REGION_NAMES = { [1] = "US", [2] = "KR", [3] = "EU", [4] = "TW", [5] = "CN" }

--- How ForeverSixtyDB.savedAt reads. The Export tab shows it verbatim.
Export.SAVED_AT_FORMAT = "%Y-%m-%d %H:%M"

--- Client race token -> the site's race slug, for the eight base races. A
--- rule cannot produce this map: `Scourge` is the client's token for what
--- the site calls `undead` -- no amount of casing or hyphenation gets there
--- -- and the multi-word tokens (`NightElf`) carry no separator a rule could
--- split on either. Spike check 16 is what would confirm these against the
--- beta client; data/builds/1.60.1.69893/races.json is what confirms them
--- against the site today.
Export.RACE_SLUGS = {
	Human = "human",
	Orc = "orc",
	Dwarf = "dwarf",
	NightElf = "night-elf",
	Scourge = "undead",
	Tauren = "tauren",
	Gnome = "gnome",
	Troll = "troll",
}

--- The container API moved into C_Container; Classic Era still carries the
--- flat function on some builds. Spike check 9 says which the beta has; this
--- takes whichever exists rather than betting on one.
function Export.containerLink(bag, slot)
	if C_Container and C_Container.GetContainerItemLink then
		return C_Container.GetContainerItemLink(bag, slot)
	end
	return GetContainerItemLink and GetContainerItemLink(bag, slot) or nil
end

function Export.containerSize(bag)
	if C_Container and C_Container.GetContainerNumSlots then
		return C_Container.GetContainerNumSlots(bag) or 0
	end
	return (GetContainerNumSlots and GetContainerNumSlots(bag)) or 0
end

--- An item link's id, or nil for a link this client will not parse.
--- The lookup goes through Compat: the 1.60 client keeps
--- GetItemInfoInstant on C_Item and has no global of that name, which is
--- what the first in-game export died on. Where the client cannot answer,
--- the id is read out of the link text itself.
local function itemIdOf(link)
	if link == nil then
		return nil
	end
	local id = Compat.itemInfoInstant(link)
	if id ~= nil then
		return tonumber(id)
	end
	return tonumber(link:match("item:(%d+)"))
end

--- True when the item can be worn. An export that carried reagents, quest
--- items and bags would be several times longer for nothing: the planner
--- has no slot to put any of them in, `INVTYPE_BAG` included.
local function isEquippable(link)
	local _, _, _, equipSlot = Compat.itemInfoInstant(link)
	return equipSlot ~= nil and equipSlot ~= "" and equipSlot ~= "INVTYPE_NON_EQUIP"
		and equipSlot ~= "INVTYPE_BAG"
end

--- Public because the Export tab counts what is carried and banked
--- without rebuilding the whole string to do it.
function Export.itemsInBags(bags)
	local items = {}
	for _, bag in ipairs(bags) do
		for slot = 1, Export.containerSize(bag) do
			local link = Export.containerLink(bag, slot)
			local id = itemIdOf(link)
			if id ~= nil and isEquippable(link) then
				items[#items + 1] = { itemId = id }
			end
		end
	end
	return items
end

function Export.equippedSlots()
	local slots = {}
	for _, entry in ipairs(Export.INVENTORY_SLOTS) do
		local id = itemIdOf(GetInventoryItemLink("player", entry.id))
		if id ~= nil then
			slots[#slots + 1] = { slot = entry.slot, itemId = id }
		end
	end
	return slots
end

--- Shared with Gear.lua, which requires Export for the container helpers
--- already and has no other slug vocabulary of its own to keep in.
function Export.slugify(name)
	if name == nil or name == "" then
		return nil
	end
	return (name:lower():gsub("[^%w]+", "-"):gsub("^%-+", ""):gsub("%-+$", ""))
end

--- A fallback for a race token `Export.RACE_SLUGS` does not carry -- Forever's
--- two new races, whose client tokens are not yet known (spike check 16).
--- PascalCase gets a hyphen inserted at each internal word boundary, so
--- something like `HighOrderSkyborne` degrades to a plausible slug instead
--- of to nothing; it is not run against the eight base races, which have
--- their real slugs in the map above.
local function pascalCaseSlugify(token)
	return (token:gsub("(%l)(%u)", "%1-%2"):gsub("(%d)(%u)", "%1-%2"):lower())
end

--- The site's race slug for the client's own race. The map first; the
--- fallback only for a token the map does not carry.
local function raceSlugOf()
	local raceToken = select(2, UnitRace("player"))
	if raceToken == nil then
		return nil
	end
	return Export.RACE_SLUGS[raceToken] or pascalCaseSlugify(raceToken)
end

--- `table.pack` does not exist on the client's Lua 5.1; this is its shape
--- (positional values plus `n`, the true count including trailing nils)
--- from a single call to `...`.
local function packn(...)
	return { n = select("#", ...), ... }
end

--- Every profession slot the client reports, primaries and secondaries
--- alike. `GetProfessions()` returns five positions and any of them may be
--- nil -- an unlearned primary, most commonly -- so this reads all five by
--- count rather than iterating the return values directly: `ipairs` on a
--- plain table literal stops at the first nil and silently drops every
--- profession after it.
function Export.professionSlugs()
	if GetProfessions == nil then
		return {}
	end
	local slugs = {}
	local slots = packn(GetProfessions())
	for position = 1, slots.n do
		local index = slots[position]
		if index ~= nil then
			local name = GetProfessionInfo and GetProfessionInfo(index)
			local slug = Export.slugify(name)
			if slug ~= nil then
				slugs[#slugs + 1] = slug
			end
		end
	end
	return slugs
end

--- The character's current guild, or nil when unguilded. GetGuildInfo returns nil for
--- an unguilded character (also the shape a login before guild data has loaded would
--- produce, per the WoW API's own documented behaviour -- both read the same way here:
--- no guild= section this export). Rank name (the second return) is player-chosen free
--- text and is never read; only the index is trustworthy (spike check 23, README.md).
function Export.guildInfo()
	local name, _, rankIndex = GetGuildInfo("player")
	if name == nil then
		return nil
	end
	return { name = name, rankIndex = rankIndex }
end

--- The export string, or nil and the reason.
function Export.string(data)
	-- The class slug comes from the locale-neutral class token (see
	-- Talents.playerClassSlug); slugifying the display name on a non-English
	-- client would produce a slug data.classes does not carry. Race gets the
	-- same treatment via raceSlugOf.
	local classSlug = Talents.playerClassSlug()
	local treeRanks = Talents.treeRanks(data, classSlug)
	if treeRanks == nil then
		return nil, string.format(L.codecUnknownClass, tostring(classSlug))
	end
	return Codec.encodeFS1({
		dataBuild = data.build,
		classSlug = classSlug,
		raceSlug = raceSlugOf(),
		treeRanks = treeRanks,
		gearSlots = Export.equippedSlots(),
		bags = Export.itemsInBags(Export.CARRIED_BAGS),
		bank = Export.itemsInBags(Export.BANK_BAGS),
		professions = Export.professionSlugs(),
		guild = Export.guildInfo(),
	})
end

--- Write the record the companion reads out of SavedVariables.
---
--- The shape is companion/internal/addon/addon.go's: any table carrying a
--- string `export` with `name`, `ruleset` and `region` beside it. `ruleset`
--- is the realm name until the beta shows an API that states Forever's own
--- ruleset (spike check 12); the companion falls back to `realm` for an
--- older addon, and this writes both so neither side has to guess.
function Export.save(data)
	local code, message = Export.string(data)
	if code == nil then
		return nil, message
	end
	-- UnitName("player") always succeeds in the real client; the "player"
	-- fallback only fires in a test double (wow_mock.lua does not stub
	-- UnitName) or a future client that drops the API, so the record is
	-- still written rather than erroring.
	local name = UnitName and UnitName("player") or "player"
	local realm = GetRealmName() or ""
	local region = Export.REGION_NAMES[GetCurrentRegion and GetCurrentRegion() or 0] or ""
	ForeverSixtyDB = ForeverSixtyDB or {}
	ForeverSixtyDB.characters = ForeverSixtyDB.characters or {}
	-- Stamped on the DB, not on the character record: the Export tab shows
	-- one "last saved" line, and a player with four characters wants the
	-- last time anything was written, not the last time this one was.
	ForeverSixtyDB.savedAt = date(Export.SAVED_AT_FORMAT)
	ForeverSixtyDB.characters[region .. "/" .. realm .. "/" .. name] = {
		name = name,
		class = UnitClass("player"),
		realm = realm,
		ruleset = realm,
		region = region,
		build = data.build,
		export = code,
	}
	return code
end

ns.Export = Export
return Export
