-- addon/ForeverSixty/Window.lua
-- The one window: a title bar, a sidebar of pages, a character header,
-- and the frame the pages live in.
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
local Cards = ns.Cards or require("Cards")
local Prefs = ns.Prefs or require("Prefs")
local Talents = ns.Talents or require("Talents")
local Gear = ns.Gear or require("Gear")
local Tracker = ns.Tracker or require("Tracker")
local MinimapButton = ns.Minimap or require("Minimap")

local Window = {}

Window.FRAME_NAME = "ForeverSixtyWindow"

--- The sidebar, top to bottom. `follow` keeps its name (saved prefs and
--- /fs follow use it); its label reads Talents.
Window.TABS = {
	{ name = "overview", label = "tabOverview" },
	{ name = "follow", label = "tabFollow" },
	{ name = "gear", label = "tabGear" },
	{ name = "export", label = "tabExport" },
	{ name = "settings", label = "tabSettings" },
}

Window.VIEWS = {
	overview = ns.OverviewView or require("OverviewView"),
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

--- The addon's own version, as the TOC stamps it. Two homes for the
--- function across clients, and neither is promised.
function Window.addonVersion()
	local read = (type(C_AddOns) == "table" and C_AddOns.GetAddOnMetadata) or GetAddOnMetadata
	if type(read) ~= "function" then
		return nil
	end
	local ok, version = pcall(read, "ForeverSixty", "Version")
	return ok and version or nil
end

--- "Level 11 Gnome Warrior": the client's own localised race and class.
local function levelLine()
	local className = UnitClass("player")
	local raceName = type(UnitRace) == "function" and UnitRace("player") or nil
	return string.format(L.headerLevelLine, UnitLevel("player") or 0, raceName or "", className or "")
end

function Window.headerModel(data)
	local classSlug = Talents.playerClassSlug()
	local spec = classSlug ~= nil
		and Gear.specOf(data, classSlug, Talents.readRanks(data)) or nil
	local warning = Window.mismatchLine(data)
	return {
		name = type(UnitName) == "function" and UnitName("player") or "",
		classToken = select(2, UnitClass("player")),
		levelLine = levelLine(),
		status = warning ~= nil and L.statusOutOfDate or string.format(L.statusDataBuild, data.build),
		statusColor = warning ~= nil and "warning" or "muted",
		realmLine = string.format(L.headerRealmLevel,
			GetRealmName() or "", UnitLevel("player") or 0),
		specLine = spec or L.headerNoSpec,
		buildLine = Window.dataBuildLine(data),
		warning = Window.mismatchLine(data),
	}
end

--- The width a page lays its content out in: the window less the sidebar
--- and the page's own padding on both sides.
function Window.contentWidth()
	return Theme.SIZES.windowWidth - Theme.SIZES.sidebarWidth - Theme.SIZES.padding * 2
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
	local mark = Theme.icon(bar, "ARTWORK", Theme.MEDIA.minimapIcon, Theme.SIZES.navIcon)
	mark:SetPoint("LEFT", bar, "LEFT", Theme.SIZES.padding, 0)
	Window.title = Widgets.label(bar, L.addonName, "gold")
	Window.title:SetPoint("LEFT", mark, "RIGHT", Theme.SIZES.gap * 2, 0)
	local version = Window.addonVersion()
	Window.version = Widgets.label(bar, version and string.format(L.addonVersion, version) or "", "muted", "small")
	Window.version:SetPoint("LEFT", Window.title, "RIGHT", Theme.SIZES.gap * 2, 0)
	Window.closeButton = Widgets.closeButton(bar, function()
		Window.close()
	end)
	Window.closeButton:SetPoint("RIGHT", bar, "RIGHT", -Theme.SIZES.gap, 0)
	return bar
end

--- Where the pages start: under the title bar and the character header,
--- to the right of the sidebar. Everything is inside the window; the first
--- in-game screenshot showed tabs anchored to the bottom edge hanging half
--- outside the frame and covering a page's own buttons.
function Window.pageLeft()
	return Theme.SIZES.sidebarWidth
end

function Window.pageTop()
	return Theme.SIZES.titleBarHeight + Theme.SIZES.headerHeight
end

function Window.pageWidth()
	return Theme.SIZES.windowWidth - Window.pageLeft()
end

function Window.pageHeight()
	return Theme.SIZES.windowHeight - Window.pageTop()
end

local function hairline(frame, fromPoint, toPoint, x, y, horizontal)
	local rule = Theme.texture(frame, "ARTWORK", "border")
	rule:SetPoint(fromPoint, frame, fromPoint, x, y)
	rule:SetPoint(toPoint, frame, toPoint, horizontal and 0 or x, horizontal and y or 0)
	if horizontal then
		rule:SetHeight(Theme.SIZES.border)
	else
		rule:SetWidth(Theme.SIZES.border)
	end
	return rule
end

local function buildHeader(frame)
	local S = Theme.SIZES
	local left, top = Window.pageLeft() + S.padding, -(S.titleBarHeight + S.gap * 3)
	Window.characterName = Widgets.label(frame, "", "gold", "large")
	Window.characterName:SetPoint("TOPLEFT", frame, "TOPLEFT", left, top)
	Window.levelLine = Widgets.label(frame, "", "muted", "small")
	Window.levelLine:SetPoint("TOPLEFT", Window.characterName, "BOTTOMLEFT", 0, -S.gap)
	Window.spec = Widgets.label(frame, "", "body", "small")
	Window.spec:SetPoint("LEFT", Window.levelLine, "RIGHT", 0, 0)
	Window.status = Cards.pill(frame, "", "muted")
	Window.status:SetPoint("TOPRIGHT", frame, "TOPRIGHT", -S.padding, top)
	Window.warning = Widgets.label(frame, "", "warning", "small")
	Window.warning:SetPoint("TOPRIGHT", Window.status, "BOTTOMRIGHT", 0, -S.gap)
	Window.warning:SetWidth(Window.pageWidth() / 2)
	Window.warning:SetJustifyH("RIGHT")
	hairline(frame, "TOPLEFT", "TOPRIGHT", Window.pageLeft(), -Window.pageTop(), true)
	return frame
end

local function buildSidebar(frame)
	local S = Theme.SIZES
	local sidebar = CreateFrame("Frame", nil, frame)
	sidebar:SetPoint("TOPLEFT", frame, "TOPLEFT", 0, -S.titleBarHeight)
	sidebar:SetSize(S.sidebarWidth, S.windowHeight - S.titleBarHeight)
	Theme.texture(sidebar, "BACKGROUND", "sidebar"):SetAllPoints(sidebar)
	hairline(sidebar, "TOPRIGHT", "BOTTOMRIGHT", 0, 0, false)
	for index, tab in ipairs(Window.TABS) do
		local item = Cards.navItem(sidebar, Theme.NAV_ICONS[tab.name], L[tab.label], function()
			Window.select(tab.name)
		end)
		item:SetPoint("TOPLEFT", sidebar, "TOPLEFT", 0, -(S.gap * 2 + (index - 1) * S.navHeight))
		Window.tabs[tab.name] = item
	end
	Window.site = Widgets.label(sidebar, L.siteName, "gold", "small")
	Window.site:SetPoint("BOTTOMLEFT", sidebar, "BOTTOMLEFT", S.padding, S.padding)
	Window.sidebar = sidebar
	return sidebar
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
	Theme.shadow(frame)
	buildTitleBar(frame)
	buildSidebar(frame)
	buildHeader(frame)
	Window.restorePosition()
	Theme.makeEscapable(Window.FRAME_NAME)
	frame:Hide()
	return frame
end

function Window.mountPage(name)
	-- The holder spans the window; each page pads its own content by
	-- Theme.SIZES.padding, which is what contentWidth already allows for.
	local holder = CreateFrame("Frame", nil, Window.frame)
	holder:SetSize(Window.pageWidth(), Window.pageHeight())
	holder:SetPoint("TOPLEFT", Window.frame, "TOPLEFT", Window.pageLeft(), -Window.pageTop())
	return Window.VIEWS[name].mount(holder, Window.context())
end

local function applyHeader(data)
	local header = Window.headerModel(data)
	Window.characterName:SetText(header.name)
	Window.characterName:SetTextColor(Theme.classColor(header.classToken))
	Window.levelLine:SetText(header.levelLine)
	Window.spec:SetText(L.headerSpecSeparator .. header.specLine)
	Window.status:SetText(header.status, header.statusColor)
	Window.warning:SetText(header.warning or "")
	return header
end

function Window.select(name)
	Window.ensure()
	Window.pages[name] = Window.pages[name] or Window.mountPage(name)
	for _, tab in ipairs(Window.TABS) do
		Cards.setNavActive(Window.tabs[tab.name], tab.name == name)
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
	local wasOpen = frame:IsShown()
	frame:Show()
	if not wasOpen then
		Theme.fadeIn(frame, Theme.SIZES.fadeIn)
		Theme.playSound("open")
	end
	applyHeader(Window.data)
	local wanted = tab or Prefs.get("window", "tab")
	return Window.select(Window.VIEWS[wanted] ~= nil and wanted or Window.TABS[1].name)
end

function Window.close()
	if Window.frame ~= nil and Window.frame:IsShown() then
		Window.frame:Hide()
		Theme.playSound("close")
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
