-- addon/ForeverSixty/surfaces/BuildDialog.lua
-- The Builds window (spec 2026-11-13-addon-builds-window): the one big frame of the addon's own,
-- 740 wide. The left column is the class's builds (BuildsPane: one tap follows one) or a code
-- previewing; the right column says what comes next on the followed build (BuildsWindow's Up next),
-- takes a build in from the site (the rows the companion brought and the paste box) and sends the
-- character out to it (the copy box).
--
-- Pure model, thin frame. BuildDialog.evaluate is the paste box's whole verdict
-- (Codec.loadBuild's own answer, the codec's own reason on a refusal); pasteView,
-- refusal, preview, followingLine and nextTalent are what paste mode shows; copyModel
-- is the whole of copy mode with the clock handed in; inboxRows, rowText and rowsView are the
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
local Compat = ns.Compat or require("Compat")
local Codec = ns.Codec or require("Codec")
local Talents = ns.Talents or require("Talents")
local Export = ns.Export or require("Export")
local Follow = ns.Follow or require("Follow")
local Progress = ns.Progress or require("Progress")
local Gear = ns.Gear or require("Gear")
local Chip = ns.Chip or require("Chip")
local BuildsPane = ns.BuildsPane or require("BuildsPane")
local BuildsWindow = ns.BuildsWindow or require("BuildsWindow")
local Hub = ns.Hub or require("Hub")

local BuildDialog = {}

--- How many gear rows the preview shows at once: every item the code plans is a row, and the
--- list scrolls inside the window past this many (the wheel over it).
BuildDialog.GEAR_VISIBLE = 4
--- The most characters a refusal's lead keeps (two lines of the dialog's text).
BuildDialog.REASON_CHARS = 120
--- How far an edge may sit past the screen's and still count as on it (UI px).
BuildDialog.SCREEN_TOLERANCE = 1
--- The paste box's height in lines while empty or edited (one once a code loads).
BuildDialog.EMPTY_LINES = 2
--- The collapsed code's colour: muted text on the box's inset ground.
BuildDialog.COLLAPSED_TEXT = "muted"
--- The gear rows' pitch: two lines, the item on its quality colour and under it where it comes from
--- in full (never cut beside the name, which is what the first line gives way to).
BuildDialog.GEAR_PITCH = 32
--- The tallest the preview may stand at its worst (px): the window grows to it, never past.
BuildDialog.PREVIEW_MAX_HEIGHT = 400
--- A zero-width space: a break point the export's display may wrap at where the client cannot
--- break a run of characters itself. Never in the box's text, which Ctrl+C copies.
BuildDialog.BREAK = "\226\128\139"
--- The vertical spacing the preview is laid out with (BuildDialog.previewLayout).
BuildDialog.LAYOUT = {
	codeGap = 4, previewGap = 4, aboveReplaces = 8, buttonGap = 6,
}

--- A site row's ringed crest, smaller than the preview header's (2026-11-13: the rows sit under the
--- own builds and read as a list, not as headers).
BuildDialog.SITE_CREST = 20
--- The room between "and N more" and the paste box under it.
BuildDialog.MORE_GAP = 6

--- Injection points Options fills: the data to read and what to redraw after a build is used
--- (Options.refreshSurfaces). `say` prints the chat line (Theme.say when nil). What runs when the
--- window cannot show itself, and the tour's step aside, are the hub's (Hub.fallback, Hub.beforeOpen).
BuildDialog.data = nil
BuildDialog.afterUse = nil
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
--- paste highlighted so a second Ctrl+V replaces it. `resting` (builds wait, the player has not
--- touched the box): the box holds the keyboard for Enter and the arrow keys but is not what the
--- player is working in, so its border stays quiet until a click or a key of theirs (X.5).
function BuildDialog.pasteView(verdict, editing, focused, resting)
	local ok = verdict ~= nil and verdict.ok == true
	local refused = verdict ~= nil and not ok and verdict.reason ~= nil
	local showReason = refused and not editing
	local border = focused and not resting and "gold" or "border"
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

--- The build's trees by the points its order spends in each, biggest first: "Arms 31 / Fury 20"
--- (BuildsPane.split, the one rule).
BuildDialog.split = BuildsPane.split

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

--- The BiS pick (or one of its alternatives) in `data` that names `itemId` for any spec of
--- `classSlug`, in a fixed order (spec, band, faction, slot), or nil (L3.7 item 31: a build's
--- gear rows carry the source the item tooltip prints). The pick is Gear.decodePick's.
function BuildDialog.pickFor(data, classSlug, itemId)
	local bis = type(data) == "table" and type(data.bis) == "table" and data.bis or nil
	if bis == nil or classSlug == nil or itemId == nil then
		return nil
	end
	local function sorted(map)
		local keys = {}
		for key in pairs(map) do
			keys[#keys + 1] = key
		end
		table.sort(keys, function(left, right)
			return tostring(left) < tostring(right)
		end)
		return keys
	end
	local prefix = classSlug .. "-"
	for _, spec in ipairs(sorted(bis)) do
		if type(spec) == "string" and spec:sub(1, #prefix) == prefix and type(bis[spec]) == "table" then
			for _, band in ipairs(sorted(bis[spec])) do
				for _, faction in ipairs(sorted(bis[spec][band])) do
					local bySlot = bis[spec][band][faction]
					for _, slot in ipairs(sorted(bySlot)) do
						local pick = Gear.decodePick(bySlot[slot])
						if pick ~= nil then
							if pick.item_id == itemId then
								return pick
							end
							for _, alternative in ipairs(pick.alternatives) do
								if alternative.item_id == itemId then
									return alternative
								end
							end
						end
					end
				end
			end
		end
	end
	return nil
end

--- A site-sent name as a font string may show it: pipes doubled, so it cannot render as markup.
local function escaped(name)
	return (tostring(name or ""):gsub("|", "||"))
end

--- "Warrior · Arms 33 / Fury 16 · 51 points", the line under a build's name in the preview and in a
--- row from the site.
function BuildDialog.detailLine(build, data, withoutPoints)
	local points = #build.order
	return joined(
		Talents.titleCase(build.classSlug),
		points > 0 and BuildDialog.split(build, data) or nil,
		not withoutPoints and pointsWord(points) or nil
	)
end

--- A build from the site's title line: "Beast Mastery build · 15 points" (the day or the sender is on
--- the muted line under it). The name is cut (never the size) to what `width` leaves as `measure`
--- measures it; whole with no measure.
function BuildDialog.siteTitle(raw, points, measure, width)
	local tail = joined(nil, points)
	local room = width
	if measure ~= nil and width ~= nil then
		local used = measure(L.dialogSeparator .. tail)
		room = type(used) == "number" and width - used or width
	end
	return joined(BuildDialog.rowText(raw, measure, room), points)
end

--- The level a build is for: a point a level from 10, so its points name it (15 points are a level
--- 24 build, 51 a level 60 one); nil for a build with no points, which names no level.
function BuildDialog.buildLevel(build)
	local points = build ~= nil and build.order ~= nil and #build.order or 0
	if points <= 0 then
		return nil
	end
	return math.min(Progress.MAX_LEVEL, 9 + points)
end

--- The kinds of source whose place is worth leading with the kind: "Vendor · Darnassus".
local KIND_FIRST = { V = true, Q = true, R = true, P = true }

--- Where an item comes from, in full, as the gear row says it: "Wailing Caverns · Lady Anacondra ·
--- 20% drop", "Vendor · Darnassus", "Leatherworking · you can craft this" when the character has the
--- profession (`professions`, a set of profession slugs), "Crafted · Leatherworking" when not, and
--- "Source unknown" for an item no pick names or a pick with no source.
function BuildDialog.sourceText(pick, professions)
	local tooltip = ns.ItemTooltip or require("ItemTooltip")
	local first, second, percent = tooltip.sourceParts(pick)
	if first == nil then
		return L.dialogSourceUnknown
	end
	if pick.kind == "C" and type(pick.source) == "string" and pick.source ~= "" then
		local slug = Export.slugify(pick.source)
		if slug ~= nil and type(professions) == "table" and professions[slug] then
			return string.format(L.dialogSourceCraftable, pick.source)
		end
	elseif KIND_FIRST[pick.kind] and second ~= nil then
		first, second = second, first
	end
	return tooltip.sourceText(first, second, percent)
end

--- `text` (a source line, its parts joined by " · ") within `room` as `measure` measures it, cut
--- only between words: whole; else without the drop chance; else the leading place shortened by
--- whole words with an ellipsis, the rest kept. No measure: whole.
function BuildDialog.fitSourceText(text, measure, room)
	local function fits(value)
		local measured = measure ~= nil and room ~= nil and measure(value) or nil
		return measured == nil or measured <= room
	end
	if fits(text) then
		return text
	end
	local sep = L.dialogSeparator
	local parts, from = {}, 1
	while true do
		local at = text:find(sep, from, true)
		if at == nil then
			parts[#parts + 1] = text:sub(from)
			break
		end
		parts[#parts + 1] = text:sub(from, at - 1)
		from = at + #sep
	end
	if #parts > 1 and parts[#parts]:find("%d+%% drop$") then
		table.remove(parts)
		local line = table.concat(parts, sep)
		if fits(line) then
			return line
		end
	end
	local rest = #parts > 1 and sep .. table.concat(parts, sep, 2) or ""
	local words = {}
	for word in parts[1]:gmatch("%S+") do
		words[#words + 1] = word
	end
	for count = #words - 1, 1, -1 do
		local line = table.concat(words, " ", 1, count) .. L.ellipsis .. rest
		if fits(line) then
			return line
		end
	end
	return table.concat(parts, sep)
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
	local detail = BuildDialog.detailLine(build, data)
	local wordOf = context.slotWord or function(slot)
		return BuildDialog.slotWord(slot, {})
	end
	-- Every slot the code carries is a row (the list scrolls), the ranged slot like any other.
	local picks = {}
	for _, entry in ipairs(build.gear or {}) do
		if type(entry.itemId) == "number" then
			local pick = BuildDialog.pickFor(data, build.classSlug, entry.itemId)
			picks[#picks + 1] = { itemId = entry.itemId, slot = entry.slot, slotWord = wordOf(entry.slot),
				pick = pick, source = BuildDialog.sourceText(pick, context.professions) }
		end
	end
	local heading
	if #picks > 0 then
		heading = L.dialogGearHeading
		local own = BuildDialog.buildLevel(build)
		if own ~= nil and own < Progress.MAX_LEVEL then
			heading = string.format(L.dialogGearLevel, heading, own)
		elseif type(context.level) == "number" and context.level < Progress.MAX_LEVEL then
			heading = string.format(L.dialogGearAt, heading, Progress.MAX_LEVEL)
		end
		heading = string.format(L.dialogGearSlots, heading, #picks, #Codec.SLOTS)
	end
	return {
		classToken = build.classSlug:upper(),
		name = escaped(Follow.buildName(data, build)),
		rawName = Follow.buildName(data, build),
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

--- The Replaces sentence: "Replaces Arms raid · same points and gear · nothing changes" / "· same
--- points, different gear list" / "· first differs at point N: <talent (tree)>". Whole, never cut: the
--- frame wraps it onto a second line and the dialog grows (X.5: "Replaces Beast Ma…" decided a click).
local function replacesLine(raw, size, at, parts, sameFormat)
	if at == nil then
		return string.format(sameFormat, escaped(raw), size)
	end
	return string.format(L.dialogReplacesDiffers, escaped(raw), size, at, talentText(parts, nil, false))
end

--- The Next line: "Next talent would be: Take <talent (tree, rank)>", whole. Like the Replaces
--- sentence the frame wraps it onto a second line and the dialog grows (X.6: "(Beast M..." decided a click).
local function nextLine(parts)
	return string.format(L.dialogNextWould, string.format(L.chipTake, talentText(parts)))
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

--- The muted lines above "Use this build", each whole (the frame wraps them; X.5). `replaces`: what
--- Use replaces and what differs (same points and gear, same points with a different gear list, or
--- the first point the orders part at, with its talent); absent with nothing followed. When the
--- candidate is a strict prefix of the followed plan the sentence is "Same as your plan for 11
--- points. Yours goes on with Bestial Swiftness." instead, and `blocked` says why Use is held
--- back ("Your plan is further along."). `short`, for a candidate that ends before the points
--- `ranks` have spent and is not a prefix: "This build ends at 11 points; you have spent 15. It
--- keeps following from point 12.", above `replaces`. `nextWould`: the next talent the candidate
--- gives on this character now, whole (the frame wraps it).
function BuildDialog.previewContext(followed, candidate, ranks, data, level)
	local context = {}
	local prefix = false
	if followed ~= nil and candidate ~= nil then
		local raw = Follow.buildName(data, followed)
		local size = pointsWord(#followed.order)
		local at = BuildDialog.firstDifference(followed.order, candidate.order)
		if at == nil then
			local same = sameGear(followed, candidate) and L.dialogReplacesSameAll or L.dialogReplacesSameOrder
			context.replaces = replacesLine(raw, size, nil, nil, same)
		elseif candidate.order[at] == nil then
			prefix = true
			context.replaces = string.format(L.dialogSameAsPlan, pointsWord(#candidate.order),
				talentText(talentParts(followed, data, followed.order[at]), false, false))
			context.blocked = L.dialogUseFurther
		else
			context.replaces = replacesLine(raw, size, at, talentParts(candidate, data, candidate.order[at]), nil)
		end
	end
	local spent = Progress.pointsSpent(ranks)
	if candidate ~= nil and not prefix and #candidate.order < spent then
		context.short = string.format(L.dialogEndsShort, pointsWord(#candidate.order), spent, #candidate.order + 1)
	end
	local label, _, parts = BuildDialog.dialogTalent(candidate, ranks, data, level)
	if label ~= nil then
		context.nextWould = nextLine(parts)
	end
	return context
end

--- The preview's worst case laid out at `line` px a text line: the window's header and the body's
--- padding, the name, spec and "You can place N now." beside the class icon (a build from the site
--- takes a fourth line for the character it is for), the gear heading and GEAR_VISIBLE rows at
--- GEAR_PITCH, 8 px, `contextLines` of "Replaces" and "Next talent would be" (two; each wrapped line of
--- the replace sentence and the "ends at" sentence adds one, and the window grows by it,
--- Hub.fitHeight), the button, the body's padding and the footer. The total, and the parts by name.
function BuildDialog.previewLayout(line, contextLines, fromSite)
	local S, G, H = Theme.SIZES, BuildDialog.LAYOUT, Hub.LAYOUT
	contextLines = contextLines or 2
	local headerLines = fromSite and 4 or 3
	local parts = {
		top = H.header + H.bodyPad,
		header = math.max(S.dialogClassIcon, line * headerLines + S.gap / 2 * (headerLines - 1)) + G.previewGap,
		gearHeading = S.dialogGearRow,
		gearRows = BuildDialog.GEAR_VISIBLE * BuildDialog.GEAR_PITCH,
		aboveReplaces = G.aboveReplaces,
		context = line * contextLines + S.gap / 2 * (contextLines - 1),
		buttonGap = G.buttonGap,
		button = S.buttonHeight + S.gap,
		bottom = H.bodyPad + H.footer,
	}
	local total = 0
	for _, height in pairs(parts) do
		total = total + height
	end
	return total, parts
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
	-- The source runs on the second line from the icon's left, the whole room under the item.
	return { slot = 0, icon = S.dialogSlotColumn, name = name, nameWidth = width - name,
		source = S.dialogSlotColumn, sourceWidth = width - S.dialogSlotColumn }
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

--- "Following: Arms raid · Warrior · 51 points", nil with nothing followed. With `measure` and
--- `width` (L3.7 item 28: the line was cut at "· 51 po…"), the line that fits: the class word is
--- dropped first, then the build name is shortened by whole words (never mid-word) with an
--- ellipsis; the points are never cut.
function BuildDialog.followingLine(build, data, measure, width)
	if build == nil then
		return nil
	end
	local name = escaped(Follow.buildName(data, build))
	local class = Talents.titleCase(build.classSlug)
	local points = pointsWord(#build.order)
	local function line(shown, withClass)
		return string.format(L.dialogFollowingLine, joined(shown, withClass and class or nil, points))
	end
	local full = line(name, true)
	if type(measure) ~= "function" or type(width) ~= "number" then
		return full
	end
	local function fits(text)
		local size = measure(text)
		return type(size) ~= "number" or size <= width
	end
	if fits(full) then
		return full
	end
	local bare = line(name, false)
	if fits(bare) then
		return bare
	end
	local words = {}
	for word in name:gmatch("%S+") do
		words[#words + 1] = word
	end
	for count = #words - 1, 1, -1 do
		local shortened = line(table.concat(words, " ", 1, count) .. L.ellipsis, false)
		if fits(shortened) then
			return shortened
		end
	end
	return line(#words > 1 and (words[1] .. L.ellipsis) or name, false)
end

--- The chat line after Use: "Now following Arms raid. Take Deflection (1 of 2).", the
--- name's pipes doubled so a site-sent name cannot render as markup.
function BuildDialog.chatLine(name, nextLabel, atLevel)
	local safe = escaped(name)
	if atLevel ~= nil then
		-- No point to spend now: the level the next arrives at, and the plan's talent for it.
		if nextLabel ~= nil then
			return string.format(L.dialogNowFollowingLaterTalent, safe, atLevel, nextLabel)
		end
		return string.format(L.dialogNowFollowingLater, safe, atLevel)
	end
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

--- A time of day in the chat's own clock (Compat.clockText), every printed time's one form.
local function hhmm(fields)
	return Compat.clockText(fields.hour, fields.min)
end

--- How long ago `seconds` was, on the clock's own now: "just now" under a minute, "35 min ago",
--- "1 h 35 min ago", "3 h ago", "2 days ago". Never a time of day.
function BuildDialog.ago(seconds, clock)
	local diff = math.max(0, (clock.now or 0) - seconds)
	if diff < 60 then
		return L.agoJustNow
	elseif diff < 3600 then
		return string.format(L.agoMinutes, math.floor(diff / 60))
	elseif diff < 86400 then
		local hours, minutes = math.floor(diff / 3600), math.floor(diff % 3600 / 60)
		if minutes == 0 then
			return string.format(L.agoHours, hours)
		end
		return string.format(L.agoHoursMinutes, hours, minutes)
	end
	local days = math.floor(diff / 86400)
	return days == 1 and L.agoDay or string.format(L.agoDays, days)
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

--- What the export holds in one line: "Level 24 Hunter · 14 items · 3 professions · your guild and
--- realm are included" (a count of none is left out; "your realm is included" with no guild); nil for
--- a code that is not an export.
function BuildDialog.exportSummary(code)
	if type(code) ~= "string" then
		return nil
	end
	local ok, build = pcall(Codec.decodeFS1, code)
	if not ok or type(build) ~= "table" then
		return nil
	end
	local class = Talents.titleCase(build.classSlug)
	local who = build.level ~= nil and string.format(L.dialogSummaryWho, build.level, class)
		or string.format(L.dialogSummaryWhoNoLevel, class)
	local items = #(build.gearSlots or {})
	local professions = #(build.professions or {})
	return joined(
		who,
		items > 0 and (items == 1 and L.dialogSummaryItemOne or string.format(L.dialogSummaryItems, items)) or nil,
		professions > 0 and (professions == 1 and L.dialogSummaryProfessionOne
			or string.format(L.dialogSummaryProfessions, professions)) or nil,
		build.guild ~= nil and L.dialogSummaryShared or L.dialogSummarySharedRealm
	)
end

--- The ack's line about the export: "Site has your save from 1 h 35 min ago" when the site's copy is at or
--- after this computer's last save (the upload follows the save, so "the same export"
--- reads as not older), "Site has your save from 2 days ago. Yours is newer: log out to send it."
--- when it is older; "Not sent to the site yet · log out to send" with no ack at all (nothing has
--- reached the site); nil for an ack with no readable time.
local function seenLine(ack, savedAt, clock, copiedAt)
	if type(ack) ~= "table" then
		return L.dialogNotSent
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
		return string.format(L.dialogSiteOlder, BuildDialog.ago(seen, clock))
	end
	return string.format(L.dialogSeen, BuildDialog.ago(seen, clock))
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

--- "Sim done · your sim weights are in use" while the weights message is undated (today's
--- contract). With its `generated_at`: "Sim done from the 10 Oct 20:00 copy · …", and "· a new
--- sim is queued" when the ack's `requested_at` is later than it.
local function simDoneLine(generatedAt, ack, clock)
	local generated = BuildDialog.parseUtc(generatedAt)
	if generated == nil then
		-- Today's weights message carries no time, and a request's "queued" stays as its state
		-- after the sim is done: name no copy and no new sim (fix round 7 ruling).
		return L.dialogSimDone
	end
	local kind, time, dated = whenParts(generated, clock)
	local copy = kind == "yesterday" and string.format(L.dialogCopyYesterday, time)
		or string.format(L.dialogCopyThe, dated or time)
	local requested = type(ack) == "table" and BuildDialog.parseUtc(ack.requested_at) or nil
	if requested ~= nil and requested > generated then
		return string.format(L.dialogSimDoneQueued, copy)
	end
	return string.format(L.dialogSimDoneFrom, copy)
end

--- The whole of copy mode: the box's text and whether it is selected, the lead line under
--- the tabs (the copy claim, in the gain green, only after a real Ctrl+C, `copiedAt`), the
--- ack's lines, and the companion's explanation when no ack was ever seen. No button: the
--- box itself is what the player copies from (fix round 2, item 4).
--- `input` = { code, reason, copiedAt, ack, savedAt, clock, simInUse, simGeneratedAt, nextStep,
--- link, simFailed }. With `link` ({ url, name }, /fs link, L4.4) the box holds that character's site
--- page instead of the export, under its own two leads, and nothing about the export follows.
function BuildDialog.copyModel(input)
	local clock = input.clock
	local link = type(input.link) == "table" and input.link or nil
	local code = link ~= nil and link.url or input.code or nil
	local lead
	if not code then
		lead = { text = input.reason or "", hexKey = "muted" }
	elseif input.copiedAt ~= nil then
		lead = { text = link ~= nil and string.format(L.dialogLinkCopied, link.name)
			or string.format(L.dialogYouCopied, hhmm(clock.fields(input.copiedAt))), hexKey = "success" }
	else
		lead = { text = link ~= nil and string.format(L.dialogLinkLead, link.name) or L.dialogSelectedLead,
			hexKey = "body" }
	end
	local ack = link == nil and type(input.ack) == "table" and input.ack or nil
	local companion
	if ack == nil and link == nil then
		companion = { L.dialogNoCompanion, L.dialogCompanionWhat, string.format(L.dialogPasteAtSite, L.siteName) }
	end
	-- While the companion's sim weights are in use, the sim's answer: the copy it is from (the
	-- ack's export time; "an earlier copy" while a newer request is queued), and the chip's own
	-- line as the next step (`nextStep`); else the request's state.
	local sim, simNext
	if link == nil and input.simFailed then
		sim = L.dialogSimFailed
	elseif link == nil and input.simInUse then
		sim = simDoneLine(input.simGeneratedAt, ack, clock)
		if type(input.nextStep) == "string" and input.nextStep ~= "" then
			simNext = string.format(L.dialogNextStep, input.nextStep)
		end
	elseif ack ~= nil then
		sim = simLine(ack, clock)
	end
	return {
		text = code or "",
		-- The box's text colour: a site address is read, so it takes the tooltip text colour
		-- (L4.4 fix round 2, item 6); the export is a code, muted.
		textHex = link ~= nil and "tooltipText" or BuildDialog.COLLAPSED_TEXT,
		selected = code ~= nil and code ~= false,
		lead = lead,
		summary = link == nil and BuildDialog.exportSummary(code) or nil,
		seen = link == nil and seenLine(ack, input.savedAt, clock, input.copiedAt) or nil,
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


--- The "from the site" rows: one per build in `builds` (Follow.inboxBuilds) whose id and
--- code are both unused (`dismissed`, Follow.dismissedSet) and whose code loads for this
--- player, as { id, name, code, from, classToken, detail, target }: the name (the title the site
--- sent, verbatim, else the build's own trees), the detail "Hunter · Beast Mastery 11 · 11 points"
--- and, for a build addressed to a character, the muted line "For Bow Jackzon · Classic Beta PvP ·
--- from the site" (`player` carries the character's `name` and `realm`; a site-wide build says only
--- where it is from). Another class's build is left out without a word (it was sent for another
--- character of the account). The second value is the first other refusal met, for /fs inbox.
function BuildDialog.inboxRows(builds, dismissed, data, player, stamp)
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
					from = entry.from,
					classToken = verdict.build.classSlug:upper(),
					detail = BuildDialog.detailLine(verdict.build, data, true),
					-- "For Leafcall · Ashbringer · 12 Nov" (the day, or who pushed it) under the detail; the
					-- title line ends on "15 points" (2026-11-13: "from the site" said what the heading says).
					target = joined(BuildDialog.rowTarget(entry, player), BuildDialog.rowStamp(entry, stamp)),
					stamp = BuildDialog.rowStamp(entry, stamp),
					points = pointsWord(#verdict.build.order),
				}
			elseif refusal == nil and not otherClass then
				refusal = verdict.reason
			end
		end
	end
	return rows, refusal
end

-- The addon's own builds (spec 2026-11-13 section 1: BuildsPane) ---------------------------

--- The rows the cursor walks: the own rows first, then the site rows shown. Pure.
function BuildDialog.cursorItems(own, shown)
	local items = {}
	for _, entry in ipairs(own or {}) do
		items[#items + 1] = { own = true, entry = entry }
	end
	for _, entry in ipairs(shown or {}) do
		items[#items + 1] = { own = false, entry = entry }
	end
	return items
end

--- Where the cursor starts: the first site row when any wait, else on nothing (0): with the spec
--- undecided, Enter on open must never follow a spec. Pure.
function BuildDialog.startCursor(ownCount, siteCount)
	if (siteCount or 0) > 0 then
		return (ownCount or 0) + 1
	end
	return 0
end

--- The muted line under a row's detail: "For Bow Jackzon · Classic Beta PvP · from the site", the
--- source "pushed by <officer>" for a build an officer pushed (L3.7 item 44). The "For ..." part
--- only for a build addressed to a character (`entry.character`) while the player's name is known;
--- the realm is the one the player is on, which the build was matched to (Follow.offeredTo).
function BuildDialog.rowTarget(entry, player)
	local who, realm
	if entry.character ~= nil and type(player) == "table" and type(player.name) == "string" and player.name ~= "" then
		who = string.format(L.dialogForCharacter, player.name)
		realm = player.realm
	end
	return joined(who, realm)
end

--- When a waiting build came, or who sent it: the day the companion wrote the inbox (`day`, "22 Oct"),
--- "pushed by <officer>" for an officer's push (it carries no day); nil with neither. The rows sit
--- under "From the site", so the line never says it again.
function BuildDialog.rowStamp(entry, day)
	if type(entry.from) == "string" and entry.from ~= "" then
		return string.format(L.dialogFromPush, entry.from)
	end
	if type(day) == "string" and day ~= "" then
		return day
	end
	return nil
end

--- The day the inbox was written ("22 Oct", the player's own calendar), from its `generated_at`
--- stamp; nil when there is none or it is unreadable.
function BuildDialog.inboxDay(inbox, clock)
	local at = type(inbox) == "table" and BuildDialog.parseUtc(inbox.generated_at) or nil
	if at == nil or clock == nil then
		return nil
	end
	local fields = clock.fields(at)
	return Compat.shortDate(string.format("%04d-%02d-%02d", fields.year, fields.month, fields.day))
end

--- `raw` as a row shows it, escaped and cut with an ellipsis to `width` as `measure` measures it
--- (no measure: whole).
function BuildDialog.rowText(raw, measure, width)
	raw = tostring(raw or "")
	local bytes = #raw
	while bytes > 1 and not fitsWidth(cutName(raw, bytes), measure, width) do
		bytes = #cut(raw, bytes - 1)
	end
	return cutName(raw, bytes)
end

--- A row's height: the class icon's, or the three lines (name, detail, target) and the gap
--- between them, whichever is taller, at `line` px a text line.
function BuildDialog.rowHeight(line)
	local S = Theme.SIZES
	return math.max(BuildDialog.SITE_CREST, line * 3 + S.gap)
end

--- The rows paste mode shows: with nothing previewed, at most `limit` and "and 2 more" for the
--- rest, the one at `cursor` (the first by default) as `active`, the highlighted row Enter
--- previews; while any build previews or a pasted code is refused, none (the preview carries the
--- build's own icon, name and detail).
function BuildDialog.rowsView(rows, limit, verdict, cursor)
	if verdict ~= nil then
		return { shown = {} }
	end
	local shown = {}
	for index = 1, math.min(limit, #rows) do
		shown[index] = rows[index]
	end
	local more = #rows > limit and string.format(L.dialogInboxMore, #rows - limit) or nil
	-- A cursor of 0 sits on no site row (it is on an own row, or nowhere); nil means the first.
	local active = (cursor == nil or cursor >= 1) and shown[math.max(1, math.min(cursor or 1, #shown))] or nil
	return { shown = shown, more = more, active = active ~= nil and active.id or nil }
end

--- The cursor's row: a 2 px edge in the body colour at its left (spec 6.4: gold means one thing
--- in this dialog, the build followed). Every row hovers in the secondary buttons' hover ground,
--- the cue that it can be clicked.
BuildDialog.ROW_ACTIVE_EDGE = { width = 2, token = "body" }
BuildDialog.ROW_HOVER = Theme.CONTROL_STATES.secondary.hover

--- The no-keyboard note for /fs diag, once a session.
function BuildDialog.noteNoFocus()
	if BuildDialog.focusNoted then
		return false
	end
	BuildDialog.focusNoted = true
	Theme.note(L.diagDialogFocus)
	return true
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

--- The same clock for the other surfaces that tell a time (the site-status lines, X.3).
BuildDialog.liveClock = liveClock

--- The character logged in, for Codec.loadBuild's class check; nil where the client cannot say.
function BuildDialog.player()
	local ok, slug = pcall(Talents.playerClassSlug)
	if not ok or type(slug) ~= "string" then
		return nil
	end
	local found, realm = pcall(function()
		return type(GetRealmName) == "function" and GetRealmName() or nil
	end)
	return { classSlug = slug, classNames = Compat.classNames(), name = Compat.playerName(),
		realm = found and type(realm) == "string" and realm ~= "" and realm or nil }
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
	if type(current) == "table" and current.tag == "sim" then
		return true, current.generatedAt
	end
	return false
end

--- Whether the newest sim result for this character says it failed (L3.7 item 30: the API sends
--- `failed`), so the copy tab can say to ask again.
local function simFailed()
	local ok, found = pcall(Follow.messages, ForeverSixtyInbox, Export.characterKey(), "sim_result")
	local newest = ok and type(found) == "table" and found[1] or nil
	return type(newest) == "table" and newest.failed ~= nil and newest.failed ~= false
end

--- The chip's own line now, whatever it says (a talent, an item, or its idle line), or nil.
local function chipLine()
	local current = Chip.current
	local chosen = type(current) == "table" and current.chosen or nil
	local text = type(chosen) == "table" and chosen.text or nil
	return type(text) == "string" and text ~= "" and text or nil
end

local function currentAck()
	return Follow.messages(ForeverSixtyInbox, Export.characterKey(), "ack")[1]
end

local function currentRows()
	return BuildDialog.inboxRows(Follow.inboxBuilds(ForeverSixtyInbox, Export.characterKey()),
		Follow.dismissedSet(), dataOf(), BuildDialog.player(),
		BuildDialog.inboxDay(ForeverSixtyInbox, BuildDialog.liveClock()))
end

--- The own builds' view for this character: { classToken, title, note, offPlanTalents, question,
--- blocks }, the followed row marked; nil without a player or data. The title ("Hunter builds") and
--- the note (the level and the points) head the window; the spec question heads the build list while
--- the points do not name a tree.
local function currentOwnView()
	local player = BuildDialog.player()
	local data = dataOf()
	if player == nil or type(data) ~= "table" or type(data.defaultBuilds) ~= "table" then
		return nil
	end
	local level, ranks = liveLevel(), liveRanks(data)
	local names = type(player.classNames) == "table" and player.classNames or {}
	local className = names[player.classSlug:upper()] or Talents.titleCase(player.classSlug)
	local ok, question = pcall(Follow.specQuestion, data, player.classSlug, ranks)
	local has, split = pcall(Follow.specSplit, data, player.classSlug, ranks)
	question, split = ok and question or nil, has and split or nil
	-- PanelTab loads after this file: reached at call time, as the panel's own hover names them.
	local PanelTab = ns.PanelTab or require("PanelTab")
	return {
		classToken = player.classSlug:upper(),
		title = BuildsPane.title(className),
		question = BuildsPane.question(question, split),
		note = BuildsPane.note(level, Follow.build, ranks),
		offPlanTalents = PanelTab.offPlanTalents(Follow.build, ranks, data),
		blocks = BuildsPane.blocks(data, player.classSlug, level, Follow.build, ranks),
	}
end

--- The own view and its rows, read again (open, a refresh, after a follow).
local function readOwn()
	BuildDialog.ownView = currentOwnView()
	BuildDialog.own = BuildsPane.flatten(BuildDialog.ownView ~= nil and BuildDialog.ownView.blocks or nil)
	return BuildDialog.own
end

local SLOT_WORDS = setmetatable({}, {
	__index = function(_, name)
		return Compat.globalString(name)
	end,
})

-- Using a build --------------------------------------------------------------------------

--- Follow `code` (Follow.load saves it in the active slot, refusing another class's), mark a
--- build from the site used by its id and its code, redraw every surface and the window (the list
--- back, the box empty, "Now following ..." under it), and say so in one chat line with the next
--- talent. The window stays open (spec 4).
function BuildDialog.use(code, name, inboxId)
	local data = dataOf()
	-- A paste is adopted into its own slot, never over a restored raid build (L3.7 item 32); a build
	-- the site sent keeps the active slot.
	local slot, from
	if inboxId == nil then
		slot, from = Follow.pasteSlot()
	end
	local build = Follow.load(code, data, name, slot, BuildDialog.player())
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
	local followed = Follow.buildName(data, build)
	BuildDialog.status = { text = string.format(L.dialogNowFollowing, followed), hexKey = "success" }
	BuildDialog.afterFollow()
	local level, ranks = liveLevel(), liveRanks(data)
	local atLevel
	if type(level) == "number" and Progress.unspentPoints(level, ranks) == 0 then
		atLevel = Progress.nextPointLevel(level, ranks)
	end
	local line = BuildDialog.chatLine(followed, liveNext(build), atLevel)
	local say = BuildDialog.say or Theme.say
	say(line)
	if from ~= nil then
		say(string.format(L.dialogUsingPaste, from, from))
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
	local drawn = BuildDialog.draw()
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
	verdict.name, verdict.inboxId, verdict.target, verdict.stamp = row.name, row.id, row.target, row.stamp
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
	return BuildDialog.draw()
end

--- The rows the cursor can sit on now: the own rows, then the site rows shown.
local function currentItems()
	local own = BuildDialog.own or {}
	local limit = Theme.SIZES.dialogInboxRows
	local shown = {}
	for index = 1, math.min(limit, #(BuildDialog.rows or {})) do
		shown[index] = BuildDialog.rows[index]
	end
	return BuildDialog.cursorItems(own, shown)
end

--- After a follow from the window: the preview and the box cleared, the list and the site rows read
--- again (the followed row marked), the cursor kept on the rows there are, the window drawn.
function BuildDialog.afterFollow()
	BuildDialog.verdict, BuildDialog.editing, BuildDialog.touched = nil, false, false
	BuildDialog.rows = currentRows()
	readOwn()
	BuildDialog.cursor = math.max(0, math.min(BuildDialog.cursor or 0, #currentItems()))
	if BuildDialog.paste ~= nil then
		BuildDialog.paste.box:SetText("")
	end
	return BuildDialog.draw()
end

--- The item under the cursor: { own, entry }, or nil when the cursor is on nothing.
function BuildDialog.cursorItem()
	return currentItems()[BuildDialog.cursor or 0]
end

--- The cursor to the `index`th row (kept within the rows there are; 0 is none), and the dialog redrawn.
function BuildDialog.setCursor(index)
	local count = #currentItems()
	BuildDialog.cursor = math.max(count > 0 and 1 or 0, math.min(index, count))
	return BuildDialog.draw()
end

--- Up or Down in the paste box: the cursor moves a row while none is previewed.
function BuildDialog.moveCursor(step)
	if BuildDialog.verdict ~= nil or #currentItems() == 0 then
		return nil
	end
	return BuildDialog.setCursor((BuildDialog.cursor or 0) + step)
end

--- Enter with none previewed: the cursor's own row is followed, a site row previews, nothing
--- happens on no row.
function BuildDialog.previewCursor()
	local item = BuildDialog.cursorItem()
	if item == nil then
		return nil
	end
	if item.own then
		return BuildDialog.followOwn(item.entry)
	end
	return BuildDialog.choose(item.entry)
end

--- One tap on an own-build row (spec 2026-11-13-addon-builds-window section 4): the row already
--- followed does nothing; any other is followed (Chip.choosePath: Follow.chooseDefault with the row's
--- kind when it is not the level's own, the choice remembered for the character), the window redraws
--- with that row gold and Up next refilled, and the chat line says what is followed and the next
--- talent. A default that does not load marks the row.
function BuildDialog.followOwn(entry)
	if entry == nil then
		return nil
	end
	if entry.following then
		return Follow.build
	end
	local data = dataOf()
	Chip.data = Chip.data or data
	local build = Chip.choosePath(entry.spec, entry.choose)
	if build == nil then
		BuildDialog.ownFailed = BuildDialog.ownFailed or {}
		BuildDialog.ownFailed[entry.spec] = true
		BuildDialog.draw()
		return nil
	end
	if BuildDialog.afterUse ~= nil then
		BuildDialog.afterUse()
	end
	BuildDialog.status = nil
	BuildDialog.afterFollow()
	local level, ranks = liveLevel(), liveRanks(data)
	local atLevel
	if type(level) == "number" and Progress.unspentPoints(level, ranks) == 0 then
		atLevel = Progress.nextPointLevel(level, ranks)
	end
	local say = BuildDialog.say or Theme.say
	say(BuildDialog.chatLine(Follow.buildName(data, build), liveNext(build), atLevel))
	return build
end

--- The Use button, and Enter in the paste box: only while a preview shows, and not while the build
--- is a prefix of the plan (`useBlocked`, the reason on the screen). Loads the verdict's own
--- trimmed code, never the box's text.
function BuildDialog.useCurrent()
	local verdict = BuildDialog.verdict
	if verdict == nil or not verdict.ok or BuildDialog.useBlocked ~= nil then
		return nil
	end
	local name = verdict.name or Follow.buildName(dataOf(), verdict.build)
	return BuildDialog.use(verdict.code, name, verdict.inboxId)
end

-- The frame ------------------------------------------------------------------------------

--- The left column's width (the build list, a preview).
local function contentWidth()
	return Theme.SIZES.windowLeft
end

--- The right column's width (Up next, From the site, To the site).
local function rightWidth()
	return Theme.SIZES.windowRight
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
		local probe = text(BuildDialog.body, "body")
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

--- A measure for `region`: the width it draws `value` at.
local function widthOf(region)
	return function(value)
		region:SetText(value)
		return type(region.GetStringWidth) == "function" and region:GetStringWidth() or nil
	end
end

--- What Escape does on the Builds tab (spec section 4): a code previewing or refused goes back to the
--- build list with the paste box empty (true: handled); anything else is the hub's (it closes).
function BuildDialog.escape()
	if BuildDialog.verdict ~= nil then
		BuildDialog.clearPaste()
		return true
	end
	return false
end

--- Escape from either box: the tab's own step back, else the window closes (Hub.escape).
local function escapeFromBox()
	if not BuildDialog.escape() then
		Hub.close()
	end
end

--- The frame's own key handler (X.7): the dialog is keyboard-enabled so Escape closes it whether
--- the focus was in a box, on a row or on a tab (UISpecialFrames alone only closes it while
--- nothing else has taken the key). Escape is consumed; every other key is passed on, so Return
--- still opens chat and movement keys still walk. SetPropagateKeyboardInput is the client's call
--- for exactly this; it is a flag on our own frame, which is never protected, so it is allowed in
--- combat, and nothing here touches a client frame.
function BuildDialog.onKey(frame, key)
	local keep = key ~= "ESCAPE"
	if type(frame.SetPropagateKeyboardInput) == "function" then
		frame:SetPropagateKeyboardInput(keep)
	end
	if not keep then
		BuildDialog.escape()
	end
end

--- Escape, Enter, Tab and the arrows in either box: Escape as BuildDialog.escape, Tab hands the
--- keyboard to the other box (the copy box selects its code), Enter does the box's job.
local function bindKeys(box, onEnter)
	box:SetScript("OnEscapePressed", escapeFromBox)
	box:SetScript("OnTabPressed", function()
		BuildDialog.focusBox(BuildDialog.focus == "copy" and "paste" or "copy")
	end)
	-- Up and Down move the cursor across the build rows and the site rows (nothing previewed).
	box:SetScript("OnArrowPressed", function(_, key)
		if key == "UP" then
			BuildDialog.moveCursor(-1)
		elseif key == "DOWN" then
			BuildDialog.moveCursor(1)
		end
	end)
	box:SetScript("OnEnterPressed", onEnter)
end

--- A box set in a field: the field (the border, what a caller anchors), a clip of whole
--- lines inside it (so no glyph tops of a next line show), the box in the clip.
local function boxInField(parent, readOnly, width)
	local S = Theme.SIZES
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

--- One planned item in the preview, slot first (BuildDialog.pickRowLayout): the slot word
--- muted in its fixed column, the icon, then the name in its quality colour; hovering
--- shows the item.
local function buildPickRow(parent, width)
	local S = Theme.SIZES
	local layout = BuildDialog.pickRowLayout(width)
	local row = Theme.createFrame("Frame", nil, parent)
	row:SetSize(width, BuildDialog.GEAR_PITCH)
	row:EnableMouse(true)
	-- First line: the slot word muted in its column, the icon, the item on its quality colour.
	local slot = text(row, "muted", S.dialogSlotColumn - S.gap)
	slot:SetPoint("TOPLEFT", row, "TOPLEFT", layout.slot, -S.gap / 2)
	if type(slot.SetWordWrap) == "function" then
		slot:SetWordWrap(false)
	end
	local icon = Theme.icon(row, "ARTWORK", nil, S.iconSize)
	icon:SetPoint("TOPLEFT", row, "TOPLEFT", layout.icon, -S.gap / 2)
	local name = text(row, "body", layout.nameWidth)
	name:SetPoint("TOPLEFT", row, "TOPLEFT", layout.name, -S.gap / 2)
	if type(name.SetWordWrap) == "function" then
		name:SetWordWrap(false)
	end
	-- Second line: where the item comes from, muted and whole (BuildDialog.sourceText).
	local source = text(row, "muted", layout.sourceWidth)
	source:SetPoint("TOPLEFT", row, "TOPLEFT", layout.source, -(S.gap / 2 + S.iconSize + 1))
	if type(source.SetWordWrap) == "function" then
		source:SetWordWrap(false)
	end
	local pick = { frame = row, icon = icon, name = name, slot = slot, source = source, nameRoom = layout.nameWidth,
		sourceRoom = layout.sourceWidth }
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
	-- The ringed HD class crest (Theme.buildCrest), as every class image in the addon and on the site.
	local icon, ring = Theme.buildCrest(preview, S.dialogClassIcon)
	icon:SetPoint("TOPLEFT", preview, "TOPLEFT", 0, 0)
	local name = text(preview, "gold", inner)
	name:SetPoint("TOPLEFT", icon, "TOPRIGHT", S.gap * 2, 0)
	if type(name.SetWordWrap) == "function" then
		name:SetWordWrap(false)
	end
	local detail = text(preview, "muted", inner)
	detail:SetPoint("TOPLEFT", name, "BOTTOMLEFT", 0, -S.gap / 2)
	setMaxLines(detail, 2)
	-- A build from the site: "For Bow Jackzon · Classic Beta PvP · from the site", muted.
	local target = text(preview, "muted", inner)
	target:SetPoint("TOPLEFT", detail, "BOTTOMLEFT", 0, -S.gap / 2)
	if type(target.SetWordWrap) == "function" then
		target:SetWordWrap(false)
	end
	target:Hide()
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
	-- The gear rows: GEAR_VISIBLE of them, over the list the wheel scrolls (every item is a row).
	local picks = {}
	for index = 1, BuildDialog.GEAR_VISIBLE do
		picks[index] = buildPickRow(preview, width)
	end
	local area = Theme.createFrame("Frame", nil, preview)
	area:SetSize(width, BuildDialog.GEAR_VISIBLE * BuildDialog.GEAR_PITCH)
	local bar = Widgets.scrollbar(area, area, Theme.SIZES.gap)
	local result = {
		frame = preview, icon = icon, ring = ring, name = name, detail = detail, target = target, placeNow = placeNow,
		headingFrame = headingFrame, heading = heading, picks = picks, area = area, bar = bar,
		offset = 0, entries = {},
	}
	local function wheel(_, delta)
		BuildDialog.scrollGear(-delta)
	end
	headingFrame:EnableMouseWheel(true)
	headingFrame:SetScript("OnMouseWheel", wheel)
	for _, pick in ipairs(picks) do
		pick.frame:EnableMouseWheel(true)
		pick.frame:SetScript("OnMouseWheel", wheel)
	end
	preview:Hide()
	return result
end

--- One build from the site as a row (X.5): the class icon, the name, the detail ("Hunter · Beast
--- Mastery 11 · 11 points") and the muted "For Bow Jackzon · Classic Beta PvP · from the site" under
--- it, the same arrangement as the preview's header. The highlighted row has a gold edge outside its
--- left side and the hover ground; hovering a row highlights it, clicking previews it.
local function buildRow(content, width)
	local S = Theme.SIZES
	local inner = width - BuildDialog.SITE_CREST - S.gap * 2
	local row = { frame = Theme.createFrame("Button", nil, content) }
	row.frame:SetSize(width, BuildDialog.rowHeight(lineHeight()))
	local highlight = Theme.texture(row.frame, "BACKGROUND", BuildDialog.ROW_HOVER)
	highlight:SetAllPoints(row.frame)
	highlight:Hide()
	row.highlight = highlight
	local E = BuildDialog.ROW_ACTIVE_EDGE
	row.edge = Theme.texture(row.frame, "ARTWORK", E.token)
	row.edge:SetPoint("TOPRIGHT", row.frame, "TOPLEFT", -S.gap, 0)
	row.edge:SetPoint("BOTTOMRIGHT", row.frame, "BOTTOMLEFT", -S.gap, 0)
	row.edge:SetWidth(E.width)
	row.edge:Hide()
	row.icon, row.ring = Theme.buildCrest(row.frame, BuildDialog.SITE_CREST)
	row.icon:SetPoint("TOPLEFT", row.frame, "TOPLEFT", 0, 0)
	row.name = text(row.frame, "body", inner)
	row.name:SetPoint("TOPLEFT", row.icon, "TOPRIGHT", S.gap * 2, 0)
	row.detail = text(row.frame, "muted", inner)
	row.detail:SetPoint("TOPLEFT", row.name, "BOTTOMLEFT", 0, -S.gap / 2)
	row.target = text(row.frame, "muted", inner)
	row.target:SetPoint("TOPLEFT", row.detail, "BOTTOMLEFT", 0, -S.gap / 2)
	for _, region in ipairs({ row.name, row.detail, row.target }) do
		if type(region.SetWordWrap) == "function" then
			region:SetWordWrap(false)
		end
	end
	row.frame:SetScript("OnEnter", function()
		if row.index ~= nil and BuildDialog.cursor ~= row.index then
			BuildDialog.setCursor(row.index)
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

--- A column of the window: `x` from the frame's left, `width` wide, between the header and the footer.
local function column(parent, x, width, top, bottom)
	local area = Theme.createFrame("Frame", nil, parent)
	area:SetPoint("TOPLEFT", parent, "TOPLEFT", x, -top)
	area:SetPoint("BOTTOMLEFT", parent, "BOTTOMLEFT", x, bottom)
	area:SetWidth(width)
	return area
end

--- The window's text maker and line fitter, handed to BuildsPane and BuildsWindow.
local function makeText(parent, hexKey, width)
	return text(parent, hexKey, width)
end

--- Cut with an ellipsis to the room; a line carrying its own colour escapes (the question's split,
--- a spec's playstyle, a rank) is never cut by bytes: the client cuts it at the width.
local function fitLine(region, raw, room)
	region:SetWidth(room)
	if tostring(raw or ""):find("|c", 1, true) then
		region:SetText(raw)
	else
		region:SetText(BuildDialog.rowText(raw, widthOf(region), room))
	end
end

--- The paste side of the window: on the left the lead lines, the build list and the preview with
--- its Use button; on the right Up next, then From the site (its rows, the paste box, the refusal
--- or the status line under it).
local function buildPaste(body, top, bottom)
	local S = Theme.SIZES
	local width = contentWidth()
	local content = column(body, S.padding, width, top, bottom)
	local right = column(body, S.padding + width + S.windowGutter, rightWidth(), top, bottom)
	local paste = { content = content, right = right, rows = {}, own = {} }
	paste.following = text(content, "body", width)
	if type(paste.following.SetWordWrap) == "function" then
		paste.following:SetWordWrap(false)
	end
	paste.next = text(content, "muted", width)
	-- The own builds (BuildsPane), then the "From the site" heading over the site rows and the box.
	paste.builds = BuildsPane.build(content, width, {
		text = function(parent, hexKey)
			return text(parent, hexKey)
		end,
		fit = fitLine,
		follow = function(entry)
			return BuildDialog.followOwn(entry)
		end,
		hovered = function(index)
			if index ~= nil and BuildDialog.cursor ~= index then
				BuildDialog.setCursor(index)
			end
		end,
		hover = BuildDialog.ROW_HOVER, edge = BuildDialog.ROW_ACTIVE_EDGE,
	})
	paste.noBuilds = text(content, "muted", width)
	paste.noBuilds:SetText(L.windowNoBuilds)
	paste.upNext = BuildsWindow.buildUpNext(right, rightWidth(), { text = makeText })
	paste.site = BuildsWindow.buildSection(right, rightWidth(), { text = makeText }, L.dialogSiteHeading,
		L.windowFromSiteWhat)
	for index = 1, S.dialogInboxRows do
		paste.rows[index] = buildRow(right, rightWidth())
	end
	paste.more = text(right, "muted", rightWidth())
	local box, field, clip = boxInField(right, false, rightWidth())
	-- Ctrl+C here is not a copy out of the window; the box keeps the keyboard until Escape.
	box:SetScript("OnKeyDown", nil)
	paste.placeholder = text(clip, "muted", rightWidth() - S.gap * 4)
	paste.placeholder:SetText(L.dialogPaste)
	paste.placeholder:SetPoint("TOPLEFT", clip, "TOPLEFT", 0, 0)
	-- Once a code loads, the box reads as one muted line with an ellipsis; a click opens it.
	paste.display = text(field, "muted", rightWidth() - S.gap * 4)
	paste.display:SetPoint("LEFT", field, "LEFT", S.gap * 2, 0)
	if type(paste.display.SetWordWrap) == "function" then
		paste.display:SetWordWrap(false)
	end
	box:SetScript("OnTextChanged", function(self, userInput)
		if userInput then
			BuildDialog.touched = true
			BuildDialog.onPaste(self:GetText())
		else
			BuildDialog.draw()
		end
	end)
	local function startEditing()
		if BuildDialog.touched ~= true then
			-- A click in the box is the player working in it: its border goes gold (X.5).
			BuildDialog.touched = true
			BuildDialog.draw()
		end
		if not BuildDialog.editing and BuildDialog.verdict ~= nil then
			BuildDialog.editing = true
			BuildDialog.draw()
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
		BuildDialog.focus = "paste"
		BuildDialog.draw()
	end)
	box:SetScript("OnEditFocusLost", function()
		BuildDialog.focused = false
		if BuildDialog.editing then
			-- Done editing: a refusal shows again.
			BuildDialog.editing = false
		end
		BuildDialog.draw()
	end)
	bindKeys(box, function()
		if BuildDialog.editing and not (BuildDialog.verdict or {}).ok then
			BuildDialog.editing = false
			BuildDialog.draw()
			return
		end
		if BuildDialog.verdict == nil then
			-- Nothing is previewed: Enter acts on the cursor's row (an own row follows, a site row
			-- previews, no row does nothing).
			BuildDialog.previewCursor()
			return
		end
		BuildDialog.useCurrent()
	end)
	paste.box, paste.field, paste.clip = box, field, clip
	-- Under the box: a refusal's three lines, or the status line ("Read: ...", "Now following ...").
	paste.lead = text(right, "warning", rightWidth())
	setMaxLines(paste.lead, 2)
	paste.body = text(right, "muted", rightWidth())
	setMaxLines(paste.body, 2)
	paste.still = text(right, "muted", rightWidth())
	paste.status = text(right, "success", rightWidth())
	setMaxLines(paste.status, 2)
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
	-- Above the button, laid out bottom-up by drawPreview: the next talent the candidate gives (wrapped
	-- whole), what Use replaces and, above it, where the candidate stops short; the last two
	-- wrap whole (X.5). Beside the button, why it is held back.
	paste.nextWould = text(content, "muted", width)
	-- The only warning before a plan changes: the warning colour, not the muted body.
	paste.replaces = text(content, "warning", width)
	paste.short = text(content, "muted", width)
	paste.useReason = text(content, "muted", width - Theme.SIZES.buttonWidth - S.gap * 2)
	paste.useReason:SetPoint("LEFT", paste.use, "RIGHT", S.gap * 2, 0)
	paste.useReason:Hide()
	return paste
end

--- Show or hide the selection ground behind the export, as the box's selection stands.
local function paintSelection()
	local copy = BuildDialog.copyBox
	if copy == nil then
		return
	end
	if BuildDialog.selected and BuildDialog.exportCode ~= nil then
		copy.selection:Show()
	else
		copy.selection:Hide()
	end
end

--- To the site (spec 3.3 item 3), on the right column `content`: the heading and its explanation, the
--- copy box (BuildsWindow.COPY_LINES whole lines of the code, the rest scrolls), the status line under
--- it, then the export's facts.
local function buildCopy(content)
	local S = Theme.SIZES
	local width = rightWidth()
	local copy = { content = content }
	copy.section = BuildsWindow.buildSection(content, width, { text = makeText }, L.windowToSite, L.windowToSiteWhat)
	copy.lead = text(content, "body", width)
	setMaxLines(copy.lead, 2)
	local box, field, clip = boxInField(content, true, width)
	copy.box, copy.field, copy.clip = box, field, clip
	-- The box holds the exact export and takes Ctrl+C, unseen; what the player reads is the
	-- display over it, which wraps a run of characters with no spaces at the edge (an edit
	-- box wraps only at spaces, and the export has none).
	box:SetAlpha(0)
	-- Its selection highlight would draw over the label (round 4: a grey block): never drawn.
	if type(box.SetHighlightColor) == "function" then
		box:SetHighlightColor(Theme.rgb(Theme.HEX.shadow, 0))
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
	-- The one copy claim the dialog makes comes from a real Ctrl+C in this box. Its own handler,
	-- not chained: Widgets' release after Ctrl+C would drop the keyboard and the selection (fix
	-- round 5, item 8); a second Ctrl+C copies again, Escape closes.
	box:SetScript("OnKeyDown", function(self, key)
		local control = type(IsControlKeyDown) == "function" and IsControlKeyDown() == true
		local focused = type(self.HasFocus) == "function" and self:HasFocus() == true
		if control and key == "A" then
			BuildDialog.selected = true
		elseif BuildDialog.copyKey(key, control, focused, BuildDialog.selected == true) then
			BuildDialog.copiedAt = wallNow()
			BuildDialog.draw(false)
		end
	end)
	chain(box, "OnMouseDown", function()
		-- The press places the cursor: the export is no longer all selected...
		BuildDialog.selected = false
		paintSelection()
		BuildDialog.paintCopyStatus()
	end)
	chain(box, "OnMouseUp", function()
		-- ...and the click ends selecting it all again, for Ctrl+C.
		BuildDialog.copy()
	end)
	chain(box, "OnEditFocusGained", function()
		BuildDialog.focus = "copy"
		paintEdges(field.foreverSixtyBorder, "gold")
	end)
	chain(box, "OnEditFocusLost", function()
		BuildDialog.selected = false
		paintSelection()
		BuildDialog.paintCopyStatus()
		paintEdges(field.foreverSixtyBorder, "border")
	end)
	-- The whole export, wrapped; the wheel scrolls it a line at a time, with a thumb.
	clip:EnableMouseWheel(true)
	clip:SetScript("OnMouseWheel", function(_, delta)
		BuildDialog.scrollCopy(-delta)
	end)
	-- The Theme's thin scrollbar inside the field's right edge, clear of the text.
	copy.bar = Widgets.scrollbar(field, field, -(S.scrollbar + S.gap / 2))
	copy.summary = text(content, "muted", width)
	setMaxLines(copy.summary, 2)
	copy.summary:Hide()
	copy.seen = text(content, "muted", width)
	setMaxLines(copy.seen, 2)
	copy.sim = text(content, "muted", width)
	copy.simNext = text(content, "muted", width)
	copy.companion = {}
	for index = 1, 3 do
		copy.companion[index] = text(content, "muted", width)
	end
	return copy
end

local function onEvent(_, event, itemId)
	if event == "GET_ITEM_INFO_RECEIVED" and itemId ~= nil and BuildDialog.pending[itemId] then
		BuildDialog.pending[itemId] = nil
		if Hub.isShown() and Hub.active == "builds" then
			BuildDialog.draw(false)
		end
	end
end

--- The Builds tab's body, built into the hub's content frame `body` on the window's first open: the
--- two columns with a hairline between them (spec section 3).
function BuildDialog.buildInto(body)
	local S, H = Theme.SIZES, Hub.LAYOUT
	BuildDialog.body, BuildDialog.frame = body, Hub.frame
	local divider = Theme.texture(body, "ARTWORK", "border")
	divider:SetWidth(1)
	local ruleX = S.padding + contentWidth() + S.windowGutter / 2
	divider:SetPoint("TOPLEFT", body, "TOPLEFT", ruleX, 0)
	divider:SetPoint("BOTTOMLEFT", body, "BOTTOMLEFT", ruleX, 0)
	BuildDialog.paste = buildPaste(body, H.bodyPad, H.bodyPad)
	BuildDialog.copyBox = buildCopy(BuildDialog.paste.right)
	body:SetScript("OnEvent", onEvent)
	Theme.registerEvent(body, "GET_ITEM_INFO_RECEIVED")
	return body
end

--- The tab leaving the screen (another tab, or the window hiding): the keyboard given back, a copy
--- claim and a link ended, the status line cleared.
function BuildDialog.onHide()
	if BuildDialog.paste ~= nil then
		BuildDialog.paste.box:ClearFocus()
		BuildDialog.copyBox.box:ClearFocus()
	end
	BuildDialog.copiedAt = nil
	BuildDialog.link = nil
	BuildDialog.status = nil
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


--- The row's source beside its name (L3.7 item 31): the item tooltip's own text fitted to what the
--- name leaves (at most 60% of the row is the name's), the boss never cut (ItemTooltip.fitSource);
--- hidden where no pick names the item or nothing fits.
function BuildDialog.fitSource(pick, entry, room)
	room = room or pick.sourceRoom
	pick.source:SetWidth(room)
	pick.source:SetText(BuildDialog.fitSourceText(entry.source or L.dialogSourceUnknown, widthOf(pick.source), room))
	pick.source:Show()
end

--- Draw the gear rows from `preview.entries` at `preview.offset` (the wheel's): each row its slot,
--- icon, item on its quality colour and, under it, the whole source; the scrollbar beside them when
--- the list is longer than the rows shown. The height the rows take (the window's, whole rows).
function BuildDialog.paintGear()
	local paste = BuildDialog.paste
	local preview = paste ~= nil and paste.preview or nil
	if preview == nil then
		return 0
	end
	local S = Theme.SIZES
	local total, visible = #preview.entries, BuildDialog.GEAR_VISIBLE
	local scrolls = total > visible
	local shown = math.min(total, visible)
	local width = contentWidth() - (scrolls and S.scrollbar + S.gap * 2 or 0)
	local layout = BuildDialog.pickRowLayout(width)
	preview.area:ClearAllPoints()
	preview.area:SetPoint("TOPLEFT", preview.frame, "TOPLEFT", 0, -(preview.top or 0))
	preview.area:SetSize(width, math.max(1, shown * BuildDialog.GEAR_PITCH))
	if total > 0 then
		preview.area:Show()
	else
		preview.area:Hide()
	end
	for index, pick in ipairs(preview.picks) do
		local entry = preview.entries[(preview.offset or 0) + index]
		pick.itemId = entry ~= nil and entry.itemId or nil
		if entry == nil or index > shown then
			pick.frame:Hide()
		else
			pick.frame:SetWidth(width)
			pick.nameRoom, pick.sourceRoom = layout.nameWidth, layout.sourceWidth
			pick.name:SetWidth(layout.nameWidth)
			pick.slot:SetText(entry.slotWord)
			pick.icon:SetTexture(Compat.itemIcon(entry.itemId) or Theme.UNKNOWN_ICON)
			local name = Compat.itemName(entry.itemId)
			if name == nil then
				BuildDialog.pending[entry.itemId] = true
			end
			pick.name:SetText(name or "")
			pick.name:SetTextColor(Theme.qualityColor(Compat.itemQuality(entry.itemId)))
			BuildDialog.fitSource(pick, entry)
			pick.frame:ClearAllPoints()
			pick.frame:SetPoint("TOPLEFT", preview.frame, "TOPLEFT", 0,
				-((preview.top or 0) + (index - 1) * BuildDialog.GEAR_PITCH))
			pick.frame:Show()
		end
	end
	local range = math.max(0, total - visible)
	preview.bar:Update(total > 0 and math.min(1, visible / total) or 1, range > 0 and (preview.offset or 0) / range or 0)
	return shown * BuildDialog.GEAR_PITCH
end

--- The wheel over the gear list: `rows` rows down (negative: up), within the list. The new offset.
function BuildDialog.scrollGear(rows)
	local paste = BuildDialog.paste
	local preview = paste ~= nil and paste.preview or nil
	if preview == nil then
		return 0
	end
	local range = math.max(0, #preview.entries - BuildDialog.GEAR_VISIBLE)
	preview.offset = math.max(0, math.min(range, (preview.offset or 0) + rows))
	BuildDialog.paintGear()
	return preview.offset
end

local function drawPreview(verdict, y)
	local S = Theme.SIZES
	local paste = BuildDialog.paste
	local preview = paste.preview
	local data = dataOf()
	local ranks = liveRanks(data)
	local professions = {}
	local known, slugs = pcall(Export.professionSlugs)
	for _, slug in ipairs(known and slugs or {}) do
		professions[slug] = true
	end
	local model = BuildDialog.preview(verdict.build, data, {
		level = liveLevel(),
		unspent = Progress.unspentPoints(liveLevel(), ranks),
		professions = professions,
		slotWord = function(slot)
			return BuildDialog.slotWord(slot, SLOT_WORDS)
		end,
	})
	Theme.paintCrest(preview.icon, preview.ring, model.classToken)
	if verdict.inboxId ~= nil then
		preview.name:SetText(BuildDialog.siteTitle(model.rawName, pointsWord(#verdict.build.order),
			widthOf(preview.name), contentWidth() - S.dialogClassIcon - S.gap * 2))
		model.detail = BuildDialog.detailLine(verdict.build, data, true)
	else
		preview.name:SetText(model.name)
	end
	preview.detail:SetText(model.detail)
	local textY = textHeight(preview.name) + S.gap / 2 + textHeight(preview.detail)
	-- A build from the site names the character it is for, as its row did.
	local lastLine = preview.detail
	if type(verdict.target) == "string" and verdict.target ~= "" then
		preview.target:SetText(BuildDialog.rowText(verdict.target, widthOf(preview.target),
			contentWidth() - S.dialogClassIcon - S.gap * 2))
		preview.target:Show()
		textY = textY + S.gap / 2 + textHeight(preview.target)
		lastLine = preview.target
	else
		preview.target:Hide()
	end
	if model.placeNow ~= nil then
		preview.placeNow:SetText(model.placeNow)
		preview.placeNow:ClearAllPoints()
		preview.placeNow:SetPoint("TOPLEFT", lastLine, "BOTTOMLEFT", 0, -S.gap / 2)
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
	preview.entries, preview.top = model.picks, headY
	-- A different build starts at the top of its list; a redraw of the same one (an item's name arriving) keeps
	-- where the wheel left it.
	if preview.scrolledFor ~= verdict then
		preview.offset, preview.scrolledFor = 0, verdict
	end
	preview.offset = math.min(preview.offset or 0, math.max(0, #model.picks - BuildDialog.GEAR_VISIBLE))
	headY = headY + BuildDialog.paintGear()
	preview.frame:SetHeight(headY)
	stack(preview.frame, paste.content, y, headY)
	-- What Use replaces and how the orders compare, and the next talent the candidate gives.
	-- A paste goes to the leveling slot while the raid slot is active, so it replaces that slot's
	-- build (or nothing), never the raid build Follow.build holds.
	local followed = Follow.build
	if verdict.inboxId == nil then
		local slot, from = Follow.pasteSlot()
		if from ~= nil then
			followed = Follow.slotBuild(data, slot)
		end
	end
	local ok, context = pcall(BuildDialog.previewContext, followed, verdict.build, ranks, data, liveLevel())
	if not ok then
		context = {}
	end
	-- The three lines above the button, bottom-up from it: the next talent, what Use replaces, and
	-- where the build stops short. The sentences wrap whole; their height is what the dialog grows by.
	local G = BuildDialog.LAYOUT
	local anchor, gap, contextHeight = paste.use, G.buttonGap, 0
	for _, key in ipairs({ "nextWould", "replaces", "short" }) do
		local region = paste[key]
		if context[key] ~= nil then
			region:SetText(context[key])
			region:ClearAllPoints()
			region:SetPoint("BOTTOMLEFT", anchor, "TOPLEFT", 0, gap)
			region:Show()
			contextHeight = contextHeight + (anchor == paste.use and 0 or S.gap / 2) + textHeight(region)
			anchor, gap = region, S.gap / 2
		else
			region:SetText("")
			region:Hide()
		end
	end
	-- A build that is a prefix of the plan cannot be used, and says why beside the button.
	BuildDialog.useBlocked = context.blocked
	Widgets.setEnabled(paste.use, context.blocked == nil)
	if context.blocked ~= nil then
		paste.useReason:SetText(context.blocked)
		paste.useReason:Show()
	else
		paste.useReason:Hide()
	end
	return y + headY + G.aboveReplaces + contextHeight + G.buttonGap + S.buttonHeight + S.gap
end

--- The hub's header for the Builds tab (spec 3.1): the class crest, "Hunter builds" and the panel's
--- figures, the off-plan talents on the figures' hover.
local function drawHeader()
	local view = BuildDialog.ownView
	local player = BuildDialog.player()
	Hub.setHeader({
		classToken = view ~= nil and view.classToken or (player ~= nil and player.classSlug:upper() or nil),
		title = view ~= nil and view.title or nil,
		note = view ~= nil and view.note or nil,
		noteLines = view ~= nil and view.offPlanTalents or nil,
	})
end

--- The left column (spec 3.2): the lead lines for a build that is not a row, the build list (the spec
--- question over it while the points name no tree), or the preview of a code with its Use button.
--- The y it used.
local function drawLeft(view, verdict)
	local S = Theme.SIZES
	local paste = BuildDialog.paste
	local content = paste.content
	local data = dataOf()
	for _, region in ipairs({ paste.following, paste.next, paste.replaces, paste.short, paste.nextWould,
		paste.useReason, paste.noBuilds }) do
		region:Hide()
	end
	paste.preview.frame:Hide()
	-- Never hidden and shown again in one draw: the mouse-down on Use takes the box's focus and
	-- redraws, and a Button hidden while held loses its click (BuildsPane.draw's own rule).
	Theme.setShown({ paste.use }, view.preview == true)
	BuildDialog.useBlocked = nil
	if view.preview then
		BuildsPane.draw(paste.builds, nil, 0, {})
		return drawPreview(verdict, 0)
	end
	local own = BuildDialog.own or {}
	local followedOwn = false
	for _, entry in ipairs(own) do
		followedOwn = followedOwn or entry.following == true
	end
	local y = 0
	-- "Following: ..." only for a build that is not one of the rows (a paste, a site build).
	local line = not followedOwn
		and BuildDialog.followingLine(Follow.build, data, widthOf(paste.following), contentWidth()) or nil
	if line ~= nil then
		paste.following:SetText(line)
		y = stack(paste.following, content, y, textHeight(paste.following) + S.gap / 2)
		local talentLine = BuildDialog.nextTalentLine(liveDialogNext(Follow.build))
		if talentLine ~= nil then
			paste.next:SetText(talentLine)
			y = stack(paste.next, content, y, textHeight(paste.next))
		end
		y = y + S.gap * 2
	end
	y = BuildsPane.draw(paste.builds, BuildDialog.ownView, y,
		{ line = lineHeight(), cursor = BuildDialog.cursor, failed = BuildDialog.ownFailed })
	if #own == 0 then
		y = stack(paste.noBuilds, content, y, textHeight(paste.noBuilds))
	end
	return y
end

--- The site rows under From the site, at `y`: the crest, the name and its points, the split, who it is
--- for; "and N more" past dialogInboxRows. The y under them.
local function drawSiteRows(verdict, y)
	local S = Theme.SIZES
	local paste = BuildDialog.paste
	local own = BuildDialog.own or {}
	local siteCursor = (BuildDialog.cursor or 0) - #own
	local rows = BuildDialog.rowsView(BuildDialog.rows or {}, S.dialogInboxRows, verdict,
		#own > 0 and siteCursor or BuildDialog.cursor)
	local rowHeight = BuildDialog.rowHeight(lineHeight())
	local inner = rightWidth() - BuildDialog.SITE_CREST - S.gap * 2
	for index, row in ipairs(paste.rows) do
		local entry = rows.shown[index]
		row.entry, row.index = entry, entry ~= nil and (#own + index) or nil
		if entry == nil then
			row.frame:Hide()
		else
			Theme.paintCrest(row.icon, row.ring, entry.classToken)
			row.selected = rows.active == entry.id
			-- The cursor never turns a name gold: gold is the build followed (spec 6.4).
			row.name:SetTextColor(Theme.rgb(Theme.HEX.body))
			if entry.points ~= nil then
				row.name:SetText(BuildDialog.siteTitle(entry.name, entry.points, widthOf(row.name), inner))
			else
				row.name:SetText(BuildDialog.rowText(entry.name, widthOf(row.name), inner))
			end
			row.detail:SetText(BuildDialog.rowText(entry.detail, widthOf(row.detail), inner))
			row.target:SetText(BuildDialog.rowText(entry.target, widthOf(row.target), inner))
			Theme.setShown({ row.edge, row.highlight }, row.selected)
			y = stack(row.frame, paste.right, y, rowHeight + S.gap)
		end
	end
	if rows.more ~= nil then
		-- Under the site rows' text, not their crests.
		paste.more:SetText(rows.more)
		paste.more:ClearAllPoints()
		paste.more:SetPoint("TOPLEFT", paste.right, "TOPLEFT", BuildDialog.SITE_CREST + S.gap * 2, -y)
		paste.more:Show()
		y = y + textHeight(paste.more) + BuildDialog.MORE_GAP
	else
		paste.more:Hide()
		if #rows.shown > 0 then
			y = y + S.gap
		end
	end
	return y
end

--- The paste box at `y`, in whole lines: two while empty or edited, one muted line once a code loads.
--- A site row previewing keeps it unseen and deaf to the mouse (it still holds the keyboard for Enter
--- and Escape). The y under it.
local function drawPasteBox(view, verdict, y)
	local S = Theme.SIZES
	local paste = BuildDialog.paste
	local first, advance = lineMetrics()
	local boxText = paste.box:GetText() or ""
	local fieldHeight = sizeField(paste.field, paste.clip, paste.box, first + (view.lines - 1) * advance)
	local fromSite = verdict ~= nil and verdict.inboxId ~= nil
	paste.field:SetAlpha(fromSite and 0 or 1)
	paste.field:EnableMouse(not fromSite)
	y = stack(paste.field, paste.right, y, fieldHeight)
	if fromSite then
		-- Nothing under an unseen box: the status line would hang in the air.
		y = y - fieldHeight
	end
	paste.box:EnableMouse(not view.collapsed)
	-- The selection under a collapsed line is kept (a next Ctrl+V replaces it) but never drawn.
	if type(paste.box.SetHighlightColor) == "function" then
		if view.hideSelection then
			paste.box:SetHighlightColor(Theme.rgb(Theme.HEX.shadow, 0))
		else
			local r, g, b = Theme.rgb(Theme.HEX.gold)
			paste.box:SetHighlightColor(r, g, b, 0.35)
		end
	end
	if view.collapsed then
		paste.box:SetAlpha(0)
		local shown = verdict.code or boxText:match("^%s*(.-)%s*$"):gsub("%s+", " ")
		paste.display:SetTextColor(Theme.rgb(Theme.HEX[BuildDialog.COLLAPSED_TEXT]))
		paste.display:SetText(Chip.fit(shown, widthOf(paste.display), rightWidth() - S.gap * 4))
		paste.display:Show()
		paste.placeholder:Hide()
	else
		paste.box:SetAlpha(1)
		paste.display:Hide()
		Theme.setShown({ paste.placeholder }, boxText == "")
	end
	paintEdges(paste.field.foreverSixtyBorder, view.border)
	return y
end

--- Under the paste box, at `y`: a refusal's lines (the codec's first sentence in the warning colour,
--- the rest muted, "You are still following ..."), else "Read: <name>. Check it on the left." while a
--- paste previews, else the status line a follow left ("Now following ..."). The y under them.
local function drawUnderBox(view, verdict, y)
	local S, W = Theme.SIZES, BuildsWindow.LAYOUT
	local paste = BuildDialog.paste
	local data = dataOf()
	for _, region in ipairs({ paste.lead, paste.body, paste.still, paste.status }) do
		region:Hide()
	end
	if view.reason then
		local followed = Follow.build ~= nil and Follow.buildName(data, Follow.build) or nil
		local refusal = BuildDialog.refusal(verdict.reason, followed and followed:gsub("|", "||") or nil)
		paste.lead:SetText(refusal.lead)
		y = stack(paste.lead, paste.right, y + W.statusGap, textHeight(paste.lead))
		if refusal.body ~= nil then
			paste.body:SetText(refusal.body)
			y = stack(paste.body, paste.right, y + S.gap / 2, textHeight(paste.body))
		end
		if refusal.following ~= nil then
			paste.still:SetText(refusal.following)
			y = stack(paste.still, paste.right, y + S.gap / 2, textHeight(paste.still))
		end
		return y
	end
	local status = BuildDialog.status
	if view.preview and verdict.inboxId == nil then
		local name = verdict.name or Follow.buildName(data, verdict.build)
		status = { text = string.format(L.windowRead, name), hexKey = "success" }
	elseif view.preview then
		status = nil
	end
	if status ~= nil then
		paste.status:SetText(status.text)
		paste.status:SetTextColor(Theme.rgb(Theme.HEX[status.hexKey or "success"]))
		y = stack(paste.status, paste.right, y + W.statusGap, textHeight(paste.status))
	end
	return y
end

--- The right column (spec 3.3): Up next, From the site, To the site. The y it used.
local function drawRight(view, verdict, select)
	local W = BuildsWindow.LAYOUT
	local paste = BuildDialog.paste
	local data = dataOf()
	local build = Follow.build
	local rows = build ~= nil and BuildsWindow.upNext(build, liveRanks(data), liveLevel(), data) or nil
	local name = build ~= nil and Follow.buildName(data, build) or nil
	local y = BuildsWindow.drawUpNext(paste.upNext, paste.right, rightWidth(), 0, rows, name, fitLine, textHeight)
	y = BuildsWindow.drawSection(paste.site, paste.right, rightWidth(), y + W.sectionGap, nil, fitLine, textHeight)
	y = drawSiteRows(verdict, y)
	y = drawPasteBox(view, verdict, y)
	y = drawUnderBox(view, verdict, y)
	return BuildDialog.drawCopy(select, y + W.sectionGap)
end

--- The footer's hints for what Escape and Enter do now (BuildsWindow.footer), Enter's only while the
--- paste box has the keyboard.
local function drawFooter(view, verdict)
	local item = verdict == nil and BuildDialog.cursorItem() or nil
	local cursor = item ~= nil and (item.own and "own" or "site") or nil
	local escape, enter = BuildsWindow.footer({
		previewing = view.preview == true, useBlocked = BuildDialog.useBlocked ~= nil, cursor = cursor,
	})
	-- Enter in the copy box selects the code again: nothing to hint.
	Hub.setHints(escape, BuildDialog.focus ~= "copy" and enter or nil)
end

--- Redraw the Builds tab from the state: the header, both columns, the footer, and the window's
--- height for the taller column. `select` true selects the copy box's code (Tab, /fs copy); otherwise
--- its text is only refreshed. The verdict.
function BuildDialog.draw(select)
	local paste = BuildDialog.paste
	if paste == nil then
		return nil
	end
	local verdict = BuildDialog.verdict
	local waiting = #(BuildDialog.rows or {}) > 0
	local resting = waiting and verdict == nil and BuildDialog.touched ~= true
	local view = BuildDialog.pasteView(verdict, BuildDialog.editing == true, BuildDialog.focused == true, resting)
	drawHeader()
	local left = drawLeft(view, verdict)
	local right = drawRight(view, verdict, select)
	Hub.fitHeight(math.max(left, right))
	drawFooter(view, verdict)
	return verdict
end

--- Back to the build list: no verdict, an empty paste box with the keyboard.
function BuildDialog.clearPaste()
	BuildDialog.verdict, BuildDialog.editing, BuildDialog.touched = nil, false, false
	if BuildDialog.paste ~= nil then
		BuildDialog.paste.box:SetText("")
		BuildDialog.draw()
		BuildDialog.focusBox("paste")
	end
end

--- Select all of the paste box (a second paste replaces the code rather than appending).
function BuildDialog.selectPaste()
	local box = BuildDialog.paste.box
	box:HighlightText()
	box:SetFocus()
end

-- To the site -----------------------------------------------------------------------------

local function outboxName()
	local ok, name = pcall(Export.outboxCharacter)
	return ok and name or nil
end

--- Scroll the export by `lines` whole lines inside its field, and move the thumb.
function BuildDialog.scrollCopy(lines)
	local copy = BuildDialog.copyBox
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

--- The copy model for now: the character's export (guarded: a client answer the export cannot encode
--- must never leave the section half drawn), or /fs link's page.
local function currentCopyModel()
	local link = BuildDialog.link
	if link ~= nil then
		local model = BuildDialog.copyModel({ link = link, copiedAt = BuildDialog.copiedAt, clock = liveClock() })
		BuildDialog.exportCode = model.text
		return model
	end
	local ran, code, reason = pcall(Export.string, dataOf())
	if not ran then
		code, reason = nil, nil
	end
	BuildDialog.exportCode = code
	local simInUse, simGeneratedAt = simWeightsInUse()
	return BuildDialog.copyModel({
		code = code,
		reason = reason,
		copiedAt = BuildDialog.copiedAt,
		ack = currentAck(),
		savedAt = BuildDialog.savedAtFor(ForeverSixtyDB, outboxName()),
		clock = liveClock(),
		simInUse = simInUse,
		simGeneratedAt = simGeneratedAt,
		nextStep = chipLine(),
		simFailed = simFailed(),
	})
end

--- The status line under the copy box (spec 3.3 item 3): the model's lead after a copy or for a link;
--- with the code there and not selected, "Click the box, then press Ctrl+C." muted. Pure.
function BuildDialog.copyStatus(model, selected)
	local lead = model.lead
	if model.text ~= "" and lead.hexKey == "body" and not selected and model.summary ~= nil then
		return { text = L.windowClickToSelect, hexKey = "muted" }
	end
	return lead
end

--- Draw To the site at `y` on the right column: the heading and its explanation, the copy box in
--- BuildsWindow.COPY_LINES whole lines (the rest scrolls), the status line, then the export's facts
--- (the summary, the site's ack, the sim's lines, the companion's lines when no ack was ever seen).
--- `select` true selects the code and gives the box the keyboard. The y under it.
function BuildDialog.drawCopy(select, y)
	local copy = BuildDialog.copyBox
	local S, W = Theme.SIZES, BuildsWindow.LAYOUT
	local content, width = copy.content, rightWidth()
	local model = currentCopyModel()
	y = BuildsWindow.drawSection(copy.section, content, width, y or 0, nil, fitLine, textHeight)
	local first, advance = lineMetrics()
	if model.text ~= "" then
		copy.display:SetText(copy.wrapsRuns and model.text or BuildDialog.breakable(model.text))
		copy.display:SetTextColor(Theme.rgb(Theme.HEX[model.textHex]))
		copy.lines = BuildDialog.boxLines(textHeight(copy.display), first, advance,
			first + (BuildsWindow.COPY_LINES - 1) * advance)
		local fieldHeight = sizeField(copy.field, copy.clip, copy.box, copy.lines.height)
		copy.box:SetHeight(first + (copy.lines.total - 1) * advance)
		y = stack(copy.field, content, y, fieldHeight) + W.statusGap
		if select == true and model.selected then
			BuildDialog.copyOffset = 0
			Widgets.selectText(copy.box, model.text)
			BuildDialog.selected = true
		elseif copy.box.foreverSixtyValue ~= model.text then
			BuildDialog.copyOffset = 0
			Widgets.setText(copy.box, model.text)
			BuildDialog.selected = false
		end
		BuildDialog.scrollCopy(0)
	else
		copy.field:Hide()
		Widgets.setText(copy.box, "")
		BuildDialog.selected = false
	end
	paintSelection()
	BuildDialog.copyModelNow = model
	-- The status line takes the room of its taller state (selected or not), so a click never moves
	-- the facts under it.
	local tallest = 0
	for _, selected in ipairs({ true, false }) do
		copy.lead:SetText(BuildDialog.copyStatus(model, selected).text)
		tallest = math.max(tallest, textHeight(copy.lead))
	end
	BuildDialog.paintCopyStatus()
	y = stack(copy.lead, content, y, tallest)
	local facts = {
		{ copy.summary, model.summary }, { copy.seen, model.seen }, { copy.sim, model.sim },
		{ copy.simNext, model.simNext },
	}
	for index = 1, #copy.companion do
		facts[#facts + 1] = { copy.companion[index], model.companion ~= nil and model.companion[index] or nil }
	end
	for _, fact in ipairs(facts) do
		if fact[2] ~= nil then
			fact[1]:SetText(fact[2])
			y = stack(fact[1], content, y + S.gap / 2, textHeight(fact[1]))
		else
			fact[1]:Hide()
		end
	end
	return y
end

--- The copy box's status line for the selection now (BuildDialog.copyStatus over the last model).
function BuildDialog.paintCopyStatus()
	local copy, model = BuildDialog.copyBox, BuildDialog.copyModelNow
	if copy == nil or model == nil then
		return nil
	end
	local lead = BuildDialog.copyStatus(model, BuildDialog.selected == true)
	copy.lead:SetText(lead.text)
	copy.lead:SetTextColor(Theme.rgb(Theme.HEX[lead.hexKey]))
	return lead
end

--- Select the export again for Ctrl+C (a click in the box, Enter, the deferred refocus). It
--- claims nothing itself.
function BuildDialog.copy()
	local copy = BuildDialog.copyBox
	if BuildDialog.exportCode == nil or copy == nil then
		return nil
	end
	Widgets.selectText(copy.box, BuildDialog.exportCode)
	BuildDialog.selected = true
	paintSelection()
	BuildDialog.paintCopyStatus()
	return BuildDialog.exportCode
end

-- Opening and closing --------------------------------------------------------------------

--- Hand the keyboard to a box: "copy" selects the code in the copy box (for Ctrl+C), "paste" gives the
--- paste box the keys (Ctrl+V, the arrows, Enter). The box now holding the keyboard.
function BuildDialog.focusBox(which)
	if BuildDialog.paste == nil then
		return nil
	end
	BuildDialog.focus = which == "copy" and "copy" or "paste"
	if BuildDialog.focus == "copy" and BuildDialog.copy() ~= nil then
		BuildDialog.paste.box:ClearFocus()
	else
		BuildDialog.focus = "paste"
		BuildDialog.copyBox.box:ClearFocus()
		BuildDialog.paste.box:SetFocus()
	end
	-- The footer says what Enter does in the box that has the keyboard now.
	BuildDialog.draw(false)
	return BuildDialog.focus
end

--- The Builds tab coming on screen (Hub.select): the state reset (nothing previewed, the box empty),
--- the site rows and the list read, the cursor on the first site row when any wait, the tab drawn,
--- then the keyboard to the box `options.focus` names ("copy": the copy box selected; else the paste
--- box). One frame later the box takes the keyboard again (the chat box that ran /fs lets go of it
--- after the command) and the render check runs.
function BuildDialog.onShow(options)
	BuildDialog.link = options.link
	BuildDialog.verdict, BuildDialog.editing = nil, false
	BuildDialog.touched, BuildDialog.ownFailed, BuildDialog.status = false, nil, nil
	BuildDialog.rows = currentRows()
	readOwn()
	BuildDialog.cursor = BuildDialog.startCursor(#BuildDialog.own, #BuildDialog.rows)
	BuildDialog.paste.box:SetText("")
	local focus = (options.focus == "copy" or options.link ~= nil) and "copy" or "paste"
	BuildDialog.draw(false)
	BuildDialog.focusBox(focus)
	local frame = BuildDialog.frame
	local deferred = Theme.after(0, function()
		if frame:IsShown() and Hub.active == "builds" then
			BuildDialog.focusBox(focus)
			BuildDialog.selfCheck()
		end
	end)
	if not deferred then
		BuildDialog.selfCheck()
	end
end

--- Open the window on the Builds tab with the keyboard in `focus`'s box ("paste" by default, or
--- "copy"). `link` is BuildDialog.openLink's. False when the window turned itself off earlier (the
--- fallback ran instead).
function BuildDialog.open(focus, link)
	if Hub.disabled and link ~= nil then
		-- The caller prints the address instead; the command list is not what /fs link asked for.
		return false
	end
	return Hub.open("builds", { focus = focus, link = link }) == "builds"
end

--- /fs link (L4.4): the copy box holding a character's site page (`url`) instead of the export,
--- selected for Ctrl+C, under "<name>'s page on foreversixty.gg. Press Ctrl+C, then paste it in your
--- browser." Closing the window goes back to the export. False when the window cannot show itself
--- (the caller prints the address instead).
function BuildDialog.openLink(url, name)
	return BuildDialog.open("copy", { url = url, name = name })
end

--- The site rows and the list read again, and the tab redrawn while it shows (L3.7 item 44: an
--- officer's push just arrived).
function BuildDialog.refreshRows()
	BuildDialog.rows = currentRows()
	readOwn()
	BuildDialog.cursor = math.max(0, math.min(BuildDialog.cursor or 0, #currentItems()))
	if Hub.isShown() and Hub.active == "builds" then
		BuildDialog.draw(false)
	end
end

--- Hide the window and give the keyboard back (the hub's OnHide runs onHide too).
function BuildDialog.close()
	Hub.close()
end

--- Open on the Builds tab, or close when it is the tab showing.
function BuildDialog.toggle()
	if Hub.isShown() and Hub.active == "builds" then
		Hub.close()
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
		Hub.disabled = true
		Hub.close()
		Theme.note(L.diagDialogOff)
		if Hub.fallback ~= nil then
			Hub.fallback()
		end
		return false
	end
	local box = BuildDialog.paste.box
	if BuildDialog.focus == "paste" and type(box.HasFocus) == "function" and box:HasFocus() ~= true then
		BuildDialog.noteNoFocus()
		return false
	end
	return true
end

-- The hub's first tab (spec section 3): the window opens on it until other tabs have content.
Hub.register({
	id = "builds",
	label = L.hubTabBuilds,
	order = 1,
	build = BuildDialog.buildInto,
	show = BuildDialog.onShow,
	hide = BuildDialog.onHide,
	escape = BuildDialog.escape,
})

ns.BuildDialog = BuildDialog
return BuildDialog
