-- addon/ForeverSixty/Tooltip.lua
-- The item tooltip section, to the standard docs/tenets.md sets: a single
-- brand header, a coloured upgrade verdict, a structured BiS row (an icon
-- and a quality-coloured link on the left, a tag on the right), and a
-- muted source line -- never a bare "Source: Crafted".
--
-- Pure functions first (verdict, bisLines, sections) -- what the specs
-- cover without a real GameTooltip. The guarded hooks that draw them onto
-- an actual tooltip (Tooltip.register() and friends) come after.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Talents = ns.Talents or require("Talents")
local Gear = ns.Gear or require("Gear")
local Export = ns.Export or require("Export")
local Compat = ns.Compat or require("Compat")
local Theme = ns.Theme or require("Theme")
local Prefs = ns.Prefs or require("Prefs")

local Tooltip = {}

--- The site slot name -> the client's equipped-item slot id, built once
--- from Export's own table. Export.lua is off limits to edit in this lane
--- (file ownership), so this reads its exported table rather than keeping
--- a second copy of the numbers.
local SLOT_IDS = {}
for _, entry in ipairs(Export.INVENTORY_SLOTS) do
	SLOT_IDS[entry.slot] = entry.id
end

--- An item link's id. Mirrors the local itemIdOf in Export.lua, which is
--- private to that file and off limits to edit here; kept to three lines
--- so the day the two can share one function costs little.
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
--- unit's item, or nothing equipped there). Iterates every SLOT_IDS entry
--- rather than reading a "which slot was this tooltip set from" API,
--- because no such API is documented for SetInventoryItem's item variant;
--- comparing the exact link (enchant and suffix included) against what is
--- worn is exact where it matters -- whether THIS item is THIS character's
--- current pick for the slot.
function Tooltip.equippedSlotFor(itemLink)
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

--- pipeline.addonbis.BIS_LEVEL_BANDS' own spacing: one BiS band every ten levels.
local BIS_BAND_STEP = 10

--- The highest band key at or below `level` in a band-keyed table (bands
--- run 20..60 step 10, and a build in progress may skip one) -- shared by
--- the BiS lookup and the verdict's own band-weights lookup, so both
--- round down to the same rung. nil below `floor`, or with no table (or no
--- entry at all) to look in.
local function highestBandAtOrBelow(bands, level, floor)
	if type(bands) ~= "table" or type(level) ~= "number" or level < floor then
		return nil
	end
	local rounded = math.min(60, math.floor(level / BIS_BAND_STEP) * BIS_BAND_STEP)
	for candidate = rounded, floor, -BIS_BAND_STEP do
		if bands[candidate] ~= nil then
			return candidate
		end
	end
	return nil
end

--- Below this level a character has no BiS band yet (design: "below 20
--- shows nothing" -- pipeline.addonbis.BIS_LEVEL_BANDS' own floor).
Tooltip.BIS_MIN_LEVEL = 20

--- This spec's stat weights at `level`: the nightly-measured band weights
--- (ns.Data.bis_weights, lane addon-tooltip-polish item 4) when the
--- leveling-bis run has actually measured this spec/band, else the
--- curated static table (ns.Data.weights) -- the same "band data first,
--- static fallback" rule Tooltip.bisLines already follows for gear picks.
--- A band entry that exists but measured nothing significant (weights ==
--- {}) is "not measured yet", not "measured as zero", so it also falls
--- through to the static table.
function Tooltip.weightsFor(data, spec, level)
	if spec == nil then
		return nil
	end
	local bisWeights = type(data.bis_weights) == "table" and data.bis_weights[spec] or nil
	local band = highestBandAtOrBelow(bisWeights, level, Tooltip.BIS_MIN_LEVEL)
	local measured = band ~= nil and bisWeights[band] or nil
	if type(measured) == "table" and next(measured) ~= nil then
		return measured
	end
	return data.weights[spec]
end

--- Below this percentage magnitude a delta reads as noise, not a verdict --
--- lane addon-tooltip-polish's own call; nothing in the design states a
--- number, and this is small enough that two items only a stray point of a
--- minor stat apart still read as a Sidegrade rather than a coin-flip
--- Upgrade/Downgrade.
Tooltip.SIDEGRADE_THRESHOLD_PERCENT = 0.5

--- { kind, color, text } for a slot's own best candidate delta, kind one
--- of "upgrade" | "downgrade" | "sidegrade", color one of Theme.HEX's own
--- keys. With something actually worn in the best slot (equippedScore ~=
--- 0), the verdict is a percentage of that baseline, which is what the
--- owner's screenshot review asked for: a chest with more agility than
--- the one worn must obviously read as an upgrade. With nothing worn
--- there at all (an empty slot, or a worn item that scores exactly zero
--- under these weights), a percentage has no baseline to be a percentage
--- OF, so this falls back to the older absolute-points wording instead of
--- dividing by zero.
local function verdictFor(slot, delta, equippedScore)
	if equippedScore ~= 0 then
		local percent = delta / math.abs(equippedScore) * 100
		if math.abs(percent) < Tooltip.SIDEGRADE_THRESHOLD_PERCENT then
			return { kind = "sidegrade", color = "muted", text = L.tooltipVerdictSidegrade }
		end
		if percent > 0 then
			return {
				kind = "upgrade", color = "success",
				text = string.format(L.tooltipVerdictUpgrade, percent),
			}
		end
		return {
			kind = "downgrade", color = "warning",
			text = string.format(L.tooltipVerdictDowngrade, -percent),
		}
	end
	if delta > 0 then
		return {
			kind = "upgrade", color = "success",
			text = string.format(L.tooltipUpgrade, slot, delta),
		}
	end
	return { kind = "sidegrade", color = "muted", text = L.tooltipNotUpgrade }
end

--- The upgrade verdict for a hovered item, scored against whatever is worn
--- in each slot the item could fill -- the best (highest-delta) slot wins
--- when more than one fits (rings, trinkets, one- and two-handers). Gear's
--- own scoring; no second scorer. Uses the player's own class (there may
--- be no build loaded at all). nil with nothing to say: no resolvable
--- class, no item link, no data table, no stat weights for the spec at
--- any band, or an item with no equip slot at all (a reagent, say).
function Tooltip.verdict(data, itemLink)
	local classSlug = Talents.playerClassSlug()
	if classSlug == nil or itemLink == nil or data == nil then
		return nil
	end
	local ranks = Talents.readRanks(data)
	local spec = Gear.specOf(data, classSlug, ranks)
	local weights = Tooltip.weightsFor(data, spec, Compat.playerLevel())
	if weights == nil or next(weights) == nil then
		return nil
	end
	local _, _, _, equipLocation = Compat.itemInfoInstant(itemLink)
	local slots = Gear.SLOTS_BY_EQUIP_LOCATION[equipLocation]
	if slots == nil then
		return nil
	end
	local itemScore = Gear.score(Gear.statsOf(itemLink), weights)
	local best
	for _, slot in ipairs(slots) do
		local equippedScore = Gear.score(Gear.statsOf(equippedLinkFor(slot)), weights)
		local delta = itemScore - equippedScore
		if best == nil or delta > best.delta then
			best = { slot = slot, delta = delta, equippedScore = equippedScore }
		end
	end
	if best == nil then
		return nil
	end
	return verdictFor(best.slot, best.delta, best.equippedScore)
end

--- "<Stat> capped: more is wasted" for each stat this item carries that
--- `weightsMessage` (a "weights" inbox message -- Codec.inboxMessages(
--- inbox, key, "weights"), section 4's breakpoint call-out) already
--- names in its `caps` list. With no message, or one carrying no caps
--- at all, this answers no lines rather than guessing -- the same "no
--- data, no line" rule Tooltip.verdict already follows for a spec with
--- no weights.
function Tooltip.capLines(itemLink, weightsMessage)
	local lines = {}
	if itemLink == nil or type(weightsMessage) ~= "table" or type(weightsMessage.caps) ~= "table" then
		return lines
	end
	local capped = {}
	for _, stat in ipairs(weightsMessage.caps) do
		capped[stat] = true
	end
	for stat in pairs(Gear.statsOf(itemLink)) do
		if capped[stat] then
			lines[#lines + 1] = string.format(L.tooltipCapped, stat)
		end
	end
	-- Gear.statsOf is a plain table, so iterating it in pairs() order
	-- is not deterministic; sorted once here rather than left to
	-- whatever pairs() happened to visit first, so two runs against
	-- the same item read the same tooltip.
	table.sort(lines)
	return lines
end

--- Site slot name -> the Blizzard PaperDoll button's global frame name
--- (Gear.lua's own slot vocabulary). Shirt and Tabard carry no stats and
--- are not in that vocabulary at all, so they get no BiS hover.
Tooltip.BIS_SLOT_BUTTONS = {
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

--- data/pipeline/addonlua.py's SOURCE_KIND_CODES, decoded back for the
--- muted source line. Kept beside BIS_SLOT_BUTTONS as the addon's own
--- half of a contract data/pipeline/addonbis.py's SOURCE_KIND_CODES
--- states in full; tests/tooltip_bis_spec.lua pins this table's key set
--- against that same set of source kinds independently, the way
--- rotation_spec.lua pins LEVEL_BANDS against ladder.go's literal.
Tooltip.SOURCE_KIND_NAMES = {
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

--- The highest bis band <= level, tolerating gaps in a partially-generated
--- spec (the nightly bis.yml workflow may not have finished every band
--- yet). Below BIS_MIN_LEVEL, or with no bis table for this spec at all:
--- nil, nil.
local function bisBandFor(specBis, level)
	local band = highestBandAtOrBelow(specBis, level, Tooltip.BIS_MIN_LEVEL)
	if band == nil then
		return nil, nil
	end
	return band, specBis[band]
end

--- "alliance"/"horde", the bis table's own faction keys, or nil on a test
--- double with no UnitFactionGroup.
local function playerFaction()
	if type(UnitFactionGroup) ~= "function" then
		return nil
	end
	local token = UnitFactionGroup("player")
	if token == "Alliance" then
		return "alliance"
	end
	if token == "Horde" then
		return "horde"
	end
	return nil
end

--- The item ids newly best at `band` for `faction`, or {} with no
--- bis_new table for this spec/band/faction at all.
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

--- The BiS row's own right-aligned tag, as { text, color }: "equipped"
--- (green) beats "new at <band>" (gold) beats the plain "BiS · lvl <band>"
--- (muted) default -- an item already worn cannot also be new to put on.
local function bisTag(itemId, band, equippedItemId, newIds)
	if equippedItemId == itemId then
		return { text = L.tooltipBisEquipped, color = "success" }
	end
	if isAmong(newIds, itemId) then
		return { text = string.format(L.tooltipBisNew, band), color = "gold" }
	end
	return { text = string.format(L.tooltipBisDefaultTag, band), color = "muted" }
end

--- The item's link once GetItemInfo knows it, else the "item:<id>" form
--- (still a valid hyperlink target) while the client fills its cache --
--- Compat.displayLink's own rule, kept as a named function here since
--- design/lane-bis-hover-addon.md's own tests call it by this name.
function Tooltip.bisItemLink(itemId)
	return Compat.displayLink(itemId)
end

--- The muted source line -- "<kind> · <place>", e.g. "Dungeon · Wailing
--- Caverns · Mutanus the Devourer" or "Crafted · Leatherworking" -- from
--- the pick's source-kind code and its short place label (addonbis.py's
--- own `source`, positional element 3). nil with no label at all: a bare
--- "Source: Crafted" is never shown again (lane addon-tooltip-polish item
--- 1), so a pick this data has no place for simply gets no source line,
--- rather than the kind alone.
local function sourceLine(sourceCode, sourceLabel)
	if type(sourceLabel) ~= "string" or sourceLabel == "" then
		return nil
	end
	local kindName = Tooltip.SOURCE_KIND_NAMES[sourceCode] or sourceCode
	-- "Wailing Caverns: Mutanus the Devourer" -> "Wailing Caverns · Mutanus
	-- the Devourer": report.go's own "instance: boss" punctuation, recut to
	-- the addon's own separator so the whole line reads as one style.
	return string.format(L.tooltipBisSourceLabel, kindName, (sourceLabel:gsub(": ", " · ")))
end

--- The BiS hover's inline icon size -- a touch larger than a list row's
--- (Theme.SIZES.iconSize), since a tooltip line has more room than a
--- packed row.
Tooltip.BIS_ICON_SIZE = 18

--- The BiS section for one site slot, as a list of draw ops (Tooltip.draw
--- is what turns these into real tooltip calls): a "doubleline" op for the
--- pick itself (icon + quality-coloured link on the left, the tag on the
--- right), then plain "line" ops for the item level (once cached) and the
--- muted source line (when the data has one). Empty (no ops) while
--- InCombatLockdown, with no resolvable class or spec, no bis band at or
--- above BIS_MIN_LEVEL for this character's level, or the band names
--- nothing for this slot -- design/lane-bis-hover-addon.md item 2's own
--- rules.
---
--- The second return is the recommended item ({ itemId, link }), for the
--- caller to show the client's own item tooltip for it beside GameTooltip
--- (Theme.showCompareTooltip) -- nil on every path that answers no
--- section at all, so a caller need only check it, not re-derive it.
function Tooltip.bisLines(data, slot)
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
	local band, bandData = bisBandFor(specBis, Compat.playerLevel())
	if band == nil then
		return {}
	end
	local faction = playerFaction()
	local factionBis = faction ~= nil and bandData[faction] or nil
	local entry = factionBis ~= nil and factionBis[slot] or nil
	if entry == nil then
		return {}
	end
	-- entry is { itemId, sourceKindCode, sourceLabel } -- addonlua.py's
	-- positional, unnamed-field encoding (the lane brief's size budget);
	-- sourceLabel is left out of the table literal entirely for a pick
	-- with none, so it comes back nil rather than "".
	local itemId, sourceCode, sourceLabel = entry[1], entry[2], entry[3]
	local tag = bisTag(itemId, band, itemIdOf(equippedLinkFor(slot)), newAtBand(data, spec, band, faction))
	local link = Tooltip.bisItemLink(itemId)
	local icon = Theme.inlineIcon(Compat.itemIcon(itemId), Tooltip.BIS_ICON_SIZE)
	local ops = {
		{ kind = "doubleline", left = icon .. " " .. link, right = tag.text, rightColor = tag.color },
	}
	local level = Compat.itemLevel(itemId)
	if level ~= nil then
		ops[#ops + 1] = { kind = "line", text = string.format(L.tooltipBisItemLevel, level), color = "muted" }
	end
	local source = sourceLine(sourceCode, sourceLabel)
	if source ~= nil then
		ops[#ops + 1] = { kind = "line", text = source, color = "muted" }
	end
	return ops, { itemId = itemId, link = link }
end

--- The site slot an item's hover should draw its BiS row for --
--- Gear.SLOTS_BY_EQUIP_LOCATION's own candidate list for the item's
--- equip location (round-2 section 10 ruling 1, adopted): ANY item that maps
--- to a tracked slot gets the row, not only the exact item currently
--- worn there. A location with more than one candidate slot (rings,
--- trinkets, one- and two-handers) prefers the slot the item is
--- actually equipped in, when it is equipped in one of them, so a ring
--- worn in finger2 draws finger2's own pick rather than finger1's; the
--- first candidate otherwise, which is the only slot there is for every
--- other equip location. nil for an item whose equip location this
--- band does not track at all (a shirt, a tabard, a bag, ammo), or with
--- no link to look up.
function Tooltip.bisSlotFor(itemLink)
	if itemLink == nil then
		return nil
	end
	local _, _, _, equipLocation = Compat.itemInfoInstant(itemLink)
	local slots = Gear.SLOTS_BY_EQUIP_LOCATION[equipLocation]
	if slots == nil then
		return nil
	end
	local equippedSlot = Tooltip.equippedSlotFor(itemLink)
	for _, slot in ipairs(slots) do
		if slot == equippedSlot then
			return slot
		end
	end
	return slots[1]
end

--- Every op this addon ever draws for one item, 0 or more: the upgrade
--- verdict first (Tooltip.verdict, colour by meaning), then the BiS
--- section for the slot Tooltip.bisSlotFor resolves -- any item whose
--- equip location maps to a tracked slot, not only the exact-equipped
--- link (round-2 section 10 ruling 1: a bag item being compared, an empty
--- character-pane slot and the worn piece itself all draw the same way
--- now) -- then any capped-stat call-outs. Pure; Tooltip.draw is what
--- turns this into real tooltip calls. weightsMessage is optional (nil
--- is "no weights message for this character yet").
---
--- The second return is the BiS section's own recommended item, exactly
--- as Tooltip.bisLines answers it (nil with no BiS section at all), so a
--- caller need only check it, not re-derive it.
function Tooltip.sections(data, itemLink, weightsMessage)
	local ops = {}
	local verdict = Tooltip.verdict(data, itemLink)
	if verdict ~= nil then
		ops[#ops + 1] = { kind = "line", text = verdict.text, color = verdict.color }
	end
	local target
	if data ~= nil then
		local slot = Tooltip.bisSlotFor(itemLink)
		if slot ~= nil then
			local bisOps, bisTarget = Tooltip.bisLines(data, slot)
			for _, op in ipairs(bisOps) do
				ops[#ops + 1] = op
			end
			target = bisTarget
		end
	end
	for _, capLine in ipairs(Tooltip.capLines(itemLink, weightsMessage)) do
		ops[#ops + 1] = { kind = "line", text = capLine, color = "body" }
	end
	return ops, target
end

--- The generated data table, set once by Options.register() exactly like
--- Window.data and MinimapButton.data -- never required directly, so a
--- spec can hand the hook a fixture instead of the real Data.lua, and the
--- hook does nothing (rather than erroring) before login has set it.
Tooltip.data = nil

--- The "weights" inbox message for the character currently logged in --
--- set once, the same way and by the same caller as Tooltip.data, once
--- something reads the companion's inbox at login and resolves it with
--- Codec.inboxMessages(inbox, Export.characterKey(), "weights") to the
--- one message (if any) addressed to this character. Wiring that call
--- is a documented hook rather than done in this file: it belongs where
--- Tooltip.data itself is set (Options.register(), a different lane's
--- file this wave), not in Tooltip.lua, which only ever reads the
--- result. Nil until then, and Tooltip.capLines answers no lines --
--- exactly this feature's behaviour before the hook is wired.
Tooltip.weightsMessage = nil

--- Computed once per item link for the session: the mouse crossing the
--- same item repeatedly must not re-run the scoring walk every time.
Tooltip.cache = {}

function Tooltip.resetCache()
	Tooltip.cache = {}
	return Tooltip.cache
end

--- Tooltip.sections(Tooltip.data, itemLink, Tooltip.weightsMessage),
--- cached by link. Caching the ops AND the target together (one call)
--- rather than the ops alone is what lets the item hook and the compare
--- tooltip share one scoring walk instead of two -- see git history for
--- the version that scored bisLines a second time just to read `target`.
local function cachedSections(itemLink)
	local cached = Tooltip.cache[itemLink]
	if cached ~= nil then
		return cached.ops, cached.target
	end
	local ops, target = Tooltip.sections(Tooltip.data, itemLink, Tooltip.weightsMessage)
	Tooltip.cache[itemLink] = { ops = ops, target = target }
	return ops, target
end

--- The BiS hover currently on screen, if it names an item id the client
--- had to ask for (Compat.requestedItems) -- set by whichever hook below
--- drew it. GET_ITEM_INFO_RECEIVED (Tooltip.onItemInfoReceived) redraws
--- it in place once that id's data arrives, rather than asking the
--- player to move the mouse away and back to see the icon, the coloured
--- link and the item level that were missing the first time.
Tooltip.activeHover = nil

--- Colour one op's field by Theme.HEX key, falling back to body text for
--- an op that names no colour at all.
local function opColor(key)
	return Theme.rgb(Theme.HEX[key or "body"])
end

--- The header strip icon's inline size -- Theme.SIZES.rotationHeaderIcon
--- is the same "a small icon beside a header line" job the rotation
--- card's own strip already does; reused rather than a second constant
--- for one more small header icon.
Tooltip.HEADER_ICON_SIZE = Theme.SIZES.rotationHeaderIcon

--- Draws `ops` (Tooltip.sections' or Tooltip.bisLines' own output) onto a
--- real tooltip, with the brand header first -- the addon's small icon
--- texture beside "Forever Sixty" in the brand gold, docs/tenets.md's own
--- ask -- and never drawn at all when there is nothing to say. Shared by
--- the item hook and the empty-slot section (both hover paths item 2
--- covers), so the two draw exactly the same way.
local function drawOps(tooltip, ops)
	if #ops == 0 then
		return
	end
	local hr, hg, hb = opColor("gold")
	tooltip:AddLine(Theme.inlineIcon(Theme.MEDIA.minimapIcon, Tooltip.HEADER_ICON_SIZE) .. " " .. L.addonName,
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

local function addLines(tooltip, itemLink)
	local ops, target = cachedSections(itemLink)
	if #ops == 0 then
		return
	end
	drawOps(tooltip, ops)
	tooltip:Show()
	if target == nil then
		return
	end
	Theme.showCompareTooltip(target.itemId, target.link)
	Tooltip.activeHover = {
		itemId = target.itemId,
		redraw = function()
			-- Re-running Blizzard's own SetHyperlink re-fires the
			-- TooltipDataProcessor/OnTooltipSetItem hook that called
			-- Tooltip.onTooltip in the first place, so the tooltip is
			-- rebuilt from scratch with the data that just arrived
			-- rather than appending a second copy of this addon's lines.
			if type(tooltip.SetHyperlink) == "function"
				and type(tooltip.IsShown) == "function" and tooltip:IsShown() then
				tooltip:SetHyperlink(itemLink)
			end
		end,
	}
end

--- One guarded body for both hook shapes below. A Lua error inside a
--- tooltip hook breaks every tooltip in the game (found in game, twice),
--- so a failure here turns the hook off instead of raising a second time.
function Tooltip.onTooltip(tooltip, itemLink)
	if Tooltip.disabled or itemLink == nil or not Prefs.flag("tooltip") then
		return
	end
	local ok, err = pcall(addLines, tooltip, itemLink)
	if not ok then
		Tooltip.disabled = true
		Theme.note(string.format(L.diagTooltipHookFailed, tostring(err)))
	end
end

function Tooltip.itemLinkFrom(tooltip)
	if type(tooltip) ~= "table" or type(tooltip.GetItem) ~= "function" then
		return nil
	end
	local ok, _, link = pcall(tooltip.GetItem, tooltip)
	if not ok then
		return nil
	end
	return link
end

--- The Ratings line on a player's tooltip: the Raider.IO moment. Same
--- guard as the item hook: a failure turns this hook off, never raises.
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

function Tooltip.onUnitTooltip(tooltip, unit)
	if Tooltip.unitDisabled or unit == nil or not Prefs.flag("tooltip") then
		return
	end
	local ok, err = pcall(addUnitLines, tooltip, unit)
	if not ok then
		Tooltip.unitDisabled = true
		Theme.note(string.format(L.diagTooltipHookFailed, tostring(err)))
	end
end

function Tooltip.unitFrom(tooltip)
	if type(tooltip) ~= "table" or type(tooltip.GetUnit) ~= "function" then
		return nil
	end
	local ok, _, unit = pcall(tooltip.GetUnit, tooltip)
	return ok and unit or nil
end

function Tooltip.hasUnitProcessor()
	return Tooltip.hasProcessor() and Enum.TooltipDataType.Unit ~= nil
end

--- The unit hook, registered beside the item hook by the same rule.
function Tooltip.registerUnit()
	if Tooltip.unitRegistered then
		return Tooltip.unitHow
	end
	Tooltip.unitRegistered = true
	if Tooltip.hasUnitProcessor() then
		TooltipDataProcessor.AddTooltipPostCall(Enum.TooltipDataType.Unit, function(tooltip)
			Tooltip.onUnitTooltip(tooltip, Tooltip.unitFrom(tooltip))
		end)
		Tooltip.unitHow = "processor"
		return Tooltip.unitHow
	end
	if type(GameTooltip) == "table" and type(GameTooltip.HookScript) == "function" then
		GameTooltip:HookScript("OnTooltipSetUnit", function(tooltip)
			Tooltip.onUnitTooltip(tooltip, Tooltip.unitFrom(tooltip))
		end)
		Tooltip.unitHow = "legacy"
		return Tooltip.unitHow
	end
	Tooltip.unitHow = nil
	return nil
end

--- The Blizzard slot-button frame this character's own BIS_SLOT_BUTTONS
--- names for `slot`, or nil (this client renamed it, or a test double
--- never defined it).
local function slotButton(slot)
	local buttonName = Tooltip.BIS_SLOT_BUTTONS[slot]
	return buttonName ~= nil and _G[buttonName] or nil
end

--- The site slot `owner` (a GameTooltip owner frame) names, or nil: the
--- reverse of slotButton, by frame identity rather than by name -- this
--- client's slot buttons are whatever BIS_SLOT_BUTTONS says they are,
--- named global or not.
local function slotForOwner(owner)
	if owner == nil then
		return nil
	end
	for slot in pairs(Tooltip.BIS_SLOT_BUTTONS) do
		if slotButton(slot) == owner then
			return slot
		end
	end
	return nil
end

--- Whether Tooltip.appendEmptySlotSection has already drawn its lines onto
--- a given tooltip since its last OnTooltipCleared -- item 2's own "never
--- appended twice" mark. Kept in a weak-keyed table of this file's own
--- rather than as a field on the tooltip itself: GameTooltip is a global
--- this addon only ever reads (luacheck rightly flags an assignment into
--- one of its fields as writing a read-only global), and a test double is
--- a plain table a spec may reuse across examples. Weak keys let a
--- tooltip this addon no longer references (a test's own, once the spec
--- ends) be collected rather than pinned here forever.
local appendedSection = setmetatable({}, { __mode = "k" })

--- An EMPTY slot's own BiS section, appended onto `tooltip` (already
--- owned and, for a fresh hover, already cleared by the caller) unless it
--- already carries one -- appendedSection, cleared by
--- Tooltip.onEmptySlotTooltipCleared below, is the mark item 2's own fix
--- turns on so Blizzard's own OnUpdate re-render (which rebuilds
--- GameTooltip without going through this addon's OnEnter hook at all)
--- gets exactly one copy of this section, never zero and never two.
function Tooltip.appendEmptySlotSection(tooltip, slot, button)
	if appendedSection[tooltip] then
		return
	end
	local ops, target = Tooltip.bisLines(Tooltip.data, slot)
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
	Tooltip.activeHover = {
		itemId = target.itemId,
		redraw = function()
			appendedSection[tooltip] = nil
			Tooltip.onSlotEnter(slot, button)
		end,
	}
end

--- The character-frame slot buttons' own OnEnter: an EMPTY slot's BiS
--- section (design/lane-bis-hover-addon.md item 2). A slot that DOES have
--- an item is left alone here: its section already came from the item
--- hook, and drawing it twice is not this hook's job.
function Tooltip.onSlotEnter(slot, button)
	if Tooltip.slotDisabled or not Prefs.flag("tooltip") then
		return
	end
	if equippedLinkFor(slot) ~= nil then
		return
	end
	local ok, err = pcall(function()
		GameTooltip:SetOwner(button, "ANCHOR_RIGHT")
		GameTooltip:ClearLines()
		-- Cleared explicitly rather than only relying on OnTooltipCleared
		-- firing for this same-owner SetOwner/ClearLines pair -- unverified
		-- without a live client (see the lane report).
		appendedSection[GameTooltip] = nil
		Tooltip.appendEmptySlotSection(GameTooltip, slot, button)
	end)
	if not ok then
		Tooltip.slotDisabled = true
		Theme.note(string.format(L.diagTooltipHookFailed, tostring(err)))
	end
end

--- Hooks every BIS_SLOT_BUTTONS global this client actually has (a test
--- double, or a future client that renames one, simply gets no hook for
--- that slot rather than an error). Idempotent, like registerUnit.
function Tooltip.registerBisSlotButtons()
	if Tooltip.bisButtonsRegistered then
		return
	end
	Tooltip.bisButtonsRegistered = true
	for slot, buttonName in pairs(Tooltip.BIS_SLOT_BUTTONS) do
		local button = _G[buttonName]
		if type(button) == "table" and type(button.HookScript) == "function" then
			button:HookScript("OnEnter", function(self)
				Tooltip.onSlotEnter(slot, self)
			end)
		end
	end
end

--- GameTooltip's own OnTooltipCleared: clears the "already appended" mark
--- (item 2's own fix) so the NEXT OnShow -- whether this addon's own
--- Tooltip.onSlotEnter or Blizzard's OnUpdate re-render -- knows it must
--- draw the section again.
function Tooltip.onEmptySlotTooltipCleared(tooltip)
	appendedSection[tooltip] = nil
end

--- GameTooltip's own OnShow: the fix for the empty-slot flash
--- (design/lane-addon-tooltip-polish.md item 2). Blizzard's
--- PaperDollItemSlotButton_OnUpdate re-runs the hovered slot's OnEnter
--- logic on a timer WITHOUT firing the button's OnEnter script -- this
--- addon's own hook in registerBisSlotButtons never sees that re-render at
--- all -- so GameTooltip is rebuilt with none of this addon's lines a
--- moment after they were drawn. OnShow fires for every one of those
--- re-renders (and for this addon's own first draw, which is what the
--- "already appended" mark is for): whenever GameTooltip's owner is one
--- of the 17 slot buttons and it carries no item, this appends the
--- section again -- a no-op when it is already there.
function Tooltip.onEmptySlotTooltipShow(tooltip)
	if Tooltip.slotDisabled or not Prefs.flag("tooltip") then
		return
	end
	local owner = type(tooltip.GetOwner) == "function" and tooltip:GetOwner() or nil
	local slot = slotForOwner(owner)
	if slot == nil or equippedLinkFor(slot) ~= nil then
		return
	end
	local ok, err = pcall(Tooltip.appendEmptySlotSection, tooltip, slot, owner)
	if not ok then
		Tooltip.slotDisabled = true
		Theme.note(string.format(L.diagTooltipHookFailed, tostring(err)))
	end
end

--- Hooks GameTooltip's OnTooltipCleared/OnShow once (item 2's own fix).
--- Idempotent, like registerBisSlotButtons.
function Tooltip.registerEmptySlotRefresh()
	if Tooltip.emptySlotRefreshRegistered then
		return
	end
	Tooltip.emptySlotRefreshRegistered = true
	if type(GameTooltip) ~= "table" or type(GameTooltip.HookScript) ~= "function" then
		return
	end
	GameTooltip:HookScript("OnTooltipCleared", Tooltip.onEmptySlotTooltipCleared)
	GameTooltip:HookScript("OnShow", Tooltip.onEmptySlotTooltipShow)
end

--- Redraws whichever hover Tooltip.activeHover names, once GET_ITEM_INFO_
--- RECEIVED confirms the client now has that id's data -- only while it
--- is still the id showing, so an id arriving after the player moved on
--- redraws nothing.
function Tooltip.onItemInfoReceived(itemId, success)
	if not success or Tooltip.activeHover == nil or Tooltip.activeHover.itemId ~= itemId then
		return
	end
	Tooltip.resetCache()
	Tooltip.activeHover.redraw()
end

--- The one GET_ITEM_INFO_RECEIVED listener this addon needs, own frame
--- rather than Options.EVENTS: that file is a different lane's this wave
--- (Tooltip.weightsMessage's own comment), and this event is Tooltip's
--- alone to act on. Idempotent, like register().
function Tooltip.registerItemInfoRefresh()
	if Tooltip.itemInfoFrame ~= nil then
		return Tooltip.itemInfoFrame
	end
	local frame = CreateFrame("Frame", nil, UIParent)
	Theme.registerEvent(frame, "GET_ITEM_INFO_RECEIVED")
	frame:SetScript("OnEvent", function(_, _, itemId, success)
		Tooltip.onItemInfoReceived(itemId, success)
	end)
	Tooltip.itemInfoFrame = frame
	return frame
end

function Tooltip.hasProcessor()
	return type(TooltipDataProcessor) == "table"
		and type(TooltipDataProcessor.AddTooltipPostCall) == "function"
		and type(Enum) == "table"
		and type(Enum.TooltipDataType) == "table"
		and Enum.TooltipDataType.Item ~= nil
end

--- Hooks exactly one of the two tooltip shapes, never both: the modern
--- processor when the client has it, else the legacy script. Idempotent,
--- so Options.register() can call it plainly every load.
function Tooltip.register()
	if Tooltip.registered then
		return Tooltip.how
	end
	Tooltip.registered = true
	Tooltip.registerUnit()
	Tooltip.registerBisSlotButtons()
	Tooltip.registerEmptySlotRefresh()
	Tooltip.registerItemInfoRefresh()
	if Tooltip.hasProcessor() then
		TooltipDataProcessor.AddTooltipPostCall(Enum.TooltipDataType.Item, function(tooltip)
			Tooltip.onTooltip(tooltip, Tooltip.itemLinkFrom(tooltip))
		end)
		Tooltip.how = "processor"
		return Tooltip.how
	end
	if type(GameTooltip) == "table" and type(GameTooltip.HookScript) == "function" then
		GameTooltip:HookScript("OnTooltipSetItem", function(tooltip)
			Tooltip.onTooltip(tooltip, Tooltip.itemLinkFrom(tooltip))
		end)
		Tooltip.how = "legacy"
		return Tooltip.how
	end
	Theme.note(L.diagNoTooltipHook)
	Tooltip.how = nil
	return nil
end

ns.Tooltip = Tooltip
return Tooltip
