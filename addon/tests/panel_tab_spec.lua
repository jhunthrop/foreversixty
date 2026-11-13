-- addon/tests/panel_tab_spec.lua
-- The character pane's "Forever Sixty" button (design/specs/2026-11-08-addon-ux.md 4.6): an
-- addon-owned button right of the client's last tab that opens the addon panel docked to the pane,
-- and the readers the panel shares (the next targets, the tooltip's own Next lines and the
-- off-plan talents). The taint rule: nothing is written into a client frame, the recorder stays empty.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local function line(spellId, name, condition)
	return { spellId = spellId, name = name, condition = condition }
end

local DATA = {
	build = "1.60.1.69893",
	classes = { warrior = { tabs = {
		{ name = "Arms", talents = {
			{ name = "A", tier = 1, column = 1, maxRank = 5 }, { name = "B", tier = 1, column = 2, maxRank = 1 },
		} },
		{ name = "Fury", talents = { { name = "C", tier = 1, column = 1, maxRank = 5 } } },
		{ name = "Protection", talents = {} },
	} } },
	weights = { ["warrior-arms"] = { strength = 2, attack_power = 1 } },
	dpsPerPoint = { ["warrior-arms"] = { reference = "attack_power", dps = 0.1, healing = false, weaponDps = 0.5,
		[20] = { reference = "attack_power", dps = 0.1, weaponDps = 0.5 },
		[60] = { reference = "attack_power", dps = 0.1, weaponDps = 0.5 } } },
	bis = { ["warrior-arms"] = {
		[20] = { alliance = {
			chest = "300|D|Wailing Caverns: Mutanus the Devourer|0.33",
			legs = "400|V|Bernard Brubaker",
		} },
		[60] = { alliance = {
			chest = "300|D|Wailing Caverns: Mutanus the Devourer|0.33",
			legs = "400|V|Bernard Brubaker",
		} },
	} },
	bisNames = { [300] = "Mutant Scale Breastplate", [400] = "Chausses of Westfall" },
	rotations = { ["warrior-arms"] = {
		{ level = 20, lines = {
			line(772, "Rend", "Rend kept up (on targets that live long enough). Then the rest."),
			line(5308, "Execute", "Execute below 20% health. It ends the fight."),
			line(1464, "Slam", "Slam inside the swing window."),
		} },
	} },
	abilityLevels = { [772] = 4, [5308] = 24, [1464] = 30 },
}

local function link(id, name)
	return "|Hitem:" .. id .. "::::::::21:::::|h[" .. name .. "]|h"
end

local WORN_CHEST = link(200, "Worn Chest")
local THIN_LEGS = link(410, "Thin Leggings")
local PICK_CHEST = link(300, "Mutant Scale Breastplate")
local PICK_LEGS = link(400, "Chausses of Westfall")

local ITEMS = {
	[WORN_CHEST] = { id = 200, slot = "INVTYPE_CHEST", class = 4, sub = 3, name = "Worn Chest",
		stats = { ITEM_MOD_STRENGTH_SHORT = 5 } },
	[THIN_LEGS] = { id = 410, slot = "INVTYPE_LEGS", class = 4, sub = 3, name = "Thin Leggings",
		stats = { ITEM_MOD_STRENGTH_SHORT = 4 } },
	[PICK_CHEST] = { id = 300, slot = "INVTYPE_CHEST", class = 4, sub = 3, name = "Mutant Scale Breastplate",
		stats = { ITEM_MOD_STRENGTH_SHORT = 10 } },
	[PICK_LEGS] = { id = 400, slot = "INVTYPE_LEGS", class = 4, sub = 3, name = "Chausses of Westfall",
		stats = { ITEM_MOD_STRENGTH_SHORT = 20 } },
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

--- Item ids the client has not sent yet (GetItemInfo answers nil for them).
local HIDDEN = {}

local SLOT_IDS = { chest = 5, legs = 7 }

--- The plan: Arms A three times and Fury C once; the character has A 2, B 1 and C 1:
--- 3 of 4 points on plan, B off plan.
local BUILD = { classSlug = "warrior", gear = {}, order = {
	{ tab = 1, tier = 1, column = 1 }, { tab = 1, tier = 1, column = 1 }, { tab = 2, tier = 1, column = 1 },
	{ tab = 1, tier = 1, column = 1 },
} }

describe("PanelTab", function()
	local PanelTab, ItemTooltip, Follow, Theme
	local state

	--- A client frame given a rect: GetLeft answers a number, as the client's does for a frame
	--- it has laid out (fix round 2: a hidden or unplaced frame answers nil).
	local function withRect(frame, left)
		frame.GetLeft = function()
			return left
		end
		return frame
	end

	--- A recorded point equals `expected` element by element, frames by identity (fix round 2:
	--- assert.are.same compares tables deeply, and every watched client frame is an empty proxy,
	--- so any two of them compared the same).
	local function assertPoint(expected, actual)
		assert.is_not_nil(actual)
		assert.are.equal(#expected, #actual)
		for index = 1, #expected do
			assert.is_true(rawequal(expected[index], actual[index]), "point element " .. index)
		end
	end

	--- `worn` is { [slot] = link }; `options` { level, inbox, noPane, tabs }. Answers the pane.
	local function start(worn, options)
		options = options or {}
		local equipped = {}
		for slot, itemLink in pairs(worn or {}) do
			equipped[SLOT_IDS[slot]] = itemLink
		end
		local itemStats = {}
		for itemLink, info in pairs(ITEMS) do
			itemStats[itemLink] = info.stats
		end
		state = mock.install({
			class = { name = "Warrior", token = "WARRIOR" },
			race = { name = "Human", token = "Human" },
			realm = "Ashbringer",
			region = 1,
			playerName = "Bob",
			level = options.level or 21,
			equipped = equipped,
			itemStats = itemStats,
			talents = {
				{ name = "Arms", talents = {
					{ name = "A", tier = 1, column = 1, rank = 2, maxRank = 5 },
					{ name = "B", tier = 1, column = 2, rank = 1, maxRank = 1 },
				} },
				{ name = "Fury", talents = { { name = "C", tier = 1, column = 1, rank = 1, maxRank = 5 } } },
				{ name = "Protection", talents = {} },
			},
			globals = {
				UnitFactionGroup = function() return "Alliance" end,
				GetItemInfoInstant = function(item)
					local info = lookup(item)
					if info == nil then
						return nil
					end
					return info.id, "Armor", "Mail", info.slot, nil, info.class, info.sub
				end,
				GetItemInfo = function(item)
					local info = lookup(item)
					if info == nil or HIDDEN[info.id] then
						return nil
					end
					return info.name, info.link, 2
				end,
				GetItemIconByID = function(id)
					return "icon-" .. tostring(id)
				end,
				GetSpellTexture = function(id)
					return "spell-" .. tostring(id)
				end,
				GetSpellSubtext = function()
					return "Rank 2"
				end,
				time = function()
					return 0
				end,
			},
		})
		_G.ForeverSixtyInbox = options.inbox
		Theme = require("Theme")
		Theme.reset()
		helper.load("Compat")
		helper.load("Export")
		helper.load("Gear")
		helper.load("Talents")
		helper.load("Prefs")
		Follow = helper.load("Follow")
		helper.load("Progress")
		helper.load("Rotation")
		helper.load("Widgets")
		ItemTooltip = helper.load("ItemTooltip")
		helper.load("Chip")
		helper.load("BuildDialog")
		PanelTab = helper.load("PanelTab")
		Follow.build = BUILD
		ItemTooltip.data = DATA
		PanelTab.data = DATA
		local pane
		if not options.noPane then
			pane = _G.CreateFrame("Frame", "CharacterFrame", _G.UIParent)
			for index = 1, options.tabs or 3 do
				withRect(_G.CreateFrame("Button", "CharacterFrameTab" .. index, pane), 20 + 60 * index)
			end
		end
		return pane
	end

	after_each(function()
		mock.uninstall()
		for id in pairs(HIDDEN) do
			HIDDEN[id] = nil
		end
	end)

	local function click(button)
		button:GetScript("OnClick")(button, "LeftButton")
	end

	--- A worn item's section as a stub tooltip answers it: a Next line naming `itemId`
	--- at `gain`, and its source line.
	local function section(itemId, gain, source, pending)
		return { lines = {
			{ kind = "next", itemId = itemId, upgrade = { name = "Item " .. itemId, hex = "1eff00", gain = gain,
				unit = "dps" } },
			{ kind = "source", text = source or ("Source " .. itemId), pick = { item_id = itemId } },
		}, pending = pending or {} }
	end

	-- Ruling 2: one verdict source. A target is the very Next line the worn item's own
	-- tooltip prints (ItemTooltip.sectionFor), with its source line.
	describe("targets", function()
		it("lists, per worn slot, the pick and gain the tooltip's Next line prints, with its source line", function()
			start({ chest = WORN_CHEST, legs = THIN_LEGS })
			local base = ItemTooltip.baseContext(DATA)
			local rows = PanelTab.targets(base, function(itemLink)
				return ItemTooltip.sectionFor(itemLink, base)
			end)
			assert.are.equal(2, #rows)
			for _, row in ipairs(rows) do
				local worn = row.slot == "chest" and WORN_CHEST or THIN_LEGS
				local lines = ItemTooltip.section(ItemTooltip.contextFor(DATA, worn))
				local nextLine, sourceLine
				for _, entry in ipairs(lines) do
					if entry.kind == "next" then
						nextLine = entry
					elseif entry.kind == "source" then
						sourceLine = entry
					end
				end
				assert.are.equal(nextLine.itemId, row.itemId)
				assert.are.equal(nextLine.upgrade.gain, row.gain)
				assert.are.equal(nextLine.upgrade.name, row.name)
				assert.are.equal(sourceLine.text, row.source)
			end
			-- Legs gain 3.2 DPS (16 Str x 2 x 0.1), the chest 1.0: the biggest first.
			assert.are.equal("legs", rows[1].slot)
			assert.are.equal(3.2, rows[1].gain)
			assert.are.equal("chest", rows[2].slot)
			assert.are.equal(ItemTooltip.sourceLine(require("Gear").pick("warrior-arms", 20, "alliance", "chest", DATA)),
				rows[2].source)
		end)

		it("leaves a slot whose hover prints no Next line, and an empty slot", function()
			start({ chest = PICK_CHEST })
			local base = ItemTooltip.baseContext(DATA)
			local rows = PanelTab.targets(base, function(itemLink)
				return ItemTooltip.sectionFor(itemLink, base)
			end)
			assert.are.same({}, rows)
		end)

		it("keeps five, biggest gain first, a melee spec's weapons before armour at an equal gain", function()
			start()
			local sections = {
				legs = section(1, 5.0), chest = section(2, 2.0), main_hand = section(3, 2.0), head = section(4, 1.5),
				feet = section(5, 1.2), hands = section(6, 1.1),
			}
			local equipped, bySlot = {}, {}
			for slot in pairs(sections) do
				equipped[slot] = { link = "worn-" .. slot, itemId = 900 }
				bySlot["worn-" .. slot] = sections[slot]
			end
			local function order(spec)
				local rows = PanelTab.targets({ spec = spec, equipped = equipped, isCached = function() return true end },
					function(itemLink) return bySlot[itemLink] end)
				local slots = {}
				for index, row in ipairs(rows) do
					slots[index] = row.slot
				end
				return slots
			end
			assert.are.same({ "legs", "main_hand", "chest", "head", "feet" }, order("warrior-arms"))
			assert.are.same({ "legs", "chest", "main_hand", "head", "feet" }, order("mage-frost"))
		end)

		it("names a pick the client has not sent from the data, and waits on it", function()
			start()
			local equipped = { chest = { link = "worn-chest", itemId = 900 } }
			local rows, waiting = PanelTab.targets({ spec = "warrior-arms", equipped = equipped,
				isCached = function() return true end }, function()
				return section(300, 1.0, nil, { [300] = true })
			end)
			assert.are.equal(1, #rows)
			assert.is_true(rows[1].pending)
			assert.is_true(waiting[300])
		end)

		it("waits on a worn item the client has not sent rather than list a Next line it does not mean", function()
			start()
			local equipped = { chest = { link = "worn-chest", itemId = 900 } }
			local rows, waiting = PanelTab.targets({ spec = "warrior-arms", equipped = equipped,
				isCached = function(id) return id ~= 900 end }, function()
				return section(300, 1.0)
			end)
			assert.are.same({}, rows)
			assert.is_true(waiting[900])
		end)

		it("lists a pick two slots name once, at its bigger gain", function()
			start()
			local equipped = { finger1 = { link = "ring-1", itemId = 901 }, finger2 = { link = "ring-2", itemId = 902 } }
			local rows = PanelTab.targets({ spec = "warrior-arms", equipped = equipped,
				isCached = function() return true end }, function(itemLink)
				return section(500, itemLink == "ring-1" and 1.0 or 2.0)
			end)
			assert.are.equal(1, #rows)
			assert.are.equal("finger2", rows[1].slot)
			assert.are.equal(2.0, rows[1].gain)
		end)
	end)

	-- Ruling 4: the split in the dialog's words, progress and the off-plan count.

	describe("the button", function()
		it("is made on the first refresh that finds the pane, right of the client's last tab", function()
			local pane = start({ chest = WORN_CHEST })
			PanelTab.refresh(DATA)
			local button = PanelTab.button
			assert.is_not_nil(button)
			assert.are.equal(pane, button.parent)
			assertPoint({ "LEFT", _G.CharacterFrameTab3, "RIGHT" }, { table.unpack(button.points[1], 1, 3) })
			assert.are.same({}, state.clientWrites)
		end)

		it("sits under the pane's bottom right when no client tab has a rect", function()
			local pane = start({}, { tabs = 0 })
			PanelTab.refresh(DATA)
			assertPoint({ "TOPRIGHT", pane, "BOTTOMRIGHT" }, { table.unpack(PanelTab.button.points[1], 1, 3) })
		end)

		it("makes nothing without a pane, and nothing in combat until combat ends", function()
			start({}, { noPane = true })
			PanelTab.refresh(DATA)
			assert.is_nil(PanelTab.button)
			start({})
			Theme.inCombat = function() return true end
			PanelTab.refresh(DATA)
			assert.is_nil(PanelTab.button)
			assert.is_true(PanelTab.deferred)
			Theme.inCombat = function() return false end
			PanelTab.onCombatEnd()
			assert.is_not_nil(PanelTab.button)
		end)

		-- design/specs/2026-11-13-addon-hub.md 5: the button is the door to the window's Gear tab.
		it("opens the Forever Sixty window on Gear, and closes it when Gear shows", function()
			start({})
			PanelTab.refresh(DATA)
			local Hub = require("Hub")
			local asked = {}
			local toggle = Hub.toggle
			Hub.toggle = function(id)
				asked[#asked + 1] = id
				return id
			end
			click(PanelTab.button)
			Hub.toggle = toggle
			assert.are.same({ "gear" }, asked)
			assert.is_nil(_G.ForeverSixtyPanelFlyout)
		end)

		it("says what it opens on hover", function()
			start({})
			local lines = PanelTab.buttonHint()
			assert.are.equal(2, #lines)
			assert.is_truthy(lines[2]:find(L.panelTabOpens, 1, true))
		end)

		it("reports the panel open or closed in /fs diag", function()
			start({})
			assert.are.equal(string.format(L.diagPanelState, L.diagPanelClosed), PanelTab.diagLine())
		end)
	end)
end)
