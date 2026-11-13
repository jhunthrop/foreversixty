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
		it("lists the keyring's item ids, sorted and once each, from KEYRING_CONTAINER", function()
			install({ bags = { [-2] = { KEY, "|Hitem:12382|h", KEY } }, globals = { KEYRING_CONTAINER = -2 } })
			assert.are.same({ 12382, 13704 }, RaidPrep.keyring(Export.CARRIED_BAGS, container(Export)))
		end)

		it("takes Enum.BagIndex.Keyring when the constant is absent", function()
			install({ bags = { [-2] = { KEY } }, globals = { Enum = { BagIndex = { Keyring = -2 } } } })
			assert.are.same({ 13704 }, RaidPrep.keyring(Export.CARRIED_BAGS, container(Export)))
		end)

		it("adds the Key-class items in the bags, and only those, keyring or none", function()
			install({ bags = { [0] = { KEY, FLASK, HELM } } })
			assert.are.same({ 13704 }, RaidPrep.keyring(Export.CARRIED_BAGS, container(Export)))
			install()
			assert.are.same({}, RaidPrep.keyring(Export.CARRIED_BAGS, container(Export)))
		end)

		it("is nil on a client with no container API at all", function()
			install()
			_G.C_Container, _G.GetContainerNumSlots = nil, nil
			assert.is_nil(RaidPrep.keyring(Export.CARRIED_BAGS, container(Export)))
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
			assert.are.same(
				{ { questId = 6502, done = true }, { questId = 6602, done = false } },
				RaidPrep.quests())
		end)

		it("falls back to the old global", function()
			install({ globals = { IsQuestFlaggedCompleted = function(id)
				return id == 6602
			end } })
			assert.are.same(
				{ { questId = 6502, done = false }, { questId = 6602, done = true } },
				RaidPrep.quests())
		end)
	end)

	describe("owned", function()
		local AMULET = "|Hitem:16309|h"

		it("says whether the Drakefire Amulet is worn or in the bags", function()
			install({ itemStats = { [AMULET] = { __itemId = 16309, __class = 4, __slot = "INVTYPE_NECK" } } })
			assert.are.same({ { itemId = 16309, held = false } }, RaidPrep.owned(Export.CARRIED_BAGS, container(Export)))
			install({ equipped = { [2] = AMULET },
				itemStats = { [AMULET] = { __itemId = 16309, __class = 4, __slot = "INVTYPE_NECK" } } })
			assert.are.same({ { itemId = 16309, held = true } }, RaidPrep.owned(Export.CARRIED_BAGS, container(Export)))
			install({ bags = { [3] = { AMULET } },
				itemStats = { [AMULET] = { __itemId = 16309, __class = 4, __slot = "INVTYPE_NECK" } } })
			assert.are.same({ { itemId = 16309, held = true } }, RaidPrep.owned(Export.CARRIED_BAGS, container(Export)))
		end)

		it("is nil on a client with no container API at all", function()
			install()
			_G.C_Container, _G.GetContainerNumSlots = nil, nil
			assert.is_nil(RaidPrep.owned(Export.CARRIED_BAGS, container(Export)))
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
			local want = "|level=60|school=52:0:0:0:0|supplies=20452:5|locks=13704"
				.. "|quests=6502:0,6602:0|owned=16309:0|who="
			assert.is_truthy(code:find(want,
				1, true), code)
		end)

		it("writes no school or quests where the client cannot answer, and empty bag reads", function()
			install()
			local code = assert(Export.string(DATA))
			assert.is_nil(code:find("school=", 1, true))
			assert.is_nil(code:find("quests=", 1, true))
			assert.is_truthy(code:find("|supplies=|locks=|owned=16309:0|", 1, true), code)
		end)
	end)

	-- The hub's Guild tab, Your prep (spec 2026-11-13-addon-hub 4.5): what tonight's raid asks of you.
	describe("prep for a raid", function()
		local NEEDS = {
			encounterNeeds = {
				[10184] = { resist = "fire", consumables = { "fire_protection" } },
				[11583] = { item = 15138 },
				[12118] = { resist = "fire", consumables = { "fire_protection" } },
			},
			consumables = { [13510] = "flask", [20452] = "food", [13457] = "fire_protection", [9187] = "elixir" },
		}

		it("asks Onyxia's Lair for the amulet, fire resistance and a fire potion, by the zone's bosses", function()
			install()
			local needs = RaidPrep.needsOf(NEEDS, "Onyxia's Lair")
			assert.are.same({}, needs.items)
			assert.are.same({ "fire" }, needs.resists)
			assert.are.same({ "fire_protection" }, needs.kinds)
			local rows = RaidPrep.prep({ data = NEEDS, zone = "Onyxia's Lair", where = { [16309] = "bags" },
				resistances = { 42, 10, 0, 15, 0 }, counts = { flask = 1, food = 20 } })
			assert.are.same({ kind = "door", itemId = 16309, state = "bags" }, rows[1])
			assert.are.same({ kind = "resist", school = "fire", value = 42 }, rows[2])
			assert.are.same({ kind = "carried", counts = { { kind = "flask", count = 1 }, { kind = "food", count = 20 } } },
				rows[3])
			assert.are.same({ kind = "missing", needs = { { kinds = { "fire_protection" }, itemId = 13457 } } }, rows[4])
		end)

		it("asks Blackwing Lair for the Onyxia Scale Cloak for Nefarian, worn or not", function()
			install()
			local rows = RaidPrep.prep({ data = NEEDS, zone = "Blackwing Lair", where = {}, counts = { elixir = 2, food = 5 } })
			assert.are.same({ kind = "item", itemId = 15138, npc = 11583, state = "missing" }, rows[1])
			assert.are.same({ kind = "packed" }, rows[3])
			assert.are.equal("Nefarian", RaidPrep.BOSS_NAMES[11583])
			rows = RaidPrep.prep({ data = NEEDS, zone = "Blackwing Lair", where = { [15138] = "worn" } })
			assert.are.same({ { kind = "item", itemId = 15138, npc = 11583, state = "worn" } }, rows)
		end)

		it("counts a flask or an elixir as one need and buff food as the other", function()
			install()
			local rows = RaidPrep.prep({ data = NEEDS, zone = "Molten Core", counts = {} })
			assert.are.same({ kind = "missing", needs = { { kinds = { "flask", "elixir" } }, { kinds = { "food" } },
				{ kinds = { "fire_protection" }, itemId = 13457 } } }, rows[1])
		end)

		it("leaves out what the client could not read, and says nothing for a raid it has no needs for", function()
			install()
			assert.are.same({}, RaidPrep.prep({ data = NEEDS, zone = "Onyxia's Lair" }))
			assert.are.same({}, RaidPrep.prep({ data = NEEDS, zone = "Zul'Gurub", where = {}, counts = {} }))
		end)

		it("reads the bags, the gear and the resistances for the summary, worn before the bags", function()
			local AMULET, POTION = "|Hitem:16309|h", "|Hitem:13457|h"
			install({ bags = { [0] = { AMULET, POTION, FLASK } }, itemStats = {
				[AMULET] = { __itemId = 16309, __class = 4, __slot = "INVTYPE_NECK" },
				[POTION] = { __itemId = 13457, __class = 0, __slot = "" },
				[FLASK] = { __itemId = 13510, __class = 0, __slot = "" },
			}, globals = { UnitResistance = function(_, school)
				return 0, school == 2 and 55 or 0
			end } })
			local rows = RaidPrep.summary("Onyxia's Lair", NEEDS, Export.CARRIED_BAGS, container(Export))
			assert.are.same({ kind = "door", itemId = 16309, state = "bags" }, rows[1])
			assert.are.same({ kind = "resist", school = "fire", value = 55 }, rows[2])
			assert.are.same({ kind = "carried", counts = { { kind = "flask", count = 1 }, { kind = "fire_protection",
				count = 1 } } }, rows[3])
			assert.are.same({ kind = "missing", needs = { { kinds = { "food" } } } }, rows[4])
			assert.are.same({ fire_protection = 1, flask = 1 }, RaidPrep.kindCounts({ { itemId = 13457, count = 1 },
				{ itemId = 13510, count = 1 }, { itemId = 1, count = 9 } }, NEEDS.consumables))
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
			owned = { { itemId = 16309, held = true } },
		}
		local code = Codec.encodeFS1(build)
		assert.are.equal(HEAD .. "|school=52:10:0:15:5|supplies=13510:1,20452:20|locks=13704"
			.. "|quests=6502:1,6602:0|owned=16309:1",
			code)
		local decoded = assert(Codec.decodeFS1(code))
		assert.are.same(build.resistances, decoded.resistances)
		assert.are.same(build.supplies, decoded.supplies)
		assert.are.same(build.keyring, decoded.keyring)
		assert.are.same(build.quests, decoded.quests)
		assert.are.same(build.owned, decoded.owned)
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
