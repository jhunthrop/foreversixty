-- addon/tests/group_panel_spec.lua
-- The group panel and the ready check (spec sections 4.8, 4.9, 5.4; task L5.2): an addon-owned
-- "Forever Sixty" button under the client's raid frame (RaidFrame, the child FriendsFrame shows
-- on its Raid tab) and an addon-owned flyout, 480 wide, beside it: one row per party or raid
-- unit, the player included (spec, level, item level, gear gap, Forever Score, enchants, addon
-- present), the weakest slot named under each row, and after a ready check who lacks a flask,
-- food, an imbue or a class buff and who covers each of the encounter's needs. Ruling 1
-- (taint): nothing is written into a client frame, the recorder stays empty in every example.
local helper = require("spec_helper")
local mock = require("wow_mock")
local L = require("Locale")

local QUALITY = {
	[1] = { r = 1, g = 1, b = 1 }, [2] = { r = 0.12, g = 1, b = 0 }, [3] = { r = 0, g = 0.44, b = 0.87 },
	[4] = { r = 0.64, g = 0.21, b = 0.93 }, [5] = { r = 1, g = 0.5, b = 0 },
}

local CLASS_COLORS = {
	PRIEST = { r = 1, g = 1, b = 1 }, WARRIOR = { r = 0.78, g = 0.61, b = 0.43 },
	MAGE = { r = 0.25, g = 0.78, b = 0.92 }, DRUID = { r = 1, g = 0.49, b = 0.04 },
}

--- The client's spell names for the ids the data and the check read.
local SPELLS = {
	[11551] = "Battle Shout", [9885] = "Mark of the Wild", [10157] = "Arcane Intellect",
	[10938] = "Power Word: Fortitude", [20217] = "Blessing of Kings", [11597] = "Sunder Armor",
	[11717] = "Curse of Recklessness", [9907] = "Faerie Fire",
	[17626] = "Flask of the Titans", [17627] = "Distilled Wisdom", [17628] = "Supreme Power",
	[17629] = "Chromatic Resistance", [17624] = "Petrification", [19705] = "Well Fed",
	-- The group variants (fix round 1, D2).
	[21849] = "Gift of the Wild", [21850] = "Gift of the Wild", [21562] = "Prayer of Fortitude",
	[21564] = "Prayer of Fortitude", [450086] = "Prayer of Fortitude", [23028] = "Arcane Brilliance",
	[25898] = "Greater Blessing of Kings",
}

--- An aura as UnitAura answers it: the name and the spell id.
local function aura(id, name)
	return { name = name or SPELLS[id], spellId = id }
end

--- The nightly record (contract 5, format 2): Thorgrim epic 1610 with his best parses, Elwen
--- rare 980, the player with no logs.
local function nightly()
	return {
		format = 2,
		generated = os.date("!%Y-%m-%d"),
		characters = {
			["Thorgrim-ClassicBetaPvP"] = { score = 1610, tier = "epic", as_of = "2026-11-04", fights = 20,
				spec = "warrior-arms", best_parses = { ["Lucifron"] = 71, ["Garr"] = 64 } },
			["Elwen-ClassicBetaPvP"] = { score = 980, tier = "rare", as_of = "2026-11-04", fights = 9 },
		},
		guilds = {},
	}
end

--- The party: the player (a priest), Thorgrim (a warrior who runs the addon) and Elwen (a mage
--- who does not, inspected).
local PARTY = {
	{ unit = "player", name = "Tester", class = "PRIEST", className = "Priest", level = 60 },
	{ unit = "party1", name = "Thorgrim", class = "WARRIOR", className = "Warrior", level = 60 },
	{ unit = "party2", name = "Elwen", class = "MAGE", className = "Mage", level = 58 },
}

--- The data the panel reads: the classes' trees, an arcane mage's weights, DPS per point and
--- band-60 chest pick, and two bosses' needs (Lucifron in Molten Core, Onyxia).
local function groupData()
	return {
		build = "1.60.1.69893",
		classes = {
			priest = { tabs = { { name = "Discipline", talents = {} }, { name = "Holy", talents = {} },
				{ name = "Shadow", talents = {} } } },
			warrior = { tabs = { { name = "Arms", talents = {} }, { name = "Fury", talents = {} },
				{ name = "Protection", talents = {} } } },
			mage = { tabs = { { name = "Arcane", talents = {} }, { name = "Fire", talents = {} },
				{ name = "Frost", talents = {} } } },
		},
		weights = { ["mage-arcane"] = { intellect = 1 } },
		dpsPerPoint = { ["mage-arcane"] = { reference = "spell_power", dps = 0.5, weaponDps = 0, healing = false,
			[50] = { reference = "spell_power", dps = 0.5, weaponDps = 0 } },
			-- Arms has a DPS per point too, so a live warrior's gap reads in DPS beside the mage's.
			["warrior-arms"] = { reference = "attack_power", dps = 0.1, weaponDps = 0, healing = false,
				[60] = { reference = "attack_power", dps = 0.1, weaponDps = 0 } } },
		bis = { ["mage-arcane"] = { [50] = { alliance = { chest = "100|D|Lucifron" } } } },
		bisNames = { [100] = "Robes of Arcana" },
		encounterNeeds = {
			[12118] = { buffs = { 10157, 10938, 9885, 11551 }, debuffs = { 11597, 9907 }, resist = "fire" },
			[10184] = { buffs = { 10157, 10938 }, debuffs = { 11717 } },
		},
	}
end

describe("GroupPanel", function()
	local GroupPanel, Theme, Comms, Inspect, Prefs, Chip
	local state, clock, sent, notified, combat, units, auras, gear, groupList

	--- A recorded point equals `expected` element by element, frames by identity.
	local function assertPoint(expected, actual)
		assert.is_not_nil(actual)
		assert.are.equal(#expected, #actual)
		for index = 1, #expected do
			assert.is_true(rawequal(expected[index], actual[index]), "point element " .. index)
		end
	end

	--- A client frame given a rect: GetLeft answers a number (L4.3 fix round 2). Set from the spec,
	--- so the recorder does not count it.
	local function withRect(frame)
		frame.GetLeft = function()
			return 100
		end
		frame.GetRight = function()
			return 100 + frame:GetWidth()
		end
		return frame
	end

	--- A client frame's shown flag changed with no script run (fix round 2): a state the panel was
	--- not told of, which a refresh point must read for itself. Set from the spec, not recorded.
	local function unheard(frame, shown)
		frame.shown = shown
	end

	--- The client before its drawn window exists (fix round 2, N3/N4): FriendsFrame on UIParent,
	--- hidden, RaidFrame in it, no SocialUIFrame yet. Answers { friends, raid }.
	local function beforeSocialWindow()
		local friends = _G.CreateFrame("Frame", "FriendsFrame", _G.UIParent)
		local raid = _G.CreateFrame("Frame", "RaidFrame", friends)
		friends:Hide()
		return { friends = friends, raid = raid }
	end

	--- The client's drawn Social window, made now (shown, with a rect).
	local function socialWindow()
		local social = withRect(_G.CreateFrame("Frame", "SocialUIFrame", _G.UIParent))
		social:SetSize(400, 512)
		return social
	end

	--- The group the client answers. `options`: { members, raid, solo, count, auras, instance,
	--- leader, noLeaderApi, weaponEnchant, noAuraApi, auraUtil, frames, noFrame, noSpellNames, data,
	--- sharingOff }. `frames` (fix round 1, D1): "live" (the default, the 1.60 client: a drawn
	--- SocialUIFrame with a rect holding the rectless logical FriendsFrame and RaidFrame),
	--- "friends" (no SocialUIFrame; FriendsFrame with a rect), "raid" (only RaidFrame with a rect).
	--- Answers { social, friends, raid, anchor } (the frame the panel should anchor to).
	local function start(options)
		options = options or {}
		-- A second start in one example: the first client is checked and taken down first.
		if state ~= nil then
			assert.are.same({}, state.clientWrites)
			mock.uninstall()
		end
		clock, sent, notified, combat = 1000, {}, {}, false
		groupList = {}
		for _, entry in ipairs(options.solo and { PARTY[1] } or options.members or PARTY) do
			groupList[#groupList + 1] = entry
		end
		units = {}
		for _, entry in ipairs(groupList) do
			units[entry.unit] = entry
		end
		auras = options.auras or {}
		gear = options.gear or {}
		local function member(unit)
			return units[unit]
		end
		local globals = {
			IsInRaid = function()
				return options.raid == true
			end,
			IsInGroup = function()
				return not options.solo
			end,
			GetNumGroupMembers = function()
				return options.count or (options.solo and 0 or #groupList)
			end,
			GetNumSubgroupMembers = function()
				return options.solo and 0 or #groupList - 1
			end,
			UnitName = function(unit)
				local entry = member(unit)
				return entry and entry.name or nil, nil
			end,
			UnitClass = function(unit)
				local entry = member(unit)
				return entry and entry.className or nil, entry and entry.class or nil
			end,
			UnitLevel = function(unit)
				local entry = member(unit)
				return entry and entry.level or 0
			end,
			UnitIsConnected = function(unit)
				local entry = member(unit)
				return entry ~= nil and entry.offline ~= true
			end,
			UnitIsUnit = function(left, right)
				local a, b = member(left), member(right)
				return a ~= nil and a == b
			end,
			UnitExists = function(unit)
				return member(unit) ~= nil
			end,
			UnitIsPlayer = function(unit)
				return member(unit) ~= nil
			end,
			UnitGUID = function(unit)
				local entry = member(unit)
				return entry and "Player-" .. entry.name or nil
			end,
			UnitIsGroupLeader = function(unit)
				return unit == "player" and options.leader == true
			end,
			NotifyInspect = function(unit)
				notified[#notified + 1] = unit
			end,
			CheckInteractDistance = function()
				return true
			end,
			CanInspect = function()
				return true
			end,
			GetInventoryItemLink = function(unit, slot)
				local entry = member(unit)
				return entry and (gear[entry.name] or {})[slot] or nil
			end,
			GetDetailedItemLevelInfo = function(link)
				local info = state.itemStats[link]
				return info and info.__level or nil
			end,
			GetInstanceInfo = function()
				local instance = options.instance or { "Elwynn Forest", "none", 0 }
				return instance[1], instance[2], 0, "", 40, 0, false, instance[3]
			end,
			GetWeaponEnchantInfo = function()
				return options.weaponEnchant == true
			end,
			InCombatLockdown = function()
				return combat
			end,
			GetTime = function()
				return clock
			end,
			C_ChatInfo = { SendAddonMessage = function(prefix, text, channel)
				sent[#sent + 1] = { prefix = prefix, text = text, channel = channel }
				return true
			end },
			RAID_CLASS_COLORS = CLASS_COLORS,
			ITEM_QUALITY_COLORS = QUALITY,
			GetNormalizedRealmName = function()
				return "ClassicBetaPvP"
			end,
			UnitFactionGroup = function()
				return "Alliance"
			end,
		}
		if options.noLeaderApi then
			globals.UnitIsGroupLeader = nil
		end
		if not options.noSpellNames then
			globals.C_Spell = { GetSpellName = function(id)
				return SPELLS[id]
			end }
		end
		local function readAura(unit, index)
			local entry = member(unit)
			local list = entry and auras[entry.name] or {}
			local found = list[index]
			if found == nil then
				return nil
			end
			return found.name, "icon", 0, nil, 0, 0, "player", false, false, found.spellId
		end
		if options.auraUtil then
			globals.AuraUtil = { ForEachAura = function(unit, filter, _, action, packed)
				assert.are.equal("HELPFUL", filter)
				assert.is_true(packed)
				local entry = member(unit)
				for _, found in ipairs(entry and auras[entry.name] or {}) do
					if action({ name = found.name, spellId = found.spellId }) then
						return
					end
				end
			end }
		elseif not options.noAuraApi then
			globals.UnitAura = readAura
		end
		state = mock.install({
			class = { name = "Priest", token = "PRIEST" },
			realm = "Classic Beta PvP",
			region = 1,
			playerName = "Tester",
			level = 60,
			globals = globals,
		})
		-- The client's clock as date("%H:%M") reads it.
		state.date = "12:00"
		_G.ForeverSixtyData = options.data == nil and nightly() or options.data or nil
		_G.ForeverSixtyDB = {}
		_G.ForeverSixtyInbox = nil
		Theme = require("Theme")
		Theme.reset()
		Prefs = helper.load("Prefs")
		helper.load("Compat")
		helper.load("Export")
		helper.load("Score")
		helper.load("Follow")
		helper.load("Gear")
		helper.load("Progress")
		Comms = helper.load("Comms")
		helper.load("Widgets")
		helper.load("ItemTooltip")
		helper.load("PlayerTooltip")
		Inspect = helper.load("Inspect")
		Chip = helper.load("Chip")
		GroupPanel = helper.load("GroupPanel")
		if options.sharingOff then
			Prefs.setFlag("broadcast", false)
		end
		GroupPanel.data = groupData()
		Comms.data = GroupPanel.data
		local frames = {}
		if not options.noFrame then
			local kind = options.frames or "live"
			local holder = _G.UIParent
			if kind == "live" then
				-- 1.60: the drawn window; FriendsFrame and RaidFrame show but have no rect (width 0).
				frames.social = withRect(_G.CreateFrame("Frame", "SocialUIFrame", _G.UIParent))
				frames.social:SetSize(400, 512)
				holder = frames.social
			end
			frames.friends = _G.CreateFrame("Frame", "FriendsFrame", holder)
			frames.raid = _G.CreateFrame("Frame", "RaidFrame", frames.friends)
			if kind == "friends" then
				withRect(frames.friends):SetSize(384, 512)
			elseif kind == "raid" then
				withRect(frames.raid):SetSize(338, 424)
			end
			frames.anchor = frames.social or kind == "friends" and frames.friends or frames.raid
		end
		return frames
	end

	after_each(function()
		if state ~= nil then
			assert.are.same({}, state.clientWrites)
		end
		_G.ForeverSixtyData = nil
		_G.ForeverSixtyInbox = nil
		mock.uninstall()
		state = nil
	end)

	local function click(button)
		button:GetScript("OnClick")(button, "LeftButton")
	end

	--- Opens the flyout with a click on the button; answers it.
	local function open()
		GroupPanel.refresh()
		click(_G[GroupPanel.BUTTON_NAME])
		return _G[GroupPanel.FLYOUT_NAME]
	end

	--- An inspect record for Elwen as INSPECT_READY leaves it in the cache: an arcane mage at
	--- item level 52 wearing a Worn Robe with no enchant (+ a wand), the record fresh now.
	local function inspectElwen()
		Inspect.cache["Elwen"] = {
			key = "Elwen", name = "Elwen", guid = "Player-Elwen", class = "mage", spec = "mage-arcane",
			level = 58, itemLevel = 52, at = clock,
			items = { chest = "item:50:0", ranged = "item:60:0" },
			enchants = { chest = false, ranged = false },
		}
	end

	--- The client's item cache for the inspect's links and the band's chest pick.
	local function itemCache()
		state.itemStats["item:50:0"] = { __itemId = 50, __slot = "INVTYPE_CHEST", __name = "Worn Robe",
			ITEM_MOD_INTELLECT_SHORT = 4 }
		state.itemStats[50] = { __itemId = 50, __slot = "INVTYPE_CHEST", __name = "Worn Robe" }
		state.itemStats["item:60:0"] = { __itemId = 60, __slot = "INVTYPE_RANGEDRIGHT", __name = "Lesser Wand" }
		state.itemStats[100] = { __itemId = 100, __slot = "INVTYPE_ROBE", __name = "Robes of Arcana",
			ITEM_MOD_INTELLECT_SHORT = 28 }
	end

	--- Thorgrim's addon said hello: arms, a 14 DPS gap, his raid build.
	local function thorgrimHello()
		Comms.seen["Thorgrim-ClassicBetaPvP"] = { level = 60, spec = "warrior-arms", gearGap = 14,
			buildName = "Arms raid", scoreDate = "2026-11-04", channel = "PARTY", time = clock }
	end

	local function rowFor(rows, short)
		for _, row in ipairs(rows) do
			if row.short == short then
				return row
			end
		end
		return nil
	end

	local function names(rows)
		local list = {}
		for index, row in ipairs(rows) do
			list[index] = row.short
		end
		return list
	end

	--- The drawn row frame whose name cell reads `short`.
	local function drawnRow(short)
		for _, row in ipairs(GroupPanel.rowFrames) do
			if row.frame:IsShown() and row.cells.name:GetText() == short then
				return row
			end
		end
		return nil
	end

	local function colorOf(region)
		local call = mock.lastCall(region, "SetTextColor")
		return call ~= nil and Theme.hexOf(call[1], call[2], call[3]) or nil
	end

	-- Ruling 2: the rows, a pure join of what the client, the inspect, the record and the addon say.
	describe("the rows", function()
		it("joins each unit with its inspect, its record and its addon, the player included", function()
			start()
			itemCache()
			inspectElwen()
			thorgrimHello()
			local rows, counts = GroupPanel.rows(GroupPanel.members(GroupPanel.data))
			assert.are.same({ members = 3, inspected = 2, live = 2 }, counts)
			local thorgrim = rowFor(rows, "Thorgrim")
			assert.are.equal("WARRIOR", thorgrim.classToken)
			assert.are.equal("warrior-arms", thorgrim.spec)
			assert.are.equal(60, thorgrim.level)
			assert.are.equal(14, thorgrim.gap)
			assert.is_true(thorgrim.live)
			assert.are.equal(1610, thorgrim.score)
			assert.are.equal("epic", thorgrim.tier)
			assert.are.same({ boss = "Lucifron", percentile = 71 }, thorgrim.bestParse)
			assert.is_nil(thorgrim.ilvl)
			assert.is_false(thorgrim.inspected)
			local elwen = rowFor(rows, "Elwen")
			assert.are.equal("mage-arcane", elwen.spec)
			assert.are.equal(58, elwen.level)
			assert.are.equal(52, elwen.ilvl)
			-- Two slots worn bare; the wand is not an enchantable slot.
			assert.are.equal(1, elwen.enchants)
			assert.is_true(elwen.inspected)
			assert.is_false(elwen.live)
			assert.are.equal(980, elwen.score)
			-- The player's own row reads this client: live, and inspected from its own gear.
			local own = rowFor(rows, "Tester")
			assert.is_true(own.own)
			assert.is_true(own.live)
			assert.is_true(own.inspected)
			assert.is_nil(own.score)
		end)

		it("measures an inspected member's gap in their spec and names the weakest slot under the row", function()
			start()
			itemCache()
			inspectElwen()
			local rows = GroupPanel.rows(GroupPanel.members(GroupPanel.data))
			local elwen = rowFor(rows, "Elwen")
			-- Worn Robe 4 Int against the band's Robes of Arcana 28 Int at 0.5 DPS a point: +12 DPS.
			assert.are.equal(12, elwen.gap)
			assert.are.equal("dps", elwen.unit)
			assert.are.equal("Chest: Worn Robe · next Robes of Arcana +12.0 DPS", GroupPanel.subLine(elwen))
			-- The same one target and scoring as the item tooltip and the hello (Progress.gearPlan).
			local Progress, Gear = require("Progress"), require("Gear")
			local current = Gear.currentWeights(GroupPanel.data, "mage-arcane", nil, nil, 58)
			local plan = Progress.gearPlan(GroupPanel.data, "mage-arcane", 58, "alliance",
				Gear.equippedFrom(Inspect.cache["Elwen"].items), current)
			assert.are.same({ 12, "dps", "chest", 100 }, { plan.gap, plan.unit, plan.weakest.slot, plan.weakest.itemId })
		end)

		-- Fix round 1: a client answering GetNumGroupMembers 1 for a player alone is not a group.
		it("counts nobody out of a group, even where the client answers one member", function()
			start({ solo = true, count = 1 })
			assert.are.same({}, GroupPanel.groupUnits())
			assert.is_false(GroupPanel.inGroup())
			assert.are.same({}, GroupPanel.members(GroupPanel.data))
		end)

		-- Fix round 1: the Gap of a live member is their addon's; the weakest line is the inspect's,
		-- tagged "(inspect)" where the two disagree, and absent without an inspect.
		it("tags the weakest line (inspect) where a live gap and the inspected one differ, none without an inspect",
			function()
				start()
				itemCache()
				inspectElwen()
				Comms.seen["Elwen-ClassicBetaPvP"] = { level = 58, spec = "mage-arcane", gearGap = 3, time = clock }
				local elwen = rowFor(GroupPanel.rows(GroupPanel.members(GroupPanel.data)), "Elwen")
				assert.are.equal(3, elwen.gap)
				assert.are.equal("Chest: Worn Robe · next Robes of Arcana +12.0 DPS " .. L.groupWeakestInspect,
					GroupPanel.subLine(elwen))
				-- The same gap both ways: no tag.
				Comms.seen["Elwen-ClassicBetaPvP"].gearGap = 12
				elwen = rowFor(GroupPanel.rows(GroupPanel.members(GroupPanel.data)), "Elwen")
				assert.are.equal("Chest: Worn Robe · next Robes of Arcana +12.0 DPS", GroupPanel.subLine(elwen))
				-- Live without an inspect: the gap, and no weakest line.
				thorgrimHello()
				local thorgrim = rowFor(GroupPanel.rows(GroupPanel.members(GroupPanel.data)), "Thorgrim")
				assert.are.equal(14, thorgrim.gap)
				assert.are.equal("", GroupPanel.subLine(thorgrim))
			end)

		-- Fix round 1: an inspect that came back with no item links is no gear to score.
		it("ignores a cached inspect with no item links rather than scoring an empty set", function()
			start()
			itemCache()
			Inspect.cache["Elwen"] = { key = "Elwen", name = "Elwen", guid = "Player-Elwen", class = "mage",
				spec = "mage-arcane", level = 58, itemLevel = nil, at = clock, items = {}, enchants = {} }
			local members = GroupPanel.members(GroupPanel.data)
			local elwen = rowFor(GroupPanel.rows(members), "Elwen")
			assert.is_nil(elwen.gap)
			assert.is_nil(elwen.weakest)
			assert.is_nil(elwen.enchants)
			assert.are.equal("", GroupPanel.subLine(elwen))
			assert.are.equal("", GroupPanel.cells(elwen).enchants.text)
		end)

		it("takes a member's addon over the inspect for the spec and the gap", function()
			start()
			itemCache()
			inspectElwen()
			Comms.seen["Elwen-ClassicBetaPvP"] = { level = 58, spec = "mage-fire", gearGap = 3, buildName = "Fire",
				time = clock }
			local elwen = rowFor(GroupPanel.rows(GroupPanel.members(GroupPanel.data)), "Elwen")
			assert.are.equal("mage-fire", elwen.spec)
			assert.are.equal(3, elwen.gap)
			assert.are.equal("Fire", elwen.buildName)
			assert.is_true(elwen.live)
		end)

		it("sorts by Gap high to low by default, blanks last, and by any header either way", function()
			start()
			itemCache()
			inspectElwen()
			thorgrimHello()
			local members = GroupPanel.members(GroupPanel.data)
			-- Thorgrim 14, Elwen 12, the player blank (nothing to score).
			assert.are.same({ "Thorgrim", "Elwen", "Tester" }, names(GroupPanel.rows(members)))
			assert.are.same({ "Elwen", "Thorgrim", "Tester" },
				names(GroupPanel.rows(members, { key = "gap", descending = false })))
			assert.are.same({ "Elwen", "Tester", "Thorgrim" },
				names(GroupPanel.rows(members, { key = "name", descending = false })))
			assert.are.same({ "Thorgrim", "Elwen", "Tester" },
				names(GroupPanel.rows(members, { key = "score", descending = true })))
			assert.are.same({ "Tester", "Thorgrim", "Elwen" },
				names(GroupPanel.rows(members, { key = "level", descending = true })))
			assert.are.same({ key = "gap", descending = false }, GroupPanel.nextSort(nil, "gap"))
			assert.are.same({ key = "name", descending = false }, GroupPanel.nextSort(nil, "name"))
			assert.are.same({ key = "ilvl", descending = true }, GroupPanel.nextSort(nil, "ilvl"))
		end)

		it("reads inspect paused, muted, for a member not yet inspected while in combat", function()
			start()
			combat = true
			local rows = GroupPanel.rows(GroupPanel.members(GroupPanel.data), nil, { paused = Inspect.isPaused() })
			assert.are.equal(L.groupInspectPaused, GroupPanel.subLine(rowFor(rows, "Elwen")))
			assert.are.equal(L.groupInspectPaused, GroupPanel.subLine(rowFor(rows, "Thorgrim")))
			-- The player's own row reads this client and never waits on an inspect.
			assert.are.equal("", GroupPanel.subLine(rowFor(rows, "Tester")))
		end)

		it("tags a row under the score floor, muted, and declines nobody", function()
			start()
			itemCache()
			inspectElwen()
			local rows = GroupPanel.rows(GroupPanel.members(GroupPanel.data), nil, { floor = 1200 })
			assert.is_true(rowFor(rows, "Elwen").below)
			assert.is_false(rowFor(rows, "Thorgrim").below)
			-- No score is no verdict.
			assert.is_false(rowFor(rows, "Tester").below)
			assert.are.equal(string.format(L.groupBelow, 1200) .. L.groupSubJoin
				.. "Chest: Worn Robe · next Robes of Arcana +12.0 DPS", GroupPanel.subLine(rowFor(rows, "Elwen")))
			assert.are.equal(3, #rows)
		end)
	end)

	describe("the cells", function()
		it("draws the name in its class colour, enchants as N missing muted or the client's tick, live or blank",
			function()
				start()
				itemCache()
				inspectElwen()
				thorgrimHello()
				local rows = GroupPanel.rows(GroupPanel.members(GroupPanel.data))
				local elwen = GroupPanel.cells(rowFor(rows, "Elwen"))
				assert.are.equal("Elwen", elwen.name.text)
				assert.are.equal(Theme.panelClassHex("MAGE"), elwen.name.hex)
				assert.are.equal("Arcane", elwen.spec.text)
				assert.are.equal("58", elwen.level.text)
				assert.are.equal("52", elwen.ilvl.text)
				assert.are.equal(string.format(L.groupGapDps, 12), elwen.gap.text)
				assert.are.equal("980", elwen.score.text)
				assert.are.equal(Theme.HEX.rareBright, elwen.score.hex)
				assert.are.equal(string.format(L.groupEnchantsMissing, 1), elwen.enchants.text)
				assert.are.equal(Theme.HEX.muted, elwen.enchants.hex)
				assert.are.equal("", elwen.addon.text)
				local thorgrim = GroupPanel.cells(rowFor(rows, "Thorgrim"))
				assert.are.equal(L.groupAddonLive, thorgrim.addon.text)
				-- His live gap in the unit his spec has (Comms.gapUnit: a DPS per point for Arms).
				assert.are.equal(string.format(L.groupGapDps, 14), thorgrim.gap.text)
				-- A spec with no DPS per point reads in points; 0 muted.
				local points = GroupPanel.cells({ short = "Fenna", gap = 9, unit = "pts", live = true })
				assert.are.equal(string.format(L.groupGapPts, 9), points.gap.text)
				local even = GroupPanel.cells({ short = "Fenna", gap = 0, unit = "dps", live = true })
				assert.are.equal(Theme.HEX.muted, even.gap.hex)
				-- Not inspected: no enchant verdict at all.
				assert.are.equal("", thorgrim.enchants.text)
				Inspect.cache["Elwen"].enchants = { chest = true, ranged = false }
				local fixed = GroupPanel.cells(rowFor(GroupPanel.rows(GroupPanel.members(GroupPanel.data)), "Elwen"))
				assert.are.equal(Theme.inlineIcon(Theme.MEDIA.gearTick, Theme.SIZES.groupTick), fixed.enchants.text)
			end)

		it("reads no logs muted for a member the readable record lacks, and says so on the row's hover", function()
			start()
			thorgrimHello()
			local rows = GroupPanel.rows(GroupPanel.members(GroupPanel.data), nil, { ratings = true })
			local own = GroupPanel.cells(rowFor(rows, "Tester"))
			assert.are.equal(L.groupNoLogs, own.score.text)
			assert.are.equal(Theme.HEX.muted, own.score.hex)
			local hint = table.concat(GroupPanel.rowHint(rowFor(rows, "Thorgrim")), "\n")
			assert.is_truthy(hint:find(string.format(L.groupBestParse, "71st", "Lucifron"), 1, true))
			assert.is_truthy(hint:find(string.format(L.groupBuild, "Arms raid"), 1, true))
		end)
	end)

	describe("the frames", function()
		-- Fix round 1, D1: on 1.60 RaidFrame and FriendsFrame show but have no rect; the drawn
		-- Social window is SocialUIFrame. The panel anchors to the first of SocialUIFrame,
		-- FriendsFrame and RaidFrame that shows with a rect, parented to it.
		it("puts the button under the Social window's bottom right and the flyout at its top right, 480 wide",
			function()
				local frames = start()
				local social = frames.social
				GroupPanel.refresh()
				local button = _G[GroupPanel.BUTTON_NAME]
				assert.are.equal(social, button.parent)
				assertPoint({ "TOPRIGHT", social, "BOTTOMRIGHT", 0, 0 }, button.points[#button.points])
				assert.are.equal(L.groupPanelLabel, button:GetText())
				assert.are.same({ Theme.SIZES.panelTabWidth, Theme.SIZES.panelTabHeight },
					{ button:GetWidth(), button:GetHeight() })
				assert.are.equal(social, GroupPanel.anchor())
				assert.is_true(GroupPanel.tabOpen())
				-- The flyout waits for the first click.
				assert.is_nil(_G[GroupPanel.FLYOUT_NAME])
				local flyout = open()
				assert.are.equal(social, flyout.parent)
				assert.are.equal(480, flyout:GetWidth())
				assertPoint({ "TOPLEFT", social, "TOPRIGHT", Theme.SIZES.panelOffsetX, Theme.SIZES.panelOffsetY },
					flyout.points[#flyout.points])
				assert.is_true(flyout:IsShown())
				assert.is_true(flyout:IsVisible())
				click(button)
				assert.is_false(flyout:IsShown())
				assert.are.equal(0, mock.countCalls(flyout, "SetScript"))
			end)

		it("answers /fs group with the Raid tab open by showing the flyout and no hint (the 1.60 shape)", function()
			start()
			GroupPanel.refresh()
			assert.are.same({}, GroupPanel.command())
			local flyout = _G[GroupPanel.FLYOUT_NAME]
			assert.is_true(flyout:IsVisible())
			assert.is_truthy(GroupPanel.diagLine():find(L.diagPanelOpen, 1, true))
		end)

		it("anchors to FriendsFrame where it has a rect and there is no SocialUIFrame", function()
			local frames = start({ frames = "friends" })
			assert.is_nil(_G.SocialUIFrame)
			GroupPanel.refresh()
			local button = _G[GroupPanel.BUTTON_NAME]
			assert.are.equal(frames.friends, button.parent)
			assertPoint({ "TOPRIGHT", frames.friends, "BOTTOMRIGHT", 0, 0 }, button.points[#button.points])
			assert.are.same({}, GroupPanel.command())
			local flyout = _G[GroupPanel.FLYOUT_NAME]
			assert.are.equal(frames.friends, flyout.parent)
			assertPoint({ "TOPLEFT", frames.friends, "TOPRIGHT", Theme.SIZES.panelOffsetX, Theme.SIZES.panelOffsetY },
				flyout.points[#flyout.points])
		end)

		it("anchors to RaidFrame where only it has a rect", function()
			local frames = start({ frames = "raid" })
			GroupPanel.refresh()
			local button = _G[GroupPanel.BUTTON_NAME]
			assert.are.equal(frames.raid, button.parent)
			assertPoint({ "TOPRIGHT", frames.raid, "BOTTOMRIGHT", 0, 0 }, button.points[#button.points])
			assert.are.same({}, GroupPanel.command())
			assert.are.equal(frames.raid, _G[GroupPanel.FLYOUT_NAME].parent)
		end)

		it("makes no button without a social frame; /fs group opens the flyout on UIParent at the chip's side",
			function()
				start({ noFrame = true })
				assert.has_no.errors(function()
					GroupPanel.refresh()
				end)
				assert.is_nil(_G[GroupPanel.BUTTON_NAME])
				local chip = _G.CreateFrame("Button", "ForeverSixtyChip", _G.UIParent)
				Chip.frame = chip
				assert.are.same({}, GroupPanel.command())
				local flyout = _G[GroupPanel.FLYOUT_NAME]
				assert.are.equal(_G.UIParent, flyout.parent)
				assertPoint({ "TOPRIGHT", chip, "TOPLEFT", -Theme.SIZES.gap, 0 }, flyout.points[#flyout.points])
				assert.is_true(flyout:IsShown())
				GroupPanel.command()
				assert.is_false(flyout:IsShown())
			end)

		it("says where it shows when /fs group runs with the Social window closed, and shows there", function()
			local frames = start()
			frames.social:Hide()
			GroupPanel.refresh()
			assert.is_nil(GroupPanel.anchor())
			assert.is_false(GroupPanel.tabOpen())
			assert.are.same({ L.groupOpensOnRaidTab }, GroupPanel.command())
			local flyout = _G[GroupPanel.FLYOUT_NAME]
			-- Parented to the window, so it shows with it.
			assert.are.equal(frames.social, flyout.parent)
			assert.is_true(flyout:IsShown())
			assert.is_false(flyout:IsVisible())
			-- The window opens on the Raid tab: the show scripts (the button's, the sentinel's) place and
			-- draw the flyout (the mock runs them down from the window, fix round 2).
			local draws = GroupPanel.draws
			frames.social:Show()
			assert.is_true(GroupPanel.draws > draws)
			assertPoint({ "TOPLEFT", frames.social, "TOPRIGHT", Theme.SIZES.panelOffsetX, Theme.SIZES.panelOffsetY },
				flyout.points[#flyout.points])
			assert.are.same({}, GroupPanel.command())
			assert.is_false(flyout:IsShown())
		end)

		it("prints the hint on another Social tab (RaidFrame hidden), never on the Raid tab", function()
			local frames = start()
			frames.raid:Hide()
			GroupPanel.refresh()
			assert.is_false(GroupPanel.tabOpen())
			assert.are.same({ L.groupOpensOnRaidTab }, GroupPanel.command())
			GroupPanel.command()
			frames.raid:Show()
			assert.is_true(GroupPanel.tabOpen())
			assert.are.same({}, GroupPanel.command())
		end)

		it("reads the Raid tab at every refresh point: the button rests on another tab, wakes on the Raid tab",
			function()
				local frames = start()
				GroupPanel.refresh()
				local button = _G[GroupPanel.BUTTON_NAME]
				assert.are.equal(1, mock.lastCall(button, "SetAlpha")[1])
				-- Another Social tab the panel was not told of: the button rests (no alpha, no mouse) but
				-- keeps its own shown flag, so its show script still runs the next time the window opens.
				unheard(frames.raid, false)
				GroupPanel.onRosterUpdate()
				assert.are.equal(0, mock.lastCall(button, "SetAlpha")[1])
				assert.is_false(mock.lastCall(button, "EnableMouse")[1])
				assert.is_true(button:IsShown())
				-- Back on the Raid tab: an inspect arriving wakes it.
				unheard(frames.raid, true)
				GroupPanel.onInspectReady({ key = "Elwen" })
				assert.are.equal(1, mock.lastCall(button, "SetAlpha")[1])
				assert.is_true(mock.lastCall(button, "EnableMouse")[1])
				-- PLAYER_ENTERING_WORLD (Options' refresh) reads it too.
				unheard(frames.raid, false)
				GroupPanel.refresh()
				assert.are.equal(0, mock.lastCall(button, "SetAlpha")[1])
				-- And the button's own show script.
				unheard(frames.raid, true)
				button:GetScript("OnShow")(button)
				assert.are.equal(1, mock.lastCall(button, "SetAlpha")[1])
				-- The window closed: nothing to read, the button is left as it was (it hides with the window).
				frames.social:Hide()
				local alphas = mock.countCalls(button, "SetAlpha")
				unheard(frames.raid, false)
				GroupPanel.refresh()
				assert.are.equal(alphas, mock.countCalls(button, "SetAlpha"))
				assert.is_true(button:IsShown())
			end)

		it("runs a parent's Show and Hide down to its shown children's OnShow and OnHide (the mock)", function()
			start()
			local parent = _G.CreateFrame("Frame", nil, _G.UIParent)
			local child = _G.CreateFrame("Frame", nil, parent)
			local hidden = _G.CreateFrame("Frame", nil, parent)
			hidden:Hide()
			local seen = {}
			child:SetScript("OnShow", function()
				seen[#seen + 1] = "show"
			end)
			child:SetScript("OnHide", function()
				seen[#seen + 1] = "hide"
			end)
			hidden:SetScript("OnShow", function()
				seen[#seen + 1] = "hidden shown"
			end)
			parent:Hide()
			parent:Show()
			-- Already visible: nothing runs.
			parent:Show()
			assert.are.same({ "hide", "show" }, seen)
			-- A move under a hidden frame hides it; the old parent lets it go.
			local other = _G.CreateFrame("Frame", nil, _G.UIParent)
			other:Hide()
			child:SetParent(other)
			assert.are.same({ "hide", "show", "hide" }, seen)
			parent:Hide()
			parent:Show()
			assert.are.same({ "hide", "show", "hide" }, seen)
			for _, kid in ipairs(parent.children) do
				assert.is_false(rawequal(kid, child))
			end
			other:Show()
			assert.are.same({ "hide", "show", "hide", "show" }, seen)
		end)

		-- Fix round 2, N1: the addon's own sentinel frame, parented to RaidFrame, shows and hides with
		-- the Raid tab (RaidFrame has no rect, but its children still show with it); its own scripts
		-- read the tab again. No hook on a client frame.
		it("sees a tab switch in the open Social window: Friends to Raid wakes the button, Raid to Friends rests it",
			function()
				local frames = start()
				-- Login with the window closed: the button and the sentinel are made.
				frames.social:Hide()
				GroupPanel.refresh()
				local button = _G[GroupPanel.BUTTON_NAME]
				local sentinel = GroupPanel.sentinel
				assert.is_not_nil(sentinel)
				assert.are.equal(frames.raid, sentinel.parent)
				assert.is_function(sentinel:GetScript("OnShow"))
				assert.is_function(sentinel:GetScript("OnHide"))
				-- The window opens on the Friends tab: the button rests.
				frames.raid:Hide()
				frames.social:Show()
				assert.are.equal(0, mock.lastCall(button, "SetAlpha")[1])
				assert.is_false(mock.lastCall(button, "EnableMouse")[1])
				-- The Raid tab clicked: it wakes.
				frames.raid:Show()
				assert.are.equal(1, mock.lastCall(button, "SetAlpha")[1])
				assert.is_true(mock.lastCall(button, "EnableMouse")[1])
				-- Back to Friends: it rests again, its shown flag kept.
				frames.raid:Hide()
				assert.are.equal(0, mock.lastCall(button, "SetAlpha")[1])
				assert.is_false(mock.lastCall(button, "EnableMouse")[1])
				assert.is_true(button:IsShown())
				-- Raid once more.
				frames.raid:Show()
				assert.are.equal(1, mock.lastCall(button, "SetAlpha")[1])
				-- The window closes from the Raid tab: nothing is read, the button hides with it.
				local alphas = mock.countCalls(button, "SetAlpha")
				frames.social:Hide()
				assert.are.equal(alphas, mock.countCalls(button, "SetAlpha"))
				assert.is_false(button:IsVisible())
			end)

		-- Fix round 2, N3: the drawn window made only after login.
		it("moves to a SocialUIFrame made after login: when the window first shows, and on a refresh", function()
			start({ noFrame = true })
			local frames = beforeSocialWindow()
			GroupPanel.refresh()
			local button = _G[GroupPanel.BUTTON_NAME]
			assert.are.equal(frames.friends, button.parent)
			assert.are.equal(frames.raid, GroupPanel.sentinel.parent)
			-- /fs group with nothing shown: the flyout waits on FriendsFrame.
			assert.are.same({ L.groupOpensOnRaidTab }, GroupPanel.command())
			local flyout = _G[GroupPanel.FLYOUT_NAME]
			assert.are.equal(frames.friends, flyout.parent)
			-- The client makes its window, puts FriendsFrame in it and opens it on the Raid tab.
			local social = socialWindow()
			frames.friends:SetParent(social)
			frames.friends:Show()
			assert.are.equal(social, button.parent)
			assertPoint({ "TOPRIGHT", social, "BOTTOMRIGHT", 0, 0 }, button.points[#button.points])
			assert.are.equal(social, flyout.parent)
			assertPoint({ "TOPLEFT", social, "TOPRIGHT", Theme.SIZES.panelOffsetX, Theme.SIZES.panelOffsetY },
				flyout.points[#flyout.points])
			assert.are.equal(1, mock.lastCall(button, "SetAlpha")[1])
			assert.is_true(flyout:IsVisible())
			-- The same window found by a refresh alone (no script told the panel of it).
			start({ noFrame = true })
			frames = beforeSocialWindow()
			GroupPanel.refresh()
			button = _G[GroupPanel.BUTTON_NAME]
			social = socialWindow()
			frames.friends:SetParent(social)
			unheard(frames.friends, true)
			assert.are.equal(frames.friends, button.parent)
			GroupPanel.refresh()
			assert.are.equal(social, button.parent)
			assertPoint({ "TOPRIGHT", social, "BOTTOMRIGHT", 0, 0 }, button.points[#button.points])
		end)

		-- Fix round 2, N4: attach records the new home before SetParent, so a show script the move
		-- itself runs (the client runs it inside SetParent) finds the move made and moves nothing.
		it("moves the button once when the move itself shows it, and the old parent lets it go", function()
			start({ noFrame = true })
			local frames = beforeSocialWindow()
			GroupPanel.refresh()
			local button = _G[GroupPanel.BUTTON_NAME]
			assert.are.equal(frames.friends, button.parent)
			assert.is_false(button:IsVisible())
			-- The drawn window shows while FriendsFrame stays hidden: the move shows the button.
			local social = socialWindow()
			GroupPanel.refresh()
			assert.are.equal(social, button.parent)
			assert.is_true(button:IsVisible())
			assert.are.equal(1, mock.countCalls(button, "SetParent"))
			local held = 0
			for _, kid in ipairs(social.children) do
				if rawequal(kid, button) then
					held = held + 1
				end
			end
			assert.are.equal(1, held)
			for _, kid in ipairs(frames.friends.children) do
				assert.is_false(rawequal(kid, button))
			end
		end)

		it("fits the rows to the Social window's height, not the rectless raid frame's", function()
			local frames = start()
			frames.social:SetHeight(130)
			open()
			-- 130 less the header, the counts, the headers and the padding leaves room for two rows.
			assert.are.equal(GroupPanel.fitRows(130, 0), GroupPanel.last.visible)
			assert.are.equal(2, GroupPanel.last.visible)
		end)

		it("shows one muted line, Not in a group., and nothing else out of a group", function()
			start({ solo = true })
			local flyout = open()
			assert.are.equal(L.groupNotInGroup, GroupPanel.countsLine:GetText())
			assert.are.equal(Theme.HEX.muted, colorOf(GroupPanel.countsLine))
			for _, header in pairs(GroupPanel.headers) do
				assert.is_false(header.button:IsShown())
			end
			for _, row in ipairs(GroupPanel.rowFrames) do
				assert.is_false(row.frame:IsShown())
			end
			assert.is_false(GroupPanel.holder:IsShown())
			assert.is_false(GroupPanel.checkFrame:IsShown())
			assert.are.equal(GroupPanel.soloHeight(), flyout:GetHeight())
			-- A dry run alone still keeps the flyout to its one line.
			GroupPanel.handleReady()
			assert.is_false(GroupPanel.checkFrame:IsShown())
			assert.are.equal(L.groupNotInGroup, GroupPanel.countsLine:GetText())
		end)

		it("draws a row per unit, the player included, Gap high to low, the weakest slot under each", function()
			start()
			itemCache()
			inspectElwen()
			thorgrimHello()
			open()
			assert.are.equal(string.format(L.groupCounts, 3, 2, 2), GroupPanel.countsLine:GetText())
			local thorgrim, elwen, own = drawnRow("Thorgrim"), drawnRow("Elwen"), drawnRow("Tester")
			assert.is_not_nil(thorgrim)
			assert.is_not_nil(own)
			assert.are.equal("Arms", thorgrim.cells.spec:GetText())
			assert.are.equal("12 DPS", elwen.cells.gap:GetText())
			assert.are.equal("Chest: Worn Robe · next Robes of Arcana +12.0 DPS", elwen.sub:GetText())
			assert.are.equal(Theme.HEX.muted, colorOf(elwen.sub))
			-- Rows top to bottom: the biggest gap first.
			local order = {}
			for _, row in ipairs(GroupPanel.rowFrames) do
				if row.frame:IsShown() then
					order[#order + 1] = row.cells.name:GetText()
				end
			end
			assert.are.same({ "Thorgrim", "Elwen", "Tester" }, order)
			for key, header in pairs(GroupPanel.headers) do
				assert.is_true(header.button:IsShown(), key)
			end
			-- A header click sorts and the next turns it round.
			click(GroupPanel.headers.name.button)
			assert.are.same({ key = "name", descending = false }, GroupPanel.sort)
			click(GroupPanel.headers.name.button)
			assert.are.same({ key = "name", descending = true }, GroupPanel.sort)
		end)

		it("dims an offline member's row once", function()
			local offline = { PARTY[1], PARTY[2], { unit = "party2", name = "Elwen", class = "MAGE", className = "Mage",
				level = 58, offline = true } }
			start({ members = offline })
			open()
			assert.are.equal(Theme.ALPHA.dim, mock.lastCall(drawnRow("Elwen").frame, "SetAlpha")[1])
			assert.are.equal(1, mock.lastCall(drawnRow("Thorgrim").frame, "SetAlpha")[1])
		end)

		it("queues an inspect for every member on open and on GROUP_ROSTER_UPDATE, and redraws on INSPECT_READY",
			function()
				start()
				itemCache()
				open()
				-- One NotifyInspect at a time (Inspect's queue); the player is never inspected.
				assert.are.same({ "party1" }, notified)
				assert.are.equal(1, #Inspect.queue)
				state.itemStats["item:70:0"] = { __itemId = 70, __slot = "INVTYPE_CHEST", __level = 66 }
				gear.Thorgrim = { [5] = "item:70:0" }
				local draws = GroupPanel.draws
				local record = Inspect.onEvent("INSPECT_READY", "Player-Thorgrim")
				GroupPanel.onInspectReady(record)
				assert.is_true(GroupPanel.draws > draws)
				assert.are.equal("66", drawnRow("Thorgrim").cells.ilvl:GetText())
				-- A new member arrives: GROUP_ROSTER_UPDATE asks for them too.
				units.party3 = { unit = "party3", name = "Fenna", class = "DRUID", className = "Druid", level = 60 }
				groupList[#groupList + 1] = units.party3
				Inspect.lastAt = nil
				Inspect.pending = nil
				Inspect.queue = {}
				GroupPanel.onRosterUpdate()
				local asked = {}
				for _, entry in ipairs(Inspect.queue) do
					asked[entry.unit] = true
				end
				if Inspect.pending ~= nil then
					asked[Inspect.pending.unit] = true
				end
				assert.is_true(asked.party3 == true)
			end)

		it("redraws an open flyout on a hello and a snapshot!", function()
			start()
			GroupPanel.register()
			open()
			local draws = GroupPanel.draws
			Comms.onEvent("FS1", "1|hello|60|warrior-arms|14|Arms raid|2026-11-04", "GUILD", "Thorgrim-ClassicBetaPvP")
			assert.are.equal(draws + 1, GroupPanel.draws)
			Comms.onEvent("FS1", "1|snapshot!|58|mage-fire|3|Fire|-", "PARTY", "Elwen-ClassicBetaPvP")
			assert.are.equal(draws + 2, GroupPanel.draws)
			assert.are.equal(L.groupAddonLive, drawnRow("Elwen").cells.addon:GetText())
		end)

		it("asks the group for snapshots on open, at most once a minute, never with sharing off", function()
			start()
			open()
			assert.are.same({ { prefix = "FS1", text = "1|snapshot?", channel = "PARTY" } }, sent)
			local button = _G[GroupPanel.BUTTON_NAME]
			click(button)
			clock = clock + 59
			click(button)
			assert.are.equal(1, #sent)
			click(button)
			clock = clock + 1
			click(button)
			assert.are.equal(2, #sent)
			start({ raid = true, sharingOff = true })
			open()
			assert.are.equal(0, #sent)
		end)

		it("never writes into a client frame through any path, and never polls or hooks", function()
			local frames = start({ leader = true })
			itemCache()
			inspectElwen()
			GroupPanel.register()
			open()
			GroupPanel.onRosterUpdate()
			GroupPanel.onReadyCheck()
			GroupPanel.onAura("party1")
			mock.runTimers(state)
			GroupPanel.onReadyCheckFinished()
			click(GroupPanel.headers.score.button)
			GroupPanel.handleFloor("1200")
			-- Tab switches and the window closing and opening (the sentinel's scripts, fix round 2).
			frames.raid:Hide()
			frames.raid:Show()
			frames.social:Hide()
			frames.social:Show()
			click(_G[GroupPanel.BUTTON_NAME])
			assert.are.same({}, state.clientWrites)
			local file = assert(io.open("ForeverSixty/surfaces/GroupPanel.lua", "r"))
			local source = file:read("a")
			file:close()
			for _, call in ipairs({ "HookScript(", "hooksecurefunc(", "PanelTemplates_", '"OnUpdate"', "StaticPopup",
				"ToggleFriendsFrame", "ToggleRaidFrame", "DoReadyCheck", "ShowUIPanel" }) do
				assert.is_nil(source:find(call, 1, true), call)
			end
			-- The scripts are on the addon's own frames only: the button's show script and the
			-- sentinel's show and hide scripts (fix round 2, N1).
			local _, onShows = source:gsub('"OnShow"', "")
			assert.are.equal(2, onShows)
			local _, onHides = source:gsub('"OnHide"', "")
			assert.are.equal(1, onHides)
			assert.is_truthy(source:find('button:SetScript("OnShow"', 1, true))
			assert.is_truthy(source:find('sentinel:SetScript("OnShow"', 1, true))
			assert.is_truthy(source:find('sentinel:SetScript("OnHide"', 1, true))
		end)

		it("watches RaidFrame, FriendsFrame, PartyFrame and every CompactRaidFrame frame as client frames", function()
			for _, name in ipairs({ "RaidFrame", "FriendsFrame", "PartyFrame", "CompactRaidFrameContainer",
				"CompactRaidFrame1" }) do
				assert.is_true(mock.isClientName(name), name)
			end
			assert.is_false(mock.isClientName("ForeverSixtyGroupFlyout"))
		end)

		it("turns itself off on a failure, hides what it drew and says so once in /fs diag", function()
			start()
			GroupPanel.refresh()
			GroupPanel.draw = function()
				error("boom")
			end
			click(_G[GroupPanel.BUTTON_NAME])
			assert.is_true(GroupPanel.disabled)
			assert.is_false(_G[GroupPanel.BUTTON_NAME]:IsShown())
			local notes = Theme.diagnostics()
			assert.is_truthy(notes[#notes]:find(L.diagGroupPanelFailed:match("^[^%%]+"), 1, true))
		end)
	end)

	-- Ruling 3: the ready check.
	describe("the ready check", function()
		--- Every unit's auras for the classic case: the player has Fortitude, Arcane Intellect
		--- (rank 1, by name) and food, but no flask; Thorgrim Fortitude, Battle Shout, food and no
		--- flask; Elwen a flask, food and Arcane Intellect, but no Fortitude.
		local function buffed()
			return {
				Tester = { aura(10938), aura(1459, "Arcane Intellect"), aura(19705) },
				Thorgrim = { aura(10938), aura(11551), aura(19705) },
				Elwen = { aura(17627), aura(19705), aura(10157) },
			}
		end

		it("reads each unit's auras against the instance's needs: who lacks a flask, food, an imbue or a buff",
			function()
				start({ auras = buffed(), instance = { "Molten Core", "raid", 409 } })
				gear.Tester = { [16] = "item:80:0" }
				local result = GroupPanel.onReadyCheck()
				assert.are.same({
					string.format(L.groupMissing, "Tester", "flask, imbue"),
					string.format(L.groupMissing, "Thorgrim", "flask"),
					string.format(L.groupMissing, "Elwen", "Power Word: Fortitude"),
				}, GroupPanel.missingLines(result))
			end)

		it("lists each need of the bosses with who covers it, nobody where nobody does, and the resist", function()
			start({ auras = buffed(), instance = { "Molten Core", "raid", 409 } })
			local result = GroupPanel.onReadyCheck()
			assert.are.same({
				string.format(L.groupCoverage, "Arcane Intellect", "Elwen"),
				string.format(L.groupCoverage, "Power Word: Fortitude", "Tester"),
				string.format(L.groupCoverageNobody, "Mark of the Wild"),
				string.format(L.groupCoverage, "Battle Shout", "Thorgrim"),
				string.format(L.groupCoverage, "Sunder Armor", "Thorgrim"),
				string.format(L.groupCoverageNobody, "Faerie Fire"),
				string.format(L.groupCoverageNobody, string.format(L.groupResist, L.groupSchools.fire)),
			}, GroupPanel.coverageLines(result))
		end)

		it("takes the generic lists outside an instance, the bosses' own inside one", function()
			start()
			local data = GroupPanel.data
			local outside = GroupPanel.needsFor(data, GroupPanel.instanceKey())
			assert.are.same({ 10157, 10938, 9885, 11551 }, outside.buffs)
			assert.are.same({ 11717, 11597, 9907 }, outside.debuffs)
			assert.is_nil(outside.resist)
			assert.are.same({ buffs = { 10157, 10938 }, debuffs = { 11717 }, generic = false },
				GroupPanel.needsFor(data, "Onyxia's Lair"))
			assert.are.equal("Molten Core", GroupPanel.instanceKey("Molten Core", "raid", 409))
			assert.are.equal("Molten Core", GroupPanel.instanceKey("Geschmolzener Kern", "raid", 409))
			assert.is_nil(GroupPanel.instanceKey("Elwynn Forest", "none", 0))
		end)

		it("reads auras through AuraUtil.ForEachAura where the client has it", function()
			start({ auras = buffed(), auraUtil = true, instance = { "Molten Core", "raid", 409 } })
			assert.is_nil(_G.UnitAura)
			local lines = GroupPanel.missingLines(GroupPanel.onReadyCheck())
			assert.are.equal(string.format(L.groupMissing, "Elwen", "Power Word: Fortitude"), lines[3])
		end)

		it("names a need from the addon's own words where the client cannot name the spell, never by id", function()
			start({ auras = buffed(), noSpellNames = true })
			local lines = GroupPanel.coverageLines(GroupPanel.onReadyCheck())
			assert.are.equal(string.format(L.groupCoverage, "Arcane Intellect", "Elwen"), lines[1])
			for _, line in ipairs(lines) do
				assert.is_nil(line:find("%d%d%d%d"), line)
			end
		end)

		it("claims nothing missing where the client cannot read auras", function()
			start({ noAuraApi = true })
			local result = GroupPanel.onReadyCheck()
			assert.are.same({}, GroupPanel.missingLines(result))
			assert.is_true(#GroupPanel.coverageLines(result) > 0)
		end)

		it("gives the chip Ready check: N missing (names) until READY_CHECK_FINISHED and 30 s more", function()
			start({ auras = buffed(), instance = { "Molten Core", "raid", 409 } })
			local refreshes = 0
			Chip.refresh = function()
				refreshes = refreshes + 1
			end
			GroupPanel.onReadyCheck()
			assert.are.equal(1, refreshes)
			local summary = string.format(L.groupChipSummary, 3, "Tester, Thorgrim, Elwen")
			local model = Chip.model({ level = 30, build = { name = "Arms leveling", order = { {}, {}, {} } },
				progress = { matched = 3, total = 3 }, unspent = 0, now = clock, shown = true, combatHide = true,
				raidCheck = Chip.raidCheck })
			assert.are.equal("raidcheck", model.chosen.kind)
			assert.are.equal(summary, model.chosen.text)
			-- The check stands while it runs, past the old 60 s.
			assert.is_true(Chip.raidCheckLive(Chip.raidCheck, clock + 90))
			clock = clock + 35
			GroupPanel.onReadyCheckFinished()
			assert.is_true(Chip.raidCheckLive(Chip.raidCheck, clock + 30))
			assert.is_false(Chip.raidCheckLive(Chip.raidCheck, clock + 31))
			-- A refresh waits for the expiry.
			assert.are.equal(30, state.timers[#state.timers].seconds)
		end)

		it("re-reads the auras on UNIT_AURA while the check stands, throttled: a flask taken leaves the list",
			function()
				start({ auras = buffed(), instance = { "Molten Core", "raid", 409 } })
				gear.Tester = { [16] = "item:80:0" }
				GroupPanel.onReadyCheck()
				table.insert(auras.Thorgrim, aura(17626))
				state.timers = {}
				GroupPanel.onAura("party1")
				GroupPanel.onAura("party1")
				GroupPanel.onAura("target")
				assert.are.equal(1, #state.timers)
				mock.runTimers(state)
				assert.are.same({
					string.format(L.groupMissing, "Tester", "flask, imbue"),
					string.format(L.groupMissing, "Elwen", "Power Word: Fortitude"),
				}, GroupPanel.missingLines(GroupPanel.result))
				assert.are.equal(string.format(L.groupChipSummary, 2, "Tester, Elwen"), Chip.raidCheck.text)
				-- No check standing: an aura change does nothing.
				clock = clock + 35
				GroupPanel.onReadyCheckFinished()
				clock = clock + 31
				state.timers = {}
				GroupPanel.onAura("party1")
				assert.are.equal(0, #state.timers)
			end)

		-- Fix round 2, N2: UNIT_AURA reads the Raid tab only at the throttled re-read, and only when it
		-- changed since the last read (the sentinel hears the tab switches themselves).
		it("reads the Raid tab on UNIT_AURA only after the throttle, and only when it changed", function()
			local frames = start({ auras = buffed(), instance = { "Molten Core", "raid", 409 } })
			GroupPanel.refresh()
			local button = _G[GroupPanel.BUTTON_NAME]
			GroupPanel.onReadyCheck()
			local alphas = mock.countCalls(button, "SetAlpha")
			state.timers = {}
			GroupPanel.onAura("party1")
			GroupPanel.onAura("party2")
			GroupPanel.onAura("player")
			assert.are.equal(alphas, mock.countCalls(button, "SetAlpha"))
			mock.runTimers(state)
			-- The tab is as last read: nothing to do.
			assert.are.equal(alphas, mock.countCalls(button, "SetAlpha"))
			-- Another tab the panel was not told of: the next throttled read rests the button, once.
			unheard(frames.raid, false)
			GroupPanel.onAura("party1")
			assert.are.equal(alphas, mock.countCalls(button, "SetAlpha"))
			mock.runTimers(state)
			assert.are.equal(alphas + 1, mock.countCalls(button, "SetAlpha"))
			assert.are.equal(0, mock.lastCall(button, "SetAlpha")[1])
			GroupPanel.onAura("party1")
			mock.runTimers(state)
			assert.are.equal(alphas + 1, mock.countCalls(button, "SetAlpha"))
			-- No check standing: an aura change reads nothing.
			unheard(frames.raid, true)
			clock = clock + 35
			GroupPanel.onReadyCheckFinished()
			clock = clock + 31
			state.timers = {}
			GroupPanel.onAura("party1")
			mock.runTimers(state)
			assert.are.equal(alphas + 1, mock.countCalls(button, "SetAlpha"))
		end)

		it("draws the check under the rows: the missing list and the coverage list", function()
			start({ auras = buffed(), instance = { "Molten Core", "raid", 409 } })
			open()
			GroupPanel.onReadyCheck()
			assert.is_true(GroupPanel.checkFrame:IsShown())
			local texts = {}
			for _, line in ipairs(GroupPanel.checkLines) do
				if line:IsShown() then
					texts[#texts + 1] = line:GetText()
				end
			end
			assert.are.equal(string.format(L.groupCheckTitle, "12:00"), texts[1])
			assert.are.equal(string.format(L.groupMissing, "Tester", "flask"), texts[2])
			assert.is_truthy(table.concat(texts, "\n"):find(L.groupCoverageTitle, 1, true))
			assert.is_truthy(table.concat(texts, "\n"):find(string.format(L.groupCoverage, "Sunder Armor", "Thorgrim"),
				1, true))
		end)

		it("dry-runs on /fs ready outside a group: the player alone, the lines in chat and the chip", function()
			-- The player is a priest: Fortitude is a buff they can give, so they are asked for it too.
			start({ solo = true, auras = { Tester = { aura(19705), aura(10938) } } })
			local lines = GroupPanel.handleReady()
			assert.are.same({ L.groupReadyDry, string.format(L.groupMissing, "Tester", "flask") }, lines)
			-- Fix round 2: the player alone, the chip names what they lack, not "1 missing (Tester)".
			assert.are.equal(string.format(L.groupChipSummaryYou, "flask"), Chip.raidCheck.text)
			assert.are.equal("Ready check: you are missing flask", Chip.raidCheck.text)
			assert.is_nil(Chip.raidCheck.open)
			start({ solo = true, auras = { Tester = { aura(19705), aura(17627), aura(10938) } } })
			assert.are.same({ L.groupReadyDry, L.groupReadyNone }, GroupPanel.handleReady())
			-- The chip holds the check (fix round 1) but has nothing to show.
			assert.are.same({}, Chip.raidCheck.missing)
			assert.is_false(Chip.raidCheckLive(Chip.raidCheck, clock))
		end)

		-- Fix round 1, D3: a panel that turned itself off never claims a clean check.
		it("says the group panel is off on /fs ready when it has turned itself off, and claims nothing", function()
			start({ solo = true, auras = { Tester = { aura(19705) } } })
			GroupPanel.disabled = true
			assert.are.same({ L.groupPanelOff }, GroupPanel.handleReady())
		end)

		-- Fix round 1, D2: a group buff is the single-target need it stands for.
		it("counts Gift of the Wild, Prayer of Fortitude, Arcane Brilliance and Greater Blessing of Kings as their needs",
			function()
				local raid = {
					{ unit = "player", name = "Tester", class = "PRIEST", className = "Priest", level = 60 },
					{ unit = "party1", name = "Thorgrim", class = "WARRIOR", className = "Warrior", level = 60 },
					{ unit = "party2", name = "Elwen", class = "MAGE", className = "Mage", level = 60 },
					{ unit = "party3", name = "Fenna", class = "DRUID", className = "Druid", level = 60 },
					{ unit = "party4", name = "Brannoc", class = "PALADIN", className = "Paladin", level = 60 },
				}
				local function groupBuffed(extra)
					local list = { aura(17627), aura(19705), aura(11551), aura(21850), aura(450086), aura(23028),
						aura(25898) }
					list[#list + 1] = extra
					return list
				end
				local everyone = {}
				for _, entry in ipairs(raid) do
					everyone[entry.name] = groupBuffed()
				end
				-- One member's group buffs are read by name only (another rank's id).
				everyone.Fenna = { aura(17627), aura(19705), aura(11551), aura(21849, "Gift of the Wild"),
					aura(21562, "Prayer of Fortitude"), aura(99999, "Arcane Brilliance"),
					aura(99998, "Greater Blessing of Kings") }
				local data = groupData()
				data.encounterNeeds[12118].buffs = { 10157, 10938, 9885, 11551, 20217 }
				start({ members = raid, auras = everyone, instance = { "Molten Core", "raid", 409 }, data = nil })
				GroupPanel.data = data
				local result = GroupPanel.onReadyCheck()
				assert.are.same({}, GroupPanel.missingLines(result))
				-- The variants by id, as the table carries them.
				assert.are.same({ 21849, 21850 }, GroupPanel.VARIANTS[9885])
				assert.are.same({ 21562, 21564, 450086 }, GroupPanel.VARIANTS[10938])
				assert.are.same({ 23028 }, GroupPanel.VARIANTS[10157])
				assert.are.same({ 25898 }, GroupPanel.VARIANTS[20217])
				-- Without the group buff, the single-target need is still asked for.
				everyone.Elwen = { aura(17627), aura(19705), aura(11551), aura(21850), aura(23028), aura(25898) }
				result = GroupPanel.onReadyCheck()
				assert.are.same({ string.format(L.groupMissing, "Elwen", "Power Word: Fortitude") },
					GroupPanel.missingLines(result))
			end)

		-- Fix round 1: a check that first found nobody missing still shows when an aura goes.
		it("puts the chip line up when a later aura loss adds someone to a check that found nobody", function()
			local full = {
				Tester = { aura(10938), aura(19705), aura(17627), aura(10157) },
				Thorgrim = { aura(10938), aura(11551), aura(19705), aura(17626) },
				Elwen = { aura(17627), aura(19705), aura(10157), aura(10938) },
			}
			start({ auras = full, instance = { "Molten Core", "raid", 409 } })
			GroupPanel.onReadyCheck()
			assert.are.same({}, GroupPanel.missingLines(GroupPanel.result))
			assert.is_not_nil(Chip.raidCheck)
			assert.is_false(Chip.raidCheckLive(Chip.raidCheck, clock))
			-- Thorgrim's flask runs out while the check stands.
			table.remove(auras.Thorgrim, 4)
			state.timers = {}
			GroupPanel.onAura("party1")
			mock.runTimers(state)
			assert.are.equal(string.format(L.groupChipSummary, 1, "Thorgrim"), Chip.raidCheck.text)
			assert.is_true(Chip.raidCheckLive(Chip.raidCheck, clock))
			local model = Chip.model({ level = 30, build = { name = "Arms leveling", order = { {}, {}, {} } },
				progress = { matched = 3, total = 3 }, unspent = 0, now = clock, shown = true, combatHide = true,
				raidCheck = Chip.raidCheck })
			assert.are.equal("raidcheck", model.chosen.kind)
		end)
	end)

	-- Ruling 4: the floor.
	describe("the score floor", function()
		it("sets Prefs.groupFloor for the leader, clears it with /fs floor, and tells anyone else", function()
			start({ leader = true })
			assert.are.same({ string.format(L.groupFloorSet, 1200) }, GroupPanel.handleFloor("1200"))
			assert.are.equal(1200, Prefs.get("groupFloor"))
			assert.are.same({ L.groupFloorCleared }, GroupPanel.handleFloor(""))
			assert.are.equal(0, Prefs.get("groupFloor"))
			assert.are.same({ L.groupFloorUsage }, GroupPanel.handleFloor("lots"))
			start({ leader = false })
			assert.are.same({ L.groupFloorLeader }, GroupPanel.handleFloor("1200"))
			assert.are.equal(0, Prefs.get("groupFloor"))
			-- Before the group forms, the player sets the floor for the listing.
			start({ solo = true })
			assert.are.same({ string.format(L.groupFloorSet, 900) }, GroupPanel.handleFloor("900"))
		end)

		-- Fix round 1: a client with no UnitIsGroupLeader: a leader only out of a group.
		it("takes the player as leader only out of a group where the client has no UnitIsGroupLeader", function()
			start({ noLeaderApi = true })
			assert.is_nil(_G.UnitIsGroupLeader)
			assert.are.same({ L.groupFloorLeader }, GroupPanel.handleFloor("1200"))
			assert.are.equal(0, Prefs.get("groupFloor"))
			start({ noLeaderApi = true, solo = true })
			assert.are.same({ string.format(L.groupFloorSet, 1200) }, GroupPanel.handleFloor("1200"))
			-- A client that answers one member for a player alone is still solo (IsInGroup).
			start({ noLeaderApi = true, solo = true, count = 1 })
			assert.are.same({ string.format(L.groupFloorSet, 1200) }, GroupPanel.handleFloor("1200"))
		end)

		it("tags rows under the saved floor in the flyout", function()
			start({ leader = true })
			itemCache()
			inspectElwen()
			GroupPanel.handleFloor("1200")
			open()
			local sub = drawnRow("Elwen").sub:GetText()
			assert.are.equal(1, (sub:find(string.format(L.groupBelow, 1200), 1, true)))
			assert.are.equal("", drawnRow("Thorgrim").sub:GetText())
		end)
	end)

	-- Ruling 5: /fs diag.
	describe("/fs diag", function()
		-- Fix round 2: out of a group the count is the player ("1 member (you)"), as the dry run counts.
		it("prints a closed line with the player alone counted and no check outside a group", function()
			start({ solo = true, noFrame = true })
			assert.are.equal(string.format(L.diagGroupPanelYou, L.diagPanelClosed, L.diagGroupNone),
				GroupPanel.diagLine())
			assert.are.equal("Group panel: closed, 1 member (you), ready check none", GroupPanel.diagLine())
		end)

		-- Fix round 2: a dry run alone names what the player lacks, not "1 missing".
		it("says what the player lacks after a dry run alone, and that they lack nothing", function()
			-- A priest with Fortitude on: no flask, no food.
			start({ solo = true, auras = { Tester = { aura(10938) } } })
			open()
			GroupPanel.handleReady()
			assert.are.equal(string.format(L.diagGroupPanelYou, L.diagPanelOpen,
				string.format(L.diagGroupCheckYou, "12:00", "flask, food")), GroupPanel.diagLine())
			assert.are.equal("Group panel: open, 1 member (you), ready check 12:00: you are missing flask, food",
				GroupPanel.diagLine())
			assert.are.equal("Ready check: you are missing flask, food", Chip.raidCheck.text)
			start({ solo = true, auras = { Tester = { aura(10938), aura(19705), aura(17627) } } })
			GroupPanel.handleReady()
			assert.are.equal(string.format(L.diagGroupPanelYou, L.diagPanelClosed,
				string.format(L.diagGroupCheckYouNone, "12:00")), GroupPanel.diagLine())
		end)

		it("counts the members, the inspected and the live, and the last check's time and missing", function()
			start({ auras = {
				Tester = { aura(10938), aura(19705), aura(17627), aura(10157) },
				Thorgrim = { aura(10938), aura(11551), aura(19705) },
				Elwen = { aura(17627), aura(19705), aura(10157), aura(10938) },
			} })
			itemCache()
			inspectElwen()
			thorgrimHello()
			open()
			GroupPanel.onReadyCheck()
			assert.are.equal(string.format(L.diagGroupPanel, L.diagPanelOpen, 3, 2, 2,
				string.format(L.diagGroupCheck, "12:00", 1)), GroupPanel.diagLine())
		end)
	end)

	-- Ruling 6: the README's probes fit the chat box (it cuts /run lines at ~200 characters).
	it("keeps every /run and /dump probe in the group panel's README rows under 180 characters", function()
		local file = assert(io.open("README.md", "r"))
		local count = 0
		for row in file:lines() do
			if row:find("(L5.2)", 1, true) and row:match("^| %d+ |") then
				for probe in row:gmatch("`(/%a+ [^`]*)`") do
					count = count + 1
					assert.is_true(#probe < 180, #probe .. ": " .. probe)
				end
			end
		end
		file:close()
		assert.is_true(count >= 5)
	end)
end)
