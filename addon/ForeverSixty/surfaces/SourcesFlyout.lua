-- addon/ForeverSixty/surfaces/SourcesFlyout.lua
-- The Sources flyout (design/specs/2026-11-13-addon-source-filter.md 4): the panel's foot line's
-- Change opens it beside the panel. One switch per source kind in the web BiS page's order and words,
-- each with the count of gear slots that have an option from that kind at the character's band;
-- Default (only while a choice is saved) clears the choice. Every click saves at once and every
-- surface re-picks (Sources.listeners); there is no Apply.
--
-- Pure model, thin frame: SourcesFlyout.model says what shows; the frame is built on the first open,
-- never at login, and closes on its X, Escape, Change again, the panel closing and combat.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Sources = ns.Sources or require("Sources")

local SourcesFlyout = {}

SourcesFlyout.FRAME_NAME = "ForeverSixtySourcesFlyout"
SourcesFlyout.WIDTH = 232
--- The gap to the panel's edge, and how close to the screen's right edge the panel may sit before
--- the flyout opens on its left instead.
SourcesFlyout.GAP = 4
SourcesFlyout.RIGHT_ROOM = 260
SourcesFlyout.PAD = 12
SourcesFlyout.ROW = 22
SourcesFlyout.TITLE = 22
SourcesFlyout.FOOT = 30

--- The band the flyout reads and writes (the panel's), and the counts there; SourcesFlyout.context
--- is the panel's injection point: function() -> band, counts.
SourcesFlyout.context = function()
	return nil, nil
end

local function shownState()
	local band, counts = SourcesFlyout.context()
	return Sources.state(band), counts, band
end

--- Lay the frame out for the model now.
local function draw()
	local frame = SourcesFlyout.frame
	if frame == nil then
		return nil
	end
	local state, counts = shownState()
	local model = Sources.choiceRows(state, counts)
	for index, row in ipairs(model.rows) do
		local parts = SourcesFlyout.rows[index]
		parts.toggle:SetChecked(row.checked)
		parts.count:SetText(row.count)
		parts.count:SetTextColor(Theme.rgb(Theme.HEX.umberMuted))
		parts.tip = row.tip
	end
	Theme.setShown({ SourcesFlyout.defaultButton }, model.showDefault)
	return model
end

--- Place the flyout beside `panel`: its right edge with a 4 px gap, its left edge when the panel
--- sits within RIGHT_ROOM of the screen's right edge.
function SourcesFlyout.place(panel)
	local frame = SourcesFlyout.frame
	frame:ClearAllPoints()
	local right = type(panel.GetRight) == "function" and panel:GetRight() or nil
	local screen = type(UIParent) == "table" and type(UIParent.GetRight) == "function" and UIParent:GetRight() or nil
	if type(right) == "number" and type(screen) == "number" and screen - right < SourcesFlyout.RIGHT_ROOM then
		frame:SetPoint("TOPRIGHT", panel, "TOPLEFT", -SourcesFlyout.GAP, 0)
		return "left"
	end
	frame:SetPoint("TOPLEFT", panel, "TOPRIGHT", SourcesFlyout.GAP, 0)
	return "right"
end

local function addRow(frame, index, kind)
	local y = SourcesFlyout.PAD + SourcesFlyout.TITLE + (index - 1) * SourcesFlyout.ROW
	local toggle = Widgets.toggle(frame, Sources.label(kind), false, function()
		local band = select(3, shownState())
		Sources.toggle(kind, band)
	end)
	toggle.frame:SetWidth(SourcesFlyout.WIDTH - SourcesFlyout.PAD * 2)
	toggle.frame:SetPoint("TOPLEFT", frame, "TOPLEFT", SourcesFlyout.PAD, -y)
	local count = Widgets.label(frame, "", "umberMuted", "small")
	count:SetJustifyH("RIGHT")
	count:SetPoint("RIGHT", toggle.frame, "RIGHT", 0, 0)
	local parts = { toggle = toggle, count = count }
	-- The switch keeps its own hover look; the row adds its tooltip after it (our own frame, so its
	-- scripts are ours to wrap).
	local enter, leave = toggle.frame:GetScript("OnEnter"), toggle.frame:GetScript("OnLeave")
	toggle.frame:SetScript("OnEnter", function(self, ...)
		enter(self, ...)
		if parts.tip ~= nil then
			Theme.showLines(self, { Sources.label(kind), Theme.inlineColor("tooltipMuted", parts.tip) }, "ANCHOR_BOTTOM")
		end
	end)
	toggle.frame:SetScript("OnLeave", function(self, ...)
		leave(self, ...)
		Theme.hideTooltip()
	end)
	return parts
end

--- The frame, built on the first open: the panel's own ground and edge, the title and its X, a row
--- per kind, Default and the foot line.
function SourcesFlyout.ensure()
	if SourcesFlyout.frame ~= nil then
		return SourcesFlyout.frame
	end
	local height = SourcesFlyout.PAD * 2 + SourcesFlyout.TITLE + #Sources.KINDS * SourcesFlyout.ROW
		+ SourcesFlyout.FOOT
	local frame = Widgets.umberFrame(UIParent, SourcesFlyout.WIDTH, height, SourcesFlyout.FRAME_NAME)
	frame:SetFrameStrata("MEDIUM")
	frame:SetClampedToScreen(true)
	frame:EnableMouse(true)
	frame:Hide()
	SourcesFlyout.frame = frame
	Theme.makeEscapable(SourcesFlyout.FRAME_NAME)
	local title = Widgets.label(frame, L.sourcesTitle, "bone", "normal")
	Theme.titleFont(title, Theme.SIZES.headingText)
	title:SetPoint("TOPLEFT", frame, "TOPLEFT", SourcesFlyout.PAD, -SourcesFlyout.PAD)
	local close = Widgets.closeButton(frame, SourcesFlyout.close)
	close:SetPoint("TOPRIGHT", frame, "TOPRIGHT", -2, -2)
	SourcesFlyout.closeButton = close
	SourcesFlyout.rows = {}
	for index, kind in ipairs(Sources.KINDS) do
		SourcesFlyout.rows[index] = addRow(frame, index, kind)
	end
	local footTop = SourcesFlyout.PAD + SourcesFlyout.TITLE + #Sources.KINDS * SourcesFlyout.ROW + 4
	local default = Widgets.textButton(frame, L.sourcesDefault, function()
		Sources.reset()
	end)
	default:SetPoint("TOPLEFT", frame, "TOPLEFT", SourcesFlyout.PAD, -footTop)
	SourcesFlyout.defaultButton = default
	local foot = Widgets.label(frame, L.sourcesFlyoutFoot, "umberMuted", "small")
	foot:SetPoint("BOTTOMLEFT", frame, "BOTTOMLEFT", SourcesFlyout.PAD, SourcesFlyout.PAD)
	foot:SetWidth(SourcesFlyout.WIDTH - SourcesFlyout.PAD * 2)
	SourcesFlyout.footLabel = foot
	return frame
end

function SourcesFlyout.isOpen()
	return SourcesFlyout.frame ~= nil and SourcesFlyout.frame:IsShown() == true
end

--- Open beside `panel` (never in combat). True when it is open afterwards.
function SourcesFlyout.open(panel)
	if Theme.inCombat() or panel == nil then
		return false
	end
	local frame = SourcesFlyout.ensure()
	SourcesFlyout.place(panel)
	draw()
	frame:Show()
	return true
end

function SourcesFlyout.close()
	if SourcesFlyout.frame ~= nil then
		SourcesFlyout.frame:Hide()
	end
end

--- Change: open it, or close it when it is open. True when it is open afterwards.
function SourcesFlyout.toggle(panel)
	if SourcesFlyout.isOpen() then
		SourcesFlyout.close()
		return false
	end
	return SourcesFlyout.open(panel)
end

--- Redraw the open flyout (a change made here or on the Settings page).
function SourcesFlyout.refresh()
	if SourcesFlyout.isOpen() then
		draw()
	end
end

ns.SourcesFlyout = SourcesFlyout
return SourcesFlyout
