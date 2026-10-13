-- addon/ForeverSixty/Rotation.lua
-- The per-level rotation lines (design section 2 item 3, section 6 Wave B):
-- data/pipeline/addonrotation.py resolves data/curated/apl/<spec>.json
-- against spellranks.json into ns.Data.rotations[spec], one band per level
-- in its own LEVEL_BANDS (10, 20, 30, 38, 40, 50, 60). This module picks
-- the right band for a character's level and slices it for novice
-- (top four lines) or advanced (every line) display. Rotation.newAbilities
-- fed the level-up rotation toast, which went with the chip (L3.1). The
-- character-panel tab's rotation region (L4.3) reads Rotation.model with
-- Rotation.plainCondition and Rotation.learnedAt.
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

--- A rotation line's condition in plain words (L4.3 ruling 3): parentheticals stripped (a
--- nested one whole), and an aside set off by " -- " with them (the data's editorial asides:
--- "-- the cited guide's own rotation step, not skipped"); cut at the first sentence end; ended
--- with a full stop. A colon is kept: what follows one is often the rule itself ("Flash Heal
--- ...: it is only for a tank that is genuinely low"). With `measure` and `width` (the
--- panel's), never more than two lines: the second is cut on a word with an ellipsis.
--- Answers the text and the line count it wraps to (nil with nothing to measure with); nil
--- for no condition.
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
	-- The second line, a word at a time, until it and the ellipsis fit.
	local words = {}
	for word in lines[2]:gmatch("%S+") do
		words[#words + 1] = word
	end
	local second = ""
	for count = #words, 1, -1 do
		second = table.concat(words, " ", 1, count) .. ELLIPSIS
		local size = measure(second)
		if type(size) ~= "number" or size <= width then
			break
		end
	end
	return lines[1] .. " " .. second, 2
end

ns.Rotation = Rotation
return Rotation
