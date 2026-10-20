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

	local function fire(event)
		Options.onEvent(Options.frame, event)
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
		helper.load("GroupPanel")
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
		_G.ForeverSixtyDB, _G.ForeverSixtyInbox = nil, nil
		mock.uninstall()
	end)

	describe("the Site line", function()
		it("says the site has not seen the character when no ack exists", function()
			inbox()
			assert.are.equal(L.diagSiteNever, SiteStatus.siteLine(SiteStatus.facts()))
			assert.are.equal("Site: not seen yet \226\128\148 is the companion running and signed in?", L.diagSiteNever)
		end)

		it("says when the newest ack was and how many characters the site has", function()
			inbox(ack(2 * 3600), ack(3600, "US/Ashbringer/Alice"), ack(5 * 3600, "US/Ashbringer/Cleo"))
			assert.are.equal("Site: seen 10:00 (3 characters)", SiteStatus.siteLine(SiteStatus.facts()))
		end)

		it("reads the site's ruleset key as this beta character", function()
			_G.ForeverSixtyDB = { characters = { ["us/Classic Beta PvP/Bow Jackzon"] = {
				name = "Bow Jackzon", class = "Warrior", realm = "Classic Beta PvP", region = "us" } } }
			state.realm, state.playerName = "Classic Beta PvP", "Bow Jackzon"
			inbox(ack(2 * 3600, "us/pvp/bow-jackzon"))
			assert.are.equal("Site: seen 10:00 (1 character)", SiteStatus.siteLine(SiteStatus.facts()))
		end)

		it("says one character in the singular", function()
			inbox(ack(2 * 3600))
			assert.are.equal("Site: seen 10:00 (1 character)", SiteStatus.siteLine(SiteStatus.facts()))
		end)

		it("says the day and the days gone once the newest ack is older than 24 hours", function()
			inbox(ack(3 * 86400))
			assert.are.equal("Site: last seen 16 Oct, 3 days ago", SiteStatus.siteLine(SiteStatus.facts()))
			inbox(ack(86400 + 60))
			assert.are.equal("Site: last seen 18 Oct, 1 day ago", SiteStatus.siteLine(SiteStatus.facts()))
		end)

		it("reads the newest ack for this character, not another's", function()
			inbox(ack(10 * 3600), ack(3600, "US/Ashbringer/Alice"))
			assert.are.equal("Site: seen 02:00 (2 characters)", SiteStatus.siteLine(SiteStatus.facts()))
		end)
	end)

	describe("the Inbox line", function()
		it("is empty with nothing for this character", function()
			inbox(ack(3600), build("Fury", FURY, "US/Ashbringer/Alice"))
			assert.are.equal("Inbox: empty", SiteStatus.inboxLine(DATA))
			_G.ForeverSixtyInbox = nil
			assert.are.equal("Inbox: empty", SiteStatus.inboxLine(DATA))
		end)

		it("counts a build waiting, named, and the messages", function()
			inbox(ack(3600), build("Human Warrior", ARMS))
			assert.are.equal("Inbox: 1 build waiting (Human Warrior) \194\183 0 messages", SiteStatus.inboxLine(DATA))
			inbox(build("Human Warrior", ARMS), { type = "last_raid", character = KEY, date = "2026-10-18", bosses = {} })
			assert.are.equal("Inbox: 1 build waiting (Human Warrior) \194\183 1 message", SiteStatus.inboxLine(DATA))
		end)

		it("names several builds and says messages alone without builds", function()
			inbox(build("Arms raid", ARMS), build("Fury leveling", FURY))
			assert.are.equal("Inbox: 2 builds waiting (Fury leveling, Arms raid) \194\183 0 messages",
				SiteStatus.inboxLine(DATA))
			inbox({ type = "last_raid", character = KEY, date = "2026-10-18", bosses = {} })
			assert.are.equal("Inbox: no builds waiting \194\183 1 message", SiteStatus.inboxLine(DATA))
		end)

		it("does not count a dismissed or used build", function()
			inbox(build("Arms raid", ARMS))
			require("Follow").dismissInbox(ARMS)
			assert.are.equal("Inbox: empty", SiteStatus.inboxLine(DATA))
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
			assert.are.equal("Site: seen 11:00 (1 character)", list[export + 1])
			assert.are.equal("Inbox: 1 build waiting (Arms raid) \194\183 0 messages", list[export + 2])
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
		local function login()
			fire("PLAYER_ENTERING_WORLD")
			mock.runTimers(state)
		end

		local function only(text)
			assert.are.same({ Theme.chatLine(text) }, state.printed)
		end

		it("says the site has not seen the character, when it never has", function()
			record(GEAR)
			inbox()
			login()
			only(L.siteLoginNever)
			assert.are.equal("the site has not seen this character yet. Keep the companion running and signed in; "
				.. "log out once from the world.", L.siteLoginNever)
		end)

		it("says the gear is on its way when the site last saw the character before it had gear", function()
			record(GEAR, { gearAt = "2026-10-19 11:00" })
			inbox(ack(2 * 3600), build("Arms raid", ARMS))
			login()
			only(L.siteLoginGear)
			assert.are.equal("your gear is on its way to the site; check it after your next logout.", L.siteLoginGear)
		end)

		it("says nothing of the gear once the site has seen the record that has it", function()
			record(GEAR, { gearAt = "2026-10-19 09:00" })
			inbox(ack(2 * 3600))
			login()
			assert.are.same({}, state.printed)
		end)

		it("says a build is waiting, by name", function()
			record(GEAR)
			inbox(ack(3600), build("Human Warrior", ARMS))
			login()
			only(string.format(L.siteLoginBuild, "Human Warrior"))
			assert.are.equal("a build from the site is waiting: Human Warrior. /fs opens it.",
				string.format(L.siteLoginBuild, "Human Warrior"))
		end)

		it("says several builds wait in the one line", function()
			record(GEAR)
			inbox(ack(3600), build("Arms raid", ARMS), build("Fury leveling", FURY))
			login()
			only(string.format(L.siteLoginBuilds, 2, "Fury leveling, Arms raid"))
		end)

		it("says the companion has gone quiet after three days, with the day it last picked up", function()
			record(GEAR)
			inbox(ack(5 * 86400))
			login()
			only(string.format(L.siteLoginSilent, "14 Oct"))
			assert.are.equal("the companion has not picked up your data since 14 Oct. Start it, or paste an export at "
				.. "foreversixty.gg/setup.", string.format(L.siteLoginSilent, "14 Oct"))
		end)

		it("says nothing at two days, or when all is well", function()
			record(GEAR)
			inbox(ack(2 * 86400))
			login()
			assert.are.same({}, state.printed)
		end)

		it("is the first that applies, never two lines", function()
			-- Never seen and a build waiting: only the first.
			record(GEAR)
			inbox(build("Arms raid", ARMS))
			login()
			only(L.siteLoginNever)
			-- Silent for days and a build waiting: the build.
			state.printed = {}
			SiteStatus.announced = nil
			inbox(ack(6 * 86400), build("Arms raid", ARMS))
			login()
			only(string.format(L.siteLoginBuild, "Arms raid"))
		end)

		it("waits five seconds, after the export's refresh, and says it once a session", function()
			record(GEAR)
			inbox()
			fire("PLAYER_ENTERING_WORLD")
			assert.are.same({}, state.printed)
			assert.are.equal(2, timersAt(5))
			fire("PLAYER_ENTERING_WORLD")
			assert.are.equal(3, timersAt(5))
			mock.runTimers(state)
			only(L.siteLoginNever)
			fire("PLAYER_ENTERING_WORLD")
			mock.runTimers(state)
			only(L.siteLoginNever)
		end)

		it("reads the record after the refresh has run", function()
			-- The record gains gear in the world's first build, before the line is worked out.
			record(NO_GEAR, { gearAt = nil })
			_G.ForeverSixtyDB.characters[KEY].gearAt = nil
			inbox(ack(2 * 3600))
			state.equipped = { [1] = HEAD }
			state.date = "2026-10-19 11:59"
			login()
			local mine = _G.ForeverSixtyDB.characters[KEY]
			assert.is_true(require("Export").hasGear(mine.export))
			assert.are.equal("2026-10-19 11:59", mine.gearAt)
			only(L.siteLoginGear)
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
			assert.is_truthy(joined:find("A build from the site is waiting. Click to open.", 1, true))
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
