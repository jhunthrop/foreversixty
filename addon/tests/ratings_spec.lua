-- addon/tests/ratings_spec.lua
-- Ratings is the thin adapter over Score for the older surfaces. The data addon may be
-- absent, newer than this addon understands, stale, or missing the character asked about;
-- every one of those reads as "no data" with a reason, never as an error.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")
local Compat = require("Compat")

describe("Ratings", function()
	local Ratings

	local function start(data, install)
		mock.install(install or {
			class = { name = "Warrior", token = "WARRIOR" },
			realm = "Ashbringer",
			region = 1,
		})
		_G.ForeverSixtyData = data
		Ratings = helper.load("Ratings")
		return Ratings
	end

	after_each(function()
		_G.ForeverSixtyData = nil
		mock.uninstall()
	end)

	local function fresh(overrides)
		local data = {
			format = 2,
			generated = os.date("!%Y-%m-%d"),
			characters = {
				["Thoradin-Ashbringer"] = {
					score = 1830, tier = "epic", as_of = "2026-11-04",
					components = { gear = 412, experience = 620, parses = 540, attendance = 180, discipline = 78 },
					fights = 24,
				},
			},
			guilds = {
				["Iron Vanguard"] = { standing = 3, nights = 12, roster = 31, members = { "Thoradin-Ashbringer" } },
			},
		}
		for key, value in pairs(overrides or {}) do
			data[key] = value
		end
		return data
	end

	describe("status", function()
		it("says the data addon is not installed when the global is absent", function()
			start(nil)
			local status = Ratings.status()
			assert.is_false(status.available)
			assert.are.equal(L.ratingsNotInstalled, status.reason)
		end)

		it("refuses a format newer than it understands, naming the fix", function()
			start(fresh({ format = 3 }))
			local status = Ratings.status()
			assert.is_false(status.available)
			assert.are.equal(L.ratingsUpdateAddon, status.reason)
		end)

		it("is available with the generation date when the file is fresh", function()
			start(fresh())
			local status = Ratings.status()
			assert.is_true(status.available)
			assert.is_false(status.stale)
			assert.is_truthy(status.generatedLine:find("Data from", 1, true))
			assert.are.equal(
				string.format(L.ratingsGenerated, Compat.shortDate(os.date("!%Y-%m-%d"))),
				status.generatedLine
			)
			assert.is_nil(status.staleLine)
		end)

		it("marks a file older than the freshness window as stale but still usable", function()
			start(fresh({ generated = "2026-01-01" }))
			local status = Ratings.status()
			assert.is_true(status.available)
			assert.is_true(status.stale)
			assert.are.equal(L.ratingsStale, status.staleLine)
		end)

		it("treats a malformed table as not available rather than erroring", function()
			start({ format = "one", characters = "nope" })
			assert.is_false(Ratings.status().available)
			start({})
			assert.is_false(Ratings.status().available)
		end)
	end)

	describe("lookups", function()
		it("finds the player's own score from the client's realm", function()
			start(fresh())
			local card = Ratings.forCharacter("Thoradin")
			assert.are.equal(1830, card.rating)
			assert.are.equal(24, card.fights)
			assert.are.equal(5, #card.components)
			assert.are.equal("gear", card.components[1].key)
			assert.are.equal(L.scoreGear, card.components[1].label)
			assert.are.equal(412, card.components[1].score)
		end)

		it("finds a character on another realm when the realm is given", function()
			start(fresh({ characters = { ["Bob-Whitemane"] = { score = 500, tier = "uncommon", fights = 3 } } }))
			assert.are.equal(500, Ratings.forCharacter("Bob", "Whitemane").rating)
			assert.are.equal(500, Ratings.forCharacter("Bob-Whitemane").rating)
			assert.is_nil(Ratings.forCharacter("Bob"))
		end)

		it("answers nil for an unrated character and when the data is unavailable", function()
			start(fresh())
			assert.is_nil(Ratings.forCharacter("Nobody"))
			start(fresh({ characters = { ["Thoradin-Ashbringer"] = { score = 0, fights = 0 } } }))
			assert.is_nil(Ratings.forCharacter("Thoradin"))
			start(nil)
			assert.is_nil(Ratings.forCharacter("Thoradin"))
		end)

		it("finds the guild", function()
			start(fresh())
			local guild = Ratings.guild("Iron Vanguard")
			assert.are.equal(string.format(L.scoreGuildStanding, 3), guild.progress)
			assert.are.equal(12, guild.nights)
			assert.are.equal("Thoradin-Ashbringer", guild.members[1])
			assert.are.equal(12, Ratings.forGuild("Iron Vanguard").nights)
			assert.is_nil(Ratings.guild("Nobody"))
		end)

		it("splits a cross-realm unit name into name and realm", function()
			start(fresh())
			local name, realm = Ratings.splitUnitName("Bob-Whitemane")
			assert.are.equal("Bob", name)
			assert.are.equal("Whitemane", realm)
			name, realm = Ratings.splitUnitName("Thoradin")
			assert.are.equal("Thoradin", name)
			assert.is_nil(realm)
		end)
	end)

	it("words a rating for a tooltip line, with its tier and date", function()
		start(fresh())
		local line = Ratings.tooltipLine(Ratings.forCharacter("Thoradin"))
		assert.are.equal("Forever Score 1830 \194\183 Epic \194\183 as of 4 Nov", line)
		assert.is_nil(Ratings.tooltipLine(nil))
	end)
end)
