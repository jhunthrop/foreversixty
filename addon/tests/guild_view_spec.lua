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
		_G.ForeverSixtyInbox = nil
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

	it("has no state line at all with no inbox message", function()
		start(DATA, { name = "Iron Vanguard", rankName = "Officer", rankIndex = 1 })
		local model = GuildView.summary()
		assert.is_nil(model.stateLine)
	end)

	it("shows claim state and the approval count from a matching guild inbox message", function()
		start(DATA, { name = "Iron Vanguard", rankName = "Officer", rankIndex = 1 })
		_G.ForeverSixtyInbox = { messages = {
			{ type = "guild", character = "us/ashbringer/tester", guild_name = "Iron Vanguard",
				claim_state = "claimed", pending_approvals = 3, rank = "officer" },
		} }
		local model = GuildView.summary()
		assert.are.equal(L.guildClaimClaimed .. " · " .. string.format(L.guildPendingApprovals, 3), model.stateLine)
	end)

	it("omits the approval count when there are none pending", function()
		start(DATA, { name = "Iron Vanguard", rankName = "Officer", rankIndex = 1 })
		_G.ForeverSixtyInbox = { messages = {
			{ type = "guild", character = "us/ashbringer/tester", guild_name = "Iron Vanguard",
				claim_state = "unclaimed", pending_approvals = 0, rank = "officer" },
		} }
		local model = GuildView.summary()
		assert.are.equal(L.guildClaimUnclaimed, model.stateLine)
	end)

	-- Round-3 owner ruling 8: a member sees none of the private state
	-- block, not even the claim-state half, regardless of what the
	-- message itself carries.
	it("shows nothing at all to a member, even with claim info and a nonzero approval count", function()
		start(DATA, { name = "Iron Vanguard", rankName = "Member", rankIndex = 5 })
		_G.ForeverSixtyInbox = { messages = {
			{ type = "guild", character = "us/ashbringer/tester", guild_name = "Iron Vanguard",
				claim_state = "claimed", pending_approvals = 3, rank = "member" },
		} }
		assert.is_nil(GuildView.summary().stateLine)
	end)

	it("shows the private state block to a leader, the same as an officer", function()
		start(DATA, { name = "Iron Vanguard", rankName = "Leader", rankIndex = 0 })
		_G.ForeverSixtyInbox = { messages = {
			{ type = "guild", character = "us/ashbringer/tester", guild_name = "Iron Vanguard",
				claim_state = "claimed", pending_approvals = 2, rank = "leader" },
		} }
		local model = GuildView.summary()
		assert.are.equal(L.guildClaimClaimed .. " · " .. string.format(L.guildPendingApprovals, 2), model.stateLine)
	end)

	it("ignores a guild message addressed to a different character", function()
		start(DATA, { name = "Iron Vanguard", rankName = "Officer", rankIndex = 1 })
		_G.ForeverSixtyInbox = { messages = {
			{ type = "guild", character = "us/ashbringer/someone-else", guild_name = "Iron Vanguard",
				claim_state = "claimed", pending_approvals = 3, rank = "officer" },
		} }
		assert.is_nil(GuildView.summary().stateLine)
	end)

	it("ignores a guild message about a different guild -- a stale sync after a transfer", function()
		start(DATA, { name = "Iron Vanguard", rankName = "Officer", rankIndex = 1 })
		_G.ForeverSixtyInbox = { messages = {
			{ type = "guild", character = "us/ashbringer/tester", guild_name = "Old Guild",
				claim_state = "claimed", pending_approvals = 3, rank = "officer" },
		} }
		assert.is_nil(GuildView.summary().stateLine)
	end)

	it("ignores a message of a different type entirely", function()
		start(DATA, { name = "Iron Vanguard", rankName = "Officer", rankIndex = 1 })
		_G.ForeverSixtyInbox = { messages = {
			{ type = "weights", character = "us/ashbringer/tester", spec = "arms" },
		} }
		assert.is_nil(GuildView.summary().stateLine)
	end)

	it("mounts and refreshes without error in every state", function()
		start(DATA, { name = "Iron Vanguard", rankName = "Member", rankIndex = 5 })
		local view = GuildView.mount(_G.CreateFrame("Frame"), { contentWidth = 538, select = function() end })
		view.refresh()
		_G.ForeverSixtyData = nil
		view.refresh()
	end)

	it("draws the state line onto the frame and hides it again once the message is gone", function()
		start(DATA, { name = "Iron Vanguard", rankName = "Officer", rankIndex = 1 })
		_G.ForeverSixtyInbox = { messages = {
			{ type = "guild", character = "us/ashbringer/tester", guild_name = "Iron Vanguard",
				claim_state = "claimed", pending_approvals = 1, rank = "officer" },
		} }
		local view = GuildView.mount(_G.CreateFrame("Frame"), { contentWidth = 538, select = function() end })
		assert.is_true(view.state:IsShown())
		assert.are.equal(L.guildClaimClaimed .. " · " .. string.format(L.guildPendingApprovals, 1), view.state:GetText())

		_G.ForeverSixtyInbox = nil
		view.refresh()
		assert.is_false(view.state:IsShown())
	end)
end)

-- The member standing line (wave-1 scope item 4), shown first on every
-- panel: the player's biggest gear gap (real, this addon's own
-- Gear.upgrades) and their item-level rank among guildmates (not real --
-- ForeverSixtyData's own schema carries neither item_level nor spec per
-- member yet, so this says exactly that rather than inventing a rank).
describe("GuildView.standingLine", function()
	local GuildView, Follow

	local BUILD_DATA = {
		build = "1.60.1.69893",
		classes = { paladin = { tabs = {
			{ name = "Holy", talents = {} },
			{ name = "Protection", talents = {} },
			{ name = "Retribution", talents = {} },
		} } },
		weights = { ["paladin-holy"] = { strength = 1.0 } },
	}

	local function start(overrides)
		local state = {
			class = { name = "Paladin", token = "PALADIN" },
			talents = {
				{ name = "Holy", talents = {} },
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			},
		}
		for key, value in pairs(overrides or {}) do
			state[key] = value
		end
		mock.install(state)
		helper.load("Theme").reset()
		helper.load("Widgets")
		helper.load("Cards")
		helper.load("Export")
		helper.load("Gear")
		helper.load("Talents")
		Follow = helper.load("Follow")
		GuildView = helper.load("GuildView")
	end

	after_each(function()
		mock.uninstall()
	end)

	it("says the gear gap needs a build, and that the rank is not in the data addon, with nothing loaded", function()
		start()
		local line = GuildView.standingLine(BUILD_DATA, nil)
		assert.are.equal(L.guildStandingNoRank .. " · " .. L.guildStandingGapNoBuild, line)
	end)

	it("says nothing beats the plan once a build is loaded with no upgrades waiting", function()
		start()
		local build = assert(Follow.load("FSB1:1.60.1.69893:paladin:111:head=10:strength=10", BUILD_DATA))
		local line = GuildView.standingLine(BUILD_DATA, build)
		assert.are.equal(L.guildStandingNoRank .. " · " .. L.guildStandingGapNone, line)
	end)

	it("names the real top gear gap once the bags (or what is worn) beat the plan", function()
		start({
			equipped = { [1] = "|Hitem:11|h" }, -- slot id 1 is head
			itemStats = { ["|Hitem:11|h"] = { __itemId = 11, __slot = "INVTYPE_HEAD", ITEM_MOD_STRENGTH_SHORT = 30 } },
		})
		local build = assert(Follow.load("FSB1:1.60.1.69893:paladin:111:head=10:strength=10", BUILD_DATA))
		local line = GuildView.standingLine(BUILD_DATA, build)
		assert.are.equal(L.guildStandingNoRank .. " · " .. string.format(L.guildStandingGap, "head", 20), line)
	end)

	it("wires into GuildView.summary as memberStanding, every panel's own line", function()
		mock.install({
			class = { name = "Paladin", token = "PALADIN" },
			race = { name = "Human", token = "Human" },
			realm = "Ashbringer", region = 1,
			guild = { name = "Iron Vanguard", rankName = "Member", rankIndex = 5 },
			talents = {
				{ name = "Holy", talents = {} }, { name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			},
		})
		helper.load("Theme").reset()
		helper.load("Widgets")
		helper.load("Cards")
		helper.load("Export")
		helper.load("Gear")
		helper.load("Talents")
		helper.load("Ratings")
		Follow = helper.load("Follow")
		GuildView = helper.load("GuildView")
		local model = GuildView.summary(BUILD_DATA, nil)
		assert.are.equal(L.guildStandingNoRank .. " · " .. L.guildStandingGapNoBuild, model.memberStanding)
	end)

	it("leaves memberStanding nil for a caller that passes no build data at all (back-compat)", function()
		mock.install({
			class = { name = "Paladin", token = "PALADIN" },
			guild = { name = "Iron Vanguard", rankName = "Member", rankIndex = 5 },
		})
		helper.load("Theme").reset()
		helper.load("Widgets")
		helper.load("Cards")
		helper.load("Export")
		helper.load("Gear")
		helper.load("Talents")
		helper.load("Ratings")
		Follow = helper.load("Follow")
		GuildView = helper.load("GuildView")
		assert.is_nil(GuildView.summary().memberStanding)
	end)

	it("shows the line first on the mounted page, ahead of the roster", function()
		mock.install({
			class = { name = "Paladin", token = "PALADIN" },
			race = { name = "Human", token = "Human" },
			realm = "Ashbringer", region = 1,
			guild = { name = "Iron Vanguard", rankName = "Member", rankIndex = 5 },
			talents = {
				{ name = "Holy", talents = {} }, { name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			},
		})
		helper.load("Theme").reset()
		helper.load("Widgets")
		helper.load("Cards")
		helper.load("Export")
		helper.load("Gear")
		helper.load("Talents")
		helper.load("Ratings")
		Follow = helper.load("Follow")
		GuildView = helper.load("GuildView")
		local view = GuildView.mount(_G.CreateFrame("Frame"), {
			data = BUILD_DATA, contentWidth = 538, select = function() end,
		})
		assert.is_true(view.memberStanding:IsShown())
		assert.are.equal(L.guildStandingNoRank .. " · " .. L.guildStandingGapNoBuild, view.memberStanding:GetText())
	end)
end)
