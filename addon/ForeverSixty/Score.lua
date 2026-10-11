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

--- The record's components in the order a tooltip lists them, with their labels.
Score.COMPONENTS = {
	{ key = "gear", label = "scoreGear" },
	{ key = "experience", label = "scoreExperience" },
	{ key = "parses", label = "scoreParses" },
	{ key = "attendance", label = "scoreAttendance" },
	{ key = "discipline", label = "scoreDiscipline" },
}

--- The client's item-quality index for each tier name, and the tier's label.
local TIERS = {
	common = { quality = 1, label = "scoreTierCommon" },
	uncommon = { quality = 2, label = "scoreTierUncommon" },
	rare = { quality = 3, label = "scoreTierRare" },
	epic = { quality = 4, label = "scoreTierEpic" },
	legendary = { quality = 5, label = "scoreTierLegendary" },
}

-- The client has no `os` table; its clock is the `time` global. The specs run on
-- plain Lua, where only os.time exists.
local clockTime = type(time) == "function" and time or os.time

--- Seconds since the epoch at 00:00 UTC on a YYYY-MM-DD date, or nil for any
--- other text. Plain arithmetic (days from civil), so the host's time zone
--- never enters into it.
function Score.dateSeconds(text)
	local y, m, d = tostring(text or ""):match("^(%d%d%d%d)%-(%d%d)%-(%d%d)$")
	if y == nil then
		return nil
	end
	y, m, d = tonumber(y), tonumber(m), tonumber(d)
	if m < 1 or m > 12 or d < 1 or d > 31 then
		return nil
	end
	if m <= 2 then
		y = y - 1
	end
	local era = math.floor(y / 400)
	local yoe = y - era * 400
	local doy = math.floor((153 * (m + (m > 2 and -3 or 9)) + 2) / 5) + d - 1
	local doe = yoe * 365 + math.floor(yoe / 4) - math.floor(yoe / 100) + doy
	return (era * 146097 + doe - 719468) * 86400
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

local function unavailable(reason)
	noteOnce(reason)
	return { available = false, reason = reason }
end

--- Whether the record can be used, and why not when it cannot:
--- { available, generated, stale, reason }.
function Score.status()
	local data = _G.ForeverSixtyData
	if data == nil then
		return unavailable(L.ratingsNotInstalled)
	end
	if not wellFormed(data) then
		return unavailable(L.ratingsUnreadable)
	end
	if data.format > Score.FORMAT then
		return unavailable(L.ratingsUpdateAddon)
	end
	if data.format < Score.FORMAT then
		-- An older data addon keys its characters differently; reading it as this
		-- format would only find nobody.
		return unavailable(L.ratingsUnreadable)
	end
	local written = Score.dateSeconds(data.generated)
	local stale = written == nil or clockTime() - written > Score.STALE_AFTER_SECONDS
	return { available = true, generated = data.generated, stale = stale }
end

local function usable()
	local status = Score.status()
	return status.available and _G.ForeverSixtyData or nil
end

local function ownRealm()
	return type(GetRealmName) == "function" and GetRealmName() or nil
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

local function tierLabel(tier)
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
	if type(row) ~= "table" or type(row.score) ~= "number" then
		return nil
	end
	if row.score == 0 and (row.fights or 0) == 0 then
		return nil
	end
	local record = copy(row)
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
	local tier = tierLabel(record.tier)
	if tier ~= nil then
		head[#head + 1] = tier
	end
	if record.asOfText ~= nil then
		head[#head + 1] = record.asOfText
	end
	local lines = { table.concat(head, L.exportTreeSeparator) }
	local components = type(record.components) == "table" and record.components or {}
	for _, component in ipairs(Score.COMPONENTS) do
		local value = components[component.key]
		if type(value) == "number" then
			lines[#lines + 1] = string.format(L.scoreComponent, L[component.label], value)
		end
	end
	return lines
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
