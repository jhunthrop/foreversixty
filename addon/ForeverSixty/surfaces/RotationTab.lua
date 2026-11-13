-- addon/ForeverSixty/surfaces/RotationTab.lua
-- The Rotation tab of the Forever Sixty window (design/specs/2026-11-13-addon-hub.md 4.4,
-- hub-rotation.png): what do I press, and where did I lose damage last raid?
--
-- Left: Priority, the spec's lines at a level (the stepper walks the levels the priority changes at,
-- the player's own by default), each the spell's icon, its name and its condition in plain words, a
-- line the player cannot have trained yet dimmed with "at level N"; under it the stat weights the
-- verdicts read (Gear.currentWeights) with where they come from. Right: How you played it (the inbox
-- last_raid: the mean followed, every boss least followed first with its biggest loss) and Lost across
-- the night (the losses summed by kind), with a link to Me.
--
-- Pure models first (RotationTab.priority, played, weightRows, model), then the frames.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Compat = ns.Compat or require("Compat")
local Talents = ns.Talents or require("Talents")
local Export = ns.Export or require("Export")
local Follow = ns.Follow or require("Follow")
local Gear = ns.Gear or require("Gear")
local Rules = ns.Rules or require("Rules")
local LastRaid = ns.LastRaid or require("LastRaid")
local Progress = ns.Progress or require("Progress")
local Rotation = ns.Rotation or require("Rotation")
local Hub = ns.Hub or require("Hub")
local BuildsWindow = ns.BuildsWindow or require("BuildsWindow")

local RotationTab = {}

--- Pixel layout: a priority row's pitch for a one-line condition and each line more, the number
--- column, the spell icon and the gap after it; a table row and its head; the stepper's buttons and
--- the level between them; the room between sections; the columns of the weights and the boss tables.
RotationTab.LAYOUT = {
	row = 36, rowLine = 15, number = 16, icon = 24, iconGap = 8,
	tableRow = 18, tableHead = 16, step = 20, stepLevel = 64, stepGap = 4, sectionGap = 20,
	weightStat = 232, weightValue = 60,
	boss = 136, followed = 56, cellGap = 8,
}
--- A condition wraps to at most this many lines (Rotation.plainCondition cuts the rest).
RotationTab.CONDITION_LINES = 2
--- The most bosses listed; the rest are counted under the table (spec 3.1: a tab never scrolls).
RotationTab.MAX_BOSSES = 12
--- The most stat weights listed, the largest first.
RotationTab.MAX_WEIGHTS = 6
--- Attack power counts for ranged attacks too: where a spec weighs both alike, one row says it.
RotationTab.SAME_AS = { attack_power = "ranged_attack_power", ranged_attack_power = "attack_power" }
--- The weights keys that are haste: shown apart, per 1% (the SimC convention, spec 4.3).
RotationTab.HASTE = { haste = true, melee_haste = true, spell_haste = true }

--- Injection point: the shipped data (ns.Data when nil); tests set it.
RotationTab.data = nil

-- Pure models --------------------------------------------------------------------------------

--- The stepper's stops: the levels `spec`'s priority changes at and the player's own level, ascending
--- and unique.
function RotationTab.stops(data, spec, level)
	local seen, stops = {}, {}
	local levels = Rotation.changeLevels(data, spec)
	if type(level) == "number" then
		levels[#levels + 1] = level
	end
	for _, value in ipairs(levels) do
		if not seen[value] then
			seen[value] = true
			stops[#stops + 1] = value
		end
	end
	table.sort(stops)
	return stops
end

--- The stop before (`step` -1) or after (+1) `shown` in `stops`, or nil at that end.
function RotationTab.step(stops, shown, step)
	local found
	for _, value in ipairs(stops) do
		if (step < 0 and value < shown) or (step > 0 and value > shown and found == nil) then
			found = value
		end
	end
	return found
end

--- A condition as its line shows it: Rotation.plainCondition's words, sentence case, no full stop.
function RotationTab.conditionText(text, measure, width)
	local plain, lines = Rotation.plainCondition(text, measure, width)
	if plain == nil then
		return nil, 0
	end
	plain = plain:gsub("%.$", "")
	return plain:sub(1, 1):upper() .. plain:sub(2), lines or 1
end

--- The priority at `shown` (the stepper's level) for a character at `level`: each line { number,
--- spellId, name, condition, conditionLines, full, at }, `at` the level a line the character cannot
--- have trained yet is learned at (Rotation.learnedAt above `level`), else nil. Empty without a band.
function RotationTab.priority(data, spec, shown, level, measure, width)
	local band = Rotation.bandAt(data, spec, shown)
	local rows = {}
	for index, line in ipairs(band ~= nil and band.lines or {}) do
		local condition, lines = RotationTab.conditionText(line.condition, measure, width)
		local learned = Rotation.learnedAt(line.spellId, data)
		rows[index] = { number = index, spellId = line.spellId, name = line.name, condition = condition,
			conditionLines = condition ~= nil and lines or 0, full = RotationTab.conditionText(line.condition),
			at = (learned ~= nil and type(level) == "number" and learned > level) and learned or nil }
	end
	return rows
end

--- A weight per rating point: a rating stat's weight is per 1% (Rules.RATING_PER_PERCENT).
local function perPoint(stat, weight)
	return weight / (Rules.RATING_PER_PERCENT[stat] or 1)
end

--- The stat's name in the web's words (L.rotStatNames), else the tooltip's.
local function statLabel(stat)
	return L.rotStatNames[stat] or L.statNames[stat] or stat
end

--- Each weighted stat but haste, per rating point, and haste's own weight: a stat the spec weighs
--- exactly as its twin (SAME_AS) is left to the twin that is the reference.
local function perPointRows(current)
	local rows, haste = {}, nil
	for stat, weight in pairs(current.weights) do
		local twin = RotationTab.SAME_AS[stat]
		local duplicate = twin ~= nil and current.weights[twin] == weight and current.reference ~= stat
		if type(weight) == "number" and weight > 0 and not duplicate then
			if RotationTab.HASTE[stat] then
				haste = math.max(haste or 0, weight)
			else
				rows[#rows + 1] = { stat = stat, point = perPoint(stat, weight) }
			end
		end
	end
	table.sort(rows, function(a, b)
		if a.point ~= b.point then
			return a.point > b.point
		end
		return a.stat < b.stat
	end)
	return rows, haste
end

--- Where the weights come from: the companion's sim and its day; else the band they were measured at,
--- said only when it is not the level the stepper shows (`shown`).
local function sourceNote(current, shown)
	if current.tag == "sim" then
		local stamp = type(current.generatedAt) == "string" and Compat.shortDate(current.generatedAt:sub(1, 10)) or nil
		return stamp ~= nil and string.format(L.rotWeightsSimOn, stamp) or L.rotWeightsSim
	end
	local band = current.measuredBand or Progress.MAX_LEVEL
	return band ~= shown and string.format(L.rotWeightsShipped, band) or nil
end

--- A stat's row label: a rating stat says so ("Hit rating"), its hover gives the conversion.
local function rowLabel(stat)
	local label = statLabel(stat)
	return Rules.RATING_PER_PERCENT[stat] ~= nil and string.format(L.rotRating, label) or label
end

local function rowHover(stat)
	local factor = Rules.RATING_PER_PERCENT[stat]
	if factor == nil then
		return nil
	end
	local word = statLabel(stat):lower()
	return { string.format(L.rotRatingConversion, factor, word, word) }
end

--- The weights section from Gear.currentWeights' answer `current`, or nil: { rows = { stat, label,
--- scale, value, capped } with `scale` per point over the top stat's (the top stat is 1.00) and `value`
--- DPS per point (nil without a DPS figure), largest first, at most MAX_WEIGHTS; haste (DPS per 1%,
--- apart); unit ("DPS", "HPS"); note (the source) }.
function RotationTab.weightRows(current, shown)
	if type(current) ~= "table" or type(current.weights) ~= "table" then
		return nil
	end
	local rows, haste = perPointRows(current)
	if #rows == 0 then
		return nil
	end
	local dps, top, capped = tonumber(current.dpsPerPoint), rows[1].point, {}
	for _, stat in ipairs(type(current.caps) == "table" and current.caps or {}) do
		capped[stat] = true
	end
	local kept = {}
	for index = 1, math.min(#rows, RotationTab.MAX_WEIGHTS) do
		local row = rows[index]
		kept[index] = { stat = row.stat, label = rowLabel(row.stat), hover = rowHover(row.stat), scale = row.point / top,
			value = dps ~= nil and row.point * dps or nil, capped = capped[row.stat] == true }
	end
	return { rows = kept, haste = (haste ~= nil and dps ~= nil) and haste * dps or nil, note = sourceNote(current, shown),
		unit = current.healing and L.rotHps or L.rotDps }
end

--- The tab's whole model from `input` { data, spec, followed, buildName, level, shown, lastRaid,
--- companion, weights, measure, width }: { spec, specName, followed, buildName, companion, stops,
--- shown, previous, next, lines, played, weights }.
function RotationTab.model(input)
	local stops = RotationTab.stops(input.data, input.spec, input.level)
	local shown = input.shown or input.level or stops[1]
	return {
		spec = input.spec, specName = Follow.specName(input.spec), followed = input.followed,
		buildName = input.buildName, companion = input.companion, stops = stops, shown = shown,
		previous = RotationTab.step(stops, shown, -1), next = RotationTab.step(stops, shown, 1),
		lines = RotationTab.priority(input.data, input.spec, shown, input.level, input.measure, input.width),
		played = LastRaid.played(input.lastRaid, RotationTab.MAX_BOSSES),
		weights = RotationTab.weightRows(input.weights, shown),
	}
end

--- The header's figures (spec 3.2): the followed build, then "followed 85% on Tuesday" while a last
--- raid has a mean. With no build the Priority heading names the spec, so the header does not.
function RotationTab.figuresOf(model)
	local parts = { model.buildName }
	local played = model.played
	if played ~= nil and played.mean ~= nil and played.day ~= nil then
		parts[#parts + 1] = string.format(L.rotFiguresFollowed, played.mean, played.day)
	end
	return parts
end

-- Live inputs --------------------------------------------------------------------------------

local function dataOf()
	return RotationTab.data or Hub.data or ns.Data
end

--- The live inputs of RotationTab.model, or nil when the client cannot say who is logged in.
function RotationTab.inputs(shown, measure, width)
	local data = dataOf()
	local classSlug = Talents.playerClassSlug()
	if type(data) ~= "table" or classSlug == nil then
		return nil
	end
	local ranks = Talents.readRanks(data)
	local build, level, key = Follow.build, Compat.playerLevel(), Export.characterKey()
	local spec, followed = Rotation.specFor(data, build, ranks, classSlug)
	if spec == nil then
		return nil
	end
	return { data = data, spec = spec, followed = followed, level = level, shown = shown,
		buildName = followed and Follow.buildName(data, build) or nil,
		lastRaid = LastRaid.newest(_G.ForeverSixtyInbox, key), companion = type(_G.ForeverSixtyInbox) == "table",
		-- The weights at the level the stepper shows, so the priority and its stats read one level.
		weights = Gear.currentWeights(data, spec, _G.ForeverSixtyInbox, key, shown or level),
		measure = measure, width = width }
end

--- The header's figures for the hub (Hub.register `figures`).
function RotationTab.figures()
	local input = RotationTab.inputs()
	return input ~= nil and RotationTab.figuresOf(RotationTab.model(input)) or {}
end

-- Frames -------------------------------------------------------------------------------------

local G = RotationTab.LAYOUT

--- A line at the window's text size: `fontKey` "highlight" is the semibold face; never wraps.
local function text(parent, hexKey, width, fontKey)
	local region = Theme.fontString(parent, "OVERLAY", fontKey or "normal")
	local path, _, flags = region:GetFont()
	if path ~= nil then
		region:SetFont(path, Theme.SIZES.dialogText, flags)
	end
	region:SetTextColor(Theme.rgb(Theme.HEX[hexKey or "body"]))
	region:SetJustifyH("LEFT")
	if width ~= nil then
		region:SetWidth(width)
	end
	if type(region.SetWordWrap) == "function" then
		region:SetWordWrap(false)
	end
	return region
end

local DEPS = { text = text }

local function textHeight(region)
	local height = type(region.GetStringHeight) == "function" and region:GetStringHeight() or nil
	return (type(height) == "number" and height > 0) and height or G.rowLine
end

local function fit(region, raw, room)
	region:SetWidth(room)
	region:SetText(raw or "")
end

local function place(region, parent, x, y)
	region:ClearAllPoints()
	region:SetPoint("TOPLEFT", parent, "TOPLEFT", x, -y)
	region:Show()
end

local function placeRight(region, parent, right, y)
	region:ClearAllPoints()
	region:SetPoint("TOPRIGHT", parent, "TOPLEFT", right, -y)
	region:SetJustifyH("RIGHT")
	region:Show()
end

--- A hover area over a row: `lines()` for a plain tooltip, or `spellId` for the client's own.
local function hoverArea(parent)
	local over = Theme.createFrame("Frame", nil, parent)
	over:EnableMouse(true)
	over:SetScript("OnEnter", function(self)
		if self.spellId ~= nil and type(GameTooltip) == "table" and type(GameTooltip.SetSpellByID) == "function" then
			GameTooltip:SetOwner(self, "ANCHOR_RIGHT")
			GameTooltip:SetSpellByID(self.spellId)
			GameTooltip:Show()
		elseif type(self.lines) == "table" then
			Theme.showLines(self, self.lines, "ANCHOR_RIGHT")
		end
	end)
	over:SetScript("OnLeave", function()
		Theme.hideTooltip()
	end)
	return over
end

local function setHover(over, parent, y, width, height)
	over:ClearAllPoints()
	over:SetPoint("TOPLEFT", parent, "TOPLEFT", 0, -y)
	over:SetSize(width, height)
	over:Show()
end

--- A column of the body: `x` from the body's left, `width` wide, from the body padding down.
local function column(body, x, width)
	local frame = Theme.createFrame("Frame", nil, body)
	frame:SetPoint("TOPLEFT", body, "TOPLEFT", x, -Hub.LAYOUT.bodyPad)
	frame:SetSize(width, 1)
	return frame
end

--- Get row `index` from `pool`, made by `make(parent)` the first time.
local function pooled(pool, index, make, parent)
	pool[index] = pool[index] or make(parent)
	return pool[index]
end

local function hideFrom(pool, first)
	for index = first, #pool do
		for _, region in pairs(pool[index]) do
			region:Hide()
		end
	end
end

local function makePriorityRow(parent)
	local width = Theme.SIZES.windowLeft - G.number - G.icon - G.iconGap * 2
	local row = { number = text(parent, "muted", G.number), icon = Theme.newTexture(parent, "ARTWORK"),
		name = text(parent, "strong", width, "highlight"), at = text(parent, "muted"),
		condition = text(parent, "muted", width), over = hoverArea(parent) }
	row.number:SetJustifyH("CENTER")
	row.icon:SetSize(G.icon, G.icon)
	if type(row.condition.SetWordWrap) == "function" then
		row.condition:SetWordWrap(true)
	end
	if type(row.condition.SetMaxLines) == "function" then
		row.condition:SetMaxLines(RotationTab.CONDITION_LINES)
	end
	return row
end

--- One priority row at `y`; the y under it.
local function drawPriorityRow(row, line, parent, y)
	local width = Theme.SIZES.windowLeft
	local x = G.number + G.iconGap
	local height = G.row + G.rowLine * math.max(0, line.conditionLines - 1)
	local alpha = line.at ~= nil and Theme.ALPHA.dim or 1
	row.number:SetText(tostring(line.number))
	place(row.number, parent, 0, y + (G.icon - textHeight(row.number)) / 2)
	row.icon:SetTexture(Compat.spellTexture(line.spellId) or Theme.UNKNOWN_ICON)
	place(row.icon, parent, x, y)
	row.name:SetText(line.name)
	place(row.name, parent, x + G.icon + G.iconGap, y - 1)
	row.at:SetText(line.at ~= nil and string.format(L.rotAtLevel, line.at) or "")
	placeRight(row.at, parent, width, y - 1)
	row.condition:SetText(line.condition or "")
	place(row.condition, parent, x + G.icon + G.iconGap, y - 1 + G.rowLine)
	for _, region in ipairs({ row.number, row.icon, row.name }) do
		region:SetAlpha(alpha)
	end
	row.over.spellId, row.over.lines = line.spellId, nil
	setHover(row.over, parent, y, width, height)
	return y + height
end

--- The stepper right of the Priority heading: the client buttons and the level between them.
local function buildStepper(view, parent)
	local stepper = {}
	stepper.previous = Widgets.clientButton(parent, L.rotPrevious, function()
		RotationTab.show({ shown = view.model and view.model.previous })
	end)
	stepper.nextButton = Widgets.clientButton(parent, L.rotNext, function()
		RotationTab.show({ shown = view.model and view.model.next })
	end)
	for _, button in ipairs({ stepper.previous, stepper.nextButton }) do
		button:SetSize(G.step, G.step)
	end
	stepper.level = text(parent, "strong", G.stepLevel, "highlight")
	stepper.level:SetJustifyH("CENTER")
	return stepper
end

local function drawStepper(stepper, parent, model, y)
	local right = Theme.SIZES.windowLeft
	stepper.nextButton:ClearAllPoints()
	stepper.nextButton:SetPoint("TOPRIGHT", parent, "TOPLEFT", right, -y + 1)
	stepper.level:ClearAllPoints()
	stepper.level:SetPoint("RIGHT", stepper.nextButton, "LEFT", -G.stepGap, 0)
	stepper.level:SetText(string.format(L.rotLevel, model.shown))
	stepper.previous:ClearAllPoints()
	stepper.previous:SetPoint("RIGHT", stepper.level, "LEFT", -G.stepGap, 0)
	for _, pair in ipairs({ { stepper.previous, model.previous }, { stepper.nextButton, model.next } }) do
		local button, target = pair[1], pair[2]
		if target ~= nil then
			button:Enable()
		else
			button:Disable()
		end
		button:SetAlpha(target ~= nil and 1 or Theme.ALPHA.disabled)
		button:Show()
	end
	stepper.level:Show()
end

--- A table row of three cells: the first left aligned, the second right aligned, the third either.
local function makeTableRow(parent)
	return { first = text(parent, "strong"), second = text(parent, "strong"), third = text(parent, "muted"),
		over = hoverArea(parent) }
end

local function paintCell(region, cell)
	region:SetText(cell ~= nil and cell[1] or "")
	region:SetTextColor(Theme.rgb(Theme.HEX[cell ~= nil and cell[2] or "muted"]))
end

--- Draw `row` at `y` with `layout` { firstWidth, second (right edge), third (left x or right edge),
--- thirdRight }; the y under it.
local function drawTableRow(row, parent, y, cells, layout, hover)
	paintCell(row.first, cells[1])
	fit(row.first, cells[1] and cells[1][1], layout.firstWidth)
	place(row.first, parent, 0, y)
	paintCell(row.second, cells[2])
	placeRight(row.second, parent, layout.second, y)
	paintCell(row.third, cells[3])
	if layout.thirdRight then
		placeRight(row.third, parent, layout.third, y)
	else
		fit(row.third, cells[3] and cells[3][1], layout.thirdWidth)
		place(row.third, parent, layout.third, y)
	end
	if hover ~= nil then
		row.over.spellId, row.over.lines = nil, hover
		setHover(row.over, parent, y, layout.width, G.tableRow)
	else
		row.over:Hide()
	end
	return y + G.tableRow
end

local function weightsLayout()
	local width = Theme.SIZES.windowLeft
	return { width = width, firstWidth = G.weightStat, second = G.weightStat + G.weightValue,
		third = width, thirdRight = true }
end

local function bossLayout()
	local width = Theme.SIZES.windowRight
	local third = G.boss + G.followed + G.cellGap
	return { width = width, firstWidth = G.boss, second = G.boss + G.followed, third = third,
		thirdWidth = width - third }
end

local function lostLayout()
	local width = Theme.SIZES.windowRight
	return { width = width, firstWidth = width - G.followed, second = width, third = width, thirdRight = true }
end

--- Draw rows `entries` (each a list of cells and an optional hover) from `pool` at `y`; the y under them.
local function drawRows(pool, parent, y, entries, layout)
	for index, entry in ipairs(entries) do
		y = drawTableRow(pooled(pool, index, makeTableRow, parent), parent, y, entry.cells, layout, entry.hover)
	end
	hideFrom(pool, #entries + 1)
	return y
end

--- The weights section on the left at `y`; the y under it (y itself without weights).
local function drawWeights(view, weights, y)
	local parent, section = view.left, view.weights
	for _, region in ipairs({ section.heading, section.note, section.haste }) do
		region:Hide()
	end
	hideFrom(view.weightRows, 1)
	if weights == nil then
		return y
	end
	y = BuildsWindow.drawSection(section, parent, Theme.SIZES.windowLeft, y + G.sectionGap, weights.note, fit,
		textHeight)
	local entries = { { cells = { { L.rotStat, "muted" }, { L.rotWeight, "muted" },
		{ string.format(L.rotPerPoint, weights.unit), "muted" } } } }
	for _, row in ipairs(weights.rows) do
		local scale = row.capped and L.rotCapped or string.format("%.2f", row.scale)
		entries[#entries + 1] = { hover = row.hover, cells = { { row.label, "body" },
			{ scale, row.capped and "muted" or "strong" },
			{ row.value ~= nil and string.format("%.2f", row.value) or L.rotNone, "muted" } } }
	end
	y = drawRows(view.weightRows, parent, y, entries, weightsLayout())
	if weights.haste ~= nil then
		section.haste:SetText(string.format(L.rotHaste, weights.haste, weights.unit))
		place(section.haste, parent, 0, y + 2)
		y = y + 2 + textHeight(section.haste)
	end
	return y
end

--- The left column: Priority with the stepper, the lines, then the stat weights at the level shown; its
--- height.
local function drawLeft(view, model)
	local parent = view.left
	local note = not model.followed and Theme.inlineColor("muted", string.format(L.rotFor, model.specName)) or nil
	view.priority.heading:SetText(L.rotPriority .. (note ~= nil and " " .. note or ""))
	local y = BuildsWindow.drawSection(view.priority, parent, Theme.SIZES.windowLeft, 0, nil, fit, textHeight)
	drawStepper(view.stepper, parent, model, 0)
	y = y + Theme.SIZES.gap
	for index, line in ipairs(model.lines) do
		y = drawPriorityRow(pooled(view.lines, index, makePriorityRow, parent), line, parent, y)
	end
	hideFrom(view.lines, #model.lines + 1)
	return drawWeights(view, model.weights, y)
end

--- How you played it's rows: the head, then each boss { name, followed, biggest loss } with its hover.
local function bossEntries(played)
	local entries = { { cells = { { L.rotBoss, "muted" }, { L.rotFollowed, "muted" }, { L.rotBiggestLoss, "muted" } } } }
	for _, boss in ipairs(played.bosses) do
		entries[#entries + 1] = { hover = boss.hover, cells = { { boss.short, "body" },
			{ boss.followed ~= nil and string.format(L.rotPercent, boss.followed) or L.rotNone,
				boss.followed ~= nil and "strong" or "muted" },
			{ boss.loss or L.rotNone, "muted" } } }
	end
	if played.mean ~= nil then
		-- The night's mean as the table's last row, as a meter's total line.
		entries[#entries + 1] = { cells = { { string.format(L.rotAllBosses, played.counted), "muted" },
			{ string.format(L.rotPercent, played.mean), "strong" }, nil } }
	end
	return entries
end

local function lostEntries(played)
	local entries = {}
	for _, kind in ipairs(played.lost) do
		entries[#entries + 1] = { cells = { { kind.label, "body" }, { kind.value, "warning" } } }
	end
	return entries
end

--- A muted or gold line of the right column at `y`; the y under it.
local function drawLine(region, parent, value, y)
	region:SetText(value)
	place(region, parent, 0, y)
	return y + textHeight(region)
end

local function hideRight(view)
	for _, region in ipairs({ view.summary, view.more, view.lost.heading, view.lost.note, view.meLink,
		view.meLink.label }) do
		region:Hide()
	end
	hideFrom(view.bossRows, 1)
	hideFrom(view.lostRows, 1)
end

--- How you played it, Lost across the night and the link to Me; the y under them.
local function drawPlayed(view, played, y)
	local parent, width = view.right, Theme.SIZES.windowRight
	if played.mean == nil then
		y = drawLine(view.summary, parent, L.rotNoFigures, y)
	end
	y = drawRows(view.bossRows, parent, y + Theme.SIZES.gap, bossEntries(played), bossLayout())
	if played.more > 0 then
		y = drawLine(view.more, parent, string.format(L.rotMore, played.more), y)
	end
	if #played.lost > 0 then
		y = BuildsWindow.drawSection(view.lost, parent, width, y + G.sectionGap, nil, fit, textHeight)
		y = drawRows(view.lostRows, parent, y, lostEntries(played), lostLayout())
	end
	if Hub.isAvailable(Hub.tab("me")) then
		y = drawLine(view.meLink.label, parent, L.rotOnMe, y + Theme.SIZES.gap)
		view.meLink:SetAllPoints(view.meLink.label)
		view.meLink:Show()
	end
	return y
end

--- The right column: the night, or the one line without one; its height.
local function drawRight(view, model)
	local played = model.played
	hideRight(view)
	local y = BuildsWindow.drawSection(view.played, view.right, Theme.SIZES.windowRight, 0,
		played ~= nil and played.note or nil, fit, textHeight)
	if played == nil then
		return drawLine(view.summary, view.right, model.companion and L.rotNoRaidLogged or L.meNoRaid, y)
	end
	return drawPlayed(view, played, y)
end

--- The gold link to Me under the right column (spec 3.5: gold text for going somewhere).
local function buildMeLink(parent)
	local link = Theme.createFrame("Button", nil, parent)
	link.label = text(parent, "gold")
	link:SetScript("OnClick", function()
		Hub.open("me")
	end)
	link:SetScript("OnEnter", function()
		link.label:SetTextColor(Theme.rgb(Theme.HEX.goldHover))
	end)
	link:SetScript("OnLeave", function()
		link.label:SetTextColor(Theme.rgb(Theme.HEX.gold))
	end)
	link:Hide()
	return link
end

--- The tab's body, once (Hub.register `build`): the two columns and the rule between them.
function RotationTab.build(body)
	local S = Theme.SIZES
	local divider = Theme.texture(body, "ARTWORK", "border")
	divider:SetWidth(Theme.wholePixels(divider, S.border))
	local ruleX = S.padding + S.windowLeft + S.windowGutter / 2
	divider:SetPoint("TOPLEFT", body, "TOPLEFT", ruleX, 0)
	divider:SetPoint("BOTTOMLEFT", body, "BOTTOMLEFT", ruleX, 0)
	local view = { body = body, lines = {}, weightRows = {}, bossRows = {}, lostRows = {} }
	view.left = column(body, S.padding, S.windowLeft)
	view.right = column(body, S.padding + S.windowLeft + S.windowGutter, S.windowRight)
	view.priority = BuildsWindow.buildSection(view.left, S.windowLeft, DEPS, L.rotPriority, nil)
	view.stepper = buildStepper(view, view.left)
	view.weights = BuildsWindow.buildSection(view.left, S.windowLeft, DEPS, L.rotWeights, nil)
	view.weights.haste = text(view.left, "muted", S.windowLeft)
	view.played = BuildsWindow.buildSection(view.right, S.windowRight, DEPS, L.rotPlayed, nil)
	view.lost = BuildsWindow.buildSection(view.right, S.windowRight, DEPS, L.rotLost, nil)
	view.summary = text(view.right, "body", S.windowRight)
	if type(view.summary.SetWordWrap) == "function" then
		view.summary:SetWordWrap(true)
	end
	view.more = text(view.right, "muted", S.windowRight)
	view.meLink = buildMeLink(view.right)
	view.measure = Theme.measurer(text(body, "muted"))
	RotationTab.view = view
	return body
end

--- The tab comes on screen (Hub.register `show`): the player's level, or `options.shown` (the stepper).
function RotationTab.show(options)
	local view = RotationTab.view
	if view == nil then
		return nil
	end
	local width = Theme.SIZES.windowLeft - G.number - G.icon - G.iconGap * 2
	local input = RotationTab.inputs(options ~= nil and options.shown or nil, view.measure, width)
	if input == nil then
		return nil
	end
	local model = RotationTab.model(input)
	view.model = model
	local left = drawLeft(view, model)
	local right = drawRight(view, model)
	Hub.setHints(nil, nil)
	Hub.fitHeight(math.max(left, right))
	return model
end

--- The tab leaves the screen: a hover it opened goes with it.
function RotationTab.hide()
	Theme.hideTooltip()
end

Hub.register({
	id = "rotation",
	label = L.hubTabRotation,
	build = RotationTab.build,
	show = RotationTab.show,
	hide = RotationTab.hide,
	figures = RotationTab.figures,
})

ns.RotationTab = RotationTab
return RotationTab
