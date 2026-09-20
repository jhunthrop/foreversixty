local helper = require("spec_helper")
local mock = require("wow_mock")

local DATA = {
	build = "1.60.1.69893",
	classes = {
		paladin = {
			tabs = {
				{
					name = "Holy",
					talents = {
						{ name = "Improved Holy Strike", tier = 1, column = 1, maxRank = 2 },
						{ name = "Divine Strength", tier = 1, column = 2, maxRank = 5 },
						{ name = "Healing Light", tier = 2, column = 1, maxRank = 3 },
					},
				},
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			},
		},
	},
	weights = {},
}

describe("Talents", function()
	local Talents

	before_each(function()
		Talents = helper.load("Talents")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("reads the client's ranks keyed by tier and column", function()
		mock.install({
			talents = {
				{
					name = "Holy",
					talents = {
						{ name = "Divine Strength", tier = 1, column = 2, rank = 5, maxRank = 5 },
						{ name = "Improved Holy Strike", tier = 1, column = 1, rank = 1, maxRank = 2 },
					},
				},
			},
		})
		assert.are.same({ [1] = { ["1:2"] = 5, ["1:1"] = 1 } }, Talents.readRanks())
	end)

	it("emits ranks in Data.lua's order, not the client's index order", function()
		-- The client lists Divine Strength first; Data.lua lists Improved
		-- Holy Strike first. The export must follow Data.lua, or the site
		-- decodes the ranks onto the wrong talents.
		mock.install({
			talents = {
				{
					name = "Holy",
					talents = {
						{ name = "Divine Strength", tier = 1, column = 2, rank = 5, maxRank = 5 },
						{ name = "Improved Holy Strike", tier = 1, column = 1, rank = 1, maxRank = 2 },
						{ name = "Healing Light", tier = 2, column = 1, rank = 3, maxRank = 3 },
					},
				},
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			},
		})
		assert.are.same({ { 1, 5, 3 }, {}, {} }, Talents.treeRanks(DATA, "paladin"))
	end)

	it("reports a talent the client has no cell for as zero, not as nil", function()
		mock.install({
			talents = {
				{ name = "Holy", talents = {
					{ name = "Divine Strength", tier = 1, column = 2, rank = 2, maxRank = 5 },
				} },
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			},
		})
		-- A nil in the middle of the array would truncate the encoded tree.
		assert.are.same({ { 0, 2, 0 }, {}, {} }, Talents.treeRanks(DATA, "paladin"))
	end)

	it("finds a talent by its cell", function()
		local talent = Talents.cellOf(DATA, "paladin", 1, 2, 1)
		assert.are.equal("Healing Light", talent.name)
		assert.is_nil(Talents.cellOf(DATA, "paladin", 1, 7, 4))
	end)

	it("returns nothing for a class Data.lua does not carry", function()
		assert.is_nil(Talents.treeRanks(DATA, "shaman"))
	end)
end)
