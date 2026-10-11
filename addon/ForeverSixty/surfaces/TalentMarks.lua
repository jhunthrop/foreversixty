-- addon/ForeverSixty/surfaces/TalentMarks.lua
-- The talent marks (spec section 4.2): the plan drawn on the client's own
-- talent window.
--
-- Three marks, all on the window TalentGlow.watch hooks: the pulsing gold
-- ring on the next node (TalentGlow's, drawn first), a small order number
-- on every node the plan touches, and a 24 px strip along the window's
-- bottom edge with the build, the plan's count, the level the next point
-- arrives at and the points spent off plan. A point off plan turns that
-- count warning-coloured; nothing pops up and nothing sounds.
--
-- Pure rules, thin frames: every number comes from Progress (Progress.cells
-- for the badges, Progress.build for the strip); TalentMarks.stripText and
-- TalentMarks.badgeState turn them into what is drawn. Frames are built on
-- first need, never at load, and the badges are pooled: reopening the
-- window makes no new frames.
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
--- How far above the talent window the strip's frame sits, as the ring's.
TalentMarks.STRIP_LEVEL = 10

--- The badge pool, reused across refreshes; badgesOff once the render check
--- turned them off for the session; placed is the last placement's count.
TalentMarks.badges = {}
TalentMarks.badgesOff = false
TalentMarks.placed = nil

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

-- Frames ----------------------------------------------------------------------

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

--- One badge: its own frame (raised over the node's art and the ring in
--- placeBadge), a dark square with a 1 px border and the digits centred.
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
--- the first time the pool reaches that size.
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
	local state = TalentMarks.badgeState(cell)
	badge.digits:SetTextColor(Theme.rgb(Theme.HEX[state == "taken" and "muted" or "gold"]))
	badge:Show()
end

--- A badge for every cell whose node button the trait mapping finds
--- (TalentGlow.traitButton, the one lookup that does not depend on the open
--- tab). Records how many found none for TalentMarks.selfCheck.
local function placeBadges(frame, frameName, cells)
	TalentMarks.placed = nil
	if TalentMarks.badgesOff or TalentMarks.CLASSIC_FRAMES[frameName] then
		return
	end
	local missing, used = 0, 0
	for _, cell in ipairs(cells) do
		local button = cell.node ~= nil and TalentGlow.traitButton(frame, cell.node) or nil
		if button == nil then
			missing = missing + 1
		else
			used = used + 1
			placeBadge(used, button, cell)
		end
	end
	TalentMarks.placed = { missing = missing, total = #cells }
end

--- The strip sits inside the window's border: on the window's own inset
--- frame when it has one, else talentStripInset in from its edges.
local function anchorStrip(strip, frame)
	local S = Theme.SIZES
	strip:ClearAllPoints()
	local inset = frame.Inset
	if type(inset) == "table" and type(inset.GetObjectType) == "function" then
		strip:SetPoint("BOTTOMLEFT", inset, "BOTTOMLEFT", 0, 0)
		strip:SetPoint("BOTTOMRIGHT", inset, "BOTTOMRIGHT", 0, 0)
	else
		local edge = S.talentStripInset
		strip:SetPoint("BOTTOMLEFT", frame, "BOTTOMLEFT", edge, edge)
		strip:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", -edge, edge)
	end
	strip:SetHeight(S.talentStrip)
	if type(frame.GetFrameLevel) == "function" and type(strip.SetFrameLevel) == "function" then
		strip:SetFrameLevel((frame:GetFrameLevel() or 0) + TalentMarks.STRIP_LEVEL)
	end
end

--- The strip, built the first time a window with a build opens, and moved
--- to whichever talent window is open.
local function ensureStrip(frame)
	local S = Theme.SIZES
	local strip = TalentMarks.strip
	if strip == nil then
		strip = CreateFrame("Frame", nil, frame)
		local ground = Theme.texture(strip, "BACKGROUND", "background", Theme.ALPHA.strip)
		ground:SetAllPoints(strip)
		local label = Widgets.label(strip, "", "body", "small")
		label:SetPoint("LEFT", strip, "LEFT", S.talentStripPadding, 0)
		label:SetPoint("RIGHT", strip, "RIGHT", -S.talentStripPadding, 0)
		if type(label.SetWordWrap) == "function" then
			label:SetWordWrap(false)
		end
		strip.label = label
		TalentMarks.strip = strip
		TalentMarks.stripParent = nil
	end
	if TalentMarks.stripParent ~= frame then
		if type(strip.SetParent) == "function" then
			strip:SetParent(frame)
		end
		anchorStrip(strip, frame)
		TalentMarks.stripParent = frame
	end
	return strip
end

local function hideStrip()
	if TalentMarks.strip ~= nil then
		TalentMarks.strip:Hide()
	end
end

--- The render check (spec section 2.4): every badge found a button. When
--- one did not (a cell with no trait node counts as one), every badge is
--- hidden, the badges stay off for the session and /fs diag says how many
--- were missing, once. The ring and the strip are unaffected. True when it
--- passes or nothing was placed.
function TalentMarks.selfCheck()
	if TalentMarks.badgesOff then
		return false
	end
	local placed = TalentMarks.placed
	if placed == nil or placed.missing == 0 then
		return true
	end
	TalentMarks.badgesOff = true
	hideBadges()
	Theme.note(string.format(L.diagTalentBadges, placed.missing, placed.total))
	return false
end

--- Take every mark off: the window closed.
function TalentMarks.clear()
	TalentGlow.clear()
	hideBadges()
	hideStrip()
	return nil
end

--- Redraw the marks for `data`: the ring (TalentGlow), then the badges,
--- then the strip. A closed talent window is a no-op; with no build loaded
--- the strip and the badges are absent. Badge numbers are the plan's own
--- order indexes, fixed by the plan; only their colour and the ring move.
function TalentMarks.refresh(data)
	TalentGlow.refresh(data)
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
	placeBadges(frame, frameName, Progress.cells(build, ranks, data))
	TalentMarks.selfCheck()
	local text = TalentMarks.stripText(build, Progress.build(build, ranks, Chip.level()))
	local strip = ensureStrip(frame)
	strip.label:SetText(text)
	strip:Show()
	return text
end

ns.TalentMarks = TalentMarks
return TalentMarks
