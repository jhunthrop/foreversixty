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

--- The strip's whole line, as the fitter gives it with no measure (fix
--- round 5: one source, no separate stripText), and whether it warns.
local function stripLine(name, prog, state)
	local segments, warned = TalentMarks.stripSegments(name, prog, state)
	if segments == nil then
		return nil, false
	end
	return (TalentMarks.fitStrip(segments, nil, nil)), warned
end

-- Fix round 2 (personas on marks2-*): "next at level 21" at level 21 with
-- twelve points to spend read as "later"; the strip says "spend 12 now"
-- while points wait. Fix round 3: "off plan" (spec section 4.2), muted at 0
-- and a warning from 1; no "click Apply Changes" (the client's lit button says it).
local MUTED, WARN, GOLD, BODY = "|cff9a9484", "|cffff6b5c", "|cffe5b955", "|cffe9e4d8"

describe("the strip's line (TalentMarks.stripSegments)", function()
	it("says spend N now while points wait, with no level", function()
		local text, warned = stripLine("Arms / Fury build", progress(0, 51, 21, 0), { unspent = 12 })
		assert.are.equal("Arms / Fury build · 0 of 51 points · spend 12 now · " .. MUTED .. "0 off plan|r", text)
		assert.is_false(warned)
	end)

	it("names the level the next point arrives at when none waits", function()
		local text = stripLine("Arms leveling", progress(11, 51, 22, 0), { unspent = 0 })
		assert.are.equal("Arms leveling · 11 of 51 points · next point at level 22 · " .. MUTED .. "0 off plan|r",
			text)
	end)

	it("colours the off-plan count as a warning from one point", function()
		local text, warned = stripLine("Arms leveling", progress(11, 51, 22, 1), { unspent = 0 })
		assert.are.equal("Arms leveling · 11 of 51 points · next point at level 22 · " .. WARN .. "1 off plan|r",
			text)
		assert.is_true(warned)
	end)

	it("leaves the level out where no level follows (the level cap)", function()
		local text = stripLine("Arms leveling", progress(49, 51, nil, 2), { unspent = 0 })
		assert.are.equal("Arms leveling · 49 of 51 points · " .. WARN .. "2 off plan|r", text)
	end)

	it("says complete once the build has no point left, spendable points or not", function()
		local text = stripLine("Arms leveling", progress(51, 51, nil, 0, true), { unspent = 3 })
		assert.are.equal("Arms leveling · 51 of 51 points · complete · " .. MUTED .. "0 off plan|r", text)
	end)

	it("says nothing about Apply Changes while a click is staged", function()
		local text = stripLine("Arms leveling", progress(12, 51, 21, 0), { unspent = 0, staged = true })
		assert.are.equal("Arms leveling · 12 of 51 points · next point at level 21 · " .. MUTED .. "0 off plan|r", text)
	end)

	it("has nothing to say without a build", function()
		assert.is_nil(stripLine(nil, progress(0, 0, nil, 0), {}))
	end)

	-- Fix round 4: while the open window hides the chip, the strip leads with its line.
	it("leads with the chip's next talent while the window hides the chip", function()
		local text = stripLine("Arms / Fury build", progress(0, 51, 21, 0),
			{ unspent = 12, next = "Improved Heroic Strike (1 of 3)" })
		assert.are.equal("Next: Improved Heroic Strike (1 of 3) · Arms / Fury build · 0 of 51 points · spend 12 now · "
			.. MUTED .. "0 off plan|r", text)
		assert.are.equal("Next: %s", L.talentStripNextTalent)
	end)

	it("keeps its wording in Locale", function()
		assert.are.equal("spend %d now", L.talentStripSpend)
		assert.are.equal("next point at level %d", L.talentStripNext)
		assert.are.equal("%d of %d points", L.talentStripCount)
		assert.are.equal("%d off plan", L.talentStripOffPlan)
		assert.is_nil(L.talentStripApply)
	end)
end)

-- Fix round 3 (seen in game: "click Apply Chan…" cut by the pill's cap): a
-- segment fits whole or is dropped, the action first, then the build name;
-- fix round 4: then the leading "Next: …".
describe("TalentMarks.fitStrip", function()
	local function measure(text)
		return #text
	end
	local segments = TalentMarks.stripSegments("Arms / Fury build", progress(1, 51, 21, 0), { unspent = 11 })
	local full = "Arms / Fury build · 1 of 51 points · spend 11 now · " .. MUTED .. "0 off plan|r"

	it("keeps every segment when the line fits", function()
		assert.are.equal(full, TalentMarks.fitStrip(segments, measure, #full))
		assert.are.equal(full, TalentMarks.fitStrip(segments, measure, nil))
	end)

	it("drops spend N now first, whole", function()
		local short = "Arms / Fury build · 1 of 51 points · " .. MUTED .. "0 off plan|r"
		assert.are.equal(short, TalentMarks.fitStrip(segments, measure, #full - 1))
		assert.are.equal(short, TalentMarks.fitStrip(segments, measure, #short))
	end)

	it("then the build name, never the count or the off-plan count", function()
		local shortest = "1 of 51 points · " .. MUTED .. "0 off plan|r"
		assert.are.equal(shortest, TalentMarks.fitStrip(segments, measure, #shortest))
		assert.are.equal(shortest, TalentMarks.fitStrip(segments, measure, 3))
	end)

	-- Fix round 5: a nil measure is the whole line, whatever the room.
	it("keeps the whole line when it cannot measure", function()
		assert.are.same({ full, true }, { TalentMarks.fitStrip(segments, nil, 3) })
	end)

	-- Fix round 5: the count and off-plan count never give way; when even they
	-- overflow the room, the fitter says so and the pill grows to fit them.
	it("says when even the count and off-plan count do not fit", function()
		local shortest = "1 of 51 points · " .. MUTED .. "0 off plan|r"
		assert.are.same({ shortest, true }, { TalentMarks.fitStrip(segments, measure, #shortest) })
		assert.are.same({ shortest, false }, { TalentMarks.fitStrip(segments, measure, 3) })
	end)

	it("gives up the leading Next segment last", function()
		local led = TalentMarks.stripSegments("Arms / Fury build", progress(1, 51, 21, 0),
			{ unspent = 11, next = "Improved Heroic Strike (2 of 3)" })
		local noName = "Next: Improved Heroic Strike (2 of 3) · 1 of 51 points · " .. MUTED .. "0 off plan|r"
		assert.are.equal(noName, TalentMarks.fitStrip(led, measure, #noName))
		local bare = "1 of 51 points · " .. MUTED .. "0 off plan|r"
		assert.are.equal(bare, TalentMarks.fitStrip(led, measure, #noName - 1))
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
-- Fix round 3: dimming by the cell's first point; digits never under 0.6.
-- Fix round 4 (seen in game: a bright icon washed the digit out through a
-- 70 % plate): every plate is near-opaque; dimming darkens its colour.
describe("TalentMarks.badgeLook", function()
	it("inverts the next cell's plate: solid gold, dark digits, never dimmed", function()
		assert.are.same({ plate = "gold", plateAlpha = 1, edgeAlpha = 1, digits = "onGold", digitsAlpha = 1 },
			TalentMarks.badgeLook({ next = true, have = 0, want = 5, firstIndex = 26, lastIndex = 30 }, 12))
	end)

	it("mutes a taken cell's digits on the near-opaque dark plate", function()
		assert.are.same({ plate = "background", plateAlpha = 0.9, edgeAlpha = 1, digits = "muted", digitsAlpha = 1 },
			TalentMarks.badgeLook({ next = false, have = 3, want = 3, firstIndex = 1, lastIndex = 3 }, 12))
	end)

	it("keeps a cell bright while its first point is reachable", function()
		assert.are.same({ plate = "background", plateAlpha = 0.9, edgeAlpha = 1, digits = "gold", digitsAlpha = 1 },
			TalentMarks.badgeLook({ next = false, have = 0, want = 2, firstIndex = 12, lastIndex = 13 }, 12))
	end)

	it("dims beyond the reachable points with a darker plate, never a thinner one", function()
		local look = TalentMarks.badgeLook({ next = false, have = 0, want = 3, firstIndex = 13, lastIndex = 15 }, 12)
		assert.are.same({ plate = "inset", plateAlpha = 0.9, edgeAlpha = 0.5, digits = "gold", digitsAlpha = 0.6 }, look)
	end)

	it("never draws a plate under 0.9 alpha, nor digits under 0.6", function()
		for _, cell in ipairs({
			{ next = true, have = 0, want = 1, firstIndex = 1, lastIndex = 1 },
			{ next = false, have = 1, want = 1, firstIndex = 1, lastIndex = 1 },
			{ next = false, have = 0, want = 1, firstIndex = 5, lastIndex = 5 },
			{ next = false, have = 0, want = 1, firstIndex = 50, lastIndex = 50 },
		}) do
			local look = TalentMarks.badgeLook(cell, 12)
			assert.is_true(look.plateAlpha >= 0.9)
			assert.is_true(look.digitsAlpha >= 0.6)
		end
	end)
end)

-- Fix round 4 (spec section 4.2 amended): a multi-rank cell's badge shows
-- the plan's point range, a one-point cell's its single number.
describe("TalentMarks.badgeText", function()
	it("shows the range with an en dash for a multi-rank cell", function()
		assert.are.equal("7\226\128\14711", TalentMarks.badgeText({ want = 5, firstIndex = 7, lastIndex = 11 }))
		assert.are.equal("7–11", TalentMarks.badgeText({ want = 5, firstIndex = 7, lastIndex = 11 }))
	end)

	it("shows the single number for a one-point cell", function()
		assert.are.equal("12", TalentMarks.badgeText({ want = 1, firstIndex = 12, lastIndex = 12 }))
	end)
end)

-- Fix round 3: the node tooltip gains a blank line, then "Forever Sixty" in
-- gold and the plan's ranks and points in body colour; an off-plan node
-- holding a point gets a warning line.
describe("TalentMarks.nodeLine", function()
	it("names the plan's ranks and points in the cell", function()
		assert.are.equal(GOLD .. "Forever Sixty:|r " .. BODY .. "Your build: 5 ranks, points 7 to 11|r",
			TalentMarks.nodeLine({ firstIndex = 7, lastIndex = 11, want = 5 }))
		assert.are.equal(GOLD .. "Forever Sixty:|r " .. BODY .. "Your build: 2 ranks, points 12 to 13|r",
			TalentMarks.nodeLine({ firstIndex = 12, lastIndex = 13, want = 2 }))
	end)

	it("names a single point once", function()
		assert.are.equal(GOLD .. "Forever Sixty:|r " .. BODY .. "Your build: point 12|r",
			TalentMarks.nodeLine({ firstIndex = 12, lastIndex = 12, want = 1 }))
	end)

	it("warns on a node the plan does not want", function()
		assert.are.equal(WARN .. "Forever Sixty: not in your build|r", TalentMarks.nodeLine({ offPlan = true }))
	end)
end)

-- Fix round 2: the badges line /fs diag prints beside the chip's.
describe("TalentMarks.badgesLine", function()
	it("counts what was drawn and what missed", function()
		assert.are.equal("Talent badges: 18 drawn, 0 missing", TalentMarks.badgesLine(18, 0))
	end)

	-- Fix round 3: after the latch the badges are off, not "drawn".
	it("says the badges are off after a latch, and how many found no button", function()
		assert.are.equal("Talent badges: off (2 found no button)", TalentMarks.badgesLine(16, 2, true))
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

-- Fix round 5: a line that fits takes the capped width; one that does not
-- (only the count and off-plan count left) takes its own width, past the cap.
describe("TalentMarks.pillWidth", function()
	it("caps a line that fits at the Apply Changes button", function()
		assert.are.equal(284, TalentMarks.pillWidth(300, true, 108, 400, 1200))
		assert.are.equal(124, TalentMarks.pillWidth(100, true, 108, 400, 1200))
	end)

	it("grows past the cap for a line that does not fit", function()
		assert.are.equal(324, TalentMarks.pillWidth(300, false, 108, 400, 1200))
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

-- Fix round 5: the spellbook tab of the same window is not the talent
-- window; the chip stays and the marks stay off there.
describe("TalentGlow.talentsShowing", function()
	local TalentGlow = require("TalentGlow")
	local function shown(value)
		return { IsShown = function()
			return value
		end }
	end

	it("is the talents child's own answer when the window has one", function()
		local frame = shown(true)
		frame.TalentsFrame = shown(false)
		assert.is_false(TalentGlow.talentsShowing(frame))
		frame.TalentsFrame = shown(true)
		assert.is_true(TalentGlow.talentsShowing(frame))
	end)

	it("is the window itself when it has no talents child", function()
		assert.is_true(TalentGlow.talentsShowing(shown(true)))
	end)

	it("makes openFrame answer nothing while the spellbook tab shows", function()
		local saved = _G.PlayerSpellsFrame
		local frame = shown(true)
		frame.TalentsFrame = shown(false)
		_G.PlayerSpellsFrame = frame
		local closed = TalentGlow.openFrame()
		frame.TalentsFrame = shown(true)
		local open = TalentGlow.openFrame()
		_G.PlayerSpellsFrame = saved
		assert.is_nil(closed)
		assert.are.equal(frame, open)
	end)
end)

-- Fix round 5: the ring's dark edge keeps its alpha when the pulse stops
-- (fake regions, no client mock): only the gold pulses and is reset to 1.
describe("Theme.showGlow's dark edge", function()
	local saved = {}

	local function fakeRegion()
		local region = { alphaCalls = {} }
		function region.SetPoint() end
		function region.SetHeight() end
		function region.SetWidth() end
		function region.SetFrameLevel() end
		function region.Show(self)
			self.shown = true
		end
		function region.Hide(self)
			self.shown = false
		end
		function region.SetColorTexture(self, _, _, _, a)
			self.colorAlpha = a
		end
		function region.SetAlpha(self, a)
			self.alphaCalls[#self.alphaCalls + 1] = a
		end
		function region.CreateAnimationGroup(self)
			self.animated = true
			local group = {}
			function group.CreateAnimation()
				return { SetFromAlpha = function() end, SetToAlpha = function() end, SetDuration = function() end }
			end
			function group.SetLooping() end
			function group.Play() end
			function group.Stop() end
			return group
		end
		function region.CreateTexture()
			return fakeRegion()
		end
		return region
	end

	before_each(function()
		saved.CreateFrame, saved.C_Timer = _G.CreateFrame, _G.C_Timer
		_G.CreateFrame = function()
			return fakeRegion()
		end
		-- The ten seconds pass at once: the pulse stops inside showGlow.
		_G.C_Timer = { After = function(_, action)
			action()
		end }
	end)

	after_each(function()
		_G.CreateFrame, _G.C_Timer = saved.CreateFrame, saved.C_Timer
	end)

	it("never pulses the edge nor resets its alpha, and shows and hides it with the ring", function()
		local button = fakeRegion()
		button.GetFrameLevel = function()
			return 5
		end
		Theme.showGlow(button, true, 10, { inset = 6, thickness = 3, edge = 1 })
		assert.are.equal(4, #button.foreverSixtyGlowEdge)
		for _, edge in ipairs(button.foreverSixtyGlowEdge) do
			assert.are.equal(0.6, edge.colorAlpha)
			assert.is_nil(edge.animated)
			assert.are.same({}, edge.alphaCalls)
			assert.is_true(edge.shown)
		end
		for _, gold in ipairs(button.foreverSixtyGlow) do
			assert.is_true(gold.animated)
			assert.are.equal(1, gold.alphaCalls[#gold.alphaCalls])
		end
		Theme.hideGlow(button)
		for _, edge in ipairs(button.foreverSixtyGlowEdge) do
			assert.is_false(edge.shown)
			assert.are.same({}, edge.alphaCalls)
		end
	end)
end)
