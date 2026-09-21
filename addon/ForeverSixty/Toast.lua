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

ns.Toast = Toast
return Toast
