-- addon/ForeverSixty/Score.lua
-- Reading the Forever Sixty Data addon's nightly record (contract 5,
-- docs/contracts/addon-v2.md, format 2).
--
-- The Forever Score is calculated by the site and written to one global,
-- ForeverSixtyData, every night. This file displays and dates it; it never
-- computes a number and never works out a tier (the record carries both).
-- The file may be absent, newer than this addon understands, days old, or
-- missing the character asked about; every one of those reads as "nothing
-- to say" with a reason a player can act on, never as an error.
--
-- Keys are "Name-Realm". `generated` and `as_of` are date-only strings
-- (YYYY-MM-DD), read as the start of that day in UTC.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Compat = ns.Compat or require("Compat")
local Theme = ns.Theme or require("Theme")

local Score = {}

--- The newest data format this addon understands.
Score.FORMAT = 2

--- A file older than this is shown, but marked stale.
Score.STALE_AFTER_SECONDS = 3 * 24 * 60 * 60

--- The one staleness rule (design/specs/2026-11-13-addon-hub.md 5): `dated` (epoch seconds, nil for
--- unknown) is older than Score.STALE_AFTER_SECONDS at `now`. Unknown is not stale. Pure.
function Score.isStale(dated, now)
	return type(dated) == "number" and type(now) == "number" and now - dated > Score.STALE_AFTER_SECONDS
end

--- Below this level "no logs yet" is not said (L3.7 item 34): no logs is the ordinary state of a
--- new character, and nothing is a better line than a reproach.
Score.LEVEL_FLOOR = 10

--- The record's components in the order a tooltip lists them, with their labels and the most points
--- each can give (the formula's ceilings, pinned to docs/contracts/fixtures/score-labels.json).
Score.COMPONENTS = {
	{ key = "gear", label = "scoreGear", max = 500 },
	{ key = "experience", label = "scoreExperience", max = 700 },
	{ key = "parses", label = "scoreParses", max = 600 },
	{ key = "attendance", label = "scoreAttendance", max = 200 },
	{ key = "discipline", label = "scoreDiscipline", max = 100 },
}

--- The client's item-quality index for each tier name, the tier's label and its lowest score (the
--- formula's tier floors, pinned to score-labels.json).
local TIERS = {
	common = { quality = 1, label = "scoreTierCommon", from = 0 },
	uncommon = { quality = 2, label = "scoreTierUncommon", from = 500 },
	rare = { quality = 3, label = "scoreTierRare", from = 1000 },
	epic = { quality = 4, label = "scoreTierEpic", from = 1500 },
	legendary = { quality = 5, label = "scoreTierLegendary", from = 2000 },
}
Score.TIERS = TIERS

--- A part's share of its own maximum on the item quality ladder (the web's partQuality), highest
--- floor first: 90 % or more is epic.
Score.PART_QUALITY = { { 0.9, "epic" }, { 0.75, "rare" }, { 0.5, "uncommon" } }

-- The client has no `os` table; its clock is the `time` global. The specs run on
-- plain Lua, where only os.time exists.
local clockTime = type(time) == "function" and time or os.time

--- Seconds since the epoch, the clock Score.dateSeconds is compared against.
function Score.clock()
	return clockTime()
end

--- Seconds since the epoch at 00:00 UTC on a YYYY-MM-DD date, or nil for any
--- other text. Plain arithmetic (days from civil), so the host's time zone
--- never enters into it.
function Score.dateSeconds(text)
	local days = Compat.epochDays(text)
	return days ~= nil and days * 86400 or nil
end

local function wellFormed(data)
	return type(data) == "table"
		and type(data.format) == "number"
		and type(data.characters) == "table"
		and type(data.guilds) == "table"
end

local noted = {}

--- Say once, in /fs diag, why the record cannot be used.
local function noteOnce(reason)
	if not noted[reason] then
		noted[reason] = true
		Theme.note(reason)
	end
end

--- `diag` is the /fs diag line, when the reason is worth one.
local function unavailable(reason, diag)
	if diag ~= nil then
		noteOnce(diag)
	end
	return { available = false, reason = reason }
end

--- Whether the record can be used, and why not when it cannot:
--- { available, absent, generated, stale, reason }. /fs diag says once which
--- cause it is: not installed (`absent`, a finding, since the data addon is
--- optional), or there but too old to read (a fault).
function Score.status()
	local data = _G.ForeverSixtyData
	if data == nil then
		-- Optional, so not a fault: one finding, outside "Not working:".
		if not noted[L.diagDataAddonMissingInfo] then
			noted[L.diagDataAddonMissingInfo] = true
			Theme.noteInfo(L.diagDataAddonMissingInfo)
		end
		return { available = false, absent = true, reason = L.ratingsNotInstalled }
	end
	if type(data) ~= "table" then
		-- Not a table at all (a hand edit, a half-written file): broken or missing, not "too old".
		local status = unavailable(L.ratingsBroken, L.diagDataAddonBroken)
		status.broken = true
		return status
	end
	-- When the record was made, when it says: the Settings foot dates "too old" by it.
	local generated = type(data.generated) == "string" and data.generated ~= "" and data.generated or nil
	if not wellFormed(data) then
		local status = unavailable(L.ratingsUnreadable, L.diagDataAddonOld)
		status.generated = generated
		return status
	end
	if data.format > Score.FORMAT then
		return unavailable(L.ratingsUpdateAddon, L.ratingsUpdateAddon)
	end
	if data.format < Score.FORMAT then
		-- An older data addon keys its characters differently; reading it as this
		-- format would only find nobody.
		local status = unavailable(L.ratingsUnreadable, L.diagDataAddonOld)
		status.generated = generated
		return status
	end
	local written = Score.dateSeconds(data.generated)
	local stale = written == nil or Score.isStale(written, clockTime())
	return { available = true, generated = data.generated, stale = stale }
end

--- The record's state in one word, for the lines that say it (Settings' foot, the guild
--- tab's notice; L5.1 fix round 1, item 1): "Missing" (not installed), "Current", "Newer"
--- (made for a newer addon), "Broken" (the global is not a table) or "Old" (stale, unreadable
--- or an older format). `status` is
--- Score.status()'s; Score.status() is asked when it is nil. Pure over a given status.
function Score.ratingsState(status)
	if status == nil then
		status = Score.status()
	end
	if type(status) ~= "table" or status.absent then
		return "Missing"
	elseif status.broken then
		return "Broken"
	elseif status.available and not status.stale then
		return "Current"
	elseif not status.available and status.reason == L.ratingsUpdateAddon then
		return "Newer"
	end
	return "Old"
end

local function usable()
	local status = Score.status()
	return status.available and _G.ForeverSixtyData or nil
end

local function ownRealm()
	return type(GetRealmName) == "function" and GetRealmName() or nil
end

--- Whether `name` (and `realm`, nil for this realm) is the logged-in character.
function Score.isOwn(name, realm)
	local own = Compat.playerName()
	if own == nil or name == nil then
		return false
	end
	local bare = name:match("^(.-)%-") or name
	return bare == own and (realm == nil or realm == ownRealm())
end


--- The client's colour for a tier name, as r, g, b, a; nil for a tier this
--- addon does not know, which shows the score uncoloured.
function Score.tierColor(tier)
	local entry = type(tier) == "string" and TIERS[tier] or nil
	if entry == nil then
		return nil
	end
	return Theme.qualityColor(entry.quality)
end

--- A tier's label ("Rare"), or nil for a tier this addon does not know.
function Score.tierLabel(tier)
	local entry = type(tier) == "string" and TIERS[tier] or nil
	return entry and L[entry.label] or nil
end

local function copy(record)
	local out = {}
	for key, value in pairs(record) do
		out[key] = value
	end
	return out
end

--- The inbox's `own_score` message (contract 3) as a record in the nightly file's shape, for the
--- player's own character only: a player who hid their score is left out of the nightly Data
--- addon but still sees it. nil without one, or for anyone else.
function Score.ownRecord(inbox, key)
	local Codec = ns.Codec or require("Codec")
	local message = Codec.inboxMessages(inbox, key, Score.OWN_KIND)[1]
	if type(message) ~= "table" or type(message.score) ~= "number" then
		return nil
	end
	local components = {}
	for _, component in ipairs(Score.COMPONENTS) do
		if type(message[component.key]) == "number" then
			components[component.key] = message[component.key]
		end
	end
	return { score = message.score, tier = message.tier, as_of = message.as_of, fights = message.fights,
		components = components }
end

--- The inbox message the player's own record travels in; Codec.inboxMessages reads only the kinds
--- it lists, so this module adds its own (it loads after Codec).
Score.OWN_KIND = "own_score"
;(ns.Codec or require("Codec")).MESSAGE_TYPES[Score.OWN_KIND] = true

--- The record for a character, plus `asOfText` ("as of 4 Nov") and `tierColor`
--- ({ r, g, b, a }, or nil for a tier it does not know). Nil for a character the
--- record does not have, and for one with no logs yet: never a zero.
--- `name` may be "Name-Realm" when no realm is given.
function Score.forCharacter(name, realm)
	local data = usable()
	if data == nil or type(name) ~= "string" or name == "" then
		return nil
	end
	local row
	if realm == nil and name:find("-", 1, true) then
		row = data.characters[name]
	else
		realm = realm or ownRealm()
		if type(realm) == "string" and realm ~= "" then
			row = data.characters[name .. "-" .. realm]
				or data.characters[name .. "-" .. realm:gsub("[%s%-]", "")]
		end
	end
	if (type(row) ~= "table" or type(row.score) ~= "number") and Score.isOwn(name, realm) then
		-- The player's own score, hidden from others, from the inbox (contract 3 `own_score`).
		local Export = ns.Export or require("Export")
		row = Score.ownRecord(_G.ForeverSixtyInbox, Export.characterKey())
	end
	if type(row) ~= "table" or type(row.score) ~= "number" then
		return nil
	end
	if row.score == 0 and (row.fights or 0) == 0 then
		return nil
	end
	local record = copy(row)
	-- Older than the window (L3.7 item 34): the badge shows muted. The row's own date, else the file's.
	local dated = Score.dateSeconds(row.as_of) or Score.dateSeconds(data.generated)
	record.stale = Score.isStale(dated, clockTime())
	local shortDate = Compat.shortDate(row.as_of) or Compat.shortDate(data.generated)
	record.asOfText = shortDate and string.format(L.scoreAsOf, shortDate) or nil
	local r, g, b, a = Score.tierColor(row.tier)
	record.tierColor = r ~= nil and { r = r, g = g, b = b, a = a } or nil
	return record
end

--- Tooltip lines for a record from forCharacter: "Forever Score 1830 · Epic · as of 4 Nov",
--- then one line per component the record carries. Empty for no record.
function Score.lines(record)
	if type(record) ~= "table" or type(record.score) ~= "number" then
		return {}
	end
	local head = { string.format(L.scoreLine, record.score) }
	local tier = Score.tierLabel(record.tier)
	if tier ~= nil then
		head[#head + 1] = tier
	end
	if record.asOfText ~= nil then
		head[#head + 1] = record.asOfText
	end
	local lines = { table.concat(head, L.exportTreeSeparator) }
	for _, part in ipairs(Score.components(record)) do
		lines[#lines + 1] = string.format(L.scoreComponent, part.label, part.value)
	end
	return lines
end

--- The components a record carries, in Score.COMPONENTS order, each { label, value }
--- ("Gear", 380): what the player tooltip draws as label and right-aligned value. Empty for
--- no record or none.
function Score.components(record)
	local parts = {}
	local components = type(record) == "table" and type(record.components) == "table" and record.components or {}
	for _, component in ipairs(Score.COMPONENTS) do
		local value = components[component.key]
		if type(value) == "number" then
			parts[#parts + 1] = { label = L[component.label], value = value }
		end
	end
	return parts
end

--- A tier's colour on the window's dark ground as six hex digits (Theme.panelTierHex over its quality
--- colour), or nil for a tier this addon does not know.
function Score.tierHex(tier)
	local entry = type(tier) == "string" and TIERS[tier] or nil
	if entry == nil then
		return nil
	end
	return Theme.panelTierHex(tier, Theme.QUALITY_HEX[entry.quality])
end

--- The tier above `score` and how far it is: { tier, label, from, gap } ("316 to Epic (1500)"), or nil
--- at the top tier (or without a number). Pure.
function Score.nextTier(score)
	if type(score) ~= "number" then
		return nil
	end
	local best
	for tier, entry in pairs(TIERS) do
		if entry.from > score and (best == nil or entry.from < best.from) then
			best = { tier = tier, label = L[entry.label], from = entry.from, gap = entry.from - score }
		end
	end
	return best
end

--- The tier name for a part's share `fill` (0 to 1) of its maximum, the web's partQuality. Pure.
function Score.partQuality(fill)
	for _, floor in ipairs(Score.PART_QUALITY) do
		if (fill or 0) >= floor[1] then
			return floor[2]
		end
	end
	return "common"
end

--- Every part on one scale (the web's scorePartRows): { key, label, points, max, track (the part's
--- maximum against the largest, 0 to 1), fill (its share of its own maximum, 0 to 1), quality }, in
--- Score.COMPONENTS order; a part the record lacks counts 0. Pure.
function Score.parts(record)
	local components = type(record) == "table" and type(record.components) == "table" and record.components or {}
	local longest = 0
	for _, component in ipairs(Score.COMPONENTS) do
		longest = math.max(longest, component.max)
	end
	local rows = {}
	for _, component in ipairs(Score.COMPONENTS) do
		local points = type(components[component.key]) == "number" and components[component.key] or 0
		local fill = math.max(0, math.min(1, points / component.max))
		rows[#rows + 1] = { key = component.key, label = L[component.label], points = points, max = component.max,
			track = component.max / longest, fill = fill, quality = Score.partQuality(fill) }
	end
	return rows
end

--- How many whole days old the ratings pack is at `now` (its `generated` date), or nil where it
--- does not say. Pure over a Score.status() result.
function Score.ageDays(status, now)
	local written = type(status) == "table" and Score.dateSeconds(status.generated) or nil
	if written == nil or type(now) ~= "number" then
		return nil
	end
	return math.max(0, math.floor((now - written) / 86400))
end

--- The guild's record by name, or nil.
function Score.guild(name)
	local data = usable()
	if data == nil or type(name) ~= "string" then
		return nil
	end
	local row = data.guilds[name]
	return type(row) == "table" and row or nil
end

ns.Score = Score
return Score
