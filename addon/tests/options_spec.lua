local helper = require("spec_helper")
local mock = require("wow_mock")

local DATA = {
	build = "1.60.1.69893",
	classes = { paladin = { tabs = {
		{ name = "Holy", talents = { { name = "A", tier = 1, column = 1, maxRank = 5 } } },
		{ name = "Protection", talents = {} },
		{ name = "Retribution", talents = {} },
	} } },
	weights = { ["paladin-holy"] = { spell_power = 1.0 } },
}

describe("Options", function()
	local Options
	local state

	before_each(function()
		state = mock.install({
			class = { name = "Paladin", token = "PALADIN" },
			race = { name = "Human", token = "Human" },
			realm = "Ashbringer",
			region = 1,
			talents = {
				{ name = "Holy", talents = { { name = "A", tier = 1, column = 1, rank = 5, maxRank = 5 } } },
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			},
		})
		Options = helper.load("Options")
		Options.data = DATA
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("shows the data build id with no argument", function()
		local lines = Options.handle("")
		assert.are.equal(string.format(require("Locale").dataBuild, "1.60.1.69893"), lines[1])
		assert.are.equal(require("Locale").slashHint, lines[2])
	end)

	it("warns when the client build is not the data build", function()
		_G.GetBuildInfo = function()
			return "1.61.0", "70000", "Oct 2026", 16100
		end
		local lines = Options.handle("")
		assert.is_truthy(table.concat(lines, "\n"):find("1.61.0", 1, true))
	end)

	it("exports on /fs export", function()
		local lines = Options.handle("export")
		assert.is_truthy(lines[1]:find("FS1:", 1, true))
	end)

	it("loads a build on /fs follow <code> and says how many points", function()
		local lines = Options.handle("follow FSB1:1.60.1.69893:paladin:111:")
		assert.are.equal(string.format(require("Locale").followLoaded, "paladin", 1), lines[1])
	end)

	it("passes a refusal through verbatim on a bad code", function()
		local lines = Options.handle("follow FS9:nope")
		assert.is_truthy(lines[1]:find("FS9", 1, true))
	end)

	it("shows the next point on /fs follow with no code", function()
		Options.handle("follow FSB1:1.60.1.69893:paladin:111:")
		-- Local override: the shared rank of 5 is maxed, which the /fs
		-- export example needs, but it leaves this build's one-point order
		-- already satisfied. Zero it here, for this example only, so there
		-- is an outstanding point to show.
		state.talents[1].talents[1].rank = 0
		local lines = Options.handle("follow")
		assert.is_truthy(lines[1]:find("A", 1, true))
	end)

	it("routes /fs options to the same lines as bare /fs", function()
		assert.are.same(Options.handle(""), Options.handle("options"))
	end)

	it("shows gear lines on /fs gear", function()
		Options.handle("follow FSB1:1.60.1.69893:paladin:111:")
		assert.are.same({ require("Locale").gearNone }, Options.handle("gear"))
	end)

	it("counts the builds the companion left in the inbox", function()
		_G.ForeverSixtyInbox = {
			generated_at = "2026-09-20T00:00:00Z",
			builds = {
				{ id = "a", name = "Deep Holy", character = "US/Ashbringer/Bob", code = "FSB1:1.60.1.69893:paladin:111:" },
			},
		}
		assert.are.same({ string.format(require("Locale").inboxCount, 1) }, Options.handle("inbox"))
	end)

	it("says the inbox is empty rather than nothing at all", function()
		_G.ForeverSixtyInbox = nil
		assert.are.same({ require("Locale").inboxEmpty }, Options.handle("inbox"))
	end)

	it("skips an inbox build with no code rather than counting it", function()
		_G.ForeverSixtyInbox = {
			generated_at = "2026-09-20T00:00:00Z",
			builds = {
				{ id = "a", name = "Deep Holy", character = "US/Ashbringer/Bob", code = "FSB1:1.60.1.69893:paladin:111:" },
				{ id = "b", name = "No Code Yet" },
			},
		}
		assert.are.same({ string.format(require("Locale").inboxCount, 1) }, Options.handle("inbox"))
	end)

	it("loads the first inbox build at login and never writes the inbox", function()
		_G.ForeverSixtyInbox = {
			generated_at = "2026-09-20T00:00:00Z",
			builds = { { id = "a", name = "Deep Holy", code = "FSB1:1.60.1.69893:paladin:111:" } },
		}
		local before = _G.ForeverSixtyInbox
		Options.readInbox()
		assert.are.equal(before, _G.ForeverSixtyInbox)
		assert.is_truthy(require("Follow").build)
	end)
end)
