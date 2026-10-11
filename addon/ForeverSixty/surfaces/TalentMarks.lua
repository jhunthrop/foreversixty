-- addon/ForeverSixty/surfaces/TalentMarks.lua
-- The talent marks (spec section 4.2): the plan drawn on the client's own
-- talent window.
--
-- Three marks, all on the window TalentGlow.watch hooks: the pulsing gold
-- ring on the next node (TalentGlow's, drawn first), a small order number
-- on every node the plan touches, and a 24 px strip in the window's
-- bottom-left corner with the build, the plan's count, the level the next
-- point arrives at and the points spent off plan. A point off plan turns
-- that count warning-coloured; nothing pops up and nothing sounds.
--
-- Pure rules, thin frames: every number comes from Progress (Progress.cells
-- for the badges, Progress.build for the strip); TalentMarks.stripText and
-- TalentMarks.badgeState turn them into what is drawn. Frames are built on
-- first need, never at load, and the badges are pooled: reopening the
-- window makes no new frames.
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
--- How far above its host the strip's frame sits, in the strata above the host's.
TalentMarks.STRIP_LEVEL = 10
--- The furthest right the strip reaches, as a share of the window's width,
--- when the window's Apply Changes button cannot be found.
TalentMarks.STRIP_MAX_SHARE = 0.45

--- Seconds after the show each deferred pass runs: the next frame, then
--- half a second later when the first one still missed.
TalentMarks.DEFERRED = { 0, 0.5 }

--- The client's frame strata, bottom to top.
TalentMarks.STRATA = {
	"BACKGROUND", "LOW", "MEDIUM", "HIGH", "DIALOG", "FULLSCREEN", "FULLSCREEN_DIALOG", "TOOLTIP",
}

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

--- The build's own name, or "<Class> build" for one that carries none.
function TalentMarks.buildName(build)
	if type(build.name) == "string" and build.name ~= "" then
		return build.name
	end
	return string.format(L.followBuildName, Talents.titleCase(build.classSlug))
end

--- The strip's line for a build and its Progress.build, and whether its
--- off-plan count is in the warning colour. The last " · " segment of the
--- Locale line is the count; it alone is wrapped when a point is off plan.
--- nil without a build: the strip is absent.
function TalentMarks.stripText(build, progress)
	if build == nil or progress == nil then
		return nil, false
	end
	local format
	if progress.point == nil then
		format = L.talentStripDone
	elseif progress.nextLevel ~= nil then
		format = L.talentStrip
	else
		format = L.talentStripNoNext
	end
	local head, tail = format:match("^(.*) · (.-)$")
	local offPlan = progress.offPlan or 0
	local count = string.format(tail, offPlan)
	local warned = offPlan > 0
	if warned then
		count = Theme.inlineColor("warning", count)
	end
	-- The head takes the name, the count and (when it has one) the level;
	-- string.format ignores the level where the line has no place for it.
	local text = string.format(head, TalentMarks.buildName(build), progress.matched, progress.total,
		progress.nextLevel)
	return text .. " · " .. count, warned
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

--- A badge's width for digits `digitsWidth` wide: the 14 px square, or the
--- digits plus orderBadgePad when they are wider ("51" is).
function TalentMarks.badgeWidth(digitsWidth)
	local S = Theme.SIZES
	if type(digitsWidth) ~= "number" then
		return S.orderBadge
	end
	return math.max(S.orderBadge, math.ceil(digitsWidth) + S.orderBadgePad)
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
		cap = windowWidth * TalentMarks.STRIP_MAX_SHARE - S.talentStripInset
	end
	if cap ~= nil then
		width = math.min(width, math.floor(cap))
	end
	return math.max(0, width)
end

--- The strata one step above `strata` (the top stays the top), or nil for
--- one this list does not name.
function TalentMarks.strataAbove(strata)
	local list = TalentMarks.STRATA
	for index, name in ipairs(list) do
		if name == strata then
			return list[math.min(index + 1, #list)]
		end
	end
	return nil
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

--- A font string on the client's own font at `size`, coloured `hexKey`.
local function sizedText(parent, size, hexKey)
	local text = Theme.fontString(parent, "OVERLAY", "normal")
	local fontPath, _, fontFlags = text:GetFont()
	if fontPath ~= nil then
		text:SetFont(fontPath, size, fontFlags)
	end
	text:SetTextColor(Theme.rgb(Theme.HEX[hexKey]))
	if type(text.SetShadowColor) == "function" then
		text:SetShadowColor(0, 0, 0, 1)
		text:SetShadowOffset(1, -1)
	end
	return text
end

--- The width a font string's text takes, unclipped where the client can say.
local function textWidth(text)
	return ask(text, "GetUnboundedStringWidth") or ask(text, "GetStringWidth")
end

--- One badge: its own frame (raised over the node's art and the ring in
--- placeBadge), a dark ground with a 1 px border and the digits centred.
local function newBadge(parent)
	local S = Theme.SIZES
	local badge = CreateFrame("Frame", nil, parent)
	badge:SetSize(S.orderBadge, S.orderBadge)
	local ground = Theme.texture(badge, "BACKGROUND", "background")
	ground:SetAllPoints(badge)
	Theme.outline(badge, S.border, "border")
	local digits = sizedText(badge, S.orderBadgeText, "gold")
	digits:SetPoint("CENTER", badge, "CENTER", 0, 0)
	badge.digits = digits
	badge:Hide()
	return badge
end

local function hideBadges()
	for _, badge in ipairs(TalentMarks.badges) do
		badge:Hide()
	end
end

--- Put the pool's `index`th badge on `button` for `cell`, building it only
--- the first time the pool reaches that size. It grows leftwards from its
--- top-right anchor to fit two digits.
local function placeBadge(index, button, cell)
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
	local state = TalentMarks.badgeState(cell)
	badge.digits:SetTextColor(Theme.rgb(Theme.HEX[state == "taken" and "muted" or "gold"]))
	badge:Show()
end

--- A badge for every cell whose node button the trait mapping finds
--- (TalentGlow.traitButton, the one lookup that does not depend on the open
--- tab, through the refresh's cache). Records how many found none, and
--- whether this pass may latch, for TalentMarks.selfCheck.
local function placeBadges(frame, frameName, cells, cache, final)
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
			placeBadge(used, button, cell)
		end
	end
	TalentMarks.placed = { missing = missing, total = #cells, final = final }
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
--- else talentStripInset in from the window's bottom-left corner.
local function anchorStrip(strip, frame)
	local S = Theme.SIZES
	strip:ClearAllPoints()
	if isFrame(frame.Inset) then
		strip:SetPoint("BOTTOMLEFT", frame.Inset, "BOTTOMLEFT", 0, 0)
		TalentMarks.stripAnchor = "inset"
	else
		strip:SetPoint("BOTTOMLEFT", frame, "BOTTOMLEFT", S.talentStripInset, S.talentStripInset)
		TalentMarks.stripAnchor = "edge"
	end
	strip:SetHeight(S.talentStrip)
end

--- One strata step above the host, and STRIP_LEVEL above it there, set on
--- every draw: SetParent hands a frame its new parent's strata.
local function raiseStrip(strip, host)
	local above = TalentMarks.strataAbove(ask(host, "GetFrameStrata"))
	if above ~= nil and type(strip.SetFrameStrata) == "function" then
		strip:SetFrameStrata(above)
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
		local label = Widgets.label(strip, "", "body", "small")
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
		anchorStrip(strip, frame)
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

--- Set the strip's line and fit its width: the text, kept clear of the
--- window's Apply Changes button. The label is measured unconstrained
--- (width 0 sizes a font string to its text), then held to the pill's
--- inside so a capped line is cut by the client rather than spilling.
local function drawStrip(frame, text)
	local S = Theme.SIZES
	local strip, host = ensureStrip(frame)
	local label = strip.label
	label:SetWidth(0)
	label:SetText(text)
	local talents = isFrame(frame.TalentsFrame) and frame.TalentsFrame or nil
	local apply = talents ~= nil and talents.ApplyButton or nil
	local applyLeft = isFrame(apply) and ask(apply, "GetLeft") or nil
	local frameLeft = ask(frame, "GetLeft")
	local stripLeft = ask(strip, "GetLeft") or (frameLeft ~= nil and frameLeft + S.talentStripInset or nil)
	local width = TalentMarks.stripWidth(textWidth(label), stripLeft, applyLeft, ask(frame, "GetWidth"))
	strip:SetWidth(width)
	label:SetWidth(math.max(1, width - 2 * S.talentStripPadding))
	strip:Show()
	return strip, host
end

--- Where the strip landed, once per session, for /fs diag (a finding, not
--- a fault): the anchor taken, its host, rect, strata, level, whether the
--- client calls it visible, and the talents frame's own level.
local function noteStrip(frame, strip, host)
	if TalentMarks.notedStrip then
		return
	end
	TalentMarks.notedStrip = true
	local function number(value)
		return type(value) == "number" and string.format("%.0f", value) or tostring(value)
	end
	local talents = isFrame(frame.TalentsFrame) and frame.TalentsFrame or nil
	Theme.noteInfo(string.format(L.diagTalentStrip, tostring(TalentMarks.stripAnchor),
		host == talents and "TalentsFrame" or tostring(ask(host, "GetName")),
		number(ask(strip, "GetLeft")), number(ask(strip, "GetBottom")),
		number(ask(strip, "GetWidth")), number(ask(strip, "GetHeight")),
		tostring(ask(strip, "GetFrameStrata")), number(ask(strip, "GetFrameLevel")),
		tostring(ask(strip, "IsVisible")), talents ~= nil and number(ask(talents, "GetFrameLevel")) or "none"))
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
--- then the strip, sharing one lookup per node. `pass` is the show's pass
--- (0 at once, 1 and 2 deferred), nil for an event's refresh; only the
--- final one may record a miss. A closed talent window is a no-op; with no
--- build loaded the strip and the badges are absent. Badge numbers are the
--- plan's own order indexes, fixed by the plan; only their colour and the
--- ring move. TalentMarks.lastPass says what the pass found, for runPass.
function TalentMarks.refresh(data, pass)
	TalentMarks.lastPass = nil
	local final = TalentMarks.isFinalPass(pass, TalentMarks.settled)
	local cache = {}
	local ring = TalentGlow.refresh(data, { quiet = not final, cache = cache })
	local frame, frameName = TalentGlow.openFrame()
	if frame == nil then
		return nil
	end
	hideBadges()
	local build = Follow.build
	if build == nil then
		TalentMarks.placed = nil
		hideStrip()
		return nil
	end
	local ranks = Talents.readRanks(data)
	placeBadges(frame, frameName, Progress.cells(build, ranks, data), cache, final)
	local _, decision = TalentMarks.selfCheck()
	local text = TalentMarks.stripText(build, Progress.build(build, ranks, Chip.level()))
	local strip, host = drawStrip(frame, text)
	if pass ~= 0 then
		-- After the show's own pass, when the client has laid the strip out.
		noteStrip(frame, strip, host)
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

ns.TalentMarks = TalentMarks
return TalentMarks
