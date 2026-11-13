-- addon/ForeverSixty/Sources.lua
-- The gear source filter (design/specs/2026-11-13-addon-source-filter.md), the twin of the web
-- BiS page's Sources group (web/src/lib/bis/source-filter.ts, source-choice.ts, page-state.ts and
-- doll-view.ts's cappedSlot): its eight kinds in its order, its default, its words and its re-pick.
-- Filtering re-picks, never hides: with a kind off, a slot shows the best pick from the kinds still
-- on. Gear.pick resolves every pick through Sources.resolve, so the Upgrades rows, the slot marks,
-- the chip and the tooltips follow one rule.
--
-- The choice is per character: Prefs' gearSources.byCharacter[key].on, a list of kind ids, or no
-- entry. No entry is the site's default for the band looked at (below 60 every kind but
-- Reputation, standings up to Friendly still counting; at 60 every kind). A saved list is literal.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Prefs = ns.Prefs or require("Prefs")

local Sources = {}

--- The filter row's order (source-filter.ts's SOURCE_FILTER_KINDS).
Sources.KINDS = { "dungeon", "raid", "world", "quest", "crafted", "rep", "vendor", "pvp" }

--- Data.lua's source kind letters (pipeline.addonbis.SOURCE_KIND_CODES) -> filter kind: the
--- world-drop pool's B is a World / BoE item, as on the site.
Sources.KIND_OF_CODE = { D = "dungeon", A = "raid", W = "world", B = "world", Q = "quest", C = "crafted",
	R = "rep", V = "vendor", P = "pvp" }

--- Each kind's word (source-choice.ts's sourceKindLabel), as a Locale key.
Sources.LABEL_KEYS = { dungeon = "sourcesKindDungeon", raid = "sourcesKindRaid", world = "sourcesKindWorld",
	quest = "sourcesKindQuest", crafted = "sourcesKindCrafted", rep = "sourcesKindRep",
	vendor = "sourcesKindVendor", pvp = "sourcesKindPvp" }

--- The standing letters of the packed tail (pipeline.addonlua.STANDING_CODES), their order, and
--- each standing's word as a Locale key.
Sources.STANDING_OF_CODE = { n = "neutral", f = "friendly", h = "honored", r = "revered", e = "exalted" }
local STANDING_RANK = { neutral = 0, friendly = 1, honored = 2, revered = 3, exalted = 4 }
Sources.STANDING_KEYS = { neutral = "sourcesStandingNeutral", friendly = "sourcesStandingFriendly",
	honored = "sourcesStandingHonored", revered = "sourcesStandingRevered", exalted = "sourcesStandingExalted" }

--- The leveling default keeps reputation items up to this standing (page-state.ts's defaultSources,
--- source-choice.ts's repCap).
Sources.REP_CAP = "friendly"
--- At this band and up the default is every kind (page-state.ts's MAX_BAND).
Sources.MAX_BAND = 60
--- The kinds the site's Pre-raid view turns off (pre-raid.ts's PRE_RAID_OFF).
Sources.PRE_RAID_OFF = { raid = true, pvp = true }
--- The second slot of each pair and its first (source-filter.ts's PAIR_FIRST): the second slot steers
--- clear of what the first resolved to, and no off hand sits beside a two-handed main hand.
Sources.PAIR_FIRST = { finger2 = "finger1", trinket2 = "trinket1", off_hand = "main_hand" }

local KIND_INDEX = {}
for index, kind in ipairs(Sources.KINDS) do
	KIND_INDEX[kind] = index
end

--- A kind's word ("World / BoE").
function Sources.label(kind)
	local key = Sources.LABEL_KEYS[kind]
	return key ~= nil and L[key] or kind
end

--- A standing's word ("Friendly"), nil for none.
function Sources.standingWord(standing)
	local key = Sources.STANDING_KEYS[standing or ""]
	return key ~= nil and L[key] or nil
end

--- The known kinds of `list`, once each, in the filter's order. Pure.
function Sources.ordered(list)
	local seen = {}
	for _, kind in ipairs(type(list) == "table" and list or {}) do
		if KIND_INDEX[kind] ~= nil then
			seen[kind] = true
		end
	end
	local out = {}
	for _, kind in ipairs(Sources.KINDS) do
		if seen[kind] then
			out[#out + 1] = kind
		end
	end
	return out
end

--- The kinds on with no saved choice (page-state.ts's defaultSources): below MAX_BAND every kind but
--- Reputation, at it every kind. Pure.
function Sources.defaultOn(band)
	local out = {}
	for _, kind in ipairs(Sources.KINDS) do
		if kind ~= "rep" or (type(band) == "number" and band >= Sources.MAX_BAND) then
			out[#out + 1] = kind
		end
	end
	return out
end

local function isLevelingDefault(on, band)
	if not (type(band) == "number" and band < Sources.MAX_BAND) or #on ~= #Sources.KINDS - 1 then
		return false
	end
	for _, kind in ipairs(on) do
		if kind == "rep" then
			return false
		end
	end
	return true
end

local function isPreRaid(on, band)
	if not (type(band) == "number" and band >= Sources.MAX_BAND) then
		return false
	end
	local wanted = 0
	for _, kind in ipairs(Sources.KINDS) do
		if not Sources.PRE_RAID_OFF[kind] then
			wanted = wanted + 1
		end
	end
	if #on ~= wanted then
		return false
	end
	for _, kind in ipairs(on) do
		if Sources.PRE_RAID_OFF[kind] then
			return false
		end
	end
	return true
end

--- A filter state from the kinds on and the band (source-choice.ts's sourceChoiceFor): { on (the
--- kinds, in order), set (kind -> true, Reputation included under the cap), repCap ("friendly" for
--- the leveling default only), all (every kind on: the published picks), levelingDefault, preRaid,
--- saved, id (a key for caches) }. Pure.
function Sources.stateOf(on, band, saved)
	on = Sources.ordered(on)
	local state = { on = on, set = {}, saved = saved == true }
	for _, kind in ipairs(on) do
		state.set[kind] = true
	end
	state.all = #on == #Sources.KINDS
	state.levelingDefault = isLevelingDefault(on, band)
	state.preRaid = isPreRaid(on, band)
	if state.levelingDefault then
		state.repCap = Sources.REP_CAP
		state.set.rep = true
	end
	state.id = table.concat(on, ",") .. (state.repCap ~= nil and "|" .. state.repCap or "")
	return state
end

--- Every kind on, as the published picks are ranked: the state for another player's gear (the group
--- panel, the loot council, the hello's gap), never filtered by this character's choice.
Sources.ALL = Sources.stateOf(Sources.KINDS, Sources.MAX_BAND)

--- This character's key (Export.characterKey), read late (Export loads before this file); nil where
--- the client cannot say yet, which reads as no saved choice: the band's default.
function Sources.characterKey()
	local Export = ns.Export or require("Export")
	local ok, key = pcall(Export.characterKey)
	return ok and key or nil
end

local function sameCharacter(a, b)
	local Follow = ns.Follow or require("Follow")
	return Follow.sameCharacter(a, b)
end

--- The saved entry of `key` ({ on }), or nil; `create` makes it.
local function entryOf(key, create)
	return Prefs.characterEntry(Prefs.current().gearSources.byCharacter, key, sameCharacter, create)
end

--- The saved kinds of `key` (this character's by default), in order, or nil for none saved.
function Sources.saved(key)
	local entry = entryOf(key or Sources.characterKey(), false)
	if type(entry) ~= "table" or type(entry.on) ~= "table" then
		return nil
	end
	return Sources.ordered(entry.on)
end

--- Bumped on every change, so a state cached under an older choice is never served.
Sources.version = 0
local cache = { version = -1, key = nil, states = {} }

--- The state for `band` (Sources.stateOf over the saved kinds, else the band's default) of `key`
--- (this character's by default). Cached until the choice or the character changes.
function Sources.state(band, key)
	key = key or Sources.characterKey()
	local defaulted = not (type(band) == "number" and band >= Sources.MAX_BAND)
	-- The saved table itself is part of the key: a SavedVariables reload (or a spec's fresh one) is a
	-- new table, and a state read from the old one is never served.
	local byCharacter = Prefs.current().gearSources.byCharacter
	if cache.version ~= Sources.version or cache.key ~= key or cache.byCharacter ~= byCharacter then
		cache = { version = Sources.version, key = key, byCharacter = byCharacter, states = {} }
	end
	local slot = defaulted and "below" or "cap"
	local state = cache.states[slot]
	if state == nil then
		local saved = Sources.saved(key)
		state = Sources.stateOf(saved or Sources.defaultOn(band), band, saved ~= nil)
		cache.states[slot] = state
	end
	return state
end

--- What runs after the choice changes (the panel, the flyout, the Settings group and every surface
--- that shows a pick redraw), by name: Sources.listen replaces a listener of the same name, so a
--- second registration never runs it twice.
Sources.listeners = {}

function Sources.listen(name, listener)
	Sources.listeners[name] = listener
end

local function changed()
	Sources.version = Sources.version + 1
	for _, listener in pairs(Sources.listeners) do
		listener()
	end
end

--- Save `on` as this character's choice, literally (never rewritten at a level-up).
function Sources.choose(on, key)
	local entry = entryOf(key or Sources.characterKey(), true)
	if entry == nil then
		return nil
	end
	entry.on = Sources.ordered(on)
	changed()
	return entry.on
end

--- Turn one kind on or off from what `band` shows now, and save the result.
function Sources.toggle(kind, band, key)
	local next, found = {}, false
	for _, current in ipairs(Sources.state(band, key).on) do
		if current == kind then
			found = true
		else
			next[#next + 1] = current
		end
	end
	if not found then
		next[#next + 1] = kind
	end
	return Sources.choose(next, key)
end

--- Clear this character's choice: the site's default again ("Default").
function Sources.reset(key)
	local byCharacter = Prefs.current().gearSources.byCharacter
	key = key or Sources.characterKey()
	for name, entry in pairs(byCharacter) do
		if type(entry) == "table" and (name == key or sameCharacter(name, key)) then
			byCharacter[name] = nil
		end
	end
	changed()
end

-- The words ----------------------------------------------------------------------------------------

--- The state in words (source-choice.ts's sourcesFieldLabel): "All sources", "No rep above Friendly",
--- "Pre-raid", "Dungeon only", "6 of 8 sources", "No sources". Pure.
function Sources.words(state)
	local count = #state.on
	if state.all then
		return L.sourcesAll
	elseif state.levelingDefault then
		return L.sourcesLevelingDefault
	elseif state.preRaid then
		return L.sourcesPreRaid
	elseif count == 1 then
		return string.format(L.sourcesOne, Sources.label(state.on[1]))
	elseif count == 0 then
		return L.sourcesNone
	end
	return string.format(L.sourcesSome, count, #Sources.KINDS)
end

--- The hover on the words: "Off: Reputation, PvP" (the leveling default's own line), nil at All
--- sources. Pure.
function Sources.offLine(state)
	if state.all then
		return nil
	elseif state.levelingDefault then
		return L.sourcesOffLevelingTip
	end
	local off = {}
	for _, kind in ipairs(Sources.KINDS) do
		if not state.set[kind] then
			off[#off + 1] = Sources.label(kind)
		end
	end
	return string.format(L.sourcesOffTip, table.concat(off, L.sourcesListSeparator))
end

-- The re-pick --------------------------------------------------------------------------------------

local function aboveCap(standing)
	return (STANDING_RANK[standing or ""] or math.huge) > STANDING_RANK[Sources.REP_CAP]
end

--- Whether an item of filter `kind` asking `standing` counts under `state` (doll-view.ts's capped: under
--- the cap a Reputation item with no known standing is out; any other kind only by a standing it names).
function Sources.allows(state, kind, standing)
	if not state.set[kind] then
		return false
	end
	if state.repCap == nil then
		return true
	end
	if kind == "rep" then
		return not aboveCap(standing)
	end
	return standing == nil or not aboveCap(standing)
end

--- A decoded pick's own filter kinds (source-filter.ts's pickKinds): its head's, else its code's.
function Sources.pickKinds(pick)
	local head = pick.head
	if head ~= nil and #head.kinds > 0 then
		return head.kinds
	end
	local kind = Sources.KIND_OF_CODE[pick.kind or ""]
	return kind ~= nil and { kind } or {}
end

local function blocks(notWith, mateId)
	if mateId == nil or notWith == nil then
		return false
	end
	for _, id in ipairs(notWith) do
		if id == mateId then
			return true
		end
	end
	return false
end

local function pickAllowed(pick, state)
	local standings = pick.head ~= nil and pick.head.standings or {}
	local kinds = Sources.pickKinds(pick)
	if #kinds == 0 then
		-- A kind the filter does not know is never filtered out: nothing to judge it by.
		return true
	end
	for _, kind in ipairs(kinds) do
		if Sources.allows(state, kind, standings[kind]) then
			return true
		end
	end
	return false
end

local function entryAllowed(entry, state, mateId)
	return Sources.allows(state, entry.filterKind, entry.standing) and not blocks(entry.notWith, mateId)
end

--- Re-picked answers, per entry and state: a hover never builds the same table twice. Weak keys:
--- an entry nothing holds any more drops out with its answers.
local repicked = setmetatable({}, { __mode = "k" })
local kept = setmetatable({}, { __mode = "k" })

--- The by-kind entries after `from` that `state` allows (no mate rule: the Shift block's list).
local function entriesAfter(pick, state, from, mateId)
	local out = {}
	for index = from + 1, #pick.byKind do
		local entry = pick.byKind[index]
		if #out >= 2 then
			break
		end
		if entryAllowed(entry, state, mateId) then
			out[#out + 1] = { item_id = entry.item_id, kind = entry.kind, source = entry.source }
		end
	end
	return out
end

--- A by-kind entry as a pick (Gear.decodePick's shape): `unfiltered` is the published pick it stands
--- in for, `alternatives` the next entries the state allows.
local function asPick(pick, index, state, mateId)
	local entry = pick.byKind[index]
	local byState = repicked[entry]
	if byState == nil then
		byState = {}
		repicked[entry] = byState
	end
	local cacheKey = state.id .. "#" .. tostring(mateId or "")
	if byState[cacheKey] == nil then
		byState[cacheKey] = {
			entry.item_id, entry.kind, entry.source,
			item_id = entry.item_id, kind = entry.kind, source = entry.source, standing = entry.standing,
			twoHand = entry.twoHand, alternatives = entriesAfter(pick, state, index, mateId),
			unfiltered = pick, byKind = {},
		}
	end
	return byState[cacheKey]
end

--- The published pick with only the alternatives `state` allows (an alternative's kind is its code's;
--- under the cap a Reputation alternative, whose standing the data does not carry, is out). The pick
--- itself when it keeps them all.
local function withAllowedAlternatives(pick, state)
	local alternatives, dropped = {}, false
	for _, alternative in ipairs(pick.alternatives or {}) do
		local kind = Sources.KIND_OF_CODE[alternative.kind or ""]
		if kind ~= nil and Sources.allows(state, kind, nil) then
			alternatives[#alternatives + 1] = alternative
		else
			dropped = true
		end
	end
	if not dropped then
		return pick
	end
	local byState = kept[pick]
	if byState == nil then
		byState = {}
		kept[pick] = byState
	end
	if byState[state.id] == nil then
		local copy = {}
		for field, value in pairs(pick) do
			copy[field] = value
		end
		copy.alternatives = alternatives
		byState[state.id] = copy
	end
	return byState[state.id]
end

--- One slot's pick under `state` (source-filter.ts's filteredPick): the published pick while one of
--- its kinds is on and its pair mate does not rule it out, else the best-ranked by-kind entry the state
--- allows, else nil ("No <slot> from these sources at this level"). `mate` is the pair's first slot's
--- own answer for a second slot (nil when it has none). Pure but for its caches.
function Sources.resolve(pick, state, slot, mate)
	if pick == nil then
		return nil
	end
	if state.all then
		return pick
	end
	if slot == "off_hand" and mate ~= nil and mate.twoHand then
		return nil
	end
	local mateId = mate ~= nil and mate.item_id or nil
	local notWith = pick.head ~= nil and pick.head.notWith or nil
	if pickAllowed(pick, state) and not blocks(notWith, mateId) then
		return withAllowedAlternatives(pick, state)
	end
	for index, entry in ipairs(pick.byKind or {}) do
		if entryAllowed(entry, state, mateId) then
			return asPick(pick, index, state, mateId)
		end
	end
	return nil
end

--- Gear slots with an option from each kind at a band (doll-view.ts's sourceCounts): `picks` is the
--- band's published picks by slot. Independent of what is on. Pure.
function Sources.counts(picks)
	local counts = {}
	for _, kind in ipairs(Sources.KINDS) do
		counts[kind] = 0
	end
	for _, pick in pairs(picks or {}) do
		local kinds = {}
		for _, kind in ipairs(Sources.pickKinds(pick)) do
			kinds[kind] = true
		end
		for _, entry in ipairs(pick.byKind or {}) do
			kinds[entry.filterKind] = true
		end
		for _, alternative in ipairs(pick.alternatives or {}) do
			local kind = Sources.KIND_OF_CODE[alternative.kind or ""]
			if kind ~= nil then
				kinds[kind] = true
			end
		end
		for kind in pairs(kinds) do
			counts[kind] = counts[kind] + 1
		end
	end
	return counts
end

--- What the Sources flyout and the Settings group show for a state and the band's counts (Sources.counts, nil when the band has
--- none): { rows = { { kind, label, checked, count, tip } }, showDefault }. A count reads "-" with no
--- counts; at the leveling default the Reputation row is off and says what still counts. Pure.
function Sources.choiceRows(state, counts)
	local rows = {}
	for index, kind in ipairs(Sources.KINDS) do
		local count = counts ~= nil and tostring(counts[kind] or 0) or L.sourcesCountNone
		rows[index] = { kind = kind, label = Sources.label(kind), checked = state.set[kind] == true
			and not (kind == "rep" and state.repCap ~= nil), count = count,
			tip = kind == "rep" and state.repCap ~= nil and L.sourcesRepTip or nil }
	end
	return { rows = rows, showDefault = state.saved == true }
end

ns.Sources = Sources
return Sources
