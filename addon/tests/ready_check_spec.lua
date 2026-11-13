-- addon/tests/ready_check_spec.lua
-- The ready check (ReadyCheck.lua; the group panel's until lane L7): every unit's auras read against the
-- shipped encounter needs of the instance's bosses, who lacks a flask, food, a weapon buff or a class buff,
-- each need as class presence, the chip's line while the check stands, /fs ready's dry run in chat and the
-- /fs diag line. The check draws nothing: the Officer tab's readiness shows its result.
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


describe("ReadyCheck", function()
	local ReadyCheck, Theme, Comms, Prefs
	local state, clock, sent, notified, combat, units, auras, gear, groupList

	--- The group the client answers. `options`: { members, raid, solo, count, auras, instance, leader,
	--- weaponEnchant, noAuraApi, auraUtil, noSpellNames, data, sharingOff, talents }.
	local function start(options)
		options = options or {}
		-- A second start in one example: the first client is checked and taken down first.
		if state ~= nil then
			assert.are.same({}, state.clientWrites)
			mock.uninstall()
		end
		clock, sent, notified, combat = 1000, {}, {}, false
		groupList = {}
		for _, entry in ipairs(options.solo and (options.members or { PARTY[1] }) or options.members or PARTY) do
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
			talents = options.talents or helper.decidedTalents({ "First", "Second", "Third" }),
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
		helper.load("Inspect")
		helper.load("GroupMembers")
		ReadyCheck = helper.load("ReadyCheck")
		if options.sharingOff then
			Prefs.setFlag("broadcast", false)
		end
		ReadyCheck.data = groupData()
		Comms.data = ReadyCheck.data
		return state
	end

	after_each(function()
		_G.ForeverSixtyData = nil
		_G.ForeverSixtyInbox = nil
		mock.uninstall()
		state = nil
	end)

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
				local result = ReadyCheck.onReadyCheck()
				-- Fix round 4 (E): "<Name>: missing ..."; (A) the player is a priest, who uses no imbue.
				assert.are.same({
					"Tester: missing flask",
					"Thorgrim: missing flask",
					"Elwen: missing Power Word: Fortitude",
				}, ReadyCheck.missingLines(result))
			end)

		-- Fix round 4, A: food from 35 (ReadyCheck.CONSUMABLE_LEVELS); a member below is never flagged
		-- for it, the player's own dry run included. L5.3 pre-item: the flask at level 60, or inside a
		-- raid the data knows (59 outside is not asked, 59 inside Molten Core is, 60 outside is; a 52
		-- in Scholomance, an instance the data has no bosses for, never).
		it("gates the flask at level 60 or a known raid and food at 35: a member below is never flagged", function()
			start()
			assert.are.same({ flask = 60, food = 35 }, ReadyCheck.CONSUMABLE_LEVELS)
			local needs = { buffs = {}, debuffs = {}, generic = true }
			local function lacking(level, where)
				local result = ReadyCheck.check({ { short = "Fenna", classToken = "PRIEST", level = level,
					auras = { ids = {}, names = {} } } }, where or needs, ReadyCheck.spellName)
				return result.missing[1] ~= nil and table.concat(result.missing[1].what, ", ") or "none"
			end
			assert.are.equal("none", lacking(24))
			assert.are.equal("none", lacking(34))
			assert.are.equal("food", lacking(35))
			assert.are.equal("food", lacking(49))
			assert.are.equal("food", lacking(50))
			assert.are.equal("food", lacking(59))
			assert.are.equal("flask, food", lacking(60))
			assert.are.equal("Molten Core", ReadyCheck.needsFor(ReadyCheck.data, "Molten Core").instance)
			local raid = { buffs = {}, debuffs = {}, generic = false, instance = "Molten Core" }
			assert.are.equal("flask, food", lacking(59, raid))
			assert.are.equal("food", lacking(52, { buffs = {}, debuffs = {}, generic = true }))
			assert.is_nil(ReadyCheck.needsFor(ReadyCheck.data, "Scholomance").instance)
			-- A level the client cannot say is never asked inside a raid either.
			assert.are.equal("none", lacking(nil, raid))
			-- A level the client cannot say claims nothing.
			assert.are.equal("none", lacking(nil))
			-- Through the live read: a 59 in Molten Core is asked for a flask, a 59 outside is not.
			start({ members = { { unit = "player", name = "Tester", class = "HUNTER", className = "Hunter", level = 59 } },
				solo = true, auras = { Tester = { aura(19705) } }, instance = { "Molten Core", "raid", 409 } })
			assert.are.same({ "Ready check (only you see this): you are missing flask." }, ReadyCheck.handleReady())
			start({ members = { { unit = "player", name = "Tester", class = "HUNTER", className = "Hunter", level = 59 } },
				solo = true, auras = { Tester = { aura(19705) } } })
			assert.are.same({ "Ready check (only you see this): nothing missing." }, ReadyCheck.handleReady())
			-- The player alone at 24 with nothing on: nothing missing, in chat.
			start({ solo = true, members = { { unit = "player", name = "Tester", class = "HUNTER", className = "Hunter",
				level = 24 } }, auras = { Tester = {} } })
			-- A.2: the chat line says what was checked.
			local said = "Ready check (only you see this): nothing missing at your level \194\183 "
				.. "no flask, food or buff is expected below 35"
			assert.are.same({ said }, ReadyCheck.handleReady())
		end)

		-- L3.7 item 46: the columns lead Name, Level, Spec.
		-- L3.7 item 47: enchants are checked from level 40; below it the cell is blank, the hover says why.
		-- Polish round B fix 1, item 14: with every member under 40 the Enchants header's hover says why
		-- the column is blank.
		-- L3.7 item 46: a dry run is the player's own look; the chip says only they see it. A.2: never
		-- "(test)" in a shipped build, and the chip and the chat line read alike.
		it("words a dry run's chip line as the chat line says it, never as a test, and a real check's as it was",
			function()
			local missing = { { name = "Thorgrim", what = { "flask" } } }
			assert.are.equal("Ready check (only you see this): 1 missing (Thorgrim)",
				ReadyCheck.summary({ missing = missing, dry = true }))
			assert.are.equal("Ready check: 1 missing (Thorgrim)", ReadyCheck.summary({ missing = missing }))
			assert.are.equal("Ready check (only you see this): missing flask", ReadyCheck.summary({ missing = missing,
				alone = true, dry = true }))
			assert.are.equal("Ready check: missing flask", ReadyCheck.summary({ missing = missing, alone = true }))
			assert.are.equal("Ready check (only you see this): nothing missing", ReadyCheck.summary({ missing = {},
				alone = true, dry = true }))
			assert.are.equal("Ready check (only you see this): nothing missing at your level \194\183 "
				.. "no flask, food or buff is expected below 35", ReadyCheck.summary({ missing = {},
				alone = true, dry = true, levelGated = true }))
			for _, key in ipairs({ "groupChipSummaryTest", "groupChipSummaryYouTest", "groupChipSummaryYouNoneTest",
				"groupChipSummaryYouNoneLevelTest", "groupReadyDry", "groupReadyYou", "groupReadyYouNone",
				"groupReadyYouNoneLevel" }) do
				assert.is_nil(L[key]:find("(test)", 1, true), key)
			end
		end)

		-- L3.7 item 48: nothing was asked of the player (every need gated by level): said so.
		it("marks a check where no need applied at the player's level", function()
			start()
			local needs = { buffs = {}, debuffs = {}, generic = true }
			local low = ReadyCheck.check({ { short = "Tester", classToken = "HUNTER", level = 24, own = true,
				spec = "hunter-marksmanship", auras = { ids = {}, names = {} } } }, needs, ReadyCheck.spellName)
			assert.is_true(low.levelGated)
			local at35 = ReadyCheck.check({ { short = "Tester", classToken = "HUNTER", level = 35, own = true,
				spec = "hunter-marksmanship", auras = { ids = { [19705] = true }, names = {} } } }, needs, ReadyCheck.spellName)
			assert.is_falsy(at35.levelGated)
			assert.are.same({}, at35.missing)
			assert.are.equal("Ready check (only you see this): nothing missing at your level \194\183 "
				.. "no flask, food or buff is expected below %d", L.groupReadyYouNoneLevel)
		end)

		-- L3.7 item 48: a rogue with one poisoned hand is missing the other's.
		it("names the off hand when only it lacks a weapon buff", function()
			start()
			local needs = { buffs = {}, debuffs = {}, generic = true }
			local function lack(main, off)
				local result = ReadyCheck.check({ { short = "Tester", classToken = "ROGUE", level = 60, own = true,
					spec = "rogue-combat", imbue = main, imbueOff = off,
					auras = { ids = { [17626] = true, [19705] = true }, names = {} } } }, needs, ReadyCheck.spellName)
				return result.missing[1] and table.concat(result.missing[1].what, ", ") or "none"
			end
			assert.are.equal("weapon buff (off hand)", lack(true, false))
			assert.are.equal("weapon buff", lack(false, false))
			assert.are.equal("weapon buff", lack(false, true))
			assert.are.equal("none", lack(true, true))
			assert.are.equal("none", lack(true, nil))
			assert.are.equal("weapon buff (off hand)", L.groupNeedImbueOff)
		end)

		-- L3.7 item 48: the off-hand fields of GetWeaponEnchantInfo, read only with a weapon in the off hand.
		it("reads the off hand's buff from GetWeaponEnchantInfo's fifth value, when a weapon is worn there", function()
			start()
			local saved = {}
			for _, name in ipairs({ "GetWeaponEnchantInfo", "GetInventoryItemLink", "GetItemInfoInstant" }) do
				saved[name] = _G[name]
			end
			_G.GetWeaponEnchantInfo = function()
				return true, 100, 0, 1, false, 0, 0, 0
			end
			_G.GetInventoryItemLink = function(_, slot)
				return slot == 17 and "item:80:0" or (slot == 16 and "item:81:0" or nil)
			end
			_G.GetItemInfoInstant = function(link)
				if link == "item:80:0" then
					return 80, "Weapon", "One-Handed Axes", "INVTYPE_WEAPONOFFHAND"
				end
				return 81, "Weapon", "Daggers", "INVTYPE_WEAPONMAINHAND"
			end
			assert.is_false(ReadyCheck.ownOffImbue())
			_G.GetWeaponEnchantInfo = function()
				return true, 100, 0, 1, true, 100, 0, 2
			end
			assert.is_true(ReadyCheck.ownOffImbue())
			-- A shield in the off hand has no weapon buff to ask.
			_G.GetItemInfoInstant = function()
				return 80, "Armor", "Shields", "INVTYPE_SHIELD"
			end
			assert.is_nil(ReadyCheck.ownOffImbue())
			for name, value in pairs(saved) do
				_G[name] = value
			end
		end)

		-- Fix round 4, A: an imbue is asked only of a spec that uses one; the spec is the one shown.
		it("asks a weapon buff of warriors, rogues, retribution and enhancement only; never of feral or an unknown spec",
			function()
				start()
				local needs = { buffs = {}, debuffs = {}, generic = true }
				local function asked(spec)
					local result = ReadyCheck.check({ { short = "Fenna", classToken = "WARRIOR", level = 60, spec = spec,
						imbue = false, auras = { ids = { [17626] = true, [19705] = true }, names = {} } } }, needs,
						ReadyCheck.spellName)
					return result.missing[1] ~= nil
				end
				-- Fix round 5, D3: the data's specs, matched exactly (no prefix rule).
				for _, spec in ipairs({ "warrior-arms", "warrior-fury", "warrior-protection", "rogue-assassination",
					"rogue-combat", "rogue-subtlety", "paladin-retribution", "shaman-enhancement" }) do
					assert.is_true(asked(spec), spec)
				end
				for _, spec in ipairs({ "hunter-marksmanship", "hunter-beast-mastery", "mage-arcane", "priest-holy",
					"paladin-holy", "paladin-protection", "shaman-elemental", "shaman-restoration", "druid-balance",
					"druid-restoration", "warlock-affliction", "druid-feral", "druid-feral-bear", "druid-feral-combat", "warrior",
					"rogue", "warrior-", "rogue-x" }) do
					assert.is_false(asked(spec), spec)
				end
				assert.is_false(asked(nil))
			end)

		-- Fix round 5, D3: a bear is never asked for an imbue, nor named in the disclosure.
		it("never asks a druid for a weapon buff and flags no disclosure for one", function()
			start()
			local needs = { buffs = {}, debuffs = {}, generic = true }
			local result = ReadyCheck.check({
				{ short = "Tester", classToken = "PRIEST", level = 60, own = true,
					auras = { ids = { [17627] = true, [19705] = true }, names = {} } },
				{ short = "Brannoc", classToken = "DRUID", level = 60, spec = "druid-feral-bear", imbue = false,
					auras = { ids = { [17626] = true, [19705] = true }, names = {} } },
			}, needs, ReadyCheck.spellName)
			assert.are.same({}, result.missing)
			assert.is_falsy(result.imbueOthers)
			-- L5.3 pre-item: a cat has no weapon buff either: not asked, not named in the disclosure.
			result = ReadyCheck.check({
				{ short = "Brannoc", classToken = "DRUID", level = 60, spec = "druid-feral", imbue = false,
					auras = { ids = { [17626] = true, [19705] = true }, names = {} } },
			}, needs, ReadyCheck.spellName)
			assert.are.same({}, result.missing)
			assert.is_falsy(result.imbueOthers)
			-- A retribution paladin is, and is named.
			result = ReadyCheck.check({
				{ short = "Brannoc", classToken = "PALADIN", level = 60, spec = "paladin-retribution", imbue = false,
					auras = { ids = { [17626] = true, [19705] = true }, names = {} } },
			}, needs, ReadyCheck.spellName)
			assert.is_true(result.imbueOthers)
			assert.are.same({ "Brannoc: missing weapon buff" }, ReadyCheck.missingLines(result))
			assert.are.equal("Weapon buffs are read on yourself only.", L.offImbueSelf)
			assert.are.equal("weapon buff", L.groupNeedImbue)
		end)

		-- Fix round 4, A: the player's own imbue, in the spec their row shows (a warrior here).
		it("asks the player's own imbue in the spec their row shows", function()
			local warrior = { { unit = "player", name = "Tester", class = "WARRIOR", className = "Warrior", level = 60 },
				PARTY[3] }
			start({ members = warrior, auras = { Tester = { aura(17626), aura(19705), aura(11551) },
				Elwen = { aura(17627), aura(19705), aura(10157) } } })
			gear.Tester = { [16] = "item:80:0" }
			assert.are.same({ "Tester: missing weapon buff" }, ReadyCheck.missingLines(ReadyCheck.onReadyCheck()))
		end)

		-- L3.7 item 46: the missing lines name the members in their class colour.
		it("colours the members' names in the missing lines", function()
			local warrior = { { unit = "player", name = "Tester", class = "WARRIOR", className = "Warrior", level = 60 },
				PARTY[3] }
			start({ members = warrior, auras = { Tester = { aura(17626), aura(19705), aura(11551) },
				Elwen = { aura(17627), aura(19705), aura(10157) } } })
			gear.Tester = { [16] = "item:80:0" }
			local line = ReadyCheck.missingLines(ReadyCheck.onReadyCheck(), true)[1]
			assert.are.equal(string.format(L.groupMissing, "|cff" .. Theme.panelClassHex("WARRIOR") .. "Tester|r",
				"weapon buff"), line)
		end)

		-- Fix round 4, A: Forever's Natural flasks (the items' use spells). Fix round 5, D4: the
		-- same-named 1293559-1293562 are counted too, a hedge until a client capture shows which id the
		-- worn aura has.
		it("counts Forever's Natural flasks as a flask, by both sets of their ids", function()
			start()
			for _, id in ipairs({ 1293559, 1293560, 1293561, 1293562, 1293740, 1293741, 1293742, 1293743 }) do
				local result = ReadyCheck.check({ { short = "Fenna", classToken = "PRIEST", level = 60,
					auras = { ids = { [id] = true, [19705] = true }, names = {} } } },
					{ buffs = {}, debuffs = {}, generic = true }, ReadyCheck.spellName)
				assert.are.same({}, result.missing, id)
			end
		end)

		-- Fix round 4, A: the disclosure, once under Missing, where another member uses an imbue.
		-- Fix round 4, C: coverage is class presence: "no <class>" first in the body colour, then the
		-- members of the class, muted, their names class-coloured in the flyout.
		it("lists each need as class presence: no <class> first, then <class> <Name>, and the resist", function()
			start({ auras = buffed(), instance = { "Molten Core", "raid", 409 } })
			local result = ReadyCheck.onReadyCheck()
			assert.are.same({
				"Mark of the Wild: no druid",
				"Faerie Fire: no druid",
				"Fire resistance: no paladin or shaman",
				"Arcane Intellect: mage Elwen",
				"Power Word: Fortitude: priest Tester",
				"Battle Shout: warrior Thorgrim",
				"Sunder Armor: warrior Thorgrim",
			}, ReadyCheck.coverageLines(result))
		end)

		-- Fix round 4, B: outside Molten Core, Onyxia's Lair and Blackwing Lair the check reads the
		-- class buffs and consumables only, and its title names the list.
		it("takes the generic lists outside an instance, the bosses' own inside one", function()
			start()
			local data = ReadyCheck.data
			local outside = ReadyCheck.needsFor(data, ReadyCheck.instanceKey())
			assert.are.same({ 10157, 10938, 9885, 11551 }, outside.buffs)
			assert.are.same({ 11717, 11597, 9907 }, outside.debuffs)
			assert.is_nil(outside.resist)
			assert.are.same({ buffs = { 10157, 10938 }, debuffs = { 11717 }, generic = false, instance = "Onyxia's Lair" },
				ReadyCheck.needsFor(data, "Onyxia's Lair"))
			assert.are.equal("Molten Core", ReadyCheck.instanceKey("Molten Core", "raid", 409))
			assert.are.equal("Molten Core", ReadyCheck.instanceKey("Geschmolzener Kern", "raid", 409))
			assert.is_nil(ReadyCheck.instanceKey("Elwynn Forest", "none", 0))
		end)

		it("reads auras through AuraUtil.ForEachAura where the client has it", function()
			start({ auras = buffed(), auraUtil = true, instance = { "Molten Core", "raid", 409 } })
			assert.is_nil(_G.UnitAura)
			local lines = ReadyCheck.missingLines(ReadyCheck.onReadyCheck())
			assert.are.equal(string.format(L.groupMissing, "Elwen", "Power Word: Fortitude"), lines[3])
		end)

		it("names a need from the addon's own words where the client cannot name the spell, never by id", function()
			start({ auras = buffed(), noSpellNames = true, instance = { "Molten Core", "raid", 409 } })
			local lines = ReadyCheck.coverageLines(ReadyCheck.onReadyCheck())
			assert.are.equal("Arcane Intellect: mage Elwen", lines[4])
			for _, line in ipairs(lines) do
				assert.is_nil(line:find("%d%d%d%d"), line)
			end
		end)

		it("claims nothing missing where the client cannot read auras", function()
			start({ noAuraApi = true, instance = { "Molten Core", "raid", 409 } })
			local result = ReadyCheck.onReadyCheck()
			assert.are.same({}, ReadyCheck.missingLines(result))
			assert.is_true(#ReadyCheck.coverageLines(result) > 0)
		end)

		-- Spec 2026-11-13-addon-hub 6: the ready check left the chip; its summary goes to chat.
		it("says Ready check: N missing (names) in chat and stands until READY_CHECK_FINISHED and 30 s more", function()
			start({ auras = buffed(), instance = { "Molten Core", "raid", 409 } })
			local said, say = {}, Theme.say
			Theme.say = function(text)
				said[#said + 1] = text
			end
			ReadyCheck.onReadyCheck()
			Theme.say = say
			-- Fix round 4 (E): the player first as "you", at most two names, then "+N".
			assert.are.same({ "Ready check: 3 missing (you, Thorgrim, +1)" }, said)
			-- The check stands while it runs, past a dry run's 60 s.
			clock = clock + 90
			assert.is_true(ReadyCheck.stands())
			clock = clock - 55
			ReadyCheck.onReadyCheckFinished()
			clock = clock + 30
			assert.is_true(ReadyCheck.stands())
			clock = clock + 1
			assert.is_false(ReadyCheck.stands())
		end)

		it("re-reads the auras on UNIT_AURA while the check stands, throttled: a flask taken leaves the list",
			function()
				start({ auras = buffed(), instance = { "Molten Core", "raid", 409 } })
				gear.Tester = { [16] = "item:80:0" }
				ReadyCheck.onReadyCheck()
				table.insert(auras.Thorgrim, aura(17626))
				state.timers = {}
				ReadyCheck.onAura("party1")
				ReadyCheck.onAura("party1")
				ReadyCheck.onAura("target")
				assert.are.equal(1, #state.timers)
				mock.runTimers(state)
				assert.are.same({
					"Tester: missing flask",
					"Elwen: missing Power Word: Fortitude",
				}, ReadyCheck.missingLines(ReadyCheck.result))
				assert.are.equal("Ready check: 2 missing (you, Elwen)", ReadyCheck.summary(ReadyCheck.result))
				-- No check standing: an aura change does nothing.
				clock = clock + 35
				ReadyCheck.onReadyCheckFinished()
				clock = clock + 31
				state.timers = {}
				ReadyCheck.onAura("party1")
				assert.are.equal(0, #state.timers)
			end)

		-- Fix round 2, N2: UNIT_AURA reads the Raid tab only at the throttled re-read, and only when it
		-- changed since the last read (the sentinel hears the tab switches themselves).
		-- Fix round 3 (low): a tab never read (nil) and a tab read closed (false) are the same tab.
		it("dry-runs on /fs ready outside a group: the player alone, the lines in chat", function()
			-- The player is a priest: Fortitude is a buff they can give, so they are asked for it too.
			start({ solo = true, auras = { Tester = { aura(19705), aura(10938) } } })
			local lines = ReadyCheck.handleReady()
			-- Fix round 4 (E): alone, one line in chat.
			assert.are.same({ "Ready check (only you see this): you are missing flask." }, lines)
			-- Fix round 2: the player alone, the summary names what they lack, not "1 missing (Tester)".
			assert.are.equal("Ready check (only you see this): missing flask", ReadyCheck.summary(ReadyCheck.result))
			assert.is_nil(ReadyCheck.result.open)
			start({ solo = true, auras = { Tester = { aura(19705), aura(17627), aura(10938) } } })
			assert.are.same({ "Ready check (only you see this): nothing missing." }, ReadyCheck.handleReady())
			-- The dry run stands its 60 s, then goes.
			clock = clock + ReadyCheck.STANDS_DRY
			assert.is_true(ReadyCheck.stands())
			clock = clock + 1
			assert.is_false(ReadyCheck.stands())
		end)

		-- Fix round 4 (E, B): in a group the dry run's chat names the list it read, then a line a member
		-- ("Nobody is missing anything." when clean); a clean real check says nothing in chat.
		it("dry-runs on /fs ready in a group: the header names the list, a line a member; a clean check is quiet",
			function()
				start({ auras = buffed() })
				assert.are.same({ "Ready check (only you see this), standard buffs:", "Tester: missing flask",
					"Thorgrim: missing flask", "Elwen: missing Power Word: Fortitude" }, ReadyCheck.handleReady())
				start({ auras = buffed(), instance = { "Molten Core", "raid", 409 } })
				assert.are.equal("Ready check (only you see this), Molten Core:", ReadyCheck.handleReady()[1])
				start({ auras = {
					Tester = { aura(10938), aura(19705), aura(17627), aura(10157) },
					Thorgrim = { aura(10938), aura(11551), aura(19705), aura(17626) },
					Elwen = { aura(17627), aura(19705), aura(10157), aura(10938) },
				} })
				assert.are.same({ "Ready check (only you see this), standard buffs:", "Nobody is missing anything." },
					ReadyCheck.handleReady())
				-- A real check's clean result says nothing.
				local said, say = {}, Theme.say
				Theme.say = function(text)
					said[#said + 1] = text
				end
				ReadyCheck.onReadyCheck()
				Theme.say = say
				assert.are.same({}, said)
			end)

		-- Fix round 4 (E): the chip's group line: the player first as "you", two names, then "+N".
		it("names the player first as you on the chip, at most two names, then +N", function()
			start()
			local function result(list, own)
				local missing = {}
				for index, name in ipairs(list) do
					missing[index] = { name = name, what = { "flask" }, own = name == own or nil }
				end
				return { missing = missing }
			end
			assert.are.equal("Ready check: 5 missing (you, Thorgrim, +3)",
				ReadyCheck.summary(result({ "Thorgrim", "Elwen", "Tester", "Fenna", "Brannoc" }, "Tester")))
			assert.are.equal("Ready check: 4 missing (Thorgrim, Elwen, +2)",
				ReadyCheck.summary(result({ "Thorgrim", "Elwen", "Fenna", "Brannoc" })))
			assert.are.equal("Ready check: 2 missing (Thorgrim, Elwen)", ReadyCheck.summary(result({ "Thorgrim", "Elwen" })))
			assert.are.equal("Ready check: 2 missing (you, Elwen)",
				ReadyCheck.summary(result({ "Tester", "Elwen" }, "Tester")))
			assert.are.equal("Ready check: 1 missing (Thorgrim)", ReadyCheck.summary(result({ "Thorgrim" })))
			assert.is_nil(ReadyCheck.summary(result({})))
		end)

		-- Fix round 1, D3: a panel that turned itself off never claims a clean check.
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
				ReadyCheck.data = data
				local result = ReadyCheck.onReadyCheck()
				assert.are.same({}, ReadyCheck.missingLines(result))
				-- The variants by id, as the table carries them.
				assert.are.same({ 21849, 21850 }, ReadyCheck.VARIANTS[9885])
				assert.are.same({ 21562, 21564, 450086 }, ReadyCheck.VARIANTS[10938])
				assert.are.same({ 23028 }, ReadyCheck.VARIANTS[10157])
				assert.are.same({ 25898 }, ReadyCheck.VARIANTS[20217])
				-- Without the group buff, the single-target need is still asked for.
				everyone.Elwen = { aura(17627), aura(19705), aura(11551), aura(21850), aura(23028), aura(25898) }
				result = ReadyCheck.onReadyCheck()
				assert.are.same({ string.format(L.groupMissing, "Elwen", "Power Word: Fortitude") },
					ReadyCheck.missingLines(result))
			end)

		-- Fix round 1: a check that first found nobody missing still reads an aura that goes.
		it("reads a later aura loss into a check that found nobody", function()
			local full = {
				Tester = { aura(10938), aura(19705), aura(17627), aura(10157) },
				Thorgrim = { aura(10938), aura(11551), aura(19705), aura(17626) },
				Elwen = { aura(17627), aura(19705), aura(10157), aura(10938) },
			}
			start({ auras = full, instance = { "Molten Core", "raid", 409 } })
			ReadyCheck.onReadyCheck()
			assert.are.same({}, ReadyCheck.missingLines(ReadyCheck.result))
			assert.is_true(ReadyCheck.stands())
			-- Thorgrim's flask runs out while the check stands.
			table.remove(auras.Thorgrim, 4)
			state.timers = {}
			ReadyCheck.onAura("party1")
			mock.runTimers(state)
			assert.are.equal(string.format(L.groupChipSummary, 1, "Thorgrim"), ReadyCheck.summary(ReadyCheck.result))
		end)
	end)

	-- Ruling 4: the floor.

	describe("/fs diag", function()
		it("says none before a check, what the player lacks after a dry run alone, and a group's count", function()
			start({ solo = true, auras = { Tester = { aura(10938) } } })
			assert.are.equal(L.diagReadyNone, ReadyCheck.diagLine())
			ReadyCheck.handleReady()
			assert.are.equal("Ready check 12:00: you are missing flask, food", ReadyCheck.diagLine())
			start({ solo = true, auras = { Tester = { aura(10938), aura(19705), aura(17627) } } })
			ReadyCheck.handleReady()
			assert.are.equal("Ready check 12:00: nothing missing", ReadyCheck.diagLine())
			start({ auras = { Tester = { aura(10938), aura(19705), aura(17627), aura(10157) },
				Thorgrim = { aura(10938), aura(11551), aura(19705) },
				Elwen = { aura(17627), aura(19705), aura(10157), aura(10938) } } })
			ReadyCheck.onReadyCheck()
			assert.are.equal("Ready check 12:00: 1 missing", ReadyCheck.diagLine())
		end)

		it("calls a group's check stale once the player is alone, and a dry run alone stale in a group", function()
			start({ auras = { Tester = { aura(10938) } } })
			ReadyCheck.onReadyCheck()
			local group = ReadyCheck.result
			start({ solo = true })
			ReadyCheck.result = group
			assert.are.equal(L.diagReadyNone, ReadyCheck.diagLine())
			ReadyCheck.handleReady()
			local alone = ReadyCheck.result
			start()
			ReadyCheck.result = alone
			assert.are.equal(L.diagReadyNone, ReadyCheck.diagLine())
		end)
	end)

	it("tells the Officer tab after a check is read and after it ends", function()
		start({ auras = { Tester = { aura(10938) } } })
		local told = 0
		ReadyCheck.onChange = function()
			told = told + 1
		end
		ReadyCheck.onReadyCheck()
		ReadyCheck.onReadyCheckFinished()
		ReadyCheck.handleReady()
		assert.are.equal(3, told)
	end)
end)
