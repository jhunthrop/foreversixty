-- addon/ForeverSixty/surfaces/StatWeights.lua
-- The Stat weights block of the Forever Sixty window, one component for every tab that shows it (the
-- Gear tab's right column, the Rotation tab's left): the heading with where the weights come from, the
-- Stat, Weight and DPS per point columns, and haste apart per 1%. Weights follow the SimC convention:
-- each stat per rating point over the top stat's, so the top stat reads 1.00 (design/specs/
-- 2026-11-13-addon-hub.md 4.3, 4.4).
--
-- The pure model first (StatWeights.model), then the frames (build once, draw on each show).
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Compat = ns.Compat or require("Compat")
local Rules = ns.Rules or require("Rules")
local Progress = ns.Progress or require("Progress")
local HubFrame = ns.HubFrame or require("HubFrame")
local BuildsWindow = ns.BuildsWindow or require("BuildsWindow")

local StatWeights = {}

--- Pixel layout: a row's pitch, the Weight and DPS per point columns' widths (both right aligned, the
--- per point column at the block's right edge), the room above the haste line.
StatWeights.LAYOUT = { row = 18, weight = 56, perPoint = 96, hasteGap = 2 }
--- The most rows listed, the largest first.
StatWeights.MAX_ROWS = 6
--- A weight under this share of the top stat's is left out: it moves no choice.
StatWeights.FLOOR = 0.05
--- Attack power counts for ranged attacks too: where a spec weighs both alike, one row says it.
StatWeights.SAME_AS = { attack_power = "ranged_attack_power", ranged_attack_power = "attack_power" }
--- Keys that never get a row: haste is shown apart (Rules.HASTE_WEIGHTS, per 1%), a weapon's DPS is the
--- band's weapon term, not a stat on an item line.
local NO_ROW = { haste = true, weapon_dps = true }
local HASTE = {}
for _, stat in ipairs(Rules.HASTE_WEIGHTS) do
	HASTE[stat] = true
	NO_ROW[stat] = true
end

-- The pure model ----------------------------------------------------------------------------------

--- The stat's name in the web's words (L.weightsStatNames), else the tooltip's.
local function statLabel(stat)
	return L.weightsStatNames[stat] or L.statNames[stat] or stat
end

--- A rating stat's row says so ("Hit rating"); its hover gives the conversion.
local function rowLabel(stat)
	local label = statLabel(stat)
	return Rules.RATING_PER_PERCENT[stat] ~= nil and string.format(L.weightsRating, label) or label
end

local function rowHover(stat)
	local factor = Rules.RATING_PER_PERCENT[stat]
	if factor == nil then
		return nil
	end
	local word = statLabel(stat):lower()
	return { string.format(L.weightsRatingConversion, factor, word, word) }
end

--- Each weighted stat per rating point (a rating stat's weight is per 1%, Rules.RATING_PER_PERCENT),
--- largest first; a stat weighed exactly as its twin (SAME_AS) is left to the twin that is the reference.
local function perPointRows(weights, reference)
	local rows = {}
	for stat, weight in pairs(weights) do
		local twin = StatWeights.SAME_AS[stat]
		local duplicate = twin ~= nil and weights[twin] == weight and reference ~= stat
		if type(weight) == "number" and weight > 0 and not NO_ROW[stat] and not duplicate then
			rows[#rows + 1] = { stat = stat, point = weight / (Rules.RATING_PER_PERCENT[stat] or 1) }
		end
	end
	table.sort(rows, function(a, b)
		if a.point ~= b.point then
			return a.point > b.point
		end
		return a.stat < b.stat
	end)
	return rows
end

--- Rating haste's weight per 1%: it counts at melee and spell haste both (Gear.score's rule). nil
--- without either.
local function hasteWeight(weights)
	local total
	for stat in pairs(HASTE) do
		local weight = weights[stat]
		if type(weight) == "number" and weight > 0 then
			total = (total or 0) + weight
		end
	end
	return total
end

--- Where the weights come from: the companion's sim and its day; else the band they were measured at,
--- said only when it is not the level the tab shows (`shown`).
local function sourceNote(current, shown)
	if current.tag == "sim" then
		local stamp = type(current.generatedAt) == "string" and Compat.shortDate(current.generatedAt:sub(1, 10)) or nil
		return stamp ~= nil and string.format(L.weightsSimOn, stamp) or L.weightsSim
	end
	local band = current.measuredBand or Progress.MAX_LEVEL
	return band ~= shown and string.format(L.weightsShipped, band) or nil
end

--- The block from Gear.currentWeights' answer `current` at the level the tab shows, or nil without
--- weights: { rows = { stat, label, hover, scale, value, capped }, haste, unit, note }. `scale` is the
--- stat per point over the top stat's (the top stat is 1.00), `value` DPS per point (nil without a DPS
--- figure); rows under FLOOR are left out, at most MAX_ROWS. `haste` is DPS per 1% of haste, `unit`
--- "DPS" or "HPS", `note` the source. Pure.
function StatWeights.model(current, shown)
	if type(current) ~= "table" or type(current.weights) ~= "table" then
		return nil
	end
	local rows = perPointRows(current.weights, current.reference)
	if #rows == 0 then
		return nil
	end
	local dps, top, capped = tonumber(current.dpsPerPoint), rows[1].point, {}
	for _, stat in ipairs(type(current.caps) == "table" and current.caps or {}) do
		capped[stat] = true
	end
	local kept = {}
	for _, row in ipairs(rows) do
		local scale = row.point / top
		if scale >= StatWeights.FLOOR and #kept < StatWeights.MAX_ROWS then
			kept[#kept + 1] = { stat = row.stat, label = rowLabel(row.stat), hover = rowHover(row.stat), scale = scale,
				value = dps ~= nil and row.point * dps or nil, capped = capped[row.stat] == true }
		end
	end
	local haste = hasteWeight(current.weights)
	return { rows = kept, haste = (haste ~= nil and dps ~= nil) and haste * dps or nil,
		unit = current.healing and L.weightsHps or L.weightsDps, note = sourceNote(current, shown) }
end

--- A row's three cells as { text, colour key }: the head, or a stat row.
function StatWeights.cells(model, row)
	if row == nil then
		return { { L.weightsStat, "muted" }, { L.weightsWeight, "muted" },
			{ string.format(L.weightsPerPoint, model.unit), "muted" } }
	end
	return { { row.label, "body" },
		{ row.capped and L.weightsCapped or string.format("%.2f", row.scale), row.capped and "muted" or "strong" },
		{ row.value ~= nil and string.format("%.2f", row.value) or L.weightsNone, "muted" } }
end

-- Frames -------------------------------------------------------------------------------------------

local function hoverArea(parent)
	local over = Theme.createFrame("Frame", nil, parent)
	over:EnableMouse(true)
	over:SetScript("OnEnter", function(self)
		if type(self.lines) == "table" then
			Theme.showLines(self, self.lines, "ANCHOR_RIGHT")
		end
	end)
	over:SetScript("OnLeave", function()
		Theme.hideTooltip()
	end)
	return over
end

local function makeRow(parent)
	local row = { stat = HubFrame.line(parent, "body"), weight = HubFrame.line(parent, "strong"),
		value = HubFrame.line(parent, "muted"), over = hoverArea(parent) }
	row.weight:SetJustifyH("RIGHT")
	row.value:SetJustifyH("RIGHT")
	return row
end

local function hideRow(row)
	for _, region in pairs(row) do
		region:Hide()
	end
end

--- The block on `parent` (`width` wide), built once: the section heading, the row pool, the haste line.
function StatWeights.build(parent, width)
	local section = BuildsWindow.buildSection(parent, width, { text = HubFrame.line }, L.weightsTitle, nil)
	section.parent, section.width, section.rows = parent, width, {}
	section.haste = HubFrame.line(parent, "muted", width)
	return section
end

local function paint(region, cell)
	region:SetText(cell[1])
	region:SetTextColor(Theme.rgb(Theme.HEX[cell[2]]))
end

local function drawRow(section, index, y, cells, hover)
	local G, parent, width = StatWeights.LAYOUT, section.parent, section.width
	section.rows[index] = section.rows[index] or makeRow(parent)
	local row = section.rows[index]
	paint(row.stat, cells[1])
	row.stat:SetWidth(width - G.weight - G.perPoint)
	row.stat:ClearAllPoints()
	row.stat:SetPoint("TOPLEFT", parent, "TOPLEFT", 0, -y)
	for _, pair in ipairs({ { row.weight, cells[2], width - G.perPoint }, { row.value, cells[3], width } }) do
		paint(pair[1], pair[2])
		pair[1]:ClearAllPoints()
		pair[1]:SetPoint("TOPRIGHT", parent, "TOPLEFT", pair[3], -y)
	end
	for _, region in ipairs({ row.stat, row.weight, row.value }) do
		region:Show()
	end
	row.over.lines = hover
	if hover ~= nil then
		row.over:ClearAllPoints()
		row.over:SetPoint("TOPLEFT", parent, "TOPLEFT", 0, -y)
		row.over:SetSize(width, G.row)
		row.over:Show()
	else
		row.over:Hide()
	end
	return y + G.row
end

--- Hide every part of the block.
function StatWeights.hide(section)
	section.heading:Hide()
	section.note:Hide()
	section.haste:Hide()
	for _, row in ipairs(section.rows) do
		hideRow(row)
	end
end

--- Draw `model` (StatWeights.model) at `y` with the tab's own `fit` and `textHeight` for the heading's
--- note; the y under the block, or nil with no model (the block hidden).
function StatWeights.draw(section, y, model, fit, textHeight)
	StatWeights.hide(section)
	if model == nil then
		return nil
	end
	local G = StatWeights.LAYOUT
	y = BuildsWindow.drawSection(section, section.parent, section.width, y, model.note, fit, textHeight)
	y = drawRow(section, 1, y, StatWeights.cells(model, nil), nil)
	for index, row in ipairs(model.rows) do
		y = drawRow(section, index + 1, y, StatWeights.cells(model, row), row.hover)
	end
	if model.haste ~= nil then
		section.haste:SetText(string.format(L.weightsHaste, model.haste, model.unit))
		section.haste:ClearAllPoints()
		section.haste:SetPoint("TOPLEFT", section.parent, "TOPLEFT", 0, -(y + G.hasteGap))
		section.haste:Show()
		y = y + G.hasteGap + textHeight(section.haste)
	end
	return y
end

ns.StatWeights = StatWeights
return StatWeights
