-- addon/tests/rotation_spec.lua
-- The per-level rotation lines (design section 2 item 3, section 6 Wave
-- B): ns.Data.rotations[spec] is one band per level in
-- data/pipeline/addonrotation.LEVEL_BANDS (10, 20, 30, 38, 40, 50, 60).
-- Rotation.lua picks the band for a level and slices it novice/advanced;
-- this covers that against a small hand-built rotations table, the same
-- shape addonlua.py renders into Data.lua.
local helper = require("spec_helper")
local mock = require("wow_mock")

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
end)
