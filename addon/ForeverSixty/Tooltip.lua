-- addon/ForeverSixty/Tooltip.lua
-- Two lines on an item's own tooltip: whether the loaded build wants it,
-- and whether it beats what is worn, by Gear's own scoring.
--
-- Pure functions first (plannedLine, upgradeLine, lines) -- what the specs
-- cover without a real GameTooltip. The guarded hook that draws them onto
-- an actual tooltip is Tooltip.register() and friends, added in the next
-- task.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Talents = ns.Talents or require("Talents")
local Gear = ns.Gear or require("Gear")
local Export = ns.Export or require("Export")
local Compat = ns.Compat or require("Compat")
local Theme = ns.Theme or require("Theme")
local Prefs = ns.Prefs or require("Prefs")
local Follow = ns.Follow or require("Follow")

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

--- "Planned for your <slot>" when the loaded build wants this exact item.
function Tooltip.plannedLine(build, itemId)
	if build == nil or itemId == nil then
		return nil
	end
	for _, entry in ipairs(build.gear or {}) do
		if entry.itemId == itemId then
			return string.format(L.tooltipPlanned, entry.slot)
		end
	end
	return nil
end

--- "Upgrade for <slot>: +N by our weights" or "Not an upgrade", scored
--- against whatever is worn in each slot the item could fill -- the best
--- (highest-delta) slot wins when more than one fits (rings, trinkets, one-
--- and two-handers). Gear's own scoring; no second scorer. Uses the
--- player's own class (there may be no build loaded at all), unlike
--- Gear.upgrades, which is always called with one already loaded.
function Tooltip.upgradeLine(data, itemLink)
	local classSlug = Talents.playerClassSlug()
	if classSlug == nil or itemLink == nil or data == nil then
		return nil
	end
	local ranks = Talents.readRanks(data)
	local spec = Gear.specOf(data, classSlug, ranks)
	local weights = spec ~= nil and data.weights[spec] or nil
	if weights == nil then
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
			best = { slot = slot, delta = delta }
		end
	end
	if best == nil then
		return nil
	end
	if best.delta > 0 then
		return string.format(L.tooltipUpgrade, best.slot, best.delta)
	end
	return L.tooltipNotUpgrade
end

--- "<Stat> capped: more is wasted" for each stat this item carries that
--- `weightsMessage` (a "weights" inbox message -- Codec.inboxMessages(
--- inbox, key, "weights"), section 4's breakpoint call-out) already
--- names in its `caps` list. With no message, or one carrying no caps
--- at all, this answers no lines rather than guessing -- the same "no
--- data, no line" rule Tooltip.upgradeLine already follows for a spec
--- with no weights.
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

--- Below this level a character has no ladder band yet (design: "below 10
--- shows nothing" -- pipeline.addonbis.BIS_LEVEL_BANDS' own floor).
Tooltip.BIS_MIN_LEVEL = 10

--- data/pipeline/addonlua.py's SOURCE_KIND_CODES, decoded back for the
--- advanced-detail line. Kept beside BIS_SLOT_BUTTONS as the addon's own
--- half of a contract data/pipeline/addonbis.py's SOURCE_KIND_CODES
--- states in full; tests/tooltip_bis_spec.lua pins this table's key set
--- against that same set of source kinds independently, the way
--- rotation_spec.lua pins LEVEL_BANDS against ladder.go's literal.
Tooltip.SOURCE_KIND_NAMES = {
	Q = "Quest",
	D = "Dungeon",
	C = "Crafted",
	R = "Reputation",
	P = "PvP",
	W = "World",
	A = "Raid",
}

--- The highest bis band <= level, tolerating gaps in a partially-generated
--- spec (the nightly bis.yml workflow may not have finished every band
--- yet). Below BIS_MIN_LEVEL, or with no bis table for this spec at all:
--- nil, nil.
local function bisBandFor(specBis, level)
	if type(specBis) ~= "table" or type(level) ~= "number" or level < Tooltip.BIS_MIN_LEVEL then
		return nil, nil
	end
	local rounded = math.floor(level / 5) * 5
	if rounded > 60 then
		rounded = 60
	end
	for candidate = rounded, Tooltip.BIS_MIN_LEVEL, -5 do
		if specBis[candidate] ~= nil then
			return candidate, specBis[candidate]
		end
	end
	return nil, nil
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

--- "(equipped)" beats "(new at <band>)" when both would apply -- an item
--- already worn cannot also be new to put on.
local function bisTag(itemId, band, equippedItemId, newIds)
	if equippedItemId == itemId then
		return L.tooltipBisEquipped
	end
	if isAmong(newIds, itemId) then
		return string.format(L.tooltipBisNew, band)
	end
	return ""
end

--- item id -> true once RequestLoadItemDataByID (or the legacy global) has
--- been asked for it, so re-hovering the same uncached item never asks
--- the client twice.
Tooltip.bisRequested = {}

local function requestItemLoad(itemId)
	if Tooltip.bisRequested[itemId] then
		return
	end
	Tooltip.bisRequested[itemId] = true
	local fn = (type(C_Item) == "table" and C_Item.RequestLoadItemDataByID)
		or _G.RequestLoadItemDataByID
	if type(fn) == "function" then
		fn(itemId)
	end
end

--- The item's link once GetItemInfo knows it, else the "item:<id>" form
--- (still a valid hyperlink target) while the client fills its cache --
--- RequestLoadItemDataByID/GetItemInfo asked for the id exactly once,
--- never per hover.
function Tooltip.bisItemLink(itemId)
	local _, link = Compat.itemInfo(itemId)
	if link ~= nil then
		return link
	end
	requestItemLoad(itemId)
	return "item:" .. itemId
end

--- The BiS hover section for one site slot: a header ("Best in slot ·
--- <band> · <spec>") and the slot's pick as an item link, tagged
--- "(equipped)" or "(new at <band>)"; advanced detail
--- (Prefs.flag("advancedDetail")) adds the source kind. Empty (no
--- section) while InCombatLockdown, with no resolvable class or spec, no
--- bis band at or above BIS_MIN_LEVEL for this character's level, or the
--- band names nothing for this slot -- design/lane-bis-hover-addon.md
--- item 2's own rules.
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
	-- entry is { itemId, sourceKindCode } -- addonlua.py's positional,
	-- unnamed-field encoding (the lane brief's size budget).
	local itemId, sourceCode = entry[1], entry[2]
	local tag = bisTag(itemId, band, itemIdOf(equippedLinkFor(slot)), newAtBand(data, spec, band, faction))
	local link = Tooltip.bisItemLink(itemId)
	local lines = {
		string.format(L.tooltipBisHeader, band, spec),
		tag == "" and link or (link .. " " .. tag),
	}
	if Prefs.flag("advancedDetail") then
		lines[#lines + 1] = string.format(L.tooltipBisSource, Tooltip.SOURCE_KIND_NAMES[sourceCode] or sourceCode)
	end
	return lines
end

--- Every line this addon ever adds, 0 or more. Pure; the hook this
--- file grows next only draws what this returns. weightsMessage is
--- optional (nil is "no weights message for this character yet"), so
--- every existing three-argument call site keeps behaving exactly as
--- it did before Tooltip.capLines existed.
function Tooltip.lines(data, build, itemLink, weightsMessage)
	local lines = {}
	local planned = Tooltip.plannedLine(build, itemIdOf(itemLink))
	if planned ~= nil then
		lines[#lines + 1] = planned
	end
	local upgrade = Tooltip.upgradeLine(data, itemLink)
	if upgrade ~= nil then
		lines[#lines + 1] = upgrade
	end
	for _, capLine in ipairs(Tooltip.capLines(itemLink, weightsMessage)) do
		lines[#lines + 1] = capLine
	end
	-- The BiS section only for an item that IS what this character wears
	-- right now in one of its own slots -- a bag item or someone else's
	-- gear names no site slot and gets nothing here.
	if data ~= nil then
		local slot = Tooltip.equippedSlotFor(itemLink)
		if slot ~= nil then
			for _, bisLine in ipairs(Tooltip.bisLines(data, slot)) do
				lines[#lines + 1] = bisLine
			end
		end
	end
	return lines
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

local function cachedLines(itemLink)
	local cached = Tooltip.cache[itemLink]
	if cached ~= nil then
		return cached
	end
	local lines = Tooltip.lines(Tooltip.data, Follow.build, itemLink, Tooltip.weightsMessage)
	Tooltip.cache[itemLink] = lines
	return lines
end

local function addLines(tooltip, itemLink)
	local lines = cachedLines(itemLink)
	if #lines == 0 then
		return
	end
	tooltip:AddLine(L.addonName, Theme.rgb(Theme.HEX.gold))
	for _, line in ipairs(lines) do
		tooltip:AddLine(line, Theme.rgb(Theme.HEX.body))
	end
	tooltip:Show()
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

--- An EMPTY slot's own BiS section: PaperDollItemSlotButton_OnEnter's own
--- GameTooltip:SetInventoryItem call adds nothing and never shows the
--- tooltip when the slot has no item, so OnTooltipSetItem (Tooltip.lines,
--- above) never fires for it at all -- this hook is the only place an
--- empty slot's hover gets drawn. A slot that DOES have an item is left
--- alone here: its section already came from the item hook, and drawing
--- it twice is not this hook's job.
function Tooltip.onSlotEnter(slot, button)
	if Tooltip.slotDisabled or not Prefs.flag("tooltip") then
		return
	end
	if equippedLinkFor(slot) ~= nil then
		return
	end
	local ok, err = pcall(function()
		local lines = Tooltip.bisLines(Tooltip.data, slot)
		if #lines == 0 then
			return
		end
		GameTooltip:SetOwner(button, "ANCHOR_RIGHT")
		GameTooltip:ClearLines()
		for _, line in ipairs(lines) do
			GameTooltip:AddLine(line, Theme.rgb(Theme.HEX.body))
		end
		GameTooltip:Show()
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
