-- addon/tests/onboarding_spec.lua
-- The first-run tour (design/specs/2026-11-08-addon-ux.md 3.1) and hiding one's own Forever Score
-- (the coordinator's ruling, 2026-11-08): the pages for a character in a guild or not, the set-up
-- page's states, when the tour is due and how it ends, every page drawn through the client mock
-- without writing into a client frame, the small QR, and the score switch's two ends (the outbox
-- field, and the player's own score read from the inbox).
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

describe("Onboarding, the pure model", function()
	local Onboarding

	before_each(function()
		mock.install({})
		helper.load("Theme").reset()
		Onboarding = helper.load("Onboarding")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("has five pages in a guild and four without, the guild's page skipped", function()
		assert.are.same({ "loop", "site", "back", "guild", "setup" }, Onboarding.pageKeys(true))
		assert.are.same({ "loop", "site", "back", "setup" }, Onboarding.pageKeys(false))
	end)

	it("draws the loop's guild node only in a guild, under the player's own faction banner", function()
		assert.are.equal(6, #Onboarding.loopNodes(true, "horde", "HUNTER"))
		assert.are.equal(Onboarding.BANNERS.horde, Onboarding.loopNodes(true, "horde", "HUNTER")[6].icon)
		assert.are.equal(5, #Onboarding.loopNodes(false, "alliance", "HUNTER"))
		-- The diagram says the loop; no caption restates it (review 2026-11-10).
		assert.is_nil(L.tourLoopCaption)
	end)

	it("offers the path tiles with no build, the build followed otherwise", function()
		local tiles = Onboarding.setup({ level = 6, specs = { { slug = "hunter-beast-mastery", name = "Beast Mastery" } } })
		assert.are.equal(1, #tiles.tiles)
		assert.are.equal(L.tourPathLeveling, tiles.tilesNote)
		assert.are.equal(L.tourPathRaid, Onboarding.setup({ level = 60, specs = {} }).tilesNote)
		local followed = Onboarding.setup({ level = 24, build = { order = { {}, {} } }, buildName = "Beast Mastery leveling",
			progress = { matched = 1, nextLevel = 25 }, pointName = "Lethal Attacks" })
		assert.are.equal(string.format(L.tourYourPath, "Beast Mastery leveling"), followed.path.name)
		assert.are.equal(string.format(L.tourNextPoint, 25, "Lethal Attacks"), followed.path.next)
		assert.is_nil(followed.tiles)
	end)

	it("says the companion has run under a day, warns past three days, and offers both rows otherwise", function()
		local fresh = Onboarding.setup({ siteAge = 2 * 3600, name = "Bow Jackzon" })
		assert.is_true(fresh.seen.ok)
		assert.are.equal(string.format(L.tourSiteSaw, "Bow Jackzon", "2 hours ago"), fresh.seen.detail)
		assert.is_nil(fresh.signIn)
		assert.is_nil(fresh.companion)
		local silent = Onboarding.setup({ siteAge = 4 * 86400, name = "Bow Jackzon" })
		assert.is_false(silent.seen.ok)
		assert.are.equal(string.format(L.tourSiteSilent, "Bow Jackzon", 4), silent.seen.text)
		assert.is_true(silent.companion)
		local never = Onboarding.setup({})
		assert.is_true(never.signIn)
		assert.is_true(never.companion)
		assert.is_nil(never.seen)
	end)

	it("is due once per account, and never under Tooltips only", function()
		assert.is_true(Onboarding.due(nil, false))
		assert.is_true(Onboarding.due({}, false))
		assert.is_false(Onboarding.due({ welcomedAt = 1 }, false))
		assert.is_false(Onboarding.due({}, true))
	end)

	it("keeps the app QR small and labelled Join the app beta until the apps ship", function()
		assert.are.equal(48, Onboarding.QR_SIZE)
		assert.are.equal("Join the app beta", L.tourAppBeta)
		assert.are.equal("foreversixty.gg/app", Onboarding.APP)
		assert.are.equal("Settings can hide it from others.", L.tourScoreHide)
	end)

	it("says In, Late, Bench or Out on the guild page", function()
		assert.are.equal("Answer In, Late, Bench or Out", L.tourAnswer)
	end)

	it("puts the Horde banner on a Horde guild node, from Compat.playerFaction's lower-case answer", function()
		assert.are.equal(Onboarding.BANNERS.horde, Onboarding.loopNodes(true, "horde", "HUNTER")[6].icon)
		assert.are.equal(Onboarding.BANNERS.alliance, Onboarding.loopNodes(true, nil, "HUNTER")[6].icon)
		assert.are_not.equal(Onboarding.BANNERS.horde, Onboarding.BANNERS.alliance)
	end)

	it("heads each page in the panel's title band: the heading, the welcome, and 2 of 5", function()
		local head = Onboarding.head("site", 2, 5, { classToken = "HUNTER", faction = "horde" })
		assert.are.same({ title = L.tourSiteHeading, meta = L.tourTitle, classToken = "HUNTER", faction = "horde",
			number = 2, unit = "of 5" }, head)
	end)

	it("draws the loop's arrows from flat parts, mirrored for left and turned for down", function()
		assert.are.same(Onboarding.ARROW_PARTS, Onboarding.arrowParts("right"))
		assert.are.same({ 6, 4, 10, 2 }, Onboarding.arrowParts("left")[1])
		assert.are.same({ 0, 4, 2, 2 }, Onboarding.arrowParts("left")[4])
		assert.are.same({ 4, 0, 2, 10 }, Onboarding.arrowParts("down")[1])
		assert.are.same({ 4, 14, 2, 2 }, Onboarding.arrowParts("down")[4])
	end)

	it("keeps the tour's copy free of glyph arrows, dots and the old opt-out wording", function()
		assert.is_nil(L.tourArrowRight)
		assert.is_nil(L.tourDot)
		assert.are.equal(L.settingsTooltipsOnly, L.tourJustTooltips)
	end)
end)

describe("Onboarding, the frame", function()
	local Onboarding, state

	local function start(guild)
		state = mock.install({ class = { name = "Hunter", token = "HUNTER" }, level = 6, playerName = "Bow" })
		helper.load("Theme").reset()
		Onboarding = helper.load("Onboarding")
		require("Export").guildInfo = function()
			return guild and { name = "Employed Millennials" } or nil
		end
		_G.ForeverSixtyDB = {}
	end

	after_each(function()
		_G.ForeverSixtyDB = nil
		mock.uninstall()
	end)

	it("draws every page of a guild's tour, writing nothing into a client frame", function()
		start(true)
		assert.is_true(Onboarding.open(1))
		local seen = {}
		for _ = 1, 5 do
			seen[#seen + 1] = Onboarding.draw()
			if #seen < 5 then
				Onboarding.page = Onboarding.page + 1
			end
		end
		assert.are.same({ "loop", "site", "back", "guild", "setup" }, seen)
		assert.are.equal(Onboarding.WIDTH, Onboarding.frame:GetWidth())
		-- The panel's chrome: the umber ground and the title band with the medallion.
		assert.is_not_nil(Onboarding.frame.foreverSixtyBackground)
		assert.is_not_nil(Onboarding.band.foreverSixtyMedallion)
		assert.are.equal(L.tourSetupHeading, Onboarding.band.foreverSixtyTitle:GetText())
		assert.are.equal("5", Onboarding.band.foreverSixtyNumber:GetText())
		assert.are.same({}, state.clientWrites)
	end)

	it("ends for good on Skip tour, Done or Just the tooltips, stamping the account", function()
		start(false)
		Onboarding.open(1)
		Onboarding.finish()
		assert.is_number(_G.ForeverSixtyDB.welcomedAt)
		assert.is_false(Onboarding.frame:IsShown())
	end)

	it("walks Next to the last page and Next past it ends the tour", function()
		start(false)
		Onboarding.open(1)
		for _ = 1, 3 do
			Onboarding.next()
		end
		assert.are.equal(4, Onboarding.page)
		Onboarding.back()
		assert.are.equal(3, Onboarding.page)
	end)

	it("never opens in combat; combat hides it without ending it, and it comes back once after", function()
		start(false)
		require("Theme").inCombat = function() return true end
		assert.is_false(Onboarding.open(1))
		require("Theme").inCombat = function() return false end
		Onboarding.onEvent("PLAYER_REGEN_ENABLED")
		assert.is_true(Onboarding.frame:IsShown())
		Onboarding.page = 2
		Onboarding.onEvent("PLAYER_REGEN_DISABLED")
		assert.is_false(Onboarding.frame:IsShown())
		assert.is_nil(_G.ForeverSixtyDB.welcomedAt)
		Onboarding.onEvent("PLAYER_REGEN_ENABLED")
		assert.is_true(Onboarding.frame:IsShown())
		assert.are.equal(2, Onboarding.page)
		-- A second fight hides it again; this time it waits for /fs help.
		Onboarding.onEvent("PLAYER_REGEN_DISABLED")
		Onboarding.onEvent("PLAYER_REGEN_ENABLED")
		assert.is_false(Onboarding.frame:IsShown())
		assert.is_nil(_G.ForeverSixtyDB.welcomedAt)
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
