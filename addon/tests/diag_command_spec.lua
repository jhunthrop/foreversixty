-- addon/tests/diag_command_spec.lua
-- /fs diag on a fresh session (L3.1 fix round 5): the command asks Score
-- itself, so a missing or old data addon is reported without any tooltip
-- having read it first. Pure: Options.handle("diag") with ForeverSixtyData
-- set or not, no client mock (no frame is built; the chip has not drawn).
-- The build line (L3.8) follows the verdict: nothing is worked out before login. The
-- character-panel tab's line (L4.3 ruling 7) follows the surfaces': closed, nothing counted.
-- The guild tab's line (L5.1 ruling 6) after it: closed, no rows, not an officer.
local helper = require("spec_helper")

describe("/fs diag on a fresh session", function()
	local saved, L, Options, NO_BUILD, PANEL, GUILD

	local function fresh(data)
		_G.ForeverSixtyData = data
		L = helper.load("Locale")
		require("Theme").reset()
		helper.load("Score")
		helper.load("Follow")
		Options = helper.load("Options")
		NO_BUILD = string.format(L.diagBuildNone, L.diagBuildNotYet)
		PANEL = string.format(L.diagPanel, L.diagPanelClosed, 0, 0)
		GUILD = string.format(L.diagGuildTab, L.diagPanelClosed, 0, 0, L.diagGuildTabNo)
	end

	before_each(function()
		saved = _G.ForeverSixtyData
	end)

	after_each(function()
		_G.ForeverSixtyData = saved
	end)

	it("says the data addon is not installed, as a finding under Everything is working", function()
		fresh(nil)
		assert.are.same({ L.diagAllWorking, NO_BUILD, PANEL, GUILD, L.diagDataAddonMissingInfo }, Options.handle("diag"))
	end)

	it("says the data addon is too old, under Not working", function()
		fresh({ format = 1, characters = {}, guilds = {} })
		-- L4.1 fix round 5: the status lines under their own "Working:" heading.
		assert.are.same({ L.diagHeading, L.diagDataAddonOld, L.diagWorkingHeading, NO_BUILD, PANEL, GUILD },
			Options.handle("diag"))
	end)

	it("says it once however often /fs diag is asked", function()
		fresh(nil)
		Options.handle("diag")
		assert.are.same({ L.diagAllWorking, NO_BUILD, PANEL, GUILD, L.diagDataAddonMissingInfo }, Options.handle("diag"))
	end)

	-- L3.2 fix round 2: geometry findings only under /fs diag verbose.
	it("keeps the verbose findings for /fs diag verbose", function()
		fresh(nil)
		require("Theme").noteVerbose("Talent window found: PlayerSpellsFrame (trait mapping)")
		assert.are.same({ L.diagAllWorking, NO_BUILD, PANEL, GUILD, L.diagDataAddonMissingInfo }, Options.handle("diag"))
		assert.are.same({ L.diagAllWorking, NO_BUILD, PANEL, GUILD, L.diagDataAddonMissingInfo,
			"Talent window found: PlayerSpellsFrame (trait mapping)" }, Options.handle("diag verbose"))
	end)
end)
