-- addon/tests/addon_tour_spec.lua
-- The addon-tour lane (design/specs/2026-11-08-addon-ux.md section 6, owner 2026-11-11): every tour
-- page opens on a lead; the last page asks the spec question when the points name no tree and lays
-- its tiles out for three or four specs; the tour steps aside while the Builds dialog is open and
-- comes back on the same page; the Builds dialog leads with the addon's own builds for the class,
-- the followed one marked, one tap follows another; Follow.chooseDefault answers the question over
-- a saved build too.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local DATA = {
	build = "1.60.1.70291",
	classes = { warrior = { tabs = {
		{ name = "Arms", talents = {
			{ name = "Improved Heroic Strike", tier = 1, column = 1, maxRank = 3 },
			{ name = "Deflection", tier = 1, column = 2, maxRank = 5 },
		} },
		{ name = "Fury", talents = { { name = "Cruelty", tier = 1, column = 3, maxRank = 5 } } },
		{ name = "Protection", talents = { { name = "Shield Specialization", tier = 1, column = 2, maxRank = 5 } } },
	} } },
	weights = { ["warrior-arms"] = { strength = 1 }, ["warrior-fury"] = { strength = 1 } },
	defaultBuilds = {
		["warrior-arms"] = {
			leveling = { name = "Arms leveling", order = { { 1, 1, 1 }, { 1, 1, 1 }, { 1, 1, 2 } } },
			raid = { name = "Arms raid", order = { { 1, 1, 2 }, { 2, 1, 3 } } },
		},
		["warrior-fury"] = {
			leveling = { name = "Fury leveling", order = { { 2, 1, 3 }, { 2, 1, 3 } } },
			raid = { name = "Fury raid", order = { { 2, 1, 3 } } },
		},
	},
}

local UNDECIDED = { { name = "Arms", talents = { { name = "Improved Heroic Strike", tier = 1, column = 1, rank = 1,
	maxRank = 3 } } }, { name = "Fury", talents = { { name = "Cruelty", tier = 1, column = 3, rank = 1, maxRank = 5 } } },
	{ name = "Protection", talents = {} } }

describe("the tour's pages (spec 6.3 and 6.2), the pure model", function()
	local Onboarding

	before_each(function()
		mock.install({})
		helper.load("Theme").reset()
		Onboarding = helper.load("Onboarding")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("opens every page on a lead that says what the player gets", function()
		for _, key in ipairs({ "loop", "site", "back", "guild", "setup" }) do
			local lead = Onboarding.lead(key)
			assert.is_string(lead, key)
			assert.is_truthy(#lead > 40 and #lead <= 160, key .. ": " .. #lead)
			assert.is_nil(lead:find("\226\128\147", 1, true), key)
			assert.is_nil(lead:find("\226\128\148", 1, true), key)
		end
		assert.are.equal(L.tourLoopLead, Onboarding.lead("loop"))
		-- Page 4's intro is gone: the lead carries its claim.
		assert.is_nil(L.tourGuildIntro)
		assert.are.equal("The line by your minimap.", L.tourGuildChipNote)
	end)

	it("asks the spec question over the tiles when the points name no tree, else Pick your path", function()
		local asked = Onboarding.setup({ level = 13, question = "Are you Holy or Retribution?",
			specs = { { slug = "paladin-holy", name = "Holy" } } })
		assert.are.equal("Are you Holy or Retribution?", asked.ask)
		assert.are.equal(1, #asked.tiles)
		local plain = Onboarding.setup({ level = 13, specs = { { slug = "paladin-holy", name = "Holy" } } })
		assert.are.equal(L.panelPickPath, plain.ask)
		assert.is_nil(Onboarding.setup({ level = 13, build = { order = {} }, buildName = "Holy build" }).tiles)
	end)

	it("lays three tiles out at 190 and the druid's four at 141, within the content", function()
		assert.are.equal(190, Onboarding.tileWidth(3))
		assert.are.equal(141, Onboarding.tileWidth(4))
		assert.is_true(Onboarding.tileWidth(4) * 4 + Onboarding.TILE_GAP * 3 <= Onboarding.WIDTH - 2 * Onboarding.PAD)
	end)
end)

describe("the tour steps aside for the Builds dialog (spec 6.1)", function()
	local Onboarding, BuildDialog, state

	before_each(function()
		state = mock.install({ class = { name = "Warrior", token = "WARRIOR" }, level = 13, playerName = "Bob",
			realm = "Ashbringer", region = 1, talents = UNDECIDED })
		helper.load("Theme").reset()
		_G.ForeverSixtyDB = {}
		_G.ForeverSixtyInbox = nil
		helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Cards")
		helper.load("Export")
		helper.load("Talents")
		helper.load("Gear")
		helper.load("Follow")
		helper.load("Chip")
		BuildDialog = helper.load("BuildDialog")
		BuildDialog.data = DATA
		Onboarding = helper.load("Onboarding")
		Onboarding.data = DATA
		require("Export").guildInfo = function()
			return nil
		end
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
		mock.uninstall()
	end)

	it("hides without stamping the account, and comes back on the same page when the dialog hides", function()
		assert.is_true(Onboarding.open(4))
		assert.are.equal("setup", Onboarding.draw())
		assert.is_true(BuildDialog.open("paste"))
		mock.runTimers(state)
		assert.is_false(Onboarding.frame:IsShown())
		assert.is_nil(_G.ForeverSixtyDB.welcomedAt)
		assert.is_true(BuildDialog.frame:IsShown())
		assert.is_not_nil(mock.lastCall(BuildDialog.frame, "Raise"))
		BuildDialog.close()
		assert.is_true(Onboarding.frame:IsShown())
		assert.are.equal(4, Onboarding.page)
		assert.is_nil(_G.ForeverSixtyDB.welcomedAt)
	end)

	it("does nothing when the tour is not showing", function()
		assert.is_false(Onboarding.stepAside())
		assert.is_false(Onboarding.stepBack())
		assert.is_true(BuildDialog.open("paste"))
		BuildDialog.close()
		assert.is_true(Onboarding.frame == nil or not Onboarding.frame:IsShown())
	end)

	it("draws the question, the tiles and the footer's again line on the last page for an undecided warrior",
		function()
		Onboarding.open(4)
		-- Every shown text on every frame (the mock keeps font strings in each frame's regions).
		local texts = {}
		for _, frame in ipairs(state.frames) do
			for _, region in ipairs(frame.regions or {}) do
				if type(region.GetText) == "function" and region:IsVisible() then
					texts[region:GetText() or ""] = true
				end
			end
			if type(frame.GetText) == "function" and frame:IsVisible() then
				texts[frame:GetText() or ""] = true
			end
		end
		assert.is_true(texts["Are you Arms or Fury?"] == true)
		assert.is_true(texts[L.tourAgain] == true)
		assert.is_true(texts[L.tourSetupLead] == true)
		assert.is_true(texts["Arms"] == true and texts["Fury"] == true)
		assert.is_true(Onboarding.frame:GetHeight() <= Onboarding.HEIGHT)
	end)
end)

describe("the Builds dialog's own builds (spec 6.4), the pure model", function()
	local BuildDialog

	before_each(function()
		helper.load("Export")
		helper.load("Follow")
		BuildDialog = helper.load("BuildDialog")
	end)

	it("lists one row per spec with a build, in tree order, named and split as the player reads them", function()
		local rows = BuildDialog.ownRows(DATA, "warrior", 13, nil)
		assert.are.equal(2, #rows)
		assert.are.equal("warrior-arms", rows[1].spec)
		assert.are.equal("Arms build", rows[1].name)
		assert.are.equal("Arms 3", rows[1].detail)
		assert.are.equal("leveling", rows[1].kind)
		assert.are.equal("Fury build", rows[2].name)
		assert.is_false(rows[1].following)
	end)

	it("marks the row whose points the followed build spends, and lists raid builds at 60", function()
		local followed = { classSlug = "warrior", order = { { tab = 2, tier = 1, column = 3 },
			{ tab = 2, tier = 1, column = 3 } } }
		local rows = BuildDialog.ownRows(DATA, "warrior", 13, followed)
		assert.is_false(rows[1].following)
		assert.is_true(rows[2].following)
		local raid = BuildDialog.ownRows(DATA, "warrior", 60, nil)
		assert.are.equal("raid", raid[1].kind)
		assert.are.equal("Arms 1 / Fury 1", raid[1].detail)
	end)

	it("heads the list with the class's builds, or with the question while the points name no tree", function()
		assert.are.equal("Warrior leveling builds", BuildDialog.ownHeading("Warrior", 13, nil))
		assert.are.equal("Warrior raid builds", BuildDialog.ownHeading("Warrior", 60, nil))
		assert.are.equal("Are you Arms or Fury?", BuildDialog.ownHeading("Warrior", 13, "Are you Arms or Fury?"))
	end)

	it("shows two site rows beside up to three own rows and one beside four", function()
		assert.are.equal(2, BuildDialog.siteRowLimit(3))
		assert.are.equal(1, BuildDialog.siteRowLimit(4))
	end)

	it("walks the cursor over the own rows then the site rows, starting on the first site row or nowhere", function()
		local items = BuildDialog.cursorItems({ { spec = "a" }, { spec = "b" } }, { { id = 1 } })
		assert.are.equal(3, #items)
		assert.is_true(items[1].own)
		assert.is_false(items[3].own)
		assert.are.equal(3, BuildDialog.startCursor(2, 1))
		assert.are.equal(0, BuildDialog.startCursor(2, 0))
		assert.are.equal(1, BuildDialog.startCursor(0, 2))
	end)
end)

describe("the Builds dialog's own builds (spec 6.4), the frame", function()
	local BuildDialog, Follow, Chip, state, said

	before_each(function()
		state = mock.install({ class = { name = "Warrior", token = "WARRIOR" }, level = 13, playerName = "Bob",
			realm = "Ashbringer", region = 1, talents = UNDECIDED })
		helper.load("Theme").reset()
		_G.ForeverSixtyDB = {}
		_G.ForeverSixtyInbox = nil
		helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Cards")
		helper.load("Export")
		helper.load("Talents")
		helper.load("Gear")
		Follow = helper.load("Follow")
		Chip = helper.load("Chip")
		Chip.data = DATA
		helper.load("ItemTooltip")
		BuildDialog = helper.load("BuildDialog")
		BuildDialog.data = DATA
		said = {}
		BuildDialog.say = function(line)
			said[#said + 1] = line
		end
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
		mock.uninstall()
	end)

	it("opens on the question, the own rows, Custom build and the companion line, the cursor on nothing", function()
		assert.is_true(BuildDialog.open("paste"))
		mock.runTimers(state)
		local paste = BuildDialog.paste
		assert.are.equal("Are you Arms or Fury?", paste.heading:GetText())
		assert.is_true(paste.heading:IsShown())
		assert.is_true(paste.custom:IsShown())
		assert.are.equal(L.dialogCustomHeading, paste.custom:GetText())
		assert.are.equal("Arms build", paste.own[1].name:GetText())
		assert.is_true(paste.own[1].frame:IsShown())
		assert.is_true(paste.own[2].frame:IsShown())
		assert.is_false(paste.own[3].frame:IsShown())
		assert.is_false(paste.own[1].mark:IsShown())
		assert.is_false(paste.own[1].edge:IsShown())
		assert.are.equal(0, BuildDialog.cursor)
		assert.is_false(BuildDialog.enterHint:IsShown())
		assert.is_true(paste.companion:IsShown())
	end)

	it("follows an own build on one tap, remembers the answer, says so and closes; the tour's path shows it",
		function()
		BuildDialog.open("paste")
		mock.runTimers(state)
		local build = BuildDialog.followOwn(BuildDialog.paste.own[2].entry)
		assert.is_not_nil(build)
		assert.are.equal("Fury build", Follow.buildName(DATA, Follow.build))
		assert.is_true(Follow.build.default)
		assert.are.same({ spec = "warrior-fury" }, Chip.chosenPath())
		assert.is_false(BuildDialog.frame:IsShown())
		assert.is_truthy(said[1]:find("Fury build", 1, true), said[1])
		-- Opened again: Fury is marked Following, the heading is still the question (the points are 1/1).
		BuildDialog.open("paste")
		mock.runTimers(state)
		assert.is_true(BuildDialog.paste.own[2].mark:IsShown())
		assert.is_false(BuildDialog.paste.own[1].mark:IsShown())
		assert.is_false(BuildDialog.paste.following:IsShown())
		-- Enter on an own row follows it; the hint says so.
		BuildDialog.setCursor(1)
		assert.are.equal(L.dialogEnterFollow, BuildDialog.enterHint:GetText())
		BuildDialog.previewCursor()
		assert.are.equal("Arms build", Follow.buildName(DATA, Follow.build))
	end)

	it("only closes on the row already followed", function()
		Follow.ensureDefault(DATA, { spec = "warrior-arms", classSlug = "warrior", ranks = {}, level = 13 })
		BuildDialog.open("paste")
		mock.runTimers(state)
		assert.is_true(BuildDialog.paste.own[1].mark:IsShown())
		BuildDialog.followOwn(BuildDialog.paste.own[1].entry)
		assert.is_false(BuildDialog.frame:IsShown())
		assert.are.same({}, said)
	end)
end)

describe("Follow.chooseDefault (spec 6.4)", function()
	local Follow

	before_each(function()
		_G.ForeverSixtyDB = {}
		helper.load("Codec")
		helper.load("Export")
		Follow = helper.load("Follow")
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
	end)

	it("is the default build with nothing saved, marked default", function()
		local build = Follow.chooseDefault(DATA, "warrior-fury", 13)
		assert.are.equal("Fury leveling", build.name)
		assert.is_true(build.default)
		assert.are.same({}, _G.ForeverSixtyDB)
	end)

	it("replaces a saved build in the active slot with the default's own code, so the choice survives a login",
		function()
		local key = require("Export").characterKey()
		_G.ForeverSixtyDB.follows = { [key] = { active = "raid", slots = {
			raid = { code = "FSB1:1.60.1.70291:warrior:111:", name = "Mine" } } } }
		Follow.restore(DATA)
		assert.are.equal("Mine", Follow.build.name)
		local build = Follow.chooseDefault(DATA, "warrior-fury", 13)
		assert.are.equal("Fury build", build.name)
		assert.are.equal("warrior-fury", build.spec)
		assert.are.equal("Fury build", _G.ForeverSixtyDB.follows[key].slots.raid.name)
		assert.are_not.equal("FSB1:1.60.1.70291:warrior:111:", _G.ForeverSixtyDB.follows[key].slots.raid.code)
		assert.is_nil(Follow.noBuildReason)
	end)

	it("says why when the spec has no build of that kind", function()
		assert.is_nil(Follow.chooseDefault(DATA, "warrior-protection", 13))
		assert.are.equal(string.format(L.diagBuildNoSpec, "Protection"), Follow.noBuildReason)
	end)
end)
