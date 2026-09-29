-- addon/tests/tooltip_bis_spec.lua
-- The character-frame equipment slot hover (design/lane-bis-hover-addon.md
-- item 2, restyled by lane addon-tooltip-polish item 1 onto one
-- AddDoubleLine row with a right-aligned tag and a muted source line), and
-- the empty-slot flash fix (lane addon-tooltip-polish item 2). Covers
-- Tooltip.bisLines (the pure section builder both hook paths share),
-- Tooltip.equippedSlotFor, Tooltip.bisItemLink and the three hook entry
-- points: the item hook (an equipped slot), the slot-button OnEnter hook
-- (an empty one, first hover), and GameTooltip's own OnTooltipCleared/
-- OnShow (an empty one, every hover after Blizzard's own OnUpdate
-- re-render).
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
					head = { 111, "Q" },
					chest = { 333, "D", "Wailing Caverns: Mutanus the Devourer" },
				},
			},
			[30] = {
				alliance = { head = { 222, "C", "Leatherworking" } },
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
	local Tooltip, Prefs, Theme

	--- The icon prefix bisLines puts on the item line when the client has
	--- not (or has not yet) told Compat.itemIcon an icon for it -- every
	--- fixture below leaves GetItemIconByID/GetItemIcon unmocked, so this
	--- is what every item line in this file starts with.
	local function unknownIcon()
		return Theme.inlineIcon(nil, Tooltip.BIS_ICON_SIZE)
	end

	--- A plain tooltip double, for the item hook's own tests: records both
	--- AddLine and AddDoubleLine calls in the order they happened, so a
	--- spec can check the brand header (a "line" op) and the BiS row (a
	--- "doubleline" op) sit where expected relative to each other.
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

	--- The handler HookScript recorded for exactly `script` on `frame` --
	--- mock.firstCall only ever answers the first HookScript call
	--- regardless of which script it named, which is not enough once a
	--- spec hooks more than one script on the same frame (GameTooltip's
	--- own OnTooltipCleared and OnShow, item 2's own fix).
	local function hookHandler(frame, script)
		for _, call in ipairs(frame.calls) do
			if call.method == "HookScript" and call[1] == script then
				return call[2]
			end
		end
		return nil
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
		-- state, same as Tooltip's own used to be before it moved there;
		-- reloading it fresh here is what makes "asks the client exactly
		-- once" mean once per test rather than once ever.
		helper.load("Compat")
		helper.load("Export")
		helper.load("Gear")
		helper.load("Talents")
		Prefs = helper.load("Prefs")
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

		it("shows the pick as a doubleline row with the default BiS tag", function()
			start({ level = 20, globals = {
				GetItemInfo = function(id)
					if id == 111 then
						return "Helm of the Pathfinder", "item:111:link"
					end
					return nil
				end,
			} })
			local ops = Tooltip.bisLines(DATA, "head")
			assert.are.equal(1, #ops)
			assert.are.equal("doubleline", ops[1].kind)
			assert.are.equal(unknownIcon() .. " item:111:link", ops[1].left)
			assert.are.equal(string.format(L.tooltipBisDefaultTag, 20), ops[1].right)
			assert.are.equal("muted", ops[1].rightColor)
		end)

		it("rounds down to the highest band at or below the character's level", function()
			start({ level = 25, globals = {
				GetItemInfo = function() return nil end,
			} })
			-- Level 25 has no band of its own; the highest one <= 25 among
			-- {20, 30} is 20, whose head pick is item 111.
			local ops = Tooltip.bisLines(DATA, "head")
			assert.are.equal(unknownIcon() .. " item:111", ops[1].left)
		end)

		it("reaches band 30 once the level clears it", function()
			start({ level = 35, globals = { GetItemInfo = function() return nil end } })
			-- Band 30's own head pick is item 222, not band 20's item 111.
			local ops = Tooltip.bisLines(DATA, "head")
			assert.are.equal(unknownIcon() .. " item:222", ops[1].left)
		end)

		it("falls back to the item:<id> form until GetItemInfo knows it", function()
			local requested = {}
			start({ level = 20, globals = {
				GetItemInfo = function() return nil end,
				RequestLoadItemDataByID = function(id) requested[#requested + 1] = id end,
			} })
			local ops = Tooltip.bisLines(DATA, "head")
			assert.are.equal(unknownIcon() .. " item:111", ops[1].left)
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

		it("tags the pick equipped, in green, when it is what they wear", function()
			start({ level = 20, equipped = { [1] = "item:111:link" }, globals = {
				GetItemInfo = function() return nil, "item:111:link" end,
			} })
			local ops = Tooltip.bisLines(DATA, "head")
			assert.are.equal(L.tooltipBisEquipped, ops[1].right)
			assert.are.equal("success", ops[1].rightColor)
		end)

		it("tags the pick new at band, in gold, when it is in bis_new for this band and faction", function()
			start({ level = 35, globals = { GetItemInfo = function() return nil end } })
			local ops = Tooltip.bisLines(DATA, "head")
			assert.are.equal(string.format(L.tooltipBisNew, 30), ops[1].right)
			assert.are.equal("gold", ops[1].rightColor)
		end)

		it("says nothing for a slot the band names no pick for", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil end } })
			assert.are.same({}, Tooltip.bisLines(DATA, "legs"))
		end)

		it("has only the doubleline row with nothing extra cached about the item", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil end } })
			-- head/111 carries no source label in this fixture, and
			-- GetItemInfo has no item level for it yet.
			local ops = Tooltip.bisLines(DATA, "head")
			assert.are.equal(1, #ops)
		end)

		it("says nothing with no bis table for the resolved spec", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil end } })
			assert.are.same({}, Tooltip.bisLines({ build = "x", classes = DATA.classes, weights = {} }, "head"))
		end)

		it("adds the item level once GetItemInfo has it", function()
			start({ level = 20, globals = {
				GetItemInfo = function() return "Helm", "item:111:link", 2, 45 end,
			} })
			local ops = Tooltip.bisLines(DATA, "head")
			assert.are.equal(2, #ops)
			assert.are.equal("line", ops[2].kind)
			assert.are.equal(string.format(L.tooltipBisItemLevel, 45), ops[2].text)
			assert.are.equal("muted", ops[2].color)
		end)

		it("adds the muted source line whenever the pick carries a label", function()
			-- chest/333 carries "Wailing Caverns: Mutanus the Devourer",
			-- source_kind "D" -- shown in novice mode too now, not gated
			-- behind advanced detail (lane addon-tooltip-polish item 1).
			start({ level = 20, globals = { GetItemInfo = function() return nil end } })
			local ops = Tooltip.bisLines(DATA, "chest")
			assert.are.equal(2, #ops)
			assert.are.equal(
				string.format(L.tooltipBisSourceLabel, "Dungeon", "Wailing Caverns · Mutanus the Devourer"),
				ops[2].text
			)
			assert.are.equal("muted", ops[2].color)
		end)

		it("adds no source line for a pick with no label at all", function()
			-- head/111 carries only a source kind code, no label.
			start({ level = 20, globals = { GetItemInfo = function() return "Helm", "item:111:link", 2, 45 end } })
			local ops = Tooltip.bisLines(DATA, "head")
			assert.are.equal(2, #ops) -- doubleline + item level, no source line
		end)

		it("puts the item level before the source line when both are present", function()
			start({ level = 20, globals = {
				GetItemInfo = function(id)
					if id == 333 then
						return "Robe", "item:333:link", 2, 45
					end
					return nil
				end,
			} })
			local ops = Tooltip.bisLines(DATA, "chest")
			assert.are.equal(3, #ops)
			assert.are.equal(string.format(L.tooltipBisItemLevel, 45), ops[2].text)
			assert.are.equal(
				string.format(L.tooltipBisSourceLabel, "Dungeon", "Wailing Caverns · Mutanus the Devourer"),
				ops[3].text
			)
		end)

		it("puts the recommended item's own icon inline, the neutral tile until the client knows it", function()
			start({ level = 20, globals = {
				GetItemInfo = function() return nil end,
				GetItemIconByID = function(id) return id == 111 and "Interface\\Icons\\Real" or nil end,
			} })
			local ops = Tooltip.bisLines(DATA, "head")
			assert.are.equal(Theme.inlineIcon("Interface\\Icons\\Real", Tooltip.BIS_ICON_SIZE) .. " item:111",
				ops[1].left)
		end)

		it("returns the recommended item id and link as a second value, for the compare tooltip", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil, "item:111:link" end } })
			local _, target = Tooltip.bisLines(DATA, "head")
			assert.are.equal(111, target.itemId)
			assert.are.equal("item:111:link", target.link)
		end)

		it("answers no second value on every path that shows no section", function()
			start({ level = 9 })
			local _, target = Tooltip.bisLines(DATA, "head")
			assert.is_nil(target)
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
		it("appends the BiS section after the brand header", function()
			start({
				level = 20,
				equipped = { [1] = "item:111:link" }, -- slot id 1 is head
				globals = { GetItemInfo = function() return nil, "item:111:link" end },
			})
			Tooltip.data = DATA
			local tooltip, calls = recordingTooltip()
			Tooltip.onTooltip(tooltip, "item:111:link")
			assert.are.equal("line", calls[1].kind)
			assert.is_truthy(calls[1].text:find(L.addonName, 1, true))
			assert.are.equal("doubleline", calls[2].kind)
			assert.are.equal(unknownIcon() .. " item:111:link", calls[2].left)
			assert.are.equal(L.tooltipBisEquipped, calls[2].right)
		end)

		it("adds nothing for a bag item that matches no equipped slot", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil end } })
			Tooltip.data = DATA
			local tooltip, calls = recordingTooltip()
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

		it("draws the section, with the brand header, onto GameTooltip for a genuinely empty slot", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil end } })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerBisSlotButtons()
			local handler = mock.firstCall(button, "HookScript")[2]
			handler(button)
			assert.are.equal(1, mock.countCalls(_G.GameTooltip, "SetOwner"))
			assert.are.equal(1, mock.countCalls(_G.GameTooltip, "Show"))
			-- lane addon-tooltip-polish: the empty-slot section now carries
			-- the same brand header as the equipped-item hook, for one
			-- consistent look everywhere this addon draws on a tooltip
			-- (docs/tenets.md) -- the old "Best in slot" header text that
			-- used to make this line unnecessary is gone, replaced by the
			-- row's own tag.
			local first = mock.firstCall(_G.GameTooltip, "AddLine")
			assert.is_truthy(first[1]:find(L.addonName, 1, true))
			local row = mock.firstCall(_G.GameTooltip, "AddDoubleLine")
			assert.are.equal(unknownIcon() .. " item:111", row[1])
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
			-- The first call's own SetOwner already happened (it comes
			-- before the guarded bisLines call that actually throws); the
			-- second call is what must add no more.
			local afterFirstFailure = mock.countCalls(_G.GameTooltip, "SetOwner")
			handler(button)
			assert.are.equal(afterFirstFailure, mock.countCalls(_G.GameTooltip, "SetOwner"))
		end)
	end)

	-- lane addon-tooltip-polish item 2: Blizzard's own
	-- PaperDollItemSlotButton_OnUpdate re-runs the hovered slot's OnEnter
	-- logic on a timer WITHOUT firing the button's OnEnter script at all,
	-- rebuilding GameTooltip from scratch with none of this addon's
	-- lines a moment after they were drawn. The fix is GameTooltip's own
	-- OnTooltipCleared/OnShow, which fire for that rebuild even though
	-- the button's own OnEnter script does not.
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
			start({ level = 20, globals = { GetItemInfo = function() return nil end } })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerBisSlotButtons()
			Tooltip.registerEmptySlotRefresh()
			mock.firstCall(button, "HookScript")[2](button) -- the first, ordinary hover
			local afterHover = mock.countCalls(_G.GameTooltip, "AddDoubleLine")

			-- The rebuild this addon otherwise never sees: Blizzard clears
			-- the tooltip and shows it again without the button's OnEnter
			-- script ever firing.
			_G.GameTooltip.GetOwner = function() return button end
			hookHandler(_G.GameTooltip, "OnTooltipCleared")(_G.GameTooltip)
			hookHandler(_G.GameTooltip, "OnShow")(_G.GameTooltip)

			assert.are.equal(afterHover + 1, mock.countCalls(_G.GameTooltip, "AddDoubleLine"))
		end)

		it("never appends the section twice for one tooltip cycle", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil end } })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerBisSlotButtons()
			Tooltip.registerEmptySlotRefresh()
			mock.firstCall(button, "HookScript")[2](button)
			local afterHover = mock.countCalls(_G.GameTooltip, "AddDoubleLine")

			-- OnShow firing again with no OnTooltipCleared in between (a
			-- mere re-anchor, not a rebuild) must not draw a second copy.
			_G.GameTooltip.GetOwner = function() return button end
			hookHandler(_G.GameTooltip, "OnShow")(_G.GameTooltip)

			assert.are.equal(afterHover, mock.countCalls(_G.GameTooltip, "AddDoubleLine"))
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
			assert.are.equal(0, mock.countCalls(_G.GameTooltip, "AddDoubleLine"))
		end)

		it("does nothing for a tooltip owned by something other than a slot button", function()
			start({ level = 20 })
			Tooltip.data = DATA
			Tooltip.registerEmptySlotRefresh()
			_G.GameTooltip.GetOwner = function() return _G.UIParent end
			hookHandler(_G.GameTooltip, "OnShow")(_G.GameTooltip)
			assert.are.equal(0, mock.countCalls(_G.GameTooltip, "AddDoubleLine"))
		end)

		it("disables itself after one failure rather than erroring again", function()
			start({ level = 20, globals = {
				GetItemInfo = function() return nil end,
				InCombatLockdown = function() error("boom") end,
			} })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerEmptySlotRefresh()
			_G.GameTooltip.GetOwner = function() return button end
			hookHandler(_G.GameTooltip, "OnShow")(_G.GameTooltip)
			assert.is_true(Tooltip.slotDisabled)
		end)
	end)

	-- The premium part (docs/tenets.md, lane addon-premium): the client's
	-- own item tooltip for the recommended item, beside GameTooltip, and
	-- redrawing the hover once the client's data for it arrives.
	describe("the compare tooltip and its refresh", function()
		local function headButton()
			return _G.CreateFrame("Button", "CharacterHeadSlot")
		end

		it("shows the compare tooltip beside GameTooltip for an empty slot's recommendation", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil, "item:111:link" end } })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerBisSlotButtons()
			mock.firstCall(button, "HookScript")[2](button)
			local compare = Theme.compareTooltipFrame
			assert.is_not_nil(compare)
			assert.are.equal(1, mock.countCalls(compare, "SetHyperlink"))
			assert.are.equal("item:111:link", mock.firstCall(compare, "SetHyperlink")[1])
			assert.are.equal(111, Tooltip.activeHover.itemId)
		end)

		it("shows the compare tooltip for an equipped item's own BiS recommendation", function()
			start({
				level = 20,
				equipped = { [1] = "item:111:link" },
				globals = { GetItemInfo = function() return nil, "item:111:link" end },
			})
			Tooltip.data = DATA
			local tooltip = recordingTooltip()
			Tooltip.onTooltip(tooltip, "item:111:link")
			assert.is_not_nil(Theme.compareTooltipFrame)
			assert.are.equal(1, mock.countCalls(Theme.compareTooltipFrame, "SetHyperlink"))
			assert.are.equal(111, Tooltip.activeHover.itemId)
		end)

		it("redraws the empty-slot hover once the client's data for that id arrives", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil, "item:111:link" end } })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerBisSlotButtons()
			mock.firstCall(button, "HookScript")[2](button)
			local setOwnerBefore = mock.countCalls(_G.GameTooltip, "SetOwner")
			Tooltip.onItemInfoReceived(111, true)
			assert.are.equal(setOwnerBefore + 1, mock.countCalls(_G.GameTooltip, "SetOwner"))
		end)

		it("does nothing for an id that is not the one currently on screen", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil, "item:111:link" end } })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerBisSlotButtons()
			mock.firstCall(button, "HookScript")[2](button)
			local before = mock.countCalls(_G.GameTooltip, "SetOwner")
			Tooltip.onItemInfoReceived(999, true)
			assert.are.equal(before, mock.countCalls(_G.GameTooltip, "SetOwner"))
		end)

		it("does nothing when the client answers the request with a failure", function()
			start({ level = 20, globals = { GetItemInfo = function() return nil, "item:111:link" end } })
			Tooltip.data = DATA
			local button = headButton()
			Tooltip.registerBisSlotButtons()
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
			assert.are.same({ "A", "B", "C", "D", "P", "Q", "R", "V", "W" }, names)
		end)
	end)
end)
