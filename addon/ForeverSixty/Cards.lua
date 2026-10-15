-- addon/ForeverSixty/Cards.lua
-- The larger pieces the surfaces are composed from: a card, a progress bar,
-- a status pill, the primary button and an icon tile.
--
-- Widgets holds the small controls; these are the ones that give a page
-- hierarchy. Like Widgets, nothing here names a template or a client
-- global: everything is flat textures and font strings through Theme, so
-- there is one code path and the specs cover it.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local Theme = ns.Theme or require("Theme")
local Widgets = ns.Widgets or require("Widgets")

local Cards = {}

--- A fraction a bar can draw: nil, negative and over-full all clamp.
function Cards.clamp(fraction)
	if type(fraction) ~= "number" or fraction ~= fraction then
		return 0
	end
	return math.max(0, math.min(1, fraction))
end

--- A thin bar: a track and a gold fill. SetValue takes 0..1.
function Cards.progressBar(parent, width, hexKey)
	local bar = CreateFrame("Frame", nil, parent)
	Theme.noteFrame("Progress bar", "when a card draws it")
	bar:SetSize(width, Theme.SIZES.progressHeight)
	local track = Theme.texture(bar, "BACKGROUND", "track")
	track:SetAllPoints(bar)
	local fill = Theme.texture(bar, "ARTWORK", hexKey or "gold")
	fill:SetPoint("TOPLEFT", bar, "TOPLEFT", 0, 0)
	fill:SetPoint("BOTTOMLEFT", bar, "BOTTOMLEFT", 0, 0)
	bar.foreverSixtyFill = fill
	bar.foreverSixtyWidth = width
	function bar:SetValue(fraction)
		local value = Cards.clamp(fraction)
		self.foreverSixtyValue = value
		if value <= 0 then
			self.foreverSixtyFill:Hide()
		else
			self.foreverSixtyFill:Show()
			self.foreverSixtyFill:SetWidth(math.max(1, math.floor(self.foreverSixtyWidth * value + 0.5)))
		end
		return value
	end
	bar:SetValue(0)
	return bar
end

--- A small rounded-looking label: muted by default, the warning colour
--- when something needs the player's attention.
function Cards.pill(parent, text, hexKey)
	local pill = CreateFrame("Frame", nil, parent)
	Theme.noteFrame("Pill", "when a card draws it")
	local label = Widgets.label(pill, text or "", hexKey or "muted", "small")
	label:SetPoint("CENTER", pill, "CENTER", 0, 0)
	pill.foreverSixtyLabel = label
	local background = Theme.texture(pill, "BACKGROUND", "raised")
	background:SetAllPoints(pill)
	Theme.outline(pill, Theme.SIZES.border, "border")
	function pill:SetText(value, colourKey)
		self.foreverSixtyLabel:SetText(value or "")
		self.foreverSixtyLabel:SetTextColor(Theme.rgb(Theme.HEX[colourKey or "muted"]))
		local width = type(self.foreverSixtyLabel.GetStringWidth) == "function"
			and self.foreverSixtyLabel:GetStringWidth() or 0
		self:SetSize(math.max(width, 1) + Theme.SIZES.padding, Theme.SIZES.pillHeight)
		return value
	end
	pill:SetText(text, hexKey)
	return pill
end

--- An icon with a hairline frame. SetIcon(nil) draws the neutral tile.
function Cards.iconTile(parent, size, path)
	local tile = CreateFrame("Frame", nil, parent)
	Theme.noteFrame("Icon tile", "when a card draws it")
	tile:SetSize(size, size)
	tile.foreverSixtyIcon = Theme.icon(tile, "ARTWORK", path, size)
	tile.foreverSixtyIcon:SetAllPoints(tile)
	tile.foreverSixtyBorder = Theme.outline(tile, Theme.SIZES.border, "border")
	function tile:SetIcon(value)
		self.foreverSixtyIcon:SetTexture(value or Theme.UNKNOWN_ICON)
		return value
	end
	return tile
end

--- The one button on a page the player is most likely there to press:
--- gold ground, dark label. Same click rule as Widgets.button.
function Cards.primaryButton(parent, text, onClick)
	local button = Theme.createFrame("Button", nil, parent)
	button:SetSize(Theme.SIZES.buttonWidth, Theme.SIZES.buttonHeight + Theme.SIZES.gap)
	local background = Theme.texture(button, "BACKGROUND", "gold")
	background:SetAllPoints(button)
	button.foreverSixtyGround = background
	local label = Widgets.label(button, text, "onGold", "small")
	label:SetPoint("CENTER", button, "CENTER", 0, 0)
	button.foreverSixtyLabel = label
	button.SetText = function(self, value)
		self.foreverSixtyLabel:SetText(value)
	end
	button.GetText = function(self)
		return self.foreverSixtyLabel:GetText()
	end
	local states = Theme.CONTROL_STATES.primary
	button:SetScript("OnEnter", function(self)
		if self:IsEnabled() then
			Theme.paint(background, states.hover)
		end
	end)
	button:SetScript("OnLeave", function()
		Theme.paint(background, states.rest)
	end)
	Theme.pressable(button, label, function()
		Theme.paint(background, states.pressed)
	end, function(_, over)
		Theme.paint(background, over and states.hover or states.rest)
	end)
	button:SetScript("OnClick", function(self, ...)
		if self:IsEnabled() then
			onClick(self, ...)
		end
	end)
	return button
end

--- A card: an eyebrow naming it, one strong line, one quieter line, an
--- optional bar, and an optional action in its bottom-left corner. The
--- whole card is the click target for `onClick`.
function Cards.card(parent, width, height, eyebrow, onClick)
	local S = Theme.SIZES
	local card = Theme.createFrame("Button", nil, parent)
	card:SetSize(width, height)
	local ground = Theme.texture(card, "BACKGROUND", "card")
	ground:SetAllPoints(card)
	card.foreverSixtyBorder = Theme.outline(card, S.border, "border")
	card.eyebrow = Widgets.label(card, eyebrow, "muted", "small")
	card.eyebrow:SetPoint("TOPLEFT", card, "TOPLEFT", S.padding, -S.padding)
	card.title = Widgets.label(card, "", "body")
	card.title:SetPoint("TOPLEFT", card.eyebrow, "BOTTOMLEFT", 0, -S.gap * 2)
	card.title:SetWidth(width - S.padding * 2)
	card.detail = Widgets.label(card, "", "muted", "small")
	card.detail:SetPoint("TOPLEFT", card.title, "BOTTOMLEFT", 0, -S.gap)
	card.detail:SetWidth(width - S.padding * 2)
	card.innerWidth = width - S.padding * 2
	if onClick ~= nil then
		card:SetScript("OnEnter", function()
			Theme.paint(ground, "raised")
		end)
		card:SetScript("OnLeave", function()
			Theme.paint(ground, "card")
		end)
		card:SetScript("OnClick", function(self)
			onClick(self)
		end)
	end
	return card
end

ns.Cards = Cards
return Cards
