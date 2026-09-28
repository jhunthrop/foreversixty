local helper = require("spec_helper")
local mock = require("wow_mock")

local DATA = {
	build = "1.60.1.69893",
	classes = {
		paladin = {
			tabs = {
				{ name = "Holy", talents = {
					{ name = "Improved Holy Strike", tier = 1, column = 1, maxRank = 2 },
					{ name = "Divine Strength", tier = 1, column = 2, maxRank = 5 },
					{ name = "Healing Light", tier = 2, column = 1, maxRank = 3 },
				} },
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			},
		},
	},
	weights = {},
}

-- Two points into 1:1, then one into 2:1.
local CODE = "FSB1:1.60.1.69893:paladin:111111121:"

describe("Follow", function()
	local Follow

	before_each(function()
		Follow = helper.load("Follow")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("takes the first point when nothing is spent", function()
		local build = assert(Follow.load(CODE, DATA))
		local point = Follow.nextPoint(build, { [1] = {} })
		assert.are.same({ tab = 1, tier = 1, column = 1, index = 1 }, point)
	end)

	it("skips a point the player already has", function()
		local build = assert(Follow.load(CODE, DATA))
		local point = Follow.nextPoint(build, { [1] = { ["1:1"] = 1 } })
		assert.are.same({ tab = 1, tier = 1, column = 1, index = 2 }, point)
	end)

	it("moves on once a talent is at the rank the build wants", function()
		local build = assert(Follow.load(CODE, DATA))
		local point = Follow.nextPoint(build, { [1] = { ["1:1"] = 2 } })
		assert.are.same({ tab = 1, tier = 2, column = 1, index = 3 }, point)
	end)

	it("counts a rank the player has beyond the build as spent, not as a gap", function()
		-- Overspending 1:1 to three ranks must not make the build ask for a
		-- fourth; the build wanted two and the player has them.
		local build = assert(Follow.load(CODE, DATA))
		local point = Follow.nextPoint(build, { [1] = { ["1:1"] = 3 } })
		assert.are.same({ tab = 1, tier = 2, column = 1, index = 3 }, point)
	end)

	it("reports nothing left when the build is finished", function()
		local build = assert(Follow.load(CODE, DATA))
		assert.is_nil(Follow.nextPoint(build, { [1] = { ["1:1"] = 2, ["2:1"] = 1 } }))
	end)

	it("names the talent and its tree in the line", function()
		local build = assert(Follow.load(CODE, DATA))
		local line = Follow.line(DATA, build, { [1] = {} })
		assert.is_truthy(line:find("Improved Holy Strike", 1, true))
		assert.is_truthy(line:find("Holy", 1, true))
	end)

	it("says so when the build is done rather than showing a blank line", function()
		local build = assert(Follow.load(CODE, DATA))
		local line = Follow.line(DATA, build, { [1] = { ["1:1"] = 2, ["2:1"] = 1 } })
		assert.are.equal(require("Locale").followDone, line)
	end)

	it("says so when nothing is loaded", function()
		assert.are.equal(require("Locale").followNone, Follow.line(DATA, nil, {}))
	end)

	it("passes a codec refusal through verbatim", function()
		local build, message = Follow.load("FS9:x", DATA)
		assert.is_nil(build)
		assert.is_truthy(message:find("FS9", 1, true))
	end)

	it("takes an FS1 export as a build too", function()
		local build = assert(Follow.load("FS1:1.60.1.69893:paladin:human:203/0/0:", DATA))
		assert.are.equal("FS1", build.format)
		assert.are.same({ tab = 1, tier = 1, column = 1, index = 1 }, Follow.nextPoint(build, { [1] = {} }))
	end)

	it("names an unknown cell from the locale when Data.lua has no talent for it", function()
		-- tab 1, tier 9, column 1: a cell no talent in DATA's Holy tab has.
		local build = assert(Follow.load("FSB1:1.60.1.69893:paladin:191:", DATA))
		local Locale = require("Locale")
		local line = Follow.line(DATA, build, { [1] = {} })
		local unknownName = string.format(Locale.followUnknownCell, 9, 1)
		assert.are.equal(string.format(Locale.followNext, unknownName, "Holy", 9), line)
	end)
end)

-- "the addon isn't character aware" (owner, 2026-09-28): ForeverSixtyDB.follow
-- held one build for the whole account; these cover the per-character
-- storage that replaced it, the one-time migration off the old key, and the
-- inbox filter that keeps a companion build from crossing characters.
describe("Follow, per character", function()
	local Follow

	before_each(function()
		Follow = helper.load("Follow")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("keys the saved build by the current character, the same key Export.save writes", function()
		mock.install({ realm = "Ashbringer", region = 1, playerName = "Alice" })
		Follow.load(CODE, DATA, "Deep Holy")
		assert.are.equal("US/Ashbringer/Alice", require("Export").characterKey())
		assert.are.same({ code = CODE, name = "Deep Holy" }, _G.ForeverSixtyDB.follows["US/Ashbringer/Alice"])
	end)

	it("does not offer one character's build to another", function()
		mock.install({ realm = "Ashbringer", region = 1, playerName = "Alice" })
		Follow.load(CODE, DATA, "Alice's build")

		-- A second character logging in is a fresh addon load, not a change
		-- to the same session: reload Follow so Follow.build starts nil
		-- again, the way it would after a real relog.
		_G.UnitName = function() return "Bob" end
		Follow = helper.load("Follow")
		assert.is_nil(Follow.restore(DATA))
		assert.is_nil(Follow.build)
	end)

	it("migrates the old account-wide build to the first character that restores, then deletes it", function()
		mock.install({ realm = "Ashbringer", region = 1, playerName = "Alice" })
		_G.ForeverSixtyDB = { follow = { code = CODE, name = "Old Account Build" } }
		local build = Follow.restore(DATA)
		assert.are.equal("Old Account Build", build.name)
		assert.is_nil(_G.ForeverSixtyDB.follow)
		assert.are.same({ code = CODE, name = "Old Account Build" },
			_G.ForeverSixtyDB.follows["US/Ashbringer/Alice"])
	end)

	it("does not hand the migrated build to a second character too", function()
		mock.install({ realm = "Ashbringer", region = 1, playerName = "Alice" })
		_G.ForeverSixtyDB = { follow = { code = CODE, name = "Old Account Build" } }
		Follow.restore(DATA)

		_G.UnitName = function() return "Bob" end
		Follow = helper.load("Follow")
		assert.is_nil(Follow.restore(DATA))
	end)

	it("forgets only the current character's build", function()
		mock.install({ realm = "Ashbringer", region = 1, playerName = "Alice" })
		Follow.load(CODE, DATA)
		_G.ForeverSixtyDB.follows["US/Ashbringer/Bob"] = { code = CODE, name = "Bob's" }
		Follow.forget()
		assert.is_nil(Follow.build)
		assert.is_nil(_G.ForeverSixtyDB.follows["US/Ashbringer/Alice"])
		assert.is_truthy(_G.ForeverSixtyDB.follows["US/Ashbringer/Bob"])
	end)
end)

describe("Follow.inbox", function()
	local Follow

	before_each(function()
		Follow = helper.load("Follow")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("keeps a build addressed to the current character", function()
		local usable = Follow.inbox({ builds = {
			{ id = "a", name = "Mine", character = "us/pvp/bow-jackzon", code = CODE },
		} }, "US/PvP/Bow Jackzon")
		assert.are.equal(1, #usable)
	end)

	it("keeps a build with no character at all -- site-wide, not addressed", function()
		local usable = Follow.inbox({ builds = {
			{ id = "a", name = "Anyone's", code = CODE },
		} }, "US/PvP/Bow Jackzon")
		assert.are.equal(1, #usable)
	end)

	it("drops a build addressed to a different character", function()
		local usable = Follow.inbox({ builds = {
			{ id = "a", name = "Not mine", character = "us/pvp/someone-else", code = CODE },
		} }, "US/PvP/Bow Jackzon")
		assert.are.equal(0, #usable)
	end)

	it("keeps the companion's order across a mix of addressed, site-wide and foreign builds", function()
		local usable = Follow.inbox({ builds = {
			{ id = "a", name = "Mine", character = "us/pvp/bow-jackzon", code = CODE },
			{ id = "b", name = "Foreign", character = "us/pvp/someone-else", code = CODE },
			{ id = "c", name = "Site-wide", code = CODE },
		} }, "US/PvP/Bow Jackzon")
		assert.are.same({ "a", "c" }, { usable[1].id, usable[2].id })
	end)
end)

describe("Follow.sameCharacter", function()
	local Follow

	before_each(function()
		Follow = helper.load("Follow")
	end)

	-- The rule (Do item 3): lower-case, spaces folded to hyphens, each of
	-- the three "/"-separated segments compared that way -- so the site's
	-- slugified key and the addon's own display-name key agree.
	it("matches the addon's display-name key against the site's slugified one", function()
		assert.is_true(Follow.sameCharacter("US/Ashbringer/Bow Jackzon", "us/ashbringer/bow-jackzon"))
	end)

	it("is case-insensitive on every segment", function()
		assert.is_true(Follow.sameCharacter("us/ashbringer/bow jackzon", "US/ASHBRINGER/BOW-JACKZON"))
	end)

	it("does not match a different name", function()
		assert.is_false(Follow.sameCharacter("US/Ashbringer/Bow Jackzon", "us/ashbringer/someone-else"))
	end)

	it("does not match a different realm or ruleset segment", function()
		assert.is_false(Follow.sameCharacter("US/Ashbringer/Bow Jackzon", "us/pvp/bow-jackzon"))
	end)

	it("refuses a value with no character-key shape", function()
		assert.is_false(Follow.sameCharacter("not-a-key", "US/Ashbringer/Bow Jackzon"))
		assert.is_false(Follow.sameCharacter(nil, "US/Ashbringer/Bow Jackzon"))
	end)
end)
