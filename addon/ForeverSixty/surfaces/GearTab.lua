-- addon/ForeverSixty/surfaces/GearTab.lua
-- The Gear tab's model (design/specs/2026-11-13-addon-hub.md 4.3): what you wear, what beats it and where
-- you get it. The doll by band (each slot's worn item against the band's pick: Best in slot, the gain,
-- in your bags, No enchant), the upgrades with their DPS gains and total, the hit cap under the stat
-- weights (StatWeights draws the weights themselves, the Rotation tab's block), and the source filter
-- with its counts. Pure over the inputs GearTab.read gathers once a
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
--- The Conditions control's choices in its order (the web's): the list ranked under raid buffs and
--- consumables (Data.bisRaid), then the bare character's (Data.bis). Raid-ready where the band has it.
GearTab.RAID = "raid"
GearTab.BARE = "bare"
GearTab.CONDITIONS = { GearTab.RAID, GearTab.BARE }
--- The doll's two columns, top to bottom, as the character pane draws them: the off hand follows the
--- main hand only beside a one-hander.
GearTab.LEFT = { "head", "neck", "shoulder", "back", "chest", "wrist", "main_hand", "off_hand", "ranged" }
GearTab.RIGHT = { "hands", "waist", "legs", "feet", "finger1", "finger2", "trinket1", "trinket2" }
--- The most Upgrades rows.
GearTab.UPGRADES = 4

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

-- Conditions -----------------------------------------------------------------------------------

--- Whether the data carries a Raid-ready list for `band` on the character's faction (Data.bisRaid):
--- the Conditions control shows only then.
function GearTab.hasRaidList(input, band)
	local raid = type(input.data) == "table" and type(input.data.bisRaid) == "table"
		and input.spec ~= nil and input.data.bisRaid[input.spec] or nil
	local byBand = type(raid) == "table" and band ~= nil and raid[band] or nil
	return type(byBand) == "table" and type(byBand[input.faction]) == "table"
end

--- The conditions the doll shows at `band`: `chosen` while the band has a Raid-ready list, Raid-ready
--- when nothing is chosen there, and the bare character at a band without one.
function GearTab.conditionAt(input, band, chosen)
	if not GearTab.hasRaidList(input, band) then
		return GearTab.BARE
	end
	return chosen == GearTab.BARE and GearTab.BARE or GearTab.RAID
end

local function copyOf(source)
	local copy = {}
	for key, value in pairs(source) do
		copy[key] = value
	end
	return copy
end

--- The tab's input as `condition` reads it at `band`: `input` itself for the bare character, else a copy
--- whose data names the Raid-ready list as the spec's list at that band, so every pick, gain and source
--- count below reads it through the one path (Gear.pick, Progress.slotTargets). Never changes `input` or
--- its data.
function GearTab.withCondition(input, band, condition)
	if condition ~= GearTab.RAID or not GearTab.hasRaidList(input, band) then
		return input
	end
	local bySpec = copyOf(type(input.data.bis) == "table" and input.data.bis[input.spec] or {})
	bySpec[band] = input.data.bisRaid[input.spec][band]
	local bis = copyOf(type(input.data.bis) == "table" and input.data.bis or {})
	bis[input.spec] = bySpec
	local data = copyOf(input.data)
	data.bis = bis
	local shown = copyOf(input)
	shown.data = data
	return shown
end

--- The Conditions row's hover for `condition`: the raid preset's label and its buffs, debuffs and
--- consumables (Data.bisConditions), or the bare character's line. nil without the data. Pure.
function GearTab.conditionsTip(data, spec, condition)
	if condition == GearTab.BARE then
		return { L.gearBare, { text = L.gearBareTip, wrap = true } }
	end
	local entry = type(data) == "table" and type(data.bisConditions) == "table" and data.bisConditions[spec] or nil
	if type(entry) ~= "table" then
		return nil
	end
	local lines = { entry.label or L.gearRaidReady }
	for _, group in ipairs({ { L.gearBuffs, entry.buffs }, { L.gearDebuffs, entry.debuffs },
		{ L.gearConsumes, entry.consumes } }) do
		if type(group[2]) == "string" and group[2] ~= "" then
			lines[#lines + 1] = " "
			lines[#lines + 1] = Theme.inlineColor("tooltipMuted", group[1])
			lines[#lines + 1] = { text = group[2], wrap = true }
		end
	end
	return lines
end

-- The hit cap ------------------------------------------------------------------------------------

--- A percent as the client's hit line prints it: whole bare ("9"), else one decimal ("6.5").
local function capPercent(value)
	return value == math.floor(value) and string.format("%d", value) or string.format("%.1f", value)
end

--- One part's line: "Hit: 2% to the 9% cap", or "Hit: at the 9% cap" with nothing left to gain.
local function capLine(word, part)
	if part[1] <= 0 then
		return string.format(L.gearHitCapReached, word, capPercent(part[2]))
	end
	return string.format(L.gearHitCapTo, word, capPercent(part[1]), capPercent(part[2]))
end

--- The word a weapon user's hit line leads with, from the spec's hit type (Rules.hitType): "Ranged hit"
--- for a hunter shooting from range, "Melee hit" for everyone who swings.
local function weaponHitWord(spec)
	return Rules.hitType(spec) == "ranged" and L.gearHitRangedWord or L.gearHitMeleeWord
end

--- The hit cap under the stat weights (spec 13): { lines, tip }, one line per part the band publishes
--- (Data.hitToCap[spec][condition][band][faction]): special attacks or ranged shots against the 9% cap,
--- named by the spec's hit type, white swings for a dual wielder, spells against the 16% cap, and a
--- school a talent adds hit to. A ranged hunter's priority casts nothing that rolls spell hit (Arcane
--- Shot and Serpent Sting roll ranged hit), so the spell part its traps and Volley give the engine's
--- profile is left out. The figures are the list's character's (its gear and, Raid-ready, its buffs),
--- never the addon's own. nil where the band publishes none (a healer). Pure.
function GearTab.hitCap(data, spec, band, faction, condition)
	local bySpec = type(data) == "table" and type(data.hitToCap) == "table" and spec ~= nil and data.hitToCap[spec]
		or nil
	local byCondition = type(bySpec) == "table" and bySpec[condition or GearTab.BARE] or nil
	local byBand = type(byCondition) == "table" and band ~= nil and byCondition[band] or nil
	local entry = type(byBand) == "table" and byBand[faction] or nil
	if type(entry) ~= "table" then
		return nil
	end
	local spell = entry.spell
	if entry.hit ~= nil and Rules.hitType(spec) == "ranged" then
		spell = nil
	end
	local lines = {}
	for _, part in ipairs({ { weaponHitWord(spec), entry.hit }, { L.gearHitWhiteWord, entry.white },
		{ L.gearHitSpellWord, spell } }) do
		if type(part[2]) == "table" then
			lines[#lines + 1] = capLine(part[1], part[2])
		end
	end
	if type(entry.school) == "table" and type(spell) == "table" then
		lines[#lines + 1] = capLine(string.format(L.gearHitSchoolWord, entry.school[2]),
			{ entry.school[1], spell[2] })
	end
	if #lines == 0 then
		return nil
	end
	return { lines = lines, tip = { L.gearHitCapTitle, { text = L.gearHitCapTip, wrap = true } } }
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

--- The chip's line for a bag upgrade (spec 6): "Equip Gloves of Holy Might +11.2 DPS", the gain as the
--- tooltip rounds and words it (ItemTooltip.figure), in the success colour (spec 3.5: a gain). Pure.
function GearTab.equipLine(name, gain, unit)
	local figure = ItemTooltip.figure(Progress.roundedGain(gain, unit), unit)
	return string.format(L.gearEquipLine, name, Theme.inlineColor("success", figure))
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
