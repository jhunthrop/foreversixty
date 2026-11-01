-- addon/tests/slot_marks_spec.lua
-- The paper-doll and bag marks (spec section 4.4, task L4.1): a 12 px arrow at
-- the top right of a paper-doll slot a bag item beats and of that bag item, and
-- a 12 px target corner at the bottom right of a slot whose hover names a next
-- target. One verdict source (ruling 1): the marks read the chip's own list
-- (Chip.scoredUpgrades -> Progress.bagUpgrades) and the tooltip's own Next line,
-- so the chip, the tooltip and both marks agree on one fixture. Client-mocked:
-- the slot buttons, a container frame and hooksecurefunc are plain tables.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local DATA = {
	classes = { warrior = { tabs = {
		{ name = "Arms", talents = {} },
		{ name = "Fury", talents = {} },
		{ name = "Protection", talents = {} },
	} } },
	weights = { ["warrior-arms"] = { strength = 2, attack_power = 1 } },
	dpsPerPoint = { ["warrior-arms"] = { reference = "attack_power", dps = 0.1, healing = false, weaponDps = 0.5,
		[20] = { reference = "attack_power", dps = 0.1, weaponDps = 0.5 } } },
	bis = { ["warrior-arms"] = {
		[20] = { alliance = {
			chest = "300|D|Wailing Caverns: Mutanus the Devourer|0.33",
			legs = "400|V|Bernard Brubaker",
			finger1 = "500|Q|The Ring Quest",
			finger2 = "501|Q|The Other Ring Quest",
			trinket1 = "600|W|Sludge Beast",
			trinket2 = "600|W|Sludge Beast",
		} },
		[30] = { alliance = { chest = "310|V|Bernard Brubaker" } },
	} },
	bisNames = { [300] = "Mutant Scale Breastplate", [400] = "Chausses of Westfall", [500] = "Quest Ring",
		[501] = "Other Quest Ring", [600] = "Sludge Charm", [310] = "Later Chest" },
}

local function link(id, name)
	return "|Hitem:" .. id .. "::::::::21:::::|h[" .. name .. "]|h"
end

local WORN_CHEST = link(200, "Worn Chest")
local BAG_CHEST = link(201, "Bag Chest")
local PLATE_CHEST = link(202, "Plate Chest")
local WORN_MAIN = link(210, "Short Sword")
local WORN_OFF = link(211, "Buckler")
local BAG_AXE = link(212, "Great Axe")
local PICK_CHEST = link(300, "Mutant Scale Breastplate")
local PICK_LEGS = link(400, "Chausses of Westfall")
local WORN_RING = link(500, "Quest Ring")
local THIN_LEGS = link(410, "Thin Leggings")

--- Every item the client "knows": its id, equip location, class and subclass
--- ids, and stats, keyed by link (and by id for the picks the tooltip names).
local ITEMS = {
	[WORN_CHEST] = { id = 200, slot = "INVTYPE_CHEST", class = 4, sub = 3, name = "Worn Chest",
		stats = { ITEM_MOD_STRENGTH_SHORT = 5 } },
	[BAG_CHEST] = { id = 201, slot = "INVTYPE_CHEST", class = 4, sub = 3, name = "Bag Chest",
		stats = { ITEM_MOD_STRENGTH_SHORT = 15 } },
	[PLATE_CHEST] = { id = 202, slot = "INVTYPE_CHEST", class = 4, sub = 4, name = "Plate Chest",
		stats = { ITEM_MOD_STRENGTH_SHORT = 40 } },
	[WORN_MAIN] = { id = 210, slot = "INVTYPE_WEAPON", class = 2, sub = 7, name = "Short Sword",
		stats = { ITEM_MOD_STRENGTH_SHORT = 5 } },
	[WORN_OFF] = { id = 211, slot = "INVTYPE_SHIELD", class = 4, sub = 6, name = "Buckler",
		stats = { ITEM_MOD_STRENGTH_SHORT = 3 } },
	[BAG_AXE] = { id = 212, slot = "INVTYPE_2HWEAPON", class = 2, sub = 1, name = "Great Axe",
		stats = { ITEM_MOD_STRENGTH_SHORT = 30 } },
	[PICK_CHEST] = { id = 300, slot = "INVTYPE_CHEST", class = 4, sub = 3, name = "Mutant Scale Breastplate",
		stats = { ITEM_MOD_STRENGTH_SHORT = 10 } },
	[THIN_LEGS] = { id = 410, slot = "INVTYPE_LEGS", class = 4, sub = 3, name = "Thin Leggings",
		stats = { ITEM_MOD_STRENGTH_SHORT = 4 } },
	[WORN_RING] = { id = 500, slot = "INVTYPE_FINGER", class = 4, sub = 0, name = "Quest Ring",
		stats = { ITEM_MOD_STAMINA_SHORT = 2 } },
	[PICK_LEGS] = { id = 400, slot = "INVTYPE_LEGS", class = 4, sub = 3, name = "Chausses of Westfall",
		stats = { ITEM_MOD_STRENGTH_SHORT = 4 } },
}

local BY_ID = {}
for itemLink, info in pairs(ITEMS) do
	info.link = itemLink
	BY_ID[info.id] = info
end

local function lookup(item)
	if type(item) == "number" then
		return BY_ID[item]
	end
	return ITEMS[item]
end

--- The client's own hooksecurefunc: runs `hook` after the original, never
--- instead of it, on a table's method or a global function.
local function hooksecurefunc(target, method, hook)
	if type(target) == "string" then
		target, method, hook = _G, target, method
	end
	local original = target[method]
	target[method] = function(...)
		local results = { original(...) }
		hook(...)
		return table.unpack(results)
	end
end

--- The client's English item type subtypes, by class and subclass id.
local SUBTYPES = {
	[2] = { [1] = "Two-Handed Axes", [7] = "One-Handed Swords" },
	[4] = { [0] = "Miscellaneous", [3] = "Mail", [4] = "Plate", [6] = "Shields" },
}

--- Item ids the client has not sent yet (GetItemInfo answers nil for them).
local HIDDEN = {}

local SLOT_IDS = { head = 1, chest = 5, legs = 7, finger1 = 11, finger2 = 12, trinket1 = 13, trinket2 = 14,
	main_hand = 16, off_hand = 17 }

describe("SlotMarks", function()
	local SlotMarks, ItemTooltip, Chip, Theme, Prefs, Follow
	local state

	--- `worn` is { [slot] = link }, `bag` a list of links in bag 0, `extra` more client globals.
	local function start(worn, bag, extra, level)
		local equipped = {}
		for slot, itemLink in pairs(worn or {}) do
			equipped[SLOT_IDS[slot]] = itemLink
		end
		local itemStats = {}
		for itemLink, info in pairs(ITEMS) do
			itemStats[itemLink] = info.stats
		end
		local globals = {
			UnitFactionGroup = function() return "Alliance" end,
			GetItemInfoInstant = function(item)
				local info = lookup(item)
				if info == nil then
					return nil
				end
				-- The client's shape, type and subtype as strings (fix round 1, I1).
				return info.id, info.class == 2 and "Weapon" or "Armor", SUBTYPES[info.class][info.sub] or "Miscellaneous",
					info.slot, nil, info.class, info.sub
			end,
			GetItemInfo = function(item)
				local info = lookup(item)
				if info == nil or HIDDEN[info.id] then
					return nil
				end
				return info.name, info.link, 2
			end,
			hooksecurefunc = hooksecurefunc,
		}
		for name, value in pairs(extra or {}) do
			globals[name] = value
		end
		state = mock.install({
			class = { name = "Warrior", token = "WARRIOR" },
			level = level or 21,
			equipped = equipped,
			bags = { [0] = bag or {} },
			itemStats = itemStats,
			globals = globals,
		})
		Theme = require("Theme")
		Theme.reset()
		helper.load("Compat")
		helper.load("Export")
		helper.load("Gear")
		helper.load("Talents")
		Prefs = helper.load("Prefs")
		Follow = helper.load("Follow")
		helper.load("Progress")
		ItemTooltip = helper.load("ItemTooltip")
		Chip = helper.load("Chip")
		SlotMarks = helper.load("SlotMarks")
		Follow.build = { classSlug = "warrior", gear = {}, order = {} }
		ItemTooltip.data = DATA
		SlotMarks.data = DATA
		return SlotMarks
	end

	--- A paper-doll slot button, as the client names it (CharacterChestSlot).
	local function slotButton(name)
		local button = _G.CreateFrame("Button", name)
		button.IsVisible = function() return true end
		return button
	end

	--- The client's first container frame with `count` item buttons, holding bag `bagId`.
	local function containerFrame(bagId, count)
		local frame = _G.CreateFrame("Frame", "ContainerFrame1")
		frame.GetID = function() return bagId end
		frame.GetName = function() return "ContainerFrame1" end
		frame.IsVisible = function() return true end
		frame.size = count
		for index = 1, count do
			local button = _G.CreateFrame("Button", "ContainerFrame1Item" .. index, frame)
			button.GetID = function() return index end
			button.GetParent = function() return frame end
			button.IsVisible = function() return true end
		end
		return frame
	end

	local function arrowOn(button)
		local marks = SlotMarks.marksOn(button)
		return marks ~= nil and marks.arrow ~= nil and marks.arrow:IsShown() or false
	end

	local function targetOn(button)
		local marks = SlotMarks.marksOn(button)
		return marks ~= nil and marks.target ~= nil and marks.target[1]:IsShown() and marks.target[2]:IsShown() or false
	end

	after_each(function()
		mock.uninstall()
		_G.ContainerFrame_Update = nil
		for id in pairs(HIDDEN) do
			HIDDEN[id] = nil
		end
	end)

	describe("the theme", function()
		-- L4.1 fix round 7: the arrow is drawn (SlotMarks.ARROW_PARTS); the unread texture
		-- path is gone, the atlas name kept for the atlas branch.
		it("draws a 12 px mark, the arrow the addon's own parts", function()
			start()
			assert.are.equal(12, Theme.SIZES.mark)
			assert.is_nil(Theme.MEDIA.arrow)
			assert.are.equal("bags-greenarrow", Theme.MEDIA.arrowAtlas)
		end)
	end)

	-- Ruling 1: the chip names the upgrade, the tooltip calls it an Upgrade, and the
	-- slot and the bag item carry the arrow, all from one fixture.
	describe("one verdict source", function()
		it("agrees on a bag chest: the chip's line, the tooltip's verdict, the slot and the bag item", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			-- The chip: "Equip Bag Chest, +2.0 DPS" (10 Str x 2 = 20 AP x 0.1).
			local input = Chip.inputs(DATA)
			local model = Chip.model(input)
			assert.are.equal("upgrade", model.chosen.kind)
			assert.are.equal(string.format(L.chipEquip, "Bag Chest", string.format(L.chipGainDps, 2)), model.chosen.text)
			-- The tooltip on the bag item: "Upgrade +2.0 DPS over Worn Chest".
			local lines = ItemTooltip.section(ItemTooltip.contextFor(DATA, BAG_CHEST))
			assert.are.equal(string.format(L.tooltipVerdictDps, 2, "Worn Chest"), lines[1].text)
			-- The marks: the chest slot and the bag item that beats it, nothing else.
			local marks = SlotMarks.compute(DATA)
			assert.are.same({ chest = true }, marks.arrows)
			assert.are.same({ [BAG_CHEST] = true }, marks.links)
			local chest, legs = slotButton("CharacterChestSlot"), slotButton("CharacterLegsSlot")
			local frame = containerFrame(0, 1)
			_G.ContainerFrame_Update = function() end
			SlotMarks.register()
			SlotMarks.drawPaperDoll()
			_G.ContainerFrame_Update(frame)
			assert.is_true(arrowOn(chest))
			assert.is_false(arrowOn(legs))
			assert.is_true(arrowOn(_G.ContainerFrame1Item1))
		end)

		it("marks the main hand alone for a two-hander, net of both hands as the chip and tooltip say", function()
			start({ main_hand = WORN_MAIN, off_hand = WORN_OFF }, { BAG_AXE })
			-- 30 Str against 5 + 3: 22 x 2 = 44 AP x 0.1 = 4.4 DPS.
			local model = Chip.model(Chip.inputs(DATA))
			assert.are.equal(string.format(L.chipEquip, "Great Axe", string.format(L.chipGainDps, 4.4)), model.chosen.text)
			local lines = ItemTooltip.section(ItemTooltip.contextFor(DATA, BAG_AXE))
			assert.are.equal(string.format(L.tooltipVerdictDps, 4.4, "Short Sword"), lines[1].text)
			local marks = SlotMarks.compute(DATA)
			assert.are.same({ main_hand = true }, marks.arrows)
			assert.are.same({ [BAG_AXE] = true }, marks.links)
			local main, off = slotButton("CharacterMainHandSlot"), slotButton("CharacterSecondaryHandSlot")
			SlotMarks.drawPaperDoll()
			assert.is_true(arrowOn(main))
			assert.is_false(arrowOn(off))
		end)

		-- Ruling 4: Rules.canEquip gates the chip's list, so neither mark shows.
		it("never marks a plate chest at 21: no chip line, a refusal on the tooltip, no slot or bag mark", function()
			start({ chest = WORN_CHEST }, { PLATE_CHEST })
			local model = Chip.model(Chip.inputs(DATA))
			assert.are_not.equal("upgrade", model.chosen.kind)
			local lines = ItemTooltip.section(ItemTooltip.contextFor(DATA, PLATE_CHEST))
			assert.are.equal(string.format(L.tooltipUntil, 40, "plate armor"), lines[1].text)
			local marks = SlotMarks.compute(DATA)
			assert.are.same({}, marks.arrows)
			assert.are.same({}, marks.links)
			local chest = slotButton("CharacterChestSlot")
			local frame = containerFrame(0, 1)
			_G.ContainerFrame_Update = function() end
			SlotMarks.register()
			SlotMarks.drawPaperDoll()
			_G.ContainerFrame_Update(frame)
			assert.is_false(arrowOn(chest))
			assert.is_false(arrowOn(_G.ContainerFrame1Item1))
		end)

		it("draws no arrow with no build loaded, as the chip offers none", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			Follow.build = nil
			assert.are.same({}, SlotMarks.compute(DATA).arrows)
		end)
	end)

	-- Ruling 2: the target corner, where the slot's hover names a next target that
	-- is not worn and not in the bags.
	describe("the target corner", function()
		it("marks an empty slot whose band pick is known", function()
			start({ chest = WORN_CHEST }, {})
			local marks = SlotMarks.compute(DATA)
			assert.are.equal(400, marks.targets.legs)
			local legs = slotButton("CharacterLegsSlot")
			SlotMarks.drawPaperDoll()
			assert.is_true(targetOn(legs))
			local marks2 = SlotMarks.marksOn(legs)
			assert.are.equal(12, marks2.target[1]:GetWidth())
			assert.are.equal(12, marks2.target[2]:GetHeight())
		end)

		it("marks a worn slot whose hover prints a Next upgrade line", function()
			-- Mutant Scale (10 Str) over the worn chest (5): 1.0 DPS, from the floor.
			start({ chest = WORN_CHEST }, {})
			assert.are.equal(300, SlotMarks.compute(DATA).targets.chest)
		end)

		it("leaves a slot whose pick is already worn, or carried in the bags", function()
			start({ chest = PICK_CHEST }, { PICK_LEGS })
			local marks = SlotMarks.compute(DATA)
			assert.is_nil(marks.targets.chest)
			assert.is_nil(marks.targets.legs)
		end)

		it("leaves a worn slot that holds until the next rung (no Next line on its hover)", function()
			-- The bag chest (15 Str) worn beats the pick (10): "Best until level 30".
			start({ chest = BAG_CHEST }, {})
			assert.is_nil(SlotMarks.compute(DATA).targets.chest)
		end)

		it("follows the pair rule: a ring pick worn in the other ring slot is not offered", function()
			start({ finger2 = WORN_RING }, {})
			local marks = SlotMarks.compute(DATA)
			assert.is_nil(marks.targets.finger1)
		end)

		it("offers a pick both trinket slots share once, on the first", function()
			start({}, {})
			local marks = SlotMarks.compute(DATA)
			assert.are.equal(600, marks.targets.trinket1)
			assert.is_nil(marks.targets.trinket2)
		end)
	end)

	-- Hotfix 2.0.2: below the first band (the data's bands start at 20) nothing is an upgrade yet, so
	-- no arrow, no bag arrow and no gold corner: the hover says where picks start instead.
	describe("below the first band (level 18)", function()
		it("draws no arrow, no bag mark and no target corner, and the chip has no gear line", function()
			start({ chest = WORN_CHEST }, {}, nil, 18)
			local marks = SlotMarks.compute(DATA)
			assert.are.same({}, marks.arrows)
			assert.are.same({}, marks.links)
			assert.are.same({}, marks.targets)
			local input = Chip.inputs(DATA)
			assert.are.equal(20, input.picksStart)
			local model = Chip.model(input)
			assert.are_not.equal("upgrade", model.chosen.kind)
			assert.are_not.equal("here", model.chosen.kind)
			local chest, legs = slotButton("CharacterChestSlot"), slotButton("CharacterLegsSlot")
			SlotMarks.drawPaperDoll()
			assert.is_false(arrowOn(chest))
			assert.is_false(targetOn(chest))
			assert.is_false(targetOn(legs))
		end)

		it("draws no gold corner for the first band's picks, though a bag item beating the worn one still arrows", function()
			-- The arrow is the bag-versus-worn verdict (band-independent, real leveling gear); the corner
			-- is the band's pick, which does not exist yet.
			start({ chest = WORN_CHEST }, { BAG_CHEST }, nil, 18)
			local marks = SlotMarks.compute(DATA)
			assert.are.same({ chest = true }, marks.arrows)
			assert.are.same({}, marks.targets)
		end)

		it("keeps the empty slot's hover a Picks start line, never a Next line", function()
			start({ chest = WORN_CHEST }, {}, nil, 18)
			local lines = ItemTooltip.emptySlotSection(ItemTooltip.baseContext(DATA), "legs")
			assert.are.equal("firstBand", lines[1].kind)
			assert.are.equal("at", lines[2].kind)
		end)

		it("serves the same marks from level 20", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST }, nil, 20)
			assert.are.same({ chest = true }, SlotMarks.compute(DATA).arrows)
			assert.is_nil(Chip.inputs(DATA).picksStart)
		end)
	end)

	describe("the frames", function()
		-- Fix round 2: on the live client a texture made on a slot button and an OnShow
		-- hooked on the pane tainted the pane's own Show and Hide. The marks live on addon
		-- overlays parented to the buttons; nothing is hooked or written on the pane.
		it("hooks nothing on the character pane and makes nothing before a mark is needed", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			local chest = slotButton("CharacterChestSlot")
			local paperDoll = _G.CreateFrame("Frame", "PaperDollFrame")
			local characterFrame = _G.CreateFrame("Frame", "CharacterFrame")
			SlotMarks.register()
			assert.are.equal(0, mock.countCalls(paperDoll, "HookScript"))
			assert.are.equal(0, mock.countCalls(characterFrame, "HookScript"))
			assert.is_nil(SlotMarks.marksOn(chest))
			assert.are.equal(0, #chest.children)
		end)

		it("draws on an addon overlay parented to the button, covering it, with no script of its own", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			local chest = slotButton("CharacterChestSlot")
			SlotMarks.drawPaperDoll()
			local marks = SlotMarks.marksOn(chest)
			local overlay = marks.overlay
			assert.are.equal(chest, overlay.parent)
			assert.is_nil(overlay.name)
			assert.are.same({ chest }, { mock.firstCall(overlay, "SetAllPoints")[1] })
			assert.are.equal(0, mock.countCalls(overlay, "SetScript"))
			assert.are.equal(0, mock.countCalls(overlay, "HookScript"))
			assert.are.equal(0, #chest.regions)
			assert.are.equal(overlay, marks.arrow.parent)
			assert.are.same({ "TOPRIGHT", overlay, "TOPRIGHT", 0, 0 }, marks.arrow.points[1])
			assert.are.equal(12, marks.arrow:GetWidth())
			assert.are.equal(12, marks.arrow:GetHeight())
			-- The corner, on a slot without the arrow (fix round 5: an arrow slot has none).
			local legs = slotButton("CharacterLegsSlot")
			SlotMarks.drawPaperDoll()
			local legMarks = SlotMarks.marksOn(legs)
			assert.are.equal(legMarks.overlay, legMarks.target[1].parent)
			assert.are.equal("BOTTOMRIGHT", legMarks.target[1].points[1][1])
			assert.are.equal(legMarks.overlay, legMarks.target[2].points[1][2])
		end)

		it("keeps the paper doll's marks current on the events, the pane shown or not", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			local chest = slotButton("CharacterChestSlot")
			chest.IsVisible = function() return false end
			SlotMarks.refresh(DATA)
			assert.is_true(arrowOn(chest))
		end)

		it("never writes into a client frame: no texture, script, hook or field, through every path", function()
			local mixin = { UpdateNewItem = function() end }
			start({ chest = WORN_CHEST }, { BAG_CHEST }, {
				ContainerFrameItemButtonMixin = mixin, ToggleBackpack = function() end,
			})
			for _, slotName in pairs(ItemTooltip.SLOT_BUTTONS) do
				slotButton(slotName)
			end
			_G.CreateFrame("Frame", "PaperDollFrame")
			_G.CreateFrame("Frame", "CharacterFrame")
			local frame = containerFrame(0, 1)
			frame.Items = { _G.ContainerFrame1Item1 }
			frame.UpdateItems = function() end
			_G.ContainerFrame_Update = function() end
			-- Built, drawn, refreshed, toggled, through every bag path.
			SlotMarks.register()
			SlotMarks.drawPaperDoll()
			frame:UpdateItems()
			_G.ContainerFrame_Update(frame)
			_G.ContainerFrame1Item1.UpdateNewItem = mixin.UpdateNewItem
			_G.ContainerFrame1Item1:UpdateNewItem(2)
			_G.ToggleBackpack()
			SlotMarks.refresh(DATA, "BAG_UPDATE_DELAYED")
			SlotMarks.refresh(DATA, "PLAYER_EQUIPMENT_CHANGED")
			Prefs.setFlag("tooltip", false)
			SlotMarks.refresh(DATA)
			Prefs.setFlag("tooltip", true)
			SlotMarks.refresh(DATA)
			SlotMarks.onItemInfo(400)
			SlotMarks.onCombatEnd()
			assert.is_true(arrowOn(_G.CharacterChestSlot))
			assert.is_true(arrowOn(_G.ContainerFrame1Item1))
			assert.are.same({}, state.clientWrites)
		end)

		it("the mock catches a write into a client frame (the spec above can fail)", function()
			start()
			local chest = slotButton("CharacterChestSlot")
			local probe = assert(load("local button = ...; button:CreateTexture(nil, 'OVERLAY'); button.fsMark = true",
				"@./ForeverSixty/surfaces/Probe.lua"))
			probe(chest)
			assert.are.equal(2, #state.clientWrites)
			assert.are.equal("CreateTexture", state.clientWrites[1].kind)
			assert.are.equal("field", state.clientWrites[2].kind)
		end)

		it("waits for the end of combat to make a new overlay, then draws it", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST }, { InCombatLockdown = function() return true end })
			local chest = slotButton("CharacterChestSlot")
			SlotMarks.drawPaperDoll()
			assert.is_nil(SlotMarks.marksOn(chest))
			_G.InCombatLockdown = function() return false end
			SlotMarks.onCombatEnd()
			assert.is_true(arrowOn(chest))
		end)

		-- Ruling 6: the marks follow "Gear tips on item tooltips".
		it("hides both marks with the tooltip switch off", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			local chest = slotButton("CharacterChestSlot")
			local frame = containerFrame(0, 1)
			_G.ContainerFrame_Update = function() end
			SlotMarks.register()
			SlotMarks.drawPaperDoll()
			_G.ContainerFrame_Update(frame)
			local legs = slotButton("CharacterLegsSlot")
			SlotMarks.drawPaperDoll()
			assert.is_true(arrowOn(chest))
			assert.is_true(targetOn(legs))
			Prefs.setFlag("tooltip", false)
			SlotMarks.refresh(DATA)
			assert.is_false(arrowOn(chest))
			assert.is_false(targetOn(legs))
			assert.is_false(arrowOn(_G.ContainerFrame1Item1))
		end)

		it("takes the arrows away once the upgrade is worn (refresh on the equipment event)", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			local chest = slotButton("CharacterChestSlot")
			local frame = containerFrame(0, 1)
			_G.ContainerFrame_Update = function() end
			SlotMarks.register()
			SlotMarks.drawPaperDoll()
			_G.ContainerFrame_Update(frame)
			assert.is_true(arrowOn(chest))
			-- Equip it: the bag now holds the old chest, which beats nothing.
			state.equipped[5] = BAG_CHEST
			state.bags[0] = { WORN_CHEST }
			SlotMarks.refresh(DATA)
			assert.is_false(arrowOn(chest))
			assert.is_false(arrowOn(_G.ContainerFrame1Item1))
		end)

		it("reads the client once per refresh: the marks are cached by item link until invalidated", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			local calls = 0
			local scored = Chip.scoredUpgrades
			Chip.scoredUpgrades = function(...)
				calls = calls + 1
				return scored(...)
			end
			local frame = containerFrame(0, 1)
			_G.ContainerFrame_Update = function() end
			SlotMarks.register()
			_G.ContainerFrame_Update(frame)
			_G.ContainerFrame_Update(frame)
			SlotMarks.drawPaperDoll()
			assert.are.equal(1, calls)
			SlotMarks.refresh(DATA)
			assert.are.equal(2, calls)
		end)

		it("reads no slot's section for the bags alone: the targets wait for the paper doll", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			local sections = 0
			local emptySection = ItemTooltip.emptySlotSection
			ItemTooltip.emptySlotSection = function(...)
				sections = sections + 1
				return emptySection(...)
			end
			local frame = containerFrame(0, 1)
			_G.ContainerFrame_Update = function() end
			SlotMarks.register()
			_G.ContainerFrame_Update(frame)
			assert.is_true(arrowOn(_G.ContainerFrame1Item1))
			assert.are.equal(0, sections)
			SlotMarks.drawPaperDoll()
			assert.is_true(sections > 0)
		end)

		it("never polls: no OnUpdate script anywhere in the file", function()
			local file = assert(io.open("ForeverSixty/surfaces/SlotMarks.lua", "r"))
			local source = file:read("a")
			file:close()
			assert.is_nil(source:find("OnUpdate", 1, true))
		end)

		it("hooks nothing and fails nothing on a client with no hook function", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST }, { hooksecurefunc = false })
			assert.has_no.errors(function()
				SlotMarks.register()
				SlotMarks.refresh(DATA)
			end)
			assert.are.same({}, SlotMarks.bagHooks)
		end)
	end)

	describe("the bag hook", function()
		it("reaches a later button built from the client's item button mixin (a bag addon's)", function()
			local mixin = { UpdateNewItem = function() end }
			start({ chest = WORN_CHEST }, { BAG_CHEST }, { ContainerFrameItemButtonMixin = mixin })
			SlotMarks.register()
			-- A bag addon's button, made after the hook from the mixin (Mixin copies the methods).
			local holder = _G.CreateFrame("Frame", "BaganatorBag")
			local button = _G.CreateFrame("Button", "BaganatorItem1", holder)
			button.UpdateNewItem = mixin.UpdateNewItem
			button.GetBagID = function() return 0 end
			button.GetID = function() return 1 end
			button.GetParent = function() return holder end
			button.IsVisible = function() return true end
			button:UpdateNewItem(2)
			assert.is_true(arrowOn(button))
		end)

		-- Fix round 1, M4: "not checked yet" until the bags were opened, "off" only after
		-- they were opened and the hook reached none of the bag addon's buttons.
		it("says not checked yet until the bags open, then names the bag addon the hook did not reach", function()
			start({}, {}, {
				C_AddOns = { IsAddOnLoaded = function(name) return name == "Baganator" end },
				ToggleBackpack = function() end,
			})
			SlotMarks.register()
			assert.are.equal(L.diagBagAddonUnchecked, SlotMarks.diagLine())
			_G.ToggleBackpack()
			assert.are.equal(string.format(L.diagBagAddon, "Baganator"), SlotMarks.diagLine())
		end)

		it("counts a client bag frame's own update as the bags opened", function()
			start({}, {}, { C_AddOns = { IsAddOnLoaded = function(name) return name == "Baganator" end } })
			local frame = containerFrame(0, 1)
			frame.Items = { _G.ContainerFrame1Item1 }
			frame.UpdateItems = function() end
			SlotMarks.register()
			frame:UpdateItems()
			assert.are.equal(string.format(L.diagBagAddon, "Baganator"), SlotMarks.diagLine())
		end)

		it("says nothing in /fs diag once the hook has reached the bag addon's buttons", function()
			local mixin = { UpdateNewItem = function() end }
			start({ chest = WORN_CHEST }, { BAG_CHEST }, {
				ContainerFrameItemButtonMixin = mixin,
				C_AddOns = { IsAddOnLoaded = function(name) return name == "Baganator" end },
			})
			SlotMarks.register()
			local holder = _G.CreateFrame("Frame", "BaganatorBag")
			local button = _G.CreateFrame("Button", "BaganatorItem1", holder)
			button.UpdateNewItem = mixin.UpdateNewItem
			button.GetBagID = function() return 0 end
			button.GetID = function() return 1 end
			button.GetParent = function() return holder end
			button:UpdateNewItem(2)
			assert.is_nil(SlotMarks.diagLine())
		end)

		it("says nothing in /fs diag with no bag addon", function()
			start({}, {}, { C_AddOns = { IsAddOnLoaded = function() return false end } })
			SlotMarks.register()
			assert.is_nil(SlotMarks.diagLine())
		end)
	end)

	-- Fix round 1 (the live hunter, the review). I1: GetItemInfoInstant's second return is
	-- the item type, a string; feeding the whole multi-return to tonumber made it the base.
	describe("fix round 1", function()
		it("reads a bag item's id from the client's real multi-return (I1)", function()
			start({}, {}, { GetItemInfoInstant = function()
				return 202, "Armor", "Plate", "INVTYPE_CHEST", 132000, 4, 4
			end })
			local marks
			assert.has_no.errors(function()
				marks = SlotMarks.marks({}, { chest = 202, legs = 400 }, { PLATE_CHEST })
			end)
			assert.is_nil(marks.targets.chest)
			assert.are.equal(400, marks.targets.legs)
		end)

		it("survives a full bag scan on the client's shape: no mark turns itself off (I1)", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST, PICK_LEGS })
			local chest = slotButton("CharacterChestSlot")
			SlotMarks.drawPaperDoll()
			assert.is_nil(SlotMarks.disabled)
			assert.is_true(arrowOn(chest))
		end)

		-- M1: a target whose name the client has not sent is not cached; its info event
		-- redraws the marks, and an event for an id nobody waits on does no work.
		it("draws the corner once the client sends the pick it was waiting on (M1)", function()
			start({ chest = WORN_CHEST }, {})
			local unnamed = setmetatable({ bisNames = {} }, { __index = DATA })
			SlotMarks.data, ItemTooltip.data = unnamed, unnamed
			HIDDEN[400] = true
			local legs = slotButton("CharacterLegsSlot")
			SlotMarks.drawPaperDoll()
			assert.is_false(targetOn(legs))
			assert.is_true(SlotMarks.isWaiting(400))
			HIDDEN[400] = nil
			SlotMarks.onItemInfo(400)
			assert.is_true(targetOn(legs))
			assert.is_false(SlotMarks.isWaiting(400))
		end)

		it("does no work for an item id the marks are not waiting on (M1)", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			slotButton("CharacterChestSlot")
			SlotMarks.drawPaperDoll()
			local calls = 0
			local scored = Chip.scoredUpgrades
			Chip.scoredUpgrades = function(...)
				calls = calls + 1
				return scored(...)
			end
			SlotMarks.onItemInfo(999)
			assert.are.equal(0, calls)
		end)

		-- M2: the 1.60 client has no ContainerFrame_Update; its frames' own UpdateItems over
		-- frame.Items, the combined bag, and the bag read from the button or else the frame.
		it("marks through a frame's own UpdateItems over its Items, the bag from each button (M2)", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			local frame = _G.CreateFrame("Frame", "ContainerFrame2")
			frame.GetName = function() return "ContainerFrame2" end
			frame.GetID = function() return 3 end
			frame.UpdateItems = function() end
			local button = _G.CreateFrame("Button", nil, frame)
			button.GetBagID = function() return 0 end
			button.GetID = function() return 1 end
			button.GetParent = function() return frame end
			frame.Items = { button }
			SlotMarks.register()
			-- No ContainerFrame_Update on 1.60: the frames' own UpdateItems is the path.
			assert.are.same({ "UpdateItems" }, SlotMarks.bagHooks)
			frame:UpdateItems()
			assert.is_true(arrowOn(button))
		end)

		it("marks the combined bag's items through its UpdateItems (M2)", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			local frame = _G.CreateFrame("Frame", "ContainerFrameCombinedBags")
			frame.GetName = function() return "ContainerFrameCombinedBags" end
			frame.UpdateItems = function() end
			local first = _G.CreateFrame("Button", nil, frame)
			first.GetBagID = function() return 0 end
			first.GetID = function() return 1 end
			frame.Items = { first }
			SlotMarks.register()
			frame:UpdateItems()
			assert.is_true(arrowOn(first))
		end)

		it("reads the bag from the frame when a button has no GetBagID (M2)", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			local frame = _G.CreateFrame("Frame", "ContainerFrame3")
			frame.GetName = function() return "ContainerFrame3" end
			frame.GetID = function() return 0 end
			frame.UpdateItems = function() end
			local button = _G.CreateFrame("Button", nil, frame)
			button.GetID = function() return 1 end
			frame.Items = { button }
			SlotMarks.register()
			frame:UpdateItems()
			assert.is_true(arrowOn(button))
		end)

		it("marks a client button reached through the mixin's UpdateNewItem, not counted as a bag addon's (M2)", function()
			local mixin = { UpdateNewItem = function() end }
			start({ chest = WORN_CHEST }, { BAG_CHEST }, {
				ContainerFrameItemButtonMixin = mixin,
				C_AddOns = { IsAddOnLoaded = function(name) return name == "Baganator" end },
			})
			local frame = containerFrame(0, 1)
			local button = _G.ContainerFrame1Item1
			SlotMarks.register()
			assert.are.equal("ContainerFrameItemButtonMixin.UpdateNewItem", SlotMarks.bagHooks[1])
			button.UpdateNewItem = mixin.UpdateNewItem
			button:UpdateNewItem(2)
			assert.is_true(arrowOn(button))
			assert.is_not_nil(frame)
			assert.are.equal(string.format(L.diagBagAddon, "Baganator"), SlotMarks.diagLine())
		end)

		-- M3: a carried pick the character cannot equip yet is not "carried": the corner shows.
		it("keeps the corner for a carried pick the character cannot wear yet (M3)", function()
			local plateData = setmetatable({ bis = { ["warrior-arms"] = { [20] = { alliance = {
				chest = "202|V|Plate Vendor" } } } } }, { __index = DATA })
			start({ chest = WORN_CHEST }, { PLATE_CHEST })
			SlotMarks.data, ItemTooltip.data = plateData, plateData
			local marks = SlotMarks.compute(plateData)
			assert.are.equal(202, marks.targets.chest)
			assert.is_nil(marks.arrows.chest)
			-- At 40 the plate chest is wearable: it is carried, so no corner, and it is an upgrade.
			start({ chest = WORN_CHEST }, { PLATE_CHEST }, nil, 40)
			SlotMarks.data, ItemTooltip.data = plateData, plateData
			local at40 = SlotMarks.compute(plateData)
			assert.is_nil(at40.targets.chest)
			assert.is_true(at40.arrows.chest)
		end)

		-- M4: one failure turns the marks off until /reload, says so once, never spams.
		it("turns itself off until /reload after one failure, and says so once (M4)", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			slotButton("CharacterChestSlot")
			Chip.scoredUpgrades = function()
				error("boom")
			end
			SlotMarks.drawPaperDoll()
			SlotMarks.refresh(DATA)
			SlotMarks.drawPaperDoll()
			assert.is_true(SlotMarks.disabled)
			local notes = 0
			for _, note in ipairs(Theme.diagnostics()) do
				if note:find(L.diagSlotMarksFailed:match("^(.-)%%s"), 1, true) then
					notes = notes + 1
				end
			end
			assert.are.equal(1, notes)
			assert.is_truthy(L.diagSlotMarksFailed:find("until /reload", 1, true))
		end)

		-- N1: an empty off hand under a worn two-hander says so on its hover and has no corner.
		it("draws no corner on an empty off hand under a worn two-hander (N1)", function()
			local offData = setmetatable({ bis = { ["warrior-arms"] = { [20] = { alliance = {
				off_hand = "211|V|Shield Vendor" } } } } }, { __index = DATA })
			start({ main_hand = BAG_AXE }, {})
			SlotMarks.data, ItemTooltip.data = offData, offData
			assert.is_nil(SlotMarks.compute(offData).targets.off_hand)
			start({ main_hand = WORN_MAIN }, {})
			SlotMarks.data, ItemTooltip.data = offData, offData
			assert.are.equal(211, SlotMarks.compute(offData).targets.off_hand)
		end)
	end)

	-- Fix round 3 (hunter-9/10).
	describe("fix round 3", function()
		-- 1: the corner means a next target for THIS band, never a later band's At line.
		it("draws no corner on a slot whose hover says Best until with an At line for a later band", function()
			start({ chest = BAG_CHEST }, {})
			local lines = ItemTooltip.sectionFor(BAG_CHEST, ItemTooltip.baseContext(DATA)).lines
			assert.are.equal(string.format(L.tooltipBestUntil, 30), lines[1].text)
			assert.are.equal("at", lines[2].kind)
			local chest = slotButton("CharacterChestSlot")
			SlotMarks.drawPaperDoll()
			assert.is_false(targetOn(chest))
		end)

		it("draws the corner on a slot whose hover prints a Next upgrade line", function()
			start({ chest = WORN_CHEST }, {})
			local lines = ItemTooltip.sectionFor(WORN_CHEST, ItemTooltip.baseContext(DATA)).lines
			assert.are.equal("next", lines[1].kind)
			local chest = slotButton("CharacterChestSlot")
			SlotMarks.drawPaperDoll()
			assert.is_true(targetOn(chest))
		end)

		-- A worn item the client has not sent yet scores as nothing: its section would
		-- print a Next line it does not mean. Fix round 4 (the review): no corner at all
		-- while it is on its way, never a false one that shows and then disappears; the
		-- targets wait for it and redraw when it comes.
		it("draws no corner while a worn item is on its way, and the right one once it comes", function()
			-- The worn chest (5 Str) trails the band's pick (10 Str): once known, a corner.
			start({ chest = WORN_CHEST }, {})
			HIDDEN[200] = true
			state.itemStats[WORN_CHEST] = nil
			local chest = slotButton("CharacterChestSlot")
			SlotMarks.drawPaperDoll()
			assert.is_false(targetOn(chest))
			assert.is_true(SlotMarks.isWaiting(200))
			HIDDEN[200] = nil
			state.itemStats[WORN_CHEST] = ITEMS[WORN_CHEST].stats
			SlotMarks.onItemInfo(200)
			assert.is_true(targetOn(chest))
			assert.is_false(SlotMarks.isWaiting(200))
		end)

		it("draws no corner during the wait for a worn item that holds, nor after it comes", function()
			start({ chest = BAG_CHEST }, {})
			HIDDEN[201] = true
			state.itemStats[BAG_CHEST] = nil
			local chest = slotButton("CharacterChestSlot")
			SlotMarks.drawPaperDoll()
			assert.is_false(targetOn(chest))
			HIDDEN[201] = nil
			state.itemStats[BAG_CHEST] = ITEMS[BAG_CHEST].stats
			SlotMarks.onItemInfo(201)
			assert.is_false(targetOn(chest))
		end)

		-- A band pick the client cannot show at all (ItemTooltip.missing, as the section
		-- reads it) is never waited on, so the targets are kept across bag updates.
		it("never waits on a band pick the client cannot show", function()
			start({ chest = WORN_CHEST }, {})
			HIDDEN[400] = true
			ItemTooltip.missing[400] = true
			slotButton("CharacterLegsSlot")
			SlotMarks.drawPaperDoll()
			assert.is_false(SlotMarks.isWaiting(400))
			ItemTooltip.missing[400] = nil
		end)

		-- 2: an item that fills an EMPTY slot is an upgrade below the chip's floor.
		it("marks a bag item that fills an empty slot: the chip, the tooltip and both arrows", function()
			start({ chest = WORN_CHEST }, { THIN_LEGS })
			-- 4 Str x 2 = 8 points x 0.1 = 0.8 DPS, under the 1.0 floor.
			local model = Chip.model(Chip.inputs(DATA))
			assert.are.equal(string.format(L.chipEquip, "Thin Leggings", string.format(L.chipGainDps, 0.8)),
				model.chosen.text)
			local lines = ItemTooltip.section(ItemTooltip.contextFor(DATA, THIN_LEGS))
			assert.are.same({ text = string.format(L.tooltipFills, 0.8, L.tooltipUnitDps), color = "success" },
				{ text = lines[1].text, color = lines[1].color })
			local marks = SlotMarks.compute(DATA)
			assert.is_true(marks.arrows.legs)
			assert.is_true(marks.links[THIN_LEGS])
			local legs = slotButton("CharacterLegsSlot")
			local frame = containerFrame(0, 1)
			_G.ContainerFrame_Update = function() end
			SlotMarks.register()
			SlotMarks.drawPaperDoll()
			_G.ContainerFrame_Update(frame)
			assert.is_true(arrowOn(legs))
			assert.is_true(arrowOn(_G.ContainerFrame1Item1))
		end)

		-- 4: the overlay sits one level above its button.
		it("puts the overlay one frame level above its button", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			local chest = slotButton("CharacterChestSlot")
			chest.GetFrameLevel = function() return 5 end
			SlotMarks.drawPaperDoll()
			assert.are.equal(6, mock.lastCall(SlotMarks.marksOn(chest).overlay, "SetFrameLevel")[1])
			assert.are.same({}, state.clientWrites)
		end)
	end)

	-- Fix round 4, ruling 1 (hunter-12): the hunter wears the band's own Hands pick, so the
	-- hover reads "Best until level 30 / At level 30: <the 30 band's pick> / Planned for
	-- Beast Mastery at 20". The picks are read from the shipped Data, which the nightly
	-- regenerates (the pick moved from Gloves of the Fang to the Serpent's Gloves when the
	-- 2026-10-31 pet rework changed the hunter's weights), so only the behaviour is pinned.
	-- Through the real section path with the real Data: no Next line, so no corner; the At
	-- line's later pick never becomes a target.
	describe("fix round 4: the hunter's Best until slot, from the real data", function()
		local SHIPPED = require("Data")
		local function bandPick(band, slot)
			local entry = SHIPPED.bis["hunter-beast-mastery"][band].alliance[slot]
			local id = tonumber(entry:match("^(%d+)|"))
			return id, SHIPPED.bisNames[id]
		end
		local GLOVES_ID, GLOVES_NAME = bandPick(20, "hands")
		local BELT_PICK_ID, BELT_PICK_NAME = bandPick(20, "waist")
		local BELT_PICK_LINK = string.format("|cff1eff00|Hitem:%d::::::::24:::::|h[%s]|h|r", BELT_PICK_ID, BELT_PICK_NAME)
		local _, NEXT_HANDS_NAME = bandPick(30, "hands")
		local GLOVES = string.format("|cff1eff00|Hitem:%d::::::::24:::::|h[%s]|h|r", GLOVES_ID, GLOVES_NAME)
		local BELT = "|cff1eff00|Hitem:4999::::::::24:::::|h[Azora's Will]|h|r"
		local REAL = {
			[GLOVES] = { id = GLOVES_ID, slot = "INVTYPE_HAND", name = GLOVES_NAME,
				stats = { ITEM_MOD_STRENGTH_SHORT = 4, ITEM_MOD_AGILITY_SHORT = 6 } },
			[BELT] = { id = 4999, slot = "INVTYPE_WAIST", name = "Azora's Will", stats = { ITEM_MOD_STAMINA_SHORT = 1 } },
		}
		local REAL_BY_ID = { [GLOVES_ID] = REAL[GLOVES], [4999] = REAL[BELT], [7358] = { name = "Pilferer's Gloves" },
			[BELT_PICK_ID] = { name = BELT_PICK_NAME, stats = { ITEM_MOD_AGILITY_SHORT = 11 }, link = BELT_PICK_LINK } }

		local function hunter()
			local itemStats = { [GLOVES] = REAL[GLOVES].stats, [BELT] = REAL[BELT].stats,
				[BELT_PICK_LINK] = REAL_BY_ID[BELT_PICK_ID].stats }
			mock.install({
				class = { name = "Hunter", token = "HUNTER" }, level = 24,
				equipped = { [10] = GLOVES, [6] = BELT },
				itemStats = itemStats,
				globals = {
					UnitFactionGroup = function() return "Alliance" end,
					GetItemInfoInstant = function(item)
						local info = type(item) == "number" and REAL_BY_ID[item] or REAL[item]
						if info == nil or info.slot == nil then
							return nil
						end
						return info.id, "Armor", "Leather", info.slot, nil, 4, 2
					end,
					GetItemInfo = function(item)
						local info = type(item) == "number" and REAL_BY_ID[item] or REAL[item]
						if info == nil then
							return nil
						end
						return info.name, info.link or item, 2
					end,
				},
			})
			require("Theme").reset()
			helper.load("Compat")
			helper.load("Export")
			helper.load("Gear")
			helper.load("Talents")
			helper.load("Prefs")
			helper.load("Follow")
			helper.load("Progress")
			ItemTooltip = helper.load("ItemTooltip")
			helper.load("Chip")
			SlotMarks = helper.load("SlotMarks")
			local Data = require("Data")
			ItemTooltip.data, SlotMarks.data = Data, Data
			return Data
		end

		it("draws no corner on the Hands: the section is Best until with an At line, never a Next line", function()
			local Data = hunter()
			local base = ItemTooltip.baseContext(Data)
			assert.are.equal("hunter-beast-mastery", base.spec)
			local lines = ItemTooltip.sectionFor(GLOVES, base).lines
			local texts = {}
			for index, line in ipairs(lines) do
				texts[index] = line.text
			end
			assert.are.equal(string.format(L.tooltipBestUntil, 30), texts[1])
			assert.are.equal("at", lines[2].kind)
			assert.is_truthy(texts[2]:find(NEXT_HANDS_NAME, 1, true))
			assert.are.equal(string.format(L.tooltipPlanned, "Beast Mastery", 20), texts[3])
			assert.is_nil(SlotMarks.nextOf(lines))
			assert.is_nil(SlotMarks.compute(Data).targets.hands)
			local hands = _G.CreateFrame("Button", "CharacterHandsSlot")
			SlotMarks.drawPaperDoll()
			local marks = SlotMarks.marksOn(hands)
			assert.is_true(marks == nil or marks.target == nil or not marks.target[1]:IsShown())
		end)

		-- The capture's corner sits on the slot BELOW the Hands: the Waist (crops in the
		-- fix-4 report). A worn belt the band's pick beats prints a Next line: a corner.
		it("draws the corner on the Waist, whose hover prints Next upgrade for the band's belt", function()
			local Data = hunter()
			local lines = ItemTooltip.sectionFor(BELT, ItemTooltip.baseContext(Data)).lines
			assert.are.equal("next", lines[1].kind)
			assert.are.equal(BELT_PICK_ID, lines[1].itemId)
			assert.are.equal(BELT_PICK_ID, SlotMarks.compute(Data).targets.waist)
		end)
	end)

	-- Fix round 5 (the designer's and the player's first pass).
	describe("fix round 5", function()
		-- 1 (fix round 6, hunter-19): a real 12 px arrow, not a ~6 px dot. The client's
		-- bags-greenarrow atlas drew a dot on the live client although it knows the atlas, so
		-- the arrow is the addon's own 12 x 12 glyph: an up-pointing head of six 1 px rows
		-- (2, 4 ... 12 wide) over a 4 x 6 shaft, green, each part with a 1 px black outline.
		local PROBE = { height = 22, file = 969828, leftTexCoord = 0.12890625, bottomTexCoord = 0.44921875,
			tilesHorizontally = false, tilesVertically = false }

		it("draws the addon's own 12 px up arrow, even on a client that knows the atlas", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST }, { C_Texture = {
				GetAtlasInfo = function(name) return name == "bags-greenarrow" and PROBE or nil end,
			} })
			local chest = slotButton("CharacterChestSlot")
			SlotMarks.drawPaperDoll()
			local marks = SlotMarks.marksOn(chest)
			local glyph = marks.arrow
			assert.are.equal(marks.overlay, glyph.parent)
			assert.are.same({ "TOPRIGHT", marks.overlay, "TOPRIGHT", 0, 0 }, glyph.points[1])
			assert.are.equal(12, glyph:GetWidth())
			assert.are.equal(12, glyph:GetHeight())
			local expected = {
				{ 5, 0, 2, 1 }, { 4, 1, 4, 1 }, { 3, 2, 6, 1 }, { 2, 3, 8, 1 }, { 1, 4, 10, 1 }, { 0, 5, 12, 1 },
				{ 4, 6, 4, 6 },
			}
			assert.are.same(expected, SlotMarks.ARROW_PARTS)
			assert.are.equal(#expected, #marks.arrowParts)
			local red, greenValue, blue = Theme.rgb(Theme.HEX.success)
			local green = { red, greenValue, blue }
			for index, part in ipairs(expected) do
				local fill, outline = marks.arrowParts[index].fill, marks.arrowParts[index].outline
				assert.are.same({ "TOPLEFT", glyph, "TOPLEFT", part[1], -part[2] }, fill.points[1])
				assert.are.same({ part[3], part[4] }, { fill:GetWidth(), fill:GetHeight() })
				assert.are.same({ "TOPLEFT", glyph, "TOPLEFT", part[1] - 1, -part[2] + 1 }, outline.points[1])
				assert.are.same({ part[3] + 2, part[4] + 2 }, { outline:GetWidth(), outline:GetHeight() })
				assert.are.same(green, { table.unpack(mock.lastCall(fill, "SetColorTexture"), 1, 3) })
				assert.are.same({ 0, 0, 0, 1 }, { table.unpack(mock.lastCall(outline, "SetColorTexture"), 1, 4) })
			end
			-- The bag item carries the same glyph.
			local frame = containerFrame(0, 1)
			_G.ContainerFrame_Update = function() end
			SlotMarks.register()
			_G.ContainerFrame_Update(frame)
			assert.are.equal(#expected, #SlotMarks.marksOn(_G.ContainerFrame1Item1).arrowParts)
			assert.are.same({}, state.clientWrites)
		end)

		-- The atlas branch, kept for a client whose atlas art fills its frame (Theme.ARROW_FROM_ATLAS):
		-- SetAtlas with useAtlasSize, then the explicit 12 x 12, in that order, with the probe's table.
		it("sizes the atlas after SetAtlas(name, true) when the atlas branch is on", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST }, { C_Texture = {
				GetAtlasInfo = function(name) return name == "bags-greenarrow" and PROBE or nil end,
			} })
			Theme.ARROW_FROM_ATLAS = true
			local chest = slotButton("CharacterChestSlot")
			SlotMarks.drawPaperDoll()
			Theme.ARROW_FROM_ATLAS = false
			local texture = SlotMarks.marksOn(chest).arrowParts[1].fill
			local atlas, size
			for index, call in ipairs(texture.calls) do
				if call.method == "SetAtlas" then
					atlas = atlas or index
					assert.are.same({ "bags-greenarrow", true }, { call[1], call[2] })
				elseif call.method == "SetSize" then
					size = index
					assert.are.same({ 12, 12 }, { call[1], call[2] })
				end
			end
			assert.is_true(atlas ~= nil and size ~= nil and atlas < size)
			assert.are.equal(1, #SlotMarks.marksOn(chest).arrowParts)
		end)

		-- Fix round 7 (fix-6 reviewer): a texture without SetAtlas never takes the atlas branch.
		it("draws the addon's arrow when the atlas branch is on but the texture cannot SetAtlas", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST }, { C_Texture = {
				GetAtlasInfo = function(name) return name == "bags-greenarrow" and PROBE or nil end,
			} })
			state.missingMethods.SetAtlas = true
			Theme.ARROW_FROM_ATLAS = true
			local chest = slotButton("CharacterChestSlot")
			local ok, failure = pcall(SlotMarks.drawPaperDoll)
			Theme.ARROW_FROM_ATLAS = false
			state.missingMethods.SetAtlas = nil
			assert.is_true(ok, tostring(failure))
			local parts = SlotMarks.marksOn(chest).arrowParts
			assert.are.equal(#SlotMarks.ARROW_PARTS, #parts)
			assert.is_not_nil(parts[1].outline)
		end)

		it("draws the addon's arrow when the atlas branch is on but the client has no such atlas", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			Theme.ARROW_FROM_ATLAS = true
			local chest = slotButton("CharacterChestSlot")
			SlotMarks.drawPaperDoll()
			Theme.ARROW_FROM_ATLAS = false
			assert.are.equal(#SlotMarks.ARROW_PARTS, #SlotMarks.marksOn(chest).arrowParts)
		end)

		-- Fix round 6, 4: the legend follows the marks: off with the switch, and off once the
		-- marks have turned themselves off.
		it("hides the chip's legend once the marks have turned themselves off", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			assert.is_true(Chip.inputs(DATA).marksLegend)
			Prefs.setFlag("tooltip", false)
			assert.is_false(Chip.inputs(DATA).marksLegend)
			Prefs.setFlag("tooltip", true)
			SlotMarks.disabled = true
			assert.is_false(Chip.inputs(DATA).marksLegend)
		end)

		-- 4: the bracket clears the quality border, with a shadow behind the gold.
		it("insets the bracket 2 px from the bottom right, a 0.6 black shadow a pixel behind it", function()
			start({ chest = WORN_CHEST }, {})
			local chest = slotButton("CharacterChestSlot")
			SlotMarks.drawPaperDoll()
			local marks = SlotMarks.marksOn(chest)
			local overlay = marks.overlay
			assert.are.same({ "BOTTOMRIGHT", overlay, "BOTTOMRIGHT", -2, 2 }, marks.target[1].points[1])
			assert.are.same({ "BOTTOMRIGHT", overlay, "BOTTOMRIGHT", -2, 2 }, marks.target[2].points[1])
			assert.are.same({ "BOTTOMRIGHT", overlay, "BOTTOMRIGHT", -1, 1 }, marks.target[3].points[1])
			assert.are.same({ 0, 0, 0, 0.6 }, { table.unpack(mock.lastCall(marks.target[3], "SetColorTexture"), 1, 4) })
			assert.is_true(targetOn(chest))
		end)

		-- 3: the arrow slot's hover leads with the bag fact; that slot carries no corner.
		it("draws no corner on a slot carrying the arrow, keeping the other slots' corners", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			local chest, legs = slotButton("CharacterChestSlot"), slotButton("CharacterLegsSlot")
			SlotMarks.drawPaperDoll()
			assert.is_true(arrowOn(chest))
			assert.is_false(targetOn(chest))
			assert.is_true(targetOn(legs))
		end)

		local function addedLines()
			local lines = {}
			for _, call in ipairs(_G.GameTooltip.calls) do
				if call.method == "AddLine" then
					lines[#lines + 1] = call[1]
				end
			end
			return lines
		end

		it("leads an empty slot's hover with the bag item that fills it", function()
			start({ chest = WORN_CHEST }, { THIN_LEGS })
			local legs = slotButton("CharacterLegsSlot")
			ItemTooltip.onSlotEnter("legs", legs)
			local lines = addedLines()
			local upgrade = { name = "Thin Leggings", hex = ItemTooltip.qualityHexOf(410), gain = 0.8, unit = "dps" }
			assert.are.same({ " ", lines[2], L.tooltipInYourBags, ItemTooltip.bagItemText(upgrade),
				L.tooltipRightClickEquip }, lines)
			assert.are.equal(THIN_LEGS, mock.firstCall(Theme.compareTooltipFrame, "SetHyperlink")[1])
		end)

		it("leads a beaten worn slot's hover with the bag item that beats it, no source line", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			local calls = {}
			local tooltip = {
				AddLine = function(_, text) calls[#calls + 1] = text end,
				AddDoubleLine = function(_, left) calls[#calls + 1] = left end,
				Show = function() end,
			}
			ItemTooltip.onTooltip(tooltip, WORN_CHEST)
			local upgrade = { name = "Bag Chest", hex = ItemTooltip.qualityHexOf(201), gain = 2, unit = "dps" }
			assert.are.same({ L.tooltipInYourBags, ItemTooltip.bagItemText(upgrade), L.tooltipRightClickEquip },
				{ calls[#calls - 2], calls[#calls - 1], calls[#calls] })
			for _, text in ipairs(calls) do
				assert.is_nil(text:find("Wailing Caverns", 1, true))
			end
			-- The compare tooltip beside it shows the bag item that beats it (fix round 6, 4).
			assert.are.equal(BAG_CHEST, mock.lastCall(Theme.compareTooltipFrame, "SetHyperlink")[1])
		end)
	end)
end)
