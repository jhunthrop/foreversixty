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
		assert.are.same({ active = "raid", slots = { raid = { code = CODE, name = "Deep Holy" } } },
			_G.ForeverSixtyDB.follows["US/Ashbringer/Alice"])
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
		assert.are.same(
			{ active = "raid", slots = { raid = { code = CODE, name = "Old Account Build" } } },
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

-- Named build slots and the "build arrived" banner (Wave A2,
-- docs/superpowers/specs/2026-09-28-addon-character-aware-design.md
-- section 1/2): a character keeps one build PER SLOT ("Raid", "Leveling",
-- "PvP"), switches with one click, and a companion build queued for this
-- character shows once on the Overview until loaded or dismissed.
describe("Follow, named build slots", function()
	local Follow

	before_each(function()
		mock.install({ realm = "Ashbringer", region = 1, playerName = "Alice" })
		Follow = helper.load("Follow")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("lists the three slots in order, none holding a build yet", function()
		local slots = Follow.slotsFor()
		assert.are.same({ "raid", "leveling", "pvp" }, { slots[1].id, slots[2].id, slots[3].id })
		assert.is_true(slots[1].active)
		for _, slot in ipairs(slots) do
			assert.is_false(slot.hasBuild)
		end
	end)

	it("loading a build without naming a slot fills the active one", function()
		Follow.load(CODE, DATA, "Deep Holy")
		local slots = Follow.slotsFor()
		assert.is_true(slots[1].hasBuild)
		assert.is_false(slots[2].hasBuild)
	end)

	it("loading into a named slot does not touch another slot's build", function()
		Follow.load(CODE, DATA, "Raid build", "raid")
		Follow.load(CODE, DATA, "Leveling build", "leveling")
		local key = require("Export").characterKey()
		local entry = _G.ForeverSixtyDB.follows[key]
		assert.are.equal("Raid build", entry.slots.raid.name)
		assert.are.equal("Leveling build", entry.slots.leveling.name)
	end)

	it("switching the active slot loads what that slot holds", function()
		Follow.load(CODE, DATA, "Raid build", "raid")
		Follow.load(CODE, DATA, "Leveling build", "leveling")
		local build = Follow.setActiveSlot(DATA, "leveling")
		assert.are.equal("Leveling build", build.name)
		assert.are.equal("Leveling build", Follow.build.name)
		local slots = Follow.slotsFor()
		assert.is_false(slots[1].active)
		assert.is_true(slots[2].active)
	end)

	it("switching to an empty slot clears the loaded build rather than refusing", function()
		Follow.load(CODE, DATA, "Raid build", "raid")
		local build = Follow.setActiveSlot(DATA, "pvp")
		assert.is_nil(build)
		assert.is_nil(Follow.build)
	end)

	it("forgetting one slot leaves the others alone", function()
		Follow.load(CODE, DATA, "Raid build", "raid")
		Follow.load(CODE, DATA, "Leveling build", "leveling")
		Follow.setActiveSlot(DATA, "leveling")
		Follow.forget("raid")
		local slots = Follow.slotsFor()
		assert.is_false(slots[1].hasBuild)
		assert.is_true(slots[2].hasBuild)
		-- Forgetting a slot that is not active must not clear the build
		-- that IS loaded.
		assert.is_not_nil(Follow.build)
	end)

	it("forgetting the last slot drops the character's whole entry", function()
		Follow.load(CODE, DATA, "Raid build", "raid")
		Follow.forget("raid")
		local key = require("Export").characterKey()
		assert.is_nil(_G.ForeverSixtyDB.follows[key])
	end)

	it("migrates a pre-slot single build into the default slot on restore", function()
		local key = require("Export").characterKey()
		_G.ForeverSixtyDB = { follows = { [key] = { code = CODE, name = "Old Single Build" } } }
		local build = Follow.restore(DATA)
		assert.are.equal("Old Single Build", build.name)
		local entry = _G.ForeverSixtyDB.follows[key]
		assert.are.equal("raid", entry.active)
		assert.are.equal("Old Single Build", entry.slots.raid.name)
	end)
end)

describe("Follow, the build-arrived banner", function()
	local Follow

	before_each(function()
		mock.install({ realm = "Ashbringer", region = 1, playerName = "Alice" })
		Follow = helper.load("Follow")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("finds the first undismissed inbox build addressed to this character", function()
		local inbox = { builds = { { id = "a", name = "Deep Holy", code = CODE } } }
		local pending = Follow.pendingArrival(inbox)
		assert.are.equal("a", pending.id)
	end)

	it("stops offering a build once it is dismissed", function()
		local inbox = { builds = { { id = "a", name = "Deep Holy", code = CODE } } }
		assert.is_not_nil(Follow.pendingArrival(inbox))
		Follow.dismissInbox("a")
		assert.is_nil(Follow.pendingArrival(inbox))
	end)

	it("moves on to the next build once the first is dismissed", function()
		local inbox = { builds = {
			{ id = "a", name = "First", code = CODE },
			{ id = "b", name = "Second", code = CODE },
		} }
		Follow.dismissInbox("a")
		assert.are.equal("b", Follow.pendingArrival(inbox).id)
	end)

	it("counts the talent cells two builds disagree on", function()
		local a = { order = {
			{ tab = 1, tier = 1, column = 1 }, { tab = 1, tier = 1, column = 1 },
			{ tab = 1, tier = 2, column = 1 },
		} }
		local b = { order = {
			{ tab = 1, tier = 1, column = 1 }, { tab = 1, tier = 1, column = 1 },
			{ tab = 1, tier = 1, column = 2 },
		} }
		-- 1:2:1 (a wants one, b wants none) and 1:1:2 (a wants none, b wants
		-- one) each count once; 1:1:1 agrees at two points either way.
		assert.are.equal(2, Follow.orderDiffCount(a, b))
	end)

	it("reports zero when both builds want exactly the same points", function()
		local build = assert(Follow.load(CODE, DATA))
		assert.are.equal(0, Follow.orderDiffCount(build, build))
	end)

	it("says so when this is the first build for the character", function()
		assert.are.equal(require("Locale").buildArrivedFirst, Follow.arrivalSummary(DATA, nil, CODE))
	end)

	it("says the builds match when the diff is zero", function()
		local build = assert(Follow.load(CODE, DATA))
		assert.are.equal(require("Locale").buildArrivedSame, Follow.arrivalSummary(DATA, build, CODE))
	end)

	it("counts the differing points against the currently loaded build", function()
		local current = { classSlug = "paladin", order = {
			{ tab = 1, tier = 1, column = 1 }, { tab = 1, tier = 1, column = 1 },
		} }
		-- CODE spends its second point at 1:2:1, which `current` does not.
		local summary = Follow.arrivalSummary(DATA, current, CODE)
		local Locale = require("Locale")
		assert.are.equal(string.format(Locale.buildArrivedDiff, 1), summary)
	end)
end)
