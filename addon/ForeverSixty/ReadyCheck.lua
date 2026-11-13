-- addon/ForeverSixty/ReadyCheck.lua
-- The ready check (spec 2026-11-13-addon-hub 4.6 and 6): who lacks a flask, food, a weapon buff or a class
-- buff, and who covers each need of the encounter. READY_CHECK reads every unit's auras against the shipped
-- encounter needs (Data.encounterNeeds) of the instance's bosses; `/fs ready` dry-runs it and says it in chat.
-- The chip shows its line while it stands, and the Officer tab's Tonight's readiness shows the last check.
-- The check claims only what it can: a flask from level 60 or inside a raid the data knows, food from 35, a
-- weapon buff of a spec that fights with one, read on the player alone; class buffs and consumables only
-- outside a raid the data knows. Nothing here acts: it reads and tells. Moved from the group panel (L7).
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Compat = ns.Compat or require("Compat")
local RaidPrep = ns.RaidPrep or require("RaidPrep")
local PlayerTooltip = ns.PlayerTooltip or require("PlayerTooltip")
local GroupMembers = ns.GroupMembers or require("GroupMembers")
local Chip = ns.Chip or require("Chip")

local ReadyCheck = {}

--- Seconds an aura change waits before the standing ready check is read again: a raid's
--- buffing fires UNIT_AURA by the hundred, and one read covers them all.
ReadyCheck.AURA_DELAY = 0.5

--- The raid instances whose bosses Data.encounterNeeds lists, by the instance id GetInstanceInfo
--- answers eighth (any client language) and by the English name it answers first; each with
--- its bosses' npc ids. The data carries the needs per npc id only (contract 1), so the
--- instance-to-boss map is the addon's own.
ReadyCheck.INSTANCE_IDS = { [409] = "Molten Core", [249] = "Onyxia's Lair", [469] = "Blackwing Lair" }
ReadyCheck.BOSSES = RaidPrep.BOSSES
--- Who can give or put up each need (the encounter's spell ids, contract 1), by class token:
--- the coverage list names the members of these classes.
ReadyCheck.PROVIDERS = {
	[11551] = { "WARRIOR" }, [9885] = { "DRUID" }, [10157] = { "MAGE" }, [10938] = { "PRIEST" },
	[20217] = { "PALADIN" }, [11597] = { "WARRIOR" }, [11717] = { "WARLOCK" }, [1311680] = { "WARLOCK" },
	[9907] = { "DRUID" },
}
--- The group versions of a need (fix round 1, D2), keyed by the need's spell id: a member wearing
--- one has the need, matched by id or by the client's name for it. The ids are this build's
--- (data/builds/<build>/spellranks.json): Gift of the Wild 21849/21850, Prayer of Fortitude
--- 21562/21564/450086, Arcane Brilliance 23028, Greater Blessing of Kings 25898. Battle Shout has
--- none. A data carry: encounterNeeds should carry these (lane L1).
ReadyCheck.VARIANTS = {
	[9885] = { 21849, 21850 },
	[10938] = { 21562, 21564, 450086 },
	[10157] = { 23028 },
	[20217] = { 25898 },
}
--- Who puts up a resistance aura or totem, per school.
ReadyCheck.RESIST_PROVIDERS = {
	fire = { "PALADIN", "SHAMAN" }, frost = { "PALADIN", "SHAMAN" }, shadow = { "PALADIN", "PRIEST" },
	nature = { "SHAMAN", "HUNTER" },
}
--- A class buff asked of only the classes it does anything for: Arcane Intellect of mana users,
--- Battle Shout of those who hit with attack power. Any other buff, of everyone.
ReadyCheck.BUFF_FOR = {
	[10157] = { PALADIN = true, HUNTER = true, PRIEST = true, SHAMAN = true, MAGE = true, WARLOCK = true,
		DRUID = true },
	[11551] = { WARRIOR = true, PALADIN = true, HUNTER = true, ROGUE = true, SHAMAN = true, DRUID = true },
}
--- The flasks' spell ids (Petrification, the Titans, Distilled Wisdom, Supreme Power, Chromatic
--- Resistance), matched by id or by the client's name for them; and the food buff's. Fix round 4
--- (A): Forever's four Natural flasks (Accuracy, Aggression, Precision, Swiftness), the use spells
--- of items 274273-274276 (data/builds/1.60.1.70291/simconsumes.json), as the Titans' 17626 is
--- item 13510's. Fix round 5 (D4): 1293559-1293562, the same-named Alchemy recipes (SkillLineAbility,
--- skill line 171), are kept too as a hedge: the repo has no SpellEffect data to prove which id the
--- worn aura has, and an id no aura has is harmless (the names match by `wears` either way).
ReadyCheck.FLASKS = { 17624, 17626, 17627, 17628, 17629, 1293559, 1293560, 1293561, 1293562, 1293740, 1293741,
	1293742, 1293743 }
ReadyCheck.WELL_FED = 19705
--- Fix round 4 (A): the level each consumable needs; a member below it is never asked for it.
--- L5.3 pre-item: the flask is asked at level 60, or inside a raid instance the data knows
--- (ReadyCheck.INSTANCE_IDS: Molten Core, Onyxia's Lair, Blackwing Lair) at any level the client
--- can say, so a 52 in Scholomance is never flagged.
ReadyCheck.CONSUMABLE_LEVELS = { flask = 60, food = 35 }
--- Fix round 4 (A): the specs that fight with a weapon buff (a temporary imbue): every warrior and
--- rogue spec, Retribution and Enhancement. Fix round 5 (D3): the data's spec slugs, matched
--- exactly (no prefix rule). L5.3 pre-item: no Feral (cat form has no weapon buff). Never a
--- hunter, a caster or a healer; never an unknown spec.
ReadyCheck.IMBUE_SPECS = {
	["warrior-arms"] = true, ["warrior-fury"] = true, ["warrior-protection"] = true,
	["rogue-assassination"] = true, ["rogue-combat"] = true, ["rogue-subtlety"] = true,
	["paladin-retribution"] = true, ["shaman-enhancement"] = true,
}
--- The main hand's inventory slot, for the player's own weapon imbue.
ReadyCheck.MAIN_HAND = 16
--- The off hand's slot, for the player's own off-hand weapon buff (L3.7 item 48).
ReadyCheck.OFF_HAND = 17

--- The generated data table, set by Options.
ReadyCheck.data = nil
--- The last ready check (runCheck's), or nil before one.
ReadyCheck.result = nil
--- Runs after a check was read or ended (the Officer tab redraws its readiness).
ReadyCheck.onChange = nil

local function client(name)
	local fn = _G[name]
	return type(fn) == "function" and fn or nil
end

local function try(fn, ...)
	if type(fn) ~= "function" then
		return nil
	end
	local ok, a, b, c, d, e, f, g, h = pcall(fn, ...)
	if not ok then
		return nil
	end
	return a, b, c, d, e, f, g, h
end

--- The client's name for a spell (C_Spell, else GetSpellInfo), else the addon's own word for the
--- needs it knows (L.groupSpellNames); nil for anything else: an id is never shown (contract 1).
function ReadyCheck.spellName(id)
	local spell = _G.C_Spell
	if type(spell) == "table" then
		local name = try(spell.GetSpellName, id)
		if type(name) == "string" and name ~= "" then
			return name
		end
		local info = try(spell.GetSpellInfo, id)
		if type(info) == "table" and type(info.name) == "string" and info.name ~= "" then
			return info.name
		end
	end
	local name = try(client("GetSpellInfo"), id)
	if type(name) == "string" and name ~= "" then
		return name
	end
	return L.groupSpellNames[id]
end

--- A unit's helpful auras as { ids = { [spellId] = true }, names = { [name] = true } }: through
--- AuraUtil.ForEachAura where the client has it, else UnitAura (UnitBuff on an older client);
--- nil where it can read none (the check then claims nothing for the unit).
function ReadyCheck.aurasOf(unit)
	local auras = { ids = {}, names = {} }
	local function add(name, spellId)
		if type(name) == "string" and name ~= "" then
			auras.names[name] = true
		end
		if type(spellId) == "number" then
			auras.ids[spellId] = true
		end
	end
	local util = _G.AuraUtil
	if type(util) == "table" and type(util.ForEachAura) == "function" then
		local ok = pcall(util.ForEachAura, unit, "HELPFUL", nil, function(aura)
			if type(aura) == "table" then
				add(aura.name, aura.spellId)
			end
			return false
		end, true)
		if ok then
			return auras
		end
	end
	local reader = client("UnitAura") or client("UnitBuff")
	if reader == nil then
		return nil
	end
	for index = 1, 40 do
		local ok, name, _, _, _, _, _, _, _, _, spellId = pcall(reader, unit, index, "HELPFUL")
		if not ok or name == nil then
			break
		end
		add(name, spellId)
	end
	return auras
end

--- Whether the player's main hand carries a temporary imbue (GetWeaponEnchantInfo): true or
--- false with a main-hand item worn, nil where the client cannot say or nothing is worn there.
--- Only the player's own can be read; another player's imbue is never claimed missing.
function ReadyCheck.ownImbue()
	local enchantInfo = client("GetWeaponEnchantInfo")
	if enchantInfo == nil or try(client("GetInventoryItemLink"), "player", ReadyCheck.MAIN_HAND) == nil then
		return nil
	end
	return try(enchantInfo) == true
end

--- Whether the player's off hand carries a temporary buff (GetWeaponEnchantInfo's off-hand fields):
--- true or false with a WEAPON worn in the off hand (a shield or a held item has none to ask), nil
--- where the client cannot say or no weapon is worn there (L3.7 item 48).
function ReadyCheck.ownOffImbue()
	local enchantInfo = client("GetWeaponEnchantInfo")
	local link = try(client("GetInventoryItemLink"), "player", ReadyCheck.OFF_HAND)
	if enchantInfo == nil or link == nil then
		return nil
	end
	local _, _, _, equipLoc = Compat.itemInfoInstant(link)
	if equipLoc ~= "INVTYPE_WEAPON" and equipLoc ~= "INVTYPE_WEAPONOFFHAND" and equipLoc ~= "INVTYPE_WEAPONMAINHAND" then
		return nil
	end
	local info = { try(enchantInfo) }
	if info[1] == nil and info[5] == nil then
		return nil
	end
	return info[5] == true
end

--- The instance's key in ReadyCheck.BOSSES from GetInstanceInfo's name, type and id (read from
--- the client when none is given); nil outdoors or in an instance the data has no bosses for.
function ReadyCheck.instanceKey(name, instanceType, instanceId)
	if name == nil and instanceType == nil then
		local info = { try(client("GetInstanceInfo")) }
		name, instanceType, instanceId = info[1], info[2], info[8]
	end
	if instanceType == nil or instanceType == "none" then
		return nil
	end
	local key = ReadyCheck.INSTANCE_IDS[tonumber(instanceId) or -1]
	if key ~= nil then
		return key
	end
	return ReadyCheck.BOSSES[name] ~= nil and name or nil
end

-- The ready check, pure ----------------------------------------------------------------------

--- The needs to check against (contract 1's encounterNeeds): the union of the bosses' buffs and
--- debuffs of the instance `key` (ReadyCheck.instanceKey's), in the data's order, and the school
--- they stack resistance for; outside one (or one the data has no needs for), the generic lists:
--- every boss's buffs and debuffs, by npc id, and no resistance. Pure.
function ReadyCheck.needsFor(data, key)
	local all = type(data) == "table" and type(data.encounterNeeds) == "table" and data.encounterNeeds or {}
	local npcs = {}
	for _, npc in ipairs(key ~= nil and ReadyCheck.BOSSES[key] or {}) do
		if type(all[npc]) == "table" then
			npcs[#npcs + 1] = npc
		end
	end
	local generic = #npcs == 0
	if generic then
		for npc in pairs(all) do
			npcs[#npcs + 1] = npc
		end
		table.sort(npcs)
	end
	local needs = { buffs = {}, debuffs = {}, generic = generic,
		instance = key ~= nil and ReadyCheck.BOSSES[key] ~= nil and key or nil }
	local taken = {}
	for _, npc in ipairs(npcs) do
		local entry = all[npc]
		for _, kind in ipairs({ "buffs", "debuffs" }) do
			for _, id in ipairs(type(entry[kind]) == "table" and entry[kind] or {}) do
				if not taken[id] then
					taken[id] = true
					needs[kind][#needs[kind] + 1] = id
				end
			end
		end
		if not generic and needs.resist == nil and type(entry.resist) == "string" then
			needs.resist = entry.resist
		end
	end
	return needs
end

local function wears(auras, id, name)
	return auras.ids[id] == true or (name ~= nil and auras.names[name] == true)
end

local function wearsAny(auras, ids, nameOf)
	for _, id in ipairs(ids) do
		if wears(auras, id, nameOf(id)) then
			return true
		end
	end
	return false
end

--- A need worn: its own spell (any rank, by name) or one of its group versions (D2).
local function wearsNeed(auras, id, name, nameOf)
	return wears(auras, id, name) or wearsAny(auras, ReadyCheck.VARIANTS[id] or {}, nameOf)
end

--- Whether `spec` fights with a weapon imbue (ReadyCheck.IMBUE_SPECS, exactly); false for nil. Pure.
function ReadyCheck.usesImbue(spec)
	return type(spec) == "string" and ReadyCheck.IMBUE_SPECS[spec] == true
end

--- Whether a unit at `level` is old enough for a consumable of `need` levels: a level the client
--- cannot say (nil, 0) is not; a unit too high to show (-1) is.
local function reaches(level, need)
	level = tonumber(level)
	return level ~= nil and level ~= 0 and (level < 0 or level >= need)
end

--- The ready check (ruling 3) over `members` ({ short, classToken, level, spec (the one their row
--- shows), own (the player), connected, auras (aurasOf's; nil: unread), imbue (the player's:
--- ReadyCheck.ownImbue's) }) against `needs` (ReadyCheck.needsFor's), naming spells through `nameOf`
--- (ReadyCheck.spellName); a group version of a buff (ReadyCheck.VARIANTS) counts as the buff.
--- Answers { missing = { { name, classToken, own, what = { "flask", "Arcane Intellect", ... } } },
--- coverage = { { need, classes, by = { { name, classToken } } } }, generic, imbueOthers }: a member
--- lacks a flask (from level 50) or the food buff (from 35) (fix round 4, A:
--- ReadyCheck.CONSUMABLE_LEVELS), the player's own imbue where their spec uses one, or a class buff
--- the group can give (a member of its class is here) and does them some good; a member offline
--- or whose auras cannot be read is never said to lack anything. Fix round 4 (B): with the generic
--- lists (no raid the data knows) there is no coverage. `imbueOthers`: a member other than the
--- player uses an imbue (only the player's can be read). Pure.
function ReadyCheck.check(members, needs, nameOf)
	local present = {}
	for _, member in ipairs(members) do
		if type(member.classToken) == "string" then
			present[member.classToken] = present[member.classToken] or {}
			table.insert(present[member.classToken], member.short)
		end
	end
	local function providers(classes)
		local found = {}
		for _, token in ipairs(classes or {}) do
			for _, name in ipairs(present[token] or {}) do
				found[#found + 1] = { name = name, classToken = token }
			end
		end
		return found
	end
	local result = { missing = {}, coverage = {}, generic = needs.generic == true or nil }
	local levels = ReadyCheck.CONSUMABLE_LEVELS
	for _, member in ipairs(members) do
		if member.own ~= true and ReadyCheck.usesImbue(member.spec) then
			result.imbueOthers = true
		end
		local auras = member.auras
		if member.connected ~= false and auras ~= nil then
			local what = {}
			-- How many needs applied to this member at all (L3.7 item 48): none, and a clean result
			-- reads "nothing missing at your level".
			local asked = 0
			-- L5.3 pre-item: level 60, or inside a raid the data knows (any level the client says).
			local flaskLevel = needs.instance ~= nil and 1 or levels.flask
			if reaches(member.level, flaskLevel) then
				asked = asked + 1
				if not wearsAny(auras, ReadyCheck.FLASKS, nameOf) then
					what[#what + 1] = L.groupNeedFlask
				end
			end
			if reaches(member.level, levels.food) then
				asked = asked + 1
				if not wears(auras, ReadyCheck.WELL_FED, nameOf(ReadyCheck.WELL_FED)) and not auras.names[L.groupWellFed] then
					what[#what + 1] = L.groupNeedFood
				end
			end
			if member.imbue ~= nil and ReadyCheck.usesImbue(member.spec) then
				asked = asked + 1
			end
			if member.imbue == false and ReadyCheck.usesImbue(member.spec) then
				what[#what + 1] = L.groupNeedImbue
			elseif member.imbueOff == false and ReadyCheck.usesImbue(member.spec) then
				-- The main hand is buffed, the off hand (a second weapon) is not (L3.7 item 48).
				what[#what + 1] = L.groupNeedImbueOff
			end
			for _, id in ipairs(needs.buffs or {}) do
				local name = nameOf(id)
				local forClass = ReadyCheck.BUFF_FOR[id]
				if name ~= nil and (forClass == nil or forClass[member.classToken])
					and #providers(ReadyCheck.PROVIDERS[id]) > 0 then
					asked = asked + 1
					if not wearsNeed(auras, id, name, nameOf) then
						what[#what + 1] = name
					end
				end
			end
			if member.own == true then
				result.asked = (result.asked or 0) + asked
				result.ownRead = true
			end
			if #what > 0 then
				result.missing[#result.missing + 1] = { name = member.short, classToken = member.classToken,
					own = member.own == true or nil, what = what }
			end
		end
	end
	-- Nothing applied to the player (every need gated by their level): a clean result says "at your level".
	result.levelGated = result.ownRead == true and (result.asked or 0) == 0 and #result.missing == 0 or nil
	if needs.generic then
		return result
	end
	-- A need is listed with the classes that give it; one no class is known to give says nothing.
	local function cover(need, classes)
		if need ~= nil and type(classes) == "table" and #classes > 0 then
			result.coverage[#result.coverage + 1] = { need = need, classes = classes, by = providers(classes) }
		end
	end
	for _, kind in ipairs({ "buffs", "debuffs" }) do
		for _, id in ipairs(needs[kind] or {}) do
			cover(nameOf(id), ReadyCheck.PROVIDERS[id])
		end
	end
	if needs.resist ~= nil then
		cover(string.format(L.groupResist, L.groupSchools[needs.resist] or needs.resist),
			ReadyCheck.RESIST_PROVIDERS[needs.resist])
	end
	return result
end

--- A member's name, in their class colour lifted for the panel where `coloured`.
local function nameText(name, classToken, coloured)
	if coloured and type(classToken) == "string" then
		return "|cff" .. Theme.panelClassHex(classToken) .. name .. "|r"
	end
	return name
end

--- A class's word ("druid"), by token.
local function classWord(token)
	return L.groupClassWords[token] or tostring(token):lower()
end

--- "Thorgrim: missing flask, imbue", one per member lacking anything (fix round 4, E); the names
--- class-coloured where `coloured`. Pure.
function ReadyCheck.missingLines(result, coloured)
	local lines = {}
	for _, entry in ipairs(type(result) == "table" and result.missing or {}) do
		lines[#lines + 1] = string.format(L.groupMissing, nameText(entry.name, entry.classToken, coloured),
			table.concat(entry.what, L.groupMissingJoin))
	end
	return lines
end

--- The coverage lines as class presence (fix round 4, C): each need nobody of its classes covers,
--- "Faerie Fire: no druid" ("Fire resistance: no paladin or shaman"), first; then each covered
--- need, "Sunder Armor: warrior Thorgrim", the members' names class-coloured where `coloured`.
--- Each { text, covered }. Pure.
function ReadyCheck.coverageEntries(result, coloured)
	local absent, covered = {}, {}
	for _, entry in ipairs(type(result) == "table" and result.coverage or {}) do
		if #entry.by == 0 then
			local words = {}
			for index, token in ipairs(entry.classes or {}) do
				words[index] = classWord(token)
			end
			absent[#absent + 1] = { text = string.format(L.groupCoverageNo, entry.need,
				table.concat(words, L.groupCoverageOr)), covered = false }
		else
			local names = {}
			for index, by in ipairs(entry.by) do
				names[index] = string.format(L.groupCoverageMember, classWord(by.classToken),
					nameText(by.name, by.classToken, coloured))
			end
			covered[#covered + 1] = { text = string.format(L.groupCoverageBy, entry.need,
				table.concat(names, L.groupCoverageJoin)), covered = true }
		end
	end
	for _, entry in ipairs(covered) do
		absent[#absent + 1] = entry
	end
	return absent
end

--- The coverage lines' plain text, in ReadyCheck.coverageEntries' order. Pure.
function ReadyCheck.coverageLines(result)
	local lines = {}
	for index, entry in ipairs(ReadyCheck.coverageEntries(result, false)) do
		lines[index] = entry.text
	end
	return lines
end

--- What the player lacks when the check read them alone (`result.alone`, fix round 2): "flask,
--- food"; nil when the check read a group or found nothing. Pure.
function ReadyCheck.ownLack(result)
	if type(result) ~= "table" or not result.alone then
		return nil
	end
	local entry = result.missing[1]
	if entry == nil then
		return nil
	end
	return table.concat(entry.what, L.groupMissingJoin)
end

--- The chip's one line: "Ready check: 5 missing (you, Thorgrim, +3)" (fix round 4, E: the player
--- first as "you" when among them, at most two names, then "+N"); for the player alone (fix round
--- 2; fix round 3, C3: the chip's short form) "Ready check: missing flask, food", and after a clean
--- dry run alone "Ready check: nothing missing" (fix round 4); nil with nobody missing otherwise.
--- Pure.
function ReadyCheck.summary(result)
	if type(result) ~= "table" then
		return nil
	end
	local missing = type(result.missing) == "table" and result.missing or {}
	local dry = result.dry == true
	if #missing == 0 then
		if not (result.alone and dry) then
			return nil
		end
		if result.levelGated then
			return string.format(L.groupChipSummaryYouNoneLevelTest, ReadyCheck.CONSUMABLE_LEVELS.food)
		end
		return L.groupChipSummaryYouNoneTest
	end
	local lack = ReadyCheck.ownLack(result)
	if lack ~= nil then
		return string.format(dry and L.groupChipSummaryYouTest or L.groupChipSummaryYou, lack)
	end
	local names = {}
	for _, entry in ipairs(missing) do
		if entry.own then
			table.insert(names, 1, L.groupChipYou)
		else
			names[#names + 1] = entry.name
		end
	end
	local shown = { names[1], names[2] }
	if #names > 2 then
		shown[3] = string.format(L.groupChipMore, #names - 2)
	end
	return string.format(dry and L.groupChipSummaryTest or L.groupChipSummary, #missing,
		table.concat(shown, L.groupChipJoin))
end


-- The ready check, read ------------------------------------------------------------------------

--- Whether the last check still stands at `now`: running, or within its time after it finished
--- (a dry run: Chip.RAID_CHECK_SECONDS).
local function checkStands(now)
	local result = ReadyCheck.result
	if result == nil then
		return false
	end
	if result.open then
		return now - result.at <= Chip.RAID_CHECK_OPEN
	end
	return now <= (result.untilAt or result.at + Chip.RAID_CHECK_SECONDS)
end

--- Whether a check found nothing on a dry run of the player alone: the chip then says so (fix
--- round 4, E); a real check's clean result shows nothing.
local function cleanAlone(result)
	return result.alone == true and result.dry == true and #result.missing == 0
end

--- Reads every unit's auras (the player alone out of a group) and checks them against the
--- instance's needs; the chip gets the summary. `again`: the same
--- check read anew (its time kept). Fix round 4 (A): each member's level and shown spec go in;
--- (B) `where` names the raid whose needs were read (nil: the generic lists). Answers the result.
local function runCheck(dry, again)
	local units = GroupMembers.groupUnits()
	local alone = #units == 0
	if alone then
		units = { "player" }
	end
	local data = ReadyCheck.data or ns.Data
	local _, ownRealm = Theme.playerNameRealm()
	local now = Theme.now()
	local members = {}
	for _, unit in ipairs(units) do
		local name, realm = PlayerTooltip.nameOf(unit)
		if type(name) == "string" and name ~= "" then
			local _, classToken = try(client("UnitClass"), unit)
			local member = {
				short = name, classToken = classToken, connected = try(client("UnitIsConnected"), unit) ~= false,
				auras = ReadyCheck.aurasOf(unit), level = tonumber((try(client("UnitLevel"), unit))),
			}
			if GroupMembers.isPlayer(unit) then
				member.own = true
				member.imbue = ReadyCheck.ownImbue()
				member.imbueOff = ReadyCheck.ownOffImbue()
				local ok, spec = pcall(GroupMembers.ownSpec, data)
				member.spec = ok and GroupMembers.validSpec(spec) or nil
			else
				local ok, spec = pcall(GroupMembers.shownSpec, unit, name, realm, ownRealm, now)
				member.spec = ok and spec or nil
			end
			members[#members + 1] = member
		end
	end
	local key = ReadyCheck.instanceKey()
	local needs = ReadyCheck.needsFor(data, key)
	local result = ReadyCheck.check(members, needs, ReadyCheck.spellName)
	-- Fix round 2: the player alone, so the summaries say what they lack (ReadyCheck.summary).
	result.alone = alone or nil
	result.where = not needs.generic and key or nil
	local previous = ReadyCheck.result
	if again and previous ~= nil then
		result.at, result.clock, result.open, result.untilAt, result.dry = previous.at, previous.clock, previous.open,
			previous.untilAt, previous.dry
		ReadyCheck.result = result
		Chip.updateRaidCheck(ReadyCheck.missingLines(result, true), ReadyCheck.summary(result), cleanAlone(result))
	else
		result.at, result.clock, result.open, result.dry = Theme.now(), date("%H:%M"), not dry or nil, dry or nil
		ReadyCheck.result = result
		Chip.setRaidCheck(ReadyCheck.missingLines(result, true), ReadyCheck.summary(result), not dry, cleanAlone(result))
	end
	return result
end

local function changed()
	if ReadyCheck.onChange ~= nil then
		pcall(ReadyCheck.onChange)
	end
end

--- READY_CHECK: every unit's auras read against the instance's needs; the chip gets its line until
--- READY_CHECK_FINISHED and 30 s more. Answers the result.
function ReadyCheck.onReadyCheck()
	local result = runCheck(false)
	changed()
	return result
end

--- READY_CHECK_FINISHED: the check stands 30 s more (Chip.finishRaidCheck), then goes from the chip.
function ReadyCheck.onReadyCheckFinished()
	local result = ReadyCheck.result
	if result ~= nil and result.open then
		result.open = nil
		result.untilAt = Theme.now() + Chip.RAID_CHECK_AFTER
	end
	Chip.finishRaidCheck()
	changed()
end

--- UNIT_AURA (throttled): while a check stands, a group unit's aura change has its auras read again
--- ReadyCheck.AURA_DELAY later, once for every change in between.
function ReadyCheck.onAura(unit)
	if type(unit) ~= "string" or not (unit == "player" or unit:match("^party%d+$") or unit:match("^raid%d+$")) then
		return
	end
	if ReadyCheck.auraPending or not checkStands(Theme.now()) then
		return
	end
	ReadyCheck.auraPending = true
	local function again()
		ReadyCheck.auraPending = nil
		if checkStands(Theme.now()) then
			runCheck(ReadyCheck.result.dry, true)
			changed()
		end
	end
	if not Theme.after(ReadyCheck.AURA_DELAY, again) then
		again()
	end
end

--- Whether the last check still stands now (the Officer tab shows it while it does, and after it as the
--- night's last check).
function ReadyCheck.stands()
	return checkStands(Theme.now())
end

--- /fs ready: a dry run of the check (the player alone out of a group), told in chat; the chip shows its
--- line for Chip.RAID_CHECK_SECONDS. Alone, one line ("Ready check (only you see this): you are missing
--- flask, food." or "...: nothing missing."); in a group a header naming the list read ("Ready check (only
--- you see this), Molten Core:", "..., standard buffs:"), then a line a member or "Nobody is missing
--- anything.". The lines to print.
function ReadyCheck.handleReady()
	local result = runCheck(true)
	changed()
	if result.alone then
		local lack = ReadyCheck.ownLack(result)
		return { lack ~= nil and string.format(L.groupReadyYou, lack)
			or (result.levelGated and string.format(L.groupReadyYouNoneLevel, ReadyCheck.CONSUMABLE_LEVELS.food)
				or L.groupReadyYouNone) }
	end
	local lines = { string.format(L.groupReadyDry, result.where or L.groupCheckStandard) }
	local missing = ReadyCheck.missingLines(result)
	if #missing == 0 then
		lines[2] = L.groupReadyNone
	end
	for _, line in ipairs(missing) do
		lines[#lines + 1] = line
	end
	return lines
end

--- /fs diag's line: "Ready check: none", "Ready check 21:04: 2 missing", or for a check that read the
--- player alone "Ready check 19:43: you are missing flask, food" ("nothing missing" when clean). A check
--- that read the player alone is stale once they are in a group, and a group's once they are alone.
function ReadyCheck.diagLine()
	local result = ReadyCheck.result
	local inGroup = GroupMembers.inGroup()
	if result ~= nil and result.alone and not inGroup then
		local lack = ReadyCheck.ownLack(result)
		return lack ~= nil and string.format(L.diagReadyYou, result.clock or "", lack)
			or string.format(L.diagReadyYouNone, result.clock or "")
	elseif result ~= nil and not result.alone and inGroup then
		return string.format(L.diagReadyGroup, result.clock or "", #result.missing)
	end
	return L.diagReadyNone
end

ns.ReadyCheck = ReadyCheck
return ReadyCheck
