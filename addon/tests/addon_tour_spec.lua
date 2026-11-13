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

describe("the Builds dialog's own builds (2026-11-13 spec 1), the pure model", function()
	local BuildsPane, BuildDialog

	before_each(function()
		helper.load("Export")
		helper.load("Follow")
		BuildsPane = helper.load("BuildsPane")
		BuildDialog = helper.load("BuildDialog")
	end)

	it("lists every spec's level 60 build and raid build, in tree order, labelled for what each is for",
		function()
		local blocks = BuildsPane.blocks(DATA, "warrior", 13, nil)
		assert.are.equal(2, #blocks)
		assert.are.equal("warrior-arms", blocks[1].spec)
		assert.are.equal("Arms", blocks[1].name)
		local arms = blocks[1].rows
		assert.are.equal(2, #arms)
		assert.are.equal(L.dialogKindLeveling, arms[1].label)
		assert.are.equal("Level 60 build", L.dialogKindLeveling)
		assert.are.equal("Arms 3", arms[1].detail)
		assert.are.equal(L.dialogKindRaid, arms[2].label)
		assert.are.equal("Arms 1 / Fury 1", arms[2].detail)
		assert.is_false(arms[1].following)
		-- The level's own kind is followed as the level says; the other is chosen outright.
		assert.is_nil(arms[1].choose)
		assert.are.equal("raid", arms[2].choose)
		assert.are.equal(4, #BuildsPane.flatten(blocks))
	end)

	it("shows one row for a spec whose two builds spend the same points", function()
		local data = { classes = DATA.classes, defaultBuilds = { ["warrior-fury"] = {
			leveling = { name = "Fury leveling", order = { { 2, 1, 3 } }, why = "Cruelty first." },
			raid = { name = "Fury raid", order = { { 2, 1, 3 } }, why = "Cruelty first." },
		} } }
		local rows = BuildsPane.blocks(data, "warrior", 60, nil)[1].rows
		assert.are.equal(1, #rows)
		assert.are.equal("both", rows[1].kind)
		assert.are.equal(L.dialogKindBoth, rows[1].label)
		assert.are.equal("Cruelty first.", rows[1].why)
		assert.is_nil(rows[1].choose)
	end)

	it("marks the row the followed build spends, with the next point under it", function()
		local followed = { classSlug = "warrior", order = { { tab = 2, tier = 1, column = 3 },
			{ tab = 2, tier = 1, column = 3 } } }
		local rows = BuildsPane.flatten(BuildsPane.blocks(DATA, "warrior", 13, followed, {}))
		assert.is_false(rows[1].following)
		assert.is_true(rows[3].following)
		assert.are.equal("Cruelty · " .. require("Theme").inlineColor("clientGold", string.format(L.panelSpendNow, 4)),
			rows[3].next)
		assert.is_nil(rows[1].next)
		assert.are.equal(3, BuildsPane.rowLines(rows[3]))
	end)

	it("reads the why and the spec icon from the data, and goes without either when the data has none", function()
		local data = { classes = DATA.classes, specIcons = { ["warrior-arms"] = "ability_warrior_savageblow" },
			defaultBuilds = { ["warrior-arms"] = { leveling = { name = "Arms leveling", order = { { 1, 1, 1 } },
				why = "Deep Wounds carries the leveling." } } } }
		local block = BuildsPane.blocks(data, "warrior", 20, nil)[1]
		assert.are.equal("Interface\\Icons\\ability_warrior_savageblow", block.icon)
		assert.are.equal("Deep Wounds carries the leveling.", block.rows[1].why)
		local bare = BuildsPane.blocks(DATA, "warrior", 20, nil)[1]
		assert.is_nil(bare.icon)
		assert.is_nil(bare.rows[1].why)
		assert.are.equal(2, BuildsPane.rowLines(bare.rows[1]))
	end)

	it("titles the window with the class's builds and the level, and heads the list with the question", function()
		assert.are.equal("Warrior builds", BuildsPane.title("Warrior"))
		local spentAll = { [3] = { ["1:2"] = 4 } }
		assert.are.equal("Level 13", BuildsPane.note(13, nil, spentAll))
		local followed = { classSlug = "warrior", order = { { tab = 1, tier = 1, column = 1 },
			{ tab = 1, tier = 1, column = 1 }, { tab = 1, tier = 1, column = 2 } } }
		-- Parts that add up to the level's points: on plan, off plan, the points waiting in the client gold
		-- (wow-player 2026-11-13, rounds 1 and 2).
		local Theme = require("Theme")
		assert.are.equal("Level 13 · 0 on plan · " .. Theme.inlineColor("clientGold", "4 to spend"),
			BuildsPane.note(13, followed, {}))
		assert.is_nil(BuildsPane.question(nil, "1 Arms · 1 Fury"))
		assert.are.equal("Are you Arms or Fury?", BuildsPane.question("Are you Arms or Fury?"))
		assert.are.equal("Are you Arms or Fury? · " .. Theme.inlineColor("muted", "1 Arms · 1 Fury"),
			BuildsPane.question("Are you Arms or Fury?", "1 Arms · 1 Fury"))
	end)

	it("counts the points off plan in the note, in the warning colour", function()
		local followed = { classSlug = "warrior", order = { { tab = 1, tier = 1, column = 1 },
			{ tab = 1, tier = 1, column = 1 }, { tab = 1, tier = 1, column = 2 } } }
		local ranks = { [2] = { ["1:3"] = 2 } }
		local note = BuildsPane.note(13, followed, ranks)
		assert.is_truthy(note:find(require("Theme").inlineColor("warning", string.format(L.windowOffPlanCount, 2)),
			1, true), note)
	end)

	it("titles the row tooltip with the label, then the split, the whole why, and Click to follow", function()
		local Theme = require("Theme")
		local row = { label = L.dialogKindRaid, detail = "Arms 1 / Fury 1", why = "Mortal Strike carries it.",
			following = false }
		local lines = BuildsPane.tooltipLines(row)
		assert.are.equal(Theme.inlineColor("strong", L.dialogKindRaid), lines[1])
		assert.are.equal(Theme.inlineColor("muted", "Arms 1 / Fury 1"), lines[2])
		assert.are.same({ text = Theme.inlineColor("body", "Mortal Strike carries it."), wrap = true }, lines[3])
		assert.are.equal(Theme.inlineColor("muted", L.dialogClickFollow), lines[4])
		row.following = true
		assert.are.equal(3, #BuildsPane.tooltipLines(row))
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

describe("the Builds dialog's own builds (2026-11-13 spec 1), the frame", function()
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
		helper.load("BuildsPane")
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

	it("opens on the question, every spec's two builds with Follow and From the site",
		function()
		assert.is_true(BuildDialog.open("paste"))
		mock.runTimers(state)
		local paste, pane = BuildDialog.paste, BuildDialog.paste.builds
		assert.are.equal("Are you Arms or Fury? · " .. require("Theme").inlineColor("muted", "1 Arms · 1 Fury"),
			pane.heading:GetText())
		assert.is_true(pane.heading:IsShown())
		-- The window's header carries the crest and the title (spec 2026-11-13-addon-builds-window 3.1).
		local header = require("Hub").header
		assert.is_true(header.crest:IsShown())
		assert.is_true(header.ring:IsShown())
		assert.are.equal("Warrior builds", header.title:GetText())
		assert.are.equal("Level 13 · " .. require("Theme").inlineColor("clientGold", "2 to spend"), header.note:GetText())
		assert.is_true(paste.site.heading:IsShown())
		assert.are.equal(L.dialogSiteHeading, paste.site.heading:GetText())
		assert.are.equal(L.dialogKindLeveling, pane.rows[1].label:GetText())
		assert.are.equal(L.dialogKindRaid, pane.rows[2].label:GetText())
		for index = 1, 4 do
			assert.is_true(pane.rows[index].frame:IsShown(), index)
			assert.is_true(pane.rows[index].follow:IsShown(), index)
			assert.is_false(pane.rows[index].mark:IsShown(), index)
		end
		assert.is_false(pane.rows[5].frame:IsShown())
		assert.is_true(pane.specs[1].name:IsShown())
		assert.is_true(pane.specs[2].name:IsShown())
		assert.is_false(pane.rows[1].edge:IsShown())
		assert.are.equal(0, BuildDialog.cursor)
		assert.is_false(require("Hub").footer.enter:IsShown())
		assert.are.equal(L.windowEscClose, require("Hub").footer.escape:GetText())
	end)

	it("follows an own build on one tap, remembers the answer, says so and stays open on it", function()
		BuildDialog.open("paste")
		mock.runTimers(state)
		local build = BuildDialog.followOwn(BuildDialog.paste.builds.rows[3].entry)
		assert.is_not_nil(build)
		assert.are.equal("Fury level 60 build", Follow.buildName(DATA, Follow.build))
		assert.is_true(Follow.build.default)
		assert.are.same({ spec = "warrior-fury" }, Chip.chosenPath())
		assert.is_true(BuildDialog.frame:IsShown())
		assert.is_truthy(said[1]:find("Fury level 60 build", 1, true), said[1])
		-- At once: Fury's level 60 build reads Following with its next point; the others offer Follow; Up
		-- next lists its points under its name.
		local rows = BuildDialog.paste.builds.rows
		assert.are.equal("Fury level 60 build", BuildDialog.paste.upNext.note:GetText())
		assert.is_true(BuildDialog.paste.upNext.rows[1].name:IsShown())
		assert.is_true(rows[3].mark:IsShown())
		assert.is_false(rows[3].follow:IsShown())
		assert.is_truthy(rows[3].lines[2]:GetText():find("Cruelty", 1, true))
		assert.is_true(rows[1].follow:IsShown())
		assert.is_false(BuildDialog.paste.following:IsShown())
		-- Enter on an own row follows it; the hint says so.
		BuildDialog.setCursor(1)
		assert.are.equal(L.dialogEnterFollow, require("Hub").footer.enter:GetText())
		assert.is_true(rows[1].edge:IsShown())
		BuildDialog.previewCursor()
		assert.are.equal("Arms level 60 build", Follow.buildName(DATA, Follow.build))
	end)

	it("follows the raid build below 60 and remembers the kind for the next login", function()
		BuildDialog.open("paste")
		mock.runTimers(state)
		BuildDialog.followOwn(BuildDialog.paste.builds.rows[2].entry)
		assert.are.equal("raid", Follow.build.kind)
		assert.are.equal("Arms raid build", Follow.buildName(DATA, Follow.build))
		assert.are.same({ spec = "warrior-arms", kind = "raid" }, Chip.chosenPath())
	end)

	it("lays the followed row on gold with a gold edge, and drops a playstyle that does not fit whole", function()
		Follow.ensureDefault(DATA, { spec = "warrior-arms", classSlug = "warrior", ranks = {}, level = 13 })
		state.stringWidth = function(text)
			return #text * 8
		end
		BuildDialog.open("paste")
		mock.runTimers(state)
		local pane = BuildDialog.paste.builds
		assert.is_true(pane.rows[1].followed:IsShown())
		assert.is_true(pane.rows[1].followedEdge:IsShown())
		assert.is_false(pane.rows[2].followed:IsShown())
		-- "Two-handers, Sweeping Strikes and Mortal Strike" after "Arms" runs past the room at 8 px a letter.
		assert.is_false(pane.specs[1].playstyle:IsShown())
		assert.is_true(pane.specs[2].playstyle:IsShown())
	end)

	it("does nothing on the row already followed", function()
		Follow.ensureDefault(DATA, { spec = "warrior-arms", classSlug = "warrior", ranks = {}, level = 13 })
		BuildDialog.open("paste")
		mock.runTimers(state)
		assert.is_true(BuildDialog.paste.builds.rows[1].mark:IsShown())
		local before = Follow.build
		BuildDialog.followOwn(BuildDialog.paste.builds.rows[1].entry)
		assert.is_true(BuildDialog.frame:IsShown())
		assert.are.equal(before, Follow.build)
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
		assert.are.equal("Fury level 60 build", build.name)
		assert.are.equal("warrior-fury", build.spec)
		assert.are.equal("Fury level 60 build", _G.ForeverSixtyDB.follows[key].slots.raid.name)
		assert.are_not.equal("FSB1:1.60.1.70291:warrior:111:", _G.ForeverSixtyDB.follows[key].slots.raid.code)
		assert.is_nil(Follow.noBuildReason)
	end)

	it("says why when the spec has no build of that kind", function()
		assert.is_nil(Follow.chooseDefault(DATA, "warrior-protection", 13))
		assert.are.equal(string.format(L.diagBuildNoSpec, "Protection"), Follow.noBuildReason)
	end)
end)
