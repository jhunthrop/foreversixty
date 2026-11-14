-- addon/tests/onboarding_spec.lua
-- The first open (design/specs/2026-11-13-addon-hub.md 4.1, board hub-me-setup.png): the first login of a new
-- tour version opens the window on Me with Set up in place of Right now, never in combat or an instance;
-- /fs tour brings it back. The steps' model, when it is due (a first install, a reinstall with or without its
-- SavedVariables), the worn item step 2 shows, Set up drawn in the window, and hiding one's own Forever
-- Score (the coordinator's ruling, 2026-11-08: the outbox field, and the player's own score from the inbox).
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local SPECS = { { slug = "hunter-beast-mastery", name = "Beast Mastery" },
	{ slug = "hunter-marksmanship", name = "Marksmanship" }, { slug = "hunter-survival", name = "Survival" } }
local EXAMPLE = { slot = "wrist", slotId = 9, itemId = 19587, name = "Slashclaw Bracers" }

describe("Set up, the pure model", function()
	local Onboarding

	before_each(function()
		mock.install({})
		helper.load("Theme").reset()
		Onboarding = helper.load("Onboarding")
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
		mock.uninstall()
	end)

	-- Owner 2026-11-13: deleting the addon folder keeps SavedVariables, so a reinstall never saw the
	-- tour again. The stamp is the tour's version: a changed Set up shows once more to everyone.
	it("is due until this version is seen on the account, and never under Tooltips only", function()
		assert.are.equal(3, Onboarding.TOUR_VERSION)
		-- A first install, and a reinstall that wiped SavedVariables.
		assert.is_true(Onboarding.due(nil, false))
		assert.is_true(Onboarding.due({}, false))
		-- A reinstall that kept SavedVariables from an older tour: once more.
		assert.is_true(Onboarding.due({ welcomedAt = 1 }, false))
		assert.is_true(Onboarding.due({ welcomedAt = 1, tourSeen = 2 }, false))
		assert.is_false(Onboarding.due({ welcomedAt = 1, tourSeen = Onboarding.TOUR_VERSION }, false))
		assert.is_false(Onboarding.due({}, true))
	end)

	it("stamps the account with this version and keeps the first welcome's time", function()
		_G.ForeverSixtyDB = { welcomedAt = 7 }
		Onboarding.stamp()
		assert.are.equal(Onboarding.TOUR_VERSION, _G.ForeverSixtyDB.tourSeen)
		assert.are.equal(7, _G.ForeverSixtyDB.welcomedAt)
		_G.ForeverSixtyDB = nil
		Onboarding.stamp()
		assert.is_number(_G.ForeverSixtyDB.welcomedAt)
	end)

	it("numbers three steps, each its object and the tab it opens, with nothing done yet", function()
		local model = Onboarding.steps({ specs = SPECS, classToken = "HUNTER", example = EXAMPLE, gearOpen = true })
		assert.is_nil(model.note)
		local spec, hover, app = model.rows[1], model.rows[2], model.rows[3]
		assert.are.same({ 1, 2, 3 }, { spec.number, hover.number, app.number })
		assert.are.equal(L.hubPickSpec, spec.line)
		assert.are.equal("Beast Mastery · Marksmanship · Survival", spec.sub)
		assert.are.equal("HUNTER", spec.crest)
		assert.is_nil(spec.icon)
		assert.are.equal("builds", spec.tab)
		assert.are.equal("Hover an item you wear", hover.line)
		assert.are.equal("Slashclaw Bracers, below", hover.sub)
		assert.are.equal(19587, hover.itemId)
		assert.are.equal("gear", hover.tab)
		assert.are.equal("Get the companion app", app.line)
		assert.are.equal("foreversixty.gg/app", app.link)
		assert.is_nil(app.tab)
		assert.are.same({ false, false, false }, { spec.done, hover.done, app.done })
		assert.are.equal(EXAMPLE, model.example)
	end)

	it("ticks a step once it is done: the build followed, a hover, the companion", function()
		local model = Onboarding.steps({ specs = SPECS, classToken = "HUNTER", build = {}, buildName = "Beast Mastery raid",
			specIcon = "bm", example = EXAMPLE, gearOpen = true, hovered = true, companion = true })
		assert.are.equal("3 of 3 done", model.note)
		assert.is_true(model.rows[1].done)
		assert.are.equal("Beast Mastery raid", model.rows[1].sub)
		assert.are.equal("bm", model.rows[1].icon)
		assert.is_nil(model.rows[1].crest)
		assert.is_true(model.rows[2].done)
		assert.is_true(model.rows[3].done)
		assert.is_nil(model.rows[3].link)
		assert.are.equal("1 of 3 done", Onboarding.steps({ specs = SPECS, companion = true }).note)
	end)

	it("drops step 2's Gear link below level 10 and its item with nothing worn", function()
		local hover = Onboarding.steps({ specs = SPECS, gearOpen = false }).rows[2]
		assert.is_nil(hover.tab)
		assert.is_nil(hover.sub)
		assert.is_nil(hover.itemId)
	end)
end)

describe("Set up on the Me tab, the model", function()
	local MeTab

	before_each(function()
		mock.install({ class = { name = "Hunter", token = "HUNTER" }, level = 60, playerName = "Bob" })
		helper.load("Theme")
		MeTab = helper.load("MeTab")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("takes Right now's place, leaves the companion line of the raid out and footers /fs tour", function()
		local setup = { note = "3 steps", rows = {} }
		local leads = { { tab = "gear", label = "Gear", lead = { line = "Gloves" } } }
		local model = MeTab.model({ level = 60, companion = false, pack = { absent = true }, leads = leads,
			setup = setup })
		assert.are.equal(setup, model.setup)
		assert.are.same({}, model.rightNow)
		assert.is_nil(model.lastRaid)
		assert.are.equal("/fs tour", model.footer)
		local calm = MeTab.model({ level = 60, companion = false, pack = { absent = true }, leads = leads })
		assert.are.equal(1, #calm.rightNow)
		assert.are.equal(L.meNoRaid, calm.lastRaid.empty)
	end)
end)

describe("the worn item step 2 shows", function()
	local TourExample

	before_each(function()
		mock.install({})
		helper.load("Theme").reset()
		TourExample = helper.load("TourExample")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("is the worn item Gear's top upgrade replaces, else the main hand, chest or legs, else any", function()
		local worn = { wrist = { itemId = 1 }, legs = { itemId = 2 }, head = { itemId = 3 } }
		assert.are.equal("wrist", TourExample.choose(worn, { { slot = "feet" }, { slot = "wrist" } }))
		assert.are.equal("legs", TourExample.choose(worn, {}))
		assert.are.equal("head", TourExample.choose({ head = { itemId = 3 } }, nil))
		assert.is_nil(TourExample.choose({}, {}))
	end)
end)

describe("Set up in the window", function()
	local state, Hub, MeFrame, Onboarding, Theme

	local function load()
		Theme = helper.load("Theme")
		Theme.reset()
		helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Export")
		helper.load("Follow")
		helper.load("BuildDialog")
		helper.load("MeTab")
		-- The Builds tab first (its load gives a fresh window), then Me into that same window.
		package.loaded.MeFrame = nil
		MeFrame = require("MeFrame")
		Hub = require("Hub")
		Onboarding = helper.load("Onboarding")
		Onboarding.register()
	end

	before_each(function()
		state = mock.install({ class = { name = "Hunter", token = "HUNTER" }, level = 60, playerName = "Bob",
			realm = "Ashbringer", region = 1,
			equipped = { [9] = "|cff0070dd|Hitem:19587::::::::60:::::|h[Slashclaw Bracers]|h|r" },
			itemStats = { [19587] = { __name = "Slashclaw Bracers", __quality = 3 } },
			globals = { GetItemInfoInstant = function(link)
				return tonumber(tostring(link):match("item:(%d+)"))
			end } })
		_G.ForeverSixtyDB = nil
		load()
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
		mock.uninstall()
	end)

	it("opens on Me with Set up in Right now's place, the worn item's tooltip under it, and stamps", function()
		assert.is_true(Onboarding.open())
		assert.are.equal("me", Hub.active)
		local section = Onboarding.section
		assert.are.equal(L.meSetUp, section.heading:GetText())
		assert.is_true(section.heading:IsShown())
		assert.is_false(MeFrame.parts.now.heading:IsShown())
		assert.are.equal(L.hubPickSpec, section.rows[1].line:GetText())
		-- The spec not picked yet: the class's ringed crest, never a square icon.
		assert.is_true(section.rows[1].crest:IsShown())
		assert.is_false(section.rows[1].cell:IsShown())
		assert.is_true(section.rows[2].cell:IsShown())
		assert.are.equal("Slashclaw Bracers, below", section.rows[2].sub:GetText())
		assert.is_true(section.rows[3].address:IsShown())
		assert.is_true(section.tooltip:IsShown())
		assert.are.equal(1, mock.countCalls(section.tooltip, "SetInventoryItem"))
		assert.is_true(require("ItemTooltip").embedded[section.tooltip])
		assert.are.equal(L.meTourFoot, Hub.footer.enter:GetText())
		assert.are.equal(Onboarding.TOUR_VERSION, _G.ForeverSixtyDB.tourSeen)
		assert.are.same({}, state.clientWrites)
	end)

	it("goes back to Right now once the window hides", function()
		Onboarding.open()
		Hub.close()
		assert.is_false(Onboarding.showing)
		Hub.open("me")
		assert.is_false(Onboarding.section.heading:IsShown())
		assert.is_false(Onboarding.section.tooltip:IsShown())
	end)

	it("opens a step's tab from its gold link", function()
		Onboarding.open()
		local row = Onboarding.section.rows[1]
		row.open:GetScript("OnClick")(row.open)
		assert.are.equal("builds", Hub.active)
	end)

	it("ticks step 2 after the player's own hover, never after its own picture's", function()
		Onboarding.open()
		assert.is_true(Onboarding.section.rows[2].number:IsShown())
		assert.is_true(Onboarding.noteHover())
		assert.is_true(_G.ForeverSixtyDB.tourHovered)
		assert.is_false(Onboarding.section.rows[2].number:IsShown())
		assert.is_true(Onboarding.section.rows[2].tick:IsShown())
		assert.is_false(Onboarding.noteHover())
	end)

	it("never opens in combat or an instance; it shows when combat ends or the zone changes", function()
		Theme.inCombat = function() return true end
		assert.is_false(Onboarding.open())
		assert.is_false(Hub.isShown())
		assert.is_nil(_G.ForeverSixtyDB)
		Theme.inCombat = function() return false end
		Theme.instanceName = function() return "Molten Core" end
		assert.is_false(Onboarding.onEvent("PLAYER_REGEN_ENABLED"))
		assert.are.equal("outside", Onboarding.pending)
		Theme.instanceName = function() return nil end
		assert.is_true(Onboarding.onEvent("ZONE_CHANGED_NEW_AREA"))
		assert.are.equal("me", Hub.active)
	end)

	it("opens five seconds into a fresh install's first login, and not on the next", function()
		assert.is_true(Hub.afterEnteringWorld(true, false))
		assert.is_false(Hub.isShown())
		mock.runTimers(state)
		assert.is_true(Hub.isShown())
		assert.is_true(Onboarding.section.heading:IsShown())
		Hub.close()
		Hub.afterEnteringWorld(true, false)
		mock.runTimers(state)
		assert.is_false(Hub.isShown())
		-- A loading screen asks nothing.
		assert.is_false(Hub.afterEnteringWorld(false, false))
	end)

	it("opens on a reinstall that kept an older tour's stamp, and never under Tooltips only", function()
		_G.ForeverSixtyDB = { welcomedAt = 1793000000, tourSeen = 2 }
		Hub.afterEnteringWorld(true, false)
		mock.runTimers(state)
		assert.is_true(Hub.isShown())
		Hub.close()
		_G.ForeverSixtyDB = { prefs = { tooltipsOnly = true } }
		require("Prefs").flag = function(key)
			return key == "tooltipsOnly"
		end
		Hub.afterEnteringWorld(true, false)
		mock.runTimers(state)
		assert.is_false(Hub.isShown())
	end)
end)

describe("hiding one's own Forever Score", function()
	before_each(function()
		mock.install({ playerName = "Bow", realm = "Ashbringer" })
		helper.load("Theme").reset()
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("is on by default, and the outbox says score_visible false only when it is off", function()
		local Prefs = helper.load("Prefs")
		local Export = helper.load("Export")
		assert.is_true(Prefs.DEFAULTS.showScore)
		assert.is_nil(Export.outbox({ character = "Bow-Ashbringer", code = "FS1:x", wallNow = 0 }).export.score_visible)
		assert.is_false(Export.outbox({ character = "Bow-Ashbringer", code = "FS1:x", wallNow = 0, scoreVisible = false })
			.export.score_visible)
	end)

	-- The API, the companion and the addon each pin docs/contracts/fixtures/inbox-own-score.json.
	it("reads the shared own_score fixture as a record", function()
		local json = require("dkjson")
		local text
		for _, root in ipairs({ "../docs/contracts/fixtures/", "docs/contracts/fixtures/" }) do
			local file = io.open(root .. "inbox-own-score.json", "rb")
			if file then
				text = file:read("a")
				file:close()
				break
			end
		end
		assert(text, "no inbox-own-score.json fixture")
		local message = json.decode(text)
		local Score = helper.load("Score")
		local record = Score.ownRecord({ format = 2, messages = { message } }, message.character)
		assert.are.equal(message.score, record.score)
		assert.are.equal(message.tier, record.tier)
		assert.are.equal(message.as_of, record.as_of)
		for _, part in ipairs({ "gear", "experience", "parses", "attendance", "discipline" }) do
			assert.are.equal(message[part], record.components[part], part)
		end
	end)

	it("reads the player's own score from the inbox's own_score when the Data addon leaves them out", function()
		local Score = helper.load("Score")
		local inbox = { format = 2, messages = { { type = "own_score", character = "us/normal/bow", score = 1240,
			tier = "rare", as_of = "2026-11-07", fights = 9, gear = 380, parses = 300 } } }
		local record = Score.ownRecord(inbox, "us/normal/bow")
		assert.are.equal(1240, record.score)
		assert.are.equal(380, record.components.gear)
		assert.is_nil(record.components.experience)
		assert.is_nil(Score.ownRecord({ format = 2, messages = {} }, "us/normal/bow"))
	end)
end)
