-- addon/ForeverSixty/Widgets.lua
-- The addon's own small widget set, built on Theme.
--
-- Two rules hold this file together. Every client-facing decision is
-- Theme's, so nothing here names a template, a font or a UI global. And
-- every widget works with no template at all: a panel is a background
-- texture plus four edge textures, a button is a panel with a label, a
-- tab is a button that recolours, a list is a fixed set of recycled rows
-- over a pure offset. The templates, where the client has them, are a
-- nicer skin over exactly the same behaviour -- never a second code path
-- the specs would have to cover twice.
local _, ns = ...
ns = type(ns) == "table" and ns or {}
local Theme = ns.Theme or require("Theme")

local Widgets = {}

--- Which slice of `total` items a list of `rows` rows sitting at `offset`
--- is looking at, and the offset clamped into range. `0, -1` for an empty
--- list, so `for index = first, last` runs zero times.
function Widgets.visibleRange(total, rows, offset)
	local highest = math.max(0, total - rows)
	local clamped = math.max(0, math.min(offset or 0, highest))
	if total == 0 then
		return 0, -1, 0
	end
	return clamped + 1, math.min(total, clamped + rows), clamped
end

function Widgets.panel(parent, width, height, name)
	local frame = CreateFrame("Frame", name, parent)
	frame:SetSize(width, height)
	local background = Theme.texture(frame, "BACKGROUND", "background", Theme.ALPHA.window)
	background:SetAllPoints(frame)
	frame.foreverSixtyBackground = background
	frame.foreverSixtyBorder = Theme.outline(frame, Theme.SIZES.border, "border")
	return frame
end

function Widgets.label(parent, text, hexKey, fontKey)
	local region = Theme.fontString(parent, "OVERLAY", fontKey or "normal")
	region:SetTextColor(Theme.rgb(Theme.HEX[hexKey or "body"]))
	region:SetJustifyH("LEFT")
	region:SetText(text or "")
	return region
end

--- Recolour a flat control while the cursor is over it. The control keeps
--- its resting colours in `foreverSixtySkin`, so a caller that changes them
--- (an active tab) is respected when the cursor leaves.
local function hoverable(control, background)
	local states = Theme.CONTROL_STATES.secondary
	control:SetScript("OnEnter", function(self)
		if self:IsEnabled() then
			Theme.paint(background, states.hover)
		end
	end)
	control:SetScript("OnLeave", function()
		Theme.paint(background, states.rest)
	end)
	return control
end

--- A button the player can turn off. The OnClick wrapper checks IsEnabled
--- itself rather than trusting the client to swallow the click, so the
--- combat lockout on the Equip buttons holds on every client. Drawn flat:
--- a raised panel, a hairline border, a gold label. `name` (optional) names the
--- frame, for a button a /run probe reaches (the character-panel tab's, L4.3).
function Widgets.button(parent, text, onClick, name)
	local button = Theme.createFrame("Button", name, parent)
	button:SetSize(Theme.SIZES.buttonWidth, Theme.SIZES.buttonHeight)
	local background = Theme.texture(button, "BACKGROUND", "raised")
	background:SetAllPoints(button)
	button.foreverSixtyGround = background
	button.foreverSixtyEdges = Theme.outline(button, Theme.SIZES.border, Theme.CONTROL_STATES.face.edge)
	local label = Widgets.label(button, "", Theme.CONTROL_STATES.face.label, "small")
	label:SetPoint("CENTER", button, "CENTER", 0, 0)
	button.foreverSixtyLabel = label
	button.SetText = function(self, value)
		self.foreverSixtyLabel:SetText(value)
	end
	button.GetText = function(self)
		return self.foreverSixtyLabel:GetText()
	end
	hoverable(button, background)
	Theme.pressable(button, label, function()
		Theme.paint(background, Theme.CONTROL_STATES.secondary.pressed)
	end, function(_, over)
		Theme.paint(background, over and Theme.CONTROL_STATES.secondary.hover or Theme.CONTROL_STATES.secondary.rest)
	end)
	button:SetText(text)
	button:SetScript("OnClick", function(self, ...)
		if self:IsEnabled() then
			onClick(self, ...)
		end
	end)
	return button
end

--- A button that must read as the addon's own on the client's art (the character panel's
--- "Forever Sixty" tab, L4.3 fix round 1, item 5): the panel fill under a 1 px gold edge, its
--- label in the chip's gold font (the client's own font at Theme.SIZES.chipText), `width` x
--- `height`; under the cursor the fill lifts to the hover ground and the edge and label to
--- the brighter gold. `name` (optional) names the frame for a /run probe.
function Widgets.edgedButton(parent, text, onClick, name, width, height)
	local S = Theme.SIZES
	local button = Theme.createFrame("Button", name, parent)
	button:SetSize(width or S.buttonWidth, height or S.buttonHeight)
	local ground = Theme.texture(button, "BACKGROUND", "background", Theme.ALPHA.window)
	ground:SetAllPoints(button)
	button.foreverSixtyGround = ground
	button.foreverSixtyEdges = Theme.outline(button, S.border, "gold")
	local label = Theme.fontString(button, "OVERLAY", "normal")
	local fontPath, _, fontFlags = label:GetFont()
	label:SetFont(fontPath or Theme.FALLBACK_FONT.path, S.chipText, fontFlags)
	label:SetTextColor(Theme.rgb(Theme.HEX.gold))
	label:SetPoint("CENTER", button, "CENTER", 0, 0)
	button.foreverSixtyLabel = label
	button.SetText = function(self, value)
		self.foreverSixtyLabel:SetText(value or "")
	end
	button.GetText = function(self)
		return self.foreverSixtyLabel:GetText()
	end
	--- Rest or lit, closed or open (Widgets.setEdgedOpen): the fill, the edge and the label
	--- together. Open (L4.3 fix round 3, item 3): the label in body over the inset ground of a
	--- pressed tab; lit while open, the label brightens and the inset ground stays.
	local function look(lit)
		local open = button.foreverSixtyOpen == true
		if open then
			Theme.paint(ground, Theme.CONTROL_STATES.tab.pressed)
		else
			Theme.paint(ground, lit and "hover" or "background", not lit and Theme.ALPHA.window or nil)
		end
		for _, edge in ipairs(button.foreverSixtyEdges) do
			Theme.paint(edge, lit and "goldHover" or "gold")
		end
		label:SetTextColor(Theme.rgb(Theme.HEX[lit and "goldHover" or (open and "body" or "gold")]))
	end
	button.foreverSixtyLook = look
	button:SetScript("OnEnter", function()
		look(true)
	end)
	button:SetScript("OnLeave", function()
		look(false)
	end)
	Theme.pressable(button, label, function()
		Theme.paint(ground, "inset")
	end, function(_, over)
		look(over)
	end)
	button:SetScript("OnClick", function(self, ...)
		if self:IsEnabled() then
			onClick(self, ...)
		end
	end)
	button:SetText(text)
	return button
end

--- An edged button's open state (what it opens is showing), drawn at once: lit when the
--- cursor is over it, else open or at rest.
function Widgets.setEdgedOpen(button, open)
	button.foreverSixtyOpen = open == true
	local over = type(button.IsMouseOver) == "function" and button:IsMouseOver() == true
	if type(button.foreverSixtyLook) == "function" then
		button.foreverSixtyLook(over)
	end
	return button.foreverSixtyOpen
end

--- A label the player can click, with no border or background of its own
--- -- a "+N more" or "Show fewer" affordance sitting inside a card,
--- muted at rest and gold under the cursor like a plain text link. Sizes
--- itself to the text's own width (GetStringWidth), so the caller only
--- ever positions it, never sizes it.
function Widgets.textButton(parent, text, onClick)
	local button = Theme.createFrame("Button", nil, parent)
	local label = Widgets.label(button, "", "muted", "small")
	label:SetPoint("LEFT", button, "LEFT", 0, 0)
	button.foreverSixtyLabel = label
	button:SetHeight(Theme.SIZES.rowHeight)
	button.SetText = function(self, value)
		self.foreverSixtyLabel:SetText(value or "")
		local width = type(label.GetStringWidth) == "function" and label:GetStringWidth() or 0
		self:SetWidth(math.max(width, 1))
	end
	button.GetText = function(self)
		return self.foreverSixtyLabel:GetText()
	end
	button:SetScript("OnEnter", function()
		label:SetTextColor(Theme.rgb(Theme.HEX.gold))
	end)
	button:SetScript("OnLeave", function()
		label:SetTextColor(Theme.rgb(Theme.HEX.muted))
	end)
	button:SetScript("OnClick", function(self, ...)
		if self:IsEnabled() then
			onClick(self, ...)
		end
	end)
	button:SetText(text)
	return button
end

--- The small X in the title bar: muted, body on a hover ground under the cursor
--- (Theme.CONTROL_STATES.close), inset while pressed.
function Widgets.closeButton(parent, onClick)
	local states = Theme.CONTROL_STATES.close
	local button = Theme.createFrame("Button", nil, parent)
	button:SetSize(Theme.SIZES.closeButton, Theme.SIZES.closeButton)
	local ground = Theme.texture(button, "BACKGROUND", states.hoverGround)
	ground:SetAllPoints(button)
	ground:Hide()
	button.foreverSixtyGround = ground
	local label = Widgets.label(button, "x", states.label)
	label:SetPoint("CENTER", button, "CENTER", 0, 1)
	button.foreverSixtyLabel = label
	button:SetScript("OnEnter", function(self)
		Theme.paint(ground, states.hoverGround)
		ground:Show()
		self.foreverSixtyLabel:SetTextColor(Theme.rgb(Theme.HEX[states.hoverLabel]))
	end)
	button:SetScript("OnLeave", function(self)
		ground:Hide()
		self.foreverSixtyLabel:SetTextColor(Theme.rgb(Theme.HEX[states.label]))
	end)
	Theme.pressable(button, label, function()
		Theme.paint(ground, states.pressed)
		ground:Show()
	end, function(_, over)
		Theme.paint(ground, states.hoverGround)
		if over then
			ground:Show()
		else
			ground:Hide()
		end
	end)
	button:SetScript("OnClick", function(self)
		onClick(self)
	end)
	return button
end

--- A Widgets.button's look set outright, at full alpha: enabled or not, its label and its
--- 1 px border in the given Theme.HEX keys. For a button that has acted and waits, whose
--- faded disabled look would not read (Settings' "Positions reset", fix round 2, item 4).
function Widgets.setButtonLook(button, enabled, labelKey, edgeKey)
	if enabled then
		button:Enable()
	else
		button:Disable()
	end
	button:SetAlpha(1)
	button.foreverSixtyLabel:SetTextColor(Theme.rgb(Theme.HEX[labelKey]))
	for _, edge in ipairs(button.foreverSixtyEdges or {}) do
		Theme.paint(edge, edgeKey)
	end
	return enabled
end

function Widgets.setEnabled(widget, enabled)
	if enabled then
		widget:Enable()
		widget:SetAlpha(1)
	else
		widget:Disable()
		widget:SetAlpha(Theme.ALPHA.disabled)
	end
	return enabled
end

--- A tab is a label over a thin underline. The underline shows on the
--- active tab only; the label goes gold with it.
function Widgets.tab(parent, text, onClick)
	local tab = Theme.createFrame("Button", nil, parent)
	tab:SetSize(Theme.SIZES.tabWidth, Theme.SIZES.tabHeight)
	local label = Widgets.label(tab, text, "muted", "small")
	label:SetPoint("CENTER", tab, "CENTER", 0, 0)
	tab.foreverSixtyLabel = label
	local underline = Theme.texture(tab, "OVERLAY", "gold")
	underline:SetPoint("BOTTOMLEFT", tab, "BOTTOMLEFT", 0, 0)
	underline:SetPoint("BOTTOMRIGHT", tab, "BOTTOMRIGHT", 0, 0)
	underline:SetHeight(Theme.SIZES.tabUnderline)
	underline:Hide()
	tab.foreverSixtyUnderline = underline
	local states = Theme.CONTROL_STATES.tab
	local ground = Theme.texture(tab, "BACKGROUND", states.pressed)
	ground:SetAllPoints(tab)
	ground:Hide()
	tab:SetScript("OnEnter", function(self)
		if not self.foreverSixtyActive then
			-- An inactive tab under the cursor: body label over a 2 px border-colour underline.
			self.foreverSixtyLabel:SetTextColor(Theme.rgb(Theme.HEX[states.hoverLabel]))
			Theme.paint(underline, states.hoverUnderline)
			underline:Show()
		end
	end)
	Theme.pressable(tab, label, function()
		ground:Show()
	end, function()
		ground:Hide()
	end)
	tab:SetScript("OnLeave", function(self)
		Widgets.setTabActive(self, self.foreverSixtyActive == true)
	end)
	tab:SetScript("OnClick", function(self)
		onClick(self)
	end)
	return tab
end

function Widgets.setTabActive(tab, active)
	tab.foreverSixtyLabel:SetTextColor(Theme.rgb(Theme.HEX[active and "gold" or "muted"]))
	if active then
		Theme.paint(tab.foreverSixtyUnderline, "gold")
		tab.foreverSixtyUnderline:Show()
	else
		tab.foreverSixtyUnderline:Hide()
	end
	tab.foreverSixtyActive = active
	return active
end

--- A read-only box is one the player copies out of: typing in it puts the
--- value straight back, so a stray keystroke cannot corrupt the export
--- string sitting selected under the cursor.
---
--- The box lives inside a field: a dark inset panel of exactly the size
--- asked for, which clips it. A multi-line edit box grows with its text
--- whatever height it is given, and the first in-game screenshot showed a
--- long export spilling over the labels above and the button below. The
--- field is what a caller positions, shows and hides (Widgets.field).
function Widgets.editBox(parent, width, height, readOnly, singleLine)
	local inset = Theme.SIZES.gap
	local field = Widgets.panel(parent, width, height)
	Theme.paint(field.foreverSixtyBackground, "inset")
	if type(field.SetClipsChildren) == "function" then
		field:SetClipsChildren(true)
	end
	local box = Theme.createFrame("EditBox", nil, field)
	box:SetPoint("TOPLEFT", field, "TOPLEFT", inset, -inset)
	box:SetSize(width - inset * 2, height - inset * 2)
	box:SetMultiLine(not singleLine)
	-- An edit box is created with auto-focus on and takes the keyboard the
	-- moment it exists. Turning auto-focus off does not give that focus
	-- back, so it is released here: found in game, where opening the window
	-- (which creates this box) left every keybind dead.
	box:SetAutoFocus(false)
	box:ClearFocus()
	box:SetTextColor(Theme.rgb(Theme.HEX.body))
	box.foreverSixtyField = field
	Theme.applyFont(box, "small")
	if readOnly then
		box:SetScript("OnTextChanged", function(self, userInput)
			if userInput and self:GetText() ~= (self.foreverSixtyValue or "") then
				self:SetText(self.foreverSixtyValue or "")
			end
		end)
	end
	Widgets.releaseFocusWhenDone(box)
	return box
end

--- The panel an edit box sits in: the thing to anchor, show and hide.
function Widgets.field(box)
	return box.foreverSixtyField
end

--- How long after Ctrl+C the box keeps focus: the client copies the
--- selection on the key press, so focus can go a moment later.
Widgets.COPY_FOCUS_RELEASE_SECONDS = 0.2

--- An edit box that holds keyboard focus swallows every keybind: found in
--- game, where opening the window left movement and action bars dead until
--- Escape. A box may hold focus only while the player is copying out of it,
--- so it gives focus back when it is hidden (the tab changed or the window
--- closed), on Escape, on Enter, and shortly after Ctrl+C.
function Widgets.releaseFocusWhenDone(box)
	local function release(self)
		self:ClearFocus()
	end
	box:SetScript("OnHide", release)
	box:SetScript("OnEscapePressed", release)
	box:SetScript("OnEnterPressed", release)
	box:SetScript("OnKeyDown", function(self, key)
		local copying = key == "C" and type(IsControlKeyDown) == "function" and IsControlKeyDown()
		if not copying then
			return
		end
		if type(C_Timer) == "table" and type(C_Timer.After) == "function" then
			C_Timer.After(Widgets.COPY_FOCUS_RELEASE_SECONDS, function()
				release(self)
			end)
		end
	end)
	return box
end

--- Put text in a box without touching keyboard focus. Drawing a tab uses
--- this; only an explicit copy uses selectText.
function Widgets.setText(box, text)
	box.foreverSixtyValue = text
	box:SetText(text)
	-- SetText leaves the cursor at the end, and a single-line box scrolls to
	-- its cursor, so a code longer than the box showed only its tail (found
	-- in game: the Export page's box looked empty of anything recognisable).
	-- The start is the part a player recognises -- "FS1:" -- so show that.
	if type(box.SetCursorPosition) == "function" then
		box:SetCursorPosition(0)
	end
	return text
end

--- Select the text and take focus so Ctrl+C copies it. Only ever called
--- from a click the player made.
function Widgets.selectText(box, text)
	Widgets.setText(box, text)
	box:HighlightText()
	box:SetFocus()
	return text
end

--- Show the game's own item tooltip for whatever `provider` names now.
--- The provider is a closure rather than an id so a recycled list row can
--- change what it is showing without rebinding its scripts.
function Widgets.attachTooltip(frame, provider)
	frame:SetScript("OnEnter", function(self)
		local itemId, link = provider()
		if itemId ~= nil or link ~= nil then
			Theme.showItemTooltip(self, itemId, link)
		end
	end)
	frame:SetScript("OnLeave", function()
		Theme.hideTooltip()
	end)
	return frame
end

--- One list row: icon, name, and a right-aligned number. Every list in
--- this addon -- talent order, planned gear, bag upgrades -- is made of
--- these, so a row looks the same wherever the player meets it.
function Widgets.itemRow(parent, width)
	local frame = CreateFrame("Frame", nil, parent)
	frame:SetSize(width, Theme.SIZES.rowHeight)
	frame:EnableMouse(true)
	-- Through Theme.icon so the client's baked-in bevel is cropped, the
	-- same as every other icon in the window.
	local icon = Theme.icon(frame, "ARTWORK", nil, Theme.SIZES.iconSize)
	icon:SetTexture(nil)
	icon:SetPoint("LEFT", frame, "LEFT", 0, 0)
	local text = Widgets.label(frame, "", "body", "small")
	text:SetPoint("LEFT", icon, "RIGHT", Theme.SIZES.gap, 0)
	local right = Widgets.label(frame, "", "muted", "small")
	right:SetJustifyH("RIGHT")
	right:SetPoint("RIGHT", frame, "RIGHT", 0, 0)
	local row = { frame = frame, icon = icon, text = text, right = right }
	Widgets.attachTooltip(frame, function()
		return row.itemId, row.link
	end)
	return row
end

--- One rotation-card priority row (docs/tenets.md's standard, replacing
--- the plain single-line label the card drew before): a number, the
--- ability's icon, its name, its rank in muted text beside the name, and
--- the condition in muted text on a second line. Hovering highlights the
--- row; the row for the ability the player just learned this level glows
--- (Theme.showGlow/hideGlow, on the level-up event) -- both are the
--- caller's to drive, this only builds the frame they act on.
function Widgets.rotationRow(parent, width)
	local S = Theme.SIZES
	local frame = CreateFrame("Button", nil, parent)
	frame:SetSize(width, S.rotationRowHeight)
	frame:EnableMouse(true)
	local highlight = Theme.texture(frame, "BACKGROUND", "hover")
	highlight:SetAllPoints(frame)
	highlight:Hide()
	frame:SetScript("OnEnter", function()
		highlight:Show()
	end)
	frame:SetScript("OnLeave", function()
		highlight:Hide()
	end)
	local number = Widgets.label(frame, "", "muted", "small")
	number:SetPoint("TOPLEFT", frame, "TOPLEFT", 0, 0)
	number:SetWidth(S.gap * 5)
	local icon = Theme.icon(frame, "ARTWORK", nil, S.iconSize)
	icon:SetPoint("LEFT", number, "RIGHT", S.gap, 0)
	local name = Widgets.label(frame, "", "body", "small")
	name:SetPoint("TOPLEFT", icon, "TOPRIGHT", S.gap, 0)
	local rank = Widgets.label(frame, "", "muted", "small")
	rank:SetPoint("LEFT", name, "RIGHT", S.gap, 0)
	local condition = Widgets.label(frame, "", "muted", "small")
	condition:SetPoint("TOPLEFT", name, "BOTTOMLEFT", 0, -S.gap)
	condition:SetWidth(math.max(1, width - S.iconSize - S.gap * 7))
	return { frame = frame, number = number, icon = icon, name = name, rank = rank, condition = condition }
end

--- A page taller than its holder scrolls with the mouse wheel. The
--- holder clips; the content frame inside it moves. Widgets.scrollable
--- returns the content frame to lay out on and a function to call once the
--- content's height is known. No scroll template: the same flat approach
--- as Widgets.list, so the specs cover it.
--- A thin scrollbar at the right edge of `region`, drawn on `parent`
--- and shown only while there is more than fits, so a list or page that
--- fits looks as it always did. `Update(shown, position)` takes the
--- fraction that fits (1 means everything) and how far down the view
--- is (0 to 1).
function Widgets.scrollbar(parent, region, inset)
	local S = Theme.SIZES
	inset = inset or -S.gap * 2
	local track = Theme.texture(parent, "ARTWORK", "border")
	track:SetWidth(S.scrollbar)
	track:SetPoint("TOPLEFT", region, "TOPRIGHT", inset, 0)
	track:SetPoint("BOTTOMLEFT", region, "BOTTOMRIGHT", inset, 0)
	local thumb = Theme.texture(parent, "OVERLAY", "muted")
	thumb:SetWidth(S.scrollbar)
	thumb:SetPoint("TOP", track, "TOP", 0, 0)
	local bar = { track = track, thumb = thumb, region = region }
	function bar:Update(shown, position)
		if shown >= 1 then
			self.track:Hide()
			self.thumb:Hide()
			return false
		end
		self.track:Show()
		self.thumb:Show()
		local height = self.region:GetHeight() or 0
		local thumbHeight = math.max(S.scrollbarThumb, height * shown)
		self.thumb:SetHeight(thumbHeight)
		self.thumb:ClearAllPoints()
		self.thumb:SetPoint("TOP", self.track, "TOP", 0, -(height - thumbHeight) * (position or 0))
		return true
	end
	bar:Update(1, 0)
	return bar
end

function Widgets.scrollable(holder)
	local content = CreateFrame("Frame", nil, holder)
	content:SetPoint("TOPLEFT", holder, "TOPLEFT", 0, 0)
	content:SetSize(holder:GetWidth(), holder:GetHeight())
	if type(holder.SetClipsChildren) == "function" then
		holder:SetClipsChildren(true)
	end
	local scroll = { holder = holder, content = content, offset = 0, contentHeight = holder:GetHeight() }
	scroll.bar = Widgets.scrollbar(holder, holder)
	function scroll:Range()
		return math.max(0, self.contentHeight - self.holder:GetHeight())
	end
	function scroll:ScrollTo(offset)
		self.offset = math.max(0, math.min(offset, self:Range()))
		self.content:ClearAllPoints()
		self.content:SetPoint("TOPLEFT", self.holder, "TOPLEFT", 0, self.offset)
		local range = self:Range()
		self.bar:Update(self.contentHeight > 0 and self.holder:GetHeight() / self.contentHeight or 1,
			range > 0 and self.offset / range or 0)
		return self.offset
	end
	function scroll:SetContentHeight(height)
		self.contentHeight = height
		self.content:SetHeight(math.max(height, self.holder:GetHeight()))
		return self:ScrollTo(self.offset)
	end
	holder:EnableMouseWheel(true)
	holder:SetScript("OnMouseWheel", function(_, delta)
		scroll:ScrollTo(scroll.offset - delta * Widgets.SCROLL_STEP)
	end)
	-- A holder the client resizes (an anchored canvas) keeps the range true to what shows.
	holder:SetScript("OnSizeChanged", function()
		scroll:SetContentHeight(scroll.contentHeight)
	end)
	return scroll
end

--- Pixels one notch of the wheel moves a scrollable page.
Widgets.SCROLL_STEP = 40

Widgets.List = {}
Widgets.List.__index = Widgets.List

--- A list of `rowCount` recycled rows over a pure offset. No scroll
--- template: controller ruling 3.
function Widgets.list(parent, width, rowCount, rowBuilder)
	local list = setmetatable({ items = {}, offset = 0, rows = {} }, Widgets.List)
	local frame = CreateFrame("Frame", nil, parent)
	frame:SetSize(width, rowCount * Theme.SIZES.rowHeight)
	frame:EnableMouseWheel(true)
	frame:SetScript("OnMouseWheel", function(_, delta)
		list:Scroll(-delta)
	end)
	for index = 1, rowCount do
		local row = (rowBuilder or Widgets.itemRow)(frame, width)
		row.frame:SetPoint("TOPLEFT", frame, "TOPLEFT", 0,
			-(index - 1) * Theme.SIZES.rowHeight)
		list.rows[index] = row
	end
	list.frame = frame
	-- Just outside the rows, in the page's own right padding, so no row
	-- text runs under it.
	list.bar = Widgets.scrollbar(frame, frame, Theme.SIZES.gap * 2)
	return list
end

--- `render(row, item)` fills one row; the list owns showing and hiding it.
function Widgets.List:SetRenderer(render)
	self.render = render
	return self
end

function Widgets.List:SetItems(items)
	self.items = items or {}
	self.offset = 0
	return self:Refresh()
end

function Widgets.List:Scroll(delta)
	self.offset = self.offset + delta
	return self:Refresh()
end

--- Show only the first `count` of the rows the list was built with (all
--- of them by default): a page that must share its height between two
--- lists gives one of them fewer rows without rebuilding either.
function Widgets.List:SetVisibleRows(count)
	self.visibleRows = math.max(0, math.min(count or #self.rows, #self.rows))
	self.frame:SetHeight(math.max(self.visibleRows, 1) * Theme.SIZES.rowHeight)
	return self:Refresh()
end

function Widgets.List:Refresh()
	local visible = self.visibleRows or #self.rows
	local first, last, clamped = Widgets.visibleRange(#self.items, visible, self.offset)
	self.offset = clamped
	for index, row in ipairs(self.rows) do
		local position = index + first - 1
		local item = (index <= visible and position <= last) and self.items[position] or nil
		-- Cleared before the renderer runs: a recycled row must never
		-- keep the previous item's tooltip target.
		row.itemId, row.link = nil, nil
		if item == nil then
			row.frame:Hide()
		else
			if self.render ~= nil then
				self.render(row, item)
			end
			row.frame:Show()
		end
	end
	if self.bar ~= nil then
		local total = #self.items
		local range = math.max(0, total - visible)
		self.bar:Update(total > 0 and visible / total or 1, range > 0 and self.offset / range or 0)
	end
	return self
end

--- How far right of a switch's left edge its label and hint start: the square and a gap.
function Widgets.toggleIndent()
	return Theme.SIZES.iconSize + Theme.SIZES.gap * 2
end

--- A switch row's height from its measured label and hint (fix round 2, item 3): the gap,
--- the label, the gap, the hint, the one-line rows' bottom pad (Theme.SIZES.settingsRowPad),
--- never under Theme.SIZES.settingsRow; that row where either height is unknown. Pure.
function Widgets.toggleRowHeight(labelHeight, hintHeight)
	local S = Theme.SIZES
	if type(labelHeight) ~= "number" or type(hintHeight) ~= "number" then
		return S.settingsRow
	end
	return math.max(S.settingsRow, S.gap + labelHeight + S.gap + hintHeight + S.settingsRowPad)
end

--- A font string's height as the client measured it, or nil.
local function stringHeight(region)
	if type(region.GetStringHeight) ~= "function" then
		return nil
	end
	local height = region:GetStringHeight()
	return type(height) == "number" and height or nil
end

--- A switch (L3.5 ruling 2): on, a gold square with the client's check mark tinted dark on
--- it; off, a 1 px border outline. With a `hint`, a row as tall as its label and wrapped
--- hint measure (Widgets.toggleRowHeight, at least Theme.SIZES.settingsRow; the whole row is
--- the target): the label in body, the hint muted under it, wrapped at `hintWidth`. Hover
--- and pressed come from Theme.CONTROL_STATES.toggle. `onChange` gets the new value.
function Widgets.toggle(parent, text, checked, onChange, hint, hintWidth)
	local S = Theme.SIZES
	local states = Theme.CONTROL_STATES.toggle
	local indent = Widgets.toggleIndent()
	local frame = Theme.createFrame("Button", nil, parent)
	local box = CreateFrame("Frame", nil, frame)
	box:SetSize(S.iconSize, S.iconSize)
	local fill = Theme.texture(box, "ARTWORK", states.on)
	fill:SetAllPoints(box)
	local edges = Theme.outline(box, S.border, states.edge)
	local tick = Theme.tinted(box, "OVERLAY", Theme.MEDIA.check, states.check)
	tick:SetAllPoints(box)
	local label = Widgets.label(frame, text, "body", hint ~= nil and "normal" or "small")
	local toggle = { frame = frame, box = box, fill = fill, edges = edges, tick = tick, label = label }
	if hint ~= nil then
		frame:SetSize(indent + (hintWidth or S.buttonWidth), S.settingsRow)
		box:SetPoint("TOPLEFT", frame, "TOPLEFT", 0, -S.gap)
		label:SetPoint("TOPLEFT", frame, "TOPLEFT", indent, -S.gap)
		local hintLabel = Widgets.label(frame, hint, states.hint, "small")
		hintLabel:SetPoint("TOPLEFT", label, "BOTTOMLEFT", 0, -S.gap)
		hintLabel:SetWidth(hintWidth or S.buttonWidth)
		frame:SetHeight(Widgets.toggleRowHeight(stringHeight(label), stringHeight(hintLabel)))
		toggle.hint = hintLabel
	else
		frame:SetSize(S.buttonWidth, S.rowHeight)
		box:SetPoint("LEFT", frame, "LEFT", 0, 0)
		label:SetPoint("LEFT", box, "RIGHT", S.gap, 0)
	end
	--- Paint the square for its value and the cursor: "rest", "hover" or "pressed".
	local function paint(look)
		if toggle.checked then
			local key = look == "pressed" and states.onPressed or look == "hover" and states.onHover or states.on
			Theme.paint(fill, key)
			fill:Show()
			tick:Show()
			for _, edge in ipairs(edges) do
				Theme.paint(edge, key)
			end
		else
			tick:Hide()
			if look == "pressed" then
				Theme.paint(fill, states.offPressed)
				fill:Show()
			else
				fill:Hide()
			end
			for _, edge in ipairs(edges) do
				Theme.paint(edge, look == "rest" and states.edge or states.edgeHover)
			end
		end
	end
	function toggle:SetChecked(value)
		self.checked = value and true or false
		paint(self.over and "hover" or "rest")
		return self.checked
	end
	frame:SetScript("OnEnter", function()
		toggle.over = true
		paint("hover")
	end)
	frame:SetScript("OnLeave", function()
		toggle.over = false
		paint("rest")
	end)
	Theme.pressable(frame, label, function()
		paint("pressed")
	end, function(_, over)
		toggle.over = over
		paint(over and "hover" or "rest")
	end)
	frame:SetScript("OnClick", function()
		onChange(toggle:SetChecked(not toggle.checked))
	end)
	toggle:SetChecked(checked)
	return toggle
end

ns.Widgets = Widgets
return Widgets
