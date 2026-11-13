-- addon/tests/rotation_spec.lua
-- The per-level rotation lines (design section 2 item 3, section 6 Wave
-- B): ns.Data.rotations[spec] is a band at each data/pipeline/addonrotation
-- LEVEL_BANDS rung (10, 20, 30, 38, 40, 50, 60) and at every level between
-- where a line changes.
-- Rotation.lua picks the spec and the band for a level and the levels it changes at;
-- this covers that against a small hand-built rotations table, the same
-- shape addonlua.py renders into Data.lua.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local function line(spellId, name, condition)
	return { spellId = spellId, name = name, condition = condition or "" }
end

local DATA = {
	build = "1.60.1.69893",
	classes = { warrior = { tabs = {
		{ name = "Arms", talents = {} },
		{ name = "Fury", talents = { { name = "A", tier = 1, column = 1, maxRank = 5 } } },
		{ name = "Protection", talents = {} },
	} } },
	weights = {},
	rotations = {
		["warrior-fury"] = {
			{ level = 10, lines = { line(2687, "Bloodrage", "Bloodrage on cooldown."), line(284, "Heroic Strike") } },
			{ level = 20, lines = { line(2687, "Bloodrage", "Bloodrage on cooldown."), line(285, "Heroic Strike") } },
			{ level = 40, lines = {
				line(2687, "Bloodrage", "Bloodrage on cooldown."),
				line(12328, "Death Wish", "Death Wish on cooldown."),
				line(23881, "Bloodthirst", "Bloodthirst is the highest damage per rage button."),
				line(1680, "Whirlwind", "Whirlwind while Bloodthirst is down."),
				line(1608, "Heroic Strike"),
			} },
		},
	},
}

-- 2 points into 1:1 (Fury tab, index 2 in DATA.classes.warrior.tabs).
local BUILD = { classSlug = "warrior", order = {
	{ tab = 2, tier = 1, column = 1 }, { tab = 2, tier = 1, column = 1 },
} }

-- Gear.specOf picks the tab with the most points; two points in the Fury
-- tab (index 2) is enough to read this build as warrior-fury.
local RANKS = { [2] = { ["1:1"] = 2 } }

describe("Rotation", function()
	local Rotation

	before_each(function()
		mock.install({ class = { name = "Warrior", token = "WARRIOR" } })
		helper.load("Gear")
		Rotation = helper.load("Rotation")
	end)

	after_each(function()
		mock.uninstall()
	end)

	describe("bandAt", function()
		local SPEC = "warrior-fury"

		it("picks the earliest band for a level below every rung", function()
			assert.are.equal(10, Rotation.bandAt(DATA, SPEC, 5).level)
		end)

		it("picks the highest rung at or below the level, not the nearest", function()
			assert.are.equal(20, Rotation.bandAt(DATA, SPEC, 25).level)
			assert.are.equal(40, Rotation.bandAt(DATA, SPEC, 60).level)
			assert.are.equal(5, #Rotation.bandAt(DATA, SPEC, 40).lines)
		end)

		-- Round 16: a band starts wherever a line changes (a rank row from 34, a rank learned
		-- at 34), not only at the ladder's rungs; the player reads it from its own level.
		it("reads a band that starts between two rungs from its own level", function()
			local data = { build = DATA.build, classes = DATA.classes, weights = {}, rotations = {
				[SPEC] = {
					{ level = 30, lines = { line(285, "Heroic Strike") } },
					{ level = 34, lines = { line(1608, "Heroic Strike") } },
					{ level = 38, lines = { line(1608, "Heroic Strike"), line(1680, "Whirlwind") } },
				},
			} }
			assert.are.equal(30, Rotation.bandAt(data, SPEC, 33).level)
			local at35 = Rotation.bandAt(data, SPEC, 35)
			assert.are.equal(34, at35.level)
			assert.are.equal(1608, at35.lines[1].spellId)
			-- A rank-up is no stop for the stepper; a new line is.
			assert.are.same({ 30, 38 }, Rotation.changeLevels(data, SPEC))
		end)

		it("is nothing for a spec with no curated rotation", function()
			assert.is_nil(Rotation.bandAt(DATA, "warrior-arms", 10))
			assert.are.same({}, Rotation.changeLevels(DATA, "warrior-arms"))
		end)
	end)

	describe("specFor", function()
		it("reads the build's spec, else the tree the points name", function()
			assert.are.same({ "warrior-fury", true }, { Rotation.specFor(DATA, BUILD, RANKS, "warrior") })
			assert.are.same({ "warrior-fury", false }, { Rotation.specFor(DATA, nil, RANKS, "warrior") })
		end)
	end)

	describe("plainCondition", function()
		--- One unit per character: the width the specs measure lines in.
		local function measure(text)
			return #text
		end

		it("strips parentheticals and cuts at the first sentence end", function()
			assert.are.equal("Arcane Shot as the last filler, only when mana is above a 25% reserve.",
				Rotation.plainCondition("Arcane Shot as the last filler, only when mana is above a 25% reserve "
					.. "(Multi-Shot and Serpent Sting are the shots worth saving for). It is instant."))
		end)

		it("strips a nested parenthetical whole", function()
			assert.are.equal("Mongoose Bite after a dodge.",
				Rotation.plainCondition("Mongoose Bite after a dodge (Defensive State (spell 5302) opens it)."))
		end)

		it("strips an aside set off by a double dash, and keeps a clause after a colon, often the rule itself", function()
			assert.are.equal("Serpent Sting kept up whenever it is about to fall off.",
				Rotation.plainCondition("Serpent Sting kept up whenever it is about to fall off -- the cited "
					.. "guide's own rotation step, not skipped."))
			assert.are.equal("Rend kept up while the target lives.",
				Rotation.plainCondition("Rend kept up -- an aside -- while the target lives."))
			assert.are.equal("Flash Heal is the dear one: only for a tank that is genuinely low.",
				Rotation.plainCondition("Flash Heal is the dear one: only for a tank that is genuinely low. Not more."))
		end)

		it("keeps a decimal point, and ends a condition that has no end of its own", function()
			assert.are.equal("Slam inside the 1.5 second swing window.",
				Rotation.plainCondition("Slam inside the 1.5 second swing window"))
		end)

		it("answers the line count it wraps to at the measured width", function()
			local text, lines = Rotation.plainCondition("Aimed Shot on cooldown.", measure, 40)
			assert.are.equal("Aimed Shot on cooldown.", text)
			assert.are.equal(1, lines)
			text, lines = Rotation.plainCondition("Aimed Shot on cooldown, timed so its cast does not clip the next "
				.. "auto shot.", measure, 40)
			assert.are.equal("Aimed Shot on cooldown, timed so its cast does not clip the next auto shot.", text)
			assert.are.equal(2, lines)
		end)

		it("never runs past two lines: the second is cut on a word with an ellipsis", function()
			local long = "Heroic Strike only when rage is above sixty and Bloodthirst and Whirlwind are both on "
				.. "cooldown for at least a full global."
			local text, lines = Rotation.plainCondition(long, measure, 30)
			assert.are.equal(2, lines)
			assert.are.equal(L.tooltipNamePending, text:sub(-#L.tooltipNamePending))
			-- Wrapped again at the same width, the cut text still takes two lines at most.
			local again = select(2, Rotation.plainCondition(text, measure, 30))
			assert.is_true(again <= 2)
		end)

		-- L4.3 fix round 3, item 2: the cut prefers a clause boundary in the last 40% of the second
		-- line, and never leaves a figure without the word it measures.
		it("cuts at a clause boundary when one lies in the last 40% of the second line", function()
			local text, lines = Rotation.plainCondition("Slam when the swing timer resets and rage is high, or when "
				.. "Whirlwind is ready soon.", measure, 30)
			assert.are.equal(2, lines)
			assert.are.equal("Slam when the swing timer resets and rage is high" .. L.tooltipNamePending, text)
		end)

		it("cuts on a word, never ending on a connector, when no clause boundary lies in the last 40%", function()
			local text = Rotation.plainCondition("Slam when the swing timer resets, and only while rage is above "
				.. "forty, never when Whirlwind is ready.", measure, 30)
			assert.are.equal("Slam when the swing timer resets, and only while rage" .. L.tooltipNamePending, text)
		end)

		it("goes back to the last clause boundary rather than cut a figure from its word", function()
			local text = Rotation.plainCondition("Slam when the swing timer resets and rage is above 40 rage, never "
				.. "lower.", measure, 30)
			assert.are.equal("Slam when the swing timer resets" .. L.tooltipNamePending, text)
		end)

		-- L3.7 item 27 (Tranquility "...for 10..."): with no clause boundary anywhere to go back
		-- to, a cut never lands just after a figure, nor on the word that leads into it.
		it("does not end on a figure or the word before it when the line has no clause boundary", function()
			local text = Rotation.plainCondition("Use it when the party is hurt and the heal lands for 10 and sits on "
				.. "a long cooldown after that.", measure, 30)
			assert.are.equal("Use it when the party is hurt and the heal lands" .. L.tooltipNamePending, text)
			assert.is_nil(text:find("for 10", 1, true))
		end)

		it("answers the plain words whole, with no line count, when there is nothing to measure with", function()
			local text, lines = Rotation.plainCondition("Aimed Shot on cooldown (when it is up).")
			assert.are.equal("Aimed Shot on cooldown.", text)
			assert.is_nil(lines)
		end)

		it("answers nothing for no condition", function()
			assert.is_nil(Rotation.plainCondition(nil))
			assert.is_nil(Rotation.plainCondition(""))
			assert.is_nil(Rotation.plainCondition("   "))
		end)
	end)

	-- L4.3 ruling 3: the level an ability is first learned at, from Data.abilityLevels
	-- (contract 1; every rank of an ability maps to its first rank's level).
	describe("learnedAt", function()
		it("reads Data.abilityLevels by the line's own spell id", function()
			local data = { abilityLevels = { [5308] = 24, [20662] = 24, [2687] = 10 } }
			assert.are.equal(24, Rotation.learnedAt(5308, data))
			assert.are.equal(24, Rotation.learnedAt(20662, data))
			assert.are.equal(10, Rotation.learnedAt(2687, data))
		end)

		it("answers nil for a spell the data has no level for, or data without the table", function()
			assert.is_nil(Rotation.learnedAt(12345, { abilityLevels = {} }))
			assert.is_nil(Rotation.learnedAt(5308, {}))
			assert.is_nil(Rotation.learnedAt(nil, { abilityLevels = { [5308] = 24 } }))
		end)

		it("reads the shipped data: Execute at 24 in every rank", function()
			local data = require("Data")
			assert.are.equal(24, Rotation.learnedAt(5308, data))
			assert.are.equal(24, Rotation.learnedAt(20662, data))
		end)
	end)
end)

-- A rank-choice band scoped to its own lines (round 17: rank 2 at 50-59 on line 4 only) is
-- resolved by the data job, line by line: two lines of one spell in one band carry their own
-- ids, and a race's id rides only on the line its row reaches. The addon shows each line as
-- shipped and never folds one line's rank into another's.
describe("Rotation with a band scoped to one line", function()
	local SCOPED = {
		build = DATA.build,
		classes = DATA.classes,
		weights = {},
		rotations = { ["warrior-fury"] = {
			{ level = 50, lines = {
				{ spellId = 285, name = "Heroic Strike", condition = "" },
				{ spellId = 284, name = "Heroic Strike", condition = "",
					raceSpellId = { ["undead"] = 1608 } },
			} },
		} },
	}

	local function shownFor(raceToken, factionToken)
		mock.install({
			class = { name = "Warrior", token = "WARRIOR" },
			race = raceToken and { name = raceToken, token = raceToken } or nil,
			globals = { UnitFactionGroup = function() return factionToken end },
		})
		helper.load("Compat")
		helper.load("Gear")
		local Rotation = helper.load("Rotation")
		local band = Rotation.bandAt(SCOPED, "warrior-fury", 55)
		mock.uninstall()
		return { band.lines[1].spellId, band.lines[2].spellId }
	end

	it("shows each line of one spell its own rank", function()
		assert.are.same({ 285, 284 }, shownFor("Troll", "Horde"))
	end)

	it("swaps a race's id only on the line its row reaches", function()
		assert.are.same({ 285, 1608 }, shownFor("Scourge", "Horde"))
	end)
end)

-- L3.7 pre-item (round A review): a second line whose only word that fits is a figure keeps that
-- one word rather than cutting down to none.
-- A line whose rank a race's rank-choice row changes (round 13) carries each keyed
-- race's id in raceSpellId, by the site's race slug; the band shows the player's own.
describe("Rotation with a race's rank", function()
	local SPLIT = {
		build = DATA.build,
		classes = DATA.classes,
		weights = {},
		rotations = { ["warrior-fury"] = {
			{ level = 50, lines = {
				{ spellId = 285, name = "Heroic Strike", condition = "",
					raceSpellId = { ["undead"] = 1608, ["night-elf"] = 284 } },
			} },
		} },
	}

	local function shownFor(raceToken, factionToken)
		mock.install({
			class = { name = "Warrior", token = "WARRIOR" },
			race = raceToken and { name = raceToken, token = raceToken } or nil,
			globals = { UnitFactionGroup = function() return factionToken end },
		})
		helper.load("Compat")
		helper.load("Gear")
		local Rotation = helper.load("Rotation")
		local band = Rotation.bandAt(SPLIT, "warrior-fury", 55)
		mock.uninstall()
		return band.lines[1].spellId
	end

	it("shows each keyed race its own id, by the client's race token", function()
		assert.are.equal(1608, shownFor("Scourge", "Horde"))
		assert.are.equal(284, shownFor("NightElf", "Alliance"))
	end)

	it("keeps the line's own id for a race no row names, or a client that names no race", function()
		assert.are.equal(285, shownFor("Troll", "Horde"))
		assert.are.equal(285, shownFor(nil, nil))
	end)

	it("never writes the race's id into the shipped band", function()
		shownFor("Scourge", "Horde")
		assert.are.equal(285, SPLIT.rotations["warrior-fury"][1].lines[1].spellId)
	end)
end)

describe("Compat.playerRaceSlug", function()
	local function slugFor(raceToken, factionToken)
		mock.install({
			race = raceToken and { name = raceToken, token = raceToken } or nil,
			globals = { UnitFactionGroup = function() return factionToken end },
		})
		local Compat = helper.load("Compat")
		local slug = Compat.playerRaceSlug()
		mock.uninstall()
		return slug
	end

	it("names the site's slug for each client token", function()
		assert.are.equal("undead", slugFor("Scourge", "Horde"))
		assert.are.equal("night-elf", slugFor("NightElf", "Alliance"))
		assert.are.equal("windshaper-skyborne", slugFor("Skyborne", "Horde"))
		assert.are.equal("high-order-skyborne", slugFor("Skyborne", "Alliance"))
	end)

	it("is nil where the client names no race", function()
		assert.is_nil(slugFor(nil, nil))
	end)
end)

describe("Rotation.cutCount with a figure first", function()
	local Rotation = require("Rotation")

	it("keeps one word when the only word that fits is a figure at index 1", function()
		local measure = function(text)
			return #text
		end
		-- "25%" fits, "25% reserve" does not.
		assert.are.equal(1, Rotation.cutCount({ "25%", "reserve", "is", "kept" }, "x", measure, 5))
	end)
end)
