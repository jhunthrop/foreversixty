-- addon/tests/follow_default_spec.lua
-- Follow.ensureDefault: a character with no saved build follows the data's default build for
-- its spec (contract 1, `defaultBuilds`), leveling below 60 and raid at 60, marked
-- `default = true` and never written to saved variables. Pure: the class, ranks and level are
-- handed in, and this spec never touches the client mock.
local helper = require("spec_helper")

-- A three-tree warrior, enough for Gear.specOf to name a spec.
local DATA = {
	build = "1.60.1.70291",
	classes = { warrior = { tabs = {
		{ name = "Arms", talents = {
			{ name = "Improved Heroic Strike", tier = 1, column = 1, maxRank = 3 },
			{ name = "Deflection", tier = 1, column = 2, maxRank = 5 },
		} },
		{ name = "Fury", talents = { { name = "Cruelty", tier = 1, column = 3, maxRank = 5 } } },
		{ name = "Protection", talents = { { name = "Shield Specialization", tier = 1, column = 2, maxRank = 5 } } },
	} } },
	defaultBuilds = {
		["warrior-arms"] = {
			leveling = { name = "Arms leveling", order = { { 1, 1, 1 }, { 1, 1, 1 }, { 1, 1, 2 } } },
			raid = { name = "Arms raid", order = { { 1, 1, 2 }, { 2, 1, 3 } } },
		},
		["warrior-fury"] = {
			leveling = { name = "Fury leveling", order = { { 2, 1, 3 } } },
		},
	},
}

local function cell(point)
	return { point.tab, point.tier, point.column }
end

describe("Follow.ensureDefault", function()
	local Follow
	local savedDB

	before_each(function()
		savedDB = _G.ForeverSixtyDB
		_G.ForeverSixtyDB = {}
		helper.load("Codec")
		Follow = helper.load("Follow")
	end)

	after_each(function()
		_G.ForeverSixtyDB = savedDB
	end)

	it("loads the leveling default for the spec, marked default and named from the data", function()
		local build = Follow.ensureDefault(DATA, { classSlug = "warrior", ranks = {}, level = 12 })
		assert.is_not_nil(build)
		assert.is_true(build.default)
		assert.are.equal("Arms leveling", build.name)
		assert.are.equal("warrior", build.classSlug)
		assert.are.same({ { 1, 1, 1 }, { 1, 1, 1 }, { 1, 1, 2 } }, { cell(build.order[1]), cell(build.order[2]),
			cell(build.order[3]) })
		assert.are.equal(3, #build.order)
		assert.are.equal(build, Follow.build)
	end)

	it("loads the raid default at level 60", function()
		local build = Follow.ensureDefault(DATA, { classSlug = "warrior", ranks = {}, level = 60 })
		assert.are.equal("Arms raid", build.name)
		assert.is_true(build.default)
		assert.are.same({ 2, 1, 3 }, cell(build.order[2]))
	end)

	it("follows the spec the character's points say, not the first tree", function()
		local build = Follow.ensureDefault(DATA, { classSlug = "warrior", ranks = { [2] = { ["1:3"] = 5 } }, level = 30 })
		assert.are.equal("Fury leveling", build.name)
	end)

	it("takes the class's first tree with no points spent", function()
		local build = Follow.ensureDefault(DATA, { classSlug = "warrior", ranks = {}, level = 5 })
		assert.are.equal("Arms leveling", build.name)
	end)

	it("writes nothing into saved variables", function()
		Follow.ensureDefault(DATA, { classSlug = "warrior", ranks = {}, level = 12 })
		assert.are.same({}, _G.ForeverSixtyDB)
	end)

	it("leaves a character with a saved build in any slot alone", function()
		local key = require("Export").characterKey()
		_G.ForeverSixtyDB.follows = { [key] = { active = "raid", slots = {
			pvp = { code = "FSB1:1.60.1.70291:warrior:213:", name = "Mine" },
		} } }
		Follow.build = nil
		assert.is_nil(Follow.ensureDefault(DATA, { classSlug = "warrior", ranks = {}, level = 12 }))
		assert.is_nil(Follow.build)
	end)

	it("does nothing while the data carries no default builds", function()
		local data = { build = DATA.build, classes = DATA.classes }
		assert.is_nil(Follow.ensureDefault(data, { classSlug = "warrior", ranks = {}, level = 12 }))
		assert.is_nil(Follow.build)
	end)

	it("does nothing for a spec the data has no default for", function()
		local build = Follow.ensureDefault(DATA, { classSlug = "warrior", ranks = { [3] = { ["1:2"] = 5 } }, level = 12 })
		assert.is_nil(build)
	end)

	it("does nothing for a class the data does not know", function()
		assert.is_nil(Follow.ensureDefault(DATA, { classSlug = "paladin", ranks = {}, level = 12 }))
	end)

	-- The chip's "pick a path" card (spec section 5.1): the player names the spec.
	it("loads the default for a spec the player chose, over the one the points say", function()
		local build = Follow.ensureDefault(DATA, { spec = "warrior-fury", classSlug = "warrior", ranks = {}, level = 4 })
		assert.are.equal("Fury leveling", build.name)
		assert.is_true(build.default)
		assert.are.equal(build, Follow.build)
	end)

	it("refuses a chosen spec of another class, or one with no default", function()
		assert.is_nil(Follow.ensureDefault(DATA, { spec = "mage-fire", classSlug = "warrior", ranks = {}, level = 4 }))
		assert.is_nil(Follow.ensureDefault(DATA, { spec = "warrior-protection", classSlug = "warrior", ranks = {},
			level = 4 }))
	end)

	it("still never replaces a saved build for a chosen spec", function()
		local key = require("Export").characterKey()
		_G.ForeverSixtyDB.follows = { [key] = { active = "raid", slots = {
			raid = { code = "FSB1:1.60.1.70291:warrior:213:", name = "Mine" },
		} } }
		assert.is_nil(Follow.ensureDefault(DATA, { spec = "warrior-fury", classSlug = "warrior", ranks = {}, level = 4 }))
	end)

	it("refuses a malformed order rather than raising", function()
		local data = { build = DATA.build, classes = DATA.classes, defaultBuilds = {
			["warrior-arms"] = { leveling = { name = "Broken", order = { { 1, 1 }, { 0, 1, 1 } } } },
		} }
		assert.is_nil(Follow.ensureDefault(data, { classSlug = "warrior", ranks = {}, level = 12 }))
		assert.is_nil(Follow.build)
	end)
end)
