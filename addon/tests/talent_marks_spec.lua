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
