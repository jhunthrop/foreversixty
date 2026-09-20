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
--- Bounded at MAX_CODE_LENGTH iterations: a string cannot contain more parts
--- than it has characters, so this can never fire for input a caller has
--- already checked against MAX_CODE_LENGTH (decodeFS1 and decodeFSB1 both
--- do). It exists because split is a public `_`-prefixed helper that other
--- callers may consume directly with unchecked, player-supplied text; if the
--- cap is ever hit, returning what was parsed so far is fine since the
--- grammar checks downstream will refuse the result.
local function split(text, separator)
	local parts, position = {}, 1
	for _ = 1, Codec.MAX_CODE_LENGTH do
		local start, stop = text:find(separator, position, true)
		if start == nil then
			parts[#parts + 1] = text:sub(position)
			return parts
		end
		parts[#parts + 1] = text:sub(position, start - 1)
		position = stop + 1
	end
	parts[#parts + 1] = text:sub(position)
	return parts
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

--- Every refusal below is built through this instead of a bare
--- string.format. WoW's chat frame renders `|c`, `|r` and `|H…|h` as colour
--- and hyperlink markup, and a refusal echoes fragments straight out of the
--- pasted code -- the one genuinely untrusted input in this addon -- so a
--- crafted code could otherwise make the reader's own chat frame render
--- fake markup when the refusal is printed (Options.run) or repasted
--- elsewhere. Doubling `|` is the client's own escape for a literal pipe.
---
--- This is the single choke point: it escapes each argument *before* it is
--- interpolated, never the finished message afterward, so nothing
--- downstream can double the pipes and turn `||` back into a literal `|`
--- for the player. A trusted argument (a locale constant, our own data
--- build) never contains `|`, so escaping it too is a harmless no-op --
--- there is no need to sort untrusted fragments from trusted ones here.
-- Returns the formatted message alone (not `nil, message`): every call site
-- already writes `return nil, refuse(...)`, and a function that itself
-- returns two values there would hand the `return` three, shifting `message`
-- into a value nobody reads.
local function refuse(template, ...)
	local escaped = {}
	for index = 1, select("#", ...) do
		escaped[index] = (tostring(select(index, ...)):gsub("|", "||"))
	end
	return string.format(template, table.unpack(escaped))
end

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
		return nil, refuse(message, entry)
	end
	for _, part in ipairs(parts) do
		if not isDigits(part) then
			return nil, refuse(message, entry)
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
			return nil, refuse(L.codecGearEntry, entry)
		end
		local slot, value = pieces[1], pieces[2]
		if not SLOT_SET[slot] then
			return nil, refuse(L.codecSlot, slot)
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
		return nil, refuse(L.codecWrongPrefix, named, Codec.FS1_PREFIX)
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

-- ------------------------------------------------------------------ FSB1 --
--
-- FSB1:<data-build>:<class-slug>:<order>:<gear>
--
--   <order>  one triple per point spent -- tab, tier, column -- each a single
--            base-36 character, all 1-based, concatenated with no separator.
--            A 1-based tab means a leading "0" is a malformed code rather
--            than a silently wrong tree.
--   <gear>   entries joined by "," ; each entry is
--            <slot>=<item_id>[:<stat>=<value>[;<stat>=<value>]…]
--            The stat names are the engine's own vocabulary (parity contract
--            10.8), the same names Data.lua's weights table is keyed by, so a
--            planned item and a weight meet without a translation table.

local function encodeStats(stats)
	local names = {}
	for name in pairs(stats or {}) do
		names[#names + 1] = name
	end
	-- By name, not by some vocabulary's display order: "sort by name" is a
	-- rule both this file and [web]'s TypeScript decoder can implement
	-- identically with no table shared between them, and decode never reads
	-- the order back out -- only encode needs one build to make one string.
	table.sort(names)
	local parts = {}
	for index, name in ipairs(names) do
		parts[index] = string.format("%s=%d", name, math.floor(stats[name] + 0.5))
	end
	return table.concat(parts, ";")
end

function Codec.encodeFSB1(build)
	local order = {}
	for index, point in ipairs(build.order or {}) do
		order[index] = toBase36(point.tab) .. toBase36(point.tier) .. toBase36(point.column)
	end
	local gear = {}
	for index, entry in ipairs(build.gear or {}) do
		local stats = encodeStats(entry.stats)
		gear[index] = entry.slot .. "=" .. tostring(entry.itemId)
		if stats ~= "" then
			gear[index] = gear[index] .. ":" .. stats
		end
	end
	return table.concat({
		Codec.FSB1_PREFIX,
		build.dataBuild,
		build.classSlug,
		table.concat(order),
		table.concat(gear, ","),
	}, ":")
end

local function parseOrder(field)
	local order = {}
	if field == "" then
		return order
	end
	if #field % 3 ~= 0 then
		return nil, L.codecOrderLength
	end
	for position = 1, #field, 3 do
		local triple = field:sub(position, position + 2)
		local tab = fromBase36(triple:sub(1, 1))
		local tier = fromBase36(triple:sub(2, 2))
		local column = fromBase36(triple:sub(3, 3))
		-- 1-based everywhere: a zero here is a code written against a
		-- different convention, and decoding it would put points on the
		-- wrong tree without a word.
		if tab == nil or tier == nil or column == nil or tab < 1 or tier < 1 or column < 1 then
			return nil, refuse(L.codecOrderCell, triple)
		end
		order[#order + 1] = { tab = tab, tier = tier, column = column }
	end
	return order
end

local function parseStats(field)
	local stats = {}
	if field == "" then
		return stats
	end
	for _, pair in ipairs(split(field, ";")) do
		local at = pair:find("=", 1, true)
		if at == nil then
			return nil, refuse(L.codecStatPair, pair)
		end
		local name, value = pair:sub(1, at - 1), pair:sub(at + 1)
		if name == "" or not isDigits(value) then
			return nil, refuse(L.codecStatPair, pair)
		end
		stats[name] = tonumber(value)
	end
	return stats
end

local function parseFSB1Gear(field)
	local gear = {}
	if field == "" then
		return gear
	end
	for _, entry in ipairs(split(field, ",")) do
		local at = entry:find("=", 1, true)
		if at == nil then
			return nil, refuse(L.codecGearEntry, entry)
		end
		local slot = entry:sub(1, at - 1)
		if not SLOT_SET[slot] then
			return nil, refuse(L.codecSlot, slot)
		end
		local rest = split(entry:sub(at + 1), ":")
		if not isDigits(rest[1]) then
			return nil, refuse(L.codecGearEntry, entry)
		end
		local stats, message = parseStats(table.concat(rest, ":", 2))
		if stats == nil then
			return nil, message
		end
		gear[#gear + 1] = { slot = slot, itemId = tonumber(rest[1]), stats = stats }
	end
	return gear
end

function Codec.decodeFSB1(code)
	if #code > Codec.MAX_CODE_LENGTH then
		return nil, L.codecTooLong
	end
	local parts = split(code:match("^%s*(.-)%s*$"), ":")
	if parts[1] ~= Codec.FSB1_PREFIX then
		local named = (parts[1] == nil or parts[1] == "") and L.codecUnlabelled or parts[1]
		return nil, refuse(L.codecWrongPrefix, named, Codec.FSB1_PREFIX)
	end
	if #parts < 5 then
		return nil, L.codecShort
	end
	local order, message = parseOrder(parts[4])
	if order == nil then
		return nil, message
	end
	local gear
	gear, message = parseFSB1Gear(table.concat(parts, ":", 5))
	if gear == nil then
		return nil, message
	end
	return {
		dataBuild = parts[2],
		classSlug = parts[3],
		order = order,
		gear = gear,
	}
end

--- Compare two dotted build strings numerically, field by field.
--- Returns -1, 0 or 1. A non-numeric field compares as 0, so a build string
--- from a client that changes its shape never makes the addon refuse a code
--- it could read.
local function compareBuilds(left, right)
	local a, b = split(left, "."), split(right, ".")
	for index = 1, math.max(#a, #b) do
		local x = tonumber(a[index]) or 0
		local y = tonumber(b[index]) or 0
		if x ~= y then
			return x < y and -1 or 1
		end
	end
	return 0
end

--- An order reconstructed from a final tree, lowest tier first, left to
--- right. The same rule the site's `orderFromRanks` uses, and the reason the
--- site's import box says the order is approximated: nothing in the game
--- records the order a build was actually spent in.
local function approximateOrder(tabs, treeRanks)
	local order = {}
	for tabIndex, tab in ipairs(tabs) do
		local wanted = {}
		local ranks = treeRanks[tabIndex] or {}
		for position, talent in ipairs(tab.talents) do
			local rank = math.min(ranks[position] or 0, talent.maxRank)
			if rank > 0 then
				wanted[#wanted + 1] = { talent = talent, rank = rank }
			end
		end
		table.sort(wanted, function(left, right)
			if left.talent.tier ~= right.talent.tier then
				return left.talent.tier < right.talent.tier
			end
			return left.talent.column < right.talent.column
		end)
		for _, entry in ipairs(wanted) do
			for _ = 1, entry.rank do
				order[#order + 1] = {
					tab = tabIndex,
					tier = entry.talent.tier,
					column = entry.talent.column,
				}
			end
		end
	end
	return order
end

--- Load a build from either format.
---
--- FSB1 is the site's own addon code and is taken as it stands. FS1 is the
--- character export, which the site's Top Gear also hands out as "Copy to
--- addon": it carries a final tree and no stats, so the order is
--- approximated and `statsUnknown` is set, which is what stops Gear from
--- scoring a planned item at zero and calling that a downgrade.
function Codec.loadBuild(code, data)
	-- Length-checked first, before any trimming or splitting: `code` is
	-- the one genuinely untrusted input in the system (Follow.load is
	-- called from Options.readInbox on a saved-variables file the addon
	-- does not own), so a hostile string must be refused before it is
	-- touched at all, not merely before decodeFS1/decodeFSB1 parse it.
	if #code > Codec.MAX_CODE_LENGTH then
		return nil, L.codecTooLong
	end
	-- Trimmed once, here, before the prefix is sniffed: a pasted code
	-- commonly carries leading or trailing whitespace from the chat edit
	-- box it was copied out of, and reading the prefix off the untrimmed
	-- string would route a perfectly valid FSB1 code into the FS1 branch.
	-- decodeFSB1 and decodeFS1 both trim again internally, but trimming an
	-- already-trimmed string is a no-op, not a second trim.
	code = code:match("^%s*(.-)%s*$")
	-- A plain pattern match reads element one without allocating the
	-- table `split` would build for it.
	local prefix = code:match("^([^:]*)")
	local build, message
	if prefix == Codec.FSB1_PREFIX then
		build, message = Codec.decodeFSB1(code)
		if build == nil then
			return nil, message
		end
		build.format = Codec.FSB1_PREFIX
		build.statsUnknown = false
	else
		local exported
		exported, message = Codec.decodeFS1(code)
		if exported == nil then
			return nil, message
		end
		local class = data.classes[exported.classSlug]
		if class == nil then
			return nil, refuse(L.codecUnknownClass, exported.classSlug)
		end
		local gear = {}
		for index, entry in ipairs(exported.gearSlots) do
			gear[index] = { slot = entry.slot, itemId = entry.itemId, stats = {} }
		end
		build = {
			dataBuild = exported.dataBuild,
			classSlug = exported.classSlug,
			order = approximateOrder(class.tabs, exported.treeRanks),
			gear = gear,
			format = Codec.FS1_PREFIX,
			statsUnknown = true,
		}
	end

	if data.classes[build.classSlug] == nil then
		return nil, refuse(L.codecUnknownClass, build.classSlug)
	end
	if compareBuilds(build.dataBuild, data.build) > 0 then
		return nil, refuse(L.codecNewerBuild, build.dataBuild, data.build)
	end
	return build
end

ns.Codec = Codec
return Codec
