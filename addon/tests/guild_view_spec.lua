-- addon/tests/guild_view_spec.lua
-- The Guild page: the player's guild from the export, its standing and the
-- roster's ratings from the data addon, and an honest state without it.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

describe("GuildView", function()
	local GuildView

	local function start(data, guild)
		mock.install({
			class = { name = "Warrior", token = "WARRIOR" },
			race = { name = "Gnome", token = "Gnome" },
			realm = "Ashbringer",
			region = 1,
			guild = guild,
		})
		_G.ForeverSixtyData = data
		helper.load("Theme").reset()
		helper.load("Widgets")
		helper.load("Cards")
		helper.load("Export")
		helper.load("Ratings")
		GuildView = helper.load("GuildView")
	end

	after_each(function()
		_G.ForeverSixtyData = nil
		mock.uninstall()
	end)

	local DATA = {
		format = 1,
		generated = os.date("!%Y-%m-%dT%H:%M:%SZ", os.time() - 600),
		build = "1.60.1.69893",
		characters = {
			["us:ashbringer:thoradin"] = { rating = 81, fights = 24 },
			["us:ashbringer:mira"] = { rating = 64, fights = 9 },
		},
		guilds = {
			["us:ashbringer:iron-vanguard"] = {
				name = "Iron Vanguard", progress = "7/8 BWL", nights = 12, roster = 31,
				members = { "thoradin", "mira", "nobody" },
			},
		},
	}

	it("says the character is not in a guild", function()
		start(DATA, nil)
		local model = GuildView.summary()
		assert.is_false(model.inGuild)
		assert.are.equal(L.guildNone, model.title)
	end)

	it("names the guild and rank from the client, even without the data addon", function()
		start(nil, { name = "Iron Vanguard", rankName = "Officer", rankIndex = 1 })
		local model = GuildView.summary()
		assert.is_true(model.inGuild)
		assert.are.equal("Iron Vanguard", model.title)
		assert.are.equal(string.format(L.guildRank, "Officer"), model.rankLine)
		assert.is_false(model.data.available)
		assert.are.equal(L.ratingsNotInstalled, model.data.reason)
		assert.are.equal(0, #model.roster)
	end)

	it("shows the standing and a roster sorted by rating when the data addon has the guild", function()
		start(DATA, { name = "Iron Vanguard", rankName = "Member", rankIndex = 5 })
		local model = GuildView.summary()
		assert.is_true(model.data.available)
		assert.are.equal("7/8 BWL", model.standing.progress)
		assert.are.equal(string.format(L.guildNights, 12), model.standing.nightsLine)
		assert.are.equal(3, #model.roster)
		assert.are.equal("thoradin", model.roster[1].name)
		assert.are.equal(81, model.roster[1].rating)
		assert.are.equal("mira", model.roster[2].name)
		-- An unrated member sits last and says so.
		assert.are.equal("nobody", model.roster[3].name)
		assert.is_nil(model.roster[3].rating)
		assert.are.equal(L.guildUnrated, model.roster[3].ratingLine)
	end)

	it("says the guild is not on the site yet when the data has no entry for it", function()
		start(DATA, { name = "Other Guild", rankName = "Member", rankIndex = 5 })
		local model = GuildView.summary()
		assert.is_true(model.data.available)
		assert.is_nil(model.standing)
		assert.are.equal(L.guildNotOnSite, model.standingLine)
	end)

	it("mounts and refreshes without error in every state", function()
		start(DATA, { name = "Iron Vanguard", rankName = "Member", rankIndex = 5 })
		local view = GuildView.mount(_G.CreateFrame("Frame"), { contentWidth = 538, select = function() end })
		view.refresh()
		_G.ForeverSixtyData = nil
		view.refresh()
	end)
end)
