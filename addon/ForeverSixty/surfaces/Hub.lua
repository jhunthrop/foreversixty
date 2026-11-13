-- addon/ForeverSixty/surfaces/Hub.lua
-- The Forever Sixty window (spec 2026-11-13-addon-builds-window section 3): the addon's one big frame,
-- a tabbed hub like the game's character and talent frames. The hub owns the frame and what every tab
-- shares: the header band (the class crest, a title, a line of figures), the tab strip, the footer
-- with its two key hints, where the window sits, its height, Escape and the remembered tab. Each tab
-- owns its body: it registers here (Hub.register) and builds into the content frame the hub hands it.
-- Builds (BuildDialog) is the first tab; Me, Gear, Rotation and Guild are placeholders, hidden until
-- a lane gives them content (`available`).
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Prefs = ns.Prefs or require("Prefs")

local Hub = {}

Hub.FRAME_NAME = "ForeverSixtyWindow"
--- Pixel layout: the header band and its crest, the tab strip (only while two tabs or more show), the
--- body's padding top and bottom, the footer band.
Hub.LAYOUT = { header = 56, crest = 32, crestGap = 10, strip = 28, bodyPad = 12, footer = 28, tabGap = 4 }
--- The footer's hints: the muted token at this alpha, a step dimmer than the text.
Hub.HINT_ALPHA = 0.75

--- Injection points: `fallback` runs instead when the window turned itself off this session (Options:
--- the command list); the tour fills `beforeOpen` and `afterHide` (it steps aside while the window is
--- open and comes back when it hides). `disabled` is set by a tab's render check (off screen).
Hub.fallback = nil
Hub.beforeOpen = nil
Hub.afterHide = nil
Hub.disabled = false

--- Every tab, in the order the strip shows them: { id, label, order, available() -> bool (shown on the
--- strip), build(content) (its body, once), show(options) (it comes on screen: draw, header, hints),
--- hide() (it leaves), escape() -> true when the tab stepped back itself, empty() -> a muted line
--- while it has nothing to show }. Filled by Hub.register at load.
Hub.tabs = {}

--- The window's width: the Builds tab's two columns, the gutter between them, the padding.
function Hub.width()
	local S = Theme.SIZES
	return S.padding * 2 + S.windowLeft + S.windowGutter + S.windowRight
end

-- The bodies' placement, defined with the frame below; a late registration reaches it at call time.
local placeBodies

--- Add or replace a tab (by `id`). A tab with no `available` is always shown; one whose `available`
--- answers false is left off the strip and cannot be opened. Returns the tab.
function Hub.register(tab)
	assert(type(tab) == "table" and type(tab.id) == "string", "Hub.register needs a tab with an id")
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

--- Whether `tab` shows: no `available`, or one that answers true (an error is "not available").
function Hub.isAvailable(tab)
	if type(tab) ~= "table" then
		return false
	end
	if tab.available == nil then
		return true
	end
	local ok, answer = pcall(tab.available)
	return ok and answer == true
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

local function line(parent, hexKey, width)
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
	if type(region.SetWordWrap) == "function" then
		region:SetWordWrap(false)
	end
	return region
end

local function rule(parent)
	local texture = Theme.texture(parent, "ARTWORK", "border")
	texture:SetHeight(1)
	return texture
end

--- An invisible frame over the header's note: what the active tab hands for its hover (the Builds tab:
--- the off-plan talents by name, as the panel's).
local function buildNoteHover(header)
	local over = Theme.createFrame("Frame", nil, header.bar)
	over:EnableMouse(true)
	over:SetAllPoints(header.note)
	over:SetScript("OnEnter", function(self)
		local lines = header.noteLines
		if type(lines) == "table" and #lines > 0 then
			Theme.showLines(self, lines, "ANCHOR_RIGHT")
		end
	end)
	over:SetScript("OnLeave", function()
		Theme.hideTooltip()
	end)
	return over
end

--- The header band: the ringed class crest, the title in the title face over the figures line, the
--- close X; the band drags the window; a hairline under it.
local function buildHeader(frame)
	local S, G = Theme.SIZES, Hub.LAYOUT
	local bar = Theme.createFrame("Frame", nil, frame)
	bar:SetPoint("TOPLEFT", frame, "TOPLEFT", 0, 0)
	bar:SetPoint("TOPRIGHT", frame, "TOPRIGHT", 0, 0)
	bar:SetHeight(G.header)
	bar:EnableMouse(true)
	bar:RegisterForDrag("LeftButton")
	bar:SetScript("OnDragStart", function()
		frame:StartMoving()
	end)
	bar:SetScript("OnDragStop", function()
		frame:StopMovingOrSizing()
		Hub.savePosition()
	end)
	local header = { bar = bar }
	header.crest, header.ring = Theme.buildCrest(bar, G.crest)
	header.crest:SetPoint("LEFT", bar, "LEFT", S.padding, 0)
	local textWidth = Hub.width() - S.padding * 2 - G.crest - G.crestGap - S.dialogClose - S.gap * 4
	header.title = Theme.titleFont(line(bar, "bone", textWidth), S.windowTitle)
	header.title:SetText(L.addonName)
	header.title:SetPoint("BOTTOMLEFT", header.crest, "RIGHT", G.crestGap, 1)
	header.note = line(bar, "muted")
	header.note:SetPoint("TOPLEFT", header.crest, "RIGHT", G.crestGap, -1)
	header.noteHover = buildNoteHover(header)
	local close = Widgets.closeButton(bar, function()
		Hub.close()
	end)
	close:SetSize(S.dialogClose, S.dialogClose)
	close:SetPoint("RIGHT", bar, "RIGHT", -S.gap * 3, 0)
	header.close = close
	local under = rule(bar)
	under:SetPoint("BOTTOMLEFT", bar, "BOTTOMLEFT", 0, 0)
	under:SetPoint("BOTTOMRIGHT", bar, "BOTTOMRIGHT", 0, 0)
	return header
end

--- The footer band: a hairline over it, Escape's hint at the left and Enter's at the right.
local function buildFooter(frame)
	local S, G = Theme.SIZES, Hub.LAYOUT
	local footer = {}
	local over = rule(frame)
	over:SetPoint("BOTTOMLEFT", frame, "BOTTOMLEFT", 0, G.footer)
	over:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", 0, G.footer)
	local hex = Theme.dimHex("muted", Hub.HINT_ALPHA)
	footer.escape = line(frame, "muted")
	footer.escape:SetTextColor(Theme.rgb(hex))
	footer.escape:SetPoint("LEFT", frame, "BOTTOMLEFT", S.padding, G.footer / 2)
	footer.escape:SetText(L.windowEscClose)
	footer.enter = line(frame, "muted")
	footer.enter:SetTextColor(Theme.rgb(hex))
	footer.enter:SetJustifyH("RIGHT")
	footer.enter:SetPoint("RIGHT", frame, "BOTTOMRIGHT", -S.padding, G.footer / 2)
	footer.enter:Hide()
	return footer
end

--- The tab strip under the header: one tab per registered tab (Widgets.tab, the gold underline on
--- the active one), shown only while two or more are available.
local function buildTabButton(strip, tab)
	local S, G = Theme.SIZES, Hub.LAYOUT
	local button = Widgets.tab(strip.frame, tab.label, function()
		Hub.select(tab.id)
	end)
	button:SetHeight(G.strip)
	local width = type(button.foreverSixtyLabel.GetStringWidth) == "function"
		and button.foreverSixtyLabel:GetStringWidth() or 0
	button:SetWidth(math.max(S.tabWidth, (width or 0) + S.padding * 2))
	strip.buttons[tab.id] = button
	return button
end

local function buildStrip(frame)
	local G = Hub.LAYOUT
	local strip = { buttons = {} }
	strip.frame = Theme.createFrame("Frame", nil, frame)
	strip.frame:SetPoint("TOPLEFT", frame, "TOPLEFT", 0, -G.header)
	strip.frame:SetPoint("TOPRIGHT", frame, "TOPRIGHT", 0, -G.header)
	strip.frame:SetHeight(G.strip)
	local under = rule(strip.frame)
	under:SetPoint("BOTTOMLEFT", strip.frame, "BOTTOMLEFT", 0, 0)
	under:SetPoint("BOTTOMRIGHT", strip.frame, "BOTTOMRIGHT", 0, 0)
	return strip
end

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
	buildTabButton(Hub.strip, tab)
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
	Hub.header = buildHeader(frame)
	Hub.strip = buildStrip(frame)
	Hub.footer = buildFooter(frame)
	Hub.contents, Hub.empties = {}, {}
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

--- The header for the active tab: { classToken, title, note, noteLines } (no crest without a class
--- token; the title "Forever Sixty" when the tab names none; no note line when nil).
function Hub.setHeader(header)
	local parts = Hub.header
	if parts == nil then
		return nil
	end
	header = header or {}
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

--- Grow or shrink the window to hold `body` px of the active tab's content (Hub.heightFor); it keeps
--- its top edge where it was (Hub.place). The height.
function Hub.fitHeight(body)
	if Hub.frame == nil then
		return nil
	end
	local height = Hub.heightFor(body)
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
	local point = Prefs.get("dialog", "point")
	if type(point) == "string" and point ~= "" then
		local y = Prefs.get("dialog", "y")
		if point == "CENTER" and type(y) == "number" then
			y = y - lower
		end
		frame:SetPoint(point, UIParent, point, Prefs.get("dialog", "x"), y)
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
	Prefs.set("dialog", "point", "CENTER")
	Prefs.set("dialog", "x", x - parentX)
	Prefs.set("dialog", "y", y - parentY + ((Hub.height or Theme.SIZES.dialogHeight) - Theme.SIZES.dialogHeight) / 2)
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
	Prefs.set("dialog", "tab", tab.id)
	for other, content in pairs(Hub.contents) do
		Theme.setShown({ content }, other == tab.id)
	end
	placeBodies()
	drawStrip()
	Hub.setHints(nil, nil)
	Hub.setHeader(nil)
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

--- Open the window on tab `id` (nil: the tab the player last had, else the first available), with
--- `options` for that tab. The tab's id, or nil when no tab is available.
function Hub.open(id, options)
	if Hub.disabled then
		if Hub.fallback ~= nil then
			Hub.fallback()
		end
		return nil
	end
	local target = Hub.resolveTab(id, Prefs.get("dialog", "tab"))
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

--- The key binding and the minimap's left click: close when open, else open on the remembered tab.
function Hub.toggle(id)
	if Hub.isShown() and (id == nil or id == Hub.active) then
		Hub.close()
		return nil
	end
	return Hub.open(id)
end

-- The tabs other lanes fill: each is hidden until its lane gives it content (`available`).
for index, placeholder in ipairs({
	{ id = "me", label = L.hubTabMe },
	{ id = "gear", label = L.hubTabGear },
	{ id = "rotation", label = L.hubTabRotation },
	{ id = "guild", label = L.hubTabGuild },
}) do
	placeholder.order = 10 + index
	placeholder.available = function()
		return false
	end
	Hub.register(placeholder)
end

ns.Hub = Hub
return Hub
