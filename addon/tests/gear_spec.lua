local helper = require("spec_helper")
local mock = require("wow_mock")

local DATA = {
	build = "1.60.1.69893",
	classes = {
		paladin = {
			tabs = {
				{ name = "Holy", talents = { { name = "A", tier = 1, column = 1, maxRank = 5 } } },
				{ name = "Protection", talents = { { name = "B", tier = 1, column = 1, maxRank = 5 } } },
				{ name = "Retribution", talents = { { name = "C", tier = 1, column = 1, maxRank = 5 } } },
			},
		},
	},
	weights = {
		["paladin-holy"] = { spell_power = 1.0, intellect = 0.45 },
		["paladin-protection"] = { attack_power = 1.0, stamina = 3.0 },
		["paladin-retribution"] = { attack_power = 1.0, strength = 2.0 },
	},
}

describe("Gear", function()
	local Gear

	before_each(function()
		Gear = helper.load("Gear")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("maps GetItemStats keys onto the contract vocabulary", function()
		mock.install({
			itemStats = {
				link = {
					ITEM_MOD_STRENGTH_SHORT = 18,
					ITEM_MOD_CRIT_RATING_SHORT = 1,
					RESISTANCE0_NAME = 400,
					__itemId = 1,
					__slot = "INVTYPE_HEAD",
				},
			},
		})
		assert.are.same({ strength = 18, crit = 1, armor = 400 }, Gear.statsOf("link"))
	end)

	it("leaves a stat key it does not know out rather than guessing", function()
		mock.install({ itemStats = { link = { ITEM_MOD_CR_SPEED_SHORT = 5 } } })
		assert.are.same({}, Gear.statsOf("link"))
	end)

	it("scores an item GetItemStats knows nothing about as no stats at all", function()
		mock.install({ itemStats = {} })
		assert.are.same({}, Gear.statsOf("missing"))
	end)

	it("scores by the sum of weight times stat", function()
		assert.are.equal(23 * 1.0 + 10 * 0.45, Gear.score({ spell_power = 23, intellect = 10 }, DATA.weights["paladin-holy"]))
	end)

	it("scores a stat with no weight as zero rather than refusing the item", function()
		assert.are.equal(0, Gear.score({ dodge = 9 }, DATA.weights["paladin-holy"]))
	end)

	it("reads the spec off the tree with the most points", function()
		local ranks = { [1] = { ["1:1"] = 2 }, [2] = { ["1:1"] = 5 }, [3] = { ["1:1"] = 1 } }
		assert.are.equal("paladin-protection", Gear.specOf(DATA, "paladin", ranks))
	end)

	it("resolves a tie to the first tree, as the design says", function()
		local ranks = { [1] = { ["1:1"] = 3 }, [2] = { ["1:1"] = 3 }, [3] = {} }
		assert.are.equal("paladin-holy", Gear.specOf(DATA, "paladin", ranks))
	end)

	it("lists a bagged item that beats the planned item in its slot", function()
		mock.install({
			class = { name = "Paladin", token = "PALADIN" },
			talents = {
				{ name = "Holy", talents = { { name = "A", tier = 1, column = 1, rank = 5, maxRank = 5 } } },
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			},
			bags = { [0] = { "better" } },
			itemStats = {
				better = { ITEM_MOD_SPELL_POWER_SHORT = 40, __itemId = 2, __slot = "INVTYPE_HEAD" },
			},
		})
		local build = { classSlug = "paladin", statsUnknown = false, gear = {
			{ slot = "head", itemId = 1, stats = { spell_power = 20 } },
		} }
		local upgrades = Gear.upgrades(DATA, build)
		assert.are.equal(1, #upgrades)
		assert.are.equal("head", upgrades[1].slot)
		assert.are.equal(2, upgrades[1].itemId)
		assert.are.equal(20, upgrades[1].delta)
	end)

	it("scores what the player is already wearing, not only the bags", function()
		-- The design scores equipped items too: one already on the character
		-- can beat the planned one, and hiding that tells a player to swap
		-- something they should keep.
		mock.install({
			class = { name = "Paladin", token = "PALADIN" },
			talents = {
				{ name = "Holy", talents = { { name = "A", tier = 1, column = 1, rank = 5, maxRank = 5 } } },
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			},
			equipped = { [1] = "worn" },
			itemStats = {
				worn = { ITEM_MOD_SPELL_POWER_SHORT = 35, __itemId = 3, __slot = "INVTYPE_HEAD" },
			},
		})
		local build = { classSlug = "paladin", statsUnknown = false, gear = {
			{ slot = "head", itemId = 1, stats = { spell_power = 20 } },
		} }
		local upgrades = Gear.upgrades(DATA, build)
		assert.are.equal(1, #upgrades)
		assert.are.equal(3, upgrades[1].itemId)
		assert.are.equal(15, upgrades[1].delta)
	end)

	it("says nothing about a slot whose planned item has no stats", function()
		-- An FS1 code carries no stats. Scoring the planned item at zero
		-- would make every bag item an upgrade, which is worse than silence.
		mock.install({
			class = { name = "Paladin", token = "PALADIN" },
			talents = {
				{ name = "Holy", talents = { { name = "A", tier = 1, column = 1, rank = 5, maxRank = 5 } } },
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			},
			bags = { [0] = { "better" } },
			itemStats = { better = { ITEM_MOD_SPELL_POWER_SHORT = 40, __itemId = 2, __slot = "INVTYPE_HEAD" } },
		})
		local build = { classSlug = "paladin", statsUnknown = true, gear = {
			{ slot = "head", itemId = 1, stats = {} },
		} }
		assert.are.same({}, Gear.upgrades(DATA, build))
	end)

	it("finds nothing when everything carried scores no better than what is worn", function()
		-- A different path from the empty-stats case above: here the planned
		-- item has real stats and is simply not beaten, so every candidate's
		-- delta is zero or negative and the `delta > 0` guard drops it. This
		-- is the case gearNone exists for.
		mock.install({
			class = { name = "Paladin", token = "PALADIN" },
			talents = {
				{ name = "Holy", talents = { { name = "A", tier = 1, column = 1, rank = 5, maxRank = 5 } } },
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			},
			equipped = { [1] = "worse" },
			itemStats = {
				worse = { ITEM_MOD_SPELL_POWER_SHORT = 5, __itemId = 9, __slot = "INVTYPE_HEAD" },
			},
		})
		local build = { classSlug = "paladin", statsUnknown = false, gear = {
			{ slot = "head", itemId = 1, stats = { spell_power = 20 } },
		} }
		assert.are.same({}, Gear.upgrades(DATA, build))
		assert.are.same({ require("Locale").gearNone }, Gear.lines(DATA, build))
	end)

	it("says so rather than scoring when the spec has no weights", function()
		mock.install({
			class = { name = "Rogue", token = "ROGUE" },
			talents = {},
		})
		local lines = Gear.lines(DATA, { classSlug = "rogue", gear = {} })
		assert.are.same({ require("Locale").gearNoWeights }, lines)
	end)

	it("says so when nothing is loaded", function()
		assert.are.same({ require("Locale").gearNoBuild }, Gear.lines(DATA, nil))
	end)
end)
