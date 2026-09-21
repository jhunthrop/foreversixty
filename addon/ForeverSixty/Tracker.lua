-- addon/ForeverSixty/Tracker.lua
-- The small frame that says what to spend the next point on.
--
-- It exists so a player levelling with a build open does not have to open
-- anything: two lines, draggable, lockable, and gone once the build is
-- finished. model() is the whole of what it says, so the "complete"
-- state, the counts and the hidden-with-no-build state are covered
-- without a screen.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Prefs = ns.Prefs or require("Prefs")
local Follow = ns.Follow or require("Follow")
local Talents = ns.Talents or require("Talents")

local Tracker = {}

Tracker.FRAME_NAME = "ForeverSixtyTracker"

function Tracker.model(data, build, ranks)
	if build == nil then
		return { shown = false, done = false, title = L.followNone, progress = "" }
	end
	local total = #build.order
	local point = Follow.nextPoint(build, ranks)
	if point == nil then
		return {
			shown = true,
			done = true,
			title = L.trackerComplete,
			progress = string.format(L.trackerProgress, total, total),
		}
	end
	return {
		shown = true,
		done = false,
		title = string.format(L.trackerNext, Follow.line(data, build, ranks)),
		progress = string.format(L.trackerProgress, point.index - 1, total),
	}
end

function Tracker.savePosition()
	local point, _, _, x, y = Tracker.frame:GetPoint()
	Prefs.set("tracker", "point", point or Prefs.DEFAULTS.tracker.point)
	Prefs.set("tracker", "x", x or 0)
	Prefs.set("tracker", "y", y or 0)
end

--- May assume the frame exists; callers only reach this once it does.
function Tracker.restorePosition()
	local point = Prefs.get("tracker", "point")
	Tracker.frame:ClearAllPoints()
	Tracker.frame:SetPoint(point, UIParent, point,
		Prefs.get("tracker", "x"), Prefs.get("tracker", "y"))
end

--- Built on first need, never at load.
function Tracker.ensure()
	if Tracker.frame ~= nil then
		return Tracker.frame
	end
	local frame = Widgets.panel(UIParent, Theme.SIZES.trackerWidth,
		Theme.SIZES.trackerHeight, Tracker.FRAME_NAME)
	frame:SetMovable(true)
	frame:EnableMouse(true)
	frame:SetClampedToScreen(true)
	frame:RegisterForDrag("LeftButton")
	frame:SetScript("OnDragStart", function(self)
		if not Prefs.get("tracker", "locked") then
			self:StartMoving()
		end
	end)
	frame:SetScript("OnDragStop", function(self)
		self:StopMovingOrSizing()
		Tracker.savePosition()
	end)
	Tracker.frame = frame
	Tracker.title = Widgets.label(frame, "", "gold", "small")
	Tracker.title:SetPoint("TOPLEFT", frame, "TOPLEFT",
		Theme.SIZES.gap, -Theme.SIZES.gap)
	Tracker.progress = Widgets.label(frame, "", "muted", "small")
	Tracker.progress:SetPoint("TOPLEFT", Tracker.title, "BOTTOMLEFT", 0, -Theme.SIZES.gap)
	Tracker.restorePosition()
	return frame
end

function Tracker.hide()
	if Tracker.frame ~= nil then
		Tracker.frame:Hide()
	end
	return nil
end

--- Draw the current state, building the frame only if there is something
--- to draw. A client with no C_Timer keeps "Build complete" on screen
--- instead of never showing it: the player can turn the tracker off.
function Tracker.refresh(data)
	local model = Tracker.model(data, Follow.build, Talents.readRanks(data))
	if not model.shown or not Prefs.get("tracker", "shown") then
		Tracker.hide()
		return model
	end
	local frame = Tracker.ensure()
	Tracker.title:SetText(model.title)
	Tracker.progress:SetText(model.progress)
	frame:Show()
	if model.done then
		Theme.after(Theme.SIZES.completeSeconds, Tracker.hide)
	end
	return model
end

function Tracker.setShown(shown, data)
	Prefs.set("tracker", "shown", shown)
	if not shown then
		Tracker.hide()
		return shown
	end
	Tracker.refresh(data)
	return shown
end

function Tracker.setLocked(locked)
	Prefs.set("tracker", "locked", locked)
	return locked
end

ns.Tracker = Tracker
return Tracker
