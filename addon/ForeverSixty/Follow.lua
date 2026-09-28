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

--- Named build slots (design section 1: "Raid", "Leveling", "PvP"), in
--- display order. A slot id is one of these three strings; DEFAULT_SLOT is
--- where a pre-slot single build (the old ForeverSixtyDB.follows[key] =
--- {code=, name=} shape, or the still-older account-wide
--- ForeverSixtyDB.follow) lands once this version migrates it -- "slot 1"
--- in the order above.
Follow.SLOTS = { "raid", "leveling", "pvp" }
Follow.DEFAULT_SLOT = Follow.SLOTS[1]

--- Bring one character's ForeverSixtyDB.follows[key] entry to the slotted
--- shape ({ slots = { raid = {code=,name=}, ... }, active = <slot> },
--- plus an optional `dismissed` set of inbox ids), in place where the
--- entry is already a table. A pre-slot entry ({code=, name=}) becomes
--- DEFAULT_SLOT; a missing entry becomes an empty one with nothing
--- active yet.
local function slotted(entry)
	if type(entry) ~= "table" then
		return { slots = {}, active = Follow.DEFAULT_SLOT }
	end
	if type(entry.slots) == "table" then
		entry.active = entry.active or Follow.DEFAULT_SLOT
		return entry
	end
	local slots = {}
	if entry.code ~= nil then
		slots[Follow.DEFAULT_SLOT] = { code = entry.code, name = entry.name }
	end
	return { slots = slots, active = Follow.DEFAULT_SLOT }
end

--- The current character's slotted entry, read from (and, on first touch,
--- written back into) ForeverSixtyDB.follows -- so a caller that only
--- reads (Follow.slotsFor) never needs to know about the pre-slot shape.
local function currentEntry()
	ForeverSixtyDB = ForeverSixtyDB or {}
	ForeverSixtyDB.follows = ForeverSixtyDB.follows or {}
	local key = Export.characterKey()
	local entry = slotted(ForeverSixtyDB.follows[key])
	ForeverSixtyDB.follows[key] = entry
	return entry
end

--- Load a build from a code into `slot` (the entry's active slot when
--- `slot` is nil, DEFAULT_SLOT for a character with no entry yet). `name`
--- is what to call it: neither code format carries one (controller ruling
--- 2), so it comes from the companion's inbox entry, or from nowhere for
--- a pasted code.
---
--- Saved per character (ForeverSixtyDB.follows[Export.characterKey()]):
--- the account-wide ForeverSixtyDB.follow this replaced showed the first
--- character's build to every character that logged in after it.
function Follow.load(code, data, name, slot)
	local build, message = Codec.loadBuild(code, data)
	if build == nil then
		return nil, message
	end
	build.name = name
	local entry = currentEntry()
	slot = slot or entry.active or Follow.DEFAULT_SLOT
	entry.slots[slot] = { code = code, name = name }
	entry.active = slot
	Follow.build = build
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
--- A second migration, in the same pass: a character whose entry is still
--- the pre-slot {code=,name=} shape (this version's own predecessor)
--- adopts it into DEFAULT_SLOT, the shape `slotted` above always produces.
function Follow.restore(data)
	if type(ForeverSixtyDB) ~= "table" then
		return nil
	end
	local key = Export.characterKey()
	local saved = type(ForeverSixtyDB.follows) == "table" and ForeverSixtyDB.follows[key] or nil
	if saved == nil and type(ForeverSixtyDB.follow) == "table" then
		saved = { code = ForeverSixtyDB.follow.code, name = ForeverSixtyDB.follow.name }
		ForeverSixtyDB.follow = nil
	end
	local entry = slotted(saved)
	ForeverSixtyDB.follows = ForeverSixtyDB.follows or {}
	ForeverSixtyDB.follows[key] = entry
	local active = entry.slots[entry.active]
	if type(active) ~= "table" or active.code == nil then
		return nil
	end
	return Follow.load(active.code, data, active.name, entry.active)
end

--- A slot id's display name ("raid" -> L.slotRaid), the one lookup rule
--- both Follow.slotsFor and the Talents page's slot picker use, so a slot
--- id and its Locale key can never drift into two different rules for
--- turning one into the other.
function Follow.slotName(slot)
	return L["slot" .. slot:sub(1, 1):upper() .. slot:sub(2)] or slot
end

--- Every named slot for the current character: its id, its localized
--- name, whether it holds a build, and whether it is the active one. The
--- Talents page's slot picker draws straight from this.
function Follow.slotsFor()
	local entry = currentEntry()
	local list = {}
	for index, slot in ipairs(Follow.SLOTS) do
		list[index] = {
			id = slot,
			name = Follow.slotName(slot),
			hasBuild = entry.slots[slot] ~= nil,
			active = entry.active == slot,
		}
	end
	return list
end

--- Switch the active slot for the current character and load whatever it
--- holds (nil when the slot is empty -- switching to an empty slot is not
--- a refusal, it is "nothing loaded here yet").
function Follow.setActiveSlot(data, slot)
	local entry = currentEntry()
	entry.active = slot
	local saved = entry.slots[slot]
	if type(saved) ~= "table" or saved.code == nil then
		Follow.build = nil
		return nil
	end
	local build = Codec.loadBuild(saved.code, data)
	if build == nil then
		Follow.build = nil
		return nil
	end
	build.name = saved.name
	Follow.build = build
	return build
end

--- Forget the build in `slot` (the active slot when `slot` is nil). When
--- that empties every slot the character has, the entry itself is
--- dropped -- the shape a single-slot player (everyone before named
--- slots) already expects from Follow.forget().
function Follow.forget(slot)
	local key = Export.characterKey()
	local entry = currentEntry()
	slot = slot or entry.active
	entry.slots[slot] = nil
	if slot == entry.active then
		Follow.build = nil
	end
	if next(entry.slots) == nil then
		ForeverSixtyDB.follows[key] = nil
	end
	return nil
end

--- Mark an inbox build (by its companion-assigned id) as handled, so the
--- Overview's "build arrived" banner never shows it again -- whether the
--- player loaded it or dismissed it outright, both close the news item.
function Follow.dismissInbox(id)
	if id == nil then
		return nil
	end
	local entry = currentEntry()
	entry.dismissed = entry.dismissed or {}
	entry.dismissed[id] = true
	return entry.dismissed
end

--- The first inbox build addressed to this character (Follow.inbox's own
--- rule) that has not been dismissed yet, or nil. This is the "build
--- arrived" banner's model source: a build the player already loaded or
--- closed the banner for never comes back.
function Follow.pendingArrival(inbox)
	local usable = Follow.inbox(inbox, Export.characterKey())
	local entry = currentEntry()
	local dismissed = entry.dismissed or {}
	for _, build in ipairs(usable) do
		if build.id ~= nil and not dismissed[build.id] then
			return build
		end
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

--- How many distinct talent cells `a` and `b`'s orders disagree on (either
--- side wants a different number of points there, including zero). Used
--- for the "build arrived" banner's diff summary -- a count, not a full
--- point-by-point comparison, is enough to tell a player "this is a
--- retune" from "this is a different build".
function Follow.orderDiffCount(a, b)
	local function counts(build)
		local result = {}
		for _, point in ipairs(build.order) do
			local key = cellKey(point)
			result[key] = (result[key] or 0) + 1
		end
		return result
	end
	local countsA, countsB = counts(a), counts(b)
	local seen, diff = {}, 0
	for key, count in pairs(countsA) do
		seen[key] = true
		if countsB[key] ~= count then
			diff = diff + 1
		end
	end
	for key in pairs(countsB) do
		if not seen[key] then
			diff = diff + 1
		end
	end
	return diff
end

--- The "build arrived" banner's one-line diff summary for a build waiting
--- in the inbox, against `currentBuild` (the active slot's build, or nil).
--- Decodes `code` to compare it but never loads it -- that stays the
--- player's own "Load it" click.
function Follow.arrivalSummary(data, currentBuild, code)
	local build = Codec.loadBuild(code, data)
	if build == nil then
		return nil
	end
	if currentBuild == nil then
		return L.buildArrivedFirst
	end
	local diff = Follow.orderDiffCount(currentBuild, build)
	if diff == 0 then
		return L.buildArrivedSame
	end
	return string.format(L.buildArrivedDiff, diff)
end

ns.Follow = Follow
return Follow
