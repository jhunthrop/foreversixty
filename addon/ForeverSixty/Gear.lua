-- addon/ForeverSixty/Gear.lua
-- Scoring what the player is carrying against the build they loaded.
--
-- The addon never needs the site's item database: the addon code carries the
-- planned items' stats, and everything in the bags is read from the client
-- with GetItemStats. The two meet in one vocabulary -- parity contract
-- 10.8's, which is also what Data.lua's weights are keyed by -- so no
-- translation table sits between a weight and a stat. Rules.statKey owns
-- the map from the client's stat names into it.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Export = ns.Export or require("Export")
local Talents = ns.Talents or require("Talents")
local Compat = ns.Compat or require("Compat")
local Rules = ns.Rules or require("Rules")
local Follow = ns.Follow or require("Follow")

local Gear = {}

function Gear.statsOf(link)
	local raw = link and Compat.itemStats(link) or nil
	local stats = {}
	if raw == nil then
		return stats
	end
	for key, value in pairs(raw) do
		local name = Rules.statKey(key)
		-- Skip the private keys the spec mock uses to stand in for
		-- GetItemInfoInstant, and any client key this table does not know.
		if name ~= nil and type(value) == "number" then
			stats[name] = (stats[name] or 0) + value
		end
	end
	return stats
end

--- An item's stats by id, read from the client (its link, then Gear.statsOf): nil until the
--- client holds the item. The item tooltip and the gear gap price a pick through this one read.
function Gear.itemStatsOf(itemId)
	local _, link = Compat.itemInfo(itemId)
	return link ~= nil and Gear.statsOf(link) or nil
end

--- Whether an item link carries an enchant (its enchant id, the field after
--- the item id, is not 0).
function Gear.isEnchanted(link)
	local enchant = type(link) == "string" and link:match("item:%-?%d+:(%d*)") or nil
	return tonumber(enchant or "") ~= nil and tonumber(enchant) ~= 0
end

--- A Lua pattern for a client format string with one %s, anchored, the %s
--- captured ("Enchanted: %s" -> "^Enchanted: (.+)$").
local function patternOf(format)
	local escaped = format:gsub("[%^%$%(%)%.%[%]%*%+%-%?]", "%%%0")
	return "^" .. escaped:gsub("%%s", "(.+)") .. "$"
end

local function trim(text)
	return (text:gsub("^%s+", ""):gsub("%s+$", ""))
end

--- Localized stat name (lower case) -> the stat key Rules.STAT_KEYS maps it
--- to, read from the client's own ITEM_MOD_*_SHORT globals (`globals`, _G by
--- default), so an enchant line reads in every locale.
function Gear.statNames(globals)
	globals = globals or _G
	local names = {}
	for clientKey, stat in pairs(Rules.STAT_KEYS) do
		local text = globals[clientKey]
		if type(text) == "string" and text ~= "" then
			names[text:lower()] = stat
		end
	end
	return names
end

--- What the enchant reader needs from the client: the stat names, the
--- "Enchanted: %s" line (ENCHANTED_TOOLTIP_LINE), the weapon-speed word
--- (SPEED) and the weapon-damage enchant wording. The client has no global
--- for that last one, so it is the English "Weapon Damage" (Locale), the
--- documented fallback.
function Gear.enchantWords(globals)
	globals = globals or _G
	return {
		names = Gear.statNames(globals),
		enchanted = type(globals.ENCHANTED_TOOLTIP_LINE) == "string" and globals.ENCHANTED_TOOLTIP_LINE
			or L.enchantedLine,
		speed = type(globals.SPEED) == "string" and globals.SPEED or L.enchantSpeedWord,
		weaponDamage = L.enchantWeaponDamage,
	}
end

--- The stats a worn item's enchant adds, from that item's own tooltip rows
--- ({ left, right } each, the client's text): the "Enchanted: ..." line split
--- on " and " and ", ", each part "+N <stat>" or "<stat> +N" with <stat> a
--- localized stat name, "Weapon Damage +N" read as N / the weapon's speed
--- (its "Speed 3.50" right text) of weapon_dps. Answers { stats, unparsed }
--- (unparsed when some part could not be read), or nil with no enchant line.
--- Pure over `rows` and `words` (Gear.enchantWords).
function Gear.enchantOf(rows, words)
	local pattern = patternOf(words.enchanted)
	local speedPattern = "^" .. words.speed:gsub("[%^%$%(%)%.%[%]%*%+%-%?]", "%%%0") .. " ([%d%.,]+)$"
	local payload, speed
	for _, row in ipairs(rows or {}) do
		if payload == nil and type(row.left) == "string" then
			payload = row.left:match(pattern)
		end
		if type(row.right) == "string" then
			local found = row.right:match(speedPattern)
			if found ~= nil then
				speed = tonumber((found:gsub(",", ".")))
			end
		end
	end
	if payload == nil then
		return nil
	end
	local stats, unparsed = {}, false
	for part in (payload:gsub(" and ", ",") .. ","):gmatch("([^,]*),") do
		part = trim(part)
		if part ~= "" then
			local amount, name = part:match("^%+(%d+%.?%d*) (.+)$")
			if amount == nil then
				name, amount = part:match("^(.-) %+(%d+%.?%d*)$")
			end
			amount = tonumber(amount)
			local stat = name ~= nil and words.names[name:lower()] or nil
			if amount ~= nil and name ~= nil and name:lower() == words.weaponDamage:lower() then
				if speed ~= nil and speed > 0 then
					stats.weapon_dps = (stats.weapon_dps or 0) + amount / speed
				else
					unparsed = true
				end
			elseif amount ~= nil and stat ~= nil then
				stats[stat] = (stats[stat] or 0) + amount
			else
				unparsed = true
			end
		end
	end
	return { stats = stats, unparsed = unparsed }
end

--- Worn enchants read so far, by item link (an enchant only changes with the
--- link), so the client's tooltip is read once per enchanted item; cleared
--- when it grows past a bag's worth of links.
Gear.enchantCache = {}
local ENCHANT_CACHE_LIMIT = 40
local enchantCacheSize = 0

--- The enchant on the item worn in client slot `slotId` (Gear.enchantOf of
--- its tooltip rows, Compat.inventoryTooltipRows), cached by link once the
--- enchant line was read (parsed, or parsed in part); rows that cannot be read,
--- or that carry no "Enchanted:" line yet, are asked for again next time.
function Gear.wornEnchant(slotId, link)
	local cached = Gear.enchantCache[link]
	if cached ~= nil then
		return cached
	end
	local rows = Compat.inventoryTooltipRows(slotId)
	if rows == nil or #rows == 0 then
		-- Nothing read (no client data yet): ask again next time, never cache it.
		return nil
	end
	local enchant = Gear.enchantOf(rows, Gear.enchantWords())
	if enchant == nil then
		-- The link carries an enchant (the only kind asked for here), so rows with no
		-- "Enchanted:" line are not loaded yet rather than a confirmed "none": not cached.
		return nil
	end
	if enchantCacheSize >= ENCHANT_CACHE_LIMIT then
		Gear.enchantCache, enchantCacheSize = {}, 0
	end
	Gear.enchantCache[link] = enchant
	enchantCacheSize = enchantCacheSize + 1
	return enchant
end

--- `worn` (a Gear.readEquipped entry) with its enchant (Gear.enchantOf's answer, or nil)
--- added to its stats, and `enchantUnparsed` set when part of it could not be read. Pure.
function Gear.netEnchant(worn, enchant)
	if enchant == nil then
		return worn
	end
	worn.stats = worn.stats or {}
	for stat, amount in pairs(enchant.stats or {}) do
		worn.stats[stat] = (worn.stats[stat] or 0) + amount
	end
	worn.enchantUnparsed = enchant.unparsed == true
	return worn
end

--- What the player is wearing, by the site's slot name: { itemId, link, equipLoc, stats,
--- enchantUnparsed } (equipLoc, the INVTYPE_* name, tells a two-hander in the main hand).
--- A worn item's stats include its enchant (Gear.wornEnchant), so every comparison
--- against it -- the item tooltip's verdicts and next upgrade, the chip's bag upgrades --
--- is net of the enchant the player would lose; `enchantUnparsed` says one could not be
--- read. The one thin reader of equipped items; everything that scores them takes the
--- result as a table.
function Gear.readEquipped()
	local equipped = {}
	if type(GetInventoryItemLink) ~= "function" then
		return equipped
	end
	for _, entry in ipairs(Export.INVENTORY_SLOTS) do
		local link = GetInventoryItemLink("player", entry.id)
		if link ~= nil then
			local id, _, _, equipLoc = Compat.itemInfoInstant(link)
			local worn = {
				itemId = tonumber(id),
				link = link,
				equipLoc = equipLoc,
				stats = Gear.statsOf(link),
			}
			if Gear.isEnchanted(link) then
				local enchant = Gear.wornEnchant(entry.id, link)
				-- An enchant the client's rows could not give (not loaded yet) is not
				-- counted, and says so: "(enchant not counted)".
				Gear.netEnchant(worn, enchant or { stats = {}, unparsed = true })
			end
			equipped[entry.slot] = worn
		end
	end
	return equipped
end

--- Another player's equipment, as Gear.readEquipped shapes the player's own, from the links an
--- inspect read (`items`: slot -> link; the group panel, L5.2). The worn enchant is not netted
--- in: its tooltip rows can be read only for the player's own items.
function Gear.equippedFrom(items)
	local equipped = {}
	for slot, link in pairs(type(items) == "table" and items or {}) do
		if type(link) == "string" then
			local id, _, _, equipLoc = Compat.itemInfoInstant(link)
			equipped[slot] = { itemId = tonumber(id), link = link, equipLoc = equipLoc, stats = Gear.statsOf(link) }
		end
	end
	return equipped
end

--- The weighted sum of an item's stats. A weapon's own DPS (`weapon_dps`) is never
--- summed here, whatever the weights carry: Progress.itemScore alone counts it, at
--- the band's measured weapon term (Gear.weaponTerm) and only on the slot it was
--- measured on.
function Gear.score(stats, weights)
	local total = 0
	for name, value in pairs(stats or {}) do
		if name ~= "weapon_dps" then
			total = total + value * ((weights or {})[name] or 0)
		end
	end
	return total
end

--- Whether the data names `slug` as a spec: it has a default build or weights for it.
local function namedSpec(data, slug)
	return (type(data.defaultBuilds) == "table" and data.defaultBuilds[slug] ~= nil)
		or (type(data.weights) == "table" and data.weights[slug] ~= nil)
end

--- How many of a default build's planned ranks the character already has: per cell, the lesser
--- of the ranks taken and the ranks the build's order plans (the leveling and the raid order
--- both), summed. `ranks` is { [tab] = { ["tier:column"] = rank } }.
local function overlap(builds, ranks)
	local total = 0
	for _, kind in ipairs({ "leveling", "raid" }) do
		local planned = {}
		for _, cell in ipairs(type(builds[kind]) == "table" and builds[kind].order or {}) do
			local key = tostring(cell[1]) .. ":" .. tostring(cell[2]) .. ":" .. tostring(cell[3])
			planned[key] = (planned[key] or 0) + 1
			local taken = type((ranks or {})[cell[1]]) == "table" and ranks[cell[1]][cell[2] .. ":" .. cell[3]] or 0
			if type(taken) == "number" and taken >= planned[key] then
				total = total + 1
			end
		end
	end
	return total
end

--- A tree the data splits into several specs ("druid-feral" and "druid-feral-bear" both live in
--- the Feral Combat tree, whose own slug "druid-feral-combat" names neither): the one of those
--- whose default builds the character's ranks fit best, the shortest name on a tie (the cat).
--- `slug` itself when the data names it, or names nothing in that tree.
local function treeSpec(data, classSlug, slug, ranks)
	if namedSpec(data, slug) then
		return slug
	end
	local word = slug:sub(#classSlug + 2):match("^[^-]+")
	local family = classSlug .. "-" .. (word or "")
	local candidates = {}
	for _, source in ipairs({ data.defaultBuilds, data.weights }) do
		for name in pairs(type(source) == "table" and source or {}) do
			if name == family or name:sub(1, #family + 1) == family .. "-" then
				candidates[name] = true
			end
		end
	end
	local best, bestScore
	for name in pairs(candidates) do
		local builds = type(data.defaultBuilds) == "table" and data.defaultBuilds[name] or {}
		local score = overlap(builds, ranks)
		local better = best == nil or score > bestScore
		if not better and score == bestScore then
			better = #name < #best or (#name == #best and name < best)
		end
		if better then
			best, bestScore = name, score
		end
	end
	return best or slug
end

--- The spec key: the tree with the most points, ties to the first tree (the one a character
--- with no points spent is in). Design "Stat weights" states exactly this rule; the site's own
--- spec derivation uses it too, so a build reads the same in both places. A tree the data
--- splits into two specs answers the one the ranks fit (L3.7 item 3).
function Gear.specOf(data, classSlug, ranks)
	local class = data.classes[classSlug]
	if class == nil then
		return nil
	end
	local best, bestTab = -1, 1
	for tab = 1, #class.tabs do
		local total = 0
		for _, rank in pairs(ranks[tab] or {}) do
			total = total + rank
		end
		if total > best then
			best, bestTab = total, tab
		end
	end
	-- The spec slug is the tab name lowercased, which is how
	-- curated/specs.json's spec_slug is built from the tree name.
	return treeSpec(data, classSlug, classSlug .. "-" .. Export.slugify(class.tabs[bestTab].name), ranks)
end

--- Decoded BiS picks, keyed by their packed string. Weak values: a pick
--- nothing holds any more is dropped at the next collection rather than
--- kept for the session, so hovering every slot never grows idle memory.
local decodedPicks = setmetatable({}, { __mode = "v" })

--- The `|`-separated fields of a packed pick, empty ones kept in place.
local function splitFields(packed)
	local fields, start = {}, 1
	while true do
		local stop = string.find(packed, "|", start, true)
		if stop == nil then
			fields[#fields + 1] = string.sub(packed, start)
			return fields
		end
		fields[#fields + 1] = string.sub(packed, start, stop - 1)
		start = stop + 1
	end
end

local function nonEmpty(value)
	if value == nil or value == "" then
		return nil
	end
	return value
end

--- One Data.lua BiS pick, packed by data/pipeline/addonlua.py as one string
--- ("item_id|kind|source|chance|alt1_id|alt1_kind|alt1_source|alt2_id|...",
--- trailing empty fields trimmed), decoded into contract 1's logical table:
--- { item_id, kind, source, chance, alternatives = { { item_id, kind,
--- source }, ... } }. source and chance are nil when empty, and alternatives
--- is always a list (empty when there are none). Positions 1-3 also hold id,
--- kind and source, the pick tuple's old positional meaning. The table is
--- shared between callers (memoised by string), so it must not be changed.
--- nil for anything that is not a packed pick.
function Gear.decodePick(packed)
	if type(packed) ~= "string" or packed == "" then
		return nil
	end
	local cached = decodedPicks[packed]
	if cached ~= nil then
		return cached
	end
	local fields = splitFields(packed)
	local itemId = tonumber(fields[1])
	if itemId == nil then
		return nil
	end
	local source = nonEmpty(fields[3])
	local pick = {
		itemId,
		fields[2],
		source,
		item_id = itemId,
		kind = fields[2],
		source = source,
		chance = tonumber(fields[4] or ""),
		alternatives = {},
	}
	local index = 5
	while nonEmpty(fields[index]) ~= nil do
		pick.alternatives[#pick.alternatives + 1] = {
			item_id = tonumber(fields[index]),
			kind = fields[index + 1],
			source = nonEmpty(fields[index + 2]),
		}
		index = index + 3
	end
	decodedPicks[packed] = pick
	return pick
end

--- The decoded BiS pick for one spec, level band, faction ("alliance" |
--- "horde") and slot, or nil when the data names none. `data` defaults to
--- ns.Data.
function Gear.pick(spec, band, faction, slot, data)
	data = data or ns.Data
	local bySpec = data ~= nil and type(data.bis) == "table" and data.bis[spec] or nil
	local byBand = bySpec ~= nil and bySpec[band] or nil
	local byFaction = byBand ~= nil and byBand[faction] or nil
	return Gear.decodePick(byFaction ~= nil and byFaction[slot] or nil)
end

--- Below this level a character has no BiS band yet (design: "below 20
--- shows nothing" -- pipeline.addonbis.BIS_LEVEL_BANDS' own floor).
Gear.BIS_MIN_LEVEL = 20

--- pipeline.addonbis.BIS_LEVEL_BANDS' own spacing: one BiS band every ten levels.
local BIS_BAND_STEP = 10

--- The highest band key at or below `level` in a band-keyed table (bands run
--- 20..60 step 10, and a spec in progress may skip one, since the nightly
--- bis run may not have finished every band yet), and that band's entry.
--- The one rule the item tooltip (picks and band weights) and the chip's
--- "here for you" line both read, so they always round down to the same
--- rung. nil below BIS_MIN_LEVEL, or with no table or level to look in.
function Gear.bandFor(specBis, level)
	if type(specBis) ~= "table" or type(level) ~= "number" or level < Gear.BIS_MIN_LEVEL then
		return nil
	end
	local rounded = math.min(60, math.floor(level / BIS_BAND_STEP) * BIS_BAND_STEP)
	for candidate = rounded, Gear.BIS_MIN_LEVEL, -BIS_BAND_STEP do
		if specBis[candidate] ~= nil then
			return candidate, specBis[candidate]
		end
	end
	return nil
end

--- Hotfix 2.0.2: the lowest band `specBis` has, for any character below it (the shipped bands
--- start at 20, so a level 1-19 character): the rung the gear surfaces point at as upcoming
--- ("Picks start at level 20"), never as a "Next upgrade". nil from that band up (Gear.bandFor
--- serves it), or with no table, level or band.
Gear.UPCOMING_MIN_LEVEL = 1
function Gear.firstBand(specBis, level)
	if type(specBis) ~= "table" or type(level) ~= "number" or level < Gear.UPCOMING_MIN_LEVEL then
		return nil
	end
	local lowest
	for band in pairs(specBis) do
		if type(band) == "number" and (lowest == nil or band < lowest) then
			lowest = band
		end
	end
	if lowest == nil or level >= lowest then
		return nil
	end
	return lowest
end

--- The sim's figures for `spec` at `level` (contract 1's dpsPerPoint): the band's
--- own { reference, dps, weaponDps }, read with the chip's band rule
--- (Gear.bandFor over the spec's measured bands, so 20 serves 20-29), band 20 below
--- level 20 (and with no level); the spec-level figure (the level-60 raid
--- character's) only at the cap or with no level, when no band entry is there: a leveler
--- (below 60) with no band to read gets nothing, never the level-60 figure (L3.7 item 8).
--- Second return: the band read, nil for the spec-level figure. nil for a spec with no
--- entry (tanks).
function Gear.dpsFigure(data, spec, level)
	local perSpec = spec ~= nil and type(data) == "table" and type(data.dpsPerPoint) == "table"
		and data.dpsPerPoint[spec] or nil
	if type(perSpec) ~= "table" then
		return nil
	end
	local floor = math.max(type(level) == "number" and level or Gear.BIS_MIN_LEVEL, Gear.BIS_MIN_LEVEL)
	local band, entry = Gear.bandFor(perSpec, floor)
	if type(entry) == "table" then
		return entry, band
	end
	if type(level) == "number" and level < 60 then
		return nil
	end
	return perSpec, nil
end

--- The weapon term a figure (Gear.dpsFigure) measures, as { slot, weight }: the
--- slot it was measured on (the ranged slot for a ranged-attack-power spec, else
--- the main hand; no other slot has one) and what one more point of that weapon's
--- own DPS is worth in reference-stat points, weaponDps / `dps`, so that a gain
--- counted in DPS (x `dps`) prices a weapon-DPS delta at exactly weaponDps.
--- `dps` is the DPS per point the gain is counted in (a sim's own wins); with none
--- (a gain in score) the figure's own. nil when the figure's weaponDps is 0 or
--- missing (feral druids, casters, healers) or there is no DPS per point to divide by.
function Gear.weaponTerm(figure, dps)
	local weaponDps = type(figure) == "table" and tonumber(figure.weaponDps) or nil
	if weaponDps == nil or weaponDps <= 0 then
		return nil
	end
	local perPoint = tonumber(dps) or tonumber(figure.dps)
	if perPoint == nil or perPoint <= 0 then
		return nil
	end
	return { slot = figure.reference == "ranged_attack_power" and "ranged" or "main_hand",
		weight = weaponDps / perPoint }
end

--- The client's INVTYPE_* -> the site's slot names an item may occupy.
Gear.SLOTS_BY_EQUIP_LOCATION = {
	INVTYPE_HEAD = { "head" },
	INVTYPE_NECK = { "neck" },
	INVTYPE_SHOULDER = { "shoulder" },
	INVTYPE_CLOAK = { "back" },
	INVTYPE_CHEST = { "chest" },
	INVTYPE_ROBE = { "chest" },
	INVTYPE_WRIST = { "wrist" },
	INVTYPE_HAND = { "hands" },
	INVTYPE_WAIST = { "waist" },
	INVTYPE_LEGS = { "legs" },
	INVTYPE_FEET = { "feet" },
	INVTYPE_FINGER = { "finger1", "finger2" },
	INVTYPE_TRINKET = { "trinket1", "trinket2" },
	INVTYPE_WEAPON = { "main_hand", "off_hand" },
	INVTYPE_WEAPONMAINHAND = { "main_hand" },
	INVTYPE_WEAPONOFFHAND = { "off_hand" },
	INVTYPE_2HWEAPON = { "main_hand" },
	INVTYPE_SHIELD = { "off_hand" },
	INVTYPE_HOLDABLE = { "off_hand" },
	INVTYPE_RANGED = { "ranged" },
	INVTYPE_RANGEDRIGHT = { "ranged" },
	INVTYPE_THROWN = { "ranged" },
	INVTYPE_RELIC = { "ranged" },
}

--- The site slots an item of this INVTYPE_* location may be worn in, a
--- list (two for rings, trinkets and one-hand weapons); empty for a
--- location no slot takes (a shirt, a bag, a reagent).
function Gear.slotsFor(equipLocation)
	return Gear.SLOTS_BY_EQUIP_LOCATION[equipLocation] or {}
end

--- A weights message's `weights` list ({ { stat, weight }, ... }) as a
--- stat -> weight table; a table already keyed by stat is copied as is.
local function weightTable(list)
	local weights = {}
	for key, entry in pairs(type(list) == "table" and list or {}) do
		if type(entry) == "table" and type(entry.stat) == "string" and type(entry.weight) == "number" then
			weights[entry.stat] = entry.weight
		elseif type(key) == "string" and type(entry) == "number" then
			weights[key] = entry
		end
	end
	return weights
end

--- The one weights source every verdict reads (spec section 5.2):
--- { weights, dpsPerPoint, weaponDps, weapon, tag, caps, healing, spec, reference }.
--- With a "weights" inbox message for `characterKey` and `spec` (the companion's
--- sim), its weights, its own DPS per point (none: the unit is score) and its
--- caps, tag "sim"; otherwise the shipped data.weights[spec] and the band's DPS
--- per point at `level` (Gear.dpsFigure: the spec-level figure only when no band
--- entry is there), tag "default", no caps. `weaponDps` is the band's measured
--- DPS per +1 weapon DPS and `weapon` its term for Progress (Gear.weaponTerm), in
--- the unit the gain counts in, so a sim's own DPS per point wins over the band
--- and the spec figure while a weapon-DPS delta is still worth weaponDps. A
--- message with no usable weights falls back to the shipped ones but keeps its
--- caps. nil for a spec with no weights at all. `healing` is the shipped spec's
--- own flag (HPS, not DPS). DPS per point only ever multiplies the static weights
--- or the message's own weights, never a band's measured bis_weights. Pure.
--- ("build" is the tag a build carrying its own weights will take; none
--- does yet.)
function Gear.currentWeights(data, spec, inbox, characterKey, level)
	if data == nil or spec == nil then
		return nil
	end
	local shipped = type(data.weights) == "table" and data.weights[spec] or nil
	local perSpec = type(data.dpsPerPoint) == "table" and data.dpsPerPoint[spec] or nil
	local healing = type(perSpec) == "table" and perSpec.healing == true
	local perPoint = Gear.dpsFigure(data, spec, level)
	local weaponDps = perPoint ~= nil and tonumber(perPoint.weaponDps) or nil
	local message
	for _, candidate in ipairs(Follow.messages(inbox, characterKey, "weights")) do
		if candidate.spec == spec then
			message = candidate
			break
		end
	end
	local caps = message ~= nil and type(message.caps) == "table" and message.caps or {}
	local simWeights = message ~= nil and weightTable(message.weights) or nil
	if simWeights ~= nil and next(simWeights) ~= nil then
		return {
			weights = simWeights,
			-- Only the message's own number multiplies the message's own weights
			-- (L1.3 carry-over): a sim that names none counts in score.
			dpsPerPoint = message.dps_per_point,
			weaponDps = weaponDps, weapon = Gear.weaponTerm(perPoint, message.dps_per_point),
			tag = "sim", caps = caps, healing = healing, spec = spec,
			-- When the sim ran (contract 3, once the API sends it): the build dialog names its copy.
			generatedAt = message.generated_at,
			reference = perPoint ~= nil and perPoint.reference or Gear.referenceStat(simWeights),
		}
	end
	if shipped == nil then
		return nil
	end
	return {
		weights = shipped,
		dpsPerPoint = perPoint ~= nil and perPoint.dps or nil,
		weaponDps = weaponDps, weapon = Gear.weaponTerm(perPoint),
		tag = "default", caps = caps, healing = healing, spec = spec,
		reference = perPoint ~= nil and perPoint.reference or Gear.referenceStat(shipped),
	}
end

local function plannedBySlot(build)
	local planned = {}
	for _, entry in ipairs(build.gear or {}) do
		planned[entry.slot] = entry
	end
	return planned
end

local function hasAnyStat(stats)
	return next(stats or {}) ~= nil
end

--- Every item the player can reach, equipped or bagged, with its own link.
--- The design scores both: an item already on the character can still beat the
--- planned one, and hiding that would tell a player to swap something they
--- should keep.
function Gear.candidates()
	local links = {}
	for _, entry in ipairs(Export.INVENTORY_SLOTS) do
		local link = GetInventoryItemLink("player", entry.id)
		if link ~= nil then
			links[#links + 1] = link
		end
	end
	for _, bag in ipairs(Export.CARRIED_BAGS) do
		for slot = 1, Export.containerSize(bag) do
			local link = Export.containerLink(bag, slot)
			if link ~= nil then
				links[#links + 1] = link
			end
		end
	end
	return links
end

--- The bags alone: what the player could put on. What is already worn
--- is scored on the Gear page's own slot rows, never listed here as an
--- "upgrade in your bags" with an Equip button (seen in game: the worn
--- helm offered as a +10 upgrade over itself).
function Gear.bagLinks()
	local links = {}
	for _, bag in ipairs(Export.CARRIED_BAGS) do
		for slot = 1, Export.containerSize(bag) do
			local link = Export.containerLink(bag, slot)
			if link ~= nil then
				links[#links + 1] = link
			end
		end
	end
	return links
end

--- The name of the stat a spec's weights are counted in: the one whose
--- weight is exactly 1 (the engine's own convention), or nil when the
--- table has no such reference.
function Gear.referenceStat(weights)
	for name, weight in pairs(weights or {}) do
		if weight == 1 then
			return name
		end
	end
	return nil
end

--- The class names an item's own tooltip restricts it to ("Classes: Shaman" -> { "Shaman" }), or
--- nil when it names none. The client's own format string (ITEM_CLASSES_ALLOWED) is matched, else
--- the English form. Pure over the rows.
function Gear.classLine(rows)
	local format = Compat.globalString("ITEM_CLASSES_ALLOWED") or "Classes: %s"
	local pattern = patternOf(format)
	for _, row in ipairs(rows or {}) do
		local text = type(row.left) == "string" and row.left:gsub("|c%x%x%x%x%x%x%x%x", ""):gsub("|r", "") or nil
		local list = text ~= nil and text:match(pattern) or nil
		if list ~= nil then
			local names = {}
			for name in list:gmatch("[^,]+") do
				names[#names + 1] = trim(name)
			end
			return names
		end
	end
	return nil
end

--- Class restrictions read so far, by item id (false: the tooltip was read and names none), so the
--- client's tooltip is read once per item; cleared when it grows past a few bags' worth.
Gear.classCache = {}
local CLASS_CACHE_LIMIT = 200
local classCacheSize = 0

--- The classes an item is restricted to, read from its tooltip once the client has it: a list of
--- names, false for none, nil while the tooltip cannot be read (asked again next time).
local function classesOf(link)
	local itemId = tonumber(type(link) == "string" and link:match("item:(%d+)") or nil)
	if itemId ~= nil and Gear.classCache[itemId] ~= nil then
		return Gear.classCache[itemId]
	end
	local rows = Compat.hyperlinkRows(link)
	if rows == nil or #rows == 0 then
		return nil
	end
	local names = Gear.classLine(rows) or false
	if itemId ~= nil then
		if classCacheSize >= CLASS_CACHE_LIMIT then
			Gear.classCache, classCacheSize = {}, 0
		end
		Gear.classCache[itemId] = names
		classCacheSize = classCacheSize + 1
	end
	return names
end

--- Whether a restriction list names the player's class: by the player's own class name, the
--- client's male and female names for the token, or the token itself, in any case.
local function namesClass(names, classToken)
	local wanted = {}
	local function add(name)
		if type(name) == "string" and name ~= "" then
			wanted[name:lower()] = true
		end
	end
	add(classToken)
	if type(UnitClass) == "function" then
		add((UnitClass("player")))
	end
	for _, table_ in ipairs({ Compat.classNames() or {}, _G.LOCALIZED_CLASS_NAMES_FEMALE or {} }) do
		add(table_[classToken])
	end
	for _, name in ipairs(names) do
		if wanted[name:lower()] then
			return true
		end
	end
	return false
end

--- Why the character cannot put the carried item `link` on at all, or nil when it can (the
--- chip's line, the Gear page's rows and /fs loot all ask this one function): its kind and the
--- words that finish "Not for you: ...".
---   "equip"  the client says it is not equippable (no words: it is no gear);
---   "level"  "requires level 60", the item asks more than `level`;
---   "class"  "Shaman only", the item's tooltip names classes and the character's is not one.
--- The armor and weapon proficiency by level is Rules.canEquip, per slot. Anything the client has
--- not loaded yet (no level, no tooltip) is never a refusal: a wrong "not for you" is worse than
--- saying nothing. `classToken` and `level` are the player's when nil.
function Gear.restriction(link, classToken, level)
	if link == nil then
		return nil
	end
	-- An item id reads through the hyperlink the client's tooltip takes for it.
	local hyperlink = type(link) == "number" and "item:" .. link or link
	if Compat.isEquippableItem(link) == false then
		return "equip", nil
	end
	level = level or Compat.playerLevel()
	local needs = Compat.itemMinLevel(link)
	if type(level) == "number" and needs ~= nil and needs > level then
		return "level", string.format(L.gearRequiresLevel, needs)
	end
	if classToken == nil and type(UnitClass) == "function" then
		classToken = select(2, UnitClass("player"))
	end
	local names = classToken ~= nil and classesOf(hyperlink) or nil
	if names and not namesClass(names, classToken:upper()) then
		return "class", string.format(L.gearClassOnly, table.concat(names, L.gearClassOr))
	end
	return nil
end

--- Everything the player is carrying that beats the planned item in a
--- slot it fits.
function Gear.upgrades(data, build)
	local ranks = Talents.readRanks(data)
	local spec = Follow.buildSpec(data, build) or Gear.specOf(data, build.classSlug, ranks)
	local weights = spec and data.weights[spec] or nil
	if weights == nil then
		return {}
	end
	local planned = plannedBySlot(build)
	local found = {}
	local classToken = build.classSlug ~= nil and build.classSlug:upper() or nil
	local level = Compat.playerLevel()
	for _, link in ipairs(Gear.bagLinks()) do
		-- Through Compat, like Export.lua's itemIdOf and isEquippable, so
		-- every call site finds the function in the same place on a client
		-- that moved it to C_Item.
		local id, itemType, itemSubType, equipLocation, _, classId, subClassId = Compat.itemInfoInstant(link)
		-- A piece the character cannot wear (not equippable, a level or a class it is not) is
		-- skipped silently; the armor and weapon proficiency is judged per slot below.
		local unfit = Gear.restriction(link, classToken, level) ~= nil
		for _, target in ipairs(not unfit and Gear.SLOTS_BY_EQUIP_LOCATION[equipLocation] or {}) do
			local against = planned[target]
			if against ~= nil and Rules.canEquip(classToken, level, classId or itemType, subClassId or itemSubType,
				equipLocation, target) ~= "ok" then
				against = nil
			end
			-- A planned item with no stats is an FS1 code, which carries none.
			-- Scoring it at zero would call every item an upgrade; saying
			-- nothing about that slot is honest. An item that IS the planned
			-- item scores a delta of exactly zero, which the `delta > 0` check
			-- below excludes, so a correctly geared slot never lists itself.
			if against ~= nil and hasAnyStat(against.stats) then
				local delta = Gear.score(Gear.statsOf(link), weights)
					- Gear.score(against.stats, weights)
				if delta > 0 then
					found[#found + 1] = {
						slot = target,
						itemId = tonumber(id),
						link = link,
						delta = delta,
						againstItemId = against.itemId,
					}
				end
			end
		end
	end
	table.sort(found, function(left, right)
		if left.delta ~= right.delta then
			return left.delta > right.delta
		end
		return left.slot < right.slot
	end)
	return found
end

function Gear.lines(data, build)
	if build == nil then
		return { L.gearNoBuild }
	end
	local spec = Follow.buildSpec(data, build) or Gear.specOf(data, build.classSlug, Talents.readRanks(data))
	if spec == nil or data.weights[spec] == nil then
		return { L.gearNoWeights }
	end
	local upgrades = Gear.upgrades(data, build)
	if #upgrades == 0 then
		return { L.gearNone }
	end
	local lines = {}
	for index, upgrade in ipairs(upgrades) do
		lines[index] = string.format(L.gearUpgrade, upgrade.slot, upgrade.delta, upgrade.againstItemId)
	end
	return lines
end

ns.Gear = Gear
return Gear
