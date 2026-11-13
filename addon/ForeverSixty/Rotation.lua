-- addon/ForeverSixty/Rotation.lua
-- The per-level rotation lines (design section 2 item 3, section 6 Wave B):
-- data/pipeline/addonrotation.py resolves data/curated/apl/<spec>.json
-- against spellranks.json into ns.Data.rotations[spec]: a band at each of its
-- LEVEL_BANDS rungs (10, 20, 30, 38, 40, 50, 60) and at every level between
-- where a line changes (a rank row or a rank learned mid-band). This module
-- picks the spec and the band for a level, the levels the priority changes at,
-- the level an ability is learned at, and a condition in plain words: the
-- window's Rotation tab (surfaces/RotationTab.lua, hub spec 4.4) reads them.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Gear = ns.Gear or require("Gear")
local Follow = ns.Follow or require("Follow")
local Compat = ns.Compat or require("Compat")
local Export = ns.Export or require("Export")

local Rotation = {}

--- The band whose level is the highest one <= `level`: a band starts at each
--- rung and wherever a line changes between rungs (a rank row from 34), so
--- every level reads the band that holds its own ranks. `bands[1]` (the lowest rung) covers a
--- character below it too, so a level-1 alt has something to show rather
--- than nothing before its first rung.
local function bandFor(bands, level)
	if type(bands) ~= "table" or #bands == 0 then
		return nil
	end
	local chosen = bands[1]
	for _, band in ipairs(bands) do
		if band.level <= (level or 0) then
			chosen = band
		end
	end
	return chosen
end

--- `band` as the player's race casts it: a line whose rank a race's rank-choice row
--- changes (data/curated/rank-choices.json, round 13) carries each keyed race's id in
--- `raceSpellId` by the site's race slug, and the copy shows the player's own. The
--- shipped band is never written; a race no row names, or a client that names no
--- race, keeps each line's own `spellId`.
local function ofRace(band)
	if band == nil then
		return nil
	end
	local race = Compat.playerRaceSlug()
	local lines = {}
	for index, line in ipairs(band.lines) do
		local id = race and type(line.raceSpellId) == "table" and line.raceSpellId[race] or nil
		if id ~= nil and id ~= line.spellId then
			local copy = {}
			for key, value in pairs(line) do
				copy[key] = value
			end
			copy.spellId = id
			line = copy
		end
		lines[index] = line
	end
	return { level = band.level, lines = lines }
end

--- The spec the Rotation tab shows (hub spec 4.4 and 8) and whether a followed build names it: the
--- build's own (specBands' rule), else the tree the points name (Gear.specOf), else the class's first
--- tree that has a rotation, so a character with no build still reads its class's priority. nil for a
--- class with no rotation at all.
function Rotation.specFor(data, build, ranks, classSlug)
	local rotations = type(data) == "table" and data.rotations or nil
	if type(rotations) ~= "table" then
		return nil, false
	end
	if build ~= nil then
		local spec = Follow.buildSpec(data, build) or Gear.specOf(data, build.classSlug, ranks)
		if spec ~= nil and rotations[spec] ~= nil then
			return spec, true
		end
	end
	local own = Gear.specOf(data, classSlug, ranks)
	if own ~= nil and rotations[own] ~= nil then
		return own, false
	end
	local class = type(data.classes) == "table" and data.classes[classSlug] or nil
	for _, tab in ipairs(type(class) == "table" and class.tabs or {}) do
		local spec = type(tab.name) == "string" and classSlug .. "-" .. Export.slugify(tab.name) or nil
		if spec ~= nil and rotations[spec] ~= nil then
			return spec, false
		end
	end
	return nil, false
end

--- What a band shows the player: each line's name and condition, in order (a rank-up changes neither).
local function shownKey(band)
	local parts = {}
	for _, line in ipairs(band.lines) do
		parts[#parts + 1] = tostring(line.name) .. "\n" .. tostring(line.condition)
	end
	return table.concat(parts, "\n\n")
end

--- The levels `spec`'s priority changes at (the stepper's stops, hub spec 4.4): each band's level
--- whose lines read differently from the band before it, ascending. A band that only ranks a line up
--- is no stop. Empty for a spec with no rotation.
function Rotation.changeLevels(data, spec)
	local bands = type(data) == "table" and type(data.rotations) == "table" and data.rotations[spec] or nil
	local levels, previous = {}, nil
	for _, band in ipairs(type(bands) == "table" and bands or {}) do
		local key = shownKey(band)
		if key ~= previous then
			levels[#levels + 1] = band.level
		end
		previous = key
	end
	return levels
end

--- `spec`'s band for `level` as the player's race casts it (bandFor's rule: the first band covers
--- a level below it), or nil.
function Rotation.bandAt(data, spec, level)
	local bands = type(data) == "table" and type(data.rotations) == "table" and data.rotations[spec] or nil
	return ofRace(bandFor(bands, level))
end

-- Learned levels and plain conditions -------------------------------------------

--- The level an ability is first learned at (Data.abilityLevels, contract 1: every rank of
--- an ability maps to its first rank's level, keyed by the very id a band's line carries),
--- or nil for a spell the data has no level for. `data` is Data.lua's table (ns.Data when nil).
function Rotation.learnedAt(spellId, data)
	data = data or ns.Data
	local levels = type(data) == "table" and data.abilityLevels or nil
	if spellId == nil or type(levels) ~= "table" then
		return nil
	end
	return tonumber(levels[spellId])
end

--- The ellipsis a cut condition ends with (the tooltip's own).
local ELLIPSIS = L.tooltipNamePending

--- `text` wrapped on spaces at `width` as `measure` (text -> width) reads it, greedily, the
--- way a font string of that width wraps: the lines, each a string.
local function wrap(text, measure, width)
	local lines, current = {}, nil
	for word in text:gmatch("%S+") do
		local candidate = current == nil and word or (current .. " " .. word)
		local size = measure(candidate)
		if current ~= nil and type(size) == "number" and size > width then
			lines[#lines + 1] = current
			current = word
		else
			current = candidate
		end
	end
	if current ~= nil then
		lines[#lines + 1] = current
	end
	return lines
end

--- Words that lead into an object ("for 10", "to the", "of a"): a cut never ends on one.
Rotation.LEADS = { ["for"] = true, ["to"] = true, ["of"] = true, ["at"] = true, ["by"] = true, ["with"] = true,
	["than"] = true, ["over"] = true, ["under"] = true, ["above"] = true, ["below"] = true, ["every"] = true,
	["per"] = true, ["a"] = true, ["an"] = true, ["the"] = true }

--- The words a clause boundary falls before (L4.3 fix round 3, item 2), besides a word ending
--- in ",", ";" or ":".
Rotation.CONNECTORS = { ["and"] = true, ["or"] = true, ["when"] = true, ["unless"] = true }
--- Where in the second line a clause boundary is preferred to a word cut: from this share of
--- its measured width on (the last 40%).
Rotation.CLAUSE_FROM = 0.6

local function isConnector(word)
	return word ~= nil and Rotation.CONNECTORS[word:lower()] == true
end

--- The first `count` of `words`, joined, a trailing ",", ";" or ":" dropped before the ellipsis.
local function kept(words, count)
	return (table.concat(words, " ", 1, count):gsub("[,;:]$", ""))
end

--- How many words of the second line (`words`; `following` the first word after it) a cut
--- condition keeps before its ellipsis, within `width` as `measure` reads it (fix round 3,
--- item 2): the last clause boundary (after a word ending in ",", ";" or ":", or before
--- "and", "or", "when", "unless") lying in the last 40% of the line; else a cut on a word,
--- never left on a connector, and never just after a figure ("a 25%" cut from "reserve"):
--- there the last clause boundary anywhere in the line, when it has one, else the words
--- before the figure and what leads into it.
function Rotation.cutCount(words, following, measure, width)
	local function fits(count)
		local size = measure(kept(words, count) .. ELLIPSIS)
		return type(size) ~= "number" or size <= width
	end
	local line = measure(table.concat(words, " "))
	local latest
	for count = #words, 1, -1 do
		local word = words[count]
		local after = words[count + 1] or following
		-- A boundary right after a figure or a leading word ("for 10 and ...") is no place to cut.
		local bare = not word:find("%d") and Rotation.LEADS[word:lower()] ~= true
		if not isConnector(word) and bare and (word:find("[,;:]$") ~= nil or isConnector(after)) and fits(count) then
			local size = measure(kept(words, count))
			if type(line) == "number" and type(size) == "number" and size >= Rotation.CLAUSE_FROM * line then
				return count
			end
			latest = latest or count
		end
	end
	local count = 1
	for candidate = #words, 1, -1 do
		if fits(candidate) then
			count = candidate
			break
		end
	end
	if count > 1 and isConnector(words[count]) then
		count = count - 1
	end
	if words[count]:find("%d") ~= nil then
		if latest ~= nil then
			return latest
		end
		if count == 1 then
			return 1
		end
		-- No clause boundary to go back to (Tranquility, "... every 2 seconds for 10..."): back up
		-- past the figure and the words that lead into it (L3.7 item 27).
		count = count - 1
		while count > 1 and (isConnector(words[count]) or Rotation.LEADS[words[count]:lower()] == true) do
			count = count - 1
		end
	end
	return count
end

--- A rotation line's condition in plain words (L4.3 ruling 3): parentheticals stripped (a
--- nested one whole), and an aside set off by " -- " with them (the data's editorial asides:
--- "-- the cited guide's own rotation step, not skipped"); cut at the first sentence end; ended
--- with a full stop. A colon is kept: what follows one is often the rule itself ("Flash Heal
--- ...: it is only for a tank that is genuinely low"). With `measure` and `width` (the
--- panel's), never more than two lines: the second is cut with an ellipsis, at a clause where
--- one lies near its end (Rotation.cutCount). Answers the text and the line count it wraps to
--- (nil with nothing to measure with); nil for no condition.
function Rotation.plainCondition(text, measure, width)
	if type(text) ~= "string" then
		return nil
	end
	local plain = text:gsub("%s*%b()", "")
	-- " -- an aside -- " in the middle, then a lone " -- " running to the end.
	plain = plain:gsub("%s+%-%-%s.-%s%-%-%s+", " ")
	plain = plain:gsub("%s+%-%-%s.*$", "")
	-- The first sentence end: a stop, question or exclamation mark before a space (a decimal
	-- point has a digit after it).
	local cut = plain:find("[%.!?]%s")
	if cut ~= nil then
		plain = plain:sub(1, cut - 1)
	end
	plain = plain:gsub("%s+", " "):gsub("^ ", ""):gsub(" $", "")
	plain = plain:gsub("[%.!?:;,]$", "")
	if plain == "" then
		return nil
	end
	plain = plain .. "."
	if type(measure) ~= "function" or type(width) ~= "number" or type(measure(plain)) ~= "number" then
		return plain, nil
	end
	local lines = wrap(plain, measure, width)
	if #lines <= 2 then
		return plain, #lines
	end
	local words = {}
	for word in lines[2]:gmatch("%S+") do
		words[#words + 1] = word
	end
	local count = Rotation.cutCount(words, lines[3]:match("^%S+"), measure, width)
	return lines[1] .. " " .. kept(words, count) .. ELLIPSIS, 2
end

ns.Rotation = Rotation
return Rotation
