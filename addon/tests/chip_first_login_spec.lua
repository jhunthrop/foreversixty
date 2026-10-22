-- addon/tests/chip_first_login_spec.lua
-- A.1: the Forever Sixty line is on screen at a first login and says the next step. The fresh-account
-- walk (22 Oct, level 24 High Order Skyborne hunter, the owner's saved variables, no entry for the
-- character) found nothing drawn beside the minimap though the welcome promised a line. These specs
-- run the whole boot with the walk's chip prefs: the line is drawn with the next talent in its words,
-- /fs chip hides and shows it, a saved spot with a y and no point (or a spot off the screen) never
-- leaves it off screen, an idle chip always has words, and the welcome names the line only when it
-- is there.
local helper = require("spec_helper")
local mock = require("wow_mock")

local Data = require("Data")

local SCREEN_WIDTH, SCREEN_HEIGHT = 1920, 1080

--- Trait ranks ({ [node] = rank }) for `points` points down the hunter's tab `tab`, in the data's order.
local function traitRanks(tab, points)
	local ranks = {}
	for _, talent in ipairs(Data.classes.hunter.tabs[tab].talents) do
		if points <= 0 then
			break
		end
		local rank = math.min(points, talent.maxRank)
		ranks[talent.node] = rank
		points = points - rank
	end
	return ranks
end

--- The chip prefs of the owner's saved variables the walk loaded (hiddenUntilLevel 0, shown, a saved y,
--- four alts in byCharacter and none for the character); `override` replaces fields of it.
local function walkChipPrefs(override)
	local chip = {
		hiddenUntilLevel = 0, hiddenFor = "", shown = true, pathFor = "", pathChosen = "", y = -41.9443359375,
		combatHide = true,
		byCharacter = {
			["/Classic Beta PvE 2/Lightning Bolltz"] = {},
			["/Classic Beta PvP/Lightning Boltzz"] = { pathChosen = "shaman-restoration" },
			["/Classic Beta PvE 2/Booty Twerk"] = { pathChosen = "paladin-retribution" },
			["/Classic Beta PvP/Shorty Small"] = { pathChosen = "rogue-combat" },
		},
	}
	for key, value in pairs(override or {}) do
		chip[key] = value
	end
	return chip
end

--- Boot the addon as the walk did: Bow Jackzon, level 24, the hunter's points spent, the minimap 140
--- wide, the owner's chip prefs, then PLAYER_LOGIN and the first PLAYER_ENTERING_WORLD.
local function bootWalk(chipOverride, welcomed)
	mock.install({
		class = { name = "Hunter", token = "HUNTER" },
		race = { name = "High Order Skyborne", token = "HighOrderSkyborne" },
		realm = "Classic Beta PvP",
		region = 1,
		playerName = "Bow Jackzon",
		level = 24,
		traits = { configID = 7, ranks = traitRanks(2, 15) },
	})
	_G.UIParent.width, _G.UIParent.height = SCREEN_WIDTH, SCREEN_HEIGHT
	_G.Minimap.width, _G.Minimap.height = 140, 140
	_G.ForeverSixtyDB = { ui = { chip = walkChipPrefs(chipOverride) } }
	if welcomed then
		_G.ForeverSixtyDB.welcomedAt = 1793000000
	end
	require("Theme").reset()
	for _, name in ipairs({ "Prefs", "Widgets", "Export", "Talents", "Gear", "ItemTooltip", "Follow", "TalentGlow",
		"Chip", "TalentMarks", "Minimap", "Comms", "BuildDialog", "Settings", "SiteStatus" }) do
		helper.load(name)
	end
	local Options = helper.load("Options")
	Options.register()
	Options.onEvent(Options.frame, "PLAYER_LOGIN")
	Options.onEvent(Options.frame, "PLAYER_ENTERING_WORLD", true, false)
	return Options
end

local function lastPoint(frame)
	return frame.points[#frame.points]
end

describe("the line at a first login (the fresh-account walk)", function()
	after_each(function()
		mock.uninstall()
	end)

	it("is visible and drawn, with the next talent in its words, for the walk's prefs", function()
		bootWalk()
		local Chip = require("Chip")
		assert.is_nil(Chip.disabled)
		assert.is_true(Chip.current.visible)
		assert.is_nil(Chip.current.hiddenBy)
		assert.is_true(Chip.frame:IsShown())
		assert.is_true(Chip.isDrawn())
		assert.are.equal("idle", Chip.current.chosen.kind)
		assert.is_truthy(Chip.current.chosen.text:find("^Next at %d+: ."))
		assert.are.equal(Chip.current.chosen.text, Chip.text:GetText())
		assert.is_true(Chip.text:IsShown())
		-- The words are not dimmed away with the icon.
		assert.are.equal(1, Chip.current.iconAlpha)
	end)

	it("sits beside the minimap, not at the saved y", function()
		bootWalk()
		local Chip = require("Chip")
		local point = lastPoint(Chip.frame)
		assert.are.equal("TOPRIGHT", point[1])
		assert.are.equal(_G.Minimap, point[2])
		assert.are.equal(-(140 + require("Theme").SIZES.chipMinimapGap), point[4])
		assert.are.equal(0, point[5])
	end)

	it("is raised above the client's buff icons", function()
		bootWalk()
		local Chip = require("Chip")
		local raised
		for _, call in ipairs(Chip.frame.calls) do
			if call.method == "SetFrameLevel" then
				raised = call[1]
			end
		end
		assert.are.equal(Chip.FRAME_LEVEL, raised)
	end)

	it("hides on /fs chip from a shown chip, saying until level 25, and shows again", function()
		local Options = bootWalk()
		local Chip, L = require("Chip"), require("Locale")
		assert.is_true(Chip.frame:IsShown())
		assert.are.same({ string.format(L.chipHiddenLine, 25) }, Options.handle("chip"))
		assert.is_true(Chip.isHidden())
		assert.are.equal("player", Chip.current.hiddenBy)
		assert.is_false(Chip.frame:IsShown())
		assert.are.same({ L.chipShownLine }, Options.handle("chip"))
		assert.is_false(Chip.isHidden())
		assert.is_true(Chip.frame:IsShown())
		assert.is_true(Chip.isDrawn())
	end)

	it("keeps the line when the prefs hold a y and an x but no point", function()
		bootWalk({ x = -224.7222900390625, point = "" })
		local Chip = require("Chip")
		assert.are.equal(_G.Minimap, lastPoint(Chip.frame)[2])
		assert.is_true(Chip.isDrawn())
	end)
end)

describe("Chip.place", function()
	after_each(function()
		mock.uninstall()
	end)

	local function placedAgainst(prefs)
		bootWalk(prefs)
		local Chip = require("Chip")
		Chip.place()
		return lastPoint(Chip.frame)
	end

	it("never uses a saved y without a point", function()
		local point = placedAgainst({ point = "", x = 0, y = -41.9443359375 })
		assert.are.equal(_G.Minimap, point[2])
	end)

	it("uses a dragged spot that is on the screen", function()
		local point = placedAgainst({ point = "TOPRIGHT", x = -224.7222900390625, y = -41.9443359375 })
		assert.are.equal("TOPRIGHT", point[1])
		assert.are.equal(_G.UIParent, point[2])
		assert.are.equal(-224.7222900390625, point[4])
		assert.are.equal(-41.9443359375, point[5])
	end)

	it("falls back beside the minimap for a spot outside the screen", function()
		local wide = placedAgainst({ point = "TOPRIGHT", x = -4000, y = -41 })
		assert.are.equal(_G.Minimap, wide[2])
		mock.uninstall()
		local tall = placedAgainst({ point = "TOPRIGHT", x = -224, y = 9000 })
		assert.are.equal(_G.Minimap, tall[2])
	end)

	it("puts a spot the client lays off screen back beside the minimap instead of switching the line off", function()
		bootWalk({ point = "TOPRIGHT", x = -224.7222900390625, y = -41.9443359375 })
		local Chip, Prefs, Theme, L = require("Chip"), require("Prefs"), require("Theme"), require("Locale")
		local frame = Chip.frame
		-- The client lays the dragged spot out left of the screen's edge; the default spot is fine.
		frame.GetLeft = function(own)
			return lastPoint(own)[2] == _G.UIParent and -30 or 1500
		end
		assert.is_true(Chip.selfCheck())
		assert.is_nil(Chip.disabled)
		assert.are.equal("", Prefs.get("chip", "point"))
		assert.are.equal(_G.Minimap, lastPoint(frame)[2])
		assert.is_true(frame:IsShown())
		local noted = false
		for _, line in ipairs(Theme.infoNotes()) do
			noted = noted or line == L.diagChipMoved
		end
		assert.is_true(noted)
	end)

	it("still switches the line off when even the default spot is off screen", function()
		bootWalk()
		local Chip = require("Chip")
		Chip.frame.GetLeft = function()
			return -30
		end
		assert.is_false(Chip.selfCheck())
		assert.is_true(Chip.disabled)
		assert.is_false(Chip.isDrawn())
	end)
end)

describe("the idle chip's words", function()
	local Chip, L

	setup(function()
		Chip, L = require("Chip"), require("Locale")
	end)

	local function idleInputs(overrides)
		local base = {
			level = 24, classIcon = "crest-hunter", now = 100, clock = 0, shown = true, combatHide = true,
			unspent = 0,
			build = { name = "Beast Mastery build", order = { {}, {}, {} } },
			progress = { matched = 1, total = 3, nextLevel = 25, point = { tab = 1, tier = 1, column = 2, index = 2 } },
			pointName = "Endurance Training", pointRank = { have = 0, want = 5 },
		}
		for key, value in pairs(overrides or {}) do
			-- false takes the field away (a nil cannot be set in a table constructor).
			base[key] = value or nil
		end
		return base
	end

	it("reads Next: the talent and its rank when a build is followed and the next talent is known", function()
		local chosen = Chip.model(idleInputs()).chosen
		assert.are.equal("idle", chosen.kind)
		assert.are.equal("Next: Endurance Training (1 of 5)", chosen.text)
		assert.are.equal("talents", chosen.action)
	end)

	it("reads the talent alone when the plan wants one rank of it", function()
		local model = Chip.model(idleInputs({ pointRank = { have = 0, want = 1 } }))
		assert.are.equal("Next: Endurance Training", model.chosen.text)
	end)

	it("reads Pick a build: /fs, and a click opens the build window, when nothing is followed", function()
		local chosen = Chip.model(idleInputs({ build = false, progress = { matched = 0, total = 0 },
			pointName = false, pointRank = false })).chosen
		assert.are.equal("idle", chosen.kind)
		assert.are.equal("Pick a build: /fs", chosen.text)
		assert.are.equal(L.chipIdlePick, chosen.text)
		assert.are.equal("inbox", chosen.action)
		assert.are.equal(L.chipHintBuildWindow, chosen.hint)
	end)

	it("says a build is waiting when none is followed and the site sent one", function()
		local chosen = Chip.model(idleInputs({ build = false, inboxWaiting = 1, progress = { matched = 0, total = 0 },
			pointName = false })).chosen
		assert.are.equal(L.chipIdleInbox, chosen.text)
	end)

	it("stays the class icon alone, dimmed, for a finished build with nothing to say", function()
		local model = Chip.model(idleInputs({ progress = { matched = 3, total = 3 }, pointName = false }))
		assert.are.equal("", model.chosen.text)
		assert.are.equal(Chip.IDLE_ALPHA, model.iconAlpha)
	end)

	it("is cut to fit like every other line when the talent's name is long", function()
		local text = Chip.model(idleInputs({ pointName = ("Improved Mortal Strike "):rep(3) })).chosen.text
		local fitted = Chip.fit(text, function(value)
			return #value * 6
		end, 242)
		assert.is_true(#fitted * 6 <= 242)
	end)
end)

describe("the first-login welcome", function()
	after_each(function()
		mock.uninstall()
	end)

	it("names the line beside the minimap when it is drawn", function()
		bootWalk(nil, true)
		local L = require("Locale")
		assert.are.same({ L.welcomeOn, L.welcomeNext, L.welcomeSite }, require("SiteStatus").welcomeLines())
	end)

	it("says what is there when the line is hidden or switched off", function()
		local Options = bootWalk(nil, true)
		local L, Chip = require("Locale"), require("Chip")
		Options.handle("chip")
		assert.are.same({ L.welcomeOnNoLine, L.welcomeNext, L.welcomeSite }, require("SiteStatus").welcomeLines())
		Options.handle("chip")
		Chip.disabled = true
		assert.are.same({ L.welcomeOnNoLine, L.welcomeNext, L.welcomeSite }, require("SiteStatus").welcomeLines())
		assert.is_nil(L.welcomeOnNoLine:find("minimap", 1, true))
	end)
end)
