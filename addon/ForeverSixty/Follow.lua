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

--- Load a build from a code. `name` is what to call it: neither code
--- format carries one (controller ruling 2), so it comes from the
--- companion's inbox entry, or from nowhere for a pasted code.
function Follow.load(code, data, name)
	local build, message = Codec.loadBuild(code, data)
	if build == nil then
		return nil, message
	end
	build.name = name
	Follow.build = build
	ForeverSixtyDB = ForeverSixtyDB or {}
	ForeverSixtyDB.follow = { code = code, name = name }
	return build
end

--- The build the player had loaded last session, if it still decodes.
--- A code that no longer decodes (the addon's data moved on) is dropped
--- rather than reported: nothing asked for it this session.
function Follow.restore(data)
	local saved = type(ForeverSixtyDB) == "table" and ForeverSixtyDB.follow or nil
	if type(saved) ~= "table" or saved.code == nil then
		return nil
	end
	return Follow.load(saved.code, data, saved.name)
end

function Follow.forget()
	Follow.build = nil
	if type(ForeverSixtyDB) == "table" then
		ForeverSixtyDB.follow = nil
	end
	return nil
end

--- The companion's builds that have a code, in the order it wrote them.
--- Pure: the inbox is the companion's file and this never writes it.
--- An entry with no `code` is a truncated or hand-edited file, not
--- something the companion produces, and is skipped rather than counted.
function Follow.inbox(inbox)
	local usable = {}
	if type(inbox) ~= "table" or type(inbox.builds) ~= "table" then
		return usable
	end
	for _, build in ipairs(inbox.builds) do
		if build.code ~= nil then
			usable[#usable + 1] = { id = build.id, name = build.name, code = build.code }
		end
	end
	return usable
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
	local ranks = Talents.readRanks(data)
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
function Follow.highlight(data, ranks)
	local point = Follow.nextPoint(Follow.build, ranks or Talents.readRanks(data))
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
