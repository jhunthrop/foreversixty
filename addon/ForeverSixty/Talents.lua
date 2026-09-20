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

--- A rank from the trait system. A character below the talent level has no
--- active config yet, and a node the client does not know is unranked:
--- both read as 0, never as an error.
local function traitRank(configID, node)
	if configID == nil or node == nil then
		return 0
	end
	local info = C_Traits.GetNodeInfo(configID, node)
	return info and (info.activeRank or info.ranksPurchased) or 0
end

local function traitRanks(data)
	local class = data and data.classes and data.classes[Talents.playerClassSlug()]
	if class == nil then
		return {}
	end
	local configID = C_ClassTalents.GetActiveConfigID()
	local byTab = {}
	for tabIndex, tab in ipairs(class.tabs) do
		local ranks = {}
		for _, talent in ipairs(tab.talents) do
			ranks[cellKey(talent.tier, talent.column)] = traitRank(configID, talent.node)
		end
		byTab[tabIndex] = ranks
	end
	return byTab
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
		return traitRanks(data)
	end
	if type(GetNumTalentTabs) == "function" then
		return classicRanks()
	end
	return {}
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

--- The tab name Data.lua carries, for the "next point" line's tree name.
function Talents.tabName(data, classSlug, tab)
	local class = data and data.classes and data.classes[classSlug]
	local entry = class and class.tabs[tab]
	return entry and entry.name or nil
end

ns.Talents = Talents
return Talents
