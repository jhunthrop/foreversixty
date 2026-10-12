-- addon/ForeverSixty/surfaces/ItemTooltip.lua
-- The item tooltip (spec section 4.3): after the client's own lines, a
-- "Forever Sixty" section. A header naming where the weights came from
-- ("Forever Sixty · Your sim"), the verdict against what is worn ("Upgrade
-- +4.2 DPS over Chausses of Westfall", "Sidegrade", "Not for you until 40:
-- plate armor"), any capped stat the item carries, the plan ("Planned for
-- Arms at 30"), one source line with its drop chance and two alternatives.
-- The compare tooltip beside it shows the band's pick for the slot.
--
-- Pure model, thin hook: ItemTooltip.section builds the lines from plain
-- tables (every number from Progress.slotVerdict, the weights from
-- Gear.currentWeights, the wear rules from Rules.canEquip); the guarded
-- hooks after it read the client and draw. A section with nothing to say is
-- not drawn at all (spec section 2.2).
--
-- The empty character-slot hover (ItemTooltip.bisLines and its hooks) and
-- the player tooltip's rating line (onUnitTooltip) are kept as they were.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Talents = ns.Talents or require("Talents")
local Gear = ns.Gear or require("Gear")
local Export = ns.Export or require("Export")
local Compat = ns.Compat or require("Compat")
local Theme = ns.Theme or require("Theme")
local Prefs = ns.Prefs or require("Prefs")
local Rules = ns.Rules or require("Rules")
local Follow = ns.Follow or require("Follow")
local Progress = ns.Progress or require("Progress")

local ItemTooltip = {}

--- The site slot name -> the client's equipped-item slot id, built once
--- from Export's own table.
local SLOT_IDS = {}
for _, entry in ipairs(Export.INVENTORY_SLOTS) do
	SLOT_IDS[entry.slot] = entry.id
end

--- An item link's id.
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

local function equippedLinkFor(slot)
	local slotId = SLOT_IDS[slot]
	if slotId == nil or type(GetInventoryItemLink) ~= "function" then
		return nil
	end
	return GetInventoryItemLink("player", slotId)
end

--- The site slot currently holding `itemLink`, or nil (a bag item, another
--- unit's item, or nothing equipped there), by exact link.
function ItemTooltip.equippedSlotFor(itemLink)
	if itemLink == nil then
		return nil
	end
	for slot in pairs(SLOT_IDS) do
		if equippedLinkFor(slot) == itemLink then
			return slot
		end
	end
	return nil
end

-- ---------------------------------------------------------------------
-- The section's lines, pure.
-- ---------------------------------------------------------------------

--- Below this share of the worn item's own value a gain or a loss reads as
--- a Sidegrade, both ways: two items a stray point of a minor stat apart are
--- not an upgrade either way.
ItemTooltip.SIDEGRADE_THRESHOLD_PERCENT = 0.5

--- The smallest gain each unit prints (one decimal for DPS and HPS, none
--- for a score): anything smaller would print "+0.0", so it is a Sidegrade.
local SMALLEST = { dps = 0.05, hps = 0.05, score = 0.5 }

local function round(value)
	return math.floor(value + 0.5)
end

local function unitWord(unit)
	return unit == "hps" and L.tooltipUnitHps or L.tooltipUnitDps
end

--- The verdict line for one Progress.slotVerdict answer, as { text, color },
--- or nil with nothing to say. The verdict was asked with the HOVERED item
--- as `plan` and the worn one as `worn`, so its delta is worn minus hovered
--- and the hovered item's gain is minus the delta: `better` is the worn item
--- ahead (a downgrade), `behind` the hovered one ahead or level (an
--- upgrade), `empty` nothing worn (it fills the slot), `match` the item worn.
--- `unit` is "dps", "hps" or "score"; `wornName` the worn item's name (nil
--- while the client has not sent it); `wornScore` the worn item's value in
--- the same unit, the base of the sidegrade threshold.
function ItemTooltip.verdictLine(verdict, unit, wornName, wornScore)
	if verdict == nil then
		return nil
	end
	if verdict.state == "match" then
		return { text = L.tooltipEquipped, color = "muted" }
	end
	if verdict.delta == nil or verdict.state == "noplan" then
		return nil
	end
	unit = unit or "score"
	local gain = -verdict.delta
	local size = math.abs(gain)
	if verdict.state == "empty" then
		if gain < SMALLEST[unit] then
			return nil
		end
		if unit == "score" then
			return { text = string.format(L.tooltipFillsScore, round(gain)), color = "success" }
		end
		return { text = string.format(L.tooltipFills, gain, unitWord(unit)), color = "success" }
	end
	local threshold = math.abs(wornScore or 0) * ItemTooltip.SIDEGRADE_THRESHOLD_PERCENT / 100
	if size < SMALLEST[unit] or size < threshold then
		return { text = L.tooltipSide, color = "muted" }
	end
	local name = wornName or L.tooltipNamePending
	if gain > 0 then
		if unit == "score" then
			return { text = string.format(L.tooltipVerdictScore, round(gain), name), color = "success" }
		end
		local format = unit == "hps" and L.tooltipVerdictHps or L.tooltipVerdictDps
		return { text = string.format(format, gain, name), color = "success" }
	end
	if unit == "score" then
		return { text = string.format(L.tooltipDownScore, round(size)), color = "warning" }
	end
	return { text = string.format(L.tooltipDown, size, unitWord(unit)), color = "warning" }
end

--- The line Rules.canEquip's refusal reads as, in place of the verdict:
--- "Not for you until 40: plate armor" or "Not for you: rogues cannot use
--- two-handed swords". `subclass` is canEquip's own plain-words reason. nil
--- for "ok".
function ItemTooltip.untilLine(classToken, answer, level, subclass)
	if answer == "until" then
		return string.format(L.tooltipUntil, level, subclass)
	end
	if answer == "never" then
		local class = type(classToken) == "string" and L.tooltipClassPlural[classToken:upper()] or nil
		return string.format(L.tooltipNever, class or tostring(classToken), subclass)
	end
	return nil
end

--- data/pipeline/addonlua.py's SOURCE_KIND_CODES, decoded back for the
--- source line. tests/tooltip_bis_spec.lua pins this table's key set
--- against the pipeline's set of source kinds.
ItemTooltip.SOURCE_KIND_NAMES = {
	Q = "Quest",
	V = "Vendor",
	D = "Dungeon",
	C = "Crafted",
	R = "Reputation",
	P = "PvP",
	W = "World",
	A = "Raid",
	B = "World Drop",
}

--- The kinds whose source is "<instance>: <boss>".
local INSTANCE_KINDS = { D = true, A = true }

--- A pick's drop chance as a whole percent, never 0 for a real drop.
local function percentOf(chance)
	if type(chance) ~= "number" or chance <= 0 then
		return nil
	end
	return math.max(1, round(chance * 100))
end

--- One line saying where a pick comes from: "The Deadmines · Defias
--- Overseer · 14%" for an instance drop (the source split on its first
--- ": "), "<name> · Vendor" or "<source> · Quest" for the rest (the kind
--- left out when the source already says it, "World drop"), the chance last
--- when the data has one. nil with no source.
function ItemTooltip.sourceLine(pick)
	if type(pick) ~= "table" or type(pick.source) ~= "string" or pick.source == "" then
		return nil
	end
	local first, second
	local place, boss = pick.source:match("^(.-): (.+)$")
	if INSTANCE_KINDS[pick.kind] and place ~= nil then
		first, second = place, boss
	else
		first = pick.source
		local kindName = ItemTooltip.SOURCE_KIND_NAMES[pick.kind]
		if kindName ~= nil and kindName:lower() ~= first:lower() then
			second = kindName
		end
	end
	local percent = percentOf(pick.chance)
	if second ~= nil and percent ~= nil then
		return string.format(L.tooltipSource, first, second, percent)
	end
	if second ~= nil then
		return string.format(L.tooltipSourcePair, first, second)
	end
	if percent ~= nil then
		return string.format(L.tooltipSourceChance, first, percent)
	end
	return first
end

--- "Also: <name>, <name>" from up to two names; `false` for a name the
--- client has not sent yet, which prints the pending mark. nil with none.
function ItemTooltip.alternativesLine(names)
	local shown = {}
	for index = 1, 2 do
		local name
		if names ~= nil then
			name = names[index]
		end
		if name == false then
			shown[#shown + 1] = L.tooltipNamePending
		elseif type(name) == "string" then
			shown[#shown + 1] = name
		end
	end
	if #shown == 2 then
		return string.format(L.tooltipAlternatives, shown[1], shown[2])
	end
	if #shown == 1 then
		return string.format(L.tooltipAlternativesOne, shown[1])
	end
	return nil
end

local TAG_TEXT = { sim = "tooltipTagSim", build = "tooltipTagBuild", default = "tooltipTagDefault" }

--- The section's first line: "Forever Sixty" in gold, then where the
--- weights came from, muted ("· Your sim"); the name alone with none.
function ItemTooltip.headerText(tag)
	local text = Theme.inlineColor("gold", L.addonName)
	local key = TAG_TEXT[tag]
	if key ~= nil then
		text = text .. Theme.inlineColor("muted", string.format(L.tooltipHeaderTag, L[key]))
	end
	return text
end

--- "dps", "hps" or "score": the unit Gear.currentWeights' answer counts in.
function ItemTooltip.unitOf(current)
	if current == nil or current.dpsPerPoint == nil then
		return "score"
	end
	return current.healing and "hps" or "dps"
end

--- "<stat> capped: more is wasted" for each stat the item carries that the
--- sim's weights message names as capped, sorted. Empty without caps.
local function capTexts(stats, caps)
	local lines = {}
	if type(stats) ~= "table" or type(caps) ~= "table" then
		return lines
	end
	for _, stat in ipairs(caps) do
		if type(stats[stat]) == "number" and stats[stat] > 0 then
			lines[#lines + 1] = string.format(L.tooltipCapped, stat)
		end
	end
	table.sort(lines)
	return lines
end

--- The capped-stat lines for an item link, against `current`
--- (ItemTooltip.current's shape: its `caps` list).
function ItemTooltip.capLines(itemLink, current)
	if itemLink == nil or type(current) ~= "table" then
		return {}
	end
	return capTexts(Gear.statsOf(itemLink), current.caps)
end

--- The worn item's name, from its own link ("[Name]"), else the client.
local function wornNameOf(worn, nameOf)
	if worn == nil then
		return nil
	end
	local name = type(worn.link) == "string" and worn.link:match("%[(.-)%]") or nil
	return name or nameOf(worn.itemId)
end

--- The slot the hovered item is judged in, among `slots` (each one it may
--- be worn in): the one already holding it, else an empty one, else the one
--- whose worn item scores lowest (of two rings, the weaker is the one to
--- replace).
local function chooseSlot(slots, itemId, equipped, weights)
	local best, bestScore
	for _, slot in ipairs(slots) do
		local worn = equipped[slot]
		if worn ~= nil and worn.itemId == itemId then
			return slot
		end
	end
	for _, slot in ipairs(slots) do
		local worn = equipped[slot]
		if worn == nil then
			return slot
		end
		local score = Gear.score(worn.stats, weights)
		if bestScore == nil or score < bestScore then
			best, bestScore = slot, score
		end
	end
	return best
end

--- The spec's own tree name ("Arms" for "warrior-arms").
local function specName(data, spec)
	local classSlug, rest = spec:match("^([^-]+)%-(.+)$")
	local class = data.classes ~= nil and data.classes[classSlug] or nil
	for _, tab in ipairs(class ~= nil and class.tabs or {}) do
		if Export.slugify(tab.name) == rest then
			return tab.name
		end
	end
	return Talents.titleCase(rest or spec)
end

--- The lowest band at or above the character's own (any band below 20)
--- whose pick for one of `slots` is `itemId`, or nil.
function ItemTooltip.plannedBand(data, spec, faction, slots, itemId, level)
	local specBis = spec ~= nil and type(data.bis) == "table" and data.bis[spec] or nil
	if type(specBis) ~= "table" or not faction then
		return nil
	end
	local floor = Gear.bandFor(specBis, level) or 0
	local bands = {}
	for band in pairs(specBis) do
		if type(band) == "number" and band >= floor then
			bands[#bands + 1] = band
		end
	end
	table.sort(bands)
	for _, band in ipairs(bands) do
		for _, slot in ipairs(slots) do
			local pick = Gear.pick(spec, band, faction, slot, data)
			if pick ~= nil and pick.item_id == itemId then
				return band
			end
		end
	end
	return nil
end

local function buildPlans(build, slots, itemId)
	if type(build) ~= "table" then
		return false
	end
	local wanted = {}
	for _, slot in ipairs(slots) do
		wanted[slot] = true
	end
	for _, entry in ipairs(build.gear or {}) do
		if wanted[entry.slot] and entry.itemId == itemId then
			return true
		end
	end
	return false
end

--- The section for one hovered item, as a list of { text, color } lines in
--- draw order, then the compare tooltip's item ({ itemId, link }, the
--- band's pick for the slot, or nil) and the ids whose names the client has
--- not sent yet (a set). `ctx` holds everything the hook read:
---
--- - `data`, `spec`, `classToken`, `level`, `faction`;
--- - `item`: { itemId, link, equipLocation, itemClass, itemSubClass, stats };
--- - `current`: Gear.currentWeights' answer (nil: no weights for the spec);
--- - `equipped`: Gear.readEquipped()'s table; `build`: the loaded build;
--- - `nameOf(itemId)`: the client's name for an item, nil until cached.
---
--- Rules first: an item the character cannot wear in any slot it fits gets
--- the refusal in place of the verdict, and is never called an upgrade; the
--- plan and source lines still print. One source block per tooltip: the
--- hovered item's own when it is the band's pick, else the pick's ("Next:").
function ItemTooltip.section(ctx)
	local lines, pending = {}, {}
	local item = ctx.item or {}
	local slots = Gear.slotsFor(item.equipLocation)
	if #slots == 0 then
		return lines, nil, pending
	end
	local data = ctx.data or {}
	local equipped = ctx.equipped or {}
	local nameOf = ctx.nameOf or Compat.itemName
	local function named(itemId)
		local name = nameOf(itemId)
		if name == nil then
			pending[itemId] = true
			return false
		end
		return name
	end

	local wearable, refusal = {}, nil
	for _, slot in ipairs(slots) do
		local answer, unlock, reason = Rules.canEquip(ctx.classToken, ctx.level, item.itemClass,
			item.itemSubClass, item.equipLocation, slot)
		if answer == "ok" then
			wearable[#wearable + 1] = slot
		elseif refusal == nil then
			refusal = ItemTooltip.untilLine(ctx.classToken, answer, unlock, reason)
		end
	end

	local current = ctx.current
	local weights = current and current.weights or nil
	local slot = chooseSlot(#wearable > 0 and wearable or slots, item.itemId, equipped, weights)
	if #wearable == 0 then
		if refusal ~= nil then
			lines[#lines + 1] = { text = refusal, color = "warning" }
		end
	elseif weights ~= nil then
		local worn = equipped[slot]
		-- Progress.slotVerdict(slot, plan, worn, ...) with the HOVERED item as
		-- `plan` and the worn one as `worn`: its delta is worn minus hovered,
		-- which verdictLine reads back as the hovered item's gain.
		local verdict = Progress.slotVerdict(slot, { itemId = item.itemId, stats = item.stats }, worn,
			weights, current.dpsPerPoint)
		local wornScore = worn ~= nil and Gear.score(worn.stats, weights) * (current.dpsPerPoint or 1) or nil
		local unit = ItemTooltip.unitOf(current)
		local wornName = worn ~= nil and wornNameOf(worn, nameOf) or nil
		if worn ~= nil and wornName == nil then
			pending[worn.itemId] = true
		end
		local line = ItemTooltip.verdictLine(verdict, unit, wornName, wornScore)
		if line ~= nil then
			lines[#lines + 1] = line
		end
		for _, text in ipairs(capTexts(item.stats, current.caps)) do
			lines[#lines + 1] = { text = text, color = "body" }
		end
	end

	local band = ItemTooltip.plannedBand(data, ctx.spec, ctx.faction, slots, item.itemId, ctx.level)
	if band ~= nil then
		lines[#lines + 1] = { text = string.format(L.tooltipPlanned, specName(data, ctx.spec), band), color = "gold" }
	elseif buildPlans(ctx.build, slots, item.itemId) then
		lines[#lines + 1] = { text = string.format(L.tooltipPlannedBuild, Follow.buildName(data, ctx.build)),
			color = "gold" }
	end

	local specBis = ctx.spec ~= nil and type(data.bis) == "table" and data.bis[ctx.spec] or nil
	local currentBand = ctx.faction and Gear.bandFor(specBis, ctx.level) or nil
	local pick, isPick
	if currentBand ~= nil then
		for _, candidate in ipairs(slots) do
			local found = Gear.pick(ctx.spec, currentBand, ctx.faction, candidate, data)
			if found ~= nil and found.item_id == item.itemId then
				pick, isPick = found, true
				break
			end
		end
		if pick == nil then
			pick = Gear.pick(ctx.spec, currentBand, ctx.faction, slot, data)
		end
	end
	local target
	if pick ~= nil then
		local source = ItemTooltip.sourceLine(pick)
		if isPick then
			if source ~= nil then
				lines[#lines + 1] = { text = source, color = "muted" }
			end
		else
			local name = named(pick.item_id) or L.tooltipNamePending
			lines[#lines + 1] = {
				text = source ~= nil and string.format(L.tooltipNext, name, source)
					or string.format(L.tooltipNextOnly, name),
				color = "muted",
			}
		end
		local names = {}
		for _, alternative in ipairs(pick.alternatives or {}) do
			if alternative.item_id ~= nil and alternative.item_id ~= item.itemId and #names < 2 then
				names[#names + 1] = named(alternative.item_id)
			end
		end
		local also = ItemTooltip.alternativesLine(names)
		if also ~= nil then
			lines[#lines + 1] = { text = also, color = "muted" }
		end
		target = { itemId = pick.item_id, link = Compat.displayLink(pick.item_id) }
	end
	return lines, target, pending
end

-- ---------------------------------------------------------------------
-- The empty character-slot hover (kept as it was; L4 replaces it).
-- ---------------------------------------------------------------------

--- Site slot name -> the Blizzard PaperDoll button's global frame name.
ItemTooltip.BIS_SLOT_BUTTONS = {
	head = "CharacterHeadSlot",
	neck = "CharacterNeckSlot",
	shoulder = "CharacterShoulderSlot",
	back = "CharacterBackSlot",
	chest = "CharacterChestSlot",
	wrist = "CharacterWristSlot",
	hands = "CharacterHandsSlot",
	waist = "CharacterWaistSlot",
	legs = "CharacterLegsSlot",
	feet = "CharacterFeetSlot",
	finger1 = "CharacterFinger0Slot",
	finger2 = "CharacterFinger1Slot",
	trinket1 = "CharacterTrinket0Slot",
	trinket2 = "CharacterTrinket1Slot",
	main_hand = "CharacterMainHandSlot",
	off_hand = "CharacterSecondaryHandSlot",
	ranged = "CharacterRangedSlot",
}

--- The item ids newly best at `band` for `faction`, or {}.
local function newAtBand(data, spec, band, faction)
	local bySpec = type(data.bis_new) == "table" and data.bis_new[spec] or nil
	local byBand = bySpec ~= nil and bySpec[band] or nil
	return (byBand ~= nil and byBand[faction]) or {}
end

local function isAmong(ids, itemId)
	for _, id in ipairs(ids) do
		if id == itemId then
			return true
		end
	end
	return false
end

--- The empty-slot row's right-aligned tag, as { text, color }: "equipped"
--- beats "new at <band>" beats "BiS · lvl <band>".
local function bisTag(itemId, band, equippedItemId, newIds)
	if equippedItemId == itemId then
		return { text = L.tooltipBisEquipped, color = "success" }
	end
	if isAmong(newIds, itemId) then
		return { text = string.format(L.tooltipBisNew, band), color = "gold" }
	end
	return { text = string.format(L.tooltipBisDefaultTag, band), color = "muted" }
end

--- The item's link once GetItemInfo knows it, else "item:<id>".
function ItemTooltip.bisItemLink(itemId)
	return Compat.displayLink(itemId)
end

--- The empty-slot section's muted source line, "<kind> · <place>".
local function bisSourceLine(sourceCode, sourceLabel)
	if type(sourceLabel) ~= "string" or sourceLabel == "" then
		return nil
	end
	local kindName = ItemTooltip.SOURCE_KIND_NAMES[sourceCode] or sourceCode
	return string.format(L.tooltipBisSourceLabel, kindName, (sourceLabel:gsub(": ", " · ")))
end

ItemTooltip.BIS_ICON_SIZE = 18

--- The empty-slot section for one site slot, as draw ops: a "doubleline"
--- op for the pick (icon + link left, the tag right), then the item level
--- (once cached) and the muted source line. Empty in combat, with no class,
--- spec, band or pick. The second return is the pick for the compare tooltip.
function ItemTooltip.bisLines(data, slot)
	if data == nil or slot == nil then
		return {}
	end
	if type(InCombatLockdown) == "function" and InCombatLockdown() then
		return {}
	end
	local classSlug = Talents.playerClassSlug()
	if classSlug == nil then
		return {}
	end
	local ranks = Talents.readRanks(data)
	local spec = Gear.specOf(data, classSlug, ranks)
	local specBis = spec ~= nil and type(data.bis) == "table" and data.bis[spec] or nil
	local band = Gear.bandFor(specBis, Compat.playerLevel())
	if band == nil then
		return {}
	end
	local faction = Compat.playerFaction()
	local pick = faction ~= nil and Gear.pick(spec, band, faction, slot, data) or nil
	if pick == nil then
		return {}
	end
	local itemId = pick.item_id
	local tag = bisTag(itemId, band, itemIdOf(equippedLinkFor(slot)), newAtBand(data, spec, band, faction))
	local link = ItemTooltip.bisItemLink(itemId)
	local icon = Theme.inlineIcon(Compat.itemIcon(itemId), ItemTooltip.BIS_ICON_SIZE)
	local ops = {
		{ kind = "doubleline", left = icon .. " " .. link, right = tag.text, rightColor = tag.color },
	}
	local level = Compat.itemLevel(itemId)
	if level ~= nil then
		ops[#ops + 1] = { kind = "line", text = string.format(L.tooltipBisItemLevel, level), color = "muted" }
	end
	local source = bisSourceLine(pick.kind, pick.source)
	if source ~= nil then
		ops[#ops + 1] = { kind = "line", text = source, color = "muted" }
	end
	return ops, { itemId = itemId, link = link }
end

-- ---------------------------------------------------------------------
-- The hooks.
-- ---------------------------------------------------------------------

--- The generated data table, set by Options.register().
ItemTooltip.data = nil

--- Gear.currentWeights' answer for the character and spec logged in, set by
--- Options.refreshWeightsMessage (at login, on a respec, when the inbox is
--- read). nil before that, when the hook asks Gear.currentWeights itself.
ItemTooltip.current = nil

--- The section, computed once per item link until something it reads
--- changes (Options resets it with every surface refresh).
ItemTooltip.cache = {}

function ItemTooltip.resetCache()
	ItemTooltip.cache = {}
	return ItemTooltip.cache
end

--- Everything ItemTooltip.section reads, from the client.
local function contextFor(data, itemLink)
	local id, itemType, itemSubType, equipLocation, _, classId, subClassId = Compat.itemInfoInstant(itemLink)
	local classSlug = Talents.playerClassSlug()
	local spec = classSlug ~= nil and data.classes ~= nil and Gear.specOf(data, classSlug, Talents.readRanks(data))
		or nil
	local current = ItemTooltip.current
	if current == nil and spec ~= nil then
		current = Gear.currentWeights(data, spec, ForeverSixtyInbox, Export.characterKey())
	end
	return {
		data = data, spec = spec, current = current,
		classToken = classSlug ~= nil and classSlug:upper() or nil,
		level = Compat.playerLevel(), faction = Compat.playerFaction(),
		equipped = Gear.readEquipped(), build = Follow.build,
		item = {
			itemId = tonumber(id) or itemIdOf(itemLink), link = itemLink, equipLocation = equipLocation,
			itemClass = classId or itemType, itemSubClass = subClassId or itemSubType,
			stats = Gear.statsOf(itemLink),
		},
	}
end

local function cachedSection(itemLink)
	local cached = ItemTooltip.cache[itemLink]
	if cached ~= nil then
		return cached
	end
	local data = ItemTooltip.data
	local entry = { lines = {}, pending = {} }
	if data ~= nil then
		local ctx = contextFor(data, itemLink)
		entry.lines, entry.target, entry.pending = ItemTooltip.section(ctx)
		entry.tag = ctx.current ~= nil and ctx.current.tag or nil
	end
	ItemTooltip.cache[itemLink] = entry
	return entry
end

--- The hover currently on screen, while it waits for an item the client
--- has to send (the compare tooltip's pick, a name shown as "…"):
--- GET_ITEM_INFO_RECEIVED (ItemTooltip.onItemInfoReceived) redraws it.
ItemTooltip.activeHover = nil

local function opColor(key)
	return Theme.rgb(Theme.HEX[key or "body"])
end

--- The empty-slot section's header strip icon.
ItemTooltip.HEADER_ICON_SIZE = Theme.SIZES.rotationHeaderIcon

--- Draws the empty-slot section's ops, its brand header first.
local function drawOps(tooltip, ops)
	if #ops == 0 then
		return
	end
	local hr, hg, hb = opColor("gold")
	tooltip:AddLine(Theme.inlineIcon(Theme.MEDIA.minimapIcon, ItemTooltip.HEADER_ICON_SIZE) .. " " .. L.addonName,
		hr, hg, hb)
	for _, op in ipairs(ops) do
		if op.kind == "doubleline" then
			local lr, lg, lb = opColor(op.leftColor)
			local rr, rg, rb = opColor(op.rightColor)
			tooltip:AddDoubleLine(op.left, op.right, lr, lg, lb, rr, rg, rb)
		else
			local r, g, b = opColor(op.color)
			tooltip:AddLine(op.text, r, g, b)
		end
	end
end

--- Draws the item section after the client's own lines: a blank line, the
--- header ("Forever Sixty" gold, the tag muted, the talent marks' own
--- look), then each line in its colour. Nothing at all with no lines.
local function addLines(tooltip, itemLink)
	local entry = cachedSection(itemLink)
	if #entry.lines == 0 then
		return
	end
	tooltip:AddLine(" ")
	tooltip:AddLine(ItemTooltip.headerText(entry.tag))
	for _, line in ipairs(entry.lines) do
		local r, g, b = opColor(line.color)
		tooltip:AddLine(line.text, r, g, b)
	end
	tooltip:Show()
	if entry.target ~= nil then
		Theme.showCompareTooltip(entry.target.itemId, entry.target.link)
	end
	if entry.target == nil and next(entry.pending) == nil then
		return
	end
	ItemTooltip.activeHover = {
		itemId = entry.target ~= nil and entry.target.itemId or nil,
		pending = entry.pending,
		redraw = function()
			-- Re-running SetHyperlink re-fires the hook that called
			-- onTooltip, so the tooltip is rebuilt with the data that
			-- just arrived rather than gaining a second copy of the section.
			if type(tooltip.SetHyperlink) == "function"
				and type(tooltip.IsShown) == "function" and tooltip:IsShown() then
				tooltip:SetHyperlink(itemLink)
			end
		end,
	}
end

--- One guarded body for both hook shapes. A Lua error inside a tooltip hook
--- breaks every tooltip in the game, so a failure turns the hook off.
function ItemTooltip.onTooltip(tooltip, itemLink)
	if ItemTooltip.disabled or itemLink == nil or not Prefs.flag("tooltip") then
		return
	end
	-- The compare tooltip is a GameTooltip too, so the client's processor
	-- calls this for it as well; adding lines there would open another
	-- compare tooltip for the same item, without end (seen in game as a
	-- C stack overflow).
	if tooltip == Theme.compareTooltipFrame then
		return
	end
	local ok, err = pcall(addLines, tooltip, itemLink)
	if not ok then
		ItemTooltip.disabled = true
		Theme.note(string.format(L.diagTooltipHookFailed, tostring(err)))
	end
end

function ItemTooltip.itemLinkFrom(tooltip)
	if type(tooltip) ~= "table" or type(tooltip.GetItem) ~= "function" then
		return nil
	end
	local ok, _, link = pcall(tooltip.GetItem, tooltip)
	if not ok then
		return nil
	end
	return link
end

--- The Ratings line on a player's tooltip. Same guard as the item hook.
local Ratings = ns.Ratings or require("Ratings")

local function addUnitLines(tooltip, unit)
	if type(UnitIsPlayer) == "function" and not UnitIsPlayer(unit) then
		return
	end
	local name, realm = Compat.unitName(unit)
	local card = Ratings.forCharacter(name, realm)
	if card == nil then
		return
	end
	tooltip:AddLine(Ratings.tooltipLine(card), Theme.rgb(Theme.HEX.gold))
	tooltip:Show()
end

function ItemTooltip.onUnitTooltip(tooltip, unit)
	if ItemTooltip.unitDisabled or unit == nil or not Prefs.flag("tooltip") then
		return
	end
	local ok, err = pcall(addUnitLines, tooltip, unit)
	if not ok then
		ItemTooltip.unitDisabled = true
		Theme.note(string.format(L.diagTooltipHookFailed, tostring(err)))
	end
end

function ItemTooltip.unitFrom(tooltip)
	if type(tooltip) ~= "table" or type(tooltip.GetUnit) ~= "function" then
		return nil
	end
	local ok, _, unit = pcall(tooltip.GetUnit, tooltip)
	return ok and unit or nil
end

function ItemTooltip.hasUnitProcessor()
	return ItemTooltip.hasProcessor() and Enum.TooltipDataType.Unit ~= nil
end

--- The unit hook, registered beside the item hook by the same rule.
function ItemTooltip.registerUnit()
	if ItemTooltip.unitRegistered then
		return ItemTooltip.unitHow
	end
	ItemTooltip.unitRegistered = true
	if ItemTooltip.hasUnitProcessor() then
		TooltipDataProcessor.AddTooltipPostCall(Enum.TooltipDataType.Unit, function(tooltip)
			ItemTooltip.onUnitTooltip(tooltip, ItemTooltip.unitFrom(tooltip))
		end)
		ItemTooltip.unitHow = "processor"
		return ItemTooltip.unitHow
	end
	if type(GameTooltip) == "table" and type(GameTooltip.HookScript) == "function" then
		GameTooltip:HookScript("OnTooltipSetUnit", function(tooltip)
			ItemTooltip.onUnitTooltip(tooltip, ItemTooltip.unitFrom(tooltip))
		end)
		ItemTooltip.unitHow = "legacy"
		return ItemTooltip.unitHow
	end
	ItemTooltip.unitHow = nil
	return nil
end

local function slotButton(slot)
	local buttonName = ItemTooltip.BIS_SLOT_BUTTONS[slot]
	return buttonName ~= nil and _G[buttonName] or nil
end

local function slotForOwner(owner)
	if owner == nil then
		return nil
	end
	for slot in pairs(ItemTooltip.BIS_SLOT_BUTTONS) do
		if slotButton(slot) == owner then
			return slot
		end
	end
	return nil
end

--- Whether the empty-slot section is already on a tooltip since its last
--- OnTooltipCleared (weak keys: a tooltip nothing else holds is collected).
local appendedSection = setmetatable({}, { __mode = "k" })

--- An EMPTY slot's section, appended onto `tooltip` unless already there.
function ItemTooltip.appendEmptySlotSection(tooltip, slot, button)
	if appendedSection[tooltip] then
		return
	end
	local ops, target = ItemTooltip.bisLines(ItemTooltip.data, slot)
	if #ops == 0 then
		return
	end
	drawOps(tooltip, ops)
	appendedSection[tooltip] = true
	tooltip:Show()
	if target == nil then
		return
	end
	Theme.showCompareTooltip(target.itemId, target.link)
	ItemTooltip.activeHover = {
		itemId = target.itemId,
		redraw = function()
			appendedSection[tooltip] = nil
			ItemTooltip.onSlotEnter(slot, button)
		end,
	}
end

--- The character-frame slot buttons' own OnEnter: an EMPTY slot's section.
--- A slot holding an item is the item hook's.
function ItemTooltip.onSlotEnter(slot, button)
	if ItemTooltip.slotDisabled or not Prefs.flag("tooltip") then
		return
	end
	if equippedLinkFor(slot) ~= nil then
		return
	end
	local ok, err = pcall(function()
		GameTooltip:SetOwner(button, "ANCHOR_RIGHT")
		GameTooltip:ClearLines()
		appendedSection[GameTooltip] = nil
		ItemTooltip.appendEmptySlotSection(GameTooltip, slot, button)
	end)
	if not ok then
		ItemTooltip.slotDisabled = true
		Theme.note(string.format(L.diagTooltipHookFailed, tostring(err)))
	end
end

--- Hooks every BIS_SLOT_BUTTONS global this client has. Idempotent.
function ItemTooltip.registerBisSlotButtons()
	if ItemTooltip.bisButtonsRegistered then
		return
	end
	ItemTooltip.bisButtonsRegistered = true
	for slot, buttonName in pairs(ItemTooltip.BIS_SLOT_BUTTONS) do
		local button = _G[buttonName]
		if type(button) == "table" and type(button.HookScript) == "function" then
			button:HookScript("OnEnter", function(self)
				ItemTooltip.onSlotEnter(slot, self)
			end)
		end
	end
end

function ItemTooltip.onEmptySlotTooltipCleared(tooltip)
	appendedSection[tooltip] = nil
end

--- GameTooltip's own OnShow: Blizzard's PaperDollItemSlotButton_OnUpdate
--- rebuilds the hovered slot's tooltip without firing OnEnter, so an empty
--- slot's section is appended again here (a no-op when already there).
function ItemTooltip.onEmptySlotTooltipShow(tooltip)
	if ItemTooltip.slotDisabled or not Prefs.flag("tooltip") then
		return
	end
	local owner = type(tooltip.GetOwner) == "function" and tooltip:GetOwner() or nil
	local slot = slotForOwner(owner)
	if slot == nil or equippedLinkFor(slot) ~= nil then
		return
	end
	local ok, err = pcall(ItemTooltip.appendEmptySlotSection, tooltip, slot, owner)
	if not ok then
		ItemTooltip.slotDisabled = true
		Theme.note(string.format(L.diagTooltipHookFailed, tostring(err)))
	end
end

function ItemTooltip.registerEmptySlotRefresh()
	if ItemTooltip.emptySlotRefreshRegistered then
		return
	end
	ItemTooltip.emptySlotRefreshRegistered = true
	if type(GameTooltip) ~= "table" or type(GameTooltip.HookScript) ~= "function" then
		return
	end
	GameTooltip:HookScript("OnTooltipCleared", ItemTooltip.onEmptySlotTooltipCleared)
	GameTooltip:HookScript("OnShow", ItemTooltip.onEmptySlotTooltipShow)
end

--- Redraws the hover on screen once GET_ITEM_INFO_RECEIVED brings an id it
--- waits for (its compare item, or a name it showed as "…"); an id arriving
--- after the player moved on redraws nothing.
function ItemTooltip.onItemInfoReceived(itemId, success)
	local hover = ItemTooltip.activeHover
	if not success or hover == nil then
		return
	end
	if hover.itemId ~= itemId and (hover.pending == nil or not hover.pending[itemId]) then
		return
	end
	ItemTooltip.resetCache()
	hover.redraw()
end

--- The one GET_ITEM_INFO_RECEIVED listener the tooltip needs. Idempotent.
function ItemTooltip.registerItemInfoRefresh()
	if ItemTooltip.itemInfoFrame ~= nil then
		return ItemTooltip.itemInfoFrame
	end
	local frame = CreateFrame("Frame", nil, UIParent)
	Theme.registerEvent(frame, "GET_ITEM_INFO_RECEIVED")
	frame:SetScript("OnEvent", function(_, _, itemId, success)
		ItemTooltip.onItemInfoReceived(itemId, success)
	end)
	ItemTooltip.itemInfoFrame = frame
	return frame
end

function ItemTooltip.hasProcessor()
	return type(TooltipDataProcessor) == "table"
		and type(TooltipDataProcessor.AddTooltipPostCall) == "function"
		and type(Enum) == "table"
		and type(Enum.TooltipDataType) == "table"
		and Enum.TooltipDataType.Item ~= nil
end

--- Hooks exactly one of the two tooltip shapes, never both: the modern
--- processor when the client has it, else the legacy script. Idempotent.
function ItemTooltip.register()
	if ItemTooltip.registered then
		return ItemTooltip.how
	end
	ItemTooltip.registered = true
	ItemTooltip.registerUnit()
	ItemTooltip.registerBisSlotButtons()
	ItemTooltip.registerEmptySlotRefresh()
	ItemTooltip.registerItemInfoRefresh()
	if ItemTooltip.hasProcessor() then
		TooltipDataProcessor.AddTooltipPostCall(Enum.TooltipDataType.Item, function(tooltip)
			ItemTooltip.onTooltip(tooltip, ItemTooltip.itemLinkFrom(tooltip))
		end)
		ItemTooltip.how = "processor"
		Theme.noteVerbose(string.format(L.diagItemTooltipHook, ItemTooltip.how))
		return ItemTooltip.how
	end
	if type(GameTooltip) == "table" and type(GameTooltip.HookScript) == "function" then
		GameTooltip:HookScript("OnTooltipSetItem", function(tooltip)
			ItemTooltip.onTooltip(tooltip, ItemTooltip.itemLinkFrom(tooltip))
		end)
		ItemTooltip.how = "legacy"
		Theme.noteVerbose(string.format(L.diagItemTooltipHook, ItemTooltip.how))
		return ItemTooltip.how
	end
	Theme.note(L.diagNoTooltipHook)
	ItemTooltip.how = nil
	return nil
end

ns.ItemTooltip = ItemTooltip
return ItemTooltip
