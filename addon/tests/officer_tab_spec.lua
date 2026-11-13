-- addon/tests/officer_tab_spec.lua
-- The Forever Sixty window's Officer tab (design/specs/2026-11-13-addon-hub.md 4.6, 8; lane L7): the officer
-- gate both ways and the master looter without it, the roster (raiders by rank or level 60, by sign-up then
-- Score, a heading click, the filters, the cap), Tonight's readiness from live and inspected members, the
-- site's approvals, Push a build (the guild flyout's push, rebuilt; disabled without a build), the master
-- loot session (gains largest first, Won and Att., Give twice) and the guild's rank names in the export.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local START = "2026-12-10T02:00:00Z"

--- A guild roster entry as Theme.guildRoster reads it.
local function member(name, classToken, rankIndex, level, online)
	return { name = name .. "-ClassicBetaPvP", classToken = classToken, rankIndex = rankIndex, level = level or 60,
		online = online ~= false }
end

local ROSTER = {
	member("Brannoc", "WARRIOR", 1), member("Ysolde", "PRIEST", 3), member("Kethra", "ROGUE", 3),
	member("Dravik", "WARLOCK", 3), member("Velisandre", "HUNTER", 1), member("Orwen", "PRIEST", 3, 60, false),
	member("Pell", "WARLOCK", 3), member("Alt", "MAGE", 5, 24),
}

local function records()
	return {
		["Brannoc-ClassicBetaPvP"] = { score = 1612, tier = "epic", attendance = 1.0, fights = 30,
			best_parses = { Onyxia = 91, Lucifron = 71 } },
		["Ysolde-ClassicBetaPvP"] = { score = 1490, tier = "rare", attendance = 0.96, fights = 30 },
		["Kethra-ClassicBetaPvP"] = { score = 1377, tier = "rare", attendance = 0.92, fights = 20 },
		["Dravik-ClassicBetaPvP"] = { score = 1254, tier = "rare", attendance = 0.84, fights = 20 },
		["Velisandre-ClassicBetaPvP"] = { score = 1184, tier = "rare", attendance = 0.88, fights = 20 },
		["Pell-ClassicBetaPvP"] = { score = 920, tier = "uncommon", attendance = 0.61, fights = 8 },
	}
end

local MEMBERS = {
	{ name = "Brannoc", class = "warrior", answer = "in", role = "tank" },
	{ name = "Ysolde", class = "priest", answer = "in", role = "healer" },
	{ name = "Kethra", class = "rogue", answer = "in", role = "melee" },
	{ name = "Dravik", class = "warlock", answer = "late", role = "ranged" },
	{ name = "Velisandre", class = "hunter", answer = "in", role = "ranged" },
	{ name = "Orwen", class = "priest", answer = "out", role = "healer" },
}

describe("Officer, the Officer tab's model", function()
	local Officer, now

	before_each(function()
		mock.install({ class = { name = "Hunter", token = "HUNTER" }, playerName = "Velisandre",
			globals = { date = os.date } })
		require("Theme").reset()
		helper.load("Export")
		Officer = helper.load("Officer")
		now = 100000
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
		mock.uninstall()
	end)

	--- Officer.rows' input over the roster above; `extra` overrides.
	local function input(extra)
		local recs = records()
		local base = {
			roster = ROSTER, record = function(name)
				return recs[name]
			end, seen = {
				["Kethra-ClassicBetaPvP"] = { spec = "rogue-combat", gearGap = 14, buildName = "Combat raid", time = now - 60,
					readiness = { addonVersion = "2.0.45", attuned = { ["raid:onyxias-lair"] = true }, encounterItems = {},
						supplies = { flask = 2, food = 20, fire_protection = 3 } } },
				["Dravik-ClassicBetaPvP"] = { spec = "warlock-destruction", gearGap = 18.6, buildName = "Destruction raid",
					time = now - 120, readiness = { addonVersion = "2.0.40", attuned = {}, encounterItems = {},
						supplies = { flask = 1, food = 10 } } },
				["Ysolde-ClassicBetaPvP"] = { spec = "priest-holy", gearGap = 0, time = now - 7200 },
			}, inspect = function(_, short)
				if short == "Kethra" then
					return { itemLevel = 62, items = { chest = "item:50:0" }, enchants = { chest = false } }
				end
				return nil
			end, own = { short = "Velisandre", spec = "hunter-beast-mastery", gearGap = 40.4, unit = "dps",
				buildName = "Beast Mastery raid build", readiness = { addonVersion = "2.0.45", attuned = {},
					encounterItems = {}, supplies = { flask = 1, food = 5 } } },
			members = MEMBERS, ranks = { 1, 3 }, now = now, data = {},
		}
		for key, value in pairs(extra or {}) do
			base[key] = value
		end
		return base
	end

	local function shorts(rows)
		local out = {}
		for index, row in ipairs(rows) do
			out[index] = row.short
		end
		return out
	end

	it("lists the raiders by sign-up, then Score, and counts the rest", function()
		local rows, counts = Officer.rows(input())
		assert.are.same({ "Brannoc", "Ysolde", "Kethra", "Velisandre", "Dravik", "Orwen", "Pell" }, shorts(rows))
		assert.are.same({ raiders = 7, total = 8, unsigned = 1, online = 6, signups = true }, counts)
		local brannoc = rows[1]
		assert.are.same({ 1612, "epic", 100, "in" }, { brannoc.score, brannoc.tier, brannoc.attendance, brannoc.answer })
		assert.are.same({ boss = "Onyxia", percentile = 91 }, brannoc.bestParse)
		assert.are.equal("none", rows[7].answer)
	end)

	it("counts every level 60 a raider while the officers have named no ranks", function()
		local rows = Officer.rows(input({ ranks = nil }))
		assert.are.equal(7, #rows)
		assert.is_nil(({ Officer.rows(input({ ranks = {} })) })[1][8])
		local all = Officer.rows(input({ filter = "all" }))
		assert.are.equal(8, #all)
	end)

	it("filters to the raiders not signed up and to those online", function()
		assert.are.same({ "Pell" }, shorts(Officer.rows(input({ filter = "unsigned" }))))
		local online = shorts(Officer.rows(input({ filter = "online" })))
		assert.are.equal(6, #online)
		assert.is_nil(table.concat(online, ","):find("Orwen", 1, true))
	end)

	it("sorts by a heading, numbers high to low first, a second click turning it round, blanks last", function()
		local sort = Officer.nextSort(nil, "gap")
		assert.are.same({ key = "gap", descending = true }, sort)
		local rows = Officer.rows(input({ sort = sort }))
		assert.are.same({ "Velisandre", "Dravik", "Kethra", "Ysolde" }, { rows[1].short, rows[2].short, rows[3].short,
			rows[4].short })
		sort = Officer.nextSort(sort, "gap")
		assert.is_false(sort.descending)
		rows = Officer.rows(input({ sort = sort }))
		assert.are.equal("Ysolde", rows[1].short)
		assert.are.same({ key = "name", descending = false }, Officer.nextSort(sort, "name"))
	end)

	it("reads the live addon, the snapshot's age, the gap and the inspected enchants", function()
		local rows = Officer.rows(input())
		local byName = {}
		for _, row in ipairs(rows) do
			byName[row.short] = row
		end
		assert.are.equal("live", byName.Kethra.live)
		assert.are.equal("heard", byName.Ysolde.live)
		assert.are.equal(7200, byName.Ysolde.age)
		assert.are.equal("live", byName.Velisandre.live)
		assert.is_nil(byName.Pell.live)
		assert.are.same({ 1, 62 }, { byName.Kethra.enchants, byName.Kethra.ilvl })
		assert.are.equal(18.6, byName.Dravik.gap)
	end)

	it("counts tonight's readiness over the members coming: addon, version, build, enchants, attunement, supplies",
		function()
			local rows = Officer.rows(input())
			local coming = {}
			for _, row in ipairs(rows) do
				if row.answer == "in" or row.answer == "late" then
					coming[#coming + 1] = row
				end
			end
			local ready = Officer.readiness(coming, { zone = "Onyxia's Lair", data = { encounterNeeds = {
				[10184] = { consumables = { "fire_protection" } } } }, version = "2.0.45",
				pushes = { hunter = { name = "Beast Mastery raid build" }, warlock = { name = "Affliction raid" } },
				roster = { jobs = { held = 1, total = 2 }, assignments = { { label = "Main tank", count = 1,
					holders = { "Brannoc" } }, { label = "Off tank", count = 1, holders = {} } } } })
			local byKey = {}
			for _, entry in ipairs(ready) do
				byKey[entry.key] = entry
			end
			-- Ysolde was heard two hours ago, Brannoc not at all: 3 of 5 live.
			assert.are.same({ 3, 5 }, { byKey.addon.count, byKey.addon.of })
			assert.are.same({ "Brannoc", "Ysolde" }, shorts(byKey.addon.names))
			assert.are.same({ "Dravik" }, shorts(byKey.version.names))
			-- The hunter follows the pushed build; the warlock is off it.
			assert.are.same({ 1, 2 }, { byKey.build.count, byKey.build.of })
			assert.are.same({ "Dravik" }, shorts(byKey.build.names))
			assert.are.equal(1, byKey.enchants.count)
			assert.are.same({ "Velisandre", "Dravik" }, shorts(byKey.attuned.names))
			assert.are.equal("No fire prot.", byKey.supplyfire_protection.label)
			assert.are.same({ "Velisandre", "Dravik" }, shorts(byKey.supplyfire_protection.names))
			assert.is_nil(byKey.supplyflask)
			assert.are.same({ 1, 2, { "Off tank" } }, { byKey.jobs.count, byKey.jobs.of, byKey.jobs.open })
			assert.is_nil(byKey.ready)
		end)

	it("says when nobody coming runs the addon, and shows the last ready check with what each lacks", function()
		local nobody = Officer.readiness({ { short = "Pell", classToken = "WARLOCK" } }, { zone = "Molten Core" })
		assert.is_true(nobody[1].empty)
		local ready = Officer.readiness({}, { check = { clock = "21:04", missing = { { name = "Kethra",
			classToken = "ROGUE", what = { "flask" } } }, coverage = {}, imbueOthers = true } })
		local last = ready[#ready]
		assert.are.same({ "Ready check", "Ready check 21:04" }, { last.label, last.title })
		assert.are.same({ "flask" }, last.names[1].what)
		assert.are.same({ L.offImbueSelf }, last.hint)
	end)

	it("says the site's approvals and a claim that is not settled, nothing without the message", function()
		assert.are.equal("3 approvals on the site", Officer.siteLine({ pending_approvals = 3, claim_state = "claimed" }))
		assert.are.equal("1 approval on the site · Claim contested on the site",
			Officer.siteLine({ pending_approvals = 1, claim_state = "contested" }))
		assert.are.equal(L.offNoApprovals, Officer.siteLine({ pending_approvals = 0 }))
		assert.is_nil(Officer.siteLine(nil))
	end)

	it("words the guild chat reminder with the raid, its time and how many have not signed up", function()
		assert.are.equal("Onyxia's Lair, Thu 10 Dec 20:00: 16 raiders have not signed up yet.",
			Officer.remindText("Onyxia's Lair", "Thu 10 Dec", "20:00", 16))
	end)

	it("keeps the newest push per class and knows the build each class was told to follow", function()
		_G.ForeverSixtyDB = {}
		Officer.keepPush("Thornwake Company", { "FSB1:a", "Beast Mastery raid build" }, "hunter", 10)
		Officer.keepPush("Thornwake Company", { "FSB1:b", "Arms raid" }, "warrior", 20)
		Officer.keepPush("Thornwake Company", { "FSB1:c", "Marksmanship raid" }, "hunter", 30)
		local sent = Officer.sentPushes("Thornwake Company")
		assert.are.same({ "FSB1:c", "FSB1:b" }, { sent[1].code, sent[2].code })
		local byClass = Officer.pushesByClass(sent, { { code = "FSB1:d", name = "Survival raid", at = 40 } }, "hunter")
		assert.are.equal("Survival raid", byClass.hunter.name)
		assert.are.equal("Arms raid", byClass.warrior.name)
		assert.are.equal("Marksmanship raid", Officer.pushesByClass(sent, {}, "hunter").hunter.name)
	end)

	it("cannot push without a build, and says why", function()
		local model = Officer.pushModel({}, nil, {}, {})
		assert.is_false(model.canPush)
		assert.are.equal(L.offPushNoBuild, model.why)
		assert.is_nil(Officer.pushFields({}, nil))
	end)

	it("tells an older addon version from a newer one", function()
		assert.is_true(Officer.olderVersion("2.0.40", "2.0.45"))
		assert.is_false(Officer.olderVersion("2.0.45", "2.0.45"))
		assert.is_false(Officer.olderVersion("2.1.0", "2.0.45"))
	end)
end)

describe("OfficerLoot, the master loot session", function()
	local OfficerLoot, LootCouncil, RaidDrops, state, given

	before_each(function()
		given = {}
		state = mock.install({ class = { name = "Hunter", token = "HUNTER" }, playerName = "Velisandre",
			globals = { date = os.date, GetMasterLootCandidate = function(slot, index)
				local names = { "Velisandre", "Ilvar", "Orwen" }
				return slot == 2 and names[index] or nil
			end, GiveMasterLoot = function(slot, index)
				given[#given + 1] = { slot, index }
			end } })
		_G.ForeverSixtyDB = {}
		require("Theme").reset()
		LootCouncil = helper.load("LootCouncil")
		RaidDrops = helper.load("RaidDrops")
		helper.load("GroupMembers")
		OfficerLoot = helper.load("OfficerLoot")
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
		mock.uninstall()
	end)

	local function drop()
		return { itemId = 17069, lootId = "17069-1000", at = 1000, asked = 2, answers = { x = 1, y = 2 }, slot = 2 }
	end

	local MEMBERS_HERE = {
		{ short = "Velisandre", classToken = "HUNTER", own = true, record = { attendance = 0.88 } },
		{ short = "Ilvar", classToken = "HUNTER", live = {}, record = { attendance = 0.8 } },
		{ short = "Torvald", classToken = "WARRIOR", live = {} },
		{ short = "Orwen", classToken = "PRIEST" },
	}

	it("lists the gains largest first with Won and Att., and sums up who answered, had no gain or no addon", function()
		local rows = {
			{ short = "Ilvar", classToken = "HUNTER", kind = "answer", value = 26.0, unit = "dps", currentItemId = 12651 },
			{ short = "Velisandre", classToken = "HUNTER", kind = "own", value = 21.1, unit = "dps", currentItemId = 13023,
				planned = true },
			{ short = "Torvald", classToken = "WARRIOR", kind = "answer", value = -1, unit = "dps" },
			{ short = "Orwen", classToken = "PRIEST", kind = "noaddon" },
		}
		local view = OfficerLoot.view(drop(), MEMBERS_HERE, rows, { now = 1005, won = function(short)
			return short == "Ilvar" and 1 or 0
		end })
		assert.are.same({ "Ilvar", "Velisandre" }, { view.rows[1].short, view.rows[2].short })
		assert.are.same({ "+26.0", 1, 80, true }, { view.rows[1].gain, view.rows[1].won, view.rows[1].attendance,
			view.rows[1].live })
		assert.is_true(view.rows[2].planned)
		assert.are.same({ "Torvald" }, { view.summary.noGain[1].short })
		assert.are.same({ "Orwen" }, { view.summary.noAddon[1].short })
		assert.are.same({ 2, 2 }, { view.summary.answered, view.summary.asked })
	end)

	it("lists who could wear it with no gain while nobody answered, and says so after 30 s", function()
		local rows = { { short = "Ilvar", classToken = "HUNTER", kind = "waiting" },
			{ short = "Orwen", classToken = "PRIEST", kind = "waiting" } }
		local early = OfficerLoot.view(drop(), MEMBERS_HERE, rows, { now = 1005, won = function()
			return 0
		end })
		assert.are.equal(2, #early.rows)
		assert.is_nil(early.rows[1].gain)
		assert.is_nil(early.quiet)
		local late = OfficerLoot.view(drop(), MEMBERS_HERE, rows, { now = 1031, won = function()
			return 0
		end })
		assert.is_true(late.quiet)
	end)

	it("arms Give on the first click, gives through master loot on the second, and disarms after 3 s", function()
		LootCouncil.open = true
		local item = drop()
		local row = { short = "Ilvar", gain = "+26.0 DPS" }
		assert.are.equal("armed", OfficerLoot.give(item, row, true))
		assert.are.same({}, given)
		assert.are.equal("given", OfficerLoot.give(item, row, true))
		assert.are.same({ { 2, 2 } }, given)
		assert.are.equal("Ilvar", item.givenTo)
		-- A second item: armed, then the 3 s pass.
		local other = drop()
		other.lootId = "other"
		OfficerLoot.give(other, row, false)
		mock.runTimers(state)
		assert.is_nil(OfficerLoot.armed)
	end)

	it("says why Give cannot go: the loot window closed, or a raider master loot will not give it to", function()
		LootCouncil.open = false
		local ok, why = LootCouncil.give(drop(), "Ilvar")
		assert.is_false(ok)
		assert.are.equal(L.offGiveClosed, why)
		LootCouncil.open = true
		ok, why = LootCouncil.give(drop(), "Brannoc")
		assert.is_false(ok)
		assert.are.equal(L.offGiveNotEligible, why)
	end)

	it("counts what a raider won this week from the loot chat, and the night's items given to the biggest gain",
		function()
			local nights = { { date = "2026-12-08", instance = "Molten Core", lastAt = 5000, corpses = { { items = {
				{ itemId = 1, count = 1, winners = { "Ilvar" }, given = 1, toBest = 1 },
				{ itemId = 2, count = 1, winners = { "Ilvar-ClassicBetaPvP" }, given = 1, toBest = 0 },
				{ itemId = 3, count = 1, winners = {} } } } } } }
			assert.are.equal(2, RaidDrops.wonSince(nights, "Ilvar", 1000))
			assert.are.equal(0, RaidDrops.wonSince(nights, "Ilvar", 6000))
			local line = OfficerLoot.lastNight(nights)
			assert.is_truthy(line:find(": 3 items, 1 to the biggest gain$"), line)
			nights[1].corpses[1].items[1].given, nights[1].corpses[1].items[2].given = nil, nil
			assert.is_truthy(OfficerLoot.lastNight(nights):find(": 3 items in Molten Core$"))
			assert.is_nil(OfficerLoot.lastNight({}))
		end)
end)

describe("the Officer tab in the window", function()
	local Hub, OfficerTab, state, officer, inGuild

	--- Every shown font string's text under the window, and its shown client buttons by label.
	local function drawn()
		local texts, buttons = {}, {}
		local function walk(node)
			if not node.shown then
				return
			end
			for _, region in ipairs(node.regions or {}) do
				if region.kind == "FontString" and region.shown and type(region.text) == "string" then
					texts[#texts + 1] = region.text
				end
			end
			if node.foreverSixtyLabel ~= nil and node.foreverSixtyTeal ~= nil then
				buttons[node.foreverSixtyLabel.text] = node
			end
			for _, child in ipairs(node.children or {}) do
				walk(child)
			end
		end
		walk(Hub.frame)
		return table.concat(texts, "\n"), buttons
	end

	local function start(options)
		options = options or {}
		officer = options.officer ~= false
		inGuild = true
		state = mock.install({ class = { name = "Hunter", token = "HUNTER" }, playerName = "Velisandre", level = 60,
			globals = { date = os.date, IsInGuild = function()
				return inGuild
			end, CanEditOfficerNote = function()
				return officer
			end, GetNumGuildMembers = function()
				return #ROSTER
			end, GetGuildRosterInfo = function(index)
				local entry = ROSTER[index]
				return entry.name, "Raider", entry.rankIndex, entry.level, "", "", "", "", entry.online, 0, entry.classToken
			end } })
		state.guild = { name = "Thornwake Company", rankName = "Officer", rankIndex = 1 }
		_G.ForeverSixtyDB = {}
		_G.ForeverSixtyData = { format = 2, generated = os.date("!%Y-%m-%d"), characters = records(),
			guilds = { ["Thornwake Company"] = { raider_ranks = { 1, 3 } } } }
		_G.ForeverSixtyInbox = { messages = { { type = "raid_signup", raid_id = 42, raid = "Onyxia's Lair",
			title = "Onyxia", starts_at = START, answer = "in" }, { type = "raid_roster", raid_id = 42, raid = "Onyxia's Lair",
			title = "Onyxia", starts_at = START, size = 40, signups = { ["in"] = 4, late = 1, bench = 0, out = 1, none = 1 },
			roles = {}, jobs = { held = 1, total = 2 }, members = MEMBERS,
			assignments = { { label = "Off tank", count = 1, holders = {} } } },
			{ type = "guild", pending_approvals = 3, claim_state = "claimed" } } }
		require("Theme").reset()
		helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Export")
		helper.load("Score")
		helper.load("GuildRaid")
		helper.load("LootCouncil")
		helper.load("GroupMembers")
		helper.load("ReadyCheck")
		helper.load("Officer").data = {}
		helper.load("OfficerLoot").data = {}
		OfficerTab = helper.load("OfficerTab")
		Hub = require("Hub")
		mock.enterWorld(state)
	end

	after_each(function()
		if state ~= nil then
			assert.are.same({}, state.clientWrites)
		end
		_G.ForeverSixtyDB, _G.ForeverSixtyInbox, _G.ForeverSixtyData = nil, nil, nil
		mock.uninstall()
		state = nil
	end)

	it("shows only behind the officer gate, and says Officers only. without it", function()
		start({ officer = false })
		assert.is_false(Hub.isAvailable(Hub.tab("officer")))
		assert.are.equal(L.offOnly, Hub.reason("officer"))
		mock.uninstall()
		start()
		assert.is_true(Hub.isAvailable(Hub.tab("officer")))
		assert.are.equal("officer", Hub.open("officer"))
	end)

	it("draws the roster by sign-up, readiness, sign-ups with the site's queue and Remind, and the push", function()
		start()
		Hub.open("officer")
		local text, buttons = drawn()
		for _, words in ipairs({ L.offRoster, L.offBySignup, L.offReadiness, L.guildSignups, "3 approvals on the site",
			L.offPush, L.offLoot, L.offLootIdle, "Off tank open" }) do
			assert.is_truthy(text:find(words, 1, true), words)
		end
		assert.is_truthy(text:find("7 of 7 shown", 1, true))
		assert.is_not_nil(buttons["Remind 1 in chat"])
		assert.is_not_nil(buttons[L.offSend])
		-- No build loaded: Send to guild cannot go.
		assert.is_false(buttons[L.offSend]:IsEnabled())
		assert.are.equal(L.offWhisperHint, Hub.footer.enter.text)
		assert.is_truthy(Hub.headerFor(Hub.tab("officer")).note:find("Thornwake Company · Officer", 1, true))
	end)

	it("caps the roster at what the screen holds and sorts on a heading click", function()
		start()
		assert.are.equal(OfficerTab.LAYOUT.minRows, OfficerTab.maxRows(400))
		assert.are.equal(OfficerTab.LAYOUT.maxRows, OfficerTab.maxRows(4000))
		Hub.open("officer")
		OfficerTab.sort = { key = "score", descending = true }
		OfficerTab.draw()
		assert.are.equal("Brannoc", OfficerTab.current.roster.rows[1].short)
		assert.is_truthy(drawn():find("by Score", 1, true))
	end)

	it("filters to the raiders not signed up with a click on the link", function()
		start()
		Hub.open("officer")
		OfficerTab.filter = "unsigned"
		OfficerTab.draw()
		assert.are.same({ "Pell" }, { OfficerTab.current.roster.rows[1].short })
		OfficerTab.filter = "raiders"
	end)

	it("sends the reminder to guild chat and says so in the footer", function()
		start()
		local said = {}
		_G.SendChatMessage = function(text, channel)
			said[#said + 1] = { text, channel }
		end
		Hub.open("officer")
		local _, buttons = drawn()
		buttons["Remind 1 in chat"]:GetScript("OnClick")(buttons["Remind 1 in chat"], "LeftButton")
		assert.are.equal("GUILD", said[1][2])
		assert.is_truthy(said[1][1]:find("1 raiders have not signed up yet.", 1, true))
		assert.are.equal(L.offReminded, Hub.footer.enter.text)
		_G.SendChatMessage = nil
	end)
end)

describe("the Officer tab's master loot session in the window", function()
	local Hub, OfficerTab, LootCouncil, state, officer, master

	local function start(options)
		options = options or {}
		officer, master = options.officer ~= false, true
		state = mock.install({ class = { name = "Hunter", token = "HUNTER" }, playerName = "Velisandre", level = 60,
			globals = { date = os.date, IsInGuild = function()
				return true
			end, CanEditOfficerNote = function()
				return officer
			end, GetNumGuildMembers = function()
				return 0
			end } })
		state.guild = { name = "Thornwake Company", rankName = "Raider", rankIndex = 3 }
		_G.ForeverSixtyDB = {}
		require("Theme").reset()
		helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Export")
		helper.load("Score")
		helper.load("GuildRaid")
		LootCouncil = helper.load("LootCouncil")
		LootCouncil.isMasterLooter = function()
			return master
		end
		helper.load("GroupMembers")
		helper.load("ReadyCheck")
		helper.load("Officer").data = {}
		helper.load("OfficerLoot").data = {}
		OfficerTab = helper.load("OfficerTab")
		Hub = require("Hub")
		mock.enterWorld(state)
		LootCouncil.drops = { { itemId = 17069, lootId = "17069-1", at = require("Theme").now(), asked = 2, answers = {},
			slot = 2, name = "Striker's Mark", boss = "Magmadar" } }
		LootCouncil.open = true
	end

	after_each(function()
		_G.ForeverSixtyDB = nil
		mock.uninstall()
	end)

	it("opens on the session for a master looter, Escape going back to the roster for an officer", function()
		start()
		assert.are.equal("officer", OfficerTab.openSession())
		assert.are.equal("loot", OfficerTab.mode)
		assert.is_not_nil(OfficerTab.view.texts["session.name"])
		assert.are.equal(L.offEscRoster, Hub.footer.escape.text)
		assert.are.equal(L.offMasterLooter, OfficerTab.figures()[2])
		assert.are.equal("loot", Hub.waitingOf(Hub.tab("officer"))[1].key)
		assert.is_true(OfficerTab.escape())
		assert.are.equal("roster", OfficerTab.mode)
	end)

	it("shows a master looter without the officer gate the session only, and Escape closes", function()
		start({ officer = false })
		assert.is_true(Hub.isAvailable(Hub.tab("officer")))
		assert.are.equal("officer", OfficerTab.openSession())
		assert.is_false(OfficerTab.escape())
		assert.are.equal(L.windowEscClose, Hub.footer.escape.text)
		-- The session over: the tab leaves the strip.
		LootCouncil.drops = {}
		assert.is_false(Hub.isAvailable(Hub.tab("officer")))
	end)

	it("never opens the window in combat", function()
		start()
		_G.InCombatLockdown = function()
			return true
		end
		assert.is_nil(OfficerTab.openSession())
		_G.InCombatLockdown = nil
	end)
end)

describe("the guild's rank names in the export", function()
	local Codec, Export

	before_each(function()
		mock.install({ class = { name = "Hunter", token = "HUNTER" }, globals = {
			GuildControlGetNumRanks = function()
				return 3
			end, GuildControlGetRankName = function(index)
				return ({ "Guild Master", "Officer", "Raider, core" })[index]
			end } })
		Codec = helper.load("Codec")
		Export = helper.load("Export")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("reads the rank names by rank index, the guild master first", function()
		assert.are.same({ "Guild Master", "Officer", "Raider, core" }, Export.rankNames())
	end)

	it("writes them after the guild section, URL-encoded, and reads them back", function()
		local code = Codec.encodeFS1({ dataBuild = "1", classSlug = "hunter", raceSlug = "nightelf",
			treeRanks = { {}, {}, {} }, gearSlots = {}, guild = { name = "Thornwake Company", rankIndex = 1,
				rankNames = { "Guild Master", "Officer", "Raider, core" } } })
		assert.is_truthy(code:find("|guild=Thornwake%20Company:1|guildranks=Guild%20Master,Officer,Raider%2C%20core", 1,
			true), code)
		local build = assert(Codec.decodeFS1(code))
		assert.are.same({ "Guild Master", "Officer", "Raider, core" }, build.guild.rankNames)
		assert.are.equal("Thornwake Company", build.guild.name)
	end)
end)
