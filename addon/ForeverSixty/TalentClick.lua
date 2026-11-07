-- addon/ForeverSixty/TalentClick.lua
-- The talent window, opened without tainting the action bars.
--
-- Taint (live taint.log, 7 Nov 2026): opening the talent window from addon code runs the window's
-- OnShow in tainted execution. That OnShow calls MultiActionBar_ShowAllGrids, which lays the action
-- bars out again, so MainActionBar.snappedToFrame and the bars' shownButtonContainers are written
-- tainted. The next layout in combat (PetActionBar:Update hides the bar on PLAYER_TARGET_CHANGED,
-- pet or no pet) clears MainActionBar's anchors, reads that field and is refused the SetPoint: the
-- main bar has no anchor and the bars anchored to it go with it. pcall and securecall do not help.
--
-- So on a client with the micro button every talent click is a secure one: a cover built on the
-- client's own SecureActionButtonTemplate sits over the addon's button and clicks
-- TalentMicroButton (talentsOnly: always the Talents tab), and no addon code is on the stack while
-- the window opens. Only a client without the micro button opens the window from addon code.
--
-- A cover is protected, and a frame a protected one is anchored to is restricted in combat too.
-- So a cover is a child of UIParent, never of the frame it covers, and its owner takes it away on
-- PLAYER_REGEN_DISABLED, before the lockdown starts. Nothing touches a cover in the lockdown.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Compat = ns.Compat or require("Compat")

local TalentClick = {}

TalentClick.MICRO_BUTTON = "TalentMicroButton"
TalentClick.TEMPLATE = "SecureActionButtonTemplate"
--- The cover script handlers a cover passes on to its owner.
TalentClick.SCRIPTS = { "PostClick", "OnEnter", "OnLeave", "OnDragStart", "OnDragStop" }

--- The client's talent micro button, or nil where this client has none.
function TalentClick.micro()
	return Compat.clientGlobal(TalentClick.MICRO_BUTTON, "table")
end

--- The addon-code way in, for an addon button's own click. Where the micro button exists a cover
--- takes that click, so this is only reached with the cover away (in combat) and does nothing.
--- True when the window was toggled.
function TalentClick.open()
	if TalentClick.micro() ~= nil then
		return false
	end
	return Theme.toggleTalents()
end

--- A cover, not built yet: `name` is its global name, `clicks` the one or two mouse buttons it
--- takes (RegisterForClicks), `scripts` the owner's handlers by TalentClick.SCRIPTS name. Only the left
--- button has a secure action; any other click reaches PostClick alone.
function TalentClick.new(name, clicks, scripts)
	return { name = name, clicks = clicks, scripts = scripts or {}, button = nil, target = nil }
end

local function build(cover, micro)
	local ok, button = pcall(CreateFrame, "Button", cover.name, UIParent, TalentClick.TEMPLATE)
	if not ok or type(button) ~= "table" then
		cover.refused = true
		Theme.note(L.diagNoTalentCover)
		return nil
	end
	Theme.noteFrame(cover.name, "the first talent button")
	if cover.clicks[2] ~= nil then
		button:RegisterForClicks(cover.clicks[1], cover.clicks[2])
	else
		button:RegisterForClicks(cover.clicks[1])
	end
	if cover.scripts.OnDragStart ~= nil then
		button:RegisterForDrag("LeftButton")
	end
	-- On the release, as the covered button's own click, whatever ActionButtonUseKeyDown says.
	button:SetAttribute("useOnKeyDown", false)
	-- The left button with any modifier; no other button has a type, so no other secure action.
	button:SetAttribute("*type1", "click")
	button:SetAttribute("*clickbutton1", micro)
	for _, script in ipairs(TalentClick.SCRIPTS) do
		if cover.scripts[script] ~= nil then
			button:SetScript(script, cover.scripts[script])
		end
	end
	button:Hide()
	cover.button = button
	return button
end

--- The cover's button, built on first use; nil on a client without the micro button or the
--- template.
local function ensure(cover)
	if cover.button ~= nil then
		return cover.button
	end
	local micro = TalentClick.micro()
	if micro == nil or cover.refused then
		return nil
	end
	return build(cover, micro)
end

--- The frame's strata and level, as far as the client says.
local function layer(frame)
	local strata = type(frame.GetFrameStrata) == "function" and frame:GetFrameStrata() or nil
	local level = type(frame.GetFrameLevel) == "function" and frame:GetFrameLevel() or nil
	return strata or "MEDIUM", level or 0
end

--- Put the cover over `target`, `raise` levels above it, or take it away when `target` is nil.
--- Does nothing in the lockdown: the owner calls this again when combat ends. True when the cover
--- is over `target` afterwards (always false for nil).
function TalentClick.sync(cover, target, raise)
	if Theme.inCombat() then
		return false
	end
	if target == cover.target then
		return target ~= nil
	end
	local button = target ~= nil and ensure(cover) or cover.button
	if button == nil then
		return false
	end
	button:Hide()
	button:ClearAllPoints()
	cover.target = nil
	if target ~= nil then
		local strata, level = layer(target)
		button:SetFrameStrata(strata)
		button:SetFrameLevel(level + (raise or 1))
		button:SetAllPoints(target)
		button:Show()
		cover.target = target
	end
	return target ~= nil
end

ns.TalentClick = TalentClick
return TalentClick
