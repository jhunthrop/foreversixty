-- addon/ForeverSixty/Talents.lua
-- The talent window, read into the shape the export string wants.
--
-- The client is asked only for a rank per talent. The *order* the export
-- encodes comes from Data.lua, never from the client's own index order:
-- nothing guarantees the client indexes a tab's talents the way the site's
-- talents/<class>.json lists them, and a mismatch would decode every rank
-- onto the wrong talent. Data.lua is the order; the client supplies the
-- numbers to fill it.
--
-- Two clients, one shape. The 1.60 Forever client keeps its trees in the
-- modern trait system: there is no GetTalentInfo, and a rank is read with
-- C_Traits.GetNodeInfo(configID, node) using the TraitNode id Data.lua
-- carries per talent. The classic GetTalentInfo window is kept as the
-- fallback for a client that still has it. Either way the result is
-- ranks[tab]["<tier>:<column>"], which is what Follow and Gear key on.
local _, ns = ...
ns = type(ns) == "table" and ns or {}

local Talents = {}

--- Spike check 4. Classic's talent window has three tabs per class.
Talents.TAB_COUNT = 3

local function cellKey(tier, column)
	return tier .. ":" .. column
end

Talents.cellKey = cellKey

--- The player's class as a data.classes key: UnitClass's second return is
--- the locale-neutral token, its first the display name (see Export).
function Talents.playerClassSlug()
	local classToken = select(2, UnitClass("player"))
	return classToken and classToken:lower() or nil
end

--- True when this client exposes the trait system the 1.60 trees live in.
local function hasTraitApi()
	return type(C_Traits) == "table"
		and type(C_Traits.GetNodeInfo) == "function"
		and type(C_ClassTalents) == "table"
		and type(C_ClassTalents.GetActiveConfigID) == "function"
end

--- A node's committed rank and its rank with any staged (clicked, not yet
--- applied) changes, from C_Traits.GetNodeInfo's answer. Pure. The
--- committed rank is activeRank (ranksPurchased on a client without it);
--- the staged one is currentRank, else ranksPurchased, else the committed.
--- Which fields this client fills is a spike question: /fs diag verbose
--- prints a sample (Talents.nodeSample).
function Talents.nodeRanks(info)
	if type(info) ~= "table" then
		return 0, 0
	end
	local committed = info.activeRank or info.ranksPurchased or 0
	local staged = info.currentRank or info.ranksPurchased or committed
	return committed, staged
end

--- A node's committed and staged ranks from the trait system. A character
--- below the talent level has no active config yet, and a node the client
--- does not know is unranked: both read as 0, never as an error. The first
--- answer is kept for /fs diag verbose (Talents.nodeSample).
local function traitRank(configID, node)
	if configID == nil or node == nil then
		return 0, 0
	end
	local info = C_Traits.GetNodeInfo(configID, node)
	if Talents.nodeSample == nil and type(info) == "table" then
		Talents.nodeSample = info
	end
	return Talents.nodeRanks(info)
end

--- A spell's texture, from wherever this client keeps the function.
local function spellTexture(spellID)
	if spellID == nil then
		return nil
	end
	if type(C_Spell) == "table" and type(C_Spell.GetSpellTexture) == "function" then
		return C_Spell.GetSpellTexture(spellID)
	end
	if type(GetSpellTexture) == "function" then
		return GetSpellTexture(spellID)
	end
	return nil
end

--- node -> entry -> definition -> spell -> texture. Raises on any client
--- function that is missing; iconFor turns that into nil.
local function traitIcon(node)
	local configID = C_ClassTalents.GetActiveConfigID()
	local info = C_Traits.GetNodeInfo(configID, node)
	local entryID = info.activeEntry and info.activeEntry.entryID or info.entryIDs[1]
	local entry = C_Traits.GetEntryInfo(configID, entryID)
	local definition = C_Traits.GetDefinitionInfo(entry.definitionID)
	return definition.overrideIcon or spellTexture(definition.overriddenSpellID or definition.spellID)
end

local iconCache = {}

--- The icon for a talent from Data.lua (its `node` is the client's trait
--- node id), or nil. The data file carries no icons, so this asks the
--- client, and every link of that chain may be absent on a client nobody
--- here can run: a failure is "no icon", never an error. A found icon is
--- kept for the session; a miss is asked again, because a character below
--- the talent level has no trait config yet and will have one later.
function Talents.iconFor(talent)
	local node = type(talent) == "table" and talent.node or nil
	if node == nil or not hasTraitApi() then
		return nil
	end
	if iconCache[node] ~= nil then
		return iconCache[node]
	end
	local ok, icon = pcall(traitIcon, node)
	if ok and icon ~= nil then
		iconCache[node] = icon
		return icon
	end
	return nil
end

local function traitRanks(data, staged)
	local class = data and data.classes and data.classes[Talents.playerClassSlug()]
	if class == nil then
		return {}, false
	end
	local configID = C_ClassTalents.GetActiveConfigID()
	local byTab = {}
	local differs = false
	for tabIndex, tab in ipairs(class.tabs) do
		local ranks = {}
		for _, talent in ipairs(tab.talents) do
			local committed, withStaged = traitRank(configID, talent.node)
			ranks[cellKey(talent.tier, talent.column)] = staged and withStaged or committed
			if withStaged ~= committed then
				differs = true
			end
		end
		byTab[tabIndex] = ranks
	end
	return byTab, staged and (differs or Talents.configHasStagedChanges(configID))
end

--- Whether the client says the config has clicks not yet applied; false
--- where it cannot say.
function Talents.configHasStagedChanges(configID)
	if configID == nil or type(C_Traits.ConfigHasStagedChanges) ~= "function" then
		return false
	end
	local ok, staged = pcall(C_Traits.ConfigHasStagedChanges, configID)
	return ok and staged == true
end

local function classicRanks()
	local byTab = {}
	for tab = 1, GetNumTalentTabs() or 0 do
		local ranks = {}
		for index = 1, GetNumTalents(tab) or 0 do
			local _, _, tier, column, rank = GetTalentInfo(tab, index)
			if tier ~= nil and column ~= nil then
				ranks[cellKey(tier, column)] = rank or 0
			end
		end
		byTab[tab] = ranks
	end
	return byTab
end

--- Every rank the client reports, keyed by tab then "<tier>:<column>".
--- `data` is Data.lua (or a spec's stand-in); the trait path needs its
--- node ids, the classic path ignores it. A client with neither talent
--- API reports no ranks rather than raising.
function Talents.readRanks(data)
	if hasTraitApi() then
		return (traitRanks(data, false))
	end
	if type(GetNumTalentTabs) == "function" then
		return classicRanks()
	end
	return {}
end

--- Talents.readRanks with the talent window's staged clicks counted, and
--- whether any are staged: the talent marks read this so the strip answers
--- a click before Apply Changes. The classic window has no staging.
function Talents.readStagedRanks(data)
	if hasTraitApi() then
		return traitRanks(data, true)
	end
	return Talents.readRanks(data), false
end

--- Ranks per tree, in Data.lua's talent order, ready for Codec.encodeFS1.
--- A talent the client reported no cell for is 0, never nil: a nil in the
--- middle of the array would truncate the encoded tree and shift every rank
--- after it.
function Talents.treeRanks(data, classSlug)
	local class = data and data.classes and data.classes[classSlug]
	if class == nil then
		return nil
	end
	local fromClient = Talents.readRanks(data)
	local trees = {}
	for tabIndex, tab in ipairs(class.tabs) do
		local ranks = fromClient[tabIndex] or {}
		local tree = {}
		for position, talent in ipairs(tab.talents) do
			tree[position] = ranks[cellKey(talent.tier, talent.column)] or 0
		end
		trees[tabIndex] = tree
	end
	return trees
end

--- The talent at one cell, for Follow's "next point" line.
function Talents.cellOf(data, classSlug, tab, tier, column)
	local class = data and data.classes and data.classes[classSlug]
	if class == nil then
		return nil
	end
	local entry = class.tabs[tab]
	if entry == nil then
		return nil
	end
	for _, talent in ipairs(entry.talents) do
		if talent.tier == tier and talent.column == column then
			return talent
		end
	end
	return nil
end

--- The points a character has in each tab, from Talents.readRanks' shape
--- (ranks[tab]["<tier>:<column>"]): { [tab] = points }, tabs with none left
--- out. Pure.
function Talents.tabPoints(ranks)
	local points = {}
	for tab, cells in pairs(type(ranks) == "table" and ranks or {}) do
		local total = 0
		for _, rank in pairs(type(cells) == "table" and cells or {}) do
			total = total + (tonumber(rank) or 0)
		end
		if total > 0 then
			points[tab] = total
		end
	end
	return points
end

--- How far ahead of every other tree the leading tree must be before the
--- addon names a character's spec on its own (design/specs/2026-11-08-addon-ux.md
--- section 6, owner 2026-11-11: "a wrong silent guess is worse than one question").
Talents.SURE_LEAD = 5
--- The fewest spent points that can name a tree at all: a level 10 with one
--- point is asked (the owner's words, 2026-11-11).
Talents.SURE_MIN_POINTS = 2

--- The one rule for reading a spec off a character's points (`points`,
--- Talents.tabPoints' shape). The tab the points name beyond doubt: at
--- least SURE_MIN_POINTS spent and every one of them in one tree, or the
--- leading tree holding more than half of all spent points and leading
--- every other tree by at least SURE_LEAD. Anything else, zero or one
--- point included, is nil: the player is asked. The second answer is the
--- tabs that hold points, in tab order, for the question's wording. Pure.
--- Every surface that derives a spec from points (Gear.specOf,
--- Follow.ensureDefault, Comms.ownSpec, the guild tab's own row, Inspect)
--- reads this and nothing else.
function Talents.decideTree(points)
	local spent, trees = 0, {}
	for tab, total in pairs(type(points) == "table" and points or {}) do
		if type(total) == "number" and total > 0 then
			spent = spent + total
			trees[#trees + 1] = tab
		end
	end
	table.sort(trees)
	if spent < Talents.SURE_MIN_POINTS then
		return nil, trees
	end
	local lead, second = nil, 0
	for _, tab in ipairs(trees) do
		local total = points[tab]
		if lead == nil or total > points[lead] then
			second = lead ~= nil and points[lead] or second
			lead = tab
		elseif total > second then
			second = total
		end
	end
	if second == 0 then
		return lead, trees
	end
	local top = points[lead]
	if top * 2 > spent and top - second >= Talents.SURE_LEAD then
		return lead, trees
	end
	return nil, trees
end

--- The tab name Data.lua carries, for the "next point" line's tree name.
function Talents.tabName(data, classSlug, tab)
	local class = data and data.classes and data.classes[classSlug]
	local entry = class and class.tabs[tab]
	return entry and entry.name or nil
end

--- Hyphen-separated words, each capitalised: "warrior-fury" -> "Warrior
--- Fury", "paladin" -> "Paladin". Window.specLabel strips a spec slug's
--- class half before calling this; Follow's build-name fallback (round-2
--- section 10 ruling 3) passes a bare class slug straight through. Kept here,
--- the one place both call it, rather than copied a third time -- this
--- file loads first in the TOC (before Window, before every view), so
--- nothing that needs it has a load-order problem reaching it.
function Talents.titleCase(text)
	local words = {}
	for word in (text or ""):gmatch("[^-]+") do
		words[#words + 1] = word:sub(1, 1):upper() .. word:sub(2)
	end
	return table.concat(words, " ")
end

--- The tier number the way the real talent window labels its rows --
--- Tier 1 through 8, never Data.lua's own zero-based `tier` field.
--- Round-2 section 10 ruling 4: every lookup against Data.lua (cellKey, tab
--- indices, the codec) stays zero-based; only a string a player reads
--- (the chip, the Talents page's own unknown-cell name)
--- goes through this first.
function Talents.displayTier(tier)
	-- Data.lua's tiers are already the rows a player counts (Improved
	-- Heroic Strike is tier 1 and sits on the top row of the game's own
	-- window); adding one here showed "tier 2" for it (seen in game).
	return tier or 1
end

ns.Talents = Talents
return Talents
