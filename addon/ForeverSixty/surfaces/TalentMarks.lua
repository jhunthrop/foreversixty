-- addon/ForeverSixty/surfaces/TalentMarks.lua
-- The talent marks (spec section 4.2): the plan drawn on the client's own
-- talent window.
--
-- Three marks, all on the window TalentGlow.watch hooks: the pulsing gold
-- ring on the next node (TalentGlow's, drawn first), a numbered plate on
-- every node the plan touches (the next one inverted, the ones beyond the
-- points spendable now dimmed), and a 24 px strip in the window's
-- bottom-left corner with the build, the plan's count, what to do about
-- the next point and the points spent off plan. A point off plan turns
-- that count warning-coloured; nothing pops up and nothing sounds. A
-- planned node's own tooltip gains one line naming its points in the plan.
--
-- Pure rules, thin frames: every number comes from Progress (Progress.cells
-- for the badges, Progress.build for the strip); TalentMarks.stripSegments,
-- TalentMarks.badgeLook and TalentMarks.nodeLine turn them into what is
-- drawn. Ranks are read with the window's staged clicks counted
-- (Talents.readStagedRanks), so the marks answer a click before Apply
-- Changes. Frames are built on first need, never at load, and the badges
-- are pooled: reopening the window makes no new frames.
--
-- The window's OnShow runs before its trait buttons exist (seen in game:
-- the first open after a reload found no button for any node, the second
-- found them all), so a show is drawn in passes: at once, on the next
-- frame, and once more half a second later if that one still missed.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")
local Follow = ns.Follow or require("Follow")
local Talents = ns.Talents or require("Talents")
local Progress = ns.Progress or require("Progress")
local TalentGlow = ns.TalentGlow or require("TalentGlow")
local Chip = ns.Chip or require("Chip")

local TalentMarks = {}

--- The classic talent windows: their buttons carry no trait node ids, so
--- the badges stay off there without a diag line (nothing is wrong); the
--- ring and the strip still draw.
TalentMarks.CLASSIC_FRAMES = { PlayerTalentFrame = true, TalentFrame = true }

--- How far above its node button a badge's frame sits: above the ring's
--- own frame (Theme.showGlow raises it 10), so the number reads over it.
TalentMarks.BADGE_LEVEL = 12
--- How far above its host the strip's frame sits, on the host's own strata:
--- another panel dragged over the window covers it (L3.2 fix round 2).
TalentMarks.STRIP_LEVEL = 1
--- The furthest right the strip reaches, as a share of the window's width,
--- when the window's Apply Changes button cannot be found.
TalentMarks.STRIP_MAX_SHARE = 0.45

--- Seconds after the show each deferred pass runs: the next frame, then
--- half a second later when the first one still missed.
TalentMarks.DEFERRED = { 0, 0.5 }

--- The badge pool, reused across refreshes; badgesOff once the render check
--- turned them off for the session; placed is the last placement's count.
--- generation moves on with every show and hide, so a deferred pass from an
--- earlier show does nothing; settled is true once a show's passes are done.
TalentMarks.badges = {}
TalentMarks.tipButtons = {}
TalentMarks.badgesOff = false
TalentMarks.placed = nil
TalentMarks.generation = 0
TalentMarks.settled = true

-- Pure rules ------------------------------------------------------------------

--- The strip's segments and whether its off-plan count is in the warning
--- colour. `name` is Follow.buildName's, `progress` Progress.build's, and
--- `state` { unspent, next }: the points waiting to be spent, and the
--- chip's label for the next talent while the open window hides the chip
--- (Chip.pointLabel). Each segment is { kind, text }: "next", leading,
--- "Next: Improved Heroic Strike (1 of 3)" when `next` is given; "name";
--- "count", "N of M points"; "action", "complete"
--- for a finished build, else "spend N now" while points wait (the window
--- shows the same count), else "next point at level N" (none at the level
--- cap); "offPlan", "N off plan", muted at 0 and a warning from 1. Nothing
--- about Apply Changes: the client's own lit button says it (fix round 3).
--- nil without a build: the strip is absent.
function TalentMarks.stripSegments(name, progress, state)
	if name == nil or progress == nil then
		return nil, false
	end
	state = state or {}
	local segments = {}
	if state.next ~= nil then
		segments[1] = { kind = "next", text = string.format(L.talentStripNextTalent, state.next) }
	end
	segments[#segments + 1] = { kind = "name", text = name }
	segments[#segments + 1] = { kind = "count",
		text = string.format(L.talentStripCount, progress.matched, progress.total) }
	local action
	if progress.point == nil then
		action = L.talentStripDone
	elseif (state.unspent or 0) > 0 then
		action = string.format(L.talentStripSpend, state.unspent)
	elseif progress.nextLevel ~= nil then
		action = string.format(L.talentStripNext, progress.nextLevel)
	end
	if action ~= nil then
		segments[#segments + 1] = { kind = "action", text = action }
	end
	local offPlan = progress.offPlan or 0
	local warned = offPlan > 0
	segments[#segments + 1] = { kind = "offPlan",
		text = Theme.inlineColor(warned and "warning" or "muted", string.format(L.talentStripOffPlan, offPlan)) }
	return segments, warned
end

local function joinSegments(segments, dropped)
	local parts = {}
	for _, segment in ipairs(segments) do
		if not dropped[segment.kind] then
			parts[#parts + 1] = segment.text
		end
	end
	return table.concat(parts, L.talentStripSeparator)
end

--- The order segments give way in when the line is wider than the pill
--- may grow: the action ("spend N now"), then the build name, then the
--- leading "Next: …". The count and the off-plan count always stay.
TalentMarks.STRIP_SACRIFICE = { "action", "name", "next" }

--- The strip's line: the widest line of whole segments that fits
--- `maxWidth` as `measure` (text -> width) reads it, never a segment cut
--- mid-word (seen in game: "click Apply Chan…"), and whether it fits. The
--- whole line when there is no room to keep to or no `measure`. When even
--- the count and the off-plan count alone are too wide, that line and
--- false: the pill then grows to fit it (TalentMarks.pillWidth).
function TalentMarks.fitStrip(segments, measure, maxWidth)
	local dropped = {}
	local text = joinSegments(segments, dropped)
	if maxWidth == nil or type(measure) ~= "function" then
		return text, true
	end
	for _, kind in ipairs(TalentMarks.STRIP_SACRIFICE) do
		local width = measure(text)
		if type(width) ~= "number" or width <= maxWidth then
			return text, true
		end
		dropped[kind] = true
		text = joinSegments(segments, dropped)
	end
	local width = measure(text)
	return text, type(width) ~= "number" or width <= maxWidth
end

--- A badge's state for one of Progress.cells' cells: "next" for the cell
--- the ring is on, "taken" once the character has every rank the plan
--- wants there, else "planned".
function TalentMarks.badgeState(cell)
	if cell.next then
		return "next"
	end
	if (cell.have or 0) >= (cell.want or 0) then
		return "taken"
	end
	return "planned"
end

--- What a cell's badge says (spec section 4.2, amended in L3.2 fix round
--- 4): the plan's point range for a multi-rank cell, "7–11", and the single
--- number for a one-point cell, "12".
function TalentMarks.badgeText(cell)
	if (cell.want or 1) > 1 then
		return string.format(L.talentBadgeRange, cell.firstIndex, cell.lastIndex)
	end
	return tostring(cell.lastIndex)
end

--- How a cell's badge looks, given `reach` (Progress.reachablePoints): the
--- next cell's plate inverted (solid gold, dark digits) so it and the ring
--- read as one "do this now"; a taken cell's digits muted on the dark
--- plate; a planned one's gold. Every plate is near-opaque (ALPHA.plate):
--- fix round 4, a bright icon washed a thinner plate's digit out. A planned
--- cell whose FIRST point lies beyond the points the character can spend
--- now dims by a darker plate (the inset tone), its edge at
--- ALPHA.unreachable and its digits at ALPHA.dimDigits (fix round 3: badge
--- "13" for points 12-13 stays bright at 12 reachable).
function TalentMarks.badgeLook(cell, reach)
	local A = Theme.ALPHA
	local state = TalentMarks.badgeState(cell)
	if state == "next" then
		return { plate = "gold", plateAlpha = 1, edgeAlpha = 1, digits = "onGold", digitsAlpha = 1 }
	end
	if state == "taken" then
		return { plate = "background", plateAlpha = A.plate, edgeAlpha = 1, digits = "muted", digitsAlpha = 1 }
	end
	local first = cell.firstIndex or cell.lastIndex or 0
	if type(reach) == "number" and first > reach then
		return { plate = "inset", plateAlpha = A.plate, edgeAlpha = A.unreachable,
			digits = "gold", digitsAlpha = A.dimDigits }
	end
	return { plate = "background", plateAlpha = A.plate, edgeAlpha = 1, digits = "gold", digitsAlpha = 1 }
end

--- The line a node's own tooltip gains (after a blank one): "Forever
--- Sixty:" in gold, then the plan's ranks and points there in body colour
--- ("Your build: 5 ranks, points 7 to 11", "Your build: point 12"); for a
--- node holding a point the plan never put there ({ offPlan = true }), one
--- warning line, "Forever Sixty: not in your build".
function TalentMarks.nodeLine(cell)
	if cell.offPlan then
		return Theme.inlineColor("warning", string.format(L.talentNodeOffPlan, L.addonName))
	end
	local text
	if (cell.want or 1) > 1 then
		text = string.format(L.talentNodeLine, cell.want, cell.firstIndex, cell.lastIndex)
	else
		text = string.format(L.talentNodeLineOne, cell.lastIndex)
	end
	return Theme.inlineColor("gold", string.format(L.talentNodeLabel, L.addonName)) .. " "
		.. Theme.inlineColor("body", text)
end

--- /fs diag's line for the badges' render check: what was drawn and what
--- missed, or, once the check turned them `off`, how many found no button.
function TalentMarks.badgesLine(drawn, missing, off)
	if off then
		return string.format(L.diagTalentBadgesOff, missing)
	end
	return string.format(L.diagTalentBadgesLine, drawn, missing)
end

--- What the render check does with a placement: "ok" when every badge
--- found its button (or there are none); "wait" when every node missed,
--- which is the window not built yet, never latched or noted; "retry" for a
--- partial miss before the deferred pass is done; "latch" for a partial
--- miss that persists on it.
function TalentMarks.latchDecision(missing, total, deferredPassDone)
	if (total or 0) == 0 or (missing or 0) == 0 then
		return "ok"
	end
	if missing >= total then
		return "wait"
	end
	if not deferredPassDone then
		return "retry"
	end
	return "latch"
end

--- Seconds until the show's next pass, or nil for none: after the show's
--- own pass (0) always one on the next frame; after that one (1) a second,
--- half a second on, only when a badge or the ring still missed.
function TalentMarks.nextDelay(pass, decision, ringFound)
	if pass == 0 then
		return TalentMarks.DEFERRED[1]
	end
	if pass == 1 and (decision ~= "ok" or not ringFound) then
		return TalentMarks.DEFERRED[2]
	end
	return nil
end

--- Whether a pass is the one a partial miss latches on: the last deferred
--- pass, or an event's refresh (no pass) once the show's passes are done.
function TalentMarks.isFinalPass(pass, settled)
	if pass == nil then
		return settled == true
	end
	return pass >= #TalentMarks.DEFERRED
end

--- A badge's width for digits `digitsWidth` wide: the square plate, or the
--- digits plus orderBadgePad when they are wider ("51" is).
function TalentMarks.badgeWidth(digitsWidth)
	local S = Theme.SIZES
	if type(digitsWidth) ~= "number" then
		return S.orderBadge
	end
	return math.max(S.orderBadge, math.ceil(digitsWidth) + S.orderBadgePad)
end

--- How far in from the window's left and bottom edges the strip sits:
--- talentStripInset inside the window's border art.
local function stripInset()
	local S = Theme.SIZES
	return S.talentWindowBorder + S.talentStripInset
end

--- The strip's width: its text plus talentStripPadding each side, its right
--- edge kept talentStripGap short of the Apply Changes button (`applyLeft`,
--- with the strip's own `stripLeft`, both screen x), else at most
--- STRIP_MAX_SHARE of the window's width from the window's left edge.
function TalentMarks.stripWidth(textWidth, stripLeft, applyLeft, windowWidth)
	local S = Theme.SIZES
	local width = math.ceil(textWidth or 0) + 2 * S.talentStripPadding
	local cap
	if applyLeft ~= nil and stripLeft ~= nil then
		cap = applyLeft - S.talentStripGap - stripLeft
	elseif windowWidth ~= nil then
		cap = windowWidth * TalentMarks.STRIP_MAX_SHARE - stripInset()
	end
	if cap ~= nil then
		width = math.min(width, math.floor(cap))
	end
	return math.max(0, width)
end

--- The pill's width for a line `textWidth` wide: TalentMarks.stripWidth's
--- capped width when the line fits (`fits`, from fitStrip); its own width
--- past the cap when it does not, so the count and the off-plan count (the
--- only segments left then, and short) never spill out of the pill.
function TalentMarks.pillWidth(textWidth, fits, stripLeft, applyLeft, windowWidth)
	if fits then
		return TalentMarks.stripWidth(textWidth, stripLeft, applyLeft, windowWidth)
	end
	return math.ceil(textWidth or 0) + 2 * Theme.SIZES.talentStripPadding
end

--- The strip's bottom offset from the window's bottom edge: centred on the
--- Apply Changes button beside it (its bottom and height, screen units, with
--- the window's bottom), else stripInset() up.
function TalentMarks.stripOffsetY(frameBottom, applyBottom, applyHeight)
	if type(frameBottom) ~= "number" or type(applyBottom) ~= "number" or type(applyHeight) ~= "number" then
		return stripInset()
	end
	return math.floor(applyBottom - frameBottom + (applyHeight - Theme.SIZES.talentStrip) / 2 + 0.5)
end

-- Frames ----------------------------------------------------------------------

local function isFrame(value)
	return type(value) == "table" and type(value.GetObjectType) == "function"
end

--- A method's answer, or nil where the client has no such method or it raised.
local function ask(object, method)
	if object ~= nil and type(object[method]) == "function" then
		local ok, value = pcall(object[method], object)
		if ok then
			return value
		end
	end
	return nil
end

--- Put a font string on the client's own font at `size`.
local function sized(text, size)
	local fontPath, _, fontFlags = text:GetFont()
	if fontPath ~= nil then
		text:SetFont(fontPath, size, fontFlags)
	end
	return text
end

--- The width a font string's text takes, unclipped where the client can say.
local function textWidth(text)
	return ask(text, "GetUnboundedStringWidth") or ask(text, "GetStringWidth")
end

--- A planned node's tooltip gains its line after the client draws it. The
--- cell is the one the node's badge was last placed for; hideBadges forgets
--- it, so a node the client reuses for another talent says nothing.
local function onNodeEnter(button)
	local cell = button.foreverSixtyCell
	if cell ~= nil then
		Theme.addTooltipLine(button, TalentMarks.nodeLine(cell), "body", true)
	end
end

--- Give a node button `cell` for its tooltip line, hooking it the first
--- time; TalentMarks.tipButtons remembers it so a hide forgets the cell.
local function tipOn(button, cell)
	button.foreverSixtyCell = cell
	TalentMarks.tipButtons[#TalentMarks.tipButtons + 1] = button
	if not button.foreverSixtyTipHooked then
		button.foreverSixtyTipHooked = Theme.hookScript(button, "OnEnter", onNodeEnter)
	end
end

--- One badge: its own frame (raised over the node's art and the ring in
--- placeBadge), a plate (painted per look) with a 1 px gold edge and the
--- digits centred.
local function newBadge(parent)
	local S = Theme.SIZES
	local badge = CreateFrame("Frame", nil, parent)
	badge:SetSize(S.orderBadge, S.orderBadge)
	local plate = Theme.texture(badge, "BACKGROUND", "background", Theme.ALPHA.plate)
	plate:SetAllPoints(badge)
	badge.edges = Theme.outline(badge, S.border, "gold")
	local digits = sized(Theme.fontString(badge, "OVERLAY", "normal"), S.orderBadgeText)
	digits:SetPoint("CENTER", badge, "CENTER", 0, 0)
	badge.plate, badge.digits = plate, digits
	badge:Hide()
	return badge
end

local function hideBadges()
	for _, button in ipairs(TalentMarks.tipButtons) do
		button.foreverSixtyCell = nil
	end
	TalentMarks.tipButtons = {}
	for _, badge in ipairs(TalentMarks.badges) do
		badge:Hide()
	end
end

--- Put the pool's `index`th badge on `button` for `cell`, building it only
--- the first time the pool reaches that size. It grows leftwards from its
--- top-right anchor to fit two digits.
local function placeBadge(index, button, cell, reach)
	local S = Theme.SIZES
	local badge = TalentMarks.badges[index]
	if badge == nil then
		badge = newBadge(button)
		TalentMarks.badges[index] = badge
	elseif type(badge.SetParent) == "function" then
		badge:SetParent(button)
	end
	badge:ClearAllPoints()
	badge:SetPoint("TOPRIGHT", button, "TOPRIGHT", S.orderBadgeOffset, S.orderBadgeOffset)
	if type(button.GetFrameLevel) == "function" and type(badge.SetFrameLevel) == "function" then
		badge:SetFrameLevel((button:GetFrameLevel() or 0) + TalentMarks.BADGE_LEVEL)
	end
	badge.digits:SetText(TalentMarks.badgeText(cell))
	badge:SetWidth(TalentMarks.badgeWidth(textWidth(badge.digits)))
	local look = TalentMarks.badgeLook(cell, reach)
	Theme.paint(badge.plate, look.plate, look.plateAlpha)
	for _, edge in ipairs(badge.edges) do
		edge:SetAlpha(look.edgeAlpha)
	end
	badge.digits:SetTextColor(Theme.rgb(Theme.HEX[look.digits]))
	badge.digits:SetAlpha(look.digitsAlpha)
	if type(badge.digits.SetShadowColor) == "function" then
		-- Dark digits on the gold plate take no shadow; light ones do.
		badge.digits:SetShadowColor(0, 0, 0, look.digits == "onGold" and 0 or 1)
		badge.digits:SetShadowOffset(1, -1)
	end
	tipOn(button, cell)
	badge:Show()
end

--- A badge for every cell whose node button the trait mapping finds
--- (TalentGlow.traitButton, the one lookup that does not depend on the open
--- tab, through the refresh's cache). Records how many found none, and
--- whether this pass may latch, for TalentMarks.selfCheck.
local function placeBadges(frame, frameName, cells, cache, final, reach)
	TalentMarks.placed = nil
	if TalentMarks.badgesOff or TalentMarks.CLASSIC_FRAMES[frameName] then
		return
	end
	local missing, used = 0, 0
	for _, cell in ipairs(cells) do
		local button = cell.node ~= nil and TalentGlow.traitButton(frame, cell.node, cache) or nil
		if button == nil then
			missing = missing + 1
		else
			used = used + 1
			placeBadge(used, button, cell, reach)
		end
	end
	TalentMarks.placed = { missing = missing, total = #cells, final = final, drawn = used }
end

--- A warning line on the tooltip of every node holding a point the plan
--- never put there (Progress.offPlanCells), through the same lookup.
local function markOffPlan(frame, frameName, cells, cache)
	if TalentMarks.CLASSIC_FRAMES[frameName] then
		return
	end
	for _, cell in ipairs(cells) do
		local button = cell.node ~= nil and TalentGlow.traitButton(frame, cell.node, cache) or nil
		if button ~= nil then
			tipOn(button, { offPlan = true })
		end
	end
end

--- Where the strip hangs: the trait window's talents frame when it has one
--- (its art covers the window's own children: seen in game, a strip on the
--- window never showed), else the window.
local function stripHost(frame)
	if isFrame(frame.TalentsFrame) then
		return frame.TalentsFrame
	end
	return frame
end

--- The strip's corner: on the window's own inset frame when it has one,
--- else stripInset() in from the window's left edge, centred on the Apply
--- Changes button when it can be measured. Set on every draw, since the
--- button's place is only known once the window is laid out.
local function anchorStrip(strip, frame, apply)
	local S = Theme.SIZES
	strip:ClearAllPoints()
	if isFrame(frame.Inset) then
		strip:SetPoint("BOTTOMLEFT", frame.Inset, "BOTTOMLEFT", S.talentStripInset, S.talentStripInset)
		TalentMarks.stripAnchor = "inset"
	else
		local y = TalentMarks.stripOffsetY(ask(frame, "GetBottom"), ask(apply, "GetBottom"), ask(apply, "GetHeight"))
		strip:SetPoint("BOTTOMLEFT", frame, "BOTTOMLEFT", stripInset(), y)
		TalentMarks.stripAnchor = "edge"
	end
	strip:SetHeight(S.talentStrip)
end

--- The host's own strata, STRIP_LEVEL above it: another panel over the
--- window covers the strip as it covers the window, and it hides with it.
local function raiseStrip(strip, host)
	local strata = ask(host, "GetFrameStrata")
	if type(strata) == "string" and type(strip.SetFrameStrata) == "function" then
		strip:SetFrameStrata(strata)
	end
	local level = ask(host, "GetFrameLevel")
	if type(level) == "number" and type(strip.SetFrameLevel) == "function" then
		strip:SetFrameLevel(level + TalentMarks.STRIP_LEVEL)
	end
end

--- The strip, built the first time a window with a build opens, and moved
--- to whichever talent window is open.
local function ensureStrip(frame)
	local S = Theme.SIZES
	local host = stripHost(frame)
	local strip = TalentMarks.strip
	if strip == nil then
		strip = CreateFrame("Frame", nil, host)
		local ground = Theme.texture(strip, "BACKGROUND", "background", Theme.ALPHA.strip)
		ground:SetAllPoints(strip)
		local label = sized(Widgets.label(strip, "", "body", "normal"), S.talentStripText)
		label:SetPoint("LEFT", strip, "LEFT", S.talentStripPadding, 0)
		if type(label.SetWordWrap) == "function" then
			label:SetWordWrap(false)
		end
		strip.label = label
		TalentMarks.strip = strip
		TalentMarks.stripWindow = nil
	end
	if TalentMarks.stripWindow ~= frame then
		if type(strip.SetParent) == "function" then
			strip:SetParent(host)
		end
		TalentMarks.stripWindow = frame
	end
	raiseStrip(strip, host)
	return strip, host
end

local function hideStrip()
	if TalentMarks.strip ~= nil then
		TalentMarks.strip:Hide()
	end
end

--- Set the strip's line and fit it: anchored and sized clear of the
--- window's Apply Changes button, with whole segments only
--- (TalentMarks.fitStrip). Each candidate line is measured unconstrained
--- (width 0 sizes a font string to its text); the pill then fits the line
--- that was kept.
local function drawStrip(frame, segments)
	local S = Theme.SIZES
	local strip, host = ensureStrip(frame)
	local talents = isFrame(frame.TalentsFrame) and frame.TalentsFrame or nil
	local apply = talents ~= nil and isFrame(talents.ApplyButton) and talents.ApplyButton or nil
	anchorStrip(strip, frame, apply)
	local label = strip.label
	label:SetWidth(0)
	local frameLeft = ask(frame, "GetLeft")
	local stripLeft = ask(strip, "GetLeft") or (frameLeft ~= nil and frameLeft + stripInset() or nil)
	local applyLeft, frameWidth = ask(apply, "GetLeft"), ask(frame, "GetWidth")
	-- The widest pill the corner allows, less its padding, is the text's room.
	local room = TalentMarks.stripWidth(math.huge, stripLeft, applyLeft, frameWidth)
	local maxText = room ~= math.huge and room - 2 * S.talentStripPadding or nil
	local text, fits = TalentMarks.fitStrip(segments, function(candidate)
		label:SetText(candidate)
		return textWidth(label)
	end, maxText)
	label:SetText(text)
	strip:SetWidth(TalentMarks.pillWidth(textWidth(label), fits, stripLeft, applyLeft, frameWidth))
	strip:Show()
	return strip, host, text
end

--- Where the strip landed and which rank fields the client's trait nodes
--- fill, once per session, for /fs diag verbose: the anchor taken, the
--- strip's host, rect, strata, level, whether the client calls it visible,
--- and the talents frame's own level.
local function noteFindings(frame, strip, host, staged)
	if TalentMarks.notedStrip then
		return
	end
	TalentMarks.notedStrip = true
	local function number(value)
		return type(value) == "number" and string.format("%.0f", value) or tostring(value)
	end
	local talents = isFrame(frame.TalentsFrame) and frame.TalentsFrame or nil
	Theme.noteVerbose(string.format(L.diagTalentStrip, tostring(TalentMarks.stripAnchor),
		host == talents and "TalentsFrame" or tostring(ask(host, "GetName")),
		number(ask(strip, "GetLeft")), number(ask(strip, "GetBottom")),
		number(ask(strip, "GetWidth")), number(ask(strip, "GetHeight")),
		tostring(ask(strip, "GetFrameStrata")), number(ask(strip, "GetFrameLevel")),
		tostring(ask(strip, "IsVisible")), talents ~= nil and number(ask(talents, "GetFrameLevel")) or "none"))
	local sample = Talents.nodeSample
	if type(sample) == "table" then
		Theme.noteVerbose(string.format(L.diagTalentNodeFields, tostring(sample.activeRank),
			tostring(sample.currentRank), tostring(sample.ranksPurchased), tostring(staged)))
	end
end

--- The render check (spec section 2.4): every badge found a button. A
--- partial miss that persists on the show's deferred pass (a cell with no
--- trait node counts as a miss) hides every badge, keeps them off for the
--- session and says how many were missing in /fs diag, once; every node
--- missing is the window not built yet and is neither latched nor noted.
--- The ring and the strip are unaffected. False once the badges are off;
--- the second return is TalentMarks.latchDecision's answer.
function TalentMarks.selfCheck()
	if TalentMarks.badgesOff then
		return false, "ok"
	end
	local placed = TalentMarks.placed
	if placed == nil then
		return true, "ok"
	end
	local decision = TalentMarks.latchDecision(placed.missing, placed.total, placed.final)
	if decision ~= "latch" then
		return true, decision
	end
	TalentMarks.badgesOff = true
	hideBadges()
	Theme.note(string.format(L.diagTalentBadges, placed.missing, placed.total))
	return false, decision
end

--- /fs diag's badges line ("Talent badges: 18 drawn, 0 missing") for the
--- last placement, or nil before any.
function TalentMarks.diagLine()
	local placed = TalentMarks.lastPlaced
	if placed == nil then
		return nil
	end
	return TalentMarks.badgesLine(placed.drawn, placed.missing, TalentMarks.badgesOff)
end

--- Take every mark off: the window closed. A pass still pending from the
--- show does nothing.
function TalentMarks.clear()
	TalentMarks.generation = TalentMarks.generation + 1
	TalentMarks.settled = true
	TalentGlow.clear()
	hideBadges()
	hideStrip()
	-- The chip stood aside while the window was open; it comes back.
	Chip.refresh()
	return nil
end

--- The chip's own label for the next cell (Chip.pointLabel), while the open
--- window is what hides the chip; nil otherwise or once the build is done.
local function nextLabel(data, build, cells)
	if not Chip.hiddenByWindow() then
		return nil
	end
	for _, cell in ipairs(cells) do
		if cell.next then
			local talent = Talents.cellOf(data, build.classSlug, cell.tab, cell.tier, cell.column)
			return Chip.pointLabel(Chip.talentName(talent) or L.chipNextTalent, { have = cell.have, want = cell.want })
		end
	end
	return nil
end

--- Redraw the marks for `data`: the ring (TalentGlow), then the badges,
--- then the strip, sharing one lookup per node and one read of the ranks
--- with staged clicks counted. `pass` is the show's pass (0 at once, 1 and
--- 2 deferred), nil for an event's refresh; only the final one may record
--- a miss. A closed talent window is a no-op; with no build loaded the
--- strip and the badges are absent. Badge numbers are the plan's own order
--- indexes, fixed by the plan; only their look and the ring move.
--- TalentMarks.lastPass says what the pass found, for runPass.
function TalentMarks.refresh(data, pass)
	TalentMarks.lastPass = nil
	local final = TalentMarks.isFinalPass(pass, TalentMarks.settled)
	local frame, frameName = TalentGlow.openFrame()
	local build = Follow.build
	if frame == nil or build == nil then
		-- TalentGlow takes its ring off itself: nothing to read ranks for.
		TalentGlow.refresh(data, { quiet = not final })
		if frame ~= nil then
			hideBadges()
			TalentMarks.placed = nil
			hideStrip()
		end
		return nil
	end
	local ranks, staged = Talents.readStagedRanks(data)
	local cache = {}
	local ring = TalentGlow.refresh(data, { quiet = not final, cache = cache, ranks = ranks })
	hideBadges()
	local level = Chip.level()
	local cells = Progress.cells(build, ranks, data)
	placeBadges(frame, frameName, cells, cache, final,
		Progress.reachablePoints(level, ranks))
	markOffPlan(frame, frameName, Progress.offPlanCells(build, ranks, data), cache)
	if TalentMarks.placed ~= nil then
		TalentMarks.lastPlaced = TalentMarks.placed
	end
	local _, decision = TalentMarks.selfCheck()
	local segments = TalentMarks.stripSegments(Follow.buildName(data, build), Progress.build(build, ranks, level),
		{ unspent = Progress.unspentPoints(level, ranks), next = nextLabel(data, build, cells) })
	local strip, host, text = drawStrip(frame, segments)
	if pass ~= 0 then
		-- After the show's own pass, when the client has laid the strip out.
		noteFindings(frame, strip, host, staged)
	end
	TalentMarks.lastPass = { decision = decision, ring = ring ~= nil }
	return text
end

--- One pass of a show, then the next one TalentMarks.nextDelay asks for;
--- settled once there is none (or the client cannot wait).
function TalentMarks.runPass(data, pass, generation)
	if generation ~= TalentMarks.generation then
		return nil
	end
	TalentMarks.refresh(data, pass)
	local last = TalentMarks.lastPass
	local delay = last ~= nil and TalentMarks.nextDelay(pass, last.decision, last.ring) or nil
	local waiting = delay ~= nil and Theme.after(delay, function()
		TalentMarks.runPass(data, pass + 1, generation)
	end)
	if not waiting then
		TalentMarks.settled = true
	end
	return pass
end

--- The talent window opened (TalentGlow.watch's OnShow): draw at once,
--- then again once its trait buttons exist.
function TalentMarks.onShow(data)
	TalentMarks.generation = TalentMarks.generation + 1
	TalentMarks.settled = false
	-- The ring and the strip replace the chip while the window is open.
	Chip.refresh()
	return TalentMarks.runPass(data, 0, TalentMarks.generation)
end

--- A click in the window (TRAIT_NODE_CHANGED and its kin fire once per node
--- touched): one refresh on the next frame however many arrive, now on a
--- client that cannot wait.
function TalentMarks.requestRefresh(data)
	if TalentMarks.refreshPending then
		return false
	end
	TalentMarks.refreshPending = true
	local waiting = Theme.after(0, function()
		TalentMarks.refreshPending = false
		TalentMarks.refresh(data)
	end)
	if not waiting then
		TalentMarks.refreshPending = false
		TalentMarks.refresh(data)
	end
	return true
end

ns.TalentMarks = TalentMarks
return TalentMarks
