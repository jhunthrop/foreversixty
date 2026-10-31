-- addon/tests/panel_tab_spec.lua
-- The "Forever Sixty" character-panel tab (spec section 4.5, task L4.3): an addon-owned
-- button at the character pane's bottom edge, right of the client's last tab, and an
-- addon-owned flyout at the pane's right holding up to five regions: next targets,
-- rotation, build, and at 60 the sim and the last raid. Ruling 1 (taint): nothing is
-- written into a client frame, the recorder stays empty; ruling 2: the targets are the
-- tooltip's own Next lines; regions with nothing to show are absent.
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
	local PanelTab, ItemTooltip, Follow, Export, Theme, BuildDialog
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
		Export = helper.load("Export")
		helper.load("Gear")
		helper.load("Talents")
		helper.load("Prefs")
		Follow = helper.load("Follow")
		helper.load("Progress")
		helper.load("Rotation")
		helper.load("Widgets")
		ItemTooltip = helper.load("ItemTooltip")
		helper.load("Chip")
		BuildDialog = helper.load("BuildDialog")
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

	local function kinds(model)
		local list = {}
		for index, region in ipairs(model.regions) do
			list[index] = region.kind
		end
		return list
	end

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

	-- Ruling 3: the band's priority list, plain words, abilities not learned yet muted and last.
	describe("rotation", function()
		it("lists the learned lines in priority order, then the unlearned ones marked at level N", function()
			start()
			local rows = PanelTab.rotationRows(DATA, BUILD, { [1] = { ["1:1"] = 3 } }, 21)
			local names = {}
			for index, row in ipairs(rows) do
				names[index] = row.name
			end
			assert.are.same({ "Rend", "Execute", "Slam" }, names)
			assert.are.equal(1, rows[1].number)
			assert.is_false(rows[1].muted)
			assert.are.equal("Rend kept up.", rows[1].condition)
			assert.is_nil(rows[2].number)
			assert.is_true(rows[2].muted)
			assert.are.equal(string.format(L.panelRotationAt, 24), rows[2].at)
			assert.are.equal("Execute below 20% health.", rows[2].condition)
			assert.are.equal(string.format(L.panelRotationAt, 30), rows[3].at)
		end)

		it("numbers every line once the character has learned them all", function()
			start()
			local rows = PanelTab.rotationRows(DATA, BUILD, { [1] = { ["1:1"] = 3 } }, 30)
			assert.are.same({ 1, 2, 3 }, { rows[1].number, rows[2].number, rows[3].number })
			assert.is_nil(rows[2].at)
		end)

		it("is empty with no build, or a spec with no rotation", function()
			start()
			assert.are.same({}, PanelTab.rotationRows(DATA, nil, {}, 21))
			-- A followed build names the spec (hotfix 2.0.2): a Fury plan has no rotation in the data.
			local fury = { classSlug = "warrior", order = { { tab = 2, tier = 1, column = 1 },
				{ tab = 2, tier = 1, column = 1 } } }
			assert.are.same({}, PanelTab.rotationRows(DATA, fury, {}, 21))
		end)
	end)

	-- Ruling 4: the split in the dialog's words, progress and the off-plan count.
	describe("build", function()
		it("reads the split, the points on plan and the points off it", function()
			start()
			local ranks = require("Talents").readRanks(DATA)
			local region = PanelTab.buildRegion(BUILD, DATA, ranks, 21)
			-- Fix round 3, item 1: the split reads as the plan ("Plan: Arms 3 / Fury 1").
			assert.are.equal(string.format(L.panelPlanSplit, string.format(L.dialogSplitPart, "Arms", 3)
				.. L.dialogSplitSeparator .. string.format(L.dialogSplitPart, "Fury", 1)), region.split)
			assert.are.equal(string.format(L.panelProgress, 3, 4), region.progress)
			assert.are.equal(string.format(L.panelOffPlan, 1), region.offPlan)
		end)

		it("has no off-plan line at none, and no region with no build", function()
			start()
			local region = PanelTab.buildRegion(BUILD, DATA, { [1] = { ["1:1"] = 3 }, [2] = { ["1:1"] = 1 } }, 21)
			assert.is_nil(region.offPlan)
			assert.are.equal(string.format(L.panelProgress, 4, 4), region.progress)
			assert.is_nil(PanelTab.buildRegion(nil, DATA, {}, 21))
		end)
	end)

	-- Ruling 5: at 60 only, from the inbox's own messages.
	describe("sim", function()
		local SIM = { tag = "sim", spec = "warrior-arms", dpsPerPoint = 0.42, healing = false,
			weights = { attack_power = 1, strength = 2.2, hit = 0, agility = 1.5 } }

		it("reads your DPS and the top three stat weights in DPS per point", function()
			start()
			local region = PanelTab.simRegion(SIM, { spec = "warrior-arms", dps = 412.6 }, 60, false)
			assert.are.equal(string.format(L.panelSimDps, 412.6, L.tooltipUnitDps), region.dps)
			assert.are.same({
				string.format(L.panelSimWeight, L.statNames.strength, 2.2 * 0.42, L.tooltipUnitDps),
				string.format(L.panelSimWeight, L.statNames.agility, 1.5 * 0.42, L.tooltipUnitDps),
				string.format(L.panelSimWeight, L.statNames.attack_power, 0.42, L.tooltipUnitDps),
			}, region.weights)
			assert.is_false(region.queued)
		end)

		it("leaves the DPS out for another spec's sim, and says when a sim is queued", function()
			start()
			local region = PanelTab.simRegion(SIM, { spec = "warrior-fury", dps = 500 }, 60, true)
			assert.is_nil(region.dps)
			assert.is_true(region.queued)
		end)

		it("is absent below 60 and without a weights message", function()
			start()
			assert.is_nil(PanelTab.simRegion(SIM, { spec = "warrior-arms", dps = 412.6 }, 59, false))
			assert.is_nil(PanelTab.simRegion({ tag = "default", weights = { strength = 2 } }, nil, 60, false))
			assert.is_nil(PanelTab.simRegion(nil, nil, 60, false))
		end)
	end)

	describe("last raid", function()
		local RAID = { type = "last_raid", date = "2026-12-09", bosses = {
			{ name = "Lucifron", percentile = 71, biggest_loss = "downtime", loss_seconds = 18, deaths = 0,
				consumables_missed = { "flask" }, adherence = 0.82 },
			{ name = "Magmadar", percentile = 40, deaths = 1, consumables_missed = {} },
		} }

		it("reads each boss: percentile, biggest loss, deaths, consumables missed, rotation adherence", function()
			start()
			local region = PanelTab.raidRegion(RAID, 60)
			assert.are.equal(string.format(L.panelLastRaidAsOf, "9 Dec"), region.title)
			assert.are.equal(string.format(L.panelRaidBoss, "Lucifron", "71st"), region.bosses[1].title)
			assert.are.equal(table.concat({ string.format(L.panelRaidLoss, "downtime"),
				string.format(L.panelRaidDeaths, 0), string.format(L.panelRaidMissed, "flask"),
				string.format(L.panelRaidAdherence, 82) }, L.panelRaidJoin), region.bosses[1].detail)
			assert.are.equal(string.format(L.panelRaidDeath, 1), region.bosses[2].detail)
		end)

		it("is absent below 60 and without a message", function()
			start()
			assert.is_nil(PanelTab.raidRegion(RAID, 59))
			assert.is_nil(PanelTab.raidRegion(nil, 60))
			assert.is_nil(PanelTab.raidRegion({ bosses = {} }, 60))
		end)
	end)

	describe("the model", function()
		local INBOX = { messages = {
			{ type = "weights", character = "US/Ashbringer/Bob", spec = "warrior-arms",
				weights = { { stat = "attack_power", weight = 1 }, { stat = "strength", weight = 2.2 } },
				caps = {}, dps_per_point = 0.42 },
			{ type = "sim_result", character = "US/Ashbringer/Bob", spec = "warrior-arms", dps = 412.6, error = 3.1,
				generated_at = "2026-11-04T20:55:00Z", slots = {} },
			{ type = "last_raid", character = "US/Ashbringer/Bob", date = "2026-12-09",
				bosses = { { name = "Lucifron", percentile = 71, deaths = 0, consumables_missed = {} } } },
		} }

		it("holds targets, rotation and build below 60, never the sim or the last raid", function()
			start({ chest = WORN_CHEST }, { inbox = INBOX })
			assert.are.same({ "targets", "rotation", "build" }, kinds(PanelTab.model(PanelTab.inputs(DATA))))
		end)

		it("adds the sim and the last raid at 60, read through the inbox's own readers", function()
			start({ chest = WORN_CHEST }, { level = 60, inbox = INBOX })
			local model = PanelTab.model(PanelTab.inputs(DATA))
			assert.are.same({ "targets", "rotation", "build", "sim", "raid" }, kinds(model))
			assert.are.equal(string.format(L.panelSimDps, 412.6, L.tooltipUnitDps), model.regions[4].dps)
		end)

		it("leaves out every region with nothing to show", function()
			start({ chest = PICK_CHEST })
			Follow.build = nil
			assert.are.same({}, kinds(PanelTab.model(PanelTab.inputs(DATA))))
		end)
	end)

	-- Ruling 1: an addon-owned button and flyout parented to the pane, nothing written into
	-- a client frame, no OnUpdate; ruling 7: the flyout is built on the first click.
	describe("the frames", function()
		it("puts the button at the pane's bottom edge right of the client's last shown tab", function()
			local pane = start({ chest = WORN_CHEST })
			PanelTab.refresh(DATA)
			local button = _G[PanelTab.BUTTON_NAME]
			assert.is_not_nil(button)
			assert.are.equal(pane, button.parent)
			assertPoint({ "LEFT", _G.CharacterFrameTab3, "RIGHT", Theme.SIZES.panelTabGap, 0 },
				button.points[#button.points])
			assert.are.equal(L.panelTabLabel, button:GetText())
			-- The flyout waits for the first click.
			assert.is_nil(_G[PanelTab.FLYOUT_NAME])
			_G.CharacterFrameTab3:Hide()
			PanelTab.refresh(DATA)
			assert.are.equal(_G.CharacterFrameTab2, button.points[#button.points][2])
			assert.are.same({}, state.clientWrites)
		end)

		it("falls back to the pane's bottom right with no tab to sit beside", function()
			local pane = start({}, { tabs = 0 })
			PanelTab.refresh(DATA)
			local button = _G[PanelTab.BUTTON_NAME]
			assertPoint({ "TOPRIGHT", pane, "BOTTOMRIGHT", 0, 0 }, button.points[#button.points])
			assert.are.same({}, state.clientWrites)
		end)

		it("turns itself off on a failure, hides what it drew and says so once in /fs diag", function()
			start({ chest = WORN_CHEST })
			PanelTab.refresh(DATA)
			PanelTab.draw = function()
				error("boom")
			end
			click(_G[PanelTab.BUTTON_NAME])
			assert.is_true(PanelTab.disabled)
			assert.is_false(_G[PanelTab.BUTTON_NAME]:IsShown())
			assert.is_false(_G[PanelTab.FLYOUT_NAME]:IsShown())
			local notes = Theme.diagnostics()
			assert.is_truthy(notes[#notes]:find(L.diagPanelFailed:match("^[^%%]+"), 1, true))
			PanelTab.refresh(DATA)
			assert.are.same({}, state.clientWrites)
		end)

		it("makes nothing and fails nothing on a client with no character pane", function()
			start({}, { noPane = true })
			assert.has_no.errors(function()
				PanelTab.refresh(DATA)
			end)
			assert.is_nil(_G[PanelTab.BUTTON_NAME])
			assert.are.same({}, state.clientWrites)
		end)

		it("opens the flyout at the pane's top right on a click, and closes it on the next", function()
			local pane = start({ chest = WORN_CHEST })
			PanelTab.refresh(DATA)
			click(_G[PanelTab.BUTTON_NAME])
			local flyout = _G[PanelTab.FLYOUT_NAME]
			assert.are.equal(pane, flyout.parent)
			assertPoint({ "TOPLEFT", pane, "TOPRIGHT", Theme.SIZES.panelOffsetX, Theme.SIZES.panelOffsetY },
				flyout.points[1])
			assert.is_true(flyout:IsShown())
			click(_G[PanelTab.BUTTON_NAME])
			assert.is_false(flyout:IsShown())
			assert.are.equal(0, mock.countCalls(flyout, "SetScript"))
			assert.are.same({}, state.clientWrites)
		end)

		it("sizes the flyout to its regions", function()
			start({ chest = WORN_CHEST })
			PanelTab.refresh(DATA)
			click(_G[PanelTab.BUTTON_NAME])
			local flyout = _G[PanelTab.FLYOUT_NAME]
			local full = flyout:GetHeight()
			assert.is_true(full > 0)
			Follow.build = nil
			PanelTab.refresh(DATA)
			local fewer = flyout:GetHeight()
			assert.is_true(fewer > 0 and fewer < full)
			for _, kind in ipairs({ "rotation", "build" }) do
				assert.is_false(PanelTab.regionFrames[kind]:IsShown(), kind)
			end
			assert.is_true(PanelTab.regionFrames.targets:IsShown())
		end)

		it("says there is nothing to show yet when every region is absent", function()
			start({ chest = PICK_CHEST })
			Follow.build = nil
			PanelTab.refresh(DATA)
			click(_G[PanelTab.BUTTON_NAME])
			assert.is_true(PanelTab.emptyLine:IsShown())
			assert.are.equal(L.panelEmpty, PanelTab.emptyLine:GetText())
		end)

		it("draws a target row: icon, quality-coloured name, delta, slot and source line", function()
			start({ chest = WORN_CHEST })
			PanelTab.refresh(DATA)
			click(_G[PanelTab.BUTTON_NAME])
			local row = PanelTab.targetRows[1]
			assert.are.equal("icon-300", mock.lastCall(row.icon, "SetTexture")[1])
			assert.is_truthy(row.name:GetText():find("Mutant Scale Breastplate", 1, true))
			assert.is_truthy(row.name:GetText():find("|cff", 1, true))
			assert.are.equal(ItemTooltip.figureText(1.0, "dps"), row.gain:GetText())
			-- L3.7 item 22: "· over <worn item>" is muted in the name line's empty gap (it was the end of
			-- the source line, where it rarely fitted); the source line is the tooltip's own.
			assert.are.equal(string.format(L.panelTargetDetail, L.gearSlotLabels.chest,
				"Wailing Caverns · Mutanus the Devourer · 33% drop"), row.detail:GetText())
			assert.is_truthy(row.name:GetText():find("|cff" .. Theme.HEX.muted
				.. string.format(L.panelTargetOver, "Worn Chest") .. "|r", 1, true))
		end)

		it("opens the build dialog and the copy from the build region's two buttons", function()
			start({ chest = WORN_CHEST })
			local opened = {}
			BuildDialog.open = function(mode)
				opened[#opened + 1] = mode
				return true
			end
			PanelTab.refresh(DATA)
			click(_G[PanelTab.BUTTON_NAME])
			local build = PanelTab.regionFrames.build
			assert.are.equal(L.panelUseOther, build.useOther:GetText())
			assert.are.equal(L.dialogCopy, build.copy:GetText())
			click(build.useOther)
			click(build.copy)
			assert.are.same({ "paste", "copy" }, opened)
			assert.are.same({}, state.clientWrites)
		end)

		it("queues a sim from Sim me through the outbox, then reads that it is queued", function()
			start({ chest = WORN_CHEST }, { level = 60, inbox = { messages = {
				{ type = "weights", character = "US/Ashbringer/Bob", spec = "warrior-arms",
					weights = { { stat = "strength", weight = 2 } }, caps = {}, dps_per_point = 0.1 },
			} } })
			PanelTab.refresh(DATA)
			click(_G[PanelTab.BUTTON_NAME])
			local sim = PanelTab.regionFrames.sim
			assert.is_true(sim:IsShown())
			assert.are.equal(L.panelSimMe, sim.button:GetText())
			-- The hover names the sim entitlement.
			sim.button:GetScript("OnEnter")(sim.button)
			local said = {}
			for _, call in ipairs(_G.GameTooltip.calls) do
				if call.method == "AddLine" then
					said[#said + 1] = call[1]
				end
			end
			assert.are.same(L.panelSimMeHint, said)
			click(sim.button)
			local request = Export.simRequest
			local keys = {}
			for key in pairs(request) do
				keys[#keys + 1] = key
			end
			table.sort(keys)
			assert.are.same({ "character", "gear_code", "requested_at", "spec" }, keys)
			assert.are.equal("warrior-arms", request.spec)
			assert.are.equal("FS1:", request.gear_code:sub(1, 4))
			assert.are.same(request, _G.ForeverSixtyDB.outbox.sim_request)
			assert.are.equal(L.panelSimQueued, sim.button:GetText())
			assert.are.same({}, state.clientWrites)
		end)

		it("redraws on an item it waits on, and on nothing else", function()
			HIDDEN[300] = true
			start({ chest = WORN_CHEST })
			PanelTab.refresh(DATA)
			click(_G[PanelTab.BUTTON_NAME])
			-- The pick's stats are on their way: the tooltip prints no Next line yet, so neither
			-- does the panel, and it waits on the item as the tooltip does.
			assert.is_false(PanelTab.regionFrames.targets:IsShown())
			assert.is_true(PanelTab.waiting[300])
			local draws = PanelTab.draws
			PanelTab.onItemInfo(999)
			assert.are.equal(draws, PanelTab.draws)
			HIDDEN[300] = nil
			PanelTab.onItemInfo(300)
			assert.are.equal(draws + 1, PanelTab.draws)
			assert.is_true(PanelTab.regionFrames.targets:IsShown())
			assert.is_truthy(PanelTab.targetRows[1].name:GetText():find("Mutant Scale Breastplate", 1, true))
		end)

		it("does not draw a closed flyout on an event", function()
			start({ chest = WORN_CHEST })
			PanelTab.refresh(DATA)
			click(_G[PanelTab.BUTTON_NAME])
			click(_G[PanelTab.BUTTON_NAME])
			local draws = PanelTab.draws
			PanelTab.refresh(DATA)
			assert.are.equal(draws, PanelTab.draws)
		end)

		it("prints its /fs diag line: open or closed, the targets and the rotation rows", function()
			start({ chest = WORN_CHEST })
			PanelTab.refresh(DATA)
			assert.are.equal(string.format(L.diagPanel, L.diagPanelClosed, 1, 3), PanelTab.diagLine())
			click(_G[PanelTab.BUTTON_NAME])
			assert.are.equal(string.format(L.diagPanel, L.diagPanelOpen, 1, 3), PanelTab.diagLine())
		end)

		it("never writes into a client frame through any path, and never polls or hooks", function()
			start({ chest = WORN_CHEST, legs = THIN_LEGS }, { level = 60 })
			PanelTab.refresh(DATA)
			click(_G[PanelTab.BUTTON_NAME])
			PanelTab.refresh(DATA)
			PanelTab.onItemInfo(300)
			click(_G[PanelTab.BUTTON_NAME])
			assert.are.same({}, state.clientWrites)
			local file = assert(io.open("ForeverSixty/surfaces/PanelTab.lua", "r"))
			local source = file:read("a")
			file:close()
			-- No call of these, and no script by these names (the header's prose names them).
			for _, call in ipairs({ "HookScript(", "hooksecurefunc(", "PanelTemplates_SetNumTabs(", '"OnUpdate"',
				'"OnHide"' }) do
				assert.is_nil(source:find(call, 1, true), call)
			end
			-- One OnShow (fix round 2, re-review b): on the addon's own button, never on the pane.
			local _, onShows = source:gsub('"OnShow"', "")
			assert.are.equal(1, onShows)
			assert.is_truthy(source:find('button:SetScript("OnShow"', 1, true))
		end)
	end)

	-- Fix round 1 (the code review and the live client's hunter-23 captures).
	describe("fix round 1", function()
		--- A measure of 6 units a visible byte, for the fits.
		local function sixPerByte(text)
			return #ItemTooltip.stripEscapes(text) * 6
		end

		local function open(worn, options)
			local pane = start(worn, options)
			PanelTab.refresh(DATA)
			click(_G[PanelTab.BUTTON_NAME])
			return pane, _G[PanelTab.FLYOUT_NAME]
		end

		local function regionOf(model, kind)
			for _, region in ipairs(model.regions) do
				if region.kind == kind then
					return region
				end
			end
			return nil
		end

		local function hasPoint(region, point)
			for _, candidate in ipairs(region.points) do
				local same = #candidate == #point
				for index = 1, #point do
					same = same and candidate[index] == point[index]
				end
				if same then
					return true
				end
			end
			return false
		end

		local function rgb(key, alpha)
			return { Theme.rgb(Theme.HEX[key], alpha) }
		end

		local function colorOf(region, method)
			local call = mock.lastCall(region, method or "SetTextColor")
			return { call[1], call[2], call[3], call[4] }
		end

		--- The shown texts of a region's pooled lines, top to bottom.
		local function shownTexts(region)
			local list = {}
			for _, text in ipairs(region.texts or {}) do
				if text:IsShown() then
					list[#list + 1] = text:GetText()
				end
			end
			return list
		end

		-- Item 1: the icons drew ~40 px over ~14 px rows on the live client.
		describe("row icons and row heights", function()
			it("sizes every row icon 16 x 16 at its row's top left, the text anchored right of it", function()
				open({ chest = WORN_CHEST, legs = THIN_LEGS })
				local S = Theme.SIZES
				assert.are.equal(16, S.rowIcon)
				local rows = { PanelTab.targetRows[1], PanelTab.targetRows[2] }
				for index = 1, 3 do
					rows[#rows + 1] = PanelTab.rotationWidgets[index]
				end
				for _, row in ipairs(rows) do
					local size = mock.lastCall(row.icon, "SetSize")
					assert.are.same({ S.rowIcon, S.rowIcon }, { size[1], size[2] })
					assert.is_true(hasPoint(row.icon, { "TOPLEFT", row.frame, "TOPLEFT", 0, 0 }))
					assert.is_true(hasPoint(row.icon, { "BOTTOMRIGHT", row.frame, "TOPLEFT", S.rowIcon, -S.rowIcon }))
					assertPoint({ "TOPLEFT", row.icon, "TOPRIGHT", S.gap, 0 }, row.name.points[#row.name.points])
				end
			end)

			it("gives every row a fixed height from its lines, the region title on its own row above", function()
				open({ chest = WORN_CHEST, legs = THIN_LEGS })
				local S = Theme.SIZES
				assert.are.equal(math.max(S.rowIcon, 2 * S.panelLine), PanelTab.targetRowHeight())
				assert.are.equal(math.max(S.rowIcon, 2 * S.panelLine), PanelTab.rotationRowHeight({ condition = "x." }))
				assert.are.equal(math.max(S.rowIcon, 3 * S.panelLine),
					PanelTab.rotationRowHeight({ condition = "x.", conditionLines = 2 }))
				assert.are.equal(S.rowIcon, PanelTab.rotationRowHeight({}))
				local y = S.panelTitle
				for index = 1, 2 do
					local row = PanelTab.targetRows[index]
					assert.are.equal(PanelTab.targetRowHeight(), row.frame:GetHeight())
					assertPoint({ "TOPLEFT", PanelTab.regionFrames.targets, "TOPLEFT", 0, -y },
						row.frame.points[#row.frame.points])
					y = y + PanelTab.targetRowHeight() + S.panelRowGap
				end
				local rotation = regionOf(PanelTab.last, "rotation")
				y = S.panelTitle
				for index, entry in ipairs(rotation.rows) do
					local widget = PanelTab.rotationWidgets[index]
					assert.are.equal(PanelTab.rotationRowHeight(entry), widget.frame:GetHeight())
					assertPoint({ "TOPLEFT", PanelTab.regionFrames.rotation, "TOPLEFT", 0, -y },
						widget.frame.points[#widget.frame.points])
					y = y + PanelTab.rotationRowHeight(entry) + S.panelRowGap
				end
				for _, kind in ipairs({ "targets", "rotation", "build" }) do
					local title = PanelTab.regionFrames[kind].title
					assertPoint({ "TOPLEFT", PanelTab.regionFrames[kind], "TOPLEFT", 0, 0 }, title.points[#title.points])
					assert.are.equal(PanelTab.regionHeight(regionOf(PanelTab.last, kind)),
						PanelTab.regionFrames[kind]:GetHeight())
				end
			end)
		end)

		-- Item 2: the flyout was ~220 px and the source line wrapped into the next row.
		describe("width and fitting", function()
			it("is flyoutWidth wide and fits every single-line text to its room by measure", function()
				local _, flyout = open({ chest = WORN_CHEST, legs = THIN_LEGS })
				local S = Theme.SIZES
				assert.are.equal(320, S.flyoutWidth)
				assert.are.equal(S.flyoutWidth, flyout:GetWidth())
				PanelTab.measure = sixPerByte
				PanelTab.refresh(DATA)
				for index = 1, 2 do
					local row = PanelTab.targetRows[index]
					assert.is_true(sixPerByte(row.name:GetText()) <= PanelTab.nameRoom(), row.name:GetText())
					assert.is_true(sixPerByte(row.detail:GetText()) <= PanelTab.textRoom(), row.detail:GetText())
				end
				-- The chest's source line no longer fits whole: the chance goes, then the place
				-- shortens; the boss is never cut (the tooltip's own rule).
				local chest = PanelTab.targetRows[2].detail:GetText()
				assert.is_truthy(chest:find("Mutanus the Devourer", 1, true))
				assert.is_nil(chest:find("33%", 1, true))
				for index = 1, 3 do
					local widget = PanelTab.rotationWidgets[index]
					assert.is_true(sixPerByte(widget.name:GetText()) <= PanelTab.textRoom())
				end
				for _, text in ipairs(PanelTab.regionFrames.build.texts) do
					assert.is_true(sixPerByte(text:GetText()) <= PanelTab.innerWidth())
				end
			end)

			it("keeps the target's source on one line, shortened the way the tooltip's source line is", function()
				start()
				local pick = { source = "The Deadmines: Captain Greenskin", kind = "D", chance = 0.3 }
				local entry = { slot = "waist", source = ItemTooltip.sourceLine(pick), pick = pick }
				assert.are.equal("Waist · The Deadmines · Captain Greenskin · 30% drop", PanelTab.detailText(entry, "Waist"))
				local room = sixPerByte("Waist · The Deadmines · Captain Greenskin")
				assert.are.equal("Waist · The Deadmines · Captain Greenskin",
					PanelTab.detailText(entry, "Waist", sixPerByte, room))
				local narrow = PanelTab.detailText(entry, "Waist", sixPerByte, room - 30)
				assert.is_truthy(narrow:find("Captain Greenskin", 1, true))
				assert.is_truthy(narrow:find(L.tooltipNamePending, 1, true))
				assert.is_true(sixPerByte(narrow) <= room - 30)
				-- Narrower than the boss alone: still one line, cut with the ellipsis.
				local tiny = PanelTab.detailText(entry, "Waist", sixPerByte, 90)
				assert.is_true(sixPerByte(tiny) <= 90)
				assert.is_truthy(tiny:find(L.tooltipNamePending, 1, true))
			end)
		end)

		-- Item 3: no scrolling; caps, and never taller than the pane.
		describe("height", function()
			local function rows(count)
				local list = {}
				for index = 1, count do
					list[index] = { name = "Row " .. index, condition = "Do it.", conditionLines = 1 }
				end
				return list
			end

			local function bosses(count)
				local list = {}
				for index = 1, count do
					list[index] = { title = "Boss " .. index, detail = "0 deaths" }
				end
				return list
			end

			it("shows at most eight rotation rows and six raid bosses, each with a muted N more", function()
				start()
				local model = PanelTab.model({ targets = {}, rotation = rows(11),
					raid = { kind = "raid", title = L.panelLastRaid, bosses = bosses(9) } })
				local rotation, raid = regionOf(model, "rotation"), regionOf(model, "raid")
				assert.are.equal(PanelTab.MAX_ROTATION, #rotation.rows)
				assert.are.equal(8, PanelTab.MAX_ROTATION)
				assert.are.equal(3, rotation.more)
				assert.are.equal(PanelTab.MAX_BOSSES, #raid.bosses)
				assert.are.equal(6, PanelTab.MAX_BOSSES)
				assert.are.equal(3, raid.more)
				assert.are.equal(11, model.rotation)
				local small = PanelTab.model({ targets = {}, rotation = rows(2) })
				assert.are.equal(0, regionOf(small, "rotation").more)
			end)

			it("fits under a height: the last raid's bosses go first, then rotation rows", function()
				start()
				local targets = {}
				for index = 1, 5 do
					targets[index] = { slot = "chest", itemId = index, name = "T" .. index, gain = 1, unit = "dps" }
				end
				local function model()
					return PanelTab.model({ targets = targets, rotation = rows(8),
						raid = { kind = "raid", title = L.panelLastRaid, bosses = bosses(6) } })
				end
				local full = PanelTab.flyoutHeight(model())
				local trimmed = PanelTab.fitHeight(model(), full - 1)
				assert.is_true(PanelTab.flyoutHeight(trimmed) <= full - 1)
				assert.are.equal(5, #regionOf(trimmed, "raid").bosses)
				assert.are.equal(1, regionOf(trimmed, "raid").more)
				assert.are.equal(8, #regionOf(trimmed, "rotation").rows)
				-- Lower than the flyout without its last raid: that region goes whole, then rotation rows.
				local noRaid = model()
				table.remove(noRaid.regions)
				local limit = PanelTab.flyoutHeight(noRaid) - 1
				local tighter = PanelTab.fitHeight(model(), limit)
				assert.is_nil(regionOf(tighter, "raid"))
				assert.are.equal(7, #regionOf(tighter, "rotation").rows)
				assert.are.equal(1, regionOf(tighter, "rotation").more)
				assert.are.equal(5, #regionOf(tighter, "targets").rows)
				assert.is_true(PanelTab.flyoutHeight(tighter) <= limit)
				-- No height to fit under (a client that cannot say): nothing trimmed.
				assert.are.equal(full, PanelTab.flyoutHeight(PanelTab.fitHeight(model(), nil)))
			end)

			it("draws no taller than the character pane, dropping the last raid's bosses first", function()
				local raid = { type = "last_raid", character = "US/Ashbringer/Bob", date = "2026-12-09", bosses = {} }
				for index = 1, 9 do
					raid.bosses[index] = { name = "Boss " .. index, percentile = 50, deaths = 0 }
				end
				local pane, flyout = open({ chest = WORN_CHEST, legs = THIN_LEGS }, { level = 60,
					inbox = { messages = { raid } } })
				local full = flyout:GetHeight()
				assert.are.equal(PanelTab.flyoutHeight(PanelTab.last), full)
				pane:SetHeight(full - 20)
				PanelTab.refresh(DATA)
				assert.is_true(flyout:GetHeight() <= full - 20)
				assert.are.equal(PanelTab.flyoutHeight(PanelTab.last), flyout:GetHeight())
				local drawn = regionOf(PanelTab.last, "raid")
				assert.is_true(#drawn.bosses < PanelTab.MAX_BOSSES)
				assert.are.equal(3, #regionOf(PanelTab.last, "rotation").rows)
				assert.are.equal(string.format(L.panelMore, 9 - #drawn.bosses), PanelTab.regionFrames.raid.more:GetText())
				assert.is_true(PanelTab.regionFrames.raid.more:IsShown())
				assert.are.same({}, state.clientWrites)
			end)
		end)

		-- Item 4: "at level N" from the next band (review D2).
		describe("the next band", function()
			it("adds the 30 band's Rapid Fire at level 26 for the level 24 beast mastery hunter (shipped data)", function()
				start()
				local data = require("Data")
				local build = { classSlug = "hunter", order = {}, gear = {} }
				local rows = PanelTab.rotationRows(data, build, { [1] = { ["1:1"] = 5 } }, 24)
				-- The band's own lines are the shipped rotation's (regenerated nightly from the
				-- curated file), so only their shape is pinned: numbered 1..n, no Rapid Fire.
				local last = rows[#rows]
				assert.is_true(#rows >= 2)
				for index = 1, #rows - 1 do
					assert.are.equal(index, rows[index].number)
					assert.are_not.equal("Rapid Fire", rows[index].name)
				end
				assert.are.equal("Rapid Fire", last.name)
				assert.is_nil(last.number)
				assert.is_true(last.muted)
				assert.are.equal(string.format(L.panelRotationAt, 26), last.at)
				assert.are.equal(3045, last.spellId)
			end)

			it("adds at most two, none already in the band, none learned", function()
				start()
				local data = { classes = DATA.classes, rotations = { ["warrior-arms"] = {
					{ level = 20, lines = { line(772, "Rend", "Rend kept up.") } },
					{ level = 30, lines = {
						line(773, "Rend", "Rend kept up."), line(845, "Cleave", "Cleave on two."),
						line(1464, "Slam", "Slam."), line(1680, "Whirlwind", "Whirlwind."), line(12294, "Mortal Strike", "MS."),
					} },
				} }, abilityLevels = { [772] = 4, [773] = 4, [845] = 20, [1464] = 30, [1680] = 36, [12294] = 40 } }
				local ranks = { [1] = { ["1:1"] = 3 } }
				local rows = PanelTab.rotationRows(data, BUILD, ranks, 21)
				local names = {}
				for index, row in ipairs(rows) do
					names[index] = row.name
				end
				assert.are.same({ "Rend", "Slam", "Whirlwind" }, names)
				assert.are.equal(string.format(L.panelRotationAt, 30), rows[2].at)
				assert.are.equal(string.format(L.panelRotationAt, 36), rows[3].at)
				-- At the top rung there is no next band: its own five lines, all learned.
				assert.are.equal(5, #PanelTab.rotationRows(data, BUILD, ranks, 60))
			end)
		end)

		-- Item 5: the button read as plain text (review D3, confirmed live).
		describe("the button", function()
			it("is 112 x 24 with the addon's 1 px gold edge over the panel fill and the chip's gold font", function()
				start({}, { tabs = 0 })
				PanelTab.refresh(DATA)
				local S = Theme.SIZES
				local button = _G[PanelTab.BUTTON_NAME]
				assert.are.same({ 112, 24 }, { button:GetWidth(), button:GetHeight() })
				assert.are.same({ S.panelTabWidth, S.panelTabHeight }, { button:GetWidth(), button:GetHeight() })
				assert.are.equal(4, #button.foreverSixtyEdges)
				assert.are.equal(1, S.border)
				for _, edge in ipairs(button.foreverSixtyEdges) do
					assert.are.same(rgb("gold"), colorOf(edge, "SetColorTexture"))
					local thickness = mock.lastCall(edge, "SetHeight") or mock.lastCall(edge, "SetWidth")
					assert.are.equal(S.border, thickness[1])
				end
				assert.are.same(rgb("background", Theme.ALPHA.window), colorOf(button.foreverSixtyGround, "SetColorTexture"))
				local label = button.foreverSixtyLabel
				assert.are.equal(L.panelTabLabel, label:GetText())
				assert.are.same(rgb("gold"), colorOf(label))
				assert.are.equal(S.chipText, mock.lastCall(label, "SetFont")[2])
				-- Under the cursor it brightens; it goes back when the cursor leaves.
				button:GetScript("OnEnter")(button)
				assert.are.same(rgb("goldHover"), colorOf(label))
				assert.are.same(rgb("hover"), colorOf(button.foreverSixtyGround, "SetColorTexture"))
				button:GetScript("OnLeave")(button)
				assert.are.same(rgb("gold"), colorOf(label))
				assert.are.same(rgb("background", Theme.ALPHA.window), colorOf(button.foreverSixtyGround, "SetColorTexture"))
				assert.are.same({}, state.clientWrites)
			end)

			it("sits right of the client's own tabs: TabSystem, then CharacterFrameTab<n>", function()
				local pane = start({}, { tabs = 0 })
				local S = Theme.SIZES
				local tab = withRect(_G.CreateFrame("Button", "CharacterFrameTab1", pane), 80)
				PanelTab.refresh(DATA)
				local button = _G[PanelTab.BUTTON_NAME]
				assertPoint({ "LEFT", tab, "RIGHT", S.panelTabGap, 0 }, button.points[#button.points])
				local system = withRect(_G.CreateFrame("Frame", nil, pane), 20)
				pane.TabSystem = system
				PanelTab.refresh(DATA)
				assertPoint({ "LEFT", system, "RIGHT", S.panelTabGap, 0 }, button.points[#button.points])
				assert.are.same({}, state.clientWrites)
			end)
		end)

		-- Item 6: frame-level specs for the rotation, last raid and build drawers (review D4).
		describe("the drawers", function()
			it("shows a rank on every rotation row that has one, learned or not", function()
				start()
				local function subtext(spellId)
					return spellId == 1 and "Rank 3" or nil
				end
				local saved = _G.GetSpellSubtext
				_G.GetSpellSubtext = subtext
				local learned = ItemTooltip.stripEscapes(PanelTab.rotationName({ number = 1, name = "Arcane Shot",
					spellId = 1 }))
				local later = ItemTooltip.stripEscapes(PanelTab.rotationName({ name = "Multi-Shot", spellId = 1,
					at = string.format(L.panelRotationAt, 20) }))
				local rankless = ItemTooltip.stripEscapes(PanelTab.rotationName({ name = "Kill Command", spellId = 2,
					at = string.format(L.panelRotationAt, 30) }))
				_G.GetSpellSubtext = saved
				assert.are.equal("1. Arcane Shot Rank 3", learned)
				assert.are.equal("Multi-Shot Rank 3 at level 20", later)
				assert.are.equal("Kill Command at level 30", rankless)
			end)

			-- The player's capture review: Multi-Shot showed no rank beside neighbours at Rank 1/3; its
			-- own subtext was empty, so the spellbook's highest learned rank stands in.
			it("derives a rank from the spellbook when the spell's own subtext is empty", function()
				start()
				local savedGet = _G.GetSpellSubtext
				_G.GetSpellSubtext = function()
					return nil
				end
				_G.GetNumSpellTabs = function()
					return 1
				end
				_G.GetSpellTabInfo = function()
					return "Class", "icon", 0, 3
				end
				_G.GetSpellBookItemName = function(index)
					return index <= 2 and "Multi-Shot" or "Kill Command", index == 1 and "Rank 1" or (index == 2 and "Rank 2" or "")
				end
				local derived = ItemTooltip.stripEscapes(PanelTab.rotationName({ name = "Multi-Shot", spellId = 9 }))
				local unranked = ItemTooltip.stripEscapes(PanelTab.rotationName({ name = "Kill Command", spellId = 10 }))
				_G.GetSpellSubtext, _G.GetNumSpellTabs, _G.GetSpellTabInfo, _G.GetSpellBookItemName = savedGet, nil, nil, nil
				assert.are.equal("Multi-Shot Rank 2", derived)
				assert.are.equal("Kill Command", unranked)
			end)

			-- Live bug: a level-1 rogue saw "Sinister Strike Rank 2": the rows come from the lowest band,
			-- whose spell ids are the Rank 2 spells, so the client's subtext said "Rank 2" whatever was learned.
			it("shows the spellbook's learned rank over the band spell id's subtext", function()
				start()
				local savedGet = _G.GetSpellSubtext
				_G.GetSpellSubtext = function()
					return "Rank 2"
				end
				_G.GetNumSpellTabs = function()
					return 1
				end
				_G.GetSpellTabInfo = function()
					return "Class", "icon", 0, 1
				end
				_G.GetSpellBookItemName = function()
					return "Sinister Strike", "Rank 1"
				end
				local learned = ItemTooltip.stripEscapes(PanelTab.rotationName({ number = 1, name = "Sinister Strike",
					spellId = 2098 }))
				local later = ItemTooltip.stripEscapes(PanelTab.rotationName({ name = "Eviscerate", spellId = 6761,
					at = string.format(L.panelRotationAt, 10) }))
				_G.GetSpellSubtext, _G.GetNumSpellTabs, _G.GetSpellTabInfo, _G.GetSpellBookItemName = savedGet, nil, nil, nil
				assert.are.equal("1. Sinister Strike Rank 1", learned)
				assert.are.equal("Eviscerate Rank 2 at level 10", later)
			end)

			it("shows no rank when the spellbook has no entry and the spell has no subtext", function()
				start()
				local savedGet = _G.GetSpellSubtext
				_G.GetSpellSubtext = function()
					return nil
				end
				local text = ItemTooltip.stripEscapes(PanelTab.rotationName({ name = "Kill Command", spellId = 10 }))
				_G.GetSpellSubtext = savedGet
				assert.are.equal("Kill Command", text)
			end)

			it("draws the rotation: number, name and rank, the condition under it, a later one muted", function()
				open({ chest = WORN_CHEST })
				local first, later = PanelTab.rotationWidgets[1], PanelTab.rotationWidgets[2]
				assert.are.equal("spell-772", mock.lastCall(first.icon, "SetTexture")[1])
				assert.are.equal("1. Rend Rank 2", ItemTooltip.stripEscapes(first.name:GetText()))
				assert.are.equal("Rend kept up.", first.condition:GetText())
				assert.are.same(rgb("body"), colorOf(first.name))
				assert.is_false(mock.lastCall(first.icon, "SetDesaturated")[1])
				-- L3.7 item 25: a row not learned yet keeps its rank beside "at level N" (Multi-Shot
				-- showed none), so every row that has one shows it.
				assert.are.equal("Execute Rank 2 " .. string.format(L.panelRotationAt, 24),
					ItemTooltip.stripEscapes(later.name:GetText()))
				assert.are.same(rgb("muted"), colorOf(later.name))
				assert.is_true(mock.lastCall(later.icon, "SetDesaturated")[1])
				assert.is_false(PanelTab.regionFrames.rotation.more:IsShown())
				assert.are.same({}, state.clientWrites)
			end)

			it("draws the last raid: the dated title, each boss with its muted detail, N more past six", function()
				local raid = { type = "last_raid", character = "US/Ashbringer/Bob", date = "2026-12-09", bosses = {
					{ name = "Lucifron", percentile = 71, deaths = 0, consumables_missed = { "flask" } },
				} }
				for index = 2, 8 do
					raid.bosses[index] = { name = "Boss " .. index }
				end
				open({ chest = WORN_CHEST }, { level = 60, inbox = { messages = { raid } } })
				local region = PanelTab.regionFrames.raid
				assert.is_true(region:IsShown())
				assert.are.equal(string.format(L.panelLastRaidAsOf, "9 Dec"), region.title:GetText())
				assert.are.same({
					string.format(L.panelRaidBoss, "Lucifron", "71st"),
					string.format(L.panelRaidDeaths, 0) .. L.panelRaidJoin .. string.format(L.panelRaidMissed, "flask"),
					"Boss 2", "Boss 3", "Boss 4", "Boss 5", "Boss 6",
				}, shownTexts(region))
				assert.are.same(rgb("muted"), colorOf(region.texts[2]))
				assert.is_true(region.more:IsShown())
				assert.are.equal(string.format(L.panelMore, 2), region.more:GetText())
				assert.are.same(rgb("muted"), colorOf(region.more))
				assert.are.equal(PanelTab.regionHeight(regionOf(PanelTab.last, "raid")), region:GetHeight())
				assert.are.same({}, state.clientWrites)
			end)

			it("draws the build: split, progress, off plan in the warning colour, then the two buttons under them", function()
				open({ chest = WORN_CHEST })
				local S = Theme.SIZES
				local region = PanelTab.regionFrames.build
				local entry = regionOf(PanelTab.last, "build")
				assert.are.same({ entry.split, entry.progress, entry.offPlan }, shownTexts(region))
				assert.are.same(rgb("body"), colorOf(region.texts[1]))
				assert.are.same(rgb("muted"), colorOf(region.texts[2]))
				assert.are.same(rgb("warning"), colorOf(region.texts[3]))
				local y = S.panelTitle + 3 * S.panelLine + S.panelRowGap
				assertPoint({ "TOPLEFT", region, "TOPLEFT", 0, -y }, region.useOther.points[#region.useOther.points])
				assertPoint({ "TOPRIGHT", region, "TOPRIGHT", 0, -y }, region.copy.points[#region.copy.points])
				assert.are.equal(PanelTab.regionHeight(entry), region:GetHeight())
				assert.are.equal(y + S.panelButtonHeight, region:GetHeight())
				assert.are.same({}, state.clientWrites)
			end)
		end)

		-- Item 7: no draw while hidden (review D5), and the conditions measured once.
		describe("drawing only what shows", function()
			it("skips the draw while the pane is closed", function()
				local pane, flyout = open({ chest = WORN_CHEST })
				local draws = PanelTab.draws
				pane:Hide()
				PanelTab.refresh(DATA)
				assert.are.equal(draws, PanelTab.draws)
				-- Fix round 2 (re-review b): it stays open, marked stale, for the button's OnShow.
				assert.is_true(flyout:IsShown())
				assert.is_false(flyout:IsVisible())
				assert.are.same({}, state.clientWrites)
			end)

			it("measures a condition once for each width", function()
				start()
				local calls = 0
				local function measure(text)
					calls = calls + 1
					return #text * 6
				end
				local ranks = { [1] = { ["1:1"] = 3 } }
				PanelTab.rotationRows(DATA, BUILD, ranks, 21, measure, 120)
				local first = calls
				assert.is_true(first > 0)
				local again = PanelTab.rotationRows(DATA, BUILD, ranks, 21, measure, 120)
				assert.are.equal(first, calls)
				assert.are.equal("Rend kept up.", again[1].condition)
				PanelTab.rotationRows(DATA, BUILD, ranks, 21, measure, 100)
				assert.is_true(calls > first)
			end)
		end)

		-- Item 8: a slot with the bag arrow (review D6).
		describe("the bag arrow", function()
			it("lists the slot's bag item as its hover leads with it, never the next pick", function()
				local BAG_CHEST = link(250, "Bag Chest")
				start({ chest = WORN_CHEST, legs = THIN_LEGS })
				ItemTooltip.bagUpgradeFor = function(slot)
					if slot == "chest" then
						return { link = BAG_CHEST, itemId = 250, delta = 2.0, unit = "dps", slot = "chest" }
					end
					return nil
				end
				local chest
				for _, row in ipairs(PanelTab.inputs(DATA).targets) do
					if row.slot == "chest" then
						chest = row
					end
				end
				assert.is_true(chest.bag)
				assert.are.equal(250, chest.itemId)
				assert.are.equal("Bag Chest", chest.name)
				assert.are.equal(L.tooltipRightClickEquip, chest.source)
				local lead = ItemTooltip.bagLead("chest", {})
				assert.are.equal(lead.lines[2].upgrade.gain, chest.gain)
				assert.are.equal(lead.target.itemId, chest.itemId)
				PanelTab.refresh(DATA)
				click(_G[PanelTab.BUTTON_NAME])
				local drawn
				for index = 1, 2 do
					local row = PanelTab.targetRows[index]
					if row.itemId == 250 then
						drawn = row
					end
				end
				assert.are.equal(L.panelInYourBags .. " Bag Chest", ItemTooltip.stripEscapes(drawn.name:GetText()))
				assert.are.equal(string.format(L.panelTargetDetail, L.gearSlotLabels.chest, L.tooltipRightClickEquip),
					drawn.detail:GetText())
				assert.are.equal(ItemTooltip.figureText(chest.gain, chest.unit), drawn.gain:GetText())
				assert.are.same({}, state.clientWrites)
			end)
		end)
	end)

	-- Fix round 2 (the 1.60 client's probes on fix round 1, and the fix-1 re-review).
	describe("fix round 2", function()
		--- Whether the button's last point is the pane's bottom-right fallback.
		local function atFallback(pane, button)
			local point = button.points[#button.points]
			return point ~= nil and point[1] == "TOPRIGHT" and point[2] == pane and point[3] == "BOTTOMRIGHT"
		end

		-- Item 2: PaperDollSidebarTab4 is not visible and has no rect; the button anchored to it had none.
		it("never anchors the button to the paper doll's sidebar tabs or their holder", function()
			local pane = start({}, { tabs = 0 })
			for index = 1, 4 do
				withRect(_G.CreateFrame("Button", "PaperDollSidebarTab" .. index, pane), 400 + 30 * index)
			end
			withRect(_G.CreateFrame("Frame", "PaperDollSidebarTabs", pane), 400)
			PanelTab.refresh(DATA)
			assert.is_true(atFallback(pane, _G[PanelTab.BUTTON_NAME]))
			assert.are.same({}, state.clientWrites)
		end)

		it("takes a tab only when it shows and has a rect, else falls through to the fallback", function()
			local pane = start({}, { tabs = 0 })
			local S = Theme.SIZES
			-- Shown, no rect: GetLeft answers nil.
			local tab = _G.CreateFrame("Button", "CharacterFrameTab1", pane)
			local system = _G.CreateFrame("Frame", nil, pane)
			pane.TabSystem = system
			PanelTab.refresh(DATA)
			local button = _G[PanelTab.BUTTON_NAME]
			assert.is_true(atFallback(pane, button))
			-- A rect, but hidden.
			withRect(tab, 80):Hide()
			withRect(system, 20):Hide()
			PanelTab.refresh(DATA)
			assert.is_true(atFallback(pane, button))
			-- Shown with a rect: taken.
			tab:Show()
			PanelTab.refresh(DATA)
			assertPoint({ "LEFT", tab, "RIGHT", S.panelTabGap, 0 }, button.points[#button.points])
			-- A GetLeft that is not a number, or that fails, is no rect.
			tab.GetLeft = function()
				return "647"
			end
			PanelTab.refresh(DATA)
			assert.is_true(atFallback(pane, button))
			tab.GetLeft = function()
				error("secret")
			end
			PanelTab.refresh(DATA)
			assert.is_true(atFallback(pane, button))
			assert.are.same({}, state.clientWrites)
		end)

		-- Item 1: the flyout covered CharacterFrameModeTabs, a 64 px column at the pane's right.
		it("watches CharacterFrameModeTabs as a client frame", function()
			assert.is_true(mock.isClientName("CharacterFrameModeTabs"))
			local found = false
			for _, name in ipairs(mock.CLIENT_NAMES) do
				found = found or name == "CharacterFrameModeTabs"
			end
			assert.is_true(found)
		end)

		it("opens the flyout right of the pane's mode tabs when they show with a rect", function()
			local pane = start({ chest = WORN_CHEST })
			local S = Theme.SIZES
			local modeTabs = withRect(_G.CreateFrame("Frame", "CharacterFrameModeTabs", pane), 647)
			PanelTab.refresh(DATA)
			click(_G[PanelTab.BUTTON_NAME])
			local flyout = _G[PanelTab.FLYOUT_NAME]
			assert.is_true(flyout:IsShown())
			assertPoint({ "TOPLEFT", modeTabs, "TOPRIGHT", S.panelOffsetX, 0 }, flyout.points[#flyout.points])
			assert.are.same({}, state.clientWrites)
		end)

		it("opens the flyout at the pane's top right when the mode tabs are absent, hidden or rect-less", function()
			local pane = start({ chest = WORN_CHEST })
			local S = Theme.SIZES
			local paneAnchor = { "TOPLEFT", pane, "TOPRIGHT", S.panelOffsetX, S.panelOffsetY }
			PanelTab.refresh(DATA)
			local button = _G[PanelTab.BUTTON_NAME]
			click(button)
			local flyout = _G[PanelTab.FLYOUT_NAME]
			assertPoint(paneAnchor, flyout.points[#flyout.points])
			click(button)
			-- They appear later, with no rect yet: still the pane.
			local modeTabs = _G.CreateFrame("Frame", "CharacterFrameModeTabs", pane)
			click(button)
			assertPoint(paneAnchor, flyout.points[#flyout.points])
			click(button)
			-- A rect but hidden: still the pane.
			withRect(modeTabs, 647):Hide()
			click(button)
			assertPoint(paneAnchor, flyout.points[#flyout.points])
			click(button)
			-- Shown with a rect: resolved again on this open.
			modeTabs:Show()
			click(button)
			assertPoint({ "TOPLEFT", modeTabs, "TOPRIGHT", S.panelOffsetX, 0 }, flyout.points[#flyout.points])
			assert.are.same({}, state.clientWrites)
		end)

		-- Item 3: the chat box cuts a /run line at about 200 characters.
		it("keeps every /run and /dump probe in README rows 45-47 under 180 characters", function()
			local file = assert(io.open("README.md", "r"))
			local count = 0
			for row in file:lines() do
				local number = tonumber(row:match("^| (%d+) |"))
				if number ~= nil and number >= 45 and number <= 47 then
					for probe in row:gmatch("`(/%a+ [^`]*)`") do
						count = count + 1
						assert.is_true(#probe < 180, #probe .. ": " .. probe)
					end
				end
			end
			file:close()
			assert.is_true(count >= 8)
		end)

		-- Re-review (a): the line box follows the client's measured small font, at least panelLine.
		it("measures a panel line from the small font, never under panelLine", function()
			start()
			local S = Theme.SIZES
			local function probe(height)
				return { SetText = function() end, GetStringHeight = function()
					return height
				end }
			end
			local measure = function()
				return 10
			end
			assert.are.equal(S.panelLine, PanelTab.measuredLine(probe(9.5), measure))
			assert.are.equal(S.panelLine, PanelTab.measuredLine(probe(nil), measure))
			assert.are.equal(S.panelLine, PanelTab.measuredLine(probe(12), measure))
			assert.are.equal(15, PanelTab.measuredLine(probe(14.2), measure))
			assert.are.equal(S.panelLine, PanelTab.measuredLine(nil, measure))
			-- It reads the height of the line the measurer just set.
			local set
			local region = { SetText = function(_, text)
				set = text
			end, GetStringHeight = function()
				return set ~= nil and 13 or nil
			end }
			assert.are.equal(13, PanelTab.measuredLine(region, Theme.measurer(region)))
			assert.are.equal(PanelTab.LINE_SAMPLE, set)
		end)

		it("sizes every line box and row from the measured line", function()
			start({ chest = WORN_CHEST, legs = THIN_LEGS })
			local Widgets = require("Widgets")
			local label = Widgets.label
			Widgets.label = function(...)
				local region = label(...)
				region.GetStringHeight = function()
					return 14.2
				end
				return region
			end
			PanelTab.refresh(DATA)
			click(_G[PanelTab.BUTTON_NAME])
			Widgets.label = label
			local S = Theme.SIZES
			assert.are.equal(15, PanelTab.lineHeight())
			assert.are.equal(math.max(S.rowIcon, 30), PanelTab.targetRowHeight())
			assert.are.equal(45, PanelTab.rotationRowHeight({ condition = "x.", conditionLines = 2 }))
			local row = PanelTab.targetRows[1]
			assert.are.equal(15, mock.lastCall(row.name, "SetSize")[2])
			assert.are.equal(15, mock.lastCall(row.second, "SetSize")[2])
			assert.are.equal(30, row.frame:GetHeight())
			local flyout = _G[PanelTab.FLYOUT_NAME]
			assert.are.equal(PanelTab.flyoutHeight(PanelTab.last), flyout:GetHeight())
		end)

		-- Re-review (b): an event while the pane is closed marks the open flyout stale; the
		-- button's own OnShow (it shows with the pane) redraws it once.
		it("leaves the flyout open but stale on a refresh while the pane is closed, redrawn once on show", function()
			local pane = start({ chest = WORN_CHEST })
			PanelTab.refresh(DATA)
			local button = _G[PanelTab.BUTTON_NAME]
			click(button)
			local flyout = _G[PanelTab.FLYOUT_NAME]
			local draws = PanelTab.draws
			pane:Hide()
			PanelTab.refresh(DATA)
			assert.are.equal(draws, PanelTab.draws)
			assert.is_true(flyout:IsShown())
			assert.is_true(PanelTab.stale)
			pane:Show()
			local onShow = button:GetScript("OnShow")
			assert.is_function(onShow)
			onShow(button)
			assert.are.equal(draws + 1, PanelTab.draws)
			assert.is_falsy(PanelTab.stale)
			assert.is_true(flyout:IsVisible())
			onShow(button)
			assert.are.equal(draws + 1, PanelTab.draws)
			-- A closed flyout is never drawn on show, stale or not.
			click(button)
			pane:Hide()
			PanelTab.refresh(DATA)
			pane:Show()
			onShow(button)
			assert.are.equal(draws + 1, PanelTab.draws)
			assert.are.equal(0, mock.countCalls(pane, "SetScript"))
			assert.are.equal(0, mock.countCalls(flyout, "SetScript"))
			assert.are.same({}, state.clientWrites)
		end)

		it("places the button again when the pane shows, once a tab has its rect", function()
			local pane = start({}, { tabs = 0 })
			local S = Theme.SIZES
			local tab = _G.CreateFrame("Button", "CharacterFrameTab1", pane)
			PanelTab.refresh(DATA)
			local button = _G[PanelTab.BUTTON_NAME]
			assert.is_true(atFallback(pane, button))
			withRect(tab, 80)
			button:GetScript("OnShow")(button)
			assertPoint({ "LEFT", tab, "RIGHT", S.panelTabGap, 0 }, button.points[#button.points])
			assert.are.same({}, state.clientWrites)
		end)
	end)

	-- Fix round 3 (the designer and player personas on the hunter-24 captures).
	describe("fix round 3", function()
		local function open(worn, options)
			local pane = start(worn, options)
			PanelTab.refresh(DATA)
			click(_G[PanelTab.BUTTON_NAME])
			return pane, _G[PanelTab.FLYOUT_NAME]
		end

		local function rgb(key, alpha)
			return { Theme.rgb(Theme.HEX[key], alpha) }
		end

		local function colorOf(region, method)
			local call = mock.lastCall(region, method or "SetTextColor")
			return { call[1], call[2], call[3], call[4] }
		end

		--- GameTooltip's lines since its last ClearLines: the texts, and whether each wraps.
		local function tooltip()
			local texts, wraps = {}, {}
			for _, call in ipairs(_G.GameTooltip.calls) do
				if call.method == "ClearLines" then
					texts, wraps = {}, {}
				elseif call.method == "AddLine" then
					texts[#texts + 1] = call[1]
					wraps[#wraps + 1] = call[5] == true
				end
			end
			return texts, wraps
		end

		local function regionOf(model, kind)
			for _, region in ipairs(model.regions) do
				if region.kind == kind then
					return region
				end
			end
			return nil
		end

		--- The level 24 hunter on the shipped data, read as beast mastery.
		local HUNTER = { classSlug = "hunter", order = {}, gear = {} }
		local HUNTER_RANKS = { [1] = { ["1:1"] = 5 } }

		--- A measure calibrated on the hunter-24 capture (4.6 units a character): at the 280 unit
		--- text room the Arcane Shot condition wraps after "Aimed Shot will" and after "a 25%", as drawn.
		local function hunterWidth(text)
			local visible = ItemTooltip.stripEscapes(text):gsub("[\128-\191]", "")
			return #visible * 4.6
		end

		-- Item 1: the build region reads as the plan, measured against the level.
		describe("the build region", function()
			--- A 51-point plan: Arms A thirty times, Fury C twenty-one times.
			local PLAN = { classSlug = "warrior", gear = {}, order = {} }
			for index = 1, 30 do
				PLAN.order[index] = { tab = 1, tier = 1, column = 1 }
			end
			for index = 31, 51 do
				PLAN.order[index] = { tab = 2, tier = 1, column = 1 }
			end
			local RANKS = { [1] = { ["1:1"] = 10, ["1:2"] = 1 }, [2] = { ["1:1"] = 2 } }

			it("prefixes the split with Plan:", function()
				start()
				local region = PanelTab.buildRegion(PLAN, DATA, RANKS, 24)
				assert.are.equal(string.format(L.panelPlanSplit, string.format(L.dialogSplitPart, "Arms", 30)
					.. L.dialogSplitSeparator .. string.format(L.dialogSplitPart, "Fury", 21)), region.split)
				assert.are.equal("Plan: Arms 30 / Fury 21", region.split)
			end)

			it("counts the points on plan against the points the level allows: 15 at 24", function()
				start()
				assert.are.equal("12 of 15 points on plan", PanelTab.buildRegion(PLAN, DATA, RANKS, 24).progress)
				assert.are.equal(15, PanelTab.pointsAllowed(PLAN, 24))
			end)

			-- L3.7 item 27: below level 10 no point is allowed yet, and "0 of 0 points on plan" is no
			-- line at all.
			it("names the next planned point and its level for the off-plan hover", function()
				start()
				local region = PanelTab.buildRegion(BUILD, DATA, { [1] = { ["1:1"] = 2, ["1:2"] = 1 } }, 24)
				assert.is_not_nil(region.nextPoint)
				assert.are.equal("C", region.nextPoint.name)
				assert.is_number(region.nextPoint.level)
				assert.are.same({ Theme.inlineColor("warning", region.offPlan),
					Theme.inlineColor("tooltipText", region.offPlanTalents[1]),
					Theme.inlineColor("tooltipMuted", L.panelOffPlanAction),
					Theme.inlineColor("tooltipText", string.format(L.panelOffPlanNext, region.nextPoint.level, "C")) },
					PanelTab.offPlanHint(region))
				-- Nothing to take next: the action line alone.
				local done = PanelTab.buildRegion(BUILD, DATA, { [1] = { ["1:1"] = 3, ["1:2"] = 1 }, [2] = { ["1:1"] = 1 } }, 24)
				assert.is_nil(done.nextPoint)
			end)

			it("has no progress line below level 10, where no point is allowed", function()
				start()
				local region = PanelTab.buildRegion(PLAN, DATA, RANKS, 8)
				assert.is_nil(region.progress)
				assert.are.equal(string.format(L.panelPlanSplit, "Arms 30 / Fury 21"), region.split)
				-- One line shorter than the same plan at 24 (both carry the off-plan line).
				assert.are.equal(PanelTab.regionHeight(PanelTab.buildRegion(PLAN, DATA, RANKS, 24)) - PanelTab.lineHeight(),
					PanelTab.regionHeight(region))
			end)

			-- L3.7 item 21: the panel counts the talent window's staged clicks, so "N of M on plan"
			-- follows what the player is about to commit, with a muted "(staged)" while they differ.
			it("counts the staged clicks and says (staged) while they differ from the committed", function()
				start()
				local committed = PanelTab.buildRegion(PLAN, DATA, RANKS, 24, false)
				local staged = PanelTab.buildRegion(PLAN, DATA, { [1] = { ["1:1"] = 13 }, [2] = { ["1:1"] = 2 } }, 24, true)
				assert.are.equal("12 of 15 points on plan", committed.progress)
				assert.is_nil(committed.staged)
				assert.are.equal("15 of 15 points on plan (staged)", staged.progress)
				assert.are.equal(L.panelStaged, "(staged)")
				assert.is_true(staged.staged)
			end)

			it("reads the staged ranks for the build region through the chip's own reader", function()
				start()
				local Talents = require("Talents")
				local real = Talents.readStagedRanks
				Talents.readStagedRanks = function()
					return { [1] = { ["1:1"] = 5, ["1:2"] = 1 }, [2] = { ["1:1"] = 1 } }, true
				end
				local region = PanelTab.inputs(DATA).build
				Talents.readStagedRanks = real
				assert.is_truthy(region.progress:find(L.panelStaged, 1, true))
				assert.is_true(region.staged)
				-- Committed otherwise: the ranks the talents mock holds, nothing staged.
				assert.is_nil(PanelTab.inputs(DATA).build.staged)
			end)

			-- L3.7 pre-item (round A review): a client that raises reading the staged ranks falls back
			-- to the committed ones, as the build dialog does.
			it("falls back to the committed ranks when reading the staged ones raises", function()
				start()
				local Talents = require("Talents")
				local real = Talents.readStagedRanks
				Talents.readStagedRanks = function()
					error("no staged read")
				end
				local ok, region = pcall(function()
					return PanelTab.inputs(DATA).build
				end)
				Talents.readStagedRanks = real
				assert.is_true(ok, tostring(region))
				assert.is_nil(region.staged)
			end)

			it("caps the points allowed at the plan's own total: 51 at 60", function()
				start()
				assert.are.equal("12 of 51 points on plan", PanelTab.buildRegion(PLAN, DATA, RANKS, 60).progress)
				assert.are.equal(51, PanelTab.pointsAllowed(PLAN, 60))
				-- A short plan is its own cap below 60 too; a level the client cannot say reads the plan.
				assert.are.equal(4, PanelTab.pointsAllowed(BUILD, 24))
				assert.are.equal(51, PanelTab.pointsAllowed(PLAN, nil))
				assert.are.equal(0, PanelTab.pointsAllowed(PLAN, 8))
			end)

			it("names each off-plan talent with its tree and rank, one per line", function()
				start()
				-- B is a node the plan never touches (1 of 1); C is planned once and holds 3.
				local ranks = { [1] = { ["1:1"] = 2, ["1:2"] = 1 }, [2] = { ["1:1"] = 3 } }
				local region = PanelTab.buildRegion(BUILD, DATA, ranks, 24)
				assert.are.equal(string.format(L.panelOffPlan, 3), region.offPlan)
				assert.are.same({
					string.format(L.panelOffPlanTalent, "B", "Arms", 1, 1, 1),
					string.format(L.panelOverPlanTalent, "C", "Fury", 2, 3, 5, 1),
				}, region.offPlanTalents)
				-- L3.7 item 27: the numbers a line leads with sum to the count (1 + 2 = 3 off plan): the
				-- over-plan line counts the points over the plan, not the rank held.
				assert.are.same({ "B · Arms · 1 off plan (Rank 1/1)", "C · Fury · 2 off plan (Rank 3/5, plan 1)" },
					region.offPlanTalents)
				assert.is_nil(PanelTab.buildRegion(BUILD, DATA, { [1] = { ["1:1"] = 3 }, [2] = { ["1:1"] = 1 } },
					24).offPlanTalents)
			end)

			it("keeps the off-plan count red and names the talents on its hover", function()
				open({ chest = WORN_CHEST })
				local region = PanelTab.regionFrames.build
				local entry = regionOf(PanelTab.last, "build")
				assert.are.equal(entry.offPlan, region.texts[3]:GetText())
				assert.are.same(rgb("warning"), colorOf(region.texts[3]))
				local hover = region.offPlanHover
				assert.is_true(hover:IsShown())
				assertPoint({ "TOPLEFT", region.texts[3], "TOPLEFT", 0, 0 }, hover.points[#hover.points])
				hover:GetScript("OnEnter")(hover)
				local lines = tooltip()
				assert.are.equal(Theme.inlineColor("warning", entry.offPlan), lines[1])
				for index, text in ipairs(entry.offPlanTalents) do
					assert.are.equal(Theme.inlineColor("tooltipText", text), lines[index + 1])
				end
				-- L3.7 item 23: what to do about it, then when the next planned point arrives.
				local after = 1 + #entry.offPlanTalents
				assert.are.equal(Theme.inlineColor("tooltipMuted", L.panelOffPlanAction), lines[after + 1])
				assert.are.equal("Not in your plan. Respec to fix this.", L.panelOffPlanAction)
				if entry.nextPoint ~= nil then
					assert.are.equal(Theme.inlineColor("tooltipText", string.format(L.panelOffPlanNext,
						entry.nextPoint.level, entry.nextPoint.name)), lines[after + 2])
					assert.are.equal(after + 2, #lines)
				else
					assert.are.equal(after + 1, #lines)
				end
				-- Beside the flyout, never over the Build lines the cursor is on.
				local owner = mock.lastCall(_G.GameTooltip, "SetOwner")
				assert.are.equal(PanelTab.flyout, owner[1])
				assert.are.equal("ANCHOR_RIGHT", owner[2])
				hover:GetScript("OnLeave")(hover)
				assert.is_false(_G.GameTooltip:IsShown())
				-- None off plan: no line and no hover.
				Follow.build = { classSlug = "warrior", gear = {}, order = {
					{ tab = 1, tier = 1, column = 1 }, { tab = 1, tier = 1, column = 1 }, { tab = 1, tier = 1, column = 2 },
					{ tab = 2, tier = 1, column = 1 },
				} }
				PanelTab.refresh(DATA)
				assert.is_nil(regionOf(PanelTab.last, "build").offPlan)
				assert.is_false(hover:IsShown())
				assert.are.same({}, state.clientWrites)
			end)
		end)

		-- Item 2: the rotation conditions never lose the rule.
		describe("the rotation rows", function()
			it("cuts the hunter's Arcane Shot at a clause on the real width, never mid-figure", function()
				start()
				local data = require("Data")
				local Rotation = require("Rotation")
				local rows = PanelTab.rotationRows(data, HUNTER, HUNTER_RANKS, 24, hunterWidth, PanelTab.textRoom())
				local arcane
				for _, row in ipairs(rows) do
					if row.name == "Arcane Shot" then
						arcane = row
					end
				end
				assert.are.equal(280, PanelTab.textRoom())
				assert.are.equal(2, arcane.conditionLines)
				-- The cut keeps a leading clause of the full condition and ends on the
				-- pending mark; the shipped wording is the curated note's, so it is not pinned.
				local pending = L.tooltipNamePending
				assert.are.equal(pending, arcane.condition:sub(-#pending))
				local kept = arcane.condition:sub(1, -#pending - 1)
				assert.is_true(#kept > 0)
				assert.are.equal(kept, arcane.fullCondition:sub(1, #kept))
				assert.is_truthy(arcane.fullCondition:sub(#kept + 1):match("^[%s,;:]"))
				assert.is_nil(arcane.condition:find("25%", 1, true))
				-- The full plain condition is kept for the hover.
				local raw
				for _, entry in ipairs(Rotation.model(data, HUNTER, HUNTER_RANKS, 24, true).lines) do
					if entry.name == "Arcane Shot" then
						raw = entry.condition
					end
				end
				assert.are.equal(Rotation.plainCondition(raw), arcane.fullCondition)
				assert.is_truthy(arcane.fullCondition:find("mana is above a 25% reserve", 1, true))
			end)

			it("shows the ability name and the full plain condition on a row's hover, over the hover ground", function()
				open({ chest = WORN_CHEST })
				local widget = PanelTab.rotationWidgets[1]
				assert.is_false(widget.ground:IsShown())
				assert.are.same(rgb("hover"), colorOf(widget.ground, "SetColorTexture"))
				widget.frame:GetScript("OnEnter")(widget.frame)
				assert.is_true(widget.ground:IsShown())
				local lines, wraps = tooltip()
				assert.are.same({ "Rend", Theme.inlineColor("tooltipText", "Rend kept up.") }, lines)
				assert.are.same({ false, true }, wraps)
				widget.frame:GetScript("OnLeave")(widget.frame)
				assert.is_false(widget.ground:IsShown())
				assert.is_false(_G.GameTooltip:IsShown())
				-- A row learned later adds "Learned at level N".
				local later = PanelTab.rotationWidgets[2]
				later.frame:GetScript("OnEnter")(later.frame)
				assert.are.same({ "Execute", Theme.inlineColor("tooltipText", "Execute below 20% health."),
					Theme.inlineColor("tooltipMuted", string.format(L.panelRotationLearnedAt, 24)) }, (tooltip()))
				assert.are.same({}, state.clientWrites)
			end)

			it("names the next band's row and the level it is learned at on its hover (shipped data)", function()
				start()
				local data = require("Data")
				local rows = PanelTab.rotationRows(data, HUNTER, HUNTER_RANKS, 24)
				local rapid = rows[#rows]
				assert.are.equal("Rapid Fire", rapid.name)
				assert.are.equal(26, rapid.atLevel)
				assert.are.same({ "Rapid Fire", Theme.inlineColor("tooltipText", rapid.fullCondition),
					Theme.inlineColor("tooltipMuted", "Learned at level 26") }, PanelTab.rotationHint(rapid))
			end)
		end)

		-- Item 3: the button's states and hint.
		describe("the button", function()
			it("reads closed, lit under the cursor, and open while the flyout shows", function()
				start({ chest = WORN_CHEST })
				PanelTab.refresh(DATA)
				local button = _G[PanelTab.BUTTON_NAME]
				local label, ground = button.foreverSixtyLabel, button.foreverSixtyGround
				-- Closed.
				assert.are.same(rgb("gold"), colorOf(label))
				assert.are.same(rgb("background", Theme.ALPHA.window), colorOf(ground, "SetColorTexture"))
				-- Hover.
				button:GetScript("OnEnter")(button)
				assert.are.same(rgb("goldHover"), colorOf(label))
				assert.are.same(rgb("hover"), colorOf(ground, "SetColorTexture"))
				button:GetScript("OnLeave")(button)
				-- Open: the label in body over the inset ground of a pressed tab.
				click(button)
				assert.are.equal("inset", Theme.CONTROL_STATES.tab.pressed)
				assert.are.same(rgb("body"), colorOf(label))
				assert.are.same(rgb(Theme.CONTROL_STATES.tab.pressed), colorOf(ground, "SetColorTexture"))
				-- Open under the cursor: the brighter label on the inset ground; open again on leave.
				button:GetScript("OnEnter")(button)
				assert.are.same(rgb("goldHover"), colorOf(label))
				assert.are.same(rgb("inset"), colorOf(ground, "SetColorTexture"))
				button:GetScript("OnLeave")(button)
				assert.are.same(rgb("body"), colorOf(label))
				-- Closed by its own click, and by the flyout's x.
				click(button)
				assert.are.same(rgb("gold"), colorOf(label))
				assert.are.same(rgb("background", Theme.ALPHA.window), colorOf(ground, "SetColorTexture"))
				click(button)
				assert.are.same(rgb("body"), colorOf(label))
				click(PanelTab.closeButton)
				assert.is_false(_G[PanelTab.FLYOUT_NAME]:IsShown())
				assert.are.same(rgb("gold"), colorOf(label))
				assert.are.same({}, state.clientWrites)
			end)

			it("says what it is and what a click does on its hover", function()
				start()
				PanelTab.refresh(DATA)
				local button = _G[PanelTab.BUTTON_NAME]
				button:GetScript("OnEnter")(button)
				assert.are.same({
					Theme.inlineColor("gold", "Forever Sixty"),
					Theme.inlineColor("tooltipMuted", "Your next targets, rotation and build."),
					Theme.inlineColor("tooltipMuted", "Click to open or close."),
				}, (tooltip()))
				assert.is_true(_G.GameTooltip:IsShown())
				button:GetScript("OnLeave")(button)
				assert.is_false(_G.GameTooltip:IsShown())
				assert.are.same({}, state.clientWrites)
			end)
		end)

		-- Item 4: row affordance and context.
		describe("row affordance", function()
			it("lights a target row's hover ground and still shows the item's own tooltip", function()
				open({ chest = WORN_CHEST })
				local row = PanelTab.targetRows[1]
				assert.is_false(row.ground:IsShown())
				assert.are.same(rgb("hover"), colorOf(row.ground, "SetColorTexture"))
				row.frame:GetScript("OnEnter")(row.frame)
				assert.is_true(row.ground:IsShown())
				assert.is_truthy(mock.lastCall(_G.GameTooltip, "SetHyperlink") or mock.lastCall(_G.GameTooltip, "SetItemByID"))
				row.frame:GetScript("OnLeave")(row.frame)
				assert.is_false(row.ground:IsShown())
			end)

			-- L3.7 item 22: the name line's empty gap carries "· over <worn item>" in the muted
			-- token; short of room its worn name is cut with an ellipsis first, and the target's
			-- own name is never cut for it.
			it("puts a muted over <worn item> in the name line's gap, its worn name ellipsised first", function()
				start()
				local entry = { name = "Cord of the Wild", hex = "1eff00", worn = "Short Ash Bow", slot = "waist" }
				local function sixPerByte(text)
					return #ItemTooltip.stripEscapes(text) * 6
				end
				local name = "|cff1eff00Cord of the Wild|r"
				local over = "|cff" .. Theme.HEX.muted .. string.format(L.panelTargetOver, "Short Ash Bow") .. "|r"
				assert.are.equal(" · over Short Ash Bow", string.format(L.panelTargetOver, "Short Ash Bow"))
				assert.are.equal(name .. over, PanelTab.nameLine(entry))
				local room = sixPerByte(name .. over)
				assert.are.equal(name .. over, PanelTab.nameLine(entry, sixPerByte, room))
				-- One character short: the worn name takes an ellipsis, the target's name stays whole.
				local cut = PanelTab.nameLine(entry, sixPerByte, room - 6)
				assert.is_truthy(cut:find("Cord of the Wild", 1, true))
				assert.is_truthy(cut:find(L.tooltipNamePending, 1, true))
				assert.is_truthy(cut:find("Short Ash", 1, true))
				assert.is_true(sixPerByte(cut) <= room - 6)
				-- Not enough room for even a few letters of it: the name alone.
				assert.are.equal(name, PanelTab.nameLine(entry, sixPerByte, sixPerByte(name) + 6 * 14))
				-- A bag row (its own lead) names no worn item.
				assert.is_nil(PanelTab.nameLine({ name = "Bag Chest", worn = "X", bag = true }):find("over", 1, true))
			end)

			it("leaves the source line to the tooltip's own words, with no over", function()
				start()
				local pick = { source = "The Deadmines: Captain Greenskin", kind = "D", chance = 0.3 }
				local entry = { slot = "waist", source = ItemTooltip.sourceLine(pick), pick = pick, worn = "Short Ash Bow" }
				assert.are.equal("Waist · The Deadmines · Captain Greenskin · 30% drop", PanelTab.detailText(entry, "Waist"))
			end)

			it("lights a rotation row's hover ground the same way", function()
				open({ chest = WORN_CHEST })
				local widget = PanelTab.rotationWidgets[1]
				assert.are.equal(widget.frame, widget.ground.parent)
				widget.frame:GetScript("OnEnter")(widget.frame)
				assert.is_true(widget.ground:IsShown())
				assert.are.same(rgb("hover"), colorOf(widget.ground, "SetColorTexture"))
				widget.frame:GetScript("OnLeave")(widget.frame)
				assert.is_false(widget.ground:IsShown())
			end)

			it("gives the close x a hit rectangle of at least 20 x 20", function()
				open({ chest = WORN_CHEST })
				local close = PanelTab.closeButton
				assert.are.equal(_G[PanelTab.FLYOUT_NAME], close.parent)
				assert.is_true(close:GetWidth() >= 20)
				assert.is_true(close:GetHeight() >= 20)
			end)
		end)

		-- Item 5: hierarchy and the unlearned row.
		describe("hierarchy", function()
			it("draws a 1 px border rule above the Rotation and Build titles, inside the 10 px gap", function()
				open({ chest = WORN_CHEST })
				local S = Theme.SIZES
				assert.are.equal(10, S.panelRegionGap)
				for _, kind in ipairs({ "rotation", "build" }) do
					local region = PanelTab.regionFrames[kind]
					local rule = region.rule
					assert.is_not_nil(rule, kind)
					assert.are.equal(region, rule.parent)
					assert.are.same(rgb("border"), colorOf(rule, "SetColorTexture"))
					assert.are.same({ PanelTab.innerWidth(), S.border }, { rule.width, rule.height })
					assertPoint({ "TOPLEFT", region, "TOPLEFT", 0, S.panelRegionGap / 2 }, rule.points[#rule.points])
				end
				assert.is_nil(PanelTab.regionFrames.targets.rule)
				-- The gap stays: the rotation region still sits a region gap under the targets.
				local targets, rotation = PanelTab.regionFrames.targets, PanelTab.regionFrames.rotation
				local below = -targets.points[#targets.points][5] + targets:GetHeight()
				assert.are.equal(below + S.panelRegionGap, -rotation.points[#rotation.points][5])
			end)

			it("draws the at-level rows at the dim alpha, icon, name and condition, still last and unnumbered", function()
				open({ chest = WORN_CHEST })
				local learned, later = PanelTab.rotationWidgets[1], PanelTab.rotationWidgets[2]
				assert.are.equal(Theme.ALPHA.dim, mock.lastCall(later.icon, "SetAlpha")[1])
				-- L3.7 item 24 and the designer's pre-item: the name and the condition line are lifted
				-- to the muted token at full alpha (dimmed again they were too dark to read); only the
				-- icon stays dim.
				assert.are.equal(1, mock.lastCall(later.name, "SetAlpha")[1])
				assert.are.equal(1, mock.lastCall(later.condition, "SetAlpha")[1])
				assert.are.same(rgb("muted"), colorOf(later.name))
				assert.are.same(rgb("muted"), colorOf(later.condition))
				for _, key in ipairs({ "icon", "name", "condition" }) do
					assert.are.equal(1, mock.lastCall(learned[key], "SetAlpha")[1], key)
				end
				local rows = regionOf(PanelTab.last, "rotation").rows
				assert.is_nil(rows[#rows].number)
				assert.is_true(rows[#rows].muted)
				assert.is_nil(rows[2].number)
				assert.are.equal(1, rows[1].number)
			end)
		end)

		-- Item 6: fewer than five targets.
		describe("fewer than five targets", function()
			it("adds one muted line saying the other slots are best for the level", function()
				open({ chest = WORN_CHEST, legs = THIN_LEGS })
				local S = Theme.SIZES
				local region = PanelTab.regionFrames.targets
				assert.are.equal("Nothing better for your level in other slots.", L.panelTargetsRest)
				assert.is_true(region.rest:IsShown())
				assert.are.equal(L.panelTargetsRest, region.rest:GetText())
				assert.are.same(rgb("muted"), colorOf(region.rest))
				local y = S.panelTitle + 2 * PanelTab.targetRowHeight() + 2 * S.panelRowGap
				assertPoint({ "TOPLEFT", region, "TOPLEFT", 0, -y }, region.rest.points[#region.rest.points])
				assert.are.equal(y + PanelTab.lineHeight(), region:GetHeight())
				local entry = regionOf(PanelTab.last, "targets")
				assert.is_true(entry.rest)
				assert.are.equal(PanelTab.regionHeight(entry), region:GetHeight())
			end)

			it("has no such line with five targets", function()
				start()
				local S = Theme.SIZES
				local targets = {}
				for index = 1, 5 do
					targets[index] = { slot = "chest", itemId = index, name = "T" .. index, gain = 1, unit = "dps" }
				end
				local five = PanelTab.model({ targets = targets }).regions[1]
				assert.is_false(five.rest)
				assert.are.equal(S.panelTitle + 5 * PanelTab.targetRowHeight() + 4 * S.panelRowGap,
					PanelTab.regionHeight(five))
				local four = PanelTab.model({ targets = { targets[1], targets[2], targets[3], targets[4] } }).regions[1]
				assert.is_true(four.rest)
				assert.are.equal(S.panelTitle + 4 * PanelTab.targetRowHeight() + 4 * S.panelRowGap + PanelTab.lineHeight(),
					PanelTab.regionHeight(four))
			end)
		end)

		-- Hotfix 2.0.2: from level 10, below the spec's first band (20 in the data), the band-20
		-- picks are listed as upcoming: "At level 20:" before the source, no gain, and the rest
		-- line says where the picks start.
		describe("below the first band", function()
			local function openLow(worn, options)
				local pane = start(worn, options)
				PanelTab.refresh(DATA)
				click(_G[PanelTab.BUTTON_NAME])
				return pane, _G[PanelTab.FLYOUT_NAME]
			end

			local function regionIn(model, kind)
				for _, region in ipairs(model.regions) do
					if region.kind == kind then
						return region
					end
				end
				return nil
			end

			it("lists the band-20 picks of the worn slots, upcoming and unscored", function()
				openLow({ chest = WORN_CHEST, legs = THIN_LEGS }, { level = 18 })
				local entry = regionIn(PanelTab.last, "targets")
				assert.are.equal(2, #entry.rows)
				local ids = {}
				for _, row in ipairs(entry.rows) do
					assert.is_nil(row.gain)
					assert.are.equal(20, row.upcoming)
					assert.is_nil(row.worn)
					ids[row.slot] = row.itemId
				end
				assert.are.same({ chest = 300, legs = 400 }, ids)
				assert.are.equal(20, entry.picksStart)
				assert.is_true(entry.rest)
				assert.are.equal(20, PanelTab.last.picksStart)
				assert.are.equal(2, PanelTab.last.targets)
			end)

			it("draws At level 20: before the source, no gain figure, and Picks start at level 20. under the rows", function()
				openLow({ chest = WORN_CHEST, legs = THIN_LEGS }, { level = 18 })
				local region = PanelTab.regionFrames.targets
				assert.are.equal("Picks start at level 20.", string.format(L.panelPicksStart, 20))
				assert.are.equal("At level 20:", string.format(L.panelTargetAt, 20))
				assert.are.equal("Picks start at level 20.", region.rest:GetText())
				assert.is_true(region.rest:IsShown())
				for _, row in ipairs(PanelTab.targetRows) do
					if row.frame:IsShown() then
						assert.are.equal("", row.gain:GetText())
						assert.are.equal("At level 20: ", row.detail:GetText():sub(1, 13))
						assert.is_nil(row.name:GetText():find("over", 1, true))
					end
				end
			end)

			it("keeps the region, with the line alone, when nothing worn differs", function()
				openLow({}, { level = 18 })
				local entry = regionIn(PanelTab.last, "targets")
				assert.are.equal(0, #entry.rows)
				assert.is_true(entry.rest)
				assert.are.equal("Picks start at level 20.", PanelTab.regionFrames.targets.rest:GetText())
			end)

			it("reads the plain rest line and the Next line from level 20, and picksStart below it", function()
				openLow({ chest = WORN_CHEST }, { level = 20 })
				assert.is_nil(PanelTab.last.picksStart)
				assert.are.equal(L.panelTargetsRest, PanelTab.regionFrames.targets.rest:GetText())
				openLow({ chest = WORN_CHEST }, { level = 9 })
				assert.are.equal(20, PanelTab.last.picksStart)
				openLow({ chest = WORN_CHEST }, { level = 6 })
				assert.are.equal(20, PanelTab.last.picksStart)
			end)

			it("prints the diag line with where the picks start", function()
				start({ chest = WORN_CHEST }, { level = 18 })
				local diag = PanelTab.diagLine()
				assert.is_truthy(diag:match("^Panel: closed, picks start at level 20, rotation %d+ rows$"))
				assert.are.equal("Panel: closed, picks start at level 20, rotation 3 rows",
					string.format(L.diagPanelPicksStart, L.diagPanelClosed, 20, 3))
				start({ chest = WORN_CHEST }, { level = 21 })
				assert.is_truthy(PanelTab.diagLine():match("^Panel: closed, %d+ targets, rotation %d+ rows$"))
			end)

			it("models a plain input with picksStart and no rows", function()
				start()
				local region = PanelTab.model({ targets = {}, picksStart = 20 }).regions[1]
				assert.are.equal("targets", region.kind)
				assert.is_true(region.rest)
				assert.are.equal(20, region.picksStart)
				assert.are.same({}, PanelTab.model({ targets = {} }).regions)
			end)
		end)
	end)
end)
