-- addon/tests/chip_next_step_spec.lua
-- A.2: the line names the next point when every point the character has is spent ("Next at 25:
-- Focused Fire (1 of 3)", "Build complete · next point at level 25"), keeps one steady class icon in
-- every state, wraps a ready check's words whole instead of cutting them, and sits clear of the
-- client's buff row at its default spot. The strip's words are in talent_marks_spec.lua, the chat
-- line after Use in build_dialog_spec.lua.
local helper = require("spec_helper")
local mock = require("wow_mock")

local Data = require("Data")

--- Trait ranks ({ [node] = rank }) for `points` points down the hunter's tab `tab`, in the data's order.
local function traitRanks(tab, points)
	local ranks = {}
	for _, talent in ipairs(Data.classes.hunter.tabs[tab].talents) do
		if points <= 0 then
			break
		end
		local rank = math.min(points, talent.maxRank)
		ranks[talent.node] = rank
		points = points - rank
	end
	return ranks
end

--- Boot the addon as a level 24 hunter with fifteen points spent, the minimap 140 wide, then
--- PLAYER_LOGIN and the first PLAYER_ENTERING_WORLD.
local function boot(chipPrefs, level)
	mock.install({
		class = { name = "Hunter", token = "HUNTER" },
		race = { name = "High Order Skyborne", token = "HighOrderSkyborne" },
		realm = "Classic Beta PvP",
		region = 1,
		playerName = "Bow Jackzon",
		level = level or 24,
		traits = { configID = 7, ranks = traitRanks(2, 15) },
	})
	_G.UIParent.width, _G.UIParent.height = 1920, 1080
	_G.Minimap.width, _G.Minimap.height = 140, 140
	chipPrefs = chipPrefs or { shown = true, hiddenUntilLevel = 0, hiddenFor = "", pathFor = "", pathChosen = "",
		combatHide = true, byCharacter = {} }
	_G.ForeverSixtyDB = { welcomedAt = 1793000000, ui = { chip = chipPrefs } }
	require("Theme").reset()
	for _, name in ipairs({ "Prefs", "Widgets", "Export", "Talents", "Gear", "ItemTooltip", "Follow", "TalentGlow",
		"Chip", "TalentMarks", "Minimap", "Comms", "BuildDialog", "Settings", "SiteStatus" }) do
		helper.load(name)
	end
	local Options = helper.load("Options")
	Options.register()
	Options.onEvent(Options.frame, "PLAYER_LOGIN")
	Options.onEvent(Options.frame, "PLAYER_ENTERING_WORLD", true, false)
	return Options
end

local function lastPoint(frame)
	return frame.points[#frame.points]
end

describe("the chip's words once every point is spent (A.2)", function()
	local Chip, L, Theme

	setup(function()
		Chip, L, Theme = require("Chip"), require("Locale"), require("Theme")
	end)

	after_each(function()
		mock.uninstall()
	end)

	local function inputs(overrides)
		local base = {
			level = 24, classIcon = "crest-hunter", now = 100, clock = 0, shown = true, combatHide = true,
			unspent = 0, nextLevel = 25,
			build = { name = "Beast Mastery build", order = { {}, {}, {} } },
			progress = { matched = 1, total = 3, nextLevel = 25, point = { tab = 1, tier = 1, column = 2, index = 2 } },
			pointName = "Focused Fire", pointRank = { have = 0, want = 3 },
		}
		for key, value in pairs(overrides or {}) do
			-- false takes the field away (a nil cannot be set in a table constructor).
			base[key] = value or nil
		end
		return base
	end

	it("leads the idle line with the level the next point arrives at when none is to spend", function()
		local chosen = Chip.model(inputs()).chosen
		assert.are.equal("idle", chosen.kind)
		assert.are.equal("Next at 25: Focused Fire (1 of 3)", chosen.text)
		assert.are.equal(string.format(L.chipNextAt, 25, "Focused Fire (1 of 3)"), chosen.text)
		-- At the cap no level follows: the old form.
		assert.are.equal("Next: Focused Fire (1 of 3)", Chip.model(inputs({ nextLevel = false })).chosen.text)
	end)

	it("says Build complete with the next point's level, and Build complete alone when none follows", function()
		local done = inputs({ progress = { matched = 3, total = 3 }, pointName = false, completedAt = 98 })
		local chosen = Chip.model(done).chosen
		assert.are.equal("complete", chosen.kind)
		assert.are.equal("Build complete \194\183 next point at level 25", chosen.text)
		assert.are.equal(string.format(L.chipCompleteNext, 25), chosen.text)
		-- The short form the line falls back to when it does not fit.
		assert.are.equal("Build complete", chosen.short)
		-- The plan finished and every point of the level taken: "Build complete" alone.
		assert.are.equal("Build complete", Chip.model(inputs({ progress = { matched = 3, total = 3 },
			pointName = false, completedAt = 98, nextLevel = false })).chosen.text)
		-- A point still to spend is not news of a later one.
		assert.are.equal("Build complete", Chip.model(inputs({ progress = { matched = 3, total = 3 },
			pointName = false, completedAt = 98, unspent = 1 })).chosen.text)
		assert.are.equal("Build complete", L.chipComplete)
	end)

	it("hands the live line the level the next point arrives at", function()
		boot()
		Chip = require("Chip")
		local live = Chip.inputs(require("Data"))
		-- Level 24 with fifteen spent: the next point is level 25's.
		assert.are.equal(25, live.nextLevel)
		assert.are.equal(0, live.unspent)
	end)

	it("keeps the one class icon on the line in every state", function()
		boot()
		Chip = require("Chip")
		local classIcon = Theme.classIcon("HUNTER")
		local function iconNow()
			return mock.lastCall(Chip.icon, "SetTexture")[1]
		end
		assert.are.equal(classIcon, iconNow())
		-- A ready check's tick and cross are the words' business, never the icon's.
		Chip.setRaidCheck({}, "Ready check (only you see this): nothing missing", false, true)
		assert.are.equal("raidcheck", Chip.current.chosen.kind)
		assert.are.equal(Theme.MEDIA.gearTick, Chip.current.chosen.icon)
		assert.are.equal(classIcon, iconNow())
		Chip.setRaidCheck({ "flask" }, "Ready check (only you see this): missing flask", false, false)
		assert.are.equal(classIcon, iconNow())
		Chip.setRaidCheck(nil)
		assert.are.equal(classIcon, iconNow())
	end)

	it("wraps a ready check's words inside the widest line, whole, never cut", function()
		boot()
		Chip = require("Chip")
		local text = "Ready check (only you see this): nothing missing at your level \194\183 "
			.. "no flask, food or buff is expected below 35"
		local room = Theme.SIZES.chipWide - Theme.SIZES.chipTextInset
		-- Six px a letter: far wider than the 280 px line, so it wraps at the room it is given.
		Chip.text.GetStringWidth = function(region)
			local width = utf8.len(region:GetText() or "") * 6
			return region.wrapped and math.min(width, room) or width
		end
		Chip.text.GetStringHeight = function(region)
			return math.ceil(utf8.len(region:GetText() or "") * 6 / room) * 14
		end
		Chip.text.SetWordWrap = function(region, on)
			region.wrapped = on
		end
		Chip.setRaidCheck({}, text, false, true)
		assert.are.equal(text, Chip.text:GetText())
		assert.is_true(Chip.text.wrapped)
		assert.are.equal(Theme.SIZES.chipWide, Chip.frame:GetWidth())
		assert.are.equal(room, mock.lastCall(Chip.text, "SetWidth")[1])
		assert.is_true(Chip.frame:GetHeight() > Theme.SIZES.chip)
		assert.is_nil(Chip.text:GetText():find(L.ellipsis, 1, true))
		-- A line that fits stays on one line at the ordinary height.
		Chip.setRaidCheck({}, "Ready check: nothing missing", false, true)
		assert.is_false(Chip.text.wrapped)
		assert.are.equal(Theme.SIZES.chip, Chip.frame:GetHeight())
	end)
end)

describe("the line's words at a login, for the states that had none (A.2)", function()
	after_each(function()
		mock.uninstall()
	end)

	--- A 15-point plan down the hunter's tab 2 in the data's own order: exactly the points traitRanks spends.
	local function fifteenPointBuild()
		local order, left = {}, 15
		for _, talent in ipairs(Data.classes.hunter.tabs[2].talents) do
			for _ = 1, math.min(left, talent.maxRank) do
				order[#order + 1] = { tab = 2, tier = talent.tier, column = talent.column }
			end
			left = left - math.min(left, talent.maxRank)
			if left <= 0 then
				break
			end
		end
		local Codec = require("Codec")
		return Codec.loadBuild(Codec.encodeFSB1({ dataBuild = Data.build, classSlug = "hunter", order = order,
			gear = {} }), Data)
	end

	local function words()
		local Chip = require("Chip")
		assert.is_true(Chip.text:IsShown())
		local text = Chip.text:GetText()
		assert.is_truthy(text)
		assert.is_true(#text > 0)
		return text
	end

	it("says Build complete with the next point's level for a 15-point plan with 15 points spent at level 24",
		function()
			boot()
			local Chip, Follow = require("Chip"), require("Follow")
			-- The follow state arrives after the first draw, as at a login; the line is redrawn with it.
			Follow.build = fifteenPointBuild()
			Chip.refresh()
			-- Completing it flashes the line first...
			assert.are.equal("complete", Chip.current.chosen.kind)
			assert.are.equal("Build complete \194\183 next point at level 25", words())
			-- ...and once the flash is over the idle line says the same, never an empty font string.
			Chip.completedAt = -1000
			Chip.refresh()
			assert.are.equal("idle", Chip.current.chosen.kind)
			assert.are.equal("Build complete \194\183 next point at level 25", words())
			assert.is_true(Chip.frame:GetWidth() >= require("Theme").SIZES.chip)
			assert.is_true(Chip.frame:IsShown())
			-- And the same through the model as the client feeds it.
			assert.are.equal(25, Chip.inputs(Data).nextLevel)
		end)

	it("keeps words for the first-login state of A.1: a point to spend with the default build, and Next at 25", function()
		boot(nil, 25)
		local Chip = require("Chip")
		-- Level 25 with 15 spent: one point to spend, the chip says which talent to take.
		assert.are.equal("point", Chip.current.chosen.kind)
		assert.is_truthy(words():find("^Take "))
		mock.uninstall()
		boot(nil, 24)
		assert.is_truthy(words():find("^Next at 25: ."))
	end)

	it("is empty only for a plan finished at the level cap", function()
		boot(nil, 60)
		local Chip = require("Chip")
		assert.is_nil(Chip.current.chosen.text:find("next point", 1, true))
	end)
end)

describe("the line's text always has a width, so the frame is wider than its icon", function()
	after_each(function()
		mock.uninstall()
	end)

	-- Six px a letter, a font string with no width (or a width of 0) drawing nothing: its string width is 0.
	local function measured(Chip)
		local room = require("Theme").SIZES.chipWide - require("Theme").SIZES.chipTextInset
		Chip.text.GetStringWidth = function(region)
			local width = utf8.len(region:GetText() or "") * 6
			local set = mock.lastCall(region, "SetWidth")
			if set ~= nil and set[1] == 0 then
				return 0
			end
			return region.wrapped and math.min(width, room) or width
		end
		Chip.text.GetStringHeight = function(region)
			return math.ceil(utf8.len(region:GetText() or "") * 6 / room) * 14
		end
		Chip.text.SetWordWrap = function(region, on)
			region.wrapped = on
		end
	end

	it("is wider than the 28 px icon for a short Next line", function()
		boot()
		local Chip, Theme = require("Chip"), require("Theme")
		measured(Chip)
		Chip.refresh()
		assert.is_truthy(Chip.text:GetText():find("^Next"))
		assert.is_true(Chip.text:IsShown())
		assert.is_true(Chip.frame:GetWidth() > Theme.SIZES.chip)
		for _, call in ipairs(Chip.text.calls) do
			assert.is_false(call.method == "SetWidth" and call[1] == 0)
		end
	end)

	it("is wider than the icon for the ready-check sentence, wrapped, and again for a short line after it", function()
		boot()
		local Chip, Theme = require("Chip"), require("Theme")
		measured(Chip)
		Chip.setRaidCheck({}, "Ready check (only you see this): nothing missing at your level \194\183 "
			.. "no flask, food or buff is expected below 35", false, true)
		assert.is_true(Chip.frame:GetWidth() > Theme.SIZES.chip)
		assert.is_true(Chip.text:GetStringWidth() > 0)
		Chip.setRaidCheck({}, "Ready check: nothing missing", false, true)
		assert.is_true(Chip.frame:GetWidth() > Theme.SIZES.chip)
		assert.is_true(Chip.text:GetStringWidth() > 0)
		for _, call in ipairs(Chip.text.calls) do
			assert.is_false(call.method == "SetWidth" and call[1] == 0)
		end
	end)
end)

describe("the line's default spot and the client's buff row (A.2)", function()
	after_each(function()
		mock.uninstall()
		_G.BuffFrame = nil
	end)

	local function buffs(left, right, top, bottom)
		local frame = _G.CreateFrame("Frame", "BuffFrame", _G.UIParent)
		frame.GetLeft = function() return left end
		frame.GetRight = function() return right end
		frame.GetTop = function() return top end
		frame.GetBottom = function() return bottom end
		_G.BuffFrame = frame
		return frame
	end

	local function geometry()
		_G.Minimap.GetLeft = function() return 1700 end
		_G.Minimap.GetTop = function() return 1060 end
		_G.UIParent.GetRight = function() return 1920 end
		_G.UIParent.GetTop = function() return 1080 end
	end

	local function offsets()
		local point = lastPoint(require("Chip").frame)
		return { point[1], point[2], point[3], point[4], point[5] }
	end

	it("sits just under the buffs, as plain offsets from the screen, where they are in the way", function()
		boot()
		geometry()
		local Chip, Theme = require("Chip"), require("Theme")
		assert.are.equal("minimap", Chip.place())
		-- The buffs end 10 px left of the minimap, rows from the top down: the line's rectangle meets them.
		buffs(1300, 1690, 1060, 1000)
		assert.are.equal("buffs", Chip.place())
		-- Right edge at the minimap's left less the gap, top a half gap under the buffs' bottom.
		local gap = Theme.SIZES.chipMinimapGap
		assert.are.same({ "TOPRIGHT", _G.UIParent, "TOPRIGHT", 1700 - gap - 1920, 1000 - 1080 - gap / 2 }, offsets())
		assert.is_true(Chip.onScreen(Chip.frame) ~= false)
	end)

	it("never anchors to the buff frame, so it cannot follow it off screen", function()
		boot()
		geometry()
		local frame = buffs(1300, 1690, 1060, 1000)
		local Chip = require("Chip")
		Chip.place()
		for _, point in ipairs(Chip.frame.points) do
			assert.are_not.equal(frame, point[2])
		end
	end)

	it("draws the line beside the minimap whatever the buff frame answers (nil, zero, off screen)", function()
		boot()
		geometry()
		local Chip = require("Chip")
		local function placedAt()
			assert.are.equal("minimap", Chip.place())
			assert.are.equal(_G.Minimap, lastPoint(Chip.frame)[2])
		end
		buffs(nil, nil, nil, nil)
		placedAt()
		buffs(0, 0, 0, 0)
		placedAt()
		-- Buffs hanging off the bottom of the screen: the spot under them is off screen.
		buffs(1300, 1690, 1060, -5)
		placedAt()
		buffs(1300, 1690, 1060, 5)
		placedAt()
		-- The screen's own size unknown: the plain default.
		_G.UIParent.GetRight = function() return nil end
		buffs(1300, 1690, 1060, 1000)
		placedAt()
	end)

	it("puts the line back beside the minimap, not off for the session, when the buff move lands it off screen",
		function()
			boot()
			geometry()
			local Chip = require("Chip")
			buffs(1300, 1690, 1060, 1000)
			assert.are.equal("buffs", Chip.place())
			-- The client lays that spot out left of the screen's edge.
			Chip.frame.GetLeft = function(own)
				return lastPoint(own)[2] == _G.Minimap and 1500 or -30
			end
			assert.is_true(Chip.selfCheck())
			assert.is_nil(Chip.disabled)
			assert.is_true(Chip.noBuffMove)
			assert.are.equal(_G.Minimap, lastPoint(Chip.frame)[2])
			assert.is_true(Chip.frame:IsShown())
		end)

	it("places the line again on /fs chip show, even one the render check switched off", function()
		local Options = boot()
		geometry()
		local Chip = require("Chip")
		buffs(1300, 1690, 1060, 1000)
		Chip.disabled = true
		Chip.frame:Hide()
		Options.handle("chip")
		Options.handle("chip")
		Chip = require("Chip")
		assert.is_nil(Chip.disabled)
		assert.are.equal(_G.Minimap, lastPoint(Chip.frame)[2])
		assert.is_true(Chip.frame:IsShown())
		assert.is_true(Chip.isDrawn())
	end)

	it("stays beside the minimap when the buffs are elsewhere, hidden or cannot be read", function()
		boot()
		geometry()
		local Chip = require("Chip")
		-- Far below the line, or far to its left: no overlap.
		buffs(1300, 1690, 900, 840)
		assert.is_nil(Chip.buffsInTheWay(_G.Minimap, Chip.frame))
		buffs(100, 400, 1060, 1000)
		assert.is_nil(Chip.buffsInTheWay(_G.Minimap, Chip.frame))
		local frame = buffs(1300, 1690, 1060, 1000)
		frame.IsShown = function() return false end
		assert.is_nil(Chip.buffsInTheWay(_G.Minimap, Chip.frame))
		frame.IsShown = function() return true end
		frame.GetLeft = function() return nil end
		assert.is_nil(Chip.buffsInTheWay(_G.Minimap, Chip.frame))
		assert.are.equal("minimap", Chip.place())
		_G.BuffFrame = nil
		assert.is_nil(Chip.buffsInTheWay(_G.Minimap, Chip.frame))
	end)

	it("never moves a spot the player dragged it to", function()
		boot({ shown = true, point = "TOPRIGHT", x = -300, y = -200 })
		geometry()
		buffs(1300, 1690, 1060, 1000)
		local Chip = require("Chip")
		assert.are.equal("TOPRIGHT", Chip.place())
	end)
end)
