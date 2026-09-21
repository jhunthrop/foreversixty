-- addon/ForeverSixty/views/SettingsView.lua
-- The settings, in two places, from one function.
--
-- build() is the whole page. The window's Settings tab calls it; the
-- frame registered with the game's options panel calls it. There is no
-- third copy of the list of toggles and no second place a default can
-- disagree with itself: TOGGLES is the list, Prefs is the state, and
-- ctx.onPrefChanged is how a flipped toggle reaches the surface it
-- controls.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Prefs = ns.Prefs or require("Prefs")

local SettingsView = {}

SettingsView.PANEL_NAME = "ForeverSixtySettingsPanel"

--- Each entry names either a field inside a Prefs section or a top-level
--- flag, plus the Locale key for its label. Order is the order drawn.
SettingsView.TOGGLES = {
	{ section = "minimap", key = "shown", label = "settingsMinimap" },
	{ section = "tracker", key = "shown", label = "settingsTracker" },
	{ section = "tracker", key = "locked", label = "settingsTrackerLocked" },
	{ flag = "autoSave", label = "settingsAutoSave" },
	{ flag = "chat", label = "settingsChat" },
}

--- A toggle's value can legitimately be false. An explicit if, not an
--- and/or chain, so a stored false reads as false rather than collapsing
--- to whatever the chain treats as "absent".
local function valueOf(entry)
	if entry.flag ~= nil then
		return Prefs.flag(entry.flag)
	end
	return Prefs.get(entry.section, entry.key)
end

function SettingsView.toggles()
	local rows = {}
	for index, entry in ipairs(SettingsView.TOGGLES) do
		rows[index] = {
			label = L[entry.label],
			checked = valueOf(entry),
			flag = entry.flag,
			section = entry.section,
			key = entry.key,
		}
	end
	return rows
end

function SettingsView.write(entry, value, ctx)
	if entry.flag ~= nil then
		Prefs.setFlag(entry.flag, value)
	else
		Prefs.set(entry.section, entry.key, value)
	end
	ctx.onPrefChanged(entry, value)
	return value
end

--- The page. Called once for the window tab and once for the game panel.
function SettingsView.build(parent, ctx)
	local gap, padding = Theme.SIZES.gap, Theme.SIZES.padding
	local view = { frame = parent, ctx = ctx, toggles = {} }
	view.title = Widgets.label(parent, L.settingsTitle, "gold")
	view.title:SetPoint("TOPLEFT", parent, "TOPLEFT", padding, -padding)
	local above = view.title
	for index, entry in ipairs(SettingsView.toggles()) do
		local toggle = Widgets.toggle(parent, entry.label, entry.checked, function(value)
			SettingsView.write(entry, value, ctx)
		end)
		toggle.frame:SetPoint("TOPLEFT", above, "BOTTOMLEFT", 0, -gap)
		view.toggles[index] = toggle
		above = toggle.frame
	end
	view.reset = Widgets.button(parent, L.settingsReset, function()
		Prefs.resetPositions()
		ctx.onPrefChanged({ reset = true }, true)
	end)
	view.reset:SetPoint("TOPLEFT", above, "BOTTOMLEFT", 0, -padding)
	function view.refresh()
		for index, entry in ipairs(SettingsView.toggles()) do
			view.toggles[index]:SetChecked(entry.checked)
		end
		return view
	end
	return view
end

function SettingsView.mount(parent, ctx)
	return SettingsView.build(parent, ctx)
end

--- The page in the game's own options. Built on first registration, and
--- only once: registering twice makes two entries in the player's list.
--- The content is only drawn once the client actually has somewhere to
--- put it -- a client with neither options API gets the one diagnostic
--- line and nothing more; the widgets that would fill an orphan panel
--- never get built.
function SettingsView.register(ctx)
	if SettingsView.panel ~= nil then
		return SettingsView.how
	end
	local panel = CreateFrame("Frame", SettingsView.PANEL_NAME, UIParent)
	-- The old InterfaceOptions API reads the category's label off this
	-- field; the modern one takes it as an argument. Setting both costs
	-- one line and removes the branch.
	panel.name = L.addonName
	SettingsView.panel = panel
	SettingsView.how, SettingsView.category = Theme.registerSettingsPanel(panel, L.addonName)
	if SettingsView.how ~= nil then
		SettingsView.panelView = SettingsView.build(panel, ctx)
	end
	return SettingsView.how
end

function SettingsView.open()
	return Theme.openSettingsPanel(SettingsView.category)
end

ns.SettingsView = SettingsView
return SettingsView
