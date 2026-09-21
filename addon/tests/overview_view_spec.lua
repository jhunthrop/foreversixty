-- addon/tests/overview_view_spec.lua
-- The landing page. Its model is the platform's value in one glance: where
-- the build stands and what to take next, what the gear needs, where the
-- talent points went, and whether the character has been sent to the site.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local DATA = {
	build = "1.60.1.69893",
	classes = { paladin = { tabs = {
		{ name = "Holy", talents = { { name = "Divine Strength", tier = 1, column = 1, maxRank = 5 } } },
		{ name = "Protection", talents = {} },
		{ name = "Retribution", talents = {} },
	} } },
	weights = { ["paladin-holy"] = { strength = 1.0, stamina = 0.5 } },
}

-- Three points in Divine Strength; head planned as item 10 with 10 strength.
-- An order is three digits a point: tab, tier, column.
local CODE = "FSB1:1.60.1.69893:paladin:111111111:head=10:strength=10"

local function install(overrides)
	local state = {
		class = { name = "Paladin", token = "PALADIN" },
		race = { name = "Human", token = "Human" },
		talents = {
			{ name = "Holy", talents = { { name = "Divine Strength", tier = 1, column = 1, rank = 1, maxRank = 5 } } },
			{ name = "Protection", talents = {} },
			{ name = "Retribution", talents = {} },
		},
		equipped = { [1] = "|Hitem:11|h" },
		bags = { [0] = { "|Hitem:12|h" } },
		itemStats = {
			["|Hitem:11|h"] = {
				__itemId = 11, __slot = "INVTYPE_HEAD", __name = "Worn Helm", __quality = 2,
				ITEM_MOD_STRENGTH_SHORT = 16,
			},
			["|Hitem:12|h"] = {
				__itemId = 12, __slot = "INVTYPE_HEAD", __name = "Bagged Helm", __quality = 3,
				ITEM_MOD_STRENGTH_SHORT = 30,
			},
		},
	}
	for key, value in pairs(overrides or {}) do
		state[key] = value
	end
	return mock.install(state)
end

describe("OverviewView", function()
	local Follow, OverviewView

	local function start(overrides)
		install(overrides)
		helper.load("Theme").reset()
		helper.load("Widgets")
		helper.load("Cards")
		helper.load("Export")
		helper.load("Gear")
		Follow = helper.load("Follow")
		helper.load("ExportView")
		helper.load("FollowView")
		helper.load("GearView")
		OverviewView = helper.load("OverviewView")
	end

	after_each(function()
		mock.uninstall()
	end)

	describe("with no build loaded", function()
		it("says so and points at loading one, without inventing progress", function()
			start()
			local model = OverviewView.summary(DATA)
			assert.is_false(model.build.loaded)
			assert.are.equal(0, model.build.fraction)
			assert.are.equal(L.overviewBuildNone, model.build.title)
			assert.are.equal(L.overviewBuildNoneHint, model.build.detail)
		end)

		it("still counts what is worn, and asks for a build before judging it", function()
			start()
			local model = OverviewView.summary(DATA)
			assert.are.equal(string.format(L.overviewGearSlots, 1, 17), model.gear.title)
			assert.are.equal(L.overviewGearNeedsBuild, model.gear.detail)
			assert.are.equal(0, model.gear.upgrades)
		end)
	end)

	describe("with a build loaded", function()
		it("reports progress as points taken out of points planned", function()
			start()
			assert.is_truthy(Follow.load(CODE, DATA))
			local model = OverviewView.summary(DATA)
			assert.is_true(model.build.loaded)
			assert.are.equal(1, model.build.spent)
			assert.are.equal(3, model.build.total)
			assert.is_true(math.abs(model.build.fraction - 1 / 3) < 1e-9)
			assert.is_truthy(model.build.detail:find("Divine Strength", 1, true))
		end)

		it("says the build is complete when every planned point is taken", function()
			start({ talents = {
				{ name = "Holy", talents = { { name = "Divine Strength", tier = 1, column = 1, rank = 3, maxRank = 5 } } },
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			} })
			assert.is_truthy(Follow.load(CODE, DATA))
			local model = OverviewView.summary(DATA)
			assert.are.equal(1, model.build.fraction)
			assert.are.equal(L.trackerComplete, model.build.detail)
		end)

		it("counts planned pieces already worn and upgrades waiting in the bags", function()
			start()
			assert.is_truthy(Follow.load(CODE, DATA))
			local model = OverviewView.summary(DATA)
			assert.are.equal(1, model.gear.planned)
			assert.are.equal(0, model.gear.matched)
			assert.are.equal(1, model.gear.upgrades)
			assert.are.equal(string.format(L.overviewGearUpgrades, 1), model.gear.detail)
		end)
	end)

	describe("talent points", function()
		it("gives each tree its points and its share of the biggest tree", function()
			start()
			local trees = OverviewView.summary(DATA).trees
			assert.are.equal(3, #trees)
			assert.are.equal("Holy", trees[1].name)
			assert.are.equal(1, trees[1].points)
			assert.are.equal(1, trees[1].fraction)
			assert.are.equal(0, trees[2].fraction)
		end)

		it("draws empty bars, not a division by zero, for a character with no points", function()
			start({ talents = {
				{ name = "Holy", talents = { { name = "Divine Strength", tier = 1, column = 1, rank = 0, maxRank = 5 } } },
				{ name = "Protection", talents = {} },
				{ name = "Retribution", talents = {} },
			} })
			for _, tree in ipairs(OverviewView.summary(DATA).trees) do
				assert.are.equal(0, tree.fraction)
			end
		end)
	end)

	describe("send to the site", function()
		it("offers the copy when there is a code, with when it was last saved", function()
			start()
			local sync = OverviewView.summary(DATA).sync
			assert.is_true(sync.canCopy)
			assert.is_truthy(sync.code:find("FS1:", 1, true))
			assert.are.equal(string.format(L.exportSavedAt, L.exportNotYet), sync.detail)
		end)
	end)

	it("mounts four cards and refreshes without error", function()
		start()
		local view = OverviewView.mount(_G.CreateFrame("Frame"), {
			data = DATA, contentWidth = 538, select = function() end,
		})
		assert.are.equal(4, #view.cards)
		view.refresh()
	end)
end)
