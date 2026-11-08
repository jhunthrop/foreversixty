-- addon/tests/panel_spec.lua
-- The addon panel (design/specs/2026-11-08-addon-ux.md section 4): the pure model's regions in
-- priority order (Tonight or Last raid, Upgrades, Talents or Pick your path, Rotation), the
-- Tonight answer through the companion outbox behind its feature check, the height budget's
-- shedding, and the frame drawn through the client mock without writing into a client frame.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local RaidAssign = require("RaidAssign")
local START = RaidAssign.epochOf("2026-12-10T02:00:00Z")

local function tonightInput(overrides)
	local base = {
		tonight = { raid_id = 42, raid = "Onyxia's Lair", start = START,
			jobs = { { label = "Phase 2 ranged, west side", purpose = "Keeps the whelps off the healers.",
				with = "Thornhide" } } },
		clock = START - (3 * 3600 + 12 * 60), clockText = "8:00 pm", who = "Ashvale",
	}
	for key, value in pairs(overrides or {}) do
		base[key] = value
	end
	return base
end

describe("Panel, the pure model", function()
	local Panel

	before_each(function()
		mock.install({})
		helper.load("Theme").reset()
		Panel = helper.load("Panel")
	end)

	after_each(function()
		mock.uninstall()
	end)

	describe("Tonight", function()
		it("shows from 18 hours before the start until six hours after, with the start and a countdown", function()
			local region = Panel.tonight(tonightInput())
			assert.are.equal("tonight", region.kind)
			assert.are.equal("Onyxia's Lair", region.raid)
			assert.are.equal(string.format(L.panelTonightNote, "8:00 pm", "in 3 h 12 min"), region.note)
			assert.is_nil(Panel.tonight(tonightInput({ clock = START - Panel.TONIGHT_SECONDS - 1 })))
			assert.is_nil(Panel.tonight(tonightInput({ clock = START + RaidAssign.SHOWN_FOR_SECONDS + 1 })))
			assert.are.equal(L.panelTonightNow, Panel.countdown(START, START + 60))
			assert.are.equal("in 25 min", Panel.countdown(START, START - 25 * 60))
		end)

		it("reads On the roster from the assignment alone while the contracts are missing, with no buttons", function()
			local region = Panel.tonight(tonightInput())
			assert.are.equal(L.panelOnRoster, region.answer)
			assert.is_nil(region.buttons)
			assert.are.equal(string.format(L.panelAs, "Ashvale"), region.who)
			assert.are.equal("Phase 2 ranged, west side", region.job)
			assert.are.equal(string.format(L.panelWith, "Thornhide"), region.with)
		end)

		it("offers In, Late, Bench and Out to an unanswered raid only when the companion takes answers", function()
			local signup = { raid_id = 42, answer = "none" }
			local region = Panel.tonight(tonightInput({ signup = signup, accepts = true }))
			assert.are.equal(L.panelAnswerNone, region.answer)
			assert.are.equal("clientGold", region.answerColor)
			assert.are.same({ "in", "late", "bench", "out" }, region.buttons)
			assert.is_nil(Panel.tonight(tonightInput({ signup = signup, accepts = false })).buttons)
		end)

		it("names the answer the site holds, and the role", function()
			local cases = { ["in"] = { L.panelAnswerIn, "tooltipGreen" }, late = { L.panelAnswerLate, "clientGold" },
				bench = { L.panelAnswerBench, "tooltipBody" }, out = { L.panelAnswerOut, "tooltipMuted" } }
			for answer, expected in pairs(cases) do
				local region = Panel.tonight(tonightInput({ signup = { raid_id = 42, answer = answer, role = "dps" },
					accepts = true }))
				assert.are.same(expected, { region.answer, region.answerColor }, answer)
				assert.is_nil(region.buttons)
				assert.are.equal(string.format(L.panelAsRole, "Ashvale", L.panelRoleDamage), region.who)
			end
		end)

		it("reads In, not sent yet until the inbox confirms an answer at or after the click", function()
			local pending = { raid_id = 42, answer = "in", answered_at = START - 4 * 3600 }
			local region = Panel.tonight(tonightInput({ signup = { raid_id = 42, answer = "none" }, accepts = true,
				pending = pending }))
			assert.are.equal(string.format(L.panelAnswerPending, L.panelButtonIn), region.answer)
			assert.are.equal("In, not sent yet", region.answer)
			assert.are.equal(L.panelSendsLater, region.sending)
			assert.is_nil(region.buttons)
			local confirmed = Panel.tonight(tonightInput({ accepts = true, pending = pending,
				signup = { raid_id = 42, answer = "in", answered_at = "2026-12-09T22:00:00Z" } }))
			assert.are.equal(L.panelAnswerIn, confirmed.answer)
			assert.is_nil(confirmed.sending)
		end)

		it("names one or two partners and none past that", function()
			local three = tonightInput()
			three.tonight.jobs[1].with = "A, B, C"
			assert.is_nil(Panel.tonight(three).with)
			local two = tonightInput()
			two.tonight.jobs[1].with = "A, B"
			assert.are.equal(string.format(L.panelWith, "A and B"), Panel.tonight(two).with)
		end)

		it("reads the companion's accepts list as the feature check", function()
			assert.is_true(Panel.acceptsAnswers({ accepts = { "raid_answer" } }))
			assert.is_false(Panel.acceptsAnswers({ accepts = {} }))
			assert.is_false(Panel.acceptsAnswers({}))
			assert.is_false(Panel.acceptsAnswers(nil))
		end)
	end)

	describe("the raid answer in the outbox", function()
		it("keeps the newest answer per raid and character and drops ones older than a week", function()
			local Export = require("Export")
			local now = START
			local saved = { { character = "Bob-Ash", raid_id = 42, answer = "late", answered_at = now - 100 },
				{ character = "Bob-Ash", raid_id = 7, answer = "in", answered_at = now - 8 * 86400 } }
			local added = { { character = "Bob-Ash", raid_id = 42, answer = "in", answered_at = now - 10 } }
			local list = Export.mergeRaidAnswers(saved, added, now)
			assert.are.equal(1, #list)
			assert.are.equal("in", list[1].answer)
			assert.are.equal(7 * 86400, Export.RAID_ANSWER_KEEP_SECONDS)
		end)

		it("writes raid_answer into the outbox the contract shapes, and leaves it out with none", function()
			local Export = require("Export")
			local entry = { character = "Bob-Ash", raid_id = 42, answer = "in", answered_at = 5 }
			local outbox = Export.outbox({ character = "Bob-Ash", code = "FS1:x", wallNow = 0, raidAnswers = { entry } })
			assert.are.same({ entry }, outbox.raid_answer)
			assert.is_nil(Export.outbox({ character = "Bob-Ash", code = "FS1:x", wallNow = 0, raidAnswers = {} }).raid_answer)
		end)
	end)

	describe("Upgrades", function()
		local function target(itemId, gain, extra)
			local entry = { itemId = itemId, name = "Item " .. itemId, hex = "1eff00", gain = gain, unit = "dps",
				slotWord = "Chest", source = "Wailing Caverns" }
			for key, value in pairs(extra or {}) do
				entry[key] = value
			end
			return entry
		end

		it("lists at most three, the slot and source under each, and the level an item asks for", function()
			local region = Panel.upgrades({ level = 24, targets = {
				target(1, 6.4), target(2, 4.1, { minLevel = 27 }), target(3, 3), target(4, 2) } })
			assert.are.equal(3, #region.rows)
			assert.are.equal(string.format(L.panelTargetDetail, "Chest", "Wailing Caverns"), region.rows[1].detail)
			assert.are.equal(string.format(L.panelTargetDetail, "Chest", "Wailing Caverns")
				.. string.format(L.panelAtLevel, 27), region.rows[2].detail)
			assert.is_nil(region.rest)
		end)

		it("says the other slots are best under fewer rows, every slot best with none, and where picks start", function()
			assert.are.equal(L.panelOthersBest, Panel.upgrades({ targets = { target(1, 2) } }).rest)
			assert.are.equal(L.panelAllBest, Panel.upgrades({ targets = {} }).rest)
			assert.are.equal(string.format(L.homePicksStart, 20), Panel.upgrades({ targets = {}, picksStart = 20 }).rest)
		end)

		it("carries the sim note and Sim again only once the character has a sim", function()
			assert.is_nil(Panel.upgrades({ targets = {} }).note)
			local region = Panel.upgrades({ targets = {}, sim = { age = "2 days ago" } })
			assert.are.equal(string.format(L.panelSimNote, "2 days ago"), region.note)
			assert.are.equal(L.panelSimAgain, region.simAgain)
			assert.are.equal(L.homeSimQueued, Panel.upgrades({ targets = {}, sim = { queued = true } }).simAgain)
		end)
	end)

	describe("Talents and Pick your path", function()
		local BUILD = { classSlug = "warrior", order = { {}, {}, {} } }

		it("names the next point's level and talent, its rank, the plan count and the off-plan count", function()
			local region = Panel.talents({ build = BUILD, level = 24, unspent = 0, buildName = "Arms leveling",
				pointName = "Lethal Attacks", pointRank = { have = 0, want = 5 },
				progress = { matched = 2, total = 3, offPlan = 1, nextLevel = 25, point = {} } })
			assert.are.equal(string.format(L.panelTalentAt, 25, "Lethal Attacks"), region.title)
			assert.are.equal("1 of 5", region.rank)
			assert.are.equal(string.format(L.panelOffPlanCount, 1), region.offPlan)
			assert.are.equal("Arms leveling", region.buildName)
		end)

		it("says Spend N now with points waiting", function()
			local region = Panel.talents({ build = BUILD, level = 24, unspent = 2, pointName = "X",
				progress = { matched = 0, total = 3, offPlan = 0, point = {} } })
			assert.are.equal(string.format(L.panelSpendNow, 2), region.spend)
		end)

		it("is absent for a finished build on plan; the heading and off-plan note alone off plan", function()
			assert.is_nil(Panel.talents({ build = BUILD, level = 60,
				progress = { matched = 3, total = 3, offPlan = 0 } }))
			local off = Panel.talents({ build = BUILD, level = 60, progress = { matched = 3, total = 3, offPlan = 2 } })
			assert.is_true(off.finished)
		end)

		it("offers Pick your path under 10 with no build, and nothing otherwise", function()
			local specs = { { slug = "warrior-arms", name = "Arms", playstyle = "Big hits" } }
			assert.are.equal("path", Panel.path({ level = 5, specs = specs }).kind)
			assert.is_nil(Panel.path({ level = 12, specs = specs }))
			assert.is_nil(Panel.path({ level = 5, build = BUILD, specs = specs }))
		end)
	end)

	describe("Rotation", function()
		local function rows(count, upcoming)
			local list = {}
			for index = 1, count do
				list[index] = { spellId = index, name = "Spell " .. index,
					fullCondition = "Spell " .. index .. " when ready; never at low mana." }
			end
			if upcoming then
				list[#list + 1] = { spellId = 99, name = "Rapid Fire", muted = true, atLevel = 26 }
			end
			return list
		end

		it("shows every line of the band under 60 (six at most), first clauses, and the next ability dimmed", function()
			local region = Panel.rotation({ level = 24, rows = rows(8, true) })
			assert.are.equal(Panel.MAX_ROTATION, #region.rows)
			assert.are.equal("Spell 1 when ready", region.rows[1].clause)
			assert.are.equal(string.format(L.panelAtLevelRow, 26), region.nextRow.clause)
			assert.is_nil(region.collapsed)
		end)

		it("collapses at 60 unless opened, or a new rank was learned", function()
			assert.is_true(Panel.rotation({ level = 60, rows = rows(3) }).collapsed)
			assert.is_false(Panel.rotation({ level = 60, rows = rows(3), expanded = true }).collapsed)
			assert.is_false(Panel.rotation({ level = 60, rows = rows(3), newRank = true }).collapsed)
			assert.is_nil(Panel.rotation({ level = 24, rows = {} }))
		end)
	end)

	describe("the model and its height", function()
		it("orders Tonight, Upgrades, Talents, Rotation and never passes the budget", function()
			local input = tonightInput({ level = 24, targets = {},
				build = { classSlug = "warrior", order = { {}, {} } }, unspent = 0, pointName = "X",
				progress = { matched = 1, total = 2, offPlan = 0, nextLevel = 25, point = {} },
				rows = { { spellId = 1, name = "A", fullCondition = "A." } } })
			local model = Panel.model(input)
			local kinds = {}
			for index, region in ipairs(model.regions) do
				kinds[index] = region.kind
			end
			assert.are.same({ "tonight", "upgrades", "talents", "rotation" }, kinds)
			assert.is_true(model.height <= Panel.MAX_HEIGHT)
		end)

		it("sheds the rotation to four rows, then the upgrades to two, then the next-ability row", function()
			local regions = {
				{ kind = "upgrades", rows = { {}, {}, {} }, rest = nil },
				{ kind = "rotation", rows = { {}, {}, {}, {}, {}, {} }, nextRow = {} },
			}
			local model = Panel.fit({ head = {}, regions = regions }, 200)
			assert.are.equal(Panel.SHED_ROTATION, #regions[2].rows)
			assert.are.equal(Panel.SHED_UPGRADES, #regions[1].rows)
			assert.is_nil(regions[2].nextRow)
			assert.is_true(model.height <= 200)
		end)
	end)

	it("colours a percentile as the log sites do", function()
		assert.are.equal("e5cc80", Panel.parseHex(100))
		assert.are.equal("ff8000", Panel.parseHex(96))
		assert.are.equal("0070ff", Panel.parseHex(71))
		assert.are.equal("9d9d9d", Panel.parseHex(10))
	end)
end)

describe("Panel, the frame", function()
	local Panel, state

	before_each(function()
		state = mock.install({ class = { name = "Hunter", token = "HUNTER" }, level = 24, playerName = "Ashvale" })
		helper.load("Theme").reset()
		Panel = helper.load("Panel")
	end)

	after_each(function()
		mock.uninstall()
	end)

	it("draws a full model on an addon frame, writing nothing into a client frame", function()
		local model = Panel.model(tonightInput({ level = 24, classToken = "HUNTER", fact = "Level 24 Hunter",
			signup = { raid_id = 42, answer = "none" }, accepts = true,
			targets = { { itemId = 1, name = "Bow", hex = "0070dd", gain = 3.5, unit = "dps", slotWord = "Ranged",
				source = "World drop" }, { itemId = 2, name = "Gloves", gain = 1.2, unit = "dps", bag = true, link = "x" } },
			build = { classSlug = "hunter", order = { {}, {} } }, unspent = 1, pointName = "Lethal Attacks",
			progress = { matched = 1, total = 2, offPlan = 1, nextLevel = 25, point = {} },
			rows = { { spellId = 1, name = "Aimed Shot", fullCondition = "Aimed Shot when ready." } },
			score = { score = 1240, tierColor = { r = 0, g = 0.44, b = 0.87 } } }))
		assert.has_no.errors(function()
			Panel.draw(model)
		end)
		assert.are.equal(Panel.WIDTH, Panel.frame:GetWidth())
		assert.is_true(Panel.frame:GetHeight() <= Panel.MAX_HEIGHT)
		assert.are.same({}, state.clientWrites)
	end)

	it("never opens in combat, and closes when combat starts", function()
		require("Theme").inCombat = function() return true end
		assert.is_false(Panel.open("chip"))
		require("Theme").inCombat = function() return false end
		Panel.ensure()
		Panel.frame:Show()
		Panel.onCombatStart()
		assert.is_false(Panel.frame:IsShown())
	end)

	it("goes back to its saved spot, which Reset positions forgets", function()
		Panel.ensure()
		_G.ForeverSixtyDB = { panel = { point = "TOPLEFT", x = 40, y = -60 } }
		assert.are.equal("saved", Panel.place("chip"))
		Panel.resetPosition()
		assert.is_nil(_G.ForeverSixtyDB.panel)
		assert.are.equal("center", Panel.place("key"))
		_G.ForeverSixtyDB = nil
	end)

	it("queues an answer only when the companion takes answers", function()
		_G.ForeverSixtyInbox = nil
		assert.is_false(Panel.answer(42, "in"))
		assert.is_false(Panel.answer(42, "maybe"))
		_G.ForeverSixtyInbox = nil
	end)
end)
