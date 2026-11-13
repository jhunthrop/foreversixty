-- addon/ForeverSixty/surfaces/Settings.lua
-- The Settings page (spec section 4.11, L3.5; rebuilt for the hub, design/specs/2026-11-13-addon-hub.md 7):
-- the client's own Settings, one canvas, in the Forever Sixty window's type: the title in the title face,
-- section headings as the window's (BuildsWindow.buildSection), text at the window's 12 px, no rules.
-- The /fs command list, the switches in three sections ("Tooltips only" first), Reset positions on the
-- window's one button recipe (Widgets.clientButton), and two lines at the foot: whether the addon's data
-- is for this client's patch, and the player ratings' state. The first of those in the warning colour is
-- repeated under the title, its whole sentence in that colour. Switches only: every choice that is not
-- on or off lives in the tab that owns it (the gear sources are the Gear tab's).
--
-- Pure model, thin frame. SettingsPage.SWITCHES is the whole of what the page can change
-- (each key a Prefs.DEFAULTS boolean); read, checked and apply are its only way to the prefs;
-- layout gives every number on the page from the rows' measured heights; addonDataLine and
-- ratingsLine choose the foot lines and warningLine the title's.
-- The canvas is an empty frame registered at login and built the first time the client shows
-- it, never at load.
--
-- The module's own table is SettingsPage, not Settings: `Settings` is the client's global
-- for its options API, and a local of that name in this file would shadow it.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Prefs = ns.Prefs or require("Prefs")
local Compat = ns.Compat or require("Compat")
local Codec = ns.Codec or require("Codec")
local MinimapButton = ns.Minimap or require("Minimap")
local Score = ns.Score or require("Score")

local SettingsPage = {}

SettingsPage.FRAME_NAME = "ForeverSixtySettingsPanel"

--- The switch sections, in page order: what shows in the game, the launcher, what leaves the game.
SettingsPage.SECTIONS = {
	{ id = "world", heading = L.settingsSectionWorld },
	{ id = "line", heading = L.settingsSectionLine },
	{ id = "shared", heading = L.settingsSectionShared },
}

--- The switches, in page order (L3.5 ruling 1: "Tooltips only" first), each in its section. A key is a
--- Prefs section and field ("chip.shown") or a top-level flag ("tooltip").
SettingsPage.SWITCHES = {
	{ key = "tooltipsOnly", section = "world", label = L.settingsTooltipsOnly, hint = L.settingsTooltipsOnlyHint },
	{ key = "tooltip", section = "world", label = L.settingsTooltip, hint = L.settingsTooltipHint },
	{ key = "slotMarks", section = "world", label = L.settingsSlotMarks, hint = L.settingsSlotMarksHint },
	{ key = "talentMarks", section = "world", label = L.settingsTalentMarks, hint = L.settingsTalentMarksHint },
	{ key = "badges", section = "world", label = L.settingsBadges, hint = L.settingsBadgesHint },
	{ key = "chip.shown", section = "line", label = L.settingsChipShown, hint = L.settingsChipShownHint },
	{ key = "chip.combatHide", section = "line", label = L.settingsCombatHide, hint = L.settingsCombatHideHint },
	{ key = "minimap.shown", section = "line", label = L.settingsMinimap, hint = L.settingsMinimapHint },
	{ key = "broadcast", section = "shared", label = L.settingsBroadcast, hint = L.settingsBroadcastHint },
	{ key = "showScore", section = "shared", label = L.settingsShowScore, hint = L.settingsShowScoreHint },
	{ key = "autoSave", section = "shared", label = L.settingsAutoSave, hint = L.settingsAutoSaveHint },
	{ key = "lootRares", section = "shared", label = L.settingsLootRares, hint = L.settingsLootRaresHint },
}

--- Injection points Options fills: the data the foot line reads (its `build`), and what runs after a
--- switch changes (Options redraws the surfaces, and says hello when sharing is turned on).
SettingsPage.data = nil
SettingsPage.afterChange = nil

local function split(key)
	local section, field = key:match("^([^.]+)%.([^.]+)$")
	return section, field
end

--- A switch's default, from Prefs.DEFAULTS itself (never a second copy of it).
function SettingsPage.defaultOf(key)
	local section, field = split(key)
	if section ~= nil then
		local defaults = Prefs.DEFAULTS[section]
		if type(defaults) ~= "table" then
			return nil
		end
		return defaults[field]
	end
	return Prefs.DEFAULTS[key]
end

--- The switch table with each default filled in: { key, section, label, hint, default }.
function SettingsPage.switches()
	local out = {}
	for index, switch in ipairs(SettingsPage.SWITCHES) do
		out[index] = { key = switch.key, section = switch.section, label = switch.label, hint = switch.hint,
			default = SettingsPage.defaultOf(switch.key) }
	end
	return out
end

--- The sections as drawn: { heading, first, last } with the switch indexes each holds, in page order;
--- a section with no switch is left out. Pure.
function SettingsPage.sections()
	local out = {}
	for _, section in ipairs(SettingsPage.SECTIONS) do
		local first, last
		for index, switch in ipairs(SettingsPage.SWITCHES) do
			if switch.section == section.id then
				first, last = first or index, index
			end
		end
		if first ~= nil then
			out[#out + 1] = { id = section.id, heading = section.heading, first = first, last = last }
		end
	end
	return out
end

--- A switch's value now.
function SettingsPage.read(key)
	local section, field = split(key)
	if section ~= nil then
		return Prefs.get(section, field)
	end
	return Prefs.flag(key)
end

--- What a switch shows: its pref.
function SettingsPage.checked(key)
	return SettingsPage.read(key) == true
end

--- Write a switch. The minimap button follows at once.
function SettingsPage.apply(key, value)
	value = value and true or false
	if key == "minimap.shown" then
		MinimapButton.setShown(value)
		return value
	end
	local section, field = split(key)
	if section ~= nil then
		Prefs.set(section, field, value)
	else
		Prefs.setFlag(key, value)
	end
	if key == "tooltipsOnly" and value then
		-- One step to quiet: the chip, the talent marks and the slot marks off; their own switches show
		-- it and can be turned back on one by one (spec 3.3). Turning it off changes nothing else.
		for _, quiet in ipairs(SettingsPage.TOOLTIPS_ONLY_OFF) do
			SettingsPage.apply(quiet, false)
		end
	end
	return value
end

--- The switches "Tooltips only" turns off.
SettingsPage.TOOLTIPS_ONLY_OFF = { "chip.shown", "talentMarks", "slotMarks" }

--- A toggle's click: write it, then let Options redraw.
function SettingsPage.onToggle(key, value)
	SettingsPage.apply(key, value)
	if SettingsPage.afterChange ~= nil then
		SettingsPage.afterChange(key, value)
	end
	return value
end

--- A data line from its one sentence (fix round 3, item 3): { text, label, value, tone },
--- the label the subject up to the first ": " and the value the state after it.
local function dataLine(text, tone)
	local label, value = text:match("^(.-:)%s(.*)$")
	return { text = text, label = label or "", value = value or text, tone = tone }
end

--- The foot's first line: the addon's data is for this client's patch or a later one
--- ("Addon data: current for this patch", muted), or for an older patch ("Addon data: an
--- older patch. Update the addon at foreversixty.gg.", in the warning colour). Only the patch
--- counts ("1.60.1" against the data's first three fields), never the build number, and no
--- number is printed. Nil when either side is unknown. { text, label, value, tone }. Pure.
function SettingsPage.addonDataLine(dataBuild, clientVersion)
	if type(dataBuild) ~= "string" or dataBuild == "" or type(clientVersion) ~= "string" or clientVersion == "" then
		return nil
	end
	local parts = 0
	for _ in clientVersion:gmatch("[^.]+") do
		parts = parts + 1
	end
	local dataPatch, taken = {}, 0
	for field in dataBuild:gmatch("[^.]+") do
		if taken < parts then
			taken = taken + 1
			dataPatch[taken] = field
		end
	end
	if Codec.compareBuilds(table.concat(dataPatch, "."), clientVersion) < 0 then
		return dataLine(L.settingsAddonDataOlder, "warning")
	end
	return dataLine(L.settingsAddonDataCurrent, "muted")
end

--- The foot's second line, "Player ratings (on player tooltips):" (fix rounds 2 and 3): the
--- nightly ratings data (the Forever Sixty Data addon) in the state Score.status finds it:
--- not installed (muted, it is optional), current, too old (stale, unreadable or an older
--- format), or made for a newer addon than this one. { text, label, value, tone }. Pure.
--- The state is Score.ratingsState's, which the guild tab's notice shares (L5.1 fix round 1).
local RATINGS_TONE = { Missing = "muted", Current = "muted", Newer = "warning", Old = "warning", Broken = "warning" }

function SettingsPage.ratingsLine(status)
	local state = Score.ratingsState(type(status) == "table" and status or false)
	local text = L["settingsRatings" .. state]
	-- "Old" is age-based, so it is dated by when the record was made ("made 9 Oct"), in the
	-- chip's own short date; undated when the record does not say.
	local made = state == "Old" and type(status) == "table" and Compat.shortDate(status.generated) or nil
	if made ~= nil then
		text = string.format(L.settingsRatingsOldDated, made)
	end
	local line = dataLine(text, RATINGS_TONE[state])
	line.state = state
	return line
end

--- The line under the title (fix round 3, item 2): the first foot line in the warning
--- colour, the addon data before the player ratings, its whole sentence (the frame draws it
--- in the warning colour); nil otherwise. "Not installed" is muted and stays at the foot:
--- the data addon is optional. Pure.
function SettingsPage.warningLine(addon, ratings)
	for _, entry in ipairs({ addon or false, ratings or false }) do
		if entry and entry.tone == "warning" then
			-- Both out of date: the first sentence says so of the second (L3.7 item 38).
			if entry == addon and ratings and ratings.state == "Old" then
				return entry.text .. L.settingsAlsoOld
			end
			return entry.text
		end
	end
	return nil
end

--- A foot line as drawn on a muted font string: the subject, then the state in its tone.
--- With `short` (the title already carries the whole sentence) a warning state keeps only its first sentence,
--- the state without the action: "too old (made 5 Dec)".
function SettingsPage.lineText(entry, short)
	local value = entry.value
	if short and entry.tone == "warning" then
		value = value:match("^(.-)%.%s") or value
	end
	return entry.label .. " " .. Theme.inlineColor(entry.tone, value)
end

--- How many commands the page lists before "/fs help: all the commands": the first five, so the
--- tour's own command is on the page (owner, 2026-11-13: the tour must be findable from the options).
SettingsPage.COMMANDS_SHOWN = 5

--- The commands the page lists: the first five and the line that sends the rest to /fs help (the
--- whole list is /fs help's, L.settingsCommands), so the page never ends in a half-cut row.
function SettingsPage.commandLines()
	local lines = {}
	for index = 1, SettingsPage.COMMANDS_SHOWN do
		lines[index] = L.settingsCommands[index]
	end
	lines[#lines + 1] = L.settingsCommandsMore
	return lines
end

--- A command line as drawn on a muted font string: its "/fs ...:" token in body, the words
--- after it muted (fix round 1, item 14). Pure.
function SettingsPage.commandText(line)
	local token, words = line:match("^(.-:)%s(.*)$")
	if token == nil then
		return line
	end
	return Theme.inlineColor("body", token) .. " " .. words
end

--- A switch row's height from its label's and hint's measured heights: at least the
--- one-line row, and a wrapped hint's extra lines on top (Widgets.toggleRowHeight, which the
--- toggle itself is sized by). Pure.
SettingsPage.rowHeight = Widgets.toggleRowHeight

--- Every number on the page, from the width of its backdrop, each row's measured height
--- (`rowHeights`, in switch order; Theme.SIZES.settingsRow where one is missing) and whether
--- the title carries the warning line (`warned`: true for one line, or the line's measured
--- height, never less than one line; L3.7 item 39), and the foot lines' measured heights
--- (`dataHeights`, never less than one line each): the left edge, the hint's indent and width
--- (the width minus the 40 px indent and the right padding), the width of Reset's sentence
--- and of a full line, and the y of the title, the warning line (nil without one), the Commands
--- heading and its lines (first, under the title), each switch section's heading and each switch
--- row, Reset, and the two foot lines, and the height it all needs with the foot's pad (the page
--- scrolls when the canvas is shorter). A backdrop the client has not sized yet is laid out at
--- Theme.SIZES.settingsWidth. Pure.
function SettingsPage.layout(width, rowHeights, warned, dataHeights)
	local S = Theme.SIZES
	if type(width) ~= "number" or width <= 0 then
		width = S.settingsWidth
	end
	local layout = {
		x = S.padding,
		hintIndent = Widgets.toggleIndent(),
		hintWidth = width - S.settingsHintInset - S.padding,
		-- Reset's sentence sits right of the button and wraps before the right padding.
		resetHintWidth = width - S.padding * 2 - S.buttonWidth - S.gap * 2,
		lineWidth = width - S.padding * 2,
		rows = {},
		commands = {},
		sections = {},
		data = {},
	}
	local y = S.padding
	layout.title = -y
	y = y + S.settingsTitle
	if warned then
		layout.warning = -y
		y = y + (type(warned) == "number" and math.max(warned, S.settingsLine) or S.settingsLine)
	end
	y = y + S.gap * 2
	-- A.2: the commands lead, right under the title, where the list is short and always on the first
	-- screen; the switches (the long part) follow and the page scrolls through them.
	layout.commandsHeading = -y
	y = y + S.settingsLine + S.gap
	for index = 1, #SettingsPage.commandLines() do
		layout.commands[index] = -y
		y = y + S.settingsLine
	end
	for index, section in ipairs(SettingsPage.sections()) do
		y = y + S.padding
		layout.sections[index] = -y
		y = y + S.settingsLine
		for row = section.first, section.last do
			layout.rows[row] = -y
			local height = type(rowHeights) == "table" and rowHeights[row] or nil
			y = y + ((type(height) == "number" and height > 0) and height or S.settingsRow)
		end
	end
	y = y + S.padding
	layout.reset = -y
	y = y + S.clientButtonHeight + S.padding
	for index = 1, 2 do
		layout.data[index] = -y
		-- A foot line wraps at the page's width (a too-old sentence at 12 px): its measured height, never
		-- under one line.
		local height = type(dataHeights) == "table" and dataHeights[index] or nil
		y = y + ((type(height) == "number" and height > S.settingsLine) and height or S.settingsLine)
	end
	layout.height = y + S.settingsFootPad
	return layout
end

--- The page's own controls: one toggle per switch and the Reset button (labels are regions, not
--- children). The gear sources are the Gear tab's.
function SettingsPage.expectedChildren()
	return #SettingsPage.SWITCHES + 1
end

--- How long Reset reads "Positions reset", disabled, after a click.
SettingsPage.RESET_SECONDS = 2

--- Reset's look: "Reset positions", or "Positions reset" and disabled (the window's muted disabled
--- button, as Sim again's Queued) for SettingsPage.RESET_SECONDS after a click. { label, enabled }.
--- Pure: reads SettingsPage.resetting.
function SettingsPage.resetLook()
	if SettingsPage.resetting then
		return { label = L.settingsResetDone, enabled = false }
	end
	return { label = L.settingsReset, enabled = true }
end

local function drawReset()
	local button = SettingsPage.resetButton
	if button == nil then
		return
	end
	local look = SettingsPage.resetLook()
	Widgets.resetButtonLabel(button, look.label)
	Widgets.setEnabled(button, look.enabled)
end

--- Reset positions: Prefs puts the placements back, and its onReset listeners move the
--- chip, the minimap button and the window at once (Chip.place,
--- MinimapButton.refresh, Hub.place). The button then says so for 2 s; a client
--- that cannot wait says nothing rather than stick on "Positions reset".
function SettingsPage.reset()
	local ui = Prefs.resetPositions()
	SettingsPage.resetting = true
	local waiting = Theme.after(SettingsPage.RESET_SECONDS, function()
		SettingsPage.resetting = false
		drawReset()
	end)
	if not waiting then
		SettingsPage.resetting = false
	end
	drawReset()
	return ui
end

-- The frame -------------------------------------------------------------------

--- A section heading in the window's own recipe (BuildsWindow.buildSection's: the title face, cream, 14 px).
function SettingsPage.heading(parent, text)
	return Theme.titleFont(Widgets.label(parent, text, "bone"), Theme.SIZES.windowSection)
end

local function place(region, parent, x, y)
	region:ClearAllPoints()
	region:SetPoint("TOPLEFT", parent, "TOPLEFT", x, y)
	return region
end

--- Put everything at the layout's numbers for the rows' measured heights and the warning
--- line (`warned`), and size the scroll to match.
local function arrange(warned)
	local content = SettingsPage.content
	local heights = {}
	for index, toggle in ipairs(SettingsPage.toggles) do
		heights[index] = toggle.frame:GetHeight()
	end
	if warned then
		-- The warning sentence is measured (a longer locale wraps), not assumed one line tall.
		local measured = SettingsPage.warningLabel:GetStringHeight()
		warned = type(measured) == "number" and measured > 0 and measured or true
	end
	local dataHeights = {}
	for index, label in ipairs(SettingsPage.dataLabels) do
		local measured = type(label.GetStringHeight) == "function" and label:GetStringHeight() or nil
		dataHeights[index] = type(measured) == "number" and measured or nil
	end
	local layout = SettingsPage.layout(SettingsPage.backdrop:GetWidth(), heights, warned, dataHeights)
	place(SettingsPage.title, content, layout.x, layout.title)
	place(SettingsPage.warningLabel, content, layout.x, layout.warning or layout.title)
	for index, toggle in ipairs(SettingsPage.toggles) do
		place(toggle.frame, content, layout.x, layout.rows[index])
	end
	place(SettingsPage.resetButton, content, layout.x, layout.reset)
	place(SettingsPage.commandsHeading, content, layout.x, layout.commandsHeading)
	for index, heading in ipairs(SettingsPage.sectionHeadings) do
		place(heading, content, layout.x, layout.sections[index])
	end
	for index, label in ipairs(SettingsPage.commandLabels) do
		place(label, content, layout.x, layout.commands[index])
	end
	for index, label in ipairs(SettingsPage.dataLabels) do
		place(label, content, layout.x, layout.data[index])
	end
	SettingsPage.scroll:SetContentHeight(layout.height)
	return layout
end

--- Build the page once: an opaque backdrop 8 px inside the canvas (the dialog's own, so no
--- world text shows through a hint), anchored to the canvas at both corners so it is always
--- exactly what shows (fix round 3, item 1; the size the client gave the canvas at the first
--- show is only its fallback), scrolling when the canvas is shorter than the page, and
--- everything else on its content frame, arranged once each row has measured its hint.
function SettingsPage.build(panel)
	if SettingsPage.toggles ~= nil then
		return panel
	end
	local S = Theme.SIZES
	local width, height = panel:GetWidth(), panel:GetHeight()
	width = (type(width) == "number" and width > 0) and width or S.settingsWidth + S.settingsInset * 2
	height = (type(height) == "number" and height > 0) and height or S.settingsHeight
	local backdrop = Widgets.panel(panel, width - S.settingsInset * 2, height - S.settingsInset * 2)
	backdrop:SetPoint("TOPLEFT", panel, "TOPLEFT", S.settingsInset, -S.settingsInset)
	backdrop:SetPoint("BOTTOMRIGHT", panel, "BOTTOMRIGHT", -S.settingsInset, S.settingsInset)
	SettingsPage.backdrop = backdrop
	local layout = SettingsPage.layout(backdrop:GetWidth())
	local scroll = Widgets.scrollable(backdrop)
	SettingsPage.scroll = scroll
	local content = scroll.content
	SettingsPage.content = content
	SettingsPage.title = Theme.titleFont(Widgets.label(content, L.addonName, "bone"), S.windowTitle)
	-- The first data state that is not current, repeated where the page opens (item 5).
	local warning = Widgets.label(content, "", "warning")
	warning:SetWidth(layout.lineWidth)
	warning:Hide()
	SettingsPage.warningLabel = warning
	SettingsPage.toggles = {}
	for index, switch in ipairs(SettingsPage.switches()) do
		local key = switch.key
		local toggle = Widgets.toggle(content, switch.label, SettingsPage.checked(key), function(value)
			SettingsPage.onToggle(key, value)
		end, switch.hint, layout.hintWidth)
		toggle.key = key
		SettingsPage.toggles[index] = toggle
	end
	local reset = Widgets.fitButton(Widgets.clientButton(content, L.settingsReset, SettingsPage.reset),
		S.buttonWidth)
	SettingsPage.resetButton = reset
	local resetHint = Widgets.label(content, L.settingsResetHint, "muted")
	resetHint:SetPoint("LEFT", reset, "RIGHT", S.gap * 2, 0)
	resetHint:SetWidth(layout.resetHintWidth)
	SettingsPage.commandsHeading = SettingsPage.heading(content, L.settingsCommandsHeading)
	SettingsPage.sectionHeadings = {}
	for index, section in ipairs(SettingsPage.sections()) do
		SettingsPage.sectionHeadings[index] = SettingsPage.heading(content, section.heading)
	end
	SettingsPage.commandLabels = {}
	for index, line in ipairs(SettingsPage.commandLines()) do
		SettingsPage.commandLabels[index] = Widgets.label(content, SettingsPage.commandText(line), "muted")
	end
	SettingsPage.dataLabels = {}
	for index = 1, #layout.data do
		SettingsPage.dataLabels[index] = Widgets.label(content, "", "muted")
		SettingsPage.dataLabels[index]:SetWidth(layout.lineWidth)
	end
	arrange(false)
	return panel
end

--- Show the live prefs on every toggle ("Show the chip" off while a right-click hides it),
--- Reset's look, the two foot lines for this client and, while either is in the warning colour, the
--- line under the title (the page moves down to make room for it).
function SettingsPage.refresh()
	for _, toggle in ipairs(SettingsPage.toggles or {}) do
		toggle:SetChecked(SettingsPage.checked(toggle.key))
	end
	drawReset()
	if SettingsPage.dataLabels ~= nil then
		local data = SettingsPage.data or ns.Data
		local addon = SettingsPage.addonDataLine(type(data) == "table" and data.build or nil, Compat.clientVersion())
		local ratings = SettingsPage.ratingsLine(Score.status())
		local warning = SettingsPage.warningLine(addon, ratings)
		local short = warning ~= nil
		SettingsPage.dataLabels[1]:SetText(addon ~= nil and SettingsPage.lineText(addon, short) or "")
		SettingsPage.dataLabels[2]:SetText(SettingsPage.lineText(ratings, short))
		SettingsPage.warningLabel:SetText(warning or "")
		if warning ~= nil then
			SettingsPage.warningLabel:Show()
		else
			SettingsPage.warningLabel:Hide()
		end
		arrange(warning ~= nil)
	end
end

--- The render check (spec section 2.4), on every show: the page holds every control it
--- was built with. Otherwise one /fs diag note for the session. True when it passes.
function SettingsPage.selfCheck()
	local content = SettingsPage.content
	if content == nil or SettingsPage.toggles == nil then
		return true
	end
	local count = #{ content:GetChildren() }
	if count == SettingsPage.expectedChildren() then
		return true
	end
	if not SettingsPage.drawNoted then
		SettingsPage.drawNoted = true
		Theme.note(string.format(L.diagSettingsDrawn, count, SettingsPage.expectedChildren()))
	end
	return false
end

--- Every show: built once, refreshed, opened at the top whatever scroll it was left at
--- (fix round 2, item 5), and checked.
function SettingsPage.onShow(panel)
	SettingsPage.build(panel)
	SettingsPage.refresh()
	SettingsPage.scroll:ScrollTo(0)
	SettingsPage.selfCheck()
end

--- PLAYER_LOGIN: an empty, hidden canvas in the client's Settings (the client shows it,
--- and OnShow builds it). True when the client took it; registers once.
function SettingsPage.register()
	if SettingsPage.panel ~= nil then
		return SettingsPage.how ~= nil
	end
	local panel = CreateFrame("Frame", SettingsPage.FRAME_NAME)
	Theme.noteFrame("Settings page", "login")
	panel.name = L.addonName
	panel:Hide()
	panel:SetScript("OnShow", SettingsPage.onShow)
	-- The client's Settings calls these on a canvas where it finds them: Defaults and
	-- Okay change nothing here (each switch saves as it is clicked), a refresh redraws.
	panel.OnCommit = function() end
	panel.OnDefault = function() end
	panel.OnRefresh = SettingsPage.refresh
	SettingsPage.panel = panel
	SettingsPage.how, SettingsPage.handle = Theme.registerSettingsPanel(panel, L.addonName)
	return SettingsPage.how ~= nil
end

--- Open the client's Settings on this page (the launcher's right-click, /fs options). False
--- when it cannot: no API, not registered, or the client took the call and showed nothing
--- (Theme.openSettingsPanel); a registered page that did not open is noted once.
function SettingsPage.open()
	if SettingsPage.how == nil then
		return false
	end
	if Theme.openSettingsPanel(SettingsPage.how, SettingsPage.handle, SettingsPage.panel) then
		return true
	end
	if not SettingsPage.openNoted then
		SettingsPage.openNoted = true
		Theme.note(L.diagSettingsNoOpen)
	end
	return false
end

ns.SettingsPage = SettingsPage
return SettingsPage
