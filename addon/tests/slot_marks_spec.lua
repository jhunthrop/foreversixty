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

local SLOT_IDS = { head = 1, chest = 5, legs = 7, finger1 = 11, finger2 = 12, trinket1 = 13, trinket2 = 14,
	main_hand = 16, off_hand = 17 }

describe("SlotMarks", function()
	local SlotMarks, ItemTooltip, Chip, Theme, Prefs, Follow
	local state

	--- `worn` is { [slot] = link }, `bag` a list of links in bag 0, `extra` more client globals.
	local function start(worn, bag, extra)
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
				return info.id, nil, nil, info.slot, nil, info.class, info.sub
			end,
			GetItemInfo = function(item)
				local info = lookup(item)
				if info == nil then
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
			level = 21,
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
	end)

	describe("the theme", function()
		it("draws a 12 px mark with the client's green arrow", function()
			start()
			assert.are.equal(12, Theme.SIZES.mark)
			assert.are.equal("Interface\\Buttons\\UI-MicroStream-Green", Theme.MEDIA.arrow)
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

	describe("the frames", function()
		it("creates no texture before the panel first opens", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			local chest = slotButton("CharacterChestSlot")
			local paperDoll = _G.CreateFrame("Frame", "PaperDollFrame")
			SlotMarks.register()
			assert.are.equal(0, #chest.regions)
			-- The panel's own OnShow draws them.
			local hook
			for _, call in ipairs(paperDoll.calls) do
				if call.method == "HookScript" and call[1] == "OnShow" then
					hook = call[2]
				end
			end
			assert.is_function(hook)
			hook(paperDoll)
			assert.is_true(arrowOn(chest))
		end)

		it("anchors the arrow at the top right and the target at the bottom right, 12 px", function()
			start({ chest = WORN_CHEST }, { BAG_CHEST })
			local chest = slotButton("CharacterChestSlot")
			SlotMarks.drawPaperDoll()
			local marks = SlotMarks.marksOn(chest)
			assert.are.same({ "TOPRIGHT", chest, "TOPRIGHT", 0, 0 }, marks.arrow.points[1])
			assert.are.equal(12, marks.arrow:GetWidth())
			assert.are.equal(12, marks.arrow:GetHeight())
			assert.are.equal("BOTTOMRIGHT", marks.target[1].points[1][1])
			assert.are.equal("BOTTOMRIGHT", marks.target[2].points[1][1])
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
			assert.is_true(arrowOn(chest))
			assert.is_true(targetOn(chest))
			Prefs.setFlag("tooltip", false)
			SlotMarks.refresh(DATA)
			assert.is_false(arrowOn(chest))
			assert.is_false(targetOn(chest))
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

		it("names the bag addon in /fs diag while the hook has not reached its buttons", function()
			start({}, {}, { C_AddOns = { IsAddOnLoaded = function(name) return name == "Baganator" end } })
			SlotMarks.register()
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
end)
