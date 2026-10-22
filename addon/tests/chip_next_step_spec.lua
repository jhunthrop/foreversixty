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
local function boot(chipPrefs)
	mock.install({
		class = { name = "Hunter", token = "HUNTER" },
		race = { name = "High Order Skyborne", token = "HighOrderSkyborne" },
		realm = "Classic Beta PvP",
		region = 1,
		playerName = "Bow Jackzon",
		level = 24,
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
	end

	it("sits just under the buff frame where the buffs are in the way, and beside the minimap otherwise", function()
		boot()
		geometry()
		local Chip, Theme = require("Chip"), require("Theme")
		assert.are.equal("minimap", Chip.place())
		-- The buffs end 10 px left of the minimap, rows from the top down: the line's rectangle meets them.
		local frame = buffs(1300, 1690, 1060, 1000)
		assert.are.equal(frame, Chip.buffsInTheWay(_G.Minimap, Chip.frame))
		assert.are.equal("buffs", Chip.place())
		local point = lastPoint(Chip.frame)
		assert.are.same({ "TOPRIGHT", frame, "BOTTOMRIGHT" }, { point[1], point[2], point[3] })
		assert.are.equal(-Theme.SIZES.chipMinimapGap / 2, point[5])
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
