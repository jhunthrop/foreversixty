-- addon/ForeverSixty/RaidPrep.lua
-- What a raid leader checks before the pull, read from the client for the export (owner
-- decision D14): resistance totals, the consumables in the bags, the keyring and the
-- attunement quests and items. Codec.encodeFS1 writes them as the school=, supplies=,
-- locks=, quests= and owned= sections; the site names the consumables and judges the readiness.
-- RaidPrep.summary reads the same for the hub's Guild tab: what tonight's raid asks of the player.
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

--- Where each of `ids` is: "worn", "bags" (the carried bags) or nil (not held), by id; nil without a
--- container API (an empty read would prove nothing). Worn wins over the bags.
function RaidPrep.locations(ids, bags, container)
	if not RaidPrep.hasContainers() then
		return nil
	end
	local wanted = {}
	for _, id in ipairs(ids) do
		wanted[id] = true
	end
	local where = {}
	for _, bag in ipairs(bags) do
		for slot = 1, container.size(bag) do
			local id = itemOf(container.link(bag, slot))
			if id ~= nil and wanted[id] then
				where[id] = "bags"
			end
		end
	end
	for slot = 1, RaidPrep.EQUIPPED_SLOTS do
		local id = itemOf(GetInventoryItemLink("player", slot))
		if id ~= nil and wanted[id] then
			where[id] = "worn"
		end
	end
	return where
end

--- Each of `ids` and whether it is worn or in `bags`: { { itemId, held } } in `ids`' order, nil
--- without a container API (an empty read would prove nothing).
function RaidPrep.held(ids, bags, container)
	local where = RaidPrep.locations(ids, bags, container)
	if where == nil then
		return nil
	end
	local out = {}
	for index, id in ipairs(ids) do
		out[index] = { itemId = id, held = where[id] ~= nil }
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

-- What tonight's raid asks of the player (the hub's Guild tab, Your prep; spec 2026-11-13-addon-hub 4.5)

--- The raids whose bosses Data.encounterNeeds lists, by zone name (the site's `raid` and the client's
--- instance name), each with its bosses' npc ids. The data carries the needs per npc id only
--- (contract 1), so the zone-to-boss map is the addon's own; the Officer tab's readiness reads it too.
RaidPrep.BOSSES = {
	["Molten Core"] = { 12118, 11982, 12259, 12057, 12056, 12264, 12098, 11988, 12018, 11502 },
	["Onyxia's Lair"] = { 10184 },
	["Blackwing Lair"] = { 12435, 13020, 12017, 11983, 14601, 11981, 14020, 11583 },
}

--- The boss a carried item is for, by npc id: only the bosses whose need names an item (Nefarian's
--- Shadow Flame asks for the Onyxia Scale Cloak).
RaidPrep.BOSS_NAMES = { [11583] = "Nefarian" }

--- The item a raid's door asks the player to carry, by zone: the Drakefire Amulet for Onyxia's Lair.
RaidPrep.ZONE_ITEMS = { ["Onyxia's Lair"] = RaidPrep.ATTUNEMENT_ITEMS[1] }

--- The supplies every raid asks for (the site's raid-prep check): a flask or a raid elixir, and
--- buff food; each need is met by any one of its kinds (Data.consumables' kinds).
RaidPrep.RAID_SUPPLIES = { { "flask", "elixir" }, { "food" } }

--- The potion that answers a protection kind an encounter names: the greater potion of its school.
RaidPrep.PROTECTION_ITEMS = { fire_protection = 13457, frost_protection = 13456, nature_protection = 13458,
	shadow_protection = 13459, arcane_protection = 13461 }

--- RaidPrep.resistances' order, by school name (encounterNeeds' `resist`).
RaidPrep.SCHOOLS = { "fire", "nature", "frost", "shadow", "arcane" }

local function schoolIndex(school)
	for index, name in ipairs(RaidPrep.SCHOOLS) do
		if name == school then
			return index
		end
	end
	return nil
end

--- The zone's needs from the data: { items = { { itemId, npc } }, resists = { school }, kinds = { kind } },
--- each once, in the bosses' order. Pure.
function RaidPrep.needsOf(data, zone)
	local all = type(data) == "table" and type(data.encounterNeeds) == "table" and data.encounterNeeds or {}
	local needs, seen = { items = {}, resists = {}, kinds = {} }, {}
	local function once(list, key, value)
		if not seen[key] then
			seen[key] = true
			list[#list + 1] = value
		end
	end
	for _, npc in ipairs(RaidPrep.BOSSES[zone] or {}) do
		local entry = type(all[npc]) == "table" and all[npc] or {}
		if type(entry.item) == "number" then
			once(needs.items, "item:" .. entry.item, { itemId = entry.item, npc = npc })
		end
		if type(entry.resist) == "string" and schoolIndex(entry.resist) ~= nil then
			once(needs.resists, "resist:" .. entry.resist, entry.resist)
		end
		for _, kind in ipairs(type(entry.consumables) == "table" and entry.consumables or {}) do
			once(needs.kinds, "kind:" .. kind, kind)
		end
	end
	return needs
end

--- Your prep for a raid in `zone`, from what the client read. `input`: { data, zone, where (item id ->
--- "worn" or "bags", RaidPrep.locations; nil unread), resistances (RaidPrep.resistances'), counts (kind ->
--- count of Data.consumables' kinds in the bags; nil unread) }. Rows in order, each
--- { kind = "door" | "item" | "resist" | "carried" | "missing" | "packed", ... }: the door's item and
--- where it is (`state` "worn", "bags" or "missing"), each encounter item the same with the boss's
--- `npc`, each resistance with its `value`, the supplies carried (`counts`: { { kind, count } }), then
--- the missing supplies (`needs`: { { kinds } or { itemId } }) or, with every need met, "packed".
--- Rows the client could not read are left out. Pure.
function RaidPrep.prep(input)
	local zone, rows = input.zone, {}
	local where = input.where
	local doorItem = RaidPrep.ZONE_ITEMS[zone]
	if doorItem ~= nil and where ~= nil then
		rows[#rows + 1] = { kind = "door", itemId = doorItem, state = where[doorItem] or "missing" }
	end
	local needs = RaidPrep.needsOf(input.data, zone)
	if where ~= nil then
		for _, item in ipairs(needs.items) do
			rows[#rows + 1] = { kind = "item", itemId = item.itemId, npc = item.npc, state = where[item.itemId] or "missing" }
		end
	end
	if type(input.resistances) == "table" then
		for _, school in ipairs(needs.resists) do
			rows[#rows + 1] = { kind = "resist", school = school, value = input.resistances[schoolIndex(school)] or 0 }
		end
	end
	local counts = input.counts
	if type(counts) == "table" and RaidPrep.BOSSES[zone] ~= nil then
		local missing = {}
		for _, kinds in ipairs(RaidPrep.RAID_SUPPLIES) do
			local have = 0
			for _, kind in ipairs(kinds) do
				have = have + (counts[kind] or 0)
			end
			if have == 0 then
				missing[#missing + 1] = { kinds = kinds }
			end
		end
		for _, kind in ipairs(needs.kinds) do
			if (counts[kind] or 0) == 0 then
				missing[#missing + 1] = { kinds = { kind }, itemId = RaidPrep.PROTECTION_ITEMS[kind] }
			end
		end
		local carried = {}
		for _, kinds in ipairs(RaidPrep.RAID_SUPPLIES) do
			for _, kind in ipairs(kinds) do
				if (counts[kind] or 0) > 0 then
					carried[#carried + 1] = { kind = kind, count = counts[kind] }
				end
			end
		end
		for _, kind in ipairs(needs.kinds) do
			if (counts[kind] or 0) > 0 then
				carried[#carried + 1] = { kind = kind, count = counts[kind] }
			end
		end
		if #carried > 0 then
			rows[#rows + 1] = { kind = "carried", counts = carried }
		end
		rows[#rows + 1] = #missing > 0 and { kind = "missing", needs = missing } or { kind = "packed" }
	end
	return rows
end

--- Counts by kind of `stacks` (RaidPrep.supplies') that `catalogue` (Data.consumables, item id -> kind)
--- names; nil for nil stacks. Pure.
function RaidPrep.kindCounts(stacks, catalogue)
	if stacks == nil then
		return nil
	end
	local counts = {}
	for _, stack in ipairs(stacks) do
		local kind = type(catalogue) == "table" and catalogue[stack.itemId] or nil
		if kind ~= nil then
			counts[kind] = (counts[kind] or 0) + (tonumber(stack.count) or 0)
		end
	end
	return counts
end

--- Your prep for a raid in `zone`, read now (RaidPrep.prep's rows): the door's and the encounters'
--- items worn or in `bags`, the resistance totals, the supplies counted by kind. `container` is
--- Export's { size, link }.
function RaidPrep.summary(zone, data, bags, container)
	local needs = RaidPrep.needsOf(data, zone)
	local ids = {}
	if RaidPrep.ZONE_ITEMS[zone] ~= nil then
		ids[#ids + 1] = RaidPrep.ZONE_ITEMS[zone]
	end
	for _, item in ipairs(needs.items) do
		ids[#ids + 1] = item.itemId
	end
	local catalogue = type(data) == "table" and data.consumables or nil
	return RaidPrep.prep({
		data = data, zone = zone, where = RaidPrep.locations(ids, bags, container),
		resistances = RaidPrep.resistances(),
		counts = RaidPrep.kindCounts(RaidPrep.supplies(bags, container), catalogue),
	})
end

ns.RaidPrep = RaidPrep
return RaidPrep
