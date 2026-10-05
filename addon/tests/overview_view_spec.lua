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
			assert.are.equal(OverviewView.FIGURE_PLACEHOLDER, model.build.figure)
			assert.are.equal(L.overviewBuildNone, model.build.caption)
			assert.are.equal(L.overviewBuildNoneHint, model.build.detail)
		end)

		-- Round-2 pass: the raw "N of 17 slots filled" figure is gone
		-- entirely -- there is no plan yet to count pieces against, so
		-- the figure area shows the placeholder rather than a number
		-- that answers a question nobody asked.
		it("asks for a build before judging gear, inventing no figure to show instead", function()
			start()
			local model = OverviewView.summary(DATA)
			assert.are.equal(OverviewView.FIGURE_PLACEHOLDER, model.gear.figure)
			assert.are.equal("", model.gear.caption)
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
			-- Round-2 pass: the big figure is the spent count alone --
			-- the plan's own 11-point target lives only in the caption,
			-- never duplicated on the Talent points card.
			assert.are.equal("1", model.build.figure)
			assert.are.equal(string.format(L.overviewBuildProgress, 1, 3), model.build.caption)
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

		it("counts planned pieces already worn, reading the exact model the Gear page shows", function()
			-- Round-3 blocker fix (finding 1): this card's own upgrade
			-- count is never a second, independently-counted figure --
			-- it is #rows.upgrades, the same list the Gear page's own
			-- "UPGRADES IN YOUR BAGS" section renders (GearView.rows),
			-- which here is both the bagged Bagged Helm (delta 20) AND
			-- the already-worn Worn Helm (delta 6, which also beats the
			-- plan) -- Gear.upgrades scans what is worn too, so a slot
			-- the player already has right is never hidden, on either
			-- surface.
			start()
			assert.is_truthy(Follow.load(CODE, DATA))
			local model = OverviewView.summary(DATA)
			assert.are.equal(1, model.gear.planned)
			assert.are.equal(0, model.gear.matched)
			assert.are.equal(2, model.gear.upgrades)
			assert.are.equal(string.format(L.overviewGearUpgrades, 2), model.gear.detail)
			-- Round-2 pass: the figure is the matched-pieces count, one
			-- labelled denominator against the plan, not the raw slots
			-- a prior pass dropped from this card entirely.
			assert.are.equal("0", model.gear.figure)
			assert.are.equal(string.format(L.overviewGearMatched, 0, 1), model.gear.caption)
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

		-- design section 4.5.2: the one card about the character's own
		-- spent choices paints a class-coloured bar, not plain gold.
		it("paints each tree's bar in the player's own class colour", function()
			start()
			assert.is_truthy(Follow.load(CODE, DATA))
			local Theme = require("Theme")
			local view = OverviewView.mount(_G.CreateFrame("Frame"), {
				data = DATA, contentWidth = 538, select = function() end,
			})
			local expected = { Theme.classColor("PALADIN") }
			local painted = mock.lastCall(view.trees.rows[1].bar.foreverSixtyFill, "SetColorTexture")
			assert.are.same(expected, { painted[1], painted[2], painted[3], painted[4] })
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

	-- Round-2 owner fix: there is always something to send -- the
	-- character itself -- so the card leads with the crest, the name
	-- and the level rather than gating its own readiness on a code.
	describe("send to the site", function()
		it("leads with the character: the crest, the class-coloured name and the level", function()
			start()
			_G.UnitName = function() return "Obnoxious Yell" end
			local sync = OverviewView.summary(DATA).sync
			assert.is_true(sync.canCopy)
			assert.is_truthy(sync.code:find("FS1:", 1, true))
			assert.are.equal("Obnoxious Yell", sync.name)
			assert.are.equal("PALADIN", sync.classToken)
			assert.are.equal(string.format(L.overviewSyncLevel, 60), sync.levelLine)
			assert.is_nil(sync.refusal)
		end)

		it("is the one card that is never blank, first run or not", function()
			-- §4.7's own naming of this card: the crest, name and level
			-- come from the client the instant the character is logged
			-- in, nothing to paste for them, even with no build loaded.
			start()
			local sync = OverviewView.summary(DATA).sync
			assert.is_truthy(sync.name ~= "")
			assert.is_true(sync.canCopy)
		end)

		it("names Export.string's own genuine refusal rather than hiding it", function()
			start({ class = { name = "Death Knight", token = "DEATHKNIGHT" } })
			local sync = OverviewView.summary(DATA).sync
			assert.is_false(sync.canCopy)
			assert.is_nil(sync.code)
			assert.is_not_nil(sync.refusal)
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

		-- The clipped-banner bug (docs/tenets.md item 1): the first in-game
		-- screenshot showed Load sitting past the window's own right edge.
		-- Window.pageWidth (the real page holder's width) and
		-- Window.contentWidth (what this file lays content out in) are
		-- pinned here as their own formulas rather than by requiring
		-- Window.lua, a different lane's file this wave.
		it("keeps every banner child inside the page holder's own width", function()
			start()
			_G.ForeverSixtyInbox = { builds = { { id = "a", name = "Deep Holy", code = CODE } } }
			local S = require("Theme").SIZES
			local holderWidth = S.windowWidth - S.sidebarWidth
			local contentWidth = holderWidth - S.padding * 2
			local holder = _G.CreateFrame("Frame")
			holder:SetSize(holderWidth, S.windowHeight - S.titleBarHeight - S.headerHeight)
			local view = OverviewView.mount(holder, {
				data = DATA, contentWidth = contentWidth, select = function() end,
				refreshEverything = function() end,
			})
			assert.is_true(view.arrival:IsShown())
			local cache = {}
			for _, child in ipairs({ view.arrival, view.arrival.load, view.arrival.dismiss }) do
				assert.is_true(helper.rightEdgeWithin(child, holder, cache) >= 0,
					"a banner child's right edge must not pass the page holder's own right edge")
			end
		end)
	end)

	-- §10 ruling 2's own fix: the dead ~84px gap above the card grid on
	-- the common case (nothing pending from the companion).
	describe("OverviewView.gridTop", function()
		it("reserves the banner's own height while something is pending", function()
			local S = require("Theme").SIZES
			assert.are.equal(S.padding + OverviewView.BANNER_HEIGHT + S.cardGap, OverviewView.gridTop(true))
		end)

		it("collapses to the page's own padding with nothing pending", function()
			local S = require("Theme").SIZES
			assert.are.equal(S.padding, OverviewView.gridTop(false))
		end)

		it("moves the whole grid, the rating card and the rotation card up once the banner hides", function()
			start()
			local view = OverviewView.mount(_G.CreateFrame("Frame"), {
				data = DATA, contentWidth = 538, select = function() end,
			})
			local shown = select(5, view.build:GetPoint())
			_G.ForeverSixtyInbox = { builds = { { id = "a", name = "Deep Holy", code = CODE } } }
			view.refresh()
			local withBanner = select(5, view.build:GetPoint())
			assert.is_true(withBanner < shown, "the grid must sit lower while the banner reserves its own space")
		end)
	end)

	-- Wave-1 scope item 5: the companion inbox's own "upgrade" message.
	describe("OverviewView.upgradeLine", function()
		it("answers nil with no upgrade message waiting", function()
			start()
			assert.is_nil(OverviewView.upgradeLine(nil, "US/Ashbringer/Tester"))
			assert.is_nil(OverviewView.upgradeLine({ messages = {} }, "US/Ashbringer/Tester"))
		end)

		it("names the item, the slot, the delta and the inbox's own generated date", function()
			-- Not yet cached (Compat.displayLink's own fallback): GetItemInfo
			-- answers nil, the same convention follow_view_spec.lua's own
			-- Top Gear upgrade-row tests use for an item the client has
			-- never seen.
			start({ globals = { GetItemInfo = function() return nil end } })
			local inbox = {
				generated_at = "2026-10-01T04:00:00Z",
				messages = {
					{ type = "upgrade", slot = "chest", item_id = 250488, delta = 14 },
				},
			}
			local line = OverviewView.upgradeLine(inbox, "US/Ashbringer/Tester")
			assert.are.equal(string.format(L.inboxUpgradeLine, "item:250488", "chest", 14, "2026-10-01"), line)
		end)

		it("takes the newest message when more than one is waiting", function()
			start()
			local inbox = {
				generated_at = "2026-10-01T04:00:00Z",
				messages = {
					{ type = "upgrade", slot = "head", item_id = 1, delta = 5 },
					{ type = "upgrade", slot = "chest", item_id = 2, delta = 14 },
				},
			}
			local line = OverviewView.upgradeLine(inbox, "US/Ashbringer/Tester")
			assert.is_truthy(line:find("chest", 1, true))
		end)

		it("ignores a message addressed to a different character", function()
			start()
			local inbox = {
				generated_at = "2026-10-01T04:00:00Z",
				messages = {
					{ type = "upgrade", character = "us/ashbringer/someone-else", slot = "chest",
						item_id = 2, delta = 14 },
				},
			}
			assert.is_nil(OverviewView.upgradeLine(inbox, "US/Ashbringer/Tester"))
		end)

		it("shows nothing on the Overview page with no message waiting", function()
			start()
			local view = OverviewView.mount(_G.CreateFrame("Frame"), {
				data = DATA, contentWidth = 538, select = function() end,
			})
			assert.is_false(view.upgradeLine:IsShown())
		end)

		it("shows the line on the Overview page once a message is waiting", function()
			start()
			_G.ForeverSixtyInbox = {
				generated_at = "2026-10-01T04:00:00Z",
				messages = { { type = "upgrade", slot = "chest", item_id = 250488, delta = 14 } },
			}
			local view = OverviewView.mount(_G.CreateFrame("Frame"), {
				data = DATA, contentWidth = 538, select = function() end,
			})
			assert.is_true(view.upgradeLine:IsShown())
			assert.is_truthy(view.upgradeLine:GetText():find("chest", 1, true))
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

		-- Round-2 owner fix: the data addon IS installed, but has not
		-- rated this character yet -- the empty state names what
		-- produces a rating, with a link, rather than a bare dead end.
		it("says what produces a rating, with a link, once the data addon has no card for this character", function()
			start()
			_G.UnitName = function() return "Thoradin" end
			_G.ForeverSixtyData = {
				format = 1, generated = os.date("!%Y-%m-%dT%H:%M:%SZ"), build = "1.60.1.69893",
				characters = {}, guilds = {},
			}
			helper.load("Ratings")
			OverviewView = helper.load("OverviewView")
			local model = OverviewView.summary(DATA).rating
			assert.is_true(model.empty)
			assert.are.equal(L.overviewRatingNone, model.detail)
			assert.are.equal(L.overviewRatingSetup, model.action)
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
			assert.are.equal(string.format(L.overviewRotationHeaderStrip, "Holy", 10, DATA.build), model.header)
			assert.are.equal(1, #model.lines)
			assert.are.equal("Holy Light", model.lines[1].name)
		end)

		it("each line carries an icon, a rank and whether it was just learned", function()
			start()
			assert.is_truthy(Follow.load(CODE, DATA))
			local model = OverviewView.summary(DATA).rotation
			local line = model.lines[1]
			assert.are.equal(1, line.spellId)
			assert.is_false(line.learned)
			assert.are.equal("", line.rank) -- no GetSpellSubtext mocked
		end)

		it("glows the row Rotation.recentlyLearned names", function()
			start()
			assert.is_truthy(Follow.load(CODE, DATA))
			require("Rotation").markLearned({ { spellId = 1, name = "Holy Light" } })
			local model = OverviewView.summary(DATA).rotation
			assert.is_true(model.lines[1].learned)
		end)

		it("widens the rank with the spell id, and a cooldown when it has one, in advanced mode", function()
			start()
			assert.is_truthy(Follow.load(CODE, DATA))
			require("Prefs").setFlag("advancedDetail", true)
			_G.GetSpellCooldown = function() return 0, 8000, 1 end
			local model = OverviewView.summary(DATA).rotation
			assert.are.equal(string.format(L.overviewRotationDetailCooldown, "", 1, 8), model.lines[1].rank)
		end)

		it("is not expandable once advanced detail already shows everything", function()
			start()
			assert.is_truthy(Follow.load(CODE, DATA))
			require("Prefs").setFlag("advancedDetail", true)
			local model = OverviewView.summary(DATA).rotation
			assert.is_false(model.expandable)
		end)

		it("expands in place when the card's own state says to, without advanced detail", function()
			start()
			-- A shallow copy of DATA with a rotation five lines deep --
			-- one more than Rotation.NOVICE_LINES -- so there is
			-- something to expand into. DATA itself only ever needs one
			-- line for its other tests, per the fixture at the top of
			-- this file.
			local data = { build = DATA.build, classes = DATA.classes, weights = DATA.weights, rotations = {
				["paladin-holy"] = {
					{ level = 10, lines = {
						{ spellId = 1, name = "A" }, { spellId = 2, name = "B" },
						{ spellId = 3, name = "C" }, { spellId = 4, name = "D" },
						{ spellId = 5, name = "E" },
					} },
				},
			} }
			assert.is_truthy(Follow.load(CODE, data))
			local collapsed = OverviewView.summary(data, false).rotation
			assert.is_true(collapsed.expandable)
			assert.are.equal(4, #collapsed.lines)
			assert.are.equal(1, collapsed.moreCount)
			assert.is_false(collapsed.expanded)
			local expanded = OverviewView.summary(data, true).rotation
			assert.are.equal(5, #expanded.lines)
			assert.are.equal(0, expanded.moreCount)
			assert.is_true(expanded.expanded)
		end)
	end)
end)
