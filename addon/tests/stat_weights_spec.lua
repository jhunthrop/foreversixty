-- addon/tests/stat_weights_spec.lua
-- The Stat weights block (surfaces/StatWeights.lua) the Gear and Rotation tabs share: every stat per rating
-- point over the top stat's (the SimC convention, the top stat 1.00), attack power and ranged attack power
-- in one row where the spec weighs them alike, haste apart per 1%, the source note, and both tabs drawing
-- the same heading, columns and rows.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local HUNTER = { agility = 1.6, attack_power = 1, crit = 18, hit = 30, intellect = 0.1, ranged_attack_power = 1,
	strength = 0.2, melee_haste = 3, spell_haste = 1 }

describe("the Stat weights block", function()
	local StatWeights

	before_each(function()
		mock.install({ class = { name = "Hunter", token = "HUNTER" }, level = 60 })
		helper.load("Theme").reset()
		StatWeights = helper.load("StatWeights")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("puts every stat per rating point over the top stat's, one attack power row, haste apart", function()
		local model = StatWeights.model({ weights = HUNTER, dpsPerPoint = 0.1351, reference = "ranged_attack_power",
			measuredBand = 60 }, 60)
		local stats, labels = {}, {}
		for index, row in ipairs(model.rows) do
			stats[index], labels[index] = row.stat, row.label
		end
		-- Hit 30 per 1% is 3 a rating point, crit 18 per 1% is 18/14 a point; intellect is under 5% of hit.
		assert.are.same({ "hit", "agility", "crit", "ranged_attack_power", "strength" }, stats)
		assert.are.same({ "Hit rating", "Agility", "Crit rating", "Ranged attack power", "Strength" }, labels)
		assert.are.equal(1, model.rows[1].scale)
		assert.are.near(1.6 / 3, model.rows[2].scale, 1e-9)
		assert.are.near(3 * 0.1351, model.rows[1].value, 1e-9)
		-- Rating haste counts at melee and spell haste both.
		assert.are.near(4 * 0.1351, model.haste, 1e-9)
		assert.are.same({ "10 hit rating is 1% hit" }, model.rows[1].hover)
		assert.is_nil(model.rows[2].hover)
		assert.are.equal("DPS", model.unit)
		assert.is_nil(model.note)
	end)

	it("says the band the weights were measured at, the companion's sim and its day, HPS for a healer", function()
		assert.are.equal("at level 60", StatWeights.model({ weights = HUNTER, measuredBand = 60 }, 34).note)
		local sim = StatWeights.model({ tag = "sim", generatedAt = "2026-12-07T21:03:00Z",
			weights = { strength = 1, hit = 20 }, caps = { "hit" } }, 60)
		assert.are.equal("your sim, 7 Dec", sim.note)
		assert.is_nil(sim.rows[1].value)
		assert.is_true(sim.rows[1].capped)
		assert.are.same({ "Capped", "muted" }, StatWeights.cells(sim, sim.rows[1])[2])
		assert.are.same({ "-", "muted" }, StatWeights.cells(sim, sim.rows[1])[3])
		assert.is_nil(sim.haste)
		assert.are.equal("HPS", StatWeights.model({ healing = true, weights = { intellect = 1 } }, 60).unit)
		assert.is_nil(StatWeights.model(nil, 60))
		assert.is_nil(StatWeights.model({ weights = { melee_haste = 2 } }, 60))
	end)

	it("keeps at most six rows, the head naming the unit per point", function()
		local many = {}
		for index, stat in ipairs({ "strength", "agility", "stamina", "intellect", "spirit", "attack_power",
			"spell_power", "healing_power", "mp5" }) do
			many[stat] = index
		end
		local model = StatWeights.model({ weights = many, dpsPerPoint = 0.1 }, 60)
		assert.are.equal(StatWeights.MAX_ROWS, #model.rows)
		assert.are.same({ { "Stat", "muted" }, { "Weight", "muted" }, { "DPS per point", "muted" } },
			StatWeights.cells(model, nil))
	end)
end)

describe("the Gear and Rotation tabs' weights", function()
	local KEY = "us/normal/velisandre"
	local SPEC = "hunter-beast-mastery"
	local DATA = {
		build = "1.60.1.70938",
		classes = { hunter = { tabs = { { name = "Beast Mastery", talents = {} }, { name = "Marksmanship", talents = {} },
			{ name = "Survival", talents = {} } } } },
		rotations = { [SPEC] = { { level = 60, lines = { { spellId = 3044, name = "Arcane Shot",
			condition = "when nothing else is ready" } } } } },
		abilityLevels = { [3044] = 6 },
		weights = { [SPEC] = HUNTER },
		dpsPerPoint = { [SPEC] = { reference = "ranged_attack_power", dps = 0.1351, weaponDps = 1.17,
			[60] = { reference = "ranged_attack_power", dps = 0.1351, weaponDps = 1.17 } } },
	}

	local function weightTexts(section)
		local texts = { section.heading:GetText() }
		for _, row in ipairs(section.rows) do
			for _, key in ipairs({ "stat", "weight", "value" }) do
				if row[key]:IsShown() then
					texts[#texts + 1] = row[key]:GetText()
				end
			end
		end
		if section.haste:IsShown() then
			texts[#texts + 1] = section.haste:GetText()
		end
		return texts
	end

	before_each(function()
		mock.install({ class = { name = "Hunter", token = "HUNTER" }, level = 60, playerName = "Velisandre",
			realm = "Thornwake", region = 1, globals = { UnitFactionGroup = function()
				return "Horde"
			end } })
		helper.load("Theme").reset()
		_G.ForeverSixtyDB = {}
		_G.ForeverSixtyInbox = nil
		helper.load("Prefs")
		helper.load("Widgets")
		helper.load("Export").characterKey = function()
			return KEY
		end
		helper.load("Gear")
		helper.load("Follow").build = { classSlug = "hunter", spec = SPEC, name = "Beast Mastery raid", kind = "raid",
			default = true, order = {} }
		helper.load("BuildsWindow")
		helper.load("StatWeights")
		local GearTab = helper.load("GearTab")
		local RotationTab = helper.load("RotationTab")
		-- Both tabs on one hub: helper.load would make the hub again and drop the Rotation tab.
		package.loaded.GearTabView = nil
		require("GearTabView")
		local Hub = require("Hub")
		Hub.data, RotationTab.data = DATA, DATA
		local current = require("Gear").currentWeights(DATA, SPEC, nil, KEY, 60)
		GearTab.read = function()
			return { data = DATA, spec = SPEC, level = 60, faction = "horde", current = current, equipped = {},
				bagIds = {}, noEnchant = {}, isMissing = function()
					return false
				end, statsOf = function()
					return nil
				end }
		end
		local SlotMarks = require("SlotMarks")
		SlotMarks.current = function()
			return { read = { upgrades = {} }, arrows = {}, targets = {} }
		end
	end)

	after_each(function()
		_G.ForeverSixtyDB = nil
		_G.ForeverSixtyInbox = nil
		mock.uninstall()
	end)

	it("draw one block: the same heading, columns, labels and figures, attack power once", function()
		local Hub = require("Hub")
		assert.are.equal("gear", Hub.open("gear"))
		local gear = weightTexts(require("GearTabView").view.weights)
		assert.are.equal("rotation", Hub.open("rotation"))
		local rotation = weightTexts(require("RotationTab").view.weights)
		assert.are.same(rotation, gear)
		assert.are.same({ L.weightsTitle, "Stat", "Weight", "DPS per point", "Hit rating", "1.00", "0.41",
			"Agility", "0.53", "0.22", "Crit rating", "0.43", "0.17", "Ranged attack power", "0.33", "0.14",
			"Strength", "0.07", "0.03", "Haste: 0.54 DPS per 1% haste" }, gear)
	end)
end)
