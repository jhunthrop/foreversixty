-- addon/ForeverSixty/surfaces/GroupPanel.lua
-- The group panel (spec section 4.8, task L5.2): for a party or a raid, before the first pull,
-- one row per member (class and spec, level, item level, gear gap, Forever Score and best
-- parse, missing enchants, addon present) with the weakest slot named under it, and the ready
-- check: who lacks a flask, food, a weapon imbue or a class buff, and who covers each need of
-- the encounter.
--
-- Ruling 1 (taint): as the guild tab (L5.1), the surface is an addon-owned "Forever Sixty"
-- button and an addon-owned flyout. Fix round 1 (D1): on 1.60 RaidFrame and FriendsFrame show
-- but have no rect (logical frames); the drawn Social window is SocialUIFrame. So both are
-- parented to the anchor, the first of SocialUIFrame, FriendsFrame and RaidFrame that shows with
-- a rect (looked up each time, guarded), and hide with the window: the button under its bottom
-- right, the flyout TOPLEFT at its TOPRIGHT (Widgets' shared placement and measure helpers).
-- "The Raid tab is open" is the Raid tab's own frame shown with the anchor placed
-- (GroupPanel.tabOpen). Fix round 3 (C1, hunter-36): on 1.60 the global RaidFrame is the legacy
-- raid frame, shown on every tab inside a RaidParentFrame hidden on every tab, so it says nothing;
-- the Raid tab's content is SocialUIFrame.RaidFrame, an unnamed frame keyed on the window (read,
-- never written). GroupPanel.raidTab finds it, else a visible RaidFrame on a client without one.
-- The tab is read at every refresh point (open, the group and inspect events,
-- PLAYER_ENTERING_WORLD, the button's own show script) and, fix round 2 (N1), on every tab switch:
-- a tiny addon-owned sentinel frame parented to the Raid tab's frame shows and hides with it, and
-- its own show and hide scripts read the tab again (no hook). The button rests on another tab and
-- wakes on the Raid tab; leaving the Raid tab closes the flyout (fix round 3, C2). The client
-- frames are only ever read (the addon's frames are made as their children);
-- scripts sit only on the addon's own frames; nothing polls. /fs group
-- toggles the flyout from anywhere: on the Social window (saying the panel shows on the Raid tab
-- when that tab is not open), or on UIParent at the chip's side on a client with no social frame.
-- Nothing here opens a client window or starts a ready check: the panel reads, it never acts.
--
-- Ruling 2: the rows are a pure join (GroupPanel.rows) of what the client says of each unit,
-- the inspect queue's record (Inspect, L2.5: spec by talents, the gear, enchants, item level),
-- the nightly record (Score.forCharacter) and the members' own addons (Comms.seen). The gap of a
-- member without the addon is measured on the inspected gear in their spec (Progress.gearPlan,
-- the rule the item tooltip, the guild row and the hello share); the player's own row reads this
-- client (Comms.ownPlan). Ruling 3: the ready check reads every unit's auras against the
-- shipped encounter needs (Data.encounterNeeds) of the instance's bosses. Ruling 4: a score floor
-- only tags rows. Ruling 5: no switch of its own; a /fs diag line; frames built on first open.
-- Fix round 4: the check claims only what it can (a flask from level 50, food from 35, an imbue
-- of a spec that fights with one, read on the player alone), reads class buffs and consumables
-- only outside a raid the data knows, says coverage as class presence, and the flyout clears the
-- Social window's side tabs (read from its children, never written).
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Compat = ns.Compat or require("Compat")
local Follow = ns.Follow or require("Follow")
local Gear = ns.Gear or require("Gear")
local Progress = ns.Progress or require("Progress")
local Score = ns.Score or require("Score")
local Comms = ns.Comms or require("Comms")
local Talents = ns.Talents or require("Talents")
local Inspect = ns.Inspect or require("Inspect")
local Prefs = ns.Prefs or require("Prefs")
local ItemTooltip = ns.ItemTooltip or require("ItemTooltip")
local PlayerTooltip = ns.PlayerTooltip or require("PlayerTooltip")
local Chip = ns.Chip or require("Chip")
local LootCouncil = ns.LootCouncil or require("LootCouncil")

local GroupPanel = {}

local S = Theme.SIZES

--- The frames' names, for the /run probes (README rows 63-69).
GroupPanel.BUTTON_NAME = "ForeverSixtyGroupButton"
GroupPanel.FLYOUT_NAME = "ForeverSixtyGroupFlyout"
--- The client's raid frame. Fix round 3 (C1): on 1.60 the global RaidFrame is the legacy raid
--- frame (in RaidParentFrame, hidden on every tab) and shows on every tab, so it is read only on a
--- client without the Social window's own Raid tab frame, and only while it is visible.
GroupPanel.FRAME = "RaidFrame"
--- The 1.60 Social window and the key its Raid tab content frame sits under (unnamed, so a key
--- read is the only stable way in): SocialUIFrame.RaidFrame, shown only on the Raid tab.
GroupPanel.SOCIAL = "SocialUIFrame"
GroupPanel.SOCIAL_RAID_KEY = "RaidFrame"
--- The frames the button and flyout may anchor to, in order (fix round 1, D1): the drawn Social
--- window (1.60), the classic social window, the raid frame itself. The first that shows with a
--- rect is the anchor.
GroupPanel.ANCHORS = { "SocialUIFrame", "FriendsFrame", "RaidFrame" }
--- How many rows a wheel notch moves.
GroupPanel.WHEEL_STEP = 3
--- The least time, in seconds, between two snapshot? asks on the group's channel.
GroupPanel.ASK_EVERY = 60
--- Seconds an aura change waits before the standing ready check is read again: a raid's
--- buffing fires UNIT_AURA by the hundred, and one read covers them all.
GroupPanel.AURA_DELAY = 0.5
--- The text a line is measured on: a capital, an ascender and descenders.
GroupPanel.LINE_SAMPLE = "Agjy"
--- The columns, left to right (the guild tab's: the sort key, the width, whether it sorts as a
--- number and is right-aligned under its label, which ends Theme.SIZES.guildSortReserve short
--- of the edge, the sort arrow's room). The widths and the gaps fit the flyout's inner width
--- less half the scroll track; Spec keeps room for "Marksmanship".
GroupPanel.COLUMNS = {
	{ key = "name", width = 76 },
	{ key = "level", width = 40, numeric = true, right = true },
	{ key = "spec", width = 82 },
	{ key = "ilvl", width = 36, numeric = true, right = true },
	{ key = "gap", width = 50, numeric = true, right = true },
	{ key = "score", width = 46, numeric = true, right = true },
	{ key = "enchants", width = 62, numeric = true, right = true },
	{ key = "addon", width = 36 },
}

--- The raid instances whose bosses Data.encounterNeeds lists, by the instance id GetInstanceInfo
--- answers eighth (any client language) and by the English name it answers first; each with
--- its bosses' npc ids. The data carries the needs per npc id only (contract 1), so the
--- instance-to-boss map is the addon's own.
GroupPanel.INSTANCE_IDS = { [409] = "Molten Core", [249] = "Onyxia's Lair", [469] = "Blackwing Lair" }
GroupPanel.BOSSES = {
	["Molten Core"] = { 12118, 11982, 12259, 12057, 12056, 12264, 12098, 11988, 12018, 11502 },
	["Onyxia's Lair"] = { 10184 },
	["Blackwing Lair"] = { 12435, 13020, 12017, 11983, 14601, 11981, 14020, 11583 },
}
--- Who can give or put up each need (the encounter's spell ids, contract 1), by class token:
--- the coverage list names the members of these classes.
GroupPanel.PROVIDERS = {
	[11551] = { "WARRIOR" }, [9885] = { "DRUID" }, [10157] = { "MAGE" }, [10938] = { "PRIEST" },
	[20217] = { "PALADIN" }, [11597] = { "WARRIOR" }, [11717] = { "WARLOCK" }, [1311680] = { "WARLOCK" },
	[9907] = { "DRUID" },
}
--- The group versions of a need (fix round 1, D2), keyed by the need's spell id: a member wearing
--- one has the need, matched by id or by the client's name for it. The ids are this build's
--- (data/builds/<build>/spellranks.json): Gift of the Wild 21849/21850, Prayer of Fortitude
--- 21562/21564/450086, Arcane Brilliance 23028, Greater Blessing of Kings 25898. Battle Shout has
--- none. A data carry: encounterNeeds should carry these (lane L1).
GroupPanel.VARIANTS = {
	[9885] = { 21849, 21850 },
	[10938] = { 21562, 21564, 450086 },
	[10157] = { 23028 },
	[20217] = { 25898 },
}
--- Who puts up a resistance aura or totem, per school.
GroupPanel.RESIST_PROVIDERS = {
	fire = { "PALADIN", "SHAMAN" }, frost = { "PALADIN", "SHAMAN" }, shadow = { "PALADIN", "PRIEST" },
	nature = { "SHAMAN", "HUNTER" },
}
--- A class buff asked of only the classes it does anything for: Arcane Intellect of mana users,
--- Battle Shout of those who hit with attack power. Any other buff, of everyone.
GroupPanel.BUFF_FOR = {
	[10157] = { PALADIN = true, HUNTER = true, PRIEST = true, SHAMAN = true, MAGE = true, WARLOCK = true,
		DRUID = true },
	[11551] = { WARRIOR = true, PALADIN = true, HUNTER = true, ROGUE = true, SHAMAN = true, DRUID = true },
}
--- The flasks' spell ids (Petrification, the Titans, Distilled Wisdom, Supreme Power, Chromatic
--- Resistance), matched by id or by the client's name for them; and the food buff's. Fix round 4
--- (A): Forever's four Natural flasks (Accuracy, Aggression, Precision, Swiftness), the use spells
--- of items 274273-274276 (data/builds/1.60.1.70291/simconsumes.json), as the Titans' 17626 is
--- item 13510's. Fix round 5 (D4): 1293559-1293562, the same-named Alchemy recipes (SkillLineAbility,
--- skill line 171), are kept too as a hedge: the repo has no SpellEffect data to prove which id the
--- worn aura has, and an id no aura has is harmless (the names match by `wears` either way).
GroupPanel.FLASKS = { 17624, 17626, 17627, 17628, 17629, 1293559, 1293560, 1293561, 1293562, 1293740, 1293741,
	1293742, 1293743 }
GroupPanel.WELL_FED = 19705
--- Fix round 4 (A): the level each consumable needs; a member below it is never asked for it.
--- L5.3 pre-item: the flask is asked at level 60, or inside a raid instance the data knows
--- (GroupPanel.INSTANCE_IDS: Molten Core, Onyxia's Lair, Blackwing Lair) at any level the client
--- can say, so a 52 in Scholomance is never flagged.
GroupPanel.CONSUMABLE_LEVELS = { flask = 60, food = 35 }
--- Fix round 4 (A): the specs that fight with a weapon buff (a temporary imbue): every warrior and
--- rogue spec, Retribution and Enhancement. Fix round 5 (D3): the data's spec slugs, matched
--- exactly (no prefix rule). L5.3 pre-item: no Feral (cat form has no weapon buff). Never a
--- hunter, a caster or a healer; never an unknown spec.
GroupPanel.IMBUE_SPECS = {
	["warrior-arms"] = true, ["warrior-fury"] = true, ["warrior-protection"] = true,
	["rogue-assassination"] = true, ["rogue-combat"] = true, ["rogue-subtlety"] = true,
	["paladin-retribution"] = true, ["shaman-enhancement"] = true,
}
--- The main hand's inventory slot, for the player's own weapon imbue.
GroupPanel.MAIN_HAND = 16
--- The off hand's slot, for the player's own off-hand weapon buff (L3.7 item 48).
GroupPanel.OFF_HAND = 17

--- The generated data table, set by Options.register().
GroupPanel.data = nil
--- The frames, once made, and the pooled member rows.
GroupPanel.button = nil
GroupPanel.flyout = nil
--- The sentinel on the Raid tab's frame (fix round 2, N1; fix round 3: GroupPanel.raidTab), and
--- the Raid tab as last read (N2: nil before).
GroupPanel.sentinel = nil
GroupPanel.tabWasOpen = nil
GroupPanel.rowFrames = {}
GroupPanel.headers = {}
GroupPanel.checkLines = {}
--- The Loot region's pooled lines (L5.3), and its frame.
GroupPanel.lootPool = {}
--- The sort a header click chose (nil: Gap high to low), and the first row shown.
GroupPanel.sort = nil
GroupPanel.offset = 0
--- How many times the flyout was drawn (a spec reads it), and the model last drawn.
GroupPanel.draws = 0
GroupPanel.last = nil
--- The last ready check (GroupPanel.runCheck's), or nil before one.
GroupPanel.result = nil
--- The game clock at the last snapshot? sent.
GroupPanel.lastAsk = nil
--- The measure every cell is fitted with, and a text line's measured height.
GroupPanel.measure = nil
GroupPanel.line = nil

-- The client reads, guarded ---------------------------------------------------------------

--- The client's function `name`, or nil where this client has none.
local function client(name)
	local fn = _G[name]
	return type(fn) == "function" and fn or nil
end

--- `fn` called with the arguments; nil where there is no function or it raised.
local function try(fn, ...)
	if type(fn) ~= "function" then
		return nil
	end
	local ok, a, b, c, d, e, f, g, h = pcall(fn, ...)
	if not ok then
		return nil
	end
	return a, b, c, d, e, f, g, h
end

--- The unit tokens of the group, the player included: raid1..N in a raid, player and party1..N
--- in a party, none out of a group. Pure.
function GroupPanel.units(inRaid, count)
	local units = {}
	count = tonumber(count) or 0
	if count <= 0 then
		return units
	end
	if inRaid then
		for index = 1, count do
			units[#units + 1] = "raid" .. index
		end
		return units
	end
	units[1] = "player"
	for index = 1, count - 1 do
		units[#units + 1] = "party" .. index
	end
	return units
end

--- The group's unit tokens as this client sees them now (GetNumGroupMembers, IsInRaid). Fix
--- round 1: IsInGroup first, so a player alone is never a group of one (a client may answer
--- GetNumGroupMembers 1 solo); without IsInGroup a party of one is no party either.
function GroupPanel.groupUnits()
	local inGroup = client("IsInGroup")
	if inGroup ~= nil and try(inGroup) ~= true then
		return {}
	end
	local inRaid = try(client("IsInRaid")) == true
	local count = try(client("GetNumGroupMembers"))
	if count == nil and not inRaid then
		local others = try(client("GetNumSubgroupMembers"))
		count = type(others) == "number" and others > 0 and others + 1 or 0
	end
	if not inRaid and (tonumber(count) or 0) <= 1 then
		return {}
	end
	return GroupPanel.units(inRaid, count)
end

--- Whether the player is in a party or a raid.
function GroupPanel.inGroup()
	return #GroupPanel.groupUnits() > 0
end

--- Whether `unit` is the player.
local function isPlayer(unit)
	return unit == "player" or try(client("UnitIsUnit"), unit, "player") == true
end

--- The channel the group's addon messages go on: RAID or PARTY; nil out of a group.
local function groupChannel()
	if try(client("IsInRaid")) == true then
		return "RAID"
	end
	return GroupPanel.inGroup() and "PARTY" or nil
end

--- The client's name for a spell (C_Spell, else GetSpellInfo), else the addon's own word for the
--- needs it knows (L.groupSpellNames); nil for anything else: an id is never shown (contract 1).
function GroupPanel.spellName(id)
	local spell = _G.C_Spell
	if type(spell) == "table" then
		local name = try(spell.GetSpellName, id)
		if type(name) == "string" and name ~= "" then
			return name
		end
		local info = try(spell.GetSpellInfo, id)
		if type(info) == "table" and type(info.name) == "string" and info.name ~= "" then
			return info.name
		end
	end
	local name = try(client("GetSpellInfo"), id)
	if type(name) == "string" and name ~= "" then
		return name
	end
	return L.groupSpellNames[id]
end

--- A unit's helpful auras as { ids = { [spellId] = true }, names = { [name] = true } }: through
--- AuraUtil.ForEachAura where the client has it, else UnitAura (UnitBuff on an older client);
--- nil where it can read none (the check then claims nothing for the unit).
function GroupPanel.aurasOf(unit)
	local auras = { ids = {}, names = {} }
	local function add(name, spellId)
		if type(name) == "string" and name ~= "" then
			auras.names[name] = true
		end
		if type(spellId) == "number" then
			auras.ids[spellId] = true
		end
	end
	local util = _G.AuraUtil
	if type(util) == "table" and type(util.ForEachAura) == "function" then
		local ok = pcall(util.ForEachAura, unit, "HELPFUL", nil, function(aura)
			if type(aura) == "table" then
				add(aura.name, aura.spellId)
			end
			return false
		end, true)
		if ok then
			return auras
		end
	end
	local reader = client("UnitAura") or client("UnitBuff")
	if reader == nil then
		return nil
	end
	for index = 1, 40 do
		local ok, name, _, _, _, _, _, _, _, _, spellId = pcall(reader, unit, index, "HELPFUL")
		if not ok or name == nil then
			break
		end
		add(name, spellId)
	end
	return auras
end

--- Whether the player's main hand carries a temporary imbue (GetWeaponEnchantInfo): true or
--- false with a main-hand item worn, nil where the client cannot say or nothing is worn there.
--- Only the player's own can be read; another player's imbue is never claimed missing.
function GroupPanel.ownImbue()
	local enchantInfo = client("GetWeaponEnchantInfo")
	if enchantInfo == nil or try(client("GetInventoryItemLink"), "player", GroupPanel.MAIN_HAND) == nil then
		return nil
	end
	return try(enchantInfo) == true
end

--- Whether the player's off hand carries a temporary buff (GetWeaponEnchantInfo's off-hand fields):
--- true or false with a WEAPON worn in the off hand (a shield or a held item has none to ask), nil
--- where the client cannot say or no weapon is worn there (L3.7 item 48).
function GroupPanel.ownOffImbue()
	local enchantInfo = client("GetWeaponEnchantInfo")
	local link = try(client("GetInventoryItemLink"), "player", GroupPanel.OFF_HAND)
	if enchantInfo == nil or link == nil then
		return nil
	end
	local _, _, _, equipLoc = Compat.itemInfoInstant(link)
	if equipLoc ~= "INVTYPE_WEAPON" and equipLoc ~= "INVTYPE_WEAPONOFFHAND" and equipLoc ~= "INVTYPE_WEAPONMAINHAND" then
		return nil
	end
	local info = { try(enchantInfo) }
	if info[1] == nil and info[5] == nil then
		return nil
	end
	return info[5] == true
end

--- The instance's key in GroupPanel.BOSSES from GetInstanceInfo's name, type and id (read from
--- the client when none is given); nil outdoors or in an instance the data has no bosses for.
function GroupPanel.instanceKey(name, instanceType, instanceId)
	if name == nil and instanceType == nil then
		local info = { try(client("GetInstanceInfo")) }
		name, instanceType, instanceId = info[1], info[2], info[8]
	end
	if instanceType == nil or instanceType == "none" then
		return nil
	end
	local key = GroupPanel.INSTANCE_IDS[tonumber(instanceId) or -1]
	if key ~= nil then
		return key
	end
	return GroupPanel.BOSSES[name] ~= nil and name or nil
end

-- The ready check, pure ----------------------------------------------------------------------

--- The needs to check against (contract 1's encounterNeeds): the union of the bosses' buffs and
--- debuffs of the instance `key` (GroupPanel.instanceKey's), in the data's order, and the school
--- they stack resistance for; outside one (or one the data has no needs for), the generic lists:
--- every boss's buffs and debuffs, by npc id, and no resistance. Pure.
function GroupPanel.needsFor(data, key)
	local all = type(data) == "table" and type(data.encounterNeeds) == "table" and data.encounterNeeds or {}
	local npcs = {}
	for _, npc in ipairs(key ~= nil and GroupPanel.BOSSES[key] or {}) do
		if type(all[npc]) == "table" then
			npcs[#npcs + 1] = npc
		end
	end
	local generic = #npcs == 0
	if generic then
		for npc in pairs(all) do
			npcs[#npcs + 1] = npc
		end
		table.sort(npcs)
	end
	local needs = { buffs = {}, debuffs = {}, generic = generic,
		instance = key ~= nil and GroupPanel.BOSSES[key] ~= nil and key or nil }
	local taken = {}
	for _, npc in ipairs(npcs) do
		local entry = all[npc]
		for _, kind in ipairs({ "buffs", "debuffs" }) do
			for _, id in ipairs(type(entry[kind]) == "table" and entry[kind] or {}) do
				if not taken[id] then
					taken[id] = true
					needs[kind][#needs[kind] + 1] = id
				end
			end
		end
		if not generic and needs.resist == nil and type(entry.resist) == "string" then
			needs.resist = entry.resist
		end
	end
	return needs
end

local function wears(auras, id, name)
	return auras.ids[id] == true or (name ~= nil and auras.names[name] == true)
end

local function wearsAny(auras, ids, nameOf)
	for _, id in ipairs(ids) do
		if wears(auras, id, nameOf(id)) then
			return true
		end
	end
	return false
end

--- A need worn: its own spell (any rank, by name) or one of its group versions (D2).
local function wearsNeed(auras, id, name, nameOf)
	return wears(auras, id, name) or wearsAny(auras, GroupPanel.VARIANTS[id] or {}, nameOf)
end

--- Whether `spec` fights with a weapon imbue (GroupPanel.IMBUE_SPECS, exactly); false for nil. Pure.
function GroupPanel.usesImbue(spec)
	return type(spec) == "string" and GroupPanel.IMBUE_SPECS[spec] == true
end

--- Whether a unit at `level` is old enough for a consumable of `need` levels: a level the client
--- cannot say (nil, 0) is not; a unit too high to show (-1) is.
local function reaches(level, need)
	level = tonumber(level)
	return level ~= nil and level ~= 0 and (level < 0 or level >= need)
end

--- The ready check (ruling 3) over `members` ({ short, classToken, level, spec (the one their row
--- shows), own (the player), connected, auras (aurasOf's; nil: unread), imbue (the player's:
--- GroupPanel.ownImbue's) }) against `needs` (GroupPanel.needsFor's), naming spells through `nameOf`
--- (GroupPanel.spellName); a group version of a buff (GroupPanel.VARIANTS) counts as the buff.
--- Answers { missing = { { name, classToken, own, what = { "flask", "Arcane Intellect", ... } } },
--- coverage = { { need, classes, by = { { name, classToken } } } }, generic, imbueOthers }: a member
--- lacks a flask (from level 50) or the food buff (from 35) (fix round 4, A:
--- GroupPanel.CONSUMABLE_LEVELS), the player's own imbue where their spec uses one, or a class buff
--- the group can give (a member of its class is here) and does them some good; a member offline
--- or whose auras cannot be read is never said to lack anything. Fix round 4 (B): with the generic
--- lists (no raid the data knows) there is no coverage. `imbueOthers`: a member other than the
--- player uses an imbue (the flyout says only the player's can be read). Pure.
function GroupPanel.check(members, needs, nameOf)
	local present = {}
	for _, member in ipairs(members) do
		if type(member.classToken) == "string" then
			present[member.classToken] = present[member.classToken] or {}
			table.insert(present[member.classToken], member.short)
		end
	end
	local function providers(classes)
		local found = {}
		for _, token in ipairs(classes or {}) do
			for _, name in ipairs(present[token] or {}) do
				found[#found + 1] = { name = name, classToken = token }
			end
		end
		return found
	end
	local result = { missing = {}, coverage = {}, generic = needs.generic == true or nil }
	local levels = GroupPanel.CONSUMABLE_LEVELS
	for _, member in ipairs(members) do
		if member.own ~= true and GroupPanel.usesImbue(member.spec) then
			result.imbueOthers = true
		end
		local auras = member.auras
		if member.connected ~= false and auras ~= nil then
			local what = {}
			-- How many needs applied to this member at all (L3.7 item 48): none, and a clean result
			-- reads "nothing missing at your level".
			local asked = 0
			-- L5.3 pre-item: level 60, or inside a raid the data knows (any level the client says).
			local flaskLevel = needs.instance ~= nil and 1 or levels.flask
			if reaches(member.level, flaskLevel) then
				asked = asked + 1
				if not wearsAny(auras, GroupPanel.FLASKS, nameOf) then
					what[#what + 1] = L.groupNeedFlask
				end
			end
			if reaches(member.level, levels.food) then
				asked = asked + 1
				if not wears(auras, GroupPanel.WELL_FED, nameOf(GroupPanel.WELL_FED)) and not auras.names[L.groupWellFed] then
					what[#what + 1] = L.groupNeedFood
				end
			end
			if member.imbue ~= nil and GroupPanel.usesImbue(member.spec) then
				asked = asked + 1
			end
			if member.imbue == false and GroupPanel.usesImbue(member.spec) then
				what[#what + 1] = L.groupNeedImbue
			elseif member.imbueOff == false and GroupPanel.usesImbue(member.spec) then
				-- The main hand is buffed, the off hand (a second weapon) is not (L3.7 item 48).
				what[#what + 1] = L.groupNeedImbueOff
			end
			for _, id in ipairs(needs.buffs or {}) do
				local name = nameOf(id)
				local forClass = GroupPanel.BUFF_FOR[id]
				if name ~= nil and (forClass == nil or forClass[member.classToken])
					and #providers(GroupPanel.PROVIDERS[id]) > 0 then
					asked = asked + 1
					if not wearsNeed(auras, id, name, nameOf) then
						what[#what + 1] = name
					end
				end
			end
			if member.own == true then
				result.asked = (result.asked or 0) + asked
				result.ownRead = true
			end
			if #what > 0 then
				result.missing[#result.missing + 1] = { name = member.short, classToken = member.classToken,
					own = member.own == true or nil, what = what }
			end
		end
	end
	-- Nothing applied to the player (every need gated by their level): a clean result says "at your level".
	result.levelGated = result.ownRead == true and (result.asked or 0) == 0 and #result.missing == 0 or nil
	if needs.generic then
		return result
	end
	-- A need is listed with the classes that give it; one no class is known to give says nothing.
	local function cover(need, classes)
		if need ~= nil and type(classes) == "table" and #classes > 0 then
			result.coverage[#result.coverage + 1] = { need = need, classes = classes, by = providers(classes) }
		end
	end
	for _, kind in ipairs({ "buffs", "debuffs" }) do
		for _, id in ipairs(needs[kind] or {}) do
			cover(nameOf(id), GroupPanel.PROVIDERS[id])
		end
	end
	if needs.resist ~= nil then
		cover(string.format(L.groupResist, L.groupSchools[needs.resist] or needs.resist),
			GroupPanel.RESIST_PROVIDERS[needs.resist])
	end
	return result
end

--- A member's name, in their class colour lifted for the panel where `coloured` (the flyout).
local function nameText(name, classToken, coloured)
	if coloured and type(classToken) == "string" then
		return "|cff" .. Theme.panelClassHex(classToken) .. name .. "|r"
	end
	return name
end

--- A class's word ("druid"), by token.
local function classWord(token)
	return L.groupClassWords[token] or tostring(token):lower()
end

--- "Thorgrim: missing flask, imbue", one per member lacking anything (fix round 4, E); the names
--- class-coloured where `coloured` (the flyout). Pure.
function GroupPanel.missingLines(result, coloured)
	local lines = {}
	for _, entry in ipairs(type(result) == "table" and result.missing or {}) do
		lines[#lines + 1] = string.format(L.groupMissing, nameText(entry.name, entry.classToken, coloured),
			table.concat(entry.what, L.groupMissingJoin))
	end
	return lines
end

--- The coverage lines as class presence (fix round 4, C): each need nobody of its classes covers,
--- "Faerie Fire: no druid" ("Fire resistance: no paladin or shaman"), first; then each covered
--- need, "Sunder Armor: warrior Thorgrim", the members' names class-coloured where `coloured`.
--- Each { text, covered }. Pure.
function GroupPanel.coverageEntries(result, coloured)
	local absent, covered = {}, {}
	for _, entry in ipairs(type(result) == "table" and result.coverage or {}) do
		if #entry.by == 0 then
			local words = {}
			for index, token in ipairs(entry.classes or {}) do
				words[index] = classWord(token)
			end
			absent[#absent + 1] = { text = string.format(L.groupCoverageNo, entry.need,
				table.concat(words, L.groupCoverageOr)), covered = false }
		else
			local names = {}
			for index, by in ipairs(entry.by) do
				names[index] = string.format(L.groupCoverageMember, classWord(by.classToken),
					nameText(by.name, by.classToken, coloured))
			end
			covered[#covered + 1] = { text = string.format(L.groupCoverageBy, entry.need,
				table.concat(names, L.groupCoverageJoin)), covered = true }
		end
	end
	for _, entry in ipairs(covered) do
		absent[#absent + 1] = entry
	end
	return absent
end

--- The coverage lines' plain text, in GroupPanel.coverageEntries' order. Pure.
function GroupPanel.coverageLines(result)
	local lines = {}
	for index, entry in ipairs(GroupPanel.coverageEntries(result, false)) do
		lines[index] = entry.text
	end
	return lines
end

--- What the player lacks when the check read them alone (`result.alone`, fix round 2): "flask,
--- food"; nil when the check read a group or found nothing. Pure.
function GroupPanel.ownLack(result)
	if type(result) ~= "table" or not result.alone then
		return nil
	end
	local entry = result.missing[1]
	if entry == nil then
		return nil
	end
	return table.concat(entry.what, L.groupMissingJoin)
end

--- The chip's one line: "Ready check: 5 missing (you, Thorgrim, +3)" (fix round 4, E: the player
--- first as "you" when among them, at most two names, then "+N"); for the player alone (fix round
--- 2; fix round 3, C3: the chip's short form) "Ready check: missing flask, food", and after a clean
--- dry run alone "Ready check: nothing missing" (fix round 4); nil with nobody missing otherwise.
--- Pure.
function GroupPanel.summary(result)
	if type(result) ~= "table" then
		return nil
	end
	local missing = type(result.missing) == "table" and result.missing or {}
	local dry = result.dry == true
	if #missing == 0 then
		if not (result.alone and dry) then
			return nil
		end
		if result.levelGated then
			return string.format(L.groupChipSummaryYouNoneLevelTest, GroupPanel.CONSUMABLE_LEVELS.food)
		end
		return L.groupChipSummaryYouNoneTest
	end
	local lack = GroupPanel.ownLack(result)
	if lack ~= nil then
		return string.format(dry and L.groupChipSummaryYouTest or L.groupChipSummaryYou, lack)
	end
	local names = {}
	for _, entry in ipairs(missing) do
		if entry.own then
			table.insert(names, 1, L.groupChipYou)
		else
			names[#names + 1] = entry.name
		end
	end
	local shown = { names[1], names[2] }
	if #names > 2 then
		shown[3] = string.format(L.groupChipMore, #names - 2)
	end
	return string.format(dry and L.groupChipSummaryTest or L.groupChipSummary, #missing,
		table.concat(shown, L.groupChipJoin))
end

-- The rows, pure ---------------------------------------------------------------------------

local NUMERIC = {}
for _, column in ipairs(GroupPanel.COLUMNS) do
	NUMERIC[column.key] = column.numeric == true
end

--- The best of a record's best parses (contract 5's best_parses, boss name -> percentile):
--- { boss, percentile }, the highest, ties by boss name; nil with none.
local function bestParse(record)
	local parses = type(record) == "table" and record.best_parses or nil
	local best
	for boss, percentile in pairs(type(parses) == "table" and parses or {}) do
		if type(boss) == "string" and type(percentile) == "number" and (best == nil or percentile > best.percentile
			or (percentile == best.percentile and boss < best.boss)) then
			best = { boss = boss, percentile = percentile }
		end
	end
	return best
end

local function validSpec(spec)
	return type(spec) == "string" and spec ~= "" and spec ~= "-" and spec or nil
end

--- One row from one member (GroupPanel.members' shape): the spec from their addon, else the
--- inspect, else the record; the gap from their addon (in its sender's unit), else measured on
--- the inspected gear; the item level and bare enchantable slots from the inspect; the score,
--- tier and best parse from the record; `live` when their addon has spoken (the player's own
--- always). `context`: { floor, paused, ratings, data }.
local function rowOf(member, context)
	local inspect, live, record, gear = member.inspect, member.live, member.record, member.gear
	local row = {
		token = member.unit, name = member.name, short = member.short or member.name, classToken = member.classToken,
		level = type(member.level) == "number" and member.level > 0 and member.level or nil,
		connected = member.connected ~= false, own = member.own == true, live = live ~= nil, inspected = inspect ~= nil,
		record = record,
	}
	row.spec = validSpec(live ~= nil and live.spec or nil) or validSpec(inspect ~= nil and inspect.spec or nil)
		or validSpec(record ~= nil and record.spec or nil)
	if GroupPanel.hasGear(inspect) then
		row.ilvl = tonumber(inspect.itemLevel)
		row.enchants = #Inspect.missingEnchants(inspect)
		row.bareSlots = Inspect.missingEnchants(inspect)
	end
	if live ~= nil and tonumber(live.gearGap) ~= nil then
		row.gap = tonumber(live.gearGap)
		row.cold = live.cold == true or nil
		row.gapExact = tonumber(live.gapExact)
		row.unit = live.unit or Comms.gapUnit(context.data or GroupPanel.data or ns.Data, row.spec, row.level)
		-- The weakest line is the inspect's: tagged where it measures another gap than the addon's.
		if gear ~= nil and gear.weakest ~= nil and (gear.unit ~= row.unit or type(gear.gap) ~= "number"
			or math.floor(gear.gap + 0.5) ~= math.floor(row.gap + 0.5)) then
			row.weakestInspect = true
		end
	elseif gear ~= nil then
		row.gap, row.unit, row.cold = gear.gap, gear.unit, gear.cold
	end
	if live ~= nil and type(live.buildName) == "string" and live.buildName ~= "" then
		row.buildName = live.buildName
	end
	row.weakest = gear ~= nil and gear.weakest or nil
	if type(record) == "table" and type(record.score) == "number" then
		row.score, row.tier = record.score, record.tier
		row.bestParse = bestParse(record)
	end
	row.noLogs = row.score == nil and context.ratings == true
		and not (type(row.level) == "number" and row.level < Score.LEVEL_FLOOR)
	local floor = tonumber(context.floor) or 0
	row.below = floor > 0 and row.score ~= nil and row.score < floor or false
	row.floor = floor
	row.paused = not row.inspected and not row.own and context.paused == true
	return row
end

local function valueOf(row, key)
	if key == "name" then
		return row.short:lower()
	elseif key == "spec" then
		return row.spec ~= nil and Follow.specName(row.spec):lower() or nil
	elseif key == "addon" then
		return row.live and 1 or nil
	end
	return row[key]
end

--- A gap per second (DPS, or a healer's HPS) as against points.
local function perSecond(unit)
	return unit == "dps" or unit == "hps"
end

--- The order for `sort` ({ key, descending }; nil: Gap high to low): the column's values, a
--- blank last both ways; Gap puts gaps per second (DPS or HPS) before those in points, each by number. Equal
--- values go by name, A to Z.
local function comparator(sort)
	local key = sort ~= nil and sort.key or "gap"
	local descending = sort == nil or sort.descending == true
	return function(left, right)
		local a, b = valueOf(left, key), valueOf(right, key)
		if key == "gap" and a ~= nil and b ~= nil and perSecond(left.unit) ~= perSecond(right.unit) then
			return perSecond(left.unit)
		end
		if a ~= b then
			if a == nil then
				return false
			elseif b == nil then
				return true
			end
			if descending then
				return a > b
			end
			return a < b
		end
		local leftName, rightName = left.short:lower(), right.short:lower()
		if leftName ~= rightName then
			return leftName < rightName
		end
		return (left.unit or "") < (right.unit or "")
	end
end

--- The rows (ruling 2): one per member (GroupPanel.members' answer), sorted by `sort`
--- (GroupPanel.nextSort's; nil Gap high to low). `context` (optional): { floor (the score floor,
--- 0 none), paused (the inspect queue held by combat), ratings (the record is readable: "no logs"
--- for a member it lacks), data (a live gap's unit) }. Answers the rows and { members,
--- inspected, live }. Pure.
function GroupPanel.rows(members, sort, context)
	context = type(context) == "table" and context or {}
	local rows = {}
	local counts = { members = 0, inspected = 0, live = 0 }
	for _, member in ipairs(type(members) == "table" and members or {}) do
		if type(member) == "table" and type(member.name) == "string" and member.name ~= "" then
			local row = rowOf(member, context)
			rows[#rows + 1] = row
			counts.members = counts.members + 1
			counts.inspected = counts.inspected + (row.inspected and 1 or 0)
			counts.live = counts.live + (row.live and 1 or 0)
		end
	end
	table.sort(rows, comparator(sort))
	return rows, counts
end

--- The sort after a click on `key`'s header: the same column turns round; another starts high
--- to low for a number, A to Z otherwise. The default is Gap high to low.
function GroupPanel.nextSort(sort, key)
	local current = sort or { key = "gap", descending = true }
	if current.key == key then
		return { key = key, descending = not current.descending }
	end
	return { key = key, descending = NUMERIC[key] == true }
end

local function gainText(weakest)
	if weakest.unit == "score" then
		return string.format(L.tooltipNextGainScore, weakest.gain)
	end
	return string.format(L.tooltipNextGain, weakest.gain, weakest.healing and L.tooltipUnitHps or L.tooltipUnitDps)
end

--- The weakest slot in words: "Chest: Worn Robe · next Robes of Arcana +12.0 DPS" (the gain as
--- the item tooltip's Next line prints it; the addon's " · " separator, no arrow glyph the client
--- font may lack), or "Chest: empty · next ..." for a slot worn bare; nil without one. Pure.
function GroupPanel.weakestText(weakest)
	if type(weakest) ~= "table" or type(weakest.gain) ~= "number" then
		return nil
	end
	local slot = L.gearSlotLabels[weakest.slot] or weakest.slot
	local target = weakest.targetName or L.tooltipNamePending
	if weakest.wornName == nil then
		return string.format(L.groupWeakestEmpty, slot, target, gainText(weakest))
	end
	return string.format(L.groupWeakest, slot, weakest.wornName, target, gainText(weakest))
end

--- The muted line under a row: "below 1200" under the floor, then the weakest slot (from the
--- inspect; "(inspect)" after it where the member's addon reports another gap), or "inspect
--- paused" for a member not yet inspected while combat holds the queue; "" for nothing. Pure.
function GroupPanel.subLine(row)
	local parts = {}
	if row.below then
		parts[#parts + 1] = string.format(L.groupBelow, row.floor)
	end
	local weakest = GroupPanel.weakestText(row.weakest)
	if weakest ~= nil then
		if row.weakestInspect then
			weakest = weakest .. " " .. L.groupWeakestInspect
		end
		parts[#parts + 1] = weakest
	elseif row.paused then
		parts[#parts + 1] = L.groupInspectPaused
	end
	return table.concat(parts, L.groupSubJoin)
end

local function whole(value)
	return string.format("%d", math.floor(value + 0.5))
end

local function plain(text)
	return { text = text, hex = Theme.HEX.body }
end

local function muted(text)
	return { text = text, hex = Theme.HEX.muted }
end

local function blank()
	return plain("")
end

--- A row's cells, each { text, hex }: the name in its class colour lifted for the panel, the
--- spec, level, item level, the gap in its unit ("14 DPS", muted at 0), the score in its tier's
--- panel colour ("no logs" muted where the readable record lacks the member), the enchants
--- ("3 missing" muted, or the client's ready-check tick), the addon ("live"); blank where there
--- is no value. Pure.
function GroupPanel.cells(row)
	local gap = blank()
	if type(row.gap) == "number" then
		local rounded = math.floor(row.gap + 0.5)
		local healer = row.unit == "hps"
		local format = healer and L.groupGapHps or row.unit == "dps" and L.groupGapDps or L.groupGapPts
		local text = string.format(format, rounded)
		-- Under 10 per second the player's own gap shows its decimal, rounded up so it never
		-- contradicts the integer (polish round B fix 1, item 5).
		local shown, decimal = Progress.shownGap(row.gapExact)
		if perSecond(row.unit) and shown ~= nil then
			if decimal then
				text = string.format(healer and L.groupGapHpsDecimal or L.groupGapDpsDecimal, shown)
			else
				text = string.format(format, shown)
			end
		end
		if row.cold and rounded > 0 then
			text = string.format(L.gapAtLeast, text)
		end
		gap = rounded == 0 and muted(text) or plain(text)
	end
	local cells = {
		name = { text = row.short, hex = row.classToken ~= nil and Theme.panelClassHex(row.classToken) or Theme.HEX.body },
		spec = row.spec ~= nil and plain(Follow.specName(row.spec)) or blank(),
		level = type(row.level) == "number" and plain(whole(row.level)) or blank(),
		ilvl = type(row.ilvl) == "number" and plain(whole(row.ilvl)) or blank(),
		gap = gap,
		addon = row.live and plain(L.groupAddonLive) or blank(),
		enchants = blank(),
	}
	if type(row.enchants) == "number" and not GroupPanel.belowEnchantLevel(row.level) then
		cells.enchants = row.enchants > 0 and muted(string.format(L.groupEnchantsMissing, row.enchants))
			or plain(Theme.inlineIcon(Theme.MEDIA.gearTick, S.groupTick))
		cells.enchants.icon = row.enchants == 0
	end
	if type(row.score) == "number" then
		local r, g, b = Score.tierColor(row.tier)
		local hex = Theme.panelTierHex(row.tier, r ~= nil and Theme.hexOf(r, g, b) or nil)
		cells.score = { text = whole(row.score), hex = hex or Theme.HEX.body }
	elseif row.noLogs then
		cells.score = muted(L.groupNoLogs)
	else
		cells.score = blank()
	end
	return cells
end

--- The level from which enchants are checked (L3.7 item 47), and whether `level` is below it.
GroupPanel.ENCHANT_LEVEL = 40

function GroupPanel.belowEnchantLevel(level)
	return type(level) == "number" and level > 0 and level < GroupPanel.ENCHANT_LEVEL
end

--- A column header's hover lines (L.groupHeaderHints), or nil for a column with none. The Enchants
--- header adds "Enchants are checked from level 40." when every member in `rows` is under 40 (the
--- column is blank then, and this says why).
function GroupPanel.headerHint(key, rows)
	local hint = L.groupHeaderHints[key]
	if hint == nil or key ~= "enchants" or type(rows) ~= "table" or #rows == 0 then
		return hint
	end
	for _, row in ipairs(rows) do
		if not GroupPanel.belowEnchantLevel(row.level) then
			return hint
		end
	end
	local lines = {}
	for index, line in ipairs(hint) do
		lines[index] = line
	end
	lines[#lines + 1] = L.groupEnchantsFrom
	return lines
end

--- A row's hover: the full name in its class colour, the score line, the best parse, the
--- build their addon named, the bare slots, "Offline".
function GroupPanel.rowHint(row)
	local cells = GroupPanel.cells(row)
	local lines = { "|cff" .. cells.name.hex .. row.name .. "|r" }
	local scoreLines = Score.lines(row.record)
	if scoreLines[1] ~= nil then
		lines[#lines + 1] = "|cff" .. cells.score.hex .. scoreLines[1] .. "|r"
	elseif row.noLogs then
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", L.scoreNoLogs)
	end
	if row.bestParse ~= nil then
		lines[#lines + 1] = Theme.inlineColor("tooltipText", string.format(L.groupBestParse,
			Chip.ordinal(row.bestParse.percentile), row.bestParse.boss))
	end
	if row.buildName ~= nil then
		lines[#lines + 1] = Theme.inlineColor("tooltipText", string.format(L.groupBuild, row.buildName))
	end
	if type(row.bareSlots) == "table" and #row.bareSlots > 0 and not GroupPanel.belowEnchantLevel(row.level) then
		local words = {}
		for index, slot in ipairs(row.bareSlots) do
			words[index] = L.gearSlotLabels[slot] or slot
		end
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", string.format(L.groupBareSlots,
			table.concat(words, L.groupMissingJoin)))
	end
	if not row.connected then
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", L.groupOffline)
	end
	if GroupPanel.belowEnchantLevel(row.level) then
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", L.groupEnchantsFrom)
	end
	return lines
end

--- The button's hover: "Forever Sixty" in gold, then what it holds and what a click does.
function GroupPanel.buttonHint()
	local lines = { Theme.inlineColor("gold", L.addonName) }
	for _, text in ipairs(L.groupPanelHint) do
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", text)
	end
	-- L5.3: while the master looter's Loot region shows, the drop is named (the latest, coloured).
	local drops = LootCouncil.shown()
	local drop = drops[#drops]
	if drop ~= nil then
		local name, hex = LootCouncil.display(drop.itemId, drop.link)
		local named = name or L.tooltipNamePending
		lines[#lines + 1] = string.format(L.lootButtonHint, hex ~= nil and name ~= nil
			and string.format("|cff%s%s|r", hex, named) or Theme.inlineColor("tooltipText", named))
	end
	return lines
end

-- The members, read --------------------------------------------------------------------------

--- An item's name: the link's own, else the client's (asked for once), else the data's.
local function itemName(link, itemId, data)
	local named = type(link) == "string" and link:match("%[(.-)%]") or nil
	if named ~= nil and named ~= "" then
		return named
	end
	local name = itemId ~= nil and Compat.itemName(itemId) or nil
	if name == nil and type(data) == "table" and type(data.bisNames) == "table" then
		name = data.bisNames[itemId]
	end
	return name
end

--- A gear plan's weakest slot as GroupPanel.weakestText reads it: the slot, the worn item's
--- name (nil for a slot worn bare), the target's, the gain in the plan's unit.
local function weakestOf(plan, equipped, data)
	local row = plan ~= nil and plan.weakest or nil
	if row == nil or type(row.gain) ~= "number" then
		return nil
	end
	local worn = equipped[row.slot]
	local current = plan.current or {}
	return {
		slot = row.slot, gain = row.gain, unit = plan.unit ~= "pts" and "dps" or "score",
		healing = current.healing or plan.unit == "hps",
		wornName = worn ~= nil and itemName(worn.link, worn.itemId, data) or nil,
		targetName = itemName(nil, row.itemId, data),
	}
end

--- Whether an inspect record carries gear to read (fix round 1): item links and an item level. A
--- record that came back with no links (itemLevel nil) is no gear: nothing is scored on it.
function GroupPanel.hasGear(inspect)
	return type(inspect) == "table" and type(inspect.items) == "table" and next(inspect.items) ~= nil
		and inspect.itemLevel ~= nil
end

--- An inspected member's gear plan in their spec (Progress.gearPlan over the inspected links,
--- the shipped weights at their level; the player's band and faction rule): { gap, unit,
--- weakest }; nil when nothing can be scored.
function GroupPanel.memberGear(data, spec, level, faction, items)
	if spec == nil or type(items) ~= "table" or next(items) == nil then
		return nil
	end
	local equipped = Gear.equippedFrom(items)
	local current = Gear.currentWeights(data, spec, nil, nil, level)
	local plan = Progress.gearPlan(data, spec, level, faction, equipped, current, ItemTooltip.isMissing)
	if plan == nil then
		return nil
	end
	plan.current = current
	return { gap = plan.gap, unit = plan.unit, cold = plan.cold, weakest = weakestOf(plan, equipped, data) }
end

--- The player's spec as their own row shows it (Comms.ownSpec: the build's, else the talents'
--- leading tree, Gear.specOf); the ready check asks the imbue in it too (fix round 4, A).
local function ownSpec(data)
	local ranks = type(data) == "table" and type(data.classes) == "table" and Talents.readRanks(data) or {}
	return Comms.ownSpec(data, Follow.build, Talents.playerClassSlug(), ranks)
end

--- Another member's spec as their row shows it (GroupPanel.rows: their addon's, else the
--- inspect's, else the record's); nil where none says (fix round 4, A: no imbue claimed then).
local function shownSpec(unit, name, realm, ownRealm, now)
	local live = PlayerTooltip.liveEntry(Comms.seen, name, (try(client("UnitName"), unit)), realm, ownRealm, now)
	local inspect = Inspect.get(Inspect.keyOf(unit) or name)
	local record = Score.forCharacter(name, realm)
	return validSpec(live ~= nil and live.spec or nil) or validSpec(inspect ~= nil and inspect.spec or nil)
		or validSpec(type(record) == "table" and record.spec or nil)
end

--- The player's own row's fields, read on this client (the chip's and the guild row's path:
--- Comms.ownSpec, Comms.ownPlan over what is worn, Follow.buildName): { live, inspect, gear }.
local function ownFields(data)
	local build = Follow.build
	local spec = ownSpec(data)
	local equipped = Gear.readEquipped()
	local ok, plan = pcall(Comms.ownPlan, data, spec, equipped)
	plan = ok and plan or nil
	local gear = plan ~= nil and { gap = plan.gap, unit = plan.unit, cold = plan.cold,
		weakest = weakestOf(plan, equipped, data) } or nil
	local live = {
		spec = spec, gearGap = plan ~= nil and plan.gap or nil, unit = plan ~= nil and plan.unit or nil,
		cold = plan ~= nil and plan.cold or nil, gapExact = plan ~= nil and plan.exact or nil,
		buildName = build ~= nil and type(data) == "table" and Follow.buildName(data, build) or nil,
	}
	return live, Inspect.readSelf(), gear
end

--- Every unit of the group as the rows take it: { unit, name, realm, short, classToken, level,
--- connected, own, inspect (Inspect.get's fresh record; the player's own read now), live
--- (Comms.seen's fresh entry; the player's own fields), record (Score.forCharacter), gear (the
--- plan measured on the inspected gear) }. Empty out of a group.
function GroupPanel.members(data, soloOwn)
	data = data or GroupPanel.data or ns.Data
	local members = {}
	local now = Theme.now()
	local _, ownRealm = Theme.playerNameRealm()
	local faction = Compat.playerFaction()
	local units = GroupPanel.groupUnits()
	-- Solo, the panel shows the player's own row (`soloOwn`): the group of one it will become.
	if #units == 0 and soloOwn then
		units = { "player" }
	end
	for _, unit in ipairs(units) do
		local name, realm = PlayerTooltip.nameOf(unit)
		if type(name) == "string" and name ~= "" then
			local _, classToken = try(client("UnitClass"), unit)
			local connected = try(client("UnitIsConnected"), unit)
			local member = {
				unit = unit, name = realm ~= nil and name .. "-" .. realm or name, realm = realm, short = name,
				classToken = type(classToken) == "string" and classToken or nil,
				level = tonumber((try(client("UnitLevel"), unit))),
				connected = connected ~= false, own = isPlayer(unit),
				record = Score.forCharacter(name, realm),
			}
			if member.own then
				member.live, member.inspect, member.gear = ownFields(data)
			else
				member.inspect = Inspect.get(Inspect.keyOf(unit) or name)
				member.live = PlayerTooltip.liveEntry(Comms.seen, name, (try(client("UnitName"), unit)), realm, ownRealm,
					now)
				if GroupPanel.hasGear(member.inspect) then
					local spec = validSpec(member.inspect.spec) or validSpec(member.live ~= nil and member.live.spec or nil)
					local ok, gear = pcall(GroupPanel.memberGear, data, spec, member.level, faction, member.inspect.items)
					member.gear = ok and gear or nil
				end
			end
			members[#members + 1] = member
		end
	end
	return members
end

--- The rows as drawn now: the group's members joined and sorted, with the floor, the inspect
--- pause and whether the record is readable. Answers the rows and the counts.
function GroupPanel.model(soloOwn)
	local data = GroupPanel.data or ns.Data
	return GroupPanel.rows(GroupPanel.members(data, soloOwn), GroupPanel.sort, {
		floor = Prefs.get("groupFloor"), paused = Inspect.isPaused(), ratings = Score.status().available == true,
		data = data,
	})
end

-- The ready check, read ------------------------------------------------------------------------

--- Whether the last check still stands at `now`: running, or within its time after it finished
--- (a dry run: Chip.RAID_CHECK_SECONDS).
local function checkStands(now)
	local result = GroupPanel.result
	if result == nil then
		return false
	end
	if result.open then
		return now - result.at <= Chip.RAID_CHECK_OPEN
	end
	return now <= (result.untilAt or result.at + Chip.RAID_CHECK_SECONDS)
end

--- Whether a check found nothing on a dry run of the player alone: the chip then says so (fix
--- round 4, E); a real check's clean result shows nothing.
local function cleanAlone(result)
	return result.alone == true and result.dry == true and #result.missing == 0
end

--- Reads every unit's auras (the player alone out of a group) and checks them against the
--- instance's needs; the chip gets the summary and the open flyout redraws. `again`: the same
--- check read anew (its time kept). Fix round 4 (A): each member's level and shown spec go in;
--- (B) `where` names the raid whose needs were read (nil: the generic lists). Answers the result.
local function runCheck(dry, again)
	local units = GroupPanel.groupUnits()
	local alone = #units == 0
	if alone then
		units = { "player" }
	end
	local data = GroupPanel.data or ns.Data
	local _, ownRealm = Theme.playerNameRealm()
	local now = Theme.now()
	local members = {}
	for _, unit in ipairs(units) do
		local name, realm = PlayerTooltip.nameOf(unit)
		if type(name) == "string" and name ~= "" then
			local _, classToken = try(client("UnitClass"), unit)
			local member = {
				short = name, classToken = classToken, connected = try(client("UnitIsConnected"), unit) ~= false,
				auras = GroupPanel.aurasOf(unit), level = tonumber((try(client("UnitLevel"), unit))),
			}
			if isPlayer(unit) then
				member.own = true
				member.imbue = GroupPanel.ownImbue()
				member.imbueOff = GroupPanel.ownOffImbue()
				local ok, spec = pcall(ownSpec, data)
				member.spec = ok and validSpec(spec) or nil
			else
				local ok, spec = pcall(shownSpec, unit, name, realm, ownRealm, now)
				member.spec = ok and spec or nil
			end
			members[#members + 1] = member
		end
	end
	local key = GroupPanel.instanceKey()
	local needs = GroupPanel.needsFor(data, key)
	local result = GroupPanel.check(members, needs, GroupPanel.spellName)
	-- Fix round 2: the player alone, so the summaries say what they lack (GroupPanel.summary).
	result.alone = alone or nil
	result.where = not needs.generic and key or nil
	local previous = GroupPanel.result
	if again and previous ~= nil then
		result.at, result.clock, result.open, result.untilAt, result.dry = previous.at, previous.clock, previous.open,
			previous.untilAt, previous.dry
		GroupPanel.result = result
		Chip.updateRaidCheck(GroupPanel.missingLines(result, true), GroupPanel.summary(result), cleanAlone(result))
	else
		result.at, result.clock, result.open, result.dry = Theme.now(), date("%H:%M"), not dry or nil, dry or nil
		GroupPanel.result = result
		Chip.setRaidCheck(GroupPanel.missingLines(result, true), GroupPanel.summary(result), not dry, cleanAlone(result))
	end
	return result
end

-- The Loot region, pure ------------------------------------------------------------------------------

--- One row of the master looter's list as drawn (L5.3): `row` is LootCouncil.rows'. The name in the
--- member's class colour, the value ("+6.3 DPS", with "(inspect)" or "(from the site)" after it, or "no
--- addon" / "can't score" / "waiting" muted, the first two with a `hint` for the hover), the worn item
--- muted ("empty" for a bare slot) and a gold "on their list" tag; `plain` is the same as one chat line.
--- `nameOf` names an item (Compat.itemName). Pure.
function GroupPanel.lootCell(row, nameOf, now)
	local cell = { kind = "row", name = row.short,
		nameHex = row.classToken ~= nil and Theme.panelClassHex(row.classToken) or Theme.HEX.body,
		value = "", valueHex = Theme.HEX.muted, worn = "", planned = row.planned == true }
	if row.kind == "noaddon" then
		cell.value, cell.hint = L.lootNoAddon, { L.lootNoAddonHint }
	elseif row.kind == "cantscore" then
		cell.value, cell.hint = L.lootCantScore, { L.lootCantScoreHint }
	elseif row.kind == "waiting" then
		cell.value = L.lootWaiting
	elseif row.alreadyWorn then
		cell.value = row.slotLabel ~= nil and string.format(L.lootAlreadyWornSlot, row.slotLabel) or L.lootAlreadyWorn
		cell.valueHex = Theme.HEX.body
		cell.planned = false
	else
		cell.value = LootCouncil.valueText(row.value, row.unit, row.healing)
		if row.kind == "inspect" then
			-- With the inspect's age where it is known (L3.7 item 51): "(inspect, 12 min ago)".
			local age = type(now) == "number" and type(row.inspectAt) == "number" and math.max(0, now - row.inspectAt) or nil
			local words
			if age ~= nil then
				words = age < 60 and L.lootAgeNow or (age < 3600 and string.format(L.lootAgeMin, math.floor(age / 60))
					or string.format(L.lootAgeHour, math.floor(age / 3600)))
			end
			cell.value = cell.value .. " " .. (words ~= nil and string.format(L.lootInspectAge, words) or L.lootInspect)
		elseif row.site then
			cell.value = cell.value .. " " .. L.lootSite
		end
		cell.valueHex = Theme.HEX.body
		if row.currentItemId == 0 then
			cell.worn = L.lootEmpty
		elseif row.currentItemId ~= nil then
			cell.worn = nameOf(row.currentItemId) or L.tooltipNamePending
		end
	end
	local parts = { cell.name, cell.value }
	if cell.worn ~= "" then
		parts[#parts + 1] = cell.worn
	end
	if cell.planned then
		parts[#parts + 1] = L.lootOnTheirList
	end
	cell.plain = table.concat(parts, L.groupSubJoin)
	return cell
end

--- The Loot region's entries (L5.3) for `drops` (LootCouncil.shown's): per drop a title ("Drop: <item>"
--- in the item's colour), then its rows (LootCouncil.rows over `members`, GroupPanel.members' shape),
--- at most `lootMaxRows` and then a muted "and N more". `resolvers` as LootCouncil.rows takes them;
--- `nameOf` names an item. Pure.
function GroupPanel.lootEntries(drops, members, now, resolvers, nameOf)
	local entries = {}
	for index, drop in ipairs(drops) do
		local name, hex = LootCouncil.display(drop.itemId, drop.link)
		if name ~= nil then
			drop.name, drop.hex = name, hex
		end
		local shown = name or L.tooltipNamePending
		if hex ~= nil and name ~= nil then
			shown = string.format("|cff%s%s|r", hex, name)
		end
		local title = string.format(L.lootDrop, shown)
		local left = LootCouncil.countdown(drop, now)
		if left ~= nil then
			title = string.format(L.lootCountdown, title, left)
		end
		entries[#entries + 1] = { kind = "title", text = title, plain = string.format(L.lootDrop, shown),
			gap = index > 1 or nil }
		local rows = LootCouncil.rows(drop, members, now, resolvers)
		for position, row in ipairs(rows) do
			if position > S.lootMaxRows then
				entries[#entries + 1] = { kind = "more", plain = string.format(L.lootMore, #rows - S.lootMaxRows),
					text = string.format(L.lootMore, #rows - S.lootMaxRows) }
				break
			end
			entries[#entries + 1] = GroupPanel.lootCell(row, nameOf, now)
		end
	end
	return entries
end

-- The heights, pure ----------------------------------------------------------------------------

--- A text line's height: the measured one, else `panelLine`.
function GroupPanel.lineHeight()
	return GroupPanel.line or S.panelLine
end

--- One member row: the cells' line (`guildRow` at least) and the muted line under it.
function GroupPanel.rowHeight()
	return math.max(S.guildRow, GroupPanel.lineHeight()) + GroupPanel.lineHeight()
end

--- From one member row's top to the next's (fix round 4, G): the row and `panelRowGap`.
function GroupPanel.rowStep()
	return GroupPanel.rowHeight() + S.panelRowGap
end

--- `count` member rows with the gap between them.
function GroupPanel.rowsHeight(count)
	return count * GroupPanel.rowHeight() + math.max(0, count - 1) * S.panelRowGap
end

--- Where the count line (or "Not in a group.") sits, from the flyout's top.
function GroupPanel.countsTop()
	return S.panelPadding + S.panelHeader + S.panelRowGap
end

--- The solo line under "Not in a group." (fix round 4, E): L.groupSoloHint's two sentences on one
--- line where `measure` (GroupPanel.measure) finds it within `room`, else a sentence a line, so
--- "/fs ready tries the check on you." is never cut. Pure.
function GroupPanel.soloHint(measure, room)
	local oneLine = table.concat(L.groupSoloHint, " ")
	if ItemTooltip.fitPlain(oneLine, measure, room) == oneLine then
		return { oneLine }
	end
	local lines = {}
	for index, text in ipairs(L.groupSoloHint) do
		lines[index] = text
	end
	return lines
end

--- Out of a group: the header, "Not in a group." and the solo line (`hintLines` lines; as many as
--- GroupPanel.soloHint gives now when not said).
function GroupPanel.soloHeight(hintLines)
	hintLines = hintLines or #GroupPanel.soloHint(GroupPanel.measure, GroupPanel.innerWidth())
	return GroupPanel.countsTop() + GroupPanel.lineHeight() * (1 + hintLines) + S.panelPadding
end

--- Where the column headers sit, and the rows under them.
function GroupPanel.headerTop()
	return GroupPanel.countsTop() + GroupPanel.lineHeight() + S.panelRowGap
end

function GroupPanel.rowsTop()
	return GroupPanel.headerTop() + S.guildHeaderRow
end

--- `entries` ({ text, color }) onto `lines`, at most `groupMissingLines`, the last then saying
--- "and N more".
local function capped(lines, entries)
	local room = S.groupMissingLines
	for index, entry in ipairs(entries) do
		if index < room or #entries == room then
			lines[#lines + 1] = entry
		elseif index == room then
			lines[#lines + 1] = { text = string.format(L.groupMissingMore, #entries - room + 1), color = "muted" }
		end
	end
end

--- The ready check's lines as drawn under the rows: its title naming the list it read (fix round
--- 4, B: "Ready check · Molten Core · 21:04", "Ready check · standard buffs · 21:04"), the missing
--- lines, names class-coloured (at most `groupMissingLines`, the last saying "and N more") or
--- "Nobody is missing anything.", once (fix round 4, A) where another member fights with an imbue
--- "Weapon imbues can only be read on yourself." muted, then the coverage title and lines (fix
--- round 4, C: "no <class>" in the body colour first, the covered muted; capped the same); each
--- { text, color }. Empty before a check. Pure.
function GroupPanel.checkText(result)
	if type(result) ~= "table" then
		return {}
	end
	local where = result.where or L.groupCheckStandard
	local lines = { { text = string.format(L.groupCheckTitle, where, result.clock or ""), color = "gold" } }
	local missing = {}
	for index, line in ipairs(GroupPanel.missingLines(result, true)) do
		missing[index] = { text = line, color = "body" }
	end
	if #missing == 0 then
		lines[#lines + 1] = { text = L.groupAllSet, color = "muted" }
	end
	capped(lines, missing)
	if result.imbueOthers and not result.alone then
		lines[#lines + 1] = { text = L.groupImbueSelfOnly, color = "muted" }
	end
	local coverage = {}
	for index, entry in ipairs(GroupPanel.coverageEntries(result, true)) do
		coverage[index] = { text = entry.text, color = entry.covered and "muted" or "body" }
	end
	if #coverage > 0 then
		lines[#lines + 1] = { text = L.groupCoverageTitle, color = "gold", gap = true }
		capped(lines, coverage)
	end
	return lines
end

--- The check region's height for `lines` (GroupPanel.checkText's), its gap above included; 0
--- for none.
function GroupPanel.checkHeight(lines)
	if #lines == 0 then
		return 0
	end
	local height = S.panelRegionGap
	for _, line in ipairs(lines) do
		height = height + GroupPanel.lineHeight() + (line.gap and S.panelRowGap or 0)
	end
	return height
end

--- The Loot region's height for `entries` (GroupPanel.lootEntries'), its gap above included; 0 for none.
function GroupPanel.lootHeight(entries)
	if #entries == 0 then
		return 0
	end
	local height = S.panelRegionGap
	for _, entry in ipairs(entries) do
		height = height + GroupPanel.lineHeight() + (entry.gap and S.panelRowGap or 0)
	end
	return height
end

--- How many member rows show at once under `limit` (the anchor's height: the Social window's on
--- 1.60; the flyout is never taller) with `extra` (the check region) below them:
--- `groupVisibleRows` at most, one at least; all of them where the client cannot say.
function GroupPanel.fitRows(limit, extra)
	if type(limit) ~= "number" or limit <= 0 then
		return S.groupVisibleRows
	end
	local fixed = GroupPanel.rowsTop() + (extra or 0) + S.panelPadding
	-- Fix round 4 (G): n rows take n steps less the last gap.
	local room = math.floor((limit - fixed + S.panelRowGap) / GroupPanel.rowStep())
	return math.max(1, math.min(S.groupVisibleRows, room))
end

-- The frames -------------------------------------------------------------------------------------

--- The client's global (legacy, on 1.60) raid frame, or nil (looked up each time, guarded).
local function pane()
	return Widgets.clientFrame(GroupPanel.FRAME)
end

--- The Social window's own Raid tab frame (SocialUIFrame.RaidFrame, unnamed), read through a
--- guarded key lookup: a table with IsShown and GetObjectType methods (fix round 4: a frame), else
--- nil. Read only, never written.
local function socialRaidTab()
	local social = Widgets.clientFrame(GroupPanel.SOCIAL)
	if social == nil then
		return nil
	end
	local ok, frame = pcall(function()
		return social[GroupPanel.SOCIAL_RAID_KEY]
	end)
	if ok and type(frame) == "table" and type(frame.IsShown) == "function"
		and type(frame.GetObjectType) == "function" then
		return frame
	end
	return nil
end

--- The frame that shows with the Raid tab (fix round 3, C1): SocialUIFrame.RaidFrame where the
--- client has it (1.60), else the global RaidFrame only while it is visible (a client without the
--- Social window's own frame), else nil. Looked up each time.
function GroupPanel.raidTab()
	local frame = socialRaidTab()
	if frame ~= nil then
		return frame
	end
	local raid = pane()
	if raid ~= nil and Widgets.isVisible(raid) then
		return raid
	end
	return nil
end

--- The anchor (fix round 1, D1): the first of GroupPanel.ANCHORS that shows with a rect
--- (Widgets.placed), looked up each time, guarded; nil while the window is closed.
function GroupPanel.anchor()
	for _, name in ipairs(GroupPanel.ANCHORS) do
		local frame = Widgets.clientFrame(name)
		if frame ~= nil and Widgets.placed(frame) then
			return frame
		end
	end
	return nil
end

--- Where the button and flyout live: the anchor, else (the window closed) the first of
--- GroupPanel.ANCHORS this client has, so they show when the window does; nil with none.
function GroupPanel.home()
	local anchor = GroupPanel.anchor()
	if anchor ~= nil then
		return anchor
	end
	for _, name in ipairs(GroupPanel.ANCHORS) do
		local frame = Widgets.clientFrame(name)
		if frame ~= nil then
			return frame
		end
	end
	return nil
end

--- Whether the Raid tab is the one open: the anchor placed and the Raid tab's frame
--- (GroupPanel.raidTab) shown.
function GroupPanel.tabOpen()
	return GroupPanel.anchor() ~= nil and Widgets.isShown(GroupPanel.raidTab())
end

--- One of the addon's own frames moved to `holder` when it lives elsewhere (SetParent on the
--- addon's frame only); `key` names where it was put. Fix round 2 (N4): the new home is recorded
--- before SetParent, which can run the frame's own show script (the move can make it visible) and
--- so come back here; that call finds the move made and moves nothing.
local function attach(frame, holder, key)
	if GroupPanel[key] ~= holder then
		GroupPanel[key] = holder
		frame:SetParent(holder)
	end
end

local function guarded(action, ...)
	return Widgets.guarded(GroupPanel, L.diagGroupPanelFailed, action, ...)
end

--- How far the Social window's side tabs reach past its right edge (fix round 4, D): the tab
--- buttons are unnamed children of SocialUIFrame, so the farthest GetRight of its children
--- (GetChildren, a read) that have a rect and end past the window's own right edge, the panel's
--- own button and flyout aside; nil where none does or the client cannot say. Read only, guarded.
--- Fix round 5 (D1): only a shown child whose left edge is at the window's right edge (less 1)
--- and at most `groupSideTabMaxWidth` wide is a side tab (the client's window art ends 1 past the
--- edge; a hidden child ends 407 past it), and the reach is clamped to that width.
local function sideTabsPast(frame)
	local own = Widgets.frameEdge(frame, "GetRight")
	if own == nil then
		return nil
	end
	local ok, children = pcall(function()
		return { frame:GetChildren() }
	end)
	if not ok or type(children) ~= "table" then
		return nil
	end
	local most = S.groupSideTabMaxWidth
	local farthest
	for _, child in ipairs(children) do
		if type(child) == "table" and not rawequal(child, GroupPanel.flyout) and not rawequal(child, GroupPanel.button)
			and Widgets.isShown(child) and Widgets.hasRect(child) then
			local left = Widgets.frameEdge(child, "GetLeft")
			local right = Widgets.frameEdge(child, "GetRight")
			if left ~= nil and right ~= nil and left >= own - 1 and right - left <= most and right > own
				and (farthest == nil or right > farthest) then
				farthest = right
			end
		end
	end
	return farthest ~= nil and math.min(farthest - own, most) or nil
end

--- The flyout on its home (GroupPanel.home: the anchor, else the closed window), TOPLEFT at its
--- TOPRIGHT (Widgets.placeFlyout); on the Social window (fix round 4, D) moved right past its side
--- tabs and clear of them by `guildSideClear`, as the guild tab's flyout is (GuildTab's
--- placeFlyout), or `groupSideFallback` across where they cannot be read; with no social frame
--- at all, on UIParent: its TOPRIGHT at the chip's TOPLEFT, else at the screen's top right.
function GroupPanel.placeFlyout(flyout)
	local frame = GroupPanel.home()
	if frame ~= nil then
		attach(flyout, frame, "flyoutHome")
		local x = S.panelOffsetX
		if rawequal(frame, Widgets.clientFrame(GroupPanel.SOCIAL)) then
			local past = sideTabsPast(frame)
			x = x + (past ~= nil and past + S.guildSideClear or S.groupSideFallback)
		end
		Widgets.placeFlyout(flyout, frame, nil, x, S.panelOffsetY)
		return
	end
	attach(flyout, UIParent, "flyoutHome")
	flyout:ClearAllPoints()
	if Chip.frame ~= nil then
		flyout:SetPoint("TOPRIGHT", Chip.frame, "TOPLEFT", -S.gap, 0)
	else
		flyout:SetPoint("TOPRIGHT", UIParent, "TOPRIGHT", Chip.FALLBACK_X, Chip.FALLBACK_Y)
	end
end

--- The "Forever Sixty" button under its home's bottom right (the Social window's on 1.60), the
--- panel tab's look (Widgets.edgedButton) and hover hint. Made the first time a refresh finds a
--- social frame (never in combat: it waits for its end); moved to the anchor when the window
--- shows on another. nil without one.
function GroupPanel.ensureButton()
	local frame = GroupPanel.home()
	if frame == nil then
		return nil
	end
	local button = GroupPanel.button
	if button == nil then
		if Theme.inCombat() then
			GroupPanel.deferred = true
			return nil
		end
		button = Widgets.edgedButton(frame, L.groupPanelLabel, function()
			GroupPanel.toggle()
		end, GroupPanel.BUTTON_NAME, S.panelTabWidth, S.panelTabHeight)
		GroupPanel.buttonHome = frame
		-- The addon's own button, a child of the window: it shows each time the window does.
		button:SetScript("OnShow", function(self)
			GroupPanel.onButtonShow(self)
		end)
		Widgets.withHint(button, GroupPanel.buttonHint)
		button:ClearAllPoints()
		button:SetPoint("TOPRIGHT", frame, "BOTTOMRIGHT", 0, 0)
		-- A capture helper for a /run (README rows 65-69): the rows as the flyout draws them now.
		button.rows = function()
			return GroupPanel.model()
		end
		GroupPanel.button = button
	elseif GroupPanel.buttonHome ~= frame then
		attach(button, frame, "buttonHome")
		button:ClearAllPoints()
		button:SetPoint("TOPRIGHT", frame, "BOTTOMRIGHT", 0, 0)
	end
	return button
end

--- The sentinel (fix round 2, N1): a tiny addon-owned frame parented to the Raid tab's frame
--- (GroupPanel.raidTab; fix round 3: SocialUIFrame.RaidFrame on 1.60). That frame shows and hides
--- with the Raid tab, and so does its child, so the sentinel's own show and hide scripts
--- (GroupPanel.onTab, GroupPanel.onTabLeft) hear every tab switch inside the open Social window,
--- which nothing else does. Scripts on the addon's own frame only, no hook. Made the first time a
--- refresh finds a Raid tab frame (never in combat: it waits for its end, as the button does);
--- moved when that frame changes (a SocialUIFrame made late, or the legacy RaidFrame giving way to
--- the window's own). Kept where it is while there is none; nil before the first.
function GroupPanel.ensureSentinel()
	local raid = GroupPanel.raidTab()
	if raid == nil then
		return GroupPanel.sentinel
	end
	local sentinel = GroupPanel.sentinel
	if sentinel == nil then
		if Theme.inCombat() then
			GroupPanel.deferred = true
			return nil
		end
		sentinel = CreateFrame("Frame", nil, raid)
		Theme.noteFrame("Group panel watcher", "the raid window first shown")
		sentinel:SetSize(1, 1)
		GroupPanel.sentinelHome = raid
		sentinel:SetScript("OnShow", function()
			GroupPanel.onTab()
		end)
		sentinel:SetScript("OnHide", function()
			GroupPanel.onTabLeft()
		end)
		GroupPanel.sentinel = sentinel
	elseif GroupPanel.sentinelHome ~= raid then
		attach(sentinel, raid, "sentinelHome")
	end
	return sentinel
end

--- The Raid tab read again (fix round 1, D1), at every refresh point and on every tab switch (the
--- sentinel, fix round 2): with the window open, the button is up on the Raid tab and rests on
--- another (no alpha, no mouse; its own shown flag kept, so its show script still runs when the
--- window next opens); with the window closed it is left as it is (it hides with the window). The
--- sentinel and the button made first where they can be. What it read is kept
--- (GroupPanel.tabWasOpen) so UNIT_AURA reads again only on a change (N2).
--- The Raid tab button's label: "Forever Sixty" in a group, "Forever Sixty · solo" alone (the
--- button is wider for it), set whenever the tab is read.
function GroupPanel.labelButton(button)
	local solo = not GroupPanel.inGroup()
	local label = solo and L.groupPanelSoloLabel or L.groupPanelLabel
	if button:GetText() ~= label then
		button:SetText(label)
	end
	button:SetWidth(solo and S.panelTabWidth + S.groupSoloTabExtra or S.panelTabWidth)
end

local function syncTab()
	GroupPanel.ensureSentinel()
	local button = GroupPanel.ensureButton()
	if button == nil then
		return
	end
	GroupPanel.labelButton(button)
	local open = GroupPanel.tabOpen()
	GroupPanel.tabWasOpen = open
	if GroupPanel.anchor() == nil then
		return
	end
	button:SetAlpha(open and 1 or 0)
	button:EnableMouse(open)
end

--- The inner width: the flyout less its padding.
function GroupPanel.innerWidth()
	return S.groupFlyoutWidth - S.panelPadding * 2
end

--- A column's x within a row.
function GroupPanel.columnX(index)
	local x = 0
	for before = 1, index - 1 do
		x = x + GroupPanel.COLUMNS[before].width + S.guildColumnGap
	end
	return x
end

--- The room a column's text has (the guild tab's rule).
function GroupPanel.textWidth(column)
	return column.right and column.width - S.guildSortReserve or column.width
end

local function fitted(text, room)
	return ItemTooltip.fitPlain(text, GroupPanel.measure, room)
end

local function sized(region, width)
	region:SetSize(width, GroupPanel.lineHeight())
	if type(region.SetWordWrap) == "function" then
		region:SetWordWrap(false)
	end
	return region
end

--- One pooled member row: a Button the width of the rows, a hover ground, a cell per column on
--- the top line and the muted line under it.
local function rowFrame(index)
	local row = GroupPanel.rowFrames[index]
	if row ~= nil then
		return row
	end
	local frame = CreateFrame("Button", nil, GroupPanel.content)
	Theme.noteFrame("Group panel row", "the group panel opened")
	frame:SetSize(GroupPanel.innerWidth(), GroupPanel.rowHeight())
	frame:EnableMouse(true)
	local ground = Theme.texture(frame, "BACKGROUND", "hover")
	ground:SetAllPoints(frame)
	ground:Hide()
	row = { frame = frame, ground = ground, cells = {} }
	local top = math.max(S.guildRow, GroupPanel.lineHeight())
	for column, spec in ipairs(GroupPanel.COLUMNS) do
		local cell = sized(Widgets.label(frame, "", "body", "small"), GroupPanel.textWidth(spec))
		cell:SetPoint("TOPLEFT", frame, "TOPLEFT", GroupPanel.columnX(column),
			-math.floor((top - GroupPanel.lineHeight()) / 2))
		if spec.right then
			cell:SetJustifyH("RIGHT")
		end
		row.cells[spec.key] = cell
	end
	row.sub = sized(Widgets.label(frame, "", "muted", "small"), GroupPanel.innerWidth())
	row.sub:SetPoint("TOPLEFT", frame, "TOPLEFT", 0, -top)
	frame:SetScript("OnEnter", function(self)
		ground:Show()
		if row.entry ~= nil then
			Theme.showLines(self, GroupPanel.rowHint(row.entry), "ANCHOR_RIGHT")
		end
	end)
	frame:SetScript("OnLeave", function()
		ground:Hide()
		Theme.hideTooltip()
	end)
	GroupPanel.rowFrames[index] = row
	return row
end

--- A header's label colour (Theme.CONTROL_STATES.header): the sorted column gold, another muted,
--- brightened under the cursor.
local function paintHeader(key)
	local header = GroupPanel.headers[key]
	local states = Theme.CONTROL_STATES.header
	local sorted = GroupPanel.sort ~= nil and GroupPanel.sort.key or "gap"
	local look = key == sorted and states.active or header.over and states.hoverLabel or states.label
	header.label:SetTextColor(Theme.rgb(Theme.HEX[look]))
	header.look = look
end

--- The column headers (the guild tab's look): an addon Button per column, its label, the client's
--- sort arrow, an inset ground while held, a hover where the column needs saying.
local function makeHeaders(flyout)
	local states = Theme.CONTROL_STATES.header
	for index, column in ipairs(GroupPanel.COLUMNS) do
		local key = column.key
		local button = CreateFrame("Button", nil, flyout)
		Theme.noteFrame("Group panel header", "the group panel opened")
		button:SetSize(column.width, S.guildHeaderRow)
		local ground = Theme.texture(button, "BACKGROUND", states.pressed)
		ground:SetAllPoints(button)
		ground:Hide()
		local label = Widgets.label(button, L.groupColumns[key], states.label, "small")
		if column.right then
			label:SetJustifyH("RIGHT")
			label:SetPoint("RIGHT", button, "RIGHT", -S.guildSortReserve, 0)
		else
			label:SetPoint("LEFT", button, "LEFT", 0, 0)
		end
		local arrow = Theme.tinted(button, "OVERLAY", Theme.MEDIA.sortArrow, states.active)
		arrow:SetSize(S.guildSortArrow, S.guildSortArrow)
		arrow:SetPoint("LEFT", label, "RIGHT", 1, 0)
		arrow:Hide()
		local hint = L.groupHeaderHints[key]
		local header = { button = button, label = label, arrow = arrow, ground = ground }
		GroupPanel.headers[key] = header
		button:SetScript("OnEnter", function(self)
			header.over = true
			paintHeader(key)
			-- Read as the pointer enters, from the rows drawn last (the Enchants hint depends on them).
			local shown = GroupPanel.headerHint(key, GroupPanel.last ~= nil and GroupPanel.last.rows or nil)
			if shown ~= nil then
				local lines = {}
				for line, text in ipairs(shown) do
					lines[line] = Theme.inlineColor(line == 1 and "tooltipText" or "tooltipMuted", text)
				end
				Theme.showLines(self, lines, "ANCHOR_TOP")
			end
		end)
		button:SetScript("OnLeave", function()
			header.over = false
			paintHeader(key)
			if hint ~= nil then
				Theme.hideTooltip()
			end
		end)
		Theme.pressable(button, label, function()
			ground:Show()
		end, function()
			ground:Hide()
		end)
		button:SetScript("OnClick", function()
			GroupPanel.setSort(key)
		end)
		button:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding + GroupPanel.columnX(index), -GroupPanel.headerTop())
	end
end

--- Hides the flyout; the button reads closed.
local function closeFlyout()
	if GroupPanel.flyout ~= nil then
		GroupPanel.flyout:Hide()
	end
	GroupPanel.syncButton()
end

--- The flyout, made on the first open: an addon panel `groupFlyoutWidth` wide parented to its home
--- (GroupPanel.home; UIParent on a client with no social frame), hidden until shown. It carries no script; its
--- close button, headers, rows and wheel carry their own.
function GroupPanel.ensureFlyout()
	if GroupPanel.flyout ~= nil then
		return GroupPanel.flyout
	end
	local home = GroupPanel.home() or UIParent
	local flyout = Widgets.panel(home, S.groupFlyoutWidth, S.panelPadding * 2, GroupPanel.FLYOUT_NAME)
	GroupPanel.flyoutHome = home
	flyout:EnableMouse(true)
	flyout:Hide()
	GroupPanel.flyout = flyout
	local header = Widgets.label(flyout, L.groupPanelHeader, "gold", "normal")
	header:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding, -S.panelPadding)
	local close = Widgets.closeButton(flyout, function()
		closeFlyout()
	end)
	close:SetPoint("TOPRIGHT", flyout, "TOPRIGHT", -S.gap, -S.gap)
	-- The measure helper's own string, never shown.
	local probe = Widgets.label(flyout, "", "muted", "small")
	probe:Hide()
	GroupPanel.measure = ItemTooltip.visibleMeasure(Theme.measurer(probe))
	GroupPanel.line = Widgets.measuredLine(probe, GroupPanel.measure, GroupPanel.LINE_SAMPLE, S.panelLine)
	local counts = sized(Widgets.label(flyout, "", "muted", "small"), GroupPanel.innerWidth())
	counts:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding, -GroupPanel.countsTop())
	GroupPanel.countsLine = counts
	-- Fix round 4 (E): the solo line under "Not in a group.", a label for each line it may take.
	GroupPanel.soloLines = {}
	for index = 1, #L.groupSoloHint do
		local line = sized(Widgets.label(flyout, "", "muted", "small"), GroupPanel.innerWidth())
		line:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding,
			-(GroupPanel.countsTop() + index * GroupPanel.lineHeight()))
		line:Hide()
		GroupPanel.soloLines[index] = line
	end
	makeHeaders(flyout)
	-- The rows: an addon-owned ScrollFrame clipping a content frame of pooled rows; the wheel
	-- moves the first row shown, three a notch.
	local holder = CreateFrame("ScrollFrame", nil, flyout)
	Theme.noteFrame("Group panel list", "the group panel opened")
	holder:SetSize(GroupPanel.innerWidth(), GroupPanel.rowHeight())
	holder:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding, -GroupPanel.rowsTop())
	local content = CreateFrame("Frame", nil, holder)
	content:SetSize(GroupPanel.innerWidth(), GroupPanel.rowHeight())
	if type(holder.SetScrollChild) == "function" then
		holder:SetScrollChild(content)
	end
	holder:EnableMouseWheel(true)
	holder:SetScript("OnMouseWheel", function(_, delta)
		GroupPanel.scroll(-delta * GroupPanel.WHEEL_STEP)
	end)
	GroupPanel.holder, GroupPanel.content = holder, content
	GroupPanel.bar = Widgets.scrollbar(holder, holder, -S.guildScrollbar / 2, S.guildScrollbar, "scrollTrack")
	-- The ready check under the rows: a 1 px rule, then its lines (pooled).
	local check = CreateFrame("Frame", nil, flyout)
	Theme.noteFrame("Group panel ready check", "the group panel opened")
	check:SetWidth(GroupPanel.innerWidth())
	local rule = Theme.texture(check, "ARTWORK", "border")
	rule:SetSize(GroupPanel.innerWidth(), S.border)
	rule:SetPoint("TOPLEFT", check, "TOPLEFT", 0, S.panelRegionGap / 2)
	check:Hide()
	GroupPanel.checkFrame = check
	-- The loot council helper's Loot region under it (L5.3): a rule, then pooled lines.
	local loot = CreateFrame("Frame", nil, flyout)
	Theme.noteFrame("Group panel loot", "the group panel opened")
	loot:SetWidth(GroupPanel.innerWidth())
	local lootRule = Theme.texture(loot, "ARTWORK", "border")
	lootRule:SetSize(GroupPanel.innerWidth(), S.border)
	lootRule:SetPoint("TOPLEFT", loot, "TOPLEFT", 0, S.panelRegionGap / 2)
	loot:Hide()
	GroupPanel.lootFrame = loot
	return flyout
end

--- The headers' look: the sorted column's label gold with its arrow, the others muted.
local function drawHeaders(shown)
	local key = GroupPanel.sort ~= nil and GroupPanel.sort.key or "gap"
	local descending = GroupPanel.sort == nil or GroupPanel.sort.descending == true
	for columnKey, header in pairs(GroupPanel.headers) do
		if shown then
			header.button:Show()
		else
			header.button:Hide()
		end
		paintHeader(columnKey)
		if columnKey == key then
			local coords = Theme.SORT_COORDS[descending and "descending" or "ascending"]
			header.arrow:SetTexCoord(coords[1], coords[2], coords[3], coords[4])
			header.arrow:Show()
		else
			header.arrow:Hide()
		end
	end
end

--- Draws the visible slice of the last model's rows from GroupPanel.offset (clamped): each row at
--- one alpha (Theme.ALPHA.dim offline), its cells, the muted line under it, the scrollbar.
function GroupPanel.renderRows()
	local model = GroupPanel.last
	if model == nil or GroupPanel.flyout == nil then
		return
	end
	local rows, visible = model.rows, model.visible
	local first, last, offset = Widgets.visibleRange(#rows, visible, GroupPanel.offset)
	GroupPanel.offset = offset
	-- Fix round 4 (G): `panelRowGap` between rows; the muted line stays tight under its row.
	local height = GroupPanel.rowStep()
	local drawn = 0
	for index = first, last do
		drawn = drawn + 1
		local entry = rows[index]
		local row = rowFrame(drawn)
		row.entry = entry
		row.frame:ClearAllPoints()
		row.frame:SetPoint("TOPLEFT", GroupPanel.content, "TOPLEFT", 0, -(drawn - 1) * height)
		row.frame:SetAlpha(entry.connected and 1 or Theme.ALPHA.dim)
		local cells = GroupPanel.cells(entry)
		for _, column in ipairs(GroupPanel.COLUMNS) do
			local cell, value = row.cells[column.key], cells[column.key]
			cell:SetText(value.icon and value.text or fitted(value.text, GroupPanel.textWidth(column)))
			cell:SetTextColor(Theme.rgb(value.hex))
		end
		row.sub:SetText(fitted(GroupPanel.subLine(entry), GroupPanel.innerWidth()))
		row.sub:SetTextColor(Theme.rgb(Theme.HEX.muted))
		row.frame:Show()
	end
	for index = drawn + 1, #GroupPanel.rowFrames do
		local row = GroupPanel.rowFrames[index]
		row.entry = nil
		row.frame:Hide()
	end
	local range = math.max(0, #rows - visible)
	GroupPanel.bar:Update(#rows > 0 and math.min(1, visible / #rows) or 1, range > 0 and offset / range or 0)
end

--- One pooled line of the check region.
local function checkLine(index)
	local line = GroupPanel.checkLines[index]
	if line == nil then
		line = sized(Widgets.label(GroupPanel.checkFrame, "", "body", "small"), GroupPanel.innerWidth())
		GroupPanel.checkLines[index] = line
	end
	return line
end

--- Lays the check region out under the rows at `y` from `lines`; answers the y below it.
local function drawCheck(flyout, lines, y)
	local region = GroupPanel.checkFrame
	if #lines == 0 then
		region:Hide()
		for _, line in ipairs(GroupPanel.checkLines) do
			line:Hide()
		end
		return y
	end
	y = y + S.panelRegionGap
	region:ClearAllPoints()
	region:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding, -y)
	local inner = 0
	for index, entry in ipairs(lines) do
		if entry.gap then
			inner = inner + S.panelRowGap
		end
		local line = checkLine(index)
		line:ClearAllPoints()
		line:SetPoint("TOPLEFT", region, "TOPLEFT", 0, -inner)
		line:SetText(fitted(entry.text, GroupPanel.innerWidth()))
		line:SetTextColor(Theme.rgb(Theme.HEX[entry.color]))
		line:Show()
		inner = inner + GroupPanel.lineHeight()
	end
	for index = #lines + 1, #GroupPanel.checkLines do
		GroupPanel.checkLines[index]:Hide()
	end
	region:SetHeight(inner)
	region:Show()
	return y + inner
end

--- One pooled line of the Loot region: a title label and the four cells of a member's row.
local function lootLine(index)
	local line = GroupPanel.lootPool[index]
	if line ~= nil then
		return line
	end
	local parent, width = GroupPanel.lootFrame, GroupPanel.innerWidth()
	local gap = S.lootColumnGap
	local wornX = S.lootNameWidth + S.lootValueWidth + gap * 2
	local wornWidth = width - wornX - S.lootPlannedWidth - gap
	line = {
		title = sized(Widgets.label(parent, "", "body", "small"), width),
		name = sized(Widgets.label(parent, "", "body", "small"), S.lootNameWidth),
		value = sized(Widgets.label(parent, "", "body", "small"), S.lootValueWidth),
		worn = sized(Widgets.label(parent, "", "muted", "small"), wornWidth),
		planned = sized(Widgets.label(parent, "", "gold", "small"), S.lootPlannedWidth),
		x = { 0, S.lootNameWidth + gap, wornX, width - S.lootPlannedWidth },
		wornWidth = wornWidth,
	}
	-- The hover over a row that has one ("no addon" / "can't score"): an addon frame laid over the row.
	local hit = CreateFrame("Frame", nil, parent)
	Theme.noteFrame("Group panel hover area", "the group panel opened")
	hit:EnableMouse(true)
	hit:SetSize(width, GroupPanel.lineHeight())
	hit:SetScript("OnEnter", function(self)
		if line.hint ~= nil then
			Theme.showLines(self, line.hint, "ANCHOR_RIGHT")
		end
	end)
	hit:SetScript("OnLeave", Theme.hideTooltip)
	hit:Hide()
	line.hit = hit
	GroupPanel.lootPool[index] = line
	return line
end

local LOOT_KEYS = { "title", "name", "value", "worn", "planned" }

--- Lays the Loot region out under the check at `y` from `entries`; answers the y below it.
local function drawLoot(flyout, entries, y)
	local region = GroupPanel.lootFrame
	if #entries == 0 then
		region:Hide()
		for _, line in ipairs(GroupPanel.lootPool) do
			for _, key in ipairs(LOOT_KEYS) do
				line[key]:Hide()
			end
			line.hit:Hide()
			line.hint = nil
		end
		return y
	end
	y = y + S.panelRegionGap
	region:ClearAllPoints()
	region:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding, -y)
	local inner = 0
	local function put(label, text, hex, x, room)
		label:ClearAllPoints()
		label:SetPoint("TOPLEFT", region, "TOPLEFT", x, -inner)
		label:SetText(fitted(text, room))
		label:SetTextColor(Theme.rgb(hex))
		label:Show()
	end
	for index, entry in ipairs(entries) do
		if entry.gap then
			inner = inner + S.panelRowGap
		end
		local line = lootLine(index)
		for _, key in ipairs(LOOT_KEYS) do
			line[key]:Hide()
		end
		line.hint = entry.hint
		if entry.hint ~= nil then
			line.hit:ClearAllPoints()
			line.hit:SetPoint("TOPLEFT", region, "TOPLEFT", 0, -inner)
			line.hit:Show()
		else
			line.hit:Hide()
		end
		if entry.kind == "row" then
			put(line.name, entry.name, entry.nameHex, line.x[1], S.lootNameWidth)
			put(line.value, entry.value, entry.valueHex, line.x[2], S.lootValueWidth)
			if entry.worn ~= "" then
				put(line.worn, entry.worn, Theme.HEX.muted, line.x[3], line.wornWidth)
			end
			if entry.planned then
				put(line.planned, L.lootOnTheirList, Theme.HEX.gold, line.x[4], S.lootPlannedWidth)
			end
		else
			put(line.title, entry.text, entry.kind == "more" and Theme.HEX.muted or Theme.HEX.body, 0,
				GroupPanel.innerWidth())
		end
		inner = inner + GroupPanel.lineHeight()
	end
	for index = #entries + 1, #GroupPanel.lootPool do
		for _, key in ipairs(LOOT_KEYS) do
			GroupPanel.lootPool[index][key]:Hide()
		end
		GroupPanel.lootPool[index].hit:Hide()
		GroupPanel.lootPool[index].hint = nil
	end
	region:SetHeight(inner)
	region:Show()
	return y + inner
end

--- The solo line (fix round 4, E): `lines` (GroupPanel.soloHint's) drawn muted, the rest hidden;
--- none in a group.
local function drawSolo(lines, top)
	for index, label in ipairs(GroupPanel.soloLines) do
		local text = lines[index]
		if text ~= nil then
			-- Under the player's own row (`top`), else under the count line.
			label:ClearAllPoints()
			local y = top ~= nil and top + (index - 1) * GroupPanel.lineHeight()
				or GroupPanel.countsTop() + index * GroupPanel.lineHeight()
			label:SetPoint("TOPLEFT", GroupPanel.flyout, "TOPLEFT", S.panelPadding, -y)
			label:SetText(fitted(text, GroupPanel.innerWidth()))
			label:SetTextColor(Theme.rgb(Theme.HEX.muted))
			label:Show()
		else
			label:Hide()
		end
	end
end

--- The Loot region's entries now (L5.3): the drops the master looter asked about (LootCouncil.shown), each
--- member's answer, the player's own value computed here, and a silent member scored from their inspect
--- once LootCouncil.WAIT seconds passed. Empty while no drop shows.
function GroupPanel.lootNow()
	local drops = LootCouncil.shown()
	if #drops == 0 then
		return {}
	end
	local data = GroupPanel.data or ns.Data
	local faction = Compat.playerFaction()
	local resolvers = {
		-- A healer's answer reads in HPS: their spec, from the hello or the inspect, in the data's table.
		healing = function(member)
			local spec = validSpec(member.live ~= nil and member.live.spec or nil)
				or validSpec(member.inspect ~= nil and member.inspect.spec or nil)
			return spec ~= nil and Comms.gapUnit(data, spec) == "hps" or nil
		end,
		own = function(drop)
			if drop.ownResult == nil then
				local ok, result = pcall(LootCouncil.ownValue, data, drop.itemId, true)
				drop.ownResult = ok and result or nil
			end
			return drop.ownResult
		end,
		inspect = function(member, drop)
			if not GroupPanel.hasGear(member.inspect) then
				return nil
			end
			local spec = validSpec(member.inspect.spec) or validSpec(member.live ~= nil and member.live.spec or nil)
			if spec == nil then
				return nil
			end
			drop.inspected = drop.inspected or {}
			local kept = drop.inspected[member.name]
			if kept ~= nil and kept.at == member.inspect.at then
				return kept.result
			end
			local ok, result = pcall(LootCouncil.valueFor, data, drop.itemId, {
				spec = spec, level = member.level, classToken = member.classToken, faction = faction,
				equipped = Gear.equippedFrom(member.inspect.items),
				current = Gear.currentWeights(data, spec, nil, nil, member.level), isMissing = ItemTooltip.isMissing,
			})
			result = ok and result or nil
			if result ~= nil then
				result.at = member.inspect.at
			end
			drop.inspected[member.name] = { at = member.inspect.at, result = result }
			return result
		end,
	}
	return GroupPanel.lootEntries(drops, GroupPanel.members(data), Theme.now(), resolvers, Compat.itemName)
end

--- The Loot region as chat lines (the chip's click).
function GroupPanel.lootChatLines()
	local lines = {}
	for _, entry in ipairs(GroupPanel.lootNow()) do
		lines[#lines + 1] = entry.plain
	end
	return lines
end

--- Anything the Loot region shows changed (LootCouncil.onChange): the chip's line follows and an open
--- flyout redraws. The flyout is never opened here: nothing acts on the player's behalf.
function GroupPanel.onLoot()
	guarded(function()
		Chip.setLoot(LootCouncil.chipText())
	end)
	GroupPanel.redraw()
end

--- Draws the flyout: out of a group "Not in a group." and the muted solo line (fix round 4, E); in
--- one the count line, the headers, the rows (as many as fit under the anchor's height) and the
--- last ready check; sized to them.
function GroupPanel.draw()
	local flyout = GroupPanel.flyout
	if flyout == nil then
		return nil
	end
	GroupPanel.stale = nil
	GroupPanel.draws = GroupPanel.draws + 1
	if not GroupPanel.inGroup() then
		local ownRows = GroupPanel.model(true)
		local hint = GroupPanel.soloHint(GroupPanel.measure, GroupPanel.innerWidth())
		GroupPanel.countsLine:SetText(fitted(L.groupNotInGroup, GroupPanel.innerWidth()))
		-- L5.3 pre-item: the line itself in the body colour; the hint lines stay muted.
		GroupPanel.countsLine:SetTextColor(Theme.rgb(Theme.HEX.body))
		drawCheck(flyout, {}, 0)
		drawLoot(flyout, {}, 0)
		if #ownRows > 0 then
			-- The player's own row, with "Invite someone and their row appears here." under it.
			GroupPanel.last = { rows = ownRows, counts = { members = 1, inspected = 0, live = 0 }, visible = 1,
				solo = true }
			drawHeaders(true)
			local holder = GroupPanel.holder
			holder:SetHeight(GroupPanel.rowsHeight(1))
			GroupPanel.content:SetHeight(GroupPanel.rowsHeight(1))
			holder:Show()
			GroupPanel.renderRows()
			local below = GroupPanel.rowsTop() + GroupPanel.rowsHeight(1) + S.panelRowGap
			drawSolo(hint, below)
			flyout:SetHeight(below + GroupPanel.lineHeight() * #hint + S.panelPadding)
			return GroupPanel.last
		end
		GroupPanel.last = { rows = {}, counts = { members = 0, inspected = 0, live = 0 }, visible = 0, solo = true }
		drawSolo(hint)
		drawHeaders(false)
		GroupPanel.holder:Hide()
		for _, row in ipairs(GroupPanel.rowFrames) do
			row.entry = nil
			row.frame:Hide()
		end
		flyout:SetHeight(GroupPanel.soloHeight(#hint))
		return GroupPanel.last
	end
	drawSolo({})
	local rows, counts = GroupPanel.model()
	local lines = GroupPanel.checkText(GroupPanel.result)
	local loot = GroupPanel.lootNow()
	local visible = GroupPanel.fitRows(Widgets.frameHeight(GroupPanel.anchor()),
		GroupPanel.checkHeight(lines) + GroupPanel.lootHeight(loot))
	GroupPanel.last = { rows = rows, counts = counts, visible = visible }
	GroupPanel.countsLine:SetText(fitted(string.format(L.groupCounts, counts.members, counts.inspected, counts.live),
		GroupPanel.innerWidth()))
	GroupPanel.countsLine:SetTextColor(Theme.rgb(Theme.HEX.muted))
	drawHeaders(true)
	local shownRows = math.max(1, math.min(visible, #rows))
	local holder = GroupPanel.holder
	holder:SetHeight(GroupPanel.rowsHeight(shownRows))
	GroupPanel.content:SetHeight(GroupPanel.rowsHeight(shownRows))
	holder:Show()
	GroupPanel.renderRows()
	local y = drawCheck(flyout, lines, GroupPanel.rowsTop() + GroupPanel.rowsHeight(shownRows))
	y = drawLoot(flyout, loot, y)
	flyout:SetHeight(y + S.panelPadding)
	return GroupPanel.last
end

--- Redraws an open flyout while it is on screen; one the closed window hides is marked stale and
--- redrawn when the window shows again (the button's show script).
local function redraw()
	local flyout = GroupPanel.flyout
	if flyout == nil or not flyout:IsShown() then
		return
	end
	if Widgets.isVisible(flyout) then
		GroupPanel.draw()
	else
		GroupPanel.stale = true
	end
end

--- An open flyout redrawn, guarded.
function GroupPanel.redraw()
	guarded(redraw)
end

--- Queues an inspect for every member not freshly inspected (Inspect, L2.5: one at a time, two
--- seconds apart, held in combat); the player is never inspected.
local function inspectGroup()
	for _, unit in ipairs(GroupPanel.groupUnits()) do
		if not isPlayer(unit) then
			Inspect.request(unit)
		end
	end
end

--- On each open: the inspects queued and one snapshot? on the group's channel (through Comms,
--- at most once a minute, never while "Share my build with my guild" is off), so members with
--- the addon answer and their rows go live.
function GroupPanel.askLive()
	inspectGroup()
	local channel = groupChannel()
	local now = Theme.now()
	local due = GroupPanel.lastAsk == nil or now - GroupPanel.lastAsk >= GroupPanel.ASK_EVERY
	if channel ~= nil and Comms.sharing() and due then
		if Comms.send("snapshot?", {}, channel) then
			GroupPanel.lastAsk = now
		end
	end
end

--- The button reads open while the flyout is.
function GroupPanel.syncButton()
	local button, flyout = GroupPanel.button, GroupPanel.flyout
	if button ~= nil then
		Widgets.setEdgedOpen(button, flyout ~= nil and flyout:IsShown())
	end
end

--- Opens the flyout: the Raid tab read, placed, shown, the inspects and snapshots asked for, drawn
--- (or drawn when the window next shows).
local function openFlyout(flyout)
	syncTab()
	GroupPanel.placeFlyout(flyout)
	flyout:Show()
	GroupPanel.syncButton()
	GroupPanel.askLive()
	if Widgets.isVisible(flyout) then
		GroupPanel.draw()
	else
		GroupPanel.stale = true
	end
end

--- The button's click: the flyout opens (made on the first click) or closes.
function GroupPanel.toggle()
	return guarded(function()
		local flyout = GroupPanel.ensureFlyout()
		if flyout:IsShown() then
			closeFlyout()
			return false
		end
		openFlyout(flyout)
		return true
	end)
end

--- /fs group (ruling 1): the same toggle from anywhere. With the Raid tab not open (the window
--- closed, or another tab) the flyout opens on the window and a line says where the panel shows
--- (nothing here opens a client window); never with the Raid tab open, nor on a client with no
--- social frame (the flyout is then at the chip's side). Fix round 4 (F): with the window drawn on
--- another tab the flyout it opened is closed again (the line says where the panel shows); with
--- the window closed it stays open to show when the window does. The lines to print.
function GroupPanel.command()
	-- Solo there is nothing to toggle: the line says where the panel is and when it fills.
	if not GroupPanel.inGroup() then
		return { L.groupSoloCommand }
	end
	local lines = guarded(function()
		local flyout = GroupPanel.ensureFlyout()
		if flyout:IsShown() then
			closeFlyout()
			return {}
		end
		openFlyout(flyout)
		if GroupPanel.home() ~= nil and not GroupPanel.tabOpen() then
			if GroupPanel.anchor() ~= nil then
				closeFlyout()
			end
			return { L.groupOpensOnRaidTab }
		end
		return {}
	end)
	return lines or {}
end

--- A header's click: sorts by that column, or turns it round; back to the first row.
function GroupPanel.setSort(key)
	guarded(function()
		GroupPanel.sort = GroupPanel.nextSort(GroupPanel.sort, key)
		GroupPanel.offset = 0
		redraw()
	end)
end

--- The wheel: `delta` rows down (up when negative), clamped.
function GroupPanel.scroll(delta)
	guarded(function()
		GroupPanel.offset = GroupPanel.offset + delta
		GroupPanel.renderRows()
	end)
end

--- Every surface's refresh (Options; PLAYER_ENTERING_WORLD among them): the button made when a
--- social frame exists, the Raid tab read, an open flyout redrawn.
function GroupPanel.refresh(data)
	if data ~= nil then
		GroupPanel.data = data
	end
	guarded(function()
		syncTab()
		redraw()
	end)
end

--- GROUP_ROSTER_UPDATE: an invited member is inspected on arrival (ruling 4); the Raid tab read;
--- an open flyout redraws.
function GroupPanel.onRosterUpdate()
	guarded(function()
		inspectGroup()
		syncTab()
		redraw()
	end)
end

--- INSPECT_READY brought a record: the Raid tab read; an open flyout redraws.
function GroupPanel.onInspectReady(record)
	if record ~= nil then
		guarded(function()
			syncTab()
			redraw()
		end)
	end
end

--- Combat began or ended: the inspect queue pauses or carries on, and the rows waiting on it
--- say so. Ended: the button a combat lockdown held back is made.
function GroupPanel.onCombat(inCombat)
	if not inCombat and GroupPanel.deferred then
		GroupPanel.deferred = nil
		GroupPanel.refresh()
		return
	end
	guarded(redraw)
end

--- A hello or snapshot! arrived (Comms has put it in Comms.seen): an open flyout redraws.
function GroupPanel.onLive()
	guarded(redraw)
end

--- READY_CHECK (ruling 3): every unit's auras read against the instance's needs; the chip gets
--- its line until READY_CHECK_FINISHED and 30 s more; an open flyout shows the lists. Answers
--- the result (nil when the panel has turned itself off).
function GroupPanel.onReadyCheck()
	return guarded(function()
		local result = runCheck(false)
		redraw()
		return result
	end)
end

--- READY_CHECK_FINISHED: the check stands 30 s more (Chip.finishRaidCheck), then goes from the chip.
function GroupPanel.onReadyCheckFinished()
	guarded(function()
		local result = GroupPanel.result
		if result ~= nil and result.open then
			result.open = nil
			result.untilAt = Theme.now() + Chip.RAID_CHECK_AFTER
		end
		Chip.finishRaidCheck()
		redraw()
	end)
end

--- UNIT_AURA (throttled): while a check stands, a group unit's aura change has its auras read
--- again GroupPanel.AURA_DELAY later, once for every change in between. Fix round 2 (N2): the
--- Raid tab is read there, after the throttle, and only when it differs from the last read
--- (GroupPanel.tabWasOpen); the sentinel hears the tab switches themselves.
function GroupPanel.onAura(unit)
	if type(unit) ~= "string" or not (unit == "player" or unit:match("^party%d+$") or unit:match("^raid%d+$")) then
		return
	end
	if GroupPanel.auraPending or not checkStands(Theme.now()) then
		return
	end
	GroupPanel.auraPending = true
	local function again()
		GroupPanel.auraPending = nil
		guarded(function()
			-- Fix round 3: a tab never read (nil) and one read closed (false) are the same.
			if (GroupPanel.tabOpen() or false) ~= (GroupPanel.tabWasOpen or false) then
				syncTab()
			end
			if checkStands(Theme.now()) then
				runCheck(GroupPanel.result.dry, true)
				redraw()
			end
		end)
	end
	if not Theme.after(GroupPanel.AURA_DELAY, again) then
		again()
	end
end

--- /fs ready: a dry run of the check (the player alone out of a group), told in chat; the chip
--- shows its line for Chip.RAID_CHECK_SECONDS. A panel that has turned itself off says so and
--- claims no clean check (fix round 1, D3). Fix round 4 (E, B): alone, one line ("Ready check
--- (only you see this): you are missing flask, food." or "...: nothing missing."); in a group a
--- header naming the list read ("Ready check (only you see this), Molten Core:", "..., standard
--- buffs:"), then a line a member or "Nobody is missing anything.". The lines to print.
function GroupPanel.handleReady()
	local result = guarded(function()
		local result = runCheck(true)
		redraw()
		return result
	end)
	if result == nil then
		return { L.groupPanelOff }
	end
	if result.alone then
		local lack = GroupPanel.ownLack(result)
		return { lack ~= nil and string.format(L.groupReadyYou, lack)
			or (result.levelGated and string.format(L.groupReadyYouNoneLevel, GroupPanel.CONSUMABLE_LEVELS.food)
				or L.groupReadyYouNone) }
	end
	local lines = { string.format(L.groupReadyDry, result.where or L.groupCheckStandard) }
	local missing = GroupPanel.missingLines(result)
	if #missing == 0 then
		lines[2] = L.groupReadyNone
	end
	for _, line in ipairs(missing) do
		lines[#lines + 1] = line
	end
	return lines
end

--- /fs floor N (ruling 4): the group's leader (or a player not yet in a group, setting it for the
--- listing) sets the score floor (Prefs.groupFloor); /fs floor alone clears it. On a client with
--- no UnitIsGroupLeader, the player is the leader only out of a group (fix round 1). The panel
--- only tags; it never declines or whispers. The lines to print.
function GroupPanel.handleFloor(rest)
	local typed = tostring(rest or ""):match("^%s*(.-)%s*$")
	local leader = not GroupPanel.inGroup()
		or (client("UnitIsGroupLeader") ~= nil and Theme.isGroupLeader("player"))
	if not leader then
		return { L.groupFloorLeader }
	end
	if typed == "" then
		Prefs.setFlag("groupFloor", 0)
		GroupPanel.redraw()
		return { L.groupFloorCleared }
	end
	local floor = typed:match("^%d+$") and tonumber(typed) or nil
	if floor == nil or floor <= 0 then
		return { L.groupFloorUsage }
	end
	Prefs.setFlag("groupFloor", floor)
	GroupPanel.redraw()
	return { string.format(L.groupFloorSet, floor) }
end

--- The Raid tab read, the button and an open flyout moved to the anchor (a SocialUIFrame the
--- client made after login among them, N3), and the flyout redrawn once when it went stale while
--- the window was closed. Fix round 4 (F): an open flyout found on a window drawn on another tab
--- (/fs group with the window closed, which then opened on Friends) is closed.
local function reread()
	syncTab()
	local flyout = GroupPanel.flyout
	if flyout == nil or not flyout:IsShown() then
		return
	end
	GroupPanel.placeFlyout(flyout)
	if GroupPanel.anchor() ~= nil and not GroupPanel.tabOpen() then
		closeFlyout()
		return
	end
	if GroupPanel.stale and Widgets.isVisible(flyout) then
		GroupPanel.draw()
	end
end

--- The button's show script (it shows with the window): the Raid tab and the window read again.
function GroupPanel.onButtonShow()
	guarded(reread)
end

--- The sentinel's show script (fix round 2, N1: the Raid tab's frame shown, the Raid tab opened or
--- the window opened on it): the Raid tab and the window read again.
function GroupPanel.onTab()
	guarded(reread)
end

--- The sentinel's hide script (the Raid tab left, or the window closed): the Raid tab and the
--- window read again; fix round 3 (C2): an open flyout closes when the window is still drawn on
--- another tab (never when the window closed: it shows with the window again, as /fs group with
--- the window closed has it). It is not opened again on the way back.
function GroupPanel.onTabLeft()
	guarded(function()
		reread()
		local flyout = GroupPanel.flyout
		if flyout and flyout:IsShown() and GroupPanel.anchor() ~= nil and not GroupPanel.tabOpen() then
			closeFlyout()
		end
	end)
end

--- Listens for the members' hellos and snapshot! answers, once (Options.register).
function GroupPanel.register()
	if GroupPanel.registered then
		return
	end
	GroupPanel.registered = true
	LootCouncil.onChange = GroupPanel.onLoot
	Chip.lootLines = GroupPanel.lootChatLines
	Comms.on("hello", function()
		GroupPanel.onLive()
	end)
	Comms.on("snapshot!", function()
		GroupPanel.onLive()
	end)
end

--- /fs diag's line (ruling 5): "Group panel: open, 5 members (3 inspected, 2 live), ready check
--- 21:04: 2 missing" (open: the flyout shows and is visible), counted now as the flyout counts.
--- Fix round 2: out of a group a check that read the player alone says what they lack: "ready
--- check 19:43: you are missing flask, food" (fix round 3: "nothing missing" when clean); fix
--- round 4 (E): out of a group the line reads "not in a group". Fix round 3: a check that read the
--- player alone is stale once they are in a group, and reads "none"; fix round 4 (F): so is a
--- group's check once the player is alone.
function GroupPanel.diagLine()
	local open = Widgets.isShown(GroupPanel.flyout) and Widgets.isVisible(GroupPanel.flyout)
	local state = open and L.diagPanelOpen or L.diagPanelClosed
	local result = GroupPanel.result
	local inGroup = GroupPanel.inGroup()
	local check = L.diagGroupNone
	if result ~= nil and result.alone then
		if not inGroup then
			local lack = GroupPanel.ownLack(result)
			check = lack ~= nil and string.format(L.diagGroupCheckYou, result.clock or "", lack)
				or string.format(L.diagGroupCheckYouNone, result.clock or "")
		end
	elseif result ~= nil and inGroup then
		check = string.format(L.diagGroupCheck, result.clock or "", #result.missing)
	end
	if not inGroup then
		return string.format(L.diagGroupPanelYou, state, check)
	end
	local counts = { members = 0, inspected = 0, live = 0 }
	local ok, _, found = pcall(GroupPanel.model)
	if ok and type(found) == "table" then
		counts = found
	end
	return string.format(L.diagGroupPanel, state, counts.members, counts.inspected, counts.live, check)
end

ns.GroupPanel = GroupPanel
return GroupPanel
