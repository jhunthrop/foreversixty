-- addon/ForeverSixty/surfaces/HubFrame.lua
-- The Forever Sixty window's chrome (design/specs/2026-11-13-addon-hub.md 3.2 to 3.4): the header band,
-- the tab strip and its buttons, the footer, and the text line and hairline they are drawn with. Frames
-- only: what they show and when is surfaces/Hub.lua's, handed in as `Hub`.
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

--- The header band: the ringed class crest, the title in the title face over the figures line, the
--- close X; the band drags the window; a hairline under it.
function HubFrame.header(frame, Hub)
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
	header.title = Theme.titleFont(HubFrame.line(bar, "bone", textWidth), S.windowTitle)
	header.title:SetText(L.addonName)
	header.title:SetPoint("BOTTOMLEFT", header.crest, "RIGHT", G.crestGap, 1)
	header.note = HubFrame.line(bar, "muted")
	header.note:SetPoint("TOPLEFT", header.crest, "RIGHT", G.crestGap, -1)
	header.noteHover = buildNoteHover(header)
	local close = Widgets.closeButton(bar, function()
		Hub.close()
	end)
	close:SetSize(S.dialogClose, S.dialogClose)
	close:SetPoint("RIGHT", bar, "RIGHT", -S.gap * 3, 0)
	header.close = close
	local under = HubFrame.rule(bar)
	under:SetPoint("BOTTOMLEFT", bar, "BOTTOMLEFT", 0, 0)
	under:SetPoint("BOTTOMRIGHT", bar, "BOTTOMRIGHT", 0, 0)
	return header
end

--- The footer band: a hairline over it, Escape's hint at the left and what a click or Enter does now at
--- the right.
function HubFrame.footer(frame, Hub)
	local S, G = Theme.SIZES, Hub.LAYOUT
	local footer = {}
	local over = HubFrame.rule(frame)
	over:SetPoint("BOTTOMLEFT", frame, "BOTTOMLEFT", 0, G.footer)
	over:SetPoint("BOTTOMRIGHT", frame, "BOTTOMRIGHT", 0, G.footer)
	local hex = Theme.dimHex("muted", Hub.HINT_ALPHA)
	footer.escape = HubFrame.line(frame, "muted")
	footer.escape:SetTextColor(Theme.rgb(hex))
	footer.escape:SetPoint("LEFT", frame, "BOTTOMLEFT", S.padding, G.footer / 2)
	footer.escape:SetText(L.windowEscClose)
	footer.enter = HubFrame.line(frame, "muted")
	footer.enter:SetTextColor(Theme.rgb(hex))
	footer.enter:SetJustifyH("RIGHT")
	footer.enter:SetPoint("RIGHT", frame, "BOTTOMRIGHT", -S.padding, G.footer / 2)
	footer.enter:Hide()
	return footer
end

--- A tab's button on the strip (Widgets.tab: the gold underline on the active one), at least
--- Theme.SIZES.tabWidth wide, with the dot Hub shows while the tab has something waiting.
function HubFrame.tabButton(strip, tab, Hub)
	local S, G = Theme.SIZES, Hub.LAYOUT
	local button = Widgets.tab(strip.frame, tab.label, function()
		Hub.select(tab.id)
	end)
	button:SetHeight(G.strip)
	local width = type(button.foreverSixtyLabel.GetStringWidth) == "function"
		and button.foreverSixtyLabel:GetStringWidth() or 0
	button:SetWidth(math.max(S.tabWidth, (width or 0) + S.padding * 2))
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

--- The tab strip under the header, a hairline under it; Hub shows it only while two tabs or more do.
function HubFrame.strip(frame, Hub)
	local G = Hub.LAYOUT
	local strip = { buttons = {} }
	strip.frame = Theme.createFrame("Frame", nil, frame)
	strip.frame:SetPoint("TOPLEFT", frame, "TOPLEFT", 0, -G.header)
	strip.frame:SetPoint("TOPRIGHT", frame, "TOPRIGHT", 0, -G.header)
	strip.frame:SetHeight(G.strip)
	local under = HubFrame.rule(strip.frame)
	under:SetPoint("BOTTOMLEFT", strip.frame, "BOTTOMLEFT", 0, 0)
	under:SetPoint("BOTTOMRIGHT", strip.frame, "BOTTOMRIGHT", 0, 0)
	return strip
end

ns.HubFrame = HubFrame
return HubFrame
