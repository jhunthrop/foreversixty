-- addon/ForeverSixty/surfaces/HubFrame.lua
-- The Forever Sixty window's chrome (design/specs/2026-11-13-addon-hub.md 3.2 to 3.4, in the site's brand:
-- 2026-11-14-addon-brand.md section 2): the earth ground, the painted header band, the tab strip as the
-- site's nav bar and its buttons, the footer plate, and the text line and hairline they are drawn with.
-- Frames only: what they show and when is surfaces/Hub.lua's, handed in as `Hub`.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local L = ns.L or require("Locale")
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")

local HubFrame = {}

function HubFrame.line(parent, hexKey, width)
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

--- A hairline in the border colour, one whole pixel thick (Theme.wholePixels).
function HubFrame.rule(parent)
	local texture = Theme.texture(parent, "ARTWORK", "border")
	texture:SetHeight(Theme.wholePixels(texture, Theme.SIZES.border))
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

--- The window's ground (spec 2.2): the painted earth tile repeated across the window over a flat earth
--- (what shows if the tile does not load), and the window's edge: black outside, the plates' warm line
--- inside it.
function HubFrame.ground(frame)
	local S = Theme.SIZES
	local flat = frame.foreverSixtyBackground
	Theme.paint(flat, "ground", 1)
	local tile = Theme.newTexture(frame, "BACKGROUND", 1)
	tile:SetAllPoints(frame)
	tile:SetTexture(Theme.MEDIA.ground, "REPEAT", "REPEAT")
	frame.foreverSixtyGround = tile
	for _, edge in ipairs(frame.foreverSixtyBorder) do
		Theme.paint(edge, "shadow")
	end
	local inset = S.border
	for _, edge in ipairs(Theme.EDGES) do
		local texture = Theme.texture(frame, "OVERLAY", "plateEdge")
		local x = edge.from:find("RIGHT") and -inset or inset
		local y = edge.from:find("BOTTOM") and inset or -inset
		local toX = edge.to:find("RIGHT") and -inset or inset
		local toY = edge.to:find("BOTTOM") and inset or -inset
		texture:SetPoint(edge.from, frame, edge.from, x, y)
		texture:SetPoint(edge.to, frame, edge.to, toX, toY)
		if edge.horizontal then
			texture:SetHeight(Theme.wholePixels(texture, S.border))
		else
			texture:SetWidth(Theme.wholePixels(texture, S.border))
		end
	end
	HubFrame.fitGround(frame, frame:GetWidth(), frame:GetHeight())
	return tile
end

--- The earth tile repeats every Theme.SIZES.groundRepeat px over a `width` by `height` window.
function HubFrame.fitGround(frame, width, height)
	local tile = frame.foreverSixtyGround
	local size = Theme.SIZES.groundRepeat
	if tile ~= nil and type(tile.SetTexCoord) == "function" then
		tile:SetTexCoord(0, (width or size) / size, 0, (height or size) / size)
	end
	return tile
end

--- A white-to-transparent fade across `texture` from `key` (its token's alpha) at the left to nothing at
--- the right, where the client draws gradients with alpha; else the texture stays hidden.
local function fadeRight(texture, key)
	if type(CreateColor) ~= "function" or type(texture.SetGradient) ~= "function" then
		texture:Hide()
		return false
	end
	Theme.paintRGB(texture, 1, 1, 1, 1)
	local r, g, b, a = Theme.rgb(Theme.HEX[key], Theme.HEX_ALPHA[key])
	local ok = pcall(texture.SetGradient, texture, "HORIZONTAL", CreateColor(r, g, b, a), CreateColor(r, g, b, 0))
	Theme.setShown({ texture }, ok)
	return ok
end

--- The header band (spec 2.3): the tab's painted art under its veil and a fade behind the crest, the
--- ringed class crest, the title in the title face over the figures line, the close X, and the
--- hairline where the band meets the strip. The band drags the window.
function HubFrame.header(frame, Hub)
	local S, G = Theme.SIZES, Hub.LAYOUT
	local bar = Theme.createFrame("Frame", nil, frame)
	bar:SetPoint("TOPLEFT", frame, "TOPLEFT", S.border, -S.border)
	bar:SetPoint("TOPRIGHT", frame, "TOPRIGHT", -S.border, -S.border)
	bar:SetHeight(G.header - S.border)
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
	header.art = Theme.newTexture(bar, "BACKGROUND", 0)
	header.art:SetAllPoints(bar)
	header.veil = Theme.texture(bar, "BACKGROUND", "veilBand")
	header.veil:SetAllPoints(bar)
	header.fade = Theme.newTexture(bar, "BORDER", 0)
	header.fade:SetPoint("TOPLEFT", bar, "TOPLEFT", 0, 0)
	header.fade:SetPoint("BOTTOMLEFT", bar, "BOTTOMLEFT", 0, 0)
	header.fade:SetWidth(S.headerFade)
	fadeRight(header.fade, "headerFade")
	header.crest, header.ring = Theme.buildCrest(bar, G.crest)
	header.crest:SetPoint("LEFT", bar, "LEFT", S.hubSide, 0)
	local textWidth = Hub.width() - S.hubSide * 2 - G.crest - G.crestGap - S.closeButton - S.gap * 4
	header.title = Theme.titleFont(HubFrame.line(bar, "strong", textWidth), S.windowTitle)
	header.title:SetText(L.addonName)
	header.title:SetPoint("BOTTOMLEFT", header.crest, "RIGHT", G.crestGap, 1)
	header.note = HubFrame.line(bar, "headerNote")
	header.note:SetTextColor(Theme.rgb(Theme.HEX.headerNote, Theme.HEX_ALPHA.headerNote))
	header.note:SetPoint("TOPLEFT", header.crest, "RIGHT", G.crestGap, -1)
	header.noteHover = buildNoteHover(header)
	local close = Widgets.closeButton(bar, function()
		Hub.close()
	end)
	close:SetPoint("RIGHT", bar, "RIGHT", -S.hubSide, 0)
	header.close = close
	header.join = Theme.texture(bar, "ARTWORK", "headerJoin")
	header.join:SetPoint("BOTTOMLEFT", bar, "BOTTOMLEFT", 0, 0)
	header.join:SetPoint("BOTTOMRIGHT", bar, "BOTTOMRIGHT", 0, 0)
	header.join:SetHeight(Theme.wholePixels(header.join, S.border))
	return header
end

--- The header's art, veil and join line for a tab's `band` (Hub.BANDS).
function HubFrame.paintBand(header, band)
	header.art:SetTexture(Theme.MEDIA.bands[band.art])
	Theme.paint(header.veil, band.veil)
	Theme.paint(header.join, band.join)
	return band
end

--- The footer plate (spec 2.5): the plates' warm line over it, Escape's hint at the left and what a
--- click or Enter does now at the right.
function HubFrame.footer(frame, Hub)
	local S, G = Theme.SIZES, Hub.LAYOUT
	local footer = {}
	footer.plate = Theme.texture(frame, "BORDER", "footer")
	footer.plate:SetPoint("BOTTOMLEFT", frame, "BOTTOMLEFT", S.border, S.border)
	footer.plate:SetPoint("TOPRIGHT", frame, "BOTTOMRIGHT", -S.border, G.footer)
	local over = HubFrame.rule(frame)
	Theme.paint(over, "plateEdge")
	over:SetPoint("BOTTOMLEFT", frame, "BOTTOMLEFT", 0, G.footer)
	over:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", 0, G.footer)
	local hex = Theme.dimHex("muted", Hub.HINT_ALPHA)
	footer.escape = HubFrame.line(frame, "muted")
	footer.escape:SetTextColor(Theme.rgb(hex))
	footer.escape:SetPoint("LEFT", frame, "BOTTOMLEFT", S.hubSide, G.footer / 2)
	footer.escape:SetText(L.windowEscClose)
	footer.enter = HubFrame.line(frame, "muted")
	footer.enter:SetTextColor(Theme.rgb(hex))
	footer.enter:SetJustifyH("RIGHT")
	footer.enter:SetPoint("RIGHT", frame, "BOTTOMRIGHT", -S.hubSide, G.footer / 2)
	footer.enter:Hide()
	return footer
end

--- The site glyph on each tab (spec section 3): Me the character, Builds the planner, Gear best in slot,
--- Rotation the sim, Guild the guild door, Officer the two figures.
HubFrame.TAB_GLYPHS = {
	me = "character", builds = "planner", gear = "bis", rotation = "sim", guild = "guild", officer = "manage",
}

--- A tab's button on the strip (Widgets.tab: the site glyph and label, the accent underline on the active
--- one), one fixed width from its glyph, its label and the slot of the dot Hub shows while the tab has
--- something waiting, so no state moves the strip.
function HubFrame.tabButton(strip, tab, Hub)
	local G = Hub.LAYOUT
	local button = Widgets.tab(strip.frame, tab.label, function()
		Hub.select(tab.id)
	end, HubFrame.TAB_GLYPHS[tab.id])
	button:SetHeight(G.strip)
	button:SetWidth(Widgets.tabWidth(button) + G.dot + G.dotGap)
	-- The dot after the label while the tab has something waiting (spec 3.3): 6 px, the client gold.
	local dot = Theme.texture(button, "OVERLAY", "clientGold")
	dot:SetSize(G.dot, G.dot)
	dot:SetPoint("LEFT", button.foreverSixtyLabel, "RIGHT", G.dotGap, 0)
	if type(dot.SetMask) == "function" then
		dot:SetMask(Theme.MEDIA.minimapMask)
	end
	dot:Hide()
	button.foreverSixtyDot = dot
	strip.buttons[tab.id] = button
	return button
end

--- The tab strip under the header as the site's nav bar (spec 2.4): the nav bar's ground, the seal and the
--- wordmark at its left, the warm line under it. Hub shows it only while two tabs or more do.
function HubFrame.strip(frame, Hub)
	local S, G = Theme.SIZES, Hub.LAYOUT
	local strip = { buttons = {} }
	strip.frame = Theme.createFrame("Frame", nil, frame)
	strip.frame:SetPoint("TOPLEFT", frame, "TOPLEFT", S.border, -G.header)
	strip.frame:SetPoint("TOPRIGHT", frame, "TOPRIGHT", -S.border, -G.header)
	strip.frame:SetHeight(G.strip)
	local ground = Theme.texture(strip.frame, "BACKGROUND", "navBar")
	ground:SetAllPoints(strip.frame)
	strip.seal = Theme.newTexture(strip.frame, "ARTWORK")
	strip.seal:SetTexture(Theme.MEDIA.seal)
	strip.seal:SetSize(G.seal, G.seal)
	strip.seal:SetPoint("LEFT", strip.frame, "LEFT", G.sealLeft, 0)
	strip.wordmark = Theme.titleFont(HubFrame.line(strip.frame, "bone"), G.wordmark)
	strip.wordmark:SetText(L.addonName)
	strip.wordmark:SetPoint("LEFT", strip.seal, "RIGHT", G.wordmarkGap, 0)
	strip.lead = G.sealLeft + G.seal + G.tabsGap
	local under = HubFrame.rule(strip.frame)
	Theme.paint(under, "plateEdge")
	under:SetPoint("BOTTOMLEFT", strip.frame, "BOTTOMLEFT", 0, 0)
	under:SetPoint("BOTTOMRIGHT", strip.frame, "BOTTOMRIGHT", 0, 0)
	return strip
end

--- Where the strip's first tab starts for the `buttons` showing: after the wordmark while the tabs still
--- fit beside it, else after the seal alone (owner ruling 2026-11-14: the wordmark only if no tab label is
--- cut). Shows or hides the wordmark; the x.
function HubFrame.stripLead(strip, buttons, Hub)
	local G = Hub.LAYOUT
	local tabs = 0
	for _, button in ipairs(buttons) do
		tabs = tabs + (button:GetWidth() or 0) + G.tabGap
	end
	local mark = type(strip.wordmark.GetStringWidth) == "function" and strip.wordmark:GetStringWidth() or 0
	local withMark = G.sealLeft + G.seal + G.wordmarkGap + math.ceil(mark or 0) + G.tabsGap
	local fits = (mark or 0) > 0 and withMark + tabs <= Hub.width() - G.sealLeft
	Theme.setShown({ strip.wordmark }, fits)
	strip.lead = fits and withMark or G.sealLeft + G.seal + G.tabsGap
	return strip.lead
end

ns.HubFrame = HubFrame
return HubFrame
