-- addon/tests/raid_prep_spec.lua
-- The raid-prep reads (D14) and their FS1 sections: resistance totals, the consumables in the
-- bags, the keyring and the attunement quests. Each reader answers nil on a client without
-- its API (no section, "not tracked" on the site), and an empty list when it read nothing.
local helper = require("spec_helper")
local mock = require("wow_mock")

local DATA = {
	build = "1.60.1.70790",
	classes = { warrior = { tabs = {
		{ name = "Arms", talents = {} }, { name = "Fury", talents = {} }, { name = "Protection", talents = {} },
	} } },
	weights = {},
}

local FLASK, FOOD, HELM, KEY = "|Hitem:13510|h", "|Hitem:20452|h", "|Hitem:16866|h", "|Hitem:13704|h"

local function items()
	return {
		[FLASK] = { __itemId = 13510, __class = 0, __slot = "" },
		[FOOD] = { __itemId = 20452, __class = 0, __slot = "" },
		[HELM] = { __itemId = 16866, __class = 4, __slot = "INVTYPE_HEAD" },
		[KEY] = { __itemId = 13704, __class = 13, __slot = "" },
	}
end

local function install(overrides)
	local state = {
		talents = DATA.classes.warrior.tabs,
		class = { name = "Warrior", token = "WARRIOR" },
		race = { name = "Human", token = "Human" },
		realm = "Ashbringer",
		region = 1,
		bags = {},
		itemStats = items(),
		globals = {},
	}
	for key, value in pairs(overrides or {}) do
		state[key] = value
	end
	return mock.install(state)
end

--- The mock's container API with a stack count per slot (bag -> slot -> count).
local function withStacks(counts)
	_G.C_Container.GetContainerItemInfo = function(bag, slot)
		local count = counts[bag] and counts[bag][slot]
		return count and { stackCount = count } or nil
	end
end

local function container(Export)
	return { size = Export.containerSize, link = Export.containerLink }
end

describe("RaidPrep", function()
	local RaidPrep, Export

	before_each(function()
		RaidPrep = helper.load("RaidPrep")
		Export = helper.load("Export")
	end)

	after_each(function()
		mock.uninstall()
	end)

	describe("resistances", function()
		it("is nil on a client without UnitResistance", function()
			install()
			assert.is_nil(RaidPrep.resistances())
		end)

		it("reads each school's total, the second return, in fire nature frost shadow arcane order", function()
			local asked = {}
			install({ globals = { UnitResistance = function(unit, school)
				asked[#asked + 1] = unit .. ":" .. school
				local totals = { [2] = 85, [3] = 10, [4] = 0, [5] = 15, [6] = -5 }
				return 0, totals[school], 0, 0
			end } })
			assert.are.same({ 85, 10, 0, 15, 0 }, RaidPrep.resistances())
			assert.are.same({ "player:2", "player:3", "player:4", "player:5", "player:6" }, asked)
		end)
	end)

	describe("supplies", function()
		it("counts every Consumable-class stack in the carried bags, summed per item, and nothing else", function()
			install({ bags = { [0] = { FLASK, HELM, FOOD }, [1] = { FOOD } } })
			withStacks({ [0] = { [1] = 2, [3] = 6 }, [1] = { [1] = 14 } })
			assert.are.same({ { itemId = 13510, count = 2 }, { itemId = 20452, count = 20 } },
				RaidPrep.supplies(Export.CARRIED_BAGS, container(Export)))
		end)

		it("reads an empty list from empty bags, and one per slot without a stack count", function()
			install()
			assert.are.same({}, RaidPrep.supplies(Export.CARRIED_BAGS, container(Export)))
			install({ bags = { [0] = { FLASK } } })
			assert.are.same({ { itemId = 13510, count = 1 } }, RaidPrep.supplies(Export.CARRIED_BAGS, container(Export)))
		end)

		it("is nil on a client with no container API at all", function()
			install()
			_G.C_Container, _G.GetContainerNumSlots = nil, nil
			assert.is_nil(RaidPrep.supplies(Export.CARRIED_BAGS, container(Export)))
		end)
	end)

	describe("keyring", function()
		it("is nil on a client with no keyring bag id", function()
			install({ bags = { [-2] = { KEY } } })
			assert.is_nil(RaidPrep.keyring(container(Export)))
		end)

		it("lists the keyring's item ids, sorted and once each, from KEYRING_CONTAINER", function()
			install({ bags = { [-2] = { KEY, "|Hitem:12382|h", KEY } }, globals = { KEYRING_CONTAINER = -2 } })
			assert.are.same({ 12382, 13704 }, RaidPrep.keyring(container(Export)))
		end)

		it("takes Enum.BagIndex.Keyring when the constant is absent, and is nil for a keyring of no slots", function()
			install({ bags = { [-2] = { KEY } }, globals = { Enum = { BagIndex = { Keyring = -2 } } } })
			assert.are.same({ 13704 }, RaidPrep.keyring(container(Export)))
			install({ globals = { KEYRING_CONTAINER = -2 } })
			assert.is_nil(RaidPrep.keyring(container(Export)))
		end)
	end)

	describe("quests", function()
		it("is nil on a client with no quest-completed check", function()
			install()
			assert.is_nil(RaidPrep.quests())
		end)

		it("asks C_QuestLog about each attunement quest", function()
			install({ globals = { C_QuestLog = { IsQuestFlaggedCompleted = function(id)
				return id == 6502
			end } } })
			assert.are.same({ { questId = 6502, done = true }, { questId = 6602, done = false } }, RaidPrep.quests())
		end)

		it("falls back to the old global", function()
			install({ globals = { IsQuestFlaggedCompleted = function(id)
				return id == 6602
			end } })
			assert.are.same({ { questId = 6502, done = false }, { questId = 6602, done = true } }, RaidPrep.quests())
		end)
	end)

	describe("the export", function()
		it("writes school, supplies, locks and quests after professions, before who", function()
			install({
				bags = { [0] = { FOOD }, [-2] = { KEY } },
				globals = {
					KEYRING_CONTAINER = -2,
					UnitResistance = function(_, school)
						return 0, school == 2 and 52 or 0
					end,
					C_QuestLog = { IsQuestFlaggedCompleted = function()
						return false
					end },
				},
			})
			withStacks({ [0] = { [1] = 5 } })
			local code = assert(Export.string(DATA))
			assert.is_truthy(code:find("|level=60|school=52:0:0:0:0|supplies=20452:5|locks=13704|quests=6502:0,6602:0|who=",
				1, true), code)
		end)

		it("writes none of them where the client can answer none (the bags excepted)", function()
			install()
			local code = assert(Export.string(DATA))
			assert.is_nil(code:find("school=", 1, true))
			assert.is_nil(code:find("locks=", 1, true))
			assert.is_nil(code:find("quests=", 1, true))
			assert.is_truthy(code:find("|supplies=|", 1, true), code)
		end)
	end)
end)

describe("Codec FS1 raid-prep sections", function()
	local Codec

	before_each(function()
		Codec = helper.load("Codec")
	end)

	local HEAD = "FS1:1.60.1.70790:warrior:human:0/0/0:head=16866"

	it("round-trips every section", function()
		local build = {
			dataBuild = "1.60.1.70790", classSlug = "warrior", raceSlug = "human",
			treeRanks = { {}, {}, {} }, gearSlots = { { slot = "head", itemId = 16866 } },
			resistances = { 52, 10, 0, 15, 5 },
			supplies = { { itemId = 13510, count = 1 }, { itemId = 20452, count = 20 } },
			keyring = { 13704 },
			quests = { { questId = 6502, done = true }, { questId = 6602, done = false } },
		}
		local code = Codec.encodeFS1(build)
		assert.are.equal(HEAD .. "|school=52:10:0:15:5|supplies=13510:1,20452:20|locks=13704|quests=6502:1,6602:0", code)
		local decoded = assert(Codec.decodeFS1(code))
		assert.are.same(build.resistances, decoded.resistances)
		assert.are.same(build.supplies, decoded.supplies)
		assert.are.same(build.keyring, decoded.keyring)
		assert.are.same(build.quests, decoded.quests)
		assert.are.same({}, decoded.ignored)
	end)

	it("writes an empty section for an empty read, and reads it back as empty, not absent", function()
		local code = Codec.encodeFS1({ dataBuild = "b", classSlug = "warrior", raceSlug = "human",
			treeRanks = { {}, {}, {} }, gearSlots = {}, supplies = {}, keyring = {}, quests = {} })
		assert.are.equal("FS1:b:warrior:human:0/0/0:|supplies=|locks=|quests=", code)
		local decoded = assert(Codec.decodeFS1(code))
		assert.are.same({}, decoded.supplies)
		assert.are.same({}, decoded.keyring)
		assert.are.same({}, decoded.quests)
	end)

	it("drops a malformed entry instead of refusing the code", function()
		local code = HEAD .. "|school=1:2:3|supplies=x:1,13510:2|locks=a,13704|quests=6502:2,6602:1"
		local decoded = assert(Codec.decodeFS1(code))
		assert.is_nil(decoded.resistances)
		assert.are.same({ { itemId = 13510, count = 2 } }, decoded.supplies)
		assert.are.same({ 13704 }, decoded.keyring)
		assert.are.same({ { questId = 6602, done = true } }, decoded.quests)
	end)

	it("leaves every field nil on an export that predates them", function()
		local decoded = assert(Codec.decodeFS1(HEAD .. "|level=60"))
		assert.is_nil(decoded.resistances)
		assert.is_nil(decoded.supplies)
		assert.is_nil(decoded.keyring)
		assert.is_nil(decoded.quests)
	end)
end)
