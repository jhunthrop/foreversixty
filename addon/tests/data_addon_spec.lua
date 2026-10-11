-- addon/tests/data_addon_spec.lua
-- The nightly data addon's output loaded exactly as a packaged Data.lua would be, then
-- read back through the real reader. The format 1 sample (checked in lockstep with
-- api/internal/dataaddon's TestGoldenDataLuaIsByteExact, until Lane L6 moves the job to
-- format 2) is refused with one line; the format 2 sample is read through Ratings, the
-- adapter over Score that the older surfaces still use.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

describe("the data addon's output, read by Ratings.lua", function()
	local Ratings

	local function start(sample, realm, region)
		mock.install({ class = { name = "Warrior", token = "WARRIOR" }, realm = realm, region = region })
		_G.ForeverSixtyData = nil
		dofile("tests/fixtures/" .. sample)
		Ratings = helper.load("Ratings")
	end

	after_each(function()
		_G.ForeverSixtyData = nil
		mock.uninstall()
	end)

	describe("the format 1 sample", function()
		it("is refused with the update-both-addons line, not read as if it were format 2", function()
			start("data_addon_sample.lua", "Normal", 1)
			local status = Ratings.status()
			assert.is_false(status.available)
			assert.are.equal(L.ratingsUnreadable, status.reason)
			assert.is_nil(Ratings.forCharacter("Thoradin"))
		end)
	end)

	describe("the format 2 sample", function()
		it("is available", function()
			start("data_addon_sample_v2.lua", "Normal", 1)
			assert.is_true(Ratings.status().available)
		end)

		it("reads Thoradin's card with every component present", function()
			start("data_addon_sample_v2.lua", "Normal", 1)
			local card = Ratings.forCharacter("Thoradin")
			assert.are.equal(1230, card.rating)
			assert.are.equal(2, card.fights)
			local byKey = {}
			for _, c in ipairs(card.components) do
				byKey[c.key] = c.score
			end
			assert.are.equal(400, byKey.gear)
			assert.are.equal(380, byKey.experience)
			assert.are.equal(250, byKey.parses)
			assert.are.equal(150, byKey.attendance)
			assert.are.equal(50, byKey.discipline)
		end)

		it("reads a non-ASCII name on another realm with only the components it carries", function()
			start("data_addon_sample_v2.lua", "Pvp", 3)
			local card = Ratings.forCharacter("Mörk")
			assert.are.equal(480, card.rating)
			assert.are.equal(1, card.fights)
			local byKey = {}
			for _, c in ipairs(card.components) do
				byKey[c.key] = c.score
			end
			assert.are.equal(200, byKey.gear)
			assert.are.equal(280, byKey.experience)
			assert.is_nil(byKey.parses)
		end)

		it("reads the guild's standing, nights, roster and members", function()
			start("data_addon_sample_v2.lua", "Normal", 1)
			local guild = Ratings.forGuild("Iron Vanguard")
			assert.are.equal(string.format(L.scoreGuildStanding, 3), guild.progress)
			assert.are.equal(2, guild.nights)
			assert.are.equal(2, guild.roster)
			assert.are.equal("Thoradin-Normal", guild.members[1])
		end)

		it("shows nothing for a character with no logs yet", function()
			start("data_addon_sample_v2.lua", "Normal", 1)
			assert.is_nil(Ratings.forCharacter("Newcomer"))
			assert.is_nil(Ratings.forCharacter("Hiddenhero"))
		end)
	end)
end)
