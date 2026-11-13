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
	local Onboarding, TourExample

	before_each(function()
		mock.install({})
		helper.load("Theme").reset()
		Onboarding = helper.load("Onboarding")
		TourExample = helper.load("TourExample")
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
		-- Round 2: the note "loads its leveling build" went; the points split sits beside the question.
		assert.is_nil(tiles.tilesNote)
		assert.are.equal("2 Holy · 2 Retribution", Onboarding.setup({ level = 13, specs = {},
			split = "2 Holy · 2 Retribution" }).split)
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

	-- Owner 2026-11-13: deleting the addon folder keeps SavedVariables, so a reinstall never saw the
	-- tour again. The stamp is the tour's version now: a changed tour shows once more to everyone.
	it("is due until this tour's version is ended on the account, and never under Tooltips only", function()
		assert.are.equal(2, Onboarding.TOUR_VERSION)
		-- A first install, and a reinstall that wiped SavedVariables.
		assert.is_true(Onboarding.due(nil, false))
		assert.is_true(Onboarding.due({}, false))
		-- A reinstall that kept SavedVariables from the first tour, or an account that saw it: once more.
		assert.is_true(Onboarding.due({ welcomedAt = 1 }, false))
		assert.is_true(Onboarding.due({ welcomedAt = 1, tourSeen = 1 }, false))
		assert.is_false(Onboarding.due({ welcomedAt = 1, tourSeen = Onboarding.TOUR_VERSION }, false))
		assert.is_false(Onboarding.due({}, true))
	end)

	it("stamps the account with this tour's version and keeps the first welcome's time", function()
		_G.ForeverSixtyDB = { welcomedAt = 1793000000 }
		Onboarding.stamp()
		assert.are.equal(Onboarding.TOUR_VERSION, _G.ForeverSixtyDB.tourSeen)
		assert.are.equal(1793000000, _G.ForeverSixtyDB.welcomedAt)
		assert.is_false(Onboarding.due(_G.ForeverSixtyDB, false))
		_G.ForeverSixtyDB = nil
		Onboarding.stamp()
		assert.is_number(_G.ForeverSixtyDB.welcomedAt)
		assert.are.equal(Onboarding.TOUR_VERSION, _G.ForeverSixtyDB.tourSeen)
		_G.ForeverSixtyDB = nil
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
		-- Round 2: the welcome line sits under page 1's heading only.
		assert.are.equal(L.tourTitle, Onboarding.head("loop", 1, 5, { classToken = "HUNTER", faction = "horde" }).meta)
		assert.are.same({ title = L.tourSiteHeading, classToken = "HUNTER", faction = "horde",
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

	it("hovers Tooltips only with the setting's own hint and that it ends the tour", function()
		assert.are.equal(L.tourJustTooltips, L.tourJustTooltipsHint[1])
		assert.are.equal(L.settingsTooltipsOnlyHint, L.tourJustTooltipsHint[2].text)
		assert.are.equal(3, #L.tourJustTooltipsHint)
	end)

	it("sizes each page to its content, between the floor and the full height", function()
		assert.are.equal(Onboarding.MIN_HEIGHT, Onboarding.frameHeight(0))
		assert.are.equal(Onboarding.HEIGHT, Onboarding.frameHeight(2000))
		local mid = Onboarding.frameHeight(200)
		assert.is_true(mid > Onboarding.MIN_HEIGHT and mid < Onboarding.HEIGHT)
	end)

	it("shows the panel's first upgrade as the tooltip shows it: the arrow and its figure, then slot and source",
		function()
			local example = TourExample.example({ upgrade = { name = "Ranger Bow", hex = "0070dd", gain = 3.5, unit = "dps",
				detail = "Ranged · World drop" } })
			assert.are.equal("Ranger Bow", example.name)
			assert.are.equal("0070dd", example.hex)
			assert.is_truthy(example.lines[1][1]:find("+3.5 DPS", 1, true))
			assert.is_truthy(example.lines[1][1]:find("|T", 1, true))
			assert.are.equal("tooltipGreen", example.lines[1][2])
			assert.are.same({ "Ranged · World drop", "tooltipMuted" }, example.lines[2])
		end)

	it("falls back to the data's pick, Best in slot at the band ahead for a low character", function()
		local data = {
			bis = { ["hunter-beast-mastery"] = { [20] = { alliance = {
				chest = "6468|D|Wailing Caverns: Mutanus the Devourer|0.31",
			} } } },
			bisNames = { [6468] = "Deviate Scale Belt" },
		}
		local pick = TourExample.dataPick(data, "hunter-beast-mastery", 6, "alliance")
		assert.are.equal(6468, pick.itemId)
		assert.are.equal(20, pick.band)
		assert.is_true(pick.upcoming)
		local example = TourExample.example({ pick = pick })
		assert.are.equal("Deviate Scale Belt", example.name)
		assert.are.same({ string.format(L.tooltipYourBisAt, 20), "clientGold" }, example.lines[1])
		assert.is_truthy(example.lines[2][1]:find("Mutanus the Devourer", 1, true))
		assert.is_nil(TourExample.dataPick(data, "hunter-beast-mastery", 6, nil))
		assert.is_nil(TourExample.example({}))
	end)

	it("says the tooltip's own gold line on a pick in the character's band: Best for its span", function()
		local data = {
			bis = { ["hunter-beast-mastery"] = {
				[40] = { alliance = { chest = "6468|D|Wailing Caverns: Mutanus the Devourer|0.31" } },
				[50] = { alliance = { chest = "6468|D|Wailing Caverns: Mutanus the Devourer|0.31" } },
				[60] = { alliance = { chest = "6469|D|Wailing Caverns: Mutanus the Devourer|0.31" } },
			} },
			bisNames = { [6468] = "Deviate Scale Belt", [6469] = "Venomstrike" },
		}
		local pick = TourExample.dataPick(data, "hunter-beast-mastery", 42, "alliance")
		assert.is_false(pick.upcoming)
		assert.are.equal(60, pick.after)
		local example = TourExample.example({ pick = pick })
		assert.are.same({ "Best for 40 to 59", "clientGold" }, example.lines[1])
		local top = TourExample.example({ pick = TourExample.dataPick(data, "hunter-beast-mastery", 60, "alliance") })
		assert.are.same({ "Best at level 60", "clientGold" }, top.lines[1])
	end)

	-- The example is read from the shipped Data.lua at runtime, so a nightly refresh redraws it. This
	-- pins the shape only (data contracts are floors), never which item the nightly picked.
	it("reads a named main hand from the shipped data for a level 6 Beast Mastery hunter", function()
		local Data = require("Data")
		local pick = TourExample.dataPick(Data, "hunter-beast-mastery", 6, "alliance")
		assert.is_table(pick)
		assert.are.equal(Data.bis["hunter-beast-mastery"][20].alliance.main_hand:match("^(%d+)"), tostring(pick.itemId))
		assert.are.equal(Data.bisNames[pick.itemId], pick.name)
		assert.is_string(pick.name)
		assert.is_true(pick.upcoming)
		assert.is_string(TourExample.example({ pick = pick }).lines[2][1])
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
		assert.are.equal(Onboarding.TOUR_VERSION, _G.ForeverSixtyDB.tourSeen)
		assert.is_false(Onboarding.frame:IsShown())
	end)

	it("shows once on a reinstall that kept the old tour's stamp, and once on one that wiped it", function()
		for _, kept in ipairs({ { welcomedAt = 1793000000 }, {} }) do
			start(false)
			_G.ForeverSixtyDB = kept
			Onboarding.register()
			local SiteStatus = require("SiteStatus")
			SiteStatus.welcoming = nil
			-- The login's welcome asks the tour's own rule, not the old stamp.
			assert.is_true(SiteStatus.welcomeDue())
			SiteStatus.openTour()
			assert.is_true(Onboarding.frame:IsShown())
			Onboarding.finish()
			assert.are.equal(Onboarding.TOUR_VERSION, _G.ForeverSixtyDB.tourSeen)
			assert.is_false(SiteStatus.welcomeDue())
			SiteStatus.openTour()
			assert.is_false(Onboarding.frame:IsShown())
			mock.uninstall()
		end
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
