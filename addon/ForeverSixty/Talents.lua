-- addon/ForeverSixty/Talents.lua
-- The talent window, read into the shape the export string wants.
--
-- The client is asked only for tier, column and rank. The *order* the export
-- encodes comes from Data.lua, never from GetTalentInfo's own index order:
-- nothing guarantees the client indexes a tab's talents the way the site's
-- talents/<class>.json lists them, and a mismatch would decode every rank
-- onto the wrong talent. Data.lua is the order; the client supplies the
-- numbers to fill it.
local _, ns = ...
ns = type(ns) == "table" and ns or {}

local Talents = {}

--- Spike check 4. Classic's talent window has three tabs per class.
Talents.TAB_COUNT = 3

local function cellKey(tier, column)
	return tier .. ":" .. column
end

Talents.cellKey = cellKey

--- Every rank the client reports, keyed by tab then "<tier>:<column>".
function Talents.readRanks()
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

--- Ranks per tree, in Data.lua's talent order, ready for Codec.encodeFS1.
--- A talent the client reported no cell for is 0, never nil: a nil in the
--- middle of the array would truncate the encoded tree and shift every rank
--- after it.
function Talents.treeRanks(data, classSlug)
	local class = data and data.classes and data.classes[classSlug]
	if class == nil then
		return nil
	end
	local fromClient = Talents.readRanks()
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
