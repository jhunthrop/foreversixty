-- addon/tests/reward_marks_spec.lua
-- Quest reward, roll window and dungeon line (spec sections 4.1 and 4.4, task L4.2).
-- "Best for Arms" on the quest reward the weights prefer, "Vendor" on each when none
-- is of any use; the roll frame carries the tooltip's verdict line beside the item;
-- the chip's "here" line names the band's picks that drop in the instance entered.
-- One verdict source (ruling 1): the reward label, the roll line and the tooltip's
-- verdict read ItemTooltip.judge, the part of the section that scores the item, so
-- they agree on one fixture. Client-mocked: the reward buttons, the roll frames,
-- GetQuestItemLink, GetLootRollItemLink and GetInstanceInfo are plain tables.
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
		} },
	} },
	bisNames = { [300] = "Mutant Scale Breastplate" },
}

local function link(id, name)
	return "|Hitem:" .. id .. "::::::::21:::::|h[" .. name .. "]|h"
end

local WORN_CHEST = link(200, "Worn Chest")
local BAG_CHEST = link(201, "Bag Chest")
local PLATE_CHEST = link(202, "Plate Chest")
local RAG_CHEST = link(203, "Rag Chest")
local LIGHT_VEST = link(204, "Light Vest")
local OLD_CHEST = "|Hitem:205::::::::21:::::|h|h"
local LONG_CHEST = link(206, "Chestguard of the Unreasonably Long Defias Overseer Name")
local FILL_VEST = link(207, "Fill Vest")
local FAIR_CLOAK = link(222, "Fair Cloak")
local WORN_BACK = link(220, "Worn Cloak")
local THIN_CLOAK = link(221, "Thin Cloak")
local WORN_FEET = link(230, "Worn Boots")
local SOFT_BOOTS = link(231, "Soft Boots")
local WORN_AXE = link(240, "Worn Axe")
local BUCKLER = link(241, "Buckler")
local POTION = link(250, "Healing Potion")
local PICK_CHEST = link(300, "Mutant Scale Breastplate")

--- Every item the client "knows", keyed by link: id, equip location, class and
--- subclass ids, and stats (10 Str is 20 AP, 2.0 DPS at 0.1 per AP).
local ITEMS = {
	[WORN_CHEST] = { id = 200, slot = "INVTYPE_CHEST", class = 4, sub = 3, name = "Worn Chest",
		stats = { ITEM_MOD_STRENGTH_SHORT = 5 } },
	[BAG_CHEST] = { id = 201, slot = "INVTYPE_CHEST", class = 4, sub = 3, name = "Bag Chest",
		stats = { ITEM_MOD_STRENGTH_SHORT = 15 } },
	[PLATE_CHEST] = { id = 202, slot = "INVTYPE_CHEST", class = 4, sub = 4, name = "Plate Chest",
		stats = { ITEM_MOD_STRENGTH_SHORT = 40 } },
	[RAG_CHEST] = { id = 203, slot = "INVTYPE_CHEST", class = 4, sub = 3, name = "Rag Chest",
		stats = { ITEM_MOD_STAMINA_SHORT = 1 } },
	[LIGHT_VEST] = { id = 204, slot = "INVTYPE_CHEST", class = 4, sub = 3, name = "Light Vest",
		stats = { ITEM_MOD_STRENGTH_SHORT = 2 } },
	-- A worn chest whose link carries no name: its name comes from the client (fix round 1).
	[OLD_CHEST] = { id = 205, slot = "INVTYPE_CHEST", class = 4, sub = 3, name = "Old Chest",
		stats = { ITEM_MOD_STRENGTH_SHORT = 5 } },
	[LONG_CHEST] = { id = 206, slot = "INVTYPE_CHEST", class = 4, sub = 3,
		name = "Chestguard of the Unreasonably Long Defias Overseer Name", stats = { ITEM_MOD_STRENGTH_SHORT = 5 } },
	[FILL_VEST] = { id = 207, slot = "INVTYPE_CHEST", class = 4, sub = 3, name = "Fill Vest",
		stats = { ITEM_MOD_STRENGTH_SHORT = 4 } },
	-- 3 Str and 3 AP over the worn cloak's 3 Str: +0.3 DPS, a Sidegrade on its tooltip.
	[FAIR_CLOAK] = { id = 222, slot = "INVTYPE_CLOAK", class = 4, sub = 1, name = "Fair Cloak",
		stats = { ITEM_MOD_STRENGTH_SHORT = 3, ITEM_MOD_ATTACK_POWER_SHORT = 3 } },
	[WORN_BACK] = { id = 220, slot = "INVTYPE_CLOAK", class = 4, sub = 1, name = "Worn Cloak",
		stats = { ITEM_MOD_STRENGTH_SHORT = 3 } },
	[THIN_CLOAK] = { id = 221, slot = "INVTYPE_CLOAK", class = 4, sub = 1, name = "Thin Cloak",
		stats = { ITEM_MOD_STRENGTH_SHORT = 2 } },
	[WORN_FEET] = { id = 230, slot = "INVTYPE_FEET", class = 4, sub = 3, name = "Worn Boots",
		stats = { ITEM_MOD_STRENGTH_SHORT = 4 } },
	[SOFT_BOOTS] = { id = 231, slot = "INVTYPE_FEET", class = 4, sub = 3, name = "Soft Boots",
		stats = { ITEM_MOD_STRENGTH_SHORT = 4 } },
	[WORN_AXE] = { id = 240, slot = "INVTYPE_2HWEAPON", class = 2, sub = 1, name = "Worn Axe",
		stats = { ITEM_MOD_STRENGTH_SHORT = 20 } },
	[BUCKLER] = { id = 241, slot = "INVTYPE_SHIELD", class = 4, sub = 6, name = "Buckler",
		stats = { ITEM_MOD_STRENGTH_SHORT = 3 } },
	[POTION] = { id = 250, slot = "", class = 0, sub = 1, name = "Healing Potion", stats = {} },
	[PICK_CHEST] = { id = 300, slot = "INVTYPE_CHEST", class = 4, sub = 3, name = "Mutant Scale Breastplate",
		stats = { ITEM_MOD_STRENGTH_SHORT = 10 } },
}

local BY_ID = {}
for itemLink, info in pairs(ITEMS) do
	info.link = itemLink
	BY_ID[info.id] = info
end

--- A Theme.HEX key's red, green and blue (Theme.rgb's alpha left off).
local function rgb3(key)
	local Theme = require("Theme")
	local r, g, b = Theme.rgb(Theme.HEX[key])
	return { r, g, b }
end

local function lookup(item)
	if type(item) == "number" then
		return BY_ID[item]
	end
	return ITEMS[item]
end

local SUBTYPES = {
	[0] = { [1] = "Potion" },
	[2] = { [1] = "Two-Handed Axes" },
	[4] = { [1] = "Cloth", [3] = "Mail", [4] = "Plate", [6] = "Shields" },
}

--- Item ids the client has not sent yet (GetItemInfo answers nil for them).
local HIDDEN = {}

local SLOT_IDS = { chest = 5, feet = 8, back = 15, main_hand = 16, off_hand = 17 }

--- The client's own hooksecurefunc: runs `hook` after the original.
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

describe("RewardMarks", function()
	local RewardMarks, ItemTooltip, Prefs, Follow
	local state

	--- `worn` is { [slot] = link }; `extra` more client globals.
	local function start(worn, extra, level, mockOptions)
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
			itemStats = itemStats,
			globals = globals,
			stringWidth = mockOptions ~= nil and mockOptions.stringWidth or nil,
		})
		require("Theme").reset()
		helper.load("Compat")
		helper.load("Export")
		helper.load("Gear")
		helper.load("Talents")
		Prefs = helper.load("Prefs")
		Follow = helper.load("Follow")
		helper.load("Progress")
		ItemTooltip = helper.load("ItemTooltip")
		helper.load("Chip")
		RewardMarks = helper.load("RewardMarks")
		Follow.build = nil
		ItemTooltip.data = DATA
		RewardMarks.data = DATA
		return RewardMarks
	end

	--- The quest frame's reward panel offering `choices` (links), on the client's
	--- QuestInfo reward buttons (`prefix` "QuestInfoRewardsFrameQuestInfoItem", or the
	--- classic "QuestRewardItem"), each marked a choice with its index as its id.
	local function offer(choices, prefix)
		_G.GetNumQuestChoices = function() return #choices end
		_G.GetQuestItemLink = function(kind, index)
			return kind == "choice" and choices[index] or nil
		end
		state.extraGlobals[#state.extraGlobals + 1] = "GetNumQuestChoices"
		state.extraGlobals[#state.extraGlobals + 1] = "GetQuestItemLink"
		local buttons = {}
		for index = 1, #choices do
			buttons[index] = _G.CreateFrame("Button", (prefix or "QuestInfoRewardsFrameQuestInfoItem") .. index)
		end
		-- The client's own fields, set by the client (not the addon): written raw.
		for index = 1, #buttons do
			local raw = state.widgets[(prefix or "QuestInfoRewardsFrameQuestInfoItem") .. index]
			if prefix ~= "QuestRewardItem" then
				raw.type = "choice"
			end
			raw.GetID = function() return index end
			raw.IsVisible = function() return true end
			raw.GetFrameLevel = function() return 3 end
		end
		return buttons
	end

	--- The raw table behind a client frame's proxy (the client's own writes).
	local function raw(name)
		return state.widgets[name]
	end

	--- The client's first roll frame, showing roll `rollId` of `itemLink`.
	local function rollFrame(rollId, itemLink, name)
		name = name or "GroupLootFrame1"
		_G.GetLootRollItemLink = function(id)
			return id == rollId and itemLink or nil
		end
		state.extraGlobals[#state.extraGlobals + 1] = "GetLootRollItemLink"
		local frame = _G.CreateFrame("Frame", name)
		local inner = raw(name)
		inner.rollID = rollId
		inner.IsVisible = function() return true end
		inner.GetFrameLevel = function() return 4 end
		return frame
	end

	local function labelText(button)
		local label = RewardMarks.labelOn(button)
		if label == nil or not label.text:IsShown() or not label.overlay:IsShown() then
			return nil
		end
		return label.text:GetText()
	end

	local function lineText(frame)
		local line = RewardMarks.lineOn(frame)
		if line == nil or not line.text:IsShown() then
			return nil
		end
		return line.text:GetText()
	end

	after_each(function()
		mock.uninstall()
		for id in pairs(HIDDEN) do
			HIDDEN[id] = nil
		end
	end)

	-- Ruling 1: the reward label, the roll line and the tooltip's verdict agree.
	describe("one verdict source", function()
		it("agrees on a chest: the tooltip's Upgrade, the roll line, and Best on the reward", function()
			start({ chest = WORN_CHEST, back = WORN_BACK, feet = WORN_FEET })
			local lines = ItemTooltip.section(ItemTooltip.contextFor(DATA, BAG_CHEST))
			local verdict = string.format(L.tooltipVerdictDps, 2, "Worn Chest")
			assert.are.equal(verdict, lines[1].text)
			local judged = RewardMarks.judge(BAG_CHEST, ItemTooltip.baseContext(DATA))
			assert.are.equal(verdict, judged.primary.text)
			assert.are.equal(2, judged.gain)
			local frame = rollFrame(7, BAG_CHEST)
			RewardMarks.onEvent("START_LOOT_ROLL", 7, 60)
			assert.are.equal(verdict, lineText(frame))
			local buttons = offer({ THIN_CLOAK, BAG_CHEST, SOFT_BOOTS })
			RewardMarks.onEvent("QUEST_COMPLETE")
			mock.runTimers(state)
			assert.are.equal(string.format(L.rewardBest, "Arms"), labelText(buttons[2]))
		end)

		it("reads the section's own scoring: ItemTooltip.judge is what the section prints first", function()
			start({ chest = WORN_CHEST })
			local ctx = ItemTooltip.contextFor(DATA, PLATE_CHEST)
			local judged = ItemTooltip.judge(ctx, {})
			local lines = ItemTooltip.section(ctx)
			assert.are.equal(string.format(L.tooltipUntil, 40, "plate armor"), judged.primary.text)
			assert.are.equal(lines[1].text, judged.primary.text)
			assert.is_true(judged.gear)
			assert.are.equal(0, #judged.wearable)
		end)
	end)

	-- Ruling 2: the quest reward labels.
	describe("the quest reward", function()
		it("labels the one upgrade among three Best for Arms, and nothing else", function()
			start({ chest = WORN_CHEST, back = WORN_BACK, feet = WORN_FEET })
			local buttons = offer({ THIN_CLOAK, BAG_CHEST, SOFT_BOOTS })
			RewardMarks.onEvent("QUEST_COMPLETE")
			mock.runTimers(state)
			assert.is_nil(labelText(buttons[1]))
			assert.are.equal("Best for Arms", labelText(buttons[2]))
			assert.is_nil(labelText(buttons[3]))
			local label = RewardMarks.labelOn(buttons[2])
			assert.are.same(rgb3("gold"), { table.unpack(mock.lastCall(label.text, "SetTextColor"), 1, 3) })
		end)

		it("says Vendor on each of three when none is of any use", function()
			start({ chest = WORN_CHEST, back = WORN_BACK, feet = WORN_FEET })
			local buttons = offer({ THIN_CLOAK, RAG_CHEST, SOFT_BOOTS })
			RewardMarks.onEvent("QUEST_COMPLETE")
			mock.runTimers(state)
			for index = 1, 3 do
				assert.are.equal(L.rewardVendor, labelText(buttons[index]))
			end
		end)

		it("never calls a plate reward at 21 best, however it scores: the mail one is", function()
			start({ chest = WORN_CHEST, back = WORN_BACK })
			local buttons = offer({ PLATE_CHEST, BAG_CHEST, THIN_CLOAK })
			RewardMarks.onEvent("QUEST_COMPLETE")
			mock.runTimers(state)
			assert.is_nil(labelText(buttons[1]))
			assert.are.equal("Best for Arms", labelText(buttons[2]))
			assert.is_nil(labelText(buttons[3]))
		end)

		it("calls a plate reward at 21 Vendor with the rest when nothing else is of use", function()
			start({ chest = WORN_CHEST, back = WORN_BACK, feet = WORN_FEET })
			local buttons = offer({ PLATE_CHEST, THIN_CLOAK, SOFT_BOOTS })
			RewardMarks.onEvent("QUEST_COMPLETE")
			mock.runTimers(state)
			for index = 1, 3 do
				assert.are.equal(L.rewardVendor, labelText(buttons[index]))
			end
		end)

		it("says no Vendor while a choice is not gear (a potion may be of use)", function()
			start({ chest = WORN_CHEST, back = WORN_BACK })
			local buttons = offer({ THIN_CLOAK, POTION, RAG_CHEST })
			RewardMarks.onEvent("QUEST_COMPLETE")
			mock.runTimers(state)
			for index = 1, 3 do
				assert.is_nil(labelText(buttons[index]))
			end
		end)

		it("calls an off hand under a worn two-hander of no use", function()
			start({ main_hand = WORN_AXE, back = WORN_BACK })
			local labels = RewardMarks.rewardLabels({
				RewardMarks.judge(BUCKLER, ItemTooltip.baseContext(DATA)),
				RewardMarks.judge(THIN_CLOAK, ItemTooltip.baseContext(DATA)),
			}, 2, "Arms")
			assert.are.equal(L.rewardVendor, labels[1].text)
			assert.are.equal(L.rewardVendor, labels[2].text)
		end)

		it("calls a choice that fills an empty slot best at any gain that shows (+0.4 DPS)", function()
			start({ back = WORN_BACK })
			local buttons = offer({ THIN_CLOAK, LIGHT_VEST })
			RewardMarks.onEvent("QUEST_COMPLETE")
			mock.runTimers(state)
			assert.is_nil(labelText(buttons[1]))
			assert.are.equal("Best for Arms", labelText(buttons[2]))
		end)

		it("labels nothing with no spec weights to score by", function()
			start({ chest = WORN_CHEST })
			ItemTooltip.data = { classes = DATA.classes }
			RewardMarks.data = ItemTooltip.data
			local buttons = offer({ RAG_CHEST, THIN_CLOAK })
			RewardMarks.onEvent("QUEST_COMPLETE")
			mock.runTimers(state)
			assert.is_nil(labelText(buttons[1]))
			assert.is_nil(labelText(buttons[2]))
		end)

		it("waits for a choice the client has not sent, then labels on its item info", function()
			start({ chest = WORN_CHEST, back = WORN_BACK })
			HIDDEN[201] = true
			local buttons = offer({ THIN_CLOAK, BAG_CHEST })
			RewardMarks.onEvent("QUEST_COMPLETE")
			mock.runTimers(state)
			assert.is_nil(labelText(buttons[1]))
			assert.is_nil(labelText(buttons[2]))
			assert.is_true(RewardMarks.isWaiting(201))
			HIDDEN[201] = nil
			RewardMarks.onItemInfo(201)
			assert.are.equal("Best for Arms", labelText(buttons[2]))
		end)

		it("reads the classic reward buttons where the client has those", function()
			start({ chest = WORN_CHEST, back = WORN_BACK })
			local buttons = offer({ THIN_CLOAK, BAG_CHEST }, "QuestRewardItem")
			RewardMarks.onEvent("QUEST_COMPLETE")
			mock.runTimers(state)
			assert.are.equal("Best for Arms", labelText(buttons[2]))
		end)

		it("draws on the next frame, after the client has laid the buttons out", function()
			start({ chest = WORN_CHEST, back = WORN_BACK })
			local buttons = offer({ THIN_CLOAK, BAG_CHEST })
			RewardMarks.onEvent("QUEST_COMPLETE")
			assert.is_nil(RewardMarks.labelOn(buttons[2]))
			assert.are.equal(1, mock.runTimers(state))
			assert.are.equal("Best for Arms", labelText(buttons[2]))
		end)

		it("takes the labels away when the quest frame closes, and off the quest log's own display", function()
			start({ chest = WORN_CHEST, back = WORN_BACK }, { QuestInfo_Display = function() end })
			_G.CreateFrame("Frame", "QuestInfoFrame")
			local buttons = offer({ THIN_CLOAK, BAG_CHEST })
			RewardMarks.register()
			RewardMarks.onEvent("QUEST_COMPLETE")
			mock.runTimers(state)
			assert.are.equal("Best for Arms", labelText(buttons[2]))
			RewardMarks.onEvent("QUEST_FINISHED")
			assert.is_nil(labelText(buttons[2]))
			-- The quest log shows the same buttons for another quest: no label there.
			RewardMarks.onEvent("QUEST_COMPLETE")
			mock.runTimers(state)
			assert.are.equal("Best for Arms", labelText(buttons[2]))
			raw("QuestInfoFrame").questLog = true
			_G.QuestInfo_Display()
			assert.is_nil(labelText(buttons[2]))
		end)

		it("draws on an addon overlay parented to the reward button, a level above it, on a dark plate", function()
			start({ chest = WORN_CHEST, back = WORN_BACK })
			local buttons = offer({ THIN_CLOAK, BAG_CHEST })
			RewardMarks.onEvent("QUEST_COMPLETE")
			mock.runTimers(state)
			local label = RewardMarks.labelOn(buttons[2])
			assert.are.equal(buttons[2], label.overlay.parent)
			assert.is_nil(label.overlay.name)
			assert.are.equal(4, mock.lastCall(label.overlay, "SetFrameLevel")[1])
			assert.are.equal(0, mock.countCalls(label.overlay, "SetScript"))
			assert.are.equal(label.overlay, label.text.parent)
			assert.are.equal(label.overlay, label.plate.parent)
			assert.are.equal(0, #raw("QuestInfoRewardsFrameQuestInfoItem2").regions)
		end)
	end)

	-- Ruling 3: the roll window carries the tooltip's verdict line.
	describe("the roll window", function()
		it("says Upgrade with the gain and the worn item's name", function()
			start({ chest = WORN_CHEST })
			local frame = rollFrame(3, BAG_CHEST)
			RewardMarks.onEvent("START_LOOT_ROLL", 3, 60)
			assert.are.equal("Upgrade +2.0 DPS over Worn Chest", lineText(frame))
		end)

		it("says Sidegrade for an item no better than what is worn", function()
			start({ feet = WORN_FEET })
			local frame = rollFrame(3, SOFT_BOOTS)
			RewardMarks.onEvent("START_LOOT_ROLL", 3, 60)
			assert.are.equal(L.tooltipSide, lineText(frame))
		end)

		it("says Downgrade for an item a point or more worse", function()
			start({ chest = WORN_CHEST })
			local frame = rollFrame(3, RAG_CHEST)
			RewardMarks.onEvent("START_LOOT_ROLL", 3, 60)
			assert.are.equal(string.format(L.tooltipDown, 1, L.tooltipUnitDps, "Worn Chest"), lineText(frame))
		end)

		it("says Not for you until 40 on plate at 21", function()
			start({ chest = WORN_CHEST })
			local frame = rollFrame(3, PLATE_CHEST)
			RewardMarks.onEvent("START_LOOT_ROLL", 3, 60)
			assert.are.equal("Not for you until 40: plate armor", lineText(frame))
		end)

		it("says Not with your two-hander on a shield under a worn two-hander", function()
			start({ main_hand = WORN_AXE })
			local frame = rollFrame(3, BUCKLER)
			RewardMarks.onEvent("START_LOOT_ROLL", 3, 60)
			assert.are.equal(L.tooltipTwoHander, lineText(frame))
		end)

		it("says nothing on an item that is not gear", function()
			start({ chest = WORN_CHEST })
			local frame = rollFrame(3, POTION)
			RewardMarks.onEvent("START_LOOT_ROLL", 3, 60)
			assert.is_nil(lineText(frame))
		end)

		it("finds the frame on the next frame when the client opens it after the event", function()
			start({ chest = WORN_CHEST })
			_G.GetLootRollItemLink = function() return BAG_CHEST end
			state.extraGlobals[#state.extraGlobals + 1] = "GetLootRollItemLink"
			RewardMarks.onEvent("START_LOOT_ROLL", 9, 60)
			local frame = _G.CreateFrame("Frame", "GroupLootFrame2")
			raw("GroupLootFrame2").rollID = 9
			raw("GroupLootFrame2").IsVisible = function() return true end
			mock.runTimers(state)
			assert.are.equal("Upgrade +2.0 DPS over Worn Chest", lineText(frame))
		end)

		it("reads the modern container's pooled roll frames", function()
			start({ chest = WORN_CHEST })
			_G.GetLootRollItemLink = function() return BAG_CHEST end
			state.extraGlobals[#state.extraGlobals + 1] = "GetLootRollItemLink"
			local container = _G.CreateFrame("Frame", "GroupLootContainer")
			local pooled = _G.CreateFrame("Frame", nil, container)
			pooled.rollID = 4
			pooled.IsVisible = function() return true end
			raw("GroupLootContainer").rollFrames = { pooled }
			RewardMarks.onEvent("START_LOOT_ROLL", 4, 60)
			assert.are.equal("Upgrade +2.0 DPS over Worn Chest", lineText(pooled))
		end)

		it("draws beside the frame on an addon overlay that hides with it, in the verdict's colour", function()
			start({ chest = WORN_CHEST })
			local frame = rollFrame(3, BAG_CHEST)
			RewardMarks.onEvent("START_LOOT_ROLL", 3, 60)
			local line = RewardMarks.lineOn(frame)
			assert.are.equal(frame, line.overlay.parent)
			assert.is_nil(line.overlay.name)
			assert.are.same({ frame }, { mock.firstCall(line.overlay, "SetAllPoints")[1] })
			assert.are.equal(0, mock.countCalls(line.overlay, "SetScript"))
			assert.are.same({ "LEFT", line.overlay, "RIGHT" }, { table.unpack(line.plate.points[1], 1, 3) })
			assert.are.same(rgb3("success"), { table.unpack(mock.lastCall(line.text, "SetTextColor"), 1, 3) })
			assert.are.equal(0, #raw("GroupLootFrame1").regions)
		end)

		it("redraws a reused frame for its new roll", function()
			start({ chest = WORN_CHEST })
			local frame = rollFrame(3, BAG_CHEST)
			RewardMarks.onEvent("START_LOOT_ROLL", 3, 60)
			raw("GroupLootFrame1").rollID = 5
			_G.GetLootRollItemLink = function(id) return id == 5 and PLATE_CHEST or nil end
			RewardMarks.onEvent("START_LOOT_ROLL", 5, 60)
			assert.are.equal("Not for you until 40: plate armor", lineText(frame))
		end)
	end)

	-- Fix round 1, item 1: Best needs the tooltip's own Upgrade (the 1.0 floor over a worn
	-- item, any gain that shows for an empty slot); a 0 < gain < 1 choice gets no label.
	describe("fix round 1: Best needs an upgrade", function()
		it("gives a lone +0.3 choice (a Sidegrade on its tooltip) neither Best nor Vendor", function()
			start({ chest = WORN_CHEST, back = WORN_BACK, feet = WORN_FEET })
			local lines = ItemTooltip.section(ItemTooltip.contextFor(DATA, FAIR_CLOAK))
			assert.are.equal(string.format(L.tooltipSideGain, "+", 0.3, L.tooltipUnitDps), lines[1].text)
			local buttons = offer({ FAIR_CLOAK, RAG_CHEST, SOFT_BOOTS })
			RewardMarks.onEvent("QUEST_COMPLETE")
			mock.runTimers(state)
			for index = 1, 3 do
				assert.is_nil(labelText(buttons[index]))
			end
		end)

		it("puts Best on the +2.0 upgrade beside a +0.3 sidegrade", function()
			start({ chest = WORN_CHEST, back = WORN_BACK })
			local buttons = offer({ FAIR_CLOAK, BAG_CHEST })
			RewardMarks.onEvent("QUEST_COMPLETE")
			mock.runTimers(state)
			assert.is_nil(labelText(buttons[1]))
			assert.are.equal("Best for Arms", labelText(buttons[2]))
		end)

		it("calls an empty-slot +0.8 fill best: it is an upgrade on its tooltip", function()
			start({ back = WORN_BACK })
			local judged = RewardMarks.judge(FILL_VEST, ItemTooltip.baseContext(DATA))
			assert.are.equal(0.8, judged.gain)
			assert.is_true(judged.upgrade)
			assert.are.equal("success", judged.primary.color)
			local buttons = offer({ THIN_CLOAK, FILL_VEST })
			RewardMarks.onEvent("QUEST_COMPLETE")
			mock.runTimers(state)
			assert.are.equal("Best for Arms", labelText(buttons[2]))
		end)
	end)

	describe("fix round 1: the roll line", function()
		-- Item 2: a worn item whose name the client has not sent reads "…" until it comes.
		it("redraws when the worn item's name arrives", function()
			start({ chest = OLD_CHEST })
			HIDDEN[205] = true
			local frame = rollFrame(3, BAG_CHEST)
			RewardMarks.onEvent("START_LOOT_ROLL", 3, 60)
			assert.are.equal(string.format(L.tooltipVerdictDps, 2, L.tooltipNamePending), lineText(frame))
			assert.is_true(RewardMarks.isWaiting(205))
			HIDDEN[205] = nil
			RewardMarks.onItemInfo(205)
			assert.are.equal(string.format(L.tooltipVerdictDps, 2, "Old Chest"), lineText(frame))
			-- L3.7 item 20: the name-arrival redraw wrote nothing into the client's frame either.
			assert.are.same({}, state.clientWrites)
		end)

		-- Item 3: the 240 px fit, measured on the line's own font string.
		it("fits a long worn name within the 240 px line: the name on its own line, ellipsised", function()
			local function widest(text)
				local most = 0
				for part in (text or ""):gmatch("[^\n]+") do
					most = math.max(most, #part * 6)
				end
				return most
			end
			-- L3.7 item 20: the text measure is the mock's own (state.stringWidth), not a wrapped
			-- CreateFrame, so every font string the addon makes answers it.
			start({ chest = LONG_CHEST }, nil, nil, { stringWidth = widest })
			local frame = rollFrame(3, BAG_CHEST)
			RewardMarks.onEvent("START_LOOT_ROLL", 3, 60)
			local text = lineText(frame)
			local head, name = text:match("^(.-)\n(.+)$")
			assert.are.equal("Upgrade +2.0 DPS over", head)
			assert.are.equal(L.tooltipNamePending, name:sub(-#L.tooltipNamePending))
			assert.is_true(widest(text) <= 240)
			local plate = RewardMarks.lineOn(frame).plate
			assert.is_true(plate:GetWidth() <= 240 + require("Theme").SIZES.labelPad * 4)
		end)

		-- Item 4: a reused frame that no longer shows the roll drawn on it loses its line.
		it("clears a reused frame's old line when the new roll's frame cannot be found", function()
			start({ chest = WORN_CHEST })
			local frame = rollFrame(3, BAG_CHEST)
			RewardMarks.onEvent("START_LOOT_ROLL", 3, 60)
			assert.is_not_nil(lineText(frame))
			raw("GroupLootFrame1").rollID = 6
			RewardMarks.onEvent("START_LOOT_ROLL", 5, 60)
			mock.runTimers(state)
			assert.is_nil(lineText(frame))
			assert.are.same({}, state.clientWrites)
		end)

		it("clears a reused frame's old line when its new roll has no item link", function()
			start({ chest = WORN_CHEST })
			local frame = rollFrame(3, BAG_CHEST)
			RewardMarks.onEvent("START_LOOT_ROLL", 3, 60)
			raw("GroupLootFrame1").rollID = 5
			RewardMarks.onEvent("START_LOOT_ROLL", 5, 60)
			assert.is_nil(lineText(frame))
			assert.are.same({}, state.clientWrites)
		end)
	end)

	-- Item 5: the unread event table is gone.
	it("keeps no unread event list (Options names the events)", function()
		start()
		assert.is_nil(RewardMarks.EVENTS)
	end)

	-- Ruling 6: the labels follow "Gear tips on item tooltips".
	describe("the switch", function()
		it("hides every label and line with the tooltip switch off, and draws none", function()
			start({ chest = WORN_CHEST, back = WORN_BACK })
			local buttons = offer({ THIN_CLOAK, BAG_CHEST })
			local frame = rollFrame(3, BAG_CHEST)
			RewardMarks.onEvent("QUEST_COMPLETE")
			mock.runTimers(state)
			RewardMarks.onEvent("START_LOOT_ROLL", 3, 60)
			assert.is_not_nil(labelText(buttons[2]))
			assert.is_not_nil(lineText(frame))
			Prefs.setFlag("tooltip", false)
			RewardMarks.refresh(DATA)
			assert.is_nil(labelText(buttons[2]))
			assert.is_nil(lineText(frame))
			RewardMarks.onEvent("START_LOOT_ROLL", 3, 60)
			assert.is_nil(lineText(frame))
			Prefs.setFlag("tooltip", true)
			RewardMarks.refresh(DATA)
			assert.are.equal("Best for Arms", labelText(buttons[2]))
			assert.are.equal("Upgrade +2.0 DPS over Worn Chest", lineText(frame))
		end)
	end)

	-- The L4.1 taint rule: nothing is ever written into a client frame.
	describe("the client's frames", function()
		it("never writes into a reward button or a roll frame, through every path", function()
			start({ chest = WORN_CHEST, back = WORN_BACK }, { QuestInfo_Display = function() end })
			_G.CreateFrame("Frame", "QuestInfoFrame")
			local buttons = offer({ THIN_CLOAK, BAG_CHEST })
			local frame = rollFrame(3, BAG_CHEST)
			RewardMarks.register()
			RewardMarks.onEvent("QUEST_COMPLETE")
			mock.runTimers(state)
			RewardMarks.onEvent("QUEST_ITEM_UPDATE")
			mock.runTimers(state)
			_G.QuestInfo_Display()
			RewardMarks.onEvent("START_LOOT_ROLL", 3, 60)
			RewardMarks.onEvent("QUEST_FINISHED")
			Prefs.setFlag("tooltip", false)
			RewardMarks.refresh(DATA)
			Prefs.setFlag("tooltip", true)
			RewardMarks.refresh(DATA)
			RewardMarks.onItemInfo(201)
			assert.is_not_nil(RewardMarks.labelOn(buttons[2]))
			assert.is_not_nil(lineText(frame))
			assert.are.same({}, state.clientWrites)
		end)

		it("the mock watches the reward buttons and the roll frames", function()
			start()
			for _, name in ipairs({ "QuestInfoRewardsFrameQuestInfoItem1", "QuestRewardItem1", "GroupLootFrame1",
				"GroupLootContainer", "QuestInfoFrame" }) do
				local frame = _G.CreateFrame("Frame", name)
				local probe = assert(load("local frame = ...; frame:CreateFontString(nil, 'OVERLAY'); frame.fsLabel = 1",
					"@./ForeverSixty/surfaces/Probe.lua"))
				probe(frame)
			end
			assert.are.equal(10, #state.clientWrites)
		end)

		it("turns itself off on a failure and says so once, never a Lua error per event", function()
			start({ chest = WORN_CHEST })
			_G.GetLootRollItemLink = function() error("boom") end
			state.extraGlobals[#state.extraGlobals + 1] = "GetLootRollItemLink"
			rollFrame(3, BAG_CHEST)
			_G.GetLootRollItemLink = function() error("boom") end
			assert.has_no.errors(function()
				RewardMarks.onEvent("START_LOOT_ROLL", 3, 60)
				RewardMarks.onEvent("START_LOOT_ROLL", 3, 60)
			end)
			assert.is_true(RewardMarks.disabled)
			local notes = 0
			for _, line in ipairs(require("Theme").diagnostics()) do
				if line:find(L.diagRewardMarksFailed:sub(1, 20), 1, true) then
					notes = notes + 1
				end
			end
			assert.are.equal(1, notes)
		end)
	end)
end)

-- Ruling 4: the dungeon line, from the real data. The chip's "here" state reads
-- Progress.picksHere for the instance entered, names a pick the client has not sent
-- from Data.bisNames, and its click prints the picks with their bosses and chances.
describe("the dungeon line", function()
	local Chip
	local state

	local function enter(classToken, className, level, instance, instanceType, names)
		state = mock.install({
			class = { name = className, token = classToken }, level = level,
			globals = {
				UnitFactionGroup = function() return "Alliance" end,
				GetInstanceInfo = function() return instance, instanceType end,
				GetItemInfo = function(item)
					local name = (names or {})[item]
					if name == nil then
						return nil
					end
					return name, "|Hitem:" .. item .. "|h[" .. name .. "]|h", 2
				end,
				GetItemIcon = function(item) return "icon-" .. tostring(item) end,
			},
		})
		require("Theme").reset()
		helper.load("Compat")
		helper.load("Export")
		helper.load("Gear")
		helper.load("Talents")
		helper.load("Prefs")
		helper.load("Follow").build = nil
		helper.load("Progress")
		helper.load("ItemTooltip")
		Chip = helper.load("Chip")
		return require("Data")
	end

	after_each(function()
		mock.uninstall()
	end)

	it("names the warrior's band-20 picks in The Deadmines from the data, none cached yet", function()
		local Data = enter("WARRIOR", "Warrior", 21, "The Deadmines", "party")
		local model = Chip.model(Chip.inputs(Data))
		assert.are.equal("here", model.chosen.kind)
		assert.are.equal("Here for you: Smite's Mighty Hammer (Mr. Smite), 4 more", model.chosen.text)
		assert.is_true(Chip.isPending(7230))
		assert.are.equal("icon-7230", model.chosen.icon)
	end)

	-- The hunter's picks are read from the shipped Data, which the nightly re-ranks, so only
	-- the behaviour is pinned: the chip names the first of the band's Deadmines picks with its
	-- boss and counts the rest.
	it("names the hunter's band-20 pick in The Deadmines", function()
		local Data = enter("HUNTER", "Hunter", 24, "The Deadmines", "party")
		local bosses = {}
		for _, entry in pairs(Data.bis["hunter-beast-mastery"][20].alliance) do
			local id, boss = entry:match("^(%d+)|%a|The Deadmines: ([^|]+)")
			if id ~= nil then
				bosses[tonumber(id)] = boss
			end
		end
		local picks = require("Progress").picksHere(Data, "hunter-beast-mastery", 20, "alliance", "The Deadmines", {})
		assert(#picks > 0, "the shipped Data has no band-20 Deadmines pick for the hunter")
		local first = picks[1].itemId
		assert.is_not_nil(bosses[first])
		local more = #picks > 1 and string.format(", %d more", #picks - 1) or ""
		local model = Chip.model(Chip.inputs(Data))
		assert.are.equal("here", model.chosen.kind)
		assert.are.equal(string.format("Here for you: %s (%s)%s", Data.bisNames[first], bosses[first], more),
			model.chosen.text)
	end)

	it("takes the client's own name once it has it", function()
		local Data = enter("WARRIOR", "Warrior", 21, "The Deadmines", "party", { [7230] = "Smite's Hammer (client)" })
		local model = Chip.model(Chip.inputs(Data))
		assert.are.equal("Here for you: Smite's Hammer (client) (Mr. Smite), 4 more", model.chosen.text)
		assert.is_false(Chip.isPending(7230))
	end)

	it("has no here state outside an instance", function()
		local Data = enter("WARRIOR", "Warrior", 21, "Elwynn Forest", "none")
		local model = Chip.model(Chip.inputs(Data))
		assert.are_not.equal("here", model.chosen.kind)
		for _, chipState in ipairs(model.states) do
			assert.are_not.equal("here", chipState.kind)
		end
	end)

	it("keeps the count whole when the line is ellipsised by measured width", function()
		local Data = enter("WARRIOR", "Warrior", 21, "The Deadmines", "party")
		local chosen = Chip.model(Chip.inputs(Data)).chosen
		local function measure(text)
			return #text * 6
		end
		local fitted = Chip.fit(chosen.text, measure, 30 * 6, chosen.short, chosen.keep)
		assert.are.equal(", 4 more", fitted:sub(-#", 4 more"))
		assert.is_truthy(fitted:find(L.ellipsis, 1, true))
		assert.is_true(measure(fitted) <= 30 * 6)
		-- With room for the name but not the boss, the boss goes first.
		assert.are.equal("Here for you: Smite's Mighty Hammer, 4 more",
			Chip.fit(chosen.text, measure, 44 * 6, chosen.short, chosen.keep))
	end)

	it("prints the picks with their bosses and chances on a click", function()
		local Data = enter("WARRIOR", "Warrior", 21, "The Deadmines", "party")
		Chip.state(Data)
		Chip.printHere()
		local printed = table.concat(state.printed, "\n")
		assert.is_truthy(printed:find(string.format(L.chipHereHeading, "The Deadmines"), 1, true))
		assert.is_truthy(printed:find("Smite's Mighty Hammer · Mr. Smite · 20% drop", 1, true))
		assert.is_truthy(printed:find("Blackened Defias Gloves · Defias Overseer · 3% drop", 1, true))
		assert.are.equal(6, #state.printed)
	end)
end)
