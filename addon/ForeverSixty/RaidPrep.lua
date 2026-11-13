-- addon/ForeverSixty/RaidPrep.lua
-- What a raid leader checks before the pull, read from the client for the export (owner
-- decision D14): resistance totals, the consumables in the bags, the keyring and the
-- attunement quests and items. Codec.encodeFS1 writes them as the school=, supplies=,
-- locks=, quests= and owned= sections; the site names the consumables and judges the readiness.
--
-- Each reader answers nil when this client has no API for it, and Export then writes no
-- section: an absent section reads "not tracked" on the site, never "missing". A reader the
-- client can answer returns its list even when empty, because "read, and nothing there" is a
-- real answer (no flask in the bags).
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local Compat = ns.Compat or require("Compat")

local RaidPrep = {}

--- UnitResistance's school indices, in the school= section's order: fire, nature, frost,
--- shadow, arcane (0 is armor and 1 holy, which nothing reads). Spike check 98.
RaidPrep.RESISTANCE_SCHOOLS = { 2, 3, 4, 5, 6 }

--- The attunement quests the export asks the client about, each the last quest of its chain.
--- Onyxia's Lair: Drakefire Amulet (6502, Alliance: Stormwind Rendezvous 6402, The Great
--- Masquerade 6403, The Dragon's Eye 6501, then 6502) and Blood of the Black Dragon Champion
--- (6602, Horde, after Ascension... 6601). Every id
--- is a row of the 1.60.1.70790 client's QuestV2 table (wago.tools db2/QuestV2, which carries no
--- titles); titles and chain order are cmangos classic-db quest_template's (PrevQuestId).
--- Spike check 101 confirms them in game.
RaidPrep.ATTUNEMENT_QUESTS = { 6502, 6602 }

--- The items an attunement chain hands out that the export says are held or not, worn or in
--- the carried bags: the Drakefire Amulet (16309), which both Onyxia chains reward.
RaidPrep.ATTUNEMENT_ITEMS = { 16309 }

--- The item class GetItemInfoInstant reports for food, drink, potions, elixirs and flasks
--- (Enum.ItemClass.Consumable). Every stack of it in the carried bags is exported with its
--- count, and the site's catalogue names what counts as a flask, an elixir or buff food, so
--- a new Forever consumable needs a site change, never an addon release.
RaidPrep.CONSUMABLE_CLASS = 0

--- The item class of keys (Enum.ItemClass.Key): the Skeleton Key and the Key to the City.
RaidPrep.KEY_CLASS = 13

--- The keyring's bag id: the client's KEYRING_CONTAINER (-2 on Classic), else the Enum.
--- nil on a client with neither: the keys in the bags are then all locks= carries.
function RaidPrep.keyringBag()
	if type(KEYRING_CONTAINER) == "number" then
		return KEYRING_CONTAINER
	end
	local bags = type(Enum) == "table" and Enum.BagIndex or nil
	return type(bags) == "table" and bags.Keyring or nil
end

--- The resistance totals { fire, nature, frost, shadow, arcane } (UnitResistance's second
--- return, base plus gear and buffs), or nil when the client has no UnitResistance.
function RaidPrep.resistances()
	if type(UnitResistance) ~= "function" then
		return nil
	end
	local totals = {}
	for index, school in ipairs(RaidPrep.RESISTANCE_SCHOOLS) do
		local _, total = UnitResistance("player", school)
		if type(total) ~= "number" then
			return nil
		end
		totals[index] = math.max(0, math.floor(total))
	end
	return totals
end

--- How many items sit in one bag slot: C_Container.GetContainerItemInfo's stackCount,
--- else the legacy global's second return, else 1 (the slot holds something).
function RaidPrep.stackCount(bag, slot)
	if C_Container and C_Container.GetContainerItemInfo then
		local info = C_Container.GetContainerItemInfo(bag, slot)
		if type(info) == "table" and type(info.stackCount) == "number" then
			return info.stackCount
		end
	elseif GetContainerItemInfo then
		local _, count = GetContainerItemInfo(bag, slot)
		if type(count) == "number" then
			return count
		end
	end
	return 1
end

--- True when the client has a container API at all: without one an empty read proves nothing.
function RaidPrep.hasContainers()
	return (C_Container ~= nil and C_Container.GetContainerNumSlots ~= nil) or GetContainerNumSlots ~= nil
end

--- A bag item's id and item class, read from the link; nil for an empty slot or a link the
--- client cannot parse (the id then falls back to the link text, the class stays unknown).
local function itemOf(link)
	if link == nil then
		return nil
	end
	local id, _, _, _, _, classId = Compat.itemInfoInstant(link)
	return tonumber(id) or tonumber(link:match("item:(%d+)")), classId
end

--- `ids` (a set) as a sorted list.
local function sortedKeys(set)
	local out = {}
	for id in pairs(set) do
		out[#out + 1] = id
	end
	table.sort(out)
	return out
end

--- Every Consumable-class stack in `bags`, counts summed per item, sorted by id:
--- { { itemId, count } }. `container` is Export's { size = fn(bag), link = fn(bag, slot) }.
--- nil without a container API.
function RaidPrep.supplies(bags, container)
	if not RaidPrep.hasContainers() then
		return nil
	end
	local counts = {}
	for _, bag in ipairs(bags) do
		for slot = 1, container.size(bag) do
			local id, classId = itemOf(container.link(bag, slot))
			if id ~= nil and classId == RaidPrep.CONSUMABLE_CLASS then
				counts[id] = (counts[id] or 0) + RaidPrep.stackCount(bag, slot)
			end
		end
	end
	local out = {}
	for _, id in ipairs(sortedKeys(counts)) do
		out[#out + 1] = { itemId = id, count = counts[id] }
	end
	return out
end

--- The keys the character carries: every item on the keyring plus every Key-class item in
--- `bags` (a full keyring spills keys into the bags, and a client without a keyring keeps
--- them there), sorted and once each. Bags hold only gear in bags=, so a key there is seen
--- only here. nil without a container API.
function RaidPrep.keyring(bags, container)
	if not RaidPrep.hasContainers() then
		return nil
	end
	local ids = {}
	local keyring = RaidPrep.keyringBag()
	if keyring ~= nil then
		for slot = 1, container.size(keyring) do
			local id = itemOf(container.link(keyring, slot))
			if id ~= nil then
				ids[id] = true
			end
		end
	end
	for _, bag in ipairs(bags) do
		for slot = 1, container.size(bag) do
			local id, classId = itemOf(container.link(bag, slot))
			if id ~= nil and classId == RaidPrep.KEY_CLASS then
				ids[id] = true
			end
		end
	end
	return sortedKeys(ids)
end

--- The client's quest-completed check: C_QuestLog's, else the old global; nil without one.
local function questCheck()
	if type(C_QuestLog) == "table" and type(C_QuestLog.IsQuestFlaggedCompleted) == "function" then
		return C_QuestLog.IsQuestFlaggedCompleted
	end
	if type(IsQuestFlaggedCompleted) == "function" then
		return IsQuestFlaggedCompleted
	end
	return nil
end

--- Each attunement quest and whether the client flags it done: { { questId, done } }, nil
--- when the client has no quest-completed check.
function RaidPrep.quests()
	local done = questCheck()
	if done == nil then
		return nil
	end
	local out = {}
	for index, id in ipairs(RaidPrep.ATTUNEMENT_QUESTS) do
		out[index] = { questId = id, done = done(id) == true }
	end
	return out
end

--- Each of `ids` and whether it is worn or in `bags`: { { itemId, held } } in `ids`' order, nil
--- without a container API (an empty read would prove nothing).
function RaidPrep.held(ids, bags, container)
	if not RaidPrep.hasContainers() then
		return nil
	end
	local held = {}
	for slot = 1, RaidPrep.EQUIPPED_SLOTS do
		local id = itemOf(GetInventoryItemLink("player", slot))
		if id ~= nil then
			held[id] = true
		end
	end
	for _, bag in ipairs(bags) do
		for slot = 1, container.size(bag) do
			local id = itemOf(container.link(bag, slot))
			if id ~= nil then
				held[id] = true
			end
		end
	end
	local out = {}
	for index, id in ipairs(ids) do
		out[index] = { itemId = id, held = held[id] == true }
	end
	return out
end

--- Each attunement item and whether it is worn or in `bags`: RaidPrep.held of ATTUNEMENT_ITEMS.
function RaidPrep.owned(bags, container)
	return RaidPrep.held(RaidPrep.ATTUNEMENT_ITEMS, bags, container)
end

--- The inventory slots GetInventoryItemLink answers for worn gear (1 head to 19 tabard).
RaidPrep.EQUIPPED_SLOTS = 19

--- Every raid-prep field Codec.encodeFS1 takes, read now.
function RaidPrep.read(bags, container)
	return {
		owned = RaidPrep.owned(bags, container),
		resistances = RaidPrep.resistances(),
		supplies = RaidPrep.supplies(bags, container),
		keyring = RaidPrep.keyring(bags, container),
		quests = RaidPrep.quests(),
	}
end

ns.RaidPrep = RaidPrep
return RaidPrep
