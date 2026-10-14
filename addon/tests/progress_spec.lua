-- addon/tests/progress_spec.lua
-- Progress: every count, delta, verdict and gap the addon shows, computed
-- once (spec section 2.1). Pure: builds, ranks, equipped items and bag
-- items are plain tables here, and this spec never touches the client mock.
local Progress = require("Progress")
local L = require("Locale")

local function cell(tab, tier, column)
	return { tab = tab, tier = tier, column = column }
end

-- Forever's warrior cells (Data.lua): Arms tab 1, Fury tab 2.
local IHS, DEFLECTION, REND, OVERPOWER = cell(1, 1, 1), cell(1, 1, 2), cell(1, 1, 3), cell(1, 2, 4)
local CRUELTY = cell(2, 1, 3)

--- A 51-point warrior order: Improved Heroic Strike twice, then Deflection,
--- Rend, Overpower and Cruelty, then the rest of the build.
local function warriorBuild()
	local order = {}
	local function take(point, times)
		for _ = 1, times do
			order[#order + 1] = point
		end
	end
	take(IHS, 2)
	take(DEFLECTION, 5)
	take(REND, 3)
	take(OVERPOWER, 2)
	take(CRUELTY, 5)
	take(cell(1, 3, 2), 1)
	take(cell(1, 3, 3), 3)
	take(cell(1, 4, 2), 2)
	take(cell(1, 4, 3), 5)
	take(cell(2, 2, 2), 5)
	take(cell(2, 2, 3), 5)
	take(cell(2, 3, 1), 5)
	take(cell(2, 3, 2), 5)
	take(cell(2, 4, 1), 3)
	return { classSlug = "warrior", order = order }
end

-- IHS 2, Rend 3, Overpower 2, Cruelty 4.
local RANKS = {
	[1] = { ["1:1"] = 2, ["1:3"] = 3, ["2:4"] = 2 },
	[2] = { ["1:3"] = 4 },
}

describe("Progress.build", function()
	it("counts the 51-point warrior fixture: 11 matched, none off plan, Deflection next at 12", function()
		local build = warriorBuild()
		assert.are.equal(51, #build.order)
		local progress = Progress.build(build, RANKS)
		assert.are.equal(11, progress.matched)
		assert.are.equal(51, progress.total)
		assert.are.equal(0, progress.offPlan)
		assert.are.equal(3, progress.point.index)
		assert.are.same({ tab = 1, tier = 1, column = 2, index = 3 }, progress.point)
		assert.are.equal(12, progress.nextLevel)
		assert.is_false(progress.done)
	end)

	it("counts a rank over the plan as off plan, not as matched", function()
		local ranks = { [1] = { ["1:1"] = 3, ["1:3"] = 3, ["2:4"] = 2 }, [2] = { ["1:3"] = 4 } }
		local progress = Progress.build(warriorBuild(), ranks)
		assert.are.equal(11, progress.matched)
		assert.are.equal(1, progress.offPlan)
	end)

	it("is done, with no next point or level, once every point is taken", function()
		local progress = Progress.build({ order = { IHS } }, { [1] = { ["1:1"] = 1 } })
		assert.is_true(progress.done)
		assert.is_nil(progress.point)
		assert.is_nil(progress.nextLevel)
	end)

	it("says nothing is matched or done without a build", function()
		local progress = Progress.build(nil, RANKS)
		assert.are.equal(0, progress.matched)
		assert.are.equal(0, progress.total)
		assert.is_nil(progress.point)
		assert.is_false(progress.done)
	end)
end)

-- Fix round 2 (seen in game): a level-21 warrior with 12 spent read "Next at 16",
-- the plan's order index of the first unmet cell (#7) and not the level the
-- character's next point arrives at.
describe("Progress.nextPointLevel", function()
	-- IHS 2, Deflection 4 of 5, Rend 3, Overpower 2, Cruelty 1: twelve spent,
	-- and the first unmet cell in the plan's order is #7 (the fifth Deflection).
	local OUT_OF_ORDER = { [1] = { ["1:1"] = 2, ["1:2"] = 4, ["1:3"] = 3, ["2:4"] = 2 }, [2] = { ["1:3"] = 1 } }

	it("is the next level when no point is unspent, counted from the points spent", function()
		assert.are.equal(22, Progress.nextPointLevel(21, OUT_OF_ORDER))
		assert.are.equal(10, Progress.nextPointLevel(9, {}))
		assert.are.equal(10, Progress.nextPointLevel(4, {}))
	end)

	-- Fix round 3 (re-review): no "Next at 61".
	it("has no next level at the level cap", function()
		assert.is_nil(Progress.nextPointLevel(60, OUT_OF_ORDER))
		assert.is_nil(Progress.build(warriorBuild(), OUT_OF_ORDER, 60).nextLevel)
	end)

	it("is the level now when a point is unspent", function()
		assert.are.equal(22, Progress.nextPointLevel(22, OUT_OF_ORDER))
	end)

	it("sets Progress.build's nextLevel when the level is given, not the plan's order index", function()
		local progress = Progress.build(warriorBuild(), OUT_OF_ORDER, 21)
		assert.are.equal(7, progress.point.index)
		assert.are.equal(22, progress.nextLevel)
		assert.are.equal(16, Progress.levelForPoint(progress.point.index).level)
	end)
end)

describe("Progress.levelForPoint", function()
	it("takes point N at level 9 + N", function()
		assert.are.same({ level = 12, text = "Take at level 12" }, Progress.levelForPoint(3))
		assert.are.equal("Take at level %d", L.progressTakeAt)
	end)
end)

-- strength 2, attack power 1: a plan of 10 strength scores 20.
local WEIGHTS = { strength = 2, attack_power = 1 }
local PLAN = { itemId = 100, stats = { strength = 10 } }

describe("Progress.slotVerdict", function()
	it("calls the same item on both sides a match, even when the plan's stats read better", function()
		local verdict = Progress.slotVerdict("head", PLAN, { itemId = 100, stats = { strength = 5 } }, WEIGHTS)
		assert.are.equal("match", verdict.state)
		assert.are.equal(0, verdict.delta)
		assert.are.equal("head", verdict.slot)
		assert.are.equal(100, verdict.plannedItemId)
		assert.are.equal(100, verdict.equippedItemId)
	end)

	it("calls a planned slot with nothing worn empty, short by the plan's whole score", function()
		local verdict = Progress.slotVerdict("finger1", PLAN, nil, WEIGHTS)
		assert.are.equal("empty", verdict.state)
		assert.are.equal(-20, verdict.delta)
		assert.are.equal("score", verdict.unit)
		assert.is_nil(verdict.equippedItemId)
	end)

	it("reads better or behind from the sign of worn minus planned", function()
		local better = Progress.slotVerdict("chest", PLAN, { itemId = 200, stats = { strength = 15 } }, WEIGHTS)
		assert.are.equal("better", better.state)
		assert.are.equal(10, better.delta)
		local behind = Progress.slotVerdict("chest", PLAN, { itemId = 201, stats = { strength = 5 } }, WEIGHTS)
		assert.are.equal("behind", behind.state)
		assert.are.equal(-10, behind.delta)
	end)

	it("says noplan for a worn item the plan never named, with no delta", function()
		local verdict = Progress.slotVerdict("back", nil, { itemId = 300, stats = { strength = 3 } }, WEIGHTS)
		assert.are.equal("noplan", verdict.state)
		assert.is_nil(verdict.delta)
	end)

	it("gives no delta when the planned side carries no stats", function()
		local worn = { itemId = 200, stats = { strength = 15 } }
		local verdict = Progress.slotVerdict("chest", { itemId = 100 }, worn, WEIGHTS)
		assert.is_nil(verdict.delta)
		assert.are.equal("behind", verdict.state)
	end)

	it("converts the delta to DPS when a DPS per point is given", function()
		local verdict = Progress.slotVerdict("finger1", PLAN, nil, WEIGHTS, 0.5)
		assert.are.equal(-10, verdict.delta)
		assert.are.equal("dps", verdict.unit)
	end)
end)

local DATA = {
	classes = {
		warrior = { tabs = { { name = "Arms" }, { name = "Fury" }, { name = "Protection" } } },
		mage = { tabs = { { name = "Arcane" }, { name = "Fire" }, { name = "Frost" } } },
	},
	weights = {
		["warrior-arms"] = WEIGHTS,
		["mage-arcane"] = { intellect = 1 },
	},
}

local ARMS_RANKS = { [1] = { ["1:1"] = 3 } }

local function slotsOf(rows)
	local names = {}
	for index, row in ipairs(rows) do
		names[index] = row.slot
	end
	return names
end

-- L3.7 item 43: Progress.slots and Progress.biggestGap were read by nothing but this spec; the
-- reference versions live here now, over Progress.slotOrder and Progress.slotVerdict.
local function slotRows(data, build, ranks, equipped, dpsPerPoint)
	local Gear = require("Gear")
	local spec = build ~= nil and Gear.specOf(data, build.classSlug, ranks) or nil
	local weights = spec ~= nil and data.weights ~= nil and data.weights[spec] or nil
	local planned = {}
	for _, entry in ipairs(build ~= nil and build.gear or {}) do
		planned[entry.slot] = entry
	end
	local rows = {}
	for _, slot in ipairs(Progress.slotOrder(spec)) do
		local verdict = Progress.slotVerdict(slot, planned[slot], equipped[slot], weights, dpsPerPoint)
		if verdict ~= nil then
			rows[#rows + 1] = verdict
		end
	end
	return rows
end

local function biggestGapOf(rows)
	local worst
	for _, row in ipairs(rows or {}) do
		if row.delta ~= nil and row.delta < 0 and (worst == nil or row.delta < worst.delta) then
			worst = row
		end
	end
	if worst == nil then
		return nil
	end
	return { slot = worst.slot, delta = worst.delta, item = worst.plannedItemId }
end

describe("Progress.slotVerdict over the slot order (the reference rows)", function()
	local gear = {
		{ slot = "head", itemId = 1, stats = { strength = 4 } },
		{ slot = "main_hand", itemId = 2, stats = { strength = 8 } },
		{ slot = "finger1", itemId = 3, stats = { strength = 2 } },
	}
	local equipped = {
		head = { itemId = 1, stats = { strength = 4 } },
		chest = { itemId = 9, stats = { strength = 6 } },
		main_hand = { itemId = 5, stats = { strength = 3 } },
	}

	it("puts the weapons first for a melee spec, then the paper doll", function()
		local rows = slotRows(DATA, { classSlug = "warrior", gear = gear }, ARMS_RANKS, equipped)
		assert.are.same({ "main_hand", "head", "chest", "finger1" }, slotsOf(rows))
		assert.are.equal("behind", rows[1].state)
		assert.are.equal(-10, rows[1].delta)
		assert.are.equal("match", rows[2].state)
		assert.are.equal("noplan", rows[3].state)
		assert.are.equal("empty", rows[4].state)
	end)

	it("keeps the paper-doll order for a caster", function()
		local rows = slotRows(DATA, { classSlug = "mage", gear = gear }, ARMS_RANKS, equipped)
		assert.are.same({ "head", "chest", "finger1", "main_hand" }, slotsOf(rows))
	end)

	it("reads every delta in DPS when a DPS per point is given", function()
		local rows = slotRows(DATA, { classSlug = "warrior", gear = gear }, ARMS_RANKS, equipped, 2)
		for _, row in ipairs(rows) do
			assert.are.equal("dps", row.unit)
		end
		assert.are.equal(-20, rows[1].delta)
	end)

	it("lists what is worn as noplan without a build", function()
		local rows = slotRows(DATA, nil, ARMS_RANKS, equipped)
		assert.are.same({ "head", "chest", "main_hand" }, slotsOf(rows))
		assert.are.equal("noplan", rows[1].state)
	end)
end)

describe("the biggest gap over the slot rows", function()
	it("names the empty ring over a behind chest when the ring is further behind", function()
		local rows = {
			Progress.slotVerdict("chest", PLAN, { itemId = 201, stats = { strength = 5 } }, WEIGHTS),
			Progress.slotVerdict("finger1", { itemId = 400, stats = { strength = 10 } }, nil, WEIGHTS),
		}
		assert.are.same({ slot = "finger1", delta = -20, item = 400 }, biggestGapOf(rows))
	end)

	it("names the chest when the chest is further behind", function()
		local rows = {
			Progress.slotVerdict("chest", PLAN, { itemId = 201, stats = {} }, WEIGHTS),
			Progress.slotVerdict("finger1", { itemId = 400, stats = { strength = 5 } }, nil, WEIGHTS),
		}
		assert.are.equal("chest", biggestGapOf(rows).slot)
	end)

	it("is nil when nothing is behind the plan", function()
		local rows = {
			Progress.slotVerdict("head", PLAN, { itemId = 100, stats = {} }, WEIGHTS),
			Progress.slotVerdict("chest", PLAN, { itemId = 200, stats = { strength = 15 } }, WEIGHTS),
		}
		assert.is_nil(biggestGapOf(rows))
	end)
end)

describe("Progress.bagUpgrades", function()
	local build = { classSlug = "warrior", gear = {} }
	local equipped = {
		head = { itemId = 10, link = "|Hitem:10|h[Worn Helm]|h", stats = { strength = 5 } },
		finger1 = { itemId = 20, link = "|Hitem:20|h[Worn Ring]|h", stats = { strength = 1 } },
		finger2 = { itemId = 21, link = "|Hitem:21|h[Other Ring]|h", stats = { strength = 2 } },
	}
	local function bagItem(id, name, equipLocation, strength)
		return {
			link = "|Hitem:" .. id .. "|h[" .. name .. "]|h", itemId = id,
			equipLocation = equipLocation, stats = { strength = strength },
		}
	end

	it("lists bag items that beat the worn item, biggest first", function()
		local bag = {
			bagItem(11, "Better Helm", "INVTYPE_HEAD", 6),
			bagItem(12, "Worse Helm", "INVTYPE_HEAD", 2),
			bagItem(22, "Big Ring", "INVTYPE_FINGER", 4),
		}
		assert.are.same({
			{ link = "|Hitem:22|h[Big Ring]|h", itemId = 22, slot = "finger1", delta = 6, unit = "score" },
			{ link = "|Hitem:11|h[Better Helm]|h", itemId = 11, slot = "head", delta = 2, unit = "score" },
		}, Progress.bagUpgrades(DATA, build, ARMS_RANKS, bag, equipped))
	end)

	-- Contract 10's rules engine: a warrior wears plate from 40 and never a libram.
	it("never offers an item the character cannot wear yet, or ever", function()
		local plate = bagItem(30, "Plate Chest", "INVTYPE_CHEST", 40)
		plate.itemClass, plate.itemSubClass = 4, 4
		local libram = bagItem(31, "Libram", "INVTYPE_RELIC", 40)
		libram.itemClass, libram.itemSubClass = 4, 7
		local mail = bagItem(32, "Mail Chest", "INVTYPE_CHEST", 3)
		mail.itemClass, mail.itemSubClass = 4, 3
		local found = Progress.bagUpgrades(DATA, build, ARMS_RANKS, { plate, libram, mail }, equipped, nil, 20)
		assert.are.equal(1, #found)
		assert.are.equal(32, found[1].itemId)
		local at40 = Progress.bagUpgrades(DATA, build, ARMS_RANKS, { plate }, equipped, nil, 40)
		assert.are.equal(30, at40[1].itemId)
	end)

	it("scores by the weights it is handed, the sim's when a weights message is in use (L3.4 fix 6, 11)", function()
		local bag = { bagItem(11, "Better Helm", "INVTYPE_HEAD", 6) }
		assert.are.equal(1, #Progress.bagUpgrades(DATA, build, ARMS_RANKS, bag, equipped))
		-- A sim that weighs only stamina: strength gains nothing, so nothing upgrades.
		local current = { weights = { stamina = 1 }, tag = "sim" }
		assert.are.same({}, Progress.bagUpgrades(DATA, build, ARMS_RANKS, bag, equipped, nil, nil, current))
	end)

	-- L3.4 fix 7, 2: the chip's own wiring (Chip.scoredUpgrades): Gear.currentWeights, the sim's
	-- DPS per point while a weights message is in use, else the shipped one.
	describe("the chip's bag upgrades", function()
		local Chip = require("Chip")
		local KEY = "US/Ashbringer/Bob"
		local helm = function()
			return { bagItem(11, "Strong Helm", "INVTYPE_HEAD", 15) }
		end

		it("scores by the sim's weights while a weights message is in use", function()
			local shipped = Chip.scoredUpgrades(DATA, "warrior-arms", nil, nil, KEY, build, ARMS_RANKS, helm(), equipped, 20)
			assert.are.equal(1, #shipped)
			local inbox = { messages = { { type = "weights", spec = "warrior-arms",
				weights = { { stat = "stamina", weight = 1 } } } } }
			assert.are.same({}, Chip.scoredUpgrades(DATA, "warrior-arms", nil, inbox, KEY, build, ARMS_RANKS, helm(),
				equipped, 20))
		end)

		it("counts in the sim's DPS per point, not the shipped one", function()
			local inbox = { messages = { { type = "weights", spec = "warrior-arms", dps_per_point = 0.5,
				weights = { { stat = "strength", weight = 1 } } } } }
			local found = Chip.scoredUpgrades(DATA, "warrior-arms", { dps = 3 }, inbox, KEY, build, ARMS_RANKS, helm(),
				equipped, 20)
			assert.are.equal(1, #found)
			assert.are.equal("dps", found[1].unit)
			assert.are.equal(5, found[1].delta)
			local plain = Chip.scoredUpgrades(DATA, "warrior-arms", { dps = 3 }, nil, KEY, build, ARMS_RANKS, helm(),
				equipped, 20)
			assert.are.equal(60, plain[1].delta)
		end)

		-- L3.6 ruling 1: the band's DPS per point and its measured weapon term, on
		-- the main hand only; the chip and the tooltip read the same figures.
		it("prices a bag weapon at the band's figures, the weapon term on the main hand", function()
			local Gear = require("Gear")
			local banded = { classes = DATA.classes, weights = DATA.weights, dpsPerPoint = { ["warrior-arms"] = {
				reference = "attack_power", dps = 0.2, healing = false, weaponDps = 3,
				[20] = { reference = "attack_power", dps = 0.05, weaponDps = 0.5 },
				[30] = { reference = "attack_power", dps = 0.1, weaponDps = 1 } } } }
			local worn = { main_hand = { itemId = 60, link = "item:60", equipLoc = "INVTYPE_WEAPON",
				stats = { weapon_dps = 10 } } }
			local sword = { link = "|Hitem:63|h[Sword]|h", itemId = 63, equipLocation = "INVTYPE_WEAPON", itemClass = 2,
				itemSubClass = 7, stats = { strength = 5, weapon_dps = 20 } }
			local function at(level)
				return Chip.scoredUpgrades(banded, "warrior-arms", Gear.dpsFigure(banded, "warrior-arms", level), nil, KEY,
					build, ARMS_RANKS, { sword }, worn, level)[1]
			end
			-- Band 20: 5 Str x 2 = 10 AP x 0.05 = 0.5 DPS, + 10 weapon DPS x 0.5 = 5.0: 5.5 DPS.
			assert.are.equal("main_hand", at(25).slot)
			assert.is_true(math.abs(at(25).delta - 5.5) < 1e-9)
			-- Band 30: 10 x 0.1 + 10 x 1 = 11 DPS.
			assert.is_true(math.abs(at(35).delta - 11) < 1e-9)
			-- Below level 20, band 20; never the spec-level 0.2 / 3.
			assert.is_true(math.abs(at(15).delta - 5.5) < 1e-9)
		end)
	end)

	it("never lists an item that is already worn", function()
		local bag = { bagItem(10, "Worn Helm", "INVTYPE_HEAD", 50) }
		local found = Progress.bagUpgrades(DATA, build, ARMS_RANKS, bag, equipped)
		for _, upgrade in ipairs(found) do
			for _, worn in pairs(equipped) do
				assert.are_not.equal(worn.link, upgrade.link)
			end
		end
		assert.are.equal(0, #found)
	end)

	-- Seen in game (fix round 1): a shield offered to a two-handed Arms warrior.
	it("never offers an off hand or a shield while a two-hander is worn", function()
		local twoHanded = { main_hand = { itemId = 50, equipLoc = "INVTYPE_2HWEAPON", stats = { strength = 10 } } }
		local shield = bagItem(51, "Burnished Shield", "INVTYPE_SHIELD", 5)
		shield.itemClass, shield.itemSubClass = 4, 6
		local dagger = bagItem(52, "Dagger", "INVTYPE_WEAPON", 30)
		dagger.itemClass, dagger.itemSubClass = 2, 15
		local found = Progress.bagUpgrades(DATA, build, ARMS_RANKS, { shield, dagger }, twoHanded, nil, 20)
		assert.are.equal(1, #found)
		assert.are.equal("main_hand", found[1].slot)
		assert.are.equal(40, found[1].delta)
	end)

	it("scores a two-hander against the main hand and the off hand together", function()
		local pair = {
			main_hand = { itemId = 60, equipLoc = "INVTYPE_WEAPON", stats = { strength = 5 } },
			off_hand = { itemId = 61, equipLoc = "INVTYPE_SHIELD", stats = { strength = 3 } },
		}
		local axe = bagItem(62, "Great Axe", "INVTYPE_2HWEAPON", 12)
		axe.itemClass, axe.itemSubClass = 2, 1
		local found = Progress.bagUpgrades(DATA, build, ARMS_RANKS, { axe }, pair, nil, 20)
		assert.are.equal(1, #found)
		assert.are.equal(24 - 10 - 6, found[1].delta)
	end)

	it("does not call a gain under 1 in its unit an upgrade", function()
		local bag = { bagItem(11, "Better Helm", "INVTYPE_HEAD", 6) }
		assert.are.same({}, Progress.bagUpgrades(DATA, build, ARMS_RANKS, bag, equipped, 0.2))
	end)

	-- L4.1 fix round 3, 2 (hunter-10): an item that fills an EMPTY slot is always an
	-- upgrade, whatever its gain; a sidegrade over a WORN item stays under the floor.
	it("calls an item that fills an empty slot an upgrade below the floor", function()
		-- 2 Str x 2 = 4 points x 0.2 = 0.8 DPS into the empty legs slot.
		local legs = bagItem(13, "Thin Leggings", "INVTYPE_LEGS", 2)
		local found = Progress.bagUpgrades(DATA, build, ARMS_RANKS, { legs }, equipped, 0.2)
		assert.are.equal(1, #found)
		assert.are.equal("legs", found[1].slot)
		assert.is_true(math.abs(found[1].delta - 0.8) < 1e-9)
		-- The same gain over a worn helm is still no upgrade.
		local helm = bagItem(11, "Slightly Better Helm", "INVTYPE_HEAD", 7)
		assert.are.same({}, Progress.bagUpgrades(DATA, build, ARMS_RANKS, { helm }, equipped, 0.2))
	end)

	it("never calls an item that adds nothing to an empty slot an upgrade", function()
		local legs = bagItem(14, "Plain Leggings", "INVTYPE_LEGS", 0)
		assert.are.same({}, Progress.bagUpgrades(DATA, build, ARMS_RANKS, { legs }, equipped, 0.2))
	end)

	it("converts each delta to DPS when a DPS per point is given", function()
		local bag = { bagItem(11, "Better Helm", "INVTYPE_HEAD", 6) }
		local found = Progress.bagUpgrades(DATA, build, ARMS_RANKS, bag, equipped, 0.5)
		assert.are.equal(1, found[1].delta)
		assert.are.equal("dps", found[1].unit)
	end)
end)

describe("Progress.picksHere", function()
	-- Packed picks as Data.lua stores them (contract 1), under the real data's
	-- numeric band key.
	local HERE_DATA = {
		bis = {
			["warrior-arms"] = {
				[20] = {
					alliance = {
						head = "10399|D|The Deadmines: Edwin VanCleef",
						chest = "2317|C|Leatherworking",
						finger1 = "1156|D|The Deadmines: Gilnid",
						finger2 = "1156|D|The Deadmines: Gilnid",
						main_hand = "5187|D|The Deadmines: Rhahk'Zor",
						legs = "6087|D|The Deadmines",
						feet = "6089|D|Wailing Caverns: Mutanus the Devourer",
					},
				},
			},
		},
	}

	it("lists the band's picks that drop here, weapons first for a melee spec", function()
		local picks = Progress.picksHere(HERE_DATA, "warrior-arms", 20, "alliance", "The Deadmines", {})
		assert.are.same({
			{ slot = "main_hand", itemId = 5187, source = "The Deadmines: Rhahk'Zor" },
			{ slot = "head", itemId = 10399, source = "The Deadmines: Edwin VanCleef" },
			{ slot = "legs", itemId = 6087, source = "The Deadmines" },
			{ slot = "finger1", itemId = 1156, source = "The Deadmines: Gilnid" },
		}, picks)
	end)

	-- L4.2 ruling 4: the chip's click prints each pick with its boss and its chance.
	it("carries each pick's drop chance", function()
		local data = { bis = { ["warrior-arms"] = { [20] = { alliance = {
			main_hand = "7230|D|The Deadmines: Mr. Smite|0.2",
			legs = "6087|D|The Deadmines",
		} } } } }
		local picks = Progress.picksHere(data, "warrior-arms", 20, "alliance", "The Deadmines", {})
		assert.are.same({
			{ slot = "main_hand", itemId = 7230, source = "The Deadmines: Mr. Smite", chance = 0.2 },
			{ slot = "legs", itemId = 6087, source = "The Deadmines" },
		}, picks)
	end)

	it("leaves out a pick worn in the other slot of a pair", function()
		local picks = Progress.picksHere(HERE_DATA, "warrior-arms", 20, "alliance", "The Deadmines",
			{ finger2 = { itemId = 1156 } })
		for _, pick in ipairs(picks) do
			assert.are_not.equal(1156, pick.itemId)
		end
		assert.are.equal(3, #picks)
	end)

	it("leaves out a slot where the pick is already worn", function()
		local equipped = { main_hand = { itemId = 5187 }, head = { itemId = 1 } }
		local picks = Progress.picksHere(HERE_DATA, "warrior-arms", 20, "alliance", "The Deadmines", equipped)
		assert.are.equal("head", picks[1].slot)
		assert.are.equal(3, #picks)
	end)

	it("lists nothing for another instance's name, a missing band, or no instance", function()
		assert.are.same({}, Progress.picksHere(HERE_DATA, "warrior-arms", 20, "alliance", "The Dead", {}))
		assert.are.same({}, Progress.picksHere(HERE_DATA, "warrior-arms", 30, "alliance", "The Deadmines", {}))
		assert.are.same({}, Progress.picksHere(HERE_DATA, "warrior-arms", 20, "horde", "The Deadmines", {}))
		assert.are.same({}, Progress.picksHere(HERE_DATA, "warrior-arms", 20, "alliance", nil, {}))
		assert.are.same({}, Progress.picksHere(HERE_DATA, "warrior-arms", nil, "alliance", "The Deadmines", {}))
	end)

	-- L4.1 ruling 5: an empty paper-doll slot's pick follows picksHere's pair rule,
	-- read from the one function both call.
	describe("Progress.slotPicks", function()
		it("gives each slot the band's pick, wherever it drops", function()
			local picks = Progress.slotPicks(HERE_DATA, "warrior-arms", 20, "alliance", {})
			assert.are.equal(2317, picks.chest.item_id)
			assert.are.equal(6089, picks.feet.item_id)
			assert.are.equal(5187, picks.main_hand.item_id)
		end)

		it("offers a pick two slots share once, on the first in order", function()
			local picks = Progress.slotPicks(HERE_DATA, "warrior-arms", 20, "alliance", {})
			assert.are.equal(1156, picks.finger1.item_id)
			assert.is_nil(picks.finger2)
		end)

		it("leaves out a pick worn in any slot, the other of a pair included", function()
			local picks = Progress.slotPicks(HERE_DATA, "warrior-arms", 20, "alliance",
				{ finger2 = { itemId = 1156 }, head = { itemId = 10399 } })
			assert.is_nil(picks.finger1)
			assert.is_nil(picks.finger2)
			assert.is_nil(picks.head)
			assert.are.equal(2317, picks.chest.item_id)
		end)

		it("gives nothing with no band, faction or spec", function()
			assert.are.same({}, (Progress.slotPicks(HERE_DATA, "warrior-arms", nil, "alliance", {})))
			assert.are.same({}, (Progress.slotPicks(HERE_DATA, "warrior-arms", 20, nil, {})))
			assert.are.same({}, (Progress.slotPicks(HERE_DATA, nil, 20, "alliance", {})))
		end)
	end)
end)

describe("Progress.unspentPoints", function()
	-- Classic grants one point a level from 10 (fix round 1: the client's trait
	-- reader answered a positive count with "Unspent Talents 0" on screen).
	local TWELVE_SPENT = { [1] = { ["1:1"] = 3, ["1:3"] = 3, ["2:4"] = 2 }, [2] = { ["1:3"] = 4 } }

	it("has none to spend at 21 with twelve spent, and one at 22", function()
		assert.are.equal(0, Progress.unspentPoints(21, TWELVE_SPENT))
		assert.are.equal(1, Progress.unspentPoints(22, TWELVE_SPENT))
	end)

	it("has none below 10, one at 10, and never fewer than none", function()
		assert.are.equal(0, Progress.unspentPoints(9, {}))
		assert.are.equal(1, Progress.unspentPoints(10, {}))
		assert.are.equal(0, Progress.unspentPoints(10, TWELVE_SPENT))
		assert.are.equal(0, Progress.unspentPoints(nil, {}))
	end)
end)

-- Task L3.2 ruling 1: the talent marks draw from this list only.
describe("Progress.cells", function()
	-- Arms 1:1 carries a trait node; Fury 1:3 is a data gap (no talent there).
	local TALENTS = { classes = { warrior = { tabs = {
		{ name = "Arms", talents = { { name = "Improved Heroic Strike", tier = 1, column = 1, maxRank = 3, node = 101 } } },
		{ name = "Fury", talents = {} },
	} } } }
	-- Improved Heroic Strike, Cruelty, Improved Heroic Strike again.
	local BUILD = { classSlug = "warrior", order = { IHS, CRUELTY, IHS } }

	it("lists each cell once, in plan order, with its last point's index, rank and node", function()
		local cells = Progress.cells(BUILD, { [1] = { ["1:1"] = 1 } }, TALENTS)
		assert.are.same({
			{ tab = 1, tier = 1, column = 1, want = 2, have = 1, firstIndex = 1, lastIndex = 3, node = 101, next = false,
				runs = { { 1, 1 }, { 3, 3 } } },
			{ tab = 2, tier = 1, column = 3, want = 1, have = 0, firstIndex = 2, lastIndex = 2, node = nil, next = true,
				runs = { { 2, 2 } } },
		}, cells)
	end)

	it("marks no cell next once the build is done, and reads a missing tab as no rank", function()
		local cells = Progress.cells(BUILD, { [1] = { ["1:1"] = 2 }, [2] = { ["1:3"] = 1 } }, TALENTS)
		assert.is_false(cells[1].next)
		assert.is_false(cells[2].next)
		assert.are.equal(1, cells[2].have)
		assert.are.equal(0, Progress.cells(BUILD, {}, TALENTS)[1].have)
	end)

	it("is empty without a build, and has no nodes without data", function()
		assert.are.same({}, Progress.cells(nil, {}, TALENTS))
		assert.is_nil(Progress.cells(BUILD, {}, nil)[1].node)
	end)
end)

-- Task L3.2 fix (runs), user-reported on the hunter (hunter-13-*): the saved
-- 31/20 build takes Lethal Attacks (Marksmanship 1/3) at points 8-11 and 36.
-- The 51 points as tab, tier, column triples, decoded from the SavedVariables code.
local HUNTER_ORDER = "112x5 121x2 213x4 113x3 132 133x5 143x5 144 152 154x2 163x5 172 213 221x3 222x5 231x2 242 232x4"

--- The hunter's saved build, its order expanded from HUNTER_ORDER.
local function hunterBuild()
	local order = {}
	for triple in HUNTER_ORDER:gmatch("%S+") do
		local tab, tier, column, times = triple:match("^(%d)(%d)(%d)x?(%d*)$")
		for _ = 1, tonumber(times) or 1 do
			order[#order + 1] = cell(tonumber(tab), tonumber(tier), tonumber(column))
		end
	end
	return { classSlug = "hunter", order = order }
end

describe("Progress.runsFor", function()
	it("pins the hunter's order: 51 points, Lethal Attacks in runs 8-11 and 36", function()
		local build = hunterBuild()
		assert.are.equal(51, #build.order)
		assert.are.same({ { 8, 11 }, { 36, 36 } }, Progress.runsFor(build, 2, 1, 3))
	end)

	it("keeps a consecutive talent one run", function()
		local build = hunterBuild()
		assert.are.same({ { 1, 5 } }, Progress.runsFor(build, 1, 1, 2))
		assert.are.same({ { 37, 39 } }, Progress.runsFor(build, 2, 2, 1))
		assert.are.same({ { 47, 47 } }, Progress.runsFor(build, 2, 4, 2))
	end)

	it("splits three runs in plan order", function()
		local build = { classSlug = "warrior", order = { IHS, IHS, DEFLECTION, IHS, REND, REND, IHS, IHS } }
		assert.are.same({ { 1, 2 }, { 4, 4 }, { 7, 8 } }, Progress.runsFor(build, 1, 1, 1))
	end)

	it("is empty for a cell the plan never touches, and without a build", function()
		assert.are.same({}, Progress.runsFor(hunterBuild(), 3, 1, 1))
		assert.are.same({}, Progress.runsFor(nil, 2, 1, 3))
	end)
end)

describe("Progress.nextRun", function()
	local RUNS = { { 8, 11 }, { 36, 36 } }

	it("is the first run while any of its ranks is missing", function()
		assert.are.same({ 8, 11 }, Progress.nextRun(RUNS, 0))
		assert.are.same({ 8, 11 }, Progress.nextRun(RUNS, 3))
		assert.are.same({ 8, 11 }, Progress.nextRun(RUNS, nil))
	end)

	it("moves to the next run once a run's ranks are all taken", function()
		assert.are.same({ 36, 36 }, Progress.nextRun(RUNS, 4))
	end)

	it("is nil once every rank is taken, or overspent, and for no runs", function()
		assert.is_nil(Progress.nextRun(RUNS, 5))
		assert.is_nil(Progress.nextRun(RUNS, 6))
		assert.is_nil(Progress.nextRun({}, 0))
		assert.is_nil(Progress.nextRun(nil, 0))
	end)

	it("reads the hunter's Lethal Attacks cell from Progress.cells", function()
		local ranks = { [1] = { ["1:2"] = 5, ["2:1"] = 2 }, [2] = { ["1:3"] = 4 } }
		local lethal
		for _, entry in ipairs(Progress.cells(hunterBuild(), ranks, nil)) do
			if entry.tab == 2 and entry.tier == 1 and entry.column == 3 then
				lethal = entry
			end
		end
		assert.are.same({ { 8, 11 }, { 36, 36 } }, lethal.runs)
		assert.are.same({ 36, 36 }, Progress.nextRun(lethal.runs, lethal.have))
	end)
end)

-- Fix round 2: the plan's points the character can reach now, spent or spendable.
describe("Progress.reachablePoints", function()
	it("is the points spent plus the points waiting", function()
		assert.are.equal(12, Progress.reachablePoints(21, {}))
		assert.are.equal(12, Progress.reachablePoints(21, { [1] = { ["1:1"] = 12 } }))
		assert.are.equal(13, Progress.reachablePoints(22, { [1] = { ["1:1"] = 12 } }))
		assert.are.equal(0, Progress.reachablePoints(5, {}))
	end)
end)

-- L3.2 fix round 3: the nodes holding a point the plan never wanted there.
describe("Progress.offPlanCells", function()
	local TREES = { classes = { warrior = { tabs = {
		{ name = "Arms", talents = {
			{ name = "Improved Heroic Strike", tier = 1, column = 1, maxRank = 3, node = 101 },
			{ name = "Deflection", tier = 1, column = 2, maxRank = 5, node = 102 },
		} },
		{ name = "Fury", talents = { { name = "Cruelty", tier = 1, column = 3, maxRank = 5, node = 201 } } },
	} } } }
	local BUILD = { classSlug = "warrior", order = { IHS, IHS } }

	it("lists every ranked node the plan does not touch, with its node id", function()
		local ranks = { [1] = { ["1:1"] = 3, ["1:2"] = 1 }, [2] = { ["1:3"] = 2 } }
		assert.are.same({
			{ tab = 1, tier = 1, column = 2, have = 1, node = 102 },
			{ tab = 2, tier = 1, column = 3, have = 2, node = 201 },
		}, Progress.offPlanCells(BUILD, ranks, TREES))
	end)

	it("is empty with nothing off plan, no build or no data", function()
		assert.are.same({}, Progress.offPlanCells(BUILD, { [1] = { ["1:1"] = 2, ["1:2"] = 0 } }, TREES))
		assert.are.same({}, Progress.offPlanCells(nil, {}, TREES))
		assert.are.same({}, Progress.offPlanCells(BUILD, { [1] = { ["1:2"] = 1 } }, nil))
	end)
end)

-- L3.3 fix round 1: the two-hander rule the chip and the item tooltip share.
describe("Progress.weaponVerdict", function()
	local HANDS = { strength = 2 }
	local SMITE = { itemId = 70, equipLoc = "INVTYPE_2HWEAPON", stats = { strength = 20 } }

	it("blocks the off hand under a worn two-hander", function()
		local shield = { itemId = 80, equipLocation = "INVTYPE_SHIELD", stats = { strength = 3 } }
		local verdict = Progress.weaponVerdict("off_hand", shield, SMITE, nil, HANDS)
		assert.are.equal("blocked", verdict.state)
		assert.is_nil(verdict.delta)
	end)

	it("scores a one-hander in the main hand against the two-hander", function()
		local mace = { itemId = 81, equipLocation = "INVTYPE_WEAPON", stats = { strength = 4 } }
		local verdict = Progress.weaponVerdict("main_hand", mace, SMITE, nil, HANDS, 0.5)
		assert.are.equal("better", verdict.state)
		assert.are.equal(16, verdict.delta)
	end)

	it("scores a two-hander net of both hands", function()
		local main = { itemId = 71, equipLoc = "INVTYPE_WEAPON", stats = { strength = 6 } }
		local off = { itemId = 72, equipLoc = "INVTYPE_SHIELD", stats = { strength = 4 } }
		local axe = { itemId = 82, equipLocation = "INVTYPE_2HWEAPON", stats = { strength = 12 } }
		local verdict = Progress.weaponVerdict("main_hand", axe, main, off, HANDS)
		assert.are.equal("behind", verdict.state)
		assert.are.equal(-4, verdict.delta)
		assert.are.equal(71, verdict.equippedItemId)
	end)

	it("is the plain slot verdict for every other case", function()
		local mace = { itemId = 81, equipLocation = "INVTYPE_WEAPON", stats = { strength = 4 } }
		assert.are.same(Progress.slotVerdict("off_hand", mace, nil, HANDS),
			Progress.weaponVerdict("off_hand", mace, nil, nil, HANDS))
	end)
end)

describe("Progress.slotVerdict, by link", function()
	it("scores a second copy with a different link rather than calling it a match", function()
		local worn = { itemId = 50, link = "item:50:a", stats = { strength = 4 } }
		local copy = { itemId = 50, link = "item:50:b", stats = { strength = 6 } }
		assert.are.equal("behind", Progress.slotVerdict("chest", copy, worn, { strength = 1 }).state)
		assert.are.equal("match", Progress.slotVerdict("chest", worn, worn, { strength = 1 }).state)
		assert.are.equal("match", Progress.slotVerdict("chest", { itemId = 50, stats = {} }, worn, { strength = 1 }).state)
	end)
end)

-- L5.1 fix round 3: one target per slot, shared by the item tooltip's Next line and the gear gap
-- (Comms.gearGap). Fix round 4: that target is the band's pick (Gear.pick); a build's gear list is
-- a snapshot of equipment, never a target, so the helpers take no build.
describe("Progress.slotTargets", function()
	local TARGET_DATA = {
		bis = { ["warrior-arms"] = { [20] = { alliance = {
			chest = "100|D|The Deadmines: Defias Overseer",
			finger1 = "200|Q|A ring quest",
			finger2 = "201|Q|Another ring quest",
			off_hand = "300|V|A shield vendor",
			legs = "400|D|Gone|0.1|401|V|A trouser vendor",
		} } } },
	}
	-- Whole points: no DPS per point.
	local CURRENT = { weights = { strength = 1 }, spec = "warrior-arms" }
	local STATS = { [100] = { strength = 10 }, [110] = { strength = 30 }, [200] = { strength = 5 },
		[201] = { strength = 5 }, [300] = { strength = 3 }, [401] = { strength = 7 } }
	local function statsOf(itemId)
		return STATS[itemId]
	end
	local function never()
		return false
	end
	local function targets(equipped, current, isMissing, stats)
		local rows = Progress.slotTargets(TARGET_DATA, "warrior-arms", 20, "alliance", equipped,
			current == nil and CURRENT or current, stats or statsOf, isMissing or never)
		local bySlot = {}
		for _, row in ipairs(rows) do
			bySlot[row.slot] = row
		end
		return bySlot, rows
	end
	local CHEST = { itemId = 50, link = "item:50", equipLoc = "INVTYPE_CHEST", stats = { strength = 4 } }

	it("scores each slot's band pick over what is worn, rounded as the tooltip shows it", function()
		local bySlot = targets({ chest = CHEST })
		assert.are.same({ 100, 6 }, { bySlot.chest.itemId, bySlot.chest.gain })
		-- An empty slot is a fill: the pick's whole score.
		assert.are.same({ 200, 5 }, { bySlot.finger1.itemId, bySlot.finger1.gain })
		-- A slot with no target is left out.
		assert.is_nil(bySlot.head)
	end)

	it("ignores a build's gear list: each slot's target is the band's pick, its source and alternatives kept",
		function()
			-- Fix round 4: a build code's gear list is the equipment at export, never a plan. The
			-- helpers read no build, so a slot no band pick names has no target.
			local bySlot = targets({ chest = CHEST })
			assert.are.same({ 100, 6 }, { bySlot.chest.itemId, bySlot.chest.gain })
			assert.is_nil(bySlot.wrist)
			local chest = Progress.slotTarget(TARGET_DATA, "warrior-arms", 20, "alliance", "chest")
			assert.are.same({ 100, "D", "The Deadmines: Defias Overseer" }, { chest.item_id, chest.kind, chest.source })
			assert.is_nil(chest.planned)
			local legs = Progress.slotTarget(TARGET_DATA, "warrior-arms", 20, "alliance", "legs")
			assert.are.same({ 400, "Gone", 401, "A trouser vendor" }, { legs.item_id, legs.source,
				legs.alternatives[1].item_id, legs.alternatives[1].source })
			assert.is_nil(Progress.slotTarget(TARGET_DATA, "warrior-arms", nil, "alliance", "legs"))
			assert.is_nil(Progress.slotTarget(TARGET_DATA, "warrior-arms", 20, nil, "legs"))
			assert.is_nil(Progress.slotTarget(TARGET_DATA, nil, 20, "alliance", "legs"))
			-- The worn item's own hover reads the same target (Progress.targetFor).
			local target = Progress.targetFor(TARGET_DATA, "warrior-arms", 20, "alliance", { "chest" }, "chest", 50,
				{ chest = CHEST }, never)
			assert.are.equal(100, target.pick.item_id)
			assert.is_nil(target.isPick)
		end)

	it("holds a slot whose worn item is a pick of its pair, or whose pick is worn in the pair", function()
		local ring = { itemId = 201, link = "item:201", equipLoc = "INVTYPE_FINGER", stats = { strength = 5 } }
		local bySlot = targets({ finger1 = ring })
		assert.are.same({ 0, true }, { bySlot.finger1.gain, bySlot.finger1.holds })
		-- The empty finger2's pick is worn in finger1: nothing to go and get there.
		assert.is_nil(bySlot.finger2)
		bySlot = targets({ finger2 = { itemId = 202, link = "item:202", equipLoc = "INVTYPE_FINGER",
			stats = { strength = 1 } }, finger1 = { itemId = 201, link = "item:201", equipLoc = "INVTYPE_FINGER",
			stats = { strength = 5 } } })
		assert.are.same({ 0, true }, { bySlot.finger2.gain, bySlot.finger2.holds })
	end)

	it("leaves a hand the worn two-hander holds unscored", function()
		local axe = { itemId = 70, link = "item:70", equipLoc = "INVTYPE_2HWEAPON", stats = { strength = 20 } }
		local bySlot = targets({ main_hand = axe })
		assert.are.equal(300, bySlot.off_hand.itemId)
		assert.is_nil(bySlot.off_hand.gain)
	end)

	it("goes to the first alternative the client can show when it cannot show the pick", function()
		local bySlot = targets({}, nil, function(itemId)
			return itemId == 400
		end)
		assert.are.same({ 401, 7 }, { bySlot.legs.itemId, bySlot.legs.gain })
	end)

	it("counts a pick with no stats (not cached, no DPS in the data) or no weights as unscored", function()
		local bySlot = targets({ chest = CHEST }, nil, nil, function()
			return nil
		end)
		assert.are.equal(100, bySlot.chest.itemId)
		assert.is_nil(bySlot.chest.gain)
		local _, rows = targets({ chest = CHEST }, false)
		for _, row in ipairs(rows) do
			assert.is_nil(row.gain)
		end
	end)

	it("prices a weapon the client has not cached from Data.bisDps, as the tooltip does", function()
		local told
		local function standIn(itemId)
			told = itemId
		end
		local stats, cold = Progress.pickStats({ bisDps = { [9001] = 7.5 } }, function()
			return {}
		end, 9001, standIn)
		assert.are.same({ weapon_dps = 7.5 }, stats)
		-- L3.7 item 6: the second answer says the stats are the data's stand-in (a lower bound),
		-- so a surface that prints a gain withholds it.
		assert.is_true(cold)
		assert.are.equal(9001, told)
		told = nil
		local own, ownCold = Progress.pickStats({ bisDps = { [9001] = 7.5 } }, function()
			return { agility = 3 }
		end, 9001, standIn)
		assert.are.same({ agility = 3 }, own)
		assert.is_falsy(ownCold)
		assert.is_nil(told)
		-- Fix round 4: no build entry's own stats to fall back on: the client's, else the data's DPS.
		assert.is_nil(Progress.pickStats({}, function()
			return nil
		end, 9001, standIn))
	end)
end)

-- L3.7 pre-item (round A review): only the hovered copy is left out of the partner search, so
-- a second copy of the same one-hander in the bags can be the off-hand partner.
describe("Progress.pairPartner", function()
	local weights = { strength = 1 }
	local function sword(link)
		return { link = link, itemId = 7, equipLocation = "INVTYPE_WEAPON", stats = { strength = 5 } }
	end
	local wornMain = { itemId = 9, equipLoc = "INVTYPE_2HWEAPON", stats = { strength = 20 } }

	it("leaves out the hovered item's own link only", function()
		local hovered = sword("item:7:a")
		assert.is_nil(Progress.pairPartner(hovered, wornMain, { "item:7:a" }, weights))
		local partner = Progress.pairPartner(hovered, wornMain, { sword("item:7:a"), sword("item:7:b") }, weights)
		assert.is_not_nil(partner)
		assert.are.equal("item:7:b", partner.item.link)
	end)
end)

-- The designer's pre-item (round A review): a gap whose biggest gain comes from a weapon priced from
-- Data.bisDps (the client's stats not here yet) is a lower bound, and says so (plan.cold).
describe("Progress.gearPlan, a cold weapon", function()
	local DATA_COLD = {
		bis = { ["warrior-arms"] = { [20] = { alliance = {
			main_hand = "500|V|A weapon vendor",
			chest = "100|D|The Deadmines: Defias Overseer",
		} } } },
		bisDps = { [500] = 10 },
	}
	local CURRENT = { weights = { strength = 1 }, spec = "warrior-arms", weapon = { slot = "main_hand", weight = 1 } }
	local function never()
		return false
	end
	local function plan(chestStats, weaponStats)
		local equipped = {
			main_hand = { itemId = 60, link = "item:60", equipLoc = "INVTYPE_WEAPON", stats = { weapon_dps = 4 } },
			chest = { itemId = 50, link = "item:50", equipLoc = "INVTYPE_CHEST", stats = { strength = 4 } },
		}
		local stats = { [100] = chestStats, [500] = weaponStats }
		local rows = Progress.slotTargets(DATA_COLD, "warrior-arms", 20, "alliance", equipped, CURRENT,
			function(itemId)
				return stats[itemId]
			end, never)
		return Progress.gearPlan(DATA_COLD, "warrior-arms", 20, "alliance", equipped, CURRENT, never), rows
	end

	it("marks the row priced from bisDps as cold, and a plan whose biggest gain is that row", function()
		local gap, rows = plan({ strength = 6 }, nil)
		assert.is_true(gap.cold)
		local bySlot = {}
		for _, row in ipairs(rows) do
			bySlot[row.slot] = row
		end
		assert.is_true(bySlot.main_hand.cold)
		assert.is_nil(bySlot.chest.cold)
	end)
end)
