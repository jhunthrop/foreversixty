-- addon/ForeverSixty/surfaces/Chip.lua
-- The chip (design/specs/2026-11-13-addon-hub.md 6): the addon's one launcher, beside the minimap. The
-- class crest in its ring and one line naming the most urgent thing (Hub.urgent's first item, with
-- tonight's raid ahead of it once you answered), "+N" muted for the others, at most 280 px with the
-- figure kept. A click opens the window on the tab the line names (again on that tab: it closes); the
-- hover lists everything waiting; a right-click opens Settings; shift-drag moves it. In combat the line
-- reads "After combat" and a click does nothing. The minimap button and the addon compartment entry are
-- its twins: the same Chip.click and the same Chip.hoverLines.
--
-- Pure model, thin frame: Chip.model turns a plain table of inputs into what the chip says and opens;
-- Chip.hoverLines turns a model into the hover; Chip.inputs gathers the table from the hub and the
-- client; Chip.refresh draws it. The frame is built on first need, never at load.
--
-- The file also keeps the small shared rules other surfaces read through it (the talent's name and
-- rank, the spec picked, the cut that fits a line, a percentile's ordinal).
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Prefs = ns.Prefs or require("Prefs")
local Compat = ns.Compat or require("Compat")
local Talents = ns.Talents or require("Talents")
local Export = ns.Export or require("Export")
local Follow = ns.Follow or require("Follow")
local Gear = ns.Gear or require("Gear")
local Progress = ns.Progress or require("Progress")
local TalentClick = ns.TalentClick or require("TalentClick")
local TalentGlow = ns.TalentGlow or require("TalentGlow")

local Chip = {}

Chip.FRAME_NAME = "ForeverSixtyChip"
--- The secure cover that opens the talent window from the chip (Chip.syncTalentButton).
Chip.TALENT_BUTTON_NAME = "ForeverSixtyChipTalents"
--- How far over the chip the secure button sits.
Chip.TALENT_BUTTON_RAISE = 20
--- The chip's alpha in combat when it stays shown (the "Hide in combat" switch off).
Chip.COMBAT_ALPHA = 0.25
--- The first-run pulse (2026-11-08-addon-ux.md 3.2): a gold ring breathing 0.3 to 1.0 and back over
--- PULSE_SECONDS, PULSE_CYCLES times after each login, until the chip is first clicked or hovered
--- (ForeverSixtyDB.chipSeenAt), giving up after PULSE_LOGINS logins (ForeverSixtyDB.chipPulses).
Chip.PULSE_SECONDS = 1.6
Chip.PULSE_CYCLES = 5
Chip.PULSE_LOGINS = 3
Chip.PULSE_LOW = 0.3
--- Where the chip goes on a client with no minimap frame to sit beside.
Chip.FALLBACK_X = -200
Chip.FALLBACK_Y = -16
--- The chip's frame level inside its strata: over the buff icons and the minimap's own furniture.
Chip.FRAME_LEVEL = 100
--- The waiting item whose click is the talent window too: the Builds tab's points to spend.
Chip.TALENT_ITEM = { tab = "builds", key = "spend" }

--- Injection points Options fills: what to redraw after a spec is picked (every surface, so the
--- talent marks follow the new build too), and the Settings page (a right-click).
Chip.afterChoice = nil
Chip.openSettings = nil

--- A module loaded after this file in the TOC (the hub and its tabs), looked up when it is needed;
--- specs load modules apart, so they are required there.
local function late(name)
	local module = ns[name] or package.loaded[name]
	if module == nil and type(require) == "function" then
		local ok, loaded = pcall(require, name)
		module = ok and loaded or nil
	end
	return type(module) == "table" and module or nil
end

-- Shared rules other surfaces read ------------------------------------------------------------

--- A talent's name, or nil for one the data cannot name; never a cell like "3:2".
function Chip.talentName(talent)
	if type(talent) == "table" and type(talent.name) == "string" and talent.name ~= "" then
		return talent.name
	end
	return nil
end

--- The next talent with the rank a click takes, "Improved Heroic Strike (1 of 3)", where the plan
--- wants more than one rank there and the cell is not yet full (`rank` is { have, want } from
--- Progress.cells, counting a click staged in the talent window); else the name. The talent strip's
--- leading talent and the Builds tab read it.
function Chip.pointLabel(name, rank)
	if type(rank) == "table" and (rank.want or 0) > 1 and (rank.have or 0) < rank.want then
		return string.format(L.chipPointRank, name, (rank.have or 0) + 1, rank.want)
	end
	return name
end

--- The plan's { have, want } at the committed next point's cell, `have` read from `ranks`
--- (Talents.readStagedRanks' in the game, so a staged click counts); nil without a point. Pure.
function Chip.pointRankAt(build, ranks, data, point)
	if point == nil then
		return nil
	end
	for _, cell in ipairs(Progress.cells(build, ranks, data)) do
		if cell.tab == point.tab and cell.tier == point.tier and cell.column == point.column then
			return { have = cell.have, want = cell.want }
		end
	end
	return nil
end

--- "71st": a percentile as a player reads it, rounded to a whole number.
function Chip.ordinal(number)
	local whole = math.floor((tonumber(number) or 0) + 0.5)
	local suffix = L.ordinalOther
	local lastTwo = whole % 100
	if lastTwo < 11 or lastTwo > 13 then
		suffix = L.ordinalSuffixes[whole % 10] or L.ordinalOther
	end
	return string.format(L.ordinal, whole, suffix)
end

--- `text` without its last character: the trailing UTF-8 continuation bytes and the byte that leads
--- them, so a name is never cut mid-letter.
local function dropLastCharacter(text)
	local stop = #text
	while stop > 0 do
		local byte = text:byte(stop)
		stop = stop - 1
		if byte < 128 or byte >= 192 then
			break
		end
	end
	return text:sub(1, stop)
end

--- `text` cut character by character until it fits, ended with L.ellipsis (then `tail`, kept whole)
--- and never a dangling space or comma; "" when nothing but the ellipsis would.
local function ellipsise(text, measure, maxWidth, tail)
	local body = text
	while body ~= "" do
		body = dropLastCharacter(body)
		local trimmed = body:gsub("[%s,]+$", "")
		if trimmed == "" then
			return ""
		end
		local candidate = trimmed .. L.ellipsis .. (tail or "")
		local width = measure(candidate)
		if width == nil or width <= maxWidth then
			return candidate
		end
	end
	return ""
end

local function fits(text, measure, maxWidth)
	local width = measure(text)
	return width == nil or width <= maxWidth
end

--- The line that fits `maxWidth` as `measure` measures it, in two stages: the whole `text` when it
--- fits; else `short` when there is one and it fits; only past that, cut with an ellipsis. `keep` is
--- the end of `short` that is never cut (a figure, " +11.2"): the part before it gives way instead. A
--- measure that answers nothing counts as fitting.
function Chip.fit(text, measure, maxWidth, short, keep)
	text = text or ""
	if fits(text, measure, maxWidth) then
		return text
	end
	if short == nil or short == "" then
		return ellipsise(text, measure, maxWidth)
	end
	if fits(short, measure, maxWidth) then
		return short
	end
	if keep ~= nil and keep ~= "" and #keep < #short and short:sub(-#keep) == keep then
		local kept = ellipsise(short:sub(1, #short - #keep), measure, maxWidth, keep)
		if kept ~= "" then
			return kept
		end
		-- Not even the figure fits beside a letter of the name: cut the short line whole.
	end
	return ellipsise(short, measure, maxWidth)
end

--- The tree a spec slug's part ("feral-bear") belongs to: the tab whose own slug starts with the same
--- word ("feral-combat"), else after every tree.
local function treeIndex(class, part)
	local word = part:match("^[^-]+")
	for index, tab in ipairs(type(class) == "table" and class.tabs or {}) do
		if Export.slugify(tab.name or ""):match("^[^-]+") == word then
			return index
		end
	end
	return math.huge
end

--- The specs a class can pick: every spec the data has a default build for, in talent-tree order,
--- each { slug, name, playstyle, next }. A spec that shares a tree ("feral-bear") sits beside it.
--- `next` (the talent the character would take next on that spec's default, Follow.defaultNext) is
--- filled only when `ranks` is given.
function Chip.pathSpecs(data, classSlug, ranks, level)
	local list = {}
	local builds = type(data) == "table" and type(data.defaultBuilds) == "table" and data.defaultBuilds or nil
	if builds == nil or type(classSlug) ~= "string" then
		return list
	end
	local class = type(data.classes) == "table" and data.classes[classSlug] or nil
	local prefix = classSlug .. "-"
	for slug in pairs(builds) do
		if type(slug) == "string" and slug:sub(1, #prefix) == prefix then
			local part = slug:sub(#prefix + 1)
			list[#list + 1] = {
				slug = slug,
				name = Talents.titleCase(part),
				playstyle = L.specPlaystyle[slug],
				tree = treeIndex(class, part),
				next = ranks ~= nil and Follow.defaultNext(data, slug, ranks, level) or nil,
			}
		end
	end
	table.sort(list, function(left, right)
		if left.tree ~= right.tree then
			return left.tree < right.tree
		end
		return left.slug < right.slug
	end)
	return list
end

--- The player's class token ("ROGUE"), or nil.
function Chip.classToken()
	local slug = Talents.playerClassSlug()
	return slug ~= nil and slug:upper() or nil
end

--- The level, as the last PLAYER_LEVEL_UP said it if the client has not caught up yet (UnitLevel can
--- still answer the old level during the event).
function Chip.level()
	return math.max(Compat.playerLevel() or 0, Chip.eventLevel or 0)
end

--- PLAYER_LEVEL_UP: remember the new level.
function Chip.onLevelUp(level)
	if type(level) == "number" then
		Chip.eventLevel = level
	end
end

--- This character's saved choices (Prefs.chipCharacter: a key saved under the old empty region still
--- counts), or nil; `create` makes it.
local function characterPrefs(create)
	return Prefs.chipCharacter(Export.characterKey(), Follow.sameCharacter, create)
end

--- What Follow.ensureDefault takes at login for a character that picked a spec: { spec = slug, kind },
--- or nil to let the points decide (kind nil: the level decides). The pref is account-wide, so it
--- counts only for the character that chose it, and only for its class.
function Chip.chosenPath()
	local classSlug = Talents.playerClassSlug()
	local mine = characterPrefs(false)
	local slug = mine ~= nil and mine.pathChosen or nil
	if type(slug) ~= "string" or slug == "" then
		return nil
	end
	if classSlug ~= nil and slug:sub(1, #classSlug + 1) ~= classSlug .. "-" then
		return nil
	end
	local kind = mine.pathKind
	return { spec = slug, kind = (kind == "leveling" or kind == "raid") and kind or nil }
end

--- The player's answer to the spec question: follow the default build for `slug` (Follow.chooseDefault,
--- over a saved build too) and remember the choice for this character. The build followed.
function Chip.choosePath(slug, kind)
	local build = Follow.chooseDefault(Chip.data or ns.Data, slug, nil, kind)
	local mine = characterPrefs(true)
	if mine ~= nil then
		-- The kind only when the player picked one (the Builds tab's rows); else the level decides.
		mine.pathChosen, mine.pathKind = slug, kind
	end
	if Chip.afterChoice ~= nil then
		Chip.afterChoice()
	else
		Chip.refresh()
	end
	return build
end

--- The bag upgrades scored as the tooltip scores them (Gear.currentWeights at `level`): the sim's
--- weights and its DPS per point while a weights message for `spec` is in use, else the shipped
--- weights and the band's DPS per point (`perPoint`, Gear.dpsFigure's). Pure over its arguments.
function Chip.scoredUpgrades(data, spec, perPoint, inbox, key, build, ranks, bagLinks, equipped, level)
	local current = Gear.currentWeights(data, spec, inbox, key, level)
	local dps = type(perPoint) == "table" and tonumber(perPoint.dps) or nil
	if type(current) == "table" and (current.tag == "sim" or current.dpsPerPoint ~= nil) then
		dps = current.dpsPerPoint
	end
	return Progress.bagUpgrades(data, build, ranks, bagLinks, equipped, dps, level, current)
end

-- The launcher (pure) ----------------------------------------------------------------------------

--- The end of a line that is never cut: the answer or count after " · " ("· Sign up", "· 2 of 2
--- answered"), else a trailing figure (" +11.2", in its colour or not). nil when the line has neither.
--- Pure.
function Chip.figureOf(line)
	if type(line) ~= "string" then
		return nil
	end
	return line:match("( · .*)$") or line:match("( %+%d[%d%.]*[^·]*)$")
		or line:match("( |c%x%x%x%x%x%x%x%x%+%d[^·]*)$")
end

--- What the chip lists, in order: tonight's raid once you answered (`raidDay`, GuildRaid.tonightLine's;
--- it is not waiting, so it carries no dot, but it leads the chip on the day), then Hub.urgent's items.
--- Each { tab, key, line }. Pure.
function Chip.entries(urgent, raidDay)
	local entries = {}
	if type(raidDay) == "table" and type(raidDay.line) == "string" then
		entries[1] = { tab = "guild", key = "tonight", line = raidDay.line }
	end
	for _, item in ipairs(type(urgent) == "table" and urgent.items or {}) do
		entries[#entries + 1] = { tab = item.tab, key = item.key, line = item.line }
	end
	return entries
end

--- The chip from a plain table: { entries (Chip.entries'), landing (the tab an open with nothing
--- waiting lands on, for the hover), starts (tonight's raid's countdown, for the hover), talentsOpen,
--- inCombat, combatHide, shown }. Answers { line, muted, more, opens (the tab a click asks for; nil: the
--- hub's own landing), hintTab (the tab the hover names), entries, starts, talent (the click opens the
--- talent window too: a point to spend, the window not open yet), visible, hiddenBy, alpha, mouse }. Pure.
function Chip.model(input)
	local entries = input.entries or {}
	local first = entries[1]
	local model = { entries = entries, starts = input.starts, more = 0, alpha = 1, mouse = true, visible = true }
	if input.inCombat then
		model.line, model.muted, model.alpha, model.mouse = L.chipAfterCombat, true, Chip.COMBAT_ALPHA, false
	elseif first == nil then
		model.line, model.muted, model.hintTab = L.addonName, true, input.landing
	else
		model.line, model.opens, model.hintTab, model.more = first.line, first.tab, first.tab, #entries - 1
		model.talent = first.tab == Chip.TALENT_ITEM.tab and first.key == Chip.TALENT_ITEM.key
			and not input.talentsOpen
	end
	if input.shown == false then
		model.visible, model.hiddenBy = false, "player"
	elseif input.inCombat and input.combatHide ~= false then
		model.visible, model.hiddenBy = false, "combat"
	end
	return model
end

--- A hover hint: the client tooltip's muted colour.
local function hinted(text)
	return Theme.inlineColor("tooltipMuted", text)
end

--- The hover (spec 6): titled Forever Sixty, every item in order with the tab it opens on the right
--- (`labelOf` gives a tab's label), how long until tonight's raid under its first line, then what a
--- click does ("Click: open on Guild", the tab the line names, and the talent window with a point to
--- spend), the right-click and, when there is one, `dragHint` (the chip's "Shift-drag: move"). Pure.
function Chip.hoverLines(model, labelOf, dragHint)
	local lines = { L.addonName }
	local startsShown = false
	for _, entry in ipairs(model.entries or {}) do
		lines[#lines + 1] = { text = entry.line, right = labelOf ~= nil and labelOf(entry.tab) or nil }
		if entry.tab == "guild" and model.starts ~= nil and not startsShown then
			startsShown = true
			lines[#lines + 1] = hinted(string.format(L.chipStarts, model.starts))
		end
	end
	local label = model.hintTab ~= nil and labelOf ~= nil and labelOf(model.hintTab) or nil
	local click = model.talent and L.chipHintOpenTalents or L.chipHintOpenOn
	lines[#lines + 1] = hinted(label ~= nil and string.format(click, label) or L.chipHintOpen)
	lines[#lines + 1] = hinted(L.chipHintSettings)
	if dragHint ~= nil then
		lines[#lines + 1] = hinted(dragHint)
	end
	return lines
end

-- The live inputs ------------------------------------------------------------------------------

--- Tonight's raid (GuildRaid.tonightLine: its countdown, and its line once nothing waits on it), while
--- Guild is a tab.
function Chip.raidDay(hub)
	local guildRaid = late("GuildRaid")
	if guildRaid == nil or hub == nil or not hub.isAvailable(hub.tab("guild")) then
		return nil
	end
	local ok, line = pcall(function()
		return guildRaid.tonightLine(guildRaid.liveRaids(), guildRaid.context())
	end)
	return ok and line or nil
end

--- The model's inputs, read from the hub and the client.
function Chip.inputs()
	local hub = late("Hub")
	local urgent = hub ~= nil and hub.urgent() or nil
	local raidDay = Chip.raidDay(hub)
	return {
		entries = Chip.entries(urgent, raidDay),
		starts = raidDay ~= nil and raidDay.starts or nil,
		landing = hub ~= nil and hub.resolveTab(hub.landing(), Prefs.get("hub", "tab")) or nil,
		talentsOpen = TalentGlow.openFrame() ~= nil,
		inCombat = Chip.combat == true or Theme.inCombat(),
		combatHide = Prefs.get("chip", "combatHide"),
		shown = Prefs.get("chip", "shown"),
	}
end

--- A tab's label on the strip ("Guild"), for the hover.
function Chip.tabLabel(id)
	local hub = late("Hub")
	local tab = hub ~= nil and hub.tab(id) or nil
	return tab ~= nil and tab.label or nil
end

--- The hover for the model now, with `dragHint` (the chip's, the minimap button's, or none).
function Chip.lines(dragHint)
	if Chip.current == nil then
		Chip.current = Chip.model(Chip.inputs())
	end
	return Chip.hoverLines(Chip.current, Chip.tabLabel, dragHint)
end

--- A click on the chip, the minimap button or the compartment entry: right opens Settings; left opens
--- the window on the tab the line names (Hub.toggle: again on that tab, it closes), never in combat.
function Chip.click(mouseButton)
	Theme.hideTooltip()
	Chip.markSeen()
	if mouseButton == "RightButton" then
		if Chip.openSettings ~= nil then
			Chip.openSettings()
		end
		return nil
	end
	local hub = late("Hub")
	if hub == nil or Chip.combat or Theme.inCombat() then
		return nil
	end
	local model = Chip.current or Chip.model(Chip.inputs())
	if model.talent then
		-- A client without the talent micro button has no secure cover: the window opens from here
		-- (TalentClick.open does nothing where the cover took the click).
		TalentClick.open()
	end
	local opened = hub.toggle(model.opens)
	Chip.refresh()
	return opened
end

--- Whether the player has put the chip away (the Settings switch).
function Chip.isHidden()
	return Prefs.get("chip", "shown") == false
end

--- PLAYER_REGEN_DISABLED / ENABLED: InCombatLockdown is not yet true while the first of them fires,
--- so the chip keeps its own flag.
function Chip.setCombat(inCombat)
	Chip.combat = inCombat == true
	if Chip.combat then
		Chip.stopPulse()
	end
	-- Combat is starting and the lockdown is not on yet: the last moment the secure button can go.
	Chip.syncTalentButton()
end

-- Where it sits --------------------------------------------------------------------------------

--- Whether a shift-dragged spot can be trusted: a named anchor point and two numbers inside the screen
--- (UIParent's size, when the client has told it). A spot that fails this is not drawn at: the chip
--- sits beside the minimap instead, so a bad number can never put it off screen.
function Chip.savedSpotUsable(point, x, y)
	if type(point) ~= "string" or point == "" or type(x) ~= "number" or type(y) ~= "number" then
		return false
	end
	local parent = _G.UIParent
	local width = type(parent) == "table" and type(parent.GetWidth) == "function" and parent:GetWidth() or nil
	local height = type(parent) == "table" and type(parent.GetHeight) == "function" and parent:GetHeight() or nil
	if type(width) == "number" and width > 0 and math.abs(x) >= width then
		return false
	end
	if type(height) == "number" and height > 0 and math.abs(y) >= height then
		return false
	end
	return true
end

--- Whether the chip is drawn now: built, shown, and not turned off for the session.
function Chip.isDrawn()
	return not Chip.disabled and Chip.frame ~= nil and Chip.frame:IsShown() == true
end

local function edges(region)
	local read, left, right, top, bottom = pcall(function()
		return region:GetLeft(), region:GetRight(), region:GetTop(), region:GetBottom()
	end)
	if not read or type(left) ~= "number" or type(right) ~= "number" or type(top) ~= "number"
		or type(bottom) ~= "number" then
		return nil
	end
	return left, right, top, bottom
end

--- Where to put the chip under the client's buff frame when that frame sits where the chip's default
--- spot would be: { x, y } offsets of the chip's top-right corner from UIParent's, plain numbers, never
--- an anchor to the buff frame. nil when the buffs are elsewhere, not drawn, unreadable, or the spot
--- under them would not be on screen. Only asked, never written.
function Chip.buffsInTheWay(minimap, frame)
	local buffs = _G.BuffFrame
	if Chip.noBuffMove or type(buffs) ~= "table" or type(buffs.GetLeft) ~= "function"
		or type(minimap.GetLeft) ~= "function" then
		return nil
	end
	local ok, shown = pcall(function()
		return type(buffs.IsShown) ~= "function" or buffs:IsShown()
	end)
	local left, right, top, bottom = edges(buffs)
	local minimapLeft, minimapTop = minimap:GetLeft(), minimap:GetTop()
	if not ok or not shown or left == nil or type(minimapLeft) ~= "number" or type(minimapTop) ~= "number" then
		return nil
	end
	local S = Theme.SIZES
	local width, height = frame:GetWidth(), frame:GetHeight()
	height = type(height) == "number" and height or S.chip
	local chipRight = minimapLeft - S.chipMinimapGap
	local chipLeft = chipRight - (type(width) == "number" and width or S.chip)
	if not (right > chipLeft and left < chipRight and top > minimapTop - height and bottom < minimapTop) then
		return nil
	end
	local parent = _G.UIParent
	local parentRight = type(parent) == "table" and type(parent.GetRight) == "function" and parent:GetRight() or nil
	local parentTop = type(parent) == "table" and type(parent.GetTop) == "function" and parent:GetTop() or nil
	if type(parentRight) ~= "number" or type(parentTop) ~= "number" or bottom <= 0 or bottom >= top then
		return nil
	end
	local x = chipRight - parentRight
	local under = bottom - S.chipMinimapGap / 2
	-- The spot must leave the whole chip on screen (its height below the top edge, its right edge in).
	if x > 0 or chipLeft < 0 or under - height < 0 then
		return nil
	end
	return { x = x, y = under - parentTop }
end

--- Where the chip sits: the saved spot after a shift-drag, else just left of the minimap, its top edge
--- level with the minimap's; where the client's buff icons are in that spot, just under them instead.
function Chip.place()
	local frame = Chip.frame
	if frame == nil then
		return nil
	end
	frame:ClearAllPoints()
	local point = Prefs.get("chip", "point")
	if Chip.savedSpotUsable(point, Prefs.get("chip", "x"), Prefs.get("chip", "y")) then
		frame:SetPoint(point, UIParent, point, Prefs.get("chip", "x"), Prefs.get("chip", "y"))
		return point
	end
	local minimap = _G.Minimap
	local width = type(minimap) == "table" and type(minimap.GetWidth) == "function" and minimap:GetWidth() or nil
	if type(width) == "number" and width > 0 then
		local under = Chip.buffsInTheWay(minimap, frame)
		if under ~= nil then
			frame:SetPoint("TOPRIGHT", UIParent, "TOPRIGHT", under.x, under.y)
			return "buffs"
		end
		frame:SetPoint("TOPRIGHT", minimap, "TOPRIGHT", -(width + Theme.SIZES.chipMinimapGap), 0)
		return "minimap"
	end
	frame:SetPoint("TOPRIGHT", UIParent, "TOPRIGHT", Chip.FALLBACK_X, Chip.FALLBACK_Y)
	return "fallback"
end

-- Reset positions moves the chip at once, without a reload.
Prefs.onReset[#Prefs.onReset + 1] = function()
	Chip.place()
end

--- After a shift-drag: the chip's top-right corner against UIParent's, so it keeps growing leftwards
--- from where the player put it.
function Chip.savePosition()
	local frame = Chip.frame
	local right, top = frame:GetRight(), frame:GetTop()
	local parentRight, parentTop = UIParent:GetRight(), UIParent:GetTop()
	if right == nil or top == nil or parentRight == nil or parentTop == nil then
		return nil
	end
	Prefs.set("chip", "point", "TOPRIGHT")
	Prefs.set("chip", "x", right - parentRight)
	Prefs.set("chip", "y", top - parentTop)
	Chip.place()
	return true
end

-- The mouse ------------------------------------------------------------------------------------

function Chip.onClick(_, mouseButton)
	Chip.click(mouseButton)
end

function Chip.onEnter(frame)
	Chip.markSeen()
	-- Read again first: the countdown and the items are as of this hover.
	Chip.refresh()
	Theme.showLines(frame, Chip.lines(L.chipHintDrag), "ANCHOR_BOTTOMLEFT")
end

function Chip.onLeave()
	Theme.hideTooltip()
end

function Chip.onDragStart(frame)
	if Theme.shiftDown() then
		Chip.dragging = true
		Theme.hideTooltip()
		frame:StartMoving()
	end
end

function Chip.onDragStop(frame)
	if Chip.dragging then
		Chip.dragging = false
		frame:StopMovingOrSizing()
		Chip.savePosition()
	end
end

-- The talent click ----------------------------------------------------------------------------
-- While the line is a point to spend, a click opens the talent window as well as the Builds tab. The
-- window opens only through a secure cover (TalentClick), never from the chip's own code.

--- After the cover's click: the left click was the client's own (the talent window); the chip's own
--- click follows (the Builds tab), and a right click is the chip's right click.
function Chip.onTalentCoverPostClick(_, mouseButton)
	Chip.click(mouseButton)
end

Chip.talentCover = TalentClick.new(Chip.TALENT_BUTTON_NAME, { "LeftButtonUp", "RightButtonUp" }, {
	PostClick = function(self, mouseButton)
		Chip.onTalentCoverPostClick(self, mouseButton)
	end,
	OnEnter = function()
		Chip.onEnter(Chip.frame)
	end,
	OnLeave = function()
		Chip.onLeave()
	end,
	OnDragStart = function()
		Chip.onDragStart(Chip.frame)
	end,
	OnDragStop = function()
		Chip.onDragStop(Chip.frame)
	end,
})

--- Whether the cover belongs over the chip now: the chip is up, takes the mouse, its line is a point
--- to spend with the talent window not open (a click never closes it), and combat is not on.
local function talentCoverWanted()
	local frame, model = Chip.frame, Chip.current
	return frame ~= nil and frame:IsShown() == true and model ~= nil and model.talent == true
		and model.mouse ~= false and not (Chip.combat or Chip.disabled or Chip.inCinematic)
end

--- Put the cover over the chip, or take it away (TalentClick.sync: never in the lockdown).
function Chip.syncTalentButton()
	local target = talentCoverWanted() and Chip.frame or nil
	return TalentClick.sync(Chip.talentCover, target, Chip.TALENT_BUTTON_RAISE)
end

local function onEvent(_, event)
	if event == "CINEMATIC_START" then
		Chip.inCinematic = true
		if Chip.frame ~= nil then
			Chip.frame:Hide()
		end
		Chip.syncTalentButton()
	elseif event == "CINEMATIC_STOP" then
		Chip.inCinematic = false
		Chip.refresh()
	end
end

-- The frame ------------------------------------------------------------------------------------

--- A font string on the client's own font at `size`, coloured `hexKey`.
local function sizedText(parent, size, hexKey)
	local text = Theme.fontString(parent, "OVERLAY", "normal")
	local fontPath, _, fontFlags = text:GetFont()
	if fontPath ~= nil then
		text:SetFont(fontPath, size, fontFlags)
	end
	text:SetTextColor(Theme.rgb(Theme.HEX[hexKey]))
	return text
end

--- Built on first need, never at load: the raised ground with the hairline edge, the 22 px crest in
--- its ring, and the line.
function Chip.ensure()
	if Chip.frame ~= nil then
		return Chip.frame
	end
	local S = Theme.SIZES
	local frame = CreateFrame("Button", Chip.FRAME_NAME, UIParent)
	Theme.noteFrame("Chip", "login")
	frame:SetSize(S.chip, S.chip)
	frame:SetFrameStrata("MEDIUM")
	-- Above the client's own buff icons, which sit in the same strata beside the minimap.
	if type(frame.SetFrameLevel) == "function" then
		frame:SetFrameLevel(Chip.FRAME_LEVEL)
	end
	frame:SetMovable(true)
	frame:SetClampedToScreen(true)
	frame:EnableMouse(true)
	frame:RegisterForClicks("LeftButtonUp", "RightButtonUp")
	frame:RegisterForDrag("LeftButton")
	Theme.texture(frame, "BACKGROUND", "raised"):SetAllPoints(frame)
	Theme.outline(frame, S.border, "border")
	Chip.crest, Chip.ring = Theme.buildCrest(frame, S.chipCrest)
	Chip.crest:SetPoint("LEFT", frame, "LEFT", S.chipCrestInset, 0)
	local text = sizedText(frame, S.chipText, "body")
	text:SetJustifyH("LEFT")
	if type(text.SetWordWrap) == "function" then
		text:SetWordWrap(false)
	end
	text:SetPoint("LEFT", Chip.crest, "RIGHT", S.chipTextGap, 0)
	Chip.text = text
	-- "+N", muted and apart from the line, so it never reads as part of the answer.
	local more = sizedText(frame, S.chipText, "muted")
	more:SetPoint("LEFT", text, "RIGHT", S.chipMoreGap, 0)
	Chip.moreCount = more
	frame:SetScript("OnClick", Chip.onClick)
	frame:SetScript("OnEnter", Chip.onEnter)
	frame:SetScript("OnLeave", Chip.onLeave)
	frame:SetScript("OnDragStart", Chip.onDragStart)
	frame:SetScript("OnDragStop", Chip.onDragStop)
	frame:SetScript("OnEvent", onEvent)
	Theme.registerEvent(frame, "CINEMATIC_START")
	Theme.registerEvent(frame, "CINEMATIC_STOP")
	Chip.frame = frame
	Chip.place()
	return frame
end

-- The first-run pulse --------------------------------------------------------------------------

--- Whether the first-run pulse plays now: the chip never clicked or hovered on this account
--- (`db.chipSeenAt` unset), fewer than PULSE_LOGINS logins pulsed already, not this session yet, not in
--- combat and the chip on screen. Pure.
function Chip.pulseDue(db, inCombat, visible, playedThisSession)
	if type(db) ~= "table" or db.chipSeenAt ~= nil or playedThisSession or inCombat or not visible then
		return false
	end
	return (tonumber(db.chipPulses) or 0) < Chip.PULSE_LOGINS
end

--- The chip was clicked or hovered: the first-run pulse stops now and never plays again on the account.
function Chip.markSeen()
	if type(ForeverSixtyDB) == "table" and ForeverSixtyDB.chipSeenAt == nil then
		ForeverSixtyDB.chipSeenAt = type(time) == "function" and time() or 0
	end
	Chip.stopPulse()
end

--- The gold ring round the whole chip, breathing through one AnimationGroup of 2 x PULSE_CYCLES ordered
--- fades, which ends on its own and hides the ring: no OnUpdate, no timer. Built the first time it
--- plays; nil on a client without animations, which shows no ring at all.
local function pulseRing(frame)
	if Chip.pulseFrame ~= nil then
		return Chip.pulseFrame
	end
	if type(frame.CreateAnimationGroup) ~= "function" then
		return nil
	end
	local S = Theme.SIZES
	local ring = CreateFrame("Frame", nil, frame)
	Theme.noteFrame("Chip first-run pulse", "the first login before the chip is clicked")
	ring:SetPoint("TOPLEFT", frame, "TOPLEFT", -S.chipRingInset, S.chipRingInset)
	ring:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", S.chipRingInset, -S.chipRingInset)
	Theme.outline(ring, S.chipRingThickness, "clientGold")
	local group = type(ring.CreateAnimationGroup) == "function" and ring:CreateAnimationGroup() or nil
	if type(group) ~= "table" or type(group.CreateAnimation) ~= "function" then
		return nil
	end
	for step = 1, Chip.PULSE_CYCLES * 2 do
		local fade = group:CreateAnimation("Alpha")
		local rising = step % 2 == 1
		fade:SetFromAlpha(rising and Chip.PULSE_LOW or 1)
		fade:SetToAlpha(rising and 1 or Chip.PULSE_LOW)
		fade:SetDuration(Chip.PULSE_SECONDS / 2)
		fade:SetOrder(step)
		if type(fade.SetSmoothing) == "function" then
			fade:SetSmoothing("IN_OUT")
		end
	end
	group:SetScript("OnFinished", function()
		ring:Hide()
	end)
	ring:Hide()
	Chip.pulseFrame, Chip.pulseGroup = ring, group
	return ring
end

--- Plays the first-run pulse once this session when it is due (Chip.pulseDue), counting the login; it
--- ends by itself after PULSE_CYCLES breaths. True when it started.
function Chip.startPulse(model)
	local db = ForeverSixtyDB
	if Chip.frame == nil or not Chip.pulseDue(db, model ~= nil and model.alpha < 1, Chip.frame:IsShown(),
		Chip.pulsedThisSession) then
		return false
	end
	local ring = pulseRing(Chip.frame)
	if ring == nil then
		return false
	end
	Chip.pulsedThisSession = true
	db.chipPulses = (tonumber(db.chipPulses) or 0) + 1
	ring:Show()
	Chip.pulseGroup:Play()
	return true
end

--- Stops the first-run pulse and hides its ring (a click, a hover, combat or the chip hidden).
function Chip.stopPulse()
	if Chip.pulseGroup ~= nil then
		Chip.pulseGroup:Stop()
	end
	if Chip.pulseFrame ~= nil then
		Chip.pulseFrame:Hide()
	end
end

-- Drawing --------------------------------------------------------------------------------------

--- The line the chip shows for `model` inside `room` as `measure` measures it: the line's figure is
--- never cut (an item name gives way first). Also answers whether the line was cut.
function Chip.shownText(model, measure, room)
	local fitted = Chip.fit(model.line, measure, room, model.line, Chip.figureOf(model.line))
	return fitted, fitted ~= model.line
end

--- The muted count after the line: "+2", or "" with nothing else waiting.
function Chip.moreText(model)
	return model.more > 0 and string.format(L.chipMore, model.more) or ""
end

local function draw(model)
	local S = Theme.SIZES
	local frame = Chip.ensure()
	Theme.paintCrest(Chip.crest, Chip.ring, Chip.classToken())
	local count = Chip.moreText(model)
	Chip.moreCount:SetText(count)
	local countWidth = count ~= "" and math.ceil(Chip.moreCount:GetStringWidth() or 0) + S.chipMoreGap or 0
	local room = S.chipWide - S.chipTextInset - countWidth
	local shown, shortened = Chip.shownText(model, Theme.measurer(Chip.text), room)
	Chip.text:SetText(shown)
	Chip.text:SetTextColor(Theme.rgb(Theme.HEX[model.muted and "muted" or "body"]))
	local textWidth = Chip.text:GetStringWidth()
	local width = type(textWidth) == "number"
		and math.min(S.chipWide, math.ceil(textWidth) + countWidth + S.chipTextInset) or S.chipWide
	-- The chip grows leftwards from its top-right anchor.
	frame:SetWidth(width)
	frame:SetHeight(S.chip)
	-- The default spot is read again now the chip has its size: under the buff row where the buffs are
	-- in the way (Chip.place).
	if not Chip.savedSpotUsable(Prefs.get("chip", "point"), Prefs.get("chip", "x"), Prefs.get("chip", "y")) then
		Chip.place()
	end
	Chip.lastDraw = { width = width, shortened = shortened }
	frame:SetAlpha(model.alpha)
	frame:EnableMouse(model.mouse)
	frame:Show()
end

--- Redraw the chip (`data`, the shipped data, kept for the spec choice).
function Chip.refresh(data)
	if data ~= nil then
		Chip.data = data
	end
	local model = Chip.model(Chip.inputs())
	Chip.current = model
	if Chip.disabled then
		return model
	end
	if not model.visible or Chip.inCinematic then
		Chip.stopPulse()
		if Chip.frame ~= nil then
			Chip.frame:Hide()
		end
		Chip.syncTalentButton()
		return model
	end
	draw(model)
	Chip.startPulse(model)
	Chip.selfCheck()
	Chip.syncTalentButton()
	return model
end

--- Whether the frame's edges the client has resolved lie inside the screen. An edge it has not
--- resolved yet (nil, before the first layout) counts as inside.
function Chip.onScreen(frame)
	local left, right = frame:GetLeft(), frame:GetRight()
	local top = frame:GetTop()
	local parent = _G.UIParent
	local parentRight = type(parent) == "table" and type(parent.GetRight) == "function" and parent:GetRight() or nil
	local parentTop = type(parent) == "table" and type(parent.GetTop) == "function" and parent:GetTop() or nil
	if type(left) == "number" and left < 0 then
		return false
	end
	if type(right) == "number" and type(parentRight) == "number" and right > parentRight + 1 then
		return false
	end
	if type(top) == "number" and type(parentTop) == "number" and top > parentTop + 1 then
		return false
	end
	return true
end

--- The render check: the line fits the widest chip and the chip is on screen. A dragged spot off
--- screen goes back beside the minimap once; a spot under the buffs off screen goes back beside the
--- minimap for the session; past that the chip turns itself off for the session and says why in
--- /fs diag. True when it passes or there is nothing shown.
function Chip.selfCheck()
	local frame = Chip.frame
	if Chip.disabled or frame == nil or not frame:IsShown() then
		return not Chip.disabled
	end
	local S = Theme.SIZES
	local width = Chip.text:GetStringWidth()
	local fitsWidth = type(width) ~= "number" or width <= S.chipWide - S.chipTextInset
	local onScreen = Chip.onScreen(frame)
	if fitsWidth and not onScreen and Chip.savedSpotUsable(Prefs.get("chip", "point"), Prefs.get("chip", "x"),
		Prefs.get("chip", "y")) then
		Prefs.set("chip", "point", "")
		Chip.place()
		Theme.noteInfo(L.diagChipMoved)
		onScreen = Chip.onScreen(frame)
	end
	if fitsWidth and not onScreen and not Chip.noBuffMove then
		Chip.noBuffMove = true
		Chip.place()
		onScreen = Chip.onScreen(frame)
	end
	if fitsWidth and onScreen then
		return true
	end
	Chip.disabled = true
	frame:Hide()
	Theme.note(L.diagChipOff)
	return false
end

--- Why a hidden chip is hidden, for /fs diag: in combat, during a cinematic, or switched off.
function Chip.hiddenLine(hiddenBy, inCinematic)
	if inCinematic then
		return L.diagChipCinematic
	end
	if hiddenBy == "combat" then
		return L.diagChipCombat
	end
	return L.diagChipHidden
end

--- /fs diag's line for the chip: shown and whether its line was cut, or why it is hidden; nil before it
--- has drawn, or once it has turned itself off (that is already a /fs diag note).
function Chip.diagLine()
	if Chip.disabled or Chip.frame == nil then
		return nil
	end
	if not Chip.frame:IsShown() or Chip.lastDraw == nil then
		return Chip.hiddenLine(Chip.current ~= nil and Chip.current.hiddenBy or nil, Chip.inCinematic)
	end
	return Chip.lastDraw.shortened and L.diagChipShortened or L.diagChipShown
end

ns.Chip = Chip
return Chip
