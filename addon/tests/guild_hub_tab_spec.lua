-- addon/tests/guild_hub_tab_spec.lua
-- The Forever Sixty window's Guild tab (design/specs/2026-11-13-addon-hub.md 4.5, 8): the raids the inbox
-- names with your answer and the four answers behind the companion's feature check, your jobs with and
-- without a partner, the guild's sign-ups from raid_roster (left out without one), Next raids, the guild's
-- raiders by Forever Score (raider ranks, else every level 60), what the tab has waiting, and the frame:
-- raid day answered and not, a queued answer's footer, no raid from the site, not in a guild.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local KEY_NAME = "Velisandre"
local START = "2026-12-10T02:00:00Z"
local LATER = "2026-12-13T01:30:00Z"

local function signup(extra)
	local message = { type = "raid_signup", raid_id = 42, raid = "Onyxia's Lair", title = "Onyxia",
		starts_at = START, answer = "none" }
	for key, value in pairs(extra or {}) do
		message[key] = value
	end
	return message
end

local function job(extra)
	local message = { type = "raid_assignment", raid_id = 42, raid = "Onyxia's Lair", title = "Onyxia",
		starts_at = START, encounter = "Onyxia", slot = "whelps", label = "Phase 2 ranged, west side",
		purpose = "Fire on the whelps from the west wall when she lifts off." }
	for key, value in pairs(extra or {}) do
		message[key] = value
	end
	return message
end

local function roster(extra)
	local message = { type = "raid_roster", raid_id = 42, raid = "Onyxia's Lair", title = "Onyxia",
		starts_at = START, size = 40, signups = { ["in"] = 11, late = 1, bench = 2, out = 1, none = 16 },
		roles = { { role = "tank", coming = 2, wanted = 4 }, { role = "healer", coming = 4, wanted = 10 },
			{ role = "melee", coming = 2, wanted = 12 }, { role = "ranged", coming = 4 } },
		jobs = { held = 1, total = 2 } }
	for key, value in pairs(extra or {}) do
		message[key] = value
	end
	return message
end

describe("GuildRaid, the Guild tab's model", function()
	local GuildRaid, Theme, startAt

	before_each(function()
		mock.install({ class = { name = "Hunter", token = "HUNTER" }, playerName = KEY_NAME,
			globals = { date = os.date } })
		Theme = require("Theme")
		Theme.reset()
		helper.load("Export")
		GuildRaid = helper.load("GuildRaid")
		startAt = require("RaidAssign").epochOf(START)
	end)

	after_each(function()
		mock.uninstall()
	end)

	local function ctx(extra)
		local base = { now = startAt - (3 * 3600 + 12 * 60), accepts = true }
		for key, value in pairs(extra or {}) do
			base[key] = value
		end
		return base
	end

	it("lists the inbox's raids soonest first with their sign-up, jobs and roster, none that are over", function()
		local inbox = { messages = { signup({ raid_id = 43, raid = "Molten Core", starts_at = LATER }), signup(),
			job(), roster() } }
		local raids = GuildRaid.raids(inbox, nil, startAt - 3600)
		assert.are.equal(2, #raids)
		assert.are.same({ 42, 43 }, { raids[1].raidId, raids[2].raidId })
		assert.are.equal(1, #raids[1].jobs)
		assert.are.equal(40, raids[1].roster.size)
		assert.is_nil(raids[2].roster)
		assert.are.equal(1, #GuildRaid.raids(inbox, nil, startAt + 7 * 3600))
		-- An older companion: the assignment alone still names the raid.
		assert.are.equal(42, GuildRaid.raids({ messages = { job() } }, nil, startAt)[1].raidId)
	end)

	it("reads the answer: a queued one newer than the site's, else the site's, Not signed up in client gold", function()
		local answered = signup({ answer = "in", role = "dps", answered_at = "2026-12-08T19:04:00Z" })
		local words = GuildRaid.answerWords(GuildRaid.answerOf(answered, nil))
		assert.are.same({ "In", "success" }, { words.word, words.hex })
		assert.is_truthy(words.detail:find("^as DPS, sent %a%a%a 19:04$"), words.detail)
		local queued = GuildRaid.answerOf(answered, { answer = "late", answered_at = startAt - 60 })
		assert.is_true(queued.pending)
		assert.are.equal("as DPS, not sent yet", GuildRaid.answerWords(queued).detail)
		local old = GuildRaid.answerOf(answered, { answer = "out", answered_at = 1 })
		assert.are.equal("in", old.answer)
		local none = GuildRaid.answerWords(GuildRaid.answerOf(signup(), nil))
		assert.are.same({ L.guildNotSigned, "clientGold" }, { none.word, none.hex })
		assert.are.equal(L.guildOnRoster, GuildRaid.answerWords(GuildRaid.answerOf(nil, nil)).word)
	end)

	it("heads a raid within 18 hours Tonight with a countdown, a later one Next raid without", function()
		local raids = GuildRaid.raids({ messages = { signup() } }, nil, 0)
		local tonight = GuildRaid.tonight(raids[1], ctx())
		assert.are.equal(L.guildTonight, tonight.heading)
		assert.is_truthy(tonight.when:find(" local · in 3 h 12 min$"), tonight.when)
		assert.is_truthy(tonight.when:find("^%a%a%a 10 Dec, "), tonight.when)
		local early = GuildRaid.tonight(raids[1], ctx({ now = startAt - GuildRaid.TONIGHT_SECONDS - 1 }))
		assert.are.equal(L.guildNextRaid, early.heading)
		assert.is_falsy(early.when:find(" in ", 1, true))
		assert.are.equal(L.guildNow, GuildRaid.countdown(startAt, startAt + 60))
		assert.are.equal("in 25 min", GuildRaid.countdown(startAt, startAt - 25 * 60))
	end)

	it("offers the four answers only when the companion takes them, else says to sign up on the site", function()
		local raids = GuildRaid.raids({ messages = { signup({ answer = "late" }) } }, nil, 0)
		local open = GuildRaid.tonight(raids[1], ctx())
		assert.are.same(GuildRaid.ANSWERS, open.buttons)
		assert.are.equal("late", open.selected)
		local closed = GuildRaid.tonight(raids[1], ctx({ accepts = false }))
		assert.is_nil(closed.buttons)
		assert.are.equal(L.guildChangeOnSite, closed.onSite)
		local unsigned = GuildRaid.raids({ messages = { signup() } }, nil, 0)
		assert.are.equal(L.guildSignUpOnSite, GuildRaid.tonight(unsigned[1], ctx({ accepts = false })).onSite)
		assert.is_true(GuildRaid.acceptsAnswers({ accepts = { "raid_answer" } }))
		assert.is_false(GuildRaid.acceptsAnswers({ accepts = {} }))
		assert.is_false(GuildRaid.acceptsAnswers(nil))
	end)

	it("gives each job its label and purpose, the partners by name, none without one", function()
		local raids = GuildRaid.raids({ messages = { signup(), job({ with = "Mirelle, Brannoc", encounter = "Onyxia" }),
			job({ slot = "soul", label = "Soulstones", purpose = "Soulstone a healer.", with = nil }) } }, nil, 0)
		local jobs = GuildRaid.tonight(raids[1], ctx()).jobs
		assert.are.equal(2, #jobs)
		assert.are.same({ "Mirelle", "Brannoc" }, jobs[1].with)
		assert.are.same({}, jobs[2].with)
		assert.are.equal("Soulstones", jobs[2].label)
		assert.is_nil(jobs[1].encounter)
		local mc = GuildRaid.raids({ messages = { job({ title = "Molten Core", encounter = "Magmadar" }) } }, nil, 0)
		assert.are.equal("Magmadar", GuildRaid.tonight(mc[1], ctx()).jobs[1].encounter)
	end)

	it("counts the sign-ups in the web's words, each role against the comp target, short ones marked", function()
		local raids = GuildRaid.raids({ messages = { signup(), roster() } }, nil, 0)
		local model = GuildRaid.signups(raids[1].roster)
		assert.are.equal("12 of 40 coming", model.note)
		assert.are.same({ "11 In", "1 Late", "2 Bench", "1 Out", "16 not signed up" }, {
			model.counts[1].text, model.counts[2].text, model.counts[3].text, model.counts[4].text, model.counts[5].text })
		assert.are.same({ "Tanks 2 of 4", "warning" }, { model.roles[1].text, model.roles[1].hex })
		assert.are.same({ "Ranged 4", "muted" }, { model.roles[4].text, model.roles[4].hex })
		local met = GuildRaid.signups(GuildRaid.raids({ messages = { signup(), roster({ roles = {
			{ role = "tank", coming = 4, wanted = 4 } } }) } }, nil, 0)[1].roster)
		assert.are.equal("success", met.roles[1].hex)
		local unsized = GuildRaid.raids({ messages = { signup(), roster({ size = false }) } }, nil, 0)[1].roster
		assert.are.equal("12 coming", GuildRaid.signups(unsized).note)
		assert.is_nil(GuildRaid.signups(nil))
	end)

	it("lists Next raids with your answer only when there is more than the one on top", function()
		local one = GuildRaid.raids({ messages = { signup() } }, nil, 0)
		assert.is_nil(GuildRaid.nextRaids(one, 42))
		local two = GuildRaid.raids({ messages = { signup({ answer = "in", role = "dps" }),
			signup({ raid_id = 43, raid = "Molten Core", starts_at = LATER }) } }, nil, 0)
		local rows = GuildRaid.nextRaids(two, 42)
		assert.are.equal(2, #rows)
		assert.are.same({ "In as DPS", "success", true }, { rows[1].answer.word, rows[1].answer.hex, rows[1].shown })
		assert.are.same({ L.guildNotSigned, "clientGold", false }, { rows[2].answer.word, rows[2].answer.hex,
			rows[2].shown })
		assert.is_nil(rows[2].coming)
	end)

	describe("the guild's raiders", function()
		local function member(name, level, rankIndex)
			return { name = name .. "-Realm", level = level, rankIndex = rankIndex, classToken = "WARRIOR" }
		end
		local RECORDS = {
			["Brannoc-Realm"] = { score = 1612, tier = "epic", attendance = 1 },
			["Ysolde-Realm"] = { score = 1490, tier = "rare", attendance = 0.96 },
			["Torvald-Realm"] = { score = 1421, tier = "rare", attendance = 0.94 },
			["Kethra-Realm"] = { score = 1377, tier = "rare", attendance = 0.92 },
			["Mirelle-Realm"] = { score = 1302, tier = "rare", attendance = 0.9 },
			["Velisandre-Realm"] = { score = 1184, tier = "rare", attendance = 0.88 },
			["Alt-Realm"] = { score = 1700, tier = "epic" },
		}
		local ROSTER = { member("Brannoc", 60, 1), member("Ysolde", 60, 2), member("Torvald", 60, 3),
			member("Kethra", 60, 3), member("Mirelle", 60, 3), member("Velisandre", 60, 3), member("Alt", 60, 5),
			member("Young", 42, 3) }
		local function record(name)
			return RECORDS[name]
		end

		it("counts every level 60 while the officers have named no raider ranks, your row under the top five", function()
			local standings = GuildRaid.standings(ROSTER, record, nil, "Velisandre")
			assert.are.equal(7, standings.of)
			assert.are.equal("Alt", standings.top[1].short)
			assert.are.equal(GuildRaid.TOP, #standings.top)
			assert.are.equal(7, standings.place)
			assert.are.same({ "Velisandre", 88 }, { standings.own.short, standings.own.attendance })
		end)

		it("counts only the raider ranks once the officers set them on the site", function()
			local standings = GuildRaid.standings(ROSTER, record, { 1, 2, 3 }, "Velisandre")
			assert.are.equal(7, standings.of)
			assert.are.equal("Brannoc", standings.top[1].short)
			assert.are.equal(6, standings.place)
			assert.are.equal("Velisandre", standings.own.short)
			local mine = GuildRaid.standings(ROSTER, record, { 1, 2, 3 }, "Mirelle")
			assert.are.equal(5, mine.place)
			assert.is_nil(mine.own)
		end)

		it("says the guild's standing and raid nights on one line, and an old pack's age", function()
			assert.are.equal("3rd of all guilds · 4 raid nights in 90 days",
				GuildRaid.guildLine({ standing = 3, nights = 4, attendance = 0.8 }))
			assert.are.equal(Theme.inlineColor("warning", "Ratings 8 days old"), GuildRaid.guildLine(nil, 8))
			assert.are.equal("3rd of all guilds · " .. Theme.inlineColor("warning", "Ratings 8 days old"),
				GuildRaid.guildLine({ standing = 3, nights = 4 }, 8))
			assert.is_nil(GuildRaid.guildLine({}))
		end)
	end)

	it("waits on tonight's raid not signed up and on tonight's jobs, nothing on a raid further out", function()
		local unsigned = GuildRaid.raids({ messages = { signup(), job() } }, nil, 0)
		local items = GuildRaid.waitingOf(unsigned, ctx())
		assert.are.same({ "signup", "job" }, { items[1].key, items[2].key })
		assert.is_truthy(items[1].line:find("^Onyxia's Lair .- · " .. Theme.inlineColor("clientGold", "Sign up"), 1),
			items[1].line)
		local answered = GuildRaid.raids({ messages = { signup({ answer = "in" }) } }, nil, 0)
		assert.are.same({}, GuildRaid.waitingOf(answered, ctx()))
		assert.are.same({}, GuildRaid.waitingOf(unsigned, ctx({ now = startAt - GuildRaid.TONIGHT_SECONDS - 60 })))
		local queued = GuildRaid.waitingOf(GuildRaid.raids({ messages = { signup() } }, nil, 0), ctx({
			pending = function()
				return { answer = "in", answered_at = startAt - 60 }
			end }))
		assert.are.same({}, queued)
	end)

	it("leads Me's Right now with the soonest raid, your answer in its colour and your job", function()
		local lead = GuildRaid.leadOf(GuildRaid.raids({ messages = { signup({ answer = "in" }), job() } }, nil, 0), ctx())
		assert.is_truthy(lead.line:find("^Onyxia's Lair .- · " .. Theme.inlineColor("success", "In") .. "$"), lead.line)
		assert.are.equal(L.guildYourJob .. " Phase 2 ranged, west side", lead.sub)
		assert.are.equal(GuildRaid.RAID_ICONS["Onyxia's Lair"], lead.icon)
		assert.is_nil(GuildRaid.leadOf({}, ctx()))
	end)

	it("keeps the newest outbox answer per raid and character and drops ones older than a week", function()
		local Export = require("Export")
		local saved = { { character = "Bob-Ash", raid_id = 42, answer = "late", answered_at = startAt - 100 },
			{ character = "Bob-Ash", raid_id = 7, answer = "in", answered_at = startAt - 8 * 86400 } }
		local added = { { character = "Bob-Ash", raid_id = 42, answer = "in", answered_at = startAt - 10 } }
		local list = Export.mergeRaidAnswers(saved, added, startAt)
		assert.are.equal(1, #list)
		assert.are.equal("in", list[1].answer)
		local entry = { character = "Bob-Ash", raid_id = 42, answer = "in", answered_at = 5 }
		assert.are.same({ entry }, Export.outbox({ character = "Bob-Ash", code = "FS1:x", wallNow = 0,
			raidAnswers = { entry } }).raid_answer)
	end)
end)

describe("the Guild tab in the window", function()
	local Hub, GuildHubTab, GuildRaid, state, inGuild, startAt

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

	local function start(inbox, options)
		options = options or {}
		inGuild = options.guildless ~= true
		state = mock.install({ class = { name = "Hunter", token = "HUNTER" }, playerName = KEY_NAME, level = 60,
			globals = { date = os.date, IsInGuild = function()
				return inGuild
			end } })
		state.guild = inGuild and { name = "Thornwake Company", rankName = "Officer", rankIndex = 1 } or nil
		_G.ForeverSixtyDB = {}
		_G.ForeverSixtyInbox = inbox
		require("Theme").reset()
		helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Export")
		helper.load("Score").clock = function()
			return startAt - (3 * 3600 + 12 * 60)
		end
		GuildRaid = helper.load("GuildRaid")
		GuildHubTab = helper.load("GuildHubTab")
		GuildHubTab.data = {}
		Hub = require("Hub")
		mock.enterWorld(state)
	end

	before_each(function()
		startAt = require("RaidAssign").epochOf(START)
	end)

	after_each(function()
		if state ~= nil then
			assert.are.same({}, state.clientWrites)
		end
		_G.ForeverSixtyDB, _G.ForeverSixtyInbox = nil, nil
		mock.uninstall()
		state = nil
	end)

	it("shows tonight's raid not signed up with the four answers, the job and its partner, and the sign-ups", function()
		start({ accepts = { "raid_answer" }, messages = { signup(), job({ with = "Mirelle" }), roster() } })
		assert.are.equal("guild", Hub.open("guild"))
		local text, buttons = drawn()
		for _, label in ipairs({ "In", "Late", "Bench", "Out" }) do
			assert.is_not_nil(buttons[label], label)
			assert.is_false(buttons[label].foreverSixtySelected, label)
		end
		assert.is_truthy(text:find(L.guildTonight, 1, true))
		assert.is_truthy(text:find("Onyxia's Lair", 1, true))
		assert.is_truthy(text:find(L.guildNotSigned, 1, true))
		assert.is_truthy(text:find("Phase 2 ranged, west side", 1, true))
		assert.is_truthy(text:find("With |cff", 1, true))
		assert.is_truthy(text:find("12 of 40 coming", 1, true))
		assert.is_truthy(text:find("Tanks 2 of 4", 1, true))
		assert.is_falsy(text:find("jobs filled", 1, true))
		assert.are.equal("Velisandre", Hub.headerFor(Hub.tab("guild")).title)
		assert.is_truthy(Hub.headerFor(Hub.tab("guild")).note:find("Thornwake Company · Officer", 1, true))
		assert.are.same({ "signup", "job" }, { Hub.waitingOf(Hub.tab("guild"))[1].key,
			Hub.waitingOf(Hub.tab("guild"))[2].key })
	end)

	it("queues an answer in the outbox, the pressed one teal, the footer saying when it sends", function()
		start({ accepts = { "raid_answer" }, messages = { signup() } })
		Hub.open("guild")
		local _, buttons = drawn()
		buttons.Late:GetScript("OnClick")(buttons.Late, "LeftButton")
		local text, after = drawn()
		assert.is_true(after.Late.foreverSixtySelected)
		assert.is_truthy(text:find(L.guildNotSent, 1, true))
		assert.are.equal(L.guildSendsLater, Hub.footer.enter.text)
		assert.are.equal("late", GuildRaid.pendingAnswer(42).answer)
		assert.are.same({}, Hub.waitingOf(Hub.tab("guild")))
	end)

	it("shows an answered raid, the job alone, and leaves Sign-ups out without raid_roster", function()
		start({ accepts = { "raid_answer" }, messages = { signup({ answer = "in", role = "dps",
			answered_at = "2026-12-08T19:04:00Z" }), job() } })
		Hub.open("guild")
		local text, buttons = drawn()
		assert.is_true(buttons.In.foreverSixtySelected)
		assert.is_truthy(text:find("as DPS, sent", 1, true))
		assert.is_falsy(text:find("With ", 1, true))
		assert.is_falsy(text:find(L.guildSignups, 1, true))
	end)

	it("says to sign up on the site when the companion does not take answers", function()
		start({ messages = { signup() } })
		Hub.open("guild")
		local text, buttons = drawn()
		assert.is_nil(buttons.In)
		assert.is_truthy(text:find(L.guildSignUpOnSite, 1, true))
		assert.is_false(GuildRaid.answer(42, "in"))
	end)

	it("says no raid from the site without a companion, and none from the guild with one", function()
		start(nil)
		Hub.open("guild")
		assert.is_truthy(drawn():find(L.guildNoRaid, 1, true))
		mock.uninstall()
		start({ messages = {} })
		Hub.open("guild")
		assert.is_truthy(drawn():find(string.format(L.guildNoRaids, "Thornwake Company"), 1, true))
	end)

	it("puts a Next raids row's raid on top with a click", function()
		start({ accepts = { "raid_answer" }, messages = { signup(),
			signup({ raid_id = 43, raid = "Molten Core", title = "Molten Core", starts_at = LATER }) } })
		Hub.open("guild")
		assert.is_truthy(drawn():find(L.guildNext, 1, true))
		GuildHubTab.select(43)
		assert.are.equal(L.guildNextRaid, GuildHubTab.current.tonight.heading)
		assert.are.equal("Molten Core", GuildHubTab.current.tonight.raid)
	end)

	it("names the build the officers sent, who sent it and the day, and whether you follow it", function()
		start(nil)
		local sent = GuildHubTab.officersSent({ { code = "FSB1:x", name = "Beast Mastery raid build", from = "Brannoc-Realm",
			at = startAt } }, function()
			return true
		end)
		assert.are.same({ "Beast Mastery raid build", "Brannoc-Realm", true }, { sent.name, sent.from, sent.following })
		assert.are.equal(GuildRaid.weekday(startAt), sent.day)
		assert.is_nil(GuildHubTab.officersSent({}, function() end))
	end)

	it("leaves the tab off the strip outside a guild", function()
		start(nil, { guildless = true })
		assert.is_false(Hub.isAvailable(Hub.tab("guild")))
		assert.are.equal(L.hubNotInGuild, Hub.reason("guild"))
	end)
end)
