-- addon/ForeverSixty/Rotation.lua
-- The per-level rotation lines (design section 2 item 3, section 6 Wave B):
-- data/pipeline/addonrotation.py resolves data/curated/apl/<spec>.json
-- against spellranks.json into ns.Data.rotations[spec], one band per level
-- in its own LEVEL_BANDS (10, 20, 30, 38, 40, 50, 60). This module picks
-- the right band for a character's level and slices it for novice
-- (top four lines) or advanced (every line) display; Toast.lua uses
-- Rotation.newAbilities for the level-up rotation toast.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
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
--- character read from at `previousLevel` (by spellId) -- design section
--- 2's rotation toast: "Level 20: Sunder Armor opens your rotation now."
--- Compared by NAME, not spellId: a line's spellId is already the
--- highest rank learned by that band (pipeline.addonrotation's own
--- resolving), so the same ability ranking up between two bands (Heroic
--- Strike rank 2 to rank 3, say) would otherwise misread as a brand new
--- ability entering the rotation -- that rank-up is what the existing
--- talent-point toast already announces, and is not this toast's job to
--- repeat.
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

ns.Rotation = Rotation
return Rotation
