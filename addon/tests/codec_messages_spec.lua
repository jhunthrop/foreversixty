-- addon/tests/codec_messages_spec.lua
-- Codec.inboxMessages: Wave C's typed inbox messages, read the same way
-- Follow.inbox already reads builds -- addressed to one character, or to
-- none (site-wide) -- but filtered by `type` as well, since the inbox now
-- carries three kinds in one list.
local helper = require("spec_helper")

describe("Codec.inboxMessages", function()
	local Codec

	before_each(function()
		Codec = helper.load("Codec")
	end)

	it("keeps a message of the requested kind addressed to the current character", function()
		local usable = Codec.inboxMessages({ messages = {
			{ type = "guild", character = "us/pvp/bow-jackzon", guild_name = "Sanguine" },
		} }, "US/PvP/Bow Jackzon", "guild")
		assert.are.equal(1, #usable)
		assert.are.equal("Sanguine", usable[1].guild_name)
	end)

	it("keeps a message with no character at all -- site-wide, not addressed", function()
		local usable = Codec.inboxMessages({ messages = {
			{ type = "weights", spec = "fury" },
		} }, "US/PvP/Bow Jackzon", "weights")
		assert.are.equal(1, #usable)
	end)

	it("drops a message addressed to a different character", function()
		local usable = Codec.inboxMessages({ messages = {
			{ type = "upgrade", character = "us/pvp/someone-else", slot = "chest" },
		} }, "US/PvP/Bow Jackzon", "upgrade")
		assert.are.equal(0, #usable)
	end)

	it("drops a message of a different kind, even addressed to this character", function()
		local usable = Codec.inboxMessages({ messages = {
			{ type = "guild", character = "us/pvp/bow-jackzon", guild_name = "Sanguine" },
		} }, "US/PvP/Bow Jackzon", "upgrade")
		assert.are.equal(0, #usable)
	end)

	it("filters a mix of kinds, characters, and keeps the companion's order", function()
		local usable = Codec.inboxMessages({ messages = {
			{ type = "upgrade", character = "us/pvp/bow-jackzon", slot = "chest" },
			{ type = "guild", character = "us/pvp/bow-jackzon", guild_name = "Sanguine" },
			{ type = "upgrade", character = "us/pvp/someone-else", slot = "legs" },
			{ type = "upgrade", slot = "feet" },
		} }, "US/PvP/Bow Jackzon", "upgrade")
		assert.are.same({ "chest", "feet" }, { usable[1].slot, usable[2].slot })
	end)

	it("reports nothing for an inbox that is not there", function()
		assert.are.same({}, Codec.inboxMessages(nil, "US/PvP/Bow Jackzon", "guild"))
		assert.are.same({}, Codec.inboxMessages({}, "US/PvP/Bow Jackzon", "guild"))
		assert.are.same({}, Codec.inboxMessages({ messages = {} }, "US/PvP/Bow Jackzon", "guild"))
	end)

	it("reports nothing for a kind this addon build does not recognise", function()
		local usable = Codec.inboxMessages({ messages = {
			{ type = "future_kind", character = "us/pvp/bow-jackzon" },
		} }, "US/PvP/Bow Jackzon", "future_kind")
		assert.are.same({}, usable)
	end)

	it("matches a character key the same loose way Follow.sameCharacter does", function()
		local usable = Codec.inboxMessages({ messages = {
			{ type = "guild", character = "US/Ashbringer/Bow Jackzon", guild_name = "Sanguine" },
		} }, "us/ashbringer/bow-jackzon", "guild")
		assert.are.equal(1, #usable)
	end)

	it("matches the site's ruleset segment against the beta's realm name", function()
		local inbox = { messages = {
			{ type = "guild", character = "us/pvp/bow-jackzon", guild_name = "Sanguine" },
			{ type = "guild", character = "us/normal/bow-jackzon", guild_name = "Other" },
		} }
		local usable = Codec.inboxMessages(inbox, "us/Classic Beta PvP/Bow Jackzon", "guild")
		assert.are.equal(1, #usable)
		assert.are.equal("Sanguine", usable[1].guild_name)
	end)
end)

-- Codec.newestFirst: the inbox readers (Follow.messages, and through it
-- Gear.currentWeights for the item tooltip) want the companion's own append order reversed, since nothing on
-- the wire timestamps one message against another.
describe("Codec.newestFirst", function()
	local Codec

	before_each(function()
		Codec = helper.load("Codec")
	end)

	it("reverses the list, oldest last becomes first", function()
		local reversed = Codec.newestFirst({ { id = "a" }, { id = "b" }, { id = "c" } })
		assert.are.same({ "c", "b", "a" }, { reversed[1].id, reversed[2].id, reversed[3].id })
	end)

	it("returns a new table rather than mutating the one it was given", function()
		local original = { { id = "a" }, { id = "b" } }
		local reversed = Codec.newestFirst(original)
		assert.are.equal("a", original[1].id)
		assert.are_not.equal(original, reversed)
	end)

	it("answers an empty list for an empty list", function()
		assert.are.same({}, Codec.newestFirst({}))
	end)

	it("answers the one entry for a single-message list", function()
		local reversed = Codec.newestFirst({ { id = "only" } })
		assert.are.equal("only", reversed[1].id)
	end)
end)

describe("Codec.inboxMessages, the last raid (contract 3)", function()
	it("reads a last_raid message for the current character", function()
		local Codec = helper.load("Codec")
		local usable = Codec.inboxMessages({ messages = {
			{ type = "last_raid", character = "us/pvp/bow-jackzon", bosses = { { name = "Lucifron", percentile = 71 } } },
		} }, "US/PvP/Bow Jackzon", "last_raid")
		assert.are.equal(1, #usable)
		assert.are.equal("Lucifron", usable[1].bosses[1].name)
	end)

	it("reads the raid's name and zone key, each nil when the message lacks it", function()
		local Codec = helper.load("Codec")
		assert.are.same({ "Molten Core", "Molten Core" }, { Codec.lastRaidZone({ type = "last_raid",
			zone = "Molten Core", zone_key = "Molten Core" }) })
		assert.are.same({ "Molten Core" }, { Codec.lastRaidZone({ type = "last_raid", zone = "Molten Core" }) })
		assert.are.same({}, { Codec.lastRaidZone({ type = "last_raid", zone = "", zone_key = 7 }) })
		assert.are.same({}, { Codec.lastRaidZone({ type = "loot_values", zone = "Molten Core" }) })
		assert.are.same({}, { Codec.lastRaidZone(nil) })
	end)
end)

-- The raid_roster message (hub L9): the shared fixture is an officer's, so it carries names; a
-- member's has the counts alone. Codec.raidRosters reads them for the Guild and Officer tabs.
describe("Codec.raidRosters (contract 3)", function()
	local json = require("dkjson")
	local Codec
	local KEY = "us/normal/obnoxious"

	local function fixture()
		for _, root in ipairs({ "../docs/contracts/fixtures/", "docs/contracts/fixtures/" }) do
			local file = io.open(root .. "inbox-raid-roster.json", "rb")
			if file then
				local text = file:read("a")
				file:close()
				local decoded, _, err = json.decode(text)
				assert(decoded, err)
				return decoded
			end
		end
		error("no inbox-raid-roster.json fixture")
	end

	before_each(function()
		Codec = helper.load("Codec")
	end)

	it("reads the officer's fixture with its counts, roles, jobs and names", function()
		local rosters = Codec.raidRosters({ messages = { fixture() } }, KEY)
		assert.are.equal(1, #rosters)
		local roster = rosters[1]
		assert.are.equal(42, roster.raidId)
		assert.are.equal("Onyxia's Lair", roster.raid)
		assert.are.equal(40, roster.size)
		assert.are.same({ ["in"] = 10, late = 2, bench = 1, out = 3, none = 16 }, roster.signups)
		assert.are.equal(12, roster.coming)
		assert.are.same({ role = "healer", coming = 4, wanted = 8 }, roster.roles[2])
		assert.are.same({ held = 1, total = 2 }, roster.jobs)
		assert.are.equal(3, #roster.members)
		assert.are.same({ name = "Quenna", class = "priest", answer = "none", role = "healer" }, roster.members[3])
		assert.are.same({ "Thornhide" }, roster.assignments[1].holders)
		assert.are.same({}, roster.assignments[2].holders)
	end)

	it("gives a member's roster no names, and a role without a target no wanted", function()
		local message = fixture()
		message.members, message.assignments = nil, nil
		for _, role in ipairs(message.roles) do
			role.wanted = nil
		end
		local roster = Codec.raidRosters({ messages = { message } }, KEY)[1]
		assert.is_nil(roster.members)
		assert.is_nil(roster.assignments)
		assert.is_nil(roster.roles[1].wanted)
		assert.are.equal(2, roster.roles[1].coming)
	end)

	it("lists the soonest raid first and leaves out one it cannot read or another character's", function()
		local later, sooner, broken, other = fixture(), fixture(), fixture(), fixture()
		later.raid_id, later.starts_at = 43, "2026-12-17T02:00:00Z"
		broken.raid_id, broken.signups = 44, { ["in"] = "ten" }
		other.raid_id, other.character = 45, "us/normal/someone-else"
		local rosters = Codec.raidRosters({ messages = { later, broken, sooner, other } }, KEY)
		assert.are.equal(2, #rosters)
		assert.are.equal(42, rosters[1].raidId)
		assert.are.equal(43, rosters[2].raidId)
	end)

	it("answers an empty list with no inbox", function()
		assert.are.same({}, Codec.raidRosters(nil, KEY))
	end)
end)
