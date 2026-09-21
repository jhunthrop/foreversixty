-- addon/ForeverSixty/Toast.lua
-- The level-up / next-talent-point toast: a few seconds of "Level 12.
-- Take Improved Rend, rank 2 of 3" at the top of the screen.
--
-- model() is pure and is what the specs cover. Two triggers call into it,
-- wired in Options.lua: PLAYER_LEVEL_UP always fires (see onLevelUp in the
-- next task), and a talent-ish event fires only when a fresh read of the
-- unspent-point count rose since the last one -- the design's own explicit
-- fallback for a client whose trait config will not answer that question
-- is to skip that path and rely on PLAYER_LEVEL_UP alone.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Follow = ns.Follow or require("Follow")
local Talents = ns.Talents or require("Talents")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Prefs = ns.Prefs or require("Prefs")
local Window = ns.Window or require("Window")

local Toast = {}

--- The message for `level` and the next point in `build`. nil when there
--- is nothing worth telling the player: no build loaded, or it is done.
function Toast.model(data, build, ranks, level)
	if build == nil then
		return nil
	end
	local point = Follow.nextPoint(build, ranks)
	if point == nil then
		return nil
	end
	local talent = Talents.cellOf(data, build.classSlug, point.tab, point.tier, point.column)
	local have = (ranks[point.tab] or {})[Talents.cellKey(point.tier, point.column)] or 0
	local name = talent and talent.name or string.format(L.followUnknownCell, point.tier, point.column)
	local maxRank = talent and talent.maxRank or (have + 1)
	return { text = string.format(L.toastMessage, level or 0, name, have + 1, maxRank) }
end

--- Best-effort unspent talent point count from the trait system. Neither
--- field name is confirmed against the 1.60 client (README spike row 24);
--- a shape this client does not use reads as "cannot tell" -- the design's
--- own explicit fallback, under which the talent-event trigger never fires
--- and the toast is carried by PLAYER_LEVEL_UP alone.
function Toast.unspentPoints()
	if type(C_ClassTalents) ~= "table" or type(C_ClassTalents.GetActiveConfigID) ~= "function" then
		return nil
	end
	local configID = C_ClassTalents.GetActiveConfigID()
	if configID == nil or type(C_Traits) ~= "table" or type(C_Traits.GetConfigInfo) ~= "function"
		or type(C_Traits.GetTreeInfo) ~= "function" then
		return nil
	end
	local ok, info = pcall(C_Traits.GetConfigInfo, configID)
	if not ok or type(info) ~= "table" or type(info.treeIDs) ~= "table" then
		return nil
	end
	local total, found = 0, false
	for _, treeID in ipairs(info.treeIDs) do
		local okTree, treeInfo = pcall(C_Traits.GetTreeInfo, configID, treeID)
		if okTree and type(treeInfo) == "table" then
			local points = treeInfo.pointsAvailable or treeInfo.unspentPoints or treeInfo.points
			if type(points) == "number" then
				total, found = total + points, true
			end
		end
	end
	if not found then
		return nil
	end
	return total
end

--- Pure: whether a fresh read is worth a toast. Either side nil means
--- "cannot tell" and never fires; only a confirmed rise does.
function Toast.grewSince(previous, current)
	return previous ~= nil and current ~= nil and current > previous
end

Toast.FRAME_NAME = "ForeverSixtyToast"
Toast.WIDTH = 320
Toast.HEIGHT = 32
Toast.FADE_SECONDS = 5
Toast.TOP_OFFSET = -80

function Toast.ensure()
	if Toast.frame ~= nil then
		return Toast.frame
	end
	local frame = Widgets.panel(UIParent, Toast.WIDTH, Toast.HEIGHT, Toast.FRAME_NAME)
	frame:SetFrameStrata("HIGH")
	frame:SetPoint("TOP", UIParent, "TOP", 0, Toast.TOP_OFFSET)
	frame:EnableMouse(true)
	frame:SetScript("OnMouseUp", function()
		Window.open("follow")
	end)
	Toast.frame = frame
	Toast.text = Widgets.label(frame, "", "gold", "small")
	Toast.text:SetPoint("CENTER", frame, "CENTER", 0, 0)
	frame:Hide()
	return frame
end

function Toast.hide()
	if Toast.frame ~= nil then
		Toast.frame:Hide()
	end
	return nil
end

--- Show `model` now, or hold it until combat ends: the design's combat
--- rule applies to a toast popping up mid-fight just as much as to a
--- protected action.
function Toast.show(model)
	if model == nil or not Prefs.flag("toast") then
		return nil
	end
	if Theme.inCombat() then
		Toast.pending = model
		return nil
	end
	local frame = Toast.ensure()
	Toast.text:SetText(model.text)
	frame:Show()
	Theme.after(Toast.FADE_SECONDS, Toast.hide)
	return model
end

function Toast.flushPending()
	local model = Toast.pending
	Toast.pending = nil
	if model ~= nil then
		Toast.show(model)
	end
	return model
end

local function currentLevel()
	if type(UnitLevel) ~= "function" then
		return nil
	end
	return UnitLevel("player")
end

--- Always fires (if a build is loaded): PLAYER_LEVEL_UP is never
--- ambiguous. Resets the baseline so the very next refresh() does not
--- immediately fire again for the same point.
function Toast.onLevelUp(data, level)
	Toast.baseline = Toast.unspentPoints()
	return Toast.show(Toast.model(data, Follow.build, Talents.readRanks(data), level or currentLevel()))
end

--- Called on every talent-ish event; fires only on a confirmed rise in the
--- unspent count.
function Toast.refresh(data)
	local current = Toast.unspentPoints()
	local fire = Toast.grewSince(Toast.baseline, current)
	Toast.baseline = current
	if not fire then
		return nil
	end
	return Toast.show(Toast.model(data, Follow.build, Talents.readRanks(data), currentLevel()))
end

ns.Toast = Toast
return Toast
