-- addon/ForeverSixty/surfaces/Hub.lua
-- The Forever Sixty window (design/specs/2026-11-13-addon-hub.md section 3; the frame and the Builds tab:
-- 2026-11-13-addon-builds-window.md): the addon's one framed window, tabbed like the game's character
-- and talent frames. The hub owns everything the tabs share: the header (the class crest, the character's
-- name, the figures line), the tab strip and its dots, the footer, the height, where the window sits,
-- Escape, which tab an open lands on, and the one officer gate. Each tab owns only its body.
--
-- The API the tab lanes build on (keep it stable; docs/superpowers/plans/2026-11-13-addon-hub.md, "How to
-- add a tab"):
--
--   Hub.register { id, label, build(content), show(options), hide(), escape(), empty(), figures(),
--                  waiting(), lead(), available(), order }
--     id        one of Hub.ORDER's keys (me, builds, gear, rotation, guild, officer); `order` and the
--               shared availability rule (Hub.RULES, spec 3.3) come from the id. Registering an id again
--               replaces the tab (the placeholders below are replaced this way).
--     build     once, into the tab's own content frame (the body under the header and the strip).
--     show      the tab comes on screen with the options Hub.open passed: draw, Hub.setHints, Hub.fitHeight.
--     hide      it leaves (another tab, or the window closed).
--     escape    true when the tab stepped back itself (a preview closed); else Escape closes the window.
--     empty     a muted line shown instead of the body while the tab has nothing to show, or nil.
--     figures   at most Hub.FIGURES_MAX parts for the header after `Level 60 Hunter` (spec 3.2), and
--               optionally the lines the figures' hover shows. The title is always the character's name.
--     waiting   the items this tab has for the player now, each { key, line, fact }: `key` names the kind
--               ("spend"), `line` is what the chip says, `fact` changes when the item is new again
--               (default: the line). While it returns any, the strip draws the tab's dot.
--     lead      the line the tab leads with on the Me tab's Right now (spec 4.1), from the tab's own
--               function so the two never disagree: { line, sub, icon (a texture path) or itemId }, or
--               nil. A tab without one shows its first waiting line there, if any.
--     available an extra rule on top of the shared one (both must pass); an error counts as false.
--
--   Hub.open(id, options) / Hub.toggle(id) / Hub.close() / Hub.isShown() / Hub.select(id, options)
--     open with no id lands on Hub.landing() (the first unseen waiting item's tab), else the last tab.
--   Hub.refresh()              a tab's figures or waiting items changed: header and strip redrawn.
--   Hub.setHints(escape, right) the footer's two lines; Hub.fitHeight(body) the tab's body height.
--   Hub.urgent()               { tab, line, more, items } (the first waiting item, seen or not), or nil.
--   Hub.reason(id)             why tab `id` cannot open now (one line), or nil when it can.
--   Hub.officer()              the player passes the officer gate (the client's officer-note right).
--   Hub.officerRank(rankIndex) a guild rank passes it (a received push: the sender's rank).
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Prefs = ns.Prefs or require("Prefs")
local Compat = ns.Compat or require("Compat")
local Talents = ns.Talents or require("Talents")
local HubFrame = ns.HubFrame or require("HubFrame")

local Hub = {}

Hub.FRAME_NAME = "ForeverSixtyWindow"
--- Pixel layout: the header band and its crest, the tab strip (only while two tabs or more show), the
--- body's padding top and bottom, the footer band, the gap between tabs, the tab dot and its gap.
Hub.LAYOUT = { header = 56, crest = 32, crestGap = 10, strip = 28, bodyPad = 12, footer = 28, tabGap = 4,
	dot = 6, dotGap = 4 }
--- The footer's hints: the muted token at this alpha, a step dimmer than the text.
Hub.HINT_ALPHA = 0.75

--- The strip's order (spec 3.3): Me, Builds, Gear, Rotation, Guild, Officer.
Hub.ORDER = { me = 1, builds = 2, gear = 3, rotation = 4, guild = 5, officer = 6 }
--- Which tab's waiting items come first (spec 3.6): the raid tonight, a loot session, the talents, a bag
--- upgrade, a raid night not yet seen. Rotation has none.
Hub.URGENT_ORDER = { "guild", "officer", "builds", "gear", "me" }
--- The header's figures after `Level 60 Hunter`: at most this many parts.
Hub.FIGURES_MAX = 2
--- Gear shows from this level (spec 3.3).
Hub.GEAR_MIN_LEVEL = 10
--- The rank flag of the client's guild control that is "Edit officer note" (GuildControlGetRankFlags).
Hub.EDIT_OFFICER_NOTE_FLAG = 12

--- The shipped data (ns.Data), for the rules that read it; Options sets it.
Hub.data = nil

--- Injection points: `fallback` runs instead when the window turned itself off this session (Options:
--- the command list); the tour fills `beforeOpen` and `afterHide` (it steps aside while the window is
--- open and comes back when it hides). `disabled` is set by a tab's render check (off screen).
Hub.fallback = nil
Hub.beforeOpen = nil
Hub.afterHide = nil
Hub.disabled = false

--- Every tab, in strip order. Filled by Hub.register at load.
Hub.tabs = {}

-- The officer gate ---------------------------------------------------------------------------------

local function ask(fn, ...)
	if type(fn) ~= "function" then
		return nil
	end
	local ok, answer = pcall(fn, ...)
	return ok and answer or nil
end

--- Whether the player passes the officer gate: the client's own officer-note right
--- (C_GuildInfo.CanEditOfficerNote, else the global CanEditOfficerNote). A guild that gives a rank
--- officer-note rights trusts it with the roster (spec 3.3); false outside a guild or where the client
--- cannot say.
function Hub.officer()
	local info = _G.C_GuildInfo
	local answer
	if type(info) == "table" then
		answer = ask(info.CanEditOfficerNote)
	end
	if answer == nil then
		answer = ask(_G.CanEditOfficerNote)
	end
	return answer == true
end

--- Whether guild rank `rankIndex` (0 is the guild master) passes the same gate: the guild master
--- always; the player's own rank as Hub.officer answers; another rank by its "Edit officer note" flag
--- where the client lets this player read it (C_GuildInfo.GuildControlGetRankFlags, else the global),
--- false where it does not.
function Hub.officerRank(rankIndex)
	if type(rankIndex) ~= "number" then
		return false
	end
	if rankIndex == 0 then
		return true
	end
	local _, own = Theme.playerGuild()
	if own == rankIndex then
		return Hub.officer()
	end
	local info = _G.C_GuildInfo
	local flags = type(info) == "table" and ask(info.GuildControlGetRankFlags, rankIndex + 1) or nil
	if type(flags) ~= "table" then
		flags = { ask(_G.GuildControlGetRankFlags, rankIndex + 1) }
	end
	return flags[Hub.EDIT_OFFICER_NOTE_FLAG] == true
end

-- Availability ---------------------------------------------------------------------------------------

--- Whether the shipped data has a rotation for any spec of `classSlug` (Data.rotations keys are
--- "<class>-<spec>"). Pure.
function Hub.hasRotation(data, classSlug)
	local rotations = type(data) == "table" and data.rotations or nil
	if type(rotations) ~= "table" or type(classSlug) ~= "string" then
		return false
	end
	local prefix = classSlug .. "-"
	for spec in pairs(rotations) do
		if type(spec) == "string" and spec:sub(1, #prefix) == prefix then
			return true
		end
	end
	return false
end

--- Whether the player runs a master loot session now (the Officer tab shows its Master loot section
--- to a master looter without the gate).
function Hub.masterLooting()
	local LootCouncil = ns.LootCouncil or package.loaded.LootCouncil
	if type(LootCouncil) ~= "table" then
		return false
	end
	return ask(LootCouncil.isMasterLooter) == true and ask(LootCouncil.active) == true
end

--- The shared rule of each tab that is not always shown (spec 3.3), and the line that says why not.
Hub.RULES = {
	gear = function()
		local level = Compat.playerLevel()
		return type(level) == "number" and level >= Hub.GEAR_MIN_LEVEL
	end,
	rotation = function()
		return Hub.hasRotation(Hub.data or ns.Data, ask(Talents.playerClassSlug))
	end,
	guild = function()
		return Theme.inGuild() == true
	end,
	officer = function()
		return (Theme.inGuild() == true and Hub.officer()) or Hub.masterLooting()
	end,
}
Hub.REASONS = {
	gear = string.format(L.hubGearLevel, Hub.GEAR_MIN_LEVEL),
	rotation = L.hubNoRotation,
	guild = L.hubNotInGuild,
	officer = L.offOnly,
}

local function passes(check)
	if check == nil then
		return true
	end
	local ok, answer = pcall(check)
	return ok and answer == true
end

--- The window's width: the Builds tab's two columns, the gutter between them, the padding.
function Hub.width()
	local S = Theme.SIZES
	return S.padding * 2 + S.windowLeft + S.windowGutter + S.windowRight
end

-- The bodies' placement, defined with the frame below; a late registration reaches it at call time.
local placeBodies

--- Add or replace a tab (by `id`), its order from Hub.ORDER unless it names one. Returns the tab.
function Hub.register(tab)
	assert(type(tab) == "table" and type(tab.id) == "string", "Hub.register needs a tab with an id")
	tab.order = tab.order or Hub.ORDER[tab.id]
	for index, existing in ipairs(Hub.tabs) do
		if existing.id == tab.id then
			Hub.tabs[index] = tab
			if Hub.frame ~= nil then
				-- A lane filling a placeholder after the window was built: a fresh body and button.
				Hub.contents[tab.id]:Hide()
				Hub.strip.buttons[tab.id]:Hide()
				Hub.addTabFrames(tab)
				placeBodies()
			end
			return tab
		end
	end
	Hub.tabs[#Hub.tabs + 1] = tab
	table.sort(Hub.tabs, function(a, b)
		return (a.order or 99) < (b.order or 99)
	end)
	if Hub.frame ~= nil then
		-- A tab registered after the window was built gets its body and its strip button now.
		Hub.addTabFrames(tab)
	end
	return tab
end

--- Whether `tab` shows: not a placeholder, its id's shared rule (Hub.RULES) and its own `available`
--- both pass (an error is "not available").
function Hub.isAvailable(tab)
	if type(tab) ~= "table" or tab.placeholder == true then
		return false
	end
	return passes(Hub.RULES[tab.id]) and passes(tab.available)
end

--- Why tab `id` cannot open now, in one line (`/fs gear` below level 10: "Gear opens at level 10."),
--- or nil when it can.
function Hub.reason(id)
	local tab = Hub.tab(id)
	if tab == nil or Hub.isAvailable(tab) then
		return nil
	end
	if tab.placeholder ~= true and not passes(Hub.RULES[tab.id]) and Hub.REASONS[tab.id] ~= nil then
		return Hub.REASONS[tab.id]
	end
	return L.hubTabSoon
end

--- The tabs the strip shows now, in order. Pure over the registry.
function Hub.visibleTabs()
	local shown = {}
	for _, tab in ipairs(Hub.tabs) do
		if Hub.isAvailable(tab) then
			shown[#shown + 1] = tab
		end
	end
	return shown
end

--- The tab `id`, or nil.
function Hub.tab(id)
	for _, tab in ipairs(Hub.tabs) do
		if tab.id == id then
			return tab
		end
	end
	return nil
end

--- Which tab an open lands on: the one asked for when it shows, else the one the player last had
--- (`remembered`) when it shows, else the first that shows; nil when none does. Pure over the registry.
function Hub.resolveTab(asked, remembered)
	for _, id in ipairs({ asked or false, remembered or false }) do
		local tab = type(id) == "string" and id ~= "" and Hub.tab(id) or nil
		if tab ~= nil and Hub.isAvailable(tab) then
			return tab.id
		end
	end
	local first = Hub.visibleTabs()[1]
	return first ~= nil and first.id or nil
end

-- The header contract (spec 3.2) -------------------------------------------------------------------

--- `Level 60 Hunter`: the first part of the figures line on every tab; nil where the client cannot say.
function Hub.identity()
	local level = Compat.playerLevel()
	local _, className = Compat.playerRaceClass()
	if type(level) ~= "number" or type(className) ~= "string" then
		return nil
	end
	return string.format(L.hubIdentity, level, className)
end

--- A tab's figures: at most Hub.FIGURES_MAX non-empty parts, and the lines its hover shows (or nil).
function Hub.figuresOf(tab)
	if type(tab) ~= "table" or tab.figures == nil then
		return {}, nil
	end
	local ok, parts, lines = pcall(tab.figures)
	local kept = {}
	if ok and type(parts) == "table" then
		for _, part in ipairs(parts) do
			if type(part) == "string" and part ~= "" and #kept < Hub.FIGURES_MAX then
				kept[#kept + 1] = part
			end
		end
	end
	return kept, ok and type(lines) == "table" and lines or nil
end

--- The header for `tab`: { classToken, title, note, noteLines }. The title is the character's name on
--- every tab (the addon's name where the client gives none); the note is the identity, then the tab's
--- figures, joined as the site joins them.
function Hub.headerFor(tab)
	local _, _, classToken = Compat.playerRaceClass()
	local parts, lines = Hub.figuresOf(tab)
	local note = { Hub.identity() }
	for _, part in ipairs(parts) do
		note[#note + 1] = part
	end
	return { classToken = classToken, title = Compat.playerName() or L.addonName,
		note = #note > 0 and table.concat(note, L.dialogSeparator) or nil, noteLines = lines }
end

-- What waits, and which tab an open lands on (spec 3.6) -----------------------------------------------

local function characterKey()
	local Export = ns.Export or package.loaded.Export
	local key = type(Export) == "table" and ask(Export.characterKey) or nil
	return key or Compat.playerName() or ""
end

--- A tab's waiting items, each { tab, key, line, fact } (an item without a key or line is dropped; the
--- fact defaults to the line); none for a tab that is not available.
function Hub.waitingOf(tab)
	if not Hub.isAvailable(tab) or tab.waiting == nil then
		return {}
	end
	local ok, items = pcall(tab.waiting)
	local kept = {}
	for _, item in ipairs(ok and type(items) == "table" and items or {}) do
		if type(item) == "table" and type(item.key) == "string" and type(item.line) == "string" then
			kept[#kept + 1] = { tab = tab.id, key = item.key, line = item.line, fact = tostring(item.fact or item.line) }
		end
	end
	return kept
end

--- Every waiting item, the Hub.URGENT_ORDER tabs first in that order, then any other tab in strip order.
function Hub.items()
	local ordered, placed = {}, {}
	for _, id in ipairs(Hub.URGENT_ORDER) do
		placed[id] = true
		ordered[#ordered + 1] = Hub.tab(id)
	end
	for _, tab in ipairs(Hub.tabs) do
		if not placed[tab.id] then
			ordered[#ordered + 1] = tab
		end
	end
	local items = {}
	for _, tab in ipairs(ordered) do
		for _, item in ipairs(Hub.waitingOf(tab)) do
			items[#items + 1] = item
		end
	end
	return items
end

--- The most urgent thing, seen or not (the chip's line): { tab, line, more = the others' count, items },
--- or nil with nothing waiting.
function Hub.urgent()
	local items = Hub.items()
	if #items == 0 then
		return nil
	end
	return { tab = items[1].tab, line = items[1].line, more = #items - 1, items = items }
end

local function seenKey(item)
	return characterKey() .. "|" .. item.tab .. ":" .. item.key
end

--- Whether the player opened `item`'s tab since its fact last changed.
function Hub.isSeen(item)
	return Prefs.get("hub", "seen")[seenKey(item)] == item.fact
end

--- Opening tab `id` marks what it has waiting seen; an item comes back when its fact changes.
function Hub.markSeen(id)
	local seen = Prefs.get("hub", "seen")
	for _, item in ipairs(Hub.waitingOf(Hub.tab(id))) do
		seen[seenKey(item)] = item.fact
	end
end

--- Where an open with no tab asked lands: the first unseen item's tab, or nil (the last tab, then).
function Hub.landing()
	for _, item in ipairs(Hub.items()) do
		if not Hub.isSeen(item) then
			return item.tab
		end
	end
	return nil
end

--- The strip's height: none with a single tab showing (one tab is not a choice), else Hub.LAYOUT.strip.
function Hub.stripHeight()
	return #Hub.visibleTabs() > 1 and Hub.LAYOUT.strip or 0
end

--- From the window's top to its body: the header and the strip.
function Hub.bodyTop()
	return Hub.LAYOUT.header + Hub.stripHeight()
end

--- The window's height for `body` px of content, never under Theme.SIZES.dialogHeight. Pure.
function Hub.heightFor(body)
	local G = Hub.LAYOUT
	return math.max(Theme.SIZES.dialogHeight, math.ceil(Hub.bodyTop() + G.bodyPad * 2 + (body or 0) + G.footer))
end

-- The frame ------------------------------------------------------------------------------------

local line = HubFrame.line

--- Lay the strip out for the tabs available now: left to right from the padding, the others hidden.
local function drawStrip()
	local strip, G = Hub.strip, Hub.LAYOUT
	local shown = Hub.visibleTabs()
	Theme.setShown({ strip.frame }, #shown > 1)
	for _, button in pairs(strip.buttons) do
		button:Hide()
	end
	local previous
	for _, tab in ipairs(shown) do
		local button = strip.buttons[tab.id]
		if button ~= nil then
			button:ClearAllPoints()
			if previous == nil then
				button:SetPoint("BOTTOMLEFT", strip.frame, "BOTTOMLEFT", Theme.SIZES.padding, 0)
			else
				button:SetPoint("LEFT", previous, "RIGHT", G.tabGap, 0)
			end
			Widgets.setTabActive(button, tab.id == Hub.active)
			Theme.setShown({ button.foreverSixtyDot }, #Hub.waitingOf(tab) > 0)
			button:Show()
			previous = button
		end
	end
end

--- The body under the header and the strip, above the footer: each tab's content frame fills it.
function placeBodies()
	local G = Hub.LAYOUT
	for _, tab in ipairs(Hub.tabs) do
		local content = Hub.contents[tab.id]
		if content ~= nil then
			content:ClearAllPoints()
			content:SetPoint("TOPLEFT", Hub.frame, "TOPLEFT", 0, -Hub.bodyTop())
			content:SetPoint("BOTTOMRIGHT", Hub.frame, "BOTTOMRIGHT", 0, G.footer)
		end
	end
end

--- A tab's content frame and strip button, its body built into the frame by the tab itself.
function Hub.addTabFrames(tab)
	local S = Theme.SIZES
	local content = Theme.createFrame("Frame", nil, Hub.frame)
	content:Hide()
	Hub.contents[tab.id] = content
	-- The tab's empty state: one muted line at the body's top left (Hub.select shows it).
	local empty = line(content, "muted", Hub.width() - S.padding * 2)
	if type(empty.SetWordWrap) == "function" then
		empty:SetWordWrap(true)
	end
	empty:SetPoint("TOPLEFT", content, "TOPLEFT", S.padding, -Hub.LAYOUT.bodyPad)
	empty:Hide()
	Hub.empties[tab.id] = empty
	HubFrame.tabButton(Hub.strip, tab, Hub)
	if tab.build ~= nil then
		tab.build(content)
	end
	return content
end

--- The window, built on the first open: the frame, the header, the strip, the footer and one content
--- frame per tab, each tab building its own body into it.
function Hub.ensure()
	if Hub.frame ~= nil then
		return Hub.frame
	end
	local S = Theme.SIZES
	local frame = Widgets.panel(UIParent, Hub.width(), S.dialogHeight, Hub.FRAME_NAME)
	frame:Hide()
	frame:SetFrameStrata("DIALOG")
	frame:SetClampedToScreen(true)
	frame:SetMovable(true)
	frame:EnableMouse(true)
	Hub.frame = frame
	Hub.height = S.dialogHeight
	Hub.header = HubFrame.header(frame, Hub)
	Hub.strip = HubFrame.strip(frame, Hub)
	Hub.footer = HubFrame.footer(frame, Hub)
	Hub.contents, Hub.empties, Hub.bodies = {}, {}, {}
	for _, tab in ipairs(Hub.tabs) do
		Hub.addTabFrames(tab)
	end
	placeBodies()
	frame:SetScript("OnHide", function()
		local tab = Hub.tab(Hub.active)
		if tab ~= nil and tab.hide ~= nil then
			tab.hide()
		end
		Hub.footer.enter:Hide()
		if Hub.afterHide ~= nil then
			Hub.afterHide()
		end
	end)
	-- Above whatever else sits on the DIALOG strata (the tour steps aside, spec 6.1).
	if type(frame.SetToplevel) == "function" then
		frame:SetToplevel(true)
	end
	-- Escape from any focus state, every other key passed on (Hub.onKey).
	if type(frame.EnableKeyboard) == "function" then
		frame:EnableKeyboard(true)
	end
	frame:SetScript("OnKeyDown", Hub.onKey)
	Theme.makeEscapable(Hub.FRAME_NAME)
	return frame
end

-- The header and the footer, as the active tab sets them ----------------------------------------

--- Draw `header` (Hub.headerFor): no crest without a class token, no note line without a note.
local function drawHeader(header)
	local parts = Hub.header
	if header.classToken ~= nil then
		Theme.paintCrest(parts.crest, parts.ring, header.classToken)
		parts.crest:Show()
		parts.ring:Show()
	else
		parts.crest:Hide()
		parts.ring:Hide()
	end
	parts.title:SetText(header.title or L.addonName)
	parts.noteLines = header.noteLines
	parts.note:SetText(header.note or "")
	Theme.setShown({ parts.note, parts.noteHover }, header.note ~= nil and header.note ~= "")
	return header
end

--- A tab's figures or waiting items changed: the header and the strip (its dots) drawn again while the
--- window shows. The header drawn, or nil while it is hidden.
function Hub.refresh()
	if not Hub.isShown() then
		return nil
	end
	drawStrip()
	return drawHeader(Hub.headerFor(Hub.tab(Hub.active)))
end

--- The footer's two hints: what Escape does (Escape closes, by default) and what Enter does (none: hidden).
function Hub.setHints(escape, enter)
	local footer = Hub.footer
	if footer == nil then
		return nil
	end
	footer.escape:SetText(escape or L.windowEscClose)
	if enter ~= nil and enter ~= "" then
		footer.enter:SetText(enter)
		footer.enter:Show()
	else
		footer.enter:Hide()
	end
	return escape, enter
end

--- The active tab's body is `body` px tall: the window takes the tallest body any available tab has
--- reported (spec 3.1: switching tabs never moves the footer; a shorter tab ends in ground), never under
--- Hub.heightFor's floor, and keeps its top edge where it was (Hub.place). The height.
function Hub.fitHeight(body)
	if Hub.frame == nil then
		return nil
	end
	Hub.bodies[Hub.active or ""] = body or 0
	local tallest = 0
	for _, tab in ipairs(Hub.visibleTabs()) do
		tallest = math.max(tallest, Hub.bodies[tab.id] or 0)
	end
	local height = Hub.heightFor(tallest)
	if height ~= Hub.height then
		Hub.height = height
		Hub.frame:SetHeight(height)
		Hub.place()
	end
	return height
end

-- Where it sits ----------------------------------------------------------------------------------

--- Where the window sits: the saved spot after a drag, else the middle of the screen. Taller than its
--- ordinary height, it keeps its top edge where it was: a centred frame moves down by half of what it grew.
function Hub.place()
	local frame = Hub.frame
	if frame == nil then
		return nil
	end
	frame:ClearAllPoints()
	local lower = ((Hub.height or Theme.SIZES.dialogHeight) - Theme.SIZES.dialogHeight) / 2
	local point = Prefs.get("hub", "point")
	if type(point) == "string" and point ~= "" then
		local y = Prefs.get("hub", "y")
		if point == "CENTER" and type(y) == "number" then
			y = y - lower
		end
		frame:SetPoint(point, UIParent, point, Prefs.get("hub", "x"), y)
		return point
	end
	frame:SetPoint("CENTER", UIParent, "CENTER", 0, -lower)
	return "CENTER"
end

-- Reset positions moves the window at once, open or not.
Prefs.onReset[#Prefs.onReset + 1] = function()
	Hub.place()
end

--- After a drag by the header: the window's centre against the screen's, at its ordinary height.
function Hub.savePosition()
	local frame = Hub.frame
	local x, y = frame:GetCenter()
	local parentX, parentY = UIParent:GetCenter()
	if x == nil or y == nil or parentX == nil or parentY == nil then
		return nil
	end
	Prefs.set("hub", "point", "CENTER")
	Prefs.set("hub", "x", x - parentX)
	Prefs.set("hub", "y", y - parentY + ((Hub.height or Theme.SIZES.dialogHeight) - Theme.SIZES.dialogHeight) / 2)
	Hub.place()
	return true
end

-- Tabs, opening and closing ----------------------------------------------------------------------

--- Show tab `id` (when it is available): the other tabs' bodies hidden and told so, the strip redrawn,
--- the choice remembered, the tab shown with `options`. The tab's id, or nil.
function Hub.select(id, options)
	local tab = Hub.tab(id)
	if Hub.frame == nil or tab == nil or not Hub.isAvailable(tab) then
		return nil
	end
	local previous = Hub.tab(Hub.active)
	if previous ~= nil and previous ~= tab and previous.hide ~= nil then
		previous.hide()
	end
	Hub.active = tab.id
	Prefs.set("hub", "tab", tab.id)
	Hub.markSeen(tab.id)
	for other, content in pairs(Hub.contents) do
		Theme.setShown({ content }, other == tab.id)
	end
	placeBodies()
	drawStrip()
	Hub.setHints(nil, nil)
	local emptyText = Hub.emptyText(tab)
	local empty = Hub.empties[tab.id]
	if emptyText ~= nil then
		empty:SetText(emptyText)
		empty:Show()
		Hub.fitHeight(type(empty.GetStringHeight) == "function" and empty:GetStringHeight() or 0)
	else
		empty:Hide()
		if tab.show ~= nil then
			tab.show(options or {})
		end
	end
	drawHeader(Hub.headerFor(tab))
	return tab.id
end

--- What a tab says while it has nothing to show (its `empty()`, e.g. "No raid logged yet."), or nil
--- when it has content; an error is nil (the tab draws as usual).
function Hub.emptyText(tab)
	if type(tab) ~= "table" or tab.empty == nil then
		return nil
	end
	local ok, text = pcall(tab.empty)
	return ok and type(text) == "string" and text ~= "" and text or nil
end

--- The window's frame key handler: Escape goes to the active tab first (a preview it closes, say), then
--- closes the window; every other key is passed on, so the player can walk with the window open.
--- SetPropagateKeyboardInput is a flag on our own frame, never protected, so it is allowed in combat.
function Hub.onKey(frame, key)
	local keep = key ~= "ESCAPE"
	if type(frame.SetPropagateKeyboardInput) == "function" then
		frame:SetPropagateKeyboardInput(keep)
	end
	if not keep then
		Hub.escape()
	end
end

--- Escape: the active tab's own step back when it has one, else the window closes.
function Hub.escape()
	local tab = Hub.tab(Hub.active)
	if tab ~= nil and tab.escape ~= nil and tab.escape() == true then
		return "tab"
	end
	Hub.close()
	return "closed"
end

--- Open the window on tab `id` (nil: Hub.landing's tab, else the tab the player last had, else the first
--- available), with `options` for that tab. The tab's id, or nil when no tab is available.
function Hub.open(id, options)
	if Hub.disabled then
		if Hub.fallback ~= nil then
			Hub.fallback()
		end
		return nil
	end
	local target = Hub.resolveTab(id or Hub.landing(), Prefs.get("hub", "tab"))
	if target == nil then
		return nil
	end
	if Hub.beforeOpen ~= nil then
		Hub.beforeOpen()
	end
	local frame = Hub.ensure()
	Hub.place()
	frame:Show()
	if type(frame.Raise) == "function" then
		frame:Raise()
	end
	return Hub.select(target, options)
end

--- Whether the window is on screen now.
function Hub.isShown()
	return Hub.frame ~= nil and Hub.frame:IsShown() == true
end

--- Hide the window (the active tab gives the keyboard back as it hides).
function Hub.close()
	if Hub.frame ~= nil and Hub.frame:IsShown() then
		Hub.frame:Hide()
	end
end

--- `/fs`, the key binding and the minimap's left click: close when open (on tab `id`, when one is
--- named), else open (no id: on Hub.landing's tab).
function Hub.toggle(id)
	if Hub.isShown() and (id == nil or id == Hub.active) then
		Hub.close()
		return nil
	end
	return Hub.open(id)
end

-- The tabs other lanes fill, kept off the strip until their lane registers the id again with content.
for _, id in ipairs({ "me", "gear", "rotation", "guild", "officer" }) do
	local key = "hubTab" .. id:sub(1, 1):upper() .. id:sub(2)
	Hub.register({ id = id, label = L[key], placeholder = true })
end

ns.Hub = Hub
return Hub
