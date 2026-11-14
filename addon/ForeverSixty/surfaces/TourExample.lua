-- addon/ForeverSixty/surfaces/TourExample.lua
-- Set up's picture (design/specs/2026-11-13-addon-hub.md 4.1, board hub-me-setup.png): step 2 says "Hover an
-- item you wear" and names one, and under the steps the client's own tooltip of that worn item shows with
-- the line Forever Sixty adds to it, exactly as a hover would draw it. The item is the worn one Gear's top
-- upgrade would replace (its tooltip names that upgrade), else the main hand, chest or legs, else any worn
-- item. Pure choice plus one reader and one embedded tooltip.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local Theme = ns.Theme or require("Theme")
local Compat = ns.Compat or require("Compat")
local Export = ns.Export or require("Export")
local ItemTooltip = ns.ItemTooltip or require("ItemTooltip")
local GearTab = ns.GearTab or require("GearTab")

local TourExample = {}

TourExample.TOOLTIP_NAME = "ForeverSixtySetUpTooltip"
--- The slots tried, in order, when no upgrade names a worn item.
TourExample.EXAMPLE_SLOTS = { "main_hand", "chest", "legs" }

local SLOT_IDS = {}
for _, entry in ipairs(Export.INVENTORY_SLOTS) do
	SLOT_IDS[entry.slot] = entry.id
end

--- The worn slot to show: the first upgrade row (GearTab.upgrades' order) over a worn item, else the first
--- of EXAMPLE_SLOTS worn, else the first worn slot in the paper doll's order; nil with nothing worn. Pure on
--- `equipped` (slot -> { itemId }) and `upgradeRows` (cells with `slot`).
function TourExample.choose(equipped, upgradeRows)
	equipped = equipped or {}
	for _, cell in ipairs(upgradeRows or {}) do
		if equipped[cell.slot] ~= nil then
			return cell.slot
		end
	end
	for _, slot in ipairs(TourExample.EXAMPLE_SLOTS) do
		if equipped[slot] ~= nil then
			return slot
		end
	end
	for _, entry in ipairs(Export.INVENTORY_SLOTS) do
		if equipped[entry.slot] ~= nil then
			return entry.slot
		end
	end
	return nil
end

--- Gear's upgrade rows for the character's own band, or none without a spec or a list.
local function upgradeRows(input)
	if input.spec == nil then
		return {}
	end
	local band = GearTab.ownBand(input)
	if band == nil or not GearTab.hasList(input, band) then
		return {}
	end
	return GearTab.upgrades(GearTab.doll(input, band)).rows
end

--- The example from the client: { slot, slotId, itemId, name }, or nil with nothing worn or a read that
--- failed (Set up then draws step 2 without its item).
function TourExample.read(data)
	local ok, input = pcall(GearTab.read, data)
	if not ok or type(input) ~= "table" then
		return nil
	end
	local scored, rows = pcall(upgradeRows, input)
	local slot = TourExample.choose(input.equipped, scored and rows or {})
	if slot == nil then
		return nil
	end
	local itemId = input.equipped[slot].itemId
	return { slot = slot, slotId = SLOT_IDS[slot], itemId = itemId, name = Compat.itemName(itemId) }
end

--- The embedded tooltip, once, in `parent`: the client's own GameTooltip look, on the window's strata,
--- never opening a compare tooltip (ItemTooltip.embed).
function TourExample.ensureTooltip(parent)
	if TourExample.tooltip ~= nil then
		return TourExample.tooltip
	end
	local tooltip = Theme.createFrame("GameTooltip", TourExample.TOOLTIP_NAME, parent, "gameTooltip")
	tooltip:Hide()
	ItemTooltip.embed(tooltip)
	TourExample.tooltip = tooltip
	return tooltip
end

--- Shows the worn item in `slotId` on `tooltip` at (`x`, `y`) in `parent`, at least `width` wide; the
--- tooltip's height, 0 when the client has nothing in that slot.
function TourExample.showTooltip(tooltip, parent, slotId, x, y, width)
	if tooltip == nil or slotId == nil or type(tooltip.SetInventoryItem) ~= "function" then
		return 0
	end
	tooltip:SetOwner(parent, "ANCHOR_NONE")
	tooltip:ClearAllPoints()
	-- Placed before it is filled: the item hook skips an unplaced ANCHOR_NONE tooltip as a scanner.
	tooltip:SetPoint("TOPLEFT", parent, "TOPLEFT", x, -y)
	if type(tooltip.SetFrameStrata) == "function" and type(parent.GetFrameStrata) == "function" then
		tooltip:SetFrameStrata(parent:GetFrameStrata())
	end
	if type(tooltip.SetMinimumWidth) == "function" then
		tooltip:SetMinimumWidth(width)
	end
	if not tooltip:SetInventoryItem("player", slotId) and type(tooltip.NumLines) == "function"
		and tooltip:NumLines() == 0 then
		tooltip:Hide()
		return 0
	end
	tooltip:Show()
	local height = type(tooltip.GetHeight) == "function" and tooltip:GetHeight() or 0
	return type(height) == "number" and math.ceil(height) or 0
end

function TourExample.hideTooltip(tooltip)
	if tooltip ~= nil then
		tooltip:Hide()
	end
end

ns.TourExample = TourExample
return TourExample
