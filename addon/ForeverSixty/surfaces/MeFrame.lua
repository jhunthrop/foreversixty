-- addon/ForeverSixty/surfaces/MeFrame.lua
-- The Me tab's frame (design/specs/2026-11-13-addon-hub.md 4.1, board hub-me.png): draws MeTab.model into
-- the hub's two columns, left 388 (the score, the last raid; with no site activity yet Where you stand and the
-- site line, 2026-11-14-me-empty.md), right 288 (Right now, the sim; or Gear at the band), and registers
-- the tab. Every region is built once on the first open and reused on every draw; sections are a heading
-- and rows, no boxes. Items show their icon in the quality edge with the client's tooltip and shift-click
-- links them; every Right now line and every gold link goes somewhere.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Tokens = ns.Tokens or require("Tokens")
local Widgets = ns.Widgets or require("Widgets")
local Compat = ns.Compat or require("Compat")
local Export = ns.Export or require("Export")
local Follow = ns.Follow or require("Follow")
local HubFrame = ns.HubFrame or require("HubFrame")
local Hub = ns.Hub or require("Hub")
local HubPlates = ns.HubPlates or require("HubPlates")
local BuildsWindow = ns.BuildsWindow or require("BuildsWindow")
local BuildDialog = ns.BuildDialog or require("BuildDialog")
local MeTab = ns.MeTab or require("MeTab")

local MeFrame = {}

--- Pixel layout. `line` a text line; `partRow` and `bossRow` a row of the score's parts and the raid's
--- table; `partLabel` where a part's track starts, `trackMax` the longest track (Experience, 700 points),
--- `track` its thickness; `big` the score's line and `bigSize` its type (24 px); `simSize` the DPS's type;
--- `sectionGap` between two sections; `nowRow` a Right now row (two lines), `nowIcon` its icon, `nowLink`
--- the room its tab link keeps; `gainIcon` and `gainRow` a gain's icon and row, `gainValue` its DPS column;
--- `underHead` the room under a section heading; `cell` an icon's edge.
MeFrame.LAYOUT = { line = 15, partRow = 18, partLabel = 86, trackMax = 130, track = 8, big = 30, bigSize = 24,
	simSize = 14, bossRow = 18, sectionGap = Theme.SIZES.plateSectionGap, chevron = 10, chevronGap = 2,
	nowRow = 38, nowIcon = 20, nowGap = 8, nowLink = 76,
	gainIcon = 16, gainGap = 6, gainRow = 20, gainValue = 72, valueWidth = 70, reasonGap = 6, summaryGap = 6,
	underHead = 4, cell = 1 }
--- The last raid's columns (x, width, right-aligned) summing to the left column, 388 (spec 3.5).
MeFrame.COLUMNS = {
	{ key = "name", head = "meBoss", x = 0, width = 112 },
	{ key = "parse", head = "meParse", x = 112, width = 33, right = true },
	{ key = "loss", head = "meLoss", x = 151, width = 128 },
	{ key = "died", head = "meDied", x = 279, width = 28, right = true },
	{ key = "consumables", head = "meConsumables", x = 312, width = 76 },
}

local function textHeight(region)
	local height = type(region.GetStringHeight) == "function" and region:GetStringHeight() or nil
	return (type(height) == "number" and height > 0) and height or MeFrame.LAYOUT.line
end

local function fit(region, raw, width)
	region:SetWidth(width)
	region:SetText(raw or "")
end

local function place(region, parent, x, y)
	region:ClearAllPoints()
	region:SetPoint("TOPLEFT", parent, "TOPLEFT", x, -y)
	region:Show()
end

local function text(parent, hexKey, width)
	return HubFrame.line(parent, hexKey, width)
end

local function hideAll(regions)
	for _, region in ipairs(regions) do
		region:Hide()
	end
end

--- A text link: the label in the link colour, white under the cursor, a click runs `onClick`. With
--- `chevron`, the site's chevron after it says the row opens somewhere (spec 5: Right now).
local function link(parent, label, onClick, chevron)
	local G = MeFrame.LAYOUT
	local button = Theme.createFrame("Button", nil, parent)
	local region = text(button, "link")
	region:SetPoint("LEFT", button, "LEFT", 0, 0)
	button.foreverSixtyText = region
	local glyph = chevron and Theme.glyph(button, "ARTWORK", "chevron-right", G.chevron, "link") or nil
	if glyph ~= nil then
		glyph:SetPoint("LEFT", region, "RIGHT", G.chevronGap, 0)
	end
	button:SetHeight(G.line)
	button.SetLabel = function(self, value)
		region:SetText(value or "")
		local width = type(region.GetStringWidth) == "function" and region:GetStringWidth() or 0
		local extra = glyph ~= nil and G.chevronGap + G.chevron or 0
		self:SetWidth(math.max(1, math.ceil((width or 0) + extra)))
	end
	local function paint(key)
		region:SetTextColor(Theme.rgb(Theme.HEX[key]))
		if glyph ~= nil then
			Theme.tint(glyph, key)
		end
	end
	button:SetScript("OnEnter", function()
		paint("linkHover")
	end)
	button:SetScript("OnLeave", function()
		paint("link")
	end)
	button:SetScript("OnClick", function(self, ...)
		onClick(self, ...)
	end)
	button:SetLabel(label)
	return button
end

--- An item's icon in its quality edge: the client's item tooltip on hover, shift-click links it.
local function itemCell(parent, size)
	local cell = Theme.createFrame("Button", nil, parent)
	cell:SetSize(size, size)
	cell.icon = Theme.icon(cell, "ARTWORK", nil, size)
	cell.icon:SetAllPoints(cell)
	cell.edges = Theme.outline(cell, MeFrame.LAYOUT.cell, "border")
	Widgets.attachTooltip(cell, function()
		return cell.itemId, nil
	end)
	cell:SetScript("OnClick", function(self)
		local linkText = self.itemId ~= nil and select(2, Compat.itemInfo(self.itemId)) or nil
		local modified, insert = _G.IsModifiedClick, _G.ChatEdit_InsertLink
		if linkText ~= nil and type(modified) == "function" and modified("CHATLINK") and type(insert) == "function" then
			insert(linkText)
		end
	end)
	return cell
end

--- Point `cell` at `itemId` (or a plain icon `path`): its icon and its edge in the item's quality colour.
local function paintCell(cell, itemId, path)
	cell.itemId = itemId
	cell.icon:SetTexture(itemId ~= nil and Compat.itemIcon(itemId) or path or Theme.UNKNOWN_ICON)
	local quality = itemId ~= nil and Compat.itemQuality(itemId) or nil
	local hex = quality ~= nil and Theme.QUALITY_HEX[quality] or Theme.HEX.border
	for _, edge in ipairs(cell.edges) do
		Theme.paintRGB(edge, Theme.rgb(hex))
	end
end

-- Building --------------------------------------------------------------------------------------------

local function buildScore(parent, width)
	local G = MeFrame.LAYOUT
	local score = BuildsWindow.buildSection(parent, width, { text = text }, L.meScore, nil)
	score.line = text(parent, "muted", width)
	score.number = text(parent, "body")
	Theme.setFace(score.number, Tokens.type.addon.strong, G.bigSize)
	score.tier = text(parent, "body")
	score.toNext = text(parent, "muted")
	score.parts = {}
	for index = 1, 5 do
		local part = {}
		part.label = text(parent, "muted", G.partLabel)
		part.track = Theme.texture(parent, "ARTWORK", "track")
		part.fill = Theme.texture(parent, "OVERLAY", "body")
		part.reason = text(parent, "muted")
		part.value = text(parent, "body", G.valueWidth)
		part.value:SetJustifyH("RIGHT")
		score.parts[index] = part
	end
	return score
end

local function buildRaid(parent, width)
	local raid = BuildsWindow.buildSection(parent, width, { text = text }, L.meLastRaid, nil)
	raid.heads, raid.rows = {}, {}
	for index, column in ipairs(MeFrame.COLUMNS) do
		local head = text(parent, "muted", column.width)
		head:SetText(L[column.head])
		if column.right then
			head:SetJustifyH("RIGHT")
		end
		raid.heads[index] = head
	end
	for index = 1, MeTab.MAX_BOSSES do
		-- Every other boss row on the zebra tint (spec 2.5).
		local row = { zebra = index % 2 == 1 and Theme.texture(parent, "BACKGROUND", "zebra") or nil }
		for _, column in ipairs(MeFrame.COLUMNS) do
			row[column.key] = text(parent, column.key == "name" and "body" or "muted", column.width)
			if column.right then
				row[column.key]:SetJustifyH("RIGHT")
			end
		end
		raid.rows[index] = row
	end
	raid.summary = text(parent, "body")
	raid.empty = text(parent, "muted", width)
	raid.fullLog = link(parent, L.meFullLog, function()
		MeFrame.openLog()
	end)
	raid.app = link(parent, L.meAppLink, function()
		MeFrame.openApp()
	end)
	return raid
end

local function buildNow(parent, width)
	local G = MeFrame.LAYOUT
	local now = BuildsWindow.buildSection(parent, width, { text = text }, L.meRightNow, nil)
	now.rows = {}
	for index = 1, MeTab.MAX_NOW do
		local row = {}
		row.cell = itemCell(parent, G.nowIcon)
		local room = width - G.nowIcon - G.nowGap - G.nowLink
		row.line = text(parent, "body", room)
		row.sub = text(parent, "muted", room)
		row.link = link(parent, "", function()
			MeFrame.stopPulse()
			if row.tab ~= nil then
				Hub.select(row.tab)
			end
		end, true)
		now.rows[index] = row
	end
	-- The next action's accent bar on the first row's plate edge (spec 9: it pulses, then holds).
	now.bar = Theme.texture(parent, "ARTWORK", "accent")
	now.bar:SetSize(Theme.SIZES.tabUnderline, G.nowRow - G.nowGap)
	now.bar:Hide()
	return now
end

local function buildSim(parent, width)
	local G = MeFrame.LAYOUT
	local sim = BuildsWindow.buildSection(parent, width, { text = text }, L.meSim, nil)
	sim.dps = text(parent, "strong")
	Theme.setFace(sim.dps, Tokens.type.addon.strong, G.simSize)
	sim.error = text(parent, "muted")
	sim.label = text(parent, "muted", width)
	sim.label:SetText(L.meBiggestGains)
	sim.gains = {}
	for index = 1, MeTab.MAX_GAINS do
		local value = text(parent, "success", G.gainValue)
		value:SetJustifyH("RIGHT")
		sim.gains[index] = { cell = itemCell(parent, G.gainIcon),
			name = text(parent, "body", width - G.gainIcon - G.gainGap * 2 - G.gainValue), text = value }
	end
	sim.date = text(parent, "muted")
	sim.button = Widgets.clientButton(parent, L.meSimAgain, function()
		MeFrame.simAgain()
	end)
	return sim
end

--- Where you stand (no site activity yet): a row each of label, track and fill, icon, value and note.
local function buildStanding(parent, width)
	local G = MeFrame.LAYOUT
	local standing = BuildsWindow.buildSection(parent, width, { text = text }, L.meStanding, nil)
	standing.rows = {}
	for index = 1, MeTab.MAX_STANDING do
		local row = { label = text(parent, "muted", G.partLabel), value = text(parent, "body"),
			note = text(parent, "muted", G.valueWidth), track = Theme.texture(parent, "ARTWORK", "track"),
			fill = Theme.texture(parent, "OVERLAY", "body"), cell = itemCell(parent, G.gainIcon) }
		row.note:SetJustifyH("RIGHT")
		standing.rows[index] = row
	end
	return standing
end

--- Gear at the band (no site activity yet): each upgrade as a Right now row (icon, name and where, the gain
--- at the right), then the Gear tab's own enchant rows (GearEnchants), or the empty line.
local function buildGear(parent, width)
	local G = MeFrame.LAYOUT
	local gear = BuildsWindow.buildSection(parent, width, { text = text }, L.meGear, nil)
	gear.rows = {}
	for index = 1, MeTab.MAX_LOCAL_GAINS do
		local value = text(parent, "success", G.gainValue)
		value:SetJustifyH("RIGHT")
		local room = width - G.nowIcon - G.nowGap - G.gainValue
		gear.rows[index] = { cell = itemCell(parent, G.nowIcon), name = text(parent, "body", room),
			where = text(parent, "muted", room + G.gainValue), text = value }
	end
	gear.enchants = (ns.GearEnchants or require("GearEnchants")).build(parent, width, MeTab.MAX_LOCAL_ENCHANTS)
	gear.empty = text(parent, "muted", width)
	return gear
end

--- The one line for what the site adds, and the companion app's address under it.
local function buildSite(parent, width)
	local line = text(parent, "muted", width)
	if line.SetWordWrap then line:SetWordWrap(true) end
	return { line = line, link = link(parent, L.meAppLink, function()
		MeFrame.openApp()
	end) }
end

--- The tab's body, once, into the hub's content frame: the two columns on their plates.
function MeFrame.build(content)
	local S = Theme.SIZES
	local left, right = Hub.columns(content)
	MeFrame.parts = {
		content = content, left = left, right = right,
		score = buildScore(left, S.windowLeft), raid = buildRaid(left, S.windowLeft),
		standing = buildStanding(left, S.windowLeft), gear = buildGear(right, S.windowRight),
		site = buildSite(left, S.windowLeft),
		now = buildNow(right, S.windowRight), sim = buildSim(right, S.windowRight),
	}
	-- Right now is the tab's lead plate (spec 2.5).
	MeFrame.parts.now.plateLook = HubPlates.LEAD
	return content
end

-- Drawing ------------------------------------------------------------------------------------------------

--- A section's heading (the Builds tab's, BuildsWindow.drawSection) and its note at `y`; the y under it.
local function heading(section, parent, width, y, note)
	return BuildsWindow.drawSection(section, parent, width, y, note, fit, textHeight) + MeFrame.LAYOUT.underHead
end

local function hideSection(section)
	section.heading:Hide()
	section.note:Hide()
end

local function drawParts(score, parent, model, y)
	local G = MeFrame.LAYOUT
	for index, part in ipairs(score.parts) do
		local row = model.parts[index]
		if row == nil then
			hideAll({ part.label, part.track, part.fill, part.reason, part.value })
		else
			local top = y + (index - 1) * G.partRow
			local mid = top + (G.partRow - G.line) / 2
			part.label:SetText(row.label)
			place(part.label, parent, 0, mid)
			local length = math.floor(G.trackMax * row.track + 0.5)
			part.track:SetSize(length, G.track)
			place(part.track, parent, G.partLabel, top + (G.partRow - G.track) / 2)
			local filled = math.floor(length * row.fill + 0.5)
			Theme.paint(part.fill, "accent")
			Theme.setShown({ part.fill }, filled > 0)
			if filled > 0 then
				part.fill:SetSize(filled, G.track)
				place(part.fill, parent, G.partLabel, top + (G.partRow - G.track) / 2)
			end
			if row.reason ~= nil then
				part.reason:SetText(row.reason)
				place(part.reason, parent, G.partLabel + length + G.reasonGap, mid)
			else
				part.reason:Hide()
			end
			part.value:SetText(row.points .. " " .. Theme.inlineColor("muted", string.format(L.meOfMax, row.max)))
			place(part.value, parent, Theme.SIZES.windowLeft - G.valueWidth, mid)
		end
	end
	return y + #model.parts * G.partRow
end

local function scoreNote(model)
	if model.note == nil then
		return nil
	end
	local parts = {}
	for index, part in ipairs(model.note) do
		parts[index] = (index == 1 and model.noteWarn) and Theme.inlineColor("warning", part) or part
	end
	return table.concat(parts, L.dialogSeparator)
end

--- The Forever Score at `y`; the y under it (y itself when the section is left out).
local function drawScore(score, parent, model, y)
	local G, width = MeFrame.LAYOUT, Theme.SIZES.windowLeft
	hideAll({ score.line, score.number, score.tier, score.toNext })
	if model == nil then
		hideSection(score)
		drawParts(score, parent, { parts = {} }, y)
		return y
	end
	if model.line ~= nil then
		y = heading(score, parent, width, y, nil)
		drawParts(score, parent, { parts = {} }, y)
		score.line:SetText(model.line)
		place(score.line, parent, 0, y + 2)
		return y + 2 + textHeight(score.line) + G.sectionGap
	end
	y = heading(score, parent, width, y, scoreNote(model))
	score.number:SetText(tostring(model.total))
	score.number:SetTextColor(Theme.rgb(model.tierHex or Theme.HEX.body))
	place(score.number, parent, 0, y + (G.big - textHeight(score.number)) / 2)
	local x = math.ceil(score.number:GetStringWidth() or 0) + G.reasonGap
	score.tier:SetText(model.tierLabel or "")
	score.tier:SetTextColor(Theme.rgb(model.tierHex or Theme.HEX.body))
	place(score.tier, parent, x, y + G.big - G.line - 2)
	x = x + math.ceil(score.tier:GetStringWidth() or 0) + G.reasonGap
	score.toNext:SetText(L.dialogSeparator:gsub("^%s+", "") .. model.toNext)
	place(score.toNext, parent, x, y + G.big - G.line - 2)
	y = drawParts(score, parent, model, y + G.big)
	return y + G.sectionGap
end

local function drawBossRows(raid, parent, rows, y)
	local G = MeFrame.LAYOUT
	for index, row in ipairs(raid.rows) do
		local entry = rows[index]
		if entry == nil then
			hideAll({ row.name, row.parse, row.loss, row.died, row.consumables, row.zebra })
		else
			if row.zebra ~= nil then
				row.zebra:ClearAllPoints()
				row.zebra:SetPoint("TOPLEFT", parent, "TOPLEFT", 0, -(y + (index - 1) * G.bossRow))
				row.zebra:SetSize(Theme.SIZES.windowLeft, G.bossRow)
				row.zebra:Show()
			end
			local top = y + (index - 1) * G.bossRow + (G.bossRow - G.line) / 2
			for _, column in ipairs(MeFrame.COLUMNS) do
				place(row[column.key], parent, column.x, top)
			end
			fit(row.name, entry.name, MeFrame.COLUMNS[1].width)
			row.parse:SetText(tostring(entry.parse))
			row.parse:SetTextColor(Theme.rgb(entry.parseHex))
			fit(row.loss, entry.loss or "-", MeFrame.COLUMNS[3].width)
			row.died:SetText(tostring(entry.died))
			row.died:SetTextColor(Theme.rgb(Theme.HEX[entry.died > 0 and "warning" or "muted"]))
			fit(row.consumables, entry.consumables, MeFrame.COLUMNS[5].width)
			row.consumables:SetTextColor(Theme.rgb(Theme.HEX[entry.missed and "warning" or "muted"]))
		end
	end
	return y + #rows * G.bossRow
end

--- The last raid at `y`; the y under it.
local function drawRaid(raid, parent, model, y)
	local G, width = MeFrame.LAYOUT, Theme.SIZES.windowLeft
	hideAll({ raid.summary, raid.empty, raid.fullLog, raid.app })
	hideAll(raid.heads)
	if model == nil then
		hideSection(raid)
		drawBossRows(raid, parent, {}, y)
		return y
	end
	y = heading(raid, parent, width, y, model.note)
	if model.empty ~= nil then
		drawBossRows(raid, parent, {}, y)
		raid.empty:SetText(model.empty)
		place(raid.empty, parent, 0, y + 2)
		y = y + 2 + textHeight(raid.empty)
		raid.app:SetLabel(model.link)
		place(raid.app, parent, 0, y + 4)
		return y + 4 + G.line + G.sectionGap
	end
	for index, column in ipairs(MeFrame.COLUMNS) do
		place(raid.heads[index], parent, column.x, y + (G.bossRow - G.line) / 2)
	end
	y = drawBossRows(raid, parent, model.rows, y + G.bossRow)
	local summary = model.summary
	raid.summary:SetText(L.meMedian .. " " .. "|cff" .. summary.medianHex .. summary.median .. "|r"
		.. Theme.inlineColor("muted", L.dialogSeparator .. summary.rest .. (model.url and L.dialogSeparator or "")))
	place(raid.summary, parent, 0, y + G.summaryGap)
	if model.url ~= nil then
		MeFrame.logUrl, MeFrame.logDay = model.url, model.day
		local x = math.ceil(raid.summary:GetStringWidth() or 0)
		place(raid.fullLog, parent, x, y + G.summaryGap)
	end
	return y + G.summaryGap + G.line + G.sectionGap
end

--- A section's heading at `y` while it has a model (the y under it), else the heading hidden and `y`.
local function headingIf(section, parent, width, y, model, note)
	if model == nil then
		hideSection(section)
		return y
	end
	return heading(section, parent, width, y, note)
end

local function drawNow(now, parent, rows, y)
	local G, width = MeFrame.LAYOUT, Theme.SIZES.windowRight
	y = headingIf(now, parent, width, y, #rows > 0 and rows or nil, nil)
	for index, row in ipairs(now.rows) do
		local entry = rows[index]
		row.tab = entry ~= nil and entry.tab or nil
		if entry == nil then
			hideAll({ row.cell, row.line, row.sub, row.link })
		else
			local top = y + (index - 1) * G.nowRow
			if entry.itemId ~= nil or entry.icon ~= nil then
				paintCell(row.cell, entry.itemId, entry.icon)
				place(row.cell, parent, 0, top + 2)
			else
				row.cell:Hide()
			end
			-- A row with no icon (a tab's waiting line) starts at the column's edge.
			local x = (entry.itemId ~= nil or entry.icon ~= nil) and G.nowIcon + G.nowGap or 0
			local room = width - x - G.nowLink
			fit(row.line, entry.urgent and entry.sub == nil and Theme.inlineColor("clientGold", entry.line) or entry.line,
				room)
			place(row.line, parent, x, top)
			if entry.sub ~= nil then
				fit(row.sub, entry.sub, room)
				place(row.sub, parent, x, top + G.line)
			else
				row.sub:Hide()
			end
			row.link:SetLabel(entry.label)
			row.link:ClearAllPoints()
			row.link:SetPoint("TOPRIGHT", parent, "TOPLEFT", width, -top)
			row.link:Show()
		end
	end
	MeFrame.drawBar(now, parent, rows[1], y)
	if #rows == 0 then
		return y
	end
	return y + #rows * G.nowRow + G.sectionGap - (G.nowRow - 2 * G.line)
end

--- The first Right now row's accent bar at `y` (hidden with no row): it pulses on the window's first
--- show in a session (Theme.pulseNext), with the dot of the tab the row leads to, then holds.
function MeFrame.drawBar(now, parent, first, y)
	if first == nil then
		now.bar:Hide()
		return nil
	end
	now.bar:ClearAllPoints()
	now.bar:SetPoint("TOPLEFT", parent, "TOPLEFT", -Theme.SIZES.platePad, -y)
	now.bar:Show()
	if not MeFrame.pulsed then
		MeFrame.pulsed = true
		local strip = Hub.strip
		local button = strip ~= nil and first.tab ~= nil and strip.buttons[first.tab] or nil
		MeFrame.pulsing = { now.bar, button ~= nil and button.foreverSixtyDot or nil }
		Theme.pulseNext(MeFrame.pulsing)
	end
	return now.bar
end

--- The pulse ends at once (the row clicked, the tab left): the bar and the dot at the steady alpha.
function MeFrame.stopPulse()
	Theme.stopNext(MeFrame.pulsing or {})
	MeFrame.pulsing = nil
end

--- The item's name in its quality colour: the client's, else the shipped list's while the client loads it.
local function itemName(itemId)
	local data = MeFrame.data or ns.Data
	local names = type(data) == "table" and type(data.bisNames) == "table" and data.bisNames or {}
	local name = Compat.itemName(itemId) or names[itemId] or ""
	local quality = Compat.itemQuality(itemId)
	return name, quality ~= nil and Theme.QUALITY_HEX[quality] or Theme.HEX.body
end

--- "Biggest gains", then a row each: the item's icon and name, its DPS at the right. The y under them.
local function drawGains(sim, parent, gains, y)
	local G, width = MeFrame.LAYOUT, Theme.SIZES.windowRight
	Theme.setShown({ sim.label }, #gains > 0)
	local top = y + G.line + 2
	for index, gain in ipairs(sim.gains) do
		local entry = gains[index]
		if entry == nil then
			hideAll({ gain.cell, gain.name, gain.text })
		else
			local rowTop = top + (index - 1) * G.gainRow
			paintCell(gain.cell, entry.itemId)
			place(gain.cell, parent, 0, rowTop + (G.gainRow - G.gainIcon) / 2)
			local name, hex = itemName(entry.itemId)
			fit(gain.name, name, width - G.gainIcon - G.gainGap * 2 - G.gainValue)
			gain.name:SetTextColor(Theme.rgb(hex))
			place(gain.name, parent, G.gainIcon + G.gainGap, rowTop + (G.gainRow - G.line) / 2)
			gain.text:SetText(string.format(L.meGain, entry.gain))
			place(gain.text, parent, width - G.gainValue, rowTop + (G.gainRow - G.line) / 2)
		end
	end
	if #gains == 0 then
		return y
	end
	place(sim.label, parent, 0, y)
	return top + #gains * G.gainRow
end

--- The sim at `y`; the y under it.
local function drawSim(sim, parent, model, y)
	local G, width = MeFrame.LAYOUT, Theme.SIZES.windowRight
	hideAll({ sim.dps, sim.error, sim.date })
	-- The button stays shown across redraws (it may be under the cursor); it hides only when it goes.
	Theme.setShown({ sim.button }, model ~= nil and model.button ~= nil)
	if model == nil then
		hideSection(sim)
		drawGains(sim, parent, {}, y)
		return y
	end
	y = heading(sim, parent, width, y, nil)
	if model.dps ~= nil then
		sim.dps:SetText(model.dps)
		place(sim.dps, parent, 0, y + 2)
		if model.error ~= nil then
			sim.error:SetText(model.error)
			sim.error:ClearAllPoints()
			sim.error:SetPoint("BOTTOMLEFT", sim.dps, "BOTTOMRIGHT", G.reasonGap, 0)
			sim.error:Show()
		end
		y = y + 2 + textHeight(sim.dps) + 4
	end
	y = drawGains(sim, parent, model.gains, y)
	local foot = y + 4
	if model.date ~= nil then
		sim.date:SetText(model.date)
		place(sim.date, parent, 0, foot + (Theme.SIZES.clientButtonHeight - G.line) / 2)
	end
	if model.button ~= nil then
		Widgets.resetButtonLabel(sim.button, model.button)
		Widgets.fitButton(sim.button)
		Widgets.setEnabled(sim.button, not model.queued)
		sim.button:ClearAllPoints()
		sim.button:SetPoint("TOPRIGHT", parent, "TOPLEFT", width, -foot)
		sim.button:Show()
	end
	if model.date ~= nil or model.button ~= nil then
		y = foot + Theme.SIZES.clientButtonHeight
	end
	return y + G.sectionGap
end

-- No site activity yet -------------------------------------------------------------------------------------

--- One Where you stand row at `top`: the label, then the experience track or the icon and value, the note.
local function drawStandingRow(row, parent, entry, top)
	local G, width = MeFrame.LAYOUT, Theme.SIZES.windowLeft
	local mid = top + (G.partRow - G.line) / 2
	place(row.label, parent, 0, mid)
	row.label:SetText(entry.label)
	local x = G.partLabel
	Theme.setShown({ row.track, row.fill }, entry.fill ~= nil)
	if entry.fill ~= nil then
		row.track:SetSize(G.trackMax, G.track)
		place(row.track, parent, x, top + (G.partRow - G.track) / 2)
		local filled = math.floor(G.trackMax * entry.fill + 0.5)
		Theme.setShown({ row.fill }, filled > 0)
		row.fill:SetSize(math.max(filled, 1), G.track)
		row.fill:ClearAllPoints()
		row.fill:SetPoint("TOPLEFT", parent, "TOPLEFT", x, -(top + (G.partRow - G.track) / 2))
		x = x + G.trackMax + G.reasonGap
	end
	Theme.setShown({ row.cell }, entry.icon ~= nil)
	if entry.icon ~= nil then
		paintCell(row.cell, nil, entry.icon)
		place(row.cell, parent, x, top + (G.partRow - G.gainIcon) / 2)
		x = x + G.gainIcon + G.gainGap
	end
	fit(row.value, entry.value, width - x - G.valueWidth - G.reasonGap)
	row.value:SetTextColor(Theme.rgb(Theme.HEX[entry.muted and "muted" or "body"]))
	place(row.value, parent, x, mid)
	Theme.setShown({ row.note }, entry.note ~= nil)
	if entry.note ~= nil then
		row.note:SetText(entry.note)
		place(row.note, parent, width - G.valueWidth, mid)
	end
end

--- Where you stand at `y`; the y under it.
local function drawStanding(standing, parent, model, y)
	local G, rows = MeFrame.LAYOUT, model ~= nil and model.rows or {}
	y = headingIf(standing, parent, Theme.SIZES.windowLeft, y, model, nil)
	for index, row in ipairs(standing.rows) do
		if rows[index] == nil then
			hideAll({ row.label, row.value, row.note, row.track, row.fill, row.cell })
		else
			drawStandingRow(row, parent, rows[index], y + (index - 1) * G.partRow)
		end
	end
	return model == nil and y or y + #rows * G.partRow + G.sectionGap
end

--- Gear at `y` in `parent`, `width` wide: the upgrades, then the enchant rows or the empty line; the y under it.
local function drawLocalGear(gear, parent, model, y, width)
	local G, rows = MeFrame.LAYOUT, model ~= nil and model.rows or {}
	y = headingIf(gear, parent, width, y, model, model ~= nil and model.note or nil)
	for index, row in ipairs(gear.rows) do
		local entry, top = rows[index], y + (index - 1) * G.nowRow
		if entry == nil then
			hideAll({ row.cell, row.name, row.where, row.text })
		else
			paintCell(row.cell, entry.itemId)
			place(row.cell, parent, 0, top + 2)
			local name, hex = itemName(entry.itemId)
			fit(row.name, name, width - G.nowIcon - G.nowGap - G.gainValue)
			row.name:SetTextColor(Theme.rgb(hex))
			place(row.name, parent, G.nowIcon + G.nowGap, top)
			-- The drop chance goes first when the source does not fit; the boss stays whole.
			fit(row.where, entry.where or "", width - G.nowIcon - G.nowGap)
			if (row.where:GetStringWidth() or 0) > width - G.nowIcon - G.nowGap then
				row.where:SetText(entry.whereShort or entry.where)
			end
			place(row.where, parent, G.nowIcon + G.nowGap, top + G.line)
			row.text:SetText(entry.gain)
			place(row.text, parent, width - G.gainValue, top)
		end
	end
	y = y + #rows * G.nowRow - (#rows > 0 and G.nowRow - 2 * G.line or 0)
	local enchants = model ~= nil and model.enchants or {}
	local GearEnchants = ns.GearEnchants or require("GearEnchants")
	y = GearEnchants.draw(gear.enchants, parent, enchants, y + (#enchants > 0 and 4 or 0), GearEnchants.shortWords)
	Theme.setShown({ gear.empty }, model ~= nil and model.empty ~= nil)
	if model ~= nil and model.empty ~= nil then
		gear.empty:SetText(model.empty)
		place(gear.empty, parent, 0, y + 2)
		y = y + 2 + G.line
	end
	return model == nil and y or y + G.sectionGap
end

--- The site line at `y` and the app's address under it; the y under them.
local function drawSite(site, parent, model, y)
	Theme.setShown({ site.line, site.link }, model ~= nil)
	if model == nil then
		return y
	end
	site.line:SetText(model.line)
	place(site.line, parent, 0, y)
	y = y + textHeight(site.line)
	place(site.link, parent, 0, y + 4)
	return y + 4 + MeFrame.LAYOUT.line + MeFrame.LAYOUT.sectionGap
end

--- Draw `model` (MeTab.model): the columns, the footer's line, the window's height. The body's height.
function MeFrame.draw(model)
	local parts = MeFrame.parts
	if parts == nil then
		return 0
	end
	local G = MeFrame.LAYOUT
	local left = drawScore(parts.score, parts.left, model.score, 0)
	left = drawRaid(parts.raid, parts.left, model.lastRaid, left)
	left = drawStanding(parts.standing, parts.left, model.standing, left)
	local right = drawNow(parts.now, parts.right, model.rightNow, 0)
	if model.setup ~= nil and MeFrame.setup ~= nil then
		right = MeFrame.setup.draw(parts.right, Theme.SIZES.windowRight, model.setup, right)
	elseif MeFrame.setup ~= nil then
		MeFrame.setup.hide()
	end
	-- Gear (no site activity yet) sits on the right above the sim, unless the left column is shorter by more
	-- than a row with it: measured where it is drawn, then moved (never hidden) when it goes left.
	local geared = drawLocalGear(parts.gear, parts.right, model.gear, right, Theme.SIZES.windowRight)
	local sim = drawSim(parts.sim, parts.right, model.sim, right) - right
	if math.max(left + geared - right, right + sim) + G.nowRow < math.max(left, geared + sim) then
		left, right = drawLocalGear(parts.gear, parts.left, model.gear, left, Theme.SIZES.windowLeft), right + sim
	else
		right = drawSim(parts.sim, parts.right, model.sim, geared)
	end
	left = drawSite(parts.site, parts.left, model.site, left)
	-- Every section ends in its gap to the next; the last one's is not body.
	local body = math.max(0, math.max(left, right) - G.sectionGap)
	Hub.setHints(nil, model.footer)
	Hub.fitHeight(body)
	return body
end

-- The tab -------------------------------------------------------------------------------------------------

-- The pieces Set up (Onboarding, the first open) draws with, so its section reads as the tab's own.
MeFrame.kit = {
	text = text, link = link, itemCell = itemCell, paintCell = paintCell, heading = heading, place = place,
	fit = fit, hideAll = hideAll, hideSection = hideSection,
	section = function(parent, width, title)
		return BuildsWindow.buildSection(parent, width, { text = text }, title, nil)
	end,
}

--- Injection: the first open's Set up (Onboarding sets it): `model(data)` while it shows, else nil;
--- `draw(parent, width, model, y)` and `hide()` in the right column, in Right now's place.
MeFrame.setup = nil

--- The live inputs.
function MeFrame.inputs()
	return MeTab.inputs({ Hub = Hub, Export = Export, Follow = Follow, BuildDialog = BuildDialog,
		data = MeFrame.data or ns.Data })
end

function MeFrame.show()
	local input = MeFrame.inputs()
	local setup = MeFrame.setup
	input.setup = setup ~= nil and setup.model(MeFrame.data or ns.Data) or nil
	MeFrame.model = MeTab.model(input)
	return MeFrame.draw(MeFrame.model)
end

--- Draw again while the tab shows (a sim queued, the inbox read).
function MeFrame.refresh()
	if Hub.isShown() and Hub.active == "me" then
		MeFrame.show()
		Hub.refresh()
	end
end

--- Full log: the report's address in the Builds tab's copy box (as /fs link does).
function MeFrame.openLog()
	if MeFrame.logUrl == nil then
		return false
	end
	local day = MeFrame.logDay or ""
	return BuildDialog.openLink(MeFrame.logUrl, Compat.playerName() or "",
		{ lead = string.format(L.meLogLead, day), copied = string.format(L.meLogCopied, day) })
end

--- The companion app's address in the copy box.
function MeFrame.openApp()
	return BuildDialog.openLink(L.meAppLink, Compat.playerName() or "", { lead = L.meAppLead, copied = L.meAppCopied })
end

--- Sim again (contract 4): the sim request set as Export.outbox's `simRequest` part and written into the
--- outbox at once, for the spec the followed build is for. False when one already waits this session, or
--- there is no export or spec to send.
function MeFrame.simAgain()
	local data = MeFrame.data or ns.Data
	local spec = Follow.buildSpec(data, Follow.build)
	if Export.simRequest ~= nil or spec == nil then
		return false
	end
	local code = Export.string(data)
	if code == nil then
		return false
	end
	Export.simRequest = {
		character = Export.outboxCharacter(),
		spec = spec,
		requested_at = date(Export.SAVED_AT_FORMAT),
		gear_code = code,
	}
	Export.saveOutbox(data, code)
	MeFrame.refresh()
	return true
end

function MeFrame.figures()
	local guild = Theme.playerGuild()
	return MeTab.figures({ guild = guild, buildName = type(Follow.build) == "table" and Follow.build.name or nil })
end

function MeFrame.waiting()
	local key = Export.characterKey()
	local found = (ns.Codec or require("Codec")).inboxMessages(_G.ForeverSixtyInbox, key, "last_raid")
	return MeTab.waitingItems(found[#found], (ns.Score or require("Score")).clock(), Hub.isSeen)
end

Hub.register({
	id = "me",
	label = L.hubTabMe,
	build = MeFrame.build,
	show = MeFrame.show,
	hide = MeFrame.stopPulse,
	figures = MeFrame.figures,
	waiting = MeFrame.waiting,
})

ns.MeFrame = MeFrame
return MeFrame
