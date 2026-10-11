-- addon/tests/talent_marks_spec.lua
-- The talent marks (spec section 4.2): the strip's line and each badge's
-- state, the two pure rules the frames read. Pure: hand-built tables only,
-- no client mock and no frame.
local TalentMarks = require("TalentMarks")
local L = require("Locale")
local Theme = require("Theme")

local ARMS = { name = "Arms leveling", classSlug = "warrior", order = {} }

local function progress(matched, total, nextLevel, offPlan, done)
	return { matched = matched, total = total, nextLevel = nextLevel, offPlan = offPlan,
		point = (not done) and { index = matched + 1 } or nil, done = done == true }
end

describe("TalentMarks.stripText", function()
	it("reads the plan's count, the next level and nothing off plan", function()
		local text, warned = TalentMarks.stripText(ARMS, progress(11, 51, 21, 0))
		assert.are.equal("Arms leveling · 11 of 51 · next at level 21 · 0 off plan", text)
		assert.is_false(warned)
	end)

	it("colours the off-plan count in the warning colour, and only that segment", function()
		local text, warned = TalentMarks.stripText(ARMS, progress(11, 51, 22, 1))
		assert.are.equal("Arms leveling · 11 of 51 · next at level 22 · |cffff6b5c1 off plan|r", text)
		assert.is_true(warned)
	end)

	it("says complete once the build has no point left", function()
		local text = TalentMarks.stripText(ARMS, progress(51, 51, nil, 0, true))
		assert.are.equal("Arms leveling · 51 of 51 · complete · 0 off plan", text)
	end)

	it("leaves the level out where no level follows (the level cap)", function()
		local text, warned = TalentMarks.stripText(ARMS, progress(49, 51, nil, 2))
		assert.are.equal("Arms leveling · 49 of 51 · |cffff6b5c2 off plan|r", text)
		assert.is_true(warned)
	end)

	it("names a build with no name after its class", function()
		local text = TalentMarks.stripText({ classSlug = "warrior", order = {} }, progress(0, 51, 10, 0))
		assert.are.equal("Warrior build · 0 of 51 · next at level 10 · 0 off plan", text)
	end)

	it("has nothing to say without a build", function()
		assert.is_nil(TalentMarks.stripText(nil, progress(0, 0, nil, 0)))
	end)

	it("keeps the plan's own wording in Locale", function()
		assert.are.equal("%s · %d of %d · next at level %d · %d off plan", L.talentStrip)
		assert.are.equal("%s · %d of %d · complete · %d off plan", L.talentStripDone)
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
	it("has the 14 px badge with 10 px digits and the 24 px strip", function()
		assert.are.equal(14, Theme.SIZES.orderBadge)
		assert.are.equal(10, Theme.SIZES.orderBadgeText)
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
	it("is the 14 px square for one digit, and digits plus 4 when wider", function()
		assert.are.equal(14, TalentMarks.badgeWidth(6))
		assert.are.equal(14, TalentMarks.badgeWidth(10))
		assert.are.equal(16, TalentMarks.badgeWidth(11.3))
		assert.are.equal(14, TalentMarks.badgeWidth(nil))
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
		-- 0.45 * 600 = 270, less the strip's own 8 px inset.
		assert.are.equal(262, TalentMarks.stripWidth(300, 108, nil, 600))
		assert.are.equal(324, TalentMarks.stripWidth(300, nil, nil, nil))
	end)
end)

describe("TalentMarks.strataAbove", function()
	it("is one strata step above, and the top stays the top", function()
		assert.are.equal("HIGH", TalentMarks.strataAbove("MEDIUM"))
		assert.are.equal("DIALOG", TalentMarks.strataAbove("HIGH"))
		assert.are.equal("TOOLTIP", TalentMarks.strataAbove("TOOLTIP"))
		assert.is_nil(TalentMarks.strataAbove(nil))
		assert.is_nil(TalentMarks.strataAbove("PARENT"))
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
