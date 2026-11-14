-- addon/ForeverSixty/surfaces/ItemTooltip.lua
-- The item tooltip (design/specs/2026-11-08-addon-ux.md section 2). After the client's own
-- lines, with no blank row and no header, one answer and at most one supporting line, in the
-- client's own colours: an upgrade in green behind the green arrow ("Upgrade +2.2 DPS"), the
-- band's pick in gold ("Your best in slot until level 30"), information in white ("Next
-- upgrade: Mutant Scale Breastplate +6.1 DPS"), no or detail in grey ("Sidegrade", "Not for
-- you until level 40"). Everything else (the breakdown, alternatives, the build's gear list,
-- caps, the enchant note and where the weights came from) is drawn only while Shift is held,
-- and holding or letting go of Shift redraws the tooltip (ItemTooltip.onModifierChanged).
--
-- Pure model, thin hook: ItemTooltip.section builds the lines from plain tables (every
-- number from Progress.itemVerdict and Progress.roundedGain, the weights from
-- Gear.currentWeights, the wear rules from Rules.canEquip); each line carries its `role`
-- ("verdict", "support" or "shift"). The guarded hooks after it read the client and draw. A
-- section with nothing to say is not drawn at all (G5).
--
-- An EMPTY character slot shows the same grammar with the band's pick for the slot
-- (ItemTooltip.emptySlotSection). A player's tooltip is surfaces/PlayerTooltip.lua's.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Talents = ns.Talents or require("Talents")
local Gear = ns.Gear or require("Gear")
local Sources = ns.Sources or require("Sources")
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

local function colored(name, hex)
	if hex == nil then
		return name
	end
	return "|cff" .. hex .. name .. "|r"
end

local function green(text)
	return Theme.inlineColor("tooltipGreen", text)
end

--- A signed figure: "+2.2 DPS", "−1.3 HPS", "+3 score". Pure.
function ItemTooltip.figure(gain, unit)
	local sign = gain < 0 and L.tooltipMinus or "+"
	local size = math.abs(gain)
	if unit == "score" then
		return string.format(L.tooltipFigureScore, sign, math.floor(size + 0.5))
	end
	return string.format(L.tooltipFigure, sign, size, unitWord(unit))
end

--- The verdict line for one Progress.itemVerdict answer (spec 2.2 cases 6, 7, 9, 10, 12, 13),
--- as { text, color, upgrade?, gain?, sideFigure?, fit? }, or nil with nothing to say. The
--- verdict was asked with the HOVERED item as `plan` and the worn one as `worn`, so the hovered
--- item's gain is minus its delta, rounded as shown (Progress.roundedGain) against the one floor
--- (Progress.MIN_GAIN, 1.0 in the unit). `unit` is "dps", "hps" or "score"; `wornName` the worn
--- item's name (nil while the client has not sent it). `opts` (optional):
---
--- - `over`: name the worn item ("+2.2 DPS over <worn>": a ring or a trinket, where the slot
---   is ambiguous), in its quality colour `wornHex`;
--- - `requiredLevel`, `level`: an item above the character's level reads "+3.5 DPS at level 27".
---
--- An upgrade (or any gain into an empty slot) is the green arrow and its figure, in green; a
--- loss from the floor is its figure, "−1.3 DPS", and anything between "Sidegrade" (its size kept
--- as `sideFigure` for Shift), both grey. `match` says nothing (the client's compare tooltip
--- says Equipped); `blocked` reads "Not with your two-hander".
function ItemTooltip.verdictLine(verdict, unit, wornName, opts)
	if verdict == nil or verdict.state == "match" then
		return nil
	end
	if verdict.state == "blocked" then
		return { text = L.tooltipTwoHander, color = "muted" }
	end
	if verdict.delta == nil or verdict.state == "noplan" then
		return nil
	end
	unit = unit or "score"
	opts = opts or {}
	local gain = Progress.roundedGain(-verdict.delta, unit == "score" and "score" or "dps")
	if gain >= Progress.MIN_GAIN or (verdict.state == "empty" and gain > 0) then
		local shown = ItemTooltip.figure(gain, unit)
		local line = { color = "tooltipGreen", upgrade = true, gain = gain }
		if type(opts.requiredLevel) == "number" and type(opts.level) == "number" and opts.requiredLevel > opts.level then
			line.text = Theme.arrowMarkup() .. string.format(L.tooltipUpgradeAt, shown, opts.requiredLevel)
		elseif opts.over and wornName ~= nil and verdict.state ~= "empty" then
			local function build(name)
				return Theme.arrowMarkup() .. string.format(L.tooltipUpgradeOver, shown, colored(name, opts.wornHex))
			end
			line.text, line.fit = build(wornName), { name = wornName, build = build }
		else
			line.text = Theme.arrowMarkup() .. string.format(L.tooltipUpgrade, shown)
		end
		return line
	end
	if verdict.state == "empty" then
		return nil
	end
	if gain <= -Progress.MIN_GAIN then
		return { text = string.format(L.tooltipDown, ItemTooltip.figure(gain, unit)), color = "muted", gain = gain }
	end
	local line = { text = L.tooltipSide, color = "muted", side = true, gain = gain }
	if unit ~= "score" and gain ~= 0 then
		line.sideFigure = ItemTooltip.figure(gain, unit)
	end
	return line
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

--- The same, for the character-panel tab's target names (L3.7 item 22).
ItemTooltip.shorten = shorten


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

--- A line within `room`: whole when it fits; a line carrying `fit` ({ name, build }: an item
--- name inside it) shortens that name, a character at a time, before anything else is cut (the
--- figure and the label never are); any other line is cut by measured width. One line.
function ItemTooltip.fitNamed(line, measure, room)
	local fits = fitter(measure, room)
	if fits(line.text) then
		return line.text
	end
	if type(line.fit) == "table" then
		return shorten(line.fit.build, line.fit.name, fits, 1)
	end
	return ItemTooltip.fitPlain(line.text, measure, room)
end

local function capitalised(text)
	return (tostring(text):gsub("^%l", string.upper))
end

--- The class plural for a class token ("shamans"), or the token itself.
local function classPlural(classToken)
	local class = type(classToken) == "string" and L.tooltipClassPlural[classToken:upper()] or nil
	return class or tostring(classToken)
end

--- The grey verdict Rules.canEquip's refusal reads as (spec 2.2 cases 4 and 5): "Mail armor at
--- level 40" (`subclass` canEquip's own plain-words reason), "Not for shamans". nil for "ok".
function ItemTooltip.untilLine(classToken, answer, level, subclass)
	if answer == "until" then
		return string.format(L.tooltipNotUntil, capitalised(subclass), level)
	end
	if answer == "never" then
		return string.format(L.tooltipNotFor, classPlural(classToken))
	end
	return nil
end

--- The refusal's reason, for Shift: "Rogues cannot use two-handed swords", "Shamans cannot
--- dual wield". `subclass` is canEquip's own plain-words reason. nil for "ok", and for "until"
--- (its verdict line already names the reason).
function ItemTooltip.refusalReason(classToken, answer, subclass)
	if answer == "never" then
		local class = capitalised(classPlural(classToken))
		if subclass == "dual wielding" then
			return string.format(L.tooltipReasonTwoWeapons, class)
		end
		return string.format(L.tooltipReasonNever, class, subclass)
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
--- chance last when the data has one. nil with no source. (`boe` is accepted and unused:
--- "check the AH" was cut, spec 1.1.)
function ItemTooltip.sourceLine(pick, boe)
	local first, second, percent = ItemTooltip.sourceParts(pick, boe)
	return ItemTooltip.sourceText(first, second, percent)
end

--- A pick's source as its parts: the place (or the source), the boss (or the
--- kind's name, nil when the source already says it) and the drop chance as a
--- whole percent (nil for no drop). Nothing with no source, but a crafted pick.
function ItemTooltip.sourceParts(pick, _boe)
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
		-- A reputation or vendor item that asks a standing names it instead of its kind
		-- ("Silverwing Sentinels · Friendly"): the standing is the gate (source filter spec 7).
		local standing = Sources.standingWord(pick.standing)
		if standing ~= nil then
			second = standing
		elseif kindName ~= nil and kindName:lower() ~= first:lower() then
			second = kindName
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

--- An alternative's gain as printed after its name, muted, with its unit: "+1.1 DPS" ("+12
--- score"), the minus sign for a loss. nil with no gain.
function ItemTooltip.alternativeGain(gain, unit)
	if gain == nil then
		return nil
	end
	return ItemTooltip.figure(gain, unit or "score")
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

--- Ranged weapon subclass ids -> the type's word in the game's own terms (the owner, 2026-11-08:
--- name the type, never explain what it implies, so "(gun)", never "(needs bullets)").
ItemTooltip.RANGED_TYPE_KEYS = {
	[2] = "tooltipTypeBow", [3] = "tooltipTypeGun", [16] = "tooltipTypeThrown", [18] = "tooltipTypeCrossbow",
	[19] = "tooltipTypeWand",
}

--- The weapon subclass id of thrown weapons (the client's).
ItemTooltip.THROWN_SUBCLASS = 16

--- "(gun)", "(bow)", "(crossbow)", "(thrown)" or "(wand)" for a ranged item whose type differs from
--- the worn ranged weapon's, from the two items' subclass ids (`wornSub`, `itemSub`); nil when they
--- match or the item is no ranged type this table knows. Nothing worn: the type is named. Pure.
function ItemTooltip.typeTag(wornSub, itemSub)
	local key = ItemTooltip.RANGED_TYPE_KEYS[tonumber(itemSub)]
	if key == nil or (wornSub ~= nil and tonumber(wornSub) == tonumber(itemSub)) then
		return nil
	end
	return L[key]
end

--- The alternatives as printed lines (spec 2.3 item 2), for `entries` ({ name, hex, gain,
--- unit, flag } each; `false` for a name not sent yet, which is skipped, never "…"): one per
--- line, at most two, "Also good: <item> +2.2 (needs bullets)", the second indented with
--- spaces to sit under the first item's name (as `measure` reads the label; four spaces with
--- nothing to measure with). A row too wide shortens its name first. Empty with none.
function ItemTooltip.fitAlternatives(entries, measure, room)
	local shown = {}
	for _, entry in ipairs(entries or {}) do
		if type(entry) == "table" and #shown < 2 then
			shown[#shown + 1] = entry
		end
	end
	if #shown == 0 then
		return {}
	end
	local indent = "    "
	if type(measure) == "function" then
		local label, space = measure(L.tooltipAlternativesLabel .. " "), measure(" ")
		if type(label) == "number" and type(space) == "number" and space > 0 then
			indent = string.rep(" ", math.floor(label / space + 0.5))
		end
	end
	local fits = fitter(measure, room)
	local rows = {}
	for index, entry in ipairs(shown) do
		local function build(name)
			local text = entryText(entry, name)
			return index == 1 and string.format(L.tooltipAlternativesOne, text) or indent .. text
		end
		rows[index] = shorten(build, entry.name, fits, 1)
	end
	return rows
end
--- The last Shift line (spec 2.3 item 6), dimmer: "Forever Sixty · Beast Mastery estimate",
--- "Forever Sixty · Beast Mastery, your sim" (or "your build"); "Forever Sixty" alone with no
--- spec. Pure.
function ItemTooltip.attributionText(specWord, tag)
	if specWord == nil then
		return L.addonName
	end
	local format = tag == "sim" and L.tooltipAttrSim or tag == "build" and L.tooltipAttrBuild or L.tooltipAttrEstimate
	return string.format(L.tooltipAttribution, string.format(format, specWord))
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

--- The Shift line naming the published pick a source choice set aside (source filter spec 6):
--- "All sources: <name> +0.3 DPS", the gap left out when there is none to price.
--- `name` is already in its quality colour. Pure.
function ItemTooltip.setAsideText(name, gap, unit)
	if type(gap) ~= "number" or gap <= 0 then
		return string.format(L.tipSetAsideBare, name)
	end
	local figure = unit == "score" and string.format(L.tipSetAsideScore, gap)
		or string.format(L.tipSetAsideDps, gap, unitWord(unit))
	return string.format(L.tipSetAside, name, figure)
end

--- The empty slot's line when the sources on leave it without an option: "No finger from these
--- sources at this level". Pure.
function ItemTooltip.noSlotFromSourcesText(slot)
	return string.format(L.tipNoSlotFromSources, string.lower(L.gearSlotLabels[slot] or slot))
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

--- A stat in the player's words: "weapon DPS" for a weapon's own DPS, else the
--- client's own name for it ("Agility") (its ITEM_MOD_*_SHORT global: the generic one first
--- where a stat has several, STAT_WORD_KEYS; else the alphabetical first), else
--- Locale's. A caster never reads "Melee Hit Rating" for hit.
function ItemTooltip.statWord(stat)
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

--- Display order: the five primary stats first, then Armor, then the rest, a weapon's DPS last.
local BREAKDOWN_ORDER = {
	strength = 1, agility = 2, stamina = 3, intellect = 4, spirit = 5, armor = 6, weapon_dps = 99,
}

--- A stat's word in the breakdown: Armor by Locale (the client's armor global is a sentence
--- in some locales), the rest by ItemTooltip.statWord.
local function breakdownWord(stat)
	if stat == "armor" then
		return L.tooltipArmor
	end
	return ItemTooltip.statWord(stat)
end

--- The breakdown's parts (spec 2.3 item 1): every stat that moves, signed, in the player's
--- words, in a fixed order: "+7 Agi", "+2 Sta", "+6 Armor", "+3.4 weapon DPS". Pure.
function ItemTooltip.breakdownParts(contributions)
	local chosen = {}
	for index, entry in ipairs(contributions or {}) do
		chosen[index] = entry
	end
	table.sort(chosen, function(left, right)
		local leftRank, rightRank = BREAKDOWN_ORDER[left.stat] or 50, BREAKDOWN_ORDER[right.stat] or 50
		if leftRank ~= rightRank then
			return leftRank < rightRank
		end
		if math.abs(left.weighted) ~= math.abs(right.weighted) then
			return math.abs(left.weighted) > math.abs(right.weighted)
		end
		return left.stat < right.stat
	end)
	local parts = {}
	for _, entry in ipairs(chosen) do
		-- The weapon's own DPS is left out (the owner, 2026-11-08): beside "+3.5 DPS" on the headline,
		-- "+3.4 weapon DPS" read as the same number twice. The breakdown names the stats that change.
		if entry.stat ~= "weapon_dps" then
			local size = math.abs(entry.delta)
			local amount
			if math.abs(size - math.floor(size + 0.5)) < 1e-9 then
				amount = string.format("%d", math.floor(size + 0.5))
			else
				amount = string.format("%.1f", size)
			end
			local sign = entry.delta < 0 and L.tooltipMinus or "+"
			parts[#parts + 1] = string.format(L.tooltipBreakdownPart, sign .. amount, breakdownWord(entry.stat))
		end
	end
	return parts
end

--- The breakdown as one line, "+7 Agi, +2 Sta, +6 Armor"; nil with nothing that moves.
function ItemTooltip.breakdownText(contributions)
	local parts = ItemTooltip.breakdownParts(contributions)
	return #parts > 0 and table.concat(parts, L.tooltipBreakdownJoin) or nil
end

--- The breakdown within `room` (led by `of`, the item it describes, when given): as many
--- parts on a line as fit, the rest on the next line
--- (nothing is dropped; the tooltip never widens). A part wider than the room alone is cut.
function ItemTooltip.fitBreakdown(contributions, measure, room, of)
	local fits = fitter(measure, room)
	local rows, current = {}, nil
	local parts = ItemTooltip.breakdownParts(contributions)
	if of ~= nil and parts[1] ~= nil then
		parts[1] = string.format(L.tooltipBreakdownOf, of, parts[1])
	end
	for _, part in ipairs(parts) do
		local joined = current ~= nil and (current .. L.tooltipBreakdownJoin .. part) or part
		if current == nil or fits(joined) then
			current = joined
		else
			rows[#rows + 1] = current
			current = part
		end
	end
	if current ~= nil then
		rows[#rows + 1] = current
	end
	for index, row in ipairs(rows) do
		rows[index] = ItemTooltip.fitPlain(row, measure, room)
	end
	return rows
end

--- The breakdown line for a verdict that carries a figure, or nil. `of` (optional): the
--- coloured name of the item it describes when that is not the hovered one ("Ranger Bow:
--- +1 Agility").
local function breakdownLine(verdict, of)
	local contributions = verdict ~= nil and verdict.contributions or nil
	if contributions == nil or #contributions == 0 then
		return nil
	end
	local text = ItemTooltip.breakdownText(contributions)
	return { kind = "breakdown", role = "shift", contributions = contributions, color = "muted", of = of,
		text = of ~= nil and string.format(L.tooltipBreakdownOf, of, text) or text }
end

--- "Next upgrade: <name> +6.1 DPS" (spec 2.2 case 16): the label in the line's white, the name
--- as given (in its quality colour), the gain green, already rounded and at the floor or
--- above. With no gain (an empty slot: nothing worn to gain over; a weapon pick whose stats
--- the client has not sent, case 17) the label and the name. `tag` (a thrown weapon's
--- "(thrown)") follows the name, muted.
function ItemTooltip.nextLine(name, gain, unit, tag)
	if tag ~= nil then
		name = name .. " " .. muted(tag)
	end
	if gain == nil then
		return L.tooltipNextLabel .. " " .. name
	end
	return L.tooltipNextLabel .. " " .. name .. " " .. green(ItemTooltip.figure(gain, unit))
end

--- The Next line within `room`: one line, the name shortened before the gain is ever cut.
--- `upgrade` is { name, hex, gain, unit, tag }.
function ItemTooltip.fitNext(upgrade, measure, room)
	return { shorten(function(name)
		return ItemTooltip.nextLine(colored(name, upgrade.hex), upgrade.gain, upgrade.unit, upgrade.tag)
	end, upgrade.name, fitter(measure, room), 1) }
end

--- A band in the player's words: its levels up to the band before `nextBand` ("40 to 49"; a band
--- the nightly skipped is spanned), or "level 60" at the top (no `nextBand`). Pure.
function ItemTooltip.bandWord(band, nextBand)
	if nextBand == nil then
		return string.format(L.tooltipBandCap, band)
	end
	return string.format(L.tooltipBandSpan, band, nextBand - 1)
end

--- The gold line on the band's own pick (spec 2.2 cases 8 and 18): "Best for 40 to 49", "Best at
--- 60" in the top band. Pure.
function ItemTooltip.bestText(band, nextBand)
	if nextBand == nil then
		return string.format(L.tooltipBestAtCap, band)
	end
	return string.format(L.tooltipBestFor, ItemTooltip.bandWord(band, nextBand))
end

--- The white line on a worn item the band's pick does not beat (case 19), the pick named: "Better
--- than the 40 to 49 pick: <pick>" when the pick loses to it, "As good as the 40 to 49 pick:
--- <pick>" when the pick gains nothing or under the floor, or is unscored. `name` may carry its
--- colour. Pure.
function ItemTooltip.holdText(band, nextBand, name, gain)
	local format = type(gain) == "number" and gain < 0 and L.tooltipBeatsPick or L.tooltipMatchesPick
	return string.format(format, ItemTooltip.bandWord(band, nextBand), name)
end

--- The first band after `band` whose pick for `slot` is not `itemId`: the band the Best and At
--- lines turn on (a pick that holds for two bands running is "Best for 30 to 49", and the At line
--- names the band that changes it), or nil when no later band does.
function ItemTooltip.bandAfter(data, spec, faction, slot, band, itemId)
	local specBis = spec ~= nil and type(data) == "table" and type(data.bis) == "table" and data.bis[spec] or nil
	local candidate = ItemTooltip.nextBand(specBis, band)
	while candidate ~= nil do
		local pick = Gear.pick(spec, candidate, faction, slot, data)
		if pick == nil or pick.item_id ~= itemId then
			return candidate
		end
		candidate = ItemTooltip.nextBand(specBis, candidate)
	end
	return nil
end

--- "At level 30: <name>" (spec 2.2 cases 18 to 20), the name in its quality colour: the next
--- rung's pick under a worn item. Its source and gain wait under Shift.
function ItemTooltip.atText(at, name)
	return string.format(L.tooltipAtBandName, at.band, colored(name or at.name, at.hex))
end

--- The At line within `room`: the name shortened only when the line would widen the tooltip.
function ItemTooltip.fitAt(at, measure, room)
	return { shorten(function(name) return ItemTooltip.atText(at, name) end, at.name, fitter(measure, room), 1) }
end

--- The arrow slot's line (case 15): the green arrow, "In your bags: <item> +0.8 DPS", the name
--- in its quality colour and the gain green; the name shortened before the gain is ever cut.
--- `upgrade` is { name, hex, gain, unit }.
function ItemTooltip.bagItemText(upgrade, name)
	return Theme.arrowMarkup() .. string.format(L.tooltipInBags, colored(name or upgrade.name, upgrade.hex),
		green(ItemTooltip.figure(upgrade.gain, upgrade.unit)))
end

function ItemTooltip.fitBagLine(upgrade, measure, room)
	local fits = fitter(measure, room)
	return shorten(function(name) return ItemTooltip.bagItemText(upgrade, name) end, upgrade.name, fits, 1)
end

--- The section of a paper-doll slot that carries the arrow (case 15): the bag fact alone on
--- one line, "In your bags: <item> +0.8 DPS" (the gain the chip's list gives,
--- Progress.bagUpgrades, as shown), then under Shift where the weights came from. The compare
--- tooltip shows the bag item, titled "Next upgrade". `found` is the chip's list entry ({ link,
--- itemId, delta, unit }); `head` the hover's own header ({ specWord, tag, equipped, slot }).
function ItemTooltip.bagSection(found, head, unit, qualityHex)
	local name = type(found.link) == "string" and found.link:match("%[(.-)%]") or Compat.itemName(found.itemId)
	if name == nil then
		return nil
	end
	local shownUnit = found.unit == "dps" and (unit == "hps" and "hps" or "dps") or "score"
	local upgrade = { name = name, hex = (qualityHex or ItemTooltip.qualityHexOf)(found.itemId), unit = shownUnit,
		gain = Progress.roundedGain(found.delta, found.unit == "score" and "score" or "dps") }
	head = head or {}
	return {
		lines = {
			{ kind = "bag", role = "verdict", upgrade = upgrade, color = "tooltipBody",
				text = ItemTooltip.bagItemText(upgrade) },
			{ kind = "attribution", role = "shift", color = "tooltipDim",
				text = ItemTooltip.attributionText(head.specWord, head.tag) },
		},
		target = { itemId = found.itemId, link = found.link, label = L.tooltipCompareLabel },
		pending = {},
		head = head,
	}
end

--- Injection point (surfaces/SlotMarks.lua sets it): the chip's list entry that puts an
--- arrow on a paper-doll slot, or nil. The slot hovers read it to lead with the bag fact.
ItemTooltip.bagUpgradeFor = nil

--- The slots whose verdict names the worn item it replaces: two of each are worn, so "+2.2
--- DPS" alone would not say which.
local OVER_SLOTS = { finger1 = true, finger2 = true, trinket1 = true, trinket2 = true }

--- The client reads `ctx` falls back to when a caller hands none in (Gear.itemStatsOf: the
--- gear gap's own read).
local function itemStatsOf(itemId)
	return Gear.itemStatsOf(itemId)
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
--- its own alternatives). nil when it can show none. Progress.reachablePick, the
--- gear gap's own rule (L5.1 fix round 3).
local reachablePick = Progress.reachablePick

--- The part of the section that scores one item: the one verdict the item tooltip, the quest
--- reward labels and the roll line all read, so they can never disagree. `ctx` is
--- ItemTooltip.section's; `pending` (optional) gathers the worn item whose name the client has
--- not sent. Answers { gear, slots, wornSlot, wearable, usable, refusal, reason, blocked,
--- heldByTwoHander, slot, verdict, primary, unit, pairLine }:
---
--- - `gear`: the item goes in a slot the plan covers (none: nothing else is set);
--- - `wearable`: the slots Rules.canEquip allows now (none: `refusal`, "Not for you until
---   level 40", is the primary line, grey, its `reason` "Mail armor at 40" for Shift);
--- - `heldByTwoHander`: every allowed slot is a hand the worn two-hander holds ("Not with your
---   two-hander");
--- - `slot` the slot it is judged in, `verdict` Progress.itemVerdict's answer there (on an item
---   not worn, with weights), and `primary` the verdict line (ItemTooltip.verdictLine: the 1.0
---   floor, the empty-slot fill, "over <worn>" where two slots take it, the required level);
--- - `pairLine`: a one-hander under a worn two-hander whose best partner in the bags makes the
---   pair an upgrade, "Upgrade +14.0 DPS as a pair with <item>" (case 7), which then leads.
function ItemTooltip.judge(ctx, pending)
	pending = pending or {}
	local item = ctx.item or {}
	local slots = Gear.slotsFor(item.equipLocation)
	if #slots == 0 then
		return { gear = false, slots = slots, wearable = {}, usable = {} }
	end
	local equipped = ctx.equipped or {}
	local nameOf = ctx.nameOf or Compat.itemName
	local qualityHex = ctx.qualityHex or qualityHexOf
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
	local wearable, refusal, reason = {}, nil, nil
	for _, slot in ipairs(slots) do
		local answer, unlock, why = Rules.canEquip(ctx.classToken, ctx.level, item.itemClass,
			item.itemSubClass, item.equipLocation, slot)
		if answer == "ok" then
			wearable[#wearable + 1] = slot
		elseif refusal == nil then
			refusal = ItemTooltip.untilLine(ctx.classToken, answer, unlock, why)
			reason = ItemTooltip.refusalReason(ctx.classToken, answer, why)
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
		primary = refusal ~= nil and { kind = "refusal", text = refusal, color = "muted", reason = reason } or nil
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
		-- "over <worn>" only for rings and trinkets, where two slots take the item.
		primary = ItemTooltip.verdictLine(verdict, unit, wornName, {
			over = #usable > 1 and OVER_SLOTS[slot] == true,
			wornHex = worn ~= nil and qualityHex(worn.itemId) or nil,
			requiredLevel = item.requiredLevel, level = ctx.level,
		})
		if primary ~= nil then
			primary.kind, primary.verdict, primary.unit, primary.wornName = "verdict", verdict, unit, wornName
			primary.breakdown = breakdownLine(verdict)
		end
	end
	-- A one-hander under a worn two-hander, by a class that can dual wield: the swap also needs a
	-- second weapon, so the verdict is the pair's with the best one carried (case 7).
	local pairLine
	local wornMain = equipped.main_hand
	if primary ~= nil and primary.kind == "verdict" and slot == "main_hand" and ctx.bags ~= nil
		and wornMain ~= nil and wornMain.equipLoc == "INVTYPE_2HWEAPON"
		and (item.equipLocation == "INVTYPE_WEAPON" or item.equipLocation == "INVTYPE_WEAPONMAINHAND")
		and Rules.canEquip(ctx.classToken, ctx.level, item.itemClass, item.itemSubClass, "INVTYPE_WEAPON",
			"off_hand") == "ok" then
		-- The bags are read only here: a function reads them now, a table is the list itself.
		local bags = type(ctx.bags) == "function" and ctx.bags() or ctx.bags
		local partner = Progress.pairPartner(hovered, wornMain, bags, weights, perPoint, weapon, function(candidate)
			return Rules.canEquip(ctx.classToken, ctx.level, candidate.itemClass, candidate.itemSubClass,
				candidate.equipLocation, "off_hand") == "ok"
		end)
		if partner ~= nil then
			local partnerId = partner.item.itemId
			local partnerName = nameOf(partnerId)
			local pairVerdict = ItemTooltip.verdictLine(partner, unit, nil)
			if partnerName == nil then
				pending[partnerId] = true
			elseif pairVerdict ~= nil and pairVerdict.upgrade then
				local shown = ItemTooltip.figure(pairVerdict.gain, unit)
				local hex = qualityHex(partnerId)
				local function build(name)
					return Theme.arrowMarkup() .. string.format(L.tooltipUpgradePair, shown, colored(name, hex))
				end
				pairLine = { kind = "pair", role = "verdict", color = "tooltipGreen", upgrade = true,
					gain = pairVerdict.gain, partner = partnerId, text = build(partnerName),
					fit = { name = partnerName, build = build }, breakdown = breakdownLine(partner) }
			end
		end
	end
	return {
		gear = true, slots = slots, wornSlot = wornSlot, wearable = wearable, usable = usable,
		refusal = refusal, reason = reason, blocked = blocked, heldByTwoHander = heldByTwoHander, slot = slot,
		verdict = verdict, primary = primary, unit = unit, pairLine = pairLine,
	}
end

--- A line with its role: "verdict" (the answer), "support" (at most one under it) or "shift"
--- (drawn only while Shift is held).
local function as(role, line)
	if line ~= nil then
		line.role = role
	end
	return line
end

--- The section for one hovered item (spec 2.2 and 2.3): the lines ({ text, color, role, kind?
--- }) in draw order (the verdict, at most one support line, then the Shift block in the spec's
--- order: the figure's breakdown and the reasons, alternatives, the build's gear list, caps,
--- the enchant note, the attribution), the compare tooltip's item ({ itemId, link, label } for
--- the next upgrade, only once the client holds it), the ids whose names the client has not
--- sent yet (a set: the hover redraws when one arrives), and the head ({ specWord, tag,
--- equipped, slot }). `ctx` holds what the hook read:
---
--- - `data`, `spec`, `classToken`, `level`, `faction`;
--- - `item`: { itemId, link, equipLocation, itemClass, itemSubClass, stats, requiredLevel };
--- - `current`: Gear.currentWeights' answer (nil: no weights for the spec);
--- - `equipped`: Gear.readEquipped()'s table; `build`: the loaded build;
--- - `nameOf(id)`, `statsOf(id)`, `isCached(id)`, `qualityHex(id)`, `isMissing(id)`: the
---   client's name, stats, cache state, quality colour and "cannot show it at all".
---
--- A pick (or an alternative, or the next rung's pick) the client has not sent yet is named
--- from the data (Data.bisNames, never "…") and, for a weapon, waited on until the client's
--- stats arrive; its id stays pending so the hover redraws when they do.
function ItemTooltip.section(ctx)
	local pending = {}
	local item = ctx.item or {}
	local slots = Gear.slotsFor(item.equipLocation)
	if #slots == 0 then
		return {}, nil, pending, nil
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
	local unit = ItemTooltip.unitOf(current)
	local function pickName(itemId)
		return pickNameOf(data, nameOf, pending, itemId)
	end
	--- A weapon priced from the data's own DPS waits for the client's stats (pending, as above).
	local function standIn(itemId)
		pending[itemId] = true
	end
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
	--- The ranged type's word ("(gun)", "(thrown)") for an item named in the ranged slot whose
	--- type differs from the worn ranged weapon's (ItemTooltip.typeTag), else nil.
	local function typeTagOf(itemId)
		if slot ~= "ranged" then
			return nil
		end
		return ItemTooltip.typeTag(wornSub, subClassOf(itemId))
	end
	local specWord = specName(data, ctx.spec)
	local head = { specWord = specWord, tag = current ~= nil and current.tag or nil,
		equipped = wornSlot ~= nil, slot = wornSlot }
	local verdict = judged.pairLine or judged.primary
	-- A hand the worn two-hander holds says only that (case 6).
	if heldByTwoHander then
		return { as("verdict", verdict) }, nil, pending, head
	end

	-- Shift's first group: what the verdict's figure is made of, and why not.
	local reasons = {}
	if verdict ~= nil and verdict.reason ~= nil then
		reasons[#reasons + 1] = { kind = "reason", text = verdict.reason, color = "muted" }
	end
	if verdict ~= nil and verdict.sideFigure ~= nil then
		reasons[#reasons + 1] = { kind = "figure", text = verdict.sideFigure, color = "muted" }
	end
	if verdict ~= nil and verdict.breakdown ~= nil then
		reasons[#reasons + 1] = verdict.breakdown
	end
	-- A two-hander's figure is net of both worn hands; the verdict names neither.
	if verdict ~= nil and verdict.kind == "verdict" and item.equipLocation == "INVTYPE_2HWEAPON"
		and equipped.main_hand ~= nil and equipped.off_hand ~= nil then
		local main, off = wornNameOf(equipped.main_hand, nameOf), wornNameOf(equipped.off_hand, nameOf)
		if main ~= nil and off ~= nil then
			reasons[#reasons + 1] = { kind = "compared", color = "muted",
				text = string.format(L.tooltipCompared, main, off) }
		end
	end
	local caps = {}
	if #usable > 0 and current ~= nil then
		for _, text in ipairs(capTexts(item.stats, current.caps)) do
			caps[#caps + 1] = { kind = "cap", text = text, color = "muted" }
		end
	end

	local specBis = ctx.spec ~= nil and type(data.bis) == "table" and data.bis[ctx.spec] or nil
	local currentBand = ctx.faction and Gear.bandFor(specBis, ctx.level) or nil
	local nextBand = ItemTooltip.nextBand(specBis, currentBand)
	-- Below the spec's first band, that band's picks are upcoming.
	local firstBand = ctx.faction and currentBand == nil and Gear.firstBand(specBis, ctx.level) or nil
	if firstBand ~= nil and wornSlot == nil
		and ItemTooltip.plannedBand(data, ctx.spec, ctx.faction, slots, item.itemId, ctx.level) == firstBand then
		-- Case 11: a bag item the first band names is planned, never scored against a band the
		-- character is not in.
		verdict = { kind = "planned", band = firstBand, color = "clientGold",
			text = string.format(L.tooltipYourBisAt, firstBand) }
		reasons = {}
	end
	--- A candidate's gain over what is worn in the slot (net of the worn enchant), rounded as
	--- shown; nil with no figure; third answer true while a weapon waits on the client's stats.
	local function gainOf(itemId)
		local stats, cold
		if weights ~= nil then
			stats, cold = Progress.pickStats(data, statsOf, itemId, standIn)
		end
		if cold then
			return nil, nil, true
		end
		return Progress.pickGain(slot, itemId, stats, equipped, current)
	end
	-- The slot's one target (Progress.targetFor): the band's pick; `unshowable` when the data
	-- names one the client can show none of. The gear gap measures the same target.
	local aimed = Progress.targetFor(data, ctx.spec, currentBand, ctx.faction, slots, slot,
		item.itemId, equipped, isMissing)
	local pick, isPick, pickWorn, unshowable = aimed.pick, aimed.isPick, aimed.worn, aimed.unshowable
	--- The slot of the item's own whose band pick is `itemId` (a ring's second finger), else `slot`.
	local function pickSlotOf(itemId)
		for _, candidate in ipairs(slots) do
			local found = Gear.pick(ctx.spec, currentBand, ctx.faction, candidate, data)
			if found ~= nil and found.item_id == itemId then
				return candidate
			end
		end
		return slot
	end

	--- The published pick a source choice set aside for `pick` (a re-picked slot only), as its Shift
	--- line: its name and how much more it gains than `pick` (`pickGain`), nil when the pick is the
	--- published one or the client has not sent its name yet.
	local function setAsideLine(chosen, pickGain)
		local best = chosen.unfiltered
		if best == nil or best.item_id == chosen.item_id then
			return nil
		end
		local name = pickName(best.item_id)
		if name == nil then
			return nil
		end
		local bestGain = gainOf(best.item_id)
		local gap = type(bestGain) == "number" and type(pickGain) == "number" and bestGain - pickGain or nil
		return { kind = "setAside", itemId = best.item_id, color = "muted",
			text = ItemTooltip.setAsideText(colored(name, qualityHex(best.item_id)), gap, unit) }
	end

	-- The pick's own line: its gain over what is worn, from the floor up (`nextLine`), or why
	-- what is worn holds (`holds`, with `holdGain` the pick's gain under the floor).
	local nextLine, target, sourceLine, setAside
	local holds, holdGain = false, nil
	if pick ~= nil and not isPick and not pickWorn then
		local name = pickName(pick.item_id)
		local gain, nextVerdict, cold = gainOf(pick.item_id)
		local tag = typeTagOf(pick.item_id)
		if cold and name ~= nil then
			local upgrade = { name = name, hex = qualityHex(pick.item_id), unit = unit, tag = tag }
			nextLine = { kind = "next", itemId = pick.item_id, upgrade = upgrade, color = "tooltipBody",
				text = ItemTooltip.nextLine(colored(name, upgrade.hex), nil, unit, tag) }
		elseif gain ~= nil and gain >= Progress.MIN_GAIN and name ~= nil then
			local worn = equipped[slot]
			local upgrade = { name = name, hex = qualityHex(pick.item_id), gain = gain, unit = unit, tag = tag,
				unparsed = worn ~= nil and worn.enchantUnparsed == true }
			nextLine = { kind = "next", itemId = pick.item_id, upgrade = upgrade, color = "tooltipBody",
				text = ItemTooltip.nextLine(colored(name, upgrade.hex), gain, unit, tag),
				breakdown = breakdownLine(nextVerdict, colored(name, upgrade.hex)) }
			if isCached(pick.item_id) then
				target = { itemId = pick.item_id, link = Compat.displayLink(pick.item_id), label = L.tooltipCompareLabel }
			end
		elseif name == nil and (gain == nil or gain >= Progress.MIN_GAIN) then
			-- Not sent yet: wait for it (the hover redraws), never print "…".
			pending[pick.item_id] = true
		elseif gain == nil and weights ~= nil and not isCached(pick.item_id) then
			-- Named from the data, its stats on their way: wait, never a verdict nobody worked out.
			pending[pick.item_id] = true
		else
			holds, holdGain = wornSlot ~= nil, gain
		end
		if nextLine ~= nil then
			local boe = isBoe(pick.item_id)
			local source = ItemTooltip.sourceLine(pick, boe)
			if source ~= nil then
				sourceLine = { kind = "source", pick = pick, boe = boe, text = source, color = "muted" }
			end
			setAside = setAsideLine(pick, nextLine.upgrade.gain)
		end
	elseif pick ~= nil and pickWorn then
		holds = wornSlot ~= nil
	end
	-- Alternatives: the other ways to the pick, under a worn item's next upgrade, or under the
	-- hovered pick itself.
	local alternatives
	local altSource = (wornSlot ~= nil and nextLine ~= nil and pick) or (wornSlot == nil and isPick and pick) or nil
	if altSource ~= nil then
		local entries = {}
		for _, alternative in ipairs(altSource.alternatives or {}) do
			local id = alternative.item_id
			if id ~= nil and id ~= item.itemId and #entries < 2 and not isMissing(id) then
				local name, gain = pickName(id), gainOf(id)
				if name == nil then
					pending[id] = true
				elseif gain == nil or gain >= Progress.MIN_GAIN then
					-- Under the floor it is a sidegrade, never "good": left out, as on the verdict line.
					entries[#entries + 1] = { name = name, hex = qualityHex(id), gain = gain, unit = unit,
						flag = typeTagOf(id) }
				end
			end
		end
		if #entries > 0 then
			alternatives = { kind = "alternatives", entries = entries, color = "muted",
				text = table.concat(ItemTooltip.fitAlternatives(entries), "\n") }
		end
	end
	--- The next rung's pick for the slot, as the At line (`band` the rung, `first` below the
	--- first band), with its source and gain kept for Shift.
	local function atLineFor(band, first)
		local upcoming
		if first then
			upcoming = Gear.pick(ctx.spec, band, ctx.faction, slot, data)
			if upcoming ~= nil and upcoming.item_id == item.itemId then
				upcoming = nil
			end
		else
			upcoming = Gear.pick(ctx.spec, band, ctx.faction, slot, data)
		end
		if upcoming == nil or isMissing(upcoming.item_id) then
			return nil
		end
		local name = pickName(upcoming.item_id)
		if name == nil then
			pending[upcoming.item_id] = true
			return nil
		end
		local boe = isBoe(upcoming.item_id)
		local at = { kind = "at", firstBand = first or nil, band = band, itemId = upcoming.item_id, name = name,
			hex = qualityHex(upcoming.item_id), pick = upcoming, boe = boe,
			source = ItemTooltip.sourceLine(upcoming, boe), color = "muted" }
		if not first then
			local gain = gainOf(upcoming.item_id)
			if gain ~= nil and gain >= Progress.MIN_GAIN then
				at.gain, at.unit = gain, unit
			end
		end
		at.text = ItemTooltip.atText(at)
		return at
	end
	--- The At line's own detail for Shift: its source, then its gain.
	local function atDetail(at)
		local lines = {}
		if at ~= nil and at.source ~= nil then
			lines[#lines + 1] = { kind = "source", pick = at.pick, boe = at.boe, text = at.source, color = "muted",
				detail = true }
		end
		if at ~= nil and at.gain ~= nil then
			lines[#lines + 1] = { kind = "figure", text = ItemTooltip.figure(at.gain, at.unit), color = "muted" }
		end
		return lines
	end

	--- Adds a Shift line unless the same line is already there (two crafted picks from one
	--- profession: the source is said once).
	local function addOnce(line)
		for _, earlier in ipairs(reasons) do
			if earlier.kind == line.kind and earlier.text == line.text then
				return
			end
		end
		reasons[#reasons + 1] = line
	end

	local support
	local noUpgradeText = nextBand ~= nil and string.format(L.tooltipNoUpgradeUntil, nextBand) or L.tooltipNoUpgradeSlot
	if wornSlot ~= nil then
		if unshowable then
			verdict = { kind = "hold", text = noUpgradeText, color = "tooltipBody" }
		elseif firstBand ~= nil then
			-- Case 20: below the first band, the first band's pick is upcoming.
			verdict = { kind = "firstBand", color = "muted", text = string.format(L.tooltipUntilFirstBand, firstBand) }
			support = atLineFor(firstBand, true)
			for _, line in ipairs(atDetail(support)) do
				reasons[#reasons + 1] = line
			end
		elseif nextLine ~= nil then
			-- Case 16 (17 with no figure): the band's pick as the slot's next upgrade, then the band
			-- whose pick changes (owner, 2026-11-13: always the next band's pick with its level), the
			-- upgrade's source under Shift; with no later band, its source under it.
			verdict = nextLine
			local after = ItemTooltip.bandAfter(data, ctx.spec, ctx.faction, pickSlotOf(pick.item_id), currentBand,
				pick.item_id)
			local at = after ~= nil and atLineFor(after, false) or nil
			support = at or sourceLine
			-- The published pick a source choice set aside sits directly under the source line: first
			-- behind Shift while the source is the support line, else right after the source there.
			if setAside ~= nil and at == nil then
				reasons[#reasons + 1] = setAside
			end
			if nextLine.breakdown ~= nil then
				reasons[#reasons + 1] = nextLine.breakdown
			end
			if at ~= nil and sourceLine ~= nil then
				reasons[#reasons + 1] = { kind = "source", pick = sourceLine.pick, boe = sourceLine.boe,
					text = sourceLine.text, color = "muted", detail = true }
			end
			if setAside ~= nil and at ~= nil then
				reasons[#reasons + 1] = setAside
			end
			for _, line in ipairs(atDetail(at)) do
				addOnce(line)
			end
		elseif holds then
			-- Case 18: worn and the pick; case 19: the pick gains under the floor. Both name the band
			-- the pick holds for, then the band whose pick changes (owner, 2026-11-13).
			local pickId = pick.item_id
			local after = ItemTooltip.bandAfter(data, ctx.spec, ctx.faction, pickSlotOf(pickId), currentBand, pickId)
			local name = (not isPick and not pickWorn) and pickName(pickId) or nil
			if isPick then
				verdict = { kind = "bis", text = ItemTooltip.bestText(currentBand, after), color = "clientGold" }
			elseif name == nil then
				-- The pick is worn in the item's other slot, or its name has not come yet.
				verdict = { kind = "hold", text = noUpgradeText, color = "tooltipBody" }
			else
				local hex = qualityHex(pickId)
				local function build(shown)
					return ItemTooltip.holdText(currentBand, after, colored(shown, hex), holdGain)
				end
				verdict = { kind = "hold", itemId = pickId, text = build(name), color = "tooltipBody",
					fit = { name = name, build = build } }
				-- Shift: the pick's own figure, then where it comes from.
				if holdGain ~= nil and holdGain ~= 0 then
					reasons[#reasons + 1] = { kind = "figure", color = "muted",
						text = string.format(L.tooltipPickFigure, ItemTooltip.figure(holdGain, unit)) }
				end
				local boe = isBoe(pickId)
				local source = ItemTooltip.sourceLine(pick, boe)
				if source ~= nil then
					reasons[#reasons + 1] = { kind = "source", pick = pick, boe = boe, text = source, color = "muted",
						detail = true }
				end
			end
			support = after ~= nil and atLineFor(after, false) or nil
			for _, line in ipairs(atDetail(support)) do
				addOnce(line)
			end
		elseif pick == nil and next(pending) == nil then
			-- Case 21: the band names no pick for the slot.
			verdict = { kind = "noPick", text = L.tooltipNoPick, color = "muted" }
		else
			verdict = nil
		end
	elseif verdict ~= nil and verdict.kind ~= "refusal" and verdict.kind ~= "planned" then
		if isPick then
			-- Case 8: the band's pick, gold ("Best for 40 to 49"); an upgrade keeps its figure and says
			-- so under it.
			local after = ItemTooltip.bandAfter(data, ctx.spec, ctx.faction, pickSlotOf(pick.item_id), currentBand,
				pick.item_id)
			local best = { kind = "bis", text = ItemTooltip.bestText(currentBand, after), color = "clientGold" }
			if verdict.upgrade then
				support = best
			else
				verdict = best
			end
		elseif verdict.upgrade and #caps > 0 then
			-- A cap the item's upgrade leans on changes the answer: it takes the support line.
			support = table.remove(caps, 1)
		elseif nextLine ~= nil and nextLine.upgrade.gain ~= nil and (verdict.upgrade or verdict.side)
			and (verdict.gain == nil or nextLine.upgrade.gain > verdict.gain) then
			-- Case 9: the band's pick beats it. "Better: <pick> +3.5 · <short source>".
			local upgrade = nextLine.upgrade
			local gainText = ItemTooltip.alternativeGain(upgrade.gain, unit)
			local short = ItemTooltip.shortSource(pick)
			local function build(name)
				local shown = colored(name, upgrade.hex)
				if short ~= nil then
					return string.format(L.tooltipBetter, shown, gainText, short)
				end
				return string.format(L.tooltipBetterBare, shown, gainText)
			end
			support = { kind = "better", itemId = upgrade.itemId or pick.item_id, color = "muted",
				text = build(upgrade.name), fit = { name = upgrade.name, build = build } }
		end
	elseif verdict == nil and pick == nil and not isPick and next(pending) == nil then
		-- Case 14: gear the weights score nothing on and no pick names.
		verdict = { kind = "notForBuild", text = L.tooltipNotForBuild, color = "muted" }
	end

	local lines = {}
	if verdict == nil then
		return lines, target, pending, head
	end
	lines[1] = as("verdict", verdict)
	if support ~= nil then
		lines[2] = as("support", support)
	end
	local function shift(line)
		if line ~= nil then
			lines[#lines + 1] = as("shift", line)
		end
	end
	for _, line in ipairs(reasons) do
		shift(line)
	end
	shift(alternatives)
	for _, line in ipairs(caps) do
		shift(line)
	end
	local worn = equipped[slot]
	if worn ~= nil and worn.enchantUnparsed == true and (verdict.gain ~= nil or verdict.kind == "next") then
		shift({ kind = "enchant", text = L.tooltipEnchantNote, color = "muted" })
	end
	-- Where the weights came from, last, under every answer but a refusal or "no pick".
	if verdict.kind ~= "refusal" and verdict.kind ~= "noPick" then
		shift({ kind = "attribution", text = ItemTooltip.attributionText(specWord, head.tag), color = "tooltipDim" })
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

--- The section for an EMPTY paper-doll slot (spec 2.2 case 22), in ItemTooltip.section's
--- shape (lines, the compare tooltip's item, the ids still pending, the head): the band's pick
--- for `slot` (Progress.slotPicks' pair rule), the first of its alternatives the client can
--- show when it cannot show the pick, as the worn item's Next line with no gain (nothing is
--- worn to gain over), its source under it, the attribution under Shift. Below the first
--- band, "Picks start at level 20" and the At line; with no pick, "No pick for this slot yet".
--- `ctx` is ItemTooltip.section's without `item`.
function ItemTooltip.emptySlotSection(ctx, slot)
	local lines, pending = {}, {}
	if not ctx.spec then
		return lines, nil, pending, nil
	end
	local data = ctx.data or {}
	local current = ctx.current or nil
	local equipped = ctx.equipped or {}
	local qualityHex = ctx.qualityHex or qualityHexOf
	local head = { specWord = specName(data, ctx.spec), tag = current ~= nil and current.tag or nil, equipped = false }
	if slot == "off_hand" and equipped.main_hand ~= nil and equipped.main_hand.equipLoc == "INVTYPE_2HWEAPON" then
		lines[1] = { text = L.tooltipTwoHander, color = "muted", role = "verdict" }
		return lines, nil, pending, head
	end
	local noPick = { kind = "noPick", text = L.tooltipNoPick, color = "muted", role = "verdict" }
	local specBis = type(data.bis) == "table" and data.bis[ctx.spec] or nil
	local band = ctx.faction and Gear.bandFor(specBis, ctx.level) or nil
	local first = ctx.faction and band == nil and Gear.firstBand(specBis, ctx.level) or nil
	if band == nil and first == nil then
		lines[1] = noPick
		return lines, nil, pending, head
	end
	local isMissing = ctx.isMissing or missingOf
	local nameOf = ctx.nameOf or Compat.itemName
	if first ~= nil then
		lines[1] = { kind = "firstBand", color = "muted", role = "verdict",
			text = string.format(L.tooltipUntilFirstBand, first) }
		local upcoming = reachablePick(Progress.slotPicks(data, ctx.spec, first, ctx.faction, equipped)[slot],
			isMissing)
		if upcoming ~= nil then
			local name = pickNameOf(data, nameOf, pending, upcoming.item_id)
			if name ~= nil then
				local at = { kind = "at", firstBand = true, band = first, itemId = upcoming.item_id, name = name,
					hex = qualityHex(upcoming.item_id), pick = upcoming, color = "muted", role = "support",
					source = ItemTooltip.sourceLine(upcoming) }
				at.text = ItemTooltip.atText(at)
				lines[2] = at
				if at.source ~= nil then
					lines[3] = { kind = "source", pick = upcoming, text = at.source, color = "muted", role = "shift",
						detail = true }
				end
			end
		end
		return lines, nil, pending, head
	end
	local picks = Progress.slotPicks(data, ctx.spec, band, ctx.faction, equipped)
	local pick = reachablePick(picks[slot], isMissing)
	if pick == nil then
		local _, filteredOut = Gear.pick(ctx.spec, band, ctx.faction, slot, data)
		lines[1] = filteredOut and { kind = "noPick", text = ItemTooltip.noSlotFromSourcesText(slot), color = "muted",
			role = "verdict" } or noPick
		return lines, nil, pending, head
	end
	local itemId = pick.item_id
	local name = pickNameOf(data, nameOf, pending, itemId)
	if name == nil then
		pending[itemId] = true
		return lines, nil, pending, head
	end
	local upgrade = { name = name, hex = qualityHex(itemId), unit = ItemTooltip.unitOf(current) }
	lines[1] = { kind = "next", itemId = itemId, upgrade = upgrade, color = "tooltipBody", role = "verdict",
		text = ItemTooltip.nextLine(colored(name, upgrade.hex)) }
	local source = ItemTooltip.sourceLine(pick)
	if source ~= nil then
		lines[#lines + 1] = { kind = "source", pick = pick, text = source, color = "muted", role = "support" }
	end
	lines[#lines + 1] = { kind = "attribution", role = "shift", color = "tooltipDim",
		text = ItemTooltip.attributionText(head.specWord, head.tag) }
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
local cacheSize = 0

--- Most sections kept at once (L9.2 B3): a long session hovering every item of a vendor or a
--- bank must not grow idle memory without bound. Past it the cache starts again; the next
--- hover recomputes one section.
ItemTooltip.CACHE_LIMIT = 200

function ItemTooltip.resetCache()
	ItemTooltip.cache = {}
	cacheSize = 0
	return ItemTooltip.cache
end

local function currentSpec(data)
	local classSlug = Talents.playerClassSlug()
	if classSlug == nil or data == nil or data.classes == nil then
		return nil, nil
	end
	return Follow.currentSpec(data, classSlug), classSlug
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
		-- What the bags carry, read when needed: the pair verdict for a one-hander under a worn
		-- two-hander (L3.7 item 11).
		bags = Gear.bagLinks,
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
		requiredLevel = Compat.itemMinLevel(itemLink),
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
		if cacheSize >= ItemTooltip.CACHE_LIMIT then
			ItemTooltip.resetCache()
		end
		ItemTooltip.cache[itemLink] = entry
		cacheSize = cacheSize + 1
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

--- AddLine's wrap flag on every line this addon adds to an item tooltip: wrapped at the tooltip's
--- width, never cut.
local WRAP = true

--- One section line as the lines it is drawn on: whole without `measure`, else fitted to `room`.
function ItemTooltip.fitLine(line, measure, room)
	if line.kind == "alternatives" then
		return ItemTooltip.fitAlternatives(line.entries, measure, room)
	elseif line.kind == "next" then
		return ItemTooltip.fitNext(line.upgrade, measure, room)
	elseif line.kind == "source" then
		return ItemTooltip.fitSource(line.pick, measure, room, line.boe)
	elseif line.kind == "bag" then
		return { ItemTooltip.fitBagLine(line.upgrade, measure, room) }
	elseif line.kind == "at" then
		return ItemTooltip.fitAt(line, measure, room)
	elseif line.kind == "breakdown" then
		return ItemTooltip.fitBreakdown(line.contributions, measure, room, line.of)
	end
	return { ItemTooltip.fitNamed(line, measure, room) }
end

--- Whether a section has anything drawn only while Shift is held.
function ItemTooltip.hasShift(lines)
	for _, line in ipairs(lines or {}) do
		if line.role == "shift" then
			return true
		end
	end
	return false
end

--- The lines drawn now: every line, the Shift block only while `shift` is held. Pure.
function ItemTooltip.visibleLines(lines, shift)
	local shown = {}
	for _, line in ipairs(lines or {}) do
		if line.role ~= "shift" or shift then
			shown[#shown + 1] = line
		end
	end
	return shown
end


--- Whether a section may show `tooltip`: it is up already, or its owner is on screen under the
--- mouse (the hover that just set it). A tooltip the client rebuilt while hidden (its data
--- updated after the hover left, a new item in the bags) stays hidden.
function ItemTooltip.mayShow(tooltip)
	if type(tooltip.IsShown) ~= "function" or tooltip:IsShown() then
		return true
	end
	local owner = type(tooltip.GetOwner) == "function" and tooltip:GetOwner() or nil
	return type(owner) == "table" and type(owner.IsVisible) == "function" and owner:IsVisible() == true
		and type(owner.IsMouseOver) == "function" and owner:IsMouseOver() == true
end

--- Draws a section (`entry`: { lines, target, pending, head }) directly after the lines
--- already on `tooltip`, with no blank row and no header (G1): each line whole in its colour and
--- wrapping inside the tooltip as the client's own lines do (2.0.43: a long item name or "where it
--- drops" was cut with an ellipsis), the Shift block only while Shift is held. Then the compare
--- tooltip for the entry's target, and the hover `redraw` re-runs once an item it waits for
--- arrives (ItemTooltip.onItemInfoReceived).
local function drawEntry(tooltip, entry, redraw)
	for _, line in ipairs(ItemTooltip.visibleLines(entry.lines, Theme.shiftDown())) do
		local r, g, b = opColor(line.color)
		for _, text in ipairs(ItemTooltip.fitLine(line)) do
			tooltip:AddLine(text, r, g, b, WRAP)
		end
	end
	-- A tooltip drawn inside our own window (ItemTooltip.embed) is shown by its owner, never compared.
	if ItemTooltip.embedded[tooltip] or not ItemTooltip.mayShow(tooltip) then
		return
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

--- Injections (SlotMarks sets them): whether an Alt-click on paper-doll `slot` opens the Gear tab, and
--- whether one on bag item `link` under its own button `owner` does.
ItemTooltip.altClickOpensGear = nil
ItemTooltip.bagAltClickOpensGear = nil

--- Whether `ask(...)` answers true; false for no question or a failed one (never a broken hover).
local function answers(ask, ...)
	if ask == nil then
		return false
	end
	local ok, yes = pcall(ask, ...)
	return ok and yes == true
end

--- `entry` with the Shift block's last line, the marks' one wording "Alt-click: open Gear", when the hover
--- is on a mark an Alt-click opens the Gear tab from (spec 2026-11-13-addon-hub 2, 3.7, 5): paper-doll
--- `slot`'s own marked button, or (`slot` nil) bag item `link` under its own button carrying the arrow.
--- The entry itself otherwise.
local function withAltClick(entry, slot, owner, link)
	if owner == nil then
		return entry
	end
	local opens
	if slot ~= nil then
		local buttonName = ItemTooltip.SLOT_BUTTONS[slot]
		opens = buttonName ~= nil and Compat.clientGlobal(buttonName, "table") == owner
			and answers(ItemTooltip.altClickOpensGear, slot)
	else
		opens = answers(ItemTooltip.bagAltClickOpensGear, link, owner)
	end
	if not opens then
		return entry
	end
	local copy, lines = {}, {}
	for key, value in pairs(entry) do
		copy[key] = value
	end
	for index, line in ipairs(entry.lines) do
		lines[index] = line
	end
	lines[#lines + 1] = { kind = "altClick", role = "shift", color = "tooltipDim", text = L.markAltGear }
	copy.lines = lines
	return copy
end
ItemTooltip.withAltClick = withAltClick

--- The lines dropped, in order, when a tooltip would pass the top of the screen (spec 1.1): the
--- support line first, then the Shift block's alternatives. Keys are roles or kinds.
ItemTooltip.CAP_KINDS = { { support = true }, { support = true, alternatives = true } }

--- Whether `tooltip` reaches the top of the screen (within 2 px): the client would clamp or cut
--- it. False wherever the client cannot say.
function ItemTooltip.passesScreenTop(tooltip)
	local parent = Compat.clientGlobal("UIParent", "table")
	if parent == nil or type(tooltip.GetTop) ~= "function" or type(parent.GetTop) ~= "function" then
		return false
	end
	local ok, top, screenTop = pcall(function()
		local scale = type(tooltip.GetEffectiveScale) == "function" and tooltip:GetEffectiveScale() or 1
		local parentScale = type(parent.GetEffectiveScale) == "function" and parent:GetEffectiveScale() or 1
		return (tooltip:GetTop() or 0) * scale, (parent:GetTop() or 0) * parentScale
	end)
	return ok and type(top) == "number" and type(screenTop) == "number" and screenTop > 0 and top >= screenTop - 2
end

--- `lines` without the kinds cap level `level` drops (0 drops none).
local function capLines(lines, level)
	local drop = ItemTooltip.CAP_KINDS[level]
	if drop == nil then
		return lines
	end
	local kept = {}
	for _, line in ipairs(lines) do
		if not drop[line.role] and not drop[line.kind] then
			kept[#kept + 1] = line
		end
	end
	return kept
end

--- The cap level of the rebuild in progress for this tooltip and link (0 for any other draw).
local rebuilding

local function addLines(tooltip, itemLink)
	-- A hover from before (another item) must not be redrawn onto this one.
	ItemTooltip.activeHover = nil
	local level = rebuilding ~= nil and rebuilding.tooltip == tooltip and rebuilding.link == itemLink
		and rebuilding.level or 0
	local entry = ItemTooltip.sectionFor(itemLink)
	local head = entry.head or {}
	-- The worn item of a slot a bag item beats: the bag fact leads, alone (L4.1 fix round 5).
	if head.equipped then
		entry = bagLead(head.slot, head) or entry
	end
	if #entry.lines == 0 then
		return
	end
	local owner = type(tooltip.GetOwner) == "function" and tooltip:GetOwner() or nil
	if head.equipped then
		entry = withAltClick(entry, head.slot, owner)
	else
		entry = withAltClick(entry, nil, owner, itemLink)
	end
	local shown = entry
	if level > 0 then
		shown = {}
		for key, value in pairs(entry) do
			shown[key] = value
		end
		shown.lines = capLines(entry.lines, level)
	end
	drawEntry(tooltip, shown, function()
		-- Re-running SetHyperlink re-fires the hook that called
		-- onTooltip, so the tooltip is rebuilt with the data that
		-- just arrived rather than gaining a second copy of the section.
		if type(tooltip.SetHyperlink) == "function"
			and type(tooltip.IsShown) == "function" and tooltip:IsShown() then
			tooltip:SetHyperlink(itemLink)
		end
	end)
	ItemTooltip.noteShiftHover(tooltip, entry.lines, function()
		return ItemTooltip.itemLinkFrom(tooltip) == itemLink
	end, function()
		ItemTooltip.redrawItem(tooltip, itemLink)
	end)
	-- Past the top of the screen: rebuild with the next row group dropped (L3.7 item 12).
	if level < #ItemTooltip.CAP_KINDS and ItemTooltip.passesScreenTop(tooltip)
		and type(tooltip.SetHyperlink) == "function" then
		local before = rebuilding
		rebuilding = { tooltip = tooltip, link = itemLink, level = level + 1 }
		local ok, err = pcall(tooltip.SetHyperlink, tooltip, itemLink)
		rebuilding = before
		if not ok then
			error(err, 0)
		end
	end
end

--- Whether `tooltip` is another addon's hidden scanning tooltip: owned with ANCHOR_NONE and placed
--- nowhere (a tooltip the client or an addon shows sits at a point, the default anchor included).
--- Found in game: Pawn read a looted robe through PawnPrivateTooltip1, the processor ran this hook
--- for it, and the section's compare tooltip opened beside GameTooltip on its own.
function ItemTooltip.isScanner(tooltip)
	if type(tooltip.GetAnchorType) ~= "function" or type(tooltip.GetNumPoints) ~= "function" then
		return false
	end
	local ok, anchor, points = pcall(function()
		return tooltip:GetAnchorType(), tooltip:GetNumPoints()
	end)
	return ok and anchor == "ANCHOR_NONE" and points == 0
end

--- The tooltips drawn inside the addon's own window (weak keys): Set up's picture of a worn item
--- (TourExample). The item section is drawn on them as on any hover; the owner shows them and they
--- never open a compare tooltip nor count as the player's hover (ItemTooltip.onHover).
ItemTooltip.embedded = setmetatable({}, { __mode = "k" })

--- Injection: runs after the player's own item hover drew (Onboarding.noteHover, Set up's step 2).
ItemTooltip.onHover = nil

--- Marks `tooltip` as one drawn inside our window. On a client without the tooltip processor the item
--- hook sits on GameTooltip alone, so the embedded tooltip gets its own.
function ItemTooltip.embed(tooltip)
	ItemTooltip.embedded[tooltip] = true
	if ItemTooltip.how == "legacy" and type(tooltip.HookScript) == "function" then
		tooltip:HookScript("OnTooltipSetItem", function(self)
			ItemTooltip.onTooltip(self, ItemTooltip.itemLinkFrom(self))
		end)
	end
	return tooltip
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
	if tooltip == Compat.scanFrame or ItemTooltip.isScanner(tooltip) then
		return
	end
	local ok, err = pcall(addLines, tooltip, itemLink)
	if not ok then
		ItemTooltip.disabled = true
		Theme.note(string.format(L.diagTooltipHookFailed, tostring(err)))
		return
	end
	if ItemTooltip.onHover ~= nil and not ItemTooltip.embedded[tooltip] then
		local hovered, hoverErr = pcall(ItemTooltip.onHover)
		if not hovered then
			Theme.note(string.format(L.diagOnboardingFailed, tostring(hoverErr)))
		end
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
	entry = withAltClick(entry, slot, button)
	drawEntry(tooltip, entry, redraw)
	appendedSection[tooltip] = true
	ItemTooltip.noteShiftHover(tooltip, entry.lines, function()
		return type(tooltip.GetOwner) ~= "function" or tooltip:GetOwner() == button
	end, redraw)
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
	Theme.noteFrame("Item info events", "login")
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

--- The hover whose Shift block a press or release of Shift redraws ({ tooltip, still, redraw }:
--- `still()` says whether the tooltip shows the same thing), or nil.
ItemTooltip.shiftHover = nil

--- Remembers the hover just drawn on `tooltip` when its section has a Shift block (G4), so a
--- press or release of Shift redraws it; forgets it otherwise.
function ItemTooltip.noteShiftHover(tooltip, lines, still, redraw)
	if ItemTooltip.hasShift(lines) then
		ItemTooltip.shiftHover = { tooltip = tooltip, still = still, redraw = redraw }
	elseif ItemTooltip.shiftHover ~= nil and ItemTooltip.shiftHover.tooltip == tooltip then
		ItemTooltip.shiftHover = nil
	end
end

--- Draws an item tooltip again from its own source: the client's RefreshData where it has it
--- (a bag or slot item keeps its bag lines and compare), else SetHyperlink with the link. Either
--- re-fires the hook, so the section is rebuilt rather than added twice.
function ItemTooltip.redrawItem(tooltip, itemLink)
	if type(tooltip.IsShown) ~= "function" or not tooltip:IsShown() then
		return
	end
	if type(tooltip.RefreshData) == "function" then
		tooltip:RefreshData()
	elseif type(tooltip.SetHyperlink) == "function" then
		tooltip:SetHyperlink(itemLink)
	end
end

local SHIFT_KEYS = { LSHIFT = true, RSHIFT = true }

--- MODIFIER_STATE_CHANGED: Shift going down or up over an item hover with a Shift block draws
--- it again, with or without the details (spec 2.3). Never in combat; nothing polls. True when
--- it redrew.
function ItemTooltip.onModifierChanged(key)
	local hover = ItemTooltip.shiftHover
	if not SHIFT_KEYS[key] or hover == nil or ItemTooltip.disabled or not Prefs.flag("tooltip") then
		return false
	end
	if Theme.inCombat() then
		return false
	end
	local tooltip = hover.tooltip
	if type(tooltip.IsShown) ~= "function" or not tooltip:IsShown() or not hover.still() then
		ItemTooltip.shiftHover = nil
		return false
	end
	hover.redraw()
	return true
end

--- The frame Shift is heard on. Idempotent.
function ItemTooltip.registerShiftRedraw()
	if ItemTooltip.shiftFrame ~= nil then
		return ItemTooltip.shiftFrame
	end
	local frame = CreateFrame("Frame")
	Theme.noteFrame("Item tooltip Shift", "login")
	Theme.tryRegisterEvent(frame, "MODIFIER_STATE_CHANGED")
	frame:SetScript("OnEvent", function(_, _, key)
		local ok, err = pcall(ItemTooltip.onModifierChanged, key)
		if not ok and not ItemTooltip.disabled then
			ItemTooltip.disabled = true
			Theme.note(string.format(L.diagTooltipHookFailed, tostring(err)))
		end
	end)
	ItemTooltip.shiftFrame = frame
	return frame
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
	ItemTooltip.registerShiftRedraw()
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
