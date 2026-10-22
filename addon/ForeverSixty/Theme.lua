-- addon/ForeverSixty/Theme.lua
-- The look, the measurements, and the single place that asks the 1.60
-- client what it has.
--
-- Nothing else in this addon names a frame template, a font object, a
-- client UI global or a client capability. That rule is what makes the
-- unknowns knowable: the spike table in README.md has one row per question
-- asked here, every question has a plain-texture or plain-function answer
-- for "no", and a client that answers "no" to all of them still gets a
-- working addon that merely looks plainer.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")

local Theme = {}

--- The site's palette, web/src/styles/tokens.css, as six hex digits.
Theme.HEX = {
	background = "0d111a",
	--- A field the player reads or types in: darker than the window, so it
	--- reads as set into it.
	inset = "070a10",
	--- A button at rest, and under the cursor.
	raised = "161c2b",
	hover = "1f2739",
	--- A card on a page, and the sidebar's ground.
	card = "111726",
	sidebar = "0a0e16",
	--- A progress bar's empty track.
	track = "1b2233",
	--- The primary button's label: dark, on gold.
	onGold = "15110a",
	goldHover = "f2cd74",
	success = "6fcf8e",
	shadow = "000000",
	border = "262e40",
	titleTop = "131824",
	titleBottom = "0d111a",
	gold = "e5b955",
	muted = "9a9484",
	--- Muted text on a row drawn at ALPHA.dim (an offline guildmate's "no logs", L5.1 fix
	--- round 2): `muted` reads 2.7:1 there; this reaches 3.2:1 and stays below `body`.
	mutedDimmed = "b0aa9c",
	--- Muted text on the client's own tooltip, which sits on a lighter ground than
	--- the window: at least 0.72 in every channel, lighter than the client's own grey.
	tooltipMuted = "b8b8b8",
	--- Text on the client's tooltip that must read clearly ("Your sim").
	tooltipText = "d0d0d0",
	--- The client's own tooltip body white: the addon's plain lines on a player's tooltip (A.2).
	tooltipBody = "ffffff",
	--- The client's own tooltip gold (its item-set and "Requires" lines).
	clientGold = "ffd100",
	body = "e9e4d8",
	warning = "ff6b5c",
	--- A scroll track one step lighter than `border`, so the guild tab's track reads on the
	--- panel (L5.1 fix round 1, item 7).
	scrollTrack = "323b52",
	--- The score tiers' colours lifted for the dark panel (L5.1 fix round 1, item 4): the
	--- client's rare blue (0070dd), epic purple (a335ee) and legendary orange (ff8000) read
	--- under 4.5:1 on it, or under 3:1 on an offline row. Theme.PANEL_TIER maps a tier to these.
	rareBright = "6eaeec",
	epicBright = "cd92f5",
	legendaryBright = "ff8f1f",
}

Theme.ALPHA = {
	window = 0.96,
	--- The chip's hover hints: one step dimmer than the muted info line.
	hint = 0.6,
	--- The idle chip's edge, in the muted colour, so it reads on a dark scene.
	idleEdge = 0.4,
	--- A row the player has already matched, and a disabled button.
	dim = 0.55,
	disabled = 0.4,
	shadow = 0.5,
	--- The talent window's strip ground (spec section 4.2).
	strip = 0.9,
	--- A talent badge's dark plate (near-opaque, L3.2 fix round 4: a bright
	--- icon washed a 70 % plate's digit out), a badge's edge beyond the
	--- points spendable now, and the talent ring's dark outer edge.
	plate = 0.9,
	unreachable = 0.5,
	ringEdge = 0.6,
	--- The 1 px black shadow behind the slot marks' gold corner (L4.1 fix round 5).
	markShadow = 0.6,
	--- The least a badge's digits dim beyond the reachable points: the
	--- plate dims, the number stays legible (L3.2 fix round 3).
	dimDigits = 0.6,
}

Theme.SIZES = {
	--- The paper-doll and bag marks (spec section 4.4, L4.1): the corner arrow's side, and
	--- the target corner's two gold legs, each `mark` long and `markEdge` thick.
	mark = 12,
	markEdge = 2,
	--- The target corner sits this far in from the button's bottom right corner, clear of the
	--- client's quality border (L4.1 fix round 5).
	markInset = 2,
	--- The quest reward label and the roll line (L4.2): the dark plate's padding round
	--- the text, the roll line's gap from the roll frame's right edge, and the widest the
	--- roll line grows before the verdict puts the compared name on a second line.
	labelPad = 2,
	rollGap = 6,
	rollLineWidth = 240,
	windowWidth = 720,
	windowHeight = 500,
	--- The left navigation, and one item in it.
	sidebarWidth = 150,
	navHeight = 30,
	navIcon = 16,
	navBar = 2,
	--- How far the drop shadow reaches past the window.
	shadow = 5,
	cardGap = 12,
	cardHeight = 150,
	progressHeight = 6,
	pillHeight = 16,
	--- Seconds the window takes to fade in.
	fadeIn = 0.15,
	titleBarHeight = 28,
	--- The two lines under the title bar: who this is, and which data build.
	headerHeight = 56,
	tabHeight = 24,
	tabUnderline = 2,
	closeButton = 20,
	tabWidth = 96,
	border = 1,
	padding = 16,
	gap = 4,
	rowHeight = 20,
	--- Rows per list, chosen so each page fits the window under the tab
	--- strip: the Follow list sits between the paste field and its buttons,
	--- and the Gear page stacks two lists. The lists scroll with the wheel.
	followRows = 11,
	gearSlotRows = 9,
	gearUpgradeRows = 5,
	equipButtonWidth = 72,
	buttonHeight = 22,
	buttonWidth = 150,
	editBoxHeight = 72,
	iconSize = 16,
	--- The chip (spec section 4.1): square when idle, growing leftwards to
	--- fit a line of 12 px text, up to `chipWide`. The 24 px icon sits
	--- `chipIconInset` in from the edge; the text starts `chipTextGap` past
	--- the icon and keeps the same gap on the right, so a line of width W
	--- makes a chip W + `chipTextInset` wide.
	chip = 28,
	chipWide = 280,
	chipText = 12,
	chipIcon = 24,
	chipIconInset = 2,
	chipTextGap = 6,
	chipTextInset = 38,
	--- The point state's ring, on its own raised frame round the icon: its
	--- outer edge this far outside the icon, this thick. 4 out from an icon
	--- inset 2 puts the 2 px ring just outside the chip's own 1 px edge, so
	--- the chip's border never reads as gold.
	chipRingInset = 4,
	chipRingThickness = 2,
	--- How far left of the minimap's own left edge the chip sits by default.
	chipMinimapGap = 12,
	--- The "+2" badge straddling the chip's top-right corner, and its digits.
	chipBadge = 14,
	chipBadgeText = 10,
	--- The gap under the chip before its "pick a path" card, the card's own
	--- width and inner padding, and the space between one spec and the next.
	--- The card's playstyle, tag and footer lines are `chipCardNote` px, and
	--- they and the button labels start `chipCardLabelInset` in from the button's edge.
	chipCardGap = 4,
	chipCardWidth = 240,
	chipCardNote = 11,
	chipCardLabelInset = 8,
	chipCardPadding = 12,
	chipCardRowGap = 8,
	minimapButton = 32,
	minimapIcon = 20,
	minimapRadius = 80,
	glowThickness = 3,
	--- How far outside a talent button its glow ring sits, clear of the
	--- game's own rank border.
	glowInset = 3,
	--- The talent ring's 1 px dark edge just outside its gold, so it reads
	--- on the bright talent art (L3.2 fix round 2).
	glowEdge = 1,
	--- The talent marks (spec section 4.2): the order number on a planned
	--- node, a plate 16 px high with 12 px digits, at least as wide as it is
	--- high and `orderBadgePad` wider than its digits, its top-right corner
	--- `orderBadgeOffset` up and right of the node's own; the strip, a pill
	--- in the talent window's bottom-left corner `talentStripInset` inside
	--- the window's border art (`talentWindowBorder` thick, seen in game:
	--- an 8 px offset sat on the border), its text `talentStripText` px and
	--- `talentStripPadding` in from each side, its right edge at least
	--- `talentStripGap` short of the window's Apply Changes button.
	orderBadge = 16,
	orderBadgeText = 12,
	orderBadgeOffset = 4,
	orderBadgePad = 6,
	talentStrip = 26,
	-- A.2: one size up from 12, so the strip reads at the window's distance.
	talentStripText = 13,
	talentWindowBorder = 8,
	talentStripInset = 8,
	talentStripPadding = 12,
	talentStripGap = 8,
	--- How long "Selected -- press Ctrl+C" stays, and how long "Build
	--- complete" stays on the chip before it goes idle.
	copiedSeconds = 2,
	completeSeconds = 5,
	--- The rotation card's own row height: an icon-height line for the
	--- ability's name and rank, a small-font line under it for the
	--- condition, and the gap between them.
	rotationRowHeight = 34,
	--- The rotation card's header strip icon.
	rotationHeaderIcon = 14,
	--- The character header's own circular ringed class crest (design
	--- section 2 / section 4.5.1, round-2 owner ruling), and the gap
	--- between it and the name beside it.
	headerCrest = 36,
	headerCrestGap = 10,
	--- The Overview's "Send to the site" card (round-2 owner fix): the
	--- same crest, smaller, beside the character's own name and level.
	overviewSyncCrest = 28,
	--- Gear page (design section 4.5.4, round-2/3 fixes): the slot-label
	--- column, the Planned/Equipped name columns (equal widths, sized to
	--- clear the longest band-20 item name without truncating), the
	--- green tick glyph that replaces the old "As planned" text tag, and
	--- the gap between the three columns.
	gearSlotLabelWidth = 50,
	gearNameWidth = 232,
	gearPlannedTick = 14,
	gearRowGap = 12,
	--- The one note that is still prose, not a glyph or a short tag: an
	--- unplanned-but-equipped slot's "No plan for this slot".
	gearNoPlanWidth = 96,
	--- The gold "+12" tag on a slot where what is worn beats the plan.
	gearBetterWidth = 40,
	--- The thin scrollbar beside a list or page that holds more than fits.
	scrollbar = 4,
	scrollbarThumb = 16,
	--- A tree's full width on the Overview's talent-points bars: the points
	--- that reach a tree's last row, so seven points never read as a tree
	--- completed.
	treePoints = 31,
	--- The minimap button's ring texture and how far past the minimap's
	--- own edge the button's centre sits.
	minimapRing = 54,
	minimapGround = 24,
	minimapEdge = 10,
	--- The build dialog (spec section 4.10): 420 by 300; every line that says something
	--- at `dialogText` px; the close button's hit area (boxes are sized in measured whole lines,
	--- BuildDialog.boxLines); the preview's class icon; the most "from the site" rows it
	--- lists; the gear rows' height.
	dialogWidth = 420,
	dialogHeight = 300,
	dialogText = 12,
	dialogClose = 20,
	dialogLine = 24,
	dialogClassIcon = 32,
	dialogInboxRows = 2,
	--- The gear heading's line (the rows themselves are BuildDialog.GEAR_PITCH, rowHeight).
	dialogGearRow = 14,
	--- The gear rows' fixed slot column, before the icon and the name.
	dialogSlotColumn = 72,
	--- The Settings page (spec section 4.11, L3.5): a 44 px row per switch so the whole row
	--- is a target, the least a row is; each hint wrapped to the canvas width minus its 40 px
	--- indent (and the right padding); a muted line's pitch; the title's line; the width a
	--- canvas the client has not sized yet is laid out at.
	settingsRow = 44,
	--- A row is measured from its label and hint (fix round 2, item 3): the gap above the
	--- label, the label, the gap, the hint's GetStringHeight, then this pad, the one-line
	--- rows' bottom pad (the client's 12 px label over a 10 px hint makes a 44 px row), so a
	--- wrapped hint pushes the next row down by its extra line.
	settingsRowPad = 14,
	--- Under the foot's last line, so a cut at the bottom of the scroll never lands on text
	--- (fix round 2, item 5).
	settingsFootPad = 12,
	settingsHintInset = 40,
	settingsLine = 18,
	settingsTitle = 24,
	settingsWidth = 600,
	settingsHeight = 560,
	--- How far inside the canvas the page's own opaque backdrop sits (fix round 1, item 9).
	settingsInset = 8,
	--- The character-panel tab (spec section 4.5, L4.3): the "Forever Sixty" button beside the
	--- client's tabs (its size, its gap from them) and the flyout at the pane's top right (its
	--- width, where it sits from the pane's corner, its padding, the gap between rows and
	--- between regions, a line of its small text, a region's title row, the header line, its
	--- buttons' height). Every row is a fixed height from its lines (fix round 1, item 1: the
	--- live client drew the icons ~40 px over 14 px rows): `rowIcon` square at the row's top
	--- left, the text `gap` right of it; a text line is `panelLine` high.
	panelTabWidth = 112,
	panelTabHeight = 24,
	panelTabGap = 0,
	flyoutWidth = 320,
	panelOffsetX = 0,
	panelOffsetY = 0,
	panelPadding = 10,
	panelRowGap = 4,
	panelRegionGap = 10,
	panelLine = 12,
	panelTitle = 16,
	panelHeader = 18,
	rowIcon = 16,
	panelButtonHeight = 22,
	--- The room a target row keeps at its right for the gain ("+12.3 DPS") beside the name.
	panelGainWidth = 64,
	--- The guild tab (spec section 4.7, L5.1): its flyout's width (its button is the panel
	--- tab's), one roster row, the column header row and its sort arrow, the gap between two
	--- columns, and the "Online only" row above the headers. The officer region's push
	--- button is `guildPushWidth` wide; the push confirm frame `guildConfirmWidth`.
	guildFlyoutWidth = 480,
	guildRow = 16,
	guildHeaderRow = 16,
	guildSortArrow = 8,
	guildColumnGap = 4,
	guildToolbar = 20,
	guildPushWidth = 200,
	guildConfirmWidth = 300,
	--- L5.1 fix round 1: the room every column header keeps after its label for the sort
	--- arrow (items 3 and 5; a right-aligned column's cells end where its label does), the
	--- gap the flyout keeps from the guild frame's side tabs (item 7), the guild list's scroll
	--- thumb and track width (item 7), and the room the "Online only" switch takes at the left
	--- of the count line.
	guildSortReserve = 10,
	guildSideClear = 4,
	guildScrollbar = 8,
	--- The guild tab's scroll thumb (L3.7 item 41): wider than its track, centred on it, so it
	--- reads as something to grab; the track (and the columns' room) stay as they were.
	guildScrollThumb = 10,
	guildToggleRoom = 88,
	--- The group panel (spec section 4.8, L5.2): its flyout's width (its button is the panel
	--- tab's; its headers, columns and scroll track are the guild tab's), the most member rows
	--- shown at once (each a cell line and the muted weakest-slot line under it), the ready
	--- check's missing lines before "and N more" (fix round 4, C: the coverage lines' too), and the
	--- enchant cell's tick. Fix round 4 (D): where the Social window's side tabs cannot be read, the
	--- flyout sits `groupSideFallback` right of the window's edge (past them). Fix round 5 (D1): a
	--- child of the window wider than `groupSideTabMaxWidth` is no side tab (the client's are 55),
	--- and the side tabs' reach is clamped to it.
	groupFlyoutWidth = 480,
	groupVisibleRows = 10,
	groupMissingLines = 6,
	groupTick = 12,
	groupSideFallback = 55,
	groupSideTabMaxWidth = 100,
	--- The loot council helper's Loot region in that flyout (L5.3): the columns of a member's row (the
	--- name, the value with its tag, the worn item, the muted "planned" tag; `lootColumnGap` between)
	--- and the most rows a drop lists before "and N more".
	lootNameWidth = 84,
	lootValueWidth = 130,
	lootPlannedWidth = 76,
	lootColumnGap = 6,
	lootMaxRows = 8,
}

--- The client's own font objects, and what to use when one is missing.
Theme.FONTS = {
	normal = "GameFontNormal",
	small = "GameFontNormalSmall",
	highlight = "GameFontHighlight",
	large = "GameFontNormalLarge",
}
Theme.FALLBACK_FONT = { path = "Fonts\\FRIZQT__.TTF", size = 12 }

--- Every template this addon will ever ask for, by the key callers use.
--- Empty save one entry, on purpose. The first in-game screenshots showed
--- what the client's stock templates do to this window: a red action
--- button, gold tabs hanging off the bottom edge and a one-line input
--- with the export spilling out of it, none of it the site's design.
--- Buttons, tabs and fields are drawn by Widgets from flat textures
--- instead, which also means one code path to test. A real tooltip is
--- the one widget that genuinely cannot be flat-textured -- SetHyperlink,
--- item comparison and the rest are the GameTooltip frame type's own C
--- behaviour, not something a template merely skins -- so the BiS
--- hover's second tooltip (Theme.compareTooltip) is the one caller that
--- names a real one, gracefully bare on a client without it exactly like
--- every other Theme.createFrame call. createFrame keeps the lookup so a
--- template can be named here again without touching a caller.
Theme.TEMPLATES = { gameTooltip = "GameTooltipTemplate" }

Theme.MEDIA = {
	minimapIcon = "Interface\\AddOns\\ForeverSixty\\media\\minimap",
	--- The character header's own class-colour ring (design section 2's
	--- own crest bullet): one plain white ring, class-coloured at
	--- runtime through SetVertexColor (Window.applyHeader), rather than
	--- nine near-identical ring files.
	crestRing = "Interface\\AddOns\\ForeverSixty\\media\\crests\\ring",
	--- The Gear page's own green tick glyph (round-3 fix, section 4.5.4):
	--- the client's own ready-check icon, already green, so this addon
	--- ships no new art for it.
	gearTick = "Interface\\RaidFrame\\ReadyCheck-Ready",
	--- Its red twin, for a slot worn differently from the plan.
	gearCross = "Interface\\RaidFrame\\ReadyCheck-NotReady",
	--- The minimap button's round crop and the client's own ring for it,
	--- the same two textures every LibDBIcon button on the minimap uses.
	minimapMask = "Interface\\CharacterFrame\\TempPortraitAlphaMask",
	minimapRing = "Interface\\Minimap\\MiniMap-TrackingBorder",
	minimapGround = "Interface\\Minimap\\UI-Minimap-Background",
	--- The client's own round class icons (the target frame's), cut by
	--- CLASS_ICON_TCOORDS: readable at 24 px where the crest art is not.
	classIcons = "Interface\\TargetingFrame\\UI-Classes-Circles",
	--- The client's own check mark, tinted dark on a switch's gold square (L3.5 ruling 2).
	check = "Interface\\Buttons\\UI-CheckBox-Check",
	--- The paper-doll and bag marks' corner arrow (spec section 4.4, L4.1) is drawn, the
	--- addon's own parts (SlotMarks.ARROW_PARTS). The client's own bag upgrade arrow, used
	--- instead only with Theme.ARROW_FROM_ATLAS on and where the client knows the atlas
	--- (C_Texture.GetAtlasInfo).
	arrowAtlas = "bags-greenarrow",
	--- The guild tab's column sort arrow (L5.1): the client's own, the one its Who and
	--- classic guild roster headers draw, pointing down for high to low (Theme.SORT_COORDS).
	sortArrow = "Interface\\Buttons\\UI-SortArrow",
}

--- The sort arrow's texture coordinates: the art's left 9 of 16 pixels, upright for high to
--- low (down), flipped for low to high (up), as the client's own headers cut it.
Theme.SORT_COORDS = { descending = { 0, 0.5625, 0, 1 }, ascending = { 0, 0.5625, 1, 0 } }

--- Whether the slot marks' arrow is the client's `bags-greenarrow` atlas (where the client
--- knows it) rather than the addon's own 12 px glyph (SlotMarks.ARROW_PARTS). Off: on the live
--- 1.60 client the atlas art drew a ~6 px dot in the 12 px frame (L4.1 fix round 6).
Theme.ARROW_FROM_ATLAS = false

--- The nine classes this addon ships a crest for (addon/tools/
--- make_class_crests.py), keyed by UnitClass's own second return
--- ("WARRIOR", not "warrior") -- the same token Theme.classColor and
--- Theme.SLOT_ICONS' neighbours already key by.
Theme.CREST_CLASSES = {
	WARRIOR = true, PALADIN = true, HUNTER = true, ROGUE = true, PRIEST = true,
	SHAMAN = true, MAGE = true, WARLOCK = true, DRUID = true,
}

--- The character's own crest texture for a class token, or nil for a
--- token this addon ships no crest for (a class Forever adds later, or
--- a test double with none) -- Theme.paintCrest falls back to the
--- neutral tile rather than a blank space on nil, same as every other
--- icon in this addon. The crest ships pre-cropped to a circular disc
--- (make_class_crests.py's own check of the source webp's corner alpha),
--- so unlike Theme.icon this never needs a runtime circular-crop
--- capability (MaskTexture) at all -- there is no square card left to
--- crop once the file itself is round.
function Theme.crestPath(token)
	if type(token) ~= "string" or not Theme.CREST_CLASSES[token] then
		return nil
	end
	return "Interface\\AddOns\\ForeverSixty\\media\\crests\\" .. token:lower()
end

--- The class's icon and its texture coords: the client's own round class
--- icon (Theme.MEDIA.classIcons, CLASS_ICON_TCOORDS) where the client has
--- the table, else the shipped crest, else the neutral tile.
function Theme.classIcon(token)
	local coords = type(CLASS_ICON_TCOORDS) == "table" and token ~= nil and CLASS_ICON_TCOORDS[token] or nil
	if type(coords) == "table" and #coords == 4 then
		return Theme.MEDIA.classIcons, { coords[1], coords[2], coords[3], coords[4] }
	end
	local crest = Theme.crestPath(token)
	if crest ~= nil then
		return crest, { 0, 1, 0, 1 }
	end
	return Theme.UNKNOWN_ICON, { Theme.ICON_CROP[1], Theme.ICON_CROP[2], Theme.ICON_CROP[3], Theme.ICON_CROP[4] }
end

--- A crest + ring texture pair, `size` square: the character header's
--- own 36px crest (Window.lua) and the Overview's sync card's smaller
--- one (OverviewView.lua) share this one builder rather than each
--- repeating the CreateTexture/SetPoint calls. Plain textures, not
--- Theme.icon: the crest ships pre-cropped to a circular disc, so
--- Theme.icon's own bevel-trim texture coords -- meant for square
--- inventory icons -- would cut into the circle rather than clean it up.
function Theme.buildCrest(parent, size)
	local crest = parent:CreateTexture(nil, "ARTWORK")
	crest:SetSize(size, size)
	local ring = parent:CreateTexture(nil, "OVERLAY")
	ring:SetAllPoints(crest)
	ring:SetTexture(Theme.MEDIA.crestRing)
	return crest, ring
end

--- Paints a Theme.buildCrest pair for `classToken`: the crest's own
--- texture (the neutral tile for a class this addon ships none for --
--- never a blank space) and the ring's class-colour tint (gold,
--- Theme.classColor's own fallback).
function Theme.paintCrest(crest, ring, classToken)
	crest:SetTexture(Theme.crestPath(classToken) or Theme.UNKNOWN_ICON)
	ring:SetVertexColor(Theme.classColor(classToken))
end

--- The four edges of a rectangle, for outline().
Theme.EDGES = {
	{ from = "TOPLEFT", to = "TOPRIGHT", horizontal = true },
	{ from = "BOTTOMLEFT", to = "BOTTOMRIGHT", horizontal = true },
	{ from = "TOPLEFT", to = "BOTTOMLEFT", horizontal = false },
	{ from = "TOPRIGHT", to = "BOTTOMRIGHT", horizontal = false },
}

--- What is not working on this client (Theme.note), and what was merely
--- found out about it (Theme.noteInfo). Read by /fs diag.
ns.Diagnostics = ns.Diagnostics or {}
ns.DiagnosticsInfo = ns.DiagnosticsInfo or {}
ns.DiagnosticsVerbose = ns.DiagnosticsVerbose or {}
Theme.templates = {}

function Theme.diagnostics()
	return ns.Diagnostics
end

function Theme.note(message)
	ns.Diagnostics[#ns.Diagnostics + 1] = message
	return ns.Diagnostics
end

--- A finding that is not a fault (which talent window answered, say): /fs
--- diag prints it after the report, never under "Not working:".
function Theme.noteInfo(message)
	ns.DiagnosticsInfo[#ns.DiagnosticsInfo + 1] = message
	return ns.DiagnosticsInfo
end

function Theme.infoNotes()
	return ns.DiagnosticsInfo
end

--- A finding only a tester wants (where a frame landed, which mapping
--- answered): /fs diag verbose prints it, plain /fs diag does not.
function Theme.noteVerbose(message)
	ns.DiagnosticsVerbose[#ns.DiagnosticsVerbose + 1] = message
	return ns.DiagnosticsVerbose
end

function Theme.verboseNotes()
	return ns.DiagnosticsVerbose
end

--- The addon's own frames, for /fs diag verbose (L9.2 B3): one entry per label and reason with how
--- many were made and when the first was (seconds after the first one). WoW cannot destroy a
--- frame, so what is made stays; this says what made it and why. A label made again only counts.
ns.FrameLog = ns.FrameLog or {}
local firstFrameAt = nil

function Theme.noteFrame(label, reason)
	local now = Theme.now()
	firstFrameAt = firstFrameAt or now
	for _, entry in ipairs(ns.FrameLog) do
		if entry.label == label and entry.reason == reason then
			entry.count = entry.count + 1
			return entry
		end
	end
	local entry = { label = label, reason = reason, count = 1, at = now - firstFrameAt }
	ns.FrameLog[#ns.FrameLog + 1] = entry
	return entry
end

function Theme.frameNotes()
	return ns.FrameLog
end

--- The frame log as /fs diag verbose prints it: a heading, then a line per label and reason.
function Theme.frameLines()
	local lines = {}
	if #ns.FrameLog == 0 then
		return lines
	end
	lines[1] = L.diagFramesHeading
	for _, entry in ipairs(ns.FrameLog) do
		lines[#lines + 1] = string.format(L.diagFrameLine, entry.label, entry.count > 1 and (" x" .. entry.count) or "",
			entry.reason, entry.at)
	end
	return lines
end

--- What the addon holds, in KB: the client's own count for it where it gives one, else Lua's.
function Theme.memoryKB()
	if type(GetAddOnMemoryUsage) == "function" then
		if type(UpdateAddOnMemoryUsage) == "function" then
			pcall(UpdateAddOnMemoryUsage)
		end
		local ok, kb = pcall(GetAddOnMemoryUsage, "ForeverSixty")
		if ok and type(kb) == "number" then
			return math.floor(kb + 0.5)
		end
	end
	return math.floor(collectgarbage("count") + 0.5)
end

--- The /fs diag lines: "Not working:" and only the broken things, or
--- "Everything is working."; then the surfaces' own lines (`chipLine`, a
--- line or a list of them, when given) and the findings, outside the heading.
function Theme.diagReport(broken, chipLine, info, headline)
	local lines = {}
	if #broken > 0 then
		lines[1] = L.diagHeading
		for _, note in ipairs(broken) do
			lines[#lines + 1] = note
		end
	else
		-- X.3: the green sentence only when the caller says nothing is missing; else its own line.
		lines[1] = headline or L.diagAllWorking
	end
	local status = {}
	if type(chipLine) == "table" then
		for _, line in ipairs(chipLine) do
			status[#status + 1] = line
		end
	elseif chipLine ~= nil then
		status[1] = chipLine
	end
	-- The surfaces' status lines ("Build:", "Chip:", "Talent badges:", "Bag marks:") never
	-- read as faults: under "Not working:" they get their own heading (L4.1 fix round 5).
	if #broken > 0 and #status > 0 then
		lines[#lines + 1] = L.diagWorkingHeading
	end
	for _, line in ipairs(status) do
		lines[#lines + 1] = line
	end
	for _, note in ipairs(info or {}) do
		lines[#lines + 1] = note
	end
	return lines
end

--- Forget what was learned about the client. Specs only; the game never
--- changes its capabilities inside a session.
function Theme.reset()
	Theme.templates = {}
	Theme.glowPair = nil
	Theme.addonMessagesNoted = nil
	for index = #ns.Diagnostics, 1, -1 do
		ns.Diagnostics[index] = nil
	end
	for index = #ns.DiagnosticsInfo, 1, -1 do
		ns.DiagnosticsInfo[index] = nil
	end
	for index = #ns.DiagnosticsVerbose, 1, -1 do
		ns.DiagnosticsVerbose[index] = nil
	end
	for index = #ns.FrameLog, 1, -1 do
		ns.FrameLog[index] = nil
	end
	firstFrameAt = nil
end

--- One line of the addon's own chat output: "Forever Sixty: <text>", the
--- name in gold. Every line the addon prints goes through this.
function Theme.chatLine(text)
	return string.format(L.chatLine, string.format("|cff%s%s|r", Theme.HEX.gold, L.addonName), text)
end

function Theme.say(text)
	print(Theme.chatLine(text))
	return text
end

--- `text` in a token's colour, as the client's inline escape: for one
--- segment of a line whose font string is another colour.
function Theme.inlineColor(hexKey, text)
	return "|cff" .. Theme.HEX[hexKey] .. text .. "|r"
end

--- Six hex digits for a colour given as 0..1 floats.
function Theme.hexOf(r, g, b)
	local function byte(value)
		return math.floor(math.max(0, math.min(1, value or 0)) * 255 + 0.5)
	end
	return string.format("%02x%02x%02x", byte(r), byte(g), byte(b))
end

--- `hexKey`'s colour at `alpha` over the window ground, as six hex digits: a
--- dimmer step of a token for a text escape (which carries no alpha of its own).
function Theme.dimHex(hexKey, alpha)
	local r, g, b = Theme.rgb(Theme.HEX[hexKey])
	local gr, gg, gb = Theme.rgb(Theme.HEX.background)
	return Theme.hexOf(r * alpha + gr * (1 - alpha), g * alpha + gg * (1 - alpha), b * alpha + gb * (1 - alpha))
end

--- Every button's states (L3.4 fix round 3 ruling), one table every surface reads.
--- Hover: a secondary button raised to hover, the primary gold to goldHover, an inactive tab's
--- label muted to body over a 2 px underline in the border colour, the close x muted to body on
--- a hover ground. Pressed: the primary fills goldPressed, the others inset, and every label
--- drops PRESSED_LABEL_DROP px until the mouse comes up. No focus-visible state (no button
--- takes the keyboard) and no disabled style beyond Widgets.setEnabled's alpha, save a
--- secondary button that has acted and waits (Settings' "Positions reset", fix round 2,
--- item 4): its label and its 1 px border in the muted-text grey at full alpha, so it reads.
Theme.CONTROL_STATES = {
	primary = { rest = "gold", hover = "goldHover", pressed = "goldPressed" },
	secondary = { rest = "raised", hover = "hover", pressed = "inset" },
	--- A secondary button's label and 1 px border: at rest, and while it has acted and waits.
	face = { label = "gold", edge = "border", waitingLabel = "muted", waitingEdge = "muted" },
	tab = { label = "muted", hoverLabel = "body", hoverUnderline = "border", pressed = "inset" },
	--- A sortable column header (the guild tab, L5.1 fix round 1, item 5): muted, brightened to
	--- the text white under the cursor and on an inset ground while held; the sorted column
	--- gold whatever the cursor does.
	header = { label = "muted", hoverLabel = "body", active = "gold", pressed = "inset" },
	close = { label = "muted", hoverLabel = "body", hoverGround = "hover", pressed = "inset" },
	--- A Settings switch (L3.5, fix round 1): on, a gold square (goldHover under the cursor,
	--- goldPressed held) with the check in onGold; off, a 1 px outline in the muted-text grey,
	--- body under the cursor, an inset ground while held. The hint is body text at full alpha (A.2).
	toggle = {
		on = "gold", onHover = "goldHover", onPressed = "goldPressed", check = "onGold",
		edge = "muted", edgeHover = "body", offPressed = "inset", hint = "body",
	},
}
Theme.PRESSED_LABEL_DROP = 1

--- The pressed state for any button: `onPress` and `onRelease` paint its ground, and its
--- label drops a pixel from wherever it sits now and comes back on mouse-up.
function Theme.pressable(control, label, onPress, onRelease)
	local resting
	control:SetScript("OnMouseDown", function(self)
		if type(self.IsEnabled) == "function" and not self:IsEnabled() then
			return
		end
		if label ~= nil and type(label.GetPoint) == "function" then
			resting = { label:GetPoint() }
			if resting[1] ~= nil then
				label:ClearAllPoints()
				label:SetPoint(resting[1], resting[2], resting[3], resting[4] or 0,
					(resting[5] or 0) - Theme.PRESSED_LABEL_DROP)
			end
		end
		onPress(self)
	end)
	local function settle()
		if resting ~= nil and resting[1] ~= nil then
			label:ClearAllPoints()
			label:SetPoint(resting[1], resting[2], resting[3], resting[4] or 0, resting[5] or 0)
		end
		resting = nil
	end
	control:SetScript("OnMouseUp", function(self)
		settle()
		if type(self.IsEnabled) == "function" and not self:IsEnabled() then
			-- Turned off while held (combat, say): back to rest, never to hover.
			onRelease(self, false)
			return
		end
		local over = type(self.IsMouseOver) == "function" and self:IsMouseOver() == true
		onRelease(self, over)
	end)
	-- Hidden while held: the label and the ground go back to rest.
	local hidden = control:GetScript("OnHide")
	control:SetScript("OnHide", function(self, ...)
		if resting ~= nil then
			settle()
			onRelease(self, false)
		end
		if hidden ~= nil then
			hidden(self, ...)
		end
	end)
	return control
end

--- The class colour (Theme.classColor) as six hex digits, for a chat or tooltip escape.
function Theme.classHex(token)
	local r, g, b = Theme.classColor(token)
	return Theme.hexOf(r, g, b)
end

function Theme.rgb(hex, alpha)
	return tonumber(hex:sub(1, 2), 16) / 255,
		tonumber(hex:sub(3, 4), 16) / 255,
		tonumber(hex:sub(5, 6), 16) / 255,
		alpha or 1
end

--- The primary button's pressed ground: its gold one step dimmer.
Theme.HEX.goldPressed = Theme.dimHex("gold", 0.85)

-- Contrast on the dark panel (L5.1 fix round 1, item 4) ---------------------------------------

--- The least contrast text keeps on the panel's fill (Theme.HEX.background): `text` at full
--- alpha, `dimmed` on a row drawn at Theme.ALPHA.dim (an offline guildmate).
Theme.CONTRAST = { text = 4.5, dimmed = 3 }

--- A colour's relative luminance (WCAG 2), from six hex digits.
function Theme.luminance(hex)
	local r, g, b = Theme.rgb(hex)
	local function channel(value)
		if value <= 0.03928 then
			return value / 12.92
		end
		return ((value + 0.055) / 1.055) ^ 2.4
	end
	return 0.2126 * channel(r) + 0.7152 * channel(g) + 0.0722 * channel(b)
end

--- The contrast ratio between two colours (1 to 21), each six hex digits.
function Theme.contrast(hexA, hexB)
	local a, b = Theme.luminance(hexA), Theme.luminance(hexB)
	return (math.max(a, b) + 0.05) / (math.min(a, b) + 0.05)
end

--- `hex` drawn at `alpha` over `ground` (six hex digits; the panel's fill when nil).
function Theme.overGround(hex, alpha, ground)
	local r, g, b = Theme.rgb(hex)
	local gr, gg, gb = Theme.rgb(ground or Theme.HEX.background)
	return Theme.hexOf(r * alpha + gr * (1 - alpha), g * alpha + gg * (1 - alpha), b * alpha + gb * (1 - alpha))
end

--- Whether text in `hex` reads on the panel: Theme.CONTRAST.text at full alpha and
--- Theme.CONTRAST.dimmed at Theme.ALPHA.dim.
function Theme.readsOnPanel(hex)
	local ground = Theme.HEX.background
	return Theme.contrast(hex, ground) >= Theme.CONTRAST.text
		and Theme.contrast(Theme.overGround(hex, Theme.ALPHA.dim), ground) >= Theme.CONTRAST.dimmed
end

--- `hex` mixed toward white, a twentieth at a time, until it reads on the panel
--- (Theme.readsOnPanel); a colour that already reads comes back as it is. Each answer is kept:
--- a roster of 97 asks the same few colours on every draw.
local liftCache = {}

function Theme.liftForPanel(hex)
	if liftCache[hex] == nil then
		liftCache[hex] = Theme.liftedHex(hex)
	end
	return liftCache[hex]
end

--- Theme.liftForPanel's answer, worked out.
function Theme.liftedHex(hex)
	local r, g, b = Theme.rgb(hex)
	local lifted = hex
	for step = 1, 20 do
		if Theme.readsOnPanel(lifted) then
			return lifted
		end
		local t = step / 20
		lifted = Theme.hexOf(r + (1 - r) * t, g + (1 - g) * t, b + (1 - b) * t)
	end
	return lifted
end

--- The class colours the dark panel lifts (the guild tab's names): the client's shaman blue,
--- warlock violet, druid orange, mage cyan and warrior tan, each mixed toward white just far
--- enough to keep Theme.CONTRAST on the panel, an offline row's too. Any other class takes
--- the client's colour through Theme.liftForPanel.
Theme.PANEL_CLASS = {
	SHAMAN = "6eaeec",
	WARLOCK = "a2a2f1",
	DRUID = "ff8e2a",
	MAGE = "5ccfee",
	WARRIOR = "cba378",
}

--- A class's name colour on the dark panel: Theme.PANEL_CLASS's, else the client's lifted.
function Theme.panelClassHex(token)
	local mapped = token ~= nil and Theme.PANEL_CLASS[token] or nil
	if mapped ~= nil then
		return mapped
	end
	return Theme.liftForPanel(Theme.classHex(token))
end

--- The score tiers whose client colour does not read on the dark panel, and the Theme.HEX
--- key each is drawn in there instead (common white and uncommon green read as they are).
Theme.PANEL_TIER = { rare = "rareBright", epic = "epicBright", legendary = "legendaryBright" }

--- A tier's colour on the dark panel: its Theme.PANEL_TIER token, else `hex` (the client's
--- quality colour) lifted; nil without either.
function Theme.panelTierHex(tier, hex)
	local key = type(tier) == "string" and Theme.PANEL_TIER[tier] or nil
	if key ~= nil then
		return Theme.HEX[key]
	end
	return hex ~= nil and Theme.liftForPanel(hex) or nil
end

--- Ask the client once whether it has a template for a given frame kind, by
--- trying to use it. There is no API that answers this; building one
--- throwaway frame is the question. The frame is never shown and never
--- reused -- WoW cannot destroy a frame, so this deliberately costs at most
--- one per kind/template pair. The probe is asked with the same kind the
--- real frame will be built with: a template that applies to a Button may
--- not apply to a Frame, so probing with a fixed kind would misreport for
--- any template used on a different one.
function Theme.hasTemplate(kind, name)
	Theme.templates[kind] = Theme.templates[kind] or {}
	local known = Theme.templates[kind][name]
	if known ~= nil then
		return known
	end
	local present = pcall(CreateFrame, kind, nil, UIParent, name)
	Theme.templates[kind][name] = present
	if not present then
		Theme.note(string.format(L.diagNoTemplate, name))
	end
	return present
end

--- A frame with the template if the client has it, bare if it does not.
--- The second return says which, so a caller can draw its own chrome.
--- The probe passing does not guarantee the real create call will succeed
--- too -- a real kind/parent combination the probe did not exercise can
--- still be refused -- so the real call is itself pcall'd; either failure
--- degrades to a bare frame rather than raising out of this file.
function Theme.createFrame(kind, name, parent, templateKey)
	local template = templateKey ~= nil and Theme.TEMPLATES[templateKey] or nil
	if template ~= nil and Theme.hasTemplate(kind, template) then
		local ok, frame = pcall(CreateFrame, kind, name, parent, template)
		if ok then
			return frame, true
		end
		Theme.note(string.format(L.diagNoTemplate, template))
	end
	return CreateFrame(kind, name, parent), false
end

--- A font string on the client's own font, or on the shipped TTF when the
--- font object is missing -- a font string with neither draws nothing.
--- Two distinct ways a font object can be missing: CreateFontString itself
--- can raise, or it can succeed while inheriting nothing (the likelier
--- case -- inherits is only a name, and the client silently drops it when
--- the font object it names does not exist). Both are checked.
function Theme.fontString(parent, layer, fontKey)
	local font = Theme.FONTS[fontKey] or Theme.FONTS.normal
	local ok, region = pcall(parent.CreateFontString, parent, nil, layer, font)
	if ok and region ~= nil and region:GetFont() ~= nil then
		return region
	end
	Theme.note(string.format(L.diagNoTemplate, font))
	region = (ok and region ~= nil) and region or parent:CreateFontString(nil, layer)
	region:SetFont(Theme.FALLBACK_FONT.path, Theme.FALLBACK_FONT.size)
	return region
end

--- The same two questions as fontString, for a region that already exists
--- and so could not be created with a font inherited -- the edit box is the
--- only one. SetFontObject can raise, or it can succeed while the region
--- still reports no font, because the font object it names does not exist.
function Theme.applyFont(region, fontKey)
	local font = Theme.FONTS[fontKey] or Theme.FONTS.normal
	if pcall(region.SetFontObject, region, font) and region:GetFont() ~= nil then
		return region
	end
	Theme.note(string.format(L.diagNoTemplate, font))
	region:SetFont(Theme.FALLBACK_FONT.path, Theme.FALLBACK_FONT.size)
	return region
end

--- Colour a texture from raw 0..1 floats -- a class colour, say, which
--- names no Theme.HEX key of its own. SetColorTexture is the modern
--- name; a client without it takes the colour through SetTexture's
--- four-argument form.
function Theme.paintRGB(texture, r, g, b, a)
	if type(texture.SetColorTexture) == "function" then
		texture:SetColorTexture(r, g, b, a or 1)
	else
		texture:SetTexture(r, g, b, a or 1)
	end
	return texture
end

--- Colour a texture by one of Theme.HEX's own named keys.
function Theme.paint(texture, hexKey, alpha)
	return Theme.paintRGB(texture, Theme.rgb(Theme.HEX[hexKey], alpha))
end

function Theme.texture(parent, layer, hexKey, alpha)
	return Theme.paint(parent:CreateTexture(nil, layer), hexKey, alpha)
end

--- The measure helper (L4.3): a function text -> width in UI units as `region` (a font
--- string) draws it, the unbounded width where the client has it (a font string with a
--- width set reports at most that width otherwise). The width is nil where the client
--- cannot measure, which every fit in this addon reads as "fits".
function Theme.measurer(region)
	return function(text)
		region:SetText(text)
		local width
		if type(region.GetUnboundedStringWidth) == "function" then
			width = region:GetUnboundedStringWidth()
		end
		if type(width) ~= "number" and type(region.GetStringWidth) == "function" then
			width = region:GetStringWidth()
		end
		return type(width) == "number" and width or nil
	end
end

--- One of the client's own textures (a Theme.MEDIA path) in a token's colour: the
--- texture's white is multiplied by it, so a light glyph reads in that colour.
function Theme.tinted(parent, layer, path, hexKey)
	local texture = parent:CreateTexture(nil, layer)
	texture:SetTexture(path)
	texture:SetVertexColor(Theme.rgb(Theme.HEX[hexKey]))
	return texture
end

--- A one-colour border drawn from four textures, which needs no template.
function Theme.outline(parent, thickness, hexKey)
	local edges = {}
	for index, edge in ipairs(Theme.EDGES) do
		local texture = Theme.texture(parent, "OVERLAY", hexKey)
		texture:SetPoint(edge.from, parent, edge.from, 0, 0)
		texture:SetPoint(edge.to, parent, edge.to, 0, 0)
		if edge.horizontal then
			texture:SetHeight(thickness)
		else
			texture:SetWidth(thickness)
		end
		edges[index] = texture
	end
	return edges
end

--- Two stacked half-height textures standing in for a gradient. A real
--- one needs SetGradient, whose argument shape moved between clients;
--- two flat bands read as the site's title bar and cannot break.
function Theme.gradient(parent, topKey, bottomKey)
	local top = Theme.texture(parent, "BACKGROUND", topKey)
	top:SetPoint("TOPLEFT", parent, "TOPLEFT", 0, 0)
	top:SetPoint("RIGHT", parent, "RIGHT", 0, 0)
	local bottom = Theme.texture(parent, "BACKGROUND", bottomKey)
	bottom:SetPoint("BOTTOMLEFT", parent, "BOTTOMLEFT", 0, 0)
	bottom:SetPoint("RIGHT", parent, "RIGHT", 0, 0)
	top:SetPoint("BOTTOM", parent, "CENTER", 0, 0)
	bottom:SetPoint("TOP", parent, "CENTER", 0, 0)
	return top, bottom
end

--- The edge trimmed off a game icon. Every icon in the client carries a
--- baked-in bevel; cropping it is what makes a row of them look designed.
Theme.ICON_CROP = { 0.08, 0.92, 0.08, 0.92 }

--- Icons for the sidebar. All of them are in the base client's icon set.
Theme.NAV_ICONS = {
	overview = "Interface\\Icons\\INV_Misc_Map_01",
	follow = "Interface\\Icons\\INV_Misc_Book_09",
	gear = "Interface\\Icons\\INV_Chest_Chain",
	export = "Interface\\Icons\\INV_Letter_15",
	guild = "Interface\\Icons\\INV_BannerPVP_02",
	settings = "Interface\\Icons\\INV_Misc_Gear_01",
}
Theme.UNKNOWN_ICON = "Interface\\Icons\\INV_Misc_QuestionMark"

--- The client's own empty-slot silhouettes, one per site slot name --
--- stable paths the client has carried since vanilla, not a template or a
--- capability worth probing. Shirt and Tabard carry no stats and are not
--- in Gear.lua's own slot vocabulary at all (ItemTooltip.SLOT_BUTTONS
--- skips them the same way), so they are not here either.
Theme.SLOT_ICONS = {
	head = "Interface\\PaperDollInfoFrame\\UI-PaperDoll-Slot-Head",
	neck = "Interface\\PaperDollInfoFrame\\UI-PaperDoll-Slot-Neck",
	shoulder = "Interface\\PaperDollInfoFrame\\UI-PaperDoll-Slot-Shoulder",
	back = "Interface\\PaperDollInfoFrame\\UI-PaperDoll-Slot-Chest",
	chest = "Interface\\PaperDollInfoFrame\\UI-PaperDoll-Slot-Chest",
	wrist = "Interface\\PaperDollInfoFrame\\UI-PaperDoll-Slot-Wrists",
	hands = "Interface\\PaperDollInfoFrame\\UI-PaperDoll-Slot-Hands",
	waist = "Interface\\PaperDollInfoFrame\\UI-PaperDoll-Slot-Waist",
	legs = "Interface\\PaperDollInfoFrame\\UI-PaperDoll-Slot-Legs",
	feet = "Interface\\PaperDollInfoFrame\\UI-PaperDoll-Slot-Feet",
	finger1 = "Interface\\PaperDollInfoFrame\\UI-PaperDoll-Slot-Finger",
	finger2 = "Interface\\PaperDollInfoFrame\\UI-PaperDoll-Slot-Finger",
	trinket1 = "Interface\\PaperDollInfoFrame\\UI-PaperDoll-Slot-Trinket",
	trinket2 = "Interface\\PaperDollInfoFrame\\UI-PaperDoll-Slot-Trinket",
	main_hand = "Interface\\PaperDollInfoFrame\\UI-PaperDoll-Slot-MainHand",
	off_hand = "Interface\\PaperDollInfoFrame\\UI-PaperDoll-Slot-SecondaryHand",
	ranged = "Interface\\PaperDollInfoFrame\\UI-PaperDoll-Slot-Ranged",
}

--- A game icon with its bevel cropped. `path` may be a texture path or a
--- file id; nil draws the neutral tile rather than nothing.
function Theme.icon(parent, layer, path, size)
	local texture = parent:CreateTexture(nil, layer or "ARTWORK")
	texture:SetSize(size, size)
	texture:SetTexture(path or Theme.UNKNOWN_ICON)
	if type(texture.SetTexCoord) == "function" then
		texture:SetTexCoord(Theme.ICON_CROP[1], Theme.ICON_CROP[2], Theme.ICON_CROP[3], Theme.ICON_CROP[4])
	end
	return texture
end

--- An icon as a `|T<path>:<height>:<width>|t` escape, for a text line
--- that cannot hold a real texture region -- a tooltip line's icon
--- prefix. The neutral tile when the path is not known yet, matching
--- Theme.icon's own default.
function Theme.inlineIcon(path, size)
	local pixels = size or Theme.SIZES.iconSize
	return string.format("|T%s:%d:%d|t", path or Theme.UNKNOWN_ICON, pixels, pixels)
end

--- Grey an icon out, where the client can.
function Theme.desaturate(texture, on)
	if type(texture.SetDesaturated) == "function" then
		texture:SetDesaturated(on)
	end
	return texture
end

--- Fade a frame in. A client without the helper just shows it.
function Theme.fadeIn(frame, seconds)
	if type(UIFrameFadeIn) == "function" and pcall(UIFrameFadeIn, frame, seconds, 0, 1) then
		return true
	end
	frame:SetAlpha(1)
	return false
end

--- The client's own window sounds, by name in SOUNDKIT. Silent without them.
Theme.SOUNDS = { open = "IG_CHARACTER_INFO_OPEN", close = "IG_CHARACTER_INFO_CLOSE" }

function Theme.playSound(key)
	local kit = type(SOUNDKIT) == "table" and SOUNDKIT[Theme.SOUNDS[key] or ""] or nil
	if kit ~= nil and type(PlaySound) == "function" then
		return pcall(PlaySound, kit)
	end
	return false
end

--- A soft shadow past a frame's edges: a darker frame behind it, one
--- strata step down, so it never covers the window's own content.
function Theme.shadow(frame)
	local reach = Theme.SIZES.shadow
	local shadow = Theme.texture(frame, "BACKGROUND", "shadow", Theme.ALPHA.shadow)
	shadow:SetPoint("TOPLEFT", frame, "TOPLEFT", -reach, reach)
	shadow:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", reach, -reach)
	if type(shadow.SetDrawLayer) == "function" then
		shadow:SetDrawLayer("BACKGROUND", -8)
	end
	return shadow
end

--- Let Escape close a frame. UISpecialFrames is a plain client table;
--- a client without it simply has no Escape binding for this window.
function Theme.makeEscapable(frameName)
	if type(UISpecialFrames) ~= "table" then
		return false
	end
	for _, name in ipairs(UISpecialFrames) do
		if name == frameName then
			return true
		end
	end
	table.insert(UISpecialFrames, frameName)
	return true
end

function Theme.setShown(edges, shown)
	for _, edge in ipairs(edges) do
		if shown then
			edge:Show()
		else
			edge:Hide()
		end
	end
	return edges
end

--- The class's own colour for the header name, gold when the client has no
--- table for it or does not know the token (Forever's new combinations).
function Theme.classColor(token)
	local colors = RAID_CLASS_COLORS
	local entry = type(colors) == "table" and token ~= nil and colors[token] or nil
	if entry == nil then
		return Theme.rgb(Theme.HEX.gold)
	end
	return entry.r, entry.g, entry.b, 1
end

--- An item name in its rarity colour. The client's own table when it has
--- one; body text when it does not, which is legible rather than wrong.
function Theme.qualityColor(quality)
	local colors = ITEM_QUALITY_COLORS
	local entry = type(colors) == "table" and quality ~= nil and colors[quality] or nil
	if entry == nil then
		return Theme.rgb(Theme.HEX.body)
	end
	return entry.r, entry.g, entry.b, 1
end

function Theme.hasSettingsApi()
	return type(Settings) == "table"
		and type(Settings.RegisterCanvasLayoutCategory) == "function"
		and type(Settings.RegisterAddOnCategory) == "function"
end

--- Put a panel in the game's own options, however this client does that: the modern
--- Settings canvas first, then (when that is missing or refuses) the interface options.
--- Returns which way it went, and the handle to open it with later.
function Theme.registerSettingsPanel(panel, name)
	if Theme.hasSettingsApi() then
		local ok, category = pcall(Settings.RegisterCanvasLayoutCategory, panel, name)
		if ok and category ~= nil and pcall(Settings.RegisterAddOnCategory, category) then
			return "settings", category
		end
	end
	if type(InterfaceOptions_AddCategory) == "function" and pcall(InterfaceOptions_AddCategory, panel) then
		return "interface", panel
	end
	Theme.note(L.diagNoSettingsPanel)
	return nil, nil
end

--- Whether a frame says it is shown: true or false, or nil when it cannot be asked.
local function isShown(frame)
	if type(frame) ~= "table" or type(frame.IsShown) ~= "function" then
		return nil
	end
	local ok, shown = pcall(frame.IsShown, frame)
	return ok and shown == true
end

--- Open the client's options on a panel registerSettingsPanel registered (`how` and
--- `handle` are what it answered, `panel` the page itself). True only when the options
--- window (else the page) is showing after the call: a client that takes the call and
--- shows nothing has not opened it. Never raises.
function Theme.openSettingsPanel(how, handle, panel)
	local called = false
	if how == "settings" and type(Settings) == "table" and type(Settings.OpenToCategory) == "function" then
		local id = handle
		if type(handle) == "table" then
			id = type(handle.GetID) == "function" and handle:GetID() or handle.ID
		end
		called = id ~= nil and (pcall(Settings.OpenToCategory, id))
	elseif how == "interface" and type(InterfaceOptionsFrame_OpenToCategory) == "function" then
		called = (pcall(InterfaceOptionsFrame_OpenToCategory, handle))
	end
	if not called then
		return false
	end
	local shown = isShown(how == "settings" and _G.SettingsPanel or _G.InterfaceOptionsFrame)
	if shown == nil then
		shown = isShown(panel)
	end
	return shown ~= false
end

--- Register one event, recording rather than raising when the client has
--- never heard of it. With Theme.tryRegisterEvent, the only RegisterEvent
--- calls in the addon.
function Theme.registerEvent(frame, event)
	local ok = Theme.tryRegisterEvent(frame, event)
	if not ok then
		Theme.note(string.format(L.diagNoEvent, event))
	end
	return ok
end

--- Register one event and say whether the client took it, noting nothing:
--- for a caller that words the refusal itself (the path card's click outside).
function Theme.tryRegisterEvent(frame, event)
	return (pcall(frame.RegisterEvent, frame, event))
end

--- Run `action` later. A client with no C_Timer says so rather than
--- running it now: a caller that cannot wait must show the settled state
--- immediately instead of flashing the temporary one.
function Theme.after(seconds, action)
	if type(C_Timer) == "table" and type(C_Timer.After) == "function" then
		C_Timer.After(seconds, action)
		return true
	end
	return false
end

function Theme.inCombat()
	return type(InCombatLockdown) == "function" and InCombatLockdown() == true
end

--- The instance the player is standing in: its name and type ("party",
--- "raid", ...), or nil outdoors (type "none") or on a client without
--- GetInstanceInfo.
function Theme.instanceName()
	if type(GetInstanceInfo) ~= "function" then
		return nil
	end
	local ok, name, instanceType = pcall(GetInstanceInfo)
	if not ok or type(name) ~= "string" or name == "" or instanceType == nil or instanceType == "none" then
		return nil
	end
	return name, instanceType
end

--- Whether Shift is held, for the chip's shift-drag; false where the client cannot say.
function Theme.shiftDown()
	return type(IsShiftKeyDown) == "function" and IsShiftKeyDown() == true
end

--- Open (or close) the client's talent window, whichever way this client
--- does it, and on the Talents tab: the modern player-spells frame opens
--- on whichever tab it last showed (1.60 beta: the spellbook, seen live
--- 16 Oct 2026), so the talents tab is asked for by name, then by its tab
--- id, and only a client without the modern frame gets the classic
--- ToggleTalentFrame. True when a call went through.
function Theme.toggleTalents()
	local util = type(PlayerSpellsUtil) == "table" and PlayerSpellsUtil or nil
	if util and type(util.ToggleClassTalentFrame) == "function" then
		return (pcall(util.ToggleClassTalentFrame))
	end
	if util and type(util.TogglePlayerSpellsFrame) == "function" then
		local tabs = type(util.FrameTabs) == "table" and util.FrameTabs or nil
		local talents = tabs and tabs.ClassTalents or nil
		return (pcall(util.TogglePlayerSpellsFrame, talents))
	end
	if type(ToggleTalentFrame) == "function" then
		return (pcall(ToggleTalentFrame))
	end
	return false
end

function Theme.equip(link)
	if type(C_Item) == "table" and type(C_Item.EquipItemByName) == "function" then
		C_Item.EquipItemByName(link)
		return true
	end
	if type(EquipItemByName) == "function" then
		EquipItemByName(link)
		return true
	end
	Theme.note(L.diagNoEquipApi)
	return false
end

-- Addon messages (contract 2): the client's own channel between players' addons on GUILD,
-- RAID and PARTY, not a network call. C_ChatInfo where the client has it, the old globals
-- where it does not; a client with neither is noted once in /fs diag and sends nothing.
local function noAddonMessages()
	if not Theme.addonMessagesNoted then
		Theme.addonMessagesNoted = true
		Theme.note(L.diagNoAddonMessages)
	end
	return false
end

--- Register an addon-message prefix so CHAT_MSG_ADDON delivers it. True when the client
--- took it (a prefix it already had counts).
function Theme.registerAddonPrefix(prefix)
	local register
	if type(C_ChatInfo) == "table" and type(C_ChatInfo.RegisterAddonMessagePrefix) == "function" then
		register = C_ChatInfo.RegisterAddonMessagePrefix
	elseif type(RegisterAddonMessagePrefix) == "function" then
		register = RegisterAddonMessagePrefix
	else
		return noAddonMessages()
	end
	local ok, result = pcall(register, prefix)
	return ok and result ~= false
end

--- Send one addon message. True when the client took it: a newer client answers with a
--- result code (0 is success), an older one with true, the oldest with nothing.
function Theme.sendAddonMessage(prefix, text, channel)
	local send
	if type(C_ChatInfo) == "table" and type(C_ChatInfo.SendAddonMessage) == "function" then
		send = C_ChatInfo.SendAddonMessage
	elseif type(SendAddonMessage) == "function" then
		send = SendAddonMessage
	else
		return noAddonMessages()
	end
	local ok, result = pcall(send, prefix, text, channel)
	return ok and (result == nil or result == true or result == 0)
end

--- The client's own name for a zone chat channel by its id (2 Trade, 26 LookingForGroup), or
--- nil where it cannot say: the chat badges' channel names (Compat.badgeChannelNames). Here,
--- beside the addon messages, as the one file that names C_ChatInfo.
function Theme.channelShortcut(id)
	if type(C_ChatInfo) ~= "table" or type(C_ChatInfo.GetChannelShortcutForChannelID) ~= "function" then
		return nil
	end
	local ok, name = pcall(C_ChatInfo.GetChannelShortcutForChannelID, id)
	return ok and type(name) == "string" and name ~= "" and name or nil
end

--- Seconds on the client's clock (GetTime), for rate limits; 0 where there is no clock.
function Theme.now()
	if type(GetTime) == "function" then
		return GetTime()
	end
	return 0
end

--- Whether the player is in a guild: true or false, or nil where the client cannot say.
function Theme.inGuild()
	if type(IsInGuild) == "function" then
		local ok, inGuild = pcall(IsInGuild)
		return ok and inGuild ~= nil and inGuild ~= false
	end
	if type(GetGuildInfo) == "function" then
		local ok, guild = pcall(GetGuildInfo, "player")
		return ok and guild ~= nil
	end
	return nil
end

--- A member's rank index in the player's own guild (0 is the guild master), or nil when this
--- client cannot tell. GetGuildInfo answers for a unit the player can see; the guild roster
--- for anyone else in the guild.
function Theme.guildRankOf(fullName, unit)
	if type(GetGuildInfo) ~= "function" then
		return nil
	end
	local okOwn, ownGuild = pcall(GetGuildInfo, "player")
	if not okOwn or ownGuild == nil then
		return nil
	end
	local ok, guild, _, rankIndex = pcall(GetGuildInfo, unit)
	if ok and guild == ownGuild and type(rankIndex) == "number" then
		return rankIndex
	end
	if type(GetNumGuildMembers) ~= "function" or type(GetGuildRosterInfo) ~= "function" then
		return nil
	end
	local okCount, count = pcall(GetNumGuildMembers)
	for index = 1, (okCount and tonumber(count) or 0) do
		local okRow, name, _, rank = pcall(GetGuildRosterInfo, index)
		if okRow and (name == fullName or name == unit) and type(rank) == "number" then
			return rank
		end
	end
	return nil
end

--- The player's guild and rank index (0 is the guild master), or nil where the player has no
--- guild or the client cannot say (GetGuildInfo("player"), its rank name never read).
function Theme.playerGuild()
	if type(GetGuildInfo) ~= "function" then
		return nil
	end
	local ok, name, _, rankIndex = pcall(GetGuildInfo, "player")
	if not ok or name == nil then
		return nil
	end
	return name, type(rankIndex) == "number" and rankIndex or nil
end

--- The guild roster as the client holds it (L5.1): one { name, rankIndex, level, classToken,
--- online } per GetGuildRosterInfo row, the name as the client gives it ("Name-Realm"). Empty
--- where the client has no roster API or has not sent the roster yet. Only read.
function Theme.guildRoster()
	local roster = {}
	if type(GetNumGuildMembers) ~= "function" or type(GetGuildRosterInfo) ~= "function" then
		return roster
	end
	local okCount, count = pcall(GetNumGuildMembers)
	for index = 1, (okCount and tonumber(count) or 0) do
		local ok, name, _, rankIndex, level, _, _, _, _, online, _, classToken = pcall(GetGuildRosterInfo, index)
		if ok and type(name) == "string" and name ~= "" then
			roster[#roster + 1] = {
				name = name, rankIndex = tonumber(rankIndex), level = tonumber(level),
				classToken = type(classToken) == "string" and classToken or nil, online = online == true or online == 1,
			}
		end
	end
	return roster
end

--- Asks the server for the guild roster (C_GuildInfo.GuildRoster, else the old GuildRoster):
--- GUILD_ROSTER_UPDATE answers. True when asked.
function Theme.requestGuildRoster()
	local ask
	if type(C_GuildInfo) == "table" and type(C_GuildInfo.GuildRoster) == "function" then
		ask = C_GuildInfo.GuildRoster
	elseif type(GuildRoster) == "function" then
		ask = GuildRoster
	else
		return false
	end
	return (pcall(ask))
end

--- Whether a unit leads the player's raid or party; false where the client cannot say.
function Theme.isGroupLeader(unit)
	if type(UnitIsGroupLeader) ~= "function" then
		return false
	end
	local ok, leader = pcall(UnitIsGroupLeader, unit)
	return ok and leader == true
end

--- The player's name as UnitName gives it and the realm as an addon-message sender carries
--- it ("Name-Realm", spaces and dashes removed); either nil where the client cannot say.
function Theme.playerNameRealm()
	local name = type(UnitName) == "function" and UnitName("player") or nil
	local realm
	if type(GetNormalizedRealmName) == "function" then
		realm = GetNormalizedRealmName()
	elseif type(GetRealmName) == "function" then
		realm = (GetRealmName() or ""):gsub("[%s%-]", "")
	end
	if realm == "" then
		realm = nil
	end
	return name, realm
end

--- Whether the client's overlay glow can be both put on and taken off.
--- Half a pair is no pair: a client that could show the overlay but never
--- hide it would strand a glow on a button nothing can clear, which is
--- worse than no glow at all because the player acts on it. The addon's
--- own outline, which it can always remove, is used instead, and the
--- missing half is recorded. Latched like hasTemplate: once per session,
--- not once per glow.
local function hasOverlayGlow()
	if Theme.glowPair ~= nil then
		return Theme.glowPair
	end
	local canShow = type(ActionButton_ShowOverlayGlow) == "function"
	local canHide = type(ActionButton_HideOverlayGlow) == "function"
	Theme.glowPair = canShow and canHide
	-- Neither half is the ordinary case and stays silent; only the
	-- mismatch is worth a line in /fs diag.
	if canShow ~= canHide then
		Theme.note(string.format(L.diagNoTemplate,
			canShow and "ActionButton_HideOverlayGlow" or "ActionButton_ShowOverlayGlow"))
	end
	return Theme.glowPair
end

--- How the outline glow breathes: alpha between 1 and `low`, `seconds`
--- each way (spec section 4.1: 0.4 to 1.0 over 1.2 s), so it reads as
--- "next" against the game's own static green and gold talent borders
--- rather than as one more of them.
Theme.GLOW_PULSE = { low = 0.4, seconds = 0.6 }

local function pulse(edges)
	for _, edge in pairs(edges) do
		if type(edge) == "table" and type(edge.CreateAnimationGroup) == "function" and edge.foreverSixtyPulse == nil then
			-- Every step guarded: a client (or the test double) without
			-- animations simply shows a steady outline.
			local group = edge:CreateAnimationGroup()
			local fade = type(group) == "table" and type(group.CreateAnimation) == "function"
				and group:CreateAnimation("Alpha") or nil
			if type(fade) == "table" and type(fade.SetFromAlpha) == "function" then
				fade:SetFromAlpha(1)
				fade:SetToAlpha(Theme.GLOW_PULSE.low)
				fade:SetDuration(Theme.GLOW_PULSE.seconds)
				group:SetLooping("BOUNCE")
				edge.foreverSixtyPulse = group
			end
		end
		if type(edge) == "table" and edge.foreverSixtyPulse ~= nil then
			edge.foreverSixtyPulse:Play()
		end
	end
end

local function stopPulse(edges)
	for _, edge in pairs(edges) do
		if type(edge) == "table" and edge.foreverSixtyPulse ~= nil then
			edge.foreverSixtyPulse:Stop()
		end
		if type(edge) == "table" and type(edge.SetAlpha) == "function" then
			edge:SetAlpha(1)
		end
	end
end

--- L3.7 item 5 (the 1.60 taint rules): a button the glow is put on may be a client
--- frame (the talent window's nodes), and the addon never writes a field into one. The
--- ring's parts and its pending-stop generation are kept here, by button, instead.
local glowStates = setmetatable({}, { __mode = "k" })

--- The addon-side state of `button`'s outline glow: { gold = edges, edge = edges or nil,
--- frame = the ring's frame, generation = n }, or nil before the first glow.
function Theme.glowOf(button)
	return glowStates[button]
end

local function glowState(button)
	local state = glowStates[button]
	if state == nil then
		state = { generation = 0 }
		glowStates[button] = state
	end
	return state
end

--- `outline` asks for the addon's own gold outline even on a client
--- with the overlay glow: the trait talent window's buttons showed
--- nothing under ActionButton_ShowOverlayGlow (seen in game), and an
--- outline the addon draws itself is visible wherever a frame is.
--- `staticAfter`, when given, stops the outline's pulse after that many
--- seconds (Theme.scheduleStatic) and leaves it steady at full alpha: the
--- chip's ring breathes for ten seconds, then holds still. A later showGlow
--- or hideGlow on the same button cancels the pending stop. `ring`, when
--- given ({ inset, thickness, edge }), sizes the outline instead of
--- glowInset and glowThickness, and `edge` adds that many px of dark edge
--- (Theme.ALPHA.ringEdge) just outside it, at that alpha whatever the
--- pulse does, so the ring keeps its separation at the pulse's trough; it
--- is read the first time the button glows.
function Theme.showGlow(button, outline, staticAfter, ring)
	if hasOverlayGlow() and not outline then
		ActionButton_ShowOverlayGlow(button)
		return "overlay"
	end
	local state = glowState(button)
	if state.gold == nil then
		-- Its own frame, a few levels above the button and a little wider
		-- than it: the trait window's buttons draw their rank border on
		-- child frames, which sit above any texture on the button itself
		-- (seen in game: an outline on the button never showed).
		local S = Theme.SIZES
		local inset = ring ~= nil and ring.inset or S.glowInset
		local thickness = ring ~= nil and ring.thickness or S.glowThickness
		local frame = CreateFrame("Frame", nil, button)
		frame:SetPoint("TOPLEFT", button, "TOPLEFT", -inset, inset)
		frame:SetPoint("BOTTOMRIGHT", button, "BOTTOMRIGHT", inset, -inset)
		if type(button.GetFrameLevel) == "function" and type(frame.SetFrameLevel) == "function" then
			frame:SetFrameLevel((button:GetFrameLevel() or 0) + 10)
		end
		state.frame = frame
		state.gold = Theme.outline(frame, thickness, "gold")
		local edge = ring ~= nil and ring.edge or nil
		if type(edge) == "number" and edge > 0 then
			local outer = CreateFrame("Frame", nil, frame)
			outer:SetPoint("TOPLEFT", frame, "TOPLEFT", -edge, edge)
			outer:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", edge, -edge)
			state.edge = Theme.outline(outer, edge, "shadow")
			for _, texture in ipairs(state.edge) do
				Theme.paint(texture, "shadow", Theme.ALPHA.ringEdge)
			end
		end
	end
	if state.edge ~= nil then
		Theme.setShown(state.edge, true)
	end
	Theme.setShown(state.gold, true)
	pulse(state.gold)
	if type(staticAfter) == "number" then
		Theme.scheduleStatic(button, staticAfter, function()
			stopPulse(state.gold)
		end)
	else
		Theme.cancelStatic(button)
	end
	return "texture"
end

--- Run `stop` after `seconds` (Theme.after), unless the button is shown
--- or hidden again first: each call and each Theme.cancelStatic moves the
--- button's generation on, and a callback from an older one does nothing.
function Theme.scheduleStatic(button, seconds, stop)
	local generation = Theme.cancelStatic(button)
	return Theme.after(seconds, function()
		if glowState(button).generation == generation then
			stop()
		end
	end)
end

--- Cancel a pending Theme.scheduleStatic stop; the button's new generation.
function Theme.cancelStatic(button)
	local state = glowState(button)
	state.generation = state.generation + 1
	return state.generation
end

function Theme.hideGlow(button)
	Theme.cancelStatic(button)
	local state = glowState(button)
	if state.gold ~= nil then
		stopPulse(state.gold)
		Theme.setShown(state.gold, false)
	end
	if state.edge ~= nil then
		Theme.setShown(state.edge, false)
	end
	if hasOverlayGlow() and state.gold == nil then
		ActionButton_HideOverlayGlow(button)
		return "overlay"
	end
	if hasOverlayGlow() then
		ActionButton_HideOverlayGlow(button)
	end
	return "texture"
end

--- Whether the client has a GameTooltip at all. Shared by every function
--- below that touches it, so the guard is written once.
local function hasTooltip()
	return type(GameTooltip) == "table"
end

--- The game's own item tooltip. SetHyperlink when there is a link (it
--- carries enchants and suffixes); SetItemByID for a planned item the
--- player has never seen, which has no link yet. Each is guarded on its
--- own: this file's contract is that a missing capability degrades, never
--- raises, and SetItemByID in particular is a genuinely open question on
--- the 1.60 client.
function Theme.showItemTooltip(owner, itemId, link)
	if not hasTooltip() then
		return false
	end
	GameTooltip:SetOwner(owner, "ANCHOR_RIGHT")
	if link ~= nil then
		if type(GameTooltip.SetHyperlink) ~= "function" then
			Theme.note(string.format(L.diagNoTooltipApi, "SetHyperlink"))
			return false
		end
		GameTooltip:SetHyperlink(link)
	elseif itemId ~= nil then
		if type(GameTooltip.SetItemByID) ~= "function" then
			Theme.note(string.format(L.diagNoTooltipApi, "SetItemByID"))
			return false
		end
		GameTooltip:SetItemByID(itemId)
	end
	GameTooltip:Show()
	return true
end

function Theme.hideTooltip()
	if not hasTooltip() then
		return false
	end
	GameTooltip:Hide()
	return true
end

--- A plain multi-line tooltip, for the minimap button and the chip.
--- `anchor` is GameTooltip's own anchor name, ANCHOR_LEFT when nil. A line is a string, or
--- { text = ..., wrap = true } for one the tooltip wraps at its own width rather than widen
--- to (the panel's full rotation condition, L4.3 fix round 3); its colour is its own escape.
function Theme.showLines(owner, lines, anchor)
	if not hasTooltip() then
		return false
	end
	GameTooltip:SetOwner(owner, anchor or "ANCHOR_LEFT")
	GameTooltip:ClearLines()
	for _, line in ipairs(lines) do
		if type(line) == "table" then
			GameTooltip:AddLine(line.text, 1, 1, 1, line.wrap == true)
		else
			GameTooltip:AddLine(line)
		end
	end
	GameTooltip:Show()
	return true
end

--- Add one line, in a token's colour (inline escapes in `text` win), to
--- GameTooltip while it is showing for `owner`: the client's own tooltip (a
--- talent node's, say) gains it after the client drew it, after a blank
--- line when `gap` is true. Nothing when another frame owns the tooltip.
function Theme.addTooltipLine(owner, text, hexKey, gap)
	if not hasTooltip() or type(GameTooltip.GetOwner) ~= "function" then
		return false
	end
	if GameTooltip:GetOwner() ~= owner or not GameTooltip:IsShown() then
		return false
	end
	if gap then
		GameTooltip:AddLine(" ")
	end
	GameTooltip:AddLine(text, Theme.rgb(Theme.HEX[hexKey or "gold"]))
	GameTooltip:Show()
	return true
end

--- Run `action` after one of a client frame's own scripts, without
--- replacing it. False on a frame that cannot be hooked.
function Theme.hookScript(frame, script, action)
	if type(frame) ~= "table" or type(frame.HookScript) ~= "function" then
		return false
	end
	return (pcall(frame.HookScript, frame, script, action))
end

Theme.COMPARE_TOOLTIP_NAME = "ForeverSixtyCompareTooltip"
--- Clearance between GameTooltip's right edge and the compare tooltip,
--- the same idea as the client's own shift-to-compare tooltip sitting
--- beside an item's.
Theme.COMPARE_TOOLTIP_GAP = 4

--- The BiS hover's second tooltip: the client's own item tooltip for the
--- recommended item, anchored beside GameTooltip. Lazily created on the
--- first call and reused after that -- WoW cannot destroy a frame, so
--- creating a second one on every hover would leak one per hover instead
--- of reusing the one this addon will ever need. Hidden automatically
--- whenever GameTooltip hides (hooked once, here, since this is the one
--- place that creates the frame that needs it).
function Theme.compareTooltip()
	if Theme.compareTooltipFrame ~= nil then
		return Theme.compareTooltipFrame
	end
	local frame = Theme.createFrame("GameTooltip", Theme.COMPARE_TOOLTIP_NAME, UIParent, "gameTooltip")
	if type(frame.SetFrameStrata) == "function" then
		frame:SetFrameStrata("TOOLTIP")
	end
	Theme.compareTooltipFrame = frame
	if hasTooltip() and type(GameTooltip.HookScript) == "function" then
		GameTooltip:HookScript("OnHide", function()
			Theme.hideCompareTooltip()
		end)
	end
	return frame
end

--- The compare tooltip's tab ("Next upgrade"): a small label on its top
--- edge, made the first time one is asked for. Hidden with no `label`; a
--- client whose frame cannot carry one simply shows none.
local function setCompareLabel(tooltip, label)
	local region = tooltip.foreverSixtyLabel
	if label == nil then
		if region ~= nil then
			region:Hide()
		end
		return
	end
	if region == nil then
		local ok, made = pcall(Theme.fontString, tooltip, "OVERLAY", "small")
		if not ok or made == nil then
			return
		end
		made:SetPoint("BOTTOMLEFT", tooltip, "TOPLEFT", Theme.COMPARE_TOOLTIP_GAP, 2)
		tooltip.foreverSixtyLabel = made
		region = made
	end
	region:SetText(label)
	region:Show()
end

--- Show the compare tooltip beside GameTooltip for one item: `link` when
--- it is known (it carries enchants and suffixes), else `itemId` for a
--- recommended item the player has never seen, with `label` as its tab
--- when given. Guarded exactly like Theme.showItemTooltip: a missing
--- capability degrades to "shown nothing", never raises.
function Theme.showCompareTooltip(itemId, link, label)
	if not hasTooltip() then
		return false
	end
	local tooltip = Theme.compareTooltip()
	tooltip:SetOwner(GameTooltip, "ANCHOR_NONE")
	if type(tooltip.ClearAllPoints) == "function" then
		tooltip:ClearAllPoints()
	end
	tooltip:SetPoint("TOPLEFT", GameTooltip, "TOPRIGHT", Theme.COMPARE_TOOLTIP_GAP, 0)
	if link ~= nil then
		if type(tooltip.SetHyperlink) ~= "function" then
			return false
		end
		tooltip:SetHyperlink(link)
	elseif itemId ~= nil then
		if type(tooltip.SetItemByID) ~= "function" then
			return false
		end
		tooltip:SetItemByID(itemId)
	else
		return false
	end
	setCompareLabel(tooltip, label)
	tooltip:Show()
	return true
end

function Theme.hideCompareTooltip()
	if Theme.compareTooltipFrame ~= nil then
		Theme.compareTooltipFrame:Hide()
	end
	return nil
end

ns.Theme = Theme
return Theme
