-- addon/ForeverSixty/RaidAssign.lua
-- Your job in the next raid (guild raid assignments, web layout spec 2026-11-05 §7.2). The raid
-- leader picks the names on the site; the inbox carries one `raid_assignment` message per slot
-- the logged-in character holds in its next raid (docs/contracts/addon-v2.md contract 3; the
-- companion carries the flat message through as is). The hub's Guild tab shows them (RaidAssign.jobs);
-- this module prints them once when the player zones into that raid, the one moment a window cannot
-- be opened. It never picks anything and sends nothing.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local Codec = ns.Codec or require("Codec")
local Theme = ns.Theme or require("Theme")

local RaidAssign = {}

--- The inbox message type; Codec.inboxMessages reads only the kinds it lists, so this module
--- adds its own when it loads (it loads after Codec).
RaidAssign.KIND = "raid_assignment"
Codec.MESSAGE_TYPES[RaidAssign.KIND] = true

--- A raid's jobs stop showing this long after its start: the night is over by then.
RaidAssign.SHOWN_FOR_SECONDS = 6 * 60 * 60

--- The zone events the raid-entry line listens on.
RaidAssign.EVENTS = { "PLAYER_ENTERING_WORLD", "ZONE_CHANGED_NEW_AREA" }

--- The player's words. Kept here rather than in Locale.lua so the module stands alone.
RaidAssign.COPY = {
	header = "Your jobs for %s:",
	headerBoss = "Your jobs for %s, %s:",
	line = "%s: %s",
	lineWith = "%s with %s: %s",
}

--- Seconds since the epoch for "YYYY-MM-DDTHH:MM:SSZ" (UTC), or nil for anything else. Pure:
--- the client's date functions read the local clock, so the civil-date arithmetic is done here.
function RaidAssign.epochOf(stamp)
	if type(stamp) ~= "string" then
		return nil
	end
	local y, m, d, hh, mm, ss = stamp:match("^(%d%d%d%d)%-(%d%d)%-(%d%d)T(%d%d):(%d%d):(%d%d)Z$")
	if y == nil then
		return nil
	end
	y, m, d = tonumber(y), tonumber(m), tonumber(d)
	-- Days from 1970-01-01 to the civil date (Howard Hinnant's days_from_civil).
	local shifted = m <= 2 and y - 1 or y
	local era = math.floor(shifted / 400)
	local yoe = shifted - era * 400
	local doy = math.floor((153 * (m + (m > 2 and -3 or 9)) + 2) / 5) + d - 1
	local doe = yoe * 365 + math.floor(yoe / 4) - math.floor(yoe / 100) + doy
	local days = era * 146097 + doe - 719468
	return days * 86400 + tonumber(hh) * 3600 + tonumber(mm) * 60 + tonumber(ss)
end

--- The client's clock in epoch seconds, or nil on a client without time().
function RaidAssign.now()
	return type(time) == "function" and time() or nil
end

--- The logged-in character's jobs in its next raid, in the order the site sent them: only the
--- soonest raid still on (its start no more than SHOWN_FOR_SECONDS ago). Empty when none.
function RaidAssign.jobs(inbox, key, now)
	local soonest, soonestStart
	local held = {}
	for _, message in ipairs(Codec.inboxMessages(inbox, key, RaidAssign.KIND)) do
		local start = RaidAssign.epochOf(message.starts_at)
		if start ~= nil and (now == nil or now <= start + RaidAssign.SHOWN_FOR_SECONDS) then
			held[#held + 1] = message
			if soonestStart == nil or start < soonestStart then
				soonest, soonestStart = message.raid_id, start
			end
		end
	end
	local jobs = {}
	for _, message in ipairs(held) do
		if message.raid_id == soonest then
			jobs[#jobs + 1] = message
		end
	end
	return jobs
end

--- The character's next raid for the chip's raid-day line and the Guild tab
--- (design/specs/2026-11-08-addon-ux.md 3.2; 2026-11-13-addon-hub 4.5): { raid_id, raid, title, start } from its
--- soonest assignment still on (RaidAssign.jobs), or nil with none. Until the inbox carries
--- `raid_signup`, an assignment is the only sign of a raid (it is sent only to In, Late or Bench).
function RaidAssign.tonight(inbox, key, now)
	local jobs = RaidAssign.jobs(inbox, key, now)
	local first = jobs[1]
	if first == nil then
		return nil
	end
	return { raid_id = first.raid_id, raid = first.raid, title = first.title,
		start = RaidAssign.epochOf(first.starts_at), jobs = jobs }
end

--- One line per job, under a header naming the raid: "Main tank with Thornhide: Picks her up...";
--- none without a job.
function RaidAssign.lines(jobs)
	local first = jobs[1]
	if first == nil then
		return {}
	end
	local title = tostring(first.title or first.raid or "")
	local boss = first.encounter
	local lines = {}
	if type(boss) == "string" and boss ~= "" and boss ~= title then
		lines[1] = string.format(RaidAssign.COPY.headerBoss, title, boss)
	else
		lines[1] = string.format(RaidAssign.COPY.header, title)
	end
	for _, job in ipairs(jobs) do
		local label = Theme.inlineColor("gold", tostring(job.label or job.slot or ""))
		local purpose = tostring(job.purpose or "")
		if type(job.with) == "string" and job.with ~= "" then
			lines[#lines + 1] = string.format(RaidAssign.COPY.lineWith, label, job.with, purpose)
		else
			lines[#lines + 1] = string.format(RaidAssign.COPY.line, label, purpose)
		end
	end
	return lines
end

--- A name with case, spaces and punctuation dropped, so the client's instance name and the
--- site's zone name meet ("Onyxia's Lair" and "Onyxias Lair").
function RaidAssign.fold(name)
	return (tostring(name or ""):lower():gsub("[%s%p]", ""))
end

--- The raids whose entry line has printed this session, by raid id.
RaidAssign.shown = {}

--- On a zone change: inside the raid the jobs are for, print them once per raid per session.
--- Answers the lines printed, or nil.
function RaidAssign.onZone(inbox, key, now)
	local name, instanceType = Theme.instanceName()
	if instanceType ~= "raid" then
		return nil
	end
	local jobs = RaidAssign.jobs(inbox, key, now)
	local first = jobs[1]
	if first == nil or RaidAssign.shown[first.raid_id] or RaidAssign.fold(first.raid) ~= RaidAssign.fold(name) then
		return nil
	end
	RaidAssign.shown[first.raid_id] = true
	local lines = RaidAssign.lines(jobs)
	for _, line in ipairs(lines) do
		Theme.say(line)
	end
	return lines
end

--- The character key the inbox addresses; Export loads after this file, so it is looked up late.
local function characterKey()
	local Export = ns.Export or require("Export")
	return Export.characterKey()
end

function RaidAssign.onEvent()
	return RaidAssign.onZone(ForeverSixtyInbox, characterKey(), RaidAssign.now())
end

--- Runs at file load: one hidden event frame for the zone events, where the client has frames.
function RaidAssign.register()
	if type(CreateFrame) ~= "function" then
		return nil
	end
	local frame = CreateFrame("Frame", nil, UIParent)
	frame:SetScript("OnEvent", RaidAssign.onEvent)
	for _, event in ipairs(RaidAssign.EVENTS) do
		Theme.registerEvent(frame, event)
	end
	return frame
end

RaidAssign.frame = RaidAssign.register()

ns.RaidAssign = RaidAssign
return RaidAssign
