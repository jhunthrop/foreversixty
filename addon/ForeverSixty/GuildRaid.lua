-- addon/ForeverSixty/GuildRaid.lua
-- The guild's raids as the hub's Guild tab reads them (design/specs/2026-11-13-addon-hub.md 4.5): the
-- raids the companion's inbox names (raid_signup, raid_assignment, raid_roster; contract 3), your answer
-- to each and the four answers you can give, the guild's sign-ups against the raid's comp target, and
-- what the tab has waiting for you; the guild's raiders by Forever Score and its standing, from the
-- ratings pack and the client's roster. Pure over the inbox, except the outbox reads and the answer itself:
-- an answer is queued in the companion outbox (contract 4 `raid_answer`) and offered only when the
-- companion says it uploads one (contract 3 `accepts`); the addon never touches the network.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Compat = ns.Compat or require("Compat")
local Codec = ns.Codec or require("Codec")
local Export = ns.Export or require("Export")
local Score = ns.Score or require("Score")
local RaidAssign = ns.RaidAssign or require("RaidAssign")
local Chip = ns.Chip or require("Chip")

local GuildRaid = {}

--- The shipped data (ns.Data) the outbox is written with; a spec sets it.
GuildRaid.data = nil
--- Runs after an answer is queued (the Guild tab redraws).
GuildRaid.onAnswer = nil

local SEP = L.dialogSeparator

--- The inbox message the answers read (contract 3); Codec.inboxMessages reads only the kinds it lists.
GuildRaid.SIGNUP_KIND = "raid_signup"
Codec.MESSAGE_TYPES[GuildRaid.SIGNUP_KIND] = true
--- The four answers, in the site's, the apps' and Discord's order.
GuildRaid.ANSWERS = { "in", "late", "bench", "out" }
--- Each answer's colour (the web's raid page: In green, Late the client gold, Bench plain, Out warning).
GuildRaid.ANSWER_HEX = { ["in"] = "success", late = "clientGold", bench = "body", out = "warning" }
--- A raid this close to its start (or running) is tonight's; one further out is the next raid.
GuildRaid.TONIGHT_SECONDS = 18 * 60 * 60

--- The raids' icons, by the site's zone name (the client's own instance art).
GuildRaid.RAID_ICONS = {
	["Onyxia's Lair"] = "Interface\\Icons\\INV_Misc_Head_Dragon_Black",
	["Molten Core"] = "Interface\\Icons\\Spell_Fire_Incinerate",
	["Blackwing Lair"] = "Interface\\Icons\\INV_Misc_Head_Dragon_01",
	["Zul'Gurub"] = "Interface\\Icons\\INV_Misc_Idol_03",
}
GuildRaid.RAID_ICON_DEFAULT = "Interface\\Icons\\INV_Misc_Bone_HumanSkull_01"

--- Whether the companion uploads `raid_answer` entries (the inbox's `accepts`, contract 3): the
--- feature check for the answer buttons. Pure.
function GuildRaid.acceptsAnswers(inbox)
	if type(inbox) ~= "table" or type(inbox.accepts) ~= "table" then
		return false
	end
	for _, name in ipairs(inbox.accepts) do
		if name == "raid_answer" then
			return true
		end
	end
	return false
end

--- Every raid of the character's guild the inbox names and that is not over at `now`, soonest first:
--- { raidId, raid, title, start, signup (its raid_signup), roster (Codec.raidRosters'), jobs
--- (RaidAssign.jobs' for it) }. A raid known only from an assignment (an older companion) is listed
--- too. Pure.
function GuildRaid.raids(inbox, key, now)
	local byId, list = {}, {}
	local function add(raidId, raid, title, startsAt)
		local start = RaidAssign.epochOf(startsAt)
		if raidId == nil or start == nil or (now ~= nil and now > start + RaidAssign.SHOWN_FOR_SECONDS) then
			return nil
		end
		if byId[raidId] == nil then
			byId[raidId] = { raidId = raidId, raid = raid or "", title = title, start = start, jobs = {} }
			list[#list + 1] = byId[raidId]
		end
		return byId[raidId]
	end
	for _, message in ipairs(Codec.inboxMessages(inbox, key, GuildRaid.SIGNUP_KIND)) do
		local raid = add(message.raid_id, message.raid, message.title, message.starts_at)
		if raid ~= nil and raid.signup == nil then
			raid.signup = message
		end
	end
	for _, job in ipairs(RaidAssign.jobs(inbox, key, now)) do
		local raid = add(job.raid_id, job.raid, job.title, job.starts_at)
		if raid ~= nil then
			raid.jobs[#raid.jobs + 1] = job
		end
	end
	for _, roster in ipairs(Codec.raidRosters(inbox, key)) do
		if byId[roster.raidId] ~= nil then
			byId[roster.raidId].roster = roster
		end
	end
	table.sort(list, function(left, right)
		if left.start ~= right.start then
			return left.start < right.start
		end
		return left.raidId < right.raidId
	end)
	return list
end

--- The answer that stands for a raid: a queued one newer than the site's confirmation (`pending`,
--- the outbox entry), else the site's. { answer ("in" ... "out", "none" with no answer, nil with no
--- sign-up at all), pending, at (epoch of the answer), role }. Pure.
function GuildRaid.answerOf(signup, pending)
	local confirmedAt = signup ~= nil and RaidAssign.epochOf(signup.answered_at) or nil
	local role = signup ~= nil and signup.role or nil
	if type(pending) == "table" and GuildRaid.ANSWER_HEX[pending.answer] ~= nil
		and (confirmedAt == nil or confirmedAt < (tonumber(pending.answered_at) or 0)) then
		return { answer = pending.answer, pending = true, at = tonumber(pending.answered_at), role = role }
	end
	if signup == nil then
		return { role = role }
	end
	local answer = GuildRaid.ANSWER_HEX[signup.answer] ~= nil and signup.answer or "none"
	return { answer = answer, at = answer ~= "none" and confirmedAt or nil, role = role }
end

--- "in 3 h 12 min", "in 25 min", or "now" once the raid has started. Pure.
function GuildRaid.countdown(start, now)
	local left = math.floor((start - now) / 60)
	if left <= 0 then
		return L.guildNow
	end
	if left < 60 then
		return string.format(L.guildInMinutes, left)
	end
	return string.format(L.guildInHours, math.floor(left / 60), left % 60)
end

--- "Thu 10 Dec" and "20:00" (the shared clock: "8 pm" on a 12-hour client) for epoch `seconds` on the
--- client's clock; nil where the client has no date function.
function GuildRaid.dayAndClock(seconds)
	local stamp = Compat.localTime(seconds)
	if stamp == nil then
		return nil, nil
	end
	return Compat.dayDate(string.format("%04d-%02d-%02d", stamp.year, stamp.month, stamp.day)),
		Compat.clockText(stamp.hour, stamp.min)
end

--- The weekday ("Tue") of epoch `seconds` on the client's clock, or nil.
function GuildRaid.weekday(seconds)
	local day = type(seconds) == "number" and GuildRaid.dayAndClock(seconds) or nil
	return day ~= nil and day:match("^(%a+)") or nil
end

--- The weekday and clock alone ("Tue 19:04"), for when an answer was sent.
local function sentText(seconds)
	local _, clock = GuildRaid.dayAndClock(seconds)
	local day = GuildRaid.weekday(seconds)
	return day ~= nil and (day .. " " .. clock) or nil
end

--- Whether `raid` is tonight's at `now`: within TONIGHT_SECONDS of its start, or running.
function GuildRaid.isTonight(raid, now)
	return type(raid) == "table" and type(now) == "number" and now >= raid.start - GuildRaid.TONIGHT_SECONDS
end

--- The words of an answer as the sign-up line says them: the answer in its colour, then its detail
--- muted ("as DPS, sent Tue 19:04", "not sent yet"). { word, hex, detail }. `answer` is answerOf's.
function GuildRaid.answerWords(answer)
	if answer.answer == nil then
		return { word = L.guildOnRoster, hex = "body" }
	end
	if answer.answer == "none" then
		return { word = L.guildNotSigned, hex = "clientGold" }
	end
	local parts = {}
	local role = L.guildRoles[answer.role or ""]
	if role ~= nil then
		parts[#parts + 1] = string.format(L.guildAs, role)
	end
	if answer.pending then
		parts[#parts + 1] = L.guildNotSent
	elseif type(answer.at) == "number" and sentText(answer.at) ~= nil then
		parts[#parts + 1] = string.format(L.guildSent, sentText(answer.at))
	end
	return { word = L.guildAnswers[answer.answer], hex = GuildRaid.ANSWER_HEX[answer.answer],
		detail = #parts > 0 and table.concat(parts, ", ") or nil }
end

--- The partners of a job, "Thornhide, Mirelle" split into names. Pure.
local function partners(with)
	local names = {}
	for name in tostring(with or ""):gmatch("[^,]+") do
		local trimmed = name:gsub("^%s+", ""):gsub("%s+$", "")
		if trimmed ~= "" then
			names[#names + 1] = trimmed
		end
	end
	return names
end

--- The top section for `raid` (spec 4.5 left 1). `ctx`: { now, accepts, pending (raidId -> outbox
--- entry or nil) }. { heading, raid, icon, when, answer (answerWords'), buttons (the four, or nil),
--- selected, onSite (the line in place of the answers), jobs = { { label, purpose, with = { names }, encounter } } }.
--- Pure.
function GuildRaid.tonight(raid, ctx)
	local day, clock = GuildRaid.dayAndClock(raid.start)
	local tonight = GuildRaid.isTonight(raid, ctx.now)
	local when = day ~= nil and string.format(L.guildWhen, day, clock) or nil
	if when ~= nil and tonight then
		when = when .. SEP .. GuildRaid.countdown(raid.start, ctx.now)
	end
	local answer = GuildRaid.answerOf(raid.signup, ctx.pending ~= nil and ctx.pending(raid.raidId) or nil)
	local section = {
		heading = tonight and L.guildTonight or L.guildNextRaid, raidId = raid.raidId, raid = raid.raid,
		icon = GuildRaid.RAID_ICONS[raid.raid] or GuildRaid.RAID_ICON_DEFAULT, when = when,
		answer = GuildRaid.answerWords(answer), jobs = {},
	}
	if raid.signup ~= nil and ctx.accepts then
		section.buttons = GuildRaid.ANSWERS
		section.selected = answer.answer ~= "none" and answer.answer or nil
		section.pending = answer.pending == true
	elseif raid.signup ~= nil then
		section.onSite = answer.answer ~= nil and answer.answer ~= "none" and L.guildChangeOnSite or L.guildSignUpOnSite
	end
	for _, job in ipairs(raid.jobs) do
		if type(job.label) == "string" and job.label ~= "" then
			local boss = type(job.encounter) == "string" and job.encounter ~= "" and job.encounter ~= raid.title
				and job.encounter or nil
			section.jobs[#section.jobs + 1] = { label = job.label, purpose = job.purpose, with = partners(job.with),
				encounter = boss }
		end
	end
	return section
end

--- "12 of 40 coming" (or "12 coming" for a raid with no size) from a raid_roster. Pure.
function GuildRaid.comingText(roster)
	return roster.size ~= nil and string.format(L.guildComing, roster.coming, roster.size)
		or string.format(L.guildComingAll, roster.coming)
end

--- Sign-ups (spec 4.5 left 3) from the raid's raid_roster, or nil without one (a companion older than
--- the message): { note, counts = { { text, hex } }, roles = { { text, hex } } (a group short of its
--- target in the warning colour, met in success) }. Pure.
function GuildRaid.signups(roster)
	if type(roster) ~= "table" then
		return nil
	end
	local counts = {}
	for _, answer in ipairs(GuildRaid.ANSWERS) do
		counts[#counts + 1] = { text = string.format(L.guildCount, roster.signups[answer], L.guildAnswers[answer]),
			hex = GuildRaid.ANSWER_HEX[answer] }
	end
	counts[#counts + 1] = { text = string.format(L.guildCountNone, roster.signups.none), hex = "body" }
	local roles = {}
	for _, role in ipairs(roster.roles) do
		local word = L.guildRoleGroups[role.role] or role.role
		if role.wanted ~= nil then
			-- The web's raid page marks a group short of its target; one at or past it reads as met.
			roles[#roles + 1] = { text = string.format(L.guildRoleWanted, word, role.coming, role.wanted),
				hex = role.coming >= role.wanted and "success" or "warning" }
		else
			roles[#roles + 1] = { text = string.format(L.guildRoleComing, word, role.coming), hex = "muted" }
		end
	end
	return { note = GuildRaid.comingText(roster), counts = counts, roles = #roles > 0 and roles or nil }
end

--- Next raids (spec 4.5 right 1): one row per raid, { raidId, when, raid, answer (answerWords'),
--- shown (the raid in the top section) }, at most NEXT_ROWS; nil with fewer than two raids (the one
--- raid is already on top). Pure.
function GuildRaid.nextRaids(raids, shownId, pending)
	if #raids < 2 then
		return nil
	end
	local rows = {}
	for index, raid in ipairs(raids) do
		if index > GuildRaid.NEXT_ROWS then
			break
		end
		local day, clock = GuildRaid.dayAndClock(raid.start)
		local answer = GuildRaid.answerOf(raid.signup, pending ~= nil and pending(raid.raidId) or nil)
		local words = GuildRaid.answerWords(answer)
		local role = L.guildRoles[answer.role or ""]
		if answer.answer ~= nil and answer.answer ~= "none" and answer.answer ~= "out" and role ~= nil then
			words.word = words.word .. " " .. string.format(L.guildAs, role)
		end
		rows[#rows + 1] = { raidId = raid.raidId, when = day ~= nil and (day .. " " .. clock) or "", raid = raid.raid,
			answer = words, shown = raid.raidId == shownId,
			coming = raid.roster ~= nil and GuildRaid.comingText(raid.roster) or nil }
	end
	return rows
end

--- The line the Guild tab leads with on Me's Right now (Hub's `lead`): the soonest raid, its day and
--- time and your answer in its colour, the first job under it. { line, sub, icon } or nil. Pure over
--- `raids` and `ctx`.
function GuildRaid.leadOf(raids, ctx)
	local raid = raids[1]
	if raid == nil then
		return nil
	end
	local day, clock = GuildRaid.dayAndClock(raid.start)
	local answer = GuildRaid.answerWords(GuildRaid.answerOf(raid.signup, ctx.pending ~= nil and ctx.pending(raid.raidId)
		or nil))
	local when = GuildRaid.isTonight(raid, ctx.now) and clock or (day ~= nil and (day .. " " .. clock) or nil)
	local job = raid.jobs[1]
	return {
		line = string.format(L.guildChipLine, raid.raid, when or "", Theme.inlineColor(answer.hex, answer.word)),
		sub = job ~= nil and job.label ~= nil and (L.guildYourJob .. " " .. job.label) or nil,
		icon = GuildRaid.RAID_ICONS[raid.raid] or GuildRaid.RAID_ICON_DEFAULT,
	}
end

--- What the tab has waiting for the player (spec 3.3, 3.6): tonight's raid not signed up, and the
--- jobs tonight's raid gives them. Each { key, line, fact }. Pure over `raids` and `ctx`.
function GuildRaid.waitingOf(raids, ctx)
	local raid = raids[1]
	if raid == nil or not GuildRaid.isTonight(raid, ctx.now) then
		return {}
	end
	local _, clock = GuildRaid.dayAndClock(raid.start)
	local items = {}
	local answer = GuildRaid.answerOf(raid.signup, ctx.pending ~= nil and ctx.pending(raid.raidId) or nil)
	if answer.answer == "none" then
		items[#items + 1] = { key = "signup", fact = tostring(raid.raidId),
			line = string.format(L.guildChipLine, raid.raid, clock or "", Theme.inlineColor("clientGold", L.guildSignUp)) }
	end
	local labels = {}
	for _, job in ipairs(raid.jobs) do
		labels[#labels + 1] = tostring(job.label or job.slot or "")
	end
	if #labels > 0 then
		items[#items + 1] = { key = "job", fact = raid.raidId .. ":" .. table.concat(labels, ","),
			line = string.format(L.guildChipLine, raid.raid, clock or "", Theme.inlineColor("clientGold", labels[1])) }
	end
	return items
end

--- This character's queued answer for `raidId` in the outbox, or nil.
function GuildRaid.pendingAnswer(raidId)
	local character = Export.outboxCharacter()
	local outbox = type(ForeverSixtyDB) == "table" and ForeverSixtyDB.outbox or nil
	local list = Export.mergeRaidAnswers(type(outbox) == "table" and outbox.raid_answer or nil, Export.raidAnswers)
	for _, entry in ipairs(list) do
		if entry.raid_id == raidId and entry.character == character then
			return entry
		end
	end
	return nil
end

--- Every raid the inbox names for this character now (GuildRaid.raids over the live inbox).
function GuildRaid.liveRaids()
	return GuildRaid.raids(ForeverSixtyInbox, Export.characterKey(), Score.clock())
end

--- The model's shared context from the client: { now, accepts, pending }.
function GuildRaid.context()
	return { now = Score.clock(), accepts = GuildRaid.acceptsAnswers(ForeverSixtyInbox),
		pending = GuildRaid.pendingAnswer }
end

--- An answer button: queued for the companion at once (Export.saveOutbox writes the outbox; the file
--- is written at the next reload or logout), only when the companion takes answers. False when it did
--- not queue.
function GuildRaid.answer(raidId, answer)
	if raidId == nil or GuildRaid.ANSWER_HEX[answer] == nil or not GuildRaid.acceptsAnswers(ForeverSixtyInbox) then
		return false
	end
	Export.raidAnswers = Export.mergeRaidAnswers(Export.raidAnswers, { {
		character = Export.outboxCharacter(), raid_id = raidId, answer = answer,
		answered_at = type(time) == "function" and time() or 0,
	} })
	Export.saveOutbox(GuildRaid.data or ns.Data)
	if GuildRaid.onAnswer ~= nil then
		GuildRaid.onAnswer()
	end
	return true
end

-- The guild's raiders and its standing (the ratings pack) -----------------------------------------

--- The guild's top raiders shown, and the most Next raids rows.
GuildRaid.TOP = 5
GuildRaid.NEXT_ROWS = 5

--- The level every raider is when the guild's officers have not named raider ranks on the site.
GuildRaid.RAID_LEVEL = 60

--- Whether a roster entry counts as a raider: its rank is one the guild's officers named on the site
--- (`ranks`, the ratings pack's raider_ranks), else, while they have named none, it is level 60.
function GuildRaid.isRaider(entry, ranks)
	if type(ranks) == "table" and #ranks > 0 then
		for _, rank in ipairs(ranks) do
			if rank == entry.rankIndex then
				return true
			end
		end
		return false
	end
	return entry.level == GuildRaid.RAID_LEVEL
end

--- The guild's standings (spec 4.5 right 2): the raiders (GuildRaid.isRaider) with a Forever Score,
--- best first. `record(name)` is Score.forCharacter; `own` the player's roster name. { top (at most
--- TOP rows), own (the player's row when not in the top), place, of (every raider) }; each row
--- { place, name, classToken, spec, score, tier, attendance }. Pure over its inputs.
function GuildRaid.standings(roster, record, ranks, own)
	local raiders, scored = 0, {}
	for _, entry in ipairs(roster) do
		if GuildRaid.isRaider(entry, ranks) then
			raiders = raiders + 1
			local found = record(entry.name)
			if type(found) == "table" and type(found.score) == "number" and (found.score > 0 or (found.fights or 0) > 0) then
				scored[#scored + 1] = { name = entry.name, short = (entry.name:match("^([^%-]+)") or entry.name),
					classToken = entry.classToken, spec = found.spec, score = found.score, tier = found.tier,
					attendance = type(found.attendance) == "number" and math.floor(found.attendance * 100 + 0.5) or nil }
			end
		end
	end
	table.sort(scored, function(left, right)
		if left.score ~= right.score then
			return left.score > right.score
		end
		return left.short < right.short
	end)
	local out = { top = {}, of = raiders }
	for index, row in ipairs(scored) do
		row.place = index
		if index <= GuildRaid.TOP then
			out.top[#out.top + 1] = row
		end
		if own ~= nil and row.short == own then
			out.place = index
			if index > GuildRaid.TOP then
				out.own = row
			end
		end
	end
	return out
end

--- The guild's one muted line from its record (the ratings pack): its standing among guilds, its raid
--- nights in 90 days (each raider's attendance is the table's column); "Ratings 8 days old" closes it
--- while the pack is stale. nil with nothing to say. Pure.
function GuildRaid.guildLine(record, staleDays)
	local parts = {}
	if type(record) == "table" then
		if type(record.standing) == "number" and record.standing > 0 then
			parts[#parts + 1] = string.format(L.guildStanding, Chip.ordinal(record.standing))
		end
		-- An old pack's age takes the nights' place, so the line stays one line in the column.
		if type(record.nights) == "number" and record.nights > 0 and staleDays == nil then
			parts[#parts + 1] = string.format(L.guildNights, record.nights)
		end
	end
	if type(staleDays) == "number" then
		parts[#parts + 1] = Theme.inlineColor("warning", string.format(L.guildRatingsOld, staleDays))
	end
	return #parts > 0 and table.concat(parts, SEP) or nil
end

--- The guild block's model: { name, line, note, top, own } from the ratings pack and the client's roster,
--- or nil without a usable pack (spec 8: the top five and your place are left out).
function GuildRaid.guildModel(guildName)
	local status = Score.status()
	if guildName == nil or not status.available then
		return nil
	end
	local record = Score.guild(guildName)
	local staleDays
	if status.stale then
		local written = Score.dateSeconds(status.generated)
		staleDays = written ~= nil and math.floor((Score.clock() - written) / 86400) or nil
	end
	local standings = GuildRaid.standings(Theme.guildRoster(), function(name)
		return Score.forCharacter(name)
	end, type(record) == "table" and record.raider_ranks or nil, Compat.playerName())
	local block = { name = guildName, line = GuildRaid.guildLine(record, staleDays), top = standings.top,
		own = standings.own }
	if standings.place ~= nil then
		local ranks = type(record) == "table" and record.raider_ranks or nil
		local named = type(ranks) == "table" and #ranks > 0
		block.note = string.format(named and L.guildYouAre or L.guildYouAreAtCap, Chip.ordinal(standings.place),
			standings.of)
	end
	return block
end

ns.GuildRaid = GuildRaid
return GuildRaid
