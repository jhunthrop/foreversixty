-- addon/tests/site_status_spec.lua
-- X.3: /fs diag says what the site has of this character and what the inbox holds; the one chat
-- line at login, only when true, five seconds in; the chip's hover when a build waits.
-- The 1.60 taint rules: no frame is written (the idle chip's click opens the dialog the player
-- asked for), and nothing here runs on OnUpdate.
local helper = require("spec_helper")
local mock = require("wow_mock")

local DATA = {
	build = "1.60.1.70291",
	classes = { warrior = { tabs = {
		{ name = "Arms", talents = {} }, { name = "Fury", talents = {} }, { name = "Protection", talents = {} },
	} } },
	weights = {},
}

local KEY = "us/Ashbringer/Bob"
local ARMS = "FSB1:1.60.1.70291:warrior:111:"
local FURY = "FSB1:1.60.1.70291:warrior:211:"
local GEAR = "FS1:1.60.1.70450:warrior:human:0/0/0:head=250528|level=21"
local NO_GEAR = "FS1:1.60.1.70450:warrior:human:0/0/0:|level=21"
local HEAD = "|Hitem:12640|h"

describe("what the site has and what the inbox holds", function()
	local state, Civil, NOW, L, Theme, Options, SiteStatus, Chip, BuildDialog, savedTime

	--- RFC 3339 UTC for `seconds` ago.
	local function ago(seconds)
		local at = Civil.utcFields(NOW - seconds)
		return string.format("%04d-%02d-%02dT%02d:%02d:%02dZ", at.year, at.month, at.day, at.hour, at.min, at.sec)
	end

	local function ack(seconds, character)
		return { type = "ack", character = character or KEY, export_seen_at = ago(seconds),
			requested_at = ago(seconds), sim_request_state = "done" }
	end

	local function build(name, code, character)
		return { type = "build", character = character or KEY, name = name, code = code }
	end

	local function inbox(...)
		_G.ForeverSixtyInbox = { format = 2, messages = { ... } }
	end

	local function record(export, extra)
		_G.ForeverSixtyDB = { characters = { [KEY] = {
			name = "Bob", class = "Warrior", realm = "Ashbringer", region = "us", export = export,
			builtAt = "2026-10-19 11:30", gearAt = extra and extra.gearAt,
		} } }
	end

	--- The account was welcomed at an earlier login unless a spec says otherwise (X.7: the welcome
	--- is its own describe block below).
	local welcomed = true

	local function fire(event, ...)
		if event == "PLAYER_ENTERING_WORLD" and welcomed then
			_G.ForeverSixtyDB = _G.ForeverSixtyDB or {}
			_G.ForeverSixtyDB.welcomedAt = _G.ForeverSixtyDB.welcomedAt or 1
		end
		Options.onEvent(Options.frame, event, ...)
	end

	local function timersAt(seconds)
		local count = 0
		for _, timer in ipairs(state.timers) do
			if timer.seconds == seconds then
				count = count + 1
			end
		end
		return count
	end

	before_each(function()
		state = mock.install({
			class = { name = "Warrior", token = "WARRIOR" }, race = { name = "Human", token = "Human" },
			realm = "Ashbringer", region = 1, playerName = "Bob",
			talents = { { name = "Arms", talents = {} }, { name = "Fury", talents = {} },
				{ name = "Protection", talents = {} } },
			itemStats = { [HEAD] = { __itemId = 12640, __slot = "INVTYPE_HEAD" } },
		})
		savedTime = _G.time
		Theme = require("Theme")
		Theme.reset()
		L = helper.load("Locale")
		_G.ForeverSixtyDB, _G.ForeverSixtyInbox = nil, nil
		helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Cards")
		helper.load("Export")
		helper.load("Gear")
		helper.load("ItemTooltip")
		helper.load("Follow")
		helper.load("TalentGlow")
		Chip = helper.load("Chip")
		helper.load("TalentMarks")
		helper.load("Minimap")
		helper.load("Comms")
		BuildDialog = helper.load("BuildDialog")
		Civil = BuildDialog
		NOW = Civil.civilSeconds(2026, 10, 19, 12, 0, 0)
		_G.time = function(fields)
			if fields ~= nil then
				return Civil.civilSeconds(fields.year, fields.month, fields.day, fields.hour, fields.min,
					fields.sec or 0)
			end
			return NOW
		end
		helper.load("Settings")
		helper.load("RewardMarks")
		helper.load("PanelTab")
		helper.load("GuildTab")
		helper.load("Inspect")
		helper.load("GroupMembers")
		helper.load("ReadyCheck")
		SiteStatus = helper.load("SiteStatus")
		Options = helper.load("Options")
		Options.data = DATA
		BuildDialog.data = DATA
		Options.register()
		-- Equipment that reads nothing: the export's own refresh leaves the record as it is.
		state.equipped = {}
		state.printed = {}
	end)

	after_each(function()
		_G.time = savedTime
		_G.CHAT_TIMESTAMP_FORMAT = nil
		_G.ForeverSixtyDB, _G.ForeverSixtyInbox = nil, nil
		mock.uninstall()
	end)

	describe("the Site line", function()
		it("says the site has not seen the character when no ack exists", function()
			inbox()
			assert.are.equal(L.diagSiteNever, SiteStatus.siteLine(SiteStatus.facts()))
			assert.are.equal("Site: not seen yet. Is the companion running and signed in?", L.diagSiteNever)
		end)

		it("says when the newest ack was and how many characters the site has", function()
			inbox(ack(2 * 3600), ack(3600, "US/Ashbringer/Alice"), ack(5 * 3600, "US/Ashbringer/Cleo"))
			assert.are.equal("Site: has your save from 2 h ago (3 characters)", SiteStatus.siteLine(SiteStatus.facts()))
		end)

		it("reads the site's ruleset key as this beta character", function()
			_G.ForeverSixtyDB = { characters = { ["us/Classic Beta PvP/Bow Jackzon"] = {
				name = "Bow Jackzon", class = "Warrior", realm = "Classic Beta PvP", region = "us" } } }
			state.realm, state.playerName = "Classic Beta PvP", "Bow Jackzon"
			inbox(ack(2 * 3600, "us/pvp/bow-jackzon"))
			assert.are.equal("Site: has your save from 2 h ago (1 character)", SiteStatus.siteLine(SiteStatus.facts()))
		end)

		it("says one character in the singular", function()
			inbox(ack(2 * 3600))
			assert.are.equal("Site: has your save from 2 h ago (1 character)", SiteStatus.siteLine(SiteStatus.facts()))
		end)

		it("says the day and the days gone once the newest ack is older than 24 hours", function()
			inbox(ack(3 * 86400))
			assert.are.equal("Site: last seen 16 Oct, 3 days ago", SiteStatus.siteLine(SiteStatus.facts()))
			inbox(ack(86400 + 60))
			assert.are.equal("Site: last seen 18 Oct, 1 day ago", SiteStatus.siteLine(SiteStatus.facts()))
		end)

		it("reads the newest ack for this character, not another's", function()
			inbox(ack(10 * 3600), ack(3600, "US/Ashbringer/Alice"))
			assert.are.equal("Site: has your save from 10 h ago (2 characters)", SiteStatus.siteLine(SiteStatus.facts()))
		end)

		-- A.2: an age, never a time of day, so the chat's own clock (12 or 24 hour) cannot disagree with it.
		it("says the age of the newest ack, the same whichever clock the chat shows (A.2)", function()
			inbox(ack(2 * 3600), ack(30 * 60))
			_G.CHAT_TIMESTAMP_FORMAT = nil
			assert.are.equal("Site: has your save from 30 min ago (1 character)", SiteStatus.siteLine(SiteStatus.facts()))
			_G.CHAT_TIMESTAMP_FORMAT = "%I:%M %p "
			assert.are.equal("Site: has your save from 30 min ago (1 character)", SiteStatus.siteLine(SiteStatus.facts()))
			inbox(ack(0))
			assert.are.equal("Site: has your save from just now (1 character)", SiteStatus.siteLine(SiteStatus.facts()))
			inbox(ack(95 * 60))
			assert.are.equal("Site: has your save from 1 h 35 min ago (1 character)", SiteStatus.siteLine(SiteStatus.facts()))
			NOW = Civil.civilSeconds(2026, 10, 19, 18, 25, 0)
			inbox(ack(5 * 60))
			assert.are.equal("Site: has your save from 5 min ago (1 character)", SiteStatus.siteLine(SiteStatus.facts()))
			_G.CHAT_TIMESTAMP_FORMAT = nil
			assert.are.equal("Site: has your save from 5 min ago (1 character)", SiteStatus.siteLine(SiteStatus.facts()))
		end)

		it("prints the Export line's build time and the login time in the same clock (X.5)", function()
			record(GEAR)
			_G.CHAT_TIMESTAMP_FORMAT = "%I:%M %p "
			assert.is_truthy(require("Export").diagLine():find("built 11:30 am,", 1, true))
			_G.CHAT_TIMESTAMP_FORMAT = nil
			assert.is_truthy(require("Export").diagLine():find("built 11:30,", 1, true))
			_G.CHAT_TIMESTAMP_FORMAT = "%I:%M %p "
			inbox()
			fire("PLAYER_ENTERING_WORLD", true, false)
			-- The site's clock drops ":00" on the hour (Compat.clockText, 2026-11-13).
			assert.is_truthy(SiteStatus.inboxLine(DATA):find("as of your login at 12 pm.", 1, true))
		end)
	end)

	describe("the Inbox line", function()
		it("is empty with nothing for this character, as of the login, and says /reload does not read it (X.5)", function()
			inbox(ack(3600), build("Fury", FURY, "US/Ashbringer/Alice"))
			-- No login on record yet (a first run, or an old saved variable): no time to name.
			assert.are.equal("Inbox: empty. Builds arrive when you log in, not on /reload.", SiteStatus.inboxLine(DATA))
			fire("PLAYER_ENTERING_WORLD", true, false)
			assert.are.equal("Inbox: empty, as of your login at 12:00. Builds arrive when you log in, not on /reload.",
				SiteStatus.inboxLine(DATA))
			_G.ForeverSixtyInbox = nil
			assert.are.equal("Inbox: empty, as of your login at 12:00. Builds arrive when you log in, not on /reload.",
				SiteStatus.inboxLine(DATA))
			-- A /reload keeps the login's time, and a later login moves it.
			NOW = NOW + 3600
			fire("PLAYER_ENTERING_WORLD", false, true)
			assert.are.equal("Inbox: empty, as of your login at 12:00. Builds arrive when you log in, not on /reload.",
				SiteStatus.inboxLine(DATA))
			fire("PLAYER_ENTERING_WORLD", true, false)
			assert.are.equal("Inbox: empty, as of your login at 13:00. Builds arrive when you log in, not on /reload.",
				SiteStatus.inboxLine(DATA))
		end)

		it("says the day of a login that was not today, in the chat's clock", function()
			fire("PLAYER_ENTERING_WORLD", true, false)
			NOW = NOW + 86400
			assert.is_truthy(SiteStatus.inboxLine(DATA):find("as of your login at yesterday 12:00.", 1, true))
		end)

		it("counts a build waiting, named, and the messages by kind (X.5)", function()
			inbox(ack(3600), build("Human Warrior", ARMS))
			assert.are.equal("Inbox: 1 build waiting (Human Warrior)", SiteStatus.inboxLine(DATA))
			inbox(build("Human Warrior", ARMS), { type = "guild", character = KEY, text = "Raid at 8" })
			assert.are.equal("Inbox: 1 build waiting (Human Warrior) \194\183 1 guild message", SiteStatus.inboxLine(DATA))
			inbox(build("Human Warrior", ARMS), { type = "guild", character = KEY, text = "Raid at 8" },
				{ type = "guild", text = "Bring flasks" })
			assert.are.equal("Inbox: 1 build waiting (Human Warrior) \194\183 2 guild messages",
				SiteStatus.inboxLine(DATA))
			-- Kinds that mix are only "messages".
			inbox(build("Human Warrior", ARMS), { type = "guild", character = KEY, text = "Raid at 8" },
				{ type = "last_raid", character = KEY, date = "2026-10-18", bosses = {} })
			assert.are.equal("Inbox: 1 build waiting (Human Warrior) \194\183 2 messages", SiteStatus.inboxLine(DATA))
			inbox({ type = "last_raid", character = KEY, date = "2026-10-18", bosses = {} })
			assert.are.equal("Inbox: no builds waiting \194\183 1 raid recap", SiteStatus.inboxLine(DATA))
		end)

		it("names several builds", function()
			inbox(build("Arms raid", ARMS), build("Fury leveling", FURY))
			assert.are.equal("Inbox: 2 builds waiting (Fury leveling, Arms raid)", SiteStatus.inboxLine(DATA))
			inbox(build("Arms raid", ARMS), build("Fury leveling", FURY), { type = "guild", character = KEY, text = "x" })
			assert.are.equal("Inbox: 2 builds waiting (Fury leveling, Arms raid) \194\183 1 guild message",
				SiteStatus.inboxLine(DATA))
		end)

		it("does not count a dismissed or used build", function()
			inbox(build("Arms raid", ARMS))
			require("Follow").dismissInbox(ARMS)
			assert.are.equal("Inbox: empty. Builds arrive when you log in, not on /reload.", SiteStatus.inboxLine(DATA))
		end)

		it("does not count a build for a same-named character on another realm that names its realm (X.5)", function()
			inbox(build("Theirs", ARMS, "US/Whitemane/Bob"), build("Mine", FURY, "us/normal/bob"))
			assert.are.equal("Inbox: 1 build waiting (Mine)", SiteStatus.inboxLine(DATA))
		end)
	end)

	describe("/fs diag", function()
		local function lines()
			return Options.handle("diag")
		end

		local function at(list, text)
			for index, line in ipairs(list) do
				if line == text then
					return index
				end
			end
		end

		it("puts the Site and Inbox lines right after the Export line", function()
			record(GEAR)
			inbox(ack(3600), build("Arms raid", ARMS))
			local list = lines()
			local export = at(list, require("Export").diagLine())
			assert.is_truthy(export)
			assert.are.equal("Site: has your save from 1 h ago (1 character)", list[export + 1])
			assert.are.equal("Inbox: 1 build waiting (Arms raid)", list[export + 2])
		end)

		it("says everything is working only with gear and a site that saw the character within a day", function()
			record(GEAR)
			inbox(ack(3600))
			assert.are.equal(L.diagAllWorking, lines()[1])
		end)

		it("does not say it with no ack", function()
			record(GEAR)
			inbox()
			local list = lines()
			assert.are.equal("Addon running; the site has not seen this character yet.", list[1])
			assert.is_nil(at(list, L.diagAllWorking))
		end)

		it("does not say it when the export has no gear", function()
			record(NO_GEAR)
			inbox(ack(3600))
			assert.are.equal(L.diagNotSeenHeadline, lines()[1])
			_G.ForeverSixtyDB = nil
			assert.are.equal(L.diagNotSeenHeadline, lines()[1])
		end)

		it("does not say it when the newest ack is older than a day", function()
			record(GEAR)
			inbox(ack(25 * 3600))
			local list = lines()
			assert.are.equal(L.diagNotSeenHeadline, list[1])
			assert.is_truthy(at(list, "Site: last seen 18 Oct, 1 day ago"))
		end)
	end)

	describe("the login line", function()
		--- A login from the character screen (the event's isInitialLogin), and the line's timer run.
		local function login()
			fire("PLAYER_ENTERING_WORLD", true, false)
			mock.runTimers(state)
		end

		--- A /reload (isReloadingUi), and a loading screen in the world (neither).
		local function reload()
			fire("PLAYER_ENTERING_WORLD", false, true)
			mock.runTimers(state)
		end

		local function zone()
			fire("PLAYER_ENTERING_WORLD", false, false)
			mock.runTimers(state)
		end

		local function only(text)
			assert.are.same({ Theme.chatLine(text) }, state.printed)
		end

		it("says nothing at login without builds waiting: the site and the companion are the Me tab's footer", function()
			record(GEAR, { gearAt = "2026-10-19 11:00" })
			inbox()
			login()
			assert.are.same({}, state.printed)
			inbox(ack(2 * 3600))
			login()
			assert.are.same({}, state.printed)
			inbox(ack(5 * 86400))
			login()
			assert.are.same({}, state.printed)
		end)

		it("says a build is waiting, by name and for whom (X.5)", function()
			record(GEAR)
			inbox(ack(3600), build("Human Warrior", ARMS))
			login()
			only("a build for Bob is waiting: Human Warrior. /fs opens it.")
			assert.are.equal("a build for Bob is waiting: Human Warrior. /fs opens it.",
				string.format(L.siteLoginBuild, "Bob", "Human Warrior"))
		end)

		it("says several builds wait in the one line", function()
			record(GEAR)
			inbox(ack(3600), build("Arms raid", ARMS), build("Fury leveling", FURY))
			login()
			only("2 builds for Bob are waiting: Fury leveling, Arms raid. /fs opens them.")
		end)

		it("says a waiting build even when the site has never seen the character", function()
			record(GEAR)
			inbox(build("Arms raid", ARMS))
			login()
			only(string.format(L.siteLoginBuild, "Bob", "Arms raid"))
		end)

		it("waits five seconds, after the export's refresh", function()
			record(GEAR)
			inbox(ack(3600), build("Arms raid", ARMS))
			fire("PLAYER_ENTERING_WORLD", true, false)
			assert.are.same({}, state.printed)
			assert.are.equal(2, timersAt(5))
			mock.runTimers(state)
			only(string.format(L.siteLoginBuild, "Bob", "Arms raid"))
		end)

		-- X.5: once per login. A login says its line; a /reload says it only when it changed; a
		-- loading screen never says it.
		it("says its line at every login, never after a /reload or a loading screen for the same inbox", function()
			record(GEAR)
			inbox(ack(3600), build("Arms raid", ARMS))
			login()
			only(string.format(L.siteLoginBuild, "Bob", "Arms raid"))
			state.printed = {}
			reload()
			zone()
			reload()
			assert.are.same({}, state.printed)
			-- The next login is a new reading of the inbox: the build still waits, and is said again.
			login()
			only(string.format(L.siteLoginBuild, "Bob", "Arms raid"))
		end)

		it("says a /reload's line when it is not the one said last, and then not again", function()
			record(GEAR)
			inbox(ack(3600), build("Arms raid", ARMS))
			login()
			state.printed = {}
			inbox(ack(3600), build("Arms raid", ARMS), build("Fury leveling", FURY))
			reload()
			only(string.format(L.siteLoginBuilds, 2, "Bob", "Fury leveling, Arms raid"))
			state.printed = {}
			reload()
			assert.are.same({}, state.printed)
		end)

		it("remembers the line across a /reload in the saved variable, per character", function()
			record(GEAR)
			inbox(ack(3600), build("Arms raid", ARMS))
			login()
			assert.are.equal(string.format(L.siteLoginBuild, "Bob", "Arms raid"),
				_G.ForeverSixtyDB.site[KEY].said)
			assert.is_number(_G.ForeverSixtyDB.site[KEY].loginAt)
			-- A new Lua state after the /reload: the module's own memory is gone, the saved variable stays.
			SiteStatus.announced = nil
			state.printed = {}
			reload()
			assert.are.same({}, state.printed)
		end)

		it("never says it for pressing Use this build (the line is for entering the world only)", function()
			record(GEAR)
			inbox(ack(3600), build("Arms raid", ARMS))
			login()
			state.printed = {}
			BuildDialog.data = DATA
			BuildDialog.open("paste")
			mock.runTimers(state)
			BuildDialog.useCurrent()
			BuildDialog.choose(BuildDialog.rows[1])
			BuildDialog.useCurrent()
			mock.runTimers(state)
			for _, line in ipairs(state.printed) do
				assert.is_nil(line:find("is waiting", 1, true), line)
			end
		end)

		it("treats a client that passes neither flag as a login the first time, a /reload after", function()
			record(GEAR)
			inbox(ack(3600), build("Arms raid", ARMS))
			fire("PLAYER_ENTERING_WORLD")
			mock.runTimers(state)
			only(string.format(L.siteLoginBuild, "Bob", "Arms raid"))
			state.printed = {}
			fire("PLAYER_ENTERING_WORLD")
			mock.runTimers(state)
			assert.are.same({}, state.printed)
		end)

		it("reads the ack's receive time and ignores the export's own change time", function()
			record(GEAR, { gearAt = "2026-10-19 11:00" })
			local acked = ack(10 * 60)
			acked.export_changed_at = ago(5 * 3600)
			inbox(acked)
			assert.are.equal(NOW - 600, SiteStatus.facts().seen)
			assert.are.equal("Site: has your save from 10 min ago (1 character)", SiteStatus.siteLine(SiteStatus.facts()))
		end)
	end)

	describe("the chip", function()
		local function idle(extra)
			local input = {
				level = 30, build = nil, progress = { matched = 0, total = 0 }, unspent = 0,
				classIcon = "crest-warrior", now = 100, clock = NOW, shown = true, combatHide = true,
				inboxWaiting = 1,
			}
			for key, value in pairs(extra or {}) do
				input[key] = value
			end
			return input
		end

		it("says a build is waiting in its hover, and a click opens it, when none is followed", function()
			local model = Chip.model(idle())
			assert.are.equal("idle", model.chosen.kind)
			assert.are.equal("inbox", model.chosen.action)
			local joined = table.concat(Chip.tooltipLines(model), "\n")
			assert.is_truthy(joined:find(L.chipHintBuildWindow, 1, true))
			-- The long form leads the hover when the chip had to shorten its line.
			assert.is_truthy(table.concat(Chip.tooltipLines(model, true), "\n")
				:find("A build from the site is waiting. Click to open.", 1, true))
			assert.are.equal("A build from the site is waiting. Click to open.", L.chipInboxWaiting)
		end)

		it("says nothing about the inbox with nothing waiting, or with a build followed", function()
			local none = table.concat(Chip.tooltipLines(Chip.model(idle({ inboxWaiting = 0 }))), "\n")
			assert.is_nil(none:find("waiting", 1, true))
			local followed = idle({ build = { name = "Arms leveling", order = { {}, {}, {} } } })
			assert.is_falsy(Chip.model(followed).inboxWaiting)
			assert.is_nil(table.concat(Chip.tooltipLines(Chip.model(followed)), "\n"):find("waiting", 1, true))
		end)

		it("keeps another state's click when the chip has something else to say", function()
			local model = Chip.model(idle({ lastRaid = { date = "2026-10-18", bosses = {
				{ name = "Magmadar", percentile = 88, biggest_loss = "downtime" } } } }))
			assert.are_not.equal("inbox", model.chosen.action)
			assert.is_falsy(model.inboxWaiting)
		end)

		it("counts the builds waiting for this character, for the chip's input", function()
			inbox(build("Arms raid", ARMS), build("Fury", FURY, "US/Ashbringer/Alice"))
			assert.are.equal(1, Chip.inboxWaiting(DATA))
			_G.ForeverSixtyInbox = nil
			assert.are.equal(0, Chip.inboxWaiting(DATA))
		end)

		it("opens the build dialog on the idle chip's click", function()
			inbox(build("Arms raid", ARMS))
			local handler = Chip.handlerFor({ action = "inbox" })
			assert.is_function(handler)
			handler()
			assert.is_true(BuildDialog.frame:IsShown())
		end)
	end)

	-- design/specs/2026-11-08-addon-ux.md 3.1: the first login opens the tour (no chat lines), once per
	-- account, five seconds in; the tour stamps the account when the player ends it.
	describe("the first-login tour", function()
		local opened

		local function login()
			fire("PLAYER_ENTERING_WORLD", true, false)
			mock.runTimers(state)
		end

		before_each(function()
			welcomed = false
			state.printed = {}
			opened = 0
			SiteStatus.openTour = function()
				opened = opened + 1
			end
		end)

		after_each(function()
			welcomed = true
			SiteStatus.openTour = nil
		end)

		it("opens the tour five seconds in and says no welcome in chat", function()
			record(GEAR)
			inbox(ack(3600))
			fire("PLAYER_ENTERING_WORLD", true, false)
			assert.is_true(timersAt(5) >= 1)
			assert.are.equal(0, opened)
			mock.runTimers(state)
			assert.are.equal(1, opened)
			assert.are.same({}, state.printed)
			assert.is_nil(L.welcomeOn)
			assert.is_nil(SiteStatus.welcome)
		end)

		it("opens it once when the world is entered twice before the five seconds are up", function()
			record(GEAR)
			inbox(ack(3600))
			fire("PLAYER_ENTERING_WORLD", true, false)
			fire("PLAYER_ENTERING_WORLD", true, false)
			mock.runTimers(state)
			assert.are.equal(1, opened)
		end)

		it("does not open it at a reload or a loading screen, but the inbox still speaks", function()
			record(GEAR)
			inbox(ack(3600), build("Arms raid", ARMS))
			_G.ForeverSixtyDB.welcomedAt = 1793000000
			login()
			assert.are.equal(0, opened)
			assert.are.same({ Theme.chatLine(string.format(L.siteLoginBuild, "Bob", "Arms raid")) }, state.printed)
		end)

		it("is one stamp for the account: another character that finds it set sees nothing", function()
			record(GEAR)
			inbox(ack(3600))
			_G.ForeverSixtyDB.welcomedAt = 1793000000
			login()
			assert.are.equal(0, opened)
		end)

		it("runs no OnUpdate", function()
			record(GEAR)
			login()
			for _, frame in ipairs(state.frames) do
				assert.is_nil(frame:GetScript("OnUpdate"))
			end
		end)
	end)

	describe("the record", function()
		it("stamps when it first carried gear and keeps that while it does", function()
			local Export = require("Export")
			state.date = "2026-10-19 10:00"
			Export.refresh(DATA)
			assert.is_nil(_G.ForeverSixtyDB.characters[KEY].gearAt)
			state.equipped = { [1] = HEAD }
			state.date = "2026-10-19 10:05"
			Export.refresh(DATA)
			assert.are.equal("2026-10-19 10:05", _G.ForeverSixtyDB.characters[KEY].gearAt)
			state.date = "2026-10-19 10:30"
			Export.refresh(DATA)
			assert.are.equal("2026-10-19 10:05", _G.ForeverSixtyDB.characters[KEY].gearAt)
			assert.are.equal("2026-10-19 10:30", _G.ForeverSixtyDB.characters[KEY].builtAt)
		end)
	end)
end)
