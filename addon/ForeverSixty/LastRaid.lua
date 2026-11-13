-- addon/ForeverSixty/LastRaid.lua
-- The newest last_raid inbox message (docs/contracts/addon-v2.md, contract 3) read for the window: how
-- the player followed the rotation on each boss, the biggest loss of each, the losses summed by kind,
-- and the day. The Rotation tab draws it (hub spec 4.4); Me's Right now reads its weakest boss (4.1).
-- Pure but for LastRaid.newest, which reads the inbox.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Compat = ns.Compat or require("Compat")
local Follow = ns.Follow or require("Follow")
local Score = ns.Score or require("Score")

local LastRaid = {}

--- The Unix epoch was a Thursday: the weekday index (1 Sunday) of day 0.
local EPOCH_WEEKDAY = 5
local DAY_SECONDS = 24 * 60 * 60
local WEEK_DAYS = 7
local PERCENT = 100

--- The newest last_raid message for character `key` in `inbox`, or nil.
function LastRaid.newest(inbox, key)
	return Follow.messages(inbox, key, "last_raid")[1]
end

--- The day of a YYYY-MM-DD date by name ("Tuesday"), or nil.
function LastRaid.weekday(date)
	local seconds = Score.dateSeconds(date)
	if seconds == nil then
		return nil
	end
	return L.rotWeekdays[(math.floor(seconds / DAY_SECONDS) + EPOCH_WEEKDAY - 1) % WEEK_DAYS + 1]
end

--- A loss phrase as a row starts it: "downtime" as "Downtime".
local function lossLabel(phrase)
	return phrase:sub(1, 1):upper() .. phrase:sub(2)
end

--- A boss's biggest loss in one cell: "Downtime 31 s", "Low output" (the API counts seconds only for
--- downtime and time dead), nil on a clean fight.
local function lossCell(boss)
	if type(boss.biggest_loss) ~= "string" or boss.biggest_loss == "" then
		return nil
	end
	local seconds = tonumber(boss.loss_seconds) or 0
	if seconds > 0 then
		return string.format(L.meLossSeconds, lossLabel(boss.biggest_loss), seconds)
	end
	return lossLabel(boss.biggest_loss)
end

--- The losses summed by kind in seconds, most first ("Downtime", "73 s"): only the kinds the API counts
--- seconds for (downtime, time dead); the others stay on their boss rows.
local function lostByKind(bosses)
	local kinds, order = {}, {}
	for _, boss in ipairs(bosses) do
		local phrase = type(boss) == "table" and boss.biggest_loss or nil
		if type(phrase) == "string" and phrase ~= "" then
			if kinds[phrase] == nil then
				kinds[phrase] = { label = lossLabel(phrase), seconds = 0, bosses = 0 }
				order[#order + 1] = kinds[phrase]
			end
			kinds[phrase].seconds = kinds[phrase].seconds + (tonumber(boss.loss_seconds) or 0)
			kinds[phrase].bosses = kinds[phrase].bosses + 1
		end
	end
	table.sort(order, function(a, b)
		if a.seconds ~= b.seconds then
			return a.seconds > b.seconds
		end
		if a.bosses ~= b.bosses then
			return a.bosses > b.bosses
		end
		return a.label < b.label
	end)
	local timed = {}
	for _, kind in ipairs(order) do
		if kind.seconds > 0 then
			kind.value = string.format(L.rotSeconds, kind.seconds)
			timed[#timed + 1] = kind
		end
	end
	return timed
end

--- A boss's hover: its name, followed, its parse, the loss, deaths and consumables missed.
local function bossHover(boss, followed)
	local lines = { boss.name }
	if followed ~= nil then
		lines[#lines + 1] = string.format(L.rotHoverFollowed, followed)
	end
	if type(boss.percentile) == "number" then
		lines[#lines + 1] = string.format(L.rotHoverParse, math.floor(boss.percentile + 0.5))
	end
	local loss = lossCell(boss)
	if loss ~= nil then
		lines[#lines + 1] = string.format(L.rotHoverLoss, loss)
	end
	local deaths = tonumber(boss.deaths) or 0
	if deaths > 0 then
		lines[#lines + 1] = string.format(deaths == 1 and L.rotHoverDied or L.rotHoverDiedMany, deaths)
	end
	if type(boss.consumables_missed) == "table" and #boss.consumables_missed > 0 then
		lines[#lines + 1] = string.format(L.rotHoverMissed, table.concat(boss.consumables_missed, ", "))
	end
	return lines
end

--- A boss's name as a table cell: the name before its epithet ("Golemagg the Incinerator" as
--- "Golemagg"), as the raid frames print it; the hover keeps the whole name.
function LastRaid.shortName(name)
	return name:match("^(.-) the %u") or name
end

local function percentOf(fraction)
	return math.floor(fraction * PERCENT + 0.5)
end

--- Each boss as a row { name, followed (whole percent, or nil), loss, hover }, least followed first, a
--- boss with no adherence last; and the mean of those that have one, and how many do.
local function bossRows(message)
	local rows, total, counted = {}, 0, 0
	for _, boss in ipairs(message.bosses) do
		if type(boss) == "table" and type(boss.name) == "string" then
			local adherence = tonumber(boss.adherence)
			local followed = adherence ~= nil and percentOf(adherence) or nil
			if adherence ~= nil then
				total, counted = total + adherence, counted + 1
			end
			rows[#rows + 1] = { name = boss.name, short = LastRaid.shortName(boss.name),
				followed = followed, adherence = adherence,
				loss = lossCell(boss), hover = bossHover(boss, followed) }
		end
	end
	table.sort(rows, function(a, b)
		if (a.adherence == nil) ~= (b.adherence == nil) then
			return a.adherence ~= nil
		end
		return (a.adherence or 0) < (b.adherence or 0)
	end)
	return rows, counted > 0 and percentOf(total / counted) or nil, counted
end

--- The night read for the window, or nil without a message or a boss: { note ("Tue 15 Dec"), day
--- ("Tuesday"), mean, counted, bosses (at most `limit`), more (the rest's count), lost }.
function LastRaid.played(message, limit)
	if type(message) ~= "table" or type(message.bosses) ~= "table" then
		return nil
	end
	local rows, mean, counted = bossRows(message)
	if #rows == 0 then
		return nil
	end
	local day, date = LastRaid.weekday(message.date), Compat.shortDate(message.date)
	local shown = {}
	for index = 1, math.min(#rows, limit or #rows) do
		shown[index] = rows[index]
	end
	return { note = day ~= nil and date ~= nil and (day:sub(1, 3) .. " " .. date) or date, day = day,
		mean = mean, counted = counted, bosses = shown, more = #rows - #shown, lost = lostByKind(message.bosses) }
end

--- The boss the player followed the rotation least on: { name, followed, loss, hover }, or nil.
function LastRaid.weakest(message)
	local played = LastRaid.played(message, 1)
	local first = played ~= nil and played.bosses[1] or nil
	return first ~= nil and first.followed ~= nil and first or nil
end

ns.LastRaid = LastRaid
return LastRaid
