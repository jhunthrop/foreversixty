-- addon/ForeverSixty/surfaces/OfficerTab.lua
-- The Forever Sixty window's Officer tab (design/specs/2026-11-13-addon-hub.md 4.6, boards hub-officer.png
-- and hub-officer-loot.png): who is coming tonight, are they ready, and what do I still have to chase?
--
-- Left: the Roster (fixed columns summing to the column: Raider with the class crest and the addon dot,
-- Score, Att., Sign-up, Gap, Ench., iLvl), raiders by default, a heading click sorts, the filters under it.
-- Right: Tonight's readiness, Sign-ups with the site's queue and Remind, Push a build with its in-window
-- confirm, and Master loot's one line between sessions. While the master looter runs a session the left
-- column is the session (the item, a row per member with a gain, Give twice) and the right column the
-- boss's items; Escape and the gold Roster go back. Shown behind Hub.officer(), or to the master looter
-- during a session with the session only. Model: Officer.lua, OfficerLoot.lua; drawing: HubDraw.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Compat = ns.Compat or require("Compat")
local Follow = ns.Follow or require("Follow")
local Score = ns.Score or require("Score")
local LootCouncil = ns.LootCouncil or require("LootCouncil")
local RaidDrops = ns.RaidDrops or require("RaidDrops")
local Chip = ns.Chip or require("Chip")
local GuildRaid = ns.GuildRaid or require("GuildRaid")
local ReadyCheck = ns.ReadyCheck or require("ReadyCheck")
local Officer = ns.Officer or require("Officer")
local OfficerLoot = ns.OfficerLoot or require("OfficerLoot")
local BuildsPane = ns.BuildsPane or require("BuildsPane")
local Comms = ns.Comms or require("Comms")
local HubDraw = ns.HubDraw or require("HubDraw")
local Hub = ns.Hub or require("Hub")

local OfficerTab = {}

--- The roster's columns, left to right: the sort key, the width (they and the gaps sum to the left
--- column's 388), whether the cell is right-aligned.
OfficerTab.COLUMNS = {
	{ key = "name", width = 104 }, { key = "score", width = 40, right = true },
	{ key = "attendance", width = 38, right = true }, { key = "signup", width = 72 },
	{ key = "gap", width = 38, right = true }, { key = "enchants", width = 36, right = true },
	{ key = "ilvl", width = 24, right = true },
}
--- The master loot session's columns: raider, gain, the worn item, Won, Att.; Give takes the rest.
OfficerTab.LOOT_COLUMNS = {
	{ key = "raider", width = 86 }, { key = "gain", width = 38, right = true }, { key = "wears", width = 68 },
	{ key = "won", width = 28, right = true }, { key = "att", width = 34, right = true },
}
--- Pixel layout: a roster row, the column heading row, the gap between columns, the crest and its gap,
--- the addon dot and its gap, the gap between sections, readiness's label and count columns, the push
--- icon, the session item's icon, an item row's icon, the Give button's least width, the least and most
--- roster rows, the room kept off the screen's height.
OfficerTab.LAYOUT = {
	row = 18, headRow = 16, colGap = 6, crest = 14, crestGap = 4, dot = 6, dotGap = 4, sectionGap = 20,
	readyLabel = 84, readyCount = 48, readyGap = 6, buildIcon = 20, itemIcon = 32, listIcon = 20, giveWidth = 40,
	minRows = 8, maxRows = 30, screenMargin = 160, iconGap = 8, wearsInset = 4,
}

local SEP = L.dialogSeparator

--- The roster's state: the filter, the sort, the mode ("roster" or "loot"), the push confirm open, the
--- footer's flash, the last model.
OfficerTab.filter = "raiders"
OfficerTab.sort = nil
--- The first roster row shown (a long roster pages by the window's rows).
OfficerTab.offset = 0
OfficerTab.mode = "roster"
OfficerTab.confirming = false

-- The model --------------------------------------------------------------------------------------------

--- The raid tonight's sections read: the soonest raid the inbox names (GuildRaid's), or nil.
local function nextRaid()
	return GuildRaid.liveRaids()[1]
end

--- How many roster rows the window holds: the screen's height less the window's chrome and the roster's
--- own lines, in rows, between LAYOUT.minRows and maxRows.
function OfficerTab.maxRows(screenHeight)
	local G, H = OfficerTab.LAYOUT, Hub.LAYOUT
	if type(screenHeight) ~= "number" or screenHeight <= 0 then
		return G.maxRows
	end
	local chrome = H.header + H.strip + H.bodyPad * 2 + H.footer + G.screenMargin
	local roster = HubDraw.LAYOUT.head + Theme.SIZES.gap + G.headRow + HubDraw.LAYOUT.line * 2 + 6
	return math.max(G.minRows, math.min(G.maxRows, math.floor((screenHeight - chrome - roster) / G.row)))
end

local function screenHeight()
	local ok, height = pcall(function()
		return UIParent:GetHeight()
	end)
	return ok and height or nil
end

--- The whole tab's model now: { roster = { rows, counts, shown, cap }, raid, readiness, readyNote,
--- signups, site, unsigned, push, lootLine, session }.
function OfficerTab.model()
	local raid = nextRaid()
	local input = Officer.input(raid)
	input.filter, input.sort = OfficerTab.filter, OfficerTab.sort
	local rows, counts = Officer.rows(input)
	local model = { raid = raid, roster = { rows = rows, counts = counts, cap = OfficerTab.maxRows(screenHeight()) } }
	local everyone = rows
	if OfficerTab.filter ~= "all" then
		everyone = Officer.rows({ roster = input.roster, record = input.record, seen = input.seen,
			inspect = input.inspect, own = input.own, members = input.members, ranks = input.ranks, now = input.now,
			data = input.data, filter = "all" })
	end
	local coming = {}
	for _, row in ipairs(everyone) do
		local signed = counts.signups and (row.answer == "in" or row.answer == "late")
		if signed or (not counts.signups and row.raider and row.online) then
			coming[#coming + 1] = row
		end
	end
	local zone = raid ~= nil and raid.raid or nil
	if raid ~= nil then
		local _, clock = GuildRaid.dayAndClock(raid.start)
		model.readyNote = string.format(L.offReadyNote, raid.raid, clock or "")
	end
	if not counts.signups then
		model.readyNote = L.offReadyOnline
	end
	model.readiness = Officer.readiness(coming, { zone = zone, data = input.data, pushes = Officer.pushesNow(),
		version = Comms.addonVersion(), roster = raid ~= nil and raid.roster or nil, check = ReadyCheck.result })
	model.nobodyLive = #coming > 0 and model.readiness[1].empty == true
	if raid ~= nil and raid.roster ~= nil then
		model.signups = GuildRaid.signups(raid.roster)
		model.unsigned = counts.signups and counts.unsigned or raid.roster.signups.none
		local day, clock = GuildRaid.dayAndClock(raid.start)
		model.remind = Officer.remindText(raid.raid, day, clock, model.unsigned)
	end
	model.site = Officer.siteLine(Officer.guildMessage())
	model.push = Officer.pushModel(input.data, Follow.build, Officer.sentPushes(Theme.playerGuild()), everyone)
	model.lootLine = OfficerLoot.lastNight(RaidDrops.nights())
	return model
end

--- The header's figures (spec 3.2): the guild's name and the player's rank, or Master looter in a session.
function OfficerTab.figures()
	if type(GetGuildInfo) ~= "function" then
		return {}
	end
	local ok, name, rankName = pcall(GetGuildInfo, "player")
	if not ok then
		return {}
	end
	if LootCouncil.active() and LootCouncil.isMasterLooter() then
		return { name, L.offMasterLooter }
	end
	return { name, rankName }
end

--- What waits on this tab (the dot and the chip's line): a master loot session open.
function OfficerTab.waiting()
	if not LootCouncil.active() or not LootCouncil.isMasterLooter() then
		return {}
	end
	local drops = LootCouncil.shown()
	local drop = drops[#drops]
	return { { key = "loot", fact = drop.lootId,
		line = string.format(L.offLootChip, drop.name or L.tooltipNamePending, LootCouncil.answeredText(drop)) } }
end

-- Drawing helpers -------------------------------------------------------------------------------------------

local place, hit, cover = HubDraw.place, HubDraw.hit, HubDraw.cover
local heading, textLine, nameText = HubDraw.heading, HubDraw.textLine, HubDraw.nameText

local function itemText(itemId)
	return HubDraw.itemText(itemId, OfficerTab.redraw)
end

--- A text cell at `x`, `width` wide, cut with an ellipsis, right-aligned when `right`.
local function cell(view, parent, key, x, y, width, text, hexKey, right)
	local _, region = textLine(view, parent, key, x, y, width, text, hexKey)
	region:SetJustifyH(right and "RIGHT" or "LEFT")
	return region
end

--- A button from the tab's pool by `key`, labelled `text`, its click `onClick` (set every draw).
local function button(view, parent, key, text, least, onClick)
	local made = view.buttons[key]
	if made == nil then
		made = Widgets.clientButton(parent, text, function(self)
			if self.foreverSixtyAction ~= nil then
				self.foreverSixtyAction()
			end
		end)
		view.buttons[key] = made
	end
	view.usedButtons[key] = true
	made.foreverSixtyAction = onClick
	Widgets.resetButtonLabel(made, text)
	Widgets.fitButton(made, least)
	Widgets.setSelected(made, false)
	made:Show()
	return made
end

--- The 6 px addon dot from the pool by `key`: success for live, muted for heard earlier.
local function dot(view, parent, key, state)
	local made = view.dots[key]
	if made == nil then
		made = Theme.texture(parent, "OVERLAY", "success")
		made:SetSize(OfficerTab.LAYOUT.dot, OfficerTab.LAYOUT.dot)
		if type(made.SetMask) == "function" then
			made:SetMask(Theme.MEDIA.minimapMask)
		end
		view.dots[key] = made
	end
	view.usedDots[key] = true
	Theme.paint(made, state == "live" and "success" or "muted")
	return made
end

--- "3 h", "2 days": how old a report is.
local function ageText(seconds)
	if seconds >= 2 * 86400 then
		return string.format(L.offDays, math.floor(seconds / 86400))
	elseif seconds >= 3600 then
		return string.format(L.offHours, math.floor(seconds / 3600))
	end
	return string.format(L.offMinutes, math.max(1, math.floor(seconds / 60)))
end

-- The roster -----------------------------------------------------------------------------------------------------

--- A roster row's cells, each { text, hex }. Pure.
function OfficerTab.cells(row)
	local dash = { text = L.offDash, hex = "muted" }
	local cells = { score = dash, attendance = dash, signup = dash, gap = dash, enchants = dash, ilvl = dash }
	if type(row.score) == "number" then
		cells.score = { text = string.format("%d", row.score), tierHex = Theme.panelTierHex(row.tier, Theme.HEX.body) }
	end
	if type(row.attendance) == "number" then
		cells.attendance = { text = string.format(L.guildPercent, row.attendance) }
	end
	if row.answer == "none" then
		cells.signup = { text = L.offNotSigned, hex = "muted" }
	elseif row.answer ~= nil then
		cells.signup = { text = L.guildAnswers[row.answer], hex = GuildRaid.ANSWER_HEX[row.answer] }
	end
	if type(row.gap) == "number" then
		if row.gap <= 0 then
			cells.gap = { text = L.offBest, hex = "muted" }
		elseif row.unit == "pts" then
			cells.gap = { text = string.format(L.offGapPts, math.floor(row.gap + 0.5)) }
		else
			cells.gap = { text = string.format("%.1f", row.gap) }
		end
	end
	if type(row.enchants) == "number" then
		cells.enchants = { text = tostring(row.enchants), hex = row.enchants > 0 and "warning" or "muted" }
	end
	if type(row.ilvl) == "number" then
		cells.ilvl = { text = string.format("%d", math.floor(row.ilvl + 0.5)), hex = "muted" }
	end
	return cells
end

--- A roster row's hover: the full name, level and spec, the score line, the best parse, the build, the
--- bare slots, the last report's age, offline.
function OfficerTab.rowHint(row)
	local lines = { nameText(row.name, row.classToken) }
	local spec = row.spec ~= nil and Follow.specName(row.spec) or ""
	if type(row.level) == "number" then
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", (string.format(L.offLevel, row.level, spec):gsub("%s+$", "")))
	end
	local scoreLines = Score.lines(row.record)
	if scoreLines[1] ~= nil then
		lines[#lines + 1] = scoreLines[1]
	end
	if row.bestParse ~= nil then
		lines[#lines + 1] = Theme.inlineColor("tooltipText", string.format(L.groupBestParse,
			Chip.ordinal(row.bestParse.percentile), row.bestParse.boss))
	end
	if row.buildName ~= nil then
		lines[#lines + 1] = Theme.inlineColor("tooltipText", string.format(L.groupBuild, row.buildName))
	end
	if type(row.bareSlots) == "table" and #row.bareSlots > 0 then
		local words = {}
		for index, slot in ipairs(row.bareSlots) do
			words[index] = L.gearSlotLabels[slot] or slot
		end
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", string.format(L.groupBareSlots, table.concat(words, ", ")))
	end
	if row.live == nil then
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", L.offNoAddonHeard)
	elseif row.age ~= nil and row.age > 0 then
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", string.format(L.offHeardAgo, ageText(row.age)))
	end
	if not row.online then
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", L.groupOffline)
	end
	return lines
end

local function columnX(columns, index)
	local x = 0
	for position = 1, index - 1 do
		x = x + columns[position].width + OfficerTab.LAYOUT.colGap
	end
	return x
end

--- The column headings: a click sorts by that column (or turns it round).
local function drawHeadings(view, parent, y)
	for index, column in ipairs(OfficerTab.COLUMNS) do
		local x = columnX(OfficerTab.COLUMNS, index)
		local region = cell(view, parent, "rh." .. column.key, x, y, column.width, L.offColumns[column.key], "muted",
			column.right)
		local active = OfficerTab.sort ~= nil and OfficerTab.sort.key == column.key
		region:SetTextColor(Theme.rgb(Theme.HEX[active and "gold" or "muted"]))
		local frame = hit(view, parent, "rh." .. column.key)
		frame:ClearAllPoints()
		frame:SetPoint("TOPLEFT", parent, "TOPLEFT", x, -y)
		frame:SetSize(column.width, OfficerTab.LAYOUT.headRow)
		frame:SetScript("OnEnter", nil)
		frame:SetScript("OnClick", function()
			OfficerTab.sort = Officer.nextSort(OfficerTab.sort, column.key)
			OfficerTab.offset = 0
			OfficerTab.redraw()
		end)
	end
	return y + OfficerTab.LAYOUT.headRow
end

--- One roster row at `y`.
local function drawRow(view, parent, row, index, y)
	local G = OfficerTab.LAYOUT
	local key = "r" .. index
	local old = row.age ~= nil and row.age > Officer.OLD_REPORT
	local cells = OfficerTab.cells(row)
	local top = y + (G.row - HubDraw.LAYOUT.line) / 2
	HubDraw.crest(view, parent, key, G.crest, row.classToken, 0, y + (G.row - G.crest) / 2)
	local nameX = G.crest + G.crestGap
	local room = OfficerTab.COLUMNS[1].width - nameX - (row.live ~= nil and (G.dot + G.dotGap) or 0)
	local name = cell(view, parent, key .. ".m", nameX, top, room, nameText(row.short, row.classToken), "body")
	if row.live ~= nil then
		local mark = dot(view, parent, key .. ".d", row.live)
		mark:ClearAllPoints()
		mark:SetPoint("LEFT", name, "LEFT", math.min(room, name:GetStringWidth() or room) + G.dotGap, 0)
		mark:Show()
	end
	for position = 2, #OfficerTab.COLUMNS do
		local column = OfficerTab.COLUMNS[position]
		local value = cells[column.key]
		local region = cell(view, parent, key .. "." .. column.key, columnX(OfficerTab.COLUMNS, position), top,
			column.width, value.text, value.hex or "body", column.right)
		if value.tierHex ~= nil then
			region:SetTextColor(Theme.rgb(value.tierHex))
		end
		if old and value.hex == nil then
			region:SetTextColor(Theme.rgb(Theme.HEX.muted))
		end
	end
	local frame = hit(view, parent, key)
	frame:ClearAllPoints()
	frame:SetPoint("TOPLEFT", parent, "TOPLEFT", 0, -y)
	frame:SetSize(Theme.SIZES.windowLeft, G.row)
	HubDraw.nameClicks(frame, row.name, function()
		return OfficerTab.rowHint(row)
	end)
	return y + G.row
end

--- Gold links side by side at `y`: each { key, text, onClick }, or plain muted text without onClick.
local function drawLinks(view, parent, parts, y)
	local x = 0
	for index, part in ipairs(parts) do
		if index > 1 then
			local _, sep = textLine(view, parent, "lk.s" .. index, x, y, 20, SEP, "muted")
			x = x + (sep:GetStringWidth() or 8)
		end
		local _, region = textLine(view, parent, "lk." .. part.key, x, y, Theme.SIZES.windowLeft - x, part.text,
			part.onClick ~= nil and "gold" or "muted")
		local width = region:GetStringWidth() or 0
		if part.onClick ~= nil then
			local frame = hit(view, parent, "lk." .. part.key)
			cover(frame, region, width)
			frame:SetScript("OnEnter", nil)
			frame:SetScript("OnClick", part.onClick)
		end
		x = x + width
	end
	return y + HubDraw.LAYOUT.line
end

local function setFilter(filter)
	return function()
		OfficerTab.filter, OfficerTab.offset = filter, 0
		OfficerTab.redraw()
	end
end

local function setOffset(offset)
	return function()
		OfficerTab.offset = offset
		OfficerTab.redraw()
	end
end

--- The line under the rows: "16 of 31 shown", then the filters the roster is not on.
local function drawFilters(view, parent, roster, y)
	local counts = roster.counts
	local offset = math.min(OfficerTab.offset, math.max(0, #roster.rows - 1))
	local shown = math.min(#roster.rows - offset, roster.cap)
	local total = OfficerTab.filter == "all" and counts.total or counts.raiders
	local parts = { { key = "shown", text = offset > 0 and string.format(L.offShownRange, offset + 1, offset + shown,
		total) or string.format(L.offShown, shown, total) } }
	local after = #roster.rows - offset - shown
	if after > 0 then
		parts[#parts + 1] = { key = "more", text = string.format(L.offMore, math.min(after, roster.cap)),
			onClick = setOffset(offset + shown) }
	elseif offset > 0 then
		parts[#parts + 1] = { key = "more", text = L.offFirst, onClick = setOffset(0) }
	end
	if OfficerTab.filter ~= "all" then
		parts[#parts + 1] = { key = "all", text = L.offShowAll, onClick = setFilter("all") }
	end
	if OfficerTab.filter ~= "raiders" then
		parts[#parts + 1] = { key = "raiders", text = L.offRaiders, onClick = setFilter("raiders") }
	end
	if counts.signups and OfficerTab.filter ~= "unsigned" and counts.unsigned > 0 then
		parts[#parts + 1] = { key = "unsigned", text = string.format(L.offUnsigned, counts.unsigned),
			onClick = setFilter("unsigned") }
	end
	if OfficerTab.filter ~= "online" then
		parts[#parts + 1] = { key = "online", text = string.format(L.offOnline, counts.online),
			onClick = setFilter("online") }
	end
	return drawLinks(view, parent, parts, y)
end

--- The legend: the two dots and what Gap counts in.
local function drawLegend(view, parent, y)
	local G = OfficerTab.LAYOUT
	local x = 0
	for _, entry in ipairs({ { "live", L.offLegendLive }, { "heard", L.offLegendHeard } }) do
		local mark = dot(view, parent, "lg." .. entry[1], entry[1])
		mark:ClearAllPoints()
		mark:SetPoint("TOPLEFT", parent, "TOPLEFT", x, -(y + (HubDraw.LAYOUT.line - G.dot) / 2))
		mark:Show()
		x = x + G.dot + G.dotGap
		local _, region = textLine(view, parent, "lg.t" .. entry[1], x, y, 160, entry[2], "muted")
		x = x + (region:GetStringWidth() or 60) + G.iconGap
	end
	textLine(view, parent, "lg.gap", x - G.iconGap, y, Theme.SIZES.windowLeft - x, SEP .. L.offLegendGap, "muted")
	return y + HubDraw.LAYOUT.line
end

local function sortNote(signups)
	local sort = OfficerTab.sort
	if sort == nil then
		return signups and L.offBySignup or L.offNoSignups
	end
	return string.format(L.offBy, L.offColumns[sort.key])
end

--- The left column on the roster: Roster with its headings, rows, filters and legend. Its height.
local function drawRoster(view, model)
	local parent, width = view.left, Theme.SIZES.windowLeft
	local roster = model.roster
	local y = heading(view, parent, "roster", width, 0, L.offRoster, sortNote(roster.counts.signups))
	if roster.counts.total == 0 then
		return textLine(view, parent, "noroster", 0, y, width, L.offNoRoster, "muted", true)
	end
	y = drawHeadings(view, parent, y)
	local offset = math.min(OfficerTab.offset, math.max(0, #roster.rows - 1))
	for index = offset + 1, math.min(#roster.rows, offset + roster.cap) do
		y = drawRow(view, parent, roster.rows[index], index - offset, y)
	end
	if #roster.rows == 0 then
		y = textLine(view, parent, "nobody", 0, y, width, L.offNobodyShown, "muted")
	end
	y = drawFilters(view, parent, roster, y + 4)
	return drawLegend(view, parent, y + 2)
end

-- The right column ----------------------------------------------------------------------------------------------

--- A readiness row's names, class-coloured, a count after each where it has one, the first `shown` of
--- them and a muted "+N" for the rest.
local function readyNames(entry, shown)
	if entry.open ~= nil and #entry.open > 0 then
		return Theme.inlineColor("warning", string.format(L.offReadyOpen, table.concat(entry.open, ", ")))
	end
	if entry.needs ~= nil then
		return #entry.needs > 0 and Theme.inlineColor("warning", table.concat(entry.needs, ", "))
			or Theme.inlineColor("muted", L.offReadyCovered)
	end
	local parts = {}
	for index, row in ipairs(entry.names) do
		if index > (shown or #entry.names) then
			break
		end
		local text = nameText(row.short, row.classToken)
		if entry.perName ~= nil and row[entry.perName] ~= nil then
			text = text .. " " .. row[entry.perName]
		end
		parts[#parts + 1] = text
	end
	local text = table.concat(parts, ", ")
	local rest = #entry.names - #parts
	if rest > 0 then
		text = text .. " " .. Theme.inlineColor("muted", string.format(L.offMoreNames, rest))
	elseif entry.tail ~= nil and #parts > 0 then
		text = text .. " " .. entry.tail
	end
	return text
end

--- `entry`'s names on `region`, as many whole names as fit `width`, the rest a muted "+N".
local function fitNames(region, entry, width)
	for shown = #entry.names, 1, -1 do
		region:SetText(readyNames(entry, shown))
		local drawn = type(region.GetUnboundedStringWidth) == "function" and region:GetUnboundedStringWidth()
			or region:GetStringWidth()
		if type(drawn) ~= "number" or drawn <= width then
			return
		end
	end
end

--- A readiness row's hover: every name, and what the ready check found each missing.
local function readyHint(entry)
	local lines = { entry.title or entry.label }
	for _, row in ipairs(entry.names) do
		local line = nameText(row.short, row.classToken)
		if row.what ~= nil then
			line = string.format(L.groupMissing, line, table.concat(row.what, ", "))
		elseif entry.key == "build" and row.buildName ~= nil then
			line = line .. Theme.inlineColor("tooltipMuted", SEP .. row.buildName)
		end
		lines[#lines + 1] = line
	end
	for _, line in ipairs(entry.hint or {}) do
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", line)
	end
	return lines
end

--- Tonight's readiness: label, count, names; the y under it.
local function drawReadiness(view, parent, width, model, y)
	local G = OfficerTab.LAYOUT
	y = heading(view, parent, "ready", width, y, L.offReadiness, model.readyNote)
	if model.nobodyLive then
		y = textLine(view, parent, "nobodylive", 0, y, width, L.offNobodyAddon, "muted", true)
	end
	for index, entry in ipairs(model.readiness) do
		if not (model.nobodyLive and entry.key == "addon") then
			local key = "rd" .. index
			cell(view, parent, key .. ".l", 0, y, G.readyLabel, entry.label, "muted")
			local count = entry.of ~= nil and string.format(L.offOf, entry.count, entry.of) or tostring(entry.count)
			cell(view, parent, key .. ".c", G.readyLabel, y, G.readyCount, count, entry.warn and "warning" or "strong", true)
			local x = G.readyLabel + G.readyCount + G.readyGap
			local _, names = textLine(view, parent, key .. ".v", x, y, width - x, readyNames(entry), "body")
			fitNames(names, entry, width - x)
			local frame = hit(view, parent, key)
			frame:ClearAllPoints()
			frame:SetPoint("TOPLEFT", parent, "TOPLEFT", 0, -y)
			frame:SetSize(width, HubDraw.LAYOUT.line)
			frame:SetScript("OnClick", nil)
			frame:SetScript("OnEnter", (#entry.names > 0 or #(entry.hint or {}) > 0) and function(self)
				Theme.showLines(self, readyHint(entry), "ANCHOR_RIGHT")
			end or nil)
			y = y + math.max(HubDraw.LAYOUT.line, HubDraw.textHeight(names)) + 3
		end
	end
	return y
end

--- Sign-ups: the answers, the roles, the site's queue and Remind N in chat.
local function drawSignups(view, parent, width, model, y)
	y = heading(view, parent, "signups", width, y, L.guildSignups, model.signups ~= nil and model.signups.note or nil)
	if model.signups ~= nil then
		local parts = {}
		for _, count in ipairs(model.signups.counts) do
			parts[#parts + 1] = Theme.inlineColor(count.hex, count.text)
		end
		y = textLine(view, parent, "su.c", 0, y, width, table.concat(parts, Theme.inlineColor("muted", SEP)), "body", true)
		if model.signups.roles ~= nil then
			local roles = {}
			for _, role in ipairs(model.signups.roles) do
				roles[#roles + 1] = Theme.inlineColor(role.hex, role.text)
			end
			y = textLine(view, parent, "su.r", 0, y, width, table.concat(roles, Theme.inlineColor("muted", SEP)), "body",
				true)
		end
	end
	local remind
	if model.remind ~= nil and (model.unsigned or 0) > 0 then
		remind = button(view, parent, "remind", string.format(L.offRemind, model.unsigned), 64, function()
			if Officer.remind(model.remind) then
				OfficerTab.flash(L.offReminded)
			end
		end)
	end
	local top = y + 6
	if model.site ~= nil or remind ~= nil then
		local room = width - (remind ~= nil and remind:GetWidth() + 8 or 0)
		local _, site = textLine(view, parent, "su.s", 0, top + 2, room, model.site or "", "muted", true)
		if remind ~= nil then
			remind:ClearAllPoints()
			remind:SetPoint("TOPRIGHT", parent, "TOPRIGHT", 0, -top)
		end
		y = math.max(top + 2 + HubDraw.textHeight(site), top + (remind ~= nil and Theme.SIZES.clientButtonHeight or 0))
	end
	return y
end

--- Push a build: the officer's build, when it went and how many follow it, Send to guild; the confirm
--- in place of the button once clicked.
local function drawPush(view, parent, width, push, y)
	local G = OfficerTab.LAYOUT
	y = heading(view, parent, "push", width, y, L.offPush, nil)
	local x = 0
	local path = push.spec ~= nil and BuildsPane.specIcon(Officer.data or ns.Data, push.spec) or nil
	if path ~= nil then
		local icon = HubDraw.icon(view, parent, "push", G.buildIcon)
		icon:SetTexture(path)
		place(icon, parent, 0, y + 2)
		x = G.buildIcon + G.iconGap
	end
	local top = y
	if OfficerTab.confirming and push.canPush then
		y = textLine(view, parent, "pu.q", 0, y, width, string.format(L.offPushConfirm, push.name), "body", true) + 6
		local yes = button(view, parent, "push.yes", L.offPushYes, 64, function()
			OfficerTab.confirming = false
			local fields = Officer.pushFields(Officer.data or ns.Data, Follow.build)
			local ok, why = false, L.offPushNoCode
			if fields ~= nil then
				ok, why = Officer.push(fields, Follow.build)
			end
			OfficerTab.pushFailed = not ok and string.format(L.offPushFailed, tostring(why or "")) or nil
			if ok then
				OfficerTab.flash(L.offSentToGuild)
			end
			OfficerTab.redraw()
		end)
		yes:ClearAllPoints()
		yes:SetPoint("TOPLEFT", parent, "TOPLEFT", 0, -y)
		local no = button(view, parent, "push.no", L.offCancel, 64, function()
			OfficerTab.confirming = false
			OfficerTab.redraw()
		end)
		no:ClearAllPoints()
		no:SetPoint("LEFT", yes, "RIGHT", 4, 0)
		return y + Theme.SIZES.clientButtonHeight
	end
	local send = button(view, parent, "push.send", L.offSend, 64, function()
		OfficerTab.confirming = true
		OfficerTab.redraw()
	end)
	Widgets.setEnabled(send, push.canPush)
	send:ClearAllPoints()
	send:SetPoint("TOPRIGHT", parent, "TOPRIGHT", 0, -top)
	send:SetScript("OnEnter", not push.canPush and function(self)
		Theme.showLines(self, { push.why }, "ANCHOR_RIGHT")
	end or nil)
	send:SetScript("OnLeave", function()
		Theme.hideTooltip()
	end)
	local room = width - x - send:GetWidth() - 8
	y = textLine(view, parent, "pu.n", x, y, room, push.name or push.why, push.name ~= nil and "body" or "muted")
	local sub
	if OfficerTab.pushFailed ~= nil then
		sub = Theme.inlineColor("warning", OfficerTab.pushFailed)
	elseif push.sent ~= nil then
		sub = string.format(L.offPushSent, push.sent, push.follow or 0)
	elseif push.name ~= nil then
		sub = L.offPushNotSent
	end
	if sub ~= nil then
		y = textLine(view, parent, "pu.s", x, y, room, sub, "muted")
	end
	return math.max(y, top + Theme.SIZES.clientButtonHeight)
end

--- Master loot between sessions: one muted line, and the night's count.
local function drawLootLine(view, parent, width, model, y)
	y = heading(view, parent, "lootline", width, y, L.offLoot, nil)
	local text = L.offLootIdle
	if model.lootLine ~= nil then
		text = text .. " " .. model.lootLine .. "."
	end
	return textLine(view, parent, "ll", 0, y, width, text, "muted", true)
end

local function drawRight(view, model)
	local parent, width = view.right, Theme.SIZES.windowRight
	local G = OfficerTab.LAYOUT
	local y = drawReadiness(view, parent, width, model, 0)
	if model.signups ~= nil or model.site ~= nil then
		y = drawSignups(view, parent, width, model, y + G.sectionGap - 3)
	end
	y = drawPush(view, parent, width, model.push, y + G.sectionGap)
	return drawLootLine(view, parent, width, model, y + G.sectionGap)
end

-- The master loot session -------------------------------------------------------------------------------------------

--- The summary line under the session's rows: who answered, who has no gain, can't score, no addon.
local function summaryText(view)
	local summary = view.summary
	local parts = { string.format(L.offAnswered, summary.answered, summary.asked) }
	local function names(list, format)
		if #list == 0 then
			return
		end
		local shown = {}
		for _, row in ipairs(list) do
			shown[#shown + 1] = nameText(row.short, row.classToken)
		end
		parts[#parts + 1] = string.format(format, table.concat(shown, ", "))
	end
	names(summary.noGain, L.offNoGain)
	names(summary.cantScore, L.offCantScore)
	names(summary.noAddon, L.offNoAddon)
	return table.concat(parts, Theme.inlineColor("muted", SEP))
end

--- One session row: crest, name and dot, the gain, the worn item, Won, Att., on list, Give.
local function drawLootRow(view, parent, session, row, index, y)
	local G, columns = OfficerTab.LAYOUT, OfficerTab.LOOT_COLUMNS
	local key = "lr" .. index
	local top = y + (G.row - HubDraw.LAYOUT.line) / 2 + 2
	local rowHeight = Theme.SIZES.clientButtonHeight + 4
	HubDraw.crest(view, parent, key, G.crest, row.classToken, 0, y + (rowHeight - G.crest) / 2)
	local nameX = G.crest + G.crestGap
	local room = columns[1].width - nameX - (row.live and (G.dot + G.dotGap) or 0)
	local name = cell(view, parent, key .. ".m", nameX, top, room, nameText(row.short, row.classToken), "body")
	if row.live then
		local mark = dot(view, parent, key .. ".d", "live")
		mark:ClearAllPoints()
		mark:SetPoint("LEFT", name, "LEFT", math.min(room, name:GetStringWidth() or room) + G.dotGap, 0)
		mark:Show()
	end
	cell(view, parent, key .. ".g", columnX(columns, 2), top, columns[2].width, row.gain or L.offDash,
		row.gain ~= nil and "success" or "muted", true)
	local wearsX = columnX(columns, 3) + G.wearsInset
	if row.worn ~= nil and row.worn ~= 0 then
		local icon = HubDraw.icon(view, parent, key .. ".wi", G.crest)
		icon:SetTexture(Compat.itemIcon(row.worn) or Theme.UNKNOWN_ICON)
		place(icon, parent, wearsX, y + (rowHeight - G.crest) / 2)
		local _, worn = textLine(view, parent, key .. ".w", wearsX + G.crest + G.crestGap, top,
			columns[3].width - G.wearsInset - G.crest - G.crestGap, itemText(row.worn) or L.tooltipNamePending, "body")
		local frame = hit(view, parent, key .. ".w")
		HubDraw.itemClicks(frame, row.worn)
		cover(frame, worn, columns[3].width - G.wearsInset - G.crest - G.crestGap)
	elseif row.gain ~= nil then
		cell(view, parent, key .. ".w", wearsX, top, columns[3].width - G.wearsInset, L.lootEmpty, "muted")
	end
	cell(view, parent, key .. ".n", columnX(columns, 4), top, columns[4].width, tostring(row.won), "body", true)
	cell(view, parent, key .. ".a", columnX(columns, 5), top, columns[5].width,
		row.attendance ~= nil and string.format(L.guildPercent, row.attendance) or L.offDash, "muted", true)
	local label = row.armed and string.format(L.offGiveArmed, row.short) or L.offGive
	local give = button(view, parent, key .. ".give", label, G.giveWidth, function()
		local _, why = OfficerLoot.give(session.drop, row, index == 1 and row.gain ~= nil)
		if why ~= nil then
			OfficerTab.flash(why)
		end
	end)
	Widgets.setSelected(give, row.armed)
	-- An armed Give stays inside its column: a name too long for it reads "Give?", the footer names them.
	local giveRoom = Theme.SIZES.windowLeft - columnX(columns, #columns + 1)
	if row.armed and give:GetWidth() > giveRoom then
		Widgets.resetButtonLabel(give, L.offGiveArmedShort)
		Widgets.fitButton(give, G.giveWidth)
	end
	local eligible = session.drop.slot == nil or LootCouncil.candidateIndex(session.drop.slot, row.short) ~= nil
	Widgets.setEnabled(give, eligible and session.drop.givenTo == nil)
	give:ClearAllPoints()
	give:SetPoint("TOPRIGHT", parent, "TOPRIGHT", 0, -(y + 2))
	if row.planned and not row.armed then
		local _, tag = textLine(view, parent, key .. ".p", 0, top, 60, L.offOnList, "gold")
		tag:ClearAllPoints()
		tag:SetPoint("RIGHT", give, "LEFT", -6, 0)
	end
	return y + rowHeight
end

--- The left column in a session: the item, the rows, the summary. Its height.
local function drawSession(view, session)
	local parent, width = view.left, Theme.SIZES.windowLeft
	local G, columns = OfficerTab.LAYOUT, OfficerTab.LOOT_COLUMNS
	local note = string.format(L.offLootItemOf, session.index, session.count)
	local headNote = Hub.officer() and Theme.inlineColor("gold", L.offRoster) .. Theme.inlineColor("muted", SEP .. note)
		or note
	local y = heading(view, parent, "session", width, 0, L.offLoot, headNote)
	if Hub.officer() then
		local back = hit(view, parent, "session.back")
		back:ClearAllPoints()
		back:SetPoint("TOPRIGHT", parent, "TOPRIGHT", 0, 0)
		back:SetSize(width / 2, HubDraw.LAYOUT.head)
		back:SetScript("OnEnter", nil)
		back:SetScript("OnClick", function()
			OfficerTab.mode = "roster"
			OfficerTab.redraw()
		end)
	end
	local drop = session.drop
	local icon = HubDraw.icon(view, parent, "session.icon", G.itemIcon)
	icon:SetTexture(Compat.itemIcon(drop.itemId) or Theme.UNKNOWN_ICON)
	place(icon, parent, 0, y)
	local x = G.itemIcon + G.iconGap
	local _, name = textLine(view, parent, "session.name", x, y + 1, width - x, itemText(drop.itemId)
		or L.tooltipNamePending, "body")
	local frame = hit(view, parent, "session.item")
	HubDraw.itemClicks(frame, drop.itemId)
	cover(frame, name)
	local where = {}
	where[#where + 1] = LootCouncil.itemSlotLabel(drop.itemId)
	where[#where + 1] = drop.boss
	textLine(view, parent, "session.where", x, y + HubDraw.LAYOUT.line + 2, width - x, table.concat(where, SEP), "muted")
	y = y + G.itemIcon + 8
	local unit = session.view.unit
	for index, column in ipairs(columns) do
		local text = column.key == "gain" and unit or L.offLootColumns[column.key]
		local inset = column.key == "wears" and G.wearsInset or 0
		local region = cell(view, parent, "lh." .. column.key, columnX(columns, index) + inset, y, column.width - inset,
			text, "muted", column.right)
		if column.key == "won" then
			local wonHit = hit(view, parent, "lh.won")
			cover(wonHit, region, column.width)
			wonHit:SetScript("OnClick", nil)
			wonHit:SetScript("OnEnter", function(self)
				Theme.showLines(self, { L.offWon, Theme.inlineColor("tooltipMuted", L.offWonHint) }, "ANCHOR_RIGHT")
			end)
		end
	end
	y = y + G.headRow
	for index, row in ipairs(session.view.rows) do
		y = drawLootRow(view, parent, session, row, index, y)
	end
	local summary = summaryText(session.view)
	if session.view.quiet then
		summary = L.offNoAnswers
	elseif #session.view.rows > 0 and session.view.rows[1].gain == nil then
		summary = L.offWaitingAnswers
	end
	return textLine(view, parent, "session.sum", 0, y + 4, width, summary, "muted", true)
end

--- The right column in a session: the boss's items and what became of each.
local function drawItems(view, session)
	local parent, width = view.right, Theme.SIZES.windowRight
	local G = OfficerTab.LAYOUT
	local count = #session.items
	local y = heading(view, parent, "items", width, 0, session.boss or L.offLoot,
		count == 1 and L.offOneItem or string.format(L.offItems, count))
	for index, item in ipairs(session.items) do
		local key = "it" .. index
		local icon = HubDraw.icon(view, parent, key, G.listIcon)
		icon:SetTexture(Compat.itemIcon(item.itemId) or Theme.UNKNOWN_ICON)
		place(icon, parent, 0, y + 2)
		local x = G.listIcon + G.iconGap
		local _, name = textLine(view, parent, key .. ".n", x, y, width - x, itemText(item.itemId) or L.tooltipNamePending,
			"body")
		local state
		if item.state == "given" then
			state = string.format(L.offItemTo, nameText(item.to, item.toClass), item.gain ~= nil
				and Theme.inlineColor("success", item.gain) or "")
		elseif item.state == "open" then
			state = L.offItemOpen
		else
			state = Theme.inlineColor("muted", item.answered)
		end
		textLine(view, parent, key .. ".s", x, y + HubDraw.LAYOUT.line, width - x, state, "body")
		local frame = hit(view, parent, key)
		local lootId = item.lootId
		HubDraw.itemClicks(frame, item.itemId, function()
			OfficerLoot.select(lootId)
		end)
		cover(frame, name, width - x)
		y = y + HubDraw.LAYOUT.line * 2 + 6
	end
	return y
end

-- The frame -------------------------------------------------------------------------------------------------------

--- The footer's right line now: a flash for 2 s, else the mode's hint.
local function hints(session)
	local escape = session ~= nil and Hub.officer() and L.offEscRoster or nil
	local right = session ~= nil and L.offGiveHint or L.offWhisperHint
	local armed = session ~= nil and OfficerLoot.armed or nil
	if armed ~= nil and armed.lootId == session.drop.lootId then
		right = string.format(L.offGiveAgain, armed.short, session.drop.name or L.tooltipNamePending)
	end
	if OfficerTab.flashText ~= nil then
		right = OfficerTab.flashText
	end
	Hub.setHints(escape, right)
end

--- Says `text` in the footer for 2 s (Sent to guild, why a Give did not go).
function OfficerTab.flash(text)
	local stamp = Theme.now()
	OfficerTab.flashText, OfficerTab.flashAt = text, stamp
	Theme.after(2, function()
		if OfficerTab.flashAt == stamp then
			OfficerTab.flashText = nil
			OfficerTab.redraw()
		end
	end)
	OfficerTab.redraw()
end

--- The session to draw, or nil: a session open and the tab on it (always, for a master looter without the
--- officer gate).
local function sessionToDraw()
	if OfficerTab.mode ~= "loot" and Hub.officer() then
		return nil
	end
	return OfficerLoot.session()
end

local function sweepButtons(view)
	for key, made in pairs(view.buttons) do
		if not view.usedButtons[key] then
			made:Hide()
		end
	end
	for key, made in pairs(view.dots) do
		if not view.usedDots[key] then
			made:Hide()
		end
	end
end

--- Redraw the tab: the roster and its sections, or the session; the footer; the window's height.
function OfficerTab.draw()
	local view = OfficerTab.view
	if view == nil then
		return nil
	end
	HubDraw.begin(view)
	view.usedButtons, view.usedDots = {}, {}
	local session = sessionToDraw()
	local left, right, model
	if session ~= nil then
		left, right = drawSession(view, session), drawItems(view, session)
	else
		model = OfficerTab.model()
		OfficerTab.current = model
		left, right = drawRoster(view, model), drawRight(view, model)
	end
	HubDraw.sweep(view)
	sweepButtons(view)
	view.rule:SetHeight(math.max(left, right, 1))
	hints(session)
	Hub.fitHeight(math.max(left, right))
	Hub.refresh()
	return session or model
end

--- The tab redrawn while it shows.
function OfficerTab.redraw()
	if OfficerTab.view ~= nil and OfficerTab.view.shown then
		OfficerTab.draw()
	end
end

function OfficerTab.build(content)
	local view = HubDraw.view(content)
	view.dots, view.usedButtons, view.usedDots = {}, {}, {}
	OfficerTab.view = view
	return view
end

function OfficerTab.show(options)
	local view = OfficerTab.view
	if view == nil then
		return nil
	end
	view.shown = true
	if options ~= nil and options.loot then
		OfficerTab.mode = "loot"
	end
	if Hub.officer() then
		Officer.askLive()
	end
	return OfficerTab.draw()
end

function OfficerTab.hide()
	if OfficerTab.view ~= nil then
		OfficerTab.view.shown = false
	end
	OfficerTab.confirming = false
	OfficerTab.mode = "roster"
	OfficerTab.offset = 0
	Theme.hideTooltip()
end

--- Escape: from the session back to the roster for an officer; else the window closes.
function OfficerTab.escape()
	if OfficerTab.mode == "loot" and Hub.officer() and OfficerLoot.session() ~= nil then
		OfficerTab.mode = "roster"
		OfficerTab.redraw()
		return true
	end
	if OfficerTab.confirming then
		OfficerTab.confirming = false
		OfficerTab.redraw()
		return true
	end
	return false
end

--- A new master loot session: the window opens on Officer with it, never in combat.
function OfficerTab.openSession()
	if type(InCombatLockdown) == "function" and InCombatLockdown() then
		return nil
	end
	OfficerTab.mode = "loot"
	if Hub.isShown() and Hub.active == "officer" then
		return OfficerTab.redraw()
	end
	return Hub.open("officer", { loot = true })
end

--- The events that change what the tab shows while it is open: the roster, the group, an inspect.
OfficerTab.EVENTS = { "GUILD_ROSTER_UPDATE", "GROUP_ROSTER_UPDATE", "INSPECT_READY", "LOOT_CLOSED" }

--- Runs at file load: one hidden event frame where the client has frames; Comms' live answers and the
--- ready check redraw the tab too.
function OfficerTab.register()
	OfficerLoot.onChange = OfficerTab.redraw
	OfficerLoot.openSession = OfficerTab.openSession
	ReadyCheck.onChange = OfficerTab.redraw
	Comms.on("snapshot!", OfficerTab.redraw)
	Comms.on("push", function(named, sender)
		Officer.hearPush(named, sender)
		OfficerTab.redraw()
	end)
	Comms.on("hello", OfficerTab.redraw)
	if type(CreateFrame) ~= "function" then
		return nil
	end
	local frame = CreateFrame("Frame", nil, UIParent)
	frame:SetScript("OnEvent", function()
		OfficerTab.redraw()
	end)
	for _, event in ipairs(OfficerTab.EVENTS) do
		Theme.registerEvent(frame, event)
	end
	return frame
end

OfficerTab.frame = OfficerTab.register()

Hub.register({
	id = "officer",
	label = L.hubTabOfficer,
	build = OfficerTab.build,
	show = OfficerTab.show,
	hide = OfficerTab.hide,
	escape = OfficerTab.escape,
	figures = OfficerTab.figures,
	waiting = OfficerTab.waiting,
})

ns.OfficerTab = OfficerTab
return OfficerTab
