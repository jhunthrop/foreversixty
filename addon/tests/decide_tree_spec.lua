-- addon/tests/decide_tree_spec.lua
-- The spec question (design/specs/2026-11-08-addon-ux.md section 6, owner 2026-11-11): the addon
-- names a character's tree on its own only when the points leave no doubt. One function owns the
-- rule, Talents.decideTree: sure when every spent point sits in one tree, or when the leading tree
-- holds more than half of all spent points and leads every other tree by at least SURE_LEAD;
-- everything else, zero points included, is undecided. Gear.specOf, Follow.ensureDefault,
-- Comms.ownSpec and Inspect all read it, so a 2/2 paladin is asked instead of being locked into
-- Holy by a tie to the first tree.
local helper = require("spec_helper")
local mock = require("wow_mock")

local DATA = {
	build = "1.60.1.70291",
	classes = { paladin = { tabs = {
		{ name = "Holy", talents = { { name = "Divine Strength", tier = 1, column = 1, maxRank = 5 } } },
		{ name = "Protection", talents = { { name = "Improved Devotion Aura", tier = 1, column = 1, maxRank = 5 } } },
		{ name = "Retribution", talents = { { name = "Benediction", tier = 1, column = 2, maxRank = 5 } } },
	} } },
	weights = { ["paladin-holy"] = { intellect = 1 }, ["paladin-retribution"] = { strength = 2 } },
	defaultBuilds = {
		["paladin-holy"] = { leveling = { name = "Holy leveling", order = { { 1, 1, 1 } } } },
		["paladin-retribution"] = { leveling = { name = "Retribution leveling", order = { { 3, 1, 2 } } } },
	},
}

describe("Talents.decideTree, the one rule", function()
	local Talents

	before_each(function()
		Talents = helper.load("Talents")
	end)

	it("leans toward asking: zero points, one point, a tie and an even split are undecided", function()
		assert.are.equal(2, Talents.SURE_MIN_POINTS)
		assert.is_nil(Talents.decideTree({}))
		assert.is_nil(Talents.decideTree({ [3] = 1 }))
		assert.is_nil(Talents.decideTree({ [1] = 2, [3] = 2 }))
		assert.is_nil(Talents.decideTree({ [1] = 3, [2] = 3 }))
	end)

	it("is sure when every spent point sits in one tree, from two points", function()
		assert.are.equal(3, Talents.decideTree({ [3] = 2 }))
		assert.are.equal(2, Talents.decideTree({ [2] = 14 }))
	end)

	it("is sure only when the leader holds more than half the points and leads by SURE_LEAD", function()
		assert.are.equal(5, Talents.SURE_LEAD)
		assert.are.equal(1, Talents.decideTree({ [1] = 8, [3] = 3 }))
		assert.is_nil(Talents.decideTree({ [1] = 7, [3] = 3 }))
		assert.are.equal(3, Talents.decideTree({ [3] = 14, [1] = 1 }))
		assert.are.equal(1, Talents.decideTree({ [1] = 31, [2] = 20 }))
		-- A 10/5/5: the lead is five, but half the points sit elsewhere.
		assert.is_nil(Talents.decideTree({ [1] = 10, [2] = 5, [3] = 5 }))
		assert.is_nil(Talents.decideTree({ [1] = 20, [2] = 20, [3] = 11 }))
	end)

	it("says which trees hold points, in tab order, for the question's wording", function()
		local _, trees = Talents.decideTree({ [3] = 2, [1] = 2 })
		assert.are.same({ 1, 3 }, trees)
		local _, none = Talents.decideTree({})
		assert.are.same({}, none)
	end)

	it("counts a character's ranks into points per tab", function()
		assert.are.same({ [1] = 2, [3] = 7 },
			Talents.tabPoints({ [1] = { ["1:1"] = 2 }, [3] = { ["1:2"] = 5, ["2:1"] = 2 } }))
	end)
end)

describe("Gear.specOf under the rule", function()
	local Gear

	before_each(function()
		helper.load("Export")
		Gear = helper.load("Gear")
	end)

	it("answers nil for a tie or no points, never the first tree", function()
		assert.is_nil(Gear.specOf(DATA, "paladin", {}))
		assert.is_nil(Gear.specOf(DATA, "paladin", { [1] = { ["1:1"] = 2 }, [3] = { ["1:2"] = 2 } }))
	end)

	it("answers the tree the points name beyond doubt", function()
		assert.are.equal("paladin-retribution", Gear.specOf(DATA, "paladin", { [3] = { ["1:2"] = 2 } }))
		assert.are.equal("paladin-protection",
			Gear.specOf(DATA, "paladin", { [2] = { ["1:1"] = 5 }, [1] = { ["1:1"] = 0 } }))
	end)
end)

describe("Follow.ensureDefault under the rule", function()
	local Follow, L

	before_each(function()
		_G.ForeverSixtyDB = {}
		helper.load("Codec")
		Follow = helper.load("Follow")
		L = require("Locale")
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
	end)

	it("loads no default for an undecided character and says why", function()
		assert.is_nil(Follow.ensureDefault(DATA, { classSlug = "paladin", ranks = {}, level = 10 }))
		assert.are.equal(string.format(L.diagBuildUndecided, "Are you Holy, Protection or Retribution?"),
			Follow.noBuildReason)
		local split = { [1] = { ["1:1"] = 2 }, [3] = { ["1:2"] = 2 } }
		assert.is_nil(Follow.ensureDefault(DATA, { classSlug = "paladin", ranks = split, level = 13 }))
		assert.is_nil(Follow.build)
	end)

	it("still follows the answer the player gave, and a decisive spend", function()
		local chosen = Follow.ensureDefault(DATA, { spec = "paladin-retribution", classSlug = "paladin", ranks = {},
			level = 10 })
		assert.are.equal("Retribution leveling", chosen.name)
		Follow.build = nil
		local sure = Follow.ensureDefault(DATA, { classSlug = "paladin", ranks = { [1] = { ["1:1"] = 5 } }, level = 14 })
		assert.are.equal("Holy leveling", sure.name)
	end)

	it("words the question from the trees that hold points, else every tree", function()
		assert.are.equal("Are you Holy or Retribution?", Follow.specQuestion(DATA, "paladin",
			{ [1] = { ["1:1"] = 2 }, [3] = { ["1:2"] = 2 } }))
		assert.are.equal("Are you Holy, Protection or Retribution?",
			Follow.specQuestion(DATA, "paladin", {}))
		-- One point names nothing yet; two in one tree do.
		assert.are.equal("Are you Holy, Protection or Retribution?", Follow.specQuestion(DATA, "paladin",
			{ [2] = { ["1:1"] = 1 } }))
		assert.is_nil(Follow.specQuestion(DATA, "paladin", { [2] = { ["1:1"] = 2 } }))
		assert.is_nil(Follow.specQuestion(DATA, "paladin", { [2] = { ["1:1"] = 5 } }))
	end)
end)

describe("Comms.ownSpec and Inspect under the rule", function()
	local state

	before_each(function()
		state = mock.install({
			class = { name = "Paladin", token = "PALADIN" }, level = 13, playerName = "Uther",
			talents = {
				{ name = "Holy", talents = { { name = "Divine Strength", tier = 1, column = 1, rank = 2, maxRank = 5 } } },
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = { { name = "Benediction", tier = 1, column = 2, rank = 2, maxRank = 5 } } },
			},
		})
		helper.load("Export")
		helper.load("Talents")
		helper.load("Gear")
		helper.load("Follow")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("names no spec in the hello for an undecided character", function()
		local Comms = helper.load("Comms")
		assert.is_nil(Comms.ownSpec(DATA, nil, "paladin", { [1] = { ["1:1"] = 2 }, [3] = { ["1:2"] = 2 } }))
		assert.are.equal("paladin-retribution", Comms.ownSpec(DATA, nil, "paladin", { [3] = { ["1:2"] = 2 } }))
	end)

	it("records no spec for an inspected player whose points split, and the tree when they do not", function()
		local Inspect = helper.load("Inspect")
		local points = { 2, 0, 2 }
		_G.GetTalentTabInfo = function(tab)
			return state.talents[tab].name, "icon", points[tab]
		end
		assert.is_nil(Inspect.readSelf().spec)
		points = { 0, 0, 3 }
		assert.are.equal("paladin-retribution", Inspect.readSelf().spec)
		points = { 0, 0, 1 }
		assert.is_nil(Inspect.readSelf().spec)
		points = { 8, 3, 0 }
		assert.are.equal("paladin-holy", Inspect.readSelf().spec)
		points = { 7, 3, 0 }
		assert.is_nil(Inspect.readSelf().spec)
	end)
end)
