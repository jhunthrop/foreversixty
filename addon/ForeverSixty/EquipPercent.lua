-- addon/ForeverSixty/EquipPercent.lua
-- The rating-family stats an item's on-equip spell states as a literal percent (Trueaim Gauntlets'
-- "Equip: Improves your chance to hit by 3%."), which GetItemStats never reports: the client reads
-- those from ItemSparse's stat columns, and a percent item's hit or crit is a spell, not a column
-- (data/README.md, "Hit, crit, dodge, parry and block as percentages"; items/<class>.json
-- percent_stats). Read from the item's own tooltip and handed to Gear.statsOf as the rating points
-- that buy the same percent, so Gear.score's one rating-to-percent step prices it exactly.
--
-- The 1.60 client also prints a rating stat as an equip line in the same words from its own percent
-- ("Slashclaw Bracers' 10 hit reads "Improves your chance to hit by 1.0%."), so the literal part is
-- what the tooltip states beyond what the item's rating gives: the contract's percent_stats rule,
-- `total - rating / factor`, rounded to the whole percent an on-equip spell always states.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Compat = ns.Compat or require("Compat")
local Rules = ns.Rules or require("Rules")

local EquipPercent = {}

--- A Lua pattern for one of Locale's equipPercentLines phrases, anchored, its %s captured as the
--- number (an optional percent sign after it).
local function patternOf(phrase)
	local escaped = phrase:gsub("[%^%$%(%)%.%[%]%*%+%-%?]", "%%%0")
	return "^" .. escaped:gsub("%%s", "([%%d%%.,]+)%%%%?") .. "$"
end

--- stat -> { pattern, ... } from Locale's phrases. Built once, on first need.
local patterns

local function patternsByStat()
	if patterns == nil then
		patterns = {}
		for stat, phrases in pairs(L.equipPercentLines) do
			patterns[stat] = {}
			for index, phrase in ipairs(phrases) do
				patterns[stat][index] = patternOf(phrase)
			end
		end
	end
	return patterns
end

local function plainText(text)
	return (text:gsub("|c%x%x%x%x%x%x%x%x", ""):gsub("|r", ""):gsub("^%s+", ""):gsub("%s+$", ""))
end

--- The percent each rating-family stat's equip lines state, summed: { hit = 3 } for Trueaim
--- Gauntlets' rows ({ left, right } each, the client's text). The client's "Equip:" word
--- (ITEM_SPELL_TRIGGER_ONEQUIP) is taken off first. Pure over the rows.
function EquipPercent.stated(rows)
	local trigger = Compat.globalString("ITEM_SPELL_TRIGGER_ONEQUIP") or L.equipTrigger
	local found = {}
	for _, row in ipairs(rows or {}) do
		if type(row.left) == "string" then
			local text = plainText(row.left)
			if text:sub(1, #trigger) == trigger then
				text = plainText(text:sub(#trigger + 1))
			end
			for stat, list in pairs(patternsByStat()) do
				for _, pattern in ipairs(list) do
					local amount = tonumber(((text:match(pattern) or ""):gsub(",", ".")))
					if amount ~= nil then
						found[stat] = (found[stat] or 0) + amount
					end
				end
			end
		end
	end
	return found
end

--- The literal percent each stat's on-equip spell adds, past what the item's rating points
--- (`stats`, Gear.statsOf's client read) already give: whole percents above zero only. Pure.
function EquipPercent.literal(rows, stats)
	local literal = {}
	for stat, total in pairs(EquipPercent.stated(rows)) do
		local factor = Rules.RATING_PER_PERCENT[stat] or 1
		local rest = math.floor(total - ((stats or {})[stat] or 0) / factor + 0.5)
		if rest > 0 then
			literal[stat] = rest
		end
	end
	return literal
end

--- Literal percents read so far, by item link, so an item's tooltip is read once; cleared when it
--- grows past a few bags' worth of links.
EquipPercent.cache = {}
local CACHE_LIMIT = 200
local cacheSize = 0

--- The item's on-equip literal percents (EquipPercent.literal of its own tooltip), as the rating
--- points that buy them: { hit = 30 } for Trueaim Gauntlets. Empty, and asked again next time,
--- while the client cannot give the tooltip yet.
function EquipPercent.ratingOf(link, stats)
	if type(link) ~= "string" then
		return {}
	end
	local cached = EquipPercent.cache[link]
	if cached ~= nil then
		return cached
	end
	local rows = Compat.hyperlinkRows(link)
	if rows == nil or #rows == 0 then
		return {}
	end
	local points = {}
	for stat, percent in pairs(EquipPercent.literal(rows, stats)) do
		points[stat] = percent * (Rules.RATING_PER_PERCENT[stat] or 1)
	end
	if cacheSize >= CACHE_LIMIT then
		EquipPercent.cache, cacheSize = {}, 0
	end
	EquipPercent.cache[link] = points
	cacheSize = cacheSize + 1
	return points
end

ns.EquipPercent = EquipPercent
return EquipPercent
