-- addon/ForeverSixty/Rotation.lua
-- The per-level rotation lines (design section 2 item 3, section 6 Wave B):
-- data/pipeline/addonrotation.py resolves data/curated/apl/<spec>.json
-- against spellranks.json into ns.Data.rotations[spec], one band per level
-- in its own LEVEL_BANDS (10, 20, 30, 38, 40, 50, 60). This module picks
-- the right band for a character's level and slices it for novice
-- (top four lines) or advanced (every line) display. Rotation.newAbilities
-- fed the level-up rotation toast, which went with the chip (L3.1). The
-- character-panel tab's rotation region (L4.3) reads Rotation.model with
-- Rotation.plainCondition, Rotation.learnedAt and Rotation.upcoming.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Gear = ns.Gear or require("Gear")

local Rotation = {}

--- The rotation card's novice-mode line count (design section 2 item 3:
--- "the top 4 lines"; section 6's own wording for the card).
Rotation.NOVICE_LINES = 4

--- `data.rotations[spec]` for the build's own spec (Gear.specOf's rule:
--- the tree with the most points), or nil when there is no build loaded,
--- the class has no curated rotation, or Data.lua predates this field.
local function specBands(data, build, ranks)
	if build == nil or type(data.rotations) ~= "table" then
		return nil
	end
	local spec = Gear.specOf(data, build.classSlug, ranks)
	return spec and data.rotations[spec] or nil
end

--- The band whose level is the highest one <= `level`: bands' own levels
--- (data/pipeline/addonrotation.LEVEL_BANDS) are never contiguous
--- (10, 20, 30, 38, 40, 50, 60), so every level between two rungs reads
--- the earlier rung's band. `bands[1]` (the lowest rung) covers a
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

--- The rotation card's model. `visible` is false only when the spec has
--- no rotation at all (an unwritten APL, or Data.lua predates this
--- field); a band that resolved but is genuinely empty (mage-frost below
--- its first learned rank, say) still reports visible with zero lines,
--- which the card draws as nothing rather than hiding the whole card.
function Rotation.model(data, build, ranks, level, advanced)
	local bands = specBands(data, build, ranks)
	local band = bandFor(bands, level)
	if band == nil then
		return { visible = false, lines = {} }
	end
	local lines = band.lines
	local shown = lines
	if not advanced and #lines > Rotation.NOVICE_LINES then
		shown = {}
		for index = 1, Rotation.NOVICE_LINES do
			shown[index] = lines[index]
		end
	end
	return {
		visible = true,
		level = band.level,
		lines = shown,
		hasMore = not advanced and #lines > #shown,
	}
end

--- The lines newly in `level`'s band that were not in the band the
--- character read from at `previousLevel` (by spellId): "Level 20: Sunder
--- Armor opens your rotation now" (design section 2).
--- Compared by NAME, not spellId: a line's spellId is already the
--- highest rank learned by that band (pipeline.addonrotation's own
--- resolving), so the same ability ranking up between two bands (Heroic
--- Strike rank 2 to rank 3, say) would otherwise misread as a brand new
--- ability entering the rotation -- a rank-up is the talent and spell
--- trainer's news, not the rotation's.
--- Empty (never an error) when: the spec has no rotation at all;
--- `previousLevel` is nil (nothing to diff against -- a caller with no
--- known prior level, such as the very first event after the addon
--- loads, must not read whatever band a fabricated "level 0" would
--- resolve to as if it were real prior state); or the two levels share a
--- band (nothing crossed a rung).
function Rotation.newAbilities(data, build, ranks, previousLevel, level)
	if previousLevel == nil then
		return {}
	end
	local bands = specBands(data, build, ranks)
	local before = bandFor(bands, previousLevel)
	local after = bandFor(bands, level)
	if after == nil then
		return {}
	end
	local had = {}
	if before ~= nil then
		for _, line in ipairs(before.lines) do
			had[line.name] = true
		end
	end
	local fresh = {}
	for _, line in ipairs(after.lines) do
		if not had[line.name] then
			fresh[#fresh + 1] = line
		end
	end
	return fresh
end

--- spellId -> true for every line Rotation.newAbilities last found fresh,
--- so a rotation surface can glow the matching row (docs/
--- tenets.md's standard: "the row for the ability the player just
--- learned this level glows") with a plain table
--- lookup rather than a require of a surface that loads later in the TOC.
--- Its caller calls this explicitly, right after computing `fresh` -- a
--- visible side effect at its call site, not a hidden one inside this
--- file's own pure functions above.
Rotation.recentlyLearned = {}

function Rotation.markLearned(lines)
	local set = {}
	for _, line in ipairs(lines or {}) do
		set[line.spellId] = true
	end
	Rotation.recentlyLearned = set
	return set
end

-- The character-panel tab's rotation region (L4.3) -------------------------------

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

--- The next band's lines whose ability the character has not learned yet (fix round 1, item
--- 4: the band already holds only learned abilities, so "at level N" never fired on the shipped
--- data): the rung after the one `level` reads, its lines in order whose Rotation.learnedAt is
--- above `level` and whose ability the band read now does not already hold (by name, as
--- Rotation.newAbilities compares), at most `limit`. Each { line, at }. Empty at the top rung,
--- with no build, or with no rotation for the spec.
function Rotation.upcoming(data, build, ranks, level, limit)
	local bands = specBands(data, build, ranks)
	local band = bandFor(bands, level)
	if band == nil or type(level) ~= "number" then
		return {}
	end
	local nextBand
	for _, candidate in ipairs(bands) do
		if candidate.level > band.level and (nextBand == nil or candidate.level < nextBand.level) then
			nextBand = candidate
		end
	end
	if nextBand == nil then
		return {}
	end
	local held = {}
	for _, line in ipairs(band.lines) do
		held[line.name] = true
	end
	local found = {}
	for _, line in ipairs(nextBand.lines) do
		if limit ~= nil and #found >= limit then
			break
		end
		local at = Rotation.learnedAt(line.spellId, data)
		if not held[line.name] and at ~= nil and at > level then
			held[line.name] = true
			found[#found + 1] = { line = line, at = at }
		end
	end
	return found
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
