-- addon/ForeverSixty/surfaces/Settings.lua
-- The Settings page (spec section 4.11, L3.5): the client's own Settings, one canvas.
-- Seven switches, Reset positions, the /fs command list, and two lines at the foot: whether
-- the addon's data is for this client's patch, and the player ratings' state. The first of
-- those that is not current is repeated under the title, in the warning colour.
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

--- The switches, in page order (L3.5 ruling 1, fix round 1 item 4). A key is a Prefs section
--- and field ("chip.shown") or a top-level flag ("tooltip").
SettingsPage.SWITCHES = {
	{ key = "chip.shown", label = L.settingsChipShown, hint = L.settingsChipShownHint },
	{ key = "chip.combatHide", label = L.settingsCombatHide, hint = L.settingsCombatHideHint },
	{ key = "talentMarks", label = L.settingsTalentMarks, hint = L.settingsTalentMarksHint },
	{ key = "minimap.shown", label = L.settingsMinimap, hint = L.settingsMinimapHint },
	{ key = "tooltip", label = L.settingsTooltip, hint = L.settingsTooltipHint },
	{ key = "broadcast", label = L.settingsBroadcast, hint = L.settingsBroadcastHint },
	{ key = "autoSave", label = L.settingsAutoSave, hint = L.settingsAutoSaveHint },
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
		Prefs.set("chip", "hiddenUntilLevel", 0)
		Prefs.set("chip", "hiddenFor", "")
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

--- The foot's first line: the addon's data is for this client's patch or a later one
--- ("current for this patch", muted), or for an older patch ("an older patch. Update the
--- addon.", in the warning colour). Only the patch counts ("1.60.1" against the data's first
--- three fields), never the build number, and no number is printed. Nil when either side is
--- unknown. { label, value, tone, current }. Pure.
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
		return { label = L.settingsAddonDataLabel, value = L.settingsAddonDataOlder, tone = "warning", current = false }
	end
	return { label = L.settingsAddonDataLabel, value = L.settingsAddonDataCurrent, tone = "muted", current = true }
end

--- The foot's second line, "Player ratings:" (fix round 2, item 6): the nightly ratings
--- data (the Forever Sixty Data addon) in the state Score.status finds it: not installed
--- (muted, it is optional), current, too old (stale, unreadable or an older format), or made
--- for a newer addon than this one. { label, value, tone, current }. Pure.
function SettingsPage.ratingsLine(status)
	local value, tone
	if type(status) ~= "table" or status.absent then
		value, tone = L.settingsRatingsMissing, "muted"
	elseif status.available and not status.stale then
		value, tone = L.settingsRatingsCurrent, "muted"
	elseif not status.available and status.reason == L.ratingsUpdateAddon then
		value, tone = L.settingsRatingsNewer, "warning"
	else
		value, tone = L.settingsRatingsOld, "warning"
	end
	return { label = L.settingsRatingsDataLabel, value = value, tone = tone, current = value == L.settingsRatingsCurrent }
end

--- The line under the title (fix round 2, item 5): the first foot line that is not current,
--- the addon data before the player ratings, as one sentence the frame draws in the warning
--- colour; nil while both are current (an unknown addon-data line counts as current). Pure.
function SettingsPage.warningLine(addon, ratings)
	if addon ~= nil and not addon.current then
		return addon.label .. " " .. addon.value
	end
	if ratings ~= nil and not ratings.current then
		return ratings.label .. " " .. ratings.value
	end
	return nil
end

--- A foot line as drawn on a muted font string: the label, then the state in its tone.
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
--- the title carries the warning line (`warned`): the left edge, the hint's indent and width
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
		y = y + S.settingsLine
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

--- Build the page once, at the size the client gave the canvas: an opaque backdrop 8 px
--- inside it (the dialog's own, so no world text shows through a hint), scrolling when the
--- canvas is shorter than the page, and everything else on its content frame, arranged once
--- each row has measured its wrapped hint.
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
--- Reset's look, the two foot lines for this client and, while either is not current, the
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
