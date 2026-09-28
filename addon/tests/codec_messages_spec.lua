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
end)
