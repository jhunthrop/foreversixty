-- addon/tests/character_key_spec.lua
-- The character key every saved record, inbox message and outbox entry is
-- addressed by (L3.3 fix round 1): the region it carries, an empty region
-- matching any, the move of a record saved under the old empty-region key,
-- and the outbox's own name. Pure: no client mock.
local Export = require("Export")
local Follow = require("Follow")
local Codec = require("Codec")

describe("Export.keyFor", function()
	it("carries the region, lowercased", function()
		assert.are.equal("us/Classic Beta PvP/Obnoxious Yell",
			Export.keyFor("Obnoxious Yell", "Classic Beta PvP", Export.regionSegment("US", nil)))
	end)

	it("reads the region from the client's name for it, else its number", function()
		assert.are.equal("eu", Export.regionSegment("EU", 1))
		assert.are.equal("us", Export.regionSegment(nil, 1))
		assert.are.equal("", Export.regionSegment(nil, 72))
		assert.are.equal("", Export.regionSegment("", nil))
	end)
end)

describe("an empty region", function()
	it("matches any region, on either side", function()
		assert.is_true(Follow.sameCharacter("/Classic Beta PvP/Obnoxious Yell", "us/Classic Beta PvP/Obnoxious Yell"))
		assert.is_true(Follow.sameCharacter("us/Classic Beta PvP/Obnoxious Yell", "/classic-beta-pvp/obnoxious-yell"))
		assert.is_false(Follow.sameCharacter("eu/Classic Beta PvP/Obnoxious Yell", "us/Classic Beta PvP/Obnoxious Yell"))
	end)

	it("lets the controller's inbox fixture reach the character", function()
		local inbox = { messages = { { type = "weights", character = "/Classic Beta PvP/Obnoxious Yell",
			spec = "warrior-arms" } } }
		assert.are.equal(1, #Codec.inboxMessages(inbox, "us/Classic Beta PvP/Obnoxious Yell", "weights"))
	end)
end)

describe("Export.adoptKey", function()
	it("moves a record saved under the old empty-region key to the new key", function()
		local db = {
			characters = { ["/Classic Beta PvP/Obnoxious Yell"] = { name = "Obnoxious Yell" } },
			follows = { ["/Classic Beta PvP/Obnoxious Yell"] = { active = "leveling" } },
		}
		Export.adoptKey(db, "us/Classic Beta PvP/Obnoxious Yell")
		assert.are.equal("Obnoxious Yell", db.characters["us/Classic Beta PvP/Obnoxious Yell"].name)
		assert.are.equal("leveling", db.follows["us/Classic Beta PvP/Obnoxious Yell"].active)
		assert.is_nil(db.characters["/Classic Beta PvP/Obnoxious Yell"])
		assert.is_nil(db.follows["/Classic Beta PvP/Obnoxious Yell"])
	end)

	it("never overwrites a record already under the new key, nor touches another character", function()
		local db = { characters = {
			["/Classic Beta PvP/Obnoxious Yell"] = { name = "old" },
			["us/Classic Beta PvP/Obnoxious Yell"] = { name = "new" },
			["/Classic Beta PvP/Bow Jackzon"] = { name = "other" },
		} }
		Export.adoptKey(db, "us/Classic Beta PvP/Obnoxious Yell")
		assert.are.equal("new", db.characters["us/Classic Beta PvP/Obnoxious Yell"].name)
		assert.are.equal("other", db.characters["/Classic Beta PvP/Bow Jackzon"].name)
	end)

	it("does nothing without a saved database", function()
		assert.has_no.errors(function() Export.adoptKey(nil, "us/x/y") end)
	end)
end)

describe("Export.outboxName", function()
	it("is the full name as who= writes it, then the realm with its spaces gone", function()
		assert.are.equal("Obnoxious Yell-ClassicBetaPvP", Export.outboxName("Obnoxious Yell", "Classic Beta PvP"))
	end)

	it("is the name alone with no realm", function()
		assert.are.equal("Obnoxious Yell", Export.outboxName("Obnoxious Yell", ""))
	end)
end)

-- L4.4 fix round 1 (D1): the site serves /character/<region>/<ruleset>/<slug> only, with the
-- ruleset one of normal|pvp|rp|hardcore and the slug the name alone (web/src/lib/characters.ts
-- parsePath, rulesetFromRealm, characterSlug).
describe("Export.siteLink", function()
	it("reads the ruleset from the realm's words as the site does, normal by default", function()
		local cases = {
			["Classic Beta PvP"] = "pvp", ["Classic Beta PvE"] = "normal", ["Classic Beta PvE 2"] = "normal",
			["Normal"] = "normal", ["RP Realm"] = "rp", ["Classic Roleplay"] = "rp", ["Hardcore"] = "hardcore",
			["Classic HC"] = "hardcore", ["Classic-Beta-PvP"] = "pvp", ["Ashbringer"] = "normal", [""] = "normal",
		}
		for realm, ruleset in pairs(cases) do
			assert.are.equal(ruleset, Export.rulesetFromRealm(realm), realm)
		end
		assert.are.equal("normal", Export.rulesetFromRealm(nil))
	end)

	it("reads the realm as a sender spells it, its words run together", function()
		local cases = {
			ClassicBetaPvP = "pvp", ClassicBetaPvE2 = "normal", ClassicBetaRP = "rp",
			ClassicBetaRoleplay = "rp", ClassicBetaHC = "hardcore", Whitemane = "normal",
		}
		for realm, ruleset in pairs(cases) do
			assert.are.equal(ruleset, Export.rulesetFromRealm(realm), realm)
		end
	end)

	it("slugs the name alone: lower case, each run of spaces one hyphen", function()
		assert.are.equal("bow-jackzon", Export.characterSlug("Bow Jackzon"))
		assert.are.equal("smoke-gouda", Export.characterSlug("  Smoke   Gouda "))
		assert.are.equal("jean-luc", Export.characterSlug("Jean-Luc"))
	end)

	it("is SITE/character/<region>/<ruleset>/<slug>, with no realm in the slug", function()
		assert.are.equal("foreversixty.gg/character/us/pvp/bow-jackzon",
			Export.siteLink("Bow Jackzon", "Classic Beta PvP", "us"))
		assert.are.equal("foreversixty.gg/character/eu/rp/smoke-gouda",
			Export.siteLink("Smoke Gouda", "Classic Roleplay", "eu"))
		assert.are.equal("foreversixty.gg/character/us/hardcore/facep-alm",
			Export.siteLink("Facep Alm", "Hardcore", nil))
		assert.are.equal("foreversixty.gg/character/us/normal/bob", Export.siteLink("Bob", "Ashbringer", ""))
	end)
end)
