-- addon/ForeverSixty/surfaces/MeFrame.lua
-- The Me tab's frame (design/specs/2026-11-13-addon-hub.md 4.1, board hub-me.png): draws MeTab.model into
-- the hub's two columns, left 388 (the score, the last raid), right 288 (Right now, the sim), and registers
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
	simSize = 14, bossRow = 18, sectionGap = 16, nowRow = 38, nowIcon = 20, nowGap = 8, nowLink = 64,
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

--- A gold text link: the label in gold, lifted under the cursor, a click runs `onClick`.
local function link(parent, label, onClick)
	local button = Theme.createFrame("Button", nil, parent)
	local region = text(button, "gold")
	region:SetPoint("LEFT", button, "LEFT", 0, 0)
	button.foreverSixtyText = region
	button:SetHeight(MeFrame.LAYOUT.line)
	button.SetLabel = function(self, value)
		region:SetText(value or "")
		local width = type(region.GetStringWidth) == "function" and region:GetStringWidth() or 0
		self:SetWidth(math.max(1, math.ceil(width or 0)))
	end
	button:SetScript("OnEnter", function()
		region:SetTextColor(Theme.rgb(Theme.HEX.goldHover))
	end)
	button:SetScript("OnLeave", function()
		region:SetTextColor(Theme.rgb(Theme.HEX.gold))
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
		local row = {}
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
			if row.tab ~= nil then
				Hub.select(row.tab)
			end
		end)
		now.rows[index] = row
	end
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

--- The tab's body, once, into the hub's content frame: the two columns and the rule between them.
function MeFrame.build(content)
	local S, H = Theme.SIZES, Hub.LAYOUT
	local divider = Theme.texture(content, "ARTWORK", "border")
	divider:SetWidth(Theme.wholePixels(divider, S.border))
	local ruleX = S.padding + S.windowLeft + S.windowGutter / 2
	divider:SetPoint("TOPLEFT", content, "TOPLEFT", ruleX, 0)
	divider:SetPoint("BOTTOMLEFT", content, "BOTTOMLEFT", ruleX, 0)
	local left = Theme.createFrame("Frame", nil, content)
	left:SetPoint("TOPLEFT", content, "TOPLEFT", S.padding, -H.bodyPad)
	left:SetSize(S.windowLeft, 1)
	local right = Theme.createFrame("Frame", nil, content)
	right:SetPoint("TOPLEFT", content, "TOPLEFT", S.padding + S.windowLeft + S.windowGutter, -H.bodyPad)
	right:SetSize(S.windowRight, 1)
	MeFrame.parts = {
		content = content, left = left, right = right,
		score = buildScore(left, S.windowLeft), raid = buildRaid(left, S.windowLeft),
		now = buildNow(right, S.windowRight), sim = buildSim(right, S.windowRight),
	}
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
			Theme.paintRGB(part.fill, Theme.rgb(row.hex or Theme.HEX.body))
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
			hideAll({ row.name, row.parse, row.loss, row.died, row.consumables })
		else
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

local function drawNow(now, parent, rows, y)
	local G, width = MeFrame.LAYOUT, Theme.SIZES.windowRight
	if #rows == 0 then
		hideSection(now)
	else
		y = heading(now, parent, width, y, nil)
	end
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
	if #rows == 0 then
		return y
	end
	return y + #rows * G.nowRow + G.sectionGap - (G.nowRow - 2 * G.line)
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
	hideAll({ sim.dps, sim.error, sim.date, sim.button })
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

--- Draw `model` (MeTab.model): the columns, the footer's line, the window's height. The body's height.
function MeFrame.draw(model)
	local parts = MeFrame.parts
	if parts == nil then
		return 0
	end
	local G = MeFrame.LAYOUT
	local left = drawScore(parts.score, parts.left, model.score, 0)
	left = drawRaid(parts.raid, parts.left, model.lastRaid, left)
	local right = drawNow(parts.now, parts.right, model.rightNow, 0)
	if model.setup ~= nil and MeFrame.setup ~= nil then
		right = MeFrame.setup.draw(parts.right, Theme.SIZES.windowRight, model.setup, right)
	elseif MeFrame.setup ~= nil then
		MeFrame.setup.hide()
	end
	right = drawSim(parts.sim, parts.right, model.sim, right)
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
	figures = MeFrame.figures,
	waiting = MeFrame.waiting,
})

ns.MeFrame = MeFrame
return MeFrame
