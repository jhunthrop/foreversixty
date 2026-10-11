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
	-- Through _G: busted gives a spec file its own environment, and
	-- Progress reads the client's globals from the real one.
	after_each(function()
		_G.C_ClassTalents, _G.C_Traits = nil, nil
	end)

	it("sums the points every tree has available", function()
		_G.C_ClassTalents = { GetActiveConfigID = function() return 7 end }
		_G.C_Traits = {
			GetConfigInfo = function() return { treeIDs = { 1, 2 } } end,
			GetTreeInfo = function(_, treeID) return { pointsAvailable = treeID } end,
		}
		assert.are.equal(3, Progress.unspentPoints())
	end)

	it("cannot tell on a client without the trait system", function()
		assert.is_nil(Progress.unspentPoints())
	end)
end)
