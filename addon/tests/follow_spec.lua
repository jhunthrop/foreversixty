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

	it("refresh builds one frame and keeps using it", function()
		assert(Follow.load(CODE, DATA))
		local state = mock.install({
			talents = {
				{ name = "Holy", talents = {} },
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			},
		})

		local text = Follow.refresh(DATA)
		assert.are.equal(1, #state.frames)
		assert.is_truthy(text:find("Improved Holy Strike", 1, true))
		assert.is_true(state.frames[1].shown)

		-- A second call must not build a second frame.
		local text2 = Follow.refresh(DATA)
		assert.are.equal(1, #state.frames)
		assert.are.equal(text, text2)
	end)

	it("highlight acts on the open talent frame", function()
		assert(Follow.load(CODE, DATA))
		mock.install({})
		_G.PlayerTalentFrame = { kind = "stub" }

		local result = Follow.highlight(DATA, { [1] = {} })

		_G.PlayerTalentFrame = nil
		assert.are.same({ frame = "PlayerTalentFrame", point = { tab = 1, tier = 1, column = 1, index = 1 } }, result)
	end)

	it("highlight does nothing, and does not error, when no talent frame is open", function()
		assert(Follow.load(CODE, DATA))
		mock.install({})
		_G.PlayerTalentFrame = nil
		_G.TalentFrame = nil

		assert.has_no.errors(function()
			assert.is_nil(Follow.highlight(DATA, { [1] = {} }))
		end)
	end)
end)
