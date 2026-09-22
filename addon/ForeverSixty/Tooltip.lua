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

--- Both lines this addon ever adds, 0 to 2 of them. Pure; the hook this
--- file grows next only draws what this returns.
function Tooltip.lines(data, build, itemLink)
	local lines = {}
	local planned = Tooltip.plannedLine(build, itemIdOf(itemLink))
	if planned ~= nil then
		lines[#lines + 1] = planned
	end
	local upgrade = Tooltip.upgradeLine(data, itemLink)
	if upgrade ~= nil then
		lines[#lines + 1] = upgrade
	end
	return lines
end

--- The generated data table, set once by Options.register() exactly like
--- Window.data and MinimapButton.data -- never required directly, so a
--- spec can hand the hook a fixture instead of the real Data.lua, and the
--- hook does nothing (rather than erroring) before login has set it.
Tooltip.data = nil

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
	local lines = Tooltip.lines(Tooltip.data, Follow.build, itemLink)
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
	local name = type(UnitName) == "function" and select(1, UnitName(unit)) or nil
	local realm = type(UnitName) == "function" and select(2, UnitName(unit)) or nil
	if realm == "" then
		realm = nil
	end
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
