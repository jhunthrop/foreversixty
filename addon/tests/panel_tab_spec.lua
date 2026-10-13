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
				_G.CreateFrame("Button", "CharacterFrameTab" .. index, pane)
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
			assert.are.same({}, PanelTab.rotationRows(DATA, BUILD, { [2] = { ["1:1"] = 5 } }, 21))
		end)
	end)

	-- Ruling 4: the split in the dialog's words, progress and the off-plan count.
	describe("build", function()
		it("reads the split, the points on plan and the points off it", function()
			start()
			local ranks = require("Talents").readRanks(DATA)
			local region = PanelTab.buildRegion(BUILD, DATA, ranks, 21)
			assert.are.equal(string.format(L.dialogSplitPart, "Arms", 3) .. L.dialogSplitSeparator
				.. string.format(L.dialogSplitPart, "Fury", 1), region.split)
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
			assert.are.same({ "LEFT", _G.CharacterFrameTab3, "RIGHT", Theme.SIZES.panelTabGap, 0 },
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
			assert.are.same({ "TOPRIGHT", pane, "BOTTOMRIGHT", 0, 0 }, button.points[#button.points])
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
		end)

		it("opens the flyout at the pane's top right on a click, and closes it on the next", function()
			local pane = start({ chest = WORN_CHEST })
			PanelTab.refresh(DATA)
			click(_G[PanelTab.BUTTON_NAME])
			local flyout = _G[PanelTab.FLYOUT_NAME]
			assert.are.equal(pane, flyout.parent)
			assert.are.same({ "TOPLEFT", pane, "TOPRIGHT", Theme.SIZES.panelOffsetX, Theme.SIZES.panelOffsetY },
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
			assert.are.equal(string.format(L.panelTargetDetail, L.gearSlotLabels.chest,
				"Wailing Caverns · Mutanus the Devourer · 33% drop"), row.detail:GetText())
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
				'"OnShow"', '"OnHide"' }) do
				assert.is_nil(source:find(call, 1, true), call)
			end
		end)
	end)
end)
