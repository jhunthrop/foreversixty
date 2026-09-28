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
	rotations = {
		["paladin-holy"] = {
			{ level = 10, lines = { { spellId = 1, name = "Holy Light", condition = "The whole rotation." } } },
		},
	},
}

-- Three points in Divine Strength; head planned as item 10 with 10 strength.
-- An order is three digits a point: tab, tier, column.
local CODE = "FSB1:1.60.1.69893:paladin:111111111:head=10:strength=10"

local function install(overrides)
	local state = {
		class = { name = "Paladin", token = "PALADIN" },
		race = { name = "Human", token = "Human" },
		realm = "Ashbringer",
		region = 1,
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
			-- The card's own code box is invisible, so the line under the title
			-- has to prove the code is there: its head, and how long it is.
			assert.are.equal(OverviewView.preview(sync.code), sync.detail)
			assert.is_truthy(sync.detail:find("^FS1:"))
			assert.is_truthy(sync.detail:find(string.format("(%d characters)", #sync.code), 1, true))
		end)
	end)

	it("shows the player's own rating on the sync card when the data addon has it", function()
		start()
		_G.UnitName = function() return "Thoradin" end
		_G.ForeverSixtyData = {
			format = 1, generated = os.date("!%Y-%m-%dT%H:%M:%SZ"), build = "1.60.1.69893",
			characters = { ["us:ashbringer:thoradin"] = { rating = 73, fights = 11 } }, guilds = {},
		}
		helper.load("Ratings")
		OverviewView = helper.load("OverviewView")
		assert.are.equal(string.format(L.overviewRating, 73, 11), OverviewView.summary(DATA).sync.progress)
		_G.ForeverSixtyData = nil
		helper.load("Ratings")
		OverviewView = helper.load("OverviewView")
		assert.is_nil(OverviewView.summary(DATA).sync.progress)
	end)

	it("mounts four cards and refreshes without error", function()
		start()
		local view = OverviewView.mount(_G.CreateFrame("Frame"), {
			data = DATA, contentWidth = 538, select = function() end,
		})
		assert.are.equal(4, #view.cards)
		view.refresh()
	end)

	-- The "build arrived" banner (design section 1).
	describe("the build-arrived banner", function()
		it("is hidden with nothing waiting in the inbox", function()
			start()
			assert.is_false(OverviewView.summary(DATA).arrival.visible)
		end)

		it("names the character and summarises the diff against nothing loaded", function()
			start()
			_G.ForeverSixtyInbox = { builds = { { id = "a", name = "Deep Holy", code = CODE } } }
			local model = OverviewView.summary(DATA).arrival
			assert.is_true(model.visible)
			assert.are.equal(string.format(L.buildArrivedTitle, "Tester"), model.title)
			assert.are.equal(L.buildArrivedFirst, model.diff)
		end)

		it("stops showing once the pending build is dismissed", function()
			start()
			_G.ForeverSixtyInbox = { builds = { { id = "a", name = "Deep Holy", code = CODE } } }
			assert.is_true(OverviewView.summary(DATA).arrival.visible)
			Follow.dismissInbox("a")
			assert.is_false(OverviewView.summary(DATA).arrival.visible)
		end)

		it("Load it loads the build, dismisses the entry and hides the banner", function()
			start()
			_G.ForeverSixtyInbox = { builds = { { id = "a", name = "Deep Holy", code = CODE } } }
			local refreshed = false
			local view = OverviewView.mount(_G.CreateFrame("Frame"), {
				data = DATA, contentWidth = 538, select = function() end,
				refreshEverything = function() refreshed = true end,
			})
			assert.is_true(view.arrival:IsShown())
			view.arrival.load:GetScript("OnClick")(view.arrival.load)
			assert.is_not_nil(Follow.build)
			assert.is_true(refreshed)
			assert.is_false(view.arrival:IsShown())
		end)

		it("Dismiss closes the banner without loading the build", function()
			start()
			_G.ForeverSixtyInbox = { builds = { { id = "a", name = "Deep Holy", code = CODE } } }
			local view = OverviewView.mount(_G.CreateFrame("Frame"), {
				data = DATA, contentWidth = 538, select = function() end,
				refreshEverything = function() end,
			})
			view.arrival.dismiss:GetScript("OnClick")(view.arrival.dismiss)
			assert.is_nil(Follow.build)
			assert.is_false(view.arrival:IsShown())
		end)
	end)

	-- The personal rating card (design section 3): always shows, its
	-- detail widened by the advanced-detail toggle (section 4).
	describe("the personal rating card", function()
		after_each(function()
			_G.ForeverSixtyData = nil
		end)

		it("shows the not-installed reason when the data addon is absent", function()
			start()
			local model = OverviewView.summary(DATA).rating
			assert.is_true(model.empty)
			assert.are.equal(L.ratingsNotInstalled, model.detail)
		end)

		it("shows the headline rating in novice mode, without the component breakdown", function()
			start()
			_G.UnitName = function() return "Thoradin" end
			_G.ForeverSixtyData = {
				format = 1, generated = os.date("!%Y-%m-%dT%H:%M:%SZ"), build = "1.60.1.69893",
				characters = { ["us:ashbringer:thoradin"] = { rating = 81, output = 90, fights = 24 } },
				guilds = {},
			}
			helper.load("Ratings")
			OverviewView = helper.load("OverviewView")
			local model = OverviewView.summary(DATA).rating
			assert.is_false(model.empty)
			assert.are.equal(string.format(L.overviewRatingHeadline, 81, 24), model.title)
			assert.is_nil(model.detail:find(L.ratingsOutput, 1, true))
		end)

		it("adds the top components once advanced detail is on", function()
			start()
			require("Prefs").setFlag("advancedDetail", true)
			_G.UnitName = function() return "Thoradin" end
			_G.ForeverSixtyData = {
				format = 1, generated = os.date("!%Y-%m-%dT%H:%M:%SZ"), build = "1.60.1.69893",
				characters = { ["us:ashbringer:thoradin"] = { rating = 81, output = 90, fights = 24 } },
				guilds = {},
			}
			helper.load("Ratings")
			OverviewView = helper.load("OverviewView")
			local model = OverviewView.summary(DATA).rating
			assert.is_truthy(model.detail:find(L.ratingsOutput, 1, true))
			assert.is_truthy(model.detail:find("90", 1, true))
		end)
	end)

	-- The rotation card (design section 2 item 3 / section 6 Wave B).
	describe("the rotation card", function()
		it("asks for a build when none is loaded", function()
			start()
			local model = OverviewView.summary(DATA).rotation
			assert.is_true(model.empty)
			assert.are.equal(L.overviewRotationNoBuild, model.reason)
		end)

		it("shows the current level's lines once a build is loaded", function()
			start()
			assert.is_truthy(Follow.load(CODE, DATA))
			local model = OverviewView.summary(DATA).rotation
			assert.is_false(model.empty)
			assert.are.equal(string.format(L.overviewRotationTitle, 10), model.title)
			assert.are.equal(1, #model.lines)
			assert.are.equal("Holy Light", model.lines[1].name)
		end)
	end)
end)
