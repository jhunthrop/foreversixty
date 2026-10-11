-- addon/ForeverSixty/Ratings.lua
-- A thin adapter over Score.lua for the callers that predate it: the
-- player tooltip's unit line, the Overview's rating card and the Guild page.
-- It reads nothing itself; it maps the format 2 record onto the shape those
-- callers already read (rating, fights, components in order, guild progress).
-- Lane L4.4 deletes this file once those callers read Score directly.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Compat = ns.Compat or require("Compat")
local Score = ns.Score or require("Score")

local Ratings = {}

--- The newest data format this addon understands.
Ratings.FORMAT = Score.FORMAT

--- A file older than this is shown, but marked stale.
Ratings.STALE_AFTER_SECONDS = Score.STALE_AFTER_SECONDS

--- "Bob-Whitemane" is Bob on Whitemane; "Thoradin" is on this realm.
function Ratings.splitUnitName(unit)
	local name, realm = tostring(unit or ""):match("^([^-]+)-(.+)$")
	if name == nil then
		return unit, nil
	end
	return name, realm
end

--- Whether the data can be used, and why not when it cannot, with the two
--- lines the cards print: "Data from 4 Nov", and the stale warning.
function Ratings.status()
	local status = Score.status()
	if not status.available then
		return status
	end
	local when = Compat.shortDate(status.generated) or status.generated
	status.generatedLine = (when ~= nil and when ~= "") and string.format(L.ratingsGenerated, tostring(when)) or nil
	status.staleLine = status.stale and L.ratingsStale or nil
	return status
end

--- A character's card: the score as `rating`, the components that are present
--- in order, how many fights it rests on, and the record itself. Nil when there
--- is nothing to say.
function Ratings.forCharacter(name, realm)
	local record = Score.forCharacter(name, realm)
	if record == nil then
		return nil
	end
	local components = {}
	for _, component in ipairs(Score.COMPONENTS) do
		local value = type(record.components) == "table" and record.components[component.key] or nil
		if value ~= nil then
			components[#components + 1] = { key = component.key, label = L[component.label], score = value }
		end
	end
	return { rating = record.score, fights = record.fights or 0, components = components, record = record }
end

--- The guild's record, with `progress` filled in for the Guild page's headline.
function Ratings.guild(name)
	local guild = Score.guild(name)
	if guild == nil then
		return nil
	end
	local out = {}
	for key, value in pairs(guild) do
		out[key] = value
	end
	if out.progress == nil and type(out.standing) == "number" then
		out.progress = string.format(L.scoreGuildStanding, out.standing)
	end
	return out
end

--- Kept for the callers that name the realm; a guild is found by name alone.
function Ratings.forGuild(name)
	return Ratings.guild(name)
end

--- "Forever Score 1830 · Epic · as of 4 Nov".
function Ratings.tooltipLine(card)
	if card == nil then
		return nil
	end
	return Score.lines(card.record)[1]
end

ns.Ratings = Ratings
return Ratings
