-- addon/ForeverSixty/surfaces/ItemTooltip.lua
-- The item tooltip (spec section 4.3): after the client's own lines, a
-- "Forever Sixty" section. A header naming the spec and where the weights
-- came from ("Forever Sixty · Arms · Your sim", "Equipped" right of it on
-- the item worn), then one primary line for every piece of gear: the
-- verdict against what is worn ("Upgrade +4.2 DPS over Chausses of
-- Westfall", "Sidegrade (+0.3 DPS)"), or why not ("Not for you until 40:
-- plate armor", "Not with your two-hander", "Not for your plan"), or, on
-- the item worn, the slot's next upgrade ("Next upgrade: Mutant Scale
-- Breastplate +6.1 DPS") or "Best until level 30". Then any capped stat, the plan
-- ("Planned for Arms at 30"), the source on its own line with its drop
-- chance, and two alternatives in their quality colour. The compare
-- tooltip beside it shows the next upgrade, once the client holds it.
--
-- Pure model, thin hook: ItemTooltip.section builds the lines from plain
-- tables (every number from Progress.itemVerdict and Progress.roundedGain,
-- the weights from Gear.currentWeights, the wear rules from Rules.canEquip);
-- the guarded hooks after it read the client and draw. A section with
-- nothing to say is not drawn at all (spec section 2.2).
--
-- An EMPTY character slot (L4.1) shows the same section with the band's pick
-- for the slot (ItemTooltip.emptySlotSection). A player's tooltip is
-- surfaces/PlayerTooltip.lua's (L4.4): this file draws on items only.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Talents = ns.Talents or require("Talents")
local Gear = ns.Gear or require("Gear")
local Export = ns.Export or require("Export")
local Compat = ns.Compat or require("Compat")
local Theme = ns.Theme or require("Theme")
local Prefs = ns.Prefs or require("Prefs")
local Rules = ns.Rules or require("Rules")
local Follow = ns.Follow or require("Follow")
local Progress = ns.Progress or require("Progress")

local ItemTooltip = {}

--- The site slot name -> the client's equipped-item slot id, built once
--- from Export's own table.
local SLOT_IDS = {}
for _, entry in ipairs(Export.INVENTORY_SLOTS) do
	SLOT_IDS[entry.slot] = entry.id
end

--- An item link's id.
local function itemIdOf(link)
	if link == nil then
		return nil
	end
	local id = Compat.itemInfoInstant(link)
	if id ~= nil then
		return tonumber(id)
	end
	return tonumber(link:match("item:(%d+)"))
end

local function equippedLinkFor(slot)
	local slotId = SLOT_IDS[slot]
	if slotId == nil or type(GetInventoryItemLink) ~= "function" then
		return nil
	end
	return GetInventoryItemLink("player", slotId)
end

-- ---------------------------------------------------------------------
-- The section's lines, pure.
-- ---------------------------------------------------------------------

local function unitWord(unit)
	return unit == "hps" and L.tooltipUnitHps or L.tooltipUnitDps
end

local function muted(text)
	return Theme.inlineColor("tooltipMuted", text)
end

--- The verdict line for one Progress.itemVerdict answer, as { text, color },
--- or nil with nothing to say. The verdict was asked with the HOVERED item
--- as `plan` and the worn one as `worn`, so its delta is worn minus hovered
--- and the hovered item's gain is minus the delta. `unit` is "dps", "hps"
--- or "score"; `wornName` the worn item's name (nil while the client has not
--- sent it).
---
--- One floor for the chip and the tooltip (Progress.MIN_GAIN, 1.0 in the
--- unit), read off the gain as shown (Progress.roundedGain, the chip's own):
--- a gain from it reads "Upgrade", a loss from it "Downgrade ... over
--- <worn>", anything between "Sidegrade (+0.3 DPS)" (the size dropped when it
--- shows as nothing, and for a score). An empty slot is filled, green at any
--- gain that shows (an upgrade whatever the floor, L4.1 fix round 3); `match` is
--- "Equipped", `blocked` (a hand a worn two-hander holds) "Not with your two-hander".
function ItemTooltip.verdictLine(verdict, unit, wornName)
	if verdict == nil then
		return nil
	end
	if verdict.state == "match" then
		return { text = L.tooltipEquipped, color = "muted" }
	end
	if verdict.state == "blocked" then
		return { text = L.tooltipTwoHander, color = "muted" }
	end
	if verdict.delta == nil or verdict.state == "noplan" then
		return nil
	end
	unit = unit or "score"
	local gain = Progress.roundedGain(-verdict.delta, unit == "score" and "score" or "dps")
	local size = math.abs(gain)
	if verdict.state == "empty" then
		if gain <= 0 then
			return nil
		end
		-- Any fill that shows a gain is an upgrade (L4.1 fix round 3): the chip offers
		-- it and the slot carries an arrow, so it reads green here too.
		local color = "success"
		if unit == "score" then
			return { text = string.format(L.tooltipFillsScore, gain), color = color }
		end
		return { text = string.format(L.tooltipFills, gain, unitWord(unit)), color = color }
	end
	local name = wornName or L.tooltipNamePending
	if gain >= Progress.MIN_GAIN then
		if unit == "score" then
			return { text = string.format(L.tooltipVerdictScore, gain, name), color = "success" }
		end
		local format = unit == "hps" and L.tooltipVerdictHps or L.tooltipVerdictDps
		return { text = string.format(format, gain, name), color = "success" }
	end
	if gain <= -Progress.MIN_GAIN then
		if unit == "score" then
			return { text = string.format(L.tooltipDownScore, size, name), color = "warning" }
		end
		return { text = string.format(L.tooltipDown, size, unitWord(unit), name), color = "warning" }
	end
	if unit == "score" or size == 0 then
		return { text = L.tooltipSide, color = "muted" }
	end
	local sign = gain > 0 and "+" or L.tooltipMinus
	return { text = string.format(L.tooltipSideGain, sign, size, unitWord(unit)), color = "muted" }
end

--- The UTF-8 characters of `text`, one string each: a cut never splits a
--- multibyte character.
local function chars(text)
	local list = {}
	for char in text:gmatch("[%z\1-\127\194-\244][\128-\191]*") do
		list[#list + 1] = char
	end
	return list
end

--- `name` cut to `length` characters (UTF-8 characters, not bytes) with the
--- ellipsis after it.
local function cut(name, length)
	local list = chars(name)
	return table.concat(list, "", 1, math.min(length, #list)) .. L.tooltipNamePending
end
ItemTooltip.cut = cut

--- A line's visible text: the client's colour, link and texture escapes
--- removed, so a width is measured on what is drawn.
function ItemTooltip.stripEscapes(text)
	local visible = text:gsub("|c%x%x%x%x%x%x%x%x", ""):gsub("|r", "")
	visible = visible:gsub("|H.-|h(.-)|h", "%1"):gsub("|T.-|t", ""):gsub("|A.-|a", "")
	return visible
end

--- A measure of visible text only, around `raw` (text -> width).
function ItemTooltip.visibleMeasure(raw)
	return function(text)
		return raw(ItemTooltip.stripEscapes(text))
	end
end

--- `fits(text)`: whether a line stays within `room` as `measure` reads it;
--- always true with nothing to measure with.
local function fitter(measure, room)
	return function(text)
		if type(measure) ~= "function" or room == nil then
			return true
		end
		local width = measure(text)
		return type(width) ~= "number" or width <= room
	end
end

--- `build(name)` with `name` shortened, a character at a time, until it
--- fits; never below `floor` characters (the shortest it then gives).
local function shorten(build, name, fits, floor)
	local text = build(name)
	if fits(text) then
		return text
	end
	local list = chars(name)
	for length = #list - 1, math.max(1, floor or 1), -1 do
		text = build(cut(name, length))
		if fits(text) then
			return text
		end
	end
	return text
end


--- `text` cut after `count` visible characters (UTF-8 characters; colour,
--- link and texture escapes kept whole and not counted), the ellipsis after
--- the cut and a colour still open closed.
function ItemTooltip.truncateVisible(text, count)
	local out, shown, open, index = {}, 0, false, 1
	while index <= #text do
		local escape = text:match("^|c%x%x%x%x%x%x%x%x", index) or text:match("^|r", index)
			or text:match("^|H.-|h", index) or text:match("^|h", index) or text:match("^|T.-|t", index)
			or text:match("^||", index)
		if escape ~= nil then
			if escape:sub(1, 2) == "|c" then
				open = true
			elseif escape == "|r" then
				open = false
			end
			out[#out + 1] = escape
			index = index + #escape
		else
			local char = text:match("^[%z\1-\127\194-\244][\128-\191]*", index) or text:sub(index, index)
			if shown == count then
				out[#out + 1] = L.tooltipNamePending
				if open then
					out[#out + 1] = "|r"
				end
				return table.concat(out)
			end
			out[#out + 1] = char
			shown = shown + 1
			index = index + #char
		end
	end
	return table.concat(out)
end

--- Any other line within `room` (the header, a refusal, the plan, a capped
--- stat, "Best until level 30"): whole when it fits, else cut by measured width
--- with an ellipsis, its colours kept.
function ItemTooltip.fitPlain(text, measure, room)
	local fits = fitter(measure, room)
	if text == nil or fits(text) then
		return text
	end
	local cutText = text
	for count = #chars(ItemTooltip.stripEscapes(text)) - 1, 1, -1 do
		cutText = ItemTooltip.truncateVisible(text, count)
		if fits(cutText) then
			return cutText
		end
	end
	return cutText
end

local function trimText(text)
	return (text:gsub("^%s+", ""):gsub("%s+$", ""))
end

--- The verdict as printed lines within `room`: one line when it fits; else
--- the verdict and its figure ("Downgrade −46.0 DPS over") and, on a second
--- line, the compared item's name, shortened with an ellipsis only when it
--- still does not fit on its own line, and then on its second word at the soonest.
--- One line with no measure or room. nil with no verdict line.
function ItemTooltip.fitVerdict(verdict, unit, wornName, measure, room)
	local line = ItemTooltip.verdictLine(verdict, unit, wornName)
	if line == nil then
		return nil
	end
	local fits = fitter(measure, room)
	if wornName == nil or fits(line.text) then
		return { line.text }
	end
	local head = trimText(ItemTooltip.verdictLine(verdict, unit, "").text)
	if head == line.text then
		return { line.text }
	end
	-- Never below the first word and a letter of the second: a name too wide even
	-- then is cut on its second word rather than drawn wide.
	local firstWord = wornName:match("^(%S+%s+)%S") or ""
	local name = shorten(function(text) return text end, wornName, fits, #chars(firstWord) + 1)
	return { head, name }
end

--- The line Rules.canEquip's refusal reads as, in place of the verdict:
--- "Not for you until 40: plate armor", "Not for you: rogues cannot use
--- two-handed swords", "Not for you: shamans cannot wield two weapons".
--- `subclass` is canEquip's own plain-words reason. nil for "ok".
function ItemTooltip.untilLine(classToken, answer, level, subclass)
	if answer == "until" then
		return string.format(L.tooltipUntil, level, subclass)
	end
	if answer == "never" then
		local class = type(classToken) == "string" and L.tooltipClassPlural[classToken:upper()] or nil
		class = class or tostring(classToken)
		if subclass == "dual wielding" then
			return string.format(L.tooltipNeverTwoWeapons, class)
		end
		return string.format(L.tooltipNever, class, subclass)
	end
	return nil
end

--- data/pipeline/addonlua.py's SOURCE_KIND_CODES, decoded back for the
--- source line. tests/tooltip_bis_spec.lua pins this table's key set
--- against the pipeline's set of source kinds.
ItemTooltip.SOURCE_KIND_NAMES = {
	Q = "Quest",
	V = "Vendor",
	D = "Dungeon",
	C = "Crafted",
	R = "Reputation",
	P = "PvP",
	W = "World",
	A = "Raid",
	B = "World Drop",
}

--- The kinds whose source is "<instance>: <boss>".
local INSTANCE_KINDS = { D = true, A = true }

--- A pick's drop chance as a whole percent, never 0 for a real drop.
local function percentOf(chance)
	if type(chance) ~= "number" or chance <= 0 then
		return nil
	end
	return math.max(1, math.floor(chance * 100 + 0.5))
end

--- The drop chance as printed, for the chip's dungeon list (L4.2).
ItemTooltip.percentOf = percentOf

--- One line saying where a pick comes from: "Wailing Caverns · Mutanus the
--- Devourer · 33% drop" for an instance drop (the source split on its first
--- ": "), "<name> · Vendor" or "<source> · Quest" for the rest (the kind left
--- out when the source already says it, "World drop"), "Crafted · Leatherworking"
--- for a crafted pick ("Crafted" alone when the data names no profession), the drop
--- chance last when the data has one. A world drop that binds when equipped (`boe`)
--- reads "World drop · check the AH" (L4.1 fix round 5). nil with no source.
function ItemTooltip.sourceLine(pick, boe)
	local first, second, percent = ItemTooltip.sourceParts(pick, boe)
	return ItemTooltip.sourceText(first, second, percent)
end

--- A pick's source as its parts: the place (or the source), the boss (or the
--- kind's name, nil when the source already says it) and the drop chance as a
--- whole percent (nil for no drop). Nothing with no source, but a crafted pick.
function ItemTooltip.sourceParts(pick, boe)
	if type(pick) ~= "table" then
		return nil
	end
	local crafted = ItemTooltip.SOURCE_KIND_NAMES.C
	if type(pick.source) ~= "string" or pick.source == "" then
		-- A crafted pick says so even when the data names no profession (L4.1 fix round 5).
		if pick.kind == "C" then
			return crafted, nil, nil
		end
		return nil
	end
	local first, second
	local place, boss = pick.source:match("^(.-): (.+)$")
	if INSTANCE_KINDS[pick.kind] and place ~= nil then
		first, second = place, boss
	elseif pick.kind == "C" then
		-- "Crafted · Leatherworking": the kind leads, the profession follows.
		first, second = crafted, pick.source
	else
		first = pick.source
		local kindName = ItemTooltip.SOURCE_KIND_NAMES[pick.kind]
		if kindName ~= nil and kindName:lower() ~= first:lower() then
			second = kindName
		end
		if boe and pick.kind == "B" and second == nil then
			second = L.tooltipCheckAH
		end
	end
	return first, second, percentOf(pick.chance)
end

--- The source line from its parts (ItemTooltip.sourceParts).
function ItemTooltip.sourceText(first, second, percent)
	if first == nil then
		return nil
	end
	if second ~= nil and percent ~= nil then
		return string.format(L.tooltipSourceDrop, first, second, percent)
	end
	if second ~= nil then
		return string.format(L.tooltipSourcePair, first, second)
	end
	if percent ~= nil then
		return string.format(L.tooltipSourceChance, first, percent)
	end
	return first
end

--- The source line within `room` (L4.1 fix round 5: the mob is never cut): whole when
--- it fits; else without the drop chance; else the place (the instance) shortened with
--- an ellipsis, the boss always whole ("The Dead… · Captain Greenskin"). A source with
--- no boss is cut by measured width. `prefix` (a format with one %s, optional) wraps
--- the line ("· %s", the At line's own source row); `boe` as for ItemTooltip.sourceLine.
function ItemTooltip.fitSource(pick, measure, room, boe, prefix)
	local first, second, percent = ItemTooltip.sourceParts(pick, boe)
	if first == nil then
		return {}
	end
	local function text(place, other, chance)
		local line = ItemTooltip.sourceText(place, other, chance)
		return prefix ~= nil and string.format(prefix, line) or line
	end
	local fits = fitter(measure, room)
	local whole = text(first, second, percent)
	if fits(whole) then
		return { whole }
	end
	if percent ~= nil and fits(text(first, second, nil)) then
		return { text(first, second, nil) }
	end
	if second ~= nil then
		-- The place gives way, a letter at a time; the boss never does. When even one
		-- letter of the place leaves the line too wide, the boss takes a row of its own.
		local shortened = shorten(function(name) return text(name, second, nil) end, first, fits, 1)
		if fits(shortened) then
			return { shortened }
		end
		-- The boss is cut only when its name alone is wider than the whole tooltip, as the
		-- Next line's name is: the tooltip never widens.
		local place = prefix ~= nil and string.format(prefix, first) or first
		return { ItemTooltip.fitPlain(place, measure, room),
			ItemTooltip.fitPlain(string.format(L.tooltipAtSource, second), measure, room) }
	end
	return { ItemTooltip.fitPlain(text(first, nil, percent), measure, room) }
end

--- "Also good: <name>, <name>" from up to two names (already coloured, or
--- `false` for a name the client has not sent yet: the pending mark). nil
--- with none.
function ItemTooltip.alternativesLine(names)
	local shown = {}
	for index = 1, 2 do
		local name
		if names ~= nil then
			name = names[index]
		end
		if name == false then
			shown[#shown + 1] = L.tooltipNamePending
		elseif type(name) == "string" then
			shown[#shown + 1] = name
		end
	end
	if #shown == 2 then
		return string.format(L.tooltipAlternatives, shown[1], shown[2])
	end
	if #shown == 1 then
		return string.format(L.tooltipAlternativesOne, shown[1])
	end
	return nil
end

local function colored(name, hex)
	if hex == nil then
		return name
	end
	return "|cff" .. hex .. name .. "|r"
end

--- An alternative's gain as printed after its name, muted: "+1.1" ("+12" for
--- a score), the minus sign for a loss. nil with no gain.
function ItemTooltip.alternativeGain(gain, unit)
	if gain == nil then
		return nil
	end
	local sign = gain < 0 and L.tooltipMinus or "+"
	if unit == "score" then
		return string.format(L.tooltipAltGainScore, sign, math.abs(gain))
	end
	return string.format(L.tooltipAltGain, sign, math.abs(gain))
end

--- One alternative as it is printed: its name in its quality colour (`name`
--- given: the shortened one), then its gain muted when known; the pending
--- mark for a name not sent yet. `bare`: without its tag ("(needs bullets)"), for
--- the row the tag has left (ItemTooltip.fitAlternatives).
local function entryText(entry, name, bare)
	if type(entry) ~= "table" then
		return L.tooltipNamePending
	end
	local text = colored(name or entry.name, entry.hex)
	local gain = ItemTooltip.alternativeGain(entry.gain, entry.unit)
	if gain ~= nil then
		text = string.format(L.tooltipAlternativeGain, text, muted(gain))
	end
	-- A ranged alternative of the other ammo type, kept but flagged (L4.1 fix round 5).
	if entry.flag ~= nil and not bare then
		text = text .. " " .. muted(entry.flag)
	end
	return text
end

--- Ranged weapon subclass ids -> the ammo they fire (bows and crossbows arrows, guns
--- bullets); none for thrown weapons and wands.
ItemTooltip.AMMO_BY_SUBCLASS = { [2] = "arrows", [18] = "arrows", [3] = "bullets" }

--- "(needs bullets)" for an alternative firing other ammo than the worn ranged weapon,
--- from the two items' subclass ids (`wornSub`, `altSub`); nil when they match or either
--- fires none. Pure.
function ItemTooltip.ammoFlag(wornSub, altSub)
	local worn = ItemTooltip.AMMO_BY_SUBCLASS[tonumber(wornSub)]
	local alt = ItemTooltip.AMMO_BY_SUBCLASS[tonumber(altSub)]
	if worn == nil or alt == nil or worn == alt then
		return nil
	end
	return alt == "bullets" and L.tooltipNeedsBullets or L.tooltipNeedsArrows
end

--- One alternative's own rows, indented under the label (L.tooltipAlternativesRow): the
--- whole name, gain and tag when that fits; else the tag ("(needs bullets)") on a muted
--- row of its own under the name and gain; and only when the bare name and gain are still
--- wider than the room, the name shortened with an ellipsis (L4.1 fix round 7).
local function entryRows(entry, measure, room)
	local fits = fitter(measure, room)
	local row = string.format(L.tooltipAlternativesRow, entryText(entry))
	if type(entry) ~= "table" or fits(row) then
		return { row }
	end
	local rows = { shorten(function(name)
		return string.format(L.tooltipAlternativesRow, entryText(entry, name, true))
	end, entry.name, fits, 1) }
	if entry.flag ~= nil then
		rows[2] = ItemTooltip.fitPlain(string.format(L.tooltipAlternativesRow, muted(entry.flag)), measure, room)
	end
	return rows
end

--- The alternatives as printed lines, for `entries` ({ name, hex, gain, unit, flag }
--- each: the name in quality colour, the gain over what is worn muted; or
--- false for a name not sent yet), within `room` as `measure` reads it, so the
--- tooltip never widens: both on the label's row when they fit ("Also good: A +1.1,
--- B +0.6"); else "Also good:" alone and each alternative whole on its own indented
--- row ("Lil Timmy's Peashooter +2.2 (needs bullets)", then "Cracked Blacksmith Hammer
--- +0.8"), no row ending in a comma; a row still too wide puts its tag on a muted row
--- under it, and only a bare name and gain wider than the room are cut with an ellipsis
--- (L4.1 fix round 7, hunter-20). The one line with no measure or room. Empty with no
--- entries.
function ItemTooltip.fitAlternatives(entries, measure, room)
	entries = entries or {}
	local count = math.min(2, #entries)
	if count == 0 then
		return {}
	end
	local names = {}
	for index = 1, count do
		names[index] = type(entries[index]) == "table" and entryText(entries[index]) or false
	end
	local text = ItemTooltip.alternativesLine(names)
	local fits = fitter(measure, room)
	if fits(text) then
		return { text }
	end
	local lines = { ItemTooltip.fitPlain(L.tooltipAlternativesLabel, measure, room) }
	for index = 1, count do
		for _, row in ipairs(entryRows(entries[index], measure, room)) do
			lines[#lines + 1] = row
		end
	end
	return lines
end

local TAG_TEXT = { sim = "tooltipTagSim", build = "tooltipTagBuild", default = "tooltipTagEstimate" }

--- The section's first line: "Forever Sixty" in the client's own gold, then
--- the spec's tree muted, then where the weights came from: "Your sim" in
--- the tooltip's text colour, "Estimate" muted. The name alone with neither.
function ItemTooltip.headerText(specWord, tag)
	local text = Theme.inlineColor("clientGold", L.addonName)
	if specWord ~= nil then
		text = text .. muted(L.tooltipHeaderSep .. specWord)
	end
	local key = TAG_TEXT[tag]
	if key ~= nil then
		local tagColor = tag == "default" and "tooltipMuted" or "tooltipText"
		text = text .. muted(L.tooltipHeaderSep) .. Theme.inlineColor(tagColor, L[key])
	end
	return text
end

--- The header row within `room`, as { left, right, carry }: the tag ("Your
--- sim", "Estimate") and "Forever Sixty" are never cut. (a) The whole header
--- with "Equipped" right when left + gap + right fits; (b) else the whole header
--- alone, "Equipped" on its own right-aligned row under it (`carry`: L4.1 fix rounds 5-6,
--- the spec word stays on a narrow weapon tooltip); (c) else without the spec word,
--- "Equipped" beside it when it fits; (d) else (c)'s text alone and `carry`.
--- `equipped` false: no right text at all.
function ItemTooltip.fitHeader(specWord, tag, equipped, measure, room)
	local right = equipped and L.tooltipEquipped or nil
	local function fits(left, withRight)
		if type(measure) ~= "function" or room == nil then
			return true
		end
		local width = measure(left) or 0
		if withRight ~= nil then
			width = width + ItemTooltip.PAIR_GAP + (measure(withRight) or 0)
		end
		return width <= room
	end
	local whole = ItemTooltip.headerText(specWord, tag)
	if fits(whole, right) then
		return { left = whole, right = right }
	end
	-- L4.1 fix round 5 (both personas): the header keeps the spec word on every slot, a
	-- narrow weapon tooltip's too; "Equipped" moves to the next row before the word drops.
	if fits(whole) then
		return { left = whole, carry = right ~= nil }
	end
	local short = ItemTooltip.headerText(nil, tag)
	if fits(short, right) then
		return { left = short, right = right }
	end
	return { left = short, carry = right ~= nil }
end

--- "dps", "hps" or "score": the unit Gear.currentWeights' answer counts in.
function ItemTooltip.unitOf(current)
	if not current or current.dpsPerPoint == nil then
		return "score"
	end
	return current.healing and "hps" or "dps"
end

--- "<Stat> capped: more is wasted" for each stat the item carries that the
--- sim's weights message names as capped, by its plain name, sorted.
local function capTexts(stats, caps)
	local lines = {}
	if type(stats) ~= "table" or type(caps) ~= "table" then
		return lines
	end
	for _, stat in ipairs(caps) do
		if type(stats[stat]) == "number" and stats[stat] > 0 then
			lines[#lines + 1] = string.format(L.tooltipCapped, L.statNames[stat] or stat)
		end
	end
	table.sort(lines)
	return lines
end

--- The capped-stat lines for an item link, against `current`
--- (ItemTooltip.current's shape: its `caps` list).
function ItemTooltip.capLines(itemLink, current)
	if itemLink == nil or type(current) ~= "table" then
		return {}
	end
	return capTexts(Gear.statsOf(itemLink), current.caps)
end

--- The worn item's name, from its own link ("[Name]"), else the client.
local function wornNameOf(worn, nameOf)
	if worn == nil then
		return nil
	end
	local name = type(worn.link) == "string" and worn.link:match("%[(.-)%]") or nil
	return name or nameOf(worn.itemId)
end

--- Whether an item link carries an enchant (its enchant id, the field after
--- the item id, is not 0). Enchants are not item stats, so they are not
--- counted in any gain.
function ItemTooltip.isEnchanted(link)
	return Gear.isEnchanted(link)
end

--- The slot the hovered item is judged in, among `slots` (each one it may
--- be worn in): an empty one first, else the one whose worn item scores
--- lowest (of two rings, the weaker is the one to replace).
local function chooseSlot(slots, equipped, weights, weapon)
	local best, bestScore
	for _, slot in ipairs(slots) do
		if equipped[slot] == nil then
			return slot
		end
	end
	for _, slot in ipairs(slots) do
		local score = Progress.itemScore(equipped[slot].stats, weights, slot, weapon)
		if bestScore == nil or score < bestScore then
			best, bestScore = slot, score
		end
	end
	return best
end

--- The spec's own tree name ("Arms" for "warrior-arms"), Talents.tabName's.
local function specName(data, spec)
	if spec == nil then
		return nil
	end
	local classSlug, rest = spec:match("^([^-]+)%-(.+)$")
	local class = data.classes ~= nil and data.classes[classSlug] or nil
	for tab, entry in ipairs(class ~= nil and class.tabs or {}) do
		if Export.slugify(entry.name) == rest then
			return Talents.tabName(data, classSlug, tab)
		end
	end
	return Talents.titleCase(rest or spec)
end

--- The spec's own tree name, for the quest reward's "Best for Arms" (L4.2).
ItemTooltip.specName = specName

--- The lowest band at or above the character's own (any band below 20)
--- whose pick for one of `slots` is `itemId`, or nil.
function ItemTooltip.plannedBand(data, spec, faction, slots, itemId, level)
	local specBis = spec ~= nil and type(data.bis) == "table" and data.bis[spec] or nil
	if type(specBis) ~= "table" or not faction then
		return nil
	end
	local floor = Gear.bandFor(specBis, level) or 0
	local bands = {}
	for band in pairs(specBis) do
		if type(band) == "number" and band >= floor then
			bands[#bands + 1] = band
		end
	end
	table.sort(bands)
	for _, band in ipairs(bands) do
		for _, slot in ipairs(slots) do
			local pick = Gear.pick(spec, band, faction, slot, data)
			if pick ~= nil and pick.item_id == itemId then
				return band
			end
		end
	end
	return nil
end

--- The ladder's next rung after `band` in `specBis` (at 21, band 20: 30), or
--- nil in the top band.
function ItemTooltip.nextBand(specBis, band)
	local upcoming
	for candidate in pairs(type(specBis) == "table" and specBis or {}) do
		if type(candidate) == "number" and band ~= nil and candidate > band and (upcoming == nil or candidate < upcoming) then
			upcoming = candidate
		end
	end
	return upcoming
end

--- "Best until level 30" (in the tooltip's text colour): what is worn holds until
--- the ladder's next rung; "Best for now" (muted) in the top band.
function ItemTooltip.bestLine(nextBand)
	if nextBand ~= nil then
		return { text = string.format(L.tooltipBestUntil, nextBand), color = "tooltipText" }
	end
	return { text = L.tooltipBestForNow, color = "muted" }
end

--- Where a pick comes from in a word or two: the instance for a dungeon or raid
--- drop ("Scarlet Monastery"), else the kind's own word ("Vendor", "Quest",
--- "Crafted"). No boss, no drop chance. nil with no source.
function ItemTooltip.shortSource(pick)
	if type(pick) ~= "table" or type(pick.source) ~= "string" or pick.source == "" then
		return nil
	end
	local place = pick.source:match("^(.-): .+$")
	if INSTANCE_KINDS[pick.kind] then
		return place or pick.source
	end
	return ItemTooltip.SOURCE_KIND_NAMES[pick.kind] or place or pick.source
end

--- "At level 30: <name> · <source>" (the source as the Next line's: boss and drop
--- chance included, L4.1 fix round 5) (muted, the name in its quality
--- colour, its gain over what is worn in green when it reaches the floor): the
--- next rung's pick under "Best until level 30". `withSource` false: without the
--- source; `withGain` false: without the gain.
function ItemTooltip.atText(at, withSource, withGain)
	local shown = colored(at.name, at.hex)
	if at.gain ~= nil and withGain ~= false then
		-- The next rung's pick's gain over what is worn, in the upgrade green.
		shown = shown .. " " .. ItemTooltip.figureText(at.gain, at.unit)
	end
	if at.source ~= nil and withSource ~= false then
		return string.format(L.tooltipAtBand, at.band, shown, at.source)
	end
	return string.format(L.tooltipAtBandName, at.band, shown)
end

--- The At line within `room`: whole when it fits, else the name on its line and
--- the source on a muted second line ("· Scarlet Monastery", itself fitted). The
--- name is shortened only when even the bare name is wider than the room.
function ItemTooltip.fitAt(at, measure, room)
	local fits = fitter(measure, room)
	local whole = ItemTooltip.atText(at)
	if fits(whole) then
		return { whole }
	end
	-- The source moves to its own line first, then the gain goes, and the name is
	-- cut only when even the bare name is wider than the room.
	local bare = ItemTooltip.atText(at, false)
	if not fits(bare) then
		bare = ItemTooltip.atText(at, false, false)
	end
	if not fits(bare) then
		bare = shorten(function(name)
			return ItemTooltip.atText({ band = at.band, name = name, hex = at.hex }, false)
		end, at.name, fits, 1)
	end
	if at.source == nil then
		return { bare }
	end
	-- The source row is fitted as every source line is (ItemTooltip.fitSource: the chance
	-- goes first, then the place shortens, never the boss).
	if at.pick ~= nil then
		local rows = { bare }
		for _, row in ipairs(ItemTooltip.fitSource(at.pick, measure, room, at.boe, L.tooltipAtSource)) do
			rows[#rows + 1] = row
		end
		return rows
	end
	return { bare, ItemTooltip.fitPlain(string.format(L.tooltipAtSource, at.source), measure, room) }
end

--- A figure as printed in the upgrade green: "+12.3 DPS" ("+12 score").
function ItemTooltip.figureText(gain, unit)
	local figure
	if unit == "score" then
		figure = string.format(L.tooltipNextGainScore, gain)
	else
		figure = string.format(L.tooltipNextGain, gain, unitWord(unit))
	end
	return Theme.inlineColor("success", figure)
end

--- The client global a stat's word is read from first, where several of its
--- ITEM_MOD_*_SHORT keys map to it: the generic hit and crit, never the melee,
--- ranged or spell variant.
ItemTooltip.STAT_WORD_KEYS = {
	hit = { "ITEM_MOD_HIT_RATING_SHORT" },
	crit = { "ITEM_MOD_CRIT_RATING_SHORT" },
}

--- A stat in the player's words: "Str", "Agi", "Sta", "Int", "Spi" for the
--- five primary stats (Locale), "weapon DPS" for a weapon's own DPS, else the
--- client's own name for it (its ITEM_MOD_*_SHORT global: the generic one first
--- where a stat has several, STAT_WORD_KEYS; else the alphabetical first), else
--- Locale's. A caster never reads "Melee Hit Rating" for hit.
function ItemTooltip.statWord(stat)
	if L.statAbbrev[stat] ~= nil then
		return L.statAbbrev[stat]
	end
	if stat == "weapon_dps" then
		return L.tooltipWeaponDps
	end
	local keys = {}
	for clientKey, value in pairs(Rules.STAT_KEYS) do
		if value == stat then
			keys[#keys + 1] = clientKey
		end
	end
	table.sort(keys)
	for index = #(ItemTooltip.STAT_WORD_KEYS[stat] or {}), 1, -1 do
		table.insert(keys, 1, ItemTooltip.STAT_WORD_KEYS[stat][index])
	end
	for _, clientKey in ipairs(keys) do
		if type(_G[clientKey]) == "string" and _G[clientKey] ~= "" then
			return _G[clientKey]
		end
	end
	return L.statNames[stat] or stat
end

--- Display order: the five primary stats first, then the rest, a weapon's DPS last.
local BREAKDOWN_ORDER = { strength = 1, agility = 2, stamina = 3, intellect = 4, spirit = 5, weapon_dps = 99 }

--- The breakdown line under a figure: the `count` biggest contributors to it
--- (Progress.contributions, by weighted size), signed, in the player's words,
--- in a fixed order: "(−11 Str, −4 Agi, −6.9 weapon DPS)". nil with none.
function ItemTooltip.breakdownText(contributions, count)
	local chosen = {}
	for index = 1, math.min(count or 3, #(contributions or {})) do
		chosen[index] = contributions[index]
	end
	if #chosen == 0 then
		return nil
	end
	table.sort(chosen, function(left, right)
		local leftRank, rightRank = BREAKDOWN_ORDER[left.stat] or 50, BREAKDOWN_ORDER[right.stat] or 50
		if leftRank ~= rightRank then
			return leftRank < rightRank
		end
		return math.abs(left.weighted) > math.abs(right.weighted)
	end)
	local parts = {}
	for index, entry in ipairs(chosen) do
		local size = math.abs(entry.delta)
		local amount
		if entry.stat ~= "weapon_dps" and math.abs(size - math.floor(size + 0.5)) < 1e-9 then
			amount = string.format("%d", math.floor(size + 0.5))
		else
			amount = string.format("%.1f", size)
		end
		local sign = entry.delta < 0 and L.tooltipMinus or "+"
		parts[index] = string.format(L.tooltipBreakdownPart, sign .. amount, ItemTooltip.statWord(entry.stat))
	end
	return string.format(L.tooltipBreakdown, table.concat(parts, L.tooltipBreakdownJoin))
end

--- The breakdown line within `room`: drops the third, then the second
--- contributor, then cuts the line by measured width.
function ItemTooltip.fitBreakdown(contributions, measure, room)
	local fits = fitter(measure, room)
	for count = 3, 1, -1 do
		local text = ItemTooltip.breakdownText(contributions, count)
		if text ~= nil and fits(text) then
			return { text }
		end
	end
	local one = ItemTooltip.breakdownText(contributions, 1)
	return one ~= nil and { ItemTooltip.fitPlain(one, measure, room) } or {}
end

--- The breakdown line for a verdict that carries a figure, or nil.
local function breakdownLine(verdict)
	local contributions = verdict ~= nil and verdict.contributions or nil
	if contributions == nil or #contributions == 0 then
		return nil
	end
	return { kind = "breakdown", contributions = contributions, color = "muted",
		text = ItemTooltip.breakdownText(contributions, 3) }
end

--- The gain part of the Next line: the figure in the upgrade green, then
--- "(enchant not counted)" muted when the worn item's enchant could not be
--- read (`unparsed`; a read enchant is already netted, Gear.readEquipped).
local function gainPart(gain, unit, unparsed)
	local text = ItemTooltip.figureText(gain, unit)
	if unparsed then
		text = text .. " " .. muted(L.tooltipEnchantNote)
	end
	return text
end

--- "Next upgrade: <name> +6.1 DPS": the label muted, the name as given (in its
--- quality colour), the gain in the upgrade green, already rounded
--- (Progress.roundedGain) and at the floor or above; the enchant note muted
--- after it only when the worn item's enchant could not be read. With no gain
--- (an empty slot: nothing worn to gain over, L4.1 ruling 5) the label and the name.
function ItemTooltip.nextLine(name, gain, unit, unparsed)
	if gain == nil then
		return muted(L.tooltipNextLabel) .. " " .. name
	end
	return muted(L.tooltipNextLabel) .. " " .. name .. " " .. gainPart(gain, unit, unparsed)
end

--- The Next line within `room`: one line when it fits; else "Next upgrade:
--- <name>" (the name shortened only if that line is still too wide) and, on
--- the second line, the gain (and the enchant note). The gain is never cut.
--- `upgrade` is { name, hex, gain, unit, unparsed }; with no gain, one line.
function ItemTooltip.fitNext(upgrade, measure, room)
	local one = ItemTooltip.nextLine(colored(upgrade.name, upgrade.hex), upgrade.gain, upgrade.unit, upgrade.unparsed)
	local fits = fitter(measure, room)
	if fits(one) then
		return { one }
	end
	local head = shorten(function(name) return muted(L.tooltipNextLabel) .. " " .. colored(name, upgrade.hex) end,
		upgrade.name, fits, 1)
	if upgrade.gain == nil then
		return { head }
	end
	return { head, gainPart(upgrade.gain, upgrade.unit, upgrade.unparsed) }
end

--- The plan line within `room` (L4.1 fix round 5): whole when it fits; else "Planned for
--- your build" (" at 20" when the line names a band) rather than a cut build name; that
--- too cut by measured width only when even it is wider than the room.
function ItemTooltip.fitPlanned(line, measure, room)
	local fits = fitter(measure, room)
	if fits(line.text) then
		return line.text
	end
	local short = line.band ~= nil and string.format(L.tooltipPlannedYourBuildAt, line.band) or L.tooltipPlannedYourBuild
	return ItemTooltip.fitPlain(short, measure, room)
end

--- The arrow slot's item line (L4.1 fix round 5): "<name> +0.8 DPS", the name in its
--- quality colour and the gain in the upgrade green; the name shortened before the gain is
--- ever cut. `upgrade` is { name, hex, gain, unit }.
function ItemTooltip.bagItemText(upgrade, name)
	return colored(name or upgrade.name, upgrade.hex) .. " " .. ItemTooltip.figureText(upgrade.gain, upgrade.unit)
end

function ItemTooltip.fitBagLine(upgrade, measure, room)
	local fits = fitter(measure, room)
	return shorten(function(name) return ItemTooltip.bagItemText(upgrade, name) end, upgrade.name, fits, 1)
end

--- The section of a paper-doll slot that carries the arrow (L4.1 fix round 5): the bag fact
--- leads and stands alone: "In your bags" (muted), "<item> +0.8 DPS" (the gain the chip's list
--- gives, Progress.bagUpgrades, as shown), "Right-click it in your bags to equip." (muted); no
--- source line. The compare tooltip shows the bag item, titled "Next upgrade". `found` is the
--- chip's list entry ({ link, itemId, delta, unit }); `head` the hover's own header.
function ItemTooltip.bagSection(found, head, unit, qualityHex)
	local name = type(found.link) == "string" and found.link:match("%[(.-)%]") or Compat.itemName(found.itemId)
	if name == nil then
		return nil
	end
	local shownUnit = found.unit == "dps" and (unit == "hps" and "hps" or "dps") or "score"
	local upgrade = { name = name, hex = (qualityHex or ItemTooltip.qualityHexOf)(found.itemId), unit = shownUnit,
		gain = Progress.roundedGain(found.delta, found.unit == "score" and "score" or "dps") }
	return {
		lines = {
			{ text = L.tooltipInYourBags, color = "muted" },
			{ kind = "bag", upgrade = upgrade, color = "body", text = ItemTooltip.bagItemText(upgrade) },
			{ text = L.tooltipRightClickEquip, color = "muted" },
		},
		target = { itemId = found.itemId, link = found.link, label = L.tooltipCompareLabel },
		pending = {},
		head = head,
	}
end

--- Injection point (surfaces/SlotMarks.lua sets it): the chip's list entry that puts an
--- arrow on a paper-doll slot, or nil. The slot hovers read it to lead with the bag fact.
ItemTooltip.bagUpgradeFor = nil

local function buildPlans(build, slots, itemId)
	if type(build) ~= "table" then
		return false
	end
	local wanted = {}
	for _, slot in ipairs(slots) do
		wanted[slot] = true
	end
	for _, entry in ipairs(build.gear or {}) do
		if wanted[entry.slot] and entry.itemId == itemId then
			return true
		end
	end
	return false
end

--- The client reads `ctx` falls back to when a caller hands none in.
local function itemStatsOf(itemId)
	local _, link = Compat.itemInfo(itemId)
	return link ~= nil and Gear.statsOf(link) or nil
end

local function qualityHexOf(itemId)
	local quality = Compat.itemQuality(itemId)
	if quality == nil then
		return nil
	end
	return Theme.hexOf(Theme.qualityColor(quality))
end
ItemTooltip.qualityHexOf = qualityHexOf

--- Item ids this client cannot show at all (it answered GET_ITEM_INFO_RECEIVED
--- with a failure, or says the item does not exist): a band pick the data
--- names but the client has never heard of (seen in game: The Axe of
--- Severing, a pick the data marks client_unconfirmed, never resolved).
ItemTooltip.missing = {}

local function missingOf(itemId)
	return ItemTooltip.missing[itemId] == true or Compat.itemExists(itemId) == false
end

--- Whether this client cannot show `itemId` at all (a failed GET_ITEM_INFO_RECEIVED, or
--- an id it has never heard of): the section's own rule, for the slot marks too.
ItemTooltip.isMissing = missingOf

--- A pick's name: the client's, else the data's (Data.bisNames) while the client
--- has not sent it (`pending`: the hover redraws with the client's own when it
--- comes). nil for neither.
local function pickNameOf(data, nameOf, pending, itemId)
	local name = nameOf(itemId)
	local bisNames = type(data.bisNames) == "table" and data.bisNames or {}
	if name == nil and type(bisNames[itemId]) == "string" then
		pending[itemId] = true
		return bisNames[itemId]
	end
	return name
end

--- The pick the section goes and gets: the band's pick, or, when the client
--- cannot show it, the first of its alternatives it can (with the rest as
--- its own alternatives). nil when it can show none.
local function reachablePick(pick, isMissing)
	if pick == nil or not isMissing(pick.item_id) then
		return pick
	end
	local alternatives = pick.alternatives or {}
	for index, alternative in ipairs(alternatives) do
		if alternative.item_id ~= nil and not isMissing(alternative.item_id) then
			local rest = {}
			for other = index + 1, #alternatives do
				rest[#rest + 1] = alternatives[other]
			end
			return { item_id = alternative.item_id, kind = alternative.kind, source = alternative.source,
				alternatives = rest }
		end
	end
	return nil
end

--- The part of the section that scores one item (L4.2 ruling 1): the one verdict the
--- item tooltip, the quest reward labels and the roll line all read, so they can never
--- disagree. `ctx` is ItemTooltip.section's; `pending` (optional) gathers the worn item
--- whose name the client has not sent. Answers { gear, slots, wornSlot, wearable, usable,
--- refusal, blocked, heldByTwoHander, slot, verdict, primary, unit }:
---
--- - `gear`: the item goes in a slot the plan covers (none: nothing else is set);
--- - `wearable`: the slots Rules.canEquip allows now (none: `refusal`, "Not for you until
---   40: plate armor", is the primary line, in the warning colour);
--- - `heldByTwoHander`: every allowed slot is a hand the worn two-hander holds ("Not with
---   your two-hander");
--- - `slot` the slot it is judged in, `verdict` Progress.itemVerdict's answer there (on an
---   item not worn, with weights), and `primary` the verdict line (the 1.0 floor, the
---   empty-slot fill and the two-hander pair rule, ItemTooltip.verdictLine's own).
function ItemTooltip.judge(ctx, pending)
	pending = pending or {}
	local item = ctx.item or {}
	local slots = Gear.slotsFor(item.equipLocation)
	if #slots == 0 then
		return { gear = false, slots = slots, wearable = {}, usable = {} }
	end
	local equipped = ctx.equipped or {}
	local nameOf = ctx.nameOf or Compat.itemName
	local current = ctx.current or nil
	local weights = current ~= nil and current.weights or nil
	local perPoint = current ~= nil and current.dpsPerPoint or nil
	local weapon = current ~= nil and current.weapon or nil
	local unit = ItemTooltip.unitOf(current)
	local hovered = { itemId = item.itemId, link = item.link, equipLocation = item.equipLocation, stats = item.stats }

	local wornSlot
	for _, slot in ipairs(slots) do
		local worn = equipped[slot]
		if worn ~= nil and item.link ~= nil and worn.link == item.link then
			wornSlot = slot
			break
		end
	end
	-- Rules first, then the two-hander: the slots the item may go in now.
	local wearable, refusal = {}, nil
	for _, slot in ipairs(slots) do
		local answer, unlock, reason = Rules.canEquip(ctx.classToken, ctx.level, item.itemClass,
			item.itemSubClass, item.equipLocation, slot)
		if answer == "ok" then
			wearable[#wearable + 1] = slot
		elseif refusal == nil then
			refusal = ItemTooltip.untilLine(ctx.classToken, answer, unlock, reason)
		end
	end
	local usable, blocked = {}, nil
	for _, slot in ipairs(wearable) do
		local verdict = Progress.itemVerdict(slot, hovered, equipped, weights, perPoint, weapon)
		if verdict ~= nil and verdict.state == "blocked" then
			blocked = verdict
		else
			usable[#usable + 1] = slot
		end
	end
	local heldByTwoHander = #wearable > 0 and #usable == 0
	local slot = wornSlot or chooseSlot(#usable > 0 and usable or slots, equipped, weights, weapon)

	local primary, verdict
	if #wearable == 0 then
		primary = refusal ~= nil and { text = refusal, color = "warning" } or nil
	elseif heldByTwoHander then
		primary = ItemTooltip.verdictLine(blocked, unit, nil)
	elseif wornSlot == nil and weights ~= nil then
		verdict = Progress.itemVerdict(slot, hovered, equipped, weights, perPoint, weapon)
		local worn = equipped[slot]
		if slot == "main_hand" then
			-- A two-hander scored net of both hands is named over the main hand.
			worn = equipped.main_hand or equipped.off_hand
		end
		local wornName = worn ~= nil and wornNameOf(worn, nameOf) or nil
		if worn ~= nil and wornName == nil then
			pending[worn.itemId] = true
		end
		primary = ItemTooltip.verdictLine(verdict, unit, wornName)
		if primary ~= nil then
			-- The hook fits the compared name to the tooltip's width (fitVerdict).
			primary.kind, primary.verdict, primary.unit, primary.wornName = "verdict", verdict, unit, wornName
			-- Under a figure, what it is made of (a plain "Sidegrade" has no figure).
			if primary.text ~= L.tooltipSide then
				primary.breakdown = breakdownLine(verdict)
			end
		end
	end
	return {
		gear = true, slots = slots, wornSlot = wornSlot, wearable = wearable, usable = usable,
		refusal = refusal, blocked = blocked, heldByTwoHander = heldByTwoHander, slot = slot,
		verdict = verdict, primary = primary, unit = unit,
	}
end

--- The section for one hovered item: the lines ({ text, color, kind? }) in
--- draw order, the compare tooltip's item ({ itemId, link, label } for the
--- next upgrade, only once the client holds it), the ids whose names the
--- client has not sent yet (a set: the hover redraws when one arrives), and
--- the header ({ specWord, tag, equipped }). `ctx` holds what the hook read:
---
--- - `data`, `spec`, `classToken`, `level`, `faction`;
--- - `item`: { itemId, link, equipLocation, itemClass, itemSubClass, stats };
--- - `current`: Gear.currentWeights' answer (nil: no weights for the spec);
--- - `equipped`: Gear.readEquipped()'s table; `build`: the loaded build;
--- - `nameOf(id)`, `statsOf(id)`, `isCached(id)`, `qualityHex(id)`,
---   `isMissing(id)`: the client's name, stats, cache state, quality colour
---   and "cannot show it at all" for an item.
---
--- Every piece of gear in a slot the plan covers gets a primary line: the
--- refusal (Rules first: never called an upgrade), "Not with your
--- two-hander" (alone: no plan, no next pick under it), the verdict, on the
--- item worn (the exact link) the next upgrade or "Best until N", else "Not
--- for your plan". One source block per tooltip: the hovered item's own when
--- it is the band's pick, else the next upgrade's; the next upgrade is
--- printed only with its gain from the floor up, and skips a pick the client
--- cannot show for the first alternative it can. A pick (or an alternative, or
--- the next rung's pick) the client has not sent yet is named from the data
--- (Data.bisNames, never "…") and, for a weapon, priced from its own DPS
--- (Data.bisDps) until the client's stats arrive; its id stays pending so the
--- hover redraws when they do. One with no name anywhere is waited for.
function ItemTooltip.section(ctx)
	local lines, pending = {}, {}
	local item = ctx.item or {}
	local slots = Gear.slotsFor(item.equipLocation)
	if #slots == 0 then
		return lines, nil, pending, nil
	end
	local data = ctx.data or {}
	local equipped = ctx.equipped or {}
	local nameOf = ctx.nameOf or Compat.itemName
	local statsOf = ctx.statsOf or itemStatsOf
	local isCached = ctx.isCached or Compat.isItemCached
	local qualityHex = ctx.qualityHex or qualityHexOf
	local isMissing = ctx.isMissing or missingOf
	local current = ctx.current or nil
	local weights = current ~= nil and current.weights or nil
	local perPoint = current ~= nil and current.dpsPerPoint or nil
	local weapon = current ~= nil and current.weapon or nil
	local unit = ItemTooltip.unitOf(current)
	local bisDps = type(data.bisDps) == "table" and data.bisDps or {}
	local function pickName(itemId)
		return pickNameOf(data, nameOf, pending, itemId)
	end
	--- A pick's stats: the client's, else a weapon's own DPS from the data until
	--- the client's arrive (pending, as above). nil for neither.
	local function pickStats(itemId)
		local stats = statsOf(itemId)
		if (stats == nil or next(stats) == nil) and type(bisDps[itemId]) == "number" then
			pending[itemId] = true
			return { weapon_dps = bisDps[itemId] }
		end
		return stats
	end
	--- Whether an item binds when equipped (the client's bind type 2): a world drop of one
	--- is on the auction house (L4.1 fix round 5).
	local function isBoe(itemId)
		return (ctx.bindTypeOf or Compat.itemBindType)(itemId) == 2
	end
	--- An item's subclass id (a ranged weapon's: bow, gun, crossbow), from the client.
	local function subClassOf(itemOrLink)
		if ctx.subClassOf ~= nil then
			return ctx.subClassOf(itemOrLink)
		end
		return select(7, Compat.itemInfoInstant(itemOrLink))
	end
	local wornRanged = equipped.ranged
	local wornSub = wornRanged ~= nil and subClassOf(wornRanged.link or wornRanged.itemId) or nil

	local judged = ItemTooltip.judge(ctx, pending)
	local wornSlot, usable, slot = judged.wornSlot, judged.usable, judged.slot
	local heldByTwoHander = judged.heldByTwoHander
	local head = { specWord = specName(data, ctx.spec), tag = current ~= nil and current.tag or nil,
		equipped = wornSlot ~= nil, slot = wornSlot }
	local primary = judged.primary
	-- A hand the worn two-hander holds says only that: no plan, no next pick.
	local after = {}
	if not heldByTwoHander then
		if #usable > 0 and current ~= nil then
			for _, text in ipairs(capTexts(item.stats, current.caps)) do
				after[#after + 1] = { text = text, color = "body" }
			end
		end
		local band = ItemTooltip.plannedBand(data, ctx.spec, ctx.faction, slots, item.itemId, ctx.level)
		if band ~= nil then
			after[#after + 1] = { kind = "planned", band = band, color = "gold",
				text = string.format(L.tooltipPlanned, specName(data, ctx.spec), band) }
		elseif buildPlans(ctx.build, slots, item.itemId) then
			after[#after + 1] = { kind = "planned", color = "gold",
				text = string.format(L.tooltipPlannedBuild, Follow.buildName(data, ctx.build)) }
		end
	end

	-- The source block: the band's pick for the slot.
	local specBis = ctx.spec ~= nil and type(data.bis) == "table" and data.bis[ctx.spec] or nil
	local currentBand = ctx.faction and Gear.bandFor(specBis, ctx.level) or nil
	local nextBand = ItemTooltip.nextBand(specBis, currentBand)
	local roundUnit = unit == "score" and "score" or "dps"
	--- A candidate's gain over what is worn in the slot (net of the worn
	--- enchant, Gear.readEquipped), rounded as shown; nil with no figure.
	local function gainOf(itemId)
		local stats = weights ~= nil and pickStats(itemId) or nil
		if stats == nil or next(stats) == nil then
			return nil
		end
		local wanted = { itemId = itemId, stats = stats, equipLocation = select(4, Compat.itemInfoInstant(itemId)) }
		local verdict = Progress.itemVerdict(slot, wanted, equipped, weights, perPoint, weapon)
		if verdict == nil or verdict.delta == nil then
			return nil
		end
		return Progress.roundedGain(-verdict.delta, roundUnit), verdict
	end
	local pick, isPick, pickWorn, unshowable
	if currentBand ~= nil and not heldByTwoHander then
		for _, candidate in ipairs(slots) do
			local found = Gear.pick(ctx.spec, currentBand, ctx.faction, candidate, data)
			if found ~= nil and found.item_id == item.itemId then
				pick, isPick = found, true
				break
			end
		end
		if pick == nil then
			local named = Gear.pick(ctx.spec, currentBand, ctx.faction, slot, data)
			pick = reachablePick(named, isMissing)
			-- The data names a pick but the client can show none of it.
			unshowable = named ~= nil and pick == nil
		end
		for _, candidate in ipairs(slots) do
			local worn = equipped[candidate]
			if pick ~= nil and worn ~= nil and worn.itemId == pick.item_id then
				pickWorn = true
			end
		end
	end
	local block, nextLine, target = {}, nil, nil
	local holds = false
	if unshowable and wornSlot ~= nil then
		-- What is worn holds, never a silent slot.
		primary = { text = L.tooltipBestKnown, color = "muted" }
	elseif pick ~= nil and not (isPick and wornSlot == nil) and pickWorn then
		holds = wornSlot ~= nil
	elseif pick ~= nil then
		local showBlock = true
		if not isPick then
			-- The pick's gain over what is worn, from its own stats and weapon DPS,
			-- net of the worn enchant, as shown: a next upgrade from the floor.
			local name = pickName(pick.item_id)
			local gain, nextVerdict = gainOf(pick.item_id)
			if gain ~= nil and gain >= Progress.MIN_GAIN and name ~= nil then
				local worn = equipped[slot]
				local upgrade = { name = name, hex = qualityHex(pick.item_id), gain = gain, unit = unit,
					unparsed = worn ~= nil and worn.enchantUnparsed == true }
				nextLine = { kind = "next", itemId = pick.item_id, upgrade = upgrade, color = "body",
					text = ItemTooltip.nextLine(colored(name, upgrade.hex), gain, unit, upgrade.unparsed) }
				nextLine.breakdown = breakdownLine(nextVerdict)
				if isCached(pick.item_id) then
					target = { itemId = pick.item_id, link = Compat.displayLink(pick.item_id),
						label = L.tooltipCompareLabel }
				end
			elseif name == nil and (gain == nil or gain >= Progress.MIN_GAIN) then
				-- Not sent yet: wait for it (the hover redraws), never print "…".
				pending[pick.item_id] = true
				showBlock = false
			else
				-- Under the floor, or no figure to give: nothing to go and get.
				holds = wornSlot ~= nil
				showBlock = false
			end
		end
		if showBlock then
			local boe = isBoe(pick.item_id)
			local source = ItemTooltip.sourceLine(pick, boe)
			if source ~= nil then
				block[#block + 1] = { kind = "source", pick = pick, boe = boe, text = source, color = "muted" }
			end
			local entries = {}
			-- Alternatives only under a next upgrade: they are the other ways to it.
			local alternatives = nextLine ~= nil and pick.alternatives or {}
			for _, alternative in ipairs(alternatives) do
				local id = alternative.item_id
				if id ~= nil and id ~= item.itemId and #entries < 2 and not isMissing(id) then
					local name = pickName(id)
					if name == nil then
						pending[id] = true
						entries[#entries + 1] = false
					else
						entries[#entries + 1] = { name = name, hex = qualityHex(id), gain = gainOf(id), unit = unit,
							flag = slot == "ranged" and wornSub ~= nil and ItemTooltip.ammoFlag(wornSub, subClassOf(id))
								or nil }
					end
				end
			end
			if #entries > 0 then
				block[#block + 1] = { kind = "alternatives", entries = entries, color = "muted",
					text = ItemTooltip.fitAlternatives(entries)[1] }
			end
		end
	end
	-- What is worn holds until the ladder's next rung: "Best until level 30", and
	-- muted under it the next rung's pick for the slot and where it comes from.
	local atLine
	if holds then
		primary = ItemTooltip.bestLine(nextBand)
		local upcoming = nextBand ~= nil and Gear.pick(ctx.spec, nextBand, ctx.faction, slot, data) or nil
		if upcoming ~= nil and not isMissing(upcoming.item_id) then
			local name = pickName(upcoming.item_id)
			if name == nil then
				pending[upcoming.item_id] = true
			else
				local boe = isBoe(upcoming.item_id)
				atLine = { kind = "at", band = nextBand, name = name, hex = qualityHex(upcoming.item_id),
					pick = upcoming, boe = boe, source = ItemTooltip.sourceLine(upcoming, boe), color = "muted" }
				local gain = gainOf(upcoming.item_id)
				if gain ~= nil and gain >= Progress.MIN_GAIN then
					atLine.gain, atLine.unit = gain, unit
				end
				atLine.text = ItemTooltip.atText(atLine)
			end
		end
	end

	-- On the item worn the next upgrade leads; on any other it follows the plan.
	if wornSlot ~= nil and nextLine ~= nil then
		primary = nextLine
		nextLine = nil
	end
	if primary == nil and pick == nil and wornSlot == nil and not isPick and next(pending) == nil then
		primary = { text = L.tooltipNotForPlan, color = "muted" }
	end
	-- Every worn item gets a line (L4.1 fix round 3, hunter-9-mainhand): one the band names
	-- no pick for (or below the floor, with no band at all) says so rather than a bare
	-- "Equipped". A hand the worn two-hander holds already said so above.
	if primary == nil and wornSlot ~= nil and pick == nil and not unshowable and not heldByTwoHander
		and next(pending) == nil then
		primary = { text = L.tooltipNoPick, color = "muted" }
	end
	--- A line, then the breakdown under its figure when it has one.
	local function push(line)
		lines[#lines + 1] = line
		if line.breakdown ~= nil then
			lines[#lines + 1] = line.breakdown
		end
	end
	if primary ~= nil then
		push(primary)
	end
	if atLine ~= nil then
		lines[#lines + 1] = atLine
	end
	for _, line in ipairs(after) do
		lines[#lines + 1] = line
	end
	if nextLine ~= nil then
		push(nextLine)
	end
	for _, line in ipairs(block) do
		lines[#lines + 1] = line
	end
	return lines, target, pending, head
end

-- ---------------------------------------------------------------------
-- The empty character-slot hover (L4.1 ruling 5): the section with the
-- band's pick for the slot, as the worn item prints its Next line.
-- ---------------------------------------------------------------------

--- Site slot name -> the client's paper-doll button's global frame name (shirt and
--- tabard carry no stats and are left out). The empty-slot hover and the slot marks
--- (surfaces/SlotMarks.lua) both find the client's buttons through this table.
ItemTooltip.SLOT_BUTTONS = {
	head = "CharacterHeadSlot",
	neck = "CharacterNeckSlot",
	shoulder = "CharacterShoulderSlot",
	back = "CharacterBackSlot",
	chest = "CharacterChestSlot",
	wrist = "CharacterWristSlot",
	hands = "CharacterHandsSlot",
	waist = "CharacterWaistSlot",
	legs = "CharacterLegsSlot",
	feet = "CharacterFeetSlot",
	finger1 = "CharacterFinger0Slot",
	finger2 = "CharacterFinger1Slot",
	trinket1 = "CharacterTrinket0Slot",
	trinket2 = "CharacterTrinket1Slot",
	main_hand = "CharacterMainHandSlot",
	off_hand = "CharacterSecondaryHandSlot",
	ranged = "CharacterRangedSlot",
}

--- The section for an EMPTY paper-doll slot, in ItemTooltip.section's shape (lines,
--- the compare tooltip's item, the ids still pending, the header): the band's pick
--- for `slot` (Progress.slotPicks' pair rule: a pick worn in either slot of a pair is
--- not offered, and one two slots share is offered on the first), the first of its
--- alternatives the client can show when it cannot show the pick, printed exactly as
--- the worn item prints its Next line, with no gain (nothing is worn to gain over),
--- then its source line. `ctx` is ItemTooltip.section's without `item`. No lines below
--- the band floor, with no faction, or with no pick for the slot.
function ItemTooltip.emptySlotSection(ctx, slot)
	local lines, pending = {}, {}
	if not ctx.spec then
		-- No spec to read a band from: nothing to say at all.
		return lines, nil, pending, nil
	end
	local data = ctx.data or {}
	local current = ctx.current or nil
	local equipped = ctx.equipped or {}
	local head = { specWord = specName(data, ctx.spec), tag = current ~= nil and current.tag or nil, equipped = false }
	-- An empty off hand under a worn two-hander: nothing goes there (fix round 1, N1).
	if slot == "off_hand" and equipped.main_hand ~= nil and equipped.main_hand.equipLoc == "INVTYPE_2HWEAPON" then
		lines[1] = { text = L.tooltipTwoHander, color = "muted" }
		return lines, nil, pending, head
	end
	-- Never silence (fix round 1): no pick of its own for the band (none in the data,
	-- below the floor, no faction, or the pair rule leaves it none) says so, muted.
	local noPick = { text = L.tooltipNoPick, color = "muted" }
	local specBis = type(data.bis) == "table" and data.bis[ctx.spec] or nil
	local band = ctx.faction and Gear.bandFor(specBis, ctx.level) or nil
	if band == nil then
		lines[1] = noPick
		return lines, nil, pending, head
	end
	local isMissing = ctx.isMissing or missingOf
	local picks = Progress.slotPicks(data, ctx.spec, band, ctx.faction, equipped)
	local pick = reachablePick(picks[slot], isMissing)
	if pick == nil then
		lines[1] = noPick
		return lines, nil, pending, head
	end
	local itemId = pick.item_id
	local name = pickNameOf(data, ctx.nameOf or Compat.itemName, pending, itemId)
	if name == nil then
		-- Not sent yet: wait for it (the hover redraws), never print "…".
		pending[itemId] = true
		return lines, nil, pending, head
	end
	local upgrade = { name = name, hex = (ctx.qualityHex or qualityHexOf)(itemId), unit = ItemTooltip.unitOf(current) }
	lines[1] = { kind = "next", itemId = itemId, upgrade = upgrade, color = "body",
		text = ItemTooltip.nextLine(colored(name, upgrade.hex)) }
	local source = ItemTooltip.sourceLine(pick)
	if source ~= nil then
		lines[2] = { kind = "source", pick = pick, text = source, color = "muted" }
	end
	local target
	if (ctx.isCached or Compat.isItemCached)(itemId) then
		target = { itemId = itemId, link = Compat.displayLink(itemId), label = L.tooltipCompareLabel }
	end
	return lines, target, pending, head
end

-- ---------------------------------------------------------------------
-- The hooks.
-- ---------------------------------------------------------------------

--- The generated data table, set by Options.register().
ItemTooltip.data = nil

--- Gear.currentWeights' answer for the character and spec logged in, set by
--- Options.refreshWeightsMessage (at login, on a respec, when the inbox is
--- read). nil before that, when the hook asks Gear.currentWeights itself.
ItemTooltip.current = nil

--- The section, computed once per item link until something it reads
--- changes (Options resets it with every surface refresh and every change
--- of what is worn).
ItemTooltip.cache = {}

function ItemTooltip.resetCache()
	ItemTooltip.cache = {}
	return ItemTooltip.cache
end

local function currentSpec(data)
	local classSlug = Talents.playerClassSlug()
	if classSlug == nil or data == nil or data.classes == nil then
		return nil, nil
	end
	return Gear.specOf(data, classSlug, Talents.readRanks(data)), classSlug
end

--- Everything ItemTooltip.section reads from the client but the item: the data, the
--- spec and its weights, the class, level, faction, what is worn and the build. The
--- empty-slot hover reads this alone; the slot marks read it once for every slot.
function ItemTooltip.baseContext(data)
	local spec, classSlug = currentSpec(data)
	local current = ItemTooltip.current
	if current == nil and spec ~= nil then
		current = Gear.currentWeights(data, spec, ForeverSixtyInbox, Export.characterKey(), Compat.playerLevel())
	end
	return {
		data = data, spec = spec, current = current,
		classToken = classSlug ~= nil and classSlug:upper() or nil,
		level = Compat.playerLevel(), faction = Compat.playerFaction(),
		equipped = Gear.readEquipped(), build = Follow.build,
	}
end

--- `base` (ItemTooltip.baseContext's) with the item `itemLink` read from the client.
local function withItem(base, itemLink)
	local id, itemType, itemSubType, equipLocation, _, classId, subClassId = Compat.itemInfoInstant(itemLink)
	local ctx = {}
	for key, value in pairs(base) do
		ctx[key] = value
	end
	ctx.item = {
		itemId = tonumber(id) or itemIdOf(itemLink), link = itemLink, equipLocation = equipLocation,
		itemClass = classId or itemType, itemSubClass = subClassId or itemSubType,
		stats = Gear.statsOf(itemLink),
	}
	return ctx
end

--- `base` with one item read from the client: the quest reward and roll marks judge
--- each item through it, as the hover does (L4.2).
ItemTooltip.withItem = withItem

--- Everything ItemTooltip.section reads, from the client. A field so a spec
--- can hand the hook a context of its own (no client mock).
function ItemTooltip.contextFor(data, itemLink)
	return withItem(ItemTooltip.baseContext(data), itemLink)
end

--- The section for `itemLink` ({ lines, target, pending, head }), computed once per
--- link until the cache is reset; `base` (optional, ItemTooltip.baseContext's) saves
--- reading the client again for every item when one caller asks for several (the
--- slot marks). The hover and the slot marks read this one cache, so a slot's
--- target corner and its hover's Next line can never disagree.
function ItemTooltip.sectionFor(itemLink, base)
	local cached = ItemTooltip.cache[itemLink]
	if cached ~= nil then
		return cached
	end
	local data = ItemTooltip.data
	local entry = { lines = {}, pending = {} }
	if data ~= nil then
		local ctx = base ~= nil and withItem(base, itemLink) or ItemTooltip.contextFor(data, itemLink)
		entry.lines, entry.target, entry.pending, entry.head = ItemTooltip.section(ctx)
	end
	-- A section still waiting for a name is never kept: the next hover asks
	-- again, even when the name arrived while another item was hovered.
	if next(entry.pending) == nil then
		ItemTooltip.cache[itemLink] = entry
	end
	return entry
end

--- Ask the client once, at login, for every band pick of the character's
--- spec (one per paper-doll slot, 17 at most) it does not hold yet, so the
--- next upgrade's name and compare tooltip are there on the first hover.
function ItemTooltip.warmPicks(data)
	local spec = currentSpec(data)
	local faction = Compat.playerFaction()
	local specBis = spec ~= nil and type(data.bis) == "table" and data.bis[spec] or nil
	local band = faction ~= nil and Gear.bandFor(specBis, Compat.playerLevel()) or nil
	local asked = 0
	if band == nil then
		return asked
	end
	for _, entry in ipairs(Export.INVENTORY_SLOTS) do
		local pick = Gear.pick(spec, band, faction, entry.slot, data)
		if pick ~= nil and Compat.itemExists(pick.item_id) == false then
			-- A pick this client has never heard of: the tooltip goes to its alternatives.
			ItemTooltip.missing[pick.item_id] = true
		elseif pick ~= nil and not Compat.isItemCached(pick.item_id) then
			Compat.requestItemLoad(pick.item_id)
			asked = asked + 1
		end
	end
	return asked
end

--- The hover currently on screen, while it waits for an item the client
--- has to send (the compare tooltip's pick, a name not sent yet):
--- GET_ITEM_INFO_RECEIVED (ItemTooltip.onItemInfoReceived) redraws it.
ItemTooltip.activeHover = nil

--- A line's colour by Theme.HEX key. Muted lines on the client's tooltip take
--- Theme.HEX.tooltipMuted: never darker than the client's own grey.
local function opColor(key)
	if key == "muted" then
		key = "tooltipMuted"
	end
	return Theme.rgb(Theme.HEX[key or "body"])
end

--- The gap the client keeps between a double line's left and right text.
ItemTooltip.PAIR_GAP = 20

--- A hidden font string on the tooltip's own font, made on the first hover
--- that needs one (never at load) and reused.
local measurer

--- The widest of the client's own lines: each row's left text, plus its right
--- text and the pair gap for a double line, as `measure` reads them (visible
--- text only). nil with no rows or nothing to measure with. Pure.
function ItemTooltip.widestRow(rows, measure, gap)
	if type(measure) ~= "function" then
		return nil
	end
	local widest
	for _, row in ipairs(rows or {}) do
		local width = 0
		if type(row.left) == "string" and row.left ~= "" then
			width = measure(row.left) or 0
		end
		if type(row.right) == "string" and row.right ~= "" then
			width = width + (gap or 0) + (measure(row.right) or 0)
		end
		if widest == nil or width > widest then
			widest = width
		end
	end
	return widest
end

--- The client's own lines on `tooltip` as { left, right } rows, before this
--- addon adds any: read from its TextLeft/TextRight font strings.
local function clientRows(tooltip)
	local name = type(tooltip.GetName) == "function" and tooltip:GetName() or nil
	if name == nil or type(tooltip.NumLines) ~= "function" then
		return nil
	end
	local rows = {}
	for index = 1, tooltip:NumLines() do
		local left = _G[name .. "TextLeft" .. index]
		local right = _G[name .. "TextRight" .. index]
		rows[index] = {
			left = left ~= nil and left:GetText() or nil,
			right = right ~= nil and (type(right.IsShown) ~= "function" or right:IsShown()) and right:GetText() or nil,
		}
	end
	return rows
end

--- The raw width of a text on `tooltip`'s own font: a hidden font string made
--- on the first hover that needs one (never at load), set to the tooltip's
--- body font. A field, so a spec can measure without a client.
function ItemTooltip.rawMeasure(tooltip)
	if measurer == nil then
		local frame = CreateFrame("Frame", nil, UIParent)
		frame:Hide()
		measurer = frame:CreateFontString(nil, "ARTWORK", "GameTooltipText")
	end
	local sample = _G[tooltip:GetName() .. "TextLeft2"]
	if sample ~= nil and type(sample.GetFont) == "function" then
		local fontPath, size, flags = sample:GetFont()
		if fontPath ~= nil then
			measurer:SetFont(fontPath, size, flags)
		end
	end
	return function(text)
		measurer:SetText(text)
		if type(measurer.GetUnboundedStringWidth) == "function" then
			return measurer:GetUnboundedStringWidth()
		end
		return measurer:GetStringWidth()
	end
end

--- `measure` (visible text -> width, on the tooltip's own font) and the room
--- every line of this addon is fitted to: the client's own widest line,
--- measured before the section is added, so the section never widens the
--- tooltip. nil, nil where the client cannot say (the lines are drawn whole).
local function measureFor(tooltip)
	local ok, measure, room = pcall(function()
		local rows = clientRows(tooltip)
		if rows == nil or #rows == 0 then
			return nil, nil
		end
		local visible = ItemTooltip.visibleMeasure(ItemTooltip.rawMeasure(tooltip))
		return visible, ItemTooltip.widestRow(rows, visible, ItemTooltip.PAIR_GAP)
	end)
	if not ok then
		return nil, nil
	end
	return measure, room
end

--- One section line as the lines it is drawn on, fitted to `room`.
function ItemTooltip.fitLine(line, measure, room)
	if line.kind == "alternatives" then
		return ItemTooltip.fitAlternatives(line.entries, measure, room)
	elseif line.kind == "verdict" then
		return ItemTooltip.fitVerdict(line.verdict, line.unit, line.wornName, measure, room) or { line.text }
	elseif line.kind == "next" then
		return ItemTooltip.fitNext(line.upgrade, measure, room)
	elseif line.kind == "source" then
		return ItemTooltip.fitSource(line.pick, measure, room, line.boe)
	elseif line.kind == "planned" then
		return { ItemTooltip.fitPlanned(line, measure, room) }
	elseif line.kind == "bag" then
		return { ItemTooltip.fitBagLine(line.upgrade, measure, room) }
	elseif line.kind == "at" then
		return ItemTooltip.fitAt(line, measure, room)
	elseif line.kind == "breakdown" then
		return ItemTooltip.fitBreakdown(line.contributions, measure, room)
	end
	return { ItemTooltip.fitPlain(line.text, measure, room) }
end

--- Draws a section (`entry`: { lines, target, pending, head }) after the lines already
--- on `tooltip`: a blank row, the header ("Forever Sixty" gold, the spec and the tag;
--- "Equipped" muted at its right on the item worn, or at the next row's right when the
--- header needs the room: ItemTooltip.fitHeader), then each line in its colour, every
--- one fitted to `room` as `measure` reads it (nil: drawn whole). Then the compare
--- tooltip for the entry's target, and the hover `redraw` re-runs once an item it waits
--- for arrives (ItemTooltip.onItemInfoReceived).
local function drawEntry(tooltip, entry, measure, room, redraw)
	local head = entry.head or {}
	local header = ItemTooltip.fitHeader(head.specWord, head.tag, head.equipped, measure, room)
	local mr, mg, mb = opColor("muted")
	local doubles = type(tooltip.AddDoubleLine) == "function"
	if head.equipped and doubles then
		-- Both rows are double lines, so the client never carries a right text
		-- over onto the blank row: "Equipped" sits on the header's row, or on
		-- the next one when the header needs the room (fitHeader's case c).
		tooltip:AddDoubleLine(" ", " ")
	else
		tooltip:AddLine(" ")
	end
	if header.right ~= nil and doubles then
		tooltip:AddDoubleLine(header.left, header.right, 1, 1, 1, mr, mg, mb)
	else
		tooltip:AddLine(header.left)
	end
	-- "Equipped" the header cannot hold gets its own right-aligned row directly under it,
	-- never a content line's: every content line keeps the whole room (L4.1 fix round 6:
	-- "Next upgrade: Blackene…   Equipped" cut the pick's name to make room for it).
	if header.carry and doubles then
		tooltip:AddDoubleLine(" ", L.tooltipEquipped, 1, 1, 1, mr, mg, mb)
	end
	for _, line in ipairs(entry.lines) do
		local r, g, b = opColor(line.color)
		for _, text in ipairs(ItemTooltip.fitLine(line, measure, room)) do
			tooltip:AddLine(text, r, g, b)
		end
	end
	tooltip:Show()
	if entry.target ~= nil then
		Theme.showCompareTooltip(entry.target.itemId, entry.target.link, entry.target.label)
	end
	if entry.target == nil and next(entry.pending or {}) == nil then
		return
	end
	ItemTooltip.activeHover = {
		itemId = entry.target ~= nil and entry.target.itemId or nil,
		pending = entry.pending,
		redraw = redraw,
	}
end

--- Draws the item section after the client's own lines (drawEntry), every line
--- fitted to the client's own widest line (measured before anything is added).
--- Nothing at all with no lines and nothing worn.
--- The bag fact for a paper-doll slot that carries the arrow (ItemTooltip.bagSection), or
--- nil: no arrow, no injected reader, or a reader that failed (never a broken hover).
local function bagLead(slot, head)
	local find = ItemTooltip.bagUpgradeFor
	if slot == nil or type(find) ~= "function" then
		return nil
	end
	local ok, found = pcall(find, slot)
	if not ok or type(found) ~= "table" then
		return nil
	end
	return ItemTooltip.bagSection(found, head, ItemTooltip.unitOf(ItemTooltip.current))
end

--- The hover's own bag lead, for the character panel's targets (L4.3 fix round 1, item 8): a
--- slot with the arrow lists the very bag item its hover leads with.
ItemTooltip.bagLead = bagLead

local function addLines(tooltip, itemLink)
	-- A hover from before (another item) must not be redrawn onto this one.
	ItemTooltip.activeHover = nil
	local entry = ItemTooltip.sectionFor(itemLink)
	local head = entry.head or {}
	-- The worn item of a slot a bag item beats: the bag fact leads, alone (L4.1 fix round 5).
	if head.equipped then
		entry = bagLead(head.slot, head) or entry
	end
	if #entry.lines == 0 and not head.equipped then
		return
	end
	local measure, room = measureFor(tooltip)
	-- The room every line was fitted to: the client's own width, never ours.
	ItemTooltip.lastRoom = room
	drawEntry(tooltip, entry, measure, room, function()
		-- Re-running SetHyperlink re-fires the hook that called
		-- onTooltip, so the tooltip is rebuilt with the data that
		-- just arrived rather than gaining a second copy of the section.
		if type(tooltip.SetHyperlink) == "function"
			and type(tooltip.IsShown) == "function" and tooltip:IsShown() then
			tooltip:SetHyperlink(itemLink)
		end
	end)
end

--- One guarded body for both hook shapes. A Lua error inside a tooltip hook
--- breaks every tooltip in the game, so a failure turns the hook off.
function ItemTooltip.onTooltip(tooltip, itemLink)
	if ItemTooltip.disabled or itemLink == nil or not Prefs.flag("tooltip") then
		return
	end
	-- The compare tooltip is a GameTooltip too, so the client's processor
	-- calls this for it as well; adding lines there would open another
	-- compare tooltip for the same item, without end (seen in game as a
	-- C stack overflow).
	if tooltip == Theme.compareTooltipFrame then
		return
	end
	-- The hidden tooltip a worn item's enchant is read through (Compat): reading
	-- it must never draw a section, nor read the worn items again.
	if tooltip == Compat.scanFrame then
		return
	end
	local ok, err = pcall(addLines, tooltip, itemLink)
	if not ok then
		ItemTooltip.disabled = true
		Theme.note(string.format(L.diagTooltipHookFailed, tostring(err)))
	end
end

function ItemTooltip.itemLinkFrom(tooltip)
	if type(tooltip) ~= "table" or type(tooltip.GetItem) ~= "function" then
		return nil
	end
	local ok, _, link = pcall(tooltip.GetItem, tooltip)
	if not ok then
		return nil
	end
	return link
end

--- The client's paper-doll button for a site slot, or nil where this client has none.
local function slotButton(slot)
	local buttonName = ItemTooltip.SLOT_BUTTONS[slot]
	return buttonName ~= nil and Compat.clientGlobal(buttonName, "table") or nil
end

local function slotForOwner(owner)
	if owner == nil then
		return nil
	end
	for slot in pairs(ItemTooltip.SLOT_BUTTONS) do
		if slotButton(slot) == owner then
			return slot
		end
	end
	return nil
end

--- The client's own name for an empty slot ("Head"): its button's name less
--- "Character", upper-cased (HEADSLOT), as the client's own slot tooltip reads it.
local function slotTitle(slot)
	local buttonName = ItemTooltip.SLOT_BUTTONS[slot]
	return buttonName ~= nil and Compat.globalString(buttonName:gsub("^Character", ""):upper()) or nil
end

--- Whether the empty-slot section is already on a tooltip since its last
--- OnTooltipCleared (weak keys: a tooltip nothing else holds is collected).
local appendedSection = setmetatable({}, { __mode = "k" })

--- An EMPTY slot's section (ItemTooltip.emptySlotSection), appended onto `tooltip`
--- unless already there. Drawn whole: the client's own line on an empty slot is the
--- slot's name alone, far narrower than any section line.
function ItemTooltip.appendEmptySlotSection(tooltip, slot, button)
	if appendedSection[tooltip] then
		return
	end
	local data = ItemTooltip.data
	if data == nil then
		return
	end
	local entry = {}
	entry.lines, entry.target, entry.pending, entry.head =
		ItemTooltip.emptySlotSection(ItemTooltip.baseContext(data), slot)
	ItemTooltip.activeHover = nil
	-- An empty slot a bag item fills: the bag fact leads, alone (L4.1 fix round 5).
	entry = bagLead(slot, entry.head) or entry
	local function redraw()
		ItemTooltip.onSlotEnter(slot, button, true)
	end
	if #entry.lines == 0 then
		if next(entry.pending) ~= nil then
			-- The pick's name is on its way: draw the section when it comes.
			ItemTooltip.activeHover = { pending = entry.pending, redraw = redraw }
		end
		return
	end
	drawEntry(tooltip, entry, nil, nil, redraw)
	appendedSection[tooltip] = true
end

--- The character-frame slot buttons' own OnEnter: an EMPTY slot's section. A
--- slot holding an item is the item hook's. The client's own tooltip for the slot
--- (its name) stays when the client already drew it for this button; else, or
--- with `rebuild` (an item the hover waited for arrived), the tooltip is built
--- afresh: the slot's name, then the section.
function ItemTooltip.onSlotEnter(slot, button, rebuild)
	if ItemTooltip.slotDisabled or not Prefs.flag("tooltip") then
		return
	end
	if equippedLinkFor(slot) ~= nil then
		return
	end
	local ok, err = pcall(function()
		local owned = type(GameTooltip.GetOwner) == "function" and GameTooltip:GetOwner() == button
			and type(GameTooltip.IsShown) == "function" and GameTooltip:IsShown()
		if rebuild or not owned then
			GameTooltip:SetOwner(button, "ANCHOR_RIGHT")
			GameTooltip:ClearLines()
			appendedSection[GameTooltip] = nil
			local title = slotTitle(slot)
			if title ~= nil then
				GameTooltip:AddLine(title, 1, 1, 1)
			end
		end
		ItemTooltip.appendEmptySlotSection(GameTooltip, slot, button)
	end)
	if not ok then
		ItemTooltip.slotDisabled = true
		Theme.note(string.format(L.diagTooltipHookFailed, tostring(err)))
	end
end

--- Hooks every SLOT_BUTTONS global this client has. Idempotent.
function ItemTooltip.registerSlotButtons()
	if ItemTooltip.slotButtonsRegistered then
		return
	end
	ItemTooltip.slotButtonsRegistered = true
	for slot in pairs(ItemTooltip.SLOT_BUTTONS) do
		local button = slotButton(slot)
		if button ~= nil and type(button.HookScript) == "function" then
			button:HookScript("OnEnter", function(self)
				ItemTooltip.onSlotEnter(slot, self)
			end)
		end
	end
end

function ItemTooltip.onEmptySlotTooltipCleared(tooltip)
	appendedSection[tooltip] = nil
end

--- GameTooltip's own OnShow: Blizzard's PaperDollItemSlotButton_OnUpdate
--- rebuilds the hovered slot's tooltip without firing OnEnter, so an empty
--- slot's section is appended again here (a no-op when already there).
function ItemTooltip.onEmptySlotTooltipShow(tooltip)
	if ItemTooltip.slotDisabled or not Prefs.flag("tooltip") then
		return
	end
	local owner = type(tooltip.GetOwner) == "function" and tooltip:GetOwner() or nil
	local slot = slotForOwner(owner)
	if slot == nil or equippedLinkFor(slot) ~= nil then
		return
	end
	local ok, err = pcall(ItemTooltip.appendEmptySlotSection, tooltip, slot, owner)
	if not ok then
		ItemTooltip.slotDisabled = true
		Theme.note(string.format(L.diagTooltipHookFailed, tostring(err)))
	end
end

function ItemTooltip.registerEmptySlotRefresh()
	if ItemTooltip.emptySlotRefreshRegistered then
		return
	end
	ItemTooltip.emptySlotRefreshRegistered = true
	if type(GameTooltip) ~= "table" or type(GameTooltip.HookScript) ~= "function" then
		return
	end
	GameTooltip:HookScript("OnTooltipCleared", ItemTooltip.onEmptySlotTooltipCleared)
	GameTooltip:HookScript("OnShow", ItemTooltip.onEmptySlotTooltipShow)
end

--- Redraws the hover on screen once GET_ITEM_INFO_RECEIVED brings an id it
--- waits for (its compare item, or a name it showed as "…"); an id arriving
--- after the player moved on redraws nothing.
function ItemTooltip.onItemInfoReceived(itemId, success)
	local hover = ItemTooltip.activeHover
	if not success then
		-- The client cannot show this item at all: the section skips it for the
		-- pick's alternatives, and a hover still waiting on it redraws without it.
		ItemTooltip.missing[itemId] = true
		if hover ~= nil and hover.pending ~= nil and hover.pending[itemId] then
			ItemTooltip.resetCache()
			hover.redraw()
		end
		return
	end
	if hover == nil then
		return
	end
	if hover.itemId ~= itemId and (hover.pending == nil or not hover.pending[itemId]) then
		return
	end
	ItemTooltip.resetCache()
	hover.redraw()
end

--- The one GET_ITEM_INFO_RECEIVED listener the tooltip needs. Idempotent.
function ItemTooltip.registerItemInfoRefresh()
	if ItemTooltip.itemInfoFrame ~= nil then
		return ItemTooltip.itemInfoFrame
	end
	local frame = CreateFrame("Frame", nil, UIParent)
	Theme.registerEvent(frame, "GET_ITEM_INFO_RECEIVED")
	frame:SetScript("OnEvent", function(_, _, itemId, success)
		ItemTooltip.onItemInfoReceived(itemId, success)
	end)
	ItemTooltip.itemInfoFrame = frame
	return frame
end

function ItemTooltip.hasProcessor()
	return type(TooltipDataProcessor) == "table"
		and type(TooltipDataProcessor.AddTooltipPostCall) == "function"
		and type(Enum) == "table"
		and type(Enum.TooltipDataType) == "table"
		and Enum.TooltipDataType.Item ~= nil
end

--- Hooks exactly one of the two tooltip shapes, never both: the modern
--- processor when the client has it, else the legacy script. Idempotent.
function ItemTooltip.register()
	if ItemTooltip.registered then
		return ItemTooltip.how
	end
	ItemTooltip.registered = true
	ItemTooltip.registerSlotButtons()
	ItemTooltip.registerEmptySlotRefresh()
	ItemTooltip.registerItemInfoRefresh()
	if ItemTooltip.hasProcessor() then
		TooltipDataProcessor.AddTooltipPostCall(Enum.TooltipDataType.Item, function(tooltip)
			ItemTooltip.onTooltip(tooltip, ItemTooltip.itemLinkFrom(tooltip))
		end)
		ItemTooltip.how = "processor"
		Theme.noteVerbose(string.format(L.diagItemTooltipHook, ItemTooltip.how))
		return ItemTooltip.how
	end
	if type(GameTooltip) == "table" and type(GameTooltip.HookScript) == "function" then
		GameTooltip:HookScript("OnTooltipSetItem", function(tooltip)
			ItemTooltip.onTooltip(tooltip, ItemTooltip.itemLinkFrom(tooltip))
		end)
		ItemTooltip.how = "legacy"
		Theme.noteVerbose(string.format(L.diagItemTooltipHook, ItemTooltip.how))
		return ItemTooltip.how
	end
	Theme.note(L.diagNoTooltipHook)
	ItemTooltip.how = nil
	return nil
end

ns.ItemTooltip = ItemTooltip
return ItemTooltip
