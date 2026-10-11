-- addon/ForeverSixty/surfaces/TalentMarks.lua
-- The talent marks (spec section 4.2): the plan drawn on the client's own
-- talent window.
--
-- Three marks, all on the window TalentGlow.watch hooks: the pulsing gold
-- ring on the next node (TalentGlow's, drawn first), a numbered plate on
-- every node the plan touches (the next one inverted, the ones beyond the
-- points spendable now dimmed), and a 24 px strip in the window's
-- bottom-left corner with the build, the plan's count, what to do about
-- the next point and the points spent off build. A point off build turns
-- that count warning-coloured; nothing pops up and nothing sounds. A
-- planned node's own tooltip gains one line naming its points in the plan.
--
-- Pure rules, thin frames: every number comes from Progress (Progress.cells
-- for the badges, Progress.build for the strip); TalentMarks.stripText,
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
TalentMarks.badgesOff = false
TalentMarks.placed = nil
TalentMarks.generation = 0
TalentMarks.settled = true

-- Pure rules ------------------------------------------------------------------

--- The strip's line and whether its off-build count is in the warning
--- colour. `name` is Follow.buildName's, `progress` Progress.build's, and
--- `state` { unspent, staged }: the points waiting to be spent and whether
--- an unapplied click is staged in the window. The segments: the name; "N
--- of M points"; "complete" for a finished build, else "spend N now" while
--- points wait (the window shows the same count), else "next point at
--- level N" (none at the level cap); "N off build", muted at 0 and a
--- warning from 1; and "click Apply Changes" while a staged click is on
--- build. nil without a build: the strip is absent.
function TalentMarks.stripText(name, progress, state)
	if name == nil or progress == nil then
		return nil, false
	end
	state = state or {}
	local parts = { name, string.format(L.talentStripCount, progress.matched, progress.total) }
	if progress.point == nil then
		parts[#parts + 1] = L.talentStripDone
	elseif (state.unspent or 0) > 0 then
		parts[#parts + 1] = string.format(L.talentStripSpend, state.unspent)
	elseif progress.nextLevel ~= nil then
		parts[#parts + 1] = string.format(L.talentStripNext, progress.nextLevel)
	end
	local offBuild = progress.offPlan or 0
	local warned = offBuild > 0
	parts[#parts + 1] = Theme.inlineColor(warned and "warning" or "muted",
		string.format(L.talentStripOffBuild, offBuild))
	if state.staged and not warned then
		parts[#parts + 1] = L.talentStripApply
	end
	return table.concat(parts, L.talentStripSeparator), warned
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

--- How a cell's badge looks, given `reach` (Progress.reachablePoints): the
--- next cell's plate inverted (solid gold, dark digits) so it and the ring
--- read as one "do this now"; a taken cell's digits muted on the dark
--- plate; a planned one's gold, the whole badge at half alpha when its
--- number lies beyond the points the character can spend now.
function TalentMarks.badgeLook(cell, reach)
	local A = Theme.ALPHA
	local state = TalentMarks.badgeState(cell)
	if state == "next" then
		return { plate = "gold", plateAlpha = 1, digits = "onGold", alpha = 1 }
	end
	if state == "taken" then
		return { plate = "background", plateAlpha = A.plate, digits = "muted", alpha = 1 }
	end
	local beyond = type(reach) == "number" and (cell.lastIndex or 0) > reach
	return { plate = "background", plateAlpha = A.plate, digits = "gold", alpha = beyond and A.unreachable or 1 }
end

--- The line a planned node's own tooltip gains: the plan's points in it.
function TalentMarks.nodeLine(cell)
	if (cell.want or 1) > 1 then
		return string.format(L.talentNodeLine, cell.firstIndex, cell.lastIndex)
	end
	return string.format(L.talentNodeLineOne, cell.lastIndex)
end

--- /fs diag's line for the badges' render check.
function TalentMarks.badgesLine(drawn, missing)
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
		Theme.addTooltipLine(button, TalentMarks.nodeLine(cell), "gold")
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
	Theme.outline(badge, S.border, "gold")
	local digits = sized(Theme.fontString(badge, "OVERLAY", "normal"), S.orderBadgeText)
	digits:SetPoint("CENTER", badge, "CENTER", 0, 0)
	badge.plate, badge.digits = plate, digits
	badge:Hide()
	return badge
end

local function hideBadges()
	for _, badge in ipairs(TalentMarks.badges) do
		if badge.button ~= nil then
			badge.button.foreverSixtyCell = nil
		end
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
	badge.digits:SetText(tostring(cell.lastIndex))
	badge:SetWidth(TalentMarks.badgeWidth(textWidth(badge.digits)))
	local look = TalentMarks.badgeLook(cell, reach)
	Theme.paint(badge.plate, look.plate, look.plateAlpha)
	badge.digits:SetTextColor(Theme.rgb(Theme.HEX[look.digits]))
	if type(badge.digits.SetShadowColor) == "function" then
		-- Dark digits on the gold plate take no shadow; light ones do.
		badge.digits:SetShadowColor(0, 0, 0, look.digits == "onGold" and 0 or 1)
		badge.digits:SetShadowOffset(1, -1)
	end
	badge:SetAlpha(look.alpha)
	badge.button = button
	button.foreverSixtyCell = cell
	if not button.foreverSixtyTipHooked then
		button.foreverSixtyTipHooked = Theme.hookScript(button, "OnEnter", onNodeEnter)
	end
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
--- window's Apply Changes button. The label is measured unconstrained
--- (width 0 sizes a font string to its text), then held to the pill's
--- inside so a capped line is cut by the client rather than spilling.
local function drawStrip(frame, text)
	local S = Theme.SIZES
	local strip, host = ensureStrip(frame)
	local talents = isFrame(frame.TalentsFrame) and frame.TalentsFrame or nil
	local apply = talents ~= nil and isFrame(talents.ApplyButton) and talents.ApplyButton or nil
	anchorStrip(strip, frame, apply)
	local label = strip.label
	label:SetWidth(0)
	label:SetText(text)
	local frameLeft = ask(frame, "GetLeft")
	local stripLeft = ask(strip, "GetLeft") or (frameLeft ~= nil and frameLeft + stripInset() or nil)
	local width = TalentMarks.stripWidth(textWidth(label), stripLeft, ask(apply, "GetLeft"), ask(frame, "GetWidth"))
	strip:SetWidth(width)
	label:SetWidth(math.max(1, width - 2 * S.talentStripPadding))
	strip:Show()
	return strip, host
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
	return TalentMarks.badgesLine(placed.drawn, placed.missing)
end

--- Take every mark off: the window closed. A pass still pending from the
--- show does nothing.
function TalentMarks.clear()
	TalentMarks.generation = TalentMarks.generation + 1
	TalentMarks.settled = true
	TalentGlow.clear()
	hideBadges()
	hideStrip()
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
	placeBadges(frame, frameName, Progress.cells(build, ranks, data), cache, final,
		Progress.reachablePoints(level, ranks))
	if TalentMarks.placed ~= nil then
		TalentMarks.lastPlaced = TalentMarks.placed
	end
	local _, decision = TalentMarks.selfCheck()
	local text = TalentMarks.stripText(Follow.buildName(data, build), Progress.build(build, ranks, level),
		{ unspent = Progress.unspentPoints(level, ranks), staged = staged })
	local strip, host = drawStrip(frame, text)
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
