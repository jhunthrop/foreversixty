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
--- The paste box's height in lines while empty or edited (one once a code loads).
BuildDialog.EMPTY_LINES = 2
--- Where the "Tab to switch" hint sits: the panel's bottom left.
BuildDialog.HINT_POINT = "BOTTOMLEFT"
--- The collapsed code's colour: muted text on the box's inset ground.
BuildDialog.COLLAPSED_TEXT = "muted"
--- The gear rows' pitch (Theme.SIZES.rowHeight), so their icons never touch.
BuildDialog.GEAR_PITCH = 20
--- A zero-width space: a break point the export's display may wrap at where the client cannot
--- break a run of characters itself. Never in the box's text, which Ctrl+C copies.
BuildDialog.BREAK = "\226\128\139"
--- The vertical spacing the preview is laid out with (BuildDialog.previewLayout).
BuildDialog.LAYOUT = {
	tabGap = 2, codeGap = 4, previewGap = 4, aboveReplaces = 8, buttonGap = 6,
	hintClearance = 10, hintMargin = 2,
}

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

--- How the paste box sits: two lines while empty or edited, collapsed to one muted line for a
--- preview or a refusal (opened again while the player edits), the refusal's lines and red
--- border until the box is edited, the gold border only while it has the keyboard, and a good
--- paste highlighted so a second Ctrl+V replaces it.
function BuildDialog.pasteView(verdict, editing, focused)
	local ok = verdict ~= nil and verdict.ok == true
	local refused = verdict ~= nil and not ok and verdict.reason ~= nil
	local showReason = refused and not editing
	local border = focused and "gold" or "border"
	if showReason then
		border = "warning"
	end
	local collapsed = (ok or refused) and not editing
	return {
		collapsed = collapsed,
		lines = collapsed and 1 or BuildDialog.EMPTY_LINES,
		reason = showReason,
		preview = ok,
		use = ok,
		border = border,
		selectAll = (ok or refused) and not editing,
		-- The selection stays (a next Ctrl+V replaces the code) but is never drawn under the
		-- collapsed line: drawn, it read as a solid grey bar (round 3 captures).
		hideSelection = collapsed,
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

--- A site-sent name as a font string may show it: pipes doubled, so it cannot render as markup.
local function escaped(name)
	return (tostring(name or ""):gsub("|", "||"))
end

--- What the preview shows for a build the codec loaded: the class token (for
--- Theme.classIcon), the name (Follow.buildName: the build's own, else its trees; escaped),
--- the line "Warrior · Arms 33 / Fury 16 · 51 points", "You can place 12 now." while points
--- wait (Progress.unspentPoints; nil at 0), and the code's first three gear fields with their
--- slot words under a heading ("Gear in this build · at 60" below 60). `context` =
--- { level, unspent, slotWord = function(slot) }, all optional.
function BuildDialog.preview(build, data, context)
	context = context or {}
	local points = #build.order
	local unspent = math.min(context.unspent or 0, points)
	local detail = joined(
		Talents.titleCase(build.classSlug),
		points > 0 and BuildDialog.split(build, data) or nil,
		pointsWord(points)
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
		heading = string.format(L.dialogGearSlots, heading, #(build.gear or {}), #Codec.SLOTS)
	end
	return {
		classToken = build.classSlug:upper(),
		name = escaped(Follow.buildName(data, build)),
		detail = detail,
		placeNow = unspent > 0 and string.format(L.dialogPlaceNow, unspent) or nil,
		picks = picks,
		heading = heading,
	}
end

--- The first point (1-based) where two orders put a point in different cells, or where one
--- runs out before the other; nil when they are the same order.
function BuildDialog.firstDifference(a, b)
	for index = 1, math.max(#a, #b) do
		local left, right = a[index], b[index]
		if left == nil or right == nil or left.tab ~= right.tab or left.tier ~= right.tier
			or left.column ~= right.column then
			return index
		end
	end
	return nil
end

--- A talent's parts as the dialog names it: { name, tree, have, want } (`have`/`want` only
--- where the plan wants more than one rank and a click takes the next).
local function talentParts(build, data, point, rank)
	local talent = Talents.cellOf(data, build.classSlug, point.tab, point.tier, point.column)
	local parts = {
		name = Chip.talentName(talent) or L.dialogUnnamedTalent,
		tree = Talents.tabName(data, build.classSlug, point.tab),
	}
	if type(rank) == "table" and (rank.want or 0) > 1 and (rank.have or 0) < rank.want then
		parts.have, parts.want = (rank.have or 0) + 1, rank.want
	end
	return parts
end

--- A talent's text: "Deflection (Arms, 1 of 2)", "Improved Heroic Strike (Arms)". `tree`
--- replaces the tree's word (false drops it); `withRank` false drops the rank.
local function talentText(parts, tree, withRank)
	if tree == nil then
		tree = parts.tree
	end
	local rank = withRank ~= false and parts.want ~= nil
	if tree then
		if rank then
			return string.format(L.dialogTalentTreeRank, parts.name, tree, parts.have, parts.want)
		end
		return string.format(L.dialogTalentTree, parts.name, tree)
	end
	if rank then
		return string.format(L.chipPointRank, parts.name, parts.have, parts.want)
	end
	return parts.name
end

--- The next talent the build wants on this character (`ranks`), as the dialog names it, its
--- point and its parts; nil with no build or nothing left to take.
function BuildDialog.dialogTalent(build, ranks, data, level)
	if build == nil then
		return nil
	end
	local point = Progress.build(build, ranks, level).point
	if point == nil then
		return nil
	end
	local parts = talentParts(build, data, point, Chip.pointRankAt(build, ranks, data, point))
	return talentText(parts), point, parts
end

--- Whether `text` fits `width` as `measure` measures it (no measure: it fits).
local function fitsWidth(text, measure, width)
	local measured = measure ~= nil and width ~= nil and measure(text) or nil
	return measured == nil or measured <= width
end

--- `raw` cut to `bytes` (whole letters) with an ellipsis, then escaped: a cut never leaves
--- half of a doubled pipe.
local function cutName(raw, bytes)
	if bytes >= #raw then
		return escaped(raw)
	end
	return escaped((cut(raw, bytes):gsub("%s+$", ""))) .. L.ellipsis
end

--- The build name keeps at least this many bytes before the Replaces line drops the talent's tree.
BuildDialog.NAME_FLOOR = 8

--- The Replaces line: "Replaces Arms raid · same points and gear" / "· same points, different
--- gear list" / "· first differs at point N: <talent (tree)>", fitted to `width` by cutting the
--- build name first, then dropping the talent's tree, never the talent itself.
local function replacesLine(raw, at, parts, sameFormat, measure, width)
	local function line(bytes, talent)
		if at == nil then
			return string.format(sameFormat, cutName(raw, bytes))
		end
		return string.format(L.dialogReplacesDiffers, cutName(raw, bytes), at, talent)
	end
	local talent = parts ~= nil and talentText(parts, nil, false) or nil
	local bytes = #raw
	while bytes > BuildDialog.NAME_FLOOR and not fitsWidth(line(bytes, talent), measure, width) do
		bytes = #cut(raw, bytes - 1)
	end
	if not fitsWidth(line(bytes, talent), measure, width) and parts ~= nil and parts.tree then
		talent = talentText(parts, false, false)
	end
	while bytes > 1 and not fitsWidth(line(bytes, talent), measure, width) do
		bytes = #cut(raw, bytes - 1)
	end
	return line(bytes, talent)
end

--- The Next line: "Next talent would be: Take <talent (tree, rank)>", fitted to `width` by
--- shortening the tree, then dropping the rank, then the tree, never the talent's name.
local function nextLine(parts, measure, width)
	local function line(text)
		return string.format(L.dialogNextWould, string.format(L.chipTake, text))
	end
	local full = line(talentText(parts))
	if fitsWidth(full, measure, width) then
		return full
	end
	if parts.tree then
		local tree = parts.tree
		local bytes = #tree
		while bytes > 1 do
			bytes = #cut(tree, bytes - 1)
			local candidate = line(talentText(parts, (cut(tree, bytes):gsub("%s+$", "")) .. L.ellipsis))
			if fitsWidth(candidate, measure, width) then
				return candidate
			end
		end
		local shortTree = cut(tree, 1) .. L.ellipsis
		local noRank = line(talentText(parts, shortTree, false))
		if fitsWidth(noRank, measure, width) then
			return noRank
		end
	end
	return line(talentText(parts, false, false))
end

--- Whether two builds plan the same gear: the same slots, items and stats, in order.
local function sameGear(a, b)
	local left, right = a.gear or {}, b.gear or {}
	if #left ~= #right then
		return false
	end
	for index, entry in ipairs(left) do
		local other = right[index]
		if entry.slot ~= other.slot or entry.itemId ~= other.itemId then
			return false
		end
		for stat, value in pairs(entry.stats or {}) do
			if (other.stats or {})[stat] ~= value then
				return false
			end
		end
		for stat in pairs(other.stats or {}) do
			if (entry.stats or {})[stat] == nil then
				return false
			end
		end
	end
	return true
end

--- The muted lines above "Use this build". `replaces`: what Use replaces and what differs
--- (same points and gear, same points with a different gear list, or the first point the
--- orders part at, with its talent); absent with nothing followed. `nextWould`: the next talent
--- the candidate gives on this character now. Both fitted to `width` as `measure` measures
--- them (no measure: whole).
function BuildDialog.previewContext(followed, candidate, ranks, data, level, measure, width)
	local context = {}
	if followed ~= nil and candidate ~= nil then
		local raw = Follow.buildName(data, followed)
		local at = BuildDialog.firstDifference(followed.order, candidate.order)
		if at == nil then
			local same = sameGear(followed, candidate) and L.dialogReplacesSameAll or L.dialogReplacesSameOrder
			context.replaces = replacesLine(raw, nil, nil, same, measure, width)
		else
			local cell = candidate.order[at] or followed.order[at]
			local owner = candidate.order[at] ~= nil and candidate or followed
			context.replaces = replacesLine(raw, at, talentParts(owner, data, cell), nil, measure, width)
		end
	end
	local label, _, parts = BuildDialog.dialogTalent(candidate, ranks, data, level)
	if label ~= nil then
		context.nextWould = nextLine(parts, measure, width)
	end
	return context
end

--- The preview's worst case laid out at `line` px a text line (Theme.SIZES and LAYOUT): the
--- title and tabs, a chosen inbox row or the collapsed code (never both), the name, spec and
--- "You can place N now." beside the class icon, the gear heading and three rows at
--- GEAR_PITCH, 8 px, "Replaces" and "Next talent would be", the button, and the hint under
--- it with its clearance. The total, and the parts by name.
function BuildDialog.previewLayout(line)
	local S, G = Theme.SIZES, BuildDialog.LAYOUT
	local parts = {
		top = S.titleBarHeight + S.tabHeight + G.tabGap,
		code = math.max(S.rowHeight, line + S.gap * 2) + G.codeGap,
		header = math.max(S.dialogClassIcon, line * 3 + S.gap / 2 * 2) + G.previewGap,
		gearHeading = S.dialogGearRow,
		gearRows = BuildDialog.GEAR_PICKS * BuildDialog.GEAR_PITCH,
		aboveReplaces = G.aboveReplaces,
		context = line * 2 + S.gap / 2,
		buttonGap = G.buttonGap,
		button = S.buttonHeight + S.gap,
		hintClearance = G.hintClearance,
		hint = line + G.hintMargin,
	}
	local total = 0
	for _, height in pairs(parts) do
		total = total + height
	end
	return total, parts
end

--- How many lines the copy tab's lead takes: one when both of its states (`widths`, the
--- selected line and the copied line as measured) fit `inner`, else two in both, so the box
--- under it never moves when the copy claim appears.
function BuildDialog.leadLines(widths, inner)
	for _, width in pairs(widths) do
		if type(width) == "number" and width > inner then
			return 2
		end
	end
	return 1
end

--- The export with a break point after every "," ":" "|" "=", for the display only, where the
--- client cannot wrap a run of characters itself.
function BuildDialog.breakable(text)
	return (text:gsub("([,:|=])", "%1" .. BuildDialog.BREAK))
end

--- A gear row, slot first: the slot word in a fixed column, then the icon, then the name in
--- what is left of `width`. The x of each, and the name's width.
function BuildDialog.pickRowLayout(width)
	local S = Theme.SIZES
	local name = S.dialogSlotColumn + S.iconSize + S.gap * 2
	return { slot = 0, icon = S.dialogSlotColumn, name = name, nameWidth = width - name }
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
		escaped(Follow.buildName(data, build)), Talents.titleCase(build.classSlug), pointsWord(#build.order)
	))
end

--- The chat line after Use: "Now following Arms raid. Next: Deflection (1 of 2).", the
--- name's pipes doubled so a site-sent name cannot render as markup.
function BuildDialog.chatLine(name, nextLabel)
	local safe = escaped(name)
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

--- Which day `seconds` was, in local time ("today", "yesterday" or "date"), its "14:02", and
--- "2 Nov 14:02" for a date.
local function whenParts(seconds, clock)
	local at, today = clock.fields(seconds), clock.fields(clock.now)
	local dayAt = BuildDialog.civilSeconds(at.year, at.month, at.day, 0, 0, 0)
	local dayToday = BuildDialog.civilSeconds(today.year, today.month, today.day, 0, 0, 0)
	if dayAt == dayToday then
		return "today", hhmm(at)
	elseif dayToday - dayAt == 86400 then
		return "yesterday", hhmm(at)
	end
	local shortDate = Compat.shortDate(string.format("%04d-%02d-%02d", at.year, at.month, at.day)) or ""
	return "date", hhmm(at), string.format(L.dialogOnDate, shortDate, hhmm(at))
end

--- When `seconds` was, in local time: "14:02" today, "yesterday 14:02", else "2 Nov 14:02".
function BuildDialog.when(seconds, clock)
	local kind, time, dated = whenParts(seconds, clock)
	if kind == "today" then
		return time
	elseif kind == "yesterday" then
		return string.format(L.dialogYesterday, time)
	end
	return dated
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
local function seenLine(ack, savedAt, clock, copiedAt)
	if type(ack) ~= "table" then
		return nil
	end
	local seen = BuildDialog.parseUtc(ack.export_seen_at)
	if copiedAt ~= nil and (seen == nil or seen < copiedAt) then
		-- After a real copy, until the site's ack is at or after it.
		return L.dialogWaitingSite
	end
	if seen == nil then
		return nil
	end
	local saved = BuildDialog.parseLocal(savedAt, clock)
	if saved ~= nil and seen < saved then
		return string.format(L.dialogSiteOlder, BuildDialog.when(seen, clock))
	end
	return string.format(L.dialogSeen, BuildDialog.when(seen, clock))
end

--- The ack's sim request state in words: "Sim queued for the 22:30 copy. The result arrives
--- when you next log in." (the export the ack is for), or the other three states' words;
--- nothing for a state this addon does not know.
local function simLine(ack, clock)
	local state = ack.sim_request_state
	if state == "queued" then
		local seen = BuildDialog.parseUtc(ack.export_seen_at)
		if seen ~= nil then
			local kind, time, dated = whenParts(seen, clock)
			if kind == "yesterday" then
				return string.format(L.dialogSimQueuedYesterday, time)
			end
			return string.format(L.dialogSimQueuedFor, dated or time)
		end
		return L.dialogSimQueued
	end
	if type(state) ~= "string" then
		return nil
	end
	return L.dialogSimState[state]
end

--- The whole of copy mode: the box's text and whether it is selected, the lead line under
--- the tabs (the copy claim, in the gain green, only after a real Ctrl+C, `copiedAt`), the
--- ack's lines, and the companion's explanation when no ack was ever seen. No button: the
--- box itself is what the player copies from (fix round 2, item 4).
--- `input` = { code, reason, copiedAt, ack, savedAt, clock, simInUse, upgrade }.
function BuildDialog.copyModel(input)
	local clock = input.clock
	local code = input.code or nil
	local lead
	if not code then
		lead = { text = input.reason or "", hexKey = "muted" }
	elseif input.copiedAt ~= nil then
		lead = { text = string.format(L.dialogYouCopied, hhmm(clock.fields(input.copiedAt))), hexKey = "success" }
	else
		lead = { text = L.dialogSelectedLead, hexKey = "body" }
	end
	local ack = type(input.ack) == "table" and input.ack or nil
	local companion
	if ack == nil then
		companion = { L.dialogNoCompanion, L.dialogCompanionWhat, string.format(L.dialogPasteAtSite, L.siteName) }
	end
	-- While the companion's sim weights are in use, the sim's answer: done, and the chip's
	-- best next upgrade by them (`upgrade` = { name, gain }); else the request's state.
	local sim, simNext
	if input.simInUse then
		sim = L.dialogSimDone
		if type(input.upgrade) == "table" and input.upgrade.name ~= nil then
			simNext = string.format(L.dialogBestUpgrade, escaped(input.upgrade.name), input.upgrade.gain or "")
		end
	elseif ack ~= nil then
		sim = simLine(ack, clock)
	end
	return {
		text = code or "",
		selected = code ~= nil and code ~= false,
		lead = lead,
		seen = seenLine(ack, input.savedAt, clock, input.copiedAt),
		sim = sim,
		simNext = simNext,
		companion = companion,
	}
end

--- The outbox (contract 4, last session's) export's `saved_at`, only when the outbox is
--- this character's (`outboxName`, Export.outboxCharacter).
local function outboxField(db, outboxName, part, field)
	local outbox = type(db) == "table" and db.outbox or nil
	local entry = type(outbox) == "table" and outbox[part] or nil
	if type(entry) ~= "table" or entry.character ~= outboxName or type(entry[field]) ~= "string" then
		return nil
	end
	return entry[field]
end

--- The export box in whole lines: `textHeight` is the export's measured height, `first` one
--- line's height and `advance` each further line's (both measured on the client), `room` the
--- height the box may take. { total, visible, height, scroll }: the lines the text makes, the
--- whole lines that fit (at least one), the box height for them (never a partial line), and
--- whether the rest scrolls.
function BuildDialog.boxLines(textHeight, first, advance, room)
	local total = 1
	if type(textHeight) == "number" and textHeight > first then
		total = 1 + math.floor((textHeight - first) / advance + 0.5)
	end
	local fit = 1 + math.floor(((room or first) - first) / advance)
	local visible = math.max(1, math.min(total, fit))
	return { total = total, visible = visible, height = first + (visible - 1) * advance, scroll = total > visible }
end

function BuildDialog.savedAtFor(db, outboxName)
	return outboxField(db, outboxName, "export", "saved_at")
end


-- Pure rules: builds from the site -------------------------------------------------------

--- The "from the site" rows: one per build in `builds` (Follow.inboxBuilds) whose id and
--- code are both unused (`dismissed`, Follow.dismissedSet) and whose code loads for this
--- player, as { id, name, code, split, label }, the label "Arms raid · Arms 31 / Fury 20 · from
--- the site" (BuildDialog.rowLabel fits it). Another class's build is left out without a word (it was sent
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
				local split = BuildDialog.shortSplit(verdict.build, data)
				rows[#rows + 1] = {
					id = entry.id,
					name = name,
					code = entry.code,
					split = split,
					label = BuildDialog.rowLabel(name, split),
				}
			elseif refusal == nil and not otherClass then
				refusal = verdict.reason
			end
		end
	end
	return rows, refusal
end

--- The points per tree in the players' form, in tree order: "33/16/2".
function BuildDialog.shortSplit(build, data)
	local counts = Progress.planSplit(build)
	local class = data ~= nil and data.classes ~= nil and data.classes[build.classSlug] or nil
	local trees = class ~= nil and #class.tabs or 0
	for tab in pairs(counts) do
		trees = math.max(trees, tab)
	end
	local parts = {}
	for tab = 1, trees do
		parts[tab] = tostring(counts[tab] or 0)
	end
	return table.concat(parts, "/")
end

--- A row's label: "Arms raid (old patch) · 33/16/2 · from the site", fitted to `width` as
--- `measure` measures it by shortening the name with an ellipsis, never the split; the bare
--- "Arms raid · from the site" with no split. Without a measure, the whole label.
function BuildDialog.rowLabel(name, split, measure, width)
	local raw = tostring(name or "")
	local function line(bytes)
		if split == nil or split == "" then
			return string.format(L.dialogInboxRowBare, cutName(raw, bytes))
		end
		return string.format(L.dialogInboxRow, cutName(raw, bytes), split)
	end
	local bytes = #raw
	while bytes > 1 and not fitsWidth(line(bytes), measure, width) do
		bytes = #cut(raw, bytes - 1)
	end
	return line(bytes)
end

--- The rows paste mode shows: with nothing previewed, at most `limit` and "and 2 more" for
--- the rest; while a row's build previews (`verdict.inboxId`), that row alone; while a pasted
--- code previews or is refused, none.
function BuildDialog.rowsView(rows, limit, verdict)
	if verdict ~= nil then
		local chosen = {}
		for _, row in ipairs(rows) do
			if verdict.inboxId ~= nil and row.id == verdict.inboxId then
				chosen[1] = row
			end
		end
		return { shown = chosen, active = chosen[1] ~= nil and chosen[1].id or nil }
	end
	local shown = {}
	for index = 1, math.min(limit, #rows) do
		shown[index] = rows[index]
	end
	local more = #rows > limit and string.format(L.dialogInboxMore, #rows - limit) or nil
	return { shown = shown, more = more }
end

--- The row that feeds the preview: a 2 px gold edge at its left. Every row hovers in the
--- secondary buttons' hover ground, the cue that it can be clicked.
BuildDialog.ROW_ACTIVE_EDGE = { width = 2, token = "gold" }
BuildDialog.ROW_HOVER = Theme.CONTROL_STATES.secondary.hover

--- The paste tab's label: in copy mode, with builds waiting, "1 waiting" in gold.
function BuildDialog.pasteTabLabel(mode, count)
	if mode == "copy" and (count or 0) > 0 then
		return string.format(L.dialogTabPasteCount, L.dialogTabPaste,
			Theme.inlineColor("gold", string.format(L.dialogWaitingCount, count)))
	end
	return L.dialogTabPaste
end

--- The corner hints ("Tab to switch", "Enter to use"): the muted token at HINT_ALPHA.
BuildDialog.HINT_ALPHA = 0.75

function BuildDialog.hintHex()
	return Theme.dimHex("muted", BuildDialog.HINT_ALPHA)
end

--- The no-keyboard note for /fs diag, once a session.
function BuildDialog.noteNoFocus()
	if BuildDialog.focusNoted then
		return false
	end
	BuildDialog.focusNoted = true
	Theme.note(L.diagDialogFocus)
	return true
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

--- The next talent the build wants on this character now, named as the dialog names it
--- (with its tree); nil where the client cannot say.
local function liveDialogNext(build)
	local data = dataOf()
	local ok, label = pcall(BuildDialog.dialogTalent, build, liveRanks(data), data, liveLevel())
	return ok and label or nil
end

--- Whether the item tooltip and the chip score by the companion's sim weights now.
local function simWeightsInUse()
	local ok, tooltip = pcall(function()
		return ns.ItemTooltip or require("ItemTooltip")
	end)
	local current = ok and type(tooltip) == "table" and tooltip.current or nil
	return type(current) == "table" and current.tag == "sim"
end

--- The chip's best next upgrade now, from its own inputs: { name, gain }, or nil.
local function chipUpgrade()
	local input = Chip.currentInput
	local upgrade = type(input) == "table" and type(input.upgrades) == "table" and input.upgrades[1] or nil
	if type(upgrade) ~= "table" or upgrade.name == nil then
		return nil
	end
	return { name = upgrade.name, gain = Chip.gainText(upgrade, input.healing) }
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
	local drawn = BuildDialog.drawPaste()
	local view = BuildDialog.pasteView(BuildDialog.verdict, BuildDialog.editing == true, true)
	if view.selectAll and BuildDialog.paste ~= nil then
		-- A good paste stays highlighted, so a second Ctrl+V replaces it.
		BuildDialog.paste.box:HighlightText()
	end
	return drawn
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

--- One line's height and each further line's advance at the dialog's text size, measured
--- once on the client ("Ag", then "Ag" twice), so every box is sized in whole lines: n lines
--- take first + (n - 1) * advance, never n * first (round 2: the copy box showed glyph tops of
--- a sixth line because the advance is shorter than a line's own height).
local function lineMetrics()
	if BuildDialog.lineFirst == nil then
		local probe = text(BuildDialog.frame, "body")
		local function measured(value)
			probe:SetText(value)
			local height = type(probe.GetStringHeight) == "function" and probe:GetStringHeight() or nil
			return type(height) == "number" and height > 0 and height or nil
		end
		local one, two = measured("Ag"), measured("Ag" .. string.char(10) .. "Ag")
		probe:Hide()
		local first = one or Theme.SIZES.dialogText + 2
		BuildDialog.lineFirst = first
		BuildDialog.lineAdvance = (two ~= nil and two > first) and (two - first) or first
	end
	return BuildDialog.lineFirst, BuildDialog.lineAdvance
end

local function lineHeight()
	return (lineMetrics())
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

--- Size a field to show `height` of text (whole lines, from lineMetrics or boxLines).
local function sizeField(field, clip, box, height)
	local S = Theme.SIZES
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
	-- The close X: Widgets.closeButton's 20 px hit area and its states (Theme.CONTROL_STATES).
	local close = Widgets.closeButton(bar, function()
		BuildDialog.close()
	end)
	close:SetSize(S.dialogClose, S.dialogClose)
	close:SetPoint("RIGHT", bar, "RIGHT", -S.gap * 2, 0)
	enlarge(close.foreverSixtyLabel)
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

--- One planned item in the preview, slot first (BuildDialog.pickRowLayout): the slot word
--- muted in its fixed column, the icon, then the name in its quality colour; hovering
--- shows the item.
local function buildPickRow(parent, width)
	local S = Theme.SIZES
	local layout = BuildDialog.pickRowLayout(width)
	local row = Theme.createFrame("Frame", nil, parent)
	row:SetSize(width, BuildDialog.GEAR_PITCH)
	row:EnableMouse(true)
	local slot = text(row, "muted", S.dialogSlotColumn - S.gap)
	slot:SetPoint("LEFT", row, "LEFT", layout.slot, 0)
	if type(slot.SetWordWrap) == "function" then
		slot:SetWordWrap(false)
	end
	local icon = Theme.icon(row, "ARTWORK", nil, S.iconSize)
	icon:SetPoint("LEFT", row, "LEFT", layout.icon, 0)
	local name = text(row, "body", layout.nameWidth)
	name:SetPoint("LEFT", row, "LEFT", layout.name, 0)
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
	if type(name.SetWordWrap) == "function" then
		name:SetWordWrap(false)
	end
	local detail = text(preview, "muted", inner)
	detail:SetPoint("TOPLEFT", name, "BOTTOMLEFT", 0, -S.gap / 2)
	setMaxLines(detail, 2)
	-- "You can place 12 now." on its own line, in body colour.
	local placeNow = text(preview, "body", inner)
	local headingFrame = Theme.createFrame("Frame", nil, preview)
	headingFrame:SetSize(width, S.dialogGearRow)
	headingFrame:EnableMouse(true)
	local heading = text(headingFrame, "muted", width)
	heading:SetPoint("LEFT", headingFrame, "LEFT", 0, 0)
	headingFrame:SetScript("OnEnter", function(self)
		Theme.showLines(self, { L.dialogGearHeadingTip })
	end)
	headingFrame:SetScript("OnLeave", function()
		Theme.hideTooltip()
	end)
	local picks = {}
	for index = 1, BuildDialog.GEAR_PICKS do
		picks[index] = buildPickRow(preview, width)
	end
	preview:Hide()
	return {
		frame = preview, icon = icon, name = name, detail = detail, placeNow = placeNow,
		headingFrame = headingFrame, heading = heading, picks = picks,
	}
end

local function buildRow(content)
	local S = Theme.SIZES
	local row = { frame = Theme.createFrame("Button", nil, content) }
	row.frame:SetSize(contentWidth(), S.rowHeight)
	local highlight = Theme.texture(row.frame, "BACKGROUND", BuildDialog.ROW_HOVER)
	highlight:SetAllPoints(row.frame)
	highlight:Hide()
	row.highlight = highlight
	local E = BuildDialog.ROW_ACTIVE_EDGE
	row.edge = Theme.texture(row.frame, "ARTWORK", E.token)
	row.edge:SetPoint("TOPLEFT", row.frame, "TOPLEFT", 0, 0)
	row.edge:SetPoint("BOTTOMLEFT", row.frame, "BOTTOMLEFT", 0, 0)
	row.edge:SetWidth(E.width)
	row.edge:Hide()
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

local function buildPaste(frame, top, bottom)
	local S = Theme.SIZES
	local width = contentWidth()
	local content = Theme.createFrame("Frame", nil, frame)
	content:SetPoint("TOPLEFT", frame, "TOPLEFT", S.padding, -top)
	content:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", -S.padding, bottom)
	local paste = { content = content, rows = {} }
	paste.following = text(content, "body", width)
	if type(paste.following.SetWordWrap) == "function" then
		paste.following:SetWordWrap(false)
	end
	paste.next = text(content, "muted", width)
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
	-- Once a code loads, the box reads as one muted line with an ellipsis; a click opens it.
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
	local function startEditing()
		if not BuildDialog.editing and BuildDialog.verdict ~= nil then
			BuildDialog.editing = true
			BuildDialog.drawPaste()
			box:SetFocus()
		end
	end
	box:SetScript("OnMouseDown", startEditing)
	-- While collapsed the box takes no mouse (an edit box grows with its text, and the
	-- invisible one covered the preview and the Use button, round 4); the field takes the click.
	field:EnableMouse(true)
	field:SetScript("OnMouseDown", startEditing)
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
	-- A chosen row sits where the collapsed code would, above the (then invisible) box.
	if type(field.GetFrameLevel) == "function" then
		for _, row in ipairs(paste.rows) do
			row.frame:SetFrameLevel((field:GetFrameLevel() or 0) + 3)
		end
	end
	paste.lead = text(content, "warning", width)
	setMaxLines(paste.lead, 2)
	paste.body = text(content, "body", width)
	setMaxLines(paste.body, 2)
	paste.still = text(content, "muted", width)
	paste.preview = buildPreview(content)
	paste.use = Cards.primaryButton(content, L.dialogUse, function()
		BuildDialog.useCurrent()
	end)
	enlarge(paste.use.foreverSixtyLabel)
	paste.use:SetPoint("BOTTOMLEFT", content, "BOTTOMLEFT", 0, 0)
	-- An edit box grows with its text: the open box must never sit over the preview or Use.
	local boxLevel = type(box.GetFrameLevel) == "function" and box:GetFrameLevel() or 0
	paste.preview.frame:SetFrameLevel((boxLevel or 0) + 4)
	paste.use:SetFrameLevel((boxLevel or 0) + 5)
	paste.use:Hide()
	-- Above the button: what Use replaces, and the next talent the candidate gives.
	paste.nextWould = text(content, "muted", width)
	paste.nextWould:SetPoint("BOTTOMLEFT", paste.use, "TOPLEFT", 0, BuildDialog.LAYOUT.buttonGap)
	paste.replaces = text(content, "muted", width)
	paste.replaces:SetPoint("BOTTOMLEFT", paste.nextWould, "TOPLEFT", 0, S.gap / 2)
	for _, region in ipairs({ paste.nextWould, paste.replaces }) do
		if type(region.SetWordWrap) == "function" then
			region:SetWordWrap(false)
		end
	end
	return paste
end

--- Show or hide the selection ground behind the export, as the box's selection stands.
local function paintSelection()
	local copy = BuildDialog.copyTab
	if copy == nil then
		return
	end
	if BuildDialog.selected and BuildDialog.exportCode ~= nil then
		copy.selection:Show()
	else
		copy.selection:Hide()
	end
end

local function buildCopy(frame, top, bottom)
	local S = Theme.SIZES
	local width = contentWidth()
	local content = Theme.createFrame("Frame", nil, frame)
	content:SetPoint("TOPLEFT", frame, "TOPLEFT", S.padding, -top)
	content:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", -S.padding, bottom)
	local copy = { content = content }
	copy.lead = text(content, "body", width)
	copy.lead:SetPoint("TOPLEFT", content, "TOPLEFT", 0, 0)
	setMaxLines(copy.lead, 2)
	local box, field, clip = boxInField(content, true)
	copy.box, copy.field, copy.clip = box, field, clip
	-- The box holds the exact export and takes Ctrl+C, unseen; what the player reads is the
	-- display over it, which wraps a run of characters with no spaces at the edge (an edit
	-- box wraps only at spaces, and the export has none).
	box:SetAlpha(0)
	-- Its selection highlight would draw over the label (round 4: a grey block): never drawn.
	if type(box.SetHighlightColor) == "function" then
		box:SetHighlightColor(0, 0, 0, 0)
	end
	local over = Theme.createFrame("Frame", nil, clip)
	copy.over = over
	over:SetAllPoints(clip)
	if type(box.GetFrameLevel) == "function" and type(over.SetFrameLevel) == "function" then
		over:SetFrameLevel((box:GetFrameLevel() or 0) + 2)
	end
	copy.display = text(over, "muted", width - S.gap * 4 - S.scrollbar - S.gap * 1.5)
	copy.display:SetPoint("TOPLEFT", clip, "TOPLEFT", 0, 0)
	if type(copy.display.SetWordWrap) == "function" then
		copy.display:SetWordWrap(true)
	end
	copy.wrapsRuns = type(copy.display.SetNonSpaceWrap) == "function"
	if copy.wrapsRuns then
		copy.display:SetNonSpaceWrap(true)
	end
	-- The selection, as a subtle tint behind the label.
	copy.selection = Theme.texture(over, "BACKGROUND", "hover")
	copy.selection:SetAllPoints(copy.display)
	copy.selection:Hide()
	bindKeys(box, function()
		BuildDialog.copy()
	end)
	-- The one copy claim the dialog makes comes from a real Ctrl+C in this box. Chained, so
	-- Widgets' own focus release after Ctrl+C still runs.
	-- Its own handler, not chained: Widgets' release after Ctrl+C would drop the keyboard and the
	-- selection (fix round 5, item 8); a second Ctrl+C copies again, Escape closes.
	box:SetScript("OnKeyDown", function(self, key)
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
		-- The press places the cursor: the export is no longer all selected...
		BuildDialog.selected = false
		paintSelection()
	end)
	chain(box, "OnMouseUp", function()
		-- ...and the click ends selecting it all again, for Ctrl+C.
		BuildDialog.copy()
	end)
	chain(box, "OnEditFocusGained", function()
		paintEdges(field.foreverSixtyBorder, "gold")
	end)
	chain(box, "OnEditFocusLost", function()
		BuildDialog.selected = false
		paintSelection()
		paintEdges(field.foreverSixtyBorder, "border")
	end)
	-- The whole export, wrapped; the wheel scrolls it a line at a time, with a thumb.
	clip:EnableMouseWheel(true)
	clip:SetScript("OnMouseWheel", function(_, delta)
		BuildDialog.scrollCopy(-delta)
	end)
	-- The Theme's thin scrollbar inside the field's right edge, clear of the text.
	copy.bar = Widgets.scrollbar(field, field, -(S.scrollbar + S.gap / 2))
	-- Measures the lead's two states unwrapped, so the box under it never moves.
	copy.probe = text(content, "body")
	copy.probe:Hide()
	copy.seen = text(content, "muted", width)
	setMaxLines(copy.seen, 2)
	copy.sim = text(content, "muted", width)
	copy.simNext = text(content, "muted", width)
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
		if BuildDialog.frame:IsShown() and BuildDialog.mode == "paste" then
			BuildDialog.drawPaste()
		end
	end
end

--- The dialog, built on the first open: a 420x300 panel titled "Forever Sixty", draggable
--- by its title, closed by Escape, with the two tabs under the title and "Tab to switch" at
--- its bottom left.
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
	local G = BuildDialog.LAYOUT
	tabs.paste:SetPoint("TOPLEFT", frame, "TOPLEFT", S.padding, -S.titleBarHeight)
	tabs.copy:SetPoint("LEFT", tabs.paste, "RIGHT", S.gap * 2, 0)
	BuildDialog.tabs = tabs
	-- On the dialog's bottom edge, under both tabs' content and at least 10 px below the
	-- button, a step dimmer than a tab; "Enter to use" at the right while a preview shows.
	local hint = text(frame, "muted")
	hint:SetText(L.dialogTabHint)
	hint:SetTextColor(Theme.rgb(BuildDialog.hintHex()))
	hint:SetPoint(BuildDialog.HINT_POINT, frame, BuildDialog.HINT_POINT, S.padding, G.hintMargin)
	local enter = text(frame, "muted")
	enter:SetText(L.dialogEnterHint)
	enter:SetTextColor(Theme.rgb(BuildDialog.hintHex()))
	enter:SetJustifyH("RIGHT")
	enter:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", -S.padding, G.hintMargin)
	enter:Hide()
	BuildDialog.enterHint = enter
	local top = S.titleBarHeight + S.tabHeight + G.tabGap
	local bottom = G.hintMargin + lineHeight() + G.hintClearance
	BuildDialog.contentHeight = S.dialogHeight - top - bottom
	BuildDialog.paste = buildPaste(frame, top, bottom)
	BuildDialog.copyTab = buildCopy(frame, top, bottom)
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

local function widthOf(region)
	return function(value)
		region:SetText(value)
		return type(region.GetStringWidth) == "function" and region:GetStringWidth() or nil
	end
end

local function drawPreview(verdict, y)
	local S = Theme.SIZES
	local paste = BuildDialog.paste
	local preview = paste.preview
	local data = dataOf()
	local ranks = liveRanks(data)
	local model = BuildDialog.preview(verdict.build, data, {
		level = liveLevel(),
		unspent = Progress.unspentPoints(liveLevel(), ranks),
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
	local textY = textHeight(preview.name) + S.gap / 2 + textHeight(preview.detail)
	if model.placeNow ~= nil then
		preview.placeNow:SetText(model.placeNow)
		preview.placeNow:ClearAllPoints()
		preview.placeNow:SetPoint("TOPLEFT", preview.detail, "BOTTOMLEFT", 0, -S.gap / 2)
		preview.placeNow:Show()
		textY = textY + S.gap / 2 + textHeight(preview.placeNow)
	else
		preview.placeNow:Hide()
	end
	local headY = math.max(S.dialogClassIcon, textY) + BuildDialog.LAYOUT.previewGap
	if model.heading ~= nil then
		preview.heading:SetText(model.heading)
		preview.headingFrame:ClearAllPoints()
		preview.headingFrame:SetPoint("TOPLEFT", preview.frame, "TOPLEFT", 0, -headY)
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
			pick.slot:SetText(entry.slotWord)
			pick.icon:SetTexture(Compat.itemIcon(entry.itemId) or Theme.UNKNOWN_ICON)
			local name = Compat.itemName(entry.itemId)
			if name == nil then
				BuildDialog.pending[entry.itemId] = true
			end
			pick.name:SetText(name or "")
			pick.name:SetTextColor(Theme.qualityColor(Compat.itemQuality(entry.itemId)))
			pick.frame:ClearAllPoints()
			pick.frame:SetPoint("TOPLEFT", preview.frame, "TOPLEFT", 0, -headY)
			pick.frame:Show()
			headY = headY + BuildDialog.GEAR_PITCH
		end
	end
	preview.frame:SetHeight(headY)
	stack(preview.frame, paste.content, y, headY)
	-- What Use replaces and how the orders compare, and the next talent the candidate gives.
	local ok, context = pcall(BuildDialog.previewContext, Follow.build, verdict.build, ranks, data, liveLevel(),
		widthOf(paste.replaces), contentWidth())
	if not ok then
		context = {}
	end
	for key, region in pairs({ replaces = paste.replaces, nextWould = paste.nextWould }) do
		if context[key] ~= nil then
			region:SetText(context[key])
			region:Show()
		else
			region:SetText("")
			region:Hide()
		end
	end
end

--- Redraw paste mode from the verdict: with nothing pasted, "Following" and "Next talent"
--- together above a two-line box; the rows from the site (only the chosen one while it
--- previews); the box, one muted line once a code loads; the refusal's three lines or the
--- preview with what Use replaces; and the gold Use button.
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
	for _, region in ipairs({ paste.following, paste.next, paste.more, paste.lead, paste.body, paste.still,
		paste.replaces, paste.nextWould }) do
		region:Hide()
	end
	paste.preview.frame:Hide()
	paste.use:Hide()
	BuildDialog.enterHint:Hide()
	local y = 0
	if verdict == nil then
		local line = BuildDialog.followingLine(Follow.build, data)
		if line ~= nil then
			paste.following:SetText(line)
			y = stack(paste.following, content, y, textHeight(paste.following) + S.gap / 2)
		end
		local talentLine = BuildDialog.nextTalentLine(liveDialogNext(Follow.build))
		if talentLine ~= nil then
			paste.next:SetText(talentLine)
			y = stack(paste.next, content, y, textHeight(paste.next))
		end
		if y > 0 then
			y = y + S.gap * 2
		end
	end
	local rows = BuildDialog.rowsView(BuildDialog.rows or {}, S.dialogInboxRows, verdict)
	-- A chosen row stands in for the collapsed code: one line, never both.
	local rowY = y
	local chosen = verdict ~= nil and verdict.inboxId ~= nil and #rows.shown > 0
	for index, row in ipairs(paste.rows) do
		local entry = rows.shown[index]
		row.entry = entry
		if entry == nil then
			row.frame:Hide()
		else
			row.label:SetText(BuildDialog.rowLabel(entry.name, entry.split, widthOf(row.label),
				contentWidth() - S.gap * 2))
			row.selected = verdict ~= nil and verdict.inboxId == entry.id
			if rows.active ~= nil and rows.active == entry.id then
				row.edge:Show()
			else
				row.edge:Hide()
			end
			if row.selected then
				row.highlight:Show()
			else
				row.highlight:Hide()
			end
			y = stack(row.frame, content, y, S.rowHeight)
		end
	end
	if rows.more ~= nil then
		paste.more:SetText(rows.more)
		y = stack(paste.more, content, y, textHeight(paste.more))
	end
	if #rows.shown > 0 and not chosen then
		y = y + S.gap
	end
	-- The box, in whole lines: two while empty or edited, one muted line once a code loads.
	local first, advance = lineMetrics()
	local boxText = paste.box:GetText() or ""
	local fieldHeight = sizeField(paste.field, paste.clip, paste.box, first + (view.lines - 1) * advance)
	if chosen then
		stack(paste.field, content, rowY, fieldHeight)
		paste.field:SetAlpha(0)
		y = math.max(y, rowY + fieldHeight) + BuildDialog.LAYOUT.codeGap
	else
		paste.field:SetAlpha(1)
		y = stack(paste.field, content, y, fieldHeight) + BuildDialog.LAYOUT.codeGap
	end
	paste.box:EnableMouse(not view.collapsed)
	-- The selection under a collapsed line is kept (a next Ctrl+V replaces it) but never drawn.
	if type(paste.box.SetHighlightColor) == "function" then
		if view.hideSelection then
			paste.box:SetHighlightColor(0, 0, 0, 0)
		else
			local r, g, b = Theme.rgb(Theme.HEX.gold)
			paste.box:SetHighlightColor(r, g, b, 0.35)
		end
	end
	if view.collapsed then
		paste.box:SetAlpha(0)
		local shown = verdict.code or boxText:match("^%s*(.-)%s*$"):gsub("%s+", " ")
		paste.display:SetTextColor(Theme.rgb(Theme.HEX[BuildDialog.COLLAPSED_TEXT]))
		paste.display:SetText(Chip.fit(shown, widthOf(paste.display), contentWidth() - S.gap * 4))
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
		local refusal = BuildDialog.refusal(verdict.reason, followed and followed:gsub("|", "||") or nil)
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
		drawPreview(verdict, y)
		paste.use:Show()
		if BuildDialog.mode == "paste" then
			BuildDialog.enterHint:Show()
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

--- Scroll the export by `lines` whole lines inside its field, and move the thumb.
function BuildDialog.scrollCopy(lines)
	local copy = BuildDialog.copyTab
	if copy == nil or copy.lines == nil then
		return 0
	end
	local _, advance = lineMetrics()
	local range = math.max(0, copy.lines.total - copy.lines.visible)
	BuildDialog.copyOffset = math.max(0, math.min(range, (BuildDialog.copyOffset or 0) + lines))
	copy.display:ClearAllPoints()
	copy.display:SetPoint("TOPLEFT", copy.clip, "TOPLEFT", 0, BuildDialog.copyOffset * advance)
	copy.bar:Update(copy.lines.visible / copy.lines.total, range > 0 and BuildDialog.copyOffset / range or 0)
	return BuildDialog.copyOffset
end

--- Fill copy mode from copyModel: the lead line, the whole export wrapped in its box in
--- whole lines (as many as fit above the ack's and the companion's lines; the rest scrolls),
--- selected when `select`, then the ack's lines and the companion's.
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
		clock = liveClock(),
		simInUse = simWeightsInUse(),
		upgrade = chipUpgrade(),
	})
	local content = copy.content
	copy.lead:SetText(model.lead.text)
	copy.lead:SetTextColor(Theme.rgb(Theme.HEX[model.lead.hexKey]))
	-- The lead takes the same lines whether it says "Selected" or "You copied", measured.
	local function leadWidth(value)
		copy.probe:SetText(value)
		local width = type(copy.probe.GetUnboundedStringWidth) == "function" and copy.probe:GetUnboundedStringWidth()
			or (type(copy.probe.GetStringWidth) == "function" and copy.probe:GetStringWidth() or nil)
		return width
	end
	local first, advance = lineMetrics()
	local leadLines = BuildDialog.leadLines({ leadWidth(L.dialogSelectedLead),
		leadWidth(string.format(L.dialogYouCopied, "00:00")), leadWidth(model.lead.text) }, contentWidth())
	local top = first + (leadLines - 1) * advance + S.gap * 2
	-- The lines under the box first, so the box takes only the room left. The seen line is
	-- reserved at the taller of its two states (before and after a copy), so a Ctrl+C never
	-- changes the box's height.
	local below = 0
	if model.seen ~= nil then
		local tallest = 0
		for _, line in ipairs({ L.dialogWaitingSite, model.seen }) do
			copy.seen:SetText(line)
			tallest = math.max(tallest, textHeight(copy.seen))
		end
		copy.seen:SetText(model.seen)
		below = below + tallest + S.gap
	end
	for _, entry in ipairs({ { copy.sim, model.sim }, { copy.simNext, model.simNext } }) do
		if entry[2] ~= nil then
			entry[1]:SetText(entry[2])
			below = below + textHeight(entry[1]) + S.gap
		end
	end
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
	copy.display:SetText(copy.wrapsRuns and model.text or BuildDialog.breakable(model.text))
	local room = (BuildDialog.contentHeight or S.dialogHeight) - top - S.gap * 4 - below - bottom
	copy.lines = BuildDialog.boxLines(textHeight(copy.display), first, advance, room)
	local fieldHeight = sizeField(copy.field, copy.clip, copy.box, copy.lines.height)
	copy.box:SetHeight(first + (copy.lines.total - 1) * advance)
	local y = stack(copy.field, content, top, fieldHeight) + S.gap * 2
	if select ~= false then
		BuildDialog.copyOffset = 0
		if model.selected then
			Widgets.selectText(copy.box, model.text)
			BuildDialog.selected = true
		else
			Widgets.setText(copy.box, model.text)
		end
	end
	BuildDialog.scrollCopy(0)
	paintSelection()
	copy.seen:Hide()
	copy.sim:Hide()
	copy.simNext:Hide()
	if model.seen ~= nil then
		y = stack(copy.seen, content, y, textHeight(copy.seen) + S.gap)
	end
	if model.sim ~= nil then
		y = stack(copy.sim, content, y, textHeight(copy.sim) + S.gap)
	end
	if model.simNext ~= nil then
		stack(copy.simNext, content, y, textHeight(copy.simNext))
	end
	return model
end

--- Select the export again for Ctrl+C (a click in the box, Enter, the deferred refocus). It
--- claims nothing itself.
function BuildDialog.copy()
	local copy = BuildDialog.copyTab
	if BuildDialog.exportCode == nil or copy == nil then
		return nil
	end
	Widgets.selectText(copy.box, BuildDialog.exportCode)
	BuildDialog.selected = true
	paintSelection()
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
		BuildDialog.enterHint:Hide()
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
		BuildDialog.copiedAt = nil
		BuildDialog.enterHint:Hide()
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
		BuildDialog.noteNoFocus()
		return false
	end
	return true
end

ns.BuildDialog = BuildDialog
return BuildDialog
