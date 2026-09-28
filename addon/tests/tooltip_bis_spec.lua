-- addon/tests/tooltip_bis_spec.lua
-- The character-frame equipment slot hover (design/lane-bis-hover-addon.md
-- item 2): a header and the slot's leveling BiS pick, tagged "(equipped)"
-- or "(new at <band>)", advanced detail adding the source kind. Covers
-- Tooltip.bisLines (the pure section builder both hook paths share),
-- Tooltip.equippedSlotFor, Tooltip.bisItemLink and the two hook entry
-- points: the item hook (an equipped slot) and the slot-button OnEnter
-- hook (an empty one).
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

--- paladin, tab 1 (Holy) -- the same shape tooltip_spec.lua's DATA uses,
--- so Gear.specOf's tie-break (no ranks set) resolves "paladin-holy" with
--- no talent fixture needed.
local DATA = {
	build = "1.60.1.70009",
	classes = { paladin = { tabs = {
		{ name = "Holy", talents = {} },
		{ name = "Protection", talents = {} },
		{ name = "Retribution", talents = {} },
	} } },
	weights = {},
	bis = {
		["paladin-holy"] = {
			[10] = {
				alliance = { head = { 111, "Q" }, chest = { 333, "D" } },
			},
			[30] = {
				alliance = { head = { 222, "C" } },
			},
		},
	},
	bis_new = {
		["paladin-holy"] = {
			[30] = { alliance = { 222 } },
		},
	},
}

describe("Tooltip's BiS hover section", function()
	local Tooltip, Prefs

	local function start(install)
		install = install or {}
		install.class = install.class or { name = "Paladin", token = "PALADIN" }
		install.globals = install.globals or {}
		if install.globals.UnitFactionGroup == nil then
			install.globals.UnitFactionGroup = function()
				return "Alliance"
			end
		end
		mock.install(install)
		local Theme = helper.load("Theme")
		Theme.reset()
		helper.load("Export")
		helper.load("Gear")
		helper.load("Talents")
		Prefs = helper.load("Prefs")
		helper.load("Follow")
		Tooltip = helper.load("Tooltip")
		return Tooltip
	end

	after_each(function()
		mock.uninstall()
	end)

	describe("Tooltip.bisLines", function()
		it("says nothing with no data or no slot", function()
			start({ level = 20 })
			assert.are.same({}, Tooltip.bisLines(nil, "head"))
			assert.are.same({}, Tooltip.bisLines(DATA, nil))
		end)

		it("says nothing in combat", function()
			start({ level = 20, globals = { InCombatLockdown = function() return true end } })
			assert.are.same({}, Tooltip.bisLines(DATA, "head"))
		end)

		it("says nothing below the ladder's own floor", function()
			start({ level = 9 })
			assert.are.same({}, Tooltip.bisLines(DATA, "head"))
		end)

		it("shows the header and the item link in novice mode", function()
			start({ level = 20, globals = {
				GetItemInfo = function(id)
					if id == 111 then
						return "Helm of the Pathfinder", "item:111:link"
					end
					return nil
				end,
			} })
			local lines = Tooltip.bisLines(DATA, "head")
			assert.are.equal(string.format(L.tooltipBisHeader, 10, "paladin-holy"), lines[1])
			assert.are.equal("item:111:link", lines[2])
			assert.are.equal(2, #lines)
		end)

		it("rounds down to the highest band at or below the character's level", function()
			start({ level = 25, globals = {
				GetItemInfo = function() return nil end,
			} })
			-- Level 25 has no band of its own; the highest one <= 25 among
			-- {10, 30} is 10.
			local lines = Tooltip.bisLines(DATA, "head")
			assert.are.equal(string.format(L.tooltipBisHeader, 10, "paladin-holy"), lines[1])
		end)

		it("reaches band 30 once the level clears it", function()
			start({ level = 35, globals = { GetItemInfo = function() return nil end } })
			local lines = Tooltip.bisLines(DATA, "head")
			assert.are.equal(string.format(L.tooltipBisHeader, 30, "paladin-holy"), lines[1])
		end)

		it("falls back to the item:<id> form until GetItemInfo knows it", function()
			local requested = {}
			start({ level = 20, globals = {
				GetItemInfo = function() return nil end,
				RequestLoadItemDataByID = function(id) requested[#requested + 1] = id end,
			} })
			local lines = Tooltip.bisLines(DATA, "head")
			assert.are.equal("item:111", lines[2])
			assert.are.same({ 111 }, requested)
		end)

		it("asks the client for an uncached item exactly once", function()
			local requested = {}
			start({ level = 20, globals = {
				GetItemInfo = function() return nil end,
				RequestLoadItemDataByID = function(id) requested[#requested + 1] = id end,
			} })
			Tooltip.bisLines(DATA, "head")
			Tooltip.bisLines(DATA, "head")
			assert.are.same({ 111 }, requested)
		end)

		it("tags the pick equipped when it is what they wear", function()
			start({ level = 20, equipped = { [1] = "item:111:link" }, globals = {
				GetItemInfo = function() return nil, "item:111:link" end,
			} })
			local lines = Tooltip.bisLines(DATA, "head")
			assert.is_truthy(lines[2]:find(L.tooltipBisEquipped, 1, true))
		end)

		it("tags the pick new at band when it is in bis_new for this band and faction", function()
			start({ level = 35, globals = { GetItemInfo = function() return nil end } })
			local lines = Tooltip.bisLines(DATA, "head")
			assert.is_truthy(lines[2]:find(string.format(L.tooltipBisNew, 30), 1, true))
		end)

		it("carries neither tag for a pick that is worn nowhere and not new", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil end } })
			local lines = Tooltip.bisLines(DATA, "head")
			assert.are.equal("item:111", lines[2])
		end)

		it("says nothing for a slot the band names no pick for", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil end } })
			assert.are.same({}, Tooltip.bisLines(DATA, "legs"))
		end)

		it("adds no source line in novice mode", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil end } })
			local lines = Tooltip.bisLines(DATA, "head")
			assert.are.equal(2, #lines)
		end)

		it("adds the source kind only with advanced detail on", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil end } })
			Prefs.setFlag("advancedDetail", true)
			local lines = Tooltip.bisLines(DATA, "head")
			assert.are.equal(string.format(L.tooltipBisSource, "Quest"), lines[3])
		end)

		it("says nothing with no bis table for the resolved spec", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil end } })
			assert.are.same({}, Tooltip.bisLines({ build = "x", classes = DATA.classes, weights = {} }, "head"))
		end)
	end)

	describe("Tooltip.equippedSlotFor", function()
		it("names the slot whose equipped link matches exactly", function()
			start({ equipped = { [5] = "item:chest:link" } }) -- slot id 5 is chest
			assert.are.equal("chest", Tooltip.equippedSlotFor("item:chest:link"))
		end)

		it("answers nil for a link worn nowhere", function()
			start()
			assert.is_nil(Tooltip.equippedSlotFor("item:bag:link"))
		end)

		it("answers nil with no link at all", function()
			start()
			assert.is_nil(Tooltip.equippedSlotFor(nil))
		end)
	end)

	describe("the item hook, for an equipped slot", function()
		it("appends the BiS section after the existing tooltip lines", function()
			start({
				level = 20,
				equipped = { [1] = "item:111:link" }, -- slot id 1 is head
				globals = { GetItemInfo = function() return nil, "item:111:link" end },
			})
			Tooltip.data = DATA
			local calls = {}
			local tooltip = {
				AddLine = function(_, text) calls[#calls + 1] = text end,
				Show = function() end,
			}
			Tooltip.onTooltip(tooltip, "item:111:link")
			assert.are.equal(L.addonName, calls[1])
			assert.are.equal(string.format(L.tooltipBisHeader, 10, "paladin-holy"), calls[2])
			assert.is_truthy(calls[3]:find(L.tooltipBisEquipped, 1, true))
		end)

		it("adds nothing for a bag item that matches no equipped slot", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil end } })
			Tooltip.data = DATA
			local calls = {}
			Tooltip.onTooltip(
				{ AddLine = function(_, t) calls[#calls + 1] = t end, Show = function() end },
				"item:bag:link"
			)
			assert.are.same({}, calls)
		end)
	end)

	describe("the slot-button OnEnter hook, for an empty slot", function()
		local function headButton()
			return _G.CreateFrame("Button", "CharacterHeadSlot")
		end

		it("registers OnEnter on every button global this client has", function()
			start()
			local button = headButton()
			Tooltip.registerBisSlotButtons()
			assert.are.equal(1, mock.countCalls(button, "HookScript"))
			local call = mock.firstCall(button, "HookScript")
			assert.are.equal("OnEnter", call[1])
			assert.is_function(call[2])
		end)

		it("registers only once even if called again", function()
			start()
			local button = headButton()
			Tooltip.registerBisSlotButtons()
			Tooltip.registerBisSlotButtons()
			assert.are.equal(1, mock.countCalls(button, "HookScript"))
		end)

		it("draws the section onto GameTooltip for a genuinely empty slot", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil end } })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerBisSlotButtons()
			local handler = mock.firstCall(button, "HookScript")[2]
			handler(button)
			assert.are.equal(1, mock.countCalls(_G.GameTooltip, "SetOwner"))
			assert.are.equal(1, mock.countCalls(_G.GameTooltip, "Show"))
			-- This hook owns the whole tooltip (SetOwner + ClearLines, no
			-- Blizzard content ahead of it), unlike the item hook's shared
			-- tooltip, so it needs no "Forever Sixty" marker line -- the
			-- "Best in slot" header already says whose text this is.
			local first = mock.firstCall(_G.GameTooltip, "AddLine")
			assert.are.equal(string.format(L.tooltipBisHeader, 10, "paladin-holy"), first[1])
		end)

		it("does nothing when the slot actually has an item -- the item hook's job", function()
			start({
				level = 20,
				equipped = { [1] = "item:111:link" },
				globals = { GetItemInfo = function() return nil, "item:111:link" end },
			})
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerBisSlotButtons()
			local handler = mock.firstCall(button, "HookScript")[2]
			handler(button)
			assert.are.equal(0, mock.countCalls(_G.GameTooltip, "SetOwner"))
		end)

		it("does nothing while the tooltip pref is off", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil end } })
			Tooltip.data = DATA
			Prefs.setFlag("tooltip", false)
			local button = headButton()
			Tooltip.registerBisSlotButtons()
			mock.firstCall(button, "HookScript")[2](button)
			assert.are.equal(0, mock.countCalls(_G.GameTooltip, "SetOwner"))
		end)

		it("disables itself after one failure rather than erroring again", function()
			start({ level = 20, globals = {
				GetItemInfo = function() return nil end,
				InCombatLockdown = function() error("boom") end,
			} })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerBisSlotButtons()
			local handler = mock.firstCall(button, "HookScript")[2]
			handler(button)
			assert.is_true(Tooltip.slotDisabled)
			handler(button)
			assert.are.equal(0, mock.countCalls(_G.GameTooltip, "SetOwner"))
		end)
	end)

	describe("Tooltip.SOURCE_KIND_NAMES", function()
		it("covers exactly leveling-bis's own source kinds, one code each", function()
			-- Mirrors pipeline.addonbis.SOURCE_KIND_CODES' key set
			-- (test_every_source_kind_code_is_a_single_uppercase_letter);
			-- the two are pinned independently, one per side of the contract.
			local names = {}
			for code in pairs(Tooltip.SOURCE_KIND_NAMES) do
				names[#names + 1] = code
			end
			table.sort(names)
			assert.are.same({ "A", "C", "D", "P", "Q", "R", "W" }, names)
		end)
	end)
end)
