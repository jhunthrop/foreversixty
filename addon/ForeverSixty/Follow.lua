-- addon/ForeverSixty/Follow.lua
-- The next point in a build the player loaded.
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
local Export = ns.Export or require("Export")

local Follow = {}

--- Load a build from a code. `name` is what to call it: neither code
--- format carries one (controller ruling 2), so it comes from the
--- companion's inbox entry, or from nowhere for a pasted code.
---
--- Saved per character (ForeverSixtyDB.follows[Export.characterKey()]):
--- the account-wide ForeverSixtyDB.follow this replaced showed the first
--- character's build to every character that logged in after it.
function Follow.load(code, data, name)
	local build, message = Codec.loadBuild(code, data)
	if build == nil then
		return nil, message
	end
	build.name = name
	Follow.build = build
	ForeverSixtyDB = ForeverSixtyDB or {}
	ForeverSixtyDB.follows = ForeverSixtyDB.follows or {}
	ForeverSixtyDB.follows[Export.characterKey()] = { code = code, name = name }
	return build
end

--- The build the player had loaded last session, if it still decodes.
--- A code that no longer decodes (the addon's data moved on) is dropped
--- rather than reported: nothing asked for it this session.
---
--- Migration ("the addon isn't character aware"): the first character to
--- restore after this version adopts the old account-wide
--- ForeverSixtyDB.follow, if this character has no entry of its own yet,
--- and the old key is then deleted -- so exactly one character keeps the
--- build the account already had, and no later character inherits it too.
function Follow.restore(data)
	if type(ForeverSixtyDB) ~= "table" then
		return nil
	end
	local key = Export.characterKey()
	local saved = type(ForeverSixtyDB.follows) == "table" and ForeverSixtyDB.follows[key] or nil
	if saved == nil and type(ForeverSixtyDB.follow) == "table" then
		saved = ForeverSixtyDB.follow
		ForeverSixtyDB.follow = nil
	end
	if type(saved) ~= "table" or saved.code == nil then
		return nil
	end
	local build = Follow.load(saved.code, data, saved.name)
	return build
end

function Follow.forget()
	Follow.build = nil
	if type(ForeverSixtyDB) == "table" and type(ForeverSixtyDB.follows) == "table" then
		ForeverSixtyDB.follows[Export.characterKey()] = nil
	end
	return nil
end

--- Fold one "/"-separated key segment to a form both the addon's own key
--- format and the site's agree on: lower-cased, with runs of whitespace
--- turned into a single hyphen (how the site slugifies a display name).
local function normaliseSegment(segment)
	return (segment or ""):lower():gsub("%s+", "-")
end

--- Whether `a` and `b` name the same character. Either may be the addon's
--- own key (region/realm/Display Name, Export.characterKey) or the
--- site's (region/ruleset/slugified-name, e.g. "us/pvp/bow-jackzon",
--- companion/internal/addon/addon.go's Build.Character): each of the
--- three "/"-separated segments is compared through normaliseSegment, so
--- "US/Ashbringer/Bow Jackzon" and "us/ashbringer/bow-jackzon" match.
--- The realm/ruleset segment gets no further translation -- on this
--- client the realm name doubles as the ruleset already (Export.save's
--- own `ruleset = realm`), so a straight case-fold is the whole rule.
function Follow.sameCharacter(a, b)
	if type(a) ~= "string" or type(b) ~= "string" then
		return false
	end
	local aRegion, aMid, aName = a:match("^([^/]*)/([^/]*)/(.*)$")
	local bRegion, bMid, bName = b:match("^([^/]*)/([^/]*)/(.*)$")
	if aRegion == nil or bRegion == nil then
		return false
	end
	return normaliseSegment(aRegion) == normaliseSegment(bRegion)
		and normaliseSegment(aMid) == normaliseSegment(bMid)
		and normaliseSegment(aName) == normaliseSegment(bName)
end

--- The companion's builds usable by the character at `key`, in the order
--- the companion wrote them. Pure: the inbox is the companion's file and
--- this never writes it. An entry with no `code` is a truncated or
--- hand-edited file, not something the companion produces, and is
--- skipped rather than counted. An entry with no `character` is
--- site-wide and always kept; one addressed to another character
--- (Follow.sameCharacter says no) is dropped rather than offered.
function Follow.inbox(inbox, key)
	local usable = {}
	if type(inbox) ~= "table" or type(inbox.builds) ~= "table" then
		return usable
	end
	for _, build in ipairs(inbox.builds) do
		local addressed = build.character
		if build.code ~= nil and (addressed == nil or Follow.sameCharacter(addressed, key)) then
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

ns.Follow = Follow
return Follow
