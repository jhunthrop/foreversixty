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
--- take. `nextLevel` is the level the character's next point arrives at
--- (Progress.nextPointLevel) when `level` is given; without it (the legacy
--- callers) it is the plan's own "taken at" level, 9 + the point's order
--- index, which is wrong for a character who took cells out of the plan's
--- order (seen in game: "Next at 16" at level 21). `done` is true once a
--- build has no point left (seen in game: a "2 of 51" beside eleven spent
--- points, because the old count stopped at the first unmet point).
function Progress.build(build, ranks, level)
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
	local nextLevel
	if point ~= nil and level ~= nil then
		nextLevel = Progress.nextPointLevel(level, ranks)
	elseif point ~= nil then
		nextLevel = 9 + point.index
	end
	return {
		matched = matched,
		total = #build.order,
		offPlan = math.max(0, Progress.pointsSpent(ranks) - matched),
		point = point,
		nextLevel = nextLevel,
		done = point == nil,
	}
end

--- Add order index `index` to a cell's runs: it extends the last run when it
--- follows that run's last point, else it opens a new one.
local function addToRuns(runs, index)
	local last = runs[#runs]
	if last ~= nil and last[2] == index - 1 then
		last[2] = index
	else
		runs[#runs + 1] = { index, index }
	end
end

--- The plan's points at one cell as its consecutive runs, in order:
--- { { first, last }, ... } of 1-based order indexes (Task L3.2 fix, runs:
--- the hunter's Lethal Attacks, taken at points 8-11 and 36, is
--- { { 8, 11 }, { 36, 36 } }). Empty without a build or for a cell the plan
--- never touches.
function Progress.runsFor(build, tab, tier, column)
	local runs = {}
	if build == nil then
		return runs
	end
	for index, step in ipairs(build.order) do
		if step.tab == tab and step.tier == tier and step.column == column then
			addToRuns(runs, index)
		end
	end
	return runs
end

--- The run (Progress.runsFor's) the character's next rank at the cell falls
--- in, given `have` ranks taken there: the first run whose ranks are not all
--- taken. Nil once every rank in the runs is taken (or overspent), or with no
--- runs. A run is taken whole: 2 of the 4 ranks in { 8, 11 } still name it.
function Progress.nextRun(runs, have)
	local taken = tonumber(have) or 0
	local counted = 0
	for _, run in ipairs(runs or {}) do
		counted = counted + run[2] - run[1] + 1
		if taken < counted then
			return run
		end
	end
	return nil
end

--- Every cell the plan touches, once each, in the order the plan first
--- reaches it: { tab, tier, column, want, have, firstIndex, lastIndex,
--- runs, node, next }. `want` is the points the plan puts there, `have` the
--- character's rank there (0 when none), `firstIndex` and `lastIndex` the
--- 1-based order indexes of the cell's first and last points, `runs` those
--- points as consecutive runs (Progress.runsFor's shape; the badge names the
--- next one, Progress.nextRun), `node`
--- the trait node id `data` carries for it (nil for a data gap), and `next`
--- whether the plan's next point (Progress.nextPoint) is in it. The talent
--- marks draw from this list only. Empty without a build.
function Progress.cells(build, ranks, data)
	local cells = {}
	if build == nil then
		return cells
	end
	ranks = ranks or {}
	local point = Progress.nextPoint(build, ranks)
	local byKey = {}
	for index, step in ipairs(build.order) do
		local key = cellKey(step)
		local cell = byKey[key]
		if cell == nil then
			local talent = Talents.cellOf(data, build.classSlug, step.tab, step.tier, step.column)
			cell = {
				tab = step.tab, tier = step.tier, column = step.column, want = 0, firstIndex = index, runs = {},
				have = (ranks[step.tab] or {})[Talents.cellKey(step.tier, step.column)] or 0,
				node = talent ~= nil and talent.node or nil,
				next = point ~= nil and cellKey(point) == key,
			}
			byKey[key] = cell
			cells[#cells + 1] = cell
		end
		cell.want = cell.want + 1
		cell.lastIndex = index
		addToRuns(cell.runs, index)
	end
	return cells
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

--- What one point of a weapon's own DPS is worth in `slot` under `weapon`
--- (Gear.weaponTerm's { slot, weight }: the band's measured term, in
--- reference-stat points): its weight on the slot it was measured on, 0 on any
--- other (the off hand, a hunter's melee weapon) and with no term (feral
--- druids, casters, healers, tanks).
local function weaponWeight(slot, weapon)
	if type(weapon) ~= "table" or weapon.slot ~= slot then
		return 0
	end
	return tonumber(weapon.weight) or 0
end

--- An item's score in `slot` under `weights`: its stat line (Gear.score), plus
--- its own DPS (the `weapon_dps` stat, Rules.STAT_KEYS) at the band's measured
--- weapon term (`weapon`, Gear.currentWeights' own) on the slot it was
--- measured on alone, never through attack power. A shield or a held item has
--- no DPS line, so it adds nothing. The one score the chip and the item
--- tooltip compare items by.
function Progress.itemScore(stats, weights, slot, weapon)
	local total = Gear.score(stats, weights)
	local dps = type(stats) == "table" and stats.weapon_dps or nil
	if type(dps) == "number" and dps > 0 then
		total = total + dps * weaponWeight(slot, weapon)
	end
	return total
end

--- A gain rounded once to the precision it is shown at (one decimal for
--- "dps", whole points for "score"), the same both ways from zero. The chip's
--- bag upgrades and the item tooltip both take their floor (MIN_GAIN) and
--- their word from this value, so the two can never disagree: 0.96 DPS is
--- 1.0 (an upgrade in both), 0.94 is 0.9 (in neither).
function Progress.roundedGain(value, unit)
	if value == nil then
		return nil
	end
	local size = math.abs(value)
	local rounded = unit == "score" and math.floor(size + 0.5) or math.floor(size * 10 + 0.5) / 10
	return value < 0 and -rounded or rounded
end

--- A stat's weight in `slot`: its own weight, or for a weapon's DPS
--- (`weapon_dps`) the measured weapon term on its own slot (Progress.itemScore's
--- own rule).
local function weightIn(stat, slot, weights, weapon)
	if stat == "weapon_dps" then
		return weaponWeight(slot, weapon)
	end
	return (weights or {})[stat] or 0
end

--- What a gain is made of: per stat, `plan`'s amount minus every worn
--- side's (`plan` and each of `worn` are { stats, slot }; the worn stats
--- already carry their enchant, Gear.readEquipped), as { stat, delta,
--- weighted }, sorted by the weighted size, biggest first. A stat the weights
--- do not count is left out. Weighted in score units (reference-stat points).
function Progress.contributions(plan, worn, weights, weapon)
	local stats = {}
	for stat in pairs(plan.stats or {}) do
		stats[stat] = true
	end
	for _, side in ipairs(worn) do
		for stat in pairs(side.stats or {}) do
			stats[stat] = true
		end
	end
	local list = {}
	for stat in pairs(stats) do
		local amount = (plan.stats or {})[stat] or 0
		local delta = amount
		local weighted = amount * weightIn(stat, plan.slot, weights, weapon)
		for _, side in ipairs(worn) do
			local wornAmount = (side.stats or {})[stat] or 0
			delta = delta - wornAmount
			weighted = weighted - wornAmount * weightIn(stat, side.slot, weights, weapon)
		end
		if math.abs(weighted) > 1e-9 then
			list[#list + 1] = { stat = stat, delta = delta, weighted = weighted }
		end
	end
	table.sort(list, function(left, right)
		if math.abs(left.weighted) ~= math.abs(right.weighted) then
			return math.abs(left.weighted) > math.abs(right.weighted)
		end
		return left.stat < right.stat
	end)
	return list
end

--- One slot's verdict: the planned item against what is worn there.
--- `plan` is a build.gear entry ({ itemId, stats }), `worn` an equipped
--- entry ({ itemId, link, stats }); either may be nil, not both (nil then).
---
--- - the same item id on both sides is "match", whatever the two stat
---   lines say (seen in game: a helm reading "+10 over itself"), unless both
---   sides carry a link and the links differ (a second copy is scored);
--- - a planned item with nothing worn is "empty", short by the plan's
---   whole score, so an empty ring is the biggest gap;
--- - a worn item the plan never named is "noplan";
--- - otherwise "better" when worn minus planned is above zero, else
---   "behind" (also when it cannot be scored).
---
--- `delta` is worn minus planned, nil when the planned side carries no
--- stats (an FS1 code) or there are no weights; it is in DPS when
--- `dpsPerPoint` (DPS per 1.0 of the reference stat) is given, and
--- `unit` says which. `weapon` (Gear.currentWeights' measured weapon term) lets
--- a weapon's own DPS count (Progress.itemScore). `contributions` lists what the
--- gain is made of (Progress.contributions), the planned side minus the worn side.
function Progress.slotVerdict(slot, plan, worn, weights, dpsPerPoint, weapon)
	local plannedItemId = plan ~= nil and plan.itemId or nil
	local equippedItemId = worn ~= nil and worn.itemId or nil
	if plannedItemId == nil and equippedItemId == nil then
		return nil
	end
	-- Both sides carrying a link and the links differing (a random-suffix
	-- copy, a second copy) is two items, scored, never a match.
	local sameLink = plan == nil or worn == nil or plan.link == nil or worn.link == nil or plan.link == worn.link
	local scored = weights ~= nil and plannedItemId ~= nil and hasAnyStat(plan.stats)
	local function score(stats)
		return Progress.itemScore(stats, weights, slot, weapon)
	end
	local state, delta
	if plannedItemId ~= nil and plannedItemId == equippedItemId and sameLink then
		state = "match"
		delta = scored and 0 or nil
	elseif equippedItemId == nil then
		state = "empty"
		delta = scored and -score(plan.stats) or nil
	elseif plannedItemId == nil then
		state = "noplan"
	else
		delta = scored and (score(worn.stats) - score(plan.stats)) or nil
		state = (delta ~= nil and delta > 0) and "better" or "behind"
	end
	local contributions
	if scored and state ~= "match" then
		local sides = {}
		if worn ~= nil then
			sides[1] = { stats = worn.stats, slot = slot }
		end
		contributions = Progress.contributions({ stats = plan.stats, slot = slot }, sides, weights, weapon)
	end
	return {
		slot = slot,
		contributions = contributions,
		plannedItemId = plannedItemId,
		equippedItemId = equippedItemId,
		delta = inUnit(delta, dpsPerPoint),
		state = state,
		unit = dpsPerPoint ~= nil and "dps" or "score",
	}
end

local TWO_HANDED = "INVTYPE_2HWEAPON"

--- One hand slot's verdict, with the two hands as one pair: the rule the
--- chip's bag upgrades and the item tooltip share. `hovered` is the candidate
--- ({ itemId, link, equipLocation, stats }), `wornMain` and `wornOff` what the
--- hands hold ({ itemId, link, equipLoc, stats }).
---
--- - the off hand under a worn two-hander is "blocked" (no delta): nothing
---   goes there until the two-hander comes off;
--- - a two-hander for the main hand is scored against both hands together,
---   the net of the whole swap (delta is both worn minus the two-hander);
--- - everything else is Progress.slotVerdict against that hand.
---
--- Each item's own DPS counts in the hand it is scored in, on the hand the
--- weapon term was measured on alone (Progress.itemScore).
function Progress.weaponVerdict(slot, hovered, wornMain, wornOff, weights, dpsPerPoint, weapon)
	if slot == "off_hand" and wornMain ~= nil and wornMain.equipLoc == TWO_HANDED then
		return {
			slot = slot, state = "blocked",
			plannedItemId = hovered ~= nil and hovered.itemId or nil, equippedItemId = wornMain.itemId,
			unit = dpsPerPoint ~= nil and "dps" or "score",
		}
	end
	if slot == "main_hand" and hovered ~= nil and hovered.equipLocation == TWO_HANDED and wornOff ~= nil then
		local scored = weights ~= nil and hovered.itemId ~= nil and hasAnyStat(hovered.stats)
		local delta
		if scored then
			local worn = Progress.itemScore(wornOff.stats, weights, "off_hand", weapon)
			if wornMain ~= nil then
				worn = worn + Progress.itemScore(wornMain.stats, weights, "main_hand", weapon)
			end
			delta = worn - Progress.itemScore(hovered.stats, weights, "main_hand", weapon)
		end
		return {
			slot = slot,
			contributions = scored and Progress.contributions({ stats = hovered.stats, slot = "main_hand" }, {
				{ stats = wornOff.stats, slot = "off_hand" },
				wornMain ~= nil and { stats = wornMain.stats, slot = "main_hand" } or nil,
			}, weights, weapon) or nil,
			plannedItemId = hovered.itemId,
			equippedItemId = wornMain ~= nil and wornMain.itemId or wornOff.itemId,
			delta = inUnit(delta, dpsPerPoint),
			state = (delta ~= nil and delta > 0) and "better" or "behind",
			unit = dpsPerPoint ~= nil and "dps" or "score",
		}
	end
	return Progress.slotVerdict(slot, hovered, slot == "main_hand" and wornMain or wornOff, weights, dpsPerPoint,
		weapon)
end

--- Progress.weaponVerdict for the two hands, Progress.slotVerdict for every
--- other slot: the one call the chip and the item tooltip score a candidate
--- in a slot with. `equipped` is Gear.readEquipped()'s table.
function Progress.itemVerdict(slot, hovered, equipped, weights, dpsPerPoint, weapon)
	equipped = equipped or {}
	if slot == "main_hand" or slot == "off_hand" then
		return Progress.weaponVerdict(slot, hovered, equipped.main_hand, equipped.off_hand, weights, dpsPerPoint,
			weapon)
	end
	return Progress.slotVerdict(slot, hovered, equipped[slot], weights, dpsPerPoint, weapon)
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
--- `dpsPerPoint` is given, else "score". A weapon's own DPS counts at the band's
--- measured weapon term at `level` (Gear.weaponTerm), or `current`'s.
---
--- The hands are one pair: with a two-hander worn (its `equipLoc`
--- INVTYPE_2HWEAPON), nothing goes in the off hand, and a two-hander from
--- the bags is scored against the main hand and the off hand together, the
--- net of the whole swap. A gain under MIN_GAIN in its unit is no upgrade, except
--- for an item that fills an EMPTY slot: that is one at any gain that shows.
function Progress.bagUpgrades(data, build, ranks, bagLinks, equipped, dpsPerPoint, level, current)
	local spec, weights = weightsFor(data, build, ranks)
	level = level or Compat.playerLevel()
	local weapon = Gear.weaponTerm(Gear.dpsFigure(data, spec, level), dpsPerPoint)
	-- `current` (Gear.currentWeights, as the tooltip reads it): the sim's weights and weapon term
	-- when a weights message is in use (L3.4 fix 6 ruling: one source of weights).
	if type(current) == "table" and type(current.weights) == "table" then
		weights = current.weights
		weapon = current.weapon
	end
	local unit = dpsPerPoint ~= nil and "dps" or "score"
	if weights == nil then
		return {}
	end
	equipped = equipped or {}
	local classToken = build.classSlug ~= nil and build.classSlug:upper() or nil
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
			local candidate = { itemId = item.itemId, equipLocation = item.equipLocation, stats = item.stats }
			local bestSlot, bestDelta, bestState
			for _, slot in ipairs(Gear.SLOTS_BY_EQUIP_LOCATION[item.equipLocation] or {}) do
				local wearable = Rules.canEquip(classToken, level, item.itemClass, item.itemSubClass,
					item.equipLocation, slot) == "ok"
				-- Progress.itemVerdict: the two hands as one pair (nothing in the
				-- off hand under a two-hander; a two-hander net of both).
				local verdict = wearable and Progress.itemVerdict(slot, candidate, equipped, weights, nil, weapon) or nil
				-- A "match" is the item scored against its own copy; "blocked" a
				-- hand the worn two-hander holds. The delta is worn minus candidate.
				if verdict ~= nil and verdict.delta ~= nil and verdict.state ~= "match" and verdict.state ~= "blocked" then
					local delta = -verdict.delta
					if bestDelta == nil or delta > bestDelta then
						bestSlot, bestDelta, bestState = slot, delta, verdict.state
					end
				end
			end
			local gain = bestDelta ~= nil and inUnit(bestDelta, dpsPerPoint) or nil
			-- The floor reads the gain as shown (Progress.roundedGain), as the tooltip does.
			-- An item that fills an EMPTY slot is an upgrade at any gain that shows (L4.1
			-- fix round 3: "Fills an empty slot, +0.8 DPS" with no chip line and no arrow).
			local shown = Progress.roundedGain(gain, unit)
			local fills = bestState == "empty" and shown ~= nil and shown > 0
			if shown ~= nil and (shown >= Progress.MIN_GAIN or fills) then
				found[#found + 1] = {
					link = item.link, itemId = item.itemId, slot = bestSlot,
					delta = gain,
					unit = unit,
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

--- The band's BiS pick (Gear.pick's decoded table) still to get for each slot,
--- under the pair rule: a pick worn in any slot is left out (a ring or trinket
--- has two), and an item two slots share is offered once, on the first slot in
--- Progress.slots' order. `keep(pick)` (optional) narrows the picks counted (the
--- instance's, for picksHere). { [slot] = pick }, empty with no spec, band or faction.
--- The empty paper-doll slot's hover (ItemTooltip.emptySlotSection) and picksHere
--- read this one rule.
function Progress.slotPicks(data, spec, band, faction, equipped, keep)
	local picks = {}
	if spec == nil or band == nil or faction == nil then
		return picks, {}
	end
	local seen = {}
	for _, worn in pairs(equipped or {}) do
		if type(worn) == "table" and worn.itemId ~= nil then
			seen[worn.itemId] = true
		end
	end
	local order = {}
	for _, slot in ipairs(slotOrder(spec)) do
		local pick = Gear.pick(spec, band, faction, slot, data)
		if pick ~= nil and (keep == nil or keep(pick)) and not seen[pick.item_id] then
			seen[pick.item_id] = true
			picks[slot] = pick
			order[#order + 1] = slot
		end
	end
	return picks, order
end

--- The band's BiS picks (Gear.pick) that drop in the instance the player is
--- standing in, each { slot, itemId, source, chance }, in Progress.slots' order
--- (weapons first for a melee spec), under Progress.slotPicks' pair rule. The
--- chip's "here for you" line and the panel's targets region read this.
--- Empty with no band, faction or instance, or nothing that drops here.
function Progress.picksHere(data, spec, band, faction, instanceName, equipped)
	local here = {}
	if type(instanceName) ~= "string" then
		return here
	end
	local picks, order = Progress.slotPicks(data, spec, band, faction, equipped, function(pick)
		return dropsIn(pick.source, instanceName)
	end)
	for _, slot in ipairs(order) do
		local pick = picks[slot]
		here[#here + 1] = { slot = slot, itemId = pick.item_id, source = pick.source, chance = pick.chance }
	end
	return here
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

--- The plan's points the character can reach now: those spent plus those
--- waiting to be spent at `level` (Progress.unspentPoints). The talent marks
--- dim a badge whose order index lies beyond it.
function Progress.reachablePoints(level, ranks)
	return Progress.pointsSpent(ranks) + Progress.unspentPoints(level, ranks)
end

--- The nodes holding a point the plan never put there, in `data`'s tree
--- order: { tab, tier, column, have, node }. A cell the plan touches is
--- never listed, however far over the plan it is (Progress.build counts
--- that in offPlan). The talent marks warn on these nodes' own tooltips.
function Progress.offPlanCells(build, ranks, data)
	local cells = {}
	local class = build ~= nil and data ~= nil and data.classes ~= nil and data.classes[build.classSlug] or nil
	if class == nil then
		return cells
	end
	ranks = ranks or {}
	local planned = {}
	for _, point in ipairs(build.order) do
		planned[cellKey(point)] = true
	end
	for tab, tree in ipairs(class.tabs) do
		for _, talent in ipairs(tree.talents or {}) do
			local point = { tab = tab, tier = talent.tier, column = talent.column }
			local have = (ranks[tab] or {})[Talents.cellKey(talent.tier, talent.column)] or 0
			if have > 0 and not planned[cellKey(point)] then
				cells[#cells + 1] = { tab = tab, tier = talent.tier, column = talent.column, have = have, node = talent.node }
			end
		end
	end
	return cells
end

--- The plan's points per tree: { [tab] = count }, 0 for a tree it skips.
--- Follow.buildName names an unnamed build after it.
function Progress.planSplit(build)
	local split = {}
	for _, point in ipairs(build ~= nil and build.order or {}) do
		split[point.tab] = (split[point.tab] or 0) + 1
	end
	return split
end

--- The level cap: no talent point arrives past it.
Progress.MAX_LEVEL = 60

--- The level the character's next talent point arrives at: `level` itself
--- while a point is unspent, else the later of the next level and the level
--- that grants point number pointsSpent + 1 (10 + pointsSpent). nil at the
--- level cap, where no level follows (no "Next at 61"). The "next at"
--- number; Progress.levelForPoint stays the plan's own "taken at" label.
function Progress.nextPointLevel(level, ranks)
	if type(level) == "number" and level >= Progress.MAX_LEVEL then
		return nil
	end
	if Progress.unspentPoints(level, ranks) > 0 then
		return level
	end
	return math.max((level or 0) + 1, 10 + Progress.pointsSpent(ranks))
end

ns.Progress = Progress
return Progress
