-- addon/tests/player_tooltip_spec.lua
-- The player tooltip (spec section 4.6, L4.4): after the client's own lines, the Forever
-- Score in its tier colour, a guildmate's live line, the components while Shift is held
-- (a muted hint otherwise, which also names /fs link), "no logs yet" for a character the
-- record does not have, and nothing for a non-player. It takes over the unit hook from the
-- item tooltip, registered by the same rule. Folds in what ratings_spec.lua covered for the
-- tooltip. The client-frame recorder stays empty in every example.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local DOT = "\194\183" -- the middle dot in L.exportTreeSeparator

local QUALITY = {
	[1] = { r = 1, g = 1, b = 1 }, [2] = { r = 0.12, g = 1, b = 0 }, [3] = { r = 0, g = 0.44, b = 0.87 },
	[4] = { r = 0.64, g = 0.21, b = 0.93 }, [5] = { r = 1, g = 0.5, b = 0 },
}

local function record()
	return {
		format = 2,
		generated = os.date("!%Y-%m-%d"),
		characters = {
			["Bow Jackzon-ClassicBetaPvP"] = {
				score = 1240, tier = "rare", as_of = "2026-11-04",
				components = { gear = 380, experience = 410, parses = 300, attendance = 100, discipline = 50 },
				fights = 9, profile = "/character/us/pvp/bow-jackzon",
			},
			["Facep Alm-ClassicBetaPvP"] = { score = 0, tier = "common", as_of = "2026-11-04", fights = 0 },
			["Bob-Whitemane"] = { score = 700, tier = "uncommon", as_of = "2026-11-03", fights = 2 },
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

	--- The lines AddLine drew on a fresh tooltip for `unit`: { text, r, g, b } each.
	local function hover(unit)
		local tooltip = _G.CreateFrame("GameTooltip")
		PlayerTooltip.onUnitTooltip(tooltip, unit or "mouseover")
		local lines = {}
		for _, call in ipairs(tooltip.calls) do
			if call.method == "AddLine" then
				lines[#lines + 1] = { text = call[1], r = call[2], g = call[3], b = call[4] }
			end
		end
		return lines, tooltip
	end

	local function rgbOf(hexKey)
		local r, g, b = Theme.rgb(Theme.HEX[hexKey])
		return { r, g, b }
	end

	describe("a logged character", function()
		it("leads with the score, its tier and date, in the tier's colour", function()
			start(record())
			local lines, tooltip = hover()
			assert.are.equal("Forever Score 1240 " .. DOT .. " Rare " .. DOT .. " as of 4 Nov", lines[1].text)
			assert.are.same({ 0, 0.44, 0.87 }, { lines[1].r, lines[1].g, lines[1].b })
			assert.are.equal(1, mock.countCalls(tooltip, "Show"))
		end)

		it("hides the components behind Shift, with a muted hint that names /fs link", function()
			start(record())
			local lines = hover()
			assert.are.equal(2, #lines)
			assert.are.equal(L.scoreShiftHint .. L.exportTreeSeparator .. L.scoreLinkHint, lines[2].text)
			assert.are.equal("Shift: score details " .. DOT .. " /fs link: site page", lines[2].text)
			assert.are.same(rgbOf("tooltipMuted"), { lines[2].r, lines[2].g, lines[2].b })
		end)

		it("lists the components while Shift is held, and drops the hint", function()
			start(record())
			shift = true
			local lines = hover()
			assert.are.equal(6, #lines)
			assert.are.equal("Gear 380", lines[2].text)
			assert.are.equal("Experience 410", lines[3].text)
			assert.are.equal("Discipline 50", lines[6].text)
			for index = 2, 6 do
				assert.are.same(rgbOf("tooltipText"), { lines[index].r, lines[index].g, lines[index].b })
			end
		end)

		it("offers no Shift for a record without components, only the link", function()
			start(record())
			names.mouseover = { "Bob", "Whitemane" }
			local lines = hover()
			assert.are.equal("Forever Score 700 " .. DOT .. " Uncommon " .. DOT .. " as of 3 Nov", lines[1].text)
			assert.are.equal(L.scoreLinkHint, lines[2].text)
			shift = true
			assert.are.equal(2, #hover())
		end)

		it("draws an unknown tier uncoloured, in the tooltip's text colour", function()
			local data = record()
			data.characters["Bow Jackzon-ClassicBetaPvP"].tier = "mythic"
			start(data)
			local lines = hover()
			assert.are.same(rgbOf("tooltipText"), { lines[1].r, lines[1].g, lines[1].b })
		end)
	end)

	describe("no logs", function()
		it("reads no logs yet, muted, for a character the record does not have", function()
			start(record())
			names.mouseover = { "Smoke Gouda", "" }
			local lines = hover()
			assert.are.equal(1, #lines)
			assert.are.equal("Forever Score: no logs yet", lines[1].text)
			assert.are.same(rgbOf("tooltipMuted"), { lines[1].r, lines[1].g, lines[1].b })
		end)

		it("never shows a zero for a character with no logs", function()
			start(record())
			names.mouseover = { "Facep Alm", "" }
			local lines = hover()
			assert.are.equal(1, #lines)
			assert.are.equal(L.scoreNoLogs, lines[1].text)
			assert.is_nil(lines[1].text:find("0", 1, true))
		end)

		it("says nothing about a score when the data addon is not installed or cannot be read", function()
			start(nil)
			assert.are.same({}, hover())
			start({ format = 3, generated = "2026-11-04", characters = {}, guilds = {} })
			assert.are.same({}, hover())
		end)
	end)

	describe("the live line", function()
		it("follows the score, muted, while a guildmate's hello is fresh", function()
			start(record())
			Comms.seen["Bow Jackzon-ClassicBetaPvP"] = {
				level = 24, spec = "marksmanship", gearGap = 12, time = Theme.now(), channel = "GUILD",
			}
			local lines = hover()
			assert.are.equal(3, #lines)
			assert.are.equal("Level 24 Marksmanship " .. DOT .. " gap 12 (live)", lines[2].text)
			assert.are.same(rgbOf("tooltipMuted"), { lines[2].r, lines[2].g, lines[2].b })
			assert.are.equal(L.scoreShiftHint .. L.exportTreeSeparator .. L.scoreLinkHint, lines[3].text)
		end)

		it("is current before any log exists", function()
			start(record())
			names.mouseover = { "Smoke Gouda", "" }
			Comms.seen["Smoke Gouda-ClassicBetaPvP"] = { level = 19, spec = "fury", gearGap = 3, time = Theme.now() }
			local lines = hover()
			assert.are.same({ L.scoreNoLogs, "Level 19 Fury " .. DOT .. " gap 3 (live)" },
				{ lines[1].text, lines[2].text })
		end)

		it("matches a sender heard by first name on a client with last names", function()
			start(record(), { globals = { UnitFullName = function() return "Bow", "Jackzon" end } })
			names.mouseover = { "Smoke", "Gouda" }
			Comms.seen["Smoke-ClassicBetaPvP"] = { level = 19, spec = "fury", gearGap = 3, time = Theme.now() }
			local lines = hover()
			assert.are.equal("Level 19 Fury " .. DOT .. " gap 3 (live)", lines[2].text)
		end)

		it("reads this character by its full name on a client with last names", function()
			start(record(), { globals = {
				UnitFullName = function() return "Bow", "Jackzon" end,
				UnitIsUnit = function(a, b) return a == b or (a == "target" and b == "player") end,
			} })
			names.target = { "Bow" }
			players.target = true
			assert.are.equal("Forever Score 1240 " .. DOT .. " Rare " .. DOT .. " as of 4 Nov", hover("player")[1].text)
			assert.are.equal("Forever Score 1240 " .. DOT .. " Rare " .. DOT .. " as of 4 Nov", hover("target")[1].text)
		end)

		it("ignores a sender on another realm and one heard too long ago", function()
			start(record())
			Comms.seen["Bow Jackzon-Whitemane"] = { level = 30, spec = "fury", gearGap = 1, time = Theme.now() }
			assert.are.equal(2, #hover())
			Comms.seen = { ["Bow Jackzon-ClassicBetaPvP"] = {
				level = 24, spec = "marksmanship", gearGap = 12, time = Theme.now() - Comms.SEEN_FOR - 1,
			} }
			assert.are.equal(2, #hover())
		end)

		it("leaves out a spec the sender has not got", function()
			start(record())
			Comms.seen["Bow Jackzon-ClassicBetaPvP"] = { level = 5, spec = "-", gearGap = 0, time = Theme.now() }
			assert.are.equal("Level 5 " .. DOT .. " gap 0 (live)", hover()[2].text)
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

	it("reads the score from Score, the one reader", function()
		start(record())
		assert.are.same(Score.lines(Score.forCharacter("Bow Jackzon"))[1], hover()[1].text)
	end)
end)
