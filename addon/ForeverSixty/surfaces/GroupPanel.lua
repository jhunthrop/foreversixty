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
-- "The Raid tab is open" is RaidFrame:IsShown() with the anchor placed (GroupPanel.tabOpen), read
-- at every refresh point (open, the group and inspect events, PLAYER_ENTERING_WORLD, the
-- button's own show script) and, fix round 2 (N1), on every tab switch: a tiny addon-owned
-- sentinel frame parented to RaidFrame shows and hides with the Raid tab, and its own show and
-- hide scripts read the tab again (no hook). The button rests on another tab and wakes on the
-- Raid tab. The client frames are only ever read (the addon's frames are made as their children);
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

local GroupPanel = {}

local S = Theme.SIZES

--- The frames' names, for the /run probes (README rows 63-69).
GroupPanel.BUTTON_NAME = "ForeverSixtyGroupButton"
GroupPanel.FLYOUT_NAME = "ForeverSixtyGroupFlyout"
--- The client's raid frame: the child FriendsFrame shows on its Raid tab (1.60: no RaidFrameTab).
--- On 1.60 it shows without a rect; it says only whether the Raid tab is the one open.
GroupPanel.FRAME = "RaidFrame"
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
	{ key = "spec", width = 82 },
	{ key = "level", width = 40, numeric = true, right = true },
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
--- Resistance), matched by id or by the client's name for them; and the food buff's.
GroupPanel.FLASKS = { 17624, 17626, 17627, 17628, 17629 }
GroupPanel.WELL_FED = 19705
--- The main hand's inventory slot, for the player's own weapon imbue.
GroupPanel.MAIN_HAND = 16

--- The generated data table, set by Options.register().
GroupPanel.data = nil
--- The frames, once made, and the pooled member rows.
GroupPanel.button = nil
GroupPanel.flyout = nil
--- The sentinel on RaidFrame (fix round 2, N1), and the Raid tab as last read (N2: nil before).
GroupPanel.sentinel = nil
GroupPanel.tabWasOpen = nil
GroupPanel.rowFrames = {}
GroupPanel.headers = {}
GroupPanel.checkLines = {}
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
	local needs = { buffs = {}, debuffs = {}, generic = generic }
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

--- The ready check (ruling 3) over `members` ({ short, classToken, connected, auras (aurasOf's;
--- nil: unread), imbue (the player's: GroupPanel.ownImbue's) }) against `needs`
--- (GroupPanel.needsFor's), naming spells through `nameOf` (GroupPanel.spellName); a group version
--- of a buff (GroupPanel.VARIANTS) counts as the buff. Answers
--- { missing = { { name, what = { "flask", "Arcane Intellect", ... } } }, coverage = { { need,
--- by = { names } } } }: a member lacks a flask, the food buff, the player's own imbue, or a
--- class buff the group can give (a member of its class is here) and does them some good; a
--- member offline or whose auras cannot be read is never said to lack anything. Pure.
function GroupPanel.check(members, needs, nameOf)
	local present = {}
	for _, member in ipairs(members) do
		if type(member.classToken) == "string" then
			present[member.classToken] = present[member.classToken] or {}
			table.insert(present[member.classToken], member.short)
		end
	end
	local function providers(classes)
		local names = {}
		for _, token in ipairs(classes or {}) do
			for _, name in ipairs(present[token] or {}) do
				names[#names + 1] = name
			end
		end
		return names
	end
	local result = { missing = {}, coverage = {} }
	for _, member in ipairs(members) do
		local auras = member.auras
		if member.connected ~= false and auras ~= nil then
			local what = {}
			if not wearsAny(auras, GroupPanel.FLASKS, nameOf) then
				what[#what + 1] = L.groupNeedFlask
			end
			if not wears(auras, GroupPanel.WELL_FED, nameOf(GroupPanel.WELL_FED)) and not auras.names[L.groupWellFed] then
				what[#what + 1] = L.groupNeedFood
			end
			if member.imbue == false then
				what[#what + 1] = L.groupNeedImbue
			end
			for _, id in ipairs(needs.buffs or {}) do
				local name = nameOf(id)
				local forClass = GroupPanel.BUFF_FOR[id]
				if name ~= nil and (forClass == nil or forClass[member.classToken]) and #providers(GroupPanel.PROVIDERS[id]) > 0
					and not wearsNeed(auras, id, name, nameOf) then
					what[#what + 1] = name
				end
			end
			if #what > 0 then
				result.missing[#result.missing + 1] = { name = member.short, what = what }
			end
		end
	end
	for _, kind in ipairs({ "buffs", "debuffs" }) do
		for _, id in ipairs(needs[kind] or {}) do
			local name = nameOf(id)
			if name ~= nil then
				result.coverage[#result.coverage + 1] = { need = name, by = providers(GroupPanel.PROVIDERS[id]) }
			end
		end
	end
	if needs.resist ~= nil then
		result.coverage[#result.coverage + 1] = {
			need = string.format(L.groupResist, L.groupSchools[needs.resist] or needs.resist),
			by = providers(GroupPanel.RESIST_PROVIDERS[needs.resist]),
		}
	end
	return result
end

--- "Thorgrim: flask, imbue", one per member lacking anything. Pure.
function GroupPanel.missingLines(result)
	local lines = {}
	for _, entry in ipairs(type(result) == "table" and result.missing or {}) do
		lines[#lines + 1] = string.format(L.groupMissing, entry.name, table.concat(entry.what, L.groupMissingJoin))
	end
	return lines
end

--- "Sunder Armor: covered by Thorgrim", "Faerie Fire: nobody", one per need. Pure.
function GroupPanel.coverageLines(result)
	local lines = {}
	for _, entry in ipairs(type(result) == "table" and result.coverage or {}) do
		if #entry.by > 0 then
			lines[#lines + 1] = string.format(L.groupCoverage, entry.need, table.concat(entry.by, L.groupCoverageJoin))
		else
			lines[#lines + 1] = string.format(L.groupCoverageNobody, entry.need)
		end
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

--- The chip's one line: "Ready check: 2 missing (Name, Name)"; for the player alone (fix round 2)
--- "Ready check: you are missing flask, food"; nil with nobody missing. Pure.
function GroupPanel.summary(result)
	local missing = type(result) == "table" and result.missing or {}
	if #missing == 0 then
		return nil
	end
	local lack = GroupPanel.ownLack(result)
	if lack ~= nil then
		return string.format(L.groupChipSummaryYou, lack)
	end
	local names = {}
	for index, entry in ipairs(missing) do
		names[index] = entry.name
	end
	return string.format(L.groupChipSummary, #missing, table.concat(names, L.groupChipJoin))
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
		row.unit = live.unit or Comms.gapUnit(context.data or GroupPanel.data or ns.Data, row.spec, row.level)
		-- The weakest line is the inspect's: tagged where it measures another gap than the addon's.
		if gear ~= nil and gear.weakest ~= nil and (gear.unit ~= row.unit or type(gear.gap) ~= "number"
			or math.floor(gear.gap + 0.5) ~= math.floor(row.gap + 0.5)) then
			row.weakestInspect = true
		end
	elseif gear ~= nil then
		row.gap, row.unit = gear.gap, gear.unit
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

--- The order for `sort` ({ key, descending }; nil: Gap high to low): the column's values, a
--- blank last both ways; Gap puts gaps in DPS before those in points, each by number. Equal
--- values go by name, A to Z.
local function comparator(sort)
	local key = sort ~= nil and sort.key or "gap"
	local descending = sort == nil or sort.descending == true
	return function(left, right)
		local a, b = valueOf(left, key), valueOf(right, key)
		if key == "gap" and a ~= nil and b ~= nil and (left.unit == "dps") ~= (right.unit == "dps") then
			return left.unit == "dps"
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
		local text = string.format(row.unit == "dps" and L.groupGapDps or L.groupGapPts, rounded)
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
	if type(row.enchants) == "number" then
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
	if type(row.bareSlots) == "table" and #row.bareSlots > 0 then
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
	return lines
end

--- The button's hover: "Forever Sixty" in gold, then what it holds and what a click does.
function GroupPanel.buttonHint()
	local lines = { Theme.inlineColor("gold", L.addonName) }
	for _, text in ipairs(L.groupPanelHint) do
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", text)
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
		slot = row.slot, gain = row.gain, unit = plan.unit == "dps" and "dps" or "score", healing = current.healing,
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
	return { gap = plan.gap, unit = plan.unit, weakest = weakestOf(plan, equipped, data) }
end

--- The player's own row's fields, read on this client (the chip's and the guild row's path:
--- Comms.ownSpec, Comms.ownPlan over what is worn, Follow.buildName): { live, inspect, gear }.
local function ownFields(data)
	local build = Follow.build
	local ranks = type(data) == "table" and type(data.classes) == "table" and Talents.readRanks(data) or {}
	local spec = Comms.ownSpec(data, build, Talents.playerClassSlug(), ranks)
	local equipped = Gear.readEquipped()
	local ok, plan = pcall(Comms.ownPlan, data, spec, equipped)
	plan = ok and plan or nil
	local gear = plan ~= nil and { gap = plan.gap, unit = plan.unit, weakest = weakestOf(plan, equipped, data) } or nil
	local live = {
		spec = spec, gearGap = plan ~= nil and plan.gap or nil, unit = plan ~= nil and plan.unit or nil,
		buildName = build ~= nil and type(data) == "table" and Follow.buildName(data, build) or nil,
	}
	return live, Inspect.readSelf(), gear
end

--- Every unit of the group as the rows take it: { unit, name, realm, short, classToken, level,
--- connected, own, inspect (Inspect.get's fresh record; the player's own read now), live
--- (Comms.seen's fresh entry; the player's own fields), record (Score.forCharacter), gear (the
--- plan measured on the inspected gear) }. Empty out of a group.
function GroupPanel.members(data)
	data = data or GroupPanel.data or ns.Data
	local members = {}
	local now = Theme.now()
	local _, ownRealm = Theme.playerNameRealm()
	local faction = Compat.playerFaction()
	for _, unit in ipairs(GroupPanel.groupUnits()) do
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
function GroupPanel.model()
	local data = GroupPanel.data or ns.Data
	return GroupPanel.rows(GroupPanel.members(data), GroupPanel.sort, {
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

--- Reads every unit's auras (the player alone out of a group) and checks them against the
--- instance's needs; the chip gets the summary and the open flyout redraws. `again`: the same
--- check read anew (its time kept). Answers the result.
local function runCheck(dry, again)
	local units = GroupPanel.groupUnits()
	local alone = #units == 0
	if alone then
		units = { "player" }
	end
	local members = {}
	for _, unit in ipairs(units) do
		local name = PlayerTooltip.nameOf(unit)
		if type(name) == "string" and name ~= "" then
			local _, classToken = try(client("UnitClass"), unit)
			local member = {
				short = name, classToken = classToken, connected = try(client("UnitIsConnected"), unit) ~= false,
				auras = GroupPanel.aurasOf(unit),
			}
			if isPlayer(unit) then
				member.imbue = GroupPanel.ownImbue()
			end
			members[#members + 1] = member
		end
	end
	local needs = GroupPanel.needsFor(GroupPanel.data or ns.Data, GroupPanel.instanceKey())
	local result = GroupPanel.check(members, needs, GroupPanel.spellName)
	-- Fix round 2: the player alone, so the summaries say what they lack (GroupPanel.summary).
	result.alone = alone or nil
	local previous = GroupPanel.result
	if again and previous ~= nil then
		result.at, result.clock, result.open, result.untilAt, result.dry = previous.at, previous.clock, previous.open,
			previous.untilAt, previous.dry
		GroupPanel.result = result
		Chip.updateRaidCheck(GroupPanel.missingLines(result), GroupPanel.summary(result))
	else
		result.at, result.clock, result.open, result.dry = Theme.now(), date("%H:%M"), not dry or nil, dry or nil
		GroupPanel.result = result
		Chip.setRaidCheck(GroupPanel.missingLines(result), GroupPanel.summary(result), not dry)
	end
	return result
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

--- Where the count line (or "Not in a group.") sits, from the flyout's top.
function GroupPanel.countsTop()
	return S.panelPadding + S.panelHeader + S.panelRowGap
end

--- Out of a group: the header and the one line.
function GroupPanel.soloHeight()
	return GroupPanel.countsTop() + GroupPanel.lineHeight() + S.panelPadding
end

--- Where the column headers sit, and the rows under them.
function GroupPanel.headerTop()
	return GroupPanel.countsTop() + GroupPanel.lineHeight() + S.panelRowGap
end

function GroupPanel.rowsTop()
	return GroupPanel.headerTop() + S.guildHeaderRow
end

--- The ready check's lines as drawn under the rows: its title, the missing lines (at most
--- `groupMissingLines`, the last saying "and N more") or "Nobody is missing anything.", the
--- coverage title and lines; each { text, color }. Empty before a check. Pure.
function GroupPanel.checkText(result)
	if type(result) ~= "table" then
		return {}
	end
	local lines = { { text = string.format(L.groupCheckTitle, result.clock or ""), color = "gold" } }
	local missing = GroupPanel.missingLines(result)
	if #missing == 0 then
		lines[#lines + 1] = { text = L.groupAllSet, color = "muted" }
	end
	local room = S.groupMissingLines
	for index, line in ipairs(missing) do
		if index < room or #missing == room then
			lines[#lines + 1] = { text = line, color = "body" }
		elseif index == room then
			lines[#lines + 1] = { text = string.format(L.groupMissingMore, #missing - room + 1), color = "muted" }
		end
	end
	local coverage = GroupPanel.coverageLines(result)
	if #coverage > 0 then
		lines[#lines + 1] = { text = L.groupCoverageTitle, color = "gold", gap = true }
		for index, line in ipairs(coverage) do
			lines[#lines + 1] = { text = line, color = #result.coverage[index].by > 0 and "body" or "muted" }
		end
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

--- How many member rows show at once under `limit` (the anchor's height: the Social window's on
--- 1.60; the flyout is never taller) with `extra` (the check region) below them:
--- `groupVisibleRows` at most, one at least; all of them where the client cannot say.
function GroupPanel.fitRows(limit, extra)
	if type(limit) ~= "number" or limit <= 0 then
		return S.groupVisibleRows
	end
	local fixed = GroupPanel.rowsTop() + (extra or 0) + S.panelPadding
	local room = math.floor((limit - fixed) / GroupPanel.rowHeight())
	return math.max(1, math.min(S.groupVisibleRows, room))
end

-- The frames -------------------------------------------------------------------------------------

--- The client's raid frame, or nil (looked up each time, guarded).
local function pane()
	return Widgets.clientFrame(GroupPanel.FRAME)
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

--- Whether the Raid tab is the one open: the anchor placed and RaidFrame shown.
function GroupPanel.tabOpen()
	return GroupPanel.anchor() ~= nil and Widgets.isShown(pane())
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

--- The flyout on its home (GroupPanel.home: the anchor, else the closed window), TOPLEFT at its
--- TOPRIGHT (Widgets.placeFlyout); with no social frame at all, on UIParent: its TOPRIGHT at the
--- chip's TOPLEFT, else at the screen's top right.
local function placeFlyout(flyout)
	local frame = GroupPanel.home()
	if frame ~= nil then
		attach(flyout, frame, "flyoutHome")
		Widgets.placeFlyout(flyout, frame, nil, S.panelOffsetX, S.panelOffsetY)
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

--- The sentinel (fix round 2, N1): a tiny addon-owned frame parented to RaidFrame. RaidFrame has
--- no rect on 1.60 but still shows and hides with the Raid tab, and so does its child, so the
--- sentinel's own show and hide scripts (GroupPanel.onTab) hear every tab switch inside the open
--- Social window, which nothing else does. Scripts on the addon's own frame only, no hook. Made
--- the first time a refresh finds RaidFrame (never in combat: it waits for its end, as the button
--- does); moved if the client's RaidFrame is ever another frame. nil without one.
function GroupPanel.ensureSentinel()
	local raid = pane()
	if raid == nil then
		return nil
	end
	local sentinel = GroupPanel.sentinel
	if sentinel == nil then
		if Theme.inCombat() then
			GroupPanel.deferred = true
			return nil
		end
		sentinel = CreateFrame("Frame", nil, raid)
		sentinel:SetSize(1, 1)
		GroupPanel.sentinelHome = raid
		sentinel:SetScript("OnShow", function()
			GroupPanel.onTab()
		end)
		sentinel:SetScript("OnHide", function()
			GroupPanel.onTab()
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
local function syncTab()
	GroupPanel.ensureSentinel()
	local button = GroupPanel.ensureButton()
	if button == nil then
		return
	end
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
			if hint ~= nil then
				local lines = {}
				for line, text in ipairs(hint) do
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
	makeHeaders(flyout)
	-- The rows: an addon-owned ScrollFrame clipping a content frame of pooled rows; the wheel
	-- moves the first row shown, three a notch.
	local holder = CreateFrame("ScrollFrame", nil, flyout)
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
	check:SetWidth(GroupPanel.innerWidth())
	local rule = Theme.texture(check, "ARTWORK", "border")
	rule:SetSize(GroupPanel.innerWidth(), S.border)
	rule:SetPoint("TOPLEFT", check, "TOPLEFT", 0, S.panelRegionGap / 2)
	check:Hide()
	GroupPanel.checkFrame = check
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
	local height = GroupPanel.rowHeight()
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

--- Draws the flyout: out of a group "Not in a group." alone; in one the count line, the headers,
--- the rows (as many as fit under the anchor's height) and the last ready check; sized to them.
function GroupPanel.draw()
	local flyout = GroupPanel.flyout
	if flyout == nil then
		return nil
	end
	GroupPanel.stale = nil
	GroupPanel.draws = GroupPanel.draws + 1
	if not GroupPanel.inGroup() then
		GroupPanel.last = { rows = {}, counts = { members = 0, inspected = 0, live = 0 }, visible = 0, solo = true }
		GroupPanel.countsLine:SetText(fitted(L.groupNotInGroup, GroupPanel.innerWidth()))
		GroupPanel.countsLine:SetTextColor(Theme.rgb(Theme.HEX.muted))
		drawHeaders(false)
		GroupPanel.holder:Hide()
		for _, row in ipairs(GroupPanel.rowFrames) do
			row.entry = nil
			row.frame:Hide()
		end
		drawCheck(flyout, {}, 0)
		flyout:SetHeight(GroupPanel.soloHeight())
		return GroupPanel.last
	end
	local rows, counts = GroupPanel.model()
	local lines = GroupPanel.checkText(GroupPanel.result)
	local visible = GroupPanel.fitRows(Widgets.frameHeight(GroupPanel.anchor()), GroupPanel.checkHeight(lines))
	GroupPanel.last = { rows = rows, counts = counts, visible = visible }
	GroupPanel.countsLine:SetText(fitted(string.format(L.groupCounts, counts.members, counts.inspected, counts.live),
		GroupPanel.innerWidth()))
	GroupPanel.countsLine:SetTextColor(Theme.rgb(Theme.HEX.muted))
	drawHeaders(true)
	local shownRows = math.max(1, math.min(visible, #rows))
	local holder = GroupPanel.holder
	holder:SetHeight(shownRows * GroupPanel.rowHeight())
	GroupPanel.content:SetHeight(shownRows * GroupPanel.rowHeight())
	holder:Show()
	GroupPanel.renderRows()
	local y = drawCheck(flyout, lines, GroupPanel.rowsTop() + shownRows * GroupPanel.rowHeight())
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
	placeFlyout(flyout)
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
--- social frame (the flyout is then at the chip's side). The lines to print.
function GroupPanel.command()
	local lines = guarded(function()
		local flyout = GroupPanel.ensureFlyout()
		if flyout:IsShown() then
			closeFlyout()
			return {}
		end
		openFlyout(flyout)
		if GroupPanel.home() ~= nil and not GroupPanel.tabOpen() then
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
			if GroupPanel.tabOpen() ~= GroupPanel.tabWasOpen then
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
--- claims no clean check (fix round 1, D3). The lines to print.
function GroupPanel.handleReady()
	local result = guarded(function()
		local result = runCheck(true)
		redraw()
		return result
	end)
	if result == nil then
		return { L.groupPanelOff }
	end
	local lines = { L.groupReadyDry }
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
--- the window was closed.
local function reread()
	syncTab()
	local flyout = GroupPanel.flyout
	if flyout == nil or not flyout:IsShown() then
		return
	end
	placeFlyout(flyout)
	if GroupPanel.stale and Widgets.isVisible(flyout) then
		GroupPanel.draw()
	end
end

--- The button's show script (it shows with the window): the Raid tab and the window read again.
function GroupPanel.onButtonShow()
	guarded(reread)
end

--- The sentinel's show and hide scripts (fix round 2, N1: RaidFrame shown or hidden, the Raid tab
--- opened or left, the window opened or closed): the Raid tab and the window read again.
function GroupPanel.onTab()
	guarded(reread)
end

--- Listens for the members' hellos and snapshot! answers, once (Options.register).
function GroupPanel.register()
	if GroupPanel.registered then
		return
	end
	GroupPanel.registered = true
	Comms.on("hello", function()
		GroupPanel.onLive()
	end)
	Comms.on("snapshot!", function()
		GroupPanel.onLive()
	end)
end

--- /fs diag's line (ruling 5): "Group panel: open, 5 members (3 inspected, 2 live), ready check
--- 21:04: 2 missing" (open: the flyout shows and is visible), counted now as the flyout counts.
--- Fix round 2: the count always includes the player, so out of a group it reads "1 member
--- (you)", and a check that read the player alone says what they lack: "ready check 19:43: you
--- are missing flask, food" (or "you are missing nothing").
function GroupPanel.diagLine()
	local open = Widgets.isShown(GroupPanel.flyout) and Widgets.isVisible(GroupPanel.flyout)
	local state = open and L.diagPanelOpen or L.diagPanelClosed
	local result = GroupPanel.result
	local check = L.diagGroupNone
	if result ~= nil and result.alone then
		local lack = GroupPanel.ownLack(result)
		check = lack ~= nil and string.format(L.diagGroupCheckYou, result.clock or "", lack)
			or string.format(L.diagGroupCheckYouNone, result.clock or "")
	elseif result ~= nil then
		check = string.format(L.diagGroupCheck, result.clock or "", #result.missing)
	end
	if not GroupPanel.inGroup() then
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
