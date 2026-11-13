-- addon/tests/rotation_tab_spec.lua
-- The Rotation tab of the Forever Sixty window (design/specs/2026-11-13-addon-hub.md 4.4): the stepper's
-- levels (where the priority changes, plus the player's own), a line not trained yet dimmed with its
-- level, the spec shown without a build, How you played it (the mean, least followed first) and Lost
-- across the night (summed by kind), the stat weights with their source, the header figures, and the tab
-- drawn in the window with and without a last raid.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local function line(spellId, name, condition)
	return { spellId = spellId, name = name, condition = condition }
end

local MARK = line(1130, "Hunter's Mark", "when Hunter's Mark is not on the target")
local MARK_2 = line(14323, "Hunter's Mark", "when Hunter's Mark is not on the target")
local STING = line(13549, "Serpent Sting", "when Serpent Sting has under 1.5 s left")
local ARCANE = line(3044, "Arcane Shot", "when nothing else is ready")
local MULTI = line(2643, "Multi-Shot", "on cooldown, above 25% mana (the cited guide's own step)")
local AIMED = line(19434, "Aimed Shot", "on cooldown, shared with Multi-Shot")

local DATA = {
	build = "1.60.1.70938",
	classes = { hunter = { tabs = {
		{ name = "Beast Mastery", talents = {} }, { name = "Marksmanship", talents = {} },
		{ name = "Survival", talents = {} },
	} } },
	rotations = { ["hunter-beast-mastery"] = {
		{ level = 10, lines = { MARK, STING, ARCANE } },
		-- A rank-up only: no stop.
		{ level = 22, lines = { MARK_2, STING, ARCANE } },
		{ level = 40, lines = { MARK_2, MULTI, STING, ARCANE } },
		{ level = 60, lines = { MARK_2, AIMED, MULTI, STING, ARCANE } },
	} },
	abilityLevels = { [1130] = 6, [14323] = 6, [13549] = 4, [3044] = 6, [2643] = 18, [19434] = 60 },
	weights = { ["hunter-beast-mastery"] = { agility = 1.6, attack_power = 1, crit = 18, hit = 30,
		ranged_attack_power = 1, melee_haste = 4 } },
	dpsPerPoint = { ["hunter-beast-mastery"] = { reference = "ranged_attack_power", dps = 0.2183, weaponDps = 1.95,
		[40] = { reference = "ranged_attack_power", dps = 0.1202, weaponDps = 1.06 },
		[60] = { reference = "ranged_attack_power", dps = 0.1351, weaponDps = 1.17 } } },
}

local KEY = "us/normal/velisandre"

local function boss(name, adherence, loss, seconds, extra)
	local entry = { name = name, percentile = 70, adherence = adherence, biggest_loss = loss, loss_seconds = seconds,
		deaths = 0, consumables_missed = {} }
	for field, value in pairs(extra or {}) do
		entry[field] = value
	end
	return entry
end

local LAST_RAID = { type = "last_raid", character = KEY, report_id = "r_1", date = "2026-12-08", bosses = {
	boss("Lucifron", 0.82, "downtime", 18),
	boss("Magmadar", 0.78, "downtime", 24),
	boss("Gehennas", 0.91, "low output", 0),
	boss("Garr", 0.74, "downtime", 31, { deaths = 1, consumables_missed = { "flask" } }),
	boss("Baron Geddon", nil, "dead early", 40),
	boss("Shazzrah", 0.95, nil, 0),
} }

describe("the Rotation tab", function()
	local state, RotationTab, Rotation, Follow, Hub, LastRaid

	before_each(function()
		state = mock.install({ class = { name = "Hunter", token = "HUNTER" }, level = 60, playerName = "Velisandre",
			realm = "Thornwake", region = 1 })
		helper.load("Theme").reset()
		_G.ForeverSixtyDB = {}
		_G.ForeverSixtyInbox = nil
		helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Export").characterKey = function()
			return KEY
		end
		helper.load("Gear")
		Follow = helper.load("Follow")
		Follow.build = nil
		Rotation = helper.load("Rotation")
		helper.load("BuildsWindow")
		LastRaid = helper.load("LastRaid")
		RotationTab = helper.load("RotationTab")
		Hub = require("Hub")
		Hub.data = DATA
		RotationTab.data = DATA
	end)

	after_each(function()
		Follow.build = nil
		_G.ForeverSixtyDB = nil
		_G.ForeverSixtyInbox = nil
		mock.uninstall()
	end)

	describe("the stepper", function()
		it("stops where the priority reads differently, never at a rank-up, and at the player's own level", function()
			assert.are.same({ 10, 40, 60 }, Rotation.changeLevels(DATA, "hunter-beast-mastery"))
			assert.are.same({ 10, 34, 40, 60 }, RotationTab.stops(DATA, "hunter-beast-mastery", 34))
			assert.are.same({ 10, 40, 60 }, RotationTab.stops(DATA, "hunter-beast-mastery", 60))
		end)

		it("steps to the stop before and after, nothing past the ends", function()
			local stops = { 10, 34, 40, 60 }
			assert.are.equal(10, RotationTab.step(stops, 34, -1))
			assert.are.equal(40, RotationTab.step(stops, 34, 1))
			assert.is_nil(RotationTab.step(stops, 10, -1))
			assert.is_nil(RotationTab.step(stops, 60, 1))
		end)
	end)

	describe("the priority", function()
		it("shows the band's lines in order, the condition in sentence case without its aside", function()
			local rows = RotationTab.priority(DATA, "hunter-beast-mastery", 60, 60)
			assert.are.equal(5, #rows)
			assert.are.same({ 1, "Hunter's Mark", "When Hunter's Mark is not on the target" },
				{ rows[1].number, rows[1].name, rows[1].condition })
			assert.are.equal("On cooldown, above 25% mana", rows[3].condition)
			for _, row in ipairs(rows) do
				assert.is_nil(row.at)
			end
		end)

		it("dims a line the character cannot have trained yet with the level it comes at", function()
			local rows = RotationTab.priority(DATA, "hunter-beast-mastery", 60, 34)
			assert.are.equal(60, rows[2].at)
			assert.is_nil(rows[3].at)
			assert.are.equal("at level 60", string.format(L.rotAtLevel, rows[2].at))
		end)
	end)

	describe("which spec", function()
		it("reads the followed build's spec, else the class's first tree with a rotation", function()
			local build = { classSlug = "hunter", spec = "hunter-beast-mastery", order = {} }
			assert.are.same({ "hunter-beast-mastery", true }, { Rotation.specFor(DATA, build, {}, "hunter") })
			assert.are.same({ "hunter-beast-mastery", false }, { Rotation.specFor(DATA, nil, {}, "hunter") })
			assert.are.same({ nil, false }, { Rotation.specFor(DATA, nil, {}, "mage") })
		end)
	end)

	describe("How you played it", function()
		it("lists every boss least followed first, one with no figure last, and the mean of the rest", function()
			local played = LastRaid.played(LAST_RAID, RotationTab.MAX_BOSSES)
			local names = {}
			for _, row in ipairs(played.bosses) do
				names[#names + 1] = row.name
			end
			assert.are.same({ "Garr", "Magmadar", "Lucifron", "Gehennas", "Shazzrah", "Baron Geddon" }, names)
			assert.are.equal(84, played.mean)
			assert.are.equal(5, played.counted)
			assert.are.same({ 74, "Downtime 31 s" }, { played.bosses[1].followed, played.bosses[1].loss })
			assert.are.equal("Low output", played.bosses[4].loss)
			assert.is_nil(played.bosses[5].loss)
			assert.are.equal("Golemagg", LastRaid.shortName("Golemagg the Incinerator"))
			assert.are.equal("The Prophet Skeram", LastRaid.shortName("The Prophet Skeram"))
			assert.are.equal("Sulfuron Harbinger", LastRaid.shortName("Sulfuron Harbinger"))
			assert.are.equal("Tue 8 Dec", played.note)
			assert.are.equal("Tuesday", played.day)
		end)

		it("names the raid before the day when the last raid carries its zone", function()
			local message = { type = "last_raid", date = LAST_RAID.date, zone = "Molten Core", zone_key = "Molten Core",
				bosses = LAST_RAID.bosses }
			assert.are.equal("Molten Core · Tue 8 Dec", LastRaid.played(message, RotationTab.MAX_BOSSES).note)
		end)

		it("sums the losses by kind, the most seconds first, a kind without seconds by its bosses", function()
			local lost = LastRaid.played(LAST_RAID).lost
			assert.are.same({ "Downtime", 73, "73 s" }, { lost[1].label, lost[1].seconds, lost[1].value })
			assert.are.same({ "Dead early", "40 s" }, { lost[2].label, lost[2].value })
			-- Low output counts no seconds: it stays on its boss row.
			assert.are.equal(2, #lost)
		end)

		it("names deaths and what the boss was pulled without in its hover", function()
			local garr = LastRaid.played(LAST_RAID).bosses[1]
			assert.are.same({ "Garr", "Followed 74% of the rotation", "Parse 70", "Biggest loss: Downtime 31 s",
				"Died 1 time", "Pulled without: flask" }, garr.hover)
		end)

		it("caps the bosses and counts the rest", function()
			local many = { date = "2026-12-08", bosses = {} }
			for index = 1, RotationTab.MAX_BOSSES + 3 do
				many.bosses[index] = boss("Boss " .. index, index / 100, "downtime", index)
			end
			local played = LastRaid.played(many, RotationTab.MAX_BOSSES)
			assert.are.equal(RotationTab.MAX_BOSSES, #played.bosses)
			assert.are.equal(3, played.more)
		end)

		it("is nothing without a last raid, and gives Me its weakest boss", function()
			assert.is_nil(LastRaid.played(nil))
			assert.is_nil(LastRaid.played({ date = "2026-12-08", bosses = {} }))
			assert.are.equal("Garr", LastRaid.weakest(LAST_RAID).name)
			assert.is_nil(LastRaid.weakest(nil))
		end)
	end)

	describe("the stat weights", function()
		it("puts every stat per rating point over the top stat's, haste apart per 1%, from the shipped band", function()
			local current = require("Gear").currentWeights(DATA, "hunter-beast-mastery", nil, KEY, 60)
			assert.are.equal(60, current.measuredBand)
			local weights = RotationTab.weightRows(current, 60)
			assert.are.same({ "hit", "agility", "crit", "ranged_attack_power" },
				{ weights.rows[1].stat, weights.rows[2].stat, weights.rows[3].stat, weights.rows[4].stat })
			assert.are.equal(1, weights.rows[1].scale)
			assert.are.near(1.6 / 3, weights.rows[2].scale, 1e-9)
			assert.are.near(3 * 0.1351, weights.rows[1].value, 1e-9)
			assert.are.near(4 * 0.1351, weights.haste, 1e-9)
			-- Measured at the level the stepper shows: nothing to say; at another level, the band.
			assert.is_nil(weights.note)
			assert.are.equal("at level 60", RotationTab.weightRows(current, 34).note)
			assert.are.same({ "10 hit rating is 1% hit" }, weights.rows[1].hover)
			assert.is_nil(weights.rows[2].hover)
			assert.are.same({ "Hit rating", "Agility", "Crit rating", "Ranged attack power" },
				{ weights.rows[1].label, weights.rows[2].label, weights.rows[3].label, weights.rows[4].label })
			-- Attack power weighs as ranged attack power does: one row, the reference's.
			assert.are.equal(4, #weights.rows)
			assert.are.equal("DPS", weights.unit)
		end)

		it("names the companion's sim and its day, with no DPS where the sim gives no figure", function()
			local weights = RotationTab.weightRows({ tag = "sim", generatedAt = "2026-12-07T21:03:00Z",
				weights = { strength = 1, hit = 20 }, caps = { "hit" } })
			assert.are.equal("your sim, 7 Dec", weights.note)
			assert.is_nil(weights.rows[1].value)
			assert.is_true(weights.rows[1].capped)
			assert.is_false(weights.rows[2].capped)
			assert.is_nil(weights.haste)
			assert.is_nil(RotationTab.weightRows(nil))
		end)
	end)

	describe("in the window", function()
		local function textsOf(frame)
			local found = {}
			for _, region in ipairs(frame.regions) do
				if region.kind == "FontString" and region:IsShown() and type(region.text) == "string" then
					found[region.text] = true
				end
			end
			return found
		end

		it("titles the build followed and how the last raid went, and draws both columns", function()
			Follow.build = { classSlug = "hunter", spec = "hunter-beast-mastery", name = "Beast Mastery raid",
				kind = "raid", default = true, order = {} }
			_G.ForeverSixtyInbox = { format = 2, messages = { LAST_RAID } }
			assert.are.equal("rotation", Hub.open("rotation"))
			local header = Hub.headerFor(Hub.tab("rotation"))
			assert.are.equal("Level 60 Hunter · Beast Mastery raid build · followed 84% on Tuesday", header.note)
			local view = RotationTab.view
			local left, right = textsOf(view.left), textsOf(view.right)
			assert.is_true(left["Aimed Shot"])
			assert.is_true(left[string.format(L.rotLevel, 60)])
			assert.is_true(right[string.format(L.rotAllBosses, 5)])
			assert.is_true(right["84%"])
			assert.is_true(left[L.rotWeights])
			assert.is_true(left["Hit rating"])
			assert.is_true(right["Downtime 31 s"])
			assert.is_true(right["73 s"])
			assert.is_false(view.stepper.nextButton:IsEnabled())
			assert.is_true(view.stepper.previous:IsEnabled())
			assert.are.same({}, Hub.waitingOf(Hub.tab("rotation")))
		end)

		it("steps down a level and back, and says once that no raid came in", function()
			assert.are.equal("rotation", Hub.open("rotation"))
			local view = RotationTab.view
			assert.is_true(textsOf(view.right)[L.meNoRaid])
			assert.is_true(textsOf(view.left)[L.rotPriority .. " " .. require("Theme").inlineColor("muted",
				string.format(L.rotFor, "Beast Mastery"))])
			assert.is_false(textsOf(view.right)[L.rotLost] == true)
			view.stepper.previous:GetScript("OnClick")(view.stepper.previous)
			assert.are.equal(40, view.model.shown)
			assert.is_true(textsOf(view.left)[string.format(L.rotLevel, 40)])
			assert.is_nil(textsOf(view.left)["Aimed Shot"])
			local note = Hub.headerFor(Hub.tab("rotation")).note
			assert.are.equal("Level 60 Hunter", note)
			assert.is_true(state ~= nil)
		end)

		it("says no raid is logged yet while the companion runs", function()
			_G.ForeverSixtyInbox = { format = 2, messages = {} }
			assert.are.equal("rotation", Hub.open("rotation"))
			assert.is_true(textsOf(RotationTab.view.right)[L.rotNoRaidLogged])
		end)
	end)
end)
