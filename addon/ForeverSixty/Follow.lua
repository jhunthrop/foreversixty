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
local Compat = ns.Compat or require("Compat")

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

--- Whether a slot's saved record holds a build code.
local function hasCode(saved)
	return type(saved) == "table" and saved.code ~= nil
end

--- The one definition of "a saved build" (L3.8 fix round 1): every slot of `entry` (a
--- slotted entry) holding a code, in Follow.SLOTS order, as an iterator of slot, record.
--- Follow.restore's fallback and hasSavedBuild both walk this, so they can never disagree.
local function savedSlots(entry)
	local index = 0
	return function()
		while index < #Follow.SLOTS do
			index = index + 1
			local slot = Follow.SLOTS[index]
			local saved = entry.slots[slot]
			if hasCode(saved) then
				return slot, saved
			end
		end
		return nil
	end
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
---
--- `player` (optional, Codec.loadBuild's) refuses a code for another class; the build
--- dialog passes it, a restore at login does not.
function Follow.load(code, data, name, slot, player)
	local build, message = Codec.loadBuild(code, data, player)
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

--- Where a pasted build goes (L3.7 item 32): the slot a paste is adopted into, and the slot it
--- leaves alone. A paste never overwrites a build restored in the raid slot: while the active
--- slot is the raid one and holds a build, the paste goes to the leveling slot (and becomes the
--- active one); the second answer is then "raid", the slot `/fs build raid` switches back to.
--- Otherwise the paste replaces the active slot's build, as before (the second answer is nil).
Follow.PASTE_SLOT = "leveling"
function Follow.pasteSlot()
	local entry = currentEntry()
	local active = entry.active or Follow.DEFAULT_SLOT
	if active == Follow.DEFAULT_SLOT and hasCode(entry.slots[active]) then
		return Follow.PASTE_SLOT, active
	end
	return active, nil
end

--- The build saved in `slot` for the current character, decoded and named, without making it the
--- active one or touching Follow.build; nil for an empty slot or a code that no longer decodes.
function Follow.slotBuild(data, slot)
	local saved = currentEntry().slots[slot]
	if not hasCode(saved) then
		return nil
	end
	local build = Codec.loadBuild(saved.code, data)
	if build ~= nil then
		build.name = saved.name
	end
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
	if hasCode(active) then
		local build = Follow.load(active.code, data, active.name, entry.active)
		if build ~= nil then
			Follow.noBuildReason = nil
			return build
		end
	end
	-- L3.8: an empty active slot beside a saved build (v1's slot picker could leave one; v2 has
	-- no picker to switch back), or an active one whose code no longer loads, follows the first
	-- other saved build that loads, in SLOTS order, and makes its slot the active one.
	-- Following nothing here also kept the default out (Follow.ensureDefault yields to any
	-- saved build), so the character had no build at all.
	local tried = false
	for slot, record in savedSlots(entry) do
		tried = true
		if slot ~= entry.active then
			local build = Follow.load(record.code, data, record.name, slot)
			if build ~= nil then
				Follow.noBuildReason = nil
				return build
			end
		end
	end
	-- Every saved code failed to load: /fs diag says so (the default stays out all the same).
	Follow.noBuildReason = tried and L.diagBuildSavedRefused or nil
	return nil
end

--- The level a default build switches from leveling to raid.
Follow.RAID_LEVEL = 60

--- Gear loads after this file; looked up when called, never when this file loads.
local function gearModule()
	return ns.Gear or require("Gear")
end

--- True when the current character (its own key only) has a build saved in any slot
--- (savedSlots). Reads saved variables without writing them.
local function hasSavedBuild()
	local follows = type(ForeverSixtyDB) == "table" and ForeverSixtyDB.follows or nil
	local saved = type(follows) == "table" and follows[Export.characterKey()] or nil
	return savedSlots(slotted(saved))() ~= nil
end

--- An FSB1 code for a default build's order (contract 1's {tab, tier, column} triples), or
--- nil for a triple that is not three positive whole numbers.
local function defaultCode(data, classSlug, order)
	local points = {}
	for index, triple in ipairs(type(order) == "table" and order or {}) do
		if type(triple) ~= "table" then
			return nil
		end
		for position = 1, 3 do
			local value = triple[position]
			if type(value) ~= "number" or value < 1 or value % 1 ~= 0 then
				return nil
			end
		end
		points[index] = { tab = triple[1], tier = triple[2], column = triple[3] }
	end
	return Codec.encodeFSB1({ dataBuild = data.build, classSlug = classSlug, order = points })
end

--- A spec slug as a player reads it, its class left off: "hunter-beast-mastery" ->
--- "Beast Mastery", "druid-feral-bear" -> "Feral Bear".
function Follow.specName(spec)
	local tree = type(spec) == "string" and (spec:match("^[^-]+%-(.+)$") or spec) or "?"
	return Talents.titleCase(tree)
end

--- The tree an active build is for, as a spec slug: the build's own `spec` (a data default
--- carries one), else the class tree its plan spends the most points in ("hunter-marksmanship"
--- for a Marksmanship plan, whatever is spent so far); nil when neither says. Comms.buildSpec.
function Follow.buildSpec(data, build)
	if type(build) ~= "table" then
		return nil
	end
	if type(build.spec) == "string" and build.spec ~= "" then
		return build.spec
	end
	local class = type(data) == "table" and type(data.classes) == "table" and data.classes[build.classSlug] or nil
	if type(class) ~= "table" or type(class.tabs) ~= "table" then
		return nil
	end
	local spent = {}
	for _, point in ipairs(build.order or {}) do
		if type(point) == "table" and type(point.tab) == "number" then
			spent[point.tab] = (spent[point.tab] or 0) + 1
		end
	end
	local best, bestTab = 0, nil
	for tab, entry in ipairs(class.tabs) do
		if (spent[tab] or 0) > best and type(entry) == "table" and type(entry.name) == "string" then
			best, bestTab = spent[tab], tab
		end
	end
	return bestTab ~= nil and build.classSlug .. "-" .. Export.slugify(class.tabs[bestTab].name) or nil
end

--- Hotfix 2.0.2: the spec every surface displays and scores in. Whenever a build is followed it is
--- the build's (Follow.buildSpec), so a level-6 paladin on the Retribution path card reads
--- Retribution though no talent point is spent; the talents' leading tree (Gear.specOf over the
--- real ranks, a tie or none spent the first tree) is the fallback only with no build. `classSlug`
--- and `ranks` default to the client's. nil with neither a build nor a resolvable class. Other
--- players' specs (Inspect) stay on their talents.
function Follow.currentSpec(data, classSlug, ranks)
	local spec = Follow.buildSpec(data, Follow.build)
	if spec ~= nil then
		return spec
	end
	classSlug = classSlug or Talents.playerClassSlug()
	if classSlug == nil or type(data) ~= "table" or type(data.classes) ~= "table"
		or type(data.classes[classSlug]) ~= "table" then
		return nil
	end
	local ok, talentSpec = pcall(gearModule().specOf, data, classSlug, ranks or Talents.readRanks(data))
	return ok and talentSpec or nil
end

--- A character with no build saved in any slot follows the data's default for its spec
--- (`data.defaultBuilds[<spec>]`, contract 1): leveling, raid at RAID_LEVEL. The spec is
--- Gear.specOf's, which is the tree with the most points (a tie, or none spent, the first tree;
--- a tree split into cat and bear, the one the points fit). The build is marked
--- `default = true` and named from the data, and nothing is saved: it is worked out again at
--- each login, so a data update replaces it. `character` ({ classSlug, ranks, level, spec })
--- is read from the client for any field it leaves out. `spec`, when given, is the spec the
--- player chose on the chip's "pick a path" card (spec section 5.1) and wins over the points;
--- it must be one of this class's specs with a default. The build, or nil when there is
--- nothing to load (Follow.noBuildReason then says why, for /fs diag). Idempotent: the
--- default already followed for the same spec and kind is kept, never decoded again.
function Follow.ensureDefault(data, character)
	if type(data) ~= "table" or type(data.defaultBuilds) ~= "table" then
		Follow.noBuildReason = L.diagBuildNoDefaults
		return nil
	end
	if hasSavedBuild() then
		-- A saved build is in force, or Follow.restore found none of them loads and said so.
		return nil
	end
	character = character or {}
	local classSlug = character.classSlug or Talents.playerClassSlug()
	if classSlug == nil or type(data.classes) ~= "table" or data.classes[classSlug] == nil then
		Follow.noBuildReason = L.diagBuildNoClass
		return nil
	end
	local level = character.level or Compat.playerLevel()
	local spec = character.spec
	if spec ~= nil and spec:sub(1, #classSlug + 1) ~= classSlug .. "-" then
		Follow.noBuildReason = string.format(L.diagBuildNoSpec, Follow.specName(spec))
		return nil
	end
	if spec == nil then
		spec = gearModule().specOf(data, classSlug, character.ranks or Talents.readRanks(data))
	end
	local builds = spec ~= nil and data.defaultBuilds[spec] or nil
	local kind = (level or 0) >= Follow.RAID_LEVEL and "raid" or "leveling"
	local default = type(builds) == "table" and builds[kind] or nil
	if type(default) ~= "table" then
		Follow.noBuildReason = string.format(L.diagBuildNoSpec, Follow.specName(spec))
		return nil
	end
	local current = Follow.build
	if current ~= nil and current.default and current.spec == spec and current.kind == kind then
		return current
	end
	local code = defaultCode(data, classSlug, default.order)
	local build = code ~= nil and Codec.loadBuild(code, data) or nil
	if build == nil then
		Follow.noBuildReason = string.format(L.diagBuildRefused, Follow.specName(spec))
		return nil
	end
	build.default = true
	build.name = default.name
	build.spec, build.kind = spec, kind
	Follow.build = build
	Follow.noBuildReason = nil
	return build
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

--- Mark an inbox build as handled, by its companion-assigned id or by its code (the build
--- dialog records both, so a build used once stays used after the older `builds` list,
--- which carries the ids, goes away).
function Follow.dismissInbox(id)
	if id == nil then
		return nil
	end
	local entry = currentEntry()
	entry.dismissed = entry.dismissed or {}
	entry.dismissed[id] = true
	return entry.dismissed
end

--- The inbox ids and codes the current character has used or dismissed (Follow.dismissInbox),
--- as a set. The build dialog's "from the site" rows leave these out.
function Follow.dismissedSet()
	return currentEntry().dismissed or {}
end

--- Whether `a` and `b` name the same character, the addon's own key (region/realm/Display Name)
--- or the site's (region/ruleset/name-slug) in either position: Codec.sameCharacter, the one
--- matcher, which also reads a realm name ("Classic Beta PvP") as the ruleset it spells ("pvp").
function Follow.sameCharacter(a, b)
	return Codec.sameCharacter(a, b)
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

--- Builds an officer pushed to this character (L3.7 item 44), newest first, at most
--- Follow.MAX_PUSHED, one per code: each { code, name, from }, in ForeverSixtyDB.pushed[key].
--- Nothing is followed from here; the build dialog lists them as builds to choose from.
Follow.MAX_PUSHED = 5

function Follow.pushedBuilds(key)
	local pushed = type(ForeverSixtyDB) == "table" and type(ForeverSixtyDB.pushed) == "table"
		and ForeverSixtyDB.pushed[key] or nil
	return type(pushed) == "table" and pushed or {}
end

--- Queue a pushed build for the current character; the newest first, a repeated code moved up.
function Follow.addPushed(code, name, from)
	if type(code) ~= "string" or code == "" then
		return nil
	end
	ForeverSixtyDB = ForeverSixtyDB or {}
	if type(ForeverSixtyDB.pushed) ~= "table" then
		ForeverSixtyDB.pushed = {}
	end
	local key = Export.characterKey()
	local list = ForeverSixtyDB.pushed[key]
	if type(list) ~= "table" then
		list = {}
		ForeverSixtyDB.pushed[key] = list
	end
	for index = #list, 1, -1 do
		if list[index].code == code then
			table.remove(list, index)
		end
	end
	table.insert(list, 1, { code = code, name = name, from = from })
	while #list > Follow.MAX_PUSHED do
		table.remove(list)
	end
	return list[1]
end

--- Every build the site sent the character at `key`, once per code: contract 3's `build`
--- messages first (newest first, Follow.messages), then the older top-level `builds` list
--- the companion still writes beside them (Follow.inbox). Each is { id, name, code }; the id
--- is the `builds` entry's own when the same code is there (so a build dismissed before
--- stays dismissed), else the code itself, since a `build` message carries no id. Pure.
function Follow.inboxBuilds(inbox, key)
	local legacy = Follow.inbox(inbox, key)
	local legacyId = {}
	for _, build in ipairs(legacy) do
		if build.id ~= nil and legacyId[build.code] == nil then
			legacyId[build.code] = build.id
		end
	end
	local builds, seen = {}, {}
	local function add(code, name, from)
		if type(code) ~= "string" or code == "" or seen[code] then
			return
		end
		seen[code] = true
		builds[#builds + 1] = { id = legacyId[code] or code, name = name, code = code, from = from }
	end
	-- An officer's push (L3.7 item 44) leads: the newest news.
	for _, pushed in ipairs(Follow.pushedBuilds(key)) do
		add(pushed.code, pushed.name, pushed.from)
	end
	for _, message in ipairs(Follow.messages(inbox, key, "build")) do
		add(message.code, message.name)
	end
	for _, build in ipairs(legacy) do
		add(build.code, build.name)
	end
	return builds
end

--- Messages of `kind` addressed to `key`, or to no one, newest first --
--- Codec.inboxMessages' own character/kind filter, reordered by
--- Codec.newestFirst (Wave C's messages carry no per-message timestamp,
--- so the companion's own append order is the only signal "newest" can
--- mean). Follow.inbox's own contract for builds, extended to the typed
--- inbox messages Codec.inboxMessages reads beside them.
function Follow.messages(inbox, key, kind)
	return Codec.newestFirst(Codec.inboxMessages(inbox, key, kind))
end

local function cellKey(point)
	return point.tab .. ":" .. point.tier .. ":" .. point.column
end

--- Progress owns these numbers (spec section 2.1); it loads after this
--- file, so it is looked up when called, never when this file loads.
local function progressModule()
	return ns.Progress or require("Progress")
end

--- Follow.buildName's worked-out names, per build table (weak keys: a build nothing holds
--- any more drops out), with the data they were worked out against. The chip asks on every
--- refresh; a build's order and the data's tree names never change under it.
local builtNames = setmetatable({}, { __mode = "k" })

--- What to call a build: its own name, else its two biggest trees in the
--- plan, biggest first ("Arms / Fury build", one tree "Protection build"),
--- else its class ("Warrior build") when the trees cannot be named. Every
--- surface that names a build reads this, so none can disagree.
function Follow.buildName(data, build)
	if type(build.name) == "string" and build.name ~= "" then
		return build.name
	end
	local cached = builtNames[build]
	if cached ~= nil and cached.data == data then
		return cached.name
	end
	local split = progressModule().planSplit(build)
	local tabs = {}
	for tab, count in pairs(split) do
		if count > 0 then
			tabs[#tabs + 1] = tab
		end
	end
	table.sort(tabs, function(left, right)
		if split[left] ~= split[right] then
			return split[left] > split[right]
		end
		return left < right
	end)
	local names = {}
	for index = 1, math.min(2, #tabs) do
		local name = Talents.tabName(data, build.classSlug, tabs[index])
		if name == nil then
			names = {}
			break
		end
		names[#names + 1] = name
	end
	local name
	if #names == 0 then
		name = string.format(L.followBuildName, Talents.titleCase(build.classSlug))
	else
		name = string.format(L.followBuildName, table.concat(names, L.followSplitSeparator))
	end
	builtNames[build] = { data = data, name = name }
	return name
end

--- /fs diag's one line for the build in force: "Build: Marksmanship leveling (default for
--- Marksmanship)", "Build: Arms raid (saved)", or "Build: none — <why>" (Follow.noBuildReason,
--- else "not worked out yet").
function Follow.diagLine(data)
	local build = Follow.build
	if build == nil then
		return string.format(L.diagBuildNone, Follow.noBuildReason or L.diagBuildNotYet)
	end
	if build.default then
		return string.format(L.diagBuildDefault, Follow.buildName(data, build), Follow.specName(build.spec))
	end
	return string.format(L.diagBuildSaved, Follow.buildName(data, build))
end

--- The first point in the order the player has not spent yet
--- (Progress.nextPoint).
function Follow.nextPoint(build, ranks)
	return progressModule().nextPoint(build, ranks)
end

--- Every point the character has spent (Progress.pointsSpent).
function Follow.pointsSpent(ranks)
	return progressModule().pointsSpent(ranks)
end

--- How far along the plan the character is (Progress.build).
function Follow.progress(build, ranks)
	return progressModule().build(build, ranks)
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
	-- moved. Naming the cell is more use than naming nothing. Both the
	-- fallback name and the tier below are shown to the player 1-indexed
	-- (Talents.displayTier, round-2 section 10 ruling 4) -- point.tier itself
	-- stays the raw, zero-based Data.lua value everywhere else in this
	-- file (cellKey, nextPoint's own lookups).
	local name = talent and talent.name
		or string.format(L.followUnknownCell, Talents.displayTier(point.tier), point.column)
	return string.format(L.followNext, name, tab, Talents.displayTier(point.tier))
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
