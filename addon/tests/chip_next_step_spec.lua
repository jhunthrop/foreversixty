-- addon/tests/chip_next_step_spec.lua
-- The chip as the one launcher, live (design/specs/2026-11-13-addon-hub.md 6): its line is the hub's
-- most urgent thing, a click opens the tab that line names and again closes it, a right-click opens
-- Settings, in combat it says After combat and opens nothing, the minimap button and the compartment
-- entry do exactly what it does, tonight's raid leads once you answered, and it sits clear of the
-- client's buff row at its default spot.
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
	chipPrefs = chipPrefs or { shown = true, pathFor = "", pathChosen = "",
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

local EQUIP = "Equip Gloves of Holy Might +11.2"

--- The Gear tab replaced by one that has `items` waiting (Hub.register by id, as a tab lane does).
local function gearWaiting(items)
	local Hub, L = require("Hub"), require("Locale")
	Hub.register({ id = "gear", label = L.hubTabGear, build = function() end, show = function() end,
		waiting = function()
			return items
		end })
	return Hub
end

local function click(frame, button)
	frame:GetScript("OnClick")(frame, button or "LeftButton")
end

describe("the chip as the launcher", function()
	after_each(function()
		mock.uninstall()
	end)

	it("names the most urgent thing and opens the tab it names; again on that tab, it closes", function()
		boot()
		local Hub = gearWaiting({ { key = "bag", line = EQUIP } })
		local Chip = require("Chip")
		Chip.refresh()
		assert.are.equal(EQUIP, Chip.text:GetText())
		assert.are.equal("gear", Chip.current.opens)
		click(Chip.frame)
		assert.is_true(Hub.isShown())
		assert.are.equal("gear", Hub.active)
		click(Chip.frame)
		assert.is_false(Hub.isShown())
	end)

	it("opens the tab the line names even while the window shows another", function()
		boot()
		local Hub = gearWaiting({ { key = "bag", line = EQUIP } })
		local Chip = require("Chip")
		Chip.refresh()
		Hub.open("builds")
		click(Chip.frame)
		assert.is_true(Hub.isShown())
		assert.are.equal("gear", Hub.active)
	end)

	it("counts the other things waiting after the line, muted", function()
		boot()
		gearWaiting({ { key = "bag", line = EQUIP }, { key = "bag2", line = "Equip Band of Thorns +3.1" } })
		local Chip, Theme = require("Chip"), require("Theme")
		Chip.refresh()
		assert.are.equal(EQUIP, Chip.text:GetText())
		assert.are.equal("+1", Chip.moreCount:GetText())
		assert.is_true(Chip.frame:GetWidth() > Theme.SIZES.chip)
		assert.is_true(Chip.frame:GetWidth() <= Theme.SIZES.chipWide)
	end)

	it("says Forever Sixty with nothing waiting, and a click opens the last tab", function()
		boot()
		local Chip, Hub, Prefs, L = require("Chip"), require("Hub"), require("Prefs"), require("Locale")
		Prefs.set("hub", "tab", "builds")
		Chip.refresh()
		assert.are.equal(L.addonName, Chip.text:GetText())
		click(Chip.frame)
		assert.are.equal("builds", Hub.active)
	end)

	it("opens Settings on a right-click, and opens no window", function()
		boot()
		local Chip, Hub = require("Chip"), require("Hub")
		local asked = 0
		Chip.openSettings = function()
			asked = asked + 1
		end
		Chip.refresh()
		click(Chip.frame, "RightButton")
		assert.are.equal(1, asked)
		assert.is_false(Hub.isShown())
	end)

	it("says After combat and opens nothing in combat, with the hide switch off", function()
		boot({ shown = true, combatHide = false, byCharacter = {} })
		local Hub = gearWaiting({ { key = "bag", line = EQUIP } })
		local Chip, L = require("Chip"), require("Locale")
		Chip.setCombat(true)
		Chip.refresh()
		assert.are.equal(L.chipAfterCombat, Chip.text:GetText())
		assert.are.equal(Chip.COMBAT_ALPHA, Chip.current.alpha)
		Chip.click("LeftButton")
		assert.is_false(Hub.isShown())
		Chip.setCombat(false)
		Chip.refresh()
		assert.are.equal(EQUIP, Chip.text:GetText())
	end)

	it("hovers with every item and what a click opens", function()
		boot()
		gearWaiting({ { key = "bag", line = EQUIP } })
		local Chip, L = require("Chip"), require("Locale")
		Chip.refresh()
		local lines = Chip.lines(L.chipHintDrag)
		assert.are.equal("Forever Sixty", lines[1])
		assert.are.same({ text = EQUIP, right = "Gear" }, lines[2])
		assert.is_truthy(lines[3]:find("Click: open on Gear", 1, true))
	end)

	it("leads with tonight's raid once you answered, opening Guild, with no dot on the strip", function()
		local Options = boot()
		local state = { guild = { name = "Thornwake Company", rankName = "Raider", rankIndex = 3 } }
		_G.IsInGuild = function()
			return true
		end
		_G.GetGuildInfo = function()
			return state.guild.name, state.guild.rankName, state.guild.rankIndex
		end
		_G.ForeverSixtyInbox = { accepts = { "raid_answer" }, messages = { { type = "raid_signup", raid_id = 42,
			raid = "Onyxia's Lair", title = "Onyxia", starts_at = os.date("!%Y-%m-%dT%H:%M:%SZ", os.time() + 3600),
			answer = "in", answered_at = os.date("!%Y-%m-%dT%H:%M:%SZ", os.time() - 600) } } }
		helper.load("GuildRaid")
		helper.load("GuildHubTab")
		local Chip, Hub = require("Chip"), require("Hub")
		Chip.refresh(Options.data)
		assert.is_truthy(Chip.current.line:find("^Onyxia's Lair "))
		assert.is_truthy(Chip.current.line:find("In", 1, true))
		assert.are.equal("guild", Chip.current.opens)
		assert.are.same({}, Hub.waitingOf(Hub.tab("guild")))
		-- The hover says how long until the pull, under the raid's line.
		assert.are.equal("in 1 h 0 min", Chip.current.starts)
		_G.ForeverSixtyInbox, _G.IsInGuild, _G.GetGuildInfo = nil, nil, nil
	end)
end)

describe("the minimap button and the compartment entry, the chip's twins", function()
	after_each(function()
		mock.uninstall()
	end)

	it("open the same tab and show the same hover as the chip", function()
		boot()
		local Hub = gearWaiting({ { key = "bag", line = EQUIP } })
		local Chip, MinimapButton, Prefs, L = require("Chip"), require("Minimap"), require("Prefs"), require("Locale")
		Chip.refresh()
		assert.are.equal(Chip, MinimapButton.launcher)
		Prefs.set("minimap", "shown", true)
		local button = MinimapButton.refresh()
		click(button)
		assert.are.equal("gear", Hub.active)
		click(button)
		assert.is_false(Hub.isShown())
		MinimapButton.onCompartmentClick("ForeverSixty", "LeftButton")
		assert.are.equal("gear", Hub.active)
		local chipLines, buttonLines = Chip.lines(L.chipHintDrag), MinimapButton.tooltipLines(L.chipHintDragRing)
		for index = 1, #chipLines - 1 do
			assert.are.same(chipLines[index], buttonLines[index])
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
