-- addon/ForeverSixty/surfaces/BuildDialog.lua
-- The build dialog (spec section 4.10): the one frame of the addon's own, 420x300.
-- Two tabs, one job each: paste a build in ("Paste a build"), copy the character out
-- ("Copy my character", the site's own words).
--
-- Pure model, thin frame. BuildDialog.evaluate is the paste box's whole verdict
-- (Codec.loadBuild's own answer, the codec's own reason on a refusal); pasteView,
-- refusal, preview, followingLine and nextTalent are what paste mode shows; copyModel
-- is the whole of copy mode with the clock handed in; inboxRows and rowsView are the
-- "from the site" list. The frame is built on the first open, never at load, and every
-- number on it comes from the function that owns it (Progress.planSplit for the split,
-- Progress.unspentPoints for the points to place, Follow.buildName for the name, the
-- chip's Chip.pointLabel for the next talent).
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Cards = ns.Cards or require("Cards")
local Prefs = ns.Prefs or require("Prefs")
local Compat = ns.Compat or require("Compat")
local Codec = ns.Codec or require("Codec")
local Talents = ns.Talents or require("Talents")
local Export = ns.Export or require("Export")
local Follow = ns.Follow or require("Follow")
local Progress = ns.Progress or require("Progress")
local Chip = ns.Chip or require("Chip")

local BuildDialog = {}

BuildDialog.FRAME_NAME = "ForeverSixtyBuildDialog"
--- How many of the build's planned items the preview shows: the code's first three.
BuildDialog.GEAR_PICKS = 3
--- The most characters a refusal's lead keeps (two lines of the dialog's text).
BuildDialog.REASON_CHARS = 120
--- How far an edge may sit past the screen's and still count as on it (UI px).
BuildDialog.SCREEN_TOLERANCE = 1
BuildDialog.MODES = { "paste", "copy" }

--- Injection points Options fills: the data to read, what to redraw after a build is
--- used (Options.refreshSurfaces), and what to do instead when the dialog cannot show
--- itself (Options.showHint). `say` prints the chat line (Theme.say when nil).
BuildDialog.data = nil
BuildDialog.afterUse = nil
BuildDialog.fallback = nil
BuildDialog.say = nil

--- Item ids whose names the preview is waiting on (GET_ITEM_INFO_RECEIVED redraws it).
BuildDialog.pending = {}

-- Pure rules: paste mode -----------------------------------------------------------------

--- The paste box's verdict on `text`: { ok = true, build, code } (the code trimmed, which
--- is what Use loads) when the codec loads it, else { ok = false, reason } with the codec's
--- own refusal. `player` (Codec.loadBuild's) refuses another class's code. An empty box has
--- no reason: there is nothing to refuse yet.
function BuildDialog.evaluate(text, data, player)
	if type(text) ~= "string" or not text:find("%S") then
		return { ok = false }
	end
	local trimmed = text:match("^%s*(.-)%s*$")
	local build, reason = Codec.loadBuild(trimmed, data, player)
	if build == nil then
		return { ok = false, reason = reason }
	end
	return { ok = true, build = build, code = trimmed }
end

--- How the paste box sits: collapsed to one line for a preview or a refusal (opened again
--- while the player edits), the refusal's lines and red border until the box is edited, the
--- gold border while it has the keyboard.
function BuildDialog.pasteView(verdict, editing, focused)
	local ok = verdict ~= nil and verdict.ok == true
	local refused = verdict ~= nil and not ok and verdict.reason ~= nil
	local showReason = refused and not editing
	local border = focused and "gold" or "border"
	if showReason then
		border = "warning"
	end
	return {
		collapsed = (ok or refused) and not editing,
		reason = showReason,
		preview = ok,
		use = ok,
		border = border,
	}
end

--- `text` cut to at most `limit` bytes without splitting a UTF-8 letter.
local function cut(text, limit)
	local stop = limit
	while stop > 0 do
		local byte = text:byte(stop + 1)
		if byte == nil or byte < 128 or byte >= 192 then
			break
		end
		stop = stop - 1
	end
	return text:sub(1, stop)
end

--- A refusal as the dialog shows it: the codec's first sentence as the red lead (capped at
--- REASON_CHARS, so it never runs past two lines), the rest as the body, and a muted
--- "You are still following ..." when a build is followed.
function BuildDialog.refusal(reason, followingName)
	local lead, body = reason:match("^(.-%.)%s+(.+)$")
	if lead == nil then
		lead = reason
	end
	if #lead > BuildDialog.REASON_CHARS then
		lead = cut(lead, BuildDialog.REASON_CHARS) .. L.ellipsis
	end
	return {
		lead = lead,
		body = body,
		following = followingName ~= nil and string.format(L.dialogStillFollowing, followingName) or nil,
	}
end

--- The build's trees by the points its order spends in each (Progress.planSplit), biggest
--- first: "Arms 31 / Fury 20". Empty for a build with no points.
function BuildDialog.split(build, data)
	local counts = Progress.planSplit(build)
	local tabs = {}
	for tab, count in pairs(counts) do
		if count > 0 then
			tabs[#tabs + 1] = tab
		end
	end
	table.sort(tabs, function(left, right)
		if counts[left] ~= counts[right] then
			return counts[left] > counts[right]
		end
		return left < right
	end)
	local parts = {}
	for _, tab in ipairs(tabs) do
		local name = Talents.tabName(data, build.classSlug, tab)
		if name ~= nil then
			parts[#parts + 1] = string.format(L.dialogSplitPart, name, counts[tab])
		end
	end
	return table.concat(parts, L.dialogSplitSeparator)
end

local function pointsWord(points)
	if points == 0 then
		return L.dialogNoPoints
	elseif points == 1 then
		return L.dialogOnePoint
	end
	return string.format(L.dialogPoints, points)
end

--- The parts that say something, joined by " · " (a nil or empty part is skipped).
local function joined(...)
	local kept = {}
	for index = 1, select("#", ...) do
		local part = select(index, ...)
		if part ~= nil and part ~= false and part ~= "" then
			kept[#kept + 1] = part
		end
	end
	return table.concat(kept, L.dialogSeparator)
end

--- "Arms 33 / Fury 16 / Protection 2 · 51 points": the split and the points, no class.
local function splitAndPoints(build, data)
	local points = #build.order
	return joined(points > 0 and BuildDialog.split(build, data) or nil, pointsWord(points))
end

--- The client's own word for a code slot (HEADSLOT = "Head" in `words`, a table or a
--- lookup), else the addon's own (L.gearSlotLabels).
BuildDialog.SLOT_GLOBALS = {
	head = "HEADSLOT", neck = "NECKSLOT", shoulder = "SHOULDERSLOT", back = "BACKSLOT", chest = "CHESTSLOT",
	wrist = "WRISTSLOT", hands = "HANDSSLOT", waist = "WAISTSLOT", legs = "LEGSSLOT", feet = "FEETSLOT",
	finger1 = "FINGER0SLOT", finger2 = "FINGER1SLOT", trinket1 = "TRINKET0SLOT", trinket2 = "TRINKET1SLOT",
	main_hand = "MAINHANDSLOT", off_hand = "SECONDARYHANDSLOT", ranged = "RANGEDSLOT",
}

function BuildDialog.slotWord(slot, words)
	local global = BuildDialog.SLOT_GLOBALS[slot]
	local word = global ~= nil and type(words) == "table" and words[global] or nil
	if type(word) == "string" and word ~= "" then
		return word
	end
	return L.gearSlotLabels[slot] or slot
end

--- What the preview shows for a build the codec loaded: the class token (for
--- Theme.classIcon), the name (Follow.buildName: the build's own, else its trees), the line
--- "Warrior · Arms 33 / Fury 16 · 51 points · you can place 12 now" (the last part only
--- while points wait, Progress.unspentPoints), and the code's first three gear fields with
--- their slot words under a heading ("Gear in this build · at 60" below 60). `context` =
--- { level, unspent, slotWord = function(slot) }, all optional.
function BuildDialog.preview(build, data, context)
	context = context or {}
	local points = #build.order
	local unspent = math.min(context.unspent or 0, points)
	local detail = joined(
		Talents.titleCase(build.classSlug),
		points > 0 and BuildDialog.split(build, data) or nil,
		pointsWord(points),
		unspent > 0 and string.format(L.dialogPlaceNow, unspent) or nil
	)
	local wordOf = context.slotWord or function(slot)
		return BuildDialog.slotWord(slot, {})
	end
	local picks = {}
	for _, entry in ipairs(build.gear or {}) do
		if #picks >= BuildDialog.GEAR_PICKS then
			break
		end
		if type(entry.itemId) == "number" then
			picks[#picks + 1] = { itemId = entry.itemId, slot = entry.slot, slotWord = wordOf(entry.slot) }
		end
	end
	local heading
	if #picks > 0 then
		heading = L.dialogGearHeading
		if type(context.level) == "number" and context.level < Progress.MAX_LEVEL then
			heading = string.format(L.dialogGearAt, heading, Progress.MAX_LEVEL)
		end
	end
	return {
		classToken = build.classSlug:upper(),
		name = Follow.buildName(data, build),
		detail = detail,
		picks = picks,
		heading = heading,
	}
end

--- The next talent the build wants, as the chip says it ("Deflection (1 of 2)",
--- Chip.pointLabel), or nil with no build or nothing left to take.
function BuildDialog.nextTalent(build, ranks, data, level)
	if build == nil then
		return nil
	end
	local point = Progress.build(build, ranks, level).point
	if point == nil then
		return nil
	end
	local talent = Talents.cellOf(data, build.classSlug, point.tab, point.tier, point.column)
	local name = Chip.talentName(talent) or L.chipNextTalent
	return Chip.pointLabel(name, Chip.pointRankAt(build, ranks, data, point))
end

--- "Next talent: Take Deflection (1 of 2)", the chip's own words; nil without a label.
function BuildDialog.nextTalentLine(label)
	if label == nil then
		return nil
	end
	return string.format(L.dialogNextTalent, string.format(L.chipTake, label))
end

--- "Following: Arms raid · Warrior · 51 points", nil with nothing followed.
function BuildDialog.followingLine(build, data)
	if build == nil then
		return nil
	end
	return string.format(L.dialogFollowingLine, joined(
		Follow.buildName(data, build), Talents.titleCase(build.classSlug), pointsWord(#build.order)
	))
end

--- The chat line after Use: "Now following Arms raid. Next: Deflection (1 of 2).", the
--- name's pipes doubled so a site-sent name cannot render as markup.
function BuildDialog.chatLine(name, nextLabel)
	local safe = (tostring(name or ""):gsub("|", "||"))
	if nextLabel ~= nil then
		return string.format(L.dialogNowFollowingNext, safe, nextLabel)
	end
	return string.format(L.dialogNowFollowing, safe)
end

-- Pure rules: the clock ------------------------------------------------------------------

--- Seconds since 1970-01-01 00:00 UTC for a date and time read as UTC (the proleptic
--- Gregorian day count).
function BuildDialog.civilSeconds(year, month, day, hour, minute, second)
	local y = month <= 2 and year - 1 or year
	local era = math.floor(y / 400)
	local yearOfEra = y - era * 400
	local shifted = (month + 9) % 12
	local dayOfYear = math.floor((153 * shifted + 2) / 5) + day - 1
	local dayOfEra = yearOfEra * 365 + math.floor(yearOfEra / 4) - math.floor(yearOfEra / 100) + dayOfYear
	local days = era * 146097 + dayOfEra - 719468
	return days * 86400 + (hour or 0) * 3600 + (minute or 0) * 60 + (second or 0)
end

--- The UTC date and time of `seconds`, in the shape date("*t") answers.
function BuildDialog.utcFields(seconds)
	local days = math.floor(seconds / 86400)
	local ofDay = seconds - days * 86400
	local z = days + 719468
	local era = math.floor(z / 146097)
	local dayOfEra = z - era * 146097
	local yearOfEra = math.floor((dayOfEra - math.floor(dayOfEra / 1460) + math.floor(dayOfEra / 36524)
		- math.floor(dayOfEra / 146096)) / 365)
	local dayOfYear = dayOfEra - (365 * yearOfEra + math.floor(yearOfEra / 4) - math.floor(yearOfEra / 100))
	local shifted = math.floor((5 * dayOfYear + 2) / 153)
	local day = dayOfYear - math.floor((153 * shifted + 2) / 5) + 1
	local month = shifted < 10 and shifted + 3 or shifted - 9
	local year = yearOfEra + era * 400 + (month <= 2 and 1 or 0)
	return {
		year = year, month = month, day = day,
		hour = math.floor(ofDay / 3600), min = math.floor(ofDay % 3600 / 60), sec = ofDay % 60,
	}
end

--- A clock at a fixed offset from UTC, for specs: { now, fields(seconds), epoch(fields) }.
--- The frame's clock is the client's own (date and time), so daylight saving is the client's.
function BuildDialog.utcClock(offset, now)
	offset = offset or 0
	return {
		now = now or 0,
		fields = function(seconds)
			return BuildDialog.utcFields(seconds + offset)
		end,
		epoch = function(fields)
			return BuildDialog.civilSeconds(fields.year, fields.month, fields.day, fields.hour, fields.min,
				fields.sec or 0) - offset
		end,
	}
end

local function validFields(month, day, hour, minute, second)
	return month >= 1 and month <= 12 and day >= 1 and day <= 31 and hour <= 23 and minute <= 59
		and (second or 0) <= 60
end

--- An RFC 3339 time (the ack's `export_seen_at`: "2026-11-04T14:03:00Z", with an offset
--- or fractional seconds too) as seconds since the epoch, or nil.
function BuildDialog.parseUtc(text)
	if type(text) ~= "string" then
		return nil
	end
	local year, month, day, hour, minute, second, zone =
		text:match("^(%d%d%d%d)%-(%d%d)%-(%d%d)[Tt](%d%d):(%d%d):(%d%d)%.?%d*(.*)$")
	if year == nil then
		return nil
	end
	year, month, day = tonumber(year), tonumber(month), tonumber(day)
	hour, minute, second = tonumber(hour), tonumber(minute), tonumber(second)
	if not validFields(month, day, hour, minute, second) then
		return nil
	end
	local offset
	if zone == "Z" or zone == "z" then
		offset = 0
	else
		local sign, zoneHour, zoneMinute = zone:match("^([+-])(%d%d):?(%d%d)$")
		if sign == nil then
			return nil
		end
		offset = (tonumber(zoneHour) * 3600 + tonumber(zoneMinute) * 60) * (sign == "-" and -1 or 1)
	end
	return BuildDialog.civilSeconds(year, month, day, hour, minute, second) - offset
end

--- The outbox's local wall time ("2026-11-04 14:02") as seconds since the epoch, through
--- `clock` (its epoch), or nil for anything else.
function BuildDialog.parseLocal(text, clock)
	if type(text) ~= "string" then
		return nil
	end
	local year, month, day, hour, minute = text:match("^(%d%d%d%d)%-(%d%d)%-(%d%d) (%d%d):(%d%d)$")
	if year == nil then
		return nil
	end
	year, month, day = tonumber(year), tonumber(month), tonumber(day)
	hour, minute = tonumber(hour), tonumber(minute)
	if not validFields(month, day, hour, minute, 0) then
		return nil
	end
	return clock.epoch({ year = year, month = month, day = day, hour = hour, min = minute, sec = 0 })
end

local function hhmm(fields)
	return string.format("%02d:%02d", fields.hour, fields.min)
end

--- When `seconds` was, in local time: "14:02" today, "yesterday 14:02", else "2 Nov 14:02".
function BuildDialog.when(seconds, clock)
	local at, today = clock.fields(seconds), clock.fields(clock.now)
	local dayAt = BuildDialog.civilSeconds(at.year, at.month, at.day, 0, 0, 0)
	local dayToday = BuildDialog.civilSeconds(today.year, today.month, today.day, 0, 0, 0)
	if dayAt == dayToday then
		return hhmm(at)
	elseif dayToday - dayAt == 86400 then
		return string.format(L.dialogYesterday, hhmm(at))
	end
	local shortDate = Compat.shortDate(string.format("%04d-%02d-%02d", at.year, at.month, at.day)) or ""
	return string.format(L.dialogOnDate, shortDate, hhmm(at))
end

-- Pure rules: copy mode ------------------------------------------------------------------

--- Whether a key press in the export box is a real copy: Ctrl+C while the box has the
--- keyboard and the export selected. Only this prints a copy claim.
function BuildDialog.copyKey(key, control, focused, selected)
	return key == "C" and control == true and focused == true and selected == true
end

--- The ack's line about the export: "Seen by the site 14:03" when the site's copy is at or
--- after this computer's last save (the upload follows the save, so "the same export"
--- reads as not older), "The site has your character from yesterday 22:30. Yours has
--- changed since." when it is older; nil without a readable ack.
local function seenLine(ack, savedAt, clock)
	if type(ack) ~= "table" then
		return nil
	end
	local seen = BuildDialog.parseUtc(ack.export_seen_at)
	if seen == nil then
		return nil
	end
	local saved = BuildDialog.parseLocal(savedAt, clock)
	if saved ~= nil and seen < saved then
		return string.format(L.dialogSiteOlder, BuildDialog.when(seen, clock))
	end
	return string.format(L.dialogSeen, BuildDialog.when(seen, clock))
end

--- The ack's sim request state in words: "Sim queued 13:40. The result arrives when you
--- next log in." (the time from this character's outbox request), or the other three
--- states' words; nothing for a state this addon does not know.
local function simLine(state, requestedAt, clock)
	if state == "queued" then
		local at = BuildDialog.parseLocal(requestedAt, clock)
		if at ~= nil then
			return string.format(L.dialogSimQueuedAt, hhmm(clock.fields(at)))
		end
		return L.dialogSimQueued
	end
	if type(state) ~= "string" then
		return nil
	end
	return L.dialogSimState[state]
end

--- The whole of copy mode: the box's text and whether it is selected, the lead line under
--- the tabs (the copy claim only after a real Ctrl+C, `copiedAt`), whether the button
--- works, the ack's lines, and the companion's explanation when no ack was ever seen.
--- `input` = { code, reason, copiedAt, ack, savedAt, simRequestedAt, clock }.
function BuildDialog.copyModel(input)
	local clock = input.clock
	local code = input.code or nil
	local lead
	if not code then
		lead = { text = input.reason or "", hexKey = "muted" }
	elseif input.copiedAt ~= nil then
		lead = { text = string.format(L.dialogYouCopied, hhmm(clock.fields(input.copiedAt))), hexKey = "body" }
	else
		lead = { text = L.dialogSelectedLead, hexKey = "body" }
	end
	local ack = type(input.ack) == "table" and input.ack or nil
	local companion
	if ack == nil then
		companion = { L.dialogNoCompanion, L.dialogCompanionWhat, string.format(L.dialogPasteAtSite, L.siteHost) }
	end
	return {
		text = code or "",
		selected = code ~= nil and code ~= false,
		button = code ~= nil and code ~= false,
		lead = lead,
		seen = seenLine(ack, input.savedAt, clock),
		sim = ack ~= nil and simLine(ack.sim_request_state, input.simRequestedAt, clock) or nil,
		companion = companion,
	}
end

--- The outbox (contract 4, last session's) only when it is this character's (`outboxName`,
--- Export.outboxCharacter): its export's `saved_at`, and its sim request's `requested_at`.
local function outboxField(db, outboxName, part, field)
	local outbox = type(db) == "table" and db.outbox or nil
	local entry = type(outbox) == "table" and outbox[part] or nil
	if type(entry) ~= "table" or entry.character ~= outboxName or type(entry[field]) ~= "string" then
		return nil
	end
	return entry[field]
end

function BuildDialog.savedAtFor(db, outboxName)
	return outboxField(db, outboxName, "export", "saved_at")
end

function BuildDialog.simRequestedAtFor(db, outboxName)
	return outboxField(db, outboxName, "sim_request", "requested_at")
end

-- Pure rules: builds from the site -------------------------------------------------------

--- The "from the site" rows: one per build in `builds` (Follow.inboxBuilds) whose id and
--- code are both unused (`dismissed`, Follow.dismissedSet) and whose code loads for this
--- player, as { id, name, code, label }, the label "Arms raid · Arms 33 / Fury 16 · 51
--- points · from the site". Another class's build is left out without a word (it was sent
--- for another character of the account). The second value is the first other refusal met,
--- for /fs inbox.
function BuildDialog.inboxRows(builds, dismissed, data, player)
	local rows, refusal = {}, nil
	dismissed = dismissed or {}
	for _, entry in ipairs(builds or {}) do
		if not dismissed[entry.id] and not dismissed[entry.code] then
			local verdict = BuildDialog.evaluate(entry.code, data, player)
			local otherClass = not verdict.ok and player ~= nil and BuildDialog.evaluate(entry.code, data).ok
			if verdict.ok then
				local name = type(entry.name) == "string" and entry.name ~= "" and entry.name
					or Follow.buildName(data, verdict.build)
				rows[#rows + 1] = {
					id = entry.id,
					name = name,
					code = entry.code,
					label = string.format(L.dialogInboxRow, (name:gsub("|", "||")),
						splitAndPoints(verdict.build, data)),
				}
			elseif refusal == nil and not otherClass then
				refusal = verdict.reason
			end
		end
	end
	return rows, refusal
end

--- The rows paste mode shows (at most `limit`) and "and 2 more" for the rest.
function BuildDialog.rowsView(rows, limit)
	local shown = {}
	for index = 1, math.min(limit, #rows) do
		shown[index] = rows[index]
	end
	local more = #rows > limit and string.format(L.dialogInboxMore, #rows - limit) or nil
	return { shown = shown, more = more }
end

--- The paste tab's label: in copy mode, with builds waiting, their count in gold.
function BuildDialog.pasteTabLabel(mode, count)
	if mode == "copy" and (count or 0) > 0 then
		return string.format(L.dialogTabPasteCount, L.dialogTabPaste, Theme.inlineColor("gold", tostring(count)))
	end
	return L.dialogTabPaste
end

function BuildDialog.modeOf(mode)
	return mode == "copy" and "copy" or "paste"
end

function BuildDialog.nextMode(mode)
	return BuildDialog.modeOf(mode) == "paste" and "copy" or "paste"
end

--- Whether a rect { left, right, top, bottom } lies inside `parent`'s, within
--- SCREEN_TOLERANCE; an edge either side cannot say is not held against it.
function BuildDialog.onScreen(rect, parent)
	local slack = BuildDialog.SCREEN_TOLERANCE
	local function known(edge)
		return type(rect[edge]) == "number" and type(parent[edge]) == "number"
	end
	if known("left") and rect.left < parent.left - slack then
		return false
	end
	if known("right") and rect.right > parent.right + slack then
		return false
	end
	if known("top") and rect.top > parent.top + slack then
		return false
	end
	if known("bottom") and rect.bottom < parent.bottom - slack then
		return false
	end
	return true
end

-- The live inputs ------------------------------------------------------------------------

local function dataOf()
	return BuildDialog.data or ns.Data
end

--- The wall clock, seconds since the epoch; nil on a client without one.
local function wallNow()
	if type(time) == "function" then
		return time()
	end
	return nil
end

--- The client's own clock: local fields from date("*t"), seconds from time(fields), so
--- daylight saving is the client's and never a fixed offset.
local function liveClock()
	return {
		now = wallNow() or 0,
		fields = function(seconds)
			local fields = type(date) == "function" and date("*t", seconds) or nil
			if type(fields) ~= "table" then
				return BuildDialog.utcFields(seconds)
			end
			return fields
		end,
		epoch = function(fields)
			local seconds = type(time) == "function" and time(fields) or nil
			if type(seconds) ~= "number" then
				return BuildDialog.civilSeconds(fields.year, fields.month, fields.day, fields.hour, fields.min,
					fields.sec or 0)
			end
			return seconds
		end,
	}
end

--- The character logged in, for Codec.loadBuild's class check; nil where the client cannot say.
function BuildDialog.player()
	local ok, slug = pcall(Talents.playerClassSlug)
	if not ok or type(slug) ~= "string" then
		return nil
	end
	return { classSlug = slug, classNames = Compat.classNames() }
end

local function liveLevel()
	local ok, level = pcall(Compat.playerLevel)
	return ok and level or nil
end

local function liveRanks(data)
	local ok, ranks = pcall(Talents.readStagedRanks, data)
	if ok and type(ranks) == "table" then
		return ranks
	end
	ok, ranks = pcall(Talents.readRanks, data)
	return ok and ranks or {}
end

--- The next talent the build wants on this character now; nil where the client cannot say.
local function liveNext(build)
	local data = dataOf()
	local ok, label = pcall(BuildDialog.nextTalent, build, liveRanks(data), data, liveLevel())
	return ok and label or nil
end

local function currentAck()
	return Follow.messages(ForeverSixtyInbox, Export.characterKey(), "ack")[1]
end

local function currentRows()
	return BuildDialog.inboxRows(Follow.inboxBuilds(ForeverSixtyInbox, Export.characterKey()),
		Follow.dismissedSet(), dataOf(), BuildDialog.player())
end

local SLOT_WORDS = setmetatable({}, {
	__index = function(_, name)
		return Compat.globalString(name)
	end,
})

-- Using a build --------------------------------------------------------------------------

--- Follow `code` (Follow.load saves it in the active slot, refusing another class's), mark a
--- build from the site used by its id and its code, redraw every surface, close, and say so
--- in one chat line with the next talent.
function BuildDialog.use(code, name, inboxId)
	local data = dataOf()
	local build = Follow.load(code, data, name, nil, BuildDialog.player())
	if build == nil then
		return nil
	end
	if inboxId ~= nil then
		Follow.dismissInbox(inboxId)
		Follow.dismissInbox(code)
	end
	if BuildDialog.afterUse ~= nil then
		BuildDialog.afterUse()
	end
	BuildDialog.close()
	local line = BuildDialog.chatLine(Follow.buildName(data, build), liveNext(build))
	if BuildDialog.say ~= nil then
		BuildDialog.say(line)
	else
		Theme.say(line)
	end
	return build
end

--- The player typed or pasted: judge the text and redraw. Nothing is loaded here: the
--- build the player follows, the chip and the strip stay as they were until Use.
function BuildDialog.onPaste(text)
	BuildDialog.verdict = BuildDialog.evaluate(text, dataOf(), BuildDialog.player())
	if BuildDialog.verdict.ok then
		BuildDialog.editing = false
	end
	return BuildDialog.drawPaste()
end

--- A "from the site" row was clicked: preview it in the same place a paste previews.
function BuildDialog.choose(row)
	local verdict = BuildDialog.evaluate(row.code, dataOf(), BuildDialog.player())
	verdict.name, verdict.inboxId = row.name, row.id
	if verdict.build ~= nil then
		-- The preview names it as the site does (Follow.buildName reads the build's own name).
		verdict.build.name = row.name
	end
	BuildDialog.verdict = verdict
	BuildDialog.editing = false
	if BuildDialog.paste ~= nil then
		BuildDialog.paste.box:SetText(row.code)
		BuildDialog.selectPaste()
	end
	return BuildDialog.drawPaste()
end

--- The Use button, and Enter in the paste box: only while a preview shows. Loads the
--- verdict's own trimmed code, never the box's text.
function BuildDialog.useCurrent()
	local verdict = BuildDialog.verdict
	if verdict == nil or not verdict.ok then
		return nil
	end
	local name = verdict.name or Follow.buildName(dataOf(), verdict.build)
	return BuildDialog.use(verdict.code, name, verdict.inboxId)
end

-- The frame ------------------------------------------------------------------------------

local function contentWidth()
	return Theme.SIZES.dialogWidth - Theme.SIZES.padding * 2
end

--- A font string at the dialog's text size, coloured `hexKey`.
local function text(parent, hexKey, width)
	local region = Theme.fontString(parent, "OVERLAY", "normal")
	local path, _, flags = region:GetFont()
	if path ~= nil then
		region:SetFont(path, Theme.SIZES.dialogText, flags)
	end
	region:SetTextColor(Theme.rgb(Theme.HEX[hexKey or "body"]))
	region:SetJustifyH("LEFT")
	if width ~= nil then
		region:SetWidth(width)
	end
	return region
end

--- An existing label or box set to the dialog's text size.
local function enlarge(region)
	if region == nil or type(region.GetFont) ~= "function" then
		return region
	end
	local path, _, flags = region:GetFont()
	if path ~= nil then
		region:SetFont(path, Theme.SIZES.dialogText, flags)
	end
	return region
end

local function setMaxLines(region, lines)
	if type(region.SetMaxLines) == "function" then
		region:SetMaxLines(lines)
	end
end

--- Run `action` on `script` after whatever the frame already runs there (Widgets' own
--- focus release, say): HookScript where the client has it is not relied on.
local function chain(frame, script, action)
	local previous = frame:GetScript(script)
	frame:SetScript(script, function(...)
		if previous ~= nil then
			previous(...)
		end
		action(...)
	end)
end

local function paintEdges(edges, hexKey)
	for _, edge in ipairs(edges or {}) do
		Theme.paint(edge, hexKey)
	end
end

--- One line's height at the dialog's text size, measured once on the client.
local function lineHeight()
	if BuildDialog.lineHeightPx == nil then
		local probe = text(BuildDialog.frame, "body")
		probe:SetText("Ag")
		local height = type(probe.GetStringHeight) == "function" and probe:GetStringHeight() or nil
		probe:Hide()
		BuildDialog.lineHeightPx = (type(height) == "number" and height > 0) and math.ceil(height)
			or Theme.SIZES.dialogText + 2
	end
	return BuildDialog.lineHeightPx
end

--- Where the dialog sits: the saved spot after a drag, else the middle of the screen.
function BuildDialog.place()
	local frame = BuildDialog.frame
	if frame == nil then
		return nil
	end
	frame:ClearAllPoints()
	local point = Prefs.get("dialog", "point")
	if type(point) == "string" and point ~= "" then
		frame:SetPoint(point, UIParent, point, Prefs.get("dialog", "x"), Prefs.get("dialog", "y"))
		return point
	end
	frame:SetPoint("CENTER", UIParent, "CENTER", 0, 0)
	return "CENTER"
end

-- Reset positions moves the dialog at once, open or not.
Prefs.onReset[#Prefs.onReset + 1] = function()
	BuildDialog.place()
end

--- After a drag by the title: the dialog's centre against the screen's.
function BuildDialog.savePosition()
	local frame = BuildDialog.frame
	local x, y = frame:GetCenter()
	local parentX, parentY = UIParent:GetCenter()
	if x == nil or y == nil or parentX == nil or parentY == nil then
		return nil
	end
	Prefs.set("dialog", "point", "CENTER")
	Prefs.set("dialog", "x", x - parentX)
	Prefs.set("dialog", "y", y - parentY)
	BuildDialog.place()
	return true
end

--- Escape, Enter and Tab in either box: Escape closes the dialog (the box gives up the
--- keyboard as it hides, so the player can walk), Tab switches tabs, Enter does the tab's job.
local function bindKeys(box, onEnter)
	box:SetScript("OnEscapePressed", function()
		BuildDialog.close()
	end)
	box:SetScript("OnTabPressed", function()
		BuildDialog.setMode(BuildDialog.nextMode(BuildDialog.mode))
	end)
	box:SetScript("OnEnterPressed", onEnter)
end

--- A box set in a field: the field (the border, what a caller anchors), a clip of whole
--- lines inside it (so no glyph tops of a next line show), the box in the clip.
local function boxInField(parent, readOnly)
	local S = Theme.SIZES
	local width = contentWidth()
	local box = Widgets.editBox(parent, width, S.dialogLine, readOnly, false)
	enlarge(box)
	local field = Widgets.field(box)
	local clip = Theme.createFrame("Frame", nil, field)
	clip:SetPoint("TOPLEFT", field, "TOPLEFT", S.gap * 2, -S.gap)
	clip:SetWidth(width - S.gap * 4)
	if type(clip.SetClipsChildren) == "function" then
		clip:SetClipsChildren(true)
	end
	box:SetParent(clip)
	box:ClearAllPoints()
	box:SetPoint("TOPLEFT", clip, "TOPLEFT", 0, 0)
	box:SetWidth(width - S.gap * 4)
	return box, field, clip
end

--- Size a field to `lines` whole lines of text.
local function sizeField(field, clip, box, lines)
	local S = Theme.SIZES
	local height = lines * lineHeight()
	clip:SetHeight(height)
	box:SetHeight(height)
	field:SetHeight(height + S.gap * 2)
	return height + S.gap * 2
end

local function buildTitle(frame)
	local S = Theme.SIZES
	local bar = Theme.createFrame("Frame", nil, frame)
	bar:SetPoint("TOPLEFT", frame, "TOPLEFT", 0, 0)
	bar:SetPoint("TOPRIGHT", frame, "TOPRIGHT", 0, 0)
	bar:SetHeight(S.titleBarHeight)
	bar:EnableMouse(true)
	bar:RegisterForDrag("LeftButton")
	bar:SetScript("OnDragStart", function()
		frame:StartMoving()
	end)
	bar:SetScript("OnDragStop", function()
		frame:StopMovingOrSizing()
		BuildDialog.savePosition()
	end)
	local title = text(bar, "gold")
	title:SetText(L.addonName)
	title:SetPoint("LEFT", bar, "LEFT", S.padding, 0)
	-- The close X: a 20 px hit area that lights up under the cursor.
	local close = Widgets.closeButton(bar, function()
		BuildDialog.close()
	end)
	close:SetSize(S.dialogClose, S.dialogClose)
	close:SetPoint("RIGHT", bar, "RIGHT", -S.gap * 2, 0)
	enlarge(close.foreverSixtyLabel)
	local hover = Theme.texture(close, "BACKGROUND", "hover")
	hover:SetAllPoints(close)
	hover:Hide()
	chain(close, "OnEnter", function()
		hover:Show()
	end)
	chain(close, "OnLeave", function()
		hover:Hide()
	end)
	return bar
end

local function buildTab(frame, label, mode)
	local S = Theme.SIZES
	local tab = Widgets.tab(frame, label, function()
		BuildDialog.setMode(mode)
	end)
	enlarge(tab.foreverSixtyLabel)
	local width = type(tab.foreverSixtyLabel.GetStringWidth) == "function" and tab.foreverSixtyLabel:GetStringWidth() or 0
	tab:SetWidth(math.max(S.tabWidth, width + S.padding * 2))
	return tab
end

--- One planned item in the preview: its name in its quality colour and its slot word,
--- muted, at the right; hovering shows the item.
local function buildPickRow(parent, width)
	local S = Theme.SIZES
	local row = Theme.createFrame("Frame", nil, parent)
	row:SetSize(width, S.dialogGearRow)
	row:EnableMouse(true)
	local icon = Theme.icon(row, "ARTWORK", nil, S.iconSize)
	icon:SetPoint("LEFT", row, "LEFT", 0, 0)
	local slot = text(row, "muted")
	slot:SetJustifyH("RIGHT")
	slot:SetPoint("RIGHT", row, "RIGHT", 0, 0)
	local name = text(row, "body", width - S.iconSize - S.gap * 2 - 80)
	name:SetPoint("LEFT", icon, "RIGHT", S.gap * 2, 0)
	if type(name.SetWordWrap) == "function" then
		name:SetWordWrap(false)
	end
	local pick = { frame = row, icon = icon, name = name, slot = slot }
	Widgets.attachTooltip(row, function()
		return pick.itemId, nil
	end)
	return pick
end

local function buildPreview(content)
	local S = Theme.SIZES
	local width = contentWidth()
	local inner = width - S.dialogClassIcon - S.gap * 2
	local preview = Theme.createFrame("Frame", nil, content)
	preview:SetWidth(width)
	-- The chip's own round class icon (Theme.classIcon: the client's circles, cut by coords).
	local icon = preview:CreateTexture(nil, "ARTWORK")
	icon:SetSize(S.dialogClassIcon, S.dialogClassIcon)
	icon:SetPoint("TOPLEFT", preview, "TOPLEFT", 0, 0)
	local name = text(preview, "gold", inner)
	name:SetPoint("TOPLEFT", icon, "TOPRIGHT", S.gap * 2, 0)
	local detail = text(preview, "muted", inner)
	detail:SetPoint("TOPLEFT", name, "BOTTOMLEFT", 0, -S.gap / 2)
	setMaxLines(detail, 2)
	local headingFrame = Theme.createFrame("Frame", nil, preview)
	headingFrame:SetSize(inner, S.dialogGearRow)
	headingFrame:EnableMouse(true)
	local heading = text(headingFrame, "muted", inner)
	heading:SetPoint("LEFT", headingFrame, "LEFT", 0, 0)
	headingFrame:SetScript("OnEnter", function(self)
		Theme.showLines(self, { L.dialogGearHeadingTip })
	end)
	headingFrame:SetScript("OnLeave", function()
		Theme.hideTooltip()
	end)
	local picks = {}
	for index = 1, BuildDialog.GEAR_PICKS do
		picks[index] = buildPickRow(preview, inner)
	end
	preview:Hide()
	return {
		frame = preview, icon = icon, name = name, detail = detail,
		headingFrame = headingFrame, heading = heading, picks = picks,
	}
end

local function buildRow(content)
	local S = Theme.SIZES
	local row = { frame = Theme.createFrame("Button", nil, content) }
	row.frame:SetSize(contentWidth(), S.rowHeight)
	local highlight = Theme.texture(row.frame, "BACKGROUND", "hover")
	highlight:SetAllPoints(row.frame)
	highlight:Hide()
	row.highlight = highlight
	row.label = text(row.frame, "body", contentWidth() - S.gap * 2)
	row.label:SetPoint("LEFT", row.frame, "LEFT", S.gap, 0)
	if type(row.label.SetWordWrap) == "function" then
		row.label:SetWordWrap(false)
	end
	row.frame:SetScript("OnEnter", function()
		highlight:Show()
	end)
	row.frame:SetScript("OnLeave", function()
		if not row.selected then
			highlight:Hide()
		end
	end)
	row.frame:SetScript("OnClick", function()
		if row.entry ~= nil then
			BuildDialog.choose(row.entry)
		end
	end)
	row.frame:Hide()
	return row
end

local function buildPaste(frame, top)
	local S = Theme.SIZES
	local width = contentWidth()
	local content = Theme.createFrame("Frame", nil, frame)
	content:SetPoint("TOPLEFT", frame, "TOPLEFT", S.padding, -top)
	content:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", -S.padding, S.padding)
	local paste = { content = content, rows = {} }
	paste.following = text(content, "body", width)
	if type(paste.following.SetWordWrap) == "function" then
		paste.following:SetWordWrap(false)
	end
	for index = 1, S.dialogInboxRows do
		paste.rows[index] = buildRow(content)
	end
	paste.more = text(content, "muted", width)
	local box, field, clip = boxInField(content, false)
	-- Ctrl+C here is not a copy out of the dialog; the box keeps the keyboard until Escape.
	box:SetScript("OnKeyDown", nil)
	paste.placeholder = text(clip, "muted", width - S.gap * 4)
	paste.placeholder:SetText(L.dialogPaste)
	paste.placeholder:SetPoint("TOPLEFT", clip, "TOPLEFT", 0, 0)
	-- Once a code loads, the box reads as one dim line with an ellipsis; a click opens it.
	paste.display = text(field, "muted", width - S.gap * 4)
	paste.display:SetPoint("LEFT", field, "LEFT", S.gap * 2, 0)
	if type(paste.display.SetWordWrap) == "function" then
		paste.display:SetWordWrap(false)
	end
	box:SetScript("OnTextChanged", function(self, userInput)
		if userInput then
			BuildDialog.onPaste(self:GetText())
		else
			BuildDialog.drawPaste()
		end
	end)
	box:SetScript("OnMouseDown", function()
		if not BuildDialog.editing and BuildDialog.verdict ~= nil then
			BuildDialog.editing = true
			BuildDialog.drawPaste()
		end
	end)
	box:SetScript("OnEditFocusGained", function()
		BuildDialog.focused = true
		BuildDialog.drawPaste()
	end)
	box:SetScript("OnEditFocusLost", function()
		BuildDialog.focused = false
		if BuildDialog.editing then
			-- Done editing: a refusal shows again.
			BuildDialog.editing = false
		end
		BuildDialog.drawPaste()
	end)
	bindKeys(box, function()
		if BuildDialog.editing and not (BuildDialog.verdict or {}).ok then
			BuildDialog.editing = false
			BuildDialog.drawPaste()
			return
		end
		BuildDialog.useCurrent()
	end)
	paste.box, paste.field, paste.clip = box, field, clip
	paste.lead = text(content, "warning", width)
	setMaxLines(paste.lead, 2)
	paste.body = text(content, "body", width)
	setMaxLines(paste.body, 2)
	paste.still = text(content, "muted", width)
	paste.next = text(content, "muted", width)
	paste.preview = buildPreview(content)
	paste.use = Cards.primaryButton(content, L.dialogUse, function()
		BuildDialog.useCurrent()
	end)
	enlarge(paste.use.foreverSixtyLabel)
	paste.use:SetPoint("BOTTOMLEFT", content, "BOTTOMLEFT", 0, 0)
	paste.use:Hide()
	return paste
end

local function buildCopy(frame, top)
	local S = Theme.SIZES
	local width = contentWidth()
	local content = Theme.createFrame("Frame", nil, frame)
	content:SetPoint("TOPLEFT", frame, "TOPLEFT", S.padding, -top)
	content:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", -S.padding, S.padding)
	local copy = { content = content }
	copy.lead = text(content, "body", width)
	copy.lead:SetPoint("TOPLEFT", content, "TOPLEFT", 0, 0)
	setMaxLines(copy.lead, 2)
	local box, field, clip = boxInField(content, true)
	copy.box, copy.field, copy.clip = box, field, clip
	bindKeys(box, function()
		BuildDialog.copy()
	end)
	-- The one copy claim the dialog makes comes from a real Ctrl+C in this box.
	chain(box, "OnKeyDown", function(self, key)
		local control = type(IsControlKeyDown) == "function" and IsControlKeyDown() == true
		local focused = type(self.HasFocus) == "function" and self:HasFocus() == true
		if control and key == "A" then
			BuildDialog.selected = true
		elseif BuildDialog.copyKey(key, control, focused, BuildDialog.selected == true) then
			BuildDialog.copiedAt = wallNow()
			BuildDialog.drawCopy(false)
		end
	end)
	chain(box, "OnMouseDown", function()
		-- A click places the cursor: the export is no longer all selected.
		BuildDialog.selected = false
	end)
	chain(box, "OnEditFocusGained", function()
		paintEdges(field.foreverSixtyBorder, "gold")
	end)
	chain(box, "OnEditFocusLost", function()
		BuildDialog.selected = false
		paintEdges(field.foreverSixtyBorder, "border")
	end)
	-- The whole export, wrapped; the wheel scrolls it inside its lines.
	clip:EnableMouseWheel(true)
	clip:SetScript("OnMouseWheel", function(_, delta)
		BuildDialog.scrollCopy(-delta)
	end)
	copy.measure = text(content, "body", width - S.gap * 4)
	copy.measure:Hide()
	copy.button = Widgets.button(content, L.dialogCopy, function()
		BuildDialog.copy()
	end)
	enlarge(copy.button.foreverSixtyLabel)
	copy.seen = text(content, "muted", width)
	setMaxLines(copy.seen, 2)
	copy.sim = text(content, "muted", width)
	copy.companion = {}
	for index = 1, 3 do
		copy.companion[index] = text(content, "muted", width)
	end
	content:Hide()
	return copy
end

local function onEvent(_, event, itemId)
	if event == "GET_ITEM_INFO_RECEIVED" and itemId ~= nil and BuildDialog.pending[itemId] then
		BuildDialog.pending[itemId] = nil
		if BuildDialog.frame:IsShown() then
			BuildDialog.drawPaste()
		end
	end
end

--- The dialog, built on the first open: a 420x300 panel titled "Forever Sixty",
--- draggable by its title, closed by Escape, with the two tabs and the hint under the title.
function BuildDialog.ensure()
	if BuildDialog.frame ~= nil then
		return BuildDialog.frame
	end
	local S = Theme.SIZES
	local frame = Widgets.panel(UIParent, S.dialogWidth, S.dialogHeight, BuildDialog.FRAME_NAME)
	frame:Hide()
	frame:SetFrameStrata("DIALOG")
	frame:SetClampedToScreen(true)
	frame:SetMovable(true)
	frame:EnableMouse(true)
	BuildDialog.frame = frame
	buildTitle(frame)
	local tabs = {
		paste = buildTab(frame, L.dialogTabPaste, "paste"),
		copy = buildTab(frame, L.dialogCopy, "copy"),
	}
	tabs.paste:SetPoint("TOPLEFT", frame, "TOPLEFT", S.padding, -S.titleBarHeight)
	tabs.copy:SetPoint("LEFT", tabs.paste, "RIGHT", S.gap * 2, 0)
	-- The hint sits at the right of the tab strip, so the panel keeps its 300 px.
	local hint = text(frame, "muted")
	hint:SetText(L.dialogTabHint)
	hint:SetPoint("RIGHT", frame, "TOPRIGHT", -S.padding, -(S.titleBarHeight + S.tabHeight / 2))
	BuildDialog.tabs = tabs
	local top = S.titleBarHeight + S.tabHeight + S.gap * 2
	BuildDialog.paste = buildPaste(frame, top)
	BuildDialog.copyTab = buildCopy(frame, top)
	frame:SetScript("OnHide", function()
		-- Never keep the keyboard once the dialog is gone; a copy claim ends with it.
		BuildDialog.paste.box:ClearFocus()
		BuildDialog.copyTab.box:ClearFocus()
		BuildDialog.copiedAt = nil
	end)
	frame:SetScript("OnEvent", onEvent)
	Theme.registerEvent(frame, "GET_ITEM_INFO_RECEIVED")
	Theme.makeEscapable(BuildDialog.FRAME_NAME)
	return frame
end

-- Paste mode -----------------------------------------------------------------------------

local function stack(region, content, y, height)
	region:ClearAllPoints()
	region:SetPoint("TOPLEFT", content, "TOPLEFT", 0, -y)
	region:Show()
	return y + height
end

local function textHeight(region)
	local height = type(region.GetStringHeight) == "function" and region:GetStringHeight() or nil
	return (type(height) == "number" and height > 0) and height or lineHeight()
end

local function drawPreview(build, y)
	local S = Theme.SIZES
	local paste = BuildDialog.paste
	local preview = paste.preview
	local data = dataOf()
	local model = BuildDialog.preview(build, data, {
		level = liveLevel(),
		unspent = Progress.unspentPoints(liveLevel(), liveRanks(data)),
		slotWord = function(slot)
			return BuildDialog.slotWord(slot, SLOT_WORDS)
		end,
	})
	local path, coords = Theme.classIcon(model.classToken)
	preview.icon:SetTexture(path)
	if type(preview.icon.SetTexCoord) == "function" then
		preview.icon:SetTexCoord(coords[1], coords[2], coords[3], coords[4])
	end
	preview.name:SetText(model.name)
	preview.detail:SetText(model.detail)
	local headY = math.max(S.dialogClassIcon, textHeight(preview.name) + S.gap / 2 + textHeight(preview.detail))
		+ S.gap
	if model.heading ~= nil then
		preview.heading:SetText(model.heading)
		preview.headingFrame:ClearAllPoints()
		preview.headingFrame:SetPoint("TOPLEFT", preview.frame, "TOPLEFT", S.dialogClassIcon + S.gap * 2, -headY)
		preview.headingFrame:Show()
		headY = headY + S.dialogGearRow
	else
		preview.headingFrame:Hide()
	end
	for index, pick in ipairs(preview.picks) do
		local entry = model.picks[index]
		pick.itemId = entry ~= nil and entry.itemId or nil
		if entry == nil then
			pick.frame:Hide()
		else
			pick.icon:SetTexture(Compat.itemIcon(entry.itemId) or Theme.UNKNOWN_ICON)
			local name = Compat.itemName(entry.itemId)
			if name == nil then
				BuildDialog.pending[entry.itemId] = true
			end
			pick.name:SetText(name or "")
			pick.name:SetTextColor(Theme.qualityColor(Compat.itemQuality(entry.itemId)))
			pick.slot:SetText(entry.slotWord)
			pick.frame:ClearAllPoints()
			pick.frame:SetPoint("TOPLEFT", preview.frame, "TOPLEFT", S.dialogClassIcon + S.gap * 2, -headY)
			pick.frame:Show()
			headY = headY + S.dialogGearRow
		end
	end
	preview.frame:SetHeight(headY)
	stack(preview.frame, paste.content, y, headY)
end

--- Redraw paste mode from the verdict: the following and next-talent lines in the empty
--- dialog, the rows from the site, the box (one dim line once a code loads), the refusal's
--- three lines or the preview, and the gold Use button.
function BuildDialog.drawPaste()
	local paste = BuildDialog.paste
	if paste == nil then
		return nil
	end
	local S = Theme.SIZES
	local data = dataOf()
	local verdict = BuildDialog.verdict
	local view = BuildDialog.pasteView(verdict, BuildDialog.editing == true, BuildDialog.focused == true)
	local content = paste.content
	for _, region in ipairs({ paste.following, paste.more, paste.lead, paste.body, paste.still, paste.next }) do
		region:Hide()
	end
	paste.preview.frame:Hide()
	paste.use:Hide()
	local y = 0
	if verdict == nil then
		local line = BuildDialog.followingLine(Follow.build, data)
		if line ~= nil then
			paste.following:SetText(line)
			y = stack(paste.following, content, y, textHeight(paste.following) + S.gap * 2)
		end
	end
	local rows = BuildDialog.rowsView(BuildDialog.rows or {}, S.dialogInboxRows)
	for index, row in ipairs(paste.rows) do
		local entry = rows.shown[index]
		row.entry = entry
		if entry == nil then
			row.frame:Hide()
		else
			row.label:SetText(entry.label)
			row.selected = verdict ~= nil and verdict.inboxId == entry.id
			if row.selected then
				row.highlight:Show()
			else
				row.highlight:Hide()
			end
			y = stack(row.frame, content, y, S.rowHeight)
		end
	end
	if rows.more ~= nil and verdict == nil then
		paste.more:SetText(rows.more)
		y = stack(paste.more, content, y, textHeight(paste.more))
	end
	if #rows.shown > 0 then
		y = y + S.gap
	end
	-- The box: whole lines, one dim line with an ellipsis once a code loads.
	local boxText = paste.box:GetText() or ""
	local lines = view.collapsed and 1 or S.dialogBoxLines
	local fieldHeight = sizeField(paste.field, paste.clip, paste.box, lines)
	y = stack(paste.field, content, y, fieldHeight) + S.gap * 2
	if view.collapsed then
		paste.box:SetAlpha(0)
		local measure = function(value)
			paste.display:SetText(value)
			return paste.display:GetStringWidth()
		end
		local shown = verdict.code or boxText:match("^%s*(.-)%s*$"):gsub("%s+", " ")
		paste.display:SetText(Chip.fit(shown, measure, contentWidth() - S.gap * 4))
		paste.display:Show()
		paste.placeholder:Hide()
	else
		paste.box:SetAlpha(1)
		paste.display:Hide()
		if boxText == "" then
			paste.placeholder:Show()
		else
			paste.placeholder:Hide()
		end
	end
	paintEdges(paste.field.foreverSixtyBorder, view.border)
	if view.reason then
		local followed = Follow.build ~= nil and Follow.buildName(data, Follow.build) or nil
		local refusal = BuildDialog.refusal(verdict.reason, followed)
		paste.lead:SetText(refusal.lead)
		y = stack(paste.lead, content, y, textHeight(paste.lead) + S.gap)
		if refusal.body ~= nil then
			paste.body:SetText(refusal.body)
			y = stack(paste.body, content, y, textHeight(paste.body) + S.gap)
		end
		if refusal.following ~= nil then
			paste.still:SetText(refusal.following)
			stack(paste.still, content, y, textHeight(paste.still))
		end
	elseif view.preview then
		drawPreview(verdict.build, y)
		paste.use:Show()
	elseif verdict == nil then
		local line = BuildDialog.nextTalentLine(liveNext(Follow.build))
		if line ~= nil then
			paste.next:SetText(line)
			stack(paste.next, content, y, textHeight(paste.next))
		end
	end
	return verdict
end

--- Select all of the paste box (a second paste replaces the code rather than appending).
function BuildDialog.selectPaste()
	local box = BuildDialog.paste.box
	box:HighlightText()
	box:SetFocus()
end

-- Copy mode ------------------------------------------------------------------------------

local function outboxName()
	local ok, name = pcall(Export.outboxCharacter)
	return ok and name or nil
end

--- Scroll the export by `lines` inside its field.
function BuildDialog.scrollCopy(lines)
	local copy = BuildDialog.copyTab
	if copy == nil then
		return 0
	end
	local visible = Theme.SIZES.dialogCopyLines * lineHeight()
	local total = copy.boxHeight or visible
	local range = math.max(0, total - visible)
	BuildDialog.copyOffset = math.max(0, math.min(range, (BuildDialog.copyOffset or 0) + lines * lineHeight()))
	copy.box:ClearAllPoints()
	copy.box:SetPoint("TOPLEFT", copy.clip, "TOPLEFT", 0, BuildDialog.copyOffset)
	return BuildDialog.copyOffset
end

--- Fill copy mode from copyModel: the lead line, the whole export wrapped in its box (and
--- selected, when `select`), the secondary button, the ack's lines and the companion's.
function BuildDialog.drawCopy(select)
	local copy = BuildDialog.copyTab
	local S = Theme.SIZES
	-- Guarded: a client answer the export cannot encode must never leave a half-drawn tab.
	local ran, code, reason = pcall(Export.string, dataOf())
	if not ran then
		code, reason = nil, nil
	end
	BuildDialog.exportCode = code
	local name = outboxName()
	local model = BuildDialog.copyModel({
		code = code,
		reason = reason,
		copiedAt = BuildDialog.copiedAt,
		ack = currentAck(),
		savedAt = BuildDialog.savedAtFor(ForeverSixtyDB, name),
		simRequestedAt = BuildDialog.simRequestedAtFor(ForeverSixtyDB, name),
		clock = liveClock(),
	})
	local content = copy.content
	copy.lead:SetText(model.lead.text)
	copy.lead:SetTextColor(Theme.rgb(Theme.HEX[model.lead.hexKey]))
	local y = textHeight(copy.lead) + S.gap * 2
	-- The box holds the whole export; its field shows dialogCopyLines lines of it.
	copy.measure:SetText(model.text)
	copy.boxHeight = math.max(lineHeight(), math.ceil(textHeight(copy.measure)))
	local fieldHeight = sizeField(copy.field, copy.clip, copy.box, S.dialogCopyLines)
	copy.box:SetHeight(math.max(copy.boxHeight, S.dialogCopyLines * lineHeight()))
	y = stack(copy.field, content, y, fieldHeight) + S.gap * 2
	if select ~= false then
		BuildDialog.copyOffset = 0
		BuildDialog.scrollCopy(0)
		if model.selected then
			Widgets.selectText(copy.box, model.text)
			BuildDialog.selected = true
		else
			Widgets.setText(copy.box, model.text)
		end
	end
	Widgets.setEnabled(copy.button, model.button)
	y = stack(copy.button, content, y, S.buttonHeight) + S.gap * 2
	for _, region in ipairs({ copy.seen, copy.sim }) do
		region:Hide()
	end
	if model.seen ~= nil then
		copy.seen:SetText(model.seen)
		y = stack(copy.seen, content, y, textHeight(copy.seen) + S.gap)
	end
	if model.sim ~= nil then
		copy.sim:SetText(model.sim)
		stack(copy.sim, content, y, textHeight(copy.sim))
	end
	-- The companion's lines, one sentence each, from the bottom up.
	local bottom = 0
	for index = #copy.companion, 1, -1 do
		local region = copy.companion[index]
		local line = model.companion ~= nil and model.companion[index] or nil
		if line == nil then
			region:Hide()
		else
			region:SetText(line)
			region:ClearAllPoints()
			region:SetPoint("BOTTOMLEFT", content, "BOTTOMLEFT", 0, bottom)
			region:Show()
			bottom = bottom + textHeight(region)
		end
	end
	return model
end

--- "Copy my character": select the export again for Ctrl+C. It claims nothing itself.
function BuildDialog.copy()
	local copy = BuildDialog.copyTab
	if BuildDialog.exportCode == nil or copy == nil then
		return nil
	end
	Widgets.selectText(copy.box, BuildDialog.exportCode)
	BuildDialog.selected = true
	return BuildDialog.exportCode
end

-- Opening and closing --------------------------------------------------------------------

local function refreshTabs()
	local count = #(BuildDialog.rows or {})
	local label = BuildDialog.pasteTabLabel(BuildDialog.mode, count)
	BuildDialog.tabs.paste.foreverSixtyLabel:SetText(label)
end

--- Show one tab: the paste box takes the keyboard in paste mode; in copy mode the export
--- is selected in its box.
function BuildDialog.setMode(mode)
	mode = BuildDialog.modeOf(mode)
	BuildDialog.mode = mode
	if BuildDialog.frame == nil then
		return mode
	end
	local paste, copy = BuildDialog.paste, BuildDialog.copyTab
	Widgets.setTabActive(BuildDialog.tabs.paste, mode == "paste")
	Widgets.setTabActive(BuildDialog.tabs.copy, mode == "copy")
	refreshTabs()
	if mode == "paste" then
		copy.box:ClearFocus()
		copy.content:Hide()
		paste.content:Show()
		BuildDialog.drawPaste()
		paste.box:SetFocus()
	else
		paste.box:ClearFocus()
		paste.content:Hide()
		copy.content:Show()
		-- A copy claim stays until the dialog closes (OnHide clears it), across tab switches.
		BuildDialog.drawCopy(true)
	end
	return mode
end

--- Open the dialog on `mode` ("paste" by default, or "copy"). The paste box starts empty.
--- One frame later the tab's box takes the keyboard again (the chat box that ran /fs lets
--- go of it after the command; copy mode selects the export again) and the render check
--- runs. False when the dialog turned itself off earlier (the fallback ran instead).
function BuildDialog.open(mode)
	if BuildDialog.disabled then
		if BuildDialog.fallback ~= nil then
			BuildDialog.fallback()
		end
		return false
	end
	local frame = BuildDialog.ensure()
	BuildDialog.place()
	BuildDialog.verdict, BuildDialog.editing = nil, false
	BuildDialog.rows = currentRows()
	BuildDialog.paste.box:SetText("")
	frame:Show()
	BuildDialog.setMode(mode)
	local deferred = Theme.after(0, function()
		if not frame:IsShown() then
			return
		end
		if BuildDialog.mode == "copy" then
			BuildDialog.copy()
		else
			BuildDialog.paste.box:SetFocus()
		end
		BuildDialog.selfCheck()
	end)
	if not deferred then
		BuildDialog.selfCheck()
	end
	return true
end

--- Hide the dialog and give the keyboard back (OnHide does it too, for the Escape the
--- client handles itself through UISpecialFrames).
function BuildDialog.close()
	if BuildDialog.frame ~= nil then
		BuildDialog.paste.box:ClearFocus()
		BuildDialog.copyTab.box:ClearFocus()
		BuildDialog.frame:Hide()
	end
end

--- The keybind: open on the paste tab, or close.
function BuildDialog.toggle()
	if BuildDialog.frame ~= nil and BuildDialog.frame:IsShown() then
		BuildDialog.close()
		return false
	end
	return BuildDialog.open("paste")
end

local function rectOf(region)
	return { left = region:GetLeft(), right = region:GetRight(), top = region:GetTop(), bottom = region:GetBottom() }
end

--- The render check (spec section 2.4), only while the dialog is visible. Off screen (by
--- more than a pixel): it turns itself off for the session, says why in /fs diag and the
--- fallback (the command list) runs. In paste mode without the keyboard: a /fs diag note
--- only. True when it passes.
function BuildDialog.selfCheck()
	local frame = BuildDialog.frame
	if frame == nil or type(frame.IsVisible) ~= "function" or not frame:IsVisible() then
		return true
	end
	if not BuildDialog.onScreen(rectOf(frame), rectOf(UIParent)) then
		BuildDialog.disabled = true
		frame:Hide()
		Theme.note(L.diagDialogOff)
		if BuildDialog.fallback ~= nil then
			BuildDialog.fallback()
		end
		return false
	end
	local box = BuildDialog.paste.box
	if BuildDialog.mode == "paste" and type(box.HasFocus) == "function" and box:HasFocus() ~= true then
		Theme.note(L.diagDialogFocus)
		return false
	end
	return true
end

ns.BuildDialog = BuildDialog
return BuildDialog
