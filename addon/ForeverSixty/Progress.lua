-- addon/ForeverSixty/Progress.lua
-- One function per number (spec section 2.1): every count, delta, verdict
-- and gap the addon shows is computed here, once, and every surface reads
-- the same answer -- so the talent window, the chip and the
-- paper doll can never disagree with each other or with the character.
--
-- Pure over plain tables: a build, Talents.readRanks()'s ranks, the worn
-- items keyed by slot ({ itemId, link, equipLoc, stats }) and the bag items.
-- Only a bag link handed in as a string is read from the client, through Compat.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Talents = ns.Talents or require("Talents")
local Export = ns.Export or require("Export")
local Compat = ns.Compat or require("Compat")
local Gear = ns.Gear or require("Gear")
local Rules = ns.Rules or require("Rules")

local Progress = {}

local function cellKey(point)
	return point.tab .. ":" .. point.tier .. ":" .. point.column
end

--- The first point in the order the player has not spent yet.
--- `ranks` is Talents.readRanks()'s shape: ranks[tab]["<tier>:<column>"].
--- Counting rather than comparing ranks is what makes an overspent talent
--- (three ranks where the build wanted two) read as "already done".
function Progress.nextPoint(build, ranks)
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

--- Every point the character has spent, whatever the plan says.
function Progress.pointsSpent(ranks)
	local spent = 0
	for _, tab in pairs(ranks or {}) do
		for _, rank in pairs(tab) do
			spent = spent + (tonumber(rank) or 0)
		end
	end
	return spent
end

--- The level a point in the order is taken at: talents start at 10, one
--- per level, so point N is level 9 + N.
function Progress.levelForPoint(index)
	if index == nil then
		return nil
	end
	local level = 9 + index
	return { level = level, text = string.format(L.progressTakeAt, level) }
end

--- How far along the plan the character is: `matched` is every point
--- the plan wanted that the character has (a rank over the plan's own
--- counts only up to the plan), `total` the plan's length, `offPlan`
--- the points spent where the plan wanted none, `point` the next one to
--- take and `nextLevel` the level it is taken at. `done` is true once a
--- build has no point left (seen in game: a "2 of 51" beside eleven spent
--- points, because the old count stopped at the first unmet point).
function Progress.build(build, ranks)
	if build == nil then
		return { matched = 0, total = 0, offPlan = 0, point = nil, nextLevel = nil, done = false }
	end
	ranks = ranks or {}
	local wanted = {}
	for _, point in ipairs(build.order) do
		local key = cellKey(point)
		wanted[key] = {
			tab = point.tab, tier = point.tier, column = point.column,
			want = (wanted[key] and wanted[key].want or 0) + 1,
		}
	end
	local matched = 0
	for _, cell in pairs(wanted) do
		local have = (ranks[cell.tab] or {})[Talents.cellKey(cell.tier, cell.column)] or 0
		matched = matched + math.min(have, cell.want)
	end
	local point = Progress.nextPoint(build, ranks)
	return {
		matched = matched,
		total = #build.order,
		offPlan = math.max(0, Progress.pointsSpent(ranks) - matched),
		point = point,
		nextLevel = point ~= nil and (9 + point.index) or nil,
		done = point == nil,
	}
end

local function hasAnyStat(stats)
	return type(stats) == "table" and next(stats) ~= nil
end

local function inUnit(delta, dpsPerPoint)
	if delta ~= nil and dpsPerPoint ~= nil then
		return delta * dpsPerPoint
	end
	return delta
end

--- One slot's verdict: the planned item against what is worn there.
--- `plan` is a build.gear entry ({ itemId, stats }), `worn` an equipped
--- entry ({ itemId, link, stats }); either may be nil, not both (nil then).
---
--- - the same item id on both sides is "match", whatever the two stat
---   lines say (seen in game: a helm reading "+10 over itself");
--- - a planned item with nothing worn is "empty", short by the plan's
---   whole score, so an empty ring is the biggest gap;
--- - a worn item the plan never named is "noplan";
--- - otherwise "better" when worn minus planned is above zero, else
---   "behind" (also when it cannot be scored).
---
--- `delta` is worn minus planned, nil when the planned side carries no
--- stats (an FS1 code) or there are no weights; it is in DPS when
--- `dpsPerPoint` (DPS per 1.0 of the reference stat) is given, and
--- `unit` says which.
function Progress.slotVerdict(slot, plan, worn, weights, dpsPerPoint)
	local plannedItemId = plan ~= nil and plan.itemId or nil
	local equippedItemId = worn ~= nil and worn.itemId or nil
	if plannedItemId == nil and equippedItemId == nil then
		return nil
	end
	local scored = weights ~= nil and plannedItemId ~= nil and hasAnyStat(plan.stats)
	local state, delta
	if plannedItemId ~= nil and plannedItemId == equippedItemId then
		state = "match"
		delta = scored and 0 or nil
	elseif equippedItemId == nil then
		state = "empty"
		delta = scored and -Gear.score(plan.stats, weights) or nil
	elseif plannedItemId == nil then
		state = "noplan"
	else
		delta = scored and (Gear.score(worn.stats, weights) - Gear.score(plan.stats, weights)) or nil
		state = (delta ~= nil and delta > 0) and "better" or "behind"
	end
	return {
		slot = slot,
		plannedItemId = plannedItemId,
		equippedItemId = equippedItemId,
		delta = inUnit(delta, dpsPerPoint),
		state = state,
		unit = dpsPerPoint ~= nil and "dps" or "score",
	}
end

--- Specs that fight in melee, whose weapons lead the slot list. Every
--- warrior and rogue spec; for the rest, only the melee tree (Forever's
--- hunters are melee only in survival). A feral druid's slug is
--- "druid-feral..." in every spelling the data uses, so it is a prefix.
local MELEE_CLASSES = { warrior = true, rogue = true }
local MELEE_SPECS = {
	["paladin-retribution"] = true, ["shaman-enhancement"] = true, ["hunter-survival"] = true,
}
local WEAPON_SLOTS = { main_hand = true, off_hand = true, ranged = true }

local function isMelee(spec)
	if spec == nil then
		return false
	end
	local class = spec:match("^([^-]+)")
	return MELEE_CLASSES[class] == true or MELEE_SPECS[spec] == true
		or spec:find("^druid%-feral") ~= nil
end

--- The paper doll's slot order (Export.INVENTORY_SLOTS), with main hand,
--- off hand and ranged moved to the front for a melee spec.
local function slotOrder(spec)
	local order = {}
	if isMelee(spec) then
		for _, entry in ipairs(Export.INVENTORY_SLOTS) do
			if WEAPON_SLOTS[entry.slot] then
				order[#order + 1] = entry.slot
			end
		end
		for _, entry in ipairs(Export.INVENTORY_SLOTS) do
			if not WEAPON_SLOTS[entry.slot] then
				order[#order + 1] = entry.slot
			end
		end
		return order
	end
	for _, entry in ipairs(Export.INVENTORY_SLOTS) do
		order[#order + 1] = entry.slot
	end
	return order
end

--- The spec and its weights for a build, or nil for either.
local function weightsFor(data, build, ranks)
	if build == nil then
		return nil, nil
	end
	local spec = Gear.specOf(data, build.classSlug, ranks or {})
	local weights = spec ~= nil and data.weights ~= nil and data.weights[spec] or nil
	return spec, weights
end

--- A slotVerdict for every slot the plan or the paper doll names, in
--- paper-doll order (weapons first for a melee spec). With no build,
--- every worn item reads "noplan".
function Progress.slots(data, build, ranks, equipped, dpsPerPoint)
	equipped = equipped or {}
	local spec, weights = weightsFor(data, build, ranks)
	local planned = {}
	for _, entry in ipairs(build ~= nil and build.gear or {}) do
		planned[entry.slot] = entry
	end
	local rows = {}
	for _, slot in ipairs(slotOrder(spec)) do
		local verdict = Progress.slotVerdict(slot, planned[slot], equipped[slot], weights, dpsPerPoint)
		if verdict ~= nil then
			rows[#rows + 1] = verdict
		end
	end
	return rows
end

--- The slot furthest behind the plan: the lowest delta below zero among
--- `slots` (Progress.slots' rows), first in order on a tie. `item` is the
--- planned item to go and get. nil when nothing is behind.
function Progress.biggestGap(slots)
	local worst
	for _, row in ipairs(slots or {}) do
		if row.delta ~= nil and row.delta < 0 and (worst == nil or row.delta < worst.delta) then
			worst = row
		end
	end
	if worst == nil then
		return nil
	end
	return { slot = worst.slot, delta = worst.delta, item = worst.plannedItemId }
end

--- A bag item as { link, itemId, equipLocation, itemClass, itemSubClass,
--- stats }: a plain table is taken as it is, a link string is read from the
--- client (the class and subclass ids where it gives them, else its English
--- type names, which Rules.canEquip reads too).
local function describe(item)
	if type(item) == "table" then
		return item
	end
	local id, itemType, itemSubType, equipLocation, _, classId, subClassId = Compat.itemInfoInstant(item)
	return {
		link = item, itemId = tonumber(id), equipLocation = equipLocation,
		itemClass = classId or itemType, itemSubClass = subClassId or itemSubType,
		stats = Gear.statsOf(item),
	}
end

local TWO_HANDED = "INVTYPE_2HWEAPON"

--- The smallest gain, in its own unit (1 DPS or 1 score point), that counts as an upgrade.
Progress.MIN_GAIN = 1

--- Bag items that beat what is worn in a slot they fit, each once, in the
--- slot where it gains the most; sorted by delta, biggest first. A slot
--- with nothing worn counts what is worn as zero. An item whose link is
--- worn is never listed (seen in game: the worn helm offered as an
--- upgrade over itself), nor is an item scored against its own copy, nor
--- one Forever's rules say the character cannot wear in that slot at
--- `level` (Rules.canEquip, not "ok": plate before 40, a libram on a
--- warrior). `level` is the client's when nil. `unit` is "dps" when
--- `dpsPerPoint` is given, else "score".
---
--- The hands are one pair: with a two-hander worn (its `equipLoc`
--- INVTYPE_2HWEAPON), nothing goes in the off hand, and a two-hander from
--- the bags is scored against the main hand and the off hand together, the
--- net of the whole swap. A gain under MIN_GAIN in its unit is no upgrade.
function Progress.bagUpgrades(data, build, ranks, bagLinks, equipped, dpsPerPoint, level)
	local _, weights = weightsFor(data, build, ranks)
	if weights == nil then
		return {}
	end
	equipped = equipped or {}
	level = level or Compat.playerLevel()
	local classToken = build.classSlug ~= nil and build.classSlug:upper() or nil
	local function wornScore(worn)
		return worn ~= nil and Gear.score(worn.stats, weights) or 0
	end
	local twoHanderWorn = equipped.main_hand ~= nil and equipped.main_hand.equipLoc == TWO_HANDED
	local wornLinks = {}
	for _, worn in pairs(equipped) do
		if worn.link ~= nil then
			wornLinks[worn.link] = true
		end
	end
	local found = {}
	for _, raw in ipairs(bagLinks or {}) do
		local item = describe(raw)
		if item.link == nil or not wornLinks[item.link] then
			local score = Gear.score(item.stats, weights)
			local bestSlot, bestDelta
			for _, slot in ipairs(Gear.SLOTS_BY_EQUIP_LOCATION[item.equipLocation] or {}) do
				local worn = equipped[slot]
				local wearable = Rules.canEquip(classToken, level, item.itemClass, item.itemSubClass,
					item.equipLocation, slot) == "ok"
				local blocked = slot == "off_hand" and twoHanderWorn
				if wearable and not blocked and (worn == nil or worn.itemId == nil or worn.itemId ~= item.itemId) then
					local replaced = wornScore(worn)
					if item.equipLocation == TWO_HANDED and slot == "main_hand" then
						replaced = replaced + wornScore(equipped.off_hand)
					end
					local delta = score - replaced
					if bestDelta == nil or delta > bestDelta then
						bestSlot, bestDelta = slot, delta
					end
				end
			end
			local gain = bestDelta ~= nil and inUnit(bestDelta, dpsPerPoint) or nil
			if gain ~= nil and gain >= Progress.MIN_GAIN then
				found[#found + 1] = {
					link = item.link, itemId = item.itemId, slot = bestSlot,
					delta = gain,
					unit = dpsPerPoint ~= nil and "dps" or "score",
				}
			end
		end
	end
	table.sort(found, function(left, right)
		if left.delta ~= right.delta then
			return left.delta > right.delta
		end
		return left.slot < right.slot
	end)
	return found
end

--- Whether a pick's `source` ("The Deadmines: Gilnid", or just "The
--- Deadmines") drops in the instance named `instanceName`.
local function dropsIn(source, instanceName)
	if type(source) ~= "string" then
		return false
	end
	local prefix = instanceName .. ": "
	return source == instanceName or source:sub(1, #prefix) == prefix
end

--- The band's BiS picks (Gear.pick) that drop in the instance the player is
--- standing in, each { slot, itemId, source }, in Progress.slots' order
--- (weapons first for a melee spec). A pick worn in any slot is left out (a
--- ring or trinket has two), and an item two slots share is listed once. The
--- chip's "here for you" line and the panel's targets region read this.
--- Empty with no band, faction or instance, or nothing that drops here.
function Progress.picksHere(data, spec, band, faction, instanceName, equipped)
	local picks = {}
	if spec == nil or band == nil or faction == nil or type(instanceName) ~= "string" then
		return picks
	end
	local seen = {}
	for _, worn in pairs(equipped or {}) do
		if type(worn) == "table" and worn.itemId ~= nil then
			seen[worn.itemId] = true
		end
	end
	for _, slot in ipairs(slotOrder(spec)) do
		local pick = Gear.pick(spec, band, faction, slot, data)
		if pick ~= nil and dropsIn(pick.source, instanceName) and not seen[pick.item_id] then
			seen[pick.item_id] = true
			picks[#picks + 1] = { slot = slot, itemId = pick.item_id, source = pick.source }
		end
	end
	return picks
end

--- Talent points the character has to spend at `level`: Classic grants one
--- a level from 10, so it is the points earned less the points in `ranks`
--- (a respec changes the ranks, not the rule), and never below none. Pure:
--- the client's trait reader answered a positive count beside a talent
--- window that said "Unspent Talents 0" (seen in game, L3.1 fix round 1).
function Progress.unspentPoints(level, ranks)
	if type(level) ~= "number" or level < 10 then
		return 0
	end
	return math.max(0, (level - 9) - Progress.pointsSpent(ranks))
end

ns.Progress = Progress
return Progress
