-- addon/tests/player_tooltip_spec.lua
-- The player tooltip (design/specs/2026-11-08-addon-ux.md section 2.5): after the client's own
-- lines, "Forever Score 1240" in the tier's colour, the date only once the record is more than a
-- week old, the five parts as double lines while Shift is held, and nothing at all for a
-- character with no record or no rated fights, or a non-player. No hint, no live line, no total.
-- It takes over the unit hook from the item tooltip, registered by the same rule. The
-- client-frame recorder stays empty in every example.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local QUALITY = {
	[1] = { r = 1, g = 1, b = 1 }, [2] = { r = 0.12, g = 1, b = 0 }, [3] = { r = 0, g = 0.44, b = 0.87 },
	[4] = { r = 0.64, g = 0.21, b = 0.93 }, [5] = { r = 1, g = 0.5, b = 0 },
}

--- Today's date as the record writes it: a record from today shows no date.
local TODAY = os.date("!%Y-%m-%d")

local function record()
	return {
		format = 2,
		generated = TODAY,
		characters = {
			["Bow Jackzon-ClassicBetaPvP"] = {
				score = 1240, tier = "rare", as_of = TODAY,
				components = { gear = 380, experience = 410, parses = 300, attendance = 100, discipline = 50 },
				fights = 9, profile = "/character/us/pvp/bow-jackzon",
			},
			["Facep Alm-ClassicBetaPvP"] = { score = 0, tier = "common", as_of = TODAY, fights = 0 },
			["Bob-Whitemane"] = { score = 700, tier = "uncommon", as_of = "2020-10-13", fights = 2 },
		},
		guilds = {},
	}
end
describe("PlayerTooltip", function()
	local state, PlayerTooltip, Prefs, Comms, Theme, Score
	-- What UnitName answers for each unit token: { first, second }.
	local names, players, shift

	local function start(data, install)
		install = install or {}
		install.class = { name = "Hunter", token = "HUNTER" }
		install.realm = "Classic Beta PvP"
		install.region = install.region or 1
		install.playerName = "Bow"
		install.globals = install.globals or {}
		install.globals.ITEM_QUALITY_COLORS = QUALITY
		install.globals.UnitIsPlayer = function(unit)
			return players[unit] == true
		end
		install.globals.IsShiftKeyDown = function()
			return shift
		end
		state = mock.install(install)
		_G.UnitName = function(unit)
			local entry = names[unit] or { "Tester" }
			return entry[1], entry[2]
		end
		_G.ForeverSixtyData = data
		_G.ForeverSixtyDB = {}
		Theme = require("Theme")
		Theme.reset()
		Prefs = helper.load("Prefs")
		helper.load("Export")
		Score = helper.load("Score")
		Comms = helper.load("Comms")
		PlayerTooltip = helper.load("PlayerTooltip")
		return PlayerTooltip
	end

	before_each(function()
		names = { mouseover = { "Bow Jackzon", "" }, player = { "Bow" } }
		players = { mouseover = true, player = true }
		shift = false
	end)

	after_each(function()
		assert.are.same({}, state.clientWrites)
		_G.ForeverSixtyData = nil
		mock.uninstall()
	end)

	--- The lines AddLine and AddDoubleLine drew on a fresh tooltip for `unit`, in order:
	--- { text, r, g, b } each, plus { right, rr, rg, rb } for a double line.
	local function hover(unit)
		local tooltip = _G.CreateFrame("GameTooltip")
		PlayerTooltip.onUnitTooltip(tooltip, unit or "mouseover")
		local lines = {}
		for _, call in ipairs(tooltip.calls) do
			if call.method == "AddLine" then
				lines[#lines + 1] = { text = call[1], r = call[2], g = call[3], b = call[4] }
			elseif call.method == "AddDoubleLine" then
				lines[#lines + 1] = { text = call[1], right = call[2], r = call[3], g = call[4], b = call[5],
					rr = call[6], rg = call[7], rb = call[8] }
			end
		end
		return lines, tooltip
	end

	local function rgbOf(hexKey)
		local r, g, b = Theme.rgb(Theme.HEX[hexKey])
		return { r, g, b }
	end

	describe("a logged character", function()
		it("is one line, the score in the tier's colour, no tier word and no date while it is recent", function()
			start(record())
			local lines, tooltip = hover()
			assert.are.equal(1, #lines)
			assert.are.equal(string.format(L.scoreLine, 1240), lines[1].text)
			assert.are.equal("Forever Score 1240", lines[1].text)
			assert.are.same({ Theme.rgb(require("Tokens").color.rarity.rareText) }, { lines[1].r, lines[1].g, lines[1].b, 1 })
			assert.are.equal(1, mock.countCalls(tooltip, "Show"))
		end)

		it("dates a stale record and mutes it, as the chat badge does: Forever Score 700 (13 Oct)", function()
			start(record())
			names.mouseover = { "Bob", "Whitemane" }
			local lines = hover()
			assert.are.equal(1, #lines)
			assert.are.equal(string.format(L.scoreDated, 700, "13 Oct"), lines[1].text)
			assert.are.same(rgbOf("tooltipMuted"), { lines[1].r, lines[1].g, lines[1].b })
		end)

		it("dates on the record's own as_of by the one staleness rule, after three days and not before", function()
			start(record())
			local base = { score = 900, as_of = "2026-11-01" }
			local day = 24 * 60 * 60
			local midnight = Score.dateSeconds("2026-11-01")
			assert.are.equal(3 * day, Score.STALE_AFTER_SECONDS)
			assert.is_nil(PlayerTooltip.DATED_AFTER_SECONDS)
			assert.are.equal("Forever Score 900", PlayerTooltip.scoreText(base, midnight + 3 * day))
			assert.are.equal("Forever Score 900 (1 Nov)", PlayerTooltip.scoreText(base, midnight + 3 * day + 1))
			assert.are.equal("Forever Score 900", PlayerTooltip.scoreText({ score = 900 }, midnight + 30 * day))
		end)

		it("lists the five parts as double lines while Shift is held, in the parts' own words", function()
			start(record())
			shift = true
			local lines = hover()
			assert.are.equal(6, #lines)
			local expected = {
				{ "Gear", "380" }, { "Experience", "410" }, { "Parses", "300" }, { "Attendance", "100" },
				{ L.scoreDiscipline, "50" },
			}
			for index, pair in ipairs(expected) do
				local entry = lines[index + 1]
				assert.are.same(pair, { entry.text, entry.right }, "component " .. index)
				assert.are.same(rgbOf("tooltipMuted"), { entry.r, entry.g, entry.b })
				assert.are.same(rgbOf("tooltipBody"), { entry.rr, entry.rg, entry.rb })
			end
		end)

		it("prints no total and no hint, with or without Shift", function()
			start(record())
			for _, held in ipairs({ false, true }) do
				shift = held
				for _, entry in ipairs(hover()) do
					assert.is_nil(entry.text:find("Shift", 1, true))
					assert.is_nil(entry.text:find("/fs link", 1, true))
					assert.is_nil((entry.right or ""):find("=", 1, true))
				end
			end
		end)

		it("offers no Shift block for a record without components", function()
			start(record())
			names.mouseover = { "Bob", "Whitemane" }
			shift = true
			assert.are.equal(1, #hover())
		end)

		it("draws an unknown tier in the tooltip's white", function()
			local data = record()
			data.characters["Bow Jackzon-ClassicBetaPvP"].tier = "mythic"
			start(data)
			local lines = hover()
			assert.are.same(rgbOf("tooltipBody"), { lines[1].r, lines[1].g, lines[1].b })
		end)

		it("reads this character by its full name on a client with last names", function()
			start(record(), { globals = {
				UnitFullName = function() return "Bow", "Jackzon" end,
				UnitIsUnit = function(a, b) return a == b or (a == "target" and b == "player") end,
			} })
			names.target = { "Bow" }
			players.target = true
			assert.are.equal("Forever Score 1240", hover("player")[1].text)
			assert.are.equal("Forever Score 1240", hover("target")[1].text)
		end)
	end)

	describe("nothing to say", function()
		it("says nothing for a character the record does not have", function()
			start(record())
			names.mouseover = { "Smoke Gouda", "" }
			assert.are.same({}, hover())
			shift = true
			assert.are.same({}, hover())
		end)

		it("never shows a zero for a character with no rated fights", function()
			start(record())
			names.mouseover = { "Facep Alm", "" }
			assert.are.same({}, hover())
		end)

		it("says nothing when the data addon is not installed or cannot be read", function()
			start(nil)
			assert.are.same({}, hover())
			start({ format = 3, generated = "2026-11-04", characters = {}, guilds = {} })
			assert.are.same({}, hover())
		end)

		it("is pure: no record, no lines", function()
			start(record())
			assert.are.same({}, PlayerTooltip.section(nil, true, 0))
			assert.are.same({}, PlayerTooltip.section({ tier = "rare" }, false, 0))
		end)

		it("draws no live line even while a guildmate's hello is fresh", function()
			start(record())
			Comms.seen["Bow Jackzon-ClassicBetaPvP"] = {
				level = 24, spec = "marksmanship", gearGap = 12, time = Theme.now(), channel = "GUILD",
			}
			assert.are.equal(1, #hover())
		end)
	end)

	describe("the guard", function()
		it("adds nothing for a non-player", function()
			start(record())
			players.mouseover = false
			assert.are.same({}, hover())
		end)

		it("obeys Gear tips on item tooltips: off, nothing at all", function()
			start(record())
			Prefs.setFlag("tooltip", false)
			assert.are.same({}, hover())
		end)

		it("adds nothing without a unit", function()
			start(record())
			local tooltip = _G.CreateFrame("GameTooltip")
			PlayerTooltip.onUnitTooltip(tooltip, nil)
			assert.are.equal(0, mock.countCalls(tooltip, "AddLine"))
		end)

		it("turns itself off after a failure instead of raising, noting it once", function()
			start(record())
			_G.UnitName = function()
				error("boom")
			end
			hover()
			assert.is_true(PlayerTooltip.disabled)
			assert.are.equal(1, #Theme.diagnostics())
			assert.are.same({}, hover())
		end)
	end)

	describe("the hook", function()
		it("reads the unit GetUnit hands back, and nil without one", function()
			start(record())
			assert.are.equal("mouseover", PlayerTooltip.unitFrom({ GetUnit = function() return "Bow", "mouseover" end }))
			assert.is_nil(PlayerTooltip.unitFrom({}))
			assert.is_nil(PlayerTooltip.unitFrom({ GetUnit = function() error("boom") end }))
		end)

		it("hooks the unit tooltip through the processor when the client has it", function()
			local kinds = {}
			start(record(), { globals = {
				TooltipDataProcessor = { AddTooltipPostCall = function(kind) kinds[#kinds + 1] = kind end },
				Enum = { TooltipDataType = { Item = 1, Unit = 2 } },
			} })
			assert.are.equal("processor", PlayerTooltip.registerUnit())
			assert.are.same({ 2 }, kinds)
			assert.are.equal(0, mock.countCalls(_G.GameTooltip, "HookScript"))
		end)

		it("falls back to OnTooltipSetUnit with no unit processor, once", function()
			start(record(), { globals = {
				TooltipDataProcessor = { AddTooltipPostCall = function() end },
				Enum = { TooltipDataType = { Item = 1 } },
			} })
			assert.is_false(PlayerTooltip.hasUnitProcessor())
			assert.are.equal("legacy", PlayerTooltip.registerUnit())
			assert.are.equal("legacy", PlayerTooltip.registerUnit())
			local call = mock.firstCall(_G.GameTooltip, "HookScript")
			assert.are.equal("OnTooltipSetUnit", call[1])
			assert.are.equal(1, mock.countCalls(_G.GameTooltip, "HookScript"))
		end)

		it("draws through the hook it registered", function()
			local hooked
			start(record(), { globals = {
				TooltipDataProcessor = { AddTooltipPostCall = function(_, fn) hooked = fn end },
				Enum = { TooltipDataType = { Unit = 2 } },
			} })
			PlayerTooltip.registerUnit()
			local tooltip = _G.CreateFrame("GameTooltip")
			tooltip.GetUnit = function()
				return "Bow Jackzon", "mouseover"
			end
			hooked(tooltip)
			assert.is_truthy(mock.firstCall(tooltip, "AddLine")[1]:find("Forever Score 1240", 1, true))
		end)

		it("leaves the item tooltip's registration to items alone", function()
			local kinds = {}
			start(record(), { globals = {
				TooltipDataProcessor = { AddTooltipPostCall = function(kind) kinds[#kinds + 1] = kind end },
				Enum = { TooltipDataType = { Item = 1, Unit = 2 } },
			} })
			helper.load("Gear")
			helper.load("ItemTooltip").register()
			assert.are.same({ 1 }, kinds)
			assert.is_nil(require("ItemTooltip").onUnitTooltip)
			assert.is_nil(require("ItemTooltip").registerUnit)
		end)
	end)

	describe("Shift while hovering", function()
		local function showing(unit)
			_G.GameTooltip.shown = true
			_G.GameTooltip.GetUnit = function()
				return "Bow Jackzon", unit
			end
		end

		it("redraws the hovered player's tooltip when Shift goes down or up", function()
			start(record())
			showing("mouseover")
			assert.is_true(PlayerTooltip.onModifierChanged("LSHIFT", 1))
			assert.are.equal("mouseover", mock.lastCall(_G.GameTooltip, "SetUnit")[1])
			assert.is_true(PlayerTooltip.onModifierChanged("RSHIFT", 0))
			assert.are.equal(2, mock.countCalls(_G.GameTooltip, "SetUnit"))
		end)

		it("leaves the tooltip alone for another key, a hidden tooltip or a non-player", function()
			start(record())
			showing("mouseover")
			assert.is_false(PlayerTooltip.onModifierChanged("LCTRL", 1))
			players.mouseover = false
			assert.is_false(PlayerTooltip.onModifierChanged("LSHIFT", 1))
			players.mouseover = true
			_G.GameTooltip.shown = false
			assert.is_false(PlayerTooltip.onModifierChanged("LSHIFT", 1))
			Prefs.setFlag("tooltip", false)
			_G.GameTooltip.shown = true
			assert.is_false(PlayerTooltip.onModifierChanged("LSHIFT", 1))
			assert.are.equal(0, mock.countCalls(_G.GameTooltip, "SetUnit"))
		end)

		it("listens on a frame of its own, registered once", function()
			start(record())
			local frame = PlayerTooltip.register()
			assert.are.equal(frame, PlayerTooltip.register())
			assert.is_true(frame.events.MODIFIER_STATE_CHANGED)
			assert.is_nil(frame.name)
		end)
	end)

	describe("the site link", function()
		it("is the record's own profile on the site where it has one", function()
			local data = record()
			data.characters["Bow Jackzon-ClassicBetaPvP"].profile = "/character/eu/pvp/bow-jackzon"
			start(data)
			assert.are.equal("foreversixty.gg/character/eu/pvp/bow-jackzon",
				PlayerTooltip.siteLink("Bow Jackzon", "Classic Beta PvP"))
		end)

		it("is built from the region, the realm's ruleset and the name alone otherwise", function()
			start(record())
			assert.are.equal("foreversixty.gg/character/us/pvp/smoke-gouda", PlayerTooltip.siteLink("Smoke Gouda"))
			assert.are.equal("foreversixty.gg/character/us/pvp/smoke-gouda",
				PlayerTooltip.siteLink("Smoke Gouda", "ClassicBetaPvP"))
			assert.are.equal("foreversixty.gg/character/us/pvp/facep-alm", PlayerTooltip.siteLink("Facep Alm"))
			start(nil)
			assert.are.equal("foreversixty.gg/character/us/normal/bob", PlayerTooltip.siteLink("Bob", "Whitemane"))
		end)

		it("takes the region from the client, us where it cannot name one", function()
			start(record(), { region = 3 })
			assert.are.equal("foreversixty.gg/character/eu/pvp/smoke-gouda", PlayerTooltip.siteLink("Smoke Gouda"))
			start(record(), { region = 72 })
			assert.are.equal("foreversixty.gg/character/us/pvp/smoke-gouda", PlayerTooltip.siteLink("Smoke Gouda"))
		end)
	end)

	it("reads the record from Score, the one reader", function()
		start(record())
		local record_ = Score.forCharacter("Bow Jackzon")
		assert.are.same(PlayerTooltip.scoreText(record_, Score.clock()), hover()[1].text)
	end)
end)
