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
