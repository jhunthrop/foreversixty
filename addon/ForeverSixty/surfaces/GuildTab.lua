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
--- The columns, left to right: the sort key, its width, and whether it sorts as a number
--- (a number's first click sorts high to low; a name's A to Z).
GuildTab.COLUMNS = {
	{ key = "name", width = 108 },
	{ key = "level", width = 34, numeric = true },
	{ key = "spec", width = 82 },
	{ key = "gear", width = 36, numeric = true },
	{ key = "score", width = 50, numeric = true },
	{ key = "attendance", width = 64, numeric = true },
	{ key = "gap", width = 52, numeric = true },
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

--- One row: the roster's name, class, rank and online flag; the live snapshot's level, spec,
--- gap and build where the member's addon answered (`live`); the record's score, tier, gear
--- score, attendance (a whole percent) and, without a snapshot, spec.
local function rowOf(entry, record, seen)
	local found = recordFor(record, entry.name)
	local snapshot = GuildTab.seenFor(seen, entry.name)
	local row = {
		name = entry.name, short = GuildTab.shortName(entry.name), classToken = entry.classToken,
		online = entry.online == true, rankIndex = entry.rankIndex, level = entry.level, live = snapshot ~= nil,
		record = found,
	}
	if snapshot ~= nil then
		row.level = tonumber(snapshot.level) or row.level
		row.gap = tonumber(snapshot.gearGap)
		if type(snapshot.spec) == "string" and snapshot.spec ~= "" and snapshot.spec ~= "-" then
			row.spec = snapshot.spec
		end
		if type(snapshot.buildName) == "string" and snapshot.buildName ~= "" then
			row.buildName = snapshot.buildName
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
--- by `filter` ({ onlineOnly }). Answers the rows and { total, shown, live } (live: shown rows
--- a member's addon answered for).
function GuildTab.rows(roster, record, seen, sort, filter)
	local shown, total, live = {}, 0, 0
	local onlineOnly = type(filter) == "table" and filter.onlineOnly == true
	for _, entry in ipairs(type(roster) == "table" and roster or {}) do
		if type(entry) == "table" and type(entry.name) == "string" and entry.name ~= "" then
			total = total + 1
			if not onlineOnly or entry.online == true then
				local row = rowOf(entry, record, seen)
				shown[#shown + 1] = row
				if row.live then
					live = live + 1
				end
			end
		end
	end
	table.sort(shown, comparator(sort))
	return shown, { total = total, shown = #shown, live = live }
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

--- A row's cells, each { text, hex } (`dim` on the name of a member offline): the name in its
--- class colour, the score in its tier's colour ("no logs" muted without one), attendance as
--- "80%", the gap "14 (live)"; a value nobody has a muted dash.
function GuildTab.cells(row)
	local cells = {
		name = { text = row.short, hex = row.classToken ~= nil and Theme.classHex(row.classToken) or Theme.HEX.body,
			dim = not row.online },
		level = type(row.level) == "number" and plain(whole(row.level)) or muted(L.guildTabDash),
		spec = row.spec ~= nil and plain(Follow.specName(row.spec)) or muted(L.guildTabDash),
		gear = type(row.gearScore) == "number" and plain(whole(row.gearScore)) or muted(L.guildTabDash),
		attendance = type(row.attendance) == "number" and plain(string.format(L.guildTabPercent, row.attendance))
			or muted(L.guildTabDash),
		gap = type(row.gap) == "number" and plain(string.format(L.guildTabLive, math.floor(row.gap + 0.5)))
			or muted(L.guildTabDash),
	}
	if type(row.score) == "number" then
		local r, g, b = Score.tierColor(row.tier)
		cells.score = { text = whole(row.score), hex = r ~= nil and Theme.hexOf(r, g, b) or Theme.HEX.body }
	else
		cells.score = muted(L.guildTabNoLogs)
	end
	return cells
end

--- A row's hover: the full name in its class colour, the score line (or "no logs yet"), the
--- live line and build where the member's addon answered, and "Offline".
function GuildTab.rowHint(row)
	local cells = GuildTab.cells(row)
	local lines = { "|cff" .. cells.name.hex .. row.name .. "|r" }
	local scoreLines = Score.lines(row.record)
	if scoreLines[1] ~= nil then
		lines[#lines + 1] = "|cff" .. cells.score.hex .. scoreLines[1] .. "|r"
	else
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", L.scoreNoLogs)
	end
	if row.live then
		local spec = row.spec ~= nil and Follow.specName(row.spec) or L.guildTabDash
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", string.format(L.scoreLive, row.level or 0, spec,
			math.floor((row.gap or 0) + 0.5)))
		if row.buildName ~= nil then
			lines[#lines + 1] = Theme.inlineColor("tooltipText", string.format(L.guildTabBuild, row.buildName))
		end
	end
	if not row.online then
		lines[#lines + 1] = Theme.inlineColor("tooltipMuted", L.guildTabOffline)
	end
	return lines
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

--- Whether a guild rank index may push a build: Comms' own rule (rank index PUSH_RANK or
--- lower; 0 is the guild master).
function GuildTab.isOfficer(rankIndex)
	return type(rankIndex) == "number" and rankIndex <= Comms.PUSH_RANK
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

--- Where the roster rows start, from the flyout's top: the padding, the header, the switch
--- row, the column headers.
function GuildTab.rowsTop()
	return S.panelPadding + S.panelHeader + S.panelRowGap + S.guildToolbar + S.panelRowGap + S.guildHeaderRow
end

--- The flyout's height with `rows` roster rows and `officer` (its region gap included).
function GuildTab.flyoutHeight(rows, officer)
	local height = GuildTab.rowsTop() + rows * GuildTab.rowHeight()
	if officer ~= nil then
		height = height + S.panelRegionGap + GuildTab.officerHeight(officer)
	end
	return height + S.panelPadding
end

--- How many roster rows show at once under `limit` (the guild frame's height; the flyout is
--- never taller): GuildTab.VISIBLE_ROWS at most, one at least; all of them where the client
--- cannot say.
function GuildTab.fitRows(limit, officer)
	if type(limit) ~= "number" or limit <= 0 then
		return GuildTab.VISIBLE_ROWS
	end
	local room = math.floor((limit - GuildTab.flyoutHeight(0, officer)) / GuildTab.rowHeight())
	return math.max(1, math.min(GuildTab.VISIBLE_ROWS, room))
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
--- shows with a rect (level with the frame's top); at that tab's own TOPRIGHT where the edges
--- cannot be read.
local function placeFlyout(flyout, frame)
	local column = sideColumn(frame)
	if column ~= nil then
		local right, own = Widgets.frameEdge(column, "GetRight"), Widgets.frameEdge(frame, "GetRight")
		if right ~= nil and own ~= nil then
			flyout:ClearAllPoints()
			flyout:SetPoint("TOPLEFT", frame, "TOPRIGHT", math.max(0, right - own) + S.panelOffsetX, S.panelOffsetY)
			return
		end
	end
	Widgets.placeFlyout(flyout, frame, column, S.panelOffsetX, S.panelOffsetY)
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
--- guild frame (never in combat: it waits for its end). nil without one.
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
		GuildTab.button = button
	end
	return button
end

--- The inner width: the flyout less its padding.
function GuildTab.innerWidth()
	return S.guildFlyoutWidth - S.panelPadding * 2
end

--- A cell's x within a row.
local function columnX(index)
	local x = 0
	for before = 1, index - 1 do
		x = x + GuildTab.COLUMNS[before].width + S.guildColumnGap
	end
	return x
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

--- One pooled roster row: a Button the width of the rows, a hover ground, a cell per column.
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
		local cell = sized(Widgets.label(frame, "", "body", "small"), spec.width)
		cell:SetPoint("LEFT", frame, "LEFT", columnX(column), 0)
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

--- The column headers: an addon Button per column, its label and the client's sort arrow.
local function makeHeaders(flyout)
	local y = S.panelPadding + S.panelHeader + S.panelRowGap + S.guildToolbar + S.panelRowGap
	for index, column in ipairs(GuildTab.COLUMNS) do
		local button = CreateFrame("Button", nil, flyout)
		button:SetSize(column.width, S.guildHeaderRow)
		button:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding + columnX(index), -y)
		local label = Widgets.label(button, L.guildTabColumns[column.key], "muted", "small")
		label:SetPoint("LEFT", button, "LEFT", 0, 0)
		local arrow = Theme.tinted(button, "OVERLAY", Theme.MEDIA.sortArrow, "gold")
		arrow:SetSize(S.guildSortArrow, S.guildSortArrow)
		arrow:SetPoint("LEFT", label, "RIGHT", 2, 0)
		arrow:Hide()
		button:SetScript("OnClick", function()
			GuildTab.setSort(column.key)
		end)
		GuildTab.headers[column.key] = { button = button, label = label, arrow = arrow }
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
		flyout:Hide()
		GuildTab.syncButton()
	end)
	close:SetPoint("TOPRIGHT", flyout, "TOPRIGHT", -S.gap, -S.gap)
	-- The measure helper's own string, never shown.
	local probe = Widgets.label(flyout, "", "muted", "small")
	probe:Hide()
	GuildTab.measure = ItemTooltip.visibleMeasure(Theme.measurer(probe))
	GuildTab.line = Widgets.measuredLine(probe, GuildTab.measure, GuildTab.LINE_SAMPLE, S.panelLine)
	-- The switch row: "Online only" at the left, the counts muted at the right.
	local top = S.panelPadding + S.panelHeader + S.panelRowGap
	local toggle = Widgets.toggle(flyout, L.guildTabOnlineOnly, false, function(value)
		GuildTab.setOnlineOnly(value)
	end)
	toggle.frame:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding, -top)
	GuildTab.onlineToggle = toggle
	local counts = sized(Widgets.label(flyout, "", "muted", "small"), GuildTab.innerWidth() / 2)
	counts:SetJustifyH("RIGHT")
	counts:SetPoint("TOPRIGHT", flyout, "TOPRIGHT", -S.panelPadding,
		-(top + math.floor((S.guildToolbar - GuildTab.lineHeight()) / 2)))
	GuildTab.countsLine = counts
	makeHeaders(flyout)
	-- The rows: an addon-owned ScrollFrame clipping a content frame of pooled rows; the wheel
	-- moves the first row shown, three a notch (ruling 2).
	local holder = CreateFrame("ScrollFrame", nil, flyout)
	holder:SetPoint("TOPLEFT", flyout, "TOPLEFT", S.panelPadding, -GuildTab.rowsTop())
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
	GuildTab.bar = Widgets.scrollbar(holder, holder, -S.scrollbar)
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

--- The headers' look: the sorted column's label gold with its arrow (down high to low), the
--- others muted.
local function drawHeaders()
	local key = GuildTab.sort ~= nil and GuildTab.sort.key or "score"
	local descending = GuildTab.sort == nil or GuildTab.sort.descending == true
	for columnKey, header in pairs(GuildTab.headers) do
		local active = columnKey == key
		header.label:SetTextColor(Theme.rgb(Theme.HEX[active and "gold" or "muted"]))
		if active then
			local coords = Theme.SORT_COORDS[descending and "descending" or "ascending"]
			header.arrow:SetTexCoord(coords[1], coords[2], coords[3], coords[4])
			header.arrow:Show()
		else
			header.arrow:Hide()
		end
	end
end

--- Draws the visible slice of the last model's rows from GuildTab.offset (clamped), the
--- scrollbar, and the empty line when no row shows.
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
		local cells = GuildTab.cells(entry)
		for _, column in ipairs(GuildTab.COLUMNS) do
			local cell, value = row.cells[column.key], cells[column.key]
			cell:SetText(fitted(value.text, column.width))
			cell:SetTextColor(Theme.rgb(value.hex))
			cell:SetAlpha(entry.online and 1 or Theme.ALPHA.dim)
		end
		row.frame:Show()
	end
	for index = drawn + 1, #GuildTab.rowFrames do
		local row = GuildTab.rowFrames[index]
		row.entry = nil
		row.frame:Hide()
	end
	local range = math.max(0, #rows - visible)
	GuildTab.bar:Update(#rows > 0 and math.min(1, visible / #rows) or 1, range > 0 and offset / range or 0)
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

--- Draws the flyout from the live roster, the record, the snapshots and the inbox: the
--- counts, the headers, the rows (as many as fit under the guild frame's height), the officer
--- region; sized to them.
function GuildTab.draw()
	local flyout = GuildTab.flyout
	if flyout == nil then
		return nil
	end
	local roster = Theme.guildRoster()
	local onlineOnly = GuildTab.onlineOnly
	if onlineOnly == nil then
		onlineOnly = GuildTab.defaultOnlineOnly(#roster)
	end
	local rows, counts = GuildTab.rows(roster, Score.forCharacter, Comms.seen, GuildTab.sort,
		{ onlineOnly = onlineOnly })
	local _, rank = Theme.playerGuild()
	local officer = GuildTab.officerRegion(guildMessage(), rank, Follow.build, GuildTab.data or ns.Data)
	if officer ~= nil then
		officer.pushed = GuildTab.pushed
	end
	local visible = GuildTab.fitRows(Widgets.frameHeight(pane()), officer)
	GuildTab.last = { rows = rows, counts = counts, officer = officer, visible = visible }
	GuildTab.onlineToggle:SetChecked(onlineOnly)
	GuildTab.countsLine:SetText(string.format(L.guildTabCounts, counts.shown, counts.total, counts.live))
	drawHeaders()
	local shownRows = math.max(1, math.min(visible, #rows))
	GuildTab.holder:SetHeight(shownRows * GuildTab.rowHeight())
	GuildTab.content:SetHeight(shownRows * GuildTab.rowHeight())
	GuildTab.renderRows()
	local y = drawOfficer(flyout, officer, GuildTab.rowsTop() + shownRows * GuildTab.rowHeight())
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
--- snapshots asked for; drawn) or closes.
function GuildTab.toggle()
	return guarded(function()
		local flyout = GuildTab.ensureFlyout()
		if flyout == nil then
			return false
		end
		if flyout:IsShown() then
			flyout:Hide()
			GuildTab.syncButton()
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

--- The push button: the addon's own confirm frame naming the active build.
function GuildTab.openPush()
	guarded(function()
		local fields = GuildTab.pushFields(GuildTab.data or ns.Data, Follow.build)
		if fields == nil or GuildTab.confirm == nil then
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
--- flyout and the guild frame both show), the rows counted now.
function GuildTab.diagLine()
	local ok, _, counts = pcall(function()
		local roster = Theme.guildRoster()
		local onlineOnly = GuildTab.onlineOnly
		if onlineOnly == nil then
			onlineOnly = GuildTab.defaultOnlineOnly(#roster)
		end
		return GuildTab.rows(roster, Score.forCharacter, Comms.seen, GuildTab.sort, { onlineOnly = onlineOnly })
	end)
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
