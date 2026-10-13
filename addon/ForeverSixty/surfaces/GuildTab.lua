-- addon/ForeverSixty/surfaces/GuildTab.lua
-- The guild tab (spec section 4.7, task L5.1): the guild's roster with level, spec, gear score,
-- Forever Score, attendance and gear gap, sortable; for an officer, the site's approval queue
-- and "Push the raid build to guild".
--
-- Ruling 1 (taint): the 1.60 client has no GuildFrame; its guild UI is CommunitiesFrame, and
-- it taints a frame of its own on any HookScript, CreateTexture, tab-template call or field
-- write (L4.1, L4.3). So, as the character panel tab (L4.3), the "tab" is an addon-owned
-- button and an addon-owned flyout, both parented to CommunitiesFrame (looked up each time;
-- the client loads it on demand) and anchored to it: the button under its bottom right, the
-- flyout TOPLEFT at its TOPRIGHT, right of any of its side tabs that shows with a rect. The
-- client frame is only ever read (Widgets' shared reads). Scripts sit only on the addon's own
-- frames; the flyout hides with the guild frame because it is its child. Nothing polls.
--
-- Ruling 2: the roster model is pure (GuildTab.rows: the client's roster joined with the
-- nightly record and the live snapshots, sorted and filtered), and so are the cells, the
-- officer region and the heights. The flyout may scroll (an addon-owned ScrollFrame of 20
-- pooled rows, three a wheel notch) and is never taller than the guild frame.
--
-- Ruling 3: each open asks the guild once (snapshot? on GUILD, through Comms, at most once a
-- minute, never while "Share my build with my guild" is off), so members with the addon
-- answer and their rows go live. Ruling 4: the push opens the addon's own confirm frame.
-- Ruling 5: nothing for L5.3 here yet.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Codec = ns.Codec or require("Codec")
local Export = ns.Export or require("Export")
local Follow = ns.Follow or require("Follow")
local Score = ns.Score or require("Score")
local Comms = ns.Comms or require("Comms")
local Compat = ns.Compat or require("Compat")
local Talents = ns.Talents or require("Talents")
local Gear = ns.Gear or require("Gear")
local ItemTooltip = ns.ItemTooltip or require("ItemTooltip")

local GuildTab = {}

local S = Theme.SIZES

--- The frames' names, for the /run probes (README rows 58-62).
GuildTab.BUTTON_NAME = "ForeverSixtyGuildTab"
GuildTab.FLYOUT_NAME = "ForeverSixtyGuildFlyout"
GuildTab.CONFIRM_NAME = "ForeverSixtyGuildConfirm"
--- The client's guild frame (1.60: no GuildFrame, GuildFrameTab1 or GuildRosterFrame).
GuildTab.FRAME = "CommunitiesFrame"
--- Its side tabs (chat, roster, perks, info), read as its fields and taken only when one shows
--- with a rect: the flyout opens right of the rightmost. Read, never written.
GuildTab.SIDE_TABS = { "ChatTab", "RosterTab", "GuildBenefitsTab", "GuildInfoTab" }
--- The most rows shown at once, and how many rows a wheel notch moves.
GuildTab.VISIBLE_ROWS = 20
GuildTab.WHEEL_STEP = 3
--- "Online only" starts on above this many members.
GuildTab.ONLINE_ONLY_ABOVE = 40
--- The least time, in seconds, between two roster requests and between two snapshot? asks.
GuildTab.ASK_EVERY = 60
--- The text a line is measured on: a capital, an ascender and descenders.
GuildTab.LINE_SAMPLE = "Agjy"
--- The columns, left to right: the sort key, its width, whether it sorts as a number (a
--- number's first click sorts high to low; a name's A to Z) and whether it is right-aligned
--- (fix round 1, item 3: the numbers line up under their header's label, which ends
--- Theme.SIZES.guildSortReserve short of the column's edge, the sort arrow's room). The widths
--- and the gaps fit the flyout's inner width less half the scroll track, which straddles its
--- right edge; Spec keeps room for "Marksmanship".
GuildTab.COLUMNS = {
	{ key = "name", width = 90 },
	{ key = "level", width = 40, numeric = true, right = true },
	{ key = "spec", width = 82 },
	{ key = "gear", width = 44, numeric = true, right = true },
	{ key = "score", width = 52, numeric = true, right = true },
	{ key = "attendance", width = 68, numeric = true, right = true },
	{ key = "gap", width = 52, numeric = true, right = true },
}

--- The generated data table, set by Options.register().
GuildTab.data = nil
--- The frames, once made, and the pooled roster rows.
GuildTab.button = nil
GuildTab.flyout = nil
GuildTab.rowFrames = {}
GuildTab.headers = {}
--- The sort a header click chose (nil: online first, then Score high to low), the switch
--- (nil: GuildTab.defaultOnlineOnly), and the first row shown.
GuildTab.sort = nil
GuildTab.onlineOnly = nil
GuildTab.offset = 0
--- How many times the flyout was drawn (a spec reads it), and the model last drawn.
GuildTab.draws = 0
GuildTab.last = nil
--- The game clock at the last roster request and the last snapshot? sent.
GuildTab.lastRoster = nil
GuildTab.lastAsk = nil
--- The last push: { text, failed }, drawn under the push button.
GuildTab.pushed = nil
--- The measure every cell is fitted with, and a text line's measured height; nil before the
--- flyout is made.
GuildTab.measure = nil
GuildTab.line = nil

-- The pure model ---------------------------------------------------------------------------

local NUMERIC = {}
for _, column in ipairs(GuildTab.COLUMNS) do
	NUMERIC[column.key] = column.numeric == true
end

--- "Smoke Gouda-ClassicBetaPvP" -> "Smoke Gouda"; a name with no realm as it is.
function GuildTab.shortName(name)
	return type(name) == "string" and (name:match("^([^%-]+)%-") or name) or ""
end

--- The live snapshot for a roster name (Comms.seen, keyed by sender "Name-Realm"): the exact
--- key, else the same bare name where either side carries no realm.
function GuildTab.seenFor(seen, name)
	if type(seen) ~= "table" or type(name) ~= "string" then
		return nil
	end
	if type(seen[name]) == "table" then
		return seen[name]
	end
	local short, hasRealm = GuildTab.shortName(name), name:find("-", 1, true) ~= nil
	for sender, entry in pairs(seen) do
		if type(entry) == "table" and type(sender) == "string" and GuildTab.shortName(sender) == short
			and (not hasRealm or not sender:find("-", 1, true)) then
			return entry
		end
	end
	return nil
end

--- The nightly record for a name: `record` is Score.forCharacter (a function) or a characters
--- table keyed "Name-Realm" (contract 5). A character with no logs has no record, never a zero.
local function recordFor(record, name)
	local row
	if type(record) == "function" then
		row = record(name)
	elseif type(record) == "table" then
		row = record[name]
	end
	if type(row) ~= "table" or type(row.score) ~= "number" then
		return nil
	end
	if row.score == 0 and (row.fights or 0) == 0 then
		return nil
	end
	return row
end

--- A realm with its spaces and hyphens gone, the form senders and the roster carry.
local function compact(realm)
	return (realm:gsub("[%s%-]", ""))
end

--- Whether a roster name is the player (`own`: { name, realm }): the same bare name, and the
--- same realm where both sides carry one.
local function isOwn(name, own)
	if type(own) ~= "table" or type(own.name) ~= "string"
		or GuildTab.shortName(name) ~= GuildTab.shortName(own.name) then
		return false
	end
	local realm = name:match("^[^%-]+%-(.+)$")
	return realm == nil or type(own.realm) ~= "string" or compact(realm) == compact(own.realm)
end

--- Whether a Comms.seen entry is still news at `now` (L5.1 review item 1): heard within
--- Comms.SEEN_FOR (Comms prunes only on new traffic). Any entry when `now` is not given.
local function fresh(snapshot, now)
	return type(now) ~= "number" or type(snapshot.time) ~= "number" or now - snapshot.time <= Comms.SEEN_FOR
end

--- One row: the roster's name, class, rank, level and online flag; the live fields (`live`)
--- from this client for the player (`context.own`, fix round 1 item 2) and from a fresh
--- snapshot for a guildmate online (review item 1): the spec, the gap in its unit, the build;
--- the record's score, tier, gear score, attendance (a whole percent) and, with no live spec,
--- its spec. `noLogs`: the record is readable (`context.ratings`) and has nothing for the
--- member. The roster's level wins over a snapshot's (review item 2).
local function rowOf(entry, record, seen, context)
	local found = recordFor(record, entry.name)
	local row = {
		name = entry.name, short = GuildTab.shortName(entry.name), classToken = entry.classToken,
		online = entry.online == true, rankIndex = entry.rankIndex, level = entry.level, live = false,
		record = found,
	}
	local live
	if isOwn(entry.name, context.own) then
		live, row.own = context.own, true
	elseif row.online then
		local snapshot = GuildTab.seenFor(seen, entry.name)
		if snapshot ~= nil and fresh(snapshot, context.now) then
			live = snapshot
		end
	end
	if live ~= nil then
		row.live = true
		if type(row.level) ~= "number" then
			row.level = tonumber(live.level)
		end
		row.gap = tonumber(live.gearGap)
		if type(live.spec) == "string" and live.spec ~= "" and live.spec ~= "-" then
			row.spec = live.spec
		end
		if row.gap ~= nil then
			row.unit = live.unit or Comms.gapUnit(context.data, row.spec, row.level)
		end
		if type(live.buildName) == "string" and live.buildName ~= "" then
			row.buildName = live.buildName
		end
	end
	if found ~= nil then
		row.score, row.tier = found.score, found.tier
		row.gearScore = tonumber(found.gear_score)
		if type(found.attendance) == "number" then
			row.attendance = math.floor(found.attendance * 100 + 0.5)
		end
		if row.spec == nil and type(found.spec) == "string" and found.spec ~= "" then
			row.spec = found.spec
		end
	end
	row.noLogs = found == nil and context.ratings ~= false
	return row
end

--- What a row sorts by in a column; nil (no value) sorts last both ways.
local function valueOf(row, key)
	if key == "name" then
		return row.short:lower()
	elseif key == "spec" then
		return row.spec ~= nil and Follow.specName(row.spec):lower() or nil
	elseif key == "gear" then
		return row.gearScore
	end
	return row[key]
end

--- The order for `sort` ({ key, descending }; nil: online first, then Score high to low).
--- Equal values go by name, A to Z.
local function comparator(sort)
	local key = sort ~= nil and sort.key or "score"
	local descending = sort == nil or sort.descending == true
	local onlineFirst = sort == nil
	return function(left, right)
		if onlineFirst and left.online ~= right.online then
			return left.online
		end
		local a, b = valueOf(left, key), valueOf(right, key)
		if a ~= b then
			if a == nil then
				return false
			elseif b == nil then
				return true
			end
			if descending then
				return a > b
			end
			return a < b
		end
		local leftName, rightName = left.short:lower(), right.short:lower()
		if leftName ~= rightName then
			return leftName < rightName
		end
		return left.name < right.name
	end
end

--- The roster (ruling 2): `roster` (Theme.guildRoster's { name, rankIndex, level, classToken,
--- online }) joined with `record` (Score.forCharacter, or the record's characters table) and
--- `seen` (Comms.seen), sorted by `sort` (GuildTab.nextSort's; nil the default) and filtered
--- by `filter` ({ onlineOnly }). `context` (optional): { now (Theme.now, for a snapshot's
--- age), own (GuildTab.ownLive: the player's live fields), ratings (false while the record
--- cannot be read: no "no logs" then), data (the gap's unit) }. Answers the rows and { total,
--- shown, live, online } (live: shown rows with live fields, every one of them online;
--- online: the roster's members online).
function GuildTab.rows(roster, record, seen, sort, filter, context)
	local shown, total, live, online = {}, 0, 0, 0
	local onlineOnly = type(filter) == "table" and filter.onlineOnly == true
	context = type(context) == "table" and context or {}
	for _, entry in ipairs(type(roster) == "table" and roster or {}) do
		if type(entry) == "table" and type(entry.name) == "string" and entry.name ~= "" then
			total = total + 1
			if entry.online == true then
				online = online + 1
			end
			if not onlineOnly or entry.online == true then
				local row = rowOf(entry, record, seen, context)
				shown[#shown + 1] = row
				if row.live then
					live = live + 1
				end
			end
		end
	end
	table.sort(shown, comparator(sort))
	return shown, { total = total, shown = #shown, live = live, online = online }
end

--- The sort after a click on `key`'s header: the same column turns round; another starts high
--- to low for a number, A to Z for a name. The default is Score high to low.
function GuildTab.nextSort(sort, key)
	local current = sort or { key = "score", descending = true }
	if current.key == key then
		return { key = key, descending = not current.descending }
	end
	return { key = key, descending = NUMERIC[key] == true }
end

--- Whether "Online only" starts on: above GuildTab.ONLINE_ONLY_ABOVE members.
function GuildTab.defaultOnlineOnly(total)
	return (tonumber(total) or 0) > GuildTab.ONLINE_ONLY_ABOVE
end

local function whole(value)
	return string.format("%d", math.floor(value + 0.5))
end

local function plain(text)
	return { text = text, hex = Theme.HEX.body }
end

local function muted(text)
	return { text = text, hex = Theme.HEX.muted }
end

--- A cell nobody has a value for (fix round 1, item 3): blank, never a dash.
local function blank()
	return plain("")
end

--- A row's cells, each { text, hex }: the name in its class colour lifted for the panel
--- (Theme.panelClassHex), the score in its tier's panel colour (Theme.panelTierHex; "no logs"
--- muted only where a readable record has nothing for the member), attendance as "80%", the
--- live gap in its sender's unit ("14 DPS", "9 pts"); blank where there is no value. An
--- offline row is dimmed once, as a whole, where it is drawn.
function GuildTab.cells(row)
	local gapText = row.unit == "dps" and L.guildTabGapDps or L.guildTabGapPts
	local cells = {
		name = { text = row.short, hex = row.classToken ~= nil and Theme.panelClassHex(row.classToken) or Theme.HEX.body },
		level = type(row.level) == "number" and plain(whole(row.level)) or blank(),
		spec = row.spec ~= nil and plain(Follow.specName(row.spec)) or blank(),
		gear = type(row.gearScore) == "number" and plain(whole(row.gearScore)) or blank(),
		attendance = type(row.attendance) == "number" and plain(string.format(L.guildTabPercent, row.attendance))
			or blank(),
		gap = type(row.gap) == "number" and plain(string.format(gapText, math.floor(row.gap + 0.5))) or blank(),
	}
	if type(row.score) == "number" then
		local r, g, b = Score.tierColor(row.tier)
		local hex = Theme.panelTierHex(row.tier, r ~= nil and Theme.hexOf(r, g, b) or nil)
		cells.score = { text = whole(row.score), hex = hex or Theme.HEX.body }
	elseif row.noLogs then
		cells.score = muted(L.guildTabNoLogs)
	else
		cells.score = blank()
	end
	return cells
end

--- A row's hover: the full name in its class colour, the score line (or "no logs yet" where
--- the record is readable), the live line and build where there are live fields, "Offline".
function GuildTab.rowHint(row)
	local cells = GuildTab.cells(row)
	local lines = { "|cff" .. cells.name.hex .. row.name .. "|r" }
	local scoreLines = Score.lines(row.record)
	if scoreLines[1] ~= nil then
		lines[#lines + 1] = "|cff" .. cells.score.hex .. scoreLines[1] .. "|r"
	elseif row.noLogs then
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", L.scoreNoLogs)
	end
	if row.live then
		local spec = row.spec ~= nil and Follow.specName(row.spec) or ""
		local live = string.format(L.scoreLive, row.level or 0, spec, math.floor((row.gap or 0) + 0.5))
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", (live:gsub("  +", " ")))
		if row.buildName ~= nil then
			lines[#lines + 1] = Theme.inlineColor("tooltipText", string.format(L.guildTabBuild, row.buildName))
		end
	end
	if not row.online then
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", L.guildTabOffline)
	end
	return lines
end

--- The count line (fix round 1, item 6): "3 of 97 shown · 0 of 3 online run Forever Sixty",
--- then " · sharing off" while the share switch is off (`sharing` false), then " · 1–20 of 97"
--- while the list holds more rows than show (`first` to `last` of `total`). Pure.
function GuildTab.countsText(counts, sharing, first, last, total)
	local text = string.format(L.guildTabCounts, counts.shown or 0, counts.total or 0, counts.live or 0,
		counts.online or 0)
	if sharing == false then
		text = text .. L.guildTabSharingOff
	end
	if type(first) == "number" and type(last) == "number" and type(total) == "number"
		and last - first + 1 < total then
		text = text .. string.format(L.guildTabRange, first, last, total)
	end
	return text
end

--- The one line said for the whole list while the ratings cannot be read or are old (fix
--- round 1, item 1): "Player ratings: " and the Settings foot's state for them
--- (Score.ratingsState); nil while they are current. `status` is Score.status()'s. Pure.
function GuildTab.ratingsNotice(status)
	local state = Score.ratingsState(type(status) == "table" and status or false)
	if state == "Current" then
		return nil
	end
	local sentence = L["settingsRatings" .. state]
	return string.format(L.guildTabRatings, sentence:match("^.-:%s(.*)$") or sentence)
end

--- The button's hover: "Forever Sixty" in gold, then what it holds and what a click does.
function GuildTab.buttonHint()
	local lines = { Theme.inlineColor("gold", L.addonName) }
	for _, text in ipairs(L.guildTabHint) do
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", text)
	end
	return lines
end

-- The officer region (ruling 4) ------------------------------------------------------------

local CLAIMS = { unclaimed = L.guildClaimUnclaimed, pending = L.guildClaimPending, claimed = L.guildClaimClaimed,
	contested = L.guildClaimContested }

--- Whether a guild rank index may push a build: Comms.canPush, the rule the receiver checks too
--- (the guild master, rank 0, and rank 1).
function GuildTab.isOfficer(rankIndex)
	return Comms.canPush(rankIndex)
end

--- The officer region, nil below an officer: from the inbox `guild` message (`message`,
--- optional) its approvals ("2 approvals waiting on the site", "No approvals waiting") and
--- claim state; the build a push would send (`buildName`, Follow.buildName) and whether one
--- can go (`canPush`), with a note when there is no build.
function GuildTab.officerRegion(message, rankIndex, build, data)
	if not GuildTab.isOfficer(rankIndex) then
		return nil
	end
	local region = { canPush = build ~= nil }
	if type(message) == "table" then
		local pending = tonumber(message.pending_approvals)
		if pending ~= nil then
			if pending == 0 then
				region.approvals = L.guildTabNoApprovals
			elseif pending == 1 then
				region.approvals = L.guildTabApproval
			else
				region.approvals = string.format(L.guildTabApprovals, pending)
			end
		end
		local claim = CLAIMS[message.claim_state]
		if claim ~= nil then
			region.claim = string.format(L.guildTabClaim, claim)
		end
	end
	if build ~= nil then
		region.buildName = Follow.buildName(data or {}, build)
	else
		region.note = L.guildTabPushNoBuild
	end
	return region
end

--- One UTF-8 character off the end of a name.
local function shorten(name)
	return (name:gsub("[%z\1-\127\192-\255][\128-\191]*$", ""))
end

--- The push's fields, { code, name }: the active build as an FSB1 code (contract 2; its gear
--- left out when the whole would not fit one addon message) and its name (Follow.buildName,
--- shortened to fit, never split). nil with no build or nothing that fits.
function GuildTab.pushFields(data, build)
	if type(build) ~= "table" or type(build.classSlug) ~= "string" then
		return nil
	end
	local dataBuild = build.dataBuild or (type(data) == "table" and data.build or nil)
	if dataBuild == nil then
		return nil
	end
	local parts = { dataBuild = dataBuild, classSlug = build.classSlug, order = build.order or {}, gear = build.gear }
	local fields = { Codec.encodeFSB1(parts), Follow.buildName(data or {}, build) }
	if Comms.encode("push", fields) == nil then
		parts.gear = nil
		fields[1] = Codec.encodeFSB1(parts)
	end
	while Comms.encode("push", fields) == nil and fields[2] ~= "" do
		fields[2] = shorten(fields[2])
	end
	if Comms.encode("push", fields) == nil then
		return nil
	end
	return fields
end

-- The heights, pure ------------------------------------------------------------------------

--- A text line's height: the measured one, else `panelLine`.
function GuildTab.lineHeight()
	return GuildTab.line or S.panelLine
end

--- One roster row: `guildRow`, never under a line.
function GuildTab.rowHeight()
	return math.max(S.guildRow, GuildTab.lineHeight())
end

--- The officer region's text lines ({ key, text, color }), in order.
local function officerLines(officer)
	local lines = {}
	for _, line in ipairs({ { key = "approvals", color = "body" }, { key = "claim", color = "muted" },
		{ key = "note", color = "muted" } }) do
		if officer[line.key] ~= nil then
			lines[#lines + 1] = { key = line.key, text = officer[line.key], color = line.color }
		end
	end
	return lines
end

--- The officer region's height as drawn: its title, its lines, the push button, the pushed
--- line when there is one; 0 for none.
function GuildTab.officerHeight(officer)
	if officer == nil then
		return 0
	end
	local height = S.panelTitle + #officerLines(officer) * GuildTab.lineHeight() + S.panelRowGap + S.panelButtonHeight
	if officer.pushed ~= nil then
		height = height + S.panelRowGap + GuildTab.lineHeight()
	end
	return height
end

--- Where the ratings notice sits, from the flyout's top: the padding, the header, the switch
--- row (the count line), a row gap.
function GuildTab.noticeTop()
	return S.panelPadding + S.panelHeader + S.panelRowGap + S.guildToolbar + S.panelRowGap
end

--- Where the column headers sit: under the switch row, and under the ratings notice and a row
--- gap when one shows (`notice` true; fix round 1, item 1).
function GuildTab.headerTop(notice)
	return GuildTab.noticeTop() + (notice and GuildTab.lineHeight() + S.panelRowGap or 0)
end

--- Where the roster rows start: under the column headers.
function GuildTab.rowsTop(notice)
	return GuildTab.headerTop(notice) + S.guildHeaderRow
end

--- The flyout's height with `rows` roster rows, `officer` (its region gap included) and the
--- ratings notice (`notice` true).
function GuildTab.flyoutHeight(rows, officer, notice)
	local height = GuildTab.rowsTop(notice) + rows * GuildTab.rowHeight()
	if officer ~= nil then
		height = height + S.panelRegionGap + GuildTab.officerHeight(officer)
	end
	return height + S.panelPadding
end

--- How many roster rows show at once under `limit` (the guild frame's height; the flyout is
--- never taller): GuildTab.VISIBLE_ROWS at most, one at least; all of them where the client
--- cannot say.
function GuildTab.fitRows(limit, officer, notice)
	if type(limit) ~= "number" or limit <= 0 then
		return GuildTab.VISIBLE_ROWS
	end
	local room = math.floor((limit - GuildTab.flyoutHeight(0, officer, notice)) / GuildTab.rowHeight())
	return math.max(1, math.min(GuildTab.VISIBLE_ROWS, room))
end

-- The player's own row and the model as drawn (fix round 1, item 2) -------------------------

--- The tree an active build is for, as a spec slug: the build's own `spec` (a data default
--- carries one), else the class tree its plan spends the most points in ("hunter-marksmanship"
--- for a Marksmanship plan, whatever is spent so far); nil when neither says.
function GuildTab.buildSpec(data, build)
	if type(build) ~= "table" then
		return nil
	end
	if type(build.spec) == "string" and build.spec ~= "" then
		return build.spec
	end
	local class = type(data) == "table" and type(data.classes) == "table" and data.classes[build.classSlug] or nil
	if type(class) ~= "table" or type(class.tabs) ~= "table" then
		return nil
	end
	local spent = {}
	for _, point in ipairs(build.order or {}) do
		if type(point) == "table" and type(point.tab) == "number" then
			spent[point.tab] = (spent[point.tab] or 0) + 1
		end
	end
	local best, bestTab = 0, nil
	for tab, entry in ipairs(class.tabs) do
		if (spent[tab] or 0) > best and type(entry) == "table" and type(entry.name) == "string" then
			best, bestTab = spent[tab], tab
		end
	end
	return bestTab ~= nil and build.classSlug .. "-" .. Export.slugify(class.tabs[bestTab].name) or nil
end

--- The player's own live fields, read on this client rather than heard (Comms ignores the
--- player's own messages): { name, realm, level, spec, gearGap, unit, buildName }. The spec is
--- the active build's (GuildTab.buildSpec), else the talents' (Gear.specOf, the chip's); the
--- gap and its unit are Comms.gearGap's, the one a hello sends (Progress, as the chip reads
--- it; nil without a build); the build name is the chip's (Follow.buildName). nil without a
--- name or the data.
function GuildTab.ownLive(data)
	local name = Compat.playerName()
	if name == nil or type(data) ~= "table" or type(data.classes) ~= "table" then
		return nil
	end
	local _, realm = Theme.playerNameRealm()
	local build = Follow.build
	local classSlug = Talents.playerClassSlug()
	local ranks = Talents.readRanks(data) or {}
	local talentSpec
	if classSlug ~= nil and type(data.classes[classSlug]) == "table" then
		local ok, spec = pcall(Gear.specOf, data, classSlug, ranks)
		talentSpec = ok and spec or nil
	end
	local gap, unit = Comms.gearGap(data, build, talentSpec, ranks)
	return {
		name = name, realm = realm, level = Compat.playerLevel(), spec = GuildTab.buildSpec(data, build) or talentSpec,
		gearGap = gap, unit = unit, buildName = build ~= nil and Follow.buildName(data, build) or nil,
	}
end

--- The roster as drawn now: the client's roster joined with the record, the guild's fresh
--- snapshots and the player's own live fields, sorted and filtered. Answers the rows, the
--- counts, whether Online only is on, and Score.status() (for the notice).
function GuildTab.model()
	local roster = Theme.guildRoster()
	local onlineOnly = GuildTab.onlineOnly
	if onlineOnly == nil then
		onlineOnly = GuildTab.defaultOnlineOnly(#roster)
	end
	local status = Score.status()
	local data = GuildTab.data or ns.Data
	local ok, own = pcall(GuildTab.ownLive, data)
	local rows, counts = GuildTab.rows(roster, Score.forCharacter, Comms.seen, GuildTab.sort,
		{ onlineOnly = onlineOnly },
		{ now = Theme.now(), own = ok and own or nil, ratings = status.available == true, data = data })
	return rows, counts, onlineOnly, status
end

--- Seeds one guildmate's live fields into Comms.seen, as if their snapshot! had just come, and
--- redraws an open flyout: a capture helper only (README row 60), with no effect unless called
--- (`/run ForeverSixtyGuildTab.seedLive("Name", { level = 30, spec = "warrior-arms",
--- gearGap = 14, buildName = "Arms raid" })`). True when seeded.
function GuildTab.seedLive(name, fields)
	if type(name) ~= "string" or name == "" then
		return false
	end
	fields = type(fields) == "table" and fields or {}
	Comms.seen[name] = {
		level = tonumber(fields.level), spec = fields.spec, gearGap = tonumber(fields.gearGap),
		buildName = fields.buildName, scoreDate = fields.scoreDate or "-", channel = "GUILD", time = Theme.now(),
	}
	if GuildTab.redraw ~= nil then
		GuildTab.redraw()
	end
	return true
end

-- The client reads ----------------------------------------------------------------------------

--- The client's guild frame, or nil (it loads on demand, so it is looked up each time).
local function pane()
	return Widgets.clientFrame(GuildTab.FRAME)
end

--- The rightmost of the guild frame's side tabs that shows with a rect, or nil. Each is read
--- as a field of the frame, never written.
local function sideColumn(frame)
	local found, foundRight
	for _, key in ipairs(GuildTab.SIDE_TABS) do
		local ok, tab = pcall(function()
			return frame[key]
		end)
		if ok and type(tab) == "table" and Widgets.placed(tab) then
			local right = Widgets.frameEdge(tab, "GetRight")
			if found == nil or (right ~= nil and (foundRight == nil or right > foundRight)) then
				found, foundRight = tab, right
			end
		end
	end
	return found
end

--- The flyout TOPLEFT at the guild frame's TOPRIGHT, moved right past its side tabs when one
--- shows with a rect (level with the frame's top) and clear of them by `guildSideClear` (fix
--- round 1, item 7: the tabs' art reaches past their rect); at that tab's own TOPRIGHT, the
--- same gap across, where the edges cannot be read.
local function placeFlyout(flyout, frame)
	local column = sideColumn(frame)
	if column ~= nil then
		local clear = S.panelOffsetX + S.guildSideClear
		local right, own = Widgets.frameEdge(column, "GetRight"), Widgets.frameEdge(frame, "GetRight")
		if right ~= nil and own ~= nil then
			flyout:ClearAllPoints()
			flyout:SetPoint("TOPLEFT", frame, "TOPRIGHT", math.max(0, right - own) + clear, S.panelOffsetY)
			return
		end
		Widgets.placeFlyout(flyout, frame, column, clear, S.panelOffsetY)
		return
	end
	Widgets.placeFlyout(flyout, frame, nil, S.panelOffsetX, S.panelOffsetY)
end

--- The newest inbox `guild` message for this character (the chip's own reader), or nil.
local function guildMessage()
	return Follow.messages(ForeverSixtyInbox, Export.characterKey(), "guild")[1]
end

local function guarded(action, ...)
	return Widgets.guarded(GuildTab, L.diagGuildTabFailed, action, ...)
end

-- The frames --------------------------------------------------------------------------------

--- The "Forever Sixty" button under the guild frame's bottom right, the panel tab's look
--- (Widgets.edgedButton, 112 x 24) and hover hint. Made the first time a refresh finds the
--- guild frame (never in combat: it waits for its end). nil without one. It carries the
--- capture helper GuildTab.seedLive for a /run (README row 60).
function GuildTab.ensureButton()
	local frame = pane()
	if frame == nil then
		return nil
	end
	local button = GuildTab.button
	if button == nil then
		if Theme.inCombat() then
			GuildTab.deferred = true
			return nil
		end
		button = Widgets.edgedButton(frame, L.guildTabLabel, function()
			GuildTab.toggle()
		end, GuildTab.BUTTON_NAME, S.panelTabWidth, S.panelTabHeight)
		-- The addon's own button, a child of the guild frame: it shows each time the frame does.
		button:SetScript("OnShow", function(self)
			GuildTab.onButtonShow(self)
		end)
		Widgets.withHint(button, GuildTab.buttonHint)
		button:ClearAllPoints()
		button:SetPoint("TOPRIGHT", frame, "BOTTOMRIGHT", 0, 0)
		button.seedLive = GuildTab.seedLive
		GuildTab.button = button
	end
	return button
end

--- The inner width: the flyout less its padding.
function GuildTab.innerWidth()
	return S.guildFlyoutWidth - S.panelPadding * 2
end

--- A column's x within a row.
function GuildTab.columnX(index)
	local x = 0
	for before = 1, index - 1 do
		x = x + GuildTab.COLUMNS[before].width + S.guildColumnGap
	end
	return x
end

--- The room a column's text has: a right-aligned column's ends where its header's label does,
--- `guildSortReserve` short of its edge (the arrow's room); a left-aligned one's runs to the
--- edge (its label's arrow sits inside it).
function GuildTab.textWidth(column)
	return column.right and column.width - S.guildSortReserve or column.width
end

local function fitted(text, room)
	return ItemTooltip.fitPlain(text, GuildTab.measure, room)
end

--- A font string one line high and `width` wide, never wrapping.
local function sized(region, width)
	region:SetSize(width, GuildTab.lineHeight())
	if type(region.SetWordWrap) == "function" then
		region:SetWordWrap(false)
	end
	return region
end

--- One pooled roster row: a Button the width of the rows, a hover ground, a cell per column
--- (a right-aligned column's text ends under its header's label).
local function rowFrame(index)
	local row = GuildTab.rowFrames[index]
	if row ~= nil then
		return row
	end
	local frame = CreateFrame("Button", nil, GuildTab.content)
	frame:SetSize(GuildTab.innerWidth(), GuildTab.rowHeight())
	frame:EnableMouse(true)
	local ground = Theme.texture(frame, "BACKGROUND", "hover")
	ground:SetAllPoints(frame)
	ground:Hide()
	row = { frame = frame, ground = ground, cells = {} }
	for column, spec in ipairs(GuildTab.COLUMNS) do
		local cell = sized(Widgets.label(frame, "", "body", "small"), GuildTab.textWidth(spec))
		cell:SetPoint("LEFT", frame, "LEFT", GuildTab.columnX(column), 0)
		if spec.right then
			cell:SetJustifyH("RIGHT")
		end
		row.cells[spec.key] = cell
	end
	frame:SetScript("OnEnter", function(self)
		ground:Show()
		if row.entry ~= nil then
			Theme.showLines(self, GuildTab.rowHint(row.entry), "ANCHOR_RIGHT")
		end
	end)
	frame:SetScript("OnLeave", function()
		ground:Hide()
		Theme.hideTooltip()
	end)
	GuildTab.rowFrames[index] = row
	return row
end

--- A header's label colour (Theme.CONTROL_STATES.header): the sorted column gold, another
--- muted, brightened under the cursor.
local function paintHeader(key)
	local header = GuildTab.headers[key]
	local states = Theme.CONTROL_STATES.header
	local sorted = GuildTab.sort ~= nil and GuildTab.sort.key or "score"
	local look = key == sorted and states.active or header.over and states.hoverLabel or states.label
	header.label:SetTextColor(Theme.rgb(Theme.HEX[look]))
	header.look = look
end

--- The column headers (fix round 1, items 3 and 5): an addon Button per column, its label
--- (right-aligned columns' ending `guildSortReserve` short of the edge), the client's sort
--- arrow (8 x 8) just after the label inside that room, an inset ground while held; the Gear
--- and Gap headers say on hover what they hold. Placed by layoutHeaders.
local function makeHeaders(flyout)
	local states = Theme.CONTROL_STATES.header
	for _, column in ipairs(GuildTab.COLUMNS) do
		local key = column.key
		local button = CreateFrame("Button", nil, flyout)
		button:SetSize(column.width, S.guildHeaderRow)
		local ground = Theme.texture(button, "BACKGROUND", states.pressed)
		ground:SetAllPoints(button)
		ground:Hide()
		local label = Widgets.label(button, L.guildTabColumns[key], states.label, "small")
		if column.right then
			label:SetJustifyH("RIGHT")
			label:SetPoint("RIGHT", button, "RIGHT", -S.guildSortReserve, 0)
		else
			label:SetPoint("LEFT", button, "LEFT", 0, 0)
		end
		local arrow = Theme.tinted(button, "OVERLAY", Theme.MEDIA.sortArrow, states.active)
		arrow:SetSize(S.guildSortArrow, S.guildSortArrow)
		arrow:SetPoint("LEFT", label, "RIGHT", 1, 0)
		arrow:Hide()
		local hint = L.guildTabHeaderHints[key]
		local header = { button = button, label = label, arrow = arrow, ground = ground, hint = hint }
		GuildTab.headers[key] = header
		button:SetScript("OnEnter", function(self)
			header.over = true
			paintHeader(key)
			if hint ~= nil then
				Theme.showLines(self, { Theme.inlineColor("tooltipText", hint) }, "ANCHOR_TOP")
			end
		end)
		button:SetScript("OnLeave", function()
			header.over = false
			paintHeader(key)
			if hint ~= nil then
				Theme.hideTooltip()
			end
		end)
		Theme.pressable(button, label, function()
			ground:Show()
		end, function()
			ground:Hide()
		end)
		button:SetScript("OnClick", function()
			GuildTab.setSort(key)
		end)
	end
end

--- The headers at `top` from the flyout's top (under the ratings notice when it shows).
local function layoutHeaders(flyout, top)
	for index, column in ipairs(GuildTab.COLUMNS) do
		local button = GuildTab.headers[column.key].button
		button:ClearAllPoints()
		button:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding + GuildTab.columnX(index), -top)
	end
end

--- The officer region's frame: a 1 px rule above its gold title, its lines, the push button
--- and the line under it after a push. No script of its own.
local function makeOfficer(flyout)
	local region = CreateFrame("Frame", nil, flyout)
	region:SetWidth(GuildTab.innerWidth())
	local rule = Theme.texture(region, "ARTWORK", "border")
	rule:SetSize(GuildTab.innerWidth(), S.border)
	rule:SetPoint("TOPLEFT", region, "TOPLEFT", 0, S.panelRegionGap / 2)
	region.title = sized(Widgets.label(region, L.guildTabOfficer, "gold", "small"), GuildTab.innerWidth())
	region.title:SetPoint("TOPLEFT", region, "TOPLEFT", 0, 0)
	for _, key in ipairs({ "approvals", "claim", "note", "pushed" }) do
		region[key] = sized(Widgets.label(region, "", "body", "small"), GuildTab.innerWidth())
		region[key]:Hide()
	end
	region.push = Widgets.button(region, L.guildTabPush, function()
		GuildTab.openPush()
	end)
	region.push:SetSize(S.guildPushWidth, S.panelButtonHeight)
	region:Hide()
	return region
end

--- Hides the flyout and the push confirm with it (L5.1 review item 3: a confirm left open
--- would outlive the flyout it asks about), and the button reads closed.
local function closeFlyout()
	if GuildTab.confirm ~= nil then
		GuildTab.confirm:Hide()
	end
	if GuildTab.flyout ~= nil then
		GuildTab.flyout:Hide()
	end
	GuildTab.syncButton()
end

--- The flyout, made on the first click: an addon panel `guildFlyoutWidth` wide parented to the
--- guild frame, hidden until shown; placed on each open. It carries no script; its close
--- button, switch, headers, rows, wheel and buttons carry their own.
function GuildTab.ensureFlyout()
	if GuildTab.flyout ~= nil then
		return GuildTab.flyout
	end
	local frame = pane()
	if frame == nil then
		return nil
	end
	local flyout = Widgets.panel(frame, S.guildFlyoutWidth, S.panelPadding * 2, GuildTab.FLYOUT_NAME)
	flyout:EnableMouse(true)
	flyout:Hide()
	GuildTab.flyout = flyout
	local header = Widgets.label(flyout, L.guildTabHeader, "gold", "normal")
	header:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding, -S.panelPadding)
	local close = Widgets.closeButton(flyout, function()
		closeFlyout()
	end)
	close:SetPoint("TOPRIGHT", flyout, "TOPRIGHT", -S.gap, -S.gap)
	-- The measure helper's own string, never shown.
	local probe = Widgets.label(flyout, "", "muted", "small")
	probe:Hide()
	GuildTab.measure = ItemTooltip.visibleMeasure(Theme.measurer(probe))
	GuildTab.line = Widgets.measuredLine(probe, GuildTab.measure, GuildTab.LINE_SAMPLE, S.panelLine)
	-- The switch row: "Online only" at the left, the count line muted at the right.
	local top = S.panelPadding + S.panelHeader + S.panelRowGap
	local toggle = Widgets.toggle(flyout, L.guildTabOnlineOnly, false, function(value)
		GuildTab.setOnlineOnly(value)
	end)
	toggle.frame:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding, -top)
	GuildTab.onlineToggle = toggle
	local counts = sized(Widgets.label(flyout, "", "muted", "small"), GuildTab.innerWidth() - S.guildToggleRoom)
	counts:SetJustifyH("RIGHT")
	counts:SetPoint("TOPRIGHT", flyout, "TOPRIGHT", -S.panelPadding,
		-(top + math.floor((S.guildToolbar - GuildTab.lineHeight()) / 2)))
	GuildTab.countsLine = counts
	-- Under it, one muted line while the ratings cannot be read or are old (item 1).
	local notice = sized(Widgets.label(flyout, "", "muted", "small"), GuildTab.innerWidth())
	notice:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding, -GuildTab.noticeTop())
	notice:Hide()
	GuildTab.noticeLine = notice
	makeHeaders(flyout)
	-- The rows: an addon-owned ScrollFrame clipping a content frame of pooled rows; the wheel
	-- moves the first row shown, three a notch (ruling 2).
	local holder = CreateFrame("ScrollFrame", nil, flyout)
	holder:SetSize(GuildTab.innerWidth(), GuildTab.rowHeight())
	local content = CreateFrame("Frame", nil, holder)
	content:SetSize(GuildTab.innerWidth(), GuildTab.rowHeight())
	if type(holder.SetScrollChild) == "function" then
		holder:SetScrollChild(content)
	end
	holder:EnableMouseWheel(true)
	holder:SetScript("OnMouseWheel", function(_, delta)
		GuildTab.scroll(-delta * GuildTab.WHEEL_STEP)
	end)
	GuildTab.holder, GuildTab.content = holder, content
	-- An 8-unit thumb on a track a step lighter than the border, straddling the rows' right
	-- edge (item 7).
	GuildTab.bar = Widgets.scrollbar(holder, holder, -S.guildScrollbar / 2, S.guildScrollbar, "scrollTrack")
	-- A 1 px rule in the last online row's bottom pixel when offline rows follow (item 7).
	local rule = Theme.texture(content, "ARTWORK", "border")
	rule:SetSize(GuildTab.innerWidth(), S.border)
	rule:Hide()
	GuildTab.onlineRule = rule
	local empty = sized(Widgets.label(content, "", "muted", "small"), GuildTab.innerWidth())
	empty:SetPoint("TOPLEFT", content, "TOPLEFT", 0, 0)
	empty:Hide()
	GuildTab.emptyLine = empty
	GuildTab.officerFrame = makeOfficer(flyout)
	local confirm = Widgets.confirmFrame(flyout, S.guildConfirmWidth, GuildTab.CONFIRM_NAME, L.guildTabCancel)
	confirm:SetPoint("CENTER", flyout, "CENTER", 0, 0)
	GuildTab.confirm = confirm
	return flyout
end

--- The headers' look: the sorted column's label gold with its arrow (down high to low, the
--- art flipped for low to high), the others muted (body under the cursor).
local function drawHeaders()
	local key = GuildTab.sort ~= nil and GuildTab.sort.key or "score"
	local descending = GuildTab.sort == nil or GuildTab.sort.descending == true
	for columnKey, header in pairs(GuildTab.headers) do
		paintHeader(columnKey)
		if columnKey == key then
			local coords = Theme.SORT_COORDS[descending and "descending" or "ascending"]
			header.arrow:SetTexCoord(coords[1], coords[2], coords[3], coords[4])
			header.arrow:Show()
		else
			header.arrow:Hide()
		end
	end
end

--- The index of the last online row when offline rows follow it in the default order (online
--- first), or nil.
function GuildTab.lastOnline(rows, sort)
	if sort ~= nil then
		return nil
	end
	for index = 1, #rows - 1 do
		if rows[index].online and not rows[index + 1].online then
			return index
		end
	end
	return nil
end

--- Draws the visible slice of the last model's rows from GuildTab.offset (clamped): each row
--- at one alpha (Theme.ALPHA.dim offline, item 4), the rule under the last online row, the
--- scrollbar, the count line with its range, and the empty line when no row shows.
function GuildTab.renderRows()
	local model = GuildTab.last
	if model == nil or GuildTab.flyout == nil then
		return
	end
	local rows, visible = model.rows, model.visible
	local first, last, offset = Widgets.visibleRange(#rows, visible, GuildTab.offset)
	GuildTab.offset = offset
	local height = GuildTab.rowHeight()
	local drawn = 0
	for index = first, last do
		drawn = drawn + 1
		local entry = rows[index]
		local row = rowFrame(drawn)
		row.entry = entry
		row.frame:ClearAllPoints()
		row.frame:SetPoint("TOPLEFT", GuildTab.content, "TOPLEFT", 0, -(drawn - 1) * height)
		row.frame:SetAlpha(entry.online and 1 or Theme.ALPHA.dim)
		local cells = GuildTab.cells(entry)
		for _, column in ipairs(GuildTab.COLUMNS) do
			local cell, value = row.cells[column.key], cells[column.key]
			cell:SetText(fitted(value.text, GuildTab.textWidth(column)))
			cell:SetTextColor(Theme.rgb(value.hex))
		end
		row.frame:Show()
	end
	for index = drawn + 1, #GuildTab.rowFrames do
		local row = GuildTab.rowFrames[index]
		row.entry = nil
		row.frame:Hide()
	end
	local divide = GuildTab.lastOnline(rows, GuildTab.sort)
	local rule = GuildTab.onlineRule
	if divide ~= nil and divide >= first and divide <= last then
		rule:ClearAllPoints()
		rule:SetPoint("TOPLEFT", GuildTab.content, "TOPLEFT", 0, -((divide - first + 1) * height - S.border))
		rule:Show()
	else
		rule:Hide()
	end
	local range = math.max(0, #rows - visible)
	GuildTab.bar:Update(#rows > 0 and math.min(1, visible / #rows) or 1, range > 0 and offset / range or 0)
	GuildTab.countsLine:SetText(fitted(GuildTab.countsText(model.counts, model.sharing, #rows > 0 and first or nil,
		#rows > 0 and last or nil, #rows), GuildTab.innerWidth() - S.guildToggleRoom))
	local empty = GuildTab.emptyLine
	if #rows == 0 then
		empty:SetText(fitted(model.counts.total == 0 and L.guildTabNoRoster or L.guildTabNobodyOnline,
			GuildTab.innerWidth()))
		empty:Show()
	else
		empty:Hide()
	end
end

--- Lays the officer region out under the rows at `y`; answers the y below it.
local function drawOfficer(flyout, officer, y)
	local region = GuildTab.officerFrame
	if officer == nil then
		region:Hide()
		return y
	end
	y = y + S.panelRegionGap
	region:ClearAllPoints()
	region:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding, -y)
	local inner = S.panelTitle
	local shown = {}
	for _, line in ipairs(officerLines(officer)) do
		local key = line.key
		local text = region[key]
		text:ClearAllPoints()
		text:SetPoint("TOPLEFT", region, "TOPLEFT", 0, -inner)
		text:SetText(fitted(line.text, GuildTab.innerWidth()))
		text:SetTextColor(Theme.rgb(Theme.HEX[line.color]))
		text:Show()
		shown[key] = true
		inner = inner + GuildTab.lineHeight()
	end
	for _, key in ipairs({ "approvals", "claim", "note" }) do
		if not shown[key] then
			region[key]:Hide()
		end
	end
	inner = inner + S.panelRowGap
	region.push:ClearAllPoints()
	region.push:SetPoint("TOPLEFT", region, "TOPLEFT", 0, -inner)
	Widgets.setEnabled(region.push, officer.canPush)
	inner = inner + S.panelButtonHeight
	local pushed = region.pushed
	if officer.pushed ~= nil then
		inner = inner + S.panelRowGap
		pushed:ClearAllPoints()
		pushed:SetPoint("TOPLEFT", region, "TOPLEFT", 0, -inner)
		pushed:SetText(fitted(officer.pushed.text, GuildTab.innerWidth()))
		pushed:SetTextColor(Theme.rgb(Theme.HEX[officer.pushed.failed and "warning" or "muted"]))
		pushed:Show()
		inner = inner + GuildTab.lineHeight()
	else
		pushed:Hide()
	end
	region:SetHeight(inner)
	region:Show()
	return y + inner
end

--- Draws the flyout from the live roster, the record, the snapshots, the player's own fields
--- and the inbox: the count line, the ratings notice, the headers, the rows (as many as fit
--- under the guild frame's height), the officer region; sized to them.
function GuildTab.draw()
	local flyout = GuildTab.flyout
	if flyout == nil then
		return nil
	end
	local rows, counts, onlineOnly, status = GuildTab.model()
	local notice = GuildTab.ratingsNotice(status)
	local _, rank = Theme.playerGuild()
	local officer = GuildTab.officerRegion(guildMessage(), rank, Follow.build, GuildTab.data or ns.Data)
	if officer ~= nil then
		officer.pushed = GuildTab.pushed
	end
	local hasNotice = notice ~= nil
	local visible = GuildTab.fitRows(Widgets.frameHeight(pane()), officer, hasNotice)
	GuildTab.last = { rows = rows, counts = counts, officer = officer, visible = visible, notice = notice,
		sharing = Comms.sharing() }
	GuildTab.onlineToggle:SetChecked(onlineOnly)
	if hasNotice then
		GuildTab.noticeLine:SetText(fitted(notice, GuildTab.innerWidth()))
		GuildTab.noticeLine:Show()
	else
		GuildTab.noticeLine:Hide()
	end
	layoutHeaders(flyout, GuildTab.headerTop(hasNotice))
	drawHeaders()
	local shownRows = math.max(1, math.min(visible, #rows))
	local holder = GuildTab.holder
	holder:ClearAllPoints()
	holder:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding, -GuildTab.rowsTop(hasNotice))
	holder:SetHeight(shownRows * GuildTab.rowHeight())
	GuildTab.content:SetHeight(shownRows * GuildTab.rowHeight())
	GuildTab.renderRows()
	local y = drawOfficer(flyout, officer, GuildTab.rowsTop(hasNotice) + shownRows * GuildTab.rowHeight())
	flyout:SetHeight(y + S.panelPadding)
	GuildTab.stale = nil
	GuildTab.draws = GuildTab.draws + 1
	return GuildTab.last
end

--- Redraws an open flyout while it is on screen; one the closed guild frame hides is marked
--- stale and redrawn when the frame shows again (the button's show script).
local function redraw()
	local flyout = GuildTab.flyout
	if flyout == nil or not flyout:IsShown() then
		return
	end
	if Widgets.isVisible(flyout) then
		GuildTab.draw()
	else
		GuildTab.stale = true
	end
end

--- An open flyout redrawn, guarded (GuildTab.seedLive's).
function GuildTab.redraw()
	guarded(redraw)
end

--- Asks for the roster and the guild's snapshots, each at most once a minute (rulings 2 and 3).
function GuildTab.askLive()
	local now = Theme.now()
	if GuildTab.lastRoster == nil or now - GuildTab.lastRoster >= GuildTab.ASK_EVERY then
		GuildTab.lastRoster = now
		Theme.requestGuildRoster()
	end
	if Comms.sharing() and (GuildTab.lastAsk == nil or now - GuildTab.lastAsk >= GuildTab.ASK_EVERY) then
		if Comms.send("snapshot?", {}, "GUILD") then
			GuildTab.lastAsk = now
		end
	end
end

--- The button reads open while the flyout is.
function GuildTab.syncButton()
	local button, flyout = GuildTab.button, GuildTab.flyout
	if button ~= nil then
		Widgets.setEdgedOpen(button, flyout ~= nil and flyout:IsShown())
	end
end

--- The button's click: the flyout opens (made on the first click; the roster and the
--- snapshots asked for; drawn) or closes, the push confirm with it.
function GuildTab.toggle()
	return guarded(function()
		local flyout = GuildTab.ensureFlyout()
		if flyout == nil then
			return false
		end
		if flyout:IsShown() then
			closeFlyout()
			return false
		end
		placeFlyout(flyout, pane())
		flyout:Show()
		GuildTab.syncButton()
		GuildTab.askLive()
		GuildTab.draw()
		return true
	end)
end

--- A header's click: sorts by that column, or turns it round; back to the first row.
function GuildTab.setSort(key)
	guarded(function()
		GuildTab.sort = GuildTab.nextSort(GuildTab.sort, key)
		GuildTab.offset = 0
		redraw()
	end)
end

--- The switch: Online only on or off; back to the first row.
function GuildTab.setOnlineOnly(value)
	guarded(function()
		GuildTab.onlineOnly = value == true
		GuildTab.offset = 0
		redraw()
	end)
end

--- The wheel: `delta` rows down (up when negative), clamped.
function GuildTab.scroll(delta)
	guarded(function()
		GuildTab.offset = GuildTab.offset + delta
		GuildTab.renderRows()
	end)
end

--- The push button: the addon's own confirm frame naming the active build; where no code can
--- be made, the officer region says why instead (L5.1 review item 4).
function GuildTab.openPush()
	guarded(function()
		if GuildTab.confirm == nil then
			return
		end
		local fields = GuildTab.pushFields(GuildTab.data or ns.Data, Follow.build)
		if fields == nil then
			local why = Follow.build == nil and L.guildTabPushNoBuild or L.guildTabPushNoCode
			GuildTab.pushed = { text = string.format(L.guildTabPushFailed, why), failed = true }
			redraw()
			return
		end
		local text = string.format(L.guildTabPushConfirm, fields[2])
		local room = S.guildConfirmWidth - S.padding * 2
		local width = type(GuildTab.measure) == "function" and GuildTab.measure(text) or nil
		local lines = width ~= nil and math.ceil(width / room) + 1 or 4
		Widgets.openConfirm(GuildTab.confirm, text, L.guildTabPushYes, function()
			GuildTab.push(fields)
		end, lines * GuildTab.lineHeight())
	end)
end

--- Sends the push on GUILD (Comms.send("push", { code, name })); the officer region then
--- says what went and when, or why it did not.
function GuildTab.push(fields)
	local ok, why = Comms.send("push", fields, "GUILD")
	if ok then
		GuildTab.pushed = { text = string.format(L.guildTabPushed, fields[2], date("%H:%M")) }
	else
		GuildTab.pushed = { text = string.format(L.guildTabPushFailed, tostring(why or "")), failed = true }
	end
	guarded(redraw)
	return ok
end

--- Every surface's refresh (Options): the button made when the guild frame exists, an open
--- flyout redrawn.
function GuildTab.refresh(data)
	if data ~= nil then
		GuildTab.data = data
	end
	guarded(function()
		GuildTab.ensureButton()
		redraw()
	end)
end

--- GUILD_ROSTER_UPDATE: an open flyout redraws.
function GuildTab.onRosterUpdate()
	guarded(redraw)
end

--- A snapshot! arrived (Comms has put it in Comms.seen): an open flyout redraws.
function GuildTab.onSnapshot()
	guarded(redraw)
end

--- The button's show script (it shows with the guild frame): an open flyout placed again and
--- redrawn once when an event made it stale while the frame was closed.
function GuildTab.onButtonShow()
	guarded(function()
		local frame = pane()
		local flyout = GuildTab.flyout
		if frame == nil or flyout == nil or not flyout:IsShown() then
			return
		end
		placeFlyout(flyout, frame)
		if GuildTab.stale and Widgets.isVisible(flyout) then
			GuildTab.draw()
		end
	end)
end

--- Combat ended: makes the button a combat lockdown held back.
function GuildTab.onCombatEnd()
	if not GuildTab.deferred then
		return
	end
	GuildTab.deferred = nil
	GuildTab.refresh()
end

--- Listens for the guild's snapshot! answers, once (Options.register).
function GuildTab.register()
	if GuildTab.registered then
		return
	end
	GuildTab.registered = true
	Comms.on("snapshot!", function()
		GuildTab.onSnapshot()
	end)
end

--- /fs diag's line (ruling 6): "Guild tab: open, 12 rows (2 live), officer no" (open: the
--- flyout and the guild frame both show), the rows counted now as the flyout counts them.
function GuildTab.diagLine()
	local ok, _, counts = pcall(GuildTab.model)
	if not ok or type(counts) ~= "table" then
		counts = { shown = 0, live = 0 }
	end
	local open = Widgets.isShown(GuildTab.flyout) and Widgets.isShown(pane())
	local _, rank = Theme.playerGuild()
	return string.format(L.diagGuildTab, open and L.diagPanelOpen or L.diagPanelClosed, counts.shown, counts.live,
		GuildTab.isOfficer(rank) and L.diagGuildTabYes or L.diagGuildTabNo)
end

ns.GuildTab = GuildTab
return GuildTab
