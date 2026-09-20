-- addon/ForeverSixty/Follow.lua
-- The next point in a build the player loaded, and the highlight on it.
--
-- The build carries an order -- one cell per point, in the order the build
-- spends them. The next point is the first entry in that order whose running
-- count exceeds what the player already has in that cell. Counting rather
-- than comparing ranks is what makes an overspent talent (three ranks where
-- the build wanted two) read as "already done" instead of as a gap the
-- addon keeps pointing at.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Codec = ns.Codec or require("Codec")
local Talents = ns.Talents or require("Talents")

local Follow = {}

function Follow.load(code, data)
	local build, message = Codec.loadBuild(code, data)
	if build == nil then
		return nil, message
	end
	Follow.build = build
	return build
end

local function cellKey(point)
	return point.tab .. ":" .. point.tier .. ":" .. point.column
end

--- The first point in the order the player has not spent yet.
--- `ranks` is Talents.readRanks()'s shape: ranks[tab]["<tier>:<column>"].
function Follow.nextPoint(build, ranks)
	if build == nil then
		return nil
	end
	local wanted = {}
	for index, point in ipairs(build.order) do
		local key = cellKey(point)
		wanted[key] = (wanted[key] or 0) + 1
		local have = (ranks[point.tab] or {})[Talents.cellKey(point.tier, point.column)] or 0
		if wanted[key] > have then
			return { tab = point.tab, tier = point.tier, column = point.column, index = index }
		end
	end
	return nil
end

function Follow.line(data, build, ranks)
	if build == nil then
		return L.followNone
	end
	local point = Follow.nextPoint(build, ranks)
	if point == nil then
		return L.followDone
	end
	local talent = Talents.cellOf(data, build.classSlug, point.tab, point.tier, point.column)
	local tab = Talents.tabName(data, build.classSlug, point.tab) or ""
	-- A cell Data.lua has no talent for is a code from a build whose trees
	-- moved. Naming the cell is more use than naming nothing.
	local name = talent and talent.name or string.format(L.followUnknownCell, point.tier, point.column)
	return string.format(L.followNext, name, tab, point.tier)
end

--- The frame. Created on first use, never on load: an addon that builds
--- frames at login costs every player that time whether they use it or not.
function Follow.refresh(data)
	local ranks = Talents.readRanks()
	Follow.frame = Follow.frame or CreateFrame("Frame", "ForeverSixtyFollowFrame", UIParent)
	Follow.text = Follow.text or Follow.frame:CreateFontString()
	Follow.text:SetText(Follow.line(data, Follow.build, ranks))
	Follow.frame:Show()
	Follow.highlight(data, ranks)
	return Follow.text:GetText()
end

--- Highlight the next point's button in the talent window, when it is open.
--- The frame name differs between clients (spike check 15); when neither is
--- present the line above is the whole feature and nothing errors.
Follow.TALENT_FRAME = { "PlayerTalentFrame", "TalentFrame" }

--- The returned table describing what was highlighted (or nil) is not used
--- by any caller here -- it exists so the spec can assert on the outcome
--- without a display server. Keep returning it.
function Follow.highlight(_data, ranks)
	local point = Follow.nextPoint(Follow.build, ranks or Talents.readRanks())
	if point == nil then
		return nil
	end
	for _, name in ipairs(Follow.TALENT_FRAME) do
		local frame = _G[name]
		if frame ~= nil then
			Follow.highlighted = { frame = name, point = point }
			return Follow.highlighted
		end
	end
	return nil
end

ns.Follow = Follow
return Follow
