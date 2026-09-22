-- addon/tests/data_addon_spec.lua
-- The nightly data-addon job's own golden fixture
-- (tests/fixtures/data_addon_sample.lua, checked in lockstep with
-- api/internal/dataaddon's TestGoldenDataLuaIsByteExact) loaded exactly as
-- a packaged Data.lua would be, then read back through the real reader.
local helper = require("spec_helper")
local mock = require("wow_mock")

describe("the data-addon job's output, read by Ratings.lua", function()
	local Ratings

	local function start(realm, region)
		mock.install({ class = { name = "Warrior", token = "WARRIOR" }, realm = realm, region = region })
		_G.ForeverSixtyData = nil
		dofile("tests/fixtures/data_addon_sample.lua")
		Ratings = helper.load("Ratings")
	end

	after_each(function()
		_G.ForeverSixtyData = nil
		mock.uninstall()
	end)

	it("is available", function()
		start("Normal", 1) -- Export.REGION_NAMES[1] = "US"
		assert.is_true(Ratings.status().available)
	end)

	it("reads Thoradin's card with every component present", function()
		start("Normal", 1)
		local card = Ratings.forCharacter("Thoradin")
		assert.are.equal(80, card.rating)
		assert.are.equal(2, card.fights)
		local byKey = {}
		for _, c in ipairs(card.components) do
			byKey[c.key] = c.score
		end
		assert.are.equal(91, byKey.output)
		assert.are.equal(72, byKey.survival)
		assert.are.equal(84, byKey.mechanics)
		assert.are.equal(62, byKey.utility)
		assert.are.equal(94, byKey.preparation)
		assert.are.equal(91, byKey.activity)
	end)

	it("reads a non-ASCII name on another region/ruleset with only its computable components", function()
		start("Pvp", 3) -- Export.REGION_NAMES[3] = "EU"
		local card = Ratings.forCharacter("Mörk")
		assert.are.equal(88, card.rating)
		assert.are.equal(1, card.fights)
		local byKey = {}
		for _, c in ipairs(card.components) do
			byKey[c.key] = c.score
		end
		assert.are.equal(70, byKey.output)
		assert.are.equal(70, byKey.mechanics)
		assert.is_nil(byKey.survival)
	end)

	it("reads the guild's progress, nights, roster and members", function()
		start("Normal", 1)
		local guild = Ratings.forGuild("Iron Vanguard")
		assert.are.equal("2/3", guild.progress)
		assert.are.equal(2, guild.nights)
		assert.are.equal(2, guild.roster)
		assert.are.equal("o'malley", guild.members[1])
		assert.are.equal("thoradin", guild.members[2])
	end)

	it("never sees the anonymized character or the unverified guild member", function()
		start("Normal", 1)
		assert.is_nil(Ratings.forCharacter("Hiddenhero"))
		local guild = Ratings.forGuild("Iron Vanguard")
		for _, name in ipairs(guild.members) do
			assert.are_not.equal("notyet", name)
		end
	end)
end)
