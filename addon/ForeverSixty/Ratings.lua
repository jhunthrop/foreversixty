-- addon/ForeverSixty/Ratings.lua
-- Reading the Forever Sixty Data addon.
--
-- That addon is a nightly file on one global, ForeverSixtyData, published
-- to CurseForge and Wago every night by the site. It may be absent (the
-- player installed only this addon), newer than this addon understands,
-- days old (their addon manager did not update), or missing the character
-- asked about. Every one of those reads as "no data" with a reason a
-- player can act on, never as an error: nothing here can break the window.
--
-- Keys match the site's own: "<region>:<realm>:<name>", lowercased, spaces
-- to dashes, so the nightly job writes them and this file never guesses.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Export = ns.Export or require("Export")

local Ratings = {}

--- The newest data format this addon understands.
Ratings.FORMAT = 1

--- A file older than this is shown, but marked stale.
Ratings.STALE_AFTER_SECONDS = 3 * 24 * 60 * 60

--- Component order and their labels, the site's own.
Ratings.COMPONENTS = {
	{ key = "output", label = "ratingsOutput" },
	{ key = "survival", label = "ratingsSurvival" },
	{ key = "mechanics", label = "ratingsMechanics" },
	{ key = "utility", label = "ratingsUtility" },
	{ key = "preparation", label = "ratingsPreparation" },
	{ key = "activity", label = "ratingsActivity" },
}

local function slug(text)
	return (tostring(text or ""):lower():gsub("%s+", "-"))
end

function Ratings.characterKey(region, realm, name)
	return slug(region) .. ":" .. slug(realm) .. ":" .. slug(name)
end

function Ratings.guildKey(region, realm, name)
	return Ratings.characterKey(region, realm, name)
end

--- "Bob-Whitemane" is Bob on Whitemane; "Thoradin" is on this realm.
function Ratings.splitUnitName(unit)
	local name, realm = tostring(unit or ""):match("^([^-]+)-(.+)$")
	if name == nil then
		return unit, nil
	end
	return name, realm
end

local function playerRegion()
	return Export.REGION_NAMES[type(GetCurrentRegion) == "function" and GetCurrentRegion() or 0] or ""
end

local function playerRealm()
	return type(GetRealmName) == "function" and GetRealmName() or ""
end

--- Seconds since an ISO 8601 UTC stamp, or nil for one this cannot read.
local function ageOf(generated)
	local y, mo, d, h, mi, s = tostring(generated or ""):match("^(%d%d%d%d)%-(%d%d)%-(%d%d)T(%d%d):(%d%d):(%d%d)Z$")
	if y == nil then
		return nil
	end
	-- os.time reads the table as local time; the difference from a UTC
	-- table read the same way cancels the zone out.
	local now = os.time(os.date("!*t"))
	local then_ = os.time({ year = tonumber(y), month = tonumber(mo), day = tonumber(d),
		hour = tonumber(h), min = tonumber(mi), sec = tonumber(s) })
	return now - then_
end

local function wellFormed(data)
	return type(data) == "table"
		and type(data.format) == "number"
		and type(data.characters) == "table"
		and type(data.guilds) == "table"
end

--- Whether the data can be used, and why not when it cannot.
function Ratings.status()
	local data = _G.ForeverSixtyData
	if data == nil then
		return { available = false, reason = L.ratingsNotInstalled }
	end
	if not wellFormed(data) then
		return { available = false, reason = L.ratingsUnreadable }
	end
	if data.format > Ratings.FORMAT then
		return { available = false, reason = L.ratingsUpdateAddon }
	end
	local age = ageOf(data.generated)
	local stale = age == nil or age > Ratings.STALE_AFTER_SECONDS
	return {
		available = true,
		stale = stale,
		generated = data.generated,
		generatedLine = string.format(L.ratingsGenerated, tostring(data.generated or ""):sub(1, 10)),
		staleLine = stale and L.ratingsStale or nil,
	}
end

local function usable()
	local status = Ratings.status()
	return status.available and _G.ForeverSixtyData or nil
end

--- A character's card: the rating, the six components in order, and how
--- many fights it rests on. Nil when there is nothing to say.
function Ratings.forCharacter(name, realm)
	local data = usable()
	if data == nil or name == nil then
		return nil
	end
	local row = data.characters[Ratings.characterKey(playerRegion(), realm or playerRealm(), name)]
	if type(row) ~= "table" or type(row.rating) ~= "number" then
		return nil
	end
	local components = {}
	for index, component in ipairs(Ratings.COMPONENTS) do
		components[index] = { key = component.key, label = L[component.label], score = row[component.key] }
	end
	return { rating = row.rating, fights = row.fights or 0, components = components }
end

function Ratings.forGuild(name, realm)
	local data = usable()
	if data == nil or name == nil then
		return nil
	end
	local row = data.guilds[Ratings.guildKey(playerRegion(), realm or playerRealm(), name)]
	return type(row) == "table" and row or nil
end

function Ratings.tooltipLine(card)
	if card == nil then
		return nil
	end
	return string.format(L.ratingsTooltip, card.rating, card.fights)
end

ns.Ratings = Ratings
return Ratings
