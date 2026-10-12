-- addon/ForeverSixty/surfaces/Settings.lua
-- The Settings page (spec section 4.11, L3.5): the client's own Settings, one canvas.
-- Six switches, the /fs command list, Reset positions, and one line on whether the
-- addon's data is for this client's patch.
--
-- Pure model, thin frame. SettingsPage.SWITCHES is the whole of what the page can change
-- (each key a Prefs.DEFAULTS boolean); read and apply are its only way to the prefs; layout
-- gives every number on the canvas; dataLine chooses the foot line. The canvas is an empty
-- frame registered at login and built the first time the client shows it, never at load.
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

local SettingsPage = {}

SettingsPage.FRAME_NAME = "ForeverSixtySettingsPanel"

--- The switches, in page order (L3.5 ruling 1). A key is a Prefs section and field
--- ("chip.shown") or a top-level flag ("tooltip").
SettingsPage.SWITCHES = {
	{ key = "chip.shown", label = L.settingsChipShown, hint = L.settingsChipShownHint },
	{ key = "chip.combatHide", label = L.settingsCombatHide, hint = L.settingsCombatHideHint },
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
		return type(defaults) == "table" and defaults[field] or nil
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

--- Reset positions: Prefs puts the placements back, and its onReset listeners move the
--- chip, the minimap button and the build dialog at once (Chip.place,
--- MinimapButton.refresh, BuildDialog.place).
function SettingsPage.reset()
	return Prefs.resetPositions()
end

--- The foot line: the data is for this client's patch or a later one ("current"), or for an
--- older patch than the client's. Only the patch counts ("1.60.1" against the data's first
--- three fields), never the build number, and no number is printed. Nil when either side is
--- unknown. Pure.
function SettingsPage.dataLine(dataBuild, clientVersion)
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
		return L.settingsDataOlder
	end
	return L.settingsDataCurrent
end

--- Every number on the canvas, from its width: the left edge, the hint's indent and
--- width (the canvas width minus the 40 px indent and the right padding), Reset's hint
--- width, and the y of the title, each switch row, the commands heading and lines, Reset
--- and the foot line, and the height it all needs. A canvas the client has not sized yet is
--- laid out at Theme.SIZES.settingsWidth. Pure.
function SettingsPage.layout(width)
	local S = Theme.SIZES
	if type(width) ~= "number" or width <= 0 then
		width = S.settingsWidth
	end
	local layout = {
		x = S.padding,
		hintIndent = Widgets.toggleIndent(),
		hintWidth = width - S.settingsHintInset - S.padding,
		-- Reset's hint sits right of the button and wraps before the right padding.
		resetHintWidth = width - S.padding * 2 - S.buttonWidth - S.gap * 2,
		rows = {},
		commands = {},
	}
	local y = S.padding
	layout.title = -y
	y = y + S.settingsTitle + S.gap * 2
	for index = 1, #SettingsPage.SWITCHES do
		layout.rows[index] = -y
		y = y + S.settingsRow
	end
	y = y + S.padding
	layout.commandsHeading = -y
	y = y + S.settingsLine
	for index = 1, #L.settingsCommands do
		layout.commands[index] = -y
		y = y + S.settingsLine
	end
	y = y + S.padding
	layout.reset = -y
	y = y + S.buttonHeight + S.padding
	layout.data = -y
	y = y + S.settingsLine
	layout.height = y + S.padding
	return layout
end

--- The canvas's own controls: one toggle per switch and the Reset button (labels are
--- regions, not children).
function SettingsPage.expectedChildren()
	return #SettingsPage.SWITCHES + 1
end

-- The frame -------------------------------------------------------------------

local function place(region, panel, x, y)
	region:SetPoint("TOPLEFT", panel, "TOPLEFT", x, y)
	return region
end

--- Build the canvas's contents once, at the width the client gave it.
function SettingsPage.build(panel)
	if SettingsPage.toggles ~= nil then
		return panel
	end
	local S = Theme.SIZES
	local layout = SettingsPage.layout(panel:GetWidth())
	place(Widgets.label(panel, L.addonName, "gold", "large"), panel, layout.x, layout.title)
	SettingsPage.toggles = {}
	for index, switch in ipairs(SettingsPage.switches()) do
		local key = switch.key
		local toggle = Widgets.toggle(panel, switch.label, SettingsPage.read(key), function(value)
			SettingsPage.onToggle(key, value)
		end, switch.hint, layout.hintWidth)
		toggle.key = key
		place(toggle.frame, panel, layout.x, layout.rows[index])
		SettingsPage.toggles[index] = toggle
	end
	place(Widgets.label(panel, L.settingsCommandsHeading, "body", "normal"), panel, layout.x, layout.commandsHeading)
	SettingsPage.commandLabels = {}
	for index, line in ipairs(L.settingsCommands) do
		SettingsPage.commandLabels[index] = place(Widgets.label(panel, line, "muted", "small"), panel,
			layout.x, layout.commands[index])
	end
	local reset = Widgets.button(panel, L.settingsReset, SettingsPage.reset)
	place(reset, panel, layout.x, layout.reset)
	SettingsPage.resetButton = reset
	local resetHint = Widgets.label(panel, L.settingsResetHint, "muted", "small")
	resetHint:SetPoint("LEFT", reset, "RIGHT", S.gap * 2, 0)
	resetHint:SetWidth(layout.resetHintWidth)
	local dataLabel = place(Widgets.label(panel, "", "muted", "small"), panel, layout.x, layout.data)
	SettingsPage.dataLabel = dataLabel
	return panel
end

--- Show the live prefs on every toggle, and the foot line for this client.
function SettingsPage.refresh()
	for _, toggle in ipairs(SettingsPage.toggles or {}) do
		toggle:SetChecked(SettingsPage.read(toggle.key))
	end
	if SettingsPage.dataLabel ~= nil then
		local data = SettingsPage.data or ns.Data
		local line = SettingsPage.dataLine(type(data) == "table" and data.build or nil, Compat.clientVersion())
		SettingsPage.dataLabel:SetText(line or "")
	end
end

--- The render check (spec section 2.4), on every show: the canvas holds every control it
--- was built with. Otherwise one /fs diag note for the session. True when it passes.
function SettingsPage.selfCheck()
	local panel = SettingsPage.panel
	if panel == nil or SettingsPage.toggles == nil then
		return true
	end
	local count = #{ panel:GetChildren() }
	if count == SettingsPage.expectedChildren() then
		return true
	end
	if not SettingsPage.drawNoted then
		SettingsPage.drawNoted = true
		Theme.note(string.format(L.diagSettingsDrawn, count, SettingsPage.expectedChildren()))
	end
	return false
end

function SettingsPage.onShow(panel)
	SettingsPage.build(panel)
	SettingsPage.refresh()
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

--- Open the client's Settings on this page. False when it cannot (no API, not registered,
--- or the client refused); a refusal from a registered page is noted once.
function SettingsPage.open()
	if SettingsPage.how == nil then
		return false
	end
	if Theme.openSettingsPanel(SettingsPage.how, SettingsPage.handle) then
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
