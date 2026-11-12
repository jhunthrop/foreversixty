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

	-- Sim audit round 17: every weight is per 1% (Data.lua weights from stat-weights.json, the
	-- companion's sim), while GetItemStats gives rating points, so a hit- and crit-rated item
	-- scored 10x and 14x too high. The Fury hands of 1.60.1.70755 (items/warrior.json) under
	-- stat-weights.json warrior-fury, as GetItemStats reads them.
	describe("rating stats", function()
		local FURY = { attack_power = 1, strength = 1.9, agility = 0.7, hit = 32, crit = 18, expertise = 18 }
		local HANDS = {
			annihilation = { strength = 35, stamina = 15, crit = 14, hit = 10 },
			enforcement = { agility = 20, strength = 28, stamina = 6, hit = 10 },
			batskin = { agility = 10, stamina = 9, hit = 20 },
		}

		it("reads hit and crit rating as percent before the per-1% weight", function()
			assert.is_near(35 * 1.9 + 18 + 32, Gear.score(HANDS.annihilation, FURY), 1e-9)
			assert.is_near(28 * 1.9 + 20 * 0.7 + 32, Gear.score(HANDS.enforcement, FURY), 1e-9)
			-- 10 * 0.7 + 2% hit * 32; it was 647.
			assert.is_near(71, Gear.score(HANDS.batskin, FURY), 1e-9)
		end)

		it("puts Gauntlets of Annihilation first and Primal Batskin Gloves last", function()
			local names = { "batskin", "enforcement", "annihilation" }
			table.sort(names, function(a, b)
				return Gear.score(HANDS[a], FURY) > Gear.score(HANDS[b], FURY)
			end)
			assert.are.same({ "annihilation", "enforcement", "batskin" }, names)
		end)

		it("reads expertise, dodge, parry and block at their own factors and defense one for one", function()
			local tank = { expertise = 18, dodge = 12, parry = 15, block = 5, defense = 2 }
			local stats = { expertise = 20, dodge = 12, parry = 15, block = 5, defense = 3 }
			assert.is_near(36 + 12 + 15 + 5 + 6, Gear.score(stats, tank), 1e-9)
		end)

		it("feeds rating haste to a caster's spell haste as well as melee haste", function()
			mock.install({ itemStats = { link = { ITEM_MOD_HASTE_RATING_SHORT = 10, ITEM_MOD_SPELL_POWER_SHORT = 20 } } })
			-- stat-weights.json mage-fire: spell_haste 6 per 1%; 10 haste rating is 1%.
			local fire = { spell_power = 1, spell_haste = 6 }
			assert.is_near(20 + 6, Gear.score(Gear.statsOf("link"), fire), 1e-9)
			-- A sim's weights naming both: the rating speeds attacks and casts alike.
			assert.is_near(2 + 3, Gear.score({ haste = 10 }, { melee_haste = 2, spell_haste = 3 }), 1e-9)
		end)

		it("pins its rating table to the shared contract fixture", function()
			local text
			for _, root in ipairs({ "../docs/contracts/fixtures/", "docs/contracts/fixtures/" }) do
				local file = io.open(root .. "rating-per-percent.json", "rb")
				if file then
					text = file:read("a")
					file:close()
					break
				end
			end
			assert.is_not_nil(text)
			local Rules = helper.load("Rules")
			local count = 0
			for stat, factor in pairs(Rules.RATING_PER_PERCENT) do
				count = count + 1
				assert.is_truthy(text:find(string.format('"%s": %d[,\n]', stat, factor)), stat)
			end
			local _, inFixture = text:match('"rating_per_percent": (%b{})'):gsub('"[%w_]+": %d+', "")
			assert.are.equal(inFixture, count)
		end)
	end)

	it("reads the spec off the tree the points name beyond doubt (Talents.decideTree)", function()
		local ranks = { [1] = { ["1:1"] = 1 }, [2] = { ["1:1"] = 7 }, [3] = { ["1:1"] = 1 } }
		assert.are.equal("paladin-protection", Gear.specOf(DATA, "paladin", ranks))
		-- A lead of three over nine points does not: the player is asked.
		assert.is_nil(Gear.specOf(DATA, "paladin", { [1] = { ["1:1"] = 2 }, [2] = { ["1:1"] = 5 }, [3] = { ["1:1"] = 1 } }))
	end)

	-- L3.7 item 3: the druid's Feral Combat tree is two specs in the data, cat ("druid-feral") and
	-- bear ("druid-feral-bear"); the tab's own slug ("druid-feral-combat") has neither a default
	-- build nor weights, so a feral druid with no path card got nothing.
	describe("a tree the data splits into two specs (the feral druid)", function()
		local Real = require("Data")

		it("answers the cat spec for the Feral Combat tree on a tie or with cat points", function()
			assert.are.equal("druid-feral", Gear.specOf(Real, "druid", { [2] = { ["1:1"] = 2 } }))
			assert.are.equal("druid-feral", Gear.specOf(Real, "druid", { [2] = { ["2:3"] = 2, ["5:1"] = 2 } }))
		end)

		it("answers the bear spec when the points sit in the cells only the bear build takes", function()
			assert.are.equal("druid-feral-bear",
				Gear.specOf(Real, "druid", { [2] = { ["2:4"] = 3, ["6:1"] = 5, ["6:3"] = 5 } }))
		end)

		it("leaves the other trees, and every spec the data names by its tree, as they were", function()
			assert.are.equal("druid-balance", Gear.specOf(Real, "druid", { [1] = { ["1:1"] = 4 } }))
			assert.are.equal("druid-restoration", Gear.specOf(Real, "druid", { [3] = { ["1:1"] = 4 } }))
			assert.are.equal("rogue-combat", Gear.specOf(Real, "rogue", { [2] = { ["1:1"] = 4 } }))
			assert.are.equal("paladin-protection",
				Gear.specOf(DATA, "paladin", { [2] = { ["1:1"] = 5 } }))
		end)

		it("answers the tab's own slug when the data has no spec of that name (no defaults, no weights)", function()
			local bare = { classes = { druid = { tabs = { { name = "Balance" }, { name = "Feral Combat" } } } } }
			assert.are.equal("druid-feral-combat", Gear.specOf(bare, "druid", { [2] = { ["1:1"] = 2 } }))
		end)

		it("gives a feral druid a default build at login, the cat's by default", function()
			local Follow = helper.load("Follow")
			local build = Follow.ensureDefault(Real, { classSlug = "druid", level = 20,
				ranks = { [2] = { ["1:1"] = 2 } } })
			assert.is_not_nil(build)
			assert.are.equal("druid-feral", build.spec)
		end)
	end)

	it("names no spec on a tie or with no points: the player is asked (spec section 6.2)", function()
		local ranks = { [1] = { ["1:1"] = 3 }, [2] = { ["1:1"] = 3 }, [3] = {} }
		assert.is_nil(Gear.specOf(DATA, "paladin", ranks))
		assert.is_nil(Gear.specOf(DATA, "paladin", {}))
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

	it("leaves what the player is already wearing out of the bag upgrades", function()
		-- A worn piece that beats the plan is scored on the Gear page's own
		-- slot row; listing it under "upgrades in your bags" with an Equip
		-- button told a player to equip what they had on (seen in game).
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
		assert.are.equal(0, #upgrades)
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

	it("formats upgrade lines through L.gearUpgrade with the larger delta first", function()
		-- Every other Gear.lines example in this file is a refusal path
		-- (gearNoBuild, gearNoWeights, gearNone); none of them exercises the
		-- L.gearUpgrade format path or its argument order, so a swapped
		-- %s/%+.1f/%s or a dropped table.sort in Gear.upgrades would still
		-- pass the whole file. Gear.candidates() walks equipped slots
		-- before bags, so the equipped upgrade below (head, the smaller
		-- delta) is *found* before the bagged one (legs, the larger delta)
		-- -- only the sort in Gear.upgrades puts legs first, so this line
		-- order would not hold if that sort were removed.
		mock.install({
			class = { name = "Paladin", token = "PALADIN" },
			talents = {
				{ name = "Holy", talents = { { name = "A", tier = 1, column = 1, rank = 5, maxRank = 5 } } },
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			},
			equipped = { [1] = "wornHead" },
			bags = { [0] = { "bagLegs", "bagHead" } },
			itemStats = {
				wornHead = { ITEM_MOD_SPELL_POWER_SHORT = 25, __itemId = 101, __slot = "INVTYPE_HEAD" },
				bagHead = { ITEM_MOD_SPELL_POWER_SHORT = 25, __itemId = 102, __slot = "INVTYPE_HEAD" },
				bagLegs = { ITEM_MOD_SPELL_POWER_SHORT = 50, __itemId = 202, __slot = "INVTYPE_LEGS" },
			},
		})
		local build = { classSlug = "paladin", statsUnknown = false, gear = {
			{ slot = "head", itemId = 1, stats = { spell_power = 20 } },
			{ slot = "legs", itemId = 2, stats = { spell_power = 20 } },
		} }
		local L = require("Locale")
		assert.are.same({
			string.format(L.gearUpgrade, "legs", 30, 2),
			string.format(L.gearUpgrade, "head", 5, 1),
		}, Gear.lines(DATA, build))
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

	it("carries the item link so the Equip button has something to equip", function()
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
		-- The link, not just the id: EquipItemByName needs something it can
		-- name, and an id is not that.
		assert.are.equal("better", Gear.upgrades(DATA, build)[1].link)
	end)
end)

-- Gear.bandFor: the BiS band for a level, the one rule the tooltip and the chip both read
-- (lifted from Tooltip.lua's own highestBandAtOrBelow / bisBandFor). Pure: no client mock.
describe("Gear.bandFor", function()
	local Gear
	local SPEC_BIS = { [20] = "twenty", [30] = "thirty", [50] = "fifty" }

	before_each(function()
		Gear = helper.load("Gear")
	end)

	it("rounds the level down to the highest band the spec has", function()
		assert.are.same({ 20, "twenty" }, { Gear.bandFor(SPEC_BIS, 20) })
		assert.are.same({ 30, "thirty" }, { Gear.bandFor(SPEC_BIS, 39) })
		assert.are.same({ 50, "fifty" }, { Gear.bandFor(SPEC_BIS, 60) })
	end)

	it("skips a band the data has not generated yet", function()
		assert.are.same({ 30, "thirty" }, { Gear.bandFor(SPEC_BIS, 45) })
	end)

	it("answers nothing below 20, with no table, or with no level", function()
		assert.is_nil(Gear.bandFor(SPEC_BIS, 19))
		assert.is_nil(Gear.bandFor(nil, 30))
		assert.is_nil(Gear.bandFor(SPEC_BIS, nil))
		assert.is_nil(Gear.bandFor({ [40] = "forty" }, 30))
	end)
end)

-- Gear.firstBand (hotfix 2.0.2): the lowest band the spec has, for a character from level 1
-- who is below it -- the "picks start at level N" rung the gear surfaces point at.
describe("Gear.firstBand", function()
	local Gear
	local SPEC_BIS = { [30] = "thirty", [20] = "twenty", [50] = "fifty" }

	before_each(function()
		Gear = helper.load("Gear")
	end)

	it("answers the lowest band from level 1 up to the band itself", function()
		assert.are.equal(20, Gear.firstBand(SPEC_BIS, 1))
		assert.are.equal(20, Gear.firstBand(SPEC_BIS, 6))
		assert.are.equal(20, Gear.firstBand(SPEC_BIS, 9))
		assert.are.equal(20, Gear.firstBand(SPEC_BIS, 10))
		assert.are.equal(20, Gear.firstBand(SPEC_BIS, 18))
		assert.are.equal(20, Gear.firstBand(SPEC_BIS, 19))
		assert.are.equal(30, Gear.firstBand({ [30] = "thirty" }, 25))
	end)

	it("answers nothing in or past the first band, or with no table, level or band", function()
		assert.is_nil(Gear.firstBand(SPEC_BIS, 20))
		assert.is_nil(Gear.firstBand(SPEC_BIS, 45))
		assert.is_nil(Gear.firstBand(nil, 15))
		assert.is_nil(Gear.firstBand(SPEC_BIS, nil))
		assert.is_nil(Gear.firstBand(SPEC_BIS, "6"))
		assert.is_nil(Gear.firstBand({}, 15))
	end)
end)
