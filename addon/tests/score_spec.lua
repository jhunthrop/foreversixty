-- addon/tests/score_spec.lua
-- Score.lua reads the nightly record (contract 5, docs/contracts/addon-v2.md) and never
-- computes a number. Pure: no wow_mock. ForeverSixtyData is set to the L0 golden fixture,
-- loaded with plain Lua, and read back through the real reader.
local helper = require("spec_helper")

local ROOTS = { "../docs/contracts/fixtures/", "docs/contracts/fixtures/" }

local function loadFixture()
	for _, root in ipairs(ROOTS) do
		local file = io.open(root .. "data-addon-v2.lua", "rb")
		if file then
			local text = file:read("a")
			file:close()
			local env = {}
			assert(load(text, "@" .. root .. "data-addon-v2.lua", "t", env))()
			return env.ForeverSixtyData
		end
	end
	error("fixture not found: data-addon-v2.lua")
end

local DOT = "\194\183" -- the middle dot in L.exportTreeSeparator

describe("Score", function()
	local Score, Compat, Theme, L

	local function start(data)
		_G.ForeverSixtyData = data
		L = helper.load("Locale")
		Compat = helper.load("Compat")
		Theme = helper.load("Theme")
		Score = helper.load("Score")
	end

	local function fixture(change)
		local data = loadFixture()
		if change then
			change(data)
		end
		return data
	end

	after_each(function()
		_G.ForeverSixtyData = nil
		_G.ITEM_QUALITY_COLORS = nil
	end)

	describe("Compat.shortDate", function()
		it("words a YYYY-MM-DD date as day and English month", function()
			start(nil)
			assert.are.equal("4 Nov", Compat.shortDate("2026-11-04"))
			assert.are.equal("31 Dec", Compat.shortDate("2026-12-31"))
			assert.are.equal("1 Jan", Compat.shortDate("2027-01-01"))
		end)

		it("answers nil for anything that is not a date", function()
			start(nil)
			assert.is_nil(Compat.shortDate(nil))
			assert.is_nil(Compat.shortDate(""))
			assert.is_nil(Compat.shortDate("yesterday"))
			assert.is_nil(Compat.shortDate("2026-13-04"))
		end)
	end)

	describe("the date-only stamp", function()
		it("reads a date as the start of that day, UTC", function()
			start(nil)
			assert.are.equal(1793750400, Score.dateSeconds("2026-11-04"))
			assert.are.equal(0, Score.dateSeconds("1970-01-01"))
			assert.are.equal(86400, Score.dateSeconds("1970-01-02"))
			assert.is_nil(Score.dateSeconds("2026-11-04T00:00:00Z"))
			assert.is_nil(Score.dateSeconds(nil))
		end)
	end)

	describe("status", function()
		it("says the data addon is not installed when the global is absent", function()
			start(nil)
			local status = Score.status()
			assert.is_false(status.available)
			assert.are.equal(L.ratingsNotInstalled, status.reason)
		end)

		it("refuses a format newer than 2, naming the fix", function()
			start(fixture(function(data) data.format = 3 end))
			local status = Score.status()
			assert.is_false(status.available)
			assert.are.equal(L.ratingsUpdateAddon, status.reason)
		end)

		it("refuses a format older than 2 as unreadable rather than guessing at its keys", function()
			start(fixture(function(data) data.format = 1 end))
			local status = Score.status()
			assert.is_false(status.available)
			assert.are.equal(L.ratingsUnreadable, status.reason)
		end)

		it("treats a malformed table as unreadable rather than erroring", function()
			start({ format = "two", characters = "nope" })
			assert.are.equal(L.ratingsUnreadable, Score.status().reason)
			start({})
			assert.are.equal(L.ratingsUnreadable, Score.status().reason)
			start("a string")
			assert.are.equal(L.ratingsUnreadable, Score.status().reason)
		end)

		it("is available and fresh for a file generated today", function()
			start(fixture(function(data) data.generated = os.date("!%Y-%m-%d") end))
			local status = Score.status()
			assert.is_true(status.available)
			assert.is_false(status.stale)
			assert.are.equal(os.date("!%Y-%m-%d"), status.generated)
		end)

		it("is not stale at two days and stale at five, reading the date as UTC midnight", function()
			start(fixture(function(data)
				data.generated = os.date("!%Y-%m-%d", os.time() - 2 * 86400)
			end))
			assert.is_false(Score.status().stale)
			start(fixture(function(data)
				data.generated = os.date("!%Y-%m-%d", os.time() - 5 * 86400)
			end))
			local status = Score.status()
			assert.is_true(status.available)
			assert.is_true(status.stale)
		end)

		it("marks a file with no readable date stale but still usable", function()
			start(fixture(function(data) data.generated = "2026-09-22T04:00:00Z" end))
			local status = Score.status()
			assert.is_true(status.available)
			assert.is_true(status.stale)
		end)

		-- Fix round 1: /fs diag says nothing when the data addon is simply not
		-- installed, and one plain line when it is there but cannot be read.
		-- Fix round 3: the two causes, each with its own fix.
		it("says in /fs diag that the data addon is not installed", function()
			start(nil)
			local status = Score.status()
			assert.is_true(status.absent)
			assert.are.same({ L.diagDataAddonMissing }, Theme.diagnostics())
			assert.are.equal("The Forever Sixty Data addon is not installed. Install it from the site.",
				L.diagDataAddonMissing)
		end)

		it("says in /fs diag that the data addon is too old when it cannot be read", function()
			start(fixture(function(data) data.format = 1 end))
			Score.status()
			assert.are.same({ L.diagDataAddonOld }, Theme.diagnostics())
			assert.are.equal("The Forever Sixty Data addon is too old. Update it from the site.", L.diagDataAddonOld)
		end)

		it("notes why the file is unusable once, and nothing for a usable file", function()
			start(fixture(function(data) data.format = 3 end))
			Score.status()
			Score.status()
			local notes = Theme.diagnostics()
			assert.are.equal(1, #notes)
			assert.are.equal(L.ratingsUpdateAddon, notes[1])
			start(fixture())
			Score.status()
			assert.are.equal(0, #Theme.diagnostics())
		end)
	end)

	describe("forCharacter", function()
		it("reads the fixture's record with its date and tier", function()
			start(fixture())
			local record = Score.forCharacter("Obnoxious", "Yell")
			assert.are.equal(1830, record.score)
			assert.are.equal("epic", record.tier)
			assert.are.equal("as of 4 Nov", record.asOfText)
			assert.are.equal(12, record.fights)
			assert.are.equal(412, record.components.gear)
			assert.are.equal(71, record.best_parses["Lucifron"])
			assert.are.equal("/character/us/forever/obnoxious-yell", record.profile)
		end)

		it("takes a Name-Realm unit name without a separate realm", function()
			start(fixture())
			assert.are.equal(1830, Score.forCharacter("Obnoxious-Yell").score)
		end)

		it("finds a realm with spaces in its normalized spelling", function()
			start(fixture(function(data)
				data.characters["Bob-ClassicBetaPvP"] = { score = 700, tier = "uncommon", as_of = "2026-11-03", fights = 3 }
			end))
			assert.are.equal(700, Score.forCharacter("Bob", "Classic Beta PvP").score)
		end)

		it("answers nil for a name the record does not have", function()
			start(fixture())
			assert.is_nil(Score.forCharacter("Nobody", "Yell"))
			assert.is_nil(Score.forCharacter("Obnoxious", "Elsewhere"))
			assert.is_nil(Score.forCharacter(nil, "Yell"))
		end)

		it("answers nil when the data addon is absent or too new", function()
			start(nil)
			assert.is_nil(Score.forCharacter("Obnoxious", "Yell"))
			start(fixture(function(data) data.format = 3 end))
			assert.is_nil(Score.forCharacter("Obnoxious", "Yell"))
		end)

		it("never returns a zero for a character with no logs yet", function()
			start(fixture(function(data)
				data.characters["Fresh-Yell"] = { score = 0, tier = "common", as_of = "2026-11-04", fights = 0 }
				data.characters["Blank-Yell"] = { tier = "common", fights = 0 }
			end))
			assert.is_nil(Score.forCharacter("Fresh", "Yell"))
			assert.is_nil(Score.forCharacter("Blank", "Yell"))
		end)

		it("does not change the record it read", function()
			local data = fixture()
			start(data)
			Score.forCharacter("Obnoxious", "Yell")
			assert.is_nil(data.characters["Obnoxious-Yell"].asOfText)
			assert.is_nil(data.characters["Obnoxious-Yell"].tierColor)
		end)
	end)

	describe("tierColor", function()
		it("is the client's quality colour for the tier name", function()
			_G.ITEM_QUALITY_COLORS = {
				[1] = { r = 1, g = 1, b = 1 }, [2] = { r = 0.1, g = 1, b = 0 }, [3] = { r = 0, g = 0.44, b = 0.87 },
				[4] = { r = 0.64, g = 0.21, b = 0.93 }, [5] = { r = 1, g = 0.5, b = 0 },
			}
			start(fixture())
			local r, g, b = Score.tierColor("epic")
			assert.are.same({ 0.64, 0.21, 0.93 }, { r, g, b })
			assert.are.equal(1, (Score.tierColor("common")))
			assert.are.equal(0.1, (Score.tierColor("uncommon")))
			assert.are.equal(0, (Score.tierColor("rare")))
			assert.are.equal(0.5, select(2, Score.tierColor("legendary")))
			local record = Score.forCharacter("Obnoxious", "Yell")
			assert.are.same({ r = 0.64, g = 0.21, b = 0.93, a = 1 }, record.tierColor)
		end)

		it("is nil for a missing or unknown tier, which shows the score uncoloured", function()
			start(fixture(function(data) data.characters["Obnoxious-Yell"].tier = "mythic" end))
			assert.is_nil(Score.tierColor("mythic"))
			assert.is_nil(Score.tierColor(nil))
			assert.is_nil(Score.forCharacter("Obnoxious", "Yell").tierColor)
		end)
	end)

	describe("lines", function()
		it("heads the tooltip with score, tier name and date, then the components", function()
			start(fixture())
			local lines = Score.lines(Score.forCharacter("Obnoxious", "Yell"))
			assert.are.equal("Forever Score 1830 " .. DOT .. " Epic " .. DOT .. " as of 4 Nov", lines[1])
			assert.are.equal(6, #lines)
			assert.are.equal("Gear 412", lines[2])
			assert.are.equal("Discipline 78", lines[6])
		end)

		it("drops the tier when the record has none, and is empty for no record", function()
			start(fixture(function(data) data.characters["Obnoxious-Yell"].tier = nil end))
			local lines = Score.lines(Score.forCharacter("Obnoxious", "Yell"))
			assert.are.equal("Forever Score 1830 " .. DOT .. " as of 4 Nov", lines[1])
			assert.are.same({}, Score.lines(nil))
		end)
	end)

	describe("guild", function()
		it("finds the guild record by name", function()
			start(fixture())
			local guild = Score.guild("OLYMPUS XXVII")
			assert.are.equal(3, guild.standing)
			assert.are.equal(4, guild.nights)
			assert.are.equal("Obnoxious-Yell", guild.members[1])
			assert.is_nil(Score.guild("Nobody"))
			assert.is_nil(Score.guild(nil))
		end)

		it("answers nil without usable data", function()
			start(nil)
			assert.is_nil(Score.guild("OLYMPUS XXVII"))
		end)
	end)

	it("understands format 2", function()
		start(nil)
		assert.are.equal(2, Score.FORMAT)
	end)
end)
