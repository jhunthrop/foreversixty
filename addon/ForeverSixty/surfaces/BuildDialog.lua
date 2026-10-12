-- addon/ForeverSixty/surfaces/BuildDialog.lua
-- The build dialog (spec section 4.10): the one frame of the addon's own, 420x300.
-- Two tabs, one job each: paste a build in ("Paste a build"), copy the character out
-- ("Copy my character", the site's own words).
--
-- Pure model, thin frame: BuildDialog.evaluate is the paste box's whole verdict
-- (Codec.loadBuild's own answer, and the codec's own reason on a refusal);
-- preview and split are what the preview says; copiedLine, seenLine, statusLine and
-- simStateLine are the copy tab's lines with the clock handed in; inboxRows is the
-- "from the site" list. The frame is built on the first open, never at load, and
-- every number on it comes from the function that owns it (Progress.planSplit for the
-- split, Follow.buildName for the name).
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Cards = ns.Cards or require("Cards")
local Prefs = ns.Prefs or require("Prefs")
local Compat = ns.Compat or require("Compat")
local Codec = ns.Codec or require("Codec")
local Talents = ns.Talents or require("Talents")
local Export = ns.Export or require("Export")
local Follow = ns.Follow or require("Follow")
local Progress = ns.Progress or require("Progress")

local BuildDialog = {}

BuildDialog.FRAME_NAME = "ForeverSixtyBuildDialog"
--- Seconds the "Copied 14:02" line stays before the settled line comes back.
BuildDialog.COPIED_SECONDS = 3
--- How many of the build's planned items the preview shows.
BuildDialog.GEAR_PICKS = 3
BuildDialog.MODES = { "paste", "copy" }

--- Injection points Options fills: the data to read, what to redraw after a build is
--- used (Options.refreshSurfaces), and what to do instead when the dialog cannot show
--- itself (Options.showHint).
BuildDialog.data = nil
BuildDialog.afterUse = nil
BuildDialog.fallback = nil

--- Item ids whose names the preview is waiting on (GET_ITEM_INFO_RECEIVED redraws it).
BuildDialog.pending = {}

-- Pure rules ---------------------------------------------------------------------------

--- The paste box's verdict on `text`: { ok = true, build } when the codec loads it, else
--- { ok = false, reason } with the codec's own refusal (already pipe-escaped by Codec).
--- An empty box has no reason: there is nothing to refuse yet.
function BuildDialog.evaluate(text, data)
	if type(text) ~= "string" or not text:find("%S") then
		return { ok = false }
	end
	local build, reason = Codec.loadBuild(text, data)
	if build == nil then
		return { ok = false, reason = reason }
	end
	return { ok = true, build = build }
end

--- The build's trees by the points its order spends in each (Progress.planSplit), biggest
--- first: "Arms 31 / Fury 20". Empty for a build with no points.
function BuildDialog.split(build, data)
	local counts = Progress.planSplit(build)
	local tabs = {}
	for tab, count in pairs(counts) do
		if count > 0 then
			tabs[#tabs + 1] = tab
		end
	end
	table.sort(tabs, function(left, right)
		if counts[left] ~= counts[right] then
			return counts[left] > counts[right]
		end
		return left < right
	end)
	local parts = {}
	for _, tab in ipairs(tabs) do
		local name = Talents.tabName(data, build.classSlug, tab)
		if name ~= nil then
			parts[#parts + 1] = string.format(L.dialogSplitPart, name, counts[tab])
		end
	end
	return table.concat(parts, L.dialogSplitSeparator)
end

--- What the preview shows for a build the codec loaded: the class token (for
--- Theme.classIcon), the name (Follow.buildName: the build's own, else its trees), the
--- line "Warrior · Arms 31 / Fury 20 · 51 points", and up to three planned item ids.
function BuildDialog.preview(build, data)
	local className = Talents.titleCase(build.classSlug)
	local points = #build.order
	local split = BuildDialog.split(build, data)
	local detail
	if points == 0 then
		detail = string.format(L.dialogPreviewNoPoints, className)
	elseif points == 1 then
		detail = string.format(L.dialogPreviewOne, className, split)
	else
		detail = string.format(L.dialogPreview, className, split, points)
	end
	local picks = {}
	for _, entry in ipairs(build.gear or {}) do
		if #picks >= BuildDialog.GEAR_PICKS then
			break
		end
		if type(entry.itemId) == "number" then
			picks[#picks + 1] = entry.itemId
		end
	end
	return {
		classToken = build.classSlug:upper(),
		name = Follow.buildName(data, build),
		detail = detail,
		picks = picks,
	}
end

--- Seconds since 1970-01-01 00:00 UTC for a date and time read as UTC (the proleptic
--- Gregorian day count), so no client clock call is needed to compare two times.
function BuildDialog.civilSeconds(year, month, day, hour, minute, second)
	local y = month <= 2 and year - 1 or year
	local era = math.floor(y / 400)
	local yearOfEra = y - era * 400
	local shifted = (month + 9) % 12
	local dayOfYear = math.floor((153 * shifted + 2) / 5) + day - 1
	local dayOfEra = yearOfEra * 365 + math.floor(yearOfEra / 4) - math.floor(yearOfEra / 100) + dayOfYear
	local days = era * 146097 + dayOfEra - 719468
	return days * 86400 + (hour or 0) * 3600 + (minute or 0) * 60 + (second or 0)
end

local function validFields(month, day, hour, minute, second)
	return month >= 1 and month <= 12 and day >= 1 and day <= 31 and hour <= 23 and minute <= 59
		and (second or 0) <= 60
end

--- An RFC 3339 time (the ack's `export_seen_at`: "2026-11-04T14:03:00Z", with an
--- offset or fractional seconds too) as seconds since the epoch, or nil.
function BuildDialog.parseUtc(text)
	if type(text) ~= "string" then
		return nil
	end
	local year, month, day, hour, minute, second, zone =
		text:match("^(%d%d%d%d)%-(%d%d)%-(%d%d)[Tt](%d%d):(%d%d):(%d%d)%.?%d*(.*)$")
	if year == nil then
		return nil
	end
	year, month, day = tonumber(year), tonumber(month), tonumber(day)
	hour, minute, second = tonumber(hour), tonumber(minute), tonumber(second)
	if not validFields(month, day, hour, minute, second) then
		return nil
	end
	local offset
	if zone == "Z" or zone == "z" then
		offset = 0
	else
		local sign, zoneHour, zoneMinute = zone:match("^([+-])(%d%d):?(%d%d)$")
		if sign == nil then
			return nil
		end
		offset = (tonumber(zoneHour) * 3600 + tonumber(zoneMinute) * 60) * (sign == "-" and -1 or 1)
	end
	return BuildDialog.civilSeconds(year, month, day, hour, minute, second) - offset
end

--- The outbox's `saved_at` ("2026-11-04 14:02", the player's local wall time) as seconds
--- since the epoch, given the client's offset from UTC in seconds; nil for anything else.
function BuildDialog.parseLocal(text, offset)
	if type(text) ~= "string" then
		return nil
	end
	local year, month, day, hour, minute = text:match("^(%d%d%d%d)%-(%d%d)%-(%d%d) (%d%d):(%d%d)$")
	if year == nil then
		return nil
	end
	year, month, day = tonumber(year), tonumber(month), tonumber(day)
	hour, minute = tonumber(hour), tonumber(minute)
	if not validFields(month, day, hour, minute, 0) then
		return nil
	end
	return BuildDialog.civilSeconds(year, month, day, hour, minute, 0) - (offset or 0)
end

--- A moment as the local "14:02".
function BuildDialog.clock(seconds, offset)
	local ofDay = (seconds + (offset or 0)) % 86400
	return string.format("%02d:%02d", math.floor(ofDay / 3600), math.floor(ofDay % 3600 / 60))
end

--- "Copied 14:02, paste it on the site", at `seconds` in local time.
function BuildDialog.copiedLine(seconds, offset)
	return string.format(L.dialogCopied, BuildDialog.clock(seconds, offset))
end

--- "Seen by the site 14:03" (local time) when the ack's `export_seen_at` is newer than
--- the export this computer last saved (`savedAt`, local wall time); nil when the save is
--- newer, or the ack is missing or unreadable. With no save to compare, the ack stands:
--- it is the site's own word that it has the character.
function BuildDialog.seenLine(ack, savedAt, offset)
	if type(ack) ~= "table" then
		return nil
	end
	local seen = BuildDialog.parseUtc(ack.export_seen_at)
	if seen == nil then
		return nil
	end
	local saved = BuildDialog.parseLocal(savedAt, offset)
	if saved ~= nil and seen <= saved then
		return nil
	end
	return string.format(L.dialogSeen, BuildDialog.clock(seen, offset))
end

--- The line under the copy button: { text, hexKey } or nil. The copied line in body
--- colour for COPIED_SECONDS after a copy, else the seen line muted, else nothing.
--- `input` = { now, copiedAt, offset, ack, savedAt }, times in seconds since the epoch.
function BuildDialog.statusLine(input)
	local copiedAt = input.copiedAt
	if copiedAt ~= nil and input.now - copiedAt < BuildDialog.COPIED_SECONDS then
		return { text = BuildDialog.copiedLine(copiedAt, input.offset), hexKey = "body" }
	end
	local seen = BuildDialog.seenLine(input.ack, input.savedAt, input.offset)
	if seen ~= nil then
		return { text = seen, hexKey = "muted" }
	end
	return nil
end

--- The ack's `sim_request_state` in words; nothing for a state this addon does not know.
function BuildDialog.simStateLine(state)
	if type(state) ~= "string" then
		return nil
	end
	return L.dialogSimState[state]
end

--- The outbox's `saved_at` when the outbox (contract 4, last session's) is this
--- character's (`outboxName`, Export.outboxCharacter), else nil.
function BuildDialog.savedAtFor(db, outboxName)
	local outbox = type(db) == "table" and db.outbox or nil
	local export = type(outbox) == "table" and outbox.export or nil
	if type(export) ~= "table" or export.character ~= outboxName or type(export.saved_at) ~= "string" then
		return nil
	end
	return export.saved_at
end

--- The "from the site" rows: one per build in `builds` (Follow.inboxBuilds) whose id is not
--- in `dismissed` and whose code loads, as { id, name, code, label }; a build with no name
--- is named by Follow.buildName. The second value is the first refusal met, for /fs inbox.
function BuildDialog.inboxRows(builds, dismissed, data)
	local rows, refusal = {}, nil
	dismissed = dismissed or {}
	for _, entry in ipairs(builds or {}) do
		if not dismissed[entry.id] then
			local verdict = BuildDialog.evaluate(entry.code, data)
			if verdict.ok then
				local name = type(entry.name) == "string" and entry.name ~= "" and entry.name
					or Follow.buildName(data, verdict.build)
				rows[#rows + 1] = {
					id = entry.id,
					name = name,
					code = entry.code,
					label = string.format(L.dialogInboxRow, (name:gsub("|", "||"))),
				}
			elseif refusal == nil then
				refusal = verdict.reason
			end
		end
	end
	return rows, refusal
end

function BuildDialog.modeOf(mode)
	return mode == "copy" and "copy" or "paste"
end

function BuildDialog.nextMode(mode)
	return BuildDialog.modeOf(mode) == "paste" and "copy" or "paste"
end

--- Whether a rect { left, right, top, bottom } lies wholly inside `parent`'s.
function BuildDialog.onScreen(rect, parent)
	for _, edge in ipairs({ "left", "right", "top", "bottom" }) do
		if type(rect[edge]) ~= "number" or type(parent[edge]) ~= "number" then
			return false
		end
	end
	return rect.left >= parent.left and rect.right <= parent.right
		and rect.top <= parent.top and rect.bottom >= parent.bottom
end

-- The live inputs ----------------------------------------------------------------------

local function dataOf()
	return BuildDialog.data or ns.Data
end

--- The wall clock, seconds since the epoch; nil on a client without one.
local function wallNow()
	if type(time) == "function" then
		return time()
	end
	return nil
end

--- The client's offset from UTC in seconds (local wall time minus UTC), 0 where the
--- client cannot say.
function BuildDialog.utcOffset()
	local now = wallNow()
	if now == nil or type(date) ~= "function" then
		return 0
	end
	local fields = date("*t", now)
	if type(fields) ~= "table" then
		return 0
	end
	return BuildDialog.civilSeconds(fields.year, fields.month, fields.day, fields.hour, fields.min, fields.sec) - now
end

local function currentAck()
	return Follow.messages(ForeverSixtyInbox, Export.characterKey(), "ack")[1]
end

local function currentRows()
	return BuildDialog.inboxRows(Follow.inboxBuilds(ForeverSixtyInbox, Export.characterKey()),
		Follow.dismissedSet(), dataOf())
end

-- The frame ----------------------------------------------------------------------------

local function contentWidth()
	return Theme.SIZES.dialogWidth - Theme.SIZES.padding * 2
end

--- Where the dialog sits: the saved spot after a drag, else the middle of the screen.
function BuildDialog.place()
	local frame = BuildDialog.frame
	if frame == nil then
		return nil
	end
	frame:ClearAllPoints()
	local point = Prefs.get("dialog", "point")
	if type(point) == "string" and point ~= "" then
		frame:SetPoint(point, UIParent, point, Prefs.get("dialog", "x"), Prefs.get("dialog", "y"))
		return point
	end
	frame:SetPoint("CENTER", UIParent, "CENTER", 0, 0)
	return "CENTER"
end

--- After a drag by the title: the dialog's centre against the screen's.
function BuildDialog.savePosition()
	local frame = BuildDialog.frame
	local x, y = frame:GetCenter()
	local parentX, parentY = UIParent:GetCenter()
	if x == nil or y == nil or parentX == nil or parentY == nil then
		return nil
	end
	Prefs.set("dialog", "point", "CENTER")
	Prefs.set("dialog", "x", x - parentX)
	Prefs.set("dialog", "y", y - parentY)
	BuildDialog.place()
	return true
end

--- Escape, Enter and Tab in either box: Escape closes the dialog (the box gives up the
--- keyboard as it hides, so the player can walk), Tab switches tabs, Enter does the tab's job.
local function bindKeys(box, onEnter)
	box:SetScript("OnEscapePressed", function()
		BuildDialog.close()
	end)
	box:SetScript("OnTabPressed", function()
		BuildDialog.setMode(BuildDialog.nextMode(BuildDialog.mode))
	end)
	box:SetScript("OnEnterPressed", onEnter)
end

local function buildTitle(frame)
	local S = Theme.SIZES
	local bar = Theme.createFrame("Frame", nil, frame)
	bar:SetPoint("TOPLEFT", frame, "TOPLEFT", 0, 0)
	bar:SetPoint("TOPRIGHT", frame, "TOPRIGHT", 0, 0)
	bar:SetHeight(S.titleBarHeight)
	bar:EnableMouse(true)
	bar:RegisterForDrag("LeftButton")
	bar:SetScript("OnDragStart", function()
		frame:StartMoving()
	end)
	bar:SetScript("OnDragStop", function()
		frame:StopMovingOrSizing()
		BuildDialog.savePosition()
	end)
	local title = Widgets.label(bar, L.addonName, "gold", "normal")
	title:SetPoint("LEFT", bar, "LEFT", S.padding, 0)
	local close = Widgets.closeButton(bar, function()
		BuildDialog.close()
	end)
	close:SetPoint("RIGHT", bar, "RIGHT", -S.gap * 2, 0)
	return bar
end

local function buildTab(frame, text, mode)
	local S = Theme.SIZES
	local tab = Widgets.tab(frame, text, function()
		BuildDialog.setMode(mode)
	end)
	local label = tab.foreverSixtyLabel
	local width = type(label.GetStringWidth) == "function" and label:GetStringWidth() or 0
	tab:SetWidth(math.max(S.tabWidth, width + S.padding))
	return tab
end

--- One planned item in the preview: its icon and name; hovering shows the item.
local function buildPickRow(parent)
	local S = Theme.SIZES
	local row = Theme.createFrame("Frame", nil, parent)
	row:SetSize(contentWidth() - S.dialogClassIcon - S.gap * 2, S.rowHeight)
	row:EnableMouse(true)
	local icon = Theme.icon(row, "ARTWORK", nil, S.iconSize)
	icon:SetPoint("LEFT", row, "LEFT", 0, 0)
	local name = Widgets.label(row, "", "body", "small")
	name:SetPoint("LEFT", icon, "RIGHT", S.gap * 2, 0)
	name:SetWidth(row:GetWidth() - S.iconSize - S.gap * 2)
	if type(name.SetWordWrap) == "function" then
		name:SetWordWrap(false)
	end
	local pick = { frame = row, icon = icon, name = name }
	Widgets.attachTooltip(row, function()
		return pick.itemId, nil
	end)
	return pick
end

local function buildPreview(content)
	local S = Theme.SIZES
	local preview = Theme.createFrame("Frame", nil, content)
	preview:SetSize(contentWidth(), S.dialogClassIcon + S.rowHeight * BuildDialog.GEAR_PICKS + S.gap * 2)
	local icon = Theme.icon(preview, "ARTWORK", nil, S.dialogClassIcon)
	icon:SetPoint("TOPLEFT", preview, "TOPLEFT", 0, 0)
	local name = Widgets.label(preview, "", "gold", "normal")
	name:SetPoint("TOPLEFT", icon, "TOPRIGHT", S.gap * 2, 0)
	name:SetWidth(contentWidth() - S.dialogClassIcon - S.gap * 2)
	local detail = Widgets.label(preview, "", "muted", "small")
	detail:SetPoint("TOPLEFT", name, "BOTTOMLEFT", 0, -S.gap)
	detail:SetWidth(contentWidth() - S.dialogClassIcon - S.gap * 2)
	local picks = {}
	for index = 1, BuildDialog.GEAR_PICKS do
		local pick = buildPickRow(preview)
		pick.frame:SetPoint("TOPLEFT", preview, "TOPLEFT", S.dialogClassIcon + S.gap * 2,
			-(S.dialogClassIcon + S.gap * 2 + (index - 1) * S.rowHeight))
		picks[index] = pick
	end
	preview:Hide()
	return { frame = preview, icon = icon, name = name, detail = detail, picks = picks }
end

local function buildInboxRow(content)
	local S = Theme.SIZES
	local row = { frame = Theme.createFrame("Frame", nil, content) }
	row.frame:SetSize(contentWidth(), S.buttonHeight)
	row.label = Widgets.label(row.frame, "", "body", "small")
	row.label:SetPoint("LEFT", row.frame, "LEFT", 0, 0)
	row.label:SetWidth(contentWidth() - S.dialogUseWidth - S.gap * 2)
	if type(row.label.SetWordWrap) == "function" then
		row.label:SetWordWrap(false)
	end
	row.button = Widgets.button(row.frame, L.dialogInboxUse, function()
		if row.entry ~= nil then
			BuildDialog.use(row.entry.code, row.entry.name, row.entry.id)
		end
	end)
	row.button:SetSize(S.dialogUseWidth, S.buttonHeight)
	row.button:SetPoint("RIGHT", row.frame, "RIGHT", 0, 0)
	row.frame:Hide()
	return row
end

local function buildPaste(frame, top)
	local S = Theme.SIZES
	local content = Theme.createFrame("Frame", nil, frame)
	content:SetPoint("TOPLEFT", frame, "TOPLEFT", S.padding, -top)
	content:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", -S.padding, S.padding)
	local paste = { content = content, rows = {} }
	for index = 1, S.dialogInboxRows do
		paste.rows[index] = buildInboxRow(content)
	end
	local box = Widgets.editBox(content, contentWidth(), S.dialogBox, false, false)
	-- Ctrl+C here is not a copy out of the dialog; the box keeps the keyboard until Escape.
	box:SetScript("OnKeyDown", nil)
	local field = Widgets.field(box)
	local placeholder = Widgets.label(field, L.dialogPaste, "muted", "small")
	placeholder:SetPoint("TOPLEFT", field, "TOPLEFT", S.gap * 2, -S.gap * 2)
	box:SetScript("OnTextChanged", function(self, userInput)
		if userInput then
			BuildDialog.onPaste(self:GetText())
		else
			BuildDialog.drawPaste()
		end
	end)
	bindKeys(box, function()
		BuildDialog.useCurrent()
	end)
	paste.box, paste.field, paste.placeholder = box, field, placeholder
	paste.reason = Widgets.label(content, "", "muted", "small")
	paste.reason:SetWidth(contentWidth())
	paste.reason:Hide()
	paste.preview = buildPreview(content)
	paste.use = Cards.primaryButton(content, L.dialogUse, function()
		BuildDialog.useCurrent()
	end)
	paste.use:SetPoint("BOTTOMLEFT", content, "BOTTOMLEFT", 0, 0)
	paste.use:Hide()
	return paste
end

local function buildCopy(frame, top)
	local S = Theme.SIZES
	local content = Theme.createFrame("Frame", nil, frame)
	content:SetPoint("TOPLEFT", frame, "TOPLEFT", S.padding, -top)
	content:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", -S.padding, S.padding)
	local copy = { content = content }
	local box = Widgets.editBox(content, contentWidth(), S.dialogLine, true, true)
	bindKeys(box, function()
		BuildDialog.copy()
	end)
	local field = Widgets.field(box)
	field:SetPoint("TOPLEFT", content, "TOPLEFT", 0, 0)
	copy.box = box
	copy.keys = Widgets.label(content, "", "muted", "small")
	copy.keys:SetWidth(contentWidth())
	copy.keys:SetPoint("TOPLEFT", field, "BOTTOMLEFT", 0, -S.gap * 2)
	copy.button = Cards.primaryButton(content, L.dialogCopy, function()
		BuildDialog.copy()
	end)
	copy.button:SetPoint("TOPLEFT", copy.keys, "BOTTOMLEFT", 0, -S.gap * 3)
	copy.status = Widgets.label(content, "", "body", "small")
	copy.status:SetWidth(contentWidth())
	copy.status:SetPoint("TOPLEFT", copy.button, "BOTTOMLEFT", 0, -S.gap * 3)
	copy.sim = Widgets.label(content, "", "muted", "small")
	copy.sim:SetWidth(contentWidth())
	copy.sim:SetPoint("TOPLEFT", copy.status, "BOTTOMLEFT", 0, -S.gap)
	copy.companion = Widgets.label(content, L.dialogCompanionLine, "muted", "small")
	copy.companion:SetWidth(contentWidth())
	copy.companion:SetPoint("BOTTOMLEFT", content, "BOTTOMLEFT", 0, 0)
	content:Hide()
	return copy
end

local function onEvent(_, event, itemId)
	if event == "GET_ITEM_INFO_RECEIVED" and itemId ~= nil and BuildDialog.pending[itemId] then
		BuildDialog.pending[itemId] = nil
		if BuildDialog.frame:IsShown() then
			BuildDialog.drawPaste()
		end
	end
end

--- The dialog, built on the first open: a 420x300 panel titled "Forever Sixty",
--- draggable by its title, closed by Escape, with the two tabs under the title.
function BuildDialog.ensure()
	if BuildDialog.frame ~= nil then
		return BuildDialog.frame
	end
	local S = Theme.SIZES
	local frame = Widgets.panel(UIParent, S.dialogWidth, S.dialogHeight, BuildDialog.FRAME_NAME)
	frame:Hide()
	frame:SetFrameStrata("DIALOG")
	frame:SetClampedToScreen(true)
	frame:SetMovable(true)
	frame:EnableMouse(true)
	buildTitle(frame)
	local tabs = {
		paste = buildTab(frame, L.dialogTabPaste, "paste"),
		copy = buildTab(frame, L.dialogCopy, "copy"),
	}
	tabs.paste:SetPoint("TOPLEFT", frame, "TOPLEFT", S.padding, -S.titleBarHeight)
	tabs.copy:SetPoint("LEFT", tabs.paste, "RIGHT", S.gap * 2, 0)
	BuildDialog.tabs = tabs
	local top = S.titleBarHeight + S.tabHeight + S.gap * 3
	BuildDialog.paste = buildPaste(frame, top)
	BuildDialog.copyTab = buildCopy(frame, top)
	frame:SetScript("OnHide", function()
		-- Never keep the keyboard once the dialog is gone.
		BuildDialog.paste.box:ClearFocus()
		BuildDialog.copyTab.box:ClearFocus()
		BuildDialog.copiedAt = nil
	end)
	frame:SetScript("OnEvent", onEvent)
	Theme.registerEvent(frame, "GET_ITEM_INFO_RECEIVED")
	Theme.makeEscapable(BuildDialog.FRAME_NAME)
	BuildDialog.frame = frame
	return frame
end

-- Paste mode ---------------------------------------------------------------------------

--- Lay the paste tab out top to bottom: the "from the site" rows, the box (one line once
--- a code loads), then the reason or the preview; the Use button sits at the bottom.
local function layoutPaste(collapsed)
	local S = Theme.SIZES
	local paste = BuildDialog.paste
	local y = 0
	for _, row in ipairs(paste.rows) do
		if row.frame:IsShown() then
			row.frame:ClearAllPoints()
			row.frame:SetPoint("TOPLEFT", paste.content, "TOPLEFT", 0, -y)
			y = y + S.buttonHeight + S.gap
		end
	end
	local height = collapsed and S.dialogLine or S.dialogBox
	paste.field:ClearAllPoints()
	paste.field:SetPoint("TOPLEFT", paste.content, "TOPLEFT", 0, -y)
	paste.field:SetHeight(height)
	paste.box:SetHeight(height - S.gap * 2)
	y = y + height + S.gap * 2
	paste.reason:ClearAllPoints()
	paste.reason:SetPoint("TOPLEFT", paste.content, "TOPLEFT", 0, -y)
	paste.preview.frame:ClearAllPoints()
	paste.preview.frame:SetPoint("TOPLEFT", paste.content, "TOPLEFT", 0, -y)
end

local function drawInbox()
	local rows = currentRows()
	for index, row in ipairs(BuildDialog.paste.rows) do
		local entry = rows[index]
		row.entry = entry
		if entry == nil then
			row.frame:Hide()
		else
			row.label:SetText(entry.label)
			row.frame:Show()
		end
	end
end

local function drawPreview(build)
	local preview = BuildDialog.paste.preview
	local model = BuildDialog.preview(build, dataOf())
	local path, coords = Theme.classIcon(model.classToken)
	preview.icon:SetTexture(path)
	if type(preview.icon.SetTexCoord) == "function" then
		preview.icon:SetTexCoord(coords[1], coords[2], coords[3], coords[4])
	end
	preview.name:SetText(model.name)
	preview.detail:SetText(model.detail)
	for index, pick in ipairs(preview.picks) do
		local itemId = model.picks[index]
		pick.itemId = itemId
		if itemId == nil then
			pick.frame:Hide()
		else
			pick.icon:SetTexture(Compat.itemIcon(itemId) or Theme.UNKNOWN_ICON)
			local name = Compat.itemName(itemId)
			if name == nil then
				BuildDialog.pending[itemId] = true
			end
			pick.name:SetText(name or "")
			pick.frame:Show()
		end
	end
	BuildDialog.previewName = model.name
	preview.frame:Show()
end

--- Redraw the paste tab from the last verdict.
function BuildDialog.drawPaste()
	local paste = BuildDialog.paste
	if paste == nil then
		return nil
	end
	local text = paste.box:GetText() or ""
	if text == "" then
		paste.placeholder:Show()
	else
		paste.placeholder:Hide()
	end
	local verdict = BuildDialog.verdict or { ok = false }
	if verdict.ok then
		paste.reason:Hide()
		drawPreview(verdict.build)
		paste.use:Show()
	else
		paste.preview.frame:Hide()
		paste.use:Hide()
		BuildDialog.previewName = nil
		if verdict.reason ~= nil then
			paste.reason:SetText(verdict.reason)
			paste.reason:Show()
		else
			paste.reason:Hide()
		end
	end
	layoutPaste(verdict.ok)
	return verdict
end

--- The player typed or pasted: judge the text and redraw. A refusal touches nothing
--- else: the build the player follows, the chip and the strip stay as they were.
function BuildDialog.onPaste(text)
	BuildDialog.verdict = BuildDialog.evaluate(text, dataOf())
	return BuildDialog.drawPaste()
end

--- Follow `code` (Follow.load saves it in the active slot), mark an inbox build used,
--- redraw every surface, close, and say so in one chat line.
function BuildDialog.use(code, name, inboxId)
	local build = Follow.load(code, dataOf(), name)
	if build == nil then
		return nil
	end
	if inboxId ~= nil then
		Follow.dismissInbox(inboxId)
	end
	if BuildDialog.afterUse ~= nil then
		BuildDialog.afterUse()
	end
	BuildDialog.close()
	Theme.say(string.format(L.dialogFollowing, Follow.buildName(dataOf(), build)))
	return build
end

--- The Use button, and Enter in the paste box: only while a preview shows.
function BuildDialog.useCurrent()
	local verdict = BuildDialog.verdict
	if verdict == nil or not verdict.ok or BuildDialog.paste == nil then
		return nil
	end
	return BuildDialog.use(BuildDialog.paste.box:GetText(), BuildDialog.previewName)
end

-- Copy mode ----------------------------------------------------------------------------

local function drawStatus()
	local copy = BuildDialog.copyTab
	if copy == nil then
		return nil
	end
	local ack = currentAck()
	local line = BuildDialog.statusLine({
		now = wallNow() or 0,
		copiedAt = BuildDialog.copiedAt,
		offset = BuildDialog.utcOffset(),
		ack = ack,
		savedAt = BuildDialog.savedAtFor(ForeverSixtyDB, Export.outboxCharacter()),
	})
	if line == nil then
		copy.status:Hide()
	else
		copy.status:SetText(line.text)
		copy.status:SetTextColor(Theme.rgb(Theme.HEX[line.hexKey]))
		copy.status:Show()
	end
	local sim = type(ack) == "table" and BuildDialog.simStateLine(ack.sim_request_state) or nil
	if sim == nil then
		copy.sim:Hide()
	else
		copy.sim:SetText(sim)
		copy.sim:Show()
	end
	return line
end

--- Fill the copy tab: the export selected and focused so Ctrl+C copies it, or the
--- reason there is none (then the button is off).
function BuildDialog.drawCopy()
	local copy = BuildDialog.copyTab
	-- Guarded: a client answer the export cannot encode must never leave a half-drawn tab.
	local ran, code, reason = pcall(Export.string, dataOf())
	if not ran then
		code, reason = nil, nil
	end
	BuildDialog.exportCode = code
	if code ~= nil then
		Widgets.selectText(copy.box, code)
		copy.keys:SetText(L.dialogCopyKeys)
		Widgets.setEnabled(copy.button, true)
	else
		Widgets.setText(copy.box, "")
		copy.keys:SetText(reason or "")
		Widgets.setEnabled(copy.button, false)
	end
	drawStatus()
	return code
end

--- "Copy my character": select the export again and say "Copied 14:02" for three
--- seconds. A client that cannot wait shows the settled line at once instead.
function BuildDialog.copy()
	local code = BuildDialog.exportCode
	if code == nil or BuildDialog.copyTab == nil then
		return nil
	end
	Widgets.selectText(BuildDialog.copyTab.box, code)
	BuildDialog.copiedAt = wallNow()
	BuildDialog.copyGeneration = (BuildDialog.copyGeneration or 0) + 1
	local generation = BuildDialog.copyGeneration
	local waits = Theme.after(BuildDialog.COPIED_SECONDS, function()
		if generation == BuildDialog.copyGeneration then
			drawStatus()
		end
	end)
	if not waits then
		BuildDialog.copiedAt = nil
	end
	drawStatus()
	return code
end

-- Opening and closing ------------------------------------------------------------------

--- Show one tab: the paste box takes the keyboard in paste mode; in copy mode the export
--- is selected in its box.
function BuildDialog.setMode(mode)
	mode = BuildDialog.modeOf(mode)
	BuildDialog.mode = mode
	if BuildDialog.frame == nil then
		return mode
	end
	local paste, copy = BuildDialog.paste, BuildDialog.copyTab
	Widgets.setTabActive(BuildDialog.tabs.paste, mode == "paste")
	Widgets.setTabActive(BuildDialog.tabs.copy, mode == "copy")
	if mode == "paste" then
		copy.box:ClearFocus()
		copy.content:Hide()
		paste.content:Show()
		BuildDialog.drawPaste()
		paste.box:SetFocus()
	else
		paste.box:ClearFocus()
		paste.content:Hide()
		copy.content:Show()
		BuildDialog.drawCopy()
	end
	return mode
end

--- The box of the tab showing, which holds the keyboard while the dialog is open.
local function focusedBox()
	if BuildDialog.mode == "copy" then
		return BuildDialog.copyTab.box
	end
	return BuildDialog.paste.box
end

--- Open the dialog on `mode` ("paste" by default, or "copy"). The paste box starts
--- empty. The box is focused again on the next frame, since the chat box that ran /fs
--- lets go of the keyboard after the command, and then the render check runs. False
--- when the dialog turned itself off earlier (the fallback ran instead).
function BuildDialog.open(mode)
	if BuildDialog.disabled then
		if BuildDialog.fallback ~= nil then
			BuildDialog.fallback()
		end
		return false
	end
	local frame = BuildDialog.ensure()
	BuildDialog.place()
	BuildDialog.verdict = nil
	BuildDialog.paste.box:SetText("")
	drawInbox()
	frame:Show()
	BuildDialog.setMode(mode)
	local deferred = Theme.after(0, function()
		if frame:IsShown() then
			focusedBox():SetFocus()
			BuildDialog.selfCheck()
		end
	end)
	if not deferred then
		BuildDialog.selfCheck()
	end
	return true
end

--- Hide the dialog and give the keyboard back (OnHide does it too, for the Escape the
--- client handles itself through UISpecialFrames).
function BuildDialog.close()
	if BuildDialog.frame ~= nil then
		BuildDialog.paste.box:ClearFocus()
		BuildDialog.copyTab.box:ClearFocus()
		BuildDialog.frame:Hide()
	end
end

--- The keybind: open on the paste tab, or close.
function BuildDialog.toggle()
	if BuildDialog.frame ~= nil and BuildDialog.frame:IsShown() then
		BuildDialog.close()
		return false
	end
	return BuildDialog.open("paste")
end

local function rectOf(region)
	return { left = region:GetLeft(), right = region:GetRight(), top = region:GetTop(), bottom = region:GetBottom() }
end

--- The render check (spec section 2.4): the dialog lies on screen and, in paste mode,
--- the box has the keyboard. Otherwise it turns itself off for the session, says why
--- in /fs diag and the fallback (the command list) runs. True when it passes.
function BuildDialog.selfCheck()
	local frame = BuildDialog.frame
	if frame == nil or not frame:IsShown() then
		return true
	end
	local ok = BuildDialog.onScreen(rectOf(frame), rectOf(UIParent))
	local box = BuildDialog.paste.box
	if ok and BuildDialog.mode == "paste" and type(box.HasFocus) == "function" then
		ok = box:HasFocus() == true
	end
	if ok then
		return true
	end
	BuildDialog.disabled = true
	frame:Hide()
	Theme.note(L.diagDialogOff)
	if BuildDialog.fallback ~= nil then
		BuildDialog.fallback()
	end
	return false
end

ns.BuildDialog = BuildDialog
return BuildDialog
