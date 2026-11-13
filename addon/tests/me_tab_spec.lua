-- addon/tests/me_tab_spec.lua
-- The Me tab of the Forever Sixty window (design/specs/2026-11-13-addon-hub.md 4.1 and 8): the Forever Score
-- with its five parts on one scale and the web's words, the last raid boss by boss with its summary, Right now
-- from each other tab's own lead (a link to that tab), the sim with its biggest gains and Sim again, and the
-- footer's freshness; the empty states when the ratings pack or the companion is missing; the raid night that
-- waits on Me until it is seen; the drawn tab in the hub with its header figures.
local helper = require("spec_helper")
local mock = require("wow_mock")

local RECORD = { score = 1184, tier = "rare", as_of = "2026-12-09", fights = 22,
	components = { gear = 410, experience = 380, parses = 228, attendance = 120, discipline = 46 } }

local function boss(name, percentile, loss, seconds, deaths, missed)
	return { name = name, percentile = percentile, biggest_loss = loss, loss_seconds = seconds, deaths = deaths or 0,
		consumables_missed = missed or {} }
end

local LAST_RAID = { type = "last_raid", report_id = "r_8f2a1c", date = "2026-12-08", bosses = {
	boss("Lucifron", 71, "downtime", 18, 0, { "flask" }),
	boss("Magmadar", 64, "downtime", 24, 1, { "flask" }),
	boss("Gehennas", 83, "Hunter's Mark late", 6),
	boss("Garr", 58, "downtime", 31),
	boss("Baron Geddon", 77, "Aimed Shot late", 9),
	boss("Shazzrah", 88, "Serpent Sting off", 4),
	boss("Sulfuron Harbinger", 69, "downtime", 15),
	boss("Golemagg", 92, "downtime", 5),
} }

local SIM = { type = "sim_result", dps = 941.2, error = 3.1, generated_at = "2026-12-07T20:55:00Z",
	slots = { { slot = "hands", item_id = 867, delta = 11.2 }, { slot = "feet", item_id = 16675, delta = 7.9 },
		{ slot = "trinket1", item_id = 11815, delta = 9.0 }, { slot = "neck", item_id = 18404, delta = 2.5 },
		{ slot = "waist", item_id = 1, delta = -1 } } }

local CURRENT = { available = true, generated = "2026-12-09", stale = false }
--- A key the input leaves out (a nil in a table literal would not override the default).
local NONE = {}

describe("the Me tab", function()
	local MeTab, L, Score

	local function full(extra)
		local input = {
			level = 60, record = RECORD, pack = CURRENT, packDays = 1, lastRaid = LAST_RAID, raidName = "Molten Core",
			companion = true, syncedAgo = "14 min ago", sim = SIM, upgrades = {}, spec = "hunter-beast-mastery",
			leads = {}, urgentOrder = { "guild", "officer", "builds", "gear", "me" },
		}
		for key, value in pairs(extra or {}) do
			if value == NONE then
				input[key] = nil
			else
				input[key] = value
			end
		end
		return input
	end

	before_each(function()
		mock.install({ class = { name = "Hunter", token = "HUNTER" }, level = 60, playerName = "Bob" })
		helper.load("Theme")
		L = helper.load("Locale")
		Score = helper.load("Score")
		MeTab = helper.load("MeTab")
	end)

	after_each(function()
		mock.uninstall()
	end)

	describe("the model with full data", function()
		it("shows the score in its tier, the points to the next tier, the date and fights, and every part", function()
			local score = MeTab.model(full()).score
			assert.are.equal(1184, score.total)
			assert.are.equal("Rare", score.tierLabel)
			assert.are.equal(Score.tierHex("rare"), score.tierHex)
			assert.are.equal("316 to Epic (1500)", score.toNext)
			assert.are.same({ "as of 9 Dec", "22 fights" }, score.note)
			assert.is_false(score.noteWarn)
			assert.are.equal(5, #score.parts)
			assert.are.equal("Gear", score.parts[1].label)
			assert.are.equal(500, score.parts[1].max)
			assert.are.equal(Score.tierHex("rare"), score.parts[1].hex)
			assert.are.equal("Preparation", score.parts[5].label)
			assert.are.equal("2 flasks missed", score.parts[5].reason)
			assert.is_nil(score.parts[1].reason)
		end)

		it("says Top tier at the top", function()
			local record = { score = 2100, tier = "legendary", components = {} }
			assert.are.equal(L.meTopTier, MeTab.model(full({ record = record })).score.toNext)
		end)

		it("lays the last raid out boss by boss with its note, summary and report", function()
			local raid = MeTab.model(full()).lastRaid
			assert.are.equal("Molten Core · Tue 8 Dec · 8 of 10 bosses", raid.note)
			assert.are.equal(8, #raid.rows)
			assert.are.same({ name = "Lucifron", parse = 71, parseHex = require("Theme").parseHex(71),
				loss = "Downtime 18 s", died = 0, consumables = "No flask", missed = true }, raid.rows[1])
			assert.are.equal("Hunter's Mark late 6 s", raid.rows[3].loss)
			assert.are.equal("All up", raid.rows[3].consumables)
			assert.are.equal(74, raid.summary.median)
			assert.are.equal("1 death · best Golemagg 92 · worst Garr 58", raid.summary.rest)
			assert.are.equal("foreversixty.gg/reports/r_8f2a1c", raid.url)
		end)

		it("counts bosses without a raid name it knows, and shows a loss-free boss plainly", function()
			local message = { date = "2026-12-08", bosses = { { name = "Onyxia", percentile = 99 } } }
			local raid = MeTab.lastRaid({ lastRaid = message, companion = true })
			assert.are.equal("Tue 8 Dec · 1 bosses", raid.note)
			assert.is_nil(raid.rows[1].loss)
			assert.are.equal("0 deaths · best Onyxia 99", raid.summary.rest)
			assert.is_nil(raid.url)
		end)

		it("shows the sim's DPS, error, date and its three biggest gains, largest first", function()
			local sim = MeTab.model(full()).sim
			assert.are.equal("941.2 DPS", sim.dps)
			assert.are.equal("± 3.1", sim.error)
			assert.are.equal("Simmed Mon 7 Dec", sim.date)
			assert.are.same({ { itemId = 867, gain = 11.2, slot = "hands" }, { itemId = 11815, gain = 9.0, slot = "trinket1" },
				{ itemId = 16675, gain = 7.9, slot = "feet" } }, sim.gains)
			assert.are.equal(L.meSimAgain, sim.button)
			assert.is_false(sim.queued)
		end)

		it("prices the gains from the inbox's upgrades when the run had no per-slot figures", function()
			local upgrades = { { item_id = 5, delta = 4.2, unit = "dps", slot = "chest" },
				{ item_id = 6, delta = 9, unit = "score", slot = "legs" }, { item_id = 7, delta = 1.5, slot = "feet" } }
			local sim = MeTab.model(full({ sim = { dps = 900, error = 0, slots = {} }, upgrades = upgrades })).sim
			assert.are.same({ { itemId = 5, gain = 4.2, slot = "chest" }, { itemId = 7, gain = 1.5, slot = "feet" } },
				sim.gains)
		end)

		it("says Queued while a sim request waits, and has no button without a spec", function()
			assert.are.equal(L.meSimQueued, MeTab.model(full({ queued = true })).sim.button)
			assert.is_nil(MeTab.model(full({ spec = NONE })).sim.button)
		end)

		it("says the companion's sync age in the footer", function()
			assert.are.equal("Companion synced 14 min ago", MeTab.model(full()).footer)
		end)
	end)

	describe("Right now", function()
		local function entry(tab, label, lead, waiting)
			return { tab = tab, label = label, lead = lead, waiting = waiting or {} }
		end

		it("puts the tabs with something waiting first in the hub's order, then the others, four at most", function()
			local rows = MeTab.rightNow({
				entry("builds", "Builds", { line = "Beast Mastery raid build", sub = "51 of 51 on plan", icon = "spec" }),
				entry("gear", "Gear", { line = "Gloves of Holy Might +11.2 DPS", sub = "Hands · in your bags", itemId = 867 },
					{ { key = "bag", line = "Equip Gloves of Holy Might +11.2" } }),
				entry("rotation", "Rotation", { line = "Garr: 74% followed", sub = "Your weakest boss on Tuesday" }),
				entry("guild", "Guild", nil, { { key = "tonight", line = "Onyxia's Lair 20:00 · Sign up" } }),
				entry("officer", "Officer", nil, {}),
			}, { "guild", "officer", "builds", "gear", "me" })
			assert.are.equal(4, #rows)
			assert.are.same({ "guild", "gear", "rotation", "builds" }, { rows[1].tab, rows[2].tab, rows[3].tab, rows[4].tab })
			-- A tab without a lead shows its first waiting line; a lead is the tab's own two lines.
			assert.are.equal("Onyxia's Lair 20:00 · Sign up", rows[1].line)
			assert.is_nil(rows[1].sub)
			assert.is_true(rows[1].urgent)
			assert.are.equal(867, rows[2].itemId)
			assert.are.equal("Hands · in your bags", rows[2].sub)
			assert.is_false(rows[4].urgent)
			assert.are.equal("Builds", rows[4].label)
		end)

		it("leaves out a tab with neither a lead nor anything waiting", function()
			assert.are.same({}, MeTab.rightNow({ entry("officer", "Officer") }, {}))
		end)
	end)

	describe("the empty states (spec 8)", function()
		it("says where the ratings pack comes from when it is not installed, and the footer says so", function()
			local model = MeTab.model(full({ record = NONE, pack = { available = false, absent = true } }))
			assert.are.same({ line = L.meNoPack }, model.score)
			assert.are.equal("No ratings pack · companion synced 14 min ago", model.footer)
			assert.is_not_nil(model.lastRaid)
		end)

		it("keeps the figures of a stale pack, the date and the footer in the warning colour", function()
			local Theme = require("Theme")
			local stale = { available = true, generated = "2026-12-02", stale = true }
			local model = MeTab.model(full({ pack = stale, packDays = 8 }))
			assert.are.equal(1184, model.score.total)
			assert.is_true(model.score.noteWarn)
			assert.are.equal(Theme.inlineColor("warning", "Ratings 8 days old") .. " · companion synced 14 min ago",
				model.footer)
		end)

		it("without the companion: the raid says how to get one, the sim is left out, the footer says so", function()
			local model = MeTab.model(full({ companion = false, lastRaid = NONE, sim = NONE, syncedAgo = NONE }))
			assert.are.same({ empty = L.meNoRaid, link = "foreversixty.gg/app" }, model.lastRaid)
			assert.is_nil(model.sim)
			assert.are.equal("No companion", model.footer)
			assert.are.equal(1184, model.score.total)
		end)

		it("with the companion and no raid yet, leaves the raid out (nothing to say, nothing to do)", function()
			assert.is_nil(MeTab.model(full({ lastRaid = NONE })).lastRaid)
		end)

		it("says nothing of raids or the companion below level 10", function()
			assert.is_nil(MeTab.model(full({ companion = false, lastRaid = NONE, level = 9 })).lastRaid)
		end)

		it("says No logs yet from level 10, and nothing below it", function()
			assert.are.same({ line = L.meNoLogs }, MeTab.model(full({ record = NONE, level = 10 })).score)
			assert.is_nil(MeTab.model(full({ record = NONE, level = 9 })).score)
		end)
	end)

	describe("what waits on Me", function()
		local DAY = 86400

		it("is the raid night until it is seen, in the chip's words, new again with another report", function()
			local night = Score.dateSeconds("2026-12-08") + 22 * 3600
			local items = MeTab.waitingItems(LAST_RAID, night)
			assert.are.equal(1, #items)
			assert.are.equal("Tuesday: median 74, best Golemagg", items[1].line)
			assert.are.equal("r_8f2a1c", items[1].fact)
			assert.are.same({}, MeTab.waitingItems(LAST_RAID, night, function()
				return true
			end))
			assert.are.same({}, MeTab.waitingItems(LAST_RAID, night + 3 * DAY))
			assert.are.same({}, MeTab.waitingItems(nil, night))
		end)
	end)

	it("names the weekday without the client's clock", function()
		local Compat = require("Compat")
		assert.are.equal("Tue 8 Dec", Compat.dayDate("2026-12-08"))
		assert.are.equal("Thu 29 Feb", Compat.dayDate("2024-02-29"))
		assert.are.equal("Thursday", Compat.weekdayName("1970-01-01"))
		assert.is_nil(Compat.dayDate("8 Dec"))
	end)
end)

describe("the Me tab in the window", function()
	local state, Hub, MeFrame, L

	local INBOX = { format = 2, generated_at = "2026-12-09T09:00:00Z", messages = {
		{ type = "own_score", score = 1184, tier = "rare", as_of = "2026-12-09",
			fights = 22, gear = 410, experience = 380, parses = 228, attendance = 120, discipline = 46 },
		LAST_RAID, SIM,
		{ type = "loot_values", raid = "Molten Core", items = {} },
	} }

	local function load(inbox)
		_G.ForeverSixtyInbox = inbox
		helper.load("Theme").reset()
		helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Export")
		helper.load("Follow")
		helper.load("BuildDialog")
		helper.load("Score")
		helper.load("MeTab")
		package.loaded.MeFrame = nil
		MeFrame = require("MeFrame")
		Hub = require("Hub")
		L = require("Locale")
	end

	local function open(inbox)
		load(inbox)
		return Hub.open("me")
	end

	before_each(function()
		state = mock.install({ class = { name = "Hunter", token = "HUNTER" }, level = 60, playerName = "Bob",
			realm = "Ashbringer", region = 1, guild = { name = "Thornwake Company", rankName = "Raider", rankIndex = 3 } })
		state.time = require("Score").dateSeconds("2026-12-09") + 9 * 3600 + 14 * 60
		_G.time = function()
			return state.time
		end
		_G.ForeverSixtyDB = {}
		_G.ForeverSixtyData = nil
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
		_G.ForeverSixtyInbox = nil
		_G.time = nil
		mock.uninstall()
	end)

	it("opens on Me with the guild in the header and draws the score, the raid and the footer", function()
		assert.are.equal("me", open(INBOX))
		local header = Hub.headerFor(Hub.tab("me"))
		assert.are.equal("Bob", header.title)
		assert.are.equal("Level 60 Hunter · Thornwake Company", header.note)
		local parts = MeFrame.parts
		assert.are.equal("1184", parts.score.number:GetText())
		assert.is_true(parts.score.number:IsShown())
		assert.are.equal("Lucifron", parts.raid.rows[1].name:GetText())
		assert.are.equal("Molten Core · Tue 8 Dec · 8 of 10 bosses", parts.raid.note:GetText())
		assert.is_true(parts.raid.fullLog:IsShown())
		assert.are.equal(L.meSim, parts.sim.heading:GetText())
		assert.is_true(Hub.footer.enter:IsShown())
	end)

	it("lists the raid night as waiting until Me is opened", function()
		load(INBOX)
		local waiting = Hub.waitingOf(Hub.tab("me"))
		assert.are.equal(1, #waiting)
		assert.are.equal("Tuesday: median 74, best Golemagg", waiting[1].line)
		assert.are.equal("me", Hub.landing())
		Hub.open()
		assert.are.same({}, Hub.waitingOf(Hub.tab("me")))
	end)

	it("without the companion or the pack: the pack's line, the app's link, no sim", function()
		assert.are.equal("me", open(nil))
		local parts = MeFrame.parts
		assert.are.equal(L.meNoPack, parts.score.line:GetText())
		assert.are.equal(L.meNoRaid, parts.raid.empty:GetText())
		assert.is_true(parts.raid.app:IsShown())
		assert.is_false(parts.sim.heading:IsShown())
		assert.is_false(parts.raid.rows[1].name:IsShown())
	end)

	it("leads Right now with another tab's own line, and its link opens that tab", function()
		load(INBOX)
		Hub.register({ id = "gear", label = "Gear", build = function() end, show = function() end,
			lead = function()
				return { line = "Gloves of Holy Might +11.2 DPS", sub = "Hands · in your bags", itemId = 867 }
			end })
		Hub.open("me")
		local row = MeFrame.parts.now.rows[1]
		assert.are.equal("Gloves of Holy Might +11.2 DPS", row.line:GetText())
		assert.is_true(row.link:IsShown())
		row.link:GetScript("OnClick")(row.link)
		assert.are.equal("gear", Hub.active)
	end)
end)
