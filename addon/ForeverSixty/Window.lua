-- addon/ForeverSixty/Window.lua
-- The one window: header, four tabs, and the frame they live in.
--
-- The header's lines live here rather than in Options because the TOC
-- loads this file first: Options calls Window.buildLines for what /fs
-- prints, and the window shows the same strings, so the data build and
-- its mismatch warning cannot drift between the two surfaces.
--
-- A page is built the first time its tab is selected and kept after
-- that, so opening the window costs one page, not four.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Prefs = ns.Prefs or require("Prefs")
local Talents = ns.Talents or require("Talents")
local Gear = ns.Gear or require("Gear")
local Tracker = ns.Tracker or require("Tracker")
local MinimapButton = ns.Minimap or require("Minimap")

local Window = {}

Window.FRAME_NAME = "ForeverSixtyWindow"

Window.TABS = {
	{ name = "export", label = "tabExport" },
	{ name = "follow", label = "tabFollow" },
	{ name = "gear", label = "tabGear" },
	{ name = "settings", label = "tabSettings" },
}

Window.VIEWS = {
	export = ns.ExportView or require("ExportView"),
	follow = ns.FollowView or require("FollowView"),
	gear = ns.GearView or require("GearView"),
	settings = ns.SettingsView or require("SettingsView"),
}

Window.pages = {}
Window.tabs = {}

function Window.tabIndex(name)
	for index, tab in ipairs(Window.TABS) do
		if tab.name == name then
			return index
		end
	end
	return nil
end

function Window.dataBuildLine(data)
	return string.format(L.dataBuild, data.build)
end

--- The data build is "1.60.1.69893" and GetBuildInfo's version is
--- "1.60.1", so the comparison is on the prefix: a client on 1.61 is a
--- client this data no longer describes.
function Window.mismatchLine(data)
	local version = type(GetBuildInfo) == "function" and GetBuildInfo() or nil
	if version ~= nil and data.build:sub(1, #version) ~= version then
		return string.format(L.buildMismatch, data.build, version)
	end
	return nil
end

--- What bare /fs prints, and what the header shows.
function Window.buildLines(data)
	local lines = { Window.dataBuildLine(data) }
	local mismatch = Window.mismatchLine(data)
	if mismatch ~= nil then
		lines[#lines + 1] = mismatch
	end
	lines[#lines + 1] = L.slashHint
	return lines
end

function Window.headerModel(data)
	local classSlug = Talents.playerClassSlug()
	local spec = classSlug ~= nil
		and Gear.specOf(data, classSlug, Talents.readRanks(data)) or nil
	return {
		name = type(UnitName) == "function" and UnitName("player") or "",
		classToken = select(2, UnitClass("player")),
		realmLine = string.format(L.headerRealmLevel,
			GetRealmName() or "", UnitLevel("player") or 0),
		specLine = spec or L.headerNoSpec,
		buildLine = Window.dataBuildLine(data),
		warning = Window.mismatchLine(data),
	}
end

function Window.contentWidth()
	return Theme.SIZES.windowWidth - Theme.SIZES.padding * 2
end

--- One ctx, shared by every page. `data` is refreshed on each call so a
--- page built before Options handed the window its Data.lua still reads
--- the current table through the same reference.
function Window.context()
	local ctx = Window.ctx
	if ctx == nil then
		ctx = {
			contentWidth = Window.contentWidth(),
			select = function(name)
				return Window.select(name)
			end,
			refresh = function()
				return Window.refresh()
			end,
			setTracker = function(shown)
				return Tracker.setShown(shown, Window.data)
			end,
			onPrefChanged = function(entry, value)
				return Window.onPrefChanged(entry, value)
			end,
		}
		Window.ctx = ctx
	end
	ctx.data = Window.data
	return ctx
end

function Window.savePosition()
	local point, _, _, x, y = Window.frame:GetPoint()
	Prefs.set("window", "point", point or Prefs.DEFAULTS.window.point)
	Prefs.set("window", "x", x or 0)
	Prefs.set("window", "y", y or 0)
end

function Window.restorePosition()
	local point = Prefs.get("window", "point")
	Window.frame:ClearAllPoints()
	Window.frame:SetPoint(point, UIParent, point,
		Prefs.get("window", "x"), Prefs.get("window", "y"))
end

local function buildTitleBar(frame)
	local bar = CreateFrame("Frame", nil, frame)
	bar:SetSize(Theme.SIZES.windowWidth, Theme.SIZES.titleBarHeight)
	bar:SetPoint("TOPLEFT", frame, "TOPLEFT", 0, 0)
	bar:EnableMouse(true)
	bar:RegisterForDrag("LeftButton")
	bar:SetScript("OnDragStart", function()
		frame:StartMoving()
	end)
	bar:SetScript("OnDragStop", function()
		frame:StopMovingOrSizing()
		Window.savePosition()
	end)
	Theme.gradient(bar, "titleTop", "titleBottom")
	Window.title = Widgets.label(bar, L.addonName, "gold")
	Window.title:SetPoint("LEFT", bar, "LEFT", Theme.SIZES.padding, 0)
	return bar
end

local function buildHeader(frame)
	local gap = Theme.SIZES.gap
	Window.characterName = Widgets.label(frame, "", "gold", "small")
	Window.characterName:SetPoint("TOPLEFT", frame, "TOPLEFT",
		Theme.SIZES.padding, -(Theme.SIZES.titleBarHeight + gap))
	Window.realm = Widgets.label(frame, "", "muted", "small")
	Window.realm:SetPoint("LEFT", Window.characterName, "RIGHT", gap, 0)
	Window.spec = Widgets.label(frame, "", "muted", "small")
	Window.spec:SetPoint("LEFT", Window.realm, "RIGHT", gap, 0)
	Window.build = Widgets.label(frame, "", "muted", "small")
	Window.build:SetPoint("TOPLEFT", Window.characterName, "BOTTOMLEFT", 0, -gap)
	Window.warning = Widgets.label(frame, "", "warning", "small")
	Window.warning:SetPoint("LEFT", Window.build, "RIGHT", gap, 0)
	return frame
end

local function buildTabs(frame)
	for index, tab in ipairs(Window.TABS) do
		local button = Widgets.tab(frame, L[tab.label], function()
			Window.select(tab.name)
		end)
		button:SetPoint("BOTTOMLEFT", frame, "BOTTOMLEFT",
			Theme.SIZES.padding + (index - 1) * Theme.SIZES.tabWidth, 0)
		Window.tabs[tab.name] = button
	end
	return frame
end

function Window.ensure()
	if Window.frame ~= nil then
		return Window.frame
	end
	local frame = Widgets.panel(UIParent, Theme.SIZES.windowWidth,
		Theme.SIZES.windowHeight, Window.FRAME_NAME)
	frame:SetFrameStrata("DIALOG")
	frame:SetMovable(true)
	frame:EnableMouse(true)
	frame:SetClampedToScreen(true)
	frame:RegisterForDrag("LeftButton")
	frame:SetScript("OnDragStart", function(self)
		self:StartMoving()
	end)
	frame:SetScript("OnDragStop", function(self)
		self:StopMovingOrSizing()
		Window.savePosition()
	end)
	Window.frame = frame
	buildTitleBar(frame)
	buildHeader(frame)
	buildTabs(frame)
	Window.restorePosition()
	Theme.makeEscapable(Window.FRAME_NAME)
	frame:Hide()
	return frame
end

function Window.mountPage(name)
	local holder = CreateFrame("Frame", nil, Window.frame)
	holder:SetSize(Window.contentWidth(),
		Theme.SIZES.windowHeight - Theme.SIZES.titleBarHeight - Theme.SIZES.tabHeight)
	holder:SetPoint("TOPLEFT", Window.frame, "TOPLEFT",
		Theme.SIZES.padding, -(Theme.SIZES.titleBarHeight + Theme.SIZES.padding * 3))
	return Window.VIEWS[name].mount(holder, Window.context())
end

local function applyHeader(data)
	local header = Window.headerModel(data)
	Window.characterName:SetText(header.name)
	Window.characterName:SetTextColor(Theme.classColor(header.classToken))
	Window.realm:SetText(header.realmLine)
	Window.spec:SetText(header.specLine)
	Window.build:SetText(header.buildLine)
	Window.warning:SetText(header.warning or "")
	return header
end

function Window.select(name)
	Window.ensure()
	Window.pages[name] = Window.pages[name] or Window.mountPage(name)
	for _, tab in ipairs(Window.TABS) do
		Widgets.setTabActive(Window.tabs[tab.name], tab.name == name)
		local page = Window.pages[tab.name]
		if page ~= nil then
			if tab.name == name then
				page.frame:Show()
			else
				page.frame:Hide()
			end
		end
	end
	Prefs.set("window", "tab", name)
	Window.current = name
	Window.pages[name].refresh()
	return Window.pages[name]
end

function Window.isOpen()
	return Window.frame ~= nil and Window.frame:IsShown()
end

function Window.open(tab)
	local frame = Window.ensure()
	frame:Show()
	applyHeader(Window.data)
	return Window.select(tab or Prefs.get("window", "tab"))
end

function Window.close()
	if Window.frame ~= nil then
		Window.frame:Hide()
	end
	return nil
end

function Window.toggle(tab)
	if Window.isOpen() then
		return Window.close()
	end
	return Window.open(tab)
end

--- Only the page that is showing: refreshing a hidden one costs a talent
--- read and an item sweep for something nobody is looking at.
function Window.refresh()
	if not Window.isOpen() then
		return nil
	end
	Window.context()
	applyHeader(Window.data)
	local page = Window.pages[Window.current]
	if page ~= nil then
		page.refresh()
	end
	return page
end

function Window.onPrefChanged(entry, value)
	if entry.reset then
		if Window.frame ~= nil then
			Window.restorePosition()
		end
		if Tracker.frame ~= nil then
			Tracker.restorePosition()
		end
		MinimapButton.refresh()
		return value
	end
	if entry.section == "tracker" and entry.key == "shown" then
		Tracker.setShown(value, Window.data)
	elseif entry.section == "tracker" and entry.key == "locked" then
		Tracker.setLocked(value)
	elseif entry.section == "minimap" and entry.key == "shown" then
		MinimapButton.setShown(value)
	end
	return value
end

ns.Window = Window
return Window
