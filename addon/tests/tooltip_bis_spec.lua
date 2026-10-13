-- addon/tests/tooltip_bis_spec.lua
-- The character-frame equipment slot hovers. A slot holding an item is the
-- item hook's (the section with "Equipped", the plan, the next upgrade). An
-- EMPTY slot (L4.1 ruling 5) shows the same section with the band's pick for
-- that slot, printed as the worn item's Next line and source line with no gain
-- (ItemTooltip.emptySlotSection); the old empty-slot section (bisLines, its
-- icon row and BiS tag) is gone. Covers the three hook entry points: the item
-- hook (an equipped slot), the slot-button OnEnter hook (an empty one, first
-- hover), and GameTooltip's own OnTooltipCleared/OnShow (an empty one, every
-- hover after Blizzard's own OnUpdate re-render), and the compare tooltip.
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
			[20] = {
				alliance = {
					head = "111|Q",
					-- Data.lua's packed pick (task L1.2): item_id|kind|source|
					-- chance|alternatives..., decoded by Gear.pick.
					chest = "333|D|Wailing Caverns: Mutanus the Devourer|0.14|444|Q|A Quest",
				},
			},
			[30] = {
				alliance = { head = "222|C|Leatherworking" },
			},
		},
	},
}

--- DATA with stat weights for the spec, for the item hook's next upgrade (L3.3
--- fix round 2: the next upgrade is printed only with its gain, which needs
--- weights).
local WEIGHTED = setmetatable({ weights = { ["paladin-holy"] = { intellect = 1 } } }, { __index = DATA })

--- The client's names for the band picks.
local function namedItems(id)
	if id == 111 then
		return "Helm of the Pathfinder", "item:111:link"
	end
	if id == 333 then
		return "Mutant Scale Breastplate", "item:333:link"
	end
	return nil
end

describe("Tooltip's slot hovers", function()
	local Tooltip, Prefs, Theme

	--- A plain tooltip double, for the item hook's own tests: records both
	--- AddLine and AddDoubleLine calls in the order they happened.
	local function recordingTooltip()
		local calls = {}
		local tooltip = {
			AddLine = function(_, text) calls[#calls + 1] = { kind = "line", text = text } end,
			AddDoubleLine = function(_, left, right)
				calls[#calls + 1] = { kind = "doubleline", left = left, right = right }
			end,
			Show = function() end,
		}
		return tooltip, calls
	end

	--- The handler HookScript recorded for exactly `script` on `frame`.
	local function hookHandler(frame, script)
		for _, call in ipairs(frame.calls) do
			if call.method == "HookScript" and call[1] == script then
				return call[2]
			end
		end
		return nil
	end

	--- The texts GameTooltip was given through AddLine, in order.
	local function addedLines()
		local lines = {}
		for _, call in ipairs(_G.GameTooltip.calls) do
			if call.method == "AddLine" then
				lines[#lines + 1] = call[1]
			end
		end
		return lines
	end

	local function count(list, text)
		local total = 0
		for _, entry in ipairs(list) do
			if entry == text then
				total = total + 1
			end
		end
		return total
	end

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
		Theme = helper.load("Theme")
		Theme.reset()
		-- Compat's request-dedup table (Compat.requestedItems) is module
		-- state; reloading it fresh here keeps one test's requests from the next.
		helper.load("Compat")
		helper.load("Export")
		helper.load("Gear")
		helper.load("Talents")
		Prefs = helper.load("Prefs")
		Tooltip = helper.load("ItemTooltip")
		return Tooltip
	end

	after_each(function()
		mock.uninstall()
	end)

	describe("the item hook, for an equipped slot", function()
		it("appends the section after the brand header", function()
			start({
				level = 20,
				equipped = { [1] = "item:111:link" }, -- slot id 1 is head
				itemStats = { ["item:111:link"] = { __itemId = 111, __slot = "INVTYPE_HEAD" } },
				globals = { GetItemInfo = function() return nil, "item:111:link" end },
			})
			Tooltip.data = DATA
			local tooltip, calls = recordingTooltip()
			Tooltip.onTooltip(tooltip, "item:111:link")
			-- L3.3 fix round 1: a blank line, the header with "Equipped" on its
			-- right, then "Best until level 30" (the worn item is the band's pick) and
			-- the plan. L3.3 fix round 4: the blank row is a double line too, so
			-- "Equipped" can only ever sit on the header's own row.
			assert.are.same({ kind = "doubleline", left = " ", right = " " }, calls[1])
			assert.are.equal("doubleline", calls[2].kind)
			assert.is_truthy(calls[2].left:find(L.addonName, 1, true))
			assert.are.equal(L.tooltipEquipped, calls[2].right)
			assert.are.equal(string.format(L.tooltipBestUntil, 30), calls[3].text)
			assert.are.equal(string.format(L.tooltipPlanned, "Holy", 20), calls[4].text)
		end)

		-- Round-2 ruling 1, adopted: a bag item gets the slot's next pick too.
		it("draws the section for a bag item too, once its equip location maps to a tracked slot", function()
			start({
				level = 20,
				itemStats = {
					["item:bag:link"] = { __itemId = 999, __slot = "INVTYPE_HEAD" },
					["item:111:link"] = { __itemId = 111, __slot = "INVTYPE_HEAD", ITEM_MOD_INTELLECT_SHORT = 10 },
				},
				globals = { GetItemInfo = function(id)
					if id == 111 then
						return "Helm of the Pathfinder", "item:111:link"
					end
					return nil
				end },
			})
			Tooltip.data = WEIGHTED
			local tooltip, calls = recordingTooltip()
			Tooltip.onTooltip(tooltip, "item:bag:link")
			-- L3.3 fix round 2: a blank line, the header, then the slot's next
			-- upgrade with its gain (a next upgrade is never printed without one).
			assert.are.equal(" ", calls[1].text)
			assert.is_truthy(calls[2].text:find(L.addonName, 1, true))
			assert.are.equal(Tooltip.nextLine("Helm of the Pathfinder", 10, "score"), calls[3].text)
		end)

		it("still adds nothing for an item whose equip location this band does not track at all", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil end } })
			Tooltip.data = DATA
			local tooltip, calls = recordingTooltip()
			-- No itemStats entry at all: Compat.itemInfoInstant answers no
			-- equip location, same as a shirt or a tabard hover.
			Tooltip.onTooltip(tooltip, "item:bag:link")
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
			Tooltip.registerSlotButtons()
			assert.are.equal(1, mock.countCalls(button, "HookScript"))
			local call = mock.firstCall(button, "HookScript")
			assert.are.equal("OnEnter", call[1])
			assert.is_function(call[2])
		end)

		it("registers only once even if called again", function()
			start()
			local button = headButton()
			Tooltip.registerSlotButtons()
			Tooltip.registerSlotButtons()
			assert.are.equal(1, mock.countCalls(button, "HookScript"))
		end)

		it("names every paper-doll slot's client button, shirt and tabard aside", function()
			start()
			assert.are.equal("CharacterHeadSlot", Tooltip.SLOT_BUTTONS.head)
			assert.are.equal("CharacterSecondaryHandSlot", Tooltip.SLOT_BUTTONS.off_hand)
			assert.are.equal("CharacterRangedSlot", Tooltip.SLOT_BUTTONS.ranged)
			local total = 0
			for _ in pairs(Tooltip.SLOT_BUTTONS) do
				total = total + 1
			end
			assert.are.equal(17, total)
		end)

		it("draws the section onto GameTooltip: the slot's name, a blank line, the header, the Next line", function()
			start({ level = 20, globals = { GetItemInfo = namedItems, HEADSLOT = "Head" } })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerSlotButtons()
			mock.firstCall(button, "HookScript")[2](button)
			assert.are.equal(1, mock.countCalls(_G.GameTooltip, "SetOwner"))
			assert.are.equal(1, mock.countCalls(_G.GameTooltip, "Show"))
			local lines = addedLines()
			assert.are.equal("Head", lines[1])
			assert.are.equal(" ", lines[2])
			assert.is_truthy(lines[3]:find(L.addonName, 1, true))
			-- Exactly as the worn item prints its Next line, with no gain: nothing is worn.
			assert.are.equal(Tooltip.nextLine("Helm of the Pathfinder"), lines[4])
			assert.are.equal(4, #lines)
		end)

		it("prints the pick's source on its own muted line under the Next line", function()
			start({ level = 20, globals = { GetItemInfo = namedItems } })
			Tooltip.data = DATA
			local button = _G.CreateFrame("Button", "CharacterChestSlot")
			Tooltip.registerSlotButtons()
			mock.firstCall(button, "HookScript")[2](button)
			local lines = addedLines()
			assert.are.equal(Tooltip.nextLine("Mutant Scale Breastplate"), lines[3])
			assert.are.equal(string.format(L.tooltipSourceDrop, "Wailing Caverns", "Mutanus the Devourer", 14), lines[4])
		end)

		it("keeps the client's own tooltip when the client already drew it for this slot", function()
			start({ level = 20, globals = { GetItemInfo = namedItems } })
			Tooltip.data = DATA
			local button = headButton()
			_G.GameTooltip.GetOwner = function() return button end
			Tooltip.registerSlotButtons()
			mock.firstCall(button, "HookScript")[2](button)
			assert.are.equal(0, mock.countCalls(_G.GameTooltip, "SetOwner"))
			assert.are.equal(0, mock.countCalls(_G.GameTooltip, "ClearLines"))
			assert.are.equal(1, count(addedLines(), Tooltip.nextLine("Helm of the Pathfinder")))
		end)

		it("does nothing when the slot actually has an item -- the item hook's job", function()
			start({
				level = 20,
				equipped = { [1] = "item:111:link" },
				globals = { GetItemInfo = function() return nil, "item:111:link" end },
			})
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerSlotButtons()
			local handler = mock.firstCall(button, "HookScript")[2]
			handler(button)
			assert.are.equal(0, mock.countCalls(_G.GameTooltip, "SetOwner"))
		end)

		it("does nothing while the tooltip pref is off", function()
			start({ level = 20, globals = { GetItemInfo = namedItems } })
			Tooltip.data = DATA
			Prefs.setFlag("tooltip", false)
			local button = headButton()
			Tooltip.registerSlotButtons()
			mock.firstCall(button, "HookScript")[2](button)
			assert.are.equal(0, mock.countCalls(_G.GameTooltip, "SetOwner"))
		end)

		it("says nothing below the ladder's own floor", function()
			start({ level = 9, globals = { GetItemInfo = namedItems } })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerSlotButtons()
			mock.firstCall(button, "HookScript")[2](button)
			assert.are.equal(0, count(addedLines(), Tooltip.nextLine("Helm of the Pathfinder")))
		end)

		it("disables itself after one failure rather than erroring again", function()
			start({ level = 20, globals = { GetItemInfo = function() error("boom") end } })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerSlotButtons()
			local handler = mock.firstCall(button, "HookScript")[2]
			handler(button)
			assert.is_true(Tooltip.slotDisabled)
			local afterFirstFailure = mock.countCalls(_G.GameTooltip, "SetOwner")
			handler(button)
			assert.are.equal(afterFirstFailure, mock.countCalls(_G.GameTooltip, "SetOwner"))
		end)
	end)

	-- Blizzard's own PaperDollItemSlotButton_OnUpdate re-runs the hovered slot's
	-- OnEnter logic on a timer WITHOUT firing the button's OnEnter script,
	-- rebuilding GameTooltip with none of this addon's lines. GameTooltip's own
	-- OnTooltipCleared/OnShow fire for that rebuild.
	describe("the empty-slot flash fix", function()
		local function headButton()
			return _G.CreateFrame("Button", "CharacterHeadSlot")
		end

		it("hooks OnTooltipCleared and OnShow on GameTooltip, once", function()
			start()
			Tooltip.registerEmptySlotRefresh()
			Tooltip.registerEmptySlotRefresh()
			assert.is_function(hookHandler(_G.GameTooltip, "OnTooltipCleared"))
			assert.is_function(hookHandler(_G.GameTooltip, "OnShow"))
			assert.are.equal(2, mock.countCalls(_G.GameTooltip, "HookScript"))
		end)

		it("redraws the section after Blizzard's own OnUpdate rebuild clears it", function()
			start({ level = 20, globals = { GetItemInfo = namedItems } })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerSlotButtons()
			Tooltip.registerEmptySlotRefresh()
			mock.firstCall(button, "HookScript")[2](button) -- the first, ordinary hover
			local nextText = Tooltip.nextLine("Helm of the Pathfinder")
			assert.are.equal(1, count(addedLines(), nextText))

			_G.GameTooltip.GetOwner = function() return button end
			hookHandler(_G.GameTooltip, "OnTooltipCleared")(_G.GameTooltip)
			hookHandler(_G.GameTooltip, "OnShow")(_G.GameTooltip)

			assert.are.equal(2, count(addedLines(), nextText))
		end)

		it("never appends the section twice for one tooltip cycle", function()
			start({ level = 20, globals = { GetItemInfo = namedItems } })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerSlotButtons()
			Tooltip.registerEmptySlotRefresh()
			mock.firstCall(button, "HookScript")[2](button)

			-- OnShow firing again with no OnTooltipCleared in between (a
			-- mere re-anchor, not a rebuild) must not draw a second copy.
			_G.GameTooltip.GetOwner = function() return button end
			hookHandler(_G.GameTooltip, "OnShow")(_G.GameTooltip)

			assert.are.equal(1, count(addedLines(), Tooltip.nextLine("Helm of the Pathfinder")))
		end)

		it("does nothing on OnShow when the slot actually has an item", function()
			start({
				level = 20,
				equipped = { [1] = "item:111:link" },
				globals = { GetItemInfo = function() return nil, "item:111:link" end },
			})
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerEmptySlotRefresh()
			_G.GameTooltip.GetOwner = function() return button end
			hookHandler(_G.GameTooltip, "OnShow")(_G.GameTooltip)
			assert.are.equal(0, mock.countCalls(_G.GameTooltip, "AddLine"))
		end)

		it("does nothing for a tooltip owned by something other than a slot button", function()
			start({ level = 20 })
			Tooltip.data = DATA
			Tooltip.registerEmptySlotRefresh()
			_G.GameTooltip.GetOwner = function() return _G.UIParent end
			hookHandler(_G.GameTooltip, "OnShow")(_G.GameTooltip)
			assert.are.equal(0, mock.countCalls(_G.GameTooltip, "AddLine"))
		end)

		it("disables itself after one failure rather than erroring again", function()
			start({ level = 20, globals = { GetItemInfo = function() error("boom") end } })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerEmptySlotRefresh()
			_G.GameTooltip.GetOwner = function() return button end
			hookHandler(_G.GameTooltip, "OnShow")(_G.GameTooltip)
			assert.is_true(Tooltip.slotDisabled)
		end)
	end)

	-- The client's own item tooltip for the next upgrade, beside GameTooltip,
	-- redrawn once the client's data for it arrives.
	describe("the compare tooltip and its refresh", function()
		local function headButton()
			return _G.CreateFrame("Button", "CharacterHeadSlot")
		end

		it("shows the compare tooltip beside GameTooltip for an empty slot's pick", function()
			start({ level = 20, globals = { GetItemInfo = namedItems } })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerSlotButtons()
			mock.firstCall(button, "HookScript")[2](button)
			local compare = Theme.compareTooltipFrame
			assert.is_not_nil(compare)
			assert.are.equal(1, mock.countCalls(compare, "SetHyperlink"))
			assert.are.equal("item:111:link", mock.firstCall(compare, "SetHyperlink")[1])
			assert.are.equal(111, Tooltip.activeHover.itemId)
		end)

		it("shows no compare tooltip for the worn pick, which is best for now", function()
			start({
				level = 20,
				equipped = { [1] = "item:111:link" },
				itemStats = { ["item:111:link"] = { __itemId = 111, __slot = "INVTYPE_HEAD" } },
				globals = { GetItemInfo = function() return nil, "item:111:link" end },
			})
			Tooltip.data = DATA
			local tooltip = recordingTooltip()
			Tooltip.onTooltip(tooltip, "item:111:link")
			assert.is_nil(Theme.compareTooltipFrame)
			-- L3.3 fix round 4: the hover still waits for the next rung's pick name
			-- (its "At level 30" line), never for a compare item.
			assert.is_nil(Tooltip.activeHover.itemId)
		end)

		it("shows the next upgrade's compare tooltip, with its tab, once the client holds it", function()
			start({
				level = 20,
				itemStats = {
					["item:bag:link"] = { __itemId = 999, __slot = "INVTYPE_HEAD" },
					["item:111:link"] = { __itemId = 111, __slot = "INVTYPE_HEAD", ITEM_MOD_INTELLECT_SHORT = 10 },
				},
				globals = { GetItemInfo = function(id)
					if id == 111 then
						return "Helm of the Pathfinder", "item:111:link"
					end
					return nil
				end },
			})
			Tooltip.data = WEIGHTED
			Tooltip.onTooltip(recordingTooltip(), "item:bag:link")
			assert.are.equal("item:111:link", mock.firstCall(Theme.compareTooltipFrame, "SetHyperlink")[1])
			assert.are.equal(111, Tooltip.activeHover.itemId)
		end)

		it("redraws the empty-slot hover once the client's data for that id arrives", function()
			start({ level = 20, globals = { GetItemInfo = namedItems } })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerSlotButtons()
			mock.firstCall(button, "HookScript")[2](button)
			local setOwnerBefore = mock.countCalls(_G.GameTooltip, "SetOwner")
			Tooltip.onItemInfoReceived(111, true)
			assert.are.equal(setOwnerBefore + 1, mock.countCalls(_G.GameTooltip, "SetOwner"))
		end)

		it("does nothing for an id that is not the one currently on screen", function()
			start({ level = 20, globals = { GetItemInfo = namedItems } })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerSlotButtons()
			mock.firstCall(button, "HookScript")[2](button)
			local before = mock.countCalls(_G.GameTooltip, "SetOwner")
			Tooltip.onItemInfoReceived(999, true)
			assert.are.equal(before, mock.countCalls(_G.GameTooltip, "SetOwner"))
		end)

		it("does nothing when the client answers the request with a failure", function()
			start({ level = 20, globals = { GetItemInfo = namedItems } })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerSlotButtons()
			mock.firstCall(button, "HookScript")[2](button)
			local before = mock.countCalls(_G.GameTooltip, "SetOwner")
			Tooltip.onItemInfoReceived(111, false)
			assert.are.equal(before, mock.countCalls(_G.GameTooltip, "SetOwner"))
		end)

		it("registers GET_ITEM_INFO_RECEIVED exactly once and dispatches to onItemInfoReceived", function()
			start()
			local frame = Tooltip.registerItemInfoRefresh()
			Tooltip.registerItemInfoRefresh()
			assert.are.equal(1, mock.countCalls(frame, "RegisterEvent"))
			assert.are.equal("GET_ITEM_INFO_RECEIVED", mock.firstCall(frame, "RegisterEvent")[1])
		end)
	end)

	describe("the old empty-slot section", function()
		it("is gone: no bisLines, no BiS tag strings", function()
			start()
			assert.is_nil(Tooltip.bisLines)
			assert.is_nil(Tooltip.BIS_SLOT_BUTTONS)
			assert.is_nil(L.tooltipBisDefaultTag)
			assert.is_nil(L.tooltipBisEquipped)
			assert.is_nil(L.tooltipBisNew)
			assert.is_nil(L.tooltipBisItemLevel)
			assert.is_nil(L.tooltipBisSourceLabel)
		end)
	end)

	describe("Tooltip.SOURCE_KIND_NAMES", function()
		it("covers exactly leveling-bis's own source kinds, one code each", function()
			-- Mirrors pipeline.addonbis.SOURCE_KIND_CODES' key set
			-- (test_every_source_kind_code_is_a_single_uppercase_letter);
			-- the two are pinned independently, one per side of the contract.
			start()
			local names = {}
			for code in pairs(Tooltip.SOURCE_KIND_NAMES) do
				names[#names + 1] = code
			end
			table.sort(names)
			assert.are.same({ "A", "B", "C", "D", "P", "Q", "R", "V", "W" }, names)
		end)
	end)
end)
