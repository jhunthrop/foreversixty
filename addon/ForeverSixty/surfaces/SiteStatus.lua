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

--- The messages (other than builds and acks) addressed to this character or to none, counted by
--- kind: { guild = 1 }.
local function messageCounts()
	local counts = {}
	for kind in pairs(Codec.MESSAGE_TYPES) do
		if kind ~= "ack" and kind ~= "build" then
			local found = #Codec.inboxMessages(ForeverSixtyInbox, Export.characterKey(), kind)
			if found > 0 then
				counts[kind] = found
			end
		end
	end
	return counts
end

--- The messages in words: "1 guild message", "2 guild messages" while one kind waits, "2 messages"
--- where kinds mix (or a kind has no word of its own).
local function messageWords(counts)
	local kinds, total = {}, 0
	for kind, found in pairs(counts) do
		kinds[#kinds + 1] = kind
		total = total + found
	end
	local words = #kinds == 1 and L.diagInboxKinds[kinds[1]] or nil
	if words ~= nil then
		return string.format("%d %s", total, total == 1 and words[1] or words[2])
	end
	return string.format(total == 1 and L.diagInboxMessageOne or L.diagInboxMessages, total)
end

--- What the addon remembers of this character's logins, in ForeverSixtyDB.site[key] so a /reload
--- keeps it: { loginAt (epoch of the last login), said (the login line said last), gearPending (the
--- "on its way" line was said and not yet settled) }. Nil when nothing is kept and `create` is not
--- set.
function SiteStatus.memory(create)
	if type(ForeverSixtyDB) ~= "table" then
		if not create then
			return nil
		end
		ForeverSixtyDB = {}
	end
	if type(ForeverSixtyDB.site) ~= "table" then
		if not create then
			return nil
		end
		ForeverSixtyDB.site = {}
	end
	local key = Export.characterKey()
	local mine = ForeverSixtyDB.site[key]
	if type(mine) ~= "table" and create then
		mine = {}
		ForeverSixtyDB.site[key] = mine
	end
	return type(mine) == "table" and mine or nil
end

--- Everything the lines need, read once: { seen (epoch of the newest ack or nil: the site's receive
--- time, whatever `export_changed_at` says), age (seconds), characters, gear (the record carries gear),
--- gearAt (epoch the record first had gear), builtAt (epoch the record was last built), loginAt (epoch
--- of the last login), gearPending, clock }.
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
	if mine ~= nil and type(mine.builtAt) == "string" then
		facts.builtAt = BuildDialog.parseLocal(mine.builtAt, clock)
	end
	local remembered = SiteStatus.memory(false)
	if remembered ~= nil then
		facts.loginAt, facts.gearPending = remembered.loginAt, remembered.gearPending == true
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

--- "Inbox: ..." for /fs diag. The inbox is read at login only, so an empty one says as of when (the
--- client clock at login, in the chat's own format) and that /reload does not read it again.
function SiteStatus.inboxLine(data, facts)
	local rows, counts = SiteStatus.waiting(data), messageCounts()
	if #rows == 0 and next(counts) == nil then
		facts = facts or SiteStatus.facts()
		if facts.loginAt == nil then
			return L.diagInboxEmptyNoLogin
		end
		return string.format(L.diagInboxEmpty, BuildDialog.when(facts.loginAt, facts.clock))
	end
	-- With no message to count, the builds alone (never "0 messages").
	if next(counts) == nil then
		if #rows == 1 then
			return string.format(L.diagInboxBuildAlone, nameList(rows))
		end
		return string.format(L.diagInboxBuildsAlone, #rows, nameList(rows))
	end
	local words = messageWords(counts)
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
	return { SiteStatus.siteLine(facts), SiteStatus.inboxLine(data, facts) }, SiteStatus.healthy(facts)
end

--- The one chat line at login (the text after the addon's prefix), the first that applies, or nil:
--- the site has not seen the character; the gear "on its way" (the site's last copy was seen before
--- this record first had gear and before it was last built); the builds waiting, named with the
--- character they are for; the gear the site has after all (the ack is no older than the export's
--- build time, and an "on its way" was said before: once); the companion silent for days.
function SiteStatus.loginLine(data, facts)
	facts = facts or SiteStatus.facts()
	if facts.seen == nil then
		return L.siteLoginNever, "never"
	end
	-- The site's last copy was seen before this record first carried gear, and before it was last built.
	if facts.gear and facts.gearAt ~= nil and facts.seen < facts.gearAt
		and (facts.builtAt == nil or facts.seen < facts.builtAt) then
		return L.siteLoginGear, "gear"
	end
	local rows = SiteStatus.waiting(data)
	if #rows > 0 then
		local who = BuildDialog.player()
		who = who ~= nil and who.name or nil
		who = who ~= nil and who ~= "" and who or L.siteLoginSelf
		if #rows == 1 then
			return string.format(L.siteLoginBuild, who, nameList(rows)), "build"
		end
		return string.format(L.siteLoginBuilds, #rows, who, nameList(rows)), "build"
	end
	if facts.gear and facts.gearPending and (facts.builtAt == nil or facts.seen >= facts.builtAt) then
		return string.format(L.siteLoginGearSeen, BuildDialog.when(facts.seen, facts.clock)), "gearSeen"
	end
	if facts.hasRecord and facts.age > SiteStatus.SILENT_SECONDS then
		return string.format(L.siteLoginSilent, dayOf(facts.seen, facts.clock)), "silent"
	end
	return nil
end

--- PLAYER_ENTERING_WORLD (`initial`: a login from the character screen, `reloading`: a /reload; both
--- false: a loading screen in the world, which says nothing): five seconds in (after the export's
--- refresh, queued first), say the one line that is true, if any. The rule, so the same line is
--- never said twice in a row for the same inbox: a login always says its line (the inbox was just
--- read, and a build waiting is news at every login until it is used); a /reload says it only when
--- it differs from the line said last (the client writes the old inbox back on a reload, so nothing
--- new can have arrived, and the walk's "it never arrives on /reload" must not be answered with the
--- same sentence again). Using a build says nothing here: this runs on entering the world only. A
--- client that passes neither flag counts as a login once, and as a reload after.
function SiteStatus.afterEnteringWorld(data, initial, reloading)
	local legacy = initial == nil and reloading == nil
	if initial == false and reloading == false then
		return
	end
	local login = initial == true or (legacy and not SiteStatus.announced)
	SiteStatus.announced = true
	local remembered = SiteStatus.memory(true)
	if login then
		remembered.loginAt = BuildDialog.liveClock().now
	end
	Theme.after(SiteStatus.LOGIN_DELAY, function()
		local line, kind = SiteStatus.loginLine(data)
		local again = not login and line == remembered.said
		if line ~= nil and not again then
			Theme.say(line)
			if kind == "gear" then
				remembered.gearPending = true
			elseif kind == "gearSeen" then
				remembered.gearPending = nil
			end
		end
		if line ~= nil or login then
			remembered.said = line
		end
	end)
end

ns.SiteStatus = SiteStatus
return SiteStatus
