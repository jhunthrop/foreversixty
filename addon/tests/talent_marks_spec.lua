-- addon/tests/talent_marks_spec.lua
-- The talent marks (spec section 4.2): the strip's line and each badge's
-- state, the two pure rules the frames read. Pure: hand-built tables only,
-- no client mock and no frame.
local TalentMarks = require("TalentMarks")
local L = require("Locale")
local Theme = require("Theme")

local function progress(matched, total, nextLevel, offPlan, done)
	return { matched = matched, total = total, nextLevel = nextLevel, offPlan = offPlan,
		point = (not done) and { index = matched + 1 } or nil, done = done == true }
end

-- Fix round 2 (personas on marks2-*): "next at level 21" at level 21 with
-- twelve points to spend read as "later"; the strip says "spend 12 now"
-- while points wait, and "off build" is muted at 0 and a warning from 1.
local MUTED, WARN = "|cff9a9484", "|cffff6b5c"

describe("TalentMarks.stripText", function()
	it("says spend N now while points wait, with no level", function()
		local text, warned = TalentMarks.stripText("Arms / Fury build", progress(0, 51, 21, 0), { unspent = 12 })
		assert.are.equal("Arms / Fury build · 0 of 51 points · spend 12 now · " .. MUTED .. "0 off build|r", text)
		assert.is_false(warned)
	end)

	it("names the level the next point arrives at when none waits", function()
		local text = TalentMarks.stripText("Arms leveling", progress(11, 51, 22, 0), { unspent = 0 })
		assert.are.equal("Arms leveling · 11 of 51 points · next point at level 22 · " .. MUTED .. "0 off build|r",
			text)
	end)

	it("colours the off-build count as a warning from one point", function()
		local text, warned = TalentMarks.stripText("Arms leveling", progress(11, 51, 22, 1), { unspent = 0 })
		assert.are.equal("Arms leveling · 11 of 51 points · next point at level 22 · " .. WARN .. "1 off build|r",
			text)
		assert.is_true(warned)
	end)

	it("leaves the level out where no level follows (the level cap)", function()
		local text = TalentMarks.stripText("Arms leveling", progress(49, 51, nil, 2), { unspent = 0 })
		assert.are.equal("Arms leveling · 49 of 51 points · " .. WARN .. "2 off build|r", text)
	end)

	it("says complete once the build has no point left, spendable points or not", function()
		local text = TalentMarks.stripText("Arms leveling", progress(51, 51, nil, 0, true), { unspent = 3 })
		assert.are.equal("Arms leveling · 51 of 51 points · complete · " .. MUTED .. "0 off build|r", text)
	end)

	it("asks for Apply Changes while an on-build point is staged, not while one is off build", function()
		local staged = TalentMarks.stripText("Arms leveling", progress(12, 51, 21, 0), { unspent = 0, staged = true })
		assert.are.equal("Arms leveling · 12 of 51 points · next point at level 21 · " .. MUTED
			.. "0 off build|r · click Apply Changes", staged)
		local off = TalentMarks.stripText("Arms leveling", progress(11, 51, 21, 1), { unspent = 0, staged = true })
		assert.are.equal("Arms leveling · 11 of 51 points · next point at level 21 · " .. WARN .. "1 off build|r", off)
	end)

	it("has nothing to say without a build", function()
		assert.is_nil(TalentMarks.stripText(nil, progress(0, 0, nil, 0), {}))
	end)

	it("keeps its wording in Locale", function()
		assert.are.equal("spend %d now", L.talentStripSpend)
		assert.are.equal("next point at level %d", L.talentStripNext)
		assert.are.equal("%d of %d points", L.talentStripCount)
		assert.are.equal("%d off build", L.talentStripOffBuild)
		assert.are.equal("click Apply Changes", L.talentStripApply)
	end)
end)

-- Fix round 2: an unnamed build is named after the plan's tree split.
describe("Follow.buildName", function()
	local Follow = require("Follow")
	local DATA = { classes = { warrior = { tabs = { { name = "Arms" }, { name = "Fury" }, { name = "Protection" } } } } }
	local function build(counts, name)
		local order = {}
		for tab, count in ipairs(counts) do
			for _ = 1, count do
				order[#order + 1] = { tab = tab, tier = 1, column = 1 }
			end
		end
		return { name = name, classSlug = "warrior", order = order }
	end

	it("keeps a build's own name", function()
		assert.are.equal("Arms leveling", Follow.buildName(DATA, build({ 31, 20 }, "Arms leveling")))
	end)

	it("names an unnamed build after its two biggest trees, biggest first", function()
		assert.are.equal("Arms / Fury build", Follow.buildName(DATA, build({ 31, 20 })))
		assert.are.equal("Fury / Arms build", Follow.buildName(DATA, build({ 15, 31, 5 })))
		assert.are.equal("Protection build", Follow.buildName(DATA, build({ 0, 0, 51 })))
	end)

	it("falls back to the class when the trees cannot be named", function()
		assert.are.equal("Warrior build", Follow.buildName(nil, build({ 31, 20 })))
		assert.are.equal("Warrior build", Follow.buildName(DATA, build({})))
	end)
end)

describe("TalentMarks.badgeState", function()
	it("is next for the cell the ring is on", function()
		assert.are.equal("next", TalentMarks.badgeState({ want = 5, have = 3, next = true }))
	end)

	it("is taken once the character has every rank the plan wants there, or more", function()
		assert.are.equal("taken", TalentMarks.badgeState({ want = 2, have = 2, next = false }))
		assert.are.equal("taken", TalentMarks.badgeState({ want = 2, have = 3, next = false }))
	end)

	it("is planned for a cell still to come", function()
		assert.are.equal("planned", TalentMarks.badgeState({ want = 3, have = 0, next = false }))
		assert.are.equal("planned", TalentMarks.badgeState({ want = 3, have = 1 }))
	end)
end)

describe("Theme.inlineColor", function()
	it("wraps text in the client's colour escape for a token", function()
		assert.are.equal("|cffff6b5cx|r", Theme.inlineColor("warning", "x"))
		assert.are.equal("|cffe5b955y|r", Theme.inlineColor("gold", "y"))
	end)
end)

describe("Theme.SIZES for the marks", function()
	it("has the 16 px plate with 12 px digits and the 24 px strip", function()
		assert.are.equal(16, Theme.SIZES.orderBadge)
		assert.are.equal(12, Theme.SIZES.orderBadgeText)
		assert.are.equal(24, Theme.SIZES.talentStrip)
	end)
end)

-- Fix round 1 (seen in game): on the window's first OnShow after a reload
-- the trait buttons do not exist yet, so every lookup missed and the
-- session latch turned the badges off; the second open found them all.
describe("TalentMarks.latchDecision", function()
	it("is ok when every badge found its button, or there are none", function()
		assert.are.equal("ok", TalentMarks.latchDecision(0, 18, false))
		assert.are.equal("ok", TalentMarks.latchDecision(0, 18, true))
		assert.are.equal("ok", TalentMarks.latchDecision(0, 0, true))
	end)

	it("waits, never latching, when every node is missing: the buttons are not built yet", function()
		assert.are.equal("wait", TalentMarks.latchDecision(18, 18, false))
		assert.are.equal("wait", TalentMarks.latchDecision(18, 18, true))
	end)

	it("retries a partial miss until the deferred pass is done", function()
		assert.are.equal("retry", TalentMarks.latchDecision(2, 18, false))
	end)

	it("latches only a partial miss that persists on the deferred pass", function()
		assert.are.equal("latch", TalentMarks.latchDecision(2, 18, true))
	end)
end)

describe("TalentMarks.nextDelay", function()
	it("always takes one more pass on the next frame after the show", function()
		assert.are.equal(0, TalentMarks.nextDelay(0, "ok", true))
		assert.are.equal(0, TalentMarks.nextDelay(0, "wait", false))
	end)

	it("takes a second deferred pass at half a second when the first still missed", function()
		assert.are.equal(0.5, TalentMarks.nextDelay(1, "wait", true))
		assert.are.equal(0.5, TalentMarks.nextDelay(1, "retry", true))
		assert.are.equal(0.5, TalentMarks.nextDelay(1, "ok", false))
	end)

	it("stops once the first deferred pass found everything, and after the second", function()
		assert.is_nil(TalentMarks.nextDelay(1, "ok", true))
		assert.is_nil(TalentMarks.nextDelay(2, "wait", false))
		assert.is_nil(TalentMarks.nextDelay(nil, "retry", false))
	end)

	it("counts the second deferred pass as the one a partial miss latches on", function()
		assert.is_false(TalentMarks.isFinalPass(0, false))
		assert.is_false(TalentMarks.isFinalPass(1, false))
		assert.is_true(TalentMarks.isFinalPass(2, false))
		-- An event's refresh (no pass) is final once the show's passes are done.
		assert.is_true(TalentMarks.isFinalPass(nil, true))
		assert.is_false(TalentMarks.isFinalPass(nil, false))
	end)
end)

describe("TalentMarks.badgeWidth", function()
	it("is the square plate for one digit, and digits plus 6 when wider", function()
		assert.are.equal(16, TalentMarks.badgeWidth(6))
		assert.are.equal(16, TalentMarks.badgeWidth(10))
		assert.are.equal(20, TalentMarks.badgeWidth(13.4))
		assert.are.equal(16, TalentMarks.badgeWidth(nil))
	end)
end)

-- Fix round 2: badges on a plate; the next one inverted, the unreachable dimmed.
describe("TalentMarks.badgeLook", function()
	it("inverts the next cell's plate: solid gold, dark digits, never dimmed", function()
		assert.are.same({ plate = "gold", plateAlpha = 1, digits = "onGold", alpha = 1 },
			TalentMarks.badgeLook({ next = true, have = 0, want = 5, lastIndex = 30 }, 12))
	end)

	it("mutes a taken cell's digits on the dark plate", function()
		assert.are.same({ plate = "background", plateAlpha = 0.7, digits = "muted", alpha = 1 },
			TalentMarks.badgeLook({ next = false, have = 3, want = 3, lastIndex = 3 }, 12))
	end)

	it("draws a planned cell in gold, at half alpha beyond the points spendable now", function()
		assert.are.same({ plate = "background", plateAlpha = 0.7, digits = "gold", alpha = 1 },
			TalentMarks.badgeLook({ next = false, have = 0, want = 3, lastIndex = 12 }, 12))
		assert.are.same({ plate = "background", plateAlpha = 0.7, digits = "gold", alpha = 0.5 },
			TalentMarks.badgeLook({ next = false, have = 0, want = 3, lastIndex = 13 }, 12))
	end)
end)

-- Fix round 2: one line on the client's own node tooltip.
describe("TalentMarks.nodeLine", function()
	it("names the plan's points in the cell", function()
		assert.are.equal("Forever Sixty: points 7 to 11 of your build",
			TalentMarks.nodeLine({ firstIndex = 7, lastIndex = 11, want = 5 }))
	end)

	it("names a single point once", function()
		assert.are.equal("Forever Sixty: point 3 of your build",
			TalentMarks.nodeLine({ firstIndex = 3, lastIndex = 3, want = 1 }))
	end)
end)

-- Fix round 2: the badges line /fs diag prints beside the chip's.
describe("TalentMarks.badgesLine", function()
	it("counts what was drawn and what missed", function()
		assert.are.equal("Talent badges: 18 drawn, 0 missing", TalentMarks.badgesLine(18, 0))
	end)
end)

-- Fix round 1 (seen in game): the strip must leave the client's Apply Changes button clear.
describe("TalentMarks.stripWidth", function()
	it("fits the text plus 12 px each side", function()
		assert.are.equal(324, TalentMarks.stripWidth(300, 108, 700, 1200))
	end)

	it("stops 8 px short of the Apply Changes button", function()
		-- Strip from x 108, button from x 400: at most 400 - 8 - 108 = 284.
		assert.are.equal(284, TalentMarks.stripWidth(300, 108, 400, 1200))
	end)

	it("stops at 45 % of the window's width when the button cannot be reached", function()
		-- 0.45 * 600 = 270, less the strip's own 16 px from the window's edge.
		assert.are.equal(254, TalentMarks.stripWidth(300, 108, nil, 600))
		assert.are.equal(324, TalentMarks.stripWidth(300, nil, nil, nil))
	end)
end)

-- Fix round 2: the pill sits 8 px inside the window's border, centred on Apply Changes.
describe("TalentMarks.stripOffsetY", function()
	it("centres the pill on the Apply Changes button", function()
		-- Button 22 high from 108 over a window bottom at 100: 8 + (22 - 24) / 2 = 7.
		assert.are.equal(7, TalentMarks.stripOffsetY(100, 108, 22))
	end)

	it("sits 8 px inside the window's border without the button", function()
		assert.are.equal(16, TalentMarks.stripOffsetY(100, nil, nil))
		assert.are.equal(16, TalentMarks.stripOffsetY(nil, 108, 22))
	end)
end)

-- Fix round 2: staged (not yet applied) ranks.
describe("Talents.nodeRanks", function()
	local Talents = require("Talents")

	it("reads the committed rank and the rank with staged clicks", function()
		assert.are.same({ 1, 2 }, { Talents.nodeRanks({ activeRank = 1, currentRank = 2, ranksPurchased = 2 }) })
		assert.are.same({ 2, 2 }, { Talents.nodeRanks({ activeRank = 2, ranksPurchased = 2 }) })
		assert.are.same({ 0, 1 }, { Talents.nodeRanks({ activeRank = 0, ranksPurchased = 1 }) })
		assert.are.same({ 3, 3 }, { Talents.nodeRanks({ ranksPurchased = 3 }) })
		assert.are.same({ 0, 0 }, { Talents.nodeRanks(nil) })
	end)
end)

-- Minor 6: one lookup per node per refresh, shared by the ring and the badges.
describe("TalentGlow.traitButton's per-refresh cache", function()
	local TalentGlow = require("TalentGlow")

	it("asks the client once per node, and remembers a miss too", function()
		local asked = 0
		local button = {}
		local frame = { GetTalentButtonByNodeID = function(_, node)
			asked = asked + 1
			return node == 1 and button or nil
		end }
		local cache = {}
		assert.are.equal(button, TalentGlow.traitButton(frame, 1, cache))
		assert.are.equal(button, TalentGlow.traitButton(frame, 1, cache))
		assert.is_nil(TalentGlow.traitButton(frame, 2, cache))
		assert.is_nil(TalentGlow.traitButton(frame, 2, cache))
		assert.are.equal(2, asked)
		-- Without a cache every call asks again.
		TalentGlow.traitButton(frame, 1)
		assert.are.equal(3, asked)
	end)
end)
