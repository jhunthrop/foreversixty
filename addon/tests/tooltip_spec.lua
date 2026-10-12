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

describe("Tooltip", function()
	local Theme, Tooltip

	local function start(install)
		mock.install(install or {})
		Theme = helper.load("Theme")
		Theme.reset()
		helper.load("Export")
		helper.load("Gear")
		helper.load("Prefs")
		Tooltip = helper.load("ItemTooltip")
		return Tooltip
	end

	after_each(function()
		mock.uninstall()
	end)

	-- L3.3: the second argument is ItemTooltip.current (Gear.currentWeights), whose
	-- `caps` come from the sim's weights message.
	describe("Tooltip.capLines", function()
		it("names a capped stat the item actually carries", function()
			start({
				itemStats = {
					["item:ring"] = { __itemId = 400, __slot = "INVTYPE_FINGER", ITEM_MOD_HIT_RATING_SHORT = 12 },
				},
			})
			local lines = Tooltip.capLines("item:ring", { caps = { "hit" } })
			assert.are.same({ string.format(L.tooltipCapped, "hit") }, lines)
		end)

		it("says nothing for a stat this item does not carry, even if capped", function()
			start({
				itemStats = {
					["item:ring"] = { __itemId = 400, __slot = "INVTYPE_FINGER", ITEM_MOD_CRIT_RATING_SHORT = 12 },
				},
			})
			assert.are.same({}, Tooltip.capLines("item:ring", { caps = { "hit" } }))
		end)

		it("names every capped stat the item carries, sorted", function()
			start({
				itemStats = {
					["item:ring"] = {
						__itemId = 400, __slot = "INVTYPE_FINGER",
						ITEM_MOD_HIT_RATING_SHORT = 8, ITEM_MOD_EXPERTISE_RATING_SHORT = 6,
					},
				},
			})
			local lines = Tooltip.capLines("item:ring", { caps = { "hit", "expertise" } })
			assert.are.same({
				string.format(L.tooltipCapped, "expertise"),
				string.format(L.tooltipCapped, "hit"),
			}, lines)
		end)

		it("says nothing with no current weights at all", function()
			start({ itemStats = { ["item:ring"] = { __itemId = 400, __slot = "INVTYPE_FINGER" } } })
			assert.are.same({}, Tooltip.capLines("item:ring", nil))
		end)

		it("says nothing for current weights with no caps", function()
			start({ itemStats = { ["item:ring"] = { __itemId = 400, __slot = "INVTYPE_FINGER" } } })
			assert.are.same({}, Tooltip.capLines("item:ring", { spec = "fury" }))
		end)

		it("says nothing with no item link", function()
			start()
			assert.are.same({}, Tooltip.capLines(nil, { caps = { "hit" } }))
		end)
	end)

	describe("the hook", function()
		local Prefs

		local function startHook(install)
			start(install)
			-- require, not helper.load: start() already loaded Prefs fresh,
			-- and Tooltip captured that exact instance at its own load time.
			-- Reloading it here would hand this describe block a second,
			-- disconnected copy that Tooltip never sees.
			Prefs = require("Prefs")
			return Prefs
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

		it("adds a Forever Sixty heading and the verdict onto the tooltip", function()
			startHook({
				class = { name = "Paladin", token = "PALADIN" },
				itemStats = {
					["item:111"] = { __itemId = 111, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 10 },
				},
			})
			Tooltip.data = DATA
			local calls = {}
			local tooltip = {
				AddLine = function(_, text) calls[#calls + 1] = text end,
				Show = function() end,
			}
			Tooltip.onTooltip(tooltip, "item:111")
			-- L3.3: a blank line, the header, then the verdict (nothing worn: it fills the slot).
			assert.are.equal(" ", calls[1])
			assert.is_truthy(calls[2]:find(L.addonName, 1, true))
			assert.are.equal(string.format(L.tooltipFillsScore, 10), calls[3])
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
			Tooltip.data = DATA
			local tooltip = { AddLine = function() end, Show = function() end }
			Tooltip.onTooltip(tooltip, "item:111")
			local before = Tooltip.cache["item:111"]
			Tooltip.onTooltip(tooltip, "item:111")
			assert.are.equal(before, Tooltip.cache["item:111"])
		end)

		it("disables itself after one failure rather than erroring again", function()
			startHook({
				class = { name = "Paladin", token = "PALADIN" },
				itemStats = {
					["item:111"] = { __itemId = 111, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 10 },
				},
			})
			-- A verdict, so sections is non-empty and addLines actually
			-- reaches AddLine, which is what this example throws from.
			Tooltip.data = DATA
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
			-- OnTooltipSetItem, OnTooltipSetUnit (both legacy, no
			-- processor here), and OnTooltipCleared/OnShow (item 2's
			-- empty-slot-flash fix, registerEmptySlotRefresh) -- four
			-- GameTooltip hooks total, each exactly once.
			assert.are.equal(4, mock.countCalls(_G.GameTooltip, "HookScript"))
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
				format = 2, generated = os.date("!%Y-%m-%d"),
				characters = { ["Bob-Ashbringer"] = { score = rating, tier = "epic", as_of = "2026-11-04", fights = 5 } },
				guilds = {},
			} or nil
			Theme = helper.load("Theme")
			Theme.reset()
			helper.load("Export")
			helper.load("Gear")
			helper.load("Prefs")
			helper.load("Ratings")
			Tooltip = helper.load("ItemTooltip")
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
			assert.is_truthy(call[1]:find("Forever Score 77", 1, true))
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
