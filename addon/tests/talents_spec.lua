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
						{ name = "Improved Holy Strike", tier = 1, column = 1, maxRank = 2, node = 101 },
						{ name = "Divine Strength", tier = 1, column = 2, maxRank = 5, node = 102 },
						{ name = "Healing Light", tier = 2, column = 1, maxRank = 3, node = 103 },
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

	describe("on the 1.60 client, whose trees are trait nodes", function()
		it("reads each rank by Data.lua's node id and keys it by the cell", function()
			mock.install({
				class = { name = "Paladin", token = "PALADIN" },
				traits = { configID = 7, ranks = { [102] = 5, [101] = 1 } },
			})
			assert.is_nil(GetNumTalentTabs)
			assert.are.same({ { ["1:1"] = 1, ["1:2"] = 5, ["2:1"] = 0 }, {}, {} }, Talents.readRanks(DATA))
			assert.are.same({ { 1, 5, 0 }, {}, {} }, Talents.treeRanks(DATA, "paladin"))
		end)

		it("reads a character with no talent config yet as unranked, not as an error", function()
			-- Below the talent level the client has no active config; the
			-- first beta character to run /fs was level 8.
			mock.install({
				class = { name = "Paladin", token = "PALADIN" },
				traits = { configID = nil, ranks = {} },
			})
			assert.are.same({ { ["1:1"] = 0, ["1:2"] = 0, ["2:1"] = 0 }, {}, {} }, Talents.readRanks(DATA))
		end)

		it("reads a talent without a node id as unranked", function()
			mock.install({
				class = { name = "Paladin", token = "PALADIN" },
				traits = { configID = 7, ranks = { [102] = 5 } },
			})
			local data = { classes = { paladin = { tabs = {
				{ name = "Holy", talents = { { name = "Old", tier = 1, column = 1, maxRank = 1 } } },
			} } } }
			assert.are.same({ { ["1:1"] = 0 } }, Talents.readRanks(data))
		end)

		it("reads nothing for a class Data.lua does not carry", function()
			mock.install({
				class = { name = "Shaman", token = "SHAMAN" },
				traits = { configID = 7, ranks = {} },
			})
			assert.are.same({}, Talents.readRanks(DATA))
		end)
	end)

	it("reads nothing, and does not raise, on a client with neither talent API", function()
		mock.install({})
		_G.GetNumTalentTabs = nil
		assert.are.same({}, Talents.readRanks(DATA))
	end)

	it("finds a talent by its cell", function()
		local talent = Talents.cellOf(DATA, "paladin", 1, 2, 1)
		assert.are.equal("Healing Light", talent.name)
		assert.is_nil(Talents.cellOf(DATA, "paladin", 1, 7, 4))
	end)

	it("returns nothing for a class Data.lua does not carry", function()
		assert.is_nil(Talents.treeRanks(DATA, "shaman"))
	end)
	-- The Talents page shows each talent's icon. The data file carries none,
	-- so it is asked of the client: node -> entry -> definition -> spell ->
	-- texture. Nobody can run the game here, so every link in that chain is
	-- allowed to be missing, and a missing one means "no icon", never an error.
	describe("iconFor", function()
		local function withTraits(traits)
			mock.install({
				class = { name = "Warrior", token = "WARRIOR" },
				traits = { configID = 7, ranks = {} },
			})
			_G.C_Traits.GetNodeInfo = function(_, node)
				return traits.nodes[node]
			end
			_G.C_Traits.GetEntryInfo = traits.GetEntryInfo
			_G.C_Traits.GetDefinitionInfo = traits.GetDefinitionInfo
			_G.C_Spell = traits.C_Spell
			_G.GetSpellTexture = traits.GetSpellTexture
			return helper.load("Talents")
		end

		local CHAIN = {
			nodes = { [100] = { entryIDs = { 11 }, activeEntry = { entryID = 11 } } },
			GetEntryInfo = function(_, entry)
				return entry == 11 and { definitionID = 22 } or nil
			end,
			GetDefinitionInfo = function(definition)
				return definition == 22 and { spellID = 12294 } or nil
			end,
			C_Spell = { GetSpellTexture = function(spell)
				return spell == 12294 and 132355 or nil
			end },
		}

		after_each(function()
			_G.C_Spell, _G.GetSpellTexture = nil, nil
			mock.uninstall()
		end)

		it("walks node, entry, definition and spell to a texture", function()
			assert.are.equal(132355, withTraits(CHAIN).iconFor({ node = 100 }))
		end)

		it("prefers the definition's own icon when it carries one", function()
			local chain = {}
			for key, value in pairs(CHAIN) do
				chain[key] = value
			end
			chain.GetDefinitionInfo = function()
				return { spellID = 12294, overrideIcon = 999 }
			end
			assert.are.equal(999, withTraits(chain).iconFor({ node = 100 }))
		end)

		it("uses the global spell texture function on a client without C_Spell", function()
			local chain = {}
			for key, value in pairs(CHAIN) do
				chain[key] = value
			end
			chain.C_Spell = nil
			chain.GetSpellTexture = function()
				return 4242
			end
			assert.are.equal(4242, withTraits(chain).iconFor({ node = 100 }))
		end)

		it("answers nil when any link of the chain is missing or raises", function()
			assert.is_nil(withTraits(CHAIN).iconFor({ node = 555 }))
			assert.is_nil(withTraits(CHAIN).iconFor({}))
			assert.is_nil(withTraits(CHAIN).iconFor(nil))
			local broken = { nodes = CHAIN.nodes, GetEntryInfo = function()
				error("not on this client")
			end }
			assert.is_nil(withTraits(broken).iconFor({ node = 100 }))
			local none = { nodes = CHAIN.nodes }
			assert.is_nil(withTraits(none).iconFor({ node = 100 }))
		end)

		it("asks the client once per node", function()
			local calls = 0
			local chain = {}
			for key, value in pairs(CHAIN) do
				chain[key] = value
			end
			chain.GetEntryInfo = function(_, entry)
				calls = calls + 1
				return CHAIN.GetEntryInfo(nil, entry)
			end
			local loaded = withTraits(chain)
			loaded.iconFor({ node = 100 })
			loaded.iconFor({ node = 100 })
			assert.are.equal(1, calls)
		end)

		it("answers nil on a client with no trait system at all", function()
			mock.install({ class = { name = "Warrior", token = "WARRIOR" } })
			assert.is_nil(helper.load("Talents").iconFor({ node = 100 }))
		end)
	end)
end)
