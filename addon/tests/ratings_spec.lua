-- addon/tests/ratings_spec.lua
-- Reading the Forever Sixty Data addon. It may be absent, older than this
-- addon understands, stale, or missing the character asked about; every one
-- of those reads as "no data" with a reason, never as an error.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

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
			format = 1,
			generated = os.date("!%Y-%m-%dT%H:%M:%SZ", os.time() - 3600),
			build = "1.60.1.69893",
			characters = {
				["us:ashbringer:thoradin"] = {
					rating = 81, output = 90, survival = 70, mechanics = 85,
					utility = 60, preparation = 95, activity = 88, fights = 24,
				},
			},
			guilds = {
				["us:ashbringer:iron-vanguard"] = { name = "Iron Vanguard", progress = "7/8 BWL", nights = 12, roster = 31 },
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
			start(fresh({ format = 2 }))
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
		end)

		it("marks a file older than the freshness window as stale but still usable", function()
			start(fresh({ generated = "2026-01-01T00:00:00Z" }))
			local status = Ratings.status()
			assert.is_true(status.available)
			assert.is_true(status.stale)
		end)

		it("treats a malformed table as not available rather than erroring", function()
			start({ format = "one", characters = "nope" })
			assert.is_false(Ratings.status().available)
			start({})
			assert.is_false(Ratings.status().available)
		end)
	end)

	describe("keys", function()
		it("builds the site's key from region, realm and name, lowercased and slugged", function()
			start(fresh())
			assert.are.equal("us:ashbringer:thoradin", Ratings.characterKey("US", "Ashbringer", "Thoradin"))
			assert.are.equal("eu:pyrewood-village:mörk", Ratings.characterKey("EU", "Pyrewood Village", "Mörk"))
			assert.are.equal("us:ashbringer:iron-vanguard", Ratings.guildKey("US", "Ashbringer", "Iron Vanguard"))
		end)
	end)

	describe("lookups", function()
		it("finds the player's own rating from the client's region and realm", function()
			start(fresh())
			local card = Ratings.forCharacter("Thoradin")
			assert.are.equal(81, card.rating)
			assert.are.equal(24, card.fights)
			assert.are.equal(6, #card.components)
			assert.are.equal("output", card.components[1].key)
			assert.are.equal(90, card.components[1].score)
		end)

		it("finds a character on another realm when the realm is given", function()
			start(fresh({ characters = { ["us:whitemane:bob"] = { rating = 50, fights = 3 } } }))
			assert.are.equal(50, Ratings.forCharacter("Bob", "Whitemane").rating)
			assert.is_nil(Ratings.forCharacter("Bob"))
		end)

		it("answers nil for an unrated character and when the data is unavailable", function()
			start(fresh())
			assert.is_nil(Ratings.forCharacter("Nobody"))
			start(nil)
			assert.is_nil(Ratings.forCharacter("Thoradin"))
		end)

		it("finds the guild", function()
			start(fresh())
			local guild = Ratings.forGuild("Iron Vanguard")
			assert.are.equal("7/8 BWL", guild.progress)
			assert.are.equal(12, guild.nights)
			assert.is_nil(Ratings.forGuild("Nobody"))
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

	it("words a rating for a tooltip line", function()
		start(fresh())
		local line = Ratings.tooltipLine(Ratings.forCharacter("Thoradin"))
		assert.are.equal(string.format(L.ratingsTooltip, 81, 24), line)
	end)
end)
