-- addon/ForeverSixty/surfaces/Settings.lua
-- The Settings page (spec section 4.11, L3.5): the client's own Settings, one canvas.
-- Eight switches, Reset positions, the /fs command list, and two lines at the foot: whether
-- the addon's data is for this client's patch, and the player ratings' state. The first of
-- those in the warning colour is repeated under the title, its whole sentence in that colour.
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
local Chip = ns.Chip or require("Chip")
local Score = ns.Score or require("Score")

local SettingsPage = {}

SettingsPage.FRAME_NAME = "ForeverSixtySettingsPanel"

--- The switches, in page order (L3.5 ruling 1, fix round 1 item 4; L4.4 adds the chat badges
--- after the tooltips). A key is a Prefs section and field ("chip.shown") or a top-level flag
--- ("tooltip").
SettingsPage.SWITCHES = {
	{ key = "chip.shown", label = L.settingsChipShown, hint = L.settingsChipShownHint },
	{ key = "chip.combatHide", label = L.settingsCombatHide, hint = L.settingsCombatHideHint },
	{ key = "talentMarks", label = L.settingsTalentMarks, hint = L.settingsTalentMarksHint },
	{ key = "minimap.shown", label = L.settingsMinimap, hint = L.settingsMinimapHint },
	{ key = "tooltip", label = L.settingsTooltip, hint = L.settingsTooltipHint },
	{ key = "badges", label = L.settingsBadges, hint = L.settingsBadgesHint },
	{ key = "broadcast", label = L.settingsBroadcast, hint = L.settingsBroadcastHint },
	{ key = "autoSave", label = L.settingsAutoSave, hint = L.settingsAutoSaveHint },
	{ key = "lootRares", label = L.settingsLootRares, hint = L.settingsLootRaresHint },
}

--- Injection points Options fills: the data the foot line reads (its `build`), and what
--- runs after a switch changes (Options redraws the surfaces, and says hello when sharing
--- is turned on).
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

--- The switch table with each default filled in: { key, label, hint, default }.
function SettingsPage.switches()
	local out = {}
	for index, switch in ipairs(SettingsPage.SWITCHES) do
		out[index] = { key = switch.key, label = switch.label, hint = switch.hint,
			default = SettingsPage.defaultOf(switch.key) }
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

--- What a switch shows: its pref, except that "Show the chip" reads off while a right-click
--- keeps the chip hidden, so the page never says on while the chip is away (fix 1, item 20).
function SettingsPage.checked(key)
	local value = SettingsPage.read(key)
	if key == "chip.shown" then
		return value == true and not Chip.isHidden()
	end
	return value == true
end

--- Write a switch. "Show the chip" also clears a right-click hide either way: off is off,
--- and on means the chip shows (as /fs chip does). The minimap button follows at once.
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
	if key == "chip.shown" then
		Chip.clearHidden()
	end
	return value
end

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
function SettingsPage.lineText(entry)
	return entry.label .. " " .. Theme.inlineColor(entry.tone, entry.value)
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
--- height, never less than one line; L3.7 item 39): the left edge, the hint's indent and width
--- (the width minus the 40 px indent and the right padding), the width of Reset's sentence
--- and of the divider, and the y of the title, the warning line (nil without one), each
--- switch row, Reset, the commands heading, its divider and lines, and the two foot lines,
--- and the height it all needs with the foot's pad (the page scrolls when the canvas is
--- shorter). A backdrop the client has not sized yet is laid out at
--- Theme.SIZES.settingsWidth. Pure.
function SettingsPage.layout(width, rowHeights, warned)
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
		dividerWidth = width - S.padding * 2,
		rows = {},
		commands = {},
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
	for index = 1, #SettingsPage.SWITCHES do
		layout.rows[index] = -y
		local height = type(rowHeights) == "table" and rowHeights[index] or nil
		y = y + ((type(height) == "number" and height > 0) and height or S.settingsRow)
	end
	y = y + S.padding
	layout.reset = -y
	y = y + S.buttonHeight + S.padding
	layout.commandsHeading = -y
	y = y + S.settingsLine
	layout.divider = -y
	y = y + S.border + S.gap * 2
	for index = 1, #L.settingsCommands do
		layout.commands[index] = -y
		y = y + S.settingsLine
	end
	y = y + S.padding
	for index = 1, 2 do
		layout.data[index] = -y
		y = y + S.settingsLine
	end
	layout.height = y + S.settingsFootPad
	return layout
end

--- The page's own controls: one toggle per switch and the Reset button (labels and the
--- divider are regions, not children).
function SettingsPage.expectedChildren()
	return #SettingsPage.SWITCHES + 1
end

--- How long Reset reads "Positions reset", disabled, after a click.
SettingsPage.RESET_SECONDS = 2

--- Reset's look: "Reset positions" (gold, in the border colour), or "Positions reset" and
--- disabled for SettingsPage.RESET_SECONDS after a click, its label and its 1 px border in
--- the muted-text grey at full alpha (fix round 2, item 4: the faded disabled look did not
--- read). { label, enabled, text, edge }, the colours Theme.HEX keys. Pure: reads
--- SettingsPage.resetting.
function SettingsPage.resetLook()
	local states = Theme.CONTROL_STATES.face
	if SettingsPage.resetting then
		return { label = L.settingsResetDone, enabled = false, text = states.waitingLabel, edge = states.waitingEdge }
	end
	return { label = L.settingsReset, enabled = true, text = states.label, edge = states.edge }
end

local function drawReset()
	local button = SettingsPage.resetButton
	if button == nil then
		return
	end
	local look = SettingsPage.resetLook()
	button:SetText(look.label)
	Widgets.setButtonLook(button, look.enabled, look.text, look.edge)
end

--- Reset positions: Prefs puts the placements back, and its onReset listeners move the
--- chip, the minimap button and the build dialog at once (Chip.place,
--- MinimapButton.refresh, BuildDialog.place). The button then says so for 2 s; a client
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
	local layout = SettingsPage.layout(SettingsPage.backdrop:GetWidth(), heights, warned)
	place(SettingsPage.title, content, layout.x, layout.title)
	place(SettingsPage.warningLabel, content, layout.x, layout.warning or layout.title)
	for index, toggle in ipairs(SettingsPage.toggles) do
		place(toggle.frame, content, layout.x, layout.rows[index])
	end
	place(SettingsPage.resetButton, content, layout.x, layout.reset)
	place(SettingsPage.commandsHeading, content, layout.x, layout.commandsHeading)
	place(SettingsPage.divider, content, layout.x, layout.divider)
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
	SettingsPage.title = Widgets.label(content, L.addonName, "gold", "large")
	-- The first data state that is not current, repeated where the page opens (item 5).
	local warning = Widgets.label(content, "", "warning", "small")
	warning:SetWidth(layout.dividerWidth)
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
	local reset = Widgets.button(content, L.settingsReset, SettingsPage.reset)
	SettingsPage.resetButton = reset
	local resetHint = Widgets.label(content, L.settingsResetHint, "muted", "small")
	resetHint:SetPoint("LEFT", reset, "RIGHT", S.gap * 2, 0)
	resetHint:SetWidth(layout.resetHintWidth)
	SettingsPage.commandsHeading = Widgets.label(content, L.settingsCommandsHeading, "gold", "normal")
	local divider = Theme.texture(content, "ARTWORK", "border")
	divider:SetWidth(layout.dividerWidth)
	divider:SetHeight(S.border)
	SettingsPage.divider = divider
	SettingsPage.commandLabels = {}
	for index, line in ipairs(L.settingsCommands) do
		SettingsPage.commandLabels[index] = Widgets.label(content, SettingsPage.commandText(line), "muted", "small")
	end
	SettingsPage.dataLabels = {}
	for index = 1, #layout.data do
		SettingsPage.dataLabels[index] = Widgets.label(content, "", "muted", "small")
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
		SettingsPage.dataLabels[1]:SetText(addon ~= nil and SettingsPage.lineText(addon) or "")
		SettingsPage.dataLabels[2]:SetText(SettingsPage.lineText(ratings))
		local warning = SettingsPage.warningLine(addon, ratings)
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
