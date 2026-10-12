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

describe("Progress.slots", function()
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
		local rows = Progress.slots(DATA, { classSlug = "warrior", gear = gear }, ARMS_RANKS, equipped)
		assert.are.same({ "main_hand", "head", "chest", "finger1" }, slotsOf(rows))
		assert.are.equal("behind", rows[1].state)
		assert.are.equal(-10, rows[1].delta)
		assert.are.equal("match", rows[2].state)
		assert.are.equal("noplan", rows[3].state)
		assert.are.equal("empty", rows[4].state)
	end)

	it("keeps the paper-doll order for a caster", function()
		local rows = Progress.slots(DATA, { classSlug = "mage", gear = gear }, ARMS_RANKS, equipped)
		assert.are.same({ "head", "chest", "finger1", "main_hand" }, slotsOf(rows))
	end)

	it("reads every delta in DPS when a DPS per point is given", function()
		local rows = Progress.slots(DATA, { classSlug = "warrior", gear = gear }, ARMS_RANKS, equipped, 2)
		for _, row in ipairs(rows) do
			assert.are.equal("dps", row.unit)
		end
		assert.are.equal(-20, rows[1].delta)
	end)

	it("lists what is worn as noplan without a build", function()
		local rows = Progress.slots(DATA, nil, ARMS_RANKS, equipped)
		assert.are.same({ "head", "chest", "main_hand" }, slotsOf(rows))
		assert.are.equal("noplan", rows[1].state)
	end)
end)

describe("Progress.biggestGap", function()
	it("names the empty ring over a behind chest when the ring is further behind", function()
		local rows = {
			Progress.slotVerdict("chest", PLAN, { itemId = 201, stats = { strength = 5 } }, WEIGHTS),
			Progress.slotVerdict("finger1", { itemId = 400, stats = { strength = 10 } }, nil, WEIGHTS),
		}
		assert.are.same({ slot = "finger1", delta = -20, item = 400 }, Progress.biggestGap(rows))
	end)

	it("names the chest when the chest is further behind", function()
		local rows = {
			Progress.slotVerdict("chest", PLAN, { itemId = 201, stats = {} }, WEIGHTS),
			Progress.slotVerdict("finger1", { itemId = 400, stats = { strength = 5 } }, nil, WEIGHTS),
		}
		assert.are.equal("chest", Progress.biggestGap(rows).slot)
	end)

	it("is nil when nothing is behind the plan", function()
		local rows = {
			Progress.slotVerdict("head", PLAN, { itemId = 100, stats = {} }, WEIGHTS),
			Progress.slotVerdict("chest", PLAN, { itemId = 200, stats = { strength = 15 } }, WEIGHTS),
		}
		assert.is_nil(Progress.biggestGap(rows))
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
			{ tab = 1, tier = 1, column = 1, want = 2, have = 1, firstIndex = 1, lastIndex = 3, node = 101, next = false },
			{ tab = 2, tier = 1, column = 3, want = 1, have = 0, firstIndex = 2, lastIndex = 2, node = nil, next = true },
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
