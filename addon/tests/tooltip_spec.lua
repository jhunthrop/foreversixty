local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local DATA = {
	build = "1.60.1.69893",
	classes = { paladin = { tabs = {
		{ name = "Holy", talents = {} },
		{ name = "Protection", talents = {} },
		{ name = "Retribution", talents = {} },
	} } },
	weights = { ["paladin-holy"] = { intellect = 1.0, spirit = 0.5 } },
}

local BUILD = {
	classSlug = "paladin",
	order = {},
	gear = { { slot = "chest", itemId = 111, stats = { intellect = 10 } } },
}

describe("Tooltip", function()
	local Theme, Tooltip

	local function start(install)
		mock.install(install or {})
		Theme = helper.load("Theme")
		Theme.reset()
		helper.load("Export")
		helper.load("Gear")
		helper.load("Prefs")
		helper.load("Follow")
		Tooltip = helper.load("Tooltip")
		return Tooltip
	end

	after_each(function()
		mock.uninstall()
	end)

	it("names the slot a planned item is for", function()
		start()
		assert.are.equal(string.format(L.tooltipPlanned, "chest"), Tooltip.plannedLine(BUILD, 111))
	end)

	it("says nothing for an item the build does not want", function()
		start()
		assert.is_nil(Tooltip.plannedLine(BUILD, 999))
	end)

	it("says nothing with no build loaded", function()
		start()
		assert.is_nil(Tooltip.plannedLine(nil, 111))
	end)

	it("scores an equippable item against what is worn in its slot", function()
		start({
			class = { name = "Paladin", token = "PALADIN" },
			equipped = { [5] = "item:equipped" }, -- slot id 5 is chest
			itemStats = {
				["item:candidate"] = {
					__itemId = 200, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 20,
				},
				["item:equipped"] = {
					__itemId = 100, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 5,
				},
			},
		})
		assert.are.equal(string.format(L.tooltipUpgrade, "chest", 15),
			Tooltip.upgradeLine(DATA, "item:candidate"))
	end)

	it("says not an upgrade rather than a negative number", function()
		start({
			class = { name = "Paladin", token = "PALADIN" },
			equipped = { [5] = "item:equipped" },
			itemStats = {
				["item:candidate"] = {
					__itemId = 200, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 1,
				},
				["item:equipped"] = {
					__itemId = 100, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 5,
				},
			},
		})
		assert.are.equal(L.tooltipNotUpgrade, Tooltip.upgradeLine(DATA, "item:candidate"))
	end)

	it("says nothing for an item with no equip slot at all", function()
		start({
			class = { name = "Paladin", token = "PALADIN" },
			itemStats = { ["item:reagent"] = { __itemId = 300, __slot = "" } },
		})
		assert.is_nil(Tooltip.upgradeLine(DATA, "item:reagent"))
	end)

	it("says nothing with no stat weights for the spec", function()
		start({ class = { name = "Paladin", token = "PALADIN" } })
		assert.is_nil(Tooltip.upgradeLine({ build = "x", classes = DATA.classes, weights = {} },
			"item:candidate"))
	end)

	it("says nothing rather than erroring with no data table at all", function()
		start({ class = { name = "Paladin", token = "PALADIN" } })
		assert.is_nil(Tooltip.upgradeLine(nil, "item:candidate"))
	end)

	it("combines both lines when both apply", function()
		start({
			class = { name = "Paladin", token = "PALADIN" },
			equipped = { [5] = nil },
			itemStats = {
				["item:111"] = { __itemId = 111, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 10 },
			},
		})
		local lines = Tooltip.lines(DATA, BUILD, "item:111")
		assert.are.equal(2, #lines)
		assert.are.equal(string.format(L.tooltipPlanned, "chest"), lines[1])
	end)

	it("returns an empty list rather than nil for an item with nothing to say", function()
		start({ class = { name = "Paladin", token = "PALADIN" } })
		assert.are.same({}, Tooltip.lines(DATA, nil, "item:999"))
	end)

	describe("the hook", function()
		local Prefs, Follow

		local function startHook(install)
			start(install)
			-- require, not helper.load: start() already loaded Prefs and
			-- Follow fresh (in that order, before Tooltip), and Tooltip
			-- captured those exact instances at its own load time.
			-- Reloading them here would hand this describe block a second,
			-- disconnected copy that Tooltip never sees.
			Prefs = require("Prefs")
			Follow = require("Follow")
			return Prefs, Follow
		end

		it("reads the link GetItem hands back", function()
			startHook()
			local tooltip = { GetItem = function() return "Item Name", "item:42" end }
			assert.are.equal("item:42", Tooltip.itemLinkFrom(tooltip))
		end)

		it("returns nil rather than erroring when GetItem is missing", function()
			startHook()
			assert.is_nil(Tooltip.itemLinkFrom({}))
		end)

		it("adds a Forever Sixty heading and the lines onto the tooltip", function()
			startHook({
				class = { name = "Paladin", token = "PALADIN" },
				itemStats = {
					["item:111"] = { __itemId = 111, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 10 },
				},
			})
			Tooltip.data = DATA
			Follow.build = BUILD
			local calls = {}
			local tooltip = {
				AddLine = function(_, text) calls[#calls + 1] = text end,
				Show = function() end,
			}
			Tooltip.onTooltip(tooltip, "item:111")
			assert.are.equal(L.addonName, calls[1])
			assert.are.equal(string.format(L.tooltipPlanned, "chest"), calls[2])
		end)

		it("adds nothing when the tooltip pref is off", function()
			startHook({ class = { name = "Paladin", token = "PALADIN" } })
			Prefs.setFlag("tooltip", false)
			local calls = {}
			Tooltip.onTooltip({ AddLine = function(_, t) calls[#calls + 1] = t end, Show = function() end },
				"item:111")
			assert.are.same({}, calls)
		end)

		it("never recomputes for the same link twice", function()
			startHook({
				class = { name = "Paladin", token = "PALADIN" },
				itemStats = {
					["item:111"] = { __itemId = 111, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 10 },
				},
			})
			local tooltip = { AddLine = function() end, Show = function() end }
			Tooltip.onTooltip(tooltip, "item:111")
			local before = Tooltip.cache["item:111"]
			Tooltip.onTooltip(tooltip, "item:111")
			assert.are.equal(before, Tooltip.cache["item:111"])
		end)

		it("disables itself after one failure rather than erroring again", function()
			startHook({ class = { name = "Paladin", token = "PALADIN" } })
			-- A planned item, so lines is non-empty and addLines actually
			-- reaches AddLine, which is what this example throws from.
			Follow.build = BUILD
			local tooltip = {
				AddLine = function() error("boom") end,
				Show = function() end,
			}
			Tooltip.onTooltip(tooltip, "item:111")
			assert.is_true(Tooltip.disabled)
			assert.are.equal(1, #Theme.diagnostics())
			local calls = 0
			tooltip.AddLine = function() calls = calls + 1 end
			Tooltip.onTooltip(tooltip, "item:111")
			assert.are.equal(0, calls)
		end)

		it("prefers the modern processor when the client has both", function()
			startHook({ globals = {
				TooltipDataProcessor = { AddTooltipPostCall = function() end },
				Enum = { TooltipDataType = { Item = 1 } },
			} })
			assert.is_true(Tooltip.hasProcessor())
		end)

		it("has no processor when Enum.TooltipDataType.Item is missing", function()
			startHook({ globals = { TooltipDataProcessor = { AddTooltipPostCall = function() end } } })
			assert.is_false(Tooltip.hasProcessor())
		end)

		it("registers through the processor when the client has one", function()
			local captured
			startHook({ globals = {
				TooltipDataProcessor = {
					AddTooltipPostCall = function(kind, fn) captured = { kind, fn } end,
				},
				Enum = { TooltipDataType = { Item = 1 } },
			} })
			assert.are.equal("processor", Tooltip.register())
			assert.are.equal(1, captured[1])
			assert.is_function(captured[2])
		end)

		it("falls back to the legacy hooks, item and unit, with no processor", function()
			startHook({ class = { name = "Paladin", token = "PALADIN" } })
			assert.are.equal("legacy", Tooltip.register())
			local scripts = {}
			for _, call in ipairs(_G.GameTooltip.calls) do
				if call.method == "HookScript" then
					scripts[call[1]] = true
					assert.is_function(call[2])
				end
			end
			assert.is_true(scripts.OnTooltipSetItem)
			assert.is_true(scripts.OnTooltipSetUnit)
		end)

		it("registers each hook only once", function()
			startHook()
			Tooltip.register()
			Tooltip.register()
			assert.are.equal(2, mock.countCalls(_G.GameTooltip, "HookScript"))
		end)

		it("hooks the unit tooltip through the processor when the client has it", function()
			local kinds = {}
			startHook({ globals = {
				TooltipDataProcessor = {
					AddTooltipPostCall = function(kind) kinds[#kinds + 1] = kind end,
				},
				Enum = { TooltipDataType = { Item = 1, Unit = 2 } },
			} })
			Tooltip.register()
			assert.are.same({ 2, 1 }, kinds)
		end)
	end)

	describe("the unit line", function()
		local function withData(rating)
			-- The data global must exist before Ratings and Tooltip load, in
			-- that order, so Tooltip captures the Ratings that reads it.
			mock.install({ class = { name = "Paladin", token = "PALADIN" }, realm = "Ashbringer", region = 1 })
			_G.ForeverSixtyData = rating and {
				format = 1, generated = os.date("!%Y-%m-%dT%H:%M:%SZ"), build = "1.60.1.69893",
				characters = { ["us:ashbringer:bob"] = { rating = rating, fights = 5 } }, guilds = {},
			} or nil
			Theme = helper.load("Theme")
			Theme.reset()
			helper.load("Export")
			helper.load("Gear")
			helper.load("Prefs")
			helper.load("Follow")
			helper.load("Ratings")
			Tooltip = helper.load("Tooltip")
			_G.UnitIsPlayer = function() return true end
			_G.UnitName = function(unit)
				if unit == "mouseover" then return "Bob", "" end
				return "Me", ""
			end
		end

		after_each(function()
			_G.ForeverSixtyData = nil
			_G.UnitIsPlayer = nil
		end)

		it("adds a gold rating line for a rated player", function()
			withData(77)
			local tooltip = _G.CreateFrame("GameTooltip")
			Tooltip.onUnitTooltip(tooltip, "mouseover")
			local call = mock.firstCall(tooltip, "AddLine")
			assert.is_truthy(call[1]:find("rating 77", 1, true))
		end)

		it("adds nothing for an unrated player or without the data addon", function()
			withData(nil)
			local tooltip = _G.CreateFrame("GameTooltip")
			Tooltip.onUnitTooltip(tooltip, "mouseover")
			assert.are.equal(0, mock.countCalls(tooltip, "AddLine"))
		end)

		it("turns itself off after a failure instead of raising", function()
			withData(77)
			_G.UnitName = function() error("boom") end
			local tooltip = _G.CreateFrame("GameTooltip")
			Tooltip.onUnitTooltip(tooltip, "mouseover")
			assert.is_true(Tooltip.unitDisabled)
			Tooltip.onUnitTooltip(tooltip, "mouseover")
		end)
	end)
end)
