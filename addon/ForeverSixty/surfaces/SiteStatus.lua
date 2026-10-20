-- addon/ForeverSixty/surfaces/SiteStatus.lua
-- What the site has of this character and what the inbox holds (X.3): the two /fs diag lines,
-- the headline's verdict, the one chat line at login, and the chip's "a build is waiting".
-- Reads the companion's inbox (the per-character "ack" and the builds), the character's
-- record in ForeverSixtyDB, and the clock. Writes no frame and runs no OnUpdate: the login
-- line is one timer through Theme.after.
local _, ns = ...
ns = type(ns) == "table" and ns or {}

local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Codec = ns.Codec or require("Codec")
local Follow = ns.Follow or require("Follow")
local Export = ns.Export or require("Export")
local Compat = ns.Compat or require("Compat")
local BuildDialog = ns.BuildDialog or require("BuildDialog")

local SiteStatus = {}

--- The site has seen the character "recently" within a day; the companion is "silent" past three.
SiteStatus.FRESH_SECONDS = 24 * 3600
SiteStatus.SILENT_SECONDS = 72 * 3600
--- Seconds after PLAYER_ENTERING_WORLD of the login line: with the export's first refresh.
SiteStatus.LOGIN_DELAY = 5

local function record()
	local characters = ForeverSixtyDB and ForeverSixtyDB.characters
	local mine = type(characters) == "table" and characters[Export.characterKey()] or nil
	return type(mine) == "table" and mine or nil
end

--- The newest ack for this character (the inbox is append-ordered), or nil.
local function ack()
	local ok, found = pcall(Follow.messages, ForeverSixtyInbox, Export.characterKey(), "ack")
	return ok and type(found) == "table" and found[1] or nil
end

--- How many characters the site has acknowledged: the distinct characters among the acks.
local function ackedCharacters()
	local seen, count = {}, 0
	local messages = type(ForeverSixtyInbox) == "table" and ForeverSixtyInbox.messages
	for _, message in ipairs(type(messages) == "table" and messages or {}) do
		if type(message) == "table" and message.type == "ack" then
			local who = type(message.character) == "string" and message.character:lower() or ""
			if not seen[who] then
				seen[who] = true
				count = count + 1
			end
		end
	end
	return count
end

--- The builds waiting for this character: the inbox's rows nobody has used or dismissed.
function SiteStatus.waiting(data)
	local ok, rows = pcall(function()
		return (BuildDialog.inboxRows(Follow.inboxBuilds(ForeverSixtyInbox, Export.characterKey()),
			Follow.dismissedSet(), data, BuildDialog.player()))
	end)
	return ok and type(rows) == "table" and rows or {}
end

--- How many builds wait: the chip's reading.
function SiteStatus.waitingCount(data)
	return #SiteStatus.waiting(data)
end

--- The messages (other than builds and acks) addressed to this character or to none.
local function messageCount()
	local count = 0
	for kind in pairs(Codec.MESSAGE_TYPES) do
		if kind ~= "ack" and kind ~= "build" then
			count = count + #Codec.inboxMessages(ForeverSixtyInbox, Export.characterKey(), kind)
		end
	end
	return count
end

--- Everything the lines need, read once: { seen (epoch of the newest ack or nil), age (seconds),
--- characters, gear (the record carries gear), gearAt (epoch the record first had gear), clock }.
function SiteStatus.facts(clock)
	clock = clock or BuildDialog.liveClock()
	local mine = record()
	local newest = ack()
	local seen = newest ~= nil and BuildDialog.parseUtc(newest.export_seen_at) or nil
	local facts = {
		clock = clock,
		seen = seen,
		age = seen ~= nil and clock.now - seen or nil,
		characters = ackedCharacters(),
		hasRecord = mine ~= nil,
		gear = mine ~= nil and Export.hasGear(mine.export),
	}
	if mine ~= nil and type(mine.gearAt) == "string" then
		facts.gearAt = BuildDialog.parseLocal(mine.gearAt, clock)
	end
	return facts
end

--- "12 Oct": the day of `seconds` in the clock's local time.
local function dayOf(seconds, clock)
	local at = clock.fields(seconds)
	return Compat.shortDate(string.format("%04d-%02d-%02d", at.year, at.month, at.day)) or ""
end

--- "Site: ..." for /fs diag.
function SiteStatus.siteLine(facts)
	if facts.seen == nil then
		return L.diagSiteNever
	end
	if facts.age > SiteStatus.FRESH_SECONDS then
		local days = math.floor(facts.age / 86400)
		return string.format(days == 1 and L.diagSiteOldOne or L.diagSiteOld, dayOf(facts.seen, facts.clock), days)
	end
	local stamp = BuildDialog.when(facts.seen, facts.clock)
	if facts.characters == 1 then
		return string.format(L.diagSiteSeenOne, stamp)
	end
	return string.format(L.diagSiteSeen, stamp, facts.characters)
end

local function nameList(rows)
	local names = {}
	for index, row in ipairs(rows) do
		names[index] = tostring(row.name):gsub("|", "||")
	end
	return table.concat(names, ", ")
end

--- "Inbox: ..." for /fs diag.
function SiteStatus.inboxLine(data)
	local rows, messages = SiteStatus.waiting(data), messageCount()
	if #rows == 0 and messages == 0 then
		return L.diagInboxEmpty
	end
	local words = string.format(messages == 1 and L.diagInboxMessageOne or L.diagInboxMessages, messages)
	if #rows == 0 then
		return string.format(L.diagInboxNoBuilds, words)
	elseif #rows == 1 then
		return string.format(L.diagInboxBuildOne, nameList(rows), words)
	end
	return string.format(L.diagInboxBuilds, #rows, nameList(rows), words)
end

--- Whether "Everything is working." may be said: the export has gear and the site has seen
--- this character within a day.
function SiteStatus.healthy(facts)
	return facts.gear == true and facts.seen ~= nil and facts.age <= SiteStatus.FRESH_SECONDS
end

--- The /fs diag lines after the Export line, and the headline for the all-clear.
function SiteStatus.diag(data)
	local facts = SiteStatus.facts()
	return { SiteStatus.siteLine(facts), SiteStatus.inboxLine(data) }, SiteStatus.healthy(facts)
end

--- The one chat line at login (the text after the addon's prefix), the first that applies, or nil.
function SiteStatus.loginLine(data, facts)
	facts = facts or SiteStatus.facts()
	if facts.seen == nil then
		return L.siteLoginNever
	end
	-- The site's last copy was seen before this record first carried gear.
	if facts.gear and facts.gearAt ~= nil and facts.seen < facts.gearAt then
		return L.siteLoginGear
	end
	local rows = SiteStatus.waiting(data)
	if #rows == 1 then
		return string.format(L.siteLoginBuild, nameList(rows))
	elseif #rows > 1 then
		return string.format(L.siteLoginBuilds, #rows, nameList(rows))
	end
	if facts.hasRecord and facts.age > SiteStatus.SILENT_SECONDS then
		return string.format(L.siteLoginSilent, dayOf(facts.seen, facts.clock))
	end
	return nil
end

--- PLAYER_ENTERING_WORLD: once a session, five seconds in (after the export's refresh, queued
--- first), say the one line that is true, if any.
function SiteStatus.afterEnteringWorld(data)
	if SiteStatus.announced then
		return
	end
	SiteStatus.announced = true
	Theme.after(SiteStatus.LOGIN_DELAY, function()
		local line = SiteStatus.loginLine(data)
		if line ~= nil then
			Theme.say(line)
		end
	end)
end

ns.SiteStatus = SiteStatus
return SiteStatus
