-- addon/tests/rotation_spec.lua
-- The per-level rotation lines (design section 2 item 3, section 6 Wave
-- B): ns.Data.rotations[spec] is one band per level in
-- data/pipeline/addonrotation.LEVEL_BANDS (10, 20, 30, 38, 40, 50, 60).
-- Rotation.lua picks the band for a level and slices it novice/advanced;
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

	describe("model", function()
		it("is not visible with no build loaded", function()
			local model = Rotation.model(DATA, nil, RANKS, 10, false)
			assert.is_false(model.visible)
		end)

		it("picks the earliest band for a level below every rung", function()
			local model = Rotation.model(DATA, BUILD, RANKS, 5, false)
			assert.is_true(model.visible)
			assert.are.equal(10, model.level)
		end)

		it("picks the highest rung at or below the level, not the nearest", function()
			local model = Rotation.model(DATA, BUILD, RANKS, 25, false)
			assert.are.equal(20, model.level)
			model = Rotation.model(DATA, BUILD, RANKS, 60, false)
			assert.are.equal(40, model.level)
		end)

		it("trims to the novice line count and flags that there is more", function()
			local model = Rotation.model(DATA, BUILD, RANKS, 40, false)
			assert.are.equal(Rotation.NOVICE_LINES, #model.lines)
			assert.is_true(model.hasMore)
		end)

		it("shows every line in advanced mode and never flags more", function()
			local model = Rotation.model(DATA, BUILD, RANKS, 40, true)
			assert.are.equal(5, #model.lines)
			assert.is_false(model.hasMore)
		end)

		it("is not visible for a spec with no curated rotation", function()
			local model = Rotation.model(DATA, { classSlug = "warrior", order = {} }, { [1] = {} }, 10, false)
			assert.is_false(model.visible)
		end)
	end)

	describe("newAbilities", function()
		it("is empty across two levels that resolve to the same band", function()
			local fresh = Rotation.newAbilities(DATA, BUILD, RANKS, 11, 15)
			assert.are.same({}, fresh)
		end)

		it("names the line that entered at a new rung", function()
			local fresh = Rotation.newAbilities(DATA, BUILD, RANKS, 15, 20)
			assert.are.equal(0, #fresh)
			fresh = Rotation.newAbilities(DATA, BUILD, RANKS, 20, 40)
			local names = {}
			for index, entry in ipairs(fresh) do
				names[index] = entry.name
			end
			assert.are.same({ "Death Wish", "Bloodthirst", "Whirlwind" }, names)
		end)

		it("does not re-announce an ability that only ranked up (same name, new id)", function()
			-- Heroic Strike's id changes between the level-10 and level-20
			-- bands (a rank-up, the talent-point toast's own job); it must
			-- not also show up here as though it were new.
			local fresh = Rotation.newAbilities(DATA, BUILD, RANKS, 10, 20)
			for _, entry in ipairs(fresh) do
				assert.are_not.equal("Heroic Strike", entry.name)
			end
		end)

		it("is empty with no earlier level to diff against", function()
			assert.are.same({}, Rotation.newAbilities(DATA, BUILD, RANKS, nil, 40))
		end)
	end)

	-- docs/tenets.md's standard: the rotation card glows the row for the
	-- ability the player just learned, off the toast's own event.
	describe("markLearned/recentlyLearned", function()
		it("marks every fresh line's spellId", function()
			local fresh = Rotation.newAbilities(DATA, BUILD, RANKS, 20, 40)
			Rotation.markLearned(fresh)
			assert.is_true(Rotation.recentlyLearned[12328]) -- Death Wish
			assert.is_true(Rotation.recentlyLearned[23881]) -- Bloodthirst
			assert.is_nil(Rotation.recentlyLearned[2687]) -- Bloodrage: not fresh
		end)

		it("replaces the previous mark rather than accumulating it", function()
			Rotation.markLearned(Rotation.newAbilities(DATA, BUILD, RANKS, 20, 40))
			Rotation.markLearned(Rotation.newAbilities(DATA, BUILD, RANKS, 11, 15))
			assert.are.same({}, Rotation.recentlyLearned)
		end)

		it("treats a nil list the same as an empty one", function()
			Rotation.markLearned(Rotation.newAbilities(DATA, BUILD, RANKS, 20, 40))
			Rotation.markLearned(nil)
			assert.are.same({}, Rotation.recentlyLearned)
		end)
	end)

	-- L4.3 ruling 3: the character-panel tab's rotation region prints each line's condition in
	-- plain words (parentheticals stripped, cut at the first sentence end, at most two lines of
	-- the panel's width as the panel measures it).
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

	-- L4.3 fix round 1, item 4: the next band's abilities the character has not learned yet.
	describe("upcoming", function()
		local data = {
			classes = DATA.classes, rotations = DATA.rotations,
			abilityLevels = { [2687] = 10, [12328] = 30, [23881] = 40, [1680] = 36, [1608] = 8 },
		}

		local function names(found)
			local list = {}
			for index, entry in ipairs(found) do
				list[index] = entry.line.name .. "@" .. entry.at
			end
			return list
		end

		it("lists the next rung's lines not learned yet and not in the band read now, in order, at most the limit", function()
			assert.are.same({ "Death Wish@30", "Bloodthirst@40" }, names(Rotation.upcoming(data, BUILD, RANKS, 25, 2)))
			assert.are.same({ "Death Wish@30" }, names(Rotation.upcoming(data, BUILD, RANKS, 25, 1)))
			-- Death Wish learned by 32: it is the trainer's news, not this list's.
			assert.are.same({ "Bloodthirst@40", "Whirlwind@36" }, names(Rotation.upcoming(data, BUILD, RANKS, 32, 2)))
		end)

		it("is empty at the top rung, with no build, and with no level", function()
			assert.are.same({}, Rotation.upcoming(data, BUILD, RANKS, 40, 2))
			assert.are.same({}, Rotation.upcoming(data, nil, RANKS, 25, 2))
			assert.are.same({}, Rotation.upcoming(data, BUILD, RANKS, nil, 2))
		end)
	end)
end)
