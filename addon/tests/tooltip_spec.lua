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
		Tooltip = helper.load("Tooltip")
		return Tooltip
	end

	after_each(function()
		mock.uninstall()
	end)

	describe("Tooltip.verdict", function()
		it("scores an equippable item against what is worn, as a percentage of it", function()
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
			local verdict = Tooltip.verdict(DATA, "item:candidate")
			assert.are.equal("upgrade", verdict.kind)
			assert.are.equal(string.format(L.tooltipVerdictUpgrade, 300.0), verdict.text)
		end)

		it("scores a downgrade as a percentage too, never a bare negative number", function()
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
			local verdict = Tooltip.verdict(DATA, "item:candidate")
			assert.are.equal("downgrade", verdict.kind)
			assert.are.equal(string.format(L.tooltipVerdictDowngrade, 80.0), verdict.text)
		end)

		it("calls a near-equal delta a Sidegrade rather than a coin-flip", function()
			start({
				class = { name = "Paladin", token = "PALADIN" },
				equipped = { [5] = "item:equipped" },
				itemStats = {
					["item:candidate"] = {
						__itemId = 200, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 100,
					},
					["item:equipped"] = {
						__itemId = 100, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 100,
					},
				},
			})
			local verdict = Tooltip.verdict(DATA, "item:candidate")
			assert.are.equal("sidegrade", verdict.kind)
			assert.are.equal(L.tooltipVerdictSidegrade, verdict.text)
		end)

		it("falls back to absolute points with nothing worn to be a percentage of", function()
			start({
				class = { name = "Paladin", token = "PALADIN" },
				itemStats = {
					["item:111"] = { __itemId = 111, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 10 },
				},
			})
			local verdict = Tooltip.verdict(DATA, "item:111")
			assert.are.equal("upgrade", verdict.kind)
			assert.are.equal(string.format(L.tooltipUpgrade, "chest", 10), verdict.text)
		end)

		it("says nothing for an item with no equip slot at all", function()
			start({
				class = { name = "Paladin", token = "PALADIN" },
				itemStats = { ["item:reagent"] = { __itemId = 300, __slot = "" } },
			})
			assert.is_nil(Tooltip.verdict(DATA, "item:reagent"))
		end)

		it("says nothing with no stat weights for the spec at any band", function()
			start({ class = { name = "Paladin", token = "PALADIN" } })
			assert.is_nil(Tooltip.verdict({ build = "x", classes = DATA.classes, weights = {} },
				"item:candidate"))
		end)

		it("says nothing rather than erroring with no data table at all", function()
			start({ class = { name = "Paladin", token = "PALADIN" } })
			assert.is_nil(Tooltip.verdict(nil, "item:candidate"))
		end)

		it("prefers the nightly band weights over the static table when they exist", function()
			start({
				class = { name = "Paladin", token = "PALADIN" },
				level = 20,
				equipped = { [5] = "item:equipped" },
				itemStats = {
					["item:candidate"] = {
						__itemId = 200, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 10,
					},
					["item:equipped"] = {
						__itemId = 100, __slot = "INVTYPE_CHEST", ITEM_MOD_SPIRIT_SHORT = 10,
					},
				},
			})
			-- Equal under the static table (10 intellect vs 10 spirit both
			-- worth 1.0 -- a Sidegrade), but the band table weights spirit
			-- five times higher, which only shows up if the band table is
			-- the one actually used.
			local data = {
				build = DATA.build,
				classes = DATA.classes,
				weights = { ["paladin-holy"] = { intellect = 1.0, spirit = 1.0 } },
				bis_weights = { ["paladin-holy"] = { [20] = { intellect = 1.0, spirit = 5.0 } } },
			}
			assert.are.equal("downgrade", Tooltip.verdict(data, "item:candidate").kind)
		end)

		it("falls back to the static table when the band has not measured this spec yet", function()
			start({
				class = { name = "Paladin", token = "PALADIN" },
				level = 20,
				equipped = { [5] = "item:equipped" },
				itemStats = {
					["item:candidate"] = {
						__itemId = 200, __slot = "INVTYPE_CHEST", ITEM_MOD_INTELLECT_SHORT = 10,
					},
					["item:equipped"] = {
						__itemId = 100, __slot = "INVTYPE_CHEST", ITEM_MOD_SPIRIT_SHORT = 10,
					},
				},
			})
			local data = {
				build = DATA.build,
				classes = DATA.classes,
				weights = { ["paladin-holy"] = { intellect = 1.0, spirit = 1.0 } },
				bis_weights = {},
			}
			assert.are.equal("sidegrade", Tooltip.verdict(data, "item:candidate").kind)
		end)

		it("gets a bigger-agility chest obviously right, the owner's own screenshot case", function()
			-- Tunic of Westfall (2041): agility 11, stamina 5. Dark Leather
			-- Tunic (2317, hunter-beast-mastery's own band 20 BiS chest and
			-- what this fixture wears): agility 6. hunter-beast-mastery's
			-- own measured band 20 weights (data/builds/1.60.1.70009/bis/
			-- hunter-beast-mastery.json) put agility at 2.045 and do not
			-- weight stamina at all.
			start({
				class = { name = "Hunter", token = "HUNTER" },
				level = 20,
				equipped = { [5] = "item:2317" },
				itemStats = {
					["item:2041"] = {
						__itemId = 2041, __slot = "INVTYPE_CHEST",
						ITEM_MOD_AGILITY_SHORT = 11, ITEM_MOD_STAMINA_SHORT = 5,
					},
					["item:2317"] = {
						__itemId = 2317, __slot = "INVTYPE_CHEST", ITEM_MOD_AGILITY_SHORT = 6,
					},
				},
			})
			local data = {
				build = "1.60.1.70009",
				classes = { hunter = { tabs = {
					{ name = "Beast Mastery", talents = {} },
					{ name = "Marksmanship", talents = {} },
					{ name = "Survival", talents = {} },
				} } },
				weights = {},
				bis_weights = { ["hunter-beast-mastery"] = { [20] = { agility = 2.045 } } },
			}
			local verdict = Tooltip.verdict(data, "item:2041")
			assert.are.equal("upgrade", verdict.kind)
			local expectedPercent = (11 * 2.045 - 6 * 2.045) / (6 * 2.045) * 100
			assert.are.equal(string.format(L.tooltipVerdictUpgrade, expectedPercent), verdict.text)
		end)
	end)

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

		it("says nothing with no weights message at all", function()
			start({ itemStats = { ["item:ring"] = { __itemId = 400, __slot = "INVTYPE_FINGER" } } })
			assert.are.same({}, Tooltip.capLines("item:ring", nil))
		end)

		it("says nothing for a weights message with no caps", function()
			start({ itemStats = { ["item:ring"] = { __itemId = 400, __slot = "INVTYPE_FINGER" } } })
			assert.are.same({}, Tooltip.capLines("item:ring", { spec = "fury" }))
		end)

		it("says nothing with no item link", function()
			start()
			assert.are.same({}, Tooltip.capLines(nil, { caps = { "hit" } }))
		end)
	end)

	describe("Tooltip.sections", function()
		it("folds the verdict and the capped-stat call-outs into one op list", function()
			start({
				class = { name = "Paladin", token = "PALADIN" },
				equipped = { [5] = nil },
				itemStats = {
					["item:111"] = {
						__itemId = 111, __slot = "INVTYPE_CHEST",
						ITEM_MOD_INTELLECT_SHORT = 10, ITEM_MOD_HIT_RATING_SHORT = 5,
					},
				},
			})
			local ops = Tooltip.sections(DATA, "item:111", { caps = { "hit" } })
			assert.are.equal(2, #ops)
			assert.are.equal("line", ops[1].kind)
			assert.are.equal(string.format(L.tooltipCapped, "hit"), ops[2].text)
		end)

		it("returns an empty list rather than nil for an item with nothing to say", function()
			start({ class = { name = "Paladin", token = "PALADIN" } })
			assert.are.same({}, Tooltip.sections(DATA, "item:999"))
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
			assert.is_truthy(calls[1]:find(L.addonName, 1, true))
			assert.are.equal(string.format(L.tooltipUpgrade, "chest", 10), calls[2])
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
				format = 1, generated = os.date("!%Y-%m-%dT%H:%M:%SZ"), build = "1.60.1.69893",
				characters = { ["us:ashbringer:bob"] = { rating = rating, fights = 5 } }, guilds = {},
			} or nil
			Theme = helper.load("Theme")
			Theme.reset()
			helper.load("Export")
			helper.load("Gear")
			helper.load("Prefs")
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
