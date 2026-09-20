-- addon/ForeverSixty/Codec.lua
-- The two string formats, both directions, pure: this file touches no WoW API
-- and no addon state, so busted runs it as ordinary Lua.
--
--   FS1  (addon -> site)  the character export; the site's own
--                         web/src/lib/planner/fs1.ts is the reference
--                         implementation and the shared fixture vectors are
--                         generated from it, so the two cannot drift.
--   FSB1 (site -> addon)  the build code; Follow and Gear read it. See
--                         Codec.decodeFSB1 (added in its own task).
--
-- Every refusal names its reason, from ns.L. "That code is FS2, this addon
-- reads FS1" tells someone what to do; "invalid code" does not.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")

local Codec = {}

--- The site's SLOTS order (web/src/lib/planner/types.ts). Gear encodes in
--- this order regardless of the order a caller supplies, so one build has
--- exactly one string.
Codec.SLOTS = {
	"head", "neck", "shoulder", "back", "chest", "wrist", "hands", "waist",
	"legs", "feet", "finger1", "finger2", "trinket1", "trinket2",
	"main_hand", "off_hand", "ranged",
}

local SLOT_SET = {}
for _, slot in ipairs(Codec.SLOTS) do
	SLOT_SET[slot] = true
end

Codec.FS1_PREFIX = "FS1"
Codec.FSB1_PREFIX = "FSB1"
--- Matches the site's MAX_CODE_LENGTH: a full bank is several thousand
--- characters, and this is checked before any splitting so a hostile string
--- does no real work.
Codec.MAX_CODE_LENGTH = 16384

local TREES = 3
local DIGITS = "0123456789abcdefghijklmnopqrstuvwxyz"

local function toBase36(value)
	if value < 0 then
		value = 0
	elseif value > 35 then
		value = 35
	end
	return DIGITS:sub(value + 1, value + 1)
end

local function fromBase36(char)
	local index = DIGITS:find(char, 1, true)
	if index == nil then
		local upper = DIGITS:find(char:lower(), 1, true)
		index = upper
	end
	return index and index - 1 or nil
end

--- Split on one literal separator, keeping empty fields. Lua has no split.
local function split(text, separator)
	local parts, position = {}, 1
	while true do
		local start, stop = text:find(separator, position, true)
		if start == nil then
			parts[#parts + 1] = text:sub(position)
			return parts
		end
		parts[#parts + 1] = text:sub(position, start - 1)
		position = stop + 1
	end
end

local function isDigits(text)
	return text ~= "" and text:match("^%d+$") ~= nil
end

local function urlEncode(text)
	return (text:gsub("[^%w%-%._~]", function(char)
		return string.format("%%%02X", char:byte())
	end))
end

local function urlDecode(text)
	-- A malformed escape reads as itself rather than erroring: a set name a
	-- third-party addon wrote without encoding is an honest string the player
	-- typed, not something to refuse the whole code over. This is the same
	-- call the site's decodeName makes.
	return (text:gsub("%%(%x%x)", function(hex)
		return string.char(tonumber(hex, 16))
	end))
end

-- ---------------------------------------------------------------- encode --

local function encodeTree(ranks)
	local digits = {}
	for index = 1, #ranks do
		digits[index] = toBase36(math.floor(ranks[index] + 0.5))
	end
	local text = table.concat(digits):gsub("0+$", "")
	return text == "" and "0" or text
end

local function encodeTrees(treeRanks)
	local fields = {}
	for index = 1, TREES do
		fields[index] = encodeTree(treeRanks[index] or {})
	end
	return table.concat(fields, "/")
end

local function encodeItemParts(itemId, enchant, suffix)
	if suffix then
		return string.format("%d:%d:%d", itemId, enchant or 0, suffix)
	elseif enchant then
		return string.format("%d:%d", itemId, enchant)
	end
	return tostring(itemId)
end

local function encodeGearSlots(slots)
	local bySlot = {}
	for _, entry in ipairs(slots or {}) do
		bySlot[entry.slot] = entry
	end
	local parts = {}
	for _, slot in ipairs(Codec.SLOTS) do
		local entry = bySlot[slot]
		if entry then
			parts[#parts + 1] = slot .. "=" .. encodeItemParts(entry.itemId, entry.enchant, entry.suffix)
		end
	end
	return table.concat(parts, ",")
end

local function encodeItems(items)
	local parts = {}
	for index, item in ipairs(items or {}) do
		parts[index] = encodeItemParts(item.itemId, item.enchant, item.suffix)
	end
	return table.concat(parts, ",")
end

--- FS1 version 2: the version-1 head, then the sections in the contract's
--- fixed order, each omitted when empty.
function Codec.encodeFS1(build)
	local head = table.concat({
		Codec.FS1_PREFIX,
		build.dataBuild,
		build.classSlug,
		build.raceSlug,
		encodeTrees(build.treeRanks or {}),
		encodeGearSlots(build.gearSlots),
	}, ":")

	local sections = {}
	if build.bags and #build.bags > 0 then
		sections[#sections + 1] = "bags=" .. encodeItems(build.bags)
	end
	if build.bank and #build.bank > 0 then
		sections[#sections + 1] = "bank=" .. encodeItems(build.bank)
	end
	if build.sets and #build.sets > 0 then
		local entries = {}
		for index, set in ipairs(build.sets) do
			entries[index] = urlEncode(set.name) .. "=" .. encodeGearSlots(set.gear)
		end
		sections[#sections + 1] = "sets=" .. table.concat(entries, ";")
	end
	if build.loadouts and #build.loadouts > 0 then
		local entries = {}
		for index, loadout in ipairs(build.loadouts) do
			entries[index] = urlEncode(loadout.name) .. "=" .. encodeTrees(loadout.treeRanks)
		end
		sections[#sections + 1] = "loadouts=" .. table.concat(entries, ";")
	end
	if build.professions and #build.professions > 0 then
		sections[#sections + 1] = "professions=" .. table.concat(build.professions, ",")
	end

	if #sections == 0 then
		return head
	end
	return head .. "|" .. table.concat(sections, "|")
end

-- ---------------------------------------------------------------- decode --

local function parseTrees(field)
	local treeStrings = split(field, "/")
	if #treeStrings ~= TREES then
		return nil, string.format(L.codecTrees, #treeStrings, TREES)
	end
	local treeRanks = {}
	for index, tree in ipairs(treeStrings) do
		local ranks = {}
		for position = 1, #tree do
			local digit = tree:sub(position, position)
			local rank = fromBase36(digit)
			if rank == nil then
				return nil, string.format(L.codecRank, digit)
			end
			ranks[position] = rank
		end
		treeRanks[index] = ranks
	end
	return treeRanks
end

--- `item_id[:enchant[:suffix]]`, digits only in every part.
local function parseItemParts(value, message, entry)
	local parts = split(value, ":")
	if #parts > 3 then
		return nil, string.format(message, entry)
	end
	for _, part in ipairs(parts) do
		if not isDigits(part) then
			return nil, string.format(message, entry)
		end
	end
	return {
		itemId = tonumber(parts[1]),
		enchant = parts[2] and tonumber(parts[2]) or nil,
		suffix = parts[3] and tonumber(parts[3]) or nil,
	}
end

local function parseGearList(field)
	local slots = {}
	if field == "" then
		return slots
	end
	for _, entry in ipairs(split(field, ",")) do
		local pieces = split(entry, "=")
		if #pieces ~= 2 then
			return nil, string.format(L.codecGearEntry, entry)
		end
		local slot, value = pieces[1], pieces[2]
		if not SLOT_SET[slot] then
			return nil, string.format(L.codecSlot, slot)
		end
		local item, message = parseItemParts(value, L.codecGearEntry, entry)
		if item == nil then
			return nil, message
		end
		item.slot = slot
		slots[#slots + 1] = item
	end
	return slots
end

local function parseItemList(field)
	local items = {}
	if field == "" then
		return items
	end
	for _, entry in ipairs(split(field, ",")) do
		local item, message = parseItemParts(entry, L.codecItemEntry, entry)
		if item == nil then
			return nil, message
		end
		items[#items + 1] = item
	end
	return items
end

--- `<name>=<payload>;…`, the name url-encoded so it may carry ; = and |.
local function parseNamed(field)
	local entries = {}
	if field == "" then
		return entries
	end
	for _, entry in ipairs(split(field, ";")) do
		local at = entry:find("=", 1, true)
		if at == nil then
			entries[#entries + 1] = { name = urlDecode(entry), payload = "" }
		else
			entries[#entries + 1] = {
				name = urlDecode(entry:sub(1, at - 1)),
				payload = entry:sub(at + 1),
			}
		end
	end
	return entries
end

function Codec.decodeFS1(code)
	if #code > Codec.MAX_CODE_LENGTH then
		return nil, L.codecTooLong
	end
	local trimmed = code:match("^%s*(.-)%s*$")
	local pipes = split(trimmed, "|")
	local head = table.remove(pipes, 1)

	local parts = split(head, ":")
	if parts[1] ~= Codec.FS1_PREFIX then
		local named = (parts[1] == nil or parts[1] == "") and L.codecUnlabelled or parts[1]
		return nil, string.format(L.codecWrongPrefix, named, Codec.FS1_PREFIX)
	end
	if #parts < 6 then
		return nil, L.codecShort
	end

	local treeRanks, message = parseTrees(parts[5])
	if treeRanks == nil then
		return nil, message
	end
	-- Everything from field 6 on is the gear list: an item id never contains a
	-- colon, but rejoining is what keeps the field count from mattering.
	local gearField = table.concat(parts, ":", 6)
	local gearSlots
	gearSlots, message = parseGearList(gearField)
	if gearSlots == nil then
		return nil, message
	end

	local build = {
		dataBuild = parts[2],
		classSlug = parts[3],
		raceSlug = parts[4],
		treeRanks = treeRanks,
		gearSlots = gearSlots,
		bags = {},
		bank = {},
		sets = {},
		loadouts = {},
		professions = {},
		ignored = {},
	}

	for _, section in ipairs(pipes) do
		local at = section:find("=", 1, true)
		local name = at and section:sub(1, at - 1) or section
		local field = at and section:sub(at + 1) or ""
		if name == "bags" or name == "bank" then
			local items
			items, message = parseItemList(field)
			if items == nil then
				return nil, message
			end
			build[name] = items
		elseif name == "sets" then
			for _, entry in ipairs(parseNamed(field)) do
				local gear
				gear, message = parseGearList(entry.payload)
				if gear == nil then
					return nil, message
				end
				build.sets[#build.sets + 1] = { name = entry.name, gear = gear }
			end
		elseif name == "loadouts" then
			for _, entry in ipairs(parseNamed(field)) do
				local ranks
				ranks, message = parseTrees(entry.payload)
				if ranks == nil then
					return nil, message
				end
				build.loadouts[#build.loadouts + 1] = { name = entry.name, treeRanks = ranks }
			end
		elseif name == "professions" then
			for _, slug in ipairs(split(field, ",")) do
				if slug ~= "" then
					build.professions[#build.professions + 1] = slug
				end
			end
		elseif name ~= "" then
			-- A site a version ahead of the addon is a thing that will happen;
			-- refusing its whole string would make the addon useless that day.
			build.ignored[#build.ignored + 1] = name
		end
	end

	return build
end

-- Shared with the FSB1 half of this module, added in its own task.
Codec._split = split
Codec._isDigits = isDigits
Codec._fromBase36 = fromBase36
Codec._toBase36 = toBase36
Codec._SLOT_SET = SLOT_SET

ns.Codec = Codec
return Codec
