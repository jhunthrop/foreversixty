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

ns.Tooltip = Tooltip
return Tooltip
