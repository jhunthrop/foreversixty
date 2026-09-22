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
--- `group` starts a new headed section at that entry; `hint` is the line
--- under the label that says what the setting is for.
SettingsView.TOGGLES = {
	{ section = "minimap", key = "shown", label = "settingsMinimap", hint = "settingsMinimapHint",
		group = "settingsGroupScreen" },
	{ section = "tracker", key = "shown", label = "settingsTracker", hint = "settingsTrackerHint" },
	{ section = "tracker", key = "locked", label = "settingsTrackerLocked", hint = "settingsTrackerLockedHint" },
	{ flag = "tooltip", label = "settingsTooltip", hint = "settingsTooltipHint" },
	{ flag = "toast", label = "settingsToast", hint = "settingsToastHint" },
	{ flag = "autoSave", label = "settingsAutoSave", hint = "settingsAutoSaveHint",
		group = "settingsGroupData" },
	{ flag = "chat", label = "settingsChat", hint = "settingsChatHint" },
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
			hint = entry.hint ~= nil and L[entry.hint] or nil,
			group = entry.group ~= nil and L[entry.group] or nil,
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
function SettingsView.build(holder, ctx)
	local gap, padding = Theme.SIZES.gap, Theme.SIZES.padding
	-- The list outgrew the page once every setting had a line explaining it
	-- (seen in game), so the page scrolls with the wheel.
	local scroll = Widgets.scrollable(holder)
	local parent = scroll.content
	local view = { frame = holder, ctx = ctx, toggles = {}, scroll = scroll }
	view.title = Widgets.label(parent, L.settingsTitle, "gold")
	view.title:SetPoint("TOPLEFT", parent, "TOPLEFT", padding, -padding)
	-- Every row is anchored to the page's left edge at the same x, and only
	-- its y comes from the row above. Anchoring each toggle to the indented
	-- hint under the previous one made the list staircase to the right,
	-- like sub-bullets (seen in game).
	local y = -(padding + Theme.SIZES.rowHeight)
	local function place(region, indent, gapAbove, height)
		y = y - gapAbove
		region:SetPoint("TOPLEFT", parent, "TOPLEFT", padding + indent, y)
		y = y - height
		return region
	end
	for index, entry in ipairs(SettingsView.toggles()) do
		if entry.group ~= nil then
			place(Widgets.label(parent, entry.group, "muted", "small"), 0, padding, Theme.SIZES.rowHeight)
		end
		local toggle = Widgets.toggle(parent, entry.label, entry.checked, function(value)
			SettingsView.write(entry, value, ctx)
		end)
		place(toggle.frame, 0, gap * 2, Theme.SIZES.rowHeight)
		view.toggles[index] = toggle
		if entry.hint ~= nil then
			-- Under the label, not under the tick box.
			place(Widgets.label(parent, entry.hint, "muted", "small"),
				Theme.SIZES.iconSize + gap, -gap, Theme.SIZES.rowHeight)
		end
	end
	view.reset = Widgets.button(parent, L.settingsReset, function()
		Prefs.resetPositions()
		ctx.onPrefChanged({ reset = true }, true)
	end)
	place(view.reset, 0, padding, Theme.SIZES.buttonHeight)
	scroll:SetContentHeight(-y + padding)
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
	SettingsView.how = Theme.registerSettingsPanel(panel, L.addonName)
	if SettingsView.how ~= nil then
		SettingsView.panelView = SettingsView.build(panel, ctx)
	end
	return SettingsView.how
end

ns.SettingsView = SettingsView
return SettingsView
