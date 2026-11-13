-- addon/ForeverSixty/surfaces/GearTab.lua
-- The Gear tab's model (design/specs/2026-11-13-addon-hub.md 4.3): what you wear, what beats it and where
-- you get it. The doll by band (each slot's worn item against the band's pick: Best in slot, the gain,
-- in your bags, No enchant), the upgrades with their DPS gains and total, the stat weights normalized to
-- the top stat, and the source filter with its counts. Pure over the inputs GearTab.read gathers once a
-- draw: every gain is Progress.slotTargets' (the tooltip's Next line, the slot marks and the guild gap
-- read the same one), every source line the tooltip's own (ItemTooltip.sourceParts: the drop chance only
-- where the data carries one), every pick Gear.pick under the character's source choice (Sources).
-- The frames are surfaces/GearTabView.lua's.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Compat = ns.Compat or require("Compat")
local Rules = ns.Rules or require("Rules")
local Gear = ns.Gear or require("Gear")
local Progress = ns.Progress or require("Progress")
local Sources = ns.Sources or require("Sources")
local Inspect = ns.Inspect or require("Inspect")
local ItemTooltip = ns.ItemTooltip or require("ItemTooltip")

local GearTab = {}

--- The Level control's bands (the web BiS page's), lowest first.
GearTab.BANDS = { 20, 30, 40, 50, 60 }
--- The doll's two columns, top to bottom, as the character pane draws them: the off hand follows the
--- main hand only beside a one-hander.
GearTab.LEFT = { "head", "neck", "shoulder", "back", "chest", "wrist", "main_hand", "off_hand", "ranged" }
GearTab.RIGHT = { "hands", "waist", "legs", "feet", "finger1", "finger2", "trinket1", "trinket2" }
--- A weight under this share of the top stat's is left out of the table: it moves no choice.
GearTab.WEIGHT_FLOOR = 0.05
--- The most Upgrades rows, and the most Stat weights rows.
GearTab.UPGRADES = 4
GearTab.WEIGHT_ROWS = 7
--- The weights rating haste counts at, summed (Gear.score's rule), shown apart per 1%.
local HASTE = {}
for _, stat in ipairs(Rules.HASTE_WEIGHTS) do
	HASTE[stat] = true
end

local TWO_HANDED = "INVTYPE_2HWEAPON"

-- What the client says now -----------------------------------------------------------------------

--- The ids `links` hold, as a set.
local function idsOf(links)
	local ids = {}
	for _, link in ipairs(links or {}) do
		local id = tonumber((Compat.itemInfoInstant(link)))
		if id ~= nil then
			ids[id] = true
		end
	end
	return ids
end

--- The enchantable slots the player wears an item in without an enchant, as a set (Inspect on the player).
local function unenchanted()
	local missing = {}
	for _, slot in ipairs(Inspect.missingEnchants(Inspect.readSelf())) do
		missing[slot] = true
	end
	return missing
end

--- Everything a draw reads from the client, once: the item tooltip's own context (the spec, its weights,
--- the level, the faction, what is worn), the bags' item ids and the slots missing an enchant.
function GearTab.read(data)
	data = data or ns.Data
	local base = ItemTooltip.baseContext(data)
	return {
		data = data, spec = base.spec, current = base.current, level = base.level, faction = base.faction,
		equipped = base.equipped or {}, bagIds = idsOf(Gear.bagLinks()), noEnchant = unenchanted(),
		isMissing = ItemTooltip.isMissing,
	}
end

-- The band ---------------------------------------------------------------------------------------

local function specBis(input)
	local bis = type(input.data) == "table" and input.data.bis or nil
	return input.spec ~= nil and type(bis) == "table" and bis[input.spec] or nil
end

--- The character's own band: Gear.bandFor at its level, else (below the first band) that first band.
function GearTab.ownBand(input)
	local bySpec = specBis(input)
	return Gear.bandFor(bySpec, input.level) or Gear.firstBand(bySpec, input.level)
end

--- Whether the data carries a list for `band` on the character's faction.
function GearTab.hasList(input, band)
	local bySpec = specBis(input)
	local byBand = bySpec ~= nil and band ~= nil and bySpec[band] or nil
	return type(byBand) == "table" and type(byBand[input.faction]) == "table"
end

-- The doll ---------------------------------------------------------------------------------------

local function wornName(worn)
	if worn == nil then
		return nil
	end
	local name = type(worn.link) == "string" and worn.link:match("%[(.-)%]") or nil
	return name or Compat.itemName(worn.itemId)
end

--- The slot's pick as the target names it, with its source, kind, chance and standing: the band's pick
--- under the source choice, or its first alternative the client can show (Progress.reachablePick).
local function pickOf(input, band, slot, sources, itemId)
	local pick = Progress.reachablePick(Progress.slotTarget(input.data, input.spec, band, input.faction, slot,
		sources), input.isMissing or Progress.isMissing)
	if pick ~= nil and pick.item_id == itemId then
		return pick
	end
	return { item_id = itemId }
end

--- One cell: { slot, worn = { itemId, link, name }, pick (Gear.pick's table, nil without a target),
--- gain (as the tooltip shows it), exact (unrounded), holds (you wear the pick), inBags, noEnchant }.
local function cellOf(input, band, slot, row, sources)
	local worn = input.equipped[slot]
	local cell = { slot = slot, noEnchant = input.noEnchant[slot] == true }
	if worn ~= nil then
		cell.worn = { itemId = worn.itemId, link = worn.link, name = wornName(worn) }
	end
	if row ~= nil then
		cell.pick = pickOf(input, band, slot, sources, row.itemId)
		cell.holds = row.holds == true
		cell.gain = row.gain
		cell.exact = row.holds and 0 or (row.verdict ~= nil and row.verdict.delta ~= nil and -row.verdict.delta or row.gain)
		cell.inBags = not cell.holds and input.bagIds[row.itemId] == true
	end
	return cell
end

--- Whether the main hand holds a two-hander: the off hand leaves the doll.
local function twoHanded(equipped)
	local main = equipped.main_hand
	return main ~= nil and main.equipLoc == TWO_HANDED
end

--- The doll at `band`: { band, list (the data has one for the band), left, right (cells in the doll's
--- order), cells (every cell by slot), unit }. Each slot is scored by Progress.slotTargets under the
--- character's source choice at that band. Pure over `input` (GearTab.read's).
function GearTab.doll(input, band)
	local sources = Sources.state(band)
	local list = GearTab.hasList(input, band)
	local rows = {}
	if list then
		for _, row in ipairs(Progress.slotTargets(input.data, input.spec, band, input.faction, input.equipped,
			input.current, input.statsOf, input.isMissing, sources)) do
			rows[row.slot] = row
		end
	end
	local doll = { band = band, list = list, left = {}, right = {}, cells = {}, spec = input.spec,
		unit = ItemTooltip.unitOf(input.current) }
	local hideOffHand = twoHanded(input.equipped)
	for side, slots in pairs({ left = GearTab.LEFT, right = GearTab.RIGHT }) do
		for _, slot in ipairs(slots) do
			if not (slot == "off_hand" and hideOffHand) then
				local cell = cellOf(input, band, slot, rows[slot], sources)
				doll[side][#doll[side] + 1] = cell
				doll.cells[slot] = cell
			end
		end
	end
	return doll
end

--- How the doll stands: { best (cells holding the pick), scored (cells with a target), missing (cells
--- without an enchant) }. Pure.
function GearTab.tally(doll)
	local best, scored, missing = 0, 0, 0
	for _, cell in pairs(doll.cells) do
		if cell.pick ~= nil then
			scored = scored + 1
			-- A worn item the list's pick does not beat is as good as best in slot.
			local beaten = type(cell.gain) == "number" and cell.gain > 0
			best = best + ((cell.holds or (type(cell.gain) == "number" and not beaten)) and 1 or 0)
		end
		if cell.noEnchant and cell.worn ~= nil then
			missing = missing + 1
		end
	end
	return { best = best, scored = scored, missing = missing }
end

--- The header's figures (spec 3.2): "9 of 16 best in slot", "2 enchants missing" (none: left out). Pure.
function GearTab.figures(doll)
	local tally = GearTab.tally(doll)
	local parts = {}
	if tally.scored > 0 then
		parts[#parts + 1] = string.format(L.gearBestOf, tally.best, tally.scored)
	end
	if tally.missing == 1 then
		parts[#parts + 1] = L.gearEnchantMissing
	elseif tally.missing > 1 then
		parts[#parts + 1] = string.format(L.gearEnchantsMissing, tally.missing)
	end
	return parts
end

-- Upgrades ---------------------------------------------------------------------------------------

--- The upgrades: { rows = the top GearTab.UPGRADES cells by gain, rest = the others in the same order,
--- total = every slot's unrounded gain summed, count }. A pick two slots share counts once. Pure.
function GearTab.upgrades(doll)
	local found, seen, total = {}, {}, 0
	for _, cell in pairs(doll.cells) do
		local gain = cell.gain
		if cell.pick ~= nil and not cell.holds and type(gain) == "number" and gain > 0
			and not seen[cell.pick.item_id] then
			seen[cell.pick.item_id] = true
			found[#found + 1] = cell
			total = total + (cell.exact or gain)
		end
	end
	local rank = {}
	for index, slot in ipairs(Progress.slotOrder(doll.spec)) do
		rank[slot] = index
	end
	table.sort(found, function(a, b)
		if a.gain ~= b.gain then
			return a.gain > b.gain
		end
		return (rank[a.slot] or 99) < (rank[b.slot] or 99)
	end)
	local rows, rest = {}, {}
	for index, cell in ipairs(found) do
		local list = index <= GearTab.UPGRADES and rows or rest
		list[#list + 1] = cell
	end
	return { rows = rows, rest = rest, total = total, count = #found }
end

--- Whether `text` stays within `room` as `measure` reads it (always, with nothing to measure with).
local function fitsIn(measure, room)
	return function(text)
		local width = type(measure) == "function" and room ~= nil and measure(text) or nil
		return type(width) ~= "number" or width <= room
	end
end

--- The kinds whose source already names who gives it (a mob, a vendor): the kind word is left off.
local NAMED_KINDS = { W = true, V = true }

--- An upgrade's source in the game's words: ItemTooltip.sourceParts', the kind word left off where the
--- source names a mob or a vendor ("Pix Xizzix", not "Pix Xizzix · Vendor"); a standing stays.
local function sourceParts(pick)
	local first, second, percent = ItemTooltip.sourceParts(pick)
	if second ~= nil and NAMED_KINDS[pick.kind] and Sources.standingWord(pick.standing) == nil then
		second = nil
	end
	return first, second, percent
end

--- An upgrade's whole second line: the slot and the source with its drop chance ("Feet · Scholomance ·
--- Kirtonos the Herald · 12% drop", the chance only where the data has one), or the slot and "In your
--- bags" in the client gold. The row's hover shows it whole. Pure.
function GearTab.fullWhere(cell)
	local slot = L.gearSlotLabels[cell.slot] or cell.slot
	if cell.inBags then
		return string.format(L.gearWhere, slot, Theme.inlineColor("clientGold", L.gearInBagsLead))
	end
	local first, second, percent = sourceParts(cell.pick)
	return first ~= nil and string.format(L.gearWhere, slot, ItemTooltip.sourceText(first, second, percent)) or slot
end

--- `text` cut at a word: the most whole words that fit with the ellipsis after them, or nil when not
--- even the first word does. Never a word cut in two. Pure.
function GearTab.wordCut(text, fits)
	if fits(text) then
		return text
	end
	local words = {}
	for word in text:gmatch("%S+") do
		words[#words + 1] = word
	end
	for count = #words - 1, 1, -1 do
		local cut = table.concat(words, " ", 1, count):gsub("[%s\194\183,:]+$", "") .. L.tooltipNamePending
		if fits(cut) then
			return cut
		end
	end
	return nil
end

--- The second line within `room` as `measure` reads it (both optional), whole words only: whole when it
--- fits; else the slot gives way (the icon says it); else the drop chance (the hover keeps it); else the
--- place is cut at a word, the boss never. nil when even that does not fit (the caller leaves the source
--- to the hover). Pure.
function GearTab.whereLine(cell, measure, room)
	local whole = GearTab.fullWhere(cell)
	local fits = fitsIn(measure, room)
	if cell.inBags or fits(whole) then
		return whole
	end
	local first, second, percent = sourceParts(cell.pick)
	if first == nil then
		return GearTab.wordCut(whole, fits)
	end
	for _, chance in ipairs(percent ~= nil and { percent, false } or { false }) do
		local text = ItemTooltip.sourceText(first, second, chance or nil)
		if fits(text) then
			return text
		end
	end
	if second == nil then
		return GearTab.wordCut(first, fits)
	end
	local tail = ItemTooltip.sourceText(" ", second, nil):gsub("^ ", "")
	local place = GearTab.wordCut(first, function(cut)
		return fits(cut .. tail)
	end)
	return place ~= nil and place .. tail or nil
end

--- The slots missing an enchant, in doll order, as the line under the doll says them ("Chest, Feet"), or
--- nil with none. Pure.
function GearTab.unenchantedLine(doll)
	local names = {}
	for _, cells in ipairs({ doll.left, doll.right }) do
		for _, cell in ipairs(cells) do
			if cell.noEnchant and cell.worn ~= nil then
				names[#names + 1] = L.gearSlotLabels[cell.slot] or cell.slot
			end
		end
	end
	return #names > 0 and string.format(L.gearNoEnchantSlots, table.concat(names, L.sourcesListSeparator)) or nil
end

-- Stat weights -----------------------------------------------------------------------------------

--- The stat weights (spec 4.3 item 2): { sim (the weights are the companion's sim), unit, rows = { stat,
--- word, weight (over the top stat's, so the top is 1.00), dps (per point, nil in score) }, haste (DPS
--- per 1% of haste, nil without one) }, the top GearTab.WEIGHT_ROWS by weight. A rating stat's weight is
--- per 1% (Rules.RATING_PER_PERCENT), so it is read per rating point first. nil without weights. Pure.
function GearTab.weights(current)
	if type(current) ~= "table" or type(current.weights) ~= "table" then
		return nil
	end
	local perPoint = tonumber(current.dpsPerPoint)
	local rows, top, haste = {}, 0, nil
	for stat, weight in pairs(current.weights) do
		if type(weight) == "number" and weight > 0 then
			if HASTE[stat] then
				haste = (haste or 0) + weight
			elseif stat ~= "weapon_dps" and stat ~= "haste" then
				local point = weight / (Rules.RATING_PER_PERCENT[stat] or 1)
				rows[#rows + 1] = { stat = stat, word = ItemTooltip.statWord(stat), point = point }
				top = math.max(top, point)
			end
		end
	end
	local kept = {}
	for _, row in ipairs(rows) do
		row.weight = row.point / top
		row.dps = perPoint ~= nil and row.point * perPoint or nil
		if row.weight >= GearTab.WEIGHT_FLOOR then
			kept[#kept + 1] = row
		end
	end
	rows = kept
	table.sort(rows, function(a, b)
		if a.weight ~= b.weight then
			return a.weight > b.weight
		end
		return a.word < b.word
	end)
	for index = #rows, GearTab.WEIGHT_ROWS + 1, -1 do
		rows[index] = nil
	end
	return { sim = current.tag == "sim", unit = ItemTooltip.unitOf(current), rows = rows,
		haste = haste ~= nil and perPoint ~= nil and haste * perPoint or nil }
end

-- Sources ----------------------------------------------------------------------------------------

--- Each kind's hover: the best item it gives each slot at the band ("Chest: Blackened Defias Armor").
local function kindItems(published, names)
	local items = {}
	for kind, list in pairs(Sources.kindItems(published)) do
		items[kind] = {}
		for _, entry in ipairs(list) do
			local name = Compat.itemName(entry.itemId) or names[entry.itemId]
			if name ~= nil then
				items[kind][#items[kind] + 1] = string.format(L.sourcesKindItem,
					L.gearSlotLabels[entry.slot] or entry.slot, name)
			end
		end
	end
	return items
end

--- The source filter at `band` (Sources.state): { band, state, words, off (the kinds off, nil at All
--- sources), empty (the kinds on leave every slot without an option), counts (Sources.counts over the
--- band's published picks), items (each kind's hover lines), choice (Sources.choiceRows) }. nil with no
--- spec, band or faction.
function GearTab.sources(data, spec, band, faction)
	local bySpec = spec ~= nil and type(data) == "table" and type(data.bis) == "table" and data.bis[spec] or nil
	local slots = bySpec ~= nil and band ~= nil and bySpec[band] ~= nil and bySpec[band][faction] or nil
	if type(slots) ~= "table" then
		return nil
	end
	local state = Sources.state(band)
	local published, empty = {}, not state.all
	for slot in pairs(slots) do
		published[slot] = Gear.publishedPick(spec, band, faction, slot, data)
		if empty and Gear.pick(spec, band, faction, slot, data, state) ~= nil then
			empty = false
		end
	end
	local counts = Sources.counts(published)
	return { band = band, state = state, words = Sources.words(state), off = Sources.offLine(state),
		empty = empty, counts = counts, items = kindItems(published, type(data.bisNames) == "table" and data.bisNames
			or {}), choice = Sources.choiceRows(state, counts) }
end

-- What waits ---------------------------------------------------------------------------------------

--- The chip's line for a bag upgrade (spec 6): "Equip Gloves of Holy Might +11.2", the gain as the
--- tooltip rounds it. Pure.
function GearTab.equipLine(name, gain, unit)
	local shown = Progress.roundedGain(gain, unit)
	local figure = unit == "score" and string.format(L.gearGainScoreBare, shown)
		or string.format(L.gearGainBare, shown)
	return string.format(L.gearEquipLine, name, figure)
end

--- The tab's waiting item (spec 3.6): the bag item that beats what you wear the most, from the list the
--- bag arrows and paper-doll arrows are drawn from (`upgrades`: Progress.bagUpgrades' rows). Pure.
function GearTab.waitingFrom(upgrades)
	local best = type(upgrades) == "table" and upgrades[1] or nil
	if best == nil or best.link == nil then
		return {}
	end
	local name = best.link:match("%[(.-)%]") or Compat.itemName(best.itemId)
	if name == nil then
		return {}
	end
	return { { key = "bag", line = GearTab.equipLine(name, best.delta, best.unit), fact = best.link } }
end

ns.GearTab = GearTab
return GearTab
